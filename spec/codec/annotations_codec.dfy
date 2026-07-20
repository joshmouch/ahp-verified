// AHP Dafny client — FULL Wire.Json ⇄ AnnotationsAction codec (Annotations channel).
//
// Replicates the PROVEN root-codec template (spec/fidelity.dfy: DecodeRootAction /
// EncodeRootAction / encodeRootState + RootActionRoundTrip + _NonVacuous +
// RootUnknown_EncodeVerbatim) for the Annotations channel instead of Root.
//
//   DECODE is promoted VERBATIM (logic) from the proven reducer-replay decoder in
//   spec/replay/annotations_replay.dfy — the exact code that drives 10/10 annotations
//   fixtures through the reducer. ENCODE is its inverse (canonical wire object per
//   typed variant). Together they carry the load-bearing transport property:
//   decode ∘ encode == id on the TYPED variants (AnnotationsActionRoundTrip below).
//
//   The AnUnknown leaf is carved out HONESTLY (same as RootUnknown): its raw payload
//   could itself be a typed-shaped object, so a blanket round-trip would let decode
//   re-classify it. That leaf re-encodes VERBATIM instead (the verbatim lemmas).
//
// Field accessors (field/asStr/asArr/asBool) are REUSED from Fidelity — not
// duplicated. optStr/optBool (present-key→Some / absent-key→None) are annotations-
// specific and live here, defined ON TOP of Fidelity's accessors.
include "../annotations.dfy"
include "../fidelity.dfy"

module AnnotationsCodec {
  export
    provides DecodeAnnotationsAction, decodeAnnotationsState, encodeAnnotationsState
    provides AnnotationsActionRoundTrip, AnnotationsStateRoundTrip
    reveals EncodeAnnotationsAction
    provides encodeAnnotation, encodeEntry
    provides Annotations, Wire

  import opened Annotations
  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import F = Fidelity
  import CO = ConfluxOperators

  // ── optional-scalar accessors (present JSON key → Some, absent → None) ──
  // Built on Fidelity's shared field/asStr/asBool accessors; these encode the
  // partial `updated`/opaque-field semantics (present field → set, absent → None).
  function optStr(j: Wire.Json, k: string): Option<string> {
    if F.field(j, k).Some? then Some(F.asStr(F.field(j, k))) else None
  }
  function optBool(j: Wire.Json, k: string): Option<bool> {
    if F.field(j, k).Some? then Some(F.asBool(F.field(j, k), false)) else None
  }

  // ════════════════════════════════════════════════════════════════════════
  //  DECODE — promoted from spec/replay/annotations_replay.dfy (logic verbatim;
  //  the shared accessors are Fidelity-qualified). `range`/`_meta` are opaque
  //  Option<Wire.Json>: F.field already yields Some(value) when the key is present and
  //  None when absent — exactly the present→set / absent→preserve semantics.
  // ════════════════════════════════════════════════════════════════════════
  function decodeEntry(j: Wire.Json): Entry {
    Entry(F.asStr(F.field(j, "id")), F.asStr(F.field(j, "text")), F.field(j, "_meta"))
  }
  function decodeEntries(js: seq<Wire.Json>): seq<Entry>
  { CO.Map(decodeEntry, js) }

  function decodeAnnotation(j: Wire.Json): Annotation {
    Annotation(
      F.asStr(F.field(j, "id")),
      F.asStr(F.field(j, "turnId")),
      F.asStr(F.field(j, "resource")),
      F.field(j, "range"),                               // Option<Wire.Json>, opaque
      F.asBool(F.field(j, "resolved"), false),
      decodeEntries(F.asArr(F.field(j, "entries"))),
      F.field(j, "_meta"))                               // Option<Wire.Json>, opaque
  }
  function decodeAnnotations(js: seq<Wire.Json>): seq<Annotation>
  { CO.Map(decodeAnnotation, js) }

  function decodeAnnotationsState(j: Wire.Json): AnnotationsState {
    AnnotationsState(decodeAnnotations(F.asArr(F.field(j, "annotations"))))
  }

  function DecodeAnnotationsAction(j: Wire.Json): AnnotationsAction {
    var t := F.asStr(F.field(j, "type"));
    if t == "annotations/set" then
      Set(decodeAnnotation(if F.field(j, "annotation").Some? then F.field(j, "annotation").value else Wire.JNull))
    else if t == "annotations/removed" then
      Removed(F.asStr(F.field(j, "annotationId")))
    else if t == "annotations/entrySet" then
      EntrySet(F.asStr(F.field(j, "annotationId")),
               decodeEntry(if F.field(j, "entry").Some? then F.field(j, "entry").value else Wire.JNull))
    else if t == "annotations/entryRemoved" then
      EntryRemoved(F.asStr(F.field(j, "annotationId")), F.asStr(F.field(j, "entryId")))
    else if t == "annotations/updated" then
      Updated(F.asStr(F.field(j, "annotationId")),
              optStr(j, "turnId"), optStr(j, "resource"),
              F.field(j, "range"), optBool(j, "resolved"))
    else
      AnUnknown(j)
  }

  // ════════════════════════════════════════════════════════════════════════
  //  ENCODE — the inverse of the decode above (canonical wire object per typed
  //  variant). Opaque `range`/`_meta` re-emit under their key iff Some (matching
  //  the absent→None decode); the required scalars always emit.
  // ════════════════════════════════════════════════════════════════════════
  function encodeEntry(e: Entry): Wire.Json {
    var m0 := map["id" := Wire.JStr(e.id), "text" := Wire.JStr(e.text)];
    var m1 := if e.meta.Some? then m0["_meta" := e.meta.value] else m0;
    Wire.JObj(m1)
  }
  function encodeEntries(es: seq<Entry>): seq<Wire.Json>
  { CO.Map(encodeEntry, es) }

  function encodeAnnotation(a: Annotation): Wire.Json {
    var m0 := map["id" := Wire.JStr(a.id),
                  "turnId" := Wire.JStr(a.turnId),
                  "resource" := Wire.JStr(a.resource),
                  "resolved" := Wire.JBool(a.resolved),
                  "entries" := Wire.JArr(encodeEntries(a.entries))];
    var m1 := if a.range.Some? then m0["range" := a.range.value] else m0;
    var m2 := if a.meta.Some? then m1["_meta" := a.meta.value] else m1;
    Wire.JObj(m2)
  }
  function encodeAnnotations(anns: seq<Annotation>): seq<Wire.Json>
  { CO.Map(encodeAnnotation, anns) }

  function EncodeAnnotationsAction(a: AnnotationsAction): Wire.Json {
    match a
    case Set(ann)               => Wire.JObj(map["type" := Wire.JStr("annotations/set"), "annotation" := encodeAnnotation(ann)])
    case Removed(id)            => Wire.JObj(map["type" := Wire.JStr("annotations/removed"), "annotationId" := Wire.JStr(id)])
    case EntrySet(aid, e)       => Wire.JObj(map["type" := Wire.JStr("annotations/entrySet"), "annotationId" := Wire.JStr(aid), "entry" := encodeEntry(e)])
    case EntryRemoved(aid, eid) => Wire.JObj(map["type" := Wire.JStr("annotations/entryRemoved"), "annotationId" := Wire.JStr(aid), "entryId" := Wire.JStr(eid)])
    case Updated(aid, tid, res, rng, rslv) =>
      var m0 := map["type" := Wire.JStr("annotations/updated"), "annotationId" := Wire.JStr(aid)];
      var m1 := if tid.Some?  then m0["turnId" := Wire.JStr(tid.value)]   else m0;
      var m2 := if res.Some?  then m1["resource" := Wire.JStr(res.value)] else m1;
      var m3 := if rng.Some?  then m2["range" := rng.value]          else m2;
      var m4 := if rslv.Some? then m3["resolved" := Wire.JBool(rslv.value)] else m3;
      Wire.JObj(m4)
    case AnUnknown(raw)         => raw
  }

  // ── STATE serialization (deterministic total serializer, like encodeRootState) ──
  // Used by a cross-process transport demo to emit the channel's final state as
  // canonical JSON; convergence relies on determinism, not on a state round-trip.
  function encodeAnnotationsState(s: AnnotationsState): Wire.Json {
    Wire.JObj(map["annotations" := Wire.JArr(encodeAnnotations(s.annotations))])
  }

  // ════════════════════════════════════════════════════════════════════════
  //  ROUND-TRIP — decode ∘ encode == id on the sub-codecs, then on the action.
  //  Real + quantified (∀ over the datatype) + non-vacuous (witness lemma).
  //  Induction mirrors Fidelity's Agents→Agent→Strs chain: Annotation carries a
  //  seq<Entry>, Entry carries opaque Option<Wire.Json>.
  // ════════════════════════════════════════════════════════════════════════
  lemma EntryRoundTrip(e: Entry)
    ensures decodeEntry(encodeEntry(e)) == e
  {}

  lemma EntriesRoundTrip(es: seq<Entry>)
    ensures decodeEntries(encodeEntries(es)) == es
    decreases es
  { if |es| == 0 {} else { EntryRoundTrip(es[0]); EntriesRoundTrip(es[1..]); } }

  lemma AnnotationRoundTrip(a: Annotation)
    ensures decodeAnnotation(encodeAnnotation(a)) == a
  { EntriesRoundTrip(a.entries); }

  // The load-bearing transport property: every TYPED action encodes to a wire
  // object that decodes back to itself. (AnUnknown is excluded — its raw payload
  // may itself be a typed-shaped object, so decode would re-classify it; that leaf
  // is handled by the verbatim lemmas below, not this one.)
  // Json→domain→Json fixpoint: the codec's canonical wire form re-decodes + re-encodes to ITSELF — the
  // direction the conformance suite exercises, now machine-checked. Follows from the direction-1 round-trip.
  lemma AnnotationsWireCanonicalRoundTrip(a: AnnotationsAction)
    requires !a.AnUnknown?
    ensures EncodeAnnotationsAction(DecodeAnnotationsAction(EncodeAnnotationsAction(a))) == EncodeAnnotationsAction(a)
  { AnnotationsActionRoundTrip(a); }
  lemma AnnotationsActionRoundTrip(a: AnnotationsAction)
    requires !a.AnUnknown?
    ensures DecodeAnnotationsAction(EncodeAnnotationsAction(a)) == a
  {
    match a
    case Set(ann)                          => AnnotationRoundTrip(ann);
    case Removed(id)                       =>
    case EntrySet(aid, e)                  => EntryRoundTrip(e);
    case EntryRemoved(aid, eid)            =>
    case Updated(aid, tid, res, rng, rslv) =>
  }

  // Non-vacuity witness: the round-trip precondition is inhabited by a genuine
  // typed action (a Set carrying an annotation with an entry — it exercises the
  // encodeAnnotation → encodeEntries chain), so AnnotationsActionRoundTrip is NOT
  // vacuously true.
  lemma AnnotationsActionRoundTrip_NonVacuous()
    ensures exists a: AnnotationsAction :: !a.AnUnknown? && DecodeAnnotationsAction(EncodeAnnotationsAction(a)) == a
  {
    var e := Entry("c-1", "hello", None);
    var ann := Annotation("t-1", "turn-1", "file:///src/a.ts", None, false, [e], None);
    var a := Set(ann);
    AnnotationsActionRoundTrip(a);
    assert !a.AnUnknown? && DecodeAnnotationsAction(EncodeAnnotationsAction(a)) == a;
  }

  // ── AnUnknown leaf (honest carve-out) — re-encodes VERBATIM, exactly like
  //    Fidelity's RootUnknown. This is why the round-trip lemma above excludes it:
  //    the raw Wire.Json could be a typed-shaped object that decode would re-classify. ──
  lemma AnUnknown_EncodeVerbatim(j: Wire.Json)
    ensures EncodeAnnotationsAction(AnUnknown(j)) == j {}
  lemma AnUnknownVariant_Verbatim(j: Wire.Json)
    requires DecodeAnnotationsAction(j).AnUnknown?
    ensures EncodeAnnotationsAction(DecodeAnnotationsAction(j)) == j {}

  // ════════════════════════════════════════════════════════════════════════
  //  STATE ROUND-TRIP — decode ∘ encode == id on AnnotationsState. Completes the
  //  codec's proof story in BOTH directions: the ACTION codec already round-trips
  //  (AnnotationsActionRoundTrip), and this makes the STATE serializer
  //  encodeAnnotationsState a PROVEN-INVERTIBLE function (decodeAnnotationsState
  //  is its true inverse), not merely a deterministic one.
  //
  //  Induction mirrors Fidelity's Agents→Agent→Strs chain: AnnotationsState
  //  carries a seq<Annotation>; decodeAnnotations/encodeAnnotations recurse over
  //  it, bottoming out at the proven AnnotationRoundTrip (which itself leans on
  //  EntriesRoundTrip → EntryRoundTrip).
  // ════════════════════════════════════════════════════════════════════════

  // seq<Annotation> round-trip (mirrors Fidelity.AgentsRoundTrip over decodeAgents).
  lemma AnnotationsRoundTrip(anns: seq<Annotation>)
    ensures decodeAnnotations(encodeAnnotations(anns)) == anns
    decreases anns
  {
    if |anns| == 0 {} else { AnnotationRoundTrip(anns[0]); AnnotationsRoundTrip(anns[1..]); }
  }

  // The state round-trip is UNCONDITIONAL — NO WfAnnotationsState carve-out is
  // needed. AnnotationsState's only field is a seq<Annotation>, and
  // AnnotationRoundTrip holds with no precondition: the opaque range/_meta
  // (Option<Wire.Json>) re-emit under their key iff Some and decode reads them
  // VERBATIM via F.field, so a Some(Wire.JNull) survives intact (no Option-Wire.JNull
  // collapse, unlike the Updated action's optStr/optBool scalars). Every scalar
  // is required and always emitted. Real + quantified (∀ s) + non-vacuous.
  lemma AnnotationsStateRoundTrip(s: AnnotationsState)
    ensures decodeAnnotationsState(encodeAnnotationsState(s)) == s
  {
    AnnotationsRoundTrip(s.annotations);
  }

  // Non-vacuity witness: a genuine NON-EMPTY state (one annotation carrying an
  // entry and an opaque range) round-trips, so AnnotationsStateRoundTrip is NOT
  // vacuously true — it exercises the encodeAnnotations → encodeAnnotation →
  // encodeEntries chain end to end.
  lemma AnnotationsStateRoundTrip_NonVacuous()
    ensures exists s: AnnotationsState ::
      |s.annotations| > 0 && decodeAnnotationsState(encodeAnnotationsState(s)) == s
  {
    var e := Entry("c-1", "hello", None);
    var ann := Annotation("t-1", "turn-1", "file:///src/a.ts",
                          Some(Wire.JObj(map["start" := Wire.JNull])), false, [e], None);
    var s := AnnotationsState([ann]);
    AnnotationsStateRoundTrip(s);
    assert |s.annotations| > 0 && decodeAnnotationsState(encodeAnnotationsState(s)) == s;
  }

  ghost predicate AnnotationsCodecWf(a: AnnotationsAction) { !a.AnUnknown? }
  function AnnotationsCodecC(): Wire.Codec<AnnotationsAction> {
    Wire.Codec(EncodeAnnotationsAction, (j: Wire.Json) => Wire.Some(DecodeAnnotationsAction(j)))
  }
  lemma AnnotationsCodecRoundTripsWhere()
    ensures forall a :: AnnotationsCodecWf(a) ==> AnnotationsCodecC().decode(AnnotationsCodecC().encode(a)) == Wire.Some(a)
  {
    forall a | AnnotationsCodecWf(a)
      ensures AnnotationsCodecC().decode(AnnotationsCodecC().encode(a)) == Wire.Some(a)
    { AnnotationsActionRoundTrip(a); }
  }
}
