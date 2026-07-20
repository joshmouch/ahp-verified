// Class ChatAction_CTurnComplete
// Dafny class ChatAction_CTurnComplete compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CTurnComplete extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public ChatAction_CTurnComplete (dafny.DafnySequence<? extends dafny.CodePoint> turnId) {
    super();
    this._turnId = turnId;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CTurnComplete o = (ChatAction_CTurnComplete)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 3;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CTurnComplete");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(")");
    return s.toString();
  }
}
