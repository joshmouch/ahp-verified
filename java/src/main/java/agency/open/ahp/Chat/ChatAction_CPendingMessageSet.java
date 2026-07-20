// Class ChatAction_CPendingMessageSet
// Dafny class ChatAction_CPendingMessageSet compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CPendingMessageSet extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _kind;
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public agency.open.ahp.ConfluxCodec.Json _message;
  public ChatAction_CPendingMessageSet (dafny.DafnySequence<? extends dafny.CodePoint> kind, dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.ConfluxCodec.Json message) {
    super();
    this._kind = kind;
    this._id = id;
    this._message = message;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CPendingMessageSet o = (ChatAction_CPendingMessageSet)other;
    return true && java.util.Objects.equals(this._kind, o._kind) && java.util.Objects.equals(this._id, o._id) && java.util.Objects.equals(this._message, o._message);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 25;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._kind);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._message);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CPendingMessageSet");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._kind));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._message));
    s.append(")");
    return s.toString();
  }
}
