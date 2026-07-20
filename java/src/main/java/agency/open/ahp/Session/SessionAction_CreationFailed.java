// Class SessionAction_CreationFailed
// Dafny class SessionAction_CreationFailed compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_CreationFailed extends SessionAction {
  public agency.open.ahp.ConfluxCodec.Json _error;
  public SessionAction_CreationFailed (agency.open.ahp.ConfluxCodec.Json error) {
    super();
    this._error = error;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_CreationFailed o = (SessionAction_CreationFailed)other;
    return true && java.util.Objects.equals(this._error, o._error);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._error);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.CreationFailed");
    s.append("(");
    s.append(dafny.Helpers.toString(this._error));
    s.append(")");
    return s.toString();
  }
}
