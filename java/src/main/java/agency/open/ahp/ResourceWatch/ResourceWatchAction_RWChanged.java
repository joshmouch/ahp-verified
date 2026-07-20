// Class ResourceWatchAction_RWChanged
// Dafny class ResourceWatchAction_RWChanged compiled into Java
package agency.open.ahp.ResourceWatch;

@SuppressWarnings({"unchecked", "deprecation"})
public class ResourceWatchAction_RWChanged extends ResourceWatchAction {
  public agency.open.ahp.ConfluxCodec.Json _changes;
  public ResourceWatchAction_RWChanged (agency.open.ahp.ConfluxCodec.Json changes) {
    super();
    this._changes = changes;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ResourceWatchAction_RWChanged o = (ResourceWatchAction_RWChanged)other;
    return true && java.util.Objects.equals(this._changes, o._changes);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._changes);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("ResourceWatch.ResourceWatchAction.RWChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._changes));
    s.append(")");
    return s.toString();
  }
}
