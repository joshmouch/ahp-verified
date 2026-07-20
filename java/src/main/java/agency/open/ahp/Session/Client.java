// Class Client
// Dafny class Client compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class Client {
  public dafny.DafnySequence<? extends dafny.CodePoint> _clientId;
  public agency.open.ahp.ConfluxCodec.Json _raw;
  public Client (dafny.DafnySequence<? extends dafny.CodePoint> clientId, agency.open.ahp.ConfluxCodec.Json raw) {
    this._clientId = clientId;
    this._raw = raw;
  }

  @Override
  public boolean equals(Object other) {
    if (this == other) return true;
    if (other == null) return false;
    if (getClass() != other.getClass()) return false;
    Client o = (Client)other;
    return true && java.util.Objects.equals(this._clientId, o._clientId) && java.util.Objects.equals(this._raw, o._raw);
  }
  @Override
  public int hashCode() {
    long hash = 5381;
    hash = ((hash << 5) + hash) + 0;
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._clientId);
    hash = ((hash << 5) + hash) + java.util.Objects.hashCode(this._raw);
    return (int)hash;
  }

  @Override
  public String toString() {
    StringBuilder s = new StringBuilder();
    s.append("Session.Client.Client");
    s.append("(");
    s.append(dafny.Helpers.ToStringLiteral(this._clientId));
    s.append(", ");
    s.append(dafny.Helpers.toString(this._raw));
    s.append(")");
    return s.toString();
  }
  private static final dafny.TypeDescriptor<Client> _TYPE = dafny.TypeDescriptor.<Client>referenceWithInitializer(Client.class, () -> Client.Default());
  public static dafny.TypeDescriptor<Client> _typeDescriptor() {
    return (dafny.TypeDescriptor<Client>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final Client theDefault = Client.create(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.ConfluxCodec.Json.Default());
  public static Client Default() {
    return theDefault;
  }
  public static Client create(dafny.DafnySequence<? extends dafny.CodePoint> clientId, agency.open.ahp.ConfluxCodec.Json raw) {
    return new Client(clientId, raw);
  }
  public static Client create_Client(dafny.DafnySequence<? extends dafny.CodePoint> clientId, agency.open.ahp.ConfluxCodec.Json raw) {
    return create(clientId, raw);
  }
  public boolean is_Client() { return true; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_clientId() {
    return this._clientId;
  }
  public agency.open.ahp.ConfluxCodec.Json dtor_raw() {
    return this._raw;
  }
}
