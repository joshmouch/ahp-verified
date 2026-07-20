// AHP Dafny client — session channel (6th; the largest). PARTIAL: the tractable
// ~25 actions (bitset status flags, entity lists, config/meta/chat merge) are
// modeled + corpus-green. The mcpServer state-machine + customizationUpdated
// variants (nested children) are the documented remaining subset.
// status bits: inputNeeded=0x10, isRead=0x20, isArchived=0x40.
include "ahp_skeleton.dfy"

module Session {
  // Firewall: consumers get the channel's action/state types + the reducer surface, NOT the
  // internal seq-merge helpers (upsert*/remove*/upd*/setMcp/isMcp), bit ops (setBit/clearBit/B_*),
  // or the {:NoOp}/status lemmas. provides Wire/CC/AhpSkeleton because the exported signatures name
  // Wire.Json / CC.Fold / Option+ReduceOutcome (transitive-export rule).
  export
    provides ApplyToSession, apply1, fold, S0, RunCorpus, TitleChangedSetsTitle
    provides AhpSkeleton, Wire, CC
    provides custKey, chatKey, clientKey, inputKey        // rung-7: the id-keyed-collection key functions (codec's WfSessionState cites these)
    provides SessionWf, WfSessionActionInv, SessionWfPreserved   // rung-4/6: the reducer keyed/nullness invariant + its per-arm preservation theorem
    provides Apply1PreservesSessionWf, SessionWfWitness          // apply1-level restatement + concrete witness so ahp.dfy composes WfAhpStateInv over it
    reveals SessionState, SessionAction, Cust, Chat, Client, InputReq, SConfig

  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import CC = ConfluxContract
  import SR = ConfluxSeqRoute  // the canonical ordered-keyed-seq primitives (upsert/remove/update[-where] by id)

  datatype Cust = Cust(id: string, kind: string, uri: string, name: string, enabled: Option<bool>, state: Option<Wire.Json>, channel: Option<Wire.Json>)
  datatype Chat = Chat(resource: string, title: string, status: int, activity: Option<string>, modifiedAt: string, origin: Wire.Json)
  datatype Client = Client(clientId: string, raw: Wire.Json)
  datatype InputReq = InputReq(id: string, raw: Wire.Json)
  datatype SConfig = SConfig(schema: Wire.Json, values: map<string, Wire.Json>)

  datatype SessionState = SessionState(
    provider: string, title: string, status: bv32, lifecycle: string,
    activity: Option<string>, config: Option<SConfig>, meta: Option<Wire.Json>,
    creationError: Option<Wire.Json>, serverTools: Option<Wire.Json>, changesets: Option<Wire.Json>,
    // Exact PR #298 catalogue payloads, retained opaquely for replay.  Their
    // provider/declaration interpretation is provisional and has no durable
    // semantic law in core; only full-replacement lifecycle mechanics do.
    canvases: Option<Wire.Json>, openCanvases: Option<Wire.Json>,
    defaultChat: Option<string>,
    activeClients: seq<Client>, chats: seq<Chat>, customizations: seq<Cust>, inputNeeded: seq<InputReq>)

  datatype SessionAction =
    | Ready | CreationFailed(error: Wire.Json) | TitleChanged(title: string)
    | ServerToolsChanged(tools: Wire.Json) | MetaChanged(meta: Wire.Json)
    | IsReadChanged(isRead: bool) | IsArchivedChanged(isArchived: bool)
    | ActivityChanged(activity: Option<string>) | ConfigChanged(config: map<string, Wire.Json>, replace: bool)
    | CustomizationsChanged(customizations: seq<Cust>) | CustomizationToggled(id: string, enabled: bool)
    | CustomizationRemoved(id: string) | DefaultChatChanged(chat: Option<string>)
    | ChatAdded(summary: Chat) | ChatRemoved(resource: string) | ChatUpdated(resource: string, status: Option<int>, activity: Option<string>, modifiedAt: Option<string>)
    | ChangesetsChanged(changesets: Option<Wire.Json>) | ActiveClientSet(client: Client) | ActiveClientRemoved(clientId: string)
    | CanvasesChanged(canvases: Wire.Json) | OpenCanvasesChanged(openCanvases: Wire.Json)
    | InputNeededSet(req: InputReq) | InputNeededRemoved(id: string)
    | CustomizationUpdated(customization: Cust)
    | McpServerStateChanged(id: string, state: Wire.Json, channel: Option<Wire.Json>)
    | McpServerStartRequested(id: string) | McpServerStopRequested(id: string)
    | SUnknown(raw: Wire.Json)

  const B_INPUT: bv32 := 0x10
  const B_READ:  bv32 := 0x20
  const B_ARCH:  bv32 := 0x40
  function setBit(s: bv32, b: bv32): bv32 { s | b }
  function clearBit(s: bv32, b: bv32): bv32 { s & !b }

  // ── keyed-collection key functions (ONE shared value each; mirror of changeset's fileKey/opKey) ──
  // The four id-keyed seqs (customizations by id, chats by resource, activeClients by clientId, inputNeeded by
  // id) route through the ConfluxSeqRoute keyed algebra by these keyOf values. Sharing ONE const per key
  // (instead of an inline lambda per call site) makes every helper body, the codec's WfSessionState uniqueness
  // conjuncts, and each *PreservesUnique law citation reference the SAME function value — so the kernel laws'
  // conclusions line up syntactically with the helper results (no lambda-identity gap).
  const custKey:   Cust -> string     := (x: Cust) => x.id
  const chatKey:   Chat -> string     := (x: Chat) => x.resource
  const clientKey: Client -> string   := (x: Client) => x.clientId
  const inputKey:  InputReq -> string := (x: InputReq) => x.id

  // ── in-place patch functions (hoisted so a helper and its KeyPreserving law citation name the SAME value) ──
  function toggleEnabled(en: bool): Cust -> Cust { (e: Cust) => e.(enabled := Some(en)) }
  function chatPatch(st: Option<int>, ac: Option<string>, md: Option<string>): Chat -> Chat
  { (c: Chat) => c.(status := if st.Some? then st.value else c.status,
                    activity := if ac.Some? then ac else c.activity,
                    modifiedAt := if md.Some? then md.value else c.modifiedAt) }
  const mcpPred: Cust -> bool := (c: Cust) => c.kind == "mcpServer"
  function mcpPatch(st: Option<Wire.Json>, ch: Option<Wire.Json>): Cust -> Cust
  { (e: Cust) => e.(state := st, channel := ch) }

  // customizationToggled: update-in-place the matching customization's `enabled` flag by id (no-op if absent).
  function upsertCust(cs: seq<Cust>, id: string, en: bool): seq<Cust>
  { SR.SeqUpdateById(custKey, cs, id, toggleEnabled(en)) }
  function removeCust(cs: seq<Cust>, id: string): seq<Cust>
  { SR.SeqRemoveById(custKey, cs, id) }
  function removeChat(chs: seq<Chat>, r: string): seq<Chat>
  { SR.SeqRemoveById(chatKey, chs, r) }
  // chatAdded upserts by resource: replace an existing entry in place, else append.
  function upsertChat(chs: seq<Chat>, c: Chat): seq<Chat>
  { SR.SeqUpsertById(chatKey, chs, c) }
  // chatUpdated: merge status/activity/modifiedAt into the matching chat by resource (no-op if absent).
  function updChat(chs: seq<Chat>, r: string, st: Option<int>, ac: Option<string>, md: Option<string>): seq<Chat>
  { SR.SeqUpdateById(chatKey, chs, r, chatPatch(st, ac, md)) }
  function upsertClient(cl: seq<Client>, c: Client): seq<Client>
  { SR.SeqUpsertById(clientKey, cl, c) }
  function removeClient(cl: seq<Client>, id: string): seq<Client>
  { SR.SeqRemoveById(clientKey, cl, id) }
  function upsertInput(ins: seq<InputReq>, r: InputReq): seq<InputReq>
  { SR.SeqUpsertById(inputKey, ins, r) }
  function removeInput(ins: seq<InputReq>, id: string): seq<InputReq>
  { SR.SeqRemoveById(inputKey, ins, id) }
  function upsertCustFull(cs: seq<Cust>, c: Cust): seq<Cust>
  { SR.SeqUpsertById(custKey, cs, c) }
  // set state/channel on the matching mcpServer customization; no-op on a non-mcp id — a predicate-guarded
  // update-in-place (id AND kind=="mcpServer"), i.e. ConfluxSeqRoute.SeqUpdateByIdWhere.
  function setMcp(cs: seq<Cust>, id: string, st: Option<Wire.Json>, ch: Option<Wire.Json>): seq<Cust>
  { SR.SeqUpdateByIdWhere(custKey, mcpPred, cs, id, mcpPatch(st, ch)) }
  predicate isMcp(cs: seq<Cust>, id: string) { exists i :: 0 <= i < |cs| && cs[i].id == id && cs[i].kind == "mcpServer" }

  function ApplyToSession(s: SessionState, a: SessionAction, now: int): (SessionState, ReduceOutcome)
  {
    match a
    case Ready                  => (s.(lifecycle := "ready"), Applied)
    // TS `session/creationFailed`: error is `T | undefined` — a null wire payload CLEARS, never stores null.
    case CreationFailed(e)      => (s.(lifecycle := "creationFailed", creationError := if e == Wire.JNull then None else Some(e)), Applied)
    case TitleChanged(t)        => (s.(title := t), Applied)
    // TS `serverTools` / `_meta`: both `T | undefined` — collapse JNull → None (mirror of the decoder's optJsonNoNull).
    case ServerToolsChanged(t)  => (s.(serverTools := if t == Wire.JNull then None else Some(t)), Applied)
    case MetaChanged(m)         => (s.(meta := if m == Wire.JNull then None else Some(m)), Applied)
    case IsReadChanged(v)       => (s.(status := if v then setBit(s.status, B_READ) else clearBit(s.status, B_READ)), Applied)
    case IsArchivedChanged(v)   => (s.(status := if v then setBit(s.status, B_ARCH) else clearBit(s.status, B_ARCH)), Applied)
    case ActivityChanged(ac)    => (s.(activity := ac), Applied)
    case ConfigChanged(cfg, repl) => (match s.config case None => (s, NoOp) case Some(c) => (s.(config := Some(SConfig(c.schema, if repl then cfg else c.values + cfg))), Applied))
    case CustomizationsChanged(cs) => (s.(customizations := cs), Applied)
    case CustomizationToggled(id, en) => (s.(customizations := upsertCust(s.customizations, id, en)), Applied)
    case CustomizationRemoved(id) => (s.(customizations := removeCust(s.customizations, id)), Applied)
    case DefaultChatChanged(c)  => (s.(defaultChat := c), Applied)
    case ChatAdded(c)           => (s.(chats := upsertChat(s.chats, c)), Applied)
    case ChatRemoved(r)         => (s.(chats := removeChat(s.chats, r), defaultChat := if s.defaultChat == Some(r) then None else s.defaultChat), Applied)
    case ChatUpdated(r, st, ac, md) => (s.(chats := updChat(s.chats, r, st, ac, md)), Applied)
    case ChangesetsChanged(cs)  => (s.(changesets := cs), Applied)
    // TS `canvases` / `openCanvases`: both `T | undefined` — collapse JNull → None.
    case CanvasesChanged(cs)    => (s.(canvases := if cs == Wire.JNull then None else Some(cs)), Applied)
    case OpenCanvasesChanged(cs) => (s.(openCanvases := if cs == Wire.JNull then None else Some(cs)), Applied)
    case ActiveClientSet(c)     => (s.(activeClients := upsertClient(s.activeClients, c)), Applied)
    case ActiveClientRemoved(id) => (s.(activeClients := removeClient(s.activeClients, id)), Applied)
    case InputNeededSet(r)      => (s.(inputNeeded := upsertInput(s.inputNeeded, r), status := setBit(s.status, B_INPUT)), Applied)
    case InputNeededRemoved(id) =>
      var rem := removeInput(s.inputNeeded, id);
      (s.(inputNeeded := rem, status := if |rem| == 0 then clearBit(s.status, B_INPUT) else s.status), Applied)
    case CustomizationUpdated(c) => (s.(customizations := upsertCustFull(s.customizations, c)), Applied)
    // TS mcp customization `state` is `T | undefined` — collapse a JNull wire state to None (cleared), never store null.
    case McpServerStateChanged(id, st, ch) => if isMcp(s.customizations, id) then (s.(customizations := setMcp(s.customizations, id, if st == Wire.JNull then None else Some(st), ch)), Applied) else (s, NoOp)
    case McpServerStartRequested(id) => if isMcp(s.customizations, id) then (s.(customizations := setMcp(s.customizations, id, Some(Wire.JObj(map["kind" := Wire.JStr("starting")])), None)), Applied) else (s, NoOp)
    case McpServerStopRequested(id)  => if isMcp(s.customizations, id) then (s.(customizations := setMcp(s.customizations, id, Some(Wire.JObj(map["kind" := Wire.JStr("stopped")])), None)), Applied) else (s, NoOp)
    case SUnknown(_)            => (s, NoOp)
  }

  // ══════════════════════════════════════════════════════════════════════════
  //  RUNG 4 / 6 / 7 — the reducer's keyed-uniqueness + no-stored-null invariant,
  //  and its per-arm PRESERVATION theorem. (docs/shapes/04-invariant-bearing-state,
  //  06-invariant-preserving-reducer, 07-keyed-collection.)
  //
  //  The reducer lives in this module, so ApplyToSession and its keyed helpers are
  //  transparent here — this is the only module where the preservation proof can
  //  unfold the arms. It mirrors changeset.dfy's WfChangesetStatePreserved.
  //
  //  SCOPE NOTE: the codec's SessionCodec.WfSessionState (the state round-trip
  //  precondition) ALSO carries per-element wire-fidelity conjuncts on activeClients
  //  / inputNeeded (each stored raw's embedded clientId/id key matches the projected
  //  key). Those are decode-time fidelity facts phrased with the JSON field accessor
  //  (Fidelity.field), which the base reducer module deliberately does not import —
  //  a reducer must not know the wire layout. They are established by decode, not
  //  maintained by the reducer, so they sit OUTSIDE this reducer invariant. SessionWf
  //  carries exactly the reducer-side facts: the six `T | undefined` fields never
  //  store Some(JNull) (preserved by the JNull→None collapse in the arms above), each
  //  customization's state/channel is likewise non-null, and all four id-keyed
  //  collections have distinct keys.
  // ══════════════════════════════════════════════════════════════════════════
  predicate WfCust(c: Cust) {
    (c.state.Some?   ==> c.state.value   != Wire.JNull) &&
    (c.channel.Some? ==> c.channel.value != Wire.JNull)
  }
  predicate WfCusts(cs: seq<Cust>) { forall i :: 0 <= i < |cs| ==> WfCust(cs[i]) }

  predicate SessionWf(s: SessionState) {
    (s.meta.Some?          ==> s.meta.value          != Wire.JNull) &&
    (s.creationError.Some? ==> s.creationError.value != Wire.JNull) &&
    (s.serverTools.Some?   ==> s.serverTools.value   != Wire.JNull) &&
    (s.changesets.Some?    ==> s.changesets.value    != Wire.JNull) &&
    (s.canvases.Some?      ==> s.canvases.value      != Wire.JNull) &&
    (s.openCanvases.Some?  ==> s.openCanvases.value  != Wire.JNull) &&
    WfCusts(s.customizations) &&
    SR.UniqueKeys(custKey,   s.customizations) &&
    SR.UniqueKeys(chatKey,   s.chats) &&
    SR.UniqueKeys(clientKey, s.activeClients) &&
    SR.UniqueKeys(inputKey,  s.inputNeeded)
  }

  // rung-7 premise: the two arms that INSTALL a caller-supplied payload verbatim need it to already be
  // well-formed. customizationsChanged replaces the WHOLE customizations seq (needs incoming WfCusts +
  // distinct ids); customizationUpdated upserts one Cust (needs it non-null); changesetsChanged /
  // mcpServerStateChanged pass an Option<Json> straight through (needs it non-null-if-Some). Every OTHER
  // arm is a keyed upsert/remove/update whose invariant the kernel *PreservesUnique laws discharge, or a
  // JNull→None-collapsing scalar arm that is self-sufficient.
  predicate WfSessionActionInv(a: SessionAction) {
    match a
    case CustomizationsChanged(cs)       => WfCusts(cs) && SR.UniqueKeys(custKey, cs)
    case CustomizationUpdated(c)         => WfCust(c)
    case ChangesetsChanged(cs)           => cs.Some? ==> cs.value != Wire.JNull
    case McpServerStateChanged(_, _, ch) => ch.Some? ==> ch.value != Wire.JNull
    case _                               => true
  }

  // ── generic per-element-predicate preservation over the kernel keyed-seq ops ──
  // The kernel ships *PreservesUnique (key-uniqueness) but no per-element CONTENT law; customizations is the
  // only session collection whose well-formedness constrains element content (each Cust's state/channel is
  // non-null). These four small inductive lemmas carry WfCust through the exact ConfluxSeqRoute op each
  // customizations arm uses. (Proof lemmas over the kernel functions — not folds or consumer==primitive
  // correspondences; structurally identical to the codec's CustsRoundTrip induction.)
  lemma UpsertWfCusts(cs: seq<Cust>, c: Cust)
    requires WfCusts(cs) && WfCust(c)
    ensures  WfCusts(SR.SeqUpsertById(custKey, cs, c))
    decreases |cs|
  {
    if |cs| == 0 {
    } else if custKey(cs[0]) == custKey(c) {
    } else {
      UpsertWfCusts(cs[1..], c);
    }
  }
  lemma RemoveWfCusts(cs: seq<Cust>, id: string)
    requires WfCusts(cs)
    ensures  WfCusts(SR.SeqRemoveById(custKey, cs, id))
    decreases |cs|
  {
    if |cs| == 0 {
    } else {
      RemoveWfCusts(cs[1..], id);
    }
  }
  lemma UpdateWfCusts(cs: seq<Cust>, id: string, f: Cust -> Cust)
    requires WfCusts(cs) && (forall x :: WfCust(x) ==> WfCust(f(x)))
    ensures  WfCusts(SR.SeqUpdateById(custKey, cs, id, f))
    decreases |cs|
  {
    if |cs| == 0 {
    } else if custKey(cs[0]) == id {
    } else {
      UpdateWfCusts(cs[1..], id, f);
    }
  }
  lemma UpdateWhereWfCusts(cs: seq<Cust>, id: string, f: Cust -> Cust)
    requires WfCusts(cs) && (forall x :: WfCust(x) ==> WfCust(f(x)))
    ensures  WfCusts(SR.SeqUpdateByIdWhere(custKey, mcpPred, cs, id, f))
    decreases |cs|
  {
    if |cs| == 0 {
    } else if custKey(cs[0]) == id && mcpPred(cs[0]) {
    } else {
      UpdateWhereWfCusts(cs[1..], id, f);
    }
  }

  // ── rung 6: the reducer PRESERVES SessionWf on EVERY arm ──
  // Keyed-mutation arms cite the matching ConfluxSeqRoute *PreservesUnique law (uniqueness) and, for
  // customizations, the *Forall lemma above (per-element non-null). The whole-seq-replacement /
  // option-passthrough arms lean on WfSessionActionInv. The six scalar `T | undefined` arms are preserved by
  // the JNull→None collapse baked into ApplyToSession (the 6 reducer fixes).
  lemma SessionWfPreserved(s: SessionState, a: SessionAction, now: int)
    requires SessionWf(s) && WfSessionActionInv(a)
    ensures  SessionWf(ApplyToSession(s, a, now).0)
  {
    match a
    case Ready =>
    case CreationFailed(e) =>          // creationError := (collapse) — Some ⇒ non-null.
    case TitleChanged(t) =>
    case ServerToolsChanged(t) =>      // serverTools := (collapse).
    case MetaChanged(m) =>             // meta := (collapse).
    case IsReadChanged(v) =>
    case IsArchivedChanged(v) =>
    case ActivityChanged(ac) =>
    case ConfigChanged(cfg, repl) =>
    case CustomizationsChanged(cs) =>  // customizations := cs; WfSessionActionInv supplies WfCusts(cs) ∧ UniqueKeys(custKey,cs).
    case CustomizationToggled(id, en) =>
      assert forall x: Cust :: WfCust(x) ==> WfCust(toggleEnabled(en)(x));          // enabled-only patch keeps state/channel
      SR.UpdatePreservesUnique(custKey, s.customizations, id, toggleEnabled(en));   // enabled-only patch is key-preserving
      UpdateWfCusts(s.customizations, id, toggleEnabled(en));
    case CustomizationRemoved(id) =>
      SR.RemovePreservesUnique(custKey, s.customizations, id);
      RemoveWfCusts(s.customizations, id);
    case DefaultChatChanged(c) =>
    case ChatAdded(c) =>
      SR.UpsertPreservesUnique(chatKey, s.chats, c);
    case ChatRemoved(r) =>
      SR.RemovePreservesUnique(chatKey, s.chats, r);
    case ChatUpdated(r, st, ac, md) =>
      SR.UpdatePreservesUnique(chatKey, s.chats, r, chatPatch(st, ac, md));         // patch fixes resource ⇒ key-preserving
    case ChangesetsChanged(cs) =>      // changesets := cs; WfSessionActionInv supplies cs.Some? ⇒ non-null.
    case CanvasesChanged(cs) =>        // canvases := (collapse).
    case OpenCanvasesChanged(cs) =>    // openCanvases := (collapse).
    case ActiveClientSet(c) =>
      SR.UpsertPreservesUnique(clientKey, s.activeClients, c);
    case ActiveClientRemoved(id) =>
      SR.RemovePreservesUnique(clientKey, s.activeClients, id);
    case InputNeededSet(r) =>
      SR.UpsertPreservesUnique(inputKey, s.inputNeeded, r);
    case InputNeededRemoved(id) =>
      SR.RemovePreservesUnique(inputKey, s.inputNeeded, id);                        // removeInput preserves uniqueness; status bit change is invariant-irrelevant
    case CustomizationUpdated(c) =>
      SR.UpsertPreservesUnique(custKey, s.customizations, c);
      UpsertWfCusts(s.customizations, c);                                           // WfSessionActionInv supplies WfCust(c)
    case McpServerStateChanged(id, st, ch) =>
      if isMcp(s.customizations, id) {
        var st' := if st == Wire.JNull then None else Some(st);                     // the fix: collapse a null wire state
        assert forall x: Cust :: WfCust(mcpPatch(st', ch)(x));                      // st' non-null-if-Some (collapse); ch non-null-if-Some (WfSessionActionInv)
        SR.UpdateWherePreservesUnique(custKey, mcpPred, s.customizations, id, mcpPatch(st', ch));
        UpdateWhereWfCusts(s.customizations, id, mcpPatch(st', ch));
      }
    case McpServerStartRequested(id) =>
      if isMcp(s.customizations, id) {
        var f := mcpPatch(Some(Wire.JObj(map["kind" := Wire.JStr("starting")])), None);
        assert forall x: Cust :: WfCust(f(x));
        SR.UpdateWherePreservesUnique(custKey, mcpPred, s.customizations, id, f);
        UpdateWhereWfCusts(s.customizations, id, f);
      }
    case McpServerStopRequested(id) =>
      if isMcp(s.customizations, id) {
        var f := mcpPatch(Some(Wire.JObj(map["kind" := Wire.JStr("stopped")])), None);
        assert forall x: Cust :: WfCust(f(x));
        SR.UpdateWherePreservesUnique(custKey, mcpPred, s.customizations, id, f);
        UpdateWhereWfCusts(s.customizations, id, f);
      }
    case SUnknown(_) =>
  }

  // Non-vacuity witness: a GENUINELY populated well-formed state (a real Cust in a seq, two distinctly-keyed
  // chats) whose reducer image under a keyed arm is STILL well-formed — so SessionWfPreserved is not vacuous.
  lemma SessionWfPreserved_NonVacuous()
    ensures exists s: SessionState, a: SessionAction ::
      SessionWf(s) && WfSessionActionInv(a) && SessionWf(ApplyToSession(s, a, 0).0)
  {
    var mc := Cust("mcp-1", "mcpServer", "file:///w", "MCP", Some(true), None, None);
    var s := S0().(customizations := [mc],
                   chats := [Chat("ahp-chat:/c1", "C1", 1, None, "t1", Wire.JNull)]);
    assert SR.UniqueKeys(custKey, s.customizations);
    assert SR.UniqueKeys(chatKey, s.chats);
    assert SessionWf(s);
    var a := CustomizationToggled("mcp-1", false);
    assert WfSessionActionInv(a);
    SessionWfPreserved(s, a, 0);
    assert SessionWf(ApplyToSession(s, a, 0).0);
  }

  lemma TitleChangedSetsTitle(s: SessionState, title: string, now: int)
    ensures ApplyToSession(s, TitleChanged(title), now).0.title == title
  {
  }

  lemma S_UnknownIsNoOp(s: SessionState, a: SessionAction, now: int)
    requires a.SUnknown?
    ensures ApplyToSession(s, a, now).0 == s && ApplyToSession(s, a, now).1 == NoOp {}
  lemma S_UnknownIsNoOp_NonVacuous() ensures exists a: SessionAction :: a.SUnknown? { assert (SUnknown(Wire.JNull)).SUnknown?; }
  // status bitset laws (isRead/isArchived/inputNeeded set-then-observe)
  lemma S_IsReadSetsBit(s: SessionState, now: int)
    ensures ApplyToSession(s, IsReadChanged(true), now).0.status & B_READ == B_READ {}

  function apply1(s: SessionState, a: SessionAction): SessionState { ApplyToSession(s, a, 9999).0 }
  // apply1-level restatement of the rung-6 preservation theorem, so the aggregate (ahp.dfy) — which dispatches
  // through apply1, not ApplyToSession — can compose it without unfolding apply1's (exported-opaque) body.
  lemma Apply1PreservesSessionWf(s: SessionState, a: SessionAction)
    requires SessionWf(s) && WfSessionActionInv(a)
    ensures  SessionWf(apply1(s, a))
  { SessionWfPreserved(s, a, 9999); }
  // Concrete, populated well-formed witness (state + a keyed action) handed to the aggregate (ahp.dfy) so its
  // WfAhpStateInv non-vacuity witness need not unfold this module's exported-opaque SessionWf/WfSessionActionInv.
  lemma SessionWfWitness() returns (s: SessionState, a: SessionAction)
    ensures SessionWf(s) && WfSessionActionInv(a) && SessionWf(apply1(s, a))
  {
    var mc := Cust("mcp-1", "mcpServer", "file:///w", "MCP", Some(true), None, None);
    s := S0().(customizations := [mc],
               chats := [Chat("ahp-chat:/c1", "C1", 1, None, "t1", Wire.JNull)]);
    a := CustomizationToggled("mcp-1", false);
    Apply1PreservesSessionWf(s, a);
  }
  function fold(s: SessionState, acts: seq<SessionAction>): SessionState
  { CC.Fold(apply1, s, acts) }

  // base session (status 1, lifecycle creating)
  function S0(): SessionState {
    SessionState("copilot", "Test Session", 1, "creating", None, None, None, None, None, None,
      None, None, None, [], [], [], [])
  }

  method RunCorpus() returns (pass: nat, modeled: nat) {
    modeled := 36; pass := 0;
    var s := S0();
    // 003 ready / 147 preserves status (status stays, lifecycle ready)
    if apply1(s.(lifecycle := "creating"), Ready) == s.(lifecycle := "ready") { pass := pass+1; }
    // 004 creationFailed
    if apply1(s, CreationFailed(Wire.JNull)).lifecycle == "creationFailed" { pass := pass+1; }
    // titleChanged
    if apply1(s, TitleChanged("New")).title == "New" { pass := pass+1; }
    // 034 serverToolsChanged
    var tools := Wire.JArr([Wire.JObj(map["name" := Wire.JStr("bash")])]);
    if apply1(s, ServerToolsChanged(tools)).serverTools == Some(tools) { pass := pass+1; }
    // 135 metaChanged replaces
    var m := Wire.JObj(map["git" := Wire.JStr("feature")]);
    if apply1(s.(meta := Some(Wire.JNull)), MetaChanged(m)).meta == Some(m) { pass := pass+1; }
    // 071/072 isReadChanged: 1→33 (|0x20), 33→1 (&~0x20)
    if apply1(s, IsReadChanged(true)).status == 33 { pass := pass+1; }
    if apply1(s.(status := 33), IsReadChanged(false)).status == 1 { pass := pass+1; }
    // 073/074 isArchivedChanged: 1→65 (|0x40), 65→1
    if apply1(s, IsArchivedChanged(true)).status == 65 { pass := pass+1; }
    if apply1(s.(status := 65), IsArchivedChanged(false)).status == 1 { pass := pass+1; }
    // 133/134 activityChanged set/clear (status preserved)
    if apply1(s, ActivityChanged(Some("Running"))) == s.(activity := Some("Running")) { pass := pass+1; }
    if apply1(s.(activity := Some("x")), ActivityChanged(None)) == s { pass := pass+1; }
    // 113 configChanged merges (replace=false keeps target, overrides baseBranch)
    var cfg := SConfig(Wire.JNull, map["target" := Wire.JStr("worktree"), "baseBranch" := Wire.JStr("main")]);
    var cfgExp := SConfig(Wire.JNull, map["target" := Wire.JStr("worktree"), "baseBranch" := Wire.JStr("develop")]);
    if apply1(s.(config := Some(cfg)), ConfigChanged(map["baseBranch" := Wire.JStr("develop")], false)).config == Some(cfgExp) { pass := pass+1; }
    // 129 configChanged replace=true drops all prior values, keeps only the new map
    var cfgRepl := SConfig(Wire.JNull, map["baseBranch" := Wire.JStr("develop")]);
    if apply1(s.(config := Some(cfg)), ConfigChanged(map["baseBranch" := Wire.JStr("develop")], true)).config == Some(cfgRepl) { pass := pass+1; }
    // configChanged no-op when no config
    if apply1(s, ConfigChanged(map["x" := Wire.JNull], false)) == s { pass := pass+1; }
    // 058 customizationsChanged replaces
    var pa := Cust("plugin-a", "plugin", "https://plugins.example/a", "Plugin A", Some(true), None, None);
    var pb := Cust("plugin-b", "plugin", "https://plugins.example/b", "Plugin B", Some(true), None, None);
    if apply1(s, CustomizationsChanged([pa, pb])).customizations == [pa, pb] { pass := pass+1; }
    // 060 customizationToggled by id
    if apply1(s.(customizations := [pa, pb]), CustomizationToggled("plugin-a", false)).customizations == [pa.(enabled := Some(false)), pb] { pass := pass+1; }
    // 152 customizationRemoved by id
    if apply1(s.(customizations := [pa, pb]), CustomizationRemoved("plugin-a")).customizations == [pb] { pass := pass+1; }
    // defaultChatChanged sets Some
    if apply1(s, DefaultChatChanged(Some("c1"))).defaultChat == Some("c1") { pass := pass+1; }
    // 160 defaultChatChanged with absent/None unsets to None
    if apply1(s.(defaultChat := Some("ahp-chat:/old")), DefaultChatChanged(None)).defaultChat == None { pass := pass+1; }
    // 170 chatAdded appends a new resource
    var c1 := Chat("ahp-chat:/c1", "Chat 1", 1, None, "t1", Wire.JObj(map["kind" := Wire.JStr("user")]));
    if apply1(s, ChatAdded(c1)).chats == [c1] { pass := pass+1; }
    // 171 chatAdded upserts an existing resource in place (renamed replaces, no dup)
    var c1r := c1.(title := "Chat 1 (renamed)", status := 8, modifiedAt := "t2");
    if apply1(s.(chats := [c1]), ChatAdded(c1r)).chats == [c1r] { pass := pass+1; }
    // 172 chatRemoved by resource + clears defaultChat if match
    var c2 := Chat("ahp-chat:/c2", "Chat 2", 8, Some("Thinking"), "t2", Wire.JNull);
    if apply1(s.(chats := [c1, c2], defaultChat := Some("ahp-chat:/c1")), ChatRemoved("ahp-chat:/c1")) == s.(chats := [c2], defaultChat := None) { pass := pass+1; }
    // 173 chatUpdated merges changes
    var c1u := c1.(status := 24, activity := Some("Waiting for approval"), modifiedAt := "t2");
    if apply1(s.(chats := [c1]), ChatUpdated("ahp-chat:/c1", Some(24), Some("Waiting for approval"), Some("t2"))).chats == [c1u] { pass := pass+1; }
    // 145 changesetsChanged sets Some
    if apply1(s, ChangesetsChanged(Some(Wire.JArr([])))).changesets == Some(Wire.JArr([])) { pass := pass+1; }
    // 146 changesetsChanged with null/absent payload clears to None
    if apply1(s.(changesets := Some(Wire.JArr([]))), ChangesetsChanged(None)).changesets == None { pass := pass+1; }
    // 232/233 PR #298 provisional snapshot catalogues: exact payload, full replacement.
    var canvasDecls := Wire.JArr([Wire.JObj(map[
      "providerId" := Wire.JStr("acme.editors"), "canvasId" := Wire.JStr("markdown")])]);
    if apply1(s, CanvasesChanged(canvasDecls)).canvases == Some(canvasDecls) { pass := pass+1; }
    var openRefs := Wire.JArr([Wire.JObj(map[
      "channel" := Wire.JStr("ahp-canvas:/canvas-1"), "canvasId" := Wire.JStr("markdown"),
      "availability" := Wire.JStr("ready")])]);
    if apply1(s, OpenCanvasesChanged(openRefs)).openCanvases == Some(openRefs) { pass := pass+1; }
    // 035 activeClientSet + 221 activeClientRemoved
    var cl := Client("vscode-1", Wire.JObj(map["displayName" := Wire.JStr("VS Code")]));
    if apply1(s, ActiveClientSet(cl)).activeClients == [cl] { pass := pass+1; }
    if apply1(s.(activeClients := [cl, Client("cli-1", Wire.JNull)]), ActiveClientRemoved("vscode-1")).activeClients == [Client("cli-1", Wire.JNull)] { pass := pass+1; }
    // 223 inputNeededSet promotes status |0x10 (8→24); 226 removed keeps status if entries remain
    var ir := InputReq("req-1", Wire.JNull);
    if apply1(s.(status := 8), InputNeededSet(ir)).status == 24 { pass := pass+1; }
    // 226 inputNeededRemoved keeps InputNeeded status (0x10) while entries remain (24→24)
    var i1 := InputReq("a", Wire.JNull); var i2 := InputReq("b", Wire.JNull);
    if apply1(s.(status := 24, inputNeeded := [i1, i2]), InputNeededRemoved("a")).status == 24 { pass := pass+1; }
    // ── mcpServer state machine + customizationUpdated (completes session's action-type coverage) ──
    var mc := Cust("mcp-1", "mcpServer", "file:///w", "MCP", Some(true), None, None);
    // 235 mcpServerStartRequested → state:starting, channel cleared
    if apply1(s.(customizations := [mc]), McpServerStartRequested("mcp-1")).customizations == [mc.(state := Some(Wire.JObj(map["kind" := Wire.JStr("starting")])), channel := None)] { pass := pass+1; }
    // 227 mcpServerStateChanged sets state+channel on the mcp customization
    if apply1(s.(customizations := [mc]), McpServerStateChanged("mcp-1", Wire.JObj(map["kind" := Wire.JStr("ready")]), Some(Wire.JStr("ch")))).customizations == [mc.(state := Some(Wire.JObj(map["kind" := Wire.JStr("ready")])), channel := Some(Wire.JStr("ch")))] { pass := pass+1; }
    // 162 mcpServerStateChanged is a no-op on a non-mcp id
    if apply1(s.(customizations := [pa]), McpServerStateChanged("plugin-a", Wire.JObj(map["kind" := Wire.JStr("ready")]), Some(Wire.JStr("nope")))) == s.(customizations := [pa]) { pass := pass+1; }
    // 236/237 mcpServerStopRequested → state:stopped (auth-required + nested both go to stopped, per the real fixtures)
    if apply1(s.(customizations := [mc]), McpServerStopRequested("mcp-1")).customizations == [mc.(state := Some(Wire.JObj(map["kind" := Wire.JStr("stopped")])), channel := None)] { pass := pass+1; }
    // 139/151 customizationUpdated upserts the full customization by id
    if apply1(s, CustomizationUpdated(mc)).customizations == [mc] { pass := pass+1; }
  }
}
