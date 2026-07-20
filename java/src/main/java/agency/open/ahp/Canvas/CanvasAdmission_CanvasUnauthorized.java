// Class CanvasAdmission_CanvasUnauthorized
// Dafny class CanvasAdmission_CanvasUnauthorized compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasAdmission_CanvasUnauthorized extends CanvasAdmission {
  public CanvasAdmission_CanvasUnauthorized () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasAdmission_CanvasUnauthorized o = (CanvasAdmission_CanvasUnauthorized)other;
    return true;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 3;
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.CanvasAdmission.CanvasUnauthorized");
    return s.toString();
  }
}
