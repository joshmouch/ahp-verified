// Package ConfluxSeqRoute
// Dafny module ConfluxSeqRoute compiled into Go

package ConfluxSeqRoute

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ResourceWatch "github.com/joshmouch/ahp-go/ResourceWatch"
	m__System "github.com/joshmouch/ahp-go/System_"
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
	return "ConfluxSeqRoute.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) SeqUpsertById(keyOf func(interface{}) interface{}, xs _dafny.Sequence, v interface{}) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf(v))
	} else if _dafny.AreEqual((keyOf)((xs).Select(0).(interface{})), (keyOf)(v)) {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(v), (xs).Drop(1)))
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((xs).Select(0).(interface{})))
		var _in0 func(interface{}) interface{} = keyOf
		_ = _in0
		var _in1 _dafny.Sequence = (xs).Drop(1)
		_ = _in1
		var _in2 interface{} = v
		_ = _in2
		keyOf = _in0
		xs = _in1
		v = _in2
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) SeqRemoveById(keyOf func(interface{}) interface{}, xs _dafny.Sequence, k interface{}) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else if _dafny.AreEqual((keyOf)((xs).Select(0).(interface{})), k) {
		var _in0 func(interface{}) interface{} = keyOf
		_ = _in0
		var _in1 _dafny.Sequence = (xs).Drop(1)
		_ = _in1
		var _in2 interface{} = k
		_ = _in2
		keyOf = _in0
		xs = _in1
		k = _in2
		goto TAIL_CALL_START
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((xs).Select(0).(interface{})))
		var _in3 func(interface{}) interface{} = keyOf
		_ = _in3
		var _in4 _dafny.Sequence = (xs).Drop(1)
		_ = _in4
		var _in5 interface{} = k
		_ = _in5
		keyOf = _in3
		xs = _in4
		k = _in5
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) ToMap(keyOf func(interface{}) interface{}, xs _dafny.Sequence) _dafny.Map {
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.NewMapBuilder().ToMap()
	} else {
		return (Companion_Default___.ToMap(keyOf, (xs).Drop(1))).Update((keyOf)((xs).Select(0).(interface{})), (xs).Select(0).(interface{}))
	}
}
func (_static *CompanionStruct_Default___) UniqueKeys(keyOf func(interface{}) interface{}, xs _dafny.Sequence) bool {
	return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((xs).Cardinality())), true, func(_forall_var_0 _dafny.Int) bool {
		var _0_i _dafny.Int
		_0_i = interface{}(_forall_var_0).(_dafny.Int)
		return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((xs).Cardinality())), true, func(_forall_var_1 _dafny.Int) bool {
			var _1_j _dafny.Int
			_1_j = interface{}(_forall_var_1).(_dafny.Int)
			return !(((((_0_i).Sign() != -1) && ((_0_i).Cmp(_dafny.IntOfUint32((xs).Cardinality())) < 0)) && (((_1_j).Sign() != -1) && ((_1_j).Cmp(_dafny.IntOfUint32((xs).Cardinality())) < 0))) && (_dafny.AreEqual((keyOf)((xs).Select((_0_i).Uint32()).(interface{})), (keyOf)((xs).Select((_1_j).Uint32()).(interface{}))))) || ((_0_i).Cmp(_1_j) == 0)
		})
	})
}
func (_static *CompanionStruct_Default___) InstallOp(cur m_ConfluxContract.Option, v interface{}) m_ConfluxPrune.RouteOp {
	return m_ConfluxPrune.Companion_RouteOp_.Create_Install_(v)
}
func (_static *CompanionStruct_Default___) RemoveOp(cur m_ConfluxContract.Option, v interface{}) m_ConfluxPrune.RouteOp {
	return m_ConfluxPrune.Companion_RouteOp_.Create_Remove_()
}
func (_static *CompanionStruct_Default___) SeqUpdateById(keyOf func(interface{}) interface{}, xs _dafny.Sequence, k interface{}, f func(interface{}) interface{}) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else if _dafny.AreEqual((keyOf)((xs).Select(0).(interface{})), k) {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((f)((xs).Select(0).(interface{}))), (xs).Drop(1)))
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((xs).Select(0).(interface{})))
		var _in0 func(interface{}) interface{} = keyOf
		_ = _in0
		var _in1 _dafny.Sequence = (xs).Drop(1)
		_ = _in1
		var _in2 interface{} = k
		_ = _in2
		var _in3 func(interface{}) interface{} = f
		_ = _in3
		keyOf = _in0
		xs = _in1
		k = _in2
		f = _in3
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) UpdateOp(f func(interface{}) interface{}) func(m_ConfluxContract.Option, interface{}) m_ConfluxPrune.RouteOp {
	return (func(_0_f func(interface{}) interface{}) func(m_ConfluxContract.Option, interface{}) m_ConfluxPrune.RouteOp {
		return func(_1_cur m_ConfluxContract.Option, _2_a interface{}) m_ConfluxPrune.RouteOp {
			return func() m_ConfluxPrune.RouteOp {
				var _source0 m_ConfluxContract.Option = _1_cur
				_ = _source0
				{
					if _source0.Is_Some() {
						var _3_v interface{} = _source0.Get_().(m_ConfluxContract.Option_Some).Value
						_ = _3_v
						return m_ConfluxPrune.Companion_RouteOp_.Create_Install_((_0_f)(_3_v))
					}
				}
				{
					return m_ConfluxPrune.Companion_RouteOp_.Create_NoOp_()
				}
			}()
		}
	})(f)
}
func (_static *CompanionStruct_Default___) SeqUpdateByIdWhere(keyOf func(interface{}) interface{}, pred func(interface{}) bool, xs _dafny.Sequence, k interface{}, f func(interface{}) interface{}) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else if (_dafny.AreEqual((keyOf)((xs).Select(0).(interface{})), k)) && ((pred)((xs).Select(0).(interface{}))) {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((f)((xs).Select(0).(interface{}))), (xs).Drop(1)))
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((xs).Select(0).(interface{})))
		var _in0 func(interface{}) interface{} = keyOf
		_ = _in0
		var _in1 func(interface{}) bool = pred
		_ = _in1
		var _in2 _dafny.Sequence = (xs).Drop(1)
		_ = _in2
		var _in3 interface{} = k
		_ = _in3
		var _in4 func(interface{}) interface{} = f
		_ = _in4
		keyOf = _in0
		pred = _in1
		xs = _in2
		k = _in3
		f = _in4
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) UpdateWhereOp(pred func(interface{}) bool, f func(interface{}) interface{}) func(m_ConfluxContract.Option, interface{}) m_ConfluxPrune.RouteOp {
	return (func(_0_pred func(interface{}) bool, _1_f func(interface{}) interface{}) func(m_ConfluxContract.Option, interface{}) m_ConfluxPrune.RouteOp {
		return func(_2_cur m_ConfluxContract.Option, _3_a interface{}) m_ConfluxPrune.RouteOp {
			return func() m_ConfluxPrune.RouteOp {
				var _source0 m_ConfluxContract.Option = _2_cur
				_ = _source0
				{
					if _source0.Is_Some() {
						var _4_v interface{} = _source0.Get_().(m_ConfluxContract.Option_Some).Value
						_ = _4_v
						if (_0_pred)(_4_v) {
							return m_ConfluxPrune.Companion_RouteOp_.Create_Install_((_1_f)(_4_v))
						} else {
							return m_ConfluxPrune.Companion_RouteOp_.Create_NoOp_()
						}
					}
				}
				{
					return m_ConfluxPrune.Companion_RouteOp_.Create_NoOp_()
				}
			}()
		}
	})(pred, f)
}
func (_static *CompanionStruct_Default___) SeqUpdateAllWhere(keyOf func(interface{}) interface{}, pred func(interface{}) bool, xs _dafny.Sequence, k interface{}, f func(interface{}) interface{}) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (func() _dafny.Sequence {
			if (_dafny.AreEqual((keyOf)((xs).Select(0).(interface{})), k)) && ((pred)((xs).Select(0).(interface{}))) {
				return _dafny.SeqOf((f)((xs).Select(0).(interface{})))
			}
			return _dafny.SeqOf((xs).Select(0).(interface{}))
		})())
		var _in0 func(interface{}) interface{} = keyOf
		_ = _in0
		var _in1 func(interface{}) bool = pred
		_ = _in1
		var _in2 _dafny.Sequence = (xs).Drop(1)
		_ = _in2
		var _in3 interface{} = k
		_ = _in3
		var _in4 func(interface{}) interface{} = f
		_ = _in4
		keyOf = _in0
		pred = _in1
		xs = _in2
		k = _in3
		f = _in4
		goto TAIL_CALL_START
	}
}

// End of class Default__
