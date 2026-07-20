// Class AnnotationsAction_Updated
// Dafny class AnnotationsAction_Updated compiled into Java
package agency.open.ahp.Annotations;

@SuppressWarnings({"unchecked", "deprecation"})
public class AnnotationsAction_Updated extends AnnotationsAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _annotationId;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _turnId;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _resource;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _range;
  public agency.open.ahp.AhpSkeleton.Option<Boolean> _resolved;
  public AnnotationsAction_Updated (dafny.DafnySequence<? extends dafny.CodePoint> annotationId, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> turnId, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> resource, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> range, agency.open.ahp.AhpSkeleton.Option<Boolean> resolved) {
    super();
    this._annotationId = annotationId;
    this._turnId = turnId;
    this._resource = resource;
    this._range = range;
    this._resolved = resolved;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    AnnotationsAction_Updated o = (AnnotationsAction_Updated)other;
    return true && java.util.Objects.equals(this._annotationId, o._annotationId) && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._resource, o._resource) && java.util.Objects.equals(this._range, o._range) && java.util.Objects.equals(this._resolved, o._resolved);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 4;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._annotationId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._resource);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._range);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._resolved);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Annotations.AnnotationsAction.Updated");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._annotationId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._resource));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._range));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._resolved));
    s.append(")");
    return s.toString();
  }
}
