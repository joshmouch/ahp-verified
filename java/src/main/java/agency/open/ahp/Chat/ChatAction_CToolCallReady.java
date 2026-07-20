// Class ChatAction_CToolCallReady
// Dafny class ChatAction_CToolCallReady compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CToolCallReady extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _toolCallId;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _invocationMessage;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _toolInput;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _confirmationTitle;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _riskAssessment;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _edits;
  public agency.open.ahp.AhpSkeleton.Option<Boolean> _editable;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _options;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _confirmed;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _meta;
  public ChatAction_CToolCallReady (dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> invocationMessage, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> toolInput, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> confirmationTitle, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> riskAssessment, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> edits, agency.open.ahp.AhpSkeleton.Option<Boolean> editable, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> options, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> confirmed, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    super();
    this._turnId = turnId;
    this._toolCallId = toolCallId;
    this._invocationMessage = invocationMessage;
    this._toolInput = toolInput;
    this._confirmationTitle = confirmationTitle;
    this._riskAssessment = riskAssessment;
    this._edits = edits;
    this._editable = editable;
    this._options = options;
    this._confirmed = confirmed;
    this._meta = meta;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CToolCallReady o = (ChatAction_CToolCallReady)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._toolCallId, o._toolCallId) && java.util.Objects.equals(this._invocationMessage, o._invocationMessage) && java.util.Objects.equals(this._toolInput, o._toolInput) && java.util.Objects.equals(this._confirmationTitle, o._confirmationTitle) && java.util.Objects.equals(this._riskAssessment, o._riskAssessment) && java.util.Objects.equals(this._edits, o._edits) && java.util.Objects.equals(this._editable, o._editable) && java.util.Objects.equals(this._options, o._options) && java.util.Objects.equals(this._confirmed, o._confirmed) && java.util.Objects.equals(this._meta, o._meta);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 11;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._toolCallId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._invocationMessage);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._toolInput);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._confirmationTitle);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._riskAssessment);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._edits);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._editable);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._options);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._confirmed);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._meta);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CToolCallReady");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._toolCallId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._invocationMessage));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._toolInput));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._confirmationTitle));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._riskAssessment));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._edits));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._editable));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._options));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._confirmed));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._meta));
    s.append(")");
    return s.toString();
  }
}
