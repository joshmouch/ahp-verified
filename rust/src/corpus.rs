//! The verified core's own conformance corpus.
//!
//! The Dafny core ships a corpus that replays every channel's fixture set
//! through the reducers and `expect`s a full green board. [`run`] executes that
//! corpus — `ClientMain.Main` — against this extraction.
//!
//! This is the crate's broadest correctness signal: it covers all **eight**
//! channels, including the three without a hand-written safe API, and it checks
//! the extracted reducers against expectations authored in Dafny rather than
//! against expectations re-typed here.
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

use crate::convert::to_dseq;

/// Run the verified core's conformance corpus, printing a per-channel board to
/// stdout.
///
/// # Failure behaviour
///
/// A failed corpus expectation **panics** — the Dafny `expect` becomes a Rust
/// panic in the extracted code. This function returning at all is the pass
/// signal; there is no silent-failure path and no error value to ignore.
///
/// ```no_run
/// ahp_verified::corpus::run(); // panics if any fixture is red
/// ```
pub fn run() {
    let args = to_dseq(Vec::new());
    crate::generated::ClientMain::_default::Main(&args);
}
