// C ABI bridge over the Dafny-extracted, formally-verified AHP core.
//
// Every function below is a thin marshalling shell. The decisions — how a
// reducer folds an action into a state, what a canonical JSON encoding is,
// what a command-identity digest is — are made entirely inside the
// Dafny-extracted Go packages under github.com/joshmouch/ahp-verified/go,
// which are machine-generated from the proven specification. This file
// converts C strings and integers to and from Dafny values and calls them.
// It contains no protocol logic of its own.
//
// Built with `go build -buildmode=c-archive`, this produces a native static
// archive plus a generated C header that C and C++ programs link directly.
//
// Copyright (c) Microsoft Corporation
// Copyright (c) 2026 Josh Mouch
// SPDX-License-Identifier: MIT

package main

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"runtime/cgo"
	"unsafe"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Identity "github.com/joshmouch/ahp-verified/go/ConfluxCommandIdentityCapability"
	m_ConfluxJsonText "github.com/joshmouch/ahp-verified/go/ConfluxJsonText"
	m_Session "github.com/joshmouch/ahp-verified/go/Session"
	_dafny "github.com/joshmouch/ahp-verified/go/dafny"
)

// c-archive requires a main package with an empty main.
func main() {}

// ─────────────────────────── marshalling helpers ───────────────────────────

// dseq converts a NUL-terminated C string to a Dafny string sequence.
func dseq(s *C.char) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes(C.GoString(s))
}

// cstr converts a Dafny string sequence to a malloc'd C string. The caller
// releases it with ahp_string_free.
func cstr(s _dafny.Sequence) *C.char {
	return C.CString(s.VerbatimString(false))
}

// optSeq maps a nullable C string to Option<string>: NULL becomes None.
func optSeq(s *C.char) m_AhpSkeleton.Option {
	if s == nil {
		return m_AhpSkeleton.Companion_Option_.Create_None_()
	}
	return m_AhpSkeleton.Companion_Option_.Create_Some_(dseq(s))
}

// optCStr renders an Option<string> as a C string, or NULL when it is None.
func optCStr(o m_AhpSkeleton.Option) *C.char {
	if !o.Is_Some() {
		return nil
	}
	return cstr(o.Dtor_value().(_dafny.Sequence))
}

// availabilityOf maps the wire-level availability code to the Dafny datatype.
// AHP_AVAILABILITY_READY (0) and AHP_AVAILABILITY_STALE (1) are the only
// variants the specification declares; anything else is refused as None.
func optAvailability(code C.int) m_AhpSkeleton.Option {
	switch code {
	case 0:
		return m_AhpSkeleton.Companion_Option_.Create_Some_(
			m_Canvas.Companion_CanvasAvailability_.Create_Ready_())
	case 1:
		return m_AhpSkeleton.Companion_Option_.Create_Some_(
			m_Canvas.Companion_CanvasAvailability_.Create_Stale_())
	default:
		return m_AhpSkeleton.Companion_Option_.Create_None_()
	}
}

// outcomeCode maps ReduceOutcome to the AHP_OUTCOME_* constants.
func outcomeCode(o m_AhpSkeleton.ReduceOutcome) C.int {
	switch {
	case o.Is_Applied():
		return 0
	case o.Is_NoOp():
		return 1
	default:
		return 2
	}
}

// canvasOf resolves an opaque handle back to the verified state it names.
func canvasOf(h C.uintptr_t) m_Canvas.CanvasState {
	return cgo.Handle(h).Value().(m_Canvas.CanvasState)
}

// ───────────────────────────── general exports ─────────────────────────────

//export ahp_abi_version
//
// Returns the version of THIS binding layer's C ABI. It is not a protocol
// version: it changes when the C surface below changes, and says nothing
// about the verified core.
func ahp_abi_version() *C.char {
	return C.CString("1.0.0")
}

//export ahp_string_free
//
// Releases any char* returned by this library.
func ahp_string_free(s *C.char) {
	if s != nil {
		C.free(unsafe.Pointer(s))
	}
}

//export ahp_run_corpus
//
// Runs the conformance corpus embedded in the verified Session module and
// reports how many of the modeled cases the extracted reducers reproduce.
// Writes the pass count and the modeled-case count through the out
// parameters. Returns 1 when every modeled case passed, 0 otherwise.
func ahp_run_corpus(outPassed *C.int64_t, outModeled *C.int64_t) C.int {
	passed, modeled := m_Session.Companion_Default___.RunCorpus()
	if outPassed != nil {
		*outPassed = C.int64_t(passed.Int64())
	}
	if outModeled != nil {
		*outModeled = C.int64_t(modeled.Int64())
	}
	if passed.Cmp(modeled) == 0 {
		return 1
	}
	return 0
}

// ───────────────────────── verified codec + identity ────────────────────────

//export ahp_json_canonicalize
//
// Parses JSON text with the verified parser and re-serializes it with the
// verified stringifier, yielding the core's canonical spelling of the value.
// Returns NULL when the text is not accepted by the verified parser. Release
// a non-NULL result with ahp_string_free.
func ahp_json_canonicalize(text *C.char) *C.char {
	parsed := m_ConfluxJsonText.Companion_Default___.ParseJsonChecked(dseq(text))
	if !parsed.Is_Parsed() {
		return nil
	}
	return cstr(m_ConfluxJsonText.Companion_Default___.Stringify(parsed.Dtor_value()))
}

//export ahp_sha256_digest
//
// Computes the verified command-identity digest over len bytes at data,
// returned in the core's "sha256:<hex>" spelling. Release the result with
// ahp_string_free.
func ahp_sha256_digest(data *C.char, length C.int) *C.char {
	bytes := C.GoBytes(unsafe.Pointer(data), length)
	return cstr(m_Identity.Companion_Default___.Sha256Digest(_dafny.SeqOfBytes(bytes)))
}

// ────────────────────────── Canvas channel reducer ──────────────────────────

//export ahp_canvas_new
//
// Builds a fresh CanvasState with the given identifiers, no title, activity,
// content URI or input, and Ready availability. Returns an opaque handle the
// caller releases with ahp_canvas_release.
func ahp_canvas_new(canvasID *C.char, providerID *C.char) C.uintptr_t {
	none := m_AhpSkeleton.Companion_Option_.Create_None_()
	state := m_Canvas.Companion_CanvasState_.Create_CanvasState_(
		dseq(canvasID), dseq(providerID),
		none, none, none, none,
		m_Canvas.Companion_CanvasAvailability_.Create_Ready_(),
	)
	return C.uintptr_t(cgo.NewHandle(state))
}

//export ahp_canvas_apply_updated
//
// Applies the Canvas `Updated` action through the verified reducer
// Canvas.ApplyToCanvas. Each of title, activity and contentUri may be NULL to
// leave that field untouched; availability takes AHP_AVAILABILITY_READY,
// AHP_AVAILABILITY_STALE, or AHP_AVAILABILITY_UNCHANGED. The reducer is pure —
// the input handle stays valid and unmodified, and a new handle is returned.
// When outOutcome is non-NULL it receives the reducer's AHP_OUTCOME_* verdict.
func ahp_canvas_apply_updated(
	handle C.uintptr_t,
	title *C.char,
	activity *C.char,
	contentURI *C.char,
	availability C.int,
	now C.int64_t,
	outOutcome *C.int,
) C.uintptr_t {
	action := m_Canvas.Companion_CanvasAction_.Create_Updated_(
		optSeq(title), optSeq(activity), optSeq(contentURI), optAvailability(availability),
	)
	return applyCanvas(handle, action, now, outOutcome)
}

//export ahp_canvas_apply_close_requested
//
// Applies the Canvas `CloseRequested` action through the verified reducer.
// The specification defines this as state-preserving, so the returned state
// equals the input and the outcome is AHP_OUTCOME_NOOP — that verdict comes
// from the verified reducer, not from this bridge.
func ahp_canvas_apply_close_requested(
	handle C.uintptr_t,
	now C.int64_t,
	outOutcome *C.int,
) C.uintptr_t {
	action := m_Canvas.Companion_CanvasAction_.Create_CloseRequested_()
	return applyCanvas(handle, action, now, outOutcome)
}

func applyCanvas(
	handle C.uintptr_t,
	action m_Canvas.CanvasAction,
	now C.int64_t,
	outOutcome *C.int,
) C.uintptr_t {
	result := m_Canvas.Companion_Default___.ApplyToCanvas(
		canvasOf(handle), action, _dafny.IntOfInt64(int64(now)))

	next := (*result.IndexInt(0)).(m_Canvas.CanvasState)
	if outOutcome != nil {
		*outOutcome = outcomeCode((*result.IndexInt(1)).(m_AhpSkeleton.ReduceOutcome))
	}
	return C.uintptr_t(cgo.NewHandle(next))
}

//export ahp_canvas_id
//
// Returns the canvasId field. Release with ahp_string_free.
func ahp_canvas_id(handle C.uintptr_t) *C.char {
	return cstr(canvasOf(handle).Dtor_canvasId())
}

//export ahp_canvas_provider_id
//
// Returns the providerId field. Release with ahp_string_free.
func ahp_canvas_provider_id(handle C.uintptr_t) *C.char {
	return cstr(canvasOf(handle).Dtor_providerId())
}

//export ahp_canvas_title
//
// Returns the title, or NULL when the state carries none. Release a non-NULL
// result with ahp_string_free.
func ahp_canvas_title(handle C.uintptr_t) *C.char {
	return optCStr(canvasOf(handle).Dtor_title())
}

//export ahp_canvas_activity
//
// Returns the activity, or NULL when the state carries none. Release a
// non-NULL result with ahp_string_free.
func ahp_canvas_activity(handle C.uintptr_t) *C.char {
	return optCStr(canvasOf(handle).Dtor_activity())
}

//export ahp_canvas_content_uri
//
// Returns the contentUri, or NULL when the state carries none. Release a
// non-NULL result with ahp_string_free.
func ahp_canvas_content_uri(handle C.uintptr_t) *C.char {
	return optCStr(canvasOf(handle).Dtor_contentUri())
}

//export ahp_canvas_availability
//
// Returns AHP_AVAILABILITY_READY or AHP_AVAILABILITY_STALE.
func ahp_canvas_availability(handle C.uintptr_t) C.int {
	if canvasOf(handle).Dtor_availability().Is_Ready() {
		return 0
	}
	return 1
}

//export ahp_canvas_equals
//
// Structural equality over two Canvas states, decided by the extracted
// datatype's own equality.
func ahp_canvas_equals(a C.uintptr_t, b C.uintptr_t) C.int {
	if canvasOf(a).Equals(canvasOf(b)) {
		return 1
	}
	return 0
}

//export ahp_canvas_release
//
// Releases a Canvas handle. Every handle returned by ahp_canvas_new or an
// ahp_canvas_apply_* call must be released exactly once.
func ahp_canvas_release(handle C.uintptr_t) {
	cgo.Handle(handle).Delete()
}
