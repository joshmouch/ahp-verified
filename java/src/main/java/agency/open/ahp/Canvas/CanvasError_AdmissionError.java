// Class CanvasError_AdmissionError
// Dafny class CanvasError_AdmissionError compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasError_AdmissionError extends CanvasError {
  public CanvasAdmission _reason;
  public CanvasError_AdmissionError (CanvasAdmission reason) {
    super();
    this._reason = reason;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasError_AdmissionError o = (CanvasError_AdmissionError)other;
    return true && java.util.Objects.equals(this._reason, o._reason);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._reason);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.CanvasError.AdmissionError");
    s.append("(");
    s.append(dafny.Helpers.toString(this._reason));
    s.append(")");
    return s.toString();
  }
}
