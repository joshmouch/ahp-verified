// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.AhpSkeleton;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static java.math.BigInteger Pow10(java.math.BigInteger k) {
    java.math.BigInteger _0___accumulator = java.math.BigInteger.ONE;
    TAIL_CALL_START: while (true) {
      if ((k).signum() == 0) {
        return (java.math.BigInteger.ONE).multiply(_0___accumulator);
      } else {
        _0___accumulator = (_0___accumulator).multiply(java.math.BigInteger.valueOf(10L));
        java.math.BigInteger _in0 = (k).subtract(java.math.BigInteger.ONE);
        k = _in0;
        continue TAIL_CALL_START;
      }
    }
  }
  public static Option<java.math.BigInteger> decInt(java.math.BigInteger m, java.math.BigInteger e)
  {
    if ((e).signum() != -1) {
      return Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, (m).multiply(__default.Pow10(e)));
    } else if ((dafny.DafnyEuclidean.EuclideanModulus(m, __default.Pow10((java.math.BigInteger.ZERO).subtract(e)))).signum() == 0) {
      return Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, dafny.DafnyEuclidean.EuclideanDivision(m, __default.Pow10((java.math.BigInteger.ZERO).subtract(e))));
    } else {
      return Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER);
    }
  }
  public static Option<java.math.BigInteger> jsonInt(agency.open.ahp.ConfluxCodec.Json j) {
    if ((j).is_JNum()) {
      return Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, (j).dtor_n());
    } else if ((j).is_JDec()) {
      return __default.decInt((j).dtor_mantissa(), (j).dtor_exp());
    } else {
      return Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER);
    }
  }
  public static java.math.BigInteger asInt(Option<agency.open.ahp.ConfluxCodec.Json> o) {
    if ((o).is_None()) {
      return java.math.BigInteger.ZERO;
    } else {
      Option<java.math.BigInteger> _0_r = __default.jsonInt((o).dtor_value());
      if ((_0_r).is_Some()) {
        return (_0_r).dtor_value();
      } else {
        return java.math.BigInteger.ZERO;
      }
    }
  }
  public static Option<java.math.BigInteger> optInt(Option<agency.open.ahp.ConfluxCodec.Json> o) {
    if ((o).is_Some()) {
      return __default.jsonInt((o).dtor_value());
    } else {
      return Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER);
    }
  }
  public static java.math.BigInteger EncodeStatus(int s) {
    return dafny.Helpers.unsignedToBigInteger(s);
  }
  public static dafny.Tuple2<RootState, ReduceOutcome> ApplyToRoot(RootState s, RootAction a, java.math.BigInteger now)
  {
    RootAction _source0 = a;
    if (_source0.is_RootAgentsChanged()) {
      dafny.DafnySequence<? extends AgentInfo> _0___mcc_h0 = ((agency.open.ahp.AhpSkeleton.RootAction_RootAgentsChanged)_source0)._agents;
      dafny.DafnySequence<? extends AgentInfo> _1_ag = _0___mcc_h0;
      return dafny.Tuple2.<RootState, ReduceOutcome>create(((RootState)(java.lang.Object)(dafny.Helpers.<RootState, RootState>Let(s, boxed0 -> {
  RootState _pat_let0_0 = ((RootState)(java.lang.Object)(boxed0));
  return ((RootState)(java.lang.Object)(dafny.Helpers.<RootState, RootState>Let(_pat_let0_0, boxed1 -> {
    RootState _2_dt__update__tmp_h0 = ((RootState)(java.lang.Object)(boxed1));
    return ((RootState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends AgentInfo>, RootState>Let(_1_ag, boxed2 -> {
      dafny.DafnySequence<? extends AgentInfo> _pat_let1_0 = ((dafny.DafnySequence<? extends AgentInfo>)(java.lang.Object)(boxed2));
      return ((RootState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends AgentInfo>, RootState>Let(_pat_let1_0, boxed3 -> {
        dafny.DafnySequence<? extends AgentInfo> _3_dt__update_hagents_h0 = ((dafny.DafnySequence<? extends AgentInfo>)(java.lang.Object)(boxed3));
        return RootState.create(_3_dt__update_hagents_h0, (_2_dt__update__tmp_h0).dtor_activeSessions(), (_2_dt__update__tmp_h0).dtor_terminals(), (_2_dt__update__tmp_h0).dtor_config());
      }
      )));
    }
    )));
  }
  )));
}
))), ReduceOutcome.create_Applied());
    } else if (_source0.is_RootActiveSessionsChanged()) {
      java.math.BigInteger _4___mcc_h1 = ((agency.open.ahp.AhpSkeleton.RootAction_RootActiveSessionsChanged)_source0)._activeSessions;
      java.math.BigInteger _5_n = _4___mcc_h1;
      return dafny.Tuple2.<RootState, ReduceOutcome>create(((RootState)(java.lang.Object)(dafny.Helpers.<RootState, RootState>Let(s, boxed4 -> {
  RootState _pat_let2_0 = ((RootState)(java.lang.Object)(boxed4));
  return ((RootState)(java.lang.Object)(dafny.Helpers.<RootState, RootState>Let(_pat_let2_0, boxed5 -> {
    RootState _6_dt__update__tmp_h1 = ((RootState)(java.lang.Object)(boxed5));
    return ((RootState)(java.lang.Object)(dafny.Helpers.<Option<java.math.BigInteger>, RootState>Let(Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, _5_n), boxed6 -> {
      Option<java.math.BigInteger> _pat_let3_0 = ((Option<java.math.BigInteger>)(java.lang.Object)(boxed6));
      return ((RootState)(java.lang.Object)(dafny.Helpers.<Option<java.math.BigInteger>, RootState>Let(_pat_let3_0, boxed7 -> {
        Option<java.math.BigInteger> _7_dt__update_hactiveSessions_h0 = ((Option<java.math.BigInteger>)(java.lang.Object)(boxed7));
        return RootState.create((_6_dt__update__tmp_h1).dtor_agents(), _7_dt__update_hactiveSessions_h0, (_6_dt__update__tmp_h1).dtor_terminals(), (_6_dt__update__tmp_h1).dtor_config());
      }
      )));
    }
    )));
  }
  )));
}
))), ReduceOutcome.create_Applied());
    } else if (_source0.is_RootTerminalsChanged()) {
      dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json> _8___mcc_h2 = ((agency.open.ahp.AhpSkeleton.RootAction_RootTerminalsChanged)_source0)._terminals;
      dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json> _9_t = _8___mcc_h2;
      return dafny.Tuple2.<RootState, ReduceOutcome>create(((RootState)(java.lang.Object)(dafny.Helpers.<RootState, RootState>Let(s, boxed8 -> {
  RootState _pat_let4_0 = ((RootState)(java.lang.Object)(boxed8));
  return ((RootState)(java.lang.Object)(dafny.Helpers.<RootState, RootState>Let(_pat_let4_0, boxed9 -> {
    RootState _10_dt__update__tmp_h2 = ((RootState)(java.lang.Object)(boxed9));
    return ((RootState)(java.lang.Object)(dafny.Helpers.<Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>, RootState>Let(Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_Some(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), _9_t), boxed10 -> {
      Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _pat_let5_0 = ((Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>)(java.lang.Object)(boxed10));
      return ((RootState)(java.lang.Object)(dafny.Helpers.<Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>, RootState>Let(_pat_let5_0, boxed11 -> {
        Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _11_dt__update_hterminals_h0 = ((Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>)(java.lang.Object)(boxed11));
        return RootState.create((_10_dt__update__tmp_h2).dtor_agents(), (_10_dt__update__tmp_h2).dtor_activeSessions(), _11_dt__update_hterminals_h0, (_10_dt__update__tmp_h2).dtor_config());
      }
      )));
    }
    )));
  }
  )));
}
))), ReduceOutcome.create_Applied());
    } else if (_source0.is_RootConfigChanged()) {
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _12___mcc_h3 = ((agency.open.ahp.AhpSkeleton.RootAction_RootConfigChanged)_source0)._config;
      boolean _13___mcc_h4 = ((agency.open.ahp.AhpSkeleton.RootAction_RootConfigChanged)_source0)._replace;
      boolean _14_repl = _13___mcc_h4;
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _15_cfg = _12___mcc_h3;
      Option<RootConfig> _source1 = (s).dtor_config();
      if (_source1.is_None()) {
        return dafny.Tuple2.<RootState, ReduceOutcome>create(s, ReduceOutcome.create_NoOp());
      } else {
        RootConfig _16___mcc_h6 = ((agency.open.ahp.AhpSkeleton.Option_Some<RootConfig>)_source1)._value;
        RootConfig _17_c = _16___mcc_h6;
        dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _18_newValues = ((_14_repl) ? (_15_cfg) : (dafny.DafnyMap.<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>merge((_17_c).dtor_values(), _15_cfg)));
        return dafny.Tuple2.<RootState, ReduceOutcome>create(((RootState)(java.lang.Object)(dafny.Helpers.<RootState, RootState>Let(s, boxed12 -> {
  RootState _pat_let6_0 = ((RootState)(java.lang.Object)(boxed12));
  return ((RootState)(java.lang.Object)(dafny.Helpers.<RootState, RootState>Let(_pat_let6_0, boxed13 -> {
    RootState _19_dt__update__tmp_h3 = ((RootState)(java.lang.Object)(boxed13));
    return ((RootState)(java.lang.Object)(dafny.Helpers.<Option<RootConfig>, RootState>Let(Option.<RootConfig>create_Some(RootConfig._typeDescriptor(), RootConfig.create((_17_c).dtor_schema(), _18_newValues)), boxed14 -> {
      Option<RootConfig> _pat_let7_0 = ((Option<RootConfig>)(java.lang.Object)(boxed14));
      return ((RootState)(java.lang.Object)(dafny.Helpers.<Option<RootConfig>, RootState>Let(_pat_let7_0, boxed15 -> {
        Option<RootConfig> _20_dt__update_hconfig_h0 = ((Option<RootConfig>)(java.lang.Object)(boxed15));
        return RootState.create((_19_dt__update__tmp_h3).dtor_agents(), (_19_dt__update__tmp_h3).dtor_activeSessions(), (_19_dt__update__tmp_h3).dtor_terminals(), _20_dt__update_hconfig_h0);
      }
      )));
    }
    )));
  }
  )));
}
))), ReduceOutcome.create_Applied());
      }
    } else {
      agency.open.ahp.ConfluxCodec.Json _21___mcc_h5 = ((agency.open.ahp.AhpSkeleton.RootAction_RootUnknown)_source0)._raw;
      return dafny.Tuple2.<RootState, ReduceOutcome>create(s, ReduceOutcome.create_NoOp());
    }
  }
  public static boolean IsUnknownRoot(RootAction a) {
    return (a).is_RootUnknown();
  }
  public static boolean ConfigPreconditionFails(RootState s, RootAction a)
  {
    return ((a).is_RootConfigChanged()) && (((s).dtor_config()).is_None());
  }
  public static RootState FoldRoot(RootState s, dafny.DafnySequence<? extends RootAction> actions, java.math.BigInteger now)
  {
    return agency.open.ahp.ConfluxContract.__default.<RootState, RootAction>Fold(RootState._typeDescriptor(), RootAction._typeDescriptor(), ((java.util.function.Function<java.math.BigInteger, dafny.Function2<RootState, RootAction, RootState>>)(_0_now) -> {return ((dafny.Function2<RootState, RootAction, RootState>)(_1_s_k_boxed0, _2_a_boxed0) -> {
      RootState _1_s_k = ((RootState)(java.lang.Object)(_1_s_k_boxed0));
      RootAction _2_a = ((RootAction)(java.lang.Object)(_2_a_boxed0));
      return (__default.ApplyToRoot(_1_s_k, _2_a, _0_now)).dtor__0();
    });}).apply(now), s, actions);
  }
  public static RootState apply1(RootState s, RootAction a)
  {
    return (__default.ApplyToRoot(s, a, java.math.BigInteger.valueOf(9999L))).dtor__0();
  }
  public static dafny.Tuple2 RunCorpus()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    java.math.BigInteger total = java.math.BigInteger.ZERO;
    if(true) {
      total = java.math.BigInteger.valueOf(7L);
      pass = java.math.BigInteger.ZERO;
      RootState _0_empty;
      _0_empty = RootState.create(dafny.DafnySequence.<AgentInfo> empty(AgentInfo._typeDescriptor()), Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), Option.<RootConfig>create_None(RootConfig._typeDescriptor()));
      dafny.DafnySequence<? extends AgentInfo> _1_ag;
      _1_ag = dafny.DafnySequence.<AgentInfo> of(AgentInfo._typeDescriptor(), AgentInfo.create(dafny.DafnySequence.asUnicodeString("copilot"), dafny.DafnySequence.asUnicodeString("Copilot"), dafny.DafnySequence.asUnicodeString("AI"), dafny.DafnySequence.<dafny.DafnySequence<? extends dafny.CodePoint>> empty(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))));
      if (java.util.Objects.equals(__default.apply1(_0_empty, RootAction.create_RootAgentsChanged(_1_ag)), RootState.create(_1_ag, Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), Option.<RootConfig>create_None(RootConfig._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_0_empty, RootAction.create_RootActiveSessionsChanged(java.math.BigInteger.valueOf(5L))), RootState.create(dafny.DafnySequence.<AgentInfo> empty(AgentInfo._typeDescriptor()), Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.valueOf(5L)), Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), Option.<RootConfig>create_None(RootConfig._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json> _2_terms;
      _2_terms = dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> of(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("resource"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("agenthost:/terminal/1"))) })), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("resource"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("agenthost:/terminal/2"))) })));
      if (java.util.Objects.equals(__default.apply1(_0_empty, RootAction.create_RootTerminalsChanged(_2_terms)), RootState.create(dafny.DafnySequence.<AgentInfo> empty(AgentInfo._typeDescriptor()), Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_Some(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), _2_terms), Option.<RootConfig>create_None(RootConfig._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_0_empty, RootAction.create_RootUnknown(agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("type"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("root/nonExistentAction"))) })))), _0_empty)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      agency.open.ahp.ConfluxCodec.Json _3_sch;
      _3_sch = agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("type"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("object"))) }));
      RootState _4_cfg127;
      _4_cfg127 = RootState.create(dafny.DafnySequence.<AgentInfo> empty(AgentInfo._typeDescriptor()), Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), Option.<RootConfig>create_Some(RootConfig._typeDescriptor(), RootConfig.create(_3_sch, dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("theme"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("light"))) }))));
      RootState _5_exp127;
      _5_exp127 = RootState.create(dafny.DafnySequence.<AgentInfo> empty(AgentInfo._typeDescriptor()), Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), Option.<RootConfig>create_Some(RootConfig._typeDescriptor(), RootConfig.create(_3_sch, dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("theme"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("dark"))) }))));
      if (java.util.Objects.equals(__default.apply1(_4_cfg127, RootAction.create_RootConfigChanged(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("theme"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("dark"))) }), false)), _5_exp127)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_0_empty, RootAction.create_RootConfigChanged(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("theme"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("dark"))) }), false)), _0_empty)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      RootState _6_cfg130;
      _6_cfg130 = RootState.create(dafny.DafnySequence.<AgentInfo> empty(AgentInfo._typeDescriptor()), Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), Option.<RootConfig>create_Some(RootConfig._typeDescriptor(), RootConfig.create(_3_sch, dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("theme"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("light"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("locale"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("en"))) }))));
      RootState _7_exp130;
      _7_exp130 = RootState.create(dafny.DafnySequence.<AgentInfo> empty(AgentInfo._typeDescriptor()), Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), Option.<RootConfig>create_Some(RootConfig._typeDescriptor(), RootConfig.create(_3_sch, dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("theme"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("dark"))) }))));
      if (java.util.Objects.equals(__default.apply1(_6_cfg130, RootAction.create_RootConfigChanged(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("theme"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("dark"))) }), true)), _7_exp130)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
    }
    return new dafny.Tuple2<>(pass, total);
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.AhpSkeleton._default";
  }
}
