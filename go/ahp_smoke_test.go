// Smoke tests exercising the extracted verified core directly from Go.
//
// These prove the published module is alive as a library — reducers step, the
// JSON codec round-trips, and command identity is byte-exact — not merely that
// the generated source compiles.
//
// Copyright (c) Microsoft Corporation
// Copyright (c) 2026 Josh Mouch
// SPDX-License-Identifier: MIT

package ahp_test

import (
	"testing"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_Identity "github.com/joshmouch/ahp-go/ConfluxCommandIdentityCapability"
	m_ConfluxJsonText "github.com/joshmouch/ahp-go/ConfluxJsonText"
	m_Terminal "github.com/joshmouch/ahp-go/Terminal"
	_dafny "github.com/joshmouch/ahp-go/dafny"
)

func str(s string) _dafny.Sequence { return _dafny.UnicodeSeqOfUtf8Bytes(s) }

func goStr(s _dafny.Sequence) string { return s.VerbatimString(false) }

// TestCanvasReducerSteps drives the verified Canvas channel reducer.
func TestCanvasReducerSteps(t *testing.T) {
	state := m_Canvas.Companion_CanvasState_.Create_CanvasState_(
		str("canvas-1"), str("provider-1"),
		m_AhpSkeleton.Companion_Option_.Create_None_(),
		m_AhpSkeleton.Companion_Option_.Create_None_(),
		m_AhpSkeleton.Companion_Option_.Create_None_(),
		m_AhpSkeleton.Companion_Option_.Create_None_(),
		m_Canvas.Companion_CanvasAvailability_.Create_Ready_(),
	)
	if got := goStr(state.Dtor_canvasId()); got != "canvas-1" {
		t.Fatalf("canvasId = %q, want %q", got, "canvas-1")
	}
	if state.Dtor_title().Is_Some() {
		t.Fatal("fresh state should have no title")
	}

	action := m_Canvas.Companion_CanvasAction_.Create_Updated_(
		m_AhpSkeleton.Companion_Option_.Create_Some_(str("Design Review")),
		m_AhpSkeleton.Companion_Option_.Create_None_(),
		m_AhpSkeleton.Companion_Option_.Create_None_(),
		m_AhpSkeleton.Companion_Option_.Create_None_(),
	)

	next := (*m_Canvas.Companion_Default___.ApplyToCanvas(state, action, _dafny.IntOf(1)).
		IndexInt(0)).(m_Canvas.CanvasState)

	if !next.Dtor_title().Is_Some() {
		t.Fatal("Updated action should set a title")
	}
	title := goStr(next.Dtor_title().Dtor_value().(_dafny.Sequence))
	if title != "Design Review" {
		t.Fatalf("title = %q, want %q", title, "Design Review")
	}
	// Reducers are pure: the input state must be untouched.
	if state.Dtor_title().Is_Some() {
		t.Fatal("reducer mutated its input state")
	}
}

// TestTerminalFoldIsDeterministic folds a sequence of actions twice.
func TestTerminalFoldIsDeterministic(t *testing.T) {
	initial := m_Terminal.Companion_TerminalState_.Default()
	actions := _dafny.SeqOf()

	a := m_Terminal.Companion_Default___.Fold(initial, actions)
	b := m_Terminal.Companion_Default___.Fold(initial, actions)
	if !a.Equals(b) {
		t.Fatal("Fold is not deterministic over the same inputs")
	}
}

// TestJsonRoundTrip stringifies a Json value and parses it back.
func TestJsonRoundTrip(t *testing.T) {
	fields := _dafny.NewMapBuilder().ToMap().
		UpdateUnsafe(str("channel"), m_ConfluxCodec.Companion_Json_.Create_JStr_(str("canvas"))).
		UpdateUnsafe(str("seq"), m_ConfluxCodec.Companion_Json_.Create_JNum_(_dafny.IntOf(42)))
	value := m_ConfluxCodec.Companion_Json_.Create_JObj_(fields)

	text := m_ConfluxJsonText.Companion_Default___.Stringify(value)
	parsed := m_ConfluxJsonText.Companion_Default___.ParseJsonChecked(text)
	if !parsed.Is_Parsed() {
		t.Fatalf("failed to re-parse stringified JSON: %s", goStr(text))
	}
	if !parsed.Dtor_value().Equals(value) {
		t.Fatalf("round-trip changed the value; text was %s", goStr(text))
	}
}

// TestSha256KnownAnswer pins the digest against the published SHA-256 test vector.
func TestSha256KnownAnswer(t *testing.T) {
	digest := goStr(m_Identity.Companion_Default___.Sha256Digest(_dafny.SeqOfBytes([]byte("abc"))))
	const want = "sha256:ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	if digest != want {
		t.Fatalf("Sha256Digest(\"abc\") = %s, want %s", digest, want)
	}
}

// TestCanonicalJsonAcceptsOnlyCanonicalBytes proves the canonical-form check has
// teeth: the canonical encoding is admitted, a non-canonical spelling is refused.
func TestCanonicalJsonAcceptsOnlyCanonicalBytes(t *testing.T) {
	value := m_ConfluxCodec.Companion_Json_.Create_JStr_(str("hello"))

	canonical := m_Identity.Companion_Default___.CanonicalJsonBytes(value)
	if decoded := m_Identity.Companion_Default___.DecodeCanonicalJson(canonical); !decoded.Is_Some() {
		t.Fatal("canonical bytes were rejected")
	}

	padded := _dafny.SeqOfBytes([]byte(" \"hello\""))
	if decoded := m_Identity.Companion_Default___.DecodeCanonicalJson(padded); decoded.Is_Some() {
		t.Fatal("non-canonical (leading-space) bytes were accepted")
	}

	invalidUtf8 := _dafny.SeqOfBytes([]byte{0xff, 0xfe})
	if decoded := m_Identity.Companion_Default___.DecodeCanonicalJson(invalidUtf8); decoded.Is_Some() {
		t.Fatal("invalid UTF-8 was accepted")
	}
}
