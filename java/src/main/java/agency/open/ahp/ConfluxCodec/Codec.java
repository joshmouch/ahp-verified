// Class Codec<T>
// Dafny class Codec<T> compiled into Java
package agency.open.ahp.ConfluxCodec;

@SuppressWarnings({"unchecked", "deprecation"})
public class Codec<T> {
  public java.util.function.Function<T, Json> _encode;
  public java.util.function.Function<Json, Option<T>> _decode;
  protected dafny.TypeDescriptor<T> _td_T;
  public Codec (dafny.TypeDescriptor<T> _td_T, java.util.function.Function<T, Json> encode, java.util.function.Function<Json, Option<T>> decode) {
    this._td_T = _td_T;
    this._encode = encode;
    this._decode = decode;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Codec<T> o = (Codec<T>)other;
    return true && java.util.Objects.equals(this._encode, o._encode) && java.util.Objects.equals(this._decode, o._decode);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._encode);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._decode);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("ConfluxCodec.Codec.Codec");
    s.append("(");
    s.append(dafny.Helpers.toString(this._encode));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._decode));
    s.append(")");
    return s.toString();
  }
  public static <T> dafny.TypeDescriptor<Codec<T>> _typeDescriptor(dafny.TypeDescriptor<T> _td_T) {
    return (dafny.TypeDescriptor<Codec<T>>) (dafny.TypeDescriptor<?>) dafny.TypeDescriptor.<Codec<T>>referenceWithInitializer(Codec.class, () -> Codec.<T>Default(_td_T));
  }

  public static <T> Codec<T> Default(dafny.TypeDescriptor<T> _td_T) {
    return Codec.<T>create(_td_T, ((T x21) -> Json.Default()), ((Json x22) -> Option.<T>Default(_td_T)));
  }
  public static <T> Codec<T> create(dafny.TypeDescriptor<T> _td_T, java.util.function.Function<T, Json> encode, java.util.function.Function<Json, Option<T>> decode) {
    return new Codec<T>(_td_T, encode, decode);
  }
  public static <T> Codec<T> create_Codec(dafny.TypeDescriptor<T> _td_T, java.util.function.Function<T, Json> encode, java.util.function.Function<Json, Option<T>> decode) {
    return create(_td_T, encode, decode);
  }
  public boolean is_Codec() { return true; }
  public java.util.function.Function<T, Json> dtor_encode() {
    return this._encode;
  }
  public java.util.function.Function<Json, Option<T>> dtor_decode() {
    return this._decode;
  }
}
