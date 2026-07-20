// Package ConfluxDelimited
// Dafny module ConfluxDelimited compiled into Go

package ConfluxDelimited

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
	return "ConfluxDelimited.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Intercalate(sep interface{}, xss _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xss).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else if (_dafny.IntOfUint32((xss).Cardinality())).Cmp(_dafny.One) == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (xss).Select(0).(_dafny.Sequence))
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.Companion_Sequence_.Concatenate((xss).Select(0).(_dafny.Sequence), _dafny.SeqOf(sep)))
		var _in0 interface{} = sep
		_ = _in0
		var _in1 _dafny.Sequence = (xss).Drop(1)
		_ = _in1
		sep = _in0
		xss = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Flatten(xss _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xss).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (xss).Select(0).(_dafny.Sequence))
		var _in0 _dafny.Sequence = (xss).Drop(1)
		_ = _in0
		xss = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) IndexOf(sep interface{}, xs _dafny.Sequence) _dafny.Int {
	var _0___accumulator _dafny.Int = _dafny.Zero
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return (_dafny.Zero).Plus(_0___accumulator)
	} else if _dafny.AreEqual((xs).Select(0).(interface{}), sep) {
		return (_dafny.Zero).Plus(_0___accumulator)
	} else {
		_0___accumulator = (_0___accumulator).Plus(_dafny.One)
		var _in0 interface{} = sep
		_ = _in0
		var _in1 _dafny.Sequence = (xs).Drop(1)
		_ = _in1
		sep = _in0
		xs = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Split(sep interface{}, xs _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _1_i _dafny.Int = Companion_Default___.IndexOf(sep, xs)
	_ = _1_i
	if (_1_i).Cmp(_dafny.IntOfUint32((xs).Cardinality())) == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf(xs))
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((xs).Take((_1_i).Uint32())))
		var _in0 interface{} = sep
		_ = _in0
		var _in1 _dafny.Sequence = (xs).Drop(((_1_i).Plus(_dafny.One)).Uint32())
		_ = _in1
		sep = _in0
		xs = _in1
		goto TAIL_CALL_START
	}
}

// End of class Default__
