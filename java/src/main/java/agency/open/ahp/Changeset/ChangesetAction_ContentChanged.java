// Class ChangesetAction_ContentChanged
// Dafny class ChangesetAction_ContentChanged compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChangesetAction_ContentChanged extends ChangesetAction {
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetFile>> _files;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _operations;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _error;
  public ChangesetAction_ContentChanged (agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetFile>> files, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> operations, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    super();
    this._files = files;
    this._operations = operations;
    this._error = error;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChangesetAction_ContentChanged o = (ChangesetAction_ContentChanged)other;
    return true && java.util.Objects.equals(this._files, o._files) && java.util.Objects.equals(this._operations, o._operations) && java.util.Objects.equals(this._error, o._error);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 6;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._files);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._operations);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._error);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Changeset.ChangesetAction.ContentChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._files));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._operations));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._error));
    s.append(")");
    return s.toString();
  }
}
