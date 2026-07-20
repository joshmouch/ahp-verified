// AHP Dafny client — FULL Wire.Json ⇄ ResourceWatchAction codec (resourceWatch channel).
//
// Same shape as spec/fidelity.dfy's Root codec, for the resourceWatch channel:
//   • DECODE is promoted VERBATIM from spec/replay/resourcewatch_replay.dfy's
//     `decodeResourceWatchAction` — the exact logic that drives that channel's
//     real fixtures (200/201) through the reducer.
//   • ENCODE is its inverse (canonical wire object per TYPED variant).
//   • Together they carry the load-bearing transport property: decode ∘ encode
//     == id on the typed variant (ResourceWatchActionRoundTrip below).
//   • encodeResourceWatchState is a deterministic total serializer of the
//     channel state (mirrors Fidelity.encodeRootState).
//
// The shared Wire.Json field accessors (field/asStr/asArr/asObjMap/asInt/asBool) are
// REUSED from Fidelity (imported as F), not re-duplicated.
include "../resource_watch.dfy"
include "../fidelity.dfy"

module ResourceWatchCodec {
  // ── .doo FIREWALL: explicit export set — ONLY the 6 symbols spec/ahp.dfy (RWc) consumes.
  //   EncodeResourceWatchAction is REVEALED because ahp.ResourceWatchPrefix computes
  //   typeOf(RWc.EncodeResourceWatchAction(a)) (typeOf reads the "type" field) and must
  //   unfold Encode's body to see the "resourceWatch/changed" prefix. Decode + the two
  //   state fns stay opaque: RouteResourceWatch is reflexive over Decode, and
  //   AhpStateRoundTrip only unfolds the LOCAL encodeAhpState then leans on the two
  //   round-trip lemmas (encode/decodeResourceWatchState are never unfolded downstream).
  //   provides Wire (Wire.Json in every exported signature) + ResourceWatch
  //   (ResourceWatchAction/ResourceWatchState in every signature AND the RWChanged/
  //   RWUnknown constructors in Encode's revealed body). AhpSkeleton/Fidelity are NOT
  //   re-exported — no exported signature names their types (F.* lives only in opaque bodies).
  export
    provides DecodeResourceWatchAction, encodeResourceWatchState, decodeResourceWatchState
    provides ResourceWatchActionRoundTrip, ResourceWatchStateRoundTrip
    provides Wire, ResourceWatch
    reveals EncodeResourceWatchAction

  import opened ResourceWatch
  import opened AhpSkeleton   // Option
  import Wire = ConfluxCodec   // Wire.Json, Wire.JObj, Wire.JStr, Wire.JBool, Wire.JNull (no longer re-exported by AhpSkeleton)
  import F = Fidelity          // shared accessors: field / asStr / asBool / ...

  // ════════════════════════════════════════════════════════════════════════
  //  The typed codec.
  //
  //  ResourceWatchAction has exactly two variants:
  //    | RWChanged(changes: Wire.Json)   — resourceWatch/changed passthrough notice
  //    | RWUnknown(raw: Wire.Json)       — unknown → no-op (verbatim leaf)
  //
  //  There is ONE typed variant (RWChanged); RWUnknown is the honest carve-out
  //  leaf, exactly like Fidelity's RootUnknown.
  // ════════════════════════════════════════════════════════════════════════

  // ── DECODE — promoted from resourcewatch_replay.decodeResourceWatchAction ──
  function DecodeResourceWatchAction(j: Wire.Json): ResourceWatchAction {
    var t := F.asStr(F.field(j, "type"));
    if t == "resourceWatch/changed"
    then RWChanged(if F.field(j, "changes").Some? then F.field(j, "changes").value else Wire.JNull)
    else RWUnknown(j)
  }

  // ── ENCODE — the inverse: canonical wire object per typed variant ──
  function EncodeResourceWatchAction(a: ResourceWatchAction): Wire.Json {
    match a
    case RWChanged(changes) => Wire.JObj(map["type" := Wire.JStr("resourceWatch/changed"), "changes" := changes])
    case RWUnknown(raw)     => raw
  }

  // ── ROUND-TRIP: decode ∘ encode == id on the TYPED variant ──
  // (RWUnknown is excluded — its raw payload may itself be a
  //  resourceWatch/changed-shaped object, so decode would re-classify it; that
  //  leaf is handled by the verbatim lemmas below, not this one — exactly the
  //  RootUnknown carve-out in Fidelity.)
  // Json→domain→Json fixpoint: the codec's canonical wire form re-decodes + re-encodes to ITSELF — the
  // direction the conformance suite exercises, now machine-checked. Follows from the direction-1 round-trip.
  lemma ResourceWatchWireCanonicalRoundTrip(a: ResourceWatchAction)
    requires !a.RWUnknown?
    ensures EncodeResourceWatchAction(DecodeResourceWatchAction(EncodeResourceWatchAction(a))) == EncodeResourceWatchAction(a)
  { ResourceWatchActionRoundTrip(a); }
  lemma ResourceWatchActionRoundTrip(a: ResourceWatchAction)
    requires !a.RWUnknown?
    ensures DecodeResourceWatchAction(EncodeResourceWatchAction(a)) == a
  {
    match a
    case RWChanged(changes) =>
      // EncodeResourceWatchAction(RWChanged(changes)) is a Wire.JObj carrying
      // "type" = "resourceWatch/changed" and "changes" = changes, so decode
      // takes the RWChanged branch and reads `changes` straight back.
  }

  // Non-vacuity witness: the round-trip precondition is inhabited (a genuine
  // typed action round-trips), so ResourceWatchActionRoundTrip is NOT vacuous.
  lemma ResourceWatchActionRoundTrip_NonVacuous()
    ensures exists a: ResourceWatchAction :: !a.RWUnknown? && DecodeResourceWatchAction(EncodeResourceWatchAction(a)) == a
  {
    var a := RWChanged(Wire.JObj(map["items" := Wire.JArr([Wire.JObj(map["uri" := Wire.JStr("file:///workspace/a.txt"), "type" := Wire.JStr("added")])])]));
    ResourceWatchActionRoundTrip(a);
    assert !a.RWUnknown? && DecodeResourceWatchAction(EncodeResourceWatchAction(a)) == a;
  }

  // ── unknown-variant passthrough — the RWUnknown leaf re-encodes verbatim ──
  // (honest carve-out: exactly parallel to Fidelity.RootUnknown_EncodeVerbatim /
  //  Fidelity.UnknownVariant_Verbatim)
  lemma RWUnknown_EncodeVerbatim(j: Wire.Json)
    ensures EncodeResourceWatchAction(RWUnknown(j)) == j {}
  lemma RWUnknownVariant_Verbatim(j: Wire.Json)
    requires DecodeResourceWatchAction(j).RWUnknown?
    ensures EncodeResourceWatchAction(DecodeResourceWatchAction(j)) == j {}

  // ── STATE serialization (deterministic total; mirrors encodeRootState) ──
  // ResourceWatchState models every real field of the watch-config slice
  // (root, recursive). This serializes it to canonical JSON so the host and the
  // client can each emit their final state for a cross-process convergence
  // check. DETERMINISTIC TOTAL function; the demo's convergence relies on that
  // determinism, not on a state round-trip lemma (a natural future hardening).
  function encodeResourceWatchState(s: ResourceWatchState): Wire.Json {
    Wire.JObj(map["root" := Wire.JStr(s.root), "recursive" := Wire.JBool(s.recursive)])
  }

  // ── STATE DECODE — the true inverse of encodeResourceWatchState ──
  // Promoted from spec/replay/resourcewatch_replay.decodeResourceWatchState —
  // the exact logic that drives the resourceWatch fixtures' initial/expected
  // state through the reducer. Reads each field encodeResourceWatchState emits
  // verbatim (root via asStr, recursive via asBool), rebuilding the state.
  function decodeResourceWatchState(j: Wire.Json): ResourceWatchState {
    ResourceWatchState(F.asStr(F.field(j, "root")), F.asBool(F.field(j, "recursive"), false))
  }

  // ── STATE ROUND-TRIP: decode ∘ encode == id (mirrors ResourceWatchActionRoundTrip) ──
  // Every field of ResourceWatchState (root: string, recursive: bool) is a
  // primitive that encodeResourceWatchState emits verbatim under a fixed key, so
  // decodeResourceWatchState reads each one straight back. UNCONDITIONAL — no
  // WfResourceWatchState precondition is needed: there is no Option<Wire.Json>
  // Some(Wire.JNull)-collapse field and no derived/redundant field to carve out.
  lemma ResourceWatchStateRoundTrip(s: ResourceWatchState)
    ensures decodeResourceWatchState(encodeResourceWatchState(s)) == s
  {
    // encodeResourceWatchState(s) is a Wire.JObj carrying "root" = Wire.JStr(s.root) and
    // "recursive" = Wire.JBool(s.recursive); both keys are present, so
    //   F.field(enc, "root")      = Some(Wire.JStr(s.root))      → F.asStr  = s.root
    //   F.field(enc, "recursive") = Some(Wire.JBool(s.recursive)) → F.asBool = s.recursive
    // hence decodeResourceWatchState(enc) == ResourceWatchState(s.root, s.recursive) == s.
  }

  // Non-vacuity witness: a genuine non-empty state round-trips, so
  // ResourceWatchStateRoundTrip is NOT vacuous.
  lemma ResourceWatchStateRoundTrip_NonVacuous()
    ensures exists s: ResourceWatchState :: decodeResourceWatchState(encodeResourceWatchState(s)) == s
  {
    var s := ResourceWatchState("file:///workspace/project", true);
    ResourceWatchStateRoundTrip(s);
    assert decodeResourceWatchState(encodeResourceWatchState(s)) == s;
  }

  ghost predicate ResourceWatchCodecWf(a: ResourceWatchAction) { !a.RWUnknown? }
  function ResourceWatchCodecC(): Wire.Codec<ResourceWatchAction> {
    Wire.Codec(EncodeResourceWatchAction, (j: Wire.Json) => Wire.Some(DecodeResourceWatchAction(j)))
  }
  lemma ResourceWatchCodecRoundTripsWhere()
    ensures forall a :: ResourceWatchCodecWf(a) ==> ResourceWatchCodecC().decode(ResourceWatchCodecC().encode(a)) == Wire.Some(a)
  {
    forall a | ResourceWatchCodecWf(a)
      ensures ResourceWatchCodecC().decode(ResourceWatchCodecC().encode(a)) == Wire.Some(a)
    { ResourceWatchActionRoundTrip(a); }
  }
}
