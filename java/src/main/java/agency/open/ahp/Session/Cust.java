// Class Cust
// Dafny class Cust compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class Cust {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public dafny.DafnySequence<? extends dafny.CodePoint> _kind;
  public dafny.DafnySequence<? extends dafny.CodePoint> _uri;
  public dafny.DafnySequence<? extends dafny.CodePoint> _name;
  public agency.open.ahp.AhpSkeleton.Option<Boolean> _enabled;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _state;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _channel;
  public Cust (dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> kind, dafny.DafnySequence<? extends dafny.CodePoint> uri, dafny.DafnySequence<? extends dafny.CodePoint> name, agency.open.ahp.AhpSkeleton.Option<Boolean> enabled, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> state, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> channel) {
    this._id = id;
    this._kind = kind;
    this._uri = uri;
    this._name = name;
    this._enabled = enabled;
    this._state = state;
    this._channel = channel;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Cust o = (Cust)other;
    return true && java.util.Objects.equals(this._id, o._id) && java.util.Objects.equals(this._kind, o._kind) && java.util.Objects.equals(this._uri, o._uri) && java.util.Objects.equals(this._name, o._name) && java.util.Objects.equals(this._enabled, o._enabled) && java.util.Objects.equals(this._state, o._state) && java.util.Objects.equals(this._channel, o._channel);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._kind);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._uri);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._name);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._enabled);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._state);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._channel);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.Cust.Cust");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._kind));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._uri));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._name));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._enabled));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._state));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._channel));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<Cust> _TYPE = dafny.TypeDescriptor.<Cust>referenceWithInitializer(Cust.class, () -> Cust.Default());
  public static dafny.TypeDescriptor<Cust> _typeDescriptor() {
    return (dafny.TypeDescriptor<Cust>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final Cust theDefault = Cust.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<Boolean>Default(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
  public static Cust Default() {
    return theDefault;
  }
  public static Cust create(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> kind, dafny.DafnySequence<? extends dafny.CodePoint> uri, dafny.DafnySequence<? extends dafny.CodePoint> name, agency.open.ahp.AhpSkeleton.Option<Boolean> enabled, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> state, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> channel) {
    return new Cust(id, kind, uri, name, enabled, state, channel);
  }
  public static Cust create_Cust(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> kind, dafny.DafnySequence<? extends dafny.CodePoint> uri, dafny.DafnySequence<? extends dafny.CodePoint> name, agency.open.ahp.AhpSkeleton.Option<Boolean> enabled, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> state, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> channel) {
    return create(id, kind, uri, name, enabled, state, channel);
  }
  public boolean is_Cust() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_id() {
    return this._id;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_kind() {
    return this._kind;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_uri() {
    return this._uri;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_name() {
    return this._name;
  }
  public agency.open.ahp.AhpSkeleton.Option<Boolean> dtor_enabled() {
    return this._enabled;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_state() {
    return this._state;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_channel() {
    return this._channel;
  }
}
