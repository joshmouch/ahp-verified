// Class SessionAction_ChatUpdated
// Dafny class SessionAction_ChatUpdated compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_ChatUpdated extends SessionAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _resource;
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _status;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _activity;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _modifiedAt;
  public SessionAction_ChatUpdated (dafny.DafnySequence<? extends dafny.CodePoint> resource, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> status, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> modifiedAt) {
    super();
    this._resource = resource;
    this._status = status;
    this._activity = activity;
    this._modifiedAt = modifiedAt;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_ChatUpdated o = (SessionAction_ChatUpdated)other;
    return true && java.util.Objects.equals(this._resource, o._resource) && java.util.Objects.equals(this._status, o._status) && java.util.Objects.equals(this._activity, o._activity) && java.util.Objects.equals(this._modifiedAt, o._modifiedAt);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 15;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._resource);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._status);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activity);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._modifiedAt);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.ChatUpdated");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._resource));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._status));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._activity));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._modifiedAt));
    s.append(")");
    return s.toString();
  }
}
