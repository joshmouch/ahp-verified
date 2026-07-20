// AHP Dafny client — REDUCER REPLAY harness (Phase 2), changeset channel.
//
// Mirrors root_replay.dfy exactly, for the `changeset` channel:
//
//   read bytes → Std.JSON parse → bridge to skeleton Wire.Json → decode initial/
//   actions/expected → foldCh(reducer) → assert reduced == decoded-expected
//
// Replays all 15 real changeset fixtures through the ACTUAL ApplyToChangeset
// reducer (not hand-transcribed). Build/run: needs --standard-libraries; pass
// the changeset fixture paths as args.
include "../changeset.dfy"

module ChangesetReplay {
  import opened Changeset
  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import CC = ConfluxContract   // the pinned reducer kernel: Fold (channel-generic left fold)
  import CO = ConfluxOperators
  import FileIO = Std.FileIO
  import API = Std.JSON.API
  import V = Std.JSON.Values
  import opened Std.BoundedInts

  function toJsonBytes(bs: seq<bv8>): seq<uint8> {
    seq(|bs|, i requires 0 <= i < |bs| => (bs[i] as int) as uint8)
  }

  // ── bridge Std.JSON.Values.JSON → the skeleton's Wire.Json (integer numbers; the
  //    faithful: integers stay Wire.JNum (e10==0); fractionals map to Wire.JDec(mantissa,exp)). VERBATIM. ──
  function bridge(j: V.JSON): Wire.Json
    decreases j
  {
    match j
    case Null       => Wire.JNull
    case Bool(b)    => Wire.JBool(b)
    case String(s)  => Wire.JStr(s)
    case Number(d)  => if d.e10 == 0 then Wire.JNum(d.n) else Wire.JDec(d.n, d.e10)   // faithful: integers stay Wire.JNum; fractionals (e.g. 1.5 == Decimal(15,-1)) -> Wire.JDec
    case Array(a)   => Wire.JArr(seq(|a|, i requires 0 <= i < |a| => bridge(a[i])))
    case Object(ps) => Wire.JObj(zipMap(
        seq(|ps|, i requires 0 <= i < |ps| => ps[i].0),
        seq(|ps|, i requires 0 <= i < |ps| => bridge(ps[i].1))))
  }
  function zipMap(keys: seq<string>, vals: seq<Wire.Json>): map<string, Wire.Json>
    requires |keys| == |vals|
  {
    var pairs := seq(|keys|, i requires 0 <= i < |keys| => (keys[|keys|-1-i], vals[|vals|-1-i]));
    CC.Fold((m: map<string, Wire.Json>, p: (string, Wire.Json)) => m[p.0 := p.1], map[], pairs)
  }

  // ── field accessors on skeleton Wire.Json (VERBATIM from root_replay) ──
  function field(j: Wire.Json, k: string): AhpSkeleton.Option<Wire.Json> {
    if j.JObj? && k in j.fields then Some(j.fields[k]) else None
  }
  function asStr(o: AhpSkeleton.Option<Wire.Json>): string { if o.Some? && o.value.JStr? then o.value.s else "" }
  // asInt lives in AhpSkeleton now (JNum + integer-valued JDec); in scope via `import opened AhpSkeleton`.
  function asBool(o: AhpSkeleton.Option<Wire.Json>, dflt: bool): bool { if o.Some? && o.value.JBool? then o.value.b else dflt }
  function asArr(o: AhpSkeleton.Option<Wire.Json>): seq<Wire.Json> { if o.Some? && o.value.JArr? then o.value.elems else [] }
  function asObjMap(o: AhpSkeleton.Option<Wire.Json>): map<string, Wire.Json> { if o.Some? && o.value.JObj? then o.value.fields else map[] }

  function strList(js: seq<Wire.Json>): seq<string>
  { CO.Map((j: Wire.Json) => if j.JStr? then j.s else "", js) }

  // ── decode changeset domain values from skeleton Wire.Json ──
  //   reviewed: present-and-bool → Some; absent → None (open-struct).
  //   edit: raw Wire.Json carried verbatim (reducer never inspects it; equality does).
  function decodeFile(j: Wire.Json): ChangesetFile {
    ChangesetFile(
      asStr(field(j, "id")),
      if field(j, "reviewed").Some? && field(j, "reviewed").value.JBool?
        then Some(field(j, "reviewed").value.b) else None,
      if field(j, "edit").Some? then field(j, "edit").value else Wire.JNull)
  }
  function decodeFiles(js: seq<Wire.Json>): seq<ChangesetFile>
  { CO.Map(decodeFile, js) }

  function decodeOp(j: Wire.Json): ChangesetOperation {
    ChangesetOperation(
      asStr(field(j, "id")),
      asStr(field(j, "label")),
      strList(asArr(field(j, "scopes"))),
      asStr(field(j, "status")),
      if field(j, "error").Some? then Some(field(j, "error").value) else None)
  }
  function decodeOps(js: seq<Wire.Json>): seq<ChangesetOperation>
  { CO.Map(decodeOp, js) }

  // present-and-array → Some(list); absent/null/non-array → None.
  function decodeFilesOpt(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<seq<ChangesetFile>> {
    if o.Some? && o.value.JArr? then Some(decodeFiles(o.value.elems)) else None
  }
  function decodeOpsOpt(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<seq<ChangesetOperation>> {
    if o.Some? && o.value.JArr? then Some(decodeOps(o.value.elems)) else None
  }

  function decodeChangesetState(j: Wire.Json): ChangesetState {
    ChangesetState(
      asStr(field(j, "status")),
      decodeFiles(asArr(field(j, "files"))),
      decodeOpsOpt(field(j, "operations")),
      if field(j, "error").Some? then Some(field(j, "error").value) else None)
  }

  function decodeAction(j: Wire.Json): ChangesetAction {
    var t := asStr(field(j, "type"));
    if t == "changeset/statusChanged" then
      StatusChanged(asStr(field(j, "status")),
                    if field(j, "error").Some? then Some(field(j, "error").value) else None)
    else if t == "changeset/fileSet" then
      FileSet(decodeFile(if field(j, "file").Some? then field(j, "file").value else Wire.JNull))
    else if t == "changeset/fileRemoved" then
      FileRemoved(asStr(field(j, "fileId")))
    else if t == "changeset/operationsChanged" then
      OperationsChanged(decodeOpsOpt(field(j, "operations")))
    else if t == "changeset/cleared" then
      Cleared
    else if t == "changeset/operationStatusChanged" then
      OperationStatusChanged(asStr(field(j, "operationId")), asStr(field(j, "status")),
                             if field(j, "error").Some? then Some(field(j, "error").value) else None)
    else if t == "changeset/contentChanged" then
      ContentChanged(decodeFilesOpt(field(j, "files")), decodeOpsOpt(field(j, "operations")),
                     if field(j, "error").Some? then Some(field(j, "error").value) else None)
    else if t == "changeset/filesReviewChanged" || t == "changeset/filesReviewedChanged" then
      FilesReviewedChanged(
        strList(asArr(if field(j, "files").Some? then field(j, "files") else field(j, "fileIds"))),
        asBool(field(j, "reviewed"), false))
    else
      CsUnknown(j)
  }
  function decodeActions(js: seq<Wire.Json>): seq<ChangesetAction>
  { CO.Map(decodeAction, js) }

  // ── fold the changeset reducer over a decoded action list ──
  function foldCh(s: ChangesetState, acts: seq<ChangesetAction>, now: int): ChangesetState
  { CC.Fold((st, a) => ApplyToChangeset(st, a, now).0, s, acts) }

  // ── replay one fixture file end-to-end through the reducer ──
  method replayOne(path: string) returns (isChangeset: bool, ok: bool) {
    var r := FileIO.ReadBytesFromFile(path);
    if r.Failure? { return false, false; }
    var p := API.Deserialize(toJsonBytes(r.value));
    if p.Failure? { return false, false; }
    var doc := bridge(p.value);
    if !doc.JObj? { return false, false; }
    if asStr(field(doc, "reducer")) != "changeset" { return false, false; }
    var initial := decodeChangesetState(if field(doc, "initial").Some? then field(doc, "initial").value else Wire.JNull);
    var actions := decodeActions(asArr(field(doc, "actions")));
    var expected := decodeChangesetState(if field(doc, "expected").Some? then field(doc, "expected").value else Wire.JNull);
    var reduced := foldCh(initial, actions, 9999);
    isChangeset := true;
    ok := reduced == expected;
  }

  method Main(args: seq<string>) {
    var pass := 0;
    var total := 0;
    var i := 1;
    while i < |args|
      decreases |args| - i
    {
      var isChangeset, ok := replayOne(args[i]);
      if isChangeset {
        total := total + 1;
        if ok { pass := pass + 1; } else { print "REPLAY FAIL: ", args[i], "\n"; }
      }
      i := i + 1;
    }
    print "CHANGESET REDUCER-REPLAY: ", pass, "/", total, " real fixtures replayed through the reducer, reduced == expected\n";
    expect pass == total;
    print "CHANGESET-REPLAY OK — real JSON fixtures drive the proven reducer to the expected state (extracted code)\n";
  }
}
