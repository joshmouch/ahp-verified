// Package ConfluxJsonText
// Dafny module ConfluxJsonText compiled into Go

package ConfluxJsonText

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-verified/go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-verified/go/Changeset"
	m_Chat "github.com/joshmouch/ahp-verified/go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-verified/go/ClientMain"
	m_ConfluxBatchRoute "github.com/joshmouch/ahp-verified/go/ConfluxBatchRoute"
	m_ConfluxBudgetConvergence "github.com/joshmouch/ahp-verified/go/ConfluxBudgetConvergence"
	m_ConfluxChannel "github.com/joshmouch/ahp-verified/go/ConfluxChannel"
	m_ConfluxChannelManifest "github.com/joshmouch/ahp-verified/go/ConfluxChannelManifest"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxCommand "github.com/joshmouch/ahp-verified/go/ConfluxCommand"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-verified/go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxConvergence "github.com/joshmouch/ahp-verified/go/ConfluxConvergence"
	m_ConfluxDecimalText "github.com/joshmouch/ahp-verified/go/ConfluxDecimalText"
	m_ConfluxDedupe "github.com/joshmouch/ahp-verified/go/ConfluxDedupe"
	m_ConfluxDelimited "github.com/joshmouch/ahp-verified/go/ConfluxDelimited"
	m_ConfluxDurableHistory "github.com/joshmouch/ahp-verified/go/ConfluxDurableHistory"
	m_ConfluxExtern "github.com/joshmouch/ahp-verified/go/ConfluxExtern"
	m_ConfluxFilterKeys "github.com/joshmouch/ahp-verified/go/ConfluxFilterKeys"
	m_ConfluxFixpoint "github.com/joshmouch/ahp-verified/go/ConfluxFixpoint"
	m_ConfluxGenericCodec "github.com/joshmouch/ahp-verified/go/ConfluxGenericCodec"
	m_ConfluxJoin "github.com/joshmouch/ahp-verified/go/ConfluxJoin"
	m_ConfluxKeyedOps "github.com/joshmouch/ahp-verified/go/ConfluxKeyedOps"
	m_ConfluxMapValues "github.com/joshmouch/ahp-verified/go/ConfluxMapValues"
	m_ConfluxMerge "github.com/joshmouch/ahp-verified/go/ConfluxMerge"
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-verified/go/ConfluxOrderedLog"
	m_ConfluxProduct "github.com/joshmouch/ahp-verified/go/ConfluxProduct"
	m_ConfluxPrune "github.com/joshmouch/ahp-verified/go/ConfluxPrune"
	m_ConfluxResolve "github.com/joshmouch/ahp-verified/go/ConfluxResolve"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-verified/go/ConfluxResourceCeilings"
	m_ConfluxRoute "github.com/joshmouch/ahp-verified/go/ConfluxRoute"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-verified/go/ConfluxSeqRoute"
	m_ConfluxSeqRouteMerge "github.com/joshmouch/ahp-verified/go/ConfluxSeqRouteMerge"
	m_ConfluxStateMachine "github.com/joshmouch/ahp-verified/go/ConfluxStateMachine"
	m_ConfluxSupervisedOperationResult "github.com/joshmouch/ahp-verified/go/ConfluxSupervisedOperationResult"
	m_ConfluxTree "github.com/joshmouch/ahp-verified/go/ConfluxTree"
	m_ConfluxWatermark "github.com/joshmouch/ahp-verified/go/ConfluxWatermark"
	m_ResourceWatch "github.com/joshmouch/ahp-verified/go/ResourceWatch"
	m_Session "github.com/joshmouch/ahp-verified/go/Session"
	m__System "github.com/joshmouch/ahp-verified/go/System_"
	m_Terminal "github.com/joshmouch/ahp-verified/go/Terminal"
	_dafny "github.com/joshmouch/ahp-verified/go/dafny"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_ConfluxCodec.Dummy__
var _ m_ConfluxContract.Dummy__
var _ m_AhpSkeleton.Dummy__
var _ m_ResourceWatch.Dummy__
var _ m_Canvas.Dummy__
var _ m_ConfluxOperators.Dummy__
var _ m_ConfluxPrune.Dummy__
var _ m_ConfluxSeqRoute.Dummy__
var _ m_Changeset.Dummy__
var _ m_Annotations.Dummy__
var _ m_Terminal.Dummy__
var _ m_Session.Dummy__
var _ m_ConfluxOrderedLog.Dummy__
var _ m_Chat.Dummy__
var _ m_ClientMain.Dummy__
var _ m_ConfluxSemanticGraphIdentity.Dummy__
var _ m_ConfluxContentAddress.Dummy__
var _ m_ConfluxSemanticGraphContract.Dummy__
var _ m_ConfluxResourceCeilings.Dummy__
var _ m_ConfluxSupervisedOperationResult.Dummy__
var _ m_ConfluxTree.Dummy__
var _ m_ConfluxMerge.Dummy__
var _ m_ConfluxRoute.Dummy__
var _ m_ConfluxMapValues.Dummy__
var _ m_ConfluxDelimited.Dummy__
var _ m_ConfluxFixpoint.Dummy__
var _ m_ConfluxGenericCodec.Dummy__
var _ m_ConfluxResolve.Dummy__
var _ m_ConfluxWatermark.Dummy__
var _ m_ConfluxStateMachine.Dummy__
var _ m_ConfluxProduct.Dummy__
var _ m_ConfluxJoin.Dummy__
var _ m_ConfluxDedupe.Dummy__
var _ m_ConfluxBatchRoute.Dummy__
var _ m_ConfluxSeqRouteMerge.Dummy__
var _ m_ConfluxKeyedOps.Dummy__
var _ m_ConfluxFilterKeys.Dummy__
var _ m_ConfluxChannel.Dummy__
var _ m_ConfluxDurableHistory.Dummy__
var _ m_ConfluxCommand.Dummy__
var _ m_ConfluxChannelManifest.Dummy__
var _ m_ConfluxConvergence.Dummy__
var _ m_ConfluxBudgetConvergence.Dummy__
var _ m_ConfluxExtern.Dummy__
var _ m_ConfluxDecimalText.Dummy__

type Dummy__ struct{}

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "ConfluxJsonText.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) IsWs(c _dafny.CodePoint) bool {
	return ((((c) == (_dafny.CodePoint(' '))) || ((c) == (_dafny.CodePoint('\n')))) || ((c) == (_dafny.CodePoint('\t')))) || ((c) == (_dafny.CodePoint('\r')))
}
func (_static *CompanionStruct_Default___) IsDigit(c _dafny.CodePoint) bool {
	return ((_dafny.CodePoint('0')) <= (c)) && ((c) <= (_dafny.CodePoint('9')))
}
func (_static *CompanionStruct_Default___) SkipWs(s _dafny.Sequence, pos _dafny.Int) _dafny.Int {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (Companion_Default___.IsWs((s).Select((pos).Uint32()).(_dafny.CodePoint))) {
		var _in0 _dafny.Sequence = s
		_ = _in0
		var _in1 _dafny.Int = (pos).Plus(_dafny.One)
		_ = _in1
		s = _in0
		pos = _in1
		goto TAIL_CALL_START
	} else {
		return pos
	}
}
func (_static *CompanionStruct_Default___) HexVal(c _dafny.CodePoint) _dafny.Int {
	if ((_dafny.CodePoint('0')) <= (c)) && ((c) <= (_dafny.CodePoint('9'))) {
		return (_dafny.IntOfInt32(rune(c))).Minus(_dafny.IntOfInt32(rune(_dafny.CodePoint('0'))))
	} else if ((_dafny.CodePoint('a')) <= (c)) && ((c) <= (_dafny.CodePoint('f'))) {
		return ((_dafny.IntOfInt32(rune(c))).Minus(_dafny.IntOfInt32(rune(_dafny.CodePoint('a'))))).Plus(_dafny.IntOfInt64(10))
	} else if ((_dafny.CodePoint('A')) <= (c)) && ((c) <= (_dafny.CodePoint('F'))) {
		return ((_dafny.IntOfInt32(rune(c))).Minus(_dafny.IntOfInt32(rune(_dafny.CodePoint('A'))))).Plus(_dafny.IntOfInt64(10))
	} else {
		return _dafny.IntOfInt64(-1)
	}
}
func (_static *CompanionStruct_Default___) HexDigit(n _dafny.Int) _dafny.CodePoint {
	if (n).Cmp(_dafny.IntOfInt64(10)) < 0 {
		return _dafny.CodePoint(((_dafny.IntOfInt32(rune(_dafny.CodePoint('0')))).Plus(n)).Int32())
	} else {
		return _dafny.CodePoint(((_dafny.IntOfInt32(rune(_dafny.CodePoint('A')))).Plus((n).Minus(_dafny.IntOfInt64(10)))).Int32())
	}
}
func (_static *CompanionStruct_Default___) Hex4(s _dafny.Sequence, p _dafny.Int) _dafny.Int {
	var _0_d0 _dafny.Int = Companion_Default___.HexVal((s).Select((p).Uint32()).(_dafny.CodePoint))
	_ = _0_d0
	var _1_d1 _dafny.Int = Companion_Default___.HexVal((s).Select(((p).Plus(_dafny.One)).Uint32()).(_dafny.CodePoint))
	_ = _1_d1
	var _2_d2 _dafny.Int = Companion_Default___.HexVal((s).Select(((p).Plus(_dafny.IntOfInt64(2))).Uint32()).(_dafny.CodePoint))
	_ = _2_d2
	var _3_d3 _dafny.Int = Companion_Default___.HexVal((s).Select(((p).Plus(_dafny.IntOfInt64(3))).Uint32()).(_dafny.CodePoint))
	_ = _3_d3
	if ((((_0_d0).Sign() == -1) || ((_1_d1).Sign() == -1)) || ((_2_d2).Sign() == -1)) || ((_3_d3).Sign() == -1) {
		return _dafny.IntOfInt64(-1)
	} else {
		return ((((_0_d0).Times(_dafny.IntOfInt64(4096))).Plus((_1_d1).Times(_dafny.IntOfInt64(256)))).Plus((_2_d2).Times(_dafny.IntOfInt64(16)))).Plus(_3_d3)
	}
}
func (_static *CompanionStruct_Default___) IsScalar(v _dafny.Int) bool {
	return (((v).Sign() != -1) && ((v).Cmp(_dafny.IntOfInt64(55296)) < 0)) || (((_dafny.IntOfInt64(57344)).Cmp(v) <= 0) && ((v).Cmp(_dafny.IntOfInt64(1114111)) <= 0))
}
func (_static *CompanionStruct_Default___) ScalarToChar(v _dafny.Int) _dafny.CodePoint {
	if Companion_Default___.IsScalar(v) {
		return _dafny.CodePoint((v).Int32())
	} else {
		return _dafny.CodePoint('\U0000FFFD')
	}
}
func (_static *CompanionStruct_Default___) ScanDigits(s _dafny.Sequence, pos _dafny.Int, acc _dafny.Int, count _dafny.Int) DigitCursor {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (Companion_Default___.IsDigit((s).Select((pos).Uint32()).(_dafny.CodePoint))) {
		var _in0 _dafny.Sequence = s
		_ = _in0
		var _in1 _dafny.Int = (pos).Plus(_dafny.One)
		_ = _in1
		var _in2 _dafny.Int = ((acc).Times(_dafny.IntOfInt64(10))).Plus((_dafny.IntOfInt32(rune((s).Select((pos).Uint32()).(_dafny.CodePoint)))).Minus(_dafny.IntOfInt32(rune(_dafny.CodePoint('0')))))
		_ = _in2
		var _in3 _dafny.Int = (count).Plus(_dafny.One)
		_ = _in3
		s = _in0
		pos = _in1
		acc = _in2
		count = _in3
		goto TAIL_CALL_START
	} else {
		return Companion_DigitCursor_.Create_DigitCursor_(acc, pos, count)
	}
}
func (_static *CompanionStruct_Default___) ParseStringTail(s _dafny.Sequence, pos _dafny.Int, acc _dafny.Sequence) StringCursor {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) >= 0 {
		return Companion_StringCursor_.Create_StringCursor_(acc, pos, false)
	} else if ((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('"')) {
		return Companion_StringCursor_.Create_StringCursor_(acc, (pos).Plus(_dafny.One), true)
	} else if ((s).Select((pos).Uint32()).(_dafny.CodePoint)) != (_dafny.CodePoint('\\')) /* dircomp */ {
		if (_dafny.IntOfInt32(rune((s).Select((pos).Uint32()).(_dafny.CodePoint)))).Cmp(_dafny.IntOfInt64(32)) < 0 {
			return Companion_StringCursor_.Create_StringCursor_(acc, pos, false)
		} else {
			var _in0 _dafny.Sequence = s
			_ = _in0
			var _in1 _dafny.Int = (pos).Plus(_dafny.One)
			_ = _in1
			var _in2 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf((s).Select((pos).Uint32()).(_dafny.CodePoint)))
			_ = _in2
			s = _in0
			pos = _in1
			acc = _in2
			goto TAIL_CALL_START
		}
	} else if ((pos).Plus(_dafny.One)).Cmp(_dafny.IntOfUint32((s).Cardinality())) >= 0 {
		return Companion_StringCursor_.Create_StringCursor_(acc, pos, false)
	} else {
		var _0_e _dafny.CodePoint = (s).Select(((pos).Plus(_dafny.One)).Uint32()).(_dafny.CodePoint)
		_ = _0_e
		if (((_0_e) == (_dafny.CodePoint('"'))) || ((_0_e) == (_dafny.CodePoint('\\')))) || ((_0_e) == (_dafny.CodePoint('/'))) {
			var _in3 _dafny.Sequence = s
			_ = _in3
			var _in4 _dafny.Int = (pos).Plus(_dafny.IntOfInt64(2))
			_ = _in4
			var _in5 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf(_0_e))
			_ = _in5
			s = _in3
			pos = _in4
			acc = _in5
			goto TAIL_CALL_START
		} else if (_0_e) == (_dafny.CodePoint('n')) {
			var _in6 _dafny.Sequence = s
			_ = _in6
			var _in7 _dafny.Int = (pos).Plus(_dafny.IntOfInt64(2))
			_ = _in7
			var _in8 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf(_dafny.CodePoint('\n')))
			_ = _in8
			s = _in6
			pos = _in7
			acc = _in8
			goto TAIL_CALL_START
		} else if (_0_e) == (_dafny.CodePoint('t')) {
			var _in9 _dafny.Sequence = s
			_ = _in9
			var _in10 _dafny.Int = (pos).Plus(_dafny.IntOfInt64(2))
			_ = _in10
			var _in11 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf(_dafny.CodePoint('\t')))
			_ = _in11
			s = _in9
			pos = _in10
			acc = _in11
			goto TAIL_CALL_START
		} else if (_0_e) == (_dafny.CodePoint('r')) {
			var _in12 _dafny.Sequence = s
			_ = _in12
			var _in13 _dafny.Int = (pos).Plus(_dafny.IntOfInt64(2))
			_ = _in13
			var _in14 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf(_dafny.CodePoint('\r')))
			_ = _in14
			s = _in12
			pos = _in13
			acc = _in14
			goto TAIL_CALL_START
		} else if (_0_e) == (_dafny.CodePoint('b')) {
			var _in15 _dafny.Sequence = s
			_ = _in15
			var _in16 _dafny.Int = (pos).Plus(_dafny.IntOfInt64(2))
			_ = _in16
			var _in17 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf(_dafny.CodePoint('\U00000008')))
			_ = _in17
			s = _in15
			pos = _in16
			acc = _in17
			goto TAIL_CALL_START
		} else if (_0_e) == (_dafny.CodePoint('f')) {
			var _in18 _dafny.Sequence = s
			_ = _in18
			var _in19 _dafny.Int = (pos).Plus(_dafny.IntOfInt64(2))
			_ = _in19
			var _in20 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf(_dafny.CodePoint('\U0000000C')))
			_ = _in20
			s = _in18
			pos = _in19
			acc = _in20
			goto TAIL_CALL_START
		} else if (((_0_e) == (_dafny.CodePoint('u'))) && (((pos).Plus(_dafny.IntOfInt64(6))).Cmp(_dafny.IntOfUint32((s).Cardinality())) <= 0)) && ((Companion_Default___.Hex4(s, (pos).Plus(_dafny.IntOfInt64(2)))).Sign() != -1) {
			var _1_code _dafny.Int = Companion_Default___.Hex4(s, (pos).Plus(_dafny.IntOfInt64(2)))
			_ = _1_code
			if ((((((_dafny.IntOfInt64(55296)).Cmp(_1_code) <= 0) && ((_1_code).Cmp(_dafny.IntOfInt64(56319)) <= 0)) && (((pos).Plus(_dafny.IntOfInt64(12))).Cmp(_dafny.IntOfUint32((s).Cardinality())) <= 0)) && (((s).Select(((pos).Plus(_dafny.IntOfInt64(6))).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('\\')))) && (((s).Select(((pos).Plus(_dafny.IntOfInt64(7))).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('u')))) && (((_dafny.IntOfInt64(56320)).Cmp(Companion_Default___.Hex4(s, (pos).Plus(_dafny.IntOfInt64(8)))) <= 0) && ((Companion_Default___.Hex4(s, (pos).Plus(_dafny.IntOfInt64(8)))).Cmp(_dafny.IntOfInt64(57343)) <= 0)) {
				var _2_lo _dafny.Int = Companion_Default___.Hex4(s, (pos).Plus(_dafny.IntOfInt64(8)))
				_ = _2_lo
				var _in21 _dafny.Sequence = s
				_ = _in21
				var _in22 _dafny.Int = (pos).Plus(_dafny.IntOfInt64(12))
				_ = _in22
				var _in23 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf(Companion_Default___.ScalarToChar(((_dafny.IntOfInt64(65536)).Plus(((_1_code).Minus(_dafny.IntOfInt64(55296))).Times(_dafny.IntOfInt64(1024)))).Plus((_2_lo).Minus(_dafny.IntOfInt64(56320))))))
				_ = _in23
				s = _in21
				pos = _in22
				acc = _in23
				goto TAIL_CALL_START
			} else {
				var _in24 _dafny.Sequence = s
				_ = _in24
				var _in25 _dafny.Int = (pos).Plus(_dafny.IntOfInt64(6))
				_ = _in25
				var _in26 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf(Companion_Default___.ScalarToChar(_1_code)))
				_ = _in26
				s = _in24
				pos = _in25
				acc = _in26
				goto TAIL_CALL_START
			}
		} else {
			return Companion_StringCursor_.Create_StringCursor_(acc, pos, false)
		}
	}
}
func (_static *CompanionStruct_Default___) ParseStringModel(s _dafny.Sequence, pos _dafny.Int) StringCursor {
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('"'))) {
		return Companion_Default___.ParseStringTail(s, (pos).Plus(_dafny.One), _dafny.UnicodeSeqOfUtf8Bytes(""))
	} else {
		return Companion_StringCursor_.Create_StringCursor_(_dafny.UnicodeSeqOfUtf8Bytes(""), pos, false)
	}
}
func (_static *CompanionStruct_Default___) ParseNumberModel(s _dafny.Sequence, pos0 _dafny.Int) JsonCursor {
	var _0_neg bool = ((pos0).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos0).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('-')))
	_ = _0_neg
	var _1_first _dafny.Int = (func() _dafny.Int {
		if _0_neg {
			return (pos0).Plus(_dafny.One)
		}
		return pos0
	})()
	_ = _1_first
	var _2_whole DigitCursor = Companion_Default___.ScanDigits(s, _1_first, _dafny.Zero, _dafny.Zero)
	_ = _2_whole
	if (((_2_whole).Dtor_count()).Sign() == 0) || ((((_2_whole).Dtor_count()).Cmp(_dafny.One) > 0) && (((s).Select((_1_first).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('0')))) {
		return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JNull_(), (_2_whole).Dtor_pos(), false)
	} else {
		var _3_frac DigitCursor = (func() DigitCursor {
			if (((_2_whole).Dtor_pos()).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select(((_2_whole).Dtor_pos()).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('.'))) {
				return Companion_Default___.ScanDigits(s, ((_2_whole).Dtor_pos()).Plus(_dafny.One), (_2_whole).Dtor_value(), _dafny.Zero)
			}
			return Companion_DigitCursor_.Create_DigitCursor_((_2_whole).Dtor_value(), (_2_whole).Dtor_pos(), _dafny.Zero)
		})()
		_ = _3_frac
		var _4_hadDot bool = (((_2_whole).Dtor_pos()).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select(((_2_whole).Dtor_pos()).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('.')))
		_ = _4_hadDot
		if (_4_hadDot) && (((_3_frac).Dtor_count()).Sign() == 0) {
			return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JNull_(), (_3_frac).Dtor_pos(), false)
		} else {
			var _5_expMark bool = (((_3_frac).Dtor_pos()).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && ((((s).Select(((_3_frac).Dtor_pos()).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('e'))) || (((s).Select(((_3_frac).Dtor_pos()).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('E'))))
			_ = _5_expMark
			var _6_expSignPos _dafny.Int = (func() _dafny.Int {
				if _5_expMark {
					return ((_3_frac).Dtor_pos()).Plus(_dafny.One)
				}
				return (_3_frac).Dtor_pos()
			})()
			_ = _6_expSignPos
			var _7_expNeg bool = ((_6_expSignPos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((_6_expSignPos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('-')))
			_ = _7_expNeg
			var _8_expPlus bool = ((_6_expSignPos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((_6_expSignPos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('+')))
			_ = _8_expPlus
			var _9_expFirst _dafny.Int = (func() _dafny.Int {
				if (_7_expNeg) || (_8_expPlus) {
					return (_6_expSignPos).Plus(_dafny.One)
				}
				return _6_expSignPos
			})()
			_ = _9_expFirst
			var _10_expDigits DigitCursor = (func() DigitCursor {
				if _5_expMark {
					return Companion_Default___.ScanDigits(s, _9_expFirst, _dafny.Zero, _dafny.Zero)
				}
				return Companion_DigitCursor_.Create_DigitCursor_(_dafny.Zero, (_3_frac).Dtor_pos(), _dafny.Zero)
			})()
			_ = _10_expDigits
			if (_5_expMark) && (((_10_expDigits).Dtor_count()).Sign() == 0) {
				return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JNull_(), (_10_expDigits).Dtor_pos(), false)
			} else {
				var _11_mantissa _dafny.Int = (func() _dafny.Int {
					if _0_neg {
						return (_dafny.Zero).Minus((_3_frac).Dtor_value())
					}
					return (_3_frac).Dtor_value()
				})()
				_ = _11_mantissa
				var _12_exponent _dafny.Int = ((_dafny.Zero).Minus((_3_frac).Dtor_count())).Plus((func() _dafny.Int {
					if _7_expNeg {
						return (_dafny.Zero).Minus((_10_expDigits).Dtor_value())
					}
					return (_10_expDigits).Dtor_value()
				})())
				_ = _12_exponent
				var _13_value m_ConfluxCodec.Json = (func() m_ConfluxCodec.Json {
					if (!(_4_hadDot)) && (!(_5_expMark)) {
						return m_ConfluxCodec.Companion_Json_.Create_JNum_(_11_mantissa)
					}
					return m_ConfluxCodec.Companion_Json_.Create_JDec_(_11_mantissa, _12_exponent)
				})()
				_ = _13_value
				return Companion_JsonCursor_.Create_JsonCursor_(_13_value, (_10_expDigits).Dtor_pos(), true)
			}
		}
	}
}
func (_static *CompanionStruct_Default___) ParseValueModel(s _dafny.Sequence, pos0 _dafny.Int, budget _dafny.Int) JsonCursor {
	var _0_pos _dafny.Int = Companion_Default___.SkipWs(s, pos0)
	_ = _0_pos
	if ((budget).Sign() == 0) || ((_0_pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) >= 0) {
		return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JNull_(), _0_pos, false)
	} else {
		var _1_c _dafny.CodePoint = (s).Select((_0_pos).Uint32()).(_dafny.CodePoint)
		_ = _1_c
		if (_1_c) == (_dafny.CodePoint('"')) {
			var _2_r StringCursor = Companion_Default___.ParseStringModel(s, _0_pos)
			_ = _2_r
			return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JStr_((_2_r).Dtor_value()), (_2_r).Dtor_pos(), (_2_r).Dtor_ok())
		} else if (_1_c) == (_dafny.CodePoint('{')) {
			return Companion_Default___.ParseObjectModel(s, _0_pos, budget)
		} else if (_1_c) == (_dafny.CodePoint('[')) {
			return Companion_Default___.ParseArrayModel(s, _0_pos, budget)
		} else if (((_1_c) == (_dafny.CodePoint('t'))) && (((_0_pos).Plus(_dafny.IntOfInt64(4))).Cmp(_dafny.IntOfUint32((s).Cardinality())) <= 0)) && (_dafny.Companion_Sequence_.Equal((s).Subsequence((_0_pos).Uint32(), ((_0_pos).Plus(_dafny.IntOfInt64(4))).Uint32()), _dafny.UnicodeSeqOfUtf8Bytes("true"))) {
			return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JBool_(true), (_0_pos).Plus(_dafny.IntOfInt64(4)), true)
		} else if (((_1_c) == (_dafny.CodePoint('f'))) && (((_0_pos).Plus(_dafny.IntOfInt64(5))).Cmp(_dafny.IntOfUint32((s).Cardinality())) <= 0)) && (_dafny.Companion_Sequence_.Equal((s).Subsequence((_0_pos).Uint32(), ((_0_pos).Plus(_dafny.IntOfInt64(5))).Uint32()), _dafny.UnicodeSeqOfUtf8Bytes("false"))) {
			return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JBool_(false), (_0_pos).Plus(_dafny.IntOfInt64(5)), true)
		} else if (((_1_c) == (_dafny.CodePoint('n'))) && (((_0_pos).Plus(_dafny.IntOfInt64(4))).Cmp(_dafny.IntOfUint32((s).Cardinality())) <= 0)) && (_dafny.Companion_Sequence_.Equal((s).Subsequence((_0_pos).Uint32(), ((_0_pos).Plus(_dafny.IntOfInt64(4))).Uint32()), _dafny.UnicodeSeqOfUtf8Bytes("null"))) {
			return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JNull_(), (_0_pos).Plus(_dafny.IntOfInt64(4)), true)
		} else if ((_1_c) == (_dafny.CodePoint('-'))) || (Companion_Default___.IsDigit(_1_c)) {
			return Companion_Default___.ParseNumberModel(s, _0_pos)
		} else {
			return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JNull_(), (_0_pos).Plus(_dafny.One), false)
		}
	}
}
func (_static *CompanionStruct_Default___) ParseArrayModel(s _dafny.Sequence, pos _dafny.Int, budget _dafny.Int) JsonCursor {
	var _0_next _dafny.Int = Companion_Default___.SkipWs(s, (pos).Plus(_dafny.One))
	_ = _0_next
	if ((_0_next).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((_0_next).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint(']'))) {
		return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf()), (_0_next).Plus(_dafny.One), true)
	} else {
		return Companion_Default___.ParseArrayTailModel(s, _0_next, budget, _dafny.SeqOf())
	}
}
func (_static *CompanionStruct_Default___) ParseArrayTailModel(s _dafny.Sequence, pos _dafny.Int, budget _dafny.Int, acc _dafny.Sequence) JsonCursor {
	var _0_item JsonCursor = Companion_Default___.ParseValueModel(s, pos, (budget).Minus(_dafny.One))
	_ = _0_item
	if !((_0_item).Dtor_ok()) {
		return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JArr_(acc), (_0_item).Dtor_pos(), false)
	} else {
		var _1_next _dafny.Int = Companion_Default___.SkipWs(s, (_0_item).Dtor_pos())
		_ = _1_next
		if ((_1_next).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((_1_next).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint(']'))) {
			return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf((_0_item).Dtor_value()))), (_1_next).Plus(_dafny.One), true)
		} else if (((_1_next).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((_1_next).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint(',')))) && (((_1_next).Plus(_dafny.One)).Cmp(pos) > 0) {
			return Companion_Default___.ParseArrayTailModel(s, (_1_next).Plus(_dafny.One), budget, _dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf((_0_item).Dtor_value())))
		} else {
			return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.Companion_Sequence_.Concatenate(acc, _dafny.SeqOf((_0_item).Dtor_value()))), _1_next, false)
		}
	}
}
func (_static *CompanionStruct_Default___) ParseObjectModel(s _dafny.Sequence, pos _dafny.Int, budget _dafny.Int) JsonCursor {
	var _0_next _dafny.Int = Companion_Default___.SkipWs(s, (pos).Plus(_dafny.One))
	_ = _0_next
	if ((_0_next).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((_0_next).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('}'))) {
		return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap()), (_0_next).Plus(_dafny.One), true)
	} else {
		return Companion_Default___.ParseObjectTailModel(s, _0_next, budget, _dafny.NewMapBuilder().ToMap())
	}
}
func (_static *CompanionStruct_Default___) ParseObjectTailModel(s _dafny.Sequence, pos _dafny.Int, budget _dafny.Int, acc _dafny.Map) JsonCursor {
	var _0_next _dafny.Int = Companion_Default___.SkipWs(s, pos)
	_ = _0_next
	if ((_0_next).Cmp(_dafny.IntOfUint32((s).Cardinality())) >= 0) || (((s).Select((_0_next).Uint32()).(_dafny.CodePoint)) != (_dafny.CodePoint('"')) /* dircomp */) {
		return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JObj_(acc), _0_next, false)
	} else {
		var _1_key StringCursor = Companion_Default___.ParseStringModel(s, _0_next)
		_ = _1_key
		var _2_colon _dafny.Int = Companion_Default___.SkipWs(s, (_1_key).Dtor_pos())
		_ = _2_colon
		if ((!((_1_key).Dtor_ok())) || ((_2_colon).Cmp(_dafny.IntOfUint32((s).Cardinality())) >= 0)) || (((s).Select((_2_colon).Uint32()).(_dafny.CodePoint)) != (_dafny.CodePoint(':')) /* dircomp */) {
			return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JObj_(acc), _2_colon, false)
		} else {
			var _3_value JsonCursor = Companion_Default___.ParseValueModel(s, (_2_colon).Plus(_dafny.One), (budget).Minus(_dafny.One))
			_ = _3_value
			if !((_3_value).Dtor_ok()) {
				return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JObj_(acc), (_3_value).Dtor_pos(), false)
			} else {
				var _4_after _dafny.Int = Companion_Default___.SkipWs(s, (_3_value).Dtor_pos())
				_ = _4_after
				var _5_fields _dafny.Map = (acc).Update((_1_key).Dtor_value(), (_3_value).Dtor_value())
				_ = _5_fields
				if ((_4_after).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((_4_after).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('}'))) {
					return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_5_fields), (_4_after).Plus(_dafny.One), true)
				} else if (((_4_after).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((_4_after).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint(',')))) && (((_4_after).Plus(_dafny.One)).Cmp(pos) > 0) {
					return Companion_Default___.ParseObjectTailModel(s, (_4_after).Plus(_dafny.One), budget, _5_fields)
				} else {
					return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_5_fields), _4_after, false)
				}
			}
		}
	}
}
func (_static *CompanionStruct_Default___) ParseString(s _dafny.Sequence, pos0 _dafny.Int) (_dafny.Sequence, _dafny.Int, bool) {
	var val _dafny.Sequence = _dafny.EmptySeq
	_ = val
	var pos _dafny.Int = _dafny.Zero
	_ = pos
	var ok bool = false
	_ = ok
	val = _dafny.UnicodeSeqOfUtf8Bytes("")
	ok = false
	pos = pos0
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) >= 0) || (((s).Select((pos).Uint32()).(_dafny.CodePoint)) != (_dafny.CodePoint('"')) /* dircomp */) {
		return val, pos, ok
	}
	pos = (pos).Plus(_dafny.One)
	var _0_runStart _dafny.Int
	_ = _0_runStart
	_0_runStart = pos
	for ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) != (_dafny.CodePoint('"')) /* dircomp */) {
		if (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('\\'))) && (((pos).Plus(_dafny.One)).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) {
			val = _dafny.Companion_Sequence_.Concatenate(val, (s).Subsequence((_0_runStart).Uint32(), (pos).Uint32()))
			var _1_e _dafny.CodePoint
			_ = _1_e
			_1_e = (s).Select(((pos).Plus(_dafny.One)).Uint32()).(_dafny.CodePoint)
			if (((_1_e) == (_dafny.CodePoint('"'))) || ((_1_e) == (_dafny.CodePoint('\\')))) || ((_1_e) == (_dafny.CodePoint('/'))) {
				val = _dafny.Companion_Sequence_.Concatenate(val, _dafny.SeqOf(_1_e))
				pos = (pos).Plus(_dafny.IntOfInt64(2))
			} else if (_1_e) == (_dafny.CodePoint('n')) {
				val = _dafny.Companion_Sequence_.Concatenate(val, _dafny.SeqOf(_dafny.CodePoint('\n')))
				pos = (pos).Plus(_dafny.IntOfInt64(2))
			} else if (_1_e) == (_dafny.CodePoint('t')) {
				val = _dafny.Companion_Sequence_.Concatenate(val, _dafny.SeqOf(_dafny.CodePoint('\t')))
				pos = (pos).Plus(_dafny.IntOfInt64(2))
			} else if (_1_e) == (_dafny.CodePoint('r')) {
				val = _dafny.Companion_Sequence_.Concatenate(val, _dafny.SeqOf(_dafny.CodePoint('\r')))
				pos = (pos).Plus(_dafny.IntOfInt64(2))
			} else if (_1_e) == (_dafny.CodePoint('b')) {
				val = _dafny.Companion_Sequence_.Concatenate(val, _dafny.SeqOf(_dafny.CodePoint('\U00000008')))
				pos = (pos).Plus(_dafny.IntOfInt64(2))
			} else if (_1_e) == (_dafny.CodePoint('f')) {
				val = _dafny.Companion_Sequence_.Concatenate(val, _dafny.SeqOf(_dafny.CodePoint('\U0000000C')))
				pos = (pos).Plus(_dafny.IntOfInt64(2))
			} else if (((_1_e) == (_dafny.CodePoint('u'))) && (((pos).Plus(_dafny.IntOfInt64(6))).Cmp(_dafny.IntOfUint32((s).Cardinality())) <= 0)) && ((Companion_Default___.Hex4(s, (pos).Plus(_dafny.IntOfInt64(2)))).Sign() != -1) {
				var _2_code _dafny.Int
				_ = _2_code
				_2_code = Companion_Default___.Hex4(s, (pos).Plus(_dafny.IntOfInt64(2)))
				if (((((((_dafny.IntOfInt64(55296)).Cmp(_2_code) <= 0) && ((_2_code).Cmp(_dafny.IntOfInt64(56319)) <= 0)) && (((pos).Plus(_dafny.IntOfInt64(12))).Cmp(_dafny.IntOfUint32((s).Cardinality())) <= 0)) && (((s).Select(((pos).Plus(_dafny.IntOfInt64(6))).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('\\')))) && (((s).Select(((pos).Plus(_dafny.IntOfInt64(7))).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('u')))) && ((Companion_Default___.Hex4(s, (pos).Plus(_dafny.IntOfInt64(8)))).Cmp(_dafny.IntOfInt64(56320)) >= 0)) && ((Companion_Default___.Hex4(s, (pos).Plus(_dafny.IntOfInt64(8)))).Cmp(_dafny.IntOfInt64(57343)) <= 0) {
					var _3_lo _dafny.Int
					_ = _3_lo
					_3_lo = Companion_Default___.Hex4(s, (pos).Plus(_dafny.IntOfInt64(8)))
					val = _dafny.Companion_Sequence_.Concatenate(val, _dafny.SeqOf(Companion_Default___.ScalarToChar(((_dafny.IntOfInt64(65536)).Plus(((_2_code).Minus(_dafny.IntOfInt64(55296))).Times(_dafny.IntOfInt64(1024)))).Plus((_3_lo).Minus(_dafny.IntOfInt64(56320))))))
					pos = (pos).Plus(_dafny.IntOfInt64(12))
				} else {
					val = _dafny.Companion_Sequence_.Concatenate(val, _dafny.SeqOf(Companion_Default___.ScalarToChar(_2_code)))
					pos = (pos).Plus(_dafny.IntOfInt64(6))
				}
			} else {
				return val, pos, ok
			}
			_0_runStart = pos
		} else {
			if (_dafny.IntOfInt32(rune((s).Select((pos).Uint32()).(_dafny.CodePoint)))).Cmp(_dafny.IntOfInt64(32)) < 0 {
				return val, pos, ok
			}
			pos = (pos).Plus(_dafny.One)
		}
	}
	val = _dafny.Companion_Sequence_.Concatenate(val, (s).Subsequence((_0_runStart).Uint32(), (pos).Uint32()))
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('"'))) {
		pos = (pos).Plus(_dafny.One)
		ok = true
	}
	return val, pos, ok
}
func (_static *CompanionStruct_Default___) AdvanceOver(s _dafny.Sequence, pos _dafny.Int, n _dafny.Int) _dafny.Int {
	if ((pos).Plus(n)).Cmp(_dafny.IntOfUint32((s).Cardinality())) <= 0 {
		return (pos).Plus(n)
	} else {
		return _dafny.IntOfUint32((s).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) ParseNumber(s _dafny.Sequence, pos0 _dafny.Int) (m_ConfluxCodec.Json, _dafny.Int, bool) {
	var j m_ConfluxCodec.Json = m_ConfluxCodec.Companion_Json_.Default()
	_ = j
	var pos _dafny.Int = _dafny.Zero
	_ = pos
	var ok bool = false
	_ = ok
	pos = pos0
	ok = false
	var _0_neg bool
	_ = _0_neg
	_0_neg = false
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('-'))) {
		_0_neg = true
		pos = (pos).Plus(_dafny.One)
	}
	var _1_whole _dafny.Int
	_ = _1_whole
	_1_whole = _dafny.Zero
	var _2_sawDigit bool
	_ = _2_sawDigit
	_2_sawDigit = false
	var _3_firstDigit _dafny.Int
	_ = _3_firstDigit
	_3_firstDigit = pos
	for ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (Companion_Default___.IsDigit((s).Select((pos).Uint32()).(_dafny.CodePoint))) {
		_1_whole = ((_1_whole).Times(_dafny.IntOfInt64(10))).Plus((_dafny.IntOfInt32(rune((s).Select((pos).Uint32()).(_dafny.CodePoint)))).Minus(_dafny.IntOfInt32(rune(_dafny.CodePoint('0')))))
		_2_sawDigit = true
		pos = (pos).Plus(_dafny.One)
	}
	var _4_mantissa _dafny.Int
	_ = _4_mantissa
	_4_mantissa = _1_whole
	var _5_exp _dafny.Int
	_ = _5_exp
	_5_exp = _dafny.Zero
	var _6_decimalDigits _dafny.Int
	_ = _6_decimalDigits
	_6_decimalDigits = _dafny.Zero
	if (!(_2_sawDigit)) || ((((pos).Minus(_3_firstDigit)).Cmp(_dafny.One) > 0) && (((s).Select((_3_firstDigit).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('0')))) {
		j = m_ConfluxCodec.Companion_Json_.Create_JNull_()
		return j, pos, ok
	}
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('.'))) {
		pos = (pos).Plus(_dafny.One)
		var _7_fractionStart _dafny.Int
		_ = _7_fractionStart
		_7_fractionStart = pos
		for ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (Companion_Default___.IsDigit((s).Select((pos).Uint32()).(_dafny.CodePoint))) {
			_4_mantissa = ((_4_mantissa).Times(_dafny.IntOfInt64(10))).Plus((_dafny.IntOfInt32(rune((s).Select((pos).Uint32()).(_dafny.CodePoint)))).Minus(_dafny.IntOfInt32(rune(_dafny.CodePoint('0')))))
			_6_decimalDigits = (_6_decimalDigits).Plus(_dafny.One)
			_2_sawDigit = true
			pos = (pos).Plus(_dafny.One)
		}
		if (pos).Cmp(_7_fractionStart) == 0 {
			j = m_ConfluxCodec.Companion_Json_.Create_JNull_()
			return j, pos, ok
		}
		_5_exp = (_dafny.Zero).Minus(_6_decimalDigits)
	}
	var _8_explicitExp bool
	_ = _8_explicitExp
	_8_explicitExp = false
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && ((((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('e'))) || (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('E')))) {
		_8_explicitExp = true
		pos = (pos).Plus(_dafny.One)
		var _9_expNeg bool
		_ = _9_expNeg
		_9_expNeg = false
		if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && ((((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('+'))) || (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('-')))) {
			_9_expNeg = ((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('-'))
			pos = (pos).Plus(_dafny.One)
		}
		var _10_e _dafny.Int
		_ = _10_e
		_10_e = _dafny.Zero
		var _11_exponentStart _dafny.Int
		_ = _11_exponentStart
		_11_exponentStart = pos
		for ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (Companion_Default___.IsDigit((s).Select((pos).Uint32()).(_dafny.CodePoint))) {
			_10_e = ((_10_e).Times(_dafny.IntOfInt64(10))).Plus((_dafny.IntOfInt32(rune((s).Select((pos).Uint32()).(_dafny.CodePoint)))).Minus(_dafny.IntOfInt32(rune(_dafny.CodePoint('0')))))
			pos = (pos).Plus(_dafny.One)
		}
		if (pos).Cmp(_11_exponentStart) == 0 {
			j = m_ConfluxCodec.Companion_Json_.Create_JNull_()
			return j, pos, ok
		}
		_5_exp = (_5_exp).Plus((func() _dafny.Int {
			if _9_expNeg {
				return (_dafny.Zero).Minus(_10_e)
			}
			return _10_e
		})())
	}
	var _12_signed _dafny.Int
	_ = _12_signed
	if _0_neg {
		_12_signed = (_dafny.Zero).Minus(_4_mantissa)
	} else {
		_12_signed = _4_mantissa
	}
	if !(_2_sawDigit) {
		j = m_ConfluxCodec.Companion_Json_.Create_JNull_()
	} else if ((_6_decimalDigits).Sign() == 0) && (!(_8_explicitExp)) {
		j = m_ConfluxCodec.Companion_Json_.Create_JNum_(_12_signed)
	} else {
		j = m_ConfluxCodec.Companion_Json_.Create_JDec_(_12_signed, _5_exp)
	}
	ok = true
	return j, pos, ok
}
func (_static *CompanionStruct_Default___) ParseValue(s _dafny.Sequence, pos0 _dafny.Int, budget _dafny.Int) (m_ConfluxCodec.Json, _dafny.Int, bool) {
	var j m_ConfluxCodec.Json = m_ConfluxCodec.Companion_Json_.Default()
	_ = j
	var pos _dafny.Int = _dafny.Zero
	_ = pos
	var ok bool = false
	_ = ok
	pos = Companion_Default___.SkipWs(s, pos0)
	ok = false
	if ((budget).Sign() == 0) || ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) >= 0) {
		j = m_ConfluxCodec.Companion_Json_.Create_JNull_()
		return j, pos, ok
	}
	var _0_c _dafny.CodePoint
	_ = _0_c
	_0_c = (s).Select((pos).Uint32()).(_dafny.CodePoint)
	if (_0_c) == (_dafny.CodePoint('"')) {
		var _1_v _dafny.Sequence
		_ = _1_v
		var _2_p _dafny.Int
		_ = _2_p
		var _3_good bool
		_ = _3_good
		var _out0 _dafny.Sequence
		_ = _out0
		var _out1 _dafny.Int
		_ = _out1
		var _out2 bool
		_ = _out2
		_out0, _out1, _out2 = Companion_Default___.ParseString(s, pos)
		_1_v = _out0
		_2_p = _out1
		_3_good = _out2
		j = m_ConfluxCodec.Companion_Json_.Create_JStr_(_1_v)
		pos = _2_p
		ok = _3_good
	} else if (_0_c) == (_dafny.CodePoint('{')) {
		var _4_m _dafny.Map
		_ = _4_m
		var _5_p _dafny.Int
		_ = _5_p
		var _6_good bool
		_ = _6_good
		var _out3 _dafny.Map
		_ = _out3
		var _out4 _dafny.Int
		_ = _out4
		var _out5 bool
		_ = _out5
		_out3, _out4, _out5 = Companion_Default___.ParseObject(s, pos, budget)
		_4_m = _out3
		_5_p = _out4
		_6_good = _out5
		j = m_ConfluxCodec.Companion_Json_.Create_JObj_(_4_m)
		pos = _5_p
		ok = _6_good
	} else if (_0_c) == (_dafny.CodePoint('[')) {
		var _7_a _dafny.Sequence
		_ = _7_a
		var _8_p _dafny.Int
		_ = _8_p
		var _9_good bool
		_ = _9_good
		var _out6 _dafny.Sequence
		_ = _out6
		var _out7 _dafny.Int
		_ = _out7
		var _out8 bool
		_ = _out8
		_out6, _out7, _out8 = Companion_Default___.ParseArray(s, pos, budget)
		_7_a = _out6
		_8_p = _out7
		_9_good = _out8
		j = m_ConfluxCodec.Companion_Json_.Create_JArr_(_7_a)
		pos = _8_p
		ok = _9_good
	} else if (((_0_c) == (_dafny.CodePoint('t'))) && (((pos).Plus(_dafny.IntOfInt64(4))).Cmp(_dafny.IntOfUint32((s).Cardinality())) <= 0)) && (_dafny.Companion_Sequence_.Equal((s).Subsequence((pos).Uint32(), ((pos).Plus(_dafny.IntOfInt64(4))).Uint32()), _dafny.UnicodeSeqOfUtf8Bytes("true"))) {
		j = m_ConfluxCodec.Companion_Json_.Create_JBool_(true)
		pos = (pos).Plus(_dafny.IntOfInt64(4))
		ok = true
	} else if (((_0_c) == (_dafny.CodePoint('f'))) && (((pos).Plus(_dafny.IntOfInt64(5))).Cmp(_dafny.IntOfUint32((s).Cardinality())) <= 0)) && (_dafny.Companion_Sequence_.Equal((s).Subsequence((pos).Uint32(), ((pos).Plus(_dafny.IntOfInt64(5))).Uint32()), _dafny.UnicodeSeqOfUtf8Bytes("false"))) {
		j = m_ConfluxCodec.Companion_Json_.Create_JBool_(false)
		pos = (pos).Plus(_dafny.IntOfInt64(5))
		ok = true
	} else if (((_0_c) == (_dafny.CodePoint('n'))) && (((pos).Plus(_dafny.IntOfInt64(4))).Cmp(_dafny.IntOfUint32((s).Cardinality())) <= 0)) && (_dafny.Companion_Sequence_.Equal((s).Subsequence((pos).Uint32(), ((pos).Plus(_dafny.IntOfInt64(4))).Uint32()), _dafny.UnicodeSeqOfUtf8Bytes("null"))) {
		j = m_ConfluxCodec.Companion_Json_.Create_JNull_()
		pos = (pos).Plus(_dafny.IntOfInt64(4))
		ok = true
	} else if ((_0_c) == (_dafny.CodePoint('-'))) || (Companion_Default___.IsDigit(_0_c)) {
		var _10_n m_ConfluxCodec.Json
		_ = _10_n
		var _11_p _dafny.Int
		_ = _11_p
		var _12_good bool
		_ = _12_good
		var _out9 m_ConfluxCodec.Json
		_ = _out9
		var _out10 _dafny.Int
		_ = _out10
		var _out11 bool
		_ = _out11
		_out9, _out10, _out11 = Companion_Default___.ParseNumber(s, pos)
		_10_n = _out9
		_11_p = _out10
		_12_good = _out11
		j = _10_n
		pos = _11_p
		ok = _12_good
	} else {
		j = m_ConfluxCodec.Companion_Json_.Create_JNull_()
		pos = (pos).Plus(_dafny.One)
	}
	return j, pos, ok
}
func (_static *CompanionStruct_Default___) ParseObject(s _dafny.Sequence, pos0 _dafny.Int, budget _dafny.Int) (_dafny.Map, _dafny.Int, bool) {
	var m _dafny.Map = _dafny.EmptyMap
	_ = m
	var pos _dafny.Int = _dafny.Zero
	_ = pos
	var ok bool = false
	_ = ok
	m = _dafny.NewMapBuilder().ToMap()
	pos = pos0
	ok = false
	if (budget).Sign() == 0 {
		return m, pos, ok
	}
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('{'))) {
		pos = (pos).Plus(_dafny.One)
	}
	pos = Companion_Default___.SkipWs(s, pos)
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('}'))) {
		pos = (pos).Plus(_dafny.One)
		ok = true
		return m, pos, ok
	}
	{
		for (pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0 {
			{
				var _0_before _dafny.Int
				_ = _0_before
				_0_before = pos
				pos = Companion_Default___.SkipWs(s, pos)
				if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) >= 0) || (((s).Select((pos).Uint32()).(_dafny.CodePoint)) != (_dafny.CodePoint('"')) /* dircomp */) {
					goto L0
				}
				var _1_key _dafny.Sequence
				_ = _1_key
				var _2_p1 _dafny.Int
				_ = _2_p1
				var _3_keyOk bool
				_ = _3_keyOk
				var _out0 _dafny.Sequence
				_ = _out0
				var _out1 _dafny.Int
				_ = _out1
				var _out2 bool
				_ = _out2
				_out0, _out1, _out2 = Companion_Default___.ParseString(s, pos)
				_1_key = _out0
				_2_p1 = _out1
				_3_keyOk = _out2
				pos = _2_p1
				if !(_3_keyOk) {
					return m, pos, ok
				}
				pos = Companion_Default___.SkipWs(s, pos)
				if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint(':'))) {
					pos = (pos).Plus(_dafny.One)
				} else {
					goto L0
				}
				var _4_v m_ConfluxCodec.Json
				_ = _4_v
				var _5_p2 _dafny.Int
				_ = _5_p2
				var _6_valueOk bool
				_ = _6_valueOk
				var _out3 m_ConfluxCodec.Json
				_ = _out3
				var _out4 _dafny.Int
				_ = _out4
				var _out5 bool
				_ = _out5
				_out3, _out4, _out5 = Companion_Default___.ParseValue(s, pos, (budget).Minus(_dafny.One))
				_4_v = _out3
				_5_p2 = _out4
				_6_valueOk = _out5
				pos = _5_p2
				if !(_6_valueOk) {
					return m, pos, ok
				}
				m = (m).Update(_1_key, _4_v)
				pos = Companion_Default___.SkipWs(s, pos)
				if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint(','))) {
					pos = (pos).Plus(_dafny.One)
				} else if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('}'))) {
					pos = (pos).Plus(_dafny.One)
					ok = true
					goto L0
				} else {
					goto L0
				}
				if (pos).Cmp(_0_before) <= 0 {
					goto L0
				}
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	return m, pos, ok
}
func (_static *CompanionStruct_Default___) ParseArray(s _dafny.Sequence, pos0 _dafny.Int, budget _dafny.Int) (_dafny.Sequence, _dafny.Int, bool) {
	var a _dafny.Sequence = _dafny.EmptySeq
	_ = a
	var pos _dafny.Int = _dafny.Zero
	_ = pos
	var ok bool = false
	_ = ok
	a = _dafny.SeqOf()
	pos = pos0
	ok = false
	if (budget).Sign() == 0 {
		return a, pos, ok
	}
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('['))) {
		pos = (pos).Plus(_dafny.One)
	}
	pos = Companion_Default___.SkipWs(s, pos)
	if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint(']'))) {
		pos = (pos).Plus(_dafny.One)
		ok = true
		return a, pos, ok
	}
	{
		for (pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0 {
			{
				var _0_before _dafny.Int
				_ = _0_before
				_0_before = pos
				var _1_v m_ConfluxCodec.Json
				_ = _1_v
				var _2_p _dafny.Int
				_ = _2_p
				var _3_valueOk bool
				_ = _3_valueOk
				var _out0 m_ConfluxCodec.Json
				_ = _out0
				var _out1 _dafny.Int
				_ = _out1
				var _out2 bool
				_ = _out2
				_out0, _out1, _out2 = Companion_Default___.ParseValue(s, pos, (budget).Minus(_dafny.One))
				_1_v = _out0
				_2_p = _out1
				_3_valueOk = _out2
				pos = _2_p
				if !(_3_valueOk) {
					return a, pos, ok
				}
				a = _dafny.Companion_Sequence_.Concatenate(a, _dafny.SeqOf(_1_v))
				pos = Companion_Default___.SkipWs(s, pos)
				if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint(','))) {
					pos = (pos).Plus(_dafny.One)
				} else if ((pos).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0) && (((s).Select((pos).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint(']'))) {
					pos = (pos).Plus(_dafny.One)
					ok = true
					goto L1
				} else {
					goto L1
				}
				if (pos).Cmp(_0_before) <= 0 {
					goto L1
				}
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	return a, pos, ok
}
func (_static *CompanionStruct_Default___) ParseJsonModel(s _dafny.Sequence) ParseResult {
	var _0_parsed JsonCursor = Companion_Default___.ParseValueModel(s, _dafny.Zero, Companion_Default___.PARSE__BUDGET())
	_ = _0_parsed
	var _1_end _dafny.Int = Companion_Default___.SkipWs(s, (_0_parsed).Dtor_pos())
	_ = _1_end
	if ((_0_parsed).Dtor_ok()) && ((_1_end).Cmp(_dafny.IntOfUint32((s).Cardinality())) == 0) {
		return Companion_ParseResult_.Create_Parsed_((_0_parsed).Dtor_value())
	} else {
		return Companion_ParseResult_.Create_Invalid_()
	}
}
func (_static *CompanionStruct_Default___) ParseJsonChecked(s _dafny.Sequence) ParseResult {
	var result ParseResult = Companion_ParseResult_.Default()
	_ = result
	result = Companion_Default___.ParseJsonModel(s)
	return result
}
func (_static *CompanionStruct_Default___) ParseJson(s _dafny.Sequence) m_ConfluxCodec.Json {
	var j m_ConfluxCodec.Json = m_ConfluxCodec.Companion_Json_.Default()
	_ = j
	var _0_result ParseResult
	_ = _0_result
	var _out0 ParseResult
	_ = _out0
	_out0 = Companion_Default___.ParseJsonChecked(s)
	_0_result = _out0
	if (_0_result).Is_Parsed() {
		j = (_0_result).Dtor_value()
	} else {
		j = m_ConfluxCodec.Companion_Json_.Create_JNull_()
	}
	return j
}
func (_static *CompanionStruct_Default___) CharLt(a _dafny.CodePoint, b _dafny.CodePoint) bool {
	return (_dafny.IntOfInt32(rune(a))).Cmp(_dafny.IntOfInt32(rune(b))) < 0
}
func (_static *CompanionStruct_Default___) StrLt(a _dafny.Sequence, b _dafny.Sequence) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((a).Cardinality())).Sign() == 0 {
		return (_dafny.IntOfUint32((b).Cardinality())).Sign() == 1
	} else if (_dafny.IntOfUint32((b).Cardinality())).Sign() == 0 {
		return false
	} else if ((a).Select(0).(_dafny.CodePoint)) == ((b).Select(0).(_dafny.CodePoint)) {
		var _in0 _dafny.Sequence = (a).Drop(1)
		_ = _in0
		var _in1 _dafny.Sequence = (b).Drop(1)
		_ = _in1
		a = _in0
		b = _in1
		goto TAIL_CALL_START
	} else {
		return Companion_Default___.CharLt((a).Select(0).(_dafny.CodePoint), (b).Select(0).(_dafny.CodePoint))
	}
}
func (_static *CompanionStruct_Default___) Sorted(xs _dafny.Sequence) bool {
	return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, (_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One)), true, func(_forall_var_0 _dafny.Int) bool {
		var _0_i _dafny.Int
		_0_i = interface{}(_forall_var_0).(_dafny.Int)
		return !(((_0_i).Sign() != -1) && (((_0_i).Plus(_dafny.One)).Cmp(_dafny.IntOfUint32((xs).Cardinality())) < 0)) || (Companion_Default___.StrLt((xs).Select((_0_i).Uint32()).(_dafny.Sequence), (xs).Select(((_0_i).Plus(_dafny.One)).Uint32()).(_dafny.Sequence)))
	})
}
func (_static *CompanionStruct_Default___) Elems(xs _dafny.Sequence) _dafny.Set {
	var _0___accumulator _dafny.Set = _dafny.SetOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return (_dafny.SetOf()).Union(_0___accumulator)
	} else {
		_0___accumulator = (_0___accumulator).Union(_dafny.SetOf((xs).Select(0).(_dafny.Sequence)))
		var _in0 _dafny.Sequence = (xs).Drop(1)
		_ = _in0
		xs = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) SortedKeySetRuntime(s _dafny.Set) _dafny.Sequence {
	var keys _dafny.Sequence = _dafny.EmptySeq
	_ = keys
	if (s).Equals(_dafny.SetOf()) {
		keys = _dafny.SeqOf()
	} else {
		var _0_k _dafny.Sequence
		_ = _0_k
		{
			for _iter5 := _dafny.Iterate((s).Elements()); ; {
				_assign_such_that_0, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				_0_k = interface{}(_assign_such_that_0).(_dafny.Sequence)
				if ((s).Contains(_0_k)) && (_dafny.Quantifier((s).Elements(), true, func(_forall_var_0 _dafny.Sequence) bool {
					var _1_y _dafny.Sequence
					_1_y = interface{}(_forall_var_0).(_dafny.Sequence)
					return !((s).Contains(_1_y)) || (!(Companion_Default___.StrLt(_1_y, _0_k)))
				})) {
					goto L_ASSIGN_SUCH_THAT_0
				}
			}
			panic("assign-such-that search produced no value")
			goto L_ASSIGN_SUCH_THAT_0
		}
	L_ASSIGN_SUCH_THAT_0:
		var _2_tail _dafny.Sequence
		_ = _2_tail
		var _out0 _dafny.Sequence
		_ = _out0
		_out0 = Companion_Default___.SortedKeySetRuntime((s).Difference(_dafny.SetOf(_0_k)))
		_2_tail = _out0
		keys = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_0_k), _2_tail)
	}
	return keys
}
func (_static *CompanionStruct_Default___) InsertKey(k _dafny.Sequence, xs _dafny.Sequence) _dafny.Sequence {
	var ys _dafny.Sequence = _dafny.EmptySeq
	_ = ys
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		ys = _dafny.SeqOf(k)
	} else if Companion_Default___.StrLt(k, (xs).Select(0).(_dafny.Sequence)) {
		ys = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(k), xs)
	} else {
		var _0_tail _dafny.Sequence
		_ = _0_tail
		var _out0 _dafny.Sequence
		_ = _out0
		_out0 = Companion_Default___.InsertKey(k, (xs).Drop(1))
		_0_tail = _out0
		ys = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((xs).Select(0).(_dafny.Sequence)), _0_tail)
	}
	return ys
}
func (_static *CompanionStruct_Default___) SortedKeys(m _dafny.Map) _dafny.Sequence {
	var keys _dafny.Sequence = _dafny.EmptySeq
	_ = keys
	var _out0 _dafny.Sequence
	_ = _out0
	_out0 = Companion_Default___.SortedKeySetRuntime((m).Keys())
	keys = _out0
	return keys
}
func (_static *CompanionStruct_Default___) IntText(n _dafny.Int) _dafny.Sequence {
	if (n).Sign() == -1 {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("-"), m_ConfluxDecimalText.Companion_Default___.NatText((_dafny.Zero).Minus(n)))
	} else {
		return m_ConfluxDecimalText.Companion_Default___.NatText(n)
	}
}
func (_static *CompanionStruct_Default___) EscapeChar(c _dafny.CodePoint) _dafny.Sequence {
	if (c) == (_dafny.CodePoint('"')) {
		return _dafny.UnicodeSeqOfUtf8Bytes("\\\"")
	} else if (c) == (_dafny.CodePoint('\\')) {
		return _dafny.UnicodeSeqOfUtf8Bytes("\\\\")
	} else if (c) == (_dafny.CodePoint('\n')) {
		return _dafny.UnicodeSeqOfUtf8Bytes("\\n")
	} else if (c) == (_dafny.CodePoint('\r')) {
		return _dafny.UnicodeSeqOfUtf8Bytes("\\r")
	} else if (c) == (_dafny.CodePoint('\t')) {
		return _dafny.UnicodeSeqOfUtf8Bytes("\\t")
	} else if (c) == (_dafny.CodePoint('\U00000008')) {
		return _dafny.UnicodeSeqOfUtf8Bytes("\\b")
	} else if (c) == (_dafny.CodePoint('\U0000000C')) {
		return _dafny.UnicodeSeqOfUtf8Bytes("\\f")
	} else if (_dafny.IntOfInt32(rune(c))).Cmp(_dafny.IntOfInt64(32)) < 0 {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("\\u00"), _dafny.SeqOf(Companion_Default___.HexDigit((_dafny.IntOfInt32(rune(c))).DivBy(_dafny.IntOfInt64(16))), Companion_Default___.HexDigit((_dafny.IntOfInt32(rune(c))).Modulo(_dafny.IntOfInt64(16)))))
	} else {
		return _dafny.SeqOf(c)
	}
}
func (_static *CompanionStruct_Default___) Escape(s _dafny.Sequence) _dafny.Sequence {
	var _hresult _dafny.Sequence = _dafny.EmptySeq
	_ = _hresult
	if (_dafny.IntOfUint32((s).Cardinality())).Sign() == 0 {
		_hresult = _dafny.UnicodeSeqOfUtf8Bytes("")
		return _hresult
	}
	var _0_n _dafny.Int
	_ = _0_n
	_0_n = _dafny.IntOfUint32((s).Cardinality())
	var _1_arr _dafny.Array
	_ = _1_arr
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), (_dafny.IntOfInt64(6)).Times(_0_n))
	_ = _nw0
	_1_arr = _nw0
	var _2_off _dafny.Int
	_ = _2_off
	_2_off = _dafny.Zero
	var _3_i _dafny.Int
	_ = _3_i
	_3_i = _dafny.Zero
	for (_3_i).Cmp(_0_n) < 0 {
		var _4_e _dafny.Sequence
		_ = _4_e
		_4_e = Companion_Default___.EscapeChar((s).Select((_3_i).Uint32()).(_dafny.CodePoint))
		var _5_j _dafny.Int
		_ = _5_j
		_5_j = _dafny.Zero
		for (_5_j).Cmp(_dafny.IntOfUint32((_4_e).Cardinality())) < 0 {
			var _index0 _dafny.Int = (_2_off).Plus(_5_j)
			_ = _index0
			(_1_arr).ArraySet1CodePoint((_4_e).Select((_5_j).Uint32()).(_dafny.CodePoint), (_index0).Int())
			_5_j = (_5_j).Plus(_dafny.One)
		}
		_2_off = (_2_off).Plus(_dafny.IntOfUint32((_4_e).Cardinality()))
		_3_i = (_3_i).Plus(_dafny.One)
	}
	var _6_result _dafny.Sequence
	_ = _6_result
	_6_result = _dafny.SeqCreate((_2_off).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg142 _dafny.Int) interface{} {
			return coer117(arg142)
		}
	}((func(_7_arr _dafny.Array, _8_off _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
		return func(_9_k _dafny.Int) _dafny.CodePoint {
			return (_7_arr).ArrayGet1CodePoint((_9_k).Int())
		}
	})(_1_arr, _2_off)))
	_hresult = _6_result
	return _hresult
	return _hresult
}
func (_static *CompanionStruct_Default___) AllDigits(ds _dafny.Sequence) bool {
	return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((ds).Cardinality())), true, func(_forall_var_0 _dafny.Int) bool {
		var _0_i _dafny.Int
		_0_i = interface{}(_forall_var_0).(_dafny.Int)
		return !(((_0_i).Sign() != -1) && ((_0_i).Cmp(_dafny.IntOfUint32((ds).Cardinality())) < 0)) || (Companion_Default___.IsDigit((ds).Select((_0_i).Uint32()).(_dafny.CodePoint)))
	})
}
func (_static *CompanionStruct_Default___) FoldDigits(ds _dafny.Sequence, acc _dafny.Int) _dafny.Int {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((ds).Cardinality())).Sign() == 0 {
		return acc
	} else {
		var _in0 _dafny.Sequence = (ds).Drop(1)
		_ = _in0
		var _in1 _dafny.Int = ((acc).Times(_dafny.IntOfInt64(10))).Plus((_dafny.IntOfInt32(rune((ds).Select(0).(_dafny.CodePoint)))).Minus(_dafny.IntOfInt32(rune(_dafny.CodePoint('0')))))
		_ = _in1
		ds = _in0
		acc = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) DigitStep(acc _dafny.Int, d _dafny.CodePoint) _dafny.Int {
	return ((acc).Times(_dafny.IntOfInt64(10))).Plus((_dafny.IntOfInt32(rune(d))).Minus(_dafny.IntOfInt32(rune(_dafny.CodePoint('0')))))
}
func (_static *CompanionStruct_Default___) ValueStop(suffix _dafny.Sequence) bool {
	return ((((_dafny.IntOfUint32((suffix).Cardinality())).Sign() == 0) || (((suffix).Select(0).(_dafny.CodePoint)) == (_dafny.CodePoint(',')))) || (((suffix).Select(0).(_dafny.CodePoint)) == (_dafny.CodePoint(']')))) || (((suffix).Select(0).(_dafny.CodePoint)) == (_dafny.CodePoint('}')))
}
func (_static *CompanionStruct_Default___) SignLen(n _dafny.Int) _dafny.Int {
	if (n).Sign() == -1 {
		return _dafny.One
	} else {
		return _dafny.Zero
	}
}
func (_static *CompanionStruct_Default___) Magnitude(n _dafny.Int) _dafny.Int {
	if (n).Sign() == -1 {
		return (_dafny.Zero).Minus(n)
	} else {
		return n
	}
}
func (_static *CompanionStruct_Default___) StringifySeqRuntime(xs _dafny.Sequence, budget _dafny.Int) _dafny.Sequence {
	var s _dafny.Sequence = _dafny.EmptySeq
	_ = s
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		s = _dafny.UnicodeSeqOfUtf8Bytes("")
	} else if (_dafny.IntOfUint32((xs).Cardinality())).Cmp(_dafny.One) == 0 {
		var _out0 _dafny.Sequence
		_ = _out0
		_out0 = Companion_Default___.StringifyValue((xs).Select(0).(m_ConfluxCodec.Json), budget)
		s = _out0
	} else {
		var _0_split _dafny.Int
		_ = _0_split
		_0_split = (_dafny.IntOfUint32((xs).Cardinality())).DivBy(_dafny.IntOfInt64(2))
		var _1_left _dafny.Sequence
		_ = _1_left
		var _out1 _dafny.Sequence
		_ = _out1
		_out1 = Companion_Default___.StringifySeqRuntime((xs).Take((_0_split).Uint32()), budget)
		_1_left = _out1
		var _2_right _dafny.Sequence
		_ = _2_right
		var _out2 _dafny.Sequence
		_ = _out2
		_out2 = Companion_Default___.StringifySeqRuntime((xs).Drop((_0_split).Uint32()), budget)
		_2_right = _out2
		s = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1_left, _dafny.UnicodeSeqOfUtf8Bytes(",")), _2_right)
	}
	return s
}
func (_static *CompanionStruct_Default___) StringifyFieldsRuntime(fields _dafny.Map, keys _dafny.Sequence, budget _dafny.Int) _dafny.Sequence {
	var s _dafny.Sequence = _dafny.EmptySeq
	_ = s
	if (_dafny.IntOfUint32((keys).Cardinality())).Sign() == 0 {
		s = _dafny.UnicodeSeqOfUtf8Bytes("")
	} else {
		var _0_value _dafny.Sequence
		_ = _0_value
		var _out0 _dafny.Sequence
		_ = _out0
		_out0 = Companion_Default___.StringifyValue((fields).Get((keys).Select(0).(_dafny.Sequence)).(m_ConfluxCodec.Json), budget)
		_0_value = _out0
		var _1_head _dafny.Sequence
		_ = _1_head
		_1_head = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("\""), Companion_Default___.Escape((keys).Select(0).(_dafny.Sequence))), _dafny.UnicodeSeqOfUtf8Bytes("\":")), _0_value)
		if (_dafny.IntOfUint32((keys).Cardinality())).Cmp(_dafny.One) == 0 {
			s = _1_head
		} else {
			var _2_tail _dafny.Sequence
			_ = _2_tail
			var _out1 _dafny.Sequence
			_ = _out1
			_out1 = Companion_Default___.StringifyFieldsRuntime(fields, (keys).Drop(1), budget)
			_2_tail = _out1
			s = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1_head, _dafny.UnicodeSeqOfUtf8Bytes(",")), _2_tail)
		}
	}
	return s
}
func (_static *CompanionStruct_Default___) StringifyValue(j m_ConfluxCodec.Json, budget _dafny.Int) _dafny.Sequence {
	var s _dafny.Sequence = _dafny.EmptySeq
	_ = s
	if (budget).Sign() == 0 {
		s = _dafny.UnicodeSeqOfUtf8Bytes("null")
		return s
	}
	if (j).Is_JNull() {
		s = _dafny.UnicodeSeqOfUtf8Bytes("null")
		return s
	}
	if (j).Is_JBool() {
		if (j).Dtor_b() {
			s = _dafny.UnicodeSeqOfUtf8Bytes("true")
		} else {
			s = _dafny.UnicodeSeqOfUtf8Bytes("false")
		}
		return s
	}
	if (j).Is_JNum() {
		s = Companion_Default___.IntText((j).Dtor_n())
		return s
	}
	if (j).Is_JDec() {
		s = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.IntText((j).Dtor_mantissa()), _dafny.UnicodeSeqOfUtf8Bytes("e")), Companion_Default___.IntText((j).Dtor_exp()))
		return s
	}
	if (j).Is_JStr() {
		s = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("\""), Companion_Default___.Escape((j).Dtor_s())), _dafny.UnicodeSeqOfUtf8Bytes("\""))
		return s
	}
	if (j).Is_JArr() {
		var _0_body _dafny.Sequence
		_ = _0_body
		var _out0 _dafny.Sequence
		_ = _out0
		_out0 = Companion_Default___.StringifySeqRuntime((j).Dtor_elems(), (budget).Minus(_dafny.One))
		_0_body = _out0
		s = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("["), _0_body), _dafny.UnicodeSeqOfUtf8Bytes("]"))
		return s
	}
	var _1_keys _dafny.Sequence
	_ = _1_keys
	var _out1 _dafny.Sequence
	_ = _out1
	_out1 = Companion_Default___.SortedKeys((j).Dtor_fields())
	_1_keys = _out1
	var _2_body _dafny.Sequence
	_ = _2_body
	var _out2 _dafny.Sequence
	_ = _out2
	_out2 = Companion_Default___.StringifyFieldsRuntime((j).Dtor_fields(), _1_keys, (budget).Minus(_dafny.One))
	_2_body = _out2
	s = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("{"), _2_body), _dafny.UnicodeSeqOfUtf8Bytes("}"))
	return s
}
func (_static *CompanionStruct_Default___) Stringify(j m_ConfluxCodec.Json) _dafny.Sequence {
	var s _dafny.Sequence = _dafny.EmptySeq
	_ = s
	var _out0 _dafny.Sequence
	_ = _out0
	_out0 = Companion_Default___.StringifyValue(j, Companion_Default___.STRINGIFY__BUDGET())
	s = _out0
	return s
}
func (_static *CompanionStruct_Default___) PARSE__BUDGET() _dafny.Int {
	return _dafny.IntOfInt64(4000)
}
func (_static *CompanionStruct_Default___) STRINGIFY__BUDGET() _dafny.Int {
	return _dafny.IntOfInt64(4000)
}

// End of class Default__

// Definition of datatype StringCursor
type StringCursor struct {
	Data_StringCursor_
}

func (_this StringCursor) Get_() Data_StringCursor_ {
	return _this.Data_StringCursor_
}

type Data_StringCursor_ interface {
	isStringCursor()
}

type CompanionStruct_StringCursor_ struct {
}

var Companion_StringCursor_ = CompanionStruct_StringCursor_{}

type StringCursor_StringCursor struct {
	Value _dafny.Sequence
	Pos   _dafny.Int
	Ok    bool
}

func (StringCursor_StringCursor) isStringCursor() {}

func (CompanionStruct_StringCursor_) Create_StringCursor_(Value _dafny.Sequence, Pos _dafny.Int, Ok bool) StringCursor {
	return StringCursor{StringCursor_StringCursor{Value, Pos, Ok}}
}

func (_this StringCursor) Is_StringCursor() bool {
	_, ok := _this.Get_().(StringCursor_StringCursor)
	return ok
}

func (CompanionStruct_StringCursor_) Default() StringCursor {
	return Companion_StringCursor_.Create_StringCursor_(_dafny.EmptySeq, _dafny.Zero, false)
}

func (_this StringCursor) Dtor_value() _dafny.Sequence {
	return _this.Get_().(StringCursor_StringCursor).Value
}

func (_this StringCursor) Dtor_pos() _dafny.Int {
	return _this.Get_().(StringCursor_StringCursor).Pos
}

func (_this StringCursor) Dtor_ok() bool {
	return _this.Get_().(StringCursor_StringCursor).Ok
}

func (_this StringCursor) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case StringCursor_StringCursor:
		{
			return "ConfluxJsonText.StringCursor.StringCursor" + "(" + data.Value.VerbatimString(true) + ", " + _dafny.String(data.Pos) + ", " + _dafny.String(data.Ok) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this StringCursor) Equals(other StringCursor) bool {
	switch data1 := _this.Get_().(type) {
	case StringCursor_StringCursor:
		{
			data2, ok := other.Get_().(StringCursor_StringCursor)
			return ok && data1.Value.Equals(data2.Value) && data1.Pos.Cmp(data2.Pos) == 0 && data1.Ok == data2.Ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this StringCursor) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(StringCursor)
	return ok && _this.Equals(typed)
}

func Type_StringCursor_() _dafny.TypeDescriptor {
	return type_StringCursor_{}
}

type type_StringCursor_ struct {
}

func (_this type_StringCursor_) Default() interface{} {
	return Companion_StringCursor_.Default()
}

func (_this type_StringCursor_) String() string {
	return "ConfluxJsonText.StringCursor"
}
func (_this StringCursor) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = StringCursor{}

// End of datatype StringCursor

// Definition of datatype JsonCursor
type JsonCursor struct {
	Data_JsonCursor_
}

func (_this JsonCursor) Get_() Data_JsonCursor_ {
	return _this.Data_JsonCursor_
}

type Data_JsonCursor_ interface {
	isJsonCursor()
}

type CompanionStruct_JsonCursor_ struct {
}

var Companion_JsonCursor_ = CompanionStruct_JsonCursor_{}

type JsonCursor_JsonCursor struct {
	Value m_ConfluxCodec.Json
	Pos   _dafny.Int
	Ok    bool
}

func (JsonCursor_JsonCursor) isJsonCursor() {}

func (CompanionStruct_JsonCursor_) Create_JsonCursor_(Value m_ConfluxCodec.Json, Pos _dafny.Int, Ok bool) JsonCursor {
	return JsonCursor{JsonCursor_JsonCursor{Value, Pos, Ok}}
}

func (_this JsonCursor) Is_JsonCursor() bool {
	_, ok := _this.Get_().(JsonCursor_JsonCursor)
	return ok
}

func (CompanionStruct_JsonCursor_) Default() JsonCursor {
	return Companion_JsonCursor_.Create_JsonCursor_(m_ConfluxCodec.Companion_Json_.Default(), _dafny.Zero, false)
}

func (_this JsonCursor) Dtor_value() m_ConfluxCodec.Json {
	return _this.Get_().(JsonCursor_JsonCursor).Value
}

func (_this JsonCursor) Dtor_pos() _dafny.Int {
	return _this.Get_().(JsonCursor_JsonCursor).Pos
}

func (_this JsonCursor) Dtor_ok() bool {
	return _this.Get_().(JsonCursor_JsonCursor).Ok
}

func (_this JsonCursor) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case JsonCursor_JsonCursor:
		{
			return "ConfluxJsonText.JsonCursor.JsonCursor" + "(" + _dafny.String(data.Value) + ", " + _dafny.String(data.Pos) + ", " + _dafny.String(data.Ok) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this JsonCursor) Equals(other JsonCursor) bool {
	switch data1 := _this.Get_().(type) {
	case JsonCursor_JsonCursor:
		{
			data2, ok := other.Get_().(JsonCursor_JsonCursor)
			return ok && data1.Value.Equals(data2.Value) && data1.Pos.Cmp(data2.Pos) == 0 && data1.Ok == data2.Ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this JsonCursor) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(JsonCursor)
	return ok && _this.Equals(typed)
}

func Type_JsonCursor_() _dafny.TypeDescriptor {
	return type_JsonCursor_{}
}

type type_JsonCursor_ struct {
}

func (_this type_JsonCursor_) Default() interface{} {
	return Companion_JsonCursor_.Default()
}

func (_this type_JsonCursor_) String() string {
	return "ConfluxJsonText.JsonCursor"
}
func (_this JsonCursor) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = JsonCursor{}

// End of datatype JsonCursor

// Definition of datatype DigitCursor
type DigitCursor struct {
	Data_DigitCursor_
}

func (_this DigitCursor) Get_() Data_DigitCursor_ {
	return _this.Data_DigitCursor_
}

type Data_DigitCursor_ interface {
	isDigitCursor()
}

type CompanionStruct_DigitCursor_ struct {
}

var Companion_DigitCursor_ = CompanionStruct_DigitCursor_{}

type DigitCursor_DigitCursor struct {
	Value _dafny.Int
	Pos   _dafny.Int
	Count _dafny.Int
}

func (DigitCursor_DigitCursor) isDigitCursor() {}

func (CompanionStruct_DigitCursor_) Create_DigitCursor_(Value _dafny.Int, Pos _dafny.Int, Count _dafny.Int) DigitCursor {
	return DigitCursor{DigitCursor_DigitCursor{Value, Pos, Count}}
}

func (_this DigitCursor) Is_DigitCursor() bool {
	_, ok := _this.Get_().(DigitCursor_DigitCursor)
	return ok
}

func (CompanionStruct_DigitCursor_) Default() DigitCursor {
	return Companion_DigitCursor_.Create_DigitCursor_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this DigitCursor) Dtor_value() _dafny.Int {
	return _this.Get_().(DigitCursor_DigitCursor).Value
}

func (_this DigitCursor) Dtor_pos() _dafny.Int {
	return _this.Get_().(DigitCursor_DigitCursor).Pos
}

func (_this DigitCursor) Dtor_count() _dafny.Int {
	return _this.Get_().(DigitCursor_DigitCursor).Count
}

func (_this DigitCursor) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case DigitCursor_DigitCursor:
		{
			return "ConfluxJsonText.DigitCursor.DigitCursor" + "(" + _dafny.String(data.Value) + ", " + _dafny.String(data.Pos) + ", " + _dafny.String(data.Count) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this DigitCursor) Equals(other DigitCursor) bool {
	switch data1 := _this.Get_().(type) {
	case DigitCursor_DigitCursor:
		{
			data2, ok := other.Get_().(DigitCursor_DigitCursor)
			return ok && data1.Value.Cmp(data2.Value) == 0 && data1.Pos.Cmp(data2.Pos) == 0 && data1.Count.Cmp(data2.Count) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this DigitCursor) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(DigitCursor)
	return ok && _this.Equals(typed)
}

func Type_DigitCursor_() _dafny.TypeDescriptor {
	return type_DigitCursor_{}
}

type type_DigitCursor_ struct {
}

func (_this type_DigitCursor_) Default() interface{} {
	return Companion_DigitCursor_.Default()
}

func (_this type_DigitCursor_) String() string {
	return "ConfluxJsonText.DigitCursor"
}
func (_this DigitCursor) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = DigitCursor{}

// End of datatype DigitCursor

// Definition of datatype ParseResult
type ParseResult struct {
	Data_ParseResult_
}

func (_this ParseResult) Get_() Data_ParseResult_ {
	return _this.Data_ParseResult_
}

type Data_ParseResult_ interface {
	isParseResult()
}

type CompanionStruct_ParseResult_ struct {
}

var Companion_ParseResult_ = CompanionStruct_ParseResult_{}

type ParseResult_Invalid struct {
}

func (ParseResult_Invalid) isParseResult() {}

func (CompanionStruct_ParseResult_) Create_Invalid_() ParseResult {
	return ParseResult{ParseResult_Invalid{}}
}

func (_this ParseResult) Is_Invalid() bool {
	_, ok := _this.Get_().(ParseResult_Invalid)
	return ok
}

type ParseResult_Parsed struct {
	Value m_ConfluxCodec.Json
}

func (ParseResult_Parsed) isParseResult() {}

func (CompanionStruct_ParseResult_) Create_Parsed_(Value m_ConfluxCodec.Json) ParseResult {
	return ParseResult{ParseResult_Parsed{Value}}
}

func (_this ParseResult) Is_Parsed() bool {
	_, ok := _this.Get_().(ParseResult_Parsed)
	return ok
}

func (CompanionStruct_ParseResult_) Default() ParseResult {
	return Companion_ParseResult_.Create_Invalid_()
}

func (_this ParseResult) Dtor_value() m_ConfluxCodec.Json {
	return _this.Get_().(ParseResult_Parsed).Value
}

func (_this ParseResult) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ParseResult_Invalid:
		{
			return "ConfluxJsonText.ParseResult.Invalid"
		}
	case ParseResult_Parsed:
		{
			return "ConfluxJsonText.ParseResult.Parsed" + "(" + _dafny.String(data.Value) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ParseResult) Equals(other ParseResult) bool {
	switch data1 := _this.Get_().(type) {
	case ParseResult_Invalid:
		{
			_, ok := other.Get_().(ParseResult_Invalid)
			return ok
		}
	case ParseResult_Parsed:
		{
			data2, ok := other.Get_().(ParseResult_Parsed)
			return ok && data1.Value.Equals(data2.Value)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ParseResult) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ParseResult)
	return ok && _this.Equals(typed)
}

func Type_ParseResult_() _dafny.TypeDescriptor {
	return type_ParseResult_{}
}

type type_ParseResult_ struct {
}

func (_this type_ParseResult_) Default() interface{} {
	return Companion_ParseResult_.Default()
}

func (_this type_ParseResult_) String() string {
	return "ConfluxJsonText.ParseResult"
}
func (_this ParseResult) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ParseResult{}

// End of datatype ParseResult
