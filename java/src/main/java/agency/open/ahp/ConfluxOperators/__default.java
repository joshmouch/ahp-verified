// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.ConfluxOperators;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static <__A, __B> dafny.DafnySequence<? extends __B> Map(dafny.TypeDescriptor<__A> _td___A, dafny.TypeDescriptor<__B> _td___B, java.util.function.Function<__A, __B> f, dafny.DafnySequence<? extends __A> xs)
  {
    return dafny.DafnySequence.Create(_td___B, java.math.BigInteger.valueOf((xs).length()), ((dafny.Function2<java.util.function.Function<__A, __B>, dafny.DafnySequence<? extends __A>, java.util.function.Function<java.math.BigInteger, __B>>)(_0_f, _1_xs) -> {return ((java.util.function.Function<java.math.BigInteger, __B>)(_2_i_boxed0) -> {
      java.math.BigInteger _2_i = ((java.math.BigInteger)(java.lang.Object)(_2_i_boxed0));
      return ((__B)(java.lang.Object)((_0_f).apply(((__A)(java.lang.Object)((_1_xs).select(dafny.Helpers.toInt((_2_i))))))));
    });}).apply(f, xs));
  }
  public static <__A> dafny.DafnySequence<? extends __A> Filter(dafny.TypeDescriptor<__A> _td___A, java.util.function.Function<__A, Boolean> p, dafny.DafnySequence<? extends __A> xs)
  {
    dafny.DafnySequence<? extends __A> _0___accumulator = dafny.DafnySequence.<__A> empty(_td___A);
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((xs).length())).signum() == 0) {
        return dafny.DafnySequence.<__A>concatenate(_0___accumulator, dafny.DafnySequence.<__A> empty(_td___A));
      } else {
        _0___accumulator = dafny.DafnySequence.<__A>concatenate(_0___accumulator, ((((boolean)(java.lang.Object)((p).apply(((__A)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO))))))))) ? (dafny.DafnySequence.<__A> of(_td___A, ((__A)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO))))))) : (dafny.DafnySequence.<__A> empty(_td___A))));
        java.util.function.Function<__A, Boolean> _in0 = p;
        dafny.DafnySequence<? extends __A> _in1 = (xs).drop(java.math.BigInteger.ONE);
        p = _in0;
        xs = _in1;
        continue TAIL_CALL_START;
      }
    }
  }
  public static <__A, __B, __S> __S MapFold(dafny.TypeDescriptor<__A> _td___A, dafny.TypeDescriptor<__B> _td___B, dafny.TypeDescriptor<__S> _td___S, java.util.function.Function<__A, __B> f, dafny.Function2<__S, __B, __S> r, __S init, dafny.DafnySequence<? extends __A> xs)
  {
    return agency.open.ahp.ConfluxContract.__default.<__S, __B>Fold(_td___S, _td___B, r, init, __default.<__A, __B>Map(_td___A, _td___B, f, xs));
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.ConfluxOperators._default";
  }
}
