// AHP Dafny client — the UNIFIED MULTI-CHANNEL layer.
//
// The eight channels (root + session + chat + terminal + changeset + annotations
// + resourceWatch + canvas) are each individually modelled, reduced, and codec-round-trip
// proven in their own files. THIS module composes all eight into ONE protocol:
//
//   • AhpAction  — the tagged union of every channel's Action.
//   • AhpState   — the product of every channel's State.
//   • applyAhp   — a pure, total, deterministic DISPATCH: it routes an AhpAction
//                  to the matching channel's `apply1`, updating ONLY that
//                  channel's component of the product state and leaving the other
//                  six components untouched. That per-channel isolation is exactly
//                  what makes a generic convergence proof (over applyAhp) cover a
//                  REAL mixed-channel session, not just the root channel.
//   • decodeAhpAction / encodeAhpAction — a unified codec that ROUTES by the
//                  wire "type" discriminator's channel prefix into the matching
//                  per-channel codec.
//   • encodeAhpState — serializes the whole product state (7 keys, one per
//                  channel) via each channel's own state serializer.
//
// It is a pure composition layer: it defines no new reducer or codec logic, only
// the dispatch/routing that ties the seven proven channels into one action space
// and one state space. No `assume`, no `{:axiom}`, no `requires false`.

include "ahp_skeleton.dfy"
include "fidelity.dfy"
include "session.dfy"
include "chat.dfy"
include "terminal.dfy"
include "changeset.dfy"
include "annotations.dfy"
include "resource_watch.dfy"
include "canvas.dfy"
include "codec/session_codec.dfy"
include "codec/chat_codec.dfy"
include "codec/terminal_codec.dfy"
include "codec/changeset_codec.dfy"
include "codec/annotations_codec.dfy"
include "codec/resource_watch_codec.dfy"
include "codec/canvas_codec.dfy"

module Ahp {
  // ── .doo FIREWALL: explicit export set (replaces Dafny's default export-all) ──
  // Consumers (convergence's ConvergesAhp/ReconvergesAhp + the transport demos)
  // use only the unified action/state datatypes, the pure dispatch reducer, and
  // the routing codec entry points — each called as a value/at runtime, never
  // unfolded in a proof, so the functions are `provides` (body hidden). The whole
  // round-trip proof machinery (the per-channel *Prefix/Route*/*RT lemmas, the
  // AhpAction/AhpStateRoundTrip theorems), decodeAhpState, the sub/typeOf/startsWith
  // helpers, the P_* prefix consts, typedAhp/wfAhp/WfAhpState, and the dispatch
  // isolation witnesses are all HIDDEN — no consumer needs them; the firewall bites.
  // `provides` on the imports (AhpSkeleton + the six per-channel aliases) is
  // required because the REVEALED AhpState/AhpAction fields name those modules'
  // types (RootState, S.SessionState, …); revealing a datatype forces its field
  // modules to be exported too. The codec imports (F/Sc/Cc/…) are only named in
  // the HIDDEN function bodies, so they need not be exported.
  export
    provides applyAhp, encodeAhpAction, decodeAhpAction, encodeAhpState
    provides WfAhpStateInv, WfAhpAction, WfAhpStatePreserved   // rung-4/6 aggregate: top-level reducer invariant + its preservation
    provides AhpSkeleton, S, Ch, T, Cs, An, RW, Cv
    provides Wire   // decodeAhpAction/encodeAhpAction/encodeAhpState's PROVIDED
                    // signatures name Wire.Json (param/return type); Wire.Json's owning
                    // module must be exported too (same rule as AhpSkeleton/S/…)
    reveals AhpState, AhpAction

  // Wire.Json + root types come from AhpSkeleton (opened → unqualified `apply1` is the
  // ROOT reducer). Every OTHER channel is aliased so its `apply1` stays qualified,
  // avoiding a name collision on the seven identically-named channel reducers.
  import opened AhpSkeleton
  import Wire = ConfluxCodec   // Wire.Json/Wire.JNull/Wire.JObj/... (no longer re-exported by AhpSkeleton)
  import CC = ConfluxContract
  import S  = Session
  import Ch = Chat
  import T  = Terminal
  import Cs = Changeset
  import An = Annotations
  import RW = ResourceWatch
  import Cv = Canvas

  // Codecs: Fidelity carries BOTH the shared wire accessors (field/asStr/…) AND
  // the root Wire.Json⇄Action codec; each channel has its own X-codec module.
  import F   = Fidelity
  import Sc  = SessionCodec
  import Cc  = ChatCodec
  import Tc  = TerminalCodec
  import Csc = ChangesetCodec
  import Anc = AnnotationsCodec
  import RWc = ResourceWatchCodec
  import Cvc = CanvasCodec

  // ── (1) AhpAction — the tagged union of every channel's Action ──
  datatype AhpAction =
    | ARoot(rootAction: RootAction)
    | ASession(sessionAction: S.SessionAction)
    | AChat(chatAction: Ch.ChatAction)
    | ATerminal(terminalAction: T.TerminalAction)
    | AChangeset(changesetAction: Cs.ChangesetAction)
    | AAnnotations(annotationsAction: An.AnnotationsAction)
    | AResourceWatch(resourceWatchAction: RW.ResourceWatchAction)
    | ACanvas(canvasAction: Cv.CanvasAction)

  // ── (2) AhpState — the product of every channel's State ──
  datatype AhpState = AhpState(
    root:          RootState,
    session:       S.SessionState,
    chat:          Ch.ChatState,
    terminal:      T.TerminalState,
    changeset:     Cs.ChangesetState,
    annotations:   An.AnnotationsState,
    resourceWatch: RW.ResourceWatchState,
    canvas:        Cv.CanvasState)

  // ── (3) applyAhp — pure, total, deterministic DISPATCH ──
  // Each variant updates ONLY its own component (the other six are left equal).
  // This genuinely calls the matching channel's proven `apply1`; it is not a stub.
  function applyAhp(s: AhpState, a: AhpAction): AhpState
  {
    match a
    case ARoot(ra)          => s.(root          := apply1(s.root, ra))               // AhpSkeleton.apply1 (root)
    case ASession(sa)       => s.(session       := S.apply1(s.session, sa))
    case AChat(ca)          => s.(chat          := Ch.apply1(s.chat, ca))
    case ATerminal(ta)      => s.(terminal      := T.apply1(s.terminal, ta))
    case AChangeset(csa)    => s.(changeset     := Cs.apply1(s.changeset, csa))
    case AAnnotations(ana)  => s.(annotations   := An.apply1(s.annotations, ana))
    case AResourceWatch(ra) => s.(resourceWatch := RW.apply1(s.resourceWatch, ra))
    case ACanvas(ca)        => s.(canvas        := Cv.apply1(s.canvas, ca))
  }

  function foldAhp(s: AhpState, actions: seq<AhpAction>): AhpState
  { CC.Fold(applyAhp, s, actions) }

  // ── prefix helper: is `prefix` a prefix of `s`? (pure, total) ──
  predicate startsWith(s: string, prefix: string)
  {
    |prefix| <= |s| && s[..|prefix|] == prefix
  }

  // convenience: the wire "type" discriminator string of a Wire.Json object.
  function typeOf(j: Wire.Json): string { F.asStr(F.field(j, "type")) }

  // channel prefixes (the exact strings the per-channel fixtures/codecs emit).
  const P_ROOT        : string := "root/"
  const P_SESSION     : string := "session/"
  const P_CHAT        : string := "chat/"
  const P_TERMINAL    : string := "terminal/"
  const P_CHANGESET   : string := "changeset/"
  const P_ANNOTATIONS : string := "annotations/"
  const P_RESWATCH    : string := "resourceWatch/"
  const P_RESWATCH_ALT: string := "resource-watch/"   // tolerated kebab spelling
  const P_CANVAS      : string := "canvas/"

  // ── (4) decodeAhpAction — ROUTE by the "type" discriminator's channel prefix ──
  // The matching per-channel codec then decodes the payload. The fall-through is a
  // TOTAL safe default: an unrecognized-channel envelope becomes ARoot(RootUnknown),
  // which the root reducer treats as a no-op (never a partial function).
  function decodeAhpAction(j: Wire.Json): AhpAction
  {
    var t := typeOf(j);
    if      startsWith(t, P_ROOT)        then ARoot(F.DecodeRootAction(j))
    else if startsWith(t, P_SESSION)     then ASession(Sc.DecodeSessionAction(j))
    else if startsWith(t, P_CHAT)        then AChat(Cc.DecodeChatAction(j))
    else if startsWith(t, P_TERMINAL)    then ATerminal(Tc.DecodeTerminalAction(j))
    else if startsWith(t, P_CHANGESET)   then AChangeset(Csc.DecodeChangesetAction(j))
    else if startsWith(t, P_ANNOTATIONS) then AAnnotations(Anc.DecodeAnnotationsAction(j))
    else if startsWith(t, P_RESWATCH) || startsWith(t, P_RESWATCH_ALT)
                                         then AResourceWatch(RWc.DecodeResourceWatchAction(j))
    else if startsWith(t, P_CANVAS)      then ACanvas(Cvc.DecodeCanvasAction(j))
    else ARoot(RootUnknown(j))
  }

  // ── (5) encodeAhpAction — dispatch to the matching channel's Encode ──
  function encodeAhpAction(a: AhpAction): Wire.Json
  {
    match a
    case ARoot(ra)          => F.EncodeRootAction(ra)
    case ASession(sa)       => Sc.EncodeSessionAction(sa)
    case AChat(ca)          => Cc.EncodeChatAction(ca)
    case ATerminal(ta)      => Tc.EncodeTerminalAction(ta)
    case AChangeset(csa)    => Csc.EncodeChangesetAction(csa)
    case AAnnotations(ana)  => Anc.EncodeAnnotationsAction(ana)
    case AResourceWatch(ra) => RWc.EncodeResourceWatchAction(ra)
    case ACanvas(ca)        => Cvc.EncodeCanvasAction(ca)
  }

  // ── (6) encodeAhpState — 7-keyed product serialization ──
  function encodeAhpState(s: AhpState): Wire.Json
  {
    Wire.JObj(map[
      "root"          := F.encodeRootState(s.root),
      "session"       := Sc.encodeSessionState(s.session),
      "chat"          := Cc.encodeChatState(s.chat),
      "terminal"      := Tc.encodeTerminalState(s.terminal),
      "changeset"     := Csc.encodeChangesetState(s.changeset),
      "annotations"   := Anc.encodeAnnotationsState(s.annotations),
      "resourceWatch" := RWc.encodeResourceWatchState(s.resourceWatch),
      "canvas"        := Cvc.encodeCanvasState(s.canvas)
    ])
  }

  // ── dispatch isolation witness (non-vacuity): applyAhp genuinely updates ONLY
  // the addressed component. This proves the dispatch is real, not a stub that
  // ignores the action or overwrites unrelated components. ──
  lemma ApplyAhp_SessionTouchesOnlySession(s: AhpState, sa: S.SessionAction)
    ensures applyAhp(s, ASession(sa)).session       == S.apply1(s.session, sa)
    ensures applyAhp(s, ASession(sa)).root          == s.root
    ensures applyAhp(s, ASession(sa)).chat          == s.chat
    ensures applyAhp(s, ASession(sa)).terminal      == s.terminal
    ensures applyAhp(s, ASession(sa)).changeset     == s.changeset
    ensures applyAhp(s, ASession(sa)).annotations   == s.annotations
    ensures applyAhp(s, ASession(sa)).resourceWatch == s.resourceWatch
    ensures applyAhp(s, ASession(sa)).canvas        == s.canvas
  {}
  lemma ApplyAhp_RootTouchesOnlyRoot(s: AhpState, ra: RootAction)
    ensures applyAhp(s, ARoot(ra)).root          == apply1(s.root, ra)
    ensures applyAhp(s, ARoot(ra)).session       == s.session
    ensures applyAhp(s, ARoot(ra)).chat          == s.chat
    ensures applyAhp(s, ARoot(ra)).terminal      == s.terminal
    ensures applyAhp(s, ARoot(ra)).changeset     == s.changeset
    ensures applyAhp(s, ARoot(ra)).annotations   == s.annotations
    ensures applyAhp(s, ARoot(ra)).resourceWatch == s.resourceWatch
    ensures applyAhp(s, ARoot(ra)).canvas        == s.canvas
  {}

  // ════════════════════════════════════════════════════════════════════════
  //  (7) UNIFIED ACTION ROUND-TRIP — decodeAhpAction ∘ encodeAhpAction == id
  //      on the TYPED variants of every channel.
  //
  //  Structure (per channel X, mirroring the per-channel round-trip carve-outs):
  //    • XPrefix  — the encoded wire "type" for any typed X-action starts with
  //                 X's channel prefix (and carries the distinguishing chars the
  //                 router needs to reject every EARLIER prefix in decodeAhpAction's
  //                 if-else chain). Proven by a match that unfolds EncodeX to its
  //                 concrete "channel/…" literal per variant.
  //    • RouteX   — GIVEN those distinguishing chars, decodeAhpAction routes j to
  //                 AX(DecodeX(j)). Pure character reasoning; no variant match.
  //    • XRT      — composes X's proven codec round-trip (DecodeX∘EncodeX==id) with
  //                 XPrefix + RouteX to close decodeAhpAction∘encodeAhpAction==id
  //                 for that channel.
  //
  //  The Unknown variant of each channel is EXCLUDED (its encode is the verbatim
  //  raw payload, which may itself be a typed-shaped object that the router would
  //  re-classify) — exactly the same carve-out each per-channel codec makes.
  //  Session/Chat additionally carry their codecs' well-formedness preconditions
  //  (WfSessionAction / WFPart / WFTurns / okJson) verbatim; they are genuine
  //  wire-fidelity constraints of those channels, not weakenings introduced here.
  //
  //  Router prefix order (root, session, chat, terminal, changeset, annotations,
  //  resourceWatch). No prefix is a prefix of another; the distinguishing chars:
  //    session      → [0]=='s'          (vs root 'r')
  //    chat         → [0]=='c'          (vs root 'r', session 's')
  //    terminal     → [0]=='t'          (vs r/s/c)
  //    changeset    → [0]=='c', [3]=='n'(vs chat, whose [3]=='t')
  //    annotations  → [0]=='a'          (vs r/s/c/t and changeset 'c')
  //    resourceWatch→ [0]=='r', [1]=='e'(vs root, whose [1]=='o')
  // ════════════════════════════════════════════════════════════════════════

  // ── ROOT (first prefix in the chain — no earlier prefix to reject) ──
  lemma RootPrefix(a: RootAction)
    requires !a.RootUnknown?
    ensures startsWith(typeOf(F.EncodeRootAction(a)), P_ROOT)
  {
    match a
    case RootAgentsChanged(_)         => case RootActiveSessionsChanged(_) =>
    case RootTerminalsChanged(_)      => case RootConfigChanged(_, _)      =>
    case RootUnknown(_)               =>
  }
  lemma RouteRoot(j: Wire.Json)
    requires startsWith(typeOf(j), P_ROOT)
    ensures decodeAhpAction(j) == ARoot(F.DecodeRootAction(j))
  {}
  lemma RootRT(a: RootAction)
    requires !a.RootUnknown?
    ensures decodeAhpAction(encodeAhpAction(ARoot(a))) == ARoot(a)
  { F.RootActionRoundTrip(a); RootPrefix(a); RouteRoot(encodeAhpAction(ARoot(a))); }

  // ── SESSION ([0]=='s') ──
  lemma {:isolate_assertions} SessionPrefix(a: S.SessionAction)
    requires !a.SUnknown?
    ensures var t := typeOf(Sc.EncodeSessionAction(a)); |t| >= 1 && t[0] == 's' && startsWith(t, P_SESSION)
  {
    match a
    case Ready => case CreationFailed(_) => case TitleChanged(_) => case ServerToolsChanged(_) =>
    case MetaChanged(_) => case IsReadChanged(_) => case IsArchivedChanged(_) => case ActivityChanged(_) =>
    case ConfigChanged(_, _) => case CustomizationsChanged(_) => case CustomizationToggled(_, _) =>
    case CustomizationRemoved(_) => case DefaultChatChanged(_) => case ChatAdded(_) => case ChatRemoved(_) =>
    case ChatUpdated(_, _, _, _) => case ChangesetsChanged(_) => case ActiveClientSet(_) =>
    case CanvasesChanged(_) => case OpenCanvasesChanged(_) =>
    case ActiveClientRemoved(_) => case InputNeededSet(_) => case InputNeededRemoved(_) =>
    case CustomizationUpdated(_) => case McpServerStateChanged(_, _, _) => case McpServerStartRequested(_) =>
    case McpServerStopRequested(_) => case SUnknown(_) =>
  }
  lemma RouteSession(j: Wire.Json)
    requires startsWith(typeOf(j), P_SESSION)
    requires |typeOf(j)| >= 1 && typeOf(j)[0] == 's'
    ensures decodeAhpAction(j) == ASession(Sc.DecodeSessionAction(j))
  {}
  lemma SessionRT(a: S.SessionAction)
    requires !a.SUnknown?
    requires Sc.WfSessionAction(a)
    ensures decodeAhpAction(encodeAhpAction(ASession(a))) == ASession(a)
  { Sc.SessionActionRoundTrip(a); SessionPrefix(a); RouteSession(encodeAhpAction(ASession(a))); }

  // ── CHAT ([0]=='c'; earlier prefixes root/session both differ at [0]) ──
  // NOTE: no {:isolate_assertions} here (unlike Session/other prefix lemmas). Once ChatCodec's
  // .doo firewall REVEALS EncodeChatAction via an explicit export, a per-case (isolate_assertions)
  // VC cannot unfold the 25-arm EncodeChatAction across the reveal boundary — the monolithic
  // whole-lemma VC can. Do not re-add {:isolate_assertions} unless ChatCodec drops its export.
  lemma ChatPrefix(a: Ch.ChatAction)
    requires !a.CUnknown?
    ensures var t := typeOf(Cc.EncodeChatAction(a)); |t| >= 1 && t[0] == 'c' && startsWith(t, P_CHAT)
  {
    match a
    case CDraftChanged(_) => case CActivityChanged(_) => case CTurnStarted(_, _, _) => case CTurnComplete(_) =>
    case CTurnCancelled(_) => case CUsage(_, _) => case CError(_, _) => case CResponsePart(_, _) =>
    case CDelta(_, _, _) => case CReasoning(_, _, _) => case CToolCallStart(_, _, _, _, _, _, _) =>
    case CToolCallDelta(_, _, _, _, _) => case CToolCallReady(_, _, _, _, _, _, _, _, _, _, _) => case CToolCallConfirmed(_, _, _, _, _, _, _, _, _, _) =>
    case CToolCallAuthRequired(_, _, _, _) => case CToolCallAuthResolved(_, _, _) => case CToolCallComplete(_, _, _, _, _, _, _, _, _) => case CToolCallResultConfirmed(_, _, _, _) => case CToolCallContentChanged(_, _, _, _) =>
    case CTruncated(_) => case CQueuedReordered(_) => case CTurnsLoaded(_, _) => case CInputRequested(_) =>
    case CInputAnswerChanged(_, _, _) => case CInputCompleted(_, _, _) => case CPendingMessageSet(_, _, _) =>
    case CPendingMessageRemoved(_, _) => case CUnknown(_) =>
  }
  lemma RouteChat(j: Wire.Json)
    requires startsWith(typeOf(j), P_CHAT)
    requires |typeOf(j)| >= 1 && typeOf(j)[0] == 'c'
    ensures decodeAhpAction(j) == AChat(Cc.DecodeChatAction(j))
  {}
  lemma ChatRT(a: Ch.ChatAction)
    requires !a.CUnknown?
    requires Cc.WFChatAction(a)
    requires a.CResponsePart?       ==> Cc.WFPart(a.part)
    requires a.CTurnsLoaded?        ==> Cc.WFTurns(a.loaded)
    requires a.CInputAnswerChanged? ==> Cc.okJson(a.answer)
    ensures decodeAhpAction(encodeAhpAction(AChat(a))) == AChat(a)
  { Cc.ChatActionRoundTrip(a); ChatPrefix(a); RouteChat(encodeAhpAction(AChat(a))); }

  // ── TERMINAL ([0]=='t') ──
  lemma TerminalPrefix(a: T.TerminalAction)
    requires !a.TUnknown?
    ensures var t := typeOf(Tc.EncodeTerminalAction(a)); |t| >= 1 && t[0] == 't' && startsWith(t, P_TERMINAL)
  {
    match a
    case TCwdChanged(_) => case TTitleChanged(_) => case TResized(_, _) => case TExited(_) =>
    case TData(_) => case TCleared => case TClaimed(_) => case TCommandDetectionAvailable =>
    case TCommandExecuted(_, _, _) => case TCommandFinished(_, _, _) => case TInput => case TUnknown(_) =>
  }
  lemma RouteTerminal(j: Wire.Json)
    requires startsWith(typeOf(j), P_TERMINAL)
    requires |typeOf(j)| >= 1 && typeOf(j)[0] == 't'
    ensures decodeAhpAction(j) == ATerminal(Tc.DecodeTerminalAction(j))
  {}
  lemma TerminalRT(a: T.TerminalAction)
    requires !a.TUnknown?
    ensures decodeAhpAction(encodeAhpAction(ATerminal(a))) == ATerminal(a)
  { Tc.TerminalActionRoundTrip(a); TerminalPrefix(a); RouteTerminal(encodeAhpAction(ATerminal(a))); }

  // ── CHANGESET ([0]=='c' AND [3]=='n' — to reject chat, whose [3]=='t') ──
  lemma ChangesetPrefix(a: Cs.ChangesetAction)
    requires !a.CsUnknown?
    ensures var t := typeOf(Csc.EncodeChangesetAction(a)); |t| >= 4 && t[0] == 'c' && t[3] == 'n' && startsWith(t, P_CHANGESET)
  {
    match a
    case StatusChanged(_, _) => case FileSet(_) => case FileRemoved(_) => case OperationsChanged(_) =>
    case Cleared => case OperationStatusChanged(_, _, _) => case ContentChanged(_, _, _) =>
    case FilesReviewedChanged(_, _) => case CsUnknown(_) =>
  }
  lemma RouteChangeset(j: Wire.Json)
    requires startsWith(typeOf(j), P_CHANGESET)
    requires |typeOf(j)| >= 4 && typeOf(j)[0] == 'c' && typeOf(j)[3] == 'n'
    ensures decodeAhpAction(j) == AChangeset(Csc.DecodeChangesetAction(j))
  {}
  lemma ChangesetRT(a: Cs.ChangesetAction)
    requires !a.CsUnknown?
    ensures decodeAhpAction(encodeAhpAction(AChangeset(a))) == AChangeset(a)
  { Csc.ChangesetActionRoundTrip(a); ChangesetPrefix(a); RouteChangeset(encodeAhpAction(AChangeset(a))); }

  // ── ANNOTATIONS ([0]=='a') ──
  lemma AnnotationsPrefix(a: An.AnnotationsAction)
    requires !a.AnUnknown?
    ensures var t := typeOf(Anc.EncodeAnnotationsAction(a)); |t| >= 1 && t[0] == 'a' && startsWith(t, P_ANNOTATIONS)
  {
    match a
    case Set(_) => case Removed(_) => case EntrySet(_, _) => case EntryRemoved(_, _) =>
    case Updated(_, _, _, _, _) => case AnUnknown(_) =>
  }
  lemma RouteAnnotations(j: Wire.Json)
    requires startsWith(typeOf(j), P_ANNOTATIONS)
    requires |typeOf(j)| >= 1 && typeOf(j)[0] == 'a'
    ensures decodeAhpAction(j) == AAnnotations(Anc.DecodeAnnotationsAction(j))
  {}
  lemma AnnotationsRT(a: An.AnnotationsAction)
    requires !a.AnUnknown?
    ensures decodeAhpAction(encodeAhpAction(AAnnotations(a))) == AAnnotations(a)
  { Anc.AnnotationsActionRoundTrip(a); AnnotationsPrefix(a); RouteAnnotations(encodeAhpAction(AAnnotations(a))); }

  // ── RESOURCEWATCH ([0]=='r' AND [1]=='e' — to reject root, whose [1]=='o') ──
  lemma ResourceWatchPrefix(a: RW.ResourceWatchAction)
    requires !a.RWUnknown?
    ensures var t := typeOf(RWc.EncodeResourceWatchAction(a)); |t| >= 2 && t[0] == 'r' && t[1] == 'e' && startsWith(t, P_RESWATCH)
  {
    match a
    case RWChanged(_) => case RWUnknown(_) =>
  }
  lemma RouteResourceWatch(j: Wire.Json)
    requires startsWith(typeOf(j), P_RESWATCH)
    requires |typeOf(j)| >= 2 && typeOf(j)[0] == 'r' && typeOf(j)[1] == 'e'
    ensures decodeAhpAction(j) == AResourceWatch(RWc.DecodeResourceWatchAction(j))
  {}
  lemma ResourceWatchRT(a: RW.ResourceWatchAction)
    requires !a.RWUnknown?
    ensures decodeAhpAction(encodeAhpAction(AResourceWatch(a))) == AResourceWatch(a)
  { RWc.ResourceWatchActionRoundTrip(a); ResourceWatchPrefix(a); RouteResourceWatch(encodeAhpAction(AResourceWatch(a))); }

  // ── CANVAS ([0]=='c' AND [1]=='a' — rejects chat/changeset, both "ch…") ──
  lemma CanvasPrefix(a: Cv.CanvasAction)
    requires !a.CanvasUnknown?
    ensures var t := typeOf(Cvc.EncodeCanvasAction(a));
      |t| >= 2 && t[0] == 'c' && t[1] == 'a' && startsWith(t, P_CANVAS)
  {
    match a
    case Updated(_, _, _, _) => case CloseRequested => case CanvasUnknown(_) =>
  }
  lemma RouteCanvas(j: Wire.Json)
    requires startsWith(typeOf(j), P_CANVAS)
    requires |typeOf(j)| >= 2 && typeOf(j)[0] == 'c' && typeOf(j)[1] == 'a'
    ensures decodeAhpAction(j) == ACanvas(Cvc.DecodeCanvasAction(j))
  {}
  lemma CanvasRT(a: Cv.CanvasAction)
    requires !a.CanvasUnknown?
    ensures decodeAhpAction(encodeAhpAction(ACanvas(a))) == ACanvas(a)
  { Cvc.CanvasActionRoundTrip(a); CanvasPrefix(a); RouteCanvas(encodeAhpAction(ACanvas(a))); }

  // ── typed / well-formed predicates over the UNIFIED action ──
  // typedAhp: the action is a typed (non-Unknown) variant of its channel — the
  // exact per-channel exclusion the codec round-trips require.
  predicate typedAhp(a: AhpAction) {
    match a
    case ARoot(ra)          => !ra.RootUnknown?
    case ASession(sa)       => !sa.SUnknown?
    case AChat(ca)          => !ca.CUnknown?
    case ATerminal(ta)      => !ta.TUnknown?
    case AChangeset(csa)    => !csa.CsUnknown?
    case AAnnotations(ana)  => !ana.AnUnknown?
    case AResourceWatch(ra) => !ra.RWUnknown?
    case ACanvas(ca)        => !ca.CanvasUnknown?
  }
  // wfAhp: the session/chat wire-fidelity well-formedness constraints their
  // codecs impose (verbatim); every other channel imposes none.
  predicate wfAhp(a: AhpAction) {
    match a
    case ASession(sa) => Sc.WfSessionAction(sa)
    case AChat(ca)    =>
      Cc.WFChatAction(ca) &&
      (ca.CResponsePart?       ==> Cc.WFPart(ca.part)) &&
      (ca.CTurnsLoaded?        ==> Cc.WFTurns(ca.loaded)) &&
      (ca.CInputAnswerChanged? ==> Cc.okJson(ca.answer))
    case _ => true
  }

  // ── THE UNIFIED ROUND-TRIP: decode ∘ encode == id on every typed action ──
  // Json→domain→Json fixpoint (unified dispatch): the canonical wire form re-decodes + re-encodes to ITSELF —
  // the direction the conformance suite exercises, now machine-checked. Follows from the direction-1 round-trip.
  lemma AhpWireCanonicalRoundTrip(a: AhpAction)
    requires typedAhp(a)
    requires wfAhp(a)
    ensures encodeAhpAction(decodeAhpAction(encodeAhpAction(a))) == encodeAhpAction(a)
  { AhpActionRoundTrip(a); }
  lemma AhpActionRoundTrip(a: AhpAction)
    requires typedAhp(a)
    requires wfAhp(a)
    ensures decodeAhpAction(encodeAhpAction(a)) == a
  {
    match a
    case ARoot(ra)          => RootRT(ra);
    case ASession(sa)       => SessionRT(sa);
    case AChat(ca)          => ChatRT(ca);
    case ATerminal(ta)      => TerminalRT(ta);
    case AChangeset(csa)    => ChangesetRT(csa);
    case AAnnotations(ana)  => AnnotationsRT(ana);
    case AResourceWatch(ra) => ResourceWatchRT(ra);
    case ACanvas(ca)        => CanvasRT(ca);
  }

  // Non-vacuity witness: the (typedAhp ∧ wfAhp) precondition is INHABITED by a
  // genuine mixed-channel-representative action, so AhpActionRoundTrip is NOT
  // vacuously true. Uses a terminal action (a non-Unknown typed variant that
  // exercises real routing past root/session/chat in the if-else chain).
  lemma AhpActionRoundTrip_NonVacuous()
    ensures exists a: AhpAction :: typedAhp(a) && wfAhp(a) && decodeAhpAction(encodeAhpAction(a)) == a
  {
    var a := ATerminal(T.TCwdChanged("/tmp"));
    assert typedAhp(a) && wfAhp(a);
    AhpActionRoundTrip(a);
    assert decodeAhpAction(encodeAhpAction(a)) == a;
  }

  // ════════════════════════════════════════════════════════════════════════
  //  (8) UNIFIED STATE ROUND-TRIP — decodeAhpState ∘ encodeAhpState == id.
  //
  //  The STATE-side analogue of AhpActionRoundTrip: it is the COMPOSITION of all
  //  SEVEN per-channel STATE round-trips. encodeAhpState writes one key per
  //  channel (root/session/chat/terminal/changeset/annotations/resourceWatch),
  //  each = that channel's own encodeXState; decodeAhpState reads each key back
  //  and reconstructs via that channel's own decodeXState. Because the 7 keys are
  //  distinct, each `sub(encodeAhpState(s), "<chan>")` recovers exactly the
  //  sub-object that channel's encoder wrote, so invoking each channel's proven
  //  XStateRoundTrip closes the composite decode∘encode==id.
  //
  //  Preconditions: only session + chat carry one (WfSessionState / WfChatState —
  //  the verbatim wire-fidelity constraints of those two codecs). The other five
  //  channels' state round-trips are UNCONDITIONAL, so WfAhpState omits them.
  // ════════════════════════════════════════════════════════════════════════

  // safe field-unwrap: reads key k from a Wire.Json object; Wire.JNull if absent/non-object.
  // Total (a decode function must be total); on the round-trip input each of the
  // 7 keys is present, so this returns exactly the sub-object the encoder wrote.
  function sub(j: Wire.Json, k: string): Wire.Json
  { if F.field(j, k).Some? then F.field(j, k).value else Wire.JNull }

  // decodeAhpState — the TRUE inverse of encodeAhpState. Reads the SAME 7 keys
  // encodeAhpState writes and reconstructs each component via its channel decoder.
  function decodeAhpState(j: Wire.Json): AhpState
  {
    AhpState(
      root          := F.decodeRootState(sub(j, "root")),
      session       := Sc.decodeSessionState(sub(j, "session")),
      chat          := Cc.decodeChatState(sub(j, "chat")),
      terminal      := Tc.decodeTerminalState(sub(j, "terminal")),
      changeset     := Csc.decodeChangesetState(sub(j, "changeset")),
      annotations   := Anc.decodeAnnotationsState(sub(j, "annotations")),
      resourceWatch := RWc.decodeResourceWatchState(sub(j, "resourceWatch")),
      canvas        := Cvc.decodeCanvasState(sub(j, "canvas")))
  }

  // WfAhpState — the conjunction of each component's state-codec precondition.
  // Only session + chat impose one; the other five are unconditional.
  predicate WfAhpState(s: AhpState)
  { Sc.WfSessionState(s.session) && Cc.WfChatState(s.chat) }

  // THE UNIFIED STATE PROPERTY: every well-formed AhpState serializes to a wire
  // object that decodes back to itself — the COMPOSITION of the 7 channel RTs.
  // {:isolate_assertions} splits the VC so no single query approaches the 30s
  // limit (AhpState is the 7-way product; chat's own RT already needs the split).
  lemma {:isolate_assertions} AhpStateRoundTrip(s: AhpState)
    requires WfAhpState(s)
    ensures decodeAhpState(encodeAhpState(s)) == s
  {
    var j := encodeAhpState(s);
    // Each of the 7 distinct keys recovers exactly the sub-object its encoder wrote.
    assert sub(j, "root")          == F.encodeRootState(s.root);
    assert sub(j, "session")       == Sc.encodeSessionState(s.session);
    assert sub(j, "chat")          == Cc.encodeChatState(s.chat);
    assert sub(j, "terminal")      == Tc.encodeTerminalState(s.terminal);
    assert sub(j, "changeset")     == Csc.encodeChangesetState(s.changeset);
    assert sub(j, "annotations")   == Anc.encodeAnnotationsState(s.annotations);
    assert sub(j, "resourceWatch") == RWc.encodeResourceWatchState(s.resourceWatch);
    assert sub(j, "canvas")        == Cvc.encodeCanvasState(s.canvas);
    // Compose the 7 proven per-channel STATE round-trips — one per component.
    F.RootStateRoundTrip(s.root);
    Sc.SessionStateRoundTrip(s.session);
    Cc.ChatStateRoundTrip(s.chat);
    Tc.TerminalStateRoundTrip(s.terminal);
    Csc.ChangesetStateRoundTrip(s.changeset);
    Anc.AnnotationsStateRoundTrip(s.annotations);
    RWc.ResourceWatchStateRoundTrip(s.resourceWatch);
    Cvc.CanvasStateRoundTrip(s.canvas);
  }

  // Non-vacuity witness: a GENUINELY non-empty AhpState (session + chat both
  // carry real content, resourceWatch is populated) satisfies WfAhpState and
  // round-trips — so AhpStateRoundTrip is NOT vacuously true / empty-state-only.
  lemma AhpStateRoundTrip_NonVacuous()
    ensures exists s: AhpState :: WfAhpState(s) && decodeAhpState(encodeAhpState(s)) == s
  {
    // — a real, well-formed SessionState (mirrors SessionCodec's witness) —
    var c := S.Cust("mcp-1", "mcpServer", "file:///w", "MCP", Some(true), None, None);
    var sess := S.SessionState(
      "prov", "Title", 72, "active",
      Some("thinking"),
      Some(S.SConfig(Wire.JNull, map["k" := Wire.JStr("v")])),
      None, None, None, None, None, None,
      Some("chat-1"),
      [], [S.Chat("res", "t", 1, Some("a"), "1970-01-01T00:00:02.000Z", Wire.JNull)], [c], []);
    assert Sc.WfSessionState(sess);
    // — a real, well-formed ChatState (mirrors ChatCodec's witness) —
    var t := Ch.Turn("turn-1", Wire.JStr("hello"), [Ch.PMarkdown("p1", "hi there")], "completed", None, None);
    var chat := Ch.ChatState(
      [t], "My Chat", Ch.GEN, "2026-01-01T00:00:00Z",
      Some("draft text"), None, Some(t),
      None, [Ch.QMsg("q1", Wire.JStr("queued msg"))], Some("cursor-1"));   // #338: no inputRequests field
    assert Cc.WFTurns(chat.turns) by { forall i | 0 <= i < |chat.turns| ensures Cc.WFTurn(chat.turns[i]) { assert chat.turns[i] == t; } }
    assert Cc.WfChatState(chat);
    // — the composite state: session + chat populated, resourceWatch populated —
    var s := AhpState(
      root          := RootState([], None, None, None),
      session       := sess,
      chat          := chat,
      terminal      := T.TerminalState("term", None, None, None, [], None, None, None),
      changeset     := Cs.ChangesetState("idle", [], None, None),
      annotations   := An.AnnotationsState([]),
      resourceWatch := RW.ResourceWatchState("/workspace", true),
      canvas        := Cv.CanvasState("markdown", "opaque-provider", None,
                         Some("README.md"), None, Some("ahp-session:/content/1"), Cv.Ready));
    assert WfAhpState(s);
    AhpStateRoundTrip(s);
    assert WfAhpState(s) && decodeAhpState(encodeAhpState(s)) == s;
  }

  // ════════════════════════════════════════════════════════════════════════
  //  RUNG 4 / 6 (AGGREGATE) — the top-level REDUCER INVARIANT over the whole
  //  8-channel product, and its per-channel-dispatch PRESERVATION theorem.
  //  (docs/shapes/04-invariant-bearing-state, 06-invariant-preserving-reducer.)
  //
  //  DELIBERATELY DISTINCT from WfAhpState above. WfAhpState is the codec
  //  round-trip PRECONDITION — a wire-fidelity fact the DECODER establishes (e.g.
  //  session's WfClients/WfInputs: each stored raw's embedded id matches its
  //  projected key) which the base reducers deliberately do NOT import and do NOT
  //  maintain, so it is not a reducer invariant. WfAhpStateInv is the REDUCER
  //  invariant: the keyed-collection uniqueness + no-stored-null facts each
  //  channel's ApplyToX provably preserves. It is therefore built from each
  //  channel's REDUCER-side predicate (Session.SessionWf, Chat.WfChatState,
  //  Changeset.WfChangesetState, Annotations.WfAnnotationsState) — NOT the
  //  codec-side ones (Sc.WfSessionState / Cc.WfChatState). Four channels carry a
  //  reducer invariant; root/terminal/resourceWatch/canvas carry none (their
  //  reducers are scalar / opaque-Json with nothing to keep unique), so they
  //  contribute `true` and are trivially preserved.
  // ════════════════════════════════════════════════════════════════════════
  ghost predicate WfAhpStateInv(s: AhpState) {
    S.SessionWf(s.session) &&
    Ch.WfChatState(s.chat) &&
    Cs.WfChangesetState(s.changeset) &&
    An.WfAnnotationsState(s.annotations)
  }

  // The aggregate action precondition: each channel arm carries EXACTLY its own
  // channel's rung-7 premise (the whole-seq-install arms that need incoming
  // uniqueness / non-null); every other channel contributes `true`.
  ghost predicate WfAhpAction(a: AhpAction) {
    match a
    case ASession(sa)      => S.WfSessionActionInv(sa)
    case AChat(ca)         => Ch.WfChatAction(ca)
    case AChangeset(csa)   => Cs.WfChangesetAction(csa)
    case AAnnotations(ana) => An.WfAnnotationsAction(ana)
    case _                 => true
  }

  // rung-6 AGGREGATE: applyAhp preserves the whole-product invariant. applyAhp is a
  // pure dispatch that updates ONLY the addressed component (the datatype update
  // leaves the other seven fields definitionally equal — exactly what the
  // ApplyAhp_*TouchesOnly* isolation lemmas witness), so each arm discharges the
  // touched channel's invariant by CITING that channel's apply1-level preservation
  // lemma and carries the other three channels' invariants straight through.
  lemma WfAhpStatePreserved(s: AhpState, a: AhpAction)
    requires WfAhpStateInv(s) && WfAhpAction(a)
    ensures  WfAhpStateInv(applyAhp(s, a))
  {
    match a
    case ARoot(ra)          =>                                                    // touches root only
    case ASession(sa)       => S.Apply1PreservesSessionWf(s.session, sa);
    case AChat(ca)          => Ch.Apply1PreservesWfChat(s.chat, ca);
    case ATerminal(ta)      =>                                                    // touches terminal only
    case AChangeset(csa)    => Cs.Apply1PreservesWfChangeset(s.changeset, csa);
    case AAnnotations(ana)  => An.Apply1PreservesWfAnnotations(s.annotations, ana);
    case AResourceWatch(ra) =>                                                    // touches resourceWatch only
    case ACanvas(ca)        =>                                                    // touches canvas only
  }

  // Non-vacuity witness: a GENUINELY populated well-formed product (a real keyed
  // customization + a chat in the session component) whose image under a keyed
  // session arm still satisfies the whole-product invariant — so
  // WfAhpStatePreserved is not vacuously true / empty-state-only.
  lemma WfAhpStatePreserved_NonVacuous()
    ensures exists s: AhpState, a: AhpAction ::
      WfAhpStateInv(s) && WfAhpAction(a) && WfAhpStateInv(applyAhp(s, a))
  {
    // Each channel hands back a concrete well-formed component (+ session a keyed action), so we never unfold
    // the channels' exported-opaque reducer predicates here.
    var sess, sa := S.SessionWfWitness();   // SessionWf(sess) ∧ WfSessionActionInv(sa)
    var chat := Ch.ChatWfWitness();         // WfChatState(chat)
    var cs   := Cs.ChangesetWfWitness();    // WfChangesetState(cs)
    var an   := An.AnnotationsWfWitness();  // WfAnnotationsState(an)
    var s := AhpState(
      root          := RootState([], None, None, None),
      session       := sess,
      chat          := chat,
      terminal      := T.TerminalState("term", None, None, None, [], None, None, None),
      changeset     := cs,
      annotations   := an,
      resourceWatch := RW.ResourceWatchState("/workspace", true),
      canvas        := Cv.CanvasState("markdown", "opaque-provider", None,
                         Some("README.md"), None, Some("ahp-session:/content/1"), Cv.Ready));
    assert WfAhpStateInv(s);
    var a := ASession(sa);                  // WfAhpAction(a) == S.WfSessionActionInv(sa), established by the witness
    assert WfAhpAction(a);
    WfAhpStatePreserved(s, a);
    assert WfAhpStateInv(applyAhp(s, a));
  }

  // ── REDUCTION: the unified applyAhp routing codec IS the ONE ConfluxCodec.Codec<AhpAction>, and its round-trip
  //    law IS conflux's generic RoundTrips restricted to the well-formed (typedAhp ∧ wfAhp) domain. Proven HERE,
  //    INSIDE the module, because the round-trip machinery (AhpActionRoundTrip) is deliberately sealed behind the
  //    .doo export firewall (lines 55-61) — an EXTERNAL *-codec-is-codec.dfy bridge cannot discharge it. This
  //    closes the applyAhp codec-collapse internally: the 7-channel routing codec is not a new codec shape, it
  //    is the SAME ConfluxCodec.Codec as every leaf channel, wrapped once at the dispatch layer. Not exported
  //    (kept internal, consistent with the firewall) — the fact that it reduces is proven, not the machinery.
  function AhpCodecC(): Wire.Codec<AhpAction> {
    Wire.Codec(encodeAhpAction, (j: Wire.Json) => Wire.Some(decodeAhpAction(j)))
  }
  lemma AhpCodecIsConfluxCodec()
    ensures forall a :: typedAhp(a) && wfAhp(a) ==> AhpCodecC().decode(AhpCodecC().encode(a)) == Wire.Some(a)
  {
    forall a | typedAhp(a) && wfAhp(a)
      ensures AhpCodecC().decode(AhpCodecC().encode(a)) == Wire.Some(a)
    { AhpActionRoundTrip(a); }
  }
}
