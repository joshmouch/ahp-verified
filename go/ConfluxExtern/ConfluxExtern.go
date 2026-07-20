// Package ConfluxExtern
// Dafny module ConfluxExtern compiled into Go

package ConfluxExtern

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
	m_ConfluxContentAddress "github.com/joshmouch/ahp-go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxConvergence "github.com/joshmouch/ahp-go/ConfluxConvergence"
	m_ConfluxDedupe "github.com/joshmouch/ahp-go/ConfluxDedupe"
	m_ConfluxDelimited "github.com/joshmouch/ahp-go/ConfluxDelimited"
	m_ConfluxDurableHistory "github.com/joshmouch/ahp-go/ConfluxDurableHistory"
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
var _ m_ConfluxChannel.Dummy__
var _ m_ConfluxDurableHistory.Dummy__
var _ m_ConfluxCommand.Dummy__
var _ m_ConfluxChannelManifest.Dummy__
var _ m_ConfluxConvergence.Dummy__
var _ m_ConfluxBudgetConvergence.Dummy__

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
	return "ConfluxExtern.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) PullToOption(p Pull) m_ConfluxContract.Option {
	var _source0 Pull = p
	_ = _source0
	{
		if _source0.Is_Yield() {
			var _0_a interface{} = _source0.Get_().(Pull_Yield).Action
			_ = _0_a
			return m_ConfluxContract.Companion_Option_.Create_Some_(_0_a)
		}
	}
	{
		return m_ConfluxContract.Companion_Option_.Create_None_()
	}
}
func (_static *CompanionStruct_Default___) PushToOption(p Push) m_ConfluxContract.Option {
	var _source0 Push = p
	_ = _source0
	{
		if _source0.Is_Emit() {
			var _0_e interface{} = _source0.Get_().(Push_Emit).Effect
			_ = _0_e
			return m_ConfluxContract.Companion_Option_.Create_Some_(_0_e)
		}
	}
	{
		return m_ConfluxContract.Companion_Option_.Create_None_()
	}
}
func (_static *CompanionStruct_Default___) RunSink(sink func(interface{}) Push, xs _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		var _source0 Push = (sink)((xs).Select(0).(interface{}))
		_ = _source0
		{
			if _source0.Is_Emit() {
				var _1_e interface{} = _source0.Get_().(Push_Emit).Effect
				_ = _1_e
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1_e), Companion_Default___.RunSink(sink, (xs).Drop(1)))
			}
		}
		{
			return _dafny.SeqOf()
		}
	}
}

// End of class Default__

// Definition of datatype Pull
type Pull struct {
	Data_Pull_
}

func (_this Pull) Get_() Data_Pull_ {
	return _this.Data_Pull_
}

type Data_Pull_ interface {
	isPull()
}

type CompanionStruct_Pull_ struct {
}

var Companion_Pull_ = CompanionStruct_Pull_{}

type Pull_Yield struct {
	Action interface{}
}

func (Pull_Yield) isPull() {}

func (CompanionStruct_Pull_) Create_Yield_(Action interface{}) Pull {
	return Pull{Pull_Yield{Action}}
}

func (_this Pull) Is_Yield() bool {
	_, ok := _this.Get_().(Pull_Yield)
	return ok
}

type Pull_Eof struct {
}

func (Pull_Eof) isPull() {}

func (CompanionStruct_Pull_) Create_Eof_() Pull {
	return Pull{Pull_Eof{}}
}

func (_this Pull) Is_Eof() bool {
	_, ok := _this.Get_().(Pull_Eof)
	return ok
}

func (CompanionStruct_Pull_) Default() Pull {
	return Companion_Pull_.Create_Eof_()
}

func (_this Pull) Dtor_action() interface{} {
	return _this.Get_().(Pull_Yield).Action
}

func (_this Pull) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Pull_Yield:
		{
			return "ConfluxExtern.Pull.Yield" + "(" + _dafny.String(data.Action) + ")"
		}
	case Pull_Eof:
		{
			return "ConfluxExtern.Pull.Eof"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Pull) Equals(other Pull) bool {
	switch data1 := _this.Get_().(type) {
	case Pull_Yield:
		{
			data2, ok := other.Get_().(Pull_Yield)
			return ok && _dafny.AreEqual(data1.Action, data2.Action)
		}
	case Pull_Eof:
		{
			_, ok := other.Get_().(Pull_Eof)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Pull) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Pull)
	return ok && _this.Equals(typed)
}

func Type_Pull_() _dafny.TypeDescriptor {
	return type_Pull_{}
}

type type_Pull_ struct {
}

func (_this type_Pull_) Default() interface{} {
	return Companion_Pull_.Default()
}

func (_this type_Pull_) String() string {
	return "ConfluxExtern.Pull"
}
func (_this Pull) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Pull{}

// End of datatype Pull

// Definition of datatype Push
type Push struct {
	Data_Push_
}

func (_this Push) Get_() Data_Push_ {
	return _this.Data_Push_
}

type Data_Push_ interface {
	isPush()
}

type CompanionStruct_Push_ struct {
}

var Companion_Push_ = CompanionStruct_Push_{}

type Push_Emit struct {
	Effect interface{}
}

func (Push_Emit) isPush() {}

func (CompanionStruct_Push_) Create_Emit_(Effect interface{}) Push {
	return Push{Push_Emit{Effect}}
}

func (_this Push) Is_Emit() bool {
	_, ok := _this.Get_().(Push_Emit)
	return ok
}

type Push_Closed struct {
}

func (Push_Closed) isPush() {}

func (CompanionStruct_Push_) Create_Closed_() Push {
	return Push{Push_Closed{}}
}

func (_this Push) Is_Closed() bool {
	_, ok := _this.Get_().(Push_Closed)
	return ok
}

func (CompanionStruct_Push_) Default() Push {
	return Companion_Push_.Create_Closed_()
}

func (_this Push) Dtor_effect() interface{} {
	return _this.Get_().(Push_Emit).Effect
}

func (_this Push) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Push_Emit:
		{
			return "ConfluxExtern.Push.Emit" + "(" + _dafny.String(data.Effect) + ")"
		}
	case Push_Closed:
		{
			return "ConfluxExtern.Push.Closed"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Push) Equals(other Push) bool {
	switch data1 := _this.Get_().(type) {
	case Push_Emit:
		{
			data2, ok := other.Get_().(Push_Emit)
			return ok && _dafny.AreEqual(data1.Effect, data2.Effect)
		}
	case Push_Closed:
		{
			_, ok := other.Get_().(Push_Closed)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Push) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Push)
	return ok && _this.Equals(typed)
}

func Type_Push_() _dafny.TypeDescriptor {
	return type_Push_{}
}

type type_Push_ struct {
}

func (_this type_Push_) Default() interface{} {
	return Companion_Push_.Default()
}

func (_this type_Push_) String() string {
	return "ConfluxExtern.Push"
}
func (_this Push) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Push{}

// End of datatype Push
