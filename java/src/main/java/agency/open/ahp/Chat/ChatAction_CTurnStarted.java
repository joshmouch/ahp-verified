// Class ChatAction_CTurnStarted
// Dafny class ChatAction_CTurnStarted compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CTurnStarted extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public agency.open.ahp.ConfluxCodec.Json _message;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _queuedMessageId;
  public ChatAction_CTurnStarted (dafny.DafnySequence<? extends dafny.CodePoint> turnId, agency.open.ahp.ConfluxCodec.Json message, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> queuedMessageId) {
    super();
    this._turnId = turnId;
    this._message = message;
    this._queuedMessageId = queuedMessageId;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CTurnStarted o = (ChatAction_CTurnStarted)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._message, o._message) && java.util.Objects.equals(this._queuedMessageId, o._queuedMessageId);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 2;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._message);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._queuedMessageId);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CTurnStarted");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._message));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._queuedMessageId));
    s.append(")");
    return s.toString();
  }
}
