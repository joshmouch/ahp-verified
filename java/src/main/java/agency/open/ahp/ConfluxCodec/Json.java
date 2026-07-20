// Class Json
// Dafny class Json compiled into Java
package agency.open.ahp.ConfluxCodec;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class Json {
  public Json() {
  }
  private static final dafny.TypeDescriptor<Json> _TYPE = dafny.TypeDescriptor.<Json>referenceWithInitializer(Json.class, () -> Json.Default());
  public static dafny.TypeDescriptor<Json> _typeDescriptor() {
    return (dafny.TypeDescriptor<Json>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final Json theDefault = Json.create_JNull();
  public static Json Default() {
    return theDefault;
  }
  public static Json create_JNull() {
    return new Json_JNull();
  }
  public static Json create_JBool(boolean b) {
    return new Json_JBool(b);
  }
  public static Json create_JNum(java.math.BigInteger n) {
    return new Json_JNum(n);
  }
  public static Json create_JDec(java.math.BigInteger mantissa, java.math.BigInteger exp) {
    return new Json_JDec(mantissa, exp);
  }
  public static Json create_JStr(dafny.DafnySequence<? extends dafny.CodePoint> s) {
    return new Json_JStr(s);
  }
  public static Json create_JArr(dafny.DafnySequence<? extends Json> elems) {
    return new Json_JArr(elems);
  }
  public static Json create_JObj(dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends Json> fields) {
    return new Json_JObj(fields);
  }
  public boolean is_JNull() { return this instanceof Json_JNull; }
  public boolean is_JBool() { return this instanceof Json_JBool; }
  public boolean is_JNum() { return this instanceof Json_JNum; }
  public boolean is_JDec() { return this instanceof Json_JDec; }
  public boolean is_JStr() { return this instanceof Json_JStr; }
  public boolean is_JArr() { return this instanceof Json_JArr; }
  public boolean is_JObj() { return this instanceof Json_JObj; }
  public boolean dtor_b() {
    Json d = this;
    return ((Json_JBool)d)._b;
  }
  public java.math.BigInteger dtor_n() {
    Json d = this;
    return ((Json_JNum)d)._n;
  }
  public java.math.BigInteger dtor_mantissa() {
    Json d = this;
    return ((Json_JDec)d)._mantissa;
  }
  public java.math.BigInteger dtor_exp() {
    Json d = this;
    return ((Json_JDec)d)._exp;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_s() {
    Json d = this;
    return ((Json_JStr)d)._s;
  }
  public dafny.DafnySequence<? extends Json> dtor_elems() {
    Json d = this;
    return ((Json_JArr)d)._elems;
  }
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends Json> dtor_fields() {
    Json d = this;
    return ((Json_JObj)d)._fields;
  }
}
