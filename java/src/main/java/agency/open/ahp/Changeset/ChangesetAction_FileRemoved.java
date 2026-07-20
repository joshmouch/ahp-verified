// Class ChangesetAction_FileRemoved
// Dafny class ChangesetAction_FileRemoved compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChangesetAction_FileRemoved extends ChangesetAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _fileId;
  public ChangesetAction_FileRemoved (dafny.DafnySequence<? extends dafny.CodePoint> fileId) {
    super();
    this._fileId = fileId;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChangesetAction_FileRemoved o = (ChangesetAction_FileRemoved)other;
    return true && java.util.Objects.equals(this._fileId, o._fileId);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 2;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._fileId);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Changeset.ChangesetAction.FileRemoved");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._fileId));
    s.append(")");
    return s.toString();
  }
}
