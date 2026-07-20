// AHP Dafny client — chat channel (7th; the largest at 115 pinned fixtures).
// Models the real turn/responsePart/tool-call lifecycle: a turn carries an
// ordered `parts` list (markdown | reasoning | tool-call), matching the wire
// `responseParts`. The tool-call state machine is the hardest logic in the
// protocol: start(streaming) → delta(accumulate partialInput) →
// ready(→pending-confirmation | →running when auto-confirm) →
// confirmed(→running/cancelled) → complete(→completed | →pending-result-
// confirmation) → resultConfirmed(→completed/cancelled); turnComplete
// force-cancels any still-in-progress tool call (→cancelled/skipped).
//
// `status` is a bv32 whose low activity bits (0–4) are DERIVED from live work
// via summaryStatus (Idle / Error / InProgress / InputNeeded), preserving the
// orthogonal IsRead flag (bit 5). This mirrors the canonical TS reducer
// (channels-chat/reducer.ts).
include "ahp_skeleton.dfy"

module Chat {
  // .doo export firewall: the reducer/codec/API surface only. provides the reducer fns
  // (opaque bodies — channel_laws.ChatIsChannel unfolds fold==CC.Fold the same way the
  // verified SessionIsChannel does via `provides fold`), C0/RunCorpus/GEN, and
  // Wire/CC/AhpSkeleton (transitive: exported signatures name Wire.Json / Option / ReduceOutcome).
  // reveals the 7 datatypes consumers construct / pattern-match / field-access. Internal
  // helpers (setBit/clearBit/summaryStatus/endTurn/mapTC/upsertInputPart/… + T0/TC0/run* harness) stay hidden.
  export
    provides ApplyToChat, apply1, C0, RunCorpus, GEN
    reveals fold   // channel_laws.ChatIsChannel asserts Chat.fold == CC.Fold(apply1,·) in its body
    provides AhpSkeleton, Wire, CC
    provides WfChatState, WfChatAction, Apply1PreservesWfChat, ChatWfWitness  // rung-4/6: the reducer keyed invariant + apply1-level preservation + concrete witness (ahp.dfy composes WfAhpStateInv over it)
    reveals ChatState, ChatAction, Turn, Part, ToolCall, ToolCallContributor, QMsg, InputReq

  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import CC = ConfluxContract  // the one fold — `fold` calls CC.Fold (delete-and-reuse, not a bridge)
  import Ops = ConfluxOperators
  import OL = ConfluxOrderedLog
  import SR = ConfluxSeqRoute  // the canonical ordered-keyed-seq primitives (upsert/remove/update[-first/all]-where by id)

  // A tool-call part. Fields track the transitions the reducer actually mutates.
  // `partialInput` is the streaming delta accumulator (distinct from `toolInput`,
  // which is set by ready/confirm); it is dropped on every non-streaming rebuild.
  datatype ToolCallContributor = ToolCallContributor(kind: string, raw: Wire.Json)

  datatype ToolCall = ToolCall(
    toolCallId: string, toolName: string, displayName: string, status: string,
    intention: Option<string>, contributor: Option<ToolCallContributor>, meta: Option<Wire.Json>,
    invocationMessage: string, toolInput: Option<string>, confirmationTitle: Option<Wire.Json>,
    riskAssessment: Option<Wire.Json>, edits: Option<Wire.Json>, editable: Option<bool>, options: Option<seq<Wire.Json>>,
    confirmed: Option<string>, success: Option<bool>, pastTenseMessage: Option<string>,
    reason: Option<string>, reasonMessage: Option<Wire.Json>, userSuggestion: Option<Wire.Json>,
    selectedOption: Option<Wire.Json>, content: Option<Wire.Json>, structuredContent: Option<Wire.Json>,
    error: Option<Wire.Json>, auth: Option<Wire.Json>, partialInput: Option<string>)

  datatype Part =
    | PMarkdown(id: string, content: string)
    | PReasoning(id: string, content: string)
    | PToolCall(tc: ToolCall)
    | PInputRequest(req: InputReq, response: Option<string>)   // #338: a LIVE input-request response part; `response` absent (None) while open, Some(kind) once resolved

  // A turn: user message + ordered response parts + lifecycle state + usage/error.
  datatype Turn = Turn(
    turnId: string, message: Wire.Json, parts: seq<Part>,
    state: string, usage: Option<Wire.Json>, error: Option<Wire.Json>)

  datatype QMsg = QMsg(id: string, raw: Wire.Json)
  // An outstanding input request. `open` retains every request field other
  // than the two typed keys (`id`, `answers`) so completion can record the
  // original message/questions/url/options/required flags losslessly.
  datatype InputReq = InputReq(id: string, answers: map<string, Wire.Json>, open: map<string, Wire.Json>)

  // #338: input requests are LIVE response parts in the active turn (PInputRequest), NOT a
  // separate side seq — mirroring the TS ChatState, which has no `inputRequests` field. An open
  // request lives as a PInputRequest(response=None) part; completion resolves it in place.
  datatype ChatState = ChatState(
    turns: seq<Turn>, title: string, status: bv32, modifiedAt: string,
    draft: Option<string>, activity: Option<string>, activeTurn: Option<Turn>,
    steeringMessage: Option<QMsg>, queuedMessages: seq<QMsg>, nextCursor: Option<string>)

  // status activity bits (mutually exclusive, bits 0–4) + orthogonal IsRead (bit 5).
  const IDLE: bv32 := 0x01
  const ERR:  bv32 := 0x02          // SessionStatus.Error (1 << 1)
  const GEN:  bv32 := 0x08          // SessionStatus.InProgress (1 << 3)
  const INP:  bv32 := 0x18          // SessionStatus.InputNeeded ((1<<3)|(1<<4))
  const READ: bv32 := 0x20          // SessionStatus.IsRead (1 << 5)
  const ACT:  bv32 := 0x1F          // activity mask ((1 << 5) - 1)
  function setBit(s: bv32, b: bv32): bv32 { s | b }
  function clearBit(s: bv32, b: bv32): bv32 { s & !b }

  datatype ChatAction =
    | CDraftChanged(draft: Option<string>)          // draftChanged
    | CActivityChanged(activity: Option<string>)    // activityChanged
    | CTurnStarted(turnId: string, message: Wire.Json, queuedMessageId: Option<string>)  // turnStarted → activeTurn; clears the auto-started pending message
    | CTurnComplete(turnId: string)                 // turnComplete → force-cancel + finalize (no-op on wrong id)
    | CTurnCancelled(turnId: string)                // turnCancelled → finalize with state "cancelled"
    | CUsage(turnId: string, usage: Wire.Json)           // usage → sets activeTurn.usage (no-op on wrong id)
    | CError(turnId: string, err: Wire.Json)             // error → finalize turn state "error", status Error
    | CResponsePart(turnId: string, part: Part)     // responsePart → append markdown/reasoning part
    | CDelta(turnId: string, partId: string, content: string)  // delta → accumulate content on a MARKDOWN part only
    | CToolCallStart(turnId: string, toolCallId: string, toolName: string, displayName: string, intention: Option<string>, contributor: Option<ToolCallContributor>, meta: Option<Wire.Json>)
    | CToolCallDelta(turnId: string, toolCallId: string, content: string, invocationMessage: Option<string>, meta: Option<Wire.Json>)
    | CToolCallReady(turnId: string, toolCallId: string, invocationMessage: Option<string>, toolInput: Option<string>, confirmationTitle: Option<Wire.Json>, riskAssessment: Option<Wire.Json>, edits: Option<Wire.Json>, editable: Option<bool>, options: Option<seq<Wire.Json>>, confirmed: Option<string>, meta: Option<Wire.Json>)
    | CToolCallConfirmed(turnId: string, toolCallId: string, approved: bool, confirmedVal: Option<string>, reason: Option<string>, reasonMessage: Option<Wire.Json>, userSuggestion: Option<Wire.Json>, editedToolInput: Option<string>, selectedOptionId: Option<string>, meta: Option<Wire.Json>)
    | CToolCallAuthRequired(turnId: string, toolCallId: string, auth: Wire.Json, meta: Option<Wire.Json>)
    | CToolCallAuthResolved(turnId: string, toolCallId: string, meta: Option<Wire.Json>)
    | CToolCallComplete(turnId: string, toolCallId: string, success: bool, pastTenseMessage: Option<string>, resultContent: Option<Wire.Json>, structuredContent: Option<Wire.Json>, error: Option<Wire.Json>, requiresResultConfirmation: bool, meta: Option<Wire.Json>)
    | CToolCallResultConfirmed(turnId: string, toolCallId: string, approved: bool, meta: Option<Wire.Json>)  // pending-result-confirmation → (approved) completed | (denied) cancelled
    | CToolCallContentChanged(turnId: string, toolCallId: string, newContent: Wire.Json, meta: Option<Wire.Json>)  // sets content on a RUNNING tool call
    | CReasoning(turnId: string, partId: string, content: string)  // reasoning → accumulate on a REASONING part only
    | CTruncated(upToTurnId: Option<string>)      // keep turns up to+incl turnId; None → clear all; drops active turn
    | CQueuedReordered(order: seq<string>)        // reorder queuedMessages; unlisted appended in original order
    | CTurnsLoaded(loaded: seq<Turn>, cursor: Option<string>)  // prepend older turns (dedup by id), set cursor
    | CInputRequested(req: InputReq)                                 // inputRequested → upsert request; recompute status
    | CInputAnswerChanged(requestId: string, questionId: string, answer: Option<Wire.Json>)  // set/clear a draft answer (no-op unknown request)
    | CInputCompleted(requestId: string, response: string, answers: map<string, Wire.Json>) // resolve into active-turn history; remove request (no-op unknown)
    | CPendingMessageSet(kind: string, id: string, message: Wire.Json)  // steering → steeringMessage; queued → queuedMessages upsert
    | CPendingMessageRemoved(kind: string, id: string)             // steering → clear-if-match; queued → remove by id
    | CUnknown(raw: Wire.Json)

  // ── keyed-collection key functions (ONE shared value each; docs/shapes/07) ──
  // queuedMessages is the ONE remaining id-keyed reducer seq; it routes through the ConfluxSeqRoute keyed
  // algebra by this keyOf value. Sharing ONE const (instead of an inline lambda per call site) makes the
  // helper bodies and the WfChatState / *PreservesUnique law citations reference the SAME function value, so
  // the kernel uniqueness laws' conclusions line up with the helper results syntactically (no lambda-identity gap).
  // (#338 removed the `inputRequests` seq and its rKey: input requests are now PInputRequest parts in the turn.)
  const qKey: QMsg -> string := (x: QMsg) => x.id

  // JNull-collapse: an action payload of type Wire.Json whose typed target is Option<Wire.Json> collapses an
  // incoming JNull to None (mirroring the state decoder's optJson / TS `T | undefined`, which never stores null).
  // Storing Some(JNull) would violate the codec okJson carve-out and diverge from the TS reducer.
  function jNoNull(j: Wire.Json): Option<Wire.Json> { if j.JNull? then None else Some(j) }

  // queued-message upsert / remove by id: delegate to the canonical ConfluxSeqRoute primitives.
  function upsertQ(qs: seq<QMsg>, m: QMsg): seq<QMsg>
  { SR.SeqUpsertById(qKey, qs, m) }
  function removeQ(qs: seq<QMsg>, id: string): seq<QMsg>
  { SR.SeqRemoveById(qKey, qs, id) }
  function containsQ(qs: seq<QMsg>, id: string): bool
    decreases |qs|
  { if |qs| == 0 then false else if qs[0].id == id then true else containsQ(qs[1..], id) }
  function findQ(qs: seq<QMsg>, id: string): seq<QMsg>
  {
    var found := Ops.Filter((q: QMsg) => q.id == id, qs);
    if |found| == 0 then [] else [found[0]]
  }
  // reorder: take the listed ids (present, de-duplicated) in order, then append
  // any queued messages NOT named in `order`, preserving their original order.
  function reorderQ(qs: seq<QMsg>, order: seq<string>): seq<QMsg>
  { OL.SeqReorderByKeys(qKey, order, qs) }
  // truncate: keep turns from the start up to AND INCLUDING the one matching id.
  function keepUpTo(ts: seq<Turn>, id: string): seq<Turn>
  { OL.SeqKeepThrough((t: Turn) => t.turnId, id, ts) }
  // turnsLoaded: prepend older turns not already present (dedup by id), preserving loaded order.
  predicate hasTurn(ts: seq<Turn>, id: string) { exists i :: 0 <= i < |ts| && ts[i].turnId == id }
  function dedupPrepend(loaded: seq<Turn>, existing: seq<Turn>): seq<Turn>
  { Ops.Filter((t: Turn) => !hasTurn(existing, t.turnId), loaded) + existing }
  // ── input requests as LIVE response parts (#338) ──
  // An input request lives as a PInputRequest part in the ACTIVE turn: OPEN while `response` is None,
  // resolved once `response` is Some. `isOpenInput` is the guard the keyed part operations run under
  // (matching the TS reducer's `findOpenInputRequestPart`, which filters `response === undefined`).
  predicate isOpenInput(p: Part) { p.PInputRequest? && p.response.None? }
  // Is there an OPEN input-request part with this id in the parts list? (recursive existence check,
  // same shape as anyPendingTC / containsQ — the TS `findOpenInputRequestPart` returning a match).
  function hasOpenReq(ps: seq<Part>, id: string): bool
    decreases |ps|
  { if |ps| == 0 then false
    else ((ps[0].PInputRequest? && ps[0].response.None? && ps[0].req.id == id) || hasOpenReq(ps[1..], id)) }
  // upsert-answers merge: the new request replaces the open part wholesale, EXCEPT it inherits the
  // existing part's answers when the new request brings none (TS: `answers: request.answers ?? existing`).
  function mergeInputReq(existing: InputReq, req: InputReq): InputReq
  { if |req.answers| > 0 then req else req.(answers := existing.answers) }
  // inputRequested: replace the FIRST open matching part in place (merging answers), else append a fresh
  // open part. The in-place replace delegates to the canonical ConfluxSeqRoute update-first-where primitive
  // (key = partKey, guard = isOpenInput); the append is a plain seq extension.
  function mergeInputPart(req: InputReq): Part -> Part
  { (p: Part) => match p case PInputRequest(existing, _) => PInputRequest(mergeInputReq(existing, req), None) case _ => p }
  function upsertInputPart(ps: seq<Part>, req: InputReq): seq<Part>
  { if hasOpenReq(ps, req.id)
    then SR.SeqUpdateByIdWhere(partKey, isOpenInput, ps, req.id, mergeInputPart(req))
    else ps + [PInputRequest(req, None)] }
  // answerChanged: Some(a) sets the draft on the open matching part; None removes the questionId key.
  // Update-first-where (guard = isOpenInput) — delegates to ConfluxSeqRoute.SeqUpdateByIdWhere.
  function answerInputPart(qId: string, ans: Option<Wire.Json>): Part -> Part
  { (p: Part) => match p case PInputRequest(r, resp) => PInputRequest(r.(answers := if ans.Some? then r.answers[qId := ans.value] else r.answers - {qId}), resp) case _ => p }
  function changeAnswerPart(ps: seq<Part>, reqId: string, qId: string, ans: Option<Wire.Json>): seq<Part>
  { SR.SeqUpdateByIdWhere(partKey, isOpenInput, ps, reqId, answerInputPart(qId, ans)) }
  // completed: resolve the open matching part IN PLACE — merge final answers over drafts and record the
  // response kind — so its stream position is stable (TS updates the part in place, never appends).
  // Completion answers are authoritative for matching keys; retained drafts fill keys omitted by the action.
  function mergeAnswers(drafts: map<string, Wire.Json>, completed: map<string, Wire.Json>): map<string, Wire.Json>
  { map k: string | k in drafts.Keys + completed.Keys :: if k in completed then completed[k] else drafts[k] }
  function completeInputPart(resp: string, answers: map<string, Wire.Json>): Part -> Part
  { (p: Part) => match p case PInputRequest(r, _) => PInputRequest(r.(answers := mergeAnswers(r.answers, answers)), Some(resp)) case _ => p }
  function completeAnswerPart(ps: seq<Part>, reqId: string, resp: string, answers: map<string, Wire.Json>): seq<Part>
  { SR.SeqUpdateByIdWhere(partKey, isOpenInput, ps, reqId, completeInputPart(resp, answers)) }

  predicate turnMatches(s: ChatState, id: string) { s.activeTurn.Some? && s.activeTurn.value.turnId == id }

  function optOr(o: Option<string>, d: string): string { if o.Some? then o.value else d }

  // ── status derivation (summaryStatus): the activity nibble is DERIVED from
  //    live work; IsRead and other high bits are preserved. ──
  function anyPendingTC(ps: seq<Part>): bool
    decreases |ps|
  { if |ps| == 0 then false
    else ((ps[0].PToolCall? && (ps[0].tc.status == "pending-confirmation" || ps[0].tc.status == "pending-result-confirmation" || ps[0].tc.status == "auth-required")) || anyPendingTC(ps[1..])) }
  function hasPendingTCState(s: ChatState): bool
  { s.activeTurn.Some? && anyPendingTC(s.activeTurn.value.parts) }
  // #338: any OPEN input-request part in the active turn drives InputNeeded (TS hasOpenInputRequest).
  function anyOpenInput(ps: seq<Part>): bool
    decreases |ps|
  { if |ps| == 0 then false
    else ((ps[0].PInputRequest? && ps[0].response.None?) || anyOpenInput(ps[1..])) }
  function hasOpenInputState(s: ChatState): bool
  { s.activeTurn.Some? && anyOpenInput(s.activeTurn.value.parts) }
  function activityBits(s: ChatState, isError: bool): bv32
  { if isError then ERR
    else if hasOpenInputState(s) || hasPendingTCState(s) then INP
    else if s.activeTurn.Some? then GEN
    else IDLE }
  function summaryStatus(s: ChatState, isError: bool): bv32
  { (s.status & !ACT) | activityBits(s, isError) }

  // append a text part (markdown/reasoning) to a turn.
  function appendPart(t: Turn, p: Part): Turn { t.(parts := t.parts + [p]) }

  // total id-projection over a part: pid for the two text variants, toolCallId for the tool-call variant.
  function partKey(p: Part): string
  { match p case PMarkdown(pid, _) => pid case PReasoning(pid, _) => pid case PToolCall(tc) => tc.toolCallId case PInputRequest(req, _) => req.id }

  // delta: accumulate content onto a MARKDOWN part by id (reasoning/toolcall left alone). A guarded
  // update-ALL-matches sweep — delegates to ConfluxSeqRoute.SeqUpdateAllWhere (guard = PMarkdown? ∧ pid==id).
  function deltaMarkdown(ps: seq<Part>, id: string, content: string): seq<Part>
  { SR.SeqUpdateAllWhere(partKey, (p: Part) => p.PMarkdown?, ps, id,
      (p: Part) => match p case PMarkdown(pid, c) => PMarkdown(pid, c + content) case _ => p) }
  // reasoning: accumulate content onto a REASONING part by id (markdown/toolcall left alone).
  function deltaReasoning(ps: seq<Part>, id: string, content: string): seq<Part>
  { SR.SeqUpdateAllWhere(partKey, (p: Part) => p.PReasoning?, ps, id,
      (p: Part) => match p case PReasoning(pid, c) => PReasoning(pid, c + content) case _ => p) }

  // map a tool-call part by id through a transformer (all matching tool-call parts).
  function mapTC(ps: seq<Part>, id: string, f: ToolCall -> ToolCall): seq<Part>
  { SR.SeqUpdateAllWhere(partKey, (p: Part) => p.PToolCall?, ps, id,
      (p: Part) => match p case PToolCall(tc) => PToolCall(f(tc)) case _ => p) }

  // delta on a streaming tool call: accumulate into partialInput (NOT toolInput).
  function metaOr(meta: Option<Wire.Json>, prior: Option<Wire.Json>): Option<Wire.Json>
  { if meta.Some? then meta else prior }
  function optionId(j: Wire.Json): string
  { if j.JObj? && "id" in j.fields && j.fields["id"].JStr? then j.fields["id"].s else "" }
  // Selecting a confirmation option projects an ARRAY element into the SCALAR Option<Json>
  // field `selectedOption`, so the same scalar null-collapse the reducer applies everywhere
  // else (jNoNull on content/auth/usage/error) applies at THIS array→scalar crossing: a matched
  // element that is JNull collapses to None. On real input `options` is a ConfirmationOption[]
  // of non-null objects, so this is behaviourally identical to the TS reducer's
  // `options.find(o => o.id === id)`; it just makes the crossing TOTAL (no options-non-null
  // precondition needed) instead of documenting the AHP-TS invariant as a carve-out.
  function selectOption(options: Option<seq<Wire.Json>>, id: Option<string>): Option<Wire.Json>
    decreases if options.Some? then |options.value| else 0
  {
    if options.None? || id.None? || |options.value| == 0 then None
    else if optionId(options.value[0]) == id.value then jNoNull(options.value[0])
    else selectOption(Some(options.value[1..]), id)
  }

  function deltaOne(tc: ToolCall, content: string, inv: Option<string>, meta: Option<Wire.Json>): ToolCall {
    if tc.status != "streaming" then tc
    else tc.(meta := metaOr(meta, tc.meta), partialInput := Some(optOr(tc.partialInput, "") + content), invocationMessage := optOr(inv, tc.invocationMessage))
  }
  // ready: streaming|running → (running if confirmed, else pending-confirmation).
  // Rebuilds from base: partialInput/success/pastTense/reason/content dropped.
  function readyOne(tc: ToolCall, inv: Option<string>, input: Option<string>, title: Option<Wire.Json>, risk: Option<Wire.Json>, edits: Option<Wire.Json>, editable: Option<bool>, options: Option<seq<Wire.Json>>, confirmed: Option<string>, meta: Option<Wire.Json>): ToolCall {
    if tc.status != "streaming" && tc.status != "running" && tc.status != "pending-confirmation" then tc
    else if confirmed.Some? then
      tc.(status := "running", meta := metaOr(meta, tc.meta), invocationMessage := optOr(inv, ""), toolInput := input,
          confirmationTitle := None, riskAssessment := None, edits := None, editable := None, options := None,
          confirmed := confirmed, success := None, pastTenseMessage := None, reason := None, reasonMessage := None,
          userSuggestion := None, selectedOption := None, content := None, structuredContent := None, error := None,
          auth := None, partialInput := None)
    else
      var pending := tc.status == "pending-confirmation";
      tc.(status := "pending-confirmation", meta := metaOr(meta, tc.meta), invocationMessage := optOr(inv, ""),
          toolInput := if input.Some? || !pending then input else tc.toolInput,
          confirmationTitle := if title.Some? || !pending then title else tc.confirmationTitle,
          riskAssessment := if risk.Some? || !pending then risk else tc.riskAssessment,
          edits := if edits.Some? || !pending then edits else tc.edits,
          editable := if editable.Some? || !pending then editable else tc.editable,
          options := if options.Some? || !pending then options else tc.options,
          confirmed := None, success := None, pastTenseMessage := None, reason := None, reasonMessage := None,
          userSuggestion := None, selectedOption := None, content := None, structuredContent := None, error := None,
          auth := None, partialInput := None)
  }
  // Authentication can interrupt only an actively running tool call. Its
  // existing typed metadata/result content stays intact while the status moves
  // to the input-needed class.
  function authRequiredOne(tc: ToolCall, challenge: Wire.Json, meta: Option<Wire.Json>): ToolCall {
    if tc.status != "running" || tc.contributor.None? || tc.contributor.value.kind != "mcp" then tc
    else tc.(status := "auth-required", meta := metaOr(meta, tc.meta), confirmationTitle := None,
             riskAssessment := None, edits := None, editable := None, options := None,
             success := None, pastTenseMessage := None, reason := None, reasonMessage := None,
             userSuggestion := None, structuredContent := None, error := None, auth := jNoNull(challenge), partialInput := None)
  }
  function authResolvedOne(tc: ToolCall, meta: Option<Wire.Json>): ToolCall {
    if tc.status != "auth-required" then tc
    else tc.(status := "running", meta := metaOr(meta, tc.meta), auth := None)
  }
  // confirmed: only from pending-confirmation. approved → running (editedToolInput
  // overrides toolInput); denied → cancelled with the given reason.
  function confirmOne(tc: ToolCall, approved: bool, confirmedVal: Option<string>, reason: Option<string>, reasonMessage: Option<Wire.Json>, userSuggestion: Option<Wire.Json>, editedInput: Option<string>, selectedOptionId: Option<string>, meta: Option<Wire.Json>): ToolCall {
    if tc.status != "pending-confirmation" then tc
    else if approved then
      tc.(status := "running", meta := metaOr(meta, tc.meta), toolInput := if editedInput.Some? then editedInput else tc.toolInput,
          confirmationTitle := None, riskAssessment := None, edits := None, editable := None, options := None,
          confirmed := confirmedVal, selectedOption := selectOption(tc.options, selectedOptionId),
          partialInput := None, success := None, pastTenseMessage := None, reason := None, reasonMessage := None,
          userSuggestion := None, content := None, structuredContent := None, error := None, auth := None)
    else
      tc.(status := "cancelled", meta := metaOr(meta, tc.meta), confirmationTitle := None, riskAssessment := None,
          edits := None, editable := None, options := None, confirmed := None, success := None, pastTenseMessage := None,
          reason := reason, reasonMessage := reasonMessage, userSuggestion := userSuggestion,
          selectedOption := selectOption(tc.options, selectedOptionId), content := None, structuredContent := None,
          error := None, auth := None, partialInput := None)
  }
  // complete: from running or pending-confirmation. pending-confirmation defaults
  // confirmed to "not-needed". requiresResultConfirmation → pending-result-confirmation.
  function completeOne(tc: ToolCall, ok: bool, past: Option<string>, resultContent: Option<Wire.Json>, structured: Option<Wire.Json>, err: Option<Wire.Json>, requiresResultConfirmation: bool, meta: Option<Wire.Json>): ToolCall {
    if tc.status == "auth-required" then
      if ok then tc
      else tc.(status := "completed", meta := metaOr(meta, tc.meta), confirmationTitle := None, riskAssessment := None,
               edits := None, editable := None, options := None, success := Some(false), pastTenseMessage := past,
               reason := None, reasonMessage := None, userSuggestion := None,
               content := if resultContent.Some? then resultContent else tc.content,
               structuredContent := structured, error := err, auth := None, partialInput := None)
    else if tc.status != "running" && tc.status != "pending-confirmation" then tc
    else
      var conf := if tc.status == "running" then tc.confirmed else Some("not-needed");
      var selected := if tc.status == "running" then tc.selectedOption else None;
      var newStatus := if requiresResultConfirmation then "pending-result-confirmation" else "completed";
      tc.(status := newStatus, meta := metaOr(meta, tc.meta), confirmationTitle := None, riskAssessment := None,
          edits := None, editable := None, options := None, confirmed := conf, selectedOption := selected,
          success := Some(ok), pastTenseMessage := past, reason := None, reasonMessage := None, userSuggestion := None,
          content := resultContent, structuredContent := structured, error := err, auth := None, partialInput := None)
  }
  // resultConfirmed: only from pending-result-confirmation. approved → completed
  // (keep result); denied → cancelled with reason "result-denied" (drop result).
  function resultConfirmOne(tc: ToolCall, approved: bool, meta: Option<Wire.Json>): ToolCall {
    if tc.status != "pending-result-confirmation" then tc
    else if approved then tc.(status := "completed", meta := metaOr(meta, tc.meta))
    else tc.(status := "cancelled", meta := metaOr(meta, tc.meta), reason := Some("result-denied"),
             confirmed := None, success := None, pastTenseMessage := None, content := None,
             structuredContent := None, error := None)
  }
  // contentChanged: only mutates a RUNNING tool call.
  function contentChangedOne(tc: ToolCall, c: Wire.Json, meta: Option<Wire.Json>): ToolCall {
    if tc.status != "running" then tc else tc.(meta := metaOr(meta, tc.meta), content := jNoNull(c))
  }

  // turnComplete/turnCancelled/error: force any still-in-progress tool call to
  // cancelled (reason "skipped"), rebuilt from base — streaming loses toolInput;
  // confirmed/success/pastTense/content/partialInput all dropped.
  function forceCancel(ps: seq<Part>): seq<Part>
  {
    Ops.Map((h: Part) => match h
        case PToolCall(tc) =>
          if tc.status == "completed" || tc.status == "cancelled" then h
          else PToolCall(tc.(status := "cancelled", reason := Some("skipped"),
                             toolInput := if tc.status == "streaming" then None else tc.toolInput,
                             confirmationTitle := None, riskAssessment := None, edits := None, editable := None, options := None,
                             confirmed := None, success := None, pastTenseMessage := None, reasonMessage := None,
                             userSuggestion := None, selectedOption := None, content := None, structuredContent := None,
                             error := None, auth := None, partialInput := None))
        case _ => h, ps)
  }

  // endTurn: finalize the active turn into the history, force-cancel in-progress
  // tool calls, and recompute status. #338: any unresolved input-request part stays
  // verbatim in the finalized turn transcript (forceCancel only rewrites tool-call parts),
  // matching the TS reducer — the completed turn durably records the open request.
  function endTurn(s: ChatState, id: string, turnState: string, isError: bool, err: Option<Wire.Json>): (ChatState, ReduceOutcome)
  {
    if !turnMatches(s, id) then (s, NoOp)
    else
      var t := s.activeTurn.value;
      var finalized := t.(state := turnState, parts := forceCancel(t.parts), error := err);
      var next := s.(turns := s.turns + [finalized], activeTurn := None);
      (next.(status := summaryStatus(next, isError)), Applied)
  }

  function ApplyToChat(s: ChatState, a: ChatAction, now: int): (ChatState, ReduceOutcome)
  {
    match a
    case CDraftChanged(d)     => (s.(draft := d), Applied)
    case CActivityChanged(ac) => (s.(activity := ac), Applied)
    case CTurnStarted(id, m, qid) =>
      var withTurn := s.(activeTurn := Some(Turn(id, m, [], "in-progress", None, None)));
      var s2 := withTurn.(status := clearBit(summaryStatus(withTurn, false), READ));
      var s3 := if qid.Some? && s2.steeringMessage.Some? && s2.steeringMessage.value.id == qid.value then s2.(steeringMessage := None) else s2;
      var s4 := if qid.Some? then s3.(queuedMessages := removeQ(s3.queuedMessages, qid.value)) else s3;
      (s4, Applied)
    case CTurnComplete(id)    => endTurn(s, id, "complete", false, None)
    case CTurnCancelled(id)   => endTurn(s, id, "cancelled", false, None)
    case CError(id, e)        => endTurn(s, id, "error", true, jNoNull(e))
    case CUsage(id, u)        => if turnMatches(s, id) then (s.(activeTurn := Some(s.activeTurn.value.(usage := jNoNull(u)))), Applied) else (s, NoOp)
    case CResponsePart(id, p) => if turnMatches(s, id) then (s.(activeTurn := Some(appendPart(s.activeTurn.value, p))), Applied) else (s, NoOp)
    case CDelta(id, pid, c)   => if turnMatches(s, id) then (s.(activeTurn := Some(s.activeTurn.value.(parts := deltaMarkdown(s.activeTurn.value.parts, pid, c)))), Applied) else (s, NoOp)
    case CReasoning(id, pid, c) => if turnMatches(s, id) then (s.(activeTurn := Some(s.activeTurn.value.(parts := deltaReasoning(s.activeTurn.value.parts, pid, c)))), Applied) else (s, NoOp)
    case CToolCallStart(id, tcId, tn, dn, intent, contributor, meta) =>
      if turnMatches(s, id) then
        var tc := ToolCall(tcId, tn, dn, "streaming", intent, contributor, meta, "", None, None, None, None, None, None,
                           None, None, None, None, None, None, None, None, None, None, None, None);
        (s.(activeTurn := Some(s.activeTurn.value.(parts := s.activeTurn.value.parts + [PToolCall(tc)]))), Applied)
      else (s, NoOp)
    case CToolCallDelta(id, tcId, c, inv, meta) =>
      if turnMatches(s, id) then (s.(activeTurn := Some(s.activeTurn.value.(parts := mapTC(s.activeTurn.value.parts, tcId, (x: ToolCall) => deltaOne(x, c, inv, meta))))), Applied) else (s, NoOp)
    case CToolCallReady(id, tcId, inv, input, title, risk, edits, editable, options, confirmed, meta) =>
      if turnMatches(s, id) then
        var s1 := s.(activeTurn := Some(s.activeTurn.value.(parts := mapTC(s.activeTurn.value.parts, tcId, (x: ToolCall) => readyOne(x, inv, input, title, risk, edits, editable, options, confirmed, meta)))));
        (s1.(status := summaryStatus(s1, false)), Applied)
      else (s, NoOp)
    case CToolCallConfirmed(id, tcId, approved, cv, rsn, reasonMessage, userSuggestion, edited, selectedOptionId, meta) =>
      if turnMatches(s, id) then
        var s1 := s.(activeTurn := Some(s.activeTurn.value.(parts := mapTC(s.activeTurn.value.parts, tcId, (x: ToolCall) => confirmOne(x, approved, cv, rsn, reasonMessage, userSuggestion, edited, selectedOptionId, meta)))));
        (s1.(status := summaryStatus(s1, false)), Applied)
      else (s, NoOp)
    case CToolCallAuthRequired(id, tcId, auth, meta) =>
      if turnMatches(s, id) then
        var s1 := s.(activeTurn := Some(s.activeTurn.value.(parts := mapTC(s.activeTurn.value.parts, tcId, (x: ToolCall) => authRequiredOne(x, auth, meta)))));
        (s1.(status := summaryStatus(s1, false)), Applied)
      else (s, NoOp)
    case CToolCallAuthResolved(id, tcId, meta) =>
      if turnMatches(s, id) then
        var s1 := s.(activeTurn := Some(s.activeTurn.value.(parts := mapTC(s.activeTurn.value.parts, tcId, (x: ToolCall) => authResolvedOne(x, meta)))));
        (s1.(status := summaryStatus(s1, false)), Applied)
      else (s, NoOp)
    case CToolCallComplete(id, tcId, ok, past, resultContent, structured, err, rrc, meta) =>
      if turnMatches(s, id) then
        var s1 := s.(activeTurn := Some(s.activeTurn.value.(parts := mapTC(s.activeTurn.value.parts, tcId, (x: ToolCall) => completeOne(x, ok, past, resultContent, structured, err, rrc, meta)))));
        (s1.(status := summaryStatus(s1, false)), Applied)
      else (s, NoOp)
    case CToolCallResultConfirmed(id, tcId, approved, meta) =>
      if turnMatches(s, id) then
        var s1 := s.(activeTurn := Some(s.activeTurn.value.(parts := mapTC(s.activeTurn.value.parts, tcId, (x: ToolCall) => resultConfirmOne(x, approved, meta)))));
        (s1.(status := summaryStatus(s1, false)), Applied)
      else (s, NoOp)
    case CToolCallContentChanged(id, tcId, c, meta) =>
      if turnMatches(s, id) then (s.(activeTurn := Some(s.activeTurn.value.(parts := mapTC(s.activeTurn.value.parts, tcId, (x: ToolCall) => contentChangedOne(x, c, meta))))), Applied) else (s, NoOp)
    case CTruncated(idOpt) =>
      if idOpt.Some? && !hasTurn(s.turns, idOpt.value) then (s, NoOp)
      else
        var kept := if idOpt.Some? then keepUpTo(s.turns, idOpt.value) else [];
        var next := s.(turns := kept, activeTurn := None, nextCursor := if idOpt.Some? then s.nextCursor else None);
        (next.(status := summaryStatus(next, false)), Applied)
    case CQueuedReordered(order) => (s.(queuedMessages := reorderQ(s.queuedMessages, order)), Applied)
    case CTurnsLoaded(loaded, cursor) => (s.(turns := dedupPrepend(loaded, s.turns), nextCursor := cursor), Applied)
    // #338: inputRequested upserts a LIVE input-request part into the active turn (no-op with no active turn,
    // mirroring TS `upsertInputRequestPart`). Recompute status (an open request → InputNeeded) and clear IsRead.
    case CInputRequested(req) =>
      if s.activeTurn.Some? then
        var at := s.activeTurn.value;
        var next := s.(activeTurn := Some(at.(parts := upsertInputPart(at.parts, req))));
        (next.(status := clearBit(summaryStatus(next, false), READ)), Applied)
      else (s, NoOp)
    // answerChanged patches the draft answers on the OPEN matching part in place; no-op if no such part
    // (or no active turn). Status is NOT recomputed (matches TS — an answer draft can't open/close a request).
    case CInputAnswerChanged(reqId, qId, ans) =>
      if s.activeTurn.Some? && hasOpenReq(s.activeTurn.value.parts, reqId) then
        var at := s.activeTurn.value;
        (s.(activeTurn := Some(at.(parts := changeAnswerPart(at.parts, reqId, qId, ans)))), Applied)
      else (s, NoOp)
    // completed resolves the OPEN matching part IN PLACE (stable stream position): merge final answers over
    // drafts and record the response kind. No-op for an unknown id or no active turn. Recompute status
    // (resolving the last open request drops InputNeeded back to InProgress/Idle).
    case CInputCompleted(reqId, resp, answers) =>
      if s.activeTurn.Some? && hasOpenReq(s.activeTurn.value.parts, reqId) then
        var at := s.activeTurn.value;
        var next := s.(activeTurn := Some(at.(parts := completeAnswerPart(at.parts, reqId, resp, answers))));
        (next.(status := summaryStatus(next, false)), Applied)
      else (s, NoOp)
    case CPendingMessageSet(kind, id, msg) =>
      if kind == "steering" then (s.(steeringMessage := Some(QMsg(id, msg))), Applied)
      else (s.(queuedMessages := upsertQ(s.queuedMessages, QMsg(id, msg))), Applied)
    case CPendingMessageRemoved(kind, id) =>
      if kind == "steering" then (if s.steeringMessage.Some? && s.steeringMessage.value.id == id then (s.(steeringMessage := None), Applied) else (s, NoOp))
      else (s.(queuedMessages := removeQ(s.queuedMessages, id)), Applied)
    case CUnknown(_)          => (s, NoOp)
  }

  lemma C_UnknownIsNoOp(s: ChatState, a: ChatAction, now: int)
    requires a.CUnknown?
    ensures ApplyToChat(s, a, now).0 == s && ApplyToChat(s, a, now).1 == NoOp {}
  lemma C_UnknownIsNoOp_NonVacuous() ensures exists a: ChatAction :: a.CUnknown? { assert (CUnknown(Wire.JNull)).CUnknown?; }
  // wrong-turn-id no-op guard: any turn-scoped action on a non-matching active turn is a no-op.
  lemma C_WrongTurnIsNoOp(s: ChatState, id: string, u: Wire.Json, now: int)
    requires !turnMatches(s, id)
    ensures ApplyToChat(s, CUsage(id, u), now).0 == s
    ensures ApplyToChat(s, CTurnComplete(id), now).0 == s
    ensures ApplyToChat(s, CToolCallStart(id, "x", "y", "z", None, None, None), now).0 == s {}

  // ══════════════════════════════════════════════════════════════════════════
  //  rung 4 / 6 / 7 — keyed-collection uniqueness invariant + reducer preservation
  //  (docs/shapes/07). ChatState carries ONE id-keyed REDUCER collection —
  //  queuedMessages; well-formedness is that it holds DISTINCT ids. `turns` (key
  //  turnId) and each turn's `parts` (key partKey — now including the #338
  //  input-request parts) are APPEND / find-first-where based with NO freshness
  //  guard, so they are NOT covered here — the input-request part operations
  //  (upsert-first-open / update-first-where / complete-first-open) are well-defined
  //  under find-first semantics without a uniqueness premise. See the design note by
  //  WfChatStatePreserved.
  // ══════════════════════════════════════════════════════════════════════════
  ghost predicate WfChatState(s: ChatState) {
    SR.UniqueKeys(qKey, s.queuedMessages)
  }

  // rung-7 premise carrier. Chat installs NO caller-supplied keyed seq verbatim — the queued-message arm is an
  // upsert / remove / reorder discharged by a kernel law. So (unlike changeset's OperationsChanged /
  // ContentChanged) no arm needs an incoming-uniqueness premise; kept explicit to mirror the docs/shapes/07 shape.
  ghost predicate WfChatAction(a: ChatAction) { true }
  lemma WfChatAction_NonVacuous() ensures exists a: ChatAction :: WfChatAction(a) { assert WfChatAction(CUnknown(Wire.JNull)); }

  // rung 6 — the reducer PRESERVES the queuedMessages keyed-uniqueness invariant on EVERY arm. The three
  // queued-message arms cite their matching ConfluxSeqRoute / OrderedLog law; EVERY OTHER arm — including the
  // #338 input-request arms, which touch only the active turn's `parts` — leaves queuedMessages untouched, so
  // the invariant carries straight through from the precondition.
  lemma WfChatStatePreserved(s: ChatState, a: ChatAction, now: int)
    requires WfChatState(s) && WfChatAction(a)
    ensures WfChatState(ApplyToChat(s, a, now).0)
  {
    match a
    case CTurnStarted(id, m, qid) =>
      if qid.Some? { SR.RemovePreservesUnique(qKey, s.queuedMessages, qid.value); }  // queuedMessages := removeQ(...)
    case CQueuedReordered(order) =>
      OL.ReorderPreservesUnique(qKey, order, s.queuedMessages);                      // queuedMessages := reorderQ(...)
    case CPendingMessageSet(kind, id, msg) =>
      if kind != "steering" { SR.UpsertPreservesUnique(qKey, s.queuedMessages, QMsg(id, msg)); }  // queued upsert
    case CPendingMessageRemoved(kind, id) =>
      if kind != "steering" { SR.RemovePreservesUnique(qKey, s.queuedMessages, id); }             // queued remove
    // Every remaining arm touches only scalar / activeTurn / turns fields — queuedMessages is carried through
    // unchanged, so WfChatState is preserved directly from the precondition. (The #338 CInputRequested /
    // CInputAnswerChanged / CInputCompleted arms mutate activeTurn.parts only; queuedMessages is untouched.)
    case CTurnComplete(id)  =>
    case CTurnCancelled(id) =>
    case CError(id, e)      =>
    case CTruncated(idOpt)  =>
    case CInputRequested(req) =>
    case CInputAnswerChanged(reqId, qId, ans) =>
    case CInputCompleted(reqId, resp, answers) =>
    case CDraftChanged(_) =>
    case CActivityChanged(_) =>
    case CUsage(_, _) =>
    case CResponsePart(_, _) =>
    case CDelta(_, _, _) =>
    case CReasoning(_, _, _) =>
    case CToolCallStart(_, _, _, _, _, _, _) =>
    case CToolCallDelta(_, _, _, _, _) =>
    case CToolCallReady(_, _, _, _, _, _, _, _, _, _, _) =>
    case CToolCallConfirmed(_, _, _, _, _, _, _, _, _, _) =>
    case CToolCallAuthRequired(_, _, _, _) =>
    case CToolCallAuthResolved(_, _, _) =>
    case CToolCallComplete(_, _, _, _, _, _, _, _, _) =>
    case CToolCallResultConfirmed(_, _, _, _) =>
    case CToolCallContentChanged(_, _, _, _) =>
    case CTurnsLoaded(_, _) =>
    case CUnknown(_) =>
  }

  function apply1(s: ChatState, a: ChatAction): ChatState { ApplyToChat(s, a, 9999).0 }
  // apply1-level restatement of the rung-6 preservation theorem, so the aggregate (ahp.dfy) — which dispatches
  // through apply1, not ApplyToChat — can compose it without unfolding apply1's (exported-opaque) body.
  lemma Apply1PreservesWfChat(s: ChatState, a: ChatAction)
    requires WfChatState(s) && WfChatAction(a)
    ensures  WfChatState(apply1(s, a))
  { WfChatStatePreserved(s, a, 9999); }
  // Concrete well-formed witness handed to the aggregate (ahp.dfy) so its WfAhpStateInv non-vacuity witness
  // need not unfold this module's exported-opaque WfChatState (empty keyed seqs are trivially id-unique).
  lemma ChatWfWitness() returns (s: ChatState)
    ensures WfChatState(s)
  { s := C0(); assert SR.UniqueKeys(qKey, s.queuedMessages); }

  // ══════════════════════════════════════════════════════════════════════════
  //  no-stored-null tool-call facet — the reducer-side twin of ChatCodec.WFToolCall's
  //  okJson conjuncts (mirrors session's WfCust). Every scalar Option<Json> tool-call field
  //  that maps to a TS `T | undefined` never stores Some(JNull). `options` (a Json ARRAY, not a
  //  scalar) needs NO element carve-out here: the ONLY way an options element reaches a scalar
  //  field is via `selectedOption := selectOption(options, id)`, and selectOption itself
  //  null-collapses a matched element (jNoNull), so `selectedOption` stays okJson for ANY
  //  options array — no options-non-null precondition required (see SelectOptionNoNull).
  // ══════════════════════════════════════════════════════════════════════════
  ghost predicate WfToolCallNoNull(tc: ToolCall) {
    (tc.meta.Some?              ==> tc.meta.value              != Wire.JNull) &&
    (tc.confirmationTitle.Some? ==> tc.confirmationTitle.value != Wire.JNull) &&
    (tc.riskAssessment.Some?    ==> tc.riskAssessment.value    != Wire.JNull) &&
    (tc.edits.Some?             ==> tc.edits.value             != Wire.JNull) &&
    (tc.reasonMessage.Some?     ==> tc.reasonMessage.value     != Wire.JNull) &&
    (tc.userSuggestion.Some?    ==> tc.userSuggestion.value    != Wire.JNull) &&
    (tc.selectedOption.Some?    ==> tc.selectedOption.value    != Wire.JNull) &&
    (tc.content.Some?           ==> tc.content.value           != Wire.JNull) &&
    (tc.structuredContent.Some? ==> tc.structuredContent.value != Wire.JNull) &&
    (tc.error.Some?             ==> tc.error.value             != Wire.JNull) &&
    (tc.auth.Some?              ==> tc.auth.value              != Wire.JNull)
  }

  // selectOption returns None or a jNoNull-collapsed ELEMENT of `options`; so the selected value
  // is never a stored JNull — UNCONDITIONALLY, for any options array (the matched-element arm
  // collapses a JNull to None). Structural induction mirrors selectOption's own recursion.
  lemma SelectOptionNoNull(options: Option<seq<Wire.Json>>, id: Option<string>)
    ensures var r := selectOption(options, id); r.Some? ==> r.value != Wire.JNull
    decreases if options.Some? then |options.value| else 0
  {
    if options.None? || id.None? || |options.value| == 0 {
    } else if optionId(options.value[0]) == id.value {
    } else {
      SelectOptionNoNull(Some(options.value[1..]), id);
    }
  }

  // rung-6 per-arm okJson-preservation for the CToolCallConfirmed arm (the one Wave 2 left
  // blocked): the confirm transform preserves the no-stored-null tool-call facet. The three
  // action-payload premises are EXACTLY ChatCodec.WFChatAction's CToolCallConfirmed clause
  // (okJson meta / reasonMessage / userSuggestion); every other written field is None or
  // resolves to a non-null via SelectOptionNoNull (which now holds for ANY options array,
  // since selectOption null-collapses the selected element), so no arm can introduce a stored JNull.
  lemma ConfirmOnePreservesNoNull(
      tc: ToolCall, approved: bool, confirmedVal: Option<string>, reason: Option<string>,
      reasonMessage: Option<Wire.Json>, userSuggestion: Option<Wire.Json>, editedInput: Option<string>,
      selectedOptionId: Option<string>, meta: Option<Wire.Json>)
    requires WfToolCallNoNull(tc)
    requires meta.Some?           ==> meta.value           != Wire.JNull
    requires reasonMessage.Some?  ==> reasonMessage.value  != Wire.JNull
    requires userSuggestion.Some? ==> userSuggestion.value != Wire.JNull
    ensures WfToolCallNoNull(confirmOne(tc, approved, confirmedVal, reason, reasonMessage, userSuggestion, editedInput, selectedOptionId, meta))
  {
    SelectOptionNoNull(tc.options, selectedOptionId);
  }

  // Non-vacuity: a GENUINE pending-confirmation tool call carrying a real populated `options`
  // array satisfies the facet AND its confirmed image still does (selectedOption resolves to the
  // selected non-null option) — so ConfirmOnePreservesNoNull is not vacuously true / empty-options-only.
  lemma ConfirmOnePreservesNoNull_NonVacuous()
    ensures exists tc: ToolCall ::
      (WfToolCallNoNull(tc) && tc.options.Some? && |tc.options.value| > 0
       && WfToolCallNoNull(confirmOne(tc, true, Some("user-action"), None, None, None, None, Some("approve"), None)))
  {
    var opt := Wire.JObj(map["id" := Wire.JStr("approve")]);
    var tc := TC0("tc-1", None).(status := "pending-confirmation", options := Some([opt]));
    assert WfToolCallNoNull(tc);
    ConfirmOnePreservesNoNull(tc, true, Some("user-action"), None, None, None, None, Some("approve"), None);
  }

  function fold(s: ChatState, acts: seq<ChatAction>): ChatState
  { CC.Fold(apply1, s, acts) }   // hand-rolled left-fold of apply1 DELETED — now the kernel fold

  function C0(): ChatState { ChatState([], "Chat", 1, "t1", None, None, None, None, [], None) }
  function T0(id: string): Turn { Turn(id, Wire.JNull, [], "in-progress", None, None) }
  function TDone(id: string): Turn { Turn(id, Wire.JNull, [], "complete", None, None) }
  // a fresh tool call as toolCallStart builds it.
  function TC0(id: string, intent: Option<string>): ToolCall {
    ToolCall(id, "bash", "Run Command", "streaming", intent, None, None, "", None, None, None, None, None, None,
             None, None, None, None, None, None, None, None, None, None, None, None)
  }

  // Total harness projections keep executable corpus checks free of repeated
  // Option.value / seq-index well-definedness obligations.  A missing active
  // turn, stored turn, part, or non-tool first part simply yields None.
  function activeParts(s: ChatState): seq<Part>
  { if s.activeTurn.Some? then s.activeTurn.value.parts else [] }
  function firstActiveToolCall(s: ChatState): Option<ToolCall> {
    var ps := activeParts(s);
    if |ps| > 0 && ps[0].PToolCall? then Some(ps[0].tc) else None
  }
  function firstStoredToolCall(s: ChatState): Option<ToolCall> {
    if |s.turns| == 0 then None else
      var ps := s.turns[0].parts;
      if |ps| > 0 && ps[0].PToolCall? then Some(ps[0].tc) else None
  }

  // The corpus is split into focused sub-methods so each stays small enough for
  // the verifier (the load-bearing PROOFS are the lemmas above; these methods are
  // the executable harness the extracted C#/JS run against the real fixtures).
  method runScalarTurns() returns (pass: nat) {
    pass := 0;
    var s := C0();
    // scalar / draft / activity (4)
    if apply1(s.(draft := Some("x")), CDraftChanged(None)).draft == None { pass := pass+1; }
    if apply1(s, CDraftChanged(Some("hi"))).draft == Some("hi") { pass := pass+1; }
    if apply1(s, CActivityChanged(Some("Running terminal command"))).activity == Some("Running terminal command") { pass := pass+1; }
    if apply1(s, CActivityChanged(None)).activity == None { pass := pass+1; }
    // turn lifecycle guards (4)
    if apply1(s, CTurnStarted("turn-1", Wire.JStr("Hello"), None)).activeTurn == Some(Turn("turn-1", Wire.JStr("Hello"), [], "in-progress", None, None)) { pass := pass+1; }
    var act1 := s.(activeTurn := Some(T0("t1")));
    if apply1(act1, CTurnComplete("wrong-turn")) == act1 { pass := pass+1; }                 // 018 wrong id → no-op
    if apply1(act1, CUsage("wrong-turn", Wire.JNull)) == act1 { pass := pass+1; }                 // 093 wrong id → no-op
    if apply1(act1, CToolCallStart("wrong-turn", "tc-1", "bash", "Run Command", None, None, None)) == act1 { pass := pass+1; }  // 092 wrong id → no-op
    // turnComplete/turnCancelled matching finalizes (2)
    if apply1(s.(activeTurn := Some(Turn("t1", Wire.JStr("T"), [], "in-progress", None, None))), CTurnComplete("t1"))
       == s.(activeTurn := None, turns := [Turn("t1", Wire.JStr("T"), [], "complete", None, None)]) { pass := pass+1; }
    if apply1(s.(activeTurn := Some(T0("t1"))), CTurnCancelled("t1")).turns == [Turn("t1", Wire.JNull, [], "cancelled", None, None)] { pass := pass+1; }
    // usage on matching turn / error finalizes turn (2)
    if apply1(act1, CUsage("t1", Wire.JStr("u"))).activeTurn == Some(T0("t1").(usage := Some(Wire.JStr("u")))) { pass := pass+1; }
    if apply1(act1, CError("t1", Wire.JStr("boom"))).turns == [Turn("t1", Wire.JNull, [], "error", None, Some(Wire.JStr("boom")))] { pass := pass+1; }  // 016 error finalizes
    // responsePart + delta accumulation (3)
    var withMd := apply1(act1, CResponsePart("t1", PMarkdown("md-1", "")));
    if withMd.activeTurn == Some(T0("t1").(parts := [PMarkdown("md-1", "")])) { pass := pass+1; }
    if apply1(withMd, CDelta("t1", "md-1", "Hello from chat")).activeTurn == Some(T0("t1").(parts := [PMarkdown("md-1", "Hello from chat")])) { pass := pass+1; }
    // delta targeting a reasoning part → no-op (delta only touches markdown, 102-shape)
    if apply1(act1.(activeTurn := Some(T0("t1").(parts := [PReasoning("r-1", "th")]))), CDelta("t1", "r-1", "inking")).activeTurn
       == Some(T0("t1").(parts := [PReasoning("r-1", "th")])) { pass := pass+1; }
  }

  method {:isolate_assertions} runToolCalls() returns (pass: nat) {
    pass := 0;
    var s := C0();
    var act1 := s.(activeTurn := Some(T0("t1")));
    // start appends a streaming tool call
    var started := apply1(act1, CToolCallStart("t1", "tc-1", "bash", "Run Command", Some("List the files in the current directory"), None, None));
    if started.activeTurn == Some(T0("t1").(parts := [PToolCall(TC0("tc-1", Some("List the files in the current directory")))])) { pass := pass+1; }
    // delta accumulates partialInput (NOT toolInput) + invocationMessage, status preserved
    var afterDelta := apply1(started, CToolCallDelta("t1", "tc-1", "ls -la", Some("Listing files"), None));
    if activeParts(afterDelta) == [PToolCall(TC0("tc-1", Some("List the files in the current directory")).(partialInput := Some("ls -la"), invocationMessage := "Listing files"))] { pass := pass+1; }
    // ready (needs-confirmation) → pending-confirmation, sets inv/input, drops partialInput
    var afterReady := apply1(afterDelta, CToolCallReady("t1", "tc-1", Some("Run: ls -la"), Some("ls -la"), None, None, None, None, None, None, None));
    var tcReady := firstActiveToolCall(afterReady);
    if tcReady.Some? && tcReady.value.status == "pending-confirmation" && tcReady.value.invocationMessage == "Run: ls -la" && tcReady.value.toolInput == Some("ls -la") && tcReady.value.partialInput == None { pass := pass+1; }
    // confirmed (approved) → running, sets confirmed
    var afterConf := apply1(afterReady, CToolCallConfirmed("t1", "tc-1", true, Some("user-action"), None, None, None, None, None, None));
    var tcConf := firstActiveToolCall(afterConf);
    if tcConf.Some? && tcConf.value.status == "running" && tcConf.value.confirmed == Some("user-action") { pass := pass+1; }
    // complete → completed, sets success + pastTense
    var afterComp := apply1(afterConf, CToolCallComplete("t1", "tc-1", true, Some("Ran command"), None, None, None, false, None));
    var tcComp := firstActiveToolCall(afterComp);
    if tcComp.Some? && tcComp.value.status == "completed" && tcComp.value.success == Some(true) && tcComp.value.pastTenseMessage == Some("Ran command") { pass := pass+1; }
    // 019 FULL lifecycle via fold: start→delta→ready→confirmed→complete
    var life := fold(act1, [
      CToolCallStart("t1", "tc-1", "bash", "Run Command", Some("List the files in the current directory"), None, None),
      CToolCallDelta("t1", "tc-1", "ls -la", Some("Listing files"), None),
      CToolCallReady("t1", "tc-1", Some("Run: ls -la"), Some("ls -la"), None, None, None, None, None, None, None),
      CToolCallConfirmed("t1", "tc-1", true, Some("user-action"), None, None, None, None, None, None),
      CToolCallComplete("t1", "tc-1", true, Some("Ran command"), None, None, None, false, None)]);
    var lifeTc := firstActiveToolCall(life);
    if lifeTc.Some? && lifeTc.value.status == "completed" && lifeTc.value.confirmed == Some("user-action") && lifeTc.value.success == Some(true) && lifeTc.value.toolInput == Some("ls -la") { pass := pass+1; }
    // 020 ready with auto-confirm ("not-needed") → running directly
    var auto := fold(act1, [
      CToolCallStart("t1", "tc-1", "bash", "Run Command", None, None, None),
      CToolCallReady("t1", "tc-1", Some("Run"), None, None, None, None, None, None, Some("not-needed"), None)]);
    var autoTc := firstActiveToolCall(auto);
    if autoTc.Some? && autoTc.value.status == "running" && autoTc.value.confirmed == Some("not-needed") && autoTc.value.invocationMessage == "Run" { pass := pass+1; }
    // 029/242 ready refreshes an already-pending confirmation without dropping
    // its confirmation context; status stays InputNeeded(24).
    var pendingRisk := Wire.JObj(map["status" := Wire.JStr("loading")]);
    var pendingEdits := Wire.JObj(map["items" := Wire.JArr([])]);
    var pendingOption := Wire.JObj(map["id" := Wire.JStr("approve")]);
    var pendState := act1.(activeTurn := Some(T0("t1").(parts := [PToolCall(TC0("tc-1", None).(status := "pending-confirmation", toolInput := Some("rm -rf /tmp/test"), confirmationTitle := Some(Wire.JStr("Run in terminal")), riskAssessment := Some(pendingRisk), edits := Some(pendingEdits), editable := Some(true), options := Some([pendingOption])))])), status := 24);
    var refreshed := firstActiveToolCall(apply1(pendState, CToolCallReady("t1", "tc-1", Some("Run again"), None, None, None, None, None, None, None, None)));
    if refreshed.Some? && refreshed.value.status == "pending-confirmation" && refreshed.value.invocationMessage == "Run again" &&
       refreshed.value.toolInput == Some("rm -rf /tmp/test") && refreshed.value.confirmationTitle == Some(Wire.JStr("Run in terminal")) &&
       refreshed.value.riskAssessment == Some(pendingRisk) && refreshed.value.edits == Some(pendingEdits) &&
       refreshed.value.editable == Some(true) && refreshed.value.options == Some([pendingOption]) { pass := pass+1; }
    // 017 turnComplete force-cancels an in-progress (streaming) tool call → cancelled/skipped, turn finalized
    var forced := apply1(started, CTurnComplete("t1"));
    var fTc := firstStoredToolCall(forced);
    if forced.activeTurn == None && |forced.turns| > 0 && forced.turns[0].state == "complete" && fTc.Some? && fTc.value.status == "cancelled" && fTc.value.reason == Some("skipped") { pass := pass+1; }
    // toolCallComplete on a still-streaming call → guard holds, no transition
    var ignoredComplete := apply1(started, CToolCallComplete("t1", "tc-1", true, None, None, None, None, false, None));
    if activeParts(ignoredComplete) == activeParts(started) { pass := pass+1; }
    // 248 authentication interruption is typed and drives InputNeeded while
    // preserving the running tool call's input, confirmation and content.
    var authBase := act1.(activeTurn := Some(T0("t1").(parts := [PToolCall(TC0("tc-auth", None).(status := "running", contributor := Some(ToolCallContributor("mcp", Wire.JObj(map["kind" := Wire.JStr("mcp")]))), toolInput := Some("foo"), confirmed := Some("user-action"), content := Some(Wire.JStr("partial"))))])), status := 8);
    var challenge := Wire.JObj(map["reason" := Wire.JStr("required")]);
    var authState := apply1(authBase, CToolCallAuthRequired("t1", "tc-auth", challenge, Some(Wire.JStr("auth-meta"))));
    var authTc := firstActiveToolCall(authState);
    if authState.status == 24 && authTc.Some? && authTc.value.status == "auth-required" && authTc.value.auth == Some(challenge) && authTc.value.meta == Some(Wire.JStr("auth-meta")) && authTc.value.toolInput == Some("foo") && authTc.value.confirmed == Some("user-action") && authTc.value.content == Some(Wire.JStr("partial")) { pass := pass+1; }
    // 249 resolution clears only the challenge and returns to running while
    // preserving contributor, selection/input/confirmation and partial content.
    var resolvedState := apply1(authState, CToolCallAuthResolved("t1", "tc-auth", Some(Wire.JStr("resolved-meta"))));
    var resolvedTc := firstActiveToolCall(resolvedState);
    if resolvedState.status == 8 && resolvedTc.Some? && resolvedTc.value.status == "running" && resolvedTc.value.auth == None &&
       resolvedTc.value.contributor == Some(ToolCallContributor("mcp", Wire.JObj(map["kind" := Wire.JStr("mcp")]))) && resolvedTc.value.toolInput == Some("foo") &&
       resolvedTc.value.confirmed == Some("user-action") && resolvedTc.value.content == Some(Wire.JStr("partial")) &&
       resolvedTc.value.meta == Some(Wire.JStr("resolved-meta")) { pass := pass+1; }
    // 244/245 negative invariant: running client/absent-contributor calls cannot
    // enter auth-required.
    var clientBase := authBase.(activeTurn := Some(T0("t1").(parts := [PToolCall(TC0("tc-client", None).(status := "running", contributor := Some(ToolCallContributor("client", Wire.JObj(map["kind" := Wire.JStr("client")]))), confirmed := Some("not-needed")))])));
    if apply1(clientBase, CToolCallAuthRequired("t1", "tc-client", challenge, None)) == clientBase { pass := pass+1; }
    // 253/255 a failed completion cancels the auth wait into completed and
    // ignores requiresResultConfirmation, retaining prior output.
    var authFailed := apply1(authState, CToolCallComplete("t1", "tc-auth", false, Some("Cancelled search"), None, None, Some(Wire.JStr("cancelled")), true, Some(Wire.JStr("complete-meta"))));
    var failedTc := firstActiveToolCall(authFailed);
    if authFailed.status == 8 && failedTc.Some? && failedTc.value.status == "completed" && failedTc.value.success == Some(false) && failedTc.value.pastTenseMessage == Some("Cancelled search") && failedTc.value.content == Some(Wire.JStr("partial")) { pass := pass+1; }
    // 254 a successful result cannot bypass the outstanding auth requirement.
    if apply1(authState, CToolCallComplete("t1", "tc-auth", true, Some("Should not apply"), None, None, None, false, None)) == authState { pass := pass+1; }
  }

  method runPendingHistory() returns (pass: nat) {
    pass := 0;
    var s := C0();
    var act1 := s.(activeTurn := Some(T0("t1")));
    // pending messages: steering + queued (5)
    if apply1(s, CPendingMessageSet("steering", "sm-1", Wire.JStr("Focus on tests"))).steeringMessage == Some(QMsg("sm-1", Wire.JStr("Focus on tests"))) { pass := pass+1; }
    if apply1(s, CPendingMessageRemoved("steering", "sm-unknown")) == s { pass := pass+1; }   // 043 mismatched id → no-op
    if apply1(s.(steeringMessage := Some(QMsg("sm-1", Wire.JNull))), CPendingMessageRemoved("steering", "sm-1")).steeringMessage == None { pass := pass+1; }
    var q := QMsg("q-1", Wire.JStr("m"));
    if apply1(s, CPendingMessageSet("queued", "q-1", Wire.JStr("m"))).queuedMessages == [q] { pass := pass+1; }
    if apply1(s.(queuedMessages := [q, QMsg("q-2", Wire.JNull)]), CPendingMessageRemoved("queued", "q-1")).queuedMessages == [QMsg("q-2", Wire.JNull)] { pass := pass+1; }
    // truncation / reorder / turnsLoaded / contentChanged / reasoning (7)
    var three := s.(turns := [TDone("t1"), TDone("t2"), TDone("t3")]);
    if apply1(three, CTruncated(Some("t2"))).turns == [TDone("t1"), TDone("t2")] { pass := pass+1; }              // 063 keep up-to t2
    if apply1(three.(nextCursor := Some("cur")), CTruncated(None)) == three.(turns := [], nextCursor := None) { pass := pass+1; }  // 234 clear all
    var qa := QMsg("a", Wire.JStr("A")); var qb := QMsg("b", Wire.JStr("B")); var qc := QMsg("c", Wire.JStr("C"));
    if apply1(s.(queuedMessages := [qa, qb, qc]), CQueuedReordered(["c", "a", "b"])).queuedMessages == [qc, qa, qb] { pass := pass+1; }  // 053
    if apply1(s.(turns := [TDone("t3")]), CTurnsLoaded([TDone("t1"), TDone("t2")], Some("cur-0"))) == s.(turns := [TDone("t1"), TDone("t2"), TDone("t3")], nextCursor := Some("cur-0")) { pass := pass+1; }  // 232
    if apply1(s.(turns := [TDone("t3")]), CTurnsLoaded([TDone("t1"), TDone("t2"), TDone("t3")], None)).turns == [TDone("t1"), TDone("t2"), TDone("t3")] { pass := pass+1; }  // 233 dedup
    var running := act1.(activeTurn := Some(T0("t1").(parts := [PToolCall(TC0("tc-1", None).(status := "running"))])));
    var afterCc := apply1(running, CToolCallContentChanged("t1", "tc-1", Wire.JStr("terminal-ref"), None));
    var ccTc := afterCc.activeTurn.value.parts[0];
    if ccTc.PToolCall? && ccTc.tc.status == "running" && ccTc.tc.content == Some(Wire.JStr("terminal-ref")) { pass := pass+1; }  // 084
    var withR := act1.(activeTurn := Some(T0("t1").(parts := [PReasoning("r-1", "")])));
    var r2 := fold(withR, [CReasoning("t1", "r-1", "Thinking about "), CReasoning("t1", "r-1", "the plan")]);
    if r2.activeTurn == Some(T0("t1").(parts := [PReasoning("r-1", "Thinking about the plan")])) { pass := pass+1; }  // 032
  }

  // #338: input requests are LIVE response parts in the active turn. This models the input state
  // machine over `activeTurn.parts` (matching fixtures 105/106/108/235/236/237): request → OPEN part,
  // answerChanged patches drafts, completed resolves in place; no active turn → no-op.
  method runInputFlow() returns (pass: nat) {
    pass := 0;
    var s := C0();
    var openBody := map["message" := Wire.JStr("Clarify requirements"), "questions" := Wire.JArr([Wire.JObj(map["id" := Wire.JStr("q1"), "required" := Wire.JBool(true), "options" := Wire.JArr([Wire.JStr("node")])])]), "url" := Wire.JStr("https://example.test/review")];
    var at1 := s.(activeTurn := Some(T0("turn-1")), status := 8);   // active turn, InProgress(8)
    // 106: inputRequested with an active turn appends an OPEN input-request part, status → InputNeeded(24)
    var afterReq := apply1(at1, CInputRequested(InputReq("input-1", map[], openBody)));
    if afterReq.status == 24 && afterReq.activeTurn.Some? &&
       afterReq.activeTurn.value.parts == [PInputRequest(InputReq("input-1", map[], openBody), None)] { pass := pass+1; }
    // 236: inputRequested WITHOUT an active turn is a no-op (no side-seq to record it)
    if apply1(s.(status := 8), CInputRequested(InputReq("orphan-1", map[], openBody))) == s.(status := 8) { pass := pass+1; }
    var reqState := at1.(activeTurn := Some(T0("turn-1").(parts := [PInputRequest(InputReq("input-1", map[], openBody), None)])), status := 24);
    // answerChanged patches the draft on the open part in place
    if apply1(reqState, CInputAnswerChanged("input-1", "q1", Some(Wire.JStr("node")))).activeTurn.value.parts
       == [PInputRequest(InputReq("input-1", map["q1" := Wire.JStr("node")], openBody), None)] { pass := pass+1; }
    // 109-shape: answerChanged for an unknown request → no-op
    if apply1(reqState, CInputAnswerChanged("missing", "q1", None)) == reqState { pass := pass+1; }
    // completed resolves the part IN PLACE (response recorded), status → InProgress(8) (active turn, no open request)
    var afterDone := apply1(reqState, CInputCompleted("input-1", "accept", map[]));
    if afterDone.status == 8 && afterDone.activeTurn.Some? &&
       afterDone.activeTurn.value.parts == [PInputRequest(InputReq("input-1", map[], openBody), Some("accept"))] { pass := pass+1; }
    // 237: completed for an unknown id while a request stays open → no-op (status stays InputNeeded)
    if apply1(reqState, CInputCompleted("missing", "cancel", map[])) == reqState { pass := pass+1; }
    // 105: full flow — turnStarted → inputRequested → two draft syncs → completion; drafts preserved, body retained
    var flow := fold(s, [
      CTurnStarted("turn-1", Wire.JStr("Hello"), None),
      CInputRequested(InputReq("input-1", map[], openBody)),
      CInputAnswerChanged("input-1", "q1", Some(Wire.JStr("node"))),
      CInputAnswerChanged("input-1", "q2", Some(Wire.JStr("small"))),
      CInputCompleted("input-1", "accept", map[])]);
    if flow.status == 8 && flow.activeTurn.Some? &&
       flow.activeTurn.value.parts == [PInputRequest(InputReq("input-1", map["q1" := Wire.JStr("node"), "q2" := Wire.JStr("small")], openBody), Some("accept"))] { pass := pass+1; }
    // 235: completion final answers override synced drafts; open request body retained; status → InProgress(8)
    var overrideState := at1.(activeTurn := Some(T0("turn-1").(parts := [PInputRequest(InputReq("review-1", map["approve" := Wire.JBool(true)], openBody), None)])), status := 24);
    var overridden := apply1(overrideState, CInputCompleted("review-1", "decline", map["approve" := Wire.JBool(false)]));
    if overridden.status == 8 && overridden.activeTurn.Some? &&
       overridden.activeTurn.value.parts == [PInputRequest(InputReq("review-1", map["approve" := Wire.JBool(false)], openBody), Some("decline"))] { pass := pass+1; }
  }

  method {:isolate_assertions} runResultConfirm() returns (pass: nat) {
    pass := 0;
    var s := C0();
    var act1 := s.(activeTurn := Some(T0("t1")));
    // tool-call result confirmation (4)
    var rcLife := fold(act1, [
      CToolCallStart("t1", "tc-1", "bash", "Run Command", None, None, None),
      CToolCallReady("t1", "tc-1", Some("Run"), None, None, None, None, None, None, Some("not-needed"), None),
      CToolCallComplete("t1", "tc-1", true, Some("Done"), None, None, None, false, None),
      CToolCallResultConfirmed("t1", "tc-1", true, None)]);
    var rcTc := firstActiveToolCall(rcLife);
    if rcTc.Some? && rcTc.value.status == "completed" && rcTc.value.success == Some(true) { pass := pass+1; }  // 022 completed; result-confirm on already-completed is a no-op
    var rdLife := fold(act1, [
      CToolCallStart("t1", "tc-1", "bash", "Run Command", None, None, None),
      CToolCallReady("t1", "tc-1", Some("Run"), None, None, None, None, None, None, Some("not-needed"), None),
      CToolCallComplete("t1", "tc-1", true, Some("Done"), None, None, None, true, None),
      CToolCallResultConfirmed("t1", "tc-1", false, None)]);
    var rdTc := firstActiveToolCall(rdLife);
    if rdTc.Some? && rdTc.value.status == "cancelled" && rdTc.value.confirmed == None && rdTc.value.success == None && rdTc.value.reason == Some("result-denied") { pass := pass+1; }  // 023 denied → cancelled
    var runState := act1.(activeTurn := Some(T0("t1").(parts := [PToolCall(TC0("tc-1", None).(status := "running", confirmed := Some("user-action")))])), status := 8);
    if apply1(runState, CToolCallResultConfirmed("t1", "tc-1", true, None)) == runState { pass := pass+1; }  // 098 wrong status → no-op
    var selLife := fold(act1, [
      CToolCallStart("t1", "tc-1", "bash", "Run Command", None, None, None),
      CToolCallReady("t1", "tc-1", Some("Run"), None, None, None, None, None, None, None, None),
      CToolCallConfirmed("t1", "tc-1", true, Some("user-action"), None, None, None, None, None, None),
      CToolCallComplete("t1", "tc-1", true, Some("Done"), None, None, None, false, None),
      CToolCallResultConfirmed("t1", "tc-1", true, None)]);
    var selTc := firstActiveToolCall(selLife);
    if selTc.Some? && selTc.value.status == "completed" && selTc.value.confirmed == Some("user-action") && selTc.value.success == Some(true) { pass := pass+1; }  // 131 selectedOption path
  }

  method RunCorpus() returns (pass: nat, modeled: nat) {
    modeled := 54;
    var p1 := runScalarTurns();     // 15
    var p2 := runToolCalls();       // 15
    var p3 := runPendingHistory();  // 12
    var p4 := runInputFlow();       // 8
    var p5 := runResultConfirm();   // 4
    pass := p1 + p2 + p3 + p4 + p5;
  }
}
