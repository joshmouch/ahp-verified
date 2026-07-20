// Package ConfluxDurableHistory
// Dafny module ConfluxDurableHistory compiled into Go

package ConfluxDurableHistory

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-go/Changeset"
	m_Chat "github.com/joshmouch/ahp-go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-go/ClientMain"
	m_ConfluxBatchRoute "github.com/joshmouch/ahp-go/ConfluxBatchRoute"
	m_ConfluxChannel "github.com/joshmouch/ahp-go/ConfluxChannel"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxContentAddress "github.com/joshmouch/ahp-go/ConfluxContentAddress"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxDedupe "github.com/joshmouch/ahp-go/ConfluxDedupe"
	m_ConfluxDelimited "github.com/joshmouch/ahp-go/ConfluxDelimited"
	m_ConfluxFilterKeys "github.com/joshmouch/ahp-go/ConfluxFilterKeys"
	m_ConfluxFixpoint "github.com/joshmouch/ahp-go/ConfluxFixpoint"
	m_ConfluxGenericCodec "github.com/joshmouch/ahp-go/ConfluxGenericCodec"
	m_ConfluxJoin "github.com/joshmouch/ahp-go/ConfluxJoin"
	m_ConfluxKeyedOps "github.com/joshmouch/ahp-go/ConfluxKeyedOps"
	m_ConfluxMapValues "github.com/joshmouch/ahp-go/ConfluxMapValues"
	m_ConfluxMerge "github.com/joshmouch/ahp-go/ConfluxMerge"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-go/ConfluxOrderedLog"
	m_ConfluxProduct "github.com/joshmouch/ahp-go/ConfluxProduct"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxResolve "github.com/joshmouch/ahp-go/ConfluxResolve"
	m_ConfluxResourceCeilings "github.com/joshmouch/ahp-go/ConfluxResourceCeilings"
	m_ConfluxRoute "github.com/joshmouch/ahp-go/ConfluxRoute"
	m_ConfluxSemanticGraphContract "github.com/joshmouch/ahp-go/ConfluxSemanticGraphContract"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-go/ConfluxSeqRoute"
	m_ConfluxSeqRouteMerge "github.com/joshmouch/ahp-go/ConfluxSeqRouteMerge"
	m_ConfluxStateMachine "github.com/joshmouch/ahp-go/ConfluxStateMachine"
	m_ConfluxSupervisedOperationResult "github.com/joshmouch/ahp-go/ConfluxSupervisedOperationResult"
	m_ConfluxTree "github.com/joshmouch/ahp-go/ConfluxTree"
	m_ConfluxWatermark "github.com/joshmouch/ahp-go/ConfluxWatermark"
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
var _ m_ConfluxGenericCodec.Dummy__
var _ m_ConfluxResolve.Dummy__
var _ m_ConfluxWatermark.Dummy__
var _ m_ConfluxStateMachine.Dummy__
var _ m_ConfluxProduct.Dummy__
var _ m_ConfluxJoin.Dummy__
var _ m_ConfluxDedupe.Dummy__
var _ m_ConfluxBatchRoute.Dummy__
var _ m_ConfluxSeqRouteMerge.Dummy__
var _ m_ConfluxKeyedOps.Dummy__
var _ m_ConfluxFilterKeys.Dummy__
var _ m_ConfluxChannel.Dummy__

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
	return "ConfluxDurableHistory.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Append(history _dafny.Sequence, entry interface{}) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((history), _dafny.SeqOf(entry))
}
func (_static *CompanionStruct_Default___) Replay(reducer func(interface{}, interface{}) interface{}, initial interface{}, history _dafny.Sequence) interface{} {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer90 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
		return func(arg105 interface{}, arg106 interface{}) interface{} {
			return coer90(arg105, arg106)
		}
	}(reducer), initial, (history))
}
func (_static *CompanionStruct_Default___) MakeCheckpoint(historyId _dafny.Sequence, reducerSchemaVersion _dafny.Sequence, coordinateOf func(interface{}) _dafny.Int, hashBytes func(_dafny.Sequence) _dafny.Sequence, codec StateCodec, state interface{}) ReplayCheckpoint {
	var _0_bytes _dafny.Sequence = (func(coer91 func(interface{}) _dafny.Sequence) func(interface{}) _dafny.Sequence {
		return func(arg107 interface{}) _dafny.Sequence {
			return coer91(arg107)
		}
	}((codec).Dtor_encode()))(state)
	_ = _0_bytes
	return Companion_ReplayCheckpoint_.Create_ReplayCheckpoint_(_dafny.One, historyId, (coordinateOf)(state), reducerSchemaVersion, (hashBytes)(_0_bytes), _0_bytes)
}
func (_static *CompanionStruct_Default___) RestoreCheckpoint(expectedHistoryId _dafny.Sequence, expectedReducerSchemaVersion _dafny.Sequence, coordinateOf func(interface{}) _dafny.Int, hashBytes func(_dafny.Sequence) _dafny.Sequence, codec StateCodec, checkpoint ReplayCheckpoint) m_ConfluxContract.Option {
	if (((((checkpoint).Dtor_schemaVersion()).Cmp(_dafny.One) != 0) || (!_dafny.Companion_Sequence_.Equal((checkpoint).Dtor_historyId(), expectedHistoryId))) || (!_dafny.Companion_Sequence_.Equal((checkpoint).Dtor_reducerSchemaVersion(), expectedReducerSchemaVersion))) || (!_dafny.Companion_Sequence_.Equal((checkpoint).Dtor_stateHash(), (hashBytes)((checkpoint).Dtor_stateBytes()))) {
		return m_ConfluxContract.Companion_Option_.Create_None_()
	} else {
		var _source0 m_ConfluxContract.Option = (func(coer92 func(_dafny.Sequence) m_ConfluxContract.Option) func(_dafny.Sequence) m_ConfluxContract.Option {
			return func(arg108 _dafny.Sequence) m_ConfluxContract.Option {
				return coer92(arg108)
			}
		}((codec).Dtor_decode()))((checkpoint).Dtor_stateBytes())
		_ = _source0
		{
			if _source0.Is_None() {
				return m_ConfluxContract.Companion_Option_.Create_None_()
			}
		}
		{
			var _0_state interface{} = _source0.Get_().(m_ConfluxContract.Option_Some).Value
			_ = _0_state
			if ((coordinateOf)(_0_state)).Cmp((checkpoint).Dtor_nextCoordinate()) == 0 {
				return m_ConfluxContract.Companion_Option_.Create_Some_(_0_state)
			} else {
				return m_ConfluxContract.Companion_Option_.Create_None_()
			}
		}
	}
}

// End of class Default__

// Definition of datatype History
type History struct {
	Data_History_
}

func (_this History) Get_() Data_History_ {
	return _this.Data_History_
}

type Data_History_ interface {
	isHistory()
}

type CompanionStruct_History_ struct {
}

var Companion_History_ = CompanionStruct_History_{}

type History_History struct {
	Entries _dafny.Sequence
}

func (History_History) isHistory() {}

func (CompanionStruct_History_) Create_History_(Entries _dafny.Sequence) History {
	return History{History_History{Entries}}
}

func (_this History) Is_History() bool {
	_, ok := _this.Get_().(History_History)
	return ok
}

func (CompanionStruct_History_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this History) Dtor_entries() _dafny.Sequence {
	return _this.Get_().(History_History).Entries
}

func (_this History) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case History_History:
		{
			return "ConfluxDurableHistory.History.History" + "(" + _dafny.String(data.Entries) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this History) Equals(other History) bool {
	switch data1 := _this.Get_().(type) {
	case History_History:
		{
			data2, ok := other.Get_().(History_History)
			return ok && data1.Entries.Equals(data2.Entries)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this History) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(History)
	return ok && _this.Equals(typed)
}

func Type_History_() _dafny.TypeDescriptor {
	return type_History_{}
}

type type_History_ struct {
}

func (_this type_History_) Default() interface{} {
	return Companion_History_.Default()
}

func (_this type_History_) String() string {
	return "ConfluxDurableHistory.History"
}
func (_this History) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = History{}

// End of datatype History

// Definition of datatype StateCodec
type StateCodec struct {
	Data_StateCodec_
}

func (_this StateCodec) Get_() Data_StateCodec_ {
	return _this.Data_StateCodec_
}

type Data_StateCodec_ interface {
	isStateCodec()
}

type CompanionStruct_StateCodec_ struct {
}

var Companion_StateCodec_ = CompanionStruct_StateCodec_{}

type StateCodec_StateCodec struct {
	Encode func(interface{}) _dafny.Sequence
	Decode func(_dafny.Sequence) m_ConfluxContract.Option
}

func (StateCodec_StateCodec) isStateCodec() {}

func (CompanionStruct_StateCodec_) Create_StateCodec_(Encode func(interface{}) _dafny.Sequence, Decode func(_dafny.Sequence) m_ConfluxContract.Option) StateCodec {
	return StateCodec{StateCodec_StateCodec{Encode, Decode}}
}

func (_this StateCodec) Is_StateCodec() bool {
	_, ok := _this.Get_().(StateCodec_StateCodec)
	return ok
}

func (CompanionStruct_StateCodec_) Default() StateCodec {
	return Companion_StateCodec_.Create_StateCodec_(func(interface{}) _dafny.Sequence { return _dafny.EmptySeq }, func(_dafny.Sequence) m_ConfluxContract.Option { return m_ConfluxContract.Companion_Option_.Default() })
}

func (_this StateCodec) Dtor_encode() func(interface{}) _dafny.Sequence {
	return _this.Get_().(StateCodec_StateCodec).Encode
}

func (_this StateCodec) Dtor_decode() func(_dafny.Sequence) m_ConfluxContract.Option {
	return _this.Get_().(StateCodec_StateCodec).Decode
}

func (_this StateCodec) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case StateCodec_StateCodec:
		{
			return "ConfluxDurableHistory.StateCodec.StateCodec" + "(" + _dafny.String(data.Encode) + ", " + _dafny.String(data.Decode) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this StateCodec) Equals(other StateCodec) bool {
	switch data1 := _this.Get_().(type) {
	case StateCodec_StateCodec:
		{
			data2, ok := other.Get_().(StateCodec_StateCodec)
			return ok && _dafny.AreEqual(data1.Encode, data2.Encode) && _dafny.AreEqual(data1.Decode, data2.Decode)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this StateCodec) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(StateCodec)
	return ok && _this.Equals(typed)
}

func Type_StateCodec_() _dafny.TypeDescriptor {
	return type_StateCodec_{}
}

type type_StateCodec_ struct {
}

func (_this type_StateCodec_) Default() interface{} {
	return Companion_StateCodec_.Default()
}

func (_this type_StateCodec_) String() string {
	return "ConfluxDurableHistory.StateCodec"
}
func (_this StateCodec) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = StateCodec{}

// End of datatype StateCodec

// Definition of datatype ReplayCheckpoint
type ReplayCheckpoint struct {
	Data_ReplayCheckpoint_
}

func (_this ReplayCheckpoint) Get_() Data_ReplayCheckpoint_ {
	return _this.Data_ReplayCheckpoint_
}

type Data_ReplayCheckpoint_ interface {
	isReplayCheckpoint()
}

type CompanionStruct_ReplayCheckpoint_ struct {
}

var Companion_ReplayCheckpoint_ = CompanionStruct_ReplayCheckpoint_{}

type ReplayCheckpoint_ReplayCheckpoint struct {
	SchemaVersion        _dafny.Int
	HistoryId            _dafny.Sequence
	NextCoordinate       _dafny.Int
	ReducerSchemaVersion _dafny.Sequence
	StateHash            _dafny.Sequence
	StateBytes           _dafny.Sequence
}

func (ReplayCheckpoint_ReplayCheckpoint) isReplayCheckpoint() {}

func (CompanionStruct_ReplayCheckpoint_) Create_ReplayCheckpoint_(SchemaVersion _dafny.Int, HistoryId _dafny.Sequence, NextCoordinate _dafny.Int, ReducerSchemaVersion _dafny.Sequence, StateHash _dafny.Sequence, StateBytes _dafny.Sequence) ReplayCheckpoint {
	return ReplayCheckpoint{ReplayCheckpoint_ReplayCheckpoint{SchemaVersion, HistoryId, NextCoordinate, ReducerSchemaVersion, StateHash, StateBytes}}
}

func (_this ReplayCheckpoint) Is_ReplayCheckpoint() bool {
	_, ok := _this.Get_().(ReplayCheckpoint_ReplayCheckpoint)
	return ok
}

func (CompanionStruct_ReplayCheckpoint_) Default() ReplayCheckpoint {
	return Companion_ReplayCheckpoint_.Create_ReplayCheckpoint_(_dafny.Zero, _dafny.EmptySeq, _dafny.Zero, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this ReplayCheckpoint) Dtor_schemaVersion() _dafny.Int {
	return _this.Get_().(ReplayCheckpoint_ReplayCheckpoint).SchemaVersion
}

func (_this ReplayCheckpoint) Dtor_historyId() _dafny.Sequence {
	return _this.Get_().(ReplayCheckpoint_ReplayCheckpoint).HistoryId
}

func (_this ReplayCheckpoint) Dtor_nextCoordinate() _dafny.Int {
	return _this.Get_().(ReplayCheckpoint_ReplayCheckpoint).NextCoordinate
}

func (_this ReplayCheckpoint) Dtor_reducerSchemaVersion() _dafny.Sequence {
	return _this.Get_().(ReplayCheckpoint_ReplayCheckpoint).ReducerSchemaVersion
}

func (_this ReplayCheckpoint) Dtor_stateHash() _dafny.Sequence {
	return _this.Get_().(ReplayCheckpoint_ReplayCheckpoint).StateHash
}

func (_this ReplayCheckpoint) Dtor_stateBytes() _dafny.Sequence {
	return _this.Get_().(ReplayCheckpoint_ReplayCheckpoint).StateBytes
}

func (_this ReplayCheckpoint) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ReplayCheckpoint_ReplayCheckpoint:
		{
			return "ConfluxDurableHistory.ReplayCheckpoint.ReplayCheckpoint" + "(" + _dafny.String(data.SchemaVersion) + ", " + data.HistoryId.VerbatimString(true) + ", " + _dafny.String(data.NextCoordinate) + ", " + data.ReducerSchemaVersion.VerbatimString(true) + ", " + data.StateHash.VerbatimString(true) + ", " + _dafny.String(data.StateBytes) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ReplayCheckpoint) Equals(other ReplayCheckpoint) bool {
	switch data1 := _this.Get_().(type) {
	case ReplayCheckpoint_ReplayCheckpoint:
		{
			data2, ok := other.Get_().(ReplayCheckpoint_ReplayCheckpoint)
			return ok && data1.SchemaVersion.Cmp(data2.SchemaVersion) == 0 && data1.HistoryId.Equals(data2.HistoryId) && data1.NextCoordinate.Cmp(data2.NextCoordinate) == 0 && data1.ReducerSchemaVersion.Equals(data2.ReducerSchemaVersion) && data1.StateHash.Equals(data2.StateHash) && data1.StateBytes.Equals(data2.StateBytes)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ReplayCheckpoint) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ReplayCheckpoint)
	return ok && _this.Equals(typed)
}

func Type_ReplayCheckpoint_() _dafny.TypeDescriptor {
	return type_ReplayCheckpoint_{}
}

type type_ReplayCheckpoint_ struct {
}

func (_this type_ReplayCheckpoint_) Default() interface{} {
	return Companion_ReplayCheckpoint_.Default()
}

func (_this type_ReplayCheckpoint_) String() string {
	return "ConfluxDurableHistory.ReplayCheckpoint"
}
func (_this ReplayCheckpoint) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ReplayCheckpoint{}

// End of datatype ReplayCheckpoint
