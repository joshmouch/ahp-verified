// Class AgentInfo
// Dafny class AgentInfo compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class AgentInfo {
  public dafny.DafnySequence<? extends dafny.CodePoint> _provider;
  public dafny.DafnySequence<? extends dafny.CodePoint> _displayName;
  public dafny.DafnySequence<? extends dafny.CodePoint> _description;
  public dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> _models;
  public AgentInfo (dafny.DafnySequence<? extends dafny.CodePoint> provider, dafny.DafnySequence<? extends dafny.CodePoint> displayName, dafny.DafnySequence<? extends dafny.CodePoint> description, dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> models) {
    this._provider = provider;
    this._displayName = displayName;
    this._description = description;
    this._models = models;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    AgentInfo o = (AgentInfo)other;
    return true && java.util.Objects.equals(this._provider, o._provider) && java.util.Objects.equals(this._displayName, o._displayName) && java.util.Objects.equals(this._description, o._description) && java.util.Objects.equals(this._models, o._models);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._provider);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._displayName);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._description);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._models);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.AgentInfo.AgentInfo");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._provider));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._displayName));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._description));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._models));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<AgentInfo> _TYPE = dafny.TypeDescriptor.<AgentInfo>referenceWithInitializer(AgentInfo.class, () -> AgentInfo.Default());
  public static dafny.TypeDescriptor<AgentInfo> _typeDescriptor() {
    return (dafny.TypeDescriptor<AgentInfo>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final AgentInfo theDefault = AgentInfo.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.DafnySequence<? extends dafny.CodePoint>> empty(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)));
  public static AgentInfo Default() {
    return theDefault;
  }
  public static AgentInfo create(dafny.DafnySequence<? extends dafny.CodePoint> provider, dafny.DafnySequence<? extends dafny.CodePoint> displayName, dafny.DafnySequence<? extends dafny.CodePoint> description, dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> models) {
    return new AgentInfo(provider, displayName, description, models);
  }
  public static AgentInfo create_AgentInfo(dafny.DafnySequence<? extends dafny.CodePoint> provider, dafny.DafnySequence<? extends dafny.CodePoint> displayName, dafny.DafnySequence<? extends dafny.CodePoint> description, dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> models) {
    return create(provider, displayName, description, models);
  }
  public boolean is_AgentInfo() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_provider() {
    return this._provider;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_displayName() {
    return this._displayName;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_description() {
    return this._description;
  }
  public dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> dtor_models() {
    return this._models;
  }
}
