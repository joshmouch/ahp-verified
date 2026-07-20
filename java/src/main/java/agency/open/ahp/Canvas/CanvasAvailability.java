// Class CanvasAvailability
// Dafny class CanvasAvailability compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class CanvasAvailability {
  public CanvasAvailability() {
  }
  private static final dafny.TypeDescriptor<CanvasAvailability> _TYPE = dafny.TypeDescriptor.<CanvasAvailability>referenceWithInitializer(CanvasAvailability.class, () -> CanvasAvailability.Default());
  public static dafny.TypeDescriptor<CanvasAvailability> _typeDescriptor() {
    return (dafny.TypeDescriptor<CanvasAvailability>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final CanvasAvailability theDefault = CanvasAvailability.create_Ready();
  public static CanvasAvailability Default() {
    return theDefault;
  }
  public static CanvasAvailability create_Ready() {
    return new CanvasAvailability_Ready();
  }
  public static CanvasAvailability create_Stale() {
    return new CanvasAvailability_Stale();
  }
  public boolean is_Ready() { return this instanceof CanvasAvailability_Ready; }
  public boolean is_Stale() { return this instanceof CanvasAvailability_Stale; }
  public static java.util.ArrayList<CanvasAvailability> AllSingletonConstructors() {
    java.util.ArrayList<CanvasAvailability> singleton_iterator = new java.util.ArrayList<>();
    singleton_iterator.add(new CanvasAvailability_Ready());
    singleton_iterator.add(new CanvasAvailability_Stale());
    return singleton_iterator;
  }
}
