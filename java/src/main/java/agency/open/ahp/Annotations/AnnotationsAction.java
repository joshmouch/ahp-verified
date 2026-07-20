// Class AnnotationsAction
// Dafny class AnnotationsAction compiled into Java
package agency.open.ahp.Annotations;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class AnnotationsAction {
  public AnnotationsAction() {
  }
  private static final dafny.TypeDescriptor<AnnotationsAction> _TYPE = dafny.TypeDescriptor.<AnnotationsAction>referenceWithInitializer(AnnotationsAction.class, () -> AnnotationsAction.Default());
  public static dafny.TypeDescriptor<AnnotationsAction> _typeDescriptor() {
    return (dafny.TypeDescriptor<AnnotationsAction>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final AnnotationsAction theDefault = AnnotationsAction.create_Set(Annotation.Default());
  public static AnnotationsAction Default() {
    return theDefault;
  }
  public static AnnotationsAction create_Set(Annotation annotation) {
    return new AnnotationsAction_Set(annotation);
  }
  public static AnnotationsAction create_Removed(dafny.DafnySequence<? extends dafny.CodePoint> annotationId) {
    return new AnnotationsAction_Removed(annotationId);
  }
  public static AnnotationsAction create_EntrySet(dafny.DafnySequence<? extends dafny.CodePoint> annotationId, Entry entry) {
    return new AnnotationsAction_EntrySet(annotationId, entry);
  }
  public static AnnotationsAction create_EntryRemoved(dafny.DafnySequence<? extends dafny.CodePoint> annotationId, dafny.DafnySequence<? extends dafny.CodePoint> entryId) {
    return new AnnotationsAction_EntryRemoved(annotationId, entryId);
  }
  public static AnnotationsAction create_Updated(dafny.DafnySequence<? extends dafny.CodePoint> annotationId, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> turnId, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> resource, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> range, agency.open.ahp.AhpSkeleton.Option<Boolean> resolved) {
    return new AnnotationsAction_Updated(annotationId, turnId, resource, range, resolved);
  }
  public static AnnotationsAction create_AnUnknown(agency.open.ahp.ConfluxCodec.Json raw) {
    return new AnnotationsAction_AnUnknown(raw);
  }
  public boolean is_Set() { return this instanceof AnnotationsAction_Set; }
  public boolean is_Removed() { return this instanceof AnnotationsAction_Removed; }
  public boolean is_EntrySet() { return this instanceof AnnotationsAction_EntrySet; }
  public boolean is_EntryRemoved() { return this instanceof AnnotationsAction_EntryRemoved; }
  public boolean is_Updated() { return this instanceof AnnotationsAction_Updated; }
  public boolean is_AnUnknown() { return this instanceof AnnotationsAction_AnUnknown; }
  public Annotation dtor_annotation() {
    AnnotationsAction d = this;
    return ((AnnotationsAction_Set)d)._annotation;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_annotationId() {
    AnnotationsAction d = this;
    if (d instanceof AnnotationsAction_Removed) { return ((AnnotationsAction_Removed)d)._annotationId; }
    if (d instanceof AnnotationsAction_EntrySet) { return ((AnnotationsAction_EntrySet)d)._annotationId; }
    if (d instanceof AnnotationsAction_EntryRemoved) { return ((AnnotationsAction_EntryRemoved)d)._annotationId; }
    return ((AnnotationsAction_Updated)d)._annotationId;
  }
  public Entry dtor_entry() {
    AnnotationsAction d = this;
    return ((AnnotationsAction_EntrySet)d)._entry;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_entryId() {
    AnnotationsAction d = this;
    return ((AnnotationsAction_EntryRemoved)d)._entryId;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_turnId() {
    AnnotationsAction d = this;
    return ((AnnotationsAction_Updated)d)._turnId;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_resource() {
    AnnotationsAction d = this;
    return ((AnnotationsAction_Updated)d)._resource;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_range() {
    AnnotationsAction d = this;
    return ((AnnotationsAction_Updated)d)._range;
  }
  public agency.open.ahp.AhpSkeleton.Option<Boolean> dtor_resolved() {
    AnnotationsAction d = this;
    return ((AnnotationsAction_Updated)d)._resolved;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    AnnotationsAction d = this;
    return ((AnnotationsAction_AnUnknown)d)._raw;
  }
}
