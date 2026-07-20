// Class CanvasState
// Dafny class CanvasState compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasState {
  public dafny.DafnySequence<? extends dafny.CodePoint> _canvasId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _providerId;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>> _input;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _title;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _activity;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _contentUri;
  public CanvasAvailability _availability;
  public CanvasState (dafny.DafnySequence<? extends dafny.CodePoint> canvasId, dafny.DafnySequence<? extends dafny.CodePoint> providerId, agency.open.ahp.AhpSkeleton.Option<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>> input, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> title, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> contentUri, CanvasAvailability availability) {
    this._canvasId = canvasId;
    this._providerId = providerId;
    this._input = input;
    this._title = title;
    this._activity = activity;
    this._contentUri = contentUri;
    this._availability = availability;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasState o = (CanvasState)other;
    return true && java.util.Objects.equals(this._canvasId, o._canvasId) && java.util.Objects.equals(this._providerId, o._providerId) && java.util.Objects.equals(this._input, o._input) && java.util.Objects.equals(this._title, o._title) && java.util.Objects.equals(this._activity, o._activity) && java.util.Objects.equals(this._contentUri, o._contentUri) && java.util.Objects.equals(this._availability, o._availability);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._canvasId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._providerId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._input);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._title);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activity);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._contentUri);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._availability);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.CanvasState.CanvasState");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._canvasId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._providerId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._input));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._title));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._activity));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._contentUri));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._availability));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<CanvasState> _TYPE = dafny.TypeDescriptor.<CanvasState>referenceWithInitializer(CanvasState.class, () -> CanvasState.Default());
  public static dafny.TypeDescriptor<CanvasState> _typeDescriptor() {
    return (dafny.TypeDescriptor<CanvasState>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final CanvasState theDefault = CanvasState.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>>Default(dafny.DafnyMap.<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), CanvasAvailability.Default());
  public static CanvasState Default() {
    return theDefault;
  }
  public static CanvasState create(dafny.DafnySequence<? extends dafny.CodePoint> canvasId, dafny.DafnySequence<? extends dafny.CodePoint> providerId, agency.open.ahp.AhpSkeleton.Option<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>> input, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> title, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> contentUri, CanvasAvailability availability) {
    return new CanvasState(canvasId, providerId, input, title, activity, contentUri, availability);
  }
  public static CanvasState create_CanvasState(dafny.DafnySequence<? extends dafny.CodePoint> canvasId, dafny.DafnySequence<? extends dafny.CodePoint> providerId, agency.open.ahp.AhpSkeleton.Option<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>> input, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> title, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> contentUri, CanvasAvailability availability) {
    return create(canvasId, providerId, input, title, activity, contentUri, availability);
  }
  public boolean is_CanvasState() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_canvasId() {
    return this._canvasId;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_providerId() {
    return this._providerId;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>> dtor_input() {
    return this._input;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_title() {
    return this._title;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_activity() {
    return this._activity;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_contentUri() {
    return this._contentUri;
  }
  public CanvasAvailability dtor_availability() {
    return this._availability;
  }
}
