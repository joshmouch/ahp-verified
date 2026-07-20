// Class SessionAction_McpServerStateChanged
// Dafny class SessionAction_McpServerStateChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_McpServerStateChanged extends SessionAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _id;
  public agency.open.ahp.ConfluxCodec.Json _state;
  public agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _channel;
  public SessionAction_McpServerStateChanged (dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.ConfluxCodec.Json state, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> channel) {
    super();
    this._id = id;
    this._state = state;
    this._channel = channel;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_McpServerStateChanged o = (SessionAction_McpServerStateChanged)other;
    return true && java.util.Objects.equals(this._id, o._id) && java.util.Objects.equals(this._state, o._state) && java.util.Objects.equals(this._channel, o._channel);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 24;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._id);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._state);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._channel);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.McpServerStateChanged");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._id));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._state));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._channel));
    s.append(")");
    return s.toString();
  }
}
