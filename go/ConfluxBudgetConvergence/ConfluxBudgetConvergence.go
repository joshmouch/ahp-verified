// Package ConfluxBudgetConvergence
// Dafny module ConfluxBudgetConvergence compiled into Go

package ConfluxBudgetConvergence

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-verified/go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-verified/go/Changeset"
	m_Chat "github.com/joshmouch/ahp-verified/go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-verified/go/ClientMain"
	m_ConfluxBatchRoute "github.com/joshmouch/ahp-verified/go/ConfluxBatchRoute"
	m_ConfluxChannel "github.com/joshmouch/ahp-verified/go/ConfluxChannel"
	m_ConfluxChannelManifest "github.com/joshmouch/ahp-verified/go/ConfluxChannelManifest"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxCommand "github.com/joshmouch/ahp-verified/go/ConfluxCommand"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-verified/go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxConvergence "github.com/joshmouch/ahp-verified/go/ConfluxConvergence"
	m_ConfluxDedupe "github.com/joshmouch/ahp-verified/go/ConfluxDedupe"
	m_ConfluxDelimited "github.com/joshmouch/ahp-verified/go/ConfluxDelimited"
	m_ConfluxDurableHistory "github.com/joshmouch/ahp-verified/go/ConfluxDurableHistory"
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
	return "ConfluxBudgetConvergence.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ApplyObservation(budget m_ConfluxResourceCeilings.ResourceCeilings, phase m_ConfluxResourceCeilings.ExecutionPhase, state m_ConfluxSupervisedOperationResult.SupervisionState, observation m_ConfluxResourceCeilings.ResourceObservation) m_ConfluxSupervisedOperationResult.SupervisionState {
	var _source0 m_ConfluxSupervisedOperationResult.SupervisionState = state
	_ = _source0
	{
		if _source0.Is_SupervisionTripped() {
			var _0_record m_ConfluxSupervisedOperationResult.FirstBreachV1 = _source0.Get_().(m_ConfluxSupervisedOperationResult.SupervisionState_SupervisionTripped).Record
			_ = _0_record
			return state
		}
	}
	{
		var _1_samplesSeen _dafny.Int = _source0.Get_().(m_ConfluxSupervisedOperationResult.SupervisionState_SupervisionRunning).SamplesSeen
		_ = _1_samplesSeen
		var _2_verdict m_ConfluxResourceCeilings.ResourceVerdict = m_ConfluxResourceCeilings.Companion_Default___.DecideInFlight(budget, phase, observation)
		_ = _2_verdict
		if ((_2_verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementKill_()) {
			return m_ConfluxSupervisedOperationResult.Companion_SupervisionState_.Create_SupervisionTripped_(m_ConfluxSupervisedOperationResult.Companion_FirstBreachV1_.Create_FirstBreachV1_(_1_samplesSeen, _2_verdict, ((_2_verdict).Dtor_observed()).Is_None()))
		} else {
			return m_ConfluxSupervisedOperationResult.Companion_SupervisionState_.Create_SupervisionRunning_((_1_samplesSeen).Plus(_dafny.One))
		}
	}
}
func (_static *CompanionStruct_Default___) Drive(budget m_ConfluxResourceCeilings.ResourceCeilings, phase m_ConfluxResourceCeilings.ExecutionPhase, observations _dafny.Sequence) m_ConfluxSupervisedOperationResult.SupervisionState {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer116 func(m_ConfluxSupervisedOperationResult.SupervisionState, m_ConfluxResourceCeilings.ResourceObservation) m_ConfluxSupervisedOperationResult.SupervisionState) func(interface{}, interface{}) interface{} {
		return func(arg140 interface{}, arg141 interface{}) interface{} {
			return coer116(arg140.(m_ConfluxSupervisedOperationResult.SupervisionState), arg141.(m_ConfluxResourceCeilings.ResourceObservation))
		}
	}(Companion_Default___.SupervisionReducer(budget, phase)), m_ConfluxSupervisedOperationResult.Companion_SupervisionState_.Create_SupervisionRunning_(_dafny.Zero), observations).(m_ConfluxSupervisedOperationResult.SupervisionState)
}
func (_static *CompanionStruct_Default___) SupervisionReducer(budget m_ConfluxResourceCeilings.ResourceCeilings, phase m_ConfluxResourceCeilings.ExecutionPhase) func(m_ConfluxSupervisedOperationResult.SupervisionState, m_ConfluxResourceCeilings.ResourceObservation) m_ConfluxSupervisedOperationResult.SupervisionState {
	return (func(_0_budget m_ConfluxResourceCeilings.ResourceCeilings, _1_phase m_ConfluxResourceCeilings.ExecutionPhase) func(m_ConfluxSupervisedOperationResult.SupervisionState, m_ConfluxResourceCeilings.ResourceObservation) m_ConfluxSupervisedOperationResult.SupervisionState {
		return func(_2_state m_ConfluxSupervisedOperationResult.SupervisionState, _3_observation m_ConfluxResourceCeilings.ResourceObservation) m_ConfluxSupervisedOperationResult.SupervisionState {
			return Companion_Default___.ApplyObservation(_0_budget, _1_phase, _2_state, _3_observation)
		}
	})(budget, phase)
}

// End of class Default__
