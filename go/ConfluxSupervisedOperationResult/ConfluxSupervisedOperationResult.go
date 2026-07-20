// Package ConfluxSupervisedOperationResult
// Dafny module ConfluxSupervisedOperationResult compiled into Go

package ConfluxSupervisedOperationResult

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-go/Changeset"
	m_Chat "github.com/joshmouch/ahp-go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-go/ClientMain"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-go/ConfluxOrderedLog"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-go/ConfluxResourceCeilings"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-go/ConfluxSeqRoute"
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
	return "ConfluxSupervisedOperationResult.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) V1WireNat(value _dafny.Int) bool {
	return (value).Cmp(Companion_Default___.V1MaxSafeInteger()) <= 0
}
func (_static *CompanionStruct_Default___) V1WireInt(value _dafny.Int) bool {
	return ((Companion_Default___.V1MinSafeInteger()).Cmp(value) <= 0) && ((value).Cmp(Companion_Default___.V1MaxSafeInteger()) <= 0)
}
func (_static *CompanionStruct_Default___) V1WireOptionalNat(value m_ConfluxContract.Option) bool {
	return ((value).Is_None()) || (Companion_Default___.V1WireNat((value).Dtor_value().(_dafny.Int)))
}
func (_static *CompanionStruct_Default___) V1BudgetSafeForPhase(budget m_ConfluxResourceCeilings.ResourceCeilings, phase m_ConfluxResourceCeilings.ExecutionPhase) bool {
	return ((Companion_Default___.V1WireNat(m_ConfluxResourceCeilings.Companion_Default___.TimeCeilingForPhase((budget).Dtor_time(), phase))) && (Companion_Default___.V1WireNat(m_ConfluxResourceCeilings.Companion_Default___.MemoryCeilingForPhase((budget).Dtor_memory(), phase)))) && (Companion_Default___.V1WireNat((budget).Dtor_maxOutputMb()))
}
func (_static *CompanionStruct_Default___) InclusiveWallMs(timing TimingCoverage) _dafny.Int {
	if ((timing).Dtor_completedMs()).Cmp((timing).Dtor_startedMs()) >= 0 {
		return ((timing).Dtor_completedMs()).Minus((timing).Dtor_startedMs())
	} else {
		return _dafny.Zero
	}
}
func (_static *CompanionStruct_Default___) ClampCoverage(inclusiveMs _dafny.Int, unionMs _dafny.Int) _dafny.Int {
	if (unionMs).Cmp(inclusiveMs) <= 0 {
		return unionMs
	} else {
		return inclusiveMs
	}
}
func (_static *CompanionStruct_Default___) ExclusiveWallMs(timing TimingCoverage) _dafny.Int {
	return (Companion_Default___.InclusiveWallMs(timing)).Minus(Companion_Default___.ClampCoverage(Companion_Default___.InclusiveWallMs(timing), (timing).Dtor_childUnionMs()))
}
func (_static *CompanionStruct_Default___) GapWallMs(timing TimingCoverage) _dafny.Int {
	return (Companion_Default___.InclusiveWallMs(timing)).Minus(Companion_Default___.ClampCoverage(Companion_Default___.InclusiveWallMs(timing), (timing).Dtor_observedUnionMs()))
}
func (_static *CompanionStruct_Default___) ExpandBudget(budget m_ConfluxResourceCeilings.ResourceCeilings, requestedPhase m_ConfluxResourceCeilings.ExecutionPhase) ExpandedResourceBudgetV1 {
	return Companion_ExpandedResourceBudgetV1_.Create_ExpandedResourceBudgetV1_(m_ConfluxResourceCeilings.Companion_Default___.TimeCeilingForPhase((budget).Dtor_time(), requestedPhase), m_ConfluxResourceCeilings.Companion_Default___.MemoryCeilingForPhase((budget).Dtor_memory(), requestedPhase), (budget).Dtor_maxOutputMb())
}
func (_static *CompanionStruct_Default___) StatusKind(status m_ConfluxResourceCeilings.MeasurementStatus) _dafny.Int {
	if (status).Is_MeasurementComplete() {
		return _dafny.Zero
	} else if (status).Is_MeasurementPartial() {
		return _dafny.One
	} else if (status).Is_MeasurementUnavailable() {
		return _dafny.IntOfInt64(2)
	} else {
		return _dafny.IntOfInt64(3)
	}
}
func (_static *CompanionStruct_Default___) DerivedObservationStatus(observation SupervisionObservationV1) _dafny.Int {
	if (((((observation).Dtor_peakDescendantRssMb()).Dtor_status()).Is_MeasurementComplete()) && ((((observation).Dtor_combinedOutputBytes()).Dtor_status()).Is_MeasurementComplete())) && (((observation).Dtor_processTree()).Is_ProcessTreeObserved()) {
		return _dafny.Zero
	} else {
		return _dafny.One
	}
}
func (_static *CompanionStruct_Default___) VerdictShapeValid(verdict m_ConfluxResourceCeilings.ResourceVerdict) bool {
	return (!(((verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementOk_())) || (((((verdict).Dtor_dimension()).Is_None()) && (((verdict).Dtor_ceiling()).Is_None())) && (((verdict).Dtor_observed()).Is_None()))) && (!(!((verdict).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementOk_())) || (((((verdict).Dtor_dimension()).Is_Some()) && (((verdict).Dtor_ceiling()).Is_Some())) && ((_dafny.IntOfUint32(((verdict).Dtor_reason()).Cardinality())).Sign() == 1)))
}
func (_static *CompanionStruct_Default___) FirstBreachCoherent(result SupervisedOperationResultV1) bool {
	return (((((result).Dtor_verdict()).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementOk_())) == (((result).Dtor_supervisionState()).Is_SupervisionRunning())) && (!(((result).Dtor_supervisionState()).Is_SupervisionTripped()) || ((((((((result).Dtor_supervisionState()).Dtor_record()).Dtor_verdict()).Equals((result).Dtor_verdict())) && (!(((((result).Dtor_supervisionState()).Dtor_record()).Dtor_verdict()).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementOk_()))) && (!((((result).Dtor_supervisionState()).Dtor_record()).Dtor_measurementFailure()) || ((((((result).Dtor_verdict()).Dtor_dimension()).Is_Some()) && ((((((result).Dtor_verdict()).Dtor_dimension()).Dtor_value().(m_ConfluxResourceCeilings.ResourceDimension)).Equals(m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_MemoryDimension_())) || (((((result).Dtor_verdict()).Dtor_dimension()).Dtor_value().(m_ConfluxResourceCeilings.ResourceDimension)).Equals(m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_OutputDimension_())))) && ((((result).Dtor_verdict()).Dtor_observed()).Is_None())))) && (!(!((((result).Dtor_supervisionState()).Dtor_record()).Dtor_measurementFailure())) || ((((result).Dtor_verdict()).Dtor_observed()).Is_Some()))))
}
func (_static *CompanionStruct_Default___) RequiredMeasurementsFailClosed(result SupervisedOperationResultV1) bool {
	return (!(((((result).Dtor_expandedCeilings()).Dtor_peakRssMb()).Sign() == 1) && (!(((((result).Dtor_observation()).Dtor_peakDescendantRssMb()).Dtor_status()).Is_MeasurementComplete()))) || ((!(((result).Dtor_verdict()).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementOk_())) && (((result).Dtor_wrapperExitCode()).Sign() != 0))) && (!(((((result).Dtor_expandedCeilings()).Dtor_combinedOutputMb()).Sign() == 1) && (!(((((result).Dtor_observation()).Dtor_combinedOutputBytes()).Dtor_status()).Is_MeasurementComplete()))) || ((!(((result).Dtor_verdict()).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementOk_())) && (((result).Dtor_wrapperExitCode()).Sign() != 0)))
}
func (_static *CompanionStruct_Default___) CleanupValid(cleanup CleanupStatus) bool {
	return (((cleanup).Is_CleanupNotRequired()) || ((((cleanup).Is_CleanupComplete()) && (((cleanup).Dtor_deadlineMs()).Sign() == 1)) && (((cleanup).Dtor_elapsedMs()).Cmp((cleanup).Dtor_deadlineMs()) <= 0))) || (((((cleanup).Is_CleanupFailed()) && (((cleanup).Dtor_deadlineMs()).Sign() == 1)) && (((cleanup).Dtor_elapsedMs()).Cmp((cleanup).Dtor_deadlineMs()) <= 0)) && ((_dafny.IntOfUint32(((cleanup).Dtor_reason()).Cardinality())).Sign() == 1))
}
func (_static *CompanionStruct_Default___) MeasurementStatusValid(status m_ConfluxResourceCeilings.MeasurementStatus) bool {
	return ((!((status).Is_MeasurementPartial())) && (!((status).Is_MeasurementUnavailable()))) || ((_dafny.IntOfUint32(((status).Dtor_reason()).Cardinality())).Sign() == 1)
}
func (_static *CompanionStruct_Default___) CleanupCoherent(result SupervisedOperationResultV1) bool {
	if ((result).Dtor_cleanup()).Is_CleanupNotRequired() {
		return ((result).Dtor_termination()).Equals(Companion_TerminationScope_.Create_NoTermination_())
	} else {
		return (((((result).Dtor_termination()).Equals(Companion_TerminationScope_.Create_ProcessTreeTermination_())) && ((((result).Dtor_cleanup()).Dtor_deadlineMs()).Sign() == 1)) && ((((result).Dtor_cleanup()).Dtor_elapsedMs()).Cmp(((result).Dtor_cleanup()).Dtor_deadlineMs()) <= 0)) && ((((result).Dtor_cleanup()).Is_CleanupComplete()) || ((((result).Dtor_cleanup()).Is_CleanupFailed()) && ((_dafny.IntOfUint32((((result).Dtor_cleanup()).Dtor_reason()).Cardinality())).Sign() == 1)))
	}
}
func (_static *CompanionStruct_Default___) ExpandedLimit(result SupervisedOperationResultV1, dimension m_ConfluxResourceCeilings.ResourceDimension) _dafny.Int {
	if (dimension).Equals(m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_TimeDimension_()) {
		return ((result).Dtor_expandedCeilings()).Dtor_timeMs()
	} else if (dimension).Equals(m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_MemoryDimension_()) {
		return ((result).Dtor_expandedCeilings()).Dtor_peakRssMb()
	} else {
		return ((result).Dtor_expandedCeilings()).Dtor_combinedOutputMb()
	}
}
func (_static *CompanionStruct_Default___) NumericBreachReflected(result SupervisedOperationResultV1, dimension m_ConfluxResourceCeilings.ResourceDimension, observed _dafny.Int) bool {
	if (dimension).Equals(m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_TimeDimension_()) {
		return (observed).Cmp(Companion_Default___.InclusiveWallMs(((result).Dtor_observation()).Dtor_timing())) <= 0
	} else if (dimension).Equals(m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_MemoryDimension_()) {
		return (((((result).Dtor_observation()).Dtor_peakDescendantRssMb()).Dtor_value()).Is_Some()) && ((observed).Cmp(((((result).Dtor_observation()).Dtor_peakDescendantRssMb()).Dtor_value()).Dtor_value().(_dafny.Int)) <= 0)
	} else {
		return (((((result).Dtor_observation()).Dtor_combinedOutputBytes()).Dtor_value()).Is_Some()) && ((observed).Cmp(m_ConfluxResourceCeilings.Companion_Default___.BytesToMegabytes(((((result).Dtor_observation()).Dtor_combinedOutputBytes()).Dtor_value()).Dtor_value().(_dafny.Int))) <= 0)
	}
}
func (_static *CompanionStruct_Default___) VerdictBudgetCoherent(result SupervisedOperationResultV1) bool {
	if (((result).Dtor_verdict()).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementOk_()) {
		return ((((((result).Dtor_expandedCeilings()).Dtor_timeMs()).Sign() == 0) || ((Companion_Default___.InclusiveWallMs(((result).Dtor_observation()).Dtor_timing())).Cmp(((result).Dtor_expandedCeilings()).Dtor_timeMs()) <= 0)) && (((((result).Dtor_expandedCeilings()).Dtor_peakRssMb()).Sign() == 0) || (((((((result).Dtor_observation()).Dtor_peakDescendantRssMb()).Dtor_status()).Is_MeasurementComplete()) && (((((result).Dtor_observation()).Dtor_peakDescendantRssMb()).Dtor_value()).Is_Some())) && ((((((result).Dtor_observation()).Dtor_peakDescendantRssMb()).Dtor_value()).Dtor_value().(_dafny.Int)).Cmp(((result).Dtor_expandedCeilings()).Dtor_peakRssMb()) <= 0)))) && (((((result).Dtor_expandedCeilings()).Dtor_combinedOutputMb()).Sign() == 0) || (((((((result).Dtor_observation()).Dtor_combinedOutputBytes()).Dtor_status()).Is_MeasurementComplete()) && (((((result).Dtor_observation()).Dtor_combinedOutputBytes()).Dtor_value()).Is_Some())) && ((m_ConfluxResourceCeilings.Companion_Default___.BytesToMegabytes(((((result).Dtor_observation()).Dtor_combinedOutputBytes()).Dtor_value()).Dtor_value().(_dafny.Int))).Cmp(((result).Dtor_expandedCeilings()).Dtor_combinedOutputMb()) <= 0)))
	} else {
		return (((((((result).Dtor_supervisionState()).Is_SupervisionTripped()) && ((((result).Dtor_verdict()).Dtor_dimension()).Is_Some())) && ((((result).Dtor_verdict()).Dtor_ceiling()).Is_Some())) && (((((result).Dtor_verdict()).Dtor_ceiling()).Dtor_value().(_dafny.Int)).Cmp(Companion_Default___.ExpandedLimit(result, (((result).Dtor_verdict()).Dtor_dimension()).Dtor_value().(m_ConfluxResourceCeilings.ResourceDimension))) == 0)) && (((((result).Dtor_verdict()).Dtor_ceiling()).Dtor_value().(_dafny.Int)).Sign() == 1)) && ((func() bool {
			if (((result).Dtor_supervisionState()).Dtor_record()).Dtor_measurementFailure() {
				return ((((result).Dtor_verdict()).Dtor_observed()).Is_None()) && ((func() bool {
					if ((((result).Dtor_verdict()).Dtor_dimension()).Dtor_value().(m_ConfluxResourceCeilings.ResourceDimension)).Equals(m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_MemoryDimension_()) {
						return !(((((result).Dtor_observation()).Dtor_peakDescendantRssMb()).Dtor_status()).Is_MeasurementComplete())
					}
					return (((((result).Dtor_verdict()).Dtor_dimension()).Dtor_value().(m_ConfluxResourceCeilings.ResourceDimension)).Equals(m_ConfluxResourceCeilings.Companion_ResourceDimension_.Create_OutputDimension_())) && (!(((((result).Dtor_observation()).Dtor_combinedOutputBytes()).Dtor_status()).Is_MeasurementComplete()))
				})())
			}
			return (((((result).Dtor_verdict()).Dtor_observed()).Is_Some()) && (((((result).Dtor_verdict()).Dtor_observed()).Dtor_value().(_dafny.Int)).Cmp((((result).Dtor_verdict()).Dtor_ceiling()).Dtor_value().(_dafny.Int)) > 0)) && (Companion_Default___.NumericBreachReflected(result, (((result).Dtor_verdict()).Dtor_dimension()).Dtor_value().(m_ConfluxResourceCeilings.ResourceDimension), (((result).Dtor_verdict()).Dtor_observed()).Dtor_value().(_dafny.Int)))
		})())
	}
}
func (_static *CompanionStruct_Default___) V1WireSafe(result SupervisedOperationResultV1) bool {
	return (((((((((((((((Companion_Default___.V1WireNat((result).Dtor_schemaVersion())) && (Companion_Default___.V1WireNat(((result).Dtor_expandedCeilings()).Dtor_timeMs()))) && (Companion_Default___.V1WireNat(((result).Dtor_expandedCeilings()).Dtor_peakRssMb()))) && (Companion_Default___.V1WireNat(((result).Dtor_expandedCeilings()).Dtor_combinedOutputMb()))) && (Companion_Default___.V1WireNat((((result).Dtor_observation()).Dtor_timing()).Dtor_startedMs()))) && (Companion_Default___.V1WireNat((((result).Dtor_observation()).Dtor_timing()).Dtor_completedMs()))) && (Companion_Default___.V1WireNat((((result).Dtor_observation()).Dtor_timing()).Dtor_childUnionMs()))) && (Companion_Default___.V1WireNat((((result).Dtor_observation()).Dtor_timing()).Dtor_observedUnionMs()))) && (Companion_Default___.V1WireOptionalNat((((result).Dtor_observation()).Dtor_peakDescendantRssMb()).Dtor_value()))) && (Companion_Default___.V1WireOptionalNat((((result).Dtor_observation()).Dtor_combinedOutputBytes()).Dtor_value()))) && (Companion_Default___.V1WireOptionalNat(((result).Dtor_verdict()).Dtor_ceiling()))) && (Companion_Default___.V1WireOptionalNat(((result).Dtor_verdict()).Dtor_observed()))) && ((func() bool {
		if ((result).Dtor_supervisionState()).Is_SupervisionRunning() {
			return Companion_Default___.V1WireNat(((result).Dtor_supervisionState()).Dtor_samplesSeen())
		}
		return ((((result).Dtor_supervisionState()).Dtor_record()).Dtor_sampleIndex()).Cmp(Companion_Default___.V1MaxSafeInteger()) < 0
	})())) && ((func() bool {
		if ((result).Dtor_cleanup()).Is_CleanupNotRequired() {
			return true
		}
		return (Companion_Default___.V1WireNat(((result).Dtor_cleanup()).Dtor_elapsedMs())) && (Companion_Default___.V1WireNat(((result).Dtor_cleanup()).Dtor_deadlineMs()))
	})())) && ((((result).Dtor_exit()).Is_ExitNotObserved()) || (Companion_Default___.V1WireInt(((result).Dtor_exit()).Dtor_code())))) && (Companion_Default___.V1WireInt((result).Dtor_wrapperExitCode()))
}
func (_static *CompanionStruct_Default___) ValidV1(result SupervisedOperationResultV1) bool {
	return (((((((((((((((((((((((result).Dtor_schemaVersion()).Cmp(Companion_Default___.SupervisedOperationResultSchemaVersion()) == 0) && (Companion_Default___.V1WireSafe(result))) && ((_dafny.IntOfUint32(((result).Dtor_policyId()).Cardinality())).Sign() == 1)) && (((result).Dtor_effectivePhase()).Equals(m_ConfluxResourceCeilings.Companion_Default___.EffectivePhase((result).Dtor_requestedPhase())))) && (((((result).Dtor_observation()).Dtor_timing()).Dtor_completedMs()).Cmp((((result).Dtor_observation()).Dtor_timing()).Dtor_startedMs()) >= 0)) && (m_ConfluxResourceCeilings.Companion_Default___.ValidMeasurement(((result).Dtor_observation()).Dtor_peakDescendantRssMb()))) && (m_ConfluxResourceCeilings.Companion_Default___.ValidMeasurement(((result).Dtor_observation()).Dtor_combinedOutputBytes()))) && (Companion_Default___.MeasurementStatusValid((result).Dtor_observationStatus()))) && ((Companion_Default___.StatusKind((result).Dtor_observationStatus())).Cmp(Companion_Default___.DerivedObservationStatus((result).Dtor_observation())) == 0)) && (!((((result).Dtor_observation()).Dtor_processTree()).Is_ProcessTreeObserved()) || (((((result).Dtor_observation()).Dtor_peakDescendantRssMb()).Dtor_status()).Is_MeasurementComplete()))) && (!((((result).Dtor_observation()).Dtor_processTree()).Is_ProcessTreeUnavailable()) || ((!(((((result).Dtor_observation()).Dtor_peakDescendantRssMb()).Dtor_status()).Is_MeasurementComplete())) && ((_dafny.IntOfUint32(((((result).Dtor_observation()).Dtor_processTree()).Dtor_reason()).Cardinality())).Sign() == 1)))) && ((((result).Dtor_verdict()).Dtor_phase()).Equals((result).Dtor_effectivePhase()))) && (Companion_Default___.VerdictShapeValid((result).Dtor_verdict()))) && (_dafny.Companion_Sequence_.Equal((result).Dtor_reason(), ((result).Dtor_verdict()).Dtor_reason()))) && (Companion_Default___.FirstBreachCoherent(result))) && (Companion_Default___.RequiredMeasurementsFailClosed(result))) && (Companion_Default___.VerdictBudgetCoherent(result))) && (Companion_Default___.CleanupValid((result).Dtor_cleanup()))) && (!((((result).Dtor_verdict()).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementOk_())) || (((((result).Dtor_enforcement()).Equals(Companion_EnforcementStatus_.Create_NoEnforcement_())) && (((result).Dtor_termination()).Equals(Companion_TerminationScope_.Create_NoTermination_()))) && (!(((result).Dtor_exit()).Is_ExitObserved()) || (((result).Dtor_wrapperExitCode()).Cmp(((result).Dtor_exit()).Dtor_code()) == 0))))) && (!((((result).Dtor_verdict()).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementFail_())) || (((((result).Dtor_enforcement()).Equals(Companion_EnforcementStatus_.Create_PostHocFailure_())) && (((result).Dtor_termination()).Equals(Companion_TerminationScope_.Create_NoTermination_()))) && (((result).Dtor_wrapperExitCode()).Sign() != 0)))) && (!((((result).Dtor_verdict()).Dtor_outcome()).Equals(m_ConfluxResourceCeilings.Companion_EnforcementOutcome_.Create_EnforcementKill_())) || (((((result).Dtor_enforcement()).Equals(Companion_EnforcementStatus_.Create_InFlightProcessTreeTermination_())) && (((result).Dtor_termination()).Equals(Companion_TerminationScope_.Create_ProcessTreeTermination_()))) && (((result).Dtor_wrapperExitCode()).Sign() != 0)))) && (Companion_Default___.CleanupCoherent(result))
}
func (_static *CompanionStruct_Default___) V1ReaderAccepts(schemaVersion _dafny.Int, requiredV1FieldsPresent bool, requiredV1MeaningsUnchanged bool) bool {
	return (((schemaVersion).Cmp(Companion_Default___.SupervisedOperationResultSchemaVersion()) == 0) && (requiredV1FieldsPresent)) && (requiredV1MeaningsUnchanged)
}
func (_static *CompanionStruct_Default___) PreserveProviderOutcome(provider m_ConfluxResourceCeilings.EnforcementOutcome, consumerCandidate m_ConfluxResourceCeilings.EnforcementOutcome) m_ConfluxResourceCeilings.EnforcementOutcome {
	return m_ConfluxResourceCeilings.Companion_Default___.DominantOutcome(provider, consumerCandidate)
}
func (_static *CompanionStruct_Default___) V1MaxSafeInteger() _dafny.Int {
	return _dafny.IntOfInt64(9007199254740991)
}
func (_static *CompanionStruct_Default___) V1MinSafeInteger() _dafny.Int {
	return _dafny.IntOfInt64(-9007199254740991)
}
func (_static *CompanionStruct_Default___) SupervisedOperationResultSchemaVersion() _dafny.Int {
	return _dafny.One
}

// End of class Default__

// Definition of datatype TimingCoverage
type TimingCoverage struct {
	Data_TimingCoverage_
}

func (_this TimingCoverage) Get_() Data_TimingCoverage_ {
	return _this.Data_TimingCoverage_
}

type Data_TimingCoverage_ interface {
	isTimingCoverage()
}

type CompanionStruct_TimingCoverage_ struct {
}

var Companion_TimingCoverage_ = CompanionStruct_TimingCoverage_{}

type TimingCoverage_TimingCoverage struct {
	StartedMs       _dafny.Int
	CompletedMs     _dafny.Int
	ChildUnionMs    _dafny.Int
	ObservedUnionMs _dafny.Int
}

func (TimingCoverage_TimingCoverage) isTimingCoverage() {}

func (CompanionStruct_TimingCoverage_) Create_TimingCoverage_(StartedMs _dafny.Int, CompletedMs _dafny.Int, ChildUnionMs _dafny.Int, ObservedUnionMs _dafny.Int) TimingCoverage {
	return TimingCoverage{TimingCoverage_TimingCoverage{StartedMs, CompletedMs, ChildUnionMs, ObservedUnionMs}}
}

func (_this TimingCoverage) Is_TimingCoverage() bool {
	_, ok := _this.Get_().(TimingCoverage_TimingCoverage)
	return ok
}

func (CompanionStruct_TimingCoverage_) Default() TimingCoverage {
	return Companion_TimingCoverage_.Create_TimingCoverage_(_dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this TimingCoverage) Dtor_startedMs() _dafny.Int {
	return _this.Get_().(TimingCoverage_TimingCoverage).StartedMs
}

func (_this TimingCoverage) Dtor_completedMs() _dafny.Int {
	return _this.Get_().(TimingCoverage_TimingCoverage).CompletedMs
}

func (_this TimingCoverage) Dtor_childUnionMs() _dafny.Int {
	return _this.Get_().(TimingCoverage_TimingCoverage).ChildUnionMs
}

func (_this TimingCoverage) Dtor_observedUnionMs() _dafny.Int {
	return _this.Get_().(TimingCoverage_TimingCoverage).ObservedUnionMs
}

func (_this TimingCoverage) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case TimingCoverage_TimingCoverage:
		{
			return "ConfluxSupervisedOperationResult.TimingCoverage.TimingCoverage" + "(" + _dafny.String(data.StartedMs) + ", " + _dafny.String(data.CompletedMs) + ", " + _dafny.String(data.ChildUnionMs) + ", " + _dafny.String(data.ObservedUnionMs) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this TimingCoverage) Equals(other TimingCoverage) bool {
	switch data1 := _this.Get_().(type) {
	case TimingCoverage_TimingCoverage:
		{
			data2, ok := other.Get_().(TimingCoverage_TimingCoverage)
			return ok && data1.StartedMs.Cmp(data2.StartedMs) == 0 && data1.CompletedMs.Cmp(data2.CompletedMs) == 0 && data1.ChildUnionMs.Cmp(data2.ChildUnionMs) == 0 && data1.ObservedUnionMs.Cmp(data2.ObservedUnionMs) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this TimingCoverage) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(TimingCoverage)
	return ok && _this.Equals(typed)
}

func Type_TimingCoverage_() _dafny.TypeDescriptor {
	return type_TimingCoverage_{}
}

type type_TimingCoverage_ struct {
}

func (_this type_TimingCoverage_) Default() interface{} {
	return Companion_TimingCoverage_.Default()
}

func (_this type_TimingCoverage_) String() string {
	return "ConfluxSupervisedOperationResult.TimingCoverage"
}
func (_this TimingCoverage) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = TimingCoverage{}

// End of datatype TimingCoverage

// Definition of datatype ExpandedResourceBudgetV1
type ExpandedResourceBudgetV1 struct {
	Data_ExpandedResourceBudgetV1_
}

func (_this ExpandedResourceBudgetV1) Get_() Data_ExpandedResourceBudgetV1_ {
	return _this.Data_ExpandedResourceBudgetV1_
}

type Data_ExpandedResourceBudgetV1_ interface {
	isExpandedResourceBudgetV1()
}

type CompanionStruct_ExpandedResourceBudgetV1_ struct {
}

var Companion_ExpandedResourceBudgetV1_ = CompanionStruct_ExpandedResourceBudgetV1_{}

type ExpandedResourceBudgetV1_ExpandedResourceBudgetV1 struct {
	TimeMs           _dafny.Int
	PeakRssMb        _dafny.Int
	CombinedOutputMb _dafny.Int
}

func (ExpandedResourceBudgetV1_ExpandedResourceBudgetV1) isExpandedResourceBudgetV1() {}

func (CompanionStruct_ExpandedResourceBudgetV1_) Create_ExpandedResourceBudgetV1_(TimeMs _dafny.Int, PeakRssMb _dafny.Int, CombinedOutputMb _dafny.Int) ExpandedResourceBudgetV1 {
	return ExpandedResourceBudgetV1{ExpandedResourceBudgetV1_ExpandedResourceBudgetV1{TimeMs, PeakRssMb, CombinedOutputMb}}
}

func (_this ExpandedResourceBudgetV1) Is_ExpandedResourceBudgetV1() bool {
	_, ok := _this.Get_().(ExpandedResourceBudgetV1_ExpandedResourceBudgetV1)
	return ok
}

func (CompanionStruct_ExpandedResourceBudgetV1_) Default() ExpandedResourceBudgetV1 {
	return Companion_ExpandedResourceBudgetV1_.Create_ExpandedResourceBudgetV1_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this ExpandedResourceBudgetV1) Dtor_timeMs() _dafny.Int {
	return _this.Get_().(ExpandedResourceBudgetV1_ExpandedResourceBudgetV1).TimeMs
}

func (_this ExpandedResourceBudgetV1) Dtor_peakRssMb() _dafny.Int {
	return _this.Get_().(ExpandedResourceBudgetV1_ExpandedResourceBudgetV1).PeakRssMb
}

func (_this ExpandedResourceBudgetV1) Dtor_combinedOutputMb() _dafny.Int {
	return _this.Get_().(ExpandedResourceBudgetV1_ExpandedResourceBudgetV1).CombinedOutputMb
}

func (_this ExpandedResourceBudgetV1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ExpandedResourceBudgetV1_ExpandedResourceBudgetV1:
		{
			return "ConfluxSupervisedOperationResult.ExpandedResourceBudgetV1.ExpandedResourceBudgetV1" + "(" + _dafny.String(data.TimeMs) + ", " + _dafny.String(data.PeakRssMb) + ", " + _dafny.String(data.CombinedOutputMb) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ExpandedResourceBudgetV1) Equals(other ExpandedResourceBudgetV1) bool {
	switch data1 := _this.Get_().(type) {
	case ExpandedResourceBudgetV1_ExpandedResourceBudgetV1:
		{
			data2, ok := other.Get_().(ExpandedResourceBudgetV1_ExpandedResourceBudgetV1)
			return ok && data1.TimeMs.Cmp(data2.TimeMs) == 0 && data1.PeakRssMb.Cmp(data2.PeakRssMb) == 0 && data1.CombinedOutputMb.Cmp(data2.CombinedOutputMb) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ExpandedResourceBudgetV1) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ExpandedResourceBudgetV1)
	return ok && _this.Equals(typed)
}

func Type_ExpandedResourceBudgetV1_() _dafny.TypeDescriptor {
	return type_ExpandedResourceBudgetV1_{}
}

type type_ExpandedResourceBudgetV1_ struct {
}

func (_this type_ExpandedResourceBudgetV1_) Default() interface{} {
	return Companion_ExpandedResourceBudgetV1_.Default()
}

func (_this type_ExpandedResourceBudgetV1_) String() string {
	return "ConfluxSupervisedOperationResult.ExpandedResourceBudgetV1"
}
func (_this ExpandedResourceBudgetV1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ExpandedResourceBudgetV1{}

// End of datatype ExpandedResourceBudgetV1

// Definition of datatype ProcessTreeStatus
type ProcessTreeStatus struct {
	Data_ProcessTreeStatus_
}

func (_this ProcessTreeStatus) Get_() Data_ProcessTreeStatus_ {
	return _this.Data_ProcessTreeStatus_
}

type Data_ProcessTreeStatus_ interface {
	isProcessTreeStatus()
}

type CompanionStruct_ProcessTreeStatus_ struct {
}

var Companion_ProcessTreeStatus_ = CompanionStruct_ProcessTreeStatus_{}

type ProcessTreeStatus_ProcessTreeObserved struct {
}

func (ProcessTreeStatus_ProcessTreeObserved) isProcessTreeStatus() {}

func (CompanionStruct_ProcessTreeStatus_) Create_ProcessTreeObserved_() ProcessTreeStatus {
	return ProcessTreeStatus{ProcessTreeStatus_ProcessTreeObserved{}}
}

func (_this ProcessTreeStatus) Is_ProcessTreeObserved() bool {
	_, ok := _this.Get_().(ProcessTreeStatus_ProcessTreeObserved)
	return ok
}

type ProcessTreeStatus_ProcessTreeUnavailable struct {
	Reason _dafny.Sequence
}

func (ProcessTreeStatus_ProcessTreeUnavailable) isProcessTreeStatus() {}

func (CompanionStruct_ProcessTreeStatus_) Create_ProcessTreeUnavailable_(Reason _dafny.Sequence) ProcessTreeStatus {
	return ProcessTreeStatus{ProcessTreeStatus_ProcessTreeUnavailable{Reason}}
}

func (_this ProcessTreeStatus) Is_ProcessTreeUnavailable() bool {
	_, ok := _this.Get_().(ProcessTreeStatus_ProcessTreeUnavailable)
	return ok
}

func (CompanionStruct_ProcessTreeStatus_) Default() ProcessTreeStatus {
	return Companion_ProcessTreeStatus_.Create_ProcessTreeObserved_()
}

func (_this ProcessTreeStatus) Dtor_reason() _dafny.Sequence {
	return _this.Get_().(ProcessTreeStatus_ProcessTreeUnavailable).Reason
}

func (_this ProcessTreeStatus) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ProcessTreeStatus_ProcessTreeObserved:
		{
			return "ConfluxSupervisedOperationResult.ProcessTreeStatus.ProcessTreeObserved"
		}
	case ProcessTreeStatus_ProcessTreeUnavailable:
		{
			return "ConfluxSupervisedOperationResult.ProcessTreeStatus.ProcessTreeUnavailable" + "(" + data.Reason.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ProcessTreeStatus) Equals(other ProcessTreeStatus) bool {
	switch data1 := _this.Get_().(type) {
	case ProcessTreeStatus_ProcessTreeObserved:
		{
			_, ok := other.Get_().(ProcessTreeStatus_ProcessTreeObserved)
			return ok
		}
	case ProcessTreeStatus_ProcessTreeUnavailable:
		{
			data2, ok := other.Get_().(ProcessTreeStatus_ProcessTreeUnavailable)
			return ok && data1.Reason.Equals(data2.Reason)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ProcessTreeStatus) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ProcessTreeStatus)
	return ok && _this.Equals(typed)
}

func Type_ProcessTreeStatus_() _dafny.TypeDescriptor {
	return type_ProcessTreeStatus_{}
}

type type_ProcessTreeStatus_ struct {
}

func (_this type_ProcessTreeStatus_) Default() interface{} {
	return Companion_ProcessTreeStatus_.Default()
}

func (_this type_ProcessTreeStatus_) String() string {
	return "ConfluxSupervisedOperationResult.ProcessTreeStatus"
}
func (_this ProcessTreeStatus) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ProcessTreeStatus{}

// End of datatype ProcessTreeStatus

// Definition of datatype CleanupStatus
type CleanupStatus struct {
	Data_CleanupStatus_
}

func (_this CleanupStatus) Get_() Data_CleanupStatus_ {
	return _this.Data_CleanupStatus_
}

type Data_CleanupStatus_ interface {
	isCleanupStatus()
}

type CompanionStruct_CleanupStatus_ struct {
}

var Companion_CleanupStatus_ = CompanionStruct_CleanupStatus_{}

type CleanupStatus_CleanupNotRequired struct {
}

func (CleanupStatus_CleanupNotRequired) isCleanupStatus() {}

func (CompanionStruct_CleanupStatus_) Create_CleanupNotRequired_() CleanupStatus {
	return CleanupStatus{CleanupStatus_CleanupNotRequired{}}
}

func (_this CleanupStatus) Is_CleanupNotRequired() bool {
	_, ok := _this.Get_().(CleanupStatus_CleanupNotRequired)
	return ok
}

type CleanupStatus_CleanupComplete struct {
	ElapsedMs  _dafny.Int
	DeadlineMs _dafny.Int
}

func (CleanupStatus_CleanupComplete) isCleanupStatus() {}

func (CompanionStruct_CleanupStatus_) Create_CleanupComplete_(ElapsedMs _dafny.Int, DeadlineMs _dafny.Int) CleanupStatus {
	return CleanupStatus{CleanupStatus_CleanupComplete{ElapsedMs, DeadlineMs}}
}

func (_this CleanupStatus) Is_CleanupComplete() bool {
	_, ok := _this.Get_().(CleanupStatus_CleanupComplete)
	return ok
}

type CleanupStatus_CleanupFailed struct {
	ElapsedMs  _dafny.Int
	DeadlineMs _dafny.Int
	Reason     _dafny.Sequence
}

func (CleanupStatus_CleanupFailed) isCleanupStatus() {}

func (CompanionStruct_CleanupStatus_) Create_CleanupFailed_(ElapsedMs _dafny.Int, DeadlineMs _dafny.Int, Reason _dafny.Sequence) CleanupStatus {
	return CleanupStatus{CleanupStatus_CleanupFailed{ElapsedMs, DeadlineMs, Reason}}
}

func (_this CleanupStatus) Is_CleanupFailed() bool {
	_, ok := _this.Get_().(CleanupStatus_CleanupFailed)
	return ok
}

func (CompanionStruct_CleanupStatus_) Default() CleanupStatus {
	return Companion_CleanupStatus_.Create_CleanupNotRequired_()
}

func (_this CleanupStatus) Dtor_elapsedMs() _dafny.Int {
	switch data := _this.Get_().(type) {
	case CleanupStatus_CleanupComplete:
		return data.ElapsedMs
	default:
		return data.(CleanupStatus_CleanupFailed).ElapsedMs
	}
}

func (_this CleanupStatus) Dtor_deadlineMs() _dafny.Int {
	switch data := _this.Get_().(type) {
	case CleanupStatus_CleanupComplete:
		return data.DeadlineMs
	default:
		return data.(CleanupStatus_CleanupFailed).DeadlineMs
	}
}

func (_this CleanupStatus) Dtor_reason() _dafny.Sequence {
	return _this.Get_().(CleanupStatus_CleanupFailed).Reason
}

func (_this CleanupStatus) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CleanupStatus_CleanupNotRequired:
		{
			return "ConfluxSupervisedOperationResult.CleanupStatus.CleanupNotRequired"
		}
	case CleanupStatus_CleanupComplete:
		{
			return "ConfluxSupervisedOperationResult.CleanupStatus.CleanupComplete" + "(" + _dafny.String(data.ElapsedMs) + ", " + _dafny.String(data.DeadlineMs) + ")"
		}
	case CleanupStatus_CleanupFailed:
		{
			return "ConfluxSupervisedOperationResult.CleanupStatus.CleanupFailed" + "(" + _dafny.String(data.ElapsedMs) + ", " + _dafny.String(data.DeadlineMs) + ", " + data.Reason.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CleanupStatus) Equals(other CleanupStatus) bool {
	switch data1 := _this.Get_().(type) {
	case CleanupStatus_CleanupNotRequired:
		{
			_, ok := other.Get_().(CleanupStatus_CleanupNotRequired)
			return ok
		}
	case CleanupStatus_CleanupComplete:
		{
			data2, ok := other.Get_().(CleanupStatus_CleanupComplete)
			return ok && data1.ElapsedMs.Cmp(data2.ElapsedMs) == 0 && data1.DeadlineMs.Cmp(data2.DeadlineMs) == 0
		}
	case CleanupStatus_CleanupFailed:
		{
			data2, ok := other.Get_().(CleanupStatus_CleanupFailed)
			return ok && data1.ElapsedMs.Cmp(data2.ElapsedMs) == 0 && data1.DeadlineMs.Cmp(data2.DeadlineMs) == 0 && data1.Reason.Equals(data2.Reason)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CleanupStatus) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CleanupStatus)
	return ok && _this.Equals(typed)
}

func Type_CleanupStatus_() _dafny.TypeDescriptor {
	return type_CleanupStatus_{}
}

type type_CleanupStatus_ struct {
}

func (_this type_CleanupStatus_) Default() interface{} {
	return Companion_CleanupStatus_.Default()
}

func (_this type_CleanupStatus_) String() string {
	return "ConfluxSupervisedOperationResult.CleanupStatus"
}
func (_this CleanupStatus) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CleanupStatus{}

// End of datatype CleanupStatus

// Definition of datatype ExitStatus
type ExitStatus struct {
	Data_ExitStatus_
}

func (_this ExitStatus) Get_() Data_ExitStatus_ {
	return _this.Data_ExitStatus_
}

type Data_ExitStatus_ interface {
	isExitStatus()
}

type CompanionStruct_ExitStatus_ struct {
}

var Companion_ExitStatus_ = CompanionStruct_ExitStatus_{}

type ExitStatus_ExitNotObserved struct {
}

func (ExitStatus_ExitNotObserved) isExitStatus() {}

func (CompanionStruct_ExitStatus_) Create_ExitNotObserved_() ExitStatus {
	return ExitStatus{ExitStatus_ExitNotObserved{}}
}

func (_this ExitStatus) Is_ExitNotObserved() bool {
	_, ok := _this.Get_().(ExitStatus_ExitNotObserved)
	return ok
}

type ExitStatus_ExitObserved struct {
	Code _dafny.Int
}

func (ExitStatus_ExitObserved) isExitStatus() {}

func (CompanionStruct_ExitStatus_) Create_ExitObserved_(Code _dafny.Int) ExitStatus {
	return ExitStatus{ExitStatus_ExitObserved{Code}}
}

func (_this ExitStatus) Is_ExitObserved() bool {
	_, ok := _this.Get_().(ExitStatus_ExitObserved)
	return ok
}

func (CompanionStruct_ExitStatus_) Default() ExitStatus {
	return Companion_ExitStatus_.Create_ExitNotObserved_()
}

func (_this ExitStatus) Dtor_code() _dafny.Int {
	return _this.Get_().(ExitStatus_ExitObserved).Code
}

func (_this ExitStatus) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ExitStatus_ExitNotObserved:
		{
			return "ConfluxSupervisedOperationResult.ExitStatus.ExitNotObserved"
		}
	case ExitStatus_ExitObserved:
		{
			return "ConfluxSupervisedOperationResult.ExitStatus.ExitObserved" + "(" + _dafny.String(data.Code) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ExitStatus) Equals(other ExitStatus) bool {
	switch data1 := _this.Get_().(type) {
	case ExitStatus_ExitNotObserved:
		{
			_, ok := other.Get_().(ExitStatus_ExitNotObserved)
			return ok
		}
	case ExitStatus_ExitObserved:
		{
			data2, ok := other.Get_().(ExitStatus_ExitObserved)
			return ok && data1.Code.Cmp(data2.Code) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ExitStatus) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ExitStatus)
	return ok && _this.Equals(typed)
}

func Type_ExitStatus_() _dafny.TypeDescriptor {
	return type_ExitStatus_{}
}

type type_ExitStatus_ struct {
}

func (_this type_ExitStatus_) Default() interface{} {
	return Companion_ExitStatus_.Default()
}

func (_this type_ExitStatus_) String() string {
	return "ConfluxSupervisedOperationResult.ExitStatus"
}
func (_this ExitStatus) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ExitStatus{}

// End of datatype ExitStatus

// Definition of datatype EnforcementStatus
type EnforcementStatus struct {
	Data_EnforcementStatus_
}

func (_this EnforcementStatus) Get_() Data_EnforcementStatus_ {
	return _this.Data_EnforcementStatus_
}

type Data_EnforcementStatus_ interface {
	isEnforcementStatus()
}

type CompanionStruct_EnforcementStatus_ struct {
}

var Companion_EnforcementStatus_ = CompanionStruct_EnforcementStatus_{}

type EnforcementStatus_NoEnforcement struct {
}

func (EnforcementStatus_NoEnforcement) isEnforcementStatus() {}

func (CompanionStruct_EnforcementStatus_) Create_NoEnforcement_() EnforcementStatus {
	return EnforcementStatus{EnforcementStatus_NoEnforcement{}}
}

func (_this EnforcementStatus) Is_NoEnforcement() bool {
	_, ok := _this.Get_().(EnforcementStatus_NoEnforcement)
	return ok
}

type EnforcementStatus_PostHocFailure struct {
}

func (EnforcementStatus_PostHocFailure) isEnforcementStatus() {}

func (CompanionStruct_EnforcementStatus_) Create_PostHocFailure_() EnforcementStatus {
	return EnforcementStatus{EnforcementStatus_PostHocFailure{}}
}

func (_this EnforcementStatus) Is_PostHocFailure() bool {
	_, ok := _this.Get_().(EnforcementStatus_PostHocFailure)
	return ok
}

type EnforcementStatus_InFlightProcessTreeTermination struct {
}

func (EnforcementStatus_InFlightProcessTreeTermination) isEnforcementStatus() {}

func (CompanionStruct_EnforcementStatus_) Create_InFlightProcessTreeTermination_() EnforcementStatus {
	return EnforcementStatus{EnforcementStatus_InFlightProcessTreeTermination{}}
}

func (_this EnforcementStatus) Is_InFlightProcessTreeTermination() bool {
	_, ok := _this.Get_().(EnforcementStatus_InFlightProcessTreeTermination)
	return ok
}

func (CompanionStruct_EnforcementStatus_) Default() EnforcementStatus {
	return Companion_EnforcementStatus_.Create_NoEnforcement_()
}

func (_ CompanionStruct_EnforcementStatus_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_EnforcementStatus_.Create_NoEnforcement_(), true
		case 1:
			return Companion_EnforcementStatus_.Create_PostHocFailure_(), true
		case 2:
			return Companion_EnforcementStatus_.Create_InFlightProcessTreeTermination_(), true
		default:
			return EnforcementStatus{}, false
		}
	}
}

func (_this EnforcementStatus) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case EnforcementStatus_NoEnforcement:
		{
			return "ConfluxSupervisedOperationResult.EnforcementStatus.NoEnforcement"
		}
	case EnforcementStatus_PostHocFailure:
		{
			return "ConfluxSupervisedOperationResult.EnforcementStatus.PostHocFailure"
		}
	case EnforcementStatus_InFlightProcessTreeTermination:
		{
			return "ConfluxSupervisedOperationResult.EnforcementStatus.InFlightProcessTreeTermination"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this EnforcementStatus) Equals(other EnforcementStatus) bool {
	switch _this.Get_().(type) {
	case EnforcementStatus_NoEnforcement:
		{
			_, ok := other.Get_().(EnforcementStatus_NoEnforcement)
			return ok
		}
	case EnforcementStatus_PostHocFailure:
		{
			_, ok := other.Get_().(EnforcementStatus_PostHocFailure)
			return ok
		}
	case EnforcementStatus_InFlightProcessTreeTermination:
		{
			_, ok := other.Get_().(EnforcementStatus_InFlightProcessTreeTermination)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this EnforcementStatus) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(EnforcementStatus)
	return ok && _this.Equals(typed)
}

func Type_EnforcementStatus_() _dafny.TypeDescriptor {
	return type_EnforcementStatus_{}
}

type type_EnforcementStatus_ struct {
}

func (_this type_EnforcementStatus_) Default() interface{} {
	return Companion_EnforcementStatus_.Default()
}

func (_this type_EnforcementStatus_) String() string {
	return "ConfluxSupervisedOperationResult.EnforcementStatus"
}
func (_this EnforcementStatus) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = EnforcementStatus{}

// End of datatype EnforcementStatus

// Definition of datatype TerminationScope
type TerminationScope struct {
	Data_TerminationScope_
}

func (_this TerminationScope) Get_() Data_TerminationScope_ {
	return _this.Data_TerminationScope_
}

type Data_TerminationScope_ interface {
	isTerminationScope()
}

type CompanionStruct_TerminationScope_ struct {
}

var Companion_TerminationScope_ = CompanionStruct_TerminationScope_{}

type TerminationScope_NoTermination struct {
}

func (TerminationScope_NoTermination) isTerminationScope() {}

func (CompanionStruct_TerminationScope_) Create_NoTermination_() TerminationScope {
	return TerminationScope{TerminationScope_NoTermination{}}
}

func (_this TerminationScope) Is_NoTermination() bool {
	_, ok := _this.Get_().(TerminationScope_NoTermination)
	return ok
}

type TerminationScope_ProcessTreeTermination struct {
}

func (TerminationScope_ProcessTreeTermination) isTerminationScope() {}

func (CompanionStruct_TerminationScope_) Create_ProcessTreeTermination_() TerminationScope {
	return TerminationScope{TerminationScope_ProcessTreeTermination{}}
}

func (_this TerminationScope) Is_ProcessTreeTermination() bool {
	_, ok := _this.Get_().(TerminationScope_ProcessTreeTermination)
	return ok
}

func (CompanionStruct_TerminationScope_) Default() TerminationScope {
	return Companion_TerminationScope_.Create_NoTermination_()
}

func (_ CompanionStruct_TerminationScope_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_TerminationScope_.Create_NoTermination_(), true
		case 1:
			return Companion_TerminationScope_.Create_ProcessTreeTermination_(), true
		default:
			return TerminationScope{}, false
		}
	}
}

func (_this TerminationScope) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case TerminationScope_NoTermination:
		{
			return "ConfluxSupervisedOperationResult.TerminationScope.NoTermination"
		}
	case TerminationScope_ProcessTreeTermination:
		{
			return "ConfluxSupervisedOperationResult.TerminationScope.ProcessTreeTermination"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this TerminationScope) Equals(other TerminationScope) bool {
	switch _this.Get_().(type) {
	case TerminationScope_NoTermination:
		{
			_, ok := other.Get_().(TerminationScope_NoTermination)
			return ok
		}
	case TerminationScope_ProcessTreeTermination:
		{
			_, ok := other.Get_().(TerminationScope_ProcessTreeTermination)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this TerminationScope) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(TerminationScope)
	return ok && _this.Equals(typed)
}

func Type_TerminationScope_() _dafny.TypeDescriptor {
	return type_TerminationScope_{}
}

type type_TerminationScope_ struct {
}

func (_this type_TerminationScope_) Default() interface{} {
	return Companion_TerminationScope_.Default()
}

func (_this type_TerminationScope_) String() string {
	return "ConfluxSupervisedOperationResult.TerminationScope"
}
func (_this TerminationScope) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = TerminationScope{}

// End of datatype TerminationScope

// Definition of datatype SupervisionObservationV1
type SupervisionObservationV1 struct {
	Data_SupervisionObservationV1_
}

func (_this SupervisionObservationV1) Get_() Data_SupervisionObservationV1_ {
	return _this.Data_SupervisionObservationV1_
}

type Data_SupervisionObservationV1_ interface {
	isSupervisionObservationV1()
}

type CompanionStruct_SupervisionObservationV1_ struct {
}

var Companion_SupervisionObservationV1_ = CompanionStruct_SupervisionObservationV1_{}

type SupervisionObservationV1_SupervisionObservationV1 struct {
	Timing              TimingCoverage
	PeakDescendantRssMb m_ConfluxResourceCeilings.ResourceMeasurement
	CombinedOutputBytes m_ConfluxResourceCeilings.ResourceMeasurement
	ProcessTree         ProcessTreeStatus
}

func (SupervisionObservationV1_SupervisionObservationV1) isSupervisionObservationV1() {}

func (CompanionStruct_SupervisionObservationV1_) Create_SupervisionObservationV1_(Timing TimingCoverage, PeakDescendantRssMb m_ConfluxResourceCeilings.ResourceMeasurement, CombinedOutputBytes m_ConfluxResourceCeilings.ResourceMeasurement, ProcessTree ProcessTreeStatus) SupervisionObservationV1 {
	return SupervisionObservationV1{SupervisionObservationV1_SupervisionObservationV1{Timing, PeakDescendantRssMb, CombinedOutputBytes, ProcessTree}}
}

func (_this SupervisionObservationV1) Is_SupervisionObservationV1() bool {
	_, ok := _this.Get_().(SupervisionObservationV1_SupervisionObservationV1)
	return ok
}

func (CompanionStruct_SupervisionObservationV1_) Default() SupervisionObservationV1 {
	return Companion_SupervisionObservationV1_.Create_SupervisionObservationV1_(Companion_TimingCoverage_.Default(), m_ConfluxResourceCeilings.Companion_ResourceMeasurement_.Default(), m_ConfluxResourceCeilings.Companion_ResourceMeasurement_.Default(), Companion_ProcessTreeStatus_.Default())
}

func (_this SupervisionObservationV1) Dtor_timing() TimingCoverage {
	return _this.Get_().(SupervisionObservationV1_SupervisionObservationV1).Timing
}

func (_this SupervisionObservationV1) Dtor_peakDescendantRssMb() m_ConfluxResourceCeilings.ResourceMeasurement {
	return _this.Get_().(SupervisionObservationV1_SupervisionObservationV1).PeakDescendantRssMb
}

func (_this SupervisionObservationV1) Dtor_combinedOutputBytes() m_ConfluxResourceCeilings.ResourceMeasurement {
	return _this.Get_().(SupervisionObservationV1_SupervisionObservationV1).CombinedOutputBytes
}

func (_this SupervisionObservationV1) Dtor_processTree() ProcessTreeStatus {
	return _this.Get_().(SupervisionObservationV1_SupervisionObservationV1).ProcessTree
}

func (_this SupervisionObservationV1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SupervisionObservationV1_SupervisionObservationV1:
		{
			return "ConfluxSupervisedOperationResult.SupervisionObservationV1.SupervisionObservationV1" + "(" + _dafny.String(data.Timing) + ", " + _dafny.String(data.PeakDescendantRssMb) + ", " + _dafny.String(data.CombinedOutputBytes) + ", " + _dafny.String(data.ProcessTree) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SupervisionObservationV1) Equals(other SupervisionObservationV1) bool {
	switch data1 := _this.Get_().(type) {
	case SupervisionObservationV1_SupervisionObservationV1:
		{
			data2, ok := other.Get_().(SupervisionObservationV1_SupervisionObservationV1)
			return ok && data1.Timing.Equals(data2.Timing) && data1.PeakDescendantRssMb.Equals(data2.PeakDescendantRssMb) && data1.CombinedOutputBytes.Equals(data2.CombinedOutputBytes) && data1.ProcessTree.Equals(data2.ProcessTree)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SupervisionObservationV1) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SupervisionObservationV1)
	return ok && _this.Equals(typed)
}

func Type_SupervisionObservationV1_() _dafny.TypeDescriptor {
	return type_SupervisionObservationV1_{}
}

type type_SupervisionObservationV1_ struct {
}

func (_this type_SupervisionObservationV1_) Default() interface{} {
	return Companion_SupervisionObservationV1_.Default()
}

func (_this type_SupervisionObservationV1_) String() string {
	return "ConfluxSupervisedOperationResult.SupervisionObservationV1"
}
func (_this SupervisionObservationV1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SupervisionObservationV1{}

// End of datatype SupervisionObservationV1

// Definition of datatype FirstBreachV1
type FirstBreachV1 struct {
	Data_FirstBreachV1_
}

func (_this FirstBreachV1) Get_() Data_FirstBreachV1_ {
	return _this.Data_FirstBreachV1_
}

type Data_FirstBreachV1_ interface {
	isFirstBreachV1()
}

type CompanionStruct_FirstBreachV1_ struct {
}

var Companion_FirstBreachV1_ = CompanionStruct_FirstBreachV1_{}

type FirstBreachV1_FirstBreachV1 struct {
	SampleIndex        _dafny.Int
	Verdict            m_ConfluxResourceCeilings.ResourceVerdict
	MeasurementFailure bool
}

func (FirstBreachV1_FirstBreachV1) isFirstBreachV1() {}

func (CompanionStruct_FirstBreachV1_) Create_FirstBreachV1_(SampleIndex _dafny.Int, Verdict m_ConfluxResourceCeilings.ResourceVerdict, MeasurementFailure bool) FirstBreachV1 {
	return FirstBreachV1{FirstBreachV1_FirstBreachV1{SampleIndex, Verdict, MeasurementFailure}}
}

func (_this FirstBreachV1) Is_FirstBreachV1() bool {
	_, ok := _this.Get_().(FirstBreachV1_FirstBreachV1)
	return ok
}

func (CompanionStruct_FirstBreachV1_) Default() FirstBreachV1 {
	return Companion_FirstBreachV1_.Create_FirstBreachV1_(_dafny.Zero, m_ConfluxResourceCeilings.Companion_ResourceVerdict_.Default(), false)
}

func (_this FirstBreachV1) Dtor_sampleIndex() _dafny.Int {
	return _this.Get_().(FirstBreachV1_FirstBreachV1).SampleIndex
}

func (_this FirstBreachV1) Dtor_verdict() m_ConfluxResourceCeilings.ResourceVerdict {
	return _this.Get_().(FirstBreachV1_FirstBreachV1).Verdict
}

func (_this FirstBreachV1) Dtor_measurementFailure() bool {
	return _this.Get_().(FirstBreachV1_FirstBreachV1).MeasurementFailure
}

func (_this FirstBreachV1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case FirstBreachV1_FirstBreachV1:
		{
			return "ConfluxSupervisedOperationResult.FirstBreachV1.FirstBreachV1" + "(" + _dafny.String(data.SampleIndex) + ", " + _dafny.String(data.Verdict) + ", " + _dafny.String(data.MeasurementFailure) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this FirstBreachV1) Equals(other FirstBreachV1) bool {
	switch data1 := _this.Get_().(type) {
	case FirstBreachV1_FirstBreachV1:
		{
			data2, ok := other.Get_().(FirstBreachV1_FirstBreachV1)
			return ok && data1.SampleIndex.Cmp(data2.SampleIndex) == 0 && data1.Verdict.Equals(data2.Verdict) && data1.MeasurementFailure == data2.MeasurementFailure
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this FirstBreachV1) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(FirstBreachV1)
	return ok && _this.Equals(typed)
}

func Type_FirstBreachV1_() _dafny.TypeDescriptor {
	return type_FirstBreachV1_{}
}

type type_FirstBreachV1_ struct {
}

func (_this type_FirstBreachV1_) Default() interface{} {
	return Companion_FirstBreachV1_.Default()
}

func (_this type_FirstBreachV1_) String() string {
	return "ConfluxSupervisedOperationResult.FirstBreachV1"
}
func (_this FirstBreachV1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = FirstBreachV1{}

// End of datatype FirstBreachV1

// Definition of datatype SupervisionState
type SupervisionState struct {
	Data_SupervisionState_
}

func (_this SupervisionState) Get_() Data_SupervisionState_ {
	return _this.Data_SupervisionState_
}

type Data_SupervisionState_ interface {
	isSupervisionState()
}

type CompanionStruct_SupervisionState_ struct {
}

var Companion_SupervisionState_ = CompanionStruct_SupervisionState_{}

type SupervisionState_SupervisionRunning struct {
	SamplesSeen _dafny.Int
}

func (SupervisionState_SupervisionRunning) isSupervisionState() {}

func (CompanionStruct_SupervisionState_) Create_SupervisionRunning_(SamplesSeen _dafny.Int) SupervisionState {
	return SupervisionState{SupervisionState_SupervisionRunning{SamplesSeen}}
}

func (_this SupervisionState) Is_SupervisionRunning() bool {
	_, ok := _this.Get_().(SupervisionState_SupervisionRunning)
	return ok
}

type SupervisionState_SupervisionTripped struct {
	Record FirstBreachV1
}

func (SupervisionState_SupervisionTripped) isSupervisionState() {}

func (CompanionStruct_SupervisionState_) Create_SupervisionTripped_(Record FirstBreachV1) SupervisionState {
	return SupervisionState{SupervisionState_SupervisionTripped{Record}}
}

func (_this SupervisionState) Is_SupervisionTripped() bool {
	_, ok := _this.Get_().(SupervisionState_SupervisionTripped)
	return ok
}

func (CompanionStruct_SupervisionState_) Default() SupervisionState {
	return Companion_SupervisionState_.Create_SupervisionRunning_(_dafny.Zero)
}

func (_this SupervisionState) Dtor_samplesSeen() _dafny.Int {
	return _this.Get_().(SupervisionState_SupervisionRunning).SamplesSeen
}

func (_this SupervisionState) Dtor_record() FirstBreachV1 {
	return _this.Get_().(SupervisionState_SupervisionTripped).Record
}

func (_this SupervisionState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SupervisionState_SupervisionRunning:
		{
			return "ConfluxSupervisedOperationResult.SupervisionState.SupervisionRunning" + "(" + _dafny.String(data.SamplesSeen) + ")"
		}
	case SupervisionState_SupervisionTripped:
		{
			return "ConfluxSupervisedOperationResult.SupervisionState.SupervisionTripped" + "(" + _dafny.String(data.Record) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SupervisionState) Equals(other SupervisionState) bool {
	switch data1 := _this.Get_().(type) {
	case SupervisionState_SupervisionRunning:
		{
			data2, ok := other.Get_().(SupervisionState_SupervisionRunning)
			return ok && data1.SamplesSeen.Cmp(data2.SamplesSeen) == 0
		}
	case SupervisionState_SupervisionTripped:
		{
			data2, ok := other.Get_().(SupervisionState_SupervisionTripped)
			return ok && data1.Record.Equals(data2.Record)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SupervisionState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SupervisionState)
	return ok && _this.Equals(typed)
}

func Type_SupervisionState_() _dafny.TypeDescriptor {
	return type_SupervisionState_{}
}

type type_SupervisionState_ struct {
}

func (_this type_SupervisionState_) Default() interface{} {
	return Companion_SupervisionState_.Default()
}

func (_this type_SupervisionState_) String() string {
	return "ConfluxSupervisedOperationResult.SupervisionState"
}
func (_this SupervisionState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SupervisionState{}

// End of datatype SupervisionState

// Definition of datatype SupervisedOperationResultV1
type SupervisedOperationResultV1 struct {
	Data_SupervisedOperationResultV1_
}

func (_this SupervisedOperationResultV1) Get_() Data_SupervisedOperationResultV1_ {
	return _this.Data_SupervisedOperationResultV1_
}

type Data_SupervisedOperationResultV1_ interface {
	isSupervisedOperationResultV1()
}

type CompanionStruct_SupervisedOperationResultV1_ struct {
}

var Companion_SupervisedOperationResultV1_ = CompanionStruct_SupervisedOperationResultV1_{}

type SupervisedOperationResultV1_SupervisedOperationResultV1 struct {
	SchemaVersion     _dafny.Int
	PolicyId          _dafny.Sequence
	RequestedPhase    m_ConfluxResourceCeilings.ExecutionPhase
	EffectivePhase    m_ConfluxResourceCeilings.ExecutionPhase
	ExpandedCeilings  ExpandedResourceBudgetV1
	ObservationStatus m_ConfluxResourceCeilings.MeasurementStatus
	Observation       SupervisionObservationV1
	SupervisionState  SupervisionState
	Cleanup           CleanupStatus
	Exit              ExitStatus
	Enforcement       EnforcementStatus
	Termination       TerminationScope
	WrapperExitCode   _dafny.Int
	Reason            _dafny.Sequence
	Verdict           m_ConfluxResourceCeilings.ResourceVerdict
}

func (SupervisedOperationResultV1_SupervisedOperationResultV1) isSupervisedOperationResultV1() {}

func (CompanionStruct_SupervisedOperationResultV1_) Create_SupervisedOperationResultV1_(SchemaVersion _dafny.Int, PolicyId _dafny.Sequence, RequestedPhase m_ConfluxResourceCeilings.ExecutionPhase, EffectivePhase m_ConfluxResourceCeilings.ExecutionPhase, ExpandedCeilings ExpandedResourceBudgetV1, ObservationStatus m_ConfluxResourceCeilings.MeasurementStatus, Observation SupervisionObservationV1, SupervisionState SupervisionState, Cleanup CleanupStatus, Exit ExitStatus, Enforcement EnforcementStatus, Termination TerminationScope, WrapperExitCode _dafny.Int, Reason _dafny.Sequence, Verdict m_ConfluxResourceCeilings.ResourceVerdict) SupervisedOperationResultV1 {
	return SupervisedOperationResultV1{SupervisedOperationResultV1_SupervisedOperationResultV1{SchemaVersion, PolicyId, RequestedPhase, EffectivePhase, ExpandedCeilings, ObservationStatus, Observation, SupervisionState, Cleanup, Exit, Enforcement, Termination, WrapperExitCode, Reason, Verdict}}
}

func (_this SupervisedOperationResultV1) Is_SupervisedOperationResultV1() bool {
	_, ok := _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1)
	return ok
}

func (CompanionStruct_SupervisedOperationResultV1_) Default() SupervisedOperationResultV1 {
	return Companion_SupervisedOperationResultV1_.Create_SupervisedOperationResultV1_(_dafny.Zero, _dafny.EmptySeq, m_ConfluxResourceCeilings.Companion_ExecutionPhase_.Default(), m_ConfluxResourceCeilings.Companion_ExecutionPhase_.Default(), Companion_ExpandedResourceBudgetV1_.Default(), m_ConfluxResourceCeilings.Companion_MeasurementStatus_.Default(), Companion_SupervisionObservationV1_.Default(), Companion_SupervisionState_.Default(), Companion_CleanupStatus_.Default(), Companion_ExitStatus_.Default(), Companion_EnforcementStatus_.Default(), Companion_TerminationScope_.Default(), _dafny.Zero, _dafny.EmptySeq, m_ConfluxResourceCeilings.Companion_ResourceVerdict_.Default())
}

func (_this SupervisedOperationResultV1) Dtor_schemaVersion() _dafny.Int {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).SchemaVersion
}

func (_this SupervisedOperationResultV1) Dtor_policyId() _dafny.Sequence {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).PolicyId
}

func (_this SupervisedOperationResultV1) Dtor_requestedPhase() m_ConfluxResourceCeilings.ExecutionPhase {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).RequestedPhase
}

func (_this SupervisedOperationResultV1) Dtor_effectivePhase() m_ConfluxResourceCeilings.ExecutionPhase {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).EffectivePhase
}

func (_this SupervisedOperationResultV1) Dtor_expandedCeilings() ExpandedResourceBudgetV1 {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).ExpandedCeilings
}

func (_this SupervisedOperationResultV1) Dtor_observationStatus() m_ConfluxResourceCeilings.MeasurementStatus {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).ObservationStatus
}

func (_this SupervisedOperationResultV1) Dtor_observation() SupervisionObservationV1 {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).Observation
}

func (_this SupervisedOperationResultV1) Dtor_supervisionState() SupervisionState {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).SupervisionState
}

func (_this SupervisedOperationResultV1) Dtor_cleanup() CleanupStatus {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).Cleanup
}

func (_this SupervisedOperationResultV1) Dtor_exit() ExitStatus {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).Exit
}

func (_this SupervisedOperationResultV1) Dtor_enforcement() EnforcementStatus {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).Enforcement
}

func (_this SupervisedOperationResultV1) Dtor_termination() TerminationScope {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).Termination
}

func (_this SupervisedOperationResultV1) Dtor_wrapperExitCode() _dafny.Int {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).WrapperExitCode
}

func (_this SupervisedOperationResultV1) Dtor_reason() _dafny.Sequence {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).Reason
}

func (_this SupervisedOperationResultV1) Dtor_verdict() m_ConfluxResourceCeilings.ResourceVerdict {
	return _this.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1).Verdict
}

func (_this SupervisedOperationResultV1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SupervisedOperationResultV1_SupervisedOperationResultV1:
		{
			return "ConfluxSupervisedOperationResult.SupervisedOperationResultV1.SupervisedOperationResultV1" + "(" + _dafny.String(data.SchemaVersion) + ", " + data.PolicyId.VerbatimString(true) + ", " + _dafny.String(data.RequestedPhase) + ", " + _dafny.String(data.EffectivePhase) + ", " + _dafny.String(data.ExpandedCeilings) + ", " + _dafny.String(data.ObservationStatus) + ", " + _dafny.String(data.Observation) + ", " + _dafny.String(data.SupervisionState) + ", " + _dafny.String(data.Cleanup) + ", " + _dafny.String(data.Exit) + ", " + _dafny.String(data.Enforcement) + ", " + _dafny.String(data.Termination) + ", " + _dafny.String(data.WrapperExitCode) + ", " + data.Reason.VerbatimString(true) + ", " + _dafny.String(data.Verdict) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SupervisedOperationResultV1) Equals(other SupervisedOperationResultV1) bool {
	switch data1 := _this.Get_().(type) {
	case SupervisedOperationResultV1_SupervisedOperationResultV1:
		{
			data2, ok := other.Get_().(SupervisedOperationResultV1_SupervisedOperationResultV1)
			return ok && data1.SchemaVersion.Cmp(data2.SchemaVersion) == 0 && data1.PolicyId.Equals(data2.PolicyId) && data1.RequestedPhase.Equals(data2.RequestedPhase) && data1.EffectivePhase.Equals(data2.EffectivePhase) && data1.ExpandedCeilings.Equals(data2.ExpandedCeilings) && data1.ObservationStatus.Equals(data2.ObservationStatus) && data1.Observation.Equals(data2.Observation) && data1.SupervisionState.Equals(data2.SupervisionState) && data1.Cleanup.Equals(data2.Cleanup) && data1.Exit.Equals(data2.Exit) && data1.Enforcement.Equals(data2.Enforcement) && data1.Termination.Equals(data2.Termination) && data1.WrapperExitCode.Cmp(data2.WrapperExitCode) == 0 && data1.Reason.Equals(data2.Reason) && data1.Verdict.Equals(data2.Verdict)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SupervisedOperationResultV1) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SupervisedOperationResultV1)
	return ok && _this.Equals(typed)
}

func Type_SupervisedOperationResultV1_() _dafny.TypeDescriptor {
	return type_SupervisedOperationResultV1_{}
}

type type_SupervisedOperationResultV1_ struct {
}

func (_this type_SupervisedOperationResultV1_) Default() interface{} {
	return Companion_SupervisedOperationResultV1_.Default()
}

func (_this type_SupervisedOperationResultV1_) String() string {
	return "ConfluxSupervisedOperationResult.SupervisedOperationResultV1"
}
func (_this SupervisedOperationResultV1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SupervisedOperationResultV1{}

// End of datatype SupervisedOperationResultV1
