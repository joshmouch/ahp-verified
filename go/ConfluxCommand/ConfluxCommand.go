// Package ConfluxCommand
// Dafny module ConfluxCommand compiled into Go

package ConfluxCommand

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
	m_ConfluxDurableHistory "github.com/joshmouch/ahp-go/ConfluxDurableHistory"
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
var _ m_ConfluxDurableHistory.Dummy__

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
	return "ConfluxCommand.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) CanonicalCommandCodecVersion() _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("oa-canonical-json-v1")
}
func (_static *CompanionStruct_Default___) IsLowerHexChar(c _dafny.CodePoint) bool {
	return (((_dafny.CodePoint('0')) <= (c)) && ((c) <= (_dafny.CodePoint('9')))) || (((_dafny.CodePoint('a')) <= (c)) && ((c) <= (_dafny.CodePoint('f'))))
}
func (_static *CompanionStruct_Default___) IsLowerHex(s _dafny.Sequence) bool {
	return ((_dafny.IntOfUint32((s).Cardinality())).Sign() == 0) || ((Companion_Default___.IsLowerHexChar((s).Select(0).(_dafny.CodePoint))) && (Companion_Default___.IsLowerHex((s).Drop(1))))
}
func (_static *CompanionStruct_Default___) RepeatChar(c _dafny.CodePoint, count _dafny.Int) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (count).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.UnicodeSeqOfUtf8Bytes(""))
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf(c))
		var _in0 _dafny.CodePoint = c
		_ = _in0
		var _in1 _dafny.Int = (count).Minus(_dafny.One)
		_ = _in1
		c = _in0
		count = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) ValidSha256Digest(digest _dafny.Sequence) bool {
	return (((_dafny.IntOfUint32((digest).Cardinality())).Cmp(_dafny.IntOfInt64(71)) == 0) && (_dafny.Companion_Sequence_.Equal((digest).Take(7), _dafny.UnicodeSeqOfUtf8Bytes("sha256:")))) && (Companion_Default___.IsLowerHex((digest).Drop(7)))
}
func (_static *CompanionStruct_Default___) CanonicalCommandBytes(codec CommandIdentityCodec, proposal Proposal) _dafny.Sequence {
	return (func(coer93 func(interface{}) _dafny.Sequence) func(interface{}) _dafny.Sequence {
		return func(arg109 interface{}) _dafny.Sequence {
			return coer93(arg109)
		}
	}((codec).Dtor_encode()))((proposal).Dtor_command())
}
func (_static *CompanionStruct_Default___) ValidProposalIdentity(codec CommandIdentityCodec, proposal Proposal) bool {
	return ((((_dafny.Companion_Sequence_.Equal((codec).Dtor_codecVersion(), Companion_Default___.CanonicalCommandCodecVersion())) && (_dafny.Companion_Sequence_.Equal((proposal).Dtor_codecVersion(), (codec).Dtor_codecVersion()))) && (((func(coer94 func(_dafny.Sequence) m_ConfluxContract.Option) func(_dafny.Sequence) m_ConfluxContract.Option {
		return func(arg110 _dafny.Sequence) m_ConfluxContract.Option {
			return coer94(arg110)
		}
	}((codec).Dtor_decode()))(Companion_Default___.CanonicalCommandBytes(codec, proposal))).Equals(m_ConfluxContract.Companion_Option_.Create_Some_((proposal).Dtor_command())))) && (Companion_Default___.ValidSha256Digest((proposal).Dtor_commandDigest()))) && (_dafny.Companion_Sequence_.Equal((proposal).Dtor_commandDigest(), ((codec).Dtor_digest())(Companion_Default___.CanonicalCommandBytes(codec, proposal))))
}
func (_static *CompanionStruct_Default___) IdentityOf(codec CommandIdentityCodec, proposal Proposal) ProposalIdentity {
	var _0_bytes _dafny.Sequence = Companion_Default___.CanonicalCommandBytes(codec, proposal)
	_ = _0_bytes
	return Companion_ProposalIdentity_.Create_ProposalIdentity_(_dafny.UnicodeSeqOfUtf8Bytes("conflux-command-v1"), (codec).Dtor_codecVersion(), (proposal).Dtor_channel(), (proposal).Dtor_expectedServerSeq(), ((codec).Dtor_digest())(_0_bytes), _0_bytes)
}
func (_static *CompanionStruct_Default___) FindRecord(history _dafny.Sequence, origin CommandOrigin) m_ConfluxContract.Option {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((history).Cardinality())).Sign() == 0 {
		return m_ConfluxContract.Companion_Option_.Create_None_()
	} else if (((history).Select(0).(DecisionRecord)).Dtor_origin()).Equals(origin) {
		return m_ConfluxContract.Companion_Option_.Create_Some_((history).Select(0).(DecisionRecord))
	} else {
		var _in0 _dafny.Sequence = (history).Drop(1)
		_ = _in0
		var _in1 CommandOrigin = origin
		_ = _in1
		history = _in0
		origin = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) FindDecision(history _dafny.Sequence, origin CommandOrigin) m_ConfluxContract.Option {
	var _source0 m_ConfluxContract.Option = Companion_Default___.FindRecord(history, origin)
	_ = _source0
	{
		if _source0.Is_None() {
			return m_ConfluxContract.Companion_Option_.Create_None_()
		}
	}
	{
		var _0_record DecisionRecord = _source0.Get_().(m_ConfluxContract.Option_Some).Value.(DecisionRecord)
		_ = _0_record
		return m_ConfluxContract.Companion_Option_.Create_Some_((_0_record).Dtor_result())
	}
}
func (_static *CompanionStruct_Default___) RecordDecision(history _dafny.Sequence, channel _dafny.Sequence, origin CommandOrigin, proposalIdentity ProposalIdentity, result AuthorityResult) _dafny.Sequence {
	var _0_appended _dafny.Sequence = m_ConfluxDurableHistory.Companion_Default___.Append((history), Companion_DecisionRecord_.Create_DecisionRecord_(channel, origin, proposalIdentity, result))
	_ = _0_appended
	return (_0_appended)
}
func (_static *CompanionStruct_Default___) StampActions(channel _dafny.Sequence, origin CommandOrigin, nextSeq _dafny.Int, actions _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((actions).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf(Companion_SequencedActionEnvelope_.Create_SequencedActionEnvelope_(channel, nextSeq, origin, (actions).Select(0).(interface{}))))
		var _in0 _dafny.Sequence = channel
		_ = _in0
		var _in1 CommandOrigin = origin
		_ = _in1
		var _in2 _dafny.Int = (nextSeq).Plus(_dafny.One)
		_ = _in2
		var _in3 _dafny.Sequence = (actions).Drop(1)
		_ = _in3
		channel = _in0
		origin = _in1
		nextSeq = _in2
		actions = _in3
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) EnvelopeActions(envelopes _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((envelopes).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf(((envelopes).Select(0).(SequencedActionEnvelope)).Dtor_action()))
		var _in0 _dafny.Sequence = (envelopes).Drop(1)
		_ = _in0
		envelopes = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) DecisionActions(decision AuthorityResult) _dafny.Sequence {
	var _source0 AuthorityResult = decision
	_ = _source0
	{
		if _source0.Is_Rejected() {
			return _dafny.SeqOf()
		}
	}
	{
		var _0_envelopes _dafny.Sequence = _source0.Get_().(AuthorityResult_Accepted).SequencedActions
		_ = _0_envelopes
		return _0_envelopes
	}
}
func (_static *CompanionStruct_Default___) RecordAcceptedActions(record DecisionRecord) _dafny.Sequence {
	return Companion_Default___.DecisionActions((record).Dtor_result())
}
func (_static *CompanionStruct_Default___) AcceptedHistory(history _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((history).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, Companion_Default___.RecordAcceptedActions((history).Select(0).(DecisionRecord)))
		var _in0 _dafny.Sequence = (history).Drop(1)
		_ = _in0
		history = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) ExpectedSequenceMatches(state AuthorityState, proposal Proposal) bool {
	var _source0 m_ConfluxContract.Option = (proposal).Dtor_expectedServerSeq()
	_ = _source0
	{
		if _source0.Is_None() {
			return true
		}
	}
	{
		var _0_expected _dafny.Int = _source0.Get_().(m_ConfluxContract.Option_Some).Value.(_dafny.Int)
		_ = _0_expected
		return (_0_expected).Cmp(((state).Dtor_host()).Dtor_nextSeq()) == 0
	}
}
func (_static *CompanionStruct_Default___) StaleExpectedServerSeq() RejectionReason {
	return Companion_RejectionReason_.Create_RejectionReason_(_dafny.UnicodeSeqOfUtf8Bytes("staleExpectedServerSeq"), _dafny.UnicodeSeqOfUtf8Bytes("expected sequence does not match"), m_ConfluxContract.Companion_Option_.Create_None_())
}
func (_static *CompanionStruct_Default___) IdempotencyOriginConflict(recorded ProposalIdentity, attempted ProposalIdentity) RejectionReason {
	return Companion_RejectionReason_.Create_RejectionReason_(_dafny.UnicodeSeqOfUtf8Bytes("idempotencyOriginConflict"), _dafny.UnicodeSeqOfUtf8Bytes("command origin is already bound to a different proposal"), m_ConfluxContract.Companion_Option_.Create_Some_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((recorded).Dtor_commandDigest(), _dafny.UnicodeSeqOfUtf8Bytes(":")), (attempted).Dtor_commandDigest())))
}
func (_static *CompanionStruct_Default___) InvalidProposalIdentity(codecVersion _dafny.Sequence) RejectionReason {
	return Companion_RejectionReason_.Create_RejectionReason_(_dafny.UnicodeSeqOfUtf8Bytes("invalidProposalIdentity"), _dafny.UnicodeSeqOfUtf8Bytes("command codec version, canonical encoding round-trip, or digest is invalid"), m_ConfluxContract.Companion_Option_.Create_Some_(codecVersion))
}
func (_static *CompanionStruct_Default___) EmptyAuthorizedActions() RejectionReason {
	return Companion_RejectionReason_.Create_RejectionReason_(_dafny.UnicodeSeqOfUtf8Bytes("emptyAuthorizedActions"), _dafny.UnicodeSeqOfUtf8Bytes("mutating command authority requires at least one action"), m_ConfluxContract.Companion_Option_.Create_None_())
}
func (_static *CompanionStruct_Default___) RecordRejection(state AuthorityState, proposal Proposal, identity ProposalIdentity, reason RejectionReason) ReconcileResult {
	var _0_decision AuthorityResult = Companion_AuthorityResult_.Create_Rejected_((proposal).Dtor_origin(), reason)
	_ = _0_decision
	return Companion_ReconcileResult_.Create_ReconcileResult_(Companion_AuthorityState_.Create_AuthorityState_((state).Dtor_host(), Companion_Default___.RecordDecision((state).Dtor_history(), (proposal).Dtor_channel(), (proposal).Dtor_origin(), identity, _0_decision)), _0_decision, true)
}
func (_static *CompanionStruct_Default___) Reconcile(reducer func(interface{}, interface{}) interface{}, state AuthorityState, proposal Proposal, authorizeAndLower func(m_ConfluxChannel.HostState, interface{}) Authorization, commandCodec CommandIdentityCodec) ReconcileResult {
	if !(Companion_Default___.ValidProposalIdentity(commandCodec, proposal)) {
		return Companion_ReconcileResult_.Create_ReconcileResult_(state, Companion_AuthorityResult_.Create_Rejected_((proposal).Dtor_origin(), Companion_Default___.InvalidProposalIdentity((proposal).Dtor_codecVersion())), false)
	} else {
		var _source0 m_ConfluxContract.Option = Companion_Default___.FindRecord((state).Dtor_history(), (proposal).Dtor_origin())
		_ = _source0
		{
			if _source0.Is_Some() {
				var _0_recorded DecisionRecord = _source0.Get_().(m_ConfluxContract.Option_Some).Value.(DecisionRecord)
				_ = _0_recorded
				if ((_0_recorded).Dtor_proposalIdentity()).Equals(Companion_Default___.IdentityOf(commandCodec, proposal)) {
					return Companion_ReconcileResult_.Create_ReconcileResult_(state, (_0_recorded).Dtor_result(), false)
				} else {
					return Companion_ReconcileResult_.Create_ReconcileResult_(state, Companion_AuthorityResult_.Create_Rejected_((proposal).Dtor_origin(), Companion_Default___.IdempotencyOriginConflict((_0_recorded).Dtor_proposalIdentity(), Companion_Default___.IdentityOf(commandCodec, proposal))), false)
				}
			}
		}
		{
			if !(Companion_Default___.ExpectedSequenceMatches(state, proposal)) {
				return Companion_Default___.RecordRejection(state, proposal, Companion_Default___.IdentityOf(commandCodec, proposal), Companion_Default___.StaleExpectedServerSeq())
			} else {
				var _source1 Authorization = (authorizeAndLower)((state).Dtor_host(), (proposal).Dtor_command())
				_ = _source1
				{
					if _source1.Is_Denied() {
						var _1_reason RejectionReason = _source1.Get_().(Authorization_Denied).Reason
						_ = _1_reason
						return Companion_Default___.RecordRejection(state, proposal, Companion_Default___.IdentityOf(commandCodec, proposal), _1_reason)
					}
				}
				{
					var _2_actions _dafny.Sequence = _source1.Get_().(Authorization_Allowed).Actions
					_ = _2_actions
					if (_dafny.IntOfUint32((_2_actions).Cardinality())).Sign() == 0 {
						return Companion_Default___.RecordRejection(state, proposal, Companion_Default___.IdentityOf(commandCodec, proposal), Companion_Default___.EmptyAuthorizedActions())
					} else {
						var _3_envelopes _dafny.Sequence = Companion_Default___.StampActions((proposal).Dtor_channel(), (proposal).Dtor_origin(), ((state).Dtor_host()).Dtor_nextSeq(), _2_actions)
						_ = _3_envelopes
						var _4_decision AuthorityResult = Companion_AuthorityResult_.Create_Accepted_((proposal).Dtor_origin(), _3_envelopes)
						_ = _4_decision
						return Companion_ReconcileResult_.Create_ReconcileResult_(Companion_AuthorityState_.Create_AuthorityState_(m_ConfluxChannel.Companion_Default___.AcceptFold(func(coer95 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
							return func(arg111 interface{}, arg112 interface{}) interface{} {
								return coer95(arg111, arg112)
							}
						}(reducer), (state).Dtor_host(), _2_actions), Companion_Default___.RecordDecision((state).Dtor_history(), (proposal).Dtor_channel(), (proposal).Dtor_origin(), Companion_Default___.IdentityOf(commandCodec, proposal), _4_decision)), _4_decision, true)
					}
				}
			}
		}
	}
}
func (_static *CompanionStruct_Default___) MakeCheckpoint(historyId _dafny.Sequence, reducerSchemaVersion _dafny.Sequence, hashBytes func(_dafny.Sequence) _dafny.Sequence, codec AuthorityCodec, state AuthorityState) m_ConfluxDurableHistory.ReplayCheckpoint {
	return m_ConfluxDurableHistory.Companion_Default___.MakeCheckpoint(historyId, reducerSchemaVersion, func(coer96 func(AuthorityState) _dafny.Int) func(interface{}) _dafny.Int {
		return func(arg113 interface{}) _dafny.Int {
			return coer96(arg113.(AuthorityState))
		}
	}(func(_0_authority AuthorityState) _dafny.Int {
		return ((_0_authority).Dtor_host()).Dtor_nextSeq()
	}), hashBytes, m_ConfluxDurableHistory.Companion_StateCodec_.Create_StateCodec_(func(coer97 func(AuthorityState) _dafny.Sequence) func(interface{}) _dafny.Sequence {
		return func(arg114 interface{}) _dafny.Sequence {
			return coer97(arg114.(AuthorityState))
		}
	}(func(coer98 func(AuthorityState) _dafny.Sequence) func(AuthorityState) _dafny.Sequence {
		return func(arg115 AuthorityState) _dafny.Sequence {
			return coer98(arg115)
		}
	}((codec).Dtor_encode())), func(coer99 func(_dafny.Sequence) m_ConfluxContract.Option) func(_dafny.Sequence) m_ConfluxContract.Option {
		return func(arg116 _dafny.Sequence) m_ConfluxContract.Option {
			return coer99(arg116)
		}
	}(func(coer100 func(_dafny.Sequence) m_ConfluxContract.Option) func(_dafny.Sequence) m_ConfluxContract.Option {
		return func(arg117 _dafny.Sequence) m_ConfluxContract.Option {
			return coer100(arg117)
		}
	}((codec).Dtor_decode()))), state)
}
func (_static *CompanionStruct_Default___) RestoreCheckpoint(expectedHistoryId _dafny.Sequence, expectedReducerSchemaVersion _dafny.Sequence, hashBytes func(_dafny.Sequence) _dafny.Sequence, codec AuthorityCodec, checkpoint m_ConfluxDurableHistory.ReplayCheckpoint) m_ConfluxContract.Option {
	return m_ConfluxDurableHistory.Companion_Default___.RestoreCheckpoint(expectedHistoryId, expectedReducerSchemaVersion, func(coer101 func(AuthorityState) _dafny.Int) func(interface{}) _dafny.Int {
		return func(arg118 interface{}) _dafny.Int {
			return coer101(arg118.(AuthorityState))
		}
	}(func(_0_authority AuthorityState) _dafny.Int {
		return ((_0_authority).Dtor_host()).Dtor_nextSeq()
	}), hashBytes, m_ConfluxDurableHistory.Companion_StateCodec_.Create_StateCodec_(func(coer102 func(AuthorityState) _dafny.Sequence) func(interface{}) _dafny.Sequence {
		return func(arg119 interface{}) _dafny.Sequence {
			return coer102(arg119.(AuthorityState))
		}
	}(func(coer103 func(AuthorityState) _dafny.Sequence) func(AuthorityState) _dafny.Sequence {
		return func(arg120 AuthorityState) _dafny.Sequence {
			return coer103(arg120)
		}
	}((codec).Dtor_encode())), func(coer104 func(_dafny.Sequence) m_ConfluxContract.Option) func(_dafny.Sequence) m_ConfluxContract.Option {
		return func(arg121 _dafny.Sequence) m_ConfluxContract.Option {
			return coer104(arg121)
		}
	}(func(coer105 func(_dafny.Sequence) m_ConfluxContract.Option) func(_dafny.Sequence) m_ConfluxContract.Option {
		return func(arg122 _dafny.Sequence) m_ConfluxContract.Option {
			return coer105(arg122)
		}
	}((codec).Dtor_decode()))), (checkpoint))
}
func (_static *CompanionStruct_Default___) EnvelopesPreserveOrigin(origin CommandOrigin, envelopes _dafny.Sequence) bool {
	return ((_dafny.IntOfUint32((envelopes).Cardinality())).Sign() == 0) || (((((envelopes).Select(0).(SequencedActionEnvelope)).Dtor_origin()).Equals(origin)) && (Companion_Default___.EnvelopesPreserveOrigin(origin, (envelopes).Drop(1))))
}
func (_static *CompanionStruct_Default___) ChannelsPreserved(channel _dafny.Sequence, envelopes _dafny.Sequence) bool {
	return ((_dafny.IntOfUint32((envelopes).Cardinality())).Sign() == 0) || ((_dafny.Companion_Sequence_.Equal(((envelopes).Select(0).(SequencedActionEnvelope)).Dtor_channel(), channel)) && (Companion_Default___.ChannelsPreserved(channel, (envelopes).Drop(1))))
}
func (_static *CompanionStruct_Default___) SequencesContiguous(start _dafny.Int, envelopes _dafny.Sequence) bool {
	return ((_dafny.IntOfUint32((envelopes).Cardinality())).Sign() == 0) || (((((envelopes).Select(0).(SequencedActionEnvelope)).Dtor_serverSeq()).Cmp(start) == 0) && (Companion_Default___.SequencesContiguous((start).Plus(_dafny.One), (envelopes).Drop(1))))
}

// End of class Default__

// Definition of datatype CommandOrigin
type CommandOrigin struct {
	Data_CommandOrigin_
}

func (_this CommandOrigin) Get_() Data_CommandOrigin_ {
	return _this.Data_CommandOrigin_
}

type Data_CommandOrigin_ interface {
	isCommandOrigin()
}

type CompanionStruct_CommandOrigin_ struct {
}

var Companion_CommandOrigin_ = CompanionStruct_CommandOrigin_{}

type CommandOrigin_CommandOrigin struct {
	ClientId  _dafny.Sequence
	ClientSeq _dafny.Int
}

func (CommandOrigin_CommandOrigin) isCommandOrigin() {}

func (CompanionStruct_CommandOrigin_) Create_CommandOrigin_(ClientId _dafny.Sequence, ClientSeq _dafny.Int) CommandOrigin {
	return CommandOrigin{CommandOrigin_CommandOrigin{ClientId, ClientSeq}}
}

func (_this CommandOrigin) Is_CommandOrigin() bool {
	_, ok := _this.Get_().(CommandOrigin_CommandOrigin)
	return ok
}

func (CompanionStruct_CommandOrigin_) Default() CommandOrigin {
	return Companion_CommandOrigin_.Create_CommandOrigin_(_dafny.EmptySeq, _dafny.Zero)
}

func (_this CommandOrigin) Dtor_clientId() _dafny.Sequence {
	return _this.Get_().(CommandOrigin_CommandOrigin).ClientId
}

func (_this CommandOrigin) Dtor_clientSeq() _dafny.Int {
	return _this.Get_().(CommandOrigin_CommandOrigin).ClientSeq
}

func (_this CommandOrigin) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CommandOrigin_CommandOrigin:
		{
			return "ConfluxCommand.CommandOrigin.CommandOrigin" + "(" + data.ClientId.VerbatimString(true) + ", " + _dafny.String(data.ClientSeq) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CommandOrigin) Equals(other CommandOrigin) bool {
	switch data1 := _this.Get_().(type) {
	case CommandOrigin_CommandOrigin:
		{
			data2, ok := other.Get_().(CommandOrigin_CommandOrigin)
			return ok && data1.ClientId.Equals(data2.ClientId) && data1.ClientSeq.Cmp(data2.ClientSeq) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CommandOrigin) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CommandOrigin)
	return ok && _this.Equals(typed)
}

func Type_CommandOrigin_() _dafny.TypeDescriptor {
	return type_CommandOrigin_{}
}

type type_CommandOrigin_ struct {
}

func (_this type_CommandOrigin_) Default() interface{} {
	return Companion_CommandOrigin_.Default()
}

func (_this type_CommandOrigin_) String() string {
	return "ConfluxCommand.CommandOrigin"
}
func (_this CommandOrigin) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CommandOrigin{}

// End of datatype CommandOrigin

// Definition of datatype Proposal
type Proposal struct {
	Data_Proposal_
}

func (_this Proposal) Get_() Data_Proposal_ {
	return _this.Data_Proposal_
}

type Data_Proposal_ interface {
	isProposal()
}

type CompanionStruct_Proposal_ struct {
}

var Companion_Proposal_ = CompanionStruct_Proposal_{}

type Proposal_Proposal struct {
	Channel           _dafny.Sequence
	Origin            CommandOrigin
	ExpectedServerSeq m_ConfluxContract.Option
	CodecVersion      _dafny.Sequence
	CommandDigest     _dafny.Sequence
	Command           interface{}
}

func (Proposal_Proposal) isProposal() {}

func (CompanionStruct_Proposal_) Create_Proposal_(Channel _dafny.Sequence, Origin CommandOrigin, ExpectedServerSeq m_ConfluxContract.Option, CodecVersion _dafny.Sequence, CommandDigest _dafny.Sequence, Command interface{}) Proposal {
	return Proposal{Proposal_Proposal{Channel, Origin, ExpectedServerSeq, CodecVersion, CommandDigest, Command}}
}

func (_this Proposal) Is_Proposal() bool {
	_, ok := _this.Get_().(Proposal_Proposal)
	return ok
}

func (CompanionStruct_Proposal_) Default(_default_C interface{}) Proposal {
	return Companion_Proposal_.Create_Proposal_(_dafny.EmptySeq, Companion_CommandOrigin_.Default(), m_ConfluxContract.Companion_Option_.Default(), _dafny.EmptySeq, _dafny.EmptySeq, _default_C)
}

func (_this Proposal) Dtor_channel() _dafny.Sequence {
	return _this.Get_().(Proposal_Proposal).Channel
}

func (_this Proposal) Dtor_origin() CommandOrigin {
	return _this.Get_().(Proposal_Proposal).Origin
}

func (_this Proposal) Dtor_expectedServerSeq() m_ConfluxContract.Option {
	return _this.Get_().(Proposal_Proposal).ExpectedServerSeq
}

func (_this Proposal) Dtor_codecVersion() _dafny.Sequence {
	return _this.Get_().(Proposal_Proposal).CodecVersion
}

func (_this Proposal) Dtor_commandDigest() _dafny.Sequence {
	return _this.Get_().(Proposal_Proposal).CommandDigest
}

func (_this Proposal) Dtor_command() interface{} {
	return _this.Get_().(Proposal_Proposal).Command
}

func (_this Proposal) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Proposal_Proposal:
		{
			return "ConfluxCommand.Proposal.Proposal" + "(" + data.Channel.VerbatimString(true) + ", " + _dafny.String(data.Origin) + ", " + _dafny.String(data.ExpectedServerSeq) + ", " + data.CodecVersion.VerbatimString(true) + ", " + data.CommandDigest.VerbatimString(true) + ", " + _dafny.String(data.Command) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Proposal) Equals(other Proposal) bool {
	switch data1 := _this.Get_().(type) {
	case Proposal_Proposal:
		{
			data2, ok := other.Get_().(Proposal_Proposal)
			return ok && data1.Channel.Equals(data2.Channel) && data1.Origin.Equals(data2.Origin) && data1.ExpectedServerSeq.Equals(data2.ExpectedServerSeq) && data1.CodecVersion.Equals(data2.CodecVersion) && data1.CommandDigest.Equals(data2.CommandDigest) && _dafny.AreEqual(data1.Command, data2.Command)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Proposal) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Proposal)
	return ok && _this.Equals(typed)
}

func Type_Proposal_(Type_C_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_Proposal_{Type_C_}
}

type type_Proposal_ struct {
	Type_C_ _dafny.TypeDescriptor
}

func (_this type_Proposal_) Default() interface{} {
	Type_C_ := _this.Type_C_
	_ = Type_C_
	return Companion_Proposal_.Default(Type_C_.Default())
}

func (_this type_Proposal_) String() string {
	return "ConfluxCommand.Proposal"
}
func (_this Proposal) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Proposal{}

// End of datatype Proposal

// Definition of datatype ProposalIdentity
type ProposalIdentity struct {
	Data_ProposalIdentity_
}

func (_this ProposalIdentity) Get_() Data_ProposalIdentity_ {
	return _this.Data_ProposalIdentity_
}

type Data_ProposalIdentity_ interface {
	isProposalIdentity()
}

type CompanionStruct_ProposalIdentity_ struct {
}

var Companion_ProposalIdentity_ = CompanionStruct_ProposalIdentity_{}

type ProposalIdentity_ProposalIdentity struct {
	Protocol              _dafny.Sequence
	CodecVersion          _dafny.Sequence
	Channel               _dafny.Sequence
	ExpectedServerSeq     m_ConfluxContract.Option
	CommandDigest         _dafny.Sequence
	CanonicalCommandBytes _dafny.Sequence
}

func (ProposalIdentity_ProposalIdentity) isProposalIdentity() {}

func (CompanionStruct_ProposalIdentity_) Create_ProposalIdentity_(Protocol _dafny.Sequence, CodecVersion _dafny.Sequence, Channel _dafny.Sequence, ExpectedServerSeq m_ConfluxContract.Option, CommandDigest _dafny.Sequence, CanonicalCommandBytes _dafny.Sequence) ProposalIdentity {
	return ProposalIdentity{ProposalIdentity_ProposalIdentity{Protocol, CodecVersion, Channel, ExpectedServerSeq, CommandDigest, CanonicalCommandBytes}}
}

func (_this ProposalIdentity) Is_ProposalIdentity() bool {
	_, ok := _this.Get_().(ProposalIdentity_ProposalIdentity)
	return ok
}

func (CompanionStruct_ProposalIdentity_) Default() ProposalIdentity {
	return Companion_ProposalIdentity_.Create_ProposalIdentity_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, m_ConfluxContract.Companion_Option_.Default(), _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this ProposalIdentity) Dtor_protocol() _dafny.Sequence {
	return _this.Get_().(ProposalIdentity_ProposalIdentity).Protocol
}

func (_this ProposalIdentity) Dtor_codecVersion() _dafny.Sequence {
	return _this.Get_().(ProposalIdentity_ProposalIdentity).CodecVersion
}

func (_this ProposalIdentity) Dtor_channel() _dafny.Sequence {
	return _this.Get_().(ProposalIdentity_ProposalIdentity).Channel
}

func (_this ProposalIdentity) Dtor_expectedServerSeq() m_ConfluxContract.Option {
	return _this.Get_().(ProposalIdentity_ProposalIdentity).ExpectedServerSeq
}

func (_this ProposalIdentity) Dtor_commandDigest() _dafny.Sequence {
	return _this.Get_().(ProposalIdentity_ProposalIdentity).CommandDigest
}

func (_this ProposalIdentity) Dtor_canonicalCommandBytes() _dafny.Sequence {
	return _this.Get_().(ProposalIdentity_ProposalIdentity).CanonicalCommandBytes
}

func (_this ProposalIdentity) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ProposalIdentity_ProposalIdentity:
		{
			return "ConfluxCommand.ProposalIdentity.ProposalIdentity" + "(" + data.Protocol.VerbatimString(true) + ", " + data.CodecVersion.VerbatimString(true) + ", " + data.Channel.VerbatimString(true) + ", " + _dafny.String(data.ExpectedServerSeq) + ", " + data.CommandDigest.VerbatimString(true) + ", " + _dafny.String(data.CanonicalCommandBytes) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ProposalIdentity) Equals(other ProposalIdentity) bool {
	switch data1 := _this.Get_().(type) {
	case ProposalIdentity_ProposalIdentity:
		{
			data2, ok := other.Get_().(ProposalIdentity_ProposalIdentity)
			return ok && data1.Protocol.Equals(data2.Protocol) && data1.CodecVersion.Equals(data2.CodecVersion) && data1.Channel.Equals(data2.Channel) && data1.ExpectedServerSeq.Equals(data2.ExpectedServerSeq) && data1.CommandDigest.Equals(data2.CommandDigest) && data1.CanonicalCommandBytes.Equals(data2.CanonicalCommandBytes)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ProposalIdentity) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ProposalIdentity)
	return ok && _this.Equals(typed)
}

func Type_ProposalIdentity_() _dafny.TypeDescriptor {
	return type_ProposalIdentity_{}
}

type type_ProposalIdentity_ struct {
}

func (_this type_ProposalIdentity_) Default() interface{} {
	return Companion_ProposalIdentity_.Default()
}

func (_this type_ProposalIdentity_) String() string {
	return "ConfluxCommand.ProposalIdentity"
}
func (_this ProposalIdentity) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ProposalIdentity{}

// End of datatype ProposalIdentity

// Definition of datatype CommandIdentityCodec
type CommandIdentityCodec struct {
	Data_CommandIdentityCodec_
}

func (_this CommandIdentityCodec) Get_() Data_CommandIdentityCodec_ {
	return _this.Data_CommandIdentityCodec_
}

type Data_CommandIdentityCodec_ interface {
	isCommandIdentityCodec()
}

type CompanionStruct_CommandIdentityCodec_ struct {
}

var Companion_CommandIdentityCodec_ = CompanionStruct_CommandIdentityCodec_{}

type CommandIdentityCodec_CommandIdentityCodec struct {
	CodecVersion _dafny.Sequence
	Encode       func(interface{}) _dafny.Sequence
	Decode       func(_dafny.Sequence) m_ConfluxContract.Option
	Digest       func(_dafny.Sequence) _dafny.Sequence
}

func (CommandIdentityCodec_CommandIdentityCodec) isCommandIdentityCodec() {}

func (CompanionStruct_CommandIdentityCodec_) Create_CommandIdentityCodec_(CodecVersion _dafny.Sequence, Encode func(interface{}) _dafny.Sequence, Decode func(_dafny.Sequence) m_ConfluxContract.Option, Digest func(_dafny.Sequence) _dafny.Sequence) CommandIdentityCodec {
	return CommandIdentityCodec{CommandIdentityCodec_CommandIdentityCodec{CodecVersion, Encode, Decode, Digest}}
}

func (_this CommandIdentityCodec) Is_CommandIdentityCodec() bool {
	_, ok := _this.Get_().(CommandIdentityCodec_CommandIdentityCodec)
	return ok
}

func (CompanionStruct_CommandIdentityCodec_) Default() CommandIdentityCodec {
	return Companion_CommandIdentityCodec_.Create_CommandIdentityCodec_(_dafny.EmptySeq, func(interface{}) _dafny.Sequence { return _dafny.EmptySeq }, func(_dafny.Sequence) m_ConfluxContract.Option { return m_ConfluxContract.Companion_Option_.Default() }, func(_dafny.Sequence) _dafny.Sequence { return _dafny.EmptySeq })
}

func (_this CommandIdentityCodec) Dtor_codecVersion() _dafny.Sequence {
	return _this.Get_().(CommandIdentityCodec_CommandIdentityCodec).CodecVersion
}

func (_this CommandIdentityCodec) Dtor_encode() func(interface{}) _dafny.Sequence {
	return _this.Get_().(CommandIdentityCodec_CommandIdentityCodec).Encode
}

func (_this CommandIdentityCodec) Dtor_decode() func(_dafny.Sequence) m_ConfluxContract.Option {
	return _this.Get_().(CommandIdentityCodec_CommandIdentityCodec).Decode
}

func (_this CommandIdentityCodec) Dtor_digest() func(_dafny.Sequence) _dafny.Sequence {
	return _this.Get_().(CommandIdentityCodec_CommandIdentityCodec).Digest
}

func (_this CommandIdentityCodec) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CommandIdentityCodec_CommandIdentityCodec:
		{
			return "ConfluxCommand.CommandIdentityCodec.CommandIdentityCodec" + "(" + data.CodecVersion.VerbatimString(true) + ", " + _dafny.String(data.Encode) + ", " + _dafny.String(data.Decode) + ", " + _dafny.String(data.Digest) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CommandIdentityCodec) Equals(other CommandIdentityCodec) bool {
	switch data1 := _this.Get_().(type) {
	case CommandIdentityCodec_CommandIdentityCodec:
		{
			data2, ok := other.Get_().(CommandIdentityCodec_CommandIdentityCodec)
			return ok && data1.CodecVersion.Equals(data2.CodecVersion) && _dafny.AreEqual(data1.Encode, data2.Encode) && _dafny.AreEqual(data1.Decode, data2.Decode) && _dafny.AreEqual(data1.Digest, data2.Digest)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CommandIdentityCodec) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CommandIdentityCodec)
	return ok && _this.Equals(typed)
}

func Type_CommandIdentityCodec_() _dafny.TypeDescriptor {
	return type_CommandIdentityCodec_{}
}

type type_CommandIdentityCodec_ struct {
}

func (_this type_CommandIdentityCodec_) Default() interface{} {
	return Companion_CommandIdentityCodec_.Default()
}

func (_this type_CommandIdentityCodec_) String() string {
	return "ConfluxCommand.CommandIdentityCodec"
}
func (_this CommandIdentityCodec) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CommandIdentityCodec{}

// End of datatype CommandIdentityCodec

// Definition of datatype RejectionReason
type RejectionReason struct {
	Data_RejectionReason_
}

func (_this RejectionReason) Get_() Data_RejectionReason_ {
	return _this.Data_RejectionReason_
}

type Data_RejectionReason_ interface {
	isRejectionReason()
}

type CompanionStruct_RejectionReason_ struct {
}

var Companion_RejectionReason_ = CompanionStruct_RejectionReason_{}

type RejectionReason_RejectionReason struct {
	Code          _dafny.Sequence
	Message       _dafny.Sequence
	DetailsDigest m_ConfluxContract.Option
}

func (RejectionReason_RejectionReason) isRejectionReason() {}

func (CompanionStruct_RejectionReason_) Create_RejectionReason_(Code _dafny.Sequence, Message _dafny.Sequence, DetailsDigest m_ConfluxContract.Option) RejectionReason {
	return RejectionReason{RejectionReason_RejectionReason{Code, Message, DetailsDigest}}
}

func (_this RejectionReason) Is_RejectionReason() bool {
	_, ok := _this.Get_().(RejectionReason_RejectionReason)
	return ok
}

func (CompanionStruct_RejectionReason_) Default() RejectionReason {
	return Companion_RejectionReason_.Create_RejectionReason_(_dafny.EmptySeq, _dafny.EmptySeq, m_ConfluxContract.Companion_Option_.Default())
}

func (_this RejectionReason) Dtor_code() _dafny.Sequence {
	return _this.Get_().(RejectionReason_RejectionReason).Code
}

func (_this RejectionReason) Dtor_message() _dafny.Sequence {
	return _this.Get_().(RejectionReason_RejectionReason).Message
}

func (_this RejectionReason) Dtor_detailsDigest() m_ConfluxContract.Option {
	return _this.Get_().(RejectionReason_RejectionReason).DetailsDigest
}

func (_this RejectionReason) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RejectionReason_RejectionReason:
		{
			return "ConfluxCommand.RejectionReason.RejectionReason" + "(" + data.Code.VerbatimString(true) + ", " + data.Message.VerbatimString(true) + ", " + _dafny.String(data.DetailsDigest) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RejectionReason) Equals(other RejectionReason) bool {
	switch data1 := _this.Get_().(type) {
	case RejectionReason_RejectionReason:
		{
			data2, ok := other.Get_().(RejectionReason_RejectionReason)
			return ok && data1.Code.Equals(data2.Code) && data1.Message.Equals(data2.Message) && data1.DetailsDigest.Equals(data2.DetailsDigest)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RejectionReason) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RejectionReason)
	return ok && _this.Equals(typed)
}

func Type_RejectionReason_() _dafny.TypeDescriptor {
	return type_RejectionReason_{}
}

type type_RejectionReason_ struct {
}

func (_this type_RejectionReason_) Default() interface{} {
	return Companion_RejectionReason_.Default()
}

func (_this type_RejectionReason_) String() string {
	return "ConfluxCommand.RejectionReason"
}
func (_this RejectionReason) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RejectionReason{}

// End of datatype RejectionReason

// Definition of datatype SequencedActionEnvelope
type SequencedActionEnvelope struct {
	Data_SequencedActionEnvelope_
}

func (_this SequencedActionEnvelope) Get_() Data_SequencedActionEnvelope_ {
	return _this.Data_SequencedActionEnvelope_
}

type Data_SequencedActionEnvelope_ interface {
	isSequencedActionEnvelope()
}

type CompanionStruct_SequencedActionEnvelope_ struct {
}

var Companion_SequencedActionEnvelope_ = CompanionStruct_SequencedActionEnvelope_{}

type SequencedActionEnvelope_SequencedActionEnvelope struct {
	Channel   _dafny.Sequence
	ServerSeq _dafny.Int
	Origin    CommandOrigin
	Action    interface{}
}

func (SequencedActionEnvelope_SequencedActionEnvelope) isSequencedActionEnvelope() {}

func (CompanionStruct_SequencedActionEnvelope_) Create_SequencedActionEnvelope_(Channel _dafny.Sequence, ServerSeq _dafny.Int, Origin CommandOrigin, Action interface{}) SequencedActionEnvelope {
	return SequencedActionEnvelope{SequencedActionEnvelope_SequencedActionEnvelope{Channel, ServerSeq, Origin, Action}}
}

func (_this SequencedActionEnvelope) Is_SequencedActionEnvelope() bool {
	_, ok := _this.Get_().(SequencedActionEnvelope_SequencedActionEnvelope)
	return ok
}

func (CompanionStruct_SequencedActionEnvelope_) Default(_default_A interface{}) SequencedActionEnvelope {
	return Companion_SequencedActionEnvelope_.Create_SequencedActionEnvelope_(_dafny.EmptySeq, _dafny.Zero, Companion_CommandOrigin_.Default(), _default_A)
}

func (_this SequencedActionEnvelope) Dtor_channel() _dafny.Sequence {
	return _this.Get_().(SequencedActionEnvelope_SequencedActionEnvelope).Channel
}

func (_this SequencedActionEnvelope) Dtor_serverSeq() _dafny.Int {
	return _this.Get_().(SequencedActionEnvelope_SequencedActionEnvelope).ServerSeq
}

func (_this SequencedActionEnvelope) Dtor_origin() CommandOrigin {
	return _this.Get_().(SequencedActionEnvelope_SequencedActionEnvelope).Origin
}

func (_this SequencedActionEnvelope) Dtor_action() interface{} {
	return _this.Get_().(SequencedActionEnvelope_SequencedActionEnvelope).Action
}

func (_this SequencedActionEnvelope) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SequencedActionEnvelope_SequencedActionEnvelope:
		{
			return "ConfluxCommand.SequencedActionEnvelope.SequencedActionEnvelope" + "(" + data.Channel.VerbatimString(true) + ", " + _dafny.String(data.ServerSeq) + ", " + _dafny.String(data.Origin) + ", " + _dafny.String(data.Action) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SequencedActionEnvelope) Equals(other SequencedActionEnvelope) bool {
	switch data1 := _this.Get_().(type) {
	case SequencedActionEnvelope_SequencedActionEnvelope:
		{
			data2, ok := other.Get_().(SequencedActionEnvelope_SequencedActionEnvelope)
			return ok && data1.Channel.Equals(data2.Channel) && data1.ServerSeq.Cmp(data2.ServerSeq) == 0 && data1.Origin.Equals(data2.Origin) && _dafny.AreEqual(data1.Action, data2.Action)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SequencedActionEnvelope) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SequencedActionEnvelope)
	return ok && _this.Equals(typed)
}

func Type_SequencedActionEnvelope_(Type_A_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_SequencedActionEnvelope_{Type_A_}
}

type type_SequencedActionEnvelope_ struct {
	Type_A_ _dafny.TypeDescriptor
}

func (_this type_SequencedActionEnvelope_) Default() interface{} {
	Type_A_ := _this.Type_A_
	_ = Type_A_
	return Companion_SequencedActionEnvelope_.Default(Type_A_.Default())
}

func (_this type_SequencedActionEnvelope_) String() string {
	return "ConfluxCommand.SequencedActionEnvelope"
}
func (_this SequencedActionEnvelope) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SequencedActionEnvelope{}

// End of datatype SequencedActionEnvelope

// Definition of datatype AuthorityResult
type AuthorityResult struct {
	Data_AuthorityResult_
}

func (_this AuthorityResult) Get_() Data_AuthorityResult_ {
	return _this.Data_AuthorityResult_
}

type Data_AuthorityResult_ interface {
	isAuthorityResult()
}

type CompanionStruct_AuthorityResult_ struct {
}

var Companion_AuthorityResult_ = CompanionStruct_AuthorityResult_{}

type AuthorityResult_Rejected struct {
	Origin CommandOrigin
	Reason RejectionReason
}

func (AuthorityResult_Rejected) isAuthorityResult() {}

func (CompanionStruct_AuthorityResult_) Create_Rejected_(Origin CommandOrigin, Reason RejectionReason) AuthorityResult {
	return AuthorityResult{AuthorityResult_Rejected{Origin, Reason}}
}

func (_this AuthorityResult) Is_Rejected() bool {
	_, ok := _this.Get_().(AuthorityResult_Rejected)
	return ok
}

type AuthorityResult_Accepted struct {
	Origin           CommandOrigin
	SequencedActions _dafny.Sequence
}

func (AuthorityResult_Accepted) isAuthorityResult() {}

func (CompanionStruct_AuthorityResult_) Create_Accepted_(Origin CommandOrigin, SequencedActions _dafny.Sequence) AuthorityResult {
	return AuthorityResult{AuthorityResult_Accepted{Origin, SequencedActions}}
}

func (_this AuthorityResult) Is_Accepted() bool {
	_, ok := _this.Get_().(AuthorityResult_Accepted)
	return ok
}

func (CompanionStruct_AuthorityResult_) Default() AuthorityResult {
	return Companion_AuthorityResult_.Create_Rejected_(Companion_CommandOrigin_.Default(), Companion_RejectionReason_.Default())
}

func (_this AuthorityResult) Dtor_origin() CommandOrigin {
	switch data := _this.Get_().(type) {
	case AuthorityResult_Rejected:
		return data.Origin
	default:
		return data.(AuthorityResult_Accepted).Origin
	}
}

func (_this AuthorityResult) Dtor_reason() RejectionReason {
	return _this.Get_().(AuthorityResult_Rejected).Reason
}

func (_this AuthorityResult) Dtor_sequencedActions() _dafny.Sequence {
	return _this.Get_().(AuthorityResult_Accepted).SequencedActions
}

func (_this AuthorityResult) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case AuthorityResult_Rejected:
		{
			return "ConfluxCommand.AuthorityResult.Rejected" + "(" + _dafny.String(data.Origin) + ", " + _dafny.String(data.Reason) + ")"
		}
	case AuthorityResult_Accepted:
		{
			return "ConfluxCommand.AuthorityResult.Accepted" + "(" + _dafny.String(data.Origin) + ", " + _dafny.String(data.SequencedActions) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this AuthorityResult) Equals(other AuthorityResult) bool {
	switch data1 := _this.Get_().(type) {
	case AuthorityResult_Rejected:
		{
			data2, ok := other.Get_().(AuthorityResult_Rejected)
			return ok && data1.Origin.Equals(data2.Origin) && data1.Reason.Equals(data2.Reason)
		}
	case AuthorityResult_Accepted:
		{
			data2, ok := other.Get_().(AuthorityResult_Accepted)
			return ok && data1.Origin.Equals(data2.Origin) && data1.SequencedActions.Equals(data2.SequencedActions)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this AuthorityResult) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(AuthorityResult)
	return ok && _this.Equals(typed)
}

func Type_AuthorityResult_() _dafny.TypeDescriptor {
	return type_AuthorityResult_{}
}

type type_AuthorityResult_ struct {
}

func (_this type_AuthorityResult_) Default() interface{} {
	return Companion_AuthorityResult_.Default()
}

func (_this type_AuthorityResult_) String() string {
	return "ConfluxCommand.AuthorityResult"
}
func (_this AuthorityResult) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = AuthorityResult{}

// End of datatype AuthorityResult

// Definition of datatype Authorization
type Authorization struct {
	Data_Authorization_
}

func (_this Authorization) Get_() Data_Authorization_ {
	return _this.Data_Authorization_
}

type Data_Authorization_ interface {
	isAuthorization()
}

type CompanionStruct_Authorization_ struct {
}

var Companion_Authorization_ = CompanionStruct_Authorization_{}

type Authorization_Denied struct {
	Reason RejectionReason
}

func (Authorization_Denied) isAuthorization() {}

func (CompanionStruct_Authorization_) Create_Denied_(Reason RejectionReason) Authorization {
	return Authorization{Authorization_Denied{Reason}}
}

func (_this Authorization) Is_Denied() bool {
	_, ok := _this.Get_().(Authorization_Denied)
	return ok
}

type Authorization_Allowed struct {
	Actions _dafny.Sequence
}

func (Authorization_Allowed) isAuthorization() {}

func (CompanionStruct_Authorization_) Create_Allowed_(Actions _dafny.Sequence) Authorization {
	return Authorization{Authorization_Allowed{Actions}}
}

func (_this Authorization) Is_Allowed() bool {
	_, ok := _this.Get_().(Authorization_Allowed)
	return ok
}

func (CompanionStruct_Authorization_) Default() Authorization {
	return Companion_Authorization_.Create_Denied_(Companion_RejectionReason_.Default())
}

func (_this Authorization) Dtor_reason() RejectionReason {
	return _this.Get_().(Authorization_Denied).Reason
}

func (_this Authorization) Dtor_actions() _dafny.Sequence {
	return _this.Get_().(Authorization_Allowed).Actions
}

func (_this Authorization) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Authorization_Denied:
		{
			return "ConfluxCommand.Authorization.Denied" + "(" + _dafny.String(data.Reason) + ")"
		}
	case Authorization_Allowed:
		{
			return "ConfluxCommand.Authorization.Allowed" + "(" + _dafny.String(data.Actions) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Authorization) Equals(other Authorization) bool {
	switch data1 := _this.Get_().(type) {
	case Authorization_Denied:
		{
			data2, ok := other.Get_().(Authorization_Denied)
			return ok && data1.Reason.Equals(data2.Reason)
		}
	case Authorization_Allowed:
		{
			data2, ok := other.Get_().(Authorization_Allowed)
			return ok && data1.Actions.Equals(data2.Actions)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Authorization) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Authorization)
	return ok && _this.Equals(typed)
}

func Type_Authorization_() _dafny.TypeDescriptor {
	return type_Authorization_{}
}

type type_Authorization_ struct {
}

func (_this type_Authorization_) Default() interface{} {
	return Companion_Authorization_.Default()
}

func (_this type_Authorization_) String() string {
	return "ConfluxCommand.Authorization"
}
func (_this Authorization) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Authorization{}

// End of datatype Authorization

// Definition of datatype DecisionRecord
type DecisionRecord struct {
	Data_DecisionRecord_
}

func (_this DecisionRecord) Get_() Data_DecisionRecord_ {
	return _this.Data_DecisionRecord_
}

type Data_DecisionRecord_ interface {
	isDecisionRecord()
}

type CompanionStruct_DecisionRecord_ struct {
}

var Companion_DecisionRecord_ = CompanionStruct_DecisionRecord_{}

type DecisionRecord_DecisionRecord struct {
	Channel          _dafny.Sequence
	Origin           CommandOrigin
	ProposalIdentity ProposalIdentity
	Result           AuthorityResult
}

func (DecisionRecord_DecisionRecord) isDecisionRecord() {}

func (CompanionStruct_DecisionRecord_) Create_DecisionRecord_(Channel _dafny.Sequence, Origin CommandOrigin, ProposalIdentity ProposalIdentity, Result AuthorityResult) DecisionRecord {
	return DecisionRecord{DecisionRecord_DecisionRecord{Channel, Origin, ProposalIdentity, Result}}
}

func (_this DecisionRecord) Is_DecisionRecord() bool {
	_, ok := _this.Get_().(DecisionRecord_DecisionRecord)
	return ok
}

func (CompanionStruct_DecisionRecord_) Default() DecisionRecord {
	return Companion_DecisionRecord_.Create_DecisionRecord_(_dafny.EmptySeq, Companion_CommandOrigin_.Default(), Companion_ProposalIdentity_.Default(), Companion_AuthorityResult_.Default())
}

func (_this DecisionRecord) Dtor_channel() _dafny.Sequence {
	return _this.Get_().(DecisionRecord_DecisionRecord).Channel
}

func (_this DecisionRecord) Dtor_origin() CommandOrigin {
	return _this.Get_().(DecisionRecord_DecisionRecord).Origin
}

func (_this DecisionRecord) Dtor_proposalIdentity() ProposalIdentity {
	return _this.Get_().(DecisionRecord_DecisionRecord).ProposalIdentity
}

func (_this DecisionRecord) Dtor_result() AuthorityResult {
	return _this.Get_().(DecisionRecord_DecisionRecord).Result
}

func (_this DecisionRecord) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case DecisionRecord_DecisionRecord:
		{
			return "ConfluxCommand.DecisionRecord.DecisionRecord" + "(" + data.Channel.VerbatimString(true) + ", " + _dafny.String(data.Origin) + ", " + _dafny.String(data.ProposalIdentity) + ", " + _dafny.String(data.Result) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this DecisionRecord) Equals(other DecisionRecord) bool {
	switch data1 := _this.Get_().(type) {
	case DecisionRecord_DecisionRecord:
		{
			data2, ok := other.Get_().(DecisionRecord_DecisionRecord)
			return ok && data1.Channel.Equals(data2.Channel) && data1.Origin.Equals(data2.Origin) && data1.ProposalIdentity.Equals(data2.ProposalIdentity) && data1.Result.Equals(data2.Result)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this DecisionRecord) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(DecisionRecord)
	return ok && _this.Equals(typed)
}

func Type_DecisionRecord_() _dafny.TypeDescriptor {
	return type_DecisionRecord_{}
}

type type_DecisionRecord_ struct {
}

func (_this type_DecisionRecord_) Default() interface{} {
	return Companion_DecisionRecord_.Default()
}

func (_this type_DecisionRecord_) String() string {
	return "ConfluxCommand.DecisionRecord"
}
func (_this DecisionRecord) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = DecisionRecord{}

// End of datatype DecisionRecord

// Definition of datatype DecisionHistory
type DecisionHistory struct {
	Data_DecisionHistory_
}

func (_this DecisionHistory) Get_() Data_DecisionHistory_ {
	return _this.Data_DecisionHistory_
}

type Data_DecisionHistory_ interface {
	isDecisionHistory()
}

type CompanionStruct_DecisionHistory_ struct {
}

var Companion_DecisionHistory_ = CompanionStruct_DecisionHistory_{}

type DecisionHistory_DecisionHistory struct {
	Records _dafny.Sequence
}

func (DecisionHistory_DecisionHistory) isDecisionHistory() {}

func (CompanionStruct_DecisionHistory_) Create_DecisionHistory_(Records _dafny.Sequence) DecisionHistory {
	return DecisionHistory{DecisionHistory_DecisionHistory{Records}}
}

func (_this DecisionHistory) Is_DecisionHistory() bool {
	_, ok := _this.Get_().(DecisionHistory_DecisionHistory)
	return ok
}

func (CompanionStruct_DecisionHistory_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this DecisionHistory) Dtor_records() _dafny.Sequence {
	return _this.Get_().(DecisionHistory_DecisionHistory).Records
}

func (_this DecisionHistory) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case DecisionHistory_DecisionHistory:
		{
			return "ConfluxCommand.DecisionHistory.DecisionHistory" + "(" + _dafny.String(data.Records) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this DecisionHistory) Equals(other DecisionHistory) bool {
	switch data1 := _this.Get_().(type) {
	case DecisionHistory_DecisionHistory:
		{
			data2, ok := other.Get_().(DecisionHistory_DecisionHistory)
			return ok && data1.Records.Equals(data2.Records)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this DecisionHistory) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(DecisionHistory)
	return ok && _this.Equals(typed)
}

func Type_DecisionHistory_() _dafny.TypeDescriptor {
	return type_DecisionHistory_{}
}

type type_DecisionHistory_ struct {
}

func (_this type_DecisionHistory_) Default() interface{} {
	return Companion_DecisionHistory_.Default()
}

func (_this type_DecisionHistory_) String() string {
	return "ConfluxCommand.DecisionHistory"
}
func (_this DecisionHistory) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = DecisionHistory{}

// End of datatype DecisionHistory

// Definition of datatype AuthorityState
type AuthorityState struct {
	Data_AuthorityState_
}

func (_this AuthorityState) Get_() Data_AuthorityState_ {
	return _this.Data_AuthorityState_
}

type Data_AuthorityState_ interface {
	isAuthorityState()
}

type CompanionStruct_AuthorityState_ struct {
}

var Companion_AuthorityState_ = CompanionStruct_AuthorityState_{}

type AuthorityState_AuthorityState struct {
	Host    m_ConfluxChannel.HostState
	History _dafny.Sequence
}

func (AuthorityState_AuthorityState) isAuthorityState() {}

func (CompanionStruct_AuthorityState_) Create_AuthorityState_(Host m_ConfluxChannel.HostState, History _dafny.Sequence) AuthorityState {
	return AuthorityState{AuthorityState_AuthorityState{Host, History}}
}

func (_this AuthorityState) Is_AuthorityState() bool {
	_, ok := _this.Get_().(AuthorityState_AuthorityState)
	return ok
}

func (CompanionStruct_AuthorityState_) Default(_default_S interface{}) AuthorityState {
	return Companion_AuthorityState_.Create_AuthorityState_(m_ConfluxChannel.Companion_HostState_.Default(_default_S), _dafny.EmptySeq)
}

func (_this AuthorityState) Dtor_host() m_ConfluxChannel.HostState {
	return _this.Get_().(AuthorityState_AuthorityState).Host
}

func (_this AuthorityState) Dtor_history() _dafny.Sequence {
	return _this.Get_().(AuthorityState_AuthorityState).History
}

func (_this AuthorityState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case AuthorityState_AuthorityState:
		{
			return "ConfluxCommand.AuthorityState.AuthorityState" + "(" + _dafny.String(data.Host) + ", " + _dafny.String(data.History) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this AuthorityState) Equals(other AuthorityState) bool {
	switch data1 := _this.Get_().(type) {
	case AuthorityState_AuthorityState:
		{
			data2, ok := other.Get_().(AuthorityState_AuthorityState)
			return ok && data1.Host.Equals(data2.Host) && data1.History.Equals(data2.History)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this AuthorityState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(AuthorityState)
	return ok && _this.Equals(typed)
}

func Type_AuthorityState_(Type_S_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_AuthorityState_{Type_S_}
}

type type_AuthorityState_ struct {
	Type_S_ _dafny.TypeDescriptor
}

func (_this type_AuthorityState_) Default() interface{} {
	Type_S_ := _this.Type_S_
	_ = Type_S_
	return Companion_AuthorityState_.Default(Type_S_.Default())
}

func (_this type_AuthorityState_) String() string {
	return "ConfluxCommand.AuthorityState"
}
func (_this AuthorityState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = AuthorityState{}

// End of datatype AuthorityState

// Definition of datatype ReconcileResult
type ReconcileResult struct {
	Data_ReconcileResult_
}

func (_this ReconcileResult) Get_() Data_ReconcileResult_ {
	return _this.Data_ReconcileResult_
}

type Data_ReconcileResult_ interface {
	isReconcileResult()
}

type CompanionStruct_ReconcileResult_ struct {
}

var Companion_ReconcileResult_ = CompanionStruct_ReconcileResult_{}

type ReconcileResult_ReconcileResult struct {
	Authority AuthorityState
	Decision  AuthorityResult
	IsFresh   bool
}

func (ReconcileResult_ReconcileResult) isReconcileResult() {}

func (CompanionStruct_ReconcileResult_) Create_ReconcileResult_(Authority AuthorityState, Decision AuthorityResult, IsFresh bool) ReconcileResult {
	return ReconcileResult{ReconcileResult_ReconcileResult{Authority, Decision, IsFresh}}
}

func (_this ReconcileResult) Is_ReconcileResult() bool {
	_, ok := _this.Get_().(ReconcileResult_ReconcileResult)
	return ok
}

func (CompanionStruct_ReconcileResult_) Default(_default_S interface{}) ReconcileResult {
	return Companion_ReconcileResult_.Create_ReconcileResult_(Companion_AuthorityState_.Default(_default_S), Companion_AuthorityResult_.Default(), false)
}

func (_this ReconcileResult) Dtor_authority() AuthorityState {
	return _this.Get_().(ReconcileResult_ReconcileResult).Authority
}

func (_this ReconcileResult) Dtor_decision() AuthorityResult {
	return _this.Get_().(ReconcileResult_ReconcileResult).Decision
}

func (_this ReconcileResult) Dtor_isFresh() bool {
	return _this.Get_().(ReconcileResult_ReconcileResult).IsFresh
}

func (_this ReconcileResult) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ReconcileResult_ReconcileResult:
		{
			return "ConfluxCommand.ReconcileResult.ReconcileResult" + "(" + _dafny.String(data.Authority) + ", " + _dafny.String(data.Decision) + ", " + _dafny.String(data.IsFresh) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ReconcileResult) Equals(other ReconcileResult) bool {
	switch data1 := _this.Get_().(type) {
	case ReconcileResult_ReconcileResult:
		{
			data2, ok := other.Get_().(ReconcileResult_ReconcileResult)
			return ok && data1.Authority.Equals(data2.Authority) && data1.Decision.Equals(data2.Decision) && data1.IsFresh == data2.IsFresh
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ReconcileResult) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ReconcileResult)
	return ok && _this.Equals(typed)
}

func Type_ReconcileResult_(Type_S_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_ReconcileResult_{Type_S_}
}

type type_ReconcileResult_ struct {
	Type_S_ _dafny.TypeDescriptor
}

func (_this type_ReconcileResult_) Default() interface{} {
	Type_S_ := _this.Type_S_
	_ = Type_S_
	return Companion_ReconcileResult_.Default(Type_S_.Default())
}

func (_this type_ReconcileResult_) String() string {
	return "ConfluxCommand.ReconcileResult"
}
func (_this ReconcileResult) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ReconcileResult{}

// End of datatype ReconcileResult

// Definition of datatype AuthorityCodec
type AuthorityCodec struct {
	Data_AuthorityCodec_
}

func (_this AuthorityCodec) Get_() Data_AuthorityCodec_ {
	return _this.Data_AuthorityCodec_
}

type Data_AuthorityCodec_ interface {
	isAuthorityCodec()
}

type CompanionStruct_AuthorityCodec_ struct {
}

var Companion_AuthorityCodec_ = CompanionStruct_AuthorityCodec_{}

type AuthorityCodec_AuthorityCodec struct {
	Encode func(AuthorityState) _dafny.Sequence
	Decode func(_dafny.Sequence) m_ConfluxContract.Option
}

func (AuthorityCodec_AuthorityCodec) isAuthorityCodec() {}

func (CompanionStruct_AuthorityCodec_) Create_AuthorityCodec_(Encode func(AuthorityState) _dafny.Sequence, Decode func(_dafny.Sequence) m_ConfluxContract.Option) AuthorityCodec {
	return AuthorityCodec{AuthorityCodec_AuthorityCodec{Encode, Decode}}
}

func (_this AuthorityCodec) Is_AuthorityCodec() bool {
	_, ok := _this.Get_().(AuthorityCodec_AuthorityCodec)
	return ok
}

func (CompanionStruct_AuthorityCodec_) Default() AuthorityCodec {
	return Companion_AuthorityCodec_.Create_AuthorityCodec_(func(AuthorityState) _dafny.Sequence { return _dafny.EmptySeq }, func(_dafny.Sequence) m_ConfluxContract.Option { return m_ConfluxContract.Companion_Option_.Default() })
}

func (_this AuthorityCodec) Dtor_encode() func(AuthorityState) _dafny.Sequence {
	return _this.Get_().(AuthorityCodec_AuthorityCodec).Encode
}

func (_this AuthorityCodec) Dtor_decode() func(_dafny.Sequence) m_ConfluxContract.Option {
	return _this.Get_().(AuthorityCodec_AuthorityCodec).Decode
}

func (_this AuthorityCodec) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case AuthorityCodec_AuthorityCodec:
		{
			return "ConfluxCommand.AuthorityCodec.AuthorityCodec" + "(" + _dafny.String(data.Encode) + ", " + _dafny.String(data.Decode) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this AuthorityCodec) Equals(other AuthorityCodec) bool {
	switch data1 := _this.Get_().(type) {
	case AuthorityCodec_AuthorityCodec:
		{
			data2, ok := other.Get_().(AuthorityCodec_AuthorityCodec)
			return ok && _dafny.AreEqual(data1.Encode, data2.Encode) && _dafny.AreEqual(data1.Decode, data2.Decode)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this AuthorityCodec) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(AuthorityCodec)
	return ok && _this.Equals(typed)
}

func Type_AuthorityCodec_() _dafny.TypeDescriptor {
	return type_AuthorityCodec_{}
}

type type_AuthorityCodec_ struct {
}

func (_this type_AuthorityCodec_) Default() interface{} {
	return Companion_AuthorityCodec_.Default()
}

func (_this type_AuthorityCodec_) String() string {
	return "ConfluxCommand.AuthorityCodec"
}
func (_this AuthorityCodec) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = AuthorityCodec{}

// End of datatype AuthorityCodec

// Definition of datatype AuthorityResultCodec
type AuthorityResultCodec struct {
	Data_AuthorityResultCodec_
}

func (_this AuthorityResultCodec) Get_() Data_AuthorityResultCodec_ {
	return _this.Data_AuthorityResultCodec_
}

type Data_AuthorityResultCodec_ interface {
	isAuthorityResultCodec()
}

type CompanionStruct_AuthorityResultCodec_ struct {
}

var Companion_AuthorityResultCodec_ = CompanionStruct_AuthorityResultCodec_{}

type AuthorityResultCodec_AuthorityResultCodec struct {
	Encode func(AuthorityResult) _dafny.Sequence
	Decode func(_dafny.Sequence) m_ConfluxContract.Option
}

func (AuthorityResultCodec_AuthorityResultCodec) isAuthorityResultCodec() {}

func (CompanionStruct_AuthorityResultCodec_) Create_AuthorityResultCodec_(Encode func(AuthorityResult) _dafny.Sequence, Decode func(_dafny.Sequence) m_ConfluxContract.Option) AuthorityResultCodec {
	return AuthorityResultCodec{AuthorityResultCodec_AuthorityResultCodec{Encode, Decode}}
}

func (_this AuthorityResultCodec) Is_AuthorityResultCodec() bool {
	_, ok := _this.Get_().(AuthorityResultCodec_AuthorityResultCodec)
	return ok
}

func (CompanionStruct_AuthorityResultCodec_) Default() AuthorityResultCodec {
	return Companion_AuthorityResultCodec_.Create_AuthorityResultCodec_(func(AuthorityResult) _dafny.Sequence { return _dafny.EmptySeq }, func(_dafny.Sequence) m_ConfluxContract.Option { return m_ConfluxContract.Companion_Option_.Default() })
}

func (_this AuthorityResultCodec) Dtor_encode() func(AuthorityResult) _dafny.Sequence {
	return _this.Get_().(AuthorityResultCodec_AuthorityResultCodec).Encode
}

func (_this AuthorityResultCodec) Dtor_decode() func(_dafny.Sequence) m_ConfluxContract.Option {
	return _this.Get_().(AuthorityResultCodec_AuthorityResultCodec).Decode
}

func (_this AuthorityResultCodec) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case AuthorityResultCodec_AuthorityResultCodec:
		{
			return "ConfluxCommand.AuthorityResultCodec.AuthorityResultCodec" + "(" + _dafny.String(data.Encode) + ", " + _dafny.String(data.Decode) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this AuthorityResultCodec) Equals(other AuthorityResultCodec) bool {
	switch data1 := _this.Get_().(type) {
	case AuthorityResultCodec_AuthorityResultCodec:
		{
			data2, ok := other.Get_().(AuthorityResultCodec_AuthorityResultCodec)
			return ok && _dafny.AreEqual(data1.Encode, data2.Encode) && _dafny.AreEqual(data1.Decode, data2.Decode)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this AuthorityResultCodec) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(AuthorityResultCodec)
	return ok && _this.Equals(typed)
}

func Type_AuthorityResultCodec_() _dafny.TypeDescriptor {
	return type_AuthorityResultCodec_{}
}

type type_AuthorityResultCodec_ struct {
}

func (_this type_AuthorityResultCodec_) Default() interface{} {
	return Companion_AuthorityResultCodec_.Default()
}

func (_this type_AuthorityResultCodec_) String() string {
	return "ConfluxCommand.AuthorityResultCodec"
}
func (_this AuthorityResultCodec) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = AuthorityResultCodec{}

// End of datatype AuthorityResultCodec

// Definition of datatype CommandCheckpoint
type CommandCheckpoint struct {
	Data_CommandCheckpoint_
}

func (_this CommandCheckpoint) Get_() Data_CommandCheckpoint_ {
	return _this.Data_CommandCheckpoint_
}

type Data_CommandCheckpoint_ interface {
	isCommandCheckpoint()
}

type CompanionStruct_CommandCheckpoint_ struct {
}

var Companion_CommandCheckpoint_ = CompanionStruct_CommandCheckpoint_{}

type CommandCheckpoint_CommandCheckpoint struct {
	Durable m_ConfluxDurableHistory.ReplayCheckpoint
}

func (CommandCheckpoint_CommandCheckpoint) isCommandCheckpoint() {}

func (CompanionStruct_CommandCheckpoint_) Create_CommandCheckpoint_(Durable m_ConfluxDurableHistory.ReplayCheckpoint) CommandCheckpoint {
	return CommandCheckpoint{CommandCheckpoint_CommandCheckpoint{Durable}}
}

func (_this CommandCheckpoint) Is_CommandCheckpoint() bool {
	_, ok := _this.Get_().(CommandCheckpoint_CommandCheckpoint)
	return ok
}

func (CompanionStruct_CommandCheckpoint_) Default() m_ConfluxDurableHistory.ReplayCheckpoint {
	return m_ConfluxDurableHistory.Companion_ReplayCheckpoint_.Default()
}

func (_this CommandCheckpoint) Dtor_durable() m_ConfluxDurableHistory.ReplayCheckpoint {
	return _this.Get_().(CommandCheckpoint_CommandCheckpoint).Durable
}

func (_this CommandCheckpoint) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CommandCheckpoint_CommandCheckpoint:
		{
			return "ConfluxCommand.CommandCheckpoint.CommandCheckpoint" + "(" + _dafny.String(data.Durable) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CommandCheckpoint) Equals(other CommandCheckpoint) bool {
	switch data1 := _this.Get_().(type) {
	case CommandCheckpoint_CommandCheckpoint:
		{
			data2, ok := other.Get_().(CommandCheckpoint_CommandCheckpoint)
			return ok && data1.Durable.Equals(data2.Durable)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CommandCheckpoint) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CommandCheckpoint)
	return ok && _this.Equals(typed)
}

func Type_CommandCheckpoint_() _dafny.TypeDescriptor {
	return type_CommandCheckpoint_{}
}

type type_CommandCheckpoint_ struct {
}

func (_this type_CommandCheckpoint_) Default() interface{} {
	return Companion_CommandCheckpoint_.Default()
}

func (_this type_CommandCheckpoint_) String() string {
	return "ConfluxCommand.CommandCheckpoint"
}
func (_this CommandCheckpoint) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CommandCheckpoint{}

// End of datatype CommandCheckpoint
