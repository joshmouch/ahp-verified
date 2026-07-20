// Class ReduceOutcome_OutOfScope
// Dafny class ReduceOutcome_OutOfScope compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class ReduceOutcome_OutOfScope extends ReduceOutcome {
  public ReduceOutcome_OutOfScope () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ReduceOutcome_OutOfScope o = (ReduceOutcome_OutOfScope)other;
    return true;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 2;
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.ReduceOutcome.OutOfScope");
    return s.toString();
  }
}
