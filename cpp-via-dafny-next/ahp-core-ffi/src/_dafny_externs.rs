// Rust realization of the conflux runtime's {:extern} capability modules.
//
// These are the HOST-EFFECT TRUST BOUNDARY: process execution, filesystem,
// clocks, sockets, HTTP. Dafny declares them `replaceable` / `{:extern}` and
// does NOT verify them -- every language binding of this core supplies its own
// (see the C# and JavaScript adapters under the runtime's extern/ directory,
// and the Go adapter.go files in the sibling `go/` module).
//
// This build exposes the PURE, VERIFIED REDUCERS over a C ABI. Reducers are
// total functions from (state, action) to state; they perform no I/O and never
// reach any function in this file. The stubs therefore abort loudly rather than
// return fabricated data: a silent default here would be indistinguishable from
// a real host effect and would corrupt any caller that did reach it.
//
// If you need the I/O capabilities, implement them here -- the signatures below
// are exactly the ones the translated core calls.
//
// Copyright (c) Microsoft Corporation
// Copyright (c) 2026 Josh Mouch
// SPDX-License-Identifier: MIT

#![allow(non_snake_case)]
#![allow(unused_variables)]

use ::dafny_runtime::{DafnyChar, DafnyInt, Sequence};
use ::std::rc::Rc;

/// Every stub funnels through here so the failure is unambiguous in a core dump
/// or a C++ stack trace.
#[cold]
#[inline(never)]
fn unavailable(what: &str) -> ! {
    panic!(
        "ahp-core-ffi: host capability `{what}` is not implemented in this build. \
         This build exposes the pure verified reducers only; I/O capabilities are \
         the unverified host boundary and must be supplied by the embedder."
    );
}

type Str = Sequence<DafnyChar>;
type Bytes = Sequence<u8>;

pub mod ConfluxIoCapability {
    use super::*;

    pub struct _default {}

    impl _default {
        pub fn ReadFile(path: &Str) -> (bool, Str) {
            unavailable("ConfluxIoCapability.ReadFile")
        }

        pub fn RunProcess(cwd: &Str, cmd: &Str, args: &Sequence<Str>) -> (DafnyInt, Str, Str) {
            unavailable("ConfluxIoCapability.RunProcess")
        }

        pub fn RunProcessBounded(
            cwd: &Str,
            cmd: &Str,
            args: &Sequence<Str>,
            timeLimitMs: &DafnyInt,
            memoryLimitMb: &DafnyInt,
            outputLimitMb: &DafnyInt,
        ) -> (
            DafnyInt, Str, Str, bool, bool, bool, DafnyInt, bool, DafnyInt, DafnyInt,
            bool, bool, bool, bool, DafnyInt, DafnyInt, Str, bool, DafnyInt, DafnyInt,
            bool, DafnyInt, bool, bool, DafnyInt,
        ) {
            unavailable("ConfluxIoCapability.RunProcessBounded")
        }

        pub fn MonotonicTimeMs() -> DafnyInt {
            unavailable("ConfluxIoCapability.MonotonicTimeMs")
        }

        pub fn CurrentProcessRssMb() -> (bool, DafnyInt) {
            unavailable("ConfluxIoCapability.CurrentProcessRssMb")
        }

        pub fn ProcessRssMb(proc: &DafnyInt) -> (bool, DafnyInt) {
            unavailable("ConfluxIoCapability.ProcessRssMb")
        }

        pub fn ProcessExitStatus(proc: &DafnyInt) -> (bool, bool, DafnyInt) {
            unavailable("ConfluxIoCapability.ProcessExitStatus")
        }

        pub fn PollWatch(watcher: &DafnyInt, waitMs: &DafnyInt) -> (bool, Str) {
            unavailable("ConfluxIoCapability.PollWatch")
        }

        pub fn ReadProcessOutputWithin(
            proc: &DafnyInt,
            waitMs: &DafnyInt,
        ) -> (Bytes, bool, bool, DafnyInt) {
            unavailable("ConfluxIoCapability.ReadProcessOutputWithin")
        }

        pub fn TerminateProcessTreeWithin(
            proc: &DafnyInt,
            deadlineMs: &DafnyInt,
        ) -> (bool, bool, bool, bool, DafnyInt, Str, Bytes, bool, DafnyInt, bool, DafnyInt) {
            unavailable("ConfluxIoCapability.TerminateProcessTreeWithin")
        }
    }
}

pub mod ConfluxHttpCapability {
    use super::*;

    pub struct _default {}

    impl _default {
        pub fn Respond(
            request: &DafnyInt,
            status: &DafnyInt,
            contentType: &Str,
            headers: &Sequence<Str>,
            body: &Str,
        ) {
            unavailable("ConfluxHttpCapability.Respond")
        }

        pub fn BeginStream(
            request: &DafnyInt,
            contentType: &Str,
            headers: &Sequence<Str>,
            initialBody: &Str,
        ) {
            unavailable("ConfluxHttpCapability.BeginStream")
        }

        pub fn Publish(server: &DafnyInt, payload: &Str) {
            unavailable("ConfluxHttpCapability.Publish")
        }
    }
}

pub mod ConfluxCommandIdentityCapability {
    use super::*;

    pub struct _default {}

    impl _default {
        pub fn CanonicalJsonBytes(value: &Rc<crate::ConfluxCodec::Json>) -> Bytes {
            unavailable("ConfluxCommandIdentityCapability.CanonicalJsonBytes")
        }

        pub fn DecodeCanonicalJson(
            bytes: &Bytes,
        ) -> Rc<crate::ConfluxCodec::Option<Rc<crate::ConfluxCodec::Json>>> {
            unavailable("ConfluxCommandIdentityCapability.DecodeCanonicalJson")
        }

        pub fn Sha256Digest(bytes: &Bytes) -> Str {
            unavailable("ConfluxCommandIdentityCapability.Sha256Digest")
        }
    }
}

pub mod ConfluxSocketCapability {
    // Declared so the translated module's `pub use crate::_dafny_externs::
    // ConfluxSocketCapability::*;` resolves. No member of this capability is
    // reached by the translated core.
}

pub mod ConfluxWebSocketCapability {
    // See ConfluxSocketCapability above.
}
