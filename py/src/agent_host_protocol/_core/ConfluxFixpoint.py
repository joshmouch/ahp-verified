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

# Module: ConfluxFixpoint

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Step(succ, S):
        return (S) | (succ(S))

    @staticmethod
    def Closure(succ, U, S):
        while True:
            with _dafny.label():
                d_0_next_ = default__.Step(succ, S)
                if (d_0_next_) == (S):
                    return S
                elif True:
                    in0_ = succ
                    in1_ = U
                    in2_ = d_0_next_
                    succ = in0_
                    U = in1_
                    S = in2_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def Worklist(succ, U, frontier, visited):
        while True:
            with _dafny.label():
                d_0_f_ = (frontier) - (visited)
                if (d_0_f_) == (_dafny.Set({})):
                    return visited
                elif True:
                    in0_ = succ
                    in1_ = U
                    in2_ = succ(d_0_f_)
                    in3_ = (visited) | (d_0_f_)
                    succ = in0_
                    U = in1_
                    frontier = in2_
                    visited = in3_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def StepOne(A):
        if (0) in (A):
            return _dafny.Set({1})
        elif True:
            return _dafny.Set({})

    @staticmethod
    def ClosureClamp(succ, U, S, fuel):
        while True:
            with _dafny.label():
                d_0_next_ = ((S) | (succ(S))).intersection(U)
                if ((d_0_next_) == (S)) or ((fuel) == (0)):
                    return S
                elif True:
                    in0_ = succ
                    in1_ = U
                    in2_ = d_0_next_
                    in3_ = (fuel) - (1)
                    succ = in0_
                    U = in1_
                    S = in2_
                    fuel = in3_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def ClosureTotal(succ, U, S):
        return default__.ClosureClamp(succ, U, (S).intersection(U), len(U))

    @staticmethod
    def TwoHop(s):
        return ((_dafny.Set({1}) if (0) in (s) else _dafny.Set({}))) | ((_dafny.Set({2}) if (1) in (s) else _dafny.Set({})))

