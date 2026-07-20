// core_lib.dfy — the aggregate CORE LIBRARY surface.
//
// `gen/build-doo.sh` compiles THIS to `agent-host-protocol-core.doo` so the
// exported library carries EVERY channel reducer (root + session + chat +
// terminal + changeset + annotations + resourceWatch), not just root. Downstream
// (client / host / convergence) can then fold ANY channel's state through that
// channel's `apply1` — which is what makes convergence provable per-channel.
//
// It is pure aggregation: it defines no module of its own, only `include`s the
// seven channel modules (each already verified via client_main.dfy). The include
// list mirrors client_main.dfy's (a proven-consistent include graph), minus the
// client demo module.
include "ahp_skeleton.dfy"
include "fidelity.dfy"
include "resource_watch.dfy"
include "canvas.dfy"
include "changeset.dfy"
include "annotations.dfy"
include "terminal.dfy"
include "session.dfy"
include "chat.dfy"
// per-channel Json<->Action codecs (round-trip proven) — the protocol-complete codec layer
include "codec/session_codec.dfy"
include "codec/chat_codec.dfy"
include "codec/terminal_codec.dfy"
include "codec/changeset_codec.dfy"
include "codec/annotations_codec.dfy"
include "codec/resource_watch_codec.dfy"
include "codec/canvas_codec.dfy"
// UNIFIED MULTI-CHANNEL layer: one AhpAction/AhpState composing all 7 channels,
// a pure per-channel dispatch (applyAhp), and a routing codec — so a generic
// convergence proof + the transport can cover a REAL mixed-channel session.
include "ahp.dfy"
// AHP protocol specialization for the shared ConfluxCommand authority.
include "command.dfy"
include "turn_window_proposal.dfy"
