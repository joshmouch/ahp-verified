// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.ConfluxOrderedLog;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static <__K, __V> boolean ContainsK(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, dafny.DafnySequence<? extends __V> xs, __K k)
  {
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((xs).length())).signum() == 0) {
        return false;
      } else if (java.util.Objects.equals(((__K)(java.lang.Object)((keyOf).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))), k)) {
        return true;
      } else {
        java.util.function.Function<__V, __K> _in0 = keyOf;
        dafny.DafnySequence<? extends __V> _in1 = (xs).drop(java.math.BigInteger.ONE);
        __K _in2 = k;
        keyOf = _in0;
        xs = _in1;
        k = _in2;
        continue TAIL_CALL_START;
      }
    }
  }
  public static <__K, __V> dafny.DafnySequence<? extends __V> FindK(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, dafny.DafnySequence<? extends __V> xs, __K k)
  {
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((xs).length())).signum() == 0) {
        return dafny.DafnySequence.<__V> empty(_td___V);
      } else if (java.util.Objects.equals(((__K)(java.lang.Object)((keyOf).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))), k)) {
        return dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO))))));
      } else {
        java.util.function.Function<__V, __K> _in0 = keyOf;
        dafny.DafnySequence<? extends __V> _in1 = (xs).drop(java.math.BigInteger.ONE);
        __K _in2 = k;
        keyOf = _in0;
        xs = _in1;
        k = _in2;
        continue TAIL_CALL_START;
      }
    }
  }
  public static <__K, __V> dafny.DafnySequence<? extends __V> TakeListedK(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, dafny.DafnySequence<? extends __V> xs, dafny.DafnySequence<? extends __K> order, dafny.DafnySet<? extends __K> taken)
  {
    dafny.DafnySequence<? extends __V> _0___accumulator = dafny.DafnySequence.<__V> empty(_td___V);
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((order).length())).signum() == 0) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> empty(_td___V));
      } else if ((!(taken).<__K>contains(((__K)(java.lang.Object)((order).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO))))))) && (__default.<__K, __V>ContainsK(_td___K, _td___V, keyOf, xs, ((__K)(java.lang.Object)((order).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))) {
        _0___accumulator = dafny.DafnySequence.<__V>concatenate(_0___accumulator, __default.<__K, __V>FindK(_td___K, _td___V, keyOf, xs, ((__K)(java.lang.Object)((order).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))));
        java.util.function.Function<__V, __K> _in0 = keyOf;
        dafny.DafnySequence<? extends __V> _in1 = xs;
        dafny.DafnySequence<? extends __K> _in2 = (order).drop(java.math.BigInteger.ONE);
        dafny.DafnySet<? extends __K> _in3 = dafny.DafnySet.<__K>union(taken, dafny.DafnySet.<__K> of(((__K)(java.lang.Object)((order).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))));
        keyOf = _in0;
        xs = _in1;
        order = _in2;
        taken = _in3;
        continue TAIL_CALL_START;
      } else {
        java.util.function.Function<__V, __K> _in4 = keyOf;
        dafny.DafnySequence<? extends __V> _in5 = xs;
        dafny.DafnySequence<? extends __K> _in6 = (order).drop(java.math.BigInteger.ONE);
        dafny.DafnySet<? extends __K> _in7 = taken;
        keyOf = _in4;
        xs = _in5;
        order = _in6;
        taken = _in7;
        continue TAIL_CALL_START;
      }
    }
  }
  public static <__K, __V> dafny.DafnySequence<? extends __V> RestK(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, dafny.DafnySequence<? extends __V> xs, dafny.DafnySequence<? extends __K> order)
  {
    dafny.DafnySequence<? extends __V> _0___accumulator = dafny.DafnySequence.<__V> empty(_td___V);
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((xs).length())).signum() == 0) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> empty(_td___V));
      } else {
        _0___accumulator = dafny.DafnySequence.<__V>concatenate(_0___accumulator, (((order).contains(((__K)(java.lang.Object)((keyOf).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))))) ? (dafny.DafnySequence.<__V> empty(_td___V)) : (dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))));
        java.util.function.Function<__V, __K> _in0 = keyOf;
        dafny.DafnySequence<? extends __V> _in1 = (xs).drop(java.math.BigInteger.ONE);
        dafny.DafnySequence<? extends __K> _in2 = order;
        keyOf = _in0;
        xs = _in1;
        order = _in2;
        continue TAIL_CALL_START;
      }
    }
  }
  public static <__K, __V> dafny.DafnySequence<? extends __V> SeqReorderByKeys(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, dafny.DafnySequence<? extends __K> order, dafny.DafnySequence<? extends __V> xs)
  {
    return dafny.DafnySequence.<__V>concatenate(__default.<__K, __V>TakeListedK(_td___K, _td___V, keyOf, xs, order, dafny.DafnySet.<__K> empty()), __default.<__K, __V>RestK(_td___K, _td___V, keyOf, xs, order));
  }
  public static <__K, __V> dafny.DafnySequence<? extends __V> SeqKeepThrough(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, __K k, dafny.DafnySequence<? extends __V> xs)
  {
    dafny.DafnySequence<? extends __V> _0___accumulator = dafny.DafnySequence.<__V> empty(_td___V);
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((xs).length())).signum() == 0) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> empty(_td___V));
      } else if (java.util.Objects.equals(((__K)(java.lang.Object)((keyOf).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))), k)) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))));
      } else {
        _0___accumulator = dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))));
        java.util.function.Function<__V, __K> _in0 = keyOf;
        __K _in1 = k;
        dafny.DafnySequence<? extends __V> _in2 = (xs).drop(java.math.BigInteger.ONE);
        keyOf = _in0;
        k = _in1;
        xs = _in2;
        continue TAIL_CALL_START;
      }
    }
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.ConfluxOrderedLog._default";
  }
}
