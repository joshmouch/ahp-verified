// Class SessionAction_ActiveClientRemoved
// Dafny class SessionAction_ActiveClientRemoved compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_ActiveClientRemoved extends SessionAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _clientId;
  public SessionAction_ActiveClientRemoved (dafny.DafnySequence<? extends dafny.CodePoint> clientId) {
    super();
    this._clientId = clientId;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_ActiveClientRemoved o = (SessionAction_ActiveClientRemoved)other;
    return true && java.util.Objects.equals(this._clientId, o._clientId);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 18;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._clientId);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.ActiveClientRemoved");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._clientId));
    s.append(")");
    return s.toString();
  }
}
