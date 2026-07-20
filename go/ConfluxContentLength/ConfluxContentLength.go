// Package ConfluxContentLength
// Dafny module ConfluxContentLength compiled into Go

package ConfluxContentLength

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-go/Changeset"
	m_Chat "github.com/joshmouch/ahp-go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-go/ClientMain"
	m_ConfluxBatchRoute "github.com/joshmouch/ahp-go/ConfluxBatchRoute"
	m_ConfluxBudgetConvergence "github.com/joshmouch/ahp-go/ConfluxBudgetConvergence"
	m_ConfluxChannel "github.com/joshmouch/ahp-go/ConfluxChannel"
	m_ConfluxChannelManifest "github.com/joshmouch/ahp-go/ConfluxChannelManifest"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxCommand "github.com/joshmouch/ahp-go/ConfluxCommand"
	m_ConfluxCommandIdentity "github.com/joshmouch/ahp-go/ConfluxCommandIdentity"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxConvergence "github.com/joshmouch/ahp-go/ConfluxConvergence"
	m_ConfluxDecimalText "github.com/joshmouch/ahp-go/ConfluxDecimalText"
	m_ConfluxDedupe "github.com/joshmouch/ahp-go/ConfluxDedupe"
	m_ConfluxDelimited "github.com/joshmouch/ahp-go/ConfluxDelimited"
	m_ConfluxDurableHistory "github.com/joshmouch/ahp-go/ConfluxDurableHistory"
	m_ConfluxExtern "github.com/joshmouch/ahp-go/ConfluxExtern"
	m_ConfluxFilterKeys "github.com/joshmouch/ahp-go/ConfluxFilterKeys"
	m_ConfluxFixpoint "github.com/joshmouch/ahp-go/ConfluxFixpoint"
	m_ConfluxGenericCodec "github.com/joshmouch/ahp-go/ConfluxGenericCodec"
	m_ConfluxHttpService "github.com/joshmouch/ahp-go/ConfluxHttpService"
	m_ConfluxJoin "github.com/joshmouch/ahp-go/ConfluxJoin"
	m_ConfluxJsonRpc "github.com/joshmouch/ahp-go/ConfluxJsonRpc"
	m_ConfluxJsonText "github.com/joshmouch/ahp-go/ConfluxJsonText"
	m_ConfluxKeyedOps "github.com/joshmouch/ahp-go/ConfluxKeyedOps"
	m_ConfluxMapValues "github.com/joshmouch/ahp-go/ConfluxMapValues"
	m_ConfluxMerge "github.com/joshmouch/ahp-go/ConfluxMerge"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-go/ConfluxOrderedLog"
	m_ConfluxProduct "github.com/joshmouch/ahp-go/ConfluxProduct"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxResolve "github.com/joshmouch/ahp-go/ConfluxResolve"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-go/ConfluxResourceCeilings"
	m_ConfluxResourceSupervisor "github.com/joshmouch/ahp-go/ConfluxResourceSupervisor"
	m_ConfluxRoute "github.com/joshmouch/ahp-go/ConfluxRoute"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-go/ConfluxSeqRoute"
	m_ConfluxSeqRouteMerge "github.com/joshmouch/ahp-go/ConfluxSeqRouteMerge"
	m_ConfluxSequence "github.com/joshmouch/ahp-go/ConfluxSequence"
	m_ConfluxServiceLifecycle "github.com/joshmouch/ahp-go/ConfluxServiceLifecycle"
	m_ConfluxStateMachine "github.com/joshmouch/ahp-go/ConfluxStateMachine"
	m_ConfluxSupervisedOperationResult "github.com/joshmouch/ahp-go/ConfluxSupervisedOperationResult"
	m_ConfluxTree "github.com/joshmouch/ahp-go/ConfluxTree"
	m_ConfluxWatermark "github.com/joshmouch/ahp-go/ConfluxWatermark"
	m_ResourceWatch "github.com/joshmouch/ahp-go/ResourceWatch"
	m_Session "github.com/joshmouch/ahp-go/Session"
	m__System "github.com/joshmouch/ahp-go/System_"
	m_Terminal "github.com/joshmouch/ahp-go/Terminal"
	_dafny "github.com/joshmouch/ahp-go/dafny"
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
var _ m_ConfluxJsonText.Dummy__
var _ m_ConfluxCommandIdentity.Dummy__
var _ m_ConfluxResourceSupervisor.Dummy__
var _ m_ConfluxServiceLifecycle.Dummy__
var _ m_ConfluxSequence.Dummy__
var _ m_ConfluxJsonRpc.Dummy__
var _ m_ConfluxHttpService.Dummy__

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
	return "ConfluxContentLength.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Token() _dafny.Sequence {
	return _dafny.SeqOf(uint8(67), uint8(111), uint8(110), uint8(116), uint8(101), uint8(110), uint8(116), uint8(45), uint8(76), uint8(101), uint8(110), uint8(103), uint8(116), uint8(104), uint8(58))
}
func (_static *CompanionStruct_Default___) Separator() _dafny.Sequence {
	return _dafny.SeqOf(uint8(13), uint8(10), uint8(13), uint8(10))
}
func (_static *CompanionStruct_Default___) SkipSpaces(header _dafny.Sequence, at _dafny.Int) _dafny.Int {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if ((at).Cmp(_dafny.IntOfUint32((header).Cardinality())) < 0) && ((((header).Select((at).Uint32()).(uint8)) == (uint8(32))) || (((header).Select((at).Uint32()).(uint8)) == (uint8(9)))) {
		var _in0 _dafny.Sequence = header
		_ = _in0
		var _in1 _dafny.Int = (at).Plus(_dafny.One)
		_ = _in1
		header = _in0
		at = _in1
		goto TAIL_CALL_START
	} else {
		return at
	}
}
func (_static *CompanionStruct_Default___) IsDigit(value uint8) bool {
	return ((uint8(48)) <= (value)) && ((value) <= (uint8(57)))
}
func (_static *CompanionStruct_Default___) ReadLength(header _dafny.Sequence, at _dafny.Int, value _dafny.Int, seen bool) m_ConfluxContract.Option {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if ((at).Cmp(_dafny.IntOfUint32((header).Cardinality())) < 0) && (Companion_Default___.IsDigit((header).Select((at).Uint32()).(uint8))) {
		var _in0 _dafny.Sequence = header
		_ = _in0
		var _in1 _dafny.Int = (at).Plus(_dafny.One)
		_ = _in1
		var _in2 _dafny.Int = ((value).Times(_dafny.IntOfInt64(10))).Plus((_dafny.IntOfUint8((header).Select((at).Uint32()).(uint8))).Minus(_dafny.IntOfInt64(48)))
		_ = _in2
		var _in3 bool = true
		_ = _in3
		header = _in0
		at = _in1
		value = _in2
		seen = _in3
		goto TAIL_CALL_START
	} else if seen {
		return m_ConfluxContract.Companion_Option_.Create_Some_(value)
	} else {
		return m_ConfluxContract.Companion_Option_.Create_None_()
	}
}
func (_static *CompanionStruct_Default___) Parse(header _dafny.Sequence) m_ConfluxContract.Option {
	var _0_token _dafny.Sequence = Companion_Default___.Token()
	_ = _0_token
	var _1_at _dafny.Int = m_ConfluxSequence.Companion_Default___.FindSubsequenceFrom(header, _0_token, _dafny.Zero)
	_ = _1_at
	if (_1_at).Sign() == -1 {
		return m_ConfluxContract.Companion_Option_.Create_None_()
	} else if ((_1_at).Plus(_dafny.IntOfUint32((_0_token).Cardinality()))).Cmp(_dafny.IntOfUint32((header).Cardinality())) > 0 {
		return m_ConfluxContract.Companion_Option_.Create_None_()
	} else {
		return Companion_Default___.ReadLength(header, Companion_Default___.SkipSpaces(header, (_1_at).Plus(_dafny.IntOfUint32((_0_token).Cardinality()))), _dafny.Zero, false)
	}
}
func (_static *CompanionStruct_Default___) HeaderText(length _dafny.Int) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("Content-Length: "), m_ConfluxDecimalText.Companion_Default___.NatText(length)), _dafny.UnicodeSeqOfUtf8Bytes("\r\n\r\n"))
}
func (_static *CompanionStruct_Default___) Frame(header _dafny.Sequence, body _dafny.Sequence) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(header, body)
}
func (_static *CompanionStruct_Default___) Extract(data _dafny.Sequence) (_dafny.Sequence, _dafny.Sequence) {
	var rest _dafny.Sequence = _dafny.EmptySeq
	_ = rest
	var bodies _dafny.Sequence = _dafny.EmptySeq
	_ = bodies
	var _0_separator _dafny.Sequence
	_ = _0_separator
	_0_separator = Companion_Default___.Separator()
	var _1_separatorAt _dafny.Int
	_ = _1_separatorAt
	_1_separatorAt = m_ConfluxSequence.Companion_Default___.FindSubsequenceFrom(data, _0_separator, _dafny.Zero)
	if (_1_separatorAt).Sign() == -1 {
		var _rhs0 _dafny.Sequence = data
		_ = _rhs0
		var _rhs1 _dafny.Sequence = _dafny.SeqOf()
		_ = _rhs1
		rest = _rhs0
		bodies = _rhs1
		return rest, bodies
	}
	var _2_bodyStart _dafny.Int
	_ = _2_bodyStart
	_2_bodyStart = (_1_separatorAt).Plus(_dafny.IntOfUint32((_0_separator).Cardinality()))
	var _source0 m_ConfluxContract.Option = Companion_Default___.Parse((data).Take((_1_separatorAt).Uint32()))
	_ = _source0
	{
		{
			if _source0.Is_None() {
				var _out0 _dafny.Sequence
				_ = _out0
				var _out1 _dafny.Sequence
				_ = _out1
				_out0, _out1 = Companion_Default___.Extract((data).Drop((_2_bodyStart).Uint32()))
				rest = _out0
				bodies = _out1
				goto Lmatch0
			}
		}
		{
			var _3_length _dafny.Int = _source0.Get_().(m_ConfluxContract.Option_Some).Value.(_dafny.Int)
			_ = _3_length
			if (_dafny.IntOfUint32((data).Cardinality())).Cmp((_2_bodyStart).Plus(_3_length)) < 0 {
				var _rhs2 _dafny.Sequence = data
				_ = _rhs2
				var _rhs3 _dafny.Sequence = _dafny.SeqOf()
				_ = _rhs3
				rest = _rhs2
				bodies = _rhs3
			} else {
				var _4_tailBodies _dafny.Sequence = _dafny.EmptySeq
				_ = _4_tailBodies
				var _out2 _dafny.Sequence
				_ = _out2
				var _out3 _dafny.Sequence
				_ = _out3
				_out2, _out3 = Companion_Default___.Extract((data).Drop(((_2_bodyStart).Plus(_3_length)).Uint32()))
				rest = _out2
				_4_tailBodies = _out3
				bodies = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((data).Subsequence((_2_bodyStart).Uint32(), ((_2_bodyStart).Plus(_3_length)).Uint32())), _4_tailBodies)
			}
		}
		goto Lmatch0
	}
Lmatch0:
	return rest, bodies
}

// End of class Default__
