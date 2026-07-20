// Package ConfluxStateMachine
// Dafny module ConfluxStateMachine compiled into Go

package ConfluxStateMachine

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
	m_ConfluxDelimited "github.com/joshmouch/ahp-go/ConfluxDelimited"
	m_ConfluxFixpoint "github.com/joshmouch/ahp-go/ConfluxFixpoint"
	m_ConfluxGenericCodec "github.com/joshmouch/ahp-go/ConfluxGenericCodec"
	m_ConfluxMapValues "github.com/joshmouch/ahp-go/ConfluxMapValues"
	m_ConfluxMerge "github.com/joshmouch/ahp-go/ConfluxMerge"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-go/ConfluxOrderedLog"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxResolve "github.com/joshmouch/ahp-go/ConfluxResolve"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-go/ConfluxResourceCeilings"
	m_ConfluxRoute "github.com/joshmouch/ahp-go/ConfluxRoute"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-go/ConfluxSeqRoute"
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
	return "ConfluxStateMachine.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) TransitionState(step func(interface{}, interface{}) _dafny.Tuple) func(interface{}, interface{}) interface{} {
	return (func(_0_step func(interface{}, interface{}) _dafny.Tuple) func(interface{}, interface{}) interface{} {
		return func(_1_s interface{}, _2_a interface{}) interface{} {
			return (*((_0_step)(_1_s, _2_a)).IndexInt(0))
		}
	})(step)
}
func (_static *CompanionStruct_Default___) FinalState(run TransitionRun) interface{} {
	return (run).Dtor_state()
}
func (_static *CompanionStruct_Default___) Outputs(run TransitionRun) _dafny.Sequence {
	return (run).Dtor_outputs()
}
func (_static *CompanionStruct_Default___) RunTransitions(step func(interface{}, interface{}) _dafny.Tuple, s interface{}, xs _dafny.Sequence) TransitionRun {
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return Companion_TransitionRun_.Create_TransitionRun_(s, _dafny.SeqOf())
	} else {
		var _0_so _dafny.Tuple = (step)(s, (xs).Select(0).(interface{}))
		_ = _0_so
		var _1_rest TransitionRun = Companion_Default___.RunTransitions(step, (*(_0_so).IndexInt(0)), (xs).Drop(1))
		_ = _1_rest
		return Companion_TransitionRun_.Create_TransitionRun_((_1_rest).Dtor_state(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((*(_0_so).IndexInt(1))), (_1_rest).Dtor_outputs()))
	}
}
func (_static *CompanionStruct_Default___) GuardedTransition(base func(interface{}, interface{}) _dafny.Tuple, breached func(interface{}, interface{}) bool, onBreach func(interface{}, interface{}) interface{}) func(interface{}, interface{}) _dafny.Tuple {
	return (func(_0_breached func(interface{}, interface{}) bool, _1_onBreach func(interface{}, interface{}) interface{}, _2_base func(interface{}, interface{}) _dafny.Tuple) func(interface{}, interface{}) _dafny.Tuple {
		return func(_3_s interface{}, _4_a interface{}) _dafny.Tuple {
			return (func() _dafny.Tuple {
				if (_0_breached)(_3_s, _4_a) {
					return _dafny.TupleOf(_3_s, (_1_onBreach)(_3_s, _4_a))
				}
				return (_2_base)(_3_s, _4_a)
			})()
		}
	})(breached, onBreach, base)
}

// End of class Default__

// Definition of datatype TransitionRun
type TransitionRun struct {
	Data_TransitionRun_
}

func (_this TransitionRun) Get_() Data_TransitionRun_ {
	return _this.Data_TransitionRun_
}

type Data_TransitionRun_ interface {
	isTransitionRun()
}

type CompanionStruct_TransitionRun_ struct {
}

var Companion_TransitionRun_ = CompanionStruct_TransitionRun_{}

type TransitionRun_TransitionRun struct {
	State   interface{}
	Outputs _dafny.Sequence
}

func (TransitionRun_TransitionRun) isTransitionRun() {}

func (CompanionStruct_TransitionRun_) Create_TransitionRun_(State interface{}, Outputs _dafny.Sequence) TransitionRun {
	return TransitionRun{TransitionRun_TransitionRun{State, Outputs}}
}

func (_this TransitionRun) Is_TransitionRun() bool {
	_, ok := _this.Get_().(TransitionRun_TransitionRun)
	return ok
}

func (CompanionStruct_TransitionRun_) Default(_default_S interface{}) TransitionRun {
	return Companion_TransitionRun_.Create_TransitionRun_(_default_S, _dafny.EmptySeq)
}

func (_this TransitionRun) Dtor_state() interface{} {
	return _this.Get_().(TransitionRun_TransitionRun).State
}

func (_this TransitionRun) Dtor_outputs() _dafny.Sequence {
	return _this.Get_().(TransitionRun_TransitionRun).Outputs
}

func (_this TransitionRun) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case TransitionRun_TransitionRun:
		{
			return "ConfluxStateMachine.TransitionRun.TransitionRun" + "(" + _dafny.String(data.State) + ", " + _dafny.String(data.Outputs) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this TransitionRun) Equals(other TransitionRun) bool {
	switch data1 := _this.Get_().(type) {
	case TransitionRun_TransitionRun:
		{
			data2, ok := other.Get_().(TransitionRun_TransitionRun)
			return ok && _dafny.AreEqual(data1.State, data2.State) && data1.Outputs.Equals(data2.Outputs)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this TransitionRun) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(TransitionRun)
	return ok && _this.Equals(typed)
}

func Type_TransitionRun_(Type_S_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_TransitionRun_{Type_S_}
}

type type_TransitionRun_ struct {
	Type_S_ _dafny.TypeDescriptor
}

func (_this type_TransitionRun_) Default() interface{} {
	Type_S_ := _this.Type_S_
	_ = Type_S_
	return Companion_TransitionRun_.Default(Type_S_.Default())
}

func (_this type_TransitionRun_) String() string {
	return "ConfluxStateMachine.TransitionRun"
}
func (_this TransitionRun) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = TransitionRun{}

// End of datatype TransitionRun
