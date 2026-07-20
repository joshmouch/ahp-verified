// Class SessionAction_ChatRemoved
// Dafny class SessionAction_ChatRemoved compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_ChatRemoved extends SessionAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _resource;
  public SessionAction_ChatRemoved (dafny.DafnySequence<? extends dafny.CodePoint> resource) {
    super();
    this._resource = resource;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_ChatRemoved o = (SessionAction_ChatRemoved)other;
    return true && java.util.Objects.equals(this._resource, o._resource);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 14;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._resource);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.ChatRemoved");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._resource));
    s.append(")");
    return s.toString();
  }
}
