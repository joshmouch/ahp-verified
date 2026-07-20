// Class CanvasError_ProviderError
// Dafny class CanvasError_ProviderError compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class CanvasError_ProviderError extends CanvasError {
  public java.math.BigInteger _errorCode;
  public CanvasProviderErrorData _data;
  public CanvasError_ProviderError (java.math.BigInteger errorCode, CanvasProviderErrorData data) {
    super();
    this._errorCode = errorCode;
    this._data = data;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    CanvasError_ProviderError o = (CanvasError_ProviderError)other;
    return true && java.util.Objects.equals(this._errorCode, o._errorCode) && java.util.Objects.equals(this._data, o._data);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._errorCode);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._data);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Canvas.CanvasError.ProviderError");
    s.append("(");
    s.append(dafny.Helpers.toString(this._errorCode));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._data));
    s.append(")");
    return s.toString();
  }
}
