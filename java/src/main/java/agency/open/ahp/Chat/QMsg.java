// Class QMsg
// Dafny class QMsg compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class QMsg {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public agency.open.ahp.ConfluxCodec.Json _raw;
  public QMsg (dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.ConfluxCodec.Json raw) {
    this._id = id;
    this._raw = raw;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    QMsg o = (QMsg)other;
    return true && java.util.Objects.equals(this._id, o._id) && java.util.Objects.equals(this._raw, o._raw);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._raw);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.QMsg.QMsg");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._raw));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<QMsg> _TYPE = dafny.TypeDescriptor.<QMsg>referenceWithInitializer(QMsg.class, () -> QMsg.Default());
  public static dafny.TypeDescriptor<QMsg> _typeDescriptor() {
    return (dafny.TypeDescriptor<QMsg>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final QMsg theDefault = QMsg.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.ConfluxCodec.Json.Default());
  public static QMsg Default() {
    return theDefault;
  }
  public static QMsg create(dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.ConfluxCodec.Json raw) {
    return new QMsg(id, raw);
  }
  public static QMsg create_QMsg(dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.ConfluxCodec.Json raw) {
    return create(id, raw);
  }
  public boolean is_QMsg() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_id() {
    return this._id;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    return this._raw;
  }
}
