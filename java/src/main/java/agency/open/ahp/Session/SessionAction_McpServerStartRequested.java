// Class SessionAction_McpServerStartRequested
// Dafny class SessionAction_McpServerStartRequested compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_McpServerStartRequested extends SessionAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public SessionAction_McpServerStartRequested (dafny.DafnySequence<? extends dafny.CodePoint> id) {
    super();
    this._id = id;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_McpServerStartRequested o = (SessionAction_McpServerStartRequested)other;
    return true && java.util.Objects.equals(this._id, o._id);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 25;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.McpServerStartRequested");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(")");
    return s.toString();
  }
}
