// Class ChatAction_CToolCallComplete
// Dafny class ChatAction_CToolCallComplete compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CToolCallComplete extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _toolCallId;
  public boolean _success;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pastTenseMessage;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _resultContent;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _structuredContent;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _error;
  public boolean _requiresResultConfirmation;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _meta;
  public ChatAction_CToolCallComplete (dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, boolean success, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> pastTenseMessage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> resultContent, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> structuredContent, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error, boolean requiresResultConfirmation, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    super();
    this._turnId = turnId;
    this._toolCallId = toolCallId;
    this._success = success;
    this._pastTenseMessage = pastTenseMessage;
    this._resultContent = resultContent;
    this._structuredContent = structuredContent;
    this._error = error;
    this._requiresResultConfirmation = requiresResultConfirmation;
    this._meta = meta;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CToolCallComplete o = (ChatAction_CToolCallComplete)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._toolCallId, o._toolCallId) && this._success == o._success && java.util.Objects.equals(this._pastTenseMessage, o._pastTenseMessage) && java.util.Objects.equals(this._resultContent, o._resultContent) && java.util.Objects.equals(this._structuredContent, o._structuredContent) && java.util.Objects.equals(this._error, o._error) && this._requiresResultConfirmation == o._requiresResultConfirmation && java.util.Objects.equals(this._meta, o._meta);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 15;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._toolCallId);
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._success);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._pastTenseMessage);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._resultContent);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._structuredContent);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._error);
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._requiresResultConfirmation);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._meta);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CToolCallComplete");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._toolCallId));
    s.append(", ");
    s.append(this._success);
    s.append(", ");
    s.append(dafny.Helpers.toString(this._pastTenseMessage));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._resultContent));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._structuredContent));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._error));
    s.append(", ");
    s.append(this._requiresResultConfirmation);
    s.append(", ");
    s.append(dafny.Helpers.toString(this._meta));
    s.append(")");
    return s.toString();
  }
}
