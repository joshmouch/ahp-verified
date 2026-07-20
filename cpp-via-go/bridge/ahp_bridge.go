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
// HANDLES
//   A Canvas handle is an opaque uintptr token naming an entry in the registry
//   below, NOT a pointer. Tokens start at 1 (0 is the invalid handle) and are
//   NEVER reused, so a released token can never come to name a different live
//   state. Every entry point validates its handle: an unknown, released, or
//   zero token is REFUSED (NULL / 0 / -1 per the function), never dereferenced.
//   This replaces runtime/cgo.Handle, whose Value() and Delete() PANIC on an
//   invalid handle — and a panic escaping an //export'd function terminates the
//   whole process, which a C caller has no way to defend against.
//
// PANICS NEVER CROSS THE BOUNDARY
//   Every //export'd function below recovers. A panic inside the extracted core
//   or the Go runtime is converted into that function's documented failure
//   value instead of killing the caller's process.
//
// THREAD SAFETY
//   Every entry point is safe to call concurrently, because they are SERIALIZED
//   by coreMu. This is not belt-and-braces: the extracted core is not
//   thread-safe. dafny.LazySequence.ToArray() memoizes by calling Box().Put(),
//   and the Go runtime's AtomicBox implementation (dafny.GoAtomicBox) is a
//   plain struct field with unsynchronized Get/Put. Reading one sequence from
//   two goroutines is therefore a data race on that field, and Dafny values
//   structurally share subterms, so the race is reachable from ordinary use of
//   this C ABI. Serializing is the honest fix available to a binding that
//   consumes already-extracted code; the cost is stated in the header.
//
// UTF-8 CONTRACT (shared by all three bindings in this repository)
//   Input strings MUST be well-formed UTF-8. Invalid input is REFUSED, never
//   repaired. Silently mapping malformed bytes to U+FFFD merges DISTINCT
//   inputs into ONE verified state: "c\xff" and "c\xfe" both become "c�",
//   after which the core's proven equality correctly reports two different
//   canvasIds as equal. The proof stays intact; the binding fed it the wrong
//   values. So the binding refuses instead.
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
	"sync"
	"unicode/utf8"
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

// maxDigestLen bounds ahp_sha256_digest's length argument. The true size of a
// C buffer cannot be checked against the buffer, so this refuses lengths that
// cannot name a real one before any Go allocation is sized from the number.
const maxDigestLen = 1 << 28 // 256 MiB

// ───────────────────────────── core serialization ──────────────────────────

// coreMu guards BOTH the handle registry and every call into the extracted
// core. See THREAD SAFETY above for why the core calls need it too.
var coreMu sync.Mutex

// ────────────────────────────── handle registry ────────────────────────────

var (
	canvases   = make(map[uintptr]m_Canvas.CanvasState)
	nextHandle uintptr = 1
)

// retain registers a state and returns its handle. Caller holds coreMu.
func retain(state m_Canvas.CanvasState) C.uintptr_t {
	h := nextHandle
	nextHandle++
	canvases[h] = state
	return C.uintptr_t(h)
}

// canvasOf resolves a handle to the verified state it names. The bool is false
// when the handle is zero, was never issued, or has been released. Caller holds
// coreMu.
func canvasOf(h C.uintptr_t) (m_Canvas.CanvasState, bool) {
	state, ok := canvases[uintptr(h)]
	return state, ok
}

// ─────────────────────────────── marshalling ───────────────────────────────

// dseq converts a NUL-terminated C string to a Dafny string sequence. The bool
// is false when the pointer is NULL or the bytes are not well-formed UTF-8;
// invalid input is refused, never repaired (see the UTF-8 CONTRACT above).
func dseq(s *C.char) (_dafny.Sequence, bool) {
	if s == nil {
		return nil, false
	}
	g := C.GoString(s)
	if !utf8.ValidString(g) {
		return nil, false
	}
	return _dafny.UnicodeSeqOfUtf8Bytes(g), true
}

// cstr converts a Dafny string sequence to a malloc'd C string. The caller
// releases it with ahp_string_free.
func cstr(s _dafny.Sequence) *C.char {
	return C.CString(s.VerbatimString(false))
}

// optSeq maps a nullable C string to Option<string>: NULL becomes None. The
// bool is false only for invalid UTF-8 — NULL is a legitimate "leave alone".
func optSeq(s *C.char) (m_AhpSkeleton.Option, bool) {
	if s == nil {
		return m_AhpSkeleton.Companion_Option_.Create_None_(), true
	}
	seq, ok := dseq(s)
	if !ok {
		return m_AhpSkeleton.Companion_Option_.Create_None_(), false
	}
	return m_AhpSkeleton.Companion_Option_.Create_Some_(seq), true
}

// optCStr renders an Option<string> as a C string, or NULL when it is None.
func optCStr(o m_AhpSkeleton.Option) *C.char {
	if !o.Is_Some() {
		return nil
	}
	return cstr(o.Dtor_value().(_dafny.Sequence))
}

// optAvailability maps the wire-level availability code to the Dafny datatype.
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

// ─────────────────────── panic containment at the boundary ─────────────────
//
// A panic escaping an //export'd function unwinds into C and terminates the
// process. Each export defers one of these so a panic anywhere below — in the
// extracted core, in the Dafny runtime, in cgo — becomes that function's
// documented failure value instead.

func trapStr(out **C.char) {
	if r := recover(); r != nil {
		if *out != nil {
			C.free(unsafe.Pointer(*out))
		}
		*out = nil
	}
}

func trapHandle(out *C.uintptr_t) {
	if r := recover(); r != nil {
		*out = 0
	}
}

func trapInt(out *C.int, v C.int) {
	if r := recover(); r != nil {
		*out = v
	}
}

// ───────────────────────────── general exports ─────────────────────────────

//export ahp_abi_version
//
// Returns the version of THIS binding layer's C ABI. It is not a protocol
// version: it changes when the C surface below changes, and says nothing
// about the verified core.
func ahp_abi_version() *C.char {
	return C.CString("1.1.0")
}

//export ahp_string_free
//
// Releases any char* returned by this library. NULL is accepted.
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
func ahp_run_corpus(outPassed *C.int64_t, outModeled *C.int64_t) (result C.int) {
	defer trapInt(&result, 0)
	coreMu.Lock()
	defer coreMu.Unlock()

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
// Returns NULL when text is NULL, is not valid UTF-8, or is not accepted by
// the verified parser. Release a non-NULL result with ahp_string_free.
func ahp_json_canonicalize(text *C.char) (result *C.char) {
	defer trapStr(&result)
	coreMu.Lock()
	defer coreMu.Unlock()

	seq, ok := dseq(text)
	if !ok {
		return nil
	}
	parsed := m_ConfluxJsonText.Companion_Default___.ParseJsonChecked(seq)
	if !parsed.Is_Parsed() {
		return nil
	}
	return cstr(m_ConfluxJsonText.Companion_Default___.Stringify(parsed.Dtor_value()))
}

//export ahp_sha256_digest
//
// Computes the verified command-identity digest over length bytes at data,
// returned in the core's "sha256:<hex>" spelling. Returns NULL when length is
// negative or implausibly large, or when data is NULL with a non-zero length.
// Release a non-NULL result with ahp_string_free.
func ahp_sha256_digest(data *C.char, length C.int) (result *C.char) {
	defer trapStr(&result)

	// Validated BEFORE C.GoBytes, which throws an unrecoverable runtime error
	// ("gobytes: length out of range") on a negative length and reads address
	// zero when data is NULL.
	if length < 0 || int64(length) > maxDigestLen {
		return nil
	}
	if data == nil && length != 0 {
		return nil
	}
	var bytes []byte
	if length > 0 {
		bytes = C.GoBytes(unsafe.Pointer(data), length)
	}

	coreMu.Lock()
	defer coreMu.Unlock()
	return cstr(m_Identity.Companion_Default___.Sha256Digest(_dafny.SeqOfBytes(bytes)))
}

// ────────────────────────── Canvas channel reducer ──────────────────────────

//export ahp_canvas_new
//
// Builds a fresh CanvasState with the given identifiers, no title, activity,
// content URI or input, and Ready availability. Returns an opaque handle the
// caller releases with ahp_canvas_release, or 0 when either identifier is NULL
// or is not valid UTF-8.
func ahp_canvas_new(canvasID *C.char, providerID *C.char) (result C.uintptr_t) {
	defer trapHandle(&result)
	coreMu.Lock()
	defer coreMu.Unlock()

	id, ok := dseq(canvasID)
	if !ok {
		return 0
	}
	provider, ok := dseq(providerID)
	if !ok {
		return 0
	}
	none := m_AhpSkeleton.Companion_Option_.Create_None_()
	state := m_Canvas.Companion_CanvasState_.Create_CanvasState_(
		id, provider,
		none, none, none, none,
		m_Canvas.Companion_CanvasAvailability_.Create_Ready_(),
	)
	return retain(state)
}

//export ahp_canvas_apply_updated
//
// Applies the Canvas `Updated` action through the verified reducer
// Canvas.ApplyToCanvas. Each of title, activity and contentUri may be NULL to
// leave that field untouched; availability takes AHP_AVAILABILITY_READY,
// AHP_AVAILABILITY_STALE, or AHP_AVAILABILITY_UNCHANGED. The reducer is pure —
// the input handle stays valid and unmodified, and a new handle is returned.
// When outOutcome is non-NULL it receives the reducer's AHP_OUTCOME_* verdict.
// Returns 0 when the handle is invalid or a non-NULL string is not UTF-8.
func ahp_canvas_apply_updated(
	handle C.uintptr_t,
	title *C.char,
	activity *C.char,
	contentURI *C.char,
	availability C.int,
	now C.int64_t,
	outOutcome *C.int,
) (result C.uintptr_t) {
	defer trapHandle(&result)
	coreMu.Lock()
	defer coreMu.Unlock()

	titleOpt, ok := optSeq(title)
	if !ok {
		return 0
	}
	activityOpt, ok := optSeq(activity)
	if !ok {
		return 0
	}
	uriOpt, ok := optSeq(contentURI)
	if !ok {
		return 0
	}
	action := m_Canvas.Companion_CanvasAction_.Create_Updated_(
		titleOpt, activityOpt, uriOpt, optAvailability(availability),
	)
	return applyCanvas(handle, action, now, outOutcome)
}

//export ahp_canvas_apply_close_requested
//
// Applies the Canvas `CloseRequested` action through the verified reducer.
// The specification defines this as state-preserving, so the returned state
// equals the input and the outcome is AHP_OUTCOME_NOOP — that verdict comes
// from the verified reducer, not from this bridge. Returns 0 on an invalid
// handle.
func ahp_canvas_apply_close_requested(
	handle C.uintptr_t,
	now C.int64_t,
	outOutcome *C.int,
) (result C.uintptr_t) {
	defer trapHandle(&result)
	coreMu.Lock()
	defer coreMu.Unlock()

	action := m_Canvas.Companion_CanvasAction_.Create_CloseRequested_()
	return applyCanvas(handle, action, now, outOutcome)
}

// applyCanvas runs the verified reducer. Caller holds coreMu.
func applyCanvas(
	handle C.uintptr_t,
	action m_Canvas.CanvasAction,
	now C.int64_t,
	outOutcome *C.int,
) C.uintptr_t {
	state, ok := canvasOf(handle)
	if !ok {
		return 0
	}
	result := m_Canvas.Companion_Default___.ApplyToCanvas(
		state, action, _dafny.IntOfInt64(int64(now)))

	next := (*result.IndexInt(0)).(m_Canvas.CanvasState)
	if outOutcome != nil {
		*outOutcome = outcomeCode((*result.IndexInt(1)).(m_AhpSkeleton.ReduceOutcome))
	}
	return retain(next)
}

//export ahp_canvas_id
//
// Returns the canvasId field, or NULL on an invalid handle. Release a non-NULL
// result with ahp_string_free.
func ahp_canvas_id(handle C.uintptr_t) (result *C.char) {
	defer trapStr(&result)
	coreMu.Lock()
	defer coreMu.Unlock()

	state, ok := canvasOf(handle)
	if !ok {
		return nil
	}
	return cstr(state.Dtor_canvasId())
}

//export ahp_canvas_provider_id
//
// Returns the providerId field, or NULL on an invalid handle. Release a
// non-NULL result with ahp_string_free.
func ahp_canvas_provider_id(handle C.uintptr_t) (result *C.char) {
	defer trapStr(&result)
	coreMu.Lock()
	defer coreMu.Unlock()

	state, ok := canvasOf(handle)
	if !ok {
		return nil
	}
	return cstr(state.Dtor_providerId())
}

//export ahp_canvas_title
//
// Returns the title, or NULL when the state carries none or the handle is
// invalid. Release a non-NULL result with ahp_string_free.
func ahp_canvas_title(handle C.uintptr_t) (result *C.char) {
	defer trapStr(&result)
	coreMu.Lock()
	defer coreMu.Unlock()

	state, ok := canvasOf(handle)
	if !ok {
		return nil
	}
	return optCStr(state.Dtor_title())
}

//export ahp_canvas_activity
//
// Returns the activity, or NULL when the state carries none or the handle is
// invalid. Release a non-NULL result with ahp_string_free.
func ahp_canvas_activity(handle C.uintptr_t) (result *C.char) {
	defer trapStr(&result)
	coreMu.Lock()
	defer coreMu.Unlock()

	state, ok := canvasOf(handle)
	if !ok {
		return nil
	}
	return optCStr(state.Dtor_activity())
}

//export ahp_canvas_content_uri
//
// Returns the contentUri, or NULL when the state carries none or the handle is
// invalid. Release a non-NULL result with ahp_string_free.
func ahp_canvas_content_uri(handle C.uintptr_t) (result *C.char) {
	defer trapStr(&result)
	coreMu.Lock()
	defer coreMu.Unlock()

	state, ok := canvasOf(handle)
	if !ok {
		return nil
	}
	return optCStr(state.Dtor_contentUri())
}

//export ahp_canvas_availability
//
// Returns AHP_AVAILABILITY_READY or AHP_AVAILABILITY_STALE, or -1 when the
// handle is invalid.
func ahp_canvas_availability(handle C.uintptr_t) (result C.int) {
	defer trapInt(&result, -1)
	coreMu.Lock()
	defer coreMu.Unlock()

	state, ok := canvasOf(handle)
	if !ok {
		return -1
	}
	if state.Dtor_availability().Is_Ready() {
		return 0
	}
	return 1
}

//export ahp_canvas_equals
//
// Structural equality over two Canvas states, decided by the extracted
// datatype's own equality. Returns 1 when equal, 0 when not, and -1 when
// either handle is invalid — distinguishing "not equal" from "cannot answer".
func ahp_canvas_equals(a C.uintptr_t, b C.uintptr_t) (result C.int) {
	defer trapInt(&result, -1)
	coreMu.Lock()
	defer coreMu.Unlock()

	left, ok := canvasOf(a)
	if !ok {
		return -1
	}
	right, ok := canvasOf(b)
	if !ok {
		return -1
	}
	if left.Equals(right) {
		return 1
	}
	return 0
}

//export ahp_canvas_valid
//
// Returns 1 when the handle names a live Canvas state, 0 otherwise. There is
// no other way for a C caller to test a handle.
func ahp_canvas_valid(handle C.uintptr_t) (result C.int) {
	defer trapInt(&result, 0)
	coreMu.Lock()
	defer coreMu.Unlock()

	if _, ok := canvasOf(handle); ok {
		return 1
	}
	return 0
}

//export ahp_canvas_release
//
// Releases a Canvas handle. Every handle returned by ahp_canvas_new or an
// ahp_canvas_apply_* call should be released exactly once. Releasing a handle
// that was never issued, or that was already released, is a defined no-op:
// handles are never reused, so a stale one cannot name a live state.
func ahp_canvas_release(handle C.uintptr_t) {
	defer func() { _ = recover() }()
	coreMu.Lock()
	defer coreMu.Unlock()

	delete(canvases, uintptr(handle))
}
