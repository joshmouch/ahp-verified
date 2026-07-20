// Class SessionAction_IsReadChanged
// Dafny class SessionAction_IsReadChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_IsReadChanged extends SessionAction {
  public boolean _isRead;
  public SessionAction_IsReadChanged (boolean isRead) {
    super();
    this._isRead = isRead;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_IsReadChanged o = (SessionAction_IsReadChanged)other;
    return true && this._isRead == o._isRead;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 5;
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._isRead);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.IsReadChanged");
    s.append("(");
    s.append(this._isRead);
    s.append(")");
    return s.toString();
  }
}
