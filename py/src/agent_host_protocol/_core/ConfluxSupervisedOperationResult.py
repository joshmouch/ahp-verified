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

# Module: ConfluxSupervisedOperationResult

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def V1WireNat(value):
        return (value) <= (default__.V1MaxSafeInteger)

    @staticmethod
    def V1WireInt(value):
        return ((default__.V1MinSafeInteger) <= (value)) and ((value) <= (default__.V1MaxSafeInteger))

    @staticmethod
    def V1WireOptionalNat(value):
        return ((value).is_None) or (default__.V1WireNat((value).value))

    @staticmethod
    def V1BudgetSafeForPhase(budget, phase):
        return ((default__.V1WireNat(ConfluxResourceCeilings.default__.TimeCeilingForPhase((budget).time, phase))) and (default__.V1WireNat(ConfluxResourceCeilings.default__.MemoryCeilingForPhase((budget).memory, phase)))) and (default__.V1WireNat((budget).maxOutputMb))

    @staticmethod
    def InclusiveWallMs(timing):
        if ((timing).completedMs) >= ((timing).startedMs):
            return ((timing).completedMs) - ((timing).startedMs)
        elif True:
            return 0

    @staticmethod
    def ClampCoverage(inclusiveMs, unionMs):
        if (unionMs) <= (inclusiveMs):
            return unionMs
        elif True:
            return inclusiveMs

    @staticmethod
    def ExclusiveWallMs(timing):
        return (default__.InclusiveWallMs(timing)) - (default__.ClampCoverage(default__.InclusiveWallMs(timing), (timing).childUnionMs))

    @staticmethod
    def GapWallMs(timing):
        return (default__.InclusiveWallMs(timing)) - (default__.ClampCoverage(default__.InclusiveWallMs(timing), (timing).observedUnionMs))

    @staticmethod
    def ExpandBudget(budget, requestedPhase):
        return ExpandedResourceBudgetV1_ExpandedResourceBudgetV1(ConfluxResourceCeilings.default__.TimeCeilingForPhase((budget).time, requestedPhase), ConfluxResourceCeilings.default__.MemoryCeilingForPhase((budget).memory, requestedPhase), (budget).maxOutputMb)

    @staticmethod
    def StatusKind(status):
        if (status).is_MeasurementComplete:
            return 0
        elif (status).is_MeasurementPartial:
            return 1
        elif (status).is_MeasurementUnavailable:
            return 2
        elif True:
            return 3

    @staticmethod
    def DerivedObservationStatus(observation):
        if (((((observation).peakDescendantRssMb).status).is_MeasurementComplete) and ((((observation).combinedOutputBytes).status).is_MeasurementComplete)) and (((observation).processTree).is_ProcessTreeObserved):
            return 0
        elif True:
            return 1

    @staticmethod
    def VerdictShapeValid(verdict):
        return (not (((verdict).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementOk())) or (((((verdict).dimension).is_None) and (((verdict).ceiling).is_None)) and (((verdict).observed).is_None))) and (not (((verdict).outcome) != (ConfluxResourceCeilings.EnforcementOutcome_EnforcementOk())) or (((((verdict).dimension).is_Some) and (((verdict).ceiling).is_Some)) and ((len((verdict).reason)) > (0))))

    @staticmethod
    def FirstBreachCoherent(result):
        return (((((result).verdict).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementOk())) == (((result).supervisionState).is_SupervisionRunning)) and (not (((result).supervisionState).is_SupervisionTripped) or ((((((((result).supervisionState).record).verdict) == ((result).verdict)) and ((((((result).supervisionState).record).verdict).outcome) != (ConfluxResourceCeilings.EnforcementOutcome_EnforcementOk()))) and (not ((((result).supervisionState).record).measurementFailure) or ((((((result).verdict).dimension).is_Some) and ((((((result).verdict).dimension).value) == (ConfluxResourceCeilings.ResourceDimension_MemoryDimension())) or (((((result).verdict).dimension).value) == (ConfluxResourceCeilings.ResourceDimension_OutputDimension())))) and ((((result).verdict).observed).is_None)))) and (not (not((((result).supervisionState).record).measurementFailure)) or ((((result).verdict).observed).is_Some))))

    @staticmethod
    def RequiredMeasurementsFailClosed(result):
        return (not (((((result).expandedCeilings).peakRssMb) > (0)) and (not(((((result).observation).peakDescendantRssMb).status).is_MeasurementComplete))) or (((((result).verdict).outcome) != (ConfluxResourceCeilings.EnforcementOutcome_EnforcementOk())) and (((result).wrapperExitCode) != (0)))) and (not (((((result).expandedCeilings).combinedOutputMb) > (0)) and (not(((((result).observation).combinedOutputBytes).status).is_MeasurementComplete))) or (((((result).verdict).outcome) != (ConfluxResourceCeilings.EnforcementOutcome_EnforcementOk())) and (((result).wrapperExitCode) != (0))))

    @staticmethod
    def CleanupValid(cleanup):
        return (((cleanup).is_CleanupNotRequired) or ((((cleanup).is_CleanupComplete) and (((cleanup).deadlineMs) > (0))) and (((cleanup).elapsedMs) <= ((cleanup).deadlineMs)))) or (((((cleanup).is_CleanupFailed) and (((cleanup).deadlineMs) > (0))) and (((cleanup).elapsedMs) <= ((cleanup).deadlineMs))) and ((len((cleanup).reason)) > (0)))

    @staticmethod
    def MeasurementStatusValid(status):
        return ((not((status).is_MeasurementPartial)) and (not((status).is_MeasurementUnavailable))) or ((len((status).reason)) > (0))

    @staticmethod
    def CleanupCoherent(result):
        if ((result).cleanup).is_CleanupNotRequired:
            return ((result).termination) == (TerminationScope_NoTermination())
        elif True:
            return (((((result).termination) == (TerminationScope_ProcessTreeTermination())) and ((((result).cleanup).deadlineMs) > (0))) and ((((result).cleanup).elapsedMs) <= (((result).cleanup).deadlineMs))) and ((((result).cleanup).is_CleanupComplete) or ((((result).cleanup).is_CleanupFailed) and ((len(((result).cleanup).reason)) > (0))))

    @staticmethod
    def ExpandedLimit(result, dimension):
        if (dimension) == (ConfluxResourceCeilings.ResourceDimension_TimeDimension()):
            return ((result).expandedCeilings).timeMs
        elif (dimension) == (ConfluxResourceCeilings.ResourceDimension_MemoryDimension()):
            return ((result).expandedCeilings).peakRssMb
        elif True:
            return ((result).expandedCeilings).combinedOutputMb

    @staticmethod
    def NumericBreachReflected(result, dimension, observed):
        if (dimension) == (ConfluxResourceCeilings.ResourceDimension_TimeDimension()):
            return (observed) <= (default__.InclusiveWallMs(((result).observation).timing))
        elif (dimension) == (ConfluxResourceCeilings.ResourceDimension_MemoryDimension()):
            return (((((result).observation).peakDescendantRssMb).value).is_Some) and ((observed) <= (((((result).observation).peakDescendantRssMb).value).value))
        elif True:
            return (((((result).observation).combinedOutputBytes).value).is_Some) and ((observed) <= (ConfluxResourceCeilings.default__.BytesToMegabytes(((((result).observation).combinedOutputBytes).value).value)))

    @staticmethod
    def VerdictBudgetCoherent(result):
        if (((result).verdict).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementOk()):
            return ((((((result).expandedCeilings).timeMs) == (0)) or ((default__.InclusiveWallMs(((result).observation).timing)) <= (((result).expandedCeilings).timeMs))) and (((((result).expandedCeilings).peakRssMb) == (0)) or (((((((result).observation).peakDescendantRssMb).status).is_MeasurementComplete) and (((((result).observation).peakDescendantRssMb).value).is_Some)) and ((((((result).observation).peakDescendantRssMb).value).value) <= (((result).expandedCeilings).peakRssMb))))) and (((((result).expandedCeilings).combinedOutputMb) == (0)) or (((((((result).observation).combinedOutputBytes).status).is_MeasurementComplete) and (((((result).observation).combinedOutputBytes).value).is_Some)) and ((ConfluxResourceCeilings.default__.BytesToMegabytes(((((result).observation).combinedOutputBytes).value).value)) <= (((result).expandedCeilings).combinedOutputMb))))
        elif True:
            return (((((((result).supervisionState).is_SupervisionTripped) and ((((result).verdict).dimension).is_Some)) and ((((result).verdict).ceiling).is_Some)) and (((((result).verdict).ceiling).value) == (default__.ExpandedLimit(result, (((result).verdict).dimension).value)))) and (((((result).verdict).ceiling).value) > (0))) and ((((((result).verdict).observed).is_None) and ((not(((((result).observation).peakDescendantRssMb).status).is_MeasurementComplete) if ((((result).verdict).dimension).value) == (ConfluxResourceCeilings.ResourceDimension_MemoryDimension()) else (((((result).verdict).dimension).value) == (ConfluxResourceCeilings.ResourceDimension_OutputDimension())) and (not(((((result).observation).combinedOutputBytes).status).is_MeasurementComplete)))) if (((result).supervisionState).record).measurementFailure else (((((result).verdict).observed).is_Some) and (((((result).verdict).observed).value) > ((((result).verdict).ceiling).value))) and (default__.NumericBreachReflected(result, (((result).verdict).dimension).value, (((result).verdict).observed).value))))

    @staticmethod
    def V1WireSafe(result):
        return (((((((((((((((default__.V1WireNat((result).schemaVersion)) and (default__.V1WireNat(((result).expandedCeilings).timeMs))) and (default__.V1WireNat(((result).expandedCeilings).peakRssMb))) and (default__.V1WireNat(((result).expandedCeilings).combinedOutputMb))) and (default__.V1WireNat((((result).observation).timing).startedMs))) and (default__.V1WireNat((((result).observation).timing).completedMs))) and (default__.V1WireNat((((result).observation).timing).childUnionMs))) and (default__.V1WireNat((((result).observation).timing).observedUnionMs))) and (default__.V1WireOptionalNat((((result).observation).peakDescendantRssMb).value))) and (default__.V1WireOptionalNat((((result).observation).combinedOutputBytes).value))) and (default__.V1WireOptionalNat(((result).verdict).ceiling))) and (default__.V1WireOptionalNat(((result).verdict).observed))) and ((default__.V1WireNat(((result).supervisionState).samplesSeen) if ((result).supervisionState).is_SupervisionRunning else ((((result).supervisionState).record).sampleIndex) < (default__.V1MaxSafeInteger)))) and ((True if ((result).cleanup).is_CleanupNotRequired else (default__.V1WireNat(((result).cleanup).elapsedMs)) and (default__.V1WireNat(((result).cleanup).deadlineMs))))) and ((((result).exit).is_ExitNotObserved) or (default__.V1WireInt(((result).exit).code)))) and (default__.V1WireInt((result).wrapperExitCode))

    @staticmethod
    def ValidV1(result):
        return (((((((((((((((((((((((result).schemaVersion) == (default__.SupervisedOperationResultSchemaVersion)) and (default__.V1WireSafe(result))) and ((len((result).policyId)) > (0))) and (((result).effectivePhase) == (ConfluxResourceCeilings.default__.EffectivePhase((result).requestedPhase)))) and (((((result).observation).timing).completedMs) >= ((((result).observation).timing).startedMs))) and (ConfluxResourceCeilings.default__.ValidMeasurement(((result).observation).peakDescendantRssMb))) and (ConfluxResourceCeilings.default__.ValidMeasurement(((result).observation).combinedOutputBytes))) and (default__.MeasurementStatusValid((result).observationStatus))) and ((default__.StatusKind((result).observationStatus)) == (default__.DerivedObservationStatus((result).observation)))) and (not ((((result).observation).processTree).is_ProcessTreeObserved) or (((((result).observation).peakDescendantRssMb).status).is_MeasurementComplete))) and (not ((((result).observation).processTree).is_ProcessTreeUnavailable) or ((not(((((result).observation).peakDescendantRssMb).status).is_MeasurementComplete)) and ((len((((result).observation).processTree).reason)) > (0))))) and ((((result).verdict).phase) == ((result).effectivePhase))) and (default__.VerdictShapeValid((result).verdict))) and (((result).reason) == (((result).verdict).reason))) and (default__.FirstBreachCoherent(result))) and (default__.RequiredMeasurementsFailClosed(result))) and (default__.VerdictBudgetCoherent(result))) and (default__.CleanupValid((result).cleanup))) and (not ((((result).verdict).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementOk())) or (((((result).enforcement) == (EnforcementStatus_NoEnforcement())) and (((result).termination) == (TerminationScope_NoTermination()))) and (not (((result).exit).is_ExitObserved) or (((result).wrapperExitCode) == (((result).exit).code)))))) and (not ((((result).verdict).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementFail())) or (((((result).enforcement) == (EnforcementStatus_PostHocFailure())) and (((result).termination) == (TerminationScope_NoTermination()))) and (((result).wrapperExitCode) != (0))))) and (not ((((result).verdict).outcome) == (ConfluxResourceCeilings.EnforcementOutcome_EnforcementKill())) or (((((result).enforcement) == (EnforcementStatus_InFlightProcessTreeTermination())) and (((result).termination) == (TerminationScope_ProcessTreeTermination()))) and (((result).wrapperExitCode) != (0))))) and (default__.CleanupCoherent(result))

    @staticmethod
    def V1ReaderAccepts(schemaVersion, requiredV1FieldsPresent, requiredV1MeaningsUnchanged):
        return (((schemaVersion) == (default__.SupervisedOperationResultSchemaVersion)) and (requiredV1FieldsPresent)) and (requiredV1MeaningsUnchanged)

    @staticmethod
    def PreserveProviderOutcome(provider, consumerCandidate):
        return ConfluxResourceCeilings.default__.DominantOutcome(provider, consumerCandidate)

    @_dafny.classproperty
    def V1MaxSafeInteger(instance):
        return 9007199254740991
    @_dafny.classproperty
    def V1MinSafeInteger(instance):
        return -9007199254740991
    @_dafny.classproperty
    def SupervisedOperationResultSchemaVersion(instance):
        return 1

class TimingCoverage:
    @classmethod
    def default(cls, ):
        return lambda: TimingCoverage_TimingCoverage(int(0), int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_TimingCoverage(self) -> bool:
        return isinstance(self, TimingCoverage_TimingCoverage)

class TimingCoverage_TimingCoverage(TimingCoverage, NamedTuple('TimingCoverage', [('startedMs', Any), ('completedMs', Any), ('childUnionMs', Any), ('observedUnionMs', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.TimingCoverage.TimingCoverage({_dafny.string_of(self.startedMs)}, {_dafny.string_of(self.completedMs)}, {_dafny.string_of(self.childUnionMs)}, {_dafny.string_of(self.observedUnionMs)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TimingCoverage_TimingCoverage) and self.startedMs == __o.startedMs and self.completedMs == __o.completedMs and self.childUnionMs == __o.childUnionMs and self.observedUnionMs == __o.observedUnionMs
    def __hash__(self) -> int:
        return super().__hash__()


class ExpandedResourceBudgetV1:
    @classmethod
    def default(cls, ):
        return lambda: ExpandedResourceBudgetV1_ExpandedResourceBudgetV1(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ExpandedResourceBudgetV1(self) -> bool:
        return isinstance(self, ExpandedResourceBudgetV1_ExpandedResourceBudgetV1)

class ExpandedResourceBudgetV1_ExpandedResourceBudgetV1(ExpandedResourceBudgetV1, NamedTuple('ExpandedResourceBudgetV1', [('timeMs', Any), ('peakRssMb', Any), ('combinedOutputMb', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.ExpandedResourceBudgetV1.ExpandedResourceBudgetV1({_dafny.string_of(self.timeMs)}, {_dafny.string_of(self.peakRssMb)}, {_dafny.string_of(self.combinedOutputMb)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ExpandedResourceBudgetV1_ExpandedResourceBudgetV1) and self.timeMs == __o.timeMs and self.peakRssMb == __o.peakRssMb and self.combinedOutputMb == __o.combinedOutputMb
    def __hash__(self) -> int:
        return super().__hash__()


class ProcessTreeStatus:
    @classmethod
    def default(cls, ):
        return lambda: ProcessTreeStatus_ProcessTreeObserved()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ProcessTreeObserved(self) -> bool:
        return isinstance(self, ProcessTreeStatus_ProcessTreeObserved)
    @property
    def is_ProcessTreeUnavailable(self) -> bool:
        return isinstance(self, ProcessTreeStatus_ProcessTreeUnavailable)

class ProcessTreeStatus_ProcessTreeObserved(ProcessTreeStatus, NamedTuple('ProcessTreeObserved', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.ProcessTreeStatus.ProcessTreeObserved'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProcessTreeStatus_ProcessTreeObserved)
    def __hash__(self) -> int:
        return super().__hash__()

class ProcessTreeStatus_ProcessTreeUnavailable(ProcessTreeStatus, NamedTuple('ProcessTreeUnavailable', [('reason', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.ProcessTreeStatus.ProcessTreeUnavailable({self.reason.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProcessTreeStatus_ProcessTreeUnavailable) and self.reason == __o.reason
    def __hash__(self) -> int:
        return super().__hash__()


class CleanupStatus:
    @classmethod
    def default(cls, ):
        return lambda: CleanupStatus_CleanupNotRequired()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_CleanupNotRequired(self) -> bool:
        return isinstance(self, CleanupStatus_CleanupNotRequired)
    @property
    def is_CleanupComplete(self) -> bool:
        return isinstance(self, CleanupStatus_CleanupComplete)
    @property
    def is_CleanupFailed(self) -> bool:
        return isinstance(self, CleanupStatus_CleanupFailed)

class CleanupStatus_CleanupNotRequired(CleanupStatus, NamedTuple('CleanupNotRequired', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.CleanupStatus.CleanupNotRequired'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CleanupStatus_CleanupNotRequired)
    def __hash__(self) -> int:
        return super().__hash__()

class CleanupStatus_CleanupComplete(CleanupStatus, NamedTuple('CleanupComplete', [('elapsedMs', Any), ('deadlineMs', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.CleanupStatus.CleanupComplete({_dafny.string_of(self.elapsedMs)}, {_dafny.string_of(self.deadlineMs)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CleanupStatus_CleanupComplete) and self.elapsedMs == __o.elapsedMs and self.deadlineMs == __o.deadlineMs
    def __hash__(self) -> int:
        return super().__hash__()

class CleanupStatus_CleanupFailed(CleanupStatus, NamedTuple('CleanupFailed', [('elapsedMs', Any), ('deadlineMs', Any), ('reason', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.CleanupStatus.CleanupFailed({_dafny.string_of(self.elapsedMs)}, {_dafny.string_of(self.deadlineMs)}, {self.reason.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CleanupStatus_CleanupFailed) and self.elapsedMs == __o.elapsedMs and self.deadlineMs == __o.deadlineMs and self.reason == __o.reason
    def __hash__(self) -> int:
        return super().__hash__()


class ExitStatus:
    @classmethod
    def default(cls, ):
        return lambda: ExitStatus_ExitNotObserved()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ExitNotObserved(self) -> bool:
        return isinstance(self, ExitStatus_ExitNotObserved)
    @property
    def is_ExitObserved(self) -> bool:
        return isinstance(self, ExitStatus_ExitObserved)

class ExitStatus_ExitNotObserved(ExitStatus, NamedTuple('ExitNotObserved', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.ExitStatus.ExitNotObserved'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ExitStatus_ExitNotObserved)
    def __hash__(self) -> int:
        return super().__hash__()

class ExitStatus_ExitObserved(ExitStatus, NamedTuple('ExitObserved', [('code', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.ExitStatus.ExitObserved({_dafny.string_of(self.code)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ExitStatus_ExitObserved) and self.code == __o.code
    def __hash__(self) -> int:
        return super().__hash__()


class EnforcementStatus:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [EnforcementStatus_NoEnforcement(), EnforcementStatus_PostHocFailure(), EnforcementStatus_InFlightProcessTreeTermination()]
    @classmethod
    def default(cls, ):
        return lambda: EnforcementStatus_NoEnforcement()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_NoEnforcement(self) -> bool:
        return isinstance(self, EnforcementStatus_NoEnforcement)
    @property
    def is_PostHocFailure(self) -> bool:
        return isinstance(self, EnforcementStatus_PostHocFailure)
    @property
    def is_InFlightProcessTreeTermination(self) -> bool:
        return isinstance(self, EnforcementStatus_InFlightProcessTreeTermination)

class EnforcementStatus_NoEnforcement(EnforcementStatus, NamedTuple('NoEnforcement', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.EnforcementStatus.NoEnforcement'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EnforcementStatus_NoEnforcement)
    def __hash__(self) -> int:
        return super().__hash__()

class EnforcementStatus_PostHocFailure(EnforcementStatus, NamedTuple('PostHocFailure', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.EnforcementStatus.PostHocFailure'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EnforcementStatus_PostHocFailure)
    def __hash__(self) -> int:
        return super().__hash__()

class EnforcementStatus_InFlightProcessTreeTermination(EnforcementStatus, NamedTuple('InFlightProcessTreeTermination', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.EnforcementStatus.InFlightProcessTreeTermination'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EnforcementStatus_InFlightProcessTreeTermination)
    def __hash__(self) -> int:
        return super().__hash__()


class TerminationScope:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [TerminationScope_NoTermination(), TerminationScope_ProcessTreeTermination()]
    @classmethod
    def default(cls, ):
        return lambda: TerminationScope_NoTermination()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_NoTermination(self) -> bool:
        return isinstance(self, TerminationScope_NoTermination)
    @property
    def is_ProcessTreeTermination(self) -> bool:
        return isinstance(self, TerminationScope_ProcessTreeTermination)

class TerminationScope_NoTermination(TerminationScope, NamedTuple('NoTermination', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.TerminationScope.NoTermination'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminationScope_NoTermination)
    def __hash__(self) -> int:
        return super().__hash__()

class TerminationScope_ProcessTreeTermination(TerminationScope, NamedTuple('ProcessTreeTermination', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.TerminationScope.ProcessTreeTermination'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminationScope_ProcessTreeTermination)
    def __hash__(self) -> int:
        return super().__hash__()


class SupervisionObservationV1:
    @classmethod
    def default(cls, ):
        return lambda: SupervisionObservationV1_SupervisionObservationV1(TimingCoverage.default()(), ConfluxResourceCeilings.ResourceMeasurement.default()(), ConfluxResourceCeilings.ResourceMeasurement.default()(), ProcessTreeStatus.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SupervisionObservationV1(self) -> bool:
        return isinstance(self, SupervisionObservationV1_SupervisionObservationV1)

class SupervisionObservationV1_SupervisionObservationV1(SupervisionObservationV1, NamedTuple('SupervisionObservationV1', [('timing', Any), ('peakDescendantRssMb', Any), ('combinedOutputBytes', Any), ('processTree', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.SupervisionObservationV1.SupervisionObservationV1({_dafny.string_of(self.timing)}, {_dafny.string_of(self.peakDescendantRssMb)}, {_dafny.string_of(self.combinedOutputBytes)}, {_dafny.string_of(self.processTree)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SupervisionObservationV1_SupervisionObservationV1) and self.timing == __o.timing and self.peakDescendantRssMb == __o.peakDescendantRssMb and self.combinedOutputBytes == __o.combinedOutputBytes and self.processTree == __o.processTree
    def __hash__(self) -> int:
        return super().__hash__()


class FirstBreachV1:
    @classmethod
    def default(cls, ):
        return lambda: FirstBreachV1_FirstBreachV1(int(0), ConfluxResourceCeilings.ResourceVerdict.default()(), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_FirstBreachV1(self) -> bool:
        return isinstance(self, FirstBreachV1_FirstBreachV1)

class FirstBreachV1_FirstBreachV1(FirstBreachV1, NamedTuple('FirstBreachV1', [('sampleIndex', Any), ('verdict', Any), ('measurementFailure', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.FirstBreachV1.FirstBreachV1({_dafny.string_of(self.sampleIndex)}, {_dafny.string_of(self.verdict)}, {_dafny.string_of(self.measurementFailure)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FirstBreachV1_FirstBreachV1) and self.sampleIndex == __o.sampleIndex and self.verdict == __o.verdict and self.measurementFailure == __o.measurementFailure
    def __hash__(self) -> int:
        return super().__hash__()


class SupervisionState:
    @classmethod
    def default(cls, ):
        return lambda: SupervisionState_SupervisionRunning(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SupervisionRunning(self) -> bool:
        return isinstance(self, SupervisionState_SupervisionRunning)
    @property
    def is_SupervisionTripped(self) -> bool:
        return isinstance(self, SupervisionState_SupervisionTripped)

class SupervisionState_SupervisionRunning(SupervisionState, NamedTuple('SupervisionRunning', [('samplesSeen', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.SupervisionState.SupervisionRunning({_dafny.string_of(self.samplesSeen)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SupervisionState_SupervisionRunning) and self.samplesSeen == __o.samplesSeen
    def __hash__(self) -> int:
        return super().__hash__()

class SupervisionState_SupervisionTripped(SupervisionState, NamedTuple('SupervisionTripped', [('record', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.SupervisionState.SupervisionTripped({_dafny.string_of(self.record)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SupervisionState_SupervisionTripped) and self.record == __o.record
    def __hash__(self) -> int:
        return super().__hash__()


class SupervisedOperationResultV1:
    @classmethod
    def default(cls, ):
        return lambda: SupervisedOperationResultV1_SupervisedOperationResultV1(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxResourceCeilings.ExecutionPhase.default()(), ConfluxResourceCeilings.ExecutionPhase.default()(), ExpandedResourceBudgetV1.default()(), ConfluxResourceCeilings.MeasurementStatus.default()(), SupervisionObservationV1.default()(), SupervisionState.default()(), CleanupStatus.default()(), ExitStatus.default()(), EnforcementStatus.default()(), TerminationScope.default()(), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxResourceCeilings.ResourceVerdict.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SupervisedOperationResultV1(self) -> bool:
        return isinstance(self, SupervisedOperationResultV1_SupervisedOperationResultV1)

class SupervisedOperationResultV1_SupervisedOperationResultV1(SupervisedOperationResultV1, NamedTuple('SupervisedOperationResultV1', [('schemaVersion', Any), ('policyId', Any), ('requestedPhase', Any), ('effectivePhase', Any), ('expandedCeilings', Any), ('observationStatus', Any), ('observation', Any), ('supervisionState', Any), ('cleanup', Any), ('exit', Any), ('enforcement', Any), ('termination', Any), ('wrapperExitCode', Any), ('reason', Any), ('verdict', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSupervisedOperationResult.SupervisedOperationResultV1.SupervisedOperationResultV1({_dafny.string_of(self.schemaVersion)}, {self.policyId.VerbatimString(True)}, {_dafny.string_of(self.requestedPhase)}, {_dafny.string_of(self.effectivePhase)}, {_dafny.string_of(self.expandedCeilings)}, {_dafny.string_of(self.observationStatus)}, {_dafny.string_of(self.observation)}, {_dafny.string_of(self.supervisionState)}, {_dafny.string_of(self.cleanup)}, {_dafny.string_of(self.exit)}, {_dafny.string_of(self.enforcement)}, {_dafny.string_of(self.termination)}, {_dafny.string_of(self.wrapperExitCode)}, {self.reason.VerbatimString(True)}, {_dafny.string_of(self.verdict)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SupervisedOperationResultV1_SupervisedOperationResultV1) and self.schemaVersion == __o.schemaVersion and self.policyId == __o.policyId and self.requestedPhase == __o.requestedPhase and self.effectivePhase == __o.effectivePhase and self.expandedCeilings == __o.expandedCeilings and self.observationStatus == __o.observationStatus and self.observation == __o.observation and self.supervisionState == __o.supervisionState and self.cleanup == __o.cleanup and self.exit == __o.exit and self.enforcement == __o.enforcement and self.termination == __o.termination and self.wrapperExitCode == __o.wrapperExitCode and self.reason == __o.reason and self.verdict == __o.verdict
    def __hash__(self) -> int:
        return super().__hash__()

