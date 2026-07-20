// Package ConfluxContentAddress
// Dafny module ConfluxContentAddress compiled into Go

package ConfluxContentAddress

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
	m_ConfluxSemanticGraphIdentity "github.com/joshmouch/ahp-go/ConfluxSemanticGraphIdentity"
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
var _ m_ConfluxSemanticGraphIdentity.Dummy__

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
	return "ConfluxContentAddress.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) SameContent(left VerifiedContentRef, right VerifiedContentRef) bool {
	return ((left).Dtor_address()).Equals((right).Dtor_address())
}
func (_static *CompanionStruct_Default___) SizeMatches(address ContentAddress, bytes _dafny.Sequence) bool {
	var _source0 m_ConfluxContract.Option = (address).Dtor_exactSize()
	_ = _source0
	{
		if _source0.Is_None() {
			return true
		}
	}
	{
		var _0_n _dafny.Int = _source0.Get_().(m_ConfluxContract.Option_Some).Value.(_dafny.Int)
		_ = _0_n
		return (_dafny.IntOfUint32((bytes).Cardinality())).Cmp(_0_n) == 0
	}
}
func (_static *CompanionStruct_Default___) VerifyMaterialization(reference VerifiedContentRef, bytes _dafny.Sequence, digestBytes func(DigestAlgorithm, _dafny.Sequence) _dafny.Sequence) bool {
	return (Companion_Default___.SizeMatches((reference).Dtor_address(), bytes)) && (_dafny.Companion_Sequence_.Equal((digestBytes)(((reference).Dtor_address()).Dtor_digestAlgorithm(), bytes), ((reference).Dtor_address()).Dtor_digest()))
}

// End of class Default__

// Definition of datatype DigestAlgorithm
type DigestAlgorithm struct {
	Data_DigestAlgorithm_
}

func (_this DigestAlgorithm) Get_() Data_DigestAlgorithm_ {
	return _this.Data_DigestAlgorithm_
}

type Data_DigestAlgorithm_ interface {
	isDigestAlgorithm()
}

type CompanionStruct_DigestAlgorithm_ struct {
}

var Companion_DigestAlgorithm_ = CompanionStruct_DigestAlgorithm_{}

type DigestAlgorithm_Sha256 struct {
}

func (DigestAlgorithm_Sha256) isDigestAlgorithm() {}

func (CompanionStruct_DigestAlgorithm_) Create_Sha256_() DigestAlgorithm {
	return DigestAlgorithm{DigestAlgorithm_Sha256{}}
}

func (_this DigestAlgorithm) Is_Sha256() bool {
	_, ok := _this.Get_().(DigestAlgorithm_Sha256)
	return ok
}

func (CompanionStruct_DigestAlgorithm_) Default() DigestAlgorithm {
	return Companion_DigestAlgorithm_.Create_Sha256_()
}

func (_ CompanionStruct_DigestAlgorithm_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_DigestAlgorithm_.Create_Sha256_(), true
		default:
			return DigestAlgorithm{}, false
		}
	}
}

func (_this DigestAlgorithm) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case DigestAlgorithm_Sha256:
		{
			return "ConfluxContentAddress.DigestAlgorithm.Sha256"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this DigestAlgorithm) Equals(other DigestAlgorithm) bool {
	switch _this.Get_().(type) {
	case DigestAlgorithm_Sha256:
		{
			_, ok := other.Get_().(DigestAlgorithm_Sha256)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this DigestAlgorithm) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(DigestAlgorithm)
	return ok && _this.Equals(typed)
}

func Type_DigestAlgorithm_() _dafny.TypeDescriptor {
	return type_DigestAlgorithm_{}
}

type type_DigestAlgorithm_ struct {
}

func (_this type_DigestAlgorithm_) Default() interface{} {
	return Companion_DigestAlgorithm_.Default()
}

func (_this type_DigestAlgorithm_) String() string {
	return "ConfluxContentAddress.DigestAlgorithm"
}
func (_this DigestAlgorithm) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = DigestAlgorithm{}

// End of datatype DigestAlgorithm

// Definition of datatype ContentAddress
type ContentAddress struct {
	Data_ContentAddress_
}

func (_this ContentAddress) Get_() Data_ContentAddress_ {
	return _this.Data_ContentAddress_
}

type Data_ContentAddress_ interface {
	isContentAddress()
}

type CompanionStruct_ContentAddress_ struct {
}

var Companion_ContentAddress_ = CompanionStruct_ContentAddress_{}

type ContentAddress_ContentAddress struct {
	DigestAlgorithm DigestAlgorithm
	Digest          _dafny.Sequence
	ExactSize       m_ConfluxContract.Option
}

func (ContentAddress_ContentAddress) isContentAddress() {}

func (CompanionStruct_ContentAddress_) Create_ContentAddress_(DigestAlgorithm DigestAlgorithm, Digest _dafny.Sequence, ExactSize m_ConfluxContract.Option) ContentAddress {
	return ContentAddress{ContentAddress_ContentAddress{DigestAlgorithm, Digest, ExactSize}}
}

func (_this ContentAddress) Is_ContentAddress() bool {
	_, ok := _this.Get_().(ContentAddress_ContentAddress)
	return ok
}

func (CompanionStruct_ContentAddress_) Default() ContentAddress {
	return Companion_ContentAddress_.Create_ContentAddress_(Companion_DigestAlgorithm_.Default(), _dafny.EmptySeq, m_ConfluxContract.Companion_Option_.Default())
}

func (_this ContentAddress) Dtor_digestAlgorithm() DigestAlgorithm {
	return _this.Get_().(ContentAddress_ContentAddress).DigestAlgorithm
}

func (_this ContentAddress) Dtor_digest() _dafny.Sequence {
	return _this.Get_().(ContentAddress_ContentAddress).Digest
}

func (_this ContentAddress) Dtor_exactSize() m_ConfluxContract.Option {
	return _this.Get_().(ContentAddress_ContentAddress).ExactSize
}

func (_this ContentAddress) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ContentAddress_ContentAddress:
		{
			return "ConfluxContentAddress.ContentAddress.ContentAddress" + "(" + _dafny.String(data.DigestAlgorithm) + ", " + data.Digest.VerbatimString(true) + ", " + _dafny.String(data.ExactSize) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ContentAddress) Equals(other ContentAddress) bool {
	switch data1 := _this.Get_().(type) {
	case ContentAddress_ContentAddress:
		{
			data2, ok := other.Get_().(ContentAddress_ContentAddress)
			return ok && data1.DigestAlgorithm.Equals(data2.DigestAlgorithm) && data1.Digest.Equals(data2.Digest) && data1.ExactSize.Equals(data2.ExactSize)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ContentAddress) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ContentAddress)
	return ok && _this.Equals(typed)
}

func Type_ContentAddress_() _dafny.TypeDescriptor {
	return type_ContentAddress_{}
}

type type_ContentAddress_ struct {
}

func (_this type_ContentAddress_) Default() interface{} {
	return Companion_ContentAddress_.Default()
}

func (_this type_ContentAddress_) String() string {
	return "ConfluxContentAddress.ContentAddress"
}
func (_this ContentAddress) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ContentAddress{}

// End of datatype ContentAddress

// Definition of datatype ContentLocator
type ContentLocator struct {
	Data_ContentLocator_
}

func (_this ContentLocator) Get_() Data_ContentLocator_ {
	return _this.Data_ContentLocator_
}

type Data_ContentLocator_ interface {
	isContentLocator()
}

type CompanionStruct_ContentLocator_ struct {
}

var Companion_ContentLocator_ = CompanionStruct_ContentLocator_{}

type ContentLocator_ContentLocator struct {
	Uri         _dafny.Sequence
	ContentType m_ConfluxContract.Option
	Nonce       m_ConfluxContract.Option
}

func (ContentLocator_ContentLocator) isContentLocator() {}

func (CompanionStruct_ContentLocator_) Create_ContentLocator_(Uri _dafny.Sequence, ContentType m_ConfluxContract.Option, Nonce m_ConfluxContract.Option) ContentLocator {
	return ContentLocator{ContentLocator_ContentLocator{Uri, ContentType, Nonce}}
}

func (_this ContentLocator) Is_ContentLocator() bool {
	_, ok := _this.Get_().(ContentLocator_ContentLocator)
	return ok
}

func (CompanionStruct_ContentLocator_) Default() ContentLocator {
	return Companion_ContentLocator_.Create_ContentLocator_(_dafny.EmptySeq, m_ConfluxContract.Companion_Option_.Default(), m_ConfluxContract.Companion_Option_.Default())
}

func (_this ContentLocator) Dtor_uri() _dafny.Sequence {
	return _this.Get_().(ContentLocator_ContentLocator).Uri
}

func (_this ContentLocator) Dtor_contentType() m_ConfluxContract.Option {
	return _this.Get_().(ContentLocator_ContentLocator).ContentType
}

func (_this ContentLocator) Dtor_nonce() m_ConfluxContract.Option {
	return _this.Get_().(ContentLocator_ContentLocator).Nonce
}

func (_this ContentLocator) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ContentLocator_ContentLocator:
		{
			return "ConfluxContentAddress.ContentLocator.ContentLocator" + "(" + data.Uri.VerbatimString(true) + ", " + _dafny.String(data.ContentType) + ", " + _dafny.String(data.Nonce) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ContentLocator) Equals(other ContentLocator) bool {
	switch data1 := _this.Get_().(type) {
	case ContentLocator_ContentLocator:
		{
			data2, ok := other.Get_().(ContentLocator_ContentLocator)
			return ok && data1.Uri.Equals(data2.Uri) && data1.ContentType.Equals(data2.ContentType) && data1.Nonce.Equals(data2.Nonce)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ContentLocator) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ContentLocator)
	return ok && _this.Equals(typed)
}

func Type_ContentLocator_() _dafny.TypeDescriptor {
	return type_ContentLocator_{}
}

type type_ContentLocator_ struct {
}

func (_this type_ContentLocator_) Default() interface{} {
	return Companion_ContentLocator_.Default()
}

func (_this type_ContentLocator_) String() string {
	return "ConfluxContentAddress.ContentLocator"
}
func (_this ContentLocator) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ContentLocator{}

// End of datatype ContentLocator

// Definition of datatype VerifiedContentRef
type VerifiedContentRef struct {
	Data_VerifiedContentRef_
}

func (_this VerifiedContentRef) Get_() Data_VerifiedContentRef_ {
	return _this.Data_VerifiedContentRef_
}

type Data_VerifiedContentRef_ interface {
	isVerifiedContentRef()
}

type CompanionStruct_VerifiedContentRef_ struct {
}

var Companion_VerifiedContentRef_ = CompanionStruct_VerifiedContentRef_{}

type VerifiedContentRef_VerifiedContentRef struct {
	Address ContentAddress
	Locator m_ConfluxContract.Option
}

func (VerifiedContentRef_VerifiedContentRef) isVerifiedContentRef() {}

func (CompanionStruct_VerifiedContentRef_) Create_VerifiedContentRef_(Address ContentAddress, Locator m_ConfluxContract.Option) VerifiedContentRef {
	return VerifiedContentRef{VerifiedContentRef_VerifiedContentRef{Address, Locator}}
}

func (_this VerifiedContentRef) Is_VerifiedContentRef() bool {
	_, ok := _this.Get_().(VerifiedContentRef_VerifiedContentRef)
	return ok
}

func (CompanionStruct_VerifiedContentRef_) Default() VerifiedContentRef {
	return Companion_VerifiedContentRef_.Create_VerifiedContentRef_(Companion_ContentAddress_.Default(), m_ConfluxContract.Companion_Option_.Default())
}

func (_this VerifiedContentRef) Dtor_address() ContentAddress {
	return _this.Get_().(VerifiedContentRef_VerifiedContentRef).Address
}

func (_this VerifiedContentRef) Dtor_locator() m_ConfluxContract.Option {
	return _this.Get_().(VerifiedContentRef_VerifiedContentRef).Locator
}

func (_this VerifiedContentRef) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case VerifiedContentRef_VerifiedContentRef:
		{
			return "ConfluxContentAddress.VerifiedContentRef.VerifiedContentRef" + "(" + _dafny.String(data.Address) + ", " + _dafny.String(data.Locator) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this VerifiedContentRef) Equals(other VerifiedContentRef) bool {
	switch data1 := _this.Get_().(type) {
	case VerifiedContentRef_VerifiedContentRef:
		{
			data2, ok := other.Get_().(VerifiedContentRef_VerifiedContentRef)
			return ok && data1.Address.Equals(data2.Address) && data1.Locator.Equals(data2.Locator)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this VerifiedContentRef) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(VerifiedContentRef)
	return ok && _this.Equals(typed)
}

func Type_VerifiedContentRef_() _dafny.TypeDescriptor {
	return type_VerifiedContentRef_{}
}

type type_VerifiedContentRef_ struct {
}

func (_this type_VerifiedContentRef_) Default() interface{} {
	return Companion_VerifiedContentRef_.Default()
}

func (_this type_VerifiedContentRef_) String() string {
	return "ConfluxContentAddress.VerifiedContentRef"
}
func (_this VerifiedContentRef) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = VerifiedContentRef{}

// End of datatype VerifiedContentRef

// Definition of datatype ContentCodecs
type ContentCodecs struct {
	Data_ContentCodecs_
}

func (_this ContentCodecs) Get_() Data_ContentCodecs_ {
	return _this.Data_ContentCodecs_
}

type Data_ContentCodecs_ interface {
	isContentCodecs()
}

type CompanionStruct_ContentCodecs_ struct {
}

var Companion_ContentCodecs_ = CompanionStruct_ContentCodecs_{}

type ContentCodecs_ContentCodecs struct {
	EncodeAddress func(ContentAddress) interface{}
	DecodeAddress func(interface{}) m_ConfluxContract.Option
	EncodeLocator func(ContentLocator) interface{}
	DecodeLocator func(interface{}) m_ConfluxContract.Option
	EncodeRef     func(VerifiedContentRef) interface{}
	DecodeRef     func(interface{}) m_ConfluxContract.Option
}

func (ContentCodecs_ContentCodecs) isContentCodecs() {}

func (CompanionStruct_ContentCodecs_) Create_ContentCodecs_(EncodeAddress func(ContentAddress) interface{}, DecodeAddress func(interface{}) m_ConfluxContract.Option, EncodeLocator func(ContentLocator) interface{}, DecodeLocator func(interface{}) m_ConfluxContract.Option, EncodeRef func(VerifiedContentRef) interface{}, DecodeRef func(interface{}) m_ConfluxContract.Option) ContentCodecs {
	return ContentCodecs{ContentCodecs_ContentCodecs{EncodeAddress, DecodeAddress, EncodeLocator, DecodeLocator, EncodeRef, DecodeRef}}
}

func (_this ContentCodecs) Is_ContentCodecs() bool {
	_, ok := _this.Get_().(ContentCodecs_ContentCodecs)
	return ok
}

func (CompanionStruct_ContentCodecs_) Default(_default_W interface{}) ContentCodecs {
	return Companion_ContentCodecs_.Create_ContentCodecs_(func(ContentAddress) interface{} { return _default_W }, func(interface{}) m_ConfluxContract.Option { return m_ConfluxContract.Companion_Option_.Default() }, func(ContentLocator) interface{} { return _default_W }, func(interface{}) m_ConfluxContract.Option { return m_ConfluxContract.Companion_Option_.Default() }, func(VerifiedContentRef) interface{} { return _default_W }, func(interface{}) m_ConfluxContract.Option { return m_ConfluxContract.Companion_Option_.Default() })
}

func (_this ContentCodecs) Dtor_encodeAddress() func(ContentAddress) interface{} {
	return _this.Get_().(ContentCodecs_ContentCodecs).EncodeAddress
}

func (_this ContentCodecs) Dtor_decodeAddress() func(interface{}) m_ConfluxContract.Option {
	return _this.Get_().(ContentCodecs_ContentCodecs).DecodeAddress
}

func (_this ContentCodecs) Dtor_encodeLocator() func(ContentLocator) interface{} {
	return _this.Get_().(ContentCodecs_ContentCodecs).EncodeLocator
}

func (_this ContentCodecs) Dtor_decodeLocator() func(interface{}) m_ConfluxContract.Option {
	return _this.Get_().(ContentCodecs_ContentCodecs).DecodeLocator
}

func (_this ContentCodecs) Dtor_encodeRef() func(VerifiedContentRef) interface{} {
	return _this.Get_().(ContentCodecs_ContentCodecs).EncodeRef
}

func (_this ContentCodecs) Dtor_decodeRef() func(interface{}) m_ConfluxContract.Option {
	return _this.Get_().(ContentCodecs_ContentCodecs).DecodeRef
}

func (_this ContentCodecs) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ContentCodecs_ContentCodecs:
		{
			return "ConfluxContentAddress.ContentCodecs.ContentCodecs" + "(" + _dafny.String(data.EncodeAddress) + ", " + _dafny.String(data.DecodeAddress) + ", " + _dafny.String(data.EncodeLocator) + ", " + _dafny.String(data.DecodeLocator) + ", " + _dafny.String(data.EncodeRef) + ", " + _dafny.String(data.DecodeRef) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ContentCodecs) Equals(other ContentCodecs) bool {
	switch data1 := _this.Get_().(type) {
	case ContentCodecs_ContentCodecs:
		{
			data2, ok := other.Get_().(ContentCodecs_ContentCodecs)
			return ok && _dafny.AreEqual(data1.EncodeAddress, data2.EncodeAddress) && _dafny.AreEqual(data1.DecodeAddress, data2.DecodeAddress) && _dafny.AreEqual(data1.EncodeLocator, data2.EncodeLocator) && _dafny.AreEqual(data1.DecodeLocator, data2.DecodeLocator) && _dafny.AreEqual(data1.EncodeRef, data2.EncodeRef) && _dafny.AreEqual(data1.DecodeRef, data2.DecodeRef)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ContentCodecs) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ContentCodecs)
	return ok && _this.Equals(typed)
}

func Type_ContentCodecs_(Type_W_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_ContentCodecs_{Type_W_}
}

type type_ContentCodecs_ struct {
	Type_W_ _dafny.TypeDescriptor
}

func (_this type_ContentCodecs_) Default() interface{} {
	Type_W_ := _this.Type_W_
	_ = Type_W_
	return Companion_ContentCodecs_.Default(Type_W_.Default())
}

func (_this type_ContentCodecs_) String() string {
	return "ConfluxContentAddress.ContentCodecs"
}
func (_this ContentCodecs) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ContentCodecs{}

// End of datatype ContentCodecs
