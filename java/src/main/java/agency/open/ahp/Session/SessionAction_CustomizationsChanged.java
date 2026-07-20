// Class SessionAction_CustomizationsChanged
// Dafny class SessionAction_CustomizationsChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_CustomizationsChanged extends SessionAction {
  public dafny.DafnySequence<? extends Cust> _customizations;
  public SessionAction_CustomizationsChanged (dafny.DafnySequence<? extends Cust> customizations) {
    super();
    this._customizations = customizations;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_CustomizationsChanged o = (SessionAction_CustomizationsChanged)other;
    return true && java.util.Objects.equals(this._customizations, o._customizations);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 9;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._customizations);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.CustomizationsChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._customizations));
    s.append(")");
    return s.toString();
  }
}
