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

# Module: ConfluxExtern

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def PullToOption(p):
        source0_ = p
        if True:
            if source0_.is_Yield:
                d_0_a_ = source0_.action
                return ConfluxContract.Option_Some(d_0_a_)
        if True:
            return ConfluxContract.Option_None()

    @staticmethod
    def PushToOption(p):
        source0_ = p
        if True:
            if source0_.is_Emit:
                d_0_e_ = source0_.effect
                return ConfluxContract.Option_Some(d_0_e_)
        if True:
            return ConfluxContract.Option_None()

    @staticmethod
    def RunSink(sink, xs):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif True:
                    source0_ = sink((xs)[0])
                    if True:
                        if source0_.is_Emit:
                            d_1_e_ = source0_.effect
                            return (_dafny.SeqWithoutIsStrInference([d_1_e_])) + (default__.RunSink(sink, _dafny.SeqWithoutIsStrInference((xs)[1::])))
                    if True:
                        return _dafny.SeqWithoutIsStrInference([])
                break


class Pull:
    @classmethod
    def default(cls, ):
        return lambda: Pull_Eof()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Yield(self) -> bool:
        return isinstance(self, Pull_Yield)
    @property
    def is_Eof(self) -> bool:
        return isinstance(self, Pull_Eof)

class Pull_Yield(Pull, NamedTuple('Yield', [('action', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxExtern.Pull.Yield({_dafny.string_of(self.action)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Pull_Yield) and self.action == __o.action
    def __hash__(self) -> int:
        return super().__hash__()

class Pull_Eof(Pull, NamedTuple('Eof', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxExtern.Pull.Eof'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Pull_Eof)
    def __hash__(self) -> int:
        return super().__hash__()


class Push:
    @classmethod
    def default(cls, ):
        return lambda: Push_Closed()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Emit(self) -> bool:
        return isinstance(self, Push_Emit)
    @property
    def is_Closed(self) -> bool:
        return isinstance(self, Push_Closed)

class Push_Emit(Push, NamedTuple('Emit', [('effect', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxExtern.Push.Emit({_dafny.string_of(self.effect)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Push_Emit) and self.effect == __o.effect
    def __hash__(self) -> int:
        return super().__hash__()

class Push_Closed(Push, NamedTuple('Closed', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxExtern.Push.Closed'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Push_Closed)
    def __hash__(self) -> int:
        return super().__hash__()

