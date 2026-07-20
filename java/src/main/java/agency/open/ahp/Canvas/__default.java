// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.Canvas;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static dafny.Tuple2<CanvasState, agency.open.ahp.AhpSkeleton.ReduceOutcome> ApplyToCanvas(CanvasState s, CanvasAction a, java.math.BigInteger now)
  {
    CanvasState _pat_let_tv0 = s;
    CanvasState _pat_let_tv1 = s;
    CanvasState _pat_let_tv2 = s;
    CanvasState _pat_let_tv3 = s;
    CanvasAction _source0 = a;
    if (_source0.is_Updated()) {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _0___mcc_h0 = ((agency.open.ahp.Canvas.CanvasAction_Updated)_source0)._title;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _1___mcc_h1 = ((agency.open.ahp.Canvas.CanvasAction_Updated)_source0)._activity;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _2___mcc_h2 = ((agency.open.ahp.Canvas.CanvasAction_Updated)_source0)._contentUri;
      agency.open.ahp.AhpSkeleton.Option<CanvasAvailability> _3___mcc_h3 = ((agency.open.ahp.Canvas.CanvasAction_Updated)_source0)._availability;
      agency.open.ahp.AhpSkeleton.Option<CanvasAvailability> _4_av = _3___mcc_h3;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _5_uri = _2___mcc_h2;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _6_ac = _1___mcc_h1;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _7_t = _0___mcc_h0;
      return dafny.Tuple2.<CanvasState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasState, CanvasState>Let(s, boxed16 -> {
  CanvasState _pat_let8_0 = ((CanvasState)(java.lang.Object)(boxed16));
  return ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasState, CanvasState>Let(_pat_let8_0, boxed17 -> {
    CanvasState _8_dt__update__tmp_h0 = ((CanvasState)(java.lang.Object)(boxed17));
    return ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasAvailability, CanvasState>Let((((_4_av).is_Some()) ? ((_4_av).dtor_value()) : ((_pat_let_tv0).dtor_availability())), boxed18 -> {
      CanvasAvailability _pat_let9_0 = ((CanvasAvailability)(java.lang.Object)(boxed18));
      return ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasAvailability, CanvasState>Let(_pat_let9_0, boxed19 -> {
        CanvasAvailability _9_dt__update_havailability_h0 = ((CanvasAvailability)(java.lang.Object)(boxed19));
        return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let((((_5_uri).is_Some()) ? (_5_uri) : ((_pat_let_tv1).dtor_contentUri())), boxed20 -> {
          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let10_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed20));
          return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(_pat_let10_0, boxed21 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _10_dt__update_hcontentUri_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed21));
            return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let((((_6_ac).is_Some()) ? (_6_ac) : ((_pat_let_tv2).dtor_activity())), boxed22 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let11_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed22));
              return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(_pat_let11_0, boxed23 -> {
                agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _11_dt__update_hactivity_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed23));
                return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let((((_7_t).is_Some()) ? (_7_t) : ((_pat_let_tv3).dtor_title())), boxed24 -> {
                  agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let12_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed24));
                  return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(_pat_let12_0, boxed25 -> {
                    agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _12_dt__update_htitle_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed25));
                    return CanvasState.create((_8_dt__update__tmp_h0).dtor_canvasId(), (_8_dt__update__tmp_h0).dtor_providerId(), (_8_dt__update__tmp_h0).dtor_input(), _12_dt__update_htitle_h0, _11_dt__update_hactivity_h0, _10_dt__update_hcontentUri_h0, _9_dt__update_havailability_h0);
                  }
                  )));
                }
                )));
              }
              )));
            }
            )));
          }
          )));
        }
        )));
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_CloseRequested()) {
      return dafny.Tuple2.<CanvasState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
    } else {
      agency.open.ahp.ConfluxCodec.Json _13___mcc_h4 = ((agency.open.ahp.Canvas.CanvasAction_CanvasUnknown)_source0)._raw;
      return dafny.Tuple2.<CanvasState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
    }
  }
  public static CanvasState apply1(CanvasState s, CanvasAction a)
  {
    return (__default.ApplyToCanvas(s, a, java.math.BigInteger.valueOf(9999L))).dtor__0();
  }
  public static CanvasState fold(CanvasState s, dafny.DafnySequence<? extends CanvasAction> actions)
  {
    return agency.open.ahp.ConfluxContract.__default.<CanvasState, CanvasAction>Fold(CanvasState._typeDescriptor(), CanvasAction._typeDescriptor(), __default::apply1, s, actions);
  }
  public static CanvasState C0() {
    return CanvasState.create(dafny.DafnySequence.asUnicodeString("markdown"), dafny.DafnySequence.asUnicodeString("opaque-provider"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnyMap.<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), CanvasAvailability.create_Ready());
  }
  public static boolean SameInstance(CanvasInstance a, CanvasInstance b)
  {
    return ((a).dtor_channel()).equals((b).dtor_channel());
  }
  public static ResourceReadParams ContentResourceRead(dafny.DafnySequence<? extends dafny.CodePoint> contentUri) {
    return ResourceReadParams.create(dafny.DafnySequence.asUnicodeString("ahp-root://"), contentUri);
  }
  public static CanvasAdmission ValidateCanvasAdmission(boolean supportNegotiated, boolean handlerInstalled, boolean authorized, boolean channelInUse)
  {
    if (!(supportNegotiated)) {
      return CanvasAdmission.create_CanvasUnsupported();
    } else if (!(handlerInstalled)) {
      return CanvasAdmission.create_CanvasHandlerUnavailable();
    } else if (!(authorized)) {
      return CanvasAdmission.create_CanvasUnauthorized();
    } else if (channelInUse) {
      return CanvasAdmission.create_CanvasUriCollision();
    } else {
      return CanvasAdmission.create_CanvasSupported();
    }
  }
  public static CanvasError CanvasProviderFailure(dafny.DafnySequence<? extends dafny.CodePoint> code, dafny.DafnySequence<? extends dafny.CodePoint> message)
  {
    return CanvasError.create_ProviderError(__default.CANVAS__PROVIDER__ERROR(), CanvasProviderErrorData.create(code, message));
  }
  public static CanvasState DegradeAvailability(CanvasState s) {
    CanvasState _0_dt__update__tmp_h0 = s;
    CanvasAvailability _1_dt__update_havailability_h0 = CanvasAvailability.create_Stale();
    return CanvasState.create((_0_dt__update__tmp_h0).dtor_canvasId(), (_0_dt__update__tmp_h0).dtor_providerId(), (_0_dt__update__tmp_h0).dtor_input(), (_0_dt__update__tmp_h0).dtor_title(), (_0_dt__update__tmp_h0).dtor_activity(), (_0_dt__update__tmp_h0).dtor_contentUri(), _1_dt__update_havailability_h0);
  }
  public static agency.open.ahp.AhpSkeleton.Option<CanvasState> ReconnectCanvas(CanvasState snapshot, boolean supportNegotiated)
  {
    if (supportNegotiated) {
      return agency.open.ahp.AhpSkeleton.Option.<CanvasState>create_Some(CanvasState._typeDescriptor(), snapshot);
    } else {
      return agency.open.ahp.AhpSkeleton.Option.<CanvasState>create_None(CanvasState._typeDescriptor());
    }
  }
  public static dafny.Tuple2 RunCorpus()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    java.math.BigInteger total = java.math.BigInteger.ZERO;
    if(true) {
      total = java.math.BigInteger.valueOf(5L);
      pass = java.math.BigInteger.ZERO;
      CanvasState _0_s;
      _0_s = CanvasState.create(dafny.DafnySequence.asUnicodeString("markdown"), dafny.DafnySequence.asUnicodeString("acme.editors"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnyMap.<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Draft")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("idle")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("ahp-session:/2f9c/content/canvas-1")), CanvasAvailability.create_Ready());
      CanvasState _1_all;
      _1_all = __default.apply1(_0_s, CanvasAction.create_Updated(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Published")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("saved")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("https://example.com/docs/published.html")), agency.open.ahp.AhpSkeleton.Option.<CanvasAvailability>create_Some(CanvasAvailability._typeDescriptor(), CanvasAvailability.create_Stale())));
      if (java.util.Objects.equals(_1_all, ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasState, CanvasState>Let(_0_s, boxed26 -> {
        CanvasState _pat_let13_0 = ((CanvasState)(java.lang.Object)(boxed26));
        return ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasState, CanvasState>Let(_pat_let13_0, boxed27 -> {
          CanvasState _2_dt__update__tmp_h0 = ((CanvasState)(java.lang.Object)(boxed27));
          return ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasAvailability, CanvasState>Let(CanvasAvailability.create_Stale(), boxed28 -> {
            CanvasAvailability _pat_let14_0 = ((CanvasAvailability)(java.lang.Object)(boxed28));
            return ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasAvailability, CanvasState>Let(_pat_let14_0, boxed29 -> {
              CanvasAvailability _3_dt__update_havailability_h0 = ((CanvasAvailability)(java.lang.Object)(boxed29));
              return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("https://example.com/docs/published.html")), boxed30 -> {
                agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let15_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed30));
                return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(_pat_let15_0, boxed31 -> {
                  agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _4_dt__update_hcontentUri_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed31));
                  return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("saved")), boxed32 -> {
                    agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let16_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed32));
                    return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(_pat_let16_0, boxed33 -> {
                      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _5_dt__update_hactivity_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed33));
                      return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Published")), boxed34 -> {
                        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let17_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed34));
                        return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(_pat_let17_0, boxed35 -> {
                          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _6_dt__update_htitle_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed35));
                          return CanvasState.create((_2_dt__update__tmp_h0).dtor_canvasId(), (_2_dt__update__tmp_h0).dtor_providerId(), (_2_dt__update__tmp_h0).dtor_input(), _6_dt__update_htitle_h0, _5_dt__update_hactivity_h0, _4_dt__update_hcontentUri_h0, _3_dt__update_havailability_h0);
                        }
                        )));
                      }
                      )));
                    }
                    )));
                  }
                  )));
                }
                )));
              }
              )));
            }
            )));
          }
          )));
        }
        )));
      }
      ))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      CanvasState _7_left;
      _7_left = __default.apply1(_0_s, CanvasAction.create_Updated(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Renamed")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("https://example.com/docs/renamed.html")), agency.open.ahp.AhpSkeleton.Option.<CanvasAvailability>create_None(CanvasAvailability._typeDescriptor())));
      if (java.util.Objects.equals(_7_left, ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasState, CanvasState>Let(_0_s, boxed36 -> {
        CanvasState _pat_let18_0 = ((CanvasState)(java.lang.Object)(boxed36));
        return ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasState, CanvasState>Let(_pat_let18_0, boxed37 -> {
          CanvasState _8_dt__update__tmp_h1 = ((CanvasState)(java.lang.Object)(boxed37));
          return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("https://example.com/docs/renamed.html")), boxed38 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let19_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed38));
            return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(_pat_let19_0, boxed39 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _9_dt__update_hcontentUri_h1 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed39));
              return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Renamed")), boxed40 -> {
                agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let20_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed40));
                return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(_pat_let20_0, boxed41 -> {
                  agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _10_dt__update_htitle_h1 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed41));
                  return CanvasState.create((_8_dt__update__tmp_h1).dtor_canvasId(), (_8_dt__update__tmp_h1).dtor_providerId(), (_8_dt__update__tmp_h1).dtor_input(), _10_dt__update_htitle_h1, (_8_dt__update__tmp_h1).dtor_activity(), _9_dt__update_hcontentUri_h1, (_8_dt__update__tmp_h1).dtor_availability());
                }
                )));
              }
              )));
            }
            )));
          }
          )));
        }
        )));
      }
      ))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      CanvasState _11_right;
      _11_right = __default.apply1(_0_s, CanvasAction.create_Updated(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("error")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<CanvasAvailability>create_Some(CanvasAvailability._typeDescriptor(), CanvasAvailability.create_Stale())));
      if (java.util.Objects.equals(_11_right, ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasState, CanvasState>Let(_0_s, boxed42 -> {
        CanvasState _pat_let21_0 = ((CanvasState)(java.lang.Object)(boxed42));
        return ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasState, CanvasState>Let(_pat_let21_0, boxed43 -> {
          CanvasState _12_dt__update__tmp_h2 = ((CanvasState)(java.lang.Object)(boxed43));
          return ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasAvailability, CanvasState>Let(CanvasAvailability.create_Stale(), boxed44 -> {
            CanvasAvailability _pat_let22_0 = ((CanvasAvailability)(java.lang.Object)(boxed44));
            return ((CanvasState)(java.lang.Object)(dafny.Helpers.<CanvasAvailability, CanvasState>Let(_pat_let22_0, boxed45 -> {
              CanvasAvailability _13_dt__update_havailability_h1 = ((CanvasAvailability)(java.lang.Object)(boxed45));
              return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("error")), boxed46 -> {
                agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let23_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed46));
                return ((CanvasState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, CanvasState>Let(_pat_let23_0, boxed47 -> {
                  agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _14_dt__update_hactivity_h1 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed47));
                  return CanvasState.create((_12_dt__update__tmp_h2).dtor_canvasId(), (_12_dt__update__tmp_h2).dtor_providerId(), (_12_dt__update__tmp_h2).dtor_input(), (_12_dt__update__tmp_h2).dtor_title(), _14_dt__update_hactivity_h1, (_12_dt__update__tmp_h2).dtor_contentUri(), _13_dt__update_havailability_h1);
                }
                )));
              }
              )));
            }
            )));
          }
          )));
        }
        )));
      }
      ))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_0_s, CanvasAction.create_CloseRequested()), _0_s)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_0_s, CanvasAction.create_CanvasUnknown(agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("type"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("canvas/nonExistentAction"))) })))), _0_s)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
    }
    return new dafny.Tuple2<>(pass, total);
  }
  public static java.math.BigInteger CANVAS__PROVIDER__ERROR()
  {
    return java.math.BigInteger.valueOf(-32012L);
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.Canvas._default";
  }
}
