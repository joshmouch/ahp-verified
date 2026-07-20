// Package ConfluxMerge
// Dafny module ConfluxMerge compiled into Go

package ConfluxMerge

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-go/Changeset"
	m_Chat "github.com/joshmouch/ahp-go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-go/ClientMain"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-go/ConfluxOrderedLog"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-go/ConfluxResourceCeilings"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-go/ConfluxSeqRoute"
	m_ConfluxSupervisedOperationResult "github.com/joshmouch/ahp-go/ConfluxSupervisedOperationResult"
	m_ConfluxTree "github.com/joshmouch/ahp-go/ConfluxTree"
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
	return "ConfluxMerge.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Merge(rA func(interface{}, interface{}) interface{}, rB func(interface{}, interface{}) interface{}) func(_dafny.Tuple, Either) _dafny.Tuple {
	return (func(_0_rA func(interface{}, interface{}) interface{}, _1_rB func(interface{}, interface{}) interface{}) func(_dafny.Tuple, Either) _dafny.Tuple {
		return func(_2_s _dafny.Tuple, _3_e Either) _dafny.Tuple {
			return func() _dafny.Tuple {
				var _source0 Either = _3_e
				_ = _source0
				{
					if _source0.Is_Left() {
						var _4_a interface{} = _source0.Get_().(Either_Left).L
						_ = _4_a
						return _dafny.TupleOf((_0_rA)((*(_2_s).IndexInt(0)), _4_a), (*(_2_s).IndexInt(1)))
					}
				}
				{
					var _5_b interface{} = _source0.Get_().(Either_Right).R
					_ = _5_b
					return _dafny.TupleOf((*(_2_s).IndexInt(0)), (_1_rB)((*(_2_s).IndexInt(1)), _5_b))
				}
			}()
		}
	})(rA, rB)
}

// End of class Default__

// Definition of datatype Either
type Either struct {
	Data_Either_
}

func (_this Either) Get_() Data_Either_ {
	return _this.Data_Either_
}

type Data_Either_ interface {
	isEither()
}

type CompanionStruct_Either_ struct {
}

var Companion_Either_ = CompanionStruct_Either_{}

type Either_Left struct {
	L interface{}
}

func (Either_Left) isEither() {}

func (CompanionStruct_Either_) Create_Left_(L interface{}) Either {
	return Either{Either_Left{L}}
}

func (_this Either) Is_Left() bool {
	_, ok := _this.Get_().(Either_Left)
	return ok
}

type Either_Right struct {
	R interface{}
}

func (Either_Right) isEither() {}

func (CompanionStruct_Either_) Create_Right_(R interface{}) Either {
	return Either{Either_Right{R}}
}

func (_this Either) Is_Right() bool {
	_, ok := _this.Get_().(Either_Right)
	return ok
}

func (CompanionStruct_Either_) Default(_default_A interface{}) Either {
	return Companion_Either_.Create_Left_(_default_A)
}

func (_this Either) Dtor_l() interface{} {
	return _this.Get_().(Either_Left).L
}

func (_this Either) Dtor_r() interface{} {
	return _this.Get_().(Either_Right).R
}

func (_this Either) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Either_Left:
		{
			return "ConfluxMerge.Either.Left" + "(" + _dafny.String(data.L) + ")"
		}
	case Either_Right:
		{
			return "ConfluxMerge.Either.Right" + "(" + _dafny.String(data.R) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Either) Equals(other Either) bool {
	switch data1 := _this.Get_().(type) {
	case Either_Left:
		{
			data2, ok := other.Get_().(Either_Left)
			return ok && _dafny.AreEqual(data1.L, data2.L)
		}
	case Either_Right:
		{
			data2, ok := other.Get_().(Either_Right)
			return ok && _dafny.AreEqual(data1.R, data2.R)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Either) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Either)
	return ok && _this.Equals(typed)
}

func Type_Either_(Type_A_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_Either_{Type_A_}
}

type type_Either_ struct {
	Type_A_ _dafny.TypeDescriptor
}

func (_this type_Either_) Default() interface{} {
	Type_A_ := _this.Type_A_
	_ = Type_A_
	return Companion_Either_.Default(Type_A_.Default())
}

func (_this type_Either_) String() string {
	return "ConfluxMerge.Either"
}
func (_this Either) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Either{}

// End of datatype Either
