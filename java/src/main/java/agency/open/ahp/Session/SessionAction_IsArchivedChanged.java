// Class SessionAction_IsArchivedChanged
// Dafny class SessionAction_IsArchivedChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_IsArchivedChanged extends SessionAction {
  public boolean _isArchived;
  public SessionAction_IsArchivedChanged (boolean isArchived) {
    super();
    this._isArchived = isArchived;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_IsArchivedChanged o = (SessionAction_IsArchivedChanged)other;
    return true && this._isArchived == o._isArchived;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 6;
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._isArchived);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.IsArchivedChanged");
    s.append("(");
    s.append(this._isArchived);
    s.append(")");
    return s.toString();
  }
}
