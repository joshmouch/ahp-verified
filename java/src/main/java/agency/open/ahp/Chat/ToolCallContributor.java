// Class ToolCallContributor
// Dafny class ToolCallContributor compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ToolCallContributor {
  public dafny.DafnySequence<? extends dafny.CodePoint> _kind;
  public agency.open.ahp.ConfluxCodec.Json _raw;
  public ToolCallContributor (dafny.DafnySequence<? extends dafny.CodePoint> kind, agency.open.ahp.ConfluxCodec.Json raw) {
    this._kind = kind;
    this._raw = raw;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ToolCallContributor o = (ToolCallContributor)other;
    return true && java.util.Objects.equals(this._kind, o._kind) && java.util.Objects.equals(this._raw, o._raw);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._kind);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._raw);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ToolCallContributor.ToolCallContributor");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._kind));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._raw));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<ToolCallContributor> _TYPE = dafny.TypeDescriptor.<ToolCallContributor>referenceWithInitializer(ToolCallContributor.class, () -> ToolCallContributor.Default());
  public static dafny.TypeDescriptor<ToolCallContributor> _typeDescriptor() {
    return (dafny.TypeDescriptor<ToolCallContributor>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ToolCallContributor theDefault = ToolCallContributor.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.ConfluxCodec.Json.Default());
  public static ToolCallContributor Default() {
    return theDefault;
  }
  public static ToolCallContributor create(dafny.DafnySequence<? extends dafny.CodePoint> kind, agency.open.ahp.ConfluxCodec.Json raw) {
    return new ToolCallContributor(kind, raw);
  }
  public static ToolCallContributor create_ToolCallContributor(dafny.DafnySequence<? extends dafny.CodePoint> kind, agency.open.ahp.ConfluxCodec.Json raw) {
    return create(kind, raw);
  }
  public boolean is_ToolCallContributor() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_kind() {
    return this._kind;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    return this._raw;
  }
}
