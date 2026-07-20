// Package ConfluxSeqRouteMerge
// Dafny module ConfluxSeqRouteMerge compiled into Go

package ConfluxSeqRouteMerge

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-verified/go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-verified/go/Changeset"
	m_Chat "github.com/joshmouch/ahp-verified/go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-verified/go/ClientMain"
	m_ConfluxBatchRoute "github.com/joshmouch/ahp-verified/go/ConfluxBatchRoute"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-verified/go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxDedupe "github.com/joshmouch/ahp-verified/go/ConfluxDedupe"
	m_ConfluxDelimited "github.com/joshmouch/ahp-verified/go/ConfluxDelimited"
	m_ConfluxFixpoint "github.com/joshmouch/ahp-verified/go/ConfluxFixpoint"
	m_ConfluxGenericCodec "github.com/joshmouch/ahp-verified/go/ConfluxGenericCodec"
	m_ConfluxJoin "github.com/joshmouch/ahp-verified/go/ConfluxJoin"
	m_ConfluxMapValues "github.com/joshmouch/ahp-verified/go/ConfluxMapValues"
	m_ConfluxMerge "github.com/joshmouch/ahp-verified/go/ConfluxMerge"
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-verified/go/ConfluxOrderedLog"
	m_ConfluxProduct "github.com/joshmouch/ahp-verified/go/ConfluxProduct"
	m_ConfluxPrune "github.com/joshmouch/ahp-verified/go/ConfluxPrune"
	m_ConfluxResolve "github.com/joshmouch/ahp-verified/go/ConfluxResolve"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-verified/go/ConfluxResourceCeilings"
	m_ConfluxRoute "github.com/joshmouch/ahp-verified/go/ConfluxRoute"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-verified/go/ConfluxSeqRoute"
	m_ConfluxStateMachine "github.com/joshmouch/ahp-verified/go/ConfluxStateMachine"
	m_ConfluxSupervisedOperationResult "github.com/joshmouch/ahp-verified/go/ConfluxSupervisedOperationResult"
	m_ConfluxTree "github.com/joshmouch/ahp-verified/go/ConfluxTree"
	m_ConfluxWatermark "github.com/joshmouch/ahp-verified/go/ConfluxWatermark"
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
	return "ConfluxSeqRouteMerge.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) SeqUpsertByIdWith(keyOf func(interface{}) interface{}, combine func(m_ConfluxContract.Option, interface{}) interface{}, xs _dafny.Sequence, v interface{}) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((combine)(m_ConfluxContract.Companion_Option_.Create_None_(), v)))
	} else if _dafny.AreEqual((keyOf)((xs).Select(0).(interface{})), (keyOf)(v)) {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((combine)(m_ConfluxContract.Companion_Option_.Create_Some_((xs).Select(0).(interface{})), v)), (xs).Drop(1)))
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((xs).Select(0).(interface{})))
		var _in0 func(interface{}) interface{} = keyOf
		_ = _in0
		var _in1 func(m_ConfluxContract.Option, interface{}) interface{} = combine
		_ = _in1
		var _in2 _dafny.Sequence = (xs).Drop(1)
		_ = _in2
		var _in3 interface{} = v
		_ = _in3
		keyOf = _in0
		combine = _in1
		xs = _in2
		v = _in3
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) MergeUpsertById(keyOf func(interface{}) interface{}, combine func(m_ConfluxContract.Option, interface{}) interface{}, xs _dafny.Sequence, v interface{}) _dafny.Sequence {
	return Companion_Default___.SeqUpsertByIdWith(func(coer86 func(interface{}) interface{}) func(interface{}) interface{} {
		return func(arg98 interface{}) interface{} {
			return coer86(arg98)
		}
	}(keyOf), func(coer87 func(m_ConfluxContract.Option, interface{}) interface{}) func(m_ConfluxContract.Option, interface{}) interface{} {
		return func(arg99 m_ConfluxContract.Option, arg100 interface{}) interface{} {
			return coer87(arg99, arg100)
		}
	}(combine), xs, v)
}
func (_static *CompanionStruct_Default___) UpsertWithOp(combine func(m_ConfluxContract.Option, interface{}) interface{}) func(m_ConfluxContract.Option, interface{}) m_ConfluxPrune.RouteOp {
	return (func(_0_combine func(m_ConfluxContract.Option, interface{}) interface{}) func(m_ConfluxContract.Option, interface{}) m_ConfluxPrune.RouteOp {
		return func(_1_cur m_ConfluxContract.Option, _2_a interface{}) m_ConfluxPrune.RouteOp {
			return m_ConfluxPrune.Companion_RouteOp_.Create_Install_((_0_combine)(_1_cur, _2_a))
		}
	})(combine)
}

// End of class Default__
