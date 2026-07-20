// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.ConfluxCodec;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static Json Field(Json j, dafny.DafnySequence<? extends dafny.CodePoint> key)
  {
    if (((j).is_JObj()) && (((j).dtor_fields()).<dafny.DafnySequence<? extends dafny.CodePoint>>contains(key))) {
      return ((Json)(java.lang.Object)(((j).dtor_fields()).get(key)));
    } else {
      return Json.create_JNull();
    }
  }
  public static java.math.BigInteger AsInt(Json j) {
    if ((j).is_JNum()) {
      return (j).dtor_n();
    } else {
      return java.math.BigInteger.ZERO;
    }
  }
  public static dafny.DafnySequence<? extends dafny.CodePoint> AsStr(Json j) {
    if ((j).is_JStr()) {
      return (j).dtor_s();
    } else {
      return dafny.DafnySequence.asUnicodeString("");
    }
  }
  public static java.math.BigInteger AsNat(Json j) {
    java.math.BigInteger _0_x = __default.AsInt(j);
    if ((_0_x).signum() != -1) {
      return _0_x;
    } else {
      return java.math.BigInteger.ZERO;
    }
  }
  public static Codec<Json> JsonCodec() {
    return Codec.<Json>create(Json._typeDescriptor(), ((java.util.function.Function<Json, Json>)(_0_j_boxed0) -> {
  Json _0_j = ((Json)(java.lang.Object)(_0_j_boxed0));
  return _0_j;
}), ((java.util.function.Function<Json, Option<Json>>)(_1_j_boxed0) -> {
  Json _1_j = ((Json)(java.lang.Object)(_1_j_boxed0));
  return Option.<Json>create_Some(Json._typeDescriptor(), _1_j);
}));
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.ConfluxCodec._default";
  }
}
