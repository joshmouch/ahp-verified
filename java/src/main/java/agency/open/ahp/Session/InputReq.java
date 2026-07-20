// Class InputReq
// Dafny class InputReq compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class InputReq {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public agency.open.ahp.ConfluxCodec.Json _raw;
  public InputReq (dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.ConfluxCodec.Json raw) {
    this._id = id;
    this._raw = raw;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    InputReq o = (InputReq)other;
    return true && java.util.Objects.equals(this._id, o._id) && java.util.Objects.equals(this._raw, o._raw);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._raw);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.InputReq.InputReq");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._raw));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<InputReq> _TYPE = dafny.TypeDescriptor.<InputReq>referenceWithInitializer(InputReq.class, () -> InputReq.Default());
  public static dafny.TypeDescriptor<InputReq> _typeDescriptor() {
    return (dafny.TypeDescriptor<InputReq>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final InputReq theDefault = InputReq.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.ConfluxCodec.Json.Default());
  public static InputReq Default() {
    return theDefault;
  }
  public static InputReq create(dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.ConfluxCodec.Json raw) {
    return new InputReq(id, raw);
  }
  public static InputReq create_InputReq(dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.ConfluxCodec.Json raw) {
    return create(id, raw);
  }
  public boolean is_InputReq() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_id() {
    return this._id;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    return this._raw;
  }
}
