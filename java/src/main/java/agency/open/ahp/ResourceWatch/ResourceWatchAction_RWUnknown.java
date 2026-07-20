// Class ResourceWatchAction_RWUnknown
// Dafny class ResourceWatchAction_RWUnknown compiled into Java
package agency.open.ahp.ResourceWatch;

@SuppressWarnings({"unchecked", "deprecation"})
public class ResourceWatchAction_RWUnknown extends ResourceWatchAction {
  public agency.open.ahp.ConfluxCodec.Json _raw;
  public ResourceWatchAction_RWUnknown (agency.open.ahp.ConfluxCodec.Json raw) {
    super();
    this._raw = raw;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ResourceWatchAction_RWUnknown o = (ResourceWatchAction_RWUnknown)other;
    return true && java.util.Objects.equals(this._raw, o._raw);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._raw);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("ResourceWatch.ResourceWatchAction.RWUnknown");
    s.append("(");
    s.append(dafny.Helpers.toString(this._raw));
    s.append(")");
    return s.toString();
  }
}
