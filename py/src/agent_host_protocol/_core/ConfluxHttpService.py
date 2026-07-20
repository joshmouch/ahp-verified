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
import ConfluxJsonRpc as ConfluxJsonRpc

# Module: ConfluxHttpService

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Contains(text, wanted):
        return ConfluxSequence.default__.Contains(text, wanted)

    @staticmethod
    def EndsWith(text, suffix):
        return ConfluxSequence.default__.EndsWith(text, suffix)

    @staticmethod
    def PathOnly(path):
        return _dafny.SeqWithoutIsStrInference((path)[:ConfluxDelimited.default__.IndexOf(_dafny.CodePoint('?'), path):])

    @staticmethod
    def MethodMatches(pattern, httpMethod):
        source0_ = pattern
        if True:
            if source0_.is_ExactMethod:
                d_0_value_ = source0_.value
                return (d_0_value_) == (httpMethod)
        if True:
            return True

    @staticmethod
    def PathMatches(pattern, path):
        source0_ = pattern
        if True:
            if source0_.is_ExactPath:
                d_0_value_ = source0_.value
                return (d_0_value_) == (path)
        if True:
            return True

    @staticmethod
    def HttpRouteStep(state, route, httpMethod, path):
        source0_ = state
        if True:
            if source0_.is_RouteSelected:
                d_0_handler_ = source0_.handler
                return state
        if True:
            if source0_.is_RouteSearching:
                if not(default__.PathMatches((route).pathPattern, path)):
                    return state
                elif default__.MethodMatches((route).methodPattern, httpMethod):
                    return HttpRouteFold_RouteSelected((route).handler)
                elif True:
                    return HttpRouteFold_RoutePathSeen()
        if True:
            if (default__.PathMatches((route).pathPattern, path)) and (default__.MethodMatches((route).methodPattern, httpMethod)):
                return HttpRouteFold_RouteSelected((route).handler)
            elif True:
                return state

    @staticmethod
    def FoldHttpRoutes(routes, httpMethod, path, state):
        while True:
            with _dafny.label():
                if ((len(routes)) == (0)) or ((state).is_RouteSelected):
                    return state
                elif True:
                    in0_ = _dafny.SeqWithoutIsStrInference((routes)[1::])
                    in1_ = httpMethod
                    in2_ = path
                    in3_ = default__.HttpRouteStep(state, (routes)[0], httpMethod, path)
                    routes = in0_
                    httpMethod = in1_
                    path = in2_
                    state = in3_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def ResolveHttpRoute(routes, httpMethod, rawPath):
        d_0_path_ = default__.PathOnly(rawPath)
        source0_ = default__.FoldHttpRoutes(routes, httpMethod, d_0_path_, HttpRouteFold_RouteSearching())
        if True:
            if source0_.is_RouteSelected:
                d_1_handler_ = source0_.handler
                return HttpRouteResolution_RouteMatched(d_1_handler_)
        if True:
            if source0_.is_RoutePathSeen:
                return HttpRouteResolution_RouteMethodNotAllowed()
        if True:
            return HttpRouteResolution_RouteNotFound()

    @staticmethod
    def RequestRouteStep(state, route, name):
        source0_ = state
        if True:
            if source0_.is_RequestSelected:
                d_0_handler_ = source0_.handler
                return state
        if True:
            if ((route).name) == (name):
                return RequestRouteFold_RequestSelected((route).handler)
            elif True:
                return state

    @staticmethod
    def FoldRequestRoutes(routes, name, state):
        while True:
            with _dafny.label():
                if ((len(routes)) == (0)) or ((state).is_RequestSelected):
                    return state
                elif True:
                    in0_ = _dafny.SeqWithoutIsStrInference((routes)[1::])
                    in1_ = name
                    in2_ = default__.RequestRouteStep(state, (routes)[0], name)
                    routes = in0_
                    name = in1_
                    state = in2_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def ResolveRequestRoute(routes, name):
        source0_ = default__.FoldRequestRoutes(routes, name, RequestRouteFold_RequestSearching())
        if True:
            if source0_.is_RequestSelected:
                d_0_handler_ = source0_.handler
                return RequestRouteResolution_RequestMatched(d_0_handler_)
        if True:
            return RequestRouteResolution_RequestNotFound()

    @staticmethod
    def StaticName(path):
        d_0_clean_ = default__.PathOnly(path)
        if (d_0_clean_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/"))):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "index.html"))
        elif ((len(d_0_clean_)) > (0)) and (((d_0_clean_)[0]) == (_dafny.CodePoint('/'))):
            return _dafny.SeqWithoutIsStrInference((d_0_clean_)[1::])
        elif True:
            return d_0_clean_

    @staticmethod
    def SafeStaticName(name):
        return (((name) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))) and (not(default__.Contains(name, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "..")))))) and (not(default__.Contains(name, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\\")))))

    @staticmethod
    def Mime(name):
        if default__.EndsWith(name, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ".html"))):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "text/html; charset=utf-8"))
        elif default__.EndsWith(name, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ".js"))):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "text/javascript; charset=utf-8"))
        elif default__.EndsWith(name, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ".css"))):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "text/css; charset=utf-8"))
        elif default__.EndsWith(name, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ".json"))):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "application/json; charset=utf-8"))
        elif default__.EndsWith(name, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ".svg"))):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "image/svg+xml"))
        elif default__.EndsWith(name, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ".png"))):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "image/png"))
        elif True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "application/octet-stream"))

    @staticmethod
    def EmptyChanges():
        return ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "added")): ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([])), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "removed")): ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([])), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "changed")): ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([]))}))

    @staticmethod
    def JsonArray(value):
        if (value).is_JArr:
            return (value).elems
        elif True:
            return _dafny.SeqWithoutIsStrInference([])

    @staticmethod
    def StringFieldsFrom(items, fieldName):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(items)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif True:
                    d_1_value_ = ConfluxCodec.default__.AsStr(ConfluxCodec.default__.Field((items)[0], fieldName))
                    d_0___accumulator_ = (d_0___accumulator_) + ((_dafny.SeqWithoutIsStrInference([]) if (d_1_value_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))) else _dafny.SeqWithoutIsStrInference([d_1_value_])))
                    in0_ = _dafny.SeqWithoutIsStrInference((items)[1::])
                    in1_ = fieldName
                    items = in0_
                    fieldName = in1_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def StringFields(value, fieldName):
        return default__.StringFieldsFrom(default__.JsonArray(value), fieldName)

    @staticmethod
    def ParseRpc(body):
        ingress: RpcIngress = RpcIngress.default()()
        d_0_parsed_: ConfluxCodec.Option
        out0_: ConfluxCodec.Option
        out0_ = ConfluxJsonRpc.default__.Parse(body)
        d_0_parsed_ = out0_
        if (d_0_parsed_).is_None:
            ingress = RpcIngress_RpcParseError()
        elif True:
            source0_ = (d_0_parsed_).value
            with _dafny.label("match0"):
                if True:
                    if source0_.is_Request:
                        d_1_id_ = source0_.id_
                        d_2_rpcMethod_ = source0_.rpcMethod
                        d_3_params_ = source0_.params
                        ingress = RpcIngress_RpcCall(d_1_id_, d_2_rpcMethod_, d_3_params_)
                        raise _dafny.Break("match0")
                if True:
                    if source0_.is_Invalid:
                        ingress = RpcIngress_RpcInvalid()
                        raise _dafny.Break("match0")
                if True:
                    ingress = RpcIngress_RpcNonRequest()
                pass
        return ingress

    @staticmethod
    def JsonText(value):
        text: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        out0_: _dafny.Seq
        out0_ = ConfluxJsonText.default__.Stringify(value)
        text = out0_
        return text

    @staticmethod
    def RespondJson(request, status, value):
        d_0_text_: _dafny.Seq
        out0_: _dafny.Seq
        out0_ = default__.JsonText(value)
        d_0_text_ = out0_
        ConfluxHttpCapability.default__.Respond(request, status, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "application/json; charset=utf-8")), default__.CorsHeaders, d_0_text_)

    @staticmethod
    def RespondRpcError(request, status, id_, code, message):
        default__.RespondJson(request, status, ConfluxJsonRpc.default__.FailureEnvelope(id_, code, message, ConfluxCodec.Json_JNull()))

    @staticmethod
    def RespondRawRpcResult(request, id_, resultPath, unavailableMessage):
        d_0_present_: bool
        d_1_content_: _dafny.Seq
        out0_: bool
        out1_: _dafny.Seq
        out0_, out1_ = ConfluxIoCapability.default__.ReadFile(resultPath)
        d_0_present_ = out0_
        d_1_content_ = out1_
        if d_0_present_:
            d_2_idText_: _dafny.Seq
            out2_: _dafny.Seq
            out2_ = default__.JsonText(id_)
            d_2_idText_ = out2_
            ConfluxHttpCapability.default__.Respond(request, 200, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "application/json; charset=utf-8")), default__.CorsHeaders, ((((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "{\"jsonrpc\":\"2.0\",\"id\":"))) + (d_2_idText_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ",\"result\":")))) + (d_1_content_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "}"))))
        elif True:
            default__.RespondRpcError(request, 500, id_, -32603, unavailableMessage)

    @staticmethod
    def RespondFile(request, path, contentType, headers, missingStatus, missingBody):
        d_0_present_: bool
        d_1_content_: _dafny.Seq
        out0_: bool
        out1_: _dafny.Seq
        out0_, out1_ = ConfluxIoCapability.default__.ReadFile(path)
        d_0_present_ = out0_
        d_1_content_ = out1_
        if d_0_present_:
            ConfluxHttpCapability.default__.Respond(request, 200, contentType, headers, d_1_content_)
        elif True:
            default__.RespondJson(request, missingStatus, missingBody)

    @staticmethod
    def RespondStatic(request, root, rawPath):
        d_0_name_: _dafny.Seq
        d_0_name_ = default__.StaticName(rawPath)
        if not(default__.SafeStaticName(d_0_name_)):
            default__.RespondJson(request, 404, ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "error")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "not found")))})))
        elif True:
            default__.RespondFile(request, ((root) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/")))) + (d_0_name_), default__.Mime(d_0_name_), _dafny.SeqWithoutIsStrInference([]), 404, ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "error")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "not found")))})))

    @staticmethod
    def RespondCorsPreflight(request):
        ConfluxHttpCapability.default__.Respond(request, 204, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "text/plain")), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Access-Control-Allow-Origin: *")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Access-Control-Allow-Headers: content-type"))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))

    @staticmethod
    def BeginEventStream(request, eventName, value):
        d_0_text_: _dafny.Seq
        out0_: _dafny.Seq
        out0_ = default__.JsonText(value)
        d_0_text_ = out0_
        ConfluxHttpCapability.default__.BeginStream(request, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "text/event-stream")), default__.EventHeaders, ((((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "event: "))) + (eventName)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\ndata: ")))) + (d_0_text_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n\n"))))

    @staticmethod
    def PublishEvent(server, eventName, value):
        d_0_text_: _dafny.Seq
        out0_: _dafny.Seq
        out0_ = default__.JsonText(value)
        d_0_text_ = out0_
        ConfluxHttpCapability.default__.Publish(server, ((((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "event: "))) + (eventName)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\ndata: ")))) + (d_0_text_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n\n"))))

    @staticmethod
    def RefreshJson(cwd, command, arguments, metadataPath, requiredField):
        ok: bool = False
        metadata: ConfluxCodec.Json = ConfluxCodec.Json.default()()
        error: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_0_code_: int
        d_1_output_: _dafny.Seq
        d_2_processError_: _dafny.Seq
        out0_: int
        out1_: _dafny.Seq
        out2_: _dafny.Seq
        out0_, out1_, out2_ = ConfluxIoCapability.default__.RunProcess(cwd, command, arguments)
        d_0_code_ = out0_
        d_1_output_ = out1_
        d_2_processError_ = out2_
        if (d_0_code_) != (0):
            ok = False
            metadata = ConfluxCodec.Json_JNull()
            error = d_2_processError_
            return ok, metadata, error
        d_3_present_: bool
        d_4_text_: _dafny.Seq
        out3_: bool
        out4_: _dafny.Seq
        out3_, out4_ = ConfluxIoCapability.default__.ReadFile(metadataPath)
        d_3_present_ = out3_
        d_4_text_ = out4_
        if not(d_3_present_):
            ok = False
            metadata = ConfluxCodec.Json_JNull()
            error = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "metadata unavailable"))
            return ok, metadata, error
        d_5_parsed_: ConfluxJsonText.ParseResult
        out5_: ConfluxJsonText.ParseResult
        out5_ = ConfluxJsonText.default__.ParseJsonChecked(d_4_text_)
        d_5_parsed_ = out5_
        if ((d_5_parsed_).is_Invalid) or ((ConfluxCodec.default__.AsStr(ConfluxCodec.default__.Field((d_5_parsed_).value, requiredField))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))):
            ok = False
            metadata = ConfluxCodec.Json_JNull()
            error = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "metadata invalid"))
            return ok, metadata, error
        ok = True
        metadata = (d_5_parsed_).value
        error = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        return ok, metadata, error

    @staticmethod
    def DrainWatch(watcher, budget, waitMs):
        while True:
            with _dafny.label():
                if (budget) > (0):
                    d_0_available_: bool
                    d_1_path_: _dafny.Seq
                    out0_: bool
                    out1_: _dafny.Seq
                    out0_, out1_ = ConfluxIoCapability.default__.PollWatch(watcher, waitMs)
                    d_0_available_ = out0_
                    d_1_path_ = out1_
                    if d_0_available_:
                        in0_ = watcher
                        in1_ = (budget) - (1)
                        in2_ = waitMs
                        watcher = in0_
                        budget = in1_
                        waitMs = in2_
                        raise _dafny.TailCall()
                break

    @_dafny.classproperty
    def CorsHeaders(instance):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Access-Control-Allow-Origin: *"))])
    @_dafny.classproperty
    def EventHeaders(instance):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Cache-Control: no-cache")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Access-Control-Allow-Origin: *"))])

class RpcIngress:
    @classmethod
    def default(cls, ):
        return lambda: RpcIngress_RpcCall(ConfluxCodec.Json.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxCodec.Json.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RpcCall(self) -> bool:
        return isinstance(self, RpcIngress_RpcCall)
    @property
    def is_RpcParseError(self) -> bool:
        return isinstance(self, RpcIngress_RpcParseError)
    @property
    def is_RpcInvalid(self) -> bool:
        return isinstance(self, RpcIngress_RpcInvalid)
    @property
    def is_RpcNonRequest(self) -> bool:
        return isinstance(self, RpcIngress_RpcNonRequest)

class RpcIngress_RpcCall(RpcIngress, NamedTuple('RpcCall', [('id_', Any), ('rpcMethod', Any), ('params', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.RpcIngress.RpcCall({_dafny.string_of(self.id_)}, {self.rpcMethod.VerbatimString(True)}, {_dafny.string_of(self.params)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RpcIngress_RpcCall) and self.id_ == __o.id_ and self.rpcMethod == __o.rpcMethod and self.params == __o.params
    def __hash__(self) -> int:
        return super().__hash__()

class RpcIngress_RpcParseError(RpcIngress, NamedTuple('RpcParseError', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.RpcIngress.RpcParseError'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RpcIngress_RpcParseError)
    def __hash__(self) -> int:
        return super().__hash__()

class RpcIngress_RpcInvalid(RpcIngress, NamedTuple('RpcInvalid', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.RpcIngress.RpcInvalid'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RpcIngress_RpcInvalid)
    def __hash__(self) -> int:
        return super().__hash__()

class RpcIngress_RpcNonRequest(RpcIngress, NamedTuple('RpcNonRequest', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.RpcIngress.RpcNonRequest'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RpcIngress_RpcNonRequest)
    def __hash__(self) -> int:
        return super().__hash__()


class MethodPattern:
    @classmethod
    def default(cls, ):
        return lambda: MethodPattern_ExactMethod(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ExactMethod(self) -> bool:
        return isinstance(self, MethodPattern_ExactMethod)
    @property
    def is_AnyMethod(self) -> bool:
        return isinstance(self, MethodPattern_AnyMethod)

class MethodPattern_ExactMethod(MethodPattern, NamedTuple('ExactMethod', [('value', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.MethodPattern.ExactMethod({self.value.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, MethodPattern_ExactMethod) and self.value == __o.value
    def __hash__(self) -> int:
        return super().__hash__()

class MethodPattern_AnyMethod(MethodPattern, NamedTuple('AnyMethod', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.MethodPattern.AnyMethod'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, MethodPattern_AnyMethod)
    def __hash__(self) -> int:
        return super().__hash__()


class PathPattern:
    @classmethod
    def default(cls, ):
        return lambda: PathPattern_ExactPath(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ExactPath(self) -> bool:
        return isinstance(self, PathPattern_ExactPath)
    @property
    def is_AnyPath(self) -> bool:
        return isinstance(self, PathPattern_AnyPath)

class PathPattern_ExactPath(PathPattern, NamedTuple('ExactPath', [('value', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.PathPattern.ExactPath({self.value.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, PathPattern_ExactPath) and self.value == __o.value
    def __hash__(self) -> int:
        return super().__hash__()

class PathPattern_AnyPath(PathPattern, NamedTuple('AnyPath', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.PathPattern.AnyPath'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, PathPattern_AnyPath)
    def __hash__(self) -> int:
        return super().__hash__()


class HttpRoute:
    @classmethod
    def default(cls, default_H):
        return lambda: HttpRoute_HttpRoute(MethodPattern.default()(), PathPattern.default()(), default_H())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_HttpRoute(self) -> bool:
        return isinstance(self, HttpRoute_HttpRoute)

class HttpRoute_HttpRoute(HttpRoute, NamedTuple('HttpRoute', [('methodPattern', Any), ('pathPattern', Any), ('handler', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.HttpRoute.HttpRoute({_dafny.string_of(self.methodPattern)}, {_dafny.string_of(self.pathPattern)}, {_dafny.string_of(self.handler)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, HttpRoute_HttpRoute) and self.methodPattern == __o.methodPattern and self.pathPattern == __o.pathPattern and self.handler == __o.handler
    def __hash__(self) -> int:
        return super().__hash__()


class HttpRouteFold:
    @classmethod
    def default(cls, ):
        return lambda: HttpRouteFold_RouteSearching()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RouteSearching(self) -> bool:
        return isinstance(self, HttpRouteFold_RouteSearching)
    @property
    def is_RoutePathSeen(self) -> bool:
        return isinstance(self, HttpRouteFold_RoutePathSeen)
    @property
    def is_RouteSelected(self) -> bool:
        return isinstance(self, HttpRouteFold_RouteSelected)

class HttpRouteFold_RouteSearching(HttpRouteFold, NamedTuple('RouteSearching', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.HttpRouteFold.RouteSearching'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, HttpRouteFold_RouteSearching)
    def __hash__(self) -> int:
        return super().__hash__()

class HttpRouteFold_RoutePathSeen(HttpRouteFold, NamedTuple('RoutePathSeen', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.HttpRouteFold.RoutePathSeen'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, HttpRouteFold_RoutePathSeen)
    def __hash__(self) -> int:
        return super().__hash__()

class HttpRouteFold_RouteSelected(HttpRouteFold, NamedTuple('RouteSelected', [('handler', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.HttpRouteFold.RouteSelected({_dafny.string_of(self.handler)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, HttpRouteFold_RouteSelected) and self.handler == __o.handler
    def __hash__(self) -> int:
        return super().__hash__()


class HttpRouteResolution:
    @classmethod
    def default(cls, ):
        return lambda: HttpRouteResolution_RouteMethodNotAllowed()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RouteMatched(self) -> bool:
        return isinstance(self, HttpRouteResolution_RouteMatched)
    @property
    def is_RouteMethodNotAllowed(self) -> bool:
        return isinstance(self, HttpRouteResolution_RouteMethodNotAllowed)
    @property
    def is_RouteNotFound(self) -> bool:
        return isinstance(self, HttpRouteResolution_RouteNotFound)

class HttpRouteResolution_RouteMatched(HttpRouteResolution, NamedTuple('RouteMatched', [('handler', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.HttpRouteResolution.RouteMatched({_dafny.string_of(self.handler)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, HttpRouteResolution_RouteMatched) and self.handler == __o.handler
    def __hash__(self) -> int:
        return super().__hash__()

class HttpRouteResolution_RouteMethodNotAllowed(HttpRouteResolution, NamedTuple('RouteMethodNotAllowed', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.HttpRouteResolution.RouteMethodNotAllowed'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, HttpRouteResolution_RouteMethodNotAllowed)
    def __hash__(self) -> int:
        return super().__hash__()

class HttpRouteResolution_RouteNotFound(HttpRouteResolution, NamedTuple('RouteNotFound', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.HttpRouteResolution.RouteNotFound'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, HttpRouteResolution_RouteNotFound)
    def __hash__(self) -> int:
        return super().__hash__()


class RequestRoute:
    @classmethod
    def default(cls, default_H):
        return lambda: RequestRoute_RequestRoute(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), default_H())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RequestRoute(self) -> bool:
        return isinstance(self, RequestRoute_RequestRoute)

class RequestRoute_RequestRoute(RequestRoute, NamedTuple('RequestRoute', [('name', Any), ('handler', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.RequestRoute.RequestRoute({self.name.VerbatimString(True)}, {_dafny.string_of(self.handler)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RequestRoute_RequestRoute) and self.name == __o.name and self.handler == __o.handler
    def __hash__(self) -> int:
        return super().__hash__()


class RequestRouteFold:
    @classmethod
    def default(cls, ):
        return lambda: RequestRouteFold_RequestSearching()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RequestSearching(self) -> bool:
        return isinstance(self, RequestRouteFold_RequestSearching)
    @property
    def is_RequestSelected(self) -> bool:
        return isinstance(self, RequestRouteFold_RequestSelected)

class RequestRouteFold_RequestSearching(RequestRouteFold, NamedTuple('RequestSearching', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.RequestRouteFold.RequestSearching'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RequestRouteFold_RequestSearching)
    def __hash__(self) -> int:
        return super().__hash__()

class RequestRouteFold_RequestSelected(RequestRouteFold, NamedTuple('RequestSelected', [('handler', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.RequestRouteFold.RequestSelected({_dafny.string_of(self.handler)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RequestRouteFold_RequestSelected) and self.handler == __o.handler
    def __hash__(self) -> int:
        return super().__hash__()


class RequestRouteResolution:
    @classmethod
    def default(cls, ):
        return lambda: RequestRouteResolution_RequestNotFound()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RequestMatched(self) -> bool:
        return isinstance(self, RequestRouteResolution_RequestMatched)
    @property
    def is_RequestNotFound(self) -> bool:
        return isinstance(self, RequestRouteResolution_RequestNotFound)

class RequestRouteResolution_RequestMatched(RequestRouteResolution, NamedTuple('RequestMatched', [('handler', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.RequestRouteResolution.RequestMatched({_dafny.string_of(self.handler)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RequestRouteResolution_RequestMatched) and self.handler == __o.handler
    def __hash__(self) -> int:
        return super().__hash__()

class RequestRouteResolution_RequestNotFound(RequestRouteResolution, NamedTuple('RequestNotFound', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxHttpService.RequestRouteResolution.RequestNotFound'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RequestRouteResolution_RequestNotFound)
    def __hash__(self) -> int:
        return super().__hash__()

