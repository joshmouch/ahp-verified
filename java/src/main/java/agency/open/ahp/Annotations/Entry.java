// Class Entry
// Dafny class Entry compiled into Java
package agency.open.ahp.Annotations;

@SuppressWarnings({"unchecked", "deprecation"})
public class Entry {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public dafny.DafnySequence<? extends dafny.CodePoint> _text;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _meta;
  public Entry (dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> text, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    this._id = id;
    this._text = text;
    this._meta = meta;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Entry o = (Entry)other;
    return true && java.util.Objects.equals(this._id, o._id) && java.util.Objects.equals(this._text, o._text) && java.util.Objects.equals(this._meta, o._meta);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._text);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._meta);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Annotations.Entry.Entry");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._text));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._meta));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<Entry> _TYPE = dafny.TypeDescriptor.<Entry>referenceWithInitializer(Entry.class, () -> Entry.Default());
  public static dafny.TypeDescriptor<Entry> _typeDescriptor() {
    return (dafny.TypeDescriptor<Entry>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final Entry theDefault = Entry.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
  public static Entry Default() {
    return theDefault;
  }
  public static Entry create(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> text, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return new Entry(id, text, meta);
  }
  public static Entry create_Entry(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> text, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta) {
    return create(id, text, meta);
  }
  public boolean is_Entry() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_id() {
    return this._id;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_text() {
    return this._text;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_meta() {
    return this._meta;
  }
}
