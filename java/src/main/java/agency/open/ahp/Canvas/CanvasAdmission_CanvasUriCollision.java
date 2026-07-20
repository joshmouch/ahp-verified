// Class CanvasAdmission_CanvasUriCollision
// Dafny class CanvasAdmission_CanvasUriCollision compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasAdmission_CanvasUriCollision extends CanvasAdmission {
  public CanvasAdmission_CanvasUriCollision () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasAdmission_CanvasUriCollision o = (CanvasAdmission_CanvasUriCollision)other;
    return true;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 4;
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.CanvasAdmission.CanvasUriCollision");
    return s.toString();
  }
}
