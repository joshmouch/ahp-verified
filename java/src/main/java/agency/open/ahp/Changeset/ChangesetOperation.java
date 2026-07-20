// Class ChangesetOperation
// Dafny class ChangesetOperation compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChangesetOperation {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public dafny.DafnySequence<? extends dafny.CodePoint> _label__;
  public dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> _scopes;
  public dafny.DafnySequence<? extends dafny.CodePoint> _status;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _error;
  public ChangesetOperation (dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> label__, dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> scopes, dafny.DafnySequence<? extends dafny.CodePoint> status, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    this._id = id;
    this._label__ = label__;
    this._scopes = scopes;
    this._status = status;
    this._error = error;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChangesetOperation o = (ChangesetOperation)other;
    return true && java.util.Objects.equals(this._id, o._id) && java.util.Objects.equals(this._label__, o._label__) && java.util.Objects.equals(this._scopes, o._scopes) && java.util.Objects.equals(this._status, o._status) && java.util.Objects.equals(this._error, o._error);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._label__);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._scopes);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._status);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._error);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Changeset.ChangesetOperation.ChangesetOperation");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._label__));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._scopes));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._status));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._error));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<ChangesetOperation> _TYPE = dafny.TypeDescriptor.<ChangesetOperation>referenceWithInitializer(ChangesetOperation.class, () -> ChangesetOperation.Default());
  public static dafny.TypeDescriptor<ChangesetOperation> _typeDescriptor() {
    return (dafny.TypeDescriptor<ChangesetOperation>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ChangesetOperation theDefault = ChangesetOperation.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.DafnySequence<? extends dafny.CodePoint>> empty(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
  public static ChangesetOperation Default() {
    return theDefault;
  }
  public static ChangesetOperation create(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> label__, dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> scopes, dafny.DafnySequence<? extends dafny.CodePoint> status, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    return new ChangesetOperation(id, label__, scopes, status, error);
  }
  public static ChangesetOperation create_ChangesetOperation(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> label__, dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> scopes, dafny.DafnySequence<? extends dafny.CodePoint> status, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    return create(id, label__, scopes, status, error);
  }
  public boolean is_ChangesetOperation() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_id() {
    return this._id;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_label__() {
    return this._label__;
  }
  public dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> dtor_scopes() {
    return this._scopes;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_status() {
    return this._status;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_error() {
    return this._error;
  }
}
