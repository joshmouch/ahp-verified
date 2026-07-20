// Class RouteOp_Remove<S>
// Dafny class RouteOp_Remove<S> compiled into Java
package agency.open.ahp.ConfluxPrune;

@SuppressWarnings({"unchecked", "deprecation"})
public class RouteOp_Remove<S> extends RouteOp<S> {
  public RouteOp_Remove (dafny.TypeDescriptor<S> _td_S) {
    super(_td_S);
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    RouteOp_Remove<S> o = (RouteOp_Remove<S>)other;
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
    s.append("ConfluxPrune.RouteOp.Remove");
    return s.toString();
  }
}
