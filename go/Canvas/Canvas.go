// Package Canvas
// Dafny module Canvas compiled into Go

package Canvas

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ResourceWatch "github.com/joshmouch/ahp-verified/go/ResourceWatch"
	m__System "github.com/joshmouch/ahp-verified/go/System_"
	_dafny "github.com/joshmouch/ahp-verified/go/dafny"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_ConfluxCodec.Dummy__
var _ m_ConfluxContract.Dummy__
var _ m_AhpSkeleton.Dummy__
var _ m_ResourceWatch.Dummy__

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
	return "Canvas.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ApplyToCanvas(s CanvasState, a CanvasAction, now _dafny.Int) _dafny.Tuple {
	var _pat_let_tv0 = s
	_ = _pat_let_tv0
	var _pat_let_tv1 = s
	_ = _pat_let_tv1
	var _pat_let_tv2 = s
	_ = _pat_let_tv2
	var _pat_let_tv3 = s
	_ = _pat_let_tv3
	var _source0 CanvasAction = a
	_ = _source0
	{
		if _source0.Is_Updated() {
			var _0_t m_AhpSkeleton.Option = _source0.Get_().(CanvasAction_Updated).Title
			_ = _0_t
			var _1_ac m_AhpSkeleton.Option = _source0.Get_().(CanvasAction_Updated).Activity
			_ = _1_ac
			var _2_uri m_AhpSkeleton.Option = _source0.Get_().(CanvasAction_Updated).ContentUri
			_ = _2_uri
			var _3_av m_AhpSkeleton.Option = _source0.Get_().(CanvasAction_Updated).Availability
			_ = _3_av
			return _dafny.TupleOf(func(_pat_let8_0 CanvasState) CanvasState {
				return func(_4_dt__update__tmp_h0 CanvasState) CanvasState {
					return func(_pat_let9_0 CanvasAvailability) CanvasState {
						return func(_5_dt__update_havailability_h0 CanvasAvailability) CanvasState {
							return func(_pat_let10_0 m_AhpSkeleton.Option) CanvasState {
								return func(_6_dt__update_hcontentUri_h0 m_AhpSkeleton.Option) CanvasState {
									return func(_pat_let11_0 m_AhpSkeleton.Option) CanvasState {
										return func(_7_dt__update_hactivity_h0 m_AhpSkeleton.Option) CanvasState {
											return func(_pat_let12_0 m_AhpSkeleton.Option) CanvasState {
												return func(_8_dt__update_htitle_h0 m_AhpSkeleton.Option) CanvasState {
													return Companion_CanvasState_.Create_CanvasState_((_4_dt__update__tmp_h0).Dtor_canvasId(), (_4_dt__update__tmp_h0).Dtor_providerId(), (_4_dt__update__tmp_h0).Dtor_input(), _8_dt__update_htitle_h0, _7_dt__update_hactivity_h0, _6_dt__update_hcontentUri_h0, _5_dt__update_havailability_h0)
												}(_pat_let12_0)
											}((func() m_AhpSkeleton.Option {
												if (_0_t).Is_Some() {
													return _0_t
												}
												return (_pat_let_tv3).Dtor_title()
											})())
										}(_pat_let11_0)
									}((func() m_AhpSkeleton.Option {
										if (_1_ac).Is_Some() {
											return _1_ac
										}
										return (_pat_let_tv2).Dtor_activity()
									})())
								}(_pat_let10_0)
							}((func() m_AhpSkeleton.Option {
								if (_2_uri).Is_Some() {
									return _2_uri
								}
								return (_pat_let_tv1).Dtor_contentUri()
							})())
						}(_pat_let9_0)
					}((func() CanvasAvailability {
						if (_3_av).Is_Some() {
							return (_3_av).Dtor_value().(CanvasAvailability)
						}
						return (_pat_let_tv0).Dtor_availability()
					})())
				}(_pat_let8_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_CloseRequested() {
			return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
		}
	}
	{
		return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
	}
}
func (_static *CompanionStruct_Default___) Apply1(s CanvasState, a CanvasAction) CanvasState {
	return (*(Companion_Default___.ApplyToCanvas(s, a, _dafny.IntOfInt64(9999))).IndexInt(0)).(CanvasState)
}
func (_static *CompanionStruct_Default___) Fold(s CanvasState, actions _dafny.Sequence) CanvasState {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer4 func(CanvasState, CanvasAction) CanvasState) func(interface{}, interface{}) interface{} {
		return func(arg6 interface{}, arg7 interface{}) interface{} {
			return coer4(arg6.(CanvasState), arg7.(CanvasAction))
		}
	}(Companion_Default___.Apply1), s, actions).(CanvasState)
}
func (_static *CompanionStruct_Default___) C0() CanvasState {
	return Companion_CanvasState_.Create_CanvasState_(_dafny.UnicodeSeqOfUtf8Bytes("markdown"), _dafny.UnicodeSeqOfUtf8Bytes("opaque-provider"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), Companion_CanvasAvailability_.Create_Ready_())
}
func (_static *CompanionStruct_Default___) SameInstance(a CanvasInstance, b CanvasInstance) bool {
	return _dafny.Companion_Sequence_.Equal((a).Dtor_channel(), (b).Dtor_channel())
}
func (_static *CompanionStruct_Default___) ContentResourceRead(contentUri _dafny.Sequence) ResourceReadParams {
	return Companion_ResourceReadParams_.Create_ResourceReadParams_(_dafny.UnicodeSeqOfUtf8Bytes("ahp-root://"), contentUri)
}
func (_static *CompanionStruct_Default___) ValidateCanvasAdmission(supportNegotiated bool, handlerInstalled bool, authorized bool, channelInUse bool) CanvasAdmission {
	if !(supportNegotiated) {
		return Companion_CanvasAdmission_.Create_CanvasUnsupported_()
	} else if !(handlerInstalled) {
		return Companion_CanvasAdmission_.Create_CanvasHandlerUnavailable_()
	} else if !(authorized) {
		return Companion_CanvasAdmission_.Create_CanvasUnauthorized_()
	} else if channelInUse {
		return Companion_CanvasAdmission_.Create_CanvasUriCollision_()
	} else {
		return Companion_CanvasAdmission_.Create_CanvasSupported_()
	}
}
func (_static *CompanionStruct_Default___) CanvasProviderFailure(code _dafny.Sequence, message _dafny.Sequence) CanvasError {
	return Companion_CanvasError_.Create_ProviderError_(Companion_Default___.CANVAS__PROVIDER__ERROR(), Companion_CanvasProviderErrorData_.Create_CanvasProviderErrorData_(code, message))
}
func (_static *CompanionStruct_Default___) DegradeAvailability(s CanvasState) CanvasState {
	var _0_dt__update__tmp_h0 CanvasState = s
	_ = _0_dt__update__tmp_h0
	var _1_dt__update_havailability_h0 CanvasAvailability = Companion_CanvasAvailability_.Create_Stale_()
	_ = _1_dt__update_havailability_h0
	return Companion_CanvasState_.Create_CanvasState_((_0_dt__update__tmp_h0).Dtor_canvasId(), (_0_dt__update__tmp_h0).Dtor_providerId(), (_0_dt__update__tmp_h0).Dtor_input(), (_0_dt__update__tmp_h0).Dtor_title(), (_0_dt__update__tmp_h0).Dtor_activity(), (_0_dt__update__tmp_h0).Dtor_contentUri(), _1_dt__update_havailability_h0)
}
func (_static *CompanionStruct_Default___) ReconnectCanvas(snapshot CanvasState, supportNegotiated bool) m_AhpSkeleton.Option {
	if supportNegotiated {
		return m_AhpSkeleton.Companion_Option_.Create_Some_(snapshot)
	} else {
		return m_AhpSkeleton.Companion_Option_.Create_None_()
	}
}
func (_static *CompanionStruct_Default___) RunCorpus() (_dafny.Int, _dafny.Int) {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	var total _dafny.Int = _dafny.Zero
	_ = total
	total = _dafny.IntOfInt64(5)
	pass = _dafny.Zero
	var _0_s CanvasState
	_ = _0_s
	_0_s = Companion_CanvasState_.Create_CanvasState_(_dafny.UnicodeSeqOfUtf8Bytes("markdown"), _dafny.UnicodeSeqOfUtf8Bytes("acme.editors"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Draft")), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("idle")), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("ahp-session:/2f9c/content/canvas-1")), Companion_CanvasAvailability_.Create_Ready_())
	var _1_all CanvasState
	_ = _1_all
	_1_all = Companion_Default___.Apply1(_0_s, Companion_CanvasAction_.Create_Updated_(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Published")), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("saved")), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("https://example.com/docs/published.html")), m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_CanvasAvailability_.Create_Stale_())))
	if (_1_all).Equals(func(_pat_let13_0 CanvasState) CanvasState {
		return func(_2_dt__update__tmp_h0 CanvasState) CanvasState {
			return func(_pat_let14_0 CanvasAvailability) CanvasState {
				return func(_3_dt__update_havailability_h0 CanvasAvailability) CanvasState {
					return func(_pat_let15_0 m_AhpSkeleton.Option) CanvasState {
						return func(_4_dt__update_hcontentUri_h0 m_AhpSkeleton.Option) CanvasState {
							return func(_pat_let16_0 m_AhpSkeleton.Option) CanvasState {
								return func(_5_dt__update_hactivity_h0 m_AhpSkeleton.Option) CanvasState {
									return func(_pat_let17_0 m_AhpSkeleton.Option) CanvasState {
										return func(_6_dt__update_htitle_h0 m_AhpSkeleton.Option) CanvasState {
											return Companion_CanvasState_.Create_CanvasState_((_2_dt__update__tmp_h0).Dtor_canvasId(), (_2_dt__update__tmp_h0).Dtor_providerId(), (_2_dt__update__tmp_h0).Dtor_input(), _6_dt__update_htitle_h0, _5_dt__update_hactivity_h0, _4_dt__update_hcontentUri_h0, _3_dt__update_havailability_h0)
										}(_pat_let17_0)
									}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Published")))
								}(_pat_let16_0)
							}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("saved")))
						}(_pat_let15_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("https://example.com/docs/published.html")))
				}(_pat_let14_0)
			}(Companion_CanvasAvailability_.Create_Stale_())
		}(_pat_let13_0)
	}(_0_s)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _7_left CanvasState
	_ = _7_left
	_7_left = Companion_Default___.Apply1(_0_s, Companion_CanvasAction_.Create_Updated_(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Renamed")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("https://example.com/docs/renamed.html")), m_AhpSkeleton.Companion_Option_.Create_None_()))
	if (_7_left).Equals(func(_pat_let18_0 CanvasState) CanvasState {
		return func(_8_dt__update__tmp_h1 CanvasState) CanvasState {
			return func(_pat_let19_0 m_AhpSkeleton.Option) CanvasState {
				return func(_9_dt__update_hcontentUri_h1 m_AhpSkeleton.Option) CanvasState {
					return func(_pat_let20_0 m_AhpSkeleton.Option) CanvasState {
						return func(_10_dt__update_htitle_h1 m_AhpSkeleton.Option) CanvasState {
							return Companion_CanvasState_.Create_CanvasState_((_8_dt__update__tmp_h1).Dtor_canvasId(), (_8_dt__update__tmp_h1).Dtor_providerId(), (_8_dt__update__tmp_h1).Dtor_input(), _10_dt__update_htitle_h1, (_8_dt__update__tmp_h1).Dtor_activity(), _9_dt__update_hcontentUri_h1, (_8_dt__update__tmp_h1).Dtor_availability())
						}(_pat_let20_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Renamed")))
				}(_pat_let19_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("https://example.com/docs/renamed.html")))
		}(_pat_let18_0)
	}(_0_s)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _11_right CanvasState
	_ = _11_right
	_11_right = Companion_Default___.Apply1(_0_s, Companion_CanvasAction_.Create_Updated_(m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("error")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_CanvasAvailability_.Create_Stale_())))
	if (_11_right).Equals(func(_pat_let21_0 CanvasState) CanvasState {
		return func(_12_dt__update__tmp_h2 CanvasState) CanvasState {
			return func(_pat_let22_0 CanvasAvailability) CanvasState {
				return func(_13_dt__update_havailability_h1 CanvasAvailability) CanvasState {
					return func(_pat_let23_0 m_AhpSkeleton.Option) CanvasState {
						return func(_14_dt__update_hactivity_h1 m_AhpSkeleton.Option) CanvasState {
							return Companion_CanvasState_.Create_CanvasState_((_12_dt__update__tmp_h2).Dtor_canvasId(), (_12_dt__update__tmp_h2).Dtor_providerId(), (_12_dt__update__tmp_h2).Dtor_input(), (_12_dt__update__tmp_h2).Dtor_title(), _14_dt__update_hactivity_h1, (_12_dt__update__tmp_h2).Dtor_contentUri(), _13_dt__update_havailability_h1)
						}(_pat_let23_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("error")))
				}(_pat_let22_0)
			}(Companion_CanvasAvailability_.Create_Stale_())
		}(_pat_let21_0)
	}(_0_s)) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_0_s, Companion_CanvasAction_.Create_CloseRequested_())).Equals(_0_s) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_0_s, Companion_CanvasAction_.Create_CanvasUnknown_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("type"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("canvas/nonExistentAction"))))))).Equals(_0_s) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass, total
}
func (_static *CompanionStruct_Default___) CANVAS__PROVIDER__ERROR() _dafny.Int {
	return _dafny.IntOfInt64(-32012)
}

// End of class Default__

// Definition of datatype CanvasAvailability
type CanvasAvailability struct {
	Data_CanvasAvailability_
}

func (_this CanvasAvailability) Get_() Data_CanvasAvailability_ {
	return _this.Data_CanvasAvailability_
}

type Data_CanvasAvailability_ interface {
	isCanvasAvailability()
}

type CompanionStruct_CanvasAvailability_ struct {
}

var Companion_CanvasAvailability_ = CompanionStruct_CanvasAvailability_{}

type CanvasAvailability_Ready struct {
}

func (CanvasAvailability_Ready) isCanvasAvailability() {}

func (CompanionStruct_CanvasAvailability_) Create_Ready_() CanvasAvailability {
	return CanvasAvailability{CanvasAvailability_Ready{}}
}

func (_this CanvasAvailability) Is_Ready() bool {
	_, ok := _this.Get_().(CanvasAvailability_Ready)
	return ok
}

type CanvasAvailability_Stale struct {
}

func (CanvasAvailability_Stale) isCanvasAvailability() {}

func (CompanionStruct_CanvasAvailability_) Create_Stale_() CanvasAvailability {
	return CanvasAvailability{CanvasAvailability_Stale{}}
}

func (_this CanvasAvailability) Is_Stale() bool {
	_, ok := _this.Get_().(CanvasAvailability_Stale)
	return ok
}

func (CompanionStruct_CanvasAvailability_) Default() CanvasAvailability {
	return Companion_CanvasAvailability_.Create_Ready_()
}

func (_ CompanionStruct_CanvasAvailability_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_CanvasAvailability_.Create_Ready_(), true
		case 1:
			return Companion_CanvasAvailability_.Create_Stale_(), true
		default:
			return CanvasAvailability{}, false
		}
	}
}

func (_this CanvasAvailability) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case CanvasAvailability_Ready:
		{
			return "Canvas.CanvasAvailability.Ready"
		}
	case CanvasAvailability_Stale:
		{
			return "Canvas.CanvasAvailability.Stale"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CanvasAvailability) Equals(other CanvasAvailability) bool {
	switch _this.Get_().(type) {
	case CanvasAvailability_Ready:
		{
			_, ok := other.Get_().(CanvasAvailability_Ready)
			return ok
		}
	case CanvasAvailability_Stale:
		{
			_, ok := other.Get_().(CanvasAvailability_Stale)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CanvasAvailability) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CanvasAvailability)
	return ok && _this.Equals(typed)
}

func Type_CanvasAvailability_() _dafny.TypeDescriptor {
	return type_CanvasAvailability_{}
}

type type_CanvasAvailability_ struct {
}

func (_this type_CanvasAvailability_) Default() interface{} {
	return Companion_CanvasAvailability_.Default()
}

func (_this type_CanvasAvailability_) String() string {
	return "Canvas.CanvasAvailability"
}
func (_this CanvasAvailability) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CanvasAvailability{}

// End of datatype CanvasAvailability

// Definition of datatype CanvasState
type CanvasState struct {
	Data_CanvasState_
}

func (_this CanvasState) Get_() Data_CanvasState_ {
	return _this.Data_CanvasState_
}

type Data_CanvasState_ interface {
	isCanvasState()
}

type CompanionStruct_CanvasState_ struct {
}

var Companion_CanvasState_ = CompanionStruct_CanvasState_{}

type CanvasState_CanvasState struct {
	CanvasId     _dafny.Sequence
	ProviderId   _dafny.Sequence
	Input        m_AhpSkeleton.Option
	Title        m_AhpSkeleton.Option
	Activity     m_AhpSkeleton.Option
	ContentUri   m_AhpSkeleton.Option
	Availability CanvasAvailability
}

func (CanvasState_CanvasState) isCanvasState() {}

func (CompanionStruct_CanvasState_) Create_CanvasState_(CanvasId _dafny.Sequence, ProviderId _dafny.Sequence, Input m_AhpSkeleton.Option, Title m_AhpSkeleton.Option, Activity m_AhpSkeleton.Option, ContentUri m_AhpSkeleton.Option, Availability CanvasAvailability) CanvasState {
	return CanvasState{CanvasState_CanvasState{CanvasId, ProviderId, Input, Title, Activity, ContentUri, Availability}}
}

func (_this CanvasState) Is_CanvasState() bool {
	_, ok := _this.Get_().(CanvasState_CanvasState)
	return ok
}

func (CompanionStruct_CanvasState_) Default() CanvasState {
	return Companion_CanvasState_.Create_CanvasState_(_dafny.EmptySeq, _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), Companion_CanvasAvailability_.Default())
}

func (_this CanvasState) Dtor_canvasId() _dafny.Sequence {
	return _this.Get_().(CanvasState_CanvasState).CanvasId
}

func (_this CanvasState) Dtor_providerId() _dafny.Sequence {
	return _this.Get_().(CanvasState_CanvasState).ProviderId
}

func (_this CanvasState) Dtor_input() m_AhpSkeleton.Option {
	return _this.Get_().(CanvasState_CanvasState).Input
}

func (_this CanvasState) Dtor_title() m_AhpSkeleton.Option {
	return _this.Get_().(CanvasState_CanvasState).Title
}

func (_this CanvasState) Dtor_activity() m_AhpSkeleton.Option {
	return _this.Get_().(CanvasState_CanvasState).Activity
}

func (_this CanvasState) Dtor_contentUri() m_AhpSkeleton.Option {
	return _this.Get_().(CanvasState_CanvasState).ContentUri
}

func (_this CanvasState) Dtor_availability() CanvasAvailability {
	return _this.Get_().(CanvasState_CanvasState).Availability
}

func (_this CanvasState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CanvasState_CanvasState:
		{
			return "Canvas.CanvasState.CanvasState" + "(" + data.CanvasId.VerbatimString(true) + ", " + data.ProviderId.VerbatimString(true) + ", " + _dafny.String(data.Input) + ", " + _dafny.String(data.Title) + ", " + _dafny.String(data.Activity) + ", " + _dafny.String(data.ContentUri) + ", " + _dafny.String(data.Availability) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CanvasState) Equals(other CanvasState) bool {
	switch data1 := _this.Get_().(type) {
	case CanvasState_CanvasState:
		{
			data2, ok := other.Get_().(CanvasState_CanvasState)
			return ok && data1.CanvasId.Equals(data2.CanvasId) && data1.ProviderId.Equals(data2.ProviderId) && data1.Input.Equals(data2.Input) && data1.Title.Equals(data2.Title) && data1.Activity.Equals(data2.Activity) && data1.ContentUri.Equals(data2.ContentUri) && data1.Availability.Equals(data2.Availability)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CanvasState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CanvasState)
	return ok && _this.Equals(typed)
}

func Type_CanvasState_() _dafny.TypeDescriptor {
	return type_CanvasState_{}
}

type type_CanvasState_ struct {
}

func (_this type_CanvasState_) Default() interface{} {
	return Companion_CanvasState_.Default()
}

func (_this type_CanvasState_) String() string {
	return "Canvas.CanvasState"
}
func (_this CanvasState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CanvasState{}

// End of datatype CanvasState

// Definition of datatype CanvasAction
type CanvasAction struct {
	Data_CanvasAction_
}

func (_this CanvasAction) Get_() Data_CanvasAction_ {
	return _this.Data_CanvasAction_
}

type Data_CanvasAction_ interface {
	isCanvasAction()
}

type CompanionStruct_CanvasAction_ struct {
}

var Companion_CanvasAction_ = CompanionStruct_CanvasAction_{}

type CanvasAction_Updated struct {
	Title        m_AhpSkeleton.Option
	Activity     m_AhpSkeleton.Option
	ContentUri   m_AhpSkeleton.Option
	Availability m_AhpSkeleton.Option
}

func (CanvasAction_Updated) isCanvasAction() {}

func (CompanionStruct_CanvasAction_) Create_Updated_(Title m_AhpSkeleton.Option, Activity m_AhpSkeleton.Option, ContentUri m_AhpSkeleton.Option, Availability m_AhpSkeleton.Option) CanvasAction {
	return CanvasAction{CanvasAction_Updated{Title, Activity, ContentUri, Availability}}
}

func (_this CanvasAction) Is_Updated() bool {
	_, ok := _this.Get_().(CanvasAction_Updated)
	return ok
}

type CanvasAction_CloseRequested struct {
}

func (CanvasAction_CloseRequested) isCanvasAction() {}

func (CompanionStruct_CanvasAction_) Create_CloseRequested_() CanvasAction {
	return CanvasAction{CanvasAction_CloseRequested{}}
}

func (_this CanvasAction) Is_CloseRequested() bool {
	_, ok := _this.Get_().(CanvasAction_CloseRequested)
	return ok
}

type CanvasAction_CanvasUnknown struct {
	Raw m_ConfluxCodec.Json
}

func (CanvasAction_CanvasUnknown) isCanvasAction() {}

func (CompanionStruct_CanvasAction_) Create_CanvasUnknown_(Raw m_ConfluxCodec.Json) CanvasAction {
	return CanvasAction{CanvasAction_CanvasUnknown{Raw}}
}

func (_this CanvasAction) Is_CanvasUnknown() bool {
	_, ok := _this.Get_().(CanvasAction_CanvasUnknown)
	return ok
}

func (CompanionStruct_CanvasAction_) Default() CanvasAction {
	return Companion_CanvasAction_.Create_Updated_(m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default())
}

func (_this CanvasAction) Dtor_title() m_AhpSkeleton.Option {
	return _this.Get_().(CanvasAction_Updated).Title
}

func (_this CanvasAction) Dtor_activity() m_AhpSkeleton.Option {
	return _this.Get_().(CanvasAction_Updated).Activity
}

func (_this CanvasAction) Dtor_contentUri() m_AhpSkeleton.Option {
	return _this.Get_().(CanvasAction_Updated).ContentUri
}

func (_this CanvasAction) Dtor_availability() m_AhpSkeleton.Option {
	return _this.Get_().(CanvasAction_Updated).Availability
}

func (_this CanvasAction) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(CanvasAction_CanvasUnknown).Raw
}

func (_this CanvasAction) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CanvasAction_Updated:
		{
			return "Canvas.CanvasAction.Updated" + "(" + _dafny.String(data.Title) + ", " + _dafny.String(data.Activity) + ", " + _dafny.String(data.ContentUri) + ", " + _dafny.String(data.Availability) + ")"
		}
	case CanvasAction_CloseRequested:
		{
			return "Canvas.CanvasAction.CloseRequested"
		}
	case CanvasAction_CanvasUnknown:
		{
			return "Canvas.CanvasAction.CanvasUnknown" + "(" + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CanvasAction) Equals(other CanvasAction) bool {
	switch data1 := _this.Get_().(type) {
	case CanvasAction_Updated:
		{
			data2, ok := other.Get_().(CanvasAction_Updated)
			return ok && data1.Title.Equals(data2.Title) && data1.Activity.Equals(data2.Activity) && data1.ContentUri.Equals(data2.ContentUri) && data1.Availability.Equals(data2.Availability)
		}
	case CanvasAction_CloseRequested:
		{
			_, ok := other.Get_().(CanvasAction_CloseRequested)
			return ok
		}
	case CanvasAction_CanvasUnknown:
		{
			data2, ok := other.Get_().(CanvasAction_CanvasUnknown)
			return ok && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CanvasAction) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CanvasAction)
	return ok && _this.Equals(typed)
}

func Type_CanvasAction_() _dafny.TypeDescriptor {
	return type_CanvasAction_{}
}

type type_CanvasAction_ struct {
}

func (_this type_CanvasAction_) Default() interface{} {
	return Companion_CanvasAction_.Default()
}

func (_this type_CanvasAction_) String() string {
	return "Canvas.CanvasAction"
}
func (_this CanvasAction) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CanvasAction{}

// End of datatype CanvasAction

// Definition of datatype CanvasInstance
type CanvasInstance struct {
	Data_CanvasInstance_
}

func (_this CanvasInstance) Get_() Data_CanvasInstance_ {
	return _this.Data_CanvasInstance_
}

type Data_CanvasInstance_ interface {
	isCanvasInstance()
}

type CompanionStruct_CanvasInstance_ struct {
}

var Companion_CanvasInstance_ = CompanionStruct_CanvasInstance_{}

type CanvasInstance_CanvasInstance struct {
	Channel  _dafny.Sequence
	Snapshot CanvasState
}

func (CanvasInstance_CanvasInstance) isCanvasInstance() {}

func (CompanionStruct_CanvasInstance_) Create_CanvasInstance_(Channel _dafny.Sequence, Snapshot CanvasState) CanvasInstance {
	return CanvasInstance{CanvasInstance_CanvasInstance{Channel, Snapshot}}
}

func (_this CanvasInstance) Is_CanvasInstance() bool {
	_, ok := _this.Get_().(CanvasInstance_CanvasInstance)
	return ok
}

func (CompanionStruct_CanvasInstance_) Default() CanvasInstance {
	return Companion_CanvasInstance_.Create_CanvasInstance_(_dafny.EmptySeq, Companion_CanvasState_.Default())
}

func (_this CanvasInstance) Dtor_channel() _dafny.Sequence {
	return _this.Get_().(CanvasInstance_CanvasInstance).Channel
}

func (_this CanvasInstance) Dtor_snapshot() CanvasState {
	return _this.Get_().(CanvasInstance_CanvasInstance).Snapshot
}

func (_this CanvasInstance) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CanvasInstance_CanvasInstance:
		{
			return "Canvas.CanvasInstance.CanvasInstance" + "(" + data.Channel.VerbatimString(true) + ", " + _dafny.String(data.Snapshot) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CanvasInstance) Equals(other CanvasInstance) bool {
	switch data1 := _this.Get_().(type) {
	case CanvasInstance_CanvasInstance:
		{
			data2, ok := other.Get_().(CanvasInstance_CanvasInstance)
			return ok && data1.Channel.Equals(data2.Channel) && data1.Snapshot.Equals(data2.Snapshot)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CanvasInstance) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CanvasInstance)
	return ok && _this.Equals(typed)
}

func Type_CanvasInstance_() _dafny.TypeDescriptor {
	return type_CanvasInstance_{}
}

type type_CanvasInstance_ struct {
}

func (_this type_CanvasInstance_) Default() interface{} {
	return Companion_CanvasInstance_.Default()
}

func (_this type_CanvasInstance_) String() string {
	return "Canvas.CanvasInstance"
}
func (_this CanvasInstance) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CanvasInstance{}

// End of datatype CanvasInstance

// Definition of datatype ResourceReadParams
type ResourceReadParams struct {
	Data_ResourceReadParams_
}

func (_this ResourceReadParams) Get_() Data_ResourceReadParams_ {
	return _this.Data_ResourceReadParams_
}

type Data_ResourceReadParams_ interface {
	isResourceReadParams()
}

type CompanionStruct_ResourceReadParams_ struct {
}

var Companion_ResourceReadParams_ = CompanionStruct_ResourceReadParams_{}

type ResourceReadParams_ResourceReadParams struct {
	Channel _dafny.Sequence
	Uri     _dafny.Sequence
}

func (ResourceReadParams_ResourceReadParams) isResourceReadParams() {}

func (CompanionStruct_ResourceReadParams_) Create_ResourceReadParams_(Channel _dafny.Sequence, Uri _dafny.Sequence) ResourceReadParams {
	return ResourceReadParams{ResourceReadParams_ResourceReadParams{Channel, Uri}}
}

func (_this ResourceReadParams) Is_ResourceReadParams() bool {
	_, ok := _this.Get_().(ResourceReadParams_ResourceReadParams)
	return ok
}

func (CompanionStruct_ResourceReadParams_) Default() ResourceReadParams {
	return Companion_ResourceReadParams_.Create_ResourceReadParams_(_dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this ResourceReadParams) Dtor_channel() _dafny.Sequence {
	return _this.Get_().(ResourceReadParams_ResourceReadParams).Channel
}

func (_this ResourceReadParams) Dtor_uri() _dafny.Sequence {
	return _this.Get_().(ResourceReadParams_ResourceReadParams).Uri
}

func (_this ResourceReadParams) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ResourceReadParams_ResourceReadParams:
		{
			return "Canvas.ResourceReadParams.ResourceReadParams" + "(" + data.Channel.VerbatimString(true) + ", " + data.Uri.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ResourceReadParams) Equals(other ResourceReadParams) bool {
	switch data1 := _this.Get_().(type) {
	case ResourceReadParams_ResourceReadParams:
		{
			data2, ok := other.Get_().(ResourceReadParams_ResourceReadParams)
			return ok && data1.Channel.Equals(data2.Channel) && data1.Uri.Equals(data2.Uri)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ResourceReadParams) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ResourceReadParams)
	return ok && _this.Equals(typed)
}

func Type_ResourceReadParams_() _dafny.TypeDescriptor {
	return type_ResourceReadParams_{}
}

type type_ResourceReadParams_ struct {
}

func (_this type_ResourceReadParams_) Default() interface{} {
	return Companion_ResourceReadParams_.Default()
}

func (_this type_ResourceReadParams_) String() string {
	return "Canvas.ResourceReadParams"
}
func (_this ResourceReadParams) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ResourceReadParams{}

// End of datatype ResourceReadParams

// Definition of datatype CanvasAdmission
type CanvasAdmission struct {
	Data_CanvasAdmission_
}

func (_this CanvasAdmission) Get_() Data_CanvasAdmission_ {
	return _this.Data_CanvasAdmission_
}

type Data_CanvasAdmission_ interface {
	isCanvasAdmission()
}

type CompanionStruct_CanvasAdmission_ struct {
}

var Companion_CanvasAdmission_ = CompanionStruct_CanvasAdmission_{}

type CanvasAdmission_CanvasSupported struct {
}

func (CanvasAdmission_CanvasSupported) isCanvasAdmission() {}

func (CompanionStruct_CanvasAdmission_) Create_CanvasSupported_() CanvasAdmission {
	return CanvasAdmission{CanvasAdmission_CanvasSupported{}}
}

func (_this CanvasAdmission) Is_CanvasSupported() bool {
	_, ok := _this.Get_().(CanvasAdmission_CanvasSupported)
	return ok
}

type CanvasAdmission_CanvasUnsupported struct {
}

func (CanvasAdmission_CanvasUnsupported) isCanvasAdmission() {}

func (CompanionStruct_CanvasAdmission_) Create_CanvasUnsupported_() CanvasAdmission {
	return CanvasAdmission{CanvasAdmission_CanvasUnsupported{}}
}

func (_this CanvasAdmission) Is_CanvasUnsupported() bool {
	_, ok := _this.Get_().(CanvasAdmission_CanvasUnsupported)
	return ok
}

type CanvasAdmission_CanvasHandlerUnavailable struct {
}

func (CanvasAdmission_CanvasHandlerUnavailable) isCanvasAdmission() {}

func (CompanionStruct_CanvasAdmission_) Create_CanvasHandlerUnavailable_() CanvasAdmission {
	return CanvasAdmission{CanvasAdmission_CanvasHandlerUnavailable{}}
}

func (_this CanvasAdmission) Is_CanvasHandlerUnavailable() bool {
	_, ok := _this.Get_().(CanvasAdmission_CanvasHandlerUnavailable)
	return ok
}

type CanvasAdmission_CanvasUnauthorized struct {
}

func (CanvasAdmission_CanvasUnauthorized) isCanvasAdmission() {}

func (CompanionStruct_CanvasAdmission_) Create_CanvasUnauthorized_() CanvasAdmission {
	return CanvasAdmission{CanvasAdmission_CanvasUnauthorized{}}
}

func (_this CanvasAdmission) Is_CanvasUnauthorized() bool {
	_, ok := _this.Get_().(CanvasAdmission_CanvasUnauthorized)
	return ok
}

type CanvasAdmission_CanvasUriCollision struct {
}

func (CanvasAdmission_CanvasUriCollision) isCanvasAdmission() {}

func (CompanionStruct_CanvasAdmission_) Create_CanvasUriCollision_() CanvasAdmission {
	return CanvasAdmission{CanvasAdmission_CanvasUriCollision{}}
}

func (_this CanvasAdmission) Is_CanvasUriCollision() bool {
	_, ok := _this.Get_().(CanvasAdmission_CanvasUriCollision)
	return ok
}

func (CompanionStruct_CanvasAdmission_) Default() CanvasAdmission {
	return Companion_CanvasAdmission_.Create_CanvasSupported_()
}

func (_ CompanionStruct_CanvasAdmission_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_CanvasAdmission_.Create_CanvasSupported_(), true
		case 1:
			return Companion_CanvasAdmission_.Create_CanvasUnsupported_(), true
		case 2:
			return Companion_CanvasAdmission_.Create_CanvasHandlerUnavailable_(), true
		case 3:
			return Companion_CanvasAdmission_.Create_CanvasUnauthorized_(), true
		case 4:
			return Companion_CanvasAdmission_.Create_CanvasUriCollision_(), true
		default:
			return CanvasAdmission{}, false
		}
	}
}

func (_this CanvasAdmission) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case CanvasAdmission_CanvasSupported:
		{
			return "Canvas.CanvasAdmission.CanvasSupported"
		}
	case CanvasAdmission_CanvasUnsupported:
		{
			return "Canvas.CanvasAdmission.CanvasUnsupported"
		}
	case CanvasAdmission_CanvasHandlerUnavailable:
		{
			return "Canvas.CanvasAdmission.CanvasHandlerUnavailable"
		}
	case CanvasAdmission_CanvasUnauthorized:
		{
			return "Canvas.CanvasAdmission.CanvasUnauthorized"
		}
	case CanvasAdmission_CanvasUriCollision:
		{
			return "Canvas.CanvasAdmission.CanvasUriCollision"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CanvasAdmission) Equals(other CanvasAdmission) bool {
	switch _this.Get_().(type) {
	case CanvasAdmission_CanvasSupported:
		{
			_, ok := other.Get_().(CanvasAdmission_CanvasSupported)
			return ok
		}
	case CanvasAdmission_CanvasUnsupported:
		{
			_, ok := other.Get_().(CanvasAdmission_CanvasUnsupported)
			return ok
		}
	case CanvasAdmission_CanvasHandlerUnavailable:
		{
			_, ok := other.Get_().(CanvasAdmission_CanvasHandlerUnavailable)
			return ok
		}
	case CanvasAdmission_CanvasUnauthorized:
		{
			_, ok := other.Get_().(CanvasAdmission_CanvasUnauthorized)
			return ok
		}
	case CanvasAdmission_CanvasUriCollision:
		{
			_, ok := other.Get_().(CanvasAdmission_CanvasUriCollision)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CanvasAdmission) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CanvasAdmission)
	return ok && _this.Equals(typed)
}

func Type_CanvasAdmission_() _dafny.TypeDescriptor {
	return type_CanvasAdmission_{}
}

type type_CanvasAdmission_ struct {
}

func (_this type_CanvasAdmission_) Default() interface{} {
	return Companion_CanvasAdmission_.Default()
}

func (_this type_CanvasAdmission_) String() string {
	return "Canvas.CanvasAdmission"
}
func (_this CanvasAdmission) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CanvasAdmission{}

// End of datatype CanvasAdmission

// Definition of datatype CanvasProviderErrorData
type CanvasProviderErrorData struct {
	Data_CanvasProviderErrorData_
}

func (_this CanvasProviderErrorData) Get_() Data_CanvasProviderErrorData_ {
	return _this.Data_CanvasProviderErrorData_
}

type Data_CanvasProviderErrorData_ interface {
	isCanvasProviderErrorData()
}

type CompanionStruct_CanvasProviderErrorData_ struct {
}

var Companion_CanvasProviderErrorData_ = CompanionStruct_CanvasProviderErrorData_{}

type CanvasProviderErrorData_CanvasProviderErrorData struct {
	Code    _dafny.Sequence
	Message _dafny.Sequence
}

func (CanvasProviderErrorData_CanvasProviderErrorData) isCanvasProviderErrorData() {}

func (CompanionStruct_CanvasProviderErrorData_) Create_CanvasProviderErrorData_(Code _dafny.Sequence, Message _dafny.Sequence) CanvasProviderErrorData {
	return CanvasProviderErrorData{CanvasProviderErrorData_CanvasProviderErrorData{Code, Message}}
}

func (_this CanvasProviderErrorData) Is_CanvasProviderErrorData() bool {
	_, ok := _this.Get_().(CanvasProviderErrorData_CanvasProviderErrorData)
	return ok
}

func (CompanionStruct_CanvasProviderErrorData_) Default() CanvasProviderErrorData {
	return Companion_CanvasProviderErrorData_.Create_CanvasProviderErrorData_(_dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this CanvasProviderErrorData) Dtor_code() _dafny.Sequence {
	return _this.Get_().(CanvasProviderErrorData_CanvasProviderErrorData).Code
}

func (_this CanvasProviderErrorData) Dtor_message() _dafny.Sequence {
	return _this.Get_().(CanvasProviderErrorData_CanvasProviderErrorData).Message
}

func (_this CanvasProviderErrorData) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CanvasProviderErrorData_CanvasProviderErrorData:
		{
			return "Canvas.CanvasProviderErrorData.CanvasProviderErrorData" + "(" + data.Code.VerbatimString(true) + ", " + data.Message.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CanvasProviderErrorData) Equals(other CanvasProviderErrorData) bool {
	switch data1 := _this.Get_().(type) {
	case CanvasProviderErrorData_CanvasProviderErrorData:
		{
			data2, ok := other.Get_().(CanvasProviderErrorData_CanvasProviderErrorData)
			return ok && data1.Code.Equals(data2.Code) && data1.Message.Equals(data2.Message)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CanvasProviderErrorData) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CanvasProviderErrorData)
	return ok && _this.Equals(typed)
}

func Type_CanvasProviderErrorData_() _dafny.TypeDescriptor {
	return type_CanvasProviderErrorData_{}
}

type type_CanvasProviderErrorData_ struct {
}

func (_this type_CanvasProviderErrorData_) Default() interface{} {
	return Companion_CanvasProviderErrorData_.Default()
}

func (_this type_CanvasProviderErrorData_) String() string {
	return "Canvas.CanvasProviderErrorData"
}
func (_this CanvasProviderErrorData) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CanvasProviderErrorData{}

// End of datatype CanvasProviderErrorData

// Definition of datatype CanvasError
type CanvasError struct {
	Data_CanvasError_
}

func (_this CanvasError) Get_() Data_CanvasError_ {
	return _this.Data_CanvasError_
}

type Data_CanvasError_ interface {
	isCanvasError()
}

type CompanionStruct_CanvasError_ struct {
}

var Companion_CanvasError_ = CompanionStruct_CanvasError_{}

type CanvasError_AdmissionError struct {
	Reason CanvasAdmission
}

func (CanvasError_AdmissionError) isCanvasError() {}

func (CompanionStruct_CanvasError_) Create_AdmissionError_(Reason CanvasAdmission) CanvasError {
	return CanvasError{CanvasError_AdmissionError{Reason}}
}

func (_this CanvasError) Is_AdmissionError() bool {
	_, ok := _this.Get_().(CanvasError_AdmissionError)
	return ok
}

type CanvasError_ProviderError struct {
	ErrorCode _dafny.Int
	Data      CanvasProviderErrorData
}

func (CanvasError_ProviderError) isCanvasError() {}

func (CompanionStruct_CanvasError_) Create_ProviderError_(ErrorCode _dafny.Int, Data CanvasProviderErrorData) CanvasError {
	return CanvasError{CanvasError_ProviderError{ErrorCode, Data}}
}

func (_this CanvasError) Is_ProviderError() bool {
	_, ok := _this.Get_().(CanvasError_ProviderError)
	return ok
}

func (CompanionStruct_CanvasError_) Default() CanvasError {
	return Companion_CanvasError_.Create_AdmissionError_(Companion_CanvasAdmission_.Default())
}

func (_this CanvasError) Dtor_reason() CanvasAdmission {
	return _this.Get_().(CanvasError_AdmissionError).Reason
}

func (_this CanvasError) Dtor_errorCode() _dafny.Int {
	return _this.Get_().(CanvasError_ProviderError).ErrorCode
}

func (_this CanvasError) Dtor_data() CanvasProviderErrorData {
	return _this.Get_().(CanvasError_ProviderError).Data
}

func (_this CanvasError) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CanvasError_AdmissionError:
		{
			return "Canvas.CanvasError.AdmissionError" + "(" + _dafny.String(data.Reason) + ")"
		}
	case CanvasError_ProviderError:
		{
			return "Canvas.CanvasError.ProviderError" + "(" + _dafny.String(data.ErrorCode) + ", " + _dafny.String(data.Data) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CanvasError) Equals(other CanvasError) bool {
	switch data1 := _this.Get_().(type) {
	case CanvasError_AdmissionError:
		{
			data2, ok := other.Get_().(CanvasError_AdmissionError)
			return ok && data1.Reason.Equals(data2.Reason)
		}
	case CanvasError_ProviderError:
		{
			data2, ok := other.Get_().(CanvasError_ProviderError)
			return ok && data1.ErrorCode.Cmp(data2.ErrorCode) == 0 && data1.Data.Equals(data2.Data)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CanvasError) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CanvasError)
	return ok && _this.Equals(typed)
}

func Type_CanvasError_() _dafny.TypeDescriptor {
	return type_CanvasError_{}
}

type type_CanvasError_ struct {
}

func (_this type_CanvasError_) Default() interface{} {
	return Companion_CanvasError_.Default()
}

func (_this type_CanvasError_) String() string {
	return "Canvas.CanvasError"
}
func (_this CanvasError) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CanvasError{}

// End of datatype CanvasError
