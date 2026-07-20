// Class ChatAction_CResponsePart
// Dafny class ChatAction_CResponsePart compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CResponsePart extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public Part _part;
  public ChatAction_CResponsePart (dafny.DafnySequence<? extends dafny.CodePoint> turnId, Part part) {
    super();
    this._turnId = turnId;
    this._part = part;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CResponsePart o = (ChatAction_CResponsePart)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._part, o._part);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 7;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._part);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CResponsePart");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._part));
    s.append(")");
    return s.toString();
  }
}
