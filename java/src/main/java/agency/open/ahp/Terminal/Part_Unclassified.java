// Class Part_Unclassified
// Dafny class Part_Unclassified compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class Part_Unclassified extends Part {
  public dafny.DafnySequence<? extends dafny.CodePoint> _value;
  public Part_Unclassified (dafny.DafnySequence<? extends dafny.CodePoint> value) {
    super();
    this._value = value;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Part_Unclassified o = (Part_Unclassified)other;
    return true && java.util.Objects.equals(this._value, o._value);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._value);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.Part.Unclassified");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._value));
    s.append(")");
    return s.toString();
  }
}
