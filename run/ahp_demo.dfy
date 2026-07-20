// AHP Dafny client — RUNNABLE multi-channel DEMO (standalone; NOT part of core.doo).
//
// This EXECUTES the unified multi-channel layer (module Ahp, spec/ahp.dfy) on the
// EXTRACTED code (cs + js). It does NOT re-prove anything and it does NOT
// reimplement dispatch or the codec — it CALLS the verified functions directly:
//
//   • Ahp.applyAhp       — the pure per-channel dispatch reducer
//   • Ahp.encodeAhpAction / Ahp.decodeAhpAction — the unified routing codec
//   • Ahp.encodeAhpState — the 8-keyed product-state serializer
//
// What it proves by RUNNING (not by proof):
//   1. A MIXED-channel AhpAction stream folds through Ahp.applyAhp start→final.
//   2. Every action survives decode(encode(a)) == a at RUNTIME on extracted code.
//   3. The final product state serializes via Ahp.encodeAhpState.
//
// The success marker is guarded by REAL runtime == checks: any round-trip
// mismatch prints a distinct NON-OK marker and returns WITHOUT the OK line.
// No `assume`, no `{:axiom}`. Includes spec/ahp.dfy directly (all source is local).

include "../spec/ahp.dfy"

module AhpDemo {
  import Ahp
  import opened AhpSkeleton          // Json, RootState, RootAction ctors, Option (Some/None), AgentInfo
  import Session
  import Chat
  import Terminal
  import Changeset
  import Annotations
  import ResourceWatch
  import Canvas

  // Channel label of a unified action — used to count distinct channels the
  // stream mixes (computed from the real actions, not hard-coded).
  function channelOf(a: Ahp.AhpAction): string {
    match a
    case ARoot(_)          => "root"
    case ASession(_)       => "session"
    case AChat(_)          => "chat"
    case ATerminal(_)      => "terminal"
    case AChangeset(_)     => "changeset"
    case AAnnotations(_)   => "annotations"
    case AResourceWatch(_) => "resourceWatch"
    case ACanvas(_)        => "canvas"
  }

  // The default/empty product state — one empty state per channel (8 channels).
  function InitialState(): Ahp.AhpState {
    Ahp.AhpState(
      RootState([], None, None, None),                                    // root
      Session.S0(),                                                       // session
      Chat.C0(),                                                          // chat
      Terminal.T0(),                                                      // terminal
      Changeset.ChangesetState("ready", [], None, None),                 // changeset
      Annotations.AnnotationsState([]),                                   // annotations
      ResourceWatch.ResourceWatchState("", false),                       // resourceWatch
      Canvas.C0())                                                        // canvas
  }

  method Main(args: seq<string>) {
    // ── (2) a real MIXED-channel action stream: 7 TYPED actions, 6 channels ──
    // Each action MEANINGFULLY changes its channel's slice of the product state.
    var actions: seq<Ahp.AhpAction> := [
      Ahp.ARoot(RootAgentsChanged([AgentInfo("copilot", "Copilot", "AI assistant", ["gpt-5"])])),  // root
      Ahp.ASession(Session.TitleChanged("Renamed Session")),                                        // session
      Ahp.AChat(Chat.CDraftChanged(Some("draft in flight"))),                                       // chat
      Ahp.ATerminal(Terminal.TCwdChanged("/home/user/project")),                                    // terminal
      Ahp.ARoot(RootActiveSessionsChanged(3)),                                                       // root (again)
      Ahp.AChangeset(Changeset.StatusChanged("staged", None)),                                       // changeset
      Ahp.ACanvas(Canvas.Updated(Some("README.md"), Some("editing"), None, None))                     // canvas
    ];

    // distinct channels mixed (computed from the actual stream).
    var chans: set<string> := {};
    for i := 0 to |actions| {
      chans := chans + {channelOf(actions[i])};
    }

    print "AHP MULTI-CHANNEL DEMO — executing the unified layer on extracted code\n";
    print "  stream: ", |actions|, " actions across ", |chans|, " channels\n";

    // ── (1)+(3) FOLD the mixed stream through the VERIFIED dispatch ──
    var st := InitialState();
    for i := 0 to |actions| {
      st := Ahp.applyAhp(st, actions[i]);   // real call into Ahp.applyAhp
    }
    var foldCompleted := true;

    // Non-vacuity of the fold: prove state ACTUALLY moved off the initial state.
    if st == InitialState() {
      print "NON-OK: fold produced NO state change (dispatch is a no-op stub?)\n";
      return;
    }

    // ── (4) round-trip EVERY action through the unified codec AT RUNTIME ──
    var allRT := true;
    for i := 0 to |actions| {
      var a := actions[i];
      var re := Ahp.decodeAhpAction(Ahp.encodeAhpAction(a));   // real codec calls
      if re != a {
        print "NON-OK: codec round-trip FAILED for action index ", i,
              " (channel ", channelOf(a), ")\n";
        allRT := false;
      }
    }

    // ── (5) serialize the final product state via the real state serializer ──
    var finalJson := Ahp.encodeAhpState(st);
    var stateKeys := if finalJson.JObj? then |finalJson.fields| else 0;
    print "  final state serialized to a ", stateKeys, "-key product Json\n";
    var stateOk := finalJson.JObj? && stateKeys == 8;
    if !stateOk {
      print "NON-OK: encodeAhpState did not yield the 8-key product object\n";
      return;
    }

    // ── (6) the OK marker — guarded by fold + every runtime == check ──
    if foldCompleted && allRT && stateOk {
      print "AHP MULTI-CHANNEL DEMO OK: ", |actions|,
            " mixed-channel actions folded (", |chans|,
            " channels) + codec round-tripped\n";
    } else {
      print "NON-OK: demo did not satisfy all guards\n";
    }
  }
}
