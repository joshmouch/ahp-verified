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

# Module: ConfluxSeqRouteMerge

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def SeqUpsertByIdWith(keyOf, combine, xs, v):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([combine(ConfluxContract.Option_None(), v)]))
                elif (keyOf((xs)[0])) == (keyOf(v)):
                    return (d_0___accumulator_) + ((_dafny.SeqWithoutIsStrInference([combine(ConfluxContract.Option_Some((xs)[0]), v)])) + (_dafny.SeqWithoutIsStrInference((xs)[1::])))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([(xs)[0]]))
                    in0_ = keyOf
                    in1_ = combine
                    in2_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    in3_ = v
                    keyOf = in0_
                    combine = in1_
                    xs = in2_
                    v = in3_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def MergeUpsertById(keyOf, combine, xs, v):
        return default__.SeqUpsertByIdWith(keyOf, combine, xs, v)

    @staticmethod
    def UpsertWithOp(combine):
        def lambda0_(d_0_combine_):
            def lambda1_(d_1_cur_, d_2_a_):
                return ConfluxPrune.RouteOp_Install(d_0_combine_(d_1_cur_, d_2_a_))

            return lambda1_

        return lambda0_(combine)

