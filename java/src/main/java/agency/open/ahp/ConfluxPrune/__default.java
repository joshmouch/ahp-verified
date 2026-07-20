// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.ConfluxPrune;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static <__K, __S, __A> dafny.Function2<dafny.DafnyMap<? extends __K, ? extends __S>, __A, dafny.DafnyMap<? extends __K, ? extends __S>> PrunableRoute(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__S> _td___S, dafny.TypeDescriptor<__A> _td___A, java.util.function.Function<__A, __K> keyOf, dafny.Function2<agency.open.ahp.ConfluxContract.Option<__S>, __A, RouteOp<__S>> op)
  {
    return ((dafny.Function2<java.util.function.Function<__A, __K>, dafny.Function2<agency.open.ahp.ConfluxContract.Option<__S>, __A, RouteOp<__S>>, dafny.Function2<dafny.DafnyMap<? extends __K, ? extends __S>, __A, dafny.DafnyMap<? extends __K, ? extends __S>>>)(_0_keyOf, _1_op) -> {return ((dafny.Function2<dafny.DafnyMap<? extends __K, ? extends __S>, __A, dafny.DafnyMap<? extends __K, ? extends __S>>)(_2_m_boxed0, _3_a_boxed0) -> {
      dafny.DafnyMap<? extends __K, ? extends __S> _2_m = ((dafny.DafnyMap<? extends __K, ? extends __S>)(java.lang.Object)(_2_m_boxed0));
      __A _3_a = ((__A)(java.lang.Object)(_3_a_boxed0));
      return ((dafny.DafnyMap<? extends __K, ? extends __S>)(java.lang.Object)(dafny.Helpers.<__K, dafny.DafnyMap<? extends __K, ? extends __S>>Let(((__K)(java.lang.Object)((_0_keyOf).apply(_3_a))), boxed16 -> {
        __K _pat_let8_0 = ((__K)(java.lang.Object)(boxed16));
        return ((dafny.DafnyMap<? extends __K, ? extends __S>)(java.lang.Object)(dafny.Helpers.<__K, dafny.DafnyMap<? extends __K, ? extends __S>>Let(_pat_let8_0, boxed17 -> {
          __K _4_k = ((__K)(java.lang.Object)(boxed17));
          return ((dafny.DafnyMap<? extends __K, ? extends __S>)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.ConfluxContract.Option<__S>, dafny.DafnyMap<? extends __K, ? extends __S>>Let((((_2_m).<__K>contains(_4_k)) ? (agency.open.ahp.ConfluxContract.Option.<__S>create_Some(_td___S, ((__S)(java.lang.Object)((_2_m).get(_4_k))))) : (agency.open.ahp.ConfluxContract.Option.<__S>create_None(_td___S))), boxed18 -> {
            agency.open.ahp.ConfluxContract.Option<__S> _pat_let9_0 = ((agency.open.ahp.ConfluxContract.Option<__S>)(java.lang.Object)(boxed18));
            return ((dafny.DafnyMap<? extends __K, ? extends __S>)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.ConfluxContract.Option<__S>, dafny.DafnyMap<? extends __K, ? extends __S>>Let(_pat_let9_0, boxed19 -> {
              agency.open.ahp.ConfluxContract.Option<__S> _5_cur = ((agency.open.ahp.ConfluxContract.Option<__S>)(java.lang.Object)(boxed19));
              return ((java.util.function.Function<RouteOp<__S>, dafny.DafnyMap<? extends __K, ? extends __S>>)(_source0_boxed0) -> {
                RouteOp<__S> _source0 = ((RouteOp<__S>)(java.lang.Object)(_source0_boxed0));
                if (_source0.is_NoOp()) {
                  return _2_m;
                } else if (_source0.is_Install()) {
                  __S _6___mcc_h0 = ((agency.open.ahp.ConfluxPrune.RouteOp_Install<__S>)_source0)._next;
                  __S _7_s = _6___mcc_h0;
                  return dafny.DafnyMap.<__K, __S>update(_2_m, _4_k, _7_s);
                } else {
                  return dafny.DafnyMap.<__K, __S>subtract(_2_m, dafny.DafnySet.<__K> of(_4_k));
                }
              }).apply(((RouteOp<__S>)(java.lang.Object)((_1_op).apply(_5_cur, _3_a))));
            }
            )));
          }
          )));
        }
        )));
      }
      )));
    });}).apply(keyOf, op);
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.ConfluxPrune._default";
  }
}
