// Class ResourceWatchState
// Dafny class ResourceWatchState compiled into Java
package agency.open.ahp.ResourceWatch;

@SuppressWarnings({"unchecked", "deprecation"})
public class ResourceWatchState {
  public dafny.DafnySequence<? extends dafny.CodePoint> _root;
  public boolean _recursive;
  public ResourceWatchState (dafny.DafnySequence<? extends dafny.CodePoint> root, boolean recursive) {
    this._root = root;
    this._recursive = recursive;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ResourceWatchState o = (ResourceWatchState)other;
    return true && java.util.Objects.equals(this._root, o._root) && this._recursive == o._recursive;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._root);
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._recursive);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("ResourceWatch.ResourceWatchState.ResourceWatchState");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._root));
    s.append(", ");
    s.append(this._recursive);
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<ResourceWatchState> _TYPE = dafny.TypeDescriptor.<ResourceWatchState>referenceWithInitializer(ResourceWatchState.class, () -> ResourceWatchState.Default());
  public static dafny.TypeDescriptor<ResourceWatchState> _typeDescriptor() {
    return (dafny.TypeDescriptor<ResourceWatchState>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ResourceWatchState theDefault = ResourceWatchState.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), false);
  public static ResourceWatchState Default() {
    return theDefault;
  }
  public static ResourceWatchState create(dafny.DafnySequence<? extends dafny.CodePoint> root, boolean recursive) {
    return new ResourceWatchState(root, recursive);
  }
  public static ResourceWatchState create_ResourceWatchState(dafny.DafnySequence<? extends dafny.CodePoint> root, boolean recursive) {
    return create(root, recursive);
  }
  public boolean is_ResourceWatchState() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_root() {
    return this._root;
  }
  public boolean dtor_recursive() {
    return this._recursive;
  }
}
