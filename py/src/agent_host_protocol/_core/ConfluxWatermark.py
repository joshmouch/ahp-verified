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
import ConfluxResolve as ConfluxResolve

# Module: ConfluxWatermark

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Slack(c):
        return ((c).cap) - ((c).pos)

    @staticmethod
    def TryAdvance(c):
        if ((c).pos) < ((c).cap):
            return (Cap_Cap(((c).pos) + (1), (c).cap), AdvOutcome_Advanced())
        elif True:
            return (c, AdvOutcome_Declined())

    @staticmethod
    def AdvanceN(c, n):
        while True:
            with _dafny.label():
                if (n) == (0):
                    return c
                elif True:
                    in0_ = (default__.TryAdvance(c))[0]
                    in1_ = (n) - (1)
                    c = in0_
                    n = in1_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def Raise(c, incoming):
        return Cap_Cap((c).pos, (incoming if (incoming) >= ((c).cap) else (c).cap))


class Cap:
    @classmethod
    def default(cls, ):
        return lambda: Cap_Cap(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Cap(self) -> bool:
        return isinstance(self, Cap_Cap)

class Cap_Cap(Cap, NamedTuple('Cap', [('pos', Any), ('cap', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxWatermark.Cap.Cap({_dafny.string_of(self.pos)}, {_dafny.string_of(self.cap)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Cap_Cap) and self.pos == __o.pos and self.cap == __o.cap
    def __hash__(self) -> int:
        return super().__hash__()


class AdvOutcome:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [AdvOutcome_Advanced(), AdvOutcome_Declined()]
    @classmethod
    def default(cls, ):
        return lambda: AdvOutcome_Advanced()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Advanced(self) -> bool:
        return isinstance(self, AdvOutcome_Advanced)
    @property
    def is_Declined(self) -> bool:
        return isinstance(self, AdvOutcome_Declined)

class AdvOutcome_Advanced(AdvOutcome, NamedTuple('Advanced', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxWatermark.AdvOutcome.Advanced'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AdvOutcome_Advanced)
    def __hash__(self) -> int:
        return super().__hash__()

class AdvOutcome_Declined(AdvOutcome, NamedTuple('Declined', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxWatermark.AdvOutcome.Declined'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AdvOutcome_Declined)
    def __hash__(self) -> int:
        return super().__hash__()

