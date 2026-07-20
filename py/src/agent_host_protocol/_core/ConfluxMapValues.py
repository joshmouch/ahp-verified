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

# Module: ConfluxMapValues

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def MapValues(f, m):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: TypeVar('K__')
            for compr_0_ in (m).keys.Elements:
                d_0_k_: TypeVar('K__') = compr_0_
                if (d_0_k_) in (m):
                    coll0_[d_0_k_] = f(d_0_k_, (m)[d_0_k_])
            return _dafny.Map(coll0_)
        return iife0_()
        

    @staticmethod
    def Times10(k, v):
        return (v) * (10)

