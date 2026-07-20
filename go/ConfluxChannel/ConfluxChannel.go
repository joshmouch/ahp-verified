// Package ConfluxChannel
// Dafny module ConfluxChannel compiled into Go

package ConfluxChannel

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-go/Changeset"
	m_Chat "github.com/joshmouch/ahp-go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-go/ClientMain"
	m_ConfluxBatchRoute "github.com/joshmouch/ahp-go/ConfluxBatchRoute"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxDedupe "github.com/joshmouch/ahp-go/ConfluxDedupe"
	m_ConfluxDelimited "github.com/joshmouch/ahp-go/ConfluxDelimited"
	m_ConfluxFilterKeys "github.com/joshmouch/ahp-go/ConfluxFilterKeys"
	m_ConfluxFixpoint "github.com/joshmouch/ahp-go/ConfluxFixpoint"
	m_ConfluxGenericCodec "github.com/joshmouch/ahp-go/ConfluxGenericCodec"
	m_ConfluxJoin "github.com/joshmouch/ahp-go/ConfluxJoin"
	m_ConfluxKeyedOps "github.com/joshmouch/ahp-go/ConfluxKeyedOps"
	m_ConfluxMapValues "github.com/joshmouch/ahp-go/ConfluxMapValues"
	m_ConfluxMerge "github.com/joshmouch/ahp-go/ConfluxMerge"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-go/ConfluxOrderedLog"
	m_ConfluxProduct "github.com/joshmouch/ahp-go/ConfluxProduct"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxResolve "github.com/joshmouch/ahp-go/ConfluxResolve"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-go/ConfluxResourceCeilings"
	m_ConfluxRoute "github.com/joshmouch/ahp-go/ConfluxRoute"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-go/ConfluxSeqRoute"
	m_ConfluxSeqRouteMerge "github.com/joshmouch/ahp-go/ConfluxSeqRouteMerge"
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
	return "ConfluxChannel.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Accept(r func(interface{}, interface{}) interface{}, h HostState, a interface{}) HostState {
	return Companion_HostState_.Create_HostState_((r)((h).Dtor_canonical(), a), ((h).Dtor_nextSeq()).Plus(_dafny.One))
}
func (_static *CompanionStruct_Default___) Receive(r func(interface{}, interface{}) interface{}, c ClientState, a interface{}) ClientState {
	return Companion_ClientState_.Create_ClientState_((r)((c).Dtor_mirror(), a), ((c).Dtor_lastSeq()).Plus(_dafny.One))
}
func (_static *CompanionStruct_Default___) AcceptFold(r func(interface{}, interface{}) interface{}, h HostState, actions _dafny.Sequence) HostState {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((actions).Cardinality())).Sign() == 0 {
		return h
	} else {
		var _in0 func(interface{}, interface{}) interface{} = r
		_ = _in0
		var _in1 HostState = Companion_Default___.Accept(func(coer88 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
			return func(arg101 interface{}, arg102 interface{}) interface{} {
				return coer88(arg101, arg102)
			}
		}(r), h, (actions).Select(0).(interface{}))
		_ = _in1
		var _in2 _dafny.Sequence = (actions).Drop(1)
		_ = _in2
		r = _in0
		h = _in1
		actions = _in2
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) ReceiveFold(r func(interface{}, interface{}) interface{}, c ClientState, actions _dafny.Sequence) ClientState {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((actions).Cardinality())).Sign() == 0 {
		return c
	} else {
		var _in0 func(interface{}, interface{}) interface{} = r
		_ = _in0
		var _in1 ClientState = Companion_Default___.Receive(func(coer89 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
			return func(arg103 interface{}, arg104 interface{}) interface{} {
				return coer89(arg103, arg104)
			}
		}(r), c, (actions).Select(0).(interface{}))
		_ = _in1
		var _in2 _dafny.Sequence = (actions).Drop(1)
		_ = _in2
		r = _in0
		c = _in1
		actions = _in2
		goto TAIL_CALL_START
	}
}

// End of class Default__

// Definition of datatype HostState
type HostState struct {
	Data_HostState_
}

func (_this HostState) Get_() Data_HostState_ {
	return _this.Data_HostState_
}

type Data_HostState_ interface {
	isHostState()
}

type CompanionStruct_HostState_ struct {
}

var Companion_HostState_ = CompanionStruct_HostState_{}

type HostState_HostState struct {
	Canonical interface{}
	NextSeq   _dafny.Int
}

func (HostState_HostState) isHostState() {}

func (CompanionStruct_HostState_) Create_HostState_(Canonical interface{}, NextSeq _dafny.Int) HostState {
	return HostState{HostState_HostState{Canonical, NextSeq}}
}

func (_this HostState) Is_HostState() bool {
	_, ok := _this.Get_().(HostState_HostState)
	return ok
}

func (CompanionStruct_HostState_) Default(_default_S interface{}) HostState {
	return Companion_HostState_.Create_HostState_(_default_S, _dafny.Zero)
}

func (_this HostState) Dtor_canonical() interface{} {
	return _this.Get_().(HostState_HostState).Canonical
}

func (_this HostState) Dtor_nextSeq() _dafny.Int {
	return _this.Get_().(HostState_HostState).NextSeq
}

func (_this HostState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case HostState_HostState:
		{
			return "ConfluxChannel.HostState.HostState" + "(" + _dafny.String(data.Canonical) + ", " + _dafny.String(data.NextSeq) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this HostState) Equals(other HostState) bool {
	switch data1 := _this.Get_().(type) {
	case HostState_HostState:
		{
			data2, ok := other.Get_().(HostState_HostState)
			return ok && _dafny.AreEqual(data1.Canonical, data2.Canonical) && data1.NextSeq.Cmp(data2.NextSeq) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this HostState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(HostState)
	return ok && _this.Equals(typed)
}

func Type_HostState_(Type_S_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_HostState_{Type_S_}
}

type type_HostState_ struct {
	Type_S_ _dafny.TypeDescriptor
}

func (_this type_HostState_) Default() interface{} {
	Type_S_ := _this.Type_S_
	_ = Type_S_
	return Companion_HostState_.Default(Type_S_.Default())
}

func (_this type_HostState_) String() string {
	return "ConfluxChannel.HostState"
}
func (_this HostState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = HostState{}

// End of datatype HostState

// Definition of datatype ClientState
type ClientState struct {
	Data_ClientState_
}

func (_this ClientState) Get_() Data_ClientState_ {
	return _this.Data_ClientState_
}

type Data_ClientState_ interface {
	isClientState()
}

type CompanionStruct_ClientState_ struct {
}

var Companion_ClientState_ = CompanionStruct_ClientState_{}

type ClientState_ClientState struct {
	Mirror  interface{}
	LastSeq _dafny.Int
}

func (ClientState_ClientState) isClientState() {}

func (CompanionStruct_ClientState_) Create_ClientState_(Mirror interface{}, LastSeq _dafny.Int) ClientState {
	return ClientState{ClientState_ClientState{Mirror, LastSeq}}
}

func (_this ClientState) Is_ClientState() bool {
	_, ok := _this.Get_().(ClientState_ClientState)
	return ok
}

func (CompanionStruct_ClientState_) Default(_default_S interface{}) ClientState {
	return Companion_ClientState_.Create_ClientState_(_default_S, _dafny.Zero)
}

func (_this ClientState) Dtor_mirror() interface{} {
	return _this.Get_().(ClientState_ClientState).Mirror
}

func (_this ClientState) Dtor_lastSeq() _dafny.Int {
	return _this.Get_().(ClientState_ClientState).LastSeq
}

func (_this ClientState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ClientState_ClientState:
		{
			return "ConfluxChannel.ClientState.ClientState" + "(" + _dafny.String(data.Mirror) + ", " + _dafny.String(data.LastSeq) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ClientState) Equals(other ClientState) bool {
	switch data1 := _this.Get_().(type) {
	case ClientState_ClientState:
		{
			data2, ok := other.Get_().(ClientState_ClientState)
			return ok && _dafny.AreEqual(data1.Mirror, data2.Mirror) && data1.LastSeq.Cmp(data2.LastSeq) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ClientState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ClientState)
	return ok && _this.Equals(typed)
}

func Type_ClientState_(Type_S_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_ClientState_{Type_S_}
}

type type_ClientState_ struct {
	Type_S_ _dafny.TypeDescriptor
}

func (_this type_ClientState_) Default() interface{} {
	Type_S_ := _this.Type_S_
	_ = Type_S_
	return Companion_ClientState_.Default(Type_S_.Default())
}

func (_this type_ClientState_) String() string {
	return "ConfluxChannel.ClientState"
}
func (_this ClientState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ClientState{}

// End of datatype ClientState
