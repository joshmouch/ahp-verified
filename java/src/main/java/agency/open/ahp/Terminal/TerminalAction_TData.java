// Class TerminalAction_TData
// Dafny class TerminalAction_TData compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalAction_TData extends TerminalAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _data;
  public TerminalAction_TData (dafny.DafnySequence<? extends dafny.CodePoint> data) {
    super();
    this._data = data;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalAction_TData o = (TerminalAction_TData)other;
    return true && java.util.Objects.equals(this._data, o._data);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 4;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._data);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalAction.TData");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._data));
    s.append(")");
    return s.toString();
  }
}
