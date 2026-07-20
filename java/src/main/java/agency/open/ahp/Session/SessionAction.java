// Class SessionAction
// Dafny class SessionAction compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class SessionAction {
  public SessionAction() {
  }
  private static final dafny.TypeDescriptor<SessionAction> _TYPE = dafny.TypeDescriptor.<SessionAction>referenceWithInitializer(SessionAction.class, () -> SessionAction.Default());
  public static dafny.TypeDescriptor<SessionAction> _typeDescriptor() {
    return (dafny.TypeDescriptor<SessionAction>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final SessionAction theDefault = SessionAction.create_Ready();
  public static SessionAction Default() {
    return theDefault;
  }
  public static SessionAction create_Ready() {
    return new SessionAction_Ready();
  }
  public static SessionAction create_CreationFailed(agency.open.ahp.ConfluxCodec.Json error) {
    return new SessionAction_CreationFailed(error);
  }
  public static SessionAction create_TitleChanged(dafny.DafnySequence<? extends dafny.CodePoint> title) {
    return new SessionAction_TitleChanged(title);
  }
  public static SessionAction create_ServerToolsChanged(agency.open.ahp.ConfluxCodec.Json tools) {
    return new SessionAction_ServerToolsChanged(tools);
  }
  public static SessionAction create_MetaChanged(agency.open.ahp.ConfluxCodec.Json meta) {
    return new SessionAction_MetaChanged(meta);
  }
  public static SessionAction create_IsReadChanged(boolean isRead) {
    return new SessionAction_IsReadChanged(isRead);
  }
  public static SessionAction create_IsArchivedChanged(boolean isArchived) {
    return new SessionAction_IsArchivedChanged(isArchived);
  }
  public static SessionAction create_ActivityChanged(agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity) {
    return new SessionAction_ActivityChanged(activity);
  }
  public static SessionAction create_ConfigChanged(dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> config, boolean replace) {
    return new SessionAction_ConfigChanged(config, replace);
  }
  public static SessionAction create_CustomizationsChanged(dafny.DafnySequence<? extends Cust> customizations) {
    return new SessionAction_CustomizationsChanged(customizations);
  }
  public static SessionAction create_CustomizationToggled(dafny.DafnySequence<? extends dafny.CodePoint> id, boolean enabled) {
    return new SessionAction_CustomizationToggled(id, enabled);
  }
  public static SessionAction create_CustomizationRemoved(dafny.DafnySequence<? extends dafny.CodePoint> id) {
    return new SessionAction_CustomizationRemoved(id);
  }
  public static SessionAction create_DefaultChatChanged(agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> chat) {
    return new SessionAction_DefaultChatChanged(chat);
  }
  public static SessionAction create_ChatAdded(Chat summary) {
    return new SessionAction_ChatAdded(summary);
  }
  public static SessionAction create_ChatRemoved(dafny.DafnySequence<? extends dafny.CodePoint> resource) {
    return new SessionAction_ChatRemoved(resource);
  }
  public static SessionAction create_ChatUpdated(dafny.DafnySequence<? extends dafny.CodePoint> resource, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> status, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> modifiedAt) {
    return new SessionAction_ChatUpdated(resource, status, activity, modifiedAt);
  }
  public static SessionAction create_ChangesetsChanged(agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> changesets) {
    return new SessionAction_ChangesetsChanged(changesets);
  }
  public static SessionAction create_ActiveClientSet(Client client) {
    return new SessionAction_ActiveClientSet(client);
  }
  public static SessionAction create_ActiveClientRemoved(dafny.DafnySequence<? extends dafny.CodePoint> clientId) {
    return new SessionAction_ActiveClientRemoved(clientId);
  }
  public static SessionAction create_CanvasesChanged(agency.open.ahp.ConfluxCodec.Json canvases) {
    return new SessionAction_CanvasesChanged(canvases);
  }
  public static SessionAction create_OpenCanvasesChanged(agency.open.ahp.ConfluxCodec.Json openCanvases) {
    return new SessionAction_OpenCanvasesChanged(openCanvases);
  }
  public static SessionAction create_InputNeededSet(InputReq req) {
    return new SessionAction_InputNeededSet(req);
  }
  public static SessionAction create_InputNeededRemoved(dafny.DafnySequence<? extends dafny.CodePoint> id) {
    return new SessionAction_InputNeededRemoved(id);
  }
  public static SessionAction create_CustomizationUpdated(Cust customization) {
    return new SessionAction_CustomizationUpdated(customization);
  }
  public static SessionAction create_McpServerStateChanged(dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.ConfluxCodec.Json state, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> channel) {
    return new SessionAction_McpServerStateChanged(id, state, channel);
  }
  public static SessionAction create_McpServerStartRequested(dafny.DafnySequence<? extends dafny.CodePoint> id) {
    return new SessionAction_McpServerStartRequested(id);
  }
  public static SessionAction create_McpServerStopRequested(dafny.DafnySequence<? extends dafny.CodePoint> id) {
    return new SessionAction_McpServerStopRequested(id);
  }
  public static SessionAction create_SUnknown(agency.open.ahp.ConfluxCodec.Json raw) {
    return new SessionAction_SUnknown(raw);
  }
  public boolean is_Ready() { return this instanceof SessionAction_Ready; }
  public boolean is_CreationFailed() { return this instanceof SessionAction_CreationFailed; }
  public boolean is_TitleChanged() { return this instanceof SessionAction_TitleChanged; }
  public boolean is_ServerToolsChanged() { return this instanceof SessionAction_ServerToolsChanged; }
  public boolean is_MetaChanged() { return this instanceof SessionAction_MetaChanged; }
  public boolean is_IsReadChanged() { return this instanceof SessionAction_IsReadChanged; }
  public boolean is_IsArchivedChanged() { return this instanceof SessionAction_IsArchivedChanged; }
  public boolean is_ActivityChanged() { return this instanceof SessionAction_ActivityChanged; }
  public boolean is_ConfigChanged() { return this instanceof SessionAction_ConfigChanged; }
  public boolean is_CustomizationsChanged() { return this instanceof SessionAction_CustomizationsChanged; }
  public boolean is_CustomizationToggled() { return this instanceof SessionAction_CustomizationToggled; }
  public boolean is_CustomizationRemoved() { return this instanceof SessionAction_CustomizationRemoved; }
  public boolean is_DefaultChatChanged() { return this instanceof SessionAction_DefaultChatChanged; }
  public boolean is_ChatAdded() { return this instanceof SessionAction_ChatAdded; }
  public boolean is_ChatRemoved() { return this instanceof SessionAction_ChatRemoved; }
  public boolean is_ChatUpdated() { return this instanceof SessionAction_ChatUpdated; }
  public boolean is_ChangesetsChanged() { return this instanceof SessionAction_ChangesetsChanged; }
  public boolean is_ActiveClientSet() { return this instanceof SessionAction_ActiveClientSet; }
  public boolean is_ActiveClientRemoved() { return this instanceof SessionAction_ActiveClientRemoved; }
  public boolean is_CanvasesChanged() { return this instanceof SessionAction_CanvasesChanged; }
  public boolean is_OpenCanvasesChanged() { return this instanceof SessionAction_OpenCanvasesChanged; }
  public boolean is_InputNeededSet() { return this instanceof SessionAction_InputNeededSet; }
  public boolean is_InputNeededRemoved() { return this instanceof SessionAction_InputNeededRemoved; }
  public boolean is_CustomizationUpdated() { return this instanceof SessionAction_CustomizationUpdated; }
  public boolean is_McpServerStateChanged() { return this instanceof SessionAction_McpServerStateChanged; }
  public boolean is_McpServerStartRequested() { return this instanceof SessionAction_McpServerStartRequested; }
  public boolean is_McpServerStopRequested() { return this instanceof SessionAction_McpServerStopRequested; }
  public boolean is_SUnknown() { return this instanceof SessionAction_SUnknown; }
  public agency.open.ahp.ConfluxCodec.Json dtor_error() {
    SessionAction d = this;
    return ((SessionAction_CreationFailed)d)._error;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_title() {
    SessionAction d = this;
    return ((SessionAction_TitleChanged)d)._title;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_tools() {
    SessionAction d = this;
    return ((SessionAction_ServerToolsChanged)d)._tools;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_meta() {
    SessionAction d = this;
    return ((SessionAction_MetaChanged)d)._meta;
  }
  public boolean dtor_isRead() {
    SessionAction d = this;
    return ((SessionAction_IsReadChanged)d)._isRead;
  }
  public boolean dtor_isArchived() {
    SessionAction d = this;
    return ((SessionAction_IsArchivedChanged)d)._isArchived;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_activity() {
    SessionAction d = this;
    if (d instanceof SessionAction_ActivityChanged) { return ((SessionAction_ActivityChanged)d)._activity; }
    return ((SessionAction_ChatUpdated)d)._activity;
  }
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> dtor_config() {
    SessionAction d = this;
    return ((SessionAction_ConfigChanged)d)._config;
  }
  public boolean dtor_replace() {
    SessionAction d = this;
    return ((SessionAction_ConfigChanged)d)._replace;
  }
  public dafny.DafnySequence<? extends Cust> dtor_customizations() {
    SessionAction d = this;
    return ((SessionAction_CustomizationsChanged)d)._customizations;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_id() {
    SessionAction d = this;
    if (d instanceof SessionAction_CustomizationToggled) { return ((SessionAction_CustomizationToggled)d)._id; }
    if (d instanceof SessionAction_CustomizationRemoved) { return ((SessionAction_CustomizationRemoved)d)._id; }
    if (d instanceof SessionAction_InputNeededRemoved) { return ((SessionAction_InputNeededRemoved)d)._id; }
    if (d instanceof SessionAction_McpServerStateChanged) { return ((SessionAction_McpServerStateChanged)d)._id; }
    if (d instanceof SessionAction_McpServerStartRequested) { return ((SessionAction_McpServerStartRequested)d)._id; }
    return ((SessionAction_McpServerStopRequested)d)._id;
  }
  public boolean dtor_enabled() {
    SessionAction d = this;
    return ((SessionAction_CustomizationToggled)d)._enabled;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_chat() {
    SessionAction d = this;
    return ((SessionAction_DefaultChatChanged)d)._chat;
  }
  public Chat dtor_summary() {
    SessionAction d = this;
    return ((SessionAction_ChatAdded)d)._summary;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_resource() {
    SessionAction d = this;
    if (d instanceof SessionAction_ChatRemoved) { return ((SessionAction_ChatRemoved)d)._resource; }
    return ((SessionAction_ChatUpdated)d)._resource;
  }
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> dtor_status() {
    SessionAction d = this;
    return ((SessionAction_ChatUpdated)d)._status;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_modifiedAt() {
    SessionAction d = this;
    return ((SessionAction_ChatUpdated)d)._modifiedAt;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_changesets() {
    SessionAction d = this;
    return ((SessionAction_ChangesetsChanged)d)._changesets;
  }
  public Client dtor_client() {
    SessionAction d = this;
    return ((SessionAction_ActiveClientSet)d)._client;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_clientId() {
    SessionAction d = this;
    return ((SessionAction_ActiveClientRemoved)d)._clientId;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_canvases() {
    SessionAction d = this;
    return ((SessionAction_CanvasesChanged)d)._canvases;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_openCanvases() {
    SessionAction d = this;
    return ((SessionAction_OpenCanvasesChanged)d)._openCanvases;
  }
  public InputReq dtor_req() {
    SessionAction d = this;
    return ((SessionAction_InputNeededSet)d)._req;
  }
  public Cust dtor_customization() {
    SessionAction d = this;
    return ((SessionAction_CustomizationUpdated)d)._customization;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_state() {
    SessionAction d = this;
    return ((SessionAction_McpServerStateChanged)d)._state;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_channel() {
    SessionAction d = this;
    return ((SessionAction_McpServerStateChanged)d)._channel;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    SessionAction d = this;
    return ((SessionAction_SUnknown)d)._raw;
  }
}
