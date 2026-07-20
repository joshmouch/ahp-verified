// Package ConfluxOperators
// Dafny module ConfluxOperators compiled into Go

package ConfluxOperators

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
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
	return "ConfluxOperators.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Map(f func(interface{}) interface{}, xs _dafny.Sequence) _dafny.Sequence {
	return _dafny.SeqCreate((_dafny.IntOfUint32((xs).Cardinality())).Uint32(), func(coer5 func(_dafny.Int) interface{}) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer5(arg8)
		}
	}((func(_0_f func(interface{}) interface{}, _1_xs _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(_2_i _dafny.Int) interface{} {
			return (_0_f)((_1_xs).Select((_2_i).Uint32()).(interface{}))
		}
	})(f, xs)))
}
func (_static *CompanionStruct_Default___) Filter(p func(interface{}) bool, xs _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (func() _dafny.Sequence {
			if (p)((xs).Select(0).(interface{})) {
				return _dafny.SeqOf((xs).Select(0).(interface{}))
			}
			return _dafny.SeqOf()
		})())
		var _in0 func(interface{}) bool = p
		_ = _in0
		var _in1 _dafny.Sequence = (xs).Drop(1)
		_ = _in1
		p = _in0
		xs = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) MapFold(f func(interface{}) interface{}, r func(interface{}, interface{}) interface{}, init interface{}, xs _dafny.Sequence) interface{} {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer6 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
		return func(arg9 interface{}, arg10 interface{}) interface{} {
			return coer6(arg9, arg10)
		}
	}(r), init, Companion_Default___.Map(func(coer7 func(interface{}) interface{}) func(interface{}) interface{} {
		return func(arg11 interface{}) interface{} {
			return coer7(arg11)
		}
	}(f), xs))
}

// End of class Default__
