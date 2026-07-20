// Package ConfluxTree
// Dafny module ConfluxTree compiled into Go

package ConfluxTree

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-verified/go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-verified/go/Changeset"
	m_Chat "github.com/joshmouch/ahp-verified/go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-verified/go/ClientMain"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-verified/go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-verified/go/ConfluxOrderedLog"
	m_ConfluxPrune "github.com/joshmouch/ahp-verified/go/ConfluxPrune"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-verified/go/ConfluxResourceCeilings"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-verified/go/ConfluxSeqRoute"
	m_ConfluxSupervisedOperationResult "github.com/joshmouch/ahp-verified/go/ConfluxSupervisedOperationResult"
	m_ResourceWatch "github.com/joshmouch/ahp-verified/go/ResourceWatch"
	m_Session "github.com/joshmouch/ahp-verified/go/Session"
	m__System "github.com/joshmouch/ahp-verified/go/System_"
	m_Terminal "github.com/joshmouch/ahp-verified/go/Terminal"
	_dafny "github.com/joshmouch/ahp-verified/go/dafny"
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
	return "ConfluxTree.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) RoseMap(f func(interface{}) interface{}, t Rose) Rose {
	return Companion_Rose_.Create_Rose_((f)((t).Dtor_value()), Companion_Default___.MapChildren(func(coer74 func(interface{}) interface{}) func(interface{}) interface{} {
		return func(arg83 interface{}) interface{} {
			return coer74(arg83)
		}
	}(f), (t).Dtor_children()))
}
func (_static *CompanionStruct_Default___) MapChildren(f func(interface{}) interface{}, ts _dafny.Sequence) _dafny.Sequence {
	if (_dafny.IntOfUint32((ts).Cardinality())).Sign() == 0 {
		return _dafny.SeqOf()
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.RoseMap(func(coer75 func(interface{}) interface{}) func(interface{}) interface{} {
			return func(arg84 interface{}) interface{} {
				return coer75(arg84)
			}
		}(f), (ts).Select(0).(Rose))), Companion_Default___.MapChildren(f, (ts).Drop(1)))
	}
}
func (_static *CompanionStruct_Default___) Dbl(x _dafny.Int) _dafny.Int {
	return (x).Times(_dafny.IntOfInt64(2))
}

// End of class Default__

// Definition of datatype Rose
type Rose struct {
	Data_Rose_
}

func (_this Rose) Get_() Data_Rose_ {
	return _this.Data_Rose_
}

type Data_Rose_ interface {
	isRose()
}

type CompanionStruct_Rose_ struct {
}

var Companion_Rose_ = CompanionStruct_Rose_{}

type Rose_Rose struct {
	Value    interface{}
	Children _dafny.Sequence
}

func (Rose_Rose) isRose() {}

func (CompanionStruct_Rose_) Create_Rose_(Value interface{}, Children _dafny.Sequence) Rose {
	return Rose{Rose_Rose{Value, Children}}
}

func (_this Rose) Is_Rose() bool {
	_, ok := _this.Get_().(Rose_Rose)
	return ok
}

func (CompanionStruct_Rose_) Default(_default_T interface{}) Rose {
	return Companion_Rose_.Create_Rose_(_default_T, _dafny.EmptySeq)
}

func (_this Rose) Dtor_value() interface{} {
	return _this.Get_().(Rose_Rose).Value
}

func (_this Rose) Dtor_children() _dafny.Sequence {
	return _this.Get_().(Rose_Rose).Children
}

func (_this Rose) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Rose_Rose:
		{
			return "ConfluxTree.Rose.Rose" + "(" + _dafny.String(data.Value) + ", " + _dafny.String(data.Children) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Rose) Equals(other Rose) bool {
	switch data1 := _this.Get_().(type) {
	case Rose_Rose:
		{
			data2, ok := other.Get_().(Rose_Rose)
			return ok && _dafny.AreEqual(data1.Value, data2.Value) && data1.Children.Equals(data2.Children)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Rose) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Rose)
	return ok && _this.Equals(typed)
}

func Type_Rose_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_Rose_{Type_T_}
}

type type_Rose_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_Rose_) Default() interface{} {
	Type_T_ := _this.Type_T_
	_ = Type_T_
	return Companion_Rose_.Default(Type_T_.Default())
}

func (_this type_Rose_) String() string {
	return "ConfluxTree.Rose"
}
func (_this Rose) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Rose{}

// End of datatype Rose
