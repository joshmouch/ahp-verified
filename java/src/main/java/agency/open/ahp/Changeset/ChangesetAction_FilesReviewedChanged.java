// Class ChangesetAction_FilesReviewedChanged
// Dafny class ChangesetAction_FilesReviewedChanged compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChangesetAction_FilesReviewedChanged extends ChangesetAction {
  public dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> _fileIds;
  public boolean _reviewed;
  public ChangesetAction_FilesReviewedChanged (dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> fileIds, boolean reviewed) {
    super();
    this._fileIds = fileIds;
    this._reviewed = reviewed;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChangesetAction_FilesReviewedChanged o = (ChangesetAction_FilesReviewedChanged)other;
    return true && java.util.Objects.equals(this._fileIds, o._fileIds) && this._reviewed == o._reviewed;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 7;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._fileIds);
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._reviewed);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Changeset.ChangesetAction.FilesReviewedChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._fileIds));
    s.append(", ");
    s.append(this._reviewed);
    s.append(")");
    return s.toString();
  }
}
