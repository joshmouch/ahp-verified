// Class Turn
// Dafny class Turn compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class Turn {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public agency.open.ahp.ConfluxCodec.Json _message;
  public dafny.DafnySequence<? extends Part> _parts;
  public dafny.DafnySequence<? extends dafny.CodePoint> _state;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _usage;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _error;
  public Turn (dafny.DafnySequence<? extends dafny.CodePoint> turnId, agency.open.ahp.ConfluxCodec.Json message, dafny.DafnySequence<? extends Part> parts, dafny.DafnySequence<? extends dafny.CodePoint> state, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> usage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    this._turnId = turnId;
    this._message = message;
    this._parts = parts;
    this._state = state;
    this._usage = usage;
    this._error = error;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Turn o = (Turn)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._message, o._message) && java.util.Objects.equals(this._parts, o._parts) && java.util.Objects.equals(this._state, o._state) && java.util.Objects.equals(this._usage, o._usage) && java.util.Objects.equals(this._error, o._error);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._message);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._parts);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._state);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._usage);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._error);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.Turn.Turn");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._message));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._parts));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._state));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._usage));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._error));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<Turn> _TYPE = dafny.TypeDescriptor.<Turn>referenceWithInitializer(Turn.class, () -> Turn.Default());
  public static dafny.TypeDescriptor<Turn> _typeDescriptor() {
    return (dafny.TypeDescriptor<Turn>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final Turn theDefault = Turn.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.ConfluxCodec.Json.Default(), dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
  public static Turn Default() {
    return theDefault;
  }
  public static Turn create(dafny.DafnySequence<? extends dafny.CodePoint> turnId, agency.open.ahp.ConfluxCodec.Json message, dafny.DafnySequence<? extends Part> parts, dafny.DafnySequence<? extends dafny.CodePoint> state, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> usage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    return new Turn(turnId, message, parts, state, usage, error);
  }
  public static Turn create_Turn(dafny.DafnySequence<? extends dafny.CodePoint> turnId, agency.open.ahp.ConfluxCodec.Json message, dafny.DafnySequence<? extends Part> parts, dafny.DafnySequence<? extends dafny.CodePoint> state, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> usage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error) {
    return create(turnId, message, parts, state, usage, error);
  }
  public boolean is_Turn() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_turnId() {
    return this._turnId;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_message() {
    return this._message;
  }
  public dafny.DafnySequence<? extends Part> dtor_parts() {
    return this._parts;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_state() {
    return this._state;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_usage() {
    return this._usage;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_error() {
    return this._error;
  }
}
