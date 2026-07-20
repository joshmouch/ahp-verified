// Class CanvasAdmission_CanvasUnsupported
// Dafny class CanvasAdmission_CanvasUnsupported compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasAdmission_CanvasUnsupported extends CanvasAdmission {
  public CanvasAdmission_CanvasUnsupported () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasAdmission_CanvasUnsupported o = (CanvasAdmission_CanvasUnsupported)other;
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
    s.append("Canvas.CanvasAdmission.CanvasUnsupported");
    return s.toString();
  }
}
