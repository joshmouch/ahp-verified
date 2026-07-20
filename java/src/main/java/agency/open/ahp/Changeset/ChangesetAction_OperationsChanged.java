// Class ChangesetAction_OperationsChanged
// Dafny class ChangesetAction_OperationsChanged compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChangesetAction_OperationsChanged extends ChangesetAction {
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _operations;
  public ChangesetAction_OperationsChanged (agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> operations) {
    super();
    this._operations = operations;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChangesetAction_OperationsChanged o = (ChangesetAction_OperationsChanged)other;
    return true && java.util.Objects.equals(this._operations, o._operations);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 3;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._operations);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Changeset.ChangesetAction.OperationsChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._operations));
    s.append(")");
    return s.toString();
  }
}
