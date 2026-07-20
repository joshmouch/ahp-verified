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

# Module: ConfluxStateMachine

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def TransitionState(step):
        def lambda0_(d_0_step_):
            def lambda1_(d_1_s_, d_2_a_):
                return (d_0_step_(d_1_s_, d_2_a_))[0]

            return lambda1_

        return lambda0_(step)

    @staticmethod
    def FinalState(run):
        return (run).state

    @staticmethod
    def Outputs(run):
        return (run).outputs

    @staticmethod
    def RunTransitions(step, s, xs):
        if (len(xs)) == (0):
            return TransitionRun_TransitionRun(s, _dafny.SeqWithoutIsStrInference([]))
        elif True:
            d_0_so_ = step(s, (xs)[0])
            d_1_rest_ = default__.RunTransitions(step, (d_0_so_)[0], _dafny.SeqWithoutIsStrInference((xs)[1::]))
            return TransitionRun_TransitionRun((d_1_rest_).state, (_dafny.SeqWithoutIsStrInference([(d_0_so_)[1]])) + ((d_1_rest_).outputs))

    @staticmethod
    def GuardedTransition(base, breached, onBreach):
        def lambda0_(d_0_breached_, d_1_onBreach_, d_2_base_):
            def lambda1_(d_3_s_, d_4_a_):
                return ((d_3_s_, d_1_onBreach_(d_3_s_, d_4_a_)) if d_0_breached_(d_3_s_, d_4_a_) else d_2_base_(d_3_s_, d_4_a_))

            return lambda1_

        return lambda0_(breached, onBreach, base)


class TransitionRun:
    @classmethod
    def default(cls, default_S):
        return lambda: TransitionRun_TransitionRun(default_S(), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_TransitionRun(self) -> bool:
        return isinstance(self, TransitionRun_TransitionRun)

class TransitionRun_TransitionRun(TransitionRun, NamedTuple('TransitionRun', [('state', Any), ('outputs', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxStateMachine.TransitionRun.TransitionRun({_dafny.string_of(self.state)}, {_dafny.string_of(self.outputs)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TransitionRun_TransitionRun) and self.state == __o.state and self.outputs == __o.outputs
    def __hash__(self) -> int:
        return super().__hash__()

