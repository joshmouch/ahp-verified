// Class SessionState
// Dafny class SessionState compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionState {
  public dafny.DafnySequence<? extends dafny.CodePoint> _provider;
  public dafny.DafnySequence<? extends dafny.CodePoint> _title;
  public int _status;
  public dafny.DafnySequence<? extends dafny.CodePoint> _lifecycle;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _activity;
  public agency.open.ahp.AhpSkeleton.Option<SConfig> _config;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _meta;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _creationError;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _serverTools;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _changesets;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _canvases;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _openCanvases;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _defaultChat;
  public dafny.DafnySequence<? extends Client> _activeClients;
  public dafny.DafnySequence<? extends Chat> _chats;
  public dafny.DafnySequence<? extends Cust> _customizations;
  public dafny.DafnySequence<? extends InputReq> _inputNeeded;
  public SessionState (dafny.DafnySequence<? extends dafny.CodePoint> provider, dafny.DafnySequence<? extends dafny.CodePoint> title, int status, dafny.DafnySequence<? extends dafny.CodePoint> lifecycle, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<SConfig> config, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> creationError, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> serverTools, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> changesets, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> canvases, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> openCanvases, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> defaultChat, dafny.DafnySequence<? extends Client> activeClients, dafny.DafnySequence<? extends Chat> chats, dafny.DafnySequence<? extends Cust> customizations, dafny.DafnySequence<? extends InputReq> inputNeeded) {
    this._provider = provider;
    this._title = title;
    this._status = status;
    this._lifecycle = lifecycle;
    this._activity = activity;
    this._config = config;
    this._meta = meta;
    this._creationError = creationError;
    this._serverTools = serverTools;
    this._changesets = changesets;
    this._canvases = canvases;
    this._openCanvases = openCanvases;
    this._defaultChat = defaultChat;
    this._activeClients = activeClients;
    this._chats = chats;
    this._customizations = customizations;
    this._inputNeeded = inputNeeded;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionState o = (SessionState)other;
    return true && java.util.Objects.equals(this._provider, o._provider) && java.util.Objects.equals(this._title, o._title) && this._status == o._status && java.util.Objects.equals(this._lifecycle, o._lifecycle) && java.util.Objects.equals(this._activity, o._activity) && java.util.Objects.equals(this._config, o._config) && java.util.Objects.equals(this._meta, o._meta) && java.util.Objects.equals(this._creationError, o._creationError) && java.util.Objects.equals(this._serverTools, o._serverTools) && java.util.Objects.equals(this._changesets, o._changesets) && java.util.Objects.equals(this._canvases, o._canvases) && java.util.Objects.equals(this._openCanvases, o._openCanvases) && java.util.Objects.equals(this._defaultChat, o._defaultChat) && java.util.Objects.equals(this._activeClients, o._activeClients) && java.util.Objects.equals(this._chats, o._chats) && java.util.Objects.equals(this._customizations, o._customizations) && java.util.Objects.equals(this._inputNeeded, o._inputNeeded);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._provider);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._title);
    hash = ((hash << 5) + hash) + java.lang.Integer.hashCode(this._status);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._lifecycle);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activity);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._config);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._meta);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._creationError);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._serverTools);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._changesets);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._canvases);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._openCanvases);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._defaultChat);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activeClients);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._chats);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._customizations);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._inputNeeded);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionState.SessionState");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._provider));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._title));
    s.append(", ");
    s.append(this._status);
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._lifecycle));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._activity));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._config));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._meta));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._creationError));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._serverTools));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._changesets));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._canvases));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._openCanvases));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._defaultChat));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._activeClients));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._chats));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._customizations));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._inputNeeded));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<SessionState> _TYPE = dafny.TypeDescriptor.<SessionState>referenceWithInitializer(SessionState.class, () -> SessionState.Default());
  public static dafny.TypeDescriptor<SessionState> _typeDescriptor() {
    return (dafny.TypeDescriptor<SessionState>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final SessionState theDefault = SessionState.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), 0, dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<SConfig>Default(SConfig._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), dafny.DafnySequence.<Client> empty(Client._typeDescriptor()), dafny.DafnySequence.<Chat> empty(Chat._typeDescriptor()), dafny.DafnySequence.<Cust> empty(Cust._typeDescriptor()), dafny.DafnySequence.<InputReq> empty(InputReq._typeDescriptor()));
  public static SessionState Default() {
    return theDefault;
  }
  public static SessionState create(dafny.DafnySequence<? extends dafny.CodePoint> provider, dafny.DafnySequence<? extends dafny.CodePoint> title, int status, dafny.DafnySequence<? extends dafny.CodePoint> lifecycle, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<SConfig> config, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> creationError, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> serverTools, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> changesets, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> canvases, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> openCanvases, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> defaultChat, dafny.DafnySequence<? extends Client> activeClients, dafny.DafnySequence<? extends Chat> chats, dafny.DafnySequence<? extends Cust> customizations, dafny.DafnySequence<? extends InputReq> inputNeeded) {
    return new SessionState(provider, title, status, lifecycle, activity, config, meta, creationError, serverTools, changesets, canvases, openCanvases, defaultChat, activeClients, chats, customizations, inputNeeded);
  }
  public static SessionState create_SessionState(dafny.DafnySequence<? extends dafny.CodePoint> provider, dafny.DafnySequence<? extends dafny.CodePoint> title, int status, dafny.DafnySequence<? extends dafny.CodePoint> lifecycle, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<SConfig> config, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> creationError, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> serverTools, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> changesets, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> canvases, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> openCanvases, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> defaultChat, dafny.DafnySequence<? extends Client> activeClients, dafny.DafnySequence<? extends Chat> chats, dafny.DafnySequence<? extends Cust> customizations, dafny.DafnySequence<? extends InputReq> inputNeeded) {
    return create(provider, title, status, lifecycle, activity, config, meta, creationError, serverTools, changesets, canvases, openCanvases, defaultChat, activeClients, chats, customizations, inputNeeded);
  }
  public boolean is_SessionState() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_provider() {
    return this._provider;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_title() {
    return this._title;
  }
  public int dtor_status() {
    return this._status;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_lifecycle() {
    return this._lifecycle;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_activity() {
    return this._activity;
  }
  public agency.open.ahp.AhpSkeleton.Option<SConfig> dtor_config() {
    return this._config;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_meta() {
    return this._meta;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_creationError() {
    return this._creationError;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_serverTools() {
    return this._serverTools;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_changesets() {
    return this._changesets;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_canvases() {
    return this._canvases;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_openCanvases() {
    return this._openCanvases;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_defaultChat() {
    return this._defaultChat;
  }
  public dafny.DafnySequence<? extends Client> dtor_activeClients() {
    return this._activeClients;
  }
  public dafny.DafnySequence<? extends Chat> dtor_chats() {
    return this._chats;
  }
  public dafny.DafnySequence<? extends Cust> dtor_customizations() {
    return this._customizations;
  }
  public dafny.DafnySequence<? extends InputReq> dtor_inputNeeded() {
    return this._inputNeeded;
  }
}
