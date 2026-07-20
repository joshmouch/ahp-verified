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

# Module: Canvas

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def ApplyToCanvas(s, a, now):
        pat_let_tv0_ = s
        pat_let_tv1_ = s
        pat_let_tv2_ = s
        pat_let_tv3_ = s
        source0_ = a
        if True:
            if source0_.is_Updated:
                d_0_t_ = source0_.title
                d_1_ac_ = source0_.activity
                d_2_uri_ = source0_.contentUri
                d_3_av_ = source0_.availability
                def iife0_(_pat_let8_0):
                    def iife1_(d_4_dt__update__tmp_h0_):
                        def iife2_(_pat_let9_0):
                            def iife3_(d_5_dt__update_havailability_h0_):
                                def iife4_(_pat_let10_0):
                                    def iife5_(d_6_dt__update_hcontentUri_h0_):
                                        def iife6_(_pat_let11_0):
                                            def iife7_(d_7_dt__update_hactivity_h0_):
                                                def iife8_(_pat_let12_0):
                                                    def iife9_(d_8_dt__update_htitle_h0_):
                                                        return CanvasState_CanvasState((d_4_dt__update__tmp_h0_).canvasId, (d_4_dt__update__tmp_h0_).providerId, (d_4_dt__update__tmp_h0_).input_, d_8_dt__update_htitle_h0_, d_7_dt__update_hactivity_h0_, d_6_dt__update_hcontentUri_h0_, d_5_dt__update_havailability_h0_)
                                                    return iife9_(_pat_let12_0)
                                                return iife8_((d_0_t_ if (d_0_t_).is_Some else (pat_let_tv3_).title))
                                            return iife7_(_pat_let11_0)
                                        return iife6_((d_1_ac_ if (d_1_ac_).is_Some else (pat_let_tv2_).activity))
                                    return iife5_(_pat_let10_0)
                                return iife4_((d_2_uri_ if (d_2_uri_).is_Some else (pat_let_tv1_).contentUri))
                            return iife3_(_pat_let9_0)
                        return iife2_(((d_3_av_).value if (d_3_av_).is_Some else (pat_let_tv0_).availability))
                    return iife1_(_pat_let8_0)
                return (iife0_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CloseRequested:
                return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            return (s, AhpSkeleton.ReduceOutcome_NoOp())

    @staticmethod
    def apply1(s, a):
        return (default__.ApplyToCanvas(s, a, 9999))[0]

    @staticmethod
    def fold(s, actions):
        return ConfluxContract.default__.Fold(default__.apply1, s, actions)

    @staticmethod
    def C0():
        return CanvasState_CanvasState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "markdown")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opaque-provider")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), CanvasAvailability_Ready())

    @staticmethod
    def SameInstance(a, b):
        return ((a).channel) == ((b).channel)

    @staticmethod
    def ContentResourceRead(contentUri):
        return ResourceReadParams_ResourceReadParams(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahp-root://")), contentUri)

    @staticmethod
    def ValidateCanvasAdmission(supportNegotiated, handlerInstalled, authorized, channelInUse):
        if not(supportNegotiated):
            return CanvasAdmission_CanvasUnsupported()
        elif not(handlerInstalled):
            return CanvasAdmission_CanvasHandlerUnavailable()
        elif not(authorized):
            return CanvasAdmission_CanvasUnauthorized()
        elif channelInUse:
            return CanvasAdmission_CanvasUriCollision()
        elif True:
            return CanvasAdmission_CanvasSupported()

    @staticmethod
    def CanvasProviderFailure(code, message):
        return CanvasError_ProviderError(default__.CANVAS__PROVIDER__ERROR, CanvasProviderErrorData_CanvasProviderErrorData(code, message))

    @staticmethod
    def DegradeAvailability(s):
        d_0_dt__update__tmp_h0_ = s
        d_1_dt__update_havailability_h0_ = CanvasAvailability_Stale()
        return CanvasState_CanvasState((d_0_dt__update__tmp_h0_).canvasId, (d_0_dt__update__tmp_h0_).providerId, (d_0_dt__update__tmp_h0_).input_, (d_0_dt__update__tmp_h0_).title, (d_0_dt__update__tmp_h0_).activity, (d_0_dt__update__tmp_h0_).contentUri, d_1_dt__update_havailability_h0_)

    @staticmethod
    def ReconnectCanvas(snapshot, supportNegotiated):
        if supportNegotiated:
            return AhpSkeleton.Option_Some(snapshot)
        elif True:
            return AhpSkeleton.Option_None()

    @staticmethod
    def RunCorpus():
        pass_: int = int(0)
        total: int = int(0)
        total = 5
        pass_ = 0
        d_0_s_: CanvasState
        d_0_s_ = CanvasState_CanvasState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "markdown")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acme.editors")), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Draft"))), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idle"))), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahp-session:/2f9c/content/canvas-1"))), CanvasAvailability_Ready())
        d_1_all_: CanvasState
        d_1_all_ = default__.apply1(d_0_s_, CanvasAction_Updated(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Published"))), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "saved"))), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "https://example.com/docs/published.html"))), AhpSkeleton.Option_Some(CanvasAvailability_Stale())))
        def iife0_(_pat_let13_0):
            def iife1_(d_2_dt__update__tmp_h0_):
                def iife2_(_pat_let14_0):
                    def iife3_(d_3_dt__update_havailability_h0_):
                        def iife4_(_pat_let15_0):
                            def iife5_(d_4_dt__update_hcontentUri_h0_):
                                def iife6_(_pat_let16_0):
                                    def iife7_(d_5_dt__update_hactivity_h0_):
                                        def iife8_(_pat_let17_0):
                                            def iife9_(d_6_dt__update_htitle_h0_):
                                                return CanvasState_CanvasState((d_2_dt__update__tmp_h0_).canvasId, (d_2_dt__update__tmp_h0_).providerId, (d_2_dt__update__tmp_h0_).input_, d_6_dt__update_htitle_h0_, d_5_dt__update_hactivity_h0_, d_4_dt__update_hcontentUri_h0_, d_3_dt__update_havailability_h0_)
                                            return iife9_(_pat_let17_0)
                                        return iife8_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Published"))))
                                    return iife7_(_pat_let16_0)
                                return iife6_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "saved"))))
                            return iife5_(_pat_let15_0)
                        return iife4_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "https://example.com/docs/published.html"))))
                    return iife3_(_pat_let14_0)
                return iife2_(CanvasAvailability_Stale())
            return iife1_(_pat_let13_0)
        if (d_1_all_) == (iife0_(d_0_s_)):
            pass_ = (pass_) + (1)
        d_7_left_: CanvasState
        d_7_left_ = default__.apply1(d_0_s_, CanvasAction_Updated(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Renamed"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "https://example.com/docs/renamed.html"))), AhpSkeleton.Option_None()))
        def iife10_(_pat_let18_0):
            def iife11_(d_8_dt__update__tmp_h1_):
                def iife12_(_pat_let19_0):
                    def iife13_(d_9_dt__update_hcontentUri_h1_):
                        def iife14_(_pat_let20_0):
                            def iife15_(d_10_dt__update_htitle_h1_):
                                return CanvasState_CanvasState((d_8_dt__update__tmp_h1_).canvasId, (d_8_dt__update__tmp_h1_).providerId, (d_8_dt__update__tmp_h1_).input_, d_10_dt__update_htitle_h1_, (d_8_dt__update__tmp_h1_).activity, d_9_dt__update_hcontentUri_h1_, (d_8_dt__update__tmp_h1_).availability)
                            return iife15_(_pat_let20_0)
                        return iife14_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Renamed"))))
                    return iife13_(_pat_let19_0)
                return iife12_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "https://example.com/docs/renamed.html"))))
            return iife11_(_pat_let18_0)
        if (d_7_left_) == (iife10_(d_0_s_)):
            pass_ = (pass_) + (1)
        d_11_right_: CanvasState
        d_11_right_ = default__.apply1(d_0_s_, CanvasAction_Updated(AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "error"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(CanvasAvailability_Stale())))
        def iife16_(_pat_let21_0):
            def iife17_(d_12_dt__update__tmp_h2_):
                def iife18_(_pat_let22_0):
                    def iife19_(d_13_dt__update_havailability_h1_):
                        def iife20_(_pat_let23_0):
                            def iife21_(d_14_dt__update_hactivity_h1_):
                                return CanvasState_CanvasState((d_12_dt__update__tmp_h2_).canvasId, (d_12_dt__update__tmp_h2_).providerId, (d_12_dt__update__tmp_h2_).input_, (d_12_dt__update__tmp_h2_).title, d_14_dt__update_hactivity_h1_, (d_12_dt__update__tmp_h2_).contentUri, d_13_dt__update_havailability_h1_)
                            return iife21_(_pat_let23_0)
                        return iife20_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "error"))))
                    return iife19_(_pat_let22_0)
                return iife18_(CanvasAvailability_Stale())
            return iife17_(_pat_let21_0)
        if (d_11_right_) == (iife16_(d_0_s_)):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_0_s_, CanvasAction_CloseRequested())) == (d_0_s_):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_0_s_, CanvasAction_CanvasUnknown(ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "type")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "canvas/nonExistentAction")))}))))) == (d_0_s_):
            pass_ = (pass_) + (1)
        return pass_, total

    @_dafny.classproperty
    def CANVAS__PROVIDER__ERROR(instance):
        return -32012

class CanvasAvailability:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [CanvasAvailability_Ready(), CanvasAvailability_Stale()]
    @classmethod
    def default(cls, ):
        return lambda: CanvasAvailability_Ready()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Ready(self) -> bool:
        return isinstance(self, CanvasAvailability_Ready)
    @property
    def is_Stale(self) -> bool:
        return isinstance(self, CanvasAvailability_Stale)

class CanvasAvailability_Ready(CanvasAvailability, NamedTuple('Ready', [])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasAvailability.Ready'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasAvailability_Ready)
    def __hash__(self) -> int:
        return super().__hash__()

class CanvasAvailability_Stale(CanvasAvailability, NamedTuple('Stale', [])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasAvailability.Stale'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasAvailability_Stale)
    def __hash__(self) -> int:
        return super().__hash__()


class CanvasState:
    @classmethod
    def default(cls, ):
        return lambda: CanvasState_CanvasState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), CanvasAvailability.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_CanvasState(self) -> bool:
        return isinstance(self, CanvasState_CanvasState)

class CanvasState_CanvasState(CanvasState, NamedTuple('CanvasState', [('canvasId', Any), ('providerId', Any), ('input_', Any), ('title', Any), ('activity', Any), ('contentUri', Any), ('availability', Any)])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasState.CanvasState({self.canvasId.VerbatimString(True)}, {self.providerId.VerbatimString(True)}, {_dafny.string_of(self.input_)}, {_dafny.string_of(self.title)}, {_dafny.string_of(self.activity)}, {_dafny.string_of(self.contentUri)}, {_dafny.string_of(self.availability)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasState_CanvasState) and self.canvasId == __o.canvasId and self.providerId == __o.providerId and self.input_ == __o.input_ and self.title == __o.title and self.activity == __o.activity and self.contentUri == __o.contentUri and self.availability == __o.availability
    def __hash__(self) -> int:
        return super().__hash__()


class CanvasAction:
    @classmethod
    def default(cls, ):
        return lambda: CanvasAction_Updated(AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Updated(self) -> bool:
        return isinstance(self, CanvasAction_Updated)
    @property
    def is_CloseRequested(self) -> bool:
        return isinstance(self, CanvasAction_CloseRequested)
    @property
    def is_CanvasUnknown(self) -> bool:
        return isinstance(self, CanvasAction_CanvasUnknown)

class CanvasAction_Updated(CanvasAction, NamedTuple('Updated', [('title', Any), ('activity', Any), ('contentUri', Any), ('availability', Any)])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasAction.Updated({_dafny.string_of(self.title)}, {_dafny.string_of(self.activity)}, {_dafny.string_of(self.contentUri)}, {_dafny.string_of(self.availability)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasAction_Updated) and self.title == __o.title and self.activity == __o.activity and self.contentUri == __o.contentUri and self.availability == __o.availability
    def __hash__(self) -> int:
        return super().__hash__()

class CanvasAction_CloseRequested(CanvasAction, NamedTuple('CloseRequested', [])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasAction.CloseRequested'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasAction_CloseRequested)
    def __hash__(self) -> int:
        return super().__hash__()

class CanvasAction_CanvasUnknown(CanvasAction, NamedTuple('CanvasUnknown', [('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasAction.CanvasUnknown({_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasAction_CanvasUnknown) and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()


class CanvasInstance:
    @classmethod
    def default(cls, ):
        return lambda: CanvasInstance_CanvasInstance(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), CanvasState.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_CanvasInstance(self) -> bool:
        return isinstance(self, CanvasInstance_CanvasInstance)

class CanvasInstance_CanvasInstance(CanvasInstance, NamedTuple('CanvasInstance', [('channel', Any), ('snapshot', Any)])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasInstance.CanvasInstance({self.channel.VerbatimString(True)}, {_dafny.string_of(self.snapshot)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasInstance_CanvasInstance) and self.channel == __o.channel and self.snapshot == __o.snapshot
    def __hash__(self) -> int:
        return super().__hash__()


class ResourceReadParams:
    @classmethod
    def default(cls, ):
        return lambda: ResourceReadParams_ResourceReadParams(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ResourceReadParams(self) -> bool:
        return isinstance(self, ResourceReadParams_ResourceReadParams)

class ResourceReadParams_ResourceReadParams(ResourceReadParams, NamedTuple('ResourceReadParams', [('channel', Any), ('uri', Any)])):
    def __dafnystr__(self) -> str:
        return f'Canvas.ResourceReadParams.ResourceReadParams({self.channel.VerbatimString(True)}, {self.uri.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResourceReadParams_ResourceReadParams) and self.channel == __o.channel and self.uri == __o.uri
    def __hash__(self) -> int:
        return super().__hash__()


class CanvasAdmission:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [CanvasAdmission_CanvasSupported(), CanvasAdmission_CanvasUnsupported(), CanvasAdmission_CanvasHandlerUnavailable(), CanvasAdmission_CanvasUnauthorized(), CanvasAdmission_CanvasUriCollision()]
    @classmethod
    def default(cls, ):
        return lambda: CanvasAdmission_CanvasSupported()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_CanvasSupported(self) -> bool:
        return isinstance(self, CanvasAdmission_CanvasSupported)
    @property
    def is_CanvasUnsupported(self) -> bool:
        return isinstance(self, CanvasAdmission_CanvasUnsupported)
    @property
    def is_CanvasHandlerUnavailable(self) -> bool:
        return isinstance(self, CanvasAdmission_CanvasHandlerUnavailable)
    @property
    def is_CanvasUnauthorized(self) -> bool:
        return isinstance(self, CanvasAdmission_CanvasUnauthorized)
    @property
    def is_CanvasUriCollision(self) -> bool:
        return isinstance(self, CanvasAdmission_CanvasUriCollision)

class CanvasAdmission_CanvasSupported(CanvasAdmission, NamedTuple('CanvasSupported', [])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasAdmission.CanvasSupported'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasAdmission_CanvasSupported)
    def __hash__(self) -> int:
        return super().__hash__()

class CanvasAdmission_CanvasUnsupported(CanvasAdmission, NamedTuple('CanvasUnsupported', [])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasAdmission.CanvasUnsupported'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasAdmission_CanvasUnsupported)
    def __hash__(self) -> int:
        return super().__hash__()

class CanvasAdmission_CanvasHandlerUnavailable(CanvasAdmission, NamedTuple('CanvasHandlerUnavailable', [])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasAdmission.CanvasHandlerUnavailable'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasAdmission_CanvasHandlerUnavailable)
    def __hash__(self) -> int:
        return super().__hash__()

class CanvasAdmission_CanvasUnauthorized(CanvasAdmission, NamedTuple('CanvasUnauthorized', [])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasAdmission.CanvasUnauthorized'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasAdmission_CanvasUnauthorized)
    def __hash__(self) -> int:
        return super().__hash__()

class CanvasAdmission_CanvasUriCollision(CanvasAdmission, NamedTuple('CanvasUriCollision', [])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasAdmission.CanvasUriCollision'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasAdmission_CanvasUriCollision)
    def __hash__(self) -> int:
        return super().__hash__()


class CanvasProviderErrorData:
    @classmethod
    def default(cls, ):
        return lambda: CanvasProviderErrorData_CanvasProviderErrorData(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_CanvasProviderErrorData(self) -> bool:
        return isinstance(self, CanvasProviderErrorData_CanvasProviderErrorData)

class CanvasProviderErrorData_CanvasProviderErrorData(CanvasProviderErrorData, NamedTuple('CanvasProviderErrorData', [('code', Any), ('message', Any)])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasProviderErrorData.CanvasProviderErrorData({self.code.VerbatimString(True)}, {self.message.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasProviderErrorData_CanvasProviderErrorData) and self.code == __o.code and self.message == __o.message
    def __hash__(self) -> int:
        return super().__hash__()


class CanvasError:
    @classmethod
    def default(cls, ):
        return lambda: CanvasError_AdmissionError(CanvasAdmission.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_AdmissionError(self) -> bool:
        return isinstance(self, CanvasError_AdmissionError)
    @property
    def is_ProviderError(self) -> bool:
        return isinstance(self, CanvasError_ProviderError)

class CanvasError_AdmissionError(CanvasError, NamedTuple('AdmissionError', [('reason', Any)])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasError.AdmissionError({_dafny.string_of(self.reason)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasError_AdmissionError) and self.reason == __o.reason
    def __hash__(self) -> int:
        return super().__hash__()

class CanvasError_ProviderError(CanvasError, NamedTuple('ProviderError', [('errorCode', Any), ('data', Any)])):
    def __dafnystr__(self) -> str:
        return f'Canvas.CanvasError.ProviderError({_dafny.string_of(self.errorCode)}, {_dafny.string_of(self.data)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CanvasError_ProviderError) and self.errorCode == __o.errorCode and self.data == __o.data
    def __hash__(self) -> int:
        return super().__hash__()

