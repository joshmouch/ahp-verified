// Class Annotation
// Dafny class Annotation compiled into Java
package agency.open.ahp.Annotations;

@SuppressWarnings({"unchecked", "deprecation"})
public class Annotation {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _resource;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _range;
  public boolean _resolved;
  public dafny.DafnySequence<? extends Entry> _entries;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _meta;
  public Annotation (dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> resource, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> range, boolean resolved, dafny.DafnySequence<? extends Entry> entries, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    this._id = id;
    this._turnId = turnId;
    this._resource = resource;
    this._range = range;
    this._resolved = resolved;
    this._entries = entries;
    this._meta = meta;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Annotation o = (Annotation)other;
    return true && java.util.Objects.equals(this._id, o._id) && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._resource, o._resource) && java.util.Objects.equals(this._range, o._range) && this._resolved == o._resolved && java.util.Objects.equals(this._entries, o._entries) && java.util.Objects.equals(this._meta, o._meta);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._resource);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._range);
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._resolved);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._entries);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._meta);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Annotations.Annotation.Annotation");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._resource));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._range));
    s.append(", ");
    s.append(this._resolved);
    s.append(", ");
    s.append(dafny.Helpers.toString(this._entries));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._meta));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<Annotation> _TYPE = dafny.TypeDescriptor.<Annotation>referenceWithInitializer(Annotation.class, () -> Annotation.Default());
  public static dafny.TypeDescriptor<Annotation> _typeDescriptor() {
    return (dafny.TypeDescriptor<Annotation>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final Annotation theDefault = Annotation.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), false, dafny.DafnySequence.<Entry> empty(Entry._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
  public static Annotation Default() {
    return theDefault;
  }
  public static Annotation create(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> resource, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> range, boolean resolved, dafny.DafnySequence<? extends Entry> entries, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return new Annotation(id, turnId, resource, range, resolved, entries, meta);
  }
  public static Annotation create_Annotation(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> turnId, dafny.DafnySequence<? extends dafny.CodePoint> resource, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> range, boolean resolved, dafny.DafnySequence<? extends Entry> entries, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return create(id, turnId, resource, range, resolved, entries, meta);
  }
  public boolean is_Annotation() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_id() {
    return this._id;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_turnId() {
    return this._turnId;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_resource() {
    return this._resource;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_range() {
    return this._range;
  }
  public boolean dtor_resolved() {
    return this._resolved;
  }
  public dafny.DafnySequence<? extends Entry> dtor_entries() {
    return this._entries;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_meta() {
    return this._meta;
  }
}
