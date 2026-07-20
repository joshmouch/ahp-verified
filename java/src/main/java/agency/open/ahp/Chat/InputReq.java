// Class InputReq
// Dafny class InputReq compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class InputReq {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _answers;
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _open;
  public InputReq (dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> answers, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> open) {
    this._id = id;
    this._answers = answers;
    this._open = open;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    InputReq o = (InputReq)other;
    return true && java.util.Objects.equals(this._id, o._id) && java.util.Objects.equals(this._answers, o._answers) && java.util.Objects.equals(this._open, o._open);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._answers);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._open);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.InputReq.InputReq");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._answers));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._open));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<InputReq> _TYPE = dafny.TypeDescriptor.<InputReq>referenceWithInitializer(InputReq.class, () -> InputReq.Default());
  public static dafny.TypeDescriptor<InputReq> _typeDescriptor() {
    return (dafny.TypeDescriptor<InputReq>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final InputReq theDefault = InputReq.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnyMap.<dafny.DafnySequence<? extends dafny.CodePoint>,agency.open.ahp.ConfluxCodec.Json> empty(), dafny.DafnyMap.<dafny.DafnySequence<? extends dafny.CodePoint>,agency.open.ahp.ConfluxCodec.Json> empty());
  public static InputReq Default() {
    return theDefault;
  }
  public static InputReq create(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> answers, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> open) {
    return new InputReq(id, answers, open);
  }
  public static InputReq create_InputReq(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> answers, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> open) {
    return create(id, answers, open);
  }
  public boolean is_InputReq() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_id() {
    return this._id;
  }
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> dtor_answers() {
    return this._answers;
  }
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> dtor_open() {
    return this._open;
  }
}
