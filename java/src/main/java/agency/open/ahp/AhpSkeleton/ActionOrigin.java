// Class ActionOrigin
// Dafny class ActionOrigin compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class ActionOrigin {
  public dafny.DafnySequence<? extends dafny.CodePoint> _clientId;
  public java.math.BigInteger _clientSeq;
  public ActionOrigin (dafny.DafnySequence<? extends dafny.CodePoint> clientId, java.math.BigInteger clientSeq) {
    this._clientId = clientId;
    this._clientSeq = clientSeq;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ActionOrigin o = (ActionOrigin)other;
    return true && java.util.Objects.equals(this._clientId, o._clientId) && java.util.Objects.equals(this._clientSeq, o._clientSeq);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._clientId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._clientSeq);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.ActionOrigin.ActionOrigin");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._clientId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._clientSeq));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<ActionOrigin> _TYPE = dafny.TypeDescriptor.<ActionOrigin>referenceWithInitializer(ActionOrigin.class, () -> ActionOrigin.Default());
  public static dafny.TypeDescriptor<ActionOrigin> _typeDescriptor() {
    return (dafny.TypeDescriptor<ActionOrigin>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ActionOrigin theDefault = ActionOrigin.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), java.math.BigInteger.ZERO);
  public static ActionOrigin Default() {
    return theDefault;
  }
  public static ActionOrigin create(dafny.DafnySequence<? extends dafny.CodePoint> clientId, java.math.BigInteger clientSeq) {
    return new ActionOrigin(clientId, clientSeq);
  }
  public static ActionOrigin create_ActionOrigin(dafny.DafnySequence<? extends dafny.CodePoint> clientId, java.math.BigInteger clientSeq) {
    return create(clientId, clientSeq);
  }
  public boolean is_ActionOrigin() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_clientId() {
    return this._clientId;
  }
  public java.math.BigInteger dtor_clientSeq() {
    return this._clientSeq;
  }
}
