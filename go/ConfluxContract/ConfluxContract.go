// Package ConfluxContract
// Dafny module ConfluxContract compiled into Go

package ConfluxContract

import (
	os "os"

	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m__System "github.com/joshmouch/ahp-go/System_"
	_dafny "github.com/joshmouch/ahp-go/dafny"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_ConfluxCodec.Dummy__

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
	return "ConfluxContract.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Fold(r func(interface{}, interface{}) interface{}, s interface{}, actions _dafny.Sequence) interface{} {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((actions).Cardinality())).Sign() == 0 {
		return s
	} else {
		var _in0 func(interface{}, interface{}) interface{} = r
		_ = _in0
		var _in1 interface{} = (r)(s, (actions).Select(0).(interface{}))
		_ = _in1
		var _in2 _dafny.Sequence = (actions).Drop(1)
		_ = _in2
		r = _in0
		s = _in1
		actions = _in2
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Iterate(f func(interface{}) interface{}, x interface{}, n _dafny.Int) interface{} {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (n).Sign() == 0 {
		return x
	} else {
		var _in0 func(interface{}) interface{} = f
		_ = _in0
		var _in1 interface{} = (f)(x)
		_ = _in1
		var _in2 _dafny.Int = (n).Minus(_dafny.One)
		_ = _in2
		f = _in0
		x = _in1
		n = _in2
		goto TAIL_CALL_START
	}
}

// End of class Default__

// Definition of datatype Option
type Option struct {
	Data_Option_
}

func (_this Option) Get_() Data_Option_ {
	return _this.Data_Option_
}

type Data_Option_ interface {
	isOption()
}

type CompanionStruct_Option_ struct {
}

var Companion_Option_ = CompanionStruct_Option_{}

type Option_None struct {
}

func (Option_None) isOption() {}

func (CompanionStruct_Option_) Create_None_() Option {
	return Option{Option_None{}}
}

func (_this Option) Is_None() bool {
	_, ok := _this.Get_().(Option_None)
	return ok
}

type Option_Some struct {
	Value interface{}
}

func (Option_Some) isOption() {}

func (CompanionStruct_Option_) Create_Some_(Value interface{}) Option {
	return Option{Option_Some{Value}}
}

func (_this Option) Is_Some() bool {
	_, ok := _this.Get_().(Option_Some)
	return ok
}

func (CompanionStruct_Option_) Default() Option {
	return Companion_Option_.Create_None_()
}

func (_this Option) Dtor_value() interface{} {
	return _this.Get_().(Option_Some).Value
}

func (_this Option) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Option_None:
		{
			return "ConfluxContract.Option.None"
		}
	case Option_Some:
		{
			return "ConfluxContract.Option.Some" + "(" + _dafny.String(data.Value) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Option) Equals(other Option) bool {
	switch data1 := _this.Get_().(type) {
	case Option_None:
		{
			_, ok := other.Get_().(Option_None)
			return ok
		}
	case Option_Some:
		{
			data2, ok := other.Get_().(Option_Some)
			return ok && _dafny.AreEqual(data1.Value, data2.Value)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Option) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Option)
	return ok && _this.Equals(typed)
}

func Type_Option_() _dafny.TypeDescriptor {
	return type_Option_{}
}

type type_Option_ struct {
}

func (_this type_Option_) Default() interface{} {
	return Companion_Option_.Default()
}

func (_this type_Option_) String() string {
	return "ConfluxContract.Option"
}
func (_this Option) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Option{}

// End of datatype Option
