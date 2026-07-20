// AHP Dafny client — fidelity proofs (Phase 3 start; promoted from the vaulted spikes).
// Over the client's real `AhpSkeleton.Json`. These are the wire-fidelity heart:
// bv32 bitset (fixtures 004/005), open/closed Unknown-preservation (017/019),
// unknown-variant passthrough (002/024).

include "ahp_skeleton.dfy"

module Fidelity {
  export
    provides AhpSkeleton, Wire, CO
    provides DecodeRootAction, encodeRootState, decodeRootState
    provides decodeAgent, encodeAgent, decodeAgents, encodeAgents
    provides StrsRoundTrip, AgentRoundTrip, AgentsRoundTrip, RootActionRoundTrip, RootStateRoundTrip
    provides RootUnknown_EncodeVerbatim, UnknownVariant_Verbatim
    reveals field, asStr, asBool, asArr, asObjMap, strList, encodeStrs
    reveals EncodeRootAction

  import opened AhpSkeleton
  import CO = ConfluxOperators
  // Wire.Json now comes from ConfluxCodec, not AhpSkeleton. This file's OWN field/
  // asStr/asInt/asBool/asArr (below) SHADOW ConfluxCodec's same-named accessors
  // under the bare name (Dafny: a local declaration shadows an opened-import
  // name of the same name — confirmed empirically; no duplicate-decl error).
  // Every bare call to field/asStr/... in this file keeps resolving to THIS
  // file's own definitions, unchanged. Only Wire.Json/Wire.JNull/Wire.JBool/Wire.JNum/Wire.JDec/Wire.JStr/
  // Wire.JArr/Wire.JObj come from the import.
  import Wire = ConfluxCodec

  // ── bv32 bitset (fixtures 004/005) — the resolved formulation (spike-proven) ──
  const WAITING : bv32 := 0x00000008
  const ERROR   : bv32 := 0x00000040
  const IDLE    : bv32 := 0x00000001
  const UNKNOWN_HI : bv32 := 0x80000000
  const COMBO_72  : bv32 := WAITING | ERROR
  const COMBO_UNK : bv32 := UNKNOWN_HI | WAITING | ERROR

  // forward round-trip over ALL bv32 (the law that matters; the scoped reverse per the anti-pattern).
  lemma Status_RoundTrip(s: SessionStatus) ensures (EncodeStatus(s) as bv32) == s {}
  lemma Status_WidthUniformity(s: SessionStatus) ensures 0 <= EncodeStatus(s) <= 0xFFFF_FFFF {}
  // fixture 004: flag-combo decode
  lemma Status_FlagCombo72()
    ensures (72 as bv32) == COMBO_72
    ensures COMBO_72 & WAITING == WAITING && COMBO_72 & ERROR == ERROR && COMBO_72 & IDLE == 0 {}
  // fixture 005: unknown high bit preserved (the property TS's closed const-enum cannot represent)
  lemma Status_UnknownBitPreserved()
    ensures EncodeStatus(COMBO_UNK) == 2147483720
    ensures (2147483720 as bv32) == COMBO_UNK
    ensures COMBO_UNK & UNKNOWN_HI == UNKNOWN_HI {}

  // ── open/closed decode dichotomy (fixtures 017/019) over the real Wire.JObj(map) ──
  predicate WellFormedOpen(m: map<string, Wire.Json>) { "known" in m && m["known"].JStr? }

  // OPEN: keep unknown keys → preserve verbatim.
  datatype OpenRec = OpenRec(known: string, extra: map<string, Wire.Json>)
  function DecodeOpen(m: map<string, Wire.Json>): OpenRec requires WellFormedOpen(m)
  { OpenRec(m["known"].s, m - {"known"}) }
  function EncodeOpen(r: OpenRec): map<string, Wire.Json> { r.extra["known" := Wire.JStr(r.known)] }
  lemma Open_PreservesUnknownKeys(m: map<string, Wire.Json>)
    requires WellFormedOpen(m)
    ensures EncodeOpen(DecodeOpen(m)) == m
  { assert m["known"] == Wire.JStr(m["known"].s); }

  // CLOSED: drop unknown keys → re-encode has only the known key.
  datatype ClosedRec = ClosedRec(known: string)
  function DecodeClosed(m: map<string, Wire.Json>): ClosedRec requires WellFormedOpen(m)
  { ClosedRec(m["known"].s) }
  function EncodeClosed(r: ClosedRec): map<string, Wire.Json> { map["known" := Wire.JStr(r.known)] }
  lemma Closed_DropsUnknownKeys(m: map<string, Wire.Json>)
    requires WellFormedOpen(m)
    ensures EncodeClosed(DecodeClosed(m)).Keys == {"known"}
    ensures EncodeClosed(DecodeClosed(m))["known"] == m["known"] {}

  // ════════════════════════════════════════════════════════════════════════
  //  FULL Wire.Json ⇄ RootAction codec (promoted from spec/replay/root_replay.dfy).
  //  DECODE is the exact logic that drives 7/7 root fixtures in the harness;
  //  ENCODE is its inverse (canonical wire object per typed variant). Together
  //  they carry the load-bearing transport property: decode ∘ encode == id on
  //  the typed variants (RootActionRoundTrip below).
  // ════════════════════════════════════════════════════════════════════════

  // ── field accessors on skeleton Wire.Json (pure; identical to the harness) ──
  function field(j: Wire.Json, k: string): Option<Wire.Json> {
    if j.JObj? && k in j.fields then Some(j.fields[k]) else None
  }
  function asStr(o: Option<Wire.Json>): string { if o.Some? && o.value.JStr? then o.value.s else "" }
  // asInt / optInt / jsonInt (the wire-integer accessor family — JNum + integer-valued JDec, closing the
  // JNum-only silent-drop) live in AhpSkeleton now, the ONE home shared with the reducer-replay harnesses.
  // In scope here via `import opened AhpSkeleton`; call unqualified (`asInt(...)`), NOT `F.asInt`.
  function asBool(o: Option<Wire.Json>, dflt: bool): bool { if o.Some? && o.value.JBool? then o.value.b else dflt }
  function asArr(o: Option<Wire.Json>): seq<Wire.Json> { if o.Some? && o.value.JArr? then o.value.elems else [] }
  function asObjMap(o: Option<Wire.Json>): map<string, Wire.Json> { if o.Some? && o.value.JObj? then o.value.fields else map[] }

  function strList(js: seq<Wire.Json>): seq<string>
  { CO.Map((j: Wire.Json) => if j.JStr? then j.s else "", js) }

  // ── DECODE domain values (promoted verbatim from the harness) ──
  function decodeAgent(j: Wire.Json): AgentInfo {
    AgentInfo(asStr(field(j, "provider")), asStr(field(j, "displayName")),
              asStr(field(j, "description")), strList(asArr(field(j, "models"))))
  }
  function decodeAgents(js: seq<Wire.Json>): seq<AgentInfo>
  { CO.Map(decodeAgent, js) }

  // ── ENCODE domain values (the inverse of the decode above) ──
  function encodeStrs(ss: seq<string>): seq<Wire.Json>
  { CO.Map((s: string) => Wire.JStr(s), ss) }
  function encodeAgent(a: AgentInfo): Wire.Json {
    Wire.JObj(map["provider" := Wire.JStr(a.provider),
             "displayName" := Wire.JStr(a.displayName),
             "description" := Wire.JStr(a.description),
             "models" := Wire.JArr(encodeStrs(a.models))])
  }
  function encodeAgents(ags: seq<AgentInfo>): seq<Wire.Json>
  { CO.Map(encodeAgent, ags) }

  // ── the typed codec ──
  function DecodeRootAction(j: Wire.Json): RootAction {
    var t := asStr(field(j, "type"));
    if t == "root/agentsChanged" then RootAgentsChanged(decodeAgents(asArr(field(j, "agents"))))
    else if t == "root/activeSessionsChanged" then RootActiveSessionsChanged(asInt(field(j, "activeSessions")))
    else if t == "root/terminalsChanged" then RootTerminalsChanged(asArr(field(j, "terminals")))
    else if t == "root/configChanged" then RootConfigChanged(asObjMap(field(j, "config")), asBool(field(j, "replace"), false))
    else RootUnknown(j)
  }
  function EncodeRootAction(a: RootAction): Wire.Json {
    match a
    case RootAgentsChanged(ags)       => Wire.JObj(map["type" := Wire.JStr("root/agentsChanged"), "agents" := Wire.JArr(encodeAgents(ags))])
    case RootActiveSessionsChanged(n) => Wire.JObj(map["type" := Wire.JStr("root/activeSessionsChanged"), "activeSessions" := Wire.JNum(n)])
    case RootTerminalsChanged(t)      => Wire.JObj(map["type" := Wire.JStr("root/terminalsChanged"), "terminals" := Wire.JArr(t)])
    case RootConfigChanged(cfg, repl) => Wire.JObj(map["type" := Wire.JStr("root/configChanged"), "config" := Wire.JObj(cfg), "replace" := Wire.JBool(repl)])
    case RootUnknown(raw)             => raw
  }

  // ── ROUND-TRIP proofs: decode ∘ encode == id on the sub-codecs ──
  lemma StrsRoundTrip(ss: seq<string>)
    ensures strList(encodeStrs(ss)) == ss
    decreases ss
  { if |ss| == 0 {} else { StrsRoundTrip(ss[1..]); } }

  lemma AgentRoundTrip(a: AgentInfo)
    ensures decodeAgent(encodeAgent(a)) == a
  { StrsRoundTrip(a.models); }

  lemma AgentsRoundTrip(ags: seq<AgentInfo>)
    ensures decodeAgents(encodeAgents(ags)) == ags
    decreases ags
  {
    if |ags| == 0 {} else { AgentRoundTrip(ags[0]); AgentsRoundTrip(ags[1..]); }
  }

  // The load-bearing transport property: every TYPED action encodes to a wire
  // object that decodes back to itself. (RootUnknown is excluded — its raw
  // payload may itself be a typed-shaped object, so decode would re-classify
  // it; that leaf is handled by the verbatim lemmas below, not this one.)
  // Json→domain→Json fixpoint: the codec's canonical wire form re-decodes + re-encodes to ITSELF — the
  // direction the conformance suite exercises, now machine-checked. Follows from the direction-1 round-trip.
  lemma RootWireCanonicalRoundTrip(a: RootAction)
    requires !a.RootUnknown?
    ensures EncodeRootAction(DecodeRootAction(EncodeRootAction(a))) == EncodeRootAction(a)
  { RootActionRoundTrip(a); }
  lemma RootActionRoundTrip(a: RootAction)
    requires !a.RootUnknown?
    ensures DecodeRootAction(EncodeRootAction(a)) == a
  {
    match a
    case RootAgentsChanged(ags)       => AgentsRoundTrip(ags);
    case RootActiveSessionsChanged(n) =>
    case RootTerminalsChanged(t)      =>
    case RootConfigChanged(cfg, repl) =>
  }

  // Non-vacuity witness: the round-trip precondition is inhabited (a genuine
  // typed action round-trips), so RootActionRoundTrip is NOT vacuously true.
  lemma RootActionRoundTrip_NonVacuous()
    ensures exists a: RootAction :: !a.RootUnknown? && DecodeRootAction(EncodeRootAction(a)) == a
  {
    var a := RootActiveSessionsChanged(5);
    RootActionRoundTrip(a);
    assert !a.RootUnknown? && DecodeRootAction(EncodeRootAction(a)) == a;
  }

  // ── unknown-variant passthrough (fixtures 002/024) — the RootUnknown leaf re-encodes verbatim ──
  lemma RootUnknown_EncodeVerbatim(j: Wire.Json)
    ensures EncodeRootAction(RootUnknown(j)) == j {}
  lemma UnknownVariant_Verbatim(j: Wire.Json)
    requires DecodeRootAction(j).RootUnknown?
    ensures EncodeRootAction(DecodeRootAction(j)) == j {}

  // ── STATE serialization (for the runnable cross-process transport demo) ──
  // encodeRootState serializes a RootState to canonical JSON so the host and the
  // client can each emit their final state and a cross-process e2e can compare
  // them. It is a DETERMINISTIC TOTAL function; the transport demo's convergence
  // check relies on that determinism + the convergence PROOF (mirror == canonical
  // ⇒ identical serialization), NOT on a round-trip lemma. (A decodeRootState +
  // RootState round-trip proof is a natural future hardening, not needed here.)
  function encodeConfig(c: RootConfig): Wire.Json {
    Wire.JObj(map["schema" := c.schema, "values" := Wire.JObj(c.values)])
  }
  function encodeRootState(s: RootState): Wire.Json {
    var m0 := map["agents" := Wire.JArr(encodeAgents(s.agents))];
    var m1 := if s.activeSessions.Some? then m0["activeSessions" := Wire.JNum(s.activeSessions.value)] else m0;
    var m2 := if s.terminals.Some? then m1["terminals" := Wire.JArr(s.terminals.value)] else m1;
    var m3 := if s.config.Some? then m2["config" := encodeConfig(s.config.value)] else m2;
    Wire.JObj(m3)
  }

  // ════════════════════════════════════════════════════════════════════════
  //  STATE codec inverse: decodeRootState is the TRUE inverse of encodeRootState.
  //  DECODE is the exact logic that drives the root reducer-replay harness
  //  (spec/replay/root_replay.dfy); promoted here so the STATE codec, like the
  //  ACTION codec above, carries a real round-trip proof — decode ∘ encode == id.
  //
  //  Why NO carve-out is needed (contrast the ACTION codec's `!a.RootUnknown?`):
  //  every field encodeRootState emits round-trips verbatim.
  //    • `agents`         — ALWAYS emitted; inverts via AgentsRoundTrip.
  //    • `activeSessions`  — Option<int>: Some(n) ⇒ key present as Wire.JNum(n) (a
  //         form distinguishable from absence); None ⇒ key absent. No value of
  //         the Option collapses onto the wire-form of another, so decode inverts.
  //    • `terminals`       — Option<seq<Wire.Json>>: Some(ts) ⇒ key present as Wire.JArr(ts);
  //         None ⇒ key absent. Opaque seq stored + read back verbatim.
  //    • `config`          — Option<RootConfig>: Some(c) ⇒ key present as a Wire.JObj
  //         (distinguishable from absence); None ⇒ key absent. encodeConfig ALWAYS
  //         writes both `schema` (a plain Wire.Json, verbatim — so even a Wire.JNull schema
  //         is a PRESENT key, not a collapsed absence) and `values`, so decodeConfig
  //         reconstructs the RootConfig exactly. The Option-Wire.JNull collapse that
  //         forces a Wf* precondition in a lossy codec never arises here because
  //         presence is signalled by the KEY, and schema is not an Option<Wire.Json>.
  // ════════════════════════════════════════════════════════════════════════
  function decodeConfig(o: Option<Wire.Json>): Option<RootConfig> {
    if o.None? || !o.value.JObj? then None
    else Some(RootConfig(
      if field(o.value, "schema").Some? then field(o.value, "schema").value else Wire.JNull,
      asObjMap(field(o.value, "values"))))
  }
  function decodeRootState(j: Wire.Json): RootState {
    RootState(
      decodeAgents(asArr(field(j, "agents"))),
      (var av := field(j, "activeSessions"); if av.Some? then jsonInt(av.value) else None),  // JNum OR integer-valued JDec; fractional/non-number ⇒ None (was: JNum-only silent drop)
      if field(j, "terminals").Some? && field(j, "terminals").value.JArr?
        then Some(asArr(field(j, "terminals"))) else None,
      decodeConfig(field(j, "config")))
  }

  // The load-bearing STATE property: every RootState serializes to a wire object
  // that decodes back to itself — quantified over ALL RootState, unconditionally.
  lemma RootStateRoundTrip(s: RootState)
    ensures decodeRootState(encodeRootState(s)) == s
  {
    AgentsRoundTrip(s.agents);
    // The encoded map (from encodeRootState) with its four DISTINCT string keys.
    var m0 := map["agents" := Wire.JArr(encodeAgents(s.agents))];
    var m1 := if s.activeSessions.Some? then m0["activeSessions" := Wire.JNum(s.activeSessions.value)] else m0;
    var m2 := if s.terminals.Some? then m1["terminals" := Wire.JArr(s.terminals.value)] else m1;
    var m3 := if s.config.Some? then m2["config" := encodeConfig(s.config.value)] else m2;
    // `agents` is written first and never overwritten (later keys are distinct),
    // so it survives every conditional update below.
    assert "agents" in m3 && m3["agents"] == Wire.JArr(encodeAgents(s.agents));
    // `config` sub-fields round-trip: encodeConfig always writes both keys.
    if s.config.Some? {
      assert m3["config"] == encodeConfig(s.config.value);
      assert field(encodeConfig(s.config.value), "schema") == Some(s.config.value.schema);
      assert field(encodeConfig(s.config.value), "values") == Some(Wire.JObj(s.config.value.values));
    }
  }

  // Non-vacuity witness: a genuinely non-empty state (all four fields populated,
  // agents non-empty) round-trips — so RootStateRoundTrip is NOT vacuously true.
  lemma RootStateRoundTrip_NonVacuous()
    ensures exists s: RootState ::
      s.agents != [] && decodeRootState(encodeRootState(s)) == s
  {
    var s := RootState(
      [AgentInfo("copilot", "Copilot", "AI", ["gpt-4"])],
      Some(5),
      Some([Wire.JStr("term-1")]),
      Some(RootConfig(Wire.JNull, map["theme" := Wire.JStr("dark")])));
    RootStateRoundTrip(s);
    assert s.agents != [] && decodeRootState(encodeRootState(s)) == s;
  }

  ghost predicate RootCodecWf(a: RootAction) { !a.RootUnknown? }
  function RootCodecC(): Wire.Codec<RootAction> {
    Wire.Codec(EncodeRootAction, (j: Wire.Json) => Wire.Some(DecodeRootAction(j)))
  }
  lemma RootCodecRoundTripsWhere()
    ensures forall a :: RootCodecWf(a) ==> RootCodecC().decode(RootCodecC().encode(a)) == Wire.Some(a)
  {
    forall a | RootCodecWf(a)
      ensures RootCodecC().decode(RootCodecC().encode(a)) == Wire.Some(a)
    { RootActionRoundTrip(a); }
  }
}
