// Class TerminalAction_TCommandExecuted
// Dafny class TerminalAction_TCommandExecuted compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalAction_TCommandExecuted extends TerminalAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _commandId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _commandLine;
  public java.math.BigInteger _timestamp;
  public TerminalAction_TCommandExecuted (dafny.DafnySequence<? extends dafny.CodePoint> commandId, dafny.DafnySequence<? extends dafny.CodePoint> commandLine, java.math.BigInteger timestamp) {
    super();
    this._commandId = commandId;
    this._commandLine = commandLine;
    this._timestamp = timestamp;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalAction_TCommandExecuted o = (TerminalAction_TCommandExecuted)other;
    return true && java.util.Objects.equals(this._commandId, o._commandId) && java.util.Objects.equals(this._commandLine, o._commandLine) && java.util.Objects.equals(this._timestamp, o._timestamp);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 8;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._commandId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._commandLine);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._timestamp);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalAction.TCommandExecuted");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._commandId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._commandLine));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._timestamp));
    s.append(")");
    return s.toString();
  }
}
