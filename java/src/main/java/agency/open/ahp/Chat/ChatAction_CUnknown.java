// Class ChatAction_CUnknown
// Dafny class ChatAction_CUnknown compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CUnknown extends ChatAction {
  public agency.open.ahp.ConfluxCodec.Json _raw;
  public ChatAction_CUnknown (agency.open.ahp.ConfluxCodec.Json raw) {
    super();
    this._raw = raw;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CUnknown o = (ChatAction_CUnknown)other;
    return true && java.util.Objects.equals(this._raw, o._raw);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 27;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._raw);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CUnknown");
    s.append("(");
    s.append(dafny.Helpers.toString(this._raw));
    s.append(")");
    return s.toString();
  }
}
