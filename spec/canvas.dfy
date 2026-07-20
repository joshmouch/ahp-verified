// AHP Canvas channel imported from the immutable PR #298 snapshot
// (8251dd333c9758e5cbe9e26d2972f9d2d722477d).
//
// The wire fields are mirrored exactly, but the upstream-contested providerId,
// declaration catalogue and operation/discovery interpretation is deliberately
// NOT promoted into a durable architectural theorem here.  providerId is stored
// and replayed as an opaque string.  Durable laws cover only the URI-addressed
// instance channel, sparse state updates, lifecycle signals, contentUri reuse of
// the general resourceRead command, support/authorization separation, collision
// distinctions, deterministic degradation and reconnect preservation.
include "ahp_skeleton.dfy"

module Canvas {
  export
    provides ApplyToCanvas, RunCorpus
    provides AhpSkeleton, Wire, CC
    reveals CanvasAvailability, CanvasState, CanvasAction, CanvasInstance,
            CanvasAdmission, CanvasProviderErrorData, CanvasError,
            ResourceReadParams, apply1, fold, C0,
            SameInstance, ValidateCanvasAdmission, DegradeAvailability,
            ReconnectCanvas, ContentResourceRead

  import opened AhpSkeleton
  import Wire = ConfluxCodec
  import CC = ConfluxContract

  datatype CanvasAvailability = Ready | Stale

  datatype CanvasState = CanvasState(
    canvasId: string,
    // Provisional upstream-snapshot data.  Preserved opaquely; no durable law
    // assigns issuer, routing, namespace or reconnect semantics to this value.
    providerId: string,
    input: Option<map<string, Wire.Json>>,
    title: Option<string>,
    activity: Option<string>,
    contentUri: Option<string>,
    availability: CanvasAvailability)

  datatype CanvasAction =
    | Updated(title: Option<string>, activity: Option<string>,
              contentUri: Option<string>, availability: Option<CanvasAvailability>)
    | CloseRequested
    | CanvasUnknown(raw: Wire.Json)

  function ApplyToCanvas(s: CanvasState, a: CanvasAction, now: int): (CanvasState, ReduceOutcome)
  {
    match a
    case Updated(t, ac, uri, av) =>
      (s.(title := if t.Some? then t else s.title,
          activity := if ac.Some? then ac else s.activity,
          contentUri := if uri.Some? then uri else s.contentUri,
          availability := if av.Some? then av.value else s.availability), Applied)
    case CloseRequested => (s, NoOp)
    case CanvasUnknown(_) => (s, NoOp)
  }

  function apply1(s: CanvasState, a: CanvasAction): CanvasState
  { ApplyToCanvas(s, a, 9999).0 }

  function fold(s: CanvasState, actions: seq<CanvasAction>): CanvasState
  { CC.Fold(apply1, s, actions) }

  function C0(): CanvasState
  { CanvasState("markdown", "opaque-provider", None, None, None, None, Ready) }

  // Stable instance identity is the channel URI alone.  CanvasState deliberately
  // has no instanceId; changes to any snapshot field cannot change identity.
  datatype CanvasInstance = CanvasInstance(channel: string, snapshot: CanvasState)
  predicate SameInstance(a: CanvasInstance, b: CanvasInstance)
  { a.channel == b.channel }

  lemma UriIsSoleInstanceIdentity(a: CanvasInstance, b: CanvasInstance)
    ensures SameInstance(a, b) <==> a.channel == b.channel
  {}

  lemma SnapshotCannotChangeInstanceIdentity(channel: string, before: CanvasState, after: CanvasState)
    ensures SameInstance(CanvasInstance(channel, before), CanvasInstance(channel, after))
  {}

  // The Canvas content address reuses the existing general resourceRead command.
  // There is intentionally no Canvas-specific read request in this model.
  datatype ResourceReadParams = ResourceReadParams(channel: string, uri: string)
  function ContentResourceRead(contentUri: string): ResourceReadParams
  { ResourceReadParams("ahp-root://", contentUri) }
  lemma ContentUriUsesGeneralResourceRead(contentUri: string)
    ensures ContentResourceRead(contentUri).uri == contentUri
    ensures ContentResourceRead(contentUri).channel == "ahp-root://"
  {}

  // Support, local handler installation, authorization and URI collision are
  // distinct decisions.  The order is deterministic and does not turn support
  // negotiation into authority.
  datatype CanvasAdmission =
    | CanvasSupported
    | CanvasUnsupported
    | CanvasHandlerUnavailable
    | CanvasUnauthorized
    | CanvasUriCollision

  function ValidateCanvasAdmission(
    supportNegotiated: bool,
    handlerInstalled: bool,
    authorized: bool,
    channelInUse: bool): CanvasAdmission
  {
    if !supportNegotiated then CanvasUnsupported
    else if !handlerInstalled then CanvasHandlerUnavailable
    else if !authorized then CanvasUnauthorized
    else if channelInUse then CanvasUriCollision
    else CanvasSupported
  }

  lemma SupportDoesNotGrantAuthority(handlerInstalled: bool, channelInUse: bool)
    ensures ValidateCanvasAdmission(true, handlerInstalled, false, channelInUse) != CanvasSupported
  {}

  lemma CollisionIsDistinctFromProviderFailure()
    ensures CanvasUriCollision != CanvasHandlerUnavailable
    ensures CanvasUriCollision != CanvasUnauthorized
  {}

  // Exact PR #298 provider error payload/code.  Admission failures above stay
  // distinct from a provider that was admitted and then returned -32012.
  const CANVAS_PROVIDER_ERROR: int := -32012
  datatype CanvasProviderErrorData = CanvasProviderErrorData(code: string, message: string)
  datatype CanvasError =
    | AdmissionError(reason: CanvasAdmission)
    | ProviderError(errorCode: int, data: CanvasProviderErrorData)

  function CanvasProviderFailure(code: string, message: string): CanvasError
  { ProviderError(CANVAS_PROVIDER_ERROR, CanvasProviderErrorData(code, message)) }

  lemma ProviderFailureHasExactWireCode(code: string, message: string)
    ensures CanvasProviderFailure(code, message).ProviderError?
    ensures CanvasProviderFailure(code, message).errorCode == -32012
  {}

  // Disconnect degradation changes availability only.  Reconnect restores the
  // exact retained snapshot; it does not reinterpret provisional provider data.
  function DegradeAvailability(s: CanvasState): CanvasState
  { s.(availability := Stale) }

  lemma DegradationPreservesDurableSnapshot(s: CanvasState)
    ensures DegradeAvailability(s).canvasId == s.canvasId
    ensures DegradeAvailability(s).providerId == s.providerId
    ensures DegradeAvailability(s).input == s.input
    ensures DegradeAvailability(s).title == s.title
    ensures DegradeAvailability(s).activity == s.activity
    ensures DegradeAvailability(s).contentUri == s.contentUri
    ensures DegradeAvailability(s).availability == Stale
  {}

  function ReconnectCanvas(snapshot: CanvasState, supportNegotiated: bool): Option<CanvasState>
  { if supportNegotiated then Some(snapshot) else None }

  lemma ReconnectPreservesSnapshot(snapshot: CanvasState)
    ensures ReconnectCanvas(snapshot, true) == Some(snapshot)
  {}

  lemma UpdatedPreservesAbsentFields(
    s: CanvasState, t: Option<string>, ac: Option<string>,
    uri: Option<string>, av: Option<CanvasAvailability>)
    ensures var next := apply1(s, Updated(t, ac, uri, av));
      next.canvasId == s.canvasId && next.providerId == s.providerId && next.input == s.input &&
      (!t.Some? ==> next.title == s.title) &&
      (!ac.Some? ==> next.activity == s.activity) &&
      (!uri.Some? ==> next.contentUri == s.contentUri) &&
      (!av.Some? ==> next.availability == s.availability)
  {}

  lemma CloseRequestedIsNoOp(s: CanvasState)
    ensures apply1(s, CloseRequested) == s
  {}

  lemma UnknownCanvasActionIsNoOp(s: CanvasState, raw: Wire.Json)
    ensures apply1(s, CanvasUnknown(raw)) == s
  {}

  method RunCorpus() returns (pass: nat, total: nat) {
    total := 5; pass := 0;
    var s := CanvasState("markdown", "acme.editors", None,
      Some("Draft"), Some("idle"), Some("ahp-session:/2f9c/content/canvas-1"), Ready);
    var all := apply1(s, Updated(Some("Published"), Some("saved"),
      Some("https://example.com/docs/published.html"), Some(Stale)));
    if all == s.(title := Some("Published"), activity := Some("saved"),
                 contentUri := Some("https://example.com/docs/published.html"), availability := Stale) { pass := pass + 1; }
    var left := apply1(s, Updated(Some("Renamed"), None,
      Some("https://example.com/docs/renamed.html"), None));
    if left == s.(title := Some("Renamed"), contentUri := Some("https://example.com/docs/renamed.html")) { pass := pass + 1; }
    var right := apply1(s, Updated(None, Some("error"), None, Some(Stale)));
    if right == s.(activity := Some("error"), availability := Stale) { pass := pass + 1; }
    if apply1(s, CloseRequested) == s { pass := pass + 1; }
    if apply1(s, CanvasUnknown(Wire.JObj(map["type" := Wire.JStr("canvas/nonExistentAction")]))) == s { pass := pass + 1; }
  }
}
