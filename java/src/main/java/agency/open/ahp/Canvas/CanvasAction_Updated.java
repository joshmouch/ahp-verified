// Class CanvasAction_Updated
// Dafny class CanvasAction_Updated compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasAction_Updated extends CanvasAction {
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _title;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _activity;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _contentUri;
  public agency.open.ahp.AhpSkeleton.Option<CanvasAvailability> _availability;
  public CanvasAction_Updated (agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> title, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> contentUri, agency.open.ahp.AhpSkeleton.Option<CanvasAvailability> availability) {
    super();
    this._title = title;
    this._activity = activity;
    this._contentUri = contentUri;
    this._availability = availability;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasAction_Updated o = (CanvasAction_Updated)other;
    return true && java.util.Objects.equals(this._title, o._title) && java.util.Objects.equals(this._activity, o._activity) && java.util.Objects.equals(this._contentUri, o._contentUri) && java.util.Objects.equals(this._availability, o._availability);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._title);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activity);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._contentUri);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._availability);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.CanvasAction.Updated");
    s.append("(");
    s.append(dafny.Helpers.toString(this._title));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._activity));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._contentUri));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._availability));
    s.append(")");
    return s.toString();
  }
}
