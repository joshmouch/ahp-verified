// Class ChatAction_CToolCallConfirmed
// Dafny class ChatAction_CToolCallConfirmed compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CToolCallConfirmed extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _toolCallId;
  public boolean _approved;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _confirmedVal;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _reason;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _reasonMessage;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _userSuggestion;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _editedToolInput;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _selectedOptionId;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _meta;
  public ChatAction_CToolCallConfirmed (dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, boolean approved, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> confirmedVal, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> reason, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> reasonMessage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> userSuggestion, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> editedToolInput, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> selectedOptionId, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    super();
    this._turnId = turnId;
    this._toolCallId = toolCallId;
    this._approved = approved;
    this._confirmedVal = confirmedVal;
    this._reason = reason;
    this._reasonMessage = reasonMessage;
    this._userSuggestion = userSuggestion;
    this._editedToolInput = editedToolInput;
    this._selectedOptionId = selectedOptionId;
    this._meta = meta;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CToolCallConfirmed o = (ChatAction_CToolCallConfirmed)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._toolCallId, o._toolCallId) && this._approved == o._approved && java.util.Objects.equals(this._confirmedVal, o._confirmedVal) && java.util.Objects.equals(this._reason, o._reason) && java.util.Objects.equals(this._reasonMessage, o._reasonMessage) && java.util.Objects.equals(this._userSuggestion, o._userSuggestion) && java.util.Objects.equals(this._editedToolInput, o._editedToolInput) && java.util.Objects.equals(this._selectedOptionId, o._selectedOptionId) && java.util.Objects.equals(this._meta, o._meta);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 12;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._toolCallId);
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._approved);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._confirmedVal);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._reason);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._reasonMessage);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._userSuggestion);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._editedToolInput);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._selectedOptionId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._meta);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CToolCallConfirmed");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._toolCallId));
    s.append(", ");
    s.append(this._approved);
    s.append(", ");
    s.append(dafny.Helpers.toString(this._confirmedVal));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._reason));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._reasonMessage));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._userSuggestion));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._editedToolInput));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._selectedOptionId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._meta));
    s.append(")");
    return s.toString();
  }
}
