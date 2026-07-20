// Class CanvasAction
// Dafny class CanvasAction compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class CanvasAction {
  public CanvasAction() {
  }
  private static final dafny.TypeDescriptor<CanvasAction> _TYPE = dafny.TypeDescriptor.<CanvasAction>referenceWithInitializer(CanvasAction.class, () -> CanvasAction.Default());
  public static dafny.TypeDescriptor<CanvasAction> _typeDescriptor() {
    return (dafny.TypeDescriptor<CanvasAction>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final CanvasAction theDefault = CanvasAction.create_Updated(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<CanvasAvailability>Default(CanvasAvailability._typeDescriptor()));
  public static CanvasAction Default() {
    return theDefault;
  }
  public static CanvasAction create_Updated(agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> title, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> contentUri, agency.open.ahp.AhpSkeleton.Option<CanvasAvailability> availability) {
    return new CanvasAction_Updated(title, activity, contentUri, availability);
  }
  public static CanvasAction create_CloseRequested() {
    return new CanvasAction_CloseRequested();
  }
  public static CanvasAction create_CanvasUnknown(agency.open.ahp.ConfluxCodec.Json raw) {
    return new CanvasAction_CanvasUnknown(raw);
  }
  public boolean is_Updated() { return this instanceof CanvasAction_Updated; }
  public boolean is_CloseRequested() { return this instanceof CanvasAction_CloseRequested; }
  public boolean is_CanvasUnknown() { return this instanceof CanvasAction_CanvasUnknown; }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_title() {
    CanvasAction d = this;
    return ((CanvasAction_Updated)d)._title;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_activity() {
    CanvasAction d = this;
    return ((CanvasAction_Updated)d)._activity;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_contentUri() {
    CanvasAction d = this;
    return ((CanvasAction_Updated)d)._contentUri;
  }
  public agency.open.ahp.AhpSkeleton.Option<CanvasAvailability> dtor_availability() {
    CanvasAction d = this;
    return ((CanvasAction_Updated)d)._availability;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    CanvasAction d = this;
    return ((CanvasAction_CanvasUnknown)d)._raw;
  }
}
