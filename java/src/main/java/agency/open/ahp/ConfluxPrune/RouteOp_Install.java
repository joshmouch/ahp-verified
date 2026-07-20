// Class RouteOp_Install<S>
// Dafny class RouteOp_Install<S> compiled into Java
package agency.open.ahp.ConfluxPrune;

@SuppressWarnings({"unchecked", "deprecation"})
public class RouteOp_Install<S> extends RouteOp<S> {
  public S _next;
  public RouteOp_Install (dafny.TypeDescriptor<S> _td_S, S next) {
    super(_td_S);
    this._next = next;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    RouteOp_Install<S> o = (RouteOp_Install<S>)other;
    return true && java.util.Objects.equals(this._next, o._next);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._next);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("ConfluxPrune.RouteOp.Install");
    s.append("(");
    s.append(dafny.Helpers.toString(this._next));
    s.append(")");
    return s.toString();
  }
}
