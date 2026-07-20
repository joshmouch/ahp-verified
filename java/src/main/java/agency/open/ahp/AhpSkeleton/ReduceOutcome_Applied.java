// Class ReduceOutcome_Applied
// Dafny class ReduceOutcome_Applied compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class ReduceOutcome_Applied extends ReduceOutcome {
  public ReduceOutcome_Applied () {
    super();
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ReduceOutcome_Applied o = (ReduceOutcome_Applied)other;
    return true;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("AhpSkeleton.ReduceOutcome.Applied");
    return s.toString();
  }
}
