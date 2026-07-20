// .NET realization of the pure ConfluxCommandIdentityCapability extern functions.
// ConfluxJsonText remains the sole JSON text implementation; this adapter only supplies strict UTF-8 and SHA-256.
using System;
using System.Security.Cryptography;
using System.Text;
using Rune = Dafny.Rune;

namespace ConfluxCommandIdentityCapability {
  public partial class __default {
    private static readonly UTF8Encoding StrictUtf8 = new UTF8Encoding(false, true);

    private static Dafny.ISequence<byte> Bytes(byte[] value) => Dafny.Helpers.SeqFromArray(value);
    private static Dafny.ISequence<Rune> Text(string value) =>
      Dafny.Sequence<Rune>.UnicodeFromString(value ?? "");

    public static Dafny.ISequence<byte> CanonicalJsonBytes(ConfluxCodec._IJson value) {
      var text = ConfluxJsonText.__default.Stringify(value).ToVerbatimString(false)
        .Normalize(NormalizationForm.FormC);
      return Bytes(StrictUtf8.GetBytes(text));
    }

    public static ConfluxCodec._IOption<ConfluxCodec._IJson> DecodeCanonicalJson(
        Dafny.ISequence<byte> bytes) {
      try {
        var input = bytes.CloneAsArray();
        var text = StrictUtf8.GetString(input);
        var parsed = ConfluxJsonText.__default.ParseJsonChecked(Text(text));
        if (!parsed.is_Parsed) return ConfluxCodec.Option<ConfluxCodec._IJson>.create_None();
        var canonical = CanonicalJsonBytes(parsed.dtor_value).CloneAsArray();
        if (!CryptographicOperations.FixedTimeEquals(input, canonical)) {
          return ConfluxCodec.Option<ConfluxCodec._IJson>.create_None();
        }
        return ConfluxCodec.Option<ConfluxCodec._IJson>.create_Some(parsed.dtor_value);
      } catch (DecoderFallbackException) {
        return ConfluxCodec.Option<ConfluxCodec._IJson>.create_None();
      } catch (ArgumentException) {
        return ConfluxCodec.Option<ConfluxCodec._IJson>.create_None();
      }
    }

    public static Dafny.ISequence<Rune> Sha256Digest(Dafny.ISequence<byte> bytes) {
      var digest = SHA256.HashData(bytes.CloneAsArray());
      return Text("sha256:" + Convert.ToHexString(digest).ToLowerInvariant());
    }
  }
}
