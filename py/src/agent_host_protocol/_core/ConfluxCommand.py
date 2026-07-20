import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_ as module_
import _dafny as _dafny
import System_ as System_
import ConfluxCodec as ConfluxCodec
import ConfluxContract as ConfluxContract
import AhpSkeleton as AhpSkeleton
import ResourceWatch as ResourceWatch
import Canvas as Canvas
import ConfluxOperators as ConfluxOperators
import ConfluxPrune as ConfluxPrune
import ConfluxSeqRoute as ConfluxSeqRoute
import Changeset as Changeset
import Annotations as Annotations
import Terminal as Terminal
import Session as Session
import ConfluxOrderedLog as ConfluxOrderedLog
import Chat as Chat
import ClientMain as ClientMain
import ConfluxSemanticGraphIdentity as ConfluxSemanticGraphIdentity
import ConfluxContentAddress as ConfluxContentAddress
import ConfluxSemanticGraphContract as ConfluxSemanticGraphContract
import ConfluxResourceCeilings as ConfluxResourceCeilings
import ConfluxSupervisedOperationResult as ConfluxSupervisedOperationResult
import ConfluxTree as ConfluxTree
import ConfluxMerge as ConfluxMerge
import ConfluxRoute as ConfluxRoute
import ConfluxMapValues as ConfluxMapValues
import ConfluxDelimited as ConfluxDelimited
import ConfluxFixpoint as ConfluxFixpoint
import ConfluxGenericCodec as ConfluxGenericCodec
import ConfluxResolve as ConfluxResolve
import ConfluxWatermark as ConfluxWatermark
import ConfluxStateMachine as ConfluxStateMachine
import ConfluxProduct as ConfluxProduct
import ConfluxJoin as ConfluxJoin
import ConfluxDedupe as ConfluxDedupe
import ConfluxBatchRoute as ConfluxBatchRoute
import ConfluxSeqRouteMerge as ConfluxSeqRouteMerge
import ConfluxKeyedOps as ConfluxKeyedOps
import ConfluxFilterKeys as ConfluxFilterKeys
import ConfluxChannel as ConfluxChannel
import ConfluxDurableHistory as ConfluxDurableHistory

# Module: ConfluxCommand

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def CanonicalCommandCodecVersion():
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oa-canonical-json-v1"))

    @staticmethod
    def IsLowerHexChar(c):
        return (((_dafny.CodePoint('0')) <= (c)) and ((c) <= (_dafny.CodePoint('9')))) or (((_dafny.CodePoint('a')) <= (c)) and ((c) <= (_dafny.CodePoint('f'))))

    @staticmethod
    def IsLowerHex(s):
        return ((len(s)) == (0)) or ((default__.IsLowerHexChar((s)[0])) and (default__.IsLowerHex(_dafny.SeqWithoutIsStrInference((s)[1::]))))

    @staticmethod
    def RepeatChar(c, count):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (count) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([c]))
                    in0_ = c
                    in1_ = (count) - (1)
                    c = in0_
                    count = in1_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def ValidSha256Digest(digest):
        return (((len(digest)) == (71)) and ((_dafny.SeqWithoutIsStrInference((digest)[:7:])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sha256:"))))) and (default__.IsLowerHex(_dafny.SeqWithoutIsStrInference((digest)[7::])))

    @staticmethod
    def CanonicalCommandBytes(codec, proposal):
        return (codec).encode((proposal).command)

    @staticmethod
    def ValidProposalIdentity(codec, proposal):
        return ((((((codec).codecVersion) == (default__.CanonicalCommandCodecVersion())) and (((proposal).codecVersion) == ((codec).codecVersion))) and (((codec).decode(default__.CanonicalCommandBytes(codec, proposal))) == (ConfluxContract.Option_Some((proposal).command)))) and (default__.ValidSha256Digest((proposal).commandDigest))) and (((proposal).commandDigest) == ((codec).digest(default__.CanonicalCommandBytes(codec, proposal))))

    @staticmethod
    def IdentityOf(codec, proposal):
        d_0_bytes_ = default__.CanonicalCommandBytes(codec, proposal)
        return ProposalIdentity_ProposalIdentity(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "conflux-command-v1")), (codec).codecVersion, (proposal).channel, (proposal).expectedServerSeq, (codec).digest(d_0_bytes_), d_0_bytes_)

    @staticmethod
    def FindRecord(history, origin):
        while True:
            with _dafny.label():
                if (len((history))) == (0):
                    return ConfluxContract.Option_None()
                elif ((((history))[0]).origin) == (origin):
                    return ConfluxContract.Option_Some(((history))[0])
                elif True:
                    in0_ = _dafny.SeqWithoutIsStrInference(((history))[1::])
                    in1_ = origin
                    history = in0_
                    origin = in1_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def FindDecision(history, origin):
        source0_ = default__.FindRecord(history, origin)
        if True:
            if source0_.is_None:
                return ConfluxContract.Option_None()
        if True:
            d_0_record_ = source0_.value
            return ConfluxContract.Option_Some((d_0_record_).result)

    @staticmethod
    def RecordDecision(history, channel, origin, proposalIdentity, result):
        d_0_appended_ = ConfluxDurableHistory.default__.Append((history), DecisionRecord_DecisionRecord(channel, origin, proposalIdentity, result))
        return (d_0_appended_)

    @staticmethod
    def StampActions(channel, origin, nextSeq, actions):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(actions)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([SequencedActionEnvelope_SequencedActionEnvelope(channel, nextSeq, origin, (actions)[0])]))
                    in0_ = channel
                    in1_ = origin
                    in2_ = (nextSeq) + (1)
                    in3_ = _dafny.SeqWithoutIsStrInference((actions)[1::])
                    channel = in0_
                    origin = in1_
                    nextSeq = in2_
                    actions = in3_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def EnvelopeActions(envelopes):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(envelopes)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([((envelopes)[0]).action]))
                    in0_ = _dafny.SeqWithoutIsStrInference((envelopes)[1::])
                    envelopes = in0_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def DecisionActions(decision):
        source0_ = decision
        if True:
            if source0_.is_Rejected:
                return _dafny.SeqWithoutIsStrInference([])
        if True:
            d_0_envelopes_ = source0_.sequencedActions
            return d_0_envelopes_

    @staticmethod
    def RecordAcceptedActions(record):
        return default__.DecisionActions((record).result)

    @staticmethod
    def AcceptedHistory(history):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len((history))) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (default__.RecordAcceptedActions(((history))[0]))
                    in0_ = _dafny.SeqWithoutIsStrInference(((history))[1::])
                    history = in0_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def ExpectedSequenceMatches(state, proposal):
        source0_ = (proposal).expectedServerSeq
        if True:
            if source0_.is_None:
                return True
        if True:
            d_0_expected_ = source0_.value
            return (d_0_expected_) == (((state).host).nextSeq)

    @staticmethod
    def StaleExpectedServerSeq():
        return RejectionReason_RejectionReason(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "staleExpectedServerSeq")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "expected sequence does not match")), ConfluxContract.Option_None())

    @staticmethod
    def IdempotencyOriginConflict(recorded, attempted):
        return RejectionReason_RejectionReason(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idempotencyOriginConflict")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "command origin is already bound to a different proposal")), ConfluxContract.Option_Some((((recorded).commandDigest) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ":")))) + ((attempted).commandDigest)))

    @staticmethod
    def InvalidProposalIdentity(codecVersion):
        return RejectionReason_RejectionReason(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "invalidProposalIdentity")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "command codec version, canonical encoding round-trip, or digest is invalid")), ConfluxContract.Option_Some(codecVersion))

    @staticmethod
    def EmptyAuthorizedActions():
        return RejectionReason_RejectionReason(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emptyAuthorizedActions")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mutating command authority requires at least one action")), ConfluxContract.Option_None())

    @staticmethod
    def RecordRejection(state, proposal, identity, reason):
        d_0_decision_ = AuthorityResult_Rejected((proposal).origin, reason)
        return ReconcileResult_ReconcileResult(AuthorityState_AuthorityState((state).host, default__.RecordDecision((state).history, (proposal).channel, (proposal).origin, identity, d_0_decision_)), d_0_decision_, True)

    @staticmethod
    def Reconcile(reducer, state, proposal, authorizeAndLower, commandCodec):
        if not(default__.ValidProposalIdentity(commandCodec, proposal)):
            return ReconcileResult_ReconcileResult(state, AuthorityResult_Rejected((proposal).origin, default__.InvalidProposalIdentity((proposal).codecVersion)), False)
        elif True:
            source0_ = default__.FindRecord((state).history, (proposal).origin)
            if True:
                if source0_.is_Some:
                    d_0_recorded_ = source0_.value
                    if ((d_0_recorded_).proposalIdentity) == (default__.IdentityOf(commandCodec, proposal)):
                        return ReconcileResult_ReconcileResult(state, (d_0_recorded_).result, False)
                    elif True:
                        return ReconcileResult_ReconcileResult(state, AuthorityResult_Rejected((proposal).origin, default__.IdempotencyOriginConflict((d_0_recorded_).proposalIdentity, default__.IdentityOf(commandCodec, proposal))), False)
            if True:
                if not(default__.ExpectedSequenceMatches(state, proposal)):
                    return default__.RecordRejection(state, proposal, default__.IdentityOf(commandCodec, proposal), default__.StaleExpectedServerSeq())
                elif True:
                    source1_ = authorizeAndLower((state).host, (proposal).command)
                    if True:
                        if source1_.is_Denied:
                            d_1_reason_ = source1_.reason
                            return default__.RecordRejection(state, proposal, default__.IdentityOf(commandCodec, proposal), d_1_reason_)
                    if True:
                        d_2_actions_ = source1_.actions
                        if (len(d_2_actions_)) == (0):
                            return default__.RecordRejection(state, proposal, default__.IdentityOf(commandCodec, proposal), default__.EmptyAuthorizedActions())
                        elif True:
                            d_3_envelopes_ = default__.StampActions((proposal).channel, (proposal).origin, ((state).host).nextSeq, d_2_actions_)
                            d_4_decision_ = AuthorityResult_Accepted((proposal).origin, d_3_envelopes_)
                            return ReconcileResult_ReconcileResult(AuthorityState_AuthorityState(ConfluxChannel.default__.AcceptFold(reducer, (state).host, d_2_actions_), default__.RecordDecision((state).history, (proposal).channel, (proposal).origin, default__.IdentityOf(commandCodec, proposal), d_4_decision_)), d_4_decision_, True)

    @staticmethod
    def MakeCheckpoint(historyId, reducerSchemaVersion, hashBytes, codec, state):
        def lambda0_(d_0_authority_):
            return ((d_0_authority_).host).nextSeq

        return ConfluxDurableHistory.default__.MakeCheckpoint(historyId, reducerSchemaVersion, lambda0_, hashBytes, ConfluxDurableHistory.StateCodec_StateCodec((codec).encode, (codec).decode), state)

    @staticmethod
    def RestoreCheckpoint(expectedHistoryId, expectedReducerSchemaVersion, hashBytes, codec, checkpoint):
        def lambda0_(d_0_authority_):
            return ((d_0_authority_).host).nextSeq

        return ConfluxDurableHistory.default__.RestoreCheckpoint(expectedHistoryId, expectedReducerSchemaVersion, lambda0_, hashBytes, ConfluxDurableHistory.StateCodec_StateCodec((codec).encode, (codec).decode), (checkpoint))

    @staticmethod
    def EnvelopesPreserveOrigin(origin, envelopes):
        return ((len(envelopes)) == (0)) or (((((envelopes)[0]).origin) == (origin)) and (default__.EnvelopesPreserveOrigin(origin, _dafny.SeqWithoutIsStrInference((envelopes)[1::]))))

    @staticmethod
    def ChannelsPreserved(channel, envelopes):
        return ((len(envelopes)) == (0)) or (((((envelopes)[0]).channel) == (channel)) and (default__.ChannelsPreserved(channel, _dafny.SeqWithoutIsStrInference((envelopes)[1::]))))

    @staticmethod
    def SequencesContiguous(start, envelopes):
        return ((len(envelopes)) == (0)) or (((((envelopes)[0]).serverSeq) == (start)) and (default__.SequencesContiguous((start) + (1), _dafny.SeqWithoutIsStrInference((envelopes)[1::]))))


class CommandOrigin:
    @classmethod
    def default(cls, ):
        return lambda: CommandOrigin_CommandOrigin(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_CommandOrigin(self) -> bool:
        return isinstance(self, CommandOrigin_CommandOrigin)

class CommandOrigin_CommandOrigin(CommandOrigin, NamedTuple('CommandOrigin', [('clientId', Any), ('clientSeq', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.CommandOrigin.CommandOrigin({self.clientId.VerbatimString(True)}, {_dafny.string_of(self.clientSeq)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CommandOrigin_CommandOrigin) and self.clientId == __o.clientId and self.clientSeq == __o.clientSeq
    def __hash__(self) -> int:
        return super().__hash__()


class Proposal:
    @classmethod
    def default(cls, default_C):
        return lambda: Proposal_Proposal(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), CommandOrigin.default()(), ConfluxContract.Option.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), default_C())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Proposal(self) -> bool:
        return isinstance(self, Proposal_Proposal)

class Proposal_Proposal(Proposal, NamedTuple('Proposal', [('channel', Any), ('origin', Any), ('expectedServerSeq', Any), ('codecVersion', Any), ('commandDigest', Any), ('command', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.Proposal.Proposal({self.channel.VerbatimString(True)}, {_dafny.string_of(self.origin)}, {_dafny.string_of(self.expectedServerSeq)}, {self.codecVersion.VerbatimString(True)}, {self.commandDigest.VerbatimString(True)}, {_dafny.string_of(self.command)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Proposal_Proposal) and self.channel == __o.channel and self.origin == __o.origin and self.expectedServerSeq == __o.expectedServerSeq and self.codecVersion == __o.codecVersion and self.commandDigest == __o.commandDigest and self.command == __o.command
    def __hash__(self) -> int:
        return super().__hash__()


class ProposalIdentity:
    @classmethod
    def default(cls, ):
        return lambda: ProposalIdentity_ProposalIdentity(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxContract.Option.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ProposalIdentity(self) -> bool:
        return isinstance(self, ProposalIdentity_ProposalIdentity)

class ProposalIdentity_ProposalIdentity(ProposalIdentity, NamedTuple('ProposalIdentity', [('protocol', Any), ('codecVersion', Any), ('channel', Any), ('expectedServerSeq', Any), ('commandDigest', Any), ('canonicalCommandBytes', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.ProposalIdentity.ProposalIdentity({self.protocol.VerbatimString(True)}, {self.codecVersion.VerbatimString(True)}, {self.channel.VerbatimString(True)}, {_dafny.string_of(self.expectedServerSeq)}, {self.commandDigest.VerbatimString(True)}, {_dafny.string_of(self.canonicalCommandBytes)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ProposalIdentity_ProposalIdentity) and self.protocol == __o.protocol and self.codecVersion == __o.codecVersion and self.channel == __o.channel and self.expectedServerSeq == __o.expectedServerSeq and self.commandDigest == __o.commandDigest and self.canonicalCommandBytes == __o.canonicalCommandBytes
    def __hash__(self) -> int:
        return super().__hash__()


class CommandIdentityCodec:
    @classmethod
    def default(cls, ):
        return lambda: CommandIdentityCodec_CommandIdentityCodec(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), (lambda x12: _dafny.Seq({})), (lambda x13: ConfluxContract.Option.default()()), (lambda x14: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_CommandIdentityCodec(self) -> bool:
        return isinstance(self, CommandIdentityCodec_CommandIdentityCodec)

class CommandIdentityCodec_CommandIdentityCodec(CommandIdentityCodec, NamedTuple('CommandIdentityCodec', [('codecVersion', Any), ('encode', Any), ('decode', Any), ('digest', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.CommandIdentityCodec.CommandIdentityCodec({self.codecVersion.VerbatimString(True)}, {_dafny.string_of(self.encode)}, {_dafny.string_of(self.decode)}, {_dafny.string_of(self.digest)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CommandIdentityCodec_CommandIdentityCodec) and self.codecVersion == __o.codecVersion and self.encode == __o.encode and self.decode == __o.decode and self.digest == __o.digest
    def __hash__(self) -> int:
        return super().__hash__()


class RejectionReason:
    @classmethod
    def default(cls, ):
        return lambda: RejectionReason_RejectionReason(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxContract.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_RejectionReason(self) -> bool:
        return isinstance(self, RejectionReason_RejectionReason)

class RejectionReason_RejectionReason(RejectionReason, NamedTuple('RejectionReason', [('code', Any), ('message', Any), ('detailsDigest', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.RejectionReason.RejectionReason({self.code.VerbatimString(True)}, {self.message.VerbatimString(True)}, {_dafny.string_of(self.detailsDigest)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, RejectionReason_RejectionReason) and self.code == __o.code and self.message == __o.message and self.detailsDigest == __o.detailsDigest
    def __hash__(self) -> int:
        return super().__hash__()


class SequencedActionEnvelope:
    @classmethod
    def default(cls, default_A):
        return lambda: SequencedActionEnvelope_SequencedActionEnvelope(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), CommandOrigin.default()(), default_A())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SequencedActionEnvelope(self) -> bool:
        return isinstance(self, SequencedActionEnvelope_SequencedActionEnvelope)

class SequencedActionEnvelope_SequencedActionEnvelope(SequencedActionEnvelope, NamedTuple('SequencedActionEnvelope', [('channel', Any), ('serverSeq', Any), ('origin', Any), ('action', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.SequencedActionEnvelope.SequencedActionEnvelope({self.channel.VerbatimString(True)}, {_dafny.string_of(self.serverSeq)}, {_dafny.string_of(self.origin)}, {_dafny.string_of(self.action)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SequencedActionEnvelope_SequencedActionEnvelope) and self.channel == __o.channel and self.serverSeq == __o.serverSeq and self.origin == __o.origin and self.action == __o.action
    def __hash__(self) -> int:
        return super().__hash__()


class AuthorityResult:
    @classmethod
    def default(cls, ):
        return lambda: AuthorityResult_Rejected(CommandOrigin.default()(), RejectionReason.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Rejected(self) -> bool:
        return isinstance(self, AuthorityResult_Rejected)
    @property
    def is_Accepted(self) -> bool:
        return isinstance(self, AuthorityResult_Accepted)

class AuthorityResult_Rejected(AuthorityResult, NamedTuple('Rejected', [('origin', Any), ('reason', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.AuthorityResult.Rejected({_dafny.string_of(self.origin)}, {_dafny.string_of(self.reason)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AuthorityResult_Rejected) and self.origin == __o.origin and self.reason == __o.reason
    def __hash__(self) -> int:
        return super().__hash__()

class AuthorityResult_Accepted(AuthorityResult, NamedTuple('Accepted', [('origin', Any), ('sequencedActions', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.AuthorityResult.Accepted({_dafny.string_of(self.origin)}, {_dafny.string_of(self.sequencedActions)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AuthorityResult_Accepted) and self.origin == __o.origin and self.sequencedActions == __o.sequencedActions
    def __hash__(self) -> int:
        return super().__hash__()


class Authorization:
    @classmethod
    def default(cls, ):
        return lambda: Authorization_Denied(RejectionReason.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Denied(self) -> bool:
        return isinstance(self, Authorization_Denied)
    @property
    def is_Allowed(self) -> bool:
        return isinstance(self, Authorization_Allowed)

class Authorization_Denied(Authorization, NamedTuple('Denied', [('reason', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.Authorization.Denied({_dafny.string_of(self.reason)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Authorization_Denied) and self.reason == __o.reason
    def __hash__(self) -> int:
        return super().__hash__()

class Authorization_Allowed(Authorization, NamedTuple('Allowed', [('actions', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.Authorization.Allowed({_dafny.string_of(self.actions)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Authorization_Allowed) and self.actions == __o.actions
    def __hash__(self) -> int:
        return super().__hash__()


class DecisionRecord:
    @classmethod
    def default(cls, ):
        return lambda: DecisionRecord_DecisionRecord(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), CommandOrigin.default()(), ProposalIdentity.default()(), AuthorityResult.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DecisionRecord(self) -> bool:
        return isinstance(self, DecisionRecord_DecisionRecord)

class DecisionRecord_DecisionRecord(DecisionRecord, NamedTuple('DecisionRecord', [('channel', Any), ('origin', Any), ('proposalIdentity', Any), ('result', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.DecisionRecord.DecisionRecord({self.channel.VerbatimString(True)}, {_dafny.string_of(self.origin)}, {_dafny.string_of(self.proposalIdentity)}, {_dafny.string_of(self.result)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, DecisionRecord_DecisionRecord) and self.channel == __o.channel and self.origin == __o.origin and self.proposalIdentity == __o.proposalIdentity and self.result == __o.result
    def __hash__(self) -> int:
        return super().__hash__()


class DecisionHistory:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DecisionHistory(self) -> bool:
        return isinstance(self, DecisionHistory_DecisionHistory)

class DecisionHistory_DecisionHistory(DecisionHistory, NamedTuple('DecisionHistory', [('records', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.DecisionHistory.DecisionHistory({_dafny.string_of(self.records)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, DecisionHistory_DecisionHistory) and self.records == __o.records
    def __hash__(self) -> int:
        return super().__hash__()


class AuthorityState:
    @classmethod
    def default(cls, default_S):
        return lambda: AuthorityState_AuthorityState(ConfluxChannel.HostState.default(default_S)(), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_AuthorityState(self) -> bool:
        return isinstance(self, AuthorityState_AuthorityState)

class AuthorityState_AuthorityState(AuthorityState, NamedTuple('AuthorityState', [('host', Any), ('history', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.AuthorityState.AuthorityState({_dafny.string_of(self.host)}, {_dafny.string_of(self.history)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AuthorityState_AuthorityState) and self.host == __o.host and self.history == __o.history
    def __hash__(self) -> int:
        return super().__hash__()


class ReconcileResult:
    @classmethod
    def default(cls, default_S):
        return lambda: ReconcileResult_ReconcileResult(AuthorityState.default(default_S)(), AuthorityResult.default()(), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ReconcileResult(self) -> bool:
        return isinstance(self, ReconcileResult_ReconcileResult)

class ReconcileResult_ReconcileResult(ReconcileResult, NamedTuple('ReconcileResult', [('authority', Any), ('decision', Any), ('isFresh', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.ReconcileResult.ReconcileResult({_dafny.string_of(self.authority)}, {_dafny.string_of(self.decision)}, {_dafny.string_of(self.isFresh)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ReconcileResult_ReconcileResult) and self.authority == __o.authority and self.decision == __o.decision and self.isFresh == __o.isFresh
    def __hash__(self) -> int:
        return super().__hash__()


class AuthorityCodec:
    @classmethod
    def default(cls, ):
        return lambda: AuthorityCodec_AuthorityCodec((lambda x15: _dafny.Seq({})), (lambda x16: ConfluxContract.Option.default()()))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_AuthorityCodec(self) -> bool:
        return isinstance(self, AuthorityCodec_AuthorityCodec)

class AuthorityCodec_AuthorityCodec(AuthorityCodec, NamedTuple('AuthorityCodec', [('encode', Any), ('decode', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.AuthorityCodec.AuthorityCodec({_dafny.string_of(self.encode)}, {_dafny.string_of(self.decode)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AuthorityCodec_AuthorityCodec) and self.encode == __o.encode and self.decode == __o.decode
    def __hash__(self) -> int:
        return super().__hash__()


class AuthorityResultCodec:
    @classmethod
    def default(cls, ):
        return lambda: AuthorityResultCodec_AuthorityResultCodec((lambda x17: _dafny.Seq({})), (lambda x18: ConfluxContract.Option.default()()))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_AuthorityResultCodec(self) -> bool:
        return isinstance(self, AuthorityResultCodec_AuthorityResultCodec)

class AuthorityResultCodec_AuthorityResultCodec(AuthorityResultCodec, NamedTuple('AuthorityResultCodec', [('encode', Any), ('decode', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.AuthorityResultCodec.AuthorityResultCodec({_dafny.string_of(self.encode)}, {_dafny.string_of(self.decode)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AuthorityResultCodec_AuthorityResultCodec) and self.encode == __o.encode and self.decode == __o.decode
    def __hash__(self) -> int:
        return super().__hash__()


class CommandCheckpoint:
    @classmethod
    def default(cls, ):
        return lambda: ConfluxDurableHistory.ReplayCheckpoint.default()()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_CommandCheckpoint(self) -> bool:
        return isinstance(self, CommandCheckpoint_CommandCheckpoint)

class CommandCheckpoint_CommandCheckpoint(CommandCheckpoint, NamedTuple('CommandCheckpoint', [('durable', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxCommand.CommandCheckpoint.CommandCheckpoint({_dafny.string_of(self.durable)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CommandCheckpoint_CommandCheckpoint) and self.durable == __o.durable
    def __hash__(self) -> int:
        return super().__hash__()

