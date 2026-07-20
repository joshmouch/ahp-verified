// Package ConfluxServiceLifecycle
// Dafny module ConfluxServiceLifecycle compiled into Go

package ConfluxServiceLifecycle

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
	m_ConfluxResourceSupervisor "github.com/joshmouch/ahp-verified/go/ConfluxResourceSupervisor"
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
var _ m_ConfluxResourceSupervisor.Dummy__

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
	return "ConfluxServiceLifecycle.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ValidPort(port _dafny.Int) bool {
	return ((port).Sign() == 1) && ((port).Cmp(_dafny.IntOfInt64(65535)) <= 0)
}
func (_static *CompanionStruct_Default___) ValidAbsoluteHttpPath(path _dafny.Sequence) bool {
	return ((_dafny.IntOfUint32((path).Cardinality())).Sign() == 1) && (((path).Select(0).(_dafny.CodePoint)) == (_dafny.CodePoint('/')))
}
func (_static *CompanionStruct_Default___) AllNonEmpty(values _dafny.Sequence) bool {
	return ((_dafny.IntOfUint32((values).Cardinality())).Sign() == 0) || (((_dafny.IntOfUint32(((values).Select(0).(_dafny.Sequence)).Cardinality())).Sign() == 1) && (Companion_Default___.AllNonEmpty((values).Drop(1))))
}
func (_static *CompanionStruct_Default___) ValidServiceSpec(spec ServiceSpec) bool {
	return (((((((((((((((spec).Dtor_schemaVersion()).Cmp(_dafny.One) == 0) && ((_dafny.IntOfUint32(((spec).Dtor_serviceId()).Cardinality())).Sign() == 1)) && ((_dafny.IntOfUint32(((spec).Dtor_sourceClosure()).Cardinality())).Sign() == 1)) && (Companion_Default___.AllNonEmpty((spec).Dtor_sourceClosure()))) && ((_dafny.IntOfUint32(((spec).Dtor_buildGeneration()).Cardinality())).Sign() == 1)) && ((_dafny.IntOfUint32(((spec).Dtor_runGeneration()).Cardinality())).Sign() == 1)) && ((_dafny.IntOfUint32((((spec).Dtor_healthIdentity()).Dtor_owner()).Cardinality())).Sign() == 1)) && ((_dafny.IntOfUint32((((spec).Dtor_healthIdentity()).Dtor_protocolVersion()).Cardinality())).Sign() == 1)) && (Companion_Default___.ValidPort(((spec).Dtor_endpointPolicy()).Dtor_port()))) && ((_dafny.Companion_Sequence_.Contains(((spec).Dtor_endpointPolicy()).Dtor_allowedBindKinds(), Companion_BindKind_.Create_TailscaleIp_())) || (_dafny.Companion_Sequence_.Contains(((spec).Dtor_endpointPolicy()).Dtor_allowedBindKinds(), Companion_BindKind_.Create_LoopbackBind_())))) && (Companion_Default___.ValidAbsoluteHttpPath((spec).Dtor_consolePath()))) && (Companion_Default___.ValidAbsoluteHttpPath((spec).Dtor_channelPath()))) && ((_dafny.IntOfUint32(((spec).Dtor_stateRoot()).Cardinality())).Sign() == 1)) && ((_dafny.IntOfUint32(((spec).Dtor_supervisorAdapter()).Cardinality())).Sign() == 1)
}
func (_static *CompanionStruct_Default___) ValidServiceInfo(info ServiceInfo) bool {
	return (((((((((((((info).Dtor_schemaVersion()).Cmp(_dafny.One) == 0) && ((_dafny.IntOfUint32(((info).Dtor_serviceId()).Cardinality())).Sign() == 1)) && ((_dafny.IntOfUint32((((info).Dtor_healthIdentity()).Dtor_owner()).Cardinality())).Sign() == 1)) && ((_dafny.IntOfUint32((((info).Dtor_healthIdentity()).Dtor_protocolVersion()).Cardinality())).Sign() == 1)) && ((_dafny.IntOfUint32(((info).Dtor_ownerLeaseId()).Cardinality())).Sign() == 1)) && (((info).Dtor_pid()).Sign() == 1)) && ((_dafny.IntOfUint32(((info).Dtor_generation()).Cardinality())).Sign() == 1)) && ((_dafny.IntOfUint32(((info).Dtor_bindHost()).Cardinality())).Sign() == 1)) && ((_dafny.IntOfUint32(((info).Dtor_publicHost()).Cardinality())).Sign() == 1)) && (Companion_Default___.ValidPort((info).Dtor_port()))) && ((_dafny.IntOfUint32(((info).Dtor_consoleUrl()).Cardinality())).Sign() == 1)) && ((_dafny.IntOfUint32(((info).Dtor_channelBase()).Cardinality())).Sign() == 1)
}
func (_static *CompanionStruct_Default___) ExactHealth(info ServiceInfo, observed HealthObservation) bool {
	return (((((_dafny.Companion_Sequence_.Equal((observed).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("ok"))) && (_dafny.Companion_Sequence_.Equal((observed).Dtor_owner(), ((info).Dtor_healthIdentity()).Dtor_owner()))) && (_dafny.Companion_Sequence_.Equal((observed).Dtor_protocolVersion(), ((info).Dtor_healthIdentity()).Dtor_protocolVersion()))) && (((observed).Dtor_pid()).Cmp((info).Dtor_pid()) == 0)) && (_dafny.Companion_Sequence_.Equal((observed).Dtor_generation(), (info).Dtor_generation()))) && (_dafny.Companion_Sequence_.Equal((observed).Dtor_ownerLeaseId(), (info).Dtor_ownerLeaseId()))
}
func (_static *CompanionStruct_Default___) CanUseTailscale(policy EndpointPolicy, facts TailscaleFacts) bool {
	return (((((((policy).Dtor_preferTailscaleMagicDns()) && (_dafny.Companion_Sequence_.Contains((policy).Dtor_allowedBindKinds(), Companion_BindKind_.Create_TailscaleIp_()))) && ((facts).Dtor_backendRunning())) && ((facts).Dtor_magicDnsEnabled())) && ((facts).Dtor_selfOnline())) && ((facts).Dtor_dnsNameValid())) && ((facts).Dtor_matchingIpValid())
}
func (_static *CompanionStruct_Default___) ChooseEndpoint(policy EndpointPolicy, facts TailscaleFacts) EndpointDecision {
	if Companion_Default___.CanUseTailscale(policy, facts) {
		return Companion_EndpointDecision_.Create_UseEndpoint_(Companion_Endpoint_.Create_Endpoint_((facts).Dtor_matchingIp(), (facts).Dtor_dnsName(), Companion_EndpointKind_.Create_TailscaleMagicDns_(), (policy).Dtor_port()))
	} else if ((policy).Dtor_fallbackLoopback()) && (_dafny.Companion_Sequence_.Contains((policy).Dtor_allowedBindKinds(), Companion_BindKind_.Create_LoopbackBind_())) {
		return Companion_EndpointDecision_.Create_UseEndpoint_(Companion_Endpoint_.Create_Endpoint_(_dafny.UnicodeSeqOfUtf8Bytes("127.0.0.1"), _dafny.UnicodeSeqOfUtf8Bytes("127.0.0.1"), Companion_EndpointKind_.Create_Loopback_(), (policy).Dtor_port()))
	} else {
		return Companion_EndpointDecision_.Create_EndpointUnavailable_()
	}
}
func (_static *CompanionStruct_Default___) DecideGeneration(facts GenerationFacts) GenerationDecision {
	if ((facts).Dtor_candidateBuilt()) && ((facts).Dtor_candidateHealthy()) {
		return Companion_GenerationDecision_.Create_PromoteCandidate_()
	} else if ((facts).Dtor_activeGenerationExists()) && ((facts).Dtor_activeGenerationHealthy()) {
		return Companion_GenerationDecision_.Create_RetainActive_()
	} else {
		return Companion_GenerationDecision_.Create_InitialGenerationFailed_()
	}
}
func (_static *CompanionStruct_Default___) DecideMachineOwner(facts MachineOwnerFacts) MachineOwnerDecision {
	if (((((facts).Dtor_sharedReceiptExactHealth()) && ((facts).Dtor_machineLeasePresent())) && ((facts).Dtor_machineLeaseReadable())) && ((facts).Dtor_machineLeaseMatchesReceipt())) && ((facts).Dtor_leaseOwnerAlive()) {
		return Companion_MachineOwnerDecision_.Create_ReuseExternalOwner_()
	} else if ((facts).Dtor_sharedReceiptExactHealth()) || (((facts).Dtor_machineLeasePresent()) && ((!((facts).Dtor_machineLeaseReadable())) || ((facts).Dtor_leaseOwnerAlive()))) {
		return Companion_MachineOwnerDecision_.Create_RefuseSplitBrain_()
	} else if ((facts).Dtor_sharedReceiptPresent()) || ((facts).Dtor_machineLeasePresent()) {
		return Companion_MachineOwnerDecision_.Create_RecoverStaleOwner_()
	} else {
		return Companion_MachineOwnerDecision_.Create_StartMachineOwner_()
	}
}
func (_static *CompanionStruct_Default___) DecideInstallation(facts InstallationFacts) InstallationDecision {
	if ((facts).Dtor_installed()) && ((facts).Dtor_launcherMatches()) {
		return Companion_InstallationDecision_.Create_KeepInstalled_()
	} else {
		return Companion_InstallationDecision_.Create_InstallLauncher_()
	}
}
func (_static *CompanionStruct_Default___) RequiresBuild(mode ServiceMode, facts EnsureFacts) bool {
	return (!((facts).Dtor_runtimeReady())) || (((mode).Is_Production()) && ((facts).Dtor_sourcesStale()))
}
func (_static *CompanionStruct_Default___) DecideEnsure(mode ServiceMode, facts EnsureFacts) EnsureDecision {
	if ((((facts).Dtor_bindingMatches()) && ((facts).Dtor_modeMatches())) && ((facts).Dtor_healthy())) && (!(Companion_Default___.RequiresBuild(mode, facts))) {
		return Companion_EnsureDecision_.Create_Reuse_()
	} else if Companion_Default___.RequiresBuild(mode, facts) {
		return Companion_EnsureDecision_.Create_ReplaceAndBuild_()
	} else {
		return Companion_EnsureDecision_.Create_Replace_()
	}
}

// End of class Default__

// Definition of datatype ServiceMode
type ServiceMode struct {
	Data_ServiceMode_
}

func (_this ServiceMode) Get_() Data_ServiceMode_ {
	return _this.Data_ServiceMode_
}

type Data_ServiceMode_ interface {
	isServiceMode()
}

type CompanionStruct_ServiceMode_ struct {
}

var Companion_ServiceMode_ = CompanionStruct_ServiceMode_{}

type ServiceMode_Production struct {
}

func (ServiceMode_Production) isServiceMode() {}

func (CompanionStruct_ServiceMode_) Create_Production_() ServiceMode {
	return ServiceMode{ServiceMode_Production{}}
}

func (_this ServiceMode) Is_Production() bool {
	_, ok := _this.Get_().(ServiceMode_Production)
	return ok
}

type ServiceMode_Development struct {
}

func (ServiceMode_Development) isServiceMode() {}

func (CompanionStruct_ServiceMode_) Create_Development_() ServiceMode {
	return ServiceMode{ServiceMode_Development{}}
}

func (_this ServiceMode) Is_Development() bool {
	_, ok := _this.Get_().(ServiceMode_Development)
	return ok
}

func (CompanionStruct_ServiceMode_) Default() ServiceMode {
	return Companion_ServiceMode_.Create_Production_()
}

func (_ CompanionStruct_ServiceMode_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_ServiceMode_.Create_Production_(), true
		case 1:
			return Companion_ServiceMode_.Create_Development_(), true
		default:
			return ServiceMode{}, false
		}
	}
}

func (_this ServiceMode) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case ServiceMode_Production:
		{
			return "ConfluxServiceLifecycle.ServiceMode.Production"
		}
	case ServiceMode_Development:
		{
			return "ConfluxServiceLifecycle.ServiceMode.Development"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ServiceMode) Equals(other ServiceMode) bool {
	switch _this.Get_().(type) {
	case ServiceMode_Production:
		{
			_, ok := other.Get_().(ServiceMode_Production)
			return ok
		}
	case ServiceMode_Development:
		{
			_, ok := other.Get_().(ServiceMode_Development)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ServiceMode) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ServiceMode)
	return ok && _this.Equals(typed)
}

func Type_ServiceMode_() _dafny.TypeDescriptor {
	return type_ServiceMode_{}
}

type type_ServiceMode_ struct {
}

func (_this type_ServiceMode_) Default() interface{} {
	return Companion_ServiceMode_.Default()
}

func (_this type_ServiceMode_) String() string {
	return "ConfluxServiceLifecycle.ServiceMode"
}
func (_this ServiceMode) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ServiceMode{}

// End of datatype ServiceMode

// Definition of datatype HealthIdentity
type HealthIdentity struct {
	Data_HealthIdentity_
}

func (_this HealthIdentity) Get_() Data_HealthIdentity_ {
	return _this.Data_HealthIdentity_
}

type Data_HealthIdentity_ interface {
	isHealthIdentity()
}

type CompanionStruct_HealthIdentity_ struct {
}

var Companion_HealthIdentity_ = CompanionStruct_HealthIdentity_{}

type HealthIdentity_HealthIdentity struct {
	Owner           _dafny.Sequence
	ProtocolVersion _dafny.Sequence
}

func (HealthIdentity_HealthIdentity) isHealthIdentity() {}

func (CompanionStruct_HealthIdentity_) Create_HealthIdentity_(Owner _dafny.Sequence, ProtocolVersion _dafny.Sequence) HealthIdentity {
	return HealthIdentity{HealthIdentity_HealthIdentity{Owner, ProtocolVersion}}
}

func (_this HealthIdentity) Is_HealthIdentity() bool {
	_, ok := _this.Get_().(HealthIdentity_HealthIdentity)
	return ok
}

func (CompanionStruct_HealthIdentity_) Default() HealthIdentity {
	return Companion_HealthIdentity_.Create_HealthIdentity_(_dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this HealthIdentity) Dtor_owner() _dafny.Sequence {
	return _this.Get_().(HealthIdentity_HealthIdentity).Owner
}

func (_this HealthIdentity) Dtor_protocolVersion() _dafny.Sequence {
	return _this.Get_().(HealthIdentity_HealthIdentity).ProtocolVersion
}

func (_this HealthIdentity) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case HealthIdentity_HealthIdentity:
		{
			return "ConfluxServiceLifecycle.HealthIdentity.HealthIdentity" + "(" + data.Owner.VerbatimString(true) + ", " + data.ProtocolVersion.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this HealthIdentity) Equals(other HealthIdentity) bool {
	switch data1 := _this.Get_().(type) {
	case HealthIdentity_HealthIdentity:
		{
			data2, ok := other.Get_().(HealthIdentity_HealthIdentity)
			return ok && data1.Owner.Equals(data2.Owner) && data1.ProtocolVersion.Equals(data2.ProtocolVersion)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this HealthIdentity) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(HealthIdentity)
	return ok && _this.Equals(typed)
}

func Type_HealthIdentity_() _dafny.TypeDescriptor {
	return type_HealthIdentity_{}
}

type type_HealthIdentity_ struct {
}

func (_this type_HealthIdentity_) Default() interface{} {
	return Companion_HealthIdentity_.Default()
}

func (_this type_HealthIdentity_) String() string {
	return "ConfluxServiceLifecycle.HealthIdentity"
}
func (_this HealthIdentity) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = HealthIdentity{}

// End of datatype HealthIdentity

// Definition of datatype BindKind
type BindKind struct {
	Data_BindKind_
}

func (_this BindKind) Get_() Data_BindKind_ {
	return _this.Data_BindKind_
}

type Data_BindKind_ interface {
	isBindKind()
}

type CompanionStruct_BindKind_ struct {
}

var Companion_BindKind_ = CompanionStruct_BindKind_{}

type BindKind_TailscaleIp struct {
}

func (BindKind_TailscaleIp) isBindKind() {}

func (CompanionStruct_BindKind_) Create_TailscaleIp_() BindKind {
	return BindKind{BindKind_TailscaleIp{}}
}

func (_this BindKind) Is_TailscaleIp() bool {
	_, ok := _this.Get_().(BindKind_TailscaleIp)
	return ok
}

type BindKind_LoopbackBind struct {
}

func (BindKind_LoopbackBind) isBindKind() {}

func (CompanionStruct_BindKind_) Create_LoopbackBind_() BindKind {
	return BindKind{BindKind_LoopbackBind{}}
}

func (_this BindKind) Is_LoopbackBind() bool {
	_, ok := _this.Get_().(BindKind_LoopbackBind)
	return ok
}

func (CompanionStruct_BindKind_) Default() BindKind {
	return Companion_BindKind_.Create_TailscaleIp_()
}

func (_ CompanionStruct_BindKind_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_BindKind_.Create_TailscaleIp_(), true
		case 1:
			return Companion_BindKind_.Create_LoopbackBind_(), true
		default:
			return BindKind{}, false
		}
	}
}

func (_this BindKind) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case BindKind_TailscaleIp:
		{
			return "ConfluxServiceLifecycle.BindKind.TailscaleIp"
		}
	case BindKind_LoopbackBind:
		{
			return "ConfluxServiceLifecycle.BindKind.LoopbackBind"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this BindKind) Equals(other BindKind) bool {
	switch _this.Get_().(type) {
	case BindKind_TailscaleIp:
		{
			_, ok := other.Get_().(BindKind_TailscaleIp)
			return ok
		}
	case BindKind_LoopbackBind:
		{
			_, ok := other.Get_().(BindKind_LoopbackBind)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this BindKind) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(BindKind)
	return ok && _this.Equals(typed)
}

func Type_BindKind_() _dafny.TypeDescriptor {
	return type_BindKind_{}
}

type type_BindKind_ struct {
}

func (_this type_BindKind_) Default() interface{} {
	return Companion_BindKind_.Default()
}

func (_this type_BindKind_) String() string {
	return "ConfluxServiceLifecycle.BindKind"
}
func (_this BindKind) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = BindKind{}

// End of datatype BindKind

// Definition of datatype EndpointPolicy
type EndpointPolicy struct {
	Data_EndpointPolicy_
}

func (_this EndpointPolicy) Get_() Data_EndpointPolicy_ {
	return _this.Data_EndpointPolicy_
}

type Data_EndpointPolicy_ interface {
	isEndpointPolicy()
}

type CompanionStruct_EndpointPolicy_ struct {
}

var Companion_EndpointPolicy_ = CompanionStruct_EndpointPolicy_{}

type EndpointPolicy_EndpointPolicy struct {
	Port                    _dafny.Int
	PreferTailscaleMagicDns bool
	FallbackLoopback        bool
	AllowedBindKinds        _dafny.Sequence
}

func (EndpointPolicy_EndpointPolicy) isEndpointPolicy() {}

func (CompanionStruct_EndpointPolicy_) Create_EndpointPolicy_(Port _dafny.Int, PreferTailscaleMagicDns bool, FallbackLoopback bool, AllowedBindKinds _dafny.Sequence) EndpointPolicy {
	return EndpointPolicy{EndpointPolicy_EndpointPolicy{Port, PreferTailscaleMagicDns, FallbackLoopback, AllowedBindKinds}}
}

func (_this EndpointPolicy) Is_EndpointPolicy() bool {
	_, ok := _this.Get_().(EndpointPolicy_EndpointPolicy)
	return ok
}

func (CompanionStruct_EndpointPolicy_) Default() EndpointPolicy {
	return Companion_EndpointPolicy_.Create_EndpointPolicy_(_dafny.Zero, false, false, _dafny.EmptySeq)
}

func (_this EndpointPolicy) Dtor_port() _dafny.Int {
	return _this.Get_().(EndpointPolicy_EndpointPolicy).Port
}

func (_this EndpointPolicy) Dtor_preferTailscaleMagicDns() bool {
	return _this.Get_().(EndpointPolicy_EndpointPolicy).PreferTailscaleMagicDns
}

func (_this EndpointPolicy) Dtor_fallbackLoopback() bool {
	return _this.Get_().(EndpointPolicy_EndpointPolicy).FallbackLoopback
}

func (_this EndpointPolicy) Dtor_allowedBindKinds() _dafny.Sequence {
	return _this.Get_().(EndpointPolicy_EndpointPolicy).AllowedBindKinds
}

func (_this EndpointPolicy) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case EndpointPolicy_EndpointPolicy:
		{
			return "ConfluxServiceLifecycle.EndpointPolicy.EndpointPolicy" + "(" + _dafny.String(data.Port) + ", " + _dafny.String(data.PreferTailscaleMagicDns) + ", " + _dafny.String(data.FallbackLoopback) + ", " + _dafny.String(data.AllowedBindKinds) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this EndpointPolicy) Equals(other EndpointPolicy) bool {
	switch data1 := _this.Get_().(type) {
	case EndpointPolicy_EndpointPolicy:
		{
			data2, ok := other.Get_().(EndpointPolicy_EndpointPolicy)
			return ok && data1.Port.Cmp(data2.Port) == 0 && data1.PreferTailscaleMagicDns == data2.PreferTailscaleMagicDns && data1.FallbackLoopback == data2.FallbackLoopback && data1.AllowedBindKinds.Equals(data2.AllowedBindKinds)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this EndpointPolicy) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(EndpointPolicy)
	return ok && _this.Equals(typed)
}

func Type_EndpointPolicy_() _dafny.TypeDescriptor {
	return type_EndpointPolicy_{}
}

type type_EndpointPolicy_ struct {
}

func (_this type_EndpointPolicy_) Default() interface{} {
	return Companion_EndpointPolicy_.Default()
}

func (_this type_EndpointPolicy_) String() string {
	return "ConfluxServiceLifecycle.EndpointPolicy"
}
func (_this EndpointPolicy) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = EndpointPolicy{}

// End of datatype EndpointPolicy

// Definition of datatype ServiceSpec
type ServiceSpec struct {
	Data_ServiceSpec_
}

func (_this ServiceSpec) Get_() Data_ServiceSpec_ {
	return _this.Data_ServiceSpec_
}

type Data_ServiceSpec_ interface {
	isServiceSpec()
}

type CompanionStruct_ServiceSpec_ struct {
}

var Companion_ServiceSpec_ = CompanionStruct_ServiceSpec_{}

type ServiceSpec_ServiceSpec struct {
	SchemaVersion     _dafny.Int
	ServiceId         _dafny.Sequence
	SourceClosure     _dafny.Sequence
	BuildGeneration   _dafny.Sequence
	RunGeneration     _dafny.Sequence
	HealthIdentity    HealthIdentity
	EndpointPolicy    EndpointPolicy
	ConsolePath       _dafny.Sequence
	ChannelPath       _dafny.Sequence
	StateRoot         _dafny.Sequence
	SupervisorAdapter _dafny.Sequence
}

func (ServiceSpec_ServiceSpec) isServiceSpec() {}

func (CompanionStruct_ServiceSpec_) Create_ServiceSpec_(SchemaVersion _dafny.Int, ServiceId _dafny.Sequence, SourceClosure _dafny.Sequence, BuildGeneration _dafny.Sequence, RunGeneration _dafny.Sequence, HealthIdentity HealthIdentity, EndpointPolicy EndpointPolicy, ConsolePath _dafny.Sequence, ChannelPath _dafny.Sequence, StateRoot _dafny.Sequence, SupervisorAdapter _dafny.Sequence) ServiceSpec {
	return ServiceSpec{ServiceSpec_ServiceSpec{SchemaVersion, ServiceId, SourceClosure, BuildGeneration, RunGeneration, HealthIdentity, EndpointPolicy, ConsolePath, ChannelPath, StateRoot, SupervisorAdapter}}
}

func (_this ServiceSpec) Is_ServiceSpec() bool {
	_, ok := _this.Get_().(ServiceSpec_ServiceSpec)
	return ok
}

func (CompanionStruct_ServiceSpec_) Default() ServiceSpec {
	return Companion_ServiceSpec_.Create_ServiceSpec_(_dafny.Zero, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, Companion_HealthIdentity_.Default(), Companion_EndpointPolicy_.Default(), _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this ServiceSpec) Dtor_schemaVersion() _dafny.Int {
	return _this.Get_().(ServiceSpec_ServiceSpec).SchemaVersion
}

func (_this ServiceSpec) Dtor_serviceId() _dafny.Sequence {
	return _this.Get_().(ServiceSpec_ServiceSpec).ServiceId
}

func (_this ServiceSpec) Dtor_sourceClosure() _dafny.Sequence {
	return _this.Get_().(ServiceSpec_ServiceSpec).SourceClosure
}

func (_this ServiceSpec) Dtor_buildGeneration() _dafny.Sequence {
	return _this.Get_().(ServiceSpec_ServiceSpec).BuildGeneration
}

func (_this ServiceSpec) Dtor_runGeneration() _dafny.Sequence {
	return _this.Get_().(ServiceSpec_ServiceSpec).RunGeneration
}

func (_this ServiceSpec) Dtor_healthIdentity() HealthIdentity {
	return _this.Get_().(ServiceSpec_ServiceSpec).HealthIdentity
}

func (_this ServiceSpec) Dtor_endpointPolicy() EndpointPolicy {
	return _this.Get_().(ServiceSpec_ServiceSpec).EndpointPolicy
}

func (_this ServiceSpec) Dtor_consolePath() _dafny.Sequence {
	return _this.Get_().(ServiceSpec_ServiceSpec).ConsolePath
}

func (_this ServiceSpec) Dtor_channelPath() _dafny.Sequence {
	return _this.Get_().(ServiceSpec_ServiceSpec).ChannelPath
}

func (_this ServiceSpec) Dtor_stateRoot() _dafny.Sequence {
	return _this.Get_().(ServiceSpec_ServiceSpec).StateRoot
}

func (_this ServiceSpec) Dtor_supervisorAdapter() _dafny.Sequence {
	return _this.Get_().(ServiceSpec_ServiceSpec).SupervisorAdapter
}

func (_this ServiceSpec) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ServiceSpec_ServiceSpec:
		{
			return "ConfluxServiceLifecycle.ServiceSpec.ServiceSpec" + "(" + _dafny.String(data.SchemaVersion) + ", " + data.ServiceId.VerbatimString(true) + ", " + _dafny.String(data.SourceClosure) + ", " + data.BuildGeneration.VerbatimString(true) + ", " + data.RunGeneration.VerbatimString(true) + ", " + _dafny.String(data.HealthIdentity) + ", " + _dafny.String(data.EndpointPolicy) + ", " + data.ConsolePath.VerbatimString(true) + ", " + data.ChannelPath.VerbatimString(true) + ", " + data.StateRoot.VerbatimString(true) + ", " + data.SupervisorAdapter.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ServiceSpec) Equals(other ServiceSpec) bool {
	switch data1 := _this.Get_().(type) {
	case ServiceSpec_ServiceSpec:
		{
			data2, ok := other.Get_().(ServiceSpec_ServiceSpec)
			return ok && data1.SchemaVersion.Cmp(data2.SchemaVersion) == 0 && data1.ServiceId.Equals(data2.ServiceId) && data1.SourceClosure.Equals(data2.SourceClosure) && data1.BuildGeneration.Equals(data2.BuildGeneration) && data1.RunGeneration.Equals(data2.RunGeneration) && data1.HealthIdentity.Equals(data2.HealthIdentity) && data1.EndpointPolicy.Equals(data2.EndpointPolicy) && data1.ConsolePath.Equals(data2.ConsolePath) && data1.ChannelPath.Equals(data2.ChannelPath) && data1.StateRoot.Equals(data2.StateRoot) && data1.SupervisorAdapter.Equals(data2.SupervisorAdapter)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ServiceSpec) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ServiceSpec)
	return ok && _this.Equals(typed)
}

func Type_ServiceSpec_() _dafny.TypeDescriptor {
	return type_ServiceSpec_{}
}

type type_ServiceSpec_ struct {
}

func (_this type_ServiceSpec_) Default() interface{} {
	return Companion_ServiceSpec_.Default()
}

func (_this type_ServiceSpec_) String() string {
	return "ConfluxServiceLifecycle.ServiceSpec"
}
func (_this ServiceSpec) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ServiceSpec{}

// End of datatype ServiceSpec

// Definition of datatype EndpointKind
type EndpointKind struct {
	Data_EndpointKind_
}

func (_this EndpointKind) Get_() Data_EndpointKind_ {
	return _this.Data_EndpointKind_
}

type Data_EndpointKind_ interface {
	isEndpointKind()
}

type CompanionStruct_EndpointKind_ struct {
}

var Companion_EndpointKind_ = CompanionStruct_EndpointKind_{}

type EndpointKind_TailscaleMagicDns struct {
}

func (EndpointKind_TailscaleMagicDns) isEndpointKind() {}

func (CompanionStruct_EndpointKind_) Create_TailscaleMagicDns_() EndpointKind {
	return EndpointKind{EndpointKind_TailscaleMagicDns{}}
}

func (_this EndpointKind) Is_TailscaleMagicDns() bool {
	_, ok := _this.Get_().(EndpointKind_TailscaleMagicDns)
	return ok
}

type EndpointKind_Loopback struct {
}

func (EndpointKind_Loopback) isEndpointKind() {}

func (CompanionStruct_EndpointKind_) Create_Loopback_() EndpointKind {
	return EndpointKind{EndpointKind_Loopback{}}
}

func (_this EndpointKind) Is_Loopback() bool {
	_, ok := _this.Get_().(EndpointKind_Loopback)
	return ok
}

func (CompanionStruct_EndpointKind_) Default() EndpointKind {
	return Companion_EndpointKind_.Create_TailscaleMagicDns_()
}

func (_ CompanionStruct_EndpointKind_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_EndpointKind_.Create_TailscaleMagicDns_(), true
		case 1:
			return Companion_EndpointKind_.Create_Loopback_(), true
		default:
			return EndpointKind{}, false
		}
	}
}

func (_this EndpointKind) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case EndpointKind_TailscaleMagicDns:
		{
			return "ConfluxServiceLifecycle.EndpointKind.TailscaleMagicDns"
		}
	case EndpointKind_Loopback:
		{
			return "ConfluxServiceLifecycle.EndpointKind.Loopback"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this EndpointKind) Equals(other EndpointKind) bool {
	switch _this.Get_().(type) {
	case EndpointKind_TailscaleMagicDns:
		{
			_, ok := other.Get_().(EndpointKind_TailscaleMagicDns)
			return ok
		}
	case EndpointKind_Loopback:
		{
			_, ok := other.Get_().(EndpointKind_Loopback)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this EndpointKind) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(EndpointKind)
	return ok && _this.Equals(typed)
}

func Type_EndpointKind_() _dafny.TypeDescriptor {
	return type_EndpointKind_{}
}

type type_EndpointKind_ struct {
}

func (_this type_EndpointKind_) Default() interface{} {
	return Companion_EndpointKind_.Default()
}

func (_this type_EndpointKind_) String() string {
	return "ConfluxServiceLifecycle.EndpointKind"
}
func (_this EndpointKind) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = EndpointKind{}

// End of datatype EndpointKind

// Definition of datatype Endpoint
type Endpoint struct {
	Data_Endpoint_
}

func (_this Endpoint) Get_() Data_Endpoint_ {
	return _this.Data_Endpoint_
}

type Data_Endpoint_ interface {
	isEndpoint()
}

type CompanionStruct_Endpoint_ struct {
}

var Companion_Endpoint_ = CompanionStruct_Endpoint_{}

type Endpoint_Endpoint struct {
	BindHost     _dafny.Sequence
	PublicHost   _dafny.Sequence
	EndpointKind EndpointKind
	Port         _dafny.Int
}

func (Endpoint_Endpoint) isEndpoint() {}

func (CompanionStruct_Endpoint_) Create_Endpoint_(BindHost _dafny.Sequence, PublicHost _dafny.Sequence, EndpointKind EndpointKind, Port _dafny.Int) Endpoint {
	return Endpoint{Endpoint_Endpoint{BindHost, PublicHost, EndpointKind, Port}}
}

func (_this Endpoint) Is_Endpoint() bool {
	_, ok := _this.Get_().(Endpoint_Endpoint)
	return ok
}

func (CompanionStruct_Endpoint_) Default() Endpoint {
	return Companion_Endpoint_.Create_Endpoint_(_dafny.EmptySeq, _dafny.EmptySeq, Companion_EndpointKind_.Default(), _dafny.Zero)
}

func (_this Endpoint) Dtor_bindHost() _dafny.Sequence {
	return _this.Get_().(Endpoint_Endpoint).BindHost
}

func (_this Endpoint) Dtor_publicHost() _dafny.Sequence {
	return _this.Get_().(Endpoint_Endpoint).PublicHost
}

func (_this Endpoint) Dtor_endpointKind() EndpointKind {
	return _this.Get_().(Endpoint_Endpoint).EndpointKind
}

func (_this Endpoint) Dtor_port() _dafny.Int {
	return _this.Get_().(Endpoint_Endpoint).Port
}

func (_this Endpoint) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Endpoint_Endpoint:
		{
			return "ConfluxServiceLifecycle.Endpoint.Endpoint" + "(" + data.BindHost.VerbatimString(true) + ", " + data.PublicHost.VerbatimString(true) + ", " + _dafny.String(data.EndpointKind) + ", " + _dafny.String(data.Port) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Endpoint) Equals(other Endpoint) bool {
	switch data1 := _this.Get_().(type) {
	case Endpoint_Endpoint:
		{
			data2, ok := other.Get_().(Endpoint_Endpoint)
			return ok && data1.BindHost.Equals(data2.BindHost) && data1.PublicHost.Equals(data2.PublicHost) && data1.EndpointKind.Equals(data2.EndpointKind) && data1.Port.Cmp(data2.Port) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Endpoint) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Endpoint)
	return ok && _this.Equals(typed)
}

func Type_Endpoint_() _dafny.TypeDescriptor {
	return type_Endpoint_{}
}

type type_Endpoint_ struct {
}

func (_this type_Endpoint_) Default() interface{} {
	return Companion_Endpoint_.Default()
}

func (_this type_Endpoint_) String() string {
	return "ConfluxServiceLifecycle.Endpoint"
}
func (_this Endpoint) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Endpoint{}

// End of datatype Endpoint

// Definition of datatype TailscaleFacts
type TailscaleFacts struct {
	Data_TailscaleFacts_
}

func (_this TailscaleFacts) Get_() Data_TailscaleFacts_ {
	return _this.Data_TailscaleFacts_
}

type Data_TailscaleFacts_ interface {
	isTailscaleFacts()
}

type CompanionStruct_TailscaleFacts_ struct {
}

var Companion_TailscaleFacts_ = CompanionStruct_TailscaleFacts_{}

type TailscaleFacts_TailscaleFacts struct {
	BackendRunning  bool
	MagicDnsEnabled bool
	SelfOnline      bool
	DnsName         _dafny.Sequence
	DnsNameValid    bool
	MatchingIp      _dafny.Sequence
	MatchingIpValid bool
}

func (TailscaleFacts_TailscaleFacts) isTailscaleFacts() {}

func (CompanionStruct_TailscaleFacts_) Create_TailscaleFacts_(BackendRunning bool, MagicDnsEnabled bool, SelfOnline bool, DnsName _dafny.Sequence, DnsNameValid bool, MatchingIp _dafny.Sequence, MatchingIpValid bool) TailscaleFacts {
	return TailscaleFacts{TailscaleFacts_TailscaleFacts{BackendRunning, MagicDnsEnabled, SelfOnline, DnsName, DnsNameValid, MatchingIp, MatchingIpValid}}
}

func (_this TailscaleFacts) Is_TailscaleFacts() bool {
	_, ok := _this.Get_().(TailscaleFacts_TailscaleFacts)
	return ok
}

func (CompanionStruct_TailscaleFacts_) Default() TailscaleFacts {
	return Companion_TailscaleFacts_.Create_TailscaleFacts_(false, false, false, _dafny.EmptySeq, false, _dafny.EmptySeq, false)
}

func (_this TailscaleFacts) Dtor_backendRunning() bool {
	return _this.Get_().(TailscaleFacts_TailscaleFacts).BackendRunning
}

func (_this TailscaleFacts) Dtor_magicDnsEnabled() bool {
	return _this.Get_().(TailscaleFacts_TailscaleFacts).MagicDnsEnabled
}

func (_this TailscaleFacts) Dtor_selfOnline() bool {
	return _this.Get_().(TailscaleFacts_TailscaleFacts).SelfOnline
}

func (_this TailscaleFacts) Dtor_dnsName() _dafny.Sequence {
	return _this.Get_().(TailscaleFacts_TailscaleFacts).DnsName
}

func (_this TailscaleFacts) Dtor_dnsNameValid() bool {
	return _this.Get_().(TailscaleFacts_TailscaleFacts).DnsNameValid
}

func (_this TailscaleFacts) Dtor_matchingIp() _dafny.Sequence {
	return _this.Get_().(TailscaleFacts_TailscaleFacts).MatchingIp
}

func (_this TailscaleFacts) Dtor_matchingIpValid() bool {
	return _this.Get_().(TailscaleFacts_TailscaleFacts).MatchingIpValid
}

func (_this TailscaleFacts) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case TailscaleFacts_TailscaleFacts:
		{
			return "ConfluxServiceLifecycle.TailscaleFacts.TailscaleFacts" + "(" + _dafny.String(data.BackendRunning) + ", " + _dafny.String(data.MagicDnsEnabled) + ", " + _dafny.String(data.SelfOnline) + ", " + data.DnsName.VerbatimString(true) + ", " + _dafny.String(data.DnsNameValid) + ", " + data.MatchingIp.VerbatimString(true) + ", " + _dafny.String(data.MatchingIpValid) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this TailscaleFacts) Equals(other TailscaleFacts) bool {
	switch data1 := _this.Get_().(type) {
	case TailscaleFacts_TailscaleFacts:
		{
			data2, ok := other.Get_().(TailscaleFacts_TailscaleFacts)
			return ok && data1.BackendRunning == data2.BackendRunning && data1.MagicDnsEnabled == data2.MagicDnsEnabled && data1.SelfOnline == data2.SelfOnline && data1.DnsName.Equals(data2.DnsName) && data1.DnsNameValid == data2.DnsNameValid && data1.MatchingIp.Equals(data2.MatchingIp) && data1.MatchingIpValid == data2.MatchingIpValid
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this TailscaleFacts) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(TailscaleFacts)
	return ok && _this.Equals(typed)
}

func Type_TailscaleFacts_() _dafny.TypeDescriptor {
	return type_TailscaleFacts_{}
}

type type_TailscaleFacts_ struct {
}

func (_this type_TailscaleFacts_) Default() interface{} {
	return Companion_TailscaleFacts_.Default()
}

func (_this type_TailscaleFacts_) String() string {
	return "ConfluxServiceLifecycle.TailscaleFacts"
}
func (_this TailscaleFacts) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = TailscaleFacts{}

// End of datatype TailscaleFacts

// Definition of datatype EndpointDecision
type EndpointDecision struct {
	Data_EndpointDecision_
}

func (_this EndpointDecision) Get_() Data_EndpointDecision_ {
	return _this.Data_EndpointDecision_
}

type Data_EndpointDecision_ interface {
	isEndpointDecision()
}

type CompanionStruct_EndpointDecision_ struct {
}

var Companion_EndpointDecision_ = CompanionStruct_EndpointDecision_{}

type EndpointDecision_UseEndpoint struct {
	Endpoint Endpoint
}

func (EndpointDecision_UseEndpoint) isEndpointDecision() {}

func (CompanionStruct_EndpointDecision_) Create_UseEndpoint_(Endpoint Endpoint) EndpointDecision {
	return EndpointDecision{EndpointDecision_UseEndpoint{Endpoint}}
}

func (_this EndpointDecision) Is_UseEndpoint() bool {
	_, ok := _this.Get_().(EndpointDecision_UseEndpoint)
	return ok
}

type EndpointDecision_EndpointUnavailable struct {
}

func (EndpointDecision_EndpointUnavailable) isEndpointDecision() {}

func (CompanionStruct_EndpointDecision_) Create_EndpointUnavailable_() EndpointDecision {
	return EndpointDecision{EndpointDecision_EndpointUnavailable{}}
}

func (_this EndpointDecision) Is_EndpointUnavailable() bool {
	_, ok := _this.Get_().(EndpointDecision_EndpointUnavailable)
	return ok
}

func (CompanionStruct_EndpointDecision_) Default() EndpointDecision {
	return Companion_EndpointDecision_.Create_UseEndpoint_(Companion_Endpoint_.Default())
}

func (_this EndpointDecision) Dtor_endpoint() Endpoint {
	return _this.Get_().(EndpointDecision_UseEndpoint).Endpoint
}

func (_this EndpointDecision) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case EndpointDecision_UseEndpoint:
		{
			return "ConfluxServiceLifecycle.EndpointDecision.UseEndpoint" + "(" + _dafny.String(data.Endpoint) + ")"
		}
	case EndpointDecision_EndpointUnavailable:
		{
			return "ConfluxServiceLifecycle.EndpointDecision.EndpointUnavailable"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this EndpointDecision) Equals(other EndpointDecision) bool {
	switch data1 := _this.Get_().(type) {
	case EndpointDecision_UseEndpoint:
		{
			data2, ok := other.Get_().(EndpointDecision_UseEndpoint)
			return ok && data1.Endpoint.Equals(data2.Endpoint)
		}
	case EndpointDecision_EndpointUnavailable:
		{
			_, ok := other.Get_().(EndpointDecision_EndpointUnavailable)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this EndpointDecision) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(EndpointDecision)
	return ok && _this.Equals(typed)
}

func Type_EndpointDecision_() _dafny.TypeDescriptor {
	return type_EndpointDecision_{}
}

type type_EndpointDecision_ struct {
}

func (_this type_EndpointDecision_) Default() interface{} {
	return Companion_EndpointDecision_.Default()
}

func (_this type_EndpointDecision_) String() string {
	return "ConfluxServiceLifecycle.EndpointDecision"
}
func (_this EndpointDecision) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = EndpointDecision{}

// End of datatype EndpointDecision

// Definition of datatype RefreshStatus
type RefreshStatus struct {
	Data_RefreshStatus_
}

func (_this RefreshStatus) Get_() Data_RefreshStatus_ {
	return _this.Data_RefreshStatus_
}

type Data_RefreshStatus_ interface {
	isRefreshStatus()
}

type CompanionStruct_RefreshStatus_ struct {
}

var Companion_RefreshStatus_ = CompanionStruct_RefreshStatus_{}

type RefreshStatus_RefreshStatus struct {
	Status              _dafny.Sequence
	AttemptedGeneration _dafny.Sequence
	Diagnostic          _dafny.Sequence
}

func (RefreshStatus_RefreshStatus) isRefreshStatus() {}

func (CompanionStruct_RefreshStatus_) Create_RefreshStatus_(Status _dafny.Sequence, AttemptedGeneration _dafny.Sequence, Diagnostic _dafny.Sequence) RefreshStatus {
	return RefreshStatus{RefreshStatus_RefreshStatus{Status, AttemptedGeneration, Diagnostic}}
}

func (_this RefreshStatus) Is_RefreshStatus() bool {
	_, ok := _this.Get_().(RefreshStatus_RefreshStatus)
	return ok
}

func (CompanionStruct_RefreshStatus_) Default() RefreshStatus {
	return Companion_RefreshStatus_.Create_RefreshStatus_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this RefreshStatus) Dtor_status() _dafny.Sequence {
	return _this.Get_().(RefreshStatus_RefreshStatus).Status
}

func (_this RefreshStatus) Dtor_attemptedGeneration() _dafny.Sequence {
	return _this.Get_().(RefreshStatus_RefreshStatus).AttemptedGeneration
}

func (_this RefreshStatus) Dtor_diagnostic() _dafny.Sequence {
	return _this.Get_().(RefreshStatus_RefreshStatus).Diagnostic
}

func (_this RefreshStatus) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RefreshStatus_RefreshStatus:
		{
			return "ConfluxServiceLifecycle.RefreshStatus.RefreshStatus" + "(" + data.Status.VerbatimString(true) + ", " + data.AttemptedGeneration.VerbatimString(true) + ", " + data.Diagnostic.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RefreshStatus) Equals(other RefreshStatus) bool {
	switch data1 := _this.Get_().(type) {
	case RefreshStatus_RefreshStatus:
		{
			data2, ok := other.Get_().(RefreshStatus_RefreshStatus)
			return ok && data1.Status.Equals(data2.Status) && data1.AttemptedGeneration.Equals(data2.AttemptedGeneration) && data1.Diagnostic.Equals(data2.Diagnostic)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RefreshStatus) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RefreshStatus)
	return ok && _this.Equals(typed)
}

func Type_RefreshStatus_() _dafny.TypeDescriptor {
	return type_RefreshStatus_{}
}

type type_RefreshStatus_ struct {
}

func (_this type_RefreshStatus_) Default() interface{} {
	return Companion_RefreshStatus_.Default()
}

func (_this type_RefreshStatus_) String() string {
	return "ConfluxServiceLifecycle.RefreshStatus"
}
func (_this RefreshStatus) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RefreshStatus{}

// End of datatype RefreshStatus

// Definition of datatype ServiceInfo
type ServiceInfo struct {
	Data_ServiceInfo_
}

func (_this ServiceInfo) Get_() Data_ServiceInfo_ {
	return _this.Data_ServiceInfo_
}

type Data_ServiceInfo_ interface {
	isServiceInfo()
}

type CompanionStruct_ServiceInfo_ struct {
}

var Companion_ServiceInfo_ = CompanionStruct_ServiceInfo_{}

type ServiceInfo_ServiceInfo struct {
	SchemaVersion  _dafny.Int
	ServiceId      _dafny.Sequence
	HealthIdentity HealthIdentity
	Disposition    _dafny.Sequence
	Status         _dafny.Sequence
	OwnerLeaseId   _dafny.Sequence
	Pid            _dafny.Int
	Generation     _dafny.Sequence
	BindHost       _dafny.Sequence
	PublicHost     _dafny.Sequence
	EndpointKind   EndpointKind
	Port           _dafny.Int
	ConsoleUrl     _dafny.Sequence
	ChannelBase    _dafny.Sequence
	Refresh        RefreshStatus
}

func (ServiceInfo_ServiceInfo) isServiceInfo() {}

func (CompanionStruct_ServiceInfo_) Create_ServiceInfo_(SchemaVersion _dafny.Int, ServiceId _dafny.Sequence, HealthIdentity HealthIdentity, Disposition _dafny.Sequence, Status _dafny.Sequence, OwnerLeaseId _dafny.Sequence, Pid _dafny.Int, Generation _dafny.Sequence, BindHost _dafny.Sequence, PublicHost _dafny.Sequence, EndpointKind EndpointKind, Port _dafny.Int, ConsoleUrl _dafny.Sequence, ChannelBase _dafny.Sequence, Refresh RefreshStatus) ServiceInfo {
	return ServiceInfo{ServiceInfo_ServiceInfo{SchemaVersion, ServiceId, HealthIdentity, Disposition, Status, OwnerLeaseId, Pid, Generation, BindHost, PublicHost, EndpointKind, Port, ConsoleUrl, ChannelBase, Refresh}}
}

func (_this ServiceInfo) Is_ServiceInfo() bool {
	_, ok := _this.Get_().(ServiceInfo_ServiceInfo)
	return ok
}

func (CompanionStruct_ServiceInfo_) Default() ServiceInfo {
	return Companion_ServiceInfo_.Create_ServiceInfo_(_dafny.Zero, _dafny.EmptySeq, Companion_HealthIdentity_.Default(), _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.Zero, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, Companion_EndpointKind_.Default(), _dafny.Zero, _dafny.EmptySeq, _dafny.EmptySeq, Companion_RefreshStatus_.Default())
}

func (_this ServiceInfo) Dtor_schemaVersion() _dafny.Int {
	return _this.Get_().(ServiceInfo_ServiceInfo).SchemaVersion
}

func (_this ServiceInfo) Dtor_serviceId() _dafny.Sequence {
	return _this.Get_().(ServiceInfo_ServiceInfo).ServiceId
}

func (_this ServiceInfo) Dtor_healthIdentity() HealthIdentity {
	return _this.Get_().(ServiceInfo_ServiceInfo).HealthIdentity
}

func (_this ServiceInfo) Dtor_disposition() _dafny.Sequence {
	return _this.Get_().(ServiceInfo_ServiceInfo).Disposition
}

func (_this ServiceInfo) Dtor_status() _dafny.Sequence {
	return _this.Get_().(ServiceInfo_ServiceInfo).Status
}

func (_this ServiceInfo) Dtor_ownerLeaseId() _dafny.Sequence {
	return _this.Get_().(ServiceInfo_ServiceInfo).OwnerLeaseId
}

func (_this ServiceInfo) Dtor_pid() _dafny.Int {
	return _this.Get_().(ServiceInfo_ServiceInfo).Pid
}

func (_this ServiceInfo) Dtor_generation() _dafny.Sequence {
	return _this.Get_().(ServiceInfo_ServiceInfo).Generation
}

func (_this ServiceInfo) Dtor_bindHost() _dafny.Sequence {
	return _this.Get_().(ServiceInfo_ServiceInfo).BindHost
}

func (_this ServiceInfo) Dtor_publicHost() _dafny.Sequence {
	return _this.Get_().(ServiceInfo_ServiceInfo).PublicHost
}

func (_this ServiceInfo) Dtor_endpointKind() EndpointKind {
	return _this.Get_().(ServiceInfo_ServiceInfo).EndpointKind
}

func (_this ServiceInfo) Dtor_port() _dafny.Int {
	return _this.Get_().(ServiceInfo_ServiceInfo).Port
}

func (_this ServiceInfo) Dtor_consoleUrl() _dafny.Sequence {
	return _this.Get_().(ServiceInfo_ServiceInfo).ConsoleUrl
}

func (_this ServiceInfo) Dtor_channelBase() _dafny.Sequence {
	return _this.Get_().(ServiceInfo_ServiceInfo).ChannelBase
}

func (_this ServiceInfo) Dtor_refresh() RefreshStatus {
	return _this.Get_().(ServiceInfo_ServiceInfo).Refresh
}

func (_this ServiceInfo) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ServiceInfo_ServiceInfo:
		{
			return "ConfluxServiceLifecycle.ServiceInfo.ServiceInfo" + "(" + _dafny.String(data.SchemaVersion) + ", " + data.ServiceId.VerbatimString(true) + ", " + _dafny.String(data.HealthIdentity) + ", " + data.Disposition.VerbatimString(true) + ", " + data.Status.VerbatimString(true) + ", " + data.OwnerLeaseId.VerbatimString(true) + ", " + _dafny.String(data.Pid) + ", " + data.Generation.VerbatimString(true) + ", " + data.BindHost.VerbatimString(true) + ", " + data.PublicHost.VerbatimString(true) + ", " + _dafny.String(data.EndpointKind) + ", " + _dafny.String(data.Port) + ", " + data.ConsoleUrl.VerbatimString(true) + ", " + data.ChannelBase.VerbatimString(true) + ", " + _dafny.String(data.Refresh) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ServiceInfo) Equals(other ServiceInfo) bool {
	switch data1 := _this.Get_().(type) {
	case ServiceInfo_ServiceInfo:
		{
			data2, ok := other.Get_().(ServiceInfo_ServiceInfo)
			return ok && data1.SchemaVersion.Cmp(data2.SchemaVersion) == 0 && data1.ServiceId.Equals(data2.ServiceId) && data1.HealthIdentity.Equals(data2.HealthIdentity) && data1.Disposition.Equals(data2.Disposition) && data1.Status.Equals(data2.Status) && data1.OwnerLeaseId.Equals(data2.OwnerLeaseId) && data1.Pid.Cmp(data2.Pid) == 0 && data1.Generation.Equals(data2.Generation) && data1.BindHost.Equals(data2.BindHost) && data1.PublicHost.Equals(data2.PublicHost) && data1.EndpointKind.Equals(data2.EndpointKind) && data1.Port.Cmp(data2.Port) == 0 && data1.ConsoleUrl.Equals(data2.ConsoleUrl) && data1.ChannelBase.Equals(data2.ChannelBase) && data1.Refresh.Equals(data2.Refresh)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ServiceInfo) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ServiceInfo)
	return ok && _this.Equals(typed)
}

func Type_ServiceInfo_() _dafny.TypeDescriptor {
	return type_ServiceInfo_{}
}

type type_ServiceInfo_ struct {
}

func (_this type_ServiceInfo_) Default() interface{} {
	return Companion_ServiceInfo_.Default()
}

func (_this type_ServiceInfo_) String() string {
	return "ConfluxServiceLifecycle.ServiceInfo"
}
func (_this ServiceInfo) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ServiceInfo{}

// End of datatype ServiceInfo

// Definition of datatype HealthObservation
type HealthObservation struct {
	Data_HealthObservation_
}

func (_this HealthObservation) Get_() Data_HealthObservation_ {
	return _this.Data_HealthObservation_
}

type Data_HealthObservation_ interface {
	isHealthObservation()
}

type CompanionStruct_HealthObservation_ struct {
}

var Companion_HealthObservation_ = CompanionStruct_HealthObservation_{}

type HealthObservation_HealthObservation struct {
	Status          _dafny.Sequence
	Owner           _dafny.Sequence
	ProtocolVersion _dafny.Sequence
	Pid             _dafny.Int
	Generation      _dafny.Sequence
	OwnerLeaseId    _dafny.Sequence
}

func (HealthObservation_HealthObservation) isHealthObservation() {}

func (CompanionStruct_HealthObservation_) Create_HealthObservation_(Status _dafny.Sequence, Owner _dafny.Sequence, ProtocolVersion _dafny.Sequence, Pid _dafny.Int, Generation _dafny.Sequence, OwnerLeaseId _dafny.Sequence) HealthObservation {
	return HealthObservation{HealthObservation_HealthObservation{Status, Owner, ProtocolVersion, Pid, Generation, OwnerLeaseId}}
}

func (_this HealthObservation) Is_HealthObservation() bool {
	_, ok := _this.Get_().(HealthObservation_HealthObservation)
	return ok
}

func (CompanionStruct_HealthObservation_) Default() HealthObservation {
	return Companion_HealthObservation_.Create_HealthObservation_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.Zero, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this HealthObservation) Dtor_status() _dafny.Sequence {
	return _this.Get_().(HealthObservation_HealthObservation).Status
}

func (_this HealthObservation) Dtor_owner() _dafny.Sequence {
	return _this.Get_().(HealthObservation_HealthObservation).Owner
}

func (_this HealthObservation) Dtor_protocolVersion() _dafny.Sequence {
	return _this.Get_().(HealthObservation_HealthObservation).ProtocolVersion
}

func (_this HealthObservation) Dtor_pid() _dafny.Int {
	return _this.Get_().(HealthObservation_HealthObservation).Pid
}

func (_this HealthObservation) Dtor_generation() _dafny.Sequence {
	return _this.Get_().(HealthObservation_HealthObservation).Generation
}

func (_this HealthObservation) Dtor_ownerLeaseId() _dafny.Sequence {
	return _this.Get_().(HealthObservation_HealthObservation).OwnerLeaseId
}

func (_this HealthObservation) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case HealthObservation_HealthObservation:
		{
			return "ConfluxServiceLifecycle.HealthObservation.HealthObservation" + "(" + data.Status.VerbatimString(true) + ", " + data.Owner.VerbatimString(true) + ", " + data.ProtocolVersion.VerbatimString(true) + ", " + _dafny.String(data.Pid) + ", " + data.Generation.VerbatimString(true) + ", " + data.OwnerLeaseId.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this HealthObservation) Equals(other HealthObservation) bool {
	switch data1 := _this.Get_().(type) {
	case HealthObservation_HealthObservation:
		{
			data2, ok := other.Get_().(HealthObservation_HealthObservation)
			return ok && data1.Status.Equals(data2.Status) && data1.Owner.Equals(data2.Owner) && data1.ProtocolVersion.Equals(data2.ProtocolVersion) && data1.Pid.Cmp(data2.Pid) == 0 && data1.Generation.Equals(data2.Generation) && data1.OwnerLeaseId.Equals(data2.OwnerLeaseId)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this HealthObservation) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(HealthObservation)
	return ok && _this.Equals(typed)
}

func Type_HealthObservation_() _dafny.TypeDescriptor {
	return type_HealthObservation_{}
}

type type_HealthObservation_ struct {
}

func (_this type_HealthObservation_) Default() interface{} {
	return Companion_HealthObservation_.Default()
}

func (_this type_HealthObservation_) String() string {
	return "ConfluxServiceLifecycle.HealthObservation"
}
func (_this HealthObservation) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = HealthObservation{}

// End of datatype HealthObservation

// Definition of datatype GenerationFacts
type GenerationFacts struct {
	Data_GenerationFacts_
}

func (_this GenerationFacts) Get_() Data_GenerationFacts_ {
	return _this.Data_GenerationFacts_
}

type Data_GenerationFacts_ interface {
	isGenerationFacts()
}

type CompanionStruct_GenerationFacts_ struct {
}

var Companion_GenerationFacts_ = CompanionStruct_GenerationFacts_{}

type GenerationFacts_GenerationFacts struct {
	ActiveGenerationExists  bool
	ActiveGenerationHealthy bool
	CandidateBuilt          bool
	CandidateHealthy        bool
}

func (GenerationFacts_GenerationFacts) isGenerationFacts() {}

func (CompanionStruct_GenerationFacts_) Create_GenerationFacts_(ActiveGenerationExists bool, ActiveGenerationHealthy bool, CandidateBuilt bool, CandidateHealthy bool) GenerationFacts {
	return GenerationFacts{GenerationFacts_GenerationFacts{ActiveGenerationExists, ActiveGenerationHealthy, CandidateBuilt, CandidateHealthy}}
}

func (_this GenerationFacts) Is_GenerationFacts() bool {
	_, ok := _this.Get_().(GenerationFacts_GenerationFacts)
	return ok
}

func (CompanionStruct_GenerationFacts_) Default() GenerationFacts {
	return Companion_GenerationFacts_.Create_GenerationFacts_(false, false, false, false)
}

func (_this GenerationFacts) Dtor_activeGenerationExists() bool {
	return _this.Get_().(GenerationFacts_GenerationFacts).ActiveGenerationExists
}

func (_this GenerationFacts) Dtor_activeGenerationHealthy() bool {
	return _this.Get_().(GenerationFacts_GenerationFacts).ActiveGenerationHealthy
}

func (_this GenerationFacts) Dtor_candidateBuilt() bool {
	return _this.Get_().(GenerationFacts_GenerationFacts).CandidateBuilt
}

func (_this GenerationFacts) Dtor_candidateHealthy() bool {
	return _this.Get_().(GenerationFacts_GenerationFacts).CandidateHealthy
}

func (_this GenerationFacts) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case GenerationFacts_GenerationFacts:
		{
			return "ConfluxServiceLifecycle.GenerationFacts.GenerationFacts" + "(" + _dafny.String(data.ActiveGenerationExists) + ", " + _dafny.String(data.ActiveGenerationHealthy) + ", " + _dafny.String(data.CandidateBuilt) + ", " + _dafny.String(data.CandidateHealthy) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GenerationFacts) Equals(other GenerationFacts) bool {
	switch data1 := _this.Get_().(type) {
	case GenerationFacts_GenerationFacts:
		{
			data2, ok := other.Get_().(GenerationFacts_GenerationFacts)
			return ok && data1.ActiveGenerationExists == data2.ActiveGenerationExists && data1.ActiveGenerationHealthy == data2.ActiveGenerationHealthy && data1.CandidateBuilt == data2.CandidateBuilt && data1.CandidateHealthy == data2.CandidateHealthy
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GenerationFacts) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GenerationFacts)
	return ok && _this.Equals(typed)
}

func Type_GenerationFacts_() _dafny.TypeDescriptor {
	return type_GenerationFacts_{}
}

type type_GenerationFacts_ struct {
}

func (_this type_GenerationFacts_) Default() interface{} {
	return Companion_GenerationFacts_.Default()
}

func (_this type_GenerationFacts_) String() string {
	return "ConfluxServiceLifecycle.GenerationFacts"
}
func (_this GenerationFacts) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = GenerationFacts{}

// End of datatype GenerationFacts

// Definition of datatype GenerationDecision
type GenerationDecision struct {
	Data_GenerationDecision_
}

func (_this GenerationDecision) Get_() Data_GenerationDecision_ {
	return _this.Data_GenerationDecision_
}

type Data_GenerationDecision_ interface {
	isGenerationDecision()
}

type CompanionStruct_GenerationDecision_ struct {
}

var Companion_GenerationDecision_ = CompanionStruct_GenerationDecision_{}

type GenerationDecision_PromoteCandidate struct {
}

func (GenerationDecision_PromoteCandidate) isGenerationDecision() {}

func (CompanionStruct_GenerationDecision_) Create_PromoteCandidate_() GenerationDecision {
	return GenerationDecision{GenerationDecision_PromoteCandidate{}}
}

func (_this GenerationDecision) Is_PromoteCandidate() bool {
	_, ok := _this.Get_().(GenerationDecision_PromoteCandidate)
	return ok
}

type GenerationDecision_RetainActive struct {
}

func (GenerationDecision_RetainActive) isGenerationDecision() {}

func (CompanionStruct_GenerationDecision_) Create_RetainActive_() GenerationDecision {
	return GenerationDecision{GenerationDecision_RetainActive{}}
}

func (_this GenerationDecision) Is_RetainActive() bool {
	_, ok := _this.Get_().(GenerationDecision_RetainActive)
	return ok
}

type GenerationDecision_InitialGenerationFailed struct {
}

func (GenerationDecision_InitialGenerationFailed) isGenerationDecision() {}

func (CompanionStruct_GenerationDecision_) Create_InitialGenerationFailed_() GenerationDecision {
	return GenerationDecision{GenerationDecision_InitialGenerationFailed{}}
}

func (_this GenerationDecision) Is_InitialGenerationFailed() bool {
	_, ok := _this.Get_().(GenerationDecision_InitialGenerationFailed)
	return ok
}

func (CompanionStruct_GenerationDecision_) Default() GenerationDecision {
	return Companion_GenerationDecision_.Create_PromoteCandidate_()
}

func (_ CompanionStruct_GenerationDecision_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_GenerationDecision_.Create_PromoteCandidate_(), true
		case 1:
			return Companion_GenerationDecision_.Create_RetainActive_(), true
		case 2:
			return Companion_GenerationDecision_.Create_InitialGenerationFailed_(), true
		default:
			return GenerationDecision{}, false
		}
	}
}

func (_this GenerationDecision) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case GenerationDecision_PromoteCandidate:
		{
			return "ConfluxServiceLifecycle.GenerationDecision.PromoteCandidate"
		}
	case GenerationDecision_RetainActive:
		{
			return "ConfluxServiceLifecycle.GenerationDecision.RetainActive"
		}
	case GenerationDecision_InitialGenerationFailed:
		{
			return "ConfluxServiceLifecycle.GenerationDecision.InitialGenerationFailed"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GenerationDecision) Equals(other GenerationDecision) bool {
	switch _this.Get_().(type) {
	case GenerationDecision_PromoteCandidate:
		{
			_, ok := other.Get_().(GenerationDecision_PromoteCandidate)
			return ok
		}
	case GenerationDecision_RetainActive:
		{
			_, ok := other.Get_().(GenerationDecision_RetainActive)
			return ok
		}
	case GenerationDecision_InitialGenerationFailed:
		{
			_, ok := other.Get_().(GenerationDecision_InitialGenerationFailed)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GenerationDecision) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GenerationDecision)
	return ok && _this.Equals(typed)
}

func Type_GenerationDecision_() _dafny.TypeDescriptor {
	return type_GenerationDecision_{}
}

type type_GenerationDecision_ struct {
}

func (_this type_GenerationDecision_) Default() interface{} {
	return Companion_GenerationDecision_.Default()
}

func (_this type_GenerationDecision_) String() string {
	return "ConfluxServiceLifecycle.GenerationDecision"
}
func (_this GenerationDecision) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = GenerationDecision{}

// End of datatype GenerationDecision

// Definition of datatype MachineOwnerFacts
type MachineOwnerFacts struct {
	Data_MachineOwnerFacts_
}

func (_this MachineOwnerFacts) Get_() Data_MachineOwnerFacts_ {
	return _this.Data_MachineOwnerFacts_
}

type Data_MachineOwnerFacts_ interface {
	isMachineOwnerFacts()
}

type CompanionStruct_MachineOwnerFacts_ struct {
}

var Companion_MachineOwnerFacts_ = CompanionStruct_MachineOwnerFacts_{}

type MachineOwnerFacts_MachineOwnerFacts struct {
	SharedReceiptPresent       bool
	SharedReceiptExactHealth   bool
	MachineLeasePresent        bool
	MachineLeaseReadable       bool
	MachineLeaseMatchesReceipt bool
	LeaseOwnerAlive            bool
}

func (MachineOwnerFacts_MachineOwnerFacts) isMachineOwnerFacts() {}

func (CompanionStruct_MachineOwnerFacts_) Create_MachineOwnerFacts_(SharedReceiptPresent bool, SharedReceiptExactHealth bool, MachineLeasePresent bool, MachineLeaseReadable bool, MachineLeaseMatchesReceipt bool, LeaseOwnerAlive bool) MachineOwnerFacts {
	return MachineOwnerFacts{MachineOwnerFacts_MachineOwnerFacts{SharedReceiptPresent, SharedReceiptExactHealth, MachineLeasePresent, MachineLeaseReadable, MachineLeaseMatchesReceipt, LeaseOwnerAlive}}
}

func (_this MachineOwnerFacts) Is_MachineOwnerFacts() bool {
	_, ok := _this.Get_().(MachineOwnerFacts_MachineOwnerFacts)
	return ok
}

func (CompanionStruct_MachineOwnerFacts_) Default() MachineOwnerFacts {
	return Companion_MachineOwnerFacts_.Create_MachineOwnerFacts_(false, false, false, false, false, false)
}

func (_this MachineOwnerFacts) Dtor_sharedReceiptPresent() bool {
	return _this.Get_().(MachineOwnerFacts_MachineOwnerFacts).SharedReceiptPresent
}

func (_this MachineOwnerFacts) Dtor_sharedReceiptExactHealth() bool {
	return _this.Get_().(MachineOwnerFacts_MachineOwnerFacts).SharedReceiptExactHealth
}

func (_this MachineOwnerFacts) Dtor_machineLeasePresent() bool {
	return _this.Get_().(MachineOwnerFacts_MachineOwnerFacts).MachineLeasePresent
}

func (_this MachineOwnerFacts) Dtor_machineLeaseReadable() bool {
	return _this.Get_().(MachineOwnerFacts_MachineOwnerFacts).MachineLeaseReadable
}

func (_this MachineOwnerFacts) Dtor_machineLeaseMatchesReceipt() bool {
	return _this.Get_().(MachineOwnerFacts_MachineOwnerFacts).MachineLeaseMatchesReceipt
}

func (_this MachineOwnerFacts) Dtor_leaseOwnerAlive() bool {
	return _this.Get_().(MachineOwnerFacts_MachineOwnerFacts).LeaseOwnerAlive
}

func (_this MachineOwnerFacts) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case MachineOwnerFacts_MachineOwnerFacts:
		{
			return "ConfluxServiceLifecycle.MachineOwnerFacts.MachineOwnerFacts" + "(" + _dafny.String(data.SharedReceiptPresent) + ", " + _dafny.String(data.SharedReceiptExactHealth) + ", " + _dafny.String(data.MachineLeasePresent) + ", " + _dafny.String(data.MachineLeaseReadable) + ", " + _dafny.String(data.MachineLeaseMatchesReceipt) + ", " + _dafny.String(data.LeaseOwnerAlive) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this MachineOwnerFacts) Equals(other MachineOwnerFacts) bool {
	switch data1 := _this.Get_().(type) {
	case MachineOwnerFacts_MachineOwnerFacts:
		{
			data2, ok := other.Get_().(MachineOwnerFacts_MachineOwnerFacts)
			return ok && data1.SharedReceiptPresent == data2.SharedReceiptPresent && data1.SharedReceiptExactHealth == data2.SharedReceiptExactHealth && data1.MachineLeasePresent == data2.MachineLeasePresent && data1.MachineLeaseReadable == data2.MachineLeaseReadable && data1.MachineLeaseMatchesReceipt == data2.MachineLeaseMatchesReceipt && data1.LeaseOwnerAlive == data2.LeaseOwnerAlive
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this MachineOwnerFacts) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(MachineOwnerFacts)
	return ok && _this.Equals(typed)
}

func Type_MachineOwnerFacts_() _dafny.TypeDescriptor {
	return type_MachineOwnerFacts_{}
}

type type_MachineOwnerFacts_ struct {
}

func (_this type_MachineOwnerFacts_) Default() interface{} {
	return Companion_MachineOwnerFacts_.Default()
}

func (_this type_MachineOwnerFacts_) String() string {
	return "ConfluxServiceLifecycle.MachineOwnerFacts"
}
func (_this MachineOwnerFacts) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = MachineOwnerFacts{}

// End of datatype MachineOwnerFacts

// Definition of datatype MachineOwnerDecision
type MachineOwnerDecision struct {
	Data_MachineOwnerDecision_
}

func (_this MachineOwnerDecision) Get_() Data_MachineOwnerDecision_ {
	return _this.Data_MachineOwnerDecision_
}

type Data_MachineOwnerDecision_ interface {
	isMachineOwnerDecision()
}

type CompanionStruct_MachineOwnerDecision_ struct {
}

var Companion_MachineOwnerDecision_ = CompanionStruct_MachineOwnerDecision_{}

type MachineOwnerDecision_ReuseExternalOwner struct {
}

func (MachineOwnerDecision_ReuseExternalOwner) isMachineOwnerDecision() {}

func (CompanionStruct_MachineOwnerDecision_) Create_ReuseExternalOwner_() MachineOwnerDecision {
	return MachineOwnerDecision{MachineOwnerDecision_ReuseExternalOwner{}}
}

func (_this MachineOwnerDecision) Is_ReuseExternalOwner() bool {
	_, ok := _this.Get_().(MachineOwnerDecision_ReuseExternalOwner)
	return ok
}

type MachineOwnerDecision_StartMachineOwner struct {
}

func (MachineOwnerDecision_StartMachineOwner) isMachineOwnerDecision() {}

func (CompanionStruct_MachineOwnerDecision_) Create_StartMachineOwner_() MachineOwnerDecision {
	return MachineOwnerDecision{MachineOwnerDecision_StartMachineOwner{}}
}

func (_this MachineOwnerDecision) Is_StartMachineOwner() bool {
	_, ok := _this.Get_().(MachineOwnerDecision_StartMachineOwner)
	return ok
}

type MachineOwnerDecision_RecoverStaleOwner struct {
}

func (MachineOwnerDecision_RecoverStaleOwner) isMachineOwnerDecision() {}

func (CompanionStruct_MachineOwnerDecision_) Create_RecoverStaleOwner_() MachineOwnerDecision {
	return MachineOwnerDecision{MachineOwnerDecision_RecoverStaleOwner{}}
}

func (_this MachineOwnerDecision) Is_RecoverStaleOwner() bool {
	_, ok := _this.Get_().(MachineOwnerDecision_RecoverStaleOwner)
	return ok
}

type MachineOwnerDecision_RefuseSplitBrain struct {
}

func (MachineOwnerDecision_RefuseSplitBrain) isMachineOwnerDecision() {}

func (CompanionStruct_MachineOwnerDecision_) Create_RefuseSplitBrain_() MachineOwnerDecision {
	return MachineOwnerDecision{MachineOwnerDecision_RefuseSplitBrain{}}
}

func (_this MachineOwnerDecision) Is_RefuseSplitBrain() bool {
	_, ok := _this.Get_().(MachineOwnerDecision_RefuseSplitBrain)
	return ok
}

func (CompanionStruct_MachineOwnerDecision_) Default() MachineOwnerDecision {
	return Companion_MachineOwnerDecision_.Create_ReuseExternalOwner_()
}

func (_ CompanionStruct_MachineOwnerDecision_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_MachineOwnerDecision_.Create_ReuseExternalOwner_(), true
		case 1:
			return Companion_MachineOwnerDecision_.Create_StartMachineOwner_(), true
		case 2:
			return Companion_MachineOwnerDecision_.Create_RecoverStaleOwner_(), true
		case 3:
			return Companion_MachineOwnerDecision_.Create_RefuseSplitBrain_(), true
		default:
			return MachineOwnerDecision{}, false
		}
	}
}

func (_this MachineOwnerDecision) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case MachineOwnerDecision_ReuseExternalOwner:
		{
			return "ConfluxServiceLifecycle.MachineOwnerDecision.ReuseExternalOwner"
		}
	case MachineOwnerDecision_StartMachineOwner:
		{
			return "ConfluxServiceLifecycle.MachineOwnerDecision.StartMachineOwner"
		}
	case MachineOwnerDecision_RecoverStaleOwner:
		{
			return "ConfluxServiceLifecycle.MachineOwnerDecision.RecoverStaleOwner"
		}
	case MachineOwnerDecision_RefuseSplitBrain:
		{
			return "ConfluxServiceLifecycle.MachineOwnerDecision.RefuseSplitBrain"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this MachineOwnerDecision) Equals(other MachineOwnerDecision) bool {
	switch _this.Get_().(type) {
	case MachineOwnerDecision_ReuseExternalOwner:
		{
			_, ok := other.Get_().(MachineOwnerDecision_ReuseExternalOwner)
			return ok
		}
	case MachineOwnerDecision_StartMachineOwner:
		{
			_, ok := other.Get_().(MachineOwnerDecision_StartMachineOwner)
			return ok
		}
	case MachineOwnerDecision_RecoverStaleOwner:
		{
			_, ok := other.Get_().(MachineOwnerDecision_RecoverStaleOwner)
			return ok
		}
	case MachineOwnerDecision_RefuseSplitBrain:
		{
			_, ok := other.Get_().(MachineOwnerDecision_RefuseSplitBrain)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this MachineOwnerDecision) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(MachineOwnerDecision)
	return ok && _this.Equals(typed)
}

func Type_MachineOwnerDecision_() _dafny.TypeDescriptor {
	return type_MachineOwnerDecision_{}
}

type type_MachineOwnerDecision_ struct {
}

func (_this type_MachineOwnerDecision_) Default() interface{} {
	return Companion_MachineOwnerDecision_.Default()
}

func (_this type_MachineOwnerDecision_) String() string {
	return "ConfluxServiceLifecycle.MachineOwnerDecision"
}
func (_this MachineOwnerDecision) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = MachineOwnerDecision{}

// End of datatype MachineOwnerDecision

// Definition of datatype InstallationFacts
type InstallationFacts struct {
	Data_InstallationFacts_
}

func (_this InstallationFacts) Get_() Data_InstallationFacts_ {
	return _this.Data_InstallationFacts_
}

type Data_InstallationFacts_ interface {
	isInstallationFacts()
}

type CompanionStruct_InstallationFacts_ struct {
}

var Companion_InstallationFacts_ = CompanionStruct_InstallationFacts_{}

type InstallationFacts_InstallationFacts struct {
	Installed       bool
	LauncherMatches bool
}

func (InstallationFacts_InstallationFacts) isInstallationFacts() {}

func (CompanionStruct_InstallationFacts_) Create_InstallationFacts_(Installed bool, LauncherMatches bool) InstallationFacts {
	return InstallationFacts{InstallationFacts_InstallationFacts{Installed, LauncherMatches}}
}

func (_this InstallationFacts) Is_InstallationFacts() bool {
	_, ok := _this.Get_().(InstallationFacts_InstallationFacts)
	return ok
}

func (CompanionStruct_InstallationFacts_) Default() InstallationFacts {
	return Companion_InstallationFacts_.Create_InstallationFacts_(false, false)
}

func (_this InstallationFacts) Dtor_installed() bool {
	return _this.Get_().(InstallationFacts_InstallationFacts).Installed
}

func (_this InstallationFacts) Dtor_launcherMatches() bool {
	return _this.Get_().(InstallationFacts_InstallationFacts).LauncherMatches
}

func (_this InstallationFacts) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case InstallationFacts_InstallationFacts:
		{
			return "ConfluxServiceLifecycle.InstallationFacts.InstallationFacts" + "(" + _dafny.String(data.Installed) + ", " + _dafny.String(data.LauncherMatches) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this InstallationFacts) Equals(other InstallationFacts) bool {
	switch data1 := _this.Get_().(type) {
	case InstallationFacts_InstallationFacts:
		{
			data2, ok := other.Get_().(InstallationFacts_InstallationFacts)
			return ok && data1.Installed == data2.Installed && data1.LauncherMatches == data2.LauncherMatches
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this InstallationFacts) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(InstallationFacts)
	return ok && _this.Equals(typed)
}

func Type_InstallationFacts_() _dafny.TypeDescriptor {
	return type_InstallationFacts_{}
}

type type_InstallationFacts_ struct {
}

func (_this type_InstallationFacts_) Default() interface{} {
	return Companion_InstallationFacts_.Default()
}

func (_this type_InstallationFacts_) String() string {
	return "ConfluxServiceLifecycle.InstallationFacts"
}
func (_this InstallationFacts) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = InstallationFacts{}

// End of datatype InstallationFacts

// Definition of datatype InstallationDecision
type InstallationDecision struct {
	Data_InstallationDecision_
}

func (_this InstallationDecision) Get_() Data_InstallationDecision_ {
	return _this.Data_InstallationDecision_
}

type Data_InstallationDecision_ interface {
	isInstallationDecision()
}

type CompanionStruct_InstallationDecision_ struct {
}

var Companion_InstallationDecision_ = CompanionStruct_InstallationDecision_{}

type InstallationDecision_KeepInstalled struct {
}

func (InstallationDecision_KeepInstalled) isInstallationDecision() {}

func (CompanionStruct_InstallationDecision_) Create_KeepInstalled_() InstallationDecision {
	return InstallationDecision{InstallationDecision_KeepInstalled{}}
}

func (_this InstallationDecision) Is_KeepInstalled() bool {
	_, ok := _this.Get_().(InstallationDecision_KeepInstalled)
	return ok
}

type InstallationDecision_InstallLauncher struct {
}

func (InstallationDecision_InstallLauncher) isInstallationDecision() {}

func (CompanionStruct_InstallationDecision_) Create_InstallLauncher_() InstallationDecision {
	return InstallationDecision{InstallationDecision_InstallLauncher{}}
}

func (_this InstallationDecision) Is_InstallLauncher() bool {
	_, ok := _this.Get_().(InstallationDecision_InstallLauncher)
	return ok
}

func (CompanionStruct_InstallationDecision_) Default() InstallationDecision {
	return Companion_InstallationDecision_.Create_KeepInstalled_()
}

func (_ CompanionStruct_InstallationDecision_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_InstallationDecision_.Create_KeepInstalled_(), true
		case 1:
			return Companion_InstallationDecision_.Create_InstallLauncher_(), true
		default:
			return InstallationDecision{}, false
		}
	}
}

func (_this InstallationDecision) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case InstallationDecision_KeepInstalled:
		{
			return "ConfluxServiceLifecycle.InstallationDecision.KeepInstalled"
		}
	case InstallationDecision_InstallLauncher:
		{
			return "ConfluxServiceLifecycle.InstallationDecision.InstallLauncher"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this InstallationDecision) Equals(other InstallationDecision) bool {
	switch _this.Get_().(type) {
	case InstallationDecision_KeepInstalled:
		{
			_, ok := other.Get_().(InstallationDecision_KeepInstalled)
			return ok
		}
	case InstallationDecision_InstallLauncher:
		{
			_, ok := other.Get_().(InstallationDecision_InstallLauncher)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this InstallationDecision) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(InstallationDecision)
	return ok && _this.Equals(typed)
}

func Type_InstallationDecision_() _dafny.TypeDescriptor {
	return type_InstallationDecision_{}
}

type type_InstallationDecision_ struct {
}

func (_this type_InstallationDecision_) Default() interface{} {
	return Companion_InstallationDecision_.Default()
}

func (_this type_InstallationDecision_) String() string {
	return "ConfluxServiceLifecycle.InstallationDecision"
}
func (_this InstallationDecision) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = InstallationDecision{}

// End of datatype InstallationDecision

// Definition of datatype EnsureFacts
type EnsureFacts struct {
	Data_EnsureFacts_
}

func (_this EnsureFacts) Get_() Data_EnsureFacts_ {
	return _this.Data_EnsureFacts_
}

type Data_EnsureFacts_ interface {
	isEnsureFacts()
}

type CompanionStruct_EnsureFacts_ struct {
}

var Companion_EnsureFacts_ = CompanionStruct_EnsureFacts_{}

type EnsureFacts_EnsureFacts struct {
	RuntimeReady   bool
	SourcesStale   bool
	BindingMatches bool
	ModeMatches    bool
	Healthy        bool
}

func (EnsureFacts_EnsureFacts) isEnsureFacts() {}

func (CompanionStruct_EnsureFacts_) Create_EnsureFacts_(RuntimeReady bool, SourcesStale bool, BindingMatches bool, ModeMatches bool, Healthy bool) EnsureFacts {
	return EnsureFacts{EnsureFacts_EnsureFacts{RuntimeReady, SourcesStale, BindingMatches, ModeMatches, Healthy}}
}

func (_this EnsureFacts) Is_EnsureFacts() bool {
	_, ok := _this.Get_().(EnsureFacts_EnsureFacts)
	return ok
}

func (CompanionStruct_EnsureFacts_) Default() EnsureFacts {
	return Companion_EnsureFacts_.Create_EnsureFacts_(false, false, false, false, false)
}

func (_this EnsureFacts) Dtor_runtimeReady() bool {
	return _this.Get_().(EnsureFacts_EnsureFacts).RuntimeReady
}

func (_this EnsureFacts) Dtor_sourcesStale() bool {
	return _this.Get_().(EnsureFacts_EnsureFacts).SourcesStale
}

func (_this EnsureFacts) Dtor_bindingMatches() bool {
	return _this.Get_().(EnsureFacts_EnsureFacts).BindingMatches
}

func (_this EnsureFacts) Dtor_modeMatches() bool {
	return _this.Get_().(EnsureFacts_EnsureFacts).ModeMatches
}

func (_this EnsureFacts) Dtor_healthy() bool {
	return _this.Get_().(EnsureFacts_EnsureFacts).Healthy
}

func (_this EnsureFacts) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case EnsureFacts_EnsureFacts:
		{
			return "ConfluxServiceLifecycle.EnsureFacts.EnsureFacts" + "(" + _dafny.String(data.RuntimeReady) + ", " + _dafny.String(data.SourcesStale) + ", " + _dafny.String(data.BindingMatches) + ", " + _dafny.String(data.ModeMatches) + ", " + _dafny.String(data.Healthy) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this EnsureFacts) Equals(other EnsureFacts) bool {
	switch data1 := _this.Get_().(type) {
	case EnsureFacts_EnsureFacts:
		{
			data2, ok := other.Get_().(EnsureFacts_EnsureFacts)
			return ok && data1.RuntimeReady == data2.RuntimeReady && data1.SourcesStale == data2.SourcesStale && data1.BindingMatches == data2.BindingMatches && data1.ModeMatches == data2.ModeMatches && data1.Healthy == data2.Healthy
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this EnsureFacts) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(EnsureFacts)
	return ok && _this.Equals(typed)
}

func Type_EnsureFacts_() _dafny.TypeDescriptor {
	return type_EnsureFacts_{}
}

type type_EnsureFacts_ struct {
}

func (_this type_EnsureFacts_) Default() interface{} {
	return Companion_EnsureFacts_.Default()
}

func (_this type_EnsureFacts_) String() string {
	return "ConfluxServiceLifecycle.EnsureFacts"
}
func (_this EnsureFacts) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = EnsureFacts{}

// End of datatype EnsureFacts

// Definition of datatype EnsureDecision
type EnsureDecision struct {
	Data_EnsureDecision_
}

func (_this EnsureDecision) Get_() Data_EnsureDecision_ {
	return _this.Data_EnsureDecision_
}

type Data_EnsureDecision_ interface {
	isEnsureDecision()
}

type CompanionStruct_EnsureDecision_ struct {
}

var Companion_EnsureDecision_ = CompanionStruct_EnsureDecision_{}

type EnsureDecision_Reuse struct {
}

func (EnsureDecision_Reuse) isEnsureDecision() {}

func (CompanionStruct_EnsureDecision_) Create_Reuse_() EnsureDecision {
	return EnsureDecision{EnsureDecision_Reuse{}}
}

func (_this EnsureDecision) Is_Reuse() bool {
	_, ok := _this.Get_().(EnsureDecision_Reuse)
	return ok
}

type EnsureDecision_Replace struct {
}

func (EnsureDecision_Replace) isEnsureDecision() {}

func (CompanionStruct_EnsureDecision_) Create_Replace_() EnsureDecision {
	return EnsureDecision{EnsureDecision_Replace{}}
}

func (_this EnsureDecision) Is_Replace() bool {
	_, ok := _this.Get_().(EnsureDecision_Replace)
	return ok
}

type EnsureDecision_ReplaceAndBuild struct {
}

func (EnsureDecision_ReplaceAndBuild) isEnsureDecision() {}

func (CompanionStruct_EnsureDecision_) Create_ReplaceAndBuild_() EnsureDecision {
	return EnsureDecision{EnsureDecision_ReplaceAndBuild{}}
}

func (_this EnsureDecision) Is_ReplaceAndBuild() bool {
	_, ok := _this.Get_().(EnsureDecision_ReplaceAndBuild)
	return ok
}

func (CompanionStruct_EnsureDecision_) Default() EnsureDecision {
	return Companion_EnsureDecision_.Create_Reuse_()
}

func (_ CompanionStruct_EnsureDecision_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_EnsureDecision_.Create_Reuse_(), true
		case 1:
			return Companion_EnsureDecision_.Create_Replace_(), true
		case 2:
			return Companion_EnsureDecision_.Create_ReplaceAndBuild_(), true
		default:
			return EnsureDecision{}, false
		}
	}
}

func (_this EnsureDecision) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case EnsureDecision_Reuse:
		{
			return "ConfluxServiceLifecycle.EnsureDecision.Reuse"
		}
	case EnsureDecision_Replace:
		{
			return "ConfluxServiceLifecycle.EnsureDecision.Replace"
		}
	case EnsureDecision_ReplaceAndBuild:
		{
			return "ConfluxServiceLifecycle.EnsureDecision.ReplaceAndBuild"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this EnsureDecision) Equals(other EnsureDecision) bool {
	switch _this.Get_().(type) {
	case EnsureDecision_Reuse:
		{
			_, ok := other.Get_().(EnsureDecision_Reuse)
			return ok
		}
	case EnsureDecision_Replace:
		{
			_, ok := other.Get_().(EnsureDecision_Replace)
			return ok
		}
	case EnsureDecision_ReplaceAndBuild:
		{
			_, ok := other.Get_().(EnsureDecision_ReplaceAndBuild)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this EnsureDecision) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(EnsureDecision)
	return ok && _this.Equals(typed)
}

func Type_EnsureDecision_() _dafny.TypeDescriptor {
	return type_EnsureDecision_{}
}

type type_EnsureDecision_ struct {
}

func (_this type_EnsureDecision_) Default() interface{} {
	return Companion_EnsureDecision_.Default()
}

func (_this type_EnsureDecision_) String() string {
	return "ConfluxServiceLifecycle.EnsureDecision"
}
func (_this EnsureDecision) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = EnsureDecision{}

// End of datatype EnsureDecision
