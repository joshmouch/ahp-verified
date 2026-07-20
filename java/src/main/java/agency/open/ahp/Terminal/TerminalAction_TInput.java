// Class TerminalAction_TInput
// Dafny class TerminalAction_TInput compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalAction_TInput extends TerminalAction {
  public TerminalAction_TInput () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalAction_TInput o = (TerminalAction_TInput)other;
    return true;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 10;
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalAction.TInput");
    return s.toString();
  }
}
