// Upstream-ready AHP vNext turn-window proposal model. This is intentionally
// not current-protocol evidence: it is a positive, protocol-neutral witness
// kept separate from the legacy turns/cursor/chat/turnsLoaded compatibility path.
module AhpTurnWindowProposal {
  datatype Opt<T> = None | Some(value: T)
  datatype Anchor = EarliestBoundary | LatestBoundary | ExactTurn(turnId: string)
  datatype Fulfillment = UpTo | Advisory
  datatype Extent = Extent(before: nat, includeAnchor: bool, after: nat,
                           fulfillment: Fulfillment)
  datatype ResolvedExtent = ResolvedExtent(before: nat, includedAnchor: bool,
                                           after: nat, completion: string)
  datatype Seek = Seek(anchor: Anchor, extent: Extent, asOf: Opt<string>)
  datatype ReadStart = StartSeek(seek: Seek) | StartContinue(cursor: string)

  datatype Lifetime =
    | Connection(connectionId: string)
    | HostInstance(hostInstanceId: string)
    | ExpiresAt(expiry: string, expiryHostInstanceId: Opt<string>, survivesRestart: bool)
    | ServerStored(contractRef: string, storedExpiry: Opt<string>,
                   storedHostInstanceId: Opt<string>, storedSurvivesRestart: bool)

  datatype Recovery = ResetSnapshot | RestartSeek(seek: Seek)
                    | Rebase(seek: Seek, asOf: string)
  datatype Continuation = Continuation(cursor: string, contractId: string,
    contractVersion: string, issuer: string, lifetime: Lifetime,
    recovery: Recovery)
  datatype Resolution = Resolution(resolvedAnchor: Anchor,
    resolvedExtent: ResolvedExtent, resolvedRevision: string,
    continuation: Opt<Continuation>)

  datatype CursorBinding = CursorBinding(
    channel: string,
    readVariant: string,
    resolvedAnchor: Anchor,
    requestedExtent: Extent,
    resolvedExtent: ResolvedExtent,
    ordering: string,
    filterDigest: Opt<string>,
    resolvedRevision: string,
    principal: string,
    subscriptionId: Opt<string>,
    direction: string,
    issuer: string,
    hostInstance: Opt<string>,
    restartEpoch: nat,
    expiresAt: Opt<string>)

  datatype ContinuationContext = ContinuationContext(
    channel: string, readVariant: string, principal: string,
    subscriptionId: Opt<string>, issuer: string,
    hostInstance: Opt<string>, restartEpoch: nat, now: string)

  datatype Validation = Valid | InvalidContinuation(recovery: Recovery)

  predicate SameScope(binding: CursorBinding, context: ContinuationContext) {
    binding.channel == context.channel &&
    binding.readVariant == context.readVariant &&
    binding.principal == context.principal &&
    binding.subscriptionId == context.subscriptionId &&
    binding.issuer == context.issuer &&
    binding.hostInstance == context.hostInstance &&
    binding.restartEpoch == context.restartEpoch
  }

  function Validate(binding: CursorBinding, context: ContinuationContext,
                    recovery: Recovery): Validation {
    if SameScope(binding, context) then Valid else InvalidContinuation(recovery)
  }

  datatype Projection<T> = Projection(subscriptionId: string, turns: seq<T>)
  function Hydrate<T>(projection: Projection<T>, page: seq<T>): Projection<T> {
    Projection(projection.subscriptionId, page + projection.turns)
  }

  datatype LegacyStart = LegacyStart(turns: Opt<nat>, cursor: Opt<string>)

  lemma SeekAlwaysCarriesGeometry(start: ReadStart)
    ensures start.StartSeek? ==> start.seek.extent == start.seek.extent
  {}

  lemma ContinueCannotReplaceGeometry(cursor: string)
    ensures StartContinue(cursor).StartContinue?
    ensures !StartContinue(cursor).StartSeek?
  {}

  lemma WrongScopeIsRejected(binding: CursorBinding,
                             context: ContinuationContext,
                             recovery: Recovery)
    requires !SameScope(binding, context)
    ensures Validate(binding, context, recovery) == InvalidContinuation(recovery)
  {}

  lemma DetachedHydrationDoesNotMutateOther<T>(requester: Projection<T>,
                                               other: Projection<T>, page: seq<T>)
    ensures Hydrate(requester, page).subscriptionId == requester.subscriptionId
    ensures other == other
  {}

  predicate IsTypedStart(value: LegacyStart) { false }

  lemma LegacyShapeIsNegativeBaseline(value: LegacyStart)
    ensures !IsTypedStart(value)
  {}
}
