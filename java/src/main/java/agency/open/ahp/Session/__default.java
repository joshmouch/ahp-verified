// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.Session;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static int setBit(int s, int b)
  {
    return (int) ((s) | (b));
  }
  public static int clearBit(int s, int b)
  {
    return (int) ((s) & ((int) ((~(b)))));
  }
  public static java.util.function.Function<Cust, Cust> toggleEnabled(boolean en) {
    return ((java.util.function.Function<Boolean, java.util.function.Function<Cust, Cust>>)(__0_en0) -> {boolean _0_en = ((boolean)(java.lang.Object)(__0_en0));
    return ((java.util.function.Function<Cust, Cust>)(_1_e_boxed0) -> {
      Cust _1_e = ((Cust)(java.lang.Object)(_1_e_boxed0));
      return ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_1_e, boxed200 -> {
        Cust _pat_let100_0 = ((Cust)(java.lang.Object)(boxed200));
        return ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_pat_let100_0, boxed201 -> {
          Cust _2_dt__update__tmp_h0 = ((Cust)(java.lang.Object)(boxed201));
          return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, Cust>Let(agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, _0_en), boxed202 -> {
            agency.open.ahp.AhpSkeleton.Option<Boolean> _pat_let101_0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed202));
            return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, Cust>Let(_pat_let101_0, boxed203 -> {
              agency.open.ahp.AhpSkeleton.Option<Boolean> _3_dt__update_henabled_h0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed203));
              return Cust.create((_2_dt__update__tmp_h0).dtor_id(), (_2_dt__update__tmp_h0).dtor_kind(), (_2_dt__update__tmp_h0).dtor_uri(), (_2_dt__update__tmp_h0).dtor_name(), _3_dt__update_henabled_h0, (_2_dt__update__tmp_h0).dtor_state(), (_2_dt__update__tmp_h0).dtor_channel());
            }
            )));
          }
          )));
        }
        )));
      }
      )));
    });}).apply(en);
  }
  public static java.util.function.Function<Chat, Chat> chatPatch(agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> st, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> ac, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> md)
  {
    return ((dafny.Function3<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>, java.util.function.Function<Chat, Chat>>)(_0_md, _1_ac, _2_st) -> {return ((java.util.function.Function<Chat, Chat>)(_3_c_boxed0) -> {
      Chat _3_c = ((Chat)(java.lang.Object)(_3_c_boxed0));
      return ((Chat)(java.lang.Object)(dafny.Helpers.<Chat, Chat>Let(_3_c, boxed204 -> {
        Chat _pat_let102_0 = ((Chat)(java.lang.Object)(boxed204));
        return ((Chat)(java.lang.Object)(dafny.Helpers.<Chat, Chat>Let(_pat_let102_0, boxed205 -> {
          Chat _4_dt__update__tmp_h0 = ((Chat)(java.lang.Object)(boxed205));
          return ((Chat)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, Chat>Let((((_0_md).is_Some()) ? ((_0_md).dtor_value()) : ((_3_c).dtor_modifiedAt())), boxed206 -> {
            dafny.DafnySequence<? extends dafny.CodePoint> _pat_let103_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed206));
            return ((Chat)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, Chat>Let(_pat_let103_0, boxed207 -> {
              dafny.DafnySequence<? extends dafny.CodePoint> _5_dt__update_hmodifiedAt_h0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed207));
              return ((Chat)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, Chat>Let((((_1_ac).is_Some()) ? (_1_ac) : ((_3_c).dtor_activity())), boxed208 -> {
                agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let104_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed208));
                return ((Chat)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, Chat>Let(_pat_let104_0, boxed209 -> {
                  agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _6_dt__update_hactivity_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed209));
                  return ((Chat)(java.lang.Object)(dafny.Helpers.<java.math.BigInteger, Chat>Let((((_2_st).is_Some()) ? ((_2_st).dtor_value()) : ((_3_c).dtor_status())), boxed210 -> {
                    java.math.BigInteger _pat_let105_0 = ((java.math.BigInteger)(java.lang.Object)(boxed210));
                    return ((Chat)(java.lang.Object)(dafny.Helpers.<java.math.BigInteger, Chat>Let(_pat_let105_0, boxed211 -> {
                      java.math.BigInteger _7_dt__update_hstatus_h0 = ((java.math.BigInteger)(java.lang.Object)(boxed211));
                      return Chat.create((_4_dt__update__tmp_h0).dtor_resource(), (_4_dt__update__tmp_h0).dtor_title(), _7_dt__update_hstatus_h0, _6_dt__update_hactivity_h0, _5_dt__update_hmodifiedAt_h0, (_4_dt__update__tmp_h0).dtor_origin());
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
    });}).apply(md, ac, st);
  }
  public static java.util.function.Function<Cust, Cust> mcpPatch(agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> st, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> ch)
  {
    return ((dafny.Function2<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, java.util.function.Function<Cust, Cust>>)(_0_ch, _1_st) -> {return ((java.util.function.Function<Cust, Cust>)(_2_e_boxed0) -> {
      Cust _2_e = ((Cust)(java.lang.Object)(_2_e_boxed0));
      return ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_2_e, boxed212 -> {
        Cust _pat_let106_0 = ((Cust)(java.lang.Object)(boxed212));
        return ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_pat_let106_0, boxed213 -> {
          Cust _3_dt__update__tmp_h0 = ((Cust)(java.lang.Object)(boxed213));
          return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(_0_ch, boxed214 -> {
            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let107_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed214));
            return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(_pat_let107_0, boxed215 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _4_dt__update_hchannel_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed215));
              return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(_1_st, boxed216 -> {
                agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let108_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed216));
                return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(_pat_let108_0, boxed217 -> {
                  agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _5_dt__update_hstate_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed217));
                  return Cust.create((_3_dt__update__tmp_h0).dtor_id(), (_3_dt__update__tmp_h0).dtor_kind(), (_3_dt__update__tmp_h0).dtor_uri(), (_3_dt__update__tmp_h0).dtor_name(), (_3_dt__update__tmp_h0).dtor_enabled(), _5_dt__update_hstate_h0, _4_dt__update_hchannel_h0);
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
    });}).apply(ch, st);
  }
  public static dafny.DafnySequence<? extends Cust> upsertCust(dafny.DafnySequence<? extends Cust> cs, dafny.DafnySequence<? extends dafny.CodePoint> id, boolean en)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Cust>SeqUpdateById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Cust._typeDescriptor(), __default.custKey(), cs, id, __default.toggleEnabled(en));
  }
  public static dafny.DafnySequence<? extends Cust> removeCust(dafny.DafnySequence<? extends Cust> cs, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Cust>SeqRemoveById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Cust._typeDescriptor(), __default.custKey(), cs, id);
  }
  public static dafny.DafnySequence<? extends Chat> removeChat(dafny.DafnySequence<? extends Chat> chs, dafny.DafnySequence<? extends dafny.CodePoint> r)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Chat>SeqRemoveById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Chat._typeDescriptor(), __default.chatKey(), chs, r);
  }
  public static dafny.DafnySequence<? extends Chat> upsertChat(dafny.DafnySequence<? extends Chat> chs, Chat c)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Chat>SeqUpsertById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Chat._typeDescriptor(), __default.chatKey(), chs, c);
  }
  public static dafny.DafnySequence<? extends Chat> updChat(dafny.DafnySequence<? extends Chat> chs, dafny.DafnySequence<? extends dafny.CodePoint> r, agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> st, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> ac, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> md)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Chat>SeqUpdateById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Chat._typeDescriptor(), __default.chatKey(), chs, r, __default.chatPatch(st, ac, md));
  }
  public static dafny.DafnySequence<? extends Client> upsertClient(dafny.DafnySequence<? extends Client> cl, Client c)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Client>SeqUpsertById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Client._typeDescriptor(), __default.clientKey(), cl, c);
  }
  public static dafny.DafnySequence<? extends Client> removeClient(dafny.DafnySequence<? extends Client> cl, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Client>SeqRemoveById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Client._typeDescriptor(), __default.clientKey(), cl, id);
  }
  public static dafny.DafnySequence<? extends InputReq> upsertInput(dafny.DafnySequence<? extends InputReq> ins, InputReq r)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, InputReq>SeqUpsertById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), InputReq._typeDescriptor(), __default.inputKey(), ins, r);
  }
  public static dafny.DafnySequence<? extends InputReq> removeInput(dafny.DafnySequence<? extends InputReq> ins, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, InputReq>SeqRemoveById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), InputReq._typeDescriptor(), __default.inputKey(), ins, id);
  }
  public static dafny.DafnySequence<? extends Cust> upsertCustFull(dafny.DafnySequence<? extends Cust> cs, Cust c)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Cust>SeqUpsertById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Cust._typeDescriptor(), __default.custKey(), cs, c);
  }
  public static dafny.DafnySequence<? extends Cust> setMcp(dafny.DafnySequence<? extends Cust> cs, dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> st, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> ch)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Cust>SeqUpdateByIdWhere(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Cust._typeDescriptor(), __default.custKey(), __default.mcpPred(), cs, id, __default.mcpPatch(st, ch));
  }
  public static boolean isMcp(dafny.DafnySequence<? extends Cust> cs, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return ((dafny.Function2<dafny.DafnySequence<? extends Cust>, dafny.DafnySequence<? extends dafny.CodePoint>, Boolean>)(_0_cs, _1_id) -> {return dafny.Helpers.Quantifier(dafny.Helpers.IntegerRange(java.math.BigInteger.ZERO, java.math.BigInteger.valueOf((_0_cs).length())), false, ((_exists_var_0_boxed0) -> {
      java.math.BigInteger _exists_var_0 = ((java.math.BigInteger)(java.lang.Object)(_exists_var_0_boxed0));
      java.math.BigInteger _2_i = (java.math.BigInteger)_exists_var_0;
      return ((((_2_i).signum() != -1) && ((_2_i).compareTo(java.math.BigInteger.valueOf((_0_cs).length())) < 0)) && (((((Cust)(java.lang.Object)((_0_cs).select(dafny.Helpers.toInt((_2_i)))))).dtor_id()).equals(_1_id))) && (((((Cust)(java.lang.Object)((_0_cs).select(dafny.Helpers.toInt((_2_i)))))).dtor_kind()).equals(dafny.DafnySequence.asUnicodeString("mcpServer")));
    }));}).apply(cs, id);
  }
  public static dafny.Tuple2<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome> ApplyToSession(SessionState s, SessionAction a, java.math.BigInteger now)
  {
    SessionState _pat_let_tv0 = s;
    SessionState _pat_let_tv1 = s;
    SessionState _pat_let_tv2 = s;
    SessionState _pat_let_tv3 = s;
    SessionState _pat_let_tv4 = s;
    SessionState _pat_let_tv5 = s;
    SessionState _pat_let_tv6 = s;
    SessionState _pat_let_tv7 = s;
    SessionState _pat_let_tv8 = s;
    SessionState _pat_let_tv9 = s;
    SessionState _pat_let_tv10 = s;
    SessionState _pat_let_tv11 = s;
    SessionState _pat_let_tv12 = s;
    SessionState _pat_let_tv13 = s;
    SessionState _pat_let_tv14 = s;
    SessionState _pat_let_tv15 = s;
    SessionState _pat_let_tv16 = s;
    SessionState _pat_let_tv17 = s;
    SessionState _pat_let_tv18 = s;
    SessionState _pat_let_tv19 = s;
    SessionState _pat_let_tv20 = s;
    SessionAction _source0 = a;
    if (_source0.is_Ready()) {
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed218 -> {
  SessionState _pat_let109_0 = ((SessionState)(java.lang.Object)(boxed218));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let109_0, boxed219 -> {
    SessionState _0_dt__update__tmp_h0 = ((SessionState)(java.lang.Object)(boxed219));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, SessionState>Let(dafny.DafnySequence.asUnicodeString("ready"), boxed220 -> {
      dafny.DafnySequence<? extends dafny.CodePoint> _pat_let110_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed220));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, SessionState>Let(_pat_let110_0, boxed221 -> {
        dafny.DafnySequence<? extends dafny.CodePoint> _1_dt__update_hlifecycle_h0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed221));
        return SessionState.create((_0_dt__update__tmp_h0).dtor_provider(), (_0_dt__update__tmp_h0).dtor_title(), (_0_dt__update__tmp_h0).dtor_status(), _1_dt__update_hlifecycle_h0, (_0_dt__update__tmp_h0).dtor_activity(), (_0_dt__update__tmp_h0).dtor_config(), (_0_dt__update__tmp_h0).dtor_meta(), (_0_dt__update__tmp_h0).dtor_creationError(), (_0_dt__update__tmp_h0).dtor_serverTools(), (_0_dt__update__tmp_h0).dtor_changesets(), (_0_dt__update__tmp_h0).dtor_canvases(), (_0_dt__update__tmp_h0).dtor_openCanvases(), (_0_dt__update__tmp_h0).dtor_defaultChat(), (_0_dt__update__tmp_h0).dtor_activeClients(), (_0_dt__update__tmp_h0).dtor_chats(), (_0_dt__update__tmp_h0).dtor_customizations(), (_0_dt__update__tmp_h0).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_CreationFailed()) {
      agency.open.ahp.ConfluxCodec.Json _2___mcc_h0 = ((agency.open.ahp.Session.SessionAction_CreationFailed)_source0)._error;
      agency.open.ahp.ConfluxCodec.Json _3_e = _2___mcc_h0;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed222 -> {
  SessionState _pat_let111_0 = ((SessionState)(java.lang.Object)(boxed222));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let111_0, boxed223 -> {
    SessionState _4_dt__update__tmp_h1 = ((SessionState)(java.lang.Object)(boxed223));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(((java.util.Objects.equals(_3_e, agency.open.ahp.ConfluxCodec.Json.create_JNull())) ? (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())) : (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _3_e))), boxed224 -> {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let112_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed224));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(_pat_let112_0, boxed225 -> {
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _5_dt__update_hcreationError_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed225));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, SessionState>Let(dafny.DafnySequence.asUnicodeString("creationFailed"), boxed226 -> {
          dafny.DafnySequence<? extends dafny.CodePoint> _pat_let113_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed226));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, SessionState>Let(_pat_let113_0, boxed227 -> {
            dafny.DafnySequence<? extends dafny.CodePoint> _6_dt__update_hlifecycle_h1 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed227));
            return SessionState.create((_4_dt__update__tmp_h1).dtor_provider(), (_4_dt__update__tmp_h1).dtor_title(), (_4_dt__update__tmp_h1).dtor_status(), _6_dt__update_hlifecycle_h1, (_4_dt__update__tmp_h1).dtor_activity(), (_4_dt__update__tmp_h1).dtor_config(), (_4_dt__update__tmp_h1).dtor_meta(), _5_dt__update_hcreationError_h0, (_4_dt__update__tmp_h1).dtor_serverTools(), (_4_dt__update__tmp_h1).dtor_changesets(), (_4_dt__update__tmp_h1).dtor_canvases(), (_4_dt__update__tmp_h1).dtor_openCanvases(), (_4_dt__update__tmp_h1).dtor_defaultChat(), (_4_dt__update__tmp_h1).dtor_activeClients(), (_4_dt__update__tmp_h1).dtor_chats(), (_4_dt__update__tmp_h1).dtor_customizations(), (_4_dt__update__tmp_h1).dtor_inputNeeded());
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
    } else if (_source0.is_TitleChanged()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _7___mcc_h1 = ((agency.open.ahp.Session.SessionAction_TitleChanged)_source0)._title;
      dafny.DafnySequence<? extends dafny.CodePoint> _8_t = _7___mcc_h1;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed228 -> {
  SessionState _pat_let114_0 = ((SessionState)(java.lang.Object)(boxed228));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let114_0, boxed229 -> {
    SessionState _9_dt__update__tmp_h2 = ((SessionState)(java.lang.Object)(boxed229));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, SessionState>Let(_8_t, boxed230 -> {
      dafny.DafnySequence<? extends dafny.CodePoint> _pat_let115_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed230));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, SessionState>Let(_pat_let115_0, boxed231 -> {
        dafny.DafnySequence<? extends dafny.CodePoint> _10_dt__update_htitle_h0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed231));
        return SessionState.create((_9_dt__update__tmp_h2).dtor_provider(), _10_dt__update_htitle_h0, (_9_dt__update__tmp_h2).dtor_status(), (_9_dt__update__tmp_h2).dtor_lifecycle(), (_9_dt__update__tmp_h2).dtor_activity(), (_9_dt__update__tmp_h2).dtor_config(), (_9_dt__update__tmp_h2).dtor_meta(), (_9_dt__update__tmp_h2).dtor_creationError(), (_9_dt__update__tmp_h2).dtor_serverTools(), (_9_dt__update__tmp_h2).dtor_changesets(), (_9_dt__update__tmp_h2).dtor_canvases(), (_9_dt__update__tmp_h2).dtor_openCanvases(), (_9_dt__update__tmp_h2).dtor_defaultChat(), (_9_dt__update__tmp_h2).dtor_activeClients(), (_9_dt__update__tmp_h2).dtor_chats(), (_9_dt__update__tmp_h2).dtor_customizations(), (_9_dt__update__tmp_h2).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_ServerToolsChanged()) {
      agency.open.ahp.ConfluxCodec.Json _11___mcc_h2 = ((agency.open.ahp.Session.SessionAction_ServerToolsChanged)_source0)._tools;
      agency.open.ahp.ConfluxCodec.Json _12_t = _11___mcc_h2;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed232 -> {
  SessionState _pat_let116_0 = ((SessionState)(java.lang.Object)(boxed232));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let116_0, boxed233 -> {
    SessionState _13_dt__update__tmp_h3 = ((SessionState)(java.lang.Object)(boxed233));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(((java.util.Objects.equals(_12_t, agency.open.ahp.ConfluxCodec.Json.create_JNull())) ? (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())) : (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _12_t))), boxed234 -> {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let117_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed234));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(_pat_let117_0, boxed235 -> {
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _14_dt__update_hserverTools_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed235));
        return SessionState.create((_13_dt__update__tmp_h3).dtor_provider(), (_13_dt__update__tmp_h3).dtor_title(), (_13_dt__update__tmp_h3).dtor_status(), (_13_dt__update__tmp_h3).dtor_lifecycle(), (_13_dt__update__tmp_h3).dtor_activity(), (_13_dt__update__tmp_h3).dtor_config(), (_13_dt__update__tmp_h3).dtor_meta(), (_13_dt__update__tmp_h3).dtor_creationError(), _14_dt__update_hserverTools_h0, (_13_dt__update__tmp_h3).dtor_changesets(), (_13_dt__update__tmp_h3).dtor_canvases(), (_13_dt__update__tmp_h3).dtor_openCanvases(), (_13_dt__update__tmp_h3).dtor_defaultChat(), (_13_dt__update__tmp_h3).dtor_activeClients(), (_13_dt__update__tmp_h3).dtor_chats(), (_13_dt__update__tmp_h3).dtor_customizations(), (_13_dt__update__tmp_h3).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_MetaChanged()) {
      agency.open.ahp.ConfluxCodec.Json _15___mcc_h3 = ((agency.open.ahp.Session.SessionAction_MetaChanged)_source0)._meta;
      agency.open.ahp.ConfluxCodec.Json _16_m = _15___mcc_h3;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed236 -> {
  SessionState _pat_let118_0 = ((SessionState)(java.lang.Object)(boxed236));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let118_0, boxed237 -> {
    SessionState _17_dt__update__tmp_h4 = ((SessionState)(java.lang.Object)(boxed237));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(((java.util.Objects.equals(_16_m, agency.open.ahp.ConfluxCodec.Json.create_JNull())) ? (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())) : (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _16_m))), boxed238 -> {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let119_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed238));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(_pat_let119_0, boxed239 -> {
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _18_dt__update_hmeta_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed239));
        return SessionState.create((_17_dt__update__tmp_h4).dtor_provider(), (_17_dt__update__tmp_h4).dtor_title(), (_17_dt__update__tmp_h4).dtor_status(), (_17_dt__update__tmp_h4).dtor_lifecycle(), (_17_dt__update__tmp_h4).dtor_activity(), (_17_dt__update__tmp_h4).dtor_config(), _18_dt__update_hmeta_h0, (_17_dt__update__tmp_h4).dtor_creationError(), (_17_dt__update__tmp_h4).dtor_serverTools(), (_17_dt__update__tmp_h4).dtor_changesets(), (_17_dt__update__tmp_h4).dtor_canvases(), (_17_dt__update__tmp_h4).dtor_openCanvases(), (_17_dt__update__tmp_h4).dtor_defaultChat(), (_17_dt__update__tmp_h4).dtor_activeClients(), (_17_dt__update__tmp_h4).dtor_chats(), (_17_dt__update__tmp_h4).dtor_customizations(), (_17_dt__update__tmp_h4).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_IsReadChanged()) {
      boolean _19___mcc_h4 = ((agency.open.ahp.Session.SessionAction_IsReadChanged)_source0)._isRead;
      boolean _20_v = _19___mcc_h4;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed240 -> {
  SessionState _pat_let120_0 = ((SessionState)(java.lang.Object)(boxed240));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let120_0, boxed241 -> {
    SessionState _21_dt__update__tmp_h5 = ((SessionState)(java.lang.Object)(boxed241));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(((_20_v) ? (__default.setBit((_pat_let_tv0).dtor_status(), __default.B__READ())) : (__default.clearBit((_pat_let_tv1).dtor_status(), __default.B__READ()))), boxed242 -> {
      int _pat_let121_0 = ((int)(java.lang.Object)(boxed242));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(_pat_let121_0, boxed243 -> {
        int _22_dt__update_hstatus_h0 = ((int)(java.lang.Object)(boxed243));
        return SessionState.create((_21_dt__update__tmp_h5).dtor_provider(), (_21_dt__update__tmp_h5).dtor_title(), _22_dt__update_hstatus_h0, (_21_dt__update__tmp_h5).dtor_lifecycle(), (_21_dt__update__tmp_h5).dtor_activity(), (_21_dt__update__tmp_h5).dtor_config(), (_21_dt__update__tmp_h5).dtor_meta(), (_21_dt__update__tmp_h5).dtor_creationError(), (_21_dt__update__tmp_h5).dtor_serverTools(), (_21_dt__update__tmp_h5).dtor_changesets(), (_21_dt__update__tmp_h5).dtor_canvases(), (_21_dt__update__tmp_h5).dtor_openCanvases(), (_21_dt__update__tmp_h5).dtor_defaultChat(), (_21_dt__update__tmp_h5).dtor_activeClients(), (_21_dt__update__tmp_h5).dtor_chats(), (_21_dt__update__tmp_h5).dtor_customizations(), (_21_dt__update__tmp_h5).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_IsArchivedChanged()) {
      boolean _23___mcc_h5 = ((agency.open.ahp.Session.SessionAction_IsArchivedChanged)_source0)._isArchived;
      boolean _24_v = _23___mcc_h5;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed244 -> {
  SessionState _pat_let122_0 = ((SessionState)(java.lang.Object)(boxed244));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let122_0, boxed245 -> {
    SessionState _25_dt__update__tmp_h6 = ((SessionState)(java.lang.Object)(boxed245));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(((_24_v) ? (__default.setBit((_pat_let_tv2).dtor_status(), __default.B__ARCH())) : (__default.clearBit((_pat_let_tv3).dtor_status(), __default.B__ARCH()))), boxed246 -> {
      int _pat_let123_0 = ((int)(java.lang.Object)(boxed246));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(_pat_let123_0, boxed247 -> {
        int _26_dt__update_hstatus_h1 = ((int)(java.lang.Object)(boxed247));
        return SessionState.create((_25_dt__update__tmp_h6).dtor_provider(), (_25_dt__update__tmp_h6).dtor_title(), _26_dt__update_hstatus_h1, (_25_dt__update__tmp_h6).dtor_lifecycle(), (_25_dt__update__tmp_h6).dtor_activity(), (_25_dt__update__tmp_h6).dtor_config(), (_25_dt__update__tmp_h6).dtor_meta(), (_25_dt__update__tmp_h6).dtor_creationError(), (_25_dt__update__tmp_h6).dtor_serverTools(), (_25_dt__update__tmp_h6).dtor_changesets(), (_25_dt__update__tmp_h6).dtor_canvases(), (_25_dt__update__tmp_h6).dtor_openCanvases(), (_25_dt__update__tmp_h6).dtor_defaultChat(), (_25_dt__update__tmp_h6).dtor_activeClients(), (_25_dt__update__tmp_h6).dtor_chats(), (_25_dt__update__tmp_h6).dtor_customizations(), (_25_dt__update__tmp_h6).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_ActivityChanged()) {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _27___mcc_h6 = ((agency.open.ahp.Session.SessionAction_ActivityChanged)_source0)._activity;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _28_ac = _27___mcc_h6;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed248 -> {
  SessionState _pat_let124_0 = ((SessionState)(java.lang.Object)(boxed248));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let124_0, boxed249 -> {
    SessionState _29_dt__update__tmp_h7 = ((SessionState)(java.lang.Object)(boxed249));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(_28_ac, boxed250 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let125_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed250));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(_pat_let125_0, boxed251 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _30_dt__update_hactivity_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed251));
        return SessionState.create((_29_dt__update__tmp_h7).dtor_provider(), (_29_dt__update__tmp_h7).dtor_title(), (_29_dt__update__tmp_h7).dtor_status(), (_29_dt__update__tmp_h7).dtor_lifecycle(), _30_dt__update_hactivity_h0, (_29_dt__update__tmp_h7).dtor_config(), (_29_dt__update__tmp_h7).dtor_meta(), (_29_dt__update__tmp_h7).dtor_creationError(), (_29_dt__update__tmp_h7).dtor_serverTools(), (_29_dt__update__tmp_h7).dtor_changesets(), (_29_dt__update__tmp_h7).dtor_canvases(), (_29_dt__update__tmp_h7).dtor_openCanvases(), (_29_dt__update__tmp_h7).dtor_defaultChat(), (_29_dt__update__tmp_h7).dtor_activeClients(), (_29_dt__update__tmp_h7).dtor_chats(), (_29_dt__update__tmp_h7).dtor_customizations(), (_29_dt__update__tmp_h7).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_ConfigChanged()) {
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _31___mcc_h7 = ((agency.open.ahp.Session.SessionAction_ConfigChanged)_source0)._config;
      boolean _32___mcc_h8 = ((agency.open.ahp.Session.SessionAction_ConfigChanged)_source0)._replace;
      boolean _33_repl = _32___mcc_h8;
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _34_cfg = _31___mcc_h7;
      agency.open.ahp.AhpSkeleton.Option<SConfig> _source1 = (s).dtor_config();
      if (_source1.is_None()) {
        return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      } else {
        SConfig _35___mcc_h34 = ((agency.open.ahp.AhpSkeleton.Option_Some<SConfig>)_source1)._value;
        SConfig _36_c = _35___mcc_h34;
        return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed252 -> {
  SessionState _pat_let126_0 = ((SessionState)(java.lang.Object)(boxed252));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let126_0, boxed253 -> {
    SessionState _37_dt__update__tmp_h8 = ((SessionState)(java.lang.Object)(boxed253));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<SConfig>, SessionState>Let(agency.open.ahp.AhpSkeleton.Option.<SConfig>create_Some(SConfig._typeDescriptor(), SConfig.create((_36_c).dtor_schema(), ((_33_repl) ? (_34_cfg) : (dafny.DafnyMap.<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>merge((_36_c).dtor_values(), _34_cfg))))), boxed254 -> {
      agency.open.ahp.AhpSkeleton.Option<SConfig> _pat_let127_0 = ((agency.open.ahp.AhpSkeleton.Option<SConfig>)(java.lang.Object)(boxed254));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<SConfig>, SessionState>Let(_pat_let127_0, boxed255 -> {
        agency.open.ahp.AhpSkeleton.Option<SConfig> _38_dt__update_hconfig_h0 = ((agency.open.ahp.AhpSkeleton.Option<SConfig>)(java.lang.Object)(boxed255));
        return SessionState.create((_37_dt__update__tmp_h8).dtor_provider(), (_37_dt__update__tmp_h8).dtor_title(), (_37_dt__update__tmp_h8).dtor_status(), (_37_dt__update__tmp_h8).dtor_lifecycle(), (_37_dt__update__tmp_h8).dtor_activity(), _38_dt__update_hconfig_h0, (_37_dt__update__tmp_h8).dtor_meta(), (_37_dt__update__tmp_h8).dtor_creationError(), (_37_dt__update__tmp_h8).dtor_serverTools(), (_37_dt__update__tmp_h8).dtor_changesets(), (_37_dt__update__tmp_h8).dtor_canvases(), (_37_dt__update__tmp_h8).dtor_openCanvases(), (_37_dt__update__tmp_h8).dtor_defaultChat(), (_37_dt__update__tmp_h8).dtor_activeClients(), (_37_dt__update__tmp_h8).dtor_chats(), (_37_dt__update__tmp_h8).dtor_customizations(), (_37_dt__update__tmp_h8).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      }
    } else if (_source0.is_CustomizationsChanged()) {
      dafny.DafnySequence<? extends Cust> _39___mcc_h9 = ((agency.open.ahp.Session.SessionAction_CustomizationsChanged)_source0)._customizations;
      dafny.DafnySequence<? extends Cust> _40_cs = _39___mcc_h9;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed256 -> {
  SessionState _pat_let128_0 = ((SessionState)(java.lang.Object)(boxed256));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let128_0, boxed257 -> {
    SessionState _41_dt__update__tmp_h9 = ((SessionState)(java.lang.Object)(boxed257));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_40_cs, boxed258 -> {
      dafny.DafnySequence<? extends Cust> _pat_let129_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed258));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let129_0, boxed259 -> {
        dafny.DafnySequence<? extends Cust> _42_dt__update_hcustomizations_h0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed259));
        return SessionState.create((_41_dt__update__tmp_h9).dtor_provider(), (_41_dt__update__tmp_h9).dtor_title(), (_41_dt__update__tmp_h9).dtor_status(), (_41_dt__update__tmp_h9).dtor_lifecycle(), (_41_dt__update__tmp_h9).dtor_activity(), (_41_dt__update__tmp_h9).dtor_config(), (_41_dt__update__tmp_h9).dtor_meta(), (_41_dt__update__tmp_h9).dtor_creationError(), (_41_dt__update__tmp_h9).dtor_serverTools(), (_41_dt__update__tmp_h9).dtor_changesets(), (_41_dt__update__tmp_h9).dtor_canvases(), (_41_dt__update__tmp_h9).dtor_openCanvases(), (_41_dt__update__tmp_h9).dtor_defaultChat(), (_41_dt__update__tmp_h9).dtor_activeClients(), (_41_dt__update__tmp_h9).dtor_chats(), _42_dt__update_hcustomizations_h0, (_41_dt__update__tmp_h9).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_CustomizationToggled()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _43___mcc_h10 = ((agency.open.ahp.Session.SessionAction_CustomizationToggled)_source0)._id;
      boolean _44___mcc_h11 = ((agency.open.ahp.Session.SessionAction_CustomizationToggled)_source0)._enabled;
      boolean _45_en = _44___mcc_h11;
      dafny.DafnySequence<? extends dafny.CodePoint> _46_id = _43___mcc_h10;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed260 -> {
  SessionState _pat_let130_0 = ((SessionState)(java.lang.Object)(boxed260));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let130_0, boxed261 -> {
    SessionState _47_dt__update__tmp_h10 = ((SessionState)(java.lang.Object)(boxed261));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(__default.upsertCust((_pat_let_tv4).dtor_customizations(), _46_id, _45_en), boxed262 -> {
      dafny.DafnySequence<? extends Cust> _pat_let131_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed262));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let131_0, boxed263 -> {
        dafny.DafnySequence<? extends Cust> _48_dt__update_hcustomizations_h1 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed263));
        return SessionState.create((_47_dt__update__tmp_h10).dtor_provider(), (_47_dt__update__tmp_h10).dtor_title(), (_47_dt__update__tmp_h10).dtor_status(), (_47_dt__update__tmp_h10).dtor_lifecycle(), (_47_dt__update__tmp_h10).dtor_activity(), (_47_dt__update__tmp_h10).dtor_config(), (_47_dt__update__tmp_h10).dtor_meta(), (_47_dt__update__tmp_h10).dtor_creationError(), (_47_dt__update__tmp_h10).dtor_serverTools(), (_47_dt__update__tmp_h10).dtor_changesets(), (_47_dt__update__tmp_h10).dtor_canvases(), (_47_dt__update__tmp_h10).dtor_openCanvases(), (_47_dt__update__tmp_h10).dtor_defaultChat(), (_47_dt__update__tmp_h10).dtor_activeClients(), (_47_dt__update__tmp_h10).dtor_chats(), _48_dt__update_hcustomizations_h1, (_47_dt__update__tmp_h10).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_CustomizationRemoved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _49___mcc_h12 = ((agency.open.ahp.Session.SessionAction_CustomizationRemoved)_source0)._id;
      dafny.DafnySequence<? extends dafny.CodePoint> _50_id = _49___mcc_h12;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed264 -> {
  SessionState _pat_let132_0 = ((SessionState)(java.lang.Object)(boxed264));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let132_0, boxed265 -> {
    SessionState _51_dt__update__tmp_h11 = ((SessionState)(java.lang.Object)(boxed265));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(__default.removeCust((_pat_let_tv5).dtor_customizations(), _50_id), boxed266 -> {
      dafny.DafnySequence<? extends Cust> _pat_let133_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed266));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let133_0, boxed267 -> {
        dafny.DafnySequence<? extends Cust> _52_dt__update_hcustomizations_h2 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed267));
        return SessionState.create((_51_dt__update__tmp_h11).dtor_provider(), (_51_dt__update__tmp_h11).dtor_title(), (_51_dt__update__tmp_h11).dtor_status(), (_51_dt__update__tmp_h11).dtor_lifecycle(), (_51_dt__update__tmp_h11).dtor_activity(), (_51_dt__update__tmp_h11).dtor_config(), (_51_dt__update__tmp_h11).dtor_meta(), (_51_dt__update__tmp_h11).dtor_creationError(), (_51_dt__update__tmp_h11).dtor_serverTools(), (_51_dt__update__tmp_h11).dtor_changesets(), (_51_dt__update__tmp_h11).dtor_canvases(), (_51_dt__update__tmp_h11).dtor_openCanvases(), (_51_dt__update__tmp_h11).dtor_defaultChat(), (_51_dt__update__tmp_h11).dtor_activeClients(), (_51_dt__update__tmp_h11).dtor_chats(), _52_dt__update_hcustomizations_h2, (_51_dt__update__tmp_h11).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_DefaultChatChanged()) {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _53___mcc_h13 = ((agency.open.ahp.Session.SessionAction_DefaultChatChanged)_source0)._chat;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _54_c = _53___mcc_h13;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed268 -> {
  SessionState _pat_let134_0 = ((SessionState)(java.lang.Object)(boxed268));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let134_0, boxed269 -> {
    SessionState _55_dt__update__tmp_h12 = ((SessionState)(java.lang.Object)(boxed269));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(_54_c, boxed270 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let135_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed270));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(_pat_let135_0, boxed271 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _56_dt__update_hdefaultChat_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed271));
        return SessionState.create((_55_dt__update__tmp_h12).dtor_provider(), (_55_dt__update__tmp_h12).dtor_title(), (_55_dt__update__tmp_h12).dtor_status(), (_55_dt__update__tmp_h12).dtor_lifecycle(), (_55_dt__update__tmp_h12).dtor_activity(), (_55_dt__update__tmp_h12).dtor_config(), (_55_dt__update__tmp_h12).dtor_meta(), (_55_dt__update__tmp_h12).dtor_creationError(), (_55_dt__update__tmp_h12).dtor_serverTools(), (_55_dt__update__tmp_h12).dtor_changesets(), (_55_dt__update__tmp_h12).dtor_canvases(), (_55_dt__update__tmp_h12).dtor_openCanvases(), _56_dt__update_hdefaultChat_h0, (_55_dt__update__tmp_h12).dtor_activeClients(), (_55_dt__update__tmp_h12).dtor_chats(), (_55_dt__update__tmp_h12).dtor_customizations(), (_55_dt__update__tmp_h12).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_ChatAdded()) {
      Chat _57___mcc_h14 = ((agency.open.ahp.Session.SessionAction_ChatAdded)_source0)._summary;
      Chat _58_c = _57___mcc_h14;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed272 -> {
  SessionState _pat_let136_0 = ((SessionState)(java.lang.Object)(boxed272));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let136_0, boxed273 -> {
    SessionState _59_dt__update__tmp_h13 = ((SessionState)(java.lang.Object)(boxed273));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(__default.upsertChat((_pat_let_tv6).dtor_chats(), _58_c), boxed274 -> {
      dafny.DafnySequence<? extends Chat> _pat_let137_0 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed274));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(_pat_let137_0, boxed275 -> {
        dafny.DafnySequence<? extends Chat> _60_dt__update_hchats_h0 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed275));
        return SessionState.create((_59_dt__update__tmp_h13).dtor_provider(), (_59_dt__update__tmp_h13).dtor_title(), (_59_dt__update__tmp_h13).dtor_status(), (_59_dt__update__tmp_h13).dtor_lifecycle(), (_59_dt__update__tmp_h13).dtor_activity(), (_59_dt__update__tmp_h13).dtor_config(), (_59_dt__update__tmp_h13).dtor_meta(), (_59_dt__update__tmp_h13).dtor_creationError(), (_59_dt__update__tmp_h13).dtor_serverTools(), (_59_dt__update__tmp_h13).dtor_changesets(), (_59_dt__update__tmp_h13).dtor_canvases(), (_59_dt__update__tmp_h13).dtor_openCanvases(), (_59_dt__update__tmp_h13).dtor_defaultChat(), (_59_dt__update__tmp_h13).dtor_activeClients(), _60_dt__update_hchats_h0, (_59_dt__update__tmp_h13).dtor_customizations(), (_59_dt__update__tmp_h13).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_ChatRemoved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _61___mcc_h15 = ((agency.open.ahp.Session.SessionAction_ChatRemoved)_source0)._resource;
      dafny.DafnySequence<? extends dafny.CodePoint> _62_r = _61___mcc_h15;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed276 -> {
  SessionState _pat_let138_0 = ((SessionState)(java.lang.Object)(boxed276));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let138_0, boxed277 -> {
    SessionState _63_dt__update__tmp_h14 = ((SessionState)(java.lang.Object)(boxed277));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(((java.util.Objects.equals((_pat_let_tv8).dtor_defaultChat(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), _62_r))) ? (agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))) : ((_pat_let_tv7).dtor_defaultChat())), boxed278 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let139_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed278));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(_pat_let139_0, boxed279 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _64_dt__update_hdefaultChat_h1 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed279));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(__default.removeChat((_pat_let_tv9).dtor_chats(), _62_r), boxed280 -> {
          dafny.DafnySequence<? extends Chat> _pat_let140_0 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed280));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(_pat_let140_0, boxed281 -> {
            dafny.DafnySequence<? extends Chat> _65_dt__update_hchats_h1 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed281));
            return SessionState.create((_63_dt__update__tmp_h14).dtor_provider(), (_63_dt__update__tmp_h14).dtor_title(), (_63_dt__update__tmp_h14).dtor_status(), (_63_dt__update__tmp_h14).dtor_lifecycle(), (_63_dt__update__tmp_h14).dtor_activity(), (_63_dt__update__tmp_h14).dtor_config(), (_63_dt__update__tmp_h14).dtor_meta(), (_63_dt__update__tmp_h14).dtor_creationError(), (_63_dt__update__tmp_h14).dtor_serverTools(), (_63_dt__update__tmp_h14).dtor_changesets(), (_63_dt__update__tmp_h14).dtor_canvases(), (_63_dt__update__tmp_h14).dtor_openCanvases(), _64_dt__update_hdefaultChat_h1, (_63_dt__update__tmp_h14).dtor_activeClients(), _65_dt__update_hchats_h1, (_63_dt__update__tmp_h14).dtor_customizations(), (_63_dt__update__tmp_h14).dtor_inputNeeded());
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
    } else if (_source0.is_ChatUpdated()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _66___mcc_h16 = ((agency.open.ahp.Session.SessionAction_ChatUpdated)_source0)._resource;
      agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _67___mcc_h17 = ((agency.open.ahp.Session.SessionAction_ChatUpdated)_source0)._status;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _68___mcc_h18 = ((agency.open.ahp.Session.SessionAction_ChatUpdated)_source0)._activity;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _69___mcc_h19 = ((agency.open.ahp.Session.SessionAction_ChatUpdated)_source0)._modifiedAt;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _70_md = _69___mcc_h19;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _71_ac = _68___mcc_h18;
      agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _72_st = _67___mcc_h17;
      dafny.DafnySequence<? extends dafny.CodePoint> _73_r = _66___mcc_h16;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed282 -> {
  SessionState _pat_let141_0 = ((SessionState)(java.lang.Object)(boxed282));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let141_0, boxed283 -> {
    SessionState _74_dt__update__tmp_h15 = ((SessionState)(java.lang.Object)(boxed283));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(__default.updChat((_pat_let_tv10).dtor_chats(), _73_r, _72_st, _71_ac, _70_md), boxed284 -> {
      dafny.DafnySequence<? extends Chat> _pat_let142_0 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed284));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(_pat_let142_0, boxed285 -> {
        dafny.DafnySequence<? extends Chat> _75_dt__update_hchats_h2 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed285));
        return SessionState.create((_74_dt__update__tmp_h15).dtor_provider(), (_74_dt__update__tmp_h15).dtor_title(), (_74_dt__update__tmp_h15).dtor_status(), (_74_dt__update__tmp_h15).dtor_lifecycle(), (_74_dt__update__tmp_h15).dtor_activity(), (_74_dt__update__tmp_h15).dtor_config(), (_74_dt__update__tmp_h15).dtor_meta(), (_74_dt__update__tmp_h15).dtor_creationError(), (_74_dt__update__tmp_h15).dtor_serverTools(), (_74_dt__update__tmp_h15).dtor_changesets(), (_74_dt__update__tmp_h15).dtor_canvases(), (_74_dt__update__tmp_h15).dtor_openCanvases(), (_74_dt__update__tmp_h15).dtor_defaultChat(), (_74_dt__update__tmp_h15).dtor_activeClients(), _75_dt__update_hchats_h2, (_74_dt__update__tmp_h15).dtor_customizations(), (_74_dt__update__tmp_h15).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_ChangesetsChanged()) {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _76___mcc_h20 = ((agency.open.ahp.Session.SessionAction_ChangesetsChanged)_source0)._changesets;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _77_cs = _76___mcc_h20;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed286 -> {
  SessionState _pat_let143_0 = ((SessionState)(java.lang.Object)(boxed286));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let143_0, boxed287 -> {
    SessionState _78_dt__update__tmp_h16 = ((SessionState)(java.lang.Object)(boxed287));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(_77_cs, boxed288 -> {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let144_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed288));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(_pat_let144_0, boxed289 -> {
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _79_dt__update_hchangesets_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed289));
        return SessionState.create((_78_dt__update__tmp_h16).dtor_provider(), (_78_dt__update__tmp_h16).dtor_title(), (_78_dt__update__tmp_h16).dtor_status(), (_78_dt__update__tmp_h16).dtor_lifecycle(), (_78_dt__update__tmp_h16).dtor_activity(), (_78_dt__update__tmp_h16).dtor_config(), (_78_dt__update__tmp_h16).dtor_meta(), (_78_dt__update__tmp_h16).dtor_creationError(), (_78_dt__update__tmp_h16).dtor_serverTools(), _79_dt__update_hchangesets_h0, (_78_dt__update__tmp_h16).dtor_canvases(), (_78_dt__update__tmp_h16).dtor_openCanvases(), (_78_dt__update__tmp_h16).dtor_defaultChat(), (_78_dt__update__tmp_h16).dtor_activeClients(), (_78_dt__update__tmp_h16).dtor_chats(), (_78_dt__update__tmp_h16).dtor_customizations(), (_78_dt__update__tmp_h16).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_ActiveClientSet()) {
      Client _80___mcc_h21 = ((agency.open.ahp.Session.SessionAction_ActiveClientSet)_source0)._client;
      Client _81_c = _80___mcc_h21;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed290 -> {
  SessionState _pat_let145_0 = ((SessionState)(java.lang.Object)(boxed290));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let145_0, boxed291 -> {
    SessionState _82_dt__update__tmp_h19 = ((SessionState)(java.lang.Object)(boxed291));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Client>, SessionState>Let(__default.upsertClient((_pat_let_tv11).dtor_activeClients(), _81_c), boxed292 -> {
      dafny.DafnySequence<? extends Client> _pat_let146_0 = ((dafny.DafnySequence<? extends Client>)(java.lang.Object)(boxed292));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Client>, SessionState>Let(_pat_let146_0, boxed293 -> {
        dafny.DafnySequence<? extends Client> _83_dt__update_hactiveClients_h0 = ((dafny.DafnySequence<? extends Client>)(java.lang.Object)(boxed293));
        return SessionState.create((_82_dt__update__tmp_h19).dtor_provider(), (_82_dt__update__tmp_h19).dtor_title(), (_82_dt__update__tmp_h19).dtor_status(), (_82_dt__update__tmp_h19).dtor_lifecycle(), (_82_dt__update__tmp_h19).dtor_activity(), (_82_dt__update__tmp_h19).dtor_config(), (_82_dt__update__tmp_h19).dtor_meta(), (_82_dt__update__tmp_h19).dtor_creationError(), (_82_dt__update__tmp_h19).dtor_serverTools(), (_82_dt__update__tmp_h19).dtor_changesets(), (_82_dt__update__tmp_h19).dtor_canvases(), (_82_dt__update__tmp_h19).dtor_openCanvases(), (_82_dt__update__tmp_h19).dtor_defaultChat(), _83_dt__update_hactiveClients_h0, (_82_dt__update__tmp_h19).dtor_chats(), (_82_dt__update__tmp_h19).dtor_customizations(), (_82_dt__update__tmp_h19).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_ActiveClientRemoved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _84___mcc_h22 = ((agency.open.ahp.Session.SessionAction_ActiveClientRemoved)_source0)._clientId;
      dafny.DafnySequence<? extends dafny.CodePoint> _85_id = _84___mcc_h22;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed294 -> {
  SessionState _pat_let147_0 = ((SessionState)(java.lang.Object)(boxed294));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let147_0, boxed295 -> {
    SessionState _86_dt__update__tmp_h20 = ((SessionState)(java.lang.Object)(boxed295));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Client>, SessionState>Let(__default.removeClient((_pat_let_tv12).dtor_activeClients(), _85_id), boxed296 -> {
      dafny.DafnySequence<? extends Client> _pat_let148_0 = ((dafny.DafnySequence<? extends Client>)(java.lang.Object)(boxed296));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Client>, SessionState>Let(_pat_let148_0, boxed297 -> {
        dafny.DafnySequence<? extends Client> _87_dt__update_hactiveClients_h1 = ((dafny.DafnySequence<? extends Client>)(java.lang.Object)(boxed297));
        return SessionState.create((_86_dt__update__tmp_h20).dtor_provider(), (_86_dt__update__tmp_h20).dtor_title(), (_86_dt__update__tmp_h20).dtor_status(), (_86_dt__update__tmp_h20).dtor_lifecycle(), (_86_dt__update__tmp_h20).dtor_activity(), (_86_dt__update__tmp_h20).dtor_config(), (_86_dt__update__tmp_h20).dtor_meta(), (_86_dt__update__tmp_h20).dtor_creationError(), (_86_dt__update__tmp_h20).dtor_serverTools(), (_86_dt__update__tmp_h20).dtor_changesets(), (_86_dt__update__tmp_h20).dtor_canvases(), (_86_dt__update__tmp_h20).dtor_openCanvases(), (_86_dt__update__tmp_h20).dtor_defaultChat(), _87_dt__update_hactiveClients_h1, (_86_dt__update__tmp_h20).dtor_chats(), (_86_dt__update__tmp_h20).dtor_customizations(), (_86_dt__update__tmp_h20).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_CanvasesChanged()) {
      agency.open.ahp.ConfluxCodec.Json _88___mcc_h23 = ((agency.open.ahp.Session.SessionAction_CanvasesChanged)_source0)._canvases;
      agency.open.ahp.ConfluxCodec.Json _89_cs = _88___mcc_h23;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed298 -> {
  SessionState _pat_let149_0 = ((SessionState)(java.lang.Object)(boxed298));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let149_0, boxed299 -> {
    SessionState _90_dt__update__tmp_h17 = ((SessionState)(java.lang.Object)(boxed299));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(((java.util.Objects.equals(_89_cs, agency.open.ahp.ConfluxCodec.Json.create_JNull())) ? (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())) : (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _89_cs))), boxed300 -> {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let150_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed300));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(_pat_let150_0, boxed301 -> {
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _91_dt__update_hcanvases_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed301));
        return SessionState.create((_90_dt__update__tmp_h17).dtor_provider(), (_90_dt__update__tmp_h17).dtor_title(), (_90_dt__update__tmp_h17).dtor_status(), (_90_dt__update__tmp_h17).dtor_lifecycle(), (_90_dt__update__tmp_h17).dtor_activity(), (_90_dt__update__tmp_h17).dtor_config(), (_90_dt__update__tmp_h17).dtor_meta(), (_90_dt__update__tmp_h17).dtor_creationError(), (_90_dt__update__tmp_h17).dtor_serverTools(), (_90_dt__update__tmp_h17).dtor_changesets(), _91_dt__update_hcanvases_h0, (_90_dt__update__tmp_h17).dtor_openCanvases(), (_90_dt__update__tmp_h17).dtor_defaultChat(), (_90_dt__update__tmp_h17).dtor_activeClients(), (_90_dt__update__tmp_h17).dtor_chats(), (_90_dt__update__tmp_h17).dtor_customizations(), (_90_dt__update__tmp_h17).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_OpenCanvasesChanged()) {
      agency.open.ahp.ConfluxCodec.Json _92___mcc_h24 = ((agency.open.ahp.Session.SessionAction_OpenCanvasesChanged)_source0)._openCanvases;
      agency.open.ahp.ConfluxCodec.Json _93_cs = _92___mcc_h24;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed302 -> {
  SessionState _pat_let151_0 = ((SessionState)(java.lang.Object)(boxed302));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let151_0, boxed303 -> {
    SessionState _94_dt__update__tmp_h18 = ((SessionState)(java.lang.Object)(boxed303));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(((java.util.Objects.equals(_93_cs, agency.open.ahp.ConfluxCodec.Json.create_JNull())) ? (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())) : (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _93_cs))), boxed304 -> {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let152_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed304));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(_pat_let152_0, boxed305 -> {
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _95_dt__update_hopenCanvases_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed305));
        return SessionState.create((_94_dt__update__tmp_h18).dtor_provider(), (_94_dt__update__tmp_h18).dtor_title(), (_94_dt__update__tmp_h18).dtor_status(), (_94_dt__update__tmp_h18).dtor_lifecycle(), (_94_dt__update__tmp_h18).dtor_activity(), (_94_dt__update__tmp_h18).dtor_config(), (_94_dt__update__tmp_h18).dtor_meta(), (_94_dt__update__tmp_h18).dtor_creationError(), (_94_dt__update__tmp_h18).dtor_serverTools(), (_94_dt__update__tmp_h18).dtor_changesets(), (_94_dt__update__tmp_h18).dtor_canvases(), _95_dt__update_hopenCanvases_h0, (_94_dt__update__tmp_h18).dtor_defaultChat(), (_94_dt__update__tmp_h18).dtor_activeClients(), (_94_dt__update__tmp_h18).dtor_chats(), (_94_dt__update__tmp_h18).dtor_customizations(), (_94_dt__update__tmp_h18).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_InputNeededSet()) {
      InputReq _96___mcc_h25 = ((agency.open.ahp.Session.SessionAction_InputNeededSet)_source0)._req;
      InputReq _97_r = _96___mcc_h25;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed306 -> {
  SessionState _pat_let153_0 = ((SessionState)(java.lang.Object)(boxed306));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let153_0, boxed307 -> {
    SessionState _98_dt__update__tmp_h21 = ((SessionState)(java.lang.Object)(boxed307));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(__default.setBit((_pat_let_tv13).dtor_status(), __default.B__INPUT()), boxed308 -> {
      int _pat_let154_0 = ((int)(java.lang.Object)(boxed308));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(_pat_let154_0, boxed309 -> {
        int _99_dt__update_hstatus_h2 = ((int)(java.lang.Object)(boxed309));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends InputReq>, SessionState>Let(__default.upsertInput((_pat_let_tv14).dtor_inputNeeded(), _97_r), boxed310 -> {
          dafny.DafnySequence<? extends InputReq> _pat_let155_0 = ((dafny.DafnySequence<? extends InputReq>)(java.lang.Object)(boxed310));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends InputReq>, SessionState>Let(_pat_let155_0, boxed311 -> {
            dafny.DafnySequence<? extends InputReq> _100_dt__update_hinputNeeded_h0 = ((dafny.DafnySequence<? extends InputReq>)(java.lang.Object)(boxed311));
            return SessionState.create((_98_dt__update__tmp_h21).dtor_provider(), (_98_dt__update__tmp_h21).dtor_title(), _99_dt__update_hstatus_h2, (_98_dt__update__tmp_h21).dtor_lifecycle(), (_98_dt__update__tmp_h21).dtor_activity(), (_98_dt__update__tmp_h21).dtor_config(), (_98_dt__update__tmp_h21).dtor_meta(), (_98_dt__update__tmp_h21).dtor_creationError(), (_98_dt__update__tmp_h21).dtor_serverTools(), (_98_dt__update__tmp_h21).dtor_changesets(), (_98_dt__update__tmp_h21).dtor_canvases(), (_98_dt__update__tmp_h21).dtor_openCanvases(), (_98_dt__update__tmp_h21).dtor_defaultChat(), (_98_dt__update__tmp_h21).dtor_activeClients(), (_98_dt__update__tmp_h21).dtor_chats(), (_98_dt__update__tmp_h21).dtor_customizations(), _100_dt__update_hinputNeeded_h0);
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
    } else if (_source0.is_InputNeededRemoved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _101___mcc_h26 = ((agency.open.ahp.Session.SessionAction_InputNeededRemoved)_source0)._id;
      dafny.DafnySequence<? extends dafny.CodePoint> _102_id = _101___mcc_h26;
      dafny.DafnySequence<? extends InputReq> _103_rem = __default.removeInput((s).dtor_inputNeeded(), _102_id);
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed312 -> {
  SessionState _pat_let156_0 = ((SessionState)(java.lang.Object)(boxed312));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let156_0, boxed313 -> {
    SessionState _104_dt__update__tmp_h22 = ((SessionState)(java.lang.Object)(boxed313));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let((((java.math.BigInteger.valueOf((_103_rem).length())).signum() == 0) ? (__default.clearBit((_pat_let_tv15).dtor_status(), __default.B__INPUT())) : ((_pat_let_tv16).dtor_status())), boxed314 -> {
      int _pat_let157_0 = ((int)(java.lang.Object)(boxed314));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(_pat_let157_0, boxed315 -> {
        int _105_dt__update_hstatus_h3 = ((int)(java.lang.Object)(boxed315));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends InputReq>, SessionState>Let(_103_rem, boxed316 -> {
          dafny.DafnySequence<? extends InputReq> _pat_let158_0 = ((dafny.DafnySequence<? extends InputReq>)(java.lang.Object)(boxed316));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends InputReq>, SessionState>Let(_pat_let158_0, boxed317 -> {
            dafny.DafnySequence<? extends InputReq> _106_dt__update_hinputNeeded_h1 = ((dafny.DafnySequence<? extends InputReq>)(java.lang.Object)(boxed317));
            return SessionState.create((_104_dt__update__tmp_h22).dtor_provider(), (_104_dt__update__tmp_h22).dtor_title(), _105_dt__update_hstatus_h3, (_104_dt__update__tmp_h22).dtor_lifecycle(), (_104_dt__update__tmp_h22).dtor_activity(), (_104_dt__update__tmp_h22).dtor_config(), (_104_dt__update__tmp_h22).dtor_meta(), (_104_dt__update__tmp_h22).dtor_creationError(), (_104_dt__update__tmp_h22).dtor_serverTools(), (_104_dt__update__tmp_h22).dtor_changesets(), (_104_dt__update__tmp_h22).dtor_canvases(), (_104_dt__update__tmp_h22).dtor_openCanvases(), (_104_dt__update__tmp_h22).dtor_defaultChat(), (_104_dt__update__tmp_h22).dtor_activeClients(), (_104_dt__update__tmp_h22).dtor_chats(), (_104_dt__update__tmp_h22).dtor_customizations(), _106_dt__update_hinputNeeded_h1);
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
    } else if (_source0.is_CustomizationUpdated()) {
      Cust _107___mcc_h27 = ((agency.open.ahp.Session.SessionAction_CustomizationUpdated)_source0)._customization;
      Cust _108_c = _107___mcc_h27;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed318 -> {
  SessionState _pat_let159_0 = ((SessionState)(java.lang.Object)(boxed318));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let159_0, boxed319 -> {
    SessionState _109_dt__update__tmp_h23 = ((SessionState)(java.lang.Object)(boxed319));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(__default.upsertCustFull((_pat_let_tv17).dtor_customizations(), _108_c), boxed320 -> {
      dafny.DafnySequence<? extends Cust> _pat_let160_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed320));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let160_0, boxed321 -> {
        dafny.DafnySequence<? extends Cust> _110_dt__update_hcustomizations_h3 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed321));
        return SessionState.create((_109_dt__update__tmp_h23).dtor_provider(), (_109_dt__update__tmp_h23).dtor_title(), (_109_dt__update__tmp_h23).dtor_status(), (_109_dt__update__tmp_h23).dtor_lifecycle(), (_109_dt__update__tmp_h23).dtor_activity(), (_109_dt__update__tmp_h23).dtor_config(), (_109_dt__update__tmp_h23).dtor_meta(), (_109_dt__update__tmp_h23).dtor_creationError(), (_109_dt__update__tmp_h23).dtor_serverTools(), (_109_dt__update__tmp_h23).dtor_changesets(), (_109_dt__update__tmp_h23).dtor_canvases(), (_109_dt__update__tmp_h23).dtor_openCanvases(), (_109_dt__update__tmp_h23).dtor_defaultChat(), (_109_dt__update__tmp_h23).dtor_activeClients(), (_109_dt__update__tmp_h23).dtor_chats(), _110_dt__update_hcustomizations_h3, (_109_dt__update__tmp_h23).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_McpServerStateChanged()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _111___mcc_h28 = ((agency.open.ahp.Session.SessionAction_McpServerStateChanged)_source0)._id;
      agency.open.ahp.ConfluxCodec.Json _112___mcc_h29 = ((agency.open.ahp.Session.SessionAction_McpServerStateChanged)_source0)._state;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _113___mcc_h30 = ((agency.open.ahp.Session.SessionAction_McpServerStateChanged)_source0)._channel;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _114_ch = _113___mcc_h30;
      agency.open.ahp.ConfluxCodec.Json _115_st = _112___mcc_h29;
      dafny.DafnySequence<? extends dafny.CodePoint> _116_id = _111___mcc_h28;
      if (__default.isMcp((s).dtor_customizations(), _116_id)) {
        return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed322 -> {
  SessionState _pat_let161_0 = ((SessionState)(java.lang.Object)(boxed322));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let161_0, boxed323 -> {
    SessionState _117_dt__update__tmp_h24 = ((SessionState)(java.lang.Object)(boxed323));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(__default.setMcp((_pat_let_tv18).dtor_customizations(), _116_id, ((java.util.Objects.equals(_115_st, agency.open.ahp.ConfluxCodec.Json.create_JNull())) ? (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())) : (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _115_st))), _114_ch), boxed324 -> {
      dafny.DafnySequence<? extends Cust> _pat_let162_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed324));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let162_0, boxed325 -> {
        dafny.DafnySequence<? extends Cust> _118_dt__update_hcustomizations_h4 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed325));
        return SessionState.create((_117_dt__update__tmp_h24).dtor_provider(), (_117_dt__update__tmp_h24).dtor_title(), (_117_dt__update__tmp_h24).dtor_status(), (_117_dt__update__tmp_h24).dtor_lifecycle(), (_117_dt__update__tmp_h24).dtor_activity(), (_117_dt__update__tmp_h24).dtor_config(), (_117_dt__update__tmp_h24).dtor_meta(), (_117_dt__update__tmp_h24).dtor_creationError(), (_117_dt__update__tmp_h24).dtor_serverTools(), (_117_dt__update__tmp_h24).dtor_changesets(), (_117_dt__update__tmp_h24).dtor_canvases(), (_117_dt__update__tmp_h24).dtor_openCanvases(), (_117_dt__update__tmp_h24).dtor_defaultChat(), (_117_dt__update__tmp_h24).dtor_activeClients(), (_117_dt__update__tmp_h24).dtor_chats(), _118_dt__update_hcustomizations_h4, (_117_dt__update__tmp_h24).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_McpServerStartRequested()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _119___mcc_h31 = ((agency.open.ahp.Session.SessionAction_McpServerStartRequested)_source0)._id;
      dafny.DafnySequence<? extends dafny.CodePoint> _120_id = _119___mcc_h31;
      if (__default.isMcp((s).dtor_customizations(), _120_id)) {
        return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed326 -> {
  SessionState _pat_let163_0 = ((SessionState)(java.lang.Object)(boxed326));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let163_0, boxed327 -> {
    SessionState _121_dt__update__tmp_h25 = ((SessionState)(java.lang.Object)(boxed327));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(__default.setMcp((_pat_let_tv19).dtor_customizations(), _120_id, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("starting"))) }))), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), boxed328 -> {
      dafny.DafnySequence<? extends Cust> _pat_let164_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed328));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let164_0, boxed329 -> {
        dafny.DafnySequence<? extends Cust> _122_dt__update_hcustomizations_h5 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed329));
        return SessionState.create((_121_dt__update__tmp_h25).dtor_provider(), (_121_dt__update__tmp_h25).dtor_title(), (_121_dt__update__tmp_h25).dtor_status(), (_121_dt__update__tmp_h25).dtor_lifecycle(), (_121_dt__update__tmp_h25).dtor_activity(), (_121_dt__update__tmp_h25).dtor_config(), (_121_dt__update__tmp_h25).dtor_meta(), (_121_dt__update__tmp_h25).dtor_creationError(), (_121_dt__update__tmp_h25).dtor_serverTools(), (_121_dt__update__tmp_h25).dtor_changesets(), (_121_dt__update__tmp_h25).dtor_canvases(), (_121_dt__update__tmp_h25).dtor_openCanvases(), (_121_dt__update__tmp_h25).dtor_defaultChat(), (_121_dt__update__tmp_h25).dtor_activeClients(), (_121_dt__update__tmp_h25).dtor_chats(), _122_dt__update_hcustomizations_h5, (_121_dt__update__tmp_h25).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_McpServerStopRequested()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _123___mcc_h32 = ((agency.open.ahp.Session.SessionAction_McpServerStopRequested)_source0)._id;
      dafny.DafnySequence<? extends dafny.CodePoint> _124_id = _123___mcc_h32;
      if (__default.isMcp((s).dtor_customizations(), _124_id)) {
        return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(s, boxed330 -> {
  SessionState _pat_let165_0 = ((SessionState)(java.lang.Object)(boxed330));
  return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let165_0, boxed331 -> {
    SessionState _125_dt__update__tmp_h26 = ((SessionState)(java.lang.Object)(boxed331));
    return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(__default.setMcp((_pat_let_tv20).dtor_customizations(), _124_id, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("stopped"))) }))), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), boxed332 -> {
      dafny.DafnySequence<? extends Cust> _pat_let166_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed332));
      return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let166_0, boxed333 -> {
        dafny.DafnySequence<? extends Cust> _126_dt__update_hcustomizations_h6 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed333));
        return SessionState.create((_125_dt__update__tmp_h26).dtor_provider(), (_125_dt__update__tmp_h26).dtor_title(), (_125_dt__update__tmp_h26).dtor_status(), (_125_dt__update__tmp_h26).dtor_lifecycle(), (_125_dt__update__tmp_h26).dtor_activity(), (_125_dt__update__tmp_h26).dtor_config(), (_125_dt__update__tmp_h26).dtor_meta(), (_125_dt__update__tmp_h26).dtor_creationError(), (_125_dt__update__tmp_h26).dtor_serverTools(), (_125_dt__update__tmp_h26).dtor_changesets(), (_125_dt__update__tmp_h26).dtor_canvases(), (_125_dt__update__tmp_h26).dtor_openCanvases(), (_125_dt__update__tmp_h26).dtor_defaultChat(), (_125_dt__update__tmp_h26).dtor_activeClients(), (_125_dt__update__tmp_h26).dtor_chats(), _126_dt__update_hcustomizations_h6, (_125_dt__update__tmp_h26).dtor_inputNeeded());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else {
      agency.open.ahp.ConfluxCodec.Json _127___mcc_h33 = ((agency.open.ahp.Session.SessionAction_SUnknown)_source0)._raw;
      return dafny.Tuple2.<SessionState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
    }
  }
  public static boolean WfCust(Cust c) {
    return (!(((c).dtor_state()).is_Some()) || (!java.util.Objects.equals(((c).dtor_state()).dtor_value(), agency.open.ahp.ConfluxCodec.Json.create_JNull()))) && (!(((c).dtor_channel()).is_Some()) || (!java.util.Objects.equals(((c).dtor_channel()).dtor_value(), agency.open.ahp.ConfluxCodec.Json.create_JNull())));
  }
  public static boolean WfCusts(dafny.DafnySequence<? extends Cust> cs) {
    return ((java.util.function.Function<dafny.DafnySequence<? extends Cust>, Boolean>)(_0_cs) -> {return dafny.Helpers.Quantifier(dafny.Helpers.IntegerRange(java.math.BigInteger.ZERO, java.math.BigInteger.valueOf((_0_cs).length())), true, ((_forall_var_0_boxed0) -> {
      java.math.BigInteger _forall_var_0 = ((java.math.BigInteger)(java.lang.Object)(_forall_var_0_boxed0));
      java.math.BigInteger _1_i = (java.math.BigInteger)_forall_var_0;
      return !(((_1_i).signum() != -1) && ((_1_i).compareTo(java.math.BigInteger.valueOf((_0_cs).length())) < 0)) || (__default.WfCust(((Cust)(java.lang.Object)((_0_cs).select(dafny.Helpers.toInt((_1_i)))))));
    }));}).apply(cs);
  }
  public static boolean SessionWf(SessionState s) {
    return ((((((((((!(((s).dtor_meta()).is_Some()) || (!java.util.Objects.equals(((s).dtor_meta()).dtor_value(), agency.open.ahp.ConfluxCodec.Json.create_JNull()))) && (!(((s).dtor_creationError()).is_Some()) || (!java.util.Objects.equals(((s).dtor_creationError()).dtor_value(), agency.open.ahp.ConfluxCodec.Json.create_JNull())))) && (!(((s).dtor_serverTools()).is_Some()) || (!java.util.Objects.equals(((s).dtor_serverTools()).dtor_value(), agency.open.ahp.ConfluxCodec.Json.create_JNull())))) && (!(((s).dtor_changesets()).is_Some()) || (!java.util.Objects.equals(((s).dtor_changesets()).dtor_value(), agency.open.ahp.ConfluxCodec.Json.create_JNull())))) && (!(((s).dtor_canvases()).is_Some()) || (!java.util.Objects.equals(((s).dtor_canvases()).dtor_value(), agency.open.ahp.ConfluxCodec.Json.create_JNull())))) && (!(((s).dtor_openCanvases()).is_Some()) || (!java.util.Objects.equals(((s).dtor_openCanvases()).dtor_value(), agency.open.ahp.ConfluxCodec.Json.create_JNull())))) && (__default.WfCusts((s).dtor_customizations()))) && (agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Cust>UniqueKeys(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Cust._typeDescriptor(), __default.custKey(), (s).dtor_customizations()))) && (agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Chat>UniqueKeys(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Chat._typeDescriptor(), __default.chatKey(), (s).dtor_chats()))) && (agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Client>UniqueKeys(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Client._typeDescriptor(), __default.clientKey(), (s).dtor_activeClients()))) && (agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, InputReq>UniqueKeys(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), InputReq._typeDescriptor(), __default.inputKey(), (s).dtor_inputNeeded()));
  }
  public static boolean WfSessionActionInv(SessionAction a) {
    SessionAction _source0 = a;
    if (_source0.is_Ready()) {
      return true;
    } else if (_source0.is_CreationFailed()) {
      agency.open.ahp.ConfluxCodec.Json _0___mcc_h0 = ((agency.open.ahp.Session.SessionAction_CreationFailed)_source0)._error;
      return true;
    } else if (_source0.is_TitleChanged()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _1___mcc_h2 = ((agency.open.ahp.Session.SessionAction_TitleChanged)_source0)._title;
      return true;
    } else if (_source0.is_ServerToolsChanged()) {
      agency.open.ahp.ConfluxCodec.Json _2___mcc_h4 = ((agency.open.ahp.Session.SessionAction_ServerToolsChanged)_source0)._tools;
      return true;
    } else if (_source0.is_MetaChanged()) {
      agency.open.ahp.ConfluxCodec.Json _3___mcc_h6 = ((agency.open.ahp.Session.SessionAction_MetaChanged)_source0)._meta;
      return true;
    } else if (_source0.is_IsReadChanged()) {
      boolean _4___mcc_h8 = ((agency.open.ahp.Session.SessionAction_IsReadChanged)_source0)._isRead;
      return true;
    } else if (_source0.is_IsArchivedChanged()) {
      boolean _5___mcc_h10 = ((agency.open.ahp.Session.SessionAction_IsArchivedChanged)_source0)._isArchived;
      return true;
    } else if (_source0.is_ActivityChanged()) {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _6___mcc_h12 = ((agency.open.ahp.Session.SessionAction_ActivityChanged)_source0)._activity;
      return true;
    } else if (_source0.is_ConfigChanged()) {
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _7___mcc_h14 = ((agency.open.ahp.Session.SessionAction_ConfigChanged)_source0)._config;
      boolean _8___mcc_h15 = ((agency.open.ahp.Session.SessionAction_ConfigChanged)_source0)._replace;
      return true;
    } else if (_source0.is_CustomizationsChanged()) {
      dafny.DafnySequence<? extends Cust> _9___mcc_h18 = ((agency.open.ahp.Session.SessionAction_CustomizationsChanged)_source0)._customizations;
      dafny.DafnySequence<? extends Cust> _10_cs = _9___mcc_h18;
      return (__default.WfCusts(_10_cs)) && (agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Cust>UniqueKeys(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Cust._typeDescriptor(), __default.custKey(), _10_cs));
    } else if (_source0.is_CustomizationToggled()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _11___mcc_h20 = ((agency.open.ahp.Session.SessionAction_CustomizationToggled)_source0)._id;
      boolean _12___mcc_h21 = ((agency.open.ahp.Session.SessionAction_CustomizationToggled)_source0)._enabled;
      return true;
    } else if (_source0.is_CustomizationRemoved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _13___mcc_h24 = ((agency.open.ahp.Session.SessionAction_CustomizationRemoved)_source0)._id;
      return true;
    } else if (_source0.is_DefaultChatChanged()) {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _14___mcc_h26 = ((agency.open.ahp.Session.SessionAction_DefaultChatChanged)_source0)._chat;
      return true;
    } else if (_source0.is_ChatAdded()) {
      Chat _15___mcc_h28 = ((agency.open.ahp.Session.SessionAction_ChatAdded)_source0)._summary;
      return true;
    } else if (_source0.is_ChatRemoved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _16___mcc_h30 = ((agency.open.ahp.Session.SessionAction_ChatRemoved)_source0)._resource;
      return true;
    } else if (_source0.is_ChatUpdated()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _17___mcc_h32 = ((agency.open.ahp.Session.SessionAction_ChatUpdated)_source0)._resource;
      agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _18___mcc_h33 = ((agency.open.ahp.Session.SessionAction_ChatUpdated)_source0)._status;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _19___mcc_h34 = ((agency.open.ahp.Session.SessionAction_ChatUpdated)_source0)._activity;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _20___mcc_h35 = ((agency.open.ahp.Session.SessionAction_ChatUpdated)_source0)._modifiedAt;
      return true;
    } else if (_source0.is_ChangesetsChanged()) {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _21___mcc_h40 = ((agency.open.ahp.Session.SessionAction_ChangesetsChanged)_source0)._changesets;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _22_cs = _21___mcc_h40;
      return !((_22_cs).is_Some()) || (!java.util.Objects.equals((_22_cs).dtor_value(), agency.open.ahp.ConfluxCodec.Json.create_JNull()));
    } else if (_source0.is_ActiveClientSet()) {
      Client _23___mcc_h42 = ((agency.open.ahp.Session.SessionAction_ActiveClientSet)_source0)._client;
      return true;
    } else if (_source0.is_ActiveClientRemoved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _24___mcc_h44 = ((agency.open.ahp.Session.SessionAction_ActiveClientRemoved)_source0)._clientId;
      return true;
    } else if (_source0.is_CanvasesChanged()) {
      agency.open.ahp.ConfluxCodec.Json _25___mcc_h46 = ((agency.open.ahp.Session.SessionAction_CanvasesChanged)_source0)._canvases;
      return true;
    } else if (_source0.is_OpenCanvasesChanged()) {
      agency.open.ahp.ConfluxCodec.Json _26___mcc_h48 = ((agency.open.ahp.Session.SessionAction_OpenCanvasesChanged)_source0)._openCanvases;
      return true;
    } else if (_source0.is_InputNeededSet()) {
      InputReq _27___mcc_h50 = ((agency.open.ahp.Session.SessionAction_InputNeededSet)_source0)._req;
      return true;
    } else if (_source0.is_InputNeededRemoved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _28___mcc_h52 = ((agency.open.ahp.Session.SessionAction_InputNeededRemoved)_source0)._id;
      return true;
    } else if (_source0.is_CustomizationUpdated()) {
      Cust _29___mcc_h54 = ((agency.open.ahp.Session.SessionAction_CustomizationUpdated)_source0)._customization;
      Cust _30_c = _29___mcc_h54;
      return __default.WfCust(_30_c);
    } else if (_source0.is_McpServerStateChanged()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _31___mcc_h56 = ((agency.open.ahp.Session.SessionAction_McpServerStateChanged)_source0)._id;
      agency.open.ahp.ConfluxCodec.Json _32___mcc_h57 = ((agency.open.ahp.Session.SessionAction_McpServerStateChanged)_source0)._state;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _33___mcc_h58 = ((agency.open.ahp.Session.SessionAction_McpServerStateChanged)_source0)._channel;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _34_ch = _33___mcc_h58;
      return !((_34_ch).is_Some()) || (!java.util.Objects.equals((_34_ch).dtor_value(), agency.open.ahp.ConfluxCodec.Json.create_JNull()));
    } else if (_source0.is_McpServerStartRequested()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _35___mcc_h62 = ((agency.open.ahp.Session.SessionAction_McpServerStartRequested)_source0)._id;
      return true;
    } else if (_source0.is_McpServerStopRequested()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _36___mcc_h64 = ((agency.open.ahp.Session.SessionAction_McpServerStopRequested)_source0)._id;
      return true;
    } else {
      agency.open.ahp.ConfluxCodec.Json _37___mcc_h66 = ((agency.open.ahp.Session.SessionAction_SUnknown)_source0)._raw;
      return true;
    }
  }
  public static SessionState apply1(SessionState s, SessionAction a)
  {
    return (__default.ApplyToSession(s, a, java.math.BigInteger.valueOf(9999L))).dtor__0();
  }
  public static SessionState fold(SessionState s, dafny.DafnySequence<? extends SessionAction> acts)
  {
    return agency.open.ahp.ConfluxContract.__default.<SessionState, SessionAction>Fold(SessionState._typeDescriptor(), SessionAction._typeDescriptor(), __default::apply1, s, acts);
  }
  public static SessionState S0() {
    return SessionState.create(dafny.DafnySequence.asUnicodeString("copilot"), dafny.DafnySequence.asUnicodeString("Test Session"), 1, dafny.DafnySequence.asUnicodeString("creating"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<SConfig>create_None(SConfig._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), dafny.DafnySequence.<Client> empty(Client._typeDescriptor()), dafny.DafnySequence.<Chat> empty(Chat._typeDescriptor()), dafny.DafnySequence.<Cust> empty(Cust._typeDescriptor()), dafny.DafnySequence.<InputReq> empty(InputReq._typeDescriptor()));
  }
  public static dafny.Tuple2 RunCorpus()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    java.math.BigInteger modeled = java.math.BigInteger.ZERO;
    if(true) {
      modeled = java.math.BigInteger.valueOf(36L);
      pass = java.math.BigInteger.ZERO;
      SessionState _0_s;
      _0_s = __default.S0();
      if (java.util.Objects.equals(__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed334 -> {
        SessionState _pat_let167_0 = ((SessionState)(java.lang.Object)(boxed334));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let167_0, boxed335 -> {
          SessionState _1_dt__update__tmp_h0 = ((SessionState)(java.lang.Object)(boxed335));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, SessionState>Let(dafny.DafnySequence.asUnicodeString("creating"), boxed336 -> {
            dafny.DafnySequence<? extends dafny.CodePoint> _pat_let168_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed336));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, SessionState>Let(_pat_let168_0, boxed337 -> {
              dafny.DafnySequence<? extends dafny.CodePoint> _2_dt__update_hlifecycle_h0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed337));
              return SessionState.create((_1_dt__update__tmp_h0).dtor_provider(), (_1_dt__update__tmp_h0).dtor_title(), (_1_dt__update__tmp_h0).dtor_status(), _2_dt__update_hlifecycle_h0, (_1_dt__update__tmp_h0).dtor_activity(), (_1_dt__update__tmp_h0).dtor_config(), (_1_dt__update__tmp_h0).dtor_meta(), (_1_dt__update__tmp_h0).dtor_creationError(), (_1_dt__update__tmp_h0).dtor_serverTools(), (_1_dt__update__tmp_h0).dtor_changesets(), (_1_dt__update__tmp_h0).dtor_canvases(), (_1_dt__update__tmp_h0).dtor_openCanvases(), (_1_dt__update__tmp_h0).dtor_defaultChat(), (_1_dt__update__tmp_h0).dtor_activeClients(), (_1_dt__update__tmp_h0).dtor_chats(), (_1_dt__update__tmp_h0).dtor_customizations(), (_1_dt__update__tmp_h0).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_Ready()), ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed338 -> {
        SessionState _pat_let169_0 = ((SessionState)(java.lang.Object)(boxed338));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let169_0, boxed339 -> {
          SessionState _3_dt__update__tmp_h1 = ((SessionState)(java.lang.Object)(boxed339));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, SessionState>Let(dafny.DafnySequence.asUnicodeString("ready"), boxed340 -> {
            dafny.DafnySequence<? extends dafny.CodePoint> _pat_let170_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed340));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, SessionState>Let(_pat_let170_0, boxed341 -> {
              dafny.DafnySequence<? extends dafny.CodePoint> _4_dt__update_hlifecycle_h1 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed341));
              return SessionState.create((_3_dt__update__tmp_h1).dtor_provider(), (_3_dt__update__tmp_h1).dtor_title(), (_3_dt__update__tmp_h1).dtor_status(), _4_dt__update_hlifecycle_h1, (_3_dt__update__tmp_h1).dtor_activity(), (_3_dt__update__tmp_h1).dtor_config(), (_3_dt__update__tmp_h1).dtor_meta(), (_3_dt__update__tmp_h1).dtor_creationError(), (_3_dt__update__tmp_h1).dtor_serverTools(), (_3_dt__update__tmp_h1).dtor_changesets(), (_3_dt__update__tmp_h1).dtor_canvases(), (_3_dt__update__tmp_h1).dtor_openCanvases(), (_3_dt__update__tmp_h1).dtor_defaultChat(), (_3_dt__update__tmp_h1).dtor_activeClients(), (_3_dt__update__tmp_h1).dtor_chats(), (_3_dt__update__tmp_h1).dtor_customizations(), (_3_dt__update__tmp_h1).dtor_inputNeeded());
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
      if (((__default.apply1(_0_s, SessionAction.create_CreationFailed(agency.open.ahp.ConfluxCodec.Json.create_JNull()))).dtor_lifecycle()).equals(dafny.DafnySequence.asUnicodeString("creationFailed"))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (((__default.apply1(_0_s, SessionAction.create_TitleChanged(dafny.DafnySequence.asUnicodeString("New")))).dtor_title()).equals(dafny.DafnySequence.asUnicodeString("New"))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      agency.open.ahp.ConfluxCodec.Json _5_tools;
      _5_tools = agency.open.ahp.ConfluxCodec.Json.create_JArr(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> of(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("name"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("bash"))) }))));
      if (java.util.Objects.equals((__default.apply1(_0_s, SessionAction.create_ServerToolsChanged(_5_tools))).dtor_serverTools(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _5_tools))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      agency.open.ahp.ConfluxCodec.Json _6_m;
      _6_m = agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("git"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("feature"))) }));
      if (java.util.Objects.equals((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed342 -> {
        SessionState _pat_let171_0 = ((SessionState)(java.lang.Object)(boxed342));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let171_0, boxed343 -> {
          SessionState _7_dt__update__tmp_h2 = ((SessionState)(java.lang.Object)(boxed343));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JNull()), boxed344 -> {
            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let172_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed344));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(_pat_let172_0, boxed345 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _8_dt__update_hmeta_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed345));
              return SessionState.create((_7_dt__update__tmp_h2).dtor_provider(), (_7_dt__update__tmp_h2).dtor_title(), (_7_dt__update__tmp_h2).dtor_status(), (_7_dt__update__tmp_h2).dtor_lifecycle(), (_7_dt__update__tmp_h2).dtor_activity(), (_7_dt__update__tmp_h2).dtor_config(), _8_dt__update_hmeta_h0, (_7_dt__update__tmp_h2).dtor_creationError(), (_7_dt__update__tmp_h2).dtor_serverTools(), (_7_dt__update__tmp_h2).dtor_changesets(), (_7_dt__update__tmp_h2).dtor_canvases(), (_7_dt__update__tmp_h2).dtor_openCanvases(), (_7_dt__update__tmp_h2).dtor_defaultChat(), (_7_dt__update__tmp_h2).dtor_activeClients(), (_7_dt__update__tmp_h2).dtor_chats(), (_7_dt__update__tmp_h2).dtor_customizations(), (_7_dt__update__tmp_h2).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_MetaChanged(_6_m))).dtor_meta(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _6_m))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (((__default.apply1(_0_s, SessionAction.create_IsReadChanged(true))).dtor_status()) == (33)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed346 -> {
        SessionState _pat_let173_0 = ((SessionState)(java.lang.Object)(boxed346));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let173_0, boxed347 -> {
          SessionState _9_dt__update__tmp_h3 = ((SessionState)(java.lang.Object)(boxed347));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(33, boxed348 -> {
            int _pat_let174_0 = ((int)(java.lang.Object)(boxed348));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(_pat_let174_0, boxed349 -> {
              int _10_dt__update_hstatus_h0 = ((int)(java.lang.Object)(boxed349));
              return SessionState.create((_9_dt__update__tmp_h3).dtor_provider(), (_9_dt__update__tmp_h3).dtor_title(), _10_dt__update_hstatus_h0, (_9_dt__update__tmp_h3).dtor_lifecycle(), (_9_dt__update__tmp_h3).dtor_activity(), (_9_dt__update__tmp_h3).dtor_config(), (_9_dt__update__tmp_h3).dtor_meta(), (_9_dt__update__tmp_h3).dtor_creationError(), (_9_dt__update__tmp_h3).dtor_serverTools(), (_9_dt__update__tmp_h3).dtor_changesets(), (_9_dt__update__tmp_h3).dtor_canvases(), (_9_dt__update__tmp_h3).dtor_openCanvases(), (_9_dt__update__tmp_h3).dtor_defaultChat(), (_9_dt__update__tmp_h3).dtor_activeClients(), (_9_dt__update__tmp_h3).dtor_chats(), (_9_dt__update__tmp_h3).dtor_customizations(), (_9_dt__update__tmp_h3).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_IsReadChanged(false))).dtor_status()) == (1)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (((__default.apply1(_0_s, SessionAction.create_IsArchivedChanged(true))).dtor_status()) == (65)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed350 -> {
        SessionState _pat_let175_0 = ((SessionState)(java.lang.Object)(boxed350));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let175_0, boxed351 -> {
          SessionState _11_dt__update__tmp_h4 = ((SessionState)(java.lang.Object)(boxed351));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(65, boxed352 -> {
            int _pat_let176_0 = ((int)(java.lang.Object)(boxed352));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(_pat_let176_0, boxed353 -> {
              int _12_dt__update_hstatus_h1 = ((int)(java.lang.Object)(boxed353));
              return SessionState.create((_11_dt__update__tmp_h4).dtor_provider(), (_11_dt__update__tmp_h4).dtor_title(), _12_dt__update_hstatus_h1, (_11_dt__update__tmp_h4).dtor_lifecycle(), (_11_dt__update__tmp_h4).dtor_activity(), (_11_dt__update__tmp_h4).dtor_config(), (_11_dt__update__tmp_h4).dtor_meta(), (_11_dt__update__tmp_h4).dtor_creationError(), (_11_dt__update__tmp_h4).dtor_serverTools(), (_11_dt__update__tmp_h4).dtor_changesets(), (_11_dt__update__tmp_h4).dtor_canvases(), (_11_dt__update__tmp_h4).dtor_openCanvases(), (_11_dt__update__tmp_h4).dtor_defaultChat(), (_11_dt__update__tmp_h4).dtor_activeClients(), (_11_dt__update__tmp_h4).dtor_chats(), (_11_dt__update__tmp_h4).dtor_customizations(), (_11_dt__update__tmp_h4).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_IsArchivedChanged(false))).dtor_status()) == (1)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_0_s, SessionAction.create_ActivityChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Running")))), ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed354 -> {
        SessionState _pat_let177_0 = ((SessionState)(java.lang.Object)(boxed354));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let177_0, boxed355 -> {
          SessionState _13_dt__update__tmp_h5 = ((SessionState)(java.lang.Object)(boxed355));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Running")), boxed356 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let178_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed356));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(_pat_let178_0, boxed357 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _14_dt__update_hactivity_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed357));
              return SessionState.create((_13_dt__update__tmp_h5).dtor_provider(), (_13_dt__update__tmp_h5).dtor_title(), (_13_dt__update__tmp_h5).dtor_status(), (_13_dt__update__tmp_h5).dtor_lifecycle(), _14_dt__update_hactivity_h0, (_13_dt__update__tmp_h5).dtor_config(), (_13_dt__update__tmp_h5).dtor_meta(), (_13_dt__update__tmp_h5).dtor_creationError(), (_13_dt__update__tmp_h5).dtor_serverTools(), (_13_dt__update__tmp_h5).dtor_changesets(), (_13_dt__update__tmp_h5).dtor_canvases(), (_13_dt__update__tmp_h5).dtor_openCanvases(), (_13_dt__update__tmp_h5).dtor_defaultChat(), (_13_dt__update__tmp_h5).dtor_activeClients(), (_13_dt__update__tmp_h5).dtor_chats(), (_13_dt__update__tmp_h5).dtor_customizations(), (_13_dt__update__tmp_h5).dtor_inputNeeded());
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
      if (java.util.Objects.equals(__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed358 -> {
        SessionState _pat_let179_0 = ((SessionState)(java.lang.Object)(boxed358));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let179_0, boxed359 -> {
          SessionState _15_dt__update__tmp_h6 = ((SessionState)(java.lang.Object)(boxed359));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("x")), boxed360 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let180_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed360));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(_pat_let180_0, boxed361 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _16_dt__update_hactivity_h1 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed361));
              return SessionState.create((_15_dt__update__tmp_h6).dtor_provider(), (_15_dt__update__tmp_h6).dtor_title(), (_15_dt__update__tmp_h6).dtor_status(), (_15_dt__update__tmp_h6).dtor_lifecycle(), _16_dt__update_hactivity_h1, (_15_dt__update__tmp_h6).dtor_config(), (_15_dt__update__tmp_h6).dtor_meta(), (_15_dt__update__tmp_h6).dtor_creationError(), (_15_dt__update__tmp_h6).dtor_serverTools(), (_15_dt__update__tmp_h6).dtor_changesets(), (_15_dt__update__tmp_h6).dtor_canvases(), (_15_dt__update__tmp_h6).dtor_openCanvases(), (_15_dt__update__tmp_h6).dtor_defaultChat(), (_15_dt__update__tmp_h6).dtor_activeClients(), (_15_dt__update__tmp_h6).dtor_chats(), (_15_dt__update__tmp_h6).dtor_customizations(), (_15_dt__update__tmp_h6).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_ActivityChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)))), _0_s)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      SConfig _17_cfg;
      _17_cfg = SConfig.create(agency.open.ahp.ConfluxCodec.Json.create_JNull(), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("target"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("worktree"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("baseBranch"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("main"))) }));
      SConfig _18_cfgExp;
      _18_cfgExp = SConfig.create(agency.open.ahp.ConfluxCodec.Json.create_JNull(), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("target"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("worktree"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("baseBranch"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("develop"))) }));
      SConfig _pat_let_tv0 = _17_cfg;
      if (java.util.Objects.equals((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed362 -> {
        SessionState _pat_let181_0 = ((SessionState)(java.lang.Object)(boxed362));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let181_0, boxed363 -> {
          SessionState _19_dt__update__tmp_h7 = ((SessionState)(java.lang.Object)(boxed363));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<SConfig>, SessionState>Let(agency.open.ahp.AhpSkeleton.Option.<SConfig>create_Some(SConfig._typeDescriptor(), _pat_let_tv0), boxed364 -> {
            agency.open.ahp.AhpSkeleton.Option<SConfig> _pat_let182_0 = ((agency.open.ahp.AhpSkeleton.Option<SConfig>)(java.lang.Object)(boxed364));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<SConfig>, SessionState>Let(_pat_let182_0, boxed365 -> {
              agency.open.ahp.AhpSkeleton.Option<SConfig> _20_dt__update_hconfig_h0 = ((agency.open.ahp.AhpSkeleton.Option<SConfig>)(java.lang.Object)(boxed365));
              return SessionState.create((_19_dt__update__tmp_h7).dtor_provider(), (_19_dt__update__tmp_h7).dtor_title(), (_19_dt__update__tmp_h7).dtor_status(), (_19_dt__update__tmp_h7).dtor_lifecycle(), (_19_dt__update__tmp_h7).dtor_activity(), _20_dt__update_hconfig_h0, (_19_dt__update__tmp_h7).dtor_meta(), (_19_dt__update__tmp_h7).dtor_creationError(), (_19_dt__update__tmp_h7).dtor_serverTools(), (_19_dt__update__tmp_h7).dtor_changesets(), (_19_dt__update__tmp_h7).dtor_canvases(), (_19_dt__update__tmp_h7).dtor_openCanvases(), (_19_dt__update__tmp_h7).dtor_defaultChat(), (_19_dt__update__tmp_h7).dtor_activeClients(), (_19_dt__update__tmp_h7).dtor_chats(), (_19_dt__update__tmp_h7).dtor_customizations(), (_19_dt__update__tmp_h7).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_ConfigChanged(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("baseBranch"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("develop"))) }), false))).dtor_config(), agency.open.ahp.AhpSkeleton.Option.<SConfig>create_Some(SConfig._typeDescriptor(), _18_cfgExp))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      SConfig _21_cfgRepl;
      _21_cfgRepl = SConfig.create(agency.open.ahp.ConfluxCodec.Json.create_JNull(), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("baseBranch"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("develop"))) }));
      SConfig _pat_let_tv1 = _17_cfg;
      if (java.util.Objects.equals((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed366 -> {
        SessionState _pat_let183_0 = ((SessionState)(java.lang.Object)(boxed366));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let183_0, boxed367 -> {
          SessionState _22_dt__update__tmp_h8 = ((SessionState)(java.lang.Object)(boxed367));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<SConfig>, SessionState>Let(agency.open.ahp.AhpSkeleton.Option.<SConfig>create_Some(SConfig._typeDescriptor(), _pat_let_tv1), boxed368 -> {
            agency.open.ahp.AhpSkeleton.Option<SConfig> _pat_let184_0 = ((agency.open.ahp.AhpSkeleton.Option<SConfig>)(java.lang.Object)(boxed368));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<SConfig>, SessionState>Let(_pat_let184_0, boxed369 -> {
              agency.open.ahp.AhpSkeleton.Option<SConfig> _23_dt__update_hconfig_h1 = ((agency.open.ahp.AhpSkeleton.Option<SConfig>)(java.lang.Object)(boxed369));
              return SessionState.create((_22_dt__update__tmp_h8).dtor_provider(), (_22_dt__update__tmp_h8).dtor_title(), (_22_dt__update__tmp_h8).dtor_status(), (_22_dt__update__tmp_h8).dtor_lifecycle(), (_22_dt__update__tmp_h8).dtor_activity(), _23_dt__update_hconfig_h1, (_22_dt__update__tmp_h8).dtor_meta(), (_22_dt__update__tmp_h8).dtor_creationError(), (_22_dt__update__tmp_h8).dtor_serverTools(), (_22_dt__update__tmp_h8).dtor_changesets(), (_22_dt__update__tmp_h8).dtor_canvases(), (_22_dt__update__tmp_h8).dtor_openCanvases(), (_22_dt__update__tmp_h8).dtor_defaultChat(), (_22_dt__update__tmp_h8).dtor_activeClients(), (_22_dt__update__tmp_h8).dtor_chats(), (_22_dt__update__tmp_h8).dtor_customizations(), (_22_dt__update__tmp_h8).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_ConfigChanged(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("baseBranch"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("develop"))) }), true))).dtor_config(), agency.open.ahp.AhpSkeleton.Option.<SConfig>create_Some(SConfig._typeDescriptor(), _21_cfgRepl))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_0_s, SessionAction.create_ConfigChanged(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("x"), agency.open.ahp.ConfluxCodec.Json.create_JNull()) }), false)), _0_s)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Cust _24_pa;
      _24_pa = Cust.create(dafny.DafnySequence.asUnicodeString("plugin-a"), dafny.DafnySequence.asUnicodeString("plugin"), dafny.DafnySequence.asUnicodeString("https://plugins.example/a"), dafny.DafnySequence.asUnicodeString("Plugin A"), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      Cust _25_pb;
      _25_pb = Cust.create(dafny.DafnySequence.asUnicodeString("plugin-b"), dafny.DafnySequence.asUnicodeString("plugin"), dafny.DafnySequence.asUnicodeString("https://plugins.example/b"), dafny.DafnySequence.asUnicodeString("Plugin B"), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      if (((__default.apply1(_0_s, SessionAction.create_CustomizationsChanged(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), _24_pa, _25_pb)))).dtor_customizations()).equals(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), _24_pa, _25_pb))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Cust _pat_let_tv2 = _24_pa;
      Cust _pat_let_tv3 = _25_pb;
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed370 -> {
        SessionState _pat_let185_0 = ((SessionState)(java.lang.Object)(boxed370));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let185_0, boxed371 -> {
          SessionState _26_dt__update__tmp_h9 = ((SessionState)(java.lang.Object)(boxed371));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), _pat_let_tv2, _pat_let_tv3), boxed372 -> {
            dafny.DafnySequence<? extends Cust> _pat_let186_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed372));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let186_0, boxed373 -> {
              dafny.DafnySequence<? extends Cust> _27_dt__update_hcustomizations_h0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed373));
              return SessionState.create((_26_dt__update__tmp_h9).dtor_provider(), (_26_dt__update__tmp_h9).dtor_title(), (_26_dt__update__tmp_h9).dtor_status(), (_26_dt__update__tmp_h9).dtor_lifecycle(), (_26_dt__update__tmp_h9).dtor_activity(), (_26_dt__update__tmp_h9).dtor_config(), (_26_dt__update__tmp_h9).dtor_meta(), (_26_dt__update__tmp_h9).dtor_creationError(), (_26_dt__update__tmp_h9).dtor_serverTools(), (_26_dt__update__tmp_h9).dtor_changesets(), (_26_dt__update__tmp_h9).dtor_canvases(), (_26_dt__update__tmp_h9).dtor_openCanvases(), (_26_dt__update__tmp_h9).dtor_defaultChat(), (_26_dt__update__tmp_h9).dtor_activeClients(), (_26_dt__update__tmp_h9).dtor_chats(), _27_dt__update_hcustomizations_h0, (_26_dt__update__tmp_h9).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_CustomizationToggled(dafny.DafnySequence.asUnicodeString("plugin-a"), false))).dtor_customizations()).equals(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_24_pa, boxed374 -> {
        Cust _pat_let187_0 = ((Cust)(java.lang.Object)(boxed374));
        return ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_pat_let187_0, boxed375 -> {
          Cust _28_dt__update__tmp_h10 = ((Cust)(java.lang.Object)(boxed375));
          return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, Cust>Let(agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, false), boxed376 -> {
            agency.open.ahp.AhpSkeleton.Option<Boolean> _pat_let188_0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed376));
            return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, Cust>Let(_pat_let188_0, boxed377 -> {
              agency.open.ahp.AhpSkeleton.Option<Boolean> _29_dt__update_henabled_h0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed377));
              return Cust.create((_28_dt__update__tmp_h10).dtor_id(), (_28_dt__update__tmp_h10).dtor_kind(), (_28_dt__update__tmp_h10).dtor_uri(), (_28_dt__update__tmp_h10).dtor_name(), _29_dt__update_henabled_h0, (_28_dt__update__tmp_h10).dtor_state(), (_28_dt__update__tmp_h10).dtor_channel());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), _25_pb))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Cust _pat_let_tv4 = _24_pa;
      Cust _pat_let_tv5 = _25_pb;
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed378 -> {
        SessionState _pat_let189_0 = ((SessionState)(java.lang.Object)(boxed378));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let189_0, boxed379 -> {
          SessionState _30_dt__update__tmp_h11 = ((SessionState)(java.lang.Object)(boxed379));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), _pat_let_tv4, _pat_let_tv5), boxed380 -> {
            dafny.DafnySequence<? extends Cust> _pat_let190_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed380));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let190_0, boxed381 -> {
              dafny.DafnySequence<? extends Cust> _31_dt__update_hcustomizations_h1 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed381));
              return SessionState.create((_30_dt__update__tmp_h11).dtor_provider(), (_30_dt__update__tmp_h11).dtor_title(), (_30_dt__update__tmp_h11).dtor_status(), (_30_dt__update__tmp_h11).dtor_lifecycle(), (_30_dt__update__tmp_h11).dtor_activity(), (_30_dt__update__tmp_h11).dtor_config(), (_30_dt__update__tmp_h11).dtor_meta(), (_30_dt__update__tmp_h11).dtor_creationError(), (_30_dt__update__tmp_h11).dtor_serverTools(), (_30_dt__update__tmp_h11).dtor_changesets(), (_30_dt__update__tmp_h11).dtor_canvases(), (_30_dt__update__tmp_h11).dtor_openCanvases(), (_30_dt__update__tmp_h11).dtor_defaultChat(), (_30_dt__update__tmp_h11).dtor_activeClients(), (_30_dt__update__tmp_h11).dtor_chats(), _31_dt__update_hcustomizations_h1, (_30_dt__update__tmp_h11).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_CustomizationRemoved(dafny.DafnySequence.asUnicodeString("plugin-a")))).dtor_customizations()).equals(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), _25_pb))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals((__default.apply1(_0_s, SessionAction.create_DefaultChatChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("c1"))))).dtor_defaultChat(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("c1")))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed382 -> {
        SessionState _pat_let191_0 = ((SessionState)(java.lang.Object)(boxed382));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let191_0, boxed383 -> {
          SessionState _32_dt__update__tmp_h12 = ((SessionState)(java.lang.Object)(boxed383));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("ahp-chat:/old")), boxed384 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let192_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed384));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(_pat_let192_0, boxed385 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _33_dt__update_hdefaultChat_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed385));
              return SessionState.create((_32_dt__update__tmp_h12).dtor_provider(), (_32_dt__update__tmp_h12).dtor_title(), (_32_dt__update__tmp_h12).dtor_status(), (_32_dt__update__tmp_h12).dtor_lifecycle(), (_32_dt__update__tmp_h12).dtor_activity(), (_32_dt__update__tmp_h12).dtor_config(), (_32_dt__update__tmp_h12).dtor_meta(), (_32_dt__update__tmp_h12).dtor_creationError(), (_32_dt__update__tmp_h12).dtor_serverTools(), (_32_dt__update__tmp_h12).dtor_changesets(), (_32_dt__update__tmp_h12).dtor_canvases(), (_32_dt__update__tmp_h12).dtor_openCanvases(), _33_dt__update_hdefaultChat_h0, (_32_dt__update__tmp_h12).dtor_activeClients(), (_32_dt__update__tmp_h12).dtor_chats(), (_32_dt__update__tmp_h12).dtor_customizations(), (_32_dt__update__tmp_h12).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_DefaultChatChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))))).dtor_defaultChat(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Chat _34_c1;
      _34_c1 = Chat.create(dafny.DafnySequence.asUnicodeString("ahp-chat:/c1"), dafny.DafnySequence.asUnicodeString("Chat 1"), java.math.BigInteger.ONE, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), dafny.DafnySequence.asUnicodeString("t1"), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("user"))) })));
      if (((__default.apply1(_0_s, SessionAction.create_ChatAdded(_34_c1))).dtor_chats()).equals(dafny.DafnySequence.<Chat> of(Chat._typeDescriptor(), _34_c1))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Chat _35_c1r;
      Chat _36_dt__update__tmp_h13 = _34_c1;
      dafny.DafnySequence<? extends dafny.CodePoint> _37_dt__update_hmodifiedAt_h0 = dafny.DafnySequence.asUnicodeString("t2");
      java.math.BigInteger _38_dt__update_hstatus_h2 = java.math.BigInteger.valueOf(8L);
      dafny.DafnySequence<? extends dafny.CodePoint> _39_dt__update_htitle_h0 = dafny.DafnySequence.asUnicodeString("Chat 1 (renamed)");
      _35_c1r = Chat.create((_36_dt__update__tmp_h13).dtor_resource(), _39_dt__update_htitle_h0, _38_dt__update_hstatus_h2, (_36_dt__update__tmp_h13).dtor_activity(), _37_dt__update_hmodifiedAt_h0, (_36_dt__update__tmp_h13).dtor_origin());
      Chat _pat_let_tv6 = _34_c1;
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed386 -> {
        SessionState _pat_let193_0 = ((SessionState)(java.lang.Object)(boxed386));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let193_0, boxed387 -> {
          SessionState _40_dt__update__tmp_h14 = ((SessionState)(java.lang.Object)(boxed387));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(dafny.DafnySequence.<Chat> of(Chat._typeDescriptor(), _pat_let_tv6), boxed388 -> {
            dafny.DafnySequence<? extends Chat> _pat_let194_0 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed388));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(_pat_let194_0, boxed389 -> {
              dafny.DafnySequence<? extends Chat> _41_dt__update_hchats_h0 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed389));
              return SessionState.create((_40_dt__update__tmp_h14).dtor_provider(), (_40_dt__update__tmp_h14).dtor_title(), (_40_dt__update__tmp_h14).dtor_status(), (_40_dt__update__tmp_h14).dtor_lifecycle(), (_40_dt__update__tmp_h14).dtor_activity(), (_40_dt__update__tmp_h14).dtor_config(), (_40_dt__update__tmp_h14).dtor_meta(), (_40_dt__update__tmp_h14).dtor_creationError(), (_40_dt__update__tmp_h14).dtor_serverTools(), (_40_dt__update__tmp_h14).dtor_changesets(), (_40_dt__update__tmp_h14).dtor_canvases(), (_40_dt__update__tmp_h14).dtor_openCanvases(), (_40_dt__update__tmp_h14).dtor_defaultChat(), (_40_dt__update__tmp_h14).dtor_activeClients(), _41_dt__update_hchats_h0, (_40_dt__update__tmp_h14).dtor_customizations(), (_40_dt__update__tmp_h14).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_ChatAdded(_35_c1r))).dtor_chats()).equals(dafny.DafnySequence.<Chat> of(Chat._typeDescriptor(), _35_c1r))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Chat _42_c2;
      _42_c2 = Chat.create(dafny.DafnySequence.asUnicodeString("ahp-chat:/c2"), dafny.DafnySequence.asUnicodeString("Chat 2"), java.math.BigInteger.valueOf(8L), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Thinking")), dafny.DafnySequence.asUnicodeString("t2"), agency.open.ahp.ConfluxCodec.Json.create_JNull());
      Chat _pat_let_tv7 = _34_c1;
      Chat _pat_let_tv8 = _42_c2;
      Chat _pat_let_tv9 = _42_c2;
      if (java.util.Objects.equals(__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed390 -> {
        SessionState _pat_let195_0 = ((SessionState)(java.lang.Object)(boxed390));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let195_0, boxed391 -> {
          SessionState _43_dt__update__tmp_h15 = ((SessionState)(java.lang.Object)(boxed391));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("ahp-chat:/c1")), boxed392 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let196_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed392));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(_pat_let196_0, boxed393 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _44_dt__update_hdefaultChat_h1 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed393));
              return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(dafny.DafnySequence.<Chat> of(Chat._typeDescriptor(), _pat_let_tv7, _pat_let_tv8), boxed394 -> {
                dafny.DafnySequence<? extends Chat> _pat_let197_0 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed394));
                return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(_pat_let197_0, boxed395 -> {
                  dafny.DafnySequence<? extends Chat> _45_dt__update_hchats_h1 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed395));
                  return SessionState.create((_43_dt__update__tmp_h15).dtor_provider(), (_43_dt__update__tmp_h15).dtor_title(), (_43_dt__update__tmp_h15).dtor_status(), (_43_dt__update__tmp_h15).dtor_lifecycle(), (_43_dt__update__tmp_h15).dtor_activity(), (_43_dt__update__tmp_h15).dtor_config(), (_43_dt__update__tmp_h15).dtor_meta(), (_43_dt__update__tmp_h15).dtor_creationError(), (_43_dt__update__tmp_h15).dtor_serverTools(), (_43_dt__update__tmp_h15).dtor_changesets(), (_43_dt__update__tmp_h15).dtor_canvases(), (_43_dt__update__tmp_h15).dtor_openCanvases(), _44_dt__update_hdefaultChat_h1, (_43_dt__update__tmp_h15).dtor_activeClients(), _45_dt__update_hchats_h1, (_43_dt__update__tmp_h15).dtor_customizations(), (_43_dt__update__tmp_h15).dtor_inputNeeded());
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
      ))), SessionAction.create_ChatRemoved(dafny.DafnySequence.asUnicodeString("ahp-chat:/c1"))), ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed396 -> {
        SessionState _pat_let198_0 = ((SessionState)(java.lang.Object)(boxed396));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let198_0, boxed397 -> {
          SessionState _46_dt__update__tmp_h16 = ((SessionState)(java.lang.Object)(boxed397));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), boxed398 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let199_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed398));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, SessionState>Let(_pat_let199_0, boxed399 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _47_dt__update_hdefaultChat_h2 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed399));
              return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(dafny.DafnySequence.<Chat> of(Chat._typeDescriptor(), _pat_let_tv9), boxed400 -> {
                dafny.DafnySequence<? extends Chat> _pat_let200_0 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed400));
                return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(_pat_let200_0, boxed401 -> {
                  dafny.DafnySequence<? extends Chat> _48_dt__update_hchats_h2 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed401));
                  return SessionState.create((_46_dt__update__tmp_h16).dtor_provider(), (_46_dt__update__tmp_h16).dtor_title(), (_46_dt__update__tmp_h16).dtor_status(), (_46_dt__update__tmp_h16).dtor_lifecycle(), (_46_dt__update__tmp_h16).dtor_activity(), (_46_dt__update__tmp_h16).dtor_config(), (_46_dt__update__tmp_h16).dtor_meta(), (_46_dt__update__tmp_h16).dtor_creationError(), (_46_dt__update__tmp_h16).dtor_serverTools(), (_46_dt__update__tmp_h16).dtor_changesets(), (_46_dt__update__tmp_h16).dtor_canvases(), (_46_dt__update__tmp_h16).dtor_openCanvases(), _47_dt__update_hdefaultChat_h2, (_46_dt__update__tmp_h16).dtor_activeClients(), _48_dt__update_hchats_h2, (_46_dt__update__tmp_h16).dtor_customizations(), (_46_dt__update__tmp_h16).dtor_inputNeeded());
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
      Chat _49_c1u;
      Chat _50_dt__update__tmp_h17 = _34_c1;
      dafny.DafnySequence<? extends dafny.CodePoint> _51_dt__update_hmodifiedAt_h1 = dafny.DafnySequence.asUnicodeString("t2");
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _52_dt__update_hactivity_h2 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Waiting for approval"));
      java.math.BigInteger _53_dt__update_hstatus_h3 = java.math.BigInteger.valueOf(24L);
      _49_c1u = Chat.create((_50_dt__update__tmp_h17).dtor_resource(), (_50_dt__update__tmp_h17).dtor_title(), _53_dt__update_hstatus_h3, _52_dt__update_hactivity_h2, _51_dt__update_hmodifiedAt_h1, (_50_dt__update__tmp_h17).dtor_origin());
      Chat _pat_let_tv10 = _34_c1;
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed402 -> {
        SessionState _pat_let201_0 = ((SessionState)(java.lang.Object)(boxed402));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let201_0, boxed403 -> {
          SessionState _54_dt__update__tmp_h18 = ((SessionState)(java.lang.Object)(boxed403));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(dafny.DafnySequence.<Chat> of(Chat._typeDescriptor(), _pat_let_tv10), boxed404 -> {
            dafny.DafnySequence<? extends Chat> _pat_let202_0 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed404));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Chat>, SessionState>Let(_pat_let202_0, boxed405 -> {
              dafny.DafnySequence<? extends Chat> _55_dt__update_hchats_h3 = ((dafny.DafnySequence<? extends Chat>)(java.lang.Object)(boxed405));
              return SessionState.create((_54_dt__update__tmp_h18).dtor_provider(), (_54_dt__update__tmp_h18).dtor_title(), (_54_dt__update__tmp_h18).dtor_status(), (_54_dt__update__tmp_h18).dtor_lifecycle(), (_54_dt__update__tmp_h18).dtor_activity(), (_54_dt__update__tmp_h18).dtor_config(), (_54_dt__update__tmp_h18).dtor_meta(), (_54_dt__update__tmp_h18).dtor_creationError(), (_54_dt__update__tmp_h18).dtor_serverTools(), (_54_dt__update__tmp_h18).dtor_changesets(), (_54_dt__update__tmp_h18).dtor_canvases(), (_54_dt__update__tmp_h18).dtor_openCanvases(), (_54_dt__update__tmp_h18).dtor_defaultChat(), (_54_dt__update__tmp_h18).dtor_activeClients(), _55_dt__update_hchats_h3, (_54_dt__update__tmp_h18).dtor_customizations(), (_54_dt__update__tmp_h18).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_ChatUpdated(dafny.DafnySequence.asUnicodeString("ahp-chat:/c1"), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.valueOf(24L)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Waiting for approval")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("t2"))))).dtor_chats()).equals(dafny.DafnySequence.<Chat> of(Chat._typeDescriptor(), _49_c1u))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals((__default.apply1(_0_s, SessionAction.create_ChangesetsChanged(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JArr(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> empty(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))))).dtor_changesets(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JArr(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> empty(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed406 -> {
        SessionState _pat_let203_0 = ((SessionState)(java.lang.Object)(boxed406));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let203_0, boxed407 -> {
          SessionState _56_dt__update__tmp_h19 = ((SessionState)(java.lang.Object)(boxed407));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JArr(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> empty(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), boxed408 -> {
            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let204_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed408));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, SessionState>Let(_pat_let204_0, boxed409 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _57_dt__update_hchangesets_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed409));
              return SessionState.create((_56_dt__update__tmp_h19).dtor_provider(), (_56_dt__update__tmp_h19).dtor_title(), (_56_dt__update__tmp_h19).dtor_status(), (_56_dt__update__tmp_h19).dtor_lifecycle(), (_56_dt__update__tmp_h19).dtor_activity(), (_56_dt__update__tmp_h19).dtor_config(), (_56_dt__update__tmp_h19).dtor_meta(), (_56_dt__update__tmp_h19).dtor_creationError(), (_56_dt__update__tmp_h19).dtor_serverTools(), _57_dt__update_hchangesets_h0, (_56_dt__update__tmp_h19).dtor_canvases(), (_56_dt__update__tmp_h19).dtor_openCanvases(), (_56_dt__update__tmp_h19).dtor_defaultChat(), (_56_dt__update__tmp_h19).dtor_activeClients(), (_56_dt__update__tmp_h19).dtor_chats(), (_56_dt__update__tmp_h19).dtor_customizations(), (_56_dt__update__tmp_h19).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_ChangesetsChanged(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))).dtor_changesets(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      agency.open.ahp.ConfluxCodec.Json _58_canvasDecls;
      _58_canvasDecls = agency.open.ahp.ConfluxCodec.Json.create_JArr(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> of(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("providerId"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("acme.editors"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("canvasId"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("markdown"))) }))));
      if (java.util.Objects.equals((__default.apply1(_0_s, SessionAction.create_CanvasesChanged(_58_canvasDecls))).dtor_canvases(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _58_canvasDecls))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      agency.open.ahp.ConfluxCodec.Json _59_openRefs;
      _59_openRefs = agency.open.ahp.ConfluxCodec.Json.create_JArr(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> of(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("channel"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("ahp-canvas:/canvas-1"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("canvasId"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("markdown"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("availability"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("ready"))) }))));
      if (java.util.Objects.equals((__default.apply1(_0_s, SessionAction.create_OpenCanvasesChanged(_59_openRefs))).dtor_openCanvases(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _59_openRefs))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Client _60_cl;
      _60_cl = Client.create(dafny.DafnySequence.asUnicodeString("vscode-1"), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("displayName"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("VS Code"))) })));
      if (((__default.apply1(_0_s, SessionAction.create_ActiveClientSet(_60_cl))).dtor_activeClients()).equals(dafny.DafnySequence.<Client> of(Client._typeDescriptor(), _60_cl))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Client _pat_let_tv11 = _60_cl;
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed410 -> {
        SessionState _pat_let205_0 = ((SessionState)(java.lang.Object)(boxed410));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let205_0, boxed411 -> {
          SessionState _61_dt__update__tmp_h20 = ((SessionState)(java.lang.Object)(boxed411));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Client>, SessionState>Let(dafny.DafnySequence.<Client> of(Client._typeDescriptor(), _pat_let_tv11, Client.create(dafny.DafnySequence.asUnicodeString("cli-1"), agency.open.ahp.ConfluxCodec.Json.create_JNull())), boxed412 -> {
            dafny.DafnySequence<? extends Client> _pat_let206_0 = ((dafny.DafnySequence<? extends Client>)(java.lang.Object)(boxed412));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Client>, SessionState>Let(_pat_let206_0, boxed413 -> {
              dafny.DafnySequence<? extends Client> _62_dt__update_hactiveClients_h0 = ((dafny.DafnySequence<? extends Client>)(java.lang.Object)(boxed413));
              return SessionState.create((_61_dt__update__tmp_h20).dtor_provider(), (_61_dt__update__tmp_h20).dtor_title(), (_61_dt__update__tmp_h20).dtor_status(), (_61_dt__update__tmp_h20).dtor_lifecycle(), (_61_dt__update__tmp_h20).dtor_activity(), (_61_dt__update__tmp_h20).dtor_config(), (_61_dt__update__tmp_h20).dtor_meta(), (_61_dt__update__tmp_h20).dtor_creationError(), (_61_dt__update__tmp_h20).dtor_serverTools(), (_61_dt__update__tmp_h20).dtor_changesets(), (_61_dt__update__tmp_h20).dtor_canvases(), (_61_dt__update__tmp_h20).dtor_openCanvases(), (_61_dt__update__tmp_h20).dtor_defaultChat(), _62_dt__update_hactiveClients_h0, (_61_dt__update__tmp_h20).dtor_chats(), (_61_dt__update__tmp_h20).dtor_customizations(), (_61_dt__update__tmp_h20).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_ActiveClientRemoved(dafny.DafnySequence.asUnicodeString("vscode-1")))).dtor_activeClients()).equals(dafny.DafnySequence.<Client> of(Client._typeDescriptor(), Client.create(dafny.DafnySequence.asUnicodeString("cli-1"), agency.open.ahp.ConfluxCodec.Json.create_JNull())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      InputReq _63_ir;
      _63_ir = InputReq.create(dafny.DafnySequence.asUnicodeString("req-1"), agency.open.ahp.ConfluxCodec.Json.create_JNull());
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed414 -> {
        SessionState _pat_let207_0 = ((SessionState)(java.lang.Object)(boxed414));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let207_0, boxed415 -> {
          SessionState _64_dt__update__tmp_h21 = ((SessionState)(java.lang.Object)(boxed415));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(8, boxed416 -> {
            int _pat_let208_0 = ((int)(java.lang.Object)(boxed416));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(_pat_let208_0, boxed417 -> {
              int _65_dt__update_hstatus_h4 = ((int)(java.lang.Object)(boxed417));
              return SessionState.create((_64_dt__update__tmp_h21).dtor_provider(), (_64_dt__update__tmp_h21).dtor_title(), _65_dt__update_hstatus_h4, (_64_dt__update__tmp_h21).dtor_lifecycle(), (_64_dt__update__tmp_h21).dtor_activity(), (_64_dt__update__tmp_h21).dtor_config(), (_64_dt__update__tmp_h21).dtor_meta(), (_64_dt__update__tmp_h21).dtor_creationError(), (_64_dt__update__tmp_h21).dtor_serverTools(), (_64_dt__update__tmp_h21).dtor_changesets(), (_64_dt__update__tmp_h21).dtor_canvases(), (_64_dt__update__tmp_h21).dtor_openCanvases(), (_64_dt__update__tmp_h21).dtor_defaultChat(), (_64_dt__update__tmp_h21).dtor_activeClients(), (_64_dt__update__tmp_h21).dtor_chats(), (_64_dt__update__tmp_h21).dtor_customizations(), (_64_dt__update__tmp_h21).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_InputNeededSet(_63_ir))).dtor_status()) == (24)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      InputReq _66_i1;
      _66_i1 = InputReq.create(dafny.DafnySequence.asUnicodeString("a"), agency.open.ahp.ConfluxCodec.Json.create_JNull());
      InputReq _67_i2;
      _67_i2 = InputReq.create(dafny.DafnySequence.asUnicodeString("b"), agency.open.ahp.ConfluxCodec.Json.create_JNull());
      InputReq _pat_let_tv12 = _66_i1;
      InputReq _pat_let_tv13 = _67_i2;
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed418 -> {
        SessionState _pat_let209_0 = ((SessionState)(java.lang.Object)(boxed418));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let209_0, boxed419 -> {
          SessionState _68_dt__update__tmp_h22 = ((SessionState)(java.lang.Object)(boxed419));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends InputReq>, SessionState>Let(dafny.DafnySequence.<InputReq> of(InputReq._typeDescriptor(), _pat_let_tv12, _pat_let_tv13), boxed420 -> {
            dafny.DafnySequence<? extends InputReq> _pat_let210_0 = ((dafny.DafnySequence<? extends InputReq>)(java.lang.Object)(boxed420));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends InputReq>, SessionState>Let(_pat_let210_0, boxed421 -> {
              dafny.DafnySequence<? extends InputReq> _69_dt__update_hinputNeeded_h0 = ((dafny.DafnySequence<? extends InputReq>)(java.lang.Object)(boxed421));
              return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(24, boxed422 -> {
                int _pat_let211_0 = ((int)(java.lang.Object)(boxed422));
                return ((SessionState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, SessionState>Let(_pat_let211_0, boxed423 -> {
                  int _70_dt__update_hstatus_h5 = ((int)(java.lang.Object)(boxed423));
                  return SessionState.create((_68_dt__update__tmp_h22).dtor_provider(), (_68_dt__update__tmp_h22).dtor_title(), _70_dt__update_hstatus_h5, (_68_dt__update__tmp_h22).dtor_lifecycle(), (_68_dt__update__tmp_h22).dtor_activity(), (_68_dt__update__tmp_h22).dtor_config(), (_68_dt__update__tmp_h22).dtor_meta(), (_68_dt__update__tmp_h22).dtor_creationError(), (_68_dt__update__tmp_h22).dtor_serverTools(), (_68_dt__update__tmp_h22).dtor_changesets(), (_68_dt__update__tmp_h22).dtor_canvases(), (_68_dt__update__tmp_h22).dtor_openCanvases(), (_68_dt__update__tmp_h22).dtor_defaultChat(), (_68_dt__update__tmp_h22).dtor_activeClients(), (_68_dt__update__tmp_h22).dtor_chats(), (_68_dt__update__tmp_h22).dtor_customizations(), _69_dt__update_hinputNeeded_h0);
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
      ))), SessionAction.create_InputNeededRemoved(dafny.DafnySequence.asUnicodeString("a")))).dtor_status()) == (24)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Cust _71_mc;
      _71_mc = Cust.create(dafny.DafnySequence.asUnicodeString("mcp-1"), dafny.DafnySequence.asUnicodeString("mcpServer"), dafny.DafnySequence.asUnicodeString("file:///w"), dafny.DafnySequence.asUnicodeString("MCP"), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      Cust _pat_let_tv14 = _71_mc;
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed424 -> {
        SessionState _pat_let212_0 = ((SessionState)(java.lang.Object)(boxed424));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let212_0, boxed425 -> {
          SessionState _72_dt__update__tmp_h23 = ((SessionState)(java.lang.Object)(boxed425));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), _pat_let_tv14), boxed426 -> {
            dafny.DafnySequence<? extends Cust> _pat_let213_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed426));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let213_0, boxed427 -> {
              dafny.DafnySequence<? extends Cust> _73_dt__update_hcustomizations_h2 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed427));
              return SessionState.create((_72_dt__update__tmp_h23).dtor_provider(), (_72_dt__update__tmp_h23).dtor_title(), (_72_dt__update__tmp_h23).dtor_status(), (_72_dt__update__tmp_h23).dtor_lifecycle(), (_72_dt__update__tmp_h23).dtor_activity(), (_72_dt__update__tmp_h23).dtor_config(), (_72_dt__update__tmp_h23).dtor_meta(), (_72_dt__update__tmp_h23).dtor_creationError(), (_72_dt__update__tmp_h23).dtor_serverTools(), (_72_dt__update__tmp_h23).dtor_changesets(), (_72_dt__update__tmp_h23).dtor_canvases(), (_72_dt__update__tmp_h23).dtor_openCanvases(), (_72_dt__update__tmp_h23).dtor_defaultChat(), (_72_dt__update__tmp_h23).dtor_activeClients(), (_72_dt__update__tmp_h23).dtor_chats(), _73_dt__update_hcustomizations_h2, (_72_dt__update__tmp_h23).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_McpServerStartRequested(dafny.DafnySequence.asUnicodeString("mcp-1")))).dtor_customizations()).equals(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_71_mc, boxed428 -> {
        Cust _pat_let214_0 = ((Cust)(java.lang.Object)(boxed428));
        return ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_pat_let214_0, boxed429 -> {
          Cust _74_dt__update__tmp_h24 = ((Cust)(java.lang.Object)(boxed429));
          return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed430 -> {
            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let215_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed430));
            return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(_pat_let215_0, boxed431 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _75_dt__update_hchannel_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed431));
              return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("starting"))) }))), boxed432 -> {
                agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let216_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed432));
                return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(_pat_let216_0, boxed433 -> {
                  agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _76_dt__update_hstate_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed433));
                  return Cust.create((_74_dt__update__tmp_h24).dtor_id(), (_74_dt__update__tmp_h24).dtor_kind(), (_74_dt__update__tmp_h24).dtor_uri(), (_74_dt__update__tmp_h24).dtor_name(), (_74_dt__update__tmp_h24).dtor_enabled(), _76_dt__update_hstate_h0, _75_dt__update_hchannel_h0);
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
      )))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Cust _pat_let_tv15 = _71_mc;
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed434 -> {
        SessionState _pat_let217_0 = ((SessionState)(java.lang.Object)(boxed434));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let217_0, boxed435 -> {
          SessionState _77_dt__update__tmp_h25 = ((SessionState)(java.lang.Object)(boxed435));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), _pat_let_tv15), boxed436 -> {
            dafny.DafnySequence<? extends Cust> _pat_let218_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed436));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let218_0, boxed437 -> {
              dafny.DafnySequence<? extends Cust> _78_dt__update_hcustomizations_h3 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed437));
              return SessionState.create((_77_dt__update__tmp_h25).dtor_provider(), (_77_dt__update__tmp_h25).dtor_title(), (_77_dt__update__tmp_h25).dtor_status(), (_77_dt__update__tmp_h25).dtor_lifecycle(), (_77_dt__update__tmp_h25).dtor_activity(), (_77_dt__update__tmp_h25).dtor_config(), (_77_dt__update__tmp_h25).dtor_meta(), (_77_dt__update__tmp_h25).dtor_creationError(), (_77_dt__update__tmp_h25).dtor_serverTools(), (_77_dt__update__tmp_h25).dtor_changesets(), (_77_dt__update__tmp_h25).dtor_canvases(), (_77_dt__update__tmp_h25).dtor_openCanvases(), (_77_dt__update__tmp_h25).dtor_defaultChat(), (_77_dt__update__tmp_h25).dtor_activeClients(), (_77_dt__update__tmp_h25).dtor_chats(), _78_dt__update_hcustomizations_h3, (_77_dt__update__tmp_h25).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_McpServerStateChanged(dafny.DafnySequence.asUnicodeString("mcp-1"), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("ready"))) })), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("ch")))))).dtor_customizations()).equals(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_71_mc, boxed438 -> {
        Cust _pat_let219_0 = ((Cust)(java.lang.Object)(boxed438));
        return ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_pat_let219_0, boxed439 -> {
          Cust _79_dt__update__tmp_h26 = ((Cust)(java.lang.Object)(boxed439));
          return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("ch"))), boxed440 -> {
            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let220_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed440));
            return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(_pat_let220_0, boxed441 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _80_dt__update_hchannel_h1 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed441));
              return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("ready"))) }))), boxed442 -> {
                agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let221_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed442));
                return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(_pat_let221_0, boxed443 -> {
                  agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _81_dt__update_hstate_h1 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed443));
                  return Cust.create((_79_dt__update__tmp_h26).dtor_id(), (_79_dt__update__tmp_h26).dtor_kind(), (_79_dt__update__tmp_h26).dtor_uri(), (_79_dt__update__tmp_h26).dtor_name(), (_79_dt__update__tmp_h26).dtor_enabled(), _81_dt__update_hstate_h1, _80_dt__update_hchannel_h1);
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
      )))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Cust _pat_let_tv16 = _24_pa;
      Cust _pat_let_tv17 = _24_pa;
      if (java.util.Objects.equals(__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed444 -> {
        SessionState _pat_let222_0 = ((SessionState)(java.lang.Object)(boxed444));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let222_0, boxed445 -> {
          SessionState _82_dt__update__tmp_h27 = ((SessionState)(java.lang.Object)(boxed445));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), _pat_let_tv16), boxed446 -> {
            dafny.DafnySequence<? extends Cust> _pat_let223_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed446));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let223_0, boxed447 -> {
              dafny.DafnySequence<? extends Cust> _83_dt__update_hcustomizations_h4 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed447));
              return SessionState.create((_82_dt__update__tmp_h27).dtor_provider(), (_82_dt__update__tmp_h27).dtor_title(), (_82_dt__update__tmp_h27).dtor_status(), (_82_dt__update__tmp_h27).dtor_lifecycle(), (_82_dt__update__tmp_h27).dtor_activity(), (_82_dt__update__tmp_h27).dtor_config(), (_82_dt__update__tmp_h27).dtor_meta(), (_82_dt__update__tmp_h27).dtor_creationError(), (_82_dt__update__tmp_h27).dtor_serverTools(), (_82_dt__update__tmp_h27).dtor_changesets(), (_82_dt__update__tmp_h27).dtor_canvases(), (_82_dt__update__tmp_h27).dtor_openCanvases(), (_82_dt__update__tmp_h27).dtor_defaultChat(), (_82_dt__update__tmp_h27).dtor_activeClients(), (_82_dt__update__tmp_h27).dtor_chats(), _83_dt__update_hcustomizations_h4, (_82_dt__update__tmp_h27).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_McpServerStateChanged(dafny.DafnySequence.asUnicodeString("plugin-a"), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("ready"))) })), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("nope"))))), ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed448 -> {
        SessionState _pat_let224_0 = ((SessionState)(java.lang.Object)(boxed448));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let224_0, boxed449 -> {
          SessionState _84_dt__update__tmp_h28 = ((SessionState)(java.lang.Object)(boxed449));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), _pat_let_tv17), boxed450 -> {
            dafny.DafnySequence<? extends Cust> _pat_let225_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed450));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let225_0, boxed451 -> {
              dafny.DafnySequence<? extends Cust> _85_dt__update_hcustomizations_h5 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed451));
              return SessionState.create((_84_dt__update__tmp_h28).dtor_provider(), (_84_dt__update__tmp_h28).dtor_title(), (_84_dt__update__tmp_h28).dtor_status(), (_84_dt__update__tmp_h28).dtor_lifecycle(), (_84_dt__update__tmp_h28).dtor_activity(), (_84_dt__update__tmp_h28).dtor_config(), (_84_dt__update__tmp_h28).dtor_meta(), (_84_dt__update__tmp_h28).dtor_creationError(), (_84_dt__update__tmp_h28).dtor_serverTools(), (_84_dt__update__tmp_h28).dtor_changesets(), (_84_dt__update__tmp_h28).dtor_canvases(), (_84_dt__update__tmp_h28).dtor_openCanvases(), (_84_dt__update__tmp_h28).dtor_defaultChat(), (_84_dt__update__tmp_h28).dtor_activeClients(), (_84_dt__update__tmp_h28).dtor_chats(), _85_dt__update_hcustomizations_h5, (_84_dt__update__tmp_h28).dtor_inputNeeded());
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
      Cust _pat_let_tv18 = _71_mc;
      if (((__default.apply1(((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_0_s, boxed452 -> {
        SessionState _pat_let226_0 = ((SessionState)(java.lang.Object)(boxed452));
        return ((SessionState)(java.lang.Object)(dafny.Helpers.<SessionState, SessionState>Let(_pat_let226_0, boxed453 -> {
          SessionState _86_dt__update__tmp_h29 = ((SessionState)(java.lang.Object)(boxed453));
          return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), _pat_let_tv18), boxed454 -> {
            dafny.DafnySequence<? extends Cust> _pat_let227_0 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed454));
            return ((SessionState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Cust>, SessionState>Let(_pat_let227_0, boxed455 -> {
              dafny.DafnySequence<? extends Cust> _87_dt__update_hcustomizations_h6 = ((dafny.DafnySequence<? extends Cust>)(java.lang.Object)(boxed455));
              return SessionState.create((_86_dt__update__tmp_h29).dtor_provider(), (_86_dt__update__tmp_h29).dtor_title(), (_86_dt__update__tmp_h29).dtor_status(), (_86_dt__update__tmp_h29).dtor_lifecycle(), (_86_dt__update__tmp_h29).dtor_activity(), (_86_dt__update__tmp_h29).dtor_config(), (_86_dt__update__tmp_h29).dtor_meta(), (_86_dt__update__tmp_h29).dtor_creationError(), (_86_dt__update__tmp_h29).dtor_serverTools(), (_86_dt__update__tmp_h29).dtor_changesets(), (_86_dt__update__tmp_h29).dtor_canvases(), (_86_dt__update__tmp_h29).dtor_openCanvases(), (_86_dt__update__tmp_h29).dtor_defaultChat(), (_86_dt__update__tmp_h29).dtor_activeClients(), (_86_dt__update__tmp_h29).dtor_chats(), _87_dt__update_hcustomizations_h6, (_86_dt__update__tmp_h29).dtor_inputNeeded());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), SessionAction.create_McpServerStopRequested(dafny.DafnySequence.asUnicodeString("mcp-1")))).dtor_customizations()).equals(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_71_mc, boxed456 -> {
        Cust _pat_let228_0 = ((Cust)(java.lang.Object)(boxed456));
        return ((Cust)(java.lang.Object)(dafny.Helpers.<Cust, Cust>Let(_pat_let228_0, boxed457 -> {
          Cust _88_dt__update__tmp_h30 = ((Cust)(java.lang.Object)(boxed457));
          return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed458 -> {
            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let229_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed458));
            return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(_pat_let229_0, boxed459 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _89_dt__update_hchannel_h2 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed459));
              return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("stopped"))) }))), boxed460 -> {
                agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let230_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed460));
                return ((Cust)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Cust>Let(_pat_let230_0, boxed461 -> {
                  agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _90_dt__update_hstate_h2 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed461));
                  return Cust.create((_88_dt__update__tmp_h30).dtor_id(), (_88_dt__update__tmp_h30).dtor_kind(), (_88_dt__update__tmp_h30).dtor_uri(), (_88_dt__update__tmp_h30).dtor_name(), (_88_dt__update__tmp_h30).dtor_enabled(), _90_dt__update_hstate_h2, _89_dt__update_hchannel_h2);
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
      )))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (((__default.apply1(_0_s, SessionAction.create_CustomizationUpdated(_71_mc))).dtor_customizations()).equals(dafny.DafnySequence.<Cust> of(Cust._typeDescriptor(), _71_mc))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
    }
    return new dafny.Tuple2<>(pass, modeled);
  }
  public static java.util.function.Function<Cust, dafny.DafnySequence<? extends dafny.CodePoint>> custKey()
  {
    return ((java.util.function.Function<Cust, dafny.DafnySequence<? extends dafny.CodePoint>>)(_0_x_boxed0) -> {
      Cust _0_x = ((Cust)(java.lang.Object)(_0_x_boxed0));
      return (_0_x).dtor_id();
    });
  }
  public static java.util.function.Function<Chat, dafny.DafnySequence<? extends dafny.CodePoint>> chatKey()
  {
    return ((java.util.function.Function<Chat, dafny.DafnySequence<? extends dafny.CodePoint>>)(_0_x_boxed0) -> {
      Chat _0_x = ((Chat)(java.lang.Object)(_0_x_boxed0));
      return (_0_x).dtor_resource();
    });
  }
  public static java.util.function.Function<Client, dafny.DafnySequence<? extends dafny.CodePoint>> clientKey()
  {
    return ((java.util.function.Function<Client, dafny.DafnySequence<? extends dafny.CodePoint>>)(_0_x_boxed0) -> {
      Client _0_x = ((Client)(java.lang.Object)(_0_x_boxed0));
      return (_0_x).dtor_clientId();
    });
  }
  public static java.util.function.Function<InputReq, dafny.DafnySequence<? extends dafny.CodePoint>> inputKey()
  {
    return ((java.util.function.Function<InputReq, dafny.DafnySequence<? extends dafny.CodePoint>>)(_0_x_boxed0) -> {
      InputReq _0_x = ((InputReq)(java.lang.Object)(_0_x_boxed0));
      return (_0_x).dtor_id();
    });
  }
  public static java.util.function.Function<Cust, Boolean> mcpPred()
  {
    return ((java.util.function.Function<Cust, Boolean>)(_0_c_boxed0) -> {
      Cust _0_c = ((Cust)(java.lang.Object)(_0_c_boxed0));
      return ((_0_c).dtor_kind()).equals(dafny.DafnySequence.asUnicodeString("mcpServer"));
    });
  }
  public static int B__READ()
  {
    return 32;
  }
  public static int B__ARCH()
  {
    return 64;
  }
  public static int B__INPUT()
  {
    return 16;
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.Session._default";
  }
}
