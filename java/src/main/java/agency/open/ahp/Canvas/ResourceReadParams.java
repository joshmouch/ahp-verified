// Class ResourceReadParams
// Dafny class ResourceReadParams compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class ResourceReadParams {
  public dafny.DafnySequence<? extends dafny.CodePoint> _channel;
  public dafny.DafnySequence<? extends dafny.CodePoint> _uri;
  public ResourceReadParams (dafny.DafnySequence<? extends dafny.CodePoint> channel, dafny.DafnySequence<? extends dafny.CodePoint> uri) {
    this._channel = channel;
    this._uri = uri;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ResourceReadParams o = (ResourceReadParams)other;
    return true && java.util.Objects.equals(this._channel, o._channel) && java.util.Objects.equals(this._uri, o._uri);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._channel);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._uri);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.ResourceReadParams.ResourceReadParams");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._channel));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._uri));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<ResourceReadParams> _TYPE = dafny.TypeDescriptor.<ResourceReadParams>referenceWithInitializer(ResourceReadParams.class, () -> ResourceReadParams.Default());
  public static dafny.TypeDescriptor<ResourceReadParams> _typeDescriptor() {
    return (dafny.TypeDescriptor<ResourceReadParams>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ResourceReadParams theDefault = ResourceReadParams.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR));
  public static ResourceReadParams Default() {
    return theDefault;
  }
  public static ResourceReadParams create(dafny.DafnySequence<? extends dafny.CodePoint> channel, dafny.DafnySequence<? extends dafny.CodePoint> uri) {
    return new ResourceReadParams(channel, uri);
  }
  public static ResourceReadParams create_ResourceReadParams(dafny.DafnySequence<? extends dafny.CodePoint> channel, dafny.DafnySequence<? extends dafny.CodePoint> uri) {
    return create(channel, uri);
  }
  public boolean is_ResourceReadParams() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_channel() {
    return this._channel;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_uri() {
    return this._uri;
  }
}
