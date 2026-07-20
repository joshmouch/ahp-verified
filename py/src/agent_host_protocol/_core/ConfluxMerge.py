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

# Module: ConfluxMerge

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Merge(rA, rB):
        def lambda0_(d_0_rA_, d_1_rB_):
            def lambda1_(d_2_s_, d_3_e_):
                def lambda2_():
                    source0_ = d_3_e_
                    if True:
                        if source0_.is_Left:
                            d_4_a_ = source0_.l
                            return (d_0_rA_((d_2_s_)[0], d_4_a_), (d_2_s_)[1])
                    if True:
                        d_5_b_ = source0_.r
                        return ((d_2_s_)[0], d_1_rB_((d_2_s_)[1], d_5_b_))

                return lambda2_()

            return lambda1_

        return lambda0_(rA, rB)


class Either:
    @classmethod
    def default(cls, default_A):
        return lambda: Either_Left(default_A())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Left(self) -> bool:
        return isinstance(self, Either_Left)
    @property
    def is_Right(self) -> bool:
        return isinstance(self, Either_Right)

class Either_Left(Either, NamedTuple('Left', [('l', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxMerge.Either.Left({_dafny.string_of(self.l)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Either_Left) and self.l == __o.l
    def __hash__(self) -> int:
        return super().__hash__()

class Either_Right(Either, NamedTuple('Right', [('r', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxMerge.Either.Right({_dafny.string_of(self.r)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Either_Right) and self.r == __o.r
    def __hash__(self) -> int:
        return super().__hash__()

