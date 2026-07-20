// Package ConfluxGenericCodec
// Dafny module ConfluxGenericCodec compiled into Go

package ConfluxGenericCodec

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
	m_ConfluxMapValues "github.com/joshmouch/ahp-go/ConfluxMapValues"
	m_ConfluxMerge "github.com/joshmouch/ahp-go/ConfluxMerge"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-go/ConfluxOrderedLog"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-go/ConfluxResourceCeilings"
	m_ConfluxRoute "github.com/joshmouch/ahp-go/ConfluxRoute"
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
var _ m_ConfluxRoute.Dummy__
var _ m_ConfluxMapValues.Dummy__
var _ m_ConfluxDelimited.Dummy__
var _ m_ConfluxFixpoint.Dummy__

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
	return "ConfluxGenericCodec.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) IdCodec() Codec {
	return Companion_Codec_.Create_Codec_(func(coer78 func(interface{}) interface{}) func(interface{}) interface{} {
		return func(arg87 interface{}) interface{} {
			return coer78(arg87)
		}
	}(func(_0_j interface{}) interface{} {
		return _0_j
	}), func(coer79 func(interface{}) Option) func(interface{}) Option {
		return func(arg88 interface{}) Option {
			return coer79(arg88)
		}
	}(func(_1_j interface{}) Option {
		return Companion_Option_.Create_Some_(_1_j)
	}))
}
func (_static *CompanionStruct_Default___) DoubleCodec() Codec {
	return Companion_Codec_.Create_Codec_(func(coer80 func(_dafny.Int) _dafny.Int) func(interface{}) interface{} {
		return func(arg89 interface{}) interface{} {
			return coer80(arg89.(_dafny.Int))
		}
	}(func(_0_n _dafny.Int) _dafny.Int {
		return (_0_n).Plus(_0_n)
	}), func(coer81 func(_dafny.Int) Option) func(interface{}) Option {
		return func(arg90 interface{}) Option {
			return coer81(arg90.(_dafny.Int))
		}
	}(func(_1_j _dafny.Int) Option {
		return (func() Option {
			if ((_1_j).Modulo(_dafny.IntOfInt64(2))).Sign() == 0 {
				return Companion_Option_.Create_Some_((_1_j).DivBy(_dafny.IntOfInt64(2)))
			}
			return Companion_Option_.Create_None_()
		})()
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
			return "ConfluxGenericCodec.Option.None"
		}
	case Option_Some:
		{
			return "ConfluxGenericCodec.Option.Some" + "(" + _dafny.String(data.Value) + ")"
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
	return "ConfluxGenericCodec.Option"
}
func (_this Option) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Option{}

// End of datatype Option

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
	Encode func(interface{}) interface{}
	Decode func(interface{}) Option
}

func (Codec_Codec) isCodec() {}

func (CompanionStruct_Codec_) Create_Codec_(Encode func(interface{}) interface{}, Decode func(interface{}) Option) Codec {
	return Codec{Codec_Codec{Encode, Decode}}
}

func (_this Codec) Is_Codec() bool {
	_, ok := _this.Get_().(Codec_Codec)
	return ok
}

func (CompanionStruct_Codec_) Default(_default_J interface{}) Codec {
	return Companion_Codec_.Create_Codec_(func(interface{}) interface{} { return _default_J }, func(interface{}) Option { return Companion_Option_.Default() })
}

func (_this Codec) Dtor_encode() func(interface{}) interface{} {
	return _this.Get_().(Codec_Codec).Encode
}

func (_this Codec) Dtor_decode() func(interface{}) Option {
	return _this.Get_().(Codec_Codec).Decode
}

func (_this Codec) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Codec_Codec:
		{
			return "ConfluxGenericCodec.Codec.Codec" + "(" + _dafny.String(data.Encode) + ", " + _dafny.String(data.Decode) + ")"
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

func Type_Codec_(Type_J_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_Codec_{Type_J_}
}

type type_Codec_ struct {
	Type_J_ _dafny.TypeDescriptor
}

func (_this type_Codec_) Default() interface{} {
	Type_J_ := _this.Type_J_
	_ = Type_J_
	return Companion_Codec_.Default(Type_J_.Default())
}

func (_this type_Codec_) String() string {
	return "ConfluxGenericCodec.Codec"
}
func (_this Codec) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Codec{}

// End of datatype Codec
