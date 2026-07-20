// Class Part_PToolCall
// Dafny class Part_PToolCall compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class Part_PToolCall extends Part {
  public ToolCall _tc;
  public Part_PToolCall (ToolCall tc) {
    super();
    this._tc = tc;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Part_PToolCall o = (Part_PToolCall)other;
    return true && java.util.Objects.equals(this._tc, o._tc);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 2;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._tc);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.Part.PToolCall");
    s.append("(");
    s.append(dafny.Helpers.toString(this._tc));
    s.append(")");
    return s.toString();
  }
}
