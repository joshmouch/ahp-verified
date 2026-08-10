//! Formally verified [Agent Host Protocol][ahp] channel reducers for Rust.
//!
//! Every state transition reachable from this crate's public API is
//! **machine-extracted from a proven Dafny core** — not reimplemented. The
//! Dafny source carries the proofs (626 verification units, 0 errors, 8
//! channels, zero core-owned trusted assumptions); `dafny translate rs`
//! produces the Rust in [`generated`]; the modules under [`channels`] are a
//! thin, hand-written, **safe** wrapper that marshals Rust types in and out.
//!
//! [ahp]: https://github.com/microsoft/agent-host-protocol
//!
//! # What is and is not verified
//!
//! | Layer | Verified? |
//! |---|---|
//! | Reducer logic (`apply`, `apply_all`, every state transition) | **Yes** — proven in Dafny, extracted mechanically |
//! | The Dafny-to-Rust code generator and its runtime | No — trusted toolchain |
//! | This crate's marshalling layer (`String` ⇄ Dafny sequence, etc.) | No — hand-written, but covered by the corpus tests |
//! | Host I/O capabilities (process, socket, filesystem, clock) | Not present — see [`externs`] |
//!
//! This crate exposes the **pure reducers only**. They are total functions from
//! `(state, action)` to `state`, perform no I/O, and never reach a host
//! capability. The capability stubs in [`externs`] panic loudly rather than
//! fabricate data, so a build that somehow reached one fails visibly.
//!
//! # Example
//!
//! ```
//! use ahp_verified::channels::terminal::{Terminal, TerminalAction as A};
//!
//! let t = Terminal::new()
//!     .apply(&A::TitleChanged("build".into()))
//!     .apply(&A::CwdChanged("/src".into()))
//!     .apply(&A::Resized { cols: 120, rows: 40 })
//!     .apply(&A::Data("cargo test\n".into()))
//!     .apply(&A::Exited(0));
//!
//! assert_eq!(t.title(), "build");
//! assert_eq!(t.cwd().as_deref(), Some("/src"));
//! assert_eq!(t.size(), Some((120, 40)));
//! assert_eq!(t.exit_code(), Some(0));
//! ```
//!
//! Reducers are **pure**: applying an action returns a new value and leaves the
//! receiver untouched.
//!
//! ```
//! use ahp_verified::channels::terminal::{Terminal, TerminalAction as A};
//!
//! let before = Terminal::new().apply(&A::TitleChanged("one".into()));
//! let after = before.apply(&A::TitleChanged("two".into()));
//!
//! assert_eq!(before.title(), "one"); // unchanged
//! assert_eq!(after.title(), "two");
//! ```
//!
//! # Batch application goes through the proven kernel fold
//!
//! [`Terminal::apply_all`][channels::terminal::Terminal::apply_all] and its
//! siblings call the core's own `fold`, which is defined over
//! `ConfluxContract.Fold` and proven there — rather than looping `apply` on the
//! Rust side. The two agree by construction:
//!
//! ```
//! use ahp_verified::channels::terminal::{Terminal, TerminalAction as A};
//!
//! let acts = vec![A::Data("a".into()), A::Data("b".into()), A::Data("c".into())];
//! let folded = Terminal::new().apply_all(&acts);
//! let looped = acts.iter().fold(Terminal::new(), |s, a| s.apply(a));
//! assert_eq!(folded, looped);
//! ```
//!
//! # Threading
//!
//! The extracted core is reference-counted (`Rc`), so the state types are
//! **`!Send` and `!Sync`**. Keep a state value on one thread, or send the
//! actions rather than the state. This is a property of the Dafny Rust backend
//! without `--rust-sync`, not a choice made here.
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

#![warn(missing_docs)]
#![deny(unsafe_code)]
#![warn(rustdoc::broken_intra_doc_links)]

// ---------------------------------------------------------------------------
// Machine-generated layer.
//
// Both modules below are produced by ../regenerate.sh and must not be edited by
// hand. The Dafny Rust backend assumes it owns the crate root and emits
// absolute `crate::<Module>::...` paths; regenerate.sh re-homes those under
// these two modules, so neither is glob-re-exported here. That is deliberate:
// it keeps ~60 Dafny-shaped modules and the runtime's `DafnyInt` / `Sequence` /
// `_System` out of this crate's public API and out of its semver surface.
//
// `generated` stays `pub` (but hidden) as the documented escape hatch for the
// three channels without a safe wrapper yet -- see `channels`.
// ---------------------------------------------------------------------------

#[macro_use]
#[path = "generated/dafny_runtime/mod.rs"]
#[allow(unsafe_code, missing_docs, warnings, clippy::all, clippy::pedantic)]
mod dafny_runtime;

#[doc(hidden)]
#[path = "generated/ahpcore.rs"]
#[allow(unsafe_code, missing_docs, warnings, clippy::all, clippy::pedantic)]
pub mod generated;

// ---------------------------------------------------------------------------
// Hand-written layer.
// ---------------------------------------------------------------------------

pub mod channels;
pub mod corpus;
pub mod externs;
pub mod json;

mod convert;

pub use json::Json;

/// The disposition the verified reducer assigned to an action.
///
/// Returned alongside the new state by the `reduce` method on every channel.
/// The plain `apply` methods discard it.
#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub enum Outcome {
    /// The action was recognized and the reducer processed it.
    Applied,
    /// The action was recognized but is defined to leave the state unchanged.
    NoOp,
    /// The action is outside this channel's scope.
    OutOfScope,
}

impl Outcome {
    pub(crate) fn from_core(o: &generated::AhpSkeleton::ReduceOutcome) -> Self {
        match o {
            generated::AhpSkeleton::ReduceOutcome::Applied {} => Outcome::Applied,
            generated::AhpSkeleton::ReduceOutcome::NoOp {} => Outcome::NoOp,
            generated::AhpSkeleton::ReduceOutcome::OutOfScope {} => Outcome::OutOfScope,
        }
    }
}
