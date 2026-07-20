// Class ToolCall
// Dafny class ToolCall compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ToolCall {
  public dafny.DafnySequence<? extends dafny.CodePoint> _toolCallId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _toolName;
  public dafny.DafnySequence<? extends dafny.CodePoint> _displayName;
  public dafny.DafnySequence<? extends dafny.CodePoint> _status;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _intention;
  public agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> _contributor;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _meta;
  public dafny.DafnySequence<? extends dafny.CodePoint> _invocationMessage;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _toolInput;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _confirmationTitle;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _riskAssessment;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _edits;
  public agency.open.ahp.AhpSkeleton.Option<Boolean> _editable;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _options;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _confirmed;
  public agency.open.ahp.AhpSkeleton.Option<Boolean> _success;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pastTenseMessage;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _reason;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _reasonMessage;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _userSuggestion;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _selectedOption;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _content;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _structuredContent;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _error;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _auth;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _partialInput;
  public ToolCall (dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, dafny.DafnySequence<? extends dafny.CodePoint> toolName, dafny.DafnySequence<? extends dafny.CodePoint> displayName, dafny.DafnySequence<? extends dafny.CodePoint> status, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> intention, agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> contributor, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta, dafny.DafnySequence<? extends dafny.CodePoint> invocationMessage, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> toolInput, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> confirmationTitle, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> riskAssessment, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> edits, agency.open.ahp.AhpSkeleton.Option<Boolean> editable, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> options, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> confirmed, agency.open.ahp.AhpSkeleton.Option<Boolean> success, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> pastTenseMessage, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> reason, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> reasonMessage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> userSuggestion, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> selectedOption, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> content, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> structuredContent, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> auth, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> partialInput) {
    this._toolCallId = toolCallId;
    this._toolName = toolName;
    this._displayName = displayName;
    this._status = status;
    this._intention = intention;
    this._contributor = contributor;
    this._meta = meta;
    this._invocationMessage = invocationMessage;
    this._toolInput = toolInput;
    this._confirmationTitle = confirmationTitle;
    this._riskAssessment = riskAssessment;
    this._edits = edits;
    this._editable = editable;
    this._options = options;
    this._confirmed = confirmed;
    this._success = success;
    this._pastTenseMessage = pastTenseMessage;
    this._reason = reason;
    this._reasonMessage = reasonMessage;
    this._userSuggestion = userSuggestion;
    this._selectedOption = selectedOption;
    this._content = content;
    this._structuredContent = structuredContent;
    this._error = error;
    this._auth = auth;
    this._partialInput = partialInput;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ToolCall o = (ToolCall)other;
    return true && java.util.Objects.equals(this._toolCallId, o._toolCallId) && java.util.Objects.equals(this._toolName, o._toolName) && java.util.Objects.equals(this._displayName, o._displayName) && java.util.Objects.equals(this._status, o._status) && java.util.Objects.equals(this._intention, o._intention) && java.util.Objects.equals(this._contributor, o._contributor) && java.util.Objects.equals(this._meta, o._meta) && java.util.Objects.equals(this._invocationMessage, o._invocationMessage) && java.util.Objects.equals(this._toolInput, o._toolInput) && java.util.Objects.equals(this._confirmationTitle, o._confirmationTitle) && java.util.Objects.equals(this._riskAssessment, o._riskAssessment) && java.util.Objects.equals(this._edits, o._edits) && java.util.Objects.equals(this._editable, o._editable) && java.util.Objects.equals(this._options, o._options) && java.util.Objects.equals(this._confirmed, o._confirmed) && java.util.Objects.equals(this._success, o._success) && java.util.Objects.equals(this._pastTenseMessage, o._pastTenseMessage) && java.util.Objects.equals(this._reason, o._reason) && java.util.Objects.equals(this._reasonMessage, o._reasonMessage) && java.util.Objects.equals(this._userSuggestion, o._userSuggestion) && java.util.Objects.equals(this._selectedOption, o._selectedOption) && java.util.Objects.equals(this._content, o._content) && java.util.Objects.equals(this._structuredContent, o._structuredContent) && java.util.Objects.equals(this._error, o._error) && java.util.Objects.equals(this._auth, o._auth) && java.util.Objects.equals(this._partialInput, o._partialInput);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._toolCallId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._toolName);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._displayName);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._status);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._intention);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._contributor);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._meta);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._invocationMessage);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._toolInput);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._confirmationTitle);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._riskAssessment);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._edits);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._editable);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._options);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._confirmed);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._success);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._pastTenseMessage);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._reason);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._reasonMessage);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._userSuggestion);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._selectedOption);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._content);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._structuredContent);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._error);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._auth);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._partialInput);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ToolCall.ToolCall");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._toolCallId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._toolName));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._displayName));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._status));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._intention));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._contributor));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._meta));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._invocationMessage));
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
    s.append(dafny.Helpers.toString(this._success));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._pastTenseMessage));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._reason));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._reasonMessage));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._userSuggestion));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._selectedOption));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._content));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._structuredContent));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._error));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._auth));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._partialInput));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<ToolCall> _TYPE = dafny.TypeDescriptor.<ToolCall>referenceWithInitializer(ToolCall.class, () -> ToolCall.Default());
  public static dafny.TypeDescriptor<ToolCall> _typeDescriptor() {
    return (dafny.TypeDescriptor<ToolCall>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ToolCall theDefault = ToolCall.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>Default(ToolCallContributor._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>Default(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>Default(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<Boolean>Default(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)));
  public static ToolCall Default() {
    return theDefault;
  }
  public static ToolCall create(dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, dafny.DafnySequence<? extends dafny.CodePoint> toolName, dafny.DafnySequence<? extends dafny.CodePoint> displayName, dafny.DafnySequence<? extends dafny.CodePoint> status, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> intention, agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> contributor, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta, dafny.DafnySequence<? extends dafny.CodePoint> invocationMessage, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> toolInput, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> confirmationTitle, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> riskAssessment, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> edits, agency.open.ahp.AhpSkeleton.Option<Boolean> editable, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> options, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> confirmed, agency.open.ahp.AhpSkeleton.Option<Boolean> success, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> pastTenseMessage, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> reason, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> reasonMessage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> userSuggestion, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> selectedOption, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> content, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> structuredContent, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> auth, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> partialInput) {
    return new ToolCall(toolCallId, toolName, displayName, status, intention, contributor, meta, invocationMessage, toolInput, confirmationTitle, riskAssessment, edits, editable, options, confirmed, success, pastTenseMessage, reason, reasonMessage, userSuggestion, selectedOption, content, structuredContent, error, auth, partialInput);
  }
  public static ToolCall create_ToolCall(dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, dafny.DafnySequence<? extends dafny.CodePoint> toolName, dafny.DafnySequence<? extends dafny.CodePoint> displayName, dafny.DafnySequence<? extends dafny.CodePoint> status, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> intention, agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> contributor, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta, dafny.DafnySequence<? extends dafny.CodePoint> invocationMessage, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> toolInput, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> confirmationTitle, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> riskAssessment, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> edits, agency.open.ahp.AhpSkeleton.Option<Boolean> editable, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> options, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> confirmed, agency.open.ahp.AhpSkeleton.Option<Boolean> success, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> pastTenseMessage, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> reason, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> reasonMessage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> userSuggestion, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> selectedOption, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> content, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> structuredContent, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> auth, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> partialInput) {
    return create(toolCallId, toolName, displayName, status, intention, contributor, meta, invocationMessage, toolInput, confirmationTitle, riskAssessment, edits, editable, options, confirmed, success, pastTenseMessage, reason, reasonMessage, userSuggestion, selectedOption, content, structuredContent, error, auth, partialInput);
  }
  public boolean is_ToolCall() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_toolCallId() {
    return this._toolCallId;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_toolName() {
    return this._toolName;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_displayName() {
    return this._displayName;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_status() {
    return this._status;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_intention() {
    return this._intention;
  }
  public agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> dtor_contributor() {
    return this._contributor;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_meta() {
    return this._meta;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_invocationMessage() {
    return this._invocationMessage;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_toolInput() {
    return this._toolInput;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_confirmationTitle() {
    return this._confirmationTitle;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_riskAssessment() {
    return this._riskAssessment;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_edits() {
    return this._edits;
  }
  public agency.open.ahp.AhpSkeleton.Option<Boolean> dtor_editable() {
    return this._editable;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> dtor_options() {
    return this._options;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_confirmed() {
    return this._confirmed;
  }
  public agency.open.ahp.AhpSkeleton.Option<Boolean> dtor_success() {
    return this._success;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_pastTenseMessage() {
    return this._pastTenseMessage;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_reason() {
    return this._reason;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_reasonMessage() {
    return this._reasonMessage;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_userSuggestion() {
    return this._userSuggestion;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_selectedOption() {
    return this._selectedOption;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_content() {
    return this._content;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_structuredContent() {
    return this._structuredContent;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_error() {
    return this._error;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_auth() {
    return this._auth;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_partialInput() {
    return this._partialInput;
  }
}
