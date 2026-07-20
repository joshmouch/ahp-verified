// Class SessionAction_InputNeededRemoved
// Dafny class SessionAction_InputNeededRemoved compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_InputNeededRemoved extends SessionAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public SessionAction_InputNeededRemoved (dafny.DafnySequence<? extends dafny.CodePoint> id) {
    super();
    this._id = id;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_InputNeededRemoved o = (SessionAction_InputNeededRemoved)other;
    return true && java.util.Objects.equals(this._id, o._id);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 22;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.InputNeededRemoved");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(")");
    return s.toString();
  }
}
