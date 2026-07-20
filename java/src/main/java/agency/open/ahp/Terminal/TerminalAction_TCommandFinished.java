// Class TerminalAction_TCommandFinished
// Dafny class TerminalAction_TCommandFinished compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalAction_TCommandFinished extends TerminalAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _commandId;
  public java.math.BigInteger _code;
  public java.math.BigInteger _durationMs;
  public TerminalAction_TCommandFinished (dafny.DafnySequence<? extends dafny.CodePoint> commandId, java.math.BigInteger code, java.math.BigInteger durationMs) {
    super();
    this._commandId = commandId;
    this._code = code;
    this._durationMs = durationMs;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalAction_TCommandFinished o = (TerminalAction_TCommandFinished)other;
    return true && java.util.Objects.equals(this._commandId, o._commandId) && java.util.Objects.equals(this._code, o._code) && java.util.Objects.equals(this._durationMs, o._durationMs);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 9;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._commandId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._code);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._durationMs);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalAction.TCommandFinished");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._commandId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._code));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._durationMs));
    s.append(")");
    return s.toString();
  }
}
