// Package ConfluxSequence
// Dafny module ConfluxSequence compiled into Go

package ConfluxSequence

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
	m_ConfluxJoin "github.com/joshmouch/ahp-go/ConfluxJoin"
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
	return "ConfluxSequence.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) StartsAt(haystack _dafny.Sequence, needle _dafny.Sequence, at _dafny.Int) bool {
	return (((at).Plus(_dafny.IntOfUint32((needle).Cardinality()))).Cmp(_dafny.IntOfUint32((haystack).Cardinality())) <= 0) && (_dafny.Companion_Sequence_.Equal((haystack).Subsequence((at).Uint32(), ((at).Plus(_dafny.IntOfUint32((needle).Cardinality()))).Uint32()), needle))
}
func (_static *CompanionStruct_Default___) FindSubsequenceFrom(haystack _dafny.Sequence, needle _dafny.Sequence, at _dafny.Int) _dafny.Int {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if Companion_Default___.StartsAt(haystack, needle, at) {
		return at
	} else if (at).Cmp(_dafny.IntOfUint32((haystack).Cardinality())) == 0 {
		return _dafny.IntOfInt64(-1)
	} else {
		var _in0 _dafny.Sequence = haystack
		_ = _in0
		var _in1 _dafny.Sequence = needle
		_ = _in1
		var _in2 _dafny.Int = (at).Plus(_dafny.One)
		_ = _in2
		haystack = _in0
		needle = _in1
		at = _in2
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Contains(haystack _dafny.Sequence, needle _dafny.Sequence) bool {
	return (Companion_Default___.FindSubsequenceFrom(haystack, needle, _dafny.Zero)).Sign() != -1
}
func (_static *CompanionStruct_Default___) EndsWith(haystack _dafny.Sequence, suffix _dafny.Sequence) bool {
	return ((_dafny.IntOfUint32((suffix).Cardinality())).Cmp(_dafny.IntOfUint32((haystack).Cardinality())) <= 0) && (_dafny.Companion_Sequence_.Equal((haystack).Drop(((_dafny.IntOfUint32((haystack).Cardinality())).Minus(_dafny.IntOfUint32((suffix).Cardinality()))).Uint32()), suffix))
}
func (_static *CompanionStruct_Default___) FindIndexWhere(p func(interface{}) bool, items _dafny.Sequence) _dafny.Int {
	if (_dafny.IntOfUint32((items).Cardinality())).Sign() == 0 {
		return _dafny.IntOfInt64(-1)
	} else if (p)((items).Select(0).(interface{})) {
		return _dafny.Zero
	} else {
		var _0_tail _dafny.Int = Companion_Default___.FindIndexWhere(p, (items).Drop(1))
		_ = _0_tail
		if (_0_tail).Sign() == -1 {
			return _dafny.IntOfInt64(-1)
		} else {
			return (_0_tail).Plus(_dafny.One)
		}
	}
}

// End of class Default__
