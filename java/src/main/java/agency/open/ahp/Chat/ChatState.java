// Class ChatState
// Dafny class ChatState compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatState {
  public dafny.DafnySequence<? extends Turn> _turns;
  public dafny.DafnySequence<? extends dafny.CodePoint> _title;
  public int _status;
  public dafny.DafnySequence<? extends dafny.CodePoint> _modifiedAt;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _draft;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _activity;
  public agency.open.ahp.AhpSkeleton.Option<Turn> _activeTurn;
  public agency.open.ahp.AhpSkeleton.Option<QMsg> _steeringMessage;
  public dafny.DafnySequence<? extends QMsg> _queuedMessages;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _nextCursor;
  public ChatState (dafny.DafnySequence<? extends Turn> turns, dafny.DafnySequence<? extends dafny.CodePoint> title, int status, dafny.DafnySequence<? extends dafny.CodePoint> modifiedAt, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> draft, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<Turn> activeTurn, agency.open.ahp.AhpSkeleton.Option<QMsg> steeringMessage, dafny.DafnySequence<? extends QMsg> queuedMessages, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> nextCursor) {
    this._turns = turns;
    this._title = title;
    this._status = status;
    this._modifiedAt = modifiedAt;
    this._draft = draft;
    this._activity = activity;
    this._activeTurn = activeTurn;
    this._steeringMessage = steeringMessage;
    this._queuedMessages = queuedMessages;
    this._nextCursor = nextCursor;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatState o = (ChatState)other;
    return true && java.util.Objects.equals(this._turns, o._turns) && java.util.Objects.equals(this._title, o._title) && this._status == o._status && java.util.Objects.equals(this._modifiedAt, o._modifiedAt) && java.util.Objects.equals(this._draft, o._draft) && java.util.Objects.equals(this._activity, o._activity) && java.util.Objects.equals(this._activeTurn, o._activeTurn) && java.util.Objects.equals(this._steeringMessage, o._steeringMessage) && java.util.Objects.equals(this._queuedMessages, o._queuedMessages) && java.util.Objects.equals(this._nextCursor, o._nextCursor);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turns);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._title);
    hash = ((hash << 5) + hash) + java.lang.Integer.hashCode(this._status);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._modifiedAt);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._draft);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activity);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activeTurn);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._steeringMessage);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._queuedMessages);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._nextCursor);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatState.ChatState");
    s.append("(");
    s.append(dafny.Helpers.toString(this._turns));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._title));
    s.append(", ");
    s.append(this._status);
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._modifiedAt));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._draft));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._activity));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._activeTurn));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._steeringMessage));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._queuedMessages));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._nextCursor));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<ChatState> _TYPE = dafny.TypeDescriptor.<ChatState>referenceWithInitializer(ChatState.class, () -> ChatState.Default());
  public static dafny.TypeDescriptor<ChatState> _typeDescriptor() {
    return (dafny.TypeDescriptor<ChatState>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ChatState theDefault = ChatState.create(dafny.DafnySequence.<Turn> empty(Turn._typeDescriptor()), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), 0, dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<Turn>Default(Turn._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<QMsg>Default(QMsg._typeDescriptor()), dafny.DafnySequence.<QMsg> empty(QMsg._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)));
  public static ChatState Default() {
    return theDefault;
  }
  public static ChatState create(dafny.DafnySequence<? extends Turn> turns, dafny.DafnySequence<? extends dafny.CodePoint> title, int status, dafny.DafnySequence<? extends dafny.CodePoint> modifiedAt, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> draft, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<Turn> activeTurn, agency.open.ahp.AhpSkeleton.Option<QMsg> steeringMessage, dafny.DafnySequence<? extends QMsg> queuedMessages, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> nextCursor) {
    return new ChatState(turns, title, status, modifiedAt, draft, activity, activeTurn, steeringMessage, queuedMessages, nextCursor);
  }
  public static ChatState create_ChatState(dafny.DafnySequence<? extends Turn> turns, dafny.DafnySequence<? extends dafny.CodePoint> title, int status, dafny.DafnySequence<? extends dafny.CodePoint> modifiedAt, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> draft, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<Turn> activeTurn, agency.open.ahp.AhpSkeleton.Option<QMsg> steeringMessage, dafny.DafnySequence<? extends QMsg> queuedMessages, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> nextCursor) {
    return create(turns, title, status, modifiedAt, draft, activity, activeTurn, steeringMessage, queuedMessages, nextCursor);
  }
  public boolean is_ChatState() { return true; }
  public dafny.DafnySequence<? extends Turn> dtor_turns() {
    return this._turns;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_title() {
    return this._title;
  }
  public int dtor_status() {
    return this._status;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_modifiedAt() {
    return this._modifiedAt;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_draft() {
    return this._draft;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_activity() {
    return this._activity;
  }
  public agency.open.ahp.AhpSkeleton.Option<Turn> dtor_activeTurn() {
    return this._activeTurn;
  }
  public agency.open.ahp.AhpSkeleton.Option<QMsg> dtor_steeringMessage() {
    return this._steeringMessage;
  }
  public dafny.DafnySequence<? extends QMsg> dtor_queuedMessages() {
    return this._queuedMessages;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_nextCursor() {
    return this._nextCursor;
  }
}
