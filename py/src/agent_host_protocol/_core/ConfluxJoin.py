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

# Module: ConfluxJoin

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def JoinCheckpoint(combine, primary, side):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: TypeVar('K__')
            for compr_0_ in (primary).keys.Elements:
                d_0_k_: TypeVar('K__') = compr_0_
                if (d_0_k_) in (primary):
                    coll0_[d_0_k_] = combine(d_0_k_, (primary)[d_0_k_], (ConfluxContract.Option_Some((side)[d_0_k_]) if (d_0_k_) in (side) else ConfluxContract.Option_None()))
            return _dafny.Map(coll0_)
        return iife0_()
        

