// Class RootAction_RootTerminalsChanged
// Dafny class RootAction_RootTerminalsChanged compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class RootAction_RootTerminalsChanged extends RootAction {
  public dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json> _terminals;
  public RootAction_RootTerminalsChanged (dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json> terminals) {
    super();
    this._terminals = terminals;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    RootAction_RootTerminalsChanged o = (RootAction_RootTerminalsChanged)other;
    return true && java.util.Objects.equals(this._terminals, o._terminals);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 2;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._terminals);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.RootAction.RootTerminalsChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._terminals));
    s.append(")");
    return s.toString();
  }
}
