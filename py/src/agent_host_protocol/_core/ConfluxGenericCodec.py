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

# Module: ConfluxGenericCodec

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def IdCodec():
        def lambda0_(d_0_j_):
            return d_0_j_

        def lambda1_(d_1_j_):
            return Option_Some(d_1_j_)

        return Codec_Codec(lambda0_, lambda1_)

    @staticmethod
    def DoubleCodec():
        def lambda0_(d_0_n_):
            return (d_0_n_) + (d_0_n_)

        def lambda1_(d_1_j_):
            return (Option_Some(_dafny.euclidian_division(d_1_j_, 2)) if (_dafny.euclidian_modulus(d_1_j_, 2)) == (0) else Option_None())

        return Codec_Codec(lambda0_, lambda1_)


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
        return f'ConfluxGenericCodec.Option.None'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Option_None)
    def __hash__(self) -> int:
        return super().__hash__()

class Option_Some(Option, NamedTuple('Some', [('value', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxGenericCodec.Option.Some({_dafny.string_of(self.value)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Option_Some) and self.value == __o.value
    def __hash__(self) -> int:
        return super().__hash__()


class Codec:
    @classmethod
    def default(cls, default_J):
        return lambda: Codec_Codec((lambda x8: default_J()), (lambda x9: Option.default()()))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Codec(self) -> bool:
        return isinstance(self, Codec_Codec)

class Codec_Codec(Codec, NamedTuple('Codec', [('encode', Any), ('decode', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxGenericCodec.Codec.Codec({_dafny.string_of(self.encode)}, {_dafny.string_of(self.decode)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Codec_Codec) and self.encode == __o.encode and self.decode == __o.decode
    def __hash__(self) -> int:
        return super().__hash__()

