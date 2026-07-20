// Class ChatAction
// Dafny class ChatAction compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class ChatAction {
  public ChatAction() {
  }
  private static final dafny.TypeDescriptor<ChatAction> _TYPE = dafny.TypeDescriptor.<ChatAction>referenceWithInitializer(ChatAction.class, () -> ChatAction.Default());
  public static dafny.TypeDescriptor<ChatAction> _typeDescriptor() {
    return (dafny.TypeDescriptor<ChatAction>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ChatAction theDefault = ChatAction.create_CDraftChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)));
  public static ChatAction Default() {
    return theDefault;
  }
  public static ChatAction create_CDraftChanged(agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> draft) {
    return new ChatAction_CDraftChanged(draft);
  }
  public static ChatAction create_CActivityChanged(agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity) {
    return new ChatAction_CActivityChanged(activity);
  }
  public static ChatAction create_CTurnStarted(dafny.DafnySequence<? extends dafny.CodePoint> turnId, agency.open.ahp.ConfluxCodec.Json message, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> queuedMessageId) {
    return new ChatAction_CTurnStarted(turnId, message, queuedMessageId);
  }
  public static ChatAction create_CTurnComplete(dafny.DafnySequence<? extends dafny.CodePoint> turnId) {
    return new ChatAction_CTurnComplete(turnId);
  }
  public static ChatAction create_CTurnCancelled(dafny.DafnySequence<? extends dafny.CodePoint> turnId) {
    return new ChatAction_CTurnCancelled(turnId);
  }
  public static ChatAction create_CUsage(dafny.DafnySequence<? extends dafny.CodePoint> turnId, agency.open.ahp.ConfluxCodec.Json usage) {
    return new ChatAction_CUsage(turnId, usage);
  }
  public static ChatAction create_CError(dafny.DafnySequence<? extends dafny.CodePoint> turnId, agency.open.ahp.ConfluxCodec.Json err) {
    return new ChatAction_CError(turnId, err);
  }
  public static ChatAction create_CResponsePart(dafny.DafnySequence<? extends dafny.CodePoint> turnId, Part part) {
    return new ChatAction_CResponsePart(turnId, part);
  }
  public static ChatAction create_CDelta(dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> partId, dafny.DafnySequence<? extends dafny.CodePoint> content) {
    return new ChatAction_CDelta(turnId, partId, content);
  }
  public static ChatAction create_CToolCallStart(dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, dafny.DafnySequence<? extends dafny.CodePoint> toolName, dafny.DafnySequence<? extends dafny.CodePoint> displayName, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> intention, agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> contributor, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return new ChatAction_CToolCallStart(turnId, toolCallId, toolName, displayName, intention, contributor, meta);
  }
  public static ChatAction create_CToolCallDelta(dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, dafny.DafnySequence<? extends dafny.CodePoint> content, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> invocationMessage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return new ChatAction_CToolCallDelta(turnId, toolCallId, content, invocationMessage, meta);
  }
  public static ChatAction create_CToolCallReady(dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> invocationMessage, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> toolInput, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> confirmationTitle, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> riskAssessment, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> edits, agency.open.ahp.AhpSkeleton.Option<Boolean> editable, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> options, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> confirmed, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return new ChatAction_CToolCallReady(turnId, toolCallId, invocationMessage, toolInput, confirmationTitle, riskAssessment, edits, editable, options, confirmed, meta);
  }
  public static ChatAction create_CToolCallConfirmed(dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, boolean approved, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> confirmedVal, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> reason, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> reasonMessage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> userSuggestion, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> editedToolInput, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> selectedOptionId, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return new ChatAction_CToolCallConfirmed(turnId, toolCallId, approved, confirmedVal, reason, reasonMessage, userSuggestion, editedToolInput, selectedOptionId, meta);
  }
  public static ChatAction create_CToolCallAuthRequired(dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, agency.open.ahp.ConfluxCodec.Json auth, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return new ChatAction_CToolCallAuthRequired(turnId, toolCallId, auth, meta);
  }
  public static ChatAction create_CToolCallAuthResolved(dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return new ChatAction_CToolCallAuthResolved(turnId, toolCallId, meta);
  }
  public static ChatAction create_CToolCallComplete(dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, boolean success, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> pastTenseMessage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> resultContent, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> structuredContent, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error, boolean requiresResultConfirmation, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return new ChatAction_CToolCallComplete(turnId, toolCallId, success, pastTenseMessage, resultContent, structuredContent, error, requiresResultConfirmation, meta);
  }
  public static ChatAction create_CToolCallResultConfirmed(dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, boolean approved, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return new ChatAction_CToolCallResultConfirmed(turnId, toolCallId, approved, meta);
  }
  public static ChatAction create_CToolCallContentChanged(dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> toolCallId, agency.open.ahp.ConfluxCodec.Json newContent, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return new ChatAction_CToolCallContentChanged(turnId, toolCallId, newContent, meta);
  }
  public static ChatAction create_CReasoning(dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> partId, dafny.DafnySequence<? extends dafny.CodePoint> content) {
    return new ChatAction_CReasoning(turnId, partId, content);
  }
  public static ChatAction create_CTruncated(agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> upToTurnId) {
    return new ChatAction_CTruncated(upToTurnId);
  }
  public static ChatAction create_CQueuedReordered(dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> order) {
    return new ChatAction_CQueuedReordered(order);
  }
  public static ChatAction create_CTurnsLoaded(dafny.DafnySequence<? extends Turn> loaded, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> cursor) {
    return new ChatAction_CTurnsLoaded(loaded, cursor);
  }
  public static ChatAction create_CInputRequested(InputReq req) {
    return new ChatAction_CInputRequested(req);
  }
  public static ChatAction create_CInputAnswerChanged(dafny.DafnySequence<? extends dafny.CodePoint> requestId, dafny.DafnySequence<? extends dafny.CodePoint> questionId, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> answer) {
    return new ChatAction_CInputAnswerChanged(requestId, questionId, answer);
  }
  public static ChatAction create_CInputCompleted(dafny.DafnySequence<? extends dafny.CodePoint> requestId, dafny.DafnySequence<? extends dafny.CodePoint> response, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> answers) {
    return new ChatAction_CInputCompleted(requestId, response, answers);
  }
  public static ChatAction create_CPendingMessageSet(dafny.DafnySequence<? extends dafny.CodePoint> kind, dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.ConfluxCodec.Json message) {
    return new ChatAction_CPendingMessageSet(kind, id, message);
  }
  public static ChatAction create_CPendingMessageRemoved(dafny.DafnySequence<? extends dafny.CodePoint> kind, dafny.DafnySequence<? extends dafny.CodePoint> id) {
    return new ChatAction_CPendingMessageRemoved(kind, id);
  }
  public static ChatAction create_CUnknown(agency.open.ahp.ConfluxCodec.Json raw) {
    return new ChatAction_CUnknown(raw);
  }
  public boolean is_CDraftChanged() { return this instanceof ChatAction_CDraftChanged; }
  public boolean is_CActivityChanged() { return this instanceof ChatAction_CActivityChanged; }
  public boolean is_CTurnStarted() { return this instanceof ChatAction_CTurnStarted; }
  public boolean is_CTurnComplete() { return this instanceof ChatAction_CTurnComplete; }
  public boolean is_CTurnCancelled() { return this instanceof ChatAction_CTurnCancelled; }
  public boolean is_CUsage() { return this instanceof ChatAction_CUsage; }
  public boolean is_CError() { return this instanceof ChatAction_CError; }
  public boolean is_CResponsePart() { return this instanceof ChatAction_CResponsePart; }
  public boolean is_CDelta() { return this instanceof ChatAction_CDelta; }
  public boolean is_CToolCallStart() { return this instanceof ChatAction_CToolCallStart; }
  public boolean is_CToolCallDelta() { return this instanceof ChatAction_CToolCallDelta; }
  public boolean is_CToolCallReady() { return this instanceof ChatAction_CToolCallReady; }
  public boolean is_CToolCallConfirmed() { return this instanceof ChatAction_CToolCallConfirmed; }
  public boolean is_CToolCallAuthRequired() { return this instanceof ChatAction_CToolCallAuthRequired; }
  public boolean is_CToolCallAuthResolved() { return this instanceof ChatAction_CToolCallAuthResolved; }
  public boolean is_CToolCallComplete() { return this instanceof ChatAction_CToolCallComplete; }
  public boolean is_CToolCallResultConfirmed() { return this instanceof ChatAction_CToolCallResultConfirmed; }
  public boolean is_CToolCallContentChanged() { return this instanceof ChatAction_CToolCallContentChanged; }
  public boolean is_CReasoning() { return this instanceof ChatAction_CReasoning; }
  public boolean is_CTruncated() { return this instanceof ChatAction_CTruncated; }
  public boolean is_CQueuedReordered() { return this instanceof ChatAction_CQueuedReordered; }
  public boolean is_CTurnsLoaded() { return this instanceof ChatAction_CTurnsLoaded; }
  public boolean is_CInputRequested() { return this instanceof ChatAction_CInputRequested; }
  public boolean is_CInputAnswerChanged() { return this instanceof ChatAction_CInputAnswerChanged; }
  public boolean is_CInputCompleted() { return this instanceof ChatAction_CInputCompleted; }
  public boolean is_CPendingMessageSet() { return this instanceof ChatAction_CPendingMessageSet; }
  public boolean is_CPendingMessageRemoved() { return this instanceof ChatAction_CPendingMessageRemoved; }
  public boolean is_CUnknown() { return this instanceof ChatAction_CUnknown; }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_draft() {
    ChatAction d = this;
    return ((ChatAction_CDraftChanged)d)._draft;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_activity() {
    ChatAction d = this;
    return ((ChatAction_CActivityChanged)d)._activity;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_turnId() {
    ChatAction d = this;
    if (d instanceof ChatAction_CTurnStarted) { return ((ChatAction_CTurnStarted)d)._turnId; }
    if (d instanceof ChatAction_CTurnComplete) { return ((ChatAction_CTurnComplete)d)._turnId; }
    if (d instanceof ChatAction_CTurnCancelled) { return ((ChatAction_CTurnCancelled)d)._turnId; }
    if (d instanceof ChatAction_CUsage) { return ((ChatAction_CUsage)d)._turnId; }
    if (d instanceof ChatAction_CError) { return ((ChatAction_CError)d)._turnId; }
    if (d instanceof ChatAction_CResponsePart) { return ((ChatAction_CResponsePart)d)._turnId; }
    if (d instanceof ChatAction_CDelta) { return ((ChatAction_CDelta)d)._turnId; }
    if (d instanceof ChatAction_CToolCallStart) { return ((ChatAction_CToolCallStart)d)._turnId; }
    if (d instanceof ChatAction_CToolCallDelta) { return ((ChatAction_CToolCallDelta)d)._turnId; }
    if (d instanceof ChatAction_CToolCallReady) { return ((ChatAction_CToolCallReady)d)._turnId; }
    if (d instanceof ChatAction_CToolCallConfirmed) { return ((ChatAction_CToolCallConfirmed)d)._turnId; }
    if (d instanceof ChatAction_CToolCallAuthRequired) { return ((ChatAction_CToolCallAuthRequired)d)._turnId; }
    if (d instanceof ChatAction_CToolCallAuthResolved) { return ((ChatAction_CToolCallAuthResolved)d)._turnId; }
    if (d instanceof ChatAction_CToolCallComplete) { return ((ChatAction_CToolCallComplete)d)._turnId; }
    if (d instanceof ChatAction_CToolCallResultConfirmed) { return ((ChatAction_CToolCallResultConfirmed)d)._turnId; }
    if (d instanceof ChatAction_CToolCallContentChanged) { return ((ChatAction_CToolCallContentChanged)d)._turnId; }
    return ((ChatAction_CReasoning)d)._turnId;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_message() {
    ChatAction d = this;
    if (d instanceof ChatAction_CTurnStarted) { return ((ChatAction_CTurnStarted)d)._message; }
    return ((ChatAction_CPendingMessageSet)d)._message;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_queuedMessageId() {
    ChatAction d = this;
    return ((ChatAction_CTurnStarted)d)._queuedMessageId;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_usage() {
    ChatAction d = this;
    return ((ChatAction_CUsage)d)._usage;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_err() {
    ChatAction d = this;
    return ((ChatAction_CError)d)._err;
  }
  public Part dtor_part() {
    ChatAction d = this;
    return ((ChatAction_CResponsePart)d)._part;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_partId() {
    ChatAction d = this;
    if (d instanceof ChatAction_CDelta) { return ((ChatAction_CDelta)d)._partId; }
    return ((ChatAction_CReasoning)d)._partId;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_content() {
    ChatAction d = this;
    if (d instanceof ChatAction_CDelta) { return ((ChatAction_CDelta)d)._content; }
    if (d instanceof ChatAction_CToolCallDelta) { return ((ChatAction_CToolCallDelta)d)._content; }
    return ((ChatAction_CReasoning)d)._content;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_toolCallId() {
    ChatAction d = this;
    if (d instanceof ChatAction_CToolCallStart) { return ((ChatAction_CToolCallStart)d)._toolCallId; }
    if (d instanceof ChatAction_CToolCallDelta) { return ((ChatAction_CToolCallDelta)d)._toolCallId; }
    if (d instanceof ChatAction_CToolCallReady) { return ((ChatAction_CToolCallReady)d)._toolCallId; }
    if (d instanceof ChatAction_CToolCallConfirmed) { return ((ChatAction_CToolCallConfirmed)d)._toolCallId; }
    if (d instanceof ChatAction_CToolCallAuthRequired) { return ((ChatAction_CToolCallAuthRequired)d)._toolCallId; }
    if (d instanceof ChatAction_CToolCallAuthResolved) { return ((ChatAction_CToolCallAuthResolved)d)._toolCallId; }
    if (d instanceof ChatAction_CToolCallComplete) { return ((ChatAction_CToolCallComplete)d)._toolCallId; }
    if (d instanceof ChatAction_CToolCallResultConfirmed) { return ((ChatAction_CToolCallResultConfirmed)d)._toolCallId; }
    return ((ChatAction_CToolCallContentChanged)d)._toolCallId;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_toolName() {
    ChatAction d = this;
    return ((ChatAction_CToolCallStart)d)._toolName;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_displayName() {
    ChatAction d = this;
    return ((ChatAction_CToolCallStart)d)._displayName;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_intention() {
    ChatAction d = this;
    return ((ChatAction_CToolCallStart)d)._intention;
  }
  public agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> dtor_contributor() {
    ChatAction d = this;
    return ((ChatAction_CToolCallStart)d)._contributor;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_meta() {
    ChatAction d = this;
    if (d instanceof ChatAction_CToolCallStart) { return ((ChatAction_CToolCallStart)d)._meta; }
    if (d instanceof ChatAction_CToolCallDelta) { return ((ChatAction_CToolCallDelta)d)._meta; }
    if (d instanceof ChatAction_CToolCallReady) { return ((ChatAction_CToolCallReady)d)._meta; }
    if (d instanceof ChatAction_CToolCallConfirmed) { return ((ChatAction_CToolCallConfirmed)d)._meta; }
    if (d instanceof ChatAction_CToolCallAuthRequired) { return ((ChatAction_CToolCallAuthRequired)d)._meta; }
    if (d instanceof ChatAction_CToolCallAuthResolved) { return ((ChatAction_CToolCallAuthResolved)d)._meta; }
    if (d instanceof ChatAction_CToolCallComplete) { return ((ChatAction_CToolCallComplete)d)._meta; }
    if (d instanceof ChatAction_CToolCallResultConfirmed) { return ((ChatAction_CToolCallResultConfirmed)d)._meta; }
    return ((ChatAction_CToolCallContentChanged)d)._meta;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_invocationMessage() {
    ChatAction d = this;
    if (d instanceof ChatAction_CToolCallDelta) { return ((ChatAction_CToolCallDelta)d)._invocationMessage; }
    return ((ChatAction_CToolCallReady)d)._invocationMessage;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_toolInput() {
    ChatAction d = this;
    return ((ChatAction_CToolCallReady)d)._toolInput;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_confirmationTitle() {
    ChatAction d = this;
    return ((ChatAction_CToolCallReady)d)._confirmationTitle;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_riskAssessment() {
    ChatAction d = this;
    return ((ChatAction_CToolCallReady)d)._riskAssessment;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_edits() {
    ChatAction d = this;
    return ((ChatAction_CToolCallReady)d)._edits;
  }
  public agency.open.ahp.AhpSkeleton.Option<Boolean> dtor_editable() {
    ChatAction d = this;
    return ((ChatAction_CToolCallReady)d)._editable;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> dtor_options() {
    ChatAction d = this;
    return ((ChatAction_CToolCallReady)d)._options;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_confirmed() {
    ChatAction d = this;
    return ((ChatAction_CToolCallReady)d)._confirmed;
  }
  public boolean dtor_approved() {
    ChatAction d = this;
    if (d instanceof ChatAction_CToolCallConfirmed) { return ((ChatAction_CToolCallConfirmed)d)._approved; }
    return ((ChatAction_CToolCallResultConfirmed)d)._approved;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_confirmedVal() {
    ChatAction d = this;
    return ((ChatAction_CToolCallConfirmed)d)._confirmedVal;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_reason() {
    ChatAction d = this;
    return ((ChatAction_CToolCallConfirmed)d)._reason;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_reasonMessage() {
    ChatAction d = this;
    return ((ChatAction_CToolCallConfirmed)d)._reasonMessage;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_userSuggestion() {
    ChatAction d = this;
    return ((ChatAction_CToolCallConfirmed)d)._userSuggestion;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_editedToolInput() {
    ChatAction d = this;
    return ((ChatAction_CToolCallConfirmed)d)._editedToolInput;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_selectedOptionId() {
    ChatAction d = this;
    return ((ChatAction_CToolCallConfirmed)d)._selectedOptionId;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_auth() {
    ChatAction d = this;
    return ((ChatAction_CToolCallAuthRequired)d)._auth;
  }
  public boolean dtor_success() {
    ChatAction d = this;
    return ((ChatAction_CToolCallComplete)d)._success;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_pastTenseMessage() {
    ChatAction d = this;
    return ((ChatAction_CToolCallComplete)d)._pastTenseMessage;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_resultContent() {
    ChatAction d = this;
    return ((ChatAction_CToolCallComplete)d)._resultContent;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_structuredContent() {
    ChatAction d = this;
    return ((ChatAction_CToolCallComplete)d)._structuredContent;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_error() {
    ChatAction d = this;
    return ((ChatAction_CToolCallComplete)d)._error;
  }
  public boolean dtor_requiresResultConfirmation() {
    ChatAction d = this;
    return ((ChatAction_CToolCallComplete)d)._requiresResultConfirmation;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_newContent() {
    ChatAction d = this;
    return ((ChatAction_CToolCallContentChanged)d)._newContent;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_upToTurnId() {
    ChatAction d = this;
    return ((ChatAction_CTruncated)d)._upToTurnId;
  }
  public dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> dtor_order() {
    ChatAction d = this;
    return ((ChatAction_CQueuedReordered)d)._order;
  }
  public dafny.DafnySequence<? extends Turn> dtor_loaded() {
    ChatAction d = this;
    return ((ChatAction_CTurnsLoaded)d)._loaded;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_cursor() {
    ChatAction d = this;
    return ((ChatAction_CTurnsLoaded)d)._cursor;
  }
  public InputReq dtor_req() {
    ChatAction d = this;
    return ((ChatAction_CInputRequested)d)._req;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_requestId() {
    ChatAction d = this;
    if (d instanceof ChatAction_CInputAnswerChanged) { return ((ChatAction_CInputAnswerChanged)d)._requestId; }
    return ((ChatAction_CInputCompleted)d)._requestId;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_questionId() {
    ChatAction d = this;
    return ((ChatAction_CInputAnswerChanged)d)._questionId;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_answer() {
    ChatAction d = this;
    return ((ChatAction_CInputAnswerChanged)d)._answer;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_response() {
    ChatAction d = this;
    return ((ChatAction_CInputCompleted)d)._response;
  }
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> dtor_answers() {
    ChatAction d = this;
    return ((ChatAction_CInputCompleted)d)._answers;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_kind() {
    ChatAction d = this;
    if (d instanceof ChatAction_CPendingMessageSet) { return ((ChatAction_CPendingMessageSet)d)._kind; }
    return ((ChatAction_CPendingMessageRemoved)d)._kind;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_id() {
    ChatAction d = this;
    if (d instanceof ChatAction_CPendingMessageSet) { return ((ChatAction_CPendingMessageSet)d)._id; }
    return ((ChatAction_CPendingMessageRemoved)d)._id;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    ChatAction d = this;
    return ((ChatAction_CUnknown)d)._raw;
  }
}
