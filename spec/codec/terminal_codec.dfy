// AHP Dafny client — FULL Wire.Json ⇄ TerminalAction codec (terminal channel).
//
// Same shape as spec/fidelity.dfy's Root codec, for the terminal channel:
//   • DECODE is promoted VERBATIM from spec/replay/terminal_replay.dfy's
//     `decodeTerminalAction` — the exact logic that drives that channel's 19
//     real fixtures through the reducer.
//   • ENCODE is its inverse (canonical wire object per TYPED variant).
//   • Together they carry the load-bearing transport property: decode ∘ encode
//     == id on the typed variants (TerminalActionRoundTrip below).
//   • encodeTerminalState is a deterministic total serializer of the channel
//     state (mirrors Fidelity.encodeRootState).
//
// The shared Wire.Json field accessors (field/asStr/asArr/asObjMap/asInt/asBool) are
// REUSED from Fidelity (imported as F), not re-duplicated.
include "../terminal.dfy"
include "../fidelity.dfy"

module TerminalCodec {
  // ── Export firewall ──────────────────────────────────────────────────────
  // Sole consumer is spec/ahp.dfy (alias `Tc`), which references only the action +
  // state codec entry points and the two round-trip lemmas it composes. Hidden: the
  // Part sub-codec (encodePart/decodePart/encodeParts/decodeParts), the optional-field
  // decoders (optStr/optBool), the Wire.Codec wrapper (TerminalCodecWf/TerminalCodecC/
  // TerminalCodecRoundTripsWhere), and the verbatim/canonical/non-vacuity lemmas.
  // EncodeTerminalAction is REVEALED (not provided) because ahp.dfy's TerminalPrefix
  // unfolds its body to prove typeOf(EncodeTerminalAction(a)) starts with "terminal/";
  // the other three codec fns stay opaque (consumers only compose the round-trip
  // lemmas' ensures, never unfold the bodies). provides Terminal + Wire because the
  // exported signatures name Terminal.TerminalAction/TerminalState and Wire.Json
  // (transitive-module-export rule).
  export
    provides DecodeTerminalAction, encodeTerminalState, decodeTerminalState
    provides TerminalActionRoundTrip, TerminalStateRoundTrip
    reveals EncodeTerminalAction
    provides Terminal, Wire

  import opened Terminal
  import opened AhpSkeleton   // Option
  import Wire = ConfluxCodec   // Wire.Json, Wire.JObj, Wire.JStr, Wire.JNum, Wire.JBool, Wire.JArr, Wire.JNull (no longer re-exported by AhpSkeleton)
  import F = Fidelity          // shared accessors: field / asStr / asInt / asBool / ...
  import CO = ConfluxOperators

  // ════════════════════════════════════════════════════════════════════════
  //  The typed codec.
  //
  //  TerminalAction has 12 variants (terminal.dfy):
  //    | TCwdChanged(cwd)                            — terminal/cwdChanged
  //    | TTitleChanged(title)                        — terminal/titleChanged
  //    | TResized(cols, rows)                        — terminal/resized
  //    | TExited(code)                               — terminal/exited
  //    | TData(data)                                 — terminal/data
  //    | TCleared                                    — terminal/cleared
  //    | TClaimed(claim: Wire.Json)                       — terminal/claimed
  //    | TCommandDetectionAvailable                  — terminal/commandDetectionAvailable
  //    | TCommandExecuted(commandId, commandLine, timestamp)
  //    | TCommandFinished(commandId, code, durationMs)
  //    | TInput                                      — terminal/input (no-op)
  //    | TUnknown(raw: Wire.Json)                          — unknown (verbatim leaf)
  //
  //  Eleven are TYPED variants that round-trip; TUnknown is the honest carve-out
  //  leaf, exactly like Fidelity's RootUnknown.
  // ════════════════════════════════════════════════════════════════════════

  // ── DECODE — promoted from terminal_replay.decodeTerminalAction ──
  function DecodeTerminalAction(j: Wire.Json): TerminalAction {
    var t := F.asStr(F.field(j, "type"));
    if      t == "terminal/cwdChanged"   then TCwdChanged(F.asStr(F.field(j, "cwd")))
    else if t == "terminal/titleChanged" then TTitleChanged(F.asStr(F.field(j, "title")))
    else if t == "terminal/resized"      then TResized(asInt(F.field(j, "cols")), asInt(F.field(j, "rows")))
    else if t == "terminal/exited"       then TExited(asInt(F.field(j, "exitCode")))
    else if t == "terminal/data"         then TData(F.asStr(F.field(j, "data")))
    else if t == "terminal/cleared"      then TCleared
    else if t == "terminal/claimed"      then TClaimed(if F.field(j, "claim").Some? then F.field(j, "claim").value else Wire.JNull)
    else if t == "terminal/commandDetectionAvailable" then TCommandDetectionAvailable
    else if t == "terminal/commandExecuted" then
      TCommandExecuted(F.asStr(F.field(j, "commandId")), F.asStr(F.field(j, "commandLine")), asInt(F.field(j, "timestamp")))
    else if t == "terminal/commandFinished" then
      TCommandFinished(F.asStr(F.field(j, "commandId")), asInt(F.field(j, "exitCode")), asInt(F.field(j, "durationMs")))
    else if t == "terminal/input"        then TInput
    else TUnknown(j)
  }

  // ── ENCODE — the inverse: canonical wire object per typed variant ──
  //   Each object sets a DISTINCT "type" string, so decode dispatches uniquely
  //   back to the originating variant; the payload keys mirror the decode reads.
  function EncodeTerminalAction(a: TerminalAction): Wire.Json {
    match a
    case TCwdChanged(c)      => Wire.JObj(map["type" := Wire.JStr("terminal/cwdChanged"), "cwd" := Wire.JStr(c)])
    case TTitleChanged(t)    => Wire.JObj(map["type" := Wire.JStr("terminal/titleChanged"), "title" := Wire.JStr(t)])
    case TResized(co, ro)    => Wire.JObj(map["type" := Wire.JStr("terminal/resized"), "cols" := Wire.JNum(co), "rows" := Wire.JNum(ro)])
    case TExited(code)       => Wire.JObj(map["type" := Wire.JStr("terminal/exited"), "exitCode" := Wire.JNum(code)])
    case TData(d)            => Wire.JObj(map["type" := Wire.JStr("terminal/data"), "data" := Wire.JStr(d)])
    case TCleared            => Wire.JObj(map["type" := Wire.JStr("terminal/cleared")])
    case TClaimed(c)         => Wire.JObj(map["type" := Wire.JStr("terminal/claimed"), "claim" := c])
    case TCommandDetectionAvailable => Wire.JObj(map["type" := Wire.JStr("terminal/commandDetectionAvailable")])
    case TCommandExecuted(id, cl, ts) =>
      Wire.JObj(map["type" := Wire.JStr("terminal/commandExecuted"), "commandId" := Wire.JStr(id),
               "commandLine" := Wire.JStr(cl), "timestamp" := Wire.JNum(ts)])
    case TCommandFinished(id, code, dur) =>
      Wire.JObj(map["type" := Wire.JStr("terminal/commandFinished"), "commandId" := Wire.JStr(id),
               "exitCode" := Wire.JNum(code), "durationMs" := Wire.JNum(dur)])
    case TInput              => Wire.JObj(map["type" := Wire.JStr("terminal/input")])
    case TUnknown(raw)       => raw
  }

  // ── ROUND-TRIP: decode ∘ encode == id on the TYPED variants ──
  // (TUnknown is excluded — its raw payload may itself be a terminal/*-shaped
  //  object, so decode would re-classify it; that leaf is handled by the
  //  verbatim lemmas below, not this one — exactly the RootUnknown carve-out in
  //  Fidelity. TInput is NOT excluded: it carries no payload and re-encodes to a
  //  canonical "terminal/input" object that decodes straight back to TInput.)
  // Json→domain→Json fixpoint: the codec's canonical wire form re-decodes + re-encodes to ITSELF — the
  // direction the conformance suite exercises, now machine-checked. Follows from the direction-1 round-trip.
  // (Terminal has the int fields cols/rows/exitCode/durationMs; AhpSkeleton.AsIntCanonicalizes proves a
  // non-canonical integer wire form — 5.0/1e3 → JDec — decodes to the same int, which encode re-emits as JNum.)
  lemma TerminalWireCanonicalRoundTrip(a: TerminalAction)
    requires !a.TUnknown?
    ensures EncodeTerminalAction(DecodeTerminalAction(EncodeTerminalAction(a))) == EncodeTerminalAction(a)
  { TerminalActionRoundTrip(a); }
  lemma TerminalActionRoundTrip(a: TerminalAction)
    requires !a.TUnknown?
    ensures DecodeTerminalAction(EncodeTerminalAction(a)) == a
  {
    // Each typed variant encodes to a Wire.JObj whose "type" key uniquely selects its
    // decode branch, and whose payload keys read back the same fields.
    match a
    case TCwdChanged(c)      =>
    case TTitleChanged(t)    =>
    case TResized(co, ro)    =>
    case TExited(code)       =>
    case TData(d)            =>
    case TCleared            =>
    case TClaimed(c)         =>
    case TCommandDetectionAvailable =>
    case TCommandExecuted(id, cl, ts) =>
    case TCommandFinished(id, code, dur) =>
    case TInput              =>
  }

  // Non-vacuity witness: the round-trip precondition is inhabited (a genuine
  // typed, payload-carrying action round-trips), so TerminalActionRoundTrip is
  // NOT vacuously true.
  lemma TerminalActionRoundTrip_NonVacuous()
    ensures exists a: TerminalAction :: !a.TUnknown? && DecodeTerminalAction(EncodeTerminalAction(a)) == a
  {
    var a := TCommandFinished("cmd-1", 0, 1234);
    TerminalActionRoundTrip(a);
    assert !a.TUnknown? && DecodeTerminalAction(EncodeTerminalAction(a)) == a;
  }

  // ── unknown-variant passthrough — the TUnknown leaf re-encodes verbatim ──
  // (honest carve-out: exactly parallel to Fidelity.RootUnknown_EncodeVerbatim /
  //  Fidelity.UnknownVariant_Verbatim)
  lemma TUnknown_EncodeVerbatim(j: Wire.Json)
    ensures EncodeTerminalAction(TUnknown(j)) == j {}
  lemma TUnknownVariant_Verbatim(j: Wire.Json)
    requires DecodeTerminalAction(j).TUnknown?
    ensures EncodeTerminalAction(DecodeTerminalAction(j)) == j {}

  // ── STATE serialization (deterministic total; mirrors encodeRootState) ──
  // encodeTerminalState serializes every MODELED field of the terminal slice to
  // canonical JSON so the host and the client can each emit their final state
  // for a cross-process convergence check. DETERMINISTIC TOTAL function; the
  // demo's convergence relies on that determinism, not on a state round-trip
  // lemma (a natural future hardening). Optional fields are omitted when None,
  // matching the open-struct decode in terminal_replay.decodeTerminalState.

  // ── encode a content Part — the inverse of terminal_replay.decodePart ──
  function encodePart(p: Part): Wire.Json {
    match p
    case Unclassified(v) => Wire.JObj(map["type" := Wire.JStr("unclassified"), "value" := Wire.JStr(v)])
    case Command(cid, cl, out, ts, comp, ec, dm) =>
      var m0 := map["type" := Wire.JStr("command"), "commandId" := Wire.JStr(cid),
                    "commandLine" := Wire.JStr(cl), "output" := Wire.JStr(out),
                    "timestamp" := Wire.JNum(ts), "isComplete" := Wire.JBool(comp)];
      var m1 := if ec.Some? then m0["exitCode" := Wire.JNum(ec.value)] else m0;
      var m2 := if dm.Some? then m1["durationMs" := Wire.JNum(dm.value)] else m1;
      Wire.JObj(m2)
  }
  function encodeParts(ps: seq<Part>): seq<Wire.Json>
  { CO.Map(encodePart, ps) }

  function encodeTerminalState(s: TerminalState): Wire.Json {
    var m0 := map["title" := Wire.JStr(s.title), "content" := Wire.JArr(encodeParts(s.content))];
    var m1 := if s.cwd.Some?      then m0["cwd" := Wire.JStr(s.cwd.value)]        else m0;
    var m2 := if s.cols.Some?     then m1["cols" := Wire.JNum(s.cols.value)]      else m1;
    var m3 := if s.rows.Some?     then m2["rows" := Wire.JNum(s.rows.value)]      else m2;
    var m4 := if s.claim.Some?    then m3["claim" := s.claim.value]          else m3;
    var m5 := if s.exitCode.Some? then m4["exitCode" := Wire.JNum(s.exitCode.value)] else m4;
    var m6 := if s.supportsCommandDetection.Some?
              then m5["supportsCommandDetection" := Wire.JBool(s.supportsCommandDetection.value)]
              else m5;
    Wire.JObj(m6)
  }

  // ════════════════════════════════════════════════════════════════════════
  //  STATE DECODE + ROUND-TRIP — completes the codec's proof story in BOTH
  //  directions. The ACTION codec already round-trips (TerminalActionRoundTrip);
  //  this makes the STATE serializer encodeTerminalState a PROVEN-INVERTIBLE
  //  function — decodeTerminalState is its TRUE INVERSE, not merely a
  //  deterministic serializer. Promoted from the proven decodeTerminalState in
  //  spec/replay/terminal_replay.dfy (the exact logic that decodes all 19
  //  terminal fixtures' initial/expected states), re-expressed over the shared
  //  Fidelity accessors (F.field/asStr/asInt/asBool/asArr) + the optional-field
  //  and Part decoders below. Additive only: encodeTerminalState / encodePart /
  //  encodeParts and the action codec above are UNCHANGED.
  // ════════════════════════════════════════════════════════════════════════

  // ── optional-field decoders: present+typed → Some(value); else None ──
  //    (mirror terminal_replay.optStr/optInt/optBool; chat_codec has the same trio.)
  function optStr(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<string> {
    if o.Some? && o.value.JStr? then Some(o.value.s) else None
  }
  // optInt (JNum + integer-valued JDec) lives in AhpSkeleton now; in scope via `import opened AhpSkeleton`.
  function optBool(o: AhpSkeleton.Option<Wire.Json>): AhpSkeleton.Option<bool> {
    if o.Some? && o.value.JBool? then Some(o.value.b) else None
  }

  // ── decode a content Part — the EXACT inverse of encodePart above ──
  //   "command" → Command(...); any other "type" → Unclassified(value). The
  //   optional exitCode/durationMs decode via optInt (present-key ⇔ Some), the
  //   precise inverse of encodePart's conditional m1/m2 updates.
  function decodePart(j: Wire.Json): Part {
    var t := F.asStr(F.field(j, "type"));
    if t == "command" then
      Command(
        F.asStr(F.field(j, "commandId")),
        F.asStr(F.field(j, "commandLine")),
        F.asStr(F.field(j, "output")),
        asInt(F.field(j, "timestamp")),
        F.asBool(F.field(j, "isComplete"), false),
        optInt(F.field(j, "exitCode")),      // Option<int> — absent on incomplete commands
        optInt(F.field(j, "durationMs")))    // Option<int>
    else
      Unclassified(F.asStr(F.field(j, "value")))
  }
  function decodeParts(js: seq<Wire.Json>): seq<Part>
  { CO.Map(decodePart, js) }

  // ── decode a TerminalState — reads EXACTLY the keys encodeTerminalState emits ──
  //   title  → Wire.JStr, verbatim;                 content → Wire.JArr(encodeParts), inverse via PartsRoundTrip;
  //   cwd    → present Wire.JStr  ⇔ Some / absent ⇔ None;   cols/rows/exitCode → present Wire.JNum ⇔ Some;
  //   claim  → opaque Wire.Json carried VERBATIM, present-key ⇔ Some;
  //   supportsCommandDetection → present Wire.JBool ⇔ Some.
  function decodeTerminalState(j: Wire.Json): TerminalState {
    TerminalState(
      F.asStr(F.field(j, "title")),
      optStr(F.field(j, "cwd")),
      optInt(F.field(j, "cols")),
      optInt(F.field(j, "rows")),
      decodeParts(F.asArr(F.field(j, "content"))),
      F.field(j, "claim"),                       // Option<Wire.Json>: Some(rawClaim) if present, else None
      optInt(F.field(j, "exitCode")),
      optBool(F.field(j, "supportsCommandDetection")))
  }

  // ── Part round-trip (mirrors Fidelity.AgentRoundTrip / changeset FileRoundTrip;
  //    the leaf of the seq induction). The case-split on each Option<int> mirrors
  //    FileRoundTrip's split on its conditional-map field. ──
  lemma PartRoundTrip(p: Part)
    ensures decodePart(encodePart(p)) == p
  {
    match p
    case Unclassified(v) =>
    case Command(cid, cl, out, ts, comp, ec, dm) =>
      if ec.Some? {} else {}
      if dm.Some? {} else {}
  }

  lemma PartsRoundTrip(ps: seq<Part>)
    ensures decodeParts(encodeParts(ps)) == ps
    decreases ps
  { if |ps| == 0 {} else { PartRoundTrip(ps[0]); PartsRoundTrip(ps[1..]); } }

  // ── STATE ROUND-TRIP: decodeTerminalState ∘ encodeTerminalState == id ──
  // REAL, quantified (∀ s), non-vacuous. NO WfTerminalState precondition is
  // needed — EVERY field round-trips cleanly:
  //   • title/content are required and always emitted;
  //   • cwd/cols/rows/exitCode/supportsCommandDetection are Option<scalar> whose
  //     Some value always emits under a matching-typed key, so present-key ⇔ Some;
  //   • claim is Option<Wire.Json> carried VERBATIM and keyed off PRESENCE, so it does
  //     NOT suffer the Some(Wire.JNull)→absent collapse that forces a carve-out on the
  //     action-level scalar decoders — Some(Wire.JNull) decodes back to Some(Wire.JNull),
  //     absent decodes to None. Nothing is dropped; unlike the TUnknown action
  //     leaf, the STATE needs no honest carve-out — hence no requires clause.
  lemma TerminalStateRoundTrip(s: TerminalState)
    ensures decodeTerminalState(encodeTerminalState(s)) == s
  {
    PartsRoundTrip(s.content);
  }

  // Non-vacuity witness: a GENUINE non-empty state — non-empty content (an
  // Unclassified prompt + a COMPLETED command carrying Some(exitCode)/Some(durationMs),
  // exercising the encodeParts/decodeParts induction end to end), every optional
  // field Some (including an opaque claim Wire.Json) — round-trips. So the lemma is not
  // accidentally provable only for the empty/None state.
  lemma TerminalStateRoundTrip_NonVacuous()
    ensures exists s: TerminalState ::
      |s.content| > 0 && decodeTerminalState(encodeTerminalState(s)) == s
  {
    var content := [
      Unclassified("$ "),
      Command("cmd-1", "npm test", "All tests passed\r\n", 1700000000000, true, Some(0), Some(1234))];
    var s := TerminalState(
      "zsh", Some("/tmp"), Some(80), Some(24),
      content, Some(Wire.JObj(map["kind" := Wire.JStr("client"), "clientId" := Wire.JStr("c1")])),
      Some(0), Some(true));
    TerminalStateRoundTrip(s);
    assert |s.content| > 0 && decodeTerminalState(encodeTerminalState(s)) == s;
  }

  ghost predicate TerminalCodecWf(a: TerminalAction) { !a.TUnknown? }
  function TerminalCodecC(): Wire.Codec<TerminalAction> {
    Wire.Codec(EncodeTerminalAction, (j: Wire.Json) => Wire.Some(DecodeTerminalAction(j)))
  }
  lemma TerminalCodecRoundTripsWhere()
    ensures forall a :: TerminalCodecWf(a) ==> TerminalCodecC().decode(TerminalCodecC().encode(a)) == Wire.Some(a)
  {
    forall a | TerminalCodecWf(a)
      ensures TerminalCodecC().decode(TerminalCodecC().encode(a)) == Wire.Some(a)
    { TerminalActionRoundTrip(a); }
  }
}
