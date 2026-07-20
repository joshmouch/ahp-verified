// Class SessionAction_CustomizationToggled
// Dafny class SessionAction_CustomizationToggled compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_CustomizationToggled extends SessionAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public boolean _enabled;
  public SessionAction_CustomizationToggled (dafny.DafnySequence<? extends dafny.CodePoint> id, boolean enabled) {
    super();
    this._id = id;
    this._enabled = enabled;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_CustomizationToggled o = (SessionAction_CustomizationToggled)other;
    return true && java.util.Objects.equals(this._id, o._id) && this._enabled == o._enabled;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 10;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._enabled);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.CustomizationToggled");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(this._enabled);
    s.append(")");
    return s.toString();
  }
}
