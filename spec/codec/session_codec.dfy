// AHP Dafny client — FULL Wire.Json ⇄ SessionAction codec (Session channel).
//
// Exact analogue of spec/fidelity.dfy's Root codec, for the SESSION channel:
//   • DecodeSessionAction / EncodeSessionAction  — the typed wire codec
//   • encodeSessionState                          — deterministic total serializer
//   • SessionActionRoundTrip (+ _NonVacuous)      — decode ∘ encode == id (typed)
//   • SUnknown_EncodeVerbatim / SUnknownVariant_Verbatim — the honest carve-out
//
// DECODE is PROMOTED from spec/replay/session_replay.dfy's proven `decodeSessionAction`
// (the exact logic that drives the 61 real session-reducer fixtures in the harness);
// ENCODE is its inverse (canonical wire object per typed variant). The shared field
// accessors (field/asStr/asArr/asObjMap/asInt/asBool) are REUSED from Fidelity — not
// duplicated. The session-specific optional decoders (optStr/optInt/optBool/
// optJsonNoNull) are promoted from the replay harness, since Fidelity has no analogue.
//
// ── The honest well-formedness precondition (WfSessionAction) ──
// Root's codec round-trips every typed variant UNCONDITIONALLY because its sub-codecs
// (AgentInfo) carry only strings/seqs — no null-collapsing optionals, no redundant
// stored keys. Session's decode is NOT injective on three shapes, so the round-trip is
// scoped by a stated, quantified, NON-VACUOUS precondition (exactly the invariant that
// `decodeSessionAction` itself establishes on its outputs):
//   1. Cust.state / Cust.channel and McpServerStateChanged.channel / ChangesetsChanged
//      decode through `optJsonNoNull`, which collapses BOTH absent and JSON `null` to
//      None. A `Some(Wire.JNull)` is therefore not in decode's image, so it cannot round-trip.
//   2. Client / InputReq store a redundant identity key (clientId / id) ALONGSIDE the
//      full raw Wire.Json. Encode must re-emit the raw verbatim (to preserve it), so the
//      round-trip holds iff the stored key matches the one embedded in raw — precisely
//      what decode guarantees by construction.
// This is the SAME shape as Fidelity's RootUnknown carve-out: the property holds exactly
// on the image of decode, and the precondition names that image. The witness lemma
// exercises the inductive seq path with a real Cust, so the lemma is not constants-only.

include "../session.dfy"
include "../fidelity.dfy"

module SessionCodec {
  export
    provides DecodeSessionAction, encodeSessionState, decodeSessionState
    provides WfSessionAction, SessionActionRoundTrip, SessionStateRoundTrip
    provides encodeCust, encodeCusts, encodeChat, encodeClient, encodeInputReq
    provides Session, AhpSkeleton, Wire, F, SR
    reveals EncodeSessionAction, WfSessionState, WfSessionStateWire, WfCusts, WfCust, WfClients, WfInputs

  import opened Session
  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import F = Fidelity
  import CO = ConfluxOperators
  import SR = ConfluxSeqRoute  // rung-7: UniqueKeys for the state's four id-keyed collections (keyed by Session.custKey/chatKey/clientKey/inputKey)

  // ── session-specific optional decoders (promoted from session_replay.dfy) ──
  function optStr(o: Option<Wire.Json>): Option<string> { if o.Some? && o.value.JStr? then Some(o.value.s) else None }
  // optInt (JNum + integer-valued JDec) lives in AhpSkeleton now; in scope via `import opened AhpSkeleton`.
  function optBool(o: Option<Wire.Json>): Option<bool> { if o.Some? && o.value.JBool? then Some(o.value.b) else None }
  // present-and-NON-null → Some(raw); absent OR JSON `null` → None. (mcpServer cleared
  // channels, changesets clears, and any unset state/channel decode to None here.)
  function optJsonNoNull(o: Option<Wire.Json>): Option<Wire.Json> { if o.Some? && !o.value.JNull? then Some(o.value) else None }

  // ══════════════════════════════════════════════════════════════════════════
  //  DECODE domain values (PROMOTED verbatim from session_replay.dfy; shared
  //  accessors sourced from Fidelity as F.field / F.asStr / F.asInt / F.asArr).
  // ══════════════════════════════════════════════════════════════════════════
  function decodeCust(j: Wire.Json): Cust {
    Cust(F.asStr(F.field(j, "id")), F.asStr(F.field(j, "type")), F.asStr(F.field(j, "uri")),
         F.asStr(F.field(j, "name")), optBool(F.field(j, "enabled")),
         optJsonNoNull(F.field(j, "state")), optJsonNoNull(F.field(j, "channel")))
  }
  function decodeCusts(js: seq<Wire.Json>): seq<Cust>
  { CO.Map(decodeCust, js) }

  function decodeChat(j: Wire.Json): Chat {
    Chat(F.asStr(F.field(j, "resource")), F.asStr(F.field(j, "title")), asInt(F.field(j, "status")),
         optStr(F.field(j, "activity")), F.asStr(F.field(j, "modifiedAt")),
         if F.field(j, "origin").Some? then F.field(j, "origin").value else Wire.JNull)
  }
  function decodeClient(j: Wire.Json): Client { Client(F.asStr(F.field(j, "clientId")), j) }
  function decodeInputReq(j: Wire.Json): InputReq { InputReq(F.asStr(F.field(j, "id")), j) }

  // ══════════════════════════════════════════════════════════════════════════
  //  ENCODE domain values (the INVERSE of the decode above). Optionals emit an
  //  explicit `null` placeholder when None — decode collapses absent==null==None
  //  for these fields, so a fixed-shape object is both canonical and round-trips.
  // ══════════════════════════════════════════════════════════════════════════
  function encodeCust(c: Cust): Wire.Json {
    Wire.JObj(map["id" := Wire.JStr(c.id), "type" := Wire.JStr(c.kind), "uri" := Wire.JStr(c.uri), "name" := Wire.JStr(c.name),
             "enabled" := (if c.enabled.Some? then Wire.JBool(c.enabled.value) else Wire.JNull),
             "state" := (if c.state.Some? then c.state.value else Wire.JNull),
             "channel" := (if c.channel.Some? then c.channel.value else Wire.JNull)])
  }
  function encodeCusts(cs: seq<Cust>): seq<Wire.Json>
  { CO.Map(encodeCust, cs) }

  function encodeChat(c: Chat): Wire.Json {
    Wire.JObj(map["resource" := Wire.JStr(c.resource), "title" := Wire.JStr(c.title), "status" := Wire.JNum(c.status),
             "activity" := (if c.activity.Some? then Wire.JStr(c.activity.value) else Wire.JNull),
             "modifiedAt" := Wire.JStr(c.modifiedAt), "origin" := c.origin])
  }
  function encodeChats(chs: seq<Chat>): seq<Wire.Json>
  { CO.Map(encodeChat, chs) }

  // Client / InputReq are stored as (identity-key, full-raw-Wire.Json); the raw IS the
  // canonical wire object (decode reads the key back out of it).
  function encodeClient(c: Client): Wire.Json { c.raw }
  function encodeClients(cl: seq<Client>): seq<Wire.Json>
  { CO.Map(encodeClient, cl) }

  function encodeInputReq(r: InputReq): Wire.Json { r.raw }
  function encodeInputs(ins: seq<InputReq>): seq<Wire.Json>
  { CO.Map(encodeInputReq, ins) }

  // ══════════════════════════════════════════════════════════════════════════
  //  THE TYPED CODEC — dispatch on the "type" string; unknown → SUnknown (verbatim).
  // ══════════════════════════════════════════════════════════════════════════
  function DecodeSessionAction(j: Wire.Json): SessionAction {
    var t := F.asStr(F.field(j, "type"));
    if      t == "session/ready"                  then Ready
    else if t == "session/creationFailed"         then CreationFailed(if F.field(j, "error").Some? then F.field(j, "error").value else Wire.JNull)
    else if t == "session/titleChanged"           then TitleChanged(F.asStr(F.field(j, "title")))
    else if t == "session/serverToolsChanged"     then ServerToolsChanged(if F.field(j, "tools").Some? then F.field(j, "tools").value else Wire.JNull)
    else if t == "session/metaChanged"            then MetaChanged(if F.field(j, "_meta").Some? then F.field(j, "_meta").value else Wire.JNull)
    else if t == "session/isReadChanged"          then IsReadChanged(F.asBool(F.field(j, "isRead"), false))
    else if t == "session/isArchivedChanged"      then IsArchivedChanged(F.asBool(F.field(j, "isArchived"), false))
    else if t == "session/activityChanged"        then ActivityChanged(optStr(F.field(j, "activity")))
    else if t == "session/configChanged"          then ConfigChanged(F.asObjMap(F.field(j, "config")), F.asBool(F.field(j, "replace"), false))
    else if t == "session/customizationsChanged"  then CustomizationsChanged(decodeCusts(F.asArr(F.field(j, "customizations"))))
    else if t == "session/customizationToggled"   then CustomizationToggled(F.asStr(F.field(j, "id")), F.asBool(F.field(j, "enabled"), false))
    else if t == "session/customizationRemoved"   then CustomizationRemoved(F.asStr(F.field(j, "id")))
    else if t == "session/customizationUpdated"   then CustomizationUpdated(decodeCust(if F.field(j, "customization").Some? then F.field(j, "customization").value else Wire.JNull))
    else if t == "session/defaultChatChanged"     then DefaultChatChanged(optStr(F.field(j, "defaultChat")))
    else if t == "session/chatAdded"              then ChatAdded(decodeChat(if F.field(j, "summary").Some? then F.field(j, "summary").value else Wire.JNull))
    else if t == "session/chatRemoved"            then ChatRemoved(F.asStr(F.field(j, "chat")))
    else if t == "session/chatUpdated"            then
      var ch := F.field(j, "changes");
      var chObj := if ch.Some? then ch.value else Wire.JNull;
      ChatUpdated(F.asStr(F.field(j, "chat")), optInt(F.field(chObj, "status")), optStr(F.field(chObj, "activity")), optStr(F.field(chObj, "modifiedAt")))
    else if t == "session/changesetsChanged"      then ChangesetsChanged(optJsonNoNull(F.field(j, "changesets")))
    else if t == "session/canvasesChanged"        then CanvasesChanged(if F.field(j, "canvases").Some? then F.field(j, "canvases").value else Wire.JArr([]))
    else if t == "session/openCanvasesChanged"    then OpenCanvasesChanged(if F.field(j, "openCanvases").Some? then F.field(j, "openCanvases").value else Wire.JArr([]))
    else if t == "session/activeClientSet"        then ActiveClientSet(decodeClient(if F.field(j, "activeClient").Some? then F.field(j, "activeClient").value else Wire.JNull))
    else if t == "session/activeClientRemoved"    then ActiveClientRemoved(F.asStr(F.field(j, "clientId")))
    else if t == "session/inputNeededSet"         then InputNeededSet(decodeInputReq(if F.field(j, "request").Some? then F.field(j, "request").value else Wire.JNull))
    else if t == "session/inputNeededRemoved"     then InputNeededRemoved(F.asStr(F.field(j, "id")))
    else if t == "session/mcpServerStateChanged"  then McpServerStateChanged(F.asStr(F.field(j, "id")), if F.field(j, "state").Some? then F.field(j, "state").value else Wire.JNull, optJsonNoNull(F.field(j, "channel")))
    else if t == "session/mcpServerStartRequested" then McpServerStartRequested(F.asStr(F.field(j, "id")))
    else if t == "session/mcpServerStopRequested"  then McpServerStopRequested(F.asStr(F.field(j, "id")))
    else SUnknown(j)
  }

  function EncodeSessionAction(a: SessionAction): Wire.Json {
    match a
    case Ready                          => Wire.JObj(map["type" := Wire.JStr("session/ready")])
    case CreationFailed(e)              => Wire.JObj(map["type" := Wire.JStr("session/creationFailed"), "error" := e])
    case TitleChanged(t)                => Wire.JObj(map["type" := Wire.JStr("session/titleChanged"), "title" := Wire.JStr(t)])
    case ServerToolsChanged(tls)        => Wire.JObj(map["type" := Wire.JStr("session/serverToolsChanged"), "tools" := tls])
    case MetaChanged(m)                 => Wire.JObj(map["type" := Wire.JStr("session/metaChanged"), "_meta" := m])
    case IsReadChanged(v)               => Wire.JObj(map["type" := Wire.JStr("session/isReadChanged"), "isRead" := Wire.JBool(v)])
    case IsArchivedChanged(v)           => Wire.JObj(map["type" := Wire.JStr("session/isArchivedChanged"), "isArchived" := Wire.JBool(v)])
    case ActivityChanged(ac)            => Wire.JObj(map["type" := Wire.JStr("session/activityChanged"), "activity" := (if ac.Some? then Wire.JStr(ac.value) else Wire.JNull)])
    case ConfigChanged(cfg, repl)       => Wire.JObj(map["type" := Wire.JStr("session/configChanged"), "config" := Wire.JObj(cfg), "replace" := Wire.JBool(repl)])
    case CustomizationsChanged(cs)      => Wire.JObj(map["type" := Wire.JStr("session/customizationsChanged"), "customizations" := Wire.JArr(encodeCusts(cs))])
    case CustomizationToggled(id, en)   => Wire.JObj(map["type" := Wire.JStr("session/customizationToggled"), "id" := Wire.JStr(id), "enabled" := Wire.JBool(en)])
    case CustomizationRemoved(id)       => Wire.JObj(map["type" := Wire.JStr("session/customizationRemoved"), "id" := Wire.JStr(id)])
    case DefaultChatChanged(c)          => Wire.JObj(map["type" := Wire.JStr("session/defaultChatChanged"), "defaultChat" := (if c.Some? then Wire.JStr(c.value) else Wire.JNull)])
    case ChatAdded(sum)                 => Wire.JObj(map["type" := Wire.JStr("session/chatAdded"), "summary" := encodeChat(sum)])
    case ChatRemoved(r)                 => Wire.JObj(map["type" := Wire.JStr("session/chatRemoved"), "chat" := Wire.JStr(r)])
    case ChatUpdated(r, st, ac, md)     => Wire.JObj(map["type" := Wire.JStr("session/chatUpdated"), "chat" := Wire.JStr(r),
                                                   "changes" := Wire.JObj(map["status" := (if st.Some? then Wire.JNum(st.value) else Wire.JNull),
                                                                         "activity" := (if ac.Some? then Wire.JStr(ac.value) else Wire.JNull),
                                                                         "modifiedAt" := (if md.Some? then Wire.JStr(md.value) else Wire.JNull)])])
    case ChangesetsChanged(cs)          => Wire.JObj(map["type" := Wire.JStr("session/changesetsChanged"), "changesets" := (if cs.Some? then cs.value else Wire.JNull)])
    case CanvasesChanged(cs)            => Wire.JObj(map["type" := Wire.JStr("session/canvasesChanged"), "canvases" := cs])
    case OpenCanvasesChanged(cs)        => Wire.JObj(map["type" := Wire.JStr("session/openCanvasesChanged"), "openCanvases" := cs])
    case ActiveClientSet(c)             => Wire.JObj(map["type" := Wire.JStr("session/activeClientSet"), "activeClient" := encodeClient(c)])
    case ActiveClientRemoved(id)        => Wire.JObj(map["type" := Wire.JStr("session/activeClientRemoved"), "clientId" := Wire.JStr(id)])
    case InputNeededSet(r)              => Wire.JObj(map["type" := Wire.JStr("session/inputNeededSet"), "request" := encodeInputReq(r)])
    case InputNeededRemoved(id)         => Wire.JObj(map["type" := Wire.JStr("session/inputNeededRemoved"), "id" := Wire.JStr(id)])
    case CustomizationUpdated(c)        => Wire.JObj(map["type" := Wire.JStr("session/customizationUpdated"), "customization" := encodeCust(c)])
    case McpServerStateChanged(id, st, ch) => Wire.JObj(map["type" := Wire.JStr("session/mcpServerStateChanged"), "id" := Wire.JStr(id), "state" := st,
                                                       "channel" := (if ch.Some? then ch.value else Wire.JNull)])
    case McpServerStartRequested(id)    => Wire.JObj(map["type" := Wire.JStr("session/mcpServerStartRequested"), "id" := Wire.JStr(id)])
    case McpServerStopRequested(id)     => Wire.JObj(map["type" := Wire.JStr("session/mcpServerStopRequested"), "id" := Wire.JStr(id)])
    case SUnknown(raw)                  => raw
  }

  // ══════════════════════════════════════════════════════════════════════════
  //  WELL-FORMEDNESS — exactly the invariant `decodeSessionAction` establishes on
  //  its outputs. The round-trip holds precisely on the image of decode.
  // ══════════════════════════════════════════════════════════════════════════
  predicate WfCust(c: Cust) {
    (c.state.Some? ==> c.state.value != Wire.JNull) && (c.channel.Some? ==> c.channel.value != Wire.JNull)
  }
  predicate WfCusts(cs: seq<Cust>) { forall i :: 0 <= i < |cs| ==> WfCust(cs[i]) }

  predicate WfSessionAction(a: SessionAction) {
    match a
    case CustomizationsChanged(cs)        => WfCusts(cs)
    case CustomizationUpdated(c)          => WfCust(c)
    case McpServerStateChanged(_, _, ch)  => ch.Some? ==> ch.value != Wire.JNull
    case ChangesetsChanged(cs)            => cs.Some? ==> cs.value != Wire.JNull
    case ActiveClientSet(c)               => F.asStr(F.field(c.raw, "clientId")) == c.clientId
    case InputNeededSet(r)                => F.asStr(F.field(r.raw, "id")) == r.id
    case _                                => true
  }

  // ── sub-codec round-trip proofs (decode ∘ encode == id) ──
  lemma CustRoundTrip(c: Cust)
    requires WfCust(c)
    ensures decodeCust(encodeCust(c)) == c {}

  lemma CustsRoundTrip(cs: seq<Cust>)
    requires WfCusts(cs)
    ensures decodeCusts(encodeCusts(cs)) == cs
    decreases cs
  {
    if |cs| == 0 {} else { CustRoundTrip(cs[0]); CustsRoundTrip(cs[1..]); }
  }

  lemma ChatRoundTrip(c: Chat)
    ensures decodeChat(encodeChat(c)) == c {}

  lemma ClientRoundTrip(c: Client)
    requires F.asStr(F.field(c.raw, "clientId")) == c.clientId
    ensures decodeClient(encodeClient(c)) == c {}

  lemma InputRoundTrip(r: InputReq)
    requires F.asStr(F.field(r.raw, "id")) == r.id
    ensures decodeInputReq(encodeInputReq(r)) == r {}

  // ══════════════════════════════════════════════════════════════════════════
  //  THE LOAD-BEARING TRANSPORT PROPERTY: every well-formed TYPED action encodes
  //  to a wire object that decodes back to itself. (SUnknown is excluded — its raw
  //  may itself be a typed-shaped object, so decode would re-classify it; that leaf
  //  is handled by the verbatim lemmas below, exactly as Fidelity does for RootUnknown.)
  // ══════════════════════════════════════════════════════════════════════════
  // Json→domain→Json fixpoint: the codec's canonical wire form re-decodes + re-encodes to ITSELF — the
  // direction the conformance suite exercises, now machine-checked. Follows from the direction-1 round-trip.
  lemma SessionWireCanonicalRoundTrip(a: SessionAction)
    requires !a.SUnknown?
    requires WfSessionAction(a)
    ensures EncodeSessionAction(DecodeSessionAction(EncodeSessionAction(a))) == EncodeSessionAction(a)
  { SessionActionRoundTrip(a); }
  lemma SessionActionRoundTrip(a: SessionAction)
    requires !a.SUnknown?
    requires WfSessionAction(a)
    ensures DecodeSessionAction(EncodeSessionAction(a)) == a
  {
    match a
    case Ready                          =>
    case CreationFailed(e)              =>
    case TitleChanged(t)                =>
    case ServerToolsChanged(tls)        =>
    case MetaChanged(m)                 =>
    case IsReadChanged(v)               =>
    case IsArchivedChanged(v)           =>
    case ActivityChanged(ac)            =>
    case ConfigChanged(cfg, repl)       =>
    case CustomizationsChanged(cs)      => CustsRoundTrip(cs);
    case CustomizationToggled(id, en)   =>
    case CustomizationRemoved(id)       =>
    case DefaultChatChanged(c)          =>
    case ChatAdded(sum)                 => ChatRoundTrip(sum);
    case ChatRemoved(r)                 =>
    case ChatUpdated(r, st, ac, md)     =>
    case ChangesetsChanged(cs)          =>
    case CanvasesChanged(cs)            =>
    case OpenCanvasesChanged(cs)        =>
    case ActiveClientSet(c)             => ClientRoundTrip(c);
    case ActiveClientRemoved(id)        =>
    case InputNeededSet(r)              => InputRoundTrip(r);
    case InputNeededRemoved(id)         =>
    case CustomizationUpdated(c)        => CustRoundTrip(c);
    case McpServerStateChanged(id, st, ch) =>
    case McpServerStartRequested(id)    =>
    case McpServerStopRequested(id)     =>
    case SUnknown(raw)                  =>
  }

  // Non-vacuity witness: the round-trip precondition is inhabited by a GENUINE typed
  // action carrying a real Cust in a seq (exercises the inductive CustsRoundTrip path),
  // so SessionActionRoundTrip is NOT vacuously true and NOT constants-only.
  lemma SessionActionRoundTrip_NonVacuous()
    ensures exists a: SessionAction :: !a.SUnknown? && WfSessionAction(a) && DecodeSessionAction(EncodeSessionAction(a)) == a
  {
    var c := Cust("mcp-1", "mcpServer", "file:///w", "MCP", Some(true), None, None);
    var a := CustomizationsChanged([c]);
    assert WfCust(c);
    assert WfCusts([c]);
    SessionActionRoundTrip(a);
    assert !a.SUnknown? && WfSessionAction(a) && DecodeSessionAction(EncodeSessionAction(a)) == a;
  }

  // ── unknown-variant passthrough — the SUnknown leaf re-encodes verbatim (honest
  //    carve-out; the raw payload may re-classify on decode, so it is NOT in the
  //    typed round-trip above). Mirrors Fidelity's RootUnknown_EncodeVerbatim. ──
  lemma SUnknown_EncodeVerbatim(j: Wire.Json)
    ensures EncodeSessionAction(SUnknown(j)) == j {}
  lemma SUnknownVariant_Verbatim(j: Wire.Json)
    requires DecodeSessionAction(j).SUnknown?
    ensures EncodeSessionAction(DecodeSessionAction(j)) == j {}

  // ══════════════════════════════════════════════════════════════════════════
  //  STATE serialization — deterministic TOTAL serializer of SessionState (the
  //  analogue of Fidelity.encodeRootState; used by a cross-process transport demo).
  //  Convergence relies on determinism, not a round-trip lemma.
  // ══════════════════════════════════════════════════════════════════════════
  function encodeSConfig(c: SConfig): Wire.Json {
    Wire.JObj(map["schema" := c.schema, "values" := Wire.JObj(c.values)])
  }
  function encodeSessionState(s: SessionState): Wire.Json {
    Wire.JObj(map[
      "provider" := Wire.JStr(s.provider),
      "title" := Wire.JStr(s.title),
      "status" := Wire.JNum(s.status as int),
      "lifecycle" := Wire.JStr(s.lifecycle),
      "activity" := (if s.activity.Some? then Wire.JStr(s.activity.value) else Wire.JNull),
      "config" := (if s.config.Some? then encodeSConfig(s.config.value) else Wire.JNull),
      "_meta" := (if s.meta.Some? then s.meta.value else Wire.JNull),
      "creationError" := (if s.creationError.Some? then s.creationError.value else Wire.JNull),
      "serverTools" := (if s.serverTools.Some? then s.serverTools.value else Wire.JNull),
      "changesets" := (if s.changesets.Some? then s.changesets.value else Wire.JNull),
      "canvases" := (if s.canvases.Some? then s.canvases.value else Wire.JNull),
      "openCanvases" := (if s.openCanvases.Some? then s.openCanvases.value else Wire.JNull),
      "defaultChat" := (if s.defaultChat.Some? then Wire.JStr(s.defaultChat.value) else Wire.JNull),
      "activeClients" := Wire.JArr(encodeClients(s.activeClients)),
      "chats" := Wire.JArr(encodeChats(s.chats)),
      "customizations" := Wire.JArr(encodeCusts(s.customizations)),
      "inputNeeded" := Wire.JArr(encodeInputs(s.inputNeeded))])
  }

  // ══════════════════════════════════════════════════════════════════════════
  //  STATE decode — the true INVERSE of encodeSessionState (PROMOTED from
  //  spec/replay/session_replay.dfy's proven `decodeSessionState`, the exact
  //  logic that drives the 61 real session-reducer fixtures). Shared accessors
  //  sourced from Fidelity (F.field / F.asStr / F.asInt / F.asArr / F.asObjMap).
  //  This completes the STATE codec's proof story in both directions: the ACTION
  //  codec already round-trips (SessionActionRoundTrip), and now so does STATE.
  // ══════════════════════════════════════════════════════════════════════════

  // status is a bv32 bitset in the model; the wire carries it as a small int.
  function asStatus(o: Option<Wire.Json>): bv32 {
    var n := asInt(o); if 0 <= n < 0x1_0000_0000 then n as bv32 else 0
  }

  function decodeChats(js: seq<Wire.Json>): seq<Chat>
  { CO.Map(decodeChat, js) }

  function decodeClients(js: seq<Wire.Json>): seq<Client>
  { CO.Map(decodeClient, js) }

  function decodeInputs(js: seq<Wire.Json>): seq<InputReq>
  { CO.Map(decodeInputReq, js) }

  // Option<SConfig> from a state `config` value: present Wire.JObj → Some(SConfig);
  // absent / non-object (encode emits Wire.JNull for None) → None.
  function decodeConfigOpt(o: Option<Wire.Json>): Option<SConfig> {
    if o.Some? && o.value.JObj? then
      Some(SConfig(
        if F.field(o.value, "schema").Some? then F.field(o.value, "schema").value else Wire.JNull,
        F.asObjMap(F.field(o.value, "values"))))
    else None
  }

  // The full inverse. Reads each field encodeSessionState emits, reconstructs the
  // SessionState. The model field `meta` is carried in the wire key `_meta`.
  function decodeSessionState(j: Wire.Json): SessionState {
    SessionState(
      F.asStr(F.field(j, "provider")),
      F.asStr(F.field(j, "title")),
      asStatus(F.field(j, "status")),
      F.asStr(F.field(j, "lifecycle")),
      optStr(F.field(j, "activity")),
      decodeConfigOpt(F.field(j, "config")),
      optJsonNoNull(F.field(j, "_meta")),
      optJsonNoNull(F.field(j, "creationError")),
      optJsonNoNull(F.field(j, "serverTools")),
      optJsonNoNull(F.field(j, "changesets")),
      optJsonNoNull(F.field(j, "canvases")),
      optJsonNoNull(F.field(j, "openCanvases")),
      optStr(F.field(j, "defaultChat")),
      decodeClients(F.asArr(F.field(j, "activeClients"))),
      decodeChats(F.asArr(F.field(j, "chats"))),
      decodeCusts(F.asArr(F.field(j, "customizations"))),
      decodeInputs(F.asArr(F.field(j, "inputNeeded"))))
  }

  // ── STATE well-formedness — exactly the image of `decodeSessionState`, the same
  //    honest carve-out shape as WfSessionAction / Fidelity's RootUnknown. Four
  //    Option<Wire.Json> fields (meta/creationError/serverTools/changesets) decode
  //    through optJsonNoNull, which collapses BOTH absent and JSON `null` to None;
  //    so a `Some(Wire.JNull)` is not in decode's image and cannot round-trip. Client /
  //    InputReq store a redundant identity key ALONGSIDE the full raw Wire.Json; encode
  //    re-emits the raw verbatim, so the round-trip holds iff the stored key matches
  //    the one embedded in raw — precisely what decode guarantees by construction.
  //    Every OTHER field round-trips unconditionally (Wire.JStr/Wire.JNull-disjoint optionals,
  //    bv32↔int, Wire.JObj config, and the Chat/Cust sub-codecs). ──
  predicate WfClients(cl: seq<Client>) {
    forall i :: 0 <= i < |cl| ==> F.asStr(F.field(cl[i].raw, "clientId")) == cl[i].clientId
  }
  predicate WfInputs(ins: seq<InputReq>) {
    forall i :: 0 <= i < |ins| ==> F.asStr(F.field(ins[i].raw, "id")) == ins[i].id
  }
  // The WIRE-FIDELITY facet: exactly the well-formedness the codec round-trip needs (okJson on the
  // opaque-Json fields + per-element well-formedness of the keyed seqs). It deliberately EXCLUDES the
  // rung-7 UniqueKeys conjuncts, whose quantifiers are irrelevant to encode/decode fidelity and, if
  // dragged into the round-trip VCs, blow SessionStateRoundTrip past its resource limit.
  predicate WfSessionStateWire(s: SessionState) {
    (s.meta.Some?          ==> s.meta.value          != Wire.JNull) &&
    (s.creationError.Some? ==> s.creationError.value != Wire.JNull) &&
    (s.serverTools.Some?   ==> s.serverTools.value   != Wire.JNull) &&
    (s.changesets.Some?    ==> s.changesets.value    != Wire.JNull) &&
    (s.canvases.Some?      ==> s.canvases.value      != Wire.JNull) &&
    (s.openCanvases.Some?  ==> s.openCanvases.value  != Wire.JNull) &&
    WfCusts(s.customizations) &&
    WfClients(s.activeClients) &&
    WfInputs(s.inputNeeded)
  }

  predicate WfSessionState(s: SessionState) {
    WfSessionStateWire(s) &&
    // rung-7 (docs/shapes/07): the four id-keyed collections hold DISTINCT keys — the keyed-collection
    // uniqueness invariant, keyed by the SAME Session.custKey/chatKey/clientKey/inputKey the reducer routes
    // through, so Session.SessionWfPreserved carries these exact facts across every reducer arm.
    SR.UniqueKeys(custKey,   s.customizations) &&
    SR.UniqueKeys(chatKey,   s.chats) &&
    SR.UniqueKeys(clientKey, s.activeClients) &&
    SR.UniqueKeys(inputKey,  s.inputNeeded)
  }

  // ── sub-codec round-trip proofs for the seq fields (mirror CustsRoundTrip) ──
  lemma ChatsRoundTrip(chs: seq<Chat>)
    ensures decodeChats(encodeChats(chs)) == chs
    decreases chs
  { if |chs| == 0 {} else { ChatRoundTrip(chs[0]); ChatsRoundTrip(chs[1..]); } }

  lemma ClientsRoundTrip(cl: seq<Client>)
    requires WfClients(cl)
    ensures decodeClients(encodeClients(cl)) == cl
    decreases cl
  { if |cl| == 0 {} else { ClientRoundTrip(cl[0]); ClientsRoundTrip(cl[1..]); } }

  lemma InputsRoundTrip(ins: seq<InputReq>)
    requires WfInputs(ins)
    ensures decodeInputs(encodeInputs(ins)) == ins
    decreases ins
  { if |ins| == 0 {} else { InputRoundTrip(ins[0]); InputsRoundTrip(ins[1..]); } }

  // ══════════════════════════════════════════════════════════════════════════
  //  THE LOAD-BEARING STATE TRANSPORT PROPERTY: every well-formed SessionState
  //  encodes to a wire object that decodes back to itself. REAL, quantified,
  //  non-vacuous — the STATE analogue of SessionActionRoundTrip / RootActionRoundTrip.
  // ══════════════════════════════════════════════════════════════════════════
  lemma {:isolate_assertions} SessionStateRoundTrip(s: SessionState)
    requires WfSessionStateWire(s)
    ensures decodeSessionState(encodeSessionState(s)) == s
  {
    AhpSkeleton.StatusRoundTrip(s.status);   // (s.status as int) as bv32 == s.status
    ClientsRoundTrip(s.activeClients);
    ChatsRoundTrip(s.chats);
    CustsRoundTrip(s.customizations);
    InputsRoundTrip(s.inputNeeded);

    // Keep the final datatype equality cheap and deterministic for the verifier.
    // Asking Z3 to normalize all fifteen map lookups, option decoders and four
    // sequence codecs in one postcondition query is needlessly monolithic.  The
    // isolated projection facts below expose the same proof one field at a time;
    // datatype extensionality then closes the original, unchanged contract.
    var decoded := decodeSessionState(encodeSessionState(s));
    assert decoded.provider == s.provider;
    assert decoded.title == s.title;
    assert decoded.status == s.status;
    assert decoded.lifecycle == s.lifecycle;
    assert decoded.activity == s.activity;
    assert decoded.config == s.config;
    assert decoded.meta == s.meta;
    assert decoded.creationError == s.creationError;
    assert decoded.serverTools == s.serverTools;
    assert decoded.changesets == s.changesets;
    assert decoded.canvases == s.canvases;
    assert decoded.openCanvases == s.openCanvases;
    assert decoded.defaultChat == s.defaultChat;
    assert decoded.activeClients == s.activeClients;
    assert decoded.chats == s.chats;
    assert decoded.customizations == s.customizations;
    assert decoded.inputNeeded == s.inputNeeded;
    assert decoded == s;
  }

  // Non-vacuity witness: a GENUINE non-empty state (real Cust in a seq, a Chat,
  // a Some config, a Some activity/defaultChat) satisfies the precondition and
  // round-trips — so SessionStateRoundTrip is NOT vacuous and NOT constants-only.
  lemma SessionStateRoundTrip_NonVacuous()
    ensures exists s: SessionState :: WfSessionState(s) && decodeSessionState(encodeSessionState(s)) == s
  {
    var c := Cust("mcp-1", "mcpServer", "file:///w", "MCP", Some(true), None, None);
    var s := SessionState(
      "prov", "Title", 72, "active",
      Some("thinking"),
      Some(SConfig(Wire.JNull, map["k" := Wire.JStr("v")])),
      None, None, None, None, None, None,
      Some("chat-1"),
      [], [Chat("res", "t", 1, Some("a"), "1970-01-01T00:00:02.000Z", Wire.JNull)], [c], []);
    assert WfCust(c);
    assert WfCusts([c]);
    SessionStateRoundTrip(s);
    assert WfSessionState(s) && decodeSessionState(encodeSessionState(s)) == s;
  }

  ghost predicate SessionCodecWf(a: SessionAction) { !a.SUnknown? && WfSessionAction(a) }
  function SessionCodecC(): Wire.Codec<SessionAction> {
    Wire.Codec(EncodeSessionAction, (j: Wire.Json) => Wire.Some(DecodeSessionAction(j)))
  }
  lemma SessionCodecRoundTripsWhere()
    ensures forall a :: SessionCodecWf(a) ==> SessionCodecC().decode(SessionCodecC().encode(a)) == Wire.Some(a)
  {
    forall a | SessionCodecWf(a)
      ensures SessionCodecC().decode(SessionCodecC().encode(a)) == Wire.Some(a)
    { SessionActionRoundTrip(a); }
  }
}
