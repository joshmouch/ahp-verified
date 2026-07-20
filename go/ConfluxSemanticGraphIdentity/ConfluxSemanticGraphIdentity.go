// Package ConfluxSemanticGraphIdentity
// Dafny module ConfluxSemanticGraphIdentity compiled into Go

package ConfluxSemanticGraphIdentity

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-go/Changeset"
	m_Chat "github.com/joshmouch/ahp-go/Chat"
	m_ClientMain "github.com/joshmouch/ahp-go/ClientMain"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-go/ConfluxOrderedLog"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-go/ConfluxSeqRoute"
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
	return "ConfluxSemanticGraphIdentity.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) StableId(canonicalEncode func(StableIdentityInput) _dafny.Sequence, sha256 func(_dafny.Sequence) _dafny.Sequence, input StableIdentityInput) _dafny.Sequence {
	return (sha256)(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("archai:stable-id:1.0.0\n"), (canonicalEncode)(input)))
}
func (_static *CompanionStruct_Default___) ContentRevision(kind _dafny.Sequence, schemaVersion _dafny.Sequence, canonicalSemanticContent func(interface{}) _dafny.Sequence, sha256 func(_dafny.Sequence) _dafny.Sequence, content interface{}) _dafny.Sequence {
	return (sha256)(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("archai:"), kind), _dafny.UnicodeSeqOfUtf8Bytes(":")), schemaVersion), _dafny.UnicodeSeqOfUtf8Bytes("\n")), (canonicalSemanticContent)(content)))
}
func (_static *CompanionStruct_Default___) GraphRevision(canonicalEncode func(GraphIdentityInput) _dafny.Sequence, sha256 func(_dafny.Sequence) _dafny.Sequence, input GraphIdentityInput) _dafny.Sequence {
	return (sha256)(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("archai:semantic-graph:1.0.0\n"), (canonicalEncode)(input)))
}
func (_static *CompanionStruct_Default___) GraphId(sha256 func(_dafny.Sequence) _dafny.Sequence, authorityId _dafny.Sequence, selectorKind _dafny.Sequence, selectorId _dafny.Sequence, schemaEpoch _dafny.Sequence) _dafny.Sequence {
	return (sha256)(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("archai:graph-id:1.0.0\n"), authorityId), _dafny.UnicodeSeqOfUtf8Bytes("\n")), selectorKind), _dafny.UnicodeSeqOfUtf8Bytes("\n")), selectorId), _dafny.UnicodeSeqOfUtf8Bytes("\n")), schemaEpoch))
}

// End of class Default__

// Definition of datatype StableIdentityInput
type StableIdentityInput struct {
	Data_StableIdentityInput_
}

func (_this StableIdentityInput) Get_() Data_StableIdentityInput_ {
	return _this.Data_StableIdentityInput_
}

type Data_StableIdentityInput_ interface {
	isStableIdentityInput()
}

type CompanionStruct_StableIdentityInput_ struct {
}

var Companion_StableIdentityInput_ = CompanionStruct_StableIdentityInput_{}

type StableIdentityInput_StableIdentityInput struct {
	Kind        _dafny.Sequence
	AuthorityId _dafny.Sequence
	NaturalKey  _dafny.Sequence
}

func (StableIdentityInput_StableIdentityInput) isStableIdentityInput() {}

func (CompanionStruct_StableIdentityInput_) Create_StableIdentityInput_(Kind _dafny.Sequence, AuthorityId _dafny.Sequence, NaturalKey _dafny.Sequence) StableIdentityInput {
	return StableIdentityInput{StableIdentityInput_StableIdentityInput{Kind, AuthorityId, NaturalKey}}
}

func (_this StableIdentityInput) Is_StableIdentityInput() bool {
	_, ok := _this.Get_().(StableIdentityInput_StableIdentityInput)
	return ok
}

func (CompanionStruct_StableIdentityInput_) Default() StableIdentityInput {
	return Companion_StableIdentityInput_.Create_StableIdentityInput_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this StableIdentityInput) Dtor_kind() _dafny.Sequence {
	return _this.Get_().(StableIdentityInput_StableIdentityInput).Kind
}

func (_this StableIdentityInput) Dtor_authorityId() _dafny.Sequence {
	return _this.Get_().(StableIdentityInput_StableIdentityInput).AuthorityId
}

func (_this StableIdentityInput) Dtor_naturalKey() _dafny.Sequence {
	return _this.Get_().(StableIdentityInput_StableIdentityInput).NaturalKey
}

func (_this StableIdentityInput) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case StableIdentityInput_StableIdentityInput:
		{
			return "ConfluxSemanticGraphIdentity.StableIdentityInput.StableIdentityInput" + "(" + data.Kind.VerbatimString(true) + ", " + data.AuthorityId.VerbatimString(true) + ", " + data.NaturalKey.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this StableIdentityInput) Equals(other StableIdentityInput) bool {
	switch data1 := _this.Get_().(type) {
	case StableIdentityInput_StableIdentityInput:
		{
			data2, ok := other.Get_().(StableIdentityInput_StableIdentityInput)
			return ok && data1.Kind.Equals(data2.Kind) && data1.AuthorityId.Equals(data2.AuthorityId) && data1.NaturalKey.Equals(data2.NaturalKey)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this StableIdentityInput) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(StableIdentityInput)
	return ok && _this.Equals(typed)
}

func Type_StableIdentityInput_() _dafny.TypeDescriptor {
	return type_StableIdentityInput_{}
}

type type_StableIdentityInput_ struct {
}

func (_this type_StableIdentityInput_) Default() interface{} {
	return Companion_StableIdentityInput_.Default()
}

func (_this type_StableIdentityInput_) String() string {
	return "ConfluxSemanticGraphIdentity.StableIdentityInput"
}
func (_this StableIdentityInput) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = StableIdentityInput{}

// End of datatype StableIdentityInput

// Definition of datatype GraphIdentityInput
type GraphIdentityInput struct {
	Data_GraphIdentityInput_
}

func (_this GraphIdentityInput) Get_() Data_GraphIdentityInput_ {
	return _this.Data_GraphIdentityInput_
}

type Data_GraphIdentityInput_ interface {
	isGraphIdentityInput()
}

type CompanionStruct_GraphIdentityInput_ struct {
}

var Companion_GraphIdentityInput_ = CompanionStruct_GraphIdentityInput_{}

type GraphIdentityInput_GraphIdentityInput struct {
	GraphId                   _dafny.Sequence
	AuthorityArtifactIdentity _dafny.Sequence
	CapabilityRevisions       _dafny.Sequence
	LensRevisions             _dafny.Sequence
	RecordRevisions           _dafny.Sequence
}

func (GraphIdentityInput_GraphIdentityInput) isGraphIdentityInput() {}

func (CompanionStruct_GraphIdentityInput_) Create_GraphIdentityInput_(GraphId _dafny.Sequence, AuthorityArtifactIdentity _dafny.Sequence, CapabilityRevisions _dafny.Sequence, LensRevisions _dafny.Sequence, RecordRevisions _dafny.Sequence) GraphIdentityInput {
	return GraphIdentityInput{GraphIdentityInput_GraphIdentityInput{GraphId, AuthorityArtifactIdentity, CapabilityRevisions, LensRevisions, RecordRevisions}}
}

func (_this GraphIdentityInput) Is_GraphIdentityInput() bool {
	_, ok := _this.Get_().(GraphIdentityInput_GraphIdentityInput)
	return ok
}

func (CompanionStruct_GraphIdentityInput_) Default() GraphIdentityInput {
	return Companion_GraphIdentityInput_.Create_GraphIdentityInput_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this GraphIdentityInput) Dtor_graphId() _dafny.Sequence {
	return _this.Get_().(GraphIdentityInput_GraphIdentityInput).GraphId
}

func (_this GraphIdentityInput) Dtor_authorityArtifactIdentity() _dafny.Sequence {
	return _this.Get_().(GraphIdentityInput_GraphIdentityInput).AuthorityArtifactIdentity
}

func (_this GraphIdentityInput) Dtor_capabilityRevisions() _dafny.Sequence {
	return _this.Get_().(GraphIdentityInput_GraphIdentityInput).CapabilityRevisions
}

func (_this GraphIdentityInput) Dtor_lensRevisions() _dafny.Sequence {
	return _this.Get_().(GraphIdentityInput_GraphIdentityInput).LensRevisions
}

func (_this GraphIdentityInput) Dtor_recordRevisions() _dafny.Sequence {
	return _this.Get_().(GraphIdentityInput_GraphIdentityInput).RecordRevisions
}

func (_this GraphIdentityInput) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case GraphIdentityInput_GraphIdentityInput:
		{
			return "ConfluxSemanticGraphIdentity.GraphIdentityInput.GraphIdentityInput" + "(" + data.GraphId.VerbatimString(true) + ", " + data.AuthorityArtifactIdentity.VerbatimString(true) + ", " + _dafny.String(data.CapabilityRevisions) + ", " + _dafny.String(data.LensRevisions) + ", " + _dafny.String(data.RecordRevisions) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GraphIdentityInput) Equals(other GraphIdentityInput) bool {
	switch data1 := _this.Get_().(type) {
	case GraphIdentityInput_GraphIdentityInput:
		{
			data2, ok := other.Get_().(GraphIdentityInput_GraphIdentityInput)
			return ok && data1.GraphId.Equals(data2.GraphId) && data1.AuthorityArtifactIdentity.Equals(data2.AuthorityArtifactIdentity) && data1.CapabilityRevisions.Equals(data2.CapabilityRevisions) && data1.LensRevisions.Equals(data2.LensRevisions) && data1.RecordRevisions.Equals(data2.RecordRevisions)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GraphIdentityInput) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GraphIdentityInput)
	return ok && _this.Equals(typed)
}

func Type_GraphIdentityInput_() _dafny.TypeDescriptor {
	return type_GraphIdentityInput_{}
}

type type_GraphIdentityInput_ struct {
}

func (_this type_GraphIdentityInput_) Default() interface{} {
	return Companion_GraphIdentityInput_.Default()
}

func (_this type_GraphIdentityInput_) String() string {
	return "ConfluxSemanticGraphIdentity.GraphIdentityInput"
}
func (_this GraphIdentityInput) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = GraphIdentityInput{}

// End of datatype GraphIdentityInput

// Definition of datatype NonSemanticMetadata
type NonSemanticMetadata struct {
	Data_NonSemanticMetadata_
}

func (_this NonSemanticMetadata) Get_() Data_NonSemanticMetadata_ {
	return _this.Data_NonSemanticMetadata_
}

type Data_NonSemanticMetadata_ interface {
	isNonSemanticMetadata()
}

type CompanionStruct_NonSemanticMetadata_ struct {
}

var Companion_NonSemanticMetadata_ = CompanionStruct_NonSemanticMetadata_{}

type NonSemanticMetadata_NonSemanticMetadata struct {
	AbsoluteLocator _dafny.Sequence
	GitHead         _dafny.Sequence
	ListenerPid     _dafny.Sequence
	WallClock       _dafny.Sequence
	ReceiptLocator  _dafny.Sequence
}

func (NonSemanticMetadata_NonSemanticMetadata) isNonSemanticMetadata() {}

func (CompanionStruct_NonSemanticMetadata_) Create_NonSemanticMetadata_(AbsoluteLocator _dafny.Sequence, GitHead _dafny.Sequence, ListenerPid _dafny.Sequence, WallClock _dafny.Sequence, ReceiptLocator _dafny.Sequence) NonSemanticMetadata {
	return NonSemanticMetadata{NonSemanticMetadata_NonSemanticMetadata{AbsoluteLocator, GitHead, ListenerPid, WallClock, ReceiptLocator}}
}

func (_this NonSemanticMetadata) Is_NonSemanticMetadata() bool {
	_, ok := _this.Get_().(NonSemanticMetadata_NonSemanticMetadata)
	return ok
}

func (CompanionStruct_NonSemanticMetadata_) Default() NonSemanticMetadata {
	return Companion_NonSemanticMetadata_.Create_NonSemanticMetadata_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this NonSemanticMetadata) Dtor_absoluteLocator() _dafny.Sequence {
	return _this.Get_().(NonSemanticMetadata_NonSemanticMetadata).AbsoluteLocator
}

func (_this NonSemanticMetadata) Dtor_gitHead() _dafny.Sequence {
	return _this.Get_().(NonSemanticMetadata_NonSemanticMetadata).GitHead
}

func (_this NonSemanticMetadata) Dtor_listenerPid() _dafny.Sequence {
	return _this.Get_().(NonSemanticMetadata_NonSemanticMetadata).ListenerPid
}

func (_this NonSemanticMetadata) Dtor_wallClock() _dafny.Sequence {
	return _this.Get_().(NonSemanticMetadata_NonSemanticMetadata).WallClock
}

func (_this NonSemanticMetadata) Dtor_receiptLocator() _dafny.Sequence {
	return _this.Get_().(NonSemanticMetadata_NonSemanticMetadata).ReceiptLocator
}

func (_this NonSemanticMetadata) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case NonSemanticMetadata_NonSemanticMetadata:
		{
			return "ConfluxSemanticGraphIdentity.NonSemanticMetadata.NonSemanticMetadata" + "(" + data.AbsoluteLocator.VerbatimString(true) + ", " + data.GitHead.VerbatimString(true) + ", " + data.ListenerPid.VerbatimString(true) + ", " + data.WallClock.VerbatimString(true) + ", " + data.ReceiptLocator.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this NonSemanticMetadata) Equals(other NonSemanticMetadata) bool {
	switch data1 := _this.Get_().(type) {
	case NonSemanticMetadata_NonSemanticMetadata:
		{
			data2, ok := other.Get_().(NonSemanticMetadata_NonSemanticMetadata)
			return ok && data1.AbsoluteLocator.Equals(data2.AbsoluteLocator) && data1.GitHead.Equals(data2.GitHead) && data1.ListenerPid.Equals(data2.ListenerPid) && data1.WallClock.Equals(data2.WallClock) && data1.ReceiptLocator.Equals(data2.ReceiptLocator)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this NonSemanticMetadata) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(NonSemanticMetadata)
	return ok && _this.Equals(typed)
}

func Type_NonSemanticMetadata_() _dafny.TypeDescriptor {
	return type_NonSemanticMetadata_{}
}

type type_NonSemanticMetadata_ struct {
}

func (_this type_NonSemanticMetadata_) Default() interface{} {
	return Companion_NonSemanticMetadata_.Default()
}

func (_this type_NonSemanticMetadata_) String() string {
	return "ConfluxSemanticGraphIdentity.NonSemanticMetadata"
}
func (_this NonSemanticMetadata) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = NonSemanticMetadata{}

// End of datatype NonSemanticMetadata
