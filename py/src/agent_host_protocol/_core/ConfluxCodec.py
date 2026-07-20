import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_ as module_
import _dafny as _dafny
import System_ as System_

# Module: ConfluxCodec

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Field(j, key):
        if ((j).is_JObj) and ((key) in ((j).fields)):
            return ((j).fields)[key]
        elif True:
            return Json_JNull()

    @staticmethod
    def AsInt(j):
        if (j).is_JNum:
            return (j).n
        elif True:
            return 0

    @staticmethod
    def AsStr(j):
        if (j).is_JStr:
            return (j).s
        elif True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))

    @staticmethod
    def AsNat(j):
        d_0_x_ = default__.AsInt(j)
        if (d_0_x_) >= (0):
            return d_0_x_
        elif True:
            return 0

    @staticmethod
    def JsonCodec():
        def lambda0_(d_0_j_):
            return d_0_j_

        def lambda1_(d_1_j_):
            return Option_Some(d_1_j_)

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
        return f'ConfluxCodec.Option.None'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Option_None)
    def __hash__(self) -> int:
        return super().__hash__()

class Option_Some(Option, NamedTuple('Some', [('value', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCodec.Option.Some({_dafny.string_of(self.value)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Option_Some) and self.value == __o.value
    def __hash__(self) -> int:
        return super().__hash__()


class Json:
    @classmethod
    def default(cls, ):
        return lambda: Json_JNull()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_JNull(self) -> bool:
        return isinstance(self, Json_JNull)
    @property
    def is_JBool(self) -> bool:
        return isinstance(self, Json_JBool)
    @property
    def is_JNum(self) -> bool:
        return isinstance(self, Json_JNum)
    @property
    def is_JDec(self) -> bool:
        return isinstance(self, Json_JDec)
    @property
    def is_JStr(self) -> bool:
        return isinstance(self, Json_JStr)
    @property
    def is_JArr(self) -> bool:
        return isinstance(self, Json_JArr)
    @property
    def is_JObj(self) -> bool:
        return isinstance(self, Json_JObj)

class Json_JNull(Json, NamedTuple('JNull', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCodec.Json.JNull'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Json_JNull)
    def __hash__(self) -> int:
        return super().__hash__()

class Json_JBool(Json, NamedTuple('JBool', [('b', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCodec.Json.JBool({_dafny.string_of(self.b)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Json_JBool) and self.b == __o.b
    def __hash__(self) -> int:
        return super().__hash__()

class Json_JNum(Json, NamedTuple('JNum', [('n', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCodec.Json.JNum({_dafny.string_of(self.n)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Json_JNum) and self.n == __o.n
    def __hash__(self) -> int:
        return super().__hash__()

class Json_JDec(Json, NamedTuple('JDec', [('mantissa', Any), ('exp', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCodec.Json.JDec({_dafny.string_of(self.mantissa)}, {_dafny.string_of(self.exp)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Json_JDec) and self.mantissa == __o.mantissa and self.exp == __o.exp
    def __hash__(self) -> int:
        return super().__hash__()

class Json_JStr(Json, NamedTuple('JStr', [('s', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCodec.Json.JStr({self.s.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Json_JStr) and self.s == __o.s
    def __hash__(self) -> int:
        return super().__hash__()

class Json_JArr(Json, NamedTuple('JArr', [('elems', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCodec.Json.JArr({_dafny.string_of(self.elems)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Json_JArr) and self.elems == __o.elems
    def __hash__(self) -> int:
        return super().__hash__()

class Json_JObj(Json, NamedTuple('JObj', [('fields', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCodec.Json.JObj({_dafny.string_of(self.fields)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Json_JObj) and self.fields == __o.fields
    def __hash__(self) -> int:
        return super().__hash__()


class Codec:
    @classmethod
    def default(cls, ):
        return lambda: Codec_Codec((lambda x0: Json.default()()), (lambda x1: Option.default()()))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Codec(self) -> bool:
        return isinstance(self, Codec_Codec)

class Codec_Codec(Codec, NamedTuple('Codec', [('encode', Any), ('decode', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCodec.Codec.Codec({_dafny.string_of(self.encode)}, {_dafny.string_of(self.decode)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Codec_Codec) and self.encode == __o.encode and self.decode == __o.decode
    def __hash__(self) -> int:
        return super().__hash__()

