// Class SessionAction_OpenCanvasesChanged
// Dafny class SessionAction_OpenCanvasesChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_OpenCanvasesChanged extends SessionAction {
  public agency.open.ahp.ConfluxCodec.Json _openCanvases;
  public SessionAction_OpenCanvasesChanged (agency.open.ahp.ConfluxCodec.Json openCanvases) {
    super();
    this._openCanvases = openCanvases;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_OpenCanvasesChanged o = (SessionAction_OpenCanvasesChanged)other;
    return true && java.util.Objects.equals(this._openCanvases, o._openCanvases);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 20;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._openCanvases);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.OpenCanvasesChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._openCanvases));
    s.append(")");
    return s.toString();
  }
}
