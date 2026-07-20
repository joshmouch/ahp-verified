// Package ConfluxSemanticGraphContract
// Dafny module ConfluxSemanticGraphContract compiled into Go

package ConfluxSemanticGraphContract

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
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-verified/go/ConfluxOrderedLog"
	m_ConfluxPrune "github.com/joshmouch/ahp-verified/go/ConfluxPrune"
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-verified/go/ConfluxSemanticGraphIdentity"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-verified/go/ConfluxSeqRoute"
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
	return "ConfluxSemanticGraphContract.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) CompleteWitness(c Completeness) bool {
	return !((c).Dtor_complete()) || ((_dafny.IntOfUint32(((c).Dtor_reasons()).Cardinality())).Sign() == 0)
}
func (_static *CompanionStruct_Default___) TypedRecordWellFormed(value TypedValue) bool {
	var _source0 TypedValue = value
	_ = _source0
	{
		if _source0.Is_TVList() {
			var _0_items _dafny.Sequence = _source0.Get_().(TypedValue_TVList).Items
			_ = _0_items
			return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((_0_items).Cardinality())), true, func(_forall_var_0 _dafny.Int) bool {
				var _1_i _dafny.Int
				_1_i = interface{}(_forall_var_0).(_dafny.Int)
				return !(((_1_i).Sign() != -1) && ((_1_i).Cmp(_dafny.IntOfUint32((_0_items).Cardinality())) < 0)) || (Companion_Default___.TypedRecordWellFormed((_0_items).Select((_1_i).Uint32()).(TypedValue)))
			})
		}
	}
	{
		if _source0.Is_TVRecord() {
			var _2_names _dafny.Sequence = _source0.Get_().(TypedValue_TVRecord).FieldNames
			_ = _2_names
			var _3_values _dafny.Sequence = _source0.Get_().(TypedValue_TVRecord).FieldValues
			_ = _3_values
			return (((_dafny.IntOfUint32((_2_names).Cardinality())).Cmp(_dafny.IntOfUint32((_3_values).Cardinality())) == 0) && (_dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((_2_names).Cardinality())), true, func(_forall_var_1 _dafny.Int) bool {
				var _4_i _dafny.Int
				_4_i = interface{}(_forall_var_1).(_dafny.Int)
				return _dafny.Quantifier(_dafny.IntegerRange((_4_i).Plus(_dafny.One), _dafny.IntOfUint32((_2_names).Cardinality())), true, func(_forall_var_2 _dafny.Int) bool {
					var _5_j _dafny.Int
					_5_j = interface{}(_forall_var_2).(_dafny.Int)
					return !((((_4_i).Sign() != -1) && ((_4_i).Cmp(_5_j) < 0)) && ((_5_j).Cmp(_dafny.IntOfUint32((_2_names).Cardinality())) < 0)) || (!_dafny.Companion_Sequence_.Equal((_2_names).Select((_4_i).Uint32()).(_dafny.Sequence), (_2_names).Select((_5_j).Uint32()).(_dafny.Sequence)))
				})
			}))) && (_dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((_3_values).Cardinality())), true, func(_forall_var_3 _dafny.Int) bool {
				var _6_i _dafny.Int
				_6_i = interface{}(_forall_var_3).(_dafny.Int)
				return !(((_6_i).Sign() != -1) && ((_6_i).Cmp(_dafny.IntOfUint32((_3_values).Cardinality())) < 0)) || (Companion_Default___.TypedRecordWellFormed((_3_values).Select((_6_i).Uint32()).(TypedValue)))
			}))
		}
	}
	{
		return true
	}
}
func (_static *CompanionStruct_Default___) GraphReferencesClose(graph SemanticGraphSnapshotV1) bool {
	var _pat_let_tv0 = graph
	_ = _pat_let_tv0
	var _pat_let_tv1 = graph
	_ = _pat_let_tv1
	var _pat_let_tv2 = graph
	_ = _pat_let_tv2
	var _pat_let_tv3 = graph
	_ = _pat_let_tv3
	var _0_nodeIds _dafny.Set = ((graph).Dtor_nodes()).Keys()
	_ = _0_nodeIds
	var _1_edgeIds _dafny.Set = ((graph).Dtor_edges()).Keys()
	_ = _1_edgeIds
	var _2_factIds _dafny.Set = ((graph).Dtor_facts()).Keys()
	_ = _2_factIds
	var _3_findingIds _dafny.Set = ((graph).Dtor_findings()).Keys()
	_ = _3_findingIds
	var _4_evidenceIds _dafny.Set = ((graph).Dtor_evidence()).Keys()
	_ = _4_evidenceIds
	return ((((_dafny.Quantifier((_0_nodeIds).Elements(), true, func(_forall_var_0 _dafny.Sequence) bool {
		var _5_id _dafny.Sequence
		_5_id = interface{}(_forall_var_0).(_dafny.Sequence)
		return !((_0_nodeIds).Contains(_5_id)) || (func(_pat_let464_0 SemanticNode) bool {
			return func(_6_node SemanticNode) bool {
				return ((_dafny.Companion_Sequence_.Equal((_6_node).Dtor_id(), _5_id)) && (((_6_node).Dtor_factIds()).IsSubsetOf(_2_factIds))) && (((_6_node).Dtor_evidenceIds()).IsSubsetOf(_4_evidenceIds))
			}(_pat_let464_0)
		}(((graph).Dtor_nodes()).Get(_5_id).(SemanticNode)))
	})) && (_dafny.Quantifier((_1_edgeIds).Elements(), true, func(_forall_var_1 _dafny.Sequence) bool {
		var _7_id _dafny.Sequence
		_7_id = interface{}(_forall_var_1).(_dafny.Sequence)
		return !((_1_edgeIds).Contains(_7_id)) || (func(_pat_let465_0 SemanticEdge) bool {
			return func(_8_edge SemanticEdge) bool {
				return ((((_dafny.Companion_Sequence_.Equal((_8_edge).Dtor_id(), _7_id)) && (((_pat_let_tv0).Dtor_nodes()).Contains((_8_edge).Dtor_fromId()))) && (((_pat_let_tv1).Dtor_nodes()).Contains((_8_edge).Dtor_toId()))) && (((_8_edge).Dtor_factIds()).IsSubsetOf(_2_factIds))) && (((_8_edge).Dtor_evidenceIds()).IsSubsetOf(_4_evidenceIds))
			}(_pat_let465_0)
		}(((graph).Dtor_edges()).Get(_7_id).(SemanticEdge)))
	}))) && (_dafny.Quantifier((_2_factIds).Elements(), true, func(_forall_var_2 _dafny.Sequence) bool {
		var _9_id _dafny.Sequence
		_9_id = interface{}(_forall_var_2).(_dafny.Sequence)
		return !((_2_factIds).Contains(_9_id)) || (func(_pat_let466_0 SemanticFact) bool {
			return func(_10_fact SemanticFact) bool {
				return ((((_dafny.Companion_Sequence_.Equal((_10_fact).Dtor_id(), _9_id)) && (((_pat_let_tv2).Dtor_nodes()).Contains((_10_fact).Dtor_subjectId()))) && (((_10_fact).Dtor_supportingFactIds()).IsSubsetOf(_2_factIds))) && (((_10_fact).Dtor_evidenceIds()).IsSubsetOf(_4_evidenceIds))) && (Companion_Default___.TypedRecordWellFormed((_10_fact).Dtor_value()))
			}(_pat_let466_0)
		}(((graph).Dtor_facts()).Get(_9_id).(SemanticFact)))
	}))) && (_dafny.Quantifier((_3_findingIds).Elements(), true, func(_forall_var_3 _dafny.Sequence) bool {
		var _11_id _dafny.Sequence
		_11_id = interface{}(_forall_var_3).(_dafny.Sequence)
		return !((_3_findingIds).Contains(_11_id)) || (func(_pat_let467_0 SemanticFinding) bool {
			return func(_12_finding SemanticFinding) bool {
				return (((_dafny.Companion_Sequence_.Equal((_12_finding).Dtor_id(), _11_id)) && ((_dafny.Companion_Sequence_.Equal((_12_finding).Dtor_subjectId(), _dafny.UnicodeSeqOfUtf8Bytes(""))) || (((_pat_let_tv3).Dtor_nodes()).Contains((_12_finding).Dtor_subjectId())))) && (((_12_finding).Dtor_supportingFactIds()).IsSubsetOf(_2_factIds))) && (((_12_finding).Dtor_evidenceIds()).IsSubsetOf(_4_evidenceIds))
			}(_pat_let467_0)
		}(((graph).Dtor_findings()).Get(_11_id).(SemanticFinding)))
	}))) && (_dafny.Quantifier((_4_evidenceIds).Elements(), true, func(_forall_var_4 _dafny.Sequence) bool {
		var _13_id _dafny.Sequence
		_13_id = interface{}(_forall_var_4).(_dafny.Sequence)
		return !((_4_evidenceIds).Contains(_13_id)) || (func(_pat_let468_0 EvidenceRecord) bool {
			return func(_14_item EvidenceRecord) bool {
				return (_dafny.Companion_Sequence_.Equal((_14_item).Dtor_id(), _13_id)) && (((_14_item).Dtor_parentEvidenceIds()).IsSubsetOf(_4_evidenceIds))
			}(_pat_let468_0)
		}(((graph).Dtor_evidence()).Get(_13_id).(EvidenceRecord)))
	}))
}
func (_static *CompanionStruct_Default___) GraphWellFormed(graph SemanticGraphSnapshotV1) bool {
	return ((((((_dafny.Companion_Sequence_.Equal((graph).Dtor_schemaVersion(), _dafny.UnicodeSeqOfUtf8Bytes("semantic-graph/1.0.0"))) && (!_dafny.Companion_Sequence_.Equal((graph).Dtor_graphId(), _dafny.UnicodeSeqOfUtf8Bytes("")))) && (!_dafny.Companion_Sequence_.Equal((graph).Dtor_revision(), _dafny.UnicodeSeqOfUtf8Bytes("")))) && (!_dafny.Companion_Sequence_.Equal((graph).Dtor_snapshotId(), _dafny.UnicodeSeqOfUtf8Bytes("")))) && (!_dafny.Companion_Sequence_.Equal((graph).Dtor_receiptId(), _dafny.UnicodeSeqOfUtf8Bytes("")))) && (Companion_Default___.CompleteWitness((graph).Dtor_completeness()))) && (Companion_Default___.GraphReferencesClose(graph))
}
func (_static *CompanionStruct_Default___) NonResetDeltaWellFormed(delta SemanticGraphDeltaV1) bool {
	return (((((((!((delta).Dtor_reset())) && (((delta).Dtor_resetReason()).Is_None())) && (((delta).Dtor_replacement()).Is_None())) && (((((delta).Dtor_nodes()).Dtor_added()).Keys()).IsDisjointFrom((((delta).Dtor_nodes()).Dtor_changed()).Keys()))) && (((((delta).Dtor_edges()).Dtor_added()).Keys()).IsDisjointFrom((((delta).Dtor_edges()).Dtor_changed()).Keys()))) && (((((delta).Dtor_facts()).Dtor_added()).Keys()).IsDisjointFrom((((delta).Dtor_facts()).Dtor_changed()).Keys()))) && (((((delta).Dtor_findings()).Dtor_added()).Keys()).IsDisjointFrom((((delta).Dtor_findings()).Dtor_changed()).Keys()))) && (((((delta).Dtor_evidence()).Dtor_added()).Keys()).IsDisjointFrom((((delta).Dtor_evidence()).Dtor_changed()).Keys()))
}
func (_static *CompanionStruct_Default___) ResetDeltaWellFormed(delta SemanticGraphDeltaV1) bool {
	return ((((((((delta).Dtor_reset()) && (((delta).Dtor_resetReason()).Is_Some())) && (((delta).Dtor_replacement()).Is_Some())) && (((((((delta).Dtor_nodes()).Dtor_added()).Cardinality()).Plus((((delta).Dtor_nodes()).Dtor_changed()).Cardinality())).Plus((((delta).Dtor_nodes()).Dtor_removed()).Cardinality())).Sign() == 0)) && (((((((delta).Dtor_edges()).Dtor_added()).Cardinality()).Plus((((delta).Dtor_edges()).Dtor_changed()).Cardinality())).Plus((((delta).Dtor_edges()).Dtor_removed()).Cardinality())).Sign() == 0)) && (((((((delta).Dtor_facts()).Dtor_added()).Cardinality()).Plus((((delta).Dtor_facts()).Dtor_changed()).Cardinality())).Plus((((delta).Dtor_facts()).Dtor_removed()).Cardinality())).Sign() == 0)) && (((((((delta).Dtor_findings()).Dtor_added()).Cardinality()).Plus((((delta).Dtor_findings()).Dtor_changed()).Cardinality())).Plus((((delta).Dtor_findings()).Dtor_removed()).Cardinality())).Sign() == 0)) && (((((((delta).Dtor_evidence()).Dtor_added()).Cardinality()).Plus((((delta).Dtor_evidence()).Dtor_changed()).Cardinality())).Plus((((delta).Dtor_evidence()).Dtor_removed()).Cardinality())).Sign() == 0)
}
func (_static *CompanionStruct_Default___) DeltaWellFormed(delta SemanticGraphDeltaV1) bool {
	return (((((_dafny.Companion_Sequence_.Equal((delta).Dtor_schemaVersion(), _dafny.UnicodeSeqOfUtf8Bytes("semantic-graph-delta/1.0.0"))) && (!_dafny.Companion_Sequence_.Equal((delta).Dtor_graphId(), _dafny.UnicodeSeqOfUtf8Bytes("")))) && (!_dafny.Companion_Sequence_.Equal((delta).Dtor_fromRevision(), _dafny.UnicodeSeqOfUtf8Bytes("")))) && (!_dafny.Companion_Sequence_.Equal((delta).Dtor_toRevision(), _dafny.UnicodeSeqOfUtf8Bytes("")))) && (!_dafny.Companion_Sequence_.Equal((delta).Dtor_receiptId(), _dafny.UnicodeSeqOfUtf8Bytes("")))) && ((func() bool {
		if (delta).Dtor_reset() {
			return Companion_Default___.ResetDeltaWellFormed(delta)
		}
		return Companion_Default___.NonResetDeltaWellFormed(delta)
	})())
}
func (_static *CompanionStruct_Default___) DeltaApplies(before SemanticGraphSnapshotV1, delta SemanticGraphDeltaV1, after SemanticGraphSnapshotV1) bool {
	return ((((((((Companion_Default___.GraphWellFormed(before)) && (Companion_Default___.GraphWellFormed(after))) && (Companion_Default___.DeltaWellFormed(delta))) && (!((delta).Dtor_reset()))) && (_dafny.Companion_Sequence_.Equal((before).Dtor_graphId(), (delta).Dtor_graphId()))) && (_dafny.Companion_Sequence_.Equal((after).Dtor_graphId(), (delta).Dtor_graphId()))) && (_dafny.Companion_Sequence_.Equal((before).Dtor_revision(), (delta).Dtor_fromRevision()))) && (_dafny.Companion_Sequence_.Equal((after).Dtor_revision(), (delta).Dtor_toRevision()))) && (_dafny.Companion_Sequence_.Equal((after).Dtor_previousRevision(), (before).Dtor_revision()))
}

// End of class Default__

// Definition of datatype SelectorKind
type SelectorKind struct {
	Data_SelectorKind_
}

func (_this SelectorKind) Get_() Data_SelectorKind_ {
	return _this.Data_SelectorKind_
}

type Data_SelectorKind_ interface {
	isSelectorKind()
}

type CompanionStruct_SelectorKind_ struct {
}

var Companion_SelectorKind_ = CompanionStruct_SelectorKind_{}

type SelectorKind_Workspace struct {
}

func (SelectorKind_Workspace) isSelectorKind() {}

func (CompanionStruct_SelectorKind_) Create_Workspace_() SelectorKind {
	return SelectorKind{SelectorKind_Workspace{}}
}

func (_this SelectorKind) Is_Workspace() bool {
	_, ok := _this.Get_().(SelectorKind_Workspace)
	return ok
}

type SelectorKind_RepositorySet struct {
}

func (SelectorKind_RepositorySet) isSelectorKind() {}

func (CompanionStruct_SelectorKind_) Create_RepositorySet_() SelectorKind {
	return SelectorKind{SelectorKind_RepositorySet{}}
}

func (_this SelectorKind) Is_RepositorySet() bool {
	_, ok := _this.Get_().(SelectorKind_RepositorySet)
	return ok
}

type SelectorKind_Repository struct {
}

func (SelectorKind_Repository) isSelectorKind() {}

func (CompanionStruct_SelectorKind_) Create_Repository_() SelectorKind {
	return SelectorKind{SelectorKind_Repository{}}
}

func (_this SelectorKind) Is_Repository() bool {
	_, ok := _this.Get_().(SelectorKind_Repository)
	return ok
}

type SelectorKind_Project struct {
}

func (SelectorKind_Project) isSelectorKind() {}

func (CompanionStruct_SelectorKind_) Create_Project_() SelectorKind {
	return SelectorKind{SelectorKind_Project{}}
}

func (_this SelectorKind) Is_Project() bool {
	_, ok := _this.Get_().(SelectorKind_Project)
	return ok
}

type SelectorKind_File struct {
}

func (SelectorKind_File) isSelectorKind() {}

func (CompanionStruct_SelectorKind_) Create_File_() SelectorKind {
	return SelectorKind{SelectorKind_File{}}
}

func (_this SelectorKind) Is_File() bool {
	_, ok := _this.Get_().(SelectorKind_File)
	return ok
}

type SelectorKind_Subject struct {
}

func (SelectorKind_Subject) isSelectorKind() {}

func (CompanionStruct_SelectorKind_) Create_Subject_() SelectorKind {
	return SelectorKind{SelectorKind_Subject{}}
}

func (_this SelectorKind) Is_Subject() bool {
	_, ok := _this.Get_().(SelectorKind_Subject)
	return ok
}

func (CompanionStruct_SelectorKind_) Default() SelectorKind {
	return Companion_SelectorKind_.Create_Workspace_()
}

func (_ CompanionStruct_SelectorKind_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_SelectorKind_.Create_Workspace_(), true
		case 1:
			return Companion_SelectorKind_.Create_RepositorySet_(), true
		case 2:
			return Companion_SelectorKind_.Create_Repository_(), true
		case 3:
			return Companion_SelectorKind_.Create_Project_(), true
		case 4:
			return Companion_SelectorKind_.Create_File_(), true
		case 5:
			return Companion_SelectorKind_.Create_Subject_(), true
		default:
			return SelectorKind{}, false
		}
	}
}

func (_this SelectorKind) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case SelectorKind_Workspace:
		{
			return "ConfluxSemanticGraphContract.SelectorKind.Workspace"
		}
	case SelectorKind_RepositorySet:
		{
			return "ConfluxSemanticGraphContract.SelectorKind.RepositorySet"
		}
	case SelectorKind_Repository:
		{
			return "ConfluxSemanticGraphContract.SelectorKind.Repository"
		}
	case SelectorKind_Project:
		{
			return "ConfluxSemanticGraphContract.SelectorKind.Project"
		}
	case SelectorKind_File:
		{
			return "ConfluxSemanticGraphContract.SelectorKind.File"
		}
	case SelectorKind_Subject:
		{
			return "ConfluxSemanticGraphContract.SelectorKind.Subject"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SelectorKind) Equals(other SelectorKind) bool {
	switch _this.Get_().(type) {
	case SelectorKind_Workspace:
		{
			_, ok := other.Get_().(SelectorKind_Workspace)
			return ok
		}
	case SelectorKind_RepositorySet:
		{
			_, ok := other.Get_().(SelectorKind_RepositorySet)
			return ok
		}
	case SelectorKind_Repository:
		{
			_, ok := other.Get_().(SelectorKind_Repository)
			return ok
		}
	case SelectorKind_Project:
		{
			_, ok := other.Get_().(SelectorKind_Project)
			return ok
		}
	case SelectorKind_File:
		{
			_, ok := other.Get_().(SelectorKind_File)
			return ok
		}
	case SelectorKind_Subject:
		{
			_, ok := other.Get_().(SelectorKind_Subject)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SelectorKind) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SelectorKind)
	return ok && _this.Equals(typed)
}

func Type_SelectorKind_() _dafny.TypeDescriptor {
	return type_SelectorKind_{}
}

type type_SelectorKind_ struct {
}

func (_this type_SelectorKind_) Default() interface{} {
	return Companion_SelectorKind_.Default()
}

func (_this type_SelectorKind_) String() string {
	return "ConfluxSemanticGraphContract.SelectorKind"
}
func (_this SelectorKind) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SelectorKind{}

// End of datatype SelectorKind

// Definition of datatype SemanticStatus
type SemanticStatus struct {
	Data_SemanticStatus_
}

func (_this SemanticStatus) Get_() Data_SemanticStatus_ {
	return _this.Data_SemanticStatus_
}

type Data_SemanticStatus_ interface {
	isSemanticStatus()
}

type CompanionStruct_SemanticStatus_ struct {
}

var Companion_SemanticStatus_ = CompanionStruct_SemanticStatus_{}

type SemanticStatus_Complete struct {
}

func (SemanticStatus_Complete) isSemanticStatus() {}

func (CompanionStruct_SemanticStatus_) Create_Complete_() SemanticStatus {
	return SemanticStatus{SemanticStatus_Complete{}}
}

func (_this SemanticStatus) Is_Complete() bool {
	_, ok := _this.Get_().(SemanticStatus_Complete)
	return ok
}

type SemanticStatus_Failed struct {
}

func (SemanticStatus_Failed) isSemanticStatus() {}

func (CompanionStruct_SemanticStatus_) Create_Failed_() SemanticStatus {
	return SemanticStatus{SemanticStatus_Failed{}}
}

func (_this SemanticStatus) Is_Failed() bool {
	_, ok := _this.Get_().(SemanticStatus_Failed)
	return ok
}

type SemanticStatus_Unavailable struct {
}

func (SemanticStatus_Unavailable) isSemanticStatus() {}

func (CompanionStruct_SemanticStatus_) Create_Unavailable_() SemanticStatus {
	return SemanticStatus{SemanticStatus_Unavailable{}}
}

func (_this SemanticStatus) Is_Unavailable() bool {
	_, ok := _this.Get_().(SemanticStatus_Unavailable)
	return ok
}

type SemanticStatus_Cancelled struct {
}

func (SemanticStatus_Cancelled) isSemanticStatus() {}

func (CompanionStruct_SemanticStatus_) Create_Cancelled_() SemanticStatus {
	return SemanticStatus{SemanticStatus_Cancelled{}}
}

func (_this SemanticStatus) Is_Cancelled() bool {
	_, ok := _this.Get_().(SemanticStatus_Cancelled)
	return ok
}

type SemanticStatus_Stale struct {
}

func (SemanticStatus_Stale) isSemanticStatus() {}

func (CompanionStruct_SemanticStatus_) Create_Stale_() SemanticStatus {
	return SemanticStatus{SemanticStatus_Stale{}}
}

func (_this SemanticStatus) Is_Stale() bool {
	_, ok := _this.Get_().(SemanticStatus_Stale)
	return ok
}

func (CompanionStruct_SemanticStatus_) Default() SemanticStatus {
	return Companion_SemanticStatus_.Create_Complete_()
}

func (_ CompanionStruct_SemanticStatus_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_SemanticStatus_.Create_Complete_(), true
		case 1:
			return Companion_SemanticStatus_.Create_Failed_(), true
		case 2:
			return Companion_SemanticStatus_.Create_Unavailable_(), true
		case 3:
			return Companion_SemanticStatus_.Create_Cancelled_(), true
		case 4:
			return Companion_SemanticStatus_.Create_Stale_(), true
		default:
			return SemanticStatus{}, false
		}
	}
}

func (_this SemanticStatus) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case SemanticStatus_Complete:
		{
			return "ConfluxSemanticGraphContract.SemanticStatus.Complete"
		}
	case SemanticStatus_Failed:
		{
			return "ConfluxSemanticGraphContract.SemanticStatus.Failed"
		}
	case SemanticStatus_Unavailable:
		{
			return "ConfluxSemanticGraphContract.SemanticStatus.Unavailable"
		}
	case SemanticStatus_Cancelled:
		{
			return "ConfluxSemanticGraphContract.SemanticStatus.Cancelled"
		}
	case SemanticStatus_Stale:
		{
			return "ConfluxSemanticGraphContract.SemanticStatus.Stale"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SemanticStatus) Equals(other SemanticStatus) bool {
	switch _this.Get_().(type) {
	case SemanticStatus_Complete:
		{
			_, ok := other.Get_().(SemanticStatus_Complete)
			return ok
		}
	case SemanticStatus_Failed:
		{
			_, ok := other.Get_().(SemanticStatus_Failed)
			return ok
		}
	case SemanticStatus_Unavailable:
		{
			_, ok := other.Get_().(SemanticStatus_Unavailable)
			return ok
		}
	case SemanticStatus_Cancelled:
		{
			_, ok := other.Get_().(SemanticStatus_Cancelled)
			return ok
		}
	case SemanticStatus_Stale:
		{
			_, ok := other.Get_().(SemanticStatus_Stale)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SemanticStatus) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SemanticStatus)
	return ok && _this.Equals(typed)
}

func Type_SemanticStatus_() _dafny.TypeDescriptor {
	return type_SemanticStatus_{}
}

type type_SemanticStatus_ struct {
}

func (_this type_SemanticStatus_) Default() interface{} {
	return Companion_SemanticStatus_.Default()
}

func (_this type_SemanticStatus_) String() string {
	return "ConfluxSemanticGraphContract.SemanticStatus"
}
func (_this SemanticStatus) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SemanticStatus{}

// End of datatype SemanticStatus

// Definition of datatype Availability
type Availability struct {
	Data_Availability_
}

func (_this Availability) Get_() Data_Availability_ {
	return _this.Data_Availability_
}

type Data_Availability_ interface {
	isAvailability()
}

type CompanionStruct_Availability_ struct {
}

var Companion_Availability_ = CompanionStruct_Availability_{}

type Availability_Available struct {
}

func (Availability_Available) isAvailability() {}

func (CompanionStruct_Availability_) Create_Available_() Availability {
	return Availability{Availability_Available{}}
}

func (_this Availability) Is_Available() bool {
	_, ok := _this.Get_().(Availability_Available)
	return ok
}

type Availability_Partial struct {
}

func (Availability_Partial) isAvailability() {}

func (CompanionStruct_Availability_) Create_Partial_() Availability {
	return Availability{Availability_Partial{}}
}

func (_this Availability) Is_Partial() bool {
	_, ok := _this.Get_().(Availability_Partial)
	return ok
}

type Availability_CapabilityUnavailable struct {
}

func (Availability_CapabilityUnavailable) isAvailability() {}

func (CompanionStruct_Availability_) Create_CapabilityUnavailable_() Availability {
	return Availability{Availability_CapabilityUnavailable{}}
}

func (_this Availability) Is_CapabilityUnavailable() bool {
	_, ok := _this.Get_().(Availability_CapabilityUnavailable)
	return ok
}

func (CompanionStruct_Availability_) Default() Availability {
	return Companion_Availability_.Create_Available_()
}

func (_ CompanionStruct_Availability_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_Availability_.Create_Available_(), true
		case 1:
			return Companion_Availability_.Create_Partial_(), true
		case 2:
			return Companion_Availability_.Create_CapabilityUnavailable_(), true
		default:
			return Availability{}, false
		}
	}
}

func (_this Availability) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case Availability_Available:
		{
			return "ConfluxSemanticGraphContract.Availability.Available"
		}
	case Availability_Partial:
		{
			return "ConfluxSemanticGraphContract.Availability.Partial"
		}
	case Availability_CapabilityUnavailable:
		{
			return "ConfluxSemanticGraphContract.Availability.CapabilityUnavailable"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Availability) Equals(other Availability) bool {
	switch _this.Get_().(type) {
	case Availability_Available:
		{
			_, ok := other.Get_().(Availability_Available)
			return ok
		}
	case Availability_Partial:
		{
			_, ok := other.Get_().(Availability_Partial)
			return ok
		}
	case Availability_CapabilityUnavailable:
		{
			_, ok := other.Get_().(Availability_CapabilityUnavailable)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Availability) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Availability)
	return ok && _this.Equals(typed)
}

func Type_Availability_() _dafny.TypeDescriptor {
	return type_Availability_{}
}

type type_Availability_ struct {
}

func (_this type_Availability_) Default() interface{} {
	return Companion_Availability_.Default()
}

func (_this type_Availability_) String() string {
	return "ConfluxSemanticGraphContract.Availability"
}
func (_this Availability) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Availability{}

// End of datatype Availability

// Definition of datatype CapabilityKind
type CapabilityKind struct {
	Data_CapabilityKind_
}

func (_this CapabilityKind) Get_() Data_CapabilityKind_ {
	return _this.Data_CapabilityKind_
}

type Data_CapabilityKind_ interface {
	isCapabilityKind()
}

type CompanionStruct_CapabilityKind_ struct {
}

var Companion_CapabilityKind_ = CompanionStruct_CapabilityKind_{}

type CapabilityKind_Transport struct {
}

func (CapabilityKind_Transport) isCapabilityKind() {}

func (CompanionStruct_CapabilityKind_) Create_Transport_() CapabilityKind {
	return CapabilityKind{CapabilityKind_Transport{}}
}

func (_this CapabilityKind) Is_Transport() bool {
	_, ok := _this.Get_().(CapabilityKind_Transport)
	return ok
}

type CapabilityKind_Snapshot struct {
}

func (CapabilityKind_Snapshot) isCapabilityKind() {}

func (CompanionStruct_CapabilityKind_) Create_Snapshot_() CapabilityKind {
	return CapabilityKind{CapabilityKind_Snapshot{}}
}

func (_this CapabilityKind) Is_Snapshot() bool {
	_, ok := _this.Get_().(CapabilityKind_Snapshot)
	return ok
}

type CapabilityKind_View struct {
}

func (CapabilityKind_View) isCapabilityKind() {}

func (CompanionStruct_CapabilityKind_) Create_View_() CapabilityKind {
	return CapabilityKind{CapabilityKind_View{}}
}

func (_this CapabilityKind) Is_View() bool {
	_, ok := _this.Get_().(CapabilityKind_View)
	return ok
}

type CapabilityKind_Analysis struct {
}

func (CapabilityKind_Analysis) isCapabilityKind() {}

func (CompanionStruct_CapabilityKind_) Create_Analysis_() CapabilityKind {
	return CapabilityKind{CapabilityKind_Analysis{}}
}

func (_this CapabilityKind) Is_Analysis() bool {
	_, ok := _this.Get_().(CapabilityKind_Analysis)
	return ok
}

type CapabilityKind_Projection struct {
}

func (CapabilityKind_Projection) isCapabilityKind() {}

func (CompanionStruct_CapabilityKind_) Create_Projection_() CapabilityKind {
	return CapabilityKind{CapabilityKind_Projection{}}
}

func (_this CapabilityKind) Is_Projection() bool {
	_, ok := _this.Get_().(CapabilityKind_Projection)
	return ok
}

type CapabilityKind_Cache struct {
}

func (CapabilityKind_Cache) isCapabilityKind() {}

func (CompanionStruct_CapabilityKind_) Create_Cache_() CapabilityKind {
	return CapabilityKind{CapabilityKind_Cache{}}
}

func (_this CapabilityKind) Is_Cache() bool {
	_, ok := _this.Get_().(CapabilityKind_Cache)
	return ok
}

type CapabilityKind_Lens struct {
}

func (CapabilityKind_Lens) isCapabilityKind() {}

func (CompanionStruct_CapabilityKind_) Create_Lens_() CapabilityKind {
	return CapabilityKind{CapabilityKind_Lens{}}
}

func (_this CapabilityKind) Is_Lens() bool {
	_, ok := _this.Get_().(CapabilityKind_Lens)
	return ok
}

func (CompanionStruct_CapabilityKind_) Default() CapabilityKind {
	return Companion_CapabilityKind_.Create_Transport_()
}

func (_ CompanionStruct_CapabilityKind_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_CapabilityKind_.Create_Transport_(), true
		case 1:
			return Companion_CapabilityKind_.Create_Snapshot_(), true
		case 2:
			return Companion_CapabilityKind_.Create_View_(), true
		case 3:
			return Companion_CapabilityKind_.Create_Analysis_(), true
		case 4:
			return Companion_CapabilityKind_.Create_Projection_(), true
		case 5:
			return Companion_CapabilityKind_.Create_Cache_(), true
		case 6:
			return Companion_CapabilityKind_.Create_Lens_(), true
		default:
			return CapabilityKind{}, false
		}
	}
}

func (_this CapabilityKind) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case CapabilityKind_Transport:
		{
			return "ConfluxSemanticGraphContract.CapabilityKind.Transport"
		}
	case CapabilityKind_Snapshot:
		{
			return "ConfluxSemanticGraphContract.CapabilityKind.Snapshot"
		}
	case CapabilityKind_View:
		{
			return "ConfluxSemanticGraphContract.CapabilityKind.View"
		}
	case CapabilityKind_Analysis:
		{
			return "ConfluxSemanticGraphContract.CapabilityKind.Analysis"
		}
	case CapabilityKind_Projection:
		{
			return "ConfluxSemanticGraphContract.CapabilityKind.Projection"
		}
	case CapabilityKind_Cache:
		{
			return "ConfluxSemanticGraphContract.CapabilityKind.Cache"
		}
	case CapabilityKind_Lens:
		{
			return "ConfluxSemanticGraphContract.CapabilityKind.Lens"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CapabilityKind) Equals(other CapabilityKind) bool {
	switch _this.Get_().(type) {
	case CapabilityKind_Transport:
		{
			_, ok := other.Get_().(CapabilityKind_Transport)
			return ok
		}
	case CapabilityKind_Snapshot:
		{
			_, ok := other.Get_().(CapabilityKind_Snapshot)
			return ok
		}
	case CapabilityKind_View:
		{
			_, ok := other.Get_().(CapabilityKind_View)
			return ok
		}
	case CapabilityKind_Analysis:
		{
			_, ok := other.Get_().(CapabilityKind_Analysis)
			return ok
		}
	case CapabilityKind_Projection:
		{
			_, ok := other.Get_().(CapabilityKind_Projection)
			return ok
		}
	case CapabilityKind_Cache:
		{
			_, ok := other.Get_().(CapabilityKind_Cache)
			return ok
		}
	case CapabilityKind_Lens:
		{
			_, ok := other.Get_().(CapabilityKind_Lens)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CapabilityKind) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CapabilityKind)
	return ok && _this.Equals(typed)
}

func Type_CapabilityKind_() _dafny.TypeDescriptor {
	return type_CapabilityKind_{}
}

type type_CapabilityKind_ struct {
}

func (_this type_CapabilityKind_) Default() interface{} {
	return Companion_CapabilityKind_.Default()
}

func (_this type_CapabilityKind_) String() string {
	return "ConfluxSemanticGraphContract.CapabilityKind"
}
func (_this CapabilityKind) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CapabilityKind{}

// End of datatype CapabilityKind

// Definition of datatype FactLayer
type FactLayer struct {
	Data_FactLayer_
}

func (_this FactLayer) Get_() Data_FactLayer_ {
	return _this.Data_FactLayer_
}

type Data_FactLayer_ interface {
	isFactLayer()
}

type CompanionStruct_FactLayer_ struct {
}

var Companion_FactLayer_ = CompanionStruct_FactLayer_{}

type FactLayer_Declared struct {
}

func (FactLayer_Declared) isFactLayer() {}

func (CompanionStruct_FactLayer_) Create_Declared_() FactLayer {
	return FactLayer{FactLayer_Declared{}}
}

func (_this FactLayer) Is_Declared() bool {
	_, ok := _this.Get_().(FactLayer_Declared)
	return ok
}

type FactLayer_Discovered struct {
}

func (FactLayer_Discovered) isFactLayer() {}

func (CompanionStruct_FactLayer_) Create_Discovered_() FactLayer {
	return FactLayer{FactLayer_Discovered{}}
}

func (_this FactLayer) Is_Discovered() bool {
	_, ok := _this.Get_().(FactLayer_Discovered)
	return ok
}

type FactLayer_Observed struct {
}

func (FactLayer_Observed) isFactLayer() {}

func (CompanionStruct_FactLayer_) Create_Observed_() FactLayer {
	return FactLayer{FactLayer_Observed{}}
}

func (_this FactLayer) Is_Observed() bool {
	_, ok := _this.Get_().(FactLayer_Observed)
	return ok
}

type FactLayer_Derived struct {
}

func (FactLayer_Derived) isFactLayer() {}

func (CompanionStruct_FactLayer_) Create_Derived_() FactLayer {
	return FactLayer{FactLayer_Derived{}}
}

func (_this FactLayer) Is_Derived() bool {
	_, ok := _this.Get_().(FactLayer_Derived)
	return ok
}

type FactLayer_Policy struct {
}

func (FactLayer_Policy) isFactLayer() {}

func (CompanionStruct_FactLayer_) Create_Policy_() FactLayer {
	return FactLayer{FactLayer_Policy{}}
}

func (_this FactLayer) Is_Policy() bool {
	_, ok := _this.Get_().(FactLayer_Policy)
	return ok
}

func (CompanionStruct_FactLayer_) Default() FactLayer {
	return Companion_FactLayer_.Create_Declared_()
}

func (_ CompanionStruct_FactLayer_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_FactLayer_.Create_Declared_(), true
		case 1:
			return Companion_FactLayer_.Create_Discovered_(), true
		case 2:
			return Companion_FactLayer_.Create_Observed_(), true
		case 3:
			return Companion_FactLayer_.Create_Derived_(), true
		case 4:
			return Companion_FactLayer_.Create_Policy_(), true
		default:
			return FactLayer{}, false
		}
	}
}

func (_this FactLayer) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case FactLayer_Declared:
		{
			return "ConfluxSemanticGraphContract.FactLayer.Declared"
		}
	case FactLayer_Discovered:
		{
			return "ConfluxSemanticGraphContract.FactLayer.Discovered"
		}
	case FactLayer_Observed:
		{
			return "ConfluxSemanticGraphContract.FactLayer.Observed"
		}
	case FactLayer_Derived:
		{
			return "ConfluxSemanticGraphContract.FactLayer.Derived"
		}
	case FactLayer_Policy:
		{
			return "ConfluxSemanticGraphContract.FactLayer.Policy"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this FactLayer) Equals(other FactLayer) bool {
	switch _this.Get_().(type) {
	case FactLayer_Declared:
		{
			_, ok := other.Get_().(FactLayer_Declared)
			return ok
		}
	case FactLayer_Discovered:
		{
			_, ok := other.Get_().(FactLayer_Discovered)
			return ok
		}
	case FactLayer_Observed:
		{
			_, ok := other.Get_().(FactLayer_Observed)
			return ok
		}
	case FactLayer_Derived:
		{
			_, ok := other.Get_().(FactLayer_Derived)
			return ok
		}
	case FactLayer_Policy:
		{
			_, ok := other.Get_().(FactLayer_Policy)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this FactLayer) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(FactLayer)
	return ok && _this.Equals(typed)
}

func Type_FactLayer_() _dafny.TypeDescriptor {
	return type_FactLayer_{}
}

type type_FactLayer_ struct {
}

func (_this type_FactLayer_) Default() interface{} {
	return Companion_FactLayer_.Default()
}

func (_this type_FactLayer_) String() string {
	return "ConfluxSemanticGraphContract.FactLayer"
}
func (_this FactLayer) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = FactLayer{}

// End of datatype FactLayer

// Definition of datatype FindingSeverity
type FindingSeverity struct {
	Data_FindingSeverity_
}

func (_this FindingSeverity) Get_() Data_FindingSeverity_ {
	return _this.Data_FindingSeverity_
}

type Data_FindingSeverity_ interface {
	isFindingSeverity()
}

type CompanionStruct_FindingSeverity_ struct {
}

var Companion_FindingSeverity_ = CompanionStruct_FindingSeverity_{}

type FindingSeverity_Error struct {
}

func (FindingSeverity_Error) isFindingSeverity() {}

func (CompanionStruct_FindingSeverity_) Create_Error_() FindingSeverity {
	return FindingSeverity{FindingSeverity_Error{}}
}

func (_this FindingSeverity) Is_Error() bool {
	_, ok := _this.Get_().(FindingSeverity_Error)
	return ok
}

type FindingSeverity_Warning struct {
}

func (FindingSeverity_Warning) isFindingSeverity() {}

func (CompanionStruct_FindingSeverity_) Create_Warning_() FindingSeverity {
	return FindingSeverity{FindingSeverity_Warning{}}
}

func (_this FindingSeverity) Is_Warning() bool {
	_, ok := _this.Get_().(FindingSeverity_Warning)
	return ok
}

type FindingSeverity_Information struct {
}

func (FindingSeverity_Information) isFindingSeverity() {}

func (CompanionStruct_FindingSeverity_) Create_Information_() FindingSeverity {
	return FindingSeverity{FindingSeverity_Information{}}
}

func (_this FindingSeverity) Is_Information() bool {
	_, ok := _this.Get_().(FindingSeverity_Information)
	return ok
}

type FindingSeverity_Hint struct {
}

func (FindingSeverity_Hint) isFindingSeverity() {}

func (CompanionStruct_FindingSeverity_) Create_Hint_() FindingSeverity {
	return FindingSeverity{FindingSeverity_Hint{}}
}

func (_this FindingSeverity) Is_Hint() bool {
	_, ok := _this.Get_().(FindingSeverity_Hint)
	return ok
}

func (CompanionStruct_FindingSeverity_) Default() FindingSeverity {
	return Companion_FindingSeverity_.Create_Error_()
}

func (_ CompanionStruct_FindingSeverity_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_FindingSeverity_.Create_Error_(), true
		case 1:
			return Companion_FindingSeverity_.Create_Warning_(), true
		case 2:
			return Companion_FindingSeverity_.Create_Information_(), true
		case 3:
			return Companion_FindingSeverity_.Create_Hint_(), true
		default:
			return FindingSeverity{}, false
		}
	}
}

func (_this FindingSeverity) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case FindingSeverity_Error:
		{
			return "ConfluxSemanticGraphContract.FindingSeverity.Error"
		}
	case FindingSeverity_Warning:
		{
			return "ConfluxSemanticGraphContract.FindingSeverity.Warning"
		}
	case FindingSeverity_Information:
		{
			return "ConfluxSemanticGraphContract.FindingSeverity.Information"
		}
	case FindingSeverity_Hint:
		{
			return "ConfluxSemanticGraphContract.FindingSeverity.Hint"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this FindingSeverity) Equals(other FindingSeverity) bool {
	switch _this.Get_().(type) {
	case FindingSeverity_Error:
		{
			_, ok := other.Get_().(FindingSeverity_Error)
			return ok
		}
	case FindingSeverity_Warning:
		{
			_, ok := other.Get_().(FindingSeverity_Warning)
			return ok
		}
	case FindingSeverity_Information:
		{
			_, ok := other.Get_().(FindingSeverity_Information)
			return ok
		}
	case FindingSeverity_Hint:
		{
			_, ok := other.Get_().(FindingSeverity_Hint)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this FindingSeverity) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(FindingSeverity)
	return ok && _this.Equals(typed)
}

func Type_FindingSeverity_() _dafny.TypeDescriptor {
	return type_FindingSeverity_{}
}

type type_FindingSeverity_ struct {
}

func (_this type_FindingSeverity_) Default() interface{} {
	return Companion_FindingSeverity_.Default()
}

func (_this type_FindingSeverity_) String() string {
	return "ConfluxSemanticGraphContract.FindingSeverity"
}
func (_this FindingSeverity) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = FindingSeverity{}

// End of datatype FindingSeverity

// Definition of datatype FindingStatus
type FindingStatus struct {
	Data_FindingStatus_
}

func (_this FindingStatus) Get_() Data_FindingStatus_ {
	return _this.Data_FindingStatus_
}

type Data_FindingStatus_ interface {
	isFindingStatus()
}

type CompanionStruct_FindingStatus_ struct {
}

var Companion_FindingStatus_ = CompanionStruct_FindingStatus_{}

type FindingStatus_Open struct {
}

func (FindingStatus_Open) isFindingStatus() {}

func (CompanionStruct_FindingStatus_) Create_Open_() FindingStatus {
	return FindingStatus{FindingStatus_Open{}}
}

func (_this FindingStatus) Is_Open() bool {
	_, ok := _this.Get_().(FindingStatus_Open)
	return ok
}

type FindingStatus_Resolved struct {
}

func (FindingStatus_Resolved) isFindingStatus() {}

func (CompanionStruct_FindingStatus_) Create_Resolved_() FindingStatus {
	return FindingStatus{FindingStatus_Resolved{}}
}

func (_this FindingStatus) Is_Resolved() bool {
	_, ok := _this.Get_().(FindingStatus_Resolved)
	return ok
}

type FindingStatus_FindingUnavailable struct {
}

func (FindingStatus_FindingUnavailable) isFindingStatus() {}

func (CompanionStruct_FindingStatus_) Create_FindingUnavailable_() FindingStatus {
	return FindingStatus{FindingStatus_FindingUnavailable{}}
}

func (_this FindingStatus) Is_FindingUnavailable() bool {
	_, ok := _this.Get_().(FindingStatus_FindingUnavailable)
	return ok
}

func (CompanionStruct_FindingStatus_) Default() FindingStatus {
	return Companion_FindingStatus_.Create_Open_()
}

func (_ CompanionStruct_FindingStatus_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_FindingStatus_.Create_Open_(), true
		case 1:
			return Companion_FindingStatus_.Create_Resolved_(), true
		case 2:
			return Companion_FindingStatus_.Create_FindingUnavailable_(), true
		default:
			return FindingStatus{}, false
		}
	}
}

func (_this FindingStatus) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case FindingStatus_Open:
		{
			return "ConfluxSemanticGraphContract.FindingStatus.Open"
		}
	case FindingStatus_Resolved:
		{
			return "ConfluxSemanticGraphContract.FindingStatus.Resolved"
		}
	case FindingStatus_FindingUnavailable:
		{
			return "ConfluxSemanticGraphContract.FindingStatus.FindingUnavailable"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this FindingStatus) Equals(other FindingStatus) bool {
	switch _this.Get_().(type) {
	case FindingStatus_Open:
		{
			_, ok := other.Get_().(FindingStatus_Open)
			return ok
		}
	case FindingStatus_Resolved:
		{
			_, ok := other.Get_().(FindingStatus_Resolved)
			return ok
		}
	case FindingStatus_FindingUnavailable:
		{
			_, ok := other.Get_().(FindingStatus_FindingUnavailable)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this FindingStatus) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(FindingStatus)
	return ok && _this.Equals(typed)
}

func Type_FindingStatus_() _dafny.TypeDescriptor {
	return type_FindingStatus_{}
}

type type_FindingStatus_ struct {
}

func (_this type_FindingStatus_) Default() interface{} {
	return Companion_FindingStatus_.Default()
}

func (_this type_FindingStatus_) String() string {
	return "ConfluxSemanticGraphContract.FindingStatus"
}
func (_this FindingStatus) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = FindingStatus{}

// End of datatype FindingStatus

// Definition of datatype EvidenceKind
type EvidenceKind struct {
	Data_EvidenceKind_
}

func (_this EvidenceKind) Get_() Data_EvidenceKind_ {
	return _this.Data_EvidenceKind_
}

type Data_EvidenceKind_ interface {
	isEvidenceKind()
}

type CompanionStruct_EvidenceKind_ struct {
}

var Companion_EvidenceKind_ = CompanionStruct_EvidenceKind_{}

type EvidenceKind_Source struct {
}

func (EvidenceKind_Source) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Source_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Source{}}
}

func (_this EvidenceKind) Is_Source() bool {
	_, ok := _this.Get_().(EvidenceKind_Source)
	return ok
}

type EvidenceKind_Overlay struct {
}

func (EvidenceKind_Overlay) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Overlay_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Overlay{}}
}

func (_this EvidenceKind) Is_Overlay() bool {
	_, ok := _this.Get_().(EvidenceKind_Overlay)
	return ok
}

type EvidenceKind_Directory struct {
}

func (EvidenceKind_Directory) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Directory_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Directory{}}
}

func (_this EvidenceKind) Is_Directory() bool {
	_, ok := _this.Get_().(EvidenceKind_Directory)
	return ok
}

type EvidenceKind_Config struct {
}

func (EvidenceKind_Config) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Config_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Config{}}
}

func (_this EvidenceKind) Is_Config() bool {
	_, ok := _this.Get_().(EvidenceKind_Config)
	return ok
}

type EvidenceKind_Std struct {
}

func (EvidenceKind_Std) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Std_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Std{}}
}

func (_this EvidenceKind) Is_Std() bool {
	_, ok := _this.Get_().(EvidenceKind_Std)
	return ok
}

type EvidenceKind_Doo struct {
}

func (EvidenceKind_Doo) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Doo_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Doo{}}
}

func (_this EvidenceKind) Is_Doo() bool {
	_, ok := _this.Get_().(EvidenceKind_Doo)
	return ok
}

type EvidenceKind_Syntax struct {
}

func (EvidenceKind_Syntax) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Syntax_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Syntax{}}
}

func (_this EvidenceKind) Is_Syntax() bool {
	_, ok := _this.Get_().(EvidenceKind_Syntax)
	return ok
}

type EvidenceKind_Symbol struct {
}

func (EvidenceKind_Symbol) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Symbol_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Symbol{}}
}

func (_this EvidenceKind) Is_Symbol() bool {
	_, ok := _this.Get_().(EvidenceKind_Symbol)
	return ok
}

type EvidenceKind_CallEdge struct {
}

func (EvidenceKind_CallEdge) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_CallEdge_() EvidenceKind {
	return EvidenceKind{EvidenceKind_CallEdge{}}
}

func (_this EvidenceKind) Is_CallEdge() bool {
	_, ok := _this.Get_().(EvidenceKind_CallEdge)
	return ok
}

type EvidenceKind_LexicalMatch struct {
}

func (EvidenceKind_LexicalMatch) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_LexicalMatch_() EvidenceKind {
	return EvidenceKind{EvidenceKind_LexicalMatch{}}
}

func (_this EvidenceKind) Is_LexicalMatch() bool {
	_, ok := _this.Get_().(EvidenceKind_LexicalMatch)
	return ok
}

type EvidenceKind_Plugin struct {
}

func (EvidenceKind_Plugin) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Plugin_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Plugin{}}
}

func (_this EvidenceKind) Is_Plugin() bool {
	_, ok := _this.Get_().(EvidenceKind_Plugin)
	return ok
}

type EvidenceKind_Rule struct {
}

func (EvidenceKind_Rule) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Rule_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Rule{}}
}

func (_this EvidenceKind) Is_Rule() bool {
	_, ok := _this.Get_().(EvidenceKind_Rule)
	return ok
}

type EvidenceKind_Derivation struct {
}

func (EvidenceKind_Derivation) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Derivation_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Derivation{}}
}

func (_this EvidenceKind) Is_Derivation() bool {
	_, ok := _this.Get_().(EvidenceKind_Derivation)
	return ok
}

type EvidenceKind_PolicyEvidence struct {
}

func (EvidenceKind_PolicyEvidence) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_PolicyEvidence_() EvidenceKind {
	return EvidenceKind{EvidenceKind_PolicyEvidence{}}
}

func (_this EvidenceKind) Is_PolicyEvidence() bool {
	_, ok := _this.Get_().(EvidenceKind_PolicyEvidence)
	return ok
}

type EvidenceKind_Receipt struct {
}

func (EvidenceKind_Receipt) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_Receipt_() EvidenceKind {
	return EvidenceKind{EvidenceKind_Receipt{}}
}

func (_this EvidenceKind) Is_Receipt() bool {
	_, ok := _this.Get_().(EvidenceKind_Receipt)
	return ok
}

type EvidenceKind_CacheDecision struct {
}

func (EvidenceKind_CacheDecision) isEvidenceKind() {}

func (CompanionStruct_EvidenceKind_) Create_CacheDecision_() EvidenceKind {
	return EvidenceKind{EvidenceKind_CacheDecision{}}
}

func (_this EvidenceKind) Is_CacheDecision() bool {
	_, ok := _this.Get_().(EvidenceKind_CacheDecision)
	return ok
}

func (CompanionStruct_EvidenceKind_) Default() EvidenceKind {
	return Companion_EvidenceKind_.Create_Source_()
}

func (_ CompanionStruct_EvidenceKind_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_EvidenceKind_.Create_Source_(), true
		case 1:
			return Companion_EvidenceKind_.Create_Overlay_(), true
		case 2:
			return Companion_EvidenceKind_.Create_Directory_(), true
		case 3:
			return Companion_EvidenceKind_.Create_Config_(), true
		case 4:
			return Companion_EvidenceKind_.Create_Std_(), true
		case 5:
			return Companion_EvidenceKind_.Create_Doo_(), true
		case 6:
			return Companion_EvidenceKind_.Create_Syntax_(), true
		case 7:
			return Companion_EvidenceKind_.Create_Symbol_(), true
		case 8:
			return Companion_EvidenceKind_.Create_CallEdge_(), true
		case 9:
			return Companion_EvidenceKind_.Create_LexicalMatch_(), true
		case 10:
			return Companion_EvidenceKind_.Create_Plugin_(), true
		case 11:
			return Companion_EvidenceKind_.Create_Rule_(), true
		case 12:
			return Companion_EvidenceKind_.Create_Derivation_(), true
		case 13:
			return Companion_EvidenceKind_.Create_PolicyEvidence_(), true
		case 14:
			return Companion_EvidenceKind_.Create_Receipt_(), true
		case 15:
			return Companion_EvidenceKind_.Create_CacheDecision_(), true
		default:
			return EvidenceKind{}, false
		}
	}
}

func (_this EvidenceKind) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case EvidenceKind_Source:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Source"
		}
	case EvidenceKind_Overlay:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Overlay"
		}
	case EvidenceKind_Directory:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Directory"
		}
	case EvidenceKind_Config:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Config"
		}
	case EvidenceKind_Std:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Std"
		}
	case EvidenceKind_Doo:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Doo"
		}
	case EvidenceKind_Syntax:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Syntax"
		}
	case EvidenceKind_Symbol:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Symbol"
		}
	case EvidenceKind_CallEdge:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.CallEdge"
		}
	case EvidenceKind_LexicalMatch:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.LexicalMatch"
		}
	case EvidenceKind_Plugin:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Plugin"
		}
	case EvidenceKind_Rule:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Rule"
		}
	case EvidenceKind_Derivation:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Derivation"
		}
	case EvidenceKind_PolicyEvidence:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.PolicyEvidence"
		}
	case EvidenceKind_Receipt:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.Receipt"
		}
	case EvidenceKind_CacheDecision:
		{
			return "ConfluxSemanticGraphContract.EvidenceKind.CacheDecision"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this EvidenceKind) Equals(other EvidenceKind) bool {
	switch _this.Get_().(type) {
	case EvidenceKind_Source:
		{
			_, ok := other.Get_().(EvidenceKind_Source)
			return ok
		}
	case EvidenceKind_Overlay:
		{
			_, ok := other.Get_().(EvidenceKind_Overlay)
			return ok
		}
	case EvidenceKind_Directory:
		{
			_, ok := other.Get_().(EvidenceKind_Directory)
			return ok
		}
	case EvidenceKind_Config:
		{
			_, ok := other.Get_().(EvidenceKind_Config)
			return ok
		}
	case EvidenceKind_Std:
		{
			_, ok := other.Get_().(EvidenceKind_Std)
			return ok
		}
	case EvidenceKind_Doo:
		{
			_, ok := other.Get_().(EvidenceKind_Doo)
			return ok
		}
	case EvidenceKind_Syntax:
		{
			_, ok := other.Get_().(EvidenceKind_Syntax)
			return ok
		}
	case EvidenceKind_Symbol:
		{
			_, ok := other.Get_().(EvidenceKind_Symbol)
			return ok
		}
	case EvidenceKind_CallEdge:
		{
			_, ok := other.Get_().(EvidenceKind_CallEdge)
			return ok
		}
	case EvidenceKind_LexicalMatch:
		{
			_, ok := other.Get_().(EvidenceKind_LexicalMatch)
			return ok
		}
	case EvidenceKind_Plugin:
		{
			_, ok := other.Get_().(EvidenceKind_Plugin)
			return ok
		}
	case EvidenceKind_Rule:
		{
			_, ok := other.Get_().(EvidenceKind_Rule)
			return ok
		}
	case EvidenceKind_Derivation:
		{
			_, ok := other.Get_().(EvidenceKind_Derivation)
			return ok
		}
	case EvidenceKind_PolicyEvidence:
		{
			_, ok := other.Get_().(EvidenceKind_PolicyEvidence)
			return ok
		}
	case EvidenceKind_Receipt:
		{
			_, ok := other.Get_().(EvidenceKind_Receipt)
			return ok
		}
	case EvidenceKind_CacheDecision:
		{
			_, ok := other.Get_().(EvidenceKind_CacheDecision)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this EvidenceKind) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(EvidenceKind)
	return ok && _this.Equals(typed)
}

func Type_EvidenceKind_() _dafny.TypeDescriptor {
	return type_EvidenceKind_{}
}

type type_EvidenceKind_ struct {
}

func (_this type_EvidenceKind_) Default() interface{} {
	return Companion_EvidenceKind_.Default()
}

func (_this type_EvidenceKind_) String() string {
	return "ConfluxSemanticGraphContract.EvidenceKind"
}
func (_this EvidenceKind) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = EvidenceKind{}

// End of datatype EvidenceKind

// Definition of datatype LensArchetype
type LensArchetype struct {
	Data_LensArchetype_
}

func (_this LensArchetype) Get_() Data_LensArchetype_ {
	return _this.Data_LensArchetype_
}

type Data_LensArchetype_ interface {
	isLensArchetype()
}

type CompanionStruct_LensArchetype_ struct {
}

var Companion_LensArchetype_ = CompanionStruct_LensArchetype_{}

type LensArchetype_Graph struct {
}

func (LensArchetype_Graph) isLensArchetype() {}

func (CompanionStruct_LensArchetype_) Create_Graph_() LensArchetype {
	return LensArchetype{LensArchetype_Graph{}}
}

func (_this LensArchetype) Is_Graph() bool {
	_, ok := _this.Get_().(LensArchetype_Graph)
	return ok
}

type LensArchetype_Matrix struct {
}

func (LensArchetype_Matrix) isLensArchetype() {}

func (CompanionStruct_LensArchetype_) Create_Matrix_() LensArchetype {
	return LensArchetype{LensArchetype_Matrix{}}
}

func (_this LensArchetype) Is_Matrix() bool {
	_, ok := _this.Get_().(LensArchetype_Matrix)
	return ok
}

type LensArchetype_Layered struct {
}

func (LensArchetype_Layered) isLensArchetype() {}

func (CompanionStruct_LensArchetype_) Create_Layered_() LensArchetype {
	return LensArchetype{LensArchetype_Layered{}}
}

func (_this LensArchetype) Is_Layered() bool {
	_, ok := _this.Get_().(LensArchetype_Layered)
	return ok
}

type LensArchetype_List struct {
}

func (LensArchetype_List) isLensArchetype() {}

func (CompanionStruct_LensArchetype_) Create_List_() LensArchetype {
	return LensArchetype{LensArchetype_List{}}
}

func (_this LensArchetype) Is_List() bool {
	_, ok := _this.Get_().(LensArchetype_List)
	return ok
}

type LensArchetype_Table struct {
}

func (LensArchetype_Table) isLensArchetype() {}

func (CompanionStruct_LensArchetype_) Create_Table_() LensArchetype {
	return LensArchetype{LensArchetype_Table{}}
}

func (_this LensArchetype) Is_Table() bool {
	_, ok := _this.Get_().(LensArchetype_Table)
	return ok
}

type LensArchetype_MasterDetail struct {
}

func (LensArchetype_MasterDetail) isLensArchetype() {}

func (CompanionStruct_LensArchetype_) Create_MasterDetail_() LensArchetype {
	return LensArchetype{LensArchetype_MasterDetail{}}
}

func (_this LensArchetype) Is_MasterDetail() bool {
	_, ok := _this.Get_().(LensArchetype_MasterDetail)
	return ok
}

func (CompanionStruct_LensArchetype_) Default() LensArchetype {
	return Companion_LensArchetype_.Create_Graph_()
}

func (_ CompanionStruct_LensArchetype_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_LensArchetype_.Create_Graph_(), true
		case 1:
			return Companion_LensArchetype_.Create_Matrix_(), true
		case 2:
			return Companion_LensArchetype_.Create_Layered_(), true
		case 3:
			return Companion_LensArchetype_.Create_List_(), true
		case 4:
			return Companion_LensArchetype_.Create_Table_(), true
		case 5:
			return Companion_LensArchetype_.Create_MasterDetail_(), true
		default:
			return LensArchetype{}, false
		}
	}
}

func (_this LensArchetype) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case LensArchetype_Graph:
		{
			return "ConfluxSemanticGraphContract.LensArchetype.Graph"
		}
	case LensArchetype_Matrix:
		{
			return "ConfluxSemanticGraphContract.LensArchetype.Matrix"
		}
	case LensArchetype_Layered:
		{
			return "ConfluxSemanticGraphContract.LensArchetype.Layered"
		}
	case LensArchetype_List:
		{
			return "ConfluxSemanticGraphContract.LensArchetype.List"
		}
	case LensArchetype_Table:
		{
			return "ConfluxSemanticGraphContract.LensArchetype.Table"
		}
	case LensArchetype_MasterDetail:
		{
			return "ConfluxSemanticGraphContract.LensArchetype.MasterDetail"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this LensArchetype) Equals(other LensArchetype) bool {
	switch _this.Get_().(type) {
	case LensArchetype_Graph:
		{
			_, ok := other.Get_().(LensArchetype_Graph)
			return ok
		}
	case LensArchetype_Matrix:
		{
			_, ok := other.Get_().(LensArchetype_Matrix)
			return ok
		}
	case LensArchetype_Layered:
		{
			_, ok := other.Get_().(LensArchetype_Layered)
			return ok
		}
	case LensArchetype_List:
		{
			_, ok := other.Get_().(LensArchetype_List)
			return ok
		}
	case LensArchetype_Table:
		{
			_, ok := other.Get_().(LensArchetype_Table)
			return ok
		}
	case LensArchetype_MasterDetail:
		{
			_, ok := other.Get_().(LensArchetype_MasterDetail)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this LensArchetype) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(LensArchetype)
	return ok && _this.Equals(typed)
}

func Type_LensArchetype_() _dafny.TypeDescriptor {
	return type_LensArchetype_{}
}

type type_LensArchetype_ struct {
}

func (_this type_LensArchetype_) Default() interface{} {
	return Companion_LensArchetype_.Default()
}

func (_this type_LensArchetype_) String() string {
	return "ConfluxSemanticGraphContract.LensArchetype"
}
func (_this LensArchetype) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = LensArchetype{}

// End of datatype LensArchetype

// Definition of datatype TypedValue
type TypedValue struct {
	Data_TypedValue_
}

func (_this TypedValue) Get_() Data_TypedValue_ {
	return _this.Data_TypedValue_
}

type Data_TypedValue_ interface {
	isTypedValue()
}

type CompanionStruct_TypedValue_ struct {
}

var Companion_TypedValue_ = CompanionStruct_TypedValue_{}

type TypedValue_TVNull struct {
}

func (TypedValue_TVNull) isTypedValue() {}

func (CompanionStruct_TypedValue_) Create_TVNull_() TypedValue {
	return TypedValue{TypedValue_TVNull{}}
}

func (_this TypedValue) Is_TVNull() bool {
	_, ok := _this.Get_().(TypedValue_TVNull)
	return ok
}

type TypedValue_TVBool struct {
	BoolValue bool
}

func (TypedValue_TVBool) isTypedValue() {}

func (CompanionStruct_TypedValue_) Create_TVBool_(BoolValue bool) TypedValue {
	return TypedValue{TypedValue_TVBool{BoolValue}}
}

func (_this TypedValue) Is_TVBool() bool {
	_, ok := _this.Get_().(TypedValue_TVBool)
	return ok
}

type TypedValue_TVInteger struct {
	DecimalString _dafny.Sequence
}

func (TypedValue_TVInteger) isTypedValue() {}

func (CompanionStruct_TypedValue_) Create_TVInteger_(DecimalString _dafny.Sequence) TypedValue {
	return TypedValue{TypedValue_TVInteger{DecimalString}}
}

func (_this TypedValue) Is_TVInteger() bool {
	_, ok := _this.Get_().(TypedValue_TVInteger)
	return ok
}

type TypedValue_TVDecimal struct {
	CanonicalDecimalString _dafny.Sequence
}

func (TypedValue_TVDecimal) isTypedValue() {}

func (CompanionStruct_TypedValue_) Create_TVDecimal_(CanonicalDecimalString _dafny.Sequence) TypedValue {
	return TypedValue{TypedValue_TVDecimal{CanonicalDecimalString}}
}

func (_this TypedValue) Is_TVDecimal() bool {
	_, ok := _this.Get_().(TypedValue_TVDecimal)
	return ok
}

type TypedValue_TVString struct {
	StringValue _dafny.Sequence
}

func (TypedValue_TVString) isTypedValue() {}

func (CompanionStruct_TypedValue_) Create_TVString_(StringValue _dafny.Sequence) TypedValue {
	return TypedValue{TypedValue_TVString{StringValue}}
}

func (_this TypedValue) Is_TVString() bool {
	_, ok := _this.Get_().(TypedValue_TVString)
	return ok
}

type TypedValue_TVRef struct {
	TargetId _dafny.Sequence
}

func (TypedValue_TVRef) isTypedValue() {}

func (CompanionStruct_TypedValue_) Create_TVRef_(TargetId _dafny.Sequence) TypedValue {
	return TypedValue{TypedValue_TVRef{TargetId}}
}

func (_this TypedValue) Is_TVRef() bool {
	_, ok := _this.Get_().(TypedValue_TVRef)
	return ok
}

type TypedValue_TVList struct {
	Items _dafny.Sequence
}

func (TypedValue_TVList) isTypedValue() {}

func (CompanionStruct_TypedValue_) Create_TVList_(Items _dafny.Sequence) TypedValue {
	return TypedValue{TypedValue_TVList{Items}}
}

func (_this TypedValue) Is_TVList() bool {
	_, ok := _this.Get_().(TypedValue_TVList)
	return ok
}

type TypedValue_TVRecord struct {
	FieldNames  _dafny.Sequence
	FieldValues _dafny.Sequence
}

func (TypedValue_TVRecord) isTypedValue() {}

func (CompanionStruct_TypedValue_) Create_TVRecord_(FieldNames _dafny.Sequence, FieldValues _dafny.Sequence) TypedValue {
	return TypedValue{TypedValue_TVRecord{FieldNames, FieldValues}}
}

func (_this TypedValue) Is_TVRecord() bool {
	_, ok := _this.Get_().(TypedValue_TVRecord)
	return ok
}

func (CompanionStruct_TypedValue_) Default() TypedValue {
	return Companion_TypedValue_.Create_TVNull_()
}

func (_this TypedValue) Dtor_boolValue() bool {
	return _this.Get_().(TypedValue_TVBool).BoolValue
}

func (_this TypedValue) Dtor_decimalString() _dafny.Sequence {
	return _this.Get_().(TypedValue_TVInteger).DecimalString
}

func (_this TypedValue) Dtor_canonicalDecimalString() _dafny.Sequence {
	return _this.Get_().(TypedValue_TVDecimal).CanonicalDecimalString
}

func (_this TypedValue) Dtor_stringValue() _dafny.Sequence {
	return _this.Get_().(TypedValue_TVString).StringValue
}

func (_this TypedValue) Dtor_targetId() _dafny.Sequence {
	return _this.Get_().(TypedValue_TVRef).TargetId
}

func (_this TypedValue) Dtor_items() _dafny.Sequence {
	return _this.Get_().(TypedValue_TVList).Items
}

func (_this TypedValue) Dtor_fieldNames() _dafny.Sequence {
	return _this.Get_().(TypedValue_TVRecord).FieldNames
}

func (_this TypedValue) Dtor_fieldValues() _dafny.Sequence {
	return _this.Get_().(TypedValue_TVRecord).FieldValues
}

func (_this TypedValue) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case TypedValue_TVNull:
		{
			return "ConfluxSemanticGraphContract.TypedValue.TVNull"
		}
	case TypedValue_TVBool:
		{
			return "ConfluxSemanticGraphContract.TypedValue.TVBool" + "(" + _dafny.String(data.BoolValue) + ")"
		}
	case TypedValue_TVInteger:
		{
			return "ConfluxSemanticGraphContract.TypedValue.TVInteger" + "(" + data.DecimalString.VerbatimString(true) + ")"
		}
	case TypedValue_TVDecimal:
		{
			return "ConfluxSemanticGraphContract.TypedValue.TVDecimal" + "(" + data.CanonicalDecimalString.VerbatimString(true) + ")"
		}
	case TypedValue_TVString:
		{
			return "ConfluxSemanticGraphContract.TypedValue.TVString" + "(" + data.StringValue.VerbatimString(true) + ")"
		}
	case TypedValue_TVRef:
		{
			return "ConfluxSemanticGraphContract.TypedValue.TVRef" + "(" + data.TargetId.VerbatimString(true) + ")"
		}
	case TypedValue_TVList:
		{
			return "ConfluxSemanticGraphContract.TypedValue.TVList" + "(" + _dafny.String(data.Items) + ")"
		}
	case TypedValue_TVRecord:
		{
			return "ConfluxSemanticGraphContract.TypedValue.TVRecord" + "(" + _dafny.String(data.FieldNames) + ", " + _dafny.String(data.FieldValues) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this TypedValue) Equals(other TypedValue) bool {
	switch data1 := _this.Get_().(type) {
	case TypedValue_TVNull:
		{
			_, ok := other.Get_().(TypedValue_TVNull)
			return ok
		}
	case TypedValue_TVBool:
		{
			data2, ok := other.Get_().(TypedValue_TVBool)
			return ok && data1.BoolValue == data2.BoolValue
		}
	case TypedValue_TVInteger:
		{
			data2, ok := other.Get_().(TypedValue_TVInteger)
			return ok && data1.DecimalString.Equals(data2.DecimalString)
		}
	case TypedValue_TVDecimal:
		{
			data2, ok := other.Get_().(TypedValue_TVDecimal)
			return ok && data1.CanonicalDecimalString.Equals(data2.CanonicalDecimalString)
		}
	case TypedValue_TVString:
		{
			data2, ok := other.Get_().(TypedValue_TVString)
			return ok && data1.StringValue.Equals(data2.StringValue)
		}
	case TypedValue_TVRef:
		{
			data2, ok := other.Get_().(TypedValue_TVRef)
			return ok && data1.TargetId.Equals(data2.TargetId)
		}
	case TypedValue_TVList:
		{
			data2, ok := other.Get_().(TypedValue_TVList)
			return ok && data1.Items.Equals(data2.Items)
		}
	case TypedValue_TVRecord:
		{
			data2, ok := other.Get_().(TypedValue_TVRecord)
			return ok && data1.FieldNames.Equals(data2.FieldNames) && data1.FieldValues.Equals(data2.FieldValues)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this TypedValue) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(TypedValue)
	return ok && _this.Equals(typed)
}

func Type_TypedValue_() _dafny.TypeDescriptor {
	return type_TypedValue_{}
}

type type_TypedValue_ struct {
}

func (_this type_TypedValue_) Default() interface{} {
	return Companion_TypedValue_.Default()
}

func (_this type_TypedValue_) String() string {
	return "ConfluxSemanticGraphContract.TypedValue"
}
func (_this TypedValue) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = TypedValue{}

// End of datatype TypedValue

// Definition of datatype TypedField
type TypedField struct {
	Data_TypedField_
}

func (_this TypedField) Get_() Data_TypedField_ {
	return _this.Data_TypedField_
}

type Data_TypedField_ interface {
	isTypedField()
}

type CompanionStruct_TypedField_ struct {
}

var Companion_TypedField_ = CompanionStruct_TypedField_{}

type TypedField_TypedField struct {
	Name  _dafny.Sequence
	Value TypedValue
}

func (TypedField_TypedField) isTypedField() {}

func (CompanionStruct_TypedField_) Create_TypedField_(Name _dafny.Sequence, Value TypedValue) TypedField {
	return TypedField{TypedField_TypedField{Name, Value}}
}

func (_this TypedField) Is_TypedField() bool {
	_, ok := _this.Get_().(TypedField_TypedField)
	return ok
}

func (CompanionStruct_TypedField_) Default() TypedField {
	return Companion_TypedField_.Create_TypedField_(_dafny.EmptySeq, Companion_TypedValue_.Default())
}

func (_this TypedField) Dtor_name() _dafny.Sequence {
	return _this.Get_().(TypedField_TypedField).Name
}

func (_this TypedField) Dtor_value() TypedValue {
	return _this.Get_().(TypedField_TypedField).Value
}

func (_this TypedField) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case TypedField_TypedField:
		{
			return "ConfluxSemanticGraphContract.TypedField.TypedField" + "(" + data.Name.VerbatimString(true) + ", " + _dafny.String(data.Value) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this TypedField) Equals(other TypedField) bool {
	switch data1 := _this.Get_().(type) {
	case TypedField_TypedField:
		{
			data2, ok := other.Get_().(TypedField_TypedField)
			return ok && data1.Name.Equals(data2.Name) && data1.Value.Equals(data2.Value)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this TypedField) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(TypedField)
	return ok && _this.Equals(typed)
}

func Type_TypedField_() _dafny.TypeDescriptor {
	return type_TypedField_{}
}

type type_TypedField_ struct {
}

func (_this type_TypedField_) Default() interface{} {
	return Companion_TypedField_.Default()
}

func (_this type_TypedField_) String() string {
	return "ConfluxSemanticGraphContract.TypedField"
}
func (_this TypedField) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = TypedField{}

// End of datatype TypedField

// Definition of datatype Completeness
type Completeness struct {
	Data_Completeness_
}

func (_this Completeness) Get_() Data_Completeness_ {
	return _this.Data_Completeness_
}

type Data_Completeness_ interface {
	isCompleteness()
}

type CompanionStruct_Completeness_ struct {
}

var Companion_Completeness_ = CompanionStruct_Completeness_{}

type Completeness_Completeness struct {
	Complete bool
	Reasons  _dafny.Sequence
}

func (Completeness_Completeness) isCompleteness() {}

func (CompanionStruct_Completeness_) Create_Completeness_(Complete bool, Reasons _dafny.Sequence) Completeness {
	return Completeness{Completeness_Completeness{Complete, Reasons}}
}

func (_this Completeness) Is_Completeness() bool {
	_, ok := _this.Get_().(Completeness_Completeness)
	return ok
}

func (CompanionStruct_Completeness_) Default() Completeness {
	return Companion_Completeness_.Create_Completeness_(false, _dafny.EmptySeq)
}

func (_this Completeness) Dtor_complete() bool {
	return _this.Get_().(Completeness_Completeness).Complete
}

func (_this Completeness) Dtor_reasons() _dafny.Sequence {
	return _this.Get_().(Completeness_Completeness).Reasons
}

func (_this Completeness) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Completeness_Completeness:
		{
			return "ConfluxSemanticGraphContract.Completeness.Completeness" + "(" + _dafny.String(data.Complete) + ", " + _dafny.String(data.Reasons) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Completeness) Equals(other Completeness) bool {
	switch data1 := _this.Get_().(type) {
	case Completeness_Completeness:
		{
			data2, ok := other.Get_().(Completeness_Completeness)
			return ok && data1.Complete == data2.Complete && data1.Reasons.Equals(data2.Reasons)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Completeness) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Completeness)
	return ok && _this.Equals(typed)
}

func Type_Completeness_() _dafny.TypeDescriptor {
	return type_Completeness_{}
}

type type_Completeness_ struct {
}

func (_this type_Completeness_) Default() interface{} {
	return Companion_Completeness_.Default()
}

func (_this type_Completeness_) String() string {
	return "ConfluxSemanticGraphContract.Completeness"
}
func (_this Completeness) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Completeness{}

// End of datatype Completeness

// Definition of datatype SourcePosition
type SourcePosition struct {
	Data_SourcePosition_
}

func (_this SourcePosition) Get_() Data_SourcePosition_ {
	return _this.Data_SourcePosition_
}

type Data_SourcePosition_ interface {
	isSourcePosition()
}

type CompanionStruct_SourcePosition_ struct {
}

var Companion_SourcePosition_ = CompanionStruct_SourcePosition_{}

type SourcePosition_SourcePosition struct {
	Line   _dafny.Int
	Column _dafny.Int
}

func (SourcePosition_SourcePosition) isSourcePosition() {}

func (CompanionStruct_SourcePosition_) Create_SourcePosition_(Line _dafny.Int, Column _dafny.Int) SourcePosition {
	return SourcePosition{SourcePosition_SourcePosition{Line, Column}}
}

func (_this SourcePosition) Is_SourcePosition() bool {
	_, ok := _this.Get_().(SourcePosition_SourcePosition)
	return ok
}

func (CompanionStruct_SourcePosition_) Default() SourcePosition {
	return Companion_SourcePosition_.Create_SourcePosition_(_dafny.Zero, _dafny.Zero)
}

func (_this SourcePosition) Dtor_line() _dafny.Int {
	return _this.Get_().(SourcePosition_SourcePosition).Line
}

func (_this SourcePosition) Dtor_column() _dafny.Int {
	return _this.Get_().(SourcePosition_SourcePosition).Column
}

func (_this SourcePosition) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SourcePosition_SourcePosition:
		{
			return "ConfluxSemanticGraphContract.SourcePosition.SourcePosition" + "(" + _dafny.String(data.Line) + ", " + _dafny.String(data.Column) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SourcePosition) Equals(other SourcePosition) bool {
	switch data1 := _this.Get_().(type) {
	case SourcePosition_SourcePosition:
		{
			data2, ok := other.Get_().(SourcePosition_SourcePosition)
			return ok && data1.Line.Cmp(data2.Line) == 0 && data1.Column.Cmp(data2.Column) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SourcePosition) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SourcePosition)
	return ok && _this.Equals(typed)
}

func Type_SourcePosition_() _dafny.TypeDescriptor {
	return type_SourcePosition_{}
}

type type_SourcePosition_ struct {
}

func (_this type_SourcePosition_) Default() interface{} {
	return Companion_SourcePosition_.Default()
}

func (_this type_SourcePosition_) String() string {
	return "ConfluxSemanticGraphContract.SourcePosition"
}
func (_this SourcePosition) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SourcePosition{}

// End of datatype SourcePosition

// Definition of datatype SourceRange
type SourceRange struct {
	Data_SourceRange_
}

func (_this SourceRange) Get_() Data_SourceRange_ {
	return _this.Data_SourceRange_
}

type Data_SourceRange_ interface {
	isSourceRange()
}

type CompanionStruct_SourceRange_ struct {
}

var Companion_SourceRange_ = CompanionStruct_SourceRange_{}

type SourceRange_SourceRange struct {
	Start SourcePosition
	End   SourcePosition
}

func (SourceRange_SourceRange) isSourceRange() {}

func (CompanionStruct_SourceRange_) Create_SourceRange_(Start SourcePosition, End SourcePosition) SourceRange {
	return SourceRange{SourceRange_SourceRange{Start, End}}
}

func (_this SourceRange) Is_SourceRange() bool {
	_, ok := _this.Get_().(SourceRange_SourceRange)
	return ok
}

func (CompanionStruct_SourceRange_) Default() SourceRange {
	return Companion_SourceRange_.Create_SourceRange_(Companion_SourcePosition_.Default(), Companion_SourcePosition_.Default())
}

func (_this SourceRange) Dtor_start() SourcePosition {
	return _this.Get_().(SourceRange_SourceRange).Start
}

func (_this SourceRange) Dtor_end() SourcePosition {
	return _this.Get_().(SourceRange_SourceRange).End
}

func (_this SourceRange) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SourceRange_SourceRange:
		{
			return "ConfluxSemanticGraphContract.SourceRange.SourceRange" + "(" + _dafny.String(data.Start) + ", " + _dafny.String(data.End) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SourceRange) Equals(other SourceRange) bool {
	switch data1 := _this.Get_().(type) {
	case SourceRange_SourceRange:
		{
			data2, ok := other.Get_().(SourceRange_SourceRange)
			return ok && data1.Start.Equals(data2.Start) && data1.End.Equals(data2.End)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SourceRange) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SourceRange)
	return ok && _this.Equals(typed)
}

func Type_SourceRange_() _dafny.TypeDescriptor {
	return type_SourceRange_{}
}

type type_SourceRange_ struct {
}

func (_this type_SourceRange_) Default() interface{} {
	return Companion_SourceRange_.Default()
}

func (_this type_SourceRange_) String() string {
	return "ConfluxSemanticGraphContract.SourceRange"
}
func (_this SourceRange) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SourceRange{}

// End of datatype SourceRange

// Definition of datatype CapabilityRequirement
type CapabilityRequirement struct {
	Data_CapabilityRequirement_
}

func (_this CapabilityRequirement) Get_() Data_CapabilityRequirement_ {
	return _this.Data_CapabilityRequirement_
}

type Data_CapabilityRequirement_ interface {
	isCapabilityRequirement()
}

type CompanionStruct_CapabilityRequirement_ struct {
}

var Companion_CapabilityRequirement_ = CompanionStruct_CapabilityRequirement_{}

type CapabilityRequirement_CapabilityRequirement struct {
	Id           _dafny.Sequence
	VersionRange _dafny.Sequence
}

func (CapabilityRequirement_CapabilityRequirement) isCapabilityRequirement() {}

func (CompanionStruct_CapabilityRequirement_) Create_CapabilityRequirement_(Id _dafny.Sequence, VersionRange _dafny.Sequence) CapabilityRequirement {
	return CapabilityRequirement{CapabilityRequirement_CapabilityRequirement{Id, VersionRange}}
}

func (_this CapabilityRequirement) Is_CapabilityRequirement() bool {
	_, ok := _this.Get_().(CapabilityRequirement_CapabilityRequirement)
	return ok
}

func (CompanionStruct_CapabilityRequirement_) Default() CapabilityRequirement {
	return Companion_CapabilityRequirement_.Create_CapabilityRequirement_(_dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this CapabilityRequirement) Dtor_id() _dafny.Sequence {
	return _this.Get_().(CapabilityRequirement_CapabilityRequirement).Id
}

func (_this CapabilityRequirement) Dtor_versionRange() _dafny.Sequence {
	return _this.Get_().(CapabilityRequirement_CapabilityRequirement).VersionRange
}

func (_this CapabilityRequirement) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CapabilityRequirement_CapabilityRequirement:
		{
			return "ConfluxSemanticGraphContract.CapabilityRequirement.CapabilityRequirement" + "(" + data.Id.VerbatimString(true) + ", " + data.VersionRange.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CapabilityRequirement) Equals(other CapabilityRequirement) bool {
	switch data1 := _this.Get_().(type) {
	case CapabilityRequirement_CapabilityRequirement:
		{
			data2, ok := other.Get_().(CapabilityRequirement_CapabilityRequirement)
			return ok && data1.Id.Equals(data2.Id) && data1.VersionRange.Equals(data2.VersionRange)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CapabilityRequirement) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CapabilityRequirement)
	return ok && _this.Equals(typed)
}

func Type_CapabilityRequirement_() _dafny.TypeDescriptor {
	return type_CapabilityRequirement_{}
}

type type_CapabilityRequirement_ struct {
}

func (_this type_CapabilityRequirement_) Default() interface{} {
	return Companion_CapabilityRequirement_.Default()
}

func (_this type_CapabilityRequirement_) String() string {
	return "ConfluxSemanticGraphContract.CapabilityRequirement"
}
func (_this CapabilityRequirement) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CapabilityRequirement{}

// End of datatype CapabilityRequirement

// Definition of datatype CapabilityDescriptor
type CapabilityDescriptor struct {
	Data_CapabilityDescriptor_
}

func (_this CapabilityDescriptor) Get_() Data_CapabilityDescriptor_ {
	return _this.Data_CapabilityDescriptor_
}

type Data_CapabilityDescriptor_ interface {
	isCapabilityDescriptor()
}

type CompanionStruct_CapabilityDescriptor_ struct {
}

var Companion_CapabilityDescriptor_ = CompanionStruct_CapabilityDescriptor_{}

type CapabilityDescriptor_CapabilityDescriptor struct {
	Id                   _dafny.Sequence
	Version              _dafny.Sequence
	SchemaDigest         _dafny.Sequence
	ImplementationDigest _dafny.Sequence
	Kind                 CapabilityKind
	Scopes               _dafny.Set
	Requirements         _dafny.Sequence
	Availability         Availability
	UnavailableReasons   _dafny.Sequence
	ContentRevision      _dafny.Sequence
}

func (CapabilityDescriptor_CapabilityDescriptor) isCapabilityDescriptor() {}

func (CompanionStruct_CapabilityDescriptor_) Create_CapabilityDescriptor_(Id _dafny.Sequence, Version _dafny.Sequence, SchemaDigest _dafny.Sequence, ImplementationDigest _dafny.Sequence, Kind CapabilityKind, Scopes _dafny.Set, Requirements _dafny.Sequence, Availability Availability, UnavailableReasons _dafny.Sequence, ContentRevision _dafny.Sequence) CapabilityDescriptor {
	return CapabilityDescriptor{CapabilityDescriptor_CapabilityDescriptor{Id, Version, SchemaDigest, ImplementationDigest, Kind, Scopes, Requirements, Availability, UnavailableReasons, ContentRevision}}
}

func (_this CapabilityDescriptor) Is_CapabilityDescriptor() bool {
	_, ok := _this.Get_().(CapabilityDescriptor_CapabilityDescriptor)
	return ok
}

func (CompanionStruct_CapabilityDescriptor_) Default() CapabilityDescriptor {
	return Companion_CapabilityDescriptor_.Create_CapabilityDescriptor_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, Companion_CapabilityKind_.Default(), _dafny.EmptySet, _dafny.EmptySeq, Companion_Availability_.Default(), _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this CapabilityDescriptor) Dtor_id() _dafny.Sequence {
	return _this.Get_().(CapabilityDescriptor_CapabilityDescriptor).Id
}

func (_this CapabilityDescriptor) Dtor_version() _dafny.Sequence {
	return _this.Get_().(CapabilityDescriptor_CapabilityDescriptor).Version
}

func (_this CapabilityDescriptor) Dtor_schemaDigest() _dafny.Sequence {
	return _this.Get_().(CapabilityDescriptor_CapabilityDescriptor).SchemaDigest
}

func (_this CapabilityDescriptor) Dtor_implementationDigest() _dafny.Sequence {
	return _this.Get_().(CapabilityDescriptor_CapabilityDescriptor).ImplementationDigest
}

func (_this CapabilityDescriptor) Dtor_kind() CapabilityKind {
	return _this.Get_().(CapabilityDescriptor_CapabilityDescriptor).Kind
}

func (_this CapabilityDescriptor) Dtor_scopes() _dafny.Set {
	return _this.Get_().(CapabilityDescriptor_CapabilityDescriptor).Scopes
}

func (_this CapabilityDescriptor) Dtor_requirements() _dafny.Sequence {
	return _this.Get_().(CapabilityDescriptor_CapabilityDescriptor).Requirements
}

func (_this CapabilityDescriptor) Dtor_availability() Availability {
	return _this.Get_().(CapabilityDescriptor_CapabilityDescriptor).Availability
}

func (_this CapabilityDescriptor) Dtor_unavailableReasons() _dafny.Sequence {
	return _this.Get_().(CapabilityDescriptor_CapabilityDescriptor).UnavailableReasons
}

func (_this CapabilityDescriptor) Dtor_contentRevision() _dafny.Sequence {
	return _this.Get_().(CapabilityDescriptor_CapabilityDescriptor).ContentRevision
}

func (_this CapabilityDescriptor) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CapabilityDescriptor_CapabilityDescriptor:
		{
			return "ConfluxSemanticGraphContract.CapabilityDescriptor.CapabilityDescriptor" + "(" + data.Id.VerbatimString(true) + ", " + data.Version.VerbatimString(true) + ", " + data.SchemaDigest.VerbatimString(true) + ", " + data.ImplementationDigest.VerbatimString(true) + ", " + _dafny.String(data.Kind) + ", " + _dafny.String(data.Scopes) + ", " + _dafny.String(data.Requirements) + ", " + _dafny.String(data.Availability) + ", " + _dafny.String(data.UnavailableReasons) + ", " + data.ContentRevision.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CapabilityDescriptor) Equals(other CapabilityDescriptor) bool {
	switch data1 := _this.Get_().(type) {
	case CapabilityDescriptor_CapabilityDescriptor:
		{
			data2, ok := other.Get_().(CapabilityDescriptor_CapabilityDescriptor)
			return ok && data1.Id.Equals(data2.Id) && data1.Version.Equals(data2.Version) && data1.SchemaDigest.Equals(data2.SchemaDigest) && data1.ImplementationDigest.Equals(data2.ImplementationDigest) && data1.Kind.Equals(data2.Kind) && data1.Scopes.Equals(data2.Scopes) && data1.Requirements.Equals(data2.Requirements) && data1.Availability.Equals(data2.Availability) && data1.UnavailableReasons.Equals(data2.UnavailableReasons) && data1.ContentRevision.Equals(data2.ContentRevision)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CapabilityDescriptor) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CapabilityDescriptor)
	return ok && _this.Equals(typed)
}

func Type_CapabilityDescriptor_() _dafny.TypeDescriptor {
	return type_CapabilityDescriptor_{}
}

type type_CapabilityDescriptor_ struct {
}

func (_this type_CapabilityDescriptor_) Default() interface{} {
	return Companion_CapabilityDescriptor_.Default()
}

func (_this type_CapabilityDescriptor_) String() string {
	return "ConfluxSemanticGraphContract.CapabilityDescriptor"
}
func (_this CapabilityDescriptor) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CapabilityDescriptor{}

// End of datatype CapabilityDescriptor

// Definition of datatype AuthorityDescriptor
type AuthorityDescriptor struct {
	Data_AuthorityDescriptor_
}

func (_this AuthorityDescriptor) Get_() Data_AuthorityDescriptor_ {
	return _this.Data_AuthorityDescriptor_
}

type Data_AuthorityDescriptor_ interface {
	isAuthorityDescriptor()
}

type CompanionStruct_AuthorityDescriptor_ struct {
}

var Companion_AuthorityDescriptor_ = CompanionStruct_AuthorityDescriptor_{}

type AuthorityDescriptor_AuthorityDescriptor struct {
	AuthorityId          _dafny.Sequence
	ProducerId           _dafny.Sequence
	ProducerVersion      _dafny.Sequence
	ImplementationDigest _dafny.Sequence
	SourceArtifactDigest _dafny.Sequence
	RegistryDigest       _dafny.Sequence
}

func (AuthorityDescriptor_AuthorityDescriptor) isAuthorityDescriptor() {}

func (CompanionStruct_AuthorityDescriptor_) Create_AuthorityDescriptor_(AuthorityId _dafny.Sequence, ProducerId _dafny.Sequence, ProducerVersion _dafny.Sequence, ImplementationDigest _dafny.Sequence, SourceArtifactDigest _dafny.Sequence, RegistryDigest _dafny.Sequence) AuthorityDescriptor {
	return AuthorityDescriptor{AuthorityDescriptor_AuthorityDescriptor{AuthorityId, ProducerId, ProducerVersion, ImplementationDigest, SourceArtifactDigest, RegistryDigest}}
}

func (_this AuthorityDescriptor) Is_AuthorityDescriptor() bool {
	_, ok := _this.Get_().(AuthorityDescriptor_AuthorityDescriptor)
	return ok
}

func (CompanionStruct_AuthorityDescriptor_) Default() AuthorityDescriptor {
	return Companion_AuthorityDescriptor_.Create_AuthorityDescriptor_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this AuthorityDescriptor) Dtor_authorityId() _dafny.Sequence {
	return _this.Get_().(AuthorityDescriptor_AuthorityDescriptor).AuthorityId
}

func (_this AuthorityDescriptor) Dtor_producerId() _dafny.Sequence {
	return _this.Get_().(AuthorityDescriptor_AuthorityDescriptor).ProducerId
}

func (_this AuthorityDescriptor) Dtor_producerVersion() _dafny.Sequence {
	return _this.Get_().(AuthorityDescriptor_AuthorityDescriptor).ProducerVersion
}

func (_this AuthorityDescriptor) Dtor_implementationDigest() _dafny.Sequence {
	return _this.Get_().(AuthorityDescriptor_AuthorityDescriptor).ImplementationDigest
}

func (_this AuthorityDescriptor) Dtor_sourceArtifactDigest() _dafny.Sequence {
	return _this.Get_().(AuthorityDescriptor_AuthorityDescriptor).SourceArtifactDigest
}

func (_this AuthorityDescriptor) Dtor_registryDigest() _dafny.Sequence {
	return _this.Get_().(AuthorityDescriptor_AuthorityDescriptor).RegistryDigest
}

func (_this AuthorityDescriptor) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case AuthorityDescriptor_AuthorityDescriptor:
		{
			return "ConfluxSemanticGraphContract.AuthorityDescriptor.AuthorityDescriptor" + "(" + data.AuthorityId.VerbatimString(true) + ", " + data.ProducerId.VerbatimString(true) + ", " + data.ProducerVersion.VerbatimString(true) + ", " + data.ImplementationDigest.VerbatimString(true) + ", " + data.SourceArtifactDigest.VerbatimString(true) + ", " + data.RegistryDigest.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this AuthorityDescriptor) Equals(other AuthorityDescriptor) bool {
	switch data1 := _this.Get_().(type) {
	case AuthorityDescriptor_AuthorityDescriptor:
		{
			data2, ok := other.Get_().(AuthorityDescriptor_AuthorityDescriptor)
			return ok && data1.AuthorityId.Equals(data2.AuthorityId) && data1.ProducerId.Equals(data2.ProducerId) && data1.ProducerVersion.Equals(data2.ProducerVersion) && data1.ImplementationDigest.Equals(data2.ImplementationDigest) && data1.SourceArtifactDigest.Equals(data2.SourceArtifactDigest) && data1.RegistryDigest.Equals(data2.RegistryDigest)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this AuthorityDescriptor) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(AuthorityDescriptor)
	return ok && _this.Equals(typed)
}

func Type_AuthorityDescriptor_() _dafny.TypeDescriptor {
	return type_AuthorityDescriptor_{}
}

type type_AuthorityDescriptor_ struct {
}

func (_this type_AuthorityDescriptor_) Default() interface{} {
	return Companion_AuthorityDescriptor_.Default()
}

func (_this type_AuthorityDescriptor_) String() string {
	return "ConfluxSemanticGraphContract.AuthorityDescriptor"
}
func (_this AuthorityDescriptor) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = AuthorityDescriptor{}

// End of datatype AuthorityDescriptor

// Definition of datatype EvidenceRecord
type EvidenceRecord struct {
	Data_EvidenceRecord_
}

func (_this EvidenceRecord) Get_() Data_EvidenceRecord_ {
	return _this.Data_EvidenceRecord_
}

type Data_EvidenceRecord_ interface {
	isEvidenceRecord()
}

type CompanionStruct_EvidenceRecord_ struct {
}

var Companion_EvidenceRecord_ = CompanionStruct_EvidenceRecord_{}

type EvidenceRecord_EvidenceRecord struct {
	Id                _dafny.Sequence
	SnapshotId        _dafny.Sequence
	Kind              EvidenceKind
	LogicalRootId     _dafny.Sequence
	LogicalUri        _dafny.Sequence
	Locator           _dafny.Sequence
	Range             m_ConfluxContract.Option
	SymbolId          _dafny.Sequence
	ArtifactDigest    _dafny.Sequence
	ProducerDigest    _dafny.Sequence
	Statement         _dafny.Sequence
	ParentEvidenceIds _dafny.Set
	ContentRevision   _dafny.Sequence
}

func (EvidenceRecord_EvidenceRecord) isEvidenceRecord() {}

func (CompanionStruct_EvidenceRecord_) Create_EvidenceRecord_(Id _dafny.Sequence, SnapshotId _dafny.Sequence, Kind EvidenceKind, LogicalRootId _dafny.Sequence, LogicalUri _dafny.Sequence, Locator _dafny.Sequence, Range m_ConfluxContract.Option, SymbolId _dafny.Sequence, ArtifactDigest _dafny.Sequence, ProducerDigest _dafny.Sequence, Statement _dafny.Sequence, ParentEvidenceIds _dafny.Set, ContentRevision _dafny.Sequence) EvidenceRecord {
	return EvidenceRecord{EvidenceRecord_EvidenceRecord{Id, SnapshotId, Kind, LogicalRootId, LogicalUri, Locator, Range, SymbolId, ArtifactDigest, ProducerDigest, Statement, ParentEvidenceIds, ContentRevision}}
}

func (_this EvidenceRecord) Is_EvidenceRecord() bool {
	_, ok := _this.Get_().(EvidenceRecord_EvidenceRecord)
	return ok
}

func (CompanionStruct_EvidenceRecord_) Default() EvidenceRecord {
	return Companion_EvidenceRecord_.Create_EvidenceRecord_(_dafny.EmptySeq, _dafny.EmptySeq, Companion_EvidenceKind_.Default(), _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, m_ConfluxContract.Companion_Option_.Default(), _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySet, _dafny.EmptySeq)
}

func (_this EvidenceRecord) Dtor_id() _dafny.Sequence {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).Id
}

func (_this EvidenceRecord) Dtor_snapshotId() _dafny.Sequence {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).SnapshotId
}

func (_this EvidenceRecord) Dtor_kind() EvidenceKind {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).Kind
}

func (_this EvidenceRecord) Dtor_logicalRootId() _dafny.Sequence {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).LogicalRootId
}

func (_this EvidenceRecord) Dtor_logicalUri() _dafny.Sequence {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).LogicalUri
}

func (_this EvidenceRecord) Dtor_locator() _dafny.Sequence {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).Locator
}

func (_this EvidenceRecord) Dtor_range() m_ConfluxContract.Option {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).Range
}

func (_this EvidenceRecord) Dtor_symbolId() _dafny.Sequence {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).SymbolId
}

func (_this EvidenceRecord) Dtor_artifactDigest() _dafny.Sequence {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).ArtifactDigest
}

func (_this EvidenceRecord) Dtor_producerDigest() _dafny.Sequence {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).ProducerDigest
}

func (_this EvidenceRecord) Dtor_statement() _dafny.Sequence {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).Statement
}

func (_this EvidenceRecord) Dtor_parentEvidenceIds() _dafny.Set {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).ParentEvidenceIds
}

func (_this EvidenceRecord) Dtor_contentRevision() _dafny.Sequence {
	return _this.Get_().(EvidenceRecord_EvidenceRecord).ContentRevision
}

func (_this EvidenceRecord) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case EvidenceRecord_EvidenceRecord:
		{
			return "ConfluxSemanticGraphContract.EvidenceRecord.EvidenceRecord" + "(" + data.Id.VerbatimString(true) + ", " + data.SnapshotId.VerbatimString(true) + ", " + _dafny.String(data.Kind) + ", " + data.LogicalRootId.VerbatimString(true) + ", " + data.LogicalUri.VerbatimString(true) + ", " + data.Locator.VerbatimString(true) + ", " + _dafny.String(data.Range) + ", " + data.SymbolId.VerbatimString(true) + ", " + data.ArtifactDigest.VerbatimString(true) + ", " + data.ProducerDigest.VerbatimString(true) + ", " + data.Statement.VerbatimString(true) + ", " + _dafny.String(data.ParentEvidenceIds) + ", " + data.ContentRevision.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this EvidenceRecord) Equals(other EvidenceRecord) bool {
	switch data1 := _this.Get_().(type) {
	case EvidenceRecord_EvidenceRecord:
		{
			data2, ok := other.Get_().(EvidenceRecord_EvidenceRecord)
			return ok && data1.Id.Equals(data2.Id) && data1.SnapshotId.Equals(data2.SnapshotId) && data1.Kind.Equals(data2.Kind) && data1.LogicalRootId.Equals(data2.LogicalRootId) && data1.LogicalUri.Equals(data2.LogicalUri) && data1.Locator.Equals(data2.Locator) && data1.Range.Equals(data2.Range) && data1.SymbolId.Equals(data2.SymbolId) && data1.ArtifactDigest.Equals(data2.ArtifactDigest) && data1.ProducerDigest.Equals(data2.ProducerDigest) && data1.Statement.Equals(data2.Statement) && data1.ParentEvidenceIds.Equals(data2.ParentEvidenceIds) && data1.ContentRevision.Equals(data2.ContentRevision)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this EvidenceRecord) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(EvidenceRecord)
	return ok && _this.Equals(typed)
}

func Type_EvidenceRecord_() _dafny.TypeDescriptor {
	return type_EvidenceRecord_{}
}

type type_EvidenceRecord_ struct {
}

func (_this type_EvidenceRecord_) Default() interface{} {
	return Companion_EvidenceRecord_.Default()
}

func (_this type_EvidenceRecord_) String() string {
	return "ConfluxSemanticGraphContract.EvidenceRecord"
}
func (_this EvidenceRecord) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = EvidenceRecord{}

// End of datatype EvidenceRecord

// Definition of datatype SemanticNode
type SemanticNode struct {
	Data_SemanticNode_
}

func (_this SemanticNode) Get_() Data_SemanticNode_ {
	return _this.Data_SemanticNode_
}

type Data_SemanticNode_ interface {
	isSemanticNode()
}

type CompanionStruct_SemanticNode_ struct {
}

var Companion_SemanticNode_ = CompanionStruct_SemanticNode_{}

type SemanticNode_SemanticNode struct {
	Id              _dafny.Sequence
	Kind            _dafny.Sequence
	DisplayLabel    _dafny.Sequence
	Attributes      _dafny.Sequence
	FactIds         _dafny.Set
	EvidenceIds     _dafny.Set
	ContentRevision _dafny.Sequence
}

func (SemanticNode_SemanticNode) isSemanticNode() {}

func (CompanionStruct_SemanticNode_) Create_SemanticNode_(Id _dafny.Sequence, Kind _dafny.Sequence, DisplayLabel _dafny.Sequence, Attributes _dafny.Sequence, FactIds _dafny.Set, EvidenceIds _dafny.Set, ContentRevision _dafny.Sequence) SemanticNode {
	return SemanticNode{SemanticNode_SemanticNode{Id, Kind, DisplayLabel, Attributes, FactIds, EvidenceIds, ContentRevision}}
}

func (_this SemanticNode) Is_SemanticNode() bool {
	_, ok := _this.Get_().(SemanticNode_SemanticNode)
	return ok
}

func (CompanionStruct_SemanticNode_) Default() SemanticNode {
	return Companion_SemanticNode_.Create_SemanticNode_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySet, _dafny.EmptySet, _dafny.EmptySeq)
}

func (_this SemanticNode) Dtor_id() _dafny.Sequence {
	return _this.Get_().(SemanticNode_SemanticNode).Id
}

func (_this SemanticNode) Dtor_kind() _dafny.Sequence {
	return _this.Get_().(SemanticNode_SemanticNode).Kind
}

func (_this SemanticNode) Dtor_displayLabel() _dafny.Sequence {
	return _this.Get_().(SemanticNode_SemanticNode).DisplayLabel
}

func (_this SemanticNode) Dtor_attributes() _dafny.Sequence {
	return _this.Get_().(SemanticNode_SemanticNode).Attributes
}

func (_this SemanticNode) Dtor_factIds() _dafny.Set {
	return _this.Get_().(SemanticNode_SemanticNode).FactIds
}

func (_this SemanticNode) Dtor_evidenceIds() _dafny.Set {
	return _this.Get_().(SemanticNode_SemanticNode).EvidenceIds
}

func (_this SemanticNode) Dtor_contentRevision() _dafny.Sequence {
	return _this.Get_().(SemanticNode_SemanticNode).ContentRevision
}

func (_this SemanticNode) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SemanticNode_SemanticNode:
		{
			return "ConfluxSemanticGraphContract.SemanticNode.SemanticNode" + "(" + data.Id.VerbatimString(true) + ", " + data.Kind.VerbatimString(true) + ", " + data.DisplayLabel.VerbatimString(true) + ", " + _dafny.String(data.Attributes) + ", " + _dafny.String(data.FactIds) + ", " + _dafny.String(data.EvidenceIds) + ", " + data.ContentRevision.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SemanticNode) Equals(other SemanticNode) bool {
	switch data1 := _this.Get_().(type) {
	case SemanticNode_SemanticNode:
		{
			data2, ok := other.Get_().(SemanticNode_SemanticNode)
			return ok && data1.Id.Equals(data2.Id) && data1.Kind.Equals(data2.Kind) && data1.DisplayLabel.Equals(data2.DisplayLabel) && data1.Attributes.Equals(data2.Attributes) && data1.FactIds.Equals(data2.FactIds) && data1.EvidenceIds.Equals(data2.EvidenceIds) && data1.ContentRevision.Equals(data2.ContentRevision)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SemanticNode) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SemanticNode)
	return ok && _this.Equals(typed)
}

func Type_SemanticNode_() _dafny.TypeDescriptor {
	return type_SemanticNode_{}
}

type type_SemanticNode_ struct {
}

func (_this type_SemanticNode_) Default() interface{} {
	return Companion_SemanticNode_.Default()
}

func (_this type_SemanticNode_) String() string {
	return "ConfluxSemanticGraphContract.SemanticNode"
}
func (_this SemanticNode) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SemanticNode{}

// End of datatype SemanticNode

// Definition of datatype SemanticEdge
type SemanticEdge struct {
	Data_SemanticEdge_
}

func (_this SemanticEdge) Get_() Data_SemanticEdge_ {
	return _this.Data_SemanticEdge_
}

type Data_SemanticEdge_ interface {
	isSemanticEdge()
}

type CompanionStruct_SemanticEdge_ struct {
}

var Companion_SemanticEdge_ = CompanionStruct_SemanticEdge_{}

type SemanticEdge_SemanticEdge struct {
	Id              _dafny.Sequence
	Kind            _dafny.Sequence
	FromId          _dafny.Sequence
	ToId            _dafny.Sequence
	Directed        bool
	Attributes      _dafny.Sequence
	FactIds         _dafny.Set
	EvidenceIds     _dafny.Set
	ContentRevision _dafny.Sequence
}

func (SemanticEdge_SemanticEdge) isSemanticEdge() {}

func (CompanionStruct_SemanticEdge_) Create_SemanticEdge_(Id _dafny.Sequence, Kind _dafny.Sequence, FromId _dafny.Sequence, ToId _dafny.Sequence, Directed bool, Attributes _dafny.Sequence, FactIds _dafny.Set, EvidenceIds _dafny.Set, ContentRevision _dafny.Sequence) SemanticEdge {
	return SemanticEdge{SemanticEdge_SemanticEdge{Id, Kind, FromId, ToId, Directed, Attributes, FactIds, EvidenceIds, ContentRevision}}
}

func (_this SemanticEdge) Is_SemanticEdge() bool {
	_, ok := _this.Get_().(SemanticEdge_SemanticEdge)
	return ok
}

func (CompanionStruct_SemanticEdge_) Default() SemanticEdge {
	return Companion_SemanticEdge_.Create_SemanticEdge_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, false, _dafny.EmptySeq, _dafny.EmptySet, _dafny.EmptySet, _dafny.EmptySeq)
}

func (_this SemanticEdge) Dtor_id() _dafny.Sequence {
	return _this.Get_().(SemanticEdge_SemanticEdge).Id
}

func (_this SemanticEdge) Dtor_kind() _dafny.Sequence {
	return _this.Get_().(SemanticEdge_SemanticEdge).Kind
}

func (_this SemanticEdge) Dtor_fromId() _dafny.Sequence {
	return _this.Get_().(SemanticEdge_SemanticEdge).FromId
}

func (_this SemanticEdge) Dtor_toId() _dafny.Sequence {
	return _this.Get_().(SemanticEdge_SemanticEdge).ToId
}

func (_this SemanticEdge) Dtor_directed() bool {
	return _this.Get_().(SemanticEdge_SemanticEdge).Directed
}

func (_this SemanticEdge) Dtor_attributes() _dafny.Sequence {
	return _this.Get_().(SemanticEdge_SemanticEdge).Attributes
}

func (_this SemanticEdge) Dtor_factIds() _dafny.Set {
	return _this.Get_().(SemanticEdge_SemanticEdge).FactIds
}

func (_this SemanticEdge) Dtor_evidenceIds() _dafny.Set {
	return _this.Get_().(SemanticEdge_SemanticEdge).EvidenceIds
}

func (_this SemanticEdge) Dtor_contentRevision() _dafny.Sequence {
	return _this.Get_().(SemanticEdge_SemanticEdge).ContentRevision
}

func (_this SemanticEdge) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SemanticEdge_SemanticEdge:
		{
			return "ConfluxSemanticGraphContract.SemanticEdge.SemanticEdge" + "(" + data.Id.VerbatimString(true) + ", " + data.Kind.VerbatimString(true) + ", " + data.FromId.VerbatimString(true) + ", " + data.ToId.VerbatimString(true) + ", " + _dafny.String(data.Directed) + ", " + _dafny.String(data.Attributes) + ", " + _dafny.String(data.FactIds) + ", " + _dafny.String(data.EvidenceIds) + ", " + data.ContentRevision.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SemanticEdge) Equals(other SemanticEdge) bool {
	switch data1 := _this.Get_().(type) {
	case SemanticEdge_SemanticEdge:
		{
			data2, ok := other.Get_().(SemanticEdge_SemanticEdge)
			return ok && data1.Id.Equals(data2.Id) && data1.Kind.Equals(data2.Kind) && data1.FromId.Equals(data2.FromId) && data1.ToId.Equals(data2.ToId) && data1.Directed == data2.Directed && data1.Attributes.Equals(data2.Attributes) && data1.FactIds.Equals(data2.FactIds) && data1.EvidenceIds.Equals(data2.EvidenceIds) && data1.ContentRevision.Equals(data2.ContentRevision)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SemanticEdge) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SemanticEdge)
	return ok && _this.Equals(typed)
}

func Type_SemanticEdge_() _dafny.TypeDescriptor {
	return type_SemanticEdge_{}
}

type type_SemanticEdge_ struct {
}

func (_this type_SemanticEdge_) Default() interface{} {
	return Companion_SemanticEdge_.Default()
}

func (_this type_SemanticEdge_) String() string {
	return "ConfluxSemanticGraphContract.SemanticEdge"
}
func (_this SemanticEdge) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SemanticEdge{}

// End of datatype SemanticEdge

// Definition of datatype SemanticFact
type SemanticFact struct {
	Data_SemanticFact_
}

func (_this SemanticFact) Get_() Data_SemanticFact_ {
	return _this.Data_SemanticFact_
}

type Data_SemanticFact_ interface {
	isSemanticFact()
}

type CompanionStruct_SemanticFact_ struct {
}

var Companion_SemanticFact_ = CompanionStruct_SemanticFact_{}

type SemanticFact_SemanticFact struct {
	Id                _dafny.Sequence
	AnalysisId        _dafny.Sequence
	SubjectId         _dafny.Sequence
	Key               _dafny.Sequence
	Value             TypedValue
	ValueSchema       _dafny.Sequence
	Layer             FactLayer
	RuleId            _dafny.Sequence
	SupportingFactIds _dafny.Set
	EvidenceIds       _dafny.Set
	ContentRevision   _dafny.Sequence
}

func (SemanticFact_SemanticFact) isSemanticFact() {}

func (CompanionStruct_SemanticFact_) Create_SemanticFact_(Id _dafny.Sequence, AnalysisId _dafny.Sequence, SubjectId _dafny.Sequence, Key _dafny.Sequence, Value TypedValue, ValueSchema _dafny.Sequence, Layer FactLayer, RuleId _dafny.Sequence, SupportingFactIds _dafny.Set, EvidenceIds _dafny.Set, ContentRevision _dafny.Sequence) SemanticFact {
	return SemanticFact{SemanticFact_SemanticFact{Id, AnalysisId, SubjectId, Key, Value, ValueSchema, Layer, RuleId, SupportingFactIds, EvidenceIds, ContentRevision}}
}

func (_this SemanticFact) Is_SemanticFact() bool {
	_, ok := _this.Get_().(SemanticFact_SemanticFact)
	return ok
}

func (CompanionStruct_SemanticFact_) Default() SemanticFact {
	return Companion_SemanticFact_.Create_SemanticFact_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, Companion_TypedValue_.Default(), _dafny.EmptySeq, Companion_FactLayer_.Default(), _dafny.EmptySeq, _dafny.EmptySet, _dafny.EmptySet, _dafny.EmptySeq)
}

func (_this SemanticFact) Dtor_id() _dafny.Sequence {
	return _this.Get_().(SemanticFact_SemanticFact).Id
}

func (_this SemanticFact) Dtor_analysisId() _dafny.Sequence {
	return _this.Get_().(SemanticFact_SemanticFact).AnalysisId
}

func (_this SemanticFact) Dtor_subjectId() _dafny.Sequence {
	return _this.Get_().(SemanticFact_SemanticFact).SubjectId
}

func (_this SemanticFact) Dtor_key() _dafny.Sequence {
	return _this.Get_().(SemanticFact_SemanticFact).Key
}

func (_this SemanticFact) Dtor_value() TypedValue {
	return _this.Get_().(SemanticFact_SemanticFact).Value
}

func (_this SemanticFact) Dtor_valueSchema() _dafny.Sequence {
	return _this.Get_().(SemanticFact_SemanticFact).ValueSchema
}

func (_this SemanticFact) Dtor_layer() FactLayer {
	return _this.Get_().(SemanticFact_SemanticFact).Layer
}

func (_this SemanticFact) Dtor_ruleId() _dafny.Sequence {
	return _this.Get_().(SemanticFact_SemanticFact).RuleId
}

func (_this SemanticFact) Dtor_supportingFactIds() _dafny.Set {
	return _this.Get_().(SemanticFact_SemanticFact).SupportingFactIds
}

func (_this SemanticFact) Dtor_evidenceIds() _dafny.Set {
	return _this.Get_().(SemanticFact_SemanticFact).EvidenceIds
}

func (_this SemanticFact) Dtor_contentRevision() _dafny.Sequence {
	return _this.Get_().(SemanticFact_SemanticFact).ContentRevision
}

func (_this SemanticFact) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SemanticFact_SemanticFact:
		{
			return "ConfluxSemanticGraphContract.SemanticFact.SemanticFact" + "(" + data.Id.VerbatimString(true) + ", " + data.AnalysisId.VerbatimString(true) + ", " + data.SubjectId.VerbatimString(true) + ", " + data.Key.VerbatimString(true) + ", " + _dafny.String(data.Value) + ", " + data.ValueSchema.VerbatimString(true) + ", " + _dafny.String(data.Layer) + ", " + data.RuleId.VerbatimString(true) + ", " + _dafny.String(data.SupportingFactIds) + ", " + _dafny.String(data.EvidenceIds) + ", " + data.ContentRevision.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SemanticFact) Equals(other SemanticFact) bool {
	switch data1 := _this.Get_().(type) {
	case SemanticFact_SemanticFact:
		{
			data2, ok := other.Get_().(SemanticFact_SemanticFact)
			return ok && data1.Id.Equals(data2.Id) && data1.AnalysisId.Equals(data2.AnalysisId) && data1.SubjectId.Equals(data2.SubjectId) && data1.Key.Equals(data2.Key) && data1.Value.Equals(data2.Value) && data1.ValueSchema.Equals(data2.ValueSchema) && data1.Layer.Equals(data2.Layer) && data1.RuleId.Equals(data2.RuleId) && data1.SupportingFactIds.Equals(data2.SupportingFactIds) && data1.EvidenceIds.Equals(data2.EvidenceIds) && data1.ContentRevision.Equals(data2.ContentRevision)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SemanticFact) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SemanticFact)
	return ok && _this.Equals(typed)
}

func Type_SemanticFact_() _dafny.TypeDescriptor {
	return type_SemanticFact_{}
}

type type_SemanticFact_ struct {
}

func (_this type_SemanticFact_) Default() interface{} {
	return Companion_SemanticFact_.Default()
}

func (_this type_SemanticFact_) String() string {
	return "ConfluxSemanticGraphContract.SemanticFact"
}
func (_this SemanticFact) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SemanticFact{}

// End of datatype SemanticFact

// Definition of datatype SemanticFinding
type SemanticFinding struct {
	Data_SemanticFinding_
}

func (_this SemanticFinding) Get_() Data_SemanticFinding_ {
	return _this.Data_SemanticFinding_
}

type Data_SemanticFinding_ interface {
	isSemanticFinding()
}

type CompanionStruct_SemanticFinding_ struct {
}

var Companion_SemanticFinding_ = CompanionStruct_SemanticFinding_{}

type SemanticFinding_SemanticFinding struct {
	Id                _dafny.Sequence
	AnalysisId        _dafny.Sequence
	RuleId            _dafny.Sequence
	SubjectId         _dafny.Sequence
	Severity          FindingSeverity
	Status            FindingStatus
	Code              _dafny.Sequence
	Message           _dafny.Sequence
	Expected          TypedValue
	Actual            TypedValue
	SupportingFactIds _dafny.Set
	EvidenceIds       _dafny.Set
	Blocking          bool
	ContentRevision   _dafny.Sequence
}

func (SemanticFinding_SemanticFinding) isSemanticFinding() {}

func (CompanionStruct_SemanticFinding_) Create_SemanticFinding_(Id _dafny.Sequence, AnalysisId _dafny.Sequence, RuleId _dafny.Sequence, SubjectId _dafny.Sequence, Severity FindingSeverity, Status FindingStatus, Code _dafny.Sequence, Message _dafny.Sequence, Expected TypedValue, Actual TypedValue, SupportingFactIds _dafny.Set, EvidenceIds _dafny.Set, Blocking bool, ContentRevision _dafny.Sequence) SemanticFinding {
	return SemanticFinding{SemanticFinding_SemanticFinding{Id, AnalysisId, RuleId, SubjectId, Severity, Status, Code, Message, Expected, Actual, SupportingFactIds, EvidenceIds, Blocking, ContentRevision}}
}

func (_this SemanticFinding) Is_SemanticFinding() bool {
	_, ok := _this.Get_().(SemanticFinding_SemanticFinding)
	return ok
}

func (CompanionStruct_SemanticFinding_) Default() SemanticFinding {
	return Companion_SemanticFinding_.Create_SemanticFinding_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, Companion_FindingSeverity_.Default(), Companion_FindingStatus_.Default(), _dafny.EmptySeq, _dafny.EmptySeq, Companion_TypedValue_.Default(), Companion_TypedValue_.Default(), _dafny.EmptySet, _dafny.EmptySet, false, _dafny.EmptySeq)
}

func (_this SemanticFinding) Dtor_id() _dafny.Sequence {
	return _this.Get_().(SemanticFinding_SemanticFinding).Id
}

func (_this SemanticFinding) Dtor_analysisId() _dafny.Sequence {
	return _this.Get_().(SemanticFinding_SemanticFinding).AnalysisId
}

func (_this SemanticFinding) Dtor_ruleId() _dafny.Sequence {
	return _this.Get_().(SemanticFinding_SemanticFinding).RuleId
}

func (_this SemanticFinding) Dtor_subjectId() _dafny.Sequence {
	return _this.Get_().(SemanticFinding_SemanticFinding).SubjectId
}

func (_this SemanticFinding) Dtor_severity() FindingSeverity {
	return _this.Get_().(SemanticFinding_SemanticFinding).Severity
}

func (_this SemanticFinding) Dtor_status() FindingStatus {
	return _this.Get_().(SemanticFinding_SemanticFinding).Status
}

func (_this SemanticFinding) Dtor_code() _dafny.Sequence {
	return _this.Get_().(SemanticFinding_SemanticFinding).Code
}

func (_this SemanticFinding) Dtor_message() _dafny.Sequence {
	return _this.Get_().(SemanticFinding_SemanticFinding).Message
}

func (_this SemanticFinding) Dtor_expected() TypedValue {
	return _this.Get_().(SemanticFinding_SemanticFinding).Expected
}

func (_this SemanticFinding) Dtor_actual() TypedValue {
	return _this.Get_().(SemanticFinding_SemanticFinding).Actual
}

func (_this SemanticFinding) Dtor_supportingFactIds() _dafny.Set {
	return _this.Get_().(SemanticFinding_SemanticFinding).SupportingFactIds
}

func (_this SemanticFinding) Dtor_evidenceIds() _dafny.Set {
	return _this.Get_().(SemanticFinding_SemanticFinding).EvidenceIds
}

func (_this SemanticFinding) Dtor_blocking() bool {
	return _this.Get_().(SemanticFinding_SemanticFinding).Blocking
}

func (_this SemanticFinding) Dtor_contentRevision() _dafny.Sequence {
	return _this.Get_().(SemanticFinding_SemanticFinding).ContentRevision
}

func (_this SemanticFinding) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SemanticFinding_SemanticFinding:
		{
			return "ConfluxSemanticGraphContract.SemanticFinding.SemanticFinding" + "(" + data.Id.VerbatimString(true) + ", " + data.AnalysisId.VerbatimString(true) + ", " + data.RuleId.VerbatimString(true) + ", " + data.SubjectId.VerbatimString(true) + ", " + _dafny.String(data.Severity) + ", " + _dafny.String(data.Status) + ", " + data.Code.VerbatimString(true) + ", " + data.Message.VerbatimString(true) + ", " + _dafny.String(data.Expected) + ", " + _dafny.String(data.Actual) + ", " + _dafny.String(data.SupportingFactIds) + ", " + _dafny.String(data.EvidenceIds) + ", " + _dafny.String(data.Blocking) + ", " + data.ContentRevision.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SemanticFinding) Equals(other SemanticFinding) bool {
	switch data1 := _this.Get_().(type) {
	case SemanticFinding_SemanticFinding:
		{
			data2, ok := other.Get_().(SemanticFinding_SemanticFinding)
			return ok && data1.Id.Equals(data2.Id) && data1.AnalysisId.Equals(data2.AnalysisId) && data1.RuleId.Equals(data2.RuleId) && data1.SubjectId.Equals(data2.SubjectId) && data1.Severity.Equals(data2.Severity) && data1.Status.Equals(data2.Status) && data1.Code.Equals(data2.Code) && data1.Message.Equals(data2.Message) && data1.Expected.Equals(data2.Expected) && data1.Actual.Equals(data2.Actual) && data1.SupportingFactIds.Equals(data2.SupportingFactIds) && data1.EvidenceIds.Equals(data2.EvidenceIds) && data1.Blocking == data2.Blocking && data1.ContentRevision.Equals(data2.ContentRevision)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SemanticFinding) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SemanticFinding)
	return ok && _this.Equals(typed)
}

func Type_SemanticFinding_() _dafny.TypeDescriptor {
	return type_SemanticFinding_{}
}

type type_SemanticFinding_ struct {
}

func (_this type_SemanticFinding_) Default() interface{} {
	return Companion_SemanticFinding_.Default()
}

func (_this type_SemanticFinding_) String() string {
	return "ConfluxSemanticGraphContract.SemanticFinding"
}
func (_this SemanticFinding) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SemanticFinding{}

// End of datatype SemanticFinding

// Definition of datatype GraphLensDescriptor
type GraphLensDescriptor struct {
	Data_GraphLensDescriptor_
}

func (_this GraphLensDescriptor) Get_() Data_GraphLensDescriptor_ {
	return _this.Data_GraphLensDescriptor_
}

type Data_GraphLensDescriptor_ interface {
	isGraphLensDescriptor()
}

type CompanionStruct_GraphLensDescriptor_ struct {
}

var Companion_GraphLensDescriptor_ = CompanionStruct_GraphLensDescriptor_{}

type GraphLensDescriptor_GraphLensDescriptor struct {
	Id                   _dafny.Sequence
	Version              _dafny.Sequence
	Archetype            LensArchetype
	Title                _dafny.Sequence
	NodeSelector         _dafny.Sequence
	EdgeSelector         _dafny.Sequence
	FactSelector         _dafny.Sequence
	LabelField           _dafny.Sequence
	DetailFields         _dafny.Sequence
	Columns              _dafny.Sequence
	RequiredCapabilities _dafny.Sequence
	ContentRevision      _dafny.Sequence
}

func (GraphLensDescriptor_GraphLensDescriptor) isGraphLensDescriptor() {}

func (CompanionStruct_GraphLensDescriptor_) Create_GraphLensDescriptor_(Id _dafny.Sequence, Version _dafny.Sequence, Archetype LensArchetype, Title _dafny.Sequence, NodeSelector _dafny.Sequence, EdgeSelector _dafny.Sequence, FactSelector _dafny.Sequence, LabelField _dafny.Sequence, DetailFields _dafny.Sequence, Columns _dafny.Sequence, RequiredCapabilities _dafny.Sequence, ContentRevision _dafny.Sequence) GraphLensDescriptor {
	return GraphLensDescriptor{GraphLensDescriptor_GraphLensDescriptor{Id, Version, Archetype, Title, NodeSelector, EdgeSelector, FactSelector, LabelField, DetailFields, Columns, RequiredCapabilities, ContentRevision}}
}

func (_this GraphLensDescriptor) Is_GraphLensDescriptor() bool {
	_, ok := _this.Get_().(GraphLensDescriptor_GraphLensDescriptor)
	return ok
}

func (CompanionStruct_GraphLensDescriptor_) Default() GraphLensDescriptor {
	return Companion_GraphLensDescriptor_.Create_GraphLensDescriptor_(_dafny.EmptySeq, _dafny.EmptySeq, Companion_LensArchetype_.Default(), _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this GraphLensDescriptor) Dtor_id() _dafny.Sequence {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).Id
}

func (_this GraphLensDescriptor) Dtor_version() _dafny.Sequence {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).Version
}

func (_this GraphLensDescriptor) Dtor_archetype() LensArchetype {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).Archetype
}

func (_this GraphLensDescriptor) Dtor_title() _dafny.Sequence {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).Title
}

func (_this GraphLensDescriptor) Dtor_nodeSelector() _dafny.Sequence {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).NodeSelector
}

func (_this GraphLensDescriptor) Dtor_edgeSelector() _dafny.Sequence {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).EdgeSelector
}

func (_this GraphLensDescriptor) Dtor_factSelector() _dafny.Sequence {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).FactSelector
}

func (_this GraphLensDescriptor) Dtor_labelField() _dafny.Sequence {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).LabelField
}

func (_this GraphLensDescriptor) Dtor_detailFields() _dafny.Sequence {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).DetailFields
}

func (_this GraphLensDescriptor) Dtor_columns() _dafny.Sequence {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).Columns
}

func (_this GraphLensDescriptor) Dtor_requiredCapabilities() _dafny.Sequence {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).RequiredCapabilities
}

func (_this GraphLensDescriptor) Dtor_contentRevision() _dafny.Sequence {
	return _this.Get_().(GraphLensDescriptor_GraphLensDescriptor).ContentRevision
}

func (_this GraphLensDescriptor) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case GraphLensDescriptor_GraphLensDescriptor:
		{
			return "ConfluxSemanticGraphContract.GraphLensDescriptor.GraphLensDescriptor" + "(" + data.Id.VerbatimString(true) + ", " + data.Version.VerbatimString(true) + ", " + _dafny.String(data.Archetype) + ", " + data.Title.VerbatimString(true) + ", " + data.NodeSelector.VerbatimString(true) + ", " + data.EdgeSelector.VerbatimString(true) + ", " + data.FactSelector.VerbatimString(true) + ", " + data.LabelField.VerbatimString(true) + ", " + _dafny.String(data.DetailFields) + ", " + _dafny.String(data.Columns) + ", " + _dafny.String(data.RequiredCapabilities) + ", " + data.ContentRevision.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GraphLensDescriptor) Equals(other GraphLensDescriptor) bool {
	switch data1 := _this.Get_().(type) {
	case GraphLensDescriptor_GraphLensDescriptor:
		{
			data2, ok := other.Get_().(GraphLensDescriptor_GraphLensDescriptor)
			return ok && data1.Id.Equals(data2.Id) && data1.Version.Equals(data2.Version) && data1.Archetype.Equals(data2.Archetype) && data1.Title.Equals(data2.Title) && data1.NodeSelector.Equals(data2.NodeSelector) && data1.EdgeSelector.Equals(data2.EdgeSelector) && data1.FactSelector.Equals(data2.FactSelector) && data1.LabelField.Equals(data2.LabelField) && data1.DetailFields.Equals(data2.DetailFields) && data1.Columns.Equals(data2.Columns) && data1.RequiredCapabilities.Equals(data2.RequiredCapabilities) && data1.ContentRevision.Equals(data2.ContentRevision)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GraphLensDescriptor) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GraphLensDescriptor)
	return ok && _this.Equals(typed)
}

func Type_GraphLensDescriptor_() _dafny.TypeDescriptor {
	return type_GraphLensDescriptor_{}
}

type type_GraphLensDescriptor_ struct {
}

func (_this type_GraphLensDescriptor_) Default() interface{} {
	return Companion_GraphLensDescriptor_.Default()
}

func (_this type_GraphLensDescriptor_) String() string {
	return "ConfluxSemanticGraphContract.GraphLensDescriptor"
}
func (_this GraphLensDescriptor) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = GraphLensDescriptor{}

// End of datatype GraphLensDescriptor

// Definition of datatype SemanticGraphSnapshotV1
type SemanticGraphSnapshotV1 struct {
	Data_SemanticGraphSnapshotV1_
}

func (_this SemanticGraphSnapshotV1) Get_() Data_SemanticGraphSnapshotV1_ {
	return _this.Data_SemanticGraphSnapshotV1_
}

type Data_SemanticGraphSnapshotV1_ interface {
	isSemanticGraphSnapshotV1()
}

type CompanionStruct_SemanticGraphSnapshotV1_ struct {
}

var Companion_SemanticGraphSnapshotV1_ = CompanionStruct_SemanticGraphSnapshotV1_{}

type SemanticGraphSnapshotV1_SemanticGraphSnapshotV1 struct {
	SchemaVersion    _dafny.Sequence
	GraphId          _dafny.Sequence
	Revision         _dafny.Sequence
	PreviousRevision _dafny.Sequence
	SnapshotId       _dafny.Sequence
	Authority        AuthorityDescriptor
	Capabilities     _dafny.Map
	Lenses           _dafny.Map
	Nodes            _dafny.Map
	Edges            _dafny.Map
	Facts            _dafny.Map
	Findings         _dafny.Map
	Evidence         _dafny.Map
	Completeness     Completeness
	ReceiptId        _dafny.Sequence
}

func (SemanticGraphSnapshotV1_SemanticGraphSnapshotV1) isSemanticGraphSnapshotV1() {}

func (CompanionStruct_SemanticGraphSnapshotV1_) Create_SemanticGraphSnapshotV1_(SchemaVersion _dafny.Sequence, GraphId _dafny.Sequence, Revision _dafny.Sequence, PreviousRevision _dafny.Sequence, SnapshotId _dafny.Sequence, Authority AuthorityDescriptor, Capabilities _dafny.Map, Lenses _dafny.Map, Nodes _dafny.Map, Edges _dafny.Map, Facts _dafny.Map, Findings _dafny.Map, Evidence _dafny.Map, Completeness Completeness, ReceiptId _dafny.Sequence) SemanticGraphSnapshotV1 {
	return SemanticGraphSnapshotV1{SemanticGraphSnapshotV1_SemanticGraphSnapshotV1{SchemaVersion, GraphId, Revision, PreviousRevision, SnapshotId, Authority, Capabilities, Lenses, Nodes, Edges, Facts, Findings, Evidence, Completeness, ReceiptId}}
}

func (_this SemanticGraphSnapshotV1) Is_SemanticGraphSnapshotV1() bool {
	_, ok := _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1)
	return ok
}

func (CompanionStruct_SemanticGraphSnapshotV1_) Default() SemanticGraphSnapshotV1 {
	return Companion_SemanticGraphSnapshotV1_.Create_SemanticGraphSnapshotV1_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, Companion_AuthorityDescriptor_.Default(), _dafny.EmptyMap, _dafny.EmptyMap, _dafny.EmptyMap, _dafny.EmptyMap, _dafny.EmptyMap, _dafny.EmptyMap, _dafny.EmptyMap, Companion_Completeness_.Default(), _dafny.EmptySeq)
}

func (_this SemanticGraphSnapshotV1) Dtor_schemaVersion() _dafny.Sequence {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).SchemaVersion
}

func (_this SemanticGraphSnapshotV1) Dtor_graphId() _dafny.Sequence {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).GraphId
}

func (_this SemanticGraphSnapshotV1) Dtor_revision() _dafny.Sequence {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).Revision
}

func (_this SemanticGraphSnapshotV1) Dtor_previousRevision() _dafny.Sequence {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).PreviousRevision
}

func (_this SemanticGraphSnapshotV1) Dtor_snapshotId() _dafny.Sequence {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).SnapshotId
}

func (_this SemanticGraphSnapshotV1) Dtor_authority() AuthorityDescriptor {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).Authority
}

func (_this SemanticGraphSnapshotV1) Dtor_capabilities() _dafny.Map {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).Capabilities
}

func (_this SemanticGraphSnapshotV1) Dtor_lenses() _dafny.Map {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).Lenses
}

func (_this SemanticGraphSnapshotV1) Dtor_nodes() _dafny.Map {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).Nodes
}

func (_this SemanticGraphSnapshotV1) Dtor_edges() _dafny.Map {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).Edges
}

func (_this SemanticGraphSnapshotV1) Dtor_facts() _dafny.Map {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).Facts
}

func (_this SemanticGraphSnapshotV1) Dtor_findings() _dafny.Map {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).Findings
}

func (_this SemanticGraphSnapshotV1) Dtor_evidence() _dafny.Map {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).Evidence
}

func (_this SemanticGraphSnapshotV1) Dtor_completeness() Completeness {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).Completeness
}

func (_this SemanticGraphSnapshotV1) Dtor_receiptId() _dafny.Sequence {
	return _this.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1).ReceiptId
}

func (_this SemanticGraphSnapshotV1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SemanticGraphSnapshotV1_SemanticGraphSnapshotV1:
		{
			return "ConfluxSemanticGraphContract.SemanticGraphSnapshotV1.SemanticGraphSnapshotV1" + "(" + data.SchemaVersion.VerbatimString(true) + ", " + data.GraphId.VerbatimString(true) + ", " + data.Revision.VerbatimString(true) + ", " + data.PreviousRevision.VerbatimString(true) + ", " + data.SnapshotId.VerbatimString(true) + ", " + _dafny.String(data.Authority) + ", " + _dafny.String(data.Capabilities) + ", " + _dafny.String(data.Lenses) + ", " + _dafny.String(data.Nodes) + ", " + _dafny.String(data.Edges) + ", " + _dafny.String(data.Facts) + ", " + _dafny.String(data.Findings) + ", " + _dafny.String(data.Evidence) + ", " + _dafny.String(data.Completeness) + ", " + data.ReceiptId.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SemanticGraphSnapshotV1) Equals(other SemanticGraphSnapshotV1) bool {
	switch data1 := _this.Get_().(type) {
	case SemanticGraphSnapshotV1_SemanticGraphSnapshotV1:
		{
			data2, ok := other.Get_().(SemanticGraphSnapshotV1_SemanticGraphSnapshotV1)
			return ok && data1.SchemaVersion.Equals(data2.SchemaVersion) && data1.GraphId.Equals(data2.GraphId) && data1.Revision.Equals(data2.Revision) && data1.PreviousRevision.Equals(data2.PreviousRevision) && data1.SnapshotId.Equals(data2.SnapshotId) && data1.Authority.Equals(data2.Authority) && data1.Capabilities.Equals(data2.Capabilities) && data1.Lenses.Equals(data2.Lenses) && data1.Nodes.Equals(data2.Nodes) && data1.Edges.Equals(data2.Edges) && data1.Facts.Equals(data2.Facts) && data1.Findings.Equals(data2.Findings) && data1.Evidence.Equals(data2.Evidence) && data1.Completeness.Equals(data2.Completeness) && data1.ReceiptId.Equals(data2.ReceiptId)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SemanticGraphSnapshotV1) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SemanticGraphSnapshotV1)
	return ok && _this.Equals(typed)
}

func Type_SemanticGraphSnapshotV1_() _dafny.TypeDescriptor {
	return type_SemanticGraphSnapshotV1_{}
}

type type_SemanticGraphSnapshotV1_ struct {
}

func (_this type_SemanticGraphSnapshotV1_) Default() interface{} {
	return Companion_SemanticGraphSnapshotV1_.Default()
}

func (_this type_SemanticGraphSnapshotV1_) String() string {
	return "ConfluxSemanticGraphContract.SemanticGraphSnapshotV1"
}
func (_this SemanticGraphSnapshotV1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SemanticGraphSnapshotV1{}

// End of datatype SemanticGraphSnapshotV1

// Definition of datatype ResetReason
type ResetReason struct {
	Data_ResetReason_
}

func (_this ResetReason) Get_() Data_ResetReason_ {
	return _this.Data_ResetReason_
}

type Data_ResetReason_ interface {
	isResetReason()
}

type CompanionStruct_ResetReason_ struct {
}

var Companion_ResetReason_ = CompanionStruct_ResetReason_{}

type ResetReason_UnknownRevision struct {
}

func (ResetReason_UnknownRevision) isResetReason() {}

func (CompanionStruct_ResetReason_) Create_UnknownRevision_() ResetReason {
	return ResetReason{ResetReason_UnknownRevision{}}
}

func (_this ResetReason) Is_UnknownRevision() bool {
	_, ok := _this.Get_().(ResetReason_UnknownRevision)
	return ok
}

type ResetReason_Expired struct {
}

func (ResetReason_Expired) isResetReason() {}

func (CompanionStruct_ResetReason_) Create_Expired_() ResetReason {
	return ResetReason{ResetReason_Expired{}}
}

func (_this ResetReason) Is_Expired() bool {
	_, ok := _this.Get_().(ResetReason_Expired)
	return ok
}

type ResetReason_GraphMismatch struct {
}

func (ResetReason_GraphMismatch) isResetReason() {}

func (CompanionStruct_ResetReason_) Create_GraphMismatch_() ResetReason {
	return ResetReason{ResetReason_GraphMismatch{}}
}

func (_this ResetReason) Is_GraphMismatch() bool {
	_, ok := _this.Get_().(ResetReason_GraphMismatch)
	return ok
}

type ResetReason_SchemaEpoch struct {
}

func (ResetReason_SchemaEpoch) isResetReason() {}

func (CompanionStruct_ResetReason_) Create_SchemaEpoch_() ResetReason {
	return ResetReason{ResetReason_SchemaEpoch{}}
}

func (_this ResetReason) Is_SchemaEpoch() bool {
	_, ok := _this.Get_().(ResetReason_SchemaEpoch)
	return ok
}

type ResetReason_CorruptBase struct {
}

func (ResetReason_CorruptBase) isResetReason() {}

func (CompanionStruct_ResetReason_) Create_CorruptBase_() ResetReason {
	return ResetReason{ResetReason_CorruptBase{}}
}

func (_this ResetReason) Is_CorruptBase() bool {
	_, ok := _this.Get_().(ResetReason_CorruptBase)
	return ok
}

func (CompanionStruct_ResetReason_) Default() ResetReason {
	return Companion_ResetReason_.Create_UnknownRevision_()
}

func (_ CompanionStruct_ResetReason_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_ResetReason_.Create_UnknownRevision_(), true
		case 1:
			return Companion_ResetReason_.Create_Expired_(), true
		case 2:
			return Companion_ResetReason_.Create_GraphMismatch_(), true
		case 3:
			return Companion_ResetReason_.Create_SchemaEpoch_(), true
		case 4:
			return Companion_ResetReason_.Create_CorruptBase_(), true
		default:
			return ResetReason{}, false
		}
	}
}

func (_this ResetReason) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case ResetReason_UnknownRevision:
		{
			return "ConfluxSemanticGraphContract.ResetReason.UnknownRevision"
		}
	case ResetReason_Expired:
		{
			return "ConfluxSemanticGraphContract.ResetReason.Expired"
		}
	case ResetReason_GraphMismatch:
		{
			return "ConfluxSemanticGraphContract.ResetReason.GraphMismatch"
		}
	case ResetReason_SchemaEpoch:
		{
			return "ConfluxSemanticGraphContract.ResetReason.SchemaEpoch"
		}
	case ResetReason_CorruptBase:
		{
			return "ConfluxSemanticGraphContract.ResetReason.CorruptBase"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ResetReason) Equals(other ResetReason) bool {
	switch _this.Get_().(type) {
	case ResetReason_UnknownRevision:
		{
			_, ok := other.Get_().(ResetReason_UnknownRevision)
			return ok
		}
	case ResetReason_Expired:
		{
			_, ok := other.Get_().(ResetReason_Expired)
			return ok
		}
	case ResetReason_GraphMismatch:
		{
			_, ok := other.Get_().(ResetReason_GraphMismatch)
			return ok
		}
	case ResetReason_SchemaEpoch:
		{
			_, ok := other.Get_().(ResetReason_SchemaEpoch)
			return ok
		}
	case ResetReason_CorruptBase:
		{
			_, ok := other.Get_().(ResetReason_CorruptBase)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ResetReason) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ResetReason)
	return ok && _this.Equals(typed)
}

func Type_ResetReason_() _dafny.TypeDescriptor {
	return type_ResetReason_{}
}

type type_ResetReason_ struct {
}

func (_this type_ResetReason_) Default() interface{} {
	return Companion_ResetReason_.Default()
}

func (_this type_ResetReason_) String() string {
	return "ConfluxSemanticGraphContract.ResetReason"
}
func (_this ResetReason) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ResetReason{}

// End of datatype ResetReason

// Definition of datatype GraphChanges
type GraphChanges struct {
	Data_GraphChanges_
}

func (_this GraphChanges) Get_() Data_GraphChanges_ {
	return _this.Data_GraphChanges_
}

type Data_GraphChanges_ interface {
	isGraphChanges()
}

type CompanionStruct_GraphChanges_ struct {
}

var Companion_GraphChanges_ = CompanionStruct_GraphChanges_{}

type GraphChanges_GraphChanges struct {
	Added   _dafny.Map
	Changed _dafny.Map
	Removed _dafny.Set
}

func (GraphChanges_GraphChanges) isGraphChanges() {}

func (CompanionStruct_GraphChanges_) Create_GraphChanges_(Added _dafny.Map, Changed _dafny.Map, Removed _dafny.Set) GraphChanges {
	return GraphChanges{GraphChanges_GraphChanges{Added, Changed, Removed}}
}

func (_this GraphChanges) Is_GraphChanges() bool {
	_, ok := _this.Get_().(GraphChanges_GraphChanges)
	return ok
}

func (CompanionStruct_GraphChanges_) Default() GraphChanges {
	return Companion_GraphChanges_.Create_GraphChanges_(_dafny.EmptyMap, _dafny.EmptyMap, _dafny.EmptySet)
}

func (_this GraphChanges) Dtor_added() _dafny.Map {
	return _this.Get_().(GraphChanges_GraphChanges).Added
}

func (_this GraphChanges) Dtor_changed() _dafny.Map {
	return _this.Get_().(GraphChanges_GraphChanges).Changed
}

func (_this GraphChanges) Dtor_removed() _dafny.Set {
	return _this.Get_().(GraphChanges_GraphChanges).Removed
}

func (_this GraphChanges) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case GraphChanges_GraphChanges:
		{
			return "ConfluxSemanticGraphContract.GraphChanges.GraphChanges" + "(" + _dafny.String(data.Added) + ", " + _dafny.String(data.Changed) + ", " + _dafny.String(data.Removed) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GraphChanges) Equals(other GraphChanges) bool {
	switch data1 := _this.Get_().(type) {
	case GraphChanges_GraphChanges:
		{
			data2, ok := other.Get_().(GraphChanges_GraphChanges)
			return ok && data1.Added.Equals(data2.Added) && data1.Changed.Equals(data2.Changed) && data1.Removed.Equals(data2.Removed)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GraphChanges) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GraphChanges)
	return ok && _this.Equals(typed)
}

func Type_GraphChanges_() _dafny.TypeDescriptor {
	return type_GraphChanges_{}
}

type type_GraphChanges_ struct {
}

func (_this type_GraphChanges_) Default() interface{} {
	return Companion_GraphChanges_.Default()
}

func (_this type_GraphChanges_) String() string {
	return "ConfluxSemanticGraphContract.GraphChanges"
}
func (_this GraphChanges) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = GraphChanges{}

// End of datatype GraphChanges

// Definition of datatype SemanticGraphDeltaV1
type SemanticGraphDeltaV1 struct {
	Data_SemanticGraphDeltaV1_
}

func (_this SemanticGraphDeltaV1) Get_() Data_SemanticGraphDeltaV1_ {
	return _this.Data_SemanticGraphDeltaV1_
}

type Data_SemanticGraphDeltaV1_ interface {
	isSemanticGraphDeltaV1()
}

type CompanionStruct_SemanticGraphDeltaV1_ struct {
}

var Companion_SemanticGraphDeltaV1_ = CompanionStruct_SemanticGraphDeltaV1_{}

type SemanticGraphDeltaV1_SemanticGraphDeltaV1 struct {
	SchemaVersion           _dafny.Sequence
	GraphId                 _dafny.Sequence
	FromRevision            _dafny.Sequence
	ToRevision              _dafny.Sequence
	Reset                   bool
	ResetReason             m_ConfluxContract.Option
	Replacement             m_ConfluxContract.Option
	Nodes                   GraphChanges
	Edges                   GraphChanges
	Facts                   GraphChanges
	Findings                GraphChanges
	Evidence                GraphChanges
	CapabilitiesReplacement m_ConfluxContract.Option
	LensesReplacement       m_ConfluxContract.Option
	ReceiptId               _dafny.Sequence
}

func (SemanticGraphDeltaV1_SemanticGraphDeltaV1) isSemanticGraphDeltaV1() {}

func (CompanionStruct_SemanticGraphDeltaV1_) Create_SemanticGraphDeltaV1_(SchemaVersion _dafny.Sequence, GraphId _dafny.Sequence, FromRevision _dafny.Sequence, ToRevision _dafny.Sequence, Reset bool, ResetReason m_ConfluxContract.Option, Replacement m_ConfluxContract.Option, Nodes GraphChanges, Edges GraphChanges, Facts GraphChanges, Findings GraphChanges, Evidence GraphChanges, CapabilitiesReplacement m_ConfluxContract.Option, LensesReplacement m_ConfluxContract.Option, ReceiptId _dafny.Sequence) SemanticGraphDeltaV1 {
	return SemanticGraphDeltaV1{SemanticGraphDeltaV1_SemanticGraphDeltaV1{SchemaVersion, GraphId, FromRevision, ToRevision, Reset, ResetReason, Replacement, Nodes, Edges, Facts, Findings, Evidence, CapabilitiesReplacement, LensesReplacement, ReceiptId}}
}

func (_this SemanticGraphDeltaV1) Is_SemanticGraphDeltaV1() bool {
	_, ok := _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1)
	return ok
}

func (CompanionStruct_SemanticGraphDeltaV1_) Default() SemanticGraphDeltaV1 {
	return Companion_SemanticGraphDeltaV1_.Create_SemanticGraphDeltaV1_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, false, m_ConfluxContract.Companion_Option_.Default(), m_ConfluxContract.Companion_Option_.Default(), Companion_GraphChanges_.Default(), Companion_GraphChanges_.Default(), Companion_GraphChanges_.Default(), Companion_GraphChanges_.Default(), Companion_GraphChanges_.Default(), m_ConfluxContract.Companion_Option_.Default(), m_ConfluxContract.Companion_Option_.Default(), _dafny.EmptySeq)
}

func (_this SemanticGraphDeltaV1) Dtor_schemaVersion() _dafny.Sequence {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).SchemaVersion
}

func (_this SemanticGraphDeltaV1) Dtor_graphId() _dafny.Sequence {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).GraphId
}

func (_this SemanticGraphDeltaV1) Dtor_fromRevision() _dafny.Sequence {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).FromRevision
}

func (_this SemanticGraphDeltaV1) Dtor_toRevision() _dafny.Sequence {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).ToRevision
}

func (_this SemanticGraphDeltaV1) Dtor_reset() bool {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).Reset
}

func (_this SemanticGraphDeltaV1) Dtor_resetReason() m_ConfluxContract.Option {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).ResetReason
}

func (_this SemanticGraphDeltaV1) Dtor_replacement() m_ConfluxContract.Option {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).Replacement
}

func (_this SemanticGraphDeltaV1) Dtor_nodes() GraphChanges {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).Nodes
}

func (_this SemanticGraphDeltaV1) Dtor_edges() GraphChanges {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).Edges
}

func (_this SemanticGraphDeltaV1) Dtor_facts() GraphChanges {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).Facts
}

func (_this SemanticGraphDeltaV1) Dtor_findings() GraphChanges {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).Findings
}

func (_this SemanticGraphDeltaV1) Dtor_evidence() GraphChanges {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).Evidence
}

func (_this SemanticGraphDeltaV1) Dtor_capabilitiesReplacement() m_ConfluxContract.Option {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).CapabilitiesReplacement
}

func (_this SemanticGraphDeltaV1) Dtor_lensesReplacement() m_ConfluxContract.Option {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).LensesReplacement
}

func (_this SemanticGraphDeltaV1) Dtor_receiptId() _dafny.Sequence {
	return _this.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1).ReceiptId
}

func (_this SemanticGraphDeltaV1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SemanticGraphDeltaV1_SemanticGraphDeltaV1:
		{
			return "ConfluxSemanticGraphContract.SemanticGraphDeltaV1.SemanticGraphDeltaV1" + "(" + data.SchemaVersion.VerbatimString(true) + ", " + data.GraphId.VerbatimString(true) + ", " + data.FromRevision.VerbatimString(true) + ", " + data.ToRevision.VerbatimString(true) + ", " + _dafny.String(data.Reset) + ", " + _dafny.String(data.ResetReason) + ", " + _dafny.String(data.Replacement) + ", " + _dafny.String(data.Nodes) + ", " + _dafny.String(data.Edges) + ", " + _dafny.String(data.Facts) + ", " + _dafny.String(data.Findings) + ", " + _dafny.String(data.Evidence) + ", " + _dafny.String(data.CapabilitiesReplacement) + ", " + _dafny.String(data.LensesReplacement) + ", " + data.ReceiptId.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SemanticGraphDeltaV1) Equals(other SemanticGraphDeltaV1) bool {
	switch data1 := _this.Get_().(type) {
	case SemanticGraphDeltaV1_SemanticGraphDeltaV1:
		{
			data2, ok := other.Get_().(SemanticGraphDeltaV1_SemanticGraphDeltaV1)
			return ok && data1.SchemaVersion.Equals(data2.SchemaVersion) && data1.GraphId.Equals(data2.GraphId) && data1.FromRevision.Equals(data2.FromRevision) && data1.ToRevision.Equals(data2.ToRevision) && data1.Reset == data2.Reset && data1.ResetReason.Equals(data2.ResetReason) && data1.Replacement.Equals(data2.Replacement) && data1.Nodes.Equals(data2.Nodes) && data1.Edges.Equals(data2.Edges) && data1.Facts.Equals(data2.Facts) && data1.Findings.Equals(data2.Findings) && data1.Evidence.Equals(data2.Evidence) && data1.CapabilitiesReplacement.Equals(data2.CapabilitiesReplacement) && data1.LensesReplacement.Equals(data2.LensesReplacement) && data1.ReceiptId.Equals(data2.ReceiptId)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SemanticGraphDeltaV1) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SemanticGraphDeltaV1)
	return ok && _this.Equals(typed)
}

func Type_SemanticGraphDeltaV1_() _dafny.TypeDescriptor {
	return type_SemanticGraphDeltaV1_{}
}

type type_SemanticGraphDeltaV1_ struct {
}

func (_this type_SemanticGraphDeltaV1_) Default() interface{} {
	return Companion_SemanticGraphDeltaV1_.Default()
}

func (_this type_SemanticGraphDeltaV1_) String() string {
	return "ConfluxSemanticGraphContract.SemanticGraphDeltaV1"
}
func (_this SemanticGraphDeltaV1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SemanticGraphDeltaV1{}

// End of datatype SemanticGraphDeltaV1
