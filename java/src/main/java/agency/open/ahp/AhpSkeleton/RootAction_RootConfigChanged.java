// Class RootAction_RootConfigChanged
// Dafny class RootAction_RootConfigChanged compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class RootAction_RootConfigChanged extends RootAction {
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _config;
  public boolean _replace;
  public RootAction_RootConfigChanged (dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> config, boolean replace) {
    super();
    this._config = config;
    this._replace = replace;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    RootAction_RootConfigChanged o = (RootAction_RootConfigChanged)other;
    return true && java.util.Objects.equals(this._config, o._config) && this._replace == o._replace;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 3;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._config);
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._replace);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.RootAction.RootConfigChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._config));
    s.append(", ");
    s.append(this._replace);
    s.append(")");
    return s.toString();
  }
}
