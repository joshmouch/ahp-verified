// Class CanvasAvailability_Ready
// Dafny class CanvasAvailability_Ready compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasAvailability_Ready extends CanvasAvailability {
  public CanvasAvailability_Ready () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasAvailability_Ready o = (CanvasAvailability_Ready)other;
    return true;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.CanvasAvailability.Ready");
    return s.toString();
  }
}
