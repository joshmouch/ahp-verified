// Class ChatAction_CTurnsLoaded
// Dafny class ChatAction_CTurnsLoaded compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CTurnsLoaded extends ChatAction {
  public dafny.DafnySequence<? extends Turn> _loaded;
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _cursor;
  public ChatAction_CTurnsLoaded (dafny.DafnySequence<? extends Turn> loaded, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> cursor) {
    super();
    this._loaded = loaded;
    this._cursor = cursor;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CTurnsLoaded o = (ChatAction_CTurnsLoaded)other;
    return true && java.util.Objects.equals(this._loaded, o._loaded) && java.util.Objects.equals(this._cursor, o._cursor);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 21;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._loaded);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._cursor);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CTurnsLoaded");
    s.append("(");
    s.append(dafny.Helpers.toString(this._loaded));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._cursor));
    s.append(")");
    return s.toString();
  }
}
