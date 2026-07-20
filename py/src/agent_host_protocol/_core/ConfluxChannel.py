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

# Module: ConfluxChannel

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Accept(r, h, a):
        return HostState_HostState(r((h).canonical, a), ((h).nextSeq) + (1))

    @staticmethod
    def Receive(r, c, a):
        return ClientState_ClientState(r((c).mirror, a), ((c).lastSeq) + (1))

    @staticmethod
    def AcceptFold(r, h, actions):
        while True:
            with _dafny.label():
                if (len(actions)) == (0):
                    return h
                elif True:
                    in0_ = r
                    in1_ = default__.Accept(r, h, (actions)[0])
                    in2_ = _dafny.SeqWithoutIsStrInference((actions)[1::])
                    r = in0_
                    h = in1_
                    actions = in2_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def ReceiveFold(r, c, actions):
        while True:
            with _dafny.label():
                if (len(actions)) == (0):
                    return c
                elif True:
                    in0_ = r
                    in1_ = default__.Receive(r, c, (actions)[0])
                    in2_ = _dafny.SeqWithoutIsStrInference((actions)[1::])
                    r = in0_
                    c = in1_
                    actions = in2_
                    raise _dafny.TailCall()
                break


class HostState:
    @classmethod
    def default(cls, default_S):
        return lambda: HostState_HostState(default_S(), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_HostState(self) -> bool:
        return isinstance(self, HostState_HostState)

class HostState_HostState(HostState, NamedTuple('HostState', [('canonical', Any), ('nextSeq', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxChannel.HostState.HostState({_dafny.string_of(self.canonical)}, {_dafny.string_of(self.nextSeq)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, HostState_HostState) and self.canonical == __o.canonical and self.nextSeq == __o.nextSeq
    def __hash__(self) -> int:
        return super().__hash__()


class ClientState:
    @classmethod
    def default(cls, default_S):
        return lambda: ClientState_ClientState(default_S(), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ClientState(self) -> bool:
        return isinstance(self, ClientState_ClientState)

class ClientState_ClientState(ClientState, NamedTuple('ClientState', [('mirror', Any), ('lastSeq', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxChannel.ClientState.ClientState({_dafny.string_of(self.mirror)}, {_dafny.string_of(self.lastSeq)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ClientState_ClientState) and self.mirror == __o.mirror and self.lastSeq == __o.lastSeq
    def __hash__(self) -> int:
        return super().__hash__()

