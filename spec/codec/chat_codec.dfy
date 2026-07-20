// AHP Dafny client — FULL Wire.Json ⇄ ChatAction codec (Chat channel).
//
// The exact analogue of Fidelity's Root codec (spec/fidelity.dfy), for the Chat
// channel instead of Root. DECODE is promoted from the PROVEN chat replay decoder
// (spec/replay/chat_replay.dfy `decodeChatAction`) that drives that channel's 97
// real fixtures; ENCODE is its inverse (a canonical wire object per typed variant).
// Together they carry the load-bearing transport property: decode ∘ encode == id
// on every TYPED variant (ChatActionRoundTrip below), with the CUnknown leaf carved
// out honestly (its raw payload could re-classify — same shape as RootUnknown).
//
// The shared field accessors (field/asStr/asArr/asObjMap/asInt/asBool/strList/
// encodeStrs/StrsRoundTrip) are REUSED from Fidelity (F.*), not duplicated. Only the
// Chat-specific optional-field decoders (optStr/optBool/optJson/rawJson) — which the
// Root channel doesn't need — are defined locally, matching the replay decoder verbatim.
//
// Verifies standalone: `dafny verify spec/codec/chat_codec.dfy` (no Std, no
// --standard-libraries — chat.dfy and fidelity.dfy both bottom out at ahp_skeleton.dfy).
include "../chat.dfy"
include "../fidelity.dfy"

module ChatCodec {
  // ── .doo FIREWALL: explicit export set (replaces Dafny's default export-all) ──
  // Sole consumer = spec/ahp.dfy (import Cc = ChatCodec). It UNFOLDS EncodeChatAction
  // in the ChatPrefix lemma (per-variant "type"-literal reasoning) → EncodeChatAction is
  // REVEALED; the Wf* family is unfolded in AhpStateRoundTrip_NonVacuous's witness
  // (WFTurns→WFTurn→WFParts→WFPart / okJson / WfChatState) → all REVEALED. DecodeChatAction
  // and encode/decodeChatState are called as OPAQUE values (never unfolded — the RT lemmas
  // discharge the equalities) → provides. The Encode* leaves named by EncodeChatAction's
  // revealed body (EncodePart/EncodeInputReq/EncodeTurns/jStrOpt/jJsonOpt) + WFToolCall
  // (named by WFPart's revealed body) are provides (bodies hidden, same rule as
  // ahp_skeleton's `provides ApplyToRoot`). Chat/Wire/F/AhpSkeleton are provided because
  // exported signatures name their types (ChatAction/ChatState/Turn/…, Wire.Json, Option)
  // and EncodeChatAction's revealed body names F.encodeStrs. Every sub-codec (ToolCall/Part/
  // Turn/QMsg encode+decode), per-level RT lemma, optStr/optBool/rawJson/asStatus helper,
  // and Decode* state helper is HIDDEN — the firewall bites (e.g. ChatCodec.EncodeToolCall).
  //
  // COMPANION CHANGE (required): ahp.dfy's ChatPrefix lost its {:isolate_assertions} — that
  // attribute split ChatPrefix into 25 per-case VCs, and a per-case VC cannot unfold this
  // 25-arm EncodeChatAction across the explicit-export reveal boundary (the whole-lemma VC
  // can). No export set works with {:isolate_assertions} intact; the monolithic VC is green.
  export
    provides DecodeChatAction, encodeChatState, decodeChatState
    provides ChatActionRoundTrip, ChatStateRoundTrip
    provides EncodePart, EncodeInputReq, EncodeTurns, EncodeContributor, jStrOpt, jBoolOpt, jArrOpt, jJsonOpt
    provides WFToolCall
    provides Chat, Wire, F, AhpSkeleton
    reveals EncodeChatAction
    reveals okJson, WFContributor, WFInputReq, WFChatAction, WFPart, WFParts, WFTurn, WFTurns, WfChatState

  import opened Chat
  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import F = Fidelity   // shared accessors (field/asStr/asArr/asObjMap/asBool/strList/encodeStrs)
  import CO = ConfluxOperators

  // ── Chat-specific optional-field decoders (verbatim from chat_replay.dfy;
  //    Root doesn't need these, so Fidelity doesn't define them) ──
  function optStr(o: Option<Wire.Json>): Option<string> { if o.Some? && o.value.JStr? then Some(o.value.s) else None }
  function optBool(o: Option<Wire.Json>): Option<bool> { if o.Some? && o.value.JBool? then Some(o.value.b) else None }
  function optArr(o: Option<Wire.Json>): Option<seq<Wire.Json>> { if o.Some? && o.value.JArr? then Some(o.value.elems) else None }
  // Option<Wire.Json>: present AND not null → Some(raw); absent/null → None.
  function optJson(o: Option<Wire.Json>): Option<Wire.Json> { if o.Some? && !o.value.JNull? then Some(o.value) else None }
  // raw Wire.Json field: present → value; absent → Wire.JNull.
  function rawJson(o: Option<Wire.Json>): Wire.Json { if o.Some? then o.value else Wire.JNull }

  // ── ENCODE helpers for optional values. KEY INSIGHT: encode a None as Wire.JNull
  //    (never omit the key), so decode's optStr/optBool/optJson recover it — a
  //    Wire.JNull is not Wire.JStr/Wire.JBool, and is collapsed to None by optJson. This keeps
  //    the wire object a FIXED key set (no conditional keys), which is what makes
  //    the round-trip proofs tractable. ──
  function jStrOpt(o: Option<string>): Wire.Json { if o.Some? then Wire.JStr(o.value) else Wire.JNull }
  function jBoolOpt(o: Option<bool>): Wire.Json { if o.Some? then Wire.JBool(o.value) else Wire.JNull }
  function jArrOpt(o: Option<seq<Wire.Json>>): Wire.Json { if o.Some? then Wire.JArr(o.value) else Wire.JNull }
  function jJsonOpt(o: Option<Wire.Json>): Wire.Json { if o.Some? then o.value else Wire.JNull }

  function DecodeContributor(o: Option<Wire.Json>): Option<ToolCallContributor> {
    if o.Some? && o.value.JObj? then Some(ToolCallContributor(F.asStr(F.field(o.value, "kind")), o.value)) else None
  }
  function EncodeContributor(o: Option<ToolCallContributor>): Wire.Json { if o.Some? then o.value.raw else Wire.JNull }

  // ── Well-formedness carve-out: the ONLY value that can't round-trip through an
  //    Option<Wire.Json> wire field is Some(Wire.JNull) — on the wire it is indistinguishable
  //    from absent (optJson collapses Wire.JNull → None). This is the honest wire-fidelity
  //    property (analogous to the RootUnknown re-classification carve-out): a
  //    Some(Wire.JNull) content/usage/error/answer is not representable distinctly. ──
  predicate okJson(o: Option<Wire.Json>) { !(o.Some? && o.value.JNull?) }
  predicate WFContributor(o: Option<ToolCallContributor>) {
    o.None? || (o.value.raw.JObj? && F.asStr(F.field(o.value.raw, "kind")) == o.value.kind)
  }
  predicate WFInputReq(r: InputReq) { "id" !in r.open && "answers" !in r.open }
  predicate WFToolCall(tc: ToolCall) {
    WFContributor(tc.contributor) && okJson(tc.meta) && okJson(tc.confirmationTitle) && okJson(tc.riskAssessment) &&
    okJson(tc.edits) && okJson(tc.reasonMessage) && okJson(tc.userSuggestion) && okJson(tc.selectedOption) &&
    okJson(tc.content) && okJson(tc.structuredContent) && okJson(tc.error) && okJson(tc.auth)
    // NO `options` carve-out: `options` is a Json ARRAY field, round-tripped VERBATIM by jArrOpt/optArr
    // (encode Some(xs)→JArr(xs)→decode Some(xs); None→JNull→None), so a null array ELEMENT round-trips
    // faithfully — unlike the scalar Option<Json> fields above, whose encoder can't distinguish Some(JNull)
    // from absent (hence okJson). The reducer's `selectedOption := selectOption(options, id)` write — the one
    // place an options element crosses into a scalar field — null-collapses via Chat.selectOption, so no null
    // reaches a scalar there either. ToolCallRT therefore round-trips ANY options array (null elements included).
  }
  predicate WFPart(p: Part) {
    (p.PToolCall? ==> WFToolCall(p.tc)) && (p.PInputRequest? ==> WFInputReq(p.req))
  }
  predicate WFParts(ps: seq<Part>) { forall i :: 0 <= i < |ps| ==> WFPart(ps[i]) }
  predicate WFTurn(t: Turn) { WFParts(t.parts) && okJson(t.usage) && okJson(t.error) }
  predicate WFTurns(ts: seq<Turn>) { forall i :: 0 <= i < |ts| ==> WFTurn(ts[i]) }

  // ════════════════════════════════════════════════════════════════════════
  //  Sub-codecs: ToolCall, Part, Turn, InputReq (each with its round-trip lemma).
  // ════════════════════════════════════════════════════════════════════════

  // ── ToolCall (authorization- and confirmation-complete typed projection) ──
  function DecodeToolCall(j: Wire.Json): ToolCall {
    ToolCall(
      F.asStr(F.field(j, "toolCallId")),
      F.asStr(F.field(j, "toolName")),
      F.asStr(F.field(j, "displayName")),
      F.asStr(F.field(j, "status")),
      optStr(F.field(j, "intention")),
      DecodeContributor(F.field(j, "contributor")),
      optJson(F.field(j, "_meta")),
      F.asStr(F.field(j, "invocationMessage")),
      optStr(F.field(j, "toolInput")),
      optJson(F.field(j, "confirmationTitle")),
      optJson(F.field(j, "riskAssessment")),
      optJson(F.field(j, "edits")),
      optBool(F.field(j, "editable")),
      optArr(F.field(j, "options")),
      optStr(F.field(j, "confirmed")),
      optBool(F.field(j, "success")),
      optStr(F.field(j, "pastTenseMessage")),
      optStr(F.field(j, "reason")),
      optJson(F.field(j, "reasonMessage")),
      optJson(F.field(j, "userSuggestion")),
      optJson(F.field(j, "selectedOption")),
      optJson(F.field(j, "content")),
      optJson(F.field(j, "structuredContent")),
      optJson(F.field(j, "error")),
      optJson(F.field(j, "auth")),
      optStr(F.field(j, "partialInput")))
  }
  function EncodeToolCall(tc: ToolCall): Wire.Json {
    Wire.JObj(map[
      "toolCallId" := Wire.JStr(tc.toolCallId),
      "toolName" := Wire.JStr(tc.toolName),
      "displayName" := Wire.JStr(tc.displayName),
      "status" := Wire.JStr(tc.status),
      "intention" := jStrOpt(tc.intention),
      "contributor" := EncodeContributor(tc.contributor),
      "_meta" := jJsonOpt(tc.meta),
      "invocationMessage" := Wire.JStr(tc.invocationMessage),
      "toolInput" := jStrOpt(tc.toolInput),
      "confirmationTitle" := jJsonOpt(tc.confirmationTitle),
      "riskAssessment" := jJsonOpt(tc.riskAssessment),
      "edits" := jJsonOpt(tc.edits),
      "editable" := jBoolOpt(tc.editable),
      "options" := jArrOpt(tc.options),
      "confirmed" := jStrOpt(tc.confirmed),
      "success" := jBoolOpt(tc.success),
      "pastTenseMessage" := jStrOpt(tc.pastTenseMessage),
      "reason" := jStrOpt(tc.reason),
      "reasonMessage" := jJsonOpt(tc.reasonMessage),
      "userSuggestion" := jJsonOpt(tc.userSuggestion),
      "selectedOption" := jJsonOpt(tc.selectedOption),
      "content" := jJsonOpt(tc.content),
      "structuredContent" := jJsonOpt(tc.structuredContent),
      "error" := jJsonOpt(tc.error),
      "auth" := jJsonOpt(tc.auth),
      "partialInput" := jStrOpt(tc.partialInput)])
  }
  lemma {:isolate_assertions} ToolCallRT(tc: ToolCall)
    requires WFToolCall(tc)
    ensures DecodeToolCall(EncodeToolCall(tc)) == tc
  {
    var d := DecodeToolCall(EncodeToolCall(tc));
    assert d.toolCallId == tc.toolCallId;
    assert d.toolName == tc.toolName;
    assert d.displayName == tc.displayName;
    assert d.status == tc.status;
    assert d.intention == tc.intention;
    if tc.contributor.Some? {
      assert d.contributor == tc.contributor;
    } else {
      assert d.contributor == tc.contributor;
    }
    assert d.meta == tc.meta;
    assert d.invocationMessage == tc.invocationMessage;
    assert d.toolInput == tc.toolInput;
    assert d.confirmationTitle == tc.confirmationTitle;
    assert d.riskAssessment == tc.riskAssessment;
    assert d.edits == tc.edits;
    assert d.editable == tc.editable;
    assert d.options == tc.options;
    assert d.confirmed == tc.confirmed;
    assert d.success == tc.success;
    assert d.pastTenseMessage == tc.pastTenseMessage;
    assert d.reason == tc.reason;
    assert d.reasonMessage == tc.reasonMessage;
    assert d.userSuggestion == tc.userSuggestion;
    assert d.selectedOption == tc.selectedOption;
    assert d.content == tc.content;
    assert d.structuredContent == tc.structuredContent;
    assert d.error == tc.error;
    assert d.auth == tc.auth;
    assert d.partialInput == tc.partialInput;
    assert d == tc;
  }

  // Non-vacuity for WFToolCall: a REAL tool call carrying a populated `options` array
  // satisfies it AND round-trips (ToolCallRT) — so WFToolCall is not options-absent-only,
  // and the round-trip guarantee genuinely covers tool calls that carry confirmation options.
  lemma WFToolCall_NonVacuous()
    ensures exists tc: ToolCall :: (WFToolCall(tc) && tc.options.Some? && |tc.options.value| > 0)
  {
    var opt := Wire.JObj(map["id" := Wire.JStr("approve")]);
    var tc := ToolCall("tc-1", "bash", "Run", "pending-confirmation", None, None, None, "", None, None, None, None,
                       None, Some([opt]), None, None, None, None, None, None, None, None, None, None, None, None);
    assert WFToolCall(tc);
    ToolCallRT(tc);
  }

  // Hardening witness: a tool call whose `options` array CONTAINS a JNull element — the exact
  // input the dropped precondition used to exclude — is still WFToolCall and STILL round-trips.
  // This is the concrete evidence that dropping the options-non-null conjunct broadened (not
  // weakened) the guarantee: the array (JNull element and all) is carried through verbatim.
  lemma ToolCallRT_NullOptionElement()
    ensures exists tc: ToolCall :: WFToolCall(tc) && tc.options.Some? && (exists i :: 0 <= i < |tc.options.value| && tc.options.value[i].JNull?)
                                   && DecodeToolCall(EncodeToolCall(tc)) == tc
  {
    var opt := Wire.JObj(map["id" := Wire.JStr("approve")]);
    var tc := ToolCall("tc-1", "bash", "Run", "pending-confirmation", None, None, None, "", None, None, None, None,
                       None, Some([Wire.JNull, opt]), None, None, None, None, None, None, None, None, None, None, None, None);
    assert WFToolCall(tc);
    ToolCallRT(tc);
    assert tc.options.value[0].JNull?;
  }

  // ── Part (markdown | reasoning | toolCall) ──
  function DecodePart(j: Wire.Json): Part {
    var k := F.asStr(F.field(j, "kind"));
    if k == "toolCall" then PToolCall(DecodeToolCall(rawJson(F.field(j, "toolCall"))))
    else if k == "inputRequest" then PInputRequest(DecodeInputReq(rawJson(F.field(j, "request"))), optStr(F.field(j, "response")))   // #338: `response` absent → None (open); present → Some(kind)
    else if k == "reasoning" then PReasoning(F.asStr(F.field(j, "id")), F.asStr(F.field(j, "content")))
    else PMarkdown(F.asStr(F.field(j, "id")), F.asStr(F.field(j, "content")))
  }
  function EncodePart(p: Part): Wire.Json {
    match p
    case PMarkdown(id, content)  => Wire.JObj(map["kind" := Wire.JStr("markdown"),  "id" := Wire.JStr(id), "content" := Wire.JStr(content)])
    case PReasoning(id, content) => Wire.JObj(map["kind" := Wire.JStr("reasoning"), "id" := Wire.JStr(id), "content" := Wire.JStr(content)])
    case PToolCall(tc)           => Wire.JObj(map["kind" := Wire.JStr("toolCall"), "toolCall" := EncodeToolCall(tc)])
    // #338: OMIT `response` when the request is still open (None) — on the wire an open request has no
    // `response` key; a resolved one carries its kind. optStr recovers None from the absent key.
    case PInputRequest(req, response) =>
      var base := map["kind" := Wire.JStr("inputRequest"), "request" := EncodeInputReq(req)];
      Wire.JObj(if response.Some? then base["response" := Wire.JStr(response.value)] else base)
  }
  lemma PartRT(p: Part)
    requires WFPart(p)
    ensures DecodePart(EncodePart(p)) == p
  {
    match p
    case PMarkdown(_, _)  =>
    case PReasoning(_, _) =>
    case PToolCall(tc)    => ToolCallRT(tc);
    case PInputRequest(req, response) =>
      InputReqRT(req);
      // decode reads "request" and "response"; the request map never carries a "response" key (WFInputReq
      // excludes "id"/"answers" only, but "response" is a Part-level sibling of "request", not inside it),
      // so the two decode fields invert their encoders for both response=None (key omitted) and Some.
      assert DecodePart(EncodePart(p)) == p;
  }

  function EncodeParts(ps: seq<Part>): seq<Wire.Json>
  { CO.Map(EncodePart, ps) }
  function DecodeParts(js: seq<Wire.Json>): seq<Part>
  { CO.Map(DecodePart, js) }
  lemma PartsRT(ps: seq<Part>)
    requires WFParts(ps)
    ensures DecodeParts(EncodeParts(ps)) == ps
  {
    forall i | 0 <= i < |ps| { PartRT(ps[i]); }
  }

  // ── Turn (turnId + message + ordered parts + state + usage/error) ──
  function DecodeTurn(j: Wire.Json): Turn {
    Turn(
      F.asStr(F.field(j, "id")),
      rawJson(F.field(j, "message")),
      DecodeParts(F.asArr(F.field(j, "responseParts"))),
      if F.field(j, "state").Some? && F.field(j, "state").value.JStr? then F.field(j, "state").value.s else "in-progress",
      optJson(F.field(j, "usage")),
      optJson(F.field(j, "error")))
  }
  function EncodeTurn(t: Turn): Wire.Json {
    Wire.JObj(map[
      "id" := Wire.JStr(t.turnId),
      "message" := t.message,
      "responseParts" := Wire.JArr(EncodeParts(t.parts)),
      "state" := Wire.JStr(t.state),
      "usage" := jJsonOpt(t.usage),
      "error" := jJsonOpt(t.error)])
  }
  lemma TurnRT(t: Turn)
    requires WFTurn(t)
    ensures DecodeTurn(EncodeTurn(t)) == t
  { PartsRT(t.parts); }

  function EncodeTurns(ts: seq<Turn>): seq<Wire.Json>
  { CO.Map(EncodeTurn, ts) }
  function DecodeTurns(js: seq<Wire.Json>): seq<Turn>
  { CO.Map(DecodeTurn, js) }
  lemma TurnsRT(ts: seq<Turn>)
    requires WFTurns(ts)
    ensures DecodeTurns(EncodeTurns(ts)) == ts
  {
    forall i | 0 <= i < |ts| { TurnRT(ts[i]); }
  }

  // ── InputReq: typed id/answers plus lossless open request fields ──
  function DecodeInputReq(j: Wire.Json): InputReq {
    InputReq(F.asStr(F.field(j, "id")), F.asObjMap(F.field(j, "answers")), F.asObjMap(Some(j)) - {"id", "answers"})
  }
  function EncodeInputReq(r: InputReq): Wire.Json {
    Wire.JObj(r.open["id" := Wire.JStr(r.id)]["answers" := Wire.JObj(r.answers)])
  }
  lemma InputReqRT(r: InputReq)
    requires WFInputReq(r)
    ensures DecodeInputReq(EncodeInputReq(r)) == r
  {}

  // ════════════════════════════════════════════════════════════════════════
  //  TOP-LEVEL codec: dispatch on the "type" string (DECODE promoted verbatim
  //  from chat_replay.dfy `decodeChatAction`; ENCODE the canonical inverse).
  // ════════════════════════════════════════════════════════════════════════
  function DecodeChatAction(j: Wire.Json): ChatAction {
    var t := F.asStr(F.field(j, "type"));
    if      t == "chat/draftChanged"    then CDraftChanged(optStr(F.field(j, "draft")))
    else if t == "chat/activityChanged" then CActivityChanged(optStr(F.field(j, "activity")))
    else if t == "chat/turnStarted"     then CTurnStarted(F.asStr(F.field(j, "turnId")), rawJson(F.field(j, "message")), optStr(F.field(j, "queuedMessageId")))
    else if t == "chat/turnComplete"    then CTurnComplete(F.asStr(F.field(j, "turnId")))
    else if t == "chat/turnCancelled"   then CTurnCancelled(F.asStr(F.field(j, "turnId")))
    else if t == "chat/usage"           then CUsage(F.asStr(F.field(j, "turnId")), rawJson(F.field(j, "usage")))
    else if t == "chat/error"           then CError(F.asStr(F.field(j, "turnId")), rawJson(F.field(j, "error")))
    else if t == "chat/responsePart"    then CResponsePart(F.asStr(F.field(j, "turnId")), DecodePart(rawJson(F.field(j, "part"))))
    else if t == "chat/delta"           then CDelta(F.asStr(F.field(j, "turnId")), F.asStr(F.field(j, "partId")), F.asStr(F.field(j, "content")))
    else if t == "chat/reasoning"       then CReasoning(F.asStr(F.field(j, "turnId")), F.asStr(F.field(j, "partId")), F.asStr(F.field(j, "content")))
    else if t == "chat/toolCallStart"   then
      CToolCallStart(F.asStr(F.field(j, "turnId")), F.asStr(F.field(j, "toolCallId")), F.asStr(F.field(j, "toolName")),
                     F.asStr(F.field(j, "displayName")), optStr(F.field(j, "intention")), DecodeContributor(F.field(j, "contributor")), optJson(F.field(j, "_meta")))
    else if t == "chat/toolCallDelta"   then
      CToolCallDelta(F.asStr(F.field(j, "turnId")), F.asStr(F.field(j, "toolCallId")), F.asStr(F.field(j, "content")),
                     optStr(F.field(j, "invocationMessage")), optJson(F.field(j, "_meta")))
    else if t == "chat/toolCallReady"   then
      CToolCallReady(F.asStr(F.field(j, "turnId")), F.asStr(F.field(j, "toolCallId")), optStr(F.field(j, "invocationMessage")),
                     optStr(F.field(j, "toolInput")), optJson(F.field(j, "confirmationTitle")), optJson(F.field(j, "riskAssessment")),
                     optJson(F.field(j, "edits")), optBool(F.field(j, "editable")), optArr(F.field(j, "options")),
                     optStr(F.field(j, "confirmed")), optJson(F.field(j, "_meta")))
    else if t == "chat/toolCallConfirmed" then
      CToolCallConfirmed(F.asStr(F.field(j, "turnId")), F.asStr(F.field(j, "toolCallId")),
                         F.asBool(F.field(j, "approved"), true), optStr(F.field(j, "confirmed")),
                         optStr(F.field(j, "reason")), optJson(F.field(j, "reasonMessage")), optJson(F.field(j, "userSuggestion")),
                         optStr(F.field(j, "editedToolInput")), optStr(F.field(j, "selectedOptionId")), optJson(F.field(j, "_meta")))
    else if t == "chat/toolCallAuthRequired" then
      CToolCallAuthRequired(F.asStr(F.field(j, "turnId")), F.asStr(F.field(j, "toolCallId")), rawJson(F.field(j, "auth")), optJson(F.field(j, "_meta")))
    else if t == "chat/toolCallAuthResolved" then
      CToolCallAuthResolved(F.asStr(F.field(j, "turnId")), F.asStr(F.field(j, "toolCallId")), optJson(F.field(j, "_meta")))
    else if t == "chat/toolCallComplete" then
      var result := rawJson(F.field(j, "result"));
      CToolCallComplete(F.asStr(F.field(j, "turnId")), F.asStr(F.field(j, "toolCallId")),
                        F.asBool(F.field(result, "success"), false), optStr(F.field(result, "pastTenseMessage")),
                        optJson(F.field(result, "content")), optJson(F.field(result, "structuredContent")), optJson(F.field(result, "error")),
                        F.asBool(F.field(j, "requiresResultConfirmation"), false), optJson(F.field(j, "_meta")))
    else if t == "chat/toolCallResultConfirmed" then
      CToolCallResultConfirmed(F.asStr(F.field(j, "turnId")), F.asStr(F.field(j, "toolCallId")), F.asBool(F.field(j, "approved"), false), optJson(F.field(j, "_meta")))
    else if t == "chat/toolCallContentChanged" then
      CToolCallContentChanged(F.asStr(F.field(j, "turnId")), F.asStr(F.field(j, "toolCallId")), rawJson(F.field(j, "content")), optJson(F.field(j, "_meta")))
    else if t == "chat/truncated"       then CTruncated(optStr(F.field(j, "turnId")))
    else if t == "chat/queuedMessagesReordered" then CQueuedReordered(F.strList(F.asArr(F.field(j, "order"))))
    else if t == "chat/turnsLoaded"     then
      CTurnsLoaded(DecodeTurns(F.asArr(F.field(j, "turns"))), optStr(F.field(j, "turnsNextCursor")))
    else if t == "chat/inputRequested"  then CInputRequested(DecodeInputReq(rawJson(F.field(j, "request"))))
    else if t == "chat/inputAnswerChanged" then
      CInputAnswerChanged(F.asStr(F.field(j, "requestId")), F.asStr(F.field(j, "questionId")), optJson(F.field(j, "answer")))
    else if t == "chat/inputCompleted"  then CInputCompleted(F.asStr(F.field(j, "requestId")), F.asStr(F.field(j, "response")), F.asObjMap(F.field(j, "answers")))
    else if t == "chat/pendingMessageSet" then
      CPendingMessageSet(F.asStr(F.field(j, "kind")), F.asStr(F.field(j, "id")), rawJson(F.field(j, "message")))
    else if t == "chat/pendingMessageRemoved" then
      CPendingMessageRemoved(F.asStr(F.field(j, "kind")), F.asStr(F.field(j, "id")))
    else CUnknown(j)
  }

  function EncodeChatAction(a: ChatAction): Wire.Json {
    match a
    case CDraftChanged(d)     => Wire.JObj(map["type" := Wire.JStr("chat/draftChanged"), "draft" := jStrOpt(d)])
    case CActivityChanged(ac) => Wire.JObj(map["type" := Wire.JStr("chat/activityChanged"), "activity" := jStrOpt(ac)])
    case CTurnStarted(id, m, qid) =>
      Wire.JObj(map["type" := Wire.JStr("chat/turnStarted"), "turnId" := Wire.JStr(id), "message" := m, "queuedMessageId" := jStrOpt(qid)])
    case CTurnComplete(id)    => Wire.JObj(map["type" := Wire.JStr("chat/turnComplete"), "turnId" := Wire.JStr(id)])
    case CTurnCancelled(id)   => Wire.JObj(map["type" := Wire.JStr("chat/turnCancelled"), "turnId" := Wire.JStr(id)])
    case CUsage(id, u)        => Wire.JObj(map["type" := Wire.JStr("chat/usage"), "turnId" := Wire.JStr(id), "usage" := u])
    case CError(id, e)        => Wire.JObj(map["type" := Wire.JStr("chat/error"), "turnId" := Wire.JStr(id), "error" := e])
    case CResponsePart(id, p) => Wire.JObj(map["type" := Wire.JStr("chat/responsePart"), "turnId" := Wire.JStr(id), "part" := EncodePart(p)])
    case CDelta(id, pid, c)   => Wire.JObj(map["type" := Wire.JStr("chat/delta"), "turnId" := Wire.JStr(id), "partId" := Wire.JStr(pid), "content" := Wire.JStr(c)])
    case CReasoning(id, pid, c) => Wire.JObj(map["type" := Wire.JStr("chat/reasoning"), "turnId" := Wire.JStr(id), "partId" := Wire.JStr(pid), "content" := Wire.JStr(c)])
    case CToolCallStart(id, tcId, tn, dn, intent, contributor, meta) =>
      Wire.JObj(map["type" := Wire.JStr("chat/toolCallStart"), "turnId" := Wire.JStr(id), "toolCallId" := Wire.JStr(tcId),
               "toolName" := Wire.JStr(tn), "displayName" := Wire.JStr(dn), "intention" := jStrOpt(intent), "contributor" := EncodeContributor(contributor), "_meta" := jJsonOpt(meta)])
    case CToolCallDelta(id, tcId, c, inv, meta) =>
      Wire.JObj(map["type" := Wire.JStr("chat/toolCallDelta"), "turnId" := Wire.JStr(id), "toolCallId" := Wire.JStr(tcId),
               "content" := Wire.JStr(c), "invocationMessage" := jStrOpt(inv), "_meta" := jJsonOpt(meta)])
    case CToolCallReady(id, tcId, inv, input, title, risk, edits, editable, options, confirmed, meta) =>
      Wire.JObj(map["type" := Wire.JStr("chat/toolCallReady"), "turnId" := Wire.JStr(id), "toolCallId" := Wire.JStr(tcId),
               "invocationMessage" := jStrOpt(inv), "toolInput" := jStrOpt(input), "confirmationTitle" := jJsonOpt(title),
               "riskAssessment" := jJsonOpt(risk), "edits" := jJsonOpt(edits), "editable" := jBoolOpt(editable),
               "options" := jArrOpt(options), "confirmed" := jStrOpt(confirmed), "_meta" := jJsonOpt(meta)])
    case CToolCallConfirmed(id, tcId, approved, cv, rsn, reasonMessage, userSuggestion, edited, selectedOptionId, meta) =>
      Wire.JObj(map["type" := Wire.JStr("chat/toolCallConfirmed"), "turnId" := Wire.JStr(id), "toolCallId" := Wire.JStr(tcId),
               "approved" := Wire.JBool(approved), "confirmed" := jStrOpt(cv), "reason" := jStrOpt(rsn),
               "reasonMessage" := jJsonOpt(reasonMessage), "userSuggestion" := jJsonOpt(userSuggestion),
               "editedToolInput" := jStrOpt(edited), "selectedOptionId" := jStrOpt(selectedOptionId), "_meta" := jJsonOpt(meta)])
    case CToolCallAuthRequired(id, tcId, auth, meta) =>
      Wire.JObj(map["type" := Wire.JStr("chat/toolCallAuthRequired"), "turnId" := Wire.JStr(id), "toolCallId" := Wire.JStr(tcId), "auth" := auth, "_meta" := jJsonOpt(meta)])
    case CToolCallAuthResolved(id, tcId, meta) =>
      Wire.JObj(map["type" := Wire.JStr("chat/toolCallAuthResolved"), "turnId" := Wire.JStr(id), "toolCallId" := Wire.JStr(tcId), "_meta" := jJsonOpt(meta)])
    case CToolCallComplete(id, tcId, ok, past, resultContent, structured, err, rrc, meta) =>
      Wire.JObj(map["type" := Wire.JStr("chat/toolCallComplete"), "turnId" := Wire.JStr(id), "toolCallId" := Wire.JStr(tcId),
               "result" := Wire.JObj(map["success" := Wire.JBool(ok), "pastTenseMessage" := jStrOpt(past), "content" := jJsonOpt(resultContent), "structuredContent" := jJsonOpt(structured), "error" := jJsonOpt(err)]),
               "requiresResultConfirmation" := Wire.JBool(rrc), "_meta" := jJsonOpt(meta)])
    case CToolCallResultConfirmed(id, tcId, approved, meta) =>
      Wire.JObj(map["type" := Wire.JStr("chat/toolCallResultConfirmed"), "turnId" := Wire.JStr(id), "toolCallId" := Wire.JStr(tcId), "approved" := Wire.JBool(approved), "_meta" := jJsonOpt(meta)])
    case CToolCallContentChanged(id, tcId, c, meta) =>
      Wire.JObj(map["type" := Wire.JStr("chat/toolCallContentChanged"), "turnId" := Wire.JStr(id), "toolCallId" := Wire.JStr(tcId), "content" := c, "_meta" := jJsonOpt(meta)])
    case CTruncated(idOpt)    => Wire.JObj(map["type" := Wire.JStr("chat/truncated"), "turnId" := jStrOpt(idOpt)])
    case CQueuedReordered(order) => Wire.JObj(map["type" := Wire.JStr("chat/queuedMessagesReordered"), "order" := Wire.JArr(F.encodeStrs(order))])
    case CTurnsLoaded(loaded, cursor) =>
      Wire.JObj(map["type" := Wire.JStr("chat/turnsLoaded"), "turns" := Wire.JArr(EncodeTurns(loaded)), "turnsNextCursor" := jStrOpt(cursor)])
    case CInputRequested(req) => Wire.JObj(map["type" := Wire.JStr("chat/inputRequested"), "request" := EncodeInputReq(req)])
    case CInputAnswerChanged(reqId, qId, ans) =>
      Wire.JObj(map["type" := Wire.JStr("chat/inputAnswerChanged"), "requestId" := Wire.JStr(reqId), "questionId" := Wire.JStr(qId), "answer" := jJsonOpt(ans)])
    case CInputCompleted(reqId, resp, answers) => Wire.JObj(map["type" := Wire.JStr("chat/inputCompleted"), "requestId" := Wire.JStr(reqId), "response" := Wire.JStr(resp), "answers" := Wire.JObj(answers)])
    case CPendingMessageSet(kind, id, msg) =>
      Wire.JObj(map["type" := Wire.JStr("chat/pendingMessageSet"), "kind" := Wire.JStr(kind), "id" := Wire.JStr(id), "message" := msg])
    case CPendingMessageRemoved(kind, id) =>
      Wire.JObj(map["type" := Wire.JStr("chat/pendingMessageRemoved"), "kind" := Wire.JStr(kind), "id" := Wire.JStr(id)])
    case CUnknown(raw)        => raw
  }

  // ════════════════════════════════════════════════════════════════════════
  //  ROUND-TRIP: decode ∘ encode == id on every TYPED variant. The three
  //  variants carrying Option<Wire.Json> wire fields (CResponsePart via ToolCall.content,
  //  CTurnsLoaded via Turn.usage/error/parts, CInputAnswerChanged via answer) carry
  //  a Some(Wire.JNull)-exclusion precondition (okJson) — the honest wire-fidelity
  //  carve-out. CUnknown is excluded outright (its raw could re-classify — the
  //  RootUnknown carve-out), handled by the verbatim lemmas below.
  // ════════════════════════════════════════════════════════════════════════
  // Json→domain→Json fixpoint: the codec's canonical wire form re-decodes + re-encodes to ITSELF — the
  // direction the conformance suite exercises, now machine-checked. Follows from the direction-1 round-trip.
  predicate WFChatAction(a: ChatAction) {
    (a.CToolCallStart? ==> WFContributor(a.contributor) && okJson(a.meta)) &&
    (a.CToolCallDelta? ==> okJson(a.meta)) &&
    (a.CToolCallReady? ==> okJson(a.confirmationTitle) && okJson(a.riskAssessment) && okJson(a.edits) && okJson(a.meta)) &&
    (a.CToolCallConfirmed? ==> okJson(a.reasonMessage) && okJson(a.userSuggestion) && okJson(a.meta)) &&
    (a.CToolCallAuthRequired? ==> okJson(a.meta)) &&
    (a.CToolCallAuthResolved? ==> okJson(a.meta)) &&
    (a.CToolCallComplete? ==> okJson(a.resultContent) && okJson(a.structuredContent) && okJson(a.error) && okJson(a.meta)) &&
    (a.CToolCallResultConfirmed? ==> okJson(a.meta)) &&
    (a.CToolCallContentChanged? ==> okJson(a.meta)) &&
    (a.CInputRequested? ==> WFInputReq(a.req))
  }
  lemma ChatWireCanonicalRoundTrip(a: ChatAction)
    requires !a.CUnknown?
    requires WFChatAction(a)
    requires a.CResponsePart? ==> WFPart(a.part)
    requires a.CTurnsLoaded? ==> WFTurns(a.loaded)
    requires a.CInputAnswerChanged? ==> okJson(a.answer)
    ensures EncodeChatAction(DecodeChatAction(EncodeChatAction(a))) == EncodeChatAction(a)
  { ChatActionRoundTrip(a); }
  lemma {:isolate_assertions} ChatActionRoundTrip(a: ChatAction)
    requires !a.CUnknown?
    requires WFChatAction(a)
    requires a.CResponsePart? ==> WFPart(a.part)
    requires a.CTurnsLoaded? ==> WFTurns(a.loaded)
    requires a.CInputAnswerChanged? ==> okJson(a.answer)
    ensures DecodeChatAction(EncodeChatAction(a)) == a
  {
    match a
    case CDraftChanged(_) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CActivityChanged(_) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CTurnStarted(_, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CTurnComplete(_) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CTurnCancelled(_) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CUsage(_, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CError(_, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CResponsePart(_, p) => PartRT(p); assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CDelta(_, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CReasoning(_, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CToolCallStart(_, _, _, _, _, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CToolCallDelta(_, _, _, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CToolCallReady(_, _, _, _, _, _, _, _, _, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CToolCallConfirmed(_, _, _, _, _, _, _, _, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CToolCallAuthRequired(_, _, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CToolCallAuthResolved(_, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CToolCallComplete(_, _, _, _, _, _, _, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CToolCallResultConfirmed(_, _, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CToolCallContentChanged(_, _, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CTruncated(_) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CQueuedReordered(order) => F.StrsRoundTrip(order); assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CTurnsLoaded(loaded, _) => TurnsRT(loaded); assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CInputRequested(req) => InputReqRT(req); assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CInputAnswerChanged(_, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CInputCompleted(_, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CPendingMessageSet(_, _, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CPendingMessageRemoved(_, _) => assert DecodeChatAction(EncodeChatAction(a)) == a;
    case CUnknown(_) =>   // excluded by precondition; branch is unreachable (vacuous)
  }

  // Non-vacuity witness: the round-trip precondition set is inhabited by a genuine
  // typed action (a non-trivial tool-call complete, exercising the nested "result"
  // object) — so ChatActionRoundTrip is NOT vacuously true / constants-only.
  lemma ChatActionRoundTrip_NonVacuous()
    ensures exists a: ChatAction :: !a.CUnknown? && DecodeChatAction(EncodeChatAction(a)) == a
  {
    var a := CToolCallComplete("turn-1", "tc-1", true, Some("Ran command"), None, None, None, false, None);
    ChatActionRoundTrip(a);
    assert !a.CUnknown? && DecodeChatAction(EncodeChatAction(a)) == a;
  }

  // ── unknown-variant passthrough — the CUnknown leaf re-encodes verbatim
  //    (the honest carve-out; identical in shape to Fidelity.RootUnknown_EncodeVerbatim) ──
  lemma ChatUnknown_EncodeVerbatim(j: Wire.Json)
    ensures EncodeChatAction(CUnknown(j)) == j {}
  lemma UnknownVariant_Verbatim(j: Wire.Json)
    requires DecodeChatAction(j).CUnknown?
    ensures EncodeChatAction(DecodeChatAction(j)) == j {}

  // ════════════════════════════════════════════════════════════════════════
  //  STATE serialization (the analogue of Fidelity.encodeRootState): a
  //  DETERMINISTIC TOTAL serializer of the full ChatState to canonical JSON,
  //  for the cross-process transport demo (host + client each emit final state;
  //  convergence is by determinism, not a state round-trip lemma). Every MODELED
  //  field is serialized; optional fields are omitted when None.
  // ════════════════════════════════════════════════════════════════════════
  function encodeQMsg(q: QMsg): Wire.Json { Wire.JObj(map["id" := Wire.JStr(q.id), "message" := q.raw]) }
  function encodeQMsgs(qs: seq<QMsg>): seq<Wire.Json>
  { CO.Map(encodeQMsg, qs) }

  // #338: no `inputRequests` state field — input requests live as PInputRequest parts inside the turns /
  // active turn (serialized within `responseParts` by EncodeTurn), mirroring the TS ChatState.
  function encodeChatState(s: ChatState): Wire.Json {
    var m0 := map[
      "turns" := Wire.JArr(EncodeTurns(s.turns)),
      "title" := Wire.JStr(s.title),
      "status" := Wire.JNum(s.status as int),
      "modifiedAt" := Wire.JStr(s.modifiedAt),
      "queuedMessages" := Wire.JArr(encodeQMsgs(s.queuedMessages))];
    var m1 := if s.draft.Some? then m0["draft" := Wire.JStr(s.draft.value)] else m0;
    var m2 := if s.activity.Some? then m1["activity" := Wire.JStr(s.activity.value)] else m1;
    var m3 := if s.activeTurn.Some? then m2["activeTurn" := EncodeTurn(s.activeTurn.value)] else m2;
    var m4 := if s.steeringMessage.Some? then m3["steeringMessage" := encodeQMsg(s.steeringMessage.value)] else m3;
    var m5 := if s.nextCursor.Some? then m4["turnsNextCursor" := Wire.JStr(s.nextCursor.value)] else m4;
    Wire.JObj(m5)
  }

  // ════════════════════════════════════════════════════════════════════════
  //  STATE codec (decode) — the TRUE INVERSE of encodeChatState above, promoted
  //  from the PROVEN chat replay decoder (spec/replay/chat_replay.dfy
  //  `decodeChatState`, which drives that channel's real fixtures) and re-expressed
  //  over the shared Fidelity accessors (F.*). This completes the STATE proof story
  //  in BOTH directions: encodeChatState was one-directional; ChatStateRoundTrip
  //  below proves decodeChatState ∘ encodeChatState == id (the analogue of the
  //  ACTION codec's ChatActionRoundTrip). Where encodeChatState OMITS an optional
  //  key when None, decode recovers None from the absent key (optStr/Decode* → None);
  //  where it emits the key it recovers Some — a true inverse on the FIXED-optional
  //  layout that encode produces.
  // ════════════════════════════════════════════════════════════════════════

  // status is a bv32; guard the int→bv32 conversion so the verifier proves range
  // (identical to chat_replay.dfy `asStatus`; encodeChatState emits Wire.JNum(status as int)).
  function asStatus(o: Option<Wire.Json>): bv32 {
    var n := asInt(o);
    if 0 <= n < 0x1_0000_0000 then n as bv32 else 0
  }

  // ── QMsg { id, message } — `message` is opaque Wire.Json (round-trips verbatim, even Wire.JNull) ──
  function DecodeQMsg(j: Wire.Json): QMsg { QMsg(F.asStr(F.field(j, "id")), rawJson(F.field(j, "message"))) }
  function DecodeQMsgs(js: seq<Wire.Json>): seq<QMsg>
  { CO.Map(DecodeQMsg, js) }
  lemma QMsgRT(q: QMsg)
    ensures DecodeQMsg(encodeQMsg(q)) == q
  {}
  lemma QMsgsRT(qs: seq<QMsg>)
    ensures DecodeQMsgs(encodeQMsgs(qs)) == qs
  { forall i | 0 <= i < |qs| { QMsgRT(qs[i]); } }

  // ── optional object fields: present-and-Wire.JObj → Some(decode); absent → None. encode
  //    emits a Wire.JObj (EncodeTurn / encodeQMsg both do) when Some, omits the key when None. ──
  function DecodeActiveTurn(o: Option<Wire.Json>): Option<Turn> {
    if o.Some? && o.value.JObj? then Some(DecodeTurn(o.value)) else None
  }
  function DecodeSteering(o: Option<Wire.Json>): Option<QMsg> {
    if o.Some? && o.value.JObj? then Some(DecodeQMsg(o.value)) else None
  }

  function decodeChatState(j: Wire.Json): ChatState {
    ChatState(
      DecodeTurns(F.asArr(F.field(j, "turns"))),
      F.asStr(F.field(j, "title")),
      asStatus(F.field(j, "status")),
      F.asStr(F.field(j, "modifiedAt")),
      optStr(F.field(j, "draft")),
      optStr(F.field(j, "activity")),
      DecodeActiveTurn(F.field(j, "activeTurn")),
      DecodeSteering(F.field(j, "steeringMessage")),
      DecodeQMsgs(F.asArr(F.field(j, "queuedMessages"))),
      optStr(F.field(j, "turnsNextCursor")))
  }

  // ── Well-formedness carve-out for the STATE round-trip. IDENTICAL in spirit to
  //    the ACTION codec's per-variant okJson/WFTurns preconditions: the ONLY fields
  //    that cannot round-trip are the Turn.usage/error/ToolCall.content Option<Wire.Json>
  //    wire fields carrying Some(Wire.JNull) (on the wire indistinguishable from absent —
  //    optJson collapses Wire.JNull → None). Every OTHER ChatState field round-trips
  //    UNCONDITIONALLY: title/modifiedAt (Wire.JStr), status (bv32⇄int, range-guarded),
  //    draft/activity/nextCursor (Option<string>, omit-when-None, never Wire.JNull),
  //    queuedMessages/steeringMessage (QMsg — `message` is opaque, Wire.JNull round-trips
  //    verbatim). #338: input requests are carried inside turns/activeTurn as PInputRequest parts, so
  //    their well-formedness (WFInputReq) is subsumed by WFTurns / WFTurn (WFPart covers PInputRequest).
  //    So the precondition is exactly WFTurns on the turn list plus WFTurn on the optional activeTurn. ──
  predicate WfChatState(s: ChatState) {
    WFTurns(s.turns) && (s.activeTurn.Some? ==> WFTurn(s.activeTurn.value))
  }

  // Projection lemmas deliberately keep each codec obligation in a fresh,
  // minimal solver context.  The flat state theorem below composes these
  // facts; no query has to normalize unrelated optional-map updates or nested
  // turn codecs merely to prove one field.
  lemma ChatStateTurnsRT(s: ChatState)
    requires WFTurns(s.turns)
    ensures decodeChatState(encodeChatState(s)).turns == s.turns
  { TurnsRT(s.turns); }
  lemma ChatStateTitleRT(s: ChatState)
    ensures decodeChatState(encodeChatState(s)).title == s.title
  {}
  lemma {:isolate_assertions} ChatStateStatusRT(s: ChatState)
    ensures decodeChatState(encodeChatState(s)).status == s.status
  {
    StatusRoundTrip(s.status);
    assert 0 <= s.status as int < 0x1_0000_0000;
    assert asStatus(Some(Wire.JNum(s.status as int))) == s.status;
  }
  lemma ChatStateModifiedAtRT(s: ChatState)
    ensures decodeChatState(encodeChatState(s)).modifiedAt == s.modifiedAt
  {}
  lemma ChatStateDraftRT(s: ChatState)
    ensures decodeChatState(encodeChatState(s)).draft == s.draft
  { if s.draft.Some? {} else {} }
  lemma ChatStateActivityRT(s: ChatState)
    ensures decodeChatState(encodeChatState(s)).activity == s.activity
  { if s.activity.Some? {} else {} }
  lemma ChatStateActiveTurnRT(s: ChatState)
    requires s.activeTurn.Some? ==> WFTurn(s.activeTurn.value)
    ensures decodeChatState(encodeChatState(s)).activeTurn == s.activeTurn
  { if s.activeTurn.Some? { TurnRT(s.activeTurn.value); } }
  lemma ChatStateSteeringRT(s: ChatState)
    ensures decodeChatState(encodeChatState(s)).steeringMessage == s.steeringMessage
  { if s.steeringMessage.Some? { QMsgRT(s.steeringMessage.value); } }
  lemma ChatStateQueuedRT(s: ChatState)
    ensures decodeChatState(encodeChatState(s)).queuedMessages == s.queuedMessages
  { QMsgsRT(s.queuedMessages); }
  lemma ChatStateCursorRT(s: ChatState)
    ensures decodeChatState(encodeChatState(s)).nextCursor == s.nextCursor
  { if s.nextCursor.Some? {} else {} }

  // ════════════════════════════════════════════════════════════════════════
  //  STATE ROUND-TRIP: decode ∘ encode == id on a well-formed ChatState — the
  //  STATE-side analogue of ChatActionRoundTrip. REAL, quantified, non-vacuous
  //  (witness below). Proven by discharging the seq inductions (TurnsRT/QMsgsRT/
  //  InputReqsRT) + the two optional-object element RTs + the bv32 range fact; the
  //  fixed-optional wire layout (encode always omits a None key, never emits Wire.JNull
  //  for it) makes each field lookup invert its encoder.
  // ════════════════════════════════════════════════════════════════════════
  // {:isolate_assertions} splits this proof's VC per-assertion so no single
  // solver query approaches the default 30s time-limit — ChatState is the
  // deepest state (turns→parts→toolcalls), and without the split the monolithic
  // VC flakes at the boundary. The split keeps it reliably green at any limit.
  lemma {:isolate_assertions} ChatStateRoundTrip(s: ChatState)
    requires WfChatState(s)
    ensures decodeChatState(encodeChatState(s)) == s
  {
    ChatStateTurnsRT(s);
    ChatStateTitleRT(s);
    ChatStateStatusRT(s);
    ChatStateModifiedAtRT(s);
    ChatStateDraftRT(s);
    ChatStateActivityRT(s);
    ChatStateActiveTurnRT(s);
    ChatStateSteeringRT(s);
    ChatStateQueuedRT(s);
    ChatStateCursorRT(s);
    var decoded := decodeChatState(encodeChatState(s));
    assert decoded == s;
  }

  // Non-vacuity witness: a genuinely NON-EMPTY, well-formed ChatState (a finalized
  // turn with a markdown part + a set draft + an active turn + a queued message +
  // an input request + a cursor) satisfies the precondition and round-trips — so
  // ChatStateRoundTrip is NOT vacuously true / empty-state-only.
  lemma ChatStateRoundTrip_NonVacuous()
    ensures exists s: ChatState :: WfChatState(s) && decodeChatState(encodeChatState(s)) == s
  {
    var t := Turn("turn-1", Wire.JStr("hello"), [PMarkdown("p1", "hi there")], "completed", None, None);
    assert WFTurn(t);
    // #338: the active turn carries a LIVE (open) input-request part — the state's input requests now live
    // here, so the witness exercises the PInputRequest codec path (response=None ⇒ no "response" key on wire).
    var at := Turn("turn-2", Wire.JStr("more"),
      [PInputRequest(InputReq("req-1", map["question1" := Wire.JStr("answer")], map["message" := Wire.JStr("Clarify")]), None)],
      "in-progress", None, None);
    assert WFTurn(at) by { assert WFParts(at.parts) by { forall i | 0 <= i < |at.parts| ensures WFPart(at.parts[i]) { assert at.parts[i] == at.parts[0]; } } }
    var s := ChatState(
      [t], "My Chat", GEN, "2026-01-01T00:00:00Z",
      Some("draft text"), None, Some(at),
      None, [QMsg("q1", Wire.JStr("queued msg"))], Some("cursor-1"));
    assert WFTurns(s.turns) by { forall i | 0 <= i < |s.turns| ensures WFTurn(s.turns[i]) { assert s.turns[i] == t; } }
    assert WfChatState(s);
    ChatStateRoundTrip(s);
    assert WfChatState(s) && decodeChatState(encodeChatState(s)) == s;
  }

  ghost predicate ChatCodecWf(a: ChatAction) {
    !a.CUnknown?
    && WFChatAction(a)
    && (a.CResponsePart? ==> WFPart(a.part))
    && (a.CTurnsLoaded? ==> WFTurns(a.loaded))
    && (a.CInputAnswerChanged? ==> okJson(a.answer))
  }
  function ChatCodecC(): Wire.Codec<ChatAction> {
    Wire.Codec(EncodeChatAction, (j: Wire.Json) => Wire.Some(DecodeChatAction(j)))
  }
  lemma ChatCodecRoundTripsWhere()
    ensures forall a :: ChatCodecWf(a) ==> ChatCodecC().decode(ChatCodecC().encode(a)) == Wire.Some(a)
  {
    forall a | ChatCodecWf(a)
      ensures ChatCodecC().decode(ChatCodecC().encode(a)) == Wire.Some(a)
    { ChatActionRoundTrip(a); }
  }
}
