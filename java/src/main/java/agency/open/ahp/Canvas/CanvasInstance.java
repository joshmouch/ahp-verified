// Class CanvasInstance
// Dafny class CanvasInstance compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasInstance {
  public dafny.DafnySequence<? extends dafny.CodePoint> _channel;
  public CanvasState _snapshot;
  public CanvasInstance (dafny.DafnySequence<? extends dafny.CodePoint> channel, CanvasState snapshot) {
    this._channel = channel;
    this._snapshot = snapshot;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasInstance o = (CanvasInstance)other;
    return true && java.util.Objects.equals(this._channel, o._channel) && java.util.Objects.equals(this._snapshot, o._snapshot);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._channel);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._snapshot);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.CanvasInstance.CanvasInstance");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._channel));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._snapshot));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<CanvasInstance> _TYPE = dafny.TypeDescriptor.<CanvasInstance>referenceWithInitializer(CanvasInstance.class, () -> CanvasInstance.Default());
  public static dafny.TypeDescriptor<CanvasInstance> _typeDescriptor() {
    return (dafny.TypeDescriptor<CanvasInstance>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final CanvasInstance theDefault = CanvasInstance.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), CanvasState.Default());
  public static CanvasInstance Default() {
    return theDefault;
  }
  public static CanvasInstance create(dafny.DafnySequence<? extends dafny.CodePoint> channel, CanvasState snapshot) {
    return new CanvasInstance(channel, snapshot);
  }
  public static CanvasInstance create_CanvasInstance(dafny.DafnySequence<? extends dafny.CodePoint> channel, CanvasState snapshot) {
    return create(channel, snapshot);
  }
  public boolean is_CanvasInstance() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_channel() {
    return this._channel;
  }
  public CanvasState dtor_snapshot() {
    return this._snapshot;
  }
}
