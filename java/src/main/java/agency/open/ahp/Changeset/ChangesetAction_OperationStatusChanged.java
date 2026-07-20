// Class ChangesetAction_OperationStatusChanged
// Dafny class ChangesetAction_OperationStatusChanged compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChangesetAction_OperationStatusChanged extends ChangesetAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _operationId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _status;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _error;
  public ChangesetAction_OperationStatusChanged (dafny.DafnySequence<? extends dafny.CodePoint> operationId, dafny.DafnySequence<? extends dafny.CodePoint> status, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    super();
    this._operationId = operationId;
    this._status = status;
    this._error = error;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChangesetAction_OperationStatusChanged o = (ChangesetAction_OperationStatusChanged)other;
    return true && java.util.Objects.equals(this._operationId, o._operationId) && java.util.Objects.equals(this._status, o._status) && java.util.Objects.equals(this._error, o._error);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 5;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._operationId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._status);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._error);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Changeset.ChangesetAction.OperationStatusChanged");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._operationId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._status));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._error));
    s.append(")");
    return s.toString();
  }
}
