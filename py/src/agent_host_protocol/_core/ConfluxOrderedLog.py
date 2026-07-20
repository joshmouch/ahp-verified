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

# Module: ConfluxOrderedLog

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def ContainsK(keyOf, xs, k):
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return False
                elif (keyOf((xs)[0])) == (k):
                    return True
                elif True:
                    in0_ = keyOf
                    in1_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    in2_ = k
                    keyOf = in0_
                    xs = in1_
                    k = in2_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def FindK(keyOf, xs, k):
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return _dafny.SeqWithoutIsStrInference([])
                elif (keyOf((xs)[0])) == (k):
                    return _dafny.SeqWithoutIsStrInference([(xs)[0]])
                elif True:
                    in0_ = keyOf
                    in1_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    in2_ = k
                    keyOf = in0_
                    xs = in1_
                    k = in2_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def TakeListedK(keyOf, xs, order, taken):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(order)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif (((order)[0]) not in (taken)) and (default__.ContainsK(keyOf, xs, (order)[0])):
                    d_0___accumulator_ = (d_0___accumulator_) + (default__.FindK(keyOf, xs, (order)[0]))
                    in0_ = keyOf
                    in1_ = xs
                    in2_ = _dafny.SeqWithoutIsStrInference((order)[1::])
                    in3_ = (taken) | (_dafny.Set({(order)[0]}))
                    keyOf = in0_
                    xs = in1_
                    order = in2_
                    taken = in3_
                    raise _dafny.TailCall()
                elif True:
                    in4_ = keyOf
                    in5_ = xs
                    in6_ = _dafny.SeqWithoutIsStrInference((order)[1::])
                    in7_ = taken
                    keyOf = in4_
                    xs = in5_
                    order = in6_
                    taken = in7_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def RestK(keyOf, xs, order):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + ((_dafny.SeqWithoutIsStrInference([]) if (keyOf((xs)[0])) in (order) else _dafny.SeqWithoutIsStrInference([(xs)[0]])))
                    in0_ = keyOf
                    in1_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    in2_ = order
                    keyOf = in0_
                    xs = in1_
                    order = in2_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def SeqReorderByKeys(keyOf, order, xs):
        return (default__.TakeListedK(keyOf, xs, order, _dafny.Set({}))) + (default__.RestK(keyOf, xs, order))

    @staticmethod
    def SeqKeepThrough(keyOf, k, xs):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif (keyOf((xs)[0])) == (k):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([(xs)[0]]))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([(xs)[0]]))
                    in0_ = keyOf
                    in1_ = k
                    in2_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    keyOf = in0_
                    k = in1_
                    xs = in2_
                    raise _dafny.TailCall()
                break

