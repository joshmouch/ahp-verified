// Class RouteOp_NoOp<S>
// Dafny class RouteOp_NoOp<S> compiled into Java
package agency.open.ahp.ConfluxPrune;

@SuppressWarnings({"unchecked", "deprecation"})
public class RouteOp_NoOp<S> extends RouteOp<S> {
  public RouteOp_NoOp (dafny.TypeDescriptor<S> _td_S) {
    super(_td_S);
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    RouteOp_NoOp<S> o = (RouteOp_NoOp<S>)other;
    return true;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("ConfluxPrune.RouteOp.NoOp");
    return s.toString();
  }
}
