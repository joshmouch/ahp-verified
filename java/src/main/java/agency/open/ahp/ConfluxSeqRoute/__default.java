// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.ConfluxSeqRoute;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static <__K, __V> dafny.DafnySequence<? extends __V> SeqUpsertById(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, dafny.DafnySequence<? extends __V> xs, __V v)
  {
    dafny.DafnySequence<? extends __V> _0___accumulator = dafny.DafnySequence.<__V> empty(_td___V);
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((xs).length())).signum() == 0) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> of(_td___V, v));
      } else if (java.util.Objects.equals(((__K)(java.lang.Object)((keyOf).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))), ((__K)(java.lang.Object)((keyOf).apply(v))))) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V>concatenate(dafny.DafnySequence.<__V> of(_td___V, v), (xs).drop(java.math.BigInteger.ONE)));
      } else {
        _0___accumulator = dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))));
        java.util.function.Function<__V, __K> _in0 = keyOf;
        dafny.DafnySequence<? extends __V> _in1 = (xs).drop(java.math.BigInteger.ONE);
        __V _in2 = v;
        keyOf = _in0;
        xs = _in1;
        v = _in2;
        continue TAIL_CALL_START;
      }
    }
  }
  public static <__K, __V> dafny.DafnySequence<? extends __V> SeqRemoveById(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, dafny.DafnySequence<? extends __V> xs, __K k)
  {
    dafny.DafnySequence<? extends __V> _0___accumulator = dafny.DafnySequence.<__V> empty(_td___V);
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((xs).length())).signum() == 0) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> empty(_td___V));
      } else if (java.util.Objects.equals(((__K)(java.lang.Object)((keyOf).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))), k)) {
        java.util.function.Function<__V, __K> _in0 = keyOf;
        dafny.DafnySequence<? extends __V> _in1 = (xs).drop(java.math.BigInteger.ONE);
        __K _in2 = k;
        keyOf = _in0;
        xs = _in1;
        k = _in2;
        continue TAIL_CALL_START;
      } else {
        _0___accumulator = dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))));
        java.util.function.Function<__V, __K> _in3 = keyOf;
        dafny.DafnySequence<? extends __V> _in4 = (xs).drop(java.math.BigInteger.ONE);
        __K _in5 = k;
        keyOf = _in3;
        xs = _in4;
        k = _in5;
        continue TAIL_CALL_START;
      }
    }
  }
  public static <__K, __V> dafny.DafnyMap<? extends __K, ? extends __V> ToMap(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, dafny.DafnySequence<? extends __V> xs)
  {
    if ((java.math.BigInteger.valueOf((xs).length())).signum() == 0) {
      return dafny.DafnyMap.fromElements((dafny.Tuple2<__K, __V>[])new dafny.Tuple2[]{ });
    } else {
      return dafny.DafnyMap.<__K, __V>update(__default.<__K, __V>ToMap(_td___K, _td___V, keyOf, (xs).drop(java.math.BigInteger.ONE)), ((__K)(java.lang.Object)((keyOf).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))), ((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO))))));
    }
  }
  public static <__K, __V> boolean UniqueKeys(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, dafny.DafnySequence<? extends __V> xs)
  {
    return ((dafny.Function2<dafny.DafnySequence<? extends __V>, java.util.function.Function<__V, __K>, Boolean>)(_0_xs, _1_keyOf) -> {return dafny.Helpers.Quantifier(dafny.Helpers.IntegerRange(java.math.BigInteger.ZERO, java.math.BigInteger.valueOf((_0_xs).length())), true, ((_forall_var_0_boxed0) -> {
      java.math.BigInteger _forall_var_0 = ((java.math.BigInteger)(java.lang.Object)(_forall_var_0_boxed0));
      java.math.BigInteger _2_i = (java.math.BigInteger)_forall_var_0;
      return dafny.Helpers.Quantifier(dafny.Helpers.IntegerRange(java.math.BigInteger.ZERO, java.math.BigInteger.valueOf((_0_xs).length())), true, ((_forall_var_1_boxed0) -> {
        java.math.BigInteger _forall_var_1 = ((java.math.BigInteger)(java.lang.Object)(_forall_var_1_boxed0));
        java.math.BigInteger _3_j = (java.math.BigInteger)_forall_var_1;
        return !(((((_2_i).signum() != -1) && ((_2_i).compareTo(java.math.BigInteger.valueOf((_0_xs).length())) < 0)) && (((_3_j).signum() != -1) && ((_3_j).compareTo(java.math.BigInteger.valueOf((_0_xs).length())) < 0))) && (java.util.Objects.equals(((__K)(java.lang.Object)((_1_keyOf).apply(((__V)(java.lang.Object)((_0_xs).select(dafny.Helpers.toInt((_2_i)))))))), ((__K)(java.lang.Object)((_1_keyOf).apply(((__V)(java.lang.Object)((_0_xs).select(dafny.Helpers.toInt((_3_j))))))))))) || (java.util.Objects.equals(_2_i, _3_j));
      }));
    }));}).apply(xs, keyOf);
  }
  public static <__V> agency.open.ahp.ConfluxPrune.RouteOp<__V> InstallOp(dafny.TypeDescriptor<__V> _td___V, agency.open.ahp.ConfluxContract.Option<__V> cur, __V v)
  {
    return agency.open.ahp.ConfluxPrune.RouteOp.<__V>create_Install(_td___V, v);
  }
  public static <__V> agency.open.ahp.ConfluxPrune.RouteOp<__V> RemoveOp(dafny.TypeDescriptor<__V> _td___V, agency.open.ahp.ConfluxContract.Option<__V> cur, __V v)
  {
    return agency.open.ahp.ConfluxPrune.RouteOp.<__V>create_Remove(_td___V);
  }
  public static <__K, __V> dafny.DafnySequence<? extends __V> SeqUpdateById(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, dafny.DafnySequence<? extends __V> xs, __K k, java.util.function.Function<__V, __V> f)
  {
    dafny.DafnySequence<? extends __V> _0___accumulator = dafny.DafnySequence.<__V> empty(_td___V);
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((xs).length())).signum() == 0) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> empty(_td___V));
      } else if (java.util.Objects.equals(((__K)(java.lang.Object)((keyOf).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))), k)) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V>concatenate(dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((f).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO))))))))), (xs).drop(java.math.BigInteger.ONE)));
      } else {
        _0___accumulator = dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))));
        java.util.function.Function<__V, __K> _in0 = keyOf;
        dafny.DafnySequence<? extends __V> _in1 = (xs).drop(java.math.BigInteger.ONE);
        __K _in2 = k;
        java.util.function.Function<__V, __V> _in3 = f;
        keyOf = _in0;
        xs = _in1;
        k = _in2;
        f = _in3;
        continue TAIL_CALL_START;
      }
    }
  }
  public static <__V> dafny.Function2<agency.open.ahp.ConfluxContract.Option<__V>, __V, agency.open.ahp.ConfluxPrune.RouteOp<__V>> UpdateOp(dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __V> f)
  {
    return ((java.util.function.Function<java.util.function.Function<__V, __V>, dafny.Function2<agency.open.ahp.ConfluxContract.Option<__V>, __V, agency.open.ahp.ConfluxPrune.RouteOp<__V>>>)(_0_f) -> {return ((dafny.Function2<agency.open.ahp.ConfluxContract.Option<__V>, __V, agency.open.ahp.ConfluxPrune.RouteOp<__V>>)(_1_cur_boxed0, _2_a_boxed0) -> {
      agency.open.ahp.ConfluxContract.Option<__V> _1_cur = ((agency.open.ahp.ConfluxContract.Option<__V>)(java.lang.Object)(_1_cur_boxed0));
      __V _2_a = ((__V)(java.lang.Object)(_2_a_boxed0));
      return ((java.util.function.Function<agency.open.ahp.ConfluxContract.Option<__V>, agency.open.ahp.ConfluxPrune.RouteOp<__V>>)(_source0_boxed0) -> {
        agency.open.ahp.ConfluxContract.Option<__V> _source0 = ((agency.open.ahp.ConfluxContract.Option<__V>)(java.lang.Object)(_source0_boxed0));
        if (_source0.is_None()) {
          return agency.open.ahp.ConfluxPrune.RouteOp.<__V>create_NoOp(_td___V);
        } else {
          __V _3___mcc_h0 = ((agency.open.ahp.ConfluxContract.Option_Some<__V>)_source0)._value;
          __V _4_v = _3___mcc_h0;
          return agency.open.ahp.ConfluxPrune.RouteOp.<__V>create_Install(_td___V, ((__V)(java.lang.Object)((_0_f).apply(_4_v))));
        }
      }).apply(_1_cur);
    });}).apply(f);
  }
  public static <__K, __V> dafny.DafnySequence<? extends __V> SeqUpdateByIdWhere(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, java.util.function.Function<__V, Boolean> pred, dafny.DafnySequence<? extends __V> xs, __K k, java.util.function.Function<__V, __V> f)
  {
    dafny.DafnySequence<? extends __V> _0___accumulator = dafny.DafnySequence.<__V> empty(_td___V);
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((xs).length())).signum() == 0) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> empty(_td___V));
      } else if ((java.util.Objects.equals(((__K)(java.lang.Object)((keyOf).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))), k)) && (((boolean)(java.lang.Object)((pred).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))))) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V>concatenate(dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((f).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO))))))))), (xs).drop(java.math.BigInteger.ONE)));
      } else {
        _0___accumulator = dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))));
        java.util.function.Function<__V, __K> _in0 = keyOf;
        java.util.function.Function<__V, Boolean> _in1 = pred;
        dafny.DafnySequence<? extends __V> _in2 = (xs).drop(java.math.BigInteger.ONE);
        __K _in3 = k;
        java.util.function.Function<__V, __V> _in4 = f;
        keyOf = _in0;
        pred = _in1;
        xs = _in2;
        k = _in3;
        f = _in4;
        continue TAIL_CALL_START;
      }
    }
  }
  public static <__V> dafny.Function2<agency.open.ahp.ConfluxContract.Option<__V>, __V, agency.open.ahp.ConfluxPrune.RouteOp<__V>> UpdateWhereOp(dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, Boolean> pred, java.util.function.Function<__V, __V> f)
  {
    return ((dafny.Function2<java.util.function.Function<__V, Boolean>, java.util.function.Function<__V, __V>, dafny.Function2<agency.open.ahp.ConfluxContract.Option<__V>, __V, agency.open.ahp.ConfluxPrune.RouteOp<__V>>>)(_0_pred, _1_f) -> {return ((dafny.Function2<agency.open.ahp.ConfluxContract.Option<__V>, __V, agency.open.ahp.ConfluxPrune.RouteOp<__V>>)(_2_cur_boxed0, _3_a_boxed0) -> {
      agency.open.ahp.ConfluxContract.Option<__V> _2_cur = ((agency.open.ahp.ConfluxContract.Option<__V>)(java.lang.Object)(_2_cur_boxed0));
      __V _3_a = ((__V)(java.lang.Object)(_3_a_boxed0));
      return ((java.util.function.Function<agency.open.ahp.ConfluxContract.Option<__V>, agency.open.ahp.ConfluxPrune.RouteOp<__V>>)(_source0_boxed0) -> {
        agency.open.ahp.ConfluxContract.Option<__V> _source0 = ((agency.open.ahp.ConfluxContract.Option<__V>)(java.lang.Object)(_source0_boxed0));
        if (_source0.is_None()) {
          return agency.open.ahp.ConfluxPrune.RouteOp.<__V>create_NoOp(_td___V);
        } else {
          __V _4___mcc_h0 = ((agency.open.ahp.ConfluxContract.Option_Some<__V>)_source0)._value;
          __V _5_v = _4___mcc_h0;
          if (((boolean)(java.lang.Object)((_0_pred).apply(_5_v)))) {
            return agency.open.ahp.ConfluxPrune.RouteOp.<__V>create_Install(_td___V, ((__V)(java.lang.Object)((_1_f).apply(_5_v))));
          } else {
            return agency.open.ahp.ConfluxPrune.RouteOp.<__V>create_NoOp(_td___V);
          }
        }
      }).apply(_2_cur);
    });}).apply(pred, f);
  }
  public static <__K, __V> dafny.DafnySequence<? extends __V> SeqUpdateAllWhere(dafny.TypeDescriptor<__K> _td___K, dafny.TypeDescriptor<__V> _td___V, java.util.function.Function<__V, __K> keyOf, java.util.function.Function<__V, Boolean> pred, dafny.DafnySequence<? extends __V> xs, __K k, java.util.function.Function<__V, __V> f)
  {
    dafny.DafnySequence<? extends __V> _0___accumulator = dafny.DafnySequence.<__V> empty(_td___V);
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((xs).length())).signum() == 0) {
        return dafny.DafnySequence.<__V>concatenate(_0___accumulator, dafny.DafnySequence.<__V> empty(_td___V));
      } else {
        _0___accumulator = dafny.DafnySequence.<__V>concatenate(_0___accumulator, (((java.util.Objects.equals(((__K)(java.lang.Object)((keyOf).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))), k)) && (((boolean)(java.lang.Object)((pred).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))))) ? (dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((f).apply(((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))))) : (dafny.DafnySequence.<__V> of(_td___V, ((__V)(java.lang.Object)((xs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))))));
        java.util.function.Function<__V, __K> _in0 = keyOf;
        java.util.function.Function<__V, Boolean> _in1 = pred;
        dafny.DafnySequence<? extends __V> _in2 = (xs).drop(java.math.BigInteger.ONE);
        __K _in3 = k;
        java.util.function.Function<__V, __V> _in4 = f;
        keyOf = _in0;
        pred = _in1;
        xs = _in2;
        k = _in3;
        f = _in4;
        continue TAIL_CALL_START;
      }
    }
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.ConfluxSeqRoute._default";
  }
}
