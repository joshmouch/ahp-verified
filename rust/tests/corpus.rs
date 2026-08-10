//! Runs the verified core's own conformance corpus through this extraction.
//!
//! This is the broadest signal in the suite: it covers all eight channels using
//! expectations authored in Dafny, including the three channels this crate does
//! not yet wrap in a safe API. A failed fixture panics inside the extracted
//! code, so the test fails loudly rather than returning a status nobody reads.
//!
//! Copyright (c) Microsoft Corporation
//! Copyright (c) 2026 Josh Mouch
//! SPDX-License-Identifier: MIT

#[test]
fn core_conformance_corpus_is_green() {
    ahp_verified::corpus::run();
}
