// Package ConfluxChannelManifest
// Dafny module ConfluxChannelManifest compiled into Go

package ConfluxChannelManifest

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-go/Changeset"
	m_Chat "github.com/joshmouch/ahp-go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-go/ClientMain"
	m_ConfluxBatchRoute "github.com/joshmouch/ahp-go/ConfluxBatchRoute"
	m_ConfluxChannel "github.com/joshmouch/ahp-go/ConfluxChannel"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxCommand "github.com/joshmouch/ahp-go/ConfluxCommand"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
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
	return "ConfluxChannelManifest.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Canonical(m ChannelManifest, actions _dafny.Sequence) interface{} {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer106 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
		return func(arg123 interface{}, arg124 interface{}) interface{} {
			return coer106(arg123, arg124)
		}
	}(func(coer107 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
		return func(arg125 interface{}, arg126 interface{}) interface{} {
			return coer107(arg125, arg126)
		}
	}((m).Dtor_apply())), (m).Dtor_initial(), actions)
}
func (_static *CompanionStruct_Default___) View(m ChannelManifest, actions _dafny.Sequence) interface{} {
	return (func(coer108 func(interface{}) interface{}) func(interface{}) interface{} {
		return func(arg127 interface{}) interface{} {
			return coer108(arg127)
		}
	}((m).Dtor_readModel()))(Companion_Default___.Canonical(m, actions))
}

// End of class Default__

// Definition of datatype ChannelManifest
type ChannelManifest struct {
	Data_ChannelManifest_
}

func (_this ChannelManifest) Get_() Data_ChannelManifest_ {
	return _this.Data_ChannelManifest_
}

type Data_ChannelManifest_ interface {
	isChannelManifest()
}

type CompanionStruct_ChannelManifest_ struct {
}

var Companion_ChannelManifest_ = CompanionStruct_ChannelManifest_{}

type ChannelManifest_ChannelManifest struct {
	Initial   interface{}
	Apply     func(interface{}, interface{}) interface{}
	Decode    func(interface{}) m_ConfluxContract.Option
	ReadModel func(interface{}) interface{}
}

func (ChannelManifest_ChannelManifest) isChannelManifest() {}

func (CompanionStruct_ChannelManifest_) Create_ChannelManifest_(Initial interface{}, Apply func(interface{}, interface{}) interface{}, Decode func(interface{}) m_ConfluxContract.Option, ReadModel func(interface{}) interface{}) ChannelManifest {
	return ChannelManifest{ChannelManifest_ChannelManifest{Initial, Apply, Decode, ReadModel}}
}

func (_this ChannelManifest) Is_ChannelManifest() bool {
	_, ok := _this.Get_().(ChannelManifest_ChannelManifest)
	return ok
}

func (CompanionStruct_ChannelManifest_) Default(_default_S interface{}, _default_W interface{}) ChannelManifest {
	return Companion_ChannelManifest_.Create_ChannelManifest_(_default_S, func(interface{}, interface{}) interface{} { return _default_S }, func(interface{}) m_ConfluxContract.Option { return m_ConfluxContract.Companion_Option_.Default() }, func(interface{}) interface{} { return _default_W })
}

func (_this ChannelManifest) Dtor_initial() interface{} {
	return _this.Get_().(ChannelManifest_ChannelManifest).Initial
}

func (_this ChannelManifest) Dtor_apply() func(interface{}, interface{}) interface{} {
	return _this.Get_().(ChannelManifest_ChannelManifest).Apply
}

func (_this ChannelManifest) Dtor_decode() func(interface{}) m_ConfluxContract.Option {
	return _this.Get_().(ChannelManifest_ChannelManifest).Decode
}

func (_this ChannelManifest) Dtor_readModel() func(interface{}) interface{} {
	return _this.Get_().(ChannelManifest_ChannelManifest).ReadModel
}

func (_this ChannelManifest) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ChannelManifest_ChannelManifest:
		{
			return "ConfluxChannelManifest.ChannelManifest.ChannelManifest" + "(" + _dafny.String(data.Initial) + ", " + _dafny.String(data.Apply) + ", " + _dafny.String(data.Decode) + ", " + _dafny.String(data.ReadModel) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ChannelManifest) Equals(other ChannelManifest) bool {
	switch data1 := _this.Get_().(type) {
	case ChannelManifest_ChannelManifest:
		{
			data2, ok := other.Get_().(ChannelManifest_ChannelManifest)
			return ok && _dafny.AreEqual(data1.Initial, data2.Initial) && _dafny.AreEqual(data1.Apply, data2.Apply) && _dafny.AreEqual(data1.Decode, data2.Decode) && _dafny.AreEqual(data1.ReadModel, data2.ReadModel)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ChannelManifest) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ChannelManifest)
	return ok && _this.Equals(typed)
}

func Type_ChannelManifest_(Type_S_ _dafny.TypeDescriptor, Type_W_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_ChannelManifest_{Type_S_, Type_W_}
}

type type_ChannelManifest_ struct {
	Type_S_ _dafny.TypeDescriptor
	Type_W_ _dafny.TypeDescriptor
}

func (_this type_ChannelManifest_) Default() interface{} {
	Type_S_ := _this.Type_S_
	_ = Type_S_
	Type_W_ := _this.Type_W_
	_ = Type_W_
	return Companion_ChannelManifest_.Default(Type_S_.Default(), Type_W_.Default())
}

func (_this type_ChannelManifest_) String() string {
	return "ConfluxChannelManifest.ChannelManifest"
}
func (_this ChannelManifest) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ChannelManifest{}

// End of datatype ChannelManifest
