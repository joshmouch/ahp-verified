// Class SConfig
// Dafny class SConfig compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SConfig {
  public agency.open.ahp.ConfluxCodec.Json _schema;
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _values;
  public SConfig (agency.open.ahp.ConfluxCodec.Json schema, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> values) {
    this._schema = schema;
    this._values = values;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SConfig o = (SConfig)other;
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
    s.append("Session.SConfig.SConfig");
    s.append("(");
    s.append(dafny.Helpers.toString(this._schema));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._values));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<SConfig> _TYPE = dafny.TypeDescriptor.<SConfig>referenceWithInitializer(SConfig.class, () -> SConfig.Default());
  public static dafny.TypeDescriptor<SConfig> _typeDescriptor() {
    return (dafny.TypeDescriptor<SConfig>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final SConfig theDefault = SConfig.create(agency.open.ahp.ConfluxCodec.Json.Default(), dafny.DafnyMap.<dafny.DafnySequence<? extends dafny.CodePoint>,agency.open.ahp.ConfluxCodec.Json> empty());
  public static SConfig Default() {
    return theDefault;
  }
  public static SConfig create(agency.open.ahp.ConfluxCodec.Json schema, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> values) {
    return new SConfig(schema, values);
  }
  public static SConfig create_SConfig(agency.open.ahp.ConfluxCodec.Json schema, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> values) {
    return create(schema, values);
  }
  public boolean is_SConfig() { return true; }
  public agency.open.ahp.ConfluxCodec.Json dtor_schema() {
    return this._schema;
  }
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> dtor_values() {
    return this._values;
  }
}
