// Class AnnotationsAction_Set
// Dafny class AnnotationsAction_Set compiled into Java
package agency.open.ahp.Annotations;

@SuppressWarnings({"unchecked", "deprecation"})
public class AnnotationsAction_Set extends AnnotationsAction {
  public Annotation _annotation;
  public AnnotationsAction_Set (Annotation annotation) {
    super();
    this._annotation = annotation;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    AnnotationsAction_Set o = (AnnotationsAction_Set)other;
    return true && java.util.Objects.equals(this._annotation, o._annotation);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._annotation);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Annotations.AnnotationsAction.Set");
    s.append("(");
    s.append(dafny.Helpers.toString(this._annotation));
    s.append(")");
    return s.toString();
  }
}
