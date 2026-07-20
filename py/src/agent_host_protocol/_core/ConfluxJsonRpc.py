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
import ConfluxWatermark as ConfluxWatermark
import ConfluxStateMachine as ConfluxStateMachine
import ConfluxProduct as ConfluxProduct
import ConfluxJoin as ConfluxJoin
import ConfluxDedupe as ConfluxDedupe
import ConfluxBatchRoute as ConfluxBatchRoute
import ConfluxSeqRouteMerge as ConfluxSeqRouteMerge
import ConfluxKeyedOps as ConfluxKeyedOps
import ConfluxFilterKeys as ConfluxFilterKeys
import ConfluxChannel as ConfluxChannel
import ConfluxDurableHistory as ConfluxDurableHistory
import ConfluxCommand as ConfluxCommand
import ConfluxChannelManifest as ConfluxChannelManifest
import ConfluxConvergence as ConfluxConvergence
import ConfluxBudgetConvergence as ConfluxBudgetConvergence
import ConfluxExtern as ConfluxExtern
import ConfluxCommandIdentityCapability as ConfluxCommandIdentityCapability
import ConfluxDecimalText as ConfluxDecimalText
import ConfluxJsonText as ConfluxJsonText
import ConfluxCommandIdentity as ConfluxCommandIdentity
import ConfluxIoCapability as ConfluxIoCapability
import ConfluxResourceSupervisor as ConfluxResourceSupervisor
import ConfluxSocketCapability as ConfluxSocketCapability
import ConfluxWebSocketCapability as ConfluxWebSocketCapability
import ConfluxHttpCapability as ConfluxHttpCapability
import ConfluxServiceLifecycle as ConfluxServiceLifecycle
import ConfluxSequence as ConfluxSequence

# Module: ConfluxJsonRpc

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Object(fields):
        return ConfluxCodec.Json_JObj(fields)

    @staticmethod
    def RequestEnvelope(id_, rpcMethod, params):
        return default__.Object(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsonrpc")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "2.0"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "id")): id_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "method")): ConfluxCodec.Json_JStr(rpcMethod), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "params")): params}))

    @staticmethod
    def NotificationEnvelope(rpcMethod, params):
        return default__.Object(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsonrpc")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "2.0"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "method")): ConfluxCodec.Json_JStr(rpcMethod), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "params")): params}))

    @staticmethod
    def SuccessEnvelope(id_, result):
        return default__.Object(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsonrpc")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "2.0"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "id")): id_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "result")): result}))

    @staticmethod
    def FailureEnvelope(id_, code, message, data):
        return default__.Object(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsonrpc")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "2.0"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "id")): id_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "error")): default__.Object(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "code")): ConfluxCodec.Json_JNum(code), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "message")): ConfluxCodec.Json_JStr(message), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "data")): data}))}))

    @staticmethod
    def NullResult(id_):
        return default__.SuccessEnvelope(id_, ConfluxCodec.Json_JNull())

    @staticmethod
    def Classify(j):
        d_0_id_ = ConfluxCodec.default__.Field(j, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "id")))
        d_1_rpcMethod_ = ConfluxCodec.default__.AsStr(ConfluxCodec.default__.Field(j, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "method"))))
        d_2_params_ = ConfluxCodec.default__.Field(j, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "params")))
        d_3_error_ = ConfluxCodec.default__.Field(j, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "error")))
        if ((d_1_rpcMethod_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))) and (not((d_0_id_).is_JNull)):
            return Message_Request(d_0_id_, d_1_rpcMethod_, d_2_params_)
        elif (d_1_rpcMethod_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))):
            return Message_Notification(d_1_rpcMethod_, d_2_params_)
        elif (not((d_0_id_).is_JNull)) and (not((d_3_error_).is_JNull)):
            return Message_Failure(d_0_id_, d_3_error_, ConfluxCodec.default__.AsStr(ConfluxCodec.default__.Field(d_3_error_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "message")))))
        elif not((d_0_id_).is_JNull):
            return Message_Success(d_0_id_, ConfluxCodec.default__.Field(j, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "result"))))
        elif True:
            return Message_Invalid()

    @staticmethod
    def Parse(text):
        message: ConfluxCodec.Option = ConfluxCodec.Option.default()()
        d_0_parsed_: ConfluxJsonText.ParseResult
        out0_: ConfluxJsonText.ParseResult
        out0_ = ConfluxJsonText.default__.ParseJsonChecked(text)
        d_0_parsed_ = out0_
        source0_ = d_0_parsed_
        with _dafny.label("match0"):
            if True:
                if source0_.is_Parsed:
                    d_1_j_ = source0_.value
                    message = ConfluxCodec.Option_Some(default__.Classify(d_1_j_))
                    raise _dafny.Break("match0")
            if True:
                message = ConfluxCodec.Option_None()
            pass
        return message

    @staticmethod
    def Stringify(j):
        text: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        out0_: _dafny.Seq
        out0_ = ConfluxJsonText.default__.Stringify(j)
        text = out0_
        return text


class Message:
    @classmethod
    def default(cls, ):
        return lambda: Message_Success(ConfluxCodec.Json.default()(), ConfluxCodec.Json.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Success(self) -> bool:
        return isinstance(self, Message_Success)
    @property
    def is_Failure(self) -> bool:
        return isinstance(self, Message_Failure)
    @property
    def is_Notification(self) -> bool:
        return isinstance(self, Message_Notification)
    @property
    def is_Request(self) -> bool:
        return isinstance(self, Message_Request)
    @property
    def is_Invalid(self) -> bool:
        return isinstance(self, Message_Invalid)

class Message_Success(Message, NamedTuple('Success', [('id_', Any), ('result', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxJsonRpc.Message.Success({_dafny.string_of(self.id_)}, {_dafny.string_of(self.result)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Message_Success) and self.id_ == __o.id_ and self.result == __o.result
    def __hash__(self) -> int:
        return super().__hash__()

class Message_Failure(Message, NamedTuple('Failure', [('id_', Any), ('error', Any), ('errorText', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxJsonRpc.Message.Failure({_dafny.string_of(self.id_)}, {_dafny.string_of(self.error)}, {self.errorText.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Message_Failure) and self.id_ == __o.id_ and self.error == __o.error and self.errorText == __o.errorText
    def __hash__(self) -> int:
        return super().__hash__()

class Message_Notification(Message, NamedTuple('Notification', [('rpcMethod', Any), ('params', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxJsonRpc.Message.Notification({self.rpcMethod.VerbatimString(True)}, {_dafny.string_of(self.params)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Message_Notification) and self.rpcMethod == __o.rpcMethod and self.params == __o.params
    def __hash__(self) -> int:
        return super().__hash__()

class Message_Request(Message, NamedTuple('Request', [('id_', Any), ('rpcMethod', Any), ('params', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxJsonRpc.Message.Request({_dafny.string_of(self.id_)}, {self.rpcMethod.VerbatimString(True)}, {_dafny.string_of(self.params)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Message_Request) and self.id_ == __o.id_ and self.rpcMethod == __o.rpcMethod and self.params == __o.params
    def __hash__(self) -> int:
        return super().__hash__()

class Message_Invalid(Message, NamedTuple('Invalid', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxJsonRpc.Message.Invalid'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Message_Invalid)
    def __hash__(self) -> int:
        return super().__hash__()

