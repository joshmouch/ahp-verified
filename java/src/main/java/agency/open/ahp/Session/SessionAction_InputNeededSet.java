// Class SessionAction_InputNeededSet
// Dafny class SessionAction_InputNeededSet compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_InputNeededSet extends SessionAction {
  public InputReq _req;
  public SessionAction_InputNeededSet (InputReq req) {
    super();
    this._req = req;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_InputNeededSet o = (SessionAction_InputNeededSet)other;
    return true && java.util.Objects.equals(this._req, o._req);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 21;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._req);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.InputNeededSet");
    s.append("(");
    s.append(dafny.Helpers.toString(this._req));
    s.append(")");
    return s.toString();
  }
}
