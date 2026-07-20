// Class ChangesetAction_CsUnknown
// Dafny class ChangesetAction_CsUnknown compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChangesetAction_CsUnknown extends ChangesetAction {
  public agency.open.ahp.ConfluxCodec.Json _raw;
  public ChangesetAction_CsUnknown (agency.open.ahp.ConfluxCodec.Json raw) {
    super();
    this._raw = raw;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChangesetAction_CsUnknown o = (ChangesetAction_CsUnknown)other;
    return true && java.util.Objects.equals(this._raw, o._raw);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 8;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._raw);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Changeset.ChangesetAction.CsUnknown");
    s.append("(");
    s.append(dafny.Helpers.toString(this._raw));
    s.append(")");
    return s.toString();
  }
}
