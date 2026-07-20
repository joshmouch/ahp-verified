// Class ChatAction_CDraftChanged
// Dafny class ChatAction_CDraftChanged compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CDraftChanged extends ChatAction {
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _draft;
  public ChatAction_CDraftChanged (agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> draft) {
    super();
    this._draft = draft;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CDraftChanged o = (ChatAction_CDraftChanged)other;
    return true && java.util.Objects.equals(this._draft, o._draft);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._draft);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CDraftChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._draft));
    s.append(")");
    return s.toString();
  }
}
