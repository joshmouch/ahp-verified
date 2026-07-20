// Class ChatAction_CUsage
// Dafny class ChatAction_CUsage compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public class ChatAction_CUsage extends ChatAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _turnId;
  public agency.open.ahp.ConfluxCodec.Json _usage;
  public ChatAction_CUsage (dafny.DafnySequence<? extends dafny.CodePoint> turnId, agency.open.ahp.ConfluxCodec.Json usage) {
    super();
    this._turnId = turnId;
    this._usage = usage;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    ChatAction_CUsage o = (ChatAction_CUsage)other;
    return true && java.util.Objects.equals(this._turnId, o._turnId) && java.util.Objects.equals(this._usage, o._usage);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 5;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._turnId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._usage);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Chat.ChatAction.CUsage");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._turnId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._usage));
    s.append(")");
    return s.toString();
  }
}
