// Class ChatAction_CPendingMessageRemoved
// Dafny class ChatAction_CPendingMessageRemoved compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CPendingMessageRemoved extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _kind;
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public ChatAction_CPendingMessageRemoved (dafny.DafnySequence<? extends dafny.CodePoint> kind, dafny.DafnySequence<? extends dafny.CodePoint> id) {
    super();
    this._kind = kind;
    this._id = id;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CPendingMessageRemoved o = (ChatAction_CPendingMessageRemoved)other;
    return true && java.util.Objects.equals(this._kind, o._kind) && java.util.Objects.equals(this._id, o._id);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 26;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._kind);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CPendingMessageRemoved");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._kind));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(")");
    return s.toString();
  }
}
