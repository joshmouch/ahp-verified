// Class ReduceOutcome_NoOp
// Dafny class ReduceOutcome_NoOp compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class ReduceOutcome_NoOp extends ReduceOutcome {
  public ReduceOutcome_NoOp () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ReduceOutcome_NoOp o = (ReduceOutcome_NoOp)other;
    return true;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.ReduceOutcome.NoOp");
    return s.toString();
  }
}
