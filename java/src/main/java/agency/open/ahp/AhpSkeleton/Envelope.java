// Class Envelope<A>
// Dafny class Envelope<A> compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class Envelope<A> {
  public dafny.DafnySequence<? extends dafny.CodePoint> _channel;
  public A _action;
  public java.math.BigInteger _serverSeq;
  public Option<ActionOrigin> _origin;
  public Option<dafny.DafnySequence<? extends dafny.CodePoint>> _rejectionReason;
  protected dafny.TypeDescriptor<A> _td_A;
  public Envelope (dafny.TypeDescriptor<A> _td_A, dafny.DafnySequence<? extends dafny.CodePoint> channel, A action, java.math.BigInteger serverSeq, Option<ActionOrigin> origin, Option<dafny.DafnySequence<? extends dafny.CodePoint>> rejectionReason) {
    this._td_A = _td_A;
    this._channel = channel;
    this._action = action;
    this._serverSeq = serverSeq;
    this._origin = origin;
    this._rejectionReason = rejectionReason;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Envelope<A> o = (Envelope<A>)other;
    return true && java.util.Objects.equals(this._channel, o._channel) && java.util.Objects.equals(this._action, o._action) && java.util.Objects.equals(this._serverSeq, o._serverSeq) && java.util.Objects.equals(this._origin, o._origin) && java.util.Objects.equals(this._rejectionReason, o._rejectionReason);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._channel);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._action);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._serverSeq);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._origin);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._rejectionReason);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.Envelope.Envelope");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._channel));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._action));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._serverSeq));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._origin));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._rejectionReason));
    s.append(")");
    return s.toString();
  }
  public static <A> dafny.TypeDescriptor<Envelope<A>> _typeDescriptor(dafny.TypeDescriptor<A> _td_A) {
    return (dafny.TypeDescriptor<Envelope<A>>) (dafny.TypeDescriptor<?>) dafny.TypeDescriptor.<Envelope<A>>referenceWithInitializer(Envelope.class, () -> Envelope.<A>Default(_td_A, _td_A.defaultValue()));
  }

  public static <A> Envelope<A> Default(dafny.TypeDescriptor<A> _td_A, A _default_A) {
    return Envelope.<A>create(_td_A, dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), _default_A, java.math.BigInteger.ZERO, Option.<ActionOrigin>Default(ActionOrigin._typeDescriptor()), Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)));
  }
  public static <A> Envelope<A> create(dafny.TypeDescriptor<A> _td_A, dafny.DafnySequence<? extends dafny.CodePoint> channel, A action, java.math.BigInteger serverSeq, Option<ActionOrigin> origin, Option<dafny.DafnySequence<? extends dafny.CodePoint>> rejectionReason) {
    return new Envelope<A>(_td_A, channel, action, serverSeq, origin, rejectionReason);
  }
  public static <A> Envelope<A> create_Envelope(dafny.TypeDescriptor<A> _td_A, dafny.DafnySequence<? extends dafny.CodePoint> channel, A action, java.math.BigInteger serverSeq, Option<ActionOrigin> origin, Option<dafny.DafnySequence<? extends dafny.CodePoint>> rejectionReason) {
    return create(_td_A, channel, action, serverSeq, origin, rejectionReason);
  }
  public boolean is_Envelope() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_channel() {
    return this._channel;
  }
  public A dtor_action() {
    return this._action;
  }
  public java.math.BigInteger dtor_serverSeq() {
    return this._serverSeq;
  }
  public Option<ActionOrigin> dtor_origin() {
    return this._origin;
  }
  public Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_rejectionReason() {
    return this._rejectionReason;
  }
}
