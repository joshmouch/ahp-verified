// Class SessionAction_ActivityChanged
// Dafny class SessionAction_ActivityChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_ActivityChanged extends SessionAction {
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _activity;
  public SessionAction_ActivityChanged (agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity) {
    super();
    this._activity = activity;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_ActivityChanged o = (SessionAction_ActivityChanged)other;
    return true && java.util.Objects.equals(this._activity, o._activity);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 7;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activity);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.ActivityChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._activity));
    s.append(")");
    return s.toString();
  }
}
