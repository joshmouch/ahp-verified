// Class Part_PReasoning
// Dafny class Part_PReasoning compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class Part_PReasoning extends Part {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public dafny.DafnySequence<? extends dafny.CodePoint> _content;
  public Part_PReasoning (dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> content) {
    super();
    this._id = id;
    this._content = content;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Part_PReasoning o = (Part_PReasoning)other;
    return true && java.util.Objects.equals(this._id, o._id) && java.util.Objects.equals(this._content, o._content);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._content);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.Part.PReasoning");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._content));
    s.append(")");
    return s.toString();
  }
}
