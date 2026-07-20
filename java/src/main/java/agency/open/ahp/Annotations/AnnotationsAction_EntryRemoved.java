// Class AnnotationsAction_EntryRemoved
// Dafny class AnnotationsAction_EntryRemoved compiled into Java
package agency.open.ahp.Annotations;

@SuppressWarnings({"unchecked", "deprecation"})
public class AnnotationsAction_EntryRemoved extends AnnotationsAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _annotationId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _entryId;
  public AnnotationsAction_EntryRemoved (dafny.DafnySequence<? extends dafny.CodePoint> annotationId, dafny.DafnySequence<? extends dafny.CodePoint> entryId) {
    super();
    this._annotationId = annotationId;
    this._entryId = entryId;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    AnnotationsAction_EntryRemoved o = (AnnotationsAction_EntryRemoved)other;
    return true && java.util.Objects.equals(this._annotationId, o._annotationId) && java.util.Objects.equals(this._entryId, o._entryId);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 3;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._annotationId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._entryId);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Annotations.AnnotationsAction.EntryRemoved");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._annotationId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._entryId));
    s.append(")");
    return s.toString();
  }
}
