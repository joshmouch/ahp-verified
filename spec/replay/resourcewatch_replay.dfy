// AHP Dafny client — REDUCER REPLAY harness (Phase 2), resourceWatch channel.
//
// Mirrors root_replay.dfy structurally: REPLAYS the real fixture JSON through
// the actual reducer instead of hand-transcribing:
//
//   read bytes → Std.JSON parse → bridge to skeleton Wire.Json → decode initial/
//   actions/expected → foldCh(ApplyToResourceWatch) → assert reduced == decoded-expected
//
// Build/run: needs --standard-libraries; pass the resourcewatch fixture paths as args.
include "../resource_watch.dfy"

module ResourcewatchReplay {
  import opened ResourceWatch
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

  // ── bridge Std.JSON.Values.JSON → the skeleton's Wire.Json (VERBATIM from root_replay) ──
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

  // ── decode ResourceWatchState — every modeled field (root, recursive) ──
  function decodeResourceWatchState(j: Wire.Json): ResourceWatchState {
    ResourceWatchState(asStr(field(j, "root")), asBool(field(j, "recursive"), false))
  }

  // ── decode ResourceWatchAction — dispatch on the JSON 'type' string ──
  function decodeResourceWatchAction(j: Wire.Json): ResourceWatchAction {
    var t := asStr(field(j, "type"));
    if t == "resourceWatch/changed"
    then RWChanged(if field(j, "changes").Some? then field(j, "changes").value else Wire.JNull)
    else RWUnknown(j)
  }
  function decodeActions(js: seq<Wire.Json>): seq<ResourceWatchAction>
  { CO.Map(decodeResourceWatchAction, js) }

  // ── fold the actions through the ACTUAL reducer ──
  function foldCh(s: ResourceWatchState, acts: seq<ResourceWatchAction>, now: int): ResourceWatchState
  { CC.Fold((st, a) => ApplyToResourceWatch(st, a, now).0, s, acts) }

  // ── replay one fixture file end-to-end through the reducer ──
  method replayOne(path: string) returns (ok: bool) {
    var r := FileIO.ReadBytesFromFile(path);
    if r.Failure? { return false; }
    var p := API.Deserialize(toJsonBytes(r.value));
    if p.Failure? { return false; }
    var doc := bridge(p.value);
    if !doc.JObj? { return false; }
    var initial := decodeResourceWatchState(if field(doc, "initial").Some? then field(doc, "initial").value else Wire.JNull);
    var actions := decodeActions(asArr(field(doc, "actions")));
    var expected := decodeResourceWatchState(if field(doc, "expected").Some? then field(doc, "expected").value else Wire.JNull);
    var reduced := foldCh(initial, actions, 9999);
    ok := reduced == expected;
  }

  method Main(args: seq<string>) {
    var pass := 0;
    var total := 0;
    var i := 1;
    while i < |args|
      decreases |args| - i
    {
      total := total + 1;
      var ok := replayOne(args[i]);
      if ok { pass := pass + 1; } else { print "REPLAY FAIL: ", args[i], "\n"; }
      i := i + 1;
    }
    print "RESOURCEWATCH REDUCER-REPLAY: ", pass, "/", total, " real fixtures replayed through the reducer, reduced == expected\n";
    expect pass == total;
    print "RESOURCEWATCH-REPLAY OK — real JSON fixtures drive the proven reducer to the expected state (extracted code)\n";
  }
}
