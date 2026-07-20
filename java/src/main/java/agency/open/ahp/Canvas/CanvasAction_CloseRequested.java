// Class CanvasAction_CloseRequested
// Dafny class CanvasAction_CloseRequested compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasAction_CloseRequested extends CanvasAction {
  public CanvasAction_CloseRequested () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasAction_CloseRequested o = (CanvasAction_CloseRequested)other;
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
    s.append("Canvas.CanvasAction.CloseRequested");
    return s.toString();
  }
}
