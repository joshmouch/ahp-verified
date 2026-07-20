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

# Module: ConfluxDelimited

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Intercalate(sep, xss):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xss)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif (len(xss)) == (1):
                    return (d_0___accumulator_) + ((xss)[0])
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (((xss)[0]) + (_dafny.SeqWithoutIsStrInference([sep])))
                    in0_ = sep
                    in1_ = _dafny.SeqWithoutIsStrInference((xss)[1::])
                    sep = in0_
                    xss = in1_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def Flatten(xss):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xss)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + ((xss)[0])
                    in0_ = _dafny.SeqWithoutIsStrInference((xss)[1::])
                    xss = in0_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def IndexOf(sep, xs):
        d_0___accumulator_ = 0
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (0) + (d_0___accumulator_)
                elif ((xs)[0]) == (sep):
                    return (0) + (d_0___accumulator_)
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (1)
                    in0_ = sep
                    in1_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    sep = in0_
                    xs = in1_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def Split(sep, xs):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                d_1_i_ = default__.IndexOf(sep, xs)
                if (d_1_i_) == (len(xs)):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([xs]))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference((xs)[:d_1_i_:])]))
                    in0_ = sep
                    in1_ = _dafny.SeqWithoutIsStrInference((xs)[(d_1_i_) + (1)::])
                    sep = in0_
                    xs = in1_
                    raise _dafny.TailCall()
                break

