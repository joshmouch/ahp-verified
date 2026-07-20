// Class ChatAction_CInputCompleted
// Dafny class ChatAction_CInputCompleted compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CInputCompleted extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _requestId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _response;
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _answers;
  public ChatAction_CInputCompleted (dafny.DafnySequence<? extends dafny.CodePoint> requestId, dafny.DafnySequence<? extends dafny.CodePoint> response, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> answers) {
    super();
    this._requestId = requestId;
    this._response = response;
    this._answers = answers;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CInputCompleted o = (ChatAction_CInputCompleted)other;
    return true && java.util.Objects.equals(this._requestId, o._requestId) && java.util.Objects.equals(this._response, o._response) && java.util.Objects.equals(this._answers, o._answers);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 24;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._requestId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._response);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._answers);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CInputCompleted");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._requestId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._response));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._answers));
    s.append(")");
    return s.toString();
  }
}
