// Package ConfluxRoute
// Dafny module ConfluxRoute compiled into Go

package ConfluxRoute

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
	m_ConfluxMerge "github.com/joshmouch/ahp-go/ConfluxMerge"
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
var _ m_ConfluxMerge.Dummy__

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
	return "ConfluxRoute.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Route(keyOf func(interface{}) interface{}, r func(interface{}, interface{}) interface{}, empty interface{}) func(_dafny.Map, interface{}) _dafny.Map {
	return (func(_0_keyOf func(interface{}) interface{}, _1_r func(interface{}, interface{}) interface{}, _2_empty interface{}) func(_dafny.Map, interface{}) _dafny.Map {
		return func(_3_m _dafny.Map, _4_a interface{}) _dafny.Map {
			return func(_pat_let469_0 interface{}) _dafny.Map {
				return func(_5_k interface{}) _dafny.Map {
					return (_3_m).Update(_5_k, (_1_r)((func() interface{} {
						if (_3_m).Contains(_5_k) {
							return (_3_m).Get(_5_k).(interface{})
						}
						return _2_empty
					})(), _4_a))
				}(_pat_let469_0)
			}((_0_keyOf)(_4_a))
		}
	})(keyOf, r, empty)
}
func (_static *CompanionStruct_Default___) GuardedRoute(keyOf func(interface{}) interface{}, guard func(m_ConfluxContract.Option, interface{}) m_ConfluxContract.Option) func(_dafny.Map, interface{}) _dafny.Map {
	return (func(_0_keyOf func(interface{}) interface{}, _1_guard func(m_ConfluxContract.Option, interface{}) m_ConfluxContract.Option) func(_dafny.Map, interface{}) _dafny.Map {
		return func(_2_m _dafny.Map, _3_a interface{}) _dafny.Map {
			return func(_pat_let470_0 interface{}) _dafny.Map {
				return func(_4_k interface{}) _dafny.Map {
					return func(_pat_let471_0 m_ConfluxContract.Option) _dafny.Map {
						return func(_5_cur m_ConfluxContract.Option) _dafny.Map {
							return func() _dafny.Map {
								var _source0 m_ConfluxContract.Option = (_1_guard)(_5_cur, _3_a)
								_ = _source0
								{
									if _source0.Is_None() {
										return _2_m
									}
								}
								{
									var _6_s2 interface{} = _source0.Get_().(m_ConfluxContract.Option_Some).Value
									_ = _6_s2
									return (_2_m).Update(_4_k, _6_s2)
								}
							}()
						}(_pat_let471_0)
					}((func() m_ConfluxContract.Option {
						if (_2_m).Contains(_4_k) {
							return m_ConfluxContract.Companion_Option_.Create_Some_((_2_m).Get(_4_k).(interface{}))
						}
						return m_ConfluxContract.Companion_Option_.Create_None_()
					})())
				}(_pat_let470_0)
			}((_0_keyOf)(_3_a))
		}
	})(keyOf, guard)
}

// End of class Default__
