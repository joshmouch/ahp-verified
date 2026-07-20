// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.ClientMain;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static void Main(dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> __noArgsParameter)
  {
    java.math.BigInteger _0_rootPass = java.math.BigInteger.ZERO;
    java.math.BigInteger _1_rootTotal = java.math.BigInteger.ZERO;
    java.math.BigInteger _out0 = java.math.BigInteger.ZERO;
    java.math.BigInteger _out1 = java.math.BigInteger.ZERO;
    dafny.Tuple2<java.math.BigInteger, java.math.BigInteger> _outcollector0 = agency.open.ahp.AhpSkeleton.__default.RunCorpus();
    _out0 = (java.math.BigInteger) (Object) _outcollector0.dtor__0();
    _out1 = (java.math.BigInteger) (Object) _outcollector0.dtor__1();
    _0_rootPass = _out0;
    _1_rootTotal = _out1;
    System.out.print((dafny.DafnySequence.asUnicodeString("ROOT CORPUS:          ")).verbatimString());
    System.out.print(java.lang.String.valueOf(_0_rootPass));
    System.out.print((dafny.DafnySequence.asUnicodeString("/")).verbatimString());
    System.out.print(java.lang.String.valueOf(_1_rootTotal));
    System.out.print((dafny.DafnySequence.asUnicodeString(" green against extracted code\n")).verbatimString());
    java.math.BigInteger _2_rwPass = java.math.BigInteger.ZERO;
    java.math.BigInteger _3_rwTotal = java.math.BigInteger.ZERO;
    java.math.BigInteger _out2 = java.math.BigInteger.ZERO;
    java.math.BigInteger _out3 = java.math.BigInteger.ZERO;
    dafny.Tuple2<java.math.BigInteger, java.math.BigInteger> _outcollector1 = agency.open.ahp.ResourceWatch.__default.RunCorpus();
    _out2 = (java.math.BigInteger) (Object) _outcollector1.dtor__0();
    _out3 = (java.math.BigInteger) (Object) _outcollector1.dtor__1();
    _2_rwPass = _out2;
    _3_rwTotal = _out3;
    System.out.print((dafny.DafnySequence.asUnicodeString("RESOURCEWATCH CORPUS: ")).verbatimString());
    System.out.print(java.lang.String.valueOf(_2_rwPass));
    System.out.print((dafny.DafnySequence.asUnicodeString("/")).verbatimString());
    System.out.print(java.lang.String.valueOf(_3_rwTotal));
    System.out.print((dafny.DafnySequence.asUnicodeString(" green against extracted code\n")).verbatimString());
    java.math.BigInteger _4_cvPass = java.math.BigInteger.ZERO;
    java.math.BigInteger _5_cvTotal = java.math.BigInteger.ZERO;
    java.math.BigInteger _out4 = java.math.BigInteger.ZERO;
    java.math.BigInteger _out5 = java.math.BigInteger.ZERO;
    dafny.Tuple2<java.math.BigInteger, java.math.BigInteger> _outcollector2 = agency.open.ahp.Canvas.__default.RunCorpus();
    _out4 = (java.math.BigInteger) (Object) _outcollector2.dtor__0();
    _out5 = (java.math.BigInteger) (Object) _outcollector2.dtor__1();
    _4_cvPass = _out4;
    _5_cvTotal = _out5;
    System.out.print((dafny.DafnySequence.asUnicodeString("CANVAS CORPUS:        ")).verbatimString());
    System.out.print(java.lang.String.valueOf(_4_cvPass));
    System.out.print((dafny.DafnySequence.asUnicodeString("/")).verbatimString());
    System.out.print(java.lang.String.valueOf(_5_cvTotal));
    System.out.print((dafny.DafnySequence.asUnicodeString(" green against extracted code\n")).verbatimString());
    java.math.BigInteger _6_csPass = java.math.BigInteger.ZERO;
    java.math.BigInteger _7_csTotal = java.math.BigInteger.ZERO;
    java.math.BigInteger _out6 = java.math.BigInteger.ZERO;
    java.math.BigInteger _out7 = java.math.BigInteger.ZERO;
    dafny.Tuple2<java.math.BigInteger, java.math.BigInteger> _outcollector3 = agency.open.ahp.Changeset.__default.RunCorpus();
    _out6 = (java.math.BigInteger) (Object) _outcollector3.dtor__0();
    _out7 = (java.math.BigInteger) (Object) _outcollector3.dtor__1();
    _6_csPass = _out6;
    _7_csTotal = _out7;
    System.out.print((dafny.DafnySequence.asUnicodeString("CHANGESET CORPUS:     ")).verbatimString());
    System.out.print(java.lang.String.valueOf(_6_csPass));
    System.out.print((dafny.DafnySequence.asUnicodeString("/")).verbatimString());
    System.out.print(java.lang.String.valueOf(_7_csTotal));
    System.out.print((dafny.DafnySequence.asUnicodeString(" green against extracted code\n")).verbatimString());
    java.math.BigInteger _8_anPass = java.math.BigInteger.ZERO;
    java.math.BigInteger _9_anTotal = java.math.BigInteger.ZERO;
    java.math.BigInteger _out8 = java.math.BigInteger.ZERO;
    java.math.BigInteger _out9 = java.math.BigInteger.ZERO;
    dafny.Tuple2<java.math.BigInteger, java.math.BigInteger> _outcollector4 = agency.open.ahp.Annotations.__default.RunCorpus();
    _out8 = (java.math.BigInteger) (Object) _outcollector4.dtor__0();
    _out9 = (java.math.BigInteger) (Object) _outcollector4.dtor__1();
    _8_anPass = _out8;
    _9_anTotal = _out9;
    System.out.print((dafny.DafnySequence.asUnicodeString("ANNOTATIONS CORPUS:   ")).verbatimString());
    System.out.print(java.lang.String.valueOf(_8_anPass));
    System.out.print((dafny.DafnySequence.asUnicodeString("/")).verbatimString());
    System.out.print(java.lang.String.valueOf(_9_anTotal));
    System.out.print((dafny.DafnySequence.asUnicodeString(" green against extracted code\n")).verbatimString());
    java.math.BigInteger _10_tPass = java.math.BigInteger.ZERO;
    java.math.BigInteger _11_tTotal = java.math.BigInteger.ZERO;
    java.math.BigInteger _out10 = java.math.BigInteger.ZERO;
    java.math.BigInteger _out11 = java.math.BigInteger.ZERO;
    dafny.Tuple2<java.math.BigInteger, java.math.BigInteger> _outcollector5 = agency.open.ahp.Terminal.__default.RunCorpus();
    _out10 = (java.math.BigInteger) (Object) _outcollector5.dtor__0();
    _out11 = (java.math.BigInteger) (Object) _outcollector5.dtor__1();
    _10_tPass = _out10;
    _11_tTotal = _out11;
    System.out.print((dafny.DafnySequence.asUnicodeString("TERMINAL CORPUS:      ")).verbatimString());
    System.out.print(java.lang.String.valueOf(_10_tPass));
    System.out.print((dafny.DafnySequence.asUnicodeString("/")).verbatimString());
    System.out.print(java.lang.String.valueOf(_11_tTotal));
    System.out.print((dafny.DafnySequence.asUnicodeString(" green against extracted code\n")).verbatimString());
    java.math.BigInteger _12_sPass = java.math.BigInteger.ZERO;
    java.math.BigInteger _13_sModeled = java.math.BigInteger.ZERO;
    java.math.BigInteger _out12 = java.math.BigInteger.ZERO;
    java.math.BigInteger _out13 = java.math.BigInteger.ZERO;
    dafny.Tuple2<java.math.BigInteger, java.math.BigInteger> _outcollector6 = agency.open.ahp.Session.__default.RunCorpus();
    _out12 = (java.math.BigInteger) (Object) _outcollector6.dtor__0();
    _out13 = (java.math.BigInteger) (Object) _outcollector6.dtor__1();
    _12_sPass = _out12;
    _13_sModeled = _out13;
    System.out.print((dafny.DafnySequence.asUnicodeString("SESSION CORPUS:       ")).verbatimString());
    System.out.print(java.lang.String.valueOf(_12_sPass));
    System.out.print((dafny.DafnySequence.asUnicodeString("/")).verbatimString());
    System.out.print(java.lang.String.valueOf(_13_sModeled));
    System.out.print((dafny.DafnySequence.asUnicodeString(" modeled green (of 61 total; all ~25 action TYPES now modeled)\n")).verbatimString());
    java.math.BigInteger _14_chPass = java.math.BigInteger.ZERO;
    java.math.BigInteger _15_chModeled = java.math.BigInteger.ZERO;
    java.math.BigInteger _out14 = java.math.BigInteger.ZERO;
    java.math.BigInteger _out15 = java.math.BigInteger.ZERO;
    dafny.Tuple2<java.math.BigInteger, java.math.BigInteger> _outcollector7 = agency.open.ahp.Chat.__default.RunCorpus();
    _out14 = (java.math.BigInteger) (Object) _outcollector7.dtor__0();
    _out15 = (java.math.BigInteger) (Object) _outcollector7.dtor__1();
    _14_chPass = _out14;
    _15_chModeled = _out15;
    System.out.print((dafny.DafnySequence.asUnicodeString("CHAT CORPUS:          ")).verbatimString());
    System.out.print(java.lang.String.valueOf(_14_chPass));
    System.out.print((dafny.DafnySequence.asUnicodeString("/")).verbatimString());
    System.out.print(java.lang.String.valueOf(_15_chModeled));
    System.out.print((dafny.DafnySequence.asUnicodeString(" modeled green (of 97 total; full tool-call state machine + turn lifecycle modeled)\n")).verbatimString());
    java.math.BigInteger _16_pass = java.math.BigInteger.ZERO;
    _16_pass = (((((((_0_rootPass).add(_2_rwPass)).add(_4_cvPass)).add(_6_csPass)).add(_8_anPass)).add(_10_tPass)).add(_12_sPass)).add(_14_chPass);
    java.math.BigInteger _17_total = java.math.BigInteger.ZERO;
    _17_total = (((((((_1_rootTotal).add(_3_rwTotal)).add(_5_cvTotal)).add(_7_csTotal)).add(_9_anTotal)).add(_11_tTotal)).add(_13_sModeled)).add(_15_chModeled);
    System.out.print((dafny.DafnySequence.asUnicodeString("TOTAL: ")).verbatimString());
    System.out.print(java.lang.String.valueOf(_16_pass));
    System.out.print((dafny.DafnySequence.asUnicodeString("/")).verbatimString());
    System.out.print(java.lang.String.valueOf(_17_total));
    System.out.print((dafny.DafnySequence.asUnicodeString(" corpus fixtures green (5 full AHP channels + session/chat partial)\n")).verbatimString());
    if (!(java.util.Objects.equals(_16_pass, _17_total))) {
      throw new dafny.DafnyHaltException("spec/client_main.dfy(41,4): " + (dafny.DafnySequence.asUnicodeString("expectation violation")).verbatimString());
    }
    System.out.print((dafny.DafnySequence.asUnicodeString("MULTI-CHANNEL CLIENT OK — extract(cs,js) + corpus all green\n")).verbatimString());
  }
  public static void __Main(dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> args) {
    __default.Main(args);
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.ClientMain._default";
  }
}
