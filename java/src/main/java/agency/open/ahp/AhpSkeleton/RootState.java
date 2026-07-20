// Class RootState
// Dafny class RootState compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class RootState {
  public dafny.DafnySequence<? extends AgentInfo> _agents;
  public Option<java.math.BigInteger> _activeSessions;
  public Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _terminals;
  public Option<RootConfig> _config;
  public RootState (dafny.DafnySequence<? extends AgentInfo> agents, Option<java.math.BigInteger> activeSessions, Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> terminals, Option<RootConfig> config) {
    this._agents = agents;
    this._activeSessions = activeSessions;
    this._terminals = terminals;
    this._config = config;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    RootState o = (RootState)other;
    return true && java.util.Objects.equals(this._agents, o._agents) && java.util.Objects.equals(this._activeSessions, o._activeSessions) && java.util.Objects.equals(this._terminals, o._terminals) && java.util.Objects.equals(this._config, o._config);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._agents);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activeSessions);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._terminals);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._config);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.RootState.RootState");
    s.append("(");
    s.append(dafny.Helpers.toString(this._agents));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._activeSessions));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._terminals));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._config));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<RootState> _TYPE = dafny.TypeDescriptor.<RootState>referenceWithInitializer(RootState.class, () -> RootState.Default());
  public static dafny.TypeDescriptor<RootState> _typeDescriptor() {
    return (dafny.TypeDescriptor<RootState>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final RootState theDefault = RootState.create(dafny.DafnySequence.<AgentInfo> empty(AgentInfo._typeDescriptor()), Option.<java.math.BigInteger>Default(dafny.TypeDescriptor.BIG_INTEGER), Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>Default(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), Option.<RootConfig>Default(RootConfig._typeDescriptor()));
  public static RootState Default() {
    return theDefault;
  }
  public static RootState create(dafny.DafnySequence<? extends AgentInfo> agents, Option<java.math.BigInteger> activeSessions, Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> terminals, Option<RootConfig> config) {
    return new RootState(agents, activeSessions, terminals, config);
  }
  public static RootState create_RootState(dafny.DafnySequence<? extends AgentInfo> agents, Option<java.math.BigInteger> activeSessions, Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> terminals, Option<RootConfig> config) {
    return create(agents, activeSessions, terminals, config);
  }
  public boolean is_RootState() { return true; }
  public dafny.DafnySequence<? extends AgentInfo> dtor_agents() {
    return this._agents;
  }
  public Option<java.math.BigInteger> dtor_activeSessions() {
    return this._activeSessions;
  }
  public Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> dtor_terminals() {
    return this._terminals;
  }
  public Option<RootConfig> dtor_config() {
    return this._config;
  }
}
