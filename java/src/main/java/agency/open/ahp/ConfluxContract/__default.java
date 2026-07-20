// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.ConfluxContract;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static <__S, __A> __S Fold(dafny.TypeDescriptor<__S> _td___S, dafny.TypeDescriptor<__A> _td___A, dafny.Function2<__S, __A, __S> r, __S s, dafny.DafnySequence<? extends __A> actions)
  {
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((actions).length())).signum() == 0) {
        return s;
      } else {
        dafny.Function2<__S, __A, __S> _in0 = r;
        __S _in1 = ((__S)(java.lang.Object)((r).apply(s, ((__A)(java.lang.Object)((actions).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO))))))));
        dafny.DafnySequence<? extends __A> _in2 = (actions).drop(java.math.BigInteger.ONE);
        r = _in0;
        s = _in1;
        actions = _in2;
        continue TAIL_CALL_START;
      }
    }
  }
  public static <__T> __T Iterate(dafny.TypeDescriptor<__T> _td___T, java.util.function.Function<__T, __T> f, __T x, java.math.BigInteger n)
  {
    TAIL_CALL_START: while (true) {
      if ((n).signum() == 0) {
        return x;
      } else {
        java.util.function.Function<__T, __T> _in0 = f;
        __T _in1 = ((__T)(java.lang.Object)((f).apply(x)));
        java.math.BigInteger _in2 = (n).subtract(java.math.BigInteger.ONE);
        f = _in0;
        x = _in1;
        n = _in2;
        continue TAIL_CALL_START;
      }
    }
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.ConfluxContract._default";
  }
}
