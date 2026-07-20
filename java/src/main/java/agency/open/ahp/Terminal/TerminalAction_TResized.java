// Class TerminalAction_TResized
// Dafny class TerminalAction_TResized compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalAction_TResized extends TerminalAction {
  public java.math.BigInteger _cols;
  public java.math.BigInteger _rows;
  public TerminalAction_TResized (java.math.BigInteger cols, java.math.BigInteger rows) {
    super();
    this._cols = cols;
    this._rows = rows;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalAction_TResized o = (TerminalAction_TResized)other;
    return true && java.util.Objects.equals(this._cols, o._cols) && java.util.Objects.equals(this._rows, o._rows);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 2;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._cols);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._rows);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalAction.TResized");
    s.append("(");
    s.append(dafny.Helpers.toString(this._cols));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._rows));
    s.append(")");
    return s.toString();
  }
}
