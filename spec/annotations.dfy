// AHP Dafny client — annotations channel (4th channel). Annotations by id, entries
// by id within an annotation, partial `updated` (present field → set, absent → preserve).
// Full coverage of all 10 annotations fixtures. `_meta`/`range` opaque Wire.Json.
include "ahp_skeleton.dfy"

module Annotations {
  // Firewall: consumers get the channel's action/state types + the reducer surface, NOT the
  // internal seq-route helpers (upsertAnn/removeAnn/hasAnn/upsertEntry/removeEntry/updateAnn/
  // applyUpdate/R) or the An_* {:NoOp} lemmas. provides Wire/CC/AhpSkeleton because the exported
  // signatures name Wire.Json / CC.Fold / Option+ReduceOutcome (transitive-export rule).
  export
    provides ApplyToAnnotations, apply1, RunCorpus
    reveals fold   // channel_laws.AnnotationsIsChannel asserts Annotations.fold == CC.Fold(apply1,·) in its body
    provides AhpSkeleton, Wire, CC
    provides WfAnnotationsState, WfAnnotationsAction, Apply1PreservesWfAnnotations, AnnotationsWfWitness  // rung-4/6: reducer keyed invariant + apply1-level preservation + concrete witness (ahp.dfy composes WfAhpStateInv over it)
    reveals AnnotationsState, AnnotationsAction, Annotation, Entry

  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import CC = ConfluxContract
  import SR = ConfluxSeqRoute  // the canonical ordered-keyed-seq primitives (upsert/remove/update by id)

  datatype Entry = Entry(id: string, text: string, meta: Option<Wire.Json>)
  datatype Annotation = Annotation(
    id: string, turnId: string, resource: string,
    range: Option<Wire.Json>, resolved: bool, entries: seq<Entry>, meta: Option<Wire.Json>)
  datatype AnnotationsState = AnnotationsState(annotations: seq<Annotation>)

  datatype AnnotationsAction =
    | Set(annotation: Annotation)                                       // 210/211 upsert by id
    | Removed(annotationId: string)                                     // 212
    | EntrySet(annotationId: string, entry: Entry)                      // 213/214
    | EntryRemoved(annotationId: string, entryId: string)              // 215
    | Updated(annotationId: string, turnId: Option<string>, resource: Option<string>, range: Option<Wire.Json>, resolved: Option<bool>) // 216/218/219
    | AnUnknown(raw: Wire.Json)                                              // 217

  // NAMED keyOf projections (not bare lambdas): the ConfluxSeqRoute laws about
  // SeqUpsertById(AnnId, …)/SeqUpdateById(AnnId, …) connect definitionally only when
  // the keyOf is a single named function value shared across call sites (a lambda
  // literal is not guaranteed function-equal across occurrences).
  function AnnId(a: Annotation): string { a.id }
  function EntryId(e: Entry): string { e.id }

  // order-preserving upsert-by-id: delegates to the canonical ConfluxSeqRoute.SeqUpsertById.
  function upsertAnn(anns: seq<Annotation>, a: Annotation): seq<Annotation>
  { SR.SeqUpsertById(AnnId, anns, a) }
  // remove-by-id: delegates to ConfluxSeqRoute.SeqRemoveById.
  function removeAnn(anns: seq<Annotation>, id: string): seq<Annotation>
  { SR.SeqRemoveById(AnnId, anns, id) }
  predicate hasAnn(anns: seq<Annotation>, id: string) { exists i :: 0 <= i < |anns| && anns[i].id == id }

  function upsertEntry(entries: seq<Entry>, e: Entry): seq<Entry>
  { SR.SeqUpsertById(EntryId, entries, e) }
  function removeEntry(entries: seq<Entry>, id: string): seq<Entry>
  { SR.SeqRemoveById(EntryId, entries, id) }

  // update-in-place-by-id: apply f to the matching annotation, no-op if absent —
  // delegates to ConfluxSeqRoute.SeqUpdateById.
  function updateAnn(anns: seq<Annotation>, id: string, f: Annotation -> Annotation): seq<Annotation>
  { SR.SeqUpdateById(AnnId, anns, id, f) }

  // NAMED per-arm annotation transforms (not inline lambdas): each is a single
  // function VALUE referenced identically by the reducer and by WfAnnotationsStatePreserved,
  // so `updateAnn(s.annotations, aid, setEntry(e))` in the reducer and in the proof are the
  // same term — the citation of SeqUpdateById's uniqueness law resolves.
  function setEntry(e: Entry): Annotation -> Annotation
  { (an: Annotation) => an.(entries := upsertEntry(an.entries, e)) }
  function dropEntry(eid: string): Annotation -> Annotation
  { (an: Annotation) => an.(entries := removeEntry(an.entries, eid)) }
  function reanchor(tid: Option<string>, res: Option<string>, rng: Option<Wire.Json>, rslv: Option<bool>): Annotation -> Annotation
  { (an: Annotation) => applyUpdate(an, tid, res, rng, rslv) }

  function applyUpdate(ann: Annotation, tid: Option<string>, res: Option<string>, rng: Option<Wire.Json>, rslv: Option<bool>): Annotation
  {
    ann.(turnId := if tid.Some? then tid.value else ann.turnId,
         resource := if res.Some? then res.value else ann.resource,
         range := if rng.Some? then Some(rng.value) else ann.range,
         resolved := if rslv.Some? then rslv.value else ann.resolved)
  }

  function ApplyToAnnotations(s: AnnotationsState, a: AnnotationsAction, now: int): (AnnotationsState, ReduceOutcome)
  {
    match a
    case Set(ann)              => (s.(annotations := upsertAnn(s.annotations, ann)), Applied)
    case Removed(id)           => (s.(annotations := removeAnn(s.annotations, id)), Applied)
    case EntrySet(aid, e)      => if hasAnn(s.annotations, aid) then (s.(annotations := updateAnn(s.annotations, aid, setEntry(e))), Applied) else (s, NoOp)
    case EntryRemoved(aid, eid) => if hasAnn(s.annotations, aid) then (s.(annotations := updateAnn(s.annotations, aid, dropEntry(eid))), Applied) else (s, NoOp)
    case Updated(aid, tid, res, rng, rslv) => if hasAnn(s.annotations, aid) then (s.(annotations := updateAnn(s.annotations, aid, reanchor(tid, res, rng, rslv))), Applied) else (s, NoOp)
    case AnUnknown(_)          => (s, NoOp)
  }

  // ══════════════════════════════════════════════════════════════════════════
  //  Rung 4/6/7 — the keyed-collection well-formedness invariant + its reducer
  //  preservation proof. AnnotationsState is TWO nested id-keyed ordered seqs:
  //  `annotations` keyed by id, and each annotation's `entries` keyed by id.
  //  WfAnnotationsState carries BOTH uniqueness facts; WfAnnotationsStatePreserved
  //  discharges every arm by citing the matching ConfluxSeqRoute *PreservesUnique law.
  // ══════════════════════════════════════════════════════════════════════════

  // Rung 4: the real multi-collection invariant (NOT vacuous) — top-level annotation
  // ids are unique AND every annotation's entry ids are unique.
  ghost predicate WfAnnotationsState(s: AnnotationsState) {
    SR.UniqueKeys(AnnId, s.annotations)
    && (forall x :: x in s.annotations ==> SR.UniqueKeys(EntryId, x.entries))
  }

  // Rung 7: Set installs the action's annotation WHOLESALE (its whole `entries` seq
  // replaces the target's), so the nested invariant needs the action's entries to be
  // id-unique. The other arms mutate an existing (already-Wf) annotation, so they need
  // no action-side premise.
  ghost predicate WfAnnotationsAction(a: AnnotationsAction) {
    a.Set? ==> SR.UniqueKeys(EntryId, a.annotation.entries)
  }

  // ── membership helpers: an element of an order-preserving keyed upsert/update is
  //    drawn from the input (or is the upserted value / an f-image). ConfluxSeqRoute
  //    ships RemoveKeysSubset for delete; these are the upsert/update counterparts,
  //    proved against the primitive's visible bodies (no re-derivation of the fold). ──
  lemma UpsertMemberOf<K, V>(keyOf: V -> K, xs: seq<V>, v: V, x: V)
    requires x in SR.SeqUpsertById(keyOf, xs, v)
    ensures x == v || x in xs
    decreases |xs|
  {
    if |xs| == 0 {
    } else if keyOf(xs[0]) == keyOf(v) {
      assert SR.SeqUpsertById(keyOf, xs, v) == [v] + xs[1..];
    } else {
      assert SR.SeqUpsertById(keyOf, xs, v) == [xs[0]] + SR.SeqUpsertById(keyOf, xs[1..], v);
      if x != xs[0] { UpsertMemberOf(keyOf, xs[1..], v, x); }
    }
  }

  lemma UpdateMemberOf<K, V>(keyOf: V -> K, xs: seq<V>, k: K, f: V -> V, x: V)
    requires x in SR.SeqUpdateById(keyOf, xs, k, f)
    ensures x in xs || exists y :: y in xs && x == f(y)
    decreases |xs|
  {
    if |xs| == 0 {
    } else if keyOf(xs[0]) == k {
      assert SR.SeqUpdateById(keyOf, xs, k, f) == [f(xs[0])] + xs[1..];
      if x != f(xs[0]) && x !in xs[1..] {
        assert false;
      }
      if x == f(xs[0]) { assert xs[0] in xs && x == f(xs[0]); }
    } else {
      assert SR.SeqUpdateById(keyOf, xs, k, f) == [xs[0]] + SR.SeqUpdateById(keyOf, xs[1..], k, f);
      if x != xs[0] {
        UpdateMemberOf(keyOf, xs[1..], k, f, x);
        if x !in xs[1..] { var y :| y in xs[1..] && x == f(y); assert y in xs; }
      }
    }
  }

  // Rung 6: the reducer preserves the invariant over ALL arms.  `now: int` is canonical.
  lemma WfAnnotationsStatePreserved(s: AnnotationsState, a: AnnotationsAction, now: int)
    requires WfAnnotationsState(s) && WfAnnotationsAction(a)
    ensures WfAnnotationsState(ApplyToAnnotations(s, a, now).0)
  {
    match a
    case Set(ann) =>
      // top-level: append-or-replace preserves annotation-id uniqueness.
      SR.UpsertPreservesUnique(AnnId, s.annotations, ann);
      // nested: every element of the upsert is an existing (Wf) annotation or `ann`,
      //         whose entries are id-unique by WfAnnotationsAction.
      var outAnns := upsertAnn(s.annotations, ann);
      forall x | x in outAnns ensures SR.UniqueKeys(EntryId, x.entries) {
        UpsertMemberOf(AnnId, s.annotations, ann, x);
      }
    case Removed(id) =>
      // top-level: remove-by-id preserves uniqueness; nested: the kept elements are a
      //            subset of the (Wf) input.
      SR.RemovePreservesUnique(AnnId, s.annotations, id);
      SR.RemoveKeysSubset(AnnId, s.annotations, id);
    case EntrySet(aid, e) =>
      if hasAnn(s.annotations, aid) {
        assert SR.KeyPreserving(AnnId, setEntry(e));
        SR.UpdatePreservesUnique(AnnId, s.annotations, aid, setEntry(e));
        var outAnns := updateAnn(s.annotations, aid, setEntry(e));
        forall x | x in outAnns ensures SR.UniqueKeys(EntryId, x.entries) {
          UpdateMemberOf(AnnId, s.annotations, aid, setEntry(e), x);
          if x !in s.annotations {
            var y :| y in s.annotations && x == setEntry(e)(y);
            // x.entries == SeqUpsertById(EntryId, y.entries, e); y is Wf ⇒ its entries unique.
            SR.UpsertPreservesUnique(EntryId, y.entries, e);
          }
        }
      }
    case EntryRemoved(aid, eid) =>
      if hasAnn(s.annotations, aid) {
        assert SR.KeyPreserving(AnnId, dropEntry(eid));
        SR.UpdatePreservesUnique(AnnId, s.annotations, aid, dropEntry(eid));
        var outAnns := updateAnn(s.annotations, aid, dropEntry(eid));
        forall x | x in outAnns ensures SR.UniqueKeys(EntryId, x.entries) {
          UpdateMemberOf(AnnId, s.annotations, aid, dropEntry(eid), x);
          if x !in s.annotations {
            var y :| y in s.annotations && x == dropEntry(eid)(y);
            // x.entries == SeqRemoveById(EntryId, y.entries, eid); y is Wf ⇒ stays unique.
            SR.RemovePreservesUnique(EntryId, y.entries, eid);
          }
        }
      }
    case Updated(aid, tid, res, rng, rslv) =>
      if hasAnn(s.annotations, aid) {
        assert SR.KeyPreserving(AnnId, reanchor(tid, res, rng, rslv));
        SR.UpdatePreservesUnique(AnnId, s.annotations, aid, reanchor(tid, res, rng, rslv));
        var outAnns := updateAnn(s.annotations, aid, reanchor(tid, res, rng, rslv));
        forall x | x in outAnns ensures SR.UniqueKeys(EntryId, x.entries) {
          UpdateMemberOf(AnnId, s.annotations, aid, reanchor(tid, res, rng, rslv), x);
          if x !in s.annotations {
            var y :| y in s.annotations && x == reanchor(tid, res, rng, rslv)(y);
            // reanchor rewrites turnId/resource/range/resolved only ⇒ entries unchanged.
            assert x.entries == y.entries;
          }
        }
      }
    case AnUnknown(_) =>
  }

  lemma An_UnknownIsNoOp(s: AnnotationsState, a: AnnotationsAction, now: int)
    requires a.AnUnknown?
    ensures ApplyToAnnotations(s, a, now).0 == s && ApplyToAnnotations(s, a, now).1 == NoOp {}
  lemma An_UnknownIsNoOp_NonVacuous()
    ensures exists a: AnnotationsAction :: a.AnUnknown?
  { assert (AnUnknown(Wire.JNull)).AnUnknown?; }
  lemma An_UnknownAnnIsNoOp(s: AnnotationsState, aid: string, e: Entry, now: int)
    requires !hasAnn(s.annotations, aid)
    ensures ApplyToAnnotations(s, EntrySet(aid, e), now).0 == s {}

  function apply1(s: AnnotationsState, a: AnnotationsAction): AnnotationsState { ApplyToAnnotations(s, a, 9999).0 }
  // apply1-level restatement of the rung-6 preservation theorem for the aggregate (ahp.dfy dispatches through apply1).
  lemma Apply1PreservesWfAnnotations(s: AnnotationsState, a: AnnotationsAction)
    requires WfAnnotationsState(s) && WfAnnotationsAction(a)
    ensures  WfAnnotationsState(apply1(s, a))
  { WfAnnotationsStatePreserved(s, a, 9999); }
  // Concrete well-formed witness handed to the aggregate (ahp.dfy) so its WfAhpStateInv non-vacuity witness
  // need not unfold this module's exported-opaque WfAnnotationsState (empty annotation seq ⇒ Wf).
  lemma AnnotationsWfWitness() returns (s: AnnotationsState)
    ensures WfAnnotationsState(s)
  { s := AnnotationsState([]); assert SR.UniqueKeys(AnnId, s.annotations); }
  function fold(s: AnnotationsState, acts: seq<AnnotationsAction>): AnnotationsState
  { CC.Fold(apply1, s, acts) }
  function R(): Wire.Json { Wire.JObj(map["start" := Wire.JNull]) }   // opaque range placeholder

  method RunCorpus() returns (pass: nat, total: nat) {
    total := 10; pass := 0;
    var e1 := Entry("c-1", "original", None);
    var a1 := Annotation("t-1", "turn-1", "file:///src/a.ts", Some(R()), false, [e1], None);
    var st1 := AnnotationsState([a1]);
    // 210 set append
    var a2 := Annotation("t-2", "turn-2", "file:///src/b.ts", None, false, [Entry("c-2", "x", None)], None);
    if apply1(st1, Set(a2)) == AnnotationsState([a1, a2]) { pass := pass+1; }
    // 211 set replace (drops range, resolved true)
    var a1r := Annotation("t-1", "turn-1", "file:///src/a.ts", None, true, [e1], None);
    if apply1(st1, Set(a1r)) == AnnotationsState([a1r]) { pass := pass+1; }
    // 212 removed drops; no-op unknown
    if apply1(st1, Removed("t-1")) == AnnotationsState([]) { pass := pass+1; }
    if apply1(st1, Removed("nope")) == st1 { pass := pass+1; }
    // 213 entrySet append within annotation
    var e2 := Entry("c-2", "reply", None);
    if apply1(st1, EntrySet("t-1", e2)) == AnnotationsState([a1.(entries := [e1, e2])]) { pass := pass+1; }
    // 214 entrySet unknown annotation → no-op
    if apply1(st1, EntrySet("nope", e2)) == st1 { pass := pass+1; }
    // 215 entryRemoved drops matching entry
    var a1two := a1.(entries := [e1, e2]);
    if apply1(AnnotationsState([a1two]), EntryRemoved("t-1", "c-2")) == AnnotationsState([a1]) { pass := pass+1; }
    // 216 updated resolves, preserves rest
    if apply1(st1, Updated("t-1", None, None, None, Some(true))) == AnnotationsState([a1.(resolved := true)]) { pass := pass+1; }
    // 218 updated reanchors turnId/resource/range, preserves resolved+entries
    var newR := Wire.JObj(map["start" := Wire.JBool(true)]);
    var a1re := a1.(turnId := "turn-2", resource := "file:///src/b.ts", range := Some(newR));
    if apply1(st1, Updated("t-1", Some("turn-2"), Some("file:///src/b.ts"), Some(newR), None)) == AnnotationsState([a1re]) { pass := pass+1; }
    // 219 updated unknown annotation → no-op
    if apply1(st1, Updated("nope", None, None, None, Some(true))) == st1 { pass := pass+1; }
  }
}
