// Class CanvasAdmission_CanvasHandlerUnavailable
// Dafny class CanvasAdmission_CanvasHandlerUnavailable compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasAdmission_CanvasHandlerUnavailable extends CanvasAdmission {
  public CanvasAdmission_CanvasHandlerUnavailable () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasAdmission_CanvasHandlerUnavailable o = (CanvasAdmission_CanvasHandlerUnavailable)other;
    return true;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 2;
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.CanvasAdmission.CanvasHandlerUnavailable");
    return s.toString();
  }
}
