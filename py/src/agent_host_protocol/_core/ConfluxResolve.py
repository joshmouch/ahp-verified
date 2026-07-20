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
import ConfluxFixpoint as ConfluxFixpoint
import ConfluxGenericCodec as ConfluxGenericCodec

# Module: ConfluxResolve

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def ResolveLookup(rk, snapshot):
        source0_ = rk
        if True:
            if source0_.is_None:
                return ResolveVerdict_NoKeyResolved()
        if True:
            d_0_k_ = source0_.value
            if (d_0_k_) in (snapshot):
                return ResolveVerdict_Matched((snapshot)[d_0_k_])
            elif True:
                return ResolveVerdict_KeyResolvedButAbsent()


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
        return f'ConfluxResolve.Option.None'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Option_None)
    def __hash__(self) -> int:
        return super().__hash__()

class Option_Some(Option, NamedTuple('Some', [('value', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResolve.Option.Some({_dafny.string_of(self.value)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Option_Some) and self.value == __o.value
    def __hash__(self) -> int:
        return super().__hash__()


class ResolveVerdict:
    @classmethod
    def default(cls, ):
        return lambda: ResolveVerdict_NoKeyResolved()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_NoKeyResolved(self) -> bool:
        return isinstance(self, ResolveVerdict_NoKeyResolved)
    @property
    def is_KeyResolvedButAbsent(self) -> bool:
        return isinstance(self, ResolveVerdict_KeyResolvedButAbsent)
    @property
    def is_Matched(self) -> bool:
        return isinstance(self, ResolveVerdict_Matched)

class ResolveVerdict_NoKeyResolved(ResolveVerdict, NamedTuple('NoKeyResolved', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResolve.ResolveVerdict.NoKeyResolved'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResolveVerdict_NoKeyResolved)
    def __hash__(self) -> int:
        return super().__hash__()

class ResolveVerdict_KeyResolvedButAbsent(ResolveVerdict, NamedTuple('KeyResolvedButAbsent', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResolve.ResolveVerdict.KeyResolvedButAbsent'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResolveVerdict_KeyResolvedButAbsent)
    def __hash__(self) -> int:
        return super().__hash__()

class ResolveVerdict_Matched(ResolveVerdict, NamedTuple('Matched', [('value', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxResolve.ResolveVerdict.Matched({_dafny.string_of(self.value)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResolveVerdict_Matched) and self.value == __o.value
    def __hash__(self) -> int:
        return super().__hash__()

