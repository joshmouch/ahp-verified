// Class Part_Command
// Dafny class Part_Command compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class Part_Command extends Part {
  public dafny.DafnySequence<? extends dafny.CodePoint> _commandId;
  public dafny.DafnySequence<? extends dafny.CodePoint> _commandLine;
  public dafny.DafnySequence<? extends dafny.CodePoint> _output;
  public java.math.BigInteger _timestamp;
  public boolean _isComplete;
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _exitCode;
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _durationMs;
  public Part_Command (dafny.DafnySequence<? extends dafny.CodePoint> commandId, dafny.DafnySequence<? extends dafny.CodePoint> commandLine, dafny.DafnySequence<? extends dafny.CodePoint> output, java.math.BigInteger timestamp, boolean isComplete, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> exitCode, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> durationMs) {
    super();
    this._commandId = commandId;
    this._commandLine = commandLine;
    this._output = output;
    this._timestamp = timestamp;
    this._isComplete = isComplete;
    this._exitCode = exitCode;
    this._durationMs = durationMs;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Part_Command o = (Part_Command)other;
    return true && java.util.Objects.equals(this._commandId, o._commandId) && java.util.Objects.equals(this._commandLine, o._commandLine) && java.util.Objects.equals(this._output, o._output) && java.util.Objects.equals(this._timestamp, o._timestamp) && this._isComplete == o._isComplete && java.util.Objects.equals(this._exitCode, o._exitCode) && java.util.Objects.equals(this._durationMs, o._durationMs);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._commandId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._commandLine);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._output);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._timestamp);
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._isComplete);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._exitCode);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._durationMs);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.Part.Command");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._commandId));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._commandLine));
    s.append(", ");
    s.append(dafny.Helpers.ToStringLiteral(this._output));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._timestamp));
    s.append(", ");
    s.append(this._isComplete);
    s.append(", ");
    s.append(dafny.Helpers.toString(this._exitCode));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._durationMs));
    s.append(")");
    return s.toString();
  }
}
