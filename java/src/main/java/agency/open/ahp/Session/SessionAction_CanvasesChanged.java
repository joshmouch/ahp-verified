// Class SessionAction_CanvasesChanged
// Dafny class SessionAction_CanvasesChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_CanvasesChanged extends SessionAction {
  public agency.open.ahp.ConfluxCodec.Json _canvases;
  public SessionAction_CanvasesChanged (agency.open.ahp.ConfluxCodec.Json canvases) {
    super();
    this._canvases = canvases;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_CanvasesChanged o = (SessionAction_CanvasesChanged)other;
    return true && java.util.Objects.equals(this._canvases, o._canvases);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 19;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._canvases);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.CanvasesChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._canvases));
    s.append(")");
    return s.toString();
  }
}
