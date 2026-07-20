// Class ChatAction_CQueuedReordered
// Dafny class ChatAction_CQueuedReordered compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CQueuedReordered extends ChatAction {
  public dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> _order;
  public ChatAction_CQueuedReordered (dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> order) {
    super();
    this._order = order;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CQueuedReordered o = (ChatAction_CQueuedReordered)other;
    return true && java.util.Objects.equals(this._order, o._order);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 20;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._order);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CQueuedReordered");
    s.append("(");
    s.append(dafny.Helpers.toString(this._order));
    s.append(")");
    return s.toString();
  }
}
