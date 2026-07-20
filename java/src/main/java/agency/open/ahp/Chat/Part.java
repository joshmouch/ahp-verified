// Class Part
// Dafny class Part compiled into Java
package agency.open.ahp.Chat;

@SuppressWarnings({"unchecked", "deprecation"})
public abstract class Part {
  public Part() {
  }
  private static final dafny.TypeDescriptor<Part> _TYPE = dafny.TypeDescriptor.<Part>referenceWithInitializer(Part.class, () -> Part.Default());
  public static dafny.TypeDescriptor<Part> _typeDescriptor() {
    return (dafny.TypeDescriptor<Part>) (dafny.TypeDescriptor<?>) _TYPE;
  }

  private static final Part theDefault = Part.create_PMarkdown(dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint> empty(dafny.TypeDescriptor.UNICODE_CHAR));
  public static Part Default() {
    return theDefault;
  }
  public static Part create_PMarkdown(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> content) {
    return new Part_PMarkdown(id, content);
  }
  public static Part create_PReasoning(dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> content) {
    return new Part_PReasoning(id, content);
  }
  public static Part create_PToolCall(ToolCall tc) {
    return new Part_PToolCall(tc);
  }
  public static Part create_PInputRequest(InputReq req, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> response) {
    return new Part_PInputRequest(req, response);
  }
  public boolean is_PMarkdown() { return this instanceof Part_PMarkdown; }
  public boolean is_PReasoning() { return this instanceof Part_PReasoning; }
  public boolean is_PToolCall() { return this instanceof Part_PToolCall; }
  public boolean is_PInputRequest() { return this instanceof Part_PInputRequest; }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_id() {
    Part d = this;
    if (d instanceof Part_PMarkdown) { return ((Part_PMarkdown)d)._id; }
    return ((Part_PReasoning)d)._id;
  }
  public dafny.DafnySequence<? extends dafny.CodePoint> dtor_content() {
    Part d = this;
    if (d instanceof Part_PMarkdown) { return ((Part_PMarkdown)d)._content; }
    return ((Part_PReasoning)d)._content;
  }
  public ToolCall dtor_tc() {
    Part d = this;
    return ((Part_PToolCall)d)._tc;
  }
  public InputReq dtor_req() {
    Part d = this;
    return ((Part_PInputRequest)d)._req;
  }
  public agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> dtor_response() {
    Part d = this;
    return ((Part_PInputRequest)d)._response;
  }
}
