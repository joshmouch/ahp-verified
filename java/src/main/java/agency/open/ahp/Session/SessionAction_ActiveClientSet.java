// Class SessionAction_ActiveClientSet
// Dafny class SessionAction_ActiveClientSet compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_ActiveClientSet extends SessionAction {
  public Client _client;
  public SessionAction_ActiveClientSet (Client client) {
    super();
    this._client = client;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_ActiveClientSet o = (SessionAction_ActiveClientSet)other;
    return true && java.util.Objects.equals(this._client, o._client);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 17;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._client);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.ActiveClientSet");
    s.append("(");
    s.append(dafny.Helpers.toString(this._client));
    s.append(")");
    return s.toString();
  }
}
