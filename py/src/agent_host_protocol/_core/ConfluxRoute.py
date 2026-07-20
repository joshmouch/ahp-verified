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

# Module: ConfluxRoute

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Route(keyOf, r, empty):
        def lambda0_(d_0_keyOf_, d_1_r_, d_2_empty_):
            def lambda1_(d_3_m_, d_4_a_):
                def iife0_(_pat_let469_0):
                    def iife1_(d_5_k_):
                        return (d_3_m_).set(d_5_k_, d_1_r_(((d_3_m_)[d_5_k_] if (d_5_k_) in (d_3_m_) else d_2_empty_), d_4_a_))
                    return iife1_(_pat_let469_0)
                return iife0_(d_0_keyOf_(d_4_a_))

            return lambda1_

        return lambda0_(keyOf, r, empty)

    @staticmethod
    def GuardedRoute(keyOf, guard):
        def lambda0_(d_0_keyOf_, d_1_guard_):
            def lambda1_(d_2_m_, d_3_a_):
                def iife0_(_pat_let470_0):
                    def iife1_(d_4_k_):
                        def iife2_(_pat_let471_0):
                            def iife3_(d_5_cur_):
                                def lambda2_():
                                    source0_ = d_1_guard_(d_5_cur_, d_3_a_)
                                    if True:
                                        if source0_.is_None:
                                            return d_2_m_
                                    if True:
                                        d_6_s2_ = source0_.value
                                        return (d_2_m_).set(d_4_k_, d_6_s2_)

                                return lambda2_()
                            return iife3_(_pat_let471_0)
                        return iife2_((ConfluxContract.Option_Some((d_2_m_)[d_4_k_]) if (d_4_k_) in (d_2_m_) else ConfluxContract.Option_None()))
                    return iife1_(_pat_let470_0)
                return iife0_(d_0_keyOf_(d_3_a_))

            return lambda1_

        return lambda0_(keyOf, guard)

