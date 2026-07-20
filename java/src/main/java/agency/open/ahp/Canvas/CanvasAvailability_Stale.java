// Class CanvasAvailability_Stale
// Dafny class CanvasAvailability_Stale compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasAvailability_Stale extends CanvasAvailability {
  public CanvasAvailability_Stale () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasAvailability_Stale o = (CanvasAvailability_Stale)other;
    return true;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.CanvasAvailability.Stale");
    return s.toString();
  }
}
