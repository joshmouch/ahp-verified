// Class ChangesetState
// Dafny class ChangesetState compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChangesetState {
  public dafny.DafnySequence<? extends dafny.CodePoint> _status;
  public dafny.DafnySequence<? extends ChangesetFile> _files;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _operations;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _error;
  public ChangesetState (dafny.DafnySequence<? extends dafny.CodePoint> status, dafny.DafnySequence<? extends ChangesetFile> files, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> operations, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    this._status = status;
    this._files = files;
    this._operations = operations;
    this._error = error;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChangesetState o = (ChangesetState)other;
    return true && java.util.Objects.equals(this._status, o._status) && java.util.Objects.equals(this._files, o._files) && java.util.Objects.equals(this._operations, o._operations) && java.util.Objects.equals(this._error, o._error);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._status);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._files);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._operations);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._error);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Changeset.ChangesetState.ChangesetState");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._status));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._files));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._operations));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._error));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<ChangesetState> _TYPE = dafny.TypeDescriptor.<ChangesetState>referenceWithInitializer(ChangesetState.class, () -> ChangesetState.Default());
  public static dafny.TypeDescriptor<ChangesetState> _typeDescriptor() {
    return (dafny.TypeDescriptor<ChangesetState>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ChangesetState theDefault = ChangesetState.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>Default(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
  public static ChangesetState Default() {
    return theDefault;
  }
  public static ChangesetState create(dafny.DafnySequence<? extends dafny.CodePoint> status, dafny.DafnySequence<? extends ChangesetFile> files, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> operations, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    return new ChangesetState(status, files, operations, error);
  }
  public static ChangesetState create_ChangesetState(dafny.DafnySequence<? extends dafny.CodePoint> status, dafny.DafnySequence<? extends ChangesetFile> files, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> operations, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    return create(status, files, operations, error);
  }
  public boolean is_ChangesetState() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_status() {
    return this._status;
  }
  public dafny.DafnySequence<? extends ChangesetFile> dtor_files() {
    return this._files;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> dtor_operations() {
    return this._operations;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_error() {
    return this._error;
  }
}
