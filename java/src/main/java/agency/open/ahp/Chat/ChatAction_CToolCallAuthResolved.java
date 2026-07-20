// Class ChatAction_CToolCallAuthResolved
// Dafny class ChatAction_CToolCallAuthResolved compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CToolCallAuthResolved extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _toolCallId;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _meta;
  public ChatAction_CToolCallAuthResolved (dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    super();
    this._turnId = turnId;
    this._toolCallId = toolCallId;
    this._meta = meta;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CToolCallAuthResolved o = (ChatAction_CToolCallAuthResolved)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._toolCallId, o._toolCallId) && java.util.Objects.equals(this._meta, o._meta);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 14;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._toolCallId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._meta);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CToolCallAuthResolved");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._toolCallId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._meta));
    s.append(")");
    return s.toString();
  }
}
