// Exact typed Wire.Json codec for the Canvas instance channel at pinned
// upstream PR #298 head 8251dd333c9758e5cbe9e26d2972f9d2d722477d.
// providerId is transported losslessly but remains provisional snapshot data;
// the codec proves representation fidelity, not issuer/routing semantics.
include "../canvas.dfy"
include "../fidelity.dfy"

module CanvasCodec {
  export
    provides DecodeCanvasAction, encodeCanvasState, decodeCanvasState
    provides CanvasActionRoundTrip, CanvasStateRoundTrip
    provides Wire, Canvas, AhpSkeleton
    reveals EncodeCanvasAction, OptionalStringField,
            OptionalAvailabilityField, EncodeAvailability

  import opened Canvas
  import opened AhpSkeleton
  import Wire = ConfluxCodec
  import F = Fidelity

  function DecodeAvailability(o: Option<Wire.Json>): CanvasAvailability {
    if F.asStr(o) == "stale" then Stale else Ready
  }

  function EncodeAvailability(a: CanvasAvailability): Wire.Json {
    match a case Ready => Wire.JStr("ready") case Stale => Wire.JStr("stale")
  }

  function OptStr(o: Option<Wire.Json>): Option<string> {
    if o.Some? && o.value.JStr? then Some(o.value.s) else None
  }

  function OptObj(o: Option<Wire.Json>): Option<map<string, Wire.Json>> {
    if o.Some? && o.value.JObj? then Some(o.value.fields) else None
  }

  function OptionalStringField(name: string, value: AhpSkeleton.Option<string>): map<string, Wire.Json> {
    if value.Some? then map[name := Wire.JStr(value.value)] else map[]
  }

  function OptionalAvailabilityField(value: AhpSkeleton.Option<CanvasAvailability>): map<string, Wire.Json> {
    if value.Some? then map["availability" := EncodeAvailability(value.value)] else map[]
  }

  function DecodeCanvasAction(j: Wire.Json): CanvasAction {
    var t := F.asStr(F.field(j, "type"));
    if t == "canvas/updated" then
      Updated(OptStr(F.field(j, "title")), OptStr(F.field(j, "activity")),
              OptStr(F.field(j, "contentUri")),
              if F.field(j, "availability").Some? && F.field(j, "availability").value.JStr?
              then Some(DecodeAvailability(F.field(j, "availability"))) else None)
    else if t == "canvas/closeRequested" then CloseRequested
    else CanvasUnknown(j)
  }

  function EncodeCanvasAction(a: CanvasAction): Wire.Json {
    match a
    case Updated(t, ac, uri, av) => Wire.JObj(
      map["type" := Wire.JStr("canvas/updated")] +
      OptionalStringField("title", t) +
      OptionalStringField("activity", ac) +
      OptionalStringField("contentUri", uri) +
      OptionalAvailabilityField(av))
    case CloseRequested => Wire.JObj(map["type" := Wire.JStr("canvas/closeRequested")])
    case CanvasUnknown(raw) => raw
  }

  lemma CanvasActionRoundTrip(a: CanvasAction)
    requires !a.CanvasUnknown?
    ensures DecodeCanvasAction(EncodeCanvasAction(a)) == a
  {
    match a
    case Updated(_, _, _, _) =>
    case CloseRequested =>
    case CanvasUnknown(_) =>
  }

  // C2-shaped canonical round-trip on the image of Encode: re-decoding an
  // encoded typed action and re-encoding yields the same wire. Follows directly
  // from CanvasActionRoundTrip (decode ∘ encode == id on the typed variants).
  lemma CanvasWireCanonicalRoundTrip(a: CanvasAction)
    requires !a.CanvasUnknown?
    ensures EncodeCanvasAction(DecodeCanvasAction(EncodeCanvasAction(a))) == EncodeCanvasAction(a)
  { CanvasActionRoundTrip(a); }

  lemma CanvasActionRoundTrip_NonVacuous()
    ensures exists a: CanvasAction :: (!a.CanvasUnknown? &&
      DecodeCanvasAction(EncodeCanvasAction(a)) == a)
  {
    var a := Updated(Some("Published"), None, Some("ahp-session:/content/1"), Some(Ready));
    CanvasActionRoundTrip(a);
  }

  lemma CanvasUnknown_EncodeVerbatim(j: Wire.Json)
    ensures EncodeCanvasAction(CanvasUnknown(j)) == j
  {}

  function encodeCanvasState(s: CanvasState): Wire.Json {
    Wire.JObj(
      map[
        "canvasId" := Wire.JStr(s.canvasId),
        "providerId" := Wire.JStr(s.providerId),
        "availability" := EncodeAvailability(s.availability)] +
      (if s.input.Some? then map["input" := Wire.JObj(s.input.value)] else map[]) +
      OptionalStringField("title", s.title) +
      OptionalStringField("activity", s.activity) +
      OptionalStringField("contentUri", s.contentUri))
  }

  function decodeCanvasState(j: Wire.Json): CanvasState {
    CanvasState(
      F.asStr(F.field(j, "canvasId")),
      F.asStr(F.field(j, "providerId")),
      OptObj(F.field(j, "input")),
      OptStr(F.field(j, "title")),
      OptStr(F.field(j, "activity")),
      OptStr(F.field(j, "contentUri")),
      DecodeAvailability(F.field(j, "availability")))
  }

  lemma CanvasStateRoundTrip(s: CanvasState)
    ensures decodeCanvasState(encodeCanvasState(s)) == s
  {}

  lemma CanvasStateRoundTrip_NonVacuous()
    ensures exists s: CanvasState :: decodeCanvasState(encodeCanvasState(s)) == s
  {
    var s := CanvasState("markdown", "opaque-provider",
      Some(map["path" := Wire.JStr("README.md")]), Some("README.md"),
      Some("editing"), Some("ahp-session:/content/1"), Ready);
    CanvasStateRoundTrip(s);
  }

  ghost predicate CanvasCodecWf(a: CanvasAction) { !a.CanvasUnknown? }
  function CanvasCodecC(): Wire.Codec<CanvasAction> {
    Wire.Codec(EncodeCanvasAction, (j: Wire.Json) => Wire.Some(DecodeCanvasAction(j)))
  }
  lemma CanvasCodecRoundTripsWhere()
    ensures forall a :: CanvasCodecWf(a) ==>
      CanvasCodecC().decode(CanvasCodecC().encode(a)) == Wire.Some(a)
  {
    forall a | CanvasCodecWf(a)
      ensures CanvasCodecC().decode(CanvasCodecC().encode(a)) == Wire.Some(a)
    { CanvasActionRoundTrip(a); }
  }
}
