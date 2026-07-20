// Class ChangesetAction_FileSet
// Dafny class ChangesetAction_FileSet compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChangesetAction_FileSet extends ChangesetAction {
  public ChangesetFile _file;
  public ChangesetAction_FileSet (ChangesetFile file) {
    super();
    this._file = file;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChangesetAction_FileSet o = (ChangesetAction_FileSet)other;
    return true && java.util.Objects.equals(this._file, o._file);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._file);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Changeset.ChangesetAction.FileSet");
    s.append("(");
    s.append(dafny.Helpers.toString(this._file));
    s.append(")");
    return s.toString();
  }
}
