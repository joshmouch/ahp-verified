// Class TerminalState
// Dafny class TerminalState compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalState {
  public dafny.DafnySequence<? extends dafny.CodePoint> _title;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _cwd;
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _cols;
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _rows;
  public dafny.DafnySequence<? extends Part> _content;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _claim;
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _exitCode;
  public agency.open.ahp.AhpSkeleton.Option<Boolean> _supportsCommandDetection;
  public TerminalState (dafny.DafnySequence<? extends dafny.CodePoint> title, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> cwd, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> cols, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> rows, dafny.DafnySequence<? extends Part> content, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> claim, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> exitCode, agency.open.ahp.AhpSkeleton.Option<Boolean> supportsCommandDetection) {
    this._title = title;
    this._cwd = cwd;
    this._cols = cols;
    this._rows = rows;
    this._content = content;
    this._claim = claim;
    this._exitCode = exitCode;
    this._supportsCommandDetection = supportsCommandDetection;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalState o = (TerminalState)other;
    return true && java.util.Objects.equals(this._title, o._title) && java.util.Objects.equals(this._cwd, o._cwd) && java.util.Objects.equals(this._cols, o._cols) && java.util.Objects.equals(this._rows, o._rows) && java.util.Objects.equals(this._content, o._content) && java.util.Objects.equals(this._claim, o._claim) && java.util.Objects.equals(this._exitCode, o._exitCode) && java.util.Objects.equals(this._supportsCommandDetection, o._supportsCommandDetection);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._title);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._cwd);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._cols);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._rows);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._content);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._claim);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._exitCode);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._supportsCommandDetection);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalState.TerminalState");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._title));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._cwd));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._cols));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._rows));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._content));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._claim));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._exitCode));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._supportsCommandDetection));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<TerminalState> _TYPE = dafny.TypeDescriptor.<TerminalState>referenceWithInitializer(TerminalState.class, () -> TerminalState.Default());
  public static dafny.TypeDescriptor<TerminalState> _typeDescriptor() {
    return (dafny.TypeDescriptor<TerminalState>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final TerminalState theDefault = TerminalState.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>Default(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>Default(dafny.TypeDescriptor.BIG_INTEGER), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>Default(dafny.TypeDescriptor.BIG_INTEGER), dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>Default(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>Default(dafny.TypeDescriptor.BIG_INTEGER), agency.open.ahp.AhpSkeleton.Option.<Boolean>Default(dafny.TypeDescriptor.BOOLEAN));
  public static TerminalState Default() {
    return theDefault;
  }
  public static TerminalState create(dafny.DafnySequence<? extends dafny.CodePoint> title, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> cwd, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> cols, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> rows, dafny.DafnySequence<? extends Part> content, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> claim, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> exitCode, agency.open.ahp.AhpSkeleton.Option<Boolean> supportsCommandDetection) {
    return new TerminalState(title, cwd, cols, rows, content, claim, exitCode, supportsCommandDetection);
  }
  public static TerminalState create_TerminalState(dafny.DafnySequence<? extends dafny.CodePoint> title, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> cwd, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> cols, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> rows, dafny.DafnySequence<? extends Part> content, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> claim, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> exitCode, agency.open.ahp.AhpSkeleton.Option<Boolean> supportsCommandDetection) {
    return create(title, cwd, cols, rows, content, claim, exitCode, supportsCommandDetection);
  }
  public boolean is_TerminalState() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_title() {
    return this._title;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_cwd() {
    return this._cwd;
  }
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> dtor_cols() {
    return this._cols;
  }
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> dtor_rows() {
    return this._rows;
  }
  public dafny.DafnySequence<? extends Part> dtor_content() {
    return this._content;
  }
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> dtor_claim() {
    return this._claim;
  }
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> dtor_exitCode() {
    return this._exitCode;
  }
  public agency.open.ahp.AhpSkeleton.Option<Boolean> dtor_supportsCommandDetection() {
    return this._supportsCommandDetection;
  }
}
