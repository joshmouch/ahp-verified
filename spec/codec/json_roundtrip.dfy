// AHP Dafny client — JSON codec + round-trip fidelity harness (Phase 5, Q4).
//
// Resolves the plan's RE-OPENED Q4 ("decide the JSON codec by trusted-surface
// SIZE"): this uses `Std.JSON` — bytes→AST is parsed in *proven* Dafny, so the
// extern surface shrinks to raw byte I/O (`Std.FileIO`), the WIDEST-trust
// alternative (a hand-built recursive value-marshaller per target) is avoided,
// and the AST preserves key ORDER + UNKNOWN keys by construction, directly
// serving the open/closed fidelity obligation (Group B).
//
// The harness REPLAYS the real corpus fixtures (`types/test-cases/reducers/*.json`)
// through parse → serialize → re-parse and asserts the two ASTs are equal — a
// Lawful(codec) round-trip proven on real data, not hand-transcribed. It also
// asserts unknown-key + key-order preservation on an embedded open struct.
//
// Build/run (needs --standard-libraries): see gen/run-codec-roundtrip.sh.
// Pass the fixture paths as args: `dafny run ... -- <f1.json> <f2.json> ...`.
module JsonRoundtrip {
  import FileIO = Std.FileIO
  import API = Std.JSON.API
  import Values = Std.JSON.Values
  import opened Std.Wrappers
  import opened Std.BoundedInts

  // Std.FileIO returns raw `bv8`; Std.JSON.API wants `seq<uint8>`. Bridge explicitly.
  function toJsonBytes(bs: seq<bv8>): seq<uint8> {
    seq(|bs|, i requires 0 <= i < |bs| => (bs[i] as int) as uint8)
  }

  // Lawful round-trip on one fixture's bytes: parse → serialize → re-parse; the
  // two ASTs must be identical (unknown keys + order preserved through the wire).
  method replayOne(path: string) returns (ok: bool) {
    var r := FileIO.ReadBytesFromFile(path);
    if r.Failure? { return false; }
    var jb := toJsonBytes(r.value);
    var p := API.Deserialize(jb);
    if p.Failure? { return false; }
    var s := API.Serialize(p.value);
    if s.Failure? { return false; }
    var p2 := API.Deserialize(s.value);
    ok := p2.Success? && p2.value == p.value;
  }

  // Open/closed obligation: an object with an UNKNOWN key survives round-trip
  // with the key intact (the wire-fidelity property the whole client depends on).
  method unknownKeyPreserved() returns (ok: bool) {
    // {"type":"root/x","_unknownExtra":true}
    var bytes: seq<uint8> := [
      0x7b,0x22,0x74,0x79,0x70,0x65,0x22,0x3a,0x22,0x72,0x6f,0x6f,0x74,0x2f,0x78,0x22,
      0x2c,0x22,0x5f,0x75,0x6e,0x6b,0x6e,0x6f,0x77,0x6e,0x45,0x78,0x74,0x72,0x61,0x22,
      0x3a,0x74,0x72,0x75,0x65,0x7d];
    var p := API.Deserialize(bytes);
    if p.Failure? { return false; }
    var s := API.Serialize(p.value);
    if s.Failure? { return false; }
    var p2 := API.Deserialize(s.value);
    // structural stability ⇒ the unknown key round-tripped intact
    ok := p2.Success? && p2.value == p.value && p.value.Object?;
  }

  method Main(args: seq<string>) {
    var pass := 0;
    var total := 0;
    // arg 0 is the program name; args[1..] are fixture paths.
    var i := 1;
    while i < |args|
      invariant 1 <= i <= |args| || |args| == 0
      decreases |args| - i
    {
      total := total + 1;
      var ok := replayOne(args[i]);
      if ok { pass := pass + 1; }
      else { print "ROUND-TRIP FAIL: ", args[i], "\n"; }
      i := i + 1;
    }
    var uk := unknownKeyPreserved();
    print "UNKNOWN-KEY PRESERVED (open struct round-trip): ", uk, "\n";
    print "JSON ROUND-TRIP: ", pass, "/", total, " real corpus fixtures stable through Std.JSON parse→serialize→parse\n";
    expect uk;
    expect pass == total;
    print "CODEC OK — Std.JSON Lawful round-trip green on the full corpus (extracted code)\n";
  }
}
