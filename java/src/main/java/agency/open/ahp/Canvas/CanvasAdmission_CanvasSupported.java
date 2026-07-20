// Class CanvasAdmission_CanvasSupported
// Dafny class CanvasAdmission_CanvasSupported compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasAdmission_CanvasSupported extends CanvasAdmission {
  public CanvasAdmission_CanvasSupported () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasAdmission_CanvasSupported o = (CanvasAdmission_CanvasSupported)other;
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
    s.append("Canvas.CanvasAdmission.CanvasSupported");
    return s.toString();
  }
}
