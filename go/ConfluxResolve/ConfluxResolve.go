// Package ConfluxResolve
// Dafny module ConfluxResolve compiled into Go

package ConfluxResolve

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-verified/go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-verified/go/Changeset"
	m_Chat "github.com/joshmouch/ahp-verified/go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-verified/go/ClientMain"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-verified/go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxDelimited "github.com/joshmouch/ahp-verified/go/ConfluxDelimited"
	m_ConfluxFixpoint "github.com/joshmouch/ahp-verified/go/ConfluxFixpoint"
	m_ConfluxGenericCodec "github.com/joshmouch/ahp-verified/go/ConfluxGenericCodec"
	m_ConfluxMapValues "github.com/joshmouch/ahp-verified/go/ConfluxMapValues"
	m_ConfluxMerge "github.com/joshmouch/ahp-verified/go/ConfluxMerge"
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-verified/go/ConfluxOrderedLog"
	m_ConfluxPrune "github.com/joshmouch/ahp-verified/go/ConfluxPrune"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-verified/go/ConfluxResourceCeilings"
	m_ConfluxRoute "github.com/joshmouch/ahp-verified/go/ConfluxRoute"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-verified/go/ConfluxSeqRoute"
	m_ConfluxSupervisedOperationResult "github.com/joshmouch/ahp-verified/go/ConfluxSupervisedOperationResult"
	m_ConfluxTree "github.com/joshmouch/ahp-verified/go/ConfluxTree"
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
	return "ConfluxResolve.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ResolveLookup(rk Option, snapshot _dafny.Map) ResolveVerdict {
	var _source0 Option = rk
	_ = _source0
	{
		if _source0.Is_None() {
			return Companion_ResolveVerdict_.Create_NoKeyResolved_()
		}
	}
	{
		var _0_k interface{} = _source0.Get_().(Option_Some).Value
		_ = _0_k
		if (snapshot).Contains(_0_k) {
			return Companion_ResolveVerdict_.Create_Matched_((snapshot).Get(_0_k).(interface{}))
		} else {
			return Companion_ResolveVerdict_.Create_KeyResolvedButAbsent_()
		}
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
			return "ConfluxResolve.Option.None"
		}
	case Option_Some:
		{
			return "ConfluxResolve.Option.Some" + "(" + _dafny.String(data.Value) + ")"
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
	return "ConfluxResolve.Option"
}
func (_this Option) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Option{}

// End of datatype Option

// Definition of datatype ResolveVerdict
type ResolveVerdict struct {
	Data_ResolveVerdict_
}

func (_this ResolveVerdict) Get_() Data_ResolveVerdict_ {
	return _this.Data_ResolveVerdict_
}

type Data_ResolveVerdict_ interface {
	isResolveVerdict()
}

type CompanionStruct_ResolveVerdict_ struct {
}

var Companion_ResolveVerdict_ = CompanionStruct_ResolveVerdict_{}

type ResolveVerdict_NoKeyResolved struct {
}

func (ResolveVerdict_NoKeyResolved) isResolveVerdict() {}

func (CompanionStruct_ResolveVerdict_) Create_NoKeyResolved_() ResolveVerdict {
	return ResolveVerdict{ResolveVerdict_NoKeyResolved{}}
}

func (_this ResolveVerdict) Is_NoKeyResolved() bool {
	_, ok := _this.Get_().(ResolveVerdict_NoKeyResolved)
	return ok
}

type ResolveVerdict_KeyResolvedButAbsent struct {
}

func (ResolveVerdict_KeyResolvedButAbsent) isResolveVerdict() {}

func (CompanionStruct_ResolveVerdict_) Create_KeyResolvedButAbsent_() ResolveVerdict {
	return ResolveVerdict{ResolveVerdict_KeyResolvedButAbsent{}}
}

func (_this ResolveVerdict) Is_KeyResolvedButAbsent() bool {
	_, ok := _this.Get_().(ResolveVerdict_KeyResolvedButAbsent)
	return ok
}

type ResolveVerdict_Matched struct {
	Value interface{}
}

func (ResolveVerdict_Matched) isResolveVerdict() {}

func (CompanionStruct_ResolveVerdict_) Create_Matched_(Value interface{}) ResolveVerdict {
	return ResolveVerdict{ResolveVerdict_Matched{Value}}
}

func (_this ResolveVerdict) Is_Matched() bool {
	_, ok := _this.Get_().(ResolveVerdict_Matched)
	return ok
}

func (CompanionStruct_ResolveVerdict_) Default() ResolveVerdict {
	return Companion_ResolveVerdict_.Create_NoKeyResolved_()
}

func (_this ResolveVerdict) Dtor_value() interface{} {
	return _this.Get_().(ResolveVerdict_Matched).Value
}

func (_this ResolveVerdict) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ResolveVerdict_NoKeyResolved:
		{
			return "ConfluxResolve.ResolveVerdict.NoKeyResolved"
		}
	case ResolveVerdict_KeyResolvedButAbsent:
		{
			return "ConfluxResolve.ResolveVerdict.KeyResolvedButAbsent"
		}
	case ResolveVerdict_Matched:
		{
			return "ConfluxResolve.ResolveVerdict.Matched" + "(" + _dafny.String(data.Value) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ResolveVerdict) Equals(other ResolveVerdict) bool {
	switch data1 := _this.Get_().(type) {
	case ResolveVerdict_NoKeyResolved:
		{
			_, ok := other.Get_().(ResolveVerdict_NoKeyResolved)
			return ok
		}
	case ResolveVerdict_KeyResolvedButAbsent:
		{
			_, ok := other.Get_().(ResolveVerdict_KeyResolvedButAbsent)
			return ok
		}
	case ResolveVerdict_Matched:
		{
			data2, ok := other.Get_().(ResolveVerdict_Matched)
			return ok && _dafny.AreEqual(data1.Value, data2.Value)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ResolveVerdict) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ResolveVerdict)
	return ok && _this.Equals(typed)
}

func Type_ResolveVerdict_() _dafny.TypeDescriptor {
	return type_ResolveVerdict_{}
}

type type_ResolveVerdict_ struct {
}

func (_this type_ResolveVerdict_) Default() interface{} {
	return Companion_ResolveVerdict_.Default()
}

func (_this type_ResolveVerdict_) String() string {
	return "ConfluxResolve.ResolveVerdict"
}
func (_this ResolveVerdict) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ResolveVerdict{}

// End of datatype ResolveVerdict
