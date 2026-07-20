// AHP Dafny client — REDUCER REPLAY harness (Phase 2), annotations channel.
//
// Mirrors root_replay.dfy exactly: instead of hand-transcribing each fixture,
// this REPLAYS the real annotations fixture JSON through the actual reducer:
//
//   read bytes → Std.JSON parse → bridge to skeleton Wire.Json → decode initial/
//   actions/expected → foldCh(ApplyToAnnotations) → assert reduced == decoded-expected
//
// It proves the Phase-2 reducer on REAL data (not hand-transcribed) for the
// annotations channel. Build/run: needs --standard-libraries; pass the
// annotations fixture paths as args.
include "../annotations.dfy"

module AnnotationsReplay {
  import opened Annotations
  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import CC = ConfluxContract   // the pinned reducer kernel: Fold (channel-generic left fold)
  import CO = ConfluxOperators  // the pure map primitive: Map (uniform seq->seq decode)
  import FileIO = Std.FileIO
  import API = Std.JSON.API
  import V = Std.JSON.Values
  import opened Std.BoundedInts

  function toJsonBytes(bs: seq<bv8>): seq<uint8> {
    seq(|bs|, i requires 0 <= i < |bs| => (bs[i] as int) as uint8)
  }

  // ── bridge Std.JSON.Values.JSON → the skeleton's Wire.Json (integer numbers; the
  //    faithful: integers stay Wire.JNum (e10==0); fractionals map to Wire.JDec(mantissa,exp)). VERBATIM from
  //    root_replay.dfy — channel-agnostic. ──
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
  // Pair the parallel inputs, then delegate the right-to-left map accumulation
  // to the canonical Fold. Reversing preserves the prior duplicate-key rule:
  // the earliest key wins because it is installed last.
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
  function asObjMap(o: AhpSkeleton.Option<Wire.Json>): map<string, Wire.Json> { if o.Some? && o.value.JObj? then o.value.fields else map[] }

  function strList(js: seq<Wire.Json>): seq<string>
  { CO.Map((j: Wire.Json) => if j.JStr? then j.s else "", js) }

  // ── optional-scalar accessors (present JSON key → Some, absent → None) ──
  function optStr(j: Wire.Json, k: string): AhpSkeleton.Option<string> {
    if field(j, k).Some? then Some(asStr(field(j, k))) else None
  }
  function optBool(j: Wire.Json, k: string): AhpSkeleton.Option<bool> {
    if field(j, k).Some? then Some(asBool(field(j, k), false)) else None
  }

  // ── decode annotations-channel domain values from skeleton Wire.Json ──
  // `range` and `_meta` are opaque Wire.Json (Option): field() already yields
  // Some(value) when the key is present and None when absent — the exact
  // present→set / absent→preserve semantics the model uses.
  function decodeEntry(j: Wire.Json): Entry {
    Entry(asStr(field(j, "id")), asStr(field(j, "text")), field(j, "_meta"))
  }
  function decodeEntries(js: seq<Wire.Json>): seq<Entry>
  { CO.Map(decodeEntry, js) }

  function decodeAnnotation(j: Wire.Json): Annotation {
    Annotation(
      asStr(field(j, "id")),
      asStr(field(j, "turnId")),
      asStr(field(j, "resource")),
      field(j, "range"),                               // Option<Wire.Json>, opaque
      asBool(field(j, "resolved"), false),
      decodeEntries(asArr(field(j, "entries"))),
      field(j, "_meta"))                               // Option<Wire.Json>, opaque
  }
  function decodeAnnotations(js: seq<Wire.Json>): seq<Annotation>
  { CO.Map(decodeAnnotation, js) }

  function decodeAnnotationsState(j: Wire.Json): AnnotationsState {
    AnnotationsState(decodeAnnotations(asArr(field(j, "annotations"))))
  }

  function decodeAnnotationsAction(j: Wire.Json): AnnotationsAction {
    var t := asStr(field(j, "type"));
    if t == "annotations/set" then
      Set(decodeAnnotation(if field(j, "annotation").Some? then field(j, "annotation").value else Wire.JNull))
    else if t == "annotations/removed" then
      Removed(asStr(field(j, "annotationId")))
    else if t == "annotations/entrySet" then
      EntrySet(asStr(field(j, "annotationId")),
               decodeEntry(if field(j, "entry").Some? then field(j, "entry").value else Wire.JNull))
    else if t == "annotations/entryRemoved" then
      EntryRemoved(asStr(field(j, "annotationId")), asStr(field(j, "entryId")))
    else if t == "annotations/updated" then
      Updated(asStr(field(j, "annotationId")),
              optStr(j, "turnId"), optStr(j, "resource"),
              field(j, "range"), optBool(j, "resolved"))
    else
      AnUnknown(j)
  }
  function decodeActions(js: seq<Wire.Json>): seq<AnnotationsAction>
  { CO.Map(decodeAnnotationsAction, js) }

  // ── fold the decoded actions through the ACTUAL reducer ──
  function foldCh(s: AnnotationsState, acts: seq<AnnotationsAction>, now: int): AnnotationsState
  { CC.Fold((st, a) => ApplyToAnnotations(st, a, now).0, s, acts) }

  // ── replay one fixture file end-to-end through the reducer ──
  method replayOne(path: string) returns (ok: bool) {
    var r := FileIO.ReadBytesFromFile(path);
    if r.Failure? { return false; }
    var p := API.Deserialize(toJsonBytes(r.value));
    if p.Failure? { return false; }
    var doc := bridge(p.value);
    if !doc.JObj? { return false; }
    var initial := decodeAnnotationsState(if field(doc, "initial").Some? then field(doc, "initial").value else Wire.JNull);
    var actions := decodeActions(asArr(field(doc, "actions")));
    var expected := decodeAnnotationsState(if field(doc, "expected").Some? then field(doc, "expected").value else Wire.JNull);
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
    print "ANNOTATIONS REDUCER-REPLAY: ", pass, "/", total, " real fixtures replayed through the reducer, reduced == expected\n";
    expect pass == total;
    print "ANNOTATIONS-REPLAY OK — real JSON fixtures drive the proven reducer to the expected state (extracted code)\n";
  }
}
