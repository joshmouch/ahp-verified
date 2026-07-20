// Class CanvasProviderErrorData
// Dafny class CanvasProviderErrorData compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasProviderErrorData {
  public dafny.DafnySequence<? extends dafny.CodePoint> _code;
  public dafny.DafnySequence<? extends dafny.CodePoint> _message;
  public CanvasProviderErrorData (dafny.DafnySequence<? extends dafny.CodePoint> code, dafny.DafnySequence<? extends dafny.CodePoint> message) {
    this._code = code;
    this._message = message;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasProviderErrorData o = (CanvasProviderErrorData)other;
    return true && java.util.Objects.equals(this._code, o._code) && java.util.Objects.equals(this._message, o._message);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._code);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._message);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.CanvasProviderErrorData.CanvasProviderErrorData");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._code));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._message));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<CanvasProviderErrorData> _TYPE = dafny.TypeDescriptor.<CanvasProviderErrorData>referenceWithInitializer(CanvasProviderErrorData.class, () -> CanvasProviderErrorData.Default());
  public static dafny.TypeDescriptor<CanvasProviderErrorData> _typeDescriptor() {
    return (dafny.TypeDescriptor<CanvasProviderErrorData>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final CanvasProviderErrorData theDefault = CanvasProviderErrorData.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR));
  public static CanvasProviderErrorData Default() {
    return theDefault;
  }
  public static CanvasProviderErrorData create(dafny.DafnySequence<? extends dafny.CodePoint> code, dafny.DafnySequence<? extends dafny.CodePoint> message) {
    return new CanvasProviderErrorData(code, message);
  }
  public static CanvasProviderErrorData create_CanvasProviderErrorData(dafny.DafnySequence<? extends dafny.CodePoint> code, dafny.DafnySequence<? extends dafny.CodePoint> message) {
    return create(code, message);
  }
  public boolean is_CanvasProviderErrorData() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_code() {
    return this._code;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_message() {
    return this._message;
  }
}
