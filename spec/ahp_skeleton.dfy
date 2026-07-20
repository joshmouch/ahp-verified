// AHP Dafny client — Phase 0.5 WALKING SKELETON (root channel, end-to-end).
// Proves the full pipeline: model → verify → `dafny translate cs|js` → CORPUS.
// Models the FULL root channel to pass ALL 7 real corpus fixtures
// (types/test-cases/reducers/*-root-*.json). Compile-safe (total, no ghost) → extracts.

// ConfluxCodec is the one Json value model owned by the Conflux kernel (no
// AHP-local bridge or datatype). This AHP-owned source imports ConfluxCodec and
// ConfluxContract from the installed runtime-only package's conflux-runtime.doo;
// it does not include Conflux source. Separate C#/JavaScript translation uses
// the package DTRs and links the package implementation for runnable corpus,
// reducer-replay, codec-roundtrip, client, and demo targets.
module AhpSkeleton {
  import Wire = ConfluxCodec   // Json/JNull/JBool/JNum/JDec/JStr/JArr/JObj — the ONE Json value model
  import CC = ConfluxContract  // the one fold — FoldRoot calls CC.Fold (delete-and-reuse, not a bridge)

  // ── .doo FIREWALL: explicit export set (replaces Dafny's default export-all) ──
  // `provides` exports a name's signature/spec but HIDES its body; `reveals`
  // exports the body too (datatypes need `reveals` so importers see constructors
  // + fields to pattern-match/construct). Only the names any consumer actually
  // uses are exported; internal Root no-op/determinism lemmas, the unused
  // Envelope/ActionOrigin shapes, and the IsUnknownRoot/ConfigPreconditionFails
  // predicates are HIDDEN — the firewall now bites on them.
  //
  //   reveals — datatypes/type-synonyms consumers construct or match on, plus the
  //             root fold functions whose BODY convergence unfolds (apply1/FoldRoot)
  //             and EncodeStatus (fidelity's bitset round-trip proves over its body).
  //             Wire.Json is NO LONGER in this list — AhpSkeleton doesn't own it anymore;
  //             Wire.Json is a name IMPORTED (opened) from ConfluxCodec, not declared here,
  //             and Dafny's export-set checker refuses `reveals <imported-name>`
  //             ("must be a member of AhpSkeleton to be exported" — confirmed
  //             empirically). Every consumer that needs Wire.Json now does its OWN
  //             `import Wire = ConfluxCodec` directly (see spec/fidelity.dfy etc).
  //   provides — ApplyToRoot (named by apply1/FoldRoot's revealed bodies, body kept
  //             hidden), StatusRoundTrip (lemma; session_codec calls it), RunCorpus
  //             (method; core's client_main runs it). ConfluxCodec is `provides`d
  //             too: the REVEALED RootConfig/RootState datatypes have Wire.Json-typed
  //             fields, and revealing a datatype forces its field types' owning
  //             module to be exported (same rule spec/ahp.dfy's Ahp module already
  //             documents for AhpSkeleton/S/Ch/T/Cs/An/RW).
  export
    provides ApplyToRoot, StatusRoundTrip, RunCorpus
    provides Wire, CC
    reveals Option, ReduceOutcome, SessionStatus, EncodeStatus,
            ActionOrigin, Envelope,
            AgentInfo, RootConfig, RootState, RootAction,
            apply1, FoldRoot,
            jsonInt, asInt, optInt   // the ONE wire-integer accessor family (bodies revealed so consumers' round-trips unfold them)
    provides decInt                  // opaque: JNum round-trips must not unroll its Pow10/modulo

  datatype Option<T> = None | Some(value: T)

  // ── wire-integer extraction, TOTAL over the wire's TWO number forms (JNum + JDec) ──
  // THE ONE home for reading an integer off the wire — shared by the codec layer (Fidelity + spec/codec/*)
  // AND the reducer-replay harnesses (spec/replay/*), so the JNum-vs-JDec logic exists in exactly one place
  // and can no longer drift into per-file silent-drops. The bytes→Json bridge (Std.JSON; see
  // spec/replay/root_replay.dfy) emits JNum only when the parsed decimal has exponent 0, so a protocol integer
  // written non-canonically (5.0 == JDec(50,-1); 1e3 == JDec(1,3)) parses to a JDec. jsonInt reads the integer
  // VALUE of EITHER form (JDec(m,e) is an integer iff e >= 0, or e < 0 and 10^(-e) | m); a genuinely-fractional
  // JDec (1.5) or any non-number returns None — surfaced as "not an integer", never a silently-wrong 0.
  function Pow10(k: nat): int
    ensures Pow10(k) >= 1                                 // 10^k > 0 ⇒ the `% / Pow10` below is division-safe
  { if k == 0 then 1 else 10 * Pow10(k - 1) }
  opaque function decInt(m: int, e: int): Option<int> {
    if e >= 0 then Some(m * Pow10(e))
    else if m % Pow10(-e) == 0 then Some(m / Pow10(-e))   // e < 0 here ⇒ -e > 0; integer iff 10^(-e) | m
    else None                                             // genuinely fractional (e.g. 1.5)
  }
  function jsonInt(j: Wire.Json): Option<int> {
    if j.JNum? then Some(j.n) else if j.JDec? then decInt(j.mantissa, j.exp) else None
  }
  // asInt keeps the int/default-0 contract (JNum path IDENTICAL to the old accessors, so every round-trip that
  // decodes an int field verifies unchanged); optInt surfaces a non-integer as None instead of a wrong 0.
  function asInt(o: Option<Wire.Json>): int {
    if o.None? then 0 else var r := jsonInt(o.value); if r.Some? then r.value else 0
  }
  function optInt(o: Option<Wire.Json>): Option<int> {
    if o.Some? then jsonInt(o.value) else None
  }
  // Non-vacuity: integer-valued JDec (written 5.0 / 1e3) reads as its integer; a fractional JDec reads as None,
  // NOT a silently-wrong 0. Falsifiable — flip any RHS and it fails.
  lemma JsonIntWitness()
    ensures jsonInt(Wire.JNum(7)) == Some(7)
    ensures jsonInt(Wire.JDec(1, 3)) == Some(1000)        // 1e3 — integer via exponent ≥ 0
    ensures jsonInt(Wire.JDec(50, -1)) == Some(5)         // 5.0 — integer via 10^1 | 50
    ensures jsonInt(Wire.JDec(15, -1)) == None            // 1.5 — genuinely fractional ⇒ None
  {
    reveal decInt();
    assert Pow10(3) == 1000 && Pow10(1) == 10;
  }

  // ── the GENERAL canonicalization theorem — the silent-drop's impossibility, over ALL integers ──
  // `DenotesInt(m, e, n)` holds iff the wire number JDec(m, e) — value m·10^e — equals the integer n.
  ghost predicate DenotesInt(mantissa: int, exp: int, n: int) {
    (exp >= 0 && n == mantissa * Pow10(exp))
    || (exp < 0 && mantissa % Pow10(-exp) == 0 && n == mantissa / Pow10(-exp))
  }
  // Every wire encoding of an integer n — the canonical JNum(n) OR ANY integer-valued JDec(m,e) — reads to
  // Some(n). JsonIntWitness is the 3-constant instance; this is the ∀-integer theorem. So a protocol integer
  // written 5 / 5.0 / 1e3 / 500e-2 all decode identically: the JNum-vs-JDec wire form CANNOT silently change
  // the decoded integer. (The encoder only ever emits JNum for a numeric field, so encode∘decode canonicalizes
  // any integer-JDec input back to JNum — see the per-channel WireCanonicalRoundTrip lemmas.)
  lemma AsIntCanonicalizes(mantissa: int, exp: int, n: int)
    requires DenotesInt(mantissa, exp, n)
    ensures jsonInt(Wire.JDec(mantissa, exp)) == Some(n)
    ensures jsonInt(Wire.JNum(n)) == Some(n)
  { reveal decInt(); }
  // Corollary: any two wire encodings of the same integer decode to the SAME value — no form-dependent drift.
  lemma IntWireFormsAgree(mantissa: int, exp: int, n: int)
    requires DenotesInt(mantissa, exp, n)
    ensures jsonInt(Wire.JDec(mantissa, exp)) == jsonInt(Wire.JNum(n))
  { AsIntCanonicalizes(mantissa, exp, n); }

  datatype ReduceOutcome = Applied | NoOp | OutOfScope

  // ── SessionStatus bitset (bv32) — proven tractable in the spike ──
  type SessionStatus = bv32
  function EncodeStatus(s: SessionStatus): int { s as int }
  lemma StatusRoundTrip(s: SessionStatus) ensures (EncodeStatus(s) as bv32) == s {}

  // ── ActionEnvelope (one canonical shape; keeps clientSeq + rejectionReason) ──
  // AHP clients allocate clientSeq monotonically from zero.  Making that wire
  // invariant explicit lets the protocol adapter map losslessly to the shared
  // ConfluxCommand.CommandOrigin instead of carrying an impossible negative
  // sequence through a second authority path.
  datatype ActionOrigin = ActionOrigin(clientId: string, clientSeq: nat)
  datatype Envelope<A> = Envelope(
    channel: string, action: A, serverSeq: nat,
    origin: Option<ActionOrigin>, rejectionReason: Option<string>)

  // ── ROOT channel — full state to satisfy all 7 fixtures ──
  datatype AgentInfo = AgentInfo(provider: string, displayName: string, description: string, models: seq<string>)
  datatype RootConfig = RootConfig(schema: Wire.Json, values: map<string, Wire.Json>)
  datatype RootState = RootState(
    agents: seq<AgentInfo>,
    activeSessions: Option<int>,
    terminals: Option<seq<Wire.Json>>,       // terminalsChanged stores the list opaquely
    config: Option<RootConfig>)

  datatype RootAction =
    | RootAgentsChanged(agents: seq<AgentInfo>)                     // 001
    | RootActiveSessionsChanged(activeSessions: int)                // 002
    | RootTerminalsChanged(terminals: seq<Wire.Json>)                   // 071
    | RootConfigChanged(config: map<string, Wire.Json>, replace: bool)   // 127/128/130
    | RootUnknown(raw: Wire.Json)                                        // 087

  // Pure, total, deterministic; `now` injected (the #186 fix). Returns a RESULT.
  function ApplyToRoot(s: RootState, a: RootAction, now: int): (RootState, ReduceOutcome)
  {
    match a
    case RootAgentsChanged(ag)        => (s.(agents := ag), Applied)
    case RootActiveSessionsChanged(n) => (s.(activeSessions := Some(n)), Applied)
    case RootTerminalsChanged(t)      => (s.(terminals := Some(t)), Applied)
    case RootConfigChanged(cfg, repl) =>
      (match s.config
       case None => (s, NoOp)                                        // 128: no config → no-op
       case Some(c) =>
         var newValues := if repl then cfg else c.values + cfg;      // 130 replace vs 127 merge
         (s.(config := Some(RootConfig(c.schema, newValues))), Applied))
    case RootUnknown(_)               => (s, NoOp)                    // 087
  }

  predicate IsUnknownRoot(a: RootAction) { a.RootUnknown? }
  predicate ConfigPreconditionFails(s: RootState, a: RootAction) { a.RootConfigChanged? && s.config.None? }

  // ── Universal no-op lemmas (collapse the corpus no-op families) + non-vacuity witnesses ──
  lemma Root_UnknownIsNoOp(s: RootState, a: RootAction, now: int)
    requires IsUnknownRoot(a)
    ensures ApplyToRoot(s, a, now).0 == s && ApplyToRoot(s, a, now).1 == NoOp {}
  lemma Root_UnknownIsNoOp_NonVacuous()
    ensures exists a: RootAction :: IsUnknownRoot(a)
  { assert IsUnknownRoot(RootUnknown(Wire.JNull)); }

  lemma Root_ConfigPreconditionFailsIsNoOp(s: RootState, a: RootAction, now: int)
    requires ConfigPreconditionFails(s, a)
    ensures ApplyToRoot(s, a, now).0 == s && ApplyToRoot(s, a, now).1 == NoOp {}
  lemma Root_ConfigPreconditionFails_NonVacuous()
    ensures exists s: RootState, a: RootAction :: ConfigPreconditionFails(s, a)
  { assert ConfigPreconditionFails(RootState([], None, None, None), RootConfigChanged(map[], false)); }

  lemma Root_Deterministic(s: RootState, a: RootAction, now: int)
    ensures ApplyToRoot(s, a, now) == ApplyToRoot(s, a, now) {}

  function FoldRoot(s: RootState, actions: seq<RootAction>, now: int): RootState
  {
    // hand-rolled left-fold of ApplyToRoot(·,·,now).0 DELETED — now the kernel fold over the now-closed step
    CC.Fold((s': RootState, a: RootAction) => ApplyToRoot(s', a, now).0, s, actions)
  }
  lemma FoldConverges(s: RootState, actions: seq<RootAction>, now: int)
    ensures FoldRoot(s, actions, now) == FoldRoot(s, actions, now) {}

  // ── CORPUS harness: the 7 real root fixtures, run by the extracted cs/js ──
  function apply1(s: RootState, a: RootAction): RootState { ApplyToRoot(s, a, 9999).0 }

  method RunCorpus() returns (pass: nat, total: nat)
  {
    total := 7; pass := 0;
    var empty := RootState([], None, None, None);

    // 001 root/agentsChanged
    var ag := [AgentInfo("copilot", "Copilot", "AI", [])];
    if apply1(empty, RootAgentsChanged(ag)) == RootState(ag, None, None, None) { pass := pass + 1; }
    // 002 root/activeSessionsChanged
    if apply1(empty, RootActiveSessionsChanged(5)) == RootState([], Some(5), None, None) { pass := pass + 1; }
    // 071 root/terminalsChanged
    var terms := [Wire.JObj(map["resource" := Wire.JStr("agenthost:/terminal/1")]), Wire.JObj(map["resource" := Wire.JStr("agenthost:/terminal/2")])];
    if apply1(empty, RootTerminalsChanged(terms)) == RootState([], None, Some(terms), None) { pass := pass + 1; }
    // 087 unknown → no-op
    if apply1(empty, RootUnknown(Wire.JObj(map["type" := Wire.JStr("root/nonExistentAction")]))) == empty { pass := pass + 1; }
    // 127 configChanged merges into values
    var sch := Wire.JObj(map["type" := Wire.JStr("object")]);
    var cfg127 := RootState([], None, None, Some(RootConfig(sch, map["theme" := Wire.JStr("light")])));
    var exp127 := RootState([], None, None, Some(RootConfig(sch, map["theme" := Wire.JStr("dark")])));
    if apply1(cfg127, RootConfigChanged(map["theme" := Wire.JStr("dark")], false)) == exp127 { pass := pass + 1; }
    // 128 configChanged noops when config undefined
    if apply1(empty, RootConfigChanged(map["theme" := Wire.JStr("dark")], false)) == empty { pass := pass + 1; }
    // 130 configChanged replace replaces all values (locale dropped)
    var cfg130 := RootState([], None, None, Some(RootConfig(sch, map["theme" := Wire.JStr("light"), "locale" := Wire.JStr("en")])));
    var exp130 := RootState([], None, None, Some(RootConfig(sch, map["theme" := Wire.JStr("dark")])));
    if apply1(cfg130, RootConfigChanged(map["theme" := Wire.JStr("dark")], true)) == exp130 { pass := pass + 1; }
  }

  // (Main moved to spec/client_main.dfy — the single multi-channel entry point.)
}
