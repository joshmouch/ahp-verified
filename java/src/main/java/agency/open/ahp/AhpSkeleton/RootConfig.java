// Class RootConfig
// Dafny class RootConfig compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class RootConfig {
  public agency.open.ahp.ConfluxCodec.Json _schema;
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _values;
  public RootConfig (agency.open.ahp.ConfluxCodec.Json schema, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> values) {
    this._schema = schema;
    this._values = values;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    RootConfig o = (RootConfig)other;
    return true && java.util.Objects.equals(this._schema, o._schema) && java.util.Objects.equals(this._values, o._values);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._schema);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._values);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.RootConfig.RootConfig");
    s.append("(");
    s.append(dafny.Helpers.toString(this._schema));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._values));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<RootConfig> _TYPE = dafny.TypeDescriptor.<RootConfig>referenceWithInitializer(RootConfig.class, () -> RootConfig.Default());
  public static dafny.TypeDescriptor<RootConfig> _typeDescriptor() {
    return (dafny.TypeDescriptor<RootConfig>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final RootConfig theDefault = RootConfig.create(agency.open.ahp.ConfluxCodec.Json.Default(), dafny.DafnyMap.<dafny.DafnySequence<? extends dafny.CodePoint>,agency.open.ahp.ConfluxCodec.Json> empty());
  public static RootConfig Default() {
    return theDefault;
  }
  public static RootConfig create(agency.open.ahp.ConfluxCodec.Json schema, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> values) {
    return new RootConfig(schema, values);
  }
  public static RootConfig create_RootConfig(agency.open.ahp.ConfluxCodec.Json schema, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> values) {
    return create(schema, values);
  }
  public boolean is_RootConfig() { return true; }
  public agency.open.ahp.ConfluxCodec.Json dtor_schema() {
    return this._schema;
  }
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> dtor_values() {
    return this._values;
  }
}
