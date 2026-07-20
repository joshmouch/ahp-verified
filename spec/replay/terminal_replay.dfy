// AHP Dafny client — REDUCER REPLAY harness (Phase 2), terminal channel.
//
// Mirrors root_replay.dfy: instead of hand-transcribing each fixture, it
// REPLAYS the real terminal fixture JSON through the actual reducer:
//
//   read bytes → Std.JSON parse → bridge to skeleton Wire.Json → decode initial/
//   actions/expected → foldCh(ApplyToTerminal) → assert reduced == decoded-expected
//
// The `*terminal*.json` glob also matches 071-root-terminalschanged.json (a
// ROOT-reducer fixture whose filename merely contains "terminal"); replayOne
// reads the fixture's "reducer" field and only COUNTS/replays reducer=="terminal"
// fixtures, so the total is exactly the 19 terminal-reducer fixtures.
// Build/run: needs --standard-libraries; pass the terminal fixture paths as args.
include "../terminal.dfy"

module TerminalReplay {
  import opened Terminal
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
  //    faithful: integers stay Wire.JNum (e10==0); fractionals map to Wire.JDec(mantissa,exp)). VERBATIM from root. ──
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
  // builds a map from already-bridged keys+values (no recursion into JSON). VERBATIM.
  function zipMap(keys: seq<string>, vals: seq<Wire.Json>): map<string, Wire.Json>
    requires |keys| == |vals|
  {
    var pairs := seq(|keys|, i requires 0 <= i < |keys| => (keys[|keys|-1-i], vals[|vals|-1-i]));
    CC.Fold((m: map<string, Wire.Json>, p: (string, Wire.Json)) => m[p.0 := p.1], map[], pairs)
  }

  // ── field accessors on skeleton Wire.Json — VERBATIM from root_replay.dfy ──
  function field(j: Wire.Json, k: string): AhpSkeleton.Option<Wire.Json> {
    if j.JObj? && k in j.fields then Some(j.fields[k]) else None
  }
  function asStr(o: AhpSkeleton.Option<Wire.Json>): string { if o.Some? && o.value.JStr? then o.value.s else "" }
  // asInt lives in AhpSkeleton now (JNum + integer-valued JDec); in scope via `import opened AhpSkeleton`.
  function asBool(o: AhpSkeleton.Option<Wire.Json>, dflt: bool): bool { if o.Some? && o.value.JBool? then o.value.b else dflt }
  function asArr(o: AhpSkeleton.Option<Wire.Json>): seq<Wire.Json> { if o.Some? && o.value.JArr? then o.value.elems else [] }

  // ── optional-field decoders: present+typed → Some(value); else None ──
  function optStr(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<string> {
    if o.Some? && o.value.JStr? then Some(o.value.s) else None
  }
  // optInt lives in AhpSkeleton now (JNum + integer-valued JDec); in scope via `import opened AhpSkeleton`.
  function optBool(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<bool> {
    if o.Some? && o.value.JBool? then Some(o.value.b) else None
  }

  // ── decode a content Part (unclassified | command) ──
  function decodePart(j: Wire.Json): Part {
    var t := asStr(field(j, "type"));
    if t == "command" then
      Command(
        asStr(field(j, "commandId")),
        asStr(field(j, "commandLine")),
        asStr(field(j, "output")),
        asInt(field(j, "timestamp")),
        asBool(field(j, "isComplete"), false),
        optInt(field(j, "exitCode")),      // Option<int> — absent on incomplete commands
        optInt(field(j, "durationMs")))    // Option<int>
    else
      Unclassified(asStr(field(j, "value")))   // "unclassified" (and any other) → value string
  }
  function decodeParts(js: seq<Wire.Json>): seq<Part>
  { CO.Map(decodePart, js) }

  // ── decode a TerminalState — every MODELED field; unmodeled keys are ignored (open struct) ──
  function decodeTerminalState(j: Wire.Json): TerminalState {
    TerminalState(
      asStr(field(j, "title")),
      optStr(field(j, "cwd")),
      optInt(field(j, "cols")),
      optInt(field(j, "rows")),
      decodeParts(asArr(field(j, "content"))),
      field(j, "claim"),                         // Option<Wire.Json>: Some(rawClaim) if present, else None
      optInt(field(j, "exitCode")),
      optBool(field(j, "supportsCommandDetection")))
  }

  // ── decode a TerminalAction — dispatch on the "type" string ──
  function decodeTerminalAction(j: Wire.Json): TerminalAction {
    var t := asStr(field(j, "type"));
    if      t == "terminal/cwdChanged"   then TCwdChanged(asStr(field(j, "cwd")))
    else if t == "terminal/titleChanged" then TTitleChanged(asStr(field(j, "title")))
    else if t == "terminal/resized"      then TResized(asInt(field(j, "cols")), asInt(field(j, "rows")))
    else if t == "terminal/exited"       then TExited(asInt(field(j, "exitCode")))
    else if t == "terminal/data"         then TData(asStr(field(j, "data")))
    else if t == "terminal/cleared"      then TCleared
    else if t == "terminal/claimed"      then TClaimed(if field(j, "claim").Some? then field(j, "claim").value else Wire.JNull)
    else if t == "terminal/commandDetectionAvailable" then TCommandDetectionAvailable
    else if t == "terminal/commandExecuted" then
      TCommandExecuted(asStr(field(j, "commandId")), asStr(field(j, "commandLine")), asInt(field(j, "timestamp")))
    else if t == "terminal/commandFinished" then
      TCommandFinished(asStr(field(j, "commandId")), asInt(field(j, "exitCode")), asInt(field(j, "durationMs")))
    else if t == "terminal/input"        then TInput
    else TUnknown(j)
  }
  function decodeActions(js: seq<Wire.Json>): seq<TerminalAction>
  { CO.Map(decodeTerminalAction, js) }

  // ── fold the actions through the ACTUAL reducer (now is injected; ApplyToTerminal ignores it) ──
  function foldCh(s: TerminalState, acts: seq<TerminalAction>, now: int): TerminalState
  { CC.Fold((st, a) => ApplyToTerminal(st, a, now).0, s, acts) }

  // ── replay one fixture file end-to-end through the reducer ──
  //   returns (isTerminal, ok): isTerminal is false for non-terminal-reducer fixtures
  //   (e.g. 071-root-terminalschanged) so Main does not count them toward the total.
  method replayOne(path: string) returns (isTerminal: bool, ok: bool) {
    var r := FileIO.ReadBytesFromFile(path);
    if r.Failure? { return false, false; }
    var p := API.Deserialize(toJsonBytes(r.value));
    if p.Failure? { return false, false; }
    var doc := bridge(p.value);
    if !doc.JObj? { return false, false; }
    if asStr(field(doc, "reducer")) != "terminal" { return false, false; }  // skip non-terminal fixtures
    var initial := decodeTerminalState(if field(doc, "initial").Some? then field(doc, "initial").value else Wire.JNull);
    var actions := decodeActions(asArr(field(doc, "actions")));
    var expected := decodeTerminalState(if field(doc, "expected").Some? then field(doc, "expected").value else Wire.JNull);
    var reduced := foldCh(initial, actions, 9999);
    isTerminal := true;
    ok := reduced == expected;
  }

  method Main(args: seq<string>) {
    var pass := 0;
    var total := 0;
    var i := 1;
    while i < |args|
      decreases |args| - i
    {
      var isTerminal, ok := replayOne(args[i]);
      if isTerminal {
        total := total + 1;
        if ok { pass := pass + 1; } else { print "REPLAY FAIL: ", args[i], "\n"; }
      }
      i := i + 1;
    }
    print "TERMINAL REDUCER-REPLAY: ", pass, "/", total, " real terminal fixtures replayed through the reducer, reduced == expected\n";
    expect pass == total;
    print "TERMINAL-REPLAY OK — real JSON fixtures drive the proven reducer to the expected state (extracted code)\n";
  }
}
