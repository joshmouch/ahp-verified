// Class RootAction
// Dafny class RootAction compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class RootAction {
  public RootAction() {
  }
  private static final dafny.TypeDescriptor<RootAction> _TYPE = dafny.TypeDescriptor.<RootAction>referenceWithInitializer(RootAction.class, () -> RootAction.Default());
  public static dafny.TypeDescriptor<RootAction> _typeDescriptor() {
    return (dafny.TypeDescriptor<RootAction>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final RootAction theDefault = RootAction.create_RootAgentsChanged(dafny.DafnySequence.<AgentInfo> empty(AgentInfo._typeDescriptor()));
  public static RootAction Default() {
    return theDefault;
  }
  public static RootAction create_RootAgentsChanged(dafny.DafnySequence<? extends AgentInfo> agents) {
    return new RootAction_RootAgentsChanged(agents);
  }
  public static RootAction create_RootActiveSessionsChanged(java.math.BigInteger activeSessions) {
    return new RootAction_RootActiveSessionsChanged(activeSessions);
  }
  public static RootAction create_RootTerminalsChanged(dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json> terminals) {
    return new RootAction_RootTerminalsChanged(terminals);
  }
  public static RootAction create_RootConfigChanged(dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> config, boolean replace) {
    return new RootAction_RootConfigChanged(config, replace);
  }
  public static RootAction create_RootUnknown(agency.open.ahp.ConfluxCodec.Json raw) {
    return new RootAction_RootUnknown(raw);
  }
  public boolean is_RootAgentsChanged() { return this instanceof RootAction_RootAgentsChanged; }
  public boolean is_RootActiveSessionsChanged() { return this instanceof RootAction_RootActiveSessionsChanged; }
  public boolean is_RootTerminalsChanged() { return this instanceof RootAction_RootTerminalsChanged; }
  public boolean is_RootConfigChanged() { return this instanceof RootAction_RootConfigChanged; }
  public boolean is_RootUnknown() { return this instanceof RootAction_RootUnknown; }
  public dafny.DafnySequence<? extends AgentInfo> dtor_agents() {
    RootAction d = this;
    return ((RootAction_RootAgentsChanged)d)._agents;
  }
  public java.math.BigInteger dtor_activeSessions() {
    RootAction d = this;
    return ((RootAction_RootActiveSessionsChanged)d)._activeSessions;
  }
  public dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json> dtor_terminals() {
    RootAction d = this;
    return ((RootAction_RootTerminalsChanged)d)._terminals;
  }
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> dtor_config() {
    RootAction d = this;
    return ((RootAction_RootConfigChanged)d)._config;
  }
  public boolean dtor_replace() {
    RootAction d = this;
    return ((RootAction_RootConfigChanged)d)._replace;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    RootAction d = this;
    return ((RootAction_RootUnknown)d)._raw;
  }
}
