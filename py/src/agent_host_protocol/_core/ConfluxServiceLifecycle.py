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
import ConfluxResourceSupervisor as ConfluxResourceSupervisor
import ConfluxSocketCapability as ConfluxSocketCapability
import ConfluxWebSocketCapability as ConfluxWebSocketCapability
import ConfluxHttpCapability as ConfluxHttpCapability

# Module: ConfluxServiceLifecycle

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def ValidPort(port):
        return ((0) < (port)) and ((port) <= (65535))

    @staticmethod
    def ValidAbsoluteHttpPath(path):
        return ((0) < (len(path))) and (((path)[0]) == (_dafny.CodePoint('/')))

    @staticmethod
    def AllNonEmpty(values):
        return ((len(values)) == (0)) or (((0) < (len((values)[0]))) and (default__.AllNonEmpty(_dafny.SeqWithoutIsStrInference((values)[1::]))))

    @staticmethod
    def ValidServiceSpec(spec):
        return (((((((((((((((spec).schemaVersion) == (1)) and ((0) < (len((spec).serviceId)))) and ((0) < (len((spec).sourceClosure)))) and (default__.AllNonEmpty((spec).sourceClosure))) and ((0) < (len((spec).buildGeneration)))) and ((0) < (len((spec).runGeneration)))) and ((0) < (len(((spec).healthIdentity).owner)))) and ((0) < (len(((spec).healthIdentity).protocolVersion)))) and (default__.ValidPort(((spec).endpointPolicy).port))) and (((BindKind_TailscaleIp()) in (((spec).endpointPolicy).allowedBindKinds)) or ((BindKind_LoopbackBind()) in (((spec).endpointPolicy).allowedBindKinds)))) and (default__.ValidAbsoluteHttpPath((spec).consolePath))) and (default__.ValidAbsoluteHttpPath((spec).channelPath))) and ((0) < (len((spec).stateRoot)))) and ((0) < (len((spec).supervisorAdapter)))

    @staticmethod
    def ValidServiceInfo(info):
        return (((((((((((((info).schemaVersion) == (1)) and ((0) < (len((info).serviceId)))) and ((0) < (len(((info).healthIdentity).owner)))) and ((0) < (len(((info).healthIdentity).protocolVersion)))) and ((0) < (len((info).ownerLeaseId)))) and ((0) < ((info).pid))) and ((0) < (len((info).generation)))) and ((0) < (len((info).bindHost)))) and ((0) < (len((info).publicHost)))) and (default__.ValidPort((info).port))) and ((0) < (len((info).consoleUrl)))) and ((0) < (len((info).channelBase)))

    @staticmethod
    def ExactHealth(info, observed):
        return (((((((observed).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ok")))) and (((observed).owner) == (((info).healthIdentity).owner))) and (((observed).protocolVersion) == (((info).healthIdentity).protocolVersion))) and (((observed).pid) == ((info).pid))) and (((observed).generation) == ((info).generation))) and (((observed).ownerLeaseId) == ((info).ownerLeaseId))

    @staticmethod
    def CanUseTailscale(policy, facts):
        return (((((((policy).preferTailscaleMagicDns) and ((BindKind_TailscaleIp()) in ((policy).allowedBindKinds))) and ((facts).backendRunning)) and ((facts).magicDnsEnabled)) and ((facts).selfOnline)) and ((facts).dnsNameValid)) and ((facts).matchingIpValid)

    @staticmethod
    def ChooseEndpoint(policy, facts):
        if default__.CanUseTailscale(policy, facts):
            return EndpointDecision_UseEndpoint(Endpoint_Endpoint((facts).matchingIp, (facts).dnsName, EndpointKind_TailscaleMagicDns(), (policy).port))
        elif ((policy).fallbackLoopback) and ((BindKind_LoopbackBind()) in ((policy).allowedBindKinds)):
            return EndpointDecision_UseEndpoint(Endpoint_Endpoint(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "127.0.0.1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "127.0.0.1")), EndpointKind_Loopback(), (policy).port))
        elif True:
            return EndpointDecision_EndpointUnavailable()

    @staticmethod
    def DecideGeneration(facts):
        if ((facts).candidateBuilt) and ((facts).candidateHealthy):
            return GenerationDecision_PromoteCandidate()
        elif ((facts).activeGenerationExists) and ((facts).activeGenerationHealthy):
            return GenerationDecision_RetainActive()
        elif True:
            return GenerationDecision_InitialGenerationFailed()

    @staticmethod
    def DecideMachineOwner(facts):
        if (((((facts).sharedReceiptExactHealth) and ((facts).machineLeasePresent)) and ((facts).machineLeaseReadable)) and ((facts).machineLeaseMatchesReceipt)) and ((facts).leaseOwnerAlive):
            return MachineOwnerDecision_ReuseExternalOwner()
        elif ((facts).sharedReceiptExactHealth) or (((facts).machineLeasePresent) and ((not((facts).machineLeaseReadable)) or ((facts).leaseOwnerAlive))):
            return MachineOwnerDecision_RefuseSplitBrain()
        elif ((facts).sharedReceiptPresent) or ((facts).machineLeasePresent):
            return MachineOwnerDecision_RecoverStaleOwner()
        elif True:
            return MachineOwnerDecision_StartMachineOwner()

    @staticmethod
    def DecideInstallation(facts):
        if ((facts).installed) and ((facts).launcherMatches):
            return InstallationDecision_KeepInstalled()
        elif True:
            return InstallationDecision_InstallLauncher()

    @staticmethod
    def RequiresBuild(mode, facts):
        return (not((facts).runtimeReady)) or (((mode).is_Production) and ((facts).sourcesStale))

    @staticmethod
    def DecideEnsure(mode, facts):
        if ((((facts).bindingMatches) and ((facts).modeMatches)) and ((facts).healthy)) and (not(default__.RequiresBuild(mode, facts))):
            return EnsureDecision_Reuse()
        elif default__.RequiresBuild(mode, facts):
            return EnsureDecision_ReplaceAndBuild()
        elif True:
            return EnsureDecision_Replace()


class ServiceMode:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [ServiceMode_Production(), ServiceMode_Development()]
    @classmethod
    def default(cls, ):
        return lambda: ServiceMode_Production()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Production(self) -> bool:
        return isinstance(self, ServiceMode_Production)
    @property
    def is_Development(self) -> bool:
        return isinstance(self, ServiceMode_Development)

class ServiceMode_Production(ServiceMode, NamedTuple('Production', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.ServiceMode.Production'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ServiceMode_Production)
    def __hash__(self) -> int:
        return super().__hash__()

class ServiceMode_Development(ServiceMode, NamedTuple('Development', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.ServiceMode.Development'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ServiceMode_Development)
    def __hash__(self) -> int:
        return super().__hash__()


class HealthIdentity:
    @classmethod
    def default(cls, ):
        return lambda: HealthIdentity_HealthIdentity(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_HealthIdentity(self) -> bool:
        return isinstance(self, HealthIdentity_HealthIdentity)

class HealthIdentity_HealthIdentity(HealthIdentity, NamedTuple('HealthIdentity', [('owner', Any), ('protocolVersion', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.HealthIdentity.HealthIdentity({self.owner.VerbatimString(True)}, {self.protocolVersion.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, HealthIdentity_HealthIdentity) and self.owner == __o.owner and self.protocolVersion == __o.protocolVersion
    def __hash__(self) -> int:
        return super().__hash__()


class BindKind:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [BindKind_TailscaleIp(), BindKind_LoopbackBind()]
    @classmethod
    def default(cls, ):
        return lambda: BindKind_TailscaleIp()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_TailscaleIp(self) -> bool:
        return isinstance(self, BindKind_TailscaleIp)
    @property
    def is_LoopbackBind(self) -> bool:
        return isinstance(self, BindKind_LoopbackBind)

class BindKind_TailscaleIp(BindKind, NamedTuple('TailscaleIp', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.BindKind.TailscaleIp'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, BindKind_TailscaleIp)
    def __hash__(self) -> int:
        return super().__hash__()

class BindKind_LoopbackBind(BindKind, NamedTuple('LoopbackBind', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.BindKind.LoopbackBind'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, BindKind_LoopbackBind)
    def __hash__(self) -> int:
        return super().__hash__()


class EndpointPolicy:
    @classmethod
    def default(cls, ):
        return lambda: EndpointPolicy_EndpointPolicy(int(0), False, False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_EndpointPolicy(self) -> bool:
        return isinstance(self, EndpointPolicy_EndpointPolicy)

class EndpointPolicy_EndpointPolicy(EndpointPolicy, NamedTuple('EndpointPolicy', [('port', Any), ('preferTailscaleMagicDns', Any), ('fallbackLoopback', Any), ('allowedBindKinds', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.EndpointPolicy.EndpointPolicy({_dafny.string_of(self.port)}, {_dafny.string_of(self.preferTailscaleMagicDns)}, {_dafny.string_of(self.fallbackLoopback)}, {_dafny.string_of(self.allowedBindKinds)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EndpointPolicy_EndpointPolicy) and self.port == __o.port and self.preferTailscaleMagicDns == __o.preferTailscaleMagicDns and self.fallbackLoopback == __o.fallbackLoopback and self.allowedBindKinds == __o.allowedBindKinds
    def __hash__(self) -> int:
        return super().__hash__()


class ServiceSpec:
    @classmethod
    def default(cls, ):
        return lambda: ServiceSpec_ServiceSpec(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), HealthIdentity.default()(), EndpointPolicy.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ServiceSpec(self) -> bool:
        return isinstance(self, ServiceSpec_ServiceSpec)

class ServiceSpec_ServiceSpec(ServiceSpec, NamedTuple('ServiceSpec', [('schemaVersion', Any), ('serviceId', Any), ('sourceClosure', Any), ('buildGeneration', Any), ('runGeneration', Any), ('healthIdentity', Any), ('endpointPolicy', Any), ('consolePath', Any), ('channelPath', Any), ('stateRoot', Any), ('supervisorAdapter', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.ServiceSpec.ServiceSpec({_dafny.string_of(self.schemaVersion)}, {self.serviceId.VerbatimString(True)}, {_dafny.string_of(self.sourceClosure)}, {self.buildGeneration.VerbatimString(True)}, {self.runGeneration.VerbatimString(True)}, {_dafny.string_of(self.healthIdentity)}, {_dafny.string_of(self.endpointPolicy)}, {self.consolePath.VerbatimString(True)}, {self.channelPath.VerbatimString(True)}, {self.stateRoot.VerbatimString(True)}, {self.supervisorAdapter.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ServiceSpec_ServiceSpec) and self.schemaVersion == __o.schemaVersion and self.serviceId == __o.serviceId and self.sourceClosure == __o.sourceClosure and self.buildGeneration == __o.buildGeneration and self.runGeneration == __o.runGeneration and self.healthIdentity == __o.healthIdentity and self.endpointPolicy == __o.endpointPolicy and self.consolePath == __o.consolePath and self.channelPath == __o.channelPath and self.stateRoot == __o.stateRoot and self.supervisorAdapter == __o.supervisorAdapter
    def __hash__(self) -> int:
        return super().__hash__()


class EndpointKind:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [EndpointKind_TailscaleMagicDns(), EndpointKind_Loopback()]
    @classmethod
    def default(cls, ):
        return lambda: EndpointKind_TailscaleMagicDns()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_TailscaleMagicDns(self) -> bool:
        return isinstance(self, EndpointKind_TailscaleMagicDns)
    @property
    def is_Loopback(self) -> bool:
        return isinstance(self, EndpointKind_Loopback)

class EndpointKind_TailscaleMagicDns(EndpointKind, NamedTuple('TailscaleMagicDns', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.EndpointKind.TailscaleMagicDns'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EndpointKind_TailscaleMagicDns)
    def __hash__(self) -> int:
        return super().__hash__()

class EndpointKind_Loopback(EndpointKind, NamedTuple('Loopback', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.EndpointKind.Loopback'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EndpointKind_Loopback)
    def __hash__(self) -> int:
        return super().__hash__()


class Endpoint:
    @classmethod
    def default(cls, ):
        return lambda: Endpoint_Endpoint(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), EndpointKind.default()(), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Endpoint(self) -> bool:
        return isinstance(self, Endpoint_Endpoint)

class Endpoint_Endpoint(Endpoint, NamedTuple('Endpoint', [('bindHost', Any), ('publicHost', Any), ('endpointKind', Any), ('port', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.Endpoint.Endpoint({self.bindHost.VerbatimString(True)}, {self.publicHost.VerbatimString(True)}, {_dafny.string_of(self.endpointKind)}, {_dafny.string_of(self.port)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Endpoint_Endpoint) and self.bindHost == __o.bindHost and self.publicHost == __o.publicHost and self.endpointKind == __o.endpointKind and self.port == __o.port
    def __hash__(self) -> int:
        return super().__hash__()


class TailscaleFacts:
    @classmethod
    def default(cls, ):
        return lambda: TailscaleFacts_TailscaleFacts(False, False, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_TailscaleFacts(self) -> bool:
        return isinstance(self, TailscaleFacts_TailscaleFacts)

class TailscaleFacts_TailscaleFacts(TailscaleFacts, NamedTuple('TailscaleFacts', [('backendRunning', Any), ('magicDnsEnabled', Any), ('selfOnline', Any), ('dnsName', Any), ('dnsNameValid', Any), ('matchingIp', Any), ('matchingIpValid', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.TailscaleFacts.TailscaleFacts({_dafny.string_of(self.backendRunning)}, {_dafny.string_of(self.magicDnsEnabled)}, {_dafny.string_of(self.selfOnline)}, {self.dnsName.VerbatimString(True)}, {_dafny.string_of(self.dnsNameValid)}, {self.matchingIp.VerbatimString(True)}, {_dafny.string_of(self.matchingIpValid)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TailscaleFacts_TailscaleFacts) and self.backendRunning == __o.backendRunning and self.magicDnsEnabled == __o.magicDnsEnabled and self.selfOnline == __o.selfOnline and self.dnsName == __o.dnsName and self.dnsNameValid == __o.dnsNameValid and self.matchingIp == __o.matchingIp and self.matchingIpValid == __o.matchingIpValid
    def __hash__(self) -> int:
        return super().__hash__()


class EndpointDecision:
    @classmethod
    def default(cls, ):
        return lambda: EndpointDecision_UseEndpoint(Endpoint.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_UseEndpoint(self) -> bool:
        return isinstance(self, EndpointDecision_UseEndpoint)
    @property
    def is_EndpointUnavailable(self) -> bool:
        return isinstance(self, EndpointDecision_EndpointUnavailable)

class EndpointDecision_UseEndpoint(EndpointDecision, NamedTuple('UseEndpoint', [('endpoint', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.EndpointDecision.UseEndpoint({_dafny.string_of(self.endpoint)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EndpointDecision_UseEndpoint) and self.endpoint == __o.endpoint
    def __hash__(self) -> int:
        return super().__hash__()

class EndpointDecision_EndpointUnavailable(EndpointDecision, NamedTuple('EndpointUnavailable', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.EndpointDecision.EndpointUnavailable'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EndpointDecision_EndpointUnavailable)
    def __hash__(self) -> int:
        return super().__hash__()


class RefreshStatus:
    @classmethod
    def default(cls, ):
        return lambda: RefreshStatus_RefreshStatus(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RefreshStatus(self) -> bool:
        return isinstance(self, RefreshStatus_RefreshStatus)

class RefreshStatus_RefreshStatus(RefreshStatus, NamedTuple('RefreshStatus', [('status', Any), ('attemptedGeneration', Any), ('diagnostic', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.RefreshStatus.RefreshStatus({self.status.VerbatimString(True)}, {self.attemptedGeneration.VerbatimString(True)}, {self.diagnostic.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RefreshStatus_RefreshStatus) and self.status == __o.status and self.attemptedGeneration == __o.attemptedGeneration and self.diagnostic == __o.diagnostic
    def __hash__(self) -> int:
        return super().__hash__()


class ServiceInfo:
    @classmethod
    def default(cls, ):
        return lambda: ServiceInfo_ServiceInfo(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), HealthIdentity.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), EndpointKind.default()(), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), RefreshStatus.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ServiceInfo(self) -> bool:
        return isinstance(self, ServiceInfo_ServiceInfo)

class ServiceInfo_ServiceInfo(ServiceInfo, NamedTuple('ServiceInfo', [('schemaVersion', Any), ('serviceId', Any), ('healthIdentity', Any), ('disposition', Any), ('status', Any), ('ownerLeaseId', Any), ('pid', Any), ('generation', Any), ('bindHost', Any), ('publicHost', Any), ('endpointKind', Any), ('port', Any), ('consoleUrl', Any), ('channelBase', Any), ('refresh', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.ServiceInfo.ServiceInfo({_dafny.string_of(self.schemaVersion)}, {self.serviceId.VerbatimString(True)}, {_dafny.string_of(self.healthIdentity)}, {self.disposition.VerbatimString(True)}, {self.status.VerbatimString(True)}, {self.ownerLeaseId.VerbatimString(True)}, {_dafny.string_of(self.pid)}, {self.generation.VerbatimString(True)}, {self.bindHost.VerbatimString(True)}, {self.publicHost.VerbatimString(True)}, {_dafny.string_of(self.endpointKind)}, {_dafny.string_of(self.port)}, {self.consoleUrl.VerbatimString(True)}, {self.channelBase.VerbatimString(True)}, {_dafny.string_of(self.refresh)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ServiceInfo_ServiceInfo) and self.schemaVersion == __o.schemaVersion and self.serviceId == __o.serviceId and self.healthIdentity == __o.healthIdentity and self.disposition == __o.disposition and self.status == __o.status and self.ownerLeaseId == __o.ownerLeaseId and self.pid == __o.pid and self.generation == __o.generation and self.bindHost == __o.bindHost and self.publicHost == __o.publicHost and self.endpointKind == __o.endpointKind and self.port == __o.port and self.consoleUrl == __o.consoleUrl and self.channelBase == __o.channelBase and self.refresh == __o.refresh
    def __hash__(self) -> int:
        return super().__hash__()


class HealthObservation:
    @classmethod
    def default(cls, ):
        return lambda: HealthObservation_HealthObservation(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_HealthObservation(self) -> bool:
        return isinstance(self, HealthObservation_HealthObservation)

class HealthObservation_HealthObservation(HealthObservation, NamedTuple('HealthObservation', [('status', Any), ('owner', Any), ('protocolVersion', Any), ('pid', Any), ('generation', Any), ('ownerLeaseId', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.HealthObservation.HealthObservation({self.status.VerbatimString(True)}, {self.owner.VerbatimString(True)}, {self.protocolVersion.VerbatimString(True)}, {_dafny.string_of(self.pid)}, {self.generation.VerbatimString(True)}, {self.ownerLeaseId.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, HealthObservation_HealthObservation) and self.status == __o.status and self.owner == __o.owner and self.protocolVersion == __o.protocolVersion and self.pid == __o.pid and self.generation == __o.generation and self.ownerLeaseId == __o.ownerLeaseId
    def __hash__(self) -> int:
        return super().__hash__()


class GenerationFacts:
    @classmethod
    def default(cls, ):
        return lambda: GenerationFacts_GenerationFacts(False, False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_GenerationFacts(self) -> bool:
        return isinstance(self, GenerationFacts_GenerationFacts)

class GenerationFacts_GenerationFacts(GenerationFacts, NamedTuple('GenerationFacts', [('activeGenerationExists', Any), ('activeGenerationHealthy', Any), ('candidateBuilt', Any), ('candidateHealthy', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.GenerationFacts.GenerationFacts({_dafny.string_of(self.activeGenerationExists)}, {_dafny.string_of(self.activeGenerationHealthy)}, {_dafny.string_of(self.candidateBuilt)}, {_dafny.string_of(self.candidateHealthy)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, GenerationFacts_GenerationFacts) and self.activeGenerationExists == __o.activeGenerationExists and self.activeGenerationHealthy == __o.activeGenerationHealthy and self.candidateBuilt == __o.candidateBuilt and self.candidateHealthy == __o.candidateHealthy
    def __hash__(self) -> int:
        return super().__hash__()


class GenerationDecision:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [GenerationDecision_PromoteCandidate(), GenerationDecision_RetainActive(), GenerationDecision_InitialGenerationFailed()]
    @classmethod
    def default(cls, ):
        return lambda: GenerationDecision_PromoteCandidate()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_PromoteCandidate(self) -> bool:
        return isinstance(self, GenerationDecision_PromoteCandidate)
    @property
    def is_RetainActive(self) -> bool:
        return isinstance(self, GenerationDecision_RetainActive)
    @property
    def is_InitialGenerationFailed(self) -> bool:
        return isinstance(self, GenerationDecision_InitialGenerationFailed)

class GenerationDecision_PromoteCandidate(GenerationDecision, NamedTuple('PromoteCandidate', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.GenerationDecision.PromoteCandidate'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, GenerationDecision_PromoteCandidate)
    def __hash__(self) -> int:
        return super().__hash__()

class GenerationDecision_RetainActive(GenerationDecision, NamedTuple('RetainActive', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.GenerationDecision.RetainActive'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, GenerationDecision_RetainActive)
    def __hash__(self) -> int:
        return super().__hash__()

class GenerationDecision_InitialGenerationFailed(GenerationDecision, NamedTuple('InitialGenerationFailed', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.GenerationDecision.InitialGenerationFailed'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, GenerationDecision_InitialGenerationFailed)
    def __hash__(self) -> int:
        return super().__hash__()


class MachineOwnerFacts:
    @classmethod
    def default(cls, ):
        return lambda: MachineOwnerFacts_MachineOwnerFacts(False, False, False, False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_MachineOwnerFacts(self) -> bool:
        return isinstance(self, MachineOwnerFacts_MachineOwnerFacts)

class MachineOwnerFacts_MachineOwnerFacts(MachineOwnerFacts, NamedTuple('MachineOwnerFacts', [('sharedReceiptPresent', Any), ('sharedReceiptExactHealth', Any), ('machineLeasePresent', Any), ('machineLeaseReadable', Any), ('machineLeaseMatchesReceipt', Any), ('leaseOwnerAlive', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.MachineOwnerFacts.MachineOwnerFacts({_dafny.string_of(self.sharedReceiptPresent)}, {_dafny.string_of(self.sharedReceiptExactHealth)}, {_dafny.string_of(self.machineLeasePresent)}, {_dafny.string_of(self.machineLeaseReadable)}, {_dafny.string_of(self.machineLeaseMatchesReceipt)}, {_dafny.string_of(self.leaseOwnerAlive)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, MachineOwnerFacts_MachineOwnerFacts) and self.sharedReceiptPresent == __o.sharedReceiptPresent and self.sharedReceiptExactHealth == __o.sharedReceiptExactHealth and self.machineLeasePresent == __o.machineLeasePresent and self.machineLeaseReadable == __o.machineLeaseReadable and self.machineLeaseMatchesReceipt == __o.machineLeaseMatchesReceipt and self.leaseOwnerAlive == __o.leaseOwnerAlive
    def __hash__(self) -> int:
        return super().__hash__()


class MachineOwnerDecision:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [MachineOwnerDecision_ReuseExternalOwner(), MachineOwnerDecision_StartMachineOwner(), MachineOwnerDecision_RecoverStaleOwner(), MachineOwnerDecision_RefuseSplitBrain()]
    @classmethod
    def default(cls, ):
        return lambda: MachineOwnerDecision_ReuseExternalOwner()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ReuseExternalOwner(self) -> bool:
        return isinstance(self, MachineOwnerDecision_ReuseExternalOwner)
    @property
    def is_StartMachineOwner(self) -> bool:
        return isinstance(self, MachineOwnerDecision_StartMachineOwner)
    @property
    def is_RecoverStaleOwner(self) -> bool:
        return isinstance(self, MachineOwnerDecision_RecoverStaleOwner)
    @property
    def is_RefuseSplitBrain(self) -> bool:
        return isinstance(self, MachineOwnerDecision_RefuseSplitBrain)

class MachineOwnerDecision_ReuseExternalOwner(MachineOwnerDecision, NamedTuple('ReuseExternalOwner', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.MachineOwnerDecision.ReuseExternalOwner'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, MachineOwnerDecision_ReuseExternalOwner)
    def __hash__(self) -> int:
        return super().__hash__()

class MachineOwnerDecision_StartMachineOwner(MachineOwnerDecision, NamedTuple('StartMachineOwner', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.MachineOwnerDecision.StartMachineOwner'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, MachineOwnerDecision_StartMachineOwner)
    def __hash__(self) -> int:
        return super().__hash__()

class MachineOwnerDecision_RecoverStaleOwner(MachineOwnerDecision, NamedTuple('RecoverStaleOwner', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.MachineOwnerDecision.RecoverStaleOwner'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, MachineOwnerDecision_RecoverStaleOwner)
    def __hash__(self) -> int:
        return super().__hash__()

class MachineOwnerDecision_RefuseSplitBrain(MachineOwnerDecision, NamedTuple('RefuseSplitBrain', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.MachineOwnerDecision.RefuseSplitBrain'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, MachineOwnerDecision_RefuseSplitBrain)
    def __hash__(self) -> int:
        return super().__hash__()


class InstallationFacts:
    @classmethod
    def default(cls, ):
        return lambda: InstallationFacts_InstallationFacts(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_InstallationFacts(self) -> bool:
        return isinstance(self, InstallationFacts_InstallationFacts)

class InstallationFacts_InstallationFacts(InstallationFacts, NamedTuple('InstallationFacts', [('installed', Any), ('launcherMatches', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.InstallationFacts.InstallationFacts({_dafny.string_of(self.installed)}, {_dafny.string_of(self.launcherMatches)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, InstallationFacts_InstallationFacts) and self.installed == __o.installed and self.launcherMatches == __o.launcherMatches
    def __hash__(self) -> int:
        return super().__hash__()


class InstallationDecision:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [InstallationDecision_KeepInstalled(), InstallationDecision_InstallLauncher()]
    @classmethod
    def default(cls, ):
        return lambda: InstallationDecision_KeepInstalled()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_KeepInstalled(self) -> bool:
        return isinstance(self, InstallationDecision_KeepInstalled)
    @property
    def is_InstallLauncher(self) -> bool:
        return isinstance(self, InstallationDecision_InstallLauncher)

class InstallationDecision_KeepInstalled(InstallationDecision, NamedTuple('KeepInstalled', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.InstallationDecision.KeepInstalled'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, InstallationDecision_KeepInstalled)
    def __hash__(self) -> int:
        return super().__hash__()

class InstallationDecision_InstallLauncher(InstallationDecision, NamedTuple('InstallLauncher', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.InstallationDecision.InstallLauncher'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, InstallationDecision_InstallLauncher)
    def __hash__(self) -> int:
        return super().__hash__()


class EnsureFacts:
    @classmethod
    def default(cls, ):
        return lambda: EnsureFacts_EnsureFacts(False, False, False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_EnsureFacts(self) -> bool:
        return isinstance(self, EnsureFacts_EnsureFacts)

class EnsureFacts_EnsureFacts(EnsureFacts, NamedTuple('EnsureFacts', [('runtimeReady', Any), ('sourcesStale', Any), ('bindingMatches', Any), ('modeMatches', Any), ('healthy', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.EnsureFacts.EnsureFacts({_dafny.string_of(self.runtimeReady)}, {_dafny.string_of(self.sourcesStale)}, {_dafny.string_of(self.bindingMatches)}, {_dafny.string_of(self.modeMatches)}, {_dafny.string_of(self.healthy)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EnsureFacts_EnsureFacts) and self.runtimeReady == __o.runtimeReady and self.sourcesStale == __o.sourcesStale and self.bindingMatches == __o.bindingMatches and self.modeMatches == __o.modeMatches and self.healthy == __o.healthy
    def __hash__(self) -> int:
        return super().__hash__()


class EnsureDecision:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [EnsureDecision_Reuse(), EnsureDecision_Replace(), EnsureDecision_ReplaceAndBuild()]
    @classmethod
    def default(cls, ):
        return lambda: EnsureDecision_Reuse()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Reuse(self) -> bool:
        return isinstance(self, EnsureDecision_Reuse)
    @property
    def is_Replace(self) -> bool:
        return isinstance(self, EnsureDecision_Replace)
    @property
    def is_ReplaceAndBuild(self) -> bool:
        return isinstance(self, EnsureDecision_ReplaceAndBuild)

class EnsureDecision_Reuse(EnsureDecision, NamedTuple('Reuse', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.EnsureDecision.Reuse'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EnsureDecision_Reuse)
    def __hash__(self) -> int:
        return super().__hash__()

class EnsureDecision_Replace(EnsureDecision, NamedTuple('Replace', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.EnsureDecision.Replace'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EnsureDecision_Replace)
    def __hash__(self) -> int:
        return super().__hash__()

class EnsureDecision_ReplaceAndBuild(EnsureDecision, NamedTuple('ReplaceAndBuild', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxServiceLifecycle.EnsureDecision.ReplaceAndBuild'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EnsureDecision_ReplaceAndBuild)
    def __hash__(self) -> int:
        return super().__hash__()

