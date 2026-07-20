// Class SessionAction_ServerToolsChanged
// Dafny class SessionAction_ServerToolsChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_ServerToolsChanged extends SessionAction {
  public agency.open.ahp.ConfluxCodec.Json _tools;
  public SessionAction_ServerToolsChanged (agency.open.ahp.ConfluxCodec.Json tools) {
    super();
    this._tools = tools;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_ServerToolsChanged o = (SessionAction_ServerToolsChanged)other;
    return true && java.util.Objects.equals(this._tools, o._tools);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 3;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._tools);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.ServerToolsChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._tools));
    s.append(")");
    return s.toString();
  }
}
