//! Safe Rust APIs over the verified channel reducers.
//!
//! Each submodule wraps one AHP channel. They share a shape:
//!
//! | Item | Meaning |
//! |---|---|
//! | `Channel::new()` | the core's initial state for that channel |
//! | `.apply(&action)` | one transition through the verified reducer |
//! | `.reduce(&action)` | same, plus the reducer's [`Outcome`][crate::Outcome] |
//! | `.apply_all(&actions)` | a batch, through the core's **proven kernel fold** |
//! | readout methods | typed views of the state (`String`, `Option<i64>`, …) |
//!
//! State values are immutable and cheap to clone; every transition returns a
//! new value and leaves the receiver alone.
//!
//! # Coverage
//!
//! The verified core has eight channels. Five have hand-written safe APIs here:
//!
//! | Channel | Safe API | Exercised by [`corpus`][crate::corpus] |
//! |---|---|---|
//! | Terminal | [`terminal`] | yes |
//! | Canvas | [`canvas`] | yes |
//! | Changeset | [`changeset`] | yes |
//! | Annotations | [`annotations`] | yes |
//! | ResourceWatch | [`resource_watch`] | yes |
//! | Root (AhpSkeleton) | — | yes |
//! | Session | — | yes |
//! | Chat | — | yes |
//!
//! The three without a safe API are **not** missing from the crate — their
//! reducers are extracted, verified, and run by
//! [`corpus::run`][crate::corpus::run] like every other channel. What they lack
//! is a hand-written Rust wrapper; reach them through the [`generated`][crate::generated]
//! module, at the cost of writing Dafny-shaped types by hand. Session and Chat
//! carry substantially larger action sets (14 and 32 variants) than the five
//! wrapped here, and the core models them partially — see the per-channel
//! counts printed by [`corpus::run`][crate::corpus::run].
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

pub mod annotations;
pub mod canvas;
pub mod changeset;
pub mod resource_watch;
pub mod terminal;
