// Class RootAction_RootActiveSessionsChanged
// Dafny class RootAction_RootActiveSessionsChanged compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class RootAction_RootActiveSessionsChanged extends RootAction {
  public java.math.BigInteger _activeSessions;
  public RootAction_RootActiveSessionsChanged (java.math.BigInteger activeSessions) {
    super();
    this._activeSessions = activeSessions;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    RootAction_RootActiveSessionsChanged o = (RootAction_RootActiveSessionsChanged)other;
    return true && java.util.Objects.equals(this._activeSessions, o._activeSessions);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activeSessions);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.RootAction.RootActiveSessionsChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._activeSessions));
    s.append(")");
    return s.toString();
  }
}
