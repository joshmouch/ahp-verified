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

# Module: ConfluxPrune

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def PrunableRoute(keyOf, op):
        def lambda0_(d_0_keyOf_, d_1_op_):
            def lambda1_(d_2_m_, d_3_a_):
                def iife0_(_pat_let24_0):
                    def iife1_(d_4_k_):
                        def iife2_(_pat_let25_0):
                            def iife3_(d_5_cur_):
                                def lambda2_():
                                    source0_ = d_1_op_(d_5_cur_, d_3_a_)
                                    if True:
                                        if source0_.is_NoOp:
                                            return d_2_m_
                                    if True:
                                        if source0_.is_Install:
                                            d_6_s_ = source0_.next_
                                            return (d_2_m_).set(d_4_k_, d_6_s_)
                                    if True:
                                        return (d_2_m_) - (_dafny.Set({d_4_k_}))

                                return lambda2_()
                            return iife3_(_pat_let25_0)
                        return iife2_((ConfluxContract.Option_Some((d_2_m_)[d_4_k_]) if (d_4_k_) in (d_2_m_) else ConfluxContract.Option_None()))
                    return iife1_(_pat_let24_0)
                return iife0_(d_0_keyOf_(d_3_a_))

            return lambda1_

        return lambda0_(keyOf, op)


class RouteOp:
    @classmethod
    def default(cls, ):
        return lambda: RouteOp_NoOp()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_NoOp(self) -> bool:
        return isinstance(self, RouteOp_NoOp)
    @property
    def is_Install(self) -> bool:
        return isinstance(self, RouteOp_Install)
    @property
    def is_Remove(self) -> bool:
        return isinstance(self, RouteOp_Remove)

class RouteOp_NoOp(RouteOp, NamedTuple('NoOp', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxPrune.RouteOp.NoOp'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RouteOp_NoOp)
    def __hash__(self) -> int:
        return super().__hash__()

class RouteOp_Install(RouteOp, NamedTuple('Install', [('next_', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxPrune.RouteOp.Install({_dafny.string_of(self.next_)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RouteOp_Install) and self.next_ == __o.next_
    def __hash__(self) -> int:
        return super().__hash__()

class RouteOp_Remove(RouteOp, NamedTuple('Remove', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxPrune.RouteOp.Remove'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RouteOp_Remove)
    def __hash__(self) -> int:
        return super().__hash__()

