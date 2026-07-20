// Go realization of the pure ConfluxCommandIdentityCapability extern functions.
//
// ConfluxJsonText remains the sole JSON text implementation; this adapter only
// supplies strict UTF-8 validation, NFC normalization and SHA-256. It mirrors the
// C# and JavaScript adapters shipped with the same verified core.
//
// NOTE: this file is hand-written host glue at the Dafny {:extern} trust boundary.
// It is NOT part of the formally verified core. See EXTERN.md.
//
// Copyright (c) Microsoft Corporation
// Copyright (c) 2026 Josh Mouch
// SPDX-License-Identifier: MIT

package ConfluxCommandIdentityCapability

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"

	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxJsonText "github.com/joshmouch/ahp-go/ConfluxJsonText"
	_dafny "github.com/joshmouch/ahp-go/dafny"
)

// CompanionStruct_Default___ supplies the extern module's static members.
type CompanionStruct_Default___ struct{}

// Companion_Default___ is the singleton companion the translated core calls into.
var Companion_Default___ = CompanionStruct_Default___{}

// goBytes materializes a Dafny seq<byte> as a Go byte slice.
func goBytes(seq _dafny.Sequence) []byte {
	n := seq.Cardinality()
	out := make([]byte, n)
	for i := uint32(0); i < n; i++ {
		switch v := seq.Select(i).(type) {
		case byte:
			out[i] = v
		case _dafny.Int:
			out[i] = byte(v.Int64() & 0xff)
		}
	}
	return out
}

// CanonicalJsonBytes renders a Json value as NFC-normalized UTF-8 bytes.
func (_static *CompanionStruct_Default___) CanonicalJsonBytes(value m_ConfluxCodec.Json) _dafny.Sequence {
	text := m_ConfluxJsonText.Companion_Default___.Stringify(value).VerbatimString(false)
	return _dafny.SeqOfBytes([]byte(norm.NFC.String(text)))
}

// DecodeCanonicalJson parses bytes as JSON and admits them only when they are
// already in canonical form: strict UTF-8, and byte-identical to a re-render.
func (_static *CompanionStruct_Default___) DecodeCanonicalJson(b _dafny.Sequence) m_ConfluxCodec.Option {
	none := m_ConfluxCodec.Companion_Option_.Create_None_()
	input := goBytes(b)
	if !utf8.Valid(input) {
		return none
	}
	parsed := m_ConfluxJsonText.Companion_Default___.ParseJsonChecked(
		_dafny.UnicodeSeqOfUtf8Bytes(string(input)))
	if !parsed.Is_Parsed() {
		return none
	}
	value := parsed.Dtor_value()
	if !bytes.Equal(goBytes(Companion_Default___.CanonicalJsonBytes(value)), input) {
		return none
	}
	return m_ConfluxCodec.Companion_Option_.Create_Some_(value)
}

// Sha256Digest returns the "sha256:"-prefixed lowercase hex digest of the bytes.
func (_static *CompanionStruct_Default___) Sha256Digest(b _dafny.Sequence) _dafny.Sequence {
	sum := sha256.Sum256(goBytes(b))
	return _dafny.UnicodeSeqOfUtf8Bytes("sha256:" + hex.EncodeToString(sum[:]))
}
