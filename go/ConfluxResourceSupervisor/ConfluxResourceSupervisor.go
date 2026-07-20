// Package ConfluxResourceSupervisor
// Dafny module ConfluxResourceSupervisor compiled into Go

package ConfluxResourceSupervisor

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
	m_ConfluxCommandIdentity "github.com/joshmouch/ahp-verified/go/ConfluxCommandIdentity"
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
	m_ConfluxIoCapability "github.com/joshmouch/ahp-verified/go/ConfluxIoCapability"
	m_ConfluxJoin "github.com/joshmouch/ahp-verified/go/ConfluxJoin"
	m_ConfluxJsonText "github.com/joshmouch/ahp-verified/go/ConfluxJsonText"
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
var _ m_ConfluxJsonText.Dummy__
var _ m_ConfluxCommandIdentity.Dummy__

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
	return "ConfluxResourceSupervisor.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) CleanupDiagnostic(cleanup m_ConfluxSupervisedOperationResult.CleanupStatus) _dafny.Sequence {
	if (cleanup).Is_CleanupFailed() {
		return (cleanup).Dtor_reason()
	} else {
		return _dafny.UnicodeSeqOfUtf8Bytes("")
	}
}
func (_static *CompanionStruct_Default___) TerminationFailureDiagnostic(operationSucceeded bool, handleFound bool, cleanupComplete bool, diagnostic _dafny.Sequence) _dafny.Sequence {
	if !_dafny.Companion_Sequence_.Equal(diagnostic, _dafny.UnicodeSeqOfUtf8Bytes("")) {
		return diagnostic
	} else if !(operationSucceeded) {
		return _dafny.UnicodeSeqOfUtf8Bytes("checked process-tree termination failed")
	} else if (handleFound) && (!(cleanupComplete)) {
		return _dafny.UnicodeSeqOfUtf8Bytes("checked process-tree termination was not confirmed")
	} else {
		return _dafny.UnicodeSeqOfUtf8Bytes("process handle disappeared before checked termination")
	}
}
func (_static *CompanionStruct_Default___) TerminateProcessWithin(proc _dafny.Int, deadlineMs _dafny.Int) ProcessTerminationReceipt {
	var receipt ProcessTerminationReceipt = Companion_ProcessTerminationReceipt_.Default()
	_ = receipt
	var _0_operationSucceeded bool
	_ = _0_operationSucceeded
	var _1_handleFound bool
	_ = _1_handleFound
	var _2_treeKillIssued bool
	_ = _2_treeKillIssued
	var _3_cleanupComplete bool
	_ = _3_cleanupComplete
	var _4_cleanupElapsedMs _dafny.Int
	_ = _4_cleanupElapsedMs
	var _5_diagnostic _dafny.Sequence
	_ = _5_diagnostic
	var _6_terminalChunk _dafny.Sequence
	_ = _6_terminalChunk
	var _7_terminalAvailable bool
	_ = _7_terminalAvailable
	var _8_combinedOutputBytesDelta _dafny.Int
	_ = _8_combinedOutputBytesDelta
	var _9_exitObserved bool
	_ = _9_exitObserved
	var _10_exitCode _dafny.Int
	_ = _10_exitCode
	var _out0 bool
	_ = _out0
	var _out1 bool
	_ = _out1
	var _out2 bool
	_ = _out2
	var _out3 bool
	_ = _out3
	var _out4 _dafny.Int
	_ = _out4
	var _out5 _dafny.Sequence
	_ = _out5
	var _out6 _dafny.Sequence
	_ = _out6
	var _out7 bool
	_ = _out7
	var _out8 _dafny.Int
	_ = _out8
	var _out9 bool
	_ = _out9
	var _out10 _dafny.Int
	_ = _out10
	_out0, _out1, _out2, _out3, _out4, _out5, _out6, _out7, _out8, _out9, _out10 = m_ConfluxIoCapability.Companion_Default___.TerminateProcessTreeWithin(proc, deadlineMs)
	_0_operationSucceeded = _out0
	_1_handleFound = _out1
	_2_treeKillIssued = _out2
	_3_cleanupComplete = _out3
	_4_cleanupElapsedMs = _out4
	_5_diagnostic = _out5
	_6_terminalChunk = _out6
	_7_terminalAvailable = _out7
	_8_combinedOutputBytesDelta = _out8
	_9_exitObserved = _out9
	_10_exitCode = _out10
	var _11_cleanup m_ConfluxSupervisedOperationResult.CleanupStatus
	_ = _11_cleanup
	if (_0_operationSucceeded) && (!(_1_handleFound)) {
		_11_cleanup = m_ConfluxSupervisedOperationResult.Companion_CleanupStatus_.Create_CleanupNotRequired_()
	} else if (_0_operationSucceeded) && (_3_cleanupComplete) {
		_11_cleanup = m_ConfluxSupervisedOperationResult.Companion_CleanupStatus_.Create_CleanupComplete_(Companion_Default___.NonNegative(_4_cleanupElapsedMs), deadlineMs)
	} else {
		_11_cleanup = m_ConfluxSupervisedOperationResult.Companion_CleanupStatus_.Create_CleanupFailed_(Companion_Default___.NonNegative(_4_cleanupElapsedMs), deadlineMs, Companion_Default___.TerminationFailureDiagnostic(_0_operationSucceeded, _1_handleFound, _3_cleanupComplete, _5_diagnostic))
	}
	var _12_exit m_ConfluxSupervisedOperationResult.ExitStatus
	_ = _12_exit
	if _9_exitObserved {
		_12_exit = m_ConfluxSupervisedOperationResult.Companion_ExitStatus_.Create_ExitObserved_(_10_exitCode)
	} else {
		_12_exit = m_ConfluxSupervisedOperationResult.Companion_ExitStatus_.Create_ExitNotObserved_()
	}
	receipt = Companion_ProcessTerminationReceipt_.Create_ProcessTerminationReceipt_(_1_handleFound, _2_treeKillIssued, _11_cleanup, _6_terminalChunk, _7_terminalAvailable, Companion_Default___.NonNegative(_8_combinedOutputBytesDelta), _12_exit)
	return receipt
}
func (_static *CompanionStruct_Default___) NonNegative(value _dafny.Int) _dafny.Int {
	if (value).Sign() == -1 {
		return _dafny.Zero
	} else {
		return (value)
	}
}
func (_static *CompanionStruct_Default___) FlatByteConcat(left _dafny.Sequence, right _dafny.Sequence) _dafny.Sequence {
	return _dafny.SeqCreate(((_dafny.IntOfUint32((left).Cardinality())).Plus(_dafny.IntOfUint32((right).Cardinality()))).Uint32(), func(coer122 func(_dafny.Int) uint8) func(_dafny.Int) interface{} {
		return func(arg147 _dafny.Int) interface{} {
			return coer122(arg147)
		}
	}((func(_0_left _dafny.Sequence, _1_right _dafny.Sequence) func(_dafny.Int) uint8 {
		return func(_2_i _dafny.Int) uint8 {
			return (func() uint8 {
				if (_2_i).Cmp(_dafny.IntOfUint32((_0_left).Cardinality())) < 0 {
					return (_0_left).Select((_2_i).Uint32()).(uint8)
				}
				return (_1_right).Select(((_2_i).Minus(_dafny.IntOfUint32((_0_left).Cardinality()))).Uint32()).(uint8)
			})()
		}
	})(left, right)))
}
func (_static *CompanionStruct_Default___) BytesToMegabytes(bytes _dafny.Int) _dafny.Int {
	return m_ConfluxResourceCeilings.Companion_Default___.BytesToMegabytes(bytes)
}
func (_static *CompanionStruct_Default___) DimensionFromNative(value _dafny.Int) m_ConfluxResourceCeilings.ResourceDimension {
	if (value).Sign() == 0 {
		return m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_TimeDimension_()
	} else if (value).Cmp(_dafny.One) == 0 {
		return m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_MemoryDimension_()
	} else {
		return m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_OutputDimension_()
	}
}
func (_static *CompanionStruct_Default___) WrapperExitCode(childCode _dafny.Int, verdict m_ConfluxResourceCeilings.ResourceVerdict) _dafny.Int {
	if ((verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementOk_()) {
		return childCode
	} else if ((verdict).Dtor_dimension()).Is_None() {
		if (childCode).Sign() == 0 {
			return _dafny.One
		} else {
			return childCode
		}
	} else if (((verdict).Dtor_dimension()).Dtor_value().(m_ConfluxResourceCeilings.ResourceDimension)).Equals(m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_TimeDimension_()) {
		return _dafny.IntOfInt64(124)
	} else if (((verdict).Dtor_dimension()).Dtor_value().(m_ConfluxResourceCeilings.ResourceDimension)).Equals(m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_MemoryDimension_()) {
		return _dafny.IntOfInt64(137)
	} else {
		return _dafny.IntOfInt64(125)
	}
}
func (_static *CompanionStruct_Default___) ObservationStatus(observation m_ConfluxSupervisedOperationResult.SupervisionObservationV1) m_ConfluxResourceCeilings.MeasurementStatus {
	if (((((observation).Dtor_peakDescendantRssMb()).Dtor_status()).Is_MeasurementComplete()) && ((((observation).Dtor_combinedOutputBytes()).Dtor_status()).Is_MeasurementComplete())) && (((observation).Dtor_processTree()).Is_ProcessTreeObserved()) {
		return m_ConfluxResourceCeilings.Companion_MeasurementStatus_.Create_MeasurementComplete_()
	} else {
		return m_ConfluxResourceCeilings.Companion_MeasurementStatus_.Create_MeasurementPartial_(_dafny.UnicodeSeqOfUtf8Bytes("process observation contains incomplete measurements"))
	}
}
func (_static *CompanionStruct_Default___) StartObservation() _dafny.Int {
	var startedMs _dafny.Int = _dafny.Zero
	_ = startedMs
	var _out0 _dafny.Int
	_ = _out0
	_out0 = m_ConfluxIoCapability.Companion_Default___.MonotonicTimeMs()
	startedMs = _out0
	return startedMs
}
func (_static *CompanionStruct_Default___) ObserveCurrentProcess(startedMs _dafny.Int, outputBytes m_ConfluxContract.Option, phase m_ConfluxResourceCeilings.ExecutionPhase, budget m_ConfluxResourceCeilings.ResourceCeilings) m_ConfluxResourceCeilings.ResourceObservation {
	var observation m_ConfluxResourceCeilings.ResourceObservation = m_ConfluxResourceCeilings.Companion_ResourceObservation_.Default()
	_ = observation
	var _0_now _dafny.Int
	_ = _0_now
	var _out0 _dafny.Int
	_ = _out0
	_out0 = m_ConfluxIoCapability.Companion_Default___.MonotonicTimeMs()
	_0_now = _out0
	var _1_memoryPresent bool
	_ = _1_memoryPresent
	var _2_memoryMb _dafny.Int
	_ = _2_memoryMb
	var _out1 bool
	_ = _out1
	var _out2 _dafny.Int
	_ = _out2
	_out1, _out2 = m_ConfluxIoCapability.Companion_Default___.CurrentProcessRssMb()
	_1_memoryPresent = _out1
	_2_memoryMb = _out2
	var _3_memory m_ConfluxResourceCeilings.ResourceMeasurement
	_ = _3_memory
	if _1_memoryPresent {
		_3_memory = m_ConfluxResourceCeilings.Companion_Default___.CompleteMeasurement(Companion_Default___.NonNegative(_2_memoryMb))
	} else if m_ConfluxResourceCeilings.Companion_Default___.MemoryMeasurementRequired(budget, phase) {
		_3_memory = m_ConfluxResourceCeilings.Companion_Default___.UnavailableMeasurement(_dafny.UnicodeSeqOfUtf8Bytes("current-process RSS unavailable"))
	} else {
		_3_memory = m_ConfluxResourceCeilings.Companion_Default___.NotRequestedMeasurement()
	}
	var _4_output m_ConfluxResourceCeilings.ResourceMeasurement
	_ = _4_output
	if (outputBytes).Is_Some() {
		_4_output = m_ConfluxResourceCeilings.Companion_Default___.CompleteMeasurement((outputBytes).Dtor_value().(_dafny.Int))
	} else if m_ConfluxResourceCeilings.Companion_Default___.OutputMeasurementRequired(budget) {
		_4_output = m_ConfluxResourceCeilings.Companion_Default___.UnavailableMeasurement(_dafny.UnicodeSeqOfUtf8Bytes("combined-output measurement unavailable"))
	} else {
		_4_output = m_ConfluxResourceCeilings.Companion_Default___.NotRequestedMeasurement()
	}
	observation = m_ConfluxResourceCeilings.Companion_ResourceObservation_.Create_ResourceObservation_(Companion_Default___.NonNegative((_0_now).Minus(startedMs)), _3_memory, _4_output)
	return observation
}
func (_static *CompanionStruct_Default___) ObserveCurrentProcessPostHoc(startedMs _dafny.Int, outputBytes m_ConfluxContract.Option, phase m_ConfluxResourceCeilings.ExecutionPhase, budget m_ConfluxResourceCeilings.ResourceCeilings) (m_ConfluxResourceCeilings.ResourceObservation, m_ConfluxResourceCeilings.ResourceVerdict) {
	var observation m_ConfluxResourceCeilings.ResourceObservation = m_ConfluxResourceCeilings.Companion_ResourceObservation_.Default()
	_ = observation
	var verdict m_ConfluxResourceCeilings.ResourceVerdict = m_ConfluxResourceCeilings.Companion_ResourceVerdict_.Default()
	_ = verdict
	var _out0 m_ConfluxResourceCeilings.ResourceObservation
	_ = _out0
	_out0 = Companion_Default___.ObserveCurrentProcess(startedMs, outputBytes, phase, budget)
	observation = _out0
	verdict = m_ConfluxResourceCeilings.Companion_Default___.DecidePostHoc(budget, phase, observation)
	return observation, verdict
}
func (_static *CompanionStruct_Default___) ObserveCurrentProcessPostHocV1(policyId _dafny.Sequence, startedMs _dafny.Int, outputBytes m_ConfluxContract.Option, phase m_ConfluxResourceCeilings.ExecutionPhase, budget m_ConfluxResourceCeilings.ResourceCeilings) m_ConfluxSupervisedOperationResult.SupervisedOperationResultV1 {
	var result m_ConfluxSupervisedOperationResult.SupervisedOperationResultV1 = m_ConfluxSupervisedOperationResult.Companion_SupervisedOperationResultV1_.Default()
	_ = result
	var _0_resolvedPolicyId _dafny.Sequence
	_ = _0_resolvedPolicyId
	if (_dafny.IntOfUint32((policyId).Cardinality())).Sign() == 1 {
		_0_resolvedPolicyId = policyId
	} else {
		_0_resolvedPolicyId = _dafny.UnicodeSeqOfUtf8Bytes("policy:conflux/default@1")
	}
	var _1_observation m_ConfluxResourceCeilings.ResourceObservation
	_ = _1_observation
	var _2_verdict m_ConfluxResourceCeilings.ResourceVerdict
	_ = _2_verdict
	var _out0 m_ConfluxResourceCeilings.ResourceObservation
	_ = _out0
	var _out1 m_ConfluxResourceCeilings.ResourceVerdict
	_ = _out1
	_out0, _out1 = Companion_Default___.ObserveCurrentProcessPostHoc(startedMs, outputBytes, phase, budget)
	_1_observation = _out0
	_2_verdict = _out1
	var _3_tree m_ConfluxSupervisedOperationResult.ProcessTreeStatus
	_ = _3_tree
	if (((_1_observation).Dtor_peakRssMb()).Dtor_status()).Is_MeasurementComplete() {
		_3_tree = m_ConfluxSupervisedOperationResult.Companion_ProcessTreeStatus_.Create_ProcessTreeObserved_()
	} else {
		_3_tree = m_ConfluxSupervisedOperationResult.Companion_ProcessTreeStatus_.Create_ProcessTreeUnavailable_(_dafny.UnicodeSeqOfUtf8Bytes("current-process RSS incomplete"))
	}
	var _4_supervisionObservation m_ConfluxSupervisedOperationResult.SupervisionObservationV1
	_ = _4_supervisionObservation
	_4_supervisionObservation = m_ConfluxSupervisedOperationResult.Companion_SupervisionObservationV1_.Create_SupervisionObservationV1_(m_ConfluxSupervisedOperationResult.Companion_TimingCoverage_.Create_TimingCoverage_(_dafny.Zero, (_1_observation).Dtor_wallMs(), _dafny.Zero, (_1_observation).Dtor_wallMs()), (_1_observation).Dtor_peakRssMb(), (_1_observation).Dtor_combinedOutputBytes(), _3_tree)
	var _5_state m_ConfluxSupervisedOperationResult.SupervisionState
	_ = _5_state
	if ((_2_verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementOk_()) {
		_5_state = m_ConfluxSupervisedOperationResult.Companion_SupervisionState_.Create_SupervisionRunning_(_dafny.One)
	} else {
		_5_state = m_ConfluxSupervisedOperationResult.Companion_SupervisionState_.Create_SupervisionTripped_(m_ConfluxSupervisedOperationResult.Companion_FirstBreachV1_.Create_FirstBreachV1_(_dafny.Zero, _2_verdict, ((_2_verdict).Dtor_observed()).Is_None()))
	}
	var _6_wrapperExit _dafny.Int
	_ = _6_wrapperExit
	_6_wrapperExit = Companion_Default___.WrapperExitCode(_dafny.Zero, _2_verdict)
	result = m_ConfluxSupervisedOperationResult.Companion_SupervisedOperationResultV1_.Create_SupervisedOperationResultV1_(m_ConfluxSupervisedOperationResult.Companion_Default___.SupervisedOperationResultSchemaVersion(), _0_resolvedPolicyId, phase, m_ConfluxResourceCeilings.Companion_Default___.EffectivePhase(phase), m_ConfluxSupervisedOperationResult.Companion_Default___.ExpandBudget(budget, phase), Companion_Default___.ObservationStatus(_4_supervisionObservation), _4_supervisionObservation, _5_state, m_ConfluxSupervisedOperationResult.Companion_CleanupStatus_.Create_CleanupNotRequired_(), m_ConfluxSupervisedOperationResult.Companion_ExitStatus_.Create_ExitNotObserved_(), (func() m_ConfluxSupervisedOperationResult.EnforcementStatus {
		if ((_2_verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementFail_()) {
			return m_ConfluxSupervisedOperationResult.Companion_EnforcementStatus_.Create_PostHocFailure_()
		}
		return m_ConfluxSupervisedOperationResult.Companion_EnforcementStatus_.Create_NoEnforcement_()
	})(), m_ConfluxSupervisedOperationResult.Companion_TerminationScope_.Create_NoTermination_(), _6_wrapperExit, (_2_verdict).Dtor_reason(), _2_verdict)
	return result
}
func (_static *CompanionStruct_Default___) RunProcess(cwd _dafny.Sequence, command _dafny.Sequence, arguments _dafny.Sequence, phase m_ConfluxResourceCeilings.ExecutionPhase, budget m_ConfluxResourceCeilings.ResourceCeilings) SupervisedProcessResult {
	var result SupervisedProcessResult = Companion_SupervisedProcessResult_.Default()
	_ = result
	var _out0 SupervisedProcessResult
	_ = _out0
	_out0 = Companion_Default___.RunProcessV1(_dafny.UnicodeSeqOfUtf8Bytes("policy:conflux/default@1"), cwd, command, arguments, phase, budget)
	result = _out0
	return result
}
func (_static *CompanionStruct_Default___) RunProcessV1(policyId _dafny.Sequence, cwd _dafny.Sequence, command _dafny.Sequence, arguments _dafny.Sequence, phase m_ConfluxResourceCeilings.ExecutionPhase, budget m_ConfluxResourceCeilings.ResourceCeilings) SupervisedProcessResult {
	var result SupervisedProcessResult = Companion_SupervisedProcessResult_.Default()
	_ = result
	var _0_resolvedPolicyId _dafny.Sequence
	_ = _0_resolvedPolicyId
	if (_dafny.IntOfUint32((policyId).Cardinality())).Sign() == 1 {
		_0_resolvedPolicyId = policyId
	} else {
		_0_resolvedPolicyId = _dafny.UnicodeSeqOfUtf8Bytes("policy:conflux/default@1")
	}
	var _1_timeLimit _dafny.Int
	_ = _1_timeLimit
	_1_timeLimit = m_ConfluxResourceCeilings.Companion_Default___.TimeCeilingForPhase((budget).Dtor_time(), phase)
	var _2_memoryLimit _dafny.Int
	_ = _2_memoryLimit
	_2_memoryLimit = m_ConfluxResourceCeilings.Companion_Default___.MemoryCeilingForPhase((budget).Dtor_memory(), phase)
	var _3_code _dafny.Int
	_ = _3_code
	var _4_output _dafny.Sequence
	_ = _4_output
	var _5_error _dafny.Sequence
	_ = _5_error
	var _6_timedOut bool
	_ = _6_timedOut
	var _7_memoryExceeded bool
	_ = _7_memoryExceeded
	var _8_outputExceeded bool
	_ = _8_outputExceeded
	var _9_wallMs _dafny.Int
	_ = _9_wallMs
	var _10_memoryPresent bool
	_ = _10_memoryPresent
	var _11_peakMemoryMb _dafny.Int
	_ = _11_peakMemoryMb
	var _12_combinedOutputBytes _dafny.Int
	_ = _12_combinedOutputBytes
	var _13_memoryMeasurementFailed bool
	_ = _13_memoryMeasurementFailed
	var _14_outputMeasurementFailed bool
	_ = _14_outputMeasurementFailed
	var _15_treeKillIssued bool
	_ = _15_treeKillIssued
	var _16_cleanupComplete bool
	_ = _16_cleanupComplete
	var _17_cleanupElapsedMs _dafny.Int
	_ = _17_cleanupElapsedMs
	var _18_cleanupDeadlineMs _dafny.Int
	_ = _18_cleanupDeadlineMs
	var _19_cleanupDiagnostic _dafny.Sequence
	_ = _19_cleanupDiagnostic
	var _20_firstTripPresent bool
	_ = _20_firstTripPresent
	var _21_firstTripIndex _dafny.Int
	_ = _21_firstTripIndex
	var _22_firstTripDimension _dafny.Int
	_ = _22_firstTripDimension
	var _23_firstTripObservedPresent bool
	_ = _23_firstTripObservedPresent
	var _24_firstTripObserved _dafny.Int
	_ = _24_firstTripObserved
	var _25_firstTripMeasurementFailure bool
	_ = _25_firstTripMeasurementFailure
	var _26_firstTripInFlight bool
	_ = _26_firstTripInFlight
	var _27_samplesSeen _dafny.Int
	_ = _27_samplesSeen
	var _out0 _dafny.Int
	_ = _out0
	var _out1 _dafny.Sequence
	_ = _out1
	var _out2 _dafny.Sequence
	_ = _out2
	var _out3 bool
	_ = _out3
	var _out4 bool
	_ = _out4
	var _out5 bool
	_ = _out5
	var _out6 _dafny.Int
	_ = _out6
	var _out7 bool
	_ = _out7
	var _out8 _dafny.Int
	_ = _out8
	var _out9 _dafny.Int
	_ = _out9
	var _out10 bool
	_ = _out10
	var _out11 bool
	_ = _out11
	var _out12 bool
	_ = _out12
	var _out13 bool
	_ = _out13
	var _out14 _dafny.Int
	_ = _out14
	var _out15 _dafny.Int
	_ = _out15
	var _out16 _dafny.Sequence
	_ = _out16
	var _out17 bool
	_ = _out17
	var _out18 _dafny.Int
	_ = _out18
	var _out19 _dafny.Int
	_ = _out19
	var _out20 bool
	_ = _out20
	var _out21 _dafny.Int
	_ = _out21
	var _out22 bool
	_ = _out22
	var _out23 bool
	_ = _out23
	var _out24 _dafny.Int
	_ = _out24
	_out0, _out1, _out2, _out3, _out4, _out5, _out6, _out7, _out8, _out9, _out10, _out11, _out12, _out13, _out14, _out15, _out16, _out17, _out18, _out19, _out20, _out21, _out22, _out23, _out24 = m_ConfluxIoCapability.Companion_Default___.RunProcessBounded(cwd, command, arguments, _1_timeLimit, _2_memoryLimit, (budget).Dtor_maxOutputMb())
	_3_code = _out0
	_4_output = _out1
	_5_error = _out2
	_6_timedOut = _out3
	_7_memoryExceeded = _out4
	_8_outputExceeded = _out5
	_9_wallMs = _out6
	_10_memoryPresent = _out7
	_11_peakMemoryMb = _out8
	_12_combinedOutputBytes = _out9
	_13_memoryMeasurementFailed = _out10
	_14_outputMeasurementFailed = _out11
	_15_treeKillIssued = _out12
	_16_cleanupComplete = _out13
	_17_cleanupElapsedMs = _out14
	_18_cleanupDeadlineMs = _out15
	_19_cleanupDiagnostic = _out16
	_20_firstTripPresent = _out17
	_21_firstTripIndex = _out18
	_22_firstTripDimension = _out19
	_23_firstTripObservedPresent = _out20
	_24_firstTripObserved = _out21
	_25_firstTripMeasurementFailure = _out22
	_26_firstTripInFlight = _out23
	_27_samplesSeen = _out24
	var _28_memoryMeasurement m_ConfluxResourceCeilings.ResourceMeasurement
	_ = _28_memoryMeasurement
	if _10_memoryPresent {
		if _13_memoryMeasurementFailed {
			_28_memoryMeasurement = m_ConfluxResourceCeilings.Companion_ResourceMeasurement_.Create_ResourceMeasurement_(m_ConfluxContract.Companion_Option_.Create_Some_(Companion_Default___.NonNegative(_11_peakMemoryMb)), m_ConfluxResourceCeilings.Companion_MeasurementStatus_.Create_MeasurementPartial_(_dafny.UnicodeSeqOfUtf8Bytes("one or more process-tree RSS samples failed")))
		} else {
			_28_memoryMeasurement = m_ConfluxResourceCeilings.Companion_Default___.CompleteMeasurement(Companion_Default___.NonNegative(_11_peakMemoryMb))
		}
	} else if (_2_memoryLimit).Sign() == 1 {
		_28_memoryMeasurement = m_ConfluxResourceCeilings.Companion_Default___.UnavailableMeasurement(_dafny.UnicodeSeqOfUtf8Bytes("complete process-tree RSS unavailable"))
	} else {
		_28_memoryMeasurement = m_ConfluxResourceCeilings.Companion_Default___.NotRequestedMeasurement()
	}
	var _29_outputMeasurement m_ConfluxResourceCeilings.ResourceMeasurement
	_ = _29_outputMeasurement
	if _14_outputMeasurementFailed {
		if (_12_combinedOutputBytes).Sign() != -1 {
			_29_outputMeasurement = m_ConfluxResourceCeilings.Companion_ResourceMeasurement_.Create_ResourceMeasurement_(m_ConfluxContract.Companion_Option_.Create_Some_(Companion_Default___.NonNegative(_12_combinedOutputBytes)), m_ConfluxResourceCeilings.Companion_MeasurementStatus_.Create_MeasurementPartial_(_dafny.UnicodeSeqOfUtf8Bytes("combined-output stream observation incomplete")))
		} else {
			_29_outputMeasurement = m_ConfluxResourceCeilings.Companion_Default___.UnavailableMeasurement(_dafny.UnicodeSeqOfUtf8Bytes("combined-output stream observation unavailable"))
		}
	} else {
		_29_outputMeasurement = m_ConfluxResourceCeilings.Companion_Default___.CompleteMeasurement(Companion_Default___.NonNegative(_12_combinedOutputBytes))
	}
	var _30_observation m_ConfluxResourceCeilings.ResourceObservation
	_ = _30_observation
	_30_observation = m_ConfluxResourceCeilings.Companion_ResourceObservation_.Create_ResourceObservation_(Companion_Default___.NonNegative(_9_wallMs), _28_memoryMeasurement, _29_outputMeasurement)
	var _31_nativeDimension m_ConfluxResourceCeilings.ResourceDimension
	_ = _31_nativeDimension
	_31_nativeDimension = Companion_Default___.DimensionFromNative(_22_firstTripDimension)
	var _32_nativeObservedCandidate m_ConfluxContract.Option
	_ = _32_nativeObservedCandidate
	if _23_firstTripObservedPresent {
		_32_nativeObservedCandidate = m_ConfluxContract.Companion_Option_.Create_Some_(Companion_Default___.NonNegative(_24_firstTripObserved))
	} else {
		_32_nativeObservedCandidate = m_ConfluxResourceCeilings.Companion_Default___.ObservedForDimension(_30_observation, _31_nativeDimension)
	}
	var _33_nativeMeasurementFailure bool
	_ = _33_nativeMeasurementFailure
	_33_nativeMeasurementFailure = (!(_31_nativeDimension).Equals(m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_TimeDimension_())) && ((_25_firstTripMeasurementFailure) || ((_32_nativeObservedCandidate).Is_None()))
	var _34_nativeObserved m_ConfluxContract.Option
	_ = _34_nativeObserved
	if _33_nativeMeasurementFailure {
		_34_nativeObserved = m_ConfluxContract.Companion_Option_.Create_None_()
	} else {
		_34_nativeObserved = _32_nativeObservedCandidate
	}
	var _35_nativeOutcome m_ConfluxResourceCeilings.EnforcementOutcome
	_ = _35_nativeOutcome
	if _26_firstTripInFlight {
		_35_nativeOutcome = m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementKill_()
	} else {
		_35_nativeOutcome = m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementFail_()
	}
	var _36_nativeVerdict m_ConfluxResourceCeilings.ResourceVerdict
	_ = _36_nativeVerdict
	_36_nativeVerdict = m_ConfluxResourceCeilings.Companion_Default___.RecordedViolation(_35_nativeOutcome, budget, phase, _31_nativeDimension, _34_nativeObserved, _33_nativeMeasurementFailure)
	var _37_finalVerdict m_ConfluxResourceCeilings.ResourceVerdict
	_ = _37_finalVerdict
	_37_finalVerdict = m_ConfluxResourceCeilings.Companion_Default___.DecidePostHoc(budget, phase, _30_observation)
	var _38_verdict m_ConfluxResourceCeilings.ResourceVerdict
	_ = _38_verdict
	if _20_firstTripPresent {
		_38_verdict = _36_nativeVerdict
	} else {
		_38_verdict = _37_finalVerdict
	}
	var _39_state m_ConfluxSupervisedOperationResult.SupervisionState
	_ = _39_state
	if _20_firstTripPresent {
		_39_state = m_ConfluxSupervisedOperationResult.Companion_SupervisionState_.Create_SupervisionTripped_(m_ConfluxSupervisedOperationResult.Companion_FirstBreachV1_.Create_FirstBreachV1_(Companion_Default___.NonNegative(_21_firstTripIndex), _38_verdict, _33_nativeMeasurementFailure))
	} else if !((_38_verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementOk_()) {
		_39_state = m_ConfluxSupervisedOperationResult.Companion_SupervisionState_.Create_SupervisionTripped_(m_ConfluxSupervisedOperationResult.Companion_FirstBreachV1_.Create_FirstBreachV1_(Companion_Default___.NonNegative(_27_samplesSeen), _38_verdict, ((_38_verdict).Dtor_observed()).Is_None()))
	} else {
		_39_state = m_ConfluxSupervisedOperationResult.Companion_SupervisionState_.Create_SupervisionRunning_(Companion_Default___.NonNegative(_27_samplesSeen))
	}
	var _40_processTree m_ConfluxSupervisedOperationResult.ProcessTreeStatus
	_ = _40_processTree
	if ((_28_memoryMeasurement).Dtor_status()).Is_MeasurementComplete() {
		_40_processTree = m_ConfluxSupervisedOperationResult.Companion_ProcessTreeStatus_.Create_ProcessTreeObserved_()
	} else {
		_40_processTree = m_ConfluxSupervisedOperationResult.Companion_ProcessTreeStatus_.Create_ProcessTreeUnavailable_(_dafny.UnicodeSeqOfUtf8Bytes("complete process-tree RSS observation incomplete"))
	}
	var _41_supervisionObservation m_ConfluxSupervisedOperationResult.SupervisionObservationV1
	_ = _41_supervisionObservation
	_41_supervisionObservation = m_ConfluxSupervisedOperationResult.Companion_SupervisionObservationV1_.Create_SupervisionObservationV1_(m_ConfluxSupervisedOperationResult.Companion_TimingCoverage_.Create_TimingCoverage_(_dafny.Zero, (_30_observation).Dtor_wallMs(), _dafny.Zero, (_30_observation).Dtor_wallMs()), _28_memoryMeasurement, _29_outputMeasurement, _40_processTree)
	var _42_cleanup m_ConfluxSupervisedOperationResult.CleanupStatus
	_ = _42_cleanup
	if ((_38_verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementKill_()) {
		if _16_cleanupComplete {
			_42_cleanup = m_ConfluxSupervisedOperationResult.Companion_CleanupStatus_.Create_CleanupComplete_(Companion_Default___.NonNegative(_17_cleanupElapsedMs), Companion_Default___.NonNegative(_18_cleanupDeadlineMs))
		} else {
			_42_cleanup = m_ConfluxSupervisedOperationResult.Companion_CleanupStatus_.Create_CleanupFailed_(Companion_Default___.NonNegative(_17_cleanupElapsedMs), Companion_Default___.NonNegative(_18_cleanupDeadlineMs), (func() _dafny.Sequence {
				if (_dafny.IntOfUint32((_19_cleanupDiagnostic).Cardinality())).Sign() == 1 {
					return _19_cleanupDiagnostic
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("process-tree cleanup exceeded monotonic deadline")
			})())
		}
	} else {
		_42_cleanup = m_ConfluxSupervisedOperationResult.Companion_CleanupStatus_.Create_CleanupNotRequired_()
	}
	var _43_wrapperExit _dafny.Int
	_ = _43_wrapperExit
	_43_wrapperExit = Companion_Default___.WrapperExitCode(_3_code, _38_verdict)
	var _44_supervised m_ConfluxSupervisedOperationResult.SupervisedOperationResultV1
	_ = _44_supervised
	_44_supervised = m_ConfluxSupervisedOperationResult.Companion_SupervisedOperationResultV1_.Create_SupervisedOperationResultV1_(m_ConfluxSupervisedOperationResult.Companion_Default___.SupervisedOperationResultSchemaVersion(), _0_resolvedPolicyId, phase, m_ConfluxResourceCeilings.Companion_Default___.EffectivePhase(phase), m_ConfluxSupervisedOperationResult.Companion_Default___.ExpandBudget(budget, phase), Companion_Default___.ObservationStatus(_41_supervisionObservation), _41_supervisionObservation, _39_state, _42_cleanup, m_ConfluxSupervisedOperationResult.Companion_ExitStatus_.Create_ExitObserved_(_3_code), (func() m_ConfluxSupervisedOperationResult.EnforcementStatus {
		if ((_38_verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementKill_()) {
			return m_ConfluxSupervisedOperationResult.Companion_EnforcementStatus_.Create_InFlightProcessTreeTermination_()
		}
		return (func() m_ConfluxSupervisedOperationResult.EnforcementStatus {
			if ((_38_verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementFail_()) {
				return m_ConfluxSupervisedOperationResult.Companion_EnforcementStatus_.Create_PostHocFailure_()
			}
			return m_ConfluxSupervisedOperationResult.Companion_EnforcementStatus_.Create_NoEnforcement_()
		})()
	})(), (func() m_ConfluxSupervisedOperationResult.TerminationScope {
		if ((_38_verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementKill_()) {
			return m_ConfluxSupervisedOperationResult.Companion_TerminationScope_.Create_ProcessTreeTermination_()
		}
		return m_ConfluxSupervisedOperationResult.Companion_TerminationScope_.Create_NoTermination_()
	})(), _43_wrapperExit, (_38_verdict).Dtor_reason(), _38_verdict)
	result = Companion_SupervisedProcessResult_.Create_SupervisedProcessResult_(_3_code, _4_output, _5_error, _30_observation, _38_verdict, Companion_Default___.NonNegative(_12_combinedOutputBytes), _15_treeKillIssued, _16_cleanupComplete, _44_supervised)
	return result
}
func (_static *CompanionStruct_Default___) PollProcess(proc _dafny.Int, startedMs _dafny.Int, outputBytesBefore _dafny.Int, waitMs _dafny.Int, phase m_ConfluxResourceCeilings.ExecutionPhase, budget m_ConfluxResourceCeilings.ResourceCeilings) (_dafny.Sequence, bool, ProcessReadState, m_ConfluxResourceCeilings.ResourceObservation, _dafny.Int) {
	var chunk _dafny.Sequence = _dafny.EmptySeq
	_ = chunk
	var available bool = false
	_ = available
	var state ProcessReadState = Companion_ProcessReadState_.Default()
	_ = state
	var observation m_ConfluxResourceCeilings.ResourceObservation = m_ConfluxResourceCeilings.Companion_ResourceObservation_.Default()
	_ = observation
	var outputBytesAfter _dafny.Int = _dafny.Zero
	_ = outputBytesAfter
	var _0_open bool = false
	_ = _0_open
	var _1_combinedOutputBytesDelta _dafny.Int = _dafny.Zero
	_ = _1_combinedOutputBytesDelta
	var _out0 _dafny.Sequence
	_ = _out0
	var _out1 bool
	_ = _out1
	var _out2 bool
	_ = _out2
	var _out3 _dafny.Int
	_ = _out3
	_out0, _out1, _out2, _out3 = m_ConfluxIoCapability.Companion_Default___.ReadProcessOutputWithin(proc, waitMs)
	chunk = _out0
	available = _out1
	_0_open = _out2
	_1_combinedOutputBytesDelta = _out3
	var _2_now _dafny.Int
	_ = _2_now
	var _out4 _dafny.Int
	_ = _out4
	_out4 = m_ConfluxIoCapability.Companion_Default___.MonotonicTimeMs()
	_2_now = _out4
	var _3_handleFound bool
	_ = _3_handleFound
	var _4_exitObserved bool
	_ = _4_exitObserved
	var _5_exitCode _dafny.Int
	_ = _5_exitCode
	var _out5 bool
	_ = _out5
	var _out6 bool
	_ = _out6
	var _out7 _dafny.Int
	_ = _out7
	_out5, _out6, _out7 = m_ConfluxIoCapability.Companion_Default___.ProcessExitStatus(proc)
	_3_handleFound = _out5
	_4_exitObserved = _out6
	_5_exitCode = _out7
	if _4_exitObserved {
		var _6_finalChunk _dafny.Sequence
		_ = _6_finalChunk
		var _7_finalAvailable bool
		_ = _7_finalAvailable
		var _8_finalOpen bool
		_ = _8_finalOpen
		var _9_finalOutputBytesDelta _dafny.Int
		_ = _9_finalOutputBytesDelta
		var _out8 _dafny.Sequence
		_ = _out8
		var _out9 bool
		_ = _out9
		var _out10 bool
		_ = _out10
		var _out11 _dafny.Int
		_ = _out11
		_out8, _out9, _out10, _out11 = m_ConfluxIoCapability.Companion_Default___.ReadProcessOutputWithin(proc, _dafny.Zero)
		_6_finalChunk = _out8
		_7_finalAvailable = _out9
		_8_finalOpen = _out10
		_9_finalOutputBytesDelta = _out11
		chunk = Companion_Default___.FlatByteConcat(chunk, _6_finalChunk)
		available = (available) || (_7_finalAvailable)
		_0_open = _8_finalOpen
		_1_combinedOutputBytesDelta = (_1_combinedOutputBytesDelta).Plus(_9_finalOutputBytesDelta)
	}
	outputBytesAfter = (outputBytesBefore).Plus(Companion_Default___.NonNegative(_1_combinedOutputBytesDelta))
	var _10_memoryPresent bool
	_ = _10_memoryPresent
	_10_memoryPresent = false
	var _11_memoryMb _dafny.Int
	_ = _11_memoryMb
	_11_memoryMb = _dafny.Zero
	if m_ConfluxResourceCeilings.Companion_Default___.MemoryMeasurementRequired(budget, phase) {
		var _out12 bool
		_ = _out12
		var _out13 _dafny.Int
		_ = _out13
		_out12, _out13 = m_ConfluxIoCapability.Companion_Default___.ProcessRssMb(proc)
		_10_memoryPresent = _out12
		_11_memoryMb = _out13
	}
	var _12_memory m_ConfluxResourceCeilings.ResourceMeasurement
	_ = _12_memory
	if _10_memoryPresent {
		_12_memory = m_ConfluxResourceCeilings.Companion_Default___.CompleteMeasurement(Companion_Default___.NonNegative(_11_memoryMb))
	} else if m_ConfluxResourceCeilings.Companion_Default___.MemoryMeasurementRequired(budget, phase) {
		_12_memory = m_ConfluxResourceCeilings.Companion_Default___.UnavailableMeasurement(_dafny.UnicodeSeqOfUtf8Bytes("persistent process-tree RSS unavailable"))
	} else {
		_12_memory = m_ConfluxResourceCeilings.Companion_Default___.NotRequestedMeasurement()
	}
	observation = m_ConfluxResourceCeilings.Companion_ResourceObservation_.Create_ResourceObservation_(Companion_Default___.NonNegative((_2_now).Minus(startedMs)), _12_memory, m_ConfluxResourceCeilings.Companion_Default___.CompleteMeasurement(outputBytesAfter))
	var _13_terminal bool
	_ = _13_terminal
	_13_terminal = (_4_exitObserved) || ((!(_0_open)) && (!(_3_handleFound)))
	var _14_verdict m_ConfluxResourceCeilings.ResourceVerdict
	_ = _14_verdict
	if _13_terminal {
		_14_verdict = m_ConfluxResourceCeilings.Companion_Default___.DecidePostHoc(budget, phase, observation)
	} else {
		_14_verdict = m_ConfluxResourceCeilings.Companion_Default___.DecideInFlight(budget, phase, observation)
	}
	if (_13_terminal) && (available) {
		state = Companion_ProcessReadState_.Create_ProcessReadOpen_()
	} else if _4_exitObserved {
		state = Companion_ProcessReadState_.Create_ProcessReadExited_(_5_exitCode, _14_verdict)
	} else if (!(_0_open)) && (!(_3_handleFound)) {
		state = Companion_ProcessReadState_.Create_ProcessReadClosed_()
	} else if ((_14_verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementKill_()) {
		var _15_termination ProcessTerminationReceipt
		_ = _15_termination
		var _out14 ProcessTerminationReceipt
		_ = _out14
		_out14 = Companion_Default___.TerminateProcessWithin(proc, Companion_Default___.ProcessCleanupDeadlineMs())
		_15_termination = _out14
		chunk = Companion_Default___.FlatByteConcat(chunk, (_15_termination).Dtor_terminalChunk())
		available = (available) || ((_15_termination).Dtor_terminalAvailable())
		outputBytesAfter = (outputBytesAfter).Plus((_15_termination).Dtor_combinedOutputBytesDelta())
		var _16_cleanupNow _dafny.Int
		_ = _16_cleanupNow
		var _out15 _dafny.Int
		_ = _out15
		_out15 = m_ConfluxIoCapability.Companion_Default___.MonotonicTimeMs()
		_16_cleanupNow = _out15
		observation = m_ConfluxResourceCeilings.Companion_ResourceObservation_.Create_ResourceObservation_(Companion_Default___.NonNegative((_16_cleanupNow).Minus(startedMs)), _12_memory, m_ConfluxResourceCeilings.Companion_Default___.CompleteMeasurement(outputBytesAfter))
		state = Companion_ProcessReadState_.Create_ProcessReadKilled_(_14_verdict, _15_termination)
	} else if (_0_open) || (_3_handleFound) {
		state = Companion_ProcessReadState_.Create_ProcessReadOpen_()
	} else {
		state = Companion_ProcessReadState_.Create_ProcessReadClosed_()
	}
	return chunk, available, state, observation, outputBytesAfter
}
func (_static *CompanionStruct_Default___) CancelProcess(proc _dafny.Int, startedMs _dafny.Int, outputBytesBefore _dafny.Int, phase m_ConfluxResourceCeilings.ExecutionPhase, budget m_ConfluxResourceCeilings.ResourceCeilings) ProcessCancellationResult {
	var receipt ProcessCancellationResult = Companion_ProcessCancellationResult_.Default()
	_ = receipt
	var _0_chunk _dafny.Sequence
	_ = _0_chunk
	var _1_available bool
	_ = _1_available
	var _2_open bool
	_ = _2_open
	var _3_combinedOutputBytesDelta _dafny.Int
	_ = _3_combinedOutputBytesDelta
	var _out0 _dafny.Sequence
	_ = _out0
	var _out1 bool
	_ = _out1
	var _out2 bool
	_ = _out2
	var _out3 _dafny.Int
	_ = _out3
	_out0, _out1, _out2, _out3 = m_ConfluxIoCapability.Companion_Default___.ReadProcessOutputWithin(proc, _dafny.Zero)
	_0_chunk = _out0
	_1_available = _out1
	_2_open = _out2
	_3_combinedOutputBytesDelta = _out3
	var _4_outputBytesAfter _dafny.Int
	_ = _4_outputBytesAfter
	_4_outputBytesAfter = (outputBytesBefore).Plus(Companion_Default___.NonNegative(_3_combinedOutputBytesDelta))
	var _5_now _dafny.Int
	_ = _5_now
	var _out4 _dafny.Int
	_ = _out4
	_out4 = m_ConfluxIoCapability.Companion_Default___.MonotonicTimeMs()
	_5_now = _out4
	var _6_memoryPresent bool
	_ = _6_memoryPresent
	_6_memoryPresent = false
	var _7_memoryMb _dafny.Int
	_ = _7_memoryMb
	_7_memoryMb = _dafny.Zero
	if m_ConfluxResourceCeilings.Companion_Default___.MemoryMeasurementRequired(budget, phase) {
		var _out5 bool
		_ = _out5
		var _out6 _dafny.Int
		_ = _out6
		_out5, _out6 = m_ConfluxIoCapability.Companion_Default___.ProcessRssMb(proc)
		_6_memoryPresent = _out5
		_7_memoryMb = _out6
	}
	var _8_memory m_ConfluxResourceCeilings.ResourceMeasurement
	_ = _8_memory
	if _6_memoryPresent {
		_8_memory = m_ConfluxResourceCeilings.Companion_Default___.CompleteMeasurement(Companion_Default___.NonNegative(_7_memoryMb))
	} else if m_ConfluxResourceCeilings.Companion_Default___.MemoryMeasurementRequired(budget, phase) {
		_8_memory = m_ConfluxResourceCeilings.Companion_Default___.UnavailableMeasurement(_dafny.UnicodeSeqOfUtf8Bytes("persistent process-tree RSS unavailable at cancellation"))
	} else {
		_8_memory = m_ConfluxResourceCeilings.Companion_Default___.NotRequestedMeasurement()
	}
	var _9_beforeObservation m_ConfluxResourceCeilings.ResourceObservation
	_ = _9_beforeObservation
	_9_beforeObservation = m_ConfluxResourceCeilings.Companion_ResourceObservation_.Create_ResourceObservation_(Companion_Default___.NonNegative((_5_now).Minus(startedMs)), _8_memory, m_ConfluxResourceCeilings.Companion_Default___.CompleteMeasurement(_4_outputBytesAfter))
	var _10_termination ProcessTerminationReceipt
	_ = _10_termination
	var _out7 ProcessTerminationReceipt
	_ = _out7
	_out7 = Companion_Default___.TerminateProcessWithin(proc, Companion_Default___.ProcessCleanupDeadlineMs())
	_10_termination = _out7
	_0_chunk = Companion_Default___.FlatByteConcat(_0_chunk, (_10_termination).Dtor_terminalChunk())
	_1_available = (_1_available) || ((_10_termination).Dtor_terminalAvailable())
	_4_outputBytesAfter = (_4_outputBytesAfter).Plus((_10_termination).Dtor_combinedOutputBytesDelta())
	var _11_cancellationNow _dafny.Int
	_ = _11_cancellationNow
	var _out8 _dafny.Int
	_ = _out8
	_out8 = m_ConfluxIoCapability.Companion_Default___.MonotonicTimeMs()
	_11_cancellationNow = _out8
	var _12_finalObservation m_ConfluxResourceCeilings.ResourceObservation
	_ = _12_finalObservation
	_12_finalObservation = m_ConfluxResourceCeilings.Companion_ResourceObservation_.Create_ResourceObservation_(Companion_Default___.NonNegative((_11_cancellationNow).Minus(startedMs)), _8_memory, m_ConfluxResourceCeilings.Companion_Default___.CompleteMeasurement(_4_outputBytesAfter))
	var _13_outcome ProcessCancellationOutcome
	_ = _13_outcome
	if ((_10_termination).Dtor_cleanup()).Is_CleanupFailed() {
		_13_outcome = Companion_ProcessCancellationOutcome_.Create_ProcessTerminationFailed_()
	} else if !((_10_termination).Dtor_handleFound()) {
		_13_outcome = Companion_ProcessCancellationOutcome_.Create_ProcessAlreadyClosed_()
	} else {
		_13_outcome = Companion_ProcessCancellationOutcome_.Create_ProcessCancelled_()
	}
	receipt = Companion_ProcessCancellationResult_.Create_ProcessCancellationResult_(_13_outcome, _0_chunk, _1_available, _9_beforeObservation, _12_finalObservation, _4_outputBytesAfter, _10_termination)
	return receipt
}
func (_static *CompanionStruct_Default___) ProcessCleanupDeadlineMs() _dafny.Int {
	return _dafny.IntOfInt64(2000)
}

// End of class Default__

// Definition of datatype ProcessReadState
type ProcessReadState struct {
	Data_ProcessReadState_
}

func (_this ProcessReadState) Get_() Data_ProcessReadState_ {
	return _this.Data_ProcessReadState_
}

type Data_ProcessReadState_ interface {
	isProcessReadState()
}

type CompanionStruct_ProcessReadState_ struct {
}

var Companion_ProcessReadState_ = CompanionStruct_ProcessReadState_{}

type ProcessReadState_ProcessReadOpen struct {
}

func (ProcessReadState_ProcessReadOpen) isProcessReadState() {}

func (CompanionStruct_ProcessReadState_) Create_ProcessReadOpen_() ProcessReadState {
	return ProcessReadState{ProcessReadState_ProcessReadOpen{}}
}

func (_this ProcessReadState) Is_ProcessReadOpen() bool {
	_, ok := _this.Get_().(ProcessReadState_ProcessReadOpen)
	return ok
}

type ProcessReadState_ProcessReadClosed struct {
}

func (ProcessReadState_ProcessReadClosed) isProcessReadState() {}

func (CompanionStruct_ProcessReadState_) Create_ProcessReadClosed_() ProcessReadState {
	return ProcessReadState{ProcessReadState_ProcessReadClosed{}}
}

func (_this ProcessReadState) Is_ProcessReadClosed() bool {
	_, ok := _this.Get_().(ProcessReadState_ProcessReadClosed)
	return ok
}

type ProcessReadState_ProcessReadExited struct {
	ExitCode _dafny.Int
	Verdict  m_ConfluxResourceCeilings.ResourceVerdict
}

func (ProcessReadState_ProcessReadExited) isProcessReadState() {}

func (CompanionStruct_ProcessReadState_) Create_ProcessReadExited_(ExitCode _dafny.Int, Verdict m_ConfluxResourceCeilings.ResourceVerdict) ProcessReadState {
	return ProcessReadState{ProcessReadState_ProcessReadExited{ExitCode, Verdict}}
}

func (_this ProcessReadState) Is_ProcessReadExited() bool {
	_, ok := _this.Get_().(ProcessReadState_ProcessReadExited)
	return ok
}

type ProcessReadState_ProcessReadKilled struct {
	Verdict     m_ConfluxResourceCeilings.ResourceVerdict
	Termination ProcessTerminationReceipt
}

func (ProcessReadState_ProcessReadKilled) isProcessReadState() {}

func (CompanionStruct_ProcessReadState_) Create_ProcessReadKilled_(Verdict m_ConfluxResourceCeilings.ResourceVerdict, Termination ProcessTerminationReceipt) ProcessReadState {
	return ProcessReadState{ProcessReadState_ProcessReadKilled{Verdict, Termination}}
}

func (_this ProcessReadState) Is_ProcessReadKilled() bool {
	_, ok := _this.Get_().(ProcessReadState_ProcessReadKilled)
	return ok
}

func (CompanionStruct_ProcessReadState_) Default() ProcessReadState {
	return Companion_ProcessReadState_.Create_ProcessReadOpen_()
}

func (_this ProcessReadState) Dtor_exitCode() _dafny.Int {
	return _this.Get_().(ProcessReadState_ProcessReadExited).ExitCode
}

func (_this ProcessReadState) Dtor_verdict() m_ConfluxResourceCeilings.ResourceVerdict {
	switch data := _this.Get_().(type) {
	case ProcessReadState_ProcessReadExited:
		return data.Verdict
	default:
		return data.(ProcessReadState_ProcessReadKilled).Verdict
	}
}

func (_this ProcessReadState) Dtor_termination() ProcessTerminationReceipt {
	return _this.Get_().(ProcessReadState_ProcessReadKilled).Termination
}

func (_this ProcessReadState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ProcessReadState_ProcessReadOpen:
		{
			return "ConfluxResourceSupervisor.ProcessReadState.ProcessReadOpen"
		}
	case ProcessReadState_ProcessReadClosed:
		{
			return "ConfluxResourceSupervisor.ProcessReadState.ProcessReadClosed"
		}
	case ProcessReadState_ProcessReadExited:
		{
			return "ConfluxResourceSupervisor.ProcessReadState.ProcessReadExited" + "(" + _dafny.String(data.ExitCode) + ", " + _dafny.String(data.Verdict) + ")"
		}
	case ProcessReadState_ProcessReadKilled:
		{
			return "ConfluxResourceSupervisor.ProcessReadState.ProcessReadKilled" + "(" + _dafny.String(data.Verdict) + ", " + _dafny.String(data.Termination) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ProcessReadState) Equals(other ProcessReadState) bool {
	switch data1 := _this.Get_().(type) {
	case ProcessReadState_ProcessReadOpen:
		{
			_, ok := other.Get_().(ProcessReadState_ProcessReadOpen)
			return ok
		}
	case ProcessReadState_ProcessReadClosed:
		{
			_, ok := other.Get_().(ProcessReadState_ProcessReadClosed)
			return ok
		}
	case ProcessReadState_ProcessReadExited:
		{
			data2, ok := other.Get_().(ProcessReadState_ProcessReadExited)
			return ok && data1.ExitCode.Cmp(data2.ExitCode) == 0 && data1.Verdict.Equals(data2.Verdict)
		}
	case ProcessReadState_ProcessReadKilled:
		{
			data2, ok := other.Get_().(ProcessReadState_ProcessReadKilled)
			return ok && data1.Verdict.Equals(data2.Verdict) && data1.Termination.Equals(data2.Termination)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ProcessReadState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ProcessReadState)
	return ok && _this.Equals(typed)
}

func Type_ProcessReadState_() _dafny.TypeDescriptor {
	return type_ProcessReadState_{}
}

type type_ProcessReadState_ struct {
}

func (_this type_ProcessReadState_) Default() interface{} {
	return Companion_ProcessReadState_.Default()
}

func (_this type_ProcessReadState_) String() string {
	return "ConfluxResourceSupervisor.ProcessReadState"
}
func (_this ProcessReadState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ProcessReadState{}

// End of datatype ProcessReadState

// Definition of datatype ProcessTerminationReceipt
type ProcessTerminationReceipt struct {
	Data_ProcessTerminationReceipt_
}

func (_this ProcessTerminationReceipt) Get_() Data_ProcessTerminationReceipt_ {
	return _this.Data_ProcessTerminationReceipt_
}

type Data_ProcessTerminationReceipt_ interface {
	isProcessTerminationReceipt()
}

type CompanionStruct_ProcessTerminationReceipt_ struct {
}

var Companion_ProcessTerminationReceipt_ = CompanionStruct_ProcessTerminationReceipt_{}

type ProcessTerminationReceipt_ProcessTerminationReceipt struct {
	HandleFound              bool
	TreeKillIssued           bool
	Cleanup                  m_ConfluxSupervisedOperationResult.CleanupStatus
	TerminalChunk            _dafny.Sequence
	TerminalAvailable        bool
	CombinedOutputBytesDelta _dafny.Int
	Exit                     m_ConfluxSupervisedOperationResult.ExitStatus
}

func (ProcessTerminationReceipt_ProcessTerminationReceipt) isProcessTerminationReceipt() {}

func (CompanionStruct_ProcessTerminationReceipt_) Create_ProcessTerminationReceipt_(HandleFound bool, TreeKillIssued bool, Cleanup m_ConfluxSupervisedOperationResult.CleanupStatus, TerminalChunk _dafny.Sequence, TerminalAvailable bool, CombinedOutputBytesDelta _dafny.Int, Exit m_ConfluxSupervisedOperationResult.ExitStatus) ProcessTerminationReceipt {
	return ProcessTerminationReceipt{ProcessTerminationReceipt_ProcessTerminationReceipt{HandleFound, TreeKillIssued, Cleanup, TerminalChunk, TerminalAvailable, CombinedOutputBytesDelta, Exit}}
}

func (_this ProcessTerminationReceipt) Is_ProcessTerminationReceipt() bool {
	_, ok := _this.Get_().(ProcessTerminationReceipt_ProcessTerminationReceipt)
	return ok
}

func (CompanionStruct_ProcessTerminationReceipt_) Default() ProcessTerminationReceipt {
	return Companion_ProcessTerminationReceipt_.Create_ProcessTerminationReceipt_(false, false, m_ConfluxSupervisedOperationResult.Companion_CleanupStatus_.Default(), _dafny.EmptySeq, false, _dafny.Zero, m_ConfluxSupervisedOperationResult.Companion_ExitStatus_.Default())
}

func (_this ProcessTerminationReceipt) Dtor_handleFound() bool {
	return _this.Get_().(ProcessTerminationReceipt_ProcessTerminationReceipt).HandleFound
}

func (_this ProcessTerminationReceipt) Dtor_treeKillIssued() bool {
	return _this.Get_().(ProcessTerminationReceipt_ProcessTerminationReceipt).TreeKillIssued
}

func (_this ProcessTerminationReceipt) Dtor_cleanup() m_ConfluxSupervisedOperationResult.CleanupStatus {
	return _this.Get_().(ProcessTerminationReceipt_ProcessTerminationReceipt).Cleanup
}

func (_this ProcessTerminationReceipt) Dtor_terminalChunk() _dafny.Sequence {
	return _this.Get_().(ProcessTerminationReceipt_ProcessTerminationReceipt).TerminalChunk
}

func (_this ProcessTerminationReceipt) Dtor_terminalAvailable() bool {
	return _this.Get_().(ProcessTerminationReceipt_ProcessTerminationReceipt).TerminalAvailable
}

func (_this ProcessTerminationReceipt) Dtor_combinedOutputBytesDelta() _dafny.Int {
	return _this.Get_().(ProcessTerminationReceipt_ProcessTerminationReceipt).CombinedOutputBytesDelta
}

func (_this ProcessTerminationReceipt) Dtor_exit() m_ConfluxSupervisedOperationResult.ExitStatus {
	return _this.Get_().(ProcessTerminationReceipt_ProcessTerminationReceipt).Exit
}

func (_this ProcessTerminationReceipt) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ProcessTerminationReceipt_ProcessTerminationReceipt:
		{
			return "ConfluxResourceSupervisor.ProcessTerminationReceipt.ProcessTerminationReceipt" + "(" + _dafny.String(data.HandleFound) + ", " + _dafny.String(data.TreeKillIssued) + ", " + _dafny.String(data.Cleanup) + ", " + _dafny.String(data.TerminalChunk) + ", " + _dafny.String(data.TerminalAvailable) + ", " + _dafny.String(data.CombinedOutputBytesDelta) + ", " + _dafny.String(data.Exit) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ProcessTerminationReceipt) Equals(other ProcessTerminationReceipt) bool {
	switch data1 := _this.Get_().(type) {
	case ProcessTerminationReceipt_ProcessTerminationReceipt:
		{
			data2, ok := other.Get_().(ProcessTerminationReceipt_ProcessTerminationReceipt)
			return ok && data1.HandleFound == data2.HandleFound && data1.TreeKillIssued == data2.TreeKillIssued && data1.Cleanup.Equals(data2.Cleanup) && data1.TerminalChunk.Equals(data2.TerminalChunk) && data1.TerminalAvailable == data2.TerminalAvailable && data1.CombinedOutputBytesDelta.Cmp(data2.CombinedOutputBytesDelta) == 0 && data1.Exit.Equals(data2.Exit)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ProcessTerminationReceipt) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ProcessTerminationReceipt)
	return ok && _this.Equals(typed)
}

func Type_ProcessTerminationReceipt_() _dafny.TypeDescriptor {
	return type_ProcessTerminationReceipt_{}
}

type type_ProcessTerminationReceipt_ struct {
}

func (_this type_ProcessTerminationReceipt_) Default() interface{} {
	return Companion_ProcessTerminationReceipt_.Default()
}

func (_this type_ProcessTerminationReceipt_) String() string {
	return "ConfluxResourceSupervisor.ProcessTerminationReceipt"
}
func (_this ProcessTerminationReceipt) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ProcessTerminationReceipt{}

// End of datatype ProcessTerminationReceipt

// Definition of datatype ProcessCancellationOutcome
type ProcessCancellationOutcome struct {
	Data_ProcessCancellationOutcome_
}

func (_this ProcessCancellationOutcome) Get_() Data_ProcessCancellationOutcome_ {
	return _this.Data_ProcessCancellationOutcome_
}

type Data_ProcessCancellationOutcome_ interface {
	isProcessCancellationOutcome()
}

type CompanionStruct_ProcessCancellationOutcome_ struct {
}

var Companion_ProcessCancellationOutcome_ = CompanionStruct_ProcessCancellationOutcome_{}

type ProcessCancellationOutcome_ProcessCancelled struct {
}

func (ProcessCancellationOutcome_ProcessCancelled) isProcessCancellationOutcome() {}

func (CompanionStruct_ProcessCancellationOutcome_) Create_ProcessCancelled_() ProcessCancellationOutcome {
	return ProcessCancellationOutcome{ProcessCancellationOutcome_ProcessCancelled{}}
}

func (_this ProcessCancellationOutcome) Is_ProcessCancelled() bool {
	_, ok := _this.Get_().(ProcessCancellationOutcome_ProcessCancelled)
	return ok
}

type ProcessCancellationOutcome_ProcessAlreadyClosed struct {
}

func (ProcessCancellationOutcome_ProcessAlreadyClosed) isProcessCancellationOutcome() {}

func (CompanionStruct_ProcessCancellationOutcome_) Create_ProcessAlreadyClosed_() ProcessCancellationOutcome {
	return ProcessCancellationOutcome{ProcessCancellationOutcome_ProcessAlreadyClosed{}}
}

func (_this ProcessCancellationOutcome) Is_ProcessAlreadyClosed() bool {
	_, ok := _this.Get_().(ProcessCancellationOutcome_ProcessAlreadyClosed)
	return ok
}

type ProcessCancellationOutcome_ProcessTerminationFailed struct {
}

func (ProcessCancellationOutcome_ProcessTerminationFailed) isProcessCancellationOutcome() {}

func (CompanionStruct_ProcessCancellationOutcome_) Create_ProcessTerminationFailed_() ProcessCancellationOutcome {
	return ProcessCancellationOutcome{ProcessCancellationOutcome_ProcessTerminationFailed{}}
}

func (_this ProcessCancellationOutcome) Is_ProcessTerminationFailed() bool {
	_, ok := _this.Get_().(ProcessCancellationOutcome_ProcessTerminationFailed)
	return ok
}

func (CompanionStruct_ProcessCancellationOutcome_) Default() ProcessCancellationOutcome {
	return Companion_ProcessCancellationOutcome_.Create_ProcessCancelled_()
}

func (_ CompanionStruct_ProcessCancellationOutcome_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_ProcessCancellationOutcome_.Create_ProcessCancelled_(), true
		case 1:
			return Companion_ProcessCancellationOutcome_.Create_ProcessAlreadyClosed_(), true
		case 2:
			return Companion_ProcessCancellationOutcome_.Create_ProcessTerminationFailed_(), true
		default:
			return ProcessCancellationOutcome{}, false
		}
	}
}

func (_this ProcessCancellationOutcome) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case ProcessCancellationOutcome_ProcessCancelled:
		{
			return "ConfluxResourceSupervisor.ProcessCancellationOutcome.ProcessCancelled"
		}
	case ProcessCancellationOutcome_ProcessAlreadyClosed:
		{
			return "ConfluxResourceSupervisor.ProcessCancellationOutcome.ProcessAlreadyClosed"
		}
	case ProcessCancellationOutcome_ProcessTerminationFailed:
		{
			return "ConfluxResourceSupervisor.ProcessCancellationOutcome.ProcessTerminationFailed"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ProcessCancellationOutcome) Equals(other ProcessCancellationOutcome) bool {
	switch _this.Get_().(type) {
	case ProcessCancellationOutcome_ProcessCancelled:
		{
			_, ok := other.Get_().(ProcessCancellationOutcome_ProcessCancelled)
			return ok
		}
	case ProcessCancellationOutcome_ProcessAlreadyClosed:
		{
			_, ok := other.Get_().(ProcessCancellationOutcome_ProcessAlreadyClosed)
			return ok
		}
	case ProcessCancellationOutcome_ProcessTerminationFailed:
		{
			_, ok := other.Get_().(ProcessCancellationOutcome_ProcessTerminationFailed)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ProcessCancellationOutcome) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ProcessCancellationOutcome)
	return ok && _this.Equals(typed)
}

func Type_ProcessCancellationOutcome_() _dafny.TypeDescriptor {
	return type_ProcessCancellationOutcome_{}
}

type type_ProcessCancellationOutcome_ struct {
}

func (_this type_ProcessCancellationOutcome_) Default() interface{} {
	return Companion_ProcessCancellationOutcome_.Default()
}

func (_this type_ProcessCancellationOutcome_) String() string {
	return "ConfluxResourceSupervisor.ProcessCancellationOutcome"
}
func (_this ProcessCancellationOutcome) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ProcessCancellationOutcome{}

// End of datatype ProcessCancellationOutcome

// Definition of datatype ProcessCancellationResult
type ProcessCancellationResult struct {
	Data_ProcessCancellationResult_
}

func (_this ProcessCancellationResult) Get_() Data_ProcessCancellationResult_ {
	return _this.Data_ProcessCancellationResult_
}

type Data_ProcessCancellationResult_ interface {
	isProcessCancellationResult()
}

type CompanionStruct_ProcessCancellationResult_ struct {
}

var Companion_ProcessCancellationResult_ = CompanionStruct_ProcessCancellationResult_{}

type ProcessCancellationResult_ProcessCancellationResult struct {
	Outcome             ProcessCancellationOutcome
	Chunk               _dafny.Sequence
	Available           bool
	BeforeObservation   m_ConfluxResourceCeilings.ResourceObservation
	FinalObservation    m_ConfluxResourceCeilings.ResourceObservation
	CombinedOutputBytes _dafny.Int
	Termination         ProcessTerminationReceipt
}

func (ProcessCancellationResult_ProcessCancellationResult) isProcessCancellationResult() {}

func (CompanionStruct_ProcessCancellationResult_) Create_ProcessCancellationResult_(Outcome ProcessCancellationOutcome, Chunk _dafny.Sequence, Available bool, BeforeObservation m_ConfluxResourceCeilings.ResourceObservation, FinalObservation m_ConfluxResourceCeilings.ResourceObservation, CombinedOutputBytes _dafny.Int, Termination ProcessTerminationReceipt) ProcessCancellationResult {
	return ProcessCancellationResult{ProcessCancellationResult_ProcessCancellationResult{Outcome, Chunk, Available, BeforeObservation, FinalObservation, CombinedOutputBytes, Termination}}
}

func (_this ProcessCancellationResult) Is_ProcessCancellationResult() bool {
	_, ok := _this.Get_().(ProcessCancellationResult_ProcessCancellationResult)
	return ok
}

func (CompanionStruct_ProcessCancellationResult_) Default() ProcessCancellationResult {
	return Companion_ProcessCancellationResult_.Create_ProcessCancellationResult_(Companion_ProcessCancellationOutcome_.Default(), _dafny.EmptySeq, false, m_ConfluxResourceCeilings.Companion_ResourceObservation_.Default(), m_ConfluxResourceCeilings.Companion_ResourceObservation_.Default(), _dafny.Zero, Companion_ProcessTerminationReceipt_.Default())
}

func (_this ProcessCancellationResult) Dtor_outcome() ProcessCancellationOutcome {
	return _this.Get_().(ProcessCancellationResult_ProcessCancellationResult).Outcome
}

func (_this ProcessCancellationResult) Dtor_chunk() _dafny.Sequence {
	return _this.Get_().(ProcessCancellationResult_ProcessCancellationResult).Chunk
}

func (_this ProcessCancellationResult) Dtor_available() bool {
	return _this.Get_().(ProcessCancellationResult_ProcessCancellationResult).Available
}

func (_this ProcessCancellationResult) Dtor_beforeObservation() m_ConfluxResourceCeilings.ResourceObservation {
	return _this.Get_().(ProcessCancellationResult_ProcessCancellationResult).BeforeObservation
}

func (_this ProcessCancellationResult) Dtor_finalObservation() m_ConfluxResourceCeilings.ResourceObservation {
	return _this.Get_().(ProcessCancellationResult_ProcessCancellationResult).FinalObservation
}

func (_this ProcessCancellationResult) Dtor_combinedOutputBytes() _dafny.Int {
	return _this.Get_().(ProcessCancellationResult_ProcessCancellationResult).CombinedOutputBytes
}

func (_this ProcessCancellationResult) Dtor_termination() ProcessTerminationReceipt {
	return _this.Get_().(ProcessCancellationResult_ProcessCancellationResult).Termination
}

func (_this ProcessCancellationResult) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ProcessCancellationResult_ProcessCancellationResult:
		{
			return "ConfluxResourceSupervisor.ProcessCancellationResult.ProcessCancellationResult" + "(" + _dafny.String(data.Outcome) + ", " + _dafny.String(data.Chunk) + ", " + _dafny.String(data.Available) + ", " + _dafny.String(data.BeforeObservation) + ", " + _dafny.String(data.FinalObservation) + ", " + _dafny.String(data.CombinedOutputBytes) + ", " + _dafny.String(data.Termination) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ProcessCancellationResult) Equals(other ProcessCancellationResult) bool {
	switch data1 := _this.Get_().(type) {
	case ProcessCancellationResult_ProcessCancellationResult:
		{
			data2, ok := other.Get_().(ProcessCancellationResult_ProcessCancellationResult)
			return ok && data1.Outcome.Equals(data2.Outcome) && data1.Chunk.Equals(data2.Chunk) && data1.Available == data2.Available && data1.BeforeObservation.Equals(data2.BeforeObservation) && data1.FinalObservation.Equals(data2.FinalObservation) && data1.CombinedOutputBytes.Cmp(data2.CombinedOutputBytes) == 0 && data1.Termination.Equals(data2.Termination)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ProcessCancellationResult) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ProcessCancellationResult)
	return ok && _this.Equals(typed)
}

func Type_ProcessCancellationResult_() _dafny.TypeDescriptor {
	return type_ProcessCancellationResult_{}
}

type type_ProcessCancellationResult_ struct {
}

func (_this type_ProcessCancellationResult_) Default() interface{} {
	return Companion_ProcessCancellationResult_.Default()
}

func (_this type_ProcessCancellationResult_) String() string {
	return "ConfluxResourceSupervisor.ProcessCancellationResult"
}
func (_this ProcessCancellationResult) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ProcessCancellationResult{}

// End of datatype ProcessCancellationResult

// Definition of datatype SupervisedProcessResult
type SupervisedProcessResult struct {
	Data_SupervisedProcessResult_
}

func (_this SupervisedProcessResult) Get_() Data_SupervisedProcessResult_ {
	return _this.Data_SupervisedProcessResult_
}

type Data_SupervisedProcessResult_ interface {
	isSupervisedProcessResult()
}

type CompanionStruct_SupervisedProcessResult_ struct {
}

var Companion_SupervisedProcessResult_ = CompanionStruct_SupervisedProcessResult_{}

type SupervisedProcessResult_SupervisedProcessResult struct {
	Code                _dafny.Int
	Output              _dafny.Sequence
	Error               _dafny.Sequence
	Observation         m_ConfluxResourceCeilings.ResourceObservation
	Verdict             m_ConfluxResourceCeilings.ResourceVerdict
	CombinedOutputBytes _dafny.Int
	TreeKillIssued      bool
	CleanupComplete     bool
	Supervised          m_ConfluxSupervisedOperationResult.SupervisedOperationResultV1
}

func (SupervisedProcessResult_SupervisedProcessResult) isSupervisedProcessResult() {}

func (CompanionStruct_SupervisedProcessResult_) Create_SupervisedProcessResult_(Code _dafny.Int, Output _dafny.Sequence, Error _dafny.Sequence, Observation m_ConfluxResourceCeilings.ResourceObservation, Verdict m_ConfluxResourceCeilings.ResourceVerdict, CombinedOutputBytes _dafny.Int, TreeKillIssued bool, CleanupComplete bool, Supervised m_ConfluxSupervisedOperationResult.SupervisedOperationResultV1) SupervisedProcessResult {
	return SupervisedProcessResult{SupervisedProcessResult_SupervisedProcessResult{Code, Output, Error, Observation, Verdict, CombinedOutputBytes, TreeKillIssued, CleanupComplete, Supervised}}
}

func (_this SupervisedProcessResult) Is_SupervisedProcessResult() bool {
	_, ok := _this.Get_().(SupervisedProcessResult_SupervisedProcessResult)
	return ok
}

func (CompanionStruct_SupervisedProcessResult_) Default() SupervisedProcessResult {
	return Companion_SupervisedProcessResult_.Create_SupervisedProcessResult_(_dafny.Zero, _dafny.EmptySeq, _dafny.EmptySeq, m_ConfluxResourceCeilings.Companion_ResourceObservation_.Default(), m_ConfluxResourceCeilings.Companion_ResourceVerdict_.Default(), _dafny.Zero, false, false, m_ConfluxSupervisedOperationResult.Companion_SupervisedOperationResultV1_.Default())
}

func (_this SupervisedProcessResult) Dtor_code() _dafny.Int {
	return _this.Get_().(SupervisedProcessResult_SupervisedProcessResult).Code
}

func (_this SupervisedProcessResult) Dtor_output() _dafny.Sequence {
	return _this.Get_().(SupervisedProcessResult_SupervisedProcessResult).Output
}

func (_this SupervisedProcessResult) Dtor_error() _dafny.Sequence {
	return _this.Get_().(SupervisedProcessResult_SupervisedProcessResult).Error
}

func (_this SupervisedProcessResult) Dtor_observation() m_ConfluxResourceCeilings.ResourceObservation {
	return _this.Get_().(SupervisedProcessResult_SupervisedProcessResult).Observation
}

func (_this SupervisedProcessResult) Dtor_verdict() m_ConfluxResourceCeilings.ResourceVerdict {
	return _this.Get_().(SupervisedProcessResult_SupervisedProcessResult).Verdict
}

func (_this SupervisedProcessResult) Dtor_combinedOutputBytes() _dafny.Int {
	return _this.Get_().(SupervisedProcessResult_SupervisedProcessResult).CombinedOutputBytes
}

func (_this SupervisedProcessResult) Dtor_treeKillIssued() bool {
	return _this.Get_().(SupervisedProcessResult_SupervisedProcessResult).TreeKillIssued
}

func (_this SupervisedProcessResult) Dtor_cleanupComplete() bool {
	return _this.Get_().(SupervisedProcessResult_SupervisedProcessResult).CleanupComplete
}

func (_this SupervisedProcessResult) Dtor_supervised() m_ConfluxSupervisedOperationResult.SupervisedOperationResultV1 {
	return _this.Get_().(SupervisedProcessResult_SupervisedProcessResult).Supervised
}

func (_this SupervisedProcessResult) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SupervisedProcessResult_SupervisedProcessResult:
		{
			return "ConfluxResourceSupervisor.SupervisedProcessResult.SupervisedProcessResult" + "(" + _dafny.String(data.Code) + ", " + data.Output.VerbatimString(true) + ", " + data.Error.VerbatimString(true) + ", " + _dafny.String(data.Observation) + ", " + _dafny.String(data.Verdict) + ", " + _dafny.String(data.CombinedOutputBytes) + ", " + _dafny.String(data.TreeKillIssued) + ", " + _dafny.String(data.CleanupComplete) + ", " + _dafny.String(data.Supervised) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SupervisedProcessResult) Equals(other SupervisedProcessResult) bool {
	switch data1 := _this.Get_().(type) {
	case SupervisedProcessResult_SupervisedProcessResult:
		{
			data2, ok := other.Get_().(SupervisedProcessResult_SupervisedProcessResult)
			return ok && data1.Code.Cmp(data2.Code) == 0 && data1.Output.Equals(data2.Output) && data1.Error.Equals(data2.Error) && data1.Observation.Equals(data2.Observation) && data1.Verdict.Equals(data2.Verdict) && data1.CombinedOutputBytes.Cmp(data2.CombinedOutputBytes) == 0 && data1.TreeKillIssued == data2.TreeKillIssued && data1.CleanupComplete == data2.CleanupComplete && data1.Supervised.Equals(data2.Supervised)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SupervisedProcessResult) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SupervisedProcessResult)
	return ok && _this.Equals(typed)
}

func Type_SupervisedProcessResult_() _dafny.TypeDescriptor {
	return type_SupervisedProcessResult_{}
}

type type_SupervisedProcessResult_ struct {
}

func (_this type_SupervisedProcessResult_) Default() interface{} {
	return Companion_SupervisedProcessResult_.Default()
}

func (_this type_SupervisedProcessResult_) String() string {
	return "ConfluxResourceSupervisor.SupervisedProcessResult"
}
func (_this SupervisedProcessResult) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SupervisedProcessResult{}

// End of datatype SupervisedProcessResult
