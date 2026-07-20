// Class SessionAction_ChatAdded
// Dafny class SessionAction_ChatAdded compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_ChatAdded extends SessionAction {
  public Chat _summary;
  public SessionAction_ChatAdded (Chat summary) {
    super();
    this._summary = summary;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_ChatAdded o = (SessionAction_ChatAdded)other;
    return true && java.util.Objects.equals(this._summary, o._summary);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 13;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._summary);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.ChatAdded");
    s.append("(");
    s.append(dafny.Helpers.toString(this._summary));
    s.append(")");
    return s.toString();
  }
}
