// Class SessionAction_DefaultChatChanged
// Dafny class SessionAction_DefaultChatChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_DefaultChatChanged extends SessionAction {
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _chat;
  public SessionAction_DefaultChatChanged (agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> chat) {
    super();
    this._chat = chat;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_DefaultChatChanged o = (SessionAction_DefaultChatChanged)other;
    return true && java.util.Objects.equals(this._chat, o._chat);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 12;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._chat);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.DefaultChatChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._chat));
    s.append(")");
    return s.toString();
  }
}
