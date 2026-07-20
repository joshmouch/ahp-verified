// AHP command specialization over the ONE shared Conflux command authority.
//
// This module owns only protocol types, lossless adapters, and AHP JSON.  It
// does not sequence, authorize, deduplicate, or persist commands; those
// semantics remain exclusively in ConfluxCommand.Reconcile.

include "ahp.dfy"

module AhpCommand {
  export
    provides Ahp, AhpSkeleton, Wire, Cmd, C, CO
    reveals CommandProposal, CommandDecision,
            AhpContentRef, ContentAddress, VerifiedContentRef,
            CreationOperation, CreationResponse
    reveals ToConfluxOrigin, FromConfluxOrigin,
            ToConfluxNatOption, FromConfluxNatOption,
            ToConfluxProposal, FromConfluxProposal,
            ToConfluxEnvelope, FromConfluxEnvelope,
            ToConfluxEnvelopes, FromConfluxEnvelopes,
            ToConfluxDecision, FromConfluxDecision
    reveals ToVerifiedContentRef, SemanticContentIdentity,
            ProjectCreationDecision
    provides
             EncodeAhpCommandJson, DecodeAhpCommandJson, MakeProposal,
             EncodeOrigin, DecodeOrigin,
             EncodeProposal, DecodeProposal,
             EncodeEnvelope, DecodeEnvelope,
             EncodeDecision, DecodeDecision
    provides OriginRoundTrip, ProposalRoundTrip, EnvelopeRoundTrip, DecisionRoundTrip

  import Ahp
  import opened AhpSkeleton
  import Wire = ConfluxCodec
  import C = ConfluxContract
  import F = Fidelity
  import CO = ConfluxOperators
  import Cmd = ConfluxCommand
  import Identity = ConfluxCommandIdentity

  datatype CommandProposal = CommandProposal(
    channel: string,
    origin: ActionOrigin,
    expectedServerSeq: Option<nat>,
    codecVersion: string,
    commandDigest: string,
    action: Ahp.AhpAction)

  datatype CommandDecision =
    | CommandRejected(origin: ActionOrigin, reason: Cmd.RejectionReason)
    | CommandAccepted(origin: ActionOrigin, envelopes: seq<Envelope<Ahp.AhpAction>>)

  // AHP ContentRef contributes only its optional locator. Immutable identity
  // is the separately verified content address; size/content-type/nonce drift
  // therefore cannot change semantic build identity.
  datatype AhpContentRef = AhpContentRef(uri: string, sizeHint: Option<nat>,
                                        contentType: Option<string>, nonce: Option<string>)
  datatype ContentAddress = ContentAddress(algorithm: string, digest: string)
  datatype VerifiedContentRef = VerifiedContentRef(
    address: ContentAddress, locator: C.Option<string>)

  function ToVerifiedContentRef(ref: AhpContentRef,
                                address: ContentAddress): VerifiedContentRef {
    VerifiedContentRef(address, if ref.uri == "" then C.None else C.Some(ref.uri))
  }

  function SemanticContentIdentity(ref: VerifiedContentRef): ContentAddress {
    ref.address
  }

  lemma LocatorAndNonceDoNotChangeIdentity(left: AhpContentRef,
                                           right: AhpContentRef,
                                           address: ContentAddress)
    ensures SemanticContentIdentity(ToVerifiedContentRef(left, address)) ==
            SemanticContentIdentity(ToVerifiedContentRef(right, address))
  {}

  // Collection-targeted createCharter and Commission-targeted createJudgment
  // are JSON-RPC projections of the same internal decision. Subsequent
  // dispatchAction remains an ordinary correlated command decision.
  datatype CreationOperation =
    | CreateCharter(collectionUri: string)
    | CreateJudgment(commissionUri: string)
  datatype CreationResponse =
    | CreationResult(origin: ActionOrigin, resourceUri: string)
    | CreationError(origin: ActionOrigin, reason: Cmd.RejectionReason)

  function ProjectCreationDecision(
    decision: Cmd.AuthorityResult<Ahp.AhpAction>,
    resourceUri: string): CreationResponse {
    match decision
    case Rejected(origin, reason) => CreationError(FromConfluxOrigin(origin), reason)
    case Accepted(origin, _) => CreationResult(FromConfluxOrigin(origin), resourceUri)
  }

  lemma CreationRejectionPreservesCorrelation(
    origin: Cmd.CommandOrigin, reason: Cmd.RejectionReason, resourceUri: string)
    ensures ProjectCreationDecision(Cmd.Rejected(origin, reason), resourceUri) ==
            CreationError(FromConfluxOrigin(origin), reason)
  {}

  function ToConfluxOrigin(origin: ActionOrigin): Cmd.CommandOrigin {
    Cmd.CommandOrigin(origin.clientId, origin.clientSeq)
  }

  function FromConfluxOrigin(origin: Cmd.CommandOrigin): ActionOrigin {
    ActionOrigin(origin.clientId, origin.clientSeq)
  }

  lemma OriginRoundTrip(origin: ActionOrigin)
    ensures FromConfluxOrigin(ToConfluxOrigin(origin)) == origin
  {}

  function ToConfluxNatOption(value: Option<nat>): C.Option<nat> {
    if value.Some? then C.Some(value.value) else C.None
  }

  function FromConfluxNatOption(value: C.Option<nat>): Option<nat> {
    if value.Some? then Some(value.value) else None
  }

  function ToConfluxProposal(proposal: CommandProposal): Cmd.Proposal<Ahp.AhpAction> {
    Cmd.Proposal(
      proposal.channel,
      ToConfluxOrigin(proposal.origin),
      ToConfluxNatOption(proposal.expectedServerSeq),
      proposal.codecVersion,
      proposal.commandDigest,
      proposal.action)
  }

  function FromConfluxProposal(proposal: Cmd.Proposal<Ahp.AhpAction>): CommandProposal {
    CommandProposal(
      proposal.channel,
      FromConfluxOrigin(proposal.origin),
      FromConfluxNatOption(proposal.expectedServerSeq),
      proposal.codecVersion,
      proposal.commandDigest,
      proposal.command)
  }

  function EncodeAhpCommandJson(action: Ahp.AhpAction): Wire.Json {
    Ahp.encodeAhpAction(action)
  }

  function DecodeAhpCommandJson(json: Wire.Json): Wire.Option<Ahp.AhpAction> {
    var action := Ahp.decodeAhpAction(json);
    if Ahp.encodeAhpAction(action) == json then Wire.Some(action) else Wire.None
  }

  function MakeProposal(
    channel: string,
    origin: ActionOrigin,
    expectedServerSeq: Option<nat>,
    action: Ahp.AhpAction): CommandProposal
  {
    var codec := Identity.CommandCodecFromJson(EncodeAhpCommandJson, DecodeAhpCommandJson);
    CommandProposal(channel, origin, expectedServerSeq,
      codec.codecVersion, codec.digest(codec.encode(action)), action)
  }

  lemma ProposalRoundTrip(proposal: CommandProposal)
    ensures FromConfluxProposal(ToConfluxProposal(proposal)) == proposal
  {}

  // Accepted Conflux envelopes map to ordinary AHP ActionEnvelope values with
  // an origin and no rejectionReason.  Rejection is an explicit correlated
  // CommandDecision and never receives a fabricated serverSeq.
  function FromConfluxEnvelope(
    envelope: Cmd.SequencedActionEnvelope<Ahp.AhpAction>): Envelope<Ahp.AhpAction>
  {
    Envelope(
      envelope.channel,
      envelope.action,
      envelope.serverSeq,
      Some(FromConfluxOrigin(envelope.origin)),
      None)
  }

  function ToConfluxEnvelope(
    envelope: Envelope<Ahp.AhpAction>): Cmd.SequencedActionEnvelope<Ahp.AhpAction>
  {
    Cmd.SequencedActionEnvelope(
      envelope.channel,
      envelope.serverSeq,
      if envelope.origin.Some? then ToConfluxOrigin(envelope.origin.value)
      else Cmd.CommandOrigin("", 0),
      envelope.action)
  }

  lemma EnvelopeRoundTrip(envelope: Cmd.SequencedActionEnvelope<Ahp.AhpAction>)
    ensures ToConfluxEnvelope(FromConfluxEnvelope(envelope)) == envelope
  {}

  function FromConfluxEnvelopes(
    envelopes: seq<Cmd.SequencedActionEnvelope<Ahp.AhpAction>>): seq<Envelope<Ahp.AhpAction>>
  {
    CO.Map(FromConfluxEnvelope, envelopes)
  }

  function ToConfluxEnvelopes(
    envelopes: seq<Envelope<Ahp.AhpAction>>): seq<Cmd.SequencedActionEnvelope<Ahp.AhpAction>>
  {
    CO.Map(ToConfluxEnvelope, envelopes)
  }

  lemma EnvelopesRoundTrip(envelopes: seq<Cmd.SequencedActionEnvelope<Ahp.AhpAction>>)
    ensures ToConfluxEnvelopes(FromConfluxEnvelopes(envelopes)) == envelopes
  {
    forall i | 0 <= i < |envelopes|
      ensures ToConfluxEnvelopes(FromConfluxEnvelopes(envelopes))[i] == envelopes[i]
    {
      EnvelopeRoundTrip(envelopes[i]);
    }
  }

  function FromConfluxDecision(
    decision: Cmd.AuthorityResult<Ahp.AhpAction>): CommandDecision
  {
    match decision
    case Rejected(origin, reason) =>
      CommandRejected(FromConfluxOrigin(origin), reason)
    case Accepted(origin, envelopes) =>
      CommandAccepted(FromConfluxOrigin(origin), FromConfluxEnvelopes(envelopes))
  }

  function ToConfluxDecision(
    decision: CommandDecision): Cmd.AuthorityResult<Ahp.AhpAction>
  {
    match decision
    case CommandRejected(origin, reason) =>
      Cmd.Rejected(ToConfluxOrigin(origin), reason)
    case CommandAccepted(origin, envelopes) =>
      Cmd.Accepted(ToConfluxOrigin(origin), ToConfluxEnvelopes(envelopes))
  }

  lemma DecisionRoundTrip(decision: Cmd.AuthorityResult<Ahp.AhpAction>)
    ensures ToConfluxDecision(FromConfluxDecision(decision)) == decision
  {
    match decision
    case Rejected(origin, _) => OriginRoundTrip(FromConfluxOrigin(origin));
    case Accepted(origin, envelopes) => {
      EnvelopesRoundTrip(envelopes);
    }
  }

  // ---------------- AHP-specific canonical JSON adapters ----------------

  function EncodeOrigin(origin: ActionOrigin): Wire.Json {
    Wire.JObj(map[
      "clientId" := Wire.JStr(origin.clientId),
      "clientSeq" := Wire.JNum(origin.clientSeq)])
  }

  function DecodeOrigin(j: Wire.Json): Option<ActionOrigin> {
    var clientId := F.asStr(F.field(j, "clientId"));
    var clientSeq := jsonInt(if F.field(j, "clientSeq").Some?
                             then F.field(j, "clientSeq").value else Wire.JNull);
    if j.JObj? && clientId != "" && clientSeq.Some? && clientSeq.value >= 0
    then Some(ActionOrigin(clientId, clientSeq.value as nat))
    else None
  }

  function EncodeOptionalNat(value: Option<nat>): Option<Wire.Json> {
    match value
    case None => None
    case Some(n) => Some(Wire.JNum(n))
  }

  function DecodeOptionalNat(fieldValue: Option<Wire.Json>): Option<Option<nat>> {
    match fieldValue
    case None => Some(None)
    case Some(j) =>
      var n := jsonInt(j);
      if n.Some? && n.value >= 0 then Some(Some(n.value as nat)) else None
  }

  function EncodeProposal(proposal: CommandProposal): Wire.Json {
    var fields := map[
      "channel" := Wire.JStr(proposal.channel),
      "origin" := EncodeOrigin(proposal.origin),
      "codecVersion" := Wire.JStr(proposal.codecVersion),
      "commandDigest" := Wire.JStr(proposal.commandDigest),
      "action" := Ahp.encodeAhpAction(proposal.action)];
    match EncodeOptionalNat(proposal.expectedServerSeq)
    case None => Wire.JObj(fields)
    case Some(j) => Wire.JObj(fields["expectedServerSeq" := j])
  }

  function DecodeProposal(j: Wire.Json): Option<CommandProposal> {
    var channel := F.asStr(F.field(j, "channel"));
    var origin := DecodeOrigin(if F.field(j, "origin").Some?
                               then F.field(j, "origin").value else Wire.JNull);
    var expected := DecodeOptionalNat(F.field(j, "expectedServerSeq"));
    var codecVersion := F.asStr(F.field(j, "codecVersion"));
    var commandDigest := F.asStr(F.field(j, "commandDigest"));
    var actionField := F.field(j, "action");
    if j.JObj? && channel != "" && origin.Some? && expected.Some? &&
       codecVersion != "" && commandDigest != "" && actionField.Some?
    then Some(CommandProposal(
      channel, origin.value, expected.value, codecVersion, commandDigest,
      Ahp.decodeAhpAction(actionField.value)))
    else None
  }

  function EncodeRejection(reason: Cmd.RejectionReason): Wire.Json {
    var fields := map[
      "code" := Wire.JStr(reason.code),
      "message" := Wire.JStr(reason.message)];
    match reason.detailsDigest
    case None => Wire.JObj(fields)
    case Some(digest) => Wire.JObj(fields["detailsDigest" := Wire.JStr(digest)])
  }

  function DecodeRejection(j: Wire.Json): Option<Cmd.RejectionReason> {
    var code := F.asStr(F.field(j, "code"));
    var message := F.asStr(F.field(j, "message"));
    var detailsField := F.field(j, "detailsDigest");
    var details: C.Option<string> := if detailsField.Some? && detailsField.value.JStr?
                   then C.Some(detailsField.value.s) else C.None;
    if j.JObj? && code != "" && message != ""
    then Some(Cmd.RejectionReason(code, message, details))
    else None
  }

  function EncodeEnvelope(envelope: Envelope<Ahp.AhpAction>): Wire.Json
  {
    Wire.JObj(map[
      "channel" := Wire.JStr(envelope.channel),
      "action" := Ahp.encodeAhpAction(envelope.action),
      "serverSeq" := Wire.JNum(envelope.serverSeq),
      "origin" := if envelope.origin.Some? then EncodeOrigin(envelope.origin.value) else Wire.JNull])
  }

  function DecodeEnvelope(j: Wire.Json): Option<Envelope<Ahp.AhpAction>> {
    var channel := F.asStr(F.field(j, "channel"));
    var seqValue := jsonInt(if F.field(j, "serverSeq").Some?
                            then F.field(j, "serverSeq").value else Wire.JNull);
    var origin := DecodeOrigin(if F.field(j, "origin").Some?
                               then F.field(j, "origin").value else Wire.JNull);
    var actionField := F.field(j, "action");
    if j.JObj? && channel != "" && seqValue.Some? && seqValue.value >= 0 &&
       origin.Some? && actionField.Some?
    then Some(Envelope(
      channel,
      Ahp.decodeAhpAction(actionField.value),
      seqValue.value as nat,
      Some(origin.value),
      None))
    else None
  }

  function EncodeEnvelopes(envelopes: seq<Envelope<Ahp.AhpAction>>): seq<Wire.Json>
  {
    CO.Map(EncodeEnvelope, envelopes)
  }

  function DecodeEnvelopes(js: seq<Wire.Json>): Option<seq<Envelope<Ahp.AhpAction>>> {
    var decoded := CO.Map(DecodeEnvelope, js);
    if forall i | 0 <= i < |decoded| :: decoded[i].Some?
    then Some(seq(|decoded|, i requires 0 <= i < |decoded| => decoded[i].value))
    else None
  }

  function EncodeDecision(decision: CommandDecision): Wire.Json
  {
    match decision
    case CommandRejected(origin, reason) => Wire.JObj(map[
      "type" := Wire.JStr("rejected"),
      "origin" := EncodeOrigin(origin),
      "reason" := EncodeRejection(reason)])
    case CommandAccepted(origin, envelopes) => Wire.JObj(map[
      "type" := Wire.JStr("accepted"),
      "origin" := EncodeOrigin(origin),
      "actions" := Wire.JArr(EncodeEnvelopes(envelopes))])
  }

  function DecodeDecision(j: Wire.Json): Option<CommandDecision> {
    var tag := F.asStr(F.field(j, "type"));
    var origin := DecodeOrigin(if F.field(j, "origin").Some?
                               then F.field(j, "origin").value else Wire.JNull);
    if !j.JObj? || origin.None? then None
    else if tag == "rejected" then
      var reason := DecodeRejection(if F.field(j, "reason").Some?
                                    then F.field(j, "reason").value else Wire.JNull);
      if reason.Some? then Some(CommandRejected(origin.value, reason.value)) else None
    else if tag == "accepted" then
      var actionsField := F.field(j, "actions");
      if actionsField.Some? && actionsField.value.JArr? then
        var envelopes := DecodeEnvelopes(actionsField.value.elems);
        if envelopes.Some? then Some(CommandAccepted(origin.value, envelopes.value)) else None
      else None
    else None
  }
}
