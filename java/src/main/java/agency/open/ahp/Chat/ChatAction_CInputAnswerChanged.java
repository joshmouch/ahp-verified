// Class ChatAction_CInputAnswerChanged
// Dafny class ChatAction_CInputAnswerChanged compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CInputAnswerChanged extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _requestId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _questionId;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _answer;
  public ChatAction_CInputAnswerChanged (dafny.DafnySequence<? extends dafny.CodePoint> requestId, dafny.DafnySequence<? extends dafny.CodePoint> questionId, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> answer) {
    super();
    this._requestId = requestId;
    this._questionId = questionId;
    this._answer = answer;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CInputAnswerChanged o = (ChatAction_CInputAnswerChanged)other;
    return true && java.util.Objects.equals(this._requestId, o._requestId) && java.util.Objects.equals(this._questionId, o._questionId) && java.util.Objects.equals(this._answer, o._answer);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 23;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._requestId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._questionId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._answer);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CInputAnswerChanged");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._requestId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._questionId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._answer));
    s.append(")");
    return s.toString();
  }
}
