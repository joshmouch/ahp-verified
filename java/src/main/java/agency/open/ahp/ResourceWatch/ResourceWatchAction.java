// Class ResourceWatchAction
// Dafny class ResourceWatchAction compiled into Java
package agency.open.ahp.ResourceWatch;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class ResourceWatchAction {
  public ResourceWatchAction() {
  }
  private static final dafny.TypeDescriptor<ResourceWatchAction> _TYPE = dafny.TypeDescriptor.<ResourceWatchAction>referenceWithInitializer(ResourceWatchAction.class, () -> ResourceWatchAction.Default());
  public static dafny.TypeDescriptor<ResourceWatchAction> _typeDescriptor() {
    return (dafny.TypeDescriptor<ResourceWatchAction>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ResourceWatchAction theDefault = ResourceWatchAction.create_RWChanged(agency.open.ahp.ConfluxCodec.Json.Default());
  public static ResourceWatchAction Default() {
    return theDefault;
  }
  public static ResourceWatchAction create_RWChanged(agency.open.ahp.ConfluxCodec.Json changes) {
    return new ResourceWatchAction_RWChanged(changes);
  }
  public static ResourceWatchAction create_RWUnknown(agency.open.ahp.ConfluxCodec.Json raw) {
    return new ResourceWatchAction_RWUnknown(raw);
  }
  public boolean is_RWChanged() { return this instanceof ResourceWatchAction_RWChanged; }
  public boolean is_RWUnknown() { return this instanceof ResourceWatchAction_RWUnknown; }
  public agency.open.ahp.ConfluxCodec.Json dtor_changes() {
    ResourceWatchAction d = this;
    return ((ResourceWatchAction_RWChanged)d)._changes;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    ResourceWatchAction d = this;
    return ((ResourceWatchAction_RWUnknown)d)._raw;
  }
}
