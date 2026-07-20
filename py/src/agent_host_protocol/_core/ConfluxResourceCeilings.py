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

# Module: ConfluxResourceCeilings

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def ValidMeasurement(measurement):
        return (((not (((measurement).status).is_MeasurementComplete) or (((measurement).value).is_Some)) and (not (((measurement).status).is_MeasurementPartial) or ((((measurement).value).is_Some) and ((len(((measurement).status).reason)) > (0))))) and (not (((measurement).status).is_MeasurementUnavailable) or ((len(((measurement).status).reason)) > (0)))) and (not ((((measurement).status).is_MeasurementNotRequested) or (((measurement).status).is_MeasurementUnavailable)) or (((measurement).value).is_None))

    @staticmethod
    def CompleteMeasurement(value):
        return ResourceMeasurement_ResourceMeasurement(ConfluxContract.Option_Some(value), MeasurementStatus_MeasurementComplete())

    @staticmethod
    def NotRequestedMeasurement():
        return ResourceMeasurement_ResourceMeasurement(ConfluxContract.Option_None(), MeasurementStatus_MeasurementNotRequested())

    @staticmethod
    def UnavailableMeasurement(reason):
        return ResourceMeasurement_ResourceMeasurement(ConfluxContract.Option_None(), MeasurementStatus_MeasurementUnavailable((reason if (len(reason)) > (0) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "measurement unavailable")))))

    @staticmethod
    def RequestedPhaseFor(state):
        source0_ = state
        if True:
            if source0_.is_HostPhaseUnavailable:
                return ExecutionPhase_Unknown()
        if True:
            d_0_hostGeneration_ = source0_.hostGeneration
            d_1_cachedGeneration_ = source0_.cachedGeneration
            if ((d_1_cachedGeneration_).is_Some) and (((d_1_cachedGeneration_).value) == (d_0_hostGeneration_)):
                return ExecutionPhase_Warm()
            elif True:
                return ExecutionPhase_Cold()

    @staticmethod
    def EffectivePhase(phase):
        if (phase) == (ExecutionPhase_Warm()):
            return ExecutionPhase_Warm()
        elif True:
            return ExecutionPhase_Cold()

    @staticmethod
    def TimeCeilingForPhase(budget, phase):
        if (default__.EffectivePhase(phase)) == (ExecutionPhase_Warm()):
            return (budget).warmMs
        elif True:
            return (budget).coldMs

    @staticmethod
    def MemoryCeilingForPhase(budget, phase):
        if (default__.EffectivePhase(phase)) == (ExecutionPhase_Warm()):
            return (budget).warmMb
        elif True:
            return (budget).coldMb

    @staticmethod
    def MemoryMeasurementRequired(budget, phase):
        return (default__.MemoryCeilingForPhase((budget).memory, phase)) > (0)

    @staticmethod
    def OutputMeasurementRequired(budget):
        return ((budget).maxOutputMb) > (0)

    @staticmethod
    def BytesToMegabytes(bytes_):
        if (bytes_) == (0):
            return 0
        elif True:
            return _dafny.euclidian_division(((bytes_) + ((1024) * (1024))) - (1), (1024) * (1024))

    @staticmethod
    def IsMeasurementComplete(measurement):
        return (default__.ValidMeasurement(measurement)) and (((measurement).status).is_MeasurementComplete)

    @staticmethod
    def MeasurementHasValue(measurement):
        return (default__.ValidMeasurement(measurement)) and (((measurement).value).is_Some)

    @staticmethod
    def RequiredMemoryMeasurementFailed(budget, phase, obs):
        return (default__.MemoryMeasurementRequired(budget, phase)) and (not(default__.IsMeasurementComplete((obs).peakRssMb)))

    @staticmethod
    def RequiredOutputMeasurementFailed(budget, obs):
        return (default__.OutputMeasurementRequired(budget)) and (not(default__.IsMeasurementComplete((obs).combinedOutputBytes)))

    @staticmethod
    def TimeBreached(budget, phase, obs):
        return ((default__.TimeCeilingForPhase((budget).time, phase)) > (0)) and (((obs).wallMs) > (default__.TimeCeilingForPhase((budget).time, phase)))

    @staticmethod
    def MemoryBreached(budget, phase, obs):
        return ((default__.MeasurementHasValue((obs).peakRssMb)) and ((default__.MemoryCeilingForPhase((budget).memory, phase)) > (0))) and (((((obs).peakRssMb).value).value) > (default__.MemoryCeilingForPhase((budget).memory, phase)))

    @staticmethod
    def OutputBreached(budget, obs):
        return ((default__.MeasurementHasValue((obs).combinedOutputBytes)) and (((budget).maxOutputMb) > (0))) and ((default__.BytesToMegabytes((((obs).combinedOutputBytes).value).value)) > ((budget).maxOutputMb))

    @staticmethod
    def DimensionViolated(budget, phase, obs, dimension):
        if (dimension) == (ResourceDimension_TimeDimension()):
            return default__.TimeBreached(budget, phase, obs)
        elif (dimension) == (ResourceDimension_MemoryDimension()):
            return (default__.MemoryBreached(budget, phase, obs)) or (default__.RequiredMemoryMeasurementFailed(budget, phase, obs))
        elif True:
            return (default__.OutputBreached(budget, obs)) or (default__.RequiredOutputMeasurementFailed(budget, obs))

    @staticmethod
    def ViolatedDimensions(budget, phase, obs):
        return (((_dafny.SeqWithoutIsStrInference([ResourceDimension_TimeDimension()]) if default__.TimeBreached(budget, phase, obs) else _dafny.SeqWithoutIsStrInference([]))) + ((_dafny.SeqWithoutIsStrInference([ResourceDimension_MemoryDimension()]) if (default__.MemoryBreached(budget, phase, obs)) or (default__.RequiredMemoryMeasurementFailed(budget, phase, obs)) else _dafny.SeqWithoutIsStrInference([])))) + ((_dafny.SeqWithoutIsStrInference([ResourceDimension_OutputDimension()]) if (default__.OutputBreached(budget, obs)) or (default__.RequiredOutputMeasurementFailed(budget, obs)) else _dafny.SeqWithoutIsStrInference([])))

    @staticmethod
    def BreachedDimensions(budget, phase, obs):
        return default__.ViolatedDimensions(budget, phase, obs)

    @staticmethod
    def LimitForDimension(budget, phase, dimension):
        if (dimension) == (ResourceDimension_TimeDimension()):
            return default__.TimeCeilingForPhase((budget).time, phase)
        elif (dimension) == (ResourceDimension_MemoryDimension()):
            return default__.MemoryCeilingForPhase((budget).memory, phase)
        elif True:
            return (budget).maxOutputMb

    @staticmethod
    def ObservedForDimension(obs, dimension):
        if (dimension) == (ResourceDimension_TimeDimension()):
            return ConfluxContract.Option_Some((obs).wallMs)
        elif (dimension) == (ResourceDimension_MemoryDimension()):
            return ((obs).peakRssMb).value
        elif True:
            source0_ = ((obs).combinedOutputBytes).value
            if True:
                if source0_.is_None:
                    return ConfluxContract.Option_None()
            if True:
                d_0_bytes_ = source0_.value
                return ConfluxContract.Option_Some(default__.BytesToMegabytes(d_0_bytes_))

    @staticmethod
    def Ok(phase):
        return ResourceVerdict_ResourceVerdict(EnforcementOutcome_EnforcementOk(), ConfluxContract.Option_None(), default__.EffectivePhase(phase), ConfluxContract.Option_None(), ConfluxContract.Option_None(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "within resource budget")))

    @staticmethod
    def ViolationReason(outcome, budget, phase, obs, dimension):
        d_0_prefix_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kill: ")) if (outcome) == (EnforcementOutcome_EnforcementKill()) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
        if ((dimension) == (ResourceDimension_MemoryDimension())) and (default__.RequiredMemoryMeasurementFailed(budget, phase, obs)):
            return (d_0_prefix_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "required process-tree RSS measurement unavailable")))
        elif ((dimension) == (ResourceDimension_OutputDimension())) and (default__.RequiredOutputMeasurementFailed(budget, obs)):
            return (d_0_prefix_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "required combined-output measurement unavailable")))
        elif (dimension) == (ResourceDimension_TimeDimension()):
            return (d_0_prefix_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "time budget exceeded")))
        elif (dimension) == (ResourceDimension_MemoryDimension()):
            return (d_0_prefix_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "memory budget exceeded")))
        elif True:
            return (d_0_prefix_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "output budget exceeded")))

    @staticmethod
    def Violation(outcome, budget, phase, obs, dimension):
        d_0_measurementFailure_ = (((dimension) == (ResourceDimension_MemoryDimension())) and (default__.RequiredMemoryMeasurementFailed(budget, phase, obs))) or (((dimension) == (ResourceDimension_OutputDimension())) and (default__.RequiredOutputMeasurementFailed(budget, obs)))
        return ResourceVerdict_ResourceVerdict(outcome, ConfluxContract.Option_Some(dimension), default__.EffectivePhase(phase), ConfluxContract.Option_Some(default__.LimitForDimension(budget, phase, dimension)), (ConfluxContract.Option_None() if d_0_measurementFailure_ else default__.ObservedForDimension(obs, dimension)), default__.ViolationReason(outcome, budget, phase, obs, dimension))

    @staticmethod
    def DecideAs(outcome, budget, phase, obs):
        d_0_dimensions_ = default__.ViolatedDimensions(budget, phase, obs)
        if (len(d_0_dimensions_)) == (0):
            return default__.Ok(phase)
        elif True:
            return default__.Violation(outcome, budget, phase, obs, (d_0_dimensions_)[0])

    @staticmethod
    def DecideRecorded(outcome, budget, phase, obs, dimension):
        return default__.Violation(outcome, budget, phase, obs, dimension)

    @staticmethod
    def RecordedViolationReason(outcome, dimension, measurementFailure):
        d_0_prefix_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kill: ")) if (outcome) == (EnforcementOutcome_EnforcementKill()) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
        if measurementFailure:
            if (dimension) == (ResourceDimension_MemoryDimension()):
                return (d_0_prefix_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "required process-tree RSS measurement unavailable")))
            elif True:
                return (d_0_prefix_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "required combined-output measurement unavailable")))
        elif (dimension) == (ResourceDimension_TimeDimension()):
            return (d_0_prefix_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "time budget exceeded")))
        elif (dimension) == (ResourceDimension_MemoryDimension()):
            return (d_0_prefix_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "memory budget exceeded")))
        elif True:
            return (d_0_prefix_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "output budget exceeded")))

    @staticmethod
    def RecordedViolation(outcome, budget, phase, dimension, observed, measurementFailure):
        return ResourceVerdict_ResourceVerdict(outcome, ConfluxContract.Option_Some(dimension), default__.EffectivePhase(phase), ConfluxContract.Option_Some(default__.LimitForDimension(budget, phase, dimension)), observed, default__.RecordedViolationReason(outcome, dimension, measurementFailure))

    @staticmethod
    def DecidePostHoc(budget, phase, obs):
        return default__.DecideAs(EnforcementOutcome_EnforcementFail(), budget, phase, obs)

    @staticmethod
    def DecideInFlight(budget, phase, obs):
        return default__.DecideAs(EnforcementOutcome_EnforcementKill(), budget, phase, obs)

    @staticmethod
    def MaxValue(left, right):
        if (left).is_None:
            return right
        elif (right).is_None:
            return left
        elif True:
            return ConfluxContract.Option_Some(((left).value if ((left).value) >= ((right).value) else (right).value))

    @staticmethod
    def AddValue(left, right):
        if (left).is_None:
            return right
        elif (right).is_None:
            return left
        elif True:
            return ConfluxContract.Option_Some(((left).value) + ((right).value))

    @staticmethod
    def AggregateStatus(left, right, value):
        if ((left).is_MeasurementComplete) and ((right).is_MeasurementComplete):
            return MeasurementStatus_MeasurementComplete()
        elif ((left).is_MeasurementNotRequested) and ((right).is_MeasurementNotRequested):
            return MeasurementStatus_MeasurementNotRequested()
        elif (value).is_Some:
            return MeasurementStatus_MeasurementPartial(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "provider aggregate contains incomplete measurements")))
        elif True:
            return MeasurementStatus_MeasurementUnavailable(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "provider aggregate measurements unavailable")))

    @staticmethod
    def MergeMaxMeasurement(left, right):
        d_0_value_ = default__.MaxValue((left).value, (right).value)
        return ResourceMeasurement_ResourceMeasurement(d_0_value_, default__.AggregateStatus((left).status, (right).status, d_0_value_))

    @staticmethod
    def MergeSumMeasurement(left, right):
        d_0_value_ = default__.AddValue((left).value, (right).value)
        return ResourceMeasurement_ResourceMeasurement(d_0_value_, default__.AggregateStatus((left).status, (right).status, d_0_value_))

    @staticmethod
    def MergeObservations(left, right):
        return ResourceObservation_ResourceObservation(((left).wallMs if ((left).wallMs) >= ((right).wallMs) else (right).wallMs), default__.MergeMaxMeasurement((left).peakRssMb, (right).peakRssMb), default__.MergeSumMeasurement((left).combinedOutputBytes, (right).combinedOutputBytes))

    @staticmethod
    def OutcomeRank(outcome):
        if (outcome) == (EnforcementOutcome_EnforcementOk()):
            return 0
        elif (outcome) == (EnforcementOutcome_EnforcementFail()):
            return 1
        elif True:
            return 2

    @staticmethod
    def DominantOutcome(prior, candidate):
        if (default__.OutcomeRank(prior)) >= (default__.OutcomeRank(candidate)):
            return prior
        elif True:
            return candidate

    @_dafny.classproperty
    def DefaultResourceCeilings(instance):
        return ResourceCeilings_ResourceCeilings(PhaseTimeCeilings_PhaseTimeCeilings(5000, 30000), PhaseMemoryCeilings_PhaseMemoryCeilings(1024, 2048), 128)

class ExecutionPhase:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [ExecutionPhase_Warm(), ExecutionPhase_Cold(), ExecutionPhase_Unknown()]
    @classmethod
    def default(cls, ):
        return lambda: ExecutionPhase_Warm()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Warm(self) -> bool:
        return isinstance(self, ExecutionPhase_Warm)
    @property
    def is_Cold(self) -> bool:
        return isinstance(self, ExecutionPhase_Cold)
    @property
    def is_Unknown(self) -> bool:
        return isinstance(self, ExecutionPhase_Unknown)

class ExecutionPhase_Warm(ExecutionPhase, NamedTuple('Warm', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.ExecutionPhase.Warm'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ExecutionPhase_Warm)
    def __hash__(self) -> int:
        return super().__hash__()

class ExecutionPhase_Cold(ExecutionPhase, NamedTuple('Cold', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.ExecutionPhase.Cold'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ExecutionPhase_Cold)
    def __hash__(self) -> int:
        return super().__hash__()

class ExecutionPhase_Unknown(ExecutionPhase, NamedTuple('Unknown', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.ExecutionPhase.Unknown'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ExecutionPhase_Unknown)
    def __hash__(self) -> int:
        return super().__hash__()


class PhaseTimeCeilings:
    @classmethod
    def default(cls, ):
        return lambda: PhaseTimeCeilings_PhaseTimeCeilings(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_PhaseTimeCeilings(self) -> bool:
        return isinstance(self, PhaseTimeCeilings_PhaseTimeCeilings)

class PhaseTimeCeilings_PhaseTimeCeilings(PhaseTimeCeilings, NamedTuple('PhaseTimeCeilings', [('warmMs', Any), ('coldMs', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.PhaseTimeCeilings.PhaseTimeCeilings({_dafny.string_of(self.warmMs)}, {_dafny.string_of(self.coldMs)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, PhaseTimeCeilings_PhaseTimeCeilings) and self.warmMs == __o.warmMs and self.coldMs == __o.coldMs
    def __hash__(self) -> int:
        return super().__hash__()


class PhaseMemoryCeilings:
    @classmethod
    def default(cls, ):
        return lambda: PhaseMemoryCeilings_PhaseMemoryCeilings(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_PhaseMemoryCeilings(self) -> bool:
        return isinstance(self, PhaseMemoryCeilings_PhaseMemoryCeilings)

class PhaseMemoryCeilings_PhaseMemoryCeilings(PhaseMemoryCeilings, NamedTuple('PhaseMemoryCeilings', [('warmMb', Any), ('coldMb', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.PhaseMemoryCeilings.PhaseMemoryCeilings({_dafny.string_of(self.warmMb)}, {_dafny.string_of(self.coldMb)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, PhaseMemoryCeilings_PhaseMemoryCeilings) and self.warmMb == __o.warmMb and self.coldMb == __o.coldMb
    def __hash__(self) -> int:
        return super().__hash__()


class ResourceCeilings:
    @classmethod
    def default(cls, ):
        return lambda: ResourceCeilings_ResourceCeilings(PhaseTimeCeilings.default()(), PhaseMemoryCeilings.default()(), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ResourceCeilings(self) -> bool:
        return isinstance(self, ResourceCeilings_ResourceCeilings)

class ResourceCeilings_ResourceCeilings(ResourceCeilings, NamedTuple('ResourceCeilings', [('time', Any), ('memory', Any), ('maxOutputMb', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.ResourceCeilings.ResourceCeilings({_dafny.string_of(self.time)}, {_dafny.string_of(self.memory)}, {_dafny.string_of(self.maxOutputMb)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResourceCeilings_ResourceCeilings) and self.time == __o.time and self.memory == __o.memory and self.maxOutputMb == __o.maxOutputMb
    def __hash__(self) -> int:
        return super().__hash__()


class MeasurementStatus:
    @classmethod
    def default(cls, ):
        return lambda: MeasurementStatus_MeasurementComplete()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_MeasurementComplete(self) -> bool:
        return isinstance(self, MeasurementStatus_MeasurementComplete)
    @property
    def is_MeasurementNotRequested(self) -> bool:
        return isinstance(self, MeasurementStatus_MeasurementNotRequested)
    @property
    def is_MeasurementPartial(self) -> bool:
        return isinstance(self, MeasurementStatus_MeasurementPartial)
    @property
    def is_MeasurementUnavailable(self) -> bool:
        return isinstance(self, MeasurementStatus_MeasurementUnavailable)

class MeasurementStatus_MeasurementComplete(MeasurementStatus, NamedTuple('MeasurementComplete', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.MeasurementStatus.MeasurementComplete'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, MeasurementStatus_MeasurementComplete)
    def __hash__(self) -> int:
        return super().__hash__()

class MeasurementStatus_MeasurementNotRequested(MeasurementStatus, NamedTuple('MeasurementNotRequested', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.MeasurementStatus.MeasurementNotRequested'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, MeasurementStatus_MeasurementNotRequested)
    def __hash__(self) -> int:
        return super().__hash__()

class MeasurementStatus_MeasurementPartial(MeasurementStatus, NamedTuple('MeasurementPartial', [('reason', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.MeasurementStatus.MeasurementPartial({self.reason.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, MeasurementStatus_MeasurementPartial) and self.reason == __o.reason
    def __hash__(self) -> int:
        return super().__hash__()

class MeasurementStatus_MeasurementUnavailable(MeasurementStatus, NamedTuple('MeasurementUnavailable', [('reason', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.MeasurementStatus.MeasurementUnavailable({self.reason.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, MeasurementStatus_MeasurementUnavailable) and self.reason == __o.reason
    def __hash__(self) -> int:
        return super().__hash__()


class ResourceMeasurement:
    @classmethod
    def default(cls, ):
        return lambda: ResourceMeasurement_ResourceMeasurement(ConfluxContract.Option.default()(), MeasurementStatus.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ResourceMeasurement(self) -> bool:
        return isinstance(self, ResourceMeasurement_ResourceMeasurement)

class ResourceMeasurement_ResourceMeasurement(ResourceMeasurement, NamedTuple('ResourceMeasurement', [('value', Any), ('status', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.ResourceMeasurement.ResourceMeasurement({_dafny.string_of(self.value)}, {_dafny.string_of(self.status)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResourceMeasurement_ResourceMeasurement) and self.value == __o.value and self.status == __o.status
    def __hash__(self) -> int:
        return super().__hash__()


class ResourceObservation:
    @classmethod
    def default(cls, ):
        return lambda: ResourceObservation_ResourceObservation(int(0), ResourceMeasurement.default()(), ResourceMeasurement.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ResourceObservation(self) -> bool:
        return isinstance(self, ResourceObservation_ResourceObservation)

class ResourceObservation_ResourceObservation(ResourceObservation, NamedTuple('ResourceObservation', [('wallMs', Any), ('peakRssMb', Any), ('combinedOutputBytes', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.ResourceObservation.ResourceObservation({_dafny.string_of(self.wallMs)}, {_dafny.string_of(self.peakRssMb)}, {_dafny.string_of(self.combinedOutputBytes)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResourceObservation_ResourceObservation) and self.wallMs == __o.wallMs and self.peakRssMb == __o.peakRssMb and self.combinedOutputBytes == __o.combinedOutputBytes
    def __hash__(self) -> int:
        return super().__hash__()


class ResourceDimension:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [ResourceDimension_TimeDimension(), ResourceDimension_MemoryDimension(), ResourceDimension_OutputDimension()]
    @classmethod
    def default(cls, ):
        return lambda: ResourceDimension_TimeDimension()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_TimeDimension(self) -> bool:
        return isinstance(self, ResourceDimension_TimeDimension)
    @property
    def is_MemoryDimension(self) -> bool:
        return isinstance(self, ResourceDimension_MemoryDimension)
    @property
    def is_OutputDimension(self) -> bool:
        return isinstance(self, ResourceDimension_OutputDimension)

class ResourceDimension_TimeDimension(ResourceDimension, NamedTuple('TimeDimension', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.ResourceDimension.TimeDimension'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResourceDimension_TimeDimension)
    def __hash__(self) -> int:
        return super().__hash__()

class ResourceDimension_MemoryDimension(ResourceDimension, NamedTuple('MemoryDimension', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.ResourceDimension.MemoryDimension'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResourceDimension_MemoryDimension)
    def __hash__(self) -> int:
        return super().__hash__()

class ResourceDimension_OutputDimension(ResourceDimension, NamedTuple('OutputDimension', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.ResourceDimension.OutputDimension'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResourceDimension_OutputDimension)
    def __hash__(self) -> int:
        return super().__hash__()


class EnforcementOutcome:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [EnforcementOutcome_EnforcementOk(), EnforcementOutcome_EnforcementFail(), EnforcementOutcome_EnforcementKill()]
    @classmethod
    def default(cls, ):
        return lambda: EnforcementOutcome_EnforcementOk()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_EnforcementOk(self) -> bool:
        return isinstance(self, EnforcementOutcome_EnforcementOk)
    @property
    def is_EnforcementFail(self) -> bool:
        return isinstance(self, EnforcementOutcome_EnforcementFail)
    @property
    def is_EnforcementKill(self) -> bool:
        return isinstance(self, EnforcementOutcome_EnforcementKill)

class EnforcementOutcome_EnforcementOk(EnforcementOutcome, NamedTuple('EnforcementOk', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.EnforcementOutcome.EnforcementOk'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EnforcementOutcome_EnforcementOk)
    def __hash__(self) -> int:
        return super().__hash__()

class EnforcementOutcome_EnforcementFail(EnforcementOutcome, NamedTuple('EnforcementFail', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.EnforcementOutcome.EnforcementFail'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EnforcementOutcome_EnforcementFail)
    def __hash__(self) -> int:
        return super().__hash__()

class EnforcementOutcome_EnforcementKill(EnforcementOutcome, NamedTuple('EnforcementKill', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.EnforcementOutcome.EnforcementKill'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EnforcementOutcome_EnforcementKill)
    def __hash__(self) -> int:
        return super().__hash__()


class ResourceVerdict:
    @classmethod
    def default(cls, ):
        return lambda: ResourceVerdict_ResourceVerdict(EnforcementOutcome.default()(), ConfluxContract.Option.default()(), ExecutionPhase.default()(), ConfluxContract.Option.default()(), ConfluxContract.Option.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ResourceVerdict(self) -> bool:
        return isinstance(self, ResourceVerdict_ResourceVerdict)

class ResourceVerdict_ResourceVerdict(ResourceVerdict, NamedTuple('ResourceVerdict', [('outcome', Any), ('dimension', Any), ('phase', Any), ('ceiling', Any), ('observed', Any), ('reason', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.ResourceVerdict.ResourceVerdict({_dafny.string_of(self.outcome)}, {_dafny.string_of(self.dimension)}, {_dafny.string_of(self.phase)}, {_dafny.string_of(self.ceiling)}, {_dafny.string_of(self.observed)}, {self.reason.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResourceVerdict_ResourceVerdict) and self.outcome == __o.outcome and self.dimension == __o.dimension and self.phase == __o.phase and self.ceiling == __o.ceiling and self.observed == __o.observed and self.reason == __o.reason
    def __hash__(self) -> int:
        return super().__hash__()


class CacheHostPhaseState:
    @classmethod
    def default(cls, ):
        return lambda: CacheHostPhaseState_HostPhaseUnavailable()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_HostPhaseUnavailable(self) -> bool:
        return isinstance(self, CacheHostPhaseState_HostPhaseUnavailable)
    @property
    def is_HostPhaseKnown(self) -> bool:
        return isinstance(self, CacheHostPhaseState_HostPhaseKnown)

class CacheHostPhaseState_HostPhaseUnavailable(CacheHostPhaseState, NamedTuple('HostPhaseUnavailable', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.CacheHostPhaseState.HostPhaseUnavailable'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CacheHostPhaseState_HostPhaseUnavailable)
    def __hash__(self) -> int:
        return super().__hash__()

class CacheHostPhaseState_HostPhaseKnown(CacheHostPhaseState, NamedTuple('HostPhaseKnown', [('hostGeneration', Any), ('cachedGeneration', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResourceCeilings.CacheHostPhaseState.HostPhaseKnown({_dafny.string_of(self.hostGeneration)}, {_dafny.string_of(self.cachedGeneration)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CacheHostPhaseState_HostPhaseKnown) and self.hostGeneration == __o.hostGeneration and self.cachedGeneration == __o.cachedGeneration
    def __hash__(self) -> int:
        return super().__hash__()

