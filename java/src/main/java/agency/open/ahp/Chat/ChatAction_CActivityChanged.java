// Class ChatAction_CActivityChanged
// Dafny class ChatAction_CActivityChanged compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CActivityChanged extends ChatAction {
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _activity;
  public ChatAction_CActivityChanged (agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> activity) {
    super();
    this._activity = activity;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CActivityChanged o = (ChatAction_CActivityChanged)other;
    return true && java.util.Objects.equals(this._activity, o._activity);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 1;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._activity);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CActivityChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._activity));
    s.append(")");
    return s.toString();
  }
}
