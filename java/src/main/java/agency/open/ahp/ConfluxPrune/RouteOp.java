// Class RouteOp<S>
// Dafny class RouteOp<S> compiled into Java
package agency.open.ahp.ConfluxPrune;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class RouteOp<S> {
  protected dafny.TypeDescriptor<S> _td_S;
  public RouteOp(dafny.TypeDescriptor<S> _td_S) {
    this._td_S = _td_S;
  }
  public static <S> dafny.TypeDescriptor<RouteOp<S>> _typeDescriptor(dafny.TypeDescriptor<S> _td_S) {
    return (dafny.TypeDescriptor<RouteOp<S>>) (dafny.TypeDescriptor<?>) dafny.TypeDescriptor.<RouteOp<S>>referenceWithInitializer(RouteOp.class, () -> RouteOp.<S>Default(_td_S));
  }

  public static <S> RouteOp<S> Default(dafny.TypeDescriptor<S> _td_S) {
    return RouteOp.<S>create_NoOp(_td_S);
  }
  public static <S> RouteOp<S> create_NoOp(dafny.TypeDescriptor<S> _td_S) {
    return new RouteOp_NoOp<S>(_td_S);
  }
  public static <S> RouteOp<S> create_Install(dafny.TypeDescriptor<S> _td_S, S next) {
    return new RouteOp_Install<S>(_td_S, next);
  }
  public static <S> RouteOp<S> create_Remove(dafny.TypeDescriptor<S> _td_S) {
    return new RouteOp_Remove<S>(_td_S);
  }
  public boolean is_NoOp() { return this instanceof RouteOp_NoOp; }
  public boolean is_Install() { return this instanceof RouteOp_Install; }
  public boolean is_Remove() { return this instanceof RouteOp_Remove; }
  public S dtor_next() {
    RouteOp<S> d = this;
    return ((RouteOp_Install<S>)d)._next;
  }
}
