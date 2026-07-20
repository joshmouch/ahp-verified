// Class ChatAction_CReasoning
// Dafny class ChatAction_CReasoning compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CReasoning extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _partId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _content;
  public ChatAction_CReasoning (dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> partId, dafny.DafnySequence<? extends dafny.CodePoint> content) {
    super();
    this._turnId = turnId;
    this._partId = partId;
    this._content = content;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CReasoning o = (ChatAction_CReasoning)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._partId, o._partId) && java.util.Objects.equals(this._content, o._content);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 18;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._partId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._content);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CReasoning");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._partId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._content));
    s.append(")");
    return s.toString();
  }
}
