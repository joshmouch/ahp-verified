// Package ConfluxOrderedLog
// Dafny module ConfluxOrderedLog compiled into Go

package ConfluxOrderedLog

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-verified/go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-verified/go/Changeset"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ConfluxPrune "github.com/joshmouch/ahp-verified/go/ConfluxPrune"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-verified/go/ConfluxSeqRoute"
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
	return "ConfluxOrderedLog.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ContainsK(keyOf func(interface{}) interface{}, xs _dafny.Sequence, k interface{}) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return false
	} else if _dafny.AreEqual((keyOf)((xs).Select(0).(interface{})), k) {
		return true
	} else {
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
	}
}
func (_static *CompanionStruct_Default___) FindK(keyOf func(interface{}) interface{}, xs _dafny.Sequence, k interface{}) _dafny.Sequence {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.SeqOf()
	} else if _dafny.AreEqual((keyOf)((xs).Select(0).(interface{})), k) {
		return _dafny.SeqOf((xs).Select(0).(interface{}))
	} else {
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
	}
}
func (_static *CompanionStruct_Default___) TakeListedK(keyOf func(interface{}) interface{}, xs _dafny.Sequence, order _dafny.Sequence, taken _dafny.Set) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((order).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else if (!(taken).Contains((order).Select(0).(interface{}))) && (Companion_Default___.ContainsK(func(coer44 func(interface{}) interface{}) func(interface{}) interface{} {
		return func(arg52 interface{}) interface{} {
			return coer44(arg52)
		}
	}(keyOf), xs, (order).Select(0).(interface{}))) {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, Companion_Default___.FindK(func(coer45 func(interface{}) interface{}) func(interface{}) interface{} {
			return func(arg53 interface{}) interface{} {
				return coer45(arg53)
			}
		}(keyOf), xs, (order).Select(0).(interface{})))
		var _in0 func(interface{}) interface{} = keyOf
		_ = _in0
		var _in1 _dafny.Sequence = xs
		_ = _in1
		var _in2 _dafny.Sequence = (order).Drop(1)
		_ = _in2
		var _in3 _dafny.Set = (taken).Union(_dafny.SetOf((order).Select(0).(interface{})))
		_ = _in3
		keyOf = _in0
		xs = _in1
		order = _in2
		taken = _in3
		goto TAIL_CALL_START
	} else {
		var _in4 func(interface{}) interface{} = keyOf
		_ = _in4
		var _in5 _dafny.Sequence = xs
		_ = _in5
		var _in6 _dafny.Sequence = (order).Drop(1)
		_ = _in6
		var _in7 _dafny.Set = taken
		_ = _in7
		keyOf = _in4
		xs = _in5
		order = _in6
		taken = _in7
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) RestK(keyOf func(interface{}) interface{}, xs _dafny.Sequence, order _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (func() _dafny.Sequence {
			if _dafny.Companion_Sequence_.Contains(order, (keyOf)((xs).Select(0).(interface{}))) {
				return _dafny.SeqOf()
			}
			return _dafny.SeqOf((xs).Select(0).(interface{}))
		})())
		var _in0 func(interface{}) interface{} = keyOf
		_ = _in0
		var _in1 _dafny.Sequence = (xs).Drop(1)
		_ = _in1
		var _in2 _dafny.Sequence = order
		_ = _in2
		keyOf = _in0
		xs = _in1
		order = _in2
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) SeqReorderByKeys(keyOf func(interface{}) interface{}, order _dafny.Sequence, xs _dafny.Sequence) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.TakeListedK(func(coer46 func(interface{}) interface{}) func(interface{}) interface{} {
		return func(arg54 interface{}) interface{} {
			return coer46(arg54)
		}
	}(keyOf), xs, order, _dafny.SetOf()), Companion_Default___.RestK(func(coer47 func(interface{}) interface{}) func(interface{}) interface{} {
		return func(arg55 interface{}) interface{} {
			return coer47(arg55)
		}
	}(keyOf), xs, order))
}
func (_static *CompanionStruct_Default___) SeqKeepThrough(keyOf func(interface{}) interface{}, k interface{}, xs _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else if _dafny.AreEqual((keyOf)((xs).Select(0).(interface{})), k) {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((xs).Select(0).(interface{})))
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((xs).Select(0).(interface{})))
		var _in0 func(interface{}) interface{} = keyOf
		_ = _in0
		var _in1 interface{} = k
		_ = _in1
		var _in2 _dafny.Sequence = (xs).Drop(1)
		_ = _in2
		keyOf = _in0
		k = _in1
		xs = _in2
		goto TAIL_CALL_START
	}
}

// End of class Default__
