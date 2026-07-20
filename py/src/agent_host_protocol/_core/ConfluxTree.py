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

# Module: ConfluxTree

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def RoseMap(f, t):
        return Rose_Rose(f((t).value), default__.MapChildren(f, (t).children))

    @staticmethod
    def MapChildren(f, ts):
        if (len(ts)) == (0):
            return _dafny.SeqWithoutIsStrInference([])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([default__.RoseMap(f, (ts)[0])])) + (default__.MapChildren(f, _dafny.SeqWithoutIsStrInference((ts)[1::])))

    @staticmethod
    def Dbl(x):
        return (x) * (2)


class Rose:
    @classmethod
    def default(cls, default_T):
        return lambda: Rose_Rose(default_T(), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Rose(self) -> bool:
        return isinstance(self, Rose_Rose)

class Rose_Rose(Rose, NamedTuple('Rose', [('value', Any), ('children', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxTree.Rose.Rose({_dafny.string_of(self.value)}, {_dafny.string_of(self.children)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Rose_Rose) and self.value == __o.value and self.children == __o.children
    def __hash__(self) -> int:
        return super().__hash__()

