// Class TerminalAction
// Dafny class TerminalAction compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class TerminalAction {
  public TerminalAction() {
  }
  private static final dafny.TypeDescriptor<TerminalAction> _TYPE = dafny.TypeDescriptor.<TerminalAction>referenceWithInitializer(TerminalAction.class, () -> TerminalAction.Default());
  public static dafny.TypeDescriptor<TerminalAction> _typeDescriptor() {
    return (dafny.TypeDescriptor<TerminalAction>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final TerminalAction theDefault = TerminalAction.create_TCwdChanged(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR));
  public static TerminalAction Default() {
    return theDefault;
  }
  public static TerminalAction create_TCwdChanged(dafny.DafnySequence<? extends dafny.CodePoint> cwd) {
    return new TerminalAction_TCwdChanged(cwd);
  }
  public static TerminalAction create_TTitleChanged(dafny.DafnySequence<? extends dafny.CodePoint> title) {
    return new TerminalAction_TTitleChanged(title);
  }
  public static TerminalAction create_TResized(java.math.BigInteger cols, java.math.BigInteger rows) {
    return new TerminalAction_TResized(cols, rows);
  }
  public static TerminalAction create_TExited(java.math.BigInteger code) {
    return new TerminalAction_TExited(code);
  }
  public static TerminalAction create_TData(dafny.DafnySequence<? extends dafny.CodePoint> data) {
    return new TerminalAction_TData(data);
  }
  public static TerminalAction create_TCleared() {
    return new TerminalAction_TCleared();
  }
  public static TerminalAction create_TClaimed(agency.open.ahp.ConfluxCodec.Json claim) {
    return new TerminalAction_TClaimed(claim);
  }
  public static TerminalAction create_TCommandDetectionAvailable() {
    return new TerminalAction_TCommandDetectionAvailable();
  }
  public static TerminalAction create_TCommandExecuted(dafny.DafnySequence<? extends dafny.CodePoint> commandId, dafny.DafnySequence<? extends dafny.CodePoint> commandLine, java.math.BigInteger timestamp) {
    return new TerminalAction_TCommandExecuted(commandId, commandLine, timestamp);
  }
  public static TerminalAction create_TCommandFinished(dafny.DafnySequence<? extends dafny.CodePoint> commandId, java.math.BigInteger code, java.math.BigInteger durationMs) {
    return new TerminalAction_TCommandFinished(commandId, code, durationMs);
  }
  public static TerminalAction create_TInput() {
    return new TerminalAction_TInput();
  }
  public static TerminalAction create_TUnknown(agency.open.ahp.ConfluxCodec.Json raw) {
    return new TerminalAction_TUnknown(raw);
  }
  public boolean is_TCwdChanged() { return this instanceof TerminalAction_TCwdChanged; }
  public boolean is_TTitleChanged() { return this instanceof TerminalAction_TTitleChanged; }
  public boolean is_TResized() { return this instanceof TerminalAction_TResized; }
  public boolean is_TExited() { return this instanceof TerminalAction_TExited; }
  public boolean is_TData() { return this instanceof TerminalAction_TData; }
  public boolean is_TCleared() { return this instanceof TerminalAction_TCleared; }
  public boolean is_TClaimed() { return this instanceof TerminalAction_TClaimed; }
  public boolean is_TCommandDetectionAvailable() { return this instanceof TerminalAction_TCommandDetectionAvailable; }
  public boolean is_TCommandExecuted() { return this instanceof TerminalAction_TCommandExecuted; }
  public boolean is_TCommandFinished() { return this instanceof TerminalAction_TCommandFinished; }
  public boolean is_TInput() { return this instanceof TerminalAction_TInput; }
  public boolean is_TUnknown() { return this instanceof TerminalAction_TUnknown; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_cwd() {
    TerminalAction d = this;
    return ((TerminalAction_TCwdChanged)d)._cwd;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_title() {
    TerminalAction d = this;
    return ((TerminalAction_TTitleChanged)d)._title;
  }
  public java.math.BigInteger dtor_cols() {
    TerminalAction d = this;
    return ((TerminalAction_TResized)d)._cols;
  }
  public java.math.BigInteger dtor_rows() {
    TerminalAction d = this;
    return ((TerminalAction_TResized)d)._rows;
  }
  public java.math.BigInteger dtor_code() {
    TerminalAction d = this;
    if (d instanceof TerminalAction_TExited) { return ((TerminalAction_TExited)d)._code; }
    return ((TerminalAction_TCommandFinished)d)._code;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_data() {
    TerminalAction d = this;
    return ((TerminalAction_TData)d)._data;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_claim() {
    TerminalAction d = this;
    return ((TerminalAction_TClaimed)d)._claim;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_commandId() {
    TerminalAction d = this;
    if (d instanceof TerminalAction_TCommandExecuted) { return ((TerminalAction_TCommandExecuted)d)._commandId; }
    return ((TerminalAction_TCommandFinished)d)._commandId;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_commandLine() {
    TerminalAction d = this;
    return ((TerminalAction_TCommandExecuted)d)._commandLine;
  }
  public java.math.BigInteger dtor_timestamp() {
    TerminalAction d = this;
    return ((TerminalAction_TCommandExecuted)d)._timestamp;
  }
  public java.math.BigInteger dtor_durationMs() {
    TerminalAction d = this;
    return ((TerminalAction_TCommandFinished)d)._durationMs;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    TerminalAction d = this;
    return ((TerminalAction_TUnknown)d)._raw;
  }
}
