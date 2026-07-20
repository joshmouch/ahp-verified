// Class SessionAction_CustomizationUpdated
// Dafny class SessionAction_CustomizationUpdated compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_CustomizationUpdated extends SessionAction {
  public Cust _customization;
  public SessionAction_CustomizationUpdated (Cust customization) {
    super();
    this._customization = customization;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_CustomizationUpdated o = (SessionAction_CustomizationUpdated)other;
    return true && java.util.Objects.equals(this._customization, o._customization);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 23;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._customization);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.CustomizationUpdated");
    s.append("(");
    s.append(dafny.Helpers.toString(this._customization));
    s.append(")");
    return s.toString();
  }
}
