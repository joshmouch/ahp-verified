// Class ChatAction_CInputRequested
// Dafny class ChatAction_CInputRequested compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CInputRequested extends ChatAction {
  public InputReq _req;
  public ChatAction_CInputRequested (InputReq req) {
    super();
    this._req = req;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CInputRequested o = (ChatAction_CInputRequested)other;
    return true && java.util.Objects.equals(this._req, o._req);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 22;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._req);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CInputRequested");
    s.append("(");
    s.append(dafny.Helpers.toString(this._req));
    s.append(")");
    return s.toString();
  }
}
