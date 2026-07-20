// Class Part
// Dafny class Part compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class Part {
  public Part() {
  }
  private static final dafny.TypeDescriptor<Part> _TYPE = dafny.TypeDescriptor.<Part>referenceWithInitializer(Part.class, () -> Part.Default());
  public static dafny.TypeDescriptor<Part> _typeDescriptor() {
    return (dafny.TypeDescriptor<Part>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final Part theDefault = Part.create_Unclassified(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR));
  public static Part Default() {
    return theDefault;
  }
  public static Part create_Unclassified(dafny.DafnySequence<? extends dafny.CodePoint> value) {
    return new Part_Unclassified(value);
  }
  public static Part create_Command(dafny.DafnySequence<? extends dafny.CodePoint> commandId, dafny.DafnySequence<? extends dafny.CodePoint> commandLine, dafny.DafnySequence<? extends dafny.CodePoint> output, java.math.BigInteger timestamp, boolean isComplete, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> exitCode, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> durationMs) {
    return new Part_Command(commandId, commandLine, output, timestamp, isComplete, exitCode, durationMs);
  }
  public boolean is_Unclassified() { return this instanceof Part_Unclassified; }
  public boolean is_Command() { return this instanceof Part_Command; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_value() {
    Part d = this;
    return ((Part_Unclassified)d)._value;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_commandId() {
    Part d = this;
    return ((Part_Command)d)._commandId;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_commandLine() {
    Part d = this;
    return ((Part_Command)d)._commandLine;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_output() {
    Part d = this;
    return ((Part_Command)d)._output;
  }
  public java.math.BigInteger dtor_timestamp() {
    Part d = this;
    return ((Part_Command)d)._timestamp;
  }
  public boolean dtor_isComplete() {
    Part d = this;
    return ((Part_Command)d)._isComplete;
  }
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> dtor_exitCode() {
    Part d = this;
    return ((Part_Command)d)._exitCode;
  }
  public agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> dtor_durationMs() {
    Part d = this;
    return ((Part_Command)d)._durationMs;
  }
}
