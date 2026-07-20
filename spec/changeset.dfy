// AHP Dafny client — changeset channel (3rd channel; real list-by-id reducer logic).
// Full coverage of all 15 changeset fixtures. Files/operations upsert-by-id, reviewed
// flags, error attach/clear. `edit`/`error` kept opaque (reducer never inspects them).
include "ahp_skeleton.dfy"

module Changeset {
  // Firewall: consumers get the channel's action/state datatypes + the reducer surface
  // (ApplyToChangeset/apply1/fold/RunCorpus), NOT the internal seq-merge helpers
  // (upsertFile/removeFile/setReviewed/updateOp/hasOp/anyNeedsReviewed), the E() edit
  // placeholder, or the Cs_* no-op lemmas. provides Wire/CC/AhpSkeleton because the exported
  // signatures name Wire.Json / CC.Fold / Option+ReduceOutcome (transitive-export rule).
  export
    provides ApplyToChangeset, apply1, RunCorpus
    reveals fold   // channel_laws.ChangesetIsChannel asserts Changeset.fold == CC.Fold(apply1,·) in its body
    provides AhpSkeleton, Wire, CC
    provides WfChangesetState, WfChangesetAction, Apply1PreservesWfChangeset, ChangesetWfWitness  // rung-4/6: reducer keyed invariant + apply1-level preservation + concrete witness (ahp.dfy composes WfAhpStateInv over it)
    reveals ChangesetFile, ChangesetOperation, ChangesetState, ChangesetAction

  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import Ops = ConfluxOperators
  import CC = ConfluxContract
  import SR = ConfluxSeqRoute  // the canonical ordered-keyed-seq primitives (upsert/remove/update by id)

  datatype ChangesetFile = ChangesetFile(id: string, reviewed: Option<bool>, edit: Wire.Json)
  datatype ChangesetOperation = ChangesetOperation(id: string, label_: string, scopes: seq<string>, status: string, error: Option<Wire.Json>)
  datatype ChangesetState = ChangesetState(
    status: string,
    files: seq<ChangesetFile>,
    operations: Option<seq<ChangesetOperation>>,
    error: Option<Wire.Json>)

  datatype ChangesetAction =
    | StatusChanged(status: string, error: Option<Wire.Json>)                  // 136/137
    | FileSet(file: ChangesetFile)                                        // 138/139
    | FileRemoved(fileId: string)                                         // 140
    | OperationsChanged(operations: Option<seq<ChangesetOperation>>)      // 141
    | Cleared                                                             // 142
    | OperationStatusChanged(operationId: string, status: string, error: Option<Wire.Json>)  // 155/156/157
    | ContentChanged(files: Option<seq<ChangesetFile>>, operations: Option<seq<ChangesetOperation>>, error: Option<Wire.Json>) // 158/159
    | FilesReviewedChanged(fileIds: seq<string>, reviewed: bool)          // 160/161
    | CsUnknown(raw: Wire.Json)                                                // 144

  // ── keyed-collection key functions (ONE shared value each) ──
  // The two id-keyed seqs (files, operations) route through the ConfluxSeqRoute keyed algebra by these keyOf
  // values. Sharing ONE const per key (instead of an inline lambda per call site) makes the helper bodies and
  // the WfChangesetState / *PreservesUnique law citations reference the SAME function value, so the kernel
  // uniqueness laws' conclusions line up with the helper results syntactically (no lambda-identity gap).
  const fileKey: ChangesetFile -> string := (x: ChangesetFile) => x.id
  const opKey: ChangesetOperation -> string := (x: ChangesetOperation) => x.id
  // The status/error patch applied in-place to the matching operation. Hoisted to a named function so
  // updateOp and the UpdatePreservesUnique citation both name `opSet(st, err)` — the SAME determinate value.
  function opSet(st: string, err: Option<Wire.Json>): ChangesetOperation -> ChangesetOperation
  { (e: ChangesetOperation) => e.(status := st, error := err) }

  // ── helpers (pure, total) ──
  // order-preserving upsert-by-id / remove-by-id: delegate to the canonical ConfluxSeqRoute primitives.
  function upsertFile(files: seq<ChangesetFile>, f: ChangesetFile): seq<ChangesetFile>
  { SR.SeqUpsertById(fileKey, files, f) }
  function removeFile(files: seq<ChangesetFile>, id: string): seq<ChangesetFile>
  { SR.SeqRemoveById(fileKey, files, id) }
  function setReviewed(files: seq<ChangesetFile>, ids: seq<string>, rev: bool): seq<ChangesetFile>
  {
    Ops.Map((f: ChangesetFile) => if f.id in ids then f.(reviewed := Some(rev)) else f, files)
  }
  // update-in-place-by-id: set status/error on the matching operation, no-op if absent —
  // delegates to ConfluxSeqRoute.SeqUpdateById.
  function updateOp(ops: seq<ChangesetOperation>, id: string, status: string, error: Option<Wire.Json>): seq<ChangesetOperation>
  { SR.SeqUpdateById(opKey, ops, id, opSet(status, error)) }
  predicate hasOp(ops: seq<ChangesetOperation>, id: string) {
    exists i :: 0 <= i < |ops| && ops[i].id == id
  }
  predicate anyNeedsReviewed(files: seq<ChangesetFile>, ids: seq<string>, rev: bool) {
    exists i :: 0 <= i < |files| && files[i].id in ids && files[i].reviewed != Some(rev)
  }

  // ── rung 4: the invariant-bearing well-formed state ──
  // ChangesetState carries two id-keyed collections; well-formedness is that each holds DISTINCT ids (the
  // keyed-collection uniqueness invariant, docs/shapes/07). `operations` is optional, so its uniqueness is
  // guarded by presence. (error/edit are presence-keyed opaque Json — no invariant to carry.)
  ghost predicate WfChangesetState(s: ChangesetState) {
    SR.UniqueKeys(fileKey, s.files) &&
    (s.operations.Some? ==> SR.UniqueKeys(opKey, s.operations.value))
  }

  // ── rung 7 premise: whole-seq-replacement arms must carry incoming uniqueness ──
  // The two arms that INSTALL a caller-supplied seq verbatim (OperationsChanged replaces operations;
  // ContentChanged replaces files and/or operations) cannot preserve UniqueKeys unless the incoming seq is
  // itself id-unique — the keyed-upsert/remove/update arms need no premise (the kernel laws discharge them).
  ghost predicate WfChangesetAction(a: ChangesetAction) {
    match a
    case OperationsChanged(ops) => ops.Some? ==> SR.UniqueKeys(opKey, ops.value)
    case ContentChanged(fs, ops, err) =>
      (fs.Some?  ==> SR.UniqueKeys(fileKey, fs.value)) &&
      (ops.Some? ==> SR.UniqueKeys(opKey, ops.value))
    case _ => true
  }

  function ApplyToChangeset(s: ChangesetState, a: ChangesetAction, now: int): (ChangesetState, ReduceOutcome)
  {
    match a
    case StatusChanged(st, err)   => (s.(status := st, error := err), Applied)
    case FileSet(f)               => (s.(files := upsertFile(s.files, f)), Applied)
    case FileRemoved(id)          => (s.(files := removeFile(s.files, id)), Applied)
    case OperationsChanged(ops)   => (s.(operations := ops), Applied)
    case Cleared                  => (s.(files := []), if |s.files| == 0 then NoOp else Applied)
    case OperationStatusChanged(id, st, err) =>
      (match s.operations
       case None => (s, NoOp)
       case Some(ops) => if hasOp(ops, id) then (s.(operations := Some(updateOp(ops, id, st, err))), Applied) else (s, NoOp))
    case ContentChanged(fs, ops, err) =>
      var s1 := if fs.Some? then s.(files := fs.value) else s;
      var s2 := if ops.Some? then s1.(operations := ops) else s1;
      var s3 := if err.Some? then s2.(error := err) else s2;
      (s3, Applied)
    case FilesReviewedChanged(ids, rev) =>
      if anyNeedsReviewed(s.files, ids, rev) then (s.(files := setReviewed(s.files, ids, rev)), Applied) else (s, NoOp)
    case CsUnknown(_) => (s, NoOp)
  }

  // ── rung 6: the reducer PRESERVES the keyed-uniqueness invariant on every arm ──
  // For every well-formed state and (uniqueness-carrying) action, the reduced state is still well-formed.
  // Each keyed-mutation arm is discharged by CITING the matching ConfluxSeqRoute law; the whole-seq-replacement
  // arms lean on the WfChangesetAction premise; the Ops.Map (filesReviewed) arm is discharged inline from the
  // revealed Map body + concrete id-preservation (the generic SR.MapPreservesUnique law is not yet packaged in
  // the consumed runtime .doo — see report; the arm is nonetheless closed, not assumed).
  lemma WfChangesetStatePreserved(s: ChangesetState, a: ChangesetAction, now: int)
    requires WfChangesetState(s) && WfChangesetAction(a)
    ensures WfChangesetState(ApplyToChangeset(s, a, now).0)
  {
    match a
    case StatusChanged(st, err) =>
      // files + operations untouched — invariant carried straight through.
    case FileSet(f) =>
      SR.UpsertPreservesUnique(fileKey, s.files, f);        // files := SeqUpsertById(fileKey, files, f)
    case FileRemoved(id) =>
      SR.RemovePreservesUnique(fileKey, s.files, id);       // files := SeqRemoveById(fileKey, files, id)
    case OperationsChanged(ops) =>
      // operations := ops; WfChangesetAction supplies UniqueKeys(opKey, ops.value) when Some.
    case Cleared =>
      // files := [] — vacuously unique; operations untouched.
    case OperationStatusChanged(id, st, err) =>
      if s.operations.Some? && hasOp(s.operations.value, id) {
        // operations := Some(SeqUpdateById(opKey, ops, id, opSet(st, err))); opSet fixes id ⇒ key-preserving.
        SR.UpdatePreservesUnique(opKey, s.operations.value, id, opSet(st, err));
      }
    case ContentChanged(fs, ops, err) =>
      // files/operations replaced (when Some) by WfChangesetAction-unique seqs, else carried from WfChangesetState.
    case FilesReviewedChanged(ids, rev) =>
      if anyNeedsReviewed(s.files, ids, rev) {
        var mapped := setReviewed(s.files, ids, rev);        // == Ops.Map(reviewMark, s.files), id-preserving
        assert |mapped| == |s.files|;
        forall i, j | 0 <= i < |mapped| && 0 <= j < |mapped| && fileKey(mapped[i]) == fileKey(mapped[j])
          ensures i == j
        {
          // Map body: mapped[k] == reviewMark(s.files[k]); reviewMark fixes id ⇒ fileKey(mapped[k]) == fileKey(s.files[k]).
          assert fileKey(mapped[i]) == fileKey(s.files[i]);
          assert fileKey(mapped[j]) == fileKey(s.files[j]);
          // UniqueKeys(fileKey, s.files) (from WfChangesetState) then forces i == j.
        }
        assert SR.UniqueKeys(fileKey, mapped);
      }
    case CsUnknown(raw) =>
  }

  // ── universal no-op lemmas + non-vacuity ──
  lemma Cs_UnknownIsNoOp(s: ChangesetState, a: ChangesetAction, now: int)
    requires a.CsUnknown?
    ensures ApplyToChangeset(s, a, now).0 == s && ApplyToChangeset(s, a, now).1 == NoOp {}
  lemma Cs_UnknownIsNoOp_NonVacuous()
    ensures exists a: ChangesetAction :: a.CsUnknown?
  { assert (CsUnknown(Wire.JNull)).CsUnknown?; }
  // wrong-precondition no-op: operationStatusChanged on unknown op id.
  lemma Cs_UnknownOpIsNoOp(s: ChangesetState, id: string, st: string, err: Option<Wire.Json>, now: int)
    requires s.operations.Some? && !hasOp(s.operations.value, id)
    ensures ApplyToChangeset(s, OperationStatusChanged(id, st, err), now).0 == s {}
  // Non-vacuity for the wrong-precondition no-op above. Its precondition is a
  // CONJUNCTION — operations present AND the id absent — so the `.Some?` pin
  // alone is not a witness; a present-but-empty operation list discharges both
  // at once, for any id.
  lemma Cs_UnknownOpIsNoOp_NonVacuous(id: string)
    ensures exists s: ChangesetState :: s.operations.Some? && !hasOp(s.operations.value, id)
  {
    var s := ChangesetState("idle", [], Some([]), None);
    assert s.operations.Some?;
    assert !hasOp(s.operations.value, id);
  }

  function apply1(s: ChangesetState, a: ChangesetAction): ChangesetState { ApplyToChangeset(s, a, 9999).0 }
  // apply1-level restatement of the rung-6 preservation theorem for the aggregate (ahp.dfy dispatches through apply1).
  lemma Apply1PreservesWfChangeset(s: ChangesetState, a: ChangesetAction)
    requires WfChangesetState(s) && WfChangesetAction(a)
    ensures  WfChangesetState(apply1(s, a))
  { WfChangesetStatePreserved(s, a, 9999); }
  // Concrete well-formed witness handed to the aggregate (ahp.dfy) so its WfAhpStateInv non-vacuity witness
  // need not unfold this module's exported-opaque WfChangesetState (empty files, absent operations ⇒ Wf).
  lemma ChangesetWfWitness() returns (s: ChangesetState)
    ensures WfChangesetState(s)
  { s := ChangesetState("idle", [], None, None); assert SR.UniqueKeys(fileKey, s.files); }
  function fold(s: ChangesetState, acts: seq<ChangesetAction>): ChangesetState
  { CC.Fold(apply1, s, acts) }
  function E(): Wire.Json { Wire.JObj(map["diff" := Wire.JNull]) }   // opaque edit placeholder (same in initial+expected)

  method RunCorpus() returns (pass: nat, total: nat) {
    total := 15; pass := 0;
    var fa := ChangesetFile("file:///src/a.ts", None, E());
    var fb := ChangesetFile("file:///src/b.ts", None, E());
    // 136 statusChanged to ready clears error
    if apply1(ChangesetState("computing", [], None, Some(Wire.JStr("boom"))), StatusChanged("ready", None)) == ChangesetState("ready", [], None, None) { pass := pass+1; }
    // 137 statusChanged error attaches
    if apply1(ChangesetState("ready", [], None, None), StatusChanged("error", Some(Wire.JStr("x")))) == ChangesetState("error", [], None, Some(Wire.JStr("x"))) { pass := pass+1; }
    // 138 fileSet append (unknown id)
    if apply1(ChangesetState("ready", [fa], None, None), FileSet(fb)) == ChangesetState("ready", [fa, fb], None, None) { pass := pass+1; }
    // 139 fileSet replace (id match)
    var fbv2 := ChangesetFile("file:///src/b.ts", Some(true), E());
    if apply1(ChangesetState("ready", [fa, fb], None, None), FileSet(fbv2)) == ChangesetState("ready", [fa, fbv2], None, None) { pass := pass+1; }
    // 140 fileRemoved drops match; and no-op-for-unknown checked by removeFile semantics
    if apply1(ChangesetState("ready", [fa, fb], None, None), FileRemoved("file:///src/a.ts")) == ChangesetState("ready", [fb], None, None) { pass := pass+1; }
    if apply1(ChangesetState("ready", [fa], None, None), FileRemoved("nope")) == ChangesetState("ready", [fa], None, None) { pass := pass+1; }
    // 141 operationsChanged set + clear
    var op := ChangesetOperation("create-pr", "Create Pull Request", ["changeset"], "idle", None);
    if apply1(ChangesetState("ready", [], None, None), OperationsChanged(Some([op]))) == ChangesetState("ready", [], Some([op]), None) { pass := pass+1; }
    if apply1(ChangesetState("ready", [], Some([op]), None), OperationsChanged(None)) == ChangesetState("ready", [], None, None) { pass := pass+1; }
    // 142 cleared empties; no-op on already-empty
    if apply1(ChangesetState("ready", [fa], None, None), Cleared).files == [] { pass := pass+1; }
    // 144 unknown no-op
    if apply1(ChangesetState("ready", [], None, None), CsUnknown(Wire.JObj(map["type" := Wire.JStr("changeset/nope")]))) == ChangesetState("ready", [], None, None) { pass := pass+1; }
    // 155 operationStatusChanged sets running
    var opR := op.(status := "running");
    if apply1(ChangesetState("ready", [], Some([op]), None), OperationStatusChanged("create-pr", "running", None)) == ChangesetState("ready", [], Some([opR]), None) { pass := pass+1; }
    // 157 operationStatusChanged unknown op no-op
    if apply1(ChangesetState("ready", [], Some([op]), None), OperationStatusChanged("nope", "running", None)) == ChangesetState("ready", [], Some([op]), None) { pass := pass+1; }
    // 158 contentChanged replaces files + updates operations
    if apply1(ChangesetState("ready", [fa], None, None), ContentChanged(Some([fb]), Some([op]), None)) == ChangesetState("ready", [fb], Some([op]), None) { pass := pass+1; }
    // 160 filesReviewedChanged sets on matching, skips not-listed
    var fc := ChangesetFile("file:///src/c.ts", Some(false), E());
    var fbT := ChangesetFile("file:///src/b.ts", Some(true), E());
    var faT := ChangesetFile("file:///src/a.ts", Some(true), E());
    var r160 := apply1(ChangesetState("ready", [fa, fbT, fc], None, None), FilesReviewedChanged(["file:///src/a.ts","file:///src/b.ts","file:///src/missing.ts"], true));
    if r160 == ChangesetState("ready", [faT, fbT, fc], None, None) { pass := pass+1; }
    // 161 filesReviewedChanged no-op (nothing needs changing)
    if apply1(ChangesetState("ready", [fbT], None, None), FilesReviewedChanged(["file:///src/b.ts"], true)) == ChangesetState("ready", [fbT], None, None) { pass := pass+1; }
  }
}
