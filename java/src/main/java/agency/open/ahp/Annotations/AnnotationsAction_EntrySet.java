// Class AnnotationsAction_EntrySet
// Dafny class AnnotationsAction_EntrySet compiled into Java
package agency.open.ahp.Annotations;

@SuppressWarnings({"unchecked", "deprecation"})
public class AnnotationsAction_EntrySet extends AnnotationsAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _annotationId;
  public Entry _entry;
  public AnnotationsAction_EntrySet (dafny.DafnySequence<? extends dafny.CodePoint> annotationId, Entry entry) {
    super();
    this._annotationId = annotationId;
    this._entry = entry;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    AnnotationsAction_EntrySet o = (AnnotationsAction_EntrySet)other;
    return true && java.util.Objects.equals(this._annotationId, o._annotationId) && java.util.Objects.equals(this._entry, o._entry);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 2;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._annotationId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._entry);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Annotations.AnnotationsAction.EntrySet");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._annotationId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._entry));
    s.append(")");
    return s.toString();
  }
}
