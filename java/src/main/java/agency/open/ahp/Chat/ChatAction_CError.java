// Class ChatAction_CError
// Dafny class ChatAction_CError compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CError extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public agency.open.ahp.ConfluxCodec.Json _err;
  public ChatAction_CError (dafny.DafnySequence<? extends dafny.CodePoint> turnId, agency.open.ahp.ConfluxCodec.Json err) {
    super();
    this._turnId = turnId;
    this._err = err;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CError o = (ChatAction_CError)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._err, o._err);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 6;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._err);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CError");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._err));
    s.append(")");
    return s.toString();
  }
}
