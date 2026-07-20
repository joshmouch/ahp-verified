// Class ChangesetAction
// Dafny class ChangesetAction compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class ChangesetAction {
  public ChangesetAction() {
  }
  private static final dafny.TypeDescriptor<ChangesetAction> _TYPE = dafny.TypeDescriptor.<ChangesetAction>referenceWithInitializer(ChangesetAction.class, () -> ChangesetAction.Default());
  public static dafny.TypeDescriptor<ChangesetAction> _typeDescriptor() {
    return (dafny.TypeDescriptor<ChangesetAction>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ChangesetAction theDefault = ChangesetAction.create_StatusChanged(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
  public static ChangesetAction Default() {
    return theDefault;
  }
  public static ChangesetAction create_StatusChanged(dafny.DafnySequence<? extends dafny.CodePoint> status, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    return new ChangesetAction_StatusChanged(status, error);
  }
  public static ChangesetAction create_FileSet(ChangesetFile file) {
    return new ChangesetAction_FileSet(file);
  }
  public static ChangesetAction create_FileRemoved(dafny.DafnySequence<? extends dafny.CodePoint> fileId) {
    return new ChangesetAction_FileRemoved(fileId);
  }
  public static ChangesetAction create_OperationsChanged(agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> operations) {
    return new ChangesetAction_OperationsChanged(operations);
  }
  public static ChangesetAction create_Cleared() {
    return new ChangesetAction_Cleared();
  }
  public static ChangesetAction create_OperationStatusChanged(dafny.DafnySequence<? extends dafny.CodePoint> operationId, dafny.DafnySequence<? extends dafny.CodePoint> status, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    return new ChangesetAction_OperationStatusChanged(operationId, status, error);
  }
  public static ChangesetAction create_ContentChanged(agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetFile>> files, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> operations, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    return new ChangesetAction_ContentChanged(files, operations, error);
  }
  public static ChangesetAction create_FilesReviewedChanged(dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> fileIds, boolean reviewed) {
    return new ChangesetAction_FilesReviewedChanged(fileIds, reviewed);
  }
  public static ChangesetAction create_CsUnknown(agency.open.ahp.ConfluxCodec.Json raw) {
    return new ChangesetAction_CsUnknown(raw);
  }
  public boolean is_StatusChanged() { return this instanceof ChangesetAction_StatusChanged; }
  public boolean is_FileSet() { return this instanceof ChangesetAction_FileSet; }
  public boolean is_FileRemoved() { return this instanceof ChangesetAction_FileRemoved; }
  public boolean is_OperationsChanged() { return this instanceof ChangesetAction_OperationsChanged; }
  public boolean is_Cleared() { return this instanceof ChangesetAction_Cleared; }
  public boolean is_OperationStatusChanged() { return this instanceof ChangesetAction_OperationStatusChanged; }
  public boolean is_ContentChanged() { return this instanceof ChangesetAction_ContentChanged; }
  public boolean is_FilesReviewedChanged() { return this instanceof ChangesetAction_FilesReviewedChanged; }
  public boolean is_CsUnknown() { return this instanceof ChangesetAction_CsUnknown; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_status() {
    ChangesetAction d = this;
    if (d instanceof ChangesetAction_StatusChanged) { return ((ChangesetAction_StatusChanged)d)._status; }
    return ((ChangesetAction_OperationStatusChanged)d)._status;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_error() {
    ChangesetAction d = this;
    if (d instanceof ChangesetAction_StatusChanged) { return ((ChangesetAction_StatusChanged)d)._error; }
    if (d instanceof ChangesetAction_OperationStatusChanged) { return ((ChangesetAction_OperationStatusChanged)d)._error; }
    return ((ChangesetAction_ContentChanged)d)._error;
  }
  public ChangesetFile dtor_file() {
    ChangesetAction d = this;
    return ((ChangesetAction_FileSet)d)._file;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_fileId() {
    ChangesetAction d = this;
    return ((ChangesetAction_FileRemoved)d)._fileId;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> dtor_operations() {
    ChangesetAction d = this;
    if (d instanceof ChangesetAction_OperationsChanged) { return ((ChangesetAction_OperationsChanged)d)._operations; }
    return ((ChangesetAction_ContentChanged)d)._operations;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_operationId() {
    ChangesetAction d = this;
    return ((ChangesetAction_OperationStatusChanged)d)._operationId;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetFile>> dtor_files() {
    ChangesetAction d = this;
    return ((ChangesetAction_ContentChanged)d)._files;
  }
  public dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> dtor_fileIds() {
    ChangesetAction d = this;
    return ((ChangesetAction_FilesReviewedChanged)d)._fileIds;
  }
  public boolean dtor_reviewed() {
    ChangesetAction d = this;
    return ((ChangesetAction_FilesReviewedChanged)d)._reviewed;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    ChangesetAction d = this;
    return ((ChangesetAction_CsUnknown)d)._raw;
  }
}
