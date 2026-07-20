// Real pinned Canvas fixture replay: bytes -> Std.JSON -> Wire.Json -> typed
// Canvas codec -> Conflux fold -> expected state.
include "../codec/canvas_codec.dfy"

module CanvasReplay {
  import opened Canvas
  import opened AhpSkeleton
  import Codec = CanvasCodec
  import Wire = ConfluxCodec
  import CC = ConfluxContract
  import CO = ConfluxOperators
  import FileIO = Std.FileIO
  import API = Std.JSON.API
  import V = Std.JSON.Values
  import opened Std.BoundedInts

  function toJsonBytes(bs: seq<bv8>): seq<uint8> {
    seq(|bs|, i requires 0 <= i < |bs| => (bs[i] as int) as uint8)
  }

  function bridge(j: V.JSON): Wire.Json
    decreases j
  {
    match j
    case Null       => Wire.JNull
    case Bool(b)    => Wire.JBool(b)
    case String(s)  => Wire.JStr(s)
    case Number(d)  => if d.e10 == 0 then Wire.JNum(d.n) else Wire.JDec(d.n, d.e10)
    case Array(a)   => Wire.JArr(seq(|a|, i requires 0 <= i < |a| => bridge(a[i])))
    case Object(ps) => Wire.JObj(zipMap(
      seq(|ps|, i requires 0 <= i < |ps| => ps[i].0),
      seq(|ps|, i requires 0 <= i < |ps| => bridge(ps[i].1))))
  }

  function zipMap(keys: seq<string>, vals: seq<Wire.Json>): map<string, Wire.Json>
    requires |keys| == |vals|
  {
    var pairs := seq(|keys|, i requires 0 <= i < |keys| =>
      (keys[|keys|-1-i], vals[|vals|-1-i]));
    CC.Fold((m: map<string, Wire.Json>, p: (string, Wire.Json)) =>
      m[p.0 := p.1], map[], pairs)
  }

  function field(j: Wire.Json, k: string): AhpSkeleton.Option<Wire.Json> {
    if j.JObj? && k in j.fields then Some(j.fields[k]) else None
  }
  function asArr(o: AhpSkeleton.Option<Wire.Json>): seq<Wire.Json> {
    if o.Some? && o.value.JArr? then o.value.elems else []
  }
  function asStr(o: AhpSkeleton.Option<Wire.Json>): string {
    if o.Some? && o.value.JStr? then o.value.s else ""
  }

  function decodeActions(js: seq<Wire.Json>): seq<CanvasAction>
  { CO.Map(Codec.DecodeCanvasAction, js) }

  method replayOne(path: string) returns (isCanvas: bool, ok: bool) {
    var r := FileIO.ReadBytesFromFile(path);
    if r.Failure? { return false, false; }
    var p := API.Deserialize(toJsonBytes(r.value));
    if p.Failure? { return false, false; }
    var doc := bridge(p.value);
    if !doc.JObj? { return false, false; }
    if asStr(field(doc, "reducer")) != "canvas" { return false, false; }
    var initial := Codec.decodeCanvasState(
      if field(doc, "initial").Some? then field(doc, "initial").value else Wire.JNull);
    var actions := decodeActions(asArr(field(doc, "actions")));
    var expected := Codec.decodeCanvasState(
      if field(doc, "expected").Some? then field(doc, "expected").value else Wire.JNull);
    isCanvas := true;
    ok := fold(initial, actions) == expected;
  }

  method Main(args: seq<string>) {
    var pass := 0;
    var total := 0;
    var i := 1;
    while i < |args|
      decreases |args| - i
    {
      var isCanvas, ok := replayOne(args[i]);
      if isCanvas {
        total := total + 1;
        if ok { pass := pass + 1; }
        else { print "REPLAY FAIL: ", args[i], "\n"; }
      }
      i := i + 1;
    }
    print "CANVAS REDUCER-REPLAY: ", pass, "/", total,
      " real fixtures replayed through the reducer, reduced == expected\n";
    expect pass == total;
    print "CANVAS-REPLAY OK — pinned Canvas JSON fixtures drive the proven reducer\n";
  }
}
