// Class Json_JNum
// Dafny class Json_JNum compiled into Java
package agency.open.ahp.ConfluxCodec;

@SuppressWarnings({"unchecked", "deprecation"})
public class Json_JNum extends Json {
  public java.math.BigInteger _n;
  public Json_JNum (java.math.BigInteger n) {
    super();
    this._n = n;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Json_JNum o = (Json_JNum)other;
    return true && java.util.Objects.equals(this._n, o._n);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 2;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._n);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("ConfluxCodec.Json.JNum");
    s.append("(");
    s.append(dafny.Helpers.toString(this._n));
    s.append(")");
    return s.toString();
  }
}
