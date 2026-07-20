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

# Module: ConfluxConvergence

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def FoldHost(r, h, xs):
        def lambda0_(d_0_r_):
            def lambda1_(d_1_hs_, d_2_a_):
                return ConfluxChannel.default__.Accept(d_0_r_, d_1_hs_, d_2_a_)

            return lambda1_

        return ConfluxContract.default__.Fold(lambda0_(r), h, xs)

    @staticmethod
    def FoldClient(r, c, xs):
        def lambda0_(d_0_r_):
            def lambda1_(d_1_cs_, d_2_a_):
                return ConfluxChannel.default__.Receive(d_0_r_, d_1_cs_, d_2_a_)

            return lambda1_

        return ConfluxContract.default__.Fold(lambda0_(r), c, xs)

    @staticmethod
    def InitClients(s0, n):
        return _dafny.SeqWithoutIsStrInference([ConfluxChannel.ClientState_ClientState(s0, 0) for d_0___v4_ in range(n)])

    @staticmethod
    def FanoutFold(r, clients, xs):
        return _dafny.SeqWithoutIsStrInference([default__.FoldClient(r, (clients)[d_0_i_], xs) for d_0_i_ in range(len(clients))])

    @staticmethod
    def IntAdd():
        def lambda0_(d_0_s_, d_1_a_):
            return (d_0_s_) + (d_1_a_)

        return lambda0_

