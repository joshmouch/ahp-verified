// Class SessionAction_SUnknown
// Dafny class SessionAction_SUnknown compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_SUnknown extends SessionAction {
  public agency.open.ahp.ConfluxCodec.Json _raw;
  public SessionAction_SUnknown (agency.open.ahp.ConfluxCodec.Json raw) {
    super();
    this._raw = raw;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_SUnknown o = (SessionAction_SUnknown)other;
    return true && java.util.Objects.equals(this._raw, o._raw);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 27;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._raw);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.SUnknown");
    s.append("(");
    s.append(dafny.Helpers.toString(this._raw));
    s.append(")");
    return s.toString();
  }
}
