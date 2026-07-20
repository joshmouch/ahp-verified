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

# Module: ResourceWatch

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def ApplyToResourceWatch(s, a, now):
        source0_ = a
        if True:
            if source0_.is_RWChanged:
                return (s, AhpSkeleton.ReduceOutcome_Applied())
        if True:
            return (s, AhpSkeleton.ReduceOutcome_NoOp())

    @staticmethod
    def IsUnknownRW(a):
        return (a).is_RWUnknown

    @staticmethod
    def apply1(s, a):
        return (default__.ApplyToResourceWatch(s, a, 9999))[0]

    @staticmethod
    def fold(s, acts):
        return ConfluxContract.default__.Fold(default__.apply1, s, acts)

    @staticmethod
    def RunCorpus():
        pass_: int = int(0)
        total: int = int(0)
        total = 2
        pass_ = 0
        d_0_s0_: ResourceWatchState
        d_0_s0_ = ResourceWatchState_ResourceWatchState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///workspace")), True)
        d_1_changes_: ConfluxCodec.Json
        d_1_changes_ = ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "items")): ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uri")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///workspace/a.txt"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "type")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "added")))}))]))}))
        if (default__.apply1(d_0_s0_, ResourceWatchAction_RWChanged(d_1_changes_))) == (d_0_s0_):
            pass_ = (pass_) + (1)
        d_2_s1_: ResourceWatchState
        d_2_s1_ = ResourceWatchState_ResourceWatchState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///workspace")), False)
        if (default__.apply1(d_2_s1_, ResourceWatchAction_RWUnknown(ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "type")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "resourceWatch/unknownAction")))}))))) == (d_2_s1_):
            pass_ = (pass_) + (1)
        return pass_, total


class ResourceWatchState:
    @classmethod
    def default(cls, ):
        return lambda: ResourceWatchState_ResourceWatchState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ResourceWatchState(self) -> bool:
        return isinstance(self, ResourceWatchState_ResourceWatchState)

class ResourceWatchState_ResourceWatchState(ResourceWatchState, NamedTuple('ResourceWatchState', [('root', Any), ('recursive', Any)])):
    def __dafnystr__(self) -> str:
        return f'ResourceWatch.ResourceWatchState.ResourceWatchState({self.root.VerbatimString(True)}, {_dafny.string_of(self.recursive)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResourceWatchState_ResourceWatchState) and self.root == __o.root and self.recursive == __o.recursive
    def __hash__(self) -> int:
        return super().__hash__()


class ResourceWatchAction:
    @classmethod
    def default(cls, ):
        return lambda: ResourceWatchAction_RWChanged(ConfluxCodec.Json.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RWChanged(self) -> bool:
        return isinstance(self, ResourceWatchAction_RWChanged)
    @property
    def is_RWUnknown(self) -> bool:
        return isinstance(self, ResourceWatchAction_RWUnknown)

class ResourceWatchAction_RWChanged(ResourceWatchAction, NamedTuple('RWChanged', [('changes', Any)])):
    def __dafnystr__(self) -> str:
        return f'ResourceWatch.ResourceWatchAction.RWChanged({_dafny.string_of(self.changes)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResourceWatchAction_RWChanged) and self.changes == __o.changes
    def __hash__(self) -> int:
        return super().__hash__()

class ResourceWatchAction_RWUnknown(ResourceWatchAction, NamedTuple('RWUnknown', [('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'ResourceWatch.ResourceWatchAction.RWUnknown({_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResourceWatchAction_RWUnknown) and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()

