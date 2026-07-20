// Package ConfluxPrune
// Dafny module ConfluxPrune compiled into Go

package ConfluxPrune

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ResourceWatch "github.com/joshmouch/ahp-verified/go/ResourceWatch"
	m__System "github.com/joshmouch/ahp-verified/go/System_"
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
	return "ConfluxPrune.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) PrunableRoute(keyOf func(interface{}) interface{}, op func(m_ConfluxContract.Option, interface{}) RouteOp) func(_dafny.Map, interface{}) _dafny.Map {
	return (func(_0_keyOf func(interface{}) interface{}, _1_op func(m_ConfluxContract.Option, interface{}) RouteOp) func(_dafny.Map, interface{}) _dafny.Map {
		return func(_2_m _dafny.Map, _3_a interface{}) _dafny.Map {
			return func(_pat_let24_0 interface{}) _dafny.Map {
				return func(_4_k interface{}) _dafny.Map {
					return func(_pat_let25_0 m_ConfluxContract.Option) _dafny.Map {
						return func(_5_cur m_ConfluxContract.Option) _dafny.Map {
							return func() _dafny.Map {
								var _source0 RouteOp = (_1_op)(_5_cur, _3_a)
								_ = _source0
								{
									if _source0.Is_NoOp() {
										return _2_m
									}
								}
								{
									if _source0.Is_Install() {
										var _6_s interface{} = _source0.Get_().(RouteOp_Install).Next
										_ = _6_s
										return (_2_m).Update(_4_k, _6_s)
									}
								}
								{
									return (_2_m).Subtract(_dafny.SetOf(_4_k))
								}
							}()
						}(_pat_let25_0)
					}((func() m_ConfluxContract.Option {
						if (_2_m).Contains(_4_k) {
							return m_ConfluxContract.Companion_Option_.Create_Some_((_2_m).Get(_4_k).(interface{}))
						}
						return m_ConfluxContract.Companion_Option_.Create_None_()
					})())
				}(_pat_let24_0)
			}((_0_keyOf)(_3_a))
		}
	})(keyOf, op)
}

// End of class Default__

// Definition of datatype RouteOp
type RouteOp struct {
	Data_RouteOp_
}

func (_this RouteOp) Get_() Data_RouteOp_ {
	return _this.Data_RouteOp_
}

type Data_RouteOp_ interface {
	isRouteOp()
}

type CompanionStruct_RouteOp_ struct {
}

var Companion_RouteOp_ = CompanionStruct_RouteOp_{}

type RouteOp_NoOp struct {
}

func (RouteOp_NoOp) isRouteOp() {}

func (CompanionStruct_RouteOp_) Create_NoOp_() RouteOp {
	return RouteOp{RouteOp_NoOp{}}
}

func (_this RouteOp) Is_NoOp() bool {
	_, ok := _this.Get_().(RouteOp_NoOp)
	return ok
}

type RouteOp_Install struct {
	Next interface{}
}

func (RouteOp_Install) isRouteOp() {}

func (CompanionStruct_RouteOp_) Create_Install_(Next interface{}) RouteOp {
	return RouteOp{RouteOp_Install{Next}}
}

func (_this RouteOp) Is_Install() bool {
	_, ok := _this.Get_().(RouteOp_Install)
	return ok
}

type RouteOp_Remove struct {
}

func (RouteOp_Remove) isRouteOp() {}

func (CompanionStruct_RouteOp_) Create_Remove_() RouteOp {
	return RouteOp{RouteOp_Remove{}}
}

func (_this RouteOp) Is_Remove() bool {
	_, ok := _this.Get_().(RouteOp_Remove)
	return ok
}

func (CompanionStruct_RouteOp_) Default() RouteOp {
	return Companion_RouteOp_.Create_NoOp_()
}

func (_this RouteOp) Dtor_next() interface{} {
	return _this.Get_().(RouteOp_Install).Next
}

func (_this RouteOp) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RouteOp_NoOp:
		{
			return "ConfluxPrune.RouteOp.NoOp"
		}
	case RouteOp_Install:
		{
			return "ConfluxPrune.RouteOp.Install" + "(" + _dafny.String(data.Next) + ")"
		}
	case RouteOp_Remove:
		{
			return "ConfluxPrune.RouteOp.Remove"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RouteOp) Equals(other RouteOp) bool {
	switch data1 := _this.Get_().(type) {
	case RouteOp_NoOp:
		{
			_, ok := other.Get_().(RouteOp_NoOp)
			return ok
		}
	case RouteOp_Install:
		{
			data2, ok := other.Get_().(RouteOp_Install)
			return ok && _dafny.AreEqual(data1.Next, data2.Next)
		}
	case RouteOp_Remove:
		{
			_, ok := other.Get_().(RouteOp_Remove)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RouteOp) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RouteOp)
	return ok && _this.Equals(typed)
}

func Type_RouteOp_() _dafny.TypeDescriptor {
	return type_RouteOp_{}
}

type type_RouteOp_ struct {
}

func (_this type_RouteOp_) Default() interface{} {
	return Companion_RouteOp_.Default()
}

func (_this type_RouteOp_) String() string {
	return "ConfluxPrune.RouteOp"
}
func (_this RouteOp) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RouteOp{}

// End of datatype RouteOp
