// Class SessionAction_MetaChanged
// Dafny class SessionAction_MetaChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_MetaChanged extends SessionAction {
  public agency.open.ahp.ConfluxCodec.Json _meta;
  public SessionAction_MetaChanged (agency.open.ahp.ConfluxCodec.Json meta) {
    super();
    this._meta = meta;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_MetaChanged o = (SessionAction_MetaChanged)other;
    return true && java.util.Objects.equals(this._meta, o._meta);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 4;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._meta);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.MetaChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._meta));
    s.append(")");
    return s.toString();
  }
}
