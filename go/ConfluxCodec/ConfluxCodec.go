// Package ConfluxCodec
// Dafny module ConfluxCodec compiled into Go

package ConfluxCodec

import (
	os "os"

	m__System "github.com/joshmouch/ahp-verified/go/System_"
	_dafny "github.com/joshmouch/ahp-verified/go/dafny"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__

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
	return "ConfluxCodec.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Field(j Json, key _dafny.Sequence) Json {
	if ((j).Is_JObj()) && (((j).Dtor_fields()).Contains(key)) {
		return ((j).Dtor_fields()).Get(key).(Json)
	} else {
		return Companion_Json_.Create_JNull_()
	}
}
func (_static *CompanionStruct_Default___) AsInt(j Json) _dafny.Int {
	if (j).Is_JNum() {
		return (j).Dtor_n()
	} else {
		return _dafny.Zero
	}
}
func (_static *CompanionStruct_Default___) AsStr(j Json) _dafny.Sequence {
	if (j).Is_JStr() {
		return (j).Dtor_s()
	} else {
		return _dafny.UnicodeSeqOfUtf8Bytes("")
	}
}
func (_static *CompanionStruct_Default___) AsNat(j Json) _dafny.Int {
	var _0_x _dafny.Int = Companion_Default___.AsInt(j)
	_ = _0_x
	if (_0_x).Sign() != -1 {
		return _0_x
	} else {
		return _dafny.Zero
	}
}
func (_static *CompanionStruct_Default___) JsonCodec() Codec {
	return Companion_Codec_.Create_Codec_(func(coer0 func(Json) Json) func(interface{}) Json {
		return func(arg0 interface{}) Json {
			return coer0(arg0.(Json))
		}
	}(func(_0_j Json) Json {
		return _0_j
	}), func(coer1 func(Json) Option) func(Json) Option {
		return func(arg1 Json) Option {
			return coer1(arg1)
		}
	}(func(_1_j Json) Option {
		return Companion_Option_.Create_Some_(_1_j)
	}))
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
			return "ConfluxCodec.Option.None"
		}
	case Option_Some:
		{
			return "ConfluxCodec.Option.Some" + "(" + _dafny.String(data.Value) + ")"
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
	return "ConfluxCodec.Option"
}
func (_this Option) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Option{}

// End of datatype Option

// Definition of datatype Json
type Json struct {
	Data_Json_
}

func (_this Json) Get_() Data_Json_ {
	return _this.Data_Json_
}

type Data_Json_ interface {
	isJson()
}

type CompanionStruct_Json_ struct {
}

var Companion_Json_ = CompanionStruct_Json_{}

type Json_JNull struct {
}

func (Json_JNull) isJson() {}

func (CompanionStruct_Json_) Create_JNull_() Json {
	return Json{Json_JNull{}}
}

func (_this Json) Is_JNull() bool {
	_, ok := _this.Get_().(Json_JNull)
	return ok
}

type Json_JBool struct {
	B bool
}

func (Json_JBool) isJson() {}

func (CompanionStruct_Json_) Create_JBool_(B bool) Json {
	return Json{Json_JBool{B}}
}

func (_this Json) Is_JBool() bool {
	_, ok := _this.Get_().(Json_JBool)
	return ok
}

type Json_JNum struct {
	N _dafny.Int
}

func (Json_JNum) isJson() {}

func (CompanionStruct_Json_) Create_JNum_(N _dafny.Int) Json {
	return Json{Json_JNum{N}}
}

func (_this Json) Is_JNum() bool {
	_, ok := _this.Get_().(Json_JNum)
	return ok
}

type Json_JDec struct {
	Mantissa _dafny.Int
	Exp      _dafny.Int
}

func (Json_JDec) isJson() {}

func (CompanionStruct_Json_) Create_JDec_(Mantissa _dafny.Int, Exp _dafny.Int) Json {
	return Json{Json_JDec{Mantissa, Exp}}
}

func (_this Json) Is_JDec() bool {
	_, ok := _this.Get_().(Json_JDec)
	return ok
}

type Json_JStr struct {
	S _dafny.Sequence
}

func (Json_JStr) isJson() {}

func (CompanionStruct_Json_) Create_JStr_(S _dafny.Sequence) Json {
	return Json{Json_JStr{S}}
}

func (_this Json) Is_JStr() bool {
	_, ok := _this.Get_().(Json_JStr)
	return ok
}

type Json_JArr struct {
	Elems _dafny.Sequence
}

func (Json_JArr) isJson() {}

func (CompanionStruct_Json_) Create_JArr_(Elems _dafny.Sequence) Json {
	return Json{Json_JArr{Elems}}
}

func (_this Json) Is_JArr() bool {
	_, ok := _this.Get_().(Json_JArr)
	return ok
}

type Json_JObj struct {
	Fields _dafny.Map
}

func (Json_JObj) isJson() {}

func (CompanionStruct_Json_) Create_JObj_(Fields _dafny.Map) Json {
	return Json{Json_JObj{Fields}}
}

func (_this Json) Is_JObj() bool {
	_, ok := _this.Get_().(Json_JObj)
	return ok
}

func (CompanionStruct_Json_) Default() Json {
	return Companion_Json_.Create_JNull_()
}

func (_this Json) Dtor_b() bool {
	return _this.Get_().(Json_JBool).B
}

func (_this Json) Dtor_n() _dafny.Int {
	return _this.Get_().(Json_JNum).N
}

func (_this Json) Dtor_mantissa() _dafny.Int {
	return _this.Get_().(Json_JDec).Mantissa
}

func (_this Json) Dtor_exp() _dafny.Int {
	return _this.Get_().(Json_JDec).Exp
}

func (_this Json) Dtor_s() _dafny.Sequence {
	return _this.Get_().(Json_JStr).S
}

func (_this Json) Dtor_elems() _dafny.Sequence {
	return _this.Get_().(Json_JArr).Elems
}

func (_this Json) Dtor_fields() _dafny.Map {
	return _this.Get_().(Json_JObj).Fields
}

func (_this Json) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Json_JNull:
		{
			return "ConfluxCodec.Json.JNull"
		}
	case Json_JBool:
		{
			return "ConfluxCodec.Json.JBool" + "(" + _dafny.String(data.B) + ")"
		}
	case Json_JNum:
		{
			return "ConfluxCodec.Json.JNum" + "(" + _dafny.String(data.N) + ")"
		}
	case Json_JDec:
		{
			return "ConfluxCodec.Json.JDec" + "(" + _dafny.String(data.Mantissa) + ", " + _dafny.String(data.Exp) + ")"
		}
	case Json_JStr:
		{
			return "ConfluxCodec.Json.JStr" + "(" + data.S.VerbatimString(true) + ")"
		}
	case Json_JArr:
		{
			return "ConfluxCodec.Json.JArr" + "(" + _dafny.String(data.Elems) + ")"
		}
	case Json_JObj:
		{
			return "ConfluxCodec.Json.JObj" + "(" + _dafny.String(data.Fields) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Json) Equals(other Json) bool {
	switch data1 := _this.Get_().(type) {
	case Json_JNull:
		{
			_, ok := other.Get_().(Json_JNull)
			return ok
		}
	case Json_JBool:
		{
			data2, ok := other.Get_().(Json_JBool)
			return ok && data1.B == data2.B
		}
	case Json_JNum:
		{
			data2, ok := other.Get_().(Json_JNum)
			return ok && data1.N.Cmp(data2.N) == 0
		}
	case Json_JDec:
		{
			data2, ok := other.Get_().(Json_JDec)
			return ok && data1.Mantissa.Cmp(data2.Mantissa) == 0 && data1.Exp.Cmp(data2.Exp) == 0
		}
	case Json_JStr:
		{
			data2, ok := other.Get_().(Json_JStr)
			return ok && data1.S.Equals(data2.S)
		}
	case Json_JArr:
		{
			data2, ok := other.Get_().(Json_JArr)
			return ok && data1.Elems.Equals(data2.Elems)
		}
	case Json_JObj:
		{
			data2, ok := other.Get_().(Json_JObj)
			return ok && data1.Fields.Equals(data2.Fields)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Json) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Json)
	return ok && _this.Equals(typed)
}

func Type_Json_() _dafny.TypeDescriptor {
	return type_Json_{}
}

type type_Json_ struct {
}

func (_this type_Json_) Default() interface{} {
	return Companion_Json_.Default()
}

func (_this type_Json_) String() string {
	return "ConfluxCodec.Json"
}
func (_this Json) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Json{}

// End of datatype Json

// Definition of datatype Codec
type Codec struct {
	Data_Codec_
}

func (_this Codec) Get_() Data_Codec_ {
	return _this.Data_Codec_
}

type Data_Codec_ interface {
	isCodec()
}

type CompanionStruct_Codec_ struct {
}

var Companion_Codec_ = CompanionStruct_Codec_{}

type Codec_Codec struct {
	Encode func(interface{}) Json
	Decode func(Json) Option
}

func (Codec_Codec) isCodec() {}

func (CompanionStruct_Codec_) Create_Codec_(Encode func(interface{}) Json, Decode func(Json) Option) Codec {
	return Codec{Codec_Codec{Encode, Decode}}
}

func (_this Codec) Is_Codec() bool {
	_, ok := _this.Get_().(Codec_Codec)
	return ok
}

func (CompanionStruct_Codec_) Default() Codec {
	return Companion_Codec_.Create_Codec_(func(interface{}) Json { return Companion_Json_.Default() }, func(Json) Option { return Companion_Option_.Default() })
}

func (_this Codec) Dtor_encode() func(interface{}) Json {
	return _this.Get_().(Codec_Codec).Encode
}

func (_this Codec) Dtor_decode() func(Json) Option {
	return _this.Get_().(Codec_Codec).Decode
}

func (_this Codec) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Codec_Codec:
		{
			return "ConfluxCodec.Codec.Codec" + "(" + _dafny.String(data.Encode) + ", " + _dafny.String(data.Decode) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Codec) Equals(other Codec) bool {
	switch data1 := _this.Get_().(type) {
	case Codec_Codec:
		{
			data2, ok := other.Get_().(Codec_Codec)
			return ok && _dafny.AreEqual(data1.Encode, data2.Encode) && _dafny.AreEqual(data1.Decode, data2.Decode)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Codec) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Codec)
	return ok && _this.Equals(typed)
}

func Type_Codec_() _dafny.TypeDescriptor {
	return type_Codec_{}
}

type type_Codec_ struct {
}

func (_this type_Codec_) Default() interface{} {
	return Companion_Codec_.Default()
}

func (_this type_Codec_) String() string {
	return "ConfluxCodec.Codec"
}
func (_this Codec) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Codec{}

// End of datatype Codec
