// Package ConfluxHttpService
// Dafny module ConfluxHttpService compiled into Go

package ConfluxHttpService

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-go/Changeset"
	m_Chat "github.com/joshmouch/ahp-go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-go/ClientMain"
	m_ConfluxBatchRoute "github.com/joshmouch/ahp-go/ConfluxBatchRoute"
	m_ConfluxBudgetConvergence "github.com/joshmouch/ahp-go/ConfluxBudgetConvergence"
	m_ConfluxChannel "github.com/joshmouch/ahp-go/ConfluxChannel"
	m_ConfluxChannelManifest "github.com/joshmouch/ahp-go/ConfluxChannelManifest"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxCommand "github.com/joshmouch/ahp-go/ConfluxCommand"
	m_ConfluxCommandIdentity "github.com/joshmouch/ahp-go/ConfluxCommandIdentity"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxConvergence "github.com/joshmouch/ahp-go/ConfluxConvergence"
	m_ConfluxDecimalText "github.com/joshmouch/ahp-go/ConfluxDecimalText"
	m_ConfluxDedupe "github.com/joshmouch/ahp-go/ConfluxDedupe"
	m_ConfluxDelimited "github.com/joshmouch/ahp-go/ConfluxDelimited"
	m_ConfluxDurableHistory "github.com/joshmouch/ahp-go/ConfluxDurableHistory"
	m_ConfluxExtern "github.com/joshmouch/ahp-go/ConfluxExtern"
	m_ConfluxFilterKeys "github.com/joshmouch/ahp-go/ConfluxFilterKeys"
	m_ConfluxFixpoint "github.com/joshmouch/ahp-go/ConfluxFixpoint"
	m_ConfluxGenericCodec "github.com/joshmouch/ahp-go/ConfluxGenericCodec"
	m_ConfluxHttpCapability "github.com/joshmouch/ahp-go/ConfluxHttpCapability"
	m_ConfluxIoCapability "github.com/joshmouch/ahp-go/ConfluxIoCapability"
	m_ConfluxJoin "github.com/joshmouch/ahp-go/ConfluxJoin"
	m_ConfluxJsonRpc "github.com/joshmouch/ahp-go/ConfluxJsonRpc"
	m_ConfluxJsonText "github.com/joshmouch/ahp-go/ConfluxJsonText"
	m_ConfluxKeyedOps "github.com/joshmouch/ahp-go/ConfluxKeyedOps"
	m_ConfluxMapValues "github.com/joshmouch/ahp-go/ConfluxMapValues"
	m_ConfluxMerge "github.com/joshmouch/ahp-go/ConfluxMerge"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-go/ConfluxOrderedLog"
	m_ConfluxProduct "github.com/joshmouch/ahp-go/ConfluxProduct"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxResolve "github.com/joshmouch/ahp-go/ConfluxResolve"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-go/ConfluxResourceCeilings"
	m_ConfluxResourceSupervisor "github.com/joshmouch/ahp-go/ConfluxResourceSupervisor"
	m_ConfluxRoute "github.com/joshmouch/ahp-go/ConfluxRoute"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-go/ConfluxSeqRoute"
	m_ConfluxSeqRouteMerge "github.com/joshmouch/ahp-go/ConfluxSeqRouteMerge"
	m_ConfluxSequence "github.com/joshmouch/ahp-go/ConfluxSequence"
	m_ConfluxServiceLifecycle "github.com/joshmouch/ahp-go/ConfluxServiceLifecycle"
	m_ConfluxStateMachine "github.com/joshmouch/ahp-go/ConfluxStateMachine"
	m_ConfluxSupervisedOperationResult "github.com/joshmouch/ahp-go/ConfluxSupervisedOperationResult"
	m_ConfluxTree "github.com/joshmouch/ahp-go/ConfluxTree"
	m_ConfluxWatermark "github.com/joshmouch/ahp-go/ConfluxWatermark"
	m_ResourceWatch "github.com/joshmouch/ahp-go/ResourceWatch"
	m_Session "github.com/joshmouch/ahp-go/Session"
	m__System "github.com/joshmouch/ahp-go/System_"
	m_Terminal "github.com/joshmouch/ahp-go/Terminal"
	_dafny "github.com/joshmouch/ahp-go/dafny"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_ConfluxCodec.Dummy__
var _ m_ConfluxContract.Dummy__
var _ m_AhpSkeleton.Dummy__
var _ m_ResourceWatch.Dummy__
var _ m_Canvas.Dummy__
var _ m_ConfluxOperators.Dummy__
var _ m_ConfluxPrune.Dummy__
var _ m_ConfluxSeqRoute.Dummy__
var _ m_Changeset.Dummy__
var _ m_Annotations.Dummy__
var _ m_Terminal.Dummy__
var _ m_Session.Dummy__
var _ m_ConfluxOrderedLog.Dummy__
var _ m_Chat.Dummy__
var _ m_ClientMain.Dummy__
var _ m_ConfluxSemanticGraphIdentity.Dummy__
var _ m_ConfluxContentAddress.Dummy__
var _ m_ConfluxSemanticGraphContract.Dummy__
var _ m_ConfluxResourceCeilings.Dummy__
var _ m_ConfluxSupervisedOperationResult.Dummy__
var _ m_ConfluxTree.Dummy__
var _ m_ConfluxMerge.Dummy__
var _ m_ConfluxRoute.Dummy__
var _ m_ConfluxMapValues.Dummy__
var _ m_ConfluxDelimited.Dummy__
var _ m_ConfluxFixpoint.Dummy__
var _ m_ConfluxGenericCodec.Dummy__
var _ m_ConfluxResolve.Dummy__
var _ m_ConfluxWatermark.Dummy__
var _ m_ConfluxStateMachine.Dummy__
var _ m_ConfluxProduct.Dummy__
var _ m_ConfluxJoin.Dummy__
var _ m_ConfluxDedupe.Dummy__
var _ m_ConfluxBatchRoute.Dummy__
var _ m_ConfluxSeqRouteMerge.Dummy__
var _ m_ConfluxKeyedOps.Dummy__
var _ m_ConfluxFilterKeys.Dummy__
var _ m_ConfluxChannel.Dummy__
var _ m_ConfluxDurableHistory.Dummy__
var _ m_ConfluxCommand.Dummy__
var _ m_ConfluxChannelManifest.Dummy__
var _ m_ConfluxConvergence.Dummy__
var _ m_ConfluxBudgetConvergence.Dummy__
var _ m_ConfluxExtern.Dummy__
var _ m_ConfluxDecimalText.Dummy__
var _ m_ConfluxJsonText.Dummy__
var _ m_ConfluxCommandIdentity.Dummy__
var _ m_ConfluxResourceSupervisor.Dummy__
var _ m_ConfluxServiceLifecycle.Dummy__
var _ m_ConfluxSequence.Dummy__
var _ m_ConfluxJsonRpc.Dummy__

type Dummy__ struct{}

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "ConfluxHttpService.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Contains(text _dafny.Sequence, wanted _dafny.Sequence) bool {
	return m_ConfluxSequence.Companion_Default___.Contains(text, wanted)
}
func (_static *CompanionStruct_Default___) EndsWith(text _dafny.Sequence, suffix _dafny.Sequence) bool {
	return m_ConfluxSequence.Companion_Default___.EndsWith(text, suffix)
}
func (_static *CompanionStruct_Default___) PathOnly(path _dafny.Sequence) _dafny.Sequence {
	return (path).Take((m_ConfluxDelimited.Companion_Default___.IndexOf(_dafny.CodePoint('?'), path)).Uint32())
}
func (_static *CompanionStruct_Default___) MethodMatches(pattern MethodPattern, httpMethod _dafny.Sequence) bool {
	var _source0 MethodPattern = pattern
	_ = _source0
	{
		if _source0.Is_ExactMethod() {
			var _0_value _dafny.Sequence = _source0.Get_().(MethodPattern_ExactMethod).Value
			_ = _0_value
			return _dafny.Companion_Sequence_.Equal(_0_value, httpMethod)
		}
	}
	{
		return true
	}
}
func (_static *CompanionStruct_Default___) PathMatches(pattern PathPattern, path _dafny.Sequence) bool {
	var _source0 PathPattern = pattern
	_ = _source0
	{
		if _source0.Is_ExactPath() {
			var _0_value _dafny.Sequence = _source0.Get_().(PathPattern_ExactPath).Value
			_ = _0_value
			return _dafny.Companion_Sequence_.Equal(_0_value, path)
		}
	}
	{
		return true
	}
}
func (_static *CompanionStruct_Default___) HttpRouteStep(state HttpRouteFold, route HttpRoute, httpMethod _dafny.Sequence, path _dafny.Sequence) HttpRouteFold {
	var _source0 HttpRouteFold = state
	_ = _source0
	{
		if _source0.Is_RouteSelected() {
			var _0_handler interface{} = _source0.Get_().(HttpRouteFold_RouteSelected).Handler
			_ = _0_handler
			return state
		}
	}
	{
		if _source0.Is_RouteSearching() {
			if !(Companion_Default___.PathMatches((route).Dtor_pathPattern(), path)) {
				return state
			} else if Companion_Default___.MethodMatches((route).Dtor_methodPattern(), httpMethod) {
				return Companion_HttpRouteFold_.Create_RouteSelected_((route).Dtor_handler())
			} else {
				return Companion_HttpRouteFold_.Create_RoutePathSeen_()
			}
		}
	}
	{
		if (Companion_Default___.PathMatches((route).Dtor_pathPattern(), path)) && (Companion_Default___.MethodMatches((route).Dtor_methodPattern(), httpMethod)) {
			return Companion_HttpRouteFold_.Create_RouteSelected_((route).Dtor_handler())
		} else {
			return state
		}
	}
}
func (_static *CompanionStruct_Default___) FoldHttpRoutes(routes _dafny.Sequence, httpMethod _dafny.Sequence, path _dafny.Sequence, state HttpRouteFold) HttpRouteFold {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if ((_dafny.IntOfUint32((routes).Cardinality())).Sign() == 0) || ((state).Is_RouteSelected()) {
		return state
	} else {
		var _in0 _dafny.Sequence = (routes).Drop(1)
		_ = _in0
		var _in1 _dafny.Sequence = httpMethod
		_ = _in1
		var _in2 _dafny.Sequence = path
		_ = _in2
		var _in3 HttpRouteFold = Companion_Default___.HttpRouteStep(state, (routes).Select(0).(HttpRoute), httpMethod, path)
		_ = _in3
		routes = _in0
		httpMethod = _in1
		path = _in2
		state = _in3
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) ResolveHttpRoute(routes _dafny.Sequence, httpMethod _dafny.Sequence, rawPath _dafny.Sequence) HttpRouteResolution {
	var _0_path _dafny.Sequence = Companion_Default___.PathOnly(rawPath)
	_ = _0_path
	var _source0 HttpRouteFold = Companion_Default___.FoldHttpRoutes(routes, httpMethod, _0_path, Companion_HttpRouteFold_.Create_RouteSearching_())
	_ = _source0
	{
		if _source0.Is_RouteSelected() {
			var _1_handler interface{} = _source0.Get_().(HttpRouteFold_RouteSelected).Handler
			_ = _1_handler
			return Companion_HttpRouteResolution_.Create_RouteMatched_(_1_handler)
		}
	}
	{
		if _source0.Is_RoutePathSeen() {
			return Companion_HttpRouteResolution_.Create_RouteMethodNotAllowed_()
		}
	}
	{
		return Companion_HttpRouteResolution_.Create_RouteNotFound_()
	}
}
func (_static *CompanionStruct_Default___) RequestRouteStep(state RequestRouteFold, route RequestRoute, name _dafny.Sequence) RequestRouteFold {
	var _source0 RequestRouteFold = state
	_ = _source0
	{
		if _source0.Is_RequestSelected() {
			var _0_handler interface{} = _source0.Get_().(RequestRouteFold_RequestSelected).Handler
			_ = _0_handler
			return state
		}
	}
	{
		if _dafny.Companion_Sequence_.Equal((route).Dtor_name(), name) {
			return Companion_RequestRouteFold_.Create_RequestSelected_((route).Dtor_handler())
		} else {
			return state
		}
	}
}
func (_static *CompanionStruct_Default___) FoldRequestRoutes(routes _dafny.Sequence, name _dafny.Sequence, state RequestRouteFold) RequestRouteFold {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if ((_dafny.IntOfUint32((routes).Cardinality())).Sign() == 0) || ((state).Is_RequestSelected()) {
		return state
	} else {
		var _in0 _dafny.Sequence = (routes).Drop(1)
		_ = _in0
		var _in1 _dafny.Sequence = name
		_ = _in1
		var _in2 RequestRouteFold = Companion_Default___.RequestRouteStep(state, (routes).Select(0).(RequestRoute), name)
		_ = _in2
		routes = _in0
		name = _in1
		state = _in2
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) ResolveRequestRoute(routes _dafny.Sequence, name _dafny.Sequence) RequestRouteResolution {
	var _source0 RequestRouteFold = Companion_Default___.FoldRequestRoutes(routes, name, Companion_RequestRouteFold_.Create_RequestSearching_())
	_ = _source0
	{
		if _source0.Is_RequestSelected() {
			var _0_handler interface{} = _source0.Get_().(RequestRouteFold_RequestSelected).Handler
			_ = _0_handler
			return Companion_RequestRouteResolution_.Create_RequestMatched_(_0_handler)
		}
	}
	{
		return Companion_RequestRouteResolution_.Create_RequestNotFound_()
	}
}
func (_static *CompanionStruct_Default___) StaticName(path _dafny.Sequence) _dafny.Sequence {
	var _0_clean _dafny.Sequence = Companion_Default___.PathOnly(path)
	_ = _0_clean
	if _dafny.Companion_Sequence_.Equal(_0_clean, _dafny.UnicodeSeqOfUtf8Bytes("/")) {
		return _dafny.UnicodeSeqOfUtf8Bytes("index.html")
	} else if ((_dafny.IntOfUint32((_0_clean).Cardinality())).Sign() == 1) && (((_0_clean).Select(0).(_dafny.CodePoint)) == (_dafny.CodePoint('/'))) {
		return (_0_clean).Drop(1)
	} else {
		return _0_clean
	}
}
func (_static *CompanionStruct_Default___) SafeStaticName(name _dafny.Sequence) bool {
	return ((!_dafny.Companion_Sequence_.Equal(name, _dafny.UnicodeSeqOfUtf8Bytes(""))) && (!(Companion_Default___.Contains(name, _dafny.UnicodeSeqOfUtf8Bytes(".."))))) && (!(Companion_Default___.Contains(name, _dafny.UnicodeSeqOfUtf8Bytes("\\"))))
}
func (_static *CompanionStruct_Default___) Mime(name _dafny.Sequence) _dafny.Sequence {
	if Companion_Default___.EndsWith(name, _dafny.UnicodeSeqOfUtf8Bytes(".html")) {
		return _dafny.UnicodeSeqOfUtf8Bytes("text/html; charset=utf-8")
	} else if Companion_Default___.EndsWith(name, _dafny.UnicodeSeqOfUtf8Bytes(".js")) {
		return _dafny.UnicodeSeqOfUtf8Bytes("text/javascript; charset=utf-8")
	} else if Companion_Default___.EndsWith(name, _dafny.UnicodeSeqOfUtf8Bytes(".css")) {
		return _dafny.UnicodeSeqOfUtf8Bytes("text/css; charset=utf-8")
	} else if Companion_Default___.EndsWith(name, _dafny.UnicodeSeqOfUtf8Bytes(".json")) {
		return _dafny.UnicodeSeqOfUtf8Bytes("application/json; charset=utf-8")
	} else if Companion_Default___.EndsWith(name, _dafny.UnicodeSeqOfUtf8Bytes(".svg")) {
		return _dafny.UnicodeSeqOfUtf8Bytes("image/svg+xml")
	} else if Companion_Default___.EndsWith(name, _dafny.UnicodeSeqOfUtf8Bytes(".png")) {
		return _dafny.UnicodeSeqOfUtf8Bytes("image/png")
	} else {
		return _dafny.UnicodeSeqOfUtf8Bytes("application/octet-stream")
	}
}
func (_static *CompanionStruct_Default___) EmptyChanges() m_ConfluxCodec.Json {
	return m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("added"), m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf())).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("removed"), m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf())).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("changed"), m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf())))
}
func (_static *CompanionStruct_Default___) JsonArray(value m_ConfluxCodec.Json) _dafny.Sequence {
	if (value).Is_JArr() {
		return (value).Dtor_elems()
	} else {
		return _dafny.SeqOf()
	}
}
func (_static *CompanionStruct_Default___) StringFieldsFrom(items _dafny.Sequence, fieldName _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((items).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		var _1_value _dafny.Sequence = m_ConfluxCodec.Companion_Default___.AsStr(m_ConfluxCodec.Companion_Default___.Field((items).Select(0).(m_ConfluxCodec.Json), fieldName))
		_ = _1_value
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (func() _dafny.Sequence {
			if _dafny.Companion_Sequence_.Equal(_1_value, _dafny.UnicodeSeqOfUtf8Bytes("")) {
				return _dafny.SeqOf()
			}
			return _dafny.SeqOf(_1_value)
		})())
		var _in0 _dafny.Sequence = (items).Drop(1)
		_ = _in0
		var _in1 _dafny.Sequence = fieldName
		_ = _in1
		items = _in0
		fieldName = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) StringFields(value m_ConfluxCodec.Json, fieldName _dafny.Sequence) _dafny.Sequence {
	return Companion_Default___.StringFieldsFrom(Companion_Default___.JsonArray(value), fieldName)
}
func (_static *CompanionStruct_Default___) ParseRpc(body _dafny.Sequence) RpcIngress {
	var ingress RpcIngress = Companion_RpcIngress_.Default()
	_ = ingress
	var _0_parsed m_ConfluxCodec.Option
	_ = _0_parsed
	var _out0 m_ConfluxCodec.Option
	_ = _out0
	_out0 = m_ConfluxJsonRpc.Companion_Default___.Parse(body)
	_0_parsed = _out0
	if (_0_parsed).Is_None() {
		ingress = Companion_RpcIngress_.Create_RpcParseError_()
	} else {
		var _source0 m_ConfluxJsonRpc.Message = (_0_parsed).Dtor_value().(m_ConfluxJsonRpc.Message)
		_ = _source0
		{
			{
				if _source0.Is_Request() {
					var _1_id m_ConfluxCodec.Json = _source0.Get_().(m_ConfluxJsonRpc.Message_Request).Id
					_ = _1_id
					var _2_rpcMethod _dafny.Sequence = _source0.Get_().(m_ConfluxJsonRpc.Message_Request).RpcMethod
					_ = _2_rpcMethod
					var _3_params m_ConfluxCodec.Json = _source0.Get_().(m_ConfluxJsonRpc.Message_Request).Params
					_ = _3_params
					ingress = Companion_RpcIngress_.Create_RpcCall_(_1_id, _2_rpcMethod, _3_params)
					goto Lmatch0
				}
			}
			{
				if _source0.Is_Invalid() {
					ingress = Companion_RpcIngress_.Create_RpcInvalid_()
					goto Lmatch0
				}
			}
			{
				ingress = Companion_RpcIngress_.Create_RpcNonRequest_()
			}
			goto Lmatch0
		}
	Lmatch0:
	}
	return ingress
}
func (_static *CompanionStruct_Default___) JsonText(value m_ConfluxCodec.Json) _dafny.Sequence {
	var text _dafny.Sequence = _dafny.EmptySeq
	_ = text
	var _out0 _dafny.Sequence
	_ = _out0
	_out0 = m_ConfluxJsonText.Companion_Default___.Stringify(value)
	text = _out0
	return text
}
func (_static *CompanionStruct_Default___) RespondJson(request _dafny.Int, status _dafny.Int, value m_ConfluxCodec.Json) {
	var _0_text _dafny.Sequence
	_ = _0_text
	var _out0 _dafny.Sequence
	_ = _out0
	_out0 = Companion_Default___.JsonText(value)
	_0_text = _out0
	m_ConfluxHttpCapability.Companion_Default___.Respond(request, status, _dafny.UnicodeSeqOfUtf8Bytes("application/json; charset=utf-8"), Companion_Default___.CorsHeaders(), _0_text)
}
func (_static *CompanionStruct_Default___) RespondRpcError(request _dafny.Int, status _dafny.Int, id m_ConfluxCodec.Json, code _dafny.Int, message _dafny.Sequence) {
	Companion_Default___.RespondJson(request, status, m_ConfluxJsonRpc.Companion_Default___.FailureEnvelope(id, code, message, m_ConfluxCodec.Companion_Json_.Create_JNull_()))
}
func (_static *CompanionStruct_Default___) RespondRawRpcResult(request _dafny.Int, id m_ConfluxCodec.Json, resultPath _dafny.Sequence, unavailableMessage _dafny.Sequence) {
	var _0_present bool
	_ = _0_present
	var _1_content _dafny.Sequence
	_ = _1_content
	var _out0 bool
	_ = _out0
	var _out1 _dafny.Sequence
	_ = _out1
	_out0, _out1 = m_ConfluxIoCapability.Companion_Default___.ReadFile(resultPath)
	_0_present = _out0
	_1_content = _out1
	if _0_present {
		var _2_idText _dafny.Sequence
		_ = _2_idText
		var _out2 _dafny.Sequence
		_ = _out2
		_out2 = Companion_Default___.JsonText(id)
		_2_idText = _out2
		m_ConfluxHttpCapability.Companion_Default___.Respond(request, _dafny.IntOfInt64(200), _dafny.UnicodeSeqOfUtf8Bytes("application/json; charset=utf-8"), Companion_Default___.CorsHeaders(), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("{\"jsonrpc\":\"2.0\",\"id\":"), _2_idText), _dafny.UnicodeSeqOfUtf8Bytes(",\"result\":")), _1_content), _dafny.UnicodeSeqOfUtf8Bytes("}")))
	} else {
		Companion_Default___.RespondRpcError(request, _dafny.IntOfInt64(500), id, _dafny.IntOfInt64(-32603), unavailableMessage)
	}
}
func (_static *CompanionStruct_Default___) RespondFile(request _dafny.Int, path _dafny.Sequence, contentType _dafny.Sequence, headers _dafny.Sequence, missingStatus _dafny.Int, missingBody m_ConfluxCodec.Json) {
	var _0_present bool
	_ = _0_present
	var _1_content _dafny.Sequence
	_ = _1_content
	var _out0 bool
	_ = _out0
	var _out1 _dafny.Sequence
	_ = _out1
	_out0, _out1 = m_ConfluxIoCapability.Companion_Default___.ReadFile(path)
	_0_present = _out0
	_1_content = _out1
	if _0_present {
		m_ConfluxHttpCapability.Companion_Default___.Respond(request, _dafny.IntOfInt64(200), contentType, headers, _1_content)
	} else {
		Companion_Default___.RespondJson(request, missingStatus, missingBody)
	}
}
func (_static *CompanionStruct_Default___) RespondStatic(request _dafny.Int, root _dafny.Sequence, rawPath _dafny.Sequence) {
	var _0_name _dafny.Sequence
	_ = _0_name
	_0_name = Companion_Default___.StaticName(rawPath)
	if !(Companion_Default___.SafeStaticName(_0_name)) {
		Companion_Default___.RespondJson(request, _dafny.IntOfInt64(404), m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("error"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("not found")))))
	} else {
		Companion_Default___.RespondFile(request, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(root, _dafny.UnicodeSeqOfUtf8Bytes("/")), _0_name), Companion_Default___.Mime(_0_name), _dafny.SeqOf(), _dafny.IntOfInt64(404), m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("error"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("not found")))))
	}
}
func (_static *CompanionStruct_Default___) RespondCorsPreflight(request _dafny.Int) {
	m_ConfluxHttpCapability.Companion_Default___.Respond(request, _dafny.IntOfInt64(204), _dafny.UnicodeSeqOfUtf8Bytes("text/plain"), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("Access-Control-Allow-Origin: *"), _dafny.UnicodeSeqOfUtf8Bytes("Access-Control-Allow-Headers: content-type")), _dafny.UnicodeSeqOfUtf8Bytes(""))
}
func (_static *CompanionStruct_Default___) BeginEventStream(request _dafny.Int, eventName _dafny.Sequence, value m_ConfluxCodec.Json) {
	var _0_text _dafny.Sequence
	_ = _0_text
	var _out0 _dafny.Sequence
	_ = _out0
	_out0 = Companion_Default___.JsonText(value)
	_0_text = _out0
	m_ConfluxHttpCapability.Companion_Default___.BeginStream(request, _dafny.UnicodeSeqOfUtf8Bytes("text/event-stream"), Companion_Default___.EventHeaders(), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("event: "), eventName), _dafny.UnicodeSeqOfUtf8Bytes("\ndata: ")), _0_text), _dafny.UnicodeSeqOfUtf8Bytes("\n\n")))
}
func (_static *CompanionStruct_Default___) PublishEvent(server _dafny.Int, eventName _dafny.Sequence, value m_ConfluxCodec.Json) {
	var _0_text _dafny.Sequence
	_ = _0_text
	var _out0 _dafny.Sequence
	_ = _out0
	_out0 = Companion_Default___.JsonText(value)
	_0_text = _out0
	m_ConfluxHttpCapability.Companion_Default___.Publish(server, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("event: "), eventName), _dafny.UnicodeSeqOfUtf8Bytes("\ndata: ")), _0_text), _dafny.UnicodeSeqOfUtf8Bytes("\n\n")))
}
func (_static *CompanionStruct_Default___) RefreshJson(cwd _dafny.Sequence, command _dafny.Sequence, arguments _dafny.Sequence, metadataPath _dafny.Sequence, requiredField _dafny.Sequence) (bool, m_ConfluxCodec.Json, _dafny.Sequence) {
	var ok bool = false
	_ = ok
	var metadata m_ConfluxCodec.Json = m_ConfluxCodec.Companion_Json_.Default()
	_ = metadata
	var error_ _dafny.Sequence = _dafny.EmptySeq
	_ = error_
	var _0_code _dafny.Int
	_ = _0_code
	var _1_output _dafny.Sequence
	_ = _1_output
	var _2_processError _dafny.Sequence
	_ = _2_processError
	var _out0 _dafny.Int
	_ = _out0
	var _out1 _dafny.Sequence
	_ = _out1
	var _out2 _dafny.Sequence
	_ = _out2
	_out0, _out1, _out2 = m_ConfluxIoCapability.Companion_Default___.RunProcess(cwd, command, arguments)
	_0_code = _out0
	_1_output = _out1
	_2_processError = _out2
	if (_0_code).Sign() != 0 {
		ok = false
		metadata = m_ConfluxCodec.Companion_Json_.Create_JNull_()
		error_ = _2_processError
		return ok, metadata, error_
	}
	var _3_present bool
	_ = _3_present
	var _4_text _dafny.Sequence
	_ = _4_text
	var _out3 bool
	_ = _out3
	var _out4 _dafny.Sequence
	_ = _out4
	_out3, _out4 = m_ConfluxIoCapability.Companion_Default___.ReadFile(metadataPath)
	_3_present = _out3
	_4_text = _out4
	if !(_3_present) {
		ok = false
		metadata = m_ConfluxCodec.Companion_Json_.Create_JNull_()
		error_ = _dafny.UnicodeSeqOfUtf8Bytes("metadata unavailable")
		return ok, metadata, error_
	}
	var _5_parsed m_ConfluxJsonText.ParseResult
	_ = _5_parsed
	var _out5 m_ConfluxJsonText.ParseResult
	_ = _out5
	_out5 = m_ConfluxJsonText.Companion_Default___.ParseJsonChecked(_4_text)
	_5_parsed = _out5
	if ((_5_parsed).Is_Invalid()) || (_dafny.Companion_Sequence_.Equal(m_ConfluxCodec.Companion_Default___.AsStr(m_ConfluxCodec.Companion_Default___.Field((_5_parsed).Dtor_value(), requiredField)), _dafny.UnicodeSeqOfUtf8Bytes(""))) {
		ok = false
		metadata = m_ConfluxCodec.Companion_Json_.Create_JNull_()
		error_ = _dafny.UnicodeSeqOfUtf8Bytes("metadata invalid")
		return ok, metadata, error_
	}
	ok = true
	metadata = (_5_parsed).Dtor_value()
	error_ = _dafny.UnicodeSeqOfUtf8Bytes("")
	return ok, metadata, error_
}
func (_static *CompanionStruct_Default___) DrainWatch(watcher _dafny.Int, budget _dafny.Int, waitMs _dafny.Int) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (budget).Sign() == 1 {
		var _0_available bool
		_ = _0_available
		var _1_path _dafny.Sequence
		_ = _1_path
		var _out0 bool
		_ = _out0
		var _out1 _dafny.Sequence
		_ = _out1
		_out0, _out1 = m_ConfluxIoCapability.Companion_Default___.PollWatch(watcher, waitMs)
		_0_available = _out0
		_1_path = _out1
		if _0_available {
			var _in0 _dafny.Int = watcher
			_ = _in0
			var _in1 _dafny.Int = (budget).Minus(_dafny.One)
			_ = _in1
			var _in2 _dafny.Int = waitMs
			_ = _in2
			watcher = _in0
			budget = _in1
			waitMs = _in2
			goto TAIL_CALL_START
		}
	}
}
func (_static *CompanionStruct_Default___) CorsHeaders() _dafny.Sequence {
	return _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("Access-Control-Allow-Origin: *"))
}
func (_static *CompanionStruct_Default___) EventHeaders() _dafny.Sequence {
	return _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("Cache-Control: no-cache"), _dafny.UnicodeSeqOfUtf8Bytes("Access-Control-Allow-Origin: *"))
}

// End of class Default__

// Definition of datatype RpcIngress
type RpcIngress struct {
	Data_RpcIngress_
}

func (_this RpcIngress) Get_() Data_RpcIngress_ {
	return _this.Data_RpcIngress_
}

type Data_RpcIngress_ interface {
	isRpcIngress()
}

type CompanionStruct_RpcIngress_ struct {
}

var Companion_RpcIngress_ = CompanionStruct_RpcIngress_{}

type RpcIngress_RpcCall struct {
	Id        m_ConfluxCodec.Json
	RpcMethod _dafny.Sequence
	Params    m_ConfluxCodec.Json
}

func (RpcIngress_RpcCall) isRpcIngress() {}

func (CompanionStruct_RpcIngress_) Create_RpcCall_(Id m_ConfluxCodec.Json, RpcMethod _dafny.Sequence, Params m_ConfluxCodec.Json) RpcIngress {
	return RpcIngress{RpcIngress_RpcCall{Id, RpcMethod, Params}}
}

func (_this RpcIngress) Is_RpcCall() bool {
	_, ok := _this.Get_().(RpcIngress_RpcCall)
	return ok
}

type RpcIngress_RpcParseError struct {
}

func (RpcIngress_RpcParseError) isRpcIngress() {}

func (CompanionStruct_RpcIngress_) Create_RpcParseError_() RpcIngress {
	return RpcIngress{RpcIngress_RpcParseError{}}
}

func (_this RpcIngress) Is_RpcParseError() bool {
	_, ok := _this.Get_().(RpcIngress_RpcParseError)
	return ok
}

type RpcIngress_RpcInvalid struct {
}

func (RpcIngress_RpcInvalid) isRpcIngress() {}

func (CompanionStruct_RpcIngress_) Create_RpcInvalid_() RpcIngress {
	return RpcIngress{RpcIngress_RpcInvalid{}}
}

func (_this RpcIngress) Is_RpcInvalid() bool {
	_, ok := _this.Get_().(RpcIngress_RpcInvalid)
	return ok
}

type RpcIngress_RpcNonRequest struct {
}

func (RpcIngress_RpcNonRequest) isRpcIngress() {}

func (CompanionStruct_RpcIngress_) Create_RpcNonRequest_() RpcIngress {
	return RpcIngress{RpcIngress_RpcNonRequest{}}
}

func (_this RpcIngress) Is_RpcNonRequest() bool {
	_, ok := _this.Get_().(RpcIngress_RpcNonRequest)
	return ok
}

func (CompanionStruct_RpcIngress_) Default() RpcIngress {
	return Companion_RpcIngress_.Create_RpcCall_(m_ConfluxCodec.Companion_Json_.Default(), _dafny.EmptySeq, m_ConfluxCodec.Companion_Json_.Default())
}

func (_this RpcIngress) Dtor_id() m_ConfluxCodec.Json {
	return _this.Get_().(RpcIngress_RpcCall).Id
}

func (_this RpcIngress) Dtor_rpcMethod() _dafny.Sequence {
	return _this.Get_().(RpcIngress_RpcCall).RpcMethod
}

func (_this RpcIngress) Dtor_params() m_ConfluxCodec.Json {
	return _this.Get_().(RpcIngress_RpcCall).Params
}

func (_this RpcIngress) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RpcIngress_RpcCall:
		{
			return "ConfluxHttpService.RpcIngress.RpcCall" + "(" + _dafny.String(data.Id) + ", " + data.RpcMethod.VerbatimString(true) + ", " + _dafny.String(data.Params) + ")"
		}
	case RpcIngress_RpcParseError:
		{
			return "ConfluxHttpService.RpcIngress.RpcParseError"
		}
	case RpcIngress_RpcInvalid:
		{
			return "ConfluxHttpService.RpcIngress.RpcInvalid"
		}
	case RpcIngress_RpcNonRequest:
		{
			return "ConfluxHttpService.RpcIngress.RpcNonRequest"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RpcIngress) Equals(other RpcIngress) bool {
	switch data1 := _this.Get_().(type) {
	case RpcIngress_RpcCall:
		{
			data2, ok := other.Get_().(RpcIngress_RpcCall)
			return ok && data1.Id.Equals(data2.Id) && data1.RpcMethod.Equals(data2.RpcMethod) && data1.Params.Equals(data2.Params)
		}
	case RpcIngress_RpcParseError:
		{
			_, ok := other.Get_().(RpcIngress_RpcParseError)
			return ok
		}
	case RpcIngress_RpcInvalid:
		{
			_, ok := other.Get_().(RpcIngress_RpcInvalid)
			return ok
		}
	case RpcIngress_RpcNonRequest:
		{
			_, ok := other.Get_().(RpcIngress_RpcNonRequest)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RpcIngress) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RpcIngress)
	return ok && _this.Equals(typed)
}

func Type_RpcIngress_() _dafny.TypeDescriptor {
	return type_RpcIngress_{}
}

type type_RpcIngress_ struct {
}

func (_this type_RpcIngress_) Default() interface{} {
	return Companion_RpcIngress_.Default()
}

func (_this type_RpcIngress_) String() string {
	return "ConfluxHttpService.RpcIngress"
}
func (_this RpcIngress) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RpcIngress{}

// End of datatype RpcIngress

// Definition of datatype MethodPattern
type MethodPattern struct {
	Data_MethodPattern_
}

func (_this MethodPattern) Get_() Data_MethodPattern_ {
	return _this.Data_MethodPattern_
}

type Data_MethodPattern_ interface {
	isMethodPattern()
}

type CompanionStruct_MethodPattern_ struct {
}

var Companion_MethodPattern_ = CompanionStruct_MethodPattern_{}

type MethodPattern_ExactMethod struct {
	Value _dafny.Sequence
}

func (MethodPattern_ExactMethod) isMethodPattern() {}

func (CompanionStruct_MethodPattern_) Create_ExactMethod_(Value _dafny.Sequence) MethodPattern {
	return MethodPattern{MethodPattern_ExactMethod{Value}}
}

func (_this MethodPattern) Is_ExactMethod() bool {
	_, ok := _this.Get_().(MethodPattern_ExactMethod)
	return ok
}

type MethodPattern_AnyMethod struct {
}

func (MethodPattern_AnyMethod) isMethodPattern() {}

func (CompanionStruct_MethodPattern_) Create_AnyMethod_() MethodPattern {
	return MethodPattern{MethodPattern_AnyMethod{}}
}

func (_this MethodPattern) Is_AnyMethod() bool {
	_, ok := _this.Get_().(MethodPattern_AnyMethod)
	return ok
}

func (CompanionStruct_MethodPattern_) Default() MethodPattern {
	return Companion_MethodPattern_.Create_ExactMethod_(_dafny.EmptySeq)
}

func (_this MethodPattern) Dtor_value() _dafny.Sequence {
	return _this.Get_().(MethodPattern_ExactMethod).Value
}

func (_this MethodPattern) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case MethodPattern_ExactMethod:
		{
			return "ConfluxHttpService.MethodPattern.ExactMethod" + "(" + data.Value.VerbatimString(true) + ")"
		}
	case MethodPattern_AnyMethod:
		{
			return "ConfluxHttpService.MethodPattern.AnyMethod"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this MethodPattern) Equals(other MethodPattern) bool {
	switch data1 := _this.Get_().(type) {
	case MethodPattern_ExactMethod:
		{
			data2, ok := other.Get_().(MethodPattern_ExactMethod)
			return ok && data1.Value.Equals(data2.Value)
		}
	case MethodPattern_AnyMethod:
		{
			_, ok := other.Get_().(MethodPattern_AnyMethod)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this MethodPattern) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(MethodPattern)
	return ok && _this.Equals(typed)
}

func Type_MethodPattern_() _dafny.TypeDescriptor {
	return type_MethodPattern_{}
}

type type_MethodPattern_ struct {
}

func (_this type_MethodPattern_) Default() interface{} {
	return Companion_MethodPattern_.Default()
}

func (_this type_MethodPattern_) String() string {
	return "ConfluxHttpService.MethodPattern"
}
func (_this MethodPattern) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = MethodPattern{}

// End of datatype MethodPattern

// Definition of datatype PathPattern
type PathPattern struct {
	Data_PathPattern_
}

func (_this PathPattern) Get_() Data_PathPattern_ {
	return _this.Data_PathPattern_
}

type Data_PathPattern_ interface {
	isPathPattern()
}

type CompanionStruct_PathPattern_ struct {
}

var Companion_PathPattern_ = CompanionStruct_PathPattern_{}

type PathPattern_ExactPath struct {
	Value _dafny.Sequence
}

func (PathPattern_ExactPath) isPathPattern() {}

func (CompanionStruct_PathPattern_) Create_ExactPath_(Value _dafny.Sequence) PathPattern {
	return PathPattern{PathPattern_ExactPath{Value}}
}

func (_this PathPattern) Is_ExactPath() bool {
	_, ok := _this.Get_().(PathPattern_ExactPath)
	return ok
}

type PathPattern_AnyPath struct {
}

func (PathPattern_AnyPath) isPathPattern() {}

func (CompanionStruct_PathPattern_) Create_AnyPath_() PathPattern {
	return PathPattern{PathPattern_AnyPath{}}
}

func (_this PathPattern) Is_AnyPath() bool {
	_, ok := _this.Get_().(PathPattern_AnyPath)
	return ok
}

func (CompanionStruct_PathPattern_) Default() PathPattern {
	return Companion_PathPattern_.Create_ExactPath_(_dafny.EmptySeq)
}

func (_this PathPattern) Dtor_value() _dafny.Sequence {
	return _this.Get_().(PathPattern_ExactPath).Value
}

func (_this PathPattern) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case PathPattern_ExactPath:
		{
			return "ConfluxHttpService.PathPattern.ExactPath" + "(" + data.Value.VerbatimString(true) + ")"
		}
	case PathPattern_AnyPath:
		{
			return "ConfluxHttpService.PathPattern.AnyPath"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this PathPattern) Equals(other PathPattern) bool {
	switch data1 := _this.Get_().(type) {
	case PathPattern_ExactPath:
		{
			data2, ok := other.Get_().(PathPattern_ExactPath)
			return ok && data1.Value.Equals(data2.Value)
		}
	case PathPattern_AnyPath:
		{
			_, ok := other.Get_().(PathPattern_AnyPath)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this PathPattern) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(PathPattern)
	return ok && _this.Equals(typed)
}

func Type_PathPattern_() _dafny.TypeDescriptor {
	return type_PathPattern_{}
}

type type_PathPattern_ struct {
}

func (_this type_PathPattern_) Default() interface{} {
	return Companion_PathPattern_.Default()
}

func (_this type_PathPattern_) String() string {
	return "ConfluxHttpService.PathPattern"
}
func (_this PathPattern) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = PathPattern{}

// End of datatype PathPattern

// Definition of datatype HttpRoute
type HttpRoute struct {
	Data_HttpRoute_
}

func (_this HttpRoute) Get_() Data_HttpRoute_ {
	return _this.Data_HttpRoute_
}

type Data_HttpRoute_ interface {
	isHttpRoute()
}

type CompanionStruct_HttpRoute_ struct {
}

var Companion_HttpRoute_ = CompanionStruct_HttpRoute_{}

type HttpRoute_HttpRoute struct {
	MethodPattern MethodPattern
	PathPattern   PathPattern
	Handler       interface{}
}

func (HttpRoute_HttpRoute) isHttpRoute() {}

func (CompanionStruct_HttpRoute_) Create_HttpRoute_(MethodPattern MethodPattern, PathPattern PathPattern, Handler interface{}) HttpRoute {
	return HttpRoute{HttpRoute_HttpRoute{MethodPattern, PathPattern, Handler}}
}

func (_this HttpRoute) Is_HttpRoute() bool {
	_, ok := _this.Get_().(HttpRoute_HttpRoute)
	return ok
}

func (CompanionStruct_HttpRoute_) Default(_default_H interface{}) HttpRoute {
	return Companion_HttpRoute_.Create_HttpRoute_(Companion_MethodPattern_.Default(), Companion_PathPattern_.Default(), _default_H)
}

func (_this HttpRoute) Dtor_methodPattern() MethodPattern {
	return _this.Get_().(HttpRoute_HttpRoute).MethodPattern
}

func (_this HttpRoute) Dtor_pathPattern() PathPattern {
	return _this.Get_().(HttpRoute_HttpRoute).PathPattern
}

func (_this HttpRoute) Dtor_handler() interface{} {
	return _this.Get_().(HttpRoute_HttpRoute).Handler
}

func (_this HttpRoute) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case HttpRoute_HttpRoute:
		{
			return "ConfluxHttpService.HttpRoute.HttpRoute" + "(" + _dafny.String(data.MethodPattern) + ", " + _dafny.String(data.PathPattern) + ", " + _dafny.String(data.Handler) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this HttpRoute) Equals(other HttpRoute) bool {
	switch data1 := _this.Get_().(type) {
	case HttpRoute_HttpRoute:
		{
			data2, ok := other.Get_().(HttpRoute_HttpRoute)
			return ok && data1.MethodPattern.Equals(data2.MethodPattern) && data1.PathPattern.Equals(data2.PathPattern) && _dafny.AreEqual(data1.Handler, data2.Handler)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this HttpRoute) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(HttpRoute)
	return ok && _this.Equals(typed)
}

func Type_HttpRoute_(Type_H_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_HttpRoute_{Type_H_}
}

type type_HttpRoute_ struct {
	Type_H_ _dafny.TypeDescriptor
}

func (_this type_HttpRoute_) Default() interface{} {
	Type_H_ := _this.Type_H_
	_ = Type_H_
	return Companion_HttpRoute_.Default(Type_H_.Default())
}

func (_this type_HttpRoute_) String() string {
	return "ConfluxHttpService.HttpRoute"
}
func (_this HttpRoute) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = HttpRoute{}

// End of datatype HttpRoute

// Definition of datatype HttpRouteFold
type HttpRouteFold struct {
	Data_HttpRouteFold_
}

func (_this HttpRouteFold) Get_() Data_HttpRouteFold_ {
	return _this.Data_HttpRouteFold_
}

type Data_HttpRouteFold_ interface {
	isHttpRouteFold()
}

type CompanionStruct_HttpRouteFold_ struct {
}

var Companion_HttpRouteFold_ = CompanionStruct_HttpRouteFold_{}

type HttpRouteFold_RouteSearching struct {
}

func (HttpRouteFold_RouteSearching) isHttpRouteFold() {}

func (CompanionStruct_HttpRouteFold_) Create_RouteSearching_() HttpRouteFold {
	return HttpRouteFold{HttpRouteFold_RouteSearching{}}
}

func (_this HttpRouteFold) Is_RouteSearching() bool {
	_, ok := _this.Get_().(HttpRouteFold_RouteSearching)
	return ok
}

type HttpRouteFold_RoutePathSeen struct {
}

func (HttpRouteFold_RoutePathSeen) isHttpRouteFold() {}

func (CompanionStruct_HttpRouteFold_) Create_RoutePathSeen_() HttpRouteFold {
	return HttpRouteFold{HttpRouteFold_RoutePathSeen{}}
}

func (_this HttpRouteFold) Is_RoutePathSeen() bool {
	_, ok := _this.Get_().(HttpRouteFold_RoutePathSeen)
	return ok
}

type HttpRouteFold_RouteSelected struct {
	Handler interface{}
}

func (HttpRouteFold_RouteSelected) isHttpRouteFold() {}

func (CompanionStruct_HttpRouteFold_) Create_RouteSelected_(Handler interface{}) HttpRouteFold {
	return HttpRouteFold{HttpRouteFold_RouteSelected{Handler}}
}

func (_this HttpRouteFold) Is_RouteSelected() bool {
	_, ok := _this.Get_().(HttpRouteFold_RouteSelected)
	return ok
}

func (CompanionStruct_HttpRouteFold_) Default() HttpRouteFold {
	return Companion_HttpRouteFold_.Create_RouteSearching_()
}

func (_this HttpRouteFold) Dtor_handler() interface{} {
	return _this.Get_().(HttpRouteFold_RouteSelected).Handler
}

func (_this HttpRouteFold) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case HttpRouteFold_RouteSearching:
		{
			return "ConfluxHttpService.HttpRouteFold.RouteSearching"
		}
	case HttpRouteFold_RoutePathSeen:
		{
			return "ConfluxHttpService.HttpRouteFold.RoutePathSeen"
		}
	case HttpRouteFold_RouteSelected:
		{
			return "ConfluxHttpService.HttpRouteFold.RouteSelected" + "(" + _dafny.String(data.Handler) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this HttpRouteFold) Equals(other HttpRouteFold) bool {
	switch data1 := _this.Get_().(type) {
	case HttpRouteFold_RouteSearching:
		{
			_, ok := other.Get_().(HttpRouteFold_RouteSearching)
			return ok
		}
	case HttpRouteFold_RoutePathSeen:
		{
			_, ok := other.Get_().(HttpRouteFold_RoutePathSeen)
			return ok
		}
	case HttpRouteFold_RouteSelected:
		{
			data2, ok := other.Get_().(HttpRouteFold_RouteSelected)
			return ok && _dafny.AreEqual(data1.Handler, data2.Handler)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this HttpRouteFold) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(HttpRouteFold)
	return ok && _this.Equals(typed)
}

func Type_HttpRouteFold_() _dafny.TypeDescriptor {
	return type_HttpRouteFold_{}
}

type type_HttpRouteFold_ struct {
}

func (_this type_HttpRouteFold_) Default() interface{} {
	return Companion_HttpRouteFold_.Default()
}

func (_this type_HttpRouteFold_) String() string {
	return "ConfluxHttpService.HttpRouteFold"
}
func (_this HttpRouteFold) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = HttpRouteFold{}

// End of datatype HttpRouteFold

// Definition of datatype HttpRouteResolution
type HttpRouteResolution struct {
	Data_HttpRouteResolution_
}

func (_this HttpRouteResolution) Get_() Data_HttpRouteResolution_ {
	return _this.Data_HttpRouteResolution_
}

type Data_HttpRouteResolution_ interface {
	isHttpRouteResolution()
}

type CompanionStruct_HttpRouteResolution_ struct {
}

var Companion_HttpRouteResolution_ = CompanionStruct_HttpRouteResolution_{}

type HttpRouteResolution_RouteMatched struct {
	Handler interface{}
}

func (HttpRouteResolution_RouteMatched) isHttpRouteResolution() {}

func (CompanionStruct_HttpRouteResolution_) Create_RouteMatched_(Handler interface{}) HttpRouteResolution {
	return HttpRouteResolution{HttpRouteResolution_RouteMatched{Handler}}
}

func (_this HttpRouteResolution) Is_RouteMatched() bool {
	_, ok := _this.Get_().(HttpRouteResolution_RouteMatched)
	return ok
}

type HttpRouteResolution_RouteMethodNotAllowed struct {
}

func (HttpRouteResolution_RouteMethodNotAllowed) isHttpRouteResolution() {}

func (CompanionStruct_HttpRouteResolution_) Create_RouteMethodNotAllowed_() HttpRouteResolution {
	return HttpRouteResolution{HttpRouteResolution_RouteMethodNotAllowed{}}
}

func (_this HttpRouteResolution) Is_RouteMethodNotAllowed() bool {
	_, ok := _this.Get_().(HttpRouteResolution_RouteMethodNotAllowed)
	return ok
}

type HttpRouteResolution_RouteNotFound struct {
}

func (HttpRouteResolution_RouteNotFound) isHttpRouteResolution() {}

func (CompanionStruct_HttpRouteResolution_) Create_RouteNotFound_() HttpRouteResolution {
	return HttpRouteResolution{HttpRouteResolution_RouteNotFound{}}
}

func (_this HttpRouteResolution) Is_RouteNotFound() bool {
	_, ok := _this.Get_().(HttpRouteResolution_RouteNotFound)
	return ok
}

func (CompanionStruct_HttpRouteResolution_) Default() HttpRouteResolution {
	return Companion_HttpRouteResolution_.Create_RouteMethodNotAllowed_()
}

func (_this HttpRouteResolution) Dtor_handler() interface{} {
	return _this.Get_().(HttpRouteResolution_RouteMatched).Handler
}

func (_this HttpRouteResolution) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case HttpRouteResolution_RouteMatched:
		{
			return "ConfluxHttpService.HttpRouteResolution.RouteMatched" + "(" + _dafny.String(data.Handler) + ")"
		}
	case HttpRouteResolution_RouteMethodNotAllowed:
		{
			return "ConfluxHttpService.HttpRouteResolution.RouteMethodNotAllowed"
		}
	case HttpRouteResolution_RouteNotFound:
		{
			return "ConfluxHttpService.HttpRouteResolution.RouteNotFound"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this HttpRouteResolution) Equals(other HttpRouteResolution) bool {
	switch data1 := _this.Get_().(type) {
	case HttpRouteResolution_RouteMatched:
		{
			data2, ok := other.Get_().(HttpRouteResolution_RouteMatched)
			return ok && _dafny.AreEqual(data1.Handler, data2.Handler)
		}
	case HttpRouteResolution_RouteMethodNotAllowed:
		{
			_, ok := other.Get_().(HttpRouteResolution_RouteMethodNotAllowed)
			return ok
		}
	case HttpRouteResolution_RouteNotFound:
		{
			_, ok := other.Get_().(HttpRouteResolution_RouteNotFound)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this HttpRouteResolution) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(HttpRouteResolution)
	return ok && _this.Equals(typed)
}

func Type_HttpRouteResolution_() _dafny.TypeDescriptor {
	return type_HttpRouteResolution_{}
}

type type_HttpRouteResolution_ struct {
}

func (_this type_HttpRouteResolution_) Default() interface{} {
	return Companion_HttpRouteResolution_.Default()
}

func (_this type_HttpRouteResolution_) String() string {
	return "ConfluxHttpService.HttpRouteResolution"
}
func (_this HttpRouteResolution) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = HttpRouteResolution{}

// End of datatype HttpRouteResolution

// Definition of datatype RequestRoute
type RequestRoute struct {
	Data_RequestRoute_
}

func (_this RequestRoute) Get_() Data_RequestRoute_ {
	return _this.Data_RequestRoute_
}

type Data_RequestRoute_ interface {
	isRequestRoute()
}

type CompanionStruct_RequestRoute_ struct {
}

var Companion_RequestRoute_ = CompanionStruct_RequestRoute_{}

type RequestRoute_RequestRoute struct {
	Name    _dafny.Sequence
	Handler interface{}
}

func (RequestRoute_RequestRoute) isRequestRoute() {}

func (CompanionStruct_RequestRoute_) Create_RequestRoute_(Name _dafny.Sequence, Handler interface{}) RequestRoute {
	return RequestRoute{RequestRoute_RequestRoute{Name, Handler}}
}

func (_this RequestRoute) Is_RequestRoute() bool {
	_, ok := _this.Get_().(RequestRoute_RequestRoute)
	return ok
}

func (CompanionStruct_RequestRoute_) Default(_default_H interface{}) RequestRoute {
	return Companion_RequestRoute_.Create_RequestRoute_(_dafny.EmptySeq, _default_H)
}

func (_this RequestRoute) Dtor_name() _dafny.Sequence {
	return _this.Get_().(RequestRoute_RequestRoute).Name
}

func (_this RequestRoute) Dtor_handler() interface{} {
	return _this.Get_().(RequestRoute_RequestRoute).Handler
}

func (_this RequestRoute) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RequestRoute_RequestRoute:
		{
			return "ConfluxHttpService.RequestRoute.RequestRoute" + "(" + data.Name.VerbatimString(true) + ", " + _dafny.String(data.Handler) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RequestRoute) Equals(other RequestRoute) bool {
	switch data1 := _this.Get_().(type) {
	case RequestRoute_RequestRoute:
		{
			data2, ok := other.Get_().(RequestRoute_RequestRoute)
			return ok && data1.Name.Equals(data2.Name) && _dafny.AreEqual(data1.Handler, data2.Handler)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RequestRoute) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RequestRoute)
	return ok && _this.Equals(typed)
}

func Type_RequestRoute_(Type_H_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_RequestRoute_{Type_H_}
}

type type_RequestRoute_ struct {
	Type_H_ _dafny.TypeDescriptor
}

func (_this type_RequestRoute_) Default() interface{} {
	Type_H_ := _this.Type_H_
	_ = Type_H_
	return Companion_RequestRoute_.Default(Type_H_.Default())
}

func (_this type_RequestRoute_) String() string {
	return "ConfluxHttpService.RequestRoute"
}
func (_this RequestRoute) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RequestRoute{}

// End of datatype RequestRoute

// Definition of datatype RequestRouteFold
type RequestRouteFold struct {
	Data_RequestRouteFold_
}

func (_this RequestRouteFold) Get_() Data_RequestRouteFold_ {
	return _this.Data_RequestRouteFold_
}

type Data_RequestRouteFold_ interface {
	isRequestRouteFold()
}

type CompanionStruct_RequestRouteFold_ struct {
}

var Companion_RequestRouteFold_ = CompanionStruct_RequestRouteFold_{}

type RequestRouteFold_RequestSearching struct {
}

func (RequestRouteFold_RequestSearching) isRequestRouteFold() {}

func (CompanionStruct_RequestRouteFold_) Create_RequestSearching_() RequestRouteFold {
	return RequestRouteFold{RequestRouteFold_RequestSearching{}}
}

func (_this RequestRouteFold) Is_RequestSearching() bool {
	_, ok := _this.Get_().(RequestRouteFold_RequestSearching)
	return ok
}

type RequestRouteFold_RequestSelected struct {
	Handler interface{}
}

func (RequestRouteFold_RequestSelected) isRequestRouteFold() {}

func (CompanionStruct_RequestRouteFold_) Create_RequestSelected_(Handler interface{}) RequestRouteFold {
	return RequestRouteFold{RequestRouteFold_RequestSelected{Handler}}
}

func (_this RequestRouteFold) Is_RequestSelected() bool {
	_, ok := _this.Get_().(RequestRouteFold_RequestSelected)
	return ok
}

func (CompanionStruct_RequestRouteFold_) Default() RequestRouteFold {
	return Companion_RequestRouteFold_.Create_RequestSearching_()
}

func (_this RequestRouteFold) Dtor_handler() interface{} {
	return _this.Get_().(RequestRouteFold_RequestSelected).Handler
}

func (_this RequestRouteFold) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RequestRouteFold_RequestSearching:
		{
			return "ConfluxHttpService.RequestRouteFold.RequestSearching"
		}
	case RequestRouteFold_RequestSelected:
		{
			return "ConfluxHttpService.RequestRouteFold.RequestSelected" + "(" + _dafny.String(data.Handler) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RequestRouteFold) Equals(other RequestRouteFold) bool {
	switch data1 := _this.Get_().(type) {
	case RequestRouteFold_RequestSearching:
		{
			_, ok := other.Get_().(RequestRouteFold_RequestSearching)
			return ok
		}
	case RequestRouteFold_RequestSelected:
		{
			data2, ok := other.Get_().(RequestRouteFold_RequestSelected)
			return ok && _dafny.AreEqual(data1.Handler, data2.Handler)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RequestRouteFold) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RequestRouteFold)
	return ok && _this.Equals(typed)
}

func Type_RequestRouteFold_() _dafny.TypeDescriptor {
	return type_RequestRouteFold_{}
}

type type_RequestRouteFold_ struct {
}

func (_this type_RequestRouteFold_) Default() interface{} {
	return Companion_RequestRouteFold_.Default()
}

func (_this type_RequestRouteFold_) String() string {
	return "ConfluxHttpService.RequestRouteFold"
}
func (_this RequestRouteFold) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RequestRouteFold{}

// End of datatype RequestRouteFold

// Definition of datatype RequestRouteResolution
type RequestRouteResolution struct {
	Data_RequestRouteResolution_
}

func (_this RequestRouteResolution) Get_() Data_RequestRouteResolution_ {
	return _this.Data_RequestRouteResolution_
}

type Data_RequestRouteResolution_ interface {
	isRequestRouteResolution()
}

type CompanionStruct_RequestRouteResolution_ struct {
}

var Companion_RequestRouteResolution_ = CompanionStruct_RequestRouteResolution_{}

type RequestRouteResolution_RequestMatched struct {
	Handler interface{}
}

func (RequestRouteResolution_RequestMatched) isRequestRouteResolution() {}

func (CompanionStruct_RequestRouteResolution_) Create_RequestMatched_(Handler interface{}) RequestRouteResolution {
	return RequestRouteResolution{RequestRouteResolution_RequestMatched{Handler}}
}

func (_this RequestRouteResolution) Is_RequestMatched() bool {
	_, ok := _this.Get_().(RequestRouteResolution_RequestMatched)
	return ok
}

type RequestRouteResolution_RequestNotFound struct {
}

func (RequestRouteResolution_RequestNotFound) isRequestRouteResolution() {}

func (CompanionStruct_RequestRouteResolution_) Create_RequestNotFound_() RequestRouteResolution {
	return RequestRouteResolution{RequestRouteResolution_RequestNotFound{}}
}

func (_this RequestRouteResolution) Is_RequestNotFound() bool {
	_, ok := _this.Get_().(RequestRouteResolution_RequestNotFound)
	return ok
}

func (CompanionStruct_RequestRouteResolution_) Default() RequestRouteResolution {
	return Companion_RequestRouteResolution_.Create_RequestNotFound_()
}

func (_this RequestRouteResolution) Dtor_handler() interface{} {
	return _this.Get_().(RequestRouteResolution_RequestMatched).Handler
}

func (_this RequestRouteResolution) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RequestRouteResolution_RequestMatched:
		{
			return "ConfluxHttpService.RequestRouteResolution.RequestMatched" + "(" + _dafny.String(data.Handler) + ")"
		}
	case RequestRouteResolution_RequestNotFound:
		{
			return "ConfluxHttpService.RequestRouteResolution.RequestNotFound"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RequestRouteResolution) Equals(other RequestRouteResolution) bool {
	switch data1 := _this.Get_().(type) {
	case RequestRouteResolution_RequestMatched:
		{
			data2, ok := other.Get_().(RequestRouteResolution_RequestMatched)
			return ok && _dafny.AreEqual(data1.Handler, data2.Handler)
		}
	case RequestRouteResolution_RequestNotFound:
		{
			_, ok := other.Get_().(RequestRouteResolution_RequestNotFound)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RequestRouteResolution) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RequestRouteResolution)
	return ok && _this.Equals(typed)
}

func Type_RequestRouteResolution_() _dafny.TypeDescriptor {
	return type_RequestRouteResolution_{}
}

type type_RequestRouteResolution_ struct {
}

func (_this type_RequestRouteResolution_) Default() interface{} {
	return Companion_RequestRouteResolution_.Default()
}

func (_this type_RequestRouteResolution_) String() string {
	return "ConfluxHttpService.RequestRouteResolution"
}
func (_this RequestRouteResolution) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RequestRouteResolution{}

// End of datatype RequestRouteResolution
