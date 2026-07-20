// Class RootAction_RootAgentsChanged
// Dafny class RootAction_RootAgentsChanged compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class RootAction_RootAgentsChanged extends RootAction {
  public dafny.DafnySequence<? extends AgentInfo> _agents;
  public RootAction_RootAgentsChanged (dafny.DafnySequence<? extends AgentInfo> agents) {
    super();
    this._agents = agents;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    RootAction_RootAgentsChanged o = (RootAction_RootAgentsChanged)other;
    return true && java.util.Objects.equals(this._agents, o._agents);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._agents);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.RootAction.RootAgentsChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._agents));
    s.append(")");
    return s.toString();
  }
}
