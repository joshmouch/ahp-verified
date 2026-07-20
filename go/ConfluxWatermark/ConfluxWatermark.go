// Package ConfluxWatermark
// Dafny module ConfluxWatermark compiled into Go

package ConfluxWatermark

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
	m_ConfluxResolve "github.com/joshmouch/ahp-verified/go/ConfluxResolve"
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
var _ m_ConfluxResolve.Dummy__

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
	return "ConfluxWatermark.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Slack(c Cap) _dafny.Int {
	return ((c).Dtor_cap()).Minus((c).Dtor_pos())
}
func (_static *CompanionStruct_Default___) TryAdvance(c Cap) _dafny.Tuple {
	if ((c).Dtor_pos()).Cmp((c).Dtor_cap()) < 0 {
		return _dafny.TupleOf(Companion_Cap_.Create_Cap_(((c).Dtor_pos()).Plus(_dafny.One), (c).Dtor_cap()), Companion_AdvOutcome_.Create_Advanced_())
	} else {
		return _dafny.TupleOf(c, Companion_AdvOutcome_.Create_Declined_())
	}
}
func (_static *CompanionStruct_Default___) AdvanceN(c Cap, n _dafny.Int) Cap {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (n).Sign() == 0 {
		return c
	} else {
		var _in0 Cap = (*(Companion_Default___.TryAdvance(c)).IndexInt(0)).(Cap)
		_ = _in0
		var _in1 _dafny.Int = (n).Minus(_dafny.One)
		_ = _in1
		c = _in0
		n = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Raise(c Cap, incoming _dafny.Int) Cap {
	return Companion_Cap_.Create_Cap_((c).Dtor_pos(), (func() _dafny.Int {
		if (incoming).Cmp((c).Dtor_cap()) >= 0 {
			return incoming
		}
		return (c).Dtor_cap()
	})())
}

// End of class Default__

// Definition of datatype Cap
type Cap struct {
	Data_Cap_
}

func (_this Cap) Get_() Data_Cap_ {
	return _this.Data_Cap_
}

type Data_Cap_ interface {
	isCap()
}

type CompanionStruct_Cap_ struct {
}

var Companion_Cap_ = CompanionStruct_Cap_{}

type Cap_Cap struct {
	Pos _dafny.Int
	Cap _dafny.Int
}

func (Cap_Cap) isCap() {}

func (CompanionStruct_Cap_) Create_Cap_(Pos _dafny.Int, Cap_ _dafny.Int) Cap {
	return Cap{Cap_Cap{Pos, Cap_}}
}

func (_this Cap) Is_Cap() bool {
	_, ok := _this.Get_().(Cap_Cap)
	return ok
}

func (CompanionStruct_Cap_) Default() Cap {
	return Companion_Cap_.Create_Cap_(_dafny.Zero, _dafny.Zero)
}

func (_this Cap) Dtor_pos() _dafny.Int {
	return _this.Get_().(Cap_Cap).Pos
}

func (_this Cap) Dtor_cap() _dafny.Int {
	return _this.Get_().(Cap_Cap).Cap
}

func (_this Cap) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Cap_Cap:
		{
			return "ConfluxWatermark.Cap.Cap" + "(" + _dafny.String(data.Pos) + ", " + _dafny.String(data.Cap) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Cap) Equals(other Cap) bool {
	switch data1 := _this.Get_().(type) {
	case Cap_Cap:
		{
			data2, ok := other.Get_().(Cap_Cap)
			return ok && data1.Pos.Cmp(data2.Pos) == 0 && data1.Cap.Cmp(data2.Cap) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Cap) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Cap)
	return ok && _this.Equals(typed)
}

func Type_Cap_() _dafny.TypeDescriptor {
	return type_Cap_{}
}

type type_Cap_ struct {
}

func (_this type_Cap_) Default() interface{} {
	return Companion_Cap_.Default()
}

func (_this type_Cap_) String() string {
	return "ConfluxWatermark.Cap"
}
func (_this Cap) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Cap{}

// End of datatype Cap

// Definition of datatype AdvOutcome
type AdvOutcome struct {
	Data_AdvOutcome_
}

func (_this AdvOutcome) Get_() Data_AdvOutcome_ {
	return _this.Data_AdvOutcome_
}

type Data_AdvOutcome_ interface {
	isAdvOutcome()
}

type CompanionStruct_AdvOutcome_ struct {
}

var Companion_AdvOutcome_ = CompanionStruct_AdvOutcome_{}

type AdvOutcome_Advanced struct {
}

func (AdvOutcome_Advanced) isAdvOutcome() {}

func (CompanionStruct_AdvOutcome_) Create_Advanced_() AdvOutcome {
	return AdvOutcome{AdvOutcome_Advanced{}}
}

func (_this AdvOutcome) Is_Advanced() bool {
	_, ok := _this.Get_().(AdvOutcome_Advanced)
	return ok
}

type AdvOutcome_Declined struct {
}

func (AdvOutcome_Declined) isAdvOutcome() {}

func (CompanionStruct_AdvOutcome_) Create_Declined_() AdvOutcome {
	return AdvOutcome{AdvOutcome_Declined{}}
}

func (_this AdvOutcome) Is_Declined() bool {
	_, ok := _this.Get_().(AdvOutcome_Declined)
	return ok
}

func (CompanionStruct_AdvOutcome_) Default() AdvOutcome {
	return Companion_AdvOutcome_.Create_Advanced_()
}

func (_ CompanionStruct_AdvOutcome_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_AdvOutcome_.Create_Advanced_(), true
		case 1:
			return Companion_AdvOutcome_.Create_Declined_(), true
		default:
			return AdvOutcome{}, false
		}
	}
}

func (_this AdvOutcome) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case AdvOutcome_Advanced:
		{
			return "ConfluxWatermark.AdvOutcome.Advanced"
		}
	case AdvOutcome_Declined:
		{
			return "ConfluxWatermark.AdvOutcome.Declined"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this AdvOutcome) Equals(other AdvOutcome) bool {
	switch _this.Get_().(type) {
	case AdvOutcome_Advanced:
		{
			_, ok := other.Get_().(AdvOutcome_Advanced)
			return ok
		}
	case AdvOutcome_Declined:
		{
			_, ok := other.Get_().(AdvOutcome_Declined)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this AdvOutcome) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(AdvOutcome)
	return ok && _this.Equals(typed)
}

func Type_AdvOutcome_() _dafny.TypeDescriptor {
	return type_AdvOutcome_{}
}

type type_AdvOutcome_ struct {
}

func (_this type_AdvOutcome_) Default() interface{} {
	return Companion_AdvOutcome_.Default()
}

func (_this type_AdvOutcome_) String() string {
	return "ConfluxWatermark.AdvOutcome"
}
func (_this AdvOutcome) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = AdvOutcome{}

// End of datatype AdvOutcome
