// Class CanvasAdmission
// Dafny class CanvasAdmission compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class CanvasAdmission {
  public CanvasAdmission() {
  }
  private static final dafny.TypeDescriptor<CanvasAdmission> _TYPE = dafny.TypeDescriptor.<CanvasAdmission>referenceWithInitializer(CanvasAdmission.class, () -> CanvasAdmission.Default());
  public static dafny.TypeDescriptor<CanvasAdmission> _typeDescriptor() {
    return (dafny.TypeDescriptor<CanvasAdmission>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final CanvasAdmission theDefault = CanvasAdmission.create_CanvasSupported();
  public static CanvasAdmission Default() {
    return theDefault;
  }
  public static CanvasAdmission create_CanvasSupported() {
    return new CanvasAdmission_CanvasSupported();
  }
  public static CanvasAdmission create_CanvasUnsupported() {
    return new CanvasAdmission_CanvasUnsupported();
  }
  public static CanvasAdmission create_CanvasHandlerUnavailable() {
    return new CanvasAdmission_CanvasHandlerUnavailable();
  }
  public static CanvasAdmission create_CanvasUnauthorized() {
    return new CanvasAdmission_CanvasUnauthorized();
  }
  public static CanvasAdmission create_CanvasUriCollision() {
    return new CanvasAdmission_CanvasUriCollision();
  }
  public boolean is_CanvasSupported() { return this instanceof CanvasAdmission_CanvasSupported; }
  public boolean is_CanvasUnsupported() { return this instanceof CanvasAdmission_CanvasUnsupported; }
  public boolean is_CanvasHandlerUnavailable() { return this instanceof CanvasAdmission_CanvasHandlerUnavailable; }
  public boolean is_CanvasUnauthorized() { return this instanceof CanvasAdmission_CanvasUnauthorized; }
  public boolean is_CanvasUriCollision() { return this instanceof CanvasAdmission_CanvasUriCollision; }
  public static java.util.ArrayList<CanvasAdmission> AllSingletonConstructors() {
    java.util.ArrayList<CanvasAdmission> singleton_iterator = new java.util.ArrayList<>();
    singleton_iterator.add(new CanvasAdmission_CanvasSupported());
    singleton_iterator.add(new CanvasAdmission_CanvasUnsupported());
    singleton_iterator.add(new CanvasAdmission_CanvasHandlerUnavailable());
    singleton_iterator.add(new CanvasAdmission_CanvasUnauthorized());
    singleton_iterator.add(new CanvasAdmission_CanvasUriCollision());
    return singleton_iterator;
  }
}
