// Class TerminalAction_TExited
// Dafny class TerminalAction_TExited compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalAction_TExited extends TerminalAction {
  public java.math.BigInteger _code;
  public TerminalAction_TExited (java.math.BigInteger code) {
    super();
    this._code = code;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalAction_TExited o = (TerminalAction_TExited)other;
    return true && java.util.Objects.equals(this._code, o._code);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 3;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._code);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalAction.TExited");
    s.append("(");
    s.append(dafny.Helpers.toString(this._code));
    s.append(")");
    return s.toString();
  }
}
