// Class TerminalAction_TUnknown
// Dafny class TerminalAction_TUnknown compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalAction_TUnknown extends TerminalAction {
  public agency.open.ahp.ConfluxCodec.Json _raw;
  public TerminalAction_TUnknown (agency.open.ahp.ConfluxCodec.Json raw) {
    super();
    this._raw = raw;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalAction_TUnknown o = (TerminalAction_TUnknown)other;
    return true && java.util.Objects.equals(this._raw, o._raw);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 11;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._raw);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalAction.TUnknown");
    s.append("(");
    s.append(dafny.Helpers.toString(this._raw));
    s.append(")");
    return s.toString();
  }
}
