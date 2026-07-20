// Class ChatAction_CToolCallAuthRequired
// Dafny class ChatAction_CToolCallAuthRequired compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CToolCallAuthRequired extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _toolCallId;
  public agency.open.ahp.ConfluxCodec.Json _auth;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _meta;
  public ChatAction_CToolCallAuthRequired (dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, agency.open.ahp.ConfluxCodec.Json auth, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    super();
    this._turnId = turnId;
    this._toolCallId = toolCallId;
    this._auth = auth;
    this._meta = meta;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CToolCallAuthRequired o = (ChatAction_CToolCallAuthRequired)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._toolCallId, o._toolCallId) && java.util.Objects.equals(this._auth, o._auth) && java.util.Objects.equals(this._meta, o._meta);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 13;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._toolCallId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._auth);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._meta);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CToolCallAuthRequired");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._toolCallId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._auth));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._meta));
    s.append(")");
    return s.toString();
  }
}
