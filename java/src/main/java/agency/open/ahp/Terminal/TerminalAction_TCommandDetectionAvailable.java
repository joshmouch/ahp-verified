// Class TerminalAction_TCommandDetectionAvailable
// Dafny class TerminalAction_TCommandDetectionAvailable compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class TerminalAction_TCommandDetectionAvailable extends TerminalAction {
  public TerminalAction_TCommandDetectionAvailable () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    TerminalAction_TCommandDetectionAvailable o = (TerminalAction_TCommandDetectionAvailable)other;
    return true;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 7;
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Terminal.TerminalAction.TCommandDetectionAvailable");
    return s.toString();
  }
}
