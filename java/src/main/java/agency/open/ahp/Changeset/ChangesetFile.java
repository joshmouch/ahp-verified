// Class ChangesetFile
// Dafny class ChangesetFile compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChangesetFile {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public agency.open.ahp.AhpSkeleton.Option<Boolean> _reviewed;
  public agency.open.ahp.ConfluxCodec.Json _edit;
  public ChangesetFile (dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.AhpSkeleton.Option<Boolean> reviewed, agency.open.ahp.ConfluxCodec.Json edit) {
    this._id = id;
    this._reviewed = reviewed;
    this._edit = edit;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChangesetFile o = (ChangesetFile)other;
    return true && java.util.Objects.equals(this._id, o._id) && java.util.Objects.equals(this._reviewed, o._reviewed) && java.util.Objects.equals(this._edit, o._edit);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._reviewed);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._edit);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Changeset.ChangesetFile.ChangesetFile");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._reviewed));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._edit));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<ChangesetFile> _TYPE = dafny.TypeDescriptor.<ChangesetFile>referenceWithInitializer(ChangesetFile.class, () -> ChangesetFile.Default());
  public static dafny.TypeDescriptor<ChangesetFile> _typeDescriptor() {
    return (dafny.TypeDescriptor<ChangesetFile>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ChangesetFile theDefault = ChangesetFile.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<Boolean>Default(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.ConfluxCodec.Json.Default());
  public static ChangesetFile Default() {
    return theDefault;
  }
  public static ChangesetFile create(dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.AhpSkeleton.Option<Boolean> reviewed, agency.open.ahp.ConfluxCodec.Json edit) {
    return new ChangesetFile(id, reviewed, edit);
  }
  public static ChangesetFile create_ChangesetFile(dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.AhpSkeleton.Option<Boolean> reviewed, agency.open.ahp.ConfluxCodec.Json edit) {
    return create(id, reviewed, edit);
  }
  public boolean is_ChangesetFile() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_id() {
    return this._id;
  }
  public agency.open.ahp.AhpSkeleton.Option<Boolean> dtor_reviewed() {
    return this._reviewed;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_edit() {
    return this._edit;
  }
}
