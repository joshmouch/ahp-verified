// Class Chat
// Dafny class Chat compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class Chat {
  public dafny.DafnySequence<? extends dafny.CodePoint> _resource;
  public dafny.DafnySequence<? extends dafny.CodePoint> _title;
  public java.math.BigInteger _status;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _activity;
  public dafny.DafnySequence<? extends dafny.CodePoint> _modifiedAt;
  public agency.open.ahp.ConfluxCodec.Json _origin;
  public Chat (dafny.DafnySequence<? extends dafny.CodePoint> resource, dafny.DafnySequence<? extends dafny.CodePoint> title, java.math.BigInteger status, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, dafny.DafnySequence<? extends dafny.CodePoint> modifiedAt, agency.open.ahp.ConfluxCodec.Json origin) {
    this._resource = resource;
    this._title = title;
    this._status = status;
    this._activity = activity;
    this._modifiedAt = modifiedAt;
    this._origin = origin;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Chat o = (Chat)other;
    return true && java.util.Objects.equals(this._resource, o._resource) && java.util.Objects.equals(this._title, o._title) && java.util.Objects.equals(this._status, o._status) && java.util.Objects.equals(this._activity, o._activity) && java.util.Objects.equals(this._modifiedAt, o._modifiedAt) && java.util.Objects.equals(this._origin, o._origin);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._resource);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._title);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._status);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activity);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._modifiedAt);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._origin);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.Chat.Chat");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._resource));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._title));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._status));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._activity));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._modifiedAt));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._origin));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<Chat> _TYPE = dafny.TypeDescriptor.<Chat>referenceWithInitializer(Chat.class, () -> Chat.Default());
  public static dafny.TypeDescriptor<Chat> _typeDescriptor() {
    return (dafny.TypeDescriptor<Chat>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final Chat theDefault = Chat.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), java.math.BigInteger.ZERO, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.ConfluxCodec.Json.Default());
  public static Chat Default() {
    return theDefault;
  }
  public static Chat create(dafny.DafnySequence<? extends dafny.CodePoint> resource, dafny.DafnySequence<? extends dafny.CodePoint> title, java.math.BigInteger status, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, dafny.DafnySequence<? extends dafny.CodePoint> modifiedAt, agency.open.ahp.ConfluxCodec.Json origin) {
    return new Chat(resource, title, status, activity, modifiedAt, origin);
  }
  public static Chat create_Chat(dafny.DafnySequence<? extends dafny.CodePoint> resource, dafny.DafnySequence<? extends dafny.CodePoint> title, java.math.BigInteger status, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, dafny.DafnySequence<? extends dafny.CodePoint> modifiedAt, agency.open.ahp.ConfluxCodec.Json origin) {
    return create(resource, title, status, activity, modifiedAt, origin);
  }
  public boolean is_Chat() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_resource() {
    return this._resource;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_title() {
    return this._title;
  }
  public java.math.BigInteger dtor_status() {
    return this._status;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_activity() {
    return this._activity;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_modifiedAt() {
    return this._modifiedAt;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_origin() {
    return this._origin;
  }
}
