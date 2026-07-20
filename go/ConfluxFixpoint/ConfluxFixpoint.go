// Package ConfluxFixpoint
// Dafny module ConfluxFixpoint compiled into Go

package ConfluxFixpoint

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
	return "ConfluxFixpoint.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Step(succ func(_dafny.Set) _dafny.Set, S _dafny.Set) _dafny.Set {
	return (S).Union((succ)(S))
}
func (_static *CompanionStruct_Default___) Closure(succ func(_dafny.Set) _dafny.Set, U _dafny.Set, S _dafny.Set) _dafny.Set {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _0_next _dafny.Set = Companion_Default___.Step(func(coer76 func(_dafny.Set) _dafny.Set) func(_dafny.Set) _dafny.Set {
		return func(arg85 _dafny.Set) _dafny.Set {
			return coer76(arg85)
		}
	}(succ), S)
	_ = _0_next
	if (_0_next).Equals(S) {
		return S
	} else {
		var _in0 func(_dafny.Set) _dafny.Set = succ
		_ = _in0
		var _in1 _dafny.Set = U
		_ = _in1
		var _in2 _dafny.Set = _0_next
		_ = _in2
		succ = _in0
		U = _in1
		S = _in2
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Worklist(succ func(_dafny.Set) _dafny.Set, U _dafny.Set, frontier _dafny.Set, visited _dafny.Set) _dafny.Set {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _0_f _dafny.Set = (frontier).Difference(visited)
	_ = _0_f
	if (_0_f).Equals(_dafny.SetOf()) {
		return visited
	} else {
		var _in0 func(_dafny.Set) _dafny.Set = succ
		_ = _in0
		var _in1 _dafny.Set = U
		_ = _in1
		var _in2 _dafny.Set = (succ)(_0_f)
		_ = _in2
		var _in3 _dafny.Set = (visited).Union(_0_f)
		_ = _in3
		succ = _in0
		U = _in1
		frontier = _in2
		visited = _in3
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) StepOne(A _dafny.Set) _dafny.Set {
	if (A).Contains(_dafny.Zero) {
		return _dafny.SetOf(_dafny.One)
	} else {
		return _dafny.SetOf()
	}
}
func (_static *CompanionStruct_Default___) ClosureClamp(succ func(_dafny.Set) _dafny.Set, U _dafny.Set, S _dafny.Set, fuel _dafny.Int) _dafny.Set {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _0_next _dafny.Set = ((S).Union((succ)(S))).Intersection(U)
	_ = _0_next
	if ((_0_next).Equals(S)) || ((fuel).Sign() == 0) {
		return S
	} else {
		var _in0 func(_dafny.Set) _dafny.Set = succ
		_ = _in0
		var _in1 _dafny.Set = U
		_ = _in1
		var _in2 _dafny.Set = _0_next
		_ = _in2
		var _in3 _dafny.Int = (fuel).Minus(_dafny.One)
		_ = _in3
		succ = _in0
		U = _in1
		S = _in2
		fuel = _in3
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) ClosureTotal(succ func(_dafny.Set) _dafny.Set, U _dafny.Set, S _dafny.Set) _dafny.Set {
	return Companion_Default___.ClosureClamp(func(coer77 func(_dafny.Set) _dafny.Set) func(_dafny.Set) _dafny.Set {
		return func(arg86 _dafny.Set) _dafny.Set {
			return coer77(arg86)
		}
	}(succ), U, (S).Intersection(U), (U).Cardinality())
}
func (_static *CompanionStruct_Default___) TwoHop(s _dafny.Set) _dafny.Set {
	return ((func() _dafny.Set {
		if (s).Contains(_dafny.Zero) {
			return _dafny.SetOf(_dafny.One)
		}
		return _dafny.SetOf()
	})()).Union((func() _dafny.Set {
		if (s).Contains(_dafny.One) {
			return _dafny.SetOf(_dafny.IntOfInt64(2))
		}
		return _dafny.SetOf()
	})())
}

// End of class Default__
