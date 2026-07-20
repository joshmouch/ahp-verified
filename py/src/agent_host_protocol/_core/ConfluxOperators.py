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

# Module: ConfluxOperators

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Map(f, xs):
        return _dafny.SeqWithoutIsStrInference([f((xs)[d_0_i_]) for d_0_i_ in range(len(xs))])

    @staticmethod
    def Filter(p, xs):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + ((_dafny.SeqWithoutIsStrInference([(xs)[0]]) if p((xs)[0]) else _dafny.SeqWithoutIsStrInference([])))
                    in0_ = p
                    in1_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    p = in0_
                    xs = in1_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def MapFold(f, r, init, xs):
        return ConfluxContract.default__.Fold(r, init, default__.Map(f, xs))

