// Class TerminalAction_TTitleChanged
// Dafny class TerminalAction_TTitleChanged compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalAction_TTitleChanged extends TerminalAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _title;
  public TerminalAction_TTitleChanged (dafny.DafnySequence<? extends dafny.CodePoint> title) {
    super();
    this._title = title;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalAction_TTitleChanged o = (TerminalAction_TTitleChanged)other;
    return true && java.util.Objects.equals(this._title, o._title);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._title);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalAction.TTitleChanged");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._title));
    s.append(")");
    return s.toString();
  }
}
