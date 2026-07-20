// Class ReduceOutcome
// Dafny class ReduceOutcome compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class ReduceOutcome {
  public ReduceOutcome() {
  }
  private static final dafny.TypeDescriptor<ReduceOutcome> _TYPE = dafny.TypeDescriptor.<ReduceOutcome>referenceWithInitializer(ReduceOutcome.class, () -> ReduceOutcome.Default());
  public static dafny.TypeDescriptor<ReduceOutcome> _typeDescriptor() {
    return (dafny.TypeDescriptor<ReduceOutcome>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final ReduceOutcome theDefault = ReduceOutcome.create_Applied();
  public static ReduceOutcome Default() {
    return theDefault;
  }
  public static ReduceOutcome create_Applied() {
    return new ReduceOutcome_Applied();
  }
  public static ReduceOutcome create_NoOp() {
    return new ReduceOutcome_NoOp();
  }
  public static ReduceOutcome create_OutOfScope() {
    return new ReduceOutcome_OutOfScope();
  }
  public boolean is_Applied() { return this instanceof ReduceOutcome_Applied; }
  public boolean is_NoOp() { return this instanceof ReduceOutcome_NoOp; }
  public boolean is_OutOfScope() { return this instanceof ReduceOutcome_OutOfScope; }
  public static java.util.ArrayList<ReduceOutcome> AllSingletonConstructors() {
    java.util.ArrayList<ReduceOutcome> singleton_iterator = new java.util.ArrayList<>();
    singleton_iterator.add(new ReduceOutcome_Applied());
    singleton_iterator.add(new ReduceOutcome_NoOp());
    singleton_iterator.add(new ReduceOutcome_OutOfScope());
    return singleton_iterator;
  }
}
