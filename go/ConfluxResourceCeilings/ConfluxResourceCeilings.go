// Package ConfluxResourceCeilings
// Dafny module ConfluxResourceCeilings compiled into Go

package ConfluxResourceCeilings

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-verified/go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-verified/go/Changeset"
	m_Chat "github.com/joshmouch/ahp-verified/go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-verified/go/ClientMain"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-verified/go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-verified/go/ConfluxOrderedLog"
	m_ConfluxPrune "github.com/joshmouch/ahp-verified/go/ConfluxPrune"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-verified/go/ConfluxSeqRoute"
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
	return "ConfluxResourceCeilings.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ValidMeasurement(measurement ResourceMeasurement) bool {
	return (((!(((measurement).Dtor_status()).Is_MeasurementComplete()) || (((measurement).Dtor_value()).Is_Some())) && (!(((measurement).Dtor_status()).Is_MeasurementPartial()) || ((((measurement).Dtor_value()).Is_Some()) && ((_dafny.IntOfUint32((((measurement).Dtor_status()).Dtor_reason()).Cardinality())).Sign() == 1)))) && (!(((measurement).Dtor_status()).Is_MeasurementUnavailable()) || ((_dafny.IntOfUint32((((measurement).Dtor_status()).Dtor_reason()).Cardinality())).Sign() == 1))) && (!((((measurement).Dtor_status()).Is_MeasurementNotRequested()) || (((measurement).Dtor_status()).Is_MeasurementUnavailable())) || (((measurement).Dtor_value()).Is_None()))
}
func (_static *CompanionStruct_Default___) CompleteMeasurement(value _dafny.Int) ResourceMeasurement {
	return Companion_ResourceMeasurement_.Create_ResourceMeasurement_(m_ConfluxContract.Companion_Option_.Create_Some_(value), Companion_MeasurementStatus_.Create_MeasurementComplete_())
}
func (_static *CompanionStruct_Default___) NotRequestedMeasurement() ResourceMeasurement {
	return Companion_ResourceMeasurement_.Create_ResourceMeasurement_(m_ConfluxContract.Companion_Option_.Create_None_(), Companion_MeasurementStatus_.Create_MeasurementNotRequested_())
}
func (_static *CompanionStruct_Default___) UnavailableMeasurement(reason _dafny.Sequence) ResourceMeasurement {
	return Companion_ResourceMeasurement_.Create_ResourceMeasurement_(m_ConfluxContract.Companion_Option_.Create_None_(), Companion_MeasurementStatus_.Create_MeasurementUnavailable_((func() _dafny.Sequence {
		if (_dafny.IntOfUint32((reason).Cardinality())).Sign() == 1 {
			return reason
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("measurement unavailable")
	})()))
}
func (_static *CompanionStruct_Default___) RequestedPhaseFor(state CacheHostPhaseState) ExecutionPhase {
	var _source0 CacheHostPhaseState = state
	_ = _source0
	{
		if _source0.Is_HostPhaseUnavailable() {
			return Companion_ExecutionPhase_.Create_Unknown_()
		}
	}
	{
		var _0_hostGeneration _dafny.Int = _source0.Get_().(CacheHostPhaseState_HostPhaseKnown).HostGeneration
		_ = _0_hostGeneration
		var _1_cachedGeneration m_ConfluxContract.Option = _source0.Get_().(CacheHostPhaseState_HostPhaseKnown).CachedGeneration
		_ = _1_cachedGeneration
		if ((_1_cachedGeneration).Is_Some()) && (((_1_cachedGeneration).Dtor_value().(_dafny.Int)).Cmp(_0_hostGeneration) == 0) {
			return Companion_ExecutionPhase_.Create_Warm_()
		} else {
			return Companion_ExecutionPhase_.Create_Cold_()
		}
	}
}
func (_static *CompanionStruct_Default___) EffectivePhase(phase ExecutionPhase) ExecutionPhase {
	if (phase).Equals(Companion_ExecutionPhase_.Create_Warm_()) {
		return Companion_ExecutionPhase_.Create_Warm_()
	} else {
		return Companion_ExecutionPhase_.Create_Cold_()
	}
}
func (_static *CompanionStruct_Default___) TimeCeilingForPhase(budget PhaseTimeCeilings, phase ExecutionPhase) _dafny.Int {
	if (Companion_Default___.EffectivePhase(phase)).Equals(Companion_ExecutionPhase_.Create_Warm_()) {
		return (budget).Dtor_warmMs()
	} else {
		return (budget).Dtor_coldMs()
	}
}
func (_static *CompanionStruct_Default___) MemoryCeilingForPhase(budget PhaseMemoryCeilings, phase ExecutionPhase) _dafny.Int {
	if (Companion_Default___.EffectivePhase(phase)).Equals(Companion_ExecutionPhase_.Create_Warm_()) {
		return (budget).Dtor_warmMb()
	} else {
		return (budget).Dtor_coldMb()
	}
}
func (_static *CompanionStruct_Default___) MemoryMeasurementRequired(budget ResourceCeilings, phase ExecutionPhase) bool {
	return (Companion_Default___.MemoryCeilingForPhase((budget).Dtor_memory(), phase)).Sign() == 1
}
func (_static *CompanionStruct_Default___) OutputMeasurementRequired(budget ResourceCeilings) bool {
	return ((budget).Dtor_maxOutputMb()).Sign() == 1
}
func (_static *CompanionStruct_Default___) BytesToMegabytes(bytes _dafny.Int) _dafny.Int {
	if (bytes).Sign() == 0 {
		return _dafny.Zero
	} else {
		return (((bytes).Plus((_dafny.IntOfInt64(1024)).Times(_dafny.IntOfInt64(1024)))).Minus(_dafny.One)).DivBy((_dafny.IntOfInt64(1024)).Times(_dafny.IntOfInt64(1024)))
	}
}
func (_static *CompanionStruct_Default___) IsMeasurementComplete(measurement ResourceMeasurement) bool {
	return (Companion_Default___.ValidMeasurement(measurement)) && (((measurement).Dtor_status()).Is_MeasurementComplete())
}
func (_static *CompanionStruct_Default___) MeasurementHasValue(measurement ResourceMeasurement) bool {
	return (Companion_Default___.ValidMeasurement(measurement)) && (((measurement).Dtor_value()).Is_Some())
}
func (_static *CompanionStruct_Default___) RequiredMemoryMeasurementFailed(budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation) bool {
	return (Companion_Default___.MemoryMeasurementRequired(budget, phase)) && (!(Companion_Default___.IsMeasurementComplete((obs).Dtor_peakRssMb())))
}
func (_static *CompanionStruct_Default___) RequiredOutputMeasurementFailed(budget ResourceCeilings, obs ResourceObservation) bool {
	return (Companion_Default___.OutputMeasurementRequired(budget)) && (!(Companion_Default___.IsMeasurementComplete((obs).Dtor_combinedOutputBytes())))
}
func (_static *CompanionStruct_Default___) TimeBreached(budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation) bool {
	return ((Companion_Default___.TimeCeilingForPhase((budget).Dtor_time(), phase)).Sign() == 1) && (((obs).Dtor_wallMs()).Cmp(Companion_Default___.TimeCeilingForPhase((budget).Dtor_time(), phase)) > 0)
}
func (_static *CompanionStruct_Default___) MemoryBreached(budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation) bool {
	return ((Companion_Default___.MeasurementHasValue((obs).Dtor_peakRssMb())) && ((Companion_Default___.MemoryCeilingForPhase((budget).Dtor_memory(), phase)).Sign() == 1)) && (((((obs).Dtor_peakRssMb()).Dtor_value()).Dtor_value().(_dafny.Int)).Cmp(Companion_Default___.MemoryCeilingForPhase((budget).Dtor_memory(), phase)) > 0)
}
func (_static *CompanionStruct_Default___) OutputBreached(budget ResourceCeilings, obs ResourceObservation) bool {
	return ((Companion_Default___.MeasurementHasValue((obs).Dtor_combinedOutputBytes())) && (((budget).Dtor_maxOutputMb()).Sign() == 1)) && ((Companion_Default___.BytesToMegabytes((((obs).Dtor_combinedOutputBytes()).Dtor_value()).Dtor_value().(_dafny.Int))).Cmp((budget).Dtor_maxOutputMb()) > 0)
}
func (_static *CompanionStruct_Default___) DimensionViolated(budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation, dimension ResourceDimension) bool {
	if (dimension).Equals(Companion_ResourceDimension_.Create_TimeDimension_()) {
		return Companion_Default___.TimeBreached(budget, phase, obs)
	} else if (dimension).Equals(Companion_ResourceDimension_.Create_MemoryDimension_()) {
		return (Companion_Default___.MemoryBreached(budget, phase, obs)) || (Companion_Default___.RequiredMemoryMeasurementFailed(budget, phase, obs))
	} else {
		return (Companion_Default___.OutputBreached(budget, obs)) || (Companion_Default___.RequiredOutputMeasurementFailed(budget, obs))
	}
}
func (_static *CompanionStruct_Default___) ViolatedDimensions(budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if Companion_Default___.TimeBreached(budget, phase, obs) {
			return _dafny.SeqOf(Companion_ResourceDimension_.Create_TimeDimension_())
		}
		return _dafny.SeqOf()
	})(), (func() _dafny.Sequence {
		if (Companion_Default___.MemoryBreached(budget, phase, obs)) || (Companion_Default___.RequiredMemoryMeasurementFailed(budget, phase, obs)) {
			return _dafny.SeqOf(Companion_ResourceDimension_.Create_MemoryDimension_())
		}
		return _dafny.SeqOf()
	})()), (func() _dafny.Sequence {
		if (Companion_Default___.OutputBreached(budget, obs)) || (Companion_Default___.RequiredOutputMeasurementFailed(budget, obs)) {
			return _dafny.SeqOf(Companion_ResourceDimension_.Create_OutputDimension_())
		}
		return _dafny.SeqOf()
	})())
}
func (_static *CompanionStruct_Default___) BreachedDimensions(budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation) _dafny.Sequence {
	return Companion_Default___.ViolatedDimensions(budget, phase, obs)
}
func (_static *CompanionStruct_Default___) LimitForDimension(budget ResourceCeilings, phase ExecutionPhase, dimension ResourceDimension) _dafny.Int {
	if (dimension).Equals(Companion_ResourceDimension_.Create_TimeDimension_()) {
		return Companion_Default___.TimeCeilingForPhase((budget).Dtor_time(), phase)
	} else if (dimension).Equals(Companion_ResourceDimension_.Create_MemoryDimension_()) {
		return Companion_Default___.MemoryCeilingForPhase((budget).Dtor_memory(), phase)
	} else {
		return (budget).Dtor_maxOutputMb()
	}
}
func (_static *CompanionStruct_Default___) ObservedForDimension(obs ResourceObservation, dimension ResourceDimension) m_ConfluxContract.Option {
	if (dimension).Equals(Companion_ResourceDimension_.Create_TimeDimension_()) {
		return m_ConfluxContract.Companion_Option_.Create_Some_((obs).Dtor_wallMs())
	} else if (dimension).Equals(Companion_ResourceDimension_.Create_MemoryDimension_()) {
		return ((obs).Dtor_peakRssMb()).Dtor_value()
	} else {
		var _source0 m_ConfluxContract.Option = ((obs).Dtor_combinedOutputBytes()).Dtor_value()
		_ = _source0
		{
			if _source0.Is_None() {
				return m_ConfluxContract.Companion_Option_.Create_None_()
			}
		}
		{
			var _0_bytes _dafny.Int = _source0.Get_().(m_ConfluxContract.Option_Some).Value.(_dafny.Int)
			_ = _0_bytes
			return m_ConfluxContract.Companion_Option_.Create_Some_(Companion_Default___.BytesToMegabytes(_0_bytes))
		}
	}
}
func (_static *CompanionStruct_Default___) Ok(phase ExecutionPhase) ResourceVerdict {
	return Companion_ResourceVerdict_.Create_ResourceVerdict_(Companion_EnforcementOutcome_.Create_EnforcementOk_(), m_ConfluxContract.Companion_Option_.Create_None_(), Companion_Default___.EffectivePhase(phase), m_ConfluxContract.Companion_Option_.Create_None_(), m_ConfluxContract.Companion_Option_.Create_None_(), _dafny.UnicodeSeqOfUtf8Bytes("within resource budget"))
}
func (_static *CompanionStruct_Default___) ViolationReason(outcome EnforcementOutcome, budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation, dimension ResourceDimension) _dafny.Sequence {
	var _0_prefix _dafny.Sequence = (func() _dafny.Sequence {
		if (outcome).Equals(Companion_EnforcementOutcome_.Create_EnforcementKill_()) {
			return _dafny.UnicodeSeqOfUtf8Bytes("kill: ")
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("")
	})()
	_ = _0_prefix
	if ((dimension).Equals(Companion_ResourceDimension_.Create_MemoryDimension_())) && (Companion_Default___.RequiredMemoryMeasurementFailed(budget, phase, obs)) {
		return _dafny.Companion_Sequence_.Concatenate(_0_prefix, _dafny.UnicodeSeqOfUtf8Bytes("required process-tree RSS measurement unavailable"))
	} else if ((dimension).Equals(Companion_ResourceDimension_.Create_OutputDimension_())) && (Companion_Default___.RequiredOutputMeasurementFailed(budget, obs)) {
		return _dafny.Companion_Sequence_.Concatenate(_0_prefix, _dafny.UnicodeSeqOfUtf8Bytes("required combined-output measurement unavailable"))
	} else if (dimension).Equals(Companion_ResourceDimension_.Create_TimeDimension_()) {
		return _dafny.Companion_Sequence_.Concatenate(_0_prefix, _dafny.UnicodeSeqOfUtf8Bytes("time budget exceeded"))
	} else if (dimension).Equals(Companion_ResourceDimension_.Create_MemoryDimension_()) {
		return _dafny.Companion_Sequence_.Concatenate(_0_prefix, _dafny.UnicodeSeqOfUtf8Bytes("memory budget exceeded"))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_0_prefix, _dafny.UnicodeSeqOfUtf8Bytes("output budget exceeded"))
	}
}
func (_static *CompanionStruct_Default___) Violation(outcome EnforcementOutcome, budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation, dimension ResourceDimension) ResourceVerdict {
	var _0_measurementFailure bool = (((dimension).Equals(Companion_ResourceDimension_.Create_MemoryDimension_())) && (Companion_Default___.RequiredMemoryMeasurementFailed(budget, phase, obs))) || (((dimension).Equals(Companion_ResourceDimension_.Create_OutputDimension_())) && (Companion_Default___.RequiredOutputMeasurementFailed(budget, obs)))
	_ = _0_measurementFailure
	return Companion_ResourceVerdict_.Create_ResourceVerdict_(outcome, m_ConfluxContract.Companion_Option_.Create_Some_(dimension), Companion_Default___.EffectivePhase(phase), m_ConfluxContract.Companion_Option_.Create_Some_(Companion_Default___.LimitForDimension(budget, phase, dimension)), (func() m_ConfluxContract.Option {
		if _0_measurementFailure {
			return m_ConfluxContract.Companion_Option_.Create_None_()
		}
		return Companion_Default___.ObservedForDimension(obs, dimension)
	})(), Companion_Default___.ViolationReason(outcome, budget, phase, obs, dimension))
}
func (_static *CompanionStruct_Default___) DecideAs(outcome EnforcementOutcome, budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation) ResourceVerdict {
	var _0_dimensions _dafny.Sequence = Companion_Default___.ViolatedDimensions(budget, phase, obs)
	_ = _0_dimensions
	if (_dafny.IntOfUint32((_0_dimensions).Cardinality())).Sign() == 0 {
		return Companion_Default___.Ok(phase)
	} else {
		return Companion_Default___.Violation(outcome, budget, phase, obs, (_0_dimensions).Select(0).(ResourceDimension))
	}
}
func (_static *CompanionStruct_Default___) DecideRecorded(outcome EnforcementOutcome, budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation, dimension ResourceDimension) ResourceVerdict {
	return Companion_Default___.Violation(outcome, budget, phase, obs, dimension)
}
func (_static *CompanionStruct_Default___) RecordedViolationReason(outcome EnforcementOutcome, dimension ResourceDimension, measurementFailure bool) _dafny.Sequence {
	var _0_prefix _dafny.Sequence = (func() _dafny.Sequence {
		if (outcome).Equals(Companion_EnforcementOutcome_.Create_EnforcementKill_()) {
			return _dafny.UnicodeSeqOfUtf8Bytes("kill: ")
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("")
	})()
	_ = _0_prefix
	if measurementFailure {
		if (dimension).Equals(Companion_ResourceDimension_.Create_MemoryDimension_()) {
			return _dafny.Companion_Sequence_.Concatenate(_0_prefix, _dafny.UnicodeSeqOfUtf8Bytes("required process-tree RSS measurement unavailable"))
		} else {
			return _dafny.Companion_Sequence_.Concatenate(_0_prefix, _dafny.UnicodeSeqOfUtf8Bytes("required combined-output measurement unavailable"))
		}
	} else if (dimension).Equals(Companion_ResourceDimension_.Create_TimeDimension_()) {
		return _dafny.Companion_Sequence_.Concatenate(_0_prefix, _dafny.UnicodeSeqOfUtf8Bytes("time budget exceeded"))
	} else if (dimension).Equals(Companion_ResourceDimension_.Create_MemoryDimension_()) {
		return _dafny.Companion_Sequence_.Concatenate(_0_prefix, _dafny.UnicodeSeqOfUtf8Bytes("memory budget exceeded"))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_0_prefix, _dafny.UnicodeSeqOfUtf8Bytes("output budget exceeded"))
	}
}
func (_static *CompanionStruct_Default___) RecordedViolation(outcome EnforcementOutcome, budget ResourceCeilings, phase ExecutionPhase, dimension ResourceDimension, observed m_ConfluxContract.Option, measurementFailure bool) ResourceVerdict {
	return Companion_ResourceVerdict_.Create_ResourceVerdict_(outcome, m_ConfluxContract.Companion_Option_.Create_Some_(dimension), Companion_Default___.EffectivePhase(phase), m_ConfluxContract.Companion_Option_.Create_Some_(Companion_Default___.LimitForDimension(budget, phase, dimension)), observed, Companion_Default___.RecordedViolationReason(outcome, dimension, measurementFailure))
}
func (_static *CompanionStruct_Default___) DecidePostHoc(budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation) ResourceVerdict {
	return Companion_Default___.DecideAs(Companion_EnforcementOutcome_.Create_EnforcementFail_(), budget, phase, obs)
}
func (_static *CompanionStruct_Default___) DecideInFlight(budget ResourceCeilings, phase ExecutionPhase, obs ResourceObservation) ResourceVerdict {
	return Companion_Default___.DecideAs(Companion_EnforcementOutcome_.Create_EnforcementKill_(), budget, phase, obs)
}
func (_static *CompanionStruct_Default___) MaxValue(left m_ConfluxContract.Option, right m_ConfluxContract.Option) m_ConfluxContract.Option {
	if (left).Is_None() {
		return right
	} else if (right).Is_None() {
		return left
	} else {
		return m_ConfluxContract.Companion_Option_.Create_Some_((func() _dafny.Int {
			if ((left).Dtor_value().(_dafny.Int)).Cmp((right).Dtor_value().(_dafny.Int)) >= 0 {
				return (left).Dtor_value().(_dafny.Int)
			}
			return (right).Dtor_value().(_dafny.Int)
		})())
	}
}
func (_static *CompanionStruct_Default___) AddValue(left m_ConfluxContract.Option, right m_ConfluxContract.Option) m_ConfluxContract.Option {
	if (left).Is_None() {
		return right
	} else if (right).Is_None() {
		return left
	} else {
		return m_ConfluxContract.Companion_Option_.Create_Some_(((left).Dtor_value().(_dafny.Int)).Plus((right).Dtor_value().(_dafny.Int)))
	}
}
func (_static *CompanionStruct_Default___) AggregateStatus(left MeasurementStatus, right MeasurementStatus, value m_ConfluxContract.Option) MeasurementStatus {
	if ((left).Is_MeasurementComplete()) && ((right).Is_MeasurementComplete()) {
		return Companion_MeasurementStatus_.Create_MeasurementComplete_()
	} else if ((left).Is_MeasurementNotRequested()) && ((right).Is_MeasurementNotRequested()) {
		return Companion_MeasurementStatus_.Create_MeasurementNotRequested_()
	} else if (value).Is_Some() {
		return Companion_MeasurementStatus_.Create_MeasurementPartial_(_dafny.UnicodeSeqOfUtf8Bytes("provider aggregate contains incomplete measurements"))
	} else {
		return Companion_MeasurementStatus_.Create_MeasurementUnavailable_(_dafny.UnicodeSeqOfUtf8Bytes("provider aggregate measurements unavailable"))
	}
}
func (_static *CompanionStruct_Default___) MergeMaxMeasurement(left ResourceMeasurement, right ResourceMeasurement) ResourceMeasurement {
	var _0_value m_ConfluxContract.Option = Companion_Default___.MaxValue((left).Dtor_value(), (right).Dtor_value())
	_ = _0_value
	return Companion_ResourceMeasurement_.Create_ResourceMeasurement_(_0_value, Companion_Default___.AggregateStatus((left).Dtor_status(), (right).Dtor_status(), _0_value))
}
func (_static *CompanionStruct_Default___) MergeSumMeasurement(left ResourceMeasurement, right ResourceMeasurement) ResourceMeasurement {
	var _0_value m_ConfluxContract.Option = Companion_Default___.AddValue((left).Dtor_value(), (right).Dtor_value())
	_ = _0_value
	return Companion_ResourceMeasurement_.Create_ResourceMeasurement_(_0_value, Companion_Default___.AggregateStatus((left).Dtor_status(), (right).Dtor_status(), _0_value))
}
func (_static *CompanionStruct_Default___) MergeObservations(left ResourceObservation, right ResourceObservation) ResourceObservation {
	return Companion_ResourceObservation_.Create_ResourceObservation_((func() _dafny.Int {
		if ((left).Dtor_wallMs()).Cmp((right).Dtor_wallMs()) >= 0 {
			return (left).Dtor_wallMs()
		}
		return (right).Dtor_wallMs()
	})(), Companion_Default___.MergeMaxMeasurement((left).Dtor_peakRssMb(), (right).Dtor_peakRssMb()), Companion_Default___.MergeSumMeasurement((left).Dtor_combinedOutputBytes(), (right).Dtor_combinedOutputBytes()))
}
func (_static *CompanionStruct_Default___) OutcomeRank(outcome EnforcementOutcome) _dafny.Int {
	if (outcome).Equals(Companion_EnforcementOutcome_.Create_EnforcementOk_()) {
		return _dafny.Zero
	} else if (outcome).Equals(Companion_EnforcementOutcome_.Create_EnforcementFail_()) {
		return _dafny.One
	} else {
		return _dafny.IntOfInt64(2)
	}
}
func (_static *CompanionStruct_Default___) DominantOutcome(prior EnforcementOutcome, candidate EnforcementOutcome) EnforcementOutcome {
	if (Companion_Default___.OutcomeRank(prior)).Cmp(Companion_Default___.OutcomeRank(candidate)) >= 0 {
		return prior
	} else {
		return candidate
	}
}
func (_static *CompanionStruct_Default___) DefaultResourceCeilings() ResourceCeilings {
	return Companion_ResourceCeilings_.Create_ResourceCeilings_(Companion_PhaseTimeCeilings_.Create_PhaseTimeCeilings_(_dafny.IntOfInt64(5000), _dafny.IntOfInt64(30000)), Companion_PhaseMemoryCeilings_.Create_PhaseMemoryCeilings_(_dafny.IntOfInt64(1024), _dafny.IntOfInt64(2048)), _dafny.IntOfInt64(128))
}

// End of class Default__

// Definition of datatype ExecutionPhase
type ExecutionPhase struct {
	Data_ExecutionPhase_
}

func (_this ExecutionPhase) Get_() Data_ExecutionPhase_ {
	return _this.Data_ExecutionPhase_
}

type Data_ExecutionPhase_ interface {
	isExecutionPhase()
}

type CompanionStruct_ExecutionPhase_ struct {
}

var Companion_ExecutionPhase_ = CompanionStruct_ExecutionPhase_{}

type ExecutionPhase_Warm struct {
}

func (ExecutionPhase_Warm) isExecutionPhase() {}

func (CompanionStruct_ExecutionPhase_) Create_Warm_() ExecutionPhase {
	return ExecutionPhase{ExecutionPhase_Warm{}}
}

func (_this ExecutionPhase) Is_Warm() bool {
	_, ok := _this.Get_().(ExecutionPhase_Warm)
	return ok
}

type ExecutionPhase_Cold struct {
}

func (ExecutionPhase_Cold) isExecutionPhase() {}

func (CompanionStruct_ExecutionPhase_) Create_Cold_() ExecutionPhase {
	return ExecutionPhase{ExecutionPhase_Cold{}}
}

func (_this ExecutionPhase) Is_Cold() bool {
	_, ok := _this.Get_().(ExecutionPhase_Cold)
	return ok
}

type ExecutionPhase_Unknown struct {
}

func (ExecutionPhase_Unknown) isExecutionPhase() {}

func (CompanionStruct_ExecutionPhase_) Create_Unknown_() ExecutionPhase {
	return ExecutionPhase{ExecutionPhase_Unknown{}}
}

func (_this ExecutionPhase) Is_Unknown() bool {
	_, ok := _this.Get_().(ExecutionPhase_Unknown)
	return ok
}

func (CompanionStruct_ExecutionPhase_) Default() ExecutionPhase {
	return Companion_ExecutionPhase_.Create_Warm_()
}

func (_ CompanionStruct_ExecutionPhase_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_ExecutionPhase_.Create_Warm_(), true
		case 1:
			return Companion_ExecutionPhase_.Create_Cold_(), true
		case 2:
			return Companion_ExecutionPhase_.Create_Unknown_(), true
		default:
			return ExecutionPhase{}, false
		}
	}
}

func (_this ExecutionPhase) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case ExecutionPhase_Warm:
		{
			return "ConfluxResourceCeilings.ExecutionPhase.Warm"
		}
	case ExecutionPhase_Cold:
		{
			return "ConfluxResourceCeilings.ExecutionPhase.Cold"
		}
	case ExecutionPhase_Unknown:
		{
			return "ConfluxResourceCeilings.ExecutionPhase.Unknown"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ExecutionPhase) Equals(other ExecutionPhase) bool {
	switch _this.Get_().(type) {
	case ExecutionPhase_Warm:
		{
			_, ok := other.Get_().(ExecutionPhase_Warm)
			return ok
		}
	case ExecutionPhase_Cold:
		{
			_, ok := other.Get_().(ExecutionPhase_Cold)
			return ok
		}
	case ExecutionPhase_Unknown:
		{
			_, ok := other.Get_().(ExecutionPhase_Unknown)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ExecutionPhase) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ExecutionPhase)
	return ok && _this.Equals(typed)
}

func Type_ExecutionPhase_() _dafny.TypeDescriptor {
	return type_ExecutionPhase_{}
}

type type_ExecutionPhase_ struct {
}

func (_this type_ExecutionPhase_) Default() interface{} {
	return Companion_ExecutionPhase_.Default()
}

func (_this type_ExecutionPhase_) String() string {
	return "ConfluxResourceCeilings.ExecutionPhase"
}
func (_this ExecutionPhase) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ExecutionPhase{}

// End of datatype ExecutionPhase

// Definition of datatype PhaseTimeCeilings
type PhaseTimeCeilings struct {
	Data_PhaseTimeCeilings_
}

func (_this PhaseTimeCeilings) Get_() Data_PhaseTimeCeilings_ {
	return _this.Data_PhaseTimeCeilings_
}

type Data_PhaseTimeCeilings_ interface {
	isPhaseTimeCeilings()
}

type CompanionStruct_PhaseTimeCeilings_ struct {
}

var Companion_PhaseTimeCeilings_ = CompanionStruct_PhaseTimeCeilings_{}

type PhaseTimeCeilings_PhaseTimeCeilings struct {
	WarmMs _dafny.Int
	ColdMs _dafny.Int
}

func (PhaseTimeCeilings_PhaseTimeCeilings) isPhaseTimeCeilings() {}

func (CompanionStruct_PhaseTimeCeilings_) Create_PhaseTimeCeilings_(WarmMs _dafny.Int, ColdMs _dafny.Int) PhaseTimeCeilings {
	return PhaseTimeCeilings{PhaseTimeCeilings_PhaseTimeCeilings{WarmMs, ColdMs}}
}

func (_this PhaseTimeCeilings) Is_PhaseTimeCeilings() bool {
	_, ok := _this.Get_().(PhaseTimeCeilings_PhaseTimeCeilings)
	return ok
}

func (CompanionStruct_PhaseTimeCeilings_) Default() PhaseTimeCeilings {
	return Companion_PhaseTimeCeilings_.Create_PhaseTimeCeilings_(_dafny.Zero, _dafny.Zero)
}

func (_this PhaseTimeCeilings) Dtor_warmMs() _dafny.Int {
	return _this.Get_().(PhaseTimeCeilings_PhaseTimeCeilings).WarmMs
}

func (_this PhaseTimeCeilings) Dtor_coldMs() _dafny.Int {
	return _this.Get_().(PhaseTimeCeilings_PhaseTimeCeilings).ColdMs
}

func (_this PhaseTimeCeilings) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case PhaseTimeCeilings_PhaseTimeCeilings:
		{
			return "ConfluxResourceCeilings.PhaseTimeCeilings.PhaseTimeCeilings" + "(" + _dafny.String(data.WarmMs) + ", " + _dafny.String(data.ColdMs) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this PhaseTimeCeilings) Equals(other PhaseTimeCeilings) bool {
	switch data1 := _this.Get_().(type) {
	case PhaseTimeCeilings_PhaseTimeCeilings:
		{
			data2, ok := other.Get_().(PhaseTimeCeilings_PhaseTimeCeilings)
			return ok && data1.WarmMs.Cmp(data2.WarmMs) == 0 && data1.ColdMs.Cmp(data2.ColdMs) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this PhaseTimeCeilings) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(PhaseTimeCeilings)
	return ok && _this.Equals(typed)
}

func Type_PhaseTimeCeilings_() _dafny.TypeDescriptor {
	return type_PhaseTimeCeilings_{}
}

type type_PhaseTimeCeilings_ struct {
}

func (_this type_PhaseTimeCeilings_) Default() interface{} {
	return Companion_PhaseTimeCeilings_.Default()
}

func (_this type_PhaseTimeCeilings_) String() string {
	return "ConfluxResourceCeilings.PhaseTimeCeilings"
}
func (_this PhaseTimeCeilings) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = PhaseTimeCeilings{}

// End of datatype PhaseTimeCeilings

// Definition of datatype PhaseMemoryCeilings
type PhaseMemoryCeilings struct {
	Data_PhaseMemoryCeilings_
}

func (_this PhaseMemoryCeilings) Get_() Data_PhaseMemoryCeilings_ {
	return _this.Data_PhaseMemoryCeilings_
}

type Data_PhaseMemoryCeilings_ interface {
	isPhaseMemoryCeilings()
}

type CompanionStruct_PhaseMemoryCeilings_ struct {
}

var Companion_PhaseMemoryCeilings_ = CompanionStruct_PhaseMemoryCeilings_{}

type PhaseMemoryCeilings_PhaseMemoryCeilings struct {
	WarmMb _dafny.Int
	ColdMb _dafny.Int
}

func (PhaseMemoryCeilings_PhaseMemoryCeilings) isPhaseMemoryCeilings() {}

func (CompanionStruct_PhaseMemoryCeilings_) Create_PhaseMemoryCeilings_(WarmMb _dafny.Int, ColdMb _dafny.Int) PhaseMemoryCeilings {
	return PhaseMemoryCeilings{PhaseMemoryCeilings_PhaseMemoryCeilings{WarmMb, ColdMb}}
}

func (_this PhaseMemoryCeilings) Is_PhaseMemoryCeilings() bool {
	_, ok := _this.Get_().(PhaseMemoryCeilings_PhaseMemoryCeilings)
	return ok
}

func (CompanionStruct_PhaseMemoryCeilings_) Default() PhaseMemoryCeilings {
	return Companion_PhaseMemoryCeilings_.Create_PhaseMemoryCeilings_(_dafny.Zero, _dafny.Zero)
}

func (_this PhaseMemoryCeilings) Dtor_warmMb() _dafny.Int {
	return _this.Get_().(PhaseMemoryCeilings_PhaseMemoryCeilings).WarmMb
}

func (_this PhaseMemoryCeilings) Dtor_coldMb() _dafny.Int {
	return _this.Get_().(PhaseMemoryCeilings_PhaseMemoryCeilings).ColdMb
}

func (_this PhaseMemoryCeilings) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case PhaseMemoryCeilings_PhaseMemoryCeilings:
		{
			return "ConfluxResourceCeilings.PhaseMemoryCeilings.PhaseMemoryCeilings" + "(" + _dafny.String(data.WarmMb) + ", " + _dafny.String(data.ColdMb) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this PhaseMemoryCeilings) Equals(other PhaseMemoryCeilings) bool {
	switch data1 := _this.Get_().(type) {
	case PhaseMemoryCeilings_PhaseMemoryCeilings:
		{
			data2, ok := other.Get_().(PhaseMemoryCeilings_PhaseMemoryCeilings)
			return ok && data1.WarmMb.Cmp(data2.WarmMb) == 0 && data1.ColdMb.Cmp(data2.ColdMb) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this PhaseMemoryCeilings) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(PhaseMemoryCeilings)
	return ok && _this.Equals(typed)
}

func Type_PhaseMemoryCeilings_() _dafny.TypeDescriptor {
	return type_PhaseMemoryCeilings_{}
}

type type_PhaseMemoryCeilings_ struct {
}

func (_this type_PhaseMemoryCeilings_) Default() interface{} {
	return Companion_PhaseMemoryCeilings_.Default()
}

func (_this type_PhaseMemoryCeilings_) String() string {
	return "ConfluxResourceCeilings.PhaseMemoryCeilings"
}
func (_this PhaseMemoryCeilings) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = PhaseMemoryCeilings{}

// End of datatype PhaseMemoryCeilings

// Definition of datatype ResourceCeilings
type ResourceCeilings struct {
	Data_ResourceCeilings_
}

func (_this ResourceCeilings) Get_() Data_ResourceCeilings_ {
	return _this.Data_ResourceCeilings_
}

type Data_ResourceCeilings_ interface {
	isResourceCeilings()
}

type CompanionStruct_ResourceCeilings_ struct {
}

var Companion_ResourceCeilings_ = CompanionStruct_ResourceCeilings_{}

type ResourceCeilings_ResourceCeilings struct {
	Time        PhaseTimeCeilings
	Memory      PhaseMemoryCeilings
	MaxOutputMb _dafny.Int
}

func (ResourceCeilings_ResourceCeilings) isResourceCeilings() {}

func (CompanionStruct_ResourceCeilings_) Create_ResourceCeilings_(Time PhaseTimeCeilings, Memory PhaseMemoryCeilings, MaxOutputMb _dafny.Int) ResourceCeilings {
	return ResourceCeilings{ResourceCeilings_ResourceCeilings{Time, Memory, MaxOutputMb}}
}

func (_this ResourceCeilings) Is_ResourceCeilings() bool {
	_, ok := _this.Get_().(ResourceCeilings_ResourceCeilings)
	return ok
}

func (CompanionStruct_ResourceCeilings_) Default() ResourceCeilings {
	return Companion_ResourceCeilings_.Create_ResourceCeilings_(Companion_PhaseTimeCeilings_.Default(), Companion_PhaseMemoryCeilings_.Default(), _dafny.Zero)
}

func (_this ResourceCeilings) Dtor_time() PhaseTimeCeilings {
	return _this.Get_().(ResourceCeilings_ResourceCeilings).Time
}

func (_this ResourceCeilings) Dtor_memory() PhaseMemoryCeilings {
	return _this.Get_().(ResourceCeilings_ResourceCeilings).Memory
}

func (_this ResourceCeilings) Dtor_maxOutputMb() _dafny.Int {
	return _this.Get_().(ResourceCeilings_ResourceCeilings).MaxOutputMb
}

func (_this ResourceCeilings) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ResourceCeilings_ResourceCeilings:
		{
			return "ConfluxResourceCeilings.ResourceCeilings.ResourceCeilings" + "(" + _dafny.String(data.Time) + ", " + _dafny.String(data.Memory) + ", " + _dafny.String(data.MaxOutputMb) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ResourceCeilings) Equals(other ResourceCeilings) bool {
	switch data1 := _this.Get_().(type) {
	case ResourceCeilings_ResourceCeilings:
		{
			data2, ok := other.Get_().(ResourceCeilings_ResourceCeilings)
			return ok && data1.Time.Equals(data2.Time) && data1.Memory.Equals(data2.Memory) && data1.MaxOutputMb.Cmp(data2.MaxOutputMb) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ResourceCeilings) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ResourceCeilings)
	return ok && _this.Equals(typed)
}

func Type_ResourceCeilings_() _dafny.TypeDescriptor {
	return type_ResourceCeilings_{}
}

type type_ResourceCeilings_ struct {
}

func (_this type_ResourceCeilings_) Default() interface{} {
	return Companion_ResourceCeilings_.Default()
}

func (_this type_ResourceCeilings_) String() string {
	return "ConfluxResourceCeilings.ResourceCeilings"
}
func (_this ResourceCeilings) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ResourceCeilings{}

// End of datatype ResourceCeilings

// Definition of datatype MeasurementStatus
type MeasurementStatus struct {
	Data_MeasurementStatus_
}

func (_this MeasurementStatus) Get_() Data_MeasurementStatus_ {
	return _this.Data_MeasurementStatus_
}

type Data_MeasurementStatus_ interface {
	isMeasurementStatus()
}

type CompanionStruct_MeasurementStatus_ struct {
}

var Companion_MeasurementStatus_ = CompanionStruct_MeasurementStatus_{}

type MeasurementStatus_MeasurementComplete struct {
}

func (MeasurementStatus_MeasurementComplete) isMeasurementStatus() {}

func (CompanionStruct_MeasurementStatus_) Create_MeasurementComplete_() MeasurementStatus {
	return MeasurementStatus{MeasurementStatus_MeasurementComplete{}}
}

func (_this MeasurementStatus) Is_MeasurementComplete() bool {
	_, ok := _this.Get_().(MeasurementStatus_MeasurementComplete)
	return ok
}

type MeasurementStatus_MeasurementNotRequested struct {
}

func (MeasurementStatus_MeasurementNotRequested) isMeasurementStatus() {}

func (CompanionStruct_MeasurementStatus_) Create_MeasurementNotRequested_() MeasurementStatus {
	return MeasurementStatus{MeasurementStatus_MeasurementNotRequested{}}
}

func (_this MeasurementStatus) Is_MeasurementNotRequested() bool {
	_, ok := _this.Get_().(MeasurementStatus_MeasurementNotRequested)
	return ok
}

type MeasurementStatus_MeasurementPartial struct {
	Reason _dafny.Sequence
}

func (MeasurementStatus_MeasurementPartial) isMeasurementStatus() {}

func (CompanionStruct_MeasurementStatus_) Create_MeasurementPartial_(Reason _dafny.Sequence) MeasurementStatus {
	return MeasurementStatus{MeasurementStatus_MeasurementPartial{Reason}}
}

func (_this MeasurementStatus) Is_MeasurementPartial() bool {
	_, ok := _this.Get_().(MeasurementStatus_MeasurementPartial)
	return ok
}

type MeasurementStatus_MeasurementUnavailable struct {
	Reason _dafny.Sequence
}

func (MeasurementStatus_MeasurementUnavailable) isMeasurementStatus() {}

func (CompanionStruct_MeasurementStatus_) Create_MeasurementUnavailable_(Reason _dafny.Sequence) MeasurementStatus {
	return MeasurementStatus{MeasurementStatus_MeasurementUnavailable{Reason}}
}

func (_this MeasurementStatus) Is_MeasurementUnavailable() bool {
	_, ok := _this.Get_().(MeasurementStatus_MeasurementUnavailable)
	return ok
}

func (CompanionStruct_MeasurementStatus_) Default() MeasurementStatus {
	return Companion_MeasurementStatus_.Create_MeasurementComplete_()
}

func (_this MeasurementStatus) Dtor_reason() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case MeasurementStatus_MeasurementPartial:
		return data.Reason
	default:
		return data.(MeasurementStatus_MeasurementUnavailable).Reason
	}
}

func (_this MeasurementStatus) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case MeasurementStatus_MeasurementComplete:
		{
			return "ConfluxResourceCeilings.MeasurementStatus.MeasurementComplete"
		}
	case MeasurementStatus_MeasurementNotRequested:
		{
			return "ConfluxResourceCeilings.MeasurementStatus.MeasurementNotRequested"
		}
	case MeasurementStatus_MeasurementPartial:
		{
			return "ConfluxResourceCeilings.MeasurementStatus.MeasurementPartial" + "(" + data.Reason.VerbatimString(true) + ")"
		}
	case MeasurementStatus_MeasurementUnavailable:
		{
			return "ConfluxResourceCeilings.MeasurementStatus.MeasurementUnavailable" + "(" + data.Reason.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this MeasurementStatus) Equals(other MeasurementStatus) bool {
	switch data1 := _this.Get_().(type) {
	case MeasurementStatus_MeasurementComplete:
		{
			_, ok := other.Get_().(MeasurementStatus_MeasurementComplete)
			return ok
		}
	case MeasurementStatus_MeasurementNotRequested:
		{
			_, ok := other.Get_().(MeasurementStatus_MeasurementNotRequested)
			return ok
		}
	case MeasurementStatus_MeasurementPartial:
		{
			data2, ok := other.Get_().(MeasurementStatus_MeasurementPartial)
			return ok && data1.Reason.Equals(data2.Reason)
		}
	case MeasurementStatus_MeasurementUnavailable:
		{
			data2, ok := other.Get_().(MeasurementStatus_MeasurementUnavailable)
			return ok && data1.Reason.Equals(data2.Reason)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this MeasurementStatus) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(MeasurementStatus)
	return ok && _this.Equals(typed)
}

func Type_MeasurementStatus_() _dafny.TypeDescriptor {
	return type_MeasurementStatus_{}
}

type type_MeasurementStatus_ struct {
}

func (_this type_MeasurementStatus_) Default() interface{} {
	return Companion_MeasurementStatus_.Default()
}

func (_this type_MeasurementStatus_) String() string {
	return "ConfluxResourceCeilings.MeasurementStatus"
}
func (_this MeasurementStatus) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = MeasurementStatus{}

// End of datatype MeasurementStatus

// Definition of datatype ResourceMeasurement
type ResourceMeasurement struct {
	Data_ResourceMeasurement_
}

func (_this ResourceMeasurement) Get_() Data_ResourceMeasurement_ {
	return _this.Data_ResourceMeasurement_
}

type Data_ResourceMeasurement_ interface {
	isResourceMeasurement()
}

type CompanionStruct_ResourceMeasurement_ struct {
}

var Companion_ResourceMeasurement_ = CompanionStruct_ResourceMeasurement_{}

type ResourceMeasurement_ResourceMeasurement struct {
	Value  m_ConfluxContract.Option
	Status MeasurementStatus
}

func (ResourceMeasurement_ResourceMeasurement) isResourceMeasurement() {}

func (CompanionStruct_ResourceMeasurement_) Create_ResourceMeasurement_(Value m_ConfluxContract.Option, Status MeasurementStatus) ResourceMeasurement {
	return ResourceMeasurement{ResourceMeasurement_ResourceMeasurement{Value, Status}}
}

func (_this ResourceMeasurement) Is_ResourceMeasurement() bool {
	_, ok := _this.Get_().(ResourceMeasurement_ResourceMeasurement)
	return ok
}

func (CompanionStruct_ResourceMeasurement_) Default() ResourceMeasurement {
	return Companion_ResourceMeasurement_.Create_ResourceMeasurement_(m_ConfluxContract.Companion_Option_.Default(), Companion_MeasurementStatus_.Default())
}

func (_this ResourceMeasurement) Dtor_value() m_ConfluxContract.Option {
	return _this.Get_().(ResourceMeasurement_ResourceMeasurement).Value
}

func (_this ResourceMeasurement) Dtor_status() MeasurementStatus {
	return _this.Get_().(ResourceMeasurement_ResourceMeasurement).Status
}

func (_this ResourceMeasurement) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ResourceMeasurement_ResourceMeasurement:
		{
			return "ConfluxResourceCeilings.ResourceMeasurement.ResourceMeasurement" + "(" + _dafny.String(data.Value) + ", " + _dafny.String(data.Status) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ResourceMeasurement) Equals(other ResourceMeasurement) bool {
	switch data1 := _this.Get_().(type) {
	case ResourceMeasurement_ResourceMeasurement:
		{
			data2, ok := other.Get_().(ResourceMeasurement_ResourceMeasurement)
			return ok && data1.Value.Equals(data2.Value) && data1.Status.Equals(data2.Status)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ResourceMeasurement) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ResourceMeasurement)
	return ok && _this.Equals(typed)
}

func Type_ResourceMeasurement_() _dafny.TypeDescriptor {
	return type_ResourceMeasurement_{}
}

type type_ResourceMeasurement_ struct {
}

func (_this type_ResourceMeasurement_) Default() interface{} {
	return Companion_ResourceMeasurement_.Default()
}

func (_this type_ResourceMeasurement_) String() string {
	return "ConfluxResourceCeilings.ResourceMeasurement"
}
func (_this ResourceMeasurement) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ResourceMeasurement{}

// End of datatype ResourceMeasurement

// Definition of datatype ResourceObservation
type ResourceObservation struct {
	Data_ResourceObservation_
}

func (_this ResourceObservation) Get_() Data_ResourceObservation_ {
	return _this.Data_ResourceObservation_
}

type Data_ResourceObservation_ interface {
	isResourceObservation()
}

type CompanionStruct_ResourceObservation_ struct {
}

var Companion_ResourceObservation_ = CompanionStruct_ResourceObservation_{}

type ResourceObservation_ResourceObservation struct {
	WallMs              _dafny.Int
	PeakRssMb           ResourceMeasurement
	CombinedOutputBytes ResourceMeasurement
}

func (ResourceObservation_ResourceObservation) isResourceObservation() {}

func (CompanionStruct_ResourceObservation_) Create_ResourceObservation_(WallMs _dafny.Int, PeakRssMb ResourceMeasurement, CombinedOutputBytes ResourceMeasurement) ResourceObservation {
	return ResourceObservation{ResourceObservation_ResourceObservation{WallMs, PeakRssMb, CombinedOutputBytes}}
}

func (_this ResourceObservation) Is_ResourceObservation() bool {
	_, ok := _this.Get_().(ResourceObservation_ResourceObservation)
	return ok
}

func (CompanionStruct_ResourceObservation_) Default() ResourceObservation {
	return Companion_ResourceObservation_.Create_ResourceObservation_(_dafny.Zero, Companion_ResourceMeasurement_.Default(), Companion_ResourceMeasurement_.Default())
}

func (_this ResourceObservation) Dtor_wallMs() _dafny.Int {
	return _this.Get_().(ResourceObservation_ResourceObservation).WallMs
}

func (_this ResourceObservation) Dtor_peakRssMb() ResourceMeasurement {
	return _this.Get_().(ResourceObservation_ResourceObservation).PeakRssMb
}

func (_this ResourceObservation) Dtor_combinedOutputBytes() ResourceMeasurement {
	return _this.Get_().(ResourceObservation_ResourceObservation).CombinedOutputBytes
}

func (_this ResourceObservation) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ResourceObservation_ResourceObservation:
		{
			return "ConfluxResourceCeilings.ResourceObservation.ResourceObservation" + "(" + _dafny.String(data.WallMs) + ", " + _dafny.String(data.PeakRssMb) + ", " + _dafny.String(data.CombinedOutputBytes) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ResourceObservation) Equals(other ResourceObservation) bool {
	switch data1 := _this.Get_().(type) {
	case ResourceObservation_ResourceObservation:
		{
			data2, ok := other.Get_().(ResourceObservation_ResourceObservation)
			return ok && data1.WallMs.Cmp(data2.WallMs) == 0 && data1.PeakRssMb.Equals(data2.PeakRssMb) && data1.CombinedOutputBytes.Equals(data2.CombinedOutputBytes)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ResourceObservation) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ResourceObservation)
	return ok && _this.Equals(typed)
}

func Type_ResourceObservation_() _dafny.TypeDescriptor {
	return type_ResourceObservation_{}
}

type type_ResourceObservation_ struct {
}

func (_this type_ResourceObservation_) Default() interface{} {
	return Companion_ResourceObservation_.Default()
}

func (_this type_ResourceObservation_) String() string {
	return "ConfluxResourceCeilings.ResourceObservation"
}
func (_this ResourceObservation) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ResourceObservation{}

// End of datatype ResourceObservation

// Definition of datatype ResourceDimension
type ResourceDimension struct {
	Data_ResourceDimension_
}

func (_this ResourceDimension) Get_() Data_ResourceDimension_ {
	return _this.Data_ResourceDimension_
}

type Data_ResourceDimension_ interface {
	isResourceDimension()
}

type CompanionStruct_ResourceDimension_ struct {
}

var Companion_ResourceDimension_ = CompanionStruct_ResourceDimension_{}

type ResourceDimension_TimeDimension struct {
}

func (ResourceDimension_TimeDimension) isResourceDimension() {}

func (CompanionStruct_ResourceDimension_) Create_TimeDimension_() ResourceDimension {
	return ResourceDimension{ResourceDimension_TimeDimension{}}
}

func (_this ResourceDimension) Is_TimeDimension() bool {
	_, ok := _this.Get_().(ResourceDimension_TimeDimension)
	return ok
}

type ResourceDimension_MemoryDimension struct {
}

func (ResourceDimension_MemoryDimension) isResourceDimension() {}

func (CompanionStruct_ResourceDimension_) Create_MemoryDimension_() ResourceDimension {
	return ResourceDimension{ResourceDimension_MemoryDimension{}}
}

func (_this ResourceDimension) Is_MemoryDimension() bool {
	_, ok := _this.Get_().(ResourceDimension_MemoryDimension)
	return ok
}

type ResourceDimension_OutputDimension struct {
}

func (ResourceDimension_OutputDimension) isResourceDimension() {}

func (CompanionStruct_ResourceDimension_) Create_OutputDimension_() ResourceDimension {
	return ResourceDimension{ResourceDimension_OutputDimension{}}
}

func (_this ResourceDimension) Is_OutputDimension() bool {
	_, ok := _this.Get_().(ResourceDimension_OutputDimension)
	return ok
}

func (CompanionStruct_ResourceDimension_) Default() ResourceDimension {
	return Companion_ResourceDimension_.Create_TimeDimension_()
}

func (_ CompanionStruct_ResourceDimension_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_ResourceDimension_.Create_TimeDimension_(), true
		case 1:
			return Companion_ResourceDimension_.Create_MemoryDimension_(), true
		case 2:
			return Companion_ResourceDimension_.Create_OutputDimension_(), true
		default:
			return ResourceDimension{}, false
		}
	}
}

func (_this ResourceDimension) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case ResourceDimension_TimeDimension:
		{
			return "ConfluxResourceCeilings.ResourceDimension.TimeDimension"
		}
	case ResourceDimension_MemoryDimension:
		{
			return "ConfluxResourceCeilings.ResourceDimension.MemoryDimension"
		}
	case ResourceDimension_OutputDimension:
		{
			return "ConfluxResourceCeilings.ResourceDimension.OutputDimension"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ResourceDimension) Equals(other ResourceDimension) bool {
	switch _this.Get_().(type) {
	case ResourceDimension_TimeDimension:
		{
			_, ok := other.Get_().(ResourceDimension_TimeDimension)
			return ok
		}
	case ResourceDimension_MemoryDimension:
		{
			_, ok := other.Get_().(ResourceDimension_MemoryDimension)
			return ok
		}
	case ResourceDimension_OutputDimension:
		{
			_, ok := other.Get_().(ResourceDimension_OutputDimension)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ResourceDimension) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ResourceDimension)
	return ok && _this.Equals(typed)
}

func Type_ResourceDimension_() _dafny.TypeDescriptor {
	return type_ResourceDimension_{}
}

type type_ResourceDimension_ struct {
}

func (_this type_ResourceDimension_) Default() interface{} {
	return Companion_ResourceDimension_.Default()
}

func (_this type_ResourceDimension_) String() string {
	return "ConfluxResourceCeilings.ResourceDimension"
}
func (_this ResourceDimension) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ResourceDimension{}

// End of datatype ResourceDimension

// Definition of datatype EnforcementOutcome
type EnforcementOutcome struct {
	Data_EnforcementOutcome_
}

func (_this EnforcementOutcome) Get_() Data_EnforcementOutcome_ {
	return _this.Data_EnforcementOutcome_
}

type Data_EnforcementOutcome_ interface {
	isEnforcementOutcome()
}

type CompanionStruct_EnforcementOutcome_ struct {
}

var Companion_EnforcementOutcome_ = CompanionStruct_EnforcementOutcome_{}

type EnforcementOutcome_EnforcementOk struct {
}

func (EnforcementOutcome_EnforcementOk) isEnforcementOutcome() {}

func (CompanionStruct_EnforcementOutcome_) Create_EnforcementOk_() EnforcementOutcome {
	return EnforcementOutcome{EnforcementOutcome_EnforcementOk{}}
}

func (_this EnforcementOutcome) Is_EnforcementOk() bool {
	_, ok := _this.Get_().(EnforcementOutcome_EnforcementOk)
	return ok
}

type EnforcementOutcome_EnforcementFail struct {
}

func (EnforcementOutcome_EnforcementFail) isEnforcementOutcome() {}

func (CompanionStruct_EnforcementOutcome_) Create_EnforcementFail_() EnforcementOutcome {
	return EnforcementOutcome{EnforcementOutcome_EnforcementFail{}}
}

func (_this EnforcementOutcome) Is_EnforcementFail() bool {
	_, ok := _this.Get_().(EnforcementOutcome_EnforcementFail)
	return ok
}

type EnforcementOutcome_EnforcementKill struct {
}

func (EnforcementOutcome_EnforcementKill) isEnforcementOutcome() {}

func (CompanionStruct_EnforcementOutcome_) Create_EnforcementKill_() EnforcementOutcome {
	return EnforcementOutcome{EnforcementOutcome_EnforcementKill{}}
}

func (_this EnforcementOutcome) Is_EnforcementKill() bool {
	_, ok := _this.Get_().(EnforcementOutcome_EnforcementKill)
	return ok
}

func (CompanionStruct_EnforcementOutcome_) Default() EnforcementOutcome {
	return Companion_EnforcementOutcome_.Create_EnforcementOk_()
}

func (_ CompanionStruct_EnforcementOutcome_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_EnforcementOutcome_.Create_EnforcementOk_(), true
		case 1:
			return Companion_EnforcementOutcome_.Create_EnforcementFail_(), true
		case 2:
			return Companion_EnforcementOutcome_.Create_EnforcementKill_(), true
		default:
			return EnforcementOutcome{}, false
		}
	}
}

func (_this EnforcementOutcome) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case EnforcementOutcome_EnforcementOk:
		{
			return "ConfluxResourceCeilings.EnforcementOutcome.EnforcementOk"
		}
	case EnforcementOutcome_EnforcementFail:
		{
			return "ConfluxResourceCeilings.EnforcementOutcome.EnforcementFail"
		}
	case EnforcementOutcome_EnforcementKill:
		{
			return "ConfluxResourceCeilings.EnforcementOutcome.EnforcementKill"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this EnforcementOutcome) Equals(other EnforcementOutcome) bool {
	switch _this.Get_().(type) {
	case EnforcementOutcome_EnforcementOk:
		{
			_, ok := other.Get_().(EnforcementOutcome_EnforcementOk)
			return ok
		}
	case EnforcementOutcome_EnforcementFail:
		{
			_, ok := other.Get_().(EnforcementOutcome_EnforcementFail)
			return ok
		}
	case EnforcementOutcome_EnforcementKill:
		{
			_, ok := other.Get_().(EnforcementOutcome_EnforcementKill)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this EnforcementOutcome) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(EnforcementOutcome)
	return ok && _this.Equals(typed)
}

func Type_EnforcementOutcome_() _dafny.TypeDescriptor {
	return type_EnforcementOutcome_{}
}

type type_EnforcementOutcome_ struct {
}

func (_this type_EnforcementOutcome_) Default() interface{} {
	return Companion_EnforcementOutcome_.Default()
}

func (_this type_EnforcementOutcome_) String() string {
	return "ConfluxResourceCeilings.EnforcementOutcome"
}
func (_this EnforcementOutcome) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = EnforcementOutcome{}

// End of datatype EnforcementOutcome

// Definition of datatype ResourceVerdict
type ResourceVerdict struct {
	Data_ResourceVerdict_
}

func (_this ResourceVerdict) Get_() Data_ResourceVerdict_ {
	return _this.Data_ResourceVerdict_
}

type Data_ResourceVerdict_ interface {
	isResourceVerdict()
}

type CompanionStruct_ResourceVerdict_ struct {
}

var Companion_ResourceVerdict_ = CompanionStruct_ResourceVerdict_{}

type ResourceVerdict_ResourceVerdict struct {
	Outcome   EnforcementOutcome
	Dimension m_ConfluxContract.Option
	Phase     ExecutionPhase
	Ceiling   m_ConfluxContract.Option
	Observed  m_ConfluxContract.Option
	Reason    _dafny.Sequence
}

func (ResourceVerdict_ResourceVerdict) isResourceVerdict() {}

func (CompanionStruct_ResourceVerdict_) Create_ResourceVerdict_(Outcome EnforcementOutcome, Dimension m_ConfluxContract.Option, Phase ExecutionPhase, Ceiling m_ConfluxContract.Option, Observed m_ConfluxContract.Option, Reason _dafny.Sequence) ResourceVerdict {
	return ResourceVerdict{ResourceVerdict_ResourceVerdict{Outcome, Dimension, Phase, Ceiling, Observed, Reason}}
}

func (_this ResourceVerdict) Is_ResourceVerdict() bool {
	_, ok := _this.Get_().(ResourceVerdict_ResourceVerdict)
	return ok
}

func (CompanionStruct_ResourceVerdict_) Default() ResourceVerdict {
	return Companion_ResourceVerdict_.Create_ResourceVerdict_(Companion_EnforcementOutcome_.Default(), m_ConfluxContract.Companion_Option_.Default(), Companion_ExecutionPhase_.Default(), m_ConfluxContract.Companion_Option_.Default(), m_ConfluxContract.Companion_Option_.Default(), _dafny.EmptySeq)
}

func (_this ResourceVerdict) Dtor_outcome() EnforcementOutcome {
	return _this.Get_().(ResourceVerdict_ResourceVerdict).Outcome
}

func (_this ResourceVerdict) Dtor_dimension() m_ConfluxContract.Option {
	return _this.Get_().(ResourceVerdict_ResourceVerdict).Dimension
}

func (_this ResourceVerdict) Dtor_phase() ExecutionPhase {
	return _this.Get_().(ResourceVerdict_ResourceVerdict).Phase
}

func (_this ResourceVerdict) Dtor_ceiling() m_ConfluxContract.Option {
	return _this.Get_().(ResourceVerdict_ResourceVerdict).Ceiling
}

func (_this ResourceVerdict) Dtor_observed() m_ConfluxContract.Option {
	return _this.Get_().(ResourceVerdict_ResourceVerdict).Observed
}

func (_this ResourceVerdict) Dtor_reason() _dafny.Sequence {
	return _this.Get_().(ResourceVerdict_ResourceVerdict).Reason
}

func (_this ResourceVerdict) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ResourceVerdict_ResourceVerdict:
		{
			return "ConfluxResourceCeilings.ResourceVerdict.ResourceVerdict" + "(" + _dafny.String(data.Outcome) + ", " + _dafny.String(data.Dimension) + ", " + _dafny.String(data.Phase) + ", " + _dafny.String(data.Ceiling) + ", " + _dafny.String(data.Observed) + ", " + data.Reason.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ResourceVerdict) Equals(other ResourceVerdict) bool {
	switch data1 := _this.Get_().(type) {
	case ResourceVerdict_ResourceVerdict:
		{
			data2, ok := other.Get_().(ResourceVerdict_ResourceVerdict)
			return ok && data1.Outcome.Equals(data2.Outcome) && data1.Dimension.Equals(data2.Dimension) && data1.Phase.Equals(data2.Phase) && data1.Ceiling.Equals(data2.Ceiling) && data1.Observed.Equals(data2.Observed) && data1.Reason.Equals(data2.Reason)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ResourceVerdict) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ResourceVerdict)
	return ok && _this.Equals(typed)
}

func Type_ResourceVerdict_() _dafny.TypeDescriptor {
	return type_ResourceVerdict_{}
}

type type_ResourceVerdict_ struct {
}

func (_this type_ResourceVerdict_) Default() interface{} {
	return Companion_ResourceVerdict_.Default()
}

func (_this type_ResourceVerdict_) String() string {
	return "ConfluxResourceCeilings.ResourceVerdict"
}
func (_this ResourceVerdict) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ResourceVerdict{}

// End of datatype ResourceVerdict

// Definition of datatype CacheHostPhaseState
type CacheHostPhaseState struct {
	Data_CacheHostPhaseState_
}

func (_this CacheHostPhaseState) Get_() Data_CacheHostPhaseState_ {
	return _this.Data_CacheHostPhaseState_
}

type Data_CacheHostPhaseState_ interface {
	isCacheHostPhaseState()
}

type CompanionStruct_CacheHostPhaseState_ struct {
}

var Companion_CacheHostPhaseState_ = CompanionStruct_CacheHostPhaseState_{}

type CacheHostPhaseState_HostPhaseUnavailable struct {
}

func (CacheHostPhaseState_HostPhaseUnavailable) isCacheHostPhaseState() {}

func (CompanionStruct_CacheHostPhaseState_) Create_HostPhaseUnavailable_() CacheHostPhaseState {
	return CacheHostPhaseState{CacheHostPhaseState_HostPhaseUnavailable{}}
}

func (_this CacheHostPhaseState) Is_HostPhaseUnavailable() bool {
	_, ok := _this.Get_().(CacheHostPhaseState_HostPhaseUnavailable)
	return ok
}

type CacheHostPhaseState_HostPhaseKnown struct {
	HostGeneration   _dafny.Int
	CachedGeneration m_ConfluxContract.Option
}

func (CacheHostPhaseState_HostPhaseKnown) isCacheHostPhaseState() {}

func (CompanionStruct_CacheHostPhaseState_) Create_HostPhaseKnown_(HostGeneration _dafny.Int, CachedGeneration m_ConfluxContract.Option) CacheHostPhaseState {
	return CacheHostPhaseState{CacheHostPhaseState_HostPhaseKnown{HostGeneration, CachedGeneration}}
}

func (_this CacheHostPhaseState) Is_HostPhaseKnown() bool {
	_, ok := _this.Get_().(CacheHostPhaseState_HostPhaseKnown)
	return ok
}

func (CompanionStruct_CacheHostPhaseState_) Default() CacheHostPhaseState {
	return Companion_CacheHostPhaseState_.Create_HostPhaseUnavailable_()
}

func (_this CacheHostPhaseState) Dtor_hostGeneration() _dafny.Int {
	return _this.Get_().(CacheHostPhaseState_HostPhaseKnown).HostGeneration
}

func (_this CacheHostPhaseState) Dtor_cachedGeneration() m_ConfluxContract.Option {
	return _this.Get_().(CacheHostPhaseState_HostPhaseKnown).CachedGeneration
}

func (_this CacheHostPhaseState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CacheHostPhaseState_HostPhaseUnavailable:
		{
			return "ConfluxResourceCeilings.CacheHostPhaseState.HostPhaseUnavailable"
		}
	case CacheHostPhaseState_HostPhaseKnown:
		{
			return "ConfluxResourceCeilings.CacheHostPhaseState.HostPhaseKnown" + "(" + _dafny.String(data.HostGeneration) + ", " + _dafny.String(data.CachedGeneration) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CacheHostPhaseState) Equals(other CacheHostPhaseState) bool {
	switch data1 := _this.Get_().(type) {
	case CacheHostPhaseState_HostPhaseUnavailable:
		{
			_, ok := other.Get_().(CacheHostPhaseState_HostPhaseUnavailable)
			return ok
		}
	case CacheHostPhaseState_HostPhaseKnown:
		{
			data2, ok := other.Get_().(CacheHostPhaseState_HostPhaseKnown)
			return ok && data1.HostGeneration.Cmp(data2.HostGeneration) == 0 && data1.CachedGeneration.Equals(data2.CachedGeneration)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CacheHostPhaseState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CacheHostPhaseState)
	return ok && _this.Equals(typed)
}

func Type_CacheHostPhaseState_() _dafny.TypeDescriptor {
	return type_CacheHostPhaseState_{}
}

type type_CacheHostPhaseState_ struct {
}

func (_this type_CacheHostPhaseState_) Default() interface{} {
	return Companion_CacheHostPhaseState_.Default()
}

func (_this type_CacheHostPhaseState_) String() string {
	return "ConfluxResourceCeilings.CacheHostPhaseState"
}
func (_this CacheHostPhaseState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CacheHostPhaseState{}

// End of datatype CacheHostPhaseState
