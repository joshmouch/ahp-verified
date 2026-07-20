// Class ChatAction_CTruncated
// Dafny class ChatAction_CTruncated compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CTruncated extends ChatAction {
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _upToTurnId;
  public ChatAction_CTruncated (agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> upToTurnId) {
    super();
    this._upToTurnId = upToTurnId;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CTruncated o = (ChatAction_CTruncated)other;
    return true && java.util.Objects.equals(this._upToTurnId, o._upToTurnId);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 19;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._upToTurnId);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CTruncated");
    s.append("(");
    s.append(dafny.Helpers.toString(this._upToTurnId));
    s.append(")");
    return s.toString();
  }
}
