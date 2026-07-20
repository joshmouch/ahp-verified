import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_ as module_
import _dafny as _dafny
import System_ as System_
import ConfluxCodec as ConfluxCodec

# Module: ConfluxContract

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Fold(r, s, actions):
        while True:
            with _dafny.label():
                if (len(actions)) == (0):
                    return s
                elif True:
                    in0_ = r
                    in1_ = r(s, (actions)[0])
                    in2_ = _dafny.SeqWithoutIsStrInference((actions)[1::])
                    r = in0_
                    s = in1_
                    actions = in2_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def Iterate(f, x, n):
        while True:
            with _dafny.label():
                if (n) == (0):
                    return x
                elif True:
                    in0_ = f
                    in1_ = f(x)
                    in2_ = (n) - (1)
                    f = in0_
                    x = in1_
                    n = in2_
                    raise _dafny.TailCall()
                break


class Option:
    @classmethod
    def default(cls, ):
        return lambda: Option_None()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_None(self) -> bool:
        return isinstance(self, Option_None)
    @property
    def is_Some(self) -> bool:
        return isinstance(self, Option_Some)

class Option_None(Option, NamedTuple('None_', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxContract.Option.None'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Option_None)
    def __hash__(self) -> int:
        return super().__hash__()

class Option_Some(Option, NamedTuple('Some', [('value', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxContract.Option.Some({_dafny.string_of(self.value)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Option_Some) and self.value == __o.value
    def __hash__(self) -> int:
        return super().__hash__()

