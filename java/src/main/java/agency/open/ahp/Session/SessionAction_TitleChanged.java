// Class SessionAction_TitleChanged
// Dafny class SessionAction_TitleChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_TitleChanged extends SessionAction {
  public dafny.DafnySequence<? extends dafny.CodePoint> _title;
  public SessionAction_TitleChanged (dafny.DafnySequence<? extends dafny.CodePoint> title) {
    super();
    this._title = title;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_TitleChanged o = (SessionAction_TitleChanged)other;
    return true && java.util.Objects.equals(this._title, o._title);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 2;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._title);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.TitleChanged");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._title));
    s.append(")");
    return s.toString();
  }
}
