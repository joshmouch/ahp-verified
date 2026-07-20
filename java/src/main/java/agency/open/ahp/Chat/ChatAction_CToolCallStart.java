// Class ChatAction_CToolCallStart
// Dafny class ChatAction_CToolCallStart compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CToolCallStart extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _toolCallId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _toolName;
  public dafny.DafnySequence<? extends dafny.CodePoint> _displayName;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _intention;
  public agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> _contributor;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _meta;
  public ChatAction_CToolCallStart (dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, dafny.DafnySequence<? extends dafny.CodePoint> toolName, dafny.DafnySequence<? extends dafny.CodePoint> displayName, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> intention, agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> contributor, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    super();
    this._turnId = turnId;
    this._toolCallId = toolCallId;
    this._toolName = toolName;
    this._displayName = displayName;
    this._intention = intention;
    this._contributor = contributor;
    this._meta = meta;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CToolCallStart o = (ChatAction_CToolCallStart)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._toolCallId, o._toolCallId) && java.util.Objects.equals(this._toolName, o._toolName) && java.util.Objects.equals(this._displayName, o._displayName) && java.util.Objects.equals(this._intention, o._intention) && java.util.Objects.equals(this._contributor, o._contributor) && java.util.Objects.equals(this._meta, o._meta);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 9;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._toolCallId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._toolName);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._displayName);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._intention);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._contributor);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._meta);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CToolCallStart");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._toolCallId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._toolName));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._displayName));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._intention));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._contributor));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._meta));
    s.append(")");
    return s.toString();
  }
}
