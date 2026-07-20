// Class AnnotationsAction_Removed
// Dafny class AnnotationsAction_Removed compiled into Java
package agency.open.ahp.Annotations;

@SuppressWarnings({"unchecked", "deprecation"})
public class AnnotationsAction_Removed extends AnnotationsAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _annotationId;
  public AnnotationsAction_Removed (dafny.DafnySequence<? extends dafny.CodePoint> annotationId) {
    super();
    this._annotationId = annotationId;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    AnnotationsAction_Removed o = (AnnotationsAction_Removed)other;
    return true && java.util.Objects.equals(this._annotationId, o._annotationId);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._annotationId);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Annotations.AnnotationsAction.Removed");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._annotationId));
    s.append(")");
    return s.toString();
  }
}
