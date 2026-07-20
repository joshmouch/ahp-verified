// AHP Dafny client — resourceWatch channel (proves the pattern GENERALIZES beyond root).
// A structurally-different channel: `changed` is a passthrough (state unchanged),
// `unknown` is a no-op. Full coverage of both real fixtures (200/201). Extracts + runs.

include "ahp_skeleton.dfy"

module ResourceWatch {
  export
    provides ApplyToResourceWatch, apply1, RunCorpus
    reveals fold
    provides AhpSkeleton, Wire, CC
    reveals ResourceWatchState, ResourceWatchAction

  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import CC = ConfluxContract

  datatype ResourceWatchState = ResourceWatchState(root: string, recursive: bool)

  datatype ResourceWatchAction =
    | RWChanged(changes: Wire.Json)     // resourceWatch/changed — passthrough notification
    | RWUnknown(raw: Wire.Json)         // unknown → no-op

  // Passthrough channel: neither action mutates the watch-config state slice.
  function ApplyToResourceWatch(s: ResourceWatchState, a: ResourceWatchAction, now: int): (ResourceWatchState, ReduceOutcome)
  {
    match a
    case RWChanged(_) => (s, Applied)   // 200: recognized notification, state unchanged (passthrough)
    case RWUnknown(_) => (s, NoOp)      // 201: unknown → no-op
  }

  predicate IsUnknownRW(a: ResourceWatchAction) { a.RWUnknown? }

  // No-op / passthrough lemmas + non-vacuity.
  lemma RW_UnknownIsNoOp(s: ResourceWatchState, a: ResourceWatchAction, now: int)
    requires IsUnknownRW(a)
    ensures ApplyToResourceWatch(s, a, now).0 == s && ApplyToResourceWatch(s, a, now).1 == NoOp {}
  lemma RW_UnknownIsNoOp_NonVacuous()
    ensures exists a: ResourceWatchAction :: IsUnknownRW(a)
  { assert IsUnknownRW(RWUnknown(Wire.JNull)); }
  // The passthrough invariant: `changed` never mutates the watch-config state.
  lemma RW_ChangedIsPassthrough(s: ResourceWatchState, changes: Wire.Json, now: int)
    ensures ApplyToResourceWatch(s, RWChanged(changes), now).0 == s {}

  function apply1(s: ResourceWatchState, a: ResourceWatchAction): ResourceWatchState
  { ApplyToResourceWatch(s, a, 9999).0 }
  function fold(s: ResourceWatchState, acts: seq<ResourceWatchAction>): ResourceWatchState
  { CC.Fold(apply1, s, acts) }

  method RunCorpus() returns (pass: nat, total: nat) {
    total := 2; pass := 0;
    // 200 resourceWatch/changed passthrough
    var s0 := ResourceWatchState("file:///workspace", true);
    var changes := Wire.JObj(map["items" := Wire.JArr([Wire.JObj(map["uri" := Wire.JStr("file:///workspace/a.txt"), "type" := Wire.JStr("added")])])]);
    if apply1(s0, RWChanged(changes)) == s0 { pass := pass + 1; }
    // 201 resourceWatch unknown → no-op
    var s1 := ResourceWatchState("file:///workspace", false);
    if apply1(s1, RWUnknown(Wire.JObj(map["type" := Wire.JStr("resourceWatch/unknownAction")]))) == s1 { pass := pass + 1; }
  }

  // (Main moved to spec/client_main.dfy — the single multi-channel entry point.)
}
