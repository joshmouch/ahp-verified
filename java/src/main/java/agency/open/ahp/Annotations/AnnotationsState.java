// Class AnnotationsState
// Dafny class AnnotationsState compiled into Java
package agency.open.ahp.Annotations;

@SuppressWarnings({"unchecked", "deprecation"})
public class AnnotationsState {
  public dafny.DafnySequence<? extends Annotation> _annotations;
  public AnnotationsState (dafny.DafnySequence<? extends Annotation> annotations) {
    this._annotations = annotations;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    AnnotationsState o = (AnnotationsState)other;
    return true && java.util.Objects.equals(this._annotations, o._annotations);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._annotations);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Annotations.AnnotationsState.AnnotationsState");
    s.append("(");
    s.append(dafny.Helpers.toString(this._annotations));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<dafny.DafnySequence<? extends Annotation>> _TYPE = dafny.TypeDescriptor.<dafny.DafnySequence<? extends Annotation>>referenceWithInitializer(dafny.DafnySequence.class, () -> dafny.DafnySequence.<Annotation> empty(Annotation._typeDescriptor()));
  public static dafny.TypeDescriptor<dafny.DafnySequence<? extends Annotation>> _typeDescriptor() {
    return (dafny.TypeDescriptor<dafny.DafnySequence<? extends Annotation>>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final dafny.DafnySequence<? extends Annotation> theDefault = dafny.DafnySequence.<Annotation> empty(Annotation._typeDescriptor());
  public static dafny.DafnySequence<? extends Annotation> Default() {
    return theDefault;
  }
  public static AnnotationsState create(dafny.DafnySequence<? extends Annotation> annotations) {
    return new AnnotationsState(annotations);
  }
  public static AnnotationsState create_AnnotationsState(dafny.DafnySequence<? extends Annotation> annotations) {
    return create(annotations);
  }
  public boolean is_AnnotationsState() { return true; }
  public dafny.DafnySequence<? extends Annotation> dtor_annotations() {
    return this._annotations;
  }
}
