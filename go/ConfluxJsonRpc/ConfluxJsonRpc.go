// Package ConfluxJsonRpc
// Dafny module ConfluxJsonRpc compiled into Go

package ConfluxJsonRpc

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
	m_ConfluxJoin "github.com/joshmouch/ahp-go/ConfluxJoin"
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
	return "ConfluxJsonRpc.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Object(fields _dafny.Map) m_ConfluxCodec.Json {
	return m_ConfluxCodec.Companion_Json_.Create_JObj_(fields)
}
func (_static *CompanionStruct_Default___) RequestEnvelope(id m_ConfluxCodec.Json, rpcMethod _dafny.Sequence, params m_ConfluxCodec.Json) m_ConfluxCodec.Json {
	return Companion_Default___.Object(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("jsonrpc"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("2.0"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("id"), id).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("method"), m_ConfluxCodec.Companion_Json_.Create_JStr_(rpcMethod)).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("params"), params))
}
func (_static *CompanionStruct_Default___) NotificationEnvelope(rpcMethod _dafny.Sequence, params m_ConfluxCodec.Json) m_ConfluxCodec.Json {
	return Companion_Default___.Object(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("jsonrpc"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("2.0"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("method"), m_ConfluxCodec.Companion_Json_.Create_JStr_(rpcMethod)).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("params"), params))
}
func (_static *CompanionStruct_Default___) SuccessEnvelope(id m_ConfluxCodec.Json, result m_ConfluxCodec.Json) m_ConfluxCodec.Json {
	return Companion_Default___.Object(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("jsonrpc"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("2.0"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("id"), id).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("result"), result))
}
func (_static *CompanionStruct_Default___) FailureEnvelope(id m_ConfluxCodec.Json, code _dafny.Int, message _dafny.Sequence, data m_ConfluxCodec.Json) m_ConfluxCodec.Json {
	return Companion_Default___.Object(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("jsonrpc"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("2.0"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("id"), id).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("error"), Companion_Default___.Object(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("code"), m_ConfluxCodec.Companion_Json_.Create_JNum_(code)).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("message"), m_ConfluxCodec.Companion_Json_.Create_JStr_(message)).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("data"), data))))
}
func (_static *CompanionStruct_Default___) NullResult(id m_ConfluxCodec.Json) m_ConfluxCodec.Json {
	return Companion_Default___.SuccessEnvelope(id, m_ConfluxCodec.Companion_Json_.Create_JNull_())
}
func (_static *CompanionStruct_Default___) Classify(j m_ConfluxCodec.Json) Message {
	var _0_id m_ConfluxCodec.Json = m_ConfluxCodec.Companion_Default___.Field(j, _dafny.UnicodeSeqOfUtf8Bytes("id"))
	_ = _0_id
	var _1_rpcMethod _dafny.Sequence = m_ConfluxCodec.Companion_Default___.AsStr(m_ConfluxCodec.Companion_Default___.Field(j, _dafny.UnicodeSeqOfUtf8Bytes("method")))
	_ = _1_rpcMethod
	var _2_params m_ConfluxCodec.Json = m_ConfluxCodec.Companion_Default___.Field(j, _dafny.UnicodeSeqOfUtf8Bytes("params"))
	_ = _2_params
	var _3_error m_ConfluxCodec.Json = m_ConfluxCodec.Companion_Default___.Field(j, _dafny.UnicodeSeqOfUtf8Bytes("error"))
	_ = _3_error
	if (!_dafny.Companion_Sequence_.Equal(_1_rpcMethod, _dafny.UnicodeSeqOfUtf8Bytes(""))) && (!((_0_id).Is_JNull())) {
		return Companion_Message_.Create_Request_(_0_id, _1_rpcMethod, _2_params)
	} else if !_dafny.Companion_Sequence_.Equal(_1_rpcMethod, _dafny.UnicodeSeqOfUtf8Bytes("")) {
		return Companion_Message_.Create_Notification_(_1_rpcMethod, _2_params)
	} else if (!((_0_id).Is_JNull())) && (!((_3_error).Is_JNull())) {
		return Companion_Message_.Create_Failure_(_0_id, _3_error, m_ConfluxCodec.Companion_Default___.AsStr(m_ConfluxCodec.Companion_Default___.Field(_3_error, _dafny.UnicodeSeqOfUtf8Bytes("message"))))
	} else if !((_0_id).Is_JNull()) {
		return Companion_Message_.Create_Success_(_0_id, m_ConfluxCodec.Companion_Default___.Field(j, _dafny.UnicodeSeqOfUtf8Bytes("result")))
	} else {
		return Companion_Message_.Create_Invalid_()
	}
}
func (_static *CompanionStruct_Default___) Parse(text _dafny.Sequence) m_ConfluxCodec.Option {
	var message m_ConfluxCodec.Option = m_ConfluxCodec.Companion_Option_.Default()
	_ = message
	var _0_parsed m_ConfluxJsonText.ParseResult
	_ = _0_parsed
	var _out0 m_ConfluxJsonText.ParseResult
	_ = _out0
	_out0 = m_ConfluxJsonText.Companion_Default___.ParseJsonChecked(text)
	_0_parsed = _out0
	var _source0 m_ConfluxJsonText.ParseResult = _0_parsed
	_ = _source0
	{
		{
			if _source0.Is_Parsed() {
				var _1_j m_ConfluxCodec.Json = _source0.Get_().(m_ConfluxJsonText.ParseResult_Parsed).Value
				_ = _1_j
				message = m_ConfluxCodec.Companion_Option_.Create_Some_(Companion_Default___.Classify(_1_j))
				goto Lmatch0
			}
		}
		{
			message = m_ConfluxCodec.Companion_Option_.Create_None_()
		}
		goto Lmatch0
	}
Lmatch0:
	return message
}
func (_static *CompanionStruct_Default___) Stringify(j m_ConfluxCodec.Json) _dafny.Sequence {
	var text _dafny.Sequence = _dafny.EmptySeq
	_ = text
	var _out0 _dafny.Sequence
	_ = _out0
	_out0 = m_ConfluxJsonText.Companion_Default___.Stringify(j)
	text = _out0
	return text
}

// End of class Default__

// Definition of datatype Message
type Message struct {
	Data_Message_
}

func (_this Message) Get_() Data_Message_ {
	return _this.Data_Message_
}

type Data_Message_ interface {
	isMessage()
}

type CompanionStruct_Message_ struct {
}

var Companion_Message_ = CompanionStruct_Message_{}

type Message_Success struct {
	Id     m_ConfluxCodec.Json
	Result m_ConfluxCodec.Json
}

func (Message_Success) isMessage() {}

func (CompanionStruct_Message_) Create_Success_(Id m_ConfluxCodec.Json, Result m_ConfluxCodec.Json) Message {
	return Message{Message_Success{Id, Result}}
}

func (_this Message) Is_Success() bool {
	_, ok := _this.Get_().(Message_Success)
	return ok
}

type Message_Failure struct {
	Id        m_ConfluxCodec.Json
	Error     m_ConfluxCodec.Json
	ErrorText _dafny.Sequence
}

func (Message_Failure) isMessage() {}

func (CompanionStruct_Message_) Create_Failure_(Id m_ConfluxCodec.Json, Error m_ConfluxCodec.Json, ErrorText _dafny.Sequence) Message {
	return Message{Message_Failure{Id, Error, ErrorText}}
}

func (_this Message) Is_Failure() bool {
	_, ok := _this.Get_().(Message_Failure)
	return ok
}

type Message_Notification struct {
	RpcMethod _dafny.Sequence
	Params    m_ConfluxCodec.Json
}

func (Message_Notification) isMessage() {}

func (CompanionStruct_Message_) Create_Notification_(RpcMethod _dafny.Sequence, Params m_ConfluxCodec.Json) Message {
	return Message{Message_Notification{RpcMethod, Params}}
}

func (_this Message) Is_Notification() bool {
	_, ok := _this.Get_().(Message_Notification)
	return ok
}

type Message_Request struct {
	Id        m_ConfluxCodec.Json
	RpcMethod _dafny.Sequence
	Params    m_ConfluxCodec.Json
}

func (Message_Request) isMessage() {}

func (CompanionStruct_Message_) Create_Request_(Id m_ConfluxCodec.Json, RpcMethod _dafny.Sequence, Params m_ConfluxCodec.Json) Message {
	return Message{Message_Request{Id, RpcMethod, Params}}
}

func (_this Message) Is_Request() bool {
	_, ok := _this.Get_().(Message_Request)
	return ok
}

type Message_Invalid struct {
}

func (Message_Invalid) isMessage() {}

func (CompanionStruct_Message_) Create_Invalid_() Message {
	return Message{Message_Invalid{}}
}

func (_this Message) Is_Invalid() bool {
	_, ok := _this.Get_().(Message_Invalid)
	return ok
}

func (CompanionStruct_Message_) Default() Message {
	return Companion_Message_.Create_Success_(m_ConfluxCodec.Companion_Json_.Default(), m_ConfluxCodec.Companion_Json_.Default())
}

func (_this Message) Dtor_id() m_ConfluxCodec.Json {
	switch data := _this.Get_().(type) {
	case Message_Success:
		return data.Id
	case Message_Failure:
		return data.Id
	default:
		return data.(Message_Request).Id
	}
}

func (_this Message) Dtor_result() m_ConfluxCodec.Json {
	return _this.Get_().(Message_Success).Result
}

func (_this Message) Dtor_error() m_ConfluxCodec.Json {
	return _this.Get_().(Message_Failure).Error
}

func (_this Message) Dtor_errorText() _dafny.Sequence {
	return _this.Get_().(Message_Failure).ErrorText
}

func (_this Message) Dtor_rpcMethod() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case Message_Notification:
		return data.RpcMethod
	default:
		return data.(Message_Request).RpcMethod
	}
}

func (_this Message) Dtor_params() m_ConfluxCodec.Json {
	switch data := _this.Get_().(type) {
	case Message_Notification:
		return data.Params
	default:
		return data.(Message_Request).Params
	}
}

func (_this Message) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Message_Success:
		{
			return "ConfluxJsonRpc.Message.Success" + "(" + _dafny.String(data.Id) + ", " + _dafny.String(data.Result) + ")"
		}
	case Message_Failure:
		{
			return "ConfluxJsonRpc.Message.Failure" + "(" + _dafny.String(data.Id) + ", " + _dafny.String(data.Error) + ", " + data.ErrorText.VerbatimString(true) + ")"
		}
	case Message_Notification:
		{
			return "ConfluxJsonRpc.Message.Notification" + "(" + data.RpcMethod.VerbatimString(true) + ", " + _dafny.String(data.Params) + ")"
		}
	case Message_Request:
		{
			return "ConfluxJsonRpc.Message.Request" + "(" + _dafny.String(data.Id) + ", " + data.RpcMethod.VerbatimString(true) + ", " + _dafny.String(data.Params) + ")"
		}
	case Message_Invalid:
		{
			return "ConfluxJsonRpc.Message.Invalid"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Message) Equals(other Message) bool {
	switch data1 := _this.Get_().(type) {
	case Message_Success:
		{
			data2, ok := other.Get_().(Message_Success)
			return ok && data1.Id.Equals(data2.Id) && data1.Result.Equals(data2.Result)
		}
	case Message_Failure:
		{
			data2, ok := other.Get_().(Message_Failure)
			return ok && data1.Id.Equals(data2.Id) && data1.Error.Equals(data2.Error) && data1.ErrorText.Equals(data2.ErrorText)
		}
	case Message_Notification:
		{
			data2, ok := other.Get_().(Message_Notification)
			return ok && data1.RpcMethod.Equals(data2.RpcMethod) && data1.Params.Equals(data2.Params)
		}
	case Message_Request:
		{
			data2, ok := other.Get_().(Message_Request)
			return ok && data1.Id.Equals(data2.Id) && data1.RpcMethod.Equals(data2.RpcMethod) && data1.Params.Equals(data2.Params)
		}
	case Message_Invalid:
		{
			_, ok := other.Get_().(Message_Invalid)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Message) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Message)
	return ok && _this.Equals(typed)
}

func Type_Message_() _dafny.TypeDescriptor {
	return type_Message_{}
}

type type_Message_ struct {
}

func (_this type_Message_) Default() interface{} {
	return Companion_Message_.Default()
}

func (_this type_Message_) String() string {
	return "ConfluxJsonRpc.Message"
}
func (_this Message) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Message{}

// End of datatype Message
