// Class SessionAction_ChangesetsChanged
// Dafny class SessionAction_ChangesetsChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_ChangesetsChanged extends SessionAction {
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _changesets;
  public SessionAction_ChangesetsChanged (agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> changesets) {
    super();
    this._changesets = changesets;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_ChangesetsChanged o = (SessionAction_ChangesetsChanged)other;
    return true && java.util.Objects.equals(this._changesets, o._changesets);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 16;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._changesets);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.ChangesetsChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._changesets));
    s.append(")");
    return s.toString();
  }
}
