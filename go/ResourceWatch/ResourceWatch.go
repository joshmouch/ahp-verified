// Package ResourceWatch
// Dafny module ResourceWatch compiled into Go

package ResourceWatch

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m__System "github.com/joshmouch/ahp-go/System_"
	_dafny "github.com/joshmouch/ahp-go/dafny"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_ConfluxCodec.Dummy__
var _ m_ConfluxContract.Dummy__
var _ m_AhpSkeleton.Dummy__

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
	return "ResourceWatch.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ApplyToResourceWatch(s ResourceWatchState, a ResourceWatchAction, now _dafny.Int) _dafny.Tuple {
	var _source0 ResourceWatchAction = a
	_ = _source0
	{
		if _source0.Is_RWChanged() {
			return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
	}
}
func (_static *CompanionStruct_Default___) IsUnknownRW(a ResourceWatchAction) bool {
	return (a).Is_RWUnknown()
}
func (_static *CompanionStruct_Default___) Apply1(s ResourceWatchState, a ResourceWatchAction) ResourceWatchState {
	return (*(Companion_Default___.ApplyToResourceWatch(s, a, _dafny.IntOfInt64(9999))).IndexInt(0)).(ResourceWatchState)
}
func (_static *CompanionStruct_Default___) Fold(s ResourceWatchState, acts _dafny.Sequence) ResourceWatchState {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer3 func(ResourceWatchState, ResourceWatchAction) ResourceWatchState) func(interface{}, interface{}) interface{} {
		return func(arg4 interface{}, arg5 interface{}) interface{} {
			return coer3(arg4.(ResourceWatchState), arg5.(ResourceWatchAction))
		}
	}(Companion_Default___.Apply1), s, acts).(ResourceWatchState)
}
func (_static *CompanionStruct_Default___) RunCorpus() (_dafny.Int, _dafny.Int) {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	var total _dafny.Int = _dafny.Zero
	_ = total
	total = _dafny.IntOfInt64(2)
	pass = _dafny.Zero
	var _0_s0 ResourceWatchState
	_ = _0_s0
	_0_s0 = Companion_ResourceWatchState_.Create_ResourceWatchState_(_dafny.UnicodeSeqOfUtf8Bytes("file:///workspace"), true)
	var _1_changes m_ConfluxCodec.Json
	_ = _1_changes
	_1_changes = m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("items"), m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("uri"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("file:///workspace/a.txt"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("type"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("added"))))))))
	if (Companion_Default___.Apply1(_0_s0, Companion_ResourceWatchAction_.Create_RWChanged_(_1_changes))).Equals(_0_s0) {
		pass = (pass).Plus(_dafny.One)
	}
	var _2_s1 ResourceWatchState
	_ = _2_s1
	_2_s1 = Companion_ResourceWatchState_.Create_ResourceWatchState_(_dafny.UnicodeSeqOfUtf8Bytes("file:///workspace"), false)
	if (Companion_Default___.Apply1(_2_s1, Companion_ResourceWatchAction_.Create_RWUnknown_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("type"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("resourceWatch/unknownAction"))))))).Equals(_2_s1) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass, total
}

// End of class Default__

// Definition of datatype ResourceWatchState
type ResourceWatchState struct {
	Data_ResourceWatchState_
}

func (_this ResourceWatchState) Get_() Data_ResourceWatchState_ {
	return _this.Data_ResourceWatchState_
}

type Data_ResourceWatchState_ interface {
	isResourceWatchState()
}

type CompanionStruct_ResourceWatchState_ struct {
}

var Companion_ResourceWatchState_ = CompanionStruct_ResourceWatchState_{}

type ResourceWatchState_ResourceWatchState struct {
	Root      _dafny.Sequence
	Recursive bool
}

func (ResourceWatchState_ResourceWatchState) isResourceWatchState() {}

func (CompanionStruct_ResourceWatchState_) Create_ResourceWatchState_(Root _dafny.Sequence, Recursive bool) ResourceWatchState {
	return ResourceWatchState{ResourceWatchState_ResourceWatchState{Root, Recursive}}
}

func (_this ResourceWatchState) Is_ResourceWatchState() bool {
	_, ok := _this.Get_().(ResourceWatchState_ResourceWatchState)
	return ok
}

func (CompanionStruct_ResourceWatchState_) Default() ResourceWatchState {
	return Companion_ResourceWatchState_.Create_ResourceWatchState_(_dafny.EmptySeq, false)
}

func (_this ResourceWatchState) Dtor_root() _dafny.Sequence {
	return _this.Get_().(ResourceWatchState_ResourceWatchState).Root
}

func (_this ResourceWatchState) Dtor_recursive() bool {
	return _this.Get_().(ResourceWatchState_ResourceWatchState).Recursive
}

func (_this ResourceWatchState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ResourceWatchState_ResourceWatchState:
		{
			return "ResourceWatch.ResourceWatchState.ResourceWatchState" + "(" + data.Root.VerbatimString(true) + ", " + _dafny.String(data.Recursive) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ResourceWatchState) Equals(other ResourceWatchState) bool {
	switch data1 := _this.Get_().(type) {
	case ResourceWatchState_ResourceWatchState:
		{
			data2, ok := other.Get_().(ResourceWatchState_ResourceWatchState)
			return ok && data1.Root.Equals(data2.Root) && data1.Recursive == data2.Recursive
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ResourceWatchState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ResourceWatchState)
	return ok && _this.Equals(typed)
}

func Type_ResourceWatchState_() _dafny.TypeDescriptor {
	return type_ResourceWatchState_{}
}

type type_ResourceWatchState_ struct {
}

func (_this type_ResourceWatchState_) Default() interface{} {
	return Companion_ResourceWatchState_.Default()
}

func (_this type_ResourceWatchState_) String() string {
	return "ResourceWatch.ResourceWatchState"
}
func (_this ResourceWatchState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ResourceWatchState{}

// End of datatype ResourceWatchState

// Definition of datatype ResourceWatchAction
type ResourceWatchAction struct {
	Data_ResourceWatchAction_
}

func (_this ResourceWatchAction) Get_() Data_ResourceWatchAction_ {
	return _this.Data_ResourceWatchAction_
}

type Data_ResourceWatchAction_ interface {
	isResourceWatchAction()
}

type CompanionStruct_ResourceWatchAction_ struct {
}

var Companion_ResourceWatchAction_ = CompanionStruct_ResourceWatchAction_{}

type ResourceWatchAction_RWChanged struct {
	Changes m_ConfluxCodec.Json
}

func (ResourceWatchAction_RWChanged) isResourceWatchAction() {}

func (CompanionStruct_ResourceWatchAction_) Create_RWChanged_(Changes m_ConfluxCodec.Json) ResourceWatchAction {
	return ResourceWatchAction{ResourceWatchAction_RWChanged{Changes}}
}

func (_this ResourceWatchAction) Is_RWChanged() bool {
	_, ok := _this.Get_().(ResourceWatchAction_RWChanged)
	return ok
}

type ResourceWatchAction_RWUnknown struct {
	Raw m_ConfluxCodec.Json
}

func (ResourceWatchAction_RWUnknown) isResourceWatchAction() {}

func (CompanionStruct_ResourceWatchAction_) Create_RWUnknown_(Raw m_ConfluxCodec.Json) ResourceWatchAction {
	return ResourceWatchAction{ResourceWatchAction_RWUnknown{Raw}}
}

func (_this ResourceWatchAction) Is_RWUnknown() bool {
	_, ok := _this.Get_().(ResourceWatchAction_RWUnknown)
	return ok
}

func (CompanionStruct_ResourceWatchAction_) Default() ResourceWatchAction {
	return Companion_ResourceWatchAction_.Create_RWChanged_(m_ConfluxCodec.Companion_Json_.Default())
}

func (_this ResourceWatchAction) Dtor_changes() m_ConfluxCodec.Json {
	return _this.Get_().(ResourceWatchAction_RWChanged).Changes
}

func (_this ResourceWatchAction) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(ResourceWatchAction_RWUnknown).Raw
}

func (_this ResourceWatchAction) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ResourceWatchAction_RWChanged:
		{
			return "ResourceWatch.ResourceWatchAction.RWChanged" + "(" + _dafny.String(data.Changes) + ")"
		}
	case ResourceWatchAction_RWUnknown:
		{
			return "ResourceWatch.ResourceWatchAction.RWUnknown" + "(" + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ResourceWatchAction) Equals(other ResourceWatchAction) bool {
	switch data1 := _this.Get_().(type) {
	case ResourceWatchAction_RWChanged:
		{
			data2, ok := other.Get_().(ResourceWatchAction_RWChanged)
			return ok && data1.Changes.Equals(data2.Changes)
		}
	case ResourceWatchAction_RWUnknown:
		{
			data2, ok := other.Get_().(ResourceWatchAction_RWUnknown)
			return ok && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ResourceWatchAction) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ResourceWatchAction)
	return ok && _this.Equals(typed)
}

func Type_ResourceWatchAction_() _dafny.TypeDescriptor {
	return type_ResourceWatchAction_{}
}

type type_ResourceWatchAction_ struct {
}

func (_this type_ResourceWatchAction_) Default() interface{} {
	return Companion_ResourceWatchAction_.Default()
}

func (_this type_ResourceWatchAction_) String() string {
	return "ResourceWatch.ResourceWatchAction"
}
func (_this ResourceWatchAction) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ResourceWatchAction{}

// End of datatype ResourceWatchAction
