// Class Part_PInputRequest
// Dafny class Part_PInputRequest compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class Part_PInputRequest extends Part {
  public InputReq _req;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _response;
  public Part_PInputRequest (InputReq req, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> response) {
    super();
    this._req = req;
    this._response = response;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Part_PInputRequest o = (Part_PInputRequest)other;
    return true && java.util.Objects.equals(this._req, o._req) && java.util.Objects.equals(this._response, o._response);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 3;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._req);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._response);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.Part.PInputRequest");
    s.append("(");
    s.append(dafny.Helpers.toString(this._req));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._response));
    s.append(")");
    return s.toString();
  }
}
