// Class SessionAction_ConfigChanged
// Dafny class SessionAction_ConfigChanged compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class SessionAction_ConfigChanged extends SessionAction {
  public dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _config;
  public boolean _replace;
  public SessionAction_ConfigChanged (dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> config, boolean replace) {
    super();
    this._config = config;
    this._replace = replace;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    SessionAction_ConfigChanged o = (SessionAction_ConfigChanged)other;
    return true && java.util.Objects.equals(this._config, o._config) && this._replace == o._replace;
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 8;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._config);
    hash = ((hash << 5) + hash) + Boolean.hashCode(this._replace);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.SessionAction.ConfigChanged");
    s.append("(");
    s.append(dafny.Helpers.toString(this._config));
    s.append(", ");
    s.append(this._replace);
    s.append(")");
    return s.toString();
  }
}
