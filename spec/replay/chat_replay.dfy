// AHP Dafny client — REDUCER REPLAY harness (Phase 2), chat channel.
//
// Mirrors root_replay.dfy / terminal_replay.dfy: instead of hand-transcribing
// each fixture, it REPLAYS the real chat fixture JSON through the ACTUAL reducer:
//
//   read bytes → Std.JSON parse → bridge to skeleton Wire.Json → decode initial/
//   actions/expected → foldCh(ApplyToChat) → assert reduced == decoded-expected
//
// The chat reducer models `modifiedAt` as an opaque string but does NOT thread
// the per-action ISO timestamp (the #186 clock concern — a display field, not
// reducer state-transition logic). So both sides are normalized to a constant
// modifiedAt := "N" before comparison. That is the ONLY normalization; every
// other modeled field (turns, activeTurn, status bv32, parts — including the #338
// live input-request parts, draft, activity, steeringMessage, queuedMessages,
// nextCursor) is compared for real.
//
// Only fixtures whose "reducer" field == "chat" are counted (the glob may match
// other channels' fixtures — replayOne returns isCh=false for them, like
// terminal_replay's reducer filter).
// Build/run: needs --standard-libraries; pass the fixture paths as args.
include "../chat.dfy"

module ChatReplay {
  import opened Chat
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

  // ── bridge Std.JSON.Values.JSON → the skeleton's Wire.Json — VERBATIM ──
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

  // ── field accessors on skeleton Wire.Json — VERBATIM ──
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

  // ── optional-field decoders: present+typed → Some(value); else None ──
  function optStr(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<string> {
    if o.Some? && o.value.JStr? then Some(o.value.s) else None
  }
  function optBool(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<bool> {
    if o.Some? && o.value.JBool? then Some(o.value.b) else None
  }
  function optArr(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<seq<Wire.Json>> {
    if o.Some? && o.value.JArr? then Some(o.value.elems) else None
  }
  // Option<Wire.Json>: present AND not null → Some(raw); absent/null → None.
  function optJson(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<Wire.Json> {
    if o.Some? && !o.value.JNull? then Some(o.value) else None
  }
  // raw Wire.Json field: present → value; absent → Wire.JNull.
  function rawJson(o: AhpSkeleton.Option<Wire.Json>): Wire.Json {
    if o.Some? then o.value else Wire.JNull
  }
  function decodeContributor(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<ToolCallContributor> {
    if o.Some? && o.value.JObj? then Some(ToolCallContributor(asStr(field(o.value, "kind")), o.value)) else None
  }
  // status is a bv32; guard the int→bv32 conversion so the verifier proves range.
  function asStatus(o: AhpSkeleton.Option<Wire.Json>): bv32 {
    var n := asInt(o);
    if 0 <= n < 0x1_0000_0000 then n as bv32 else 0
  }

  // ── decode a ToolCall (12 modeled fields; contributor/_meta/etc. are open-struct) ──
  function decodeToolCall(j: Wire.Json): ToolCall {
    ToolCall(
      asStr(field(j, "toolCallId")),
      asStr(field(j, "toolName")),
      asStr(field(j, "displayName")),
      asStr(field(j, "status")),
      optStr(field(j, "intention")),
      decodeContributor(field(j, "contributor")),
      optJson(field(j, "_meta")),
      asStr(field(j, "invocationMessage")),
      optStr(field(j, "toolInput")),
      optJson(field(j, "confirmationTitle")),
      optJson(field(j, "riskAssessment")),
      optJson(field(j, "edits")),
      optBool(field(j, "editable")),
      optArr(field(j, "options")),
      optStr(field(j, "confirmed")),
      optBool(field(j, "success")),
      optStr(field(j, "pastTenseMessage")),
      optStr(field(j, "reason")),
      optJson(field(j, "reasonMessage")),
      optJson(field(j, "userSuggestion")),
      optJson(field(j, "selectedOption")),
      optJson(field(j, "content")),
      optJson(field(j, "structuredContent")),
      optJson(field(j, "error")),
      optJson(field(j, "auth")),
      optStr(field(j, "partialInput")))
  }

  // ── decode a content Part (markdown | reasoning | toolCall) ──
  function decodePart(j: Wire.Json): Part {
    var k := asStr(field(j, "kind"));
    if k == "toolCall" then PToolCall(decodeToolCall(rawJson(field(j, "toolCall"))))
    else if k == "inputRequest" then PInputRequest(decodeInputReq(rawJson(field(j, "request"))), optStr(field(j, "response")))   // #338: absent `response` → None (open)
    else if k == "reasoning" then PReasoning(asStr(field(j, "id")), asStr(field(j, "content")))
    else PMarkdown(asStr(field(j, "id")), asStr(field(j, "content")))  // "markdown" (and any other) → markdown
  }
  function decodeParts(js: seq<Wire.Json>): seq<Part>
  { CO.Map(decodePart, js) }

  // ── decode a Turn. `state` defaults to "in-progress" (live activeTurn wire has
  //    no state field); finalized turns in `turns[]` carry an explicit state. ──
  function decodeTurn(j: Wire.Json): Turn {
    Turn(
      asStr(field(j, "id")),
      rawJson(field(j, "message")),
      decodeParts(asArr(field(j, "responseParts"))),
      if field(j, "state").Some? && field(j, "state").value.JStr? then field(j, "state").value.s else "in-progress",
      optJson(field(j, "usage")),
      optJson(field(j, "error")))
  }
  function decodeTurns(js: seq<Wire.Json>): seq<Turn>
  { CO.Map(decodeTurn, js) }

  function decodeActiveTurn(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<Turn> {
    if o.Some? && o.value.JObj? then Some(decodeTurn(o.value)) else None
  }

  // ── QMsg (pending/steering/queued message): { id, message } ──
  function decodeQMsg(j: Wire.Json): QMsg { QMsg(asStr(field(j, "id")), rawJson(field(j, "message"))) }
  function decodeQMsgs(js: seq<Wire.Json>): seq<QMsg>
  { CO.Map(decodeQMsg, js) }
  function decodeSteering(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<QMsg> {
    if o.Some? && o.value.JObj? then Some(decodeQMsg(o.value)) else None
  }

  // ── InputReq: typed id/answers plus every retained open request field (carried inside a PInputRequest part) ──
  function decodeInputReq(j: Wire.Json): InputReq {
    InputReq(asStr(field(j, "id")), asObjMap(field(j, "answers")), asObjMap(Some(j)) - {"id", "answers"})
  }

  // ── decode the whole ChatState — every MODELED field. #338: no `inputRequests` field — input requests
  //    are PInputRequest parts inside the turns / active turn (decoded within `responseParts`). ──
  function decodeChatState(j: Wire.Json): ChatState {
    ChatState(
      decodeTurns(asArr(field(j, "turns"))),
      asStr(field(j, "title")),
      asStatus(field(j, "status")),
      asStr(field(j, "modifiedAt")),
      optStr(field(j, "draft")),
      optStr(field(j, "activity")),
      decodeActiveTurn(field(j, "activeTurn")),
      decodeSteering(field(j, "steeringMessage")),
      decodeQMsgs(asArr(field(j, "queuedMessages"))),
      optStr(field(j, "turnsNextCursor")))
  }

  // ── decode a ChatAction — dispatch on the "type" string ──
  function decodeChatAction(j: Wire.Json): ChatAction {
    var t := asStr(field(j, "type"));
    if      t == "chat/draftChanged"    then CDraftChanged(optStr(field(j, "draft")))
    else if t == "chat/activityChanged" then CActivityChanged(optStr(field(j, "activity")))
    else if t == "chat/turnStarted"     then CTurnStarted(asStr(field(j, "turnId")), rawJson(field(j, "message")), optStr(field(j, "queuedMessageId")))
    else if t == "chat/turnComplete"    then CTurnComplete(asStr(field(j, "turnId")))
    else if t == "chat/turnCancelled"   then CTurnCancelled(asStr(field(j, "turnId")))
    else if t == "chat/usage"           then CUsage(asStr(field(j, "turnId")), rawJson(field(j, "usage")))
    else if t == "chat/error"           then CError(asStr(field(j, "turnId")), rawJson(field(j, "error")))
    else if t == "chat/responsePart"    then CResponsePart(asStr(field(j, "turnId")), decodePart(rawJson(field(j, "part"))))
    else if t == "chat/delta"           then CDelta(asStr(field(j, "turnId")), asStr(field(j, "partId")), asStr(field(j, "content")))
    else if t == "chat/reasoning"       then CReasoning(asStr(field(j, "turnId")), asStr(field(j, "partId")), asStr(field(j, "content")))
    else if t == "chat/toolCallStart"   then
      CToolCallStart(asStr(field(j, "turnId")), asStr(field(j, "toolCallId")), asStr(field(j, "toolName")),
                     asStr(field(j, "displayName")), optStr(field(j, "intention")), decodeContributor(field(j, "contributor")), optJson(field(j, "_meta")))
    else if t == "chat/toolCallDelta"   then
      CToolCallDelta(asStr(field(j, "turnId")), asStr(field(j, "toolCallId")), asStr(field(j, "content")),
                     optStr(field(j, "invocationMessage")), optJson(field(j, "_meta")))
    else if t == "chat/toolCallReady"   then
      CToolCallReady(asStr(field(j, "turnId")), asStr(field(j, "toolCallId")), optStr(field(j, "invocationMessage")),
                     optStr(field(j, "toolInput")), optJson(field(j, "confirmationTitle")), optJson(field(j, "riskAssessment")),
                     optJson(field(j, "edits")), optBool(field(j, "editable")), optArr(field(j, "options")),
                     optStr(field(j, "confirmed")), optJson(field(j, "_meta")))
    else if t == "chat/toolCallConfirmed" then
      CToolCallConfirmed(asStr(field(j, "turnId")), asStr(field(j, "toolCallId")),
                         asBool(field(j, "approved"), true), optStr(field(j, "confirmed")),
                         optStr(field(j, "reason")), optJson(field(j, "reasonMessage")), optJson(field(j, "userSuggestion")),
                         optStr(field(j, "editedToolInput")), optStr(field(j, "selectedOptionId")), optJson(field(j, "_meta")))
    else if t == "chat/toolCallAuthRequired" then
      CToolCallAuthRequired(asStr(field(j, "turnId")), asStr(field(j, "toolCallId")), rawJson(field(j, "auth")), optJson(field(j, "_meta")))
    else if t == "chat/toolCallAuthResolved" then
      CToolCallAuthResolved(asStr(field(j, "turnId")), asStr(field(j, "toolCallId")), optJson(field(j, "_meta")))
    else if t == "chat/toolCallComplete" then
      var result := rawJson(field(j, "result"));
      CToolCallComplete(asStr(field(j, "turnId")), asStr(field(j, "toolCallId")),
                        asBool(field(result, "success"), false), optStr(field(result, "pastTenseMessage")),
                        optJson(field(result, "content")), optJson(field(result, "structuredContent")), optJson(field(result, "error")),
                        asBool(field(j, "requiresResultConfirmation"), false), optJson(field(j, "_meta")))
    else if t == "chat/toolCallResultConfirmed" then
      CToolCallResultConfirmed(asStr(field(j, "turnId")), asStr(field(j, "toolCallId")), asBool(field(j, "approved"), false), optJson(field(j, "_meta")))
    else if t == "chat/toolCallContentChanged" then
      CToolCallContentChanged(asStr(field(j, "turnId")), asStr(field(j, "toolCallId")), rawJson(field(j, "content")), optJson(field(j, "_meta")))
    else if t == "chat/truncated"       then CTruncated(optStr(field(j, "turnId")))
    else if t == "chat/queuedMessagesReordered" then CQueuedReordered(strList(asArr(field(j, "order"))))
    else if t == "chat/turnsLoaded"     then
      CTurnsLoaded(decodeTurns(asArr(field(j, "turns"))), optStr(field(j, "turnsNextCursor")))
    else if t == "chat/inputRequested"  then CInputRequested(decodeInputReq(rawJson(field(j, "request"))))
    else if t == "chat/inputAnswerChanged" then
      CInputAnswerChanged(asStr(field(j, "requestId")), asStr(field(j, "questionId")), optJson(field(j, "answer")))
    else if t == "chat/inputCompleted"  then CInputCompleted(asStr(field(j, "requestId")), asStr(field(j, "response")), asObjMap(field(j, "answers")))
    else if t == "chat/pendingMessageSet" then
      CPendingMessageSet(asStr(field(j, "kind")), asStr(field(j, "id")), rawJson(field(j, "message")))
    else if t == "chat/pendingMessageRemoved" then
      CPendingMessageRemoved(asStr(field(j, "kind")), asStr(field(j, "id")))
    else CUnknown(j)
  }
  function decodeActions(js: seq<Wire.Json>): seq<ChatAction>
  { CO.Map(decodeChatAction, js) }

  // ── fold the actions through the ACTUAL reducer (now injected; ApplyToChat
  //    ignores it for state transitions — modifiedAt is not threaded) ──
  function foldCh(s: ChatState, acts: seq<ChatAction>, now: int): ChatState
  { CC.Fold((st, a) => ApplyToChat(st, a, now).0, s, acts) }

  // ── replay one fixture end-to-end. Returns (isCh, ok): isCh=false for
  //    non-chat-reducer fixtures so Main does not count them. modifiedAt is
  //    normalized to "N" on both sides (the only allowed normalization). ──
  method replayOne(path: string) returns (isCh: bool, ok: bool) {
    var r := FileIO.ReadBytesFromFile(path);
    if r.Failure? { return false, false; }
    var p := API.Deserialize(toJsonBytes(r.value));
    if p.Failure? { return false, false; }
    var doc := bridge(p.value);
    if !doc.JObj? { return false, false; }
    if asStr(field(doc, "reducer")) != "chat" { return false, false; }
    var initial := decodeChatState(rawJson(field(doc, "initial")));
    var actions := decodeActions(asArr(field(doc, "actions")));
    var expected := decodeChatState(rawJson(field(doc, "expected")));
    var reduced := foldCh(initial, actions, 9999);
    isCh := true;
    ok := reduced.(modifiedAt := "N") == expected.(modifiedAt := "N");
    if !ok {
      print "  DIFF ", path, ":";
      if reduced.turns != expected.turns { print " turns"; }
      if reduced.activeTurn != expected.activeTurn { print " activeTurn"; }
      if reduced.status != expected.status { print " status(", reduced.status as int, "vs", expected.status as int, ")"; }
      if reduced.title != expected.title { print " title"; }
      if reduced.draft != expected.draft { print " draft"; }
      if reduced.activity != expected.activity { print " activity"; }
      if reduced.steeringMessage != expected.steeringMessage { print " steeringMessage"; }
      if reduced.queuedMessages != expected.queuedMessages { print " queuedMessages"; }
      if reduced.nextCursor != expected.nextCursor { print " nextCursor"; }
      print "\n";
    }
  }

  method Main(args: seq<string>) {
    var pass := 0;
    var total := 0;
    var i := 1;
    while i < |args|
      decreases |args| - i
    {
      var isCh, ok := replayOne(args[i]);
      if isCh {
        total := total + 1;
        if ok { pass := pass + 1; } else { print "REPLAY FAIL: ", args[i], "\n"; }
      }
      i := i + 1;
    }
    print "CHAT REDUCER-REPLAY: ", pass, "/", total, " real chat fixtures replayed through the reducer, reduced == expected (modifiedAt normalized)\n";
    if pass == total {
      print "CHAT-REPLAY OK — real JSON fixtures drive the proven reducer to the expected state (extracted code)\n";
    } else {
      print "CHAT-REPLAY PARTIAL — ", (total - pass), " fixture(s) diverge from the model; see REPLAY FAIL lines above\n";
    }
  }
}
