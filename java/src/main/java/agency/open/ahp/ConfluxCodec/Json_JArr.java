// Class Json_JArr
// Dafny class Json_JArr compiled into Java
package agency.open.ahp.ConfluxCodec;

@SuppressWarnings({"unchecked", "deprecation"})
public class Json_JArr extends Json {
  public dafny.DafnySequence<? extends Json> _elems;
  public Json_JArr (dafny.DafnySequence<? extends Json> elems) {
    super();
    this._elems = elems;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Json_JArr o = (Json_JArr)other;
    return true && java.util.Objects.equals(this._elems, o._elems);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 5;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._elems);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("ConfluxCodec.Json.JArr");
    s.append("(");
    s.append(dafny.Helpers.toString(this._elems));
    s.append(")");
    return s.toString();
  }
}
