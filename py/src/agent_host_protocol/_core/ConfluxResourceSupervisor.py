import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_ as module_
import _dafny as _dafny
import System_ as System_
import ConfluxCodec as ConfluxCodec
import ConfluxContract as ConfluxContract
import AhpSkeleton as AhpSkeleton
import ResourceWatch as ResourceWatch
import Canvas as Canvas
import ConfluxOperators as ConfluxOperators
import ConfluxPrune as ConfluxPrune
import ConfluxSeqRoute as ConfluxSeqRoute
import Changeset as Changeset
import Annotations as Annotations
import Terminal as Terminal
import Session as Session
import ConfluxOrderedLog as ConfluxOrderedLog
import Chat as Chat
import ClientMain as ClientMain
import ConfluxSemanticGraphIdentity as ConfluxSemanticGraphIdentity
import ConfluxContentAddress as ConfluxContentAddress
import ConfluxSemanticGraphContract as ConfluxSemanticGraphContract
import ConfluxResourceCeilings as ConfluxResourceCeilings
import ConfluxSupervisedOperationResult as ConfluxSupervisedOperationResult
import ConfluxTree as ConfluxTree
import ConfluxMerge as ConfluxMerge
import ConfluxRoute as ConfluxRoute
import ConfluxMapValues as ConfluxMapValues
import ConfluxDelimited as ConfluxDelimited
import ConfluxFixpoint as ConfluxFixpoint
import ConfluxGenericCodec as ConfluxGenericCodec
import ConfluxResolve as ConfluxResolve
import ConfluxWatermark as ConfluxWatermark
import ConfluxStateMachine as ConfluxStateMachine
import ConfluxProduct as ConfluxProduct
import ConfluxJoin as ConfluxJoin
import ConfluxDedupe as ConfluxDedupe
import ConfluxBatchRoute as ConfluxBatchRoute
import ConfluxSeqRouteMerge as ConfluxSeqRouteMerge
import ConfluxKeyedOps as ConfluxKeyedOps
import ConfluxFilterKeys as ConfluxFilterKeys
import ConfluxChannel as ConfluxChannel
import ConfluxDurableHistory as ConfluxDurableHistory
import ConfluxCommand as ConfluxCommand
import ConfluxChannelManifest as ConfluxChannelManifest
import ConfluxConvergence as ConfluxConvergence
import ConfluxBudgetConvergence as ConfluxBudgetConvergence
import ConfluxExtern as ConfluxExtern
import ConfluxCommandIdentityCapability as ConfluxCommandIdentityCapability
import ConfluxDecimalText as ConfluxDecimalText
import ConfluxJsonText as ConfluxJsonText
import ConfluxCommandIdentity as ConfluxCommandIdentity
import ConfluxIoCapability as ConfluxIoCapability

# Module: ConfluxResourceSupervisor

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def CleanupDiagnostic(cleanup):
        if (cleanup).is_CleanupFailed:
            return (cleanup).reason
        elif True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))

    @staticmethod
    def TerminationFailureDiagnostic(operationSucceeded, handleFound, cleanupComplete, diagnostic):
        if (diagnostic) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))):
            return diagnostic
        elif not(operationSucceeded):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "checked process-tree termination failed"))
        elif (handleFound) and (not(cleanupComplete)):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "checked process-tree termination was not confirmed"))
        elif True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "process handle disappeared before checked termination"))

    @staticmethod
    def TerminateProcessWithin(proc, deadlineMs):
        receipt: ProcessTerminationReceipt = ProcessTerminationReceipt.default()()
        d_0_operationSucceeded_: bool
        d_1_handleFound_: bool
        d_2_treeKillIssued_: bool
        d_3_cleanupComplete_: bool
        d_4_cleanupElapsedMs_: int
        d_5_diagnostic_: _dafny.Seq
        d_6_terminalChunk_: _dafny.Seq
        d_7_terminalAvailable_: bool
        d_8_combinedOutputBytesDelta_: int
        d_9_exitObserved_: bool
        d_10_exitCode_: int
        out0_: bool
        out1_: bool
        out2_: bool
        out3_: bool
        out4_: int
        out5_: _dafny.Seq
        out6_: _dafny.Seq
        out7_: bool
        out8_: int
        out9_: bool
        out10_: int
        out0_, out1_, out2_, out3_, out4_, out5_, out6_, out7_, out8_, out9_, out10_ = ConfluxIoCapability.default__.TerminateProcessTreeWithin(proc, deadlineMs)
        d_0_operationSucceeded_ = out0_
        d_1_handleFound_ = out1_
        d_2_treeKillIssued_ = out2_
        d_3_cleanupComplete_ = out3_
        d_4_cleanupElapsedMs_ = out4_
        d_5_diagnostic_ = out5_
        d_6_terminalChunk_ = out6_
        d_7_terminalAvailable_ = out7_
        d_8_combinedOutputBytesDelta_ = out8_
        d_9_exitObserved_ = out9_
        d_10_exitCode_ = out10_
        d_11_cleanup_: ConfluxSupervisedOperationResult.CleanupStatus
        if (d_0_operationSucceeded_) and (not(d_1_handleFound_)):
            d_11_cleanup_ = ConfluxSupervisedOperationResult.CleanupStatus_CleanupNotRequired()
        elif (d_0_operationSucceeded_) and (d_3_cleanupComplete_):
            d_11_cleanup_ = ConfluxSupervisedOperationResult.CleanupStatus_CleanupComplete(default__.NonNegative(d_4_cleanupElapsedMs_), deadlineMs)
        elif True:
            d_11_cleanup_ = ConfluxSupervisedOperationResult.CleanupStatus_CleanupFailed(default__.NonNegative(d_4_cleanupElapsedMs_), deadlineMs, default__.TerminationFailureDiagnostic(d_0_operationSucceeded_, d_1_handleFound_, d_3_cleanupComplete_, d_5_diagnostic_))
        d_12_exit_: ConfluxSupervisedOperationResult.ExitStatus
        if d_9_exitObserved_:
            d_12_exit_ = ConfluxSupervisedOperationResult.ExitStatus_ExitObserved(d_10_exitCode_)
        elif True:
            d_12_exit_ = ConfluxSupervisedOperationResult.ExitStatus_ExitNotObserved()
        receipt = ProcessTerminationReceipt_ProcessTerminationReceipt(d_1_handleFound_, d_2_treeKillIssued_, d_11_cleanup_, d_6_terminalChunk_, d_7_terminalAvailable_, default__.NonNegative(d_8_combinedOutputBytesDelta_), d_12_exit_)
        return receipt

    @staticmethod
    def NonNegative(value):
        if (value) < (0):
            return 0
        elif True:
            return value

    @staticmethod
    def FlatByteConcat(left, right):
        return _dafny.SeqWithoutIsStrInference([((left)[d_0_i_] if (d_0_i_) < (len(left)) else (right)[(d_0_i_) - (len(left))]) for d_0_i_ in range((len(left)) + (len(right)))])

    @staticmethod
    def BytesToMegabytes(bytes_):
        return ConfluxResourceCeilings.default__.BytesToMegabytes(bytes_)

    @staticmethod
    def DimensionFromNative(value):
        if (value) == (0):
            return ConfluxResourceCeilings.ResourceDimension_TimeDimension()
        elif (value) == (1):
            return ConfluxResourceCeilings.ResourceDimension_MemoryDimension()
        elif True:
            return ConfluxResourceCeilings.ResourceDimension_OutputDimension()

    @staticmethod
    def WrapperExitCode(childCode, verdict):
        if ((verdict).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementOk()):
            return childCode
        elif ((verdict).dimension).is_None:
            if (childCode) == (0):
                return 1
            elif True:
                return childCode
        elif (((verdict).dimension).value) == (ConfluxResourceCeilings.ResourceDimension_TimeDimension()):
            return 124
        elif (((verdict).dimension).value) == (ConfluxResourceCeilings.ResourceDimension_MemoryDimension()):
            return 137
        elif True:
            return 125

    @staticmethod
    def ObservationStatus(observation):
        if (((((observation).peakDescendantRssMb).status).is_MeasurementComplete) and ((((observation).combinedOutputBytes).status).is_MeasurementComplete)) and (((observation).processTree).is_ProcessTreeObserved):
            return ConfluxResourceCeilings.MeasurementStatus_MeasurementComplete()
        elif True:
            return ConfluxResourceCeilings.MeasurementStatus_MeasurementPartial(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "process observation contains incomplete measurements")))

    @staticmethod
    def StartObservation():
        startedMs: int = int(0)
        out0_: int
        out0_ = ConfluxIoCapability.default__.MonotonicTimeMs()
        startedMs = out0_
        return startedMs

    @staticmethod
    def ObserveCurrentProcess(startedMs, outputBytes, phase, budget):
        observation: ConfluxResourceCeilings.ResourceObservation = ConfluxResourceCeilings.ResourceObservation.default()()
        d_0_now_: int
        out0_: int
        out0_ = ConfluxIoCapability.default__.MonotonicTimeMs()
        d_0_now_ = out0_
        d_1_memoryPresent_: bool
        d_2_memoryMb_: int
        out1_: bool
        out2_: int
        out1_, out2_ = ConfluxIoCapability.default__.CurrentProcessRssMb()
        d_1_memoryPresent_ = out1_
        d_2_memoryMb_ = out2_
        d_3_memory_: ConfluxResourceCeilings.ResourceMeasurement
        if d_1_memoryPresent_:
            d_3_memory_ = ConfluxResourceCeilings.default__.CompleteMeasurement(default__.NonNegative(d_2_memoryMb_))
        elif ConfluxResourceCeilings.default__.MemoryMeasurementRequired(budget, phase):
            d_3_memory_ = ConfluxResourceCeilings.default__.UnavailableMeasurement(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "current-process RSS unavailable")))
        elif True:
            d_3_memory_ = ConfluxResourceCeilings.default__.NotRequestedMeasurement()
        d_4_output_: ConfluxResourceCeilings.ResourceMeasurement
        if (outputBytes).is_Some:
            d_4_output_ = ConfluxResourceCeilings.default__.CompleteMeasurement((outputBytes).value)
        elif ConfluxResourceCeilings.default__.OutputMeasurementRequired(budget):
            d_4_output_ = ConfluxResourceCeilings.default__.UnavailableMeasurement(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "combined-output measurement unavailable")))
        elif True:
            d_4_output_ = ConfluxResourceCeilings.default__.NotRequestedMeasurement()
        observation = ConfluxResourceCeilings.ResourceObservation_ResourceObservation(default__.NonNegative((d_0_now_) - (startedMs)), d_3_memory_, d_4_output_)
        return observation

    @staticmethod
    def ObserveCurrentProcessPostHoc(startedMs, outputBytes, phase, budget):
        observation: ConfluxResourceCeilings.ResourceObservation = ConfluxResourceCeilings.ResourceObservation.default()()
        verdict: ConfluxResourceCeilings.ResourceVerdict = ConfluxResourceCeilings.ResourceVerdict.default()()
        out0_: ConfluxResourceCeilings.ResourceObservation
        out0_ = default__.ObserveCurrentProcess(startedMs, outputBytes, phase, budget)
        observation = out0_
        verdict = ConfluxResourceCeilings.default__.DecidePostHoc(budget, phase, observation)
        return observation, verdict

    @staticmethod
    def ObserveCurrentProcessPostHocV1(policyId, startedMs, outputBytes, phase, budget):
        result: ConfluxSupervisedOperationResult.SupervisedOperationResultV1 = ConfluxSupervisedOperationResult.SupervisedOperationResultV1.default()()
        d_0_resolvedPolicyId_: _dafny.Seq
        if (len(policyId)) > (0):
            d_0_resolvedPolicyId_ = policyId
        elif True:
            d_0_resolvedPolicyId_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "policy:conflux/default@1"))
        d_1_observation_: ConfluxResourceCeilings.ResourceObservation
        d_2_verdict_: ConfluxResourceCeilings.ResourceVerdict
        out0_: ConfluxResourceCeilings.ResourceObservation
        out1_: ConfluxResourceCeilings.ResourceVerdict
        out0_, out1_ = default__.ObserveCurrentProcessPostHoc(startedMs, outputBytes, phase, budget)
        d_1_observation_ = out0_
        d_2_verdict_ = out1_
        d_3_tree_: ConfluxSupervisedOperationResult.ProcessTreeStatus
        if (((d_1_observation_).peakRssMb).status).is_MeasurementComplete:
            d_3_tree_ = ConfluxSupervisedOperationResult.ProcessTreeStatus_ProcessTreeObserved()
        elif True:
            d_3_tree_ = ConfluxSupervisedOperationResult.ProcessTreeStatus_ProcessTreeUnavailable(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "current-process RSS incomplete")))
        d_4_supervisionObservation_: ConfluxSupervisedOperationResult.SupervisionObservationV1
        d_4_supervisionObservation_ = ConfluxSupervisedOperationResult.SupervisionObservationV1_SupervisionObservationV1(ConfluxSupervisedOperationResult.TimingCoverage_TimingCoverage(0, (d_1_observation_).wallMs, 0, (d_1_observation_).wallMs), (d_1_observation_).peakRssMb, (d_1_observation_).combinedOutputBytes, d_3_tree_)
        d_5_state_: ConfluxSupervisedOperationResult.SupervisionState
        if ((d_2_verdict_).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementOk()):
            d_5_state_ = ConfluxSupervisedOperationResult.SupervisionState_SupervisionRunning(1)
        elif True:
            d_5_state_ = ConfluxSupervisedOperationResult.SupervisionState_SupervisionTripped(ConfluxSupervisedOperationResult.FirstBreachV1_FirstBreachV1(0, d_2_verdict_, ((d_2_verdict_).observed).is_None))
        d_6_wrapperExit_: int
        d_6_wrapperExit_ = default__.WrapperExitCode(0, d_2_verdict_)
        result = ConfluxSupervisedOperationResult.SupervisedOperationResultV1_SupervisedOperationResultV1(ConfluxSupervisedOperationResult.default__.SupervisedOperationResultSchemaVersion, d_0_resolvedPolicyId_, phase, ConfluxResourceCeilings.default__.EffectivePhase(phase), ConfluxSupervisedOperationResult.default__.ExpandBudget(budget, phase), default__.ObservationStatus(d_4_supervisionObservation_), d_4_supervisionObservation_, d_5_state_, ConfluxSupervisedOperationResult.CleanupStatus_CleanupNotRequired(), ConfluxSupervisedOperationResult.ExitStatus_ExitNotObserved(), (ConfluxSupervisedOperationResult.EnforcementStatus_PostHocFailure() if ((d_2_verdict_).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementFail()) else ConfluxSupervisedOperationResult.EnforcementStatus_NoEnforcement()), ConfluxSupervisedOperationResult.TerminationScope_NoTermination(), d_6_wrapperExit_, (d_2_verdict_).reason, d_2_verdict_)
        return result

    @staticmethod
    def RunProcess(cwd, command, arguments, phase, budget):
        result: SupervisedProcessResult = SupervisedProcessResult.default()()
        out0_: SupervisedProcessResult
        out0_ = default__.RunProcessV1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "policy:conflux/default@1")), cwd, command, arguments, phase, budget)
        result = out0_
        return result

    @staticmethod
    def RunProcessV1(policyId, cwd, command, arguments, phase, budget):
        result: SupervisedProcessResult = SupervisedProcessResult.default()()
        d_0_resolvedPolicyId_: _dafny.Seq
        if (len(policyId)) > (0):
            d_0_resolvedPolicyId_ = policyId
        elif True:
            d_0_resolvedPolicyId_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "policy:conflux/default@1"))
        d_1_timeLimit_: int
        d_1_timeLimit_ = ConfluxResourceCeilings.default__.TimeCeilingForPhase((budget).time, phase)
        d_2_memoryLimit_: int
        d_2_memoryLimit_ = ConfluxResourceCeilings.default__.MemoryCeilingForPhase((budget).memory, phase)
        d_3_code_: int
        d_4_output_: _dafny.Seq
        d_5_error_: _dafny.Seq
        d_6_timedOut_: bool
        d_7_memoryExceeded_: bool
        d_8_outputExceeded_: bool
        d_9_wallMs_: int
        d_10_memoryPresent_: bool
        d_11_peakMemoryMb_: int
        d_12_combinedOutputBytes_: int
        d_13_memoryMeasurementFailed_: bool
        d_14_outputMeasurementFailed_: bool
        d_15_treeKillIssued_: bool
        d_16_cleanupComplete_: bool
        d_17_cleanupElapsedMs_: int
        d_18_cleanupDeadlineMs_: int
        d_19_cleanupDiagnostic_: _dafny.Seq
        d_20_firstTripPresent_: bool
        d_21_firstTripIndex_: int
        d_22_firstTripDimension_: int
        d_23_firstTripObservedPresent_: bool
        d_24_firstTripObserved_: int
        d_25_firstTripMeasurementFailure_: bool
        d_26_firstTripInFlight_: bool
        d_27_samplesSeen_: int
        out0_: int
        out1_: _dafny.Seq
        out2_: _dafny.Seq
        out3_: bool
        out4_: bool
        out5_: bool
        out6_: int
        out7_: bool
        out8_: int
        out9_: int
        out10_: bool
        out11_: bool
        out12_: bool
        out13_: bool
        out14_: int
        out15_: int
        out16_: _dafny.Seq
        out17_: bool
        out18_: int
        out19_: int
        out20_: bool
        out21_: int
        out22_: bool
        out23_: bool
        out24_: int
        out0_, out1_, out2_, out3_, out4_, out5_, out6_, out7_, out8_, out9_, out10_, out11_, out12_, out13_, out14_, out15_, out16_, out17_, out18_, out19_, out20_, out21_, out22_, out23_, out24_ = ConfluxIoCapability.default__.RunProcessBounded(cwd, command, arguments, d_1_timeLimit_, d_2_memoryLimit_, (budget).maxOutputMb)
        d_3_code_ = out0_
        d_4_output_ = out1_
        d_5_error_ = out2_
        d_6_timedOut_ = out3_
        d_7_memoryExceeded_ = out4_
        d_8_outputExceeded_ = out5_
        d_9_wallMs_ = out6_
        d_10_memoryPresent_ = out7_
        d_11_peakMemoryMb_ = out8_
        d_12_combinedOutputBytes_ = out9_
        d_13_memoryMeasurementFailed_ = out10_
        d_14_outputMeasurementFailed_ = out11_
        d_15_treeKillIssued_ = out12_
        d_16_cleanupComplete_ = out13_
        d_17_cleanupElapsedMs_ = out14_
        d_18_cleanupDeadlineMs_ = out15_
        d_19_cleanupDiagnostic_ = out16_
        d_20_firstTripPresent_ = out17_
        d_21_firstTripIndex_ = out18_
        d_22_firstTripDimension_ = out19_
        d_23_firstTripObservedPresent_ = out20_
        d_24_firstTripObserved_ = out21_
        d_25_firstTripMeasurementFailure_ = out22_
        d_26_firstTripInFlight_ = out23_
        d_27_samplesSeen_ = out24_
        d_28_memoryMeasurement_: ConfluxResourceCeilings.ResourceMeasurement
        if d_10_memoryPresent_:
            if d_13_memoryMeasurementFailed_:
                d_28_memoryMeasurement_ = ConfluxResourceCeilings.ResourceMeasurement_ResourceMeasurement(ConfluxContract.Option_Some(default__.NonNegative(d_11_peakMemoryMb_)), ConfluxResourceCeilings.MeasurementStatus_MeasurementPartial(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "one or more process-tree RSS samples failed"))))
            elif True:
                d_28_memoryMeasurement_ = ConfluxResourceCeilings.default__.CompleteMeasurement(default__.NonNegative(d_11_peakMemoryMb_))
        elif (d_2_memoryLimit_) > (0):
            d_28_memoryMeasurement_ = ConfluxResourceCeilings.default__.UnavailableMeasurement(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "complete process-tree RSS unavailable")))
        elif True:
            d_28_memoryMeasurement_ = ConfluxResourceCeilings.default__.NotRequestedMeasurement()
        d_29_outputMeasurement_: ConfluxResourceCeilings.ResourceMeasurement
        if d_14_outputMeasurementFailed_:
            if (d_12_combinedOutputBytes_) >= (0):
                d_29_outputMeasurement_ = ConfluxResourceCeilings.ResourceMeasurement_ResourceMeasurement(ConfluxContract.Option_Some(default__.NonNegative(d_12_combinedOutputBytes_)), ConfluxResourceCeilings.MeasurementStatus_MeasurementPartial(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "combined-output stream observation incomplete"))))
            elif True:
                d_29_outputMeasurement_ = ConfluxResourceCeilings.default__.UnavailableMeasurement(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "combined-output stream observation unavailable")))
        elif True:
            d_29_outputMeasurement_ = ConfluxResourceCeilings.default__.CompleteMeasurement(default__.NonNegative(d_12_combinedOutputBytes_))
        d_30_observation_: ConfluxResourceCeilings.ResourceObservation
        d_30_observation_ = ConfluxResourceCeilings.ResourceObservation_ResourceObservation(default__.NonNegative(d_9_wallMs_), d_28_memoryMeasurement_, d_29_outputMeasurement_)
        d_31_nativeDimension_: ConfluxResourceCeilings.ResourceDimension
        d_31_nativeDimension_ = default__.DimensionFromNative(d_22_firstTripDimension_)
        d_32_nativeObservedCandidate_: ConfluxContract.Option
        if d_23_firstTripObservedPresent_:
            d_32_nativeObservedCandidate_ = ConfluxContract.Option_Some(default__.NonNegative(d_24_firstTripObserved_))
        elif True:
            d_32_nativeObservedCandidate_ = ConfluxResourceCeilings.default__.ObservedForDimension(d_30_observation_, d_31_nativeDimension_)
        d_33_nativeMeasurementFailure_: bool
        d_33_nativeMeasurementFailure_ = ((d_31_nativeDimension_) != (ConfluxResourceCeilings.ResourceDimension_TimeDimension())) and ((d_25_firstTripMeasurementFailure_) or ((d_32_nativeObservedCandidate_).is_None))
        d_34_nativeObserved_: ConfluxContract.Option
        if d_33_nativeMeasurementFailure_:
            d_34_nativeObserved_ = ConfluxContract.Option_None()
        elif True:
            d_34_nativeObserved_ = d_32_nativeObservedCandidate_
        d_35_nativeOutcome_: ConfluxResourceCeilings.EnforcementOutcome
        if d_26_firstTripInFlight_:
            d_35_nativeOutcome_ = ConfluxResourceCeilings.EnforcementOutcome_EnforcementKill()
        elif True:
            d_35_nativeOutcome_ = ConfluxResourceCeilings.EnforcementOutcome_EnforcementFail()
        d_36_nativeVerdict_: ConfluxResourceCeilings.ResourceVerdict
        d_36_nativeVerdict_ = ConfluxResourceCeilings.default__.RecordedViolation(d_35_nativeOutcome_, budget, phase, d_31_nativeDimension_, d_34_nativeObserved_, d_33_nativeMeasurementFailure_)
        d_37_finalVerdict_: ConfluxResourceCeilings.ResourceVerdict
        d_37_finalVerdict_ = ConfluxResourceCeilings.default__.DecidePostHoc(budget, phase, d_30_observation_)
        d_38_verdict_: ConfluxResourceCeilings.ResourceVerdict
        if d_20_firstTripPresent_:
            d_38_verdict_ = d_36_nativeVerdict_
        elif True:
            d_38_verdict_ = d_37_finalVerdict_
        d_39_state_: ConfluxSupervisedOperationResult.SupervisionState
        if d_20_firstTripPresent_:
            d_39_state_ = ConfluxSupervisedOperationResult.SupervisionState_SupervisionTripped(ConfluxSupervisedOperationResult.FirstBreachV1_FirstBreachV1(default__.NonNegative(d_21_firstTripIndex_), d_38_verdict_, d_33_nativeMeasurementFailure_))
        elif ((d_38_verdict_).outcome) != (ConfluxResourceCeilings.EnforcementOutcome_EnforcementOk()):
            d_39_state_ = ConfluxSupervisedOperationResult.SupervisionState_SupervisionTripped(ConfluxSupervisedOperationResult.FirstBreachV1_FirstBreachV1(default__.NonNegative(d_27_samplesSeen_), d_38_verdict_, ((d_38_verdict_).observed).is_None))
        elif True:
            d_39_state_ = ConfluxSupervisedOperationResult.SupervisionState_SupervisionRunning(default__.NonNegative(d_27_samplesSeen_))
        d_40_processTree_: ConfluxSupervisedOperationResult.ProcessTreeStatus
        if ((d_28_memoryMeasurement_).status).is_MeasurementComplete:
            d_40_processTree_ = ConfluxSupervisedOperationResult.ProcessTreeStatus_ProcessTreeObserved()
        elif True:
            d_40_processTree_ = ConfluxSupervisedOperationResult.ProcessTreeStatus_ProcessTreeUnavailable(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "complete process-tree RSS observation incomplete")))
        d_41_supervisionObservation_: ConfluxSupervisedOperationResult.SupervisionObservationV1
        d_41_supervisionObservation_ = ConfluxSupervisedOperationResult.SupervisionObservationV1_SupervisionObservationV1(ConfluxSupervisedOperationResult.TimingCoverage_TimingCoverage(0, (d_30_observation_).wallMs, 0, (d_30_observation_).wallMs), d_28_memoryMeasurement_, d_29_outputMeasurement_, d_40_processTree_)
        d_42_cleanup_: ConfluxSupervisedOperationResult.CleanupStatus
        if ((d_38_verdict_).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementKill()):
            if d_16_cleanupComplete_:
                d_42_cleanup_ = ConfluxSupervisedOperationResult.CleanupStatus_CleanupComplete(default__.NonNegative(d_17_cleanupElapsedMs_), default__.NonNegative(d_18_cleanupDeadlineMs_))
            elif True:
                d_42_cleanup_ = ConfluxSupervisedOperationResult.CleanupStatus_CleanupFailed(default__.NonNegative(d_17_cleanupElapsedMs_), default__.NonNegative(d_18_cleanupDeadlineMs_), (d_19_cleanupDiagnostic_ if (len(d_19_cleanupDiagnostic_)) > (0) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "process-tree cleanup exceeded monotonic deadline"))))
        elif True:
            d_42_cleanup_ = ConfluxSupervisedOperationResult.CleanupStatus_CleanupNotRequired()
        d_43_wrapperExit_: int
        d_43_wrapperExit_ = default__.WrapperExitCode(d_3_code_, d_38_verdict_)
        d_44_supervised_: ConfluxSupervisedOperationResult.SupervisedOperationResultV1
        d_44_supervised_ = ConfluxSupervisedOperationResult.SupervisedOperationResultV1_SupervisedOperationResultV1(ConfluxSupervisedOperationResult.default__.SupervisedOperationResultSchemaVersion, d_0_resolvedPolicyId_, phase, ConfluxResourceCeilings.default__.EffectivePhase(phase), ConfluxSupervisedOperationResult.default__.ExpandBudget(budget, phase), default__.ObservationStatus(d_41_supervisionObservation_), d_41_supervisionObservation_, d_39_state_, d_42_cleanup_, ConfluxSupervisedOperationResult.ExitStatus_ExitObserved(d_3_code_), (ConfluxSupervisedOperationResult.EnforcementStatus_InFlightProcessTreeTermination() if ((d_38_verdict_).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementKill()) else (ConfluxSupervisedOperationResult.EnforcementStatus_PostHocFailure() if ((d_38_verdict_).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementFail()) else ConfluxSupervisedOperationResult.EnforcementStatus_NoEnforcement())), (ConfluxSupervisedOperationResult.TerminationScope_ProcessTreeTermination() if ((d_38_verdict_).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementKill()) else ConfluxSupervisedOperationResult.TerminationScope_NoTermination()), d_43_wrapperExit_, (d_38_verdict_).reason, d_38_verdict_)
        result = SupervisedProcessResult_SupervisedProcessResult(d_3_code_, d_4_output_, d_5_error_, d_30_observation_, d_38_verdict_, default__.NonNegative(d_12_combinedOutputBytes_), d_15_treeKillIssued_, d_16_cleanupComplete_, d_44_supervised_)
        return result

    @staticmethod
    def PollProcess(proc, startedMs, outputBytesBefore, waitMs, phase, budget):
        chunk: _dafny.Seq = _dafny.Seq({})
        available: bool = False
        state: ProcessReadState = ProcessReadState.default()()
        observation: ConfluxResourceCeilings.ResourceObservation = ConfluxResourceCeilings.ResourceObservation.default()()
        outputBytesAfter: int = int(0)
        d_0_open_: bool = False
        d_1_combinedOutputBytesDelta_: int = int(0)
        out0_: _dafny.Seq
        out1_: bool
        out2_: bool
        out3_: int
        out0_, out1_, out2_, out3_ = ConfluxIoCapability.default__.ReadProcessOutputWithin(proc, waitMs)
        chunk = out0_
        available = out1_
        d_0_open_ = out2_
        d_1_combinedOutputBytesDelta_ = out3_
        d_2_now_: int
        out4_: int
        out4_ = ConfluxIoCapability.default__.MonotonicTimeMs()
        d_2_now_ = out4_
        d_3_handleFound_: bool
        d_4_exitObserved_: bool
        d_5_exitCode_: int
        out5_: bool
        out6_: bool
        out7_: int
        out5_, out6_, out7_ = ConfluxIoCapability.default__.ProcessExitStatus(proc)
        d_3_handleFound_ = out5_
        d_4_exitObserved_ = out6_
        d_5_exitCode_ = out7_
        if d_4_exitObserved_:
            d_6_finalChunk_: _dafny.Seq
            d_7_finalAvailable_: bool
            d_8_finalOpen_: bool
            d_9_finalOutputBytesDelta_: int
            out8_: _dafny.Seq
            out9_: bool
            out10_: bool
            out11_: int
            out8_, out9_, out10_, out11_ = ConfluxIoCapability.default__.ReadProcessOutputWithin(proc, 0)
            d_6_finalChunk_ = out8_
            d_7_finalAvailable_ = out9_
            d_8_finalOpen_ = out10_
            d_9_finalOutputBytesDelta_ = out11_
            chunk = default__.FlatByteConcat(chunk, d_6_finalChunk_)
            available = (available) or (d_7_finalAvailable_)
            d_0_open_ = d_8_finalOpen_
            d_1_combinedOutputBytesDelta_ = (d_1_combinedOutputBytesDelta_) + (d_9_finalOutputBytesDelta_)
        outputBytesAfter = (outputBytesBefore) + (default__.NonNegative(d_1_combinedOutputBytesDelta_))
        d_10_memoryPresent_: bool
        d_10_memoryPresent_ = False
        d_11_memoryMb_: int
        d_11_memoryMb_ = 0
        if ConfluxResourceCeilings.default__.MemoryMeasurementRequired(budget, phase):
            out12_: bool
            out13_: int
            out12_, out13_ = ConfluxIoCapability.default__.ProcessRssMb(proc)
            d_10_memoryPresent_ = out12_
            d_11_memoryMb_ = out13_
        d_12_memory_: ConfluxResourceCeilings.ResourceMeasurement
        if d_10_memoryPresent_:
            d_12_memory_ = ConfluxResourceCeilings.default__.CompleteMeasurement(default__.NonNegative(d_11_memoryMb_))
        elif ConfluxResourceCeilings.default__.MemoryMeasurementRequired(budget, phase):
            d_12_memory_ = ConfluxResourceCeilings.default__.UnavailableMeasurement(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "persistent process-tree RSS unavailable")))
        elif True:
            d_12_memory_ = ConfluxResourceCeilings.default__.NotRequestedMeasurement()
        observation = ConfluxResourceCeilings.ResourceObservation_ResourceObservation(default__.NonNegative((d_2_now_) - (startedMs)), d_12_memory_, ConfluxResourceCeilings.default__.CompleteMeasurement(outputBytesAfter))
        d_13_terminal_: bool
        d_13_terminal_ = (d_4_exitObserved_) or ((not(d_0_open_)) and (not(d_3_handleFound_)))
        d_14_verdict_: ConfluxResourceCeilings.ResourceVerdict
        if d_13_terminal_:
            d_14_verdict_ = ConfluxResourceCeilings.default__.DecidePostHoc(budget, phase, observation)
        elif True:
            d_14_verdict_ = ConfluxResourceCeilings.default__.DecideInFlight(budget, phase, observation)
        if (d_13_terminal_) and (available):
            state = ProcessReadState_ProcessReadOpen()
        elif d_4_exitObserved_:
            state = ProcessReadState_ProcessReadExited(d_5_exitCode_, d_14_verdict_)
        elif (not(d_0_open_)) and (not(d_3_handleFound_)):
            state = ProcessReadState_ProcessReadClosed()
        elif ((d_14_verdict_).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementKill()):
            d_15_termination_: ProcessTerminationReceipt
            out14_: ProcessTerminationReceipt
            out14_ = default__.TerminateProcessWithin(proc, default__.ProcessCleanupDeadlineMs)
            d_15_termination_ = out14_
            chunk = default__.FlatByteConcat(chunk, (d_15_termination_).terminalChunk)
            available = (available) or ((d_15_termination_).terminalAvailable)
            outputBytesAfter = (outputBytesAfter) + ((d_15_termination_).combinedOutputBytesDelta)
            d_16_cleanupNow_: int
            out15_: int
            out15_ = ConfluxIoCapability.default__.MonotonicTimeMs()
            d_16_cleanupNow_ = out15_
            observation = ConfluxResourceCeilings.ResourceObservation_ResourceObservation(default__.NonNegative((d_16_cleanupNow_) - (startedMs)), d_12_memory_, ConfluxResourceCeilings.default__.CompleteMeasurement(outputBytesAfter))
            state = ProcessReadState_ProcessReadKilled(d_14_verdict_, d_15_termination_)
        elif (d_0_open_) or (d_3_handleFound_):
            state = ProcessReadState_ProcessReadOpen()
        elif True:
            state = ProcessReadState_ProcessReadClosed()
        return chunk, available, state, observation, outputBytesAfter

    @staticmethod
    def CancelProcess(proc, startedMs, outputBytesBefore, phase, budget):
        receipt: ProcessCancellationResult = ProcessCancellationResult.default()()
        d_0_chunk_: _dafny.Seq
        d_1_available_: bool
        d_2_open_: bool
        d_3_combinedOutputBytesDelta_: int
        out0_: _dafny.Seq
        out1_: bool
        out2_: bool
        out3_: int
        out0_, out1_, out2_, out3_ = ConfluxIoCapability.default__.ReadProcessOutputWithin(proc, 0)
        d_0_chunk_ = out0_
        d_1_available_ = out1_
        d_2_open_ = out2_
        d_3_combinedOutputBytesDelta_ = out3_
        d_4_outputBytesAfter_: int
        d_4_outputBytesAfter_ = (outputBytesBefore) + (default__.NonNegative(d_3_combinedOutputBytesDelta_))
        d_5_now_: int
        out4_: int
        out4_ = ConfluxIoCapability.default__.MonotonicTimeMs()
        d_5_now_ = out4_
        d_6_memoryPresent_: bool
        d_6_memoryPresent_ = False
        d_7_memoryMb_: int
        d_7_memoryMb_ = 0
        if ConfluxResourceCeilings.default__.MemoryMeasurementRequired(budget, phase):
            out5_: bool
            out6_: int
            out5_, out6_ = ConfluxIoCapability.default__.ProcessRssMb(proc)
            d_6_memoryPresent_ = out5_
            d_7_memoryMb_ = out6_
        d_8_memory_: ConfluxResourceCeilings.ResourceMeasurement
        if d_6_memoryPresent_:
            d_8_memory_ = ConfluxResourceCeilings.default__.CompleteMeasurement(default__.NonNegative(d_7_memoryMb_))
        elif ConfluxResourceCeilings.default__.MemoryMeasurementRequired(budget, phase):
            d_8_memory_ = ConfluxResourceCeilings.default__.UnavailableMeasurement(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "persistent process-tree RSS unavailable at cancellation")))
        elif True:
            d_8_memory_ = ConfluxResourceCeilings.default__.NotRequestedMeasurement()
        d_9_beforeObservation_: ConfluxResourceCeilings.ResourceObservation
        d_9_beforeObservation_ = ConfluxResourceCeilings.ResourceObservation_ResourceObservation(default__.NonNegative((d_5_now_) - (startedMs)), d_8_memory_, ConfluxResourceCeilings.default__.CompleteMeasurement(d_4_outputBytesAfter_))
        d_10_termination_: ProcessTerminationReceipt
        out7_: ProcessTerminationReceipt
        out7_ = default__.TerminateProcessWithin(proc, default__.ProcessCleanupDeadlineMs)
        d_10_termination_ = out7_
        d_0_chunk_ = default__.FlatByteConcat(d_0_chunk_, (d_10_termination_).terminalChunk)
        d_1_available_ = (d_1_available_) or ((d_10_termination_).terminalAvailable)
        d_4_outputBytesAfter_ = (d_4_outputBytesAfter_) + ((d_10_termination_).combinedOutputBytesDelta)
        d_11_cancellationNow_: int
        out8_: int
        out8_ = ConfluxIoCapability.default__.MonotonicTimeMs()
        d_11_cancellationNow_ = out8_
        d_12_finalObservation_: ConfluxResourceCeilings.ResourceObservation
        d_12_finalObservation_ = ConfluxResourceCeilings.ResourceObservation_ResourceObservation(default__.NonNegative((d_11_cancellationNow_) - (startedMs)), d_8_memory_, ConfluxResourceCeilings.default__.CompleteMeasurement(d_4_outputBytesAfter_))
        d_13_outcome_: ProcessCancellationOutcome
        if ((d_10_termination_).cleanup).is_CleanupFailed:
            d_13_outcome_ = ProcessCancellationOutcome_ProcessTerminationFailed()
        elif not((d_10_termination_).handleFound):
            d_13_outcome_ = ProcessCancellationOutcome_ProcessAlreadyClosed()
        elif True:
            d_13_outcome_ = ProcessCancellationOutcome_ProcessCancelled()
        receipt = ProcessCancellationResult_ProcessCancellationResult(d_13_outcome_, d_0_chunk_, d_1_available_, d_9_beforeObservation_, d_12_finalObservation_, d_4_outputBytesAfter_, d_10_termination_)
        return receipt

    @_dafny.classproperty
    def ProcessCleanupDeadlineMs(instance):
        return 2000

class ProcessReadState:
    @classmethod
    def default(cls, ):
        return lambda: ProcessReadState_ProcessReadOpen()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ProcessReadOpen(self) -> bool:
        return isinstance(self, ProcessReadState_ProcessReadOpen)
    @property
    def is_ProcessReadClosed(self) -> bool:
        return isinstance(self, ProcessReadState_ProcessReadClosed)
    @property
    def is_ProcessReadExited(self) -> bool:
        return isinstance(self, ProcessReadState_ProcessReadExited)
    @property
    def is_ProcessReadKilled(self) -> bool:
        return isinstance(self, ProcessReadState_ProcessReadKilled)

class ProcessReadState_ProcessReadOpen(ProcessReadState, NamedTuple('ProcessReadOpen', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceSupervisor.ProcessReadState.ProcessReadOpen'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProcessReadState_ProcessReadOpen)
    def __hash__(self) -> int:
        return super().__hash__()

class ProcessReadState_ProcessReadClosed(ProcessReadState, NamedTuple('ProcessReadClosed', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceSupervisor.ProcessReadState.ProcessReadClosed'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProcessReadState_ProcessReadClosed)
    def __hash__(self) -> int:
        return super().__hash__()

class ProcessReadState_ProcessReadExited(ProcessReadState, NamedTuple('ProcessReadExited', [('exitCode', Any), ('verdict', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceSupervisor.ProcessReadState.ProcessReadExited({_dafny.string_of(self.exitCode)}, {_dafny.string_of(self.verdict)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProcessReadState_ProcessReadExited) and self.exitCode == __o.exitCode and self.verdict == __o.verdict
    def __hash__(self) -> int:
        return super().__hash__()

class ProcessReadState_ProcessReadKilled(ProcessReadState, NamedTuple('ProcessReadKilled', [('verdict', Any), ('termination', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceSupervisor.ProcessReadState.ProcessReadKilled({_dafny.string_of(self.verdict)}, {_dafny.string_of(self.termination)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProcessReadState_ProcessReadKilled) and self.verdict == __o.verdict and self.termination == __o.termination
    def __hash__(self) -> int:
        return super().__hash__()


class ProcessTerminationReceipt:
    @classmethod
    def default(cls, ):
        return lambda: ProcessTerminationReceipt_ProcessTerminationReceipt(False, False, ConfluxSupervisedOperationResult.CleanupStatus.default()(), _dafny.Seq({}), False, int(0), ConfluxSupervisedOperationResult.ExitStatus.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ProcessTerminationReceipt(self) -> bool:
        return isinstance(self, ProcessTerminationReceipt_ProcessTerminationReceipt)

class ProcessTerminationReceipt_ProcessTerminationReceipt(ProcessTerminationReceipt, NamedTuple('ProcessTerminationReceipt', [('handleFound', Any), ('treeKillIssued', Any), ('cleanup', Any), ('terminalChunk', Any), ('terminalAvailable', Any), ('combinedOutputBytesDelta', Any), ('exit', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceSupervisor.ProcessTerminationReceipt.ProcessTerminationReceipt({_dafny.string_of(self.handleFound)}, {_dafny.string_of(self.treeKillIssued)}, {_dafny.string_of(self.cleanup)}, {_dafny.string_of(self.terminalChunk)}, {_dafny.string_of(self.terminalAvailable)}, {_dafny.string_of(self.combinedOutputBytesDelta)}, {_dafny.string_of(self.exit)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProcessTerminationReceipt_ProcessTerminationReceipt) and self.handleFound == __o.handleFound and self.treeKillIssued == __o.treeKillIssued and self.cleanup == __o.cleanup and self.terminalChunk == __o.terminalChunk and self.terminalAvailable == __o.terminalAvailable and self.combinedOutputBytesDelta == __o.combinedOutputBytesDelta and self.exit == __o.exit
    def __hash__(self) -> int:
        return super().__hash__()


class ProcessCancellationOutcome:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [ProcessCancellationOutcome_ProcessCancelled(), ProcessCancellationOutcome_ProcessAlreadyClosed(), ProcessCancellationOutcome_ProcessTerminationFailed()]
    @classmethod
    def default(cls, ):
        return lambda: ProcessCancellationOutcome_ProcessCancelled()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ProcessCancelled(self) -> bool:
        return isinstance(self, ProcessCancellationOutcome_ProcessCancelled)
    @property
    def is_ProcessAlreadyClosed(self) -> bool:
        return isinstance(self, ProcessCancellationOutcome_ProcessAlreadyClosed)
    @property
    def is_ProcessTerminationFailed(self) -> bool:
        return isinstance(self, ProcessCancellationOutcome_ProcessTerminationFailed)

class ProcessCancellationOutcome_ProcessCancelled(ProcessCancellationOutcome, NamedTuple('ProcessCancelled', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceSupervisor.ProcessCancellationOutcome.ProcessCancelled'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProcessCancellationOutcome_ProcessCancelled)
    def __hash__(self) -> int:
        return super().__hash__()

class ProcessCancellationOutcome_ProcessAlreadyClosed(ProcessCancellationOutcome, NamedTuple('ProcessAlreadyClosed', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceSupervisor.ProcessCancellationOutcome.ProcessAlreadyClosed'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProcessCancellationOutcome_ProcessAlreadyClosed)
    def __hash__(self) -> int:
        return super().__hash__()

class ProcessCancellationOutcome_ProcessTerminationFailed(ProcessCancellationOutcome, NamedTuple('ProcessTerminationFailed', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceSupervisor.ProcessCancellationOutcome.ProcessTerminationFailed'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProcessCancellationOutcome_ProcessTerminationFailed)
    def __hash__(self) -> int:
        return super().__hash__()


class ProcessCancellationResult:
    @classmethod
    def default(cls, ):
        return lambda: ProcessCancellationResult_ProcessCancellationResult(ProcessCancellationOutcome.default()(), _dafny.Seq({}), False, ConfluxResourceCeilings.ResourceObservation.default()(), ConfluxResourceCeilings.ResourceObservation.default()(), int(0), ProcessTerminationReceipt.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ProcessCancellationResult(self) -> bool:
        return isinstance(self, ProcessCancellationResult_ProcessCancellationResult)

class ProcessCancellationResult_ProcessCancellationResult(ProcessCancellationResult, NamedTuple('ProcessCancellationResult', [('outcome', Any), ('chunk', Any), ('available', Any), ('beforeObservation', Any), ('finalObservation', Any), ('combinedOutputBytes', Any), ('termination', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceSupervisor.ProcessCancellationResult.ProcessCancellationResult({_dafny.string_of(self.outcome)}, {_dafny.string_of(self.chunk)}, {_dafny.string_of(self.available)}, {_dafny.string_of(self.beforeObservation)}, {_dafny.string_of(self.finalObservation)}, {_dafny.string_of(self.combinedOutputBytes)}, {_dafny.string_of(self.termination)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProcessCancellationResult_ProcessCancellationResult) and self.outcome == __o.outcome and self.chunk == __o.chunk and self.available == __o.available and self.beforeObservation == __o.beforeObservation and self.finalObservation == __o.finalObservation and self.combinedOutputBytes == __o.combinedOutputBytes and self.termination == __o.termination
    def __hash__(self) -> int:
        return super().__hash__()


class SupervisedProcessResult:
    @classmethod
    def default(cls, ):
        return lambda: SupervisedProcessResult_SupervisedProcessResult(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxResourceCeilings.ResourceObservation.default()(), ConfluxResourceCeilings.ResourceVerdict.default()(), int(0), False, False, ConfluxSupervisedOperationResult.SupervisedOperationResultV1.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SupervisedProcessResult(self) -> bool:
        return isinstance(self, SupervisedProcessResult_SupervisedProcessResult)

class SupervisedProcessResult_SupervisedProcessResult(SupervisedProcessResult, NamedTuple('SupervisedProcessResult', [('code', Any), ('output', Any), ('error', Any), ('observation', Any), ('verdict', Any), ('combinedOutputBytes', Any), ('treeKillIssued', Any), ('cleanupComplete', Any), ('supervised', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceSupervisor.SupervisedProcessResult.SupervisedProcessResult({_dafny.string_of(self.code)}, {self.output.VerbatimString(True)}, {self.error.VerbatimString(True)}, {_dafny.string_of(self.observation)}, {_dafny.string_of(self.verdict)}, {_dafny.string_of(self.combinedOutputBytes)}, {_dafny.string_of(self.treeKillIssued)}, {_dafny.string_of(self.cleanupComplete)}, {_dafny.string_of(self.supervised)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SupervisedProcessResult_SupervisedProcessResult) and self.code == __o.code and self.output == __o.output and self.error == __o.error and self.observation == __o.observation and self.verdict == __o.verdict and self.combinedOutputBytes == __o.combinedOutputBytes and self.treeKillIssued == __o.treeKillIssued and self.cleanupComplete == __o.cleanupComplete and self.supervised == __o.supervised
    def __hash__(self) -> int:
        return super().__hash__()

