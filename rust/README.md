# ahp-verified

Formally verified [Agent Host Protocol](https://github.com/microsoft/agent-host-protocol)
channel reducers for Rust.

Every state transition this crate exposes is **machine-extracted from a proven
Dafny core** — not reimplemented. The Dafny source carries the proofs; `dafny
translate rs` produces the Rust; a thin, safe, hand-written wrapper marshals
ordinary Rust types in and out.

```toml
[dependencies]
ahp-verified = "0.1"
```

```rust
use ahp_verified::channels::terminal::{Terminal, TerminalAction as A};

let t = Terminal::new()
    .apply(&A::TitleChanged("build".into()))
    .apply(&A::CwdChanged("/src".into()))
    .apply(&A::Resized { cols: 120, rows: 40 })
    .apply(&A::Data("cargo test\n".into()))
    .apply(&A::Exited(0));

assert_eq!(t.title(), "build");
assert_eq!(t.cwd().as_deref(), Some("/src"));
assert_eq!(t.size(), Some((120, 40)));
assert_eq!(t.exit_code(), Some(0));
```

## Relationship to the upstream Rust client

This crate is **not** the AHP client. Upstream's Rust client is
[`ahp`](https://crates.io/crates/ahp), [`ahp-types`](https://crates.io/crates/ahp-types)
and [`ahp-ws`](https://crates.io/crates/ahp-ws), published by
[Connor Peet](https://github.com/connor4312) from
[microsoft/agent-host-protocol](https://github.com/microsoft/agent-host-protocol).
Use those to speak the protocol.

`ahp-verified` is a different thing: an independent, formally verified model of
the protocol's channel reducers, useful as a differential-testing oracle, as a
reference for reducer semantics, or as the state layer inside a host that wants
its channel logic to be proven rather than hand-written. The name is
deliberately outside upstream's `ahp-*` namespace to avoid any suggestion that
it is an official component.

## What is and is not verified

| Layer | Verified? |
|---|---|
| Reducer logic — every state transition reachable from this API | **Yes.** Proven in Dafny, extracted mechanically |
| The Dafny-to-Rust code generator and its runtime | No — trusted toolchain |
| This crate's marshalling layer (`String` ⇄ Dafny sequence, etc.) | No — hand-written, ~130 lines, covered by tests |
| Host I/O (process, socket, filesystem, clock) | Not present — see below |

The verified core reports **626 verification units, 0 errors, across 8
channels, with zero core-owned trusted assumptions** (16 findings, all inherited
from the vendored runtime, each name-pinned).

This crate exposes the **pure reducers only**. They are total functions from
`(state, action)` to `state`, perform no I/O, and never reach a host capability.
The capability stubs in `ahp_verified::externs` panic loudly rather than
fabricate data, so a build that somehow reached one fails visibly instead of
silently returning a plausible lie.

## API shape

Every channel follows the same shape:

| Item | Meaning |
|---|---|
| `Channel::new()` | the core's initial state for that channel |
| `.apply(&action)` | one transition through the verified reducer |
| `.reduce(&action)` | same, plus the reducer's `Outcome` (`Applied` / `NoOp` / `OutOfScope`) |
| `.apply_all(&actions)` | a batch, through the core's **proven kernel fold** |
| readouts | typed views (`String`, `Option<i64>`, `Vec<Part>`, …) |

State values are immutable and cheap to clone. Reducers are pure, so a
transition returns a new value and leaves the receiver alone:

```rust
# use ahp_verified::channels::terminal::{Terminal, TerminalAction as A};
let before = Terminal::new().apply(&A::TitleChanged("one".into()));
let after = before.apply(&A::TitleChanged("two".into()));

assert_eq!(before.title(), "one"); // unchanged
assert_eq!(after.title(), "two");
```

`apply_all` calls the core's own `fold` (defined over `ConfluxContract.Fold` and
proven there) rather than looping on the Rust side. The two agree by
construction, and a test pins that they keep agreeing.

`.reduce()` matters more than it looks. Some actions are *defined* to leave the
state unchanged, and the state alone cannot distinguish "the reducer processed
this" from "the reducer recognized this as a no-op". ResourceWatch is entirely
this case — it is a passthrough channel where the outcome is the only
observable.

## Channel coverage

The core has eight channels. Five have hand-written safe APIs:

| Channel | Safe API | In the conformance corpus |
|---|---|---|
| Terminal | `channels::terminal` | yes |
| Canvas | `channels::canvas` | yes |
| Changeset | `channels::changeset` | yes |
| Annotations | `channels::annotations` | yes |
| ResourceWatch | `channels::resource_watch` | yes |
| Root (AhpSkeleton) | — | yes |
| Session | — | yes |
| Chat | — | yes |

The three without a safe API are **not missing from the crate** — their
reducers are extracted, verified, and exercised by `corpus::run()` like every
other channel. What they lack is a hand-written Rust wrapper. They are reachable
through the hidden `generated` module at the cost of writing Dafny-shaped types
by hand. Session and Chat carry substantially larger action sets (14 and 32
variants) than the five wrapped here, and the core models them partially; the
per-channel counts are printed by `corpus::run()`.

## The conformance corpus

The Dafny core ships its own corpus, and this crate runs it:

```rust
ahp_verified::corpus::run();
```

```text
ROOT CORPUS:          7/7 green against extracted code
RESOURCEWATCH CORPUS: 2/2 green against extracted code
CANVAS CORPUS:        5/5 green against extracted code
CHANGESET CORPUS:     15/15 green against extracted code
ANNOTATIONS CORPUS:   10/10 green against extracted code
TERMINAL CORPUS:      19/19 green against extracted code
SESSION CORPUS:       36/36 modeled green (of 61 total; all ~25 action TYPES now modeled)
CHAT CORPUS:          54/54 modeled green (of 97 total; full tool-call state machine + turn lifecycle modeled)
TOTAL: 148/148 corpus fixtures green (5 full AHP channels + session/chat partial)
```

A failed fixture panics; the function returning at all is the pass signal. Note
the honest denominators: Session and Chat report *modeled* fixtures (36 of 61,
54 of 97), not full channel coverage.

## Threading

The extracted core is reference-counted (`Rc`), so state types are **`!Send` and
`!Sync`**. Keep a state value on one thread, or send actions rather than state.
This follows from the Dafny Rust backend without `--rust-sync`; it is not a
choice made here.

## Dependencies

Three, all transitively required by the vendored Dafny runtime: `once_cell`,
`num`, `itertools`. No async runtime, no I/O, no networking.

The Dafny runtime is vendored into `src/generated/dafny_runtime` rather than
taken from the crates.io `dafny-runtime`, because that crate is a third-party
vendoring (published from the AWS Database Encryption SDK) at a version that
does not correspond to the Dafny 4.11.0 code generator used here. Vendoring
keeps the runtime and the generated code exactly in step.

## Regenerating from the Dafny core

```bash
./regenerate.sh          # from the parent directory of this crate
```

The script extracts the runtime's Dafny source from its `.doo`, applies the one
source rewrite the Rust backend requires, **re-verifies** the patched runtime
(this is the step that proves the rewrite preserved every postcondition),
translates to Rust, folds the output into this crate, then builds and tests.

Everything under `src/generated/` is machine-produced and must not be hand
edited. The hand-written sources are `src/lib.rs`, `src/externs.rs`,
`src/convert.rs`, `src/json.rs`, and `src/channels/*.rs`.

The Dafny Rust backend assumes it owns the crate root — it emits absolute
`crate::<Module>::…` paths, and the runtime refers to itself the same way.
`regenerate.sh` re-homes those paths under `crate::generated` and
`crate::dafny_runtime` so that neither is glob-re-exported at the root. That
keeps ~60 Dafny-shaped modules, and the runtime's `DafnyInt` / `Sequence` /
`_System`, out of this crate's public API and off its semver surface. The
rewrites are derived from the generated source on every run, not hardcoded.

## Verification and lints

```bash
cargo test     # 50 tests
cargo clippy --all-targets   # clean
cargo doc --no-deps
cargo run --example terminal_session
```

`unsafe_code` is denied for the hand-written layer and allowed only on the two
machine-generated modules. `missing_docs` is warned on the same split.

## License

MIT. Copyright (c) Microsoft Corporation (protocol and original client),
copyright (c) 2026 Josh Mouch (verification).
