// AHP Dafny client — REDUCER REPLAY harness (Phase 2), session channel.
//
// Mirrors root_replay.dfy / changeset_replay.dfy: instead of hand-transcribing
// each fixture (Session.RunCorpus does that), it REPLAYS the real session
// fixture JSON through the actual ApplyToSession reducer:
//
//   read bytes → Std.JSON parse → bridge to skeleton Wire.Json → decode initial/
//   actions/expected → foldCh(ApplyToSession) → assert reduced == decoded-expected
//
// The `*session*.json` glob also matches fixtures whose FILENAME contains
// "session" but whose `reducer` field is chat/terminal/root (turnStarted,
// delta, truncated, input flows, terminal/claimed, root/activeSessionsChanged,
// …). replayOne reads the fixture's "reducer" field and only COUNTS/replays
// reducer=="session" fixtures — exactly the 61 session-reducer fixtures.
//
// ── modifiedAt normalization: NOT APPLICABLE to the session channel ──
// The template harnesses normalize a clock-derived `modifiedAt` display field
// (the #186 concern) before comparing. SessionState has NO such field. The only
// `modifiedAt` in this channel lives inside each `Chat` entry, and it is NOT
// clock-derived: chatAdded/chatUpdated thread the modifiedAt value straight
// from the action payload (fixtures carry explicit ISO strings like
// "1970-01-01T00:00:02.000Z"). It is therefore REAL reducer state and is
// compared for real. Normalizing it would MASK genuine differences — so this
// harness compares reduced == expected directly, with NO normalization. This is
// the only honest call for a channel with no injected-clock field.
//
// Build/run: needs --standard-libraries; pass the session fixture paths as args.
include "../session.dfy"

module SessionReplay {
  import opened Session
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

  // ── bridge Std.JSON.Values.JSON → the skeleton's Wire.Json. VERBATIM from root. ──
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

  // ── field accessors on skeleton Wire.Json — VERBATIM from root_replay.dfy ──
  function field(j: Wire.Json, k: string): AhpSkeleton.Option<Wire.Json> {
    if j.JObj? && k in j.fields then Some(j.fields[k]) else None
  }
  function asStr(o: AhpSkeleton.Option<Wire.Json>): string { if o.Some? && o.value.JStr? then o.value.s else "" }
  // asInt lives in AhpSkeleton now (JNum + integer-valued JDec); in scope via `import opened AhpSkeleton`.
  function asBool(o: AhpSkeleton.Option<Wire.Json>, dflt: bool): bool { if o.Some? && o.value.JBool? then o.value.b else dflt }
  function asArr(o: AhpSkeleton.Option<Wire.Json>): seq<Wire.Json> { if o.Some? && o.value.JArr? then o.value.elems else [] }
  function asObjMap(o: AhpSkeleton.Option<Wire.Json>): map<string, Wire.Json> { if o.Some? && o.value.JObj? then o.value.fields else map[] }

  // ── optional-field decoders: present+typed → Some; else None. VERBATIM. ──
  function optStr(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<string> {
    if o.Some? && o.value.JStr? then Some(o.value.s) else None
  }
  // optInt lives in AhpSkeleton now (JNum + integer-valued JDec); in scope via `import opened AhpSkeleton`.
  function optBool(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<bool> {
    if o.Some? && o.value.JBool? then Some(o.value.b) else None
  }
  // present-and-NON-null → Some(rawJson); absent OR JSON `null` → None. Used for
  // state/channel/meta/etc. so that an explicit `"channel": null` in expected
  // (what mcpServerStart/Stop's cleared channel produces) decodes to None,
  // matching the reducer's `None`.
  function optJsonNoNull(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<Wire.Json> {
    if o.Some? && !o.value.JNull? then Some(o.value) else None
  }
  // status is a bv32 bitset in the model; fixtures carry it as a small int.
  function asStatus(o: AhpSkeleton.Option<Wire.Json>): bv32 {
    var n := asInt(o); if 0 <= n < 0x1_0000_0000 then n as bv32 else 0
  }

  // ── decode a customization (Cust). JSON `type` → model `kind`. `state` and
  //    `channel` are the mcpServer state-machine fields; children/load/clientId
  //    are unmodeled (open-struct — dropped identically on both sides). ──
  function decodeCust(j: Wire.Json): Cust {
    Cust(asStr(field(j, "id")), asStr(field(j, "type")), asStr(field(j, "uri")),
         asStr(field(j, "name")), optBool(field(j, "enabled")),
         optJsonNoNull(field(j, "state")), optJsonNoNull(field(j, "channel")))
  }
  function decodeCusts(js: seq<Wire.Json>): seq<Cust>
  { CO.Map(decodeCust, js) }

  // ── decode a Chat summary. origin carried opaquely; modifiedAt is real state
  //    (explicit in fixtures — see the header note on the normalization). ──
  function decodeChat(j: Wire.Json): Chat {
    Chat(asStr(field(j, "resource")), asStr(field(j, "title")), asInt(field(j, "status")),
         optStr(field(j, "activity")), asStr(field(j, "modifiedAt")),
         if field(j, "origin").Some? then field(j, "origin").value else Wire.JNull)
  }
  function decodeChats(js: seq<Wire.Json>): seq<Chat>
  { CO.Map(decodeChat, js) }

  // ── decode a Client / InputReq: identity key + full raw Wire.Json (reducer keeps
  //    the whole object; equality compares it). ──
  function decodeClient(j: Wire.Json): Client { Client(asStr(field(j, "clientId")), j) }
  function decodeClients(js: seq<Wire.Json>): seq<Client>
  { CO.Map(decodeClient, js) }

  function decodeInputReq(j: Wire.Json): InputReq { InputReq(asStr(field(j, "id")), j) }
  function decodeInputs(js: seq<Wire.Json>): seq<InputReq>
  { CO.Map(decodeInputReq, js) }

  // ── decode Option<SConfig> from a state `config` object ──
  function decodeConfigOpt(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<SConfig> {
    if o.Some? && o.value.JObj? then
      Some(SConfig(
        if field(o.value, "schema").Some? then field(o.value, "schema").value else Wire.JNull,
        asObjMap(field(o.value, "values"))))
    else None
  }

  // ── decode a full SessionState — every MODELED field; unmodeled keys ignored.
  //    NOTE: the model field `meta` is carried in the wire key `_meta`. ──
  function decodeSessionState(j: Wire.Json): SessionState {
    SessionState(
      asStr(field(j, "provider")),
      asStr(field(j, "title")),
      asStatus(field(j, "status")),
      asStr(field(j, "lifecycle")),
      optStr(field(j, "activity")),
      decodeConfigOpt(field(j, "config")),
      optJsonNoNull(field(j, "_meta")),
      optJsonNoNull(field(j, "creationError")),
      optJsonNoNull(field(j, "serverTools")),
      optJsonNoNull(field(j, "changesets")),
      optJsonNoNull(field(j, "canvases")),
      optJsonNoNull(field(j, "openCanvases")),
      optStr(field(j, "defaultChat")),
      decodeClients(asArr(field(j, "activeClients"))),
      decodeChats(asArr(field(j, "chats"))),
      decodeCusts(asArr(field(j, "customizations"))),
      decodeInputs(asArr(field(j, "inputNeeded"))))
  }

  // ── decode a SessionAction — dispatch on the "type" string; unknown → SUnknown ──
  function decodeSessionAction(j: Wire.Json): SessionAction {
    var t := asStr(field(j, "type"));
    if      t == "session/ready"                  then Ready
    else if t == "session/creationFailed"         then CreationFailed(if field(j, "error").Some? then field(j, "error").value else Wire.JNull)
    else if t == "session/titleChanged"           then TitleChanged(asStr(field(j, "title")))
    else if t == "session/serverToolsChanged"     then ServerToolsChanged(if field(j, "tools").Some? then field(j, "tools").value else Wire.JNull)
    else if t == "session/metaChanged"            then MetaChanged(if field(j, "_meta").Some? then field(j, "_meta").value else Wire.JNull)
    else if t == "session/isReadChanged"          then IsReadChanged(asBool(field(j, "isRead"), false))
    else if t == "session/isArchivedChanged"      then IsArchivedChanged(asBool(field(j, "isArchived"), false))
    else if t == "session/activityChanged"        then ActivityChanged(optStr(field(j, "activity")))
    else if t == "session/configChanged"          then ConfigChanged(asObjMap(field(j, "config")), asBool(field(j, "replace"), false))
    else if t == "session/customizationsChanged"  then CustomizationsChanged(decodeCusts(asArr(field(j, "customizations"))))
    else if t == "session/customizationToggled"   then CustomizationToggled(asStr(field(j, "id")), asBool(field(j, "enabled"), false))
    else if t == "session/customizationRemoved"   then CustomizationRemoved(asStr(field(j, "id")))
    else if t == "session/customizationUpdated"   then CustomizationUpdated(decodeCust(if field(j, "customization").Some? then field(j, "customization").value else Wire.JNull))
    else if t == "session/defaultChatChanged"     then DefaultChatChanged(optStr(field(j, "defaultChat")))
    else if t == "session/chatAdded"              then ChatAdded(decodeChat(if field(j, "summary").Some? then field(j, "summary").value else Wire.JNull))
    else if t == "session/chatRemoved"            then ChatRemoved(asStr(field(j, "chat")))
    else if t == "session/chatUpdated"            then
      var ch := field(j, "changes");
      var chObj := if ch.Some? then ch.value else Wire.JNull;
      ChatUpdated(asStr(field(j, "chat")), optInt(field(chObj, "status")), optStr(field(chObj, "activity")), optStr(field(chObj, "modifiedAt")))
    else if t == "session/changesetsChanged"      then ChangesetsChanged(optJsonNoNull(field(j, "changesets")))
    else if t == "session/canvasesChanged"        then CanvasesChanged(if field(j, "canvases").Some? then field(j, "canvases").value else Wire.JArr([]))
    else if t == "session/openCanvasesChanged"    then OpenCanvasesChanged(if field(j, "openCanvases").Some? then field(j, "openCanvases").value else Wire.JArr([]))
    else if t == "session/activeClientSet"        then ActiveClientSet(decodeClient(if field(j, "activeClient").Some? then field(j, "activeClient").value else Wire.JNull))
    else if t == "session/activeClientRemoved"    then ActiveClientRemoved(asStr(field(j, "clientId")))
    else if t == "session/inputNeededSet"         then InputNeededSet(decodeInputReq(if field(j, "request").Some? then field(j, "request").value else Wire.JNull))
    else if t == "session/inputNeededRemoved"     then InputNeededRemoved(asStr(field(j, "id")))
    else if t == "session/mcpServerStateChanged"  then McpServerStateChanged(asStr(field(j, "id")), if field(j, "state").Some? then field(j, "state").value else Wire.JNull, optJsonNoNull(field(j, "channel")))
    else if t == "session/mcpServerStartRequested" then McpServerStartRequested(asStr(field(j, "id")))
    else if t == "session/mcpServerStopRequested"  then McpServerStopRequested(asStr(field(j, "id")))
    else SUnknown(j)
  }
  function decodeActions(js: seq<Wire.Json>): seq<SessionAction>
  { CO.Map(decodeSessionAction, js) }

  // ── fold the ACTUAL reducer over the decoded action list (now injected) ──
  function foldCh(s: SessionState, acts: seq<SessionAction>, now: int): SessionState
  { CC.Fold((st, a) => ApplyToSession(st, a, now).0, s, acts) }

  // ── replay one fixture end-to-end. returns (isSession, ok): isSession false
  //    for non-session-reducer fixtures so Main does not count them. ──
  method replayOne(path: string) returns (isSession: bool, ok: bool) {
    var r := FileIO.ReadBytesFromFile(path);
    if r.Failure? { return false, false; }
    var p := API.Deserialize(toJsonBytes(r.value));
    if p.Failure? { return false, false; }
    var doc := bridge(p.value);
    if !doc.JObj? { return false, false; }
    if asStr(field(doc, "reducer")) != "session" { return false, false; }  // skip non-session fixtures
    var initial := decodeSessionState(if field(doc, "initial").Some? then field(doc, "initial").value else Wire.JNull);
    var actions := decodeActions(asArr(field(doc, "actions")));
    var expected := decodeSessionState(if field(doc, "expected").Some? then field(doc, "expected").value else Wire.JNull);
    var reduced := foldCh(initial, actions, 9999);
    isSession := true;
    ok := reduced == expected;   // NO normalization — see header note.
  }

  method Main(args: seq<string>) {
    var pass := 0;
    var total := 0;
    var i := 1;
    while i < |args|
      decreases |args| - i
    {
      var isSession, ok := replayOne(args[i]);
      if isSession {
        total := total + 1;
        if ok { pass := pass + 1; } else { print "REPLAY FAIL: ", args[i], "\n"; }
      }
      i := i + 1;
    }
    print "SESSION REDUCER-REPLAY: ", pass, "/", total, " real session fixtures replayed through the reducer, reduced == expected\n";
    if pass == total {
      print "SESSION-REPLAY OK — real JSON fixtures drive the proven reducer to the expected state (extracted code)\n";
    } else {
      print "SESSION-REPLAY: ", (total - pass), " fixture(s) diverge from the reducer (see REPLAY FAIL lines above)\n";
    }
  }
}
