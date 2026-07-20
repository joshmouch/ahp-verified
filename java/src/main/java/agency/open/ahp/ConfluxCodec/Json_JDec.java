// Class Json_JDec
// Dafny class Json_JDec compiled into Java
package agency.open.ahp.ConfluxCodec;

@SuppressWarnings({"unchecked", "deprecation"})
public class Json_JDec extends Json {
  public java.math.BigInteger _mantissa;
  public java.math.BigInteger _exp;
  public Json_JDec (java.math.BigInteger mantissa, java.math.BigInteger exp) {
    super();
    this._mantissa = mantissa;
    this._exp = exp;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Json_JDec o = (Json_JDec)other;
    return true && java.util.Objects.equals(this._mantissa, o._mantissa) && java.util.Objects.equals(this._exp, o._exp);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 3;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._mantissa);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._exp);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("ConfluxCodec.Json.JDec");
    s.append("(");
    s.append(dafny.Helpers.toString(this._mantissa));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._exp));
    s.append(")");
    return s.toString();
  }
}
