// Class Json_JObj
// Dafny class Json_JObj compiled into Java
package agency.open.ahp.ConfluxCodec;

@SuppressWarnings({"unchecked", "deprecation"})
public class Json_JObj extends Json {
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends Json> _fields;
  public Json_JObj (dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends Json> fields) {
    super();
    this._fields = fields;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Json_JObj o = (Json_JObj)other;
    return true && java.util.Objects.equals(this._fields, o._fields);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 6;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._fields);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("ConfluxCodec.Json.JObj");
    s.append("(");
    s.append(dafny.Helpers.toString(this._fields));
    s.append(")");
    return s.toString();
  }
}
