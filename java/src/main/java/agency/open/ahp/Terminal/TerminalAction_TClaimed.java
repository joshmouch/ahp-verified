// Class TerminalAction_TClaimed
// Dafny class TerminalAction_TClaimed compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalAction_TClaimed extends TerminalAction {
  public agency.open.ahp.ConfluxCodec.Json _claim;
  public TerminalAction_TClaimed (agency.open.ahp.ConfluxCodec.Json claim) {
    super();
    this._claim = claim;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalAction_TClaimed o = (TerminalAction_TClaimed)other;
    return true && java.util.Objects.equals(this._claim, o._claim);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 6;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._claim);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalAction.TClaimed");
    s.append("(");
    s.append(dafny.Helpers.toString(this._claim));
    s.append(")");
    return s.toString();
  }
}
