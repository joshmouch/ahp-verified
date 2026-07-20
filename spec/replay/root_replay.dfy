// AHP Dafny client — REDUCER REPLAY harness (Phase 2), root channel.
//
// The corpus checks in ahp_skeleton.dfy hand-transcribe each fixture. THIS
// harness instead REPLAYS the real fixture JSON through the actual reducer:
//
//   read bytes → Std.JSON parse → bridge to skeleton Wire.Json → decode initial/
//   actions/expected → FoldRoot(reducer) → assert reduced == decoded-expected
//
// It proves the Phase-2 reducer harness on REAL data (not hand-transcribed) for
// one channel — establishing the decoder idiom the other 6 channels follow.
// Build/run: needs --standard-libraries; pass the root fixture paths as args.
include "../ahp_skeleton.dfy"

module RootReplay {
  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import CC = ConfluxContract
  import CO = ConfluxOperators
  import FileIO = Std.FileIO
  import API = Std.JSON.API
  import V = Std.JSON.Values
  import opened Std.BoundedInts

  function toJsonBytes(bs: seq<bv8>): seq<uint8> {
    seq(|bs|, i requires 0 <= i < |bs| => (bs[i] as int) as uint8)
  }

  // ── bridge Std.JSON.Values.JSON → the skeleton's Wire.Json (integer numbers; the
  //    faithful: integers stay Wire.JNum (e10==0); fractionals map to Wire.JDec(mantissa,exp)). ──
  function bridge(j: V.JSON): Wire.Json
    decreases j
  {
    match j
    case Null       => Wire.JNull
    case Bool(b)    => Wire.JBool(b)
    case String(s)  => Wire.JStr(s)
    case Number(d)  => if d.e10 == 0 then Wire.JNum(d.n) else Wire.JDec(d.n, d.e10)   // faithful: integers stay Wire.JNum; fractionals (e.g. 1.5 == Decimal(15,-1)) -> Wire.JDec
    // inline seq-comprehensions: a[i] / ps[i].1 are subterms of j, so the
    // recursive bridge call provably decreases — no mutual recursion needed.
    case Array(a)   => Wire.JArr(seq(|a|, i requires 0 <= i < |a| => bridge(a[i])))
    case Object(ps) => Wire.JObj(zipMap(
        seq(|ps|, i requires 0 <= i < |ps| => ps[i].0),
        seq(|ps|, i requires 0 <= i < |ps| => bridge(ps[i].1))))
  }
  // builds a map from already-bridged keys+values (no recursion into JSON).
  function zipMap(keys: seq<string>, vals: seq<Wire.Json>): map<string, Wire.Json>
    requires |keys| == |vals|
  {
    var pairs := seq(|keys|, i requires 0 <= i < |keys| => (keys[|keys|-1-i], vals[|vals|-1-i]));
    CC.Fold((m: map<string, Wire.Json>, p: (string, Wire.Json)) => m[p.0 := p.1], map[], pairs)
  }

  // ── field accessors on skeleton Wire.Json ──
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

  // ── decode domain values from skeleton Wire.Json ──
  function decodeAgent(j: Wire.Json): AgentInfo {
    AgentInfo(asStr(field(j, "provider")), asStr(field(j, "displayName")),
              asStr(field(j, "description")), strList(asArr(field(j, "models"))))
  }
  function decodeAgents(js: seq<Wire.Json>): seq<AgentInfo>
  { CO.Map(decodeAgent, js) }

  function decodeConfig(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<RootConfig> {
    if o.None? || !o.value.JObj? then None
    else Some(RootConfig(
      if field(o.value, "schema").Some? then field(o.value, "schema").value else Wire.JNull,
      asObjMap(field(o.value, "values"))))
  }

  function decodeRootState(j: Wire.Json): RootState {
    RootState(
      decodeAgents(asArr(field(j, "agents"))),
      if field(j, "activeSessions").Some? && field(j, "activeSessions").value.JNum?
        then Some(asInt(field(j, "activeSessions"))) else None,
      if field(j, "terminals").Some? && field(j, "terminals").value.JArr?
        then Some(asArr(field(j, "terminals"))) else None,
      decodeConfig(field(j, "config")))
  }

  function decodeAction(j: Wire.Json): RootAction {
    var t := asStr(field(j, "type"));
    if t == "root/agentsChanged" then RootAgentsChanged(decodeAgents(asArr(field(j, "agents"))))
    else if t == "root/activeSessionsChanged" then RootActiveSessionsChanged(asInt(field(j, "activeSessions")))
    else if t == "root/terminalsChanged" then RootTerminalsChanged(asArr(field(j, "terminals")))
    else if t == "root/configChanged" then RootConfigChanged(asObjMap(field(j, "config")), asBool(field(j, "replace"), false))
    else RootUnknown(j)
  }
  function decodeActions(js: seq<Wire.Json>): seq<RootAction>
  { CO.Map(decodeAction, js) }

  // ── replay one fixture file end-to-end through the reducer ──
  method replayOne(path: string) returns (ok: bool) {
    var r := FileIO.ReadBytesFromFile(path);
    if r.Failure? { return false; }
    var p := API.Deserialize(toJsonBytes(r.value));
    if p.Failure? { return false; }
    var doc := bridge(p.value);
    if !doc.JObj? { return false; }
    var initial := decodeRootState(if field(doc, "initial").Some? then field(doc, "initial").value else Wire.JNull);
    var actions := decodeActions(asArr(field(doc, "actions")));
    var expected := decodeRootState(if field(doc, "expected").Some? then field(doc, "expected").value else Wire.JNull);
    var reduced := FoldRoot(initial, actions, 9999);
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
    print "ROOT REDUCER-REPLAY: ", pass, "/", total, " real fixtures replayed through the reducer, reduced == expected\n";
    expect pass == total;
    print "ROOT-REPLAY OK — real JSON fixtures drive the proven reducer to the expected state (extracted code)\n";
  }
}
