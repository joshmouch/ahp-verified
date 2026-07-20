// Package Chat
// Dafny module Chat compiled into Go

package Chat

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-verified/go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-verified/go/Changeset"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ConfluxOrderedLog "github.com/joshmouch/ahp-verified/go/ConfluxOrderedLog"
	m_ConfluxPrune "github.com/joshmouch/ahp-verified/go/ConfluxPrune"
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
	return "Chat.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) SetBit(s uint32, b uint32) uint32 {
	return (s) | (b)
}
func (_static *CompanionStruct_Default___) ClearBit(s uint32, b uint32) uint32 {
	return (s) & (^(b))
}
func (_static *CompanionStruct_Default___) JNoNull(j m_ConfluxCodec.Json) m_AhpSkeleton.Option {
	if (j).Is_JNull() {
		return m_AhpSkeleton.Companion_Option_.Create_None_()
	} else {
		return m_AhpSkeleton.Companion_Option_.Create_Some_(j)
	}
}
func (_static *CompanionStruct_Default___) UpsertQ(qs _dafny.Sequence, m QMsg) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpsertById(func(coer48 func(QMsg) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg56 interface{}) interface{} {
			return coer48(arg56.(QMsg))
		}
	}(Companion_Default___.QKey()), qs, m)
}
func (_static *CompanionStruct_Default___) RemoveQ(qs _dafny.Sequence, id _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqRemoveById(func(coer49 func(QMsg) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg57 interface{}) interface{} {
			return coer49(arg57.(QMsg))
		}
	}(Companion_Default___.QKey()), qs, id)
}
func (_static *CompanionStruct_Default___) ContainsQ(qs _dafny.Sequence, id _dafny.Sequence) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((qs).Cardinality())).Sign() == 0 {
		return false
	} else if _dafny.Companion_Sequence_.Equal(((qs).Select(0).(QMsg)).Dtor_id(), id) {
		return true
	} else {
		var _in0 _dafny.Sequence = (qs).Drop(1)
		_ = _in0
		var _in1 _dafny.Sequence = id
		_ = _in1
		qs = _in0
		id = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) FindQ(qs _dafny.Sequence, id _dafny.Sequence) _dafny.Sequence {
	var _0_found _dafny.Sequence = m_ConfluxOperators.Companion_Default___.Filter(func(coer50 func(QMsg) bool) func(interface{}) bool {
		return func(arg58 interface{}) bool {
			return coer50(arg58.(QMsg))
		}
	}((func(_1_id _dafny.Sequence) func(QMsg) bool {
		return func(_2_q QMsg) bool {
			return _dafny.Companion_Sequence_.Equal((_2_q).Dtor_id(), _1_id)
		}
	})(id)), qs)
	_ = _0_found
	if (_dafny.IntOfUint32((_0_found).Cardinality())).Sign() == 0 {
		return _dafny.SeqOf()
	} else {
		return _dafny.SeqOf((_0_found).Select(0).(QMsg))
	}
}
func (_static *CompanionStruct_Default___) ReorderQ(qs _dafny.Sequence, order _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxOrderedLog.Companion_Default___.SeqReorderByKeys(func(coer51 func(QMsg) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg59 interface{}) interface{} {
			return coer51(arg59.(QMsg))
		}
	}(Companion_Default___.QKey()), order, qs)
}
func (_static *CompanionStruct_Default___) KeepUpTo(ts _dafny.Sequence, id _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxOrderedLog.Companion_Default___.SeqKeepThrough(func(coer52 func(Turn) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg60 interface{}) interface{} {
			return coer52(arg60.(Turn))
		}
	}(func(_0_t Turn) _dafny.Sequence {
		return (_0_t).Dtor_turnId()
	}), id, ts)
}
func (_static *CompanionStruct_Default___) HasTurn(ts _dafny.Sequence, id _dafny.Sequence) bool {
	return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((ts).Cardinality())), false, func(_exists_var_0 _dafny.Int) bool {
		var _0_i _dafny.Int
		_0_i = interface{}(_exists_var_0).(_dafny.Int)
		return (((_0_i).Sign() != -1) && ((_0_i).Cmp(_dafny.IntOfUint32((ts).Cardinality())) < 0)) && (_dafny.Companion_Sequence_.Equal(((ts).Select((_0_i).Uint32()).(Turn)).Dtor_turnId(), id))
	})
}
func (_static *CompanionStruct_Default___) DedupPrepend(loaded _dafny.Sequence, existing _dafny.Sequence) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(m_ConfluxOperators.Companion_Default___.Filter(func(coer53 func(Turn) bool) func(interface{}) bool {
		return func(arg61 interface{}) bool {
			return coer53(arg61.(Turn))
		}
	}((func(_0_existing _dafny.Sequence) func(Turn) bool {
		return func(_1_t Turn) bool {
			return !(Companion_Default___.HasTurn(_0_existing, (_1_t).Dtor_turnId()))
		}
	})(existing)), loaded), existing)
}
func (_static *CompanionStruct_Default___) IsOpenInput(p Part) bool {
	return ((p).Is_PInputRequest()) && (((p).Dtor_response()).Is_None())
}
func (_static *CompanionStruct_Default___) HasOpenReq(ps _dafny.Sequence, id _dafny.Sequence) bool {
	if (_dafny.IntOfUint32((ps).Cardinality())).Sign() == 0 {
		return false
	} else {
		return (((((ps).Select(0).(Part)).Is_PInputRequest()) && ((((ps).Select(0).(Part)).Dtor_response()).Is_None())) && (_dafny.Companion_Sequence_.Equal((((ps).Select(0).(Part)).Dtor_req()).Dtor_id(), id))) || (Companion_Default___.HasOpenReq((ps).Drop(1), id))
	}
}
func (_static *CompanionStruct_Default___) MergeInputReq(existing InputReq, req InputReq) InputReq {
	if (((req).Dtor_answers()).Cardinality()).Sign() == 1 {
		return req
	} else {
		var _0_dt__update__tmp_h0 InputReq = req
		_ = _0_dt__update__tmp_h0
		var _1_dt__update_hanswers_h0 _dafny.Map = (existing).Dtor_answers()
		_ = _1_dt__update_hanswers_h0
		return Companion_InputReq_.Create_InputReq_((_0_dt__update__tmp_h0).Dtor_id(), _1_dt__update_hanswers_h0, (_0_dt__update__tmp_h0).Dtor_open())
	}
}
func (_static *CompanionStruct_Default___) MergeInputPart(req InputReq) func(Part) Part {
	return (func(_0_req InputReq) func(Part) Part {
		return func(_1_p Part) Part {
			return func() Part {
				var _source0 Part = _1_p
				_ = _source0
				{
					if _source0.Is_PInputRequest() {
						var _2_existing InputReq = _source0.Get_().(Part_PInputRequest).Req
						_ = _2_existing
						return Companion_Part_.Create_PInputRequest_(Companion_Default___.MergeInputReq(_2_existing, _0_req), m_AhpSkeleton.Companion_Option_.Create_None_())
					}
				}
				{
					return _1_p
				}
			}()
		}
	})(req)
}
func (_static *CompanionStruct_Default___) UpsertInputPart(ps _dafny.Sequence, req InputReq) _dafny.Sequence {
	if Companion_Default___.HasOpenReq(ps, (req).Dtor_id()) {
		return m_ConfluxSeqRoute.Companion_Default___.SeqUpdateByIdWhere(func(coer54 func(Part) _dafny.Sequence) func(interface{}) interface{} {
			return func(arg62 interface{}) interface{} {
				return coer54(arg62.(Part))
			}
		}(Companion_Default___.PartKey), func(coer55 func(Part) bool) func(interface{}) bool {
			return func(arg63 interface{}) bool {
				return coer55(arg63.(Part))
			}
		}(Companion_Default___.IsOpenInput), ps, (req).Dtor_id(), func(coer56 func(Part) Part) func(interface{}) interface{} {
			return func(arg64 interface{}) interface{} {
				return coer56(arg64.(Part))
			}
		}(Companion_Default___.MergeInputPart(req)))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(ps, _dafny.SeqOf(Companion_Part_.Create_PInputRequest_(req, m_AhpSkeleton.Companion_Option_.Create_None_())))
	}
}
func (_static *CompanionStruct_Default___) AnswerInputPart(qId _dafny.Sequence, ans m_AhpSkeleton.Option) func(Part) Part {
	return (func(_0_ans m_AhpSkeleton.Option, _1_qId _dafny.Sequence) func(Part) Part {
		return func(_2_p Part) Part {
			return func() Part {
				var _source0 Part = _2_p
				_ = _source0
				{
					if _source0.Is_PInputRequest() {
						var _3_r InputReq = _source0.Get_().(Part_PInputRequest).Req
						_ = _3_r
						var _4_resp m_AhpSkeleton.Option = _source0.Get_().(Part_PInputRequest).Response
						_ = _4_resp
						return Companion_Part_.Create_PInputRequest_(func(_pat_let233_0 InputReq) InputReq {
							return func(_5_dt__update__tmp_h0 InputReq) InputReq {
								return func(_pat_let234_0 _dafny.Map) InputReq {
									return func(_6_dt__update_hanswers_h0 _dafny.Map) InputReq {
										return Companion_InputReq_.Create_InputReq_((_5_dt__update__tmp_h0).Dtor_id(), _6_dt__update_hanswers_h0, (_5_dt__update__tmp_h0).Dtor_open())
									}(_pat_let234_0)
								}((func() _dafny.Map {
									if (_0_ans).Is_Some() {
										return ((_3_r).Dtor_answers()).Update(_1_qId, (_0_ans).Dtor_value().(m_ConfluxCodec.Json))
									}
									return ((_3_r).Dtor_answers()).Subtract(_dafny.SetOf(_1_qId))
								})())
							}(_pat_let233_0)
						}(_3_r), _4_resp)
					}
				}
				{
					return _2_p
				}
			}()
		}
	})(ans, qId)
}
func (_static *CompanionStruct_Default___) ChangeAnswerPart(ps _dafny.Sequence, reqId _dafny.Sequence, qId _dafny.Sequence, ans m_AhpSkeleton.Option) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpdateByIdWhere(func(coer57 func(Part) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg65 interface{}) interface{} {
			return coer57(arg65.(Part))
		}
	}(Companion_Default___.PartKey), func(coer58 func(Part) bool) func(interface{}) bool {
		return func(arg66 interface{}) bool {
			return coer58(arg66.(Part))
		}
	}(Companion_Default___.IsOpenInput), ps, reqId, func(coer59 func(Part) Part) func(interface{}) interface{} {
		return func(arg67 interface{}) interface{} {
			return coer59(arg67.(Part))
		}
	}(Companion_Default___.AnswerInputPart(qId, ans)))
}
func (_static *CompanionStruct_Default___) MergeAnswers(drafts _dafny.Map, completed _dafny.Map) _dafny.Map {
	return func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((((drafts).Keys()).Union((completed).Keys())).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_k _dafny.Sequence
			_0_k = interface{}(_compr_0).(_dafny.Sequence)
			if (((drafts).Keys()).Union((completed).Keys())).Contains(_0_k) {
				_coll0.Add(_0_k, (func() m_ConfluxCodec.Json {
					if (completed).Contains(_0_k) {
						return (completed).Get(_0_k).(m_ConfluxCodec.Json)
					}
					return (drafts).Get(_0_k).(m_ConfluxCodec.Json)
				})())
			}
		}
		return _coll0.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) CompleteInputPart(resp _dafny.Sequence, answers _dafny.Map) func(Part) Part {
	return (func(_0_answers _dafny.Map, _1_resp _dafny.Sequence) func(Part) Part {
		return func(_2_p Part) Part {
			return func() Part {
				var _source0 Part = _2_p
				_ = _source0
				{
					if _source0.Is_PInputRequest() {
						var _3_r InputReq = _source0.Get_().(Part_PInputRequest).Req
						_ = _3_r
						return Companion_Part_.Create_PInputRequest_(func(_pat_let235_0 InputReq) InputReq {
							return func(_4_dt__update__tmp_h0 InputReq) InputReq {
								return func(_pat_let236_0 _dafny.Map) InputReq {
									return func(_5_dt__update_hanswers_h0 _dafny.Map) InputReq {
										return Companion_InputReq_.Create_InputReq_((_4_dt__update__tmp_h0).Dtor_id(), _5_dt__update_hanswers_h0, (_4_dt__update__tmp_h0).Dtor_open())
									}(_pat_let236_0)
								}(Companion_Default___.MergeAnswers((_3_r).Dtor_answers(), _0_answers))
							}(_pat_let235_0)
						}(_3_r), m_AhpSkeleton.Companion_Option_.Create_Some_(_1_resp))
					}
				}
				{
					return _2_p
				}
			}()
		}
	})(answers, resp)
}
func (_static *CompanionStruct_Default___) CompleteAnswerPart(ps _dafny.Sequence, reqId _dafny.Sequence, resp _dafny.Sequence, answers _dafny.Map) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpdateByIdWhere(func(coer60 func(Part) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg68 interface{}) interface{} {
			return coer60(arg68.(Part))
		}
	}(Companion_Default___.PartKey), func(coer61 func(Part) bool) func(interface{}) bool {
		return func(arg69 interface{}) bool {
			return coer61(arg69.(Part))
		}
	}(Companion_Default___.IsOpenInput), ps, reqId, func(coer62 func(Part) Part) func(interface{}) interface{} {
		return func(arg70 interface{}) interface{} {
			return coer62(arg70.(Part))
		}
	}(Companion_Default___.CompleteInputPart(resp, answers)))
}
func (_static *CompanionStruct_Default___) TurnMatches(s ChatState, id _dafny.Sequence) bool {
	return (((s).Dtor_activeTurn()).Is_Some()) && (_dafny.Companion_Sequence_.Equal((((s).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_turnId(), id))
}
func (_static *CompanionStruct_Default___) OptOr(o m_AhpSkeleton.Option, d _dafny.Sequence) _dafny.Sequence {
	if (o).Is_Some() {
		return (o).Dtor_value().(_dafny.Sequence)
	} else {
		return d
	}
}
func (_static *CompanionStruct_Default___) AnyPendingTC(ps _dafny.Sequence) bool {
	if (_dafny.IntOfUint32((ps).Cardinality())).Sign() == 0 {
		return false
	} else {
		return ((((ps).Select(0).(Part)).Is_PToolCall()) && (((_dafny.Companion_Sequence_.Equal((((ps).Select(0).(Part)).Dtor_tc()).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("pending-confirmation"))) || (_dafny.Companion_Sequence_.Equal((((ps).Select(0).(Part)).Dtor_tc()).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("pending-result-confirmation")))) || (_dafny.Companion_Sequence_.Equal((((ps).Select(0).(Part)).Dtor_tc()).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("auth-required"))))) || (Companion_Default___.AnyPendingTC((ps).Drop(1)))
	}
}
func (_static *CompanionStruct_Default___) HasPendingTCState(s ChatState) bool {
	return (((s).Dtor_activeTurn()).Is_Some()) && (Companion_Default___.AnyPendingTC((((s).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts()))
}
func (_static *CompanionStruct_Default___) AnyOpenInput(ps _dafny.Sequence) bool {
	if (_dafny.IntOfUint32((ps).Cardinality())).Sign() == 0 {
		return false
	} else {
		return ((((ps).Select(0).(Part)).Is_PInputRequest()) && ((((ps).Select(0).(Part)).Dtor_response()).Is_None())) || (Companion_Default___.AnyOpenInput((ps).Drop(1)))
	}
}
func (_static *CompanionStruct_Default___) HasOpenInputState(s ChatState) bool {
	return (((s).Dtor_activeTurn()).Is_Some()) && (Companion_Default___.AnyOpenInput((((s).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts()))
}
func (_static *CompanionStruct_Default___) ActivityBits(s ChatState, isError bool) uint32 {
	if isError {
		return Companion_Default___.ERR()
	} else if (Companion_Default___.HasOpenInputState(s)) || (Companion_Default___.HasPendingTCState(s)) {
		return Companion_Default___.INP()
	} else if ((s).Dtor_activeTurn()).Is_Some() {
		return Companion_Default___.GEN()
	} else {
		return Companion_Default___.IDLE()
	}
}
func (_static *CompanionStruct_Default___) SummaryStatus(s ChatState, isError bool) uint32 {
	return (((s).Dtor_status()) & (^(Companion_Default___.ACT()))) | (Companion_Default___.ActivityBits(s, isError))
}
func (_static *CompanionStruct_Default___) AppendPart(t Turn, p Part) Turn {
	var _0_dt__update__tmp_h0 Turn = t
	_ = _0_dt__update__tmp_h0
	var _1_dt__update_hparts_h0 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((t).Dtor_parts(), _dafny.SeqOf(p))
	_ = _1_dt__update_hparts_h0
	return Companion_Turn_.Create_Turn_((_0_dt__update__tmp_h0).Dtor_turnId(), (_0_dt__update__tmp_h0).Dtor_message(), _1_dt__update_hparts_h0, (_0_dt__update__tmp_h0).Dtor_state(), (_0_dt__update__tmp_h0).Dtor_usage(), (_0_dt__update__tmp_h0).Dtor_error())
}
func (_static *CompanionStruct_Default___) PartKey(p Part) _dafny.Sequence {
	var _source0 Part = p
	_ = _source0
	{
		if _source0.Is_PMarkdown() {
			var _0_pid _dafny.Sequence = _source0.Get_().(Part_PMarkdown).Id
			_ = _0_pid
			return _0_pid
		}
	}
	{
		if _source0.Is_PReasoning() {
			var _1_pid _dafny.Sequence = _source0.Get_().(Part_PReasoning).Id
			_ = _1_pid
			return _1_pid
		}
	}
	{
		if _source0.Is_PToolCall() {
			var _2_tc ToolCall = _source0.Get_().(Part_PToolCall).Tc
			_ = _2_tc
			return (_2_tc).Dtor_toolCallId()
		}
	}
	{
		var _3_req InputReq = _source0.Get_().(Part_PInputRequest).Req
		_ = _3_req
		return (_3_req).Dtor_id()
	}
}
func (_static *CompanionStruct_Default___) DeltaMarkdown(ps _dafny.Sequence, id _dafny.Sequence, content _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpdateAllWhere(func(coer63 func(Part) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg71 interface{}) interface{} {
			return coer63(arg71.(Part))
		}
	}(Companion_Default___.PartKey), func(coer64 func(Part) bool) func(interface{}) bool {
		return func(arg72 interface{}) bool {
			return coer64(arg72.(Part))
		}
	}(func(_0_p Part) bool {
		return (_0_p).Is_PMarkdown()
	}), ps, id, func(coer65 func(Part) Part) func(interface{}) interface{} {
		return func(arg73 interface{}) interface{} {
			return coer65(arg73.(Part))
		}
	}((func(_1_content _dafny.Sequence) func(Part) Part {
		return func(_2_p Part) Part {
			return func() Part {
				var _source0 Part = _2_p
				_ = _source0
				{
					if _source0.Is_PMarkdown() {
						var _3_pid _dafny.Sequence = _source0.Get_().(Part_PMarkdown).Id
						_ = _3_pid
						var _4_c _dafny.Sequence = _source0.Get_().(Part_PMarkdown).Content
						_ = _4_c
						return Companion_Part_.Create_PMarkdown_(_3_pid, _dafny.Companion_Sequence_.Concatenate(_4_c, _1_content))
					}
				}
				{
					return _2_p
				}
			}()
		}
	})(content)))
}
func (_static *CompanionStruct_Default___) DeltaReasoning(ps _dafny.Sequence, id _dafny.Sequence, content _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpdateAllWhere(func(coer66 func(Part) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg74 interface{}) interface{} {
			return coer66(arg74.(Part))
		}
	}(Companion_Default___.PartKey), func(coer67 func(Part) bool) func(interface{}) bool {
		return func(arg75 interface{}) bool {
			return coer67(arg75.(Part))
		}
	}(func(_0_p Part) bool {
		return (_0_p).Is_PReasoning()
	}), ps, id, func(coer68 func(Part) Part) func(interface{}) interface{} {
		return func(arg76 interface{}) interface{} {
			return coer68(arg76.(Part))
		}
	}((func(_1_content _dafny.Sequence) func(Part) Part {
		return func(_2_p Part) Part {
			return func() Part {
				var _source0 Part = _2_p
				_ = _source0
				{
					if _source0.Is_PReasoning() {
						var _3_pid _dafny.Sequence = _source0.Get_().(Part_PReasoning).Id
						_ = _3_pid
						var _4_c _dafny.Sequence = _source0.Get_().(Part_PReasoning).Content
						_ = _4_c
						return Companion_Part_.Create_PReasoning_(_3_pid, _dafny.Companion_Sequence_.Concatenate(_4_c, _1_content))
					}
				}
				{
					return _2_p
				}
			}()
		}
	})(content)))
}
func (_static *CompanionStruct_Default___) MapTC(ps _dafny.Sequence, id _dafny.Sequence, f func(ToolCall) ToolCall) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpdateAllWhere(func(coer69 func(Part) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg77 interface{}) interface{} {
			return coer69(arg77.(Part))
		}
	}(Companion_Default___.PartKey), func(coer70 func(Part) bool) func(interface{}) bool {
		return func(arg78 interface{}) bool {
			return coer70(arg78.(Part))
		}
	}(func(_0_p Part) bool {
		return (_0_p).Is_PToolCall()
	}), ps, id, func(coer71 func(Part) Part) func(interface{}) interface{} {
		return func(arg79 interface{}) interface{} {
			return coer71(arg79.(Part))
		}
	}((func(_1_f func(ToolCall) ToolCall) func(Part) Part {
		return func(_2_p Part) Part {
			return func() Part {
				var _source0 Part = _2_p
				_ = _source0
				{
					if _source0.Is_PToolCall() {
						var _3_tc ToolCall = _source0.Get_().(Part_PToolCall).Tc
						_ = _3_tc
						return Companion_Part_.Create_PToolCall_((_1_f)(_3_tc))
					}
				}
				{
					return _2_p
				}
			}()
		}
	})(f)))
}
func (_static *CompanionStruct_Default___) MetaOr(meta m_AhpSkeleton.Option, prior m_AhpSkeleton.Option) m_AhpSkeleton.Option {
	if (meta).Is_Some() {
		return meta
	} else {
		return prior
	}
}
func (_static *CompanionStruct_Default___) OptionId(j m_ConfluxCodec.Json) _dafny.Sequence {
	if (((j).Is_JObj()) && (((j).Dtor_fields()).Contains(_dafny.UnicodeSeqOfUtf8Bytes("id")))) && ((((j).Dtor_fields()).Get(_dafny.UnicodeSeqOfUtf8Bytes("id")).(m_ConfluxCodec.Json)).Is_JStr()) {
		return (((j).Dtor_fields()).Get(_dafny.UnicodeSeqOfUtf8Bytes("id")).(m_ConfluxCodec.Json)).Dtor_s()
	} else {
		return _dafny.UnicodeSeqOfUtf8Bytes("")
	}
}
func (_static *CompanionStruct_Default___) SelectOption(options m_AhpSkeleton.Option, id m_AhpSkeleton.Option) m_AhpSkeleton.Option {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (((options).Is_None()) || ((id).Is_None())) || ((_dafny.IntOfUint32(((options).Dtor_value().(_dafny.Sequence)).Cardinality())).Sign() == 0) {
		return m_AhpSkeleton.Companion_Option_.Create_None_()
	} else if _dafny.Companion_Sequence_.Equal(Companion_Default___.OptionId(((options).Dtor_value().(_dafny.Sequence)).Select(0).(m_ConfluxCodec.Json)), (id).Dtor_value().(_dafny.Sequence)) {
		return Companion_Default___.JNoNull(((options).Dtor_value().(_dafny.Sequence)).Select(0).(m_ConfluxCodec.Json))
	} else {
		var _in0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(((options).Dtor_value().(_dafny.Sequence)).Drop(1))
		_ = _in0
		var _in1 m_AhpSkeleton.Option = id
		_ = _in1
		options = _in0
		id = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) DeltaOne(tc ToolCall, content _dafny.Sequence, inv m_AhpSkeleton.Option, meta m_AhpSkeleton.Option) ToolCall {
	if !_dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("streaming")) {
		return tc
	} else {
		var _0_dt__update__tmp_h0 ToolCall = tc
		_ = _0_dt__update__tmp_h0
		var _1_dt__update_hinvocationMessage_h0 _dafny.Sequence = Companion_Default___.OptOr(inv, (tc).Dtor_invocationMessage())
		_ = _1_dt__update_hinvocationMessage_h0
		var _2_dt__update_hpartialInput_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.OptOr((tc).Dtor_partialInput(), _dafny.UnicodeSeqOfUtf8Bytes("")), content))
		_ = _2_dt__update_hpartialInput_h0
		var _3_dt__update_hmeta_h0 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
		_ = _3_dt__update_hmeta_h0
		return Companion_ToolCall_.Create_ToolCall_((_0_dt__update__tmp_h0).Dtor_toolCallId(), (_0_dt__update__tmp_h0).Dtor_toolName(), (_0_dt__update__tmp_h0).Dtor_displayName(), (_0_dt__update__tmp_h0).Dtor_status(), (_0_dt__update__tmp_h0).Dtor_intention(), (_0_dt__update__tmp_h0).Dtor_contributor(), _3_dt__update_hmeta_h0, _1_dt__update_hinvocationMessage_h0, (_0_dt__update__tmp_h0).Dtor_toolInput(), (_0_dt__update__tmp_h0).Dtor_confirmationTitle(), (_0_dt__update__tmp_h0).Dtor_riskAssessment(), (_0_dt__update__tmp_h0).Dtor_edits(), (_0_dt__update__tmp_h0).Dtor_editable(), (_0_dt__update__tmp_h0).Dtor_options(), (_0_dt__update__tmp_h0).Dtor_confirmed(), (_0_dt__update__tmp_h0).Dtor_success(), (_0_dt__update__tmp_h0).Dtor_pastTenseMessage(), (_0_dt__update__tmp_h0).Dtor_reason(), (_0_dt__update__tmp_h0).Dtor_reasonMessage(), (_0_dt__update__tmp_h0).Dtor_userSuggestion(), (_0_dt__update__tmp_h0).Dtor_selectedOption(), (_0_dt__update__tmp_h0).Dtor_content(), (_0_dt__update__tmp_h0).Dtor_structuredContent(), (_0_dt__update__tmp_h0).Dtor_error(), (_0_dt__update__tmp_h0).Dtor_auth(), _2_dt__update_hpartialInput_h0)
	}
}
func (_static *CompanionStruct_Default___) ReadyOne(tc ToolCall, inv m_AhpSkeleton.Option, input m_AhpSkeleton.Option, title m_AhpSkeleton.Option, risk m_AhpSkeleton.Option, edits m_AhpSkeleton.Option, editable m_AhpSkeleton.Option, options m_AhpSkeleton.Option, confirmed m_AhpSkeleton.Option, meta m_AhpSkeleton.Option) ToolCall {
	if ((!_dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("streaming"))) && (!_dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("running")))) && (!_dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("pending-confirmation"))) {
		return tc
	} else if (confirmed).Is_Some() {
		var _0_dt__update__tmp_h0 ToolCall = tc
		_ = _0_dt__update__tmp_h0
		var _1_dt__update_hpartialInput_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _1_dt__update_hpartialInput_h0
		var _2_dt__update_hauth_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _2_dt__update_hauth_h0
		var _3_dt__update_herror_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _3_dt__update_herror_h0
		var _4_dt__update_hstructuredContent_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _4_dt__update_hstructuredContent_h0
		var _5_dt__update_hcontent_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _5_dt__update_hcontent_h0
		var _6_dt__update_hselectedOption_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _6_dt__update_hselectedOption_h0
		var _7_dt__update_huserSuggestion_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _7_dt__update_huserSuggestion_h0
		var _8_dt__update_hreasonMessage_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _8_dt__update_hreasonMessage_h0
		var _9_dt__update_hreason_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _9_dt__update_hreason_h0
		var _10_dt__update_hpastTenseMessage_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _10_dt__update_hpastTenseMessage_h0
		var _11_dt__update_hsuccess_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _11_dt__update_hsuccess_h0
		var _12_dt__update_hconfirmed_h0 m_AhpSkeleton.Option = confirmed
		_ = _12_dt__update_hconfirmed_h0
		var _13_dt__update_hoptions_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _13_dt__update_hoptions_h0
		var _14_dt__update_heditable_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _14_dt__update_heditable_h0
		var _15_dt__update_hedits_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _15_dt__update_hedits_h0
		var _16_dt__update_hriskAssessment_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _16_dt__update_hriskAssessment_h0
		var _17_dt__update_hconfirmationTitle_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _17_dt__update_hconfirmationTitle_h0
		var _18_dt__update_htoolInput_h0 m_AhpSkeleton.Option = input
		_ = _18_dt__update_htoolInput_h0
		var _19_dt__update_hinvocationMessage_h0 _dafny.Sequence = Companion_Default___.OptOr(inv, _dafny.UnicodeSeqOfUtf8Bytes(""))
		_ = _19_dt__update_hinvocationMessage_h0
		var _20_dt__update_hmeta_h0 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
		_ = _20_dt__update_hmeta_h0
		var _21_dt__update_hstatus_h0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("running")
		_ = _21_dt__update_hstatus_h0
		return Companion_ToolCall_.Create_ToolCall_((_0_dt__update__tmp_h0).Dtor_toolCallId(), (_0_dt__update__tmp_h0).Dtor_toolName(), (_0_dt__update__tmp_h0).Dtor_displayName(), _21_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).Dtor_intention(), (_0_dt__update__tmp_h0).Dtor_contributor(), _20_dt__update_hmeta_h0, _19_dt__update_hinvocationMessage_h0, _18_dt__update_htoolInput_h0, _17_dt__update_hconfirmationTitle_h0, _16_dt__update_hriskAssessment_h0, _15_dt__update_hedits_h0, _14_dt__update_heditable_h0, _13_dt__update_hoptions_h0, _12_dt__update_hconfirmed_h0, _11_dt__update_hsuccess_h0, _10_dt__update_hpastTenseMessage_h0, _9_dt__update_hreason_h0, _8_dt__update_hreasonMessage_h0, _7_dt__update_huserSuggestion_h0, _6_dt__update_hselectedOption_h0, _5_dt__update_hcontent_h0, _4_dt__update_hstructuredContent_h0, _3_dt__update_herror_h0, _2_dt__update_hauth_h0, _1_dt__update_hpartialInput_h0)
	} else {
		var _22_pending bool = _dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("pending-confirmation"))
		_ = _22_pending
		var _23_dt__update__tmp_h1 ToolCall = tc
		_ = _23_dt__update__tmp_h1
		var _24_dt__update_hpartialInput_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _24_dt__update_hpartialInput_h1
		var _25_dt__update_hauth_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _25_dt__update_hauth_h1
		var _26_dt__update_herror_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _26_dt__update_herror_h1
		var _27_dt__update_hstructuredContent_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _27_dt__update_hstructuredContent_h1
		var _28_dt__update_hcontent_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _28_dt__update_hcontent_h1
		var _29_dt__update_hselectedOption_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _29_dt__update_hselectedOption_h1
		var _30_dt__update_huserSuggestion_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _30_dt__update_huserSuggestion_h1
		var _31_dt__update_hreasonMessage_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _31_dt__update_hreasonMessage_h1
		var _32_dt__update_hreason_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _32_dt__update_hreason_h1
		var _33_dt__update_hpastTenseMessage_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _33_dt__update_hpastTenseMessage_h1
		var _34_dt__update_hsuccess_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _34_dt__update_hsuccess_h1
		var _35_dt__update_hconfirmed_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _35_dt__update_hconfirmed_h1
		var _36_dt__update_hoptions_h1 m_AhpSkeleton.Option = (func() m_AhpSkeleton.Option {
			if ((options).Is_Some()) || (!(_22_pending)) {
				return options
			}
			return (tc).Dtor_options()
		})()
		_ = _36_dt__update_hoptions_h1
		var _37_dt__update_heditable_h1 m_AhpSkeleton.Option = (func() m_AhpSkeleton.Option {
			if ((editable).Is_Some()) || (!(_22_pending)) {
				return editable
			}
			return (tc).Dtor_editable()
		})()
		_ = _37_dt__update_heditable_h1
		var _38_dt__update_hedits_h1 m_AhpSkeleton.Option = (func() m_AhpSkeleton.Option {
			if ((edits).Is_Some()) || (!(_22_pending)) {
				return edits
			}
			return (tc).Dtor_edits()
		})()
		_ = _38_dt__update_hedits_h1
		var _39_dt__update_hriskAssessment_h1 m_AhpSkeleton.Option = (func() m_AhpSkeleton.Option {
			if ((risk).Is_Some()) || (!(_22_pending)) {
				return risk
			}
			return (tc).Dtor_riskAssessment()
		})()
		_ = _39_dt__update_hriskAssessment_h1
		var _40_dt__update_hconfirmationTitle_h1 m_AhpSkeleton.Option = (func() m_AhpSkeleton.Option {
			if ((title).Is_Some()) || (!(_22_pending)) {
				return title
			}
			return (tc).Dtor_confirmationTitle()
		})()
		_ = _40_dt__update_hconfirmationTitle_h1
		var _41_dt__update_htoolInput_h1 m_AhpSkeleton.Option = (func() m_AhpSkeleton.Option {
			if ((input).Is_Some()) || (!(_22_pending)) {
				return input
			}
			return (tc).Dtor_toolInput()
		})()
		_ = _41_dt__update_htoolInput_h1
		var _42_dt__update_hinvocationMessage_h1 _dafny.Sequence = Companion_Default___.OptOr(inv, _dafny.UnicodeSeqOfUtf8Bytes(""))
		_ = _42_dt__update_hinvocationMessage_h1
		var _43_dt__update_hmeta_h1 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
		_ = _43_dt__update_hmeta_h1
		var _44_dt__update_hstatus_h1 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("pending-confirmation")
		_ = _44_dt__update_hstatus_h1
		return Companion_ToolCall_.Create_ToolCall_((_23_dt__update__tmp_h1).Dtor_toolCallId(), (_23_dt__update__tmp_h1).Dtor_toolName(), (_23_dt__update__tmp_h1).Dtor_displayName(), _44_dt__update_hstatus_h1, (_23_dt__update__tmp_h1).Dtor_intention(), (_23_dt__update__tmp_h1).Dtor_contributor(), _43_dt__update_hmeta_h1, _42_dt__update_hinvocationMessage_h1, _41_dt__update_htoolInput_h1, _40_dt__update_hconfirmationTitle_h1, _39_dt__update_hriskAssessment_h1, _38_dt__update_hedits_h1, _37_dt__update_heditable_h1, _36_dt__update_hoptions_h1, _35_dt__update_hconfirmed_h1, _34_dt__update_hsuccess_h1, _33_dt__update_hpastTenseMessage_h1, _32_dt__update_hreason_h1, _31_dt__update_hreasonMessage_h1, _30_dt__update_huserSuggestion_h1, _29_dt__update_hselectedOption_h1, _28_dt__update_hcontent_h1, _27_dt__update_hstructuredContent_h1, _26_dt__update_herror_h1, _25_dt__update_hauth_h1, _24_dt__update_hpartialInput_h1)
	}
}
func (_static *CompanionStruct_Default___) AuthRequiredOne(tc ToolCall, challenge m_ConfluxCodec.Json, meta m_AhpSkeleton.Option) ToolCall {
	if ((!_dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("running"))) || (((tc).Dtor_contributor()).Is_None())) || (!_dafny.Companion_Sequence_.Equal((((tc).Dtor_contributor()).Dtor_value().(ToolCallContributor)).Dtor_kind(), _dafny.UnicodeSeqOfUtf8Bytes("mcp"))) {
		return tc
	} else {
		var _0_dt__update__tmp_h0 ToolCall = tc
		_ = _0_dt__update__tmp_h0
		var _1_dt__update_hpartialInput_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _1_dt__update_hpartialInput_h0
		var _2_dt__update_hauth_h0 m_AhpSkeleton.Option = Companion_Default___.JNoNull(challenge)
		_ = _2_dt__update_hauth_h0
		var _3_dt__update_herror_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _3_dt__update_herror_h0
		var _4_dt__update_hstructuredContent_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _4_dt__update_hstructuredContent_h0
		var _5_dt__update_huserSuggestion_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _5_dt__update_huserSuggestion_h0
		var _6_dt__update_hreasonMessage_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _6_dt__update_hreasonMessage_h0
		var _7_dt__update_hreason_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _7_dt__update_hreason_h0
		var _8_dt__update_hpastTenseMessage_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _8_dt__update_hpastTenseMessage_h0
		var _9_dt__update_hsuccess_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _9_dt__update_hsuccess_h0
		var _10_dt__update_hoptions_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _10_dt__update_hoptions_h0
		var _11_dt__update_heditable_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _11_dt__update_heditable_h0
		var _12_dt__update_hedits_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _12_dt__update_hedits_h0
		var _13_dt__update_hriskAssessment_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _13_dt__update_hriskAssessment_h0
		var _14_dt__update_hconfirmationTitle_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _14_dt__update_hconfirmationTitle_h0
		var _15_dt__update_hmeta_h0 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
		_ = _15_dt__update_hmeta_h0
		var _16_dt__update_hstatus_h0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("auth-required")
		_ = _16_dt__update_hstatus_h0
		return Companion_ToolCall_.Create_ToolCall_((_0_dt__update__tmp_h0).Dtor_toolCallId(), (_0_dt__update__tmp_h0).Dtor_toolName(), (_0_dt__update__tmp_h0).Dtor_displayName(), _16_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).Dtor_intention(), (_0_dt__update__tmp_h0).Dtor_contributor(), _15_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).Dtor_invocationMessage(), (_0_dt__update__tmp_h0).Dtor_toolInput(), _14_dt__update_hconfirmationTitle_h0, _13_dt__update_hriskAssessment_h0, _12_dt__update_hedits_h0, _11_dt__update_heditable_h0, _10_dt__update_hoptions_h0, (_0_dt__update__tmp_h0).Dtor_confirmed(), _9_dt__update_hsuccess_h0, _8_dt__update_hpastTenseMessage_h0, _7_dt__update_hreason_h0, _6_dt__update_hreasonMessage_h0, _5_dt__update_huserSuggestion_h0, (_0_dt__update__tmp_h0).Dtor_selectedOption(), (_0_dt__update__tmp_h0).Dtor_content(), _4_dt__update_hstructuredContent_h0, _3_dt__update_herror_h0, _2_dt__update_hauth_h0, _1_dt__update_hpartialInput_h0)
	}
}
func (_static *CompanionStruct_Default___) AuthResolvedOne(tc ToolCall, meta m_AhpSkeleton.Option) ToolCall {
	if !_dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("auth-required")) {
		return tc
	} else {
		var _0_dt__update__tmp_h0 ToolCall = tc
		_ = _0_dt__update__tmp_h0
		var _1_dt__update_hauth_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _1_dt__update_hauth_h0
		var _2_dt__update_hmeta_h0 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
		_ = _2_dt__update_hmeta_h0
		var _3_dt__update_hstatus_h0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("running")
		_ = _3_dt__update_hstatus_h0
		return Companion_ToolCall_.Create_ToolCall_((_0_dt__update__tmp_h0).Dtor_toolCallId(), (_0_dt__update__tmp_h0).Dtor_toolName(), (_0_dt__update__tmp_h0).Dtor_displayName(), _3_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).Dtor_intention(), (_0_dt__update__tmp_h0).Dtor_contributor(), _2_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).Dtor_invocationMessage(), (_0_dt__update__tmp_h0).Dtor_toolInput(), (_0_dt__update__tmp_h0).Dtor_confirmationTitle(), (_0_dt__update__tmp_h0).Dtor_riskAssessment(), (_0_dt__update__tmp_h0).Dtor_edits(), (_0_dt__update__tmp_h0).Dtor_editable(), (_0_dt__update__tmp_h0).Dtor_options(), (_0_dt__update__tmp_h0).Dtor_confirmed(), (_0_dt__update__tmp_h0).Dtor_success(), (_0_dt__update__tmp_h0).Dtor_pastTenseMessage(), (_0_dt__update__tmp_h0).Dtor_reason(), (_0_dt__update__tmp_h0).Dtor_reasonMessage(), (_0_dt__update__tmp_h0).Dtor_userSuggestion(), (_0_dt__update__tmp_h0).Dtor_selectedOption(), (_0_dt__update__tmp_h0).Dtor_content(), (_0_dt__update__tmp_h0).Dtor_structuredContent(), (_0_dt__update__tmp_h0).Dtor_error(), _1_dt__update_hauth_h0, (_0_dt__update__tmp_h0).Dtor_partialInput())
	}
}
func (_static *CompanionStruct_Default___) ConfirmOne(tc ToolCall, approved bool, confirmedVal m_AhpSkeleton.Option, reason m_AhpSkeleton.Option, reasonMessage m_AhpSkeleton.Option, userSuggestion m_AhpSkeleton.Option, editedInput m_AhpSkeleton.Option, selectedOptionId m_AhpSkeleton.Option, meta m_AhpSkeleton.Option) ToolCall {
	if !_dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("pending-confirmation")) {
		return tc
	} else if approved {
		var _0_dt__update__tmp_h0 ToolCall = tc
		_ = _0_dt__update__tmp_h0
		var _1_dt__update_hauth_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _1_dt__update_hauth_h0
		var _2_dt__update_herror_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _2_dt__update_herror_h0
		var _3_dt__update_hstructuredContent_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _3_dt__update_hstructuredContent_h0
		var _4_dt__update_hcontent_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _4_dt__update_hcontent_h0
		var _5_dt__update_huserSuggestion_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _5_dt__update_huserSuggestion_h0
		var _6_dt__update_hreasonMessage_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _6_dt__update_hreasonMessage_h0
		var _7_dt__update_hreason_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _7_dt__update_hreason_h0
		var _8_dt__update_hpastTenseMessage_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _8_dt__update_hpastTenseMessage_h0
		var _9_dt__update_hsuccess_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _9_dt__update_hsuccess_h0
		var _10_dt__update_hpartialInput_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _10_dt__update_hpartialInput_h0
		var _11_dt__update_hselectedOption_h0 m_AhpSkeleton.Option = Companion_Default___.SelectOption((tc).Dtor_options(), selectedOptionId)
		_ = _11_dt__update_hselectedOption_h0
		var _12_dt__update_hconfirmed_h0 m_AhpSkeleton.Option = confirmedVal
		_ = _12_dt__update_hconfirmed_h0
		var _13_dt__update_hoptions_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _13_dt__update_hoptions_h0
		var _14_dt__update_heditable_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _14_dt__update_heditable_h0
		var _15_dt__update_hedits_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _15_dt__update_hedits_h0
		var _16_dt__update_hriskAssessment_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _16_dt__update_hriskAssessment_h0
		var _17_dt__update_hconfirmationTitle_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _17_dt__update_hconfirmationTitle_h0
		var _18_dt__update_htoolInput_h0 m_AhpSkeleton.Option = (func() m_AhpSkeleton.Option {
			if (editedInput).Is_Some() {
				return editedInput
			}
			return (tc).Dtor_toolInput()
		})()
		_ = _18_dt__update_htoolInput_h0
		var _19_dt__update_hmeta_h0 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
		_ = _19_dt__update_hmeta_h0
		var _20_dt__update_hstatus_h0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("running")
		_ = _20_dt__update_hstatus_h0
		return Companion_ToolCall_.Create_ToolCall_((_0_dt__update__tmp_h0).Dtor_toolCallId(), (_0_dt__update__tmp_h0).Dtor_toolName(), (_0_dt__update__tmp_h0).Dtor_displayName(), _20_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).Dtor_intention(), (_0_dt__update__tmp_h0).Dtor_contributor(), _19_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).Dtor_invocationMessage(), _18_dt__update_htoolInput_h0, _17_dt__update_hconfirmationTitle_h0, _16_dt__update_hriskAssessment_h0, _15_dt__update_hedits_h0, _14_dt__update_heditable_h0, _13_dt__update_hoptions_h0, _12_dt__update_hconfirmed_h0, _9_dt__update_hsuccess_h0, _8_dt__update_hpastTenseMessage_h0, _7_dt__update_hreason_h0, _6_dt__update_hreasonMessage_h0, _5_dt__update_huserSuggestion_h0, _11_dt__update_hselectedOption_h0, _4_dt__update_hcontent_h0, _3_dt__update_hstructuredContent_h0, _2_dt__update_herror_h0, _1_dt__update_hauth_h0, _10_dt__update_hpartialInput_h0)
	} else {
		var _21_dt__update__tmp_h1 ToolCall = tc
		_ = _21_dt__update__tmp_h1
		var _22_dt__update_hpartialInput_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _22_dt__update_hpartialInput_h1
		var _23_dt__update_hauth_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _23_dt__update_hauth_h1
		var _24_dt__update_herror_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _24_dt__update_herror_h1
		var _25_dt__update_hstructuredContent_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _25_dt__update_hstructuredContent_h1
		var _26_dt__update_hcontent_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _26_dt__update_hcontent_h1
		var _27_dt__update_hselectedOption_h1 m_AhpSkeleton.Option = Companion_Default___.SelectOption((tc).Dtor_options(), selectedOptionId)
		_ = _27_dt__update_hselectedOption_h1
		var _28_dt__update_huserSuggestion_h1 m_AhpSkeleton.Option = userSuggestion
		_ = _28_dt__update_huserSuggestion_h1
		var _29_dt__update_hreasonMessage_h1 m_AhpSkeleton.Option = reasonMessage
		_ = _29_dt__update_hreasonMessage_h1
		var _30_dt__update_hreason_h1 m_AhpSkeleton.Option = reason
		_ = _30_dt__update_hreason_h1
		var _31_dt__update_hpastTenseMessage_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _31_dt__update_hpastTenseMessage_h1
		var _32_dt__update_hsuccess_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _32_dt__update_hsuccess_h1
		var _33_dt__update_hconfirmed_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _33_dt__update_hconfirmed_h1
		var _34_dt__update_hoptions_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _34_dt__update_hoptions_h1
		var _35_dt__update_heditable_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _35_dt__update_heditable_h1
		var _36_dt__update_hedits_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _36_dt__update_hedits_h1
		var _37_dt__update_hriskAssessment_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _37_dt__update_hriskAssessment_h1
		var _38_dt__update_hconfirmationTitle_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _38_dt__update_hconfirmationTitle_h1
		var _39_dt__update_hmeta_h1 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
		_ = _39_dt__update_hmeta_h1
		var _40_dt__update_hstatus_h1 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("cancelled")
		_ = _40_dt__update_hstatus_h1
		return Companion_ToolCall_.Create_ToolCall_((_21_dt__update__tmp_h1).Dtor_toolCallId(), (_21_dt__update__tmp_h1).Dtor_toolName(), (_21_dt__update__tmp_h1).Dtor_displayName(), _40_dt__update_hstatus_h1, (_21_dt__update__tmp_h1).Dtor_intention(), (_21_dt__update__tmp_h1).Dtor_contributor(), _39_dt__update_hmeta_h1, (_21_dt__update__tmp_h1).Dtor_invocationMessage(), (_21_dt__update__tmp_h1).Dtor_toolInput(), _38_dt__update_hconfirmationTitle_h1, _37_dt__update_hriskAssessment_h1, _36_dt__update_hedits_h1, _35_dt__update_heditable_h1, _34_dt__update_hoptions_h1, _33_dt__update_hconfirmed_h1, _32_dt__update_hsuccess_h1, _31_dt__update_hpastTenseMessage_h1, _30_dt__update_hreason_h1, _29_dt__update_hreasonMessage_h1, _28_dt__update_huserSuggestion_h1, _27_dt__update_hselectedOption_h1, _26_dt__update_hcontent_h1, _25_dt__update_hstructuredContent_h1, _24_dt__update_herror_h1, _23_dt__update_hauth_h1, _22_dt__update_hpartialInput_h1)
	}
}
func (_static *CompanionStruct_Default___) CompleteOne(tc ToolCall, ok bool, past m_AhpSkeleton.Option, resultContent m_AhpSkeleton.Option, structured m_AhpSkeleton.Option, err m_AhpSkeleton.Option, requiresResultConfirmation bool, meta m_AhpSkeleton.Option) ToolCall {
	if _dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("auth-required")) {
		if ok {
			return tc
		} else {
			var _0_dt__update__tmp_h0 ToolCall = tc
			_ = _0_dt__update__tmp_h0
			var _1_dt__update_hpartialInput_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
			_ = _1_dt__update_hpartialInput_h0
			var _2_dt__update_hauth_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
			_ = _2_dt__update_hauth_h0
			var _3_dt__update_herror_h0 m_AhpSkeleton.Option = err
			_ = _3_dt__update_herror_h0
			var _4_dt__update_hstructuredContent_h0 m_AhpSkeleton.Option = structured
			_ = _4_dt__update_hstructuredContent_h0
			var _5_dt__update_hcontent_h0 m_AhpSkeleton.Option = (func() m_AhpSkeleton.Option {
				if (resultContent).Is_Some() {
					return resultContent
				}
				return (tc).Dtor_content()
			})()
			_ = _5_dt__update_hcontent_h0
			var _6_dt__update_huserSuggestion_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
			_ = _6_dt__update_huserSuggestion_h0
			var _7_dt__update_hreasonMessage_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
			_ = _7_dt__update_hreasonMessage_h0
			var _8_dt__update_hreason_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
			_ = _8_dt__update_hreason_h0
			var _9_dt__update_hpastTenseMessage_h0 m_AhpSkeleton.Option = past
			_ = _9_dt__update_hpastTenseMessage_h0
			var _10_dt__update_hsuccess_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(false)
			_ = _10_dt__update_hsuccess_h0
			var _11_dt__update_hoptions_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
			_ = _11_dt__update_hoptions_h0
			var _12_dt__update_heditable_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
			_ = _12_dt__update_heditable_h0
			var _13_dt__update_hedits_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
			_ = _13_dt__update_hedits_h0
			var _14_dt__update_hriskAssessment_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
			_ = _14_dt__update_hriskAssessment_h0
			var _15_dt__update_hconfirmationTitle_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
			_ = _15_dt__update_hconfirmationTitle_h0
			var _16_dt__update_hmeta_h0 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
			_ = _16_dt__update_hmeta_h0
			var _17_dt__update_hstatus_h0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("completed")
			_ = _17_dt__update_hstatus_h0
			return Companion_ToolCall_.Create_ToolCall_((_0_dt__update__tmp_h0).Dtor_toolCallId(), (_0_dt__update__tmp_h0).Dtor_toolName(), (_0_dt__update__tmp_h0).Dtor_displayName(), _17_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).Dtor_intention(), (_0_dt__update__tmp_h0).Dtor_contributor(), _16_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).Dtor_invocationMessage(), (_0_dt__update__tmp_h0).Dtor_toolInput(), _15_dt__update_hconfirmationTitle_h0, _14_dt__update_hriskAssessment_h0, _13_dt__update_hedits_h0, _12_dt__update_heditable_h0, _11_dt__update_hoptions_h0, (_0_dt__update__tmp_h0).Dtor_confirmed(), _10_dt__update_hsuccess_h0, _9_dt__update_hpastTenseMessage_h0, _8_dt__update_hreason_h0, _7_dt__update_hreasonMessage_h0, _6_dt__update_huserSuggestion_h0, (_0_dt__update__tmp_h0).Dtor_selectedOption(), _5_dt__update_hcontent_h0, _4_dt__update_hstructuredContent_h0, _3_dt__update_herror_h0, _2_dt__update_hauth_h0, _1_dt__update_hpartialInput_h0)
		}
	} else if (!_dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("running"))) && (!_dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("pending-confirmation"))) {
		return tc
	} else {
		var _18_conf m_AhpSkeleton.Option = (func() m_AhpSkeleton.Option {
			if _dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("running")) {
				return (tc).Dtor_confirmed()
			}
			return m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("not-needed"))
		})()
		_ = _18_conf
		var _19_selected m_AhpSkeleton.Option = (func() m_AhpSkeleton.Option {
			if _dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("running")) {
				return (tc).Dtor_selectedOption()
			}
			return m_AhpSkeleton.Companion_Option_.Create_None_()
		})()
		_ = _19_selected
		var _20_newStatus _dafny.Sequence = (func() _dafny.Sequence {
			if requiresResultConfirmation {
				return _dafny.UnicodeSeqOfUtf8Bytes("pending-result-confirmation")
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("completed")
		})()
		_ = _20_newStatus
		var _21_dt__update__tmp_h1 ToolCall = tc
		_ = _21_dt__update__tmp_h1
		var _22_dt__update_hpartialInput_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _22_dt__update_hpartialInput_h1
		var _23_dt__update_hauth_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _23_dt__update_hauth_h1
		var _24_dt__update_herror_h1 m_AhpSkeleton.Option = err
		_ = _24_dt__update_herror_h1
		var _25_dt__update_hstructuredContent_h1 m_AhpSkeleton.Option = structured
		_ = _25_dt__update_hstructuredContent_h1
		var _26_dt__update_hcontent_h1 m_AhpSkeleton.Option = resultContent
		_ = _26_dt__update_hcontent_h1
		var _27_dt__update_huserSuggestion_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _27_dt__update_huserSuggestion_h1
		var _28_dt__update_hreasonMessage_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _28_dt__update_hreasonMessage_h1
		var _29_dt__update_hreason_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _29_dt__update_hreason_h1
		var _30_dt__update_hpastTenseMessage_h1 m_AhpSkeleton.Option = past
		_ = _30_dt__update_hpastTenseMessage_h1
		var _31_dt__update_hsuccess_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(ok)
		_ = _31_dt__update_hsuccess_h1
		var _32_dt__update_hselectedOption_h0 m_AhpSkeleton.Option = _19_selected
		_ = _32_dt__update_hselectedOption_h0
		var _33_dt__update_hconfirmed_h0 m_AhpSkeleton.Option = _18_conf
		_ = _33_dt__update_hconfirmed_h0
		var _34_dt__update_hoptions_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _34_dt__update_hoptions_h1
		var _35_dt__update_heditable_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _35_dt__update_heditable_h1
		var _36_dt__update_hedits_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _36_dt__update_hedits_h1
		var _37_dt__update_hriskAssessment_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _37_dt__update_hriskAssessment_h1
		var _38_dt__update_hconfirmationTitle_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _38_dt__update_hconfirmationTitle_h1
		var _39_dt__update_hmeta_h1 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
		_ = _39_dt__update_hmeta_h1
		var _40_dt__update_hstatus_h1 _dafny.Sequence = _20_newStatus
		_ = _40_dt__update_hstatus_h1
		return Companion_ToolCall_.Create_ToolCall_((_21_dt__update__tmp_h1).Dtor_toolCallId(), (_21_dt__update__tmp_h1).Dtor_toolName(), (_21_dt__update__tmp_h1).Dtor_displayName(), _40_dt__update_hstatus_h1, (_21_dt__update__tmp_h1).Dtor_intention(), (_21_dt__update__tmp_h1).Dtor_contributor(), _39_dt__update_hmeta_h1, (_21_dt__update__tmp_h1).Dtor_invocationMessage(), (_21_dt__update__tmp_h1).Dtor_toolInput(), _38_dt__update_hconfirmationTitle_h1, _37_dt__update_hriskAssessment_h1, _36_dt__update_hedits_h1, _35_dt__update_heditable_h1, _34_dt__update_hoptions_h1, _33_dt__update_hconfirmed_h0, _31_dt__update_hsuccess_h1, _30_dt__update_hpastTenseMessage_h1, _29_dt__update_hreason_h1, _28_dt__update_hreasonMessage_h1, _27_dt__update_huserSuggestion_h1, _32_dt__update_hselectedOption_h0, _26_dt__update_hcontent_h1, _25_dt__update_hstructuredContent_h1, _24_dt__update_herror_h1, _23_dt__update_hauth_h1, _22_dt__update_hpartialInput_h1)
	}
}
func (_static *CompanionStruct_Default___) ResultConfirmOne(tc ToolCall, approved bool, meta m_AhpSkeleton.Option) ToolCall {
	if !_dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("pending-result-confirmation")) {
		return tc
	} else if approved {
		var _0_dt__update__tmp_h0 ToolCall = tc
		_ = _0_dt__update__tmp_h0
		var _1_dt__update_hmeta_h0 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
		_ = _1_dt__update_hmeta_h0
		var _2_dt__update_hstatus_h0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("completed")
		_ = _2_dt__update_hstatus_h0
		return Companion_ToolCall_.Create_ToolCall_((_0_dt__update__tmp_h0).Dtor_toolCallId(), (_0_dt__update__tmp_h0).Dtor_toolName(), (_0_dt__update__tmp_h0).Dtor_displayName(), _2_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).Dtor_intention(), (_0_dt__update__tmp_h0).Dtor_contributor(), _1_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).Dtor_invocationMessage(), (_0_dt__update__tmp_h0).Dtor_toolInput(), (_0_dt__update__tmp_h0).Dtor_confirmationTitle(), (_0_dt__update__tmp_h0).Dtor_riskAssessment(), (_0_dt__update__tmp_h0).Dtor_edits(), (_0_dt__update__tmp_h0).Dtor_editable(), (_0_dt__update__tmp_h0).Dtor_options(), (_0_dt__update__tmp_h0).Dtor_confirmed(), (_0_dt__update__tmp_h0).Dtor_success(), (_0_dt__update__tmp_h0).Dtor_pastTenseMessage(), (_0_dt__update__tmp_h0).Dtor_reason(), (_0_dt__update__tmp_h0).Dtor_reasonMessage(), (_0_dt__update__tmp_h0).Dtor_userSuggestion(), (_0_dt__update__tmp_h0).Dtor_selectedOption(), (_0_dt__update__tmp_h0).Dtor_content(), (_0_dt__update__tmp_h0).Dtor_structuredContent(), (_0_dt__update__tmp_h0).Dtor_error(), (_0_dt__update__tmp_h0).Dtor_auth(), (_0_dt__update__tmp_h0).Dtor_partialInput())
	} else {
		var _3_dt__update__tmp_h1 ToolCall = tc
		_ = _3_dt__update__tmp_h1
		var _4_dt__update_herror_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _4_dt__update_herror_h0
		var _5_dt__update_hstructuredContent_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _5_dt__update_hstructuredContent_h0
		var _6_dt__update_hcontent_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _6_dt__update_hcontent_h0
		var _7_dt__update_hpastTenseMessage_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _7_dt__update_hpastTenseMessage_h0
		var _8_dt__update_hsuccess_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _8_dt__update_hsuccess_h0
		var _9_dt__update_hconfirmed_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_None_()
		_ = _9_dt__update_hconfirmed_h0
		var _10_dt__update_hreason_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("result-denied"))
		_ = _10_dt__update_hreason_h0
		var _11_dt__update_hmeta_h1 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
		_ = _11_dt__update_hmeta_h1
		var _12_dt__update_hstatus_h1 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("cancelled")
		_ = _12_dt__update_hstatus_h1
		return Companion_ToolCall_.Create_ToolCall_((_3_dt__update__tmp_h1).Dtor_toolCallId(), (_3_dt__update__tmp_h1).Dtor_toolName(), (_3_dt__update__tmp_h1).Dtor_displayName(), _12_dt__update_hstatus_h1, (_3_dt__update__tmp_h1).Dtor_intention(), (_3_dt__update__tmp_h1).Dtor_contributor(), _11_dt__update_hmeta_h1, (_3_dt__update__tmp_h1).Dtor_invocationMessage(), (_3_dt__update__tmp_h1).Dtor_toolInput(), (_3_dt__update__tmp_h1).Dtor_confirmationTitle(), (_3_dt__update__tmp_h1).Dtor_riskAssessment(), (_3_dt__update__tmp_h1).Dtor_edits(), (_3_dt__update__tmp_h1).Dtor_editable(), (_3_dt__update__tmp_h1).Dtor_options(), _9_dt__update_hconfirmed_h0, _8_dt__update_hsuccess_h0, _7_dt__update_hpastTenseMessage_h0, _10_dt__update_hreason_h0, (_3_dt__update__tmp_h1).Dtor_reasonMessage(), (_3_dt__update__tmp_h1).Dtor_userSuggestion(), (_3_dt__update__tmp_h1).Dtor_selectedOption(), _6_dt__update_hcontent_h0, _5_dt__update_hstructuredContent_h0, _4_dt__update_herror_h0, (_3_dt__update__tmp_h1).Dtor_auth(), (_3_dt__update__tmp_h1).Dtor_partialInput())
	}
}
func (_static *CompanionStruct_Default___) ContentChangedOne(tc ToolCall, c m_ConfluxCodec.Json, meta m_AhpSkeleton.Option) ToolCall {
	if !_dafny.Companion_Sequence_.Equal((tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("running")) {
		return tc
	} else {
		var _0_dt__update__tmp_h0 ToolCall = tc
		_ = _0_dt__update__tmp_h0
		var _1_dt__update_hcontent_h0 m_AhpSkeleton.Option = Companion_Default___.JNoNull(c)
		_ = _1_dt__update_hcontent_h0
		var _2_dt__update_hmeta_h0 m_AhpSkeleton.Option = Companion_Default___.MetaOr(meta, (tc).Dtor_meta())
		_ = _2_dt__update_hmeta_h0
		return Companion_ToolCall_.Create_ToolCall_((_0_dt__update__tmp_h0).Dtor_toolCallId(), (_0_dt__update__tmp_h0).Dtor_toolName(), (_0_dt__update__tmp_h0).Dtor_displayName(), (_0_dt__update__tmp_h0).Dtor_status(), (_0_dt__update__tmp_h0).Dtor_intention(), (_0_dt__update__tmp_h0).Dtor_contributor(), _2_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).Dtor_invocationMessage(), (_0_dt__update__tmp_h0).Dtor_toolInput(), (_0_dt__update__tmp_h0).Dtor_confirmationTitle(), (_0_dt__update__tmp_h0).Dtor_riskAssessment(), (_0_dt__update__tmp_h0).Dtor_edits(), (_0_dt__update__tmp_h0).Dtor_editable(), (_0_dt__update__tmp_h0).Dtor_options(), (_0_dt__update__tmp_h0).Dtor_confirmed(), (_0_dt__update__tmp_h0).Dtor_success(), (_0_dt__update__tmp_h0).Dtor_pastTenseMessage(), (_0_dt__update__tmp_h0).Dtor_reason(), (_0_dt__update__tmp_h0).Dtor_reasonMessage(), (_0_dt__update__tmp_h0).Dtor_userSuggestion(), (_0_dt__update__tmp_h0).Dtor_selectedOption(), _1_dt__update_hcontent_h0, (_0_dt__update__tmp_h0).Dtor_structuredContent(), (_0_dt__update__tmp_h0).Dtor_error(), (_0_dt__update__tmp_h0).Dtor_auth(), (_0_dt__update__tmp_h0).Dtor_partialInput())
	}
}
func (_static *CompanionStruct_Default___) ForceCancel(ps _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxOperators.Companion_Default___.Map(func(coer72 func(Part) Part) func(interface{}) interface{} {
		return func(arg80 interface{}) interface{} {
			return coer72(arg80.(Part))
		}
	}(func(_0_h Part) Part {
		return func() Part {
			var _source0 Part = _0_h
			_ = _source0
			{
				if _source0.Is_PToolCall() {
					var _1_tc ToolCall = _source0.Get_().(Part_PToolCall).Tc
					_ = _1_tc
					if (_dafny.Companion_Sequence_.Equal((_1_tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("completed"))) || (_dafny.Companion_Sequence_.Equal((_1_tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("cancelled"))) {
						return _0_h
					} else {
						return Companion_Part_.Create_PToolCall_(func(_pat_let237_0 ToolCall) ToolCall {
							return func(_2_dt__update__tmp_h0 ToolCall) ToolCall {
								return func(_pat_let238_0 m_AhpSkeleton.Option) ToolCall {
									return func(_3_dt__update_hpartialInput_h0 m_AhpSkeleton.Option) ToolCall {
										return func(_pat_let239_0 m_AhpSkeleton.Option) ToolCall {
											return func(_4_dt__update_hauth_h0 m_AhpSkeleton.Option) ToolCall {
												return func(_pat_let240_0 m_AhpSkeleton.Option) ToolCall {
													return func(_5_dt__update_herror_h0 m_AhpSkeleton.Option) ToolCall {
														return func(_pat_let241_0 m_AhpSkeleton.Option) ToolCall {
															return func(_6_dt__update_hstructuredContent_h0 m_AhpSkeleton.Option) ToolCall {
																return func(_pat_let242_0 m_AhpSkeleton.Option) ToolCall {
																	return func(_7_dt__update_hcontent_h0 m_AhpSkeleton.Option) ToolCall {
																		return func(_pat_let243_0 m_AhpSkeleton.Option) ToolCall {
																			return func(_8_dt__update_hselectedOption_h0 m_AhpSkeleton.Option) ToolCall {
																				return func(_pat_let244_0 m_AhpSkeleton.Option) ToolCall {
																					return func(_9_dt__update_huserSuggestion_h0 m_AhpSkeleton.Option) ToolCall {
																						return func(_pat_let245_0 m_AhpSkeleton.Option) ToolCall {
																							return func(_10_dt__update_hreasonMessage_h0 m_AhpSkeleton.Option) ToolCall {
																								return func(_pat_let246_0 m_AhpSkeleton.Option) ToolCall {
																									return func(_11_dt__update_hpastTenseMessage_h0 m_AhpSkeleton.Option) ToolCall {
																										return func(_pat_let247_0 m_AhpSkeleton.Option) ToolCall {
																											return func(_12_dt__update_hsuccess_h0 m_AhpSkeleton.Option) ToolCall {
																												return func(_pat_let248_0 m_AhpSkeleton.Option) ToolCall {
																													return func(_13_dt__update_hconfirmed_h0 m_AhpSkeleton.Option) ToolCall {
																														return func(_pat_let249_0 m_AhpSkeleton.Option) ToolCall {
																															return func(_14_dt__update_hoptions_h0 m_AhpSkeleton.Option) ToolCall {
																																return func(_pat_let250_0 m_AhpSkeleton.Option) ToolCall {
																																	return func(_15_dt__update_heditable_h0 m_AhpSkeleton.Option) ToolCall {
																																		return func(_pat_let251_0 m_AhpSkeleton.Option) ToolCall {
																																			return func(_16_dt__update_hedits_h0 m_AhpSkeleton.Option) ToolCall {
																																				return func(_pat_let252_0 m_AhpSkeleton.Option) ToolCall {
																																					return func(_17_dt__update_hriskAssessment_h0 m_AhpSkeleton.Option) ToolCall {
																																						return func(_pat_let253_0 m_AhpSkeleton.Option) ToolCall {
																																							return func(_18_dt__update_hconfirmationTitle_h0 m_AhpSkeleton.Option) ToolCall {
																																								return func(_pat_let254_0 m_AhpSkeleton.Option) ToolCall {
																																									return func(_19_dt__update_htoolInput_h0 m_AhpSkeleton.Option) ToolCall {
																																										return func(_pat_let255_0 m_AhpSkeleton.Option) ToolCall {
																																											return func(_20_dt__update_hreason_h0 m_AhpSkeleton.Option) ToolCall {
																																												return func(_pat_let256_0 _dafny.Sequence) ToolCall {
																																													return func(_21_dt__update_hstatus_h0 _dafny.Sequence) ToolCall {
																																														return Companion_ToolCall_.Create_ToolCall_((_2_dt__update__tmp_h0).Dtor_toolCallId(), (_2_dt__update__tmp_h0).Dtor_toolName(), (_2_dt__update__tmp_h0).Dtor_displayName(), _21_dt__update_hstatus_h0, (_2_dt__update__tmp_h0).Dtor_intention(), (_2_dt__update__tmp_h0).Dtor_contributor(), (_2_dt__update__tmp_h0).Dtor_meta(), (_2_dt__update__tmp_h0).Dtor_invocationMessage(), _19_dt__update_htoolInput_h0, _18_dt__update_hconfirmationTitle_h0, _17_dt__update_hriskAssessment_h0, _16_dt__update_hedits_h0, _15_dt__update_heditable_h0, _14_dt__update_hoptions_h0, _13_dt__update_hconfirmed_h0, _12_dt__update_hsuccess_h0, _11_dt__update_hpastTenseMessage_h0, _20_dt__update_hreason_h0, _10_dt__update_hreasonMessage_h0, _9_dt__update_huserSuggestion_h0, _8_dt__update_hselectedOption_h0, _7_dt__update_hcontent_h0, _6_dt__update_hstructuredContent_h0, _5_dt__update_herror_h0, _4_dt__update_hauth_h0, _3_dt__update_hpartialInput_h0)
																																													}(_pat_let256_0)
																																												}(_dafny.UnicodeSeqOfUtf8Bytes("cancelled"))
																																											}(_pat_let255_0)
																																										}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("skipped")))
																																									}(_pat_let254_0)
																																								}((func() m_AhpSkeleton.Option {
																																									if _dafny.Companion_Sequence_.Equal((_1_tc).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("streaming")) {
																																										return m_AhpSkeleton.Companion_Option_.Create_None_()
																																									}
																																									return (_1_tc).Dtor_toolInput()
																																								})())
																																							}(_pat_let253_0)
																																						}(m_AhpSkeleton.Companion_Option_.Create_None_())
																																					}(_pat_let252_0)
																																				}(m_AhpSkeleton.Companion_Option_.Create_None_())
																																			}(_pat_let251_0)
																																		}(m_AhpSkeleton.Companion_Option_.Create_None_())
																																	}(_pat_let250_0)
																																}(m_AhpSkeleton.Companion_Option_.Create_None_())
																															}(_pat_let249_0)
																														}(m_AhpSkeleton.Companion_Option_.Create_None_())
																													}(_pat_let248_0)
																												}(m_AhpSkeleton.Companion_Option_.Create_None_())
																											}(_pat_let247_0)
																										}(m_AhpSkeleton.Companion_Option_.Create_None_())
																									}(_pat_let246_0)
																								}(m_AhpSkeleton.Companion_Option_.Create_None_())
																							}(_pat_let245_0)
																						}(m_AhpSkeleton.Companion_Option_.Create_None_())
																					}(_pat_let244_0)
																				}(m_AhpSkeleton.Companion_Option_.Create_None_())
																			}(_pat_let243_0)
																		}(m_AhpSkeleton.Companion_Option_.Create_None_())
																	}(_pat_let242_0)
																}(m_AhpSkeleton.Companion_Option_.Create_None_())
															}(_pat_let241_0)
														}(m_AhpSkeleton.Companion_Option_.Create_None_())
													}(_pat_let240_0)
												}(m_AhpSkeleton.Companion_Option_.Create_None_())
											}(_pat_let239_0)
										}(m_AhpSkeleton.Companion_Option_.Create_None_())
									}(_pat_let238_0)
								}(m_AhpSkeleton.Companion_Option_.Create_None_())
							}(_pat_let237_0)
						}(_1_tc))
					}
				}
			}
			{
				return _0_h
			}
		}()
	}), ps)
}
func (_static *CompanionStruct_Default___) EndTurn(s ChatState, id _dafny.Sequence, turnState _dafny.Sequence, isError bool, err m_AhpSkeleton.Option) _dafny.Tuple {
	var _pat_let_tv0 = err
	_ = _pat_let_tv0
	var _pat_let_tv1 = turnState
	_ = _pat_let_tv1
	var _pat_let_tv2 = s
	_ = _pat_let_tv2
	var _pat_let_tv3 = isError
	_ = _pat_let_tv3
	if !(Companion_Default___.TurnMatches(s, id)) {
		return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
	} else {
		var _0_t Turn = ((s).Dtor_activeTurn()).Dtor_value().(Turn)
		_ = _0_t
		var _1_finalized Turn = func(_pat_let257_0 Turn) Turn {
			return func(_2_dt__update__tmp_h0 Turn) Turn {
				return func(_pat_let258_0 m_AhpSkeleton.Option) Turn {
					return func(_3_dt__update_herror_h0 m_AhpSkeleton.Option) Turn {
						return func(_pat_let259_0 _dafny.Sequence) Turn {
							return func(_4_dt__update_hparts_h0 _dafny.Sequence) Turn {
								return func(_pat_let260_0 _dafny.Sequence) Turn {
									return func(_5_dt__update_hstate_h0 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_2_dt__update__tmp_h0).Dtor_turnId(), (_2_dt__update__tmp_h0).Dtor_message(), _4_dt__update_hparts_h0, _5_dt__update_hstate_h0, (_2_dt__update__tmp_h0).Dtor_usage(), _3_dt__update_herror_h0)
									}(_pat_let260_0)
								}(_pat_let_tv1)
							}(_pat_let259_0)
						}(Companion_Default___.ForceCancel((_0_t).Dtor_parts()))
					}(_pat_let258_0)
				}(_pat_let_tv0)
			}(_pat_let257_0)
		}(_0_t)
		_ = _1_finalized
		var _6_next ChatState = func(_pat_let261_0 ChatState) ChatState {
			return func(_7_dt__update__tmp_h1 ChatState) ChatState {
				return func(_pat_let262_0 m_AhpSkeleton.Option) ChatState {
					return func(_8_dt__update_hactiveTurn_h0 m_AhpSkeleton.Option) ChatState {
						return func(_pat_let263_0 _dafny.Sequence) ChatState {
							return func(_9_dt__update_hturns_h0 _dafny.Sequence) ChatState {
								return Companion_ChatState_.Create_ChatState_(_9_dt__update_hturns_h0, (_7_dt__update__tmp_h1).Dtor_title(), (_7_dt__update__tmp_h1).Dtor_status(), (_7_dt__update__tmp_h1).Dtor_modifiedAt(), (_7_dt__update__tmp_h1).Dtor_draft(), (_7_dt__update__tmp_h1).Dtor_activity(), _8_dt__update_hactiveTurn_h0, (_7_dt__update__tmp_h1).Dtor_steeringMessage(), (_7_dt__update__tmp_h1).Dtor_queuedMessages(), (_7_dt__update__tmp_h1).Dtor_nextCursor())
							}(_pat_let263_0)
						}(_dafny.Companion_Sequence_.Concatenate((_pat_let_tv2).Dtor_turns(), _dafny.SeqOf(_1_finalized)))
					}(_pat_let262_0)
				}(m_AhpSkeleton.Companion_Option_.Create_None_())
			}(_pat_let261_0)
		}(s)
		_ = _6_next
		return _dafny.TupleOf(func(_pat_let264_0 ChatState) ChatState {
			return func(_10_dt__update__tmp_h2 ChatState) ChatState {
				return func(_pat_let265_0 uint32) ChatState {
					return func(_11_dt__update_hstatus_h0 uint32) ChatState {
						return Companion_ChatState_.Create_ChatState_((_10_dt__update__tmp_h2).Dtor_turns(), (_10_dt__update__tmp_h2).Dtor_title(), _11_dt__update_hstatus_h0, (_10_dt__update__tmp_h2).Dtor_modifiedAt(), (_10_dt__update__tmp_h2).Dtor_draft(), (_10_dt__update__tmp_h2).Dtor_activity(), (_10_dt__update__tmp_h2).Dtor_activeTurn(), (_10_dt__update__tmp_h2).Dtor_steeringMessage(), (_10_dt__update__tmp_h2).Dtor_queuedMessages(), (_10_dt__update__tmp_h2).Dtor_nextCursor())
					}(_pat_let265_0)
				}(Companion_Default___.SummaryStatus(_6_next, _pat_let_tv3))
			}(_pat_let264_0)
		}(_6_next), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
	}
}
func (_static *CompanionStruct_Default___) ApplyToChat(s ChatState, a ChatAction, now _dafny.Int) _dafny.Tuple {
	var _pat_let_tv0 = s
	_ = _pat_let_tv0
	var _pat_let_tv1 = s
	_ = _pat_let_tv1
	var _pat_let_tv2 = s
	_ = _pat_let_tv2
	var _pat_let_tv3 = s
	_ = _pat_let_tv3
	var _pat_let_tv4 = s
	_ = _pat_let_tv4
	var _pat_let_tv5 = s
	_ = _pat_let_tv5
	var _pat_let_tv6 = s
	_ = _pat_let_tv6
	var _pat_let_tv7 = s
	_ = _pat_let_tv7
	var _pat_let_tv8 = s
	_ = _pat_let_tv8
	var _pat_let_tv9 = s
	_ = _pat_let_tv9
	var _pat_let_tv10 = s
	_ = _pat_let_tv10
	var _pat_let_tv11 = s
	_ = _pat_let_tv11
	var _pat_let_tv12 = s
	_ = _pat_let_tv12
	var _pat_let_tv13 = s
	_ = _pat_let_tv13
	var _pat_let_tv14 = s
	_ = _pat_let_tv14
	var _pat_let_tv15 = s
	_ = _pat_let_tv15
	var _pat_let_tv16 = s
	_ = _pat_let_tv16
	var _pat_let_tv17 = s
	_ = _pat_let_tv17
	var _pat_let_tv18 = s
	_ = _pat_let_tv18
	var _pat_let_tv19 = s
	_ = _pat_let_tv19
	var _pat_let_tv20 = s
	_ = _pat_let_tv20
	var _pat_let_tv21 = s
	_ = _pat_let_tv21
	var _pat_let_tv22 = s
	_ = _pat_let_tv22
	var _pat_let_tv23 = s
	_ = _pat_let_tv23
	var _pat_let_tv24 = s
	_ = _pat_let_tv24
	var _pat_let_tv25 = s
	_ = _pat_let_tv25
	var _pat_let_tv26 = s
	_ = _pat_let_tv26
	var _pat_let_tv27 = s
	_ = _pat_let_tv27
	var _pat_let_tv28 = s
	_ = _pat_let_tv28
	var _source0 ChatAction = a
	_ = _source0
	{
		if _source0.Is_CDraftChanged() {
			var _0_d m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CDraftChanged).Draft
			_ = _0_d
			return _dafny.TupleOf(func(_pat_let266_0 ChatState) ChatState {
				return func(_1_dt__update__tmp_h0 ChatState) ChatState {
					return func(_pat_let267_0 m_AhpSkeleton.Option) ChatState {
						return func(_2_dt__update_hdraft_h0 m_AhpSkeleton.Option) ChatState {
							return Companion_ChatState_.Create_ChatState_((_1_dt__update__tmp_h0).Dtor_turns(), (_1_dt__update__tmp_h0).Dtor_title(), (_1_dt__update__tmp_h0).Dtor_status(), (_1_dt__update__tmp_h0).Dtor_modifiedAt(), _2_dt__update_hdraft_h0, (_1_dt__update__tmp_h0).Dtor_activity(), (_1_dt__update__tmp_h0).Dtor_activeTurn(), (_1_dt__update__tmp_h0).Dtor_steeringMessage(), (_1_dt__update__tmp_h0).Dtor_queuedMessages(), (_1_dt__update__tmp_h0).Dtor_nextCursor())
						}(_pat_let267_0)
					}(_0_d)
				}(_pat_let266_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_CActivityChanged() {
			var _3_ac m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CActivityChanged).Activity
			_ = _3_ac
			return _dafny.TupleOf(func(_pat_let268_0 ChatState) ChatState {
				return func(_4_dt__update__tmp_h1 ChatState) ChatState {
					return func(_pat_let269_0 m_AhpSkeleton.Option) ChatState {
						return func(_5_dt__update_hactivity_h0 m_AhpSkeleton.Option) ChatState {
							return Companion_ChatState_.Create_ChatState_((_4_dt__update__tmp_h1).Dtor_turns(), (_4_dt__update__tmp_h1).Dtor_title(), (_4_dt__update__tmp_h1).Dtor_status(), (_4_dt__update__tmp_h1).Dtor_modifiedAt(), (_4_dt__update__tmp_h1).Dtor_draft(), _5_dt__update_hactivity_h0, (_4_dt__update__tmp_h1).Dtor_activeTurn(), (_4_dt__update__tmp_h1).Dtor_steeringMessage(), (_4_dt__update__tmp_h1).Dtor_queuedMessages(), (_4_dt__update__tmp_h1).Dtor_nextCursor())
						}(_pat_let269_0)
					}(_3_ac)
				}(_pat_let268_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_CTurnStarted() {
			var _6_id _dafny.Sequence = _source0.Get_().(ChatAction_CTurnStarted).TurnId
			_ = _6_id
			var _7_m m_ConfluxCodec.Json = _source0.Get_().(ChatAction_CTurnStarted).Message
			_ = _7_m
			var _8_qid m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CTurnStarted).QueuedMessageId
			_ = _8_qid
			var _9_withTurn ChatState = func(_pat_let270_0 ChatState) ChatState {
				return func(_10_dt__update__tmp_h2 ChatState) ChatState {
					return func(_pat_let271_0 m_AhpSkeleton.Option) ChatState {
						return func(_11_dt__update_hactiveTurn_h0 m_AhpSkeleton.Option) ChatState {
							return Companion_ChatState_.Create_ChatState_((_10_dt__update__tmp_h2).Dtor_turns(), (_10_dt__update__tmp_h2).Dtor_title(), (_10_dt__update__tmp_h2).Dtor_status(), (_10_dt__update__tmp_h2).Dtor_modifiedAt(), (_10_dt__update__tmp_h2).Dtor_draft(), (_10_dt__update__tmp_h2).Dtor_activity(), _11_dt__update_hactiveTurn_h0, (_10_dt__update__tmp_h2).Dtor_steeringMessage(), (_10_dt__update__tmp_h2).Dtor_queuedMessages(), (_10_dt__update__tmp_h2).Dtor_nextCursor())
						}(_pat_let271_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Turn_.Create_Turn_(_6_id, _7_m, _dafny.SeqOf(), _dafny.UnicodeSeqOfUtf8Bytes("in-progress"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())))
				}(_pat_let270_0)
			}(s)
			_ = _9_withTurn
			var _12_s2 ChatState = func(_pat_let272_0 ChatState) ChatState {
				return func(_13_dt__update__tmp_h3 ChatState) ChatState {
					return func(_pat_let273_0 uint32) ChatState {
						return func(_14_dt__update_hstatus_h0 uint32) ChatState {
							return Companion_ChatState_.Create_ChatState_((_13_dt__update__tmp_h3).Dtor_turns(), (_13_dt__update__tmp_h3).Dtor_title(), _14_dt__update_hstatus_h0, (_13_dt__update__tmp_h3).Dtor_modifiedAt(), (_13_dt__update__tmp_h3).Dtor_draft(), (_13_dt__update__tmp_h3).Dtor_activity(), (_13_dt__update__tmp_h3).Dtor_activeTurn(), (_13_dt__update__tmp_h3).Dtor_steeringMessage(), (_13_dt__update__tmp_h3).Dtor_queuedMessages(), (_13_dt__update__tmp_h3).Dtor_nextCursor())
						}(_pat_let273_0)
					}(Companion_Default___.ClearBit(Companion_Default___.SummaryStatus(_9_withTurn, false), Companion_Default___.READ()))
				}(_pat_let272_0)
			}(_9_withTurn)
			_ = _12_s2
			var _15_s3 ChatState = (func() ChatState {
				if (((_8_qid).Is_Some()) && (((_12_s2).Dtor_steeringMessage()).Is_Some())) && (_dafny.Companion_Sequence_.Equal((((_12_s2).Dtor_steeringMessage()).Dtor_value().(QMsg)).Dtor_id(), (_8_qid).Dtor_value().(_dafny.Sequence))) {
					return func(_pat_let274_0 ChatState) ChatState {
						return func(_16_dt__update__tmp_h4 ChatState) ChatState {
							return func(_pat_let275_0 m_AhpSkeleton.Option) ChatState {
								return func(_17_dt__update_hsteeringMessage_h0 m_AhpSkeleton.Option) ChatState {
									return Companion_ChatState_.Create_ChatState_((_16_dt__update__tmp_h4).Dtor_turns(), (_16_dt__update__tmp_h4).Dtor_title(), (_16_dt__update__tmp_h4).Dtor_status(), (_16_dt__update__tmp_h4).Dtor_modifiedAt(), (_16_dt__update__tmp_h4).Dtor_draft(), (_16_dt__update__tmp_h4).Dtor_activity(), (_16_dt__update__tmp_h4).Dtor_activeTurn(), _17_dt__update_hsteeringMessage_h0, (_16_dt__update__tmp_h4).Dtor_queuedMessages(), (_16_dt__update__tmp_h4).Dtor_nextCursor())
								}(_pat_let275_0)
							}(m_AhpSkeleton.Companion_Option_.Create_None_())
						}(_pat_let274_0)
					}(_12_s2)
				}
				return _12_s2
			})()
			_ = _15_s3
			var _18_s4 ChatState = (func() ChatState {
				if (_8_qid).Is_Some() {
					return func(_pat_let276_0 ChatState) ChatState {
						return func(_19_dt__update__tmp_h5 ChatState) ChatState {
							return func(_pat_let277_0 _dafny.Sequence) ChatState {
								return func(_20_dt__update_hqueuedMessages_h0 _dafny.Sequence) ChatState {
									return Companion_ChatState_.Create_ChatState_((_19_dt__update__tmp_h5).Dtor_turns(), (_19_dt__update__tmp_h5).Dtor_title(), (_19_dt__update__tmp_h5).Dtor_status(), (_19_dt__update__tmp_h5).Dtor_modifiedAt(), (_19_dt__update__tmp_h5).Dtor_draft(), (_19_dt__update__tmp_h5).Dtor_activity(), (_19_dt__update__tmp_h5).Dtor_activeTurn(), (_19_dt__update__tmp_h5).Dtor_steeringMessage(), _20_dt__update_hqueuedMessages_h0, (_19_dt__update__tmp_h5).Dtor_nextCursor())
								}(_pat_let277_0)
							}(Companion_Default___.RemoveQ((_15_s3).Dtor_queuedMessages(), (_8_qid).Dtor_value().(_dafny.Sequence)))
						}(_pat_let276_0)
					}(_15_s3)
				}
				return _15_s3
			})()
			_ = _18_s4
			return _dafny.TupleOf(_18_s4, m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_CTurnComplete() {
			var _21_id _dafny.Sequence = _source0.Get_().(ChatAction_CTurnComplete).TurnId
			_ = _21_id
			return Companion_Default___.EndTurn(s, _21_id, _dafny.UnicodeSeqOfUtf8Bytes("complete"), false, m_AhpSkeleton.Companion_Option_.Create_None_())
		}
	}
	{
		if _source0.Is_CTurnCancelled() {
			var _22_id _dafny.Sequence = _source0.Get_().(ChatAction_CTurnCancelled).TurnId
			_ = _22_id
			return Companion_Default___.EndTurn(s, _22_id, _dafny.UnicodeSeqOfUtf8Bytes("cancelled"), false, m_AhpSkeleton.Companion_Option_.Create_None_())
		}
	}
	{
		if _source0.Is_CError() {
			var _23_id _dafny.Sequence = _source0.Get_().(ChatAction_CError).TurnId
			_ = _23_id
			var _24_e m_ConfluxCodec.Json = _source0.Get_().(ChatAction_CError).Err
			_ = _24_e
			return Companion_Default___.EndTurn(s, _23_id, _dafny.UnicodeSeqOfUtf8Bytes("error"), true, Companion_Default___.JNoNull(_24_e))
		}
	}
	{
		if _source0.Is_CUsage() {
			var _25_id _dafny.Sequence = _source0.Get_().(ChatAction_CUsage).TurnId
			_ = _25_id
			var _26_u m_ConfluxCodec.Json = _source0.Get_().(ChatAction_CUsage).Usage
			_ = _26_u
			if Companion_Default___.TurnMatches(s, _25_id) {
				return _dafny.TupleOf(func(_pat_let278_0 ChatState) ChatState {
					return func(_27_dt__update__tmp_h6 ChatState) ChatState {
						return func(_pat_let279_0 m_AhpSkeleton.Option) ChatState {
							return func(_30_dt__update_hactiveTurn_h1 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_27_dt__update__tmp_h6).Dtor_turns(), (_27_dt__update__tmp_h6).Dtor_title(), (_27_dt__update__tmp_h6).Dtor_status(), (_27_dt__update__tmp_h6).Dtor_modifiedAt(), (_27_dt__update__tmp_h6).Dtor_draft(), (_27_dt__update__tmp_h6).Dtor_activity(), _30_dt__update_hactiveTurn_h1, (_27_dt__update__tmp_h6).Dtor_steeringMessage(), (_27_dt__update__tmp_h6).Dtor_queuedMessages(), (_27_dt__update__tmp_h6).Dtor_nextCursor())
							}(_pat_let279_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let280_0 Turn) Turn {
							return func(_28_dt__update__tmp_h7 Turn) Turn {
								return func(_pat_let281_0 m_AhpSkeleton.Option) Turn {
									return func(_29_dt__update_husage_h0 m_AhpSkeleton.Option) Turn {
										return Companion_Turn_.Create_Turn_((_28_dt__update__tmp_h7).Dtor_turnId(), (_28_dt__update__tmp_h7).Dtor_message(), (_28_dt__update__tmp_h7).Dtor_parts(), (_28_dt__update__tmp_h7).Dtor_state(), _29_dt__update_husage_h0, (_28_dt__update__tmp_h7).Dtor_error())
									}(_pat_let281_0)
								}(Companion_Default___.JNoNull(_26_u))
							}(_pat_let280_0)
						}(((_pat_let_tv0).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let278_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CResponsePart() {
			var _31_id _dafny.Sequence = _source0.Get_().(ChatAction_CResponsePart).TurnId
			_ = _31_id
			var _32_p Part = _source0.Get_().(ChatAction_CResponsePart).Part
			_ = _32_p
			if Companion_Default___.TurnMatches(s, _31_id) {
				return _dafny.TupleOf(func(_pat_let282_0 ChatState) ChatState {
					return func(_33_dt__update__tmp_h8 ChatState) ChatState {
						return func(_pat_let283_0 m_AhpSkeleton.Option) ChatState {
							return func(_34_dt__update_hactiveTurn_h2 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_33_dt__update__tmp_h8).Dtor_turns(), (_33_dt__update__tmp_h8).Dtor_title(), (_33_dt__update__tmp_h8).Dtor_status(), (_33_dt__update__tmp_h8).Dtor_modifiedAt(), (_33_dt__update__tmp_h8).Dtor_draft(), (_33_dt__update__tmp_h8).Dtor_activity(), _34_dt__update_hactiveTurn_h2, (_33_dt__update__tmp_h8).Dtor_steeringMessage(), (_33_dt__update__tmp_h8).Dtor_queuedMessages(), (_33_dt__update__tmp_h8).Dtor_nextCursor())
							}(_pat_let283_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Default___.AppendPart(((_pat_let_tv1).Dtor_activeTurn()).Dtor_value().(Turn), _32_p)))
					}(_pat_let282_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CDelta() {
			var _35_id _dafny.Sequence = _source0.Get_().(ChatAction_CDelta).TurnId
			_ = _35_id
			var _36_pid _dafny.Sequence = _source0.Get_().(ChatAction_CDelta).PartId
			_ = _36_pid
			var _37_c _dafny.Sequence = _source0.Get_().(ChatAction_CDelta).Content
			_ = _37_c
			if Companion_Default___.TurnMatches(s, _35_id) {
				return _dafny.TupleOf(func(_pat_let284_0 ChatState) ChatState {
					return func(_38_dt__update__tmp_h9 ChatState) ChatState {
						return func(_pat_let285_0 m_AhpSkeleton.Option) ChatState {
							return func(_41_dt__update_hactiveTurn_h3 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_38_dt__update__tmp_h9).Dtor_turns(), (_38_dt__update__tmp_h9).Dtor_title(), (_38_dt__update__tmp_h9).Dtor_status(), (_38_dt__update__tmp_h9).Dtor_modifiedAt(), (_38_dt__update__tmp_h9).Dtor_draft(), (_38_dt__update__tmp_h9).Dtor_activity(), _41_dt__update_hactiveTurn_h3, (_38_dt__update__tmp_h9).Dtor_steeringMessage(), (_38_dt__update__tmp_h9).Dtor_queuedMessages(), (_38_dt__update__tmp_h9).Dtor_nextCursor())
							}(_pat_let285_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let286_0 Turn) Turn {
							return func(_39_dt__update__tmp_h10 Turn) Turn {
								return func(_pat_let287_0 _dafny.Sequence) Turn {
									return func(_40_dt__update_hparts_h0 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_39_dt__update__tmp_h10).Dtor_turnId(), (_39_dt__update__tmp_h10).Dtor_message(), _40_dt__update_hparts_h0, (_39_dt__update__tmp_h10).Dtor_state(), (_39_dt__update__tmp_h10).Dtor_usage(), (_39_dt__update__tmp_h10).Dtor_error())
									}(_pat_let287_0)
								}(Companion_Default___.DeltaMarkdown((((_pat_let_tv3).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _36_pid, _37_c))
							}(_pat_let286_0)
						}(((_pat_let_tv2).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let284_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CReasoning() {
			var _42_id _dafny.Sequence = _source0.Get_().(ChatAction_CReasoning).TurnId
			_ = _42_id
			var _43_pid _dafny.Sequence = _source0.Get_().(ChatAction_CReasoning).PartId
			_ = _43_pid
			var _44_c _dafny.Sequence = _source0.Get_().(ChatAction_CReasoning).Content
			_ = _44_c
			if Companion_Default___.TurnMatches(s, _42_id) {
				return _dafny.TupleOf(func(_pat_let288_0 ChatState) ChatState {
					return func(_45_dt__update__tmp_h11 ChatState) ChatState {
						return func(_pat_let289_0 m_AhpSkeleton.Option) ChatState {
							return func(_48_dt__update_hactiveTurn_h4 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_45_dt__update__tmp_h11).Dtor_turns(), (_45_dt__update__tmp_h11).Dtor_title(), (_45_dt__update__tmp_h11).Dtor_status(), (_45_dt__update__tmp_h11).Dtor_modifiedAt(), (_45_dt__update__tmp_h11).Dtor_draft(), (_45_dt__update__tmp_h11).Dtor_activity(), _48_dt__update_hactiveTurn_h4, (_45_dt__update__tmp_h11).Dtor_steeringMessage(), (_45_dt__update__tmp_h11).Dtor_queuedMessages(), (_45_dt__update__tmp_h11).Dtor_nextCursor())
							}(_pat_let289_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let290_0 Turn) Turn {
							return func(_46_dt__update__tmp_h12 Turn) Turn {
								return func(_pat_let291_0 _dafny.Sequence) Turn {
									return func(_47_dt__update_hparts_h1 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_46_dt__update__tmp_h12).Dtor_turnId(), (_46_dt__update__tmp_h12).Dtor_message(), _47_dt__update_hparts_h1, (_46_dt__update__tmp_h12).Dtor_state(), (_46_dt__update__tmp_h12).Dtor_usage(), (_46_dt__update__tmp_h12).Dtor_error())
									}(_pat_let291_0)
								}(Companion_Default___.DeltaReasoning((((_pat_let_tv5).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _43_pid, _44_c))
							}(_pat_let290_0)
						}(((_pat_let_tv4).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let288_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CToolCallStart() {
			var _49_id _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallStart).TurnId
			_ = _49_id
			var _50_tcId _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallStart).ToolCallId
			_ = _50_tcId
			var _51_tn _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallStart).ToolName
			_ = _51_tn
			var _52_dn _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallStart).DisplayName
			_ = _52_dn
			var _53_intent m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallStart).Intention
			_ = _53_intent
			var _54_contributor m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallStart).Contributor
			_ = _54_contributor
			var _55_meta m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallStart).Meta
			_ = _55_meta
			if Companion_Default___.TurnMatches(s, _49_id) {
				var _56_tc ToolCall = Companion_ToolCall_.Create_ToolCall_(_50_tcId, _51_tn, _52_dn, _dafny.UnicodeSeqOfUtf8Bytes("streaming"), _53_intent, _54_contributor, _55_meta, _dafny.UnicodeSeqOfUtf8Bytes(""), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())
				_ = _56_tc
				return _dafny.TupleOf(func(_pat_let292_0 ChatState) ChatState {
					return func(_57_dt__update__tmp_h13 ChatState) ChatState {
						return func(_pat_let293_0 m_AhpSkeleton.Option) ChatState {
							return func(_60_dt__update_hactiveTurn_h5 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_57_dt__update__tmp_h13).Dtor_turns(), (_57_dt__update__tmp_h13).Dtor_title(), (_57_dt__update__tmp_h13).Dtor_status(), (_57_dt__update__tmp_h13).Dtor_modifiedAt(), (_57_dt__update__tmp_h13).Dtor_draft(), (_57_dt__update__tmp_h13).Dtor_activity(), _60_dt__update_hactiveTurn_h5, (_57_dt__update__tmp_h13).Dtor_steeringMessage(), (_57_dt__update__tmp_h13).Dtor_queuedMessages(), (_57_dt__update__tmp_h13).Dtor_nextCursor())
							}(_pat_let293_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let294_0 Turn) Turn {
							return func(_58_dt__update__tmp_h14 Turn) Turn {
								return func(_pat_let295_0 _dafny.Sequence) Turn {
									return func(_59_dt__update_hparts_h2 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_58_dt__update__tmp_h14).Dtor_turnId(), (_58_dt__update__tmp_h14).Dtor_message(), _59_dt__update_hparts_h2, (_58_dt__update__tmp_h14).Dtor_state(), (_58_dt__update__tmp_h14).Dtor_usage(), (_58_dt__update__tmp_h14).Dtor_error())
									}(_pat_let295_0)
								}(_dafny.Companion_Sequence_.Concatenate((((_pat_let_tv7).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _dafny.SeqOf(Companion_Part_.Create_PToolCall_(_56_tc))))
							}(_pat_let294_0)
						}(((_pat_let_tv6).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let292_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CToolCallDelta() {
			var _61_id _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallDelta).TurnId
			_ = _61_id
			var _62_tcId _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallDelta).ToolCallId
			_ = _62_tcId
			var _63_c _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallDelta).Content
			_ = _63_c
			var _64_inv m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallDelta).InvocationMessage
			_ = _64_inv
			var _65_meta m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallDelta).Meta
			_ = _65_meta
			if Companion_Default___.TurnMatches(s, _61_id) {
				return _dafny.TupleOf(func(_pat_let296_0 ChatState) ChatState {
					return func(_66_dt__update__tmp_h15 ChatState) ChatState {
						return func(_pat_let297_0 m_AhpSkeleton.Option) ChatState {
							return func(_73_dt__update_hactiveTurn_h6 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_66_dt__update__tmp_h15).Dtor_turns(), (_66_dt__update__tmp_h15).Dtor_title(), (_66_dt__update__tmp_h15).Dtor_status(), (_66_dt__update__tmp_h15).Dtor_modifiedAt(), (_66_dt__update__tmp_h15).Dtor_draft(), (_66_dt__update__tmp_h15).Dtor_activity(), _73_dt__update_hactiveTurn_h6, (_66_dt__update__tmp_h15).Dtor_steeringMessage(), (_66_dt__update__tmp_h15).Dtor_queuedMessages(), (_66_dt__update__tmp_h15).Dtor_nextCursor())
							}(_pat_let297_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let298_0 Turn) Turn {
							return func(_67_dt__update__tmp_h16 Turn) Turn {
								return func(_pat_let299_0 _dafny.Sequence) Turn {
									return func(_72_dt__update_hparts_h3 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_67_dt__update__tmp_h16).Dtor_turnId(), (_67_dt__update__tmp_h16).Dtor_message(), _72_dt__update_hparts_h3, (_67_dt__update__tmp_h16).Dtor_state(), (_67_dt__update__tmp_h16).Dtor_usage(), (_67_dt__update__tmp_h16).Dtor_error())
									}(_pat_let299_0)
								}(Companion_Default___.MapTC((((_pat_let_tv9).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _62_tcId, (func(_68_c _dafny.Sequence, _69_inv m_AhpSkeleton.Option, _70_meta m_AhpSkeleton.Option) func(ToolCall) ToolCall {
									return func(_71_x ToolCall) ToolCall {
										return Companion_Default___.DeltaOne(_71_x, _68_c, _69_inv, _70_meta)
									}
								})(_63_c, _64_inv, _65_meta)))
							}(_pat_let298_0)
						}(((_pat_let_tv8).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let296_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CToolCallReady() {
			var _74_id _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallReady).TurnId
			_ = _74_id
			var _75_tcId _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallReady).ToolCallId
			_ = _75_tcId
			var _76_inv m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallReady).InvocationMessage
			_ = _76_inv
			var _77_input m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallReady).ToolInput
			_ = _77_input
			var _78_title m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallReady).ConfirmationTitle
			_ = _78_title
			var _79_risk m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallReady).RiskAssessment
			_ = _79_risk
			var _80_edits m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallReady).Edits
			_ = _80_edits
			var _81_editable m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallReady).Editable
			_ = _81_editable
			var _82_options m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallReady).Options
			_ = _82_options
			var _83_confirmed m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallReady).Confirmed
			_ = _83_confirmed
			var _84_meta m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallReady).Meta
			_ = _84_meta
			if Companion_Default___.TurnMatches(s, _74_id) {
				var _85_s1 ChatState = func(_pat_let300_0 ChatState) ChatState {
					return func(_86_dt__update__tmp_h17 ChatState) ChatState {
						return func(_pat_let301_0 m_AhpSkeleton.Option) ChatState {
							return func(_99_dt__update_hactiveTurn_h7 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_86_dt__update__tmp_h17).Dtor_turns(), (_86_dt__update__tmp_h17).Dtor_title(), (_86_dt__update__tmp_h17).Dtor_status(), (_86_dt__update__tmp_h17).Dtor_modifiedAt(), (_86_dt__update__tmp_h17).Dtor_draft(), (_86_dt__update__tmp_h17).Dtor_activity(), _99_dt__update_hactiveTurn_h7, (_86_dt__update__tmp_h17).Dtor_steeringMessage(), (_86_dt__update__tmp_h17).Dtor_queuedMessages(), (_86_dt__update__tmp_h17).Dtor_nextCursor())
							}(_pat_let301_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let302_0 Turn) Turn {
							return func(_87_dt__update__tmp_h18 Turn) Turn {
								return func(_pat_let303_0 _dafny.Sequence) Turn {
									return func(_98_dt__update_hparts_h4 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_87_dt__update__tmp_h18).Dtor_turnId(), (_87_dt__update__tmp_h18).Dtor_message(), _98_dt__update_hparts_h4, (_87_dt__update__tmp_h18).Dtor_state(), (_87_dt__update__tmp_h18).Dtor_usage(), (_87_dt__update__tmp_h18).Dtor_error())
									}(_pat_let303_0)
								}(Companion_Default___.MapTC((((_pat_let_tv11).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _75_tcId, (func(_88_inv m_AhpSkeleton.Option, _89_input m_AhpSkeleton.Option, _90_title m_AhpSkeleton.Option, _91_risk m_AhpSkeleton.Option, _92_edits m_AhpSkeleton.Option, _93_editable m_AhpSkeleton.Option, _94_options m_AhpSkeleton.Option, _95_confirmed m_AhpSkeleton.Option, _96_meta m_AhpSkeleton.Option) func(ToolCall) ToolCall {
									return func(_97_x ToolCall) ToolCall {
										return Companion_Default___.ReadyOne(_97_x, _88_inv, _89_input, _90_title, _91_risk, _92_edits, _93_editable, _94_options, _95_confirmed, _96_meta)
									}
								})(_76_inv, _77_input, _78_title, _79_risk, _80_edits, _81_editable, _82_options, _83_confirmed, _84_meta)))
							}(_pat_let302_0)
						}(((_pat_let_tv10).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let300_0)
				}(s)
				_ = _85_s1
				return _dafny.TupleOf(func(_pat_let304_0 ChatState) ChatState {
					return func(_100_dt__update__tmp_h19 ChatState) ChatState {
						return func(_pat_let305_0 uint32) ChatState {
							return func(_101_dt__update_hstatus_h1 uint32) ChatState {
								return Companion_ChatState_.Create_ChatState_((_100_dt__update__tmp_h19).Dtor_turns(), (_100_dt__update__tmp_h19).Dtor_title(), _101_dt__update_hstatus_h1, (_100_dt__update__tmp_h19).Dtor_modifiedAt(), (_100_dt__update__tmp_h19).Dtor_draft(), (_100_dt__update__tmp_h19).Dtor_activity(), (_100_dt__update__tmp_h19).Dtor_activeTurn(), (_100_dt__update__tmp_h19).Dtor_steeringMessage(), (_100_dt__update__tmp_h19).Dtor_queuedMessages(), (_100_dt__update__tmp_h19).Dtor_nextCursor())
							}(_pat_let305_0)
						}(Companion_Default___.SummaryStatus(_85_s1, false))
					}(_pat_let304_0)
				}(_85_s1), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CToolCallConfirmed() {
			var _102_id _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallConfirmed).TurnId
			_ = _102_id
			var _103_tcId _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallConfirmed).ToolCallId
			_ = _103_tcId
			var _104_approved bool = _source0.Get_().(ChatAction_CToolCallConfirmed).Approved
			_ = _104_approved
			var _105_cv m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallConfirmed).ConfirmedVal
			_ = _105_cv
			var _106_rsn m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallConfirmed).Reason
			_ = _106_rsn
			var _107_reasonMessage m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallConfirmed).ReasonMessage
			_ = _107_reasonMessage
			var _108_userSuggestion m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallConfirmed).UserSuggestion
			_ = _108_userSuggestion
			var _109_edited m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallConfirmed).EditedToolInput
			_ = _109_edited
			var _110_selectedOptionId m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallConfirmed).SelectedOptionId
			_ = _110_selectedOptionId
			var _111_meta m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallConfirmed).Meta
			_ = _111_meta
			if Companion_Default___.TurnMatches(s, _102_id) {
				var _112_s1 ChatState = func(_pat_let306_0 ChatState) ChatState {
					return func(_113_dt__update__tmp_h20 ChatState) ChatState {
						return func(_pat_let307_0 m_AhpSkeleton.Option) ChatState {
							return func(_125_dt__update_hactiveTurn_h8 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_113_dt__update__tmp_h20).Dtor_turns(), (_113_dt__update__tmp_h20).Dtor_title(), (_113_dt__update__tmp_h20).Dtor_status(), (_113_dt__update__tmp_h20).Dtor_modifiedAt(), (_113_dt__update__tmp_h20).Dtor_draft(), (_113_dt__update__tmp_h20).Dtor_activity(), _125_dt__update_hactiveTurn_h8, (_113_dt__update__tmp_h20).Dtor_steeringMessage(), (_113_dt__update__tmp_h20).Dtor_queuedMessages(), (_113_dt__update__tmp_h20).Dtor_nextCursor())
							}(_pat_let307_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let308_0 Turn) Turn {
							return func(_114_dt__update__tmp_h21 Turn) Turn {
								return func(_pat_let309_0 _dafny.Sequence) Turn {
									return func(_124_dt__update_hparts_h5 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_114_dt__update__tmp_h21).Dtor_turnId(), (_114_dt__update__tmp_h21).Dtor_message(), _124_dt__update_hparts_h5, (_114_dt__update__tmp_h21).Dtor_state(), (_114_dt__update__tmp_h21).Dtor_usage(), (_114_dt__update__tmp_h21).Dtor_error())
									}(_pat_let309_0)
								}(Companion_Default___.MapTC((((_pat_let_tv13).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _103_tcId, (func(_115_approved bool, _116_cv m_AhpSkeleton.Option, _117_rsn m_AhpSkeleton.Option, _118_reasonMessage m_AhpSkeleton.Option, _119_userSuggestion m_AhpSkeleton.Option, _120_edited m_AhpSkeleton.Option, _121_selectedOptionId m_AhpSkeleton.Option, _122_meta m_AhpSkeleton.Option) func(ToolCall) ToolCall {
									return func(_123_x ToolCall) ToolCall {
										return Companion_Default___.ConfirmOne(_123_x, _115_approved, _116_cv, _117_rsn, _118_reasonMessage, _119_userSuggestion, _120_edited, _121_selectedOptionId, _122_meta)
									}
								})(_104_approved, _105_cv, _106_rsn, _107_reasonMessage, _108_userSuggestion, _109_edited, _110_selectedOptionId, _111_meta)))
							}(_pat_let308_0)
						}(((_pat_let_tv12).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let306_0)
				}(s)
				_ = _112_s1
				return _dafny.TupleOf(func(_pat_let310_0 ChatState) ChatState {
					return func(_126_dt__update__tmp_h22 ChatState) ChatState {
						return func(_pat_let311_0 uint32) ChatState {
							return func(_127_dt__update_hstatus_h2 uint32) ChatState {
								return Companion_ChatState_.Create_ChatState_((_126_dt__update__tmp_h22).Dtor_turns(), (_126_dt__update__tmp_h22).Dtor_title(), _127_dt__update_hstatus_h2, (_126_dt__update__tmp_h22).Dtor_modifiedAt(), (_126_dt__update__tmp_h22).Dtor_draft(), (_126_dt__update__tmp_h22).Dtor_activity(), (_126_dt__update__tmp_h22).Dtor_activeTurn(), (_126_dt__update__tmp_h22).Dtor_steeringMessage(), (_126_dt__update__tmp_h22).Dtor_queuedMessages(), (_126_dt__update__tmp_h22).Dtor_nextCursor())
							}(_pat_let311_0)
						}(Companion_Default___.SummaryStatus(_112_s1, false))
					}(_pat_let310_0)
				}(_112_s1), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CToolCallAuthRequired() {
			var _128_id _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallAuthRequired).TurnId
			_ = _128_id
			var _129_tcId _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallAuthRequired).ToolCallId
			_ = _129_tcId
			var _130_auth m_ConfluxCodec.Json = _source0.Get_().(ChatAction_CToolCallAuthRequired).Auth
			_ = _130_auth
			var _131_meta m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallAuthRequired).Meta
			_ = _131_meta
			if Companion_Default___.TurnMatches(s, _128_id) {
				var _132_s1 ChatState = func(_pat_let312_0 ChatState) ChatState {
					return func(_133_dt__update__tmp_h23 ChatState) ChatState {
						return func(_pat_let313_0 m_AhpSkeleton.Option) ChatState {
							return func(_139_dt__update_hactiveTurn_h9 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_133_dt__update__tmp_h23).Dtor_turns(), (_133_dt__update__tmp_h23).Dtor_title(), (_133_dt__update__tmp_h23).Dtor_status(), (_133_dt__update__tmp_h23).Dtor_modifiedAt(), (_133_dt__update__tmp_h23).Dtor_draft(), (_133_dt__update__tmp_h23).Dtor_activity(), _139_dt__update_hactiveTurn_h9, (_133_dt__update__tmp_h23).Dtor_steeringMessage(), (_133_dt__update__tmp_h23).Dtor_queuedMessages(), (_133_dt__update__tmp_h23).Dtor_nextCursor())
							}(_pat_let313_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let314_0 Turn) Turn {
							return func(_134_dt__update__tmp_h24 Turn) Turn {
								return func(_pat_let315_0 _dafny.Sequence) Turn {
									return func(_138_dt__update_hparts_h6 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_134_dt__update__tmp_h24).Dtor_turnId(), (_134_dt__update__tmp_h24).Dtor_message(), _138_dt__update_hparts_h6, (_134_dt__update__tmp_h24).Dtor_state(), (_134_dt__update__tmp_h24).Dtor_usage(), (_134_dt__update__tmp_h24).Dtor_error())
									}(_pat_let315_0)
								}(Companion_Default___.MapTC((((_pat_let_tv15).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _129_tcId, (func(_135_auth m_ConfluxCodec.Json, _136_meta m_AhpSkeleton.Option) func(ToolCall) ToolCall {
									return func(_137_x ToolCall) ToolCall {
										return Companion_Default___.AuthRequiredOne(_137_x, _135_auth, _136_meta)
									}
								})(_130_auth, _131_meta)))
							}(_pat_let314_0)
						}(((_pat_let_tv14).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let312_0)
				}(s)
				_ = _132_s1
				return _dafny.TupleOf(func(_pat_let316_0 ChatState) ChatState {
					return func(_140_dt__update__tmp_h25 ChatState) ChatState {
						return func(_pat_let317_0 uint32) ChatState {
							return func(_141_dt__update_hstatus_h3 uint32) ChatState {
								return Companion_ChatState_.Create_ChatState_((_140_dt__update__tmp_h25).Dtor_turns(), (_140_dt__update__tmp_h25).Dtor_title(), _141_dt__update_hstatus_h3, (_140_dt__update__tmp_h25).Dtor_modifiedAt(), (_140_dt__update__tmp_h25).Dtor_draft(), (_140_dt__update__tmp_h25).Dtor_activity(), (_140_dt__update__tmp_h25).Dtor_activeTurn(), (_140_dt__update__tmp_h25).Dtor_steeringMessage(), (_140_dt__update__tmp_h25).Dtor_queuedMessages(), (_140_dt__update__tmp_h25).Dtor_nextCursor())
							}(_pat_let317_0)
						}(Companion_Default___.SummaryStatus(_132_s1, false))
					}(_pat_let316_0)
				}(_132_s1), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CToolCallAuthResolved() {
			var _142_id _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallAuthResolved).TurnId
			_ = _142_id
			var _143_tcId _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallAuthResolved).ToolCallId
			_ = _143_tcId
			var _144_meta m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallAuthResolved).Meta
			_ = _144_meta
			if Companion_Default___.TurnMatches(s, _142_id) {
				var _145_s1 ChatState = func(_pat_let318_0 ChatState) ChatState {
					return func(_146_dt__update__tmp_h26 ChatState) ChatState {
						return func(_pat_let319_0 m_AhpSkeleton.Option) ChatState {
							return func(_151_dt__update_hactiveTurn_h10 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_146_dt__update__tmp_h26).Dtor_turns(), (_146_dt__update__tmp_h26).Dtor_title(), (_146_dt__update__tmp_h26).Dtor_status(), (_146_dt__update__tmp_h26).Dtor_modifiedAt(), (_146_dt__update__tmp_h26).Dtor_draft(), (_146_dt__update__tmp_h26).Dtor_activity(), _151_dt__update_hactiveTurn_h10, (_146_dt__update__tmp_h26).Dtor_steeringMessage(), (_146_dt__update__tmp_h26).Dtor_queuedMessages(), (_146_dt__update__tmp_h26).Dtor_nextCursor())
							}(_pat_let319_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let320_0 Turn) Turn {
							return func(_147_dt__update__tmp_h27 Turn) Turn {
								return func(_pat_let321_0 _dafny.Sequence) Turn {
									return func(_150_dt__update_hparts_h7 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_147_dt__update__tmp_h27).Dtor_turnId(), (_147_dt__update__tmp_h27).Dtor_message(), _150_dt__update_hparts_h7, (_147_dt__update__tmp_h27).Dtor_state(), (_147_dt__update__tmp_h27).Dtor_usage(), (_147_dt__update__tmp_h27).Dtor_error())
									}(_pat_let321_0)
								}(Companion_Default___.MapTC((((_pat_let_tv17).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _143_tcId, (func(_148_meta m_AhpSkeleton.Option) func(ToolCall) ToolCall {
									return func(_149_x ToolCall) ToolCall {
										return Companion_Default___.AuthResolvedOne(_149_x, _148_meta)
									}
								})(_144_meta)))
							}(_pat_let320_0)
						}(((_pat_let_tv16).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let318_0)
				}(s)
				_ = _145_s1
				return _dafny.TupleOf(func(_pat_let322_0 ChatState) ChatState {
					return func(_152_dt__update__tmp_h28 ChatState) ChatState {
						return func(_pat_let323_0 uint32) ChatState {
							return func(_153_dt__update_hstatus_h4 uint32) ChatState {
								return Companion_ChatState_.Create_ChatState_((_152_dt__update__tmp_h28).Dtor_turns(), (_152_dt__update__tmp_h28).Dtor_title(), _153_dt__update_hstatus_h4, (_152_dt__update__tmp_h28).Dtor_modifiedAt(), (_152_dt__update__tmp_h28).Dtor_draft(), (_152_dt__update__tmp_h28).Dtor_activity(), (_152_dt__update__tmp_h28).Dtor_activeTurn(), (_152_dt__update__tmp_h28).Dtor_steeringMessage(), (_152_dt__update__tmp_h28).Dtor_queuedMessages(), (_152_dt__update__tmp_h28).Dtor_nextCursor())
							}(_pat_let323_0)
						}(Companion_Default___.SummaryStatus(_145_s1, false))
					}(_pat_let322_0)
				}(_145_s1), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CToolCallComplete() {
			var _154_id _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallComplete).TurnId
			_ = _154_id
			var _155_tcId _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallComplete).ToolCallId
			_ = _155_tcId
			var _156_ok bool = _source0.Get_().(ChatAction_CToolCallComplete).Success
			_ = _156_ok
			var _157_past m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallComplete).PastTenseMessage
			_ = _157_past
			var _158_resultContent m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallComplete).ResultContent
			_ = _158_resultContent
			var _159_structured m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallComplete).StructuredContent
			_ = _159_structured
			var _160_err m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallComplete).Error
			_ = _160_err
			var _161_rrc bool = _source0.Get_().(ChatAction_CToolCallComplete).RequiresResultConfirmation
			_ = _161_rrc
			var _162_meta m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallComplete).Meta
			_ = _162_meta
			if Companion_Default___.TurnMatches(s, _154_id) {
				var _163_s1 ChatState = func(_pat_let324_0 ChatState) ChatState {
					return func(_164_dt__update__tmp_h29 ChatState) ChatState {
						return func(_pat_let325_0 m_AhpSkeleton.Option) ChatState {
							return func(_175_dt__update_hactiveTurn_h11 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_164_dt__update__tmp_h29).Dtor_turns(), (_164_dt__update__tmp_h29).Dtor_title(), (_164_dt__update__tmp_h29).Dtor_status(), (_164_dt__update__tmp_h29).Dtor_modifiedAt(), (_164_dt__update__tmp_h29).Dtor_draft(), (_164_dt__update__tmp_h29).Dtor_activity(), _175_dt__update_hactiveTurn_h11, (_164_dt__update__tmp_h29).Dtor_steeringMessage(), (_164_dt__update__tmp_h29).Dtor_queuedMessages(), (_164_dt__update__tmp_h29).Dtor_nextCursor())
							}(_pat_let325_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let326_0 Turn) Turn {
							return func(_165_dt__update__tmp_h30 Turn) Turn {
								return func(_pat_let327_0 _dafny.Sequence) Turn {
									return func(_174_dt__update_hparts_h8 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_165_dt__update__tmp_h30).Dtor_turnId(), (_165_dt__update__tmp_h30).Dtor_message(), _174_dt__update_hparts_h8, (_165_dt__update__tmp_h30).Dtor_state(), (_165_dt__update__tmp_h30).Dtor_usage(), (_165_dt__update__tmp_h30).Dtor_error())
									}(_pat_let327_0)
								}(Companion_Default___.MapTC((((_pat_let_tv19).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _155_tcId, (func(_166_ok bool, _167_past m_AhpSkeleton.Option, _168_resultContent m_AhpSkeleton.Option, _169_structured m_AhpSkeleton.Option, _170_err m_AhpSkeleton.Option, _171_rrc bool, _172_meta m_AhpSkeleton.Option) func(ToolCall) ToolCall {
									return func(_173_x ToolCall) ToolCall {
										return Companion_Default___.CompleteOne(_173_x, _166_ok, _167_past, _168_resultContent, _169_structured, _170_err, _171_rrc, _172_meta)
									}
								})(_156_ok, _157_past, _158_resultContent, _159_structured, _160_err, _161_rrc, _162_meta)))
							}(_pat_let326_0)
						}(((_pat_let_tv18).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let324_0)
				}(s)
				_ = _163_s1
				return _dafny.TupleOf(func(_pat_let328_0 ChatState) ChatState {
					return func(_176_dt__update__tmp_h31 ChatState) ChatState {
						return func(_pat_let329_0 uint32) ChatState {
							return func(_177_dt__update_hstatus_h5 uint32) ChatState {
								return Companion_ChatState_.Create_ChatState_((_176_dt__update__tmp_h31).Dtor_turns(), (_176_dt__update__tmp_h31).Dtor_title(), _177_dt__update_hstatus_h5, (_176_dt__update__tmp_h31).Dtor_modifiedAt(), (_176_dt__update__tmp_h31).Dtor_draft(), (_176_dt__update__tmp_h31).Dtor_activity(), (_176_dt__update__tmp_h31).Dtor_activeTurn(), (_176_dt__update__tmp_h31).Dtor_steeringMessage(), (_176_dt__update__tmp_h31).Dtor_queuedMessages(), (_176_dt__update__tmp_h31).Dtor_nextCursor())
							}(_pat_let329_0)
						}(Companion_Default___.SummaryStatus(_163_s1, false))
					}(_pat_let328_0)
				}(_163_s1), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CToolCallResultConfirmed() {
			var _178_id _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallResultConfirmed).TurnId
			_ = _178_id
			var _179_tcId _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallResultConfirmed).ToolCallId
			_ = _179_tcId
			var _180_approved bool = _source0.Get_().(ChatAction_CToolCallResultConfirmed).Approved
			_ = _180_approved
			var _181_meta m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallResultConfirmed).Meta
			_ = _181_meta
			if Companion_Default___.TurnMatches(s, _178_id) {
				var _182_s1 ChatState = func(_pat_let330_0 ChatState) ChatState {
					return func(_183_dt__update__tmp_h32 ChatState) ChatState {
						return func(_pat_let331_0 m_AhpSkeleton.Option) ChatState {
							return func(_189_dt__update_hactiveTurn_h12 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_183_dt__update__tmp_h32).Dtor_turns(), (_183_dt__update__tmp_h32).Dtor_title(), (_183_dt__update__tmp_h32).Dtor_status(), (_183_dt__update__tmp_h32).Dtor_modifiedAt(), (_183_dt__update__tmp_h32).Dtor_draft(), (_183_dt__update__tmp_h32).Dtor_activity(), _189_dt__update_hactiveTurn_h12, (_183_dt__update__tmp_h32).Dtor_steeringMessage(), (_183_dt__update__tmp_h32).Dtor_queuedMessages(), (_183_dt__update__tmp_h32).Dtor_nextCursor())
							}(_pat_let331_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let332_0 Turn) Turn {
							return func(_184_dt__update__tmp_h33 Turn) Turn {
								return func(_pat_let333_0 _dafny.Sequence) Turn {
									return func(_188_dt__update_hparts_h9 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_184_dt__update__tmp_h33).Dtor_turnId(), (_184_dt__update__tmp_h33).Dtor_message(), _188_dt__update_hparts_h9, (_184_dt__update__tmp_h33).Dtor_state(), (_184_dt__update__tmp_h33).Dtor_usage(), (_184_dt__update__tmp_h33).Dtor_error())
									}(_pat_let333_0)
								}(Companion_Default___.MapTC((((_pat_let_tv21).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _179_tcId, (func(_185_approved bool, _186_meta m_AhpSkeleton.Option) func(ToolCall) ToolCall {
									return func(_187_x ToolCall) ToolCall {
										return Companion_Default___.ResultConfirmOne(_187_x, _185_approved, _186_meta)
									}
								})(_180_approved, _181_meta)))
							}(_pat_let332_0)
						}(((_pat_let_tv20).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let330_0)
				}(s)
				_ = _182_s1
				return _dafny.TupleOf(func(_pat_let334_0 ChatState) ChatState {
					return func(_190_dt__update__tmp_h34 ChatState) ChatState {
						return func(_pat_let335_0 uint32) ChatState {
							return func(_191_dt__update_hstatus_h6 uint32) ChatState {
								return Companion_ChatState_.Create_ChatState_((_190_dt__update__tmp_h34).Dtor_turns(), (_190_dt__update__tmp_h34).Dtor_title(), _191_dt__update_hstatus_h6, (_190_dt__update__tmp_h34).Dtor_modifiedAt(), (_190_dt__update__tmp_h34).Dtor_draft(), (_190_dt__update__tmp_h34).Dtor_activity(), (_190_dt__update__tmp_h34).Dtor_activeTurn(), (_190_dt__update__tmp_h34).Dtor_steeringMessage(), (_190_dt__update__tmp_h34).Dtor_queuedMessages(), (_190_dt__update__tmp_h34).Dtor_nextCursor())
							}(_pat_let335_0)
						}(Companion_Default___.SummaryStatus(_182_s1, false))
					}(_pat_let334_0)
				}(_182_s1), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CToolCallContentChanged() {
			var _192_id _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallContentChanged).TurnId
			_ = _192_id
			var _193_tcId _dafny.Sequence = _source0.Get_().(ChatAction_CToolCallContentChanged).ToolCallId
			_ = _193_tcId
			var _194_c m_ConfluxCodec.Json = _source0.Get_().(ChatAction_CToolCallContentChanged).NewContent
			_ = _194_c
			var _195_meta m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CToolCallContentChanged).Meta
			_ = _195_meta
			if Companion_Default___.TurnMatches(s, _192_id) {
				return _dafny.TupleOf(func(_pat_let336_0 ChatState) ChatState {
					return func(_196_dt__update__tmp_h35 ChatState) ChatState {
						return func(_pat_let337_0 m_AhpSkeleton.Option) ChatState {
							return func(_202_dt__update_hactiveTurn_h13 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_196_dt__update__tmp_h35).Dtor_turns(), (_196_dt__update__tmp_h35).Dtor_title(), (_196_dt__update__tmp_h35).Dtor_status(), (_196_dt__update__tmp_h35).Dtor_modifiedAt(), (_196_dt__update__tmp_h35).Dtor_draft(), (_196_dt__update__tmp_h35).Dtor_activity(), _202_dt__update_hactiveTurn_h13, (_196_dt__update__tmp_h35).Dtor_steeringMessage(), (_196_dt__update__tmp_h35).Dtor_queuedMessages(), (_196_dt__update__tmp_h35).Dtor_nextCursor())
							}(_pat_let337_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let338_0 Turn) Turn {
							return func(_197_dt__update__tmp_h36 Turn) Turn {
								return func(_pat_let339_0 _dafny.Sequence) Turn {
									return func(_201_dt__update_hparts_h10 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_197_dt__update__tmp_h36).Dtor_turnId(), (_197_dt__update__tmp_h36).Dtor_message(), _201_dt__update_hparts_h10, (_197_dt__update__tmp_h36).Dtor_state(), (_197_dt__update__tmp_h36).Dtor_usage(), (_197_dt__update__tmp_h36).Dtor_error())
									}(_pat_let339_0)
								}(Companion_Default___.MapTC((((_pat_let_tv23).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _193_tcId, (func(_198_c m_ConfluxCodec.Json, _199_meta m_AhpSkeleton.Option) func(ToolCall) ToolCall {
									return func(_200_x ToolCall) ToolCall {
										return Companion_Default___.ContentChangedOne(_200_x, _198_c, _199_meta)
									}
								})(_194_c, _195_meta)))
							}(_pat_let338_0)
						}(((_pat_let_tv22).Dtor_activeTurn()).Dtor_value().(Turn))))
					}(_pat_let336_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CTruncated() {
			var _203_idOpt m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CTruncated).UpToTurnId
			_ = _203_idOpt
			if ((_203_idOpt).Is_Some()) && (!(Companion_Default___.HasTurn((s).Dtor_turns(), (_203_idOpt).Dtor_value().(_dafny.Sequence)))) {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			} else {
				var _204_kept _dafny.Sequence = (func() _dafny.Sequence {
					if (_203_idOpt).Is_Some() {
						return Companion_Default___.KeepUpTo((s).Dtor_turns(), (_203_idOpt).Dtor_value().(_dafny.Sequence))
					}
					return _dafny.SeqOf()
				})()
				_ = _204_kept
				var _205_next ChatState = func(_pat_let340_0 ChatState) ChatState {
					return func(_206_dt__update__tmp_h37 ChatState) ChatState {
						return func(_pat_let341_0 m_AhpSkeleton.Option) ChatState {
							return func(_207_dt__update_hnextCursor_h0 m_AhpSkeleton.Option) ChatState {
								return func(_pat_let342_0 m_AhpSkeleton.Option) ChatState {
									return func(_208_dt__update_hactiveTurn_h14 m_AhpSkeleton.Option) ChatState {
										return func(_pat_let343_0 _dafny.Sequence) ChatState {
											return func(_209_dt__update_hturns_h0 _dafny.Sequence) ChatState {
												return Companion_ChatState_.Create_ChatState_(_209_dt__update_hturns_h0, (_206_dt__update__tmp_h37).Dtor_title(), (_206_dt__update__tmp_h37).Dtor_status(), (_206_dt__update__tmp_h37).Dtor_modifiedAt(), (_206_dt__update__tmp_h37).Dtor_draft(), (_206_dt__update__tmp_h37).Dtor_activity(), _208_dt__update_hactiveTurn_h14, (_206_dt__update__tmp_h37).Dtor_steeringMessage(), (_206_dt__update__tmp_h37).Dtor_queuedMessages(), _207_dt__update_hnextCursor_h0)
											}(_pat_let343_0)
										}(_204_kept)
									}(_pat_let342_0)
								}(m_AhpSkeleton.Companion_Option_.Create_None_())
							}(_pat_let341_0)
						}((func() m_AhpSkeleton.Option {
							if (_203_idOpt).Is_Some() {
								return (_pat_let_tv24).Dtor_nextCursor()
							}
							return m_AhpSkeleton.Companion_Option_.Create_None_()
						})())
					}(_pat_let340_0)
				}(s)
				_ = _205_next
				return _dafny.TupleOf(func(_pat_let344_0 ChatState) ChatState {
					return func(_210_dt__update__tmp_h38 ChatState) ChatState {
						return func(_pat_let345_0 uint32) ChatState {
							return func(_211_dt__update_hstatus_h7 uint32) ChatState {
								return Companion_ChatState_.Create_ChatState_((_210_dt__update__tmp_h38).Dtor_turns(), (_210_dt__update__tmp_h38).Dtor_title(), _211_dt__update_hstatus_h7, (_210_dt__update__tmp_h38).Dtor_modifiedAt(), (_210_dt__update__tmp_h38).Dtor_draft(), (_210_dt__update__tmp_h38).Dtor_activity(), (_210_dt__update__tmp_h38).Dtor_activeTurn(), (_210_dt__update__tmp_h38).Dtor_steeringMessage(), (_210_dt__update__tmp_h38).Dtor_queuedMessages(), (_210_dt__update__tmp_h38).Dtor_nextCursor())
							}(_pat_let345_0)
						}(Companion_Default___.SummaryStatus(_205_next, false))
					}(_pat_let344_0)
				}(_205_next), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			}
		}
	}
	{
		if _source0.Is_CQueuedReordered() {
			var _212_order _dafny.Sequence = _source0.Get_().(ChatAction_CQueuedReordered).Order
			_ = _212_order
			return _dafny.TupleOf(func(_pat_let346_0 ChatState) ChatState {
				return func(_213_dt__update__tmp_h39 ChatState) ChatState {
					return func(_pat_let347_0 _dafny.Sequence) ChatState {
						return func(_214_dt__update_hqueuedMessages_h1 _dafny.Sequence) ChatState {
							return Companion_ChatState_.Create_ChatState_((_213_dt__update__tmp_h39).Dtor_turns(), (_213_dt__update__tmp_h39).Dtor_title(), (_213_dt__update__tmp_h39).Dtor_status(), (_213_dt__update__tmp_h39).Dtor_modifiedAt(), (_213_dt__update__tmp_h39).Dtor_draft(), (_213_dt__update__tmp_h39).Dtor_activity(), (_213_dt__update__tmp_h39).Dtor_activeTurn(), (_213_dt__update__tmp_h39).Dtor_steeringMessage(), _214_dt__update_hqueuedMessages_h1, (_213_dt__update__tmp_h39).Dtor_nextCursor())
						}(_pat_let347_0)
					}(Companion_Default___.ReorderQ((_pat_let_tv25).Dtor_queuedMessages(), _212_order))
				}(_pat_let346_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_CTurnsLoaded() {
			var _215_loaded _dafny.Sequence = _source0.Get_().(ChatAction_CTurnsLoaded).Loaded
			_ = _215_loaded
			var _216_cursor m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CTurnsLoaded).Cursor
			_ = _216_cursor
			return _dafny.TupleOf(func(_pat_let348_0 ChatState) ChatState {
				return func(_217_dt__update__tmp_h40 ChatState) ChatState {
					return func(_pat_let349_0 m_AhpSkeleton.Option) ChatState {
						return func(_218_dt__update_hnextCursor_h1 m_AhpSkeleton.Option) ChatState {
							return func(_pat_let350_0 _dafny.Sequence) ChatState {
								return func(_219_dt__update_hturns_h1 _dafny.Sequence) ChatState {
									return Companion_ChatState_.Create_ChatState_(_219_dt__update_hturns_h1, (_217_dt__update__tmp_h40).Dtor_title(), (_217_dt__update__tmp_h40).Dtor_status(), (_217_dt__update__tmp_h40).Dtor_modifiedAt(), (_217_dt__update__tmp_h40).Dtor_draft(), (_217_dt__update__tmp_h40).Dtor_activity(), (_217_dt__update__tmp_h40).Dtor_activeTurn(), (_217_dt__update__tmp_h40).Dtor_steeringMessage(), (_217_dt__update__tmp_h40).Dtor_queuedMessages(), _218_dt__update_hnextCursor_h1)
								}(_pat_let350_0)
							}(Companion_Default___.DedupPrepend(_215_loaded, (_pat_let_tv26).Dtor_turns()))
						}(_pat_let349_0)
					}(_216_cursor)
				}(_pat_let348_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_CInputRequested() {
			var _220_req InputReq = _source0.Get_().(ChatAction_CInputRequested).Req
			_ = _220_req
			if ((s).Dtor_activeTurn()).Is_Some() {
				var _221_at Turn = ((s).Dtor_activeTurn()).Dtor_value().(Turn)
				_ = _221_at
				var _222_next ChatState = func(_pat_let351_0 ChatState) ChatState {
					return func(_223_dt__update__tmp_h41 ChatState) ChatState {
						return func(_pat_let352_0 m_AhpSkeleton.Option) ChatState {
							return func(_226_dt__update_hactiveTurn_h15 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_223_dt__update__tmp_h41).Dtor_turns(), (_223_dt__update__tmp_h41).Dtor_title(), (_223_dt__update__tmp_h41).Dtor_status(), (_223_dt__update__tmp_h41).Dtor_modifiedAt(), (_223_dt__update__tmp_h41).Dtor_draft(), (_223_dt__update__tmp_h41).Dtor_activity(), _226_dt__update_hactiveTurn_h15, (_223_dt__update__tmp_h41).Dtor_steeringMessage(), (_223_dt__update__tmp_h41).Dtor_queuedMessages(), (_223_dt__update__tmp_h41).Dtor_nextCursor())
							}(_pat_let352_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let353_0 Turn) Turn {
							return func(_224_dt__update__tmp_h42 Turn) Turn {
								return func(_pat_let354_0 _dafny.Sequence) Turn {
									return func(_225_dt__update_hparts_h11 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_224_dt__update__tmp_h42).Dtor_turnId(), (_224_dt__update__tmp_h42).Dtor_message(), _225_dt__update_hparts_h11, (_224_dt__update__tmp_h42).Dtor_state(), (_224_dt__update__tmp_h42).Dtor_usage(), (_224_dt__update__tmp_h42).Dtor_error())
									}(_pat_let354_0)
								}(Companion_Default___.UpsertInputPart((_221_at).Dtor_parts(), _220_req))
							}(_pat_let353_0)
						}(_221_at)))
					}(_pat_let351_0)
				}(s)
				_ = _222_next
				return _dafny.TupleOf(func(_pat_let355_0 ChatState) ChatState {
					return func(_227_dt__update__tmp_h43 ChatState) ChatState {
						return func(_pat_let356_0 uint32) ChatState {
							return func(_228_dt__update_hstatus_h8 uint32) ChatState {
								return Companion_ChatState_.Create_ChatState_((_227_dt__update__tmp_h43).Dtor_turns(), (_227_dt__update__tmp_h43).Dtor_title(), _228_dt__update_hstatus_h8, (_227_dt__update__tmp_h43).Dtor_modifiedAt(), (_227_dt__update__tmp_h43).Dtor_draft(), (_227_dt__update__tmp_h43).Dtor_activity(), (_227_dt__update__tmp_h43).Dtor_activeTurn(), (_227_dt__update__tmp_h43).Dtor_steeringMessage(), (_227_dt__update__tmp_h43).Dtor_queuedMessages(), (_227_dt__update__tmp_h43).Dtor_nextCursor())
							}(_pat_let356_0)
						}(Companion_Default___.ClearBit(Companion_Default___.SummaryStatus(_222_next, false), Companion_Default___.READ()))
					}(_pat_let355_0)
				}(_222_next), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CInputAnswerChanged() {
			var _229_reqId _dafny.Sequence = _source0.Get_().(ChatAction_CInputAnswerChanged).RequestId
			_ = _229_reqId
			var _230_qId _dafny.Sequence = _source0.Get_().(ChatAction_CInputAnswerChanged).QuestionId
			_ = _230_qId
			var _231_ans m_AhpSkeleton.Option = _source0.Get_().(ChatAction_CInputAnswerChanged).Answer
			_ = _231_ans
			if (((s).Dtor_activeTurn()).Is_Some()) && (Companion_Default___.HasOpenReq((((s).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _229_reqId)) {
				var _232_at Turn = ((s).Dtor_activeTurn()).Dtor_value().(Turn)
				_ = _232_at
				return _dafny.TupleOf(func(_pat_let357_0 ChatState) ChatState {
					return func(_233_dt__update__tmp_h44 ChatState) ChatState {
						return func(_pat_let358_0 m_AhpSkeleton.Option) ChatState {
							return func(_236_dt__update_hactiveTurn_h16 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_233_dt__update__tmp_h44).Dtor_turns(), (_233_dt__update__tmp_h44).Dtor_title(), (_233_dt__update__tmp_h44).Dtor_status(), (_233_dt__update__tmp_h44).Dtor_modifiedAt(), (_233_dt__update__tmp_h44).Dtor_draft(), (_233_dt__update__tmp_h44).Dtor_activity(), _236_dt__update_hactiveTurn_h16, (_233_dt__update__tmp_h44).Dtor_steeringMessage(), (_233_dt__update__tmp_h44).Dtor_queuedMessages(), (_233_dt__update__tmp_h44).Dtor_nextCursor())
							}(_pat_let358_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let359_0 Turn) Turn {
							return func(_234_dt__update__tmp_h45 Turn) Turn {
								return func(_pat_let360_0 _dafny.Sequence) Turn {
									return func(_235_dt__update_hparts_h12 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_234_dt__update__tmp_h45).Dtor_turnId(), (_234_dt__update__tmp_h45).Dtor_message(), _235_dt__update_hparts_h12, (_234_dt__update__tmp_h45).Dtor_state(), (_234_dt__update__tmp_h45).Dtor_usage(), (_234_dt__update__tmp_h45).Dtor_error())
									}(_pat_let360_0)
								}(Companion_Default___.ChangeAnswerPart((_232_at).Dtor_parts(), _229_reqId, _230_qId, _231_ans))
							}(_pat_let359_0)
						}(_232_at)))
					}(_pat_let357_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CInputCompleted() {
			var _237_reqId _dafny.Sequence = _source0.Get_().(ChatAction_CInputCompleted).RequestId
			_ = _237_reqId
			var _238_resp _dafny.Sequence = _source0.Get_().(ChatAction_CInputCompleted).Response
			_ = _238_resp
			var _239_answers _dafny.Map = _source0.Get_().(ChatAction_CInputCompleted).Answers
			_ = _239_answers
			if (((s).Dtor_activeTurn()).Is_Some()) && (Companion_Default___.HasOpenReq((((s).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _237_reqId)) {
				var _240_at Turn = ((s).Dtor_activeTurn()).Dtor_value().(Turn)
				_ = _240_at
				var _241_next ChatState = func(_pat_let361_0 ChatState) ChatState {
					return func(_242_dt__update__tmp_h46 ChatState) ChatState {
						return func(_pat_let362_0 m_AhpSkeleton.Option) ChatState {
							return func(_245_dt__update_hactiveTurn_h17 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_242_dt__update__tmp_h46).Dtor_turns(), (_242_dt__update__tmp_h46).Dtor_title(), (_242_dt__update__tmp_h46).Dtor_status(), (_242_dt__update__tmp_h46).Dtor_modifiedAt(), (_242_dt__update__tmp_h46).Dtor_draft(), (_242_dt__update__tmp_h46).Dtor_activity(), _245_dt__update_hactiveTurn_h17, (_242_dt__update__tmp_h46).Dtor_steeringMessage(), (_242_dt__update__tmp_h46).Dtor_queuedMessages(), (_242_dt__update__tmp_h46).Dtor_nextCursor())
							}(_pat_let362_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let363_0 Turn) Turn {
							return func(_243_dt__update__tmp_h47 Turn) Turn {
								return func(_pat_let364_0 _dafny.Sequence) Turn {
									return func(_244_dt__update_hparts_h13 _dafny.Sequence) Turn {
										return Companion_Turn_.Create_Turn_((_243_dt__update__tmp_h47).Dtor_turnId(), (_243_dt__update__tmp_h47).Dtor_message(), _244_dt__update_hparts_h13, (_243_dt__update__tmp_h47).Dtor_state(), (_243_dt__update__tmp_h47).Dtor_usage(), (_243_dt__update__tmp_h47).Dtor_error())
									}(_pat_let364_0)
								}(Companion_Default___.CompleteAnswerPart((_240_at).Dtor_parts(), _237_reqId, _238_resp, _239_answers))
							}(_pat_let363_0)
						}(_240_at)))
					}(_pat_let361_0)
				}(s)
				_ = _241_next
				return _dafny.TupleOf(func(_pat_let365_0 ChatState) ChatState {
					return func(_246_dt__update__tmp_h48 ChatState) ChatState {
						return func(_pat_let366_0 uint32) ChatState {
							return func(_247_dt__update_hstatus_h9 uint32) ChatState {
								return Companion_ChatState_.Create_ChatState_((_246_dt__update__tmp_h48).Dtor_turns(), (_246_dt__update__tmp_h48).Dtor_title(), _247_dt__update_hstatus_h9, (_246_dt__update__tmp_h48).Dtor_modifiedAt(), (_246_dt__update__tmp_h48).Dtor_draft(), (_246_dt__update__tmp_h48).Dtor_activity(), (_246_dt__update__tmp_h48).Dtor_activeTurn(), (_246_dt__update__tmp_h48).Dtor_steeringMessage(), (_246_dt__update__tmp_h48).Dtor_queuedMessages(), (_246_dt__update__tmp_h48).Dtor_nextCursor())
							}(_pat_let366_0)
						}(Companion_Default___.SummaryStatus(_241_next, false))
					}(_pat_let365_0)
				}(_241_next), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_CPendingMessageSet() {
			var _248_kind _dafny.Sequence = _source0.Get_().(ChatAction_CPendingMessageSet).Kind
			_ = _248_kind
			var _249_id _dafny.Sequence = _source0.Get_().(ChatAction_CPendingMessageSet).Id
			_ = _249_id
			var _250_msg m_ConfluxCodec.Json = _source0.Get_().(ChatAction_CPendingMessageSet).Message
			_ = _250_msg
			if _dafny.Companion_Sequence_.Equal(_248_kind, _dafny.UnicodeSeqOfUtf8Bytes("steering")) {
				return _dafny.TupleOf(func(_pat_let367_0 ChatState) ChatState {
					return func(_251_dt__update__tmp_h49 ChatState) ChatState {
						return func(_pat_let368_0 m_AhpSkeleton.Option) ChatState {
							return func(_252_dt__update_hsteeringMessage_h1 m_AhpSkeleton.Option) ChatState {
								return Companion_ChatState_.Create_ChatState_((_251_dt__update__tmp_h49).Dtor_turns(), (_251_dt__update__tmp_h49).Dtor_title(), (_251_dt__update__tmp_h49).Dtor_status(), (_251_dt__update__tmp_h49).Dtor_modifiedAt(), (_251_dt__update__tmp_h49).Dtor_draft(), (_251_dt__update__tmp_h49).Dtor_activity(), (_251_dt__update__tmp_h49).Dtor_activeTurn(), _252_dt__update_hsteeringMessage_h1, (_251_dt__update__tmp_h49).Dtor_queuedMessages(), (_251_dt__update__tmp_h49).Dtor_nextCursor())
							}(_pat_let368_0)
						}(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_QMsg_.Create_QMsg_(_249_id, _250_msg)))
					}(_pat_let367_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(func(_pat_let369_0 ChatState) ChatState {
					return func(_253_dt__update__tmp_h50 ChatState) ChatState {
						return func(_pat_let370_0 _dafny.Sequence) ChatState {
							return func(_254_dt__update_hqueuedMessages_h2 _dafny.Sequence) ChatState {
								return Companion_ChatState_.Create_ChatState_((_253_dt__update__tmp_h50).Dtor_turns(), (_253_dt__update__tmp_h50).Dtor_title(), (_253_dt__update__tmp_h50).Dtor_status(), (_253_dt__update__tmp_h50).Dtor_modifiedAt(), (_253_dt__update__tmp_h50).Dtor_draft(), (_253_dt__update__tmp_h50).Dtor_activity(), (_253_dt__update__tmp_h50).Dtor_activeTurn(), (_253_dt__update__tmp_h50).Dtor_steeringMessage(), _254_dt__update_hqueuedMessages_h2, (_253_dt__update__tmp_h50).Dtor_nextCursor())
							}(_pat_let370_0)
						}(Companion_Default___.UpsertQ((_pat_let_tv27).Dtor_queuedMessages(), Companion_QMsg_.Create_QMsg_(_249_id, _250_msg)))
					}(_pat_let369_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			}
		}
	}
	{
		if _source0.Is_CPendingMessageRemoved() {
			var _255_kind _dafny.Sequence = _source0.Get_().(ChatAction_CPendingMessageRemoved).Kind
			_ = _255_kind
			var _256_id _dafny.Sequence = _source0.Get_().(ChatAction_CPendingMessageRemoved).Id
			_ = _256_id
			if _dafny.Companion_Sequence_.Equal(_255_kind, _dafny.UnicodeSeqOfUtf8Bytes("steering")) {
				if (((s).Dtor_steeringMessage()).Is_Some()) && (_dafny.Companion_Sequence_.Equal((((s).Dtor_steeringMessage()).Dtor_value().(QMsg)).Dtor_id(), _256_id)) {
					return _dafny.TupleOf(func(_pat_let371_0 ChatState) ChatState {
						return func(_257_dt__update__tmp_h51 ChatState) ChatState {
							return func(_pat_let372_0 m_AhpSkeleton.Option) ChatState {
								return func(_258_dt__update_hsteeringMessage_h2 m_AhpSkeleton.Option) ChatState {
									return Companion_ChatState_.Create_ChatState_((_257_dt__update__tmp_h51).Dtor_turns(), (_257_dt__update__tmp_h51).Dtor_title(), (_257_dt__update__tmp_h51).Dtor_status(), (_257_dt__update__tmp_h51).Dtor_modifiedAt(), (_257_dt__update__tmp_h51).Dtor_draft(), (_257_dt__update__tmp_h51).Dtor_activity(), (_257_dt__update__tmp_h51).Dtor_activeTurn(), _258_dt__update_hsteeringMessage_h2, (_257_dt__update__tmp_h51).Dtor_queuedMessages(), (_257_dt__update__tmp_h51).Dtor_nextCursor())
								}(_pat_let372_0)
							}(m_AhpSkeleton.Companion_Option_.Create_None_())
						}(_pat_let371_0)
					}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
				} else {
					return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
				}
			} else {
				return _dafny.TupleOf(func(_pat_let373_0 ChatState) ChatState {
					return func(_259_dt__update__tmp_h52 ChatState) ChatState {
						return func(_pat_let374_0 _dafny.Sequence) ChatState {
							return func(_260_dt__update_hqueuedMessages_h3 _dafny.Sequence) ChatState {
								return Companion_ChatState_.Create_ChatState_((_259_dt__update__tmp_h52).Dtor_turns(), (_259_dt__update__tmp_h52).Dtor_title(), (_259_dt__update__tmp_h52).Dtor_status(), (_259_dt__update__tmp_h52).Dtor_modifiedAt(), (_259_dt__update__tmp_h52).Dtor_draft(), (_259_dt__update__tmp_h52).Dtor_activity(), (_259_dt__update__tmp_h52).Dtor_activeTurn(), (_259_dt__update__tmp_h52).Dtor_steeringMessage(), _260_dt__update_hqueuedMessages_h3, (_259_dt__update__tmp_h52).Dtor_nextCursor())
							}(_pat_let374_0)
						}(Companion_Default___.RemoveQ((_pat_let_tv28).Dtor_queuedMessages(), _256_id))
					}(_pat_let373_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			}
		}
	}
	{
		return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
	}
}
func (_static *CompanionStruct_Default___) Apply1(s ChatState, a ChatAction) ChatState {
	return (*(Companion_Default___.ApplyToChat(s, a, _dafny.IntOfInt64(9999))).IndexInt(0)).(ChatState)
}
func (_static *CompanionStruct_Default___) Fold(s ChatState, acts _dafny.Sequence) ChatState {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer73 func(ChatState, ChatAction) ChatState) func(interface{}, interface{}) interface{} {
		return func(arg81 interface{}, arg82 interface{}) interface{} {
			return coer73(arg81.(ChatState), arg82.(ChatAction))
		}
	}(Companion_Default___.Apply1), s, acts).(ChatState)
}
func (_static *CompanionStruct_Default___) C0() ChatState {
	return Companion_ChatState_.Create_ChatState_(_dafny.SeqOf(), _dafny.UnicodeSeqOfUtf8Bytes("Chat"), uint32(1), _dafny.UnicodeSeqOfUtf8Bytes("t1"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_None_())
}
func (_static *CompanionStruct_Default___) T0(id _dafny.Sequence) Turn {
	return Companion_Turn_.Create_Turn_(id, m_ConfluxCodec.Companion_Json_.Create_JNull_(), _dafny.SeqOf(), _dafny.UnicodeSeqOfUtf8Bytes("in-progress"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())
}
func (_static *CompanionStruct_Default___) TDone(id _dafny.Sequence) Turn {
	return Companion_Turn_.Create_Turn_(id, m_ConfluxCodec.Companion_Json_.Create_JNull_(), _dafny.SeqOf(), _dafny.UnicodeSeqOfUtf8Bytes("complete"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())
}
func (_static *CompanionStruct_Default___) TC0(id _dafny.Sequence, intent m_AhpSkeleton.Option) ToolCall {
	return Companion_ToolCall_.Create_ToolCall_(id, _dafny.UnicodeSeqOfUtf8Bytes("bash"), _dafny.UnicodeSeqOfUtf8Bytes("Run Command"), _dafny.UnicodeSeqOfUtf8Bytes("streaming"), intent, m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), _dafny.UnicodeSeqOfUtf8Bytes(""), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())
}
func (_static *CompanionStruct_Default___) ActiveParts(s ChatState) _dafny.Sequence {
	if ((s).Dtor_activeTurn()).Is_Some() {
		return (((s).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts()
	} else {
		return _dafny.SeqOf()
	}
}
func (_static *CompanionStruct_Default___) FirstActiveToolCall(s ChatState) m_AhpSkeleton.Option {
	var _0_ps _dafny.Sequence = Companion_Default___.ActiveParts(s)
	_ = _0_ps
	if ((_dafny.IntOfUint32((_0_ps).Cardinality())).Sign() == 1) && (((_0_ps).Select(0).(Part)).Is_PToolCall()) {
		return m_AhpSkeleton.Companion_Option_.Create_Some_(((_0_ps).Select(0).(Part)).Dtor_tc())
	} else {
		return m_AhpSkeleton.Companion_Option_.Create_None_()
	}
}
func (_static *CompanionStruct_Default___) FirstStoredToolCall(s ChatState) m_AhpSkeleton.Option {
	if (_dafny.IntOfUint32(((s).Dtor_turns()).Cardinality())).Sign() == 0 {
		return m_AhpSkeleton.Companion_Option_.Create_None_()
	} else {
		var _0_ps _dafny.Sequence = (((s).Dtor_turns()).Select(0).(Turn)).Dtor_parts()
		_ = _0_ps
		if ((_dafny.IntOfUint32((_0_ps).Cardinality())).Sign() == 1) && (((_0_ps).Select(0).(Part)).Is_PToolCall()) {
			return m_AhpSkeleton.Companion_Option_.Create_Some_(((_0_ps).Select(0).(Part)).Dtor_tc())
		} else {
			return m_AhpSkeleton.Companion_Option_.Create_None_()
		}
	}
}
func (_static *CompanionStruct_Default___) RunScalarTurns() _dafny.Int {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	pass = _dafny.Zero
	var _0_s ChatState
	_ = _0_s
	_0_s = Companion_Default___.C0()
	if ((Companion_Default___.Apply1(func(_pat_let375_0 ChatState) ChatState {
		return func(_1_dt__update__tmp_h0 ChatState) ChatState {
			return func(_pat_let376_0 m_AhpSkeleton.Option) ChatState {
				return func(_2_dt__update_hdraft_h0 m_AhpSkeleton.Option) ChatState {
					return Companion_ChatState_.Create_ChatState_((_1_dt__update__tmp_h0).Dtor_turns(), (_1_dt__update__tmp_h0).Dtor_title(), (_1_dt__update__tmp_h0).Dtor_status(), (_1_dt__update__tmp_h0).Dtor_modifiedAt(), _2_dt__update_hdraft_h0, (_1_dt__update__tmp_h0).Dtor_activity(), (_1_dt__update__tmp_h0).Dtor_activeTurn(), (_1_dt__update__tmp_h0).Dtor_steeringMessage(), (_1_dt__update__tmp_h0).Dtor_queuedMessages(), (_1_dt__update__tmp_h0).Dtor_nextCursor())
				}(_pat_let376_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("x")))
		}(_pat_let375_0)
	}(_0_s), Companion_ChatAction_.Create_CDraftChanged_(m_AhpSkeleton.Companion_Option_.Create_None_()))).Dtor_draft()).Equals(m_AhpSkeleton.Companion_Option_.Create_None_()) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_0_s, Companion_ChatAction_.Create_CDraftChanged_(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("hi"))))).Dtor_draft()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("hi"))) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_0_s, Companion_ChatAction_.Create_CActivityChanged_(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Running terminal command"))))).Dtor_activity()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Running terminal command"))) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_0_s, Companion_ChatAction_.Create_CActivityChanged_(m_AhpSkeleton.Companion_Option_.Create_None_()))).Dtor_activity()).Equals(m_AhpSkeleton.Companion_Option_.Create_None_()) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_0_s, Companion_ChatAction_.Create_CTurnStarted_(_dafny.UnicodeSeqOfUtf8Bytes("turn-1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("Hello")), m_AhpSkeleton.Companion_Option_.Create_None_()))).Dtor_activeTurn()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Turn_.Create_Turn_(_dafny.UnicodeSeqOfUtf8Bytes("turn-1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("Hello")), _dafny.SeqOf(), _dafny.UnicodeSeqOfUtf8Bytes("in-progress"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _3_act1 ChatState
	_ = _3_act1
	var _4_dt__update__tmp_h1 ChatState = _0_s
	_ = _4_dt__update__tmp_h1
	var _5_dt__update_hactiveTurn_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1")))
	_ = _5_dt__update_hactiveTurn_h0
	_3_act1 = Companion_ChatState_.Create_ChatState_((_4_dt__update__tmp_h1).Dtor_turns(), (_4_dt__update__tmp_h1).Dtor_title(), (_4_dt__update__tmp_h1).Dtor_status(), (_4_dt__update__tmp_h1).Dtor_modifiedAt(), (_4_dt__update__tmp_h1).Dtor_draft(), (_4_dt__update__tmp_h1).Dtor_activity(), _5_dt__update_hactiveTurn_h0, (_4_dt__update__tmp_h1).Dtor_steeringMessage(), (_4_dt__update__tmp_h1).Dtor_queuedMessages(), (_4_dt__update__tmp_h1).Dtor_nextCursor())
	if (Companion_Default___.Apply1(_3_act1, Companion_ChatAction_.Create_CTurnComplete_(_dafny.UnicodeSeqOfUtf8Bytes("wrong-turn")))).Equals(_3_act1) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_3_act1, Companion_ChatAction_.Create_CUsage_(_dafny.UnicodeSeqOfUtf8Bytes("wrong-turn"), m_ConfluxCodec.Companion_Json_.Create_JNull_()))).Equals(_3_act1) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_3_act1, Companion_ChatAction_.Create_CToolCallStart_(_dafny.UnicodeSeqOfUtf8Bytes("wrong-turn"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), _dafny.UnicodeSeqOfUtf8Bytes("bash"), _dafny.UnicodeSeqOfUtf8Bytes("Run Command"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(_3_act1) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(func(_pat_let377_0 ChatState) ChatState {
		return func(_6_dt__update__tmp_h2 ChatState) ChatState {
			return func(_pat_let378_0 m_AhpSkeleton.Option) ChatState {
				return func(_7_dt__update_hactiveTurn_h1 m_AhpSkeleton.Option) ChatState {
					return Companion_ChatState_.Create_ChatState_((_6_dt__update__tmp_h2).Dtor_turns(), (_6_dt__update__tmp_h2).Dtor_title(), (_6_dt__update__tmp_h2).Dtor_status(), (_6_dt__update__tmp_h2).Dtor_modifiedAt(), (_6_dt__update__tmp_h2).Dtor_draft(), (_6_dt__update__tmp_h2).Dtor_activity(), _7_dt__update_hactiveTurn_h1, (_6_dt__update__tmp_h2).Dtor_steeringMessage(), (_6_dt__update__tmp_h2).Dtor_queuedMessages(), (_6_dt__update__tmp_h2).Dtor_nextCursor())
				}(_pat_let378_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Turn_.Create_Turn_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("T")), _dafny.SeqOf(), _dafny.UnicodeSeqOfUtf8Bytes("in-progress"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())))
		}(_pat_let377_0)
	}(_0_s), Companion_ChatAction_.Create_CTurnComplete_(_dafny.UnicodeSeqOfUtf8Bytes("t1")))).Equals(func(_pat_let379_0 ChatState) ChatState {
		return func(_8_dt__update__tmp_h3 ChatState) ChatState {
			return func(_pat_let380_0 _dafny.Sequence) ChatState {
				return func(_9_dt__update_hturns_h0 _dafny.Sequence) ChatState {
					return func(_pat_let381_0 m_AhpSkeleton.Option) ChatState {
						return func(_10_dt__update_hactiveTurn_h2 m_AhpSkeleton.Option) ChatState {
							return Companion_ChatState_.Create_ChatState_(_9_dt__update_hturns_h0, (_8_dt__update__tmp_h3).Dtor_title(), (_8_dt__update__tmp_h3).Dtor_status(), (_8_dt__update__tmp_h3).Dtor_modifiedAt(), (_8_dt__update__tmp_h3).Dtor_draft(), (_8_dt__update__tmp_h3).Dtor_activity(), _10_dt__update_hactiveTurn_h2, (_8_dt__update__tmp_h3).Dtor_steeringMessage(), (_8_dt__update__tmp_h3).Dtor_queuedMessages(), (_8_dt__update__tmp_h3).Dtor_nextCursor())
						}(_pat_let381_0)
					}(m_AhpSkeleton.Companion_Option_.Create_None_())
				}(_pat_let380_0)
			}(_dafny.SeqOf(Companion_Turn_.Create_Turn_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("T")), _dafny.SeqOf(), _dafny.UnicodeSeqOfUtf8Bytes("complete"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())))
		}(_pat_let379_0)
	}(_0_s)) {
		pass = (pass).Plus(_dafny.One)
	}
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let382_0 ChatState) ChatState {
		return func(_11_dt__update__tmp_h4 ChatState) ChatState {
			return func(_pat_let383_0 m_AhpSkeleton.Option) ChatState {
				return func(_12_dt__update_hactiveTurn_h3 m_AhpSkeleton.Option) ChatState {
					return Companion_ChatState_.Create_ChatState_((_11_dt__update__tmp_h4).Dtor_turns(), (_11_dt__update__tmp_h4).Dtor_title(), (_11_dt__update__tmp_h4).Dtor_status(), (_11_dt__update__tmp_h4).Dtor_modifiedAt(), (_11_dt__update__tmp_h4).Dtor_draft(), (_11_dt__update__tmp_h4).Dtor_activity(), _12_dt__update_hactiveTurn_h3, (_11_dt__update__tmp_h4).Dtor_steeringMessage(), (_11_dt__update__tmp_h4).Dtor_queuedMessages(), (_11_dt__update__tmp_h4).Dtor_nextCursor())
				}(_pat_let383_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))
		}(_pat_let382_0)
	}(_0_s), Companion_ChatAction_.Create_CTurnCancelled_(_dafny.UnicodeSeqOfUtf8Bytes("t1")))).Dtor_turns(), _dafny.SeqOf(Companion_Turn_.Create_Turn_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), m_ConfluxCodec.Companion_Json_.Create_JNull_(), _dafny.SeqOf(), _dafny.UnicodeSeqOfUtf8Bytes("cancelled"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()))) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_3_act1, Companion_ChatAction_.Create_CUsage_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("u"))))).Dtor_activeTurn()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let384_0 Turn) Turn {
		return func(_13_dt__update__tmp_h5 Turn) Turn {
			return func(_pat_let385_0 m_AhpSkeleton.Option) Turn {
				return func(_14_dt__update_husage_h0 m_AhpSkeleton.Option) Turn {
					return Companion_Turn_.Create_Turn_((_13_dt__update__tmp_h5).Dtor_turnId(), (_13_dt__update__tmp_h5).Dtor_message(), (_13_dt__update__tmp_h5).Dtor_parts(), (_13_dt__update__tmp_h5).Dtor_state(), _14_dt__update_husage_h0, (_13_dt__update__tmp_h5).Dtor_error())
				}(_pat_let385_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("u"))))
		}(_pat_let384_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_3_act1, Companion_ChatAction_.Create_CError_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("boom"))))).Dtor_turns(), _dafny.SeqOf(Companion_Turn_.Create_Turn_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), m_ConfluxCodec.Companion_Json_.Create_JNull_(), _dafny.SeqOf(), _dafny.UnicodeSeqOfUtf8Bytes("error"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("boom")))))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _15_withMd ChatState
	_ = _15_withMd
	_15_withMd = Companion_Default___.Apply1(_3_act1, Companion_ChatAction_.Create_CResponsePart_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), Companion_Part_.Create_PMarkdown_(_dafny.UnicodeSeqOfUtf8Bytes("md-1"), _dafny.UnicodeSeqOfUtf8Bytes(""))))
	if ((_15_withMd).Dtor_activeTurn()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let386_0 Turn) Turn {
		return func(_16_dt__update__tmp_h6 Turn) Turn {
			return func(_pat_let387_0 _dafny.Sequence) Turn {
				return func(_17_dt__update_hparts_h0 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_16_dt__update__tmp_h6).Dtor_turnId(), (_16_dt__update__tmp_h6).Dtor_message(), _17_dt__update_hparts_h0, (_16_dt__update__tmp_h6).Dtor_state(), (_16_dt__update__tmp_h6).Dtor_usage(), (_16_dt__update__tmp_h6).Dtor_error())
				}(_pat_let387_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PMarkdown_(_dafny.UnicodeSeqOfUtf8Bytes("md-1"), _dafny.UnicodeSeqOfUtf8Bytes(""))))
		}(_pat_let386_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_15_withMd, Companion_ChatAction_.Create_CDelta_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("md-1"), _dafny.UnicodeSeqOfUtf8Bytes("Hello from chat")))).Dtor_activeTurn()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let388_0 Turn) Turn {
		return func(_18_dt__update__tmp_h7 Turn) Turn {
			return func(_pat_let389_0 _dafny.Sequence) Turn {
				return func(_19_dt__update_hparts_h1 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_18_dt__update__tmp_h7).Dtor_turnId(), (_18_dt__update__tmp_h7).Dtor_message(), _19_dt__update_hparts_h1, (_18_dt__update__tmp_h7).Dtor_state(), (_18_dt__update__tmp_h7).Dtor_usage(), (_18_dt__update__tmp_h7).Dtor_error())
				}(_pat_let389_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PMarkdown_(_dafny.UnicodeSeqOfUtf8Bytes("md-1"), _dafny.UnicodeSeqOfUtf8Bytes("Hello from chat"))))
		}(_pat_let388_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(func(_pat_let390_0 ChatState) ChatState {
		return func(_20_dt__update__tmp_h8 ChatState) ChatState {
			return func(_pat_let391_0 m_AhpSkeleton.Option) ChatState {
				return func(_23_dt__update_hactiveTurn_h4 m_AhpSkeleton.Option) ChatState {
					return Companion_ChatState_.Create_ChatState_((_20_dt__update__tmp_h8).Dtor_turns(), (_20_dt__update__tmp_h8).Dtor_title(), (_20_dt__update__tmp_h8).Dtor_status(), (_20_dt__update__tmp_h8).Dtor_modifiedAt(), (_20_dt__update__tmp_h8).Dtor_draft(), (_20_dt__update__tmp_h8).Dtor_activity(), _23_dt__update_hactiveTurn_h4, (_20_dt__update__tmp_h8).Dtor_steeringMessage(), (_20_dt__update__tmp_h8).Dtor_queuedMessages(), (_20_dt__update__tmp_h8).Dtor_nextCursor())
				}(_pat_let391_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let392_0 Turn) Turn {
				return func(_21_dt__update__tmp_h9 Turn) Turn {
					return func(_pat_let393_0 _dafny.Sequence) Turn {
						return func(_22_dt__update_hparts_h2 _dafny.Sequence) Turn {
							return Companion_Turn_.Create_Turn_((_21_dt__update__tmp_h9).Dtor_turnId(), (_21_dt__update__tmp_h9).Dtor_message(), _22_dt__update_hparts_h2, (_21_dt__update__tmp_h9).Dtor_state(), (_21_dt__update__tmp_h9).Dtor_usage(), (_21_dt__update__tmp_h9).Dtor_error())
						}(_pat_let393_0)
					}(_dafny.SeqOf(Companion_Part_.Create_PReasoning_(_dafny.UnicodeSeqOfUtf8Bytes("r-1"), _dafny.UnicodeSeqOfUtf8Bytes("th"))))
				}(_pat_let392_0)
			}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1")))))
		}(_pat_let390_0)
	}(_3_act1), Companion_ChatAction_.Create_CDelta_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("r-1"), _dafny.UnicodeSeqOfUtf8Bytes("inking")))).Dtor_activeTurn()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let394_0 Turn) Turn {
		return func(_24_dt__update__tmp_h10 Turn) Turn {
			return func(_pat_let395_0 _dafny.Sequence) Turn {
				return func(_25_dt__update_hparts_h3 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_24_dt__update__tmp_h10).Dtor_turnId(), (_24_dt__update__tmp_h10).Dtor_message(), _25_dt__update_hparts_h3, (_24_dt__update__tmp_h10).Dtor_state(), (_24_dt__update__tmp_h10).Dtor_usage(), (_24_dt__update__tmp_h10).Dtor_error())
				}(_pat_let395_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PReasoning_(_dafny.UnicodeSeqOfUtf8Bytes("r-1"), _dafny.UnicodeSeqOfUtf8Bytes("th"))))
		}(_pat_let394_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass
}
func (_static *CompanionStruct_Default___) RunToolCalls() _dafny.Int {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	pass = _dafny.Zero
	var _0_s ChatState
	_ = _0_s
	_0_s = Companion_Default___.C0()
	var _1_act1 ChatState
	_ = _1_act1
	var _2_dt__update__tmp_h0 ChatState = _0_s
	_ = _2_dt__update__tmp_h0
	var _3_dt__update_hactiveTurn_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1")))
	_ = _3_dt__update_hactiveTurn_h0
	_1_act1 = Companion_ChatState_.Create_ChatState_((_2_dt__update__tmp_h0).Dtor_turns(), (_2_dt__update__tmp_h0).Dtor_title(), (_2_dt__update__tmp_h0).Dtor_status(), (_2_dt__update__tmp_h0).Dtor_modifiedAt(), (_2_dt__update__tmp_h0).Dtor_draft(), (_2_dt__update__tmp_h0).Dtor_activity(), _3_dt__update_hactiveTurn_h0, (_2_dt__update__tmp_h0).Dtor_steeringMessage(), (_2_dt__update__tmp_h0).Dtor_queuedMessages(), (_2_dt__update__tmp_h0).Dtor_nextCursor())
	var _4_started ChatState
	_ = _4_started
	_4_started = Companion_Default___.Apply1(_1_act1, Companion_ChatAction_.Create_CToolCallStart_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), _dafny.UnicodeSeqOfUtf8Bytes("bash"), _dafny.UnicodeSeqOfUtf8Bytes("Run Command"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("List the files in the current directory")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()))
	if ((_4_started).Dtor_activeTurn()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let396_0 Turn) Turn {
		return func(_5_dt__update__tmp_h1 Turn) Turn {
			return func(_pat_let397_0 _dafny.Sequence) Turn {
				return func(_6_dt__update_hparts_h0 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_5_dt__update__tmp_h1).Dtor_turnId(), (_5_dt__update__tmp_h1).Dtor_message(), _6_dt__update_hparts_h0, (_5_dt__update__tmp_h1).Dtor_state(), (_5_dt__update__tmp_h1).Dtor_usage(), (_5_dt__update__tmp_h1).Dtor_error())
				}(_pat_let397_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PToolCall_(Companion_Default___.TC0(_dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("List the files in the current directory"))))))
		}(_pat_let396_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _7_afterDelta ChatState
	_ = _7_afterDelta
	_7_afterDelta = Companion_Default___.Apply1(_4_started, Companion_ChatAction_.Create_CToolCallDelta_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), _dafny.UnicodeSeqOfUtf8Bytes("ls -la"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Listing files")), m_AhpSkeleton.Companion_Option_.Create_None_()))
	if _dafny.Companion_Sequence_.Equal(Companion_Default___.ActiveParts(_7_afterDelta), _dafny.SeqOf(Companion_Part_.Create_PToolCall_(func(_pat_let398_0 ToolCall) ToolCall {
		return func(_8_dt__update__tmp_h2 ToolCall) ToolCall {
			return func(_pat_let399_0 _dafny.Sequence) ToolCall {
				return func(_9_dt__update_hinvocationMessage_h0 _dafny.Sequence) ToolCall {
					return func(_pat_let400_0 m_AhpSkeleton.Option) ToolCall {
						return func(_10_dt__update_hpartialInput_h0 m_AhpSkeleton.Option) ToolCall {
							return Companion_ToolCall_.Create_ToolCall_((_8_dt__update__tmp_h2).Dtor_toolCallId(), (_8_dt__update__tmp_h2).Dtor_toolName(), (_8_dt__update__tmp_h2).Dtor_displayName(), (_8_dt__update__tmp_h2).Dtor_status(), (_8_dt__update__tmp_h2).Dtor_intention(), (_8_dt__update__tmp_h2).Dtor_contributor(), (_8_dt__update__tmp_h2).Dtor_meta(), _9_dt__update_hinvocationMessage_h0, (_8_dt__update__tmp_h2).Dtor_toolInput(), (_8_dt__update__tmp_h2).Dtor_confirmationTitle(), (_8_dt__update__tmp_h2).Dtor_riskAssessment(), (_8_dt__update__tmp_h2).Dtor_edits(), (_8_dt__update__tmp_h2).Dtor_editable(), (_8_dt__update__tmp_h2).Dtor_options(), (_8_dt__update__tmp_h2).Dtor_confirmed(), (_8_dt__update__tmp_h2).Dtor_success(), (_8_dt__update__tmp_h2).Dtor_pastTenseMessage(), (_8_dt__update__tmp_h2).Dtor_reason(), (_8_dt__update__tmp_h2).Dtor_reasonMessage(), (_8_dt__update__tmp_h2).Dtor_userSuggestion(), (_8_dt__update__tmp_h2).Dtor_selectedOption(), (_8_dt__update__tmp_h2).Dtor_content(), (_8_dt__update__tmp_h2).Dtor_structuredContent(), (_8_dt__update__tmp_h2).Dtor_error(), (_8_dt__update__tmp_h2).Dtor_auth(), _10_dt__update_hpartialInput_h0)
						}(_pat_let400_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("ls -la")))
				}(_pat_let399_0)
			}(_dafny.UnicodeSeqOfUtf8Bytes("Listing files"))
		}(_pat_let398_0)
	}(Companion_Default___.TC0(_dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("List the files in the current directory"))))))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _11_afterReady ChatState
	_ = _11_afterReady
	_11_afterReady = Companion_Default___.Apply1(_7_afterDelta, Companion_ChatAction_.Create_CToolCallReady_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Run: ls -la")), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("ls -la")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()))
	var _12_tcReady m_AhpSkeleton.Option
	_ = _12_tcReady
	_12_tcReady = Companion_Default___.FirstActiveToolCall(_11_afterReady)
	if (((((_12_tcReady).Is_Some()) && (_dafny.Companion_Sequence_.Equal(((_12_tcReady).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("pending-confirmation")))) && (_dafny.Companion_Sequence_.Equal(((_12_tcReady).Dtor_value().(ToolCall)).Dtor_invocationMessage(), _dafny.UnicodeSeqOfUtf8Bytes("Run: ls -la")))) && ((((_12_tcReady).Dtor_value().(ToolCall)).Dtor_toolInput()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("ls -la"))))) && ((((_12_tcReady).Dtor_value().(ToolCall)).Dtor_partialInput()).Equals(m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	var _13_afterConf ChatState
	_ = _13_afterConf
	_13_afterConf = Companion_Default___.Apply1(_11_afterReady, Companion_ChatAction_.Create_CToolCallConfirmed_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("user-action")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()))
	var _14_tcConf m_AhpSkeleton.Option
	_ = _14_tcConf
	_14_tcConf = Companion_Default___.FirstActiveToolCall(_13_afterConf)
	if (((_14_tcConf).Is_Some()) && (_dafny.Companion_Sequence_.Equal(((_14_tcConf).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("running")))) && ((((_14_tcConf).Dtor_value().(ToolCall)).Dtor_confirmed()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("user-action")))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _15_afterComp ChatState
	_ = _15_afterComp
	_15_afterComp = Companion_Default___.Apply1(_13_afterConf, Companion_ChatAction_.Create_CToolCallComplete_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Ran command")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), false, m_AhpSkeleton.Companion_Option_.Create_None_()))
	var _16_tcComp m_AhpSkeleton.Option
	_ = _16_tcComp
	_16_tcComp = Companion_Default___.FirstActiveToolCall(_15_afterComp)
	if ((((_16_tcComp).Is_Some()) && (_dafny.Companion_Sequence_.Equal(((_16_tcComp).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("completed")))) && ((((_16_tcComp).Dtor_value().(ToolCall)).Dtor_success()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(true)))) && ((((_16_tcComp).Dtor_value().(ToolCall)).Dtor_pastTenseMessage()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Ran command")))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _17_life ChatState
	_ = _17_life
	_17_life = Companion_Default___.Fold(_1_act1, _dafny.SeqOf(Companion_ChatAction_.Create_CToolCallStart_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), _dafny.UnicodeSeqOfUtf8Bytes("bash"), _dafny.UnicodeSeqOfUtf8Bytes("Run Command"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("List the files in the current directory")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallDelta_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), _dafny.UnicodeSeqOfUtf8Bytes("ls -la"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Listing files")), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallReady_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Run: ls -la")), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("ls -la")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallConfirmed_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("user-action")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallComplete_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Ran command")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), false, m_AhpSkeleton.Companion_Option_.Create_None_())))
	var _18_lifeTc m_AhpSkeleton.Option
	_ = _18_lifeTc
	_18_lifeTc = Companion_Default___.FirstActiveToolCall(_17_life)
	if (((((_18_lifeTc).Is_Some()) && (_dafny.Companion_Sequence_.Equal(((_18_lifeTc).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("completed")))) && ((((_18_lifeTc).Dtor_value().(ToolCall)).Dtor_confirmed()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("user-action"))))) && ((((_18_lifeTc).Dtor_value().(ToolCall)).Dtor_success()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(true)))) && ((((_18_lifeTc).Dtor_value().(ToolCall)).Dtor_toolInput()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("ls -la")))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _19_auto ChatState
	_ = _19_auto
	_19_auto = Companion_Default___.Fold(_1_act1, _dafny.SeqOf(Companion_ChatAction_.Create_CToolCallStart_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), _dafny.UnicodeSeqOfUtf8Bytes("bash"), _dafny.UnicodeSeqOfUtf8Bytes("Run Command"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallReady_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Run")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("not-needed")), m_AhpSkeleton.Companion_Option_.Create_None_())))
	var _20_autoTc m_AhpSkeleton.Option
	_ = _20_autoTc
	_20_autoTc = Companion_Default___.FirstActiveToolCall(_19_auto)
	if ((((_20_autoTc).Is_Some()) && (_dafny.Companion_Sequence_.Equal(((_20_autoTc).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("running")))) && ((((_20_autoTc).Dtor_value().(ToolCall)).Dtor_confirmed()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("not-needed"))))) && (_dafny.Companion_Sequence_.Equal(((_20_autoTc).Dtor_value().(ToolCall)).Dtor_invocationMessage(), _dafny.UnicodeSeqOfUtf8Bytes("Run"))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _21_pendingRisk m_ConfluxCodec.Json
	_ = _21_pendingRisk
	_21_pendingRisk = m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("status"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("loading"))))
	var _22_pendingEdits m_ConfluxCodec.Json
	_ = _22_pendingEdits
	_22_pendingEdits = m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("items"), m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf())))
	var _23_pendingOption m_ConfluxCodec.Json
	_ = _23_pendingOption
	_23_pendingOption = m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("id"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("approve"))))
	var _pat_let_tv0 = _23_pendingOption
	_ = _pat_let_tv0
	var _pat_let_tv1 = _22_pendingEdits
	_ = _pat_let_tv1
	var _pat_let_tv2 = _21_pendingRisk
	_ = _pat_let_tv2
	var _24_pendState ChatState
	_ = _24_pendState
	var _25_dt__update__tmp_h3 ChatState = _1_act1
	_ = _25_dt__update__tmp_h3
	var _26_dt__update_hstatus_h0 uint32 = uint32(24)
	_ = _26_dt__update_hstatus_h0
	var _27_dt__update_hactiveTurn_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let401_0 Turn) Turn {
		return func(_28_dt__update__tmp_h4 Turn) Turn {
			return func(_pat_let402_0 _dafny.Sequence) Turn {
				return func(_37_dt__update_hparts_h1 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_28_dt__update__tmp_h4).Dtor_turnId(), (_28_dt__update__tmp_h4).Dtor_message(), _37_dt__update_hparts_h1, (_28_dt__update__tmp_h4).Dtor_state(), (_28_dt__update__tmp_h4).Dtor_usage(), (_28_dt__update__tmp_h4).Dtor_error())
				}(_pat_let402_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PToolCall_(func(_pat_let403_0 ToolCall) ToolCall {
				return func(_29_dt__update__tmp_h5 ToolCall) ToolCall {
					return func(_pat_let404_0 m_AhpSkeleton.Option) ToolCall {
						return func(_30_dt__update_hoptions_h0 m_AhpSkeleton.Option) ToolCall {
							return func(_pat_let405_0 m_AhpSkeleton.Option) ToolCall {
								return func(_31_dt__update_heditable_h0 m_AhpSkeleton.Option) ToolCall {
									return func(_pat_let406_0 m_AhpSkeleton.Option) ToolCall {
										return func(_32_dt__update_hedits_h0 m_AhpSkeleton.Option) ToolCall {
											return func(_pat_let407_0 m_AhpSkeleton.Option) ToolCall {
												return func(_33_dt__update_hriskAssessment_h0 m_AhpSkeleton.Option) ToolCall {
													return func(_pat_let408_0 m_AhpSkeleton.Option) ToolCall {
														return func(_34_dt__update_hconfirmationTitle_h0 m_AhpSkeleton.Option) ToolCall {
															return func(_pat_let409_0 m_AhpSkeleton.Option) ToolCall {
																return func(_35_dt__update_htoolInput_h0 m_AhpSkeleton.Option) ToolCall {
																	return func(_pat_let410_0 _dafny.Sequence) ToolCall {
																		return func(_36_dt__update_hstatus_h1 _dafny.Sequence) ToolCall {
																			return Companion_ToolCall_.Create_ToolCall_((_29_dt__update__tmp_h5).Dtor_toolCallId(), (_29_dt__update__tmp_h5).Dtor_toolName(), (_29_dt__update__tmp_h5).Dtor_displayName(), _36_dt__update_hstatus_h1, (_29_dt__update__tmp_h5).Dtor_intention(), (_29_dt__update__tmp_h5).Dtor_contributor(), (_29_dt__update__tmp_h5).Dtor_meta(), (_29_dt__update__tmp_h5).Dtor_invocationMessage(), _35_dt__update_htoolInput_h0, _34_dt__update_hconfirmationTitle_h0, _33_dt__update_hriskAssessment_h0, _32_dt__update_hedits_h0, _31_dt__update_heditable_h0, _30_dt__update_hoptions_h0, (_29_dt__update__tmp_h5).Dtor_confirmed(), (_29_dt__update__tmp_h5).Dtor_success(), (_29_dt__update__tmp_h5).Dtor_pastTenseMessage(), (_29_dt__update__tmp_h5).Dtor_reason(), (_29_dt__update__tmp_h5).Dtor_reasonMessage(), (_29_dt__update__tmp_h5).Dtor_userSuggestion(), (_29_dt__update__tmp_h5).Dtor_selectedOption(), (_29_dt__update__tmp_h5).Dtor_content(), (_29_dt__update__tmp_h5).Dtor_structuredContent(), (_29_dt__update__tmp_h5).Dtor_error(), (_29_dt__update__tmp_h5).Dtor_auth(), (_29_dt__update__tmp_h5).Dtor_partialInput())
																		}(_pat_let410_0)
																	}(_dafny.UnicodeSeqOfUtf8Bytes("pending-confirmation"))
																}(_pat_let409_0)
															}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("rm -rf /tmp/test")))
														}(_pat_let408_0)
													}(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("Run in terminal"))))
												}(_pat_let407_0)
											}(m_AhpSkeleton.Companion_Option_.Create_Some_(_pat_let_tv2))
										}(_pat_let406_0)
									}(m_AhpSkeleton.Companion_Option_.Create_Some_(_pat_let_tv1))
								}(_pat_let405_0)
							}(m_AhpSkeleton.Companion_Option_.Create_Some_(true))
						}(_pat_let404_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_pat_let_tv0)))
				}(_pat_let403_0)
			}(Companion_Default___.TC0(_dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_None_())))))
		}(_pat_let401_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))
	_ = _27_dt__update_hactiveTurn_h1
	_24_pendState = Companion_ChatState_.Create_ChatState_((_25_dt__update__tmp_h3).Dtor_turns(), (_25_dt__update__tmp_h3).Dtor_title(), _26_dt__update_hstatus_h0, (_25_dt__update__tmp_h3).Dtor_modifiedAt(), (_25_dt__update__tmp_h3).Dtor_draft(), (_25_dt__update__tmp_h3).Dtor_activity(), _27_dt__update_hactiveTurn_h1, (_25_dt__update__tmp_h3).Dtor_steeringMessage(), (_25_dt__update__tmp_h3).Dtor_queuedMessages(), (_25_dt__update__tmp_h3).Dtor_nextCursor())
	var _38_refreshed m_AhpSkeleton.Option
	_ = _38_refreshed
	_38_refreshed = Companion_Default___.FirstActiveToolCall(Companion_Default___.Apply1(_24_pendState, Companion_ChatAction_.Create_CToolCallReady_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Run again")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())))
	if (((((((((_38_refreshed).Is_Some()) && (_dafny.Companion_Sequence_.Equal(((_38_refreshed).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("pending-confirmation")))) && (_dafny.Companion_Sequence_.Equal(((_38_refreshed).Dtor_value().(ToolCall)).Dtor_invocationMessage(), _dafny.UnicodeSeqOfUtf8Bytes("Run again")))) && ((((_38_refreshed).Dtor_value().(ToolCall)).Dtor_toolInput()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("rm -rf /tmp/test"))))) && ((((_38_refreshed).Dtor_value().(ToolCall)).Dtor_confirmationTitle()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("Run in terminal")))))) && ((((_38_refreshed).Dtor_value().(ToolCall)).Dtor_riskAssessment()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_21_pendingRisk)))) && ((((_38_refreshed).Dtor_value().(ToolCall)).Dtor_edits()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_22_pendingEdits)))) && ((((_38_refreshed).Dtor_value().(ToolCall)).Dtor_editable()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(true)))) && ((((_38_refreshed).Dtor_value().(ToolCall)).Dtor_options()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_23_pendingOption)))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _39_forced ChatState
	_ = _39_forced
	_39_forced = Companion_Default___.Apply1(_4_started, Companion_ChatAction_.Create_CTurnComplete_(_dafny.UnicodeSeqOfUtf8Bytes("t1")))
	var _40_fTc m_AhpSkeleton.Option
	_ = _40_fTc
	_40_fTc = Companion_Default___.FirstStoredToolCall(_39_forced)
	if (((((((_39_forced).Dtor_activeTurn()).Equals(m_AhpSkeleton.Companion_Option_.Create_None_())) && ((_dafny.IntOfUint32(((_39_forced).Dtor_turns()).Cardinality())).Sign() == 1)) && (_dafny.Companion_Sequence_.Equal((((_39_forced).Dtor_turns()).Select(0).(Turn)).Dtor_state(), _dafny.UnicodeSeqOfUtf8Bytes("complete")))) && ((_40_fTc).Is_Some())) && (_dafny.Companion_Sequence_.Equal(((_40_fTc).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("cancelled")))) && ((((_40_fTc).Dtor_value().(ToolCall)).Dtor_reason()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("skipped")))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _41_ignoredComplete ChatState
	_ = _41_ignoredComplete
	_41_ignoredComplete = Companion_Default___.Apply1(_4_started, Companion_ChatAction_.Create_CToolCallComplete_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), false, m_AhpSkeleton.Companion_Option_.Create_None_()))
	if _dafny.Companion_Sequence_.Equal(Companion_Default___.ActiveParts(_41_ignoredComplete), Companion_Default___.ActiveParts(_4_started)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _42_authBase ChatState
	_ = _42_authBase
	var _43_dt__update__tmp_h6 ChatState = _1_act1
	_ = _43_dt__update__tmp_h6
	var _44_dt__update_hstatus_h2 uint32 = uint32(8)
	_ = _44_dt__update_hstatus_h2
	var _45_dt__update_hactiveTurn_h2 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let411_0 Turn) Turn {
		return func(_46_dt__update__tmp_h7 Turn) Turn {
			return func(_pat_let412_0 _dafny.Sequence) Turn {
				return func(_53_dt__update_hparts_h2 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_46_dt__update__tmp_h7).Dtor_turnId(), (_46_dt__update__tmp_h7).Dtor_message(), _53_dt__update_hparts_h2, (_46_dt__update__tmp_h7).Dtor_state(), (_46_dt__update__tmp_h7).Dtor_usage(), (_46_dt__update__tmp_h7).Dtor_error())
				}(_pat_let412_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PToolCall_(func(_pat_let413_0 ToolCall) ToolCall {
				return func(_47_dt__update__tmp_h8 ToolCall) ToolCall {
					return func(_pat_let414_0 m_AhpSkeleton.Option) ToolCall {
						return func(_48_dt__update_hcontent_h0 m_AhpSkeleton.Option) ToolCall {
							return func(_pat_let415_0 m_AhpSkeleton.Option) ToolCall {
								return func(_49_dt__update_hconfirmed_h0 m_AhpSkeleton.Option) ToolCall {
									return func(_pat_let416_0 m_AhpSkeleton.Option) ToolCall {
										return func(_50_dt__update_htoolInput_h1 m_AhpSkeleton.Option) ToolCall {
											return func(_pat_let417_0 m_AhpSkeleton.Option) ToolCall {
												return func(_51_dt__update_hcontributor_h0 m_AhpSkeleton.Option) ToolCall {
													return func(_pat_let418_0 _dafny.Sequence) ToolCall {
														return func(_52_dt__update_hstatus_h3 _dafny.Sequence) ToolCall {
															return Companion_ToolCall_.Create_ToolCall_((_47_dt__update__tmp_h8).Dtor_toolCallId(), (_47_dt__update__tmp_h8).Dtor_toolName(), (_47_dt__update__tmp_h8).Dtor_displayName(), _52_dt__update_hstatus_h3, (_47_dt__update__tmp_h8).Dtor_intention(), _51_dt__update_hcontributor_h0, (_47_dt__update__tmp_h8).Dtor_meta(), (_47_dt__update__tmp_h8).Dtor_invocationMessage(), _50_dt__update_htoolInput_h1, (_47_dt__update__tmp_h8).Dtor_confirmationTitle(), (_47_dt__update__tmp_h8).Dtor_riskAssessment(), (_47_dt__update__tmp_h8).Dtor_edits(), (_47_dt__update__tmp_h8).Dtor_editable(), (_47_dt__update__tmp_h8).Dtor_options(), _49_dt__update_hconfirmed_h0, (_47_dt__update__tmp_h8).Dtor_success(), (_47_dt__update__tmp_h8).Dtor_pastTenseMessage(), (_47_dt__update__tmp_h8).Dtor_reason(), (_47_dt__update__tmp_h8).Dtor_reasonMessage(), (_47_dt__update__tmp_h8).Dtor_userSuggestion(), (_47_dt__update__tmp_h8).Dtor_selectedOption(), _48_dt__update_hcontent_h0, (_47_dt__update__tmp_h8).Dtor_structuredContent(), (_47_dt__update__tmp_h8).Dtor_error(), (_47_dt__update__tmp_h8).Dtor_auth(), (_47_dt__update__tmp_h8).Dtor_partialInput())
														}(_pat_let418_0)
													}(_dafny.UnicodeSeqOfUtf8Bytes("running"))
												}(_pat_let417_0)
											}(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_ToolCallContributor_.Create_ToolCallContributor_(_dafny.UnicodeSeqOfUtf8Bytes("mcp"), m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("mcp")))))))
										}(_pat_let416_0)
									}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("foo")))
								}(_pat_let415_0)
							}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("user-action")))
						}(_pat_let414_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("partial"))))
				}(_pat_let413_0)
			}(Companion_Default___.TC0(_dafny.UnicodeSeqOfUtf8Bytes("tc-auth"), m_AhpSkeleton.Companion_Option_.Create_None_())))))
		}(_pat_let411_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))
	_ = _45_dt__update_hactiveTurn_h2
	_42_authBase = Companion_ChatState_.Create_ChatState_((_43_dt__update__tmp_h6).Dtor_turns(), (_43_dt__update__tmp_h6).Dtor_title(), _44_dt__update_hstatus_h2, (_43_dt__update__tmp_h6).Dtor_modifiedAt(), (_43_dt__update__tmp_h6).Dtor_draft(), (_43_dt__update__tmp_h6).Dtor_activity(), _45_dt__update_hactiveTurn_h2, (_43_dt__update__tmp_h6).Dtor_steeringMessage(), (_43_dt__update__tmp_h6).Dtor_queuedMessages(), (_43_dt__update__tmp_h6).Dtor_nextCursor())
	var _54_challenge m_ConfluxCodec.Json
	_ = _54_challenge
	_54_challenge = m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("reason"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("required"))))
	var _55_authState ChatState
	_ = _55_authState
	_55_authState = Companion_Default___.Apply1(_42_authBase, Companion_ChatAction_.Create_CToolCallAuthRequired_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-auth"), _54_challenge, m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("auth-meta")))))
	var _56_authTc m_AhpSkeleton.Option
	_ = _56_authTc
	_56_authTc = Companion_Default___.FirstActiveToolCall(_55_authState)
	if (((((((((_55_authState).Dtor_status()) == (uint32(24))) && ((_56_authTc).Is_Some())) && (_dafny.Companion_Sequence_.Equal(((_56_authTc).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("auth-required")))) && ((((_56_authTc).Dtor_value().(ToolCall)).Dtor_auth()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_54_challenge)))) && ((((_56_authTc).Dtor_value().(ToolCall)).Dtor_meta()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("auth-meta")))))) && ((((_56_authTc).Dtor_value().(ToolCall)).Dtor_toolInput()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("foo"))))) && ((((_56_authTc).Dtor_value().(ToolCall)).Dtor_confirmed()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("user-action"))))) && ((((_56_authTc).Dtor_value().(ToolCall)).Dtor_content()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("partial"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _57_resolvedState ChatState
	_ = _57_resolvedState
	_57_resolvedState = Companion_Default___.Apply1(_55_authState, Companion_ChatAction_.Create_CToolCallAuthResolved_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-auth"), m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("resolved-meta")))))
	var _58_resolvedTc m_AhpSkeleton.Option
	_ = _58_resolvedTc
	_58_resolvedTc = Companion_Default___.FirstActiveToolCall(_57_resolvedState)
	if ((((((((((_57_resolvedState).Dtor_status()) == (uint32(8))) && ((_58_resolvedTc).Is_Some())) && (_dafny.Companion_Sequence_.Equal(((_58_resolvedTc).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("running")))) && ((((_58_resolvedTc).Dtor_value().(ToolCall)).Dtor_auth()).Equals(m_AhpSkeleton.Companion_Option_.Create_None_()))) && ((((_58_resolvedTc).Dtor_value().(ToolCall)).Dtor_contributor()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_ToolCallContributor_.Create_ToolCallContributor_(_dafny.UnicodeSeqOfUtf8Bytes("mcp"), m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("mcp"))))))))) && ((((_58_resolvedTc).Dtor_value().(ToolCall)).Dtor_toolInput()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("foo"))))) && ((((_58_resolvedTc).Dtor_value().(ToolCall)).Dtor_confirmed()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("user-action"))))) && ((((_58_resolvedTc).Dtor_value().(ToolCall)).Dtor_content()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("partial")))))) && ((((_58_resolvedTc).Dtor_value().(ToolCall)).Dtor_meta()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("resolved-meta"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _59_clientBase ChatState
	_ = _59_clientBase
	var _60_dt__update__tmp_h9 ChatState = _42_authBase
	_ = _60_dt__update__tmp_h9
	var _61_dt__update_hactiveTurn_h3 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let419_0 Turn) Turn {
		return func(_62_dt__update__tmp_h10 Turn) Turn {
			return func(_pat_let420_0 _dafny.Sequence) Turn {
				return func(_67_dt__update_hparts_h3 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_62_dt__update__tmp_h10).Dtor_turnId(), (_62_dt__update__tmp_h10).Dtor_message(), _67_dt__update_hparts_h3, (_62_dt__update__tmp_h10).Dtor_state(), (_62_dt__update__tmp_h10).Dtor_usage(), (_62_dt__update__tmp_h10).Dtor_error())
				}(_pat_let420_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PToolCall_(func(_pat_let421_0 ToolCall) ToolCall {
				return func(_63_dt__update__tmp_h11 ToolCall) ToolCall {
					return func(_pat_let422_0 m_AhpSkeleton.Option) ToolCall {
						return func(_64_dt__update_hconfirmed_h1 m_AhpSkeleton.Option) ToolCall {
							return func(_pat_let423_0 m_AhpSkeleton.Option) ToolCall {
								return func(_65_dt__update_hcontributor_h1 m_AhpSkeleton.Option) ToolCall {
									return func(_pat_let424_0 _dafny.Sequence) ToolCall {
										return func(_66_dt__update_hstatus_h4 _dafny.Sequence) ToolCall {
											return Companion_ToolCall_.Create_ToolCall_((_63_dt__update__tmp_h11).Dtor_toolCallId(), (_63_dt__update__tmp_h11).Dtor_toolName(), (_63_dt__update__tmp_h11).Dtor_displayName(), _66_dt__update_hstatus_h4, (_63_dt__update__tmp_h11).Dtor_intention(), _65_dt__update_hcontributor_h1, (_63_dt__update__tmp_h11).Dtor_meta(), (_63_dt__update__tmp_h11).Dtor_invocationMessage(), (_63_dt__update__tmp_h11).Dtor_toolInput(), (_63_dt__update__tmp_h11).Dtor_confirmationTitle(), (_63_dt__update__tmp_h11).Dtor_riskAssessment(), (_63_dt__update__tmp_h11).Dtor_edits(), (_63_dt__update__tmp_h11).Dtor_editable(), (_63_dt__update__tmp_h11).Dtor_options(), _64_dt__update_hconfirmed_h1, (_63_dt__update__tmp_h11).Dtor_success(), (_63_dt__update__tmp_h11).Dtor_pastTenseMessage(), (_63_dt__update__tmp_h11).Dtor_reason(), (_63_dt__update__tmp_h11).Dtor_reasonMessage(), (_63_dt__update__tmp_h11).Dtor_userSuggestion(), (_63_dt__update__tmp_h11).Dtor_selectedOption(), (_63_dt__update__tmp_h11).Dtor_content(), (_63_dt__update__tmp_h11).Dtor_structuredContent(), (_63_dt__update__tmp_h11).Dtor_error(), (_63_dt__update__tmp_h11).Dtor_auth(), (_63_dt__update__tmp_h11).Dtor_partialInput())
										}(_pat_let424_0)
									}(_dafny.UnicodeSeqOfUtf8Bytes("running"))
								}(_pat_let423_0)
							}(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_ToolCallContributor_.Create_ToolCallContributor_(_dafny.UnicodeSeqOfUtf8Bytes("client"), m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("client")))))))
						}(_pat_let422_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("not-needed")))
				}(_pat_let421_0)
			}(Companion_Default___.TC0(_dafny.UnicodeSeqOfUtf8Bytes("tc-client"), m_AhpSkeleton.Companion_Option_.Create_None_())))))
		}(_pat_let419_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))
	_ = _61_dt__update_hactiveTurn_h3
	_59_clientBase = Companion_ChatState_.Create_ChatState_((_60_dt__update__tmp_h9).Dtor_turns(), (_60_dt__update__tmp_h9).Dtor_title(), (_60_dt__update__tmp_h9).Dtor_status(), (_60_dt__update__tmp_h9).Dtor_modifiedAt(), (_60_dt__update__tmp_h9).Dtor_draft(), (_60_dt__update__tmp_h9).Dtor_activity(), _61_dt__update_hactiveTurn_h3, (_60_dt__update__tmp_h9).Dtor_steeringMessage(), (_60_dt__update__tmp_h9).Dtor_queuedMessages(), (_60_dt__update__tmp_h9).Dtor_nextCursor())
	if (Companion_Default___.Apply1(_59_clientBase, Companion_ChatAction_.Create_CToolCallAuthRequired_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-client"), _54_challenge, m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(_59_clientBase) {
		pass = (pass).Plus(_dafny.One)
	}
	var _68_authFailed ChatState
	_ = _68_authFailed
	_68_authFailed = Companion_Default___.Apply1(_55_authState, Companion_ChatAction_.Create_CToolCallComplete_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-auth"), false, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Cancelled search")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("cancelled"))), true, m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("complete-meta")))))
	var _69_failedTc m_AhpSkeleton.Option
	_ = _69_failedTc
	_69_failedTc = Companion_Default___.FirstActiveToolCall(_68_authFailed)
	if (((((((_68_authFailed).Dtor_status()) == (uint32(8))) && ((_69_failedTc).Is_Some())) && (_dafny.Companion_Sequence_.Equal(((_69_failedTc).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("completed")))) && ((((_69_failedTc).Dtor_value().(ToolCall)).Dtor_success()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(false)))) && ((((_69_failedTc).Dtor_value().(ToolCall)).Dtor_pastTenseMessage()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Cancelled search"))))) && ((((_69_failedTc).Dtor_value().(ToolCall)).Dtor_content()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("partial"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_55_authState, Companion_ChatAction_.Create_CToolCallComplete_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-auth"), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Should not apply")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), false, m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(_55_authState) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass
}
func (_static *CompanionStruct_Default___) RunPendingHistory() _dafny.Int {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	pass = _dafny.Zero
	var _0_s ChatState
	_ = _0_s
	_0_s = Companion_Default___.C0()
	var _1_act1 ChatState
	_ = _1_act1
	var _2_dt__update__tmp_h0 ChatState = _0_s
	_ = _2_dt__update__tmp_h0
	var _3_dt__update_hactiveTurn_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1")))
	_ = _3_dt__update_hactiveTurn_h0
	_1_act1 = Companion_ChatState_.Create_ChatState_((_2_dt__update__tmp_h0).Dtor_turns(), (_2_dt__update__tmp_h0).Dtor_title(), (_2_dt__update__tmp_h0).Dtor_status(), (_2_dt__update__tmp_h0).Dtor_modifiedAt(), (_2_dt__update__tmp_h0).Dtor_draft(), (_2_dt__update__tmp_h0).Dtor_activity(), _3_dt__update_hactiveTurn_h0, (_2_dt__update__tmp_h0).Dtor_steeringMessage(), (_2_dt__update__tmp_h0).Dtor_queuedMessages(), (_2_dt__update__tmp_h0).Dtor_nextCursor())
	if ((Companion_Default___.Apply1(_0_s, Companion_ChatAction_.Create_CPendingMessageSet_(_dafny.UnicodeSeqOfUtf8Bytes("steering"), _dafny.UnicodeSeqOfUtf8Bytes("sm-1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("Focus on tests"))))).Dtor_steeringMessage()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_QMsg_.Create_QMsg_(_dafny.UnicodeSeqOfUtf8Bytes("sm-1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("Focus on tests"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_0_s, Companion_ChatAction_.Create_CPendingMessageRemoved_(_dafny.UnicodeSeqOfUtf8Bytes("steering"), _dafny.UnicodeSeqOfUtf8Bytes("sm-unknown")))).Equals(_0_s) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(func(_pat_let425_0 ChatState) ChatState {
		return func(_4_dt__update__tmp_h1 ChatState) ChatState {
			return func(_pat_let426_0 m_AhpSkeleton.Option) ChatState {
				return func(_5_dt__update_hsteeringMessage_h0 m_AhpSkeleton.Option) ChatState {
					return Companion_ChatState_.Create_ChatState_((_4_dt__update__tmp_h1).Dtor_turns(), (_4_dt__update__tmp_h1).Dtor_title(), (_4_dt__update__tmp_h1).Dtor_status(), (_4_dt__update__tmp_h1).Dtor_modifiedAt(), (_4_dt__update__tmp_h1).Dtor_draft(), (_4_dt__update__tmp_h1).Dtor_activity(), (_4_dt__update__tmp_h1).Dtor_activeTurn(), _5_dt__update_hsteeringMessage_h0, (_4_dt__update__tmp_h1).Dtor_queuedMessages(), (_4_dt__update__tmp_h1).Dtor_nextCursor())
				}(_pat_let426_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_QMsg_.Create_QMsg_(_dafny.UnicodeSeqOfUtf8Bytes("sm-1"), m_ConfluxCodec.Companion_Json_.Create_JNull_())))
		}(_pat_let425_0)
	}(_0_s), Companion_ChatAction_.Create_CPendingMessageRemoved_(_dafny.UnicodeSeqOfUtf8Bytes("steering"), _dafny.UnicodeSeqOfUtf8Bytes("sm-1")))).Dtor_steeringMessage()).Equals(m_AhpSkeleton.Companion_Option_.Create_None_()) {
		pass = (pass).Plus(_dafny.One)
	}
	var _6_q QMsg
	_ = _6_q
	_6_q = Companion_QMsg_.Create_QMsg_(_dafny.UnicodeSeqOfUtf8Bytes("q-1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("m")))
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_0_s, Companion_ChatAction_.Create_CPendingMessageSet_(_dafny.UnicodeSeqOfUtf8Bytes("queued"), _dafny.UnicodeSeqOfUtf8Bytes("q-1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("m"))))).Dtor_queuedMessages(), _dafny.SeqOf(_6_q)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _pat_let_tv0 = _6_q
	_ = _pat_let_tv0
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let427_0 ChatState) ChatState {
		return func(_7_dt__update__tmp_h2 ChatState) ChatState {
			return func(_pat_let428_0 _dafny.Sequence) ChatState {
				return func(_8_dt__update_hqueuedMessages_h0 _dafny.Sequence) ChatState {
					return Companion_ChatState_.Create_ChatState_((_7_dt__update__tmp_h2).Dtor_turns(), (_7_dt__update__tmp_h2).Dtor_title(), (_7_dt__update__tmp_h2).Dtor_status(), (_7_dt__update__tmp_h2).Dtor_modifiedAt(), (_7_dt__update__tmp_h2).Dtor_draft(), (_7_dt__update__tmp_h2).Dtor_activity(), (_7_dt__update__tmp_h2).Dtor_activeTurn(), (_7_dt__update__tmp_h2).Dtor_steeringMessage(), _8_dt__update_hqueuedMessages_h0, (_7_dt__update__tmp_h2).Dtor_nextCursor())
				}(_pat_let428_0)
			}(_dafny.SeqOf(_pat_let_tv0, Companion_QMsg_.Create_QMsg_(_dafny.UnicodeSeqOfUtf8Bytes("q-2"), m_ConfluxCodec.Companion_Json_.Create_JNull_())))
		}(_pat_let427_0)
	}(_0_s), Companion_ChatAction_.Create_CPendingMessageRemoved_(_dafny.UnicodeSeqOfUtf8Bytes("queued"), _dafny.UnicodeSeqOfUtf8Bytes("q-1")))).Dtor_queuedMessages(), _dafny.SeqOf(Companion_QMsg_.Create_QMsg_(_dafny.UnicodeSeqOfUtf8Bytes("q-2"), m_ConfluxCodec.Companion_Json_.Create_JNull_()))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _9_three ChatState
	_ = _9_three
	var _10_dt__update__tmp_h3 ChatState = _0_s
	_ = _10_dt__update__tmp_h3
	var _11_dt__update_hturns_h0 _dafny.Sequence = _dafny.SeqOf(Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t1")), Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t2")), Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t3")))
	_ = _11_dt__update_hturns_h0
	_9_three = Companion_ChatState_.Create_ChatState_(_11_dt__update_hturns_h0, (_10_dt__update__tmp_h3).Dtor_title(), (_10_dt__update__tmp_h3).Dtor_status(), (_10_dt__update__tmp_h3).Dtor_modifiedAt(), (_10_dt__update__tmp_h3).Dtor_draft(), (_10_dt__update__tmp_h3).Dtor_activity(), (_10_dt__update__tmp_h3).Dtor_activeTurn(), (_10_dt__update__tmp_h3).Dtor_steeringMessage(), (_10_dt__update__tmp_h3).Dtor_queuedMessages(), (_10_dt__update__tmp_h3).Dtor_nextCursor())
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_9_three, Companion_ChatAction_.Create_CTruncated_(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("t2"))))).Dtor_turns(), _dafny.SeqOf(Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t1")), Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t2")))) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(func(_pat_let429_0 ChatState) ChatState {
		return func(_12_dt__update__tmp_h4 ChatState) ChatState {
			return func(_pat_let430_0 m_AhpSkeleton.Option) ChatState {
				return func(_13_dt__update_hnextCursor_h0 m_AhpSkeleton.Option) ChatState {
					return Companion_ChatState_.Create_ChatState_((_12_dt__update__tmp_h4).Dtor_turns(), (_12_dt__update__tmp_h4).Dtor_title(), (_12_dt__update__tmp_h4).Dtor_status(), (_12_dt__update__tmp_h4).Dtor_modifiedAt(), (_12_dt__update__tmp_h4).Dtor_draft(), (_12_dt__update__tmp_h4).Dtor_activity(), (_12_dt__update__tmp_h4).Dtor_activeTurn(), (_12_dt__update__tmp_h4).Dtor_steeringMessage(), (_12_dt__update__tmp_h4).Dtor_queuedMessages(), _13_dt__update_hnextCursor_h0)
				}(_pat_let430_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("cur")))
		}(_pat_let429_0)
	}(_9_three), Companion_ChatAction_.Create_CTruncated_(m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(func(_pat_let431_0 ChatState) ChatState {
		return func(_14_dt__update__tmp_h5 ChatState) ChatState {
			return func(_pat_let432_0 m_AhpSkeleton.Option) ChatState {
				return func(_15_dt__update_hnextCursor_h1 m_AhpSkeleton.Option) ChatState {
					return func(_pat_let433_0 _dafny.Sequence) ChatState {
						return func(_16_dt__update_hturns_h1 _dafny.Sequence) ChatState {
							return Companion_ChatState_.Create_ChatState_(_16_dt__update_hturns_h1, (_14_dt__update__tmp_h5).Dtor_title(), (_14_dt__update__tmp_h5).Dtor_status(), (_14_dt__update__tmp_h5).Dtor_modifiedAt(), (_14_dt__update__tmp_h5).Dtor_draft(), (_14_dt__update__tmp_h5).Dtor_activity(), (_14_dt__update__tmp_h5).Dtor_activeTurn(), (_14_dt__update__tmp_h5).Dtor_steeringMessage(), (_14_dt__update__tmp_h5).Dtor_queuedMessages(), _15_dt__update_hnextCursor_h1)
						}(_pat_let433_0)
					}(_dafny.SeqOf())
				}(_pat_let432_0)
			}(m_AhpSkeleton.Companion_Option_.Create_None_())
		}(_pat_let431_0)
	}(_9_three)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _17_qa QMsg
	_ = _17_qa
	_17_qa = Companion_QMsg_.Create_QMsg_(_dafny.UnicodeSeqOfUtf8Bytes("a"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("A")))
	var _18_qb QMsg
	_ = _18_qb
	_18_qb = Companion_QMsg_.Create_QMsg_(_dafny.UnicodeSeqOfUtf8Bytes("b"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("B")))
	var _19_qc QMsg
	_ = _19_qc
	_19_qc = Companion_QMsg_.Create_QMsg_(_dafny.UnicodeSeqOfUtf8Bytes("c"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("C")))
	var _pat_let_tv1 = _17_qa
	_ = _pat_let_tv1
	var _pat_let_tv2 = _18_qb
	_ = _pat_let_tv2
	var _pat_let_tv3 = _19_qc
	_ = _pat_let_tv3
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let434_0 ChatState) ChatState {
		return func(_20_dt__update__tmp_h6 ChatState) ChatState {
			return func(_pat_let435_0 _dafny.Sequence) ChatState {
				return func(_21_dt__update_hqueuedMessages_h1 _dafny.Sequence) ChatState {
					return Companion_ChatState_.Create_ChatState_((_20_dt__update__tmp_h6).Dtor_turns(), (_20_dt__update__tmp_h6).Dtor_title(), (_20_dt__update__tmp_h6).Dtor_status(), (_20_dt__update__tmp_h6).Dtor_modifiedAt(), (_20_dt__update__tmp_h6).Dtor_draft(), (_20_dt__update__tmp_h6).Dtor_activity(), (_20_dt__update__tmp_h6).Dtor_activeTurn(), (_20_dt__update__tmp_h6).Dtor_steeringMessage(), _21_dt__update_hqueuedMessages_h1, (_20_dt__update__tmp_h6).Dtor_nextCursor())
				}(_pat_let435_0)
			}(_dafny.SeqOf(_pat_let_tv1, _pat_let_tv2, _pat_let_tv3))
		}(_pat_let434_0)
	}(_0_s), Companion_ChatAction_.Create_CQueuedReordered_(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("c"), _dafny.UnicodeSeqOfUtf8Bytes("a"), _dafny.UnicodeSeqOfUtf8Bytes("b"))))).Dtor_queuedMessages(), _dafny.SeqOf(_19_qc, _17_qa, _18_qb)) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(func(_pat_let436_0 ChatState) ChatState {
		return func(_22_dt__update__tmp_h7 ChatState) ChatState {
			return func(_pat_let437_0 _dafny.Sequence) ChatState {
				return func(_23_dt__update_hturns_h2 _dafny.Sequence) ChatState {
					return Companion_ChatState_.Create_ChatState_(_23_dt__update_hturns_h2, (_22_dt__update__tmp_h7).Dtor_title(), (_22_dt__update__tmp_h7).Dtor_status(), (_22_dt__update__tmp_h7).Dtor_modifiedAt(), (_22_dt__update__tmp_h7).Dtor_draft(), (_22_dt__update__tmp_h7).Dtor_activity(), (_22_dt__update__tmp_h7).Dtor_activeTurn(), (_22_dt__update__tmp_h7).Dtor_steeringMessage(), (_22_dt__update__tmp_h7).Dtor_queuedMessages(), (_22_dt__update__tmp_h7).Dtor_nextCursor())
				}(_pat_let437_0)
			}(_dafny.SeqOf(Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t3"))))
		}(_pat_let436_0)
	}(_0_s), Companion_ChatAction_.Create_CTurnsLoaded_(_dafny.SeqOf(Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t1")), Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t2"))), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("cur-0"))))).Equals(func(_pat_let438_0 ChatState) ChatState {
		return func(_24_dt__update__tmp_h8 ChatState) ChatState {
			return func(_pat_let439_0 m_AhpSkeleton.Option) ChatState {
				return func(_25_dt__update_hnextCursor_h2 m_AhpSkeleton.Option) ChatState {
					return func(_pat_let440_0 _dafny.Sequence) ChatState {
						return func(_26_dt__update_hturns_h3 _dafny.Sequence) ChatState {
							return Companion_ChatState_.Create_ChatState_(_26_dt__update_hturns_h3, (_24_dt__update__tmp_h8).Dtor_title(), (_24_dt__update__tmp_h8).Dtor_status(), (_24_dt__update__tmp_h8).Dtor_modifiedAt(), (_24_dt__update__tmp_h8).Dtor_draft(), (_24_dt__update__tmp_h8).Dtor_activity(), (_24_dt__update__tmp_h8).Dtor_activeTurn(), (_24_dt__update__tmp_h8).Dtor_steeringMessage(), (_24_dt__update__tmp_h8).Dtor_queuedMessages(), _25_dt__update_hnextCursor_h2)
						}(_pat_let440_0)
					}(_dafny.SeqOf(Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t1")), Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t2")), Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t3"))))
				}(_pat_let439_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("cur-0")))
		}(_pat_let438_0)
	}(_0_s)) {
		pass = (pass).Plus(_dafny.One)
	}
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let441_0 ChatState) ChatState {
		return func(_27_dt__update__tmp_h9 ChatState) ChatState {
			return func(_pat_let442_0 _dafny.Sequence) ChatState {
				return func(_28_dt__update_hturns_h4 _dafny.Sequence) ChatState {
					return Companion_ChatState_.Create_ChatState_(_28_dt__update_hturns_h4, (_27_dt__update__tmp_h9).Dtor_title(), (_27_dt__update__tmp_h9).Dtor_status(), (_27_dt__update__tmp_h9).Dtor_modifiedAt(), (_27_dt__update__tmp_h9).Dtor_draft(), (_27_dt__update__tmp_h9).Dtor_activity(), (_27_dt__update__tmp_h9).Dtor_activeTurn(), (_27_dt__update__tmp_h9).Dtor_steeringMessage(), (_27_dt__update__tmp_h9).Dtor_queuedMessages(), (_27_dt__update__tmp_h9).Dtor_nextCursor())
				}(_pat_let442_0)
			}(_dafny.SeqOf(Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t3"))))
		}(_pat_let441_0)
	}(_0_s), Companion_ChatAction_.Create_CTurnsLoaded_(_dafny.SeqOf(Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t1")), Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t2")), Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t3"))), m_AhpSkeleton.Companion_Option_.Create_None_()))).Dtor_turns(), _dafny.SeqOf(Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t1")), Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t2")), Companion_Default___.TDone(_dafny.UnicodeSeqOfUtf8Bytes("t3")))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _29_running ChatState
	_ = _29_running
	var _30_dt__update__tmp_h10 ChatState = _1_act1
	_ = _30_dt__update__tmp_h10
	var _31_dt__update_hactiveTurn_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let443_0 Turn) Turn {
		return func(_32_dt__update__tmp_h11 Turn) Turn {
			return func(_pat_let444_0 _dafny.Sequence) Turn {
				return func(_35_dt__update_hparts_h0 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_32_dt__update__tmp_h11).Dtor_turnId(), (_32_dt__update__tmp_h11).Dtor_message(), _35_dt__update_hparts_h0, (_32_dt__update__tmp_h11).Dtor_state(), (_32_dt__update__tmp_h11).Dtor_usage(), (_32_dt__update__tmp_h11).Dtor_error())
				}(_pat_let444_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PToolCall_(func(_pat_let445_0 ToolCall) ToolCall {
				return func(_33_dt__update__tmp_h12 ToolCall) ToolCall {
					return func(_pat_let446_0 _dafny.Sequence) ToolCall {
						return func(_34_dt__update_hstatus_h0 _dafny.Sequence) ToolCall {
							return Companion_ToolCall_.Create_ToolCall_((_33_dt__update__tmp_h12).Dtor_toolCallId(), (_33_dt__update__tmp_h12).Dtor_toolName(), (_33_dt__update__tmp_h12).Dtor_displayName(), _34_dt__update_hstatus_h0, (_33_dt__update__tmp_h12).Dtor_intention(), (_33_dt__update__tmp_h12).Dtor_contributor(), (_33_dt__update__tmp_h12).Dtor_meta(), (_33_dt__update__tmp_h12).Dtor_invocationMessage(), (_33_dt__update__tmp_h12).Dtor_toolInput(), (_33_dt__update__tmp_h12).Dtor_confirmationTitle(), (_33_dt__update__tmp_h12).Dtor_riskAssessment(), (_33_dt__update__tmp_h12).Dtor_edits(), (_33_dt__update__tmp_h12).Dtor_editable(), (_33_dt__update__tmp_h12).Dtor_options(), (_33_dt__update__tmp_h12).Dtor_confirmed(), (_33_dt__update__tmp_h12).Dtor_success(), (_33_dt__update__tmp_h12).Dtor_pastTenseMessage(), (_33_dt__update__tmp_h12).Dtor_reason(), (_33_dt__update__tmp_h12).Dtor_reasonMessage(), (_33_dt__update__tmp_h12).Dtor_userSuggestion(), (_33_dt__update__tmp_h12).Dtor_selectedOption(), (_33_dt__update__tmp_h12).Dtor_content(), (_33_dt__update__tmp_h12).Dtor_structuredContent(), (_33_dt__update__tmp_h12).Dtor_error(), (_33_dt__update__tmp_h12).Dtor_auth(), (_33_dt__update__tmp_h12).Dtor_partialInput())
						}(_pat_let446_0)
					}(_dafny.UnicodeSeqOfUtf8Bytes("running"))
				}(_pat_let445_0)
			}(Companion_Default___.TC0(_dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_None_())))))
		}(_pat_let443_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))
	_ = _31_dt__update_hactiveTurn_h1
	_29_running = Companion_ChatState_.Create_ChatState_((_30_dt__update__tmp_h10).Dtor_turns(), (_30_dt__update__tmp_h10).Dtor_title(), (_30_dt__update__tmp_h10).Dtor_status(), (_30_dt__update__tmp_h10).Dtor_modifiedAt(), (_30_dt__update__tmp_h10).Dtor_draft(), (_30_dt__update__tmp_h10).Dtor_activity(), _31_dt__update_hactiveTurn_h1, (_30_dt__update__tmp_h10).Dtor_steeringMessage(), (_30_dt__update__tmp_h10).Dtor_queuedMessages(), (_30_dt__update__tmp_h10).Dtor_nextCursor())
	var _36_afterCc ChatState
	_ = _36_afterCc
	_36_afterCc = Companion_Default___.Apply1(_29_running, Companion_ChatAction_.Create_CToolCallContentChanged_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("terminal-ref")), m_AhpSkeleton.Companion_Option_.Create_None_()))
	var _37_ccTc Part
	_ = _37_ccTc
	_37_ccTc = ((((_36_afterCc).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts()).Select(0).(Part)
	if (((_37_ccTc).Is_PToolCall()) && (_dafny.Companion_Sequence_.Equal(((_37_ccTc).Dtor_tc()).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("running")))) && ((((_37_ccTc).Dtor_tc()).Dtor_content()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("terminal-ref"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _38_withR ChatState
	_ = _38_withR
	var _39_dt__update__tmp_h13 ChatState = _1_act1
	_ = _39_dt__update__tmp_h13
	var _40_dt__update_hactiveTurn_h2 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let447_0 Turn) Turn {
		return func(_41_dt__update__tmp_h14 Turn) Turn {
			return func(_pat_let448_0 _dafny.Sequence) Turn {
				return func(_42_dt__update_hparts_h1 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_41_dt__update__tmp_h14).Dtor_turnId(), (_41_dt__update__tmp_h14).Dtor_message(), _42_dt__update_hparts_h1, (_41_dt__update__tmp_h14).Dtor_state(), (_41_dt__update__tmp_h14).Dtor_usage(), (_41_dt__update__tmp_h14).Dtor_error())
				}(_pat_let448_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PReasoning_(_dafny.UnicodeSeqOfUtf8Bytes("r-1"), _dafny.UnicodeSeqOfUtf8Bytes(""))))
		}(_pat_let447_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))
	_ = _40_dt__update_hactiveTurn_h2
	_38_withR = Companion_ChatState_.Create_ChatState_((_39_dt__update__tmp_h13).Dtor_turns(), (_39_dt__update__tmp_h13).Dtor_title(), (_39_dt__update__tmp_h13).Dtor_status(), (_39_dt__update__tmp_h13).Dtor_modifiedAt(), (_39_dt__update__tmp_h13).Dtor_draft(), (_39_dt__update__tmp_h13).Dtor_activity(), _40_dt__update_hactiveTurn_h2, (_39_dt__update__tmp_h13).Dtor_steeringMessage(), (_39_dt__update__tmp_h13).Dtor_queuedMessages(), (_39_dt__update__tmp_h13).Dtor_nextCursor())
	var _43_r2 ChatState
	_ = _43_r2
	_43_r2 = Companion_Default___.Fold(_38_withR, _dafny.SeqOf(Companion_ChatAction_.Create_CReasoning_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("r-1"), _dafny.UnicodeSeqOfUtf8Bytes("Thinking about ")), Companion_ChatAction_.Create_CReasoning_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("r-1"), _dafny.UnicodeSeqOfUtf8Bytes("the plan"))))
	if ((_43_r2).Dtor_activeTurn()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let449_0 Turn) Turn {
		return func(_44_dt__update__tmp_h15 Turn) Turn {
			return func(_pat_let450_0 _dafny.Sequence) Turn {
				return func(_45_dt__update_hparts_h2 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_44_dt__update__tmp_h15).Dtor_turnId(), (_44_dt__update__tmp_h15).Dtor_message(), _45_dt__update_hparts_h2, (_44_dt__update__tmp_h15).Dtor_state(), (_44_dt__update__tmp_h15).Dtor_usage(), (_44_dt__update__tmp_h15).Dtor_error())
				}(_pat_let450_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PReasoning_(_dafny.UnicodeSeqOfUtf8Bytes("r-1"), _dafny.UnicodeSeqOfUtf8Bytes("Thinking about the plan"))))
		}(_pat_let449_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass
}
func (_static *CompanionStruct_Default___) RunInputFlow() _dafny.Int {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	pass = _dafny.Zero
	var _0_s ChatState
	_ = _0_s
	_0_s = Companion_Default___.C0()
	var _1_openBody _dafny.Map
	_ = _1_openBody
	_1_openBody = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("message"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("Clarify requirements"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("questions"), m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("id"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("q1"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("required"), m_ConfluxCodec.Companion_Json_.Create_JBool_(true)).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("options"), m_ConfluxCodec.Companion_Json_.Create_JArr_(_dafny.SeqOf(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("node"))))))))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("url"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("https://example.test/review")))
	var _2_at1 ChatState
	_ = _2_at1
	var _3_dt__update__tmp_h0 ChatState = _0_s
	_ = _3_dt__update__tmp_h0
	var _4_dt__update_hstatus_h0 uint32 = uint32(8)
	_ = _4_dt__update_hstatus_h0
	var _5_dt__update_hactiveTurn_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("turn-1")))
	_ = _5_dt__update_hactiveTurn_h0
	_2_at1 = Companion_ChatState_.Create_ChatState_((_3_dt__update__tmp_h0).Dtor_turns(), (_3_dt__update__tmp_h0).Dtor_title(), _4_dt__update_hstatus_h0, (_3_dt__update__tmp_h0).Dtor_modifiedAt(), (_3_dt__update__tmp_h0).Dtor_draft(), (_3_dt__update__tmp_h0).Dtor_activity(), _5_dt__update_hactiveTurn_h0, (_3_dt__update__tmp_h0).Dtor_steeringMessage(), (_3_dt__update__tmp_h0).Dtor_queuedMessages(), (_3_dt__update__tmp_h0).Dtor_nextCursor())
	var _6_afterReq ChatState
	_ = _6_afterReq
	_6_afterReq = Companion_Default___.Apply1(_2_at1, Companion_ChatAction_.Create_CInputRequested_(Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.NewMapBuilder().ToMap(), _1_openBody)))
	if ((((_6_afterReq).Dtor_status()) == (uint32(24))) && (((_6_afterReq).Dtor_activeTurn()).Is_Some())) && (_dafny.Companion_Sequence_.Equal((((_6_afterReq).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _dafny.SeqOf(Companion_Part_.Create_PInputRequest_(Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.NewMapBuilder().ToMap(), _1_openBody), m_AhpSkeleton.Companion_Option_.Create_None_())))) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(func(_pat_let451_0 ChatState) ChatState {
		return func(_7_dt__update__tmp_h1 ChatState) ChatState {
			return func(_pat_let452_0 uint32) ChatState {
				return func(_8_dt__update_hstatus_h1 uint32) ChatState {
					return Companion_ChatState_.Create_ChatState_((_7_dt__update__tmp_h1).Dtor_turns(), (_7_dt__update__tmp_h1).Dtor_title(), _8_dt__update_hstatus_h1, (_7_dt__update__tmp_h1).Dtor_modifiedAt(), (_7_dt__update__tmp_h1).Dtor_draft(), (_7_dt__update__tmp_h1).Dtor_activity(), (_7_dt__update__tmp_h1).Dtor_activeTurn(), (_7_dt__update__tmp_h1).Dtor_steeringMessage(), (_7_dt__update__tmp_h1).Dtor_queuedMessages(), (_7_dt__update__tmp_h1).Dtor_nextCursor())
				}(_pat_let452_0)
			}(uint32(8))
		}(_pat_let451_0)
	}(_0_s), Companion_ChatAction_.Create_CInputRequested_(Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("orphan-1"), _dafny.NewMapBuilder().ToMap(), _1_openBody)))).Equals(func(_pat_let453_0 ChatState) ChatState {
		return func(_9_dt__update__tmp_h2 ChatState) ChatState {
			return func(_pat_let454_0 uint32) ChatState {
				return func(_10_dt__update_hstatus_h2 uint32) ChatState {
					return Companion_ChatState_.Create_ChatState_((_9_dt__update__tmp_h2).Dtor_turns(), (_9_dt__update__tmp_h2).Dtor_title(), _10_dt__update_hstatus_h2, (_9_dt__update__tmp_h2).Dtor_modifiedAt(), (_9_dt__update__tmp_h2).Dtor_draft(), (_9_dt__update__tmp_h2).Dtor_activity(), (_9_dt__update__tmp_h2).Dtor_activeTurn(), (_9_dt__update__tmp_h2).Dtor_steeringMessage(), (_9_dt__update__tmp_h2).Dtor_queuedMessages(), (_9_dt__update__tmp_h2).Dtor_nextCursor())
				}(_pat_let454_0)
			}(uint32(8))
		}(_pat_let453_0)
	}(_0_s)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _pat_let_tv0 = _1_openBody
	_ = _pat_let_tv0
	var _11_reqState ChatState
	_ = _11_reqState
	var _12_dt__update__tmp_h3 ChatState = _2_at1
	_ = _12_dt__update__tmp_h3
	var _13_dt__update_hstatus_h3 uint32 = uint32(24)
	_ = _13_dt__update_hstatus_h3
	var _14_dt__update_hactiveTurn_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let455_0 Turn) Turn {
		return func(_15_dt__update__tmp_h4 Turn) Turn {
			return func(_pat_let456_0 _dafny.Sequence) Turn {
				return func(_16_dt__update_hparts_h0 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_15_dt__update__tmp_h4).Dtor_turnId(), (_15_dt__update__tmp_h4).Dtor_message(), _16_dt__update_hparts_h0, (_15_dt__update__tmp_h4).Dtor_state(), (_15_dt__update__tmp_h4).Dtor_usage(), (_15_dt__update__tmp_h4).Dtor_error())
				}(_pat_let456_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PInputRequest_(Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.NewMapBuilder().ToMap(), _pat_let_tv0), m_AhpSkeleton.Companion_Option_.Create_None_())))
		}(_pat_let455_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("turn-1"))))
	_ = _14_dt__update_hactiveTurn_h1
	_11_reqState = Companion_ChatState_.Create_ChatState_((_12_dt__update__tmp_h3).Dtor_turns(), (_12_dt__update__tmp_h3).Dtor_title(), _13_dt__update_hstatus_h3, (_12_dt__update__tmp_h3).Dtor_modifiedAt(), (_12_dt__update__tmp_h3).Dtor_draft(), (_12_dt__update__tmp_h3).Dtor_activity(), _14_dt__update_hactiveTurn_h1, (_12_dt__update__tmp_h3).Dtor_steeringMessage(), (_12_dt__update__tmp_h3).Dtor_queuedMessages(), (_12_dt__update__tmp_h3).Dtor_nextCursor())
	if _dafny.Companion_Sequence_.Equal((((Companion_Default___.Apply1(_11_reqState, Companion_ChatAction_.Create_CInputAnswerChanged_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.UnicodeSeqOfUtf8Bytes("q1"), m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("node")))))).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _dafny.SeqOf(Companion_Part_.Create_PInputRequest_(Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("q1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("node"))), _1_openBody), m_AhpSkeleton.Companion_Option_.Create_None_()))) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_11_reqState, Companion_ChatAction_.Create_CInputAnswerChanged_(_dafny.UnicodeSeqOfUtf8Bytes("missing"), _dafny.UnicodeSeqOfUtf8Bytes("q1"), m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(_11_reqState) {
		pass = (pass).Plus(_dafny.One)
	}
	var _17_afterDone ChatState
	_ = _17_afterDone
	_17_afterDone = Companion_Default___.Apply1(_11_reqState, Companion_ChatAction_.Create_CInputCompleted_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.UnicodeSeqOfUtf8Bytes("accept"), _dafny.NewMapBuilder().ToMap()))
	if ((((_17_afterDone).Dtor_status()) == (uint32(8))) && (((_17_afterDone).Dtor_activeTurn()).Is_Some())) && (_dafny.Companion_Sequence_.Equal((((_17_afterDone).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _dafny.SeqOf(Companion_Part_.Create_PInputRequest_(Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.NewMapBuilder().ToMap(), _1_openBody), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("accept")))))) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_11_reqState, Companion_ChatAction_.Create_CInputCompleted_(_dafny.UnicodeSeqOfUtf8Bytes("missing"), _dafny.UnicodeSeqOfUtf8Bytes("cancel"), _dafny.NewMapBuilder().ToMap()))).Equals(_11_reqState) {
		pass = (pass).Plus(_dafny.One)
	}
	var _18_flow ChatState
	_ = _18_flow
	_18_flow = Companion_Default___.Fold(_0_s, _dafny.SeqOf(Companion_ChatAction_.Create_CTurnStarted_(_dafny.UnicodeSeqOfUtf8Bytes("turn-1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("Hello")), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CInputRequested_(Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.NewMapBuilder().ToMap(), _1_openBody)), Companion_ChatAction_.Create_CInputAnswerChanged_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.UnicodeSeqOfUtf8Bytes("q1"), m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("node")))), Companion_ChatAction_.Create_CInputAnswerChanged_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.UnicodeSeqOfUtf8Bytes("q2"), m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("small")))), Companion_ChatAction_.Create_CInputCompleted_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.UnicodeSeqOfUtf8Bytes("accept"), _dafny.NewMapBuilder().ToMap())))
	if ((((_18_flow).Dtor_status()) == (uint32(8))) && (((_18_flow).Dtor_activeTurn()).Is_Some())) && (_dafny.Companion_Sequence_.Equal((((_18_flow).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _dafny.SeqOf(Companion_Part_.Create_PInputRequest_(Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("input-1"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("q1"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("node"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("q2"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("small"))), _1_openBody), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("accept")))))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _pat_let_tv1 = _1_openBody
	_ = _pat_let_tv1
	var _19_overrideState ChatState
	_ = _19_overrideState
	var _20_dt__update__tmp_h5 ChatState = _2_at1
	_ = _20_dt__update__tmp_h5
	var _21_dt__update_hstatus_h4 uint32 = uint32(24)
	_ = _21_dt__update_hstatus_h4
	var _22_dt__update_hactiveTurn_h2 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let457_0 Turn) Turn {
		return func(_23_dt__update__tmp_h6 Turn) Turn {
			return func(_pat_let458_0 _dafny.Sequence) Turn {
				return func(_24_dt__update_hparts_h1 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_23_dt__update__tmp_h6).Dtor_turnId(), (_23_dt__update__tmp_h6).Dtor_message(), _24_dt__update_hparts_h1, (_23_dt__update__tmp_h6).Dtor_state(), (_23_dt__update__tmp_h6).Dtor_usage(), (_23_dt__update__tmp_h6).Dtor_error())
				}(_pat_let458_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PInputRequest_(Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("review-1"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("approve"), m_ConfluxCodec.Companion_Json_.Create_JBool_(true)), _pat_let_tv1), m_AhpSkeleton.Companion_Option_.Create_None_())))
		}(_pat_let457_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("turn-1"))))
	_ = _22_dt__update_hactiveTurn_h2
	_19_overrideState = Companion_ChatState_.Create_ChatState_((_20_dt__update__tmp_h5).Dtor_turns(), (_20_dt__update__tmp_h5).Dtor_title(), _21_dt__update_hstatus_h4, (_20_dt__update__tmp_h5).Dtor_modifiedAt(), (_20_dt__update__tmp_h5).Dtor_draft(), (_20_dt__update__tmp_h5).Dtor_activity(), _22_dt__update_hactiveTurn_h2, (_20_dt__update__tmp_h5).Dtor_steeringMessage(), (_20_dt__update__tmp_h5).Dtor_queuedMessages(), (_20_dt__update__tmp_h5).Dtor_nextCursor())
	var _25_overridden ChatState
	_ = _25_overridden
	_25_overridden = Companion_Default___.Apply1(_19_overrideState, Companion_ChatAction_.Create_CInputCompleted_(_dafny.UnicodeSeqOfUtf8Bytes("review-1"), _dafny.UnicodeSeqOfUtf8Bytes("decline"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("approve"), m_ConfluxCodec.Companion_Json_.Create_JBool_(false))))
	if ((((_25_overridden).Dtor_status()) == (uint32(8))) && (((_25_overridden).Dtor_activeTurn()).Is_Some())) && (_dafny.Companion_Sequence_.Equal((((_25_overridden).Dtor_activeTurn()).Dtor_value().(Turn)).Dtor_parts(), _dafny.SeqOf(Companion_Part_.Create_PInputRequest_(Companion_InputReq_.Create_InputReq_(_dafny.UnicodeSeqOfUtf8Bytes("review-1"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("approve"), m_ConfluxCodec.Companion_Json_.Create_JBool_(false)), _1_openBody), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("decline")))))) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass
}
func (_static *CompanionStruct_Default___) RunResultConfirm() _dafny.Int {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	pass = _dafny.Zero
	var _0_s ChatState
	_ = _0_s
	_0_s = Companion_Default___.C0()
	var _1_act1 ChatState
	_ = _1_act1
	var _2_dt__update__tmp_h0 ChatState = _0_s
	_ = _2_dt__update__tmp_h0
	var _3_dt__update_hactiveTurn_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1")))
	_ = _3_dt__update_hactiveTurn_h0
	_1_act1 = Companion_ChatState_.Create_ChatState_((_2_dt__update__tmp_h0).Dtor_turns(), (_2_dt__update__tmp_h0).Dtor_title(), (_2_dt__update__tmp_h0).Dtor_status(), (_2_dt__update__tmp_h0).Dtor_modifiedAt(), (_2_dt__update__tmp_h0).Dtor_draft(), (_2_dt__update__tmp_h0).Dtor_activity(), _3_dt__update_hactiveTurn_h0, (_2_dt__update__tmp_h0).Dtor_steeringMessage(), (_2_dt__update__tmp_h0).Dtor_queuedMessages(), (_2_dt__update__tmp_h0).Dtor_nextCursor())
	var _4_rcLife ChatState
	_ = _4_rcLife
	_4_rcLife = Companion_Default___.Fold(_1_act1, _dafny.SeqOf(Companion_ChatAction_.Create_CToolCallStart_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), _dafny.UnicodeSeqOfUtf8Bytes("bash"), _dafny.UnicodeSeqOfUtf8Bytes("Run Command"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallReady_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Run")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("not-needed")), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallComplete_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Done")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), false, m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallResultConfirmed_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_None_())))
	var _5_rcTc m_AhpSkeleton.Option
	_ = _5_rcTc
	_5_rcTc = Companion_Default___.FirstActiveToolCall(_4_rcLife)
	if (((_5_rcTc).Is_Some()) && (_dafny.Companion_Sequence_.Equal(((_5_rcTc).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("completed")))) && ((((_5_rcTc).Dtor_value().(ToolCall)).Dtor_success()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(true))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _6_rdLife ChatState
	_ = _6_rdLife
	_6_rdLife = Companion_Default___.Fold(_1_act1, _dafny.SeqOf(Companion_ChatAction_.Create_CToolCallStart_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), _dafny.UnicodeSeqOfUtf8Bytes("bash"), _dafny.UnicodeSeqOfUtf8Bytes("Run Command"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallReady_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Run")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("not-needed")), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallComplete_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Done")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), true, m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallResultConfirmed_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), false, m_AhpSkeleton.Companion_Option_.Create_None_())))
	var _7_rdTc m_AhpSkeleton.Option
	_ = _7_rdTc
	_7_rdTc = Companion_Default___.FirstActiveToolCall(_6_rdLife)
	if (((((_7_rdTc).Is_Some()) && (_dafny.Companion_Sequence_.Equal(((_7_rdTc).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("cancelled")))) && ((((_7_rdTc).Dtor_value().(ToolCall)).Dtor_confirmed()).Equals(m_AhpSkeleton.Companion_Option_.Create_None_()))) && ((((_7_rdTc).Dtor_value().(ToolCall)).Dtor_success()).Equals(m_AhpSkeleton.Companion_Option_.Create_None_()))) && ((((_7_rdTc).Dtor_value().(ToolCall)).Dtor_reason()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("result-denied")))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _8_runState ChatState
	_ = _8_runState
	var _9_dt__update__tmp_h1 ChatState = _1_act1
	_ = _9_dt__update__tmp_h1
	var _10_dt__update_hstatus_h0 uint32 = uint32(8)
	_ = _10_dt__update_hstatus_h0
	var _11_dt__update_hactiveTurn_h1 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(func(_pat_let459_0 Turn) Turn {
		return func(_12_dt__update__tmp_h2 Turn) Turn {
			return func(_pat_let460_0 _dafny.Sequence) Turn {
				return func(_16_dt__update_hparts_h0 _dafny.Sequence) Turn {
					return Companion_Turn_.Create_Turn_((_12_dt__update__tmp_h2).Dtor_turnId(), (_12_dt__update__tmp_h2).Dtor_message(), _16_dt__update_hparts_h0, (_12_dt__update__tmp_h2).Dtor_state(), (_12_dt__update__tmp_h2).Dtor_usage(), (_12_dt__update__tmp_h2).Dtor_error())
				}(_pat_let460_0)
			}(_dafny.SeqOf(Companion_Part_.Create_PToolCall_(func(_pat_let461_0 ToolCall) ToolCall {
				return func(_13_dt__update__tmp_h3 ToolCall) ToolCall {
					return func(_pat_let462_0 m_AhpSkeleton.Option) ToolCall {
						return func(_14_dt__update_hconfirmed_h0 m_AhpSkeleton.Option) ToolCall {
							return func(_pat_let463_0 _dafny.Sequence) ToolCall {
								return func(_15_dt__update_hstatus_h1 _dafny.Sequence) ToolCall {
									return Companion_ToolCall_.Create_ToolCall_((_13_dt__update__tmp_h3).Dtor_toolCallId(), (_13_dt__update__tmp_h3).Dtor_toolName(), (_13_dt__update__tmp_h3).Dtor_displayName(), _15_dt__update_hstatus_h1, (_13_dt__update__tmp_h3).Dtor_intention(), (_13_dt__update__tmp_h3).Dtor_contributor(), (_13_dt__update__tmp_h3).Dtor_meta(), (_13_dt__update__tmp_h3).Dtor_invocationMessage(), (_13_dt__update__tmp_h3).Dtor_toolInput(), (_13_dt__update__tmp_h3).Dtor_confirmationTitle(), (_13_dt__update__tmp_h3).Dtor_riskAssessment(), (_13_dt__update__tmp_h3).Dtor_edits(), (_13_dt__update__tmp_h3).Dtor_editable(), (_13_dt__update__tmp_h3).Dtor_options(), _14_dt__update_hconfirmed_h0, (_13_dt__update__tmp_h3).Dtor_success(), (_13_dt__update__tmp_h3).Dtor_pastTenseMessage(), (_13_dt__update__tmp_h3).Dtor_reason(), (_13_dt__update__tmp_h3).Dtor_reasonMessage(), (_13_dt__update__tmp_h3).Dtor_userSuggestion(), (_13_dt__update__tmp_h3).Dtor_selectedOption(), (_13_dt__update__tmp_h3).Dtor_content(), (_13_dt__update__tmp_h3).Dtor_structuredContent(), (_13_dt__update__tmp_h3).Dtor_error(), (_13_dt__update__tmp_h3).Dtor_auth(), (_13_dt__update__tmp_h3).Dtor_partialInput())
								}(_pat_let463_0)
							}(_dafny.UnicodeSeqOfUtf8Bytes("running"))
						}(_pat_let462_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("user-action")))
				}(_pat_let461_0)
			}(Companion_Default___.TC0(_dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_None_())))))
		}(_pat_let459_0)
	}(Companion_Default___.T0(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))
	_ = _11_dt__update_hactiveTurn_h1
	_8_runState = Companion_ChatState_.Create_ChatState_((_9_dt__update__tmp_h1).Dtor_turns(), (_9_dt__update__tmp_h1).Dtor_title(), _10_dt__update_hstatus_h0, (_9_dt__update__tmp_h1).Dtor_modifiedAt(), (_9_dt__update__tmp_h1).Dtor_draft(), (_9_dt__update__tmp_h1).Dtor_activity(), _11_dt__update_hactiveTurn_h1, (_9_dt__update__tmp_h1).Dtor_steeringMessage(), (_9_dt__update__tmp_h1).Dtor_queuedMessages(), (_9_dt__update__tmp_h1).Dtor_nextCursor())
	if (Companion_Default___.Apply1(_8_runState, Companion_ChatAction_.Create_CToolCallResultConfirmed_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(_8_runState) {
		pass = (pass).Plus(_dafny.One)
	}
	var _17_selLife ChatState
	_ = _17_selLife
	_17_selLife = Companion_Default___.Fold(_1_act1, _dafny.SeqOf(Companion_ChatAction_.Create_CToolCallStart_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), _dafny.UnicodeSeqOfUtf8Bytes("bash"), _dafny.UnicodeSeqOfUtf8Bytes("Run Command"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallReady_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Run")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallConfirmed_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("user-action")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallComplete_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("Done")), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), false, m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChatAction_.Create_CToolCallResultConfirmed_(_dafny.UnicodeSeqOfUtf8Bytes("t1"), _dafny.UnicodeSeqOfUtf8Bytes("tc-1"), true, m_AhpSkeleton.Companion_Option_.Create_None_())))
	var _18_selTc m_AhpSkeleton.Option
	_ = _18_selTc
	_18_selTc = Companion_Default___.FirstActiveToolCall(_17_selLife)
	if ((((_18_selTc).Is_Some()) && (_dafny.Companion_Sequence_.Equal(((_18_selTc).Dtor_value().(ToolCall)).Dtor_status(), _dafny.UnicodeSeqOfUtf8Bytes("completed")))) && ((((_18_selTc).Dtor_value().(ToolCall)).Dtor_confirmed()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("user-action"))))) && ((((_18_selTc).Dtor_value().(ToolCall)).Dtor_success()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(true))) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass
}
func (_static *CompanionStruct_Default___) RunCorpus() (_dafny.Int, _dafny.Int) {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	var modeled _dafny.Int = _dafny.Zero
	_ = modeled
	modeled = _dafny.IntOfInt64(54)
	var _0_p1 _dafny.Int
	_ = _0_p1
	var _out0 _dafny.Int
	_ = _out0
	_out0 = Companion_Default___.RunScalarTurns()
	_0_p1 = _out0
	var _1_p2 _dafny.Int
	_ = _1_p2
	var _out1 _dafny.Int
	_ = _out1
	_out1 = Companion_Default___.RunToolCalls()
	_1_p2 = _out1
	var _2_p3 _dafny.Int
	_ = _2_p3
	var _out2 _dafny.Int
	_ = _out2
	_out2 = Companion_Default___.RunPendingHistory()
	_2_p3 = _out2
	var _3_p4 _dafny.Int
	_ = _3_p4
	var _out3 _dafny.Int
	_ = _out3
	_out3 = Companion_Default___.RunInputFlow()
	_3_p4 = _out3
	var _4_p5 _dafny.Int
	_ = _4_p5
	var _out4 _dafny.Int
	_ = _out4
	_out4 = Companion_Default___.RunResultConfirm()
	_4_p5 = _out4
	pass = ((((_0_p1).Plus(_1_p2)).Plus(_2_p3)).Plus(_3_p4)).Plus(_4_p5)
	return pass, modeled
}
func (_static *CompanionStruct_Default___) QKey() func(QMsg) _dafny.Sequence {
	return func(_0_x QMsg) _dafny.Sequence {
		return (_0_x).Dtor_id()
	}
}
func (_static *CompanionStruct_Default___) ERR() uint32 {
	return uint32(2)
}
func (_static *CompanionStruct_Default___) INP() uint32 {
	return uint32(24)
}
func (_static *CompanionStruct_Default___) GEN() uint32 {
	return uint32(8)
}
func (_static *CompanionStruct_Default___) IDLE() uint32 {
	return uint32(1)
}
func (_static *CompanionStruct_Default___) ACT() uint32 {
	return uint32(31)
}
func (_static *CompanionStruct_Default___) READ() uint32 {
	return uint32(32)
}

// End of class Default__

// Definition of datatype ToolCallContributor
type ToolCallContributor struct {
	Data_ToolCallContributor_
}

func (_this ToolCallContributor) Get_() Data_ToolCallContributor_ {
	return _this.Data_ToolCallContributor_
}

type Data_ToolCallContributor_ interface {
	isToolCallContributor()
}

type CompanionStruct_ToolCallContributor_ struct {
}

var Companion_ToolCallContributor_ = CompanionStruct_ToolCallContributor_{}

type ToolCallContributor_ToolCallContributor struct {
	Kind _dafny.Sequence
	Raw  m_ConfluxCodec.Json
}

func (ToolCallContributor_ToolCallContributor) isToolCallContributor() {}

func (CompanionStruct_ToolCallContributor_) Create_ToolCallContributor_(Kind _dafny.Sequence, Raw m_ConfluxCodec.Json) ToolCallContributor {
	return ToolCallContributor{ToolCallContributor_ToolCallContributor{Kind, Raw}}
}

func (_this ToolCallContributor) Is_ToolCallContributor() bool {
	_, ok := _this.Get_().(ToolCallContributor_ToolCallContributor)
	return ok
}

func (CompanionStruct_ToolCallContributor_) Default() ToolCallContributor {
	return Companion_ToolCallContributor_.Create_ToolCallContributor_(_dafny.EmptySeq, m_ConfluxCodec.Companion_Json_.Default())
}

func (_this ToolCallContributor) Dtor_kind() _dafny.Sequence {
	return _this.Get_().(ToolCallContributor_ToolCallContributor).Kind
}

func (_this ToolCallContributor) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(ToolCallContributor_ToolCallContributor).Raw
}

func (_this ToolCallContributor) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ToolCallContributor_ToolCallContributor:
		{
			return "Chat.ToolCallContributor.ToolCallContributor" + "(" + data.Kind.VerbatimString(true) + ", " + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ToolCallContributor) Equals(other ToolCallContributor) bool {
	switch data1 := _this.Get_().(type) {
	case ToolCallContributor_ToolCallContributor:
		{
			data2, ok := other.Get_().(ToolCallContributor_ToolCallContributor)
			return ok && data1.Kind.Equals(data2.Kind) && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ToolCallContributor) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ToolCallContributor)
	return ok && _this.Equals(typed)
}

func Type_ToolCallContributor_() _dafny.TypeDescriptor {
	return type_ToolCallContributor_{}
}

type type_ToolCallContributor_ struct {
}

func (_this type_ToolCallContributor_) Default() interface{} {
	return Companion_ToolCallContributor_.Default()
}

func (_this type_ToolCallContributor_) String() string {
	return "Chat.ToolCallContributor"
}
func (_this ToolCallContributor) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ToolCallContributor{}

// End of datatype ToolCallContributor

// Definition of datatype ToolCall
type ToolCall struct {
	Data_ToolCall_
}

func (_this ToolCall) Get_() Data_ToolCall_ {
	return _this.Data_ToolCall_
}

type Data_ToolCall_ interface {
	isToolCall()
}

type CompanionStruct_ToolCall_ struct {
}

var Companion_ToolCall_ = CompanionStruct_ToolCall_{}

type ToolCall_ToolCall struct {
	ToolCallId        _dafny.Sequence
	ToolName          _dafny.Sequence
	DisplayName       _dafny.Sequence
	Status            _dafny.Sequence
	Intention         m_AhpSkeleton.Option
	Contributor       m_AhpSkeleton.Option
	Meta              m_AhpSkeleton.Option
	InvocationMessage _dafny.Sequence
	ToolInput         m_AhpSkeleton.Option
	ConfirmationTitle m_AhpSkeleton.Option
	RiskAssessment    m_AhpSkeleton.Option
	Edits             m_AhpSkeleton.Option
	Editable          m_AhpSkeleton.Option
	Options           m_AhpSkeleton.Option
	Confirmed         m_AhpSkeleton.Option
	Success           m_AhpSkeleton.Option
	PastTenseMessage  m_AhpSkeleton.Option
	Reason            m_AhpSkeleton.Option
	ReasonMessage     m_AhpSkeleton.Option
	UserSuggestion    m_AhpSkeleton.Option
	SelectedOption    m_AhpSkeleton.Option
	Content           m_AhpSkeleton.Option
	StructuredContent m_AhpSkeleton.Option
	Error             m_AhpSkeleton.Option
	Auth              m_AhpSkeleton.Option
	PartialInput      m_AhpSkeleton.Option
}

func (ToolCall_ToolCall) isToolCall() {}

func (CompanionStruct_ToolCall_) Create_ToolCall_(ToolCallId _dafny.Sequence, ToolName _dafny.Sequence, DisplayName _dafny.Sequence, Status _dafny.Sequence, Intention m_AhpSkeleton.Option, Contributor m_AhpSkeleton.Option, Meta m_AhpSkeleton.Option, InvocationMessage _dafny.Sequence, ToolInput m_AhpSkeleton.Option, ConfirmationTitle m_AhpSkeleton.Option, RiskAssessment m_AhpSkeleton.Option, Edits m_AhpSkeleton.Option, Editable m_AhpSkeleton.Option, Options m_AhpSkeleton.Option, Confirmed m_AhpSkeleton.Option, Success m_AhpSkeleton.Option, PastTenseMessage m_AhpSkeleton.Option, Reason m_AhpSkeleton.Option, ReasonMessage m_AhpSkeleton.Option, UserSuggestion m_AhpSkeleton.Option, SelectedOption m_AhpSkeleton.Option, Content m_AhpSkeleton.Option, StructuredContent m_AhpSkeleton.Option, Error m_AhpSkeleton.Option, Auth m_AhpSkeleton.Option, PartialInput m_AhpSkeleton.Option) ToolCall {
	return ToolCall{ToolCall_ToolCall{ToolCallId, ToolName, DisplayName, Status, Intention, Contributor, Meta, InvocationMessage, ToolInput, ConfirmationTitle, RiskAssessment, Edits, Editable, Options, Confirmed, Success, PastTenseMessage, Reason, ReasonMessage, UserSuggestion, SelectedOption, Content, StructuredContent, Error, Auth, PartialInput}}
}

func (_this ToolCall) Is_ToolCall() bool {
	_, ok := _this.Get_().(ToolCall_ToolCall)
	return ok
}

func (CompanionStruct_ToolCall_) Default() ToolCall {
	return Companion_ToolCall_.Create_ToolCall_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default())
}

func (_this ToolCall) Dtor_toolCallId() _dafny.Sequence {
	return _this.Get_().(ToolCall_ToolCall).ToolCallId
}

func (_this ToolCall) Dtor_toolName() _dafny.Sequence {
	return _this.Get_().(ToolCall_ToolCall).ToolName
}

func (_this ToolCall) Dtor_displayName() _dafny.Sequence {
	return _this.Get_().(ToolCall_ToolCall).DisplayName
}

func (_this ToolCall) Dtor_status() _dafny.Sequence {
	return _this.Get_().(ToolCall_ToolCall).Status
}

func (_this ToolCall) Dtor_intention() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Intention
}

func (_this ToolCall) Dtor_contributor() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Contributor
}

func (_this ToolCall) Dtor_meta() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Meta
}

func (_this ToolCall) Dtor_invocationMessage() _dafny.Sequence {
	return _this.Get_().(ToolCall_ToolCall).InvocationMessage
}

func (_this ToolCall) Dtor_toolInput() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).ToolInput
}

func (_this ToolCall) Dtor_confirmationTitle() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).ConfirmationTitle
}

func (_this ToolCall) Dtor_riskAssessment() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).RiskAssessment
}

func (_this ToolCall) Dtor_edits() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Edits
}

func (_this ToolCall) Dtor_editable() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Editable
}

func (_this ToolCall) Dtor_options() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Options
}

func (_this ToolCall) Dtor_confirmed() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Confirmed
}

func (_this ToolCall) Dtor_success() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Success
}

func (_this ToolCall) Dtor_pastTenseMessage() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).PastTenseMessage
}

func (_this ToolCall) Dtor_reason() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Reason
}

func (_this ToolCall) Dtor_reasonMessage() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).ReasonMessage
}

func (_this ToolCall) Dtor_userSuggestion() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).UserSuggestion
}

func (_this ToolCall) Dtor_selectedOption() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).SelectedOption
}

func (_this ToolCall) Dtor_content() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Content
}

func (_this ToolCall) Dtor_structuredContent() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).StructuredContent
}

func (_this ToolCall) Dtor_error() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Error
}

func (_this ToolCall) Dtor_auth() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).Auth
}

func (_this ToolCall) Dtor_partialInput() m_AhpSkeleton.Option {
	return _this.Get_().(ToolCall_ToolCall).PartialInput
}

func (_this ToolCall) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ToolCall_ToolCall:
		{
			return "Chat.ToolCall.ToolCall" + "(" + data.ToolCallId.VerbatimString(true) + ", " + data.ToolName.VerbatimString(true) + ", " + data.DisplayName.VerbatimString(true) + ", " + data.Status.VerbatimString(true) + ", " + _dafny.String(data.Intention) + ", " + _dafny.String(data.Contributor) + ", " + _dafny.String(data.Meta) + ", " + data.InvocationMessage.VerbatimString(true) + ", " + _dafny.String(data.ToolInput) + ", " + _dafny.String(data.ConfirmationTitle) + ", " + _dafny.String(data.RiskAssessment) + ", " + _dafny.String(data.Edits) + ", " + _dafny.String(data.Editable) + ", " + _dafny.String(data.Options) + ", " + _dafny.String(data.Confirmed) + ", " + _dafny.String(data.Success) + ", " + _dafny.String(data.PastTenseMessage) + ", " + _dafny.String(data.Reason) + ", " + _dafny.String(data.ReasonMessage) + ", " + _dafny.String(data.UserSuggestion) + ", " + _dafny.String(data.SelectedOption) + ", " + _dafny.String(data.Content) + ", " + _dafny.String(data.StructuredContent) + ", " + _dafny.String(data.Error) + ", " + _dafny.String(data.Auth) + ", " + _dafny.String(data.PartialInput) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ToolCall) Equals(other ToolCall) bool {
	switch data1 := _this.Get_().(type) {
	case ToolCall_ToolCall:
		{
			data2, ok := other.Get_().(ToolCall_ToolCall)
			return ok && data1.ToolCallId.Equals(data2.ToolCallId) && data1.ToolName.Equals(data2.ToolName) && data1.DisplayName.Equals(data2.DisplayName) && data1.Status.Equals(data2.Status) && data1.Intention.Equals(data2.Intention) && data1.Contributor.Equals(data2.Contributor) && data1.Meta.Equals(data2.Meta) && data1.InvocationMessage.Equals(data2.InvocationMessage) && data1.ToolInput.Equals(data2.ToolInput) && data1.ConfirmationTitle.Equals(data2.ConfirmationTitle) && data1.RiskAssessment.Equals(data2.RiskAssessment) && data1.Edits.Equals(data2.Edits) && data1.Editable.Equals(data2.Editable) && data1.Options.Equals(data2.Options) && data1.Confirmed.Equals(data2.Confirmed) && data1.Success.Equals(data2.Success) && data1.PastTenseMessage.Equals(data2.PastTenseMessage) && data1.Reason.Equals(data2.Reason) && data1.ReasonMessage.Equals(data2.ReasonMessage) && data1.UserSuggestion.Equals(data2.UserSuggestion) && data1.SelectedOption.Equals(data2.SelectedOption) && data1.Content.Equals(data2.Content) && data1.StructuredContent.Equals(data2.StructuredContent) && data1.Error.Equals(data2.Error) && data1.Auth.Equals(data2.Auth) && data1.PartialInput.Equals(data2.PartialInput)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ToolCall) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ToolCall)
	return ok && _this.Equals(typed)
}

func Type_ToolCall_() _dafny.TypeDescriptor {
	return type_ToolCall_{}
}

type type_ToolCall_ struct {
}

func (_this type_ToolCall_) Default() interface{} {
	return Companion_ToolCall_.Default()
}

func (_this type_ToolCall_) String() string {
	return "Chat.ToolCall"
}
func (_this ToolCall) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ToolCall{}

// End of datatype ToolCall

// Definition of datatype Part
type Part struct {
	Data_Part_
}

func (_this Part) Get_() Data_Part_ {
	return _this.Data_Part_
}

type Data_Part_ interface {
	isPart()
}

type CompanionStruct_Part_ struct {
}

var Companion_Part_ = CompanionStruct_Part_{}

type Part_PMarkdown struct {
	Id      _dafny.Sequence
	Content _dafny.Sequence
}

func (Part_PMarkdown) isPart() {}

func (CompanionStruct_Part_) Create_PMarkdown_(Id _dafny.Sequence, Content _dafny.Sequence) Part {
	return Part{Part_PMarkdown{Id, Content}}
}

func (_this Part) Is_PMarkdown() bool {
	_, ok := _this.Get_().(Part_PMarkdown)
	return ok
}

type Part_PReasoning struct {
	Id      _dafny.Sequence
	Content _dafny.Sequence
}

func (Part_PReasoning) isPart() {}

func (CompanionStruct_Part_) Create_PReasoning_(Id _dafny.Sequence, Content _dafny.Sequence) Part {
	return Part{Part_PReasoning{Id, Content}}
}

func (_this Part) Is_PReasoning() bool {
	_, ok := _this.Get_().(Part_PReasoning)
	return ok
}

type Part_PToolCall struct {
	Tc ToolCall
}

func (Part_PToolCall) isPart() {}

func (CompanionStruct_Part_) Create_PToolCall_(Tc ToolCall) Part {
	return Part{Part_PToolCall{Tc}}
}

func (_this Part) Is_PToolCall() bool {
	_, ok := _this.Get_().(Part_PToolCall)
	return ok
}

type Part_PInputRequest struct {
	Req      InputReq
	Response m_AhpSkeleton.Option
}

func (Part_PInputRequest) isPart() {}

func (CompanionStruct_Part_) Create_PInputRequest_(Req InputReq, Response m_AhpSkeleton.Option) Part {
	return Part{Part_PInputRequest{Req, Response}}
}

func (_this Part) Is_PInputRequest() bool {
	_, ok := _this.Get_().(Part_PInputRequest)
	return ok
}

func (CompanionStruct_Part_) Default() Part {
	return Companion_Part_.Create_PMarkdown_(_dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this Part) Dtor_id() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case Part_PMarkdown:
		return data.Id
	default:
		return data.(Part_PReasoning).Id
	}
}

func (_this Part) Dtor_content() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case Part_PMarkdown:
		return data.Content
	default:
		return data.(Part_PReasoning).Content
	}
}

func (_this Part) Dtor_tc() ToolCall {
	return _this.Get_().(Part_PToolCall).Tc
}

func (_this Part) Dtor_req() InputReq {
	return _this.Get_().(Part_PInputRequest).Req
}

func (_this Part) Dtor_response() m_AhpSkeleton.Option {
	return _this.Get_().(Part_PInputRequest).Response
}

func (_this Part) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Part_PMarkdown:
		{
			return "Chat.Part.PMarkdown" + "(" + data.Id.VerbatimString(true) + ", " + data.Content.VerbatimString(true) + ")"
		}
	case Part_PReasoning:
		{
			return "Chat.Part.PReasoning" + "(" + data.Id.VerbatimString(true) + ", " + data.Content.VerbatimString(true) + ")"
		}
	case Part_PToolCall:
		{
			return "Chat.Part.PToolCall" + "(" + _dafny.String(data.Tc) + ")"
		}
	case Part_PInputRequest:
		{
			return "Chat.Part.PInputRequest" + "(" + _dafny.String(data.Req) + ", " + _dafny.String(data.Response) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Part) Equals(other Part) bool {
	switch data1 := _this.Get_().(type) {
	case Part_PMarkdown:
		{
			data2, ok := other.Get_().(Part_PMarkdown)
			return ok && data1.Id.Equals(data2.Id) && data1.Content.Equals(data2.Content)
		}
	case Part_PReasoning:
		{
			data2, ok := other.Get_().(Part_PReasoning)
			return ok && data1.Id.Equals(data2.Id) && data1.Content.Equals(data2.Content)
		}
	case Part_PToolCall:
		{
			data2, ok := other.Get_().(Part_PToolCall)
			return ok && data1.Tc.Equals(data2.Tc)
		}
	case Part_PInputRequest:
		{
			data2, ok := other.Get_().(Part_PInputRequest)
			return ok && data1.Req.Equals(data2.Req) && data1.Response.Equals(data2.Response)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Part) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Part)
	return ok && _this.Equals(typed)
}

func Type_Part_() _dafny.TypeDescriptor {
	return type_Part_{}
}

type type_Part_ struct {
}

func (_this type_Part_) Default() interface{} {
	return Companion_Part_.Default()
}

func (_this type_Part_) String() string {
	return "Chat.Part"
}
func (_this Part) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Part{}

// End of datatype Part

// Definition of datatype Turn
type Turn struct {
	Data_Turn_
}

func (_this Turn) Get_() Data_Turn_ {
	return _this.Data_Turn_
}

type Data_Turn_ interface {
	isTurn()
}

type CompanionStruct_Turn_ struct {
}

var Companion_Turn_ = CompanionStruct_Turn_{}

type Turn_Turn struct {
	TurnId  _dafny.Sequence
	Message m_ConfluxCodec.Json
	Parts   _dafny.Sequence
	State   _dafny.Sequence
	Usage   m_AhpSkeleton.Option
	Error   m_AhpSkeleton.Option
}

func (Turn_Turn) isTurn() {}

func (CompanionStruct_Turn_) Create_Turn_(TurnId _dafny.Sequence, Message m_ConfluxCodec.Json, Parts _dafny.Sequence, State _dafny.Sequence, Usage m_AhpSkeleton.Option, Error m_AhpSkeleton.Option) Turn {
	return Turn{Turn_Turn{TurnId, Message, Parts, State, Usage, Error}}
}

func (_this Turn) Is_Turn() bool {
	_, ok := _this.Get_().(Turn_Turn)
	return ok
}

func (CompanionStruct_Turn_) Default() Turn {
	return Companion_Turn_.Create_Turn_(_dafny.EmptySeq, m_ConfluxCodec.Companion_Json_.Default(), _dafny.EmptySeq, _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default())
}

func (_this Turn) Dtor_turnId() _dafny.Sequence {
	return _this.Get_().(Turn_Turn).TurnId
}

func (_this Turn) Dtor_message() m_ConfluxCodec.Json {
	return _this.Get_().(Turn_Turn).Message
}

func (_this Turn) Dtor_parts() _dafny.Sequence {
	return _this.Get_().(Turn_Turn).Parts
}

func (_this Turn) Dtor_state() _dafny.Sequence {
	return _this.Get_().(Turn_Turn).State
}

func (_this Turn) Dtor_usage() m_AhpSkeleton.Option {
	return _this.Get_().(Turn_Turn).Usage
}

func (_this Turn) Dtor_error() m_AhpSkeleton.Option {
	return _this.Get_().(Turn_Turn).Error
}

func (_this Turn) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Turn_Turn:
		{
			return "Chat.Turn.Turn" + "(" + data.TurnId.VerbatimString(true) + ", " + _dafny.String(data.Message) + ", " + _dafny.String(data.Parts) + ", " + data.State.VerbatimString(true) + ", " + _dafny.String(data.Usage) + ", " + _dafny.String(data.Error) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Turn) Equals(other Turn) bool {
	switch data1 := _this.Get_().(type) {
	case Turn_Turn:
		{
			data2, ok := other.Get_().(Turn_Turn)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.Message.Equals(data2.Message) && data1.Parts.Equals(data2.Parts) && data1.State.Equals(data2.State) && data1.Usage.Equals(data2.Usage) && data1.Error.Equals(data2.Error)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Turn) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Turn)
	return ok && _this.Equals(typed)
}

func Type_Turn_() _dafny.TypeDescriptor {
	return type_Turn_{}
}

type type_Turn_ struct {
}

func (_this type_Turn_) Default() interface{} {
	return Companion_Turn_.Default()
}

func (_this type_Turn_) String() string {
	return "Chat.Turn"
}
func (_this Turn) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Turn{}

// End of datatype Turn

// Definition of datatype QMsg
type QMsg struct {
	Data_QMsg_
}

func (_this QMsg) Get_() Data_QMsg_ {
	return _this.Data_QMsg_
}

type Data_QMsg_ interface {
	isQMsg()
}

type CompanionStruct_QMsg_ struct {
}

var Companion_QMsg_ = CompanionStruct_QMsg_{}

type QMsg_QMsg struct {
	Id  _dafny.Sequence
	Raw m_ConfluxCodec.Json
}

func (QMsg_QMsg) isQMsg() {}

func (CompanionStruct_QMsg_) Create_QMsg_(Id _dafny.Sequence, Raw m_ConfluxCodec.Json) QMsg {
	return QMsg{QMsg_QMsg{Id, Raw}}
}

func (_this QMsg) Is_QMsg() bool {
	_, ok := _this.Get_().(QMsg_QMsg)
	return ok
}

func (CompanionStruct_QMsg_) Default() QMsg {
	return Companion_QMsg_.Create_QMsg_(_dafny.EmptySeq, m_ConfluxCodec.Companion_Json_.Default())
}

func (_this QMsg) Dtor_id() _dafny.Sequence {
	return _this.Get_().(QMsg_QMsg).Id
}

func (_this QMsg) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(QMsg_QMsg).Raw
}

func (_this QMsg) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case QMsg_QMsg:
		{
			return "Chat.QMsg.QMsg" + "(" + data.Id.VerbatimString(true) + ", " + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this QMsg) Equals(other QMsg) bool {
	switch data1 := _this.Get_().(type) {
	case QMsg_QMsg:
		{
			data2, ok := other.Get_().(QMsg_QMsg)
			return ok && data1.Id.Equals(data2.Id) && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this QMsg) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(QMsg)
	return ok && _this.Equals(typed)
}

func Type_QMsg_() _dafny.TypeDescriptor {
	return type_QMsg_{}
}

type type_QMsg_ struct {
}

func (_this type_QMsg_) Default() interface{} {
	return Companion_QMsg_.Default()
}

func (_this type_QMsg_) String() string {
	return "Chat.QMsg"
}
func (_this QMsg) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = QMsg{}

// End of datatype QMsg

// Definition of datatype InputReq
type InputReq struct {
	Data_InputReq_
}

func (_this InputReq) Get_() Data_InputReq_ {
	return _this.Data_InputReq_
}

type Data_InputReq_ interface {
	isInputReq()
}

type CompanionStruct_InputReq_ struct {
}

var Companion_InputReq_ = CompanionStruct_InputReq_{}

type InputReq_InputReq struct {
	Id      _dafny.Sequence
	Answers _dafny.Map
	Open    _dafny.Map
}

func (InputReq_InputReq) isInputReq() {}

func (CompanionStruct_InputReq_) Create_InputReq_(Id _dafny.Sequence, Answers _dafny.Map, Open _dafny.Map) InputReq {
	return InputReq{InputReq_InputReq{Id, Answers, Open}}
}

func (_this InputReq) Is_InputReq() bool {
	_, ok := _this.Get_().(InputReq_InputReq)
	return ok
}

func (CompanionStruct_InputReq_) Default() InputReq {
	return Companion_InputReq_.Create_InputReq_(_dafny.EmptySeq, _dafny.EmptyMap, _dafny.EmptyMap)
}

func (_this InputReq) Dtor_id() _dafny.Sequence {
	return _this.Get_().(InputReq_InputReq).Id
}

func (_this InputReq) Dtor_answers() _dafny.Map {
	return _this.Get_().(InputReq_InputReq).Answers
}

func (_this InputReq) Dtor_open() _dafny.Map {
	return _this.Get_().(InputReq_InputReq).Open
}

func (_this InputReq) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case InputReq_InputReq:
		{
			return "Chat.InputReq.InputReq" + "(" + data.Id.VerbatimString(true) + ", " + _dafny.String(data.Answers) + ", " + _dafny.String(data.Open) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this InputReq) Equals(other InputReq) bool {
	switch data1 := _this.Get_().(type) {
	case InputReq_InputReq:
		{
			data2, ok := other.Get_().(InputReq_InputReq)
			return ok && data1.Id.Equals(data2.Id) && data1.Answers.Equals(data2.Answers) && data1.Open.Equals(data2.Open)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this InputReq) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(InputReq)
	return ok && _this.Equals(typed)
}

func Type_InputReq_() _dafny.TypeDescriptor {
	return type_InputReq_{}
}

type type_InputReq_ struct {
}

func (_this type_InputReq_) Default() interface{} {
	return Companion_InputReq_.Default()
}

func (_this type_InputReq_) String() string {
	return "Chat.InputReq"
}
func (_this InputReq) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = InputReq{}

// End of datatype InputReq

// Definition of datatype ChatState
type ChatState struct {
	Data_ChatState_
}

func (_this ChatState) Get_() Data_ChatState_ {
	return _this.Data_ChatState_
}

type Data_ChatState_ interface {
	isChatState()
}

type CompanionStruct_ChatState_ struct {
}

var Companion_ChatState_ = CompanionStruct_ChatState_{}

type ChatState_ChatState struct {
	Turns           _dafny.Sequence
	Title           _dafny.Sequence
	Status          uint32
	ModifiedAt      _dafny.Sequence
	Draft           m_AhpSkeleton.Option
	Activity        m_AhpSkeleton.Option
	ActiveTurn      m_AhpSkeleton.Option
	SteeringMessage m_AhpSkeleton.Option
	QueuedMessages  _dafny.Sequence
	NextCursor      m_AhpSkeleton.Option
}

func (ChatState_ChatState) isChatState() {}

func (CompanionStruct_ChatState_) Create_ChatState_(Turns _dafny.Sequence, Title _dafny.Sequence, Status uint32, ModifiedAt _dafny.Sequence, Draft m_AhpSkeleton.Option, Activity m_AhpSkeleton.Option, ActiveTurn m_AhpSkeleton.Option, SteeringMessage m_AhpSkeleton.Option, QueuedMessages _dafny.Sequence, NextCursor m_AhpSkeleton.Option) ChatState {
	return ChatState{ChatState_ChatState{Turns, Title, Status, ModifiedAt, Draft, Activity, ActiveTurn, SteeringMessage, QueuedMessages, NextCursor}}
}

func (_this ChatState) Is_ChatState() bool {
	_, ok := _this.Get_().(ChatState_ChatState)
	return ok
}

func (CompanionStruct_ChatState_) Default() ChatState {
	return Companion_ChatState_.Create_ChatState_(_dafny.EmptySeq, _dafny.EmptySeq, 0, _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default())
}

func (_this ChatState) Dtor_turns() _dafny.Sequence {
	return _this.Get_().(ChatState_ChatState).Turns
}

func (_this ChatState) Dtor_title() _dafny.Sequence {
	return _this.Get_().(ChatState_ChatState).Title
}

func (_this ChatState) Dtor_status() uint32 {
	return _this.Get_().(ChatState_ChatState).Status
}

func (_this ChatState) Dtor_modifiedAt() _dafny.Sequence {
	return _this.Get_().(ChatState_ChatState).ModifiedAt
}

func (_this ChatState) Dtor_draft() m_AhpSkeleton.Option {
	return _this.Get_().(ChatState_ChatState).Draft
}

func (_this ChatState) Dtor_activity() m_AhpSkeleton.Option {
	return _this.Get_().(ChatState_ChatState).Activity
}

func (_this ChatState) Dtor_activeTurn() m_AhpSkeleton.Option {
	return _this.Get_().(ChatState_ChatState).ActiveTurn
}

func (_this ChatState) Dtor_steeringMessage() m_AhpSkeleton.Option {
	return _this.Get_().(ChatState_ChatState).SteeringMessage
}

func (_this ChatState) Dtor_queuedMessages() _dafny.Sequence {
	return _this.Get_().(ChatState_ChatState).QueuedMessages
}

func (_this ChatState) Dtor_nextCursor() m_AhpSkeleton.Option {
	return _this.Get_().(ChatState_ChatState).NextCursor
}

func (_this ChatState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ChatState_ChatState:
		{
			return "Chat.ChatState.ChatState" + "(" + _dafny.String(data.Turns) + ", " + data.Title.VerbatimString(true) + ", " + _dafny.String(data.Status) + ", " + data.ModifiedAt.VerbatimString(true) + ", " + _dafny.String(data.Draft) + ", " + _dafny.String(data.Activity) + ", " + _dafny.String(data.ActiveTurn) + ", " + _dafny.String(data.SteeringMessage) + ", " + _dafny.String(data.QueuedMessages) + ", " + _dafny.String(data.NextCursor) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ChatState) Equals(other ChatState) bool {
	switch data1 := _this.Get_().(type) {
	case ChatState_ChatState:
		{
			data2, ok := other.Get_().(ChatState_ChatState)
			return ok && data1.Turns.Equals(data2.Turns) && data1.Title.Equals(data2.Title) && data1.Status == data2.Status && data1.ModifiedAt.Equals(data2.ModifiedAt) && data1.Draft.Equals(data2.Draft) && data1.Activity.Equals(data2.Activity) && data1.ActiveTurn.Equals(data2.ActiveTurn) && data1.SteeringMessage.Equals(data2.SteeringMessage) && data1.QueuedMessages.Equals(data2.QueuedMessages) && data1.NextCursor.Equals(data2.NextCursor)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ChatState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ChatState)
	return ok && _this.Equals(typed)
}

func Type_ChatState_() _dafny.TypeDescriptor {
	return type_ChatState_{}
}

type type_ChatState_ struct {
}

func (_this type_ChatState_) Default() interface{} {
	return Companion_ChatState_.Default()
}

func (_this type_ChatState_) String() string {
	return "Chat.ChatState"
}
func (_this ChatState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ChatState{}

// End of datatype ChatState

// Definition of datatype ChatAction
type ChatAction struct {
	Data_ChatAction_
}

func (_this ChatAction) Get_() Data_ChatAction_ {
	return _this.Data_ChatAction_
}

type Data_ChatAction_ interface {
	isChatAction()
}

type CompanionStruct_ChatAction_ struct {
}

var Companion_ChatAction_ = CompanionStruct_ChatAction_{}

type ChatAction_CDraftChanged struct {
	Draft m_AhpSkeleton.Option
}

func (ChatAction_CDraftChanged) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CDraftChanged_(Draft m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CDraftChanged{Draft}}
}

func (_this ChatAction) Is_CDraftChanged() bool {
	_, ok := _this.Get_().(ChatAction_CDraftChanged)
	return ok
}

type ChatAction_CActivityChanged struct {
	Activity m_AhpSkeleton.Option
}

func (ChatAction_CActivityChanged) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CActivityChanged_(Activity m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CActivityChanged{Activity}}
}

func (_this ChatAction) Is_CActivityChanged() bool {
	_, ok := _this.Get_().(ChatAction_CActivityChanged)
	return ok
}

type ChatAction_CTurnStarted struct {
	TurnId          _dafny.Sequence
	Message         m_ConfluxCodec.Json
	QueuedMessageId m_AhpSkeleton.Option
}

func (ChatAction_CTurnStarted) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CTurnStarted_(TurnId _dafny.Sequence, Message m_ConfluxCodec.Json, QueuedMessageId m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CTurnStarted{TurnId, Message, QueuedMessageId}}
}

func (_this ChatAction) Is_CTurnStarted() bool {
	_, ok := _this.Get_().(ChatAction_CTurnStarted)
	return ok
}

type ChatAction_CTurnComplete struct {
	TurnId _dafny.Sequence
}

func (ChatAction_CTurnComplete) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CTurnComplete_(TurnId _dafny.Sequence) ChatAction {
	return ChatAction{ChatAction_CTurnComplete{TurnId}}
}

func (_this ChatAction) Is_CTurnComplete() bool {
	_, ok := _this.Get_().(ChatAction_CTurnComplete)
	return ok
}

type ChatAction_CTurnCancelled struct {
	TurnId _dafny.Sequence
}

func (ChatAction_CTurnCancelled) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CTurnCancelled_(TurnId _dafny.Sequence) ChatAction {
	return ChatAction{ChatAction_CTurnCancelled{TurnId}}
}

func (_this ChatAction) Is_CTurnCancelled() bool {
	_, ok := _this.Get_().(ChatAction_CTurnCancelled)
	return ok
}

type ChatAction_CUsage struct {
	TurnId _dafny.Sequence
	Usage  m_ConfluxCodec.Json
}

func (ChatAction_CUsage) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CUsage_(TurnId _dafny.Sequence, Usage m_ConfluxCodec.Json) ChatAction {
	return ChatAction{ChatAction_CUsage{TurnId, Usage}}
}

func (_this ChatAction) Is_CUsage() bool {
	_, ok := _this.Get_().(ChatAction_CUsage)
	return ok
}

type ChatAction_CError struct {
	TurnId _dafny.Sequence
	Err    m_ConfluxCodec.Json
}

func (ChatAction_CError) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CError_(TurnId _dafny.Sequence, Err m_ConfluxCodec.Json) ChatAction {
	return ChatAction{ChatAction_CError{TurnId, Err}}
}

func (_this ChatAction) Is_CError() bool {
	_, ok := _this.Get_().(ChatAction_CError)
	return ok
}

type ChatAction_CResponsePart struct {
	TurnId _dafny.Sequence
	Part   Part
}

func (ChatAction_CResponsePart) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CResponsePart_(TurnId _dafny.Sequence, Part Part) ChatAction {
	return ChatAction{ChatAction_CResponsePart{TurnId, Part}}
}

func (_this ChatAction) Is_CResponsePart() bool {
	_, ok := _this.Get_().(ChatAction_CResponsePart)
	return ok
}

type ChatAction_CDelta struct {
	TurnId  _dafny.Sequence
	PartId  _dafny.Sequence
	Content _dafny.Sequence
}

func (ChatAction_CDelta) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CDelta_(TurnId _dafny.Sequence, PartId _dafny.Sequence, Content _dafny.Sequence) ChatAction {
	return ChatAction{ChatAction_CDelta{TurnId, PartId, Content}}
}

func (_this ChatAction) Is_CDelta() bool {
	_, ok := _this.Get_().(ChatAction_CDelta)
	return ok
}

type ChatAction_CToolCallStart struct {
	TurnId      _dafny.Sequence
	ToolCallId  _dafny.Sequence
	ToolName    _dafny.Sequence
	DisplayName _dafny.Sequence
	Intention   m_AhpSkeleton.Option
	Contributor m_AhpSkeleton.Option
	Meta        m_AhpSkeleton.Option
}

func (ChatAction_CToolCallStart) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CToolCallStart_(TurnId _dafny.Sequence, ToolCallId _dafny.Sequence, ToolName _dafny.Sequence, DisplayName _dafny.Sequence, Intention m_AhpSkeleton.Option, Contributor m_AhpSkeleton.Option, Meta m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CToolCallStart{TurnId, ToolCallId, ToolName, DisplayName, Intention, Contributor, Meta}}
}

func (_this ChatAction) Is_CToolCallStart() bool {
	_, ok := _this.Get_().(ChatAction_CToolCallStart)
	return ok
}

type ChatAction_CToolCallDelta struct {
	TurnId            _dafny.Sequence
	ToolCallId        _dafny.Sequence
	Content           _dafny.Sequence
	InvocationMessage m_AhpSkeleton.Option
	Meta              m_AhpSkeleton.Option
}

func (ChatAction_CToolCallDelta) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CToolCallDelta_(TurnId _dafny.Sequence, ToolCallId _dafny.Sequence, Content _dafny.Sequence, InvocationMessage m_AhpSkeleton.Option, Meta m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CToolCallDelta{TurnId, ToolCallId, Content, InvocationMessage, Meta}}
}

func (_this ChatAction) Is_CToolCallDelta() bool {
	_, ok := _this.Get_().(ChatAction_CToolCallDelta)
	return ok
}

type ChatAction_CToolCallReady struct {
	TurnId            _dafny.Sequence
	ToolCallId        _dafny.Sequence
	InvocationMessage m_AhpSkeleton.Option
	ToolInput         m_AhpSkeleton.Option
	ConfirmationTitle m_AhpSkeleton.Option
	RiskAssessment    m_AhpSkeleton.Option
	Edits             m_AhpSkeleton.Option
	Editable          m_AhpSkeleton.Option
	Options           m_AhpSkeleton.Option
	Confirmed         m_AhpSkeleton.Option
	Meta              m_AhpSkeleton.Option
}

func (ChatAction_CToolCallReady) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CToolCallReady_(TurnId _dafny.Sequence, ToolCallId _dafny.Sequence, InvocationMessage m_AhpSkeleton.Option, ToolInput m_AhpSkeleton.Option, ConfirmationTitle m_AhpSkeleton.Option, RiskAssessment m_AhpSkeleton.Option, Edits m_AhpSkeleton.Option, Editable m_AhpSkeleton.Option, Options m_AhpSkeleton.Option, Confirmed m_AhpSkeleton.Option, Meta m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CToolCallReady{TurnId, ToolCallId, InvocationMessage, ToolInput, ConfirmationTitle, RiskAssessment, Edits, Editable, Options, Confirmed, Meta}}
}

func (_this ChatAction) Is_CToolCallReady() bool {
	_, ok := _this.Get_().(ChatAction_CToolCallReady)
	return ok
}

type ChatAction_CToolCallConfirmed struct {
	TurnId           _dafny.Sequence
	ToolCallId       _dafny.Sequence
	Approved         bool
	ConfirmedVal     m_AhpSkeleton.Option
	Reason           m_AhpSkeleton.Option
	ReasonMessage    m_AhpSkeleton.Option
	UserSuggestion   m_AhpSkeleton.Option
	EditedToolInput  m_AhpSkeleton.Option
	SelectedOptionId m_AhpSkeleton.Option
	Meta             m_AhpSkeleton.Option
}

func (ChatAction_CToolCallConfirmed) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CToolCallConfirmed_(TurnId _dafny.Sequence, ToolCallId _dafny.Sequence, Approved bool, ConfirmedVal m_AhpSkeleton.Option, Reason m_AhpSkeleton.Option, ReasonMessage m_AhpSkeleton.Option, UserSuggestion m_AhpSkeleton.Option, EditedToolInput m_AhpSkeleton.Option, SelectedOptionId m_AhpSkeleton.Option, Meta m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CToolCallConfirmed{TurnId, ToolCallId, Approved, ConfirmedVal, Reason, ReasonMessage, UserSuggestion, EditedToolInput, SelectedOptionId, Meta}}
}

func (_this ChatAction) Is_CToolCallConfirmed() bool {
	_, ok := _this.Get_().(ChatAction_CToolCallConfirmed)
	return ok
}

type ChatAction_CToolCallAuthRequired struct {
	TurnId     _dafny.Sequence
	ToolCallId _dafny.Sequence
	Auth       m_ConfluxCodec.Json
	Meta       m_AhpSkeleton.Option
}

func (ChatAction_CToolCallAuthRequired) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CToolCallAuthRequired_(TurnId _dafny.Sequence, ToolCallId _dafny.Sequence, Auth m_ConfluxCodec.Json, Meta m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CToolCallAuthRequired{TurnId, ToolCallId, Auth, Meta}}
}

func (_this ChatAction) Is_CToolCallAuthRequired() bool {
	_, ok := _this.Get_().(ChatAction_CToolCallAuthRequired)
	return ok
}

type ChatAction_CToolCallAuthResolved struct {
	TurnId     _dafny.Sequence
	ToolCallId _dafny.Sequence
	Meta       m_AhpSkeleton.Option
}

func (ChatAction_CToolCallAuthResolved) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CToolCallAuthResolved_(TurnId _dafny.Sequence, ToolCallId _dafny.Sequence, Meta m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CToolCallAuthResolved{TurnId, ToolCallId, Meta}}
}

func (_this ChatAction) Is_CToolCallAuthResolved() bool {
	_, ok := _this.Get_().(ChatAction_CToolCallAuthResolved)
	return ok
}

type ChatAction_CToolCallComplete struct {
	TurnId                     _dafny.Sequence
	ToolCallId                 _dafny.Sequence
	Success                    bool
	PastTenseMessage           m_AhpSkeleton.Option
	ResultContent              m_AhpSkeleton.Option
	StructuredContent          m_AhpSkeleton.Option
	Error                      m_AhpSkeleton.Option
	RequiresResultConfirmation bool
	Meta                       m_AhpSkeleton.Option
}

func (ChatAction_CToolCallComplete) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CToolCallComplete_(TurnId _dafny.Sequence, ToolCallId _dafny.Sequence, Success bool, PastTenseMessage m_AhpSkeleton.Option, ResultContent m_AhpSkeleton.Option, StructuredContent m_AhpSkeleton.Option, Error m_AhpSkeleton.Option, RequiresResultConfirmation bool, Meta m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CToolCallComplete{TurnId, ToolCallId, Success, PastTenseMessage, ResultContent, StructuredContent, Error, RequiresResultConfirmation, Meta}}
}

func (_this ChatAction) Is_CToolCallComplete() bool {
	_, ok := _this.Get_().(ChatAction_CToolCallComplete)
	return ok
}

type ChatAction_CToolCallResultConfirmed struct {
	TurnId     _dafny.Sequence
	ToolCallId _dafny.Sequence
	Approved   bool
	Meta       m_AhpSkeleton.Option
}

func (ChatAction_CToolCallResultConfirmed) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CToolCallResultConfirmed_(TurnId _dafny.Sequence, ToolCallId _dafny.Sequence, Approved bool, Meta m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CToolCallResultConfirmed{TurnId, ToolCallId, Approved, Meta}}
}

func (_this ChatAction) Is_CToolCallResultConfirmed() bool {
	_, ok := _this.Get_().(ChatAction_CToolCallResultConfirmed)
	return ok
}

type ChatAction_CToolCallContentChanged struct {
	TurnId     _dafny.Sequence
	ToolCallId _dafny.Sequence
	NewContent m_ConfluxCodec.Json
	Meta       m_AhpSkeleton.Option
}

func (ChatAction_CToolCallContentChanged) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CToolCallContentChanged_(TurnId _dafny.Sequence, ToolCallId _dafny.Sequence, NewContent m_ConfluxCodec.Json, Meta m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CToolCallContentChanged{TurnId, ToolCallId, NewContent, Meta}}
}

func (_this ChatAction) Is_CToolCallContentChanged() bool {
	_, ok := _this.Get_().(ChatAction_CToolCallContentChanged)
	return ok
}

type ChatAction_CReasoning struct {
	TurnId  _dafny.Sequence
	PartId  _dafny.Sequence
	Content _dafny.Sequence
}

func (ChatAction_CReasoning) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CReasoning_(TurnId _dafny.Sequence, PartId _dafny.Sequence, Content _dafny.Sequence) ChatAction {
	return ChatAction{ChatAction_CReasoning{TurnId, PartId, Content}}
}

func (_this ChatAction) Is_CReasoning() bool {
	_, ok := _this.Get_().(ChatAction_CReasoning)
	return ok
}

type ChatAction_CTruncated struct {
	UpToTurnId m_AhpSkeleton.Option
}

func (ChatAction_CTruncated) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CTruncated_(UpToTurnId m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CTruncated{UpToTurnId}}
}

func (_this ChatAction) Is_CTruncated() bool {
	_, ok := _this.Get_().(ChatAction_CTruncated)
	return ok
}

type ChatAction_CQueuedReordered struct {
	Order _dafny.Sequence
}

func (ChatAction_CQueuedReordered) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CQueuedReordered_(Order _dafny.Sequence) ChatAction {
	return ChatAction{ChatAction_CQueuedReordered{Order}}
}

func (_this ChatAction) Is_CQueuedReordered() bool {
	_, ok := _this.Get_().(ChatAction_CQueuedReordered)
	return ok
}

type ChatAction_CTurnsLoaded struct {
	Loaded _dafny.Sequence
	Cursor m_AhpSkeleton.Option
}

func (ChatAction_CTurnsLoaded) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CTurnsLoaded_(Loaded _dafny.Sequence, Cursor m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CTurnsLoaded{Loaded, Cursor}}
}

func (_this ChatAction) Is_CTurnsLoaded() bool {
	_, ok := _this.Get_().(ChatAction_CTurnsLoaded)
	return ok
}

type ChatAction_CInputRequested struct {
	Req InputReq
}

func (ChatAction_CInputRequested) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CInputRequested_(Req InputReq) ChatAction {
	return ChatAction{ChatAction_CInputRequested{Req}}
}

func (_this ChatAction) Is_CInputRequested() bool {
	_, ok := _this.Get_().(ChatAction_CInputRequested)
	return ok
}

type ChatAction_CInputAnswerChanged struct {
	RequestId  _dafny.Sequence
	QuestionId _dafny.Sequence
	Answer     m_AhpSkeleton.Option
}

func (ChatAction_CInputAnswerChanged) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CInputAnswerChanged_(RequestId _dafny.Sequence, QuestionId _dafny.Sequence, Answer m_AhpSkeleton.Option) ChatAction {
	return ChatAction{ChatAction_CInputAnswerChanged{RequestId, QuestionId, Answer}}
}

func (_this ChatAction) Is_CInputAnswerChanged() bool {
	_, ok := _this.Get_().(ChatAction_CInputAnswerChanged)
	return ok
}

type ChatAction_CInputCompleted struct {
	RequestId _dafny.Sequence
	Response  _dafny.Sequence
	Answers   _dafny.Map
}

func (ChatAction_CInputCompleted) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CInputCompleted_(RequestId _dafny.Sequence, Response _dafny.Sequence, Answers _dafny.Map) ChatAction {
	return ChatAction{ChatAction_CInputCompleted{RequestId, Response, Answers}}
}

func (_this ChatAction) Is_CInputCompleted() bool {
	_, ok := _this.Get_().(ChatAction_CInputCompleted)
	return ok
}

type ChatAction_CPendingMessageSet struct {
	Kind    _dafny.Sequence
	Id      _dafny.Sequence
	Message m_ConfluxCodec.Json
}

func (ChatAction_CPendingMessageSet) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CPendingMessageSet_(Kind _dafny.Sequence, Id _dafny.Sequence, Message m_ConfluxCodec.Json) ChatAction {
	return ChatAction{ChatAction_CPendingMessageSet{Kind, Id, Message}}
}

func (_this ChatAction) Is_CPendingMessageSet() bool {
	_, ok := _this.Get_().(ChatAction_CPendingMessageSet)
	return ok
}

type ChatAction_CPendingMessageRemoved struct {
	Kind _dafny.Sequence
	Id   _dafny.Sequence
}

func (ChatAction_CPendingMessageRemoved) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CPendingMessageRemoved_(Kind _dafny.Sequence, Id _dafny.Sequence) ChatAction {
	return ChatAction{ChatAction_CPendingMessageRemoved{Kind, Id}}
}

func (_this ChatAction) Is_CPendingMessageRemoved() bool {
	_, ok := _this.Get_().(ChatAction_CPendingMessageRemoved)
	return ok
}

type ChatAction_CUnknown struct {
	Raw m_ConfluxCodec.Json
}

func (ChatAction_CUnknown) isChatAction() {}

func (CompanionStruct_ChatAction_) Create_CUnknown_(Raw m_ConfluxCodec.Json) ChatAction {
	return ChatAction{ChatAction_CUnknown{Raw}}
}

func (_this ChatAction) Is_CUnknown() bool {
	_, ok := _this.Get_().(ChatAction_CUnknown)
	return ok
}

func (CompanionStruct_ChatAction_) Default() ChatAction {
	return Companion_ChatAction_.Create_CDraftChanged_(m_AhpSkeleton.Companion_Option_.Default())
}

func (_this ChatAction) Dtor_draft() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CDraftChanged).Draft
}

func (_this ChatAction) Dtor_activity() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CActivityChanged).Activity
}

func (_this ChatAction) Dtor_turnId() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case ChatAction_CTurnStarted:
		return data.TurnId
	case ChatAction_CTurnComplete:
		return data.TurnId
	case ChatAction_CTurnCancelled:
		return data.TurnId
	case ChatAction_CUsage:
		return data.TurnId
	case ChatAction_CError:
		return data.TurnId
	case ChatAction_CResponsePart:
		return data.TurnId
	case ChatAction_CDelta:
		return data.TurnId
	case ChatAction_CToolCallStart:
		return data.TurnId
	case ChatAction_CToolCallDelta:
		return data.TurnId
	case ChatAction_CToolCallReady:
		return data.TurnId
	case ChatAction_CToolCallConfirmed:
		return data.TurnId
	case ChatAction_CToolCallAuthRequired:
		return data.TurnId
	case ChatAction_CToolCallAuthResolved:
		return data.TurnId
	case ChatAction_CToolCallComplete:
		return data.TurnId
	case ChatAction_CToolCallResultConfirmed:
		return data.TurnId
	case ChatAction_CToolCallContentChanged:
		return data.TurnId
	default:
		return data.(ChatAction_CReasoning).TurnId
	}
}

func (_this ChatAction) Dtor_message() m_ConfluxCodec.Json {
	switch data := _this.Get_().(type) {
	case ChatAction_CTurnStarted:
		return data.Message
	default:
		return data.(ChatAction_CPendingMessageSet).Message
	}
}

func (_this ChatAction) Dtor_queuedMessageId() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CTurnStarted).QueuedMessageId
}

func (_this ChatAction) Dtor_usage() m_ConfluxCodec.Json {
	return _this.Get_().(ChatAction_CUsage).Usage
}

func (_this ChatAction) Dtor_err() m_ConfluxCodec.Json {
	return _this.Get_().(ChatAction_CError).Err
}

func (_this ChatAction) Dtor_part() Part {
	return _this.Get_().(ChatAction_CResponsePart).Part
}

func (_this ChatAction) Dtor_partId() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case ChatAction_CDelta:
		return data.PartId
	default:
		return data.(ChatAction_CReasoning).PartId
	}
}

func (_this ChatAction) Dtor_content() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case ChatAction_CDelta:
		return data.Content
	case ChatAction_CToolCallDelta:
		return data.Content
	default:
		return data.(ChatAction_CReasoning).Content
	}
}

func (_this ChatAction) Dtor_toolCallId() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case ChatAction_CToolCallStart:
		return data.ToolCallId
	case ChatAction_CToolCallDelta:
		return data.ToolCallId
	case ChatAction_CToolCallReady:
		return data.ToolCallId
	case ChatAction_CToolCallConfirmed:
		return data.ToolCallId
	case ChatAction_CToolCallAuthRequired:
		return data.ToolCallId
	case ChatAction_CToolCallAuthResolved:
		return data.ToolCallId
	case ChatAction_CToolCallComplete:
		return data.ToolCallId
	case ChatAction_CToolCallResultConfirmed:
		return data.ToolCallId
	default:
		return data.(ChatAction_CToolCallContentChanged).ToolCallId
	}
}

func (_this ChatAction) Dtor_toolName() _dafny.Sequence {
	return _this.Get_().(ChatAction_CToolCallStart).ToolName
}

func (_this ChatAction) Dtor_displayName() _dafny.Sequence {
	return _this.Get_().(ChatAction_CToolCallStart).DisplayName
}

func (_this ChatAction) Dtor_intention() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallStart).Intention
}

func (_this ChatAction) Dtor_contributor() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallStart).Contributor
}

func (_this ChatAction) Dtor_meta() m_AhpSkeleton.Option {
	switch data := _this.Get_().(type) {
	case ChatAction_CToolCallStart:
		return data.Meta
	case ChatAction_CToolCallDelta:
		return data.Meta
	case ChatAction_CToolCallReady:
		return data.Meta
	case ChatAction_CToolCallConfirmed:
		return data.Meta
	case ChatAction_CToolCallAuthRequired:
		return data.Meta
	case ChatAction_CToolCallAuthResolved:
		return data.Meta
	case ChatAction_CToolCallComplete:
		return data.Meta
	case ChatAction_CToolCallResultConfirmed:
		return data.Meta
	default:
		return data.(ChatAction_CToolCallContentChanged).Meta
	}
}

func (_this ChatAction) Dtor_invocationMessage() m_AhpSkeleton.Option {
	switch data := _this.Get_().(type) {
	case ChatAction_CToolCallDelta:
		return data.InvocationMessage
	default:
		return data.(ChatAction_CToolCallReady).InvocationMessage
	}
}

func (_this ChatAction) Dtor_toolInput() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallReady).ToolInput
}

func (_this ChatAction) Dtor_confirmationTitle() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallReady).ConfirmationTitle
}

func (_this ChatAction) Dtor_riskAssessment() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallReady).RiskAssessment
}

func (_this ChatAction) Dtor_edits() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallReady).Edits
}

func (_this ChatAction) Dtor_editable() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallReady).Editable
}

func (_this ChatAction) Dtor_options() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallReady).Options
}

func (_this ChatAction) Dtor_confirmed() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallReady).Confirmed
}

func (_this ChatAction) Dtor_approved() bool {
	switch data := _this.Get_().(type) {
	case ChatAction_CToolCallConfirmed:
		return data.Approved
	default:
		return data.(ChatAction_CToolCallResultConfirmed).Approved
	}
}

func (_this ChatAction) Dtor_confirmedVal() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallConfirmed).ConfirmedVal
}

func (_this ChatAction) Dtor_reason() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallConfirmed).Reason
}

func (_this ChatAction) Dtor_reasonMessage() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallConfirmed).ReasonMessage
}

func (_this ChatAction) Dtor_userSuggestion() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallConfirmed).UserSuggestion
}

func (_this ChatAction) Dtor_editedToolInput() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallConfirmed).EditedToolInput
}

func (_this ChatAction) Dtor_selectedOptionId() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallConfirmed).SelectedOptionId
}

func (_this ChatAction) Dtor_auth() m_ConfluxCodec.Json {
	return _this.Get_().(ChatAction_CToolCallAuthRequired).Auth
}

func (_this ChatAction) Dtor_success() bool {
	return _this.Get_().(ChatAction_CToolCallComplete).Success
}

func (_this ChatAction) Dtor_pastTenseMessage() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallComplete).PastTenseMessage
}

func (_this ChatAction) Dtor_resultContent() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallComplete).ResultContent
}

func (_this ChatAction) Dtor_structuredContent() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallComplete).StructuredContent
}

func (_this ChatAction) Dtor_error() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CToolCallComplete).Error
}

func (_this ChatAction) Dtor_requiresResultConfirmation() bool {
	return _this.Get_().(ChatAction_CToolCallComplete).RequiresResultConfirmation
}

func (_this ChatAction) Dtor_newContent() m_ConfluxCodec.Json {
	return _this.Get_().(ChatAction_CToolCallContentChanged).NewContent
}

func (_this ChatAction) Dtor_upToTurnId() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CTruncated).UpToTurnId
}

func (_this ChatAction) Dtor_order() _dafny.Sequence {
	return _this.Get_().(ChatAction_CQueuedReordered).Order
}

func (_this ChatAction) Dtor_loaded() _dafny.Sequence {
	return _this.Get_().(ChatAction_CTurnsLoaded).Loaded
}

func (_this ChatAction) Dtor_cursor() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CTurnsLoaded).Cursor
}

func (_this ChatAction) Dtor_req() InputReq {
	return _this.Get_().(ChatAction_CInputRequested).Req
}

func (_this ChatAction) Dtor_requestId() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case ChatAction_CInputAnswerChanged:
		return data.RequestId
	default:
		return data.(ChatAction_CInputCompleted).RequestId
	}
}

func (_this ChatAction) Dtor_questionId() _dafny.Sequence {
	return _this.Get_().(ChatAction_CInputAnswerChanged).QuestionId
}

func (_this ChatAction) Dtor_answer() m_AhpSkeleton.Option {
	return _this.Get_().(ChatAction_CInputAnswerChanged).Answer
}

func (_this ChatAction) Dtor_response() _dafny.Sequence {
	return _this.Get_().(ChatAction_CInputCompleted).Response
}

func (_this ChatAction) Dtor_answers() _dafny.Map {
	return _this.Get_().(ChatAction_CInputCompleted).Answers
}

func (_this ChatAction) Dtor_kind() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case ChatAction_CPendingMessageSet:
		return data.Kind
	default:
		return data.(ChatAction_CPendingMessageRemoved).Kind
	}
}

func (_this ChatAction) Dtor_id() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case ChatAction_CPendingMessageSet:
		return data.Id
	default:
		return data.(ChatAction_CPendingMessageRemoved).Id
	}
}

func (_this ChatAction) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(ChatAction_CUnknown).Raw
}

func (_this ChatAction) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ChatAction_CDraftChanged:
		{
			return "Chat.ChatAction.CDraftChanged" + "(" + _dafny.String(data.Draft) + ")"
		}
	case ChatAction_CActivityChanged:
		{
			return "Chat.ChatAction.CActivityChanged" + "(" + _dafny.String(data.Activity) + ")"
		}
	case ChatAction_CTurnStarted:
		{
			return "Chat.ChatAction.CTurnStarted" + "(" + data.TurnId.VerbatimString(true) + ", " + _dafny.String(data.Message) + ", " + _dafny.String(data.QueuedMessageId) + ")"
		}
	case ChatAction_CTurnComplete:
		{
			return "Chat.ChatAction.CTurnComplete" + "(" + data.TurnId.VerbatimString(true) + ")"
		}
	case ChatAction_CTurnCancelled:
		{
			return "Chat.ChatAction.CTurnCancelled" + "(" + data.TurnId.VerbatimString(true) + ")"
		}
	case ChatAction_CUsage:
		{
			return "Chat.ChatAction.CUsage" + "(" + data.TurnId.VerbatimString(true) + ", " + _dafny.String(data.Usage) + ")"
		}
	case ChatAction_CError:
		{
			return "Chat.ChatAction.CError" + "(" + data.TurnId.VerbatimString(true) + ", " + _dafny.String(data.Err) + ")"
		}
	case ChatAction_CResponsePart:
		{
			return "Chat.ChatAction.CResponsePart" + "(" + data.TurnId.VerbatimString(true) + ", " + _dafny.String(data.Part) + ")"
		}
	case ChatAction_CDelta:
		{
			return "Chat.ChatAction.CDelta" + "(" + data.TurnId.VerbatimString(true) + ", " + data.PartId.VerbatimString(true) + ", " + data.Content.VerbatimString(true) + ")"
		}
	case ChatAction_CToolCallStart:
		{
			return "Chat.ChatAction.CToolCallStart" + "(" + data.TurnId.VerbatimString(true) + ", " + data.ToolCallId.VerbatimString(true) + ", " + data.ToolName.VerbatimString(true) + ", " + data.DisplayName.VerbatimString(true) + ", " + _dafny.String(data.Intention) + ", " + _dafny.String(data.Contributor) + ", " + _dafny.String(data.Meta) + ")"
		}
	case ChatAction_CToolCallDelta:
		{
			return "Chat.ChatAction.CToolCallDelta" + "(" + data.TurnId.VerbatimString(true) + ", " + data.ToolCallId.VerbatimString(true) + ", " + data.Content.VerbatimString(true) + ", " + _dafny.String(data.InvocationMessage) + ", " + _dafny.String(data.Meta) + ")"
		}
	case ChatAction_CToolCallReady:
		{
			return "Chat.ChatAction.CToolCallReady" + "(" + data.TurnId.VerbatimString(true) + ", " + data.ToolCallId.VerbatimString(true) + ", " + _dafny.String(data.InvocationMessage) + ", " + _dafny.String(data.ToolInput) + ", " + _dafny.String(data.ConfirmationTitle) + ", " + _dafny.String(data.RiskAssessment) + ", " + _dafny.String(data.Edits) + ", " + _dafny.String(data.Editable) + ", " + _dafny.String(data.Options) + ", " + _dafny.String(data.Confirmed) + ", " + _dafny.String(data.Meta) + ")"
		}
	case ChatAction_CToolCallConfirmed:
		{
			return "Chat.ChatAction.CToolCallConfirmed" + "(" + data.TurnId.VerbatimString(true) + ", " + data.ToolCallId.VerbatimString(true) + ", " + _dafny.String(data.Approved) + ", " + _dafny.String(data.ConfirmedVal) + ", " + _dafny.String(data.Reason) + ", " + _dafny.String(data.ReasonMessage) + ", " + _dafny.String(data.UserSuggestion) + ", " + _dafny.String(data.EditedToolInput) + ", " + _dafny.String(data.SelectedOptionId) + ", " + _dafny.String(data.Meta) + ")"
		}
	case ChatAction_CToolCallAuthRequired:
		{
			return "Chat.ChatAction.CToolCallAuthRequired" + "(" + data.TurnId.VerbatimString(true) + ", " + data.ToolCallId.VerbatimString(true) + ", " + _dafny.String(data.Auth) + ", " + _dafny.String(data.Meta) + ")"
		}
	case ChatAction_CToolCallAuthResolved:
		{
			return "Chat.ChatAction.CToolCallAuthResolved" + "(" + data.TurnId.VerbatimString(true) + ", " + data.ToolCallId.VerbatimString(true) + ", " + _dafny.String(data.Meta) + ")"
		}
	case ChatAction_CToolCallComplete:
		{
			return "Chat.ChatAction.CToolCallComplete" + "(" + data.TurnId.VerbatimString(true) + ", " + data.ToolCallId.VerbatimString(true) + ", " + _dafny.String(data.Success) + ", " + _dafny.String(data.PastTenseMessage) + ", " + _dafny.String(data.ResultContent) + ", " + _dafny.String(data.StructuredContent) + ", " + _dafny.String(data.Error) + ", " + _dafny.String(data.RequiresResultConfirmation) + ", " + _dafny.String(data.Meta) + ")"
		}
	case ChatAction_CToolCallResultConfirmed:
		{
			return "Chat.ChatAction.CToolCallResultConfirmed" + "(" + data.TurnId.VerbatimString(true) + ", " + data.ToolCallId.VerbatimString(true) + ", " + _dafny.String(data.Approved) + ", " + _dafny.String(data.Meta) + ")"
		}
	case ChatAction_CToolCallContentChanged:
		{
			return "Chat.ChatAction.CToolCallContentChanged" + "(" + data.TurnId.VerbatimString(true) + ", " + data.ToolCallId.VerbatimString(true) + ", " + _dafny.String(data.NewContent) + ", " + _dafny.String(data.Meta) + ")"
		}
	case ChatAction_CReasoning:
		{
			return "Chat.ChatAction.CReasoning" + "(" + data.TurnId.VerbatimString(true) + ", " + data.PartId.VerbatimString(true) + ", " + data.Content.VerbatimString(true) + ")"
		}
	case ChatAction_CTruncated:
		{
			return "Chat.ChatAction.CTruncated" + "(" + _dafny.String(data.UpToTurnId) + ")"
		}
	case ChatAction_CQueuedReordered:
		{
			return "Chat.ChatAction.CQueuedReordered" + "(" + _dafny.String(data.Order) + ")"
		}
	case ChatAction_CTurnsLoaded:
		{
			return "Chat.ChatAction.CTurnsLoaded" + "(" + _dafny.String(data.Loaded) + ", " + _dafny.String(data.Cursor) + ")"
		}
	case ChatAction_CInputRequested:
		{
			return "Chat.ChatAction.CInputRequested" + "(" + _dafny.String(data.Req) + ")"
		}
	case ChatAction_CInputAnswerChanged:
		{
			return "Chat.ChatAction.CInputAnswerChanged" + "(" + data.RequestId.VerbatimString(true) + ", " + data.QuestionId.VerbatimString(true) + ", " + _dafny.String(data.Answer) + ")"
		}
	case ChatAction_CInputCompleted:
		{
			return "Chat.ChatAction.CInputCompleted" + "(" + data.RequestId.VerbatimString(true) + ", " + data.Response.VerbatimString(true) + ", " + _dafny.String(data.Answers) + ")"
		}
	case ChatAction_CPendingMessageSet:
		{
			return "Chat.ChatAction.CPendingMessageSet" + "(" + data.Kind.VerbatimString(true) + ", " + data.Id.VerbatimString(true) + ", " + _dafny.String(data.Message) + ")"
		}
	case ChatAction_CPendingMessageRemoved:
		{
			return "Chat.ChatAction.CPendingMessageRemoved" + "(" + data.Kind.VerbatimString(true) + ", " + data.Id.VerbatimString(true) + ")"
		}
	case ChatAction_CUnknown:
		{
			return "Chat.ChatAction.CUnknown" + "(" + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ChatAction) Equals(other ChatAction) bool {
	switch data1 := _this.Get_().(type) {
	case ChatAction_CDraftChanged:
		{
			data2, ok := other.Get_().(ChatAction_CDraftChanged)
			return ok && data1.Draft.Equals(data2.Draft)
		}
	case ChatAction_CActivityChanged:
		{
			data2, ok := other.Get_().(ChatAction_CActivityChanged)
			return ok && data1.Activity.Equals(data2.Activity)
		}
	case ChatAction_CTurnStarted:
		{
			data2, ok := other.Get_().(ChatAction_CTurnStarted)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.Message.Equals(data2.Message) && data1.QueuedMessageId.Equals(data2.QueuedMessageId)
		}
	case ChatAction_CTurnComplete:
		{
			data2, ok := other.Get_().(ChatAction_CTurnComplete)
			return ok && data1.TurnId.Equals(data2.TurnId)
		}
	case ChatAction_CTurnCancelled:
		{
			data2, ok := other.Get_().(ChatAction_CTurnCancelled)
			return ok && data1.TurnId.Equals(data2.TurnId)
		}
	case ChatAction_CUsage:
		{
			data2, ok := other.Get_().(ChatAction_CUsage)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.Usage.Equals(data2.Usage)
		}
	case ChatAction_CError:
		{
			data2, ok := other.Get_().(ChatAction_CError)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.Err.Equals(data2.Err)
		}
	case ChatAction_CResponsePart:
		{
			data2, ok := other.Get_().(ChatAction_CResponsePart)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.Part.Equals(data2.Part)
		}
	case ChatAction_CDelta:
		{
			data2, ok := other.Get_().(ChatAction_CDelta)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.PartId.Equals(data2.PartId) && data1.Content.Equals(data2.Content)
		}
	case ChatAction_CToolCallStart:
		{
			data2, ok := other.Get_().(ChatAction_CToolCallStart)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.ToolCallId.Equals(data2.ToolCallId) && data1.ToolName.Equals(data2.ToolName) && data1.DisplayName.Equals(data2.DisplayName) && data1.Intention.Equals(data2.Intention) && data1.Contributor.Equals(data2.Contributor) && data1.Meta.Equals(data2.Meta)
		}
	case ChatAction_CToolCallDelta:
		{
			data2, ok := other.Get_().(ChatAction_CToolCallDelta)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.ToolCallId.Equals(data2.ToolCallId) && data1.Content.Equals(data2.Content) && data1.InvocationMessage.Equals(data2.InvocationMessage) && data1.Meta.Equals(data2.Meta)
		}
	case ChatAction_CToolCallReady:
		{
			data2, ok := other.Get_().(ChatAction_CToolCallReady)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.ToolCallId.Equals(data2.ToolCallId) && data1.InvocationMessage.Equals(data2.InvocationMessage) && data1.ToolInput.Equals(data2.ToolInput) && data1.ConfirmationTitle.Equals(data2.ConfirmationTitle) && data1.RiskAssessment.Equals(data2.RiskAssessment) && data1.Edits.Equals(data2.Edits) && data1.Editable.Equals(data2.Editable) && data1.Options.Equals(data2.Options) && data1.Confirmed.Equals(data2.Confirmed) && data1.Meta.Equals(data2.Meta)
		}
	case ChatAction_CToolCallConfirmed:
		{
			data2, ok := other.Get_().(ChatAction_CToolCallConfirmed)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.ToolCallId.Equals(data2.ToolCallId) && data1.Approved == data2.Approved && data1.ConfirmedVal.Equals(data2.ConfirmedVal) && data1.Reason.Equals(data2.Reason) && data1.ReasonMessage.Equals(data2.ReasonMessage) && data1.UserSuggestion.Equals(data2.UserSuggestion) && data1.EditedToolInput.Equals(data2.EditedToolInput) && data1.SelectedOptionId.Equals(data2.SelectedOptionId) && data1.Meta.Equals(data2.Meta)
		}
	case ChatAction_CToolCallAuthRequired:
		{
			data2, ok := other.Get_().(ChatAction_CToolCallAuthRequired)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.ToolCallId.Equals(data2.ToolCallId) && data1.Auth.Equals(data2.Auth) && data1.Meta.Equals(data2.Meta)
		}
	case ChatAction_CToolCallAuthResolved:
		{
			data2, ok := other.Get_().(ChatAction_CToolCallAuthResolved)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.ToolCallId.Equals(data2.ToolCallId) && data1.Meta.Equals(data2.Meta)
		}
	case ChatAction_CToolCallComplete:
		{
			data2, ok := other.Get_().(ChatAction_CToolCallComplete)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.ToolCallId.Equals(data2.ToolCallId) && data1.Success == data2.Success && data1.PastTenseMessage.Equals(data2.PastTenseMessage) && data1.ResultContent.Equals(data2.ResultContent) && data1.StructuredContent.Equals(data2.StructuredContent) && data1.Error.Equals(data2.Error) && data1.RequiresResultConfirmation == data2.RequiresResultConfirmation && data1.Meta.Equals(data2.Meta)
		}
	case ChatAction_CToolCallResultConfirmed:
		{
			data2, ok := other.Get_().(ChatAction_CToolCallResultConfirmed)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.ToolCallId.Equals(data2.ToolCallId) && data1.Approved == data2.Approved && data1.Meta.Equals(data2.Meta)
		}
	case ChatAction_CToolCallContentChanged:
		{
			data2, ok := other.Get_().(ChatAction_CToolCallContentChanged)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.ToolCallId.Equals(data2.ToolCallId) && data1.NewContent.Equals(data2.NewContent) && data1.Meta.Equals(data2.Meta)
		}
	case ChatAction_CReasoning:
		{
			data2, ok := other.Get_().(ChatAction_CReasoning)
			return ok && data1.TurnId.Equals(data2.TurnId) && data1.PartId.Equals(data2.PartId) && data1.Content.Equals(data2.Content)
		}
	case ChatAction_CTruncated:
		{
			data2, ok := other.Get_().(ChatAction_CTruncated)
			return ok && data1.UpToTurnId.Equals(data2.UpToTurnId)
		}
	case ChatAction_CQueuedReordered:
		{
			data2, ok := other.Get_().(ChatAction_CQueuedReordered)
			return ok && data1.Order.Equals(data2.Order)
		}
	case ChatAction_CTurnsLoaded:
		{
			data2, ok := other.Get_().(ChatAction_CTurnsLoaded)
			return ok && data1.Loaded.Equals(data2.Loaded) && data1.Cursor.Equals(data2.Cursor)
		}
	case ChatAction_CInputRequested:
		{
			data2, ok := other.Get_().(ChatAction_CInputRequested)
			return ok && data1.Req.Equals(data2.Req)
		}
	case ChatAction_CInputAnswerChanged:
		{
			data2, ok := other.Get_().(ChatAction_CInputAnswerChanged)
			return ok && data1.RequestId.Equals(data2.RequestId) && data1.QuestionId.Equals(data2.QuestionId) && data1.Answer.Equals(data2.Answer)
		}
	case ChatAction_CInputCompleted:
		{
			data2, ok := other.Get_().(ChatAction_CInputCompleted)
			return ok && data1.RequestId.Equals(data2.RequestId) && data1.Response.Equals(data2.Response) && data1.Answers.Equals(data2.Answers)
		}
	case ChatAction_CPendingMessageSet:
		{
			data2, ok := other.Get_().(ChatAction_CPendingMessageSet)
			return ok && data1.Kind.Equals(data2.Kind) && data1.Id.Equals(data2.Id) && data1.Message.Equals(data2.Message)
		}
	case ChatAction_CPendingMessageRemoved:
		{
			data2, ok := other.Get_().(ChatAction_CPendingMessageRemoved)
			return ok && data1.Kind.Equals(data2.Kind) && data1.Id.Equals(data2.Id)
		}
	case ChatAction_CUnknown:
		{
			data2, ok := other.Get_().(ChatAction_CUnknown)
			return ok && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ChatAction) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ChatAction)
	return ok && _this.Equals(typed)
}

func Type_ChatAction_() _dafny.TypeDescriptor {
	return type_ChatAction_{}
}

type type_ChatAction_ struct {
}

func (_this type_ChatAction_) Default() interface{} {
	return Companion_ChatAction_.Default()
}

func (_this type_ChatAction_) String() string {
	return "Chat.ChatAction"
}
func (_this ChatAction) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ChatAction{}

// End of datatype ChatAction
