import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_ as module_
import _dafny as _dafny
import System_ as System_
import ConfluxCodec as ConfluxCodec
import ConfluxContract as ConfluxContract

# Module: AhpSkeleton

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Pow10(k):
        d_0___accumulator_ = 1
        while True:
            with _dafny.label():
                if (k) == (0):
                    return (1) * (d_0___accumulator_)
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) * (10)
                    in0_ = (k) - (1)
                    k = in0_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def decInt(m, e):
        if (e) >= (0):
            return Option_Some((m) * (default__.Pow10(e)))
        elif (_dafny.euclidian_modulus(m, default__.Pow10((0) - (e)))) == (0):
            return Option_Some(_dafny.euclidian_division(m, default__.Pow10((0) - (e))))
        elif True:
            return Option_None()

    @staticmethod
    def jsonInt(j):
        if (j).is_JNum:
            return Option_Some((j).n)
        elif (j).is_JDec:
            return default__.decInt((j).mantissa, (j).exp)
        elif True:
            return Option_None()

    @staticmethod
    def asInt(o):
        if (o).is_None:
            return 0
        elif True:
            d_0_r_ = default__.jsonInt((o).value)
            if (d_0_r_).is_Some:
                return (d_0_r_).value
            elif True:
                return 0

    @staticmethod
    def optInt(o):
        if (o).is_Some:
            return default__.jsonInt((o).value)
        elif True:
            return Option_None()

    @staticmethod
    def EncodeStatus(s):
        return s

    @staticmethod
    def ApplyToRoot(s, a, now):
        source0_ = a
        if True:
            if source0_.is_RootAgentsChanged:
                d_0_ag_ = source0_.agents
                def iife0_(_pat_let0_0):
                    def iife1_(d_1_dt__update__tmp_h0_):
                        def iife2_(_pat_let1_0):
                            def iife3_(d_2_dt__update_hagents_h0_):
                                return RootState_RootState(d_2_dt__update_hagents_h0_, (d_1_dt__update__tmp_h0_).activeSessions, (d_1_dt__update__tmp_h0_).terminals, (d_1_dt__update__tmp_h0_).config)
                            return iife3_(_pat_let1_0)
                        return iife2_(d_0_ag_)
                    return iife1_(_pat_let0_0)
                return (iife0_(s), ReduceOutcome_Applied())
        if True:
            if source0_.is_RootActiveSessionsChanged:
                d_3_n_ = source0_.activeSessions
                def iife4_(_pat_let2_0):
                    def iife5_(d_4_dt__update__tmp_h1_):
                        def iife6_(_pat_let3_0):
                            def iife7_(d_5_dt__update_hactiveSessions_h0_):
                                return RootState_RootState((d_4_dt__update__tmp_h1_).agents, d_5_dt__update_hactiveSessions_h0_, (d_4_dt__update__tmp_h1_).terminals, (d_4_dt__update__tmp_h1_).config)
                            return iife7_(_pat_let3_0)
                        return iife6_(Option_Some(d_3_n_))
                    return iife5_(_pat_let2_0)
                return (iife4_(s), ReduceOutcome_Applied())
        if True:
            if source0_.is_RootTerminalsChanged:
                d_6_t_ = source0_.terminals
                def iife8_(_pat_let4_0):
                    def iife9_(d_7_dt__update__tmp_h2_):
                        def iife10_(_pat_let5_0):
                            def iife11_(d_8_dt__update_hterminals_h0_):
                                return RootState_RootState((d_7_dt__update__tmp_h2_).agents, (d_7_dt__update__tmp_h2_).activeSessions, d_8_dt__update_hterminals_h0_, (d_7_dt__update__tmp_h2_).config)
                            return iife11_(_pat_let5_0)
                        return iife10_(Option_Some(d_6_t_))
                    return iife9_(_pat_let4_0)
                return (iife8_(s), ReduceOutcome_Applied())
        if True:
            if source0_.is_RootConfigChanged:
                d_9_cfg_ = source0_.config
                d_10_repl_ = source0_.replace
                source1_ = (s).config
                if True:
                    if source1_.is_None:
                        return (s, ReduceOutcome_NoOp())
                if True:
                    d_11_c_ = source1_.value
                    d_12_newValues_ = (d_9_cfg_ if d_10_repl_ else ((d_11_c_).values) | (d_9_cfg_))
                    def iife12_(_pat_let6_0):
                        def iife13_(d_13_dt__update__tmp_h3_):
                            def iife14_(_pat_let7_0):
                                def iife15_(d_14_dt__update_hconfig_h0_):
                                    return RootState_RootState((d_13_dt__update__tmp_h3_).agents, (d_13_dt__update__tmp_h3_).activeSessions, (d_13_dt__update__tmp_h3_).terminals, d_14_dt__update_hconfig_h0_)
                                return iife15_(_pat_let7_0)
                            return iife14_(Option_Some(RootConfig_RootConfig((d_11_c_).schema, d_12_newValues_)))
                        return iife13_(_pat_let6_0)
                    return (iife12_(s), ReduceOutcome_Applied())
        if True:
            return (s, ReduceOutcome_NoOp())

    @staticmethod
    def IsUnknownRoot(a):
        return (a).is_RootUnknown

    @staticmethod
    def ConfigPreconditionFails(s, a):
        return ((a).is_RootConfigChanged) and (((s).config).is_None)

    @staticmethod
    def FoldRoot(s, actions, now):
        def lambda0_(d_0_now_):
            def lambda1_(d_1_s_k_, d_2_a_):
                return (default__.ApplyToRoot(d_1_s_k_, d_2_a_, d_0_now_))[0]

            return lambda1_

        return ConfluxContract.default__.Fold(lambda0_(now), s, actions)

    @staticmethod
    def apply1(s, a):
        return (default__.ApplyToRoot(s, a, 9999))[0]

    @staticmethod
    def RunCorpus():
        pass_: int = int(0)
        total: int = int(0)
        total = 7
        pass_ = 0
        d_0_empty_: RootState
        d_0_empty_ = RootState_RootState(_dafny.SeqWithoutIsStrInference([]), Option_None(), Option_None(), Option_None())
        d_1_ag_: _dafny.Seq
        d_1_ag_ = _dafny.SeqWithoutIsStrInference([AgentInfo_AgentInfo(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "copilot")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Copilot")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "AI")), _dafny.SeqWithoutIsStrInference([]))])
        if (default__.apply1(d_0_empty_, RootAction_RootAgentsChanged(d_1_ag_))) == (RootState_RootState(d_1_ag_, Option_None(), Option_None(), Option_None())):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_0_empty_, RootAction_RootActiveSessionsChanged(5))) == (RootState_RootState(_dafny.SeqWithoutIsStrInference([]), Option_Some(5), Option_None(), Option_None())):
            pass_ = (pass_) + (1)
        d_2_terms_: _dafny.Seq
        d_2_terms_ = _dafny.SeqWithoutIsStrInference([ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "resource")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agenthost:/terminal/1")))})), ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "resource")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agenthost:/terminal/2")))}))])
        if (default__.apply1(d_0_empty_, RootAction_RootTerminalsChanged(d_2_terms_))) == (RootState_RootState(_dafny.SeqWithoutIsStrInference([]), Option_None(), Option_Some(d_2_terms_), Option_None())):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_0_empty_, RootAction_RootUnknown(ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "type")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "root/nonExistentAction")))}))))) == (d_0_empty_):
            pass_ = (pass_) + (1)
        d_3_sch_: ConfluxCodec.Json
        d_3_sch_ = ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "type")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "object")))}))
        d_4_cfg127_: RootState
        d_4_cfg127_ = RootState_RootState(_dafny.SeqWithoutIsStrInference([]), Option_None(), Option_None(), Option_Some(RootConfig_RootConfig(d_3_sch_, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "theme")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "light")))}))))
        d_5_exp127_: RootState
        d_5_exp127_ = RootState_RootState(_dafny.SeqWithoutIsStrInference([]), Option_None(), Option_None(), Option_Some(RootConfig_RootConfig(d_3_sch_, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "theme")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dark")))}))))
        if (default__.apply1(d_4_cfg127_, RootAction_RootConfigChanged(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "theme")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dark")))}), False))) == (d_5_exp127_):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_0_empty_, RootAction_RootConfigChanged(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "theme")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dark")))}), False))) == (d_0_empty_):
            pass_ = (pass_) + (1)
        d_6_cfg130_: RootState
        d_6_cfg130_ = RootState_RootState(_dafny.SeqWithoutIsStrInference([]), Option_None(), Option_None(), Option_Some(RootConfig_RootConfig(d_3_sch_, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "theme")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "light"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "locale")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "en")))}))))
        d_7_exp130_: RootState
        d_7_exp130_ = RootState_RootState(_dafny.SeqWithoutIsStrInference([]), Option_None(), Option_None(), Option_Some(RootConfig_RootConfig(d_3_sch_, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "theme")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dark")))}))))
        if (default__.apply1(d_6_cfg130_, RootAction_RootConfigChanged(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "theme")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dark")))}), True))) == (d_7_exp130_):
            pass_ = (pass_) + (1)
        return pass_, total


class Option:
    @classmethod
    def default(cls, ):
        return lambda: Option_None()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_None(self) -> bool:
        return isinstance(self, Option_None)
    @property
    def is_Some(self) -> bool:
        return isinstance(self, Option_Some)

class Option_None(Option, NamedTuple('None_', [])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.Option.None'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Option_None)
    def __hash__(self) -> int:
        return super().__hash__()

class Option_Some(Option, NamedTuple('Some', [('value', Any)])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.Option.Some({_dafny.string_of(self.value)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Option_Some) and self.value == __o.value
    def __hash__(self) -> int:
        return super().__hash__()


class ReduceOutcome:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [ReduceOutcome_Applied(), ReduceOutcome_NoOp(), ReduceOutcome_OutOfScope()]
    @classmethod
    def default(cls, ):
        return lambda: ReduceOutcome_Applied()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Applied(self) -> bool:
        return isinstance(self, ReduceOutcome_Applied)
    @property
    def is_NoOp(self) -> bool:
        return isinstance(self, ReduceOutcome_NoOp)
    @property
    def is_OutOfScope(self) -> bool:
        return isinstance(self, ReduceOutcome_OutOfScope)

class ReduceOutcome_Applied(ReduceOutcome, NamedTuple('Applied', [])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.ReduceOutcome.Applied'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ReduceOutcome_Applied)
    def __hash__(self) -> int:
        return super().__hash__()

class ReduceOutcome_NoOp(ReduceOutcome, NamedTuple('NoOp', [])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.ReduceOutcome.NoOp'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ReduceOutcome_NoOp)
    def __hash__(self) -> int:
        return super().__hash__()

class ReduceOutcome_OutOfScope(ReduceOutcome, NamedTuple('OutOfScope', [])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.ReduceOutcome.OutOfScope'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ReduceOutcome_OutOfScope)
    def __hash__(self) -> int:
        return super().__hash__()


class ActionOrigin:
    @classmethod
    def default(cls, ):
        return lambda: ActionOrigin_ActionOrigin(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ActionOrigin(self) -> bool:
        return isinstance(self, ActionOrigin_ActionOrigin)

class ActionOrigin_ActionOrigin(ActionOrigin, NamedTuple('ActionOrigin', [('clientId', Any), ('clientSeq', Any)])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.ActionOrigin.ActionOrigin({self.clientId.VerbatimString(True)}, {_dafny.string_of(self.clientSeq)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ActionOrigin_ActionOrigin) and self.clientId == __o.clientId and self.clientSeq == __o.clientSeq
    def __hash__(self) -> int:
        return super().__hash__()


class Envelope:
    @classmethod
    def default(cls, default_A):
        return lambda: Envelope_Envelope(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), default_A(), int(0), Option.default()(), Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Envelope(self) -> bool:
        return isinstance(self, Envelope_Envelope)

class Envelope_Envelope(Envelope, NamedTuple('Envelope', [('channel', Any), ('action', Any), ('serverSeq', Any), ('origin', Any), ('rejectionReason', Any)])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.Envelope.Envelope({self.channel.VerbatimString(True)}, {_dafny.string_of(self.action)}, {_dafny.string_of(self.serverSeq)}, {_dafny.string_of(self.origin)}, {_dafny.string_of(self.rejectionReason)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Envelope_Envelope) and self.channel == __o.channel and self.action == __o.action and self.serverSeq == __o.serverSeq and self.origin == __o.origin and self.rejectionReason == __o.rejectionReason
    def __hash__(self) -> int:
        return super().__hash__()


class AgentInfo:
    @classmethod
    def default(cls, ):
        return lambda: AgentInfo_AgentInfo(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_AgentInfo(self) -> bool:
        return isinstance(self, AgentInfo_AgentInfo)

class AgentInfo_AgentInfo(AgentInfo, NamedTuple('AgentInfo', [('provider', Any), ('displayName', Any), ('description', Any), ('models', Any)])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.AgentInfo.AgentInfo({self.provider.VerbatimString(True)}, {self.displayName.VerbatimString(True)}, {self.description.VerbatimString(True)}, {_dafny.string_of(self.models)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AgentInfo_AgentInfo) and self.provider == __o.provider and self.displayName == __o.displayName and self.description == __o.description and self.models == __o.models
    def __hash__(self) -> int:
        return super().__hash__()


class RootConfig:
    @classmethod
    def default(cls, ):
        return lambda: RootConfig_RootConfig(ConfluxCodec.Json.default()(), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RootConfig(self) -> bool:
        return isinstance(self, RootConfig_RootConfig)

class RootConfig_RootConfig(RootConfig, NamedTuple('RootConfig', [('schema', Any), ('values', Any)])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.RootConfig.RootConfig({_dafny.string_of(self.schema)}, {_dafny.string_of(self.values)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RootConfig_RootConfig) and self.schema == __o.schema and self.values == __o.values
    def __hash__(self) -> int:
        return super().__hash__()


class RootState:
    @classmethod
    def default(cls, ):
        return lambda: RootState_RootState(_dafny.Seq({}), Option.default()(), Option.default()(), Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RootState(self) -> bool:
        return isinstance(self, RootState_RootState)

class RootState_RootState(RootState, NamedTuple('RootState', [('agents', Any), ('activeSessions', Any), ('terminals', Any), ('config', Any)])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.RootState.RootState({_dafny.string_of(self.agents)}, {_dafny.string_of(self.activeSessions)}, {_dafny.string_of(self.terminals)}, {_dafny.string_of(self.config)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RootState_RootState) and self.agents == __o.agents and self.activeSessions == __o.activeSessions and self.terminals == __o.terminals and self.config == __o.config
    def __hash__(self) -> int:
        return super().__hash__()


class RootAction:
    @classmethod
    def default(cls, ):
        return lambda: RootAction_RootAgentsChanged(_dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RootAgentsChanged(self) -> bool:
        return isinstance(self, RootAction_RootAgentsChanged)
    @property
    def is_RootActiveSessionsChanged(self) -> bool:
        return isinstance(self, RootAction_RootActiveSessionsChanged)
    @property
    def is_RootTerminalsChanged(self) -> bool:
        return isinstance(self, RootAction_RootTerminalsChanged)
    @property
    def is_RootConfigChanged(self) -> bool:
        return isinstance(self, RootAction_RootConfigChanged)
    @property
    def is_RootUnknown(self) -> bool:
        return isinstance(self, RootAction_RootUnknown)

class RootAction_RootAgentsChanged(RootAction, NamedTuple('RootAgentsChanged', [('agents', Any)])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.RootAction.RootAgentsChanged({_dafny.string_of(self.agents)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RootAction_RootAgentsChanged) and self.agents == __o.agents
    def __hash__(self) -> int:
        return super().__hash__()

class RootAction_RootActiveSessionsChanged(RootAction, NamedTuple('RootActiveSessionsChanged', [('activeSessions', Any)])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.RootAction.RootActiveSessionsChanged({_dafny.string_of(self.activeSessions)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RootAction_RootActiveSessionsChanged) and self.activeSessions == __o.activeSessions
    def __hash__(self) -> int:
        return super().__hash__()

class RootAction_RootTerminalsChanged(RootAction, NamedTuple('RootTerminalsChanged', [('terminals', Any)])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.RootAction.RootTerminalsChanged({_dafny.string_of(self.terminals)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RootAction_RootTerminalsChanged) and self.terminals == __o.terminals
    def __hash__(self) -> int:
        return super().__hash__()

class RootAction_RootConfigChanged(RootAction, NamedTuple('RootConfigChanged', [('config', Any), ('replace', Any)])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.RootAction.RootConfigChanged({_dafny.string_of(self.config)}, {_dafny.string_of(self.replace)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RootAction_RootConfigChanged) and self.config == __o.config and self.replace == __o.replace
    def __hash__(self) -> int:
        return super().__hash__()

class RootAction_RootUnknown(RootAction, NamedTuple('RootUnknown', [('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'AhpSkeleton.RootAction.RootUnknown({_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RootAction_RootUnknown) and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()

