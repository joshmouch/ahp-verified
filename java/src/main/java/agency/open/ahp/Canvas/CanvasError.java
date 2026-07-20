// Class CanvasError
// Dafny class CanvasError compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class CanvasError {
  public CanvasError() {
  }
  private static final dafny.TypeDescriptor<CanvasError> _TYPE = dafny.TypeDescriptor.<CanvasError>referenceWithInitializer(CanvasError.class, () -> CanvasError.Default());
  public static dafny.TypeDescriptor<CanvasError> _typeDescriptor() {
    return (dafny.TypeDescriptor<CanvasError>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final CanvasError theDefault = CanvasError.create_AdmissionError(CanvasAdmission.Default());
  public static CanvasError Default() {
    return theDefault;
  }
  public static CanvasError create_AdmissionError(CanvasAdmission reason) {
    return new CanvasError_AdmissionError(reason);
  }
  public static CanvasError create_ProviderError(java.math.BigInteger errorCode, CanvasProviderErrorData data) {
    return new CanvasError_ProviderError(errorCode, data);
  }
  public boolean is_AdmissionError() { return this instanceof CanvasError_AdmissionError; }
  public boolean is_ProviderError() { return this instanceof CanvasError_ProviderError; }
  public CanvasAdmission dtor_reason() {
    CanvasError d = this;
    return ((CanvasError_AdmissionError)d)._reason;
  }
  public java.math.BigInteger dtor_errorCode() {
    CanvasError d = this;
    return ((CanvasError_ProviderError)d)._errorCode;
  }
  public CanvasProviderErrorData dtor_data() {
    CanvasError d = this;
    return ((CanvasError_ProviderError)d)._data;
  }
}
