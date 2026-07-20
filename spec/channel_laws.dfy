// CHANNEL LAWS — Campaign D: "everything is a Channel", instantiated for all 8 AHP channels.
//
// Each AHP channel is already reduced at the envelope level: its `fold` delegates to
// ConfluxContract.Fold over the retained per-action step `apply1` (= ApplyTo<Chan>(·,·,now).0).
// What was still MISSING corpus-wide (per floor/BASELINE.md — "HostCanonicalIsFold is instantiated
// NOWHERE in AHP yet") is the generic host-authority law tying that fold to the streaming envelope:
//
//   ConfluxChannel.HostCanonicalIsFold : AcceptFold(r, h, actions).canonical == Fold(r, h.canonical, actions)
//   ConfluxChannel.HostSeqAdvances     : AcceptFold(r, h, actions).nextSeq    == h.nextSeq + |actions|
//
// This file INSTANTIATES those two kernel laws (NOT re-proves them) at each channel's own `apply1`,
// establishing "this channel's host-authority canonical state IS the ConfluxContract.Fold of its action
// stream, and its assigned seq advances by exactly one per action" — the host half of the channel
// obligation. `ChannelIsFold<S,A>` proves the generic corollary ONCE by delegating to the two kernel
// lemmas; every channel is a one-line instantiation passing its own `apply1` (the retained domain step f).
//
// This is an INSTANTIATION, not a correspondence bridge: every `ensures` relates two conflux primitives
// (AcceptFold and Fold) at the channel's reducer VALUE `apply1` — there is no `consumer == primitive`
// equation (the host state `h` is a parameter, never an inline HostState(...) construction). The definitional
// tie to each channel's OWN `fold` (<Chan>.fold ≡ CC.Fold(apply1,·,·)) is discharged inside the lemma BODY,
// where it belongs — it is the proof, not the theorem.
//
// PeersConverge (host==client mirror) is NOT instantiated here: core owns no live Client seq — that seam
// lives in agent-host-protocol-client-dafny, and the N-client mirror convergence already discharges through
// ConfluxChannel.PeersConverge in agent-host-protocol-convergence-dafny/spec/convergence.dfy. Here we prove
// the host-authority (Accept/seq) half, which is exactly the part a consumer with no live client can carry.
//
// VERIFY-ONLY: consumes the core surface + the streaming envelope through the SAME --library seam
// convergence.dfy uses — `--library agent-host-protocol-core.doo --library vendor/conflux-channel.doo`
// (see gen/check-all.sh step 1d). No source `include` of the channel files (that would double-load
// ConfluxContract against the .doo copy); ConfluxChannel is verify-only, so nothing here is translated.
module ChannelLaws {
  import opened AhpSkeleton
  import Session
  import Chat
  import Terminal
  import Changeset
  import Annotations
  import ResourceWatch
  import Canvas
  import Ch = ConfluxChannel
  import CC = ConfluxContract

  // ── The generic host-authority law, proven ONCE by instantiating the two kernel lemmas ──
  // For ANY channel reducer r and ANY host h: folding `actions` reaches ConfluxContract.Fold of the same
  // stream from h.canonical, and the assigned seq advances by exactly |actions|. Both `ensures` relate two
  // kernel primitives (AcceptFold, Fold) — the "everything is a Channel" law, generic over the reducer.
  lemma ChannelIsFold<S, A>(r: CC.Reducer<S, A>, h: Ch.HostState<S>, actions: seq<A>)
    ensures Ch.AcceptFold(r, h, actions).canonical == CC.Fold(r, h.canonical, actions)
    ensures Ch.AcceptFold(r, h, actions).nextSeq == h.nextSeq + |actions|
  {
    Ch.HostCanonicalIsFold(r, h, actions);   // canonical == Fold
    Ch.HostSeqAdvances(r, h, actions);        // nextSeq == h.nextSeq + |actions|
  }

  // ── Per-channel instantiations at each channel's own retained step `apply1` ──
  // The body-level assert ties the host canonical to the channel's OWN `fold` (definitional:
  // <Chan>.fold ≡ CC.Fold(apply1,·,·)) — kept in the proof body, not the spec, since <Chan>.fold is a
  // consumer name and the theorem is the kernel law.
  // Session's export `provides fold` (body opaque via its firewall — kept as-is to leave the exported .doo
  // interface untouched), so unlike the channels below we do NOT unfold Session.fold in the body; the theorem
  // is stated at CC.Fold(Session.apply1,·,·), which Session.fold IS by its own definition.
  lemma SessionIsChannel(h: Ch.HostState<Session.SessionState>, actions: seq<Session.SessionAction>)
    ensures Ch.AcceptFold(Session.apply1, h, actions).canonical == CC.Fold(Session.apply1, h.canonical, actions)
    ensures Ch.AcceptFold(Session.apply1, h, actions).nextSeq == h.nextSeq + |actions|
  {
    ChannelIsFold(Session.apply1, h, actions);
  }

  lemma ChatIsChannel(h: Ch.HostState<Chat.ChatState>, actions: seq<Chat.ChatAction>)
    ensures Ch.AcceptFold(Chat.apply1, h, actions).canonical == CC.Fold(Chat.apply1, h.canonical, actions)
    ensures Ch.AcceptFold(Chat.apply1, h, actions).nextSeq == h.nextSeq + |actions|
  {
    ChannelIsFold(Chat.apply1, h, actions);
    assert Ch.AcceptFold(Chat.apply1, h, actions).canonical == Chat.fold(h.canonical, actions);
  }

  lemma TerminalIsChannel(h: Ch.HostState<Terminal.TerminalState>, actions: seq<Terminal.TerminalAction>)
    ensures Ch.AcceptFold(Terminal.apply1, h, actions).canonical == CC.Fold(Terminal.apply1, h.canonical, actions)
    ensures Ch.AcceptFold(Terminal.apply1, h, actions).nextSeq == h.nextSeq + |actions|
  {
    ChannelIsFold(Terminal.apply1, h, actions);
    assert Ch.AcceptFold(Terminal.apply1, h, actions).canonical == Terminal.fold(h.canonical, actions);
  }

  lemma ChangesetIsChannel(h: Ch.HostState<Changeset.ChangesetState>, actions: seq<Changeset.ChangesetAction>)
    ensures Ch.AcceptFold(Changeset.apply1, h, actions).canonical == CC.Fold(Changeset.apply1, h.canonical, actions)
    ensures Ch.AcceptFold(Changeset.apply1, h, actions).nextSeq == h.nextSeq + |actions|
  {
    ChannelIsFold(Changeset.apply1, h, actions);
    assert Ch.AcceptFold(Changeset.apply1, h, actions).canonical == Changeset.fold(h.canonical, actions);
  }

  lemma AnnotationsIsChannel(h: Ch.HostState<Annotations.AnnotationsState>, actions: seq<Annotations.AnnotationsAction>)
    ensures Ch.AcceptFold(Annotations.apply1, h, actions).canonical == CC.Fold(Annotations.apply1, h.canonical, actions)
    ensures Ch.AcceptFold(Annotations.apply1, h, actions).nextSeq == h.nextSeq + |actions|
  {
    ChannelIsFold(Annotations.apply1, h, actions);
    assert Ch.AcceptFold(Annotations.apply1, h, actions).canonical == Annotations.fold(h.canonical, actions);
  }

  lemma ResourceWatchIsChannel(h: Ch.HostState<ResourceWatch.ResourceWatchState>, actions: seq<ResourceWatch.ResourceWatchAction>)
    ensures Ch.AcceptFold(ResourceWatch.apply1, h, actions).canonical == CC.Fold(ResourceWatch.apply1, h.canonical, actions)
    ensures Ch.AcceptFold(ResourceWatch.apply1, h, actions).nextSeq == h.nextSeq + |actions|
  {
    ChannelIsFold(ResourceWatch.apply1, h, actions);
    assert Ch.AcceptFold(ResourceWatch.apply1, h, actions).canonical == ResourceWatch.fold(h.canonical, actions);
  }

  lemma CanvasIsChannel(h: Ch.HostState<Canvas.CanvasState>, actions: seq<Canvas.CanvasAction>)
    ensures Ch.AcceptFold(Canvas.apply1, h, actions).canonical == CC.Fold(Canvas.apply1, h.canonical, actions)
    ensures Ch.AcceptFold(Canvas.apply1, h, actions).nextSeq == h.nextSeq + |actions|
  {
    ChannelIsFold(Canvas.apply1, h, actions);
    assert Ch.AcceptFold(Canvas.apply1, h, actions).canonical == Canvas.fold(h.canonical, actions);
  }

  // Root: FoldRoot threads a `now` through an inline lambda (ApplyToRoot ignores it), so the host canonical
  // is stated against CC.Fold(apply1,·,·) directly — the full HostCanonicalIsFold obligation for root. The
  // FoldRoot-at-now==9999 tie is carried in convergence.dfy's HostFoldIsCoreFoldRoot (unfolds FoldRoot's body).
  lemma RootIsChannel(h: Ch.HostState<RootState>, actions: seq<RootAction>)
    ensures Ch.AcceptFold(apply1, h, actions).canonical == CC.Fold(apply1, h.canonical, actions)
    ensures Ch.AcceptFold(apply1, h, actions).nextSeq == h.nextSeq + |actions|
  {
    ChannelIsFold(apply1, h, actions);
  }

  // ── Non-vacuity witness: a concrete root host folding a real 2-action stream gets seq exactly 2. ──
  // The canonical value itself is opaque through the .doo firewall (ApplyToRoot is `provides`, body hidden),
  // so the witness asserts what the firewall permits: the assigned seq is exactly 2 after two accepts.
  // Falsifiable — change the stream length and it fails.
  lemma ChannelWitness()
    ensures Ch.AcceptFold(apply1, Ch.HostState(RootState([], None, None, None), 0),
                          [RootActiveSessionsChanged(5), RootActiveSessionsChanged(7)]).nextSeq == 2
  {
    RootIsChannel(Ch.HostState(RootState([], None, None, None), 0),
                  [RootActiveSessionsChanged(5), RootActiveSessionsChanged(7)]);
  }
}
