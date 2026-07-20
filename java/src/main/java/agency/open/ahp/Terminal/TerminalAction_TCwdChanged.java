// Class TerminalAction_TCwdChanged
// Dafny class TerminalAction_TCwdChanged compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalAction_TCwdChanged extends TerminalAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _cwd;
  public TerminalAction_TCwdChanged (dafny.DafnySequence<? extends dafny.CodePoint> cwd) {
    super();
    this._cwd = cwd;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalAction_TCwdChanged o = (TerminalAction_TCwdChanged)other;
    return true && java.util.Objects.equals(this._cwd, o._cwd);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._cwd);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalAction.TCwdChanged");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._cwd));
    s.append(")");
    return s.toString();
  }
}
