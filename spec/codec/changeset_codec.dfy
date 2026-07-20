// AHP Dafny client — FULL Wire.Json ⇄ ChangesetAction codec (Changeset channel).
//
// Replicates the PROVEN root-codec template in spec/fidelity.dfy (module Fidelity:
// DecodeRootAction / EncodeRootAction / encodeRootState + RootActionRoundTrip +
// _NonVacuous + RootUnknown_EncodeVerbatim), one channel over.
//
//   DECODE is promoted VERBATIM from spec/replay/changeset_replay.dfy's proven
//   `decodeAction` (the exact logic that drives all 15 changeset fixtures through
//   the reducer); ENCODE is its inverse (canonical wire object per typed variant).
//   Together they carry the load-bearing transport property:
//     decode ∘ encode == id  on the typed variants (ChangesetActionRoundTrip below).
//   The CsUnknown leaf is carved out HONESTLY (verbatim lemmas) — same as
//   Fidelity's RootUnknown: its raw payload could itself be a typed-shaped object,
//   so decode would re-classify it; that leaf re-encodes verbatim, it does not
//   round-trip through the typed decoder.
//
// The shared Wire.Json accessors (field/asStr/asArr/asObjMap/asInt/asBool/strList/
// encodeStrs + StrsRoundTrip) are REUSED from Fidelity (import F = Fidelity), not
// duplicated.

include "../changeset.dfy"
include "../fidelity.dfy"

module ChangesetCodec {
  export
    provides Wire, Changeset, F
    provides DecodeChangesetAction, encodeChangesetState, decodeChangesetState
    provides ChangesetActionRoundTrip, ChangesetStateRoundTrip
    provides encodeFile, encodeFiles, encodeOps
    reveals EncodeChangesetAction

  import opened Changeset
  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import F = Fidelity
  import CO = ConfluxOperators

  // ── DECODE domain values (promoted verbatim from changeset_replay.dfy) ──
  //   reviewed: present-and-bool → Some; absent → None (open-struct).
  //   edit: raw Wire.Json carried verbatim (reducer never inspects it; equality does).
  function decodeFile(j: Wire.Json): ChangesetFile {
    ChangesetFile(
      F.asStr(F.field(j, "id")),
      if F.field(j, "reviewed").Some? && F.field(j, "reviewed").value.JBool?
        then Some(F.field(j, "reviewed").value.b) else None,
      if F.field(j, "edit").Some? then F.field(j, "edit").value else Wire.JNull)
  }
  function decodeFiles(js: seq<Wire.Json>): seq<ChangesetFile>
  { CO.Map(decodeFile, js) }

  function decodeOp(j: Wire.Json): ChangesetOperation {
    ChangesetOperation(
      F.asStr(F.field(j, "id")),
      F.asStr(F.field(j, "label")),
      F.strList(F.asArr(F.field(j, "scopes"))),
      F.asStr(F.field(j, "status")),
      if F.field(j, "error").Some? then Some(F.field(j, "error").value) else None)
  }
  function decodeOps(js: seq<Wire.Json>): seq<ChangesetOperation>
  { CO.Map(decodeOp, js) }

  // present-and-array → Some(list); absent/null/non-array → None.
  function decodeFilesOpt(o: Option<Wire.Json>): Option<seq<ChangesetFile>> {
    if o.Some? && o.value.JArr? then Some(decodeFiles(o.value.elems)) else None
  }
  function decodeOpsOpt(o: Option<Wire.Json>): Option<seq<ChangesetOperation>> {
    if o.Some? && o.value.JArr? then Some(decodeOps(o.value.elems)) else None
  }

  // ── ENCODE domain values (the inverse of the decode above) ──
  function encodeFile(f: ChangesetFile): Wire.Json {
    var m0 := map["id" := Wire.JStr(f.id), "edit" := f.edit];
    var m1 := if f.reviewed.Some? then m0["reviewed" := Wire.JBool(f.reviewed.value)] else m0;
    Wire.JObj(m1)
  }
  function encodeFiles(fs: seq<ChangesetFile>): seq<Wire.Json>
  { CO.Map(encodeFile, fs) }

  function encodeOp(op: ChangesetOperation): Wire.Json {
    var m0 := map["id" := Wire.JStr(op.id), "label" := Wire.JStr(op.label_),
                  "scopes" := Wire.JArr(F.encodeStrs(op.scopes)), "status" := Wire.JStr(op.status)];
    var m1 := if op.error.Some? then m0["error" := op.error.value] else m0;
    Wire.JObj(m1)
  }
  function encodeOps(ops: seq<ChangesetOperation>): seq<Wire.Json>
  { CO.Map(encodeOp, ops) }

  // ── the typed codec ──
  function DecodeChangesetAction(j: Wire.Json): ChangesetAction {
    var t := F.asStr(F.field(j, "type"));
    if t == "changeset/statusChanged" then
      StatusChanged(F.asStr(F.field(j, "status")),
                    if F.field(j, "error").Some? then Some(F.field(j, "error").value) else None)
    else if t == "changeset/fileSet" then
      FileSet(decodeFile(if F.field(j, "file").Some? then F.field(j, "file").value else Wire.JNull))
    else if t == "changeset/fileRemoved" then
      FileRemoved(F.asStr(F.field(j, "fileId")))
    else if t == "changeset/operationsChanged" then
      OperationsChanged(decodeOpsOpt(F.field(j, "operations")))
    else if t == "changeset/cleared" then
      Cleared
    else if t == "changeset/operationStatusChanged" then
      OperationStatusChanged(F.asStr(F.field(j, "operationId")), F.asStr(F.field(j, "status")),
                             if F.field(j, "error").Some? then Some(F.field(j, "error").value) else None)
    else if t == "changeset/contentChanged" then
      ContentChanged(decodeFilesOpt(F.field(j, "files")), decodeOpsOpt(F.field(j, "operations")),
                     if F.field(j, "error").Some? then Some(F.field(j, "error").value) else None)
    else if t == "changeset/filesReviewChanged" || t == "changeset/filesReviewedChanged" then
      FilesReviewedChanged(
        F.strList(F.asArr(if F.field(j, "files").Some? then F.field(j, "files") else F.field(j, "fileIds"))),
        F.asBool(F.field(j, "reviewed"), false))
    else
      CsUnknown(j)
  }

  function EncodeChangesetAction(a: ChangesetAction): Wire.Json {
    match a
    case StatusChanged(st, err) =>
      var m0 := map["type" := Wire.JStr("changeset/statusChanged"), "status" := Wire.JStr(st)];
      Wire.JObj(if err.Some? then m0["error" := err.value] else m0)
    case FileSet(f) =>
      Wire.JObj(map["type" := Wire.JStr("changeset/fileSet"), "file" := encodeFile(f)])
    case FileRemoved(id) =>
      Wire.JObj(map["type" := Wire.JStr("changeset/fileRemoved"), "fileId" := Wire.JStr(id)])
    case OperationsChanged(ops) =>
      var m0 := map["type" := Wire.JStr("changeset/operationsChanged")];
      Wire.JObj(if ops.Some? then m0["operations" := Wire.JArr(encodeOps(ops.value))] else m0)
    case Cleared =>
      Wire.JObj(map["type" := Wire.JStr("changeset/cleared")])
    case OperationStatusChanged(id, st, err) =>
      var m0 := map["type" := Wire.JStr("changeset/operationStatusChanged"),
                    "operationId" := Wire.JStr(id), "status" := Wire.JStr(st)];
      Wire.JObj(if err.Some? then m0["error" := err.value] else m0)
    case ContentChanged(fs, ops, err) =>
      var m0 := map["type" := Wire.JStr("changeset/contentChanged")];
      var m1 := if fs.Some? then m0["files" := Wire.JArr(encodeFiles(fs.value))] else m0;
      var m2 := if ops.Some? then m1["operations" := Wire.JArr(encodeOps(ops.value))] else m1;
      Wire.JObj(if err.Some? then m2["error" := err.value] else m2)
    case FilesReviewedChanged(ids, rev) =>
      Wire.JObj(map["type" := Wire.JStr("changeset/filesReviewChanged"),
               "files" := Wire.JArr(F.encodeStrs(ids)), "reviewed" := Wire.JBool(rev)])
    case CsUnknown(raw) => raw
  }

  // ── ROUND-TRIP proofs: decode ∘ encode == id on the sub-codecs ──
  lemma FileRoundTrip(f: ChangesetFile)
    ensures decodeFile(encodeFile(f)) == f
  {
    if f.reviewed.Some? {} else {}
  }
  lemma FilesRoundTrip(fs: seq<ChangesetFile>)
    ensures decodeFiles(encodeFiles(fs)) == fs
    decreases fs
  {
    if |fs| == 0 {} else { FileRoundTrip(fs[0]); FilesRoundTrip(fs[1..]); }
  }

  lemma OpRoundTrip(op: ChangesetOperation)
    ensures decodeOp(encodeOp(op)) == op
  {
    F.StrsRoundTrip(op.scopes);
    if op.error.Some? {} else {}
  }
  lemma OpsRoundTrip(ops: seq<ChangesetOperation>)
    ensures decodeOps(encodeOps(ops)) == ops
    decreases ops
  {
    if |ops| == 0 {} else { OpRoundTrip(ops[0]); OpsRoundTrip(ops[1..]); }
  }

  // The load-bearing transport property: every TYPED action encodes to a wire
  // object that decodes back to itself. (CsUnknown is excluded — its raw payload
  // may itself be a typed-shaped object, so decode would re-classify it; that leaf
  // is handled by the verbatim lemmas below, not this one — exactly like Fidelity's
  // RootUnknown carve-out.)
  // Json→domain→Json fixpoint: the codec's canonical wire form re-decodes + re-encodes to ITSELF — the
  // direction the conformance suite exercises, now machine-checked. Follows from the direction-1 round-trip.
  lemma ChangesetWireCanonicalRoundTrip(a: ChangesetAction)
    requires !a.CsUnknown?
    ensures EncodeChangesetAction(DecodeChangesetAction(EncodeChangesetAction(a))) == EncodeChangesetAction(a)
  { ChangesetActionRoundTrip(a); }
  lemma ChangesetActionRoundTrip(a: ChangesetAction)
    requires !a.CsUnknown?
    ensures DecodeChangesetAction(EncodeChangesetAction(a)) == a
  {
    match a
    case StatusChanged(st, err)              =>
    case FileSet(f)                          => FileRoundTrip(f);
    case FileRemoved(id)                     =>
    case OperationsChanged(ops)              => if ops.Some? { OpsRoundTrip(ops.value); }
    case Cleared                             =>
    case OperationStatusChanged(id, st, err) =>
    case ContentChanged(fs, ops, err)        =>
      if fs.Some? { FilesRoundTrip(fs.value); }
      if ops.Some? { OpsRoundTrip(ops.value); }
    case FilesReviewedChanged(ids, rev)      => F.StrsRoundTrip(ids);
  }

  // Non-vacuity witness: the round-trip precondition is inhabited by a GENUINE
  // typed action carrying real payload (a seq of file ids exercised through the
  // strList/encodeStrs induction), so ChangesetActionRoundTrip is NOT vacuously true.
  lemma ChangesetActionRoundTrip_NonVacuous()
    ensures exists a: ChangesetAction :: !a.CsUnknown? && DecodeChangesetAction(EncodeChangesetAction(a)) == a
  {
    var a := FilesReviewedChanged(["file:///src/a.ts", "file:///src/b.ts"], true);
    ChangesetActionRoundTrip(a);
    assert !a.CsUnknown? && DecodeChangesetAction(EncodeChangesetAction(a)) == a;
  }

  // ── unknown-variant passthrough (fixture 144) — the CsUnknown leaf re-encodes verbatim ──
  lemma CsUnknown_EncodeVerbatim(j: Wire.Json)
    ensures EncodeChangesetAction(CsUnknown(j)) == j {}
  lemma CsUnknownVariant_Verbatim(j: Wire.Json)
    requires DecodeChangesetAction(j).CsUnknown?
    ensures EncodeChangesetAction(DecodeChangesetAction(j)) == j {}

  // ── STATE serialization (for the runnable cross-process transport demo) ──
  // encodeChangesetState serializes a ChangesetState to canonical JSON so the host
  // and the client can each emit their final state and a cross-process e2e can
  // compare them. It is a DETERMINISTIC TOTAL function (mirror of encodeRootState);
  // the transport demo's convergence check relies on that determinism, not on a
  // ChangesetState round-trip lemma (a natural future hardening, not needed here).
  function encodeChangesetState(s: ChangesetState): Wire.Json {
    var m0 := map["status" := Wire.JStr(s.status), "files" := Wire.JArr(encodeFiles(s.files))];
    var m1 := if s.operations.Some? then m0["operations" := Wire.JArr(encodeOps(s.operations.value))] else m0;
    var m2 := if s.error.Some? then m1["error" := s.error.value] else m1;
    Wire.JObj(m2)
  }

  // ── STATE round-trip: decodeChangesetState ∘ encodeChangesetState == id ──
  // The TRUE INVERSE of encodeChangesetState above. Promoted from the proven
  // decodeChangesetState in spec/replay/changeset_replay.dfy (the exact logic that
  // decodes all 15 changeset fixtures' initial/expected states), re-expressed over the
  // shared Fidelity accessors (F.field/asStr/asArr) + this module's decodeFiles/
  // decodeOpsOpt. It reads EXACTLY the keys encodeChangesetState emits:
  //   status → Wire.JStr(s.status), verbatim;
  //   files  → Wire.JArr(encodeFiles(s.files))          — inverse via FilesRoundTrip;
  //   operations → present Wire.JArr ⇔ Some / absent ⇔ None — inverse via OpsRoundTrip;
  //   error  → opaque Wire.Json carried verbatim, present-key ⇔ Some / absent ⇔ None.
  function decodeChangesetState(j: Wire.Json): ChangesetState {
    ChangesetState(
      F.asStr(F.field(j, "status")),
      decodeFiles(F.asArr(F.field(j, "files"))),
      decodeOpsOpt(F.field(j, "operations")),
      if F.field(j, "error").Some? then Some(F.field(j, "error").value) else None)
  }

  // REAL, quantified round-trip. NO Wf precondition is needed: EVERY field
  // round-trips cleanly, so the STATE codec's proof story is complete in both
  // directions (encode is total+deterministic; decode is its exact inverse).
  //
  //   error: Option<Wire.Json> does NOT suffer the Some(Wire.JNull)→absent collapse that
  //   forces a carve-out elsewhere, because encodeChangesetState emits the "error"
  //   key whenever error.Some? (INCLUDING Some(Wire.JNull)), and F.field keys off key
  //   PRESENCE — so Some(Wire.JNull) decodes back to Some(Wire.JNull) and an absent key
  //   decodes to None. operations: Option<seq<..>> is likewise present-key ⇔ Some.
  //   Nothing is dropped; unlike the CsUnknown action leaf, NO honest carve-out is
  //   required for the STATE — hence no requires clause.
  lemma ChangesetStateRoundTrip(s: ChangesetState)
    ensures decodeChangesetState(encodeChangesetState(s)) == s
  {
    FilesRoundTrip(s.files);
    if s.operations.Some? { OpsRoundTrip(s.operations.value); }
  }

  // Non-vacuity witness: a GENUINE non-empty state — non-empty files carrying a
  // reviewed flag + opaque edit, Some operations with a real op seq (exercising the
  // encodeOps/decodeOps induction), and Some error — round-trips. So the lemma is not
  // accidentally provable only for the empty/None state.
  lemma ChangesetStateRoundTrip_NonVacuous()
    ensures exists s: ChangesetState :: decodeChangesetState(encodeChangesetState(s)) == s
  {
    var f  := ChangesetFile("file:///src/a.ts", Some(true), Wire.JObj(map["diff" := Wire.JNull]));
    var op := ChangesetOperation("create-pr", "Create Pull Request", ["changeset"], "running", Some(Wire.JStr("boom")));
    var s  := ChangesetState("ready", [f], Some([op]), Some(Wire.JStr("err")));
    ChangesetStateRoundTrip(s);
    assert decodeChangesetState(encodeChangesetState(s)) == s;
  }

  ghost predicate ChangesetCodecWf(a: ChangesetAction) { !a.CsUnknown? }
  function ChangesetCodecC(): Wire.Codec<ChangesetAction> {
    Wire.Codec(EncodeChangesetAction, (j: Wire.Json) => Wire.Some(DecodeChangesetAction(j)))
  }
  lemma ChangesetCodecRoundTripsWhere()
    ensures forall a :: ChangesetCodecWf(a) ==> ChangesetCodecC().decode(ChangesetCodecC().encode(a)) == Wire.Some(a)
  {
    forall a | ChangesetCodecWf(a)
      ensures ChangesetCodecC().decode(ChangesetCodecC().encode(a)) == Wire.Some(a)
    { ChangesetActionRoundTrip(a); }
  }
}
