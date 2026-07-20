// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.Chat;

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
  public static agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> jNoNull(agency.open.ahp.ConfluxCodec.Json j) {
    if ((j).is_JNull()) {
      return agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
    } else {
      return agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), j);
    }
  }
  public static dafny.DafnySequence<? extends QMsg> upsertQ(dafny.DafnySequence<? extends QMsg> qs, QMsg m)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, QMsg>SeqUpsertById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), QMsg._typeDescriptor(), __default.qKey(), qs, m);
  }
  public static dafny.DafnySequence<? extends QMsg> removeQ(dafny.DafnySequence<? extends QMsg> qs, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, QMsg>SeqRemoveById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), QMsg._typeDescriptor(), __default.qKey(), qs, id);
  }
  public static boolean containsQ(dafny.DafnySequence<? extends QMsg> qs, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    TAIL_CALL_START: while (true) {
      if ((java.math.BigInteger.valueOf((qs).length())).signum() == 0) {
        return false;
      } else if (((((QMsg)(java.lang.Object)((qs).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).dtor_id()).equals(id)) {
        return true;
      } else {
        dafny.DafnySequence<? extends QMsg> _in0 = (qs).drop(java.math.BigInteger.ONE);
        dafny.DafnySequence<? extends dafny.CodePoint> _in1 = id;
        qs = _in0;
        id = _in1;
        continue TAIL_CALL_START;
      }
    }
  }
  public static dafny.DafnySequence<? extends QMsg> findQ(dafny.DafnySequence<? extends QMsg> qs, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    dafny.DafnySequence<? extends QMsg> _0_found = agency.open.ahp.ConfluxOperators.__default.<QMsg>Filter(QMsg._typeDescriptor(), ((java.util.function.Function<dafny.DafnySequence<? extends dafny.CodePoint>, java.util.function.Function<QMsg, Boolean>>)(_1_id) -> {return ((java.util.function.Function<QMsg, Boolean>)(_2_q_boxed0) -> {
      QMsg _2_q = ((QMsg)(java.lang.Object)(_2_q_boxed0));
      return ((_2_q).dtor_id()).equals(_1_id);
    });}).apply(id), qs);
    if ((java.math.BigInteger.valueOf((_0_found).length())).signum() == 0) {
      return dafny.DafnySequence.<QMsg> empty(QMsg._typeDescriptor());
    } else {
      return dafny.DafnySequence.<QMsg> of(QMsg._typeDescriptor(), ((QMsg)(java.lang.Object)((_0_found).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO))))));
    }
  }
  public static dafny.DafnySequence<? extends QMsg> reorderQ(dafny.DafnySequence<? extends QMsg> qs, dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> order)
  {
    return agency.open.ahp.ConfluxOrderedLog.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, QMsg>SeqReorderByKeys(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), QMsg._typeDescriptor(), __default.qKey(), order, qs);
  }
  public static dafny.DafnySequence<? extends Turn> keepUpTo(dafny.DafnySequence<? extends Turn> ts, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return agency.open.ahp.ConfluxOrderedLog.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Turn>SeqKeepThrough(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Turn._typeDescriptor(), ((java.util.function.Function<Turn, dafny.DafnySequence<? extends dafny.CodePoint>>)(_0_t_boxed0) -> {
      Turn _0_t = ((Turn)(java.lang.Object)(_0_t_boxed0));
      return (_0_t).dtor_turnId();
    }), id, ts);
  }
  public static boolean hasTurn(dafny.DafnySequence<? extends Turn> ts, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return ((dafny.Function2<dafny.DafnySequence<? extends Turn>, dafny.DafnySequence<? extends dafny.CodePoint>, Boolean>)(_0_ts, _1_id) -> {return dafny.Helpers.Quantifier(dafny.Helpers.IntegerRange(java.math.BigInteger.ZERO, java.math.BigInteger.valueOf((_0_ts).length())), false, ((_exists_var_0_boxed0) -> {
      java.math.BigInteger _exists_var_0 = ((java.math.BigInteger)(java.lang.Object)(_exists_var_0_boxed0));
      java.math.BigInteger _2_i = (java.math.BigInteger)_exists_var_0;
      return (((_2_i).signum() != -1) && ((_2_i).compareTo(java.math.BigInteger.valueOf((_0_ts).length())) < 0)) && (((((Turn)(java.lang.Object)((_0_ts).select(dafny.Helpers.toInt((_2_i)))))).dtor_turnId()).equals(_1_id));
    }));}).apply(ts, id);
  }
  public static dafny.DafnySequence<? extends Turn> dedupPrepend(dafny.DafnySequence<? extends Turn> loaded, dafny.DafnySequence<? extends Turn> existing)
  {
    return dafny.DafnySequence.<Turn>concatenate(agency.open.ahp.ConfluxOperators.__default.<Turn>Filter(Turn._typeDescriptor(), ((java.util.function.Function<dafny.DafnySequence<? extends Turn>, java.util.function.Function<Turn, Boolean>>)(_0_existing) -> {return ((java.util.function.Function<Turn, Boolean>)(_1_t_boxed0) -> {
      Turn _1_t = ((Turn)(java.lang.Object)(_1_t_boxed0));
      return !(__default.hasTurn(_0_existing, (_1_t).dtor_turnId()));
    });}).apply(existing), loaded), existing);
  }
  public static boolean isOpenInput(Part p) {
    return ((p).is_PInputRequest()) && (((p).dtor_response()).is_None());
  }
  public static boolean hasOpenReq(dafny.DafnySequence<? extends Part> ps, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    if ((java.math.BigInteger.valueOf((ps).length())).signum() == 0) {
      return false;
    } else {
      return ((((((Part)(java.lang.Object)((ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).is_PInputRequest()) && (((((Part)(java.lang.Object)((ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).dtor_response()).is_None())) && ((((((Part)(java.lang.Object)((ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).dtor_req()).dtor_id()).equals(id))) || (__default.hasOpenReq((ps).drop(java.math.BigInteger.ONE), id));
    }
  }
  public static InputReq mergeInputReq(InputReq existing, InputReq req)
  {
    if ((java.math.BigInteger.valueOf(((req).dtor_answers()).size())).signum() == 1) {
      return req;
    } else {
      InputReq _0_dt__update__tmp_h0 = req;
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _1_dt__update_hanswers_h0 = (existing).dtor_answers();
      return InputReq.create((_0_dt__update__tmp_h0).dtor_id(), _1_dt__update_hanswers_h0, (_0_dt__update__tmp_h0).dtor_open());
    }
  }
  public static java.util.function.Function<Part, Part> mergeInputPart(InputReq req) {
    return ((java.util.function.Function<InputReq, java.util.function.Function<Part, Part>>)(_0_req) -> {return ((java.util.function.Function<Part, Part>)(_1_p_boxed0) -> {
      Part _1_p = ((Part)(java.lang.Object)(_1_p_boxed0));
      return ((java.util.function.Function<Part, Part>)(_source0_boxed0) -> {
        Part _source0 = ((Part)(java.lang.Object)(_source0_boxed0));
        if (_source0.is_PMarkdown()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _2___mcc_h0 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _3___mcc_h1 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._content;
          return _1_p;
        } else if (_source0.is_PReasoning()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h4 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _5___mcc_h5 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._content;
          return _1_p;
        } else if (_source0.is_PToolCall()) {
          ToolCall _6___mcc_h8 = ((agency.open.ahp.Chat.Part_PToolCall)_source0)._tc;
          return _1_p;
        } else {
          InputReq _7___mcc_h10 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._req;
          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _8___mcc_h11 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._response;
          InputReq _9_existing = _7___mcc_h10;
          return Part.create_PInputRequest(__default.mergeInputReq(_9_existing, _0_req), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)));
        }
      }).apply(_1_p);
    });}).apply(req);
  }
  public static dafny.DafnySequence<? extends Part> upsertInputPart(dafny.DafnySequence<? extends Part> ps, InputReq req)
  {
    if (__default.hasOpenReq(ps, (req).dtor_id())) {
      return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Part>SeqUpdateByIdWhere(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Part._typeDescriptor(), __default::partKey, __default::isOpenInput, ps, (req).dtor_id(), __default.mergeInputPart(req));
    } else {
      return dafny.DafnySequence.<Part>concatenate(ps, dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PInputRequest(req, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)))));
    }
  }
  public static java.util.function.Function<Part, Part> answerInputPart(dafny.DafnySequence<? extends dafny.CodePoint> qId, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> ans)
  {
    return ((dafny.Function2<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, dafny.DafnySequence<? extends dafny.CodePoint>, java.util.function.Function<Part, Part>>)(_0_ans, _1_qId) -> {return ((java.util.function.Function<Part, Part>)(_2_p_boxed0) -> {
      Part _2_p = ((Part)(java.lang.Object)(_2_p_boxed0));
      return ((java.util.function.Function<Part, Part>)(_source0_boxed0) -> {
        Part _source0 = ((Part)(java.lang.Object)(_source0_boxed0));
        if (_source0.is_PMarkdown()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _3___mcc_h0 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h1 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._content;
          return _2_p;
        } else if (_source0.is_PReasoning()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _5___mcc_h4 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _6___mcc_h5 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._content;
          return _2_p;
        } else if (_source0.is_PToolCall()) {
          ToolCall _7___mcc_h8 = ((agency.open.ahp.Chat.Part_PToolCall)_source0)._tc;
          return _2_p;
        } else {
          InputReq _8___mcc_h10 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._req;
          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _9___mcc_h11 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._response;
          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _10_resp = _9___mcc_h11;
          InputReq _11_r = _8___mcc_h10;
          return Part.create_PInputRequest(((InputReq)(java.lang.Object)(dafny.Helpers.<InputReq, InputReq>Let(_11_r, boxed462 -> {
  InputReq _pat_let231_0 = ((InputReq)(java.lang.Object)(boxed462));
  return ((InputReq)(java.lang.Object)(dafny.Helpers.<InputReq, InputReq>Let(_pat_let231_0, boxed463 -> {
    InputReq _12_dt__update__tmp_h0 = ((InputReq)(java.lang.Object)(boxed463));
    return ((InputReq)(java.lang.Object)(dafny.Helpers.<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>, InputReq>Let((((_0_ans).is_Some()) ? (dafny.DafnyMap.<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>update((_11_r).dtor_answers(), _1_qId, (_0_ans).dtor_value())) : (dafny.DafnyMap.<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>subtract((_11_r).dtor_answers(), dafny.DafnySet.<dafny.DafnySequence<? extends dafny.CodePoint>> of(_1_qId)))), boxed464 -> {
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _pat_let232_0 = ((dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed464));
      return ((InputReq)(java.lang.Object)(dafny.Helpers.<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>, InputReq>Let(_pat_let232_0, boxed465 -> {
        dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _13_dt__update_hanswers_h0 = ((dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed465));
        return InputReq.create((_12_dt__update__tmp_h0).dtor_id(), _13_dt__update_hanswers_h0, (_12_dt__update__tmp_h0).dtor_open());
      }
      )));
    }
    )));
  }
  )));
}
))), _10_resp);
        }
      }).apply(_2_p);
    });}).apply(ans, qId);
  }
  public static dafny.DafnySequence<? extends Part> changeAnswerPart(dafny.DafnySequence<? extends Part> ps, dafny.DafnySequence<? extends dafny.CodePoint> reqId, dafny.DafnySequence<? extends dafny.CodePoint> qId, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> ans)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Part>SeqUpdateByIdWhere(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Part._typeDescriptor(), __default::partKey, __default::isOpenInput, ps, reqId, __default.answerInputPart(qId, ans));
  }
  public static dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> mergeAnswers(dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> drafts, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> completed)
  {
    return ((dafny.Function2<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>>)(_0_drafts, _1_completed) -> {return ((dafny.Function0<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>>)(() -> {
      java.util.HashMap<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json> _coll0 = new java.util.HashMap<>();
      for (dafny.DafnySequence<? extends dafny.CodePoint> _compr_0_boxed0 : (dafny.DafnySet.<dafny.DafnySequence<? extends dafny.CodePoint>>union((_0_drafts).keySet(), (_1_completed).keySet())).Elements()) {
        dafny.DafnySequence<? extends dafny.CodePoint> _compr_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(_compr_0_boxed0));
        dafny.DafnySequence<? extends dafny.CodePoint> _2_k = (dafny.DafnySequence<? extends dafny.CodePoint>)_compr_0;
        if ((dafny.DafnySet.<dafny.DafnySequence<? extends dafny.CodePoint>>union((_0_drafts).keySet(), (_1_completed).keySet())).<dafny.DafnySequence<? extends dafny.CodePoint>>contains(_2_k)) {
          _coll0.put(_2_k,(((_1_completed).<dafny.DafnySequence<? extends dafny.CodePoint>>contains(_2_k)) ? (((agency.open.ahp.ConfluxCodec.Json)(java.lang.Object)((_1_completed).get(_2_k)))) : (((agency.open.ahp.ConfluxCodec.Json)(java.lang.Object)((_0_drafts).get(_2_k))))));
        }
      }
      return new dafny.DafnyMap<dafny.DafnySequence<? extends dafny.CodePoint>,agency.open.ahp.ConfluxCodec.Json>(_coll0);
    })).apply();}).apply(drafts, completed);
  }
  public static java.util.function.Function<Part, Part> completeInputPart(dafny.DafnySequence<? extends dafny.CodePoint> resp, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> answers)
  {
    return ((dafny.Function2<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>, dafny.DafnySequence<? extends dafny.CodePoint>, java.util.function.Function<Part, Part>>)(_0_answers, _1_resp) -> {return ((java.util.function.Function<Part, Part>)(_2_p_boxed0) -> {
      Part _2_p = ((Part)(java.lang.Object)(_2_p_boxed0));
      return ((java.util.function.Function<Part, Part>)(_source0_boxed0) -> {
        Part _source0 = ((Part)(java.lang.Object)(_source0_boxed0));
        if (_source0.is_PMarkdown()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _3___mcc_h0 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h1 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._content;
          return _2_p;
        } else if (_source0.is_PReasoning()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _5___mcc_h4 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _6___mcc_h5 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._content;
          return _2_p;
        } else if (_source0.is_PToolCall()) {
          ToolCall _7___mcc_h8 = ((agency.open.ahp.Chat.Part_PToolCall)_source0)._tc;
          return _2_p;
        } else {
          InputReq _8___mcc_h10 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._req;
          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _9___mcc_h11 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._response;
          InputReq _10_r = _8___mcc_h10;
          return Part.create_PInputRequest(((InputReq)(java.lang.Object)(dafny.Helpers.<InputReq, InputReq>Let(_10_r, boxed466 -> {
  InputReq _pat_let233_0 = ((InputReq)(java.lang.Object)(boxed466));
  return ((InputReq)(java.lang.Object)(dafny.Helpers.<InputReq, InputReq>Let(_pat_let233_0, boxed467 -> {
    InputReq _11_dt__update__tmp_h0 = ((InputReq)(java.lang.Object)(boxed467));
    return ((InputReq)(java.lang.Object)(dafny.Helpers.<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>, InputReq>Let(__default.mergeAnswers((_10_r).dtor_answers(), _0_answers), boxed468 -> {
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _pat_let234_0 = ((dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed468));
      return ((InputReq)(java.lang.Object)(dafny.Helpers.<dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>, InputReq>Let(_pat_let234_0, boxed469 -> {
        dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _12_dt__update_hanswers_h0 = ((dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed469));
        return InputReq.create((_11_dt__update__tmp_h0).dtor_id(), _12_dt__update_hanswers_h0, (_11_dt__update__tmp_h0).dtor_open());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), _1_resp));
        }
      }).apply(_2_p);
    });}).apply(answers, resp);
  }
  public static dafny.DafnySequence<? extends Part> completeAnswerPart(dafny.DafnySequence<? extends Part> ps, dafny.DafnySequence<? extends dafny.CodePoint> reqId, dafny.DafnySequence<? extends dafny.CodePoint> resp, dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> answers)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Part>SeqUpdateByIdWhere(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Part._typeDescriptor(), __default::partKey, __default::isOpenInput, ps, reqId, __default.completeInputPart(resp, answers));
  }
  public static boolean turnMatches(ChatState s, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return (((s).dtor_activeTurn()).is_Some()) && (((((s).dtor_activeTurn()).dtor_value()).dtor_turnId()).equals(id));
  }
  public static dafny.DafnySequence<? extends dafny.CodePoint> optOr(agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> o, dafny.DafnySequence<? extends dafny.CodePoint> d)
  {
    if ((o).is_Some()) {
      return (o).dtor_value();
    } else {
      return d;
    }
  }
  public static boolean anyPendingTC(dafny.DafnySequence<? extends Part> ps) {
    if ((java.math.BigInteger.valueOf((ps).length())).signum() == 0) {
      return false;
    } else {
      return (((((Part)(java.lang.Object)((ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).is_PToolCall()) && ((((((((Part)(java.lang.Object)((ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).dtor_tc()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("pending-confirmation"))) || ((((((Part)(java.lang.Object)((ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).dtor_tc()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("pending-result-confirmation")))) || ((((((Part)(java.lang.Object)((ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).dtor_tc()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("auth-required"))))) || (__default.anyPendingTC((ps).drop(java.math.BigInteger.ONE)));
    }
  }
  public static boolean hasPendingTCState(ChatState s) {
    return (((s).dtor_activeTurn()).is_Some()) && (__default.anyPendingTC((((s).dtor_activeTurn()).dtor_value()).dtor_parts()));
  }
  public static boolean anyOpenInput(dafny.DafnySequence<? extends Part> ps) {
    if ((java.math.BigInteger.valueOf((ps).length())).signum() == 0) {
      return false;
    } else {
      return (((((Part)(java.lang.Object)((ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).is_PInputRequest()) && (((((Part)(java.lang.Object)((ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).dtor_response()).is_None())) || (__default.anyOpenInput((ps).drop(java.math.BigInteger.ONE)));
    }
  }
  public static boolean hasOpenInputState(ChatState s) {
    return (((s).dtor_activeTurn()).is_Some()) && (__default.anyOpenInput((((s).dtor_activeTurn()).dtor_value()).dtor_parts()));
  }
  public static int activityBits(ChatState s, boolean isError)
  {
    if (isError) {
      return __default.ERR();
    } else if ((__default.hasOpenInputState(s)) || (__default.hasPendingTCState(s))) {
      return __default.INP();
    } else if (((s).dtor_activeTurn()).is_Some()) {
      return __default.GEN();
    } else {
      return __default.IDLE();
    }
  }
  public static int summaryStatus(ChatState s, boolean isError)
  {
    return (int) (((int) (((s).dtor_status()) & ((int) ((~(__default.ACT())))))) | (__default.activityBits(s, isError)));
  }
  public static Turn appendPart(Turn t, Part p)
  {
    Turn _0_dt__update__tmp_h0 = t;
    dafny.DafnySequence<? extends Part> _1_dt__update_hparts_h0 = dafny.DafnySequence.<Part>concatenate((t).dtor_parts(), dafny.DafnySequence.<Part> of(Part._typeDescriptor(), p));
    return Turn.create((_0_dt__update__tmp_h0).dtor_turnId(), (_0_dt__update__tmp_h0).dtor_message(), _1_dt__update_hparts_h0, (_0_dt__update__tmp_h0).dtor_state(), (_0_dt__update__tmp_h0).dtor_usage(), (_0_dt__update__tmp_h0).dtor_error());
  }
  public static dafny.DafnySequence<? extends dafny.CodePoint> partKey(Part p) {
    Part _source0 = p;
    if (_source0.is_PMarkdown()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _0___mcc_h0 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._id;
      dafny.DafnySequence<? extends dafny.CodePoint> _1___mcc_h1 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._content;
      dafny.DafnySequence<? extends dafny.CodePoint> _2_pid = _0___mcc_h0;
      return _2_pid;
    } else if (_source0.is_PReasoning()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _3___mcc_h2 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._id;
      dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h3 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._content;
      dafny.DafnySequence<? extends dafny.CodePoint> _5_pid = _3___mcc_h2;
      return _5_pid;
    } else if (_source0.is_PToolCall()) {
      ToolCall _6___mcc_h4 = ((agency.open.ahp.Chat.Part_PToolCall)_source0)._tc;
      ToolCall _7_tc = _6___mcc_h4;
      return (_7_tc).dtor_toolCallId();
    } else {
      InputReq _8___mcc_h5 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._req;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _9___mcc_h6 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._response;
      InputReq _10_req = _8___mcc_h5;
      return (_10_req).dtor_id();
    }
  }
  public static dafny.DafnySequence<? extends Part> deltaMarkdown(dafny.DafnySequence<? extends Part> ps, dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> content)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Part>SeqUpdateAllWhere(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Part._typeDescriptor(), __default::partKey, ((java.util.function.Function<Part, Boolean>)(_0_p_boxed0) -> {
      Part _0_p = ((Part)(java.lang.Object)(_0_p_boxed0));
      return (_0_p).is_PMarkdown();
    }), ps, id, ((java.util.function.Function<dafny.DafnySequence<? extends dafny.CodePoint>, java.util.function.Function<Part, Part>>)(_1_content) -> {return ((java.util.function.Function<Part, Part>)(_2_p_boxed0) -> {
      Part _2_p = ((Part)(java.lang.Object)(_2_p_boxed0));
      return ((java.util.function.Function<Part, Part>)(_source0_boxed0) -> {
        Part _source0 = ((Part)(java.lang.Object)(_source0_boxed0));
        if (_source0.is_PMarkdown()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _3___mcc_h0 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h1 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._content;
          dafny.DafnySequence<? extends dafny.CodePoint> _5_c = _4___mcc_h1;
          dafny.DafnySequence<? extends dafny.CodePoint> _6_pid = _3___mcc_h0;
          return Part.create_PMarkdown(_6_pid, dafny.DafnySequence.<dafny.CodePoint>concatenate(_5_c, _1_content));
        } else if (_source0.is_PReasoning()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _7___mcc_h4 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _8___mcc_h5 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._content;
          return _2_p;
        } else if (_source0.is_PToolCall()) {
          ToolCall _9___mcc_h8 = ((agency.open.ahp.Chat.Part_PToolCall)_source0)._tc;
          return _2_p;
        } else {
          InputReq _10___mcc_h10 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._req;
          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _11___mcc_h11 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._response;
          return _2_p;
        }
      }).apply(_2_p);
    });}).apply(content));
  }
  public static dafny.DafnySequence<? extends Part> deltaReasoning(dafny.DafnySequence<? extends Part> ps, dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> content)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Part>SeqUpdateAllWhere(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Part._typeDescriptor(), __default::partKey, ((java.util.function.Function<Part, Boolean>)(_0_p_boxed0) -> {
      Part _0_p = ((Part)(java.lang.Object)(_0_p_boxed0));
      return (_0_p).is_PReasoning();
    }), ps, id, ((java.util.function.Function<dafny.DafnySequence<? extends dafny.CodePoint>, java.util.function.Function<Part, Part>>)(_1_content) -> {return ((java.util.function.Function<Part, Part>)(_2_p_boxed0) -> {
      Part _2_p = ((Part)(java.lang.Object)(_2_p_boxed0));
      return ((java.util.function.Function<Part, Part>)(_source0_boxed0) -> {
        Part _source0 = ((Part)(java.lang.Object)(_source0_boxed0));
        if (_source0.is_PMarkdown()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _3___mcc_h0 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h1 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._content;
          return _2_p;
        } else if (_source0.is_PReasoning()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _5___mcc_h4 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _6___mcc_h5 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._content;
          dafny.DafnySequence<? extends dafny.CodePoint> _7_c = _6___mcc_h5;
          dafny.DafnySequence<? extends dafny.CodePoint> _8_pid = _5___mcc_h4;
          return Part.create_PReasoning(_8_pid, dafny.DafnySequence.<dafny.CodePoint>concatenate(_7_c, _1_content));
        } else if (_source0.is_PToolCall()) {
          ToolCall _9___mcc_h8 = ((agency.open.ahp.Chat.Part_PToolCall)_source0)._tc;
          return _2_p;
        } else {
          InputReq _10___mcc_h10 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._req;
          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _11___mcc_h11 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._response;
          return _2_p;
        }
      }).apply(_2_p);
    });}).apply(content));
  }
  public static dafny.DafnySequence<? extends Part> mapTC(dafny.DafnySequence<? extends Part> ps, dafny.DafnySequence<? extends dafny.CodePoint> id, java.util.function.Function<ToolCall, ToolCall> f)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Part>SeqUpdateAllWhere(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Part._typeDescriptor(), __default::partKey, ((java.util.function.Function<Part, Boolean>)(_0_p_boxed0) -> {
      Part _0_p = ((Part)(java.lang.Object)(_0_p_boxed0));
      return (_0_p).is_PToolCall();
    }), ps, id, ((java.util.function.Function<java.util.function.Function<ToolCall, ToolCall>, java.util.function.Function<Part, Part>>)(_1_f) -> {return ((java.util.function.Function<Part, Part>)(_2_p_boxed0) -> {
      Part _2_p = ((Part)(java.lang.Object)(_2_p_boxed0));
      return ((java.util.function.Function<Part, Part>)(_source0_boxed0) -> {
        Part _source0 = ((Part)(java.lang.Object)(_source0_boxed0));
        if (_source0.is_PMarkdown()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _3___mcc_h0 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h1 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._content;
          return _2_p;
        } else if (_source0.is_PReasoning()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _5___mcc_h4 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _6___mcc_h5 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._content;
          return _2_p;
        } else if (_source0.is_PToolCall()) {
          ToolCall _7___mcc_h8 = ((agency.open.ahp.Chat.Part_PToolCall)_source0)._tc;
          ToolCall _8_tc = _7___mcc_h8;
          return Part.create_PToolCall(((ToolCall)(java.lang.Object)((_1_f).apply(_8_tc))));
        } else {
          InputReq _9___mcc_h10 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._req;
          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _10___mcc_h11 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._response;
          return _2_p;
        }
      }).apply(_2_p);
    });}).apply(f));
  }
  public static agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> metaOr(agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> prior)
  {
    if ((meta).is_Some()) {
      return meta;
    } else {
      return prior;
    }
  }
  public static dafny.DafnySequence<? extends dafny.CodePoint> optionId(agency.open.ahp.ConfluxCodec.Json j) {
    if ((((j).is_JObj()) && (((j).dtor_fields()).<dafny.DafnySequence<? extends dafny.CodePoint>>contains(dafny.DafnySequence.asUnicodeString("id")))) && ((((agency.open.ahp.ConfluxCodec.Json)(java.lang.Object)(((j).dtor_fields()).get(dafny.DafnySequence.asUnicodeString("id"))))).is_JStr())) {
      return (((agency.open.ahp.ConfluxCodec.Json)(java.lang.Object)(((j).dtor_fields()).get(dafny.DafnySequence.asUnicodeString("id"))))).dtor_s();
    } else {
      return dafny.DafnySequence.asUnicodeString("");
    }
  }
  public static agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> selectOption(agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> options, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> id)
  {
    TAIL_CALL_START: while (true) {
      if ((((options).is_None()) || ((id).is_None())) || ((java.math.BigInteger.valueOf(((options).dtor_value()).length())).signum() == 0)) {
        return agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      } else if ((__default.optionId(((agency.open.ahp.ConfluxCodec.Json)(java.lang.Object)(((options).dtor_value()).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO))))))).equals((id).dtor_value())) {
        return __default.jNoNull(((agency.open.ahp.ConfluxCodec.Json)(java.lang.Object)(((options).dtor_value()).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO))))));
      } else {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _in0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_Some(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), ((options).dtor_value()).drop(java.math.BigInteger.ONE));
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _in1 = id;
        options = _in0;
        id = _in1;
        continue TAIL_CALL_START;
      }
    }
  }
  public static ToolCall deltaOne(ToolCall tc, dafny.DafnySequence<? extends dafny.CodePoint> content, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> inv, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta)
  {
    if (!((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("streaming"))) {
      return tc;
    } else {
      ToolCall _0_dt__update__tmp_h0 = tc;
      dafny.DafnySequence<? extends dafny.CodePoint> _1_dt__update_hinvocationMessage_h0 = __default.optOr(inv, (tc).dtor_invocationMessage());
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _2_dt__update_hpartialInput_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.<dafny.CodePoint>concatenate(__default.optOr((tc).dtor_partialInput(), dafny.DafnySequence.asUnicodeString("")), content));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _3_dt__update_hmeta_h0 = __default.metaOr(meta, (tc).dtor_meta());
      return ToolCall.create((_0_dt__update__tmp_h0).dtor_toolCallId(), (_0_dt__update__tmp_h0).dtor_toolName(), (_0_dt__update__tmp_h0).dtor_displayName(), (_0_dt__update__tmp_h0).dtor_status(), (_0_dt__update__tmp_h0).dtor_intention(), (_0_dt__update__tmp_h0).dtor_contributor(), _3_dt__update_hmeta_h0, _1_dt__update_hinvocationMessage_h0, (_0_dt__update__tmp_h0).dtor_toolInput(), (_0_dt__update__tmp_h0).dtor_confirmationTitle(), (_0_dt__update__tmp_h0).dtor_riskAssessment(), (_0_dt__update__tmp_h0).dtor_edits(), (_0_dt__update__tmp_h0).dtor_editable(), (_0_dt__update__tmp_h0).dtor_options(), (_0_dt__update__tmp_h0).dtor_confirmed(), (_0_dt__update__tmp_h0).dtor_success(), (_0_dt__update__tmp_h0).dtor_pastTenseMessage(), (_0_dt__update__tmp_h0).dtor_reason(), (_0_dt__update__tmp_h0).dtor_reasonMessage(), (_0_dt__update__tmp_h0).dtor_userSuggestion(), (_0_dt__update__tmp_h0).dtor_selectedOption(), (_0_dt__update__tmp_h0).dtor_content(), (_0_dt__update__tmp_h0).dtor_structuredContent(), (_0_dt__update__tmp_h0).dtor_error(), (_0_dt__update__tmp_h0).dtor_auth(), _2_dt__update_hpartialInput_h0);
    }
  }
  public static ToolCall readyOne(ToolCall tc, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> inv, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> input, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> title, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> risk, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> edits, agency.open.ahp.AhpSkeleton.Option<Boolean> editable, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> options, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> confirmed, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta)
  {
    if (((!((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("streaming"))) && (!((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("running")))) && (!((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("pending-confirmation")))) {
      return tc;
    } else if ((confirmed).is_Some()) {
      ToolCall _0_dt__update__tmp_h0 = tc;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _1_dt__update_hpartialInput_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _2_dt__update_hauth_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _3_dt__update_herror_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _4_dt__update_hstructuredContent_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _5_dt__update_hcontent_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _6_dt__update_hselectedOption_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _7_dt__update_huserSuggestion_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _8_dt__update_hreasonMessage_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _9_dt__update_hreason_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _10_dt__update_hpastTenseMessage_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _11_dt__update_hsuccess_h0 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _12_dt__update_hconfirmed_h0 = confirmed;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _13_dt__update_hoptions_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _14_dt__update_heditable_h0 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _15_dt__update_hedits_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _16_dt__update_hriskAssessment_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _17_dt__update_hconfirmationTitle_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _18_dt__update_htoolInput_h0 = input;
      dafny.DafnySequence<? extends dafny.CodePoint> _19_dt__update_hinvocationMessage_h0 = __default.optOr(inv, dafny.DafnySequence.asUnicodeString(""));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _20_dt__update_hmeta_h0 = __default.metaOr(meta, (tc).dtor_meta());
      dafny.DafnySequence<? extends dafny.CodePoint> _21_dt__update_hstatus_h0 = dafny.DafnySequence.asUnicodeString("running");
      return ToolCall.create((_0_dt__update__tmp_h0).dtor_toolCallId(), (_0_dt__update__tmp_h0).dtor_toolName(), (_0_dt__update__tmp_h0).dtor_displayName(), _21_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).dtor_intention(), (_0_dt__update__tmp_h0).dtor_contributor(), _20_dt__update_hmeta_h0, _19_dt__update_hinvocationMessage_h0, _18_dt__update_htoolInput_h0, _17_dt__update_hconfirmationTitle_h0, _16_dt__update_hriskAssessment_h0, _15_dt__update_hedits_h0, _14_dt__update_heditable_h0, _13_dt__update_hoptions_h0, _12_dt__update_hconfirmed_h0, _11_dt__update_hsuccess_h0, _10_dt__update_hpastTenseMessage_h0, _9_dt__update_hreason_h0, _8_dt__update_hreasonMessage_h0, _7_dt__update_huserSuggestion_h0, _6_dt__update_hselectedOption_h0, _5_dt__update_hcontent_h0, _4_dt__update_hstructuredContent_h0, _3_dt__update_herror_h0, _2_dt__update_hauth_h0, _1_dt__update_hpartialInput_h0);
    } else {
      boolean _22_pending = ((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("pending-confirmation"));
      ToolCall _23_dt__update__tmp_h1 = tc;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _24_dt__update_hpartialInput_h1 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _25_dt__update_hauth_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _26_dt__update_herror_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _27_dt__update_hstructuredContent_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _28_dt__update_hcontent_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _29_dt__update_hselectedOption_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _30_dt__update_huserSuggestion_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _31_dt__update_hreasonMessage_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _32_dt__update_hreason_h1 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _33_dt__update_hpastTenseMessage_h1 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _34_dt__update_hsuccess_h1 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _35_dt__update_hconfirmed_h1 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _36_dt__update_hoptions_h1 = ((((options).is_Some()) || (!(_22_pending))) ? (options) : ((tc).dtor_options()));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _37_dt__update_heditable_h1 = ((((editable).is_Some()) || (!(_22_pending))) ? (editable) : ((tc).dtor_editable()));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _38_dt__update_hedits_h1 = ((((edits).is_Some()) || (!(_22_pending))) ? (edits) : ((tc).dtor_edits()));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _39_dt__update_hriskAssessment_h1 = ((((risk).is_Some()) || (!(_22_pending))) ? (risk) : ((tc).dtor_riskAssessment()));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _40_dt__update_hconfirmationTitle_h1 = ((((title).is_Some()) || (!(_22_pending))) ? (title) : ((tc).dtor_confirmationTitle()));
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _41_dt__update_htoolInput_h1 = ((((input).is_Some()) || (!(_22_pending))) ? (input) : ((tc).dtor_toolInput()));
      dafny.DafnySequence<? extends dafny.CodePoint> _42_dt__update_hinvocationMessage_h1 = __default.optOr(inv, dafny.DafnySequence.asUnicodeString(""));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _43_dt__update_hmeta_h1 = __default.metaOr(meta, (tc).dtor_meta());
      dafny.DafnySequence<? extends dafny.CodePoint> _44_dt__update_hstatus_h1 = dafny.DafnySequence.asUnicodeString("pending-confirmation");
      return ToolCall.create((_23_dt__update__tmp_h1).dtor_toolCallId(), (_23_dt__update__tmp_h1).dtor_toolName(), (_23_dt__update__tmp_h1).dtor_displayName(), _44_dt__update_hstatus_h1, (_23_dt__update__tmp_h1).dtor_intention(), (_23_dt__update__tmp_h1).dtor_contributor(), _43_dt__update_hmeta_h1, _42_dt__update_hinvocationMessage_h1, _41_dt__update_htoolInput_h1, _40_dt__update_hconfirmationTitle_h1, _39_dt__update_hriskAssessment_h1, _38_dt__update_hedits_h1, _37_dt__update_heditable_h1, _36_dt__update_hoptions_h1, _35_dt__update_hconfirmed_h1, _34_dt__update_hsuccess_h1, _33_dt__update_hpastTenseMessage_h1, _32_dt__update_hreason_h1, _31_dt__update_hreasonMessage_h1, _30_dt__update_huserSuggestion_h1, _29_dt__update_hselectedOption_h1, _28_dt__update_hcontent_h1, _27_dt__update_hstructuredContent_h1, _26_dt__update_herror_h1, _25_dt__update_hauth_h1, _24_dt__update_hpartialInput_h1);
    }
  }
  public static ToolCall authRequiredOne(ToolCall tc, agency.open.ahp.ConfluxCodec.Json challenge, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta)
  {
    if (((!((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("running"))) || (((tc).dtor_contributor()).is_None())) || (!((((tc).dtor_contributor()).dtor_value()).dtor_kind()).equals(dafny.DafnySequence.asUnicodeString("mcp")))) {
      return tc;
    } else {
      ToolCall _0_dt__update__tmp_h0 = tc;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _1_dt__update_hpartialInput_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _2_dt__update_hauth_h0 = __default.jNoNull(challenge);
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _3_dt__update_herror_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _4_dt__update_hstructuredContent_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _5_dt__update_huserSuggestion_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _6_dt__update_hreasonMessage_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _7_dt__update_hreason_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _8_dt__update_hpastTenseMessage_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _9_dt__update_hsuccess_h0 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _10_dt__update_hoptions_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _11_dt__update_heditable_h0 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _12_dt__update_hedits_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _13_dt__update_hriskAssessment_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _14_dt__update_hconfirmationTitle_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _15_dt__update_hmeta_h0 = __default.metaOr(meta, (tc).dtor_meta());
      dafny.DafnySequence<? extends dafny.CodePoint> _16_dt__update_hstatus_h0 = dafny.DafnySequence.asUnicodeString("auth-required");
      return ToolCall.create((_0_dt__update__tmp_h0).dtor_toolCallId(), (_0_dt__update__tmp_h0).dtor_toolName(), (_0_dt__update__tmp_h0).dtor_displayName(), _16_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).dtor_intention(), (_0_dt__update__tmp_h0).dtor_contributor(), _15_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).dtor_invocationMessage(), (_0_dt__update__tmp_h0).dtor_toolInput(), _14_dt__update_hconfirmationTitle_h0, _13_dt__update_hriskAssessment_h0, _12_dt__update_hedits_h0, _11_dt__update_heditable_h0, _10_dt__update_hoptions_h0, (_0_dt__update__tmp_h0).dtor_confirmed(), _9_dt__update_hsuccess_h0, _8_dt__update_hpastTenseMessage_h0, _7_dt__update_hreason_h0, _6_dt__update_hreasonMessage_h0, _5_dt__update_huserSuggestion_h0, (_0_dt__update__tmp_h0).dtor_selectedOption(), (_0_dt__update__tmp_h0).dtor_content(), _4_dt__update_hstructuredContent_h0, _3_dt__update_herror_h0, _2_dt__update_hauth_h0, _1_dt__update_hpartialInput_h0);
    }
  }
  public static ToolCall authResolvedOne(ToolCall tc, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta)
  {
    if (!((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("auth-required"))) {
      return tc;
    } else {
      ToolCall _0_dt__update__tmp_h0 = tc;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _1_dt__update_hauth_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _2_dt__update_hmeta_h0 = __default.metaOr(meta, (tc).dtor_meta());
      dafny.DafnySequence<? extends dafny.CodePoint> _3_dt__update_hstatus_h0 = dafny.DafnySequence.asUnicodeString("running");
      return ToolCall.create((_0_dt__update__tmp_h0).dtor_toolCallId(), (_0_dt__update__tmp_h0).dtor_toolName(), (_0_dt__update__tmp_h0).dtor_displayName(), _3_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).dtor_intention(), (_0_dt__update__tmp_h0).dtor_contributor(), _2_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).dtor_invocationMessage(), (_0_dt__update__tmp_h0).dtor_toolInput(), (_0_dt__update__tmp_h0).dtor_confirmationTitle(), (_0_dt__update__tmp_h0).dtor_riskAssessment(), (_0_dt__update__tmp_h0).dtor_edits(), (_0_dt__update__tmp_h0).dtor_editable(), (_0_dt__update__tmp_h0).dtor_options(), (_0_dt__update__tmp_h0).dtor_confirmed(), (_0_dt__update__tmp_h0).dtor_success(), (_0_dt__update__tmp_h0).dtor_pastTenseMessage(), (_0_dt__update__tmp_h0).dtor_reason(), (_0_dt__update__tmp_h0).dtor_reasonMessage(), (_0_dt__update__tmp_h0).dtor_userSuggestion(), (_0_dt__update__tmp_h0).dtor_selectedOption(), (_0_dt__update__tmp_h0).dtor_content(), (_0_dt__update__tmp_h0).dtor_structuredContent(), (_0_dt__update__tmp_h0).dtor_error(), _1_dt__update_hauth_h0, (_0_dt__update__tmp_h0).dtor_partialInput());
    }
  }
  public static ToolCall confirmOne(ToolCall tc, boolean approved, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> confirmedVal, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> reason, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> reasonMessage, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> userSuggestion, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> editedInput, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> selectedOptionId, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta)
  {
    if (!((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("pending-confirmation"))) {
      return tc;
    } else if (approved) {
      ToolCall _0_dt__update__tmp_h0 = tc;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _1_dt__update_hauth_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _2_dt__update_herror_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _3_dt__update_hstructuredContent_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _4_dt__update_hcontent_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _5_dt__update_huserSuggestion_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _6_dt__update_hreasonMessage_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _7_dt__update_hreason_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _8_dt__update_hpastTenseMessage_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _9_dt__update_hsuccess_h0 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _10_dt__update_hpartialInput_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _11_dt__update_hselectedOption_h0 = __default.selectOption((tc).dtor_options(), selectedOptionId);
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _12_dt__update_hconfirmed_h0 = confirmedVal;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _13_dt__update_hoptions_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _14_dt__update_heditable_h0 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _15_dt__update_hedits_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _16_dt__update_hriskAssessment_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _17_dt__update_hconfirmationTitle_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _18_dt__update_htoolInput_h0 = (((editedInput).is_Some()) ? (editedInput) : ((tc).dtor_toolInput()));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _19_dt__update_hmeta_h0 = __default.metaOr(meta, (tc).dtor_meta());
      dafny.DafnySequence<? extends dafny.CodePoint> _20_dt__update_hstatus_h0 = dafny.DafnySequence.asUnicodeString("running");
      return ToolCall.create((_0_dt__update__tmp_h0).dtor_toolCallId(), (_0_dt__update__tmp_h0).dtor_toolName(), (_0_dt__update__tmp_h0).dtor_displayName(), _20_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).dtor_intention(), (_0_dt__update__tmp_h0).dtor_contributor(), _19_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).dtor_invocationMessage(), _18_dt__update_htoolInput_h0, _17_dt__update_hconfirmationTitle_h0, _16_dt__update_hriskAssessment_h0, _15_dt__update_hedits_h0, _14_dt__update_heditable_h0, _13_dt__update_hoptions_h0, _12_dt__update_hconfirmed_h0, _9_dt__update_hsuccess_h0, _8_dt__update_hpastTenseMessage_h0, _7_dt__update_hreason_h0, _6_dt__update_hreasonMessage_h0, _5_dt__update_huserSuggestion_h0, _11_dt__update_hselectedOption_h0, _4_dt__update_hcontent_h0, _3_dt__update_hstructuredContent_h0, _2_dt__update_herror_h0, _1_dt__update_hauth_h0, _10_dt__update_hpartialInput_h0);
    } else {
      ToolCall _21_dt__update__tmp_h1 = tc;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _22_dt__update_hpartialInput_h1 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _23_dt__update_hauth_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _24_dt__update_herror_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _25_dt__update_hstructuredContent_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _26_dt__update_hcontent_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _27_dt__update_hselectedOption_h1 = __default.selectOption((tc).dtor_options(), selectedOptionId);
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _28_dt__update_huserSuggestion_h1 = userSuggestion;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _29_dt__update_hreasonMessage_h1 = reasonMessage;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _30_dt__update_hreason_h1 = reason;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _31_dt__update_hpastTenseMessage_h1 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _32_dt__update_hsuccess_h1 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _33_dt__update_hconfirmed_h1 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _34_dt__update_hoptions_h1 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _35_dt__update_heditable_h1 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _36_dt__update_hedits_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _37_dt__update_hriskAssessment_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _38_dt__update_hconfirmationTitle_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _39_dt__update_hmeta_h1 = __default.metaOr(meta, (tc).dtor_meta());
      dafny.DafnySequence<? extends dafny.CodePoint> _40_dt__update_hstatus_h1 = dafny.DafnySequence.asUnicodeString("cancelled");
      return ToolCall.create((_21_dt__update__tmp_h1).dtor_toolCallId(), (_21_dt__update__tmp_h1).dtor_toolName(), (_21_dt__update__tmp_h1).dtor_displayName(), _40_dt__update_hstatus_h1, (_21_dt__update__tmp_h1).dtor_intention(), (_21_dt__update__tmp_h1).dtor_contributor(), _39_dt__update_hmeta_h1, (_21_dt__update__tmp_h1).dtor_invocationMessage(), (_21_dt__update__tmp_h1).dtor_toolInput(), _38_dt__update_hconfirmationTitle_h1, _37_dt__update_hriskAssessment_h1, _36_dt__update_hedits_h1, _35_dt__update_heditable_h1, _34_dt__update_hoptions_h1, _33_dt__update_hconfirmed_h1, _32_dt__update_hsuccess_h1, _31_dt__update_hpastTenseMessage_h1, _30_dt__update_hreason_h1, _29_dt__update_hreasonMessage_h1, _28_dt__update_huserSuggestion_h1, _27_dt__update_hselectedOption_h1, _26_dt__update_hcontent_h1, _25_dt__update_hstructuredContent_h1, _24_dt__update_herror_h1, _23_dt__update_hauth_h1, _22_dt__update_hpartialInput_h1);
    }
  }
  public static ToolCall completeOne(ToolCall tc, boolean ok, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> past, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> resultContent, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> structured, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> err, boolean requiresResultConfirmation, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta)
  {
    if (((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("auth-required"))) {
      if (ok) {
        return tc;
      } else {
        ToolCall _0_dt__update__tmp_h0 = tc;
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _1_dt__update_hpartialInput_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _2_dt__update_hauth_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _3_dt__update_herror_h0 = err;
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _4_dt__update_hstructuredContent_h0 = structured;
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _5_dt__update_hcontent_h0 = (((resultContent).is_Some()) ? (resultContent) : ((tc).dtor_content()));
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _6_dt__update_huserSuggestion_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _7_dt__update_hreasonMessage_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _8_dt__update_hreason_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _9_dt__update_hpastTenseMessage_h0 = past;
        agency.open.ahp.AhpSkeleton.Option<Boolean> _10_dt__update_hsuccess_h0 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, false);
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _11_dt__update_hoptions_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
        agency.open.ahp.AhpSkeleton.Option<Boolean> _12_dt__update_heditable_h0 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _13_dt__update_hedits_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _14_dt__update_hriskAssessment_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _15_dt__update_hconfirmationTitle_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _16_dt__update_hmeta_h0 = __default.metaOr(meta, (tc).dtor_meta());
        dafny.DafnySequence<? extends dafny.CodePoint> _17_dt__update_hstatus_h0 = dafny.DafnySequence.asUnicodeString("completed");
        return ToolCall.create((_0_dt__update__tmp_h0).dtor_toolCallId(), (_0_dt__update__tmp_h0).dtor_toolName(), (_0_dt__update__tmp_h0).dtor_displayName(), _17_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).dtor_intention(), (_0_dt__update__tmp_h0).dtor_contributor(), _16_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).dtor_invocationMessage(), (_0_dt__update__tmp_h0).dtor_toolInput(), _15_dt__update_hconfirmationTitle_h0, _14_dt__update_hriskAssessment_h0, _13_dt__update_hedits_h0, _12_dt__update_heditable_h0, _11_dt__update_hoptions_h0, (_0_dt__update__tmp_h0).dtor_confirmed(), _10_dt__update_hsuccess_h0, _9_dt__update_hpastTenseMessage_h0, _8_dt__update_hreason_h0, _7_dt__update_hreasonMessage_h0, _6_dt__update_huserSuggestion_h0, (_0_dt__update__tmp_h0).dtor_selectedOption(), _5_dt__update_hcontent_h0, _4_dt__update_hstructuredContent_h0, _3_dt__update_herror_h0, _2_dt__update_hauth_h0, _1_dt__update_hpartialInput_h0);
      }
    } else if ((!((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("running"))) && (!((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("pending-confirmation")))) {
      return tc;
    } else {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _18_conf = ((((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("running"))) ? ((tc).dtor_confirmed()) : (agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("not-needed"))));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _19_selected = ((((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("running"))) ? ((tc).dtor_selectedOption()) : (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())));
      dafny.DafnySequence<? extends dafny.CodePoint> _20_newStatus = ((requiresResultConfirmation) ? (dafny.DafnySequence.asUnicodeString("pending-result-confirmation")) : (dafny.DafnySequence.asUnicodeString("completed")));
      ToolCall _21_dt__update__tmp_h1 = tc;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _22_dt__update_hpartialInput_h1 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _23_dt__update_hauth_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _24_dt__update_herror_h1 = err;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _25_dt__update_hstructuredContent_h1 = structured;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _26_dt__update_hcontent_h1 = resultContent;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _27_dt__update_huserSuggestion_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _28_dt__update_hreasonMessage_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _29_dt__update_hreason_h1 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _30_dt__update_hpastTenseMessage_h1 = past;
      agency.open.ahp.AhpSkeleton.Option<Boolean> _31_dt__update_hsuccess_h1 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, ok);
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _32_dt__update_hselectedOption_h0 = _19_selected;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _33_dt__update_hconfirmed_h0 = _18_conf;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _34_dt__update_hoptions_h1 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _35_dt__update_heditable_h1 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _36_dt__update_hedits_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _37_dt__update_hriskAssessment_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _38_dt__update_hconfirmationTitle_h1 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _39_dt__update_hmeta_h1 = __default.metaOr(meta, (tc).dtor_meta());
      dafny.DafnySequence<? extends dafny.CodePoint> _40_dt__update_hstatus_h1 = _20_newStatus;
      return ToolCall.create((_21_dt__update__tmp_h1).dtor_toolCallId(), (_21_dt__update__tmp_h1).dtor_toolName(), (_21_dt__update__tmp_h1).dtor_displayName(), _40_dt__update_hstatus_h1, (_21_dt__update__tmp_h1).dtor_intention(), (_21_dt__update__tmp_h1).dtor_contributor(), _39_dt__update_hmeta_h1, (_21_dt__update__tmp_h1).dtor_invocationMessage(), (_21_dt__update__tmp_h1).dtor_toolInput(), _38_dt__update_hconfirmationTitle_h1, _37_dt__update_hriskAssessment_h1, _36_dt__update_hedits_h1, _35_dt__update_heditable_h1, _34_dt__update_hoptions_h1, _33_dt__update_hconfirmed_h0, _31_dt__update_hsuccess_h1, _30_dt__update_hpastTenseMessage_h1, _29_dt__update_hreason_h1, _28_dt__update_hreasonMessage_h1, _27_dt__update_huserSuggestion_h1, _32_dt__update_hselectedOption_h0, _26_dt__update_hcontent_h1, _25_dt__update_hstructuredContent_h1, _24_dt__update_herror_h1, _23_dt__update_hauth_h1, _22_dt__update_hpartialInput_h1);
    }
  }
  public static ToolCall resultConfirmOne(ToolCall tc, boolean approved, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta)
  {
    if (!((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("pending-result-confirmation"))) {
      return tc;
    } else if (approved) {
      ToolCall _0_dt__update__tmp_h0 = tc;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _1_dt__update_hmeta_h0 = __default.metaOr(meta, (tc).dtor_meta());
      dafny.DafnySequence<? extends dafny.CodePoint> _2_dt__update_hstatus_h0 = dafny.DafnySequence.asUnicodeString("completed");
      return ToolCall.create((_0_dt__update__tmp_h0).dtor_toolCallId(), (_0_dt__update__tmp_h0).dtor_toolName(), (_0_dt__update__tmp_h0).dtor_displayName(), _2_dt__update_hstatus_h0, (_0_dt__update__tmp_h0).dtor_intention(), (_0_dt__update__tmp_h0).dtor_contributor(), _1_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).dtor_invocationMessage(), (_0_dt__update__tmp_h0).dtor_toolInput(), (_0_dt__update__tmp_h0).dtor_confirmationTitle(), (_0_dt__update__tmp_h0).dtor_riskAssessment(), (_0_dt__update__tmp_h0).dtor_edits(), (_0_dt__update__tmp_h0).dtor_editable(), (_0_dt__update__tmp_h0).dtor_options(), (_0_dt__update__tmp_h0).dtor_confirmed(), (_0_dt__update__tmp_h0).dtor_success(), (_0_dt__update__tmp_h0).dtor_pastTenseMessage(), (_0_dt__update__tmp_h0).dtor_reason(), (_0_dt__update__tmp_h0).dtor_reasonMessage(), (_0_dt__update__tmp_h0).dtor_userSuggestion(), (_0_dt__update__tmp_h0).dtor_selectedOption(), (_0_dt__update__tmp_h0).dtor_content(), (_0_dt__update__tmp_h0).dtor_structuredContent(), (_0_dt__update__tmp_h0).dtor_error(), (_0_dt__update__tmp_h0).dtor_auth(), (_0_dt__update__tmp_h0).dtor_partialInput());
    } else {
      ToolCall _3_dt__update__tmp_h1 = tc;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _4_dt__update_herror_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _5_dt__update_hstructuredContent_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _6_dt__update_hcontent_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor());
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _7_dt__update_hpastTenseMessage_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<Boolean> _8_dt__update_hsuccess_h0 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN);
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _9_dt__update_hconfirmed_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR));
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _10_dt__update_hreason_h0 = agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("result-denied"));
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _11_dt__update_hmeta_h1 = __default.metaOr(meta, (tc).dtor_meta());
      dafny.DafnySequence<? extends dafny.CodePoint> _12_dt__update_hstatus_h1 = dafny.DafnySequence.asUnicodeString("cancelled");
      return ToolCall.create((_3_dt__update__tmp_h1).dtor_toolCallId(), (_3_dt__update__tmp_h1).dtor_toolName(), (_3_dt__update__tmp_h1).dtor_displayName(), _12_dt__update_hstatus_h1, (_3_dt__update__tmp_h1).dtor_intention(), (_3_dt__update__tmp_h1).dtor_contributor(), _11_dt__update_hmeta_h1, (_3_dt__update__tmp_h1).dtor_invocationMessage(), (_3_dt__update__tmp_h1).dtor_toolInput(), (_3_dt__update__tmp_h1).dtor_confirmationTitle(), (_3_dt__update__tmp_h1).dtor_riskAssessment(), (_3_dt__update__tmp_h1).dtor_edits(), (_3_dt__update__tmp_h1).dtor_editable(), (_3_dt__update__tmp_h1).dtor_options(), _9_dt__update_hconfirmed_h0, _8_dt__update_hsuccess_h0, _7_dt__update_hpastTenseMessage_h0, _10_dt__update_hreason_h0, (_3_dt__update__tmp_h1).dtor_reasonMessage(), (_3_dt__update__tmp_h1).dtor_userSuggestion(), (_3_dt__update__tmp_h1).dtor_selectedOption(), _6_dt__update_hcontent_h0, _5_dt__update_hstructuredContent_h0, _4_dt__update_herror_h0, (_3_dt__update__tmp_h1).dtor_auth(), (_3_dt__update__tmp_h1).dtor_partialInput());
    }
  }
  public static ToolCall contentChangedOne(ToolCall tc, agency.open.ahp.ConfluxCodec.Json c, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> meta)
  {
    if (!((tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("running"))) {
      return tc;
    } else {
      ToolCall _0_dt__update__tmp_h0 = tc;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _1_dt__update_hcontent_h0 = __default.jNoNull(c);
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _2_dt__update_hmeta_h0 = __default.metaOr(meta, (tc).dtor_meta());
      return ToolCall.create((_0_dt__update__tmp_h0).dtor_toolCallId(), (_0_dt__update__tmp_h0).dtor_toolName(), (_0_dt__update__tmp_h0).dtor_displayName(), (_0_dt__update__tmp_h0).dtor_status(), (_0_dt__update__tmp_h0).dtor_intention(), (_0_dt__update__tmp_h0).dtor_contributor(), _2_dt__update_hmeta_h0, (_0_dt__update__tmp_h0).dtor_invocationMessage(), (_0_dt__update__tmp_h0).dtor_toolInput(), (_0_dt__update__tmp_h0).dtor_confirmationTitle(), (_0_dt__update__tmp_h0).dtor_riskAssessment(), (_0_dt__update__tmp_h0).dtor_edits(), (_0_dt__update__tmp_h0).dtor_editable(), (_0_dt__update__tmp_h0).dtor_options(), (_0_dt__update__tmp_h0).dtor_confirmed(), (_0_dt__update__tmp_h0).dtor_success(), (_0_dt__update__tmp_h0).dtor_pastTenseMessage(), (_0_dt__update__tmp_h0).dtor_reason(), (_0_dt__update__tmp_h0).dtor_reasonMessage(), (_0_dt__update__tmp_h0).dtor_userSuggestion(), (_0_dt__update__tmp_h0).dtor_selectedOption(), _1_dt__update_hcontent_h0, (_0_dt__update__tmp_h0).dtor_structuredContent(), (_0_dt__update__tmp_h0).dtor_error(), (_0_dt__update__tmp_h0).dtor_auth(), (_0_dt__update__tmp_h0).dtor_partialInput());
    }
  }
  public static dafny.DafnySequence<? extends Part> forceCancel(dafny.DafnySequence<? extends Part> ps) {
    return agency.open.ahp.ConfluxOperators.__default.<Part, Part>Map(Part._typeDescriptor(), Part._typeDescriptor(), ((java.util.function.Function<Part, Part>)(_0_h_boxed0) -> {
      Part _0_h = ((Part)(java.lang.Object)(_0_h_boxed0));
      return ((java.util.function.Function<Part, Part>)(_source0_boxed0) -> {
        Part _source0 = ((Part)(java.lang.Object)(_source0_boxed0));
        if (_source0.is_PMarkdown()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _1___mcc_h0 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _2___mcc_h1 = ((agency.open.ahp.Chat.Part_PMarkdown)_source0)._content;
          return _0_h;
        } else if (_source0.is_PReasoning()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _3___mcc_h4 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._id;
          dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h5 = ((agency.open.ahp.Chat.Part_PReasoning)_source0)._content;
          return _0_h;
        } else if (_source0.is_PToolCall()) {
          ToolCall _5___mcc_h8 = ((agency.open.ahp.Chat.Part_PToolCall)_source0)._tc;
          ToolCall _6_tc = _5___mcc_h8;
          if ((((_6_tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("completed"))) || (((_6_tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("cancelled")))) {
            return _0_h;
          } else {
            return Part.create_PToolCall(((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(_6_tc, boxed470 -> {
  ToolCall _pat_let235_0 = ((ToolCall)(java.lang.Object)(boxed470));
  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(_pat_let235_0, boxed471 -> {
    ToolCall _7_dt__update__tmp_h0 = ((ToolCall)(java.lang.Object)(boxed471));
    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), boxed472 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let236_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed472));
      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(_pat_let236_0, boxed473 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _8_dt__update_hpartialInput_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed473));
        return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed474 -> {
          agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let237_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed474));
          return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let237_0, boxed475 -> {
            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _9_dt__update_hauth_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed475));
            return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed476 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let238_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed476));
              return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let238_0, boxed477 -> {
                agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _10_dt__update_herror_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed477));
                return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed478 -> {
                  agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let239_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed478));
                  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let239_0, boxed479 -> {
                    agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _11_dt__update_hstructuredContent_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed479));
                    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed480 -> {
                      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let240_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed480));
                      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let240_0, boxed481 -> {
                        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _12_dt__update_hcontent_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed481));
                        return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed482 -> {
                          agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let241_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed482));
                          return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let241_0, boxed483 -> {
                            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _13_dt__update_hselectedOption_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed483));
                            return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed484 -> {
                              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let242_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed484));
                              return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let242_0, boxed485 -> {
                                agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _14_dt__update_huserSuggestion_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed485));
                                return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed486 -> {
                                  agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let243_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed486));
                                  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let243_0, boxed487 -> {
                                    agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _15_dt__update_hreasonMessage_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed487));
                                    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), boxed488 -> {
                                      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let244_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed488));
                                      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(_pat_let244_0, boxed489 -> {
                                        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _16_dt__update_hpastTenseMessage_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed489));
                                        return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), boxed490 -> {
                                          agency.open.ahp.AhpSkeleton.Option<Boolean> _pat_let245_0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed490));
                                          return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, ToolCall>Let(_pat_let245_0, boxed491 -> {
                                            agency.open.ahp.AhpSkeleton.Option<Boolean> _17_dt__update_hsuccess_h0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed491));
                                            return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), boxed492 -> {
                                              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let246_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed492));
                                              return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(_pat_let246_0, boxed493 -> {
                                                agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _18_dt__update_hconfirmed_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed493));
                                                return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), boxed494 -> {
                                                  agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _pat_let247_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>)(java.lang.Object)(boxed494));
                                                  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>, ToolCall>Let(_pat_let247_0, boxed495 -> {
                                                    agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _19_dt__update_hoptions_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>)(java.lang.Object)(boxed495));
                                                    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), boxed496 -> {
                                                      agency.open.ahp.AhpSkeleton.Option<Boolean> _pat_let248_0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed496));
                                                      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, ToolCall>Let(_pat_let248_0, boxed497 -> {
                                                        agency.open.ahp.AhpSkeleton.Option<Boolean> _20_dt__update_heditable_h0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed497));
                                                        return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed498 -> {
                                                          agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let249_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed498));
                                                          return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let249_0, boxed499 -> {
                                                            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _21_dt__update_hedits_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed499));
                                                            return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed500 -> {
                                                              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let250_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed500));
                                                              return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let250_0, boxed501 -> {
                                                                agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _22_dt__update_hriskAssessment_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed501));
                                                                return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), boxed502 -> {
                                                                  agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let251_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed502));
                                                                  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let251_0, boxed503 -> {
                                                                    agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _23_dt__update_hconfirmationTitle_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed503));
                                                                    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(((((_6_tc).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("streaming"))) ? (agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))) : ((_6_tc).dtor_toolInput())), boxed504 -> {
                                                                      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let252_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed504));
                                                                      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(_pat_let252_0, boxed505 -> {
                                                                        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _24_dt__update_htoolInput_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed505));
                                                                        return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("skipped")), boxed506 -> {
                                                                          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let253_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed506));
                                                                          return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(_pat_let253_0, boxed507 -> {
                                                                            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _25_dt__update_hreason_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed507));
                                                                            return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(dafny.DafnySequence.asUnicodeString("cancelled"), boxed508 -> {
                                                                              dafny.DafnySequence<? extends dafny.CodePoint> _pat_let254_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed508));
                                                                              return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(_pat_let254_0, boxed509 -> {
                                                                                dafny.DafnySequence<? extends dafny.CodePoint> _26_dt__update_hstatus_h0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed509));
                                                                                return ToolCall.create((_7_dt__update__tmp_h0).dtor_toolCallId(), (_7_dt__update__tmp_h0).dtor_toolName(), (_7_dt__update__tmp_h0).dtor_displayName(), _26_dt__update_hstatus_h0, (_7_dt__update__tmp_h0).dtor_intention(), (_7_dt__update__tmp_h0).dtor_contributor(), (_7_dt__update__tmp_h0).dtor_meta(), (_7_dt__update__tmp_h0).dtor_invocationMessage(), _24_dt__update_htoolInput_h0, _23_dt__update_hconfirmationTitle_h0, _22_dt__update_hriskAssessment_h0, _21_dt__update_hedits_h0, _20_dt__update_heditable_h0, _19_dt__update_hoptions_h0, _18_dt__update_hconfirmed_h0, _17_dt__update_hsuccess_h0, _16_dt__update_hpastTenseMessage_h0, _25_dt__update_hreason_h0, _15_dt__update_hreasonMessage_h0, _14_dt__update_huserSuggestion_h0, _13_dt__update_hselectedOption_h0, _12_dt__update_hcontent_h0, _11_dt__update_hstructuredContent_h0, _10_dt__update_herror_h0, _9_dt__update_hauth_h0, _8_dt__update_hpartialInput_h0);
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
      )));
    }
    )));
  }
  )));
}
))));
          }
        } else {
          InputReq _27___mcc_h10 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._req;
          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _28___mcc_h11 = ((agency.open.ahp.Chat.Part_PInputRequest)_source0)._response;
          return _0_h;
        }
      }).apply(_0_h);
    }), ps);
  }
  public static dafny.Tuple2<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome> endTurn(ChatState s, dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> turnState, boolean isError, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> err)
  {
    agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let_tv0 = err;
    dafny.DafnySequence<? extends dafny.CodePoint> _pat_let_tv1 = turnState;
    ChatState _pat_let_tv2 = s;
    boolean _pat_let_tv3 = isError;
    if (!(__default.turnMatches(s, id))) {
      return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
    } else {
      Turn _0_t = ((s).dtor_activeTurn()).dtor_value();
      Turn _1_finalized = ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_0_t, boxed510 -> {
        Turn _pat_let255_0 = ((Turn)(java.lang.Object)(boxed510));
        return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let255_0, boxed511 -> {
          Turn _2_dt__update__tmp_h0 = ((Turn)(java.lang.Object)(boxed511));
          return ((Turn)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Turn>Let(_pat_let_tv0, boxed512 -> {
            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let256_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed512));
            return ((Turn)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Turn>Let(_pat_let256_0, boxed513 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _3_dt__update_herror_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed513));
              return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.forceCancel((_0_t).dtor_parts()), boxed514 -> {
                dafny.DafnySequence<? extends Part> _pat_let257_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed514));
                return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let257_0, boxed515 -> {
                  dafny.DafnySequence<? extends Part> _4_dt__update_hparts_h0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed515));
                  return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, Turn>Let(_pat_let_tv1, boxed516 -> {
                    dafny.DafnySequence<? extends dafny.CodePoint> _pat_let258_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed516));
                    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, Turn>Let(_pat_let258_0, boxed517 -> {
                      dafny.DafnySequence<? extends dafny.CodePoint> _5_dt__update_hstate_h0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed517));
                      return Turn.create((_2_dt__update__tmp_h0).dtor_turnId(), (_2_dt__update__tmp_h0).dtor_message(), _4_dt__update_hparts_h0, _5_dt__update_hstate_h0, (_2_dt__update__tmp_h0).dtor_usage(), _3_dt__update_herror_h0);
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
      ChatState _6_next = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed518 -> {
        ChatState _pat_let259_0 = ((ChatState)(java.lang.Object)(boxed518));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let259_0, boxed519 -> {
          ChatState _7_dt__update__tmp_h1 = ((ChatState)(java.lang.Object)(boxed519));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_None(Turn._typeDescriptor()), boxed520 -> {
            agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let260_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed520));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let260_0, boxed521 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _8_dt__update_hactiveTurn_h0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed521));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(dafny.DafnySequence.<Turn>concatenate((_pat_let_tv2).dtor_turns(), dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), _1_finalized)), boxed522 -> {
                dafny.DafnySequence<? extends Turn> _pat_let261_0 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed522));
                return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(_pat_let261_0, boxed523 -> {
                  dafny.DafnySequence<? extends Turn> _9_dt__update_hturns_h0 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed523));
                  return ChatState.create(_9_dt__update_hturns_h0, (_7_dt__update__tmp_h1).dtor_title(), (_7_dt__update__tmp_h1).dtor_status(), (_7_dt__update__tmp_h1).dtor_modifiedAt(), (_7_dt__update__tmp_h1).dtor_draft(), (_7_dt__update__tmp_h1).dtor_activity(), _8_dt__update_hactiveTurn_h0, (_7_dt__update__tmp_h1).dtor_steeringMessage(), (_7_dt__update__tmp_h1).dtor_queuedMessages(), (_7_dt__update__tmp_h1).dtor_nextCursor());
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
      return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_6_next, boxed524 -> {
  ChatState _pat_let262_0 = ((ChatState)(java.lang.Object)(boxed524));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let262_0, boxed525 -> {
    ChatState _10_dt__update__tmp_h2 = ((ChatState)(java.lang.Object)(boxed525));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(__default.summaryStatus(_6_next, _pat_let_tv3), boxed526 -> {
      int _pat_let263_0 = ((int)(java.lang.Object)(boxed526));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let263_0, boxed527 -> {
        int _11_dt__update_hstatus_h0 = ((int)(java.lang.Object)(boxed527));
        return ChatState.create((_10_dt__update__tmp_h2).dtor_turns(), (_10_dt__update__tmp_h2).dtor_title(), _11_dt__update_hstatus_h0, (_10_dt__update__tmp_h2).dtor_modifiedAt(), (_10_dt__update__tmp_h2).dtor_draft(), (_10_dt__update__tmp_h2).dtor_activity(), (_10_dt__update__tmp_h2).dtor_activeTurn(), (_10_dt__update__tmp_h2).dtor_steeringMessage(), (_10_dt__update__tmp_h2).dtor_queuedMessages(), (_10_dt__update__tmp_h2).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    }
  }
  public static dafny.Tuple2<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome> ApplyToChat(ChatState s, ChatAction a, java.math.BigInteger now)
  {
    ChatState _pat_let_tv0 = s;
    ChatState _pat_let_tv1 = s;
    ChatState _pat_let_tv2 = s;
    ChatState _pat_let_tv3 = s;
    ChatState _pat_let_tv4 = s;
    ChatState _pat_let_tv5 = s;
    ChatState _pat_let_tv6 = s;
    ChatState _pat_let_tv7 = s;
    ChatState _pat_let_tv8 = s;
    ChatState _pat_let_tv9 = s;
    ChatState _pat_let_tv10 = s;
    ChatState _pat_let_tv11 = s;
    ChatState _pat_let_tv12 = s;
    ChatState _pat_let_tv13 = s;
    ChatState _pat_let_tv14 = s;
    ChatState _pat_let_tv15 = s;
    ChatState _pat_let_tv16 = s;
    ChatState _pat_let_tv17 = s;
    ChatState _pat_let_tv18 = s;
    ChatState _pat_let_tv19 = s;
    ChatState _pat_let_tv20 = s;
    ChatState _pat_let_tv21 = s;
    ChatState _pat_let_tv22 = s;
    ChatState _pat_let_tv23 = s;
    ChatState _pat_let_tv24 = s;
    ChatState _pat_let_tv25 = s;
    ChatState _pat_let_tv26 = s;
    ChatState _pat_let_tv27 = s;
    ChatState _pat_let_tv28 = s;
    ChatAction _source0 = a;
    if (_source0.is_CDraftChanged()) {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _0___mcc_h0 = ((agency.open.ahp.Chat.ChatAction_CDraftChanged)_source0)._draft;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _1_d = _0___mcc_h0;
      return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed528 -> {
  ChatState _pat_let264_0 = ((ChatState)(java.lang.Object)(boxed528));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let264_0, boxed529 -> {
    ChatState _2_dt__update__tmp_h0 = ((ChatState)(java.lang.Object)(boxed529));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(_1_d, boxed530 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let265_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed530));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(_pat_let265_0, boxed531 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _3_dt__update_hdraft_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed531));
        return ChatState.create((_2_dt__update__tmp_h0).dtor_turns(), (_2_dt__update__tmp_h0).dtor_title(), (_2_dt__update__tmp_h0).dtor_status(), (_2_dt__update__tmp_h0).dtor_modifiedAt(), _3_dt__update_hdraft_h0, (_2_dt__update__tmp_h0).dtor_activity(), (_2_dt__update__tmp_h0).dtor_activeTurn(), (_2_dt__update__tmp_h0).dtor_steeringMessage(), (_2_dt__update__tmp_h0).dtor_queuedMessages(), (_2_dt__update__tmp_h0).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_CActivityChanged()) {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _4___mcc_h1 = ((agency.open.ahp.Chat.ChatAction_CActivityChanged)_source0)._activity;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _5_ac = _4___mcc_h1;
      return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed532 -> {
  ChatState _pat_let266_0 = ((ChatState)(java.lang.Object)(boxed532));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let266_0, boxed533 -> {
    ChatState _6_dt__update__tmp_h1 = ((ChatState)(java.lang.Object)(boxed533));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(_5_ac, boxed534 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let267_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed534));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(_pat_let267_0, boxed535 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _7_dt__update_hactivity_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed535));
        return ChatState.create((_6_dt__update__tmp_h1).dtor_turns(), (_6_dt__update__tmp_h1).dtor_title(), (_6_dt__update__tmp_h1).dtor_status(), (_6_dt__update__tmp_h1).dtor_modifiedAt(), (_6_dt__update__tmp_h1).dtor_draft(), _7_dt__update_hactivity_h0, (_6_dt__update__tmp_h1).dtor_activeTurn(), (_6_dt__update__tmp_h1).dtor_steeringMessage(), (_6_dt__update__tmp_h1).dtor_queuedMessages(), (_6_dt__update__tmp_h1).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_CTurnStarted()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _8___mcc_h2 = ((agency.open.ahp.Chat.ChatAction_CTurnStarted)_source0)._turnId;
      agency.open.ahp.ConfluxCodec.Json _9___mcc_h3 = ((agency.open.ahp.Chat.ChatAction_CTurnStarted)_source0)._message;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _10___mcc_h4 = ((agency.open.ahp.Chat.ChatAction_CTurnStarted)_source0)._queuedMessageId;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _11_qid = _10___mcc_h4;
      agency.open.ahp.ConfluxCodec.Json _12_m = _9___mcc_h3;
      dafny.DafnySequence<? extends dafny.CodePoint> _13_id = _8___mcc_h2;
      ChatState _14_withTurn = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed536 -> {
        ChatState _pat_let268_0 = ((ChatState)(java.lang.Object)(boxed536));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let268_0, boxed537 -> {
          ChatState _15_dt__update__tmp_h2 = ((ChatState)(java.lang.Object)(boxed537));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), Turn.create(_13_id, _12_m, dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), dafny.DafnySequence.asUnicodeString("in-progress"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), boxed538 -> {
            agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let269_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed538));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let269_0, boxed539 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _16_dt__update_hactiveTurn_h0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed539));
              return ChatState.create((_15_dt__update__tmp_h2).dtor_turns(), (_15_dt__update__tmp_h2).dtor_title(), (_15_dt__update__tmp_h2).dtor_status(), (_15_dt__update__tmp_h2).dtor_modifiedAt(), (_15_dt__update__tmp_h2).dtor_draft(), (_15_dt__update__tmp_h2).dtor_activity(), _16_dt__update_hactiveTurn_h0, (_15_dt__update__tmp_h2).dtor_steeringMessage(), (_15_dt__update__tmp_h2).dtor_queuedMessages(), (_15_dt__update__tmp_h2).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      )));
      ChatState _17_s2 = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_14_withTurn, boxed540 -> {
        ChatState _pat_let270_0 = ((ChatState)(java.lang.Object)(boxed540));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let270_0, boxed541 -> {
          ChatState _18_dt__update__tmp_h3 = ((ChatState)(java.lang.Object)(boxed541));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(__default.clearBit(__default.summaryStatus(_14_withTurn, false), __default.READ()), boxed542 -> {
            int _pat_let271_0 = ((int)(java.lang.Object)(boxed542));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let271_0, boxed543 -> {
              int _19_dt__update_hstatus_h0 = ((int)(java.lang.Object)(boxed543));
              return ChatState.create((_18_dt__update__tmp_h3).dtor_turns(), (_18_dt__update__tmp_h3).dtor_title(), _19_dt__update_hstatus_h0, (_18_dt__update__tmp_h3).dtor_modifiedAt(), (_18_dt__update__tmp_h3).dtor_draft(), (_18_dt__update__tmp_h3).dtor_activity(), (_18_dt__update__tmp_h3).dtor_activeTurn(), (_18_dt__update__tmp_h3).dtor_steeringMessage(), (_18_dt__update__tmp_h3).dtor_queuedMessages(), (_18_dt__update__tmp_h3).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      )));
      ChatState _20_s3 = (((((_11_qid).is_Some()) && (((_17_s2).dtor_steeringMessage()).is_Some())) && (((((_17_s2).dtor_steeringMessage()).dtor_value()).dtor_id()).equals((_11_qid).dtor_value()))) ? (((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_17_s2, boxed544 -> {
        ChatState _pat_let272_0 = ((ChatState)(java.lang.Object)(boxed544));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let272_0, boxed545 -> {
          ChatState _21_dt__update__tmp_h4 = ((ChatState)(java.lang.Object)(boxed545));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<QMsg>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<QMsg>create_None(QMsg._typeDescriptor()), boxed546 -> {
            agency.open.ahp.AhpSkeleton.Option<QMsg> _pat_let273_0 = ((agency.open.ahp.AhpSkeleton.Option<QMsg>)(java.lang.Object)(boxed546));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<QMsg>, ChatState>Let(_pat_let273_0, boxed547 -> {
              agency.open.ahp.AhpSkeleton.Option<QMsg> _22_dt__update_hsteeringMessage_h0 = ((agency.open.ahp.AhpSkeleton.Option<QMsg>)(java.lang.Object)(boxed547));
              return ChatState.create((_21_dt__update__tmp_h4).dtor_turns(), (_21_dt__update__tmp_h4).dtor_title(), (_21_dt__update__tmp_h4).dtor_status(), (_21_dt__update__tmp_h4).dtor_modifiedAt(), (_21_dt__update__tmp_h4).dtor_draft(), (_21_dt__update__tmp_h4).dtor_activity(), (_21_dt__update__tmp_h4).dtor_activeTurn(), _22_dt__update_hsteeringMessage_h0, (_21_dt__update__tmp_h4).dtor_queuedMessages(), (_21_dt__update__tmp_h4).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      )))) : (_17_s2));
      ChatState _23_s4 = (((_11_qid).is_Some()) ? (((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_20_s3, boxed548 -> {
        ChatState _pat_let274_0 = ((ChatState)(java.lang.Object)(boxed548));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let274_0, boxed549 -> {
          ChatState _24_dt__update__tmp_h5 = ((ChatState)(java.lang.Object)(boxed549));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(__default.removeQ((_20_s3).dtor_queuedMessages(), (_11_qid).dtor_value()), boxed550 -> {
            dafny.DafnySequence<? extends QMsg> _pat_let275_0 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed550));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(_pat_let275_0, boxed551 -> {
              dafny.DafnySequence<? extends QMsg> _25_dt__update_hqueuedMessages_h0 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed551));
              return ChatState.create((_24_dt__update__tmp_h5).dtor_turns(), (_24_dt__update__tmp_h5).dtor_title(), (_24_dt__update__tmp_h5).dtor_status(), (_24_dt__update__tmp_h5).dtor_modifiedAt(), (_24_dt__update__tmp_h5).dtor_draft(), (_24_dt__update__tmp_h5).dtor_activity(), (_24_dt__update__tmp_h5).dtor_activeTurn(), (_24_dt__update__tmp_h5).dtor_steeringMessage(), _25_dt__update_hqueuedMessages_h0, (_24_dt__update__tmp_h5).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      )))) : (_20_s3));
      return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(_23_s4, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_CTurnComplete()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _26___mcc_h5 = ((agency.open.ahp.Chat.ChatAction_CTurnComplete)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _27_id = _26___mcc_h5;
      return __default.endTurn(s, _27_id, dafny.DafnySequence.asUnicodeString("complete"), false, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
    } else if (_source0.is_CTurnCancelled()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _28___mcc_h6 = ((agency.open.ahp.Chat.ChatAction_CTurnCancelled)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _29_id = _28___mcc_h6;
      return __default.endTurn(s, _29_id, dafny.DafnySequence.asUnicodeString("cancelled"), false, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
    } else if (_source0.is_CUsage()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _30___mcc_h7 = ((agency.open.ahp.Chat.ChatAction_CUsage)_source0)._turnId;
      agency.open.ahp.ConfluxCodec.Json _31___mcc_h8 = ((agency.open.ahp.Chat.ChatAction_CUsage)_source0)._usage;
      agency.open.ahp.ConfluxCodec.Json _32_u = _31___mcc_h8;
      dafny.DafnySequence<? extends dafny.CodePoint> _33_id = _30___mcc_h7;
      if (__default.turnMatches(s, _33_id)) {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed552 -> {
  ChatState _pat_let276_0 = ((ChatState)(java.lang.Object)(boxed552));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let276_0, boxed553 -> {
    ChatState _34_dt__update__tmp_h6 = ((ChatState)(java.lang.Object)(boxed553));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv0).dtor_activeTurn()).dtor_value(), boxed555 -> {
  Turn _pat_let278_0 = ((Turn)(java.lang.Object)(boxed555));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let278_0, boxed556 -> {
    Turn _35_dt__update__tmp_h7 = ((Turn)(java.lang.Object)(boxed556));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Turn>Let(__default.jNoNull(_32_u), boxed557 -> {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let279_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed557));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Turn>Let(_pat_let279_0, boxed558 -> {
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _36_dt__update_husage_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed558));
        return Turn.create((_35_dt__update__tmp_h7).dtor_turnId(), (_35_dt__update__tmp_h7).dtor_message(), (_35_dt__update__tmp_h7).dtor_parts(), (_35_dt__update__tmp_h7).dtor_state(), _36_dt__update_husage_h0, (_35_dt__update__tmp_h7).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed554 -> {
      agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let277_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed554));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let277_0, boxed559 -> {
        agency.open.ahp.AhpSkeleton.Option<Turn> _37_dt__update_hactiveTurn_h1 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed559));
        return ChatState.create((_34_dt__update__tmp_h6).dtor_turns(), (_34_dt__update__tmp_h6).dtor_title(), (_34_dt__update__tmp_h6).dtor_status(), (_34_dt__update__tmp_h6).dtor_modifiedAt(), (_34_dt__update__tmp_h6).dtor_draft(), (_34_dt__update__tmp_h6).dtor_activity(), _37_dt__update_hactiveTurn_h1, (_34_dt__update__tmp_h6).dtor_steeringMessage(), (_34_dt__update__tmp_h6).dtor_queuedMessages(), (_34_dt__update__tmp_h6).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CError()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _38___mcc_h9 = ((agency.open.ahp.Chat.ChatAction_CError)_source0)._turnId;
      agency.open.ahp.ConfluxCodec.Json _39___mcc_h10 = ((agency.open.ahp.Chat.ChatAction_CError)_source0)._err;
      agency.open.ahp.ConfluxCodec.Json _40_e = _39___mcc_h10;
      dafny.DafnySequence<? extends dafny.CodePoint> _41_id = _38___mcc_h9;
      return __default.endTurn(s, _41_id, dafny.DafnySequence.asUnicodeString("error"), true, __default.jNoNull(_40_e));
    } else if (_source0.is_CResponsePart()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _42___mcc_h11 = ((agency.open.ahp.Chat.ChatAction_CResponsePart)_source0)._turnId;
      Part _43___mcc_h12 = ((agency.open.ahp.Chat.ChatAction_CResponsePart)_source0)._part;
      Part _44_p = _43___mcc_h12;
      dafny.DafnySequence<? extends dafny.CodePoint> _45_id = _42___mcc_h11;
      if (__default.turnMatches(s, _45_id)) {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed560 -> {
  ChatState _pat_let280_0 = ((ChatState)(java.lang.Object)(boxed560));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let280_0, boxed561 -> {
    ChatState _46_dt__update__tmp_h8 = ((ChatState)(java.lang.Object)(boxed561));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), __default.appendPart(((_pat_let_tv1).dtor_activeTurn()).dtor_value(), _44_p)), boxed562 -> {
      agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let281_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed562));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let281_0, boxed563 -> {
        agency.open.ahp.AhpSkeleton.Option<Turn> _47_dt__update_hactiveTurn_h2 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed563));
        return ChatState.create((_46_dt__update__tmp_h8).dtor_turns(), (_46_dt__update__tmp_h8).dtor_title(), (_46_dt__update__tmp_h8).dtor_status(), (_46_dt__update__tmp_h8).dtor_modifiedAt(), (_46_dt__update__tmp_h8).dtor_draft(), (_46_dt__update__tmp_h8).dtor_activity(), _47_dt__update_hactiveTurn_h2, (_46_dt__update__tmp_h8).dtor_steeringMessage(), (_46_dt__update__tmp_h8).dtor_queuedMessages(), (_46_dt__update__tmp_h8).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CDelta()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _48___mcc_h13 = ((agency.open.ahp.Chat.ChatAction_CDelta)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _49___mcc_h14 = ((agency.open.ahp.Chat.ChatAction_CDelta)_source0)._partId;
      dafny.DafnySequence<? extends dafny.CodePoint> _50___mcc_h15 = ((agency.open.ahp.Chat.ChatAction_CDelta)_source0)._content;
      dafny.DafnySequence<? extends dafny.CodePoint> _51_c = _50___mcc_h15;
      dafny.DafnySequence<? extends dafny.CodePoint> _52_pid = _49___mcc_h14;
      dafny.DafnySequence<? extends dafny.CodePoint> _53_id = _48___mcc_h13;
      if (__default.turnMatches(s, _53_id)) {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed564 -> {
  ChatState _pat_let282_0 = ((ChatState)(java.lang.Object)(boxed564));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let282_0, boxed565 -> {
    ChatState _54_dt__update__tmp_h9 = ((ChatState)(java.lang.Object)(boxed565));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv2).dtor_activeTurn()).dtor_value(), boxed567 -> {
  Turn _pat_let284_0 = ((Turn)(java.lang.Object)(boxed567));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let284_0, boxed568 -> {
    Turn _55_dt__update__tmp_h10 = ((Turn)(java.lang.Object)(boxed568));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.deltaMarkdown((((_pat_let_tv3).dtor_activeTurn()).dtor_value()).dtor_parts(), _52_pid, _51_c), boxed569 -> {
      dafny.DafnySequence<? extends Part> _pat_let285_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed569));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let285_0, boxed570 -> {
        dafny.DafnySequence<? extends Part> _56_dt__update_hparts_h0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed570));
        return Turn.create((_55_dt__update__tmp_h10).dtor_turnId(), (_55_dt__update__tmp_h10).dtor_message(), _56_dt__update_hparts_h0, (_55_dt__update__tmp_h10).dtor_state(), (_55_dt__update__tmp_h10).dtor_usage(), (_55_dt__update__tmp_h10).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed566 -> {
      agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let283_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed566));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let283_0, boxed571 -> {
        agency.open.ahp.AhpSkeleton.Option<Turn> _57_dt__update_hactiveTurn_h3 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed571));
        return ChatState.create((_54_dt__update__tmp_h9).dtor_turns(), (_54_dt__update__tmp_h9).dtor_title(), (_54_dt__update__tmp_h9).dtor_status(), (_54_dt__update__tmp_h9).dtor_modifiedAt(), (_54_dt__update__tmp_h9).dtor_draft(), (_54_dt__update__tmp_h9).dtor_activity(), _57_dt__update_hactiveTurn_h3, (_54_dt__update__tmp_h9).dtor_steeringMessage(), (_54_dt__update__tmp_h9).dtor_queuedMessages(), (_54_dt__update__tmp_h9).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CToolCallStart()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _58___mcc_h16 = ((agency.open.ahp.Chat.ChatAction_CToolCallStart)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _59___mcc_h17 = ((agency.open.ahp.Chat.ChatAction_CToolCallStart)_source0)._toolCallId;
      dafny.DafnySequence<? extends dafny.CodePoint> _60___mcc_h18 = ((agency.open.ahp.Chat.ChatAction_CToolCallStart)_source0)._toolName;
      dafny.DafnySequence<? extends dafny.CodePoint> _61___mcc_h19 = ((agency.open.ahp.Chat.ChatAction_CToolCallStart)_source0)._displayName;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _62___mcc_h20 = ((agency.open.ahp.Chat.ChatAction_CToolCallStart)_source0)._intention;
      agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> _63___mcc_h21 = ((agency.open.ahp.Chat.ChatAction_CToolCallStart)_source0)._contributor;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _64___mcc_h22 = ((agency.open.ahp.Chat.ChatAction_CToolCallStart)_source0)._meta;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _65_meta = _64___mcc_h22;
      agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> _66_contributor = _63___mcc_h21;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _67_intent = _62___mcc_h20;
      dafny.DafnySequence<? extends dafny.CodePoint> _68_dn = _61___mcc_h19;
      dafny.DafnySequence<? extends dafny.CodePoint> _69_tn = _60___mcc_h18;
      dafny.DafnySequence<? extends dafny.CodePoint> _70_tcId = _59___mcc_h17;
      dafny.DafnySequence<? extends dafny.CodePoint> _71_id = _58___mcc_h16;
      if (__default.turnMatches(s, _71_id)) {
        ToolCall _72_tc = ToolCall.create(_70_tcId, _69_tn, _68_dn, dafny.DafnySequence.asUnicodeString("streaming"), _67_intent, _66_contributor, _65_meta, dafny.DafnySequence.asUnicodeString(""), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)));
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed572 -> {
  ChatState _pat_let286_0 = ((ChatState)(java.lang.Object)(boxed572));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let286_0, boxed573 -> {
    ChatState _73_dt__update__tmp_h13 = ((ChatState)(java.lang.Object)(boxed573));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv4).dtor_activeTurn()).dtor_value(), boxed575 -> {
  Turn _pat_let288_0 = ((Turn)(java.lang.Object)(boxed575));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let288_0, boxed576 -> {
    Turn _74_dt__update__tmp_h14 = ((Turn)(java.lang.Object)(boxed576));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part>concatenate((((_pat_let_tv5).dtor_activeTurn()).dtor_value()).dtor_parts(), dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PToolCall(_72_tc))), boxed577 -> {
      dafny.DafnySequence<? extends Part> _pat_let289_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed577));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let289_0, boxed578 -> {
        dafny.DafnySequence<? extends Part> _75_dt__update_hparts_h2 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed578));
        return Turn.create((_74_dt__update__tmp_h14).dtor_turnId(), (_74_dt__update__tmp_h14).dtor_message(), _75_dt__update_hparts_h2, (_74_dt__update__tmp_h14).dtor_state(), (_74_dt__update__tmp_h14).dtor_usage(), (_74_dt__update__tmp_h14).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed574 -> {
      agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let287_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed574));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let287_0, boxed579 -> {
        agency.open.ahp.AhpSkeleton.Option<Turn> _76_dt__update_hactiveTurn_h5 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed579));
        return ChatState.create((_73_dt__update__tmp_h13).dtor_turns(), (_73_dt__update__tmp_h13).dtor_title(), (_73_dt__update__tmp_h13).dtor_status(), (_73_dt__update__tmp_h13).dtor_modifiedAt(), (_73_dt__update__tmp_h13).dtor_draft(), (_73_dt__update__tmp_h13).dtor_activity(), _76_dt__update_hactiveTurn_h5, (_73_dt__update__tmp_h13).dtor_steeringMessage(), (_73_dt__update__tmp_h13).dtor_queuedMessages(), (_73_dt__update__tmp_h13).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CToolCallDelta()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _77___mcc_h23 = ((agency.open.ahp.Chat.ChatAction_CToolCallDelta)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _78___mcc_h24 = ((agency.open.ahp.Chat.ChatAction_CToolCallDelta)_source0)._toolCallId;
      dafny.DafnySequence<? extends dafny.CodePoint> _79___mcc_h25 = ((agency.open.ahp.Chat.ChatAction_CToolCallDelta)_source0)._content;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _80___mcc_h26 = ((agency.open.ahp.Chat.ChatAction_CToolCallDelta)_source0)._invocationMessage;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _81___mcc_h27 = ((agency.open.ahp.Chat.ChatAction_CToolCallDelta)_source0)._meta;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _82_meta = _81___mcc_h27;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _83_inv = _80___mcc_h26;
      dafny.DafnySequence<? extends dafny.CodePoint> _84_c = _79___mcc_h25;
      dafny.DafnySequence<? extends dafny.CodePoint> _85_tcId = _78___mcc_h24;
      dafny.DafnySequence<? extends dafny.CodePoint> _86_id = _77___mcc_h23;
      if (__default.turnMatches(s, _86_id)) {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed580 -> {
  ChatState _pat_let290_0 = ((ChatState)(java.lang.Object)(boxed580));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let290_0, boxed581 -> {
    ChatState _87_dt__update__tmp_h15 = ((ChatState)(java.lang.Object)(boxed581));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv6).dtor_activeTurn()).dtor_value(), boxed583 -> {
  Turn _pat_let292_0 = ((Turn)(java.lang.Object)(boxed583));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let292_0, boxed584 -> {
    Turn _88_dt__update__tmp_h16 = ((Turn)(java.lang.Object)(boxed584));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.mapTC((((_pat_let_tv7).dtor_activeTurn()).dtor_value()).dtor_parts(), _85_tcId, ((dafny.Function3<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, java.util.function.Function<ToolCall, ToolCall>>)(_89_c, _90_inv, _91_meta) -> {return ((java.util.function.Function<ToolCall, ToolCall>)(_92_x_boxed0) -> {
      ToolCall _92_x = ((ToolCall)(java.lang.Object)(_92_x_boxed0));
      return __default.deltaOne(_92_x, _89_c, _90_inv, _91_meta);
    });}).apply(_84_c, _83_inv, _82_meta)), boxed585 -> {
      dafny.DafnySequence<? extends Part> _pat_let293_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed585));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let293_0, boxed586 -> {
        dafny.DafnySequence<? extends Part> _93_dt__update_hparts_h3 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed586));
        return Turn.create((_88_dt__update__tmp_h16).dtor_turnId(), (_88_dt__update__tmp_h16).dtor_message(), _93_dt__update_hparts_h3, (_88_dt__update__tmp_h16).dtor_state(), (_88_dt__update__tmp_h16).dtor_usage(), (_88_dt__update__tmp_h16).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed582 -> {
      agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let291_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed582));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let291_0, boxed587 -> {
        agency.open.ahp.AhpSkeleton.Option<Turn> _94_dt__update_hactiveTurn_h6 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed587));
        return ChatState.create((_87_dt__update__tmp_h15).dtor_turns(), (_87_dt__update__tmp_h15).dtor_title(), (_87_dt__update__tmp_h15).dtor_status(), (_87_dt__update__tmp_h15).dtor_modifiedAt(), (_87_dt__update__tmp_h15).dtor_draft(), (_87_dt__update__tmp_h15).dtor_activity(), _94_dt__update_hactiveTurn_h6, (_87_dt__update__tmp_h15).dtor_steeringMessage(), (_87_dt__update__tmp_h15).dtor_queuedMessages(), (_87_dt__update__tmp_h15).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CToolCallReady()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _95___mcc_h28 = ((agency.open.ahp.Chat.ChatAction_CToolCallReady)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _96___mcc_h29 = ((agency.open.ahp.Chat.ChatAction_CToolCallReady)_source0)._toolCallId;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _97___mcc_h30 = ((agency.open.ahp.Chat.ChatAction_CToolCallReady)_source0)._invocationMessage;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _98___mcc_h31 = ((agency.open.ahp.Chat.ChatAction_CToolCallReady)_source0)._toolInput;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _99___mcc_h32 = ((agency.open.ahp.Chat.ChatAction_CToolCallReady)_source0)._confirmationTitle;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _100___mcc_h33 = ((agency.open.ahp.Chat.ChatAction_CToolCallReady)_source0)._riskAssessment;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _101___mcc_h34 = ((agency.open.ahp.Chat.ChatAction_CToolCallReady)_source0)._edits;
      agency.open.ahp.AhpSkeleton.Option<Boolean> _102___mcc_h35 = ((agency.open.ahp.Chat.ChatAction_CToolCallReady)_source0)._editable;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _103___mcc_h36 = ((agency.open.ahp.Chat.ChatAction_CToolCallReady)_source0)._options;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _104___mcc_h37 = ((agency.open.ahp.Chat.ChatAction_CToolCallReady)_source0)._confirmed;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _105___mcc_h38 = ((agency.open.ahp.Chat.ChatAction_CToolCallReady)_source0)._meta;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _106_meta = _105___mcc_h38;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _107_confirmed = _104___mcc_h37;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _108_options = _103___mcc_h36;
      agency.open.ahp.AhpSkeleton.Option<Boolean> _109_editable = _102___mcc_h35;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _110_edits = _101___mcc_h34;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _111_risk = _100___mcc_h33;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _112_title = _99___mcc_h32;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _113_input = _98___mcc_h31;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _114_inv = _97___mcc_h30;
      dafny.DafnySequence<? extends dafny.CodePoint> _115_tcId = _96___mcc_h29;
      dafny.DafnySequence<? extends dafny.CodePoint> _116_id = _95___mcc_h28;
      if (__default.turnMatches(s, _116_id)) {
        ChatState _117_s1 = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed588 -> {
          ChatState _pat_let294_0 = ((ChatState)(java.lang.Object)(boxed588));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let294_0, boxed589 -> {
            ChatState _118_dt__update__tmp_h17 = ((ChatState)(java.lang.Object)(boxed589));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv8).dtor_activeTurn()).dtor_value(), boxed591 -> {
  Turn _pat_let296_0 = ((Turn)(java.lang.Object)(boxed591));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let296_0, boxed592 -> {
    Turn _119_dt__update__tmp_h18 = ((Turn)(java.lang.Object)(boxed592));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.mapTC((((_pat_let_tv9).dtor_activeTurn()).dtor_value()).dtor_parts(), _115_tcId, ((dafny.Function9<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, agency.open.ahp.AhpSkeleton.Option<Boolean>, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, java.util.function.Function<ToolCall, ToolCall>>)(_120_inv, _121_input, _122_title, _123_risk, _124_edits, _125_editable, _126_options, _127_confirmed, _128_meta) -> {return ((java.util.function.Function<ToolCall, ToolCall>)(_129_x_boxed0) -> {
      ToolCall _129_x = ((ToolCall)(java.lang.Object)(_129_x_boxed0));
      return __default.readyOne(_129_x, _120_inv, _121_input, _122_title, _123_risk, _124_edits, _125_editable, _126_options, _127_confirmed, _128_meta);
    });}).apply(_114_inv, _113_input, _112_title, _111_risk, _110_edits, _109_editable, _108_options, _107_confirmed, _106_meta)), boxed593 -> {
      dafny.DafnySequence<? extends Part> _pat_let297_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed593));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let297_0, boxed594 -> {
        dafny.DafnySequence<? extends Part> _130_dt__update_hparts_h4 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed594));
        return Turn.create((_119_dt__update__tmp_h18).dtor_turnId(), (_119_dt__update__tmp_h18).dtor_message(), _130_dt__update_hparts_h4, (_119_dt__update__tmp_h18).dtor_state(), (_119_dt__update__tmp_h18).dtor_usage(), (_119_dt__update__tmp_h18).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed590 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let295_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed590));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let295_0, boxed595 -> {
                agency.open.ahp.AhpSkeleton.Option<Turn> _131_dt__update_hactiveTurn_h7 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed595));
                return ChatState.create((_118_dt__update__tmp_h17).dtor_turns(), (_118_dt__update__tmp_h17).dtor_title(), (_118_dt__update__tmp_h17).dtor_status(), (_118_dt__update__tmp_h17).dtor_modifiedAt(), (_118_dt__update__tmp_h17).dtor_draft(), (_118_dt__update__tmp_h17).dtor_activity(), _131_dt__update_hactiveTurn_h7, (_118_dt__update__tmp_h17).dtor_steeringMessage(), (_118_dt__update__tmp_h17).dtor_queuedMessages(), (_118_dt__update__tmp_h17).dtor_nextCursor());
              }
              )));
            }
            )));
          }
          )));
        }
        )));
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_117_s1, boxed596 -> {
  ChatState _pat_let298_0 = ((ChatState)(java.lang.Object)(boxed596));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let298_0, boxed597 -> {
    ChatState _132_dt__update__tmp_h19 = ((ChatState)(java.lang.Object)(boxed597));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(__default.summaryStatus(_117_s1, false), boxed598 -> {
      int _pat_let299_0 = ((int)(java.lang.Object)(boxed598));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let299_0, boxed599 -> {
        int _133_dt__update_hstatus_h1 = ((int)(java.lang.Object)(boxed599));
        return ChatState.create((_132_dt__update__tmp_h19).dtor_turns(), (_132_dt__update__tmp_h19).dtor_title(), _133_dt__update_hstatus_h1, (_132_dt__update__tmp_h19).dtor_modifiedAt(), (_132_dt__update__tmp_h19).dtor_draft(), (_132_dt__update__tmp_h19).dtor_activity(), (_132_dt__update__tmp_h19).dtor_activeTurn(), (_132_dt__update__tmp_h19).dtor_steeringMessage(), (_132_dt__update__tmp_h19).dtor_queuedMessages(), (_132_dt__update__tmp_h19).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CToolCallConfirmed()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _134___mcc_h39 = ((agency.open.ahp.Chat.ChatAction_CToolCallConfirmed)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _135___mcc_h40 = ((agency.open.ahp.Chat.ChatAction_CToolCallConfirmed)_source0)._toolCallId;
      boolean _136___mcc_h41 = ((agency.open.ahp.Chat.ChatAction_CToolCallConfirmed)_source0)._approved;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _137___mcc_h42 = ((agency.open.ahp.Chat.ChatAction_CToolCallConfirmed)_source0)._confirmedVal;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _138___mcc_h43 = ((agency.open.ahp.Chat.ChatAction_CToolCallConfirmed)_source0)._reason;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _139___mcc_h44 = ((agency.open.ahp.Chat.ChatAction_CToolCallConfirmed)_source0)._reasonMessage;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _140___mcc_h45 = ((agency.open.ahp.Chat.ChatAction_CToolCallConfirmed)_source0)._userSuggestion;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _141___mcc_h46 = ((agency.open.ahp.Chat.ChatAction_CToolCallConfirmed)_source0)._editedToolInput;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _142___mcc_h47 = ((agency.open.ahp.Chat.ChatAction_CToolCallConfirmed)_source0)._selectedOptionId;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _143___mcc_h48 = ((agency.open.ahp.Chat.ChatAction_CToolCallConfirmed)_source0)._meta;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _144_meta = _143___mcc_h48;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _145_selectedOptionId = _142___mcc_h47;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _146_edited = _141___mcc_h46;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _147_userSuggestion = _140___mcc_h45;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _148_reasonMessage = _139___mcc_h44;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _149_rsn = _138___mcc_h43;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _150_cv = _137___mcc_h42;
      boolean _151_approved = _136___mcc_h41;
      dafny.DafnySequence<? extends dafny.CodePoint> _152_tcId = _135___mcc_h40;
      dafny.DafnySequence<? extends dafny.CodePoint> _153_id = _134___mcc_h39;
      if (__default.turnMatches(s, _153_id)) {
        ChatState _154_s1 = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed600 -> {
          ChatState _pat_let300_0 = ((ChatState)(java.lang.Object)(boxed600));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let300_0, boxed601 -> {
            ChatState _155_dt__update__tmp_h20 = ((ChatState)(java.lang.Object)(boxed601));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv10).dtor_activeTurn()).dtor_value(), boxed603 -> {
  Turn _pat_let302_0 = ((Turn)(java.lang.Object)(boxed603));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let302_0, boxed604 -> {
    Turn _156_dt__update__tmp_h21 = ((Turn)(java.lang.Object)(boxed604));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.mapTC((((_pat_let_tv11).dtor_activeTurn()).dtor_value()).dtor_parts(), _152_tcId, ((dafny.Function8<Boolean, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, java.util.function.Function<ToolCall, ToolCall>>)(__157_approved0, _158_cv, _159_rsn, _160_reasonMessage, _161_userSuggestion, _162_edited, _163_selectedOptionId, _164_meta) -> {boolean _157_approved = ((boolean)(java.lang.Object)(__157_approved0));
    return ((java.util.function.Function<ToolCall, ToolCall>)(_165_x_boxed0) -> {
      ToolCall _165_x = ((ToolCall)(java.lang.Object)(_165_x_boxed0));
      return __default.confirmOne(_165_x, _157_approved, _158_cv, _159_rsn, _160_reasonMessage, _161_userSuggestion, _162_edited, _163_selectedOptionId, _164_meta);
    });}).apply(_151_approved, _150_cv, _149_rsn, _148_reasonMessage, _147_userSuggestion, _146_edited, _145_selectedOptionId, _144_meta)), boxed605 -> {
      dafny.DafnySequence<? extends Part> _pat_let303_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed605));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let303_0, boxed606 -> {
        dafny.DafnySequence<? extends Part> _166_dt__update_hparts_h5 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed606));
        return Turn.create((_156_dt__update__tmp_h21).dtor_turnId(), (_156_dt__update__tmp_h21).dtor_message(), _166_dt__update_hparts_h5, (_156_dt__update__tmp_h21).dtor_state(), (_156_dt__update__tmp_h21).dtor_usage(), (_156_dt__update__tmp_h21).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed602 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let301_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed602));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let301_0, boxed607 -> {
                agency.open.ahp.AhpSkeleton.Option<Turn> _167_dt__update_hactiveTurn_h8 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed607));
                return ChatState.create((_155_dt__update__tmp_h20).dtor_turns(), (_155_dt__update__tmp_h20).dtor_title(), (_155_dt__update__tmp_h20).dtor_status(), (_155_dt__update__tmp_h20).dtor_modifiedAt(), (_155_dt__update__tmp_h20).dtor_draft(), (_155_dt__update__tmp_h20).dtor_activity(), _167_dt__update_hactiveTurn_h8, (_155_dt__update__tmp_h20).dtor_steeringMessage(), (_155_dt__update__tmp_h20).dtor_queuedMessages(), (_155_dt__update__tmp_h20).dtor_nextCursor());
              }
              )));
            }
            )));
          }
          )));
        }
        )));
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_154_s1, boxed608 -> {
  ChatState _pat_let304_0 = ((ChatState)(java.lang.Object)(boxed608));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let304_0, boxed609 -> {
    ChatState _168_dt__update__tmp_h22 = ((ChatState)(java.lang.Object)(boxed609));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(__default.summaryStatus(_154_s1, false), boxed610 -> {
      int _pat_let305_0 = ((int)(java.lang.Object)(boxed610));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let305_0, boxed611 -> {
        int _169_dt__update_hstatus_h2 = ((int)(java.lang.Object)(boxed611));
        return ChatState.create((_168_dt__update__tmp_h22).dtor_turns(), (_168_dt__update__tmp_h22).dtor_title(), _169_dt__update_hstatus_h2, (_168_dt__update__tmp_h22).dtor_modifiedAt(), (_168_dt__update__tmp_h22).dtor_draft(), (_168_dt__update__tmp_h22).dtor_activity(), (_168_dt__update__tmp_h22).dtor_activeTurn(), (_168_dt__update__tmp_h22).dtor_steeringMessage(), (_168_dt__update__tmp_h22).dtor_queuedMessages(), (_168_dt__update__tmp_h22).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CToolCallAuthRequired()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _170___mcc_h49 = ((agency.open.ahp.Chat.ChatAction_CToolCallAuthRequired)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _171___mcc_h50 = ((agency.open.ahp.Chat.ChatAction_CToolCallAuthRequired)_source0)._toolCallId;
      agency.open.ahp.ConfluxCodec.Json _172___mcc_h51 = ((agency.open.ahp.Chat.ChatAction_CToolCallAuthRequired)_source0)._auth;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _173___mcc_h52 = ((agency.open.ahp.Chat.ChatAction_CToolCallAuthRequired)_source0)._meta;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _174_meta = _173___mcc_h52;
      agency.open.ahp.ConfluxCodec.Json _175_auth = _172___mcc_h51;
      dafny.DafnySequence<? extends dafny.CodePoint> _176_tcId = _171___mcc_h50;
      dafny.DafnySequence<? extends dafny.CodePoint> _177_id = _170___mcc_h49;
      if (__default.turnMatches(s, _177_id)) {
        ChatState _178_s1 = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed612 -> {
          ChatState _pat_let306_0 = ((ChatState)(java.lang.Object)(boxed612));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let306_0, boxed613 -> {
            ChatState _179_dt__update__tmp_h23 = ((ChatState)(java.lang.Object)(boxed613));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv12).dtor_activeTurn()).dtor_value(), boxed615 -> {
  Turn _pat_let308_0 = ((Turn)(java.lang.Object)(boxed615));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let308_0, boxed616 -> {
    Turn _180_dt__update__tmp_h24 = ((Turn)(java.lang.Object)(boxed616));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.mapTC((((_pat_let_tv13).dtor_activeTurn()).dtor_value()).dtor_parts(), _176_tcId, ((dafny.Function2<agency.open.ahp.ConfluxCodec.Json, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, java.util.function.Function<ToolCall, ToolCall>>)(_181_auth, _182_meta) -> {return ((java.util.function.Function<ToolCall, ToolCall>)(_183_x_boxed0) -> {
      ToolCall _183_x = ((ToolCall)(java.lang.Object)(_183_x_boxed0));
      return __default.authRequiredOne(_183_x, _181_auth, _182_meta);
    });}).apply(_175_auth, _174_meta)), boxed617 -> {
      dafny.DafnySequence<? extends Part> _pat_let309_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed617));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let309_0, boxed618 -> {
        dafny.DafnySequence<? extends Part> _184_dt__update_hparts_h6 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed618));
        return Turn.create((_180_dt__update__tmp_h24).dtor_turnId(), (_180_dt__update__tmp_h24).dtor_message(), _184_dt__update_hparts_h6, (_180_dt__update__tmp_h24).dtor_state(), (_180_dt__update__tmp_h24).dtor_usage(), (_180_dt__update__tmp_h24).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed614 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let307_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed614));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let307_0, boxed619 -> {
                agency.open.ahp.AhpSkeleton.Option<Turn> _185_dt__update_hactiveTurn_h9 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed619));
                return ChatState.create((_179_dt__update__tmp_h23).dtor_turns(), (_179_dt__update__tmp_h23).dtor_title(), (_179_dt__update__tmp_h23).dtor_status(), (_179_dt__update__tmp_h23).dtor_modifiedAt(), (_179_dt__update__tmp_h23).dtor_draft(), (_179_dt__update__tmp_h23).dtor_activity(), _185_dt__update_hactiveTurn_h9, (_179_dt__update__tmp_h23).dtor_steeringMessage(), (_179_dt__update__tmp_h23).dtor_queuedMessages(), (_179_dt__update__tmp_h23).dtor_nextCursor());
              }
              )));
            }
            )));
          }
          )));
        }
        )));
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_178_s1, boxed620 -> {
  ChatState _pat_let310_0 = ((ChatState)(java.lang.Object)(boxed620));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let310_0, boxed621 -> {
    ChatState _186_dt__update__tmp_h25 = ((ChatState)(java.lang.Object)(boxed621));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(__default.summaryStatus(_178_s1, false), boxed622 -> {
      int _pat_let311_0 = ((int)(java.lang.Object)(boxed622));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let311_0, boxed623 -> {
        int _187_dt__update_hstatus_h3 = ((int)(java.lang.Object)(boxed623));
        return ChatState.create((_186_dt__update__tmp_h25).dtor_turns(), (_186_dt__update__tmp_h25).dtor_title(), _187_dt__update_hstatus_h3, (_186_dt__update__tmp_h25).dtor_modifiedAt(), (_186_dt__update__tmp_h25).dtor_draft(), (_186_dt__update__tmp_h25).dtor_activity(), (_186_dt__update__tmp_h25).dtor_activeTurn(), (_186_dt__update__tmp_h25).dtor_steeringMessage(), (_186_dt__update__tmp_h25).dtor_queuedMessages(), (_186_dt__update__tmp_h25).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CToolCallAuthResolved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _188___mcc_h53 = ((agency.open.ahp.Chat.ChatAction_CToolCallAuthResolved)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _189___mcc_h54 = ((agency.open.ahp.Chat.ChatAction_CToolCallAuthResolved)_source0)._toolCallId;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _190___mcc_h55 = ((agency.open.ahp.Chat.ChatAction_CToolCallAuthResolved)_source0)._meta;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _191_meta = _190___mcc_h55;
      dafny.DafnySequence<? extends dafny.CodePoint> _192_tcId = _189___mcc_h54;
      dafny.DafnySequence<? extends dafny.CodePoint> _193_id = _188___mcc_h53;
      if (__default.turnMatches(s, _193_id)) {
        ChatState _194_s1 = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed624 -> {
          ChatState _pat_let312_0 = ((ChatState)(java.lang.Object)(boxed624));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let312_0, boxed625 -> {
            ChatState _195_dt__update__tmp_h26 = ((ChatState)(java.lang.Object)(boxed625));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv14).dtor_activeTurn()).dtor_value(), boxed627 -> {
  Turn _pat_let314_0 = ((Turn)(java.lang.Object)(boxed627));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let314_0, boxed628 -> {
    Turn _196_dt__update__tmp_h27 = ((Turn)(java.lang.Object)(boxed628));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.mapTC((((_pat_let_tv15).dtor_activeTurn()).dtor_value()).dtor_parts(), _192_tcId, ((java.util.function.Function<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, java.util.function.Function<ToolCall, ToolCall>>)(_197_meta) -> {return ((java.util.function.Function<ToolCall, ToolCall>)(_198_x_boxed0) -> {
      ToolCall _198_x = ((ToolCall)(java.lang.Object)(_198_x_boxed0));
      return __default.authResolvedOne(_198_x, _197_meta);
    });}).apply(_191_meta)), boxed629 -> {
      dafny.DafnySequence<? extends Part> _pat_let315_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed629));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let315_0, boxed630 -> {
        dafny.DafnySequence<? extends Part> _199_dt__update_hparts_h7 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed630));
        return Turn.create((_196_dt__update__tmp_h27).dtor_turnId(), (_196_dt__update__tmp_h27).dtor_message(), _199_dt__update_hparts_h7, (_196_dt__update__tmp_h27).dtor_state(), (_196_dt__update__tmp_h27).dtor_usage(), (_196_dt__update__tmp_h27).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed626 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let313_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed626));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let313_0, boxed631 -> {
                agency.open.ahp.AhpSkeleton.Option<Turn> _200_dt__update_hactiveTurn_h10 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed631));
                return ChatState.create((_195_dt__update__tmp_h26).dtor_turns(), (_195_dt__update__tmp_h26).dtor_title(), (_195_dt__update__tmp_h26).dtor_status(), (_195_dt__update__tmp_h26).dtor_modifiedAt(), (_195_dt__update__tmp_h26).dtor_draft(), (_195_dt__update__tmp_h26).dtor_activity(), _200_dt__update_hactiveTurn_h10, (_195_dt__update__tmp_h26).dtor_steeringMessage(), (_195_dt__update__tmp_h26).dtor_queuedMessages(), (_195_dt__update__tmp_h26).dtor_nextCursor());
              }
              )));
            }
            )));
          }
          )));
        }
        )));
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_194_s1, boxed632 -> {
  ChatState _pat_let316_0 = ((ChatState)(java.lang.Object)(boxed632));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let316_0, boxed633 -> {
    ChatState _201_dt__update__tmp_h28 = ((ChatState)(java.lang.Object)(boxed633));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(__default.summaryStatus(_194_s1, false), boxed634 -> {
      int _pat_let317_0 = ((int)(java.lang.Object)(boxed634));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let317_0, boxed635 -> {
        int _202_dt__update_hstatus_h4 = ((int)(java.lang.Object)(boxed635));
        return ChatState.create((_201_dt__update__tmp_h28).dtor_turns(), (_201_dt__update__tmp_h28).dtor_title(), _202_dt__update_hstatus_h4, (_201_dt__update__tmp_h28).dtor_modifiedAt(), (_201_dt__update__tmp_h28).dtor_draft(), (_201_dt__update__tmp_h28).dtor_activity(), (_201_dt__update__tmp_h28).dtor_activeTurn(), (_201_dt__update__tmp_h28).dtor_steeringMessage(), (_201_dt__update__tmp_h28).dtor_queuedMessages(), (_201_dt__update__tmp_h28).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CToolCallComplete()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _203___mcc_h56 = ((agency.open.ahp.Chat.ChatAction_CToolCallComplete)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _204___mcc_h57 = ((agency.open.ahp.Chat.ChatAction_CToolCallComplete)_source0)._toolCallId;
      boolean _205___mcc_h58 = ((agency.open.ahp.Chat.ChatAction_CToolCallComplete)_source0)._success;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _206___mcc_h59 = ((agency.open.ahp.Chat.ChatAction_CToolCallComplete)_source0)._pastTenseMessage;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _207___mcc_h60 = ((agency.open.ahp.Chat.ChatAction_CToolCallComplete)_source0)._resultContent;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _208___mcc_h61 = ((agency.open.ahp.Chat.ChatAction_CToolCallComplete)_source0)._structuredContent;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _209___mcc_h62 = ((agency.open.ahp.Chat.ChatAction_CToolCallComplete)_source0)._error;
      boolean _210___mcc_h63 = ((agency.open.ahp.Chat.ChatAction_CToolCallComplete)_source0)._requiresResultConfirmation;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _211___mcc_h64 = ((agency.open.ahp.Chat.ChatAction_CToolCallComplete)_source0)._meta;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _212_meta = _211___mcc_h64;
      boolean _213_rrc = _210___mcc_h63;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _214_err = _209___mcc_h62;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _215_structured = _208___mcc_h61;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _216_resultContent = _207___mcc_h60;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _217_past = _206___mcc_h59;
      boolean _218_ok = _205___mcc_h58;
      dafny.DafnySequence<? extends dafny.CodePoint> _219_tcId = _204___mcc_h57;
      dafny.DafnySequence<? extends dafny.CodePoint> _220_id = _203___mcc_h56;
      if (__default.turnMatches(s, _220_id)) {
        ChatState _221_s1 = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed636 -> {
          ChatState _pat_let318_0 = ((ChatState)(java.lang.Object)(boxed636));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let318_0, boxed637 -> {
            ChatState _222_dt__update__tmp_h29 = ((ChatState)(java.lang.Object)(boxed637));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv16).dtor_activeTurn()).dtor_value(), boxed639 -> {
  Turn _pat_let320_0 = ((Turn)(java.lang.Object)(boxed639));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let320_0, boxed640 -> {
    Turn _223_dt__update__tmp_h30 = ((Turn)(java.lang.Object)(boxed640));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.mapTC((((_pat_let_tv17).dtor_activeTurn()).dtor_value()).dtor_parts(), _219_tcId, ((dafny.Function7<Boolean, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Boolean, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, java.util.function.Function<ToolCall, ToolCall>>)(__224_ok0, _225_past, _226_resultContent, _227_structured, _228_err, __229_rrc0, _230_meta) -> {boolean _224_ok = ((boolean)(java.lang.Object)(__224_ok0));
    boolean _229_rrc = ((boolean)(java.lang.Object)(__229_rrc0));
    return ((java.util.function.Function<ToolCall, ToolCall>)(_231_x_boxed0) -> {
      ToolCall _231_x = ((ToolCall)(java.lang.Object)(_231_x_boxed0));
      return __default.completeOne(_231_x, _224_ok, _225_past, _226_resultContent, _227_structured, _228_err, _229_rrc, _230_meta);
    });}).apply(_218_ok, _217_past, _216_resultContent, _215_structured, _214_err, _213_rrc, _212_meta)), boxed641 -> {
      dafny.DafnySequence<? extends Part> _pat_let321_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed641));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let321_0, boxed642 -> {
        dafny.DafnySequence<? extends Part> _232_dt__update_hparts_h8 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed642));
        return Turn.create((_223_dt__update__tmp_h30).dtor_turnId(), (_223_dt__update__tmp_h30).dtor_message(), _232_dt__update_hparts_h8, (_223_dt__update__tmp_h30).dtor_state(), (_223_dt__update__tmp_h30).dtor_usage(), (_223_dt__update__tmp_h30).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed638 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let319_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed638));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let319_0, boxed643 -> {
                agency.open.ahp.AhpSkeleton.Option<Turn> _233_dt__update_hactiveTurn_h11 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed643));
                return ChatState.create((_222_dt__update__tmp_h29).dtor_turns(), (_222_dt__update__tmp_h29).dtor_title(), (_222_dt__update__tmp_h29).dtor_status(), (_222_dt__update__tmp_h29).dtor_modifiedAt(), (_222_dt__update__tmp_h29).dtor_draft(), (_222_dt__update__tmp_h29).dtor_activity(), _233_dt__update_hactiveTurn_h11, (_222_dt__update__tmp_h29).dtor_steeringMessage(), (_222_dt__update__tmp_h29).dtor_queuedMessages(), (_222_dt__update__tmp_h29).dtor_nextCursor());
              }
              )));
            }
            )));
          }
          )));
        }
        )));
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_221_s1, boxed644 -> {
  ChatState _pat_let322_0 = ((ChatState)(java.lang.Object)(boxed644));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let322_0, boxed645 -> {
    ChatState _234_dt__update__tmp_h31 = ((ChatState)(java.lang.Object)(boxed645));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(__default.summaryStatus(_221_s1, false), boxed646 -> {
      int _pat_let323_0 = ((int)(java.lang.Object)(boxed646));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let323_0, boxed647 -> {
        int _235_dt__update_hstatus_h5 = ((int)(java.lang.Object)(boxed647));
        return ChatState.create((_234_dt__update__tmp_h31).dtor_turns(), (_234_dt__update__tmp_h31).dtor_title(), _235_dt__update_hstatus_h5, (_234_dt__update__tmp_h31).dtor_modifiedAt(), (_234_dt__update__tmp_h31).dtor_draft(), (_234_dt__update__tmp_h31).dtor_activity(), (_234_dt__update__tmp_h31).dtor_activeTurn(), (_234_dt__update__tmp_h31).dtor_steeringMessage(), (_234_dt__update__tmp_h31).dtor_queuedMessages(), (_234_dt__update__tmp_h31).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CToolCallResultConfirmed()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _236___mcc_h65 = ((agency.open.ahp.Chat.ChatAction_CToolCallResultConfirmed)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _237___mcc_h66 = ((agency.open.ahp.Chat.ChatAction_CToolCallResultConfirmed)_source0)._toolCallId;
      boolean _238___mcc_h67 = ((agency.open.ahp.Chat.ChatAction_CToolCallResultConfirmed)_source0)._approved;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _239___mcc_h68 = ((agency.open.ahp.Chat.ChatAction_CToolCallResultConfirmed)_source0)._meta;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _240_meta = _239___mcc_h68;
      boolean _241_approved = _238___mcc_h67;
      dafny.DafnySequence<? extends dafny.CodePoint> _242_tcId = _237___mcc_h66;
      dafny.DafnySequence<? extends dafny.CodePoint> _243_id = _236___mcc_h65;
      if (__default.turnMatches(s, _243_id)) {
        ChatState _244_s1 = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed648 -> {
          ChatState _pat_let324_0 = ((ChatState)(java.lang.Object)(boxed648));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let324_0, boxed649 -> {
            ChatState _245_dt__update__tmp_h32 = ((ChatState)(java.lang.Object)(boxed649));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv18).dtor_activeTurn()).dtor_value(), boxed651 -> {
  Turn _pat_let326_0 = ((Turn)(java.lang.Object)(boxed651));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let326_0, boxed652 -> {
    Turn _246_dt__update__tmp_h33 = ((Turn)(java.lang.Object)(boxed652));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.mapTC((((_pat_let_tv19).dtor_activeTurn()).dtor_value()).dtor_parts(), _242_tcId, ((dafny.Function2<Boolean, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, java.util.function.Function<ToolCall, ToolCall>>)(__247_approved0, _248_meta) -> {boolean _247_approved = ((boolean)(java.lang.Object)(__247_approved0));
    return ((java.util.function.Function<ToolCall, ToolCall>)(_249_x_boxed0) -> {
      ToolCall _249_x = ((ToolCall)(java.lang.Object)(_249_x_boxed0));
      return __default.resultConfirmOne(_249_x, _247_approved, _248_meta);
    });}).apply(_241_approved, _240_meta)), boxed653 -> {
      dafny.DafnySequence<? extends Part> _pat_let327_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed653));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let327_0, boxed654 -> {
        dafny.DafnySequence<? extends Part> _250_dt__update_hparts_h9 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed654));
        return Turn.create((_246_dt__update__tmp_h33).dtor_turnId(), (_246_dt__update__tmp_h33).dtor_message(), _250_dt__update_hparts_h9, (_246_dt__update__tmp_h33).dtor_state(), (_246_dt__update__tmp_h33).dtor_usage(), (_246_dt__update__tmp_h33).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed650 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let325_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed650));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let325_0, boxed655 -> {
                agency.open.ahp.AhpSkeleton.Option<Turn> _251_dt__update_hactiveTurn_h12 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed655));
                return ChatState.create((_245_dt__update__tmp_h32).dtor_turns(), (_245_dt__update__tmp_h32).dtor_title(), (_245_dt__update__tmp_h32).dtor_status(), (_245_dt__update__tmp_h32).dtor_modifiedAt(), (_245_dt__update__tmp_h32).dtor_draft(), (_245_dt__update__tmp_h32).dtor_activity(), _251_dt__update_hactiveTurn_h12, (_245_dt__update__tmp_h32).dtor_steeringMessage(), (_245_dt__update__tmp_h32).dtor_queuedMessages(), (_245_dt__update__tmp_h32).dtor_nextCursor());
              }
              )));
            }
            )));
          }
          )));
        }
        )));
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_244_s1, boxed656 -> {
  ChatState _pat_let328_0 = ((ChatState)(java.lang.Object)(boxed656));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let328_0, boxed657 -> {
    ChatState _252_dt__update__tmp_h34 = ((ChatState)(java.lang.Object)(boxed657));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(__default.summaryStatus(_244_s1, false), boxed658 -> {
      int _pat_let329_0 = ((int)(java.lang.Object)(boxed658));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let329_0, boxed659 -> {
        int _253_dt__update_hstatus_h6 = ((int)(java.lang.Object)(boxed659));
        return ChatState.create((_252_dt__update__tmp_h34).dtor_turns(), (_252_dt__update__tmp_h34).dtor_title(), _253_dt__update_hstatus_h6, (_252_dt__update__tmp_h34).dtor_modifiedAt(), (_252_dt__update__tmp_h34).dtor_draft(), (_252_dt__update__tmp_h34).dtor_activity(), (_252_dt__update__tmp_h34).dtor_activeTurn(), (_252_dt__update__tmp_h34).dtor_steeringMessage(), (_252_dt__update__tmp_h34).dtor_queuedMessages(), (_252_dt__update__tmp_h34).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CToolCallContentChanged()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _254___mcc_h69 = ((agency.open.ahp.Chat.ChatAction_CToolCallContentChanged)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _255___mcc_h70 = ((agency.open.ahp.Chat.ChatAction_CToolCallContentChanged)_source0)._toolCallId;
      agency.open.ahp.ConfluxCodec.Json _256___mcc_h71 = ((agency.open.ahp.Chat.ChatAction_CToolCallContentChanged)_source0)._newContent;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _257___mcc_h72 = ((agency.open.ahp.Chat.ChatAction_CToolCallContentChanged)_source0)._meta;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _258_meta = _257___mcc_h72;
      agency.open.ahp.ConfluxCodec.Json _259_c = _256___mcc_h71;
      dafny.DafnySequence<? extends dafny.CodePoint> _260_tcId = _255___mcc_h70;
      dafny.DafnySequence<? extends dafny.CodePoint> _261_id = _254___mcc_h69;
      if (__default.turnMatches(s, _261_id)) {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed660 -> {
  ChatState _pat_let330_0 = ((ChatState)(java.lang.Object)(boxed660));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let330_0, boxed661 -> {
    ChatState _262_dt__update__tmp_h35 = ((ChatState)(java.lang.Object)(boxed661));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv20).dtor_activeTurn()).dtor_value(), boxed663 -> {
  Turn _pat_let332_0 = ((Turn)(java.lang.Object)(boxed663));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let332_0, boxed664 -> {
    Turn _263_dt__update__tmp_h36 = ((Turn)(java.lang.Object)(boxed664));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.mapTC((((_pat_let_tv21).dtor_activeTurn()).dtor_value()).dtor_parts(), _260_tcId, ((dafny.Function2<agency.open.ahp.ConfluxCodec.Json, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, java.util.function.Function<ToolCall, ToolCall>>)(_264_c, _265_meta) -> {return ((java.util.function.Function<ToolCall, ToolCall>)(_266_x_boxed0) -> {
      ToolCall _266_x = ((ToolCall)(java.lang.Object)(_266_x_boxed0));
      return __default.contentChangedOne(_266_x, _264_c, _265_meta);
    });}).apply(_259_c, _258_meta)), boxed665 -> {
      dafny.DafnySequence<? extends Part> _pat_let333_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed665));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let333_0, boxed666 -> {
        dafny.DafnySequence<? extends Part> _267_dt__update_hparts_h10 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed666));
        return Turn.create((_263_dt__update__tmp_h36).dtor_turnId(), (_263_dt__update__tmp_h36).dtor_message(), _267_dt__update_hparts_h10, (_263_dt__update__tmp_h36).dtor_state(), (_263_dt__update__tmp_h36).dtor_usage(), (_263_dt__update__tmp_h36).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed662 -> {
      agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let331_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed662));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let331_0, boxed667 -> {
        agency.open.ahp.AhpSkeleton.Option<Turn> _268_dt__update_hactiveTurn_h13 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed667));
        return ChatState.create((_262_dt__update__tmp_h35).dtor_turns(), (_262_dt__update__tmp_h35).dtor_title(), (_262_dt__update__tmp_h35).dtor_status(), (_262_dt__update__tmp_h35).dtor_modifiedAt(), (_262_dt__update__tmp_h35).dtor_draft(), (_262_dt__update__tmp_h35).dtor_activity(), _268_dt__update_hactiveTurn_h13, (_262_dt__update__tmp_h35).dtor_steeringMessage(), (_262_dt__update__tmp_h35).dtor_queuedMessages(), (_262_dt__update__tmp_h35).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CReasoning()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _269___mcc_h73 = ((agency.open.ahp.Chat.ChatAction_CReasoning)_source0)._turnId;
      dafny.DafnySequence<? extends dafny.CodePoint> _270___mcc_h74 = ((agency.open.ahp.Chat.ChatAction_CReasoning)_source0)._partId;
      dafny.DafnySequence<? extends dafny.CodePoint> _271___mcc_h75 = ((agency.open.ahp.Chat.ChatAction_CReasoning)_source0)._content;
      dafny.DafnySequence<? extends dafny.CodePoint> _272_c = _271___mcc_h75;
      dafny.DafnySequence<? extends dafny.CodePoint> _273_pid = _270___mcc_h74;
      dafny.DafnySequence<? extends dafny.CodePoint> _274_id = _269___mcc_h73;
      if (__default.turnMatches(s, _274_id)) {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed668 -> {
  ChatState _pat_let334_0 = ((ChatState)(java.lang.Object)(boxed668));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let334_0, boxed669 -> {
    ChatState _275_dt__update__tmp_h11 = ((ChatState)(java.lang.Object)(boxed669));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(((_pat_let_tv22).dtor_activeTurn()).dtor_value(), boxed671 -> {
  Turn _pat_let336_0 = ((Turn)(java.lang.Object)(boxed671));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let336_0, boxed672 -> {
    Turn _276_dt__update__tmp_h12 = ((Turn)(java.lang.Object)(boxed672));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.deltaReasoning((((_pat_let_tv23).dtor_activeTurn()).dtor_value()).dtor_parts(), _273_pid, _272_c), boxed673 -> {
      dafny.DafnySequence<? extends Part> _pat_let337_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed673));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let337_0, boxed674 -> {
        dafny.DafnySequence<? extends Part> _277_dt__update_hparts_h1 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed674));
        return Turn.create((_276_dt__update__tmp_h12).dtor_turnId(), (_276_dt__update__tmp_h12).dtor_message(), _277_dt__update_hparts_h1, (_276_dt__update__tmp_h12).dtor_state(), (_276_dt__update__tmp_h12).dtor_usage(), (_276_dt__update__tmp_h12).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed670 -> {
      agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let335_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed670));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let335_0, boxed675 -> {
        agency.open.ahp.AhpSkeleton.Option<Turn> _278_dt__update_hactiveTurn_h4 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed675));
        return ChatState.create((_275_dt__update__tmp_h11).dtor_turns(), (_275_dt__update__tmp_h11).dtor_title(), (_275_dt__update__tmp_h11).dtor_status(), (_275_dt__update__tmp_h11).dtor_modifiedAt(), (_275_dt__update__tmp_h11).dtor_draft(), (_275_dt__update__tmp_h11).dtor_activity(), _278_dt__update_hactiveTurn_h4, (_275_dt__update__tmp_h11).dtor_steeringMessage(), (_275_dt__update__tmp_h11).dtor_queuedMessages(), (_275_dt__update__tmp_h11).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CTruncated()) {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _279___mcc_h76 = ((agency.open.ahp.Chat.ChatAction_CTruncated)_source0)._upToTurnId;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _280_idOpt = _279___mcc_h76;
      if (((_280_idOpt).is_Some()) && (!(__default.hasTurn((s).dtor_turns(), (_280_idOpt).dtor_value())))) {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      } else {
        dafny.DafnySequence<? extends Turn> _281_kept = (((_280_idOpt).is_Some()) ? (__default.keepUpTo((s).dtor_turns(), (_280_idOpt).dtor_value())) : (dafny.DafnySequence.<Turn> empty(Turn._typeDescriptor())));
        ChatState _282_next = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed676 -> {
          ChatState _pat_let338_0 = ((ChatState)(java.lang.Object)(boxed676));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let338_0, boxed677 -> {
            ChatState _283_dt__update__tmp_h37 = ((ChatState)(java.lang.Object)(boxed677));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let((((_280_idOpt).is_Some()) ? ((_pat_let_tv24).dtor_nextCursor()) : (agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)))), boxed678 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let339_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed678));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(_pat_let339_0, boxed679 -> {
                agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _284_dt__update_hnextCursor_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed679));
                return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_None(Turn._typeDescriptor()), boxed680 -> {
                  agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let340_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed680));
                  return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let340_0, boxed681 -> {
                    agency.open.ahp.AhpSkeleton.Option<Turn> _285_dt__update_hactiveTurn_h14 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed681));
                    return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(_281_kept, boxed682 -> {
                      dafny.DafnySequence<? extends Turn> _pat_let341_0 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed682));
                      return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(_pat_let341_0, boxed683 -> {
                        dafny.DafnySequence<? extends Turn> _286_dt__update_hturns_h0 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed683));
                        return ChatState.create(_286_dt__update_hturns_h0, (_283_dt__update__tmp_h37).dtor_title(), (_283_dt__update__tmp_h37).dtor_status(), (_283_dt__update__tmp_h37).dtor_modifiedAt(), (_283_dt__update__tmp_h37).dtor_draft(), (_283_dt__update__tmp_h37).dtor_activity(), _285_dt__update_hactiveTurn_h14, (_283_dt__update__tmp_h37).dtor_steeringMessage(), (_283_dt__update__tmp_h37).dtor_queuedMessages(), _284_dt__update_hnextCursor_h0);
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
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_282_next, boxed684 -> {
  ChatState _pat_let342_0 = ((ChatState)(java.lang.Object)(boxed684));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let342_0, boxed685 -> {
    ChatState _287_dt__update__tmp_h38 = ((ChatState)(java.lang.Object)(boxed685));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(__default.summaryStatus(_282_next, false), boxed686 -> {
      int _pat_let343_0 = ((int)(java.lang.Object)(boxed686));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let343_0, boxed687 -> {
        int _288_dt__update_hstatus_h7 = ((int)(java.lang.Object)(boxed687));
        return ChatState.create((_287_dt__update__tmp_h38).dtor_turns(), (_287_dt__update__tmp_h38).dtor_title(), _288_dt__update_hstatus_h7, (_287_dt__update__tmp_h38).dtor_modifiedAt(), (_287_dt__update__tmp_h38).dtor_draft(), (_287_dt__update__tmp_h38).dtor_activity(), (_287_dt__update__tmp_h38).dtor_activeTurn(), (_287_dt__update__tmp_h38).dtor_steeringMessage(), (_287_dt__update__tmp_h38).dtor_queuedMessages(), (_287_dt__update__tmp_h38).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      }
    } else if (_source0.is_CQueuedReordered()) {
      dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> _289___mcc_h77 = ((agency.open.ahp.Chat.ChatAction_CQueuedReordered)_source0)._order;
      dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> _290_order = _289___mcc_h77;
      return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed688 -> {
  ChatState _pat_let344_0 = ((ChatState)(java.lang.Object)(boxed688));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let344_0, boxed689 -> {
    ChatState _291_dt__update__tmp_h39 = ((ChatState)(java.lang.Object)(boxed689));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(__default.reorderQ((_pat_let_tv25).dtor_queuedMessages(), _290_order), boxed690 -> {
      dafny.DafnySequence<? extends QMsg> _pat_let345_0 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed690));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(_pat_let345_0, boxed691 -> {
        dafny.DafnySequence<? extends QMsg> _292_dt__update_hqueuedMessages_h1 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed691));
        return ChatState.create((_291_dt__update__tmp_h39).dtor_turns(), (_291_dt__update__tmp_h39).dtor_title(), (_291_dt__update__tmp_h39).dtor_status(), (_291_dt__update__tmp_h39).dtor_modifiedAt(), (_291_dt__update__tmp_h39).dtor_draft(), (_291_dt__update__tmp_h39).dtor_activity(), (_291_dt__update__tmp_h39).dtor_activeTurn(), (_291_dt__update__tmp_h39).dtor_steeringMessage(), _292_dt__update_hqueuedMessages_h1, (_291_dt__update__tmp_h39).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_CTurnsLoaded()) {
      dafny.DafnySequence<? extends Turn> _293___mcc_h78 = ((agency.open.ahp.Chat.ChatAction_CTurnsLoaded)_source0)._loaded;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _294___mcc_h79 = ((agency.open.ahp.Chat.ChatAction_CTurnsLoaded)_source0)._cursor;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _295_cursor = _294___mcc_h79;
      dafny.DafnySequence<? extends Turn> _296_loaded = _293___mcc_h78;
      return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed692 -> {
  ChatState _pat_let346_0 = ((ChatState)(java.lang.Object)(boxed692));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let346_0, boxed693 -> {
    ChatState _297_dt__update__tmp_h40 = ((ChatState)(java.lang.Object)(boxed693));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(_295_cursor, boxed694 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let347_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed694));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(_pat_let347_0, boxed695 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _298_dt__update_hnextCursor_h1 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed695));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(__default.dedupPrepend(_296_loaded, (_pat_let_tv26).dtor_turns()), boxed696 -> {
          dafny.DafnySequence<? extends Turn> _pat_let348_0 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed696));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(_pat_let348_0, boxed697 -> {
            dafny.DafnySequence<? extends Turn> _299_dt__update_hturns_h1 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed697));
            return ChatState.create(_299_dt__update_hturns_h1, (_297_dt__update__tmp_h40).dtor_title(), (_297_dt__update__tmp_h40).dtor_status(), (_297_dt__update__tmp_h40).dtor_modifiedAt(), (_297_dt__update__tmp_h40).dtor_draft(), (_297_dt__update__tmp_h40).dtor_activity(), (_297_dt__update__tmp_h40).dtor_activeTurn(), (_297_dt__update__tmp_h40).dtor_steeringMessage(), (_297_dt__update__tmp_h40).dtor_queuedMessages(), _298_dt__update_hnextCursor_h1);
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
    } else if (_source0.is_CInputRequested()) {
      InputReq _300___mcc_h80 = ((agency.open.ahp.Chat.ChatAction_CInputRequested)_source0)._req;
      InputReq _301_req = _300___mcc_h80;
      if (((s).dtor_activeTurn()).is_Some()) {
        Turn _302_at = ((s).dtor_activeTurn()).dtor_value();
        ChatState _303_next = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed698 -> {
          ChatState _pat_let349_0 = ((ChatState)(java.lang.Object)(boxed698));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let349_0, boxed699 -> {
            ChatState _304_dt__update__tmp_h41 = ((ChatState)(java.lang.Object)(boxed699));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_302_at, boxed701 -> {
  Turn _pat_let351_0 = ((Turn)(java.lang.Object)(boxed701));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let351_0, boxed702 -> {
    Turn _305_dt__update__tmp_h42 = ((Turn)(java.lang.Object)(boxed702));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.upsertInputPart((_302_at).dtor_parts(), _301_req), boxed703 -> {
      dafny.DafnySequence<? extends Part> _pat_let352_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed703));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let352_0, boxed704 -> {
        dafny.DafnySequence<? extends Part> _306_dt__update_hparts_h11 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed704));
        return Turn.create((_305_dt__update__tmp_h42).dtor_turnId(), (_305_dt__update__tmp_h42).dtor_message(), _306_dt__update_hparts_h11, (_305_dt__update__tmp_h42).dtor_state(), (_305_dt__update__tmp_h42).dtor_usage(), (_305_dt__update__tmp_h42).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed700 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let350_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed700));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let350_0, boxed705 -> {
                agency.open.ahp.AhpSkeleton.Option<Turn> _307_dt__update_hactiveTurn_h15 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed705));
                return ChatState.create((_304_dt__update__tmp_h41).dtor_turns(), (_304_dt__update__tmp_h41).dtor_title(), (_304_dt__update__tmp_h41).dtor_status(), (_304_dt__update__tmp_h41).dtor_modifiedAt(), (_304_dt__update__tmp_h41).dtor_draft(), (_304_dt__update__tmp_h41).dtor_activity(), _307_dt__update_hactiveTurn_h15, (_304_dt__update__tmp_h41).dtor_steeringMessage(), (_304_dt__update__tmp_h41).dtor_queuedMessages(), (_304_dt__update__tmp_h41).dtor_nextCursor());
              }
              )));
            }
            )));
          }
          )));
        }
        )));
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_303_next, boxed706 -> {
  ChatState _pat_let353_0 = ((ChatState)(java.lang.Object)(boxed706));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let353_0, boxed707 -> {
    ChatState _308_dt__update__tmp_h43 = ((ChatState)(java.lang.Object)(boxed707));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(__default.clearBit(__default.summaryStatus(_303_next, false), __default.READ()), boxed708 -> {
      int _pat_let354_0 = ((int)(java.lang.Object)(boxed708));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let354_0, boxed709 -> {
        int _309_dt__update_hstatus_h8 = ((int)(java.lang.Object)(boxed709));
        return ChatState.create((_308_dt__update__tmp_h43).dtor_turns(), (_308_dt__update__tmp_h43).dtor_title(), _309_dt__update_hstatus_h8, (_308_dt__update__tmp_h43).dtor_modifiedAt(), (_308_dt__update__tmp_h43).dtor_draft(), (_308_dt__update__tmp_h43).dtor_activity(), (_308_dt__update__tmp_h43).dtor_activeTurn(), (_308_dt__update__tmp_h43).dtor_steeringMessage(), (_308_dt__update__tmp_h43).dtor_queuedMessages(), (_308_dt__update__tmp_h43).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CInputAnswerChanged()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _310___mcc_h81 = ((agency.open.ahp.Chat.ChatAction_CInputAnswerChanged)_source0)._requestId;
      dafny.DafnySequence<? extends dafny.CodePoint> _311___mcc_h82 = ((agency.open.ahp.Chat.ChatAction_CInputAnswerChanged)_source0)._questionId;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _312___mcc_h83 = ((agency.open.ahp.Chat.ChatAction_CInputAnswerChanged)_source0)._answer;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _313_ans = _312___mcc_h83;
      dafny.DafnySequence<? extends dafny.CodePoint> _314_qId = _311___mcc_h82;
      dafny.DafnySequence<? extends dafny.CodePoint> _315_reqId = _310___mcc_h81;
      if ((((s).dtor_activeTurn()).is_Some()) && (__default.hasOpenReq((((s).dtor_activeTurn()).dtor_value()).dtor_parts(), _315_reqId))) {
        Turn _316_at = ((s).dtor_activeTurn()).dtor_value();
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed710 -> {
  ChatState _pat_let355_0 = ((ChatState)(java.lang.Object)(boxed710));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let355_0, boxed711 -> {
    ChatState _317_dt__update__tmp_h44 = ((ChatState)(java.lang.Object)(boxed711));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_316_at, boxed713 -> {
  Turn _pat_let357_0 = ((Turn)(java.lang.Object)(boxed713));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let357_0, boxed714 -> {
    Turn _318_dt__update__tmp_h45 = ((Turn)(java.lang.Object)(boxed714));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.changeAnswerPart((_316_at).dtor_parts(), _315_reqId, _314_qId, _313_ans), boxed715 -> {
      dafny.DafnySequence<? extends Part> _pat_let358_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed715));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let358_0, boxed716 -> {
        dafny.DafnySequence<? extends Part> _319_dt__update_hparts_h12 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed716));
        return Turn.create((_318_dt__update__tmp_h45).dtor_turnId(), (_318_dt__update__tmp_h45).dtor_message(), _319_dt__update_hparts_h12, (_318_dt__update__tmp_h45).dtor_state(), (_318_dt__update__tmp_h45).dtor_usage(), (_318_dt__update__tmp_h45).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed712 -> {
      agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let356_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed712));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let356_0, boxed717 -> {
        agency.open.ahp.AhpSkeleton.Option<Turn> _320_dt__update_hactiveTurn_h16 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed717));
        return ChatState.create((_317_dt__update__tmp_h44).dtor_turns(), (_317_dt__update__tmp_h44).dtor_title(), (_317_dt__update__tmp_h44).dtor_status(), (_317_dt__update__tmp_h44).dtor_modifiedAt(), (_317_dt__update__tmp_h44).dtor_draft(), (_317_dt__update__tmp_h44).dtor_activity(), _320_dt__update_hactiveTurn_h16, (_317_dt__update__tmp_h44).dtor_steeringMessage(), (_317_dt__update__tmp_h44).dtor_queuedMessages(), (_317_dt__update__tmp_h44).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CInputCompleted()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _321___mcc_h84 = ((agency.open.ahp.Chat.ChatAction_CInputCompleted)_source0)._requestId;
      dafny.DafnySequence<? extends dafny.CodePoint> _322___mcc_h85 = ((agency.open.ahp.Chat.ChatAction_CInputCompleted)_source0)._response;
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _323___mcc_h86 = ((agency.open.ahp.Chat.ChatAction_CInputCompleted)_source0)._answers;
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _324_answers = _323___mcc_h86;
      dafny.DafnySequence<? extends dafny.CodePoint> _325_resp = _322___mcc_h85;
      dafny.DafnySequence<? extends dafny.CodePoint> _326_reqId = _321___mcc_h84;
      if ((((s).dtor_activeTurn()).is_Some()) && (__default.hasOpenReq((((s).dtor_activeTurn()).dtor_value()).dtor_parts(), _326_reqId))) {
        Turn _327_at = ((s).dtor_activeTurn()).dtor_value();
        ChatState _328_next = ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed718 -> {
          ChatState _pat_let359_0 = ((ChatState)(java.lang.Object)(boxed718));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let359_0, boxed719 -> {
            ChatState _329_dt__update__tmp_h46 = ((ChatState)(java.lang.Object)(boxed719));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_327_at, boxed721 -> {
  Turn _pat_let361_0 = ((Turn)(java.lang.Object)(boxed721));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let361_0, boxed722 -> {
    Turn _330_dt__update__tmp_h47 = ((Turn)(java.lang.Object)(boxed722));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(__default.completeAnswerPart((_327_at).dtor_parts(), _326_reqId, _325_resp, _324_answers), boxed723 -> {
      dafny.DafnySequence<? extends Part> _pat_let362_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed723));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let362_0, boxed724 -> {
        dafny.DafnySequence<? extends Part> _331_dt__update_hparts_h13 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed724));
        return Turn.create((_330_dt__update__tmp_h47).dtor_turnId(), (_330_dt__update__tmp_h47).dtor_message(), _331_dt__update_hparts_h13, (_330_dt__update__tmp_h47).dtor_state(), (_330_dt__update__tmp_h47).dtor_usage(), (_330_dt__update__tmp_h47).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed720 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let360_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed720));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let360_0, boxed725 -> {
                agency.open.ahp.AhpSkeleton.Option<Turn> _332_dt__update_hactiveTurn_h17 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed725));
                return ChatState.create((_329_dt__update__tmp_h46).dtor_turns(), (_329_dt__update__tmp_h46).dtor_title(), (_329_dt__update__tmp_h46).dtor_status(), (_329_dt__update__tmp_h46).dtor_modifiedAt(), (_329_dt__update__tmp_h46).dtor_draft(), (_329_dt__update__tmp_h46).dtor_activity(), _332_dt__update_hactiveTurn_h17, (_329_dt__update__tmp_h46).dtor_steeringMessage(), (_329_dt__update__tmp_h46).dtor_queuedMessages(), (_329_dt__update__tmp_h46).dtor_nextCursor());
              }
              )));
            }
            )));
          }
          )));
        }
        )));
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_328_next, boxed726 -> {
  ChatState _pat_let363_0 = ((ChatState)(java.lang.Object)(boxed726));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let363_0, boxed727 -> {
    ChatState _333_dt__update__tmp_h48 = ((ChatState)(java.lang.Object)(boxed727));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(__default.summaryStatus(_328_next, false), boxed728 -> {
      int _pat_let364_0 = ((int)(java.lang.Object)(boxed728));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let364_0, boxed729 -> {
        int _334_dt__update_hstatus_h9 = ((int)(java.lang.Object)(boxed729));
        return ChatState.create((_333_dt__update__tmp_h48).dtor_turns(), (_333_dt__update__tmp_h48).dtor_title(), _334_dt__update_hstatus_h9, (_333_dt__update__tmp_h48).dtor_modifiedAt(), (_333_dt__update__tmp_h48).dtor_draft(), (_333_dt__update__tmp_h48).dtor_activity(), (_333_dt__update__tmp_h48).dtor_activeTurn(), (_333_dt__update__tmp_h48).dtor_steeringMessage(), (_333_dt__update__tmp_h48).dtor_queuedMessages(), (_333_dt__update__tmp_h48).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_CPendingMessageSet()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _335___mcc_h87 = ((agency.open.ahp.Chat.ChatAction_CPendingMessageSet)_source0)._kind;
      dafny.DafnySequence<? extends dafny.CodePoint> _336___mcc_h88 = ((agency.open.ahp.Chat.ChatAction_CPendingMessageSet)_source0)._id;
      agency.open.ahp.ConfluxCodec.Json _337___mcc_h89 = ((agency.open.ahp.Chat.ChatAction_CPendingMessageSet)_source0)._message;
      agency.open.ahp.ConfluxCodec.Json _338_msg = _337___mcc_h89;
      dafny.DafnySequence<? extends dafny.CodePoint> _339_id = _336___mcc_h88;
      dafny.DafnySequence<? extends dafny.CodePoint> _340_kind = _335___mcc_h87;
      if ((_340_kind).equals(dafny.DafnySequence.asUnicodeString("steering"))) {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed730 -> {
  ChatState _pat_let365_0 = ((ChatState)(java.lang.Object)(boxed730));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let365_0, boxed731 -> {
    ChatState _341_dt__update__tmp_h49 = ((ChatState)(java.lang.Object)(boxed731));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<QMsg>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<QMsg>create_Some(QMsg._typeDescriptor(), QMsg.create(_339_id, _338_msg)), boxed732 -> {
      agency.open.ahp.AhpSkeleton.Option<QMsg> _pat_let366_0 = ((agency.open.ahp.AhpSkeleton.Option<QMsg>)(java.lang.Object)(boxed732));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<QMsg>, ChatState>Let(_pat_let366_0, boxed733 -> {
        agency.open.ahp.AhpSkeleton.Option<QMsg> _342_dt__update_hsteeringMessage_h1 = ((agency.open.ahp.AhpSkeleton.Option<QMsg>)(java.lang.Object)(boxed733));
        return ChatState.create((_341_dt__update__tmp_h49).dtor_turns(), (_341_dt__update__tmp_h49).dtor_title(), (_341_dt__update__tmp_h49).dtor_status(), (_341_dt__update__tmp_h49).dtor_modifiedAt(), (_341_dt__update__tmp_h49).dtor_draft(), (_341_dt__update__tmp_h49).dtor_activity(), (_341_dt__update__tmp_h49).dtor_activeTurn(), _342_dt__update_hsteeringMessage_h1, (_341_dt__update__tmp_h49).dtor_queuedMessages(), (_341_dt__update__tmp_h49).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed734 -> {
  ChatState _pat_let367_0 = ((ChatState)(java.lang.Object)(boxed734));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let367_0, boxed735 -> {
    ChatState _343_dt__update__tmp_h50 = ((ChatState)(java.lang.Object)(boxed735));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(__default.upsertQ((_pat_let_tv27).dtor_queuedMessages(), QMsg.create(_339_id, _338_msg)), boxed736 -> {
      dafny.DafnySequence<? extends QMsg> _pat_let368_0 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed736));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(_pat_let368_0, boxed737 -> {
        dafny.DafnySequence<? extends QMsg> _344_dt__update_hqueuedMessages_h2 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed737));
        return ChatState.create((_343_dt__update__tmp_h50).dtor_turns(), (_343_dt__update__tmp_h50).dtor_title(), (_343_dt__update__tmp_h50).dtor_status(), (_343_dt__update__tmp_h50).dtor_modifiedAt(), (_343_dt__update__tmp_h50).dtor_draft(), (_343_dt__update__tmp_h50).dtor_activity(), (_343_dt__update__tmp_h50).dtor_activeTurn(), (_343_dt__update__tmp_h50).dtor_steeringMessage(), _344_dt__update_hqueuedMessages_h2, (_343_dt__update__tmp_h50).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      }
    } else if (_source0.is_CPendingMessageRemoved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _345___mcc_h90 = ((agency.open.ahp.Chat.ChatAction_CPendingMessageRemoved)_source0)._kind;
      dafny.DafnySequence<? extends dafny.CodePoint> _346___mcc_h91 = ((agency.open.ahp.Chat.ChatAction_CPendingMessageRemoved)_source0)._id;
      dafny.DafnySequence<? extends dafny.CodePoint> _347_id = _346___mcc_h91;
      dafny.DafnySequence<? extends dafny.CodePoint> _348_kind = _345___mcc_h90;
      if ((_348_kind).equals(dafny.DafnySequence.asUnicodeString("steering"))) {
        if ((((s).dtor_steeringMessage()).is_Some()) && (((((s).dtor_steeringMessage()).dtor_value()).dtor_id()).equals(_347_id))) {
          return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed738 -> {
  ChatState _pat_let369_0 = ((ChatState)(java.lang.Object)(boxed738));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let369_0, boxed739 -> {
    ChatState _349_dt__update__tmp_h51 = ((ChatState)(java.lang.Object)(boxed739));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<QMsg>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<QMsg>create_None(QMsg._typeDescriptor()), boxed740 -> {
      agency.open.ahp.AhpSkeleton.Option<QMsg> _pat_let370_0 = ((agency.open.ahp.AhpSkeleton.Option<QMsg>)(java.lang.Object)(boxed740));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<QMsg>, ChatState>Let(_pat_let370_0, boxed741 -> {
        agency.open.ahp.AhpSkeleton.Option<QMsg> _350_dt__update_hsteeringMessage_h2 = ((agency.open.ahp.AhpSkeleton.Option<QMsg>)(java.lang.Object)(boxed741));
        return ChatState.create((_349_dt__update__tmp_h51).dtor_turns(), (_349_dt__update__tmp_h51).dtor_title(), (_349_dt__update__tmp_h51).dtor_status(), (_349_dt__update__tmp_h51).dtor_modifiedAt(), (_349_dt__update__tmp_h51).dtor_draft(), (_349_dt__update__tmp_h51).dtor_activity(), (_349_dt__update__tmp_h51).dtor_activeTurn(), _350_dt__update_hsteeringMessage_h2, (_349_dt__update__tmp_h51).dtor_queuedMessages(), (_349_dt__update__tmp_h51).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
        } else {
          return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
        }
      } else {
        return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(s, boxed742 -> {
  ChatState _pat_let371_0 = ((ChatState)(java.lang.Object)(boxed742));
  return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let371_0, boxed743 -> {
    ChatState _351_dt__update__tmp_h52 = ((ChatState)(java.lang.Object)(boxed743));
    return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(__default.removeQ((_pat_let_tv28).dtor_queuedMessages(), _347_id), boxed744 -> {
      dafny.DafnySequence<? extends QMsg> _pat_let372_0 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed744));
      return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(_pat_let372_0, boxed745 -> {
        dafny.DafnySequence<? extends QMsg> _352_dt__update_hqueuedMessages_h3 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed745));
        return ChatState.create((_351_dt__update__tmp_h52).dtor_turns(), (_351_dt__update__tmp_h52).dtor_title(), (_351_dt__update__tmp_h52).dtor_status(), (_351_dt__update__tmp_h52).dtor_modifiedAt(), (_351_dt__update__tmp_h52).dtor_draft(), (_351_dt__update__tmp_h52).dtor_activity(), (_351_dt__update__tmp_h52).dtor_activeTurn(), (_351_dt__update__tmp_h52).dtor_steeringMessage(), _352_dt__update_hqueuedMessages_h3, (_351_dt__update__tmp_h52).dtor_nextCursor());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      }
    } else {
      agency.open.ahp.ConfluxCodec.Json _353___mcc_h92 = ((agency.open.ahp.Chat.ChatAction_CUnknown)_source0)._raw;
      return dafny.Tuple2.<ChatState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
    }
  }
  public static ChatState apply1(ChatState s, ChatAction a)
  {
    return (__default.ApplyToChat(s, a, java.math.BigInteger.valueOf(9999L))).dtor__0();
  }
  public static ChatState fold(ChatState s, dafny.DafnySequence<? extends ChatAction> acts)
  {
    return agency.open.ahp.ConfluxContract.__default.<ChatState, ChatAction>Fold(ChatState._typeDescriptor(), ChatAction._typeDescriptor(), __default::apply1, s, acts);
  }
  public static ChatState C0() {
    return ChatState.create(dafny.DafnySequence.<Turn> empty(Turn._typeDescriptor()), dafny.DafnySequence.asUnicodeString("Chat"), 1, dafny.DafnySequence.asUnicodeString("t1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<Turn>create_None(Turn._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<QMsg>create_None(QMsg._typeDescriptor()), dafny.DafnySequence.<QMsg> empty(QMsg._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)));
  }
  public static Turn T0(dafny.DafnySequence<? extends dafny.CodePoint> id) {
    return Turn.create(id, agency.open.ahp.ConfluxCodec.Json.create_JNull(), dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), dafny.DafnySequence.asUnicodeString("in-progress"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
  }
  public static Turn TDone(dafny.DafnySequence<? extends dafny.CodePoint> id) {
    return Turn.create(id, agency.open.ahp.ConfluxCodec.Json.create_JNull(), dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), dafny.DafnySequence.asUnicodeString("complete"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
  }
  public static ToolCall TC0(dafny.DafnySequence<? extends dafny.CodePoint> id, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> intent)
  {
    return ToolCall.create(id, dafny.DafnySequence.asUnicodeString("bash"), dafny.DafnySequence.asUnicodeString("Run Command"), dafny.DafnySequence.asUnicodeString("streaming"), intent, agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>create_None(ToolCallContributor._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), dafny.DafnySequence.asUnicodeString(""), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)));
  }
  public static dafny.DafnySequence<? extends Part> activeParts(ChatState s) {
    if (((s).dtor_activeTurn()).is_Some()) {
      return (((s).dtor_activeTurn()).dtor_value()).dtor_parts();
    } else {
      return dafny.DafnySequence.<Part> empty(Part._typeDescriptor());
    }
  }
  public static agency.open.ahp.AhpSkeleton.Option<ToolCall> firstActiveToolCall(ChatState s) {
    dafny.DafnySequence<? extends Part> _0_ps = __default.activeParts(s);
    if (((java.math.BigInteger.valueOf((_0_ps).length())).signum() == 1) && ((((Part)(java.lang.Object)((_0_ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).is_PToolCall())) {
      return agency.open.ahp.AhpSkeleton.Option.<ToolCall>create_Some(ToolCall._typeDescriptor(), (((Part)(java.lang.Object)((_0_ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).dtor_tc());
    } else {
      return agency.open.ahp.AhpSkeleton.Option.<ToolCall>create_None(ToolCall._typeDescriptor());
    }
  }
  public static agency.open.ahp.AhpSkeleton.Option<ToolCall> firstStoredToolCall(ChatState s) {
    if ((java.math.BigInteger.valueOf(((s).dtor_turns()).length())).signum() == 0) {
      return agency.open.ahp.AhpSkeleton.Option.<ToolCall>create_None(ToolCall._typeDescriptor());
    } else {
      dafny.DafnySequence<? extends Part> _0_ps = (((Turn)(java.lang.Object)(((s).dtor_turns()).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).dtor_parts();
      if (((java.math.BigInteger.valueOf((_0_ps).length())).signum() == 1) && ((((Part)(java.lang.Object)((_0_ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).is_PToolCall())) {
        return agency.open.ahp.AhpSkeleton.Option.<ToolCall>create_Some(ToolCall._typeDescriptor(), (((Part)(java.lang.Object)((_0_ps).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).dtor_tc());
      } else {
        return agency.open.ahp.AhpSkeleton.Option.<ToolCall>create_None(ToolCall._typeDescriptor());
      }
    }
  }
  public static java.math.BigInteger runScalarTurns()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    if(true) {
      pass = java.math.BigInteger.ZERO;
      ChatState _0_s;
      _0_s = __default.C0();
      if (java.util.Objects.equals((__default.apply1(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed746 -> {
        ChatState _pat_let373_0 = ((ChatState)(java.lang.Object)(boxed746));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let373_0, boxed747 -> {
          ChatState _1_dt__update__tmp_h0 = ((ChatState)(java.lang.Object)(boxed747));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("x")), boxed748 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let374_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed748));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(_pat_let374_0, boxed749 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _2_dt__update_hdraft_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed749));
              return ChatState.create((_1_dt__update__tmp_h0).dtor_turns(), (_1_dt__update__tmp_h0).dtor_title(), (_1_dt__update__tmp_h0).dtor_status(), (_1_dt__update__tmp_h0).dtor_modifiedAt(), _2_dt__update_hdraft_h0, (_1_dt__update__tmp_h0).dtor_activity(), (_1_dt__update__tmp_h0).dtor_activeTurn(), (_1_dt__update__tmp_h0).dtor_steeringMessage(), (_1_dt__update__tmp_h0).dtor_queuedMessages(), (_1_dt__update__tmp_h0).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), ChatAction.create_CDraftChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))))).dtor_draft(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals((__default.apply1(_0_s, ChatAction.create_CDraftChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("hi"))))).dtor_draft(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("hi")))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals((__default.apply1(_0_s, ChatAction.create_CActivityChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Running terminal command"))))).dtor_activity(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Running terminal command")))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals((__default.apply1(_0_s, ChatAction.create_CActivityChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))))).dtor_activity(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals((__default.apply1(_0_s, ChatAction.create_CTurnStarted(dafny.DafnySequence.asUnicodeString("turn-1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("Hello")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))))).dtor_activeTurn(), agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), Turn.create(dafny.DafnySequence.asUnicodeString("turn-1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("Hello")), dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), dafny.DafnySequence.asUnicodeString("in-progress"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _3_act1;
      ChatState _4_dt__update__tmp_h1 = _0_s;
      agency.open.ahp.AhpSkeleton.Option<Turn> _5_dt__update_hactiveTurn_h0 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), __default.T0(dafny.DafnySequence.asUnicodeString("t1")));
      _3_act1 = ChatState.create((_4_dt__update__tmp_h1).dtor_turns(), (_4_dt__update__tmp_h1).dtor_title(), (_4_dt__update__tmp_h1).dtor_status(), (_4_dt__update__tmp_h1).dtor_modifiedAt(), (_4_dt__update__tmp_h1).dtor_draft(), (_4_dt__update__tmp_h1).dtor_activity(), _5_dt__update_hactiveTurn_h0, (_4_dt__update__tmp_h1).dtor_steeringMessage(), (_4_dt__update__tmp_h1).dtor_queuedMessages(), (_4_dt__update__tmp_h1).dtor_nextCursor());
      if (java.util.Objects.equals(__default.apply1(_3_act1, ChatAction.create_CTurnComplete(dafny.DafnySequence.asUnicodeString("wrong-turn"))), _3_act1)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_3_act1, ChatAction.create_CUsage(dafny.DafnySequence.asUnicodeString("wrong-turn"), agency.open.ahp.ConfluxCodec.Json.create_JNull())), _3_act1)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_3_act1, ChatAction.create_CToolCallStart(dafny.DafnySequence.asUnicodeString("wrong-turn"), dafny.DafnySequence.asUnicodeString("tc-1"), dafny.DafnySequence.asUnicodeString("bash"), dafny.DafnySequence.asUnicodeString("Run Command"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>create_None(ToolCallContributor._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), _3_act1)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed750 -> {
        ChatState _pat_let375_0 = ((ChatState)(java.lang.Object)(boxed750));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let375_0, boxed751 -> {
          ChatState _6_dt__update__tmp_h2 = ((ChatState)(java.lang.Object)(boxed751));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), Turn.create(dafny.DafnySequence.asUnicodeString("t1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("T")), dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), dafny.DafnySequence.asUnicodeString("in-progress"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), boxed752 -> {
            agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let376_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed752));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let376_0, boxed753 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _7_dt__update_hactiveTurn_h1 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed753));
              return ChatState.create((_6_dt__update__tmp_h2).dtor_turns(), (_6_dt__update__tmp_h2).dtor_title(), (_6_dt__update__tmp_h2).dtor_status(), (_6_dt__update__tmp_h2).dtor_modifiedAt(), (_6_dt__update__tmp_h2).dtor_draft(), (_6_dt__update__tmp_h2).dtor_activity(), _7_dt__update_hactiveTurn_h1, (_6_dt__update__tmp_h2).dtor_steeringMessage(), (_6_dt__update__tmp_h2).dtor_queuedMessages(), (_6_dt__update__tmp_h2).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), ChatAction.create_CTurnComplete(dafny.DafnySequence.asUnicodeString("t1"))), ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed754 -> {
        ChatState _pat_let377_0 = ((ChatState)(java.lang.Object)(boxed754));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let377_0, boxed755 -> {
          ChatState _8_dt__update__tmp_h3 = ((ChatState)(java.lang.Object)(boxed755));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), Turn.create(dafny.DafnySequence.asUnicodeString("t1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("T")), dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), dafny.DafnySequence.asUnicodeString("complete"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), boxed756 -> {
            dafny.DafnySequence<? extends Turn> _pat_let378_0 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed756));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(_pat_let378_0, boxed757 -> {
              dafny.DafnySequence<? extends Turn> _9_dt__update_hturns_h0 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed757));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_None(Turn._typeDescriptor()), boxed758 -> {
                agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let379_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed758));
                return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let379_0, boxed759 -> {
                  agency.open.ahp.AhpSkeleton.Option<Turn> _10_dt__update_hactiveTurn_h2 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed759));
                  return ChatState.create(_9_dt__update_hturns_h0, (_8_dt__update__tmp_h3).dtor_title(), (_8_dt__update__tmp_h3).dtor_status(), (_8_dt__update__tmp_h3).dtor_modifiedAt(), (_8_dt__update__tmp_h3).dtor_draft(), (_8_dt__update__tmp_h3).dtor_activity(), _10_dt__update_hactiveTurn_h2, (_8_dt__update__tmp_h3).dtor_steeringMessage(), (_8_dt__update__tmp_h3).dtor_queuedMessages(), (_8_dt__update__tmp_h3).dtor_nextCursor());
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
      if (((__default.apply1(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed760 -> {
        ChatState _pat_let380_0 = ((ChatState)(java.lang.Object)(boxed760));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let380_0, boxed761 -> {
          ChatState _11_dt__update__tmp_h4 = ((ChatState)(java.lang.Object)(boxed761));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), __default.T0(dafny.DafnySequence.asUnicodeString("t1"))), boxed762 -> {
            agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let381_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed762));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let381_0, boxed763 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _12_dt__update_hactiveTurn_h3 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed763));
              return ChatState.create((_11_dt__update__tmp_h4).dtor_turns(), (_11_dt__update__tmp_h4).dtor_title(), (_11_dt__update__tmp_h4).dtor_status(), (_11_dt__update__tmp_h4).dtor_modifiedAt(), (_11_dt__update__tmp_h4).dtor_draft(), (_11_dt__update__tmp_h4).dtor_activity(), _12_dt__update_hactiveTurn_h3, (_11_dt__update__tmp_h4).dtor_steeringMessage(), (_11_dt__update__tmp_h4).dtor_queuedMessages(), (_11_dt__update__tmp_h4).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), ChatAction.create_CTurnCancelled(dafny.DafnySequence.asUnicodeString("t1")))).dtor_turns()).equals(dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), Turn.create(dafny.DafnySequence.asUnicodeString("t1"), agency.open.ahp.ConfluxCodec.Json.create_JNull(), dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), dafny.DafnySequence.asUnicodeString("cancelled"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals((__default.apply1(_3_act1, ChatAction.create_CUsage(dafny.DafnySequence.asUnicodeString("t1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("u"))))).dtor_activeTurn(), agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed764 -> {
  Turn _pat_let382_0 = ((Turn)(java.lang.Object)(boxed764));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let382_0, boxed765 -> {
    Turn _13_dt__update__tmp_h5 = ((Turn)(java.lang.Object)(boxed765));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Turn>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("u"))), boxed766 -> {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let383_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed766));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, Turn>Let(_pat_let383_0, boxed767 -> {
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _14_dt__update_husage_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed767));
        return Turn.create((_13_dt__update__tmp_h5).dtor_turnId(), (_13_dt__update__tmp_h5).dtor_message(), (_13_dt__update__tmp_h5).dtor_parts(), (_13_dt__update__tmp_h5).dtor_state(), _14_dt__update_husage_h0, (_13_dt__update__tmp_h5).dtor_error());
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
      if (((__default.apply1(_3_act1, ChatAction.create_CError(dafny.DafnySequence.asUnicodeString("t1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("boom"))))).dtor_turns()).equals(dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), Turn.create(dafny.DafnySequence.asUnicodeString("t1"), agency.open.ahp.ConfluxCodec.Json.create_JNull(), dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), dafny.DafnySequence.asUnicodeString("error"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("boom"))))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _15_withMd;
      _15_withMd = __default.apply1(_3_act1, ChatAction.create_CResponsePart(dafny.DafnySequence.asUnicodeString("t1"), Part.create_PMarkdown(dafny.DafnySequence.asUnicodeString("md-1"), dafny.DafnySequence.asUnicodeString(""))));
      if (java.util.Objects.equals((_15_withMd).dtor_activeTurn(), agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed768 -> {
  Turn _pat_let384_0 = ((Turn)(java.lang.Object)(boxed768));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let384_0, boxed769 -> {
    Turn _16_dt__update__tmp_h6 = ((Turn)(java.lang.Object)(boxed769));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PMarkdown(dafny.DafnySequence.asUnicodeString("md-1"), dafny.DafnySequence.asUnicodeString(""))), boxed770 -> {
      dafny.DafnySequence<? extends Part> _pat_let385_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed770));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let385_0, boxed771 -> {
        dafny.DafnySequence<? extends Part> _17_dt__update_hparts_h0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed771));
        return Turn.create((_16_dt__update__tmp_h6).dtor_turnId(), (_16_dt__update__tmp_h6).dtor_message(), _17_dt__update_hparts_h0, (_16_dt__update__tmp_h6).dtor_state(), (_16_dt__update__tmp_h6).dtor_usage(), (_16_dt__update__tmp_h6).dtor_error());
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
      if (java.util.Objects.equals((__default.apply1(_15_withMd, ChatAction.create_CDelta(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("md-1"), dafny.DafnySequence.asUnicodeString("Hello from chat")))).dtor_activeTurn(), agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed772 -> {
  Turn _pat_let386_0 = ((Turn)(java.lang.Object)(boxed772));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let386_0, boxed773 -> {
    Turn _18_dt__update__tmp_h7 = ((Turn)(java.lang.Object)(boxed773));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PMarkdown(dafny.DafnySequence.asUnicodeString("md-1"), dafny.DafnySequence.asUnicodeString("Hello from chat"))), boxed774 -> {
      dafny.DafnySequence<? extends Part> _pat_let387_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed774));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let387_0, boxed775 -> {
        dafny.DafnySequence<? extends Part> _19_dt__update_hparts_h1 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed775));
        return Turn.create((_18_dt__update__tmp_h7).dtor_turnId(), (_18_dt__update__tmp_h7).dtor_message(), _19_dt__update_hparts_h1, (_18_dt__update__tmp_h7).dtor_state(), (_18_dt__update__tmp_h7).dtor_usage(), (_18_dt__update__tmp_h7).dtor_error());
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
      if (java.util.Objects.equals((__default.apply1(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_3_act1, boxed776 -> {
        ChatState _pat_let388_0 = ((ChatState)(java.lang.Object)(boxed776));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let388_0, boxed777 -> {
          ChatState _20_dt__update__tmp_h8 = ((ChatState)(java.lang.Object)(boxed777));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed779 -> {
  Turn _pat_let390_0 = ((Turn)(java.lang.Object)(boxed779));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let390_0, boxed780 -> {
    Turn _21_dt__update__tmp_h9 = ((Turn)(java.lang.Object)(boxed780));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PReasoning(dafny.DafnySequence.asUnicodeString("r-1"), dafny.DafnySequence.asUnicodeString("th"))), boxed781 -> {
      dafny.DafnySequence<? extends Part> _pat_let391_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed781));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let391_0, boxed782 -> {
        dafny.DafnySequence<? extends Part> _22_dt__update_hparts_h2 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed782));
        return Turn.create((_21_dt__update__tmp_h9).dtor_turnId(), (_21_dt__update__tmp_h9).dtor_message(), _22_dt__update_hparts_h2, (_21_dt__update__tmp_h9).dtor_state(), (_21_dt__update__tmp_h9).dtor_usage(), (_21_dt__update__tmp_h9).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
)))), boxed778 -> {
            agency.open.ahp.AhpSkeleton.Option<Turn> _pat_let389_0 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed778));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Turn>, ChatState>Let(_pat_let389_0, boxed783 -> {
              agency.open.ahp.AhpSkeleton.Option<Turn> _23_dt__update_hactiveTurn_h4 = ((agency.open.ahp.AhpSkeleton.Option<Turn>)(java.lang.Object)(boxed783));
              return ChatState.create((_20_dt__update__tmp_h8).dtor_turns(), (_20_dt__update__tmp_h8).dtor_title(), (_20_dt__update__tmp_h8).dtor_status(), (_20_dt__update__tmp_h8).dtor_modifiedAt(), (_20_dt__update__tmp_h8).dtor_draft(), (_20_dt__update__tmp_h8).dtor_activity(), _23_dt__update_hactiveTurn_h4, (_20_dt__update__tmp_h8).dtor_steeringMessage(), (_20_dt__update__tmp_h8).dtor_queuedMessages(), (_20_dt__update__tmp_h8).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), ChatAction.create_CDelta(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("r-1"), dafny.DafnySequence.asUnicodeString("inking")))).dtor_activeTurn(), agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed784 -> {
  Turn _pat_let392_0 = ((Turn)(java.lang.Object)(boxed784));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let392_0, boxed785 -> {
    Turn _24_dt__update__tmp_h10 = ((Turn)(java.lang.Object)(boxed785));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PReasoning(dafny.DafnySequence.asUnicodeString("r-1"), dafny.DafnySequence.asUnicodeString("th"))), boxed786 -> {
      dafny.DafnySequence<? extends Part> _pat_let393_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed786));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let393_0, boxed787 -> {
        dafny.DafnySequence<? extends Part> _25_dt__update_hparts_h3 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed787));
        return Turn.create((_24_dt__update__tmp_h10).dtor_turnId(), (_24_dt__update__tmp_h10).dtor_message(), _25_dt__update_hparts_h3, (_24_dt__update__tmp_h10).dtor_state(), (_24_dt__update__tmp_h10).dtor_usage(), (_24_dt__update__tmp_h10).dtor_error());
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
    }
    return pass;
  }
  public static java.math.BigInteger runToolCalls()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    if(true) {
      pass = java.math.BigInteger.ZERO;
      ChatState _0_s;
      _0_s = __default.C0();
      ChatState _1_act1;
      ChatState _2_dt__update__tmp_h0 = _0_s;
      agency.open.ahp.AhpSkeleton.Option<Turn> _3_dt__update_hactiveTurn_h0 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), __default.T0(dafny.DafnySequence.asUnicodeString("t1")));
      _1_act1 = ChatState.create((_2_dt__update__tmp_h0).dtor_turns(), (_2_dt__update__tmp_h0).dtor_title(), (_2_dt__update__tmp_h0).dtor_status(), (_2_dt__update__tmp_h0).dtor_modifiedAt(), (_2_dt__update__tmp_h0).dtor_draft(), (_2_dt__update__tmp_h0).dtor_activity(), _3_dt__update_hactiveTurn_h0, (_2_dt__update__tmp_h0).dtor_steeringMessage(), (_2_dt__update__tmp_h0).dtor_queuedMessages(), (_2_dt__update__tmp_h0).dtor_nextCursor());
      ChatState _4_started;
      _4_started = __default.apply1(_1_act1, ChatAction.create_CToolCallStart(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), dafny.DafnySequence.asUnicodeString("bash"), dafny.DafnySequence.asUnicodeString("Run Command"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("List the files in the current directory")), agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>create_None(ToolCallContributor._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())));
      if (java.util.Objects.equals((_4_started).dtor_activeTurn(), agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed788 -> {
  Turn _pat_let394_0 = ((Turn)(java.lang.Object)(boxed788));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let394_0, boxed789 -> {
    Turn _5_dt__update__tmp_h1 = ((Turn)(java.lang.Object)(boxed789));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PToolCall(__default.TC0(dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("List the files in the current directory"))))), boxed790 -> {
      dafny.DafnySequence<? extends Part> _pat_let395_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed790));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let395_0, boxed791 -> {
        dafny.DafnySequence<? extends Part> _6_dt__update_hparts_h0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed791));
        return Turn.create((_5_dt__update__tmp_h1).dtor_turnId(), (_5_dt__update__tmp_h1).dtor_message(), _6_dt__update_hparts_h0, (_5_dt__update__tmp_h1).dtor_state(), (_5_dt__update__tmp_h1).dtor_usage(), (_5_dt__update__tmp_h1).dtor_error());
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
      ChatState _7_afterDelta;
      _7_afterDelta = __default.apply1(_4_started, ChatAction.create_CToolCallDelta(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), dafny.DafnySequence.asUnicodeString("ls -la"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Listing files")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())));
      if ((__default.activeParts(_7_afterDelta)).equals(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PToolCall(((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(__default.TC0(dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("List the files in the current directory"))), boxed792 -> {
  ToolCall _pat_let396_0 = ((ToolCall)(java.lang.Object)(boxed792));
  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(_pat_let396_0, boxed793 -> {
    ToolCall _8_dt__update__tmp_h2 = ((ToolCall)(java.lang.Object)(boxed793));
    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(dafny.DafnySequence.asUnicodeString("Listing files"), boxed794 -> {
      dafny.DafnySequence<? extends dafny.CodePoint> _pat_let397_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed794));
      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(_pat_let397_0, boxed795 -> {
        dafny.DafnySequence<? extends dafny.CodePoint> _9_dt__update_hinvocationMessage_h0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed795));
        return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("ls -la")), boxed796 -> {
          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let398_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed796));
          return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(_pat_let398_0, boxed797 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _10_dt__update_hpartialInput_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed797));
            return ToolCall.create((_8_dt__update__tmp_h2).dtor_toolCallId(), (_8_dt__update__tmp_h2).dtor_toolName(), (_8_dt__update__tmp_h2).dtor_displayName(), (_8_dt__update__tmp_h2).dtor_status(), (_8_dt__update__tmp_h2).dtor_intention(), (_8_dt__update__tmp_h2).dtor_contributor(), (_8_dt__update__tmp_h2).dtor_meta(), _9_dt__update_hinvocationMessage_h0, (_8_dt__update__tmp_h2).dtor_toolInput(), (_8_dt__update__tmp_h2).dtor_confirmationTitle(), (_8_dt__update__tmp_h2).dtor_riskAssessment(), (_8_dt__update__tmp_h2).dtor_edits(), (_8_dt__update__tmp_h2).dtor_editable(), (_8_dt__update__tmp_h2).dtor_options(), (_8_dt__update__tmp_h2).dtor_confirmed(), (_8_dt__update__tmp_h2).dtor_success(), (_8_dt__update__tmp_h2).dtor_pastTenseMessage(), (_8_dt__update__tmp_h2).dtor_reason(), (_8_dt__update__tmp_h2).dtor_reasonMessage(), (_8_dt__update__tmp_h2).dtor_userSuggestion(), (_8_dt__update__tmp_h2).dtor_selectedOption(), (_8_dt__update__tmp_h2).dtor_content(), (_8_dt__update__tmp_h2).dtor_structuredContent(), (_8_dt__update__tmp_h2).dtor_error(), (_8_dt__update__tmp_h2).dtor_auth(), _10_dt__update_hpartialInput_h0);
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
))))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _11_afterReady;
      _11_afterReady = __default.apply1(_7_afterDelta, ChatAction.create_CToolCallReady(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Run: ls -la")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("ls -la")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _12_tcReady;
      _12_tcReady = __default.firstActiveToolCall(_11_afterReady);
      if ((((((_12_tcReady).is_Some()) && ((((_12_tcReady).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("pending-confirmation")))) && ((((_12_tcReady).dtor_value()).dtor_invocationMessage()).equals(dafny.DafnySequence.asUnicodeString("Run: ls -la")))) && (java.util.Objects.equals(((_12_tcReady).dtor_value()).dtor_toolInput(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("ls -la"))))) && (java.util.Objects.equals(((_12_tcReady).dtor_value()).dtor_partialInput(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _13_afterConf;
      _13_afterConf = __default.apply1(_11_afterReady, ChatAction.create_CToolCallConfirmed(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("user-action")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _14_tcConf;
      _14_tcConf = __default.firstActiveToolCall(_13_afterConf);
      if ((((_14_tcConf).is_Some()) && ((((_14_tcConf).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("running")))) && (java.util.Objects.equals(((_14_tcConf).dtor_value()).dtor_confirmed(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("user-action"))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _15_afterComp;
      _15_afterComp = __default.apply1(_13_afterConf, ChatAction.create_CToolCallComplete(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Ran command")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), false, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _16_tcComp;
      _16_tcComp = __default.firstActiveToolCall(_15_afterComp);
      if (((((_16_tcComp).is_Some()) && ((((_16_tcComp).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("completed")))) && (java.util.Objects.equals(((_16_tcComp).dtor_value()).dtor_success(), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true)))) && (java.util.Objects.equals(((_16_tcComp).dtor_value()).dtor_pastTenseMessage(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Ran command"))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _17_life;
      _17_life = __default.fold(_1_act1, dafny.DafnySequence.<ChatAction> of(ChatAction._typeDescriptor(), ChatAction.create_CToolCallStart(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), dafny.DafnySequence.asUnicodeString("bash"), dafny.DafnySequence.asUnicodeString("Run Command"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("List the files in the current directory")), agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>create_None(ToolCallContributor._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallDelta(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), dafny.DafnySequence.asUnicodeString("ls -la"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Listing files")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallReady(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Run: ls -la")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("ls -la")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallConfirmed(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("user-action")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallComplete(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Ran command")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), false, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _18_lifeTc;
      _18_lifeTc = __default.firstActiveToolCall(_17_life);
      if ((((((_18_lifeTc).is_Some()) && ((((_18_lifeTc).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("completed")))) && (java.util.Objects.equals(((_18_lifeTc).dtor_value()).dtor_confirmed(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("user-action"))))) && (java.util.Objects.equals(((_18_lifeTc).dtor_value()).dtor_success(), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true)))) && (java.util.Objects.equals(((_18_lifeTc).dtor_value()).dtor_toolInput(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("ls -la"))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _19_auto;
      _19_auto = __default.fold(_1_act1, dafny.DafnySequence.<ChatAction> of(ChatAction._typeDescriptor(), ChatAction.create_CToolCallStart(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), dafny.DafnySequence.asUnicodeString("bash"), dafny.DafnySequence.asUnicodeString("Run Command"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>create_None(ToolCallContributor._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallReady(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Run")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("not-needed")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _20_autoTc;
      _20_autoTc = __default.firstActiveToolCall(_19_auto);
      if (((((_20_autoTc).is_Some()) && ((((_20_autoTc).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("running")))) && (java.util.Objects.equals(((_20_autoTc).dtor_value()).dtor_confirmed(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("not-needed"))))) && ((((_20_autoTc).dtor_value()).dtor_invocationMessage()).equals(dafny.DafnySequence.asUnicodeString("Run")))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      agency.open.ahp.ConfluxCodec.Json _21_pendingRisk;
      _21_pendingRisk = agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("status"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("loading"))) }));
      agency.open.ahp.ConfluxCodec.Json _22_pendingEdits;
      _22_pendingEdits = agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("items"), agency.open.ahp.ConfluxCodec.Json.create_JArr(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> empty(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))) }));
      agency.open.ahp.ConfluxCodec.Json _23_pendingOption;
      _23_pendingOption = agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("id"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("approve"))) }));
      agency.open.ahp.ConfluxCodec.Json _pat_let_tv0 = _23_pendingOption;
      agency.open.ahp.ConfluxCodec.Json _pat_let_tv1 = _22_pendingEdits;
      agency.open.ahp.ConfluxCodec.Json _pat_let_tv2 = _21_pendingRisk;
      ChatState _24_pendState;
      ChatState _25_dt__update__tmp_h3 = _1_act1;
      int _26_dt__update_hstatus_h0 = 24;
      agency.open.ahp.AhpSkeleton.Option<Turn> _27_dt__update_hactiveTurn_h1 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed798 -> {
  Turn _pat_let399_0 = ((Turn)(java.lang.Object)(boxed798));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let399_0, boxed799 -> {
    Turn _28_dt__update__tmp_h4 = ((Turn)(java.lang.Object)(boxed799));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PToolCall(((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(__default.TC0(dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))), boxed801 -> {
  ToolCall _pat_let401_0 = ((ToolCall)(java.lang.Object)(boxed801));
  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(_pat_let401_0, boxed802 -> {
    ToolCall _29_dt__update__tmp_h5 = ((ToolCall)(java.lang.Object)(boxed802));
    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_Some(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> of(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _pat_let_tv0)), boxed803 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _pat_let402_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>)(java.lang.Object)(boxed803));
      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>, ToolCall>Let(_pat_let402_0, boxed804 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>> _30_dt__update_hoptions_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>)(java.lang.Object)(boxed804));
        return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), boxed805 -> {
          agency.open.ahp.AhpSkeleton.Option<Boolean> _pat_let403_0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed805));
          return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, ToolCall>Let(_pat_let403_0, boxed806 -> {
            agency.open.ahp.AhpSkeleton.Option<Boolean> _31_dt__update_heditable_h0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed806));
            return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _pat_let_tv1), boxed807 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let404_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed807));
              return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let404_0, boxed808 -> {
                agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _32_dt__update_hedits_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed808));
                return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _pat_let_tv2), boxed809 -> {
                  agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let405_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed809));
                  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let405_0, boxed810 -> {
                    agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _33_dt__update_hriskAssessment_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed810));
                    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("Run in terminal"))), boxed811 -> {
                      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let406_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed811));
                      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let406_0, boxed812 -> {
                        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _34_dt__update_hconfirmationTitle_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed812));
                        return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("rm -rf /tmp/test")), boxed813 -> {
                          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let407_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed813));
                          return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(_pat_let407_0, boxed814 -> {
                            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _35_dt__update_htoolInput_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed814));
                            return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(dafny.DafnySequence.asUnicodeString("pending-confirmation"), boxed815 -> {
                              dafny.DafnySequence<? extends dafny.CodePoint> _pat_let408_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed815));
                              return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(_pat_let408_0, boxed816 -> {
                                dafny.DafnySequence<? extends dafny.CodePoint> _36_dt__update_hstatus_h1 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed816));
                                return ToolCall.create((_29_dt__update__tmp_h5).dtor_toolCallId(), (_29_dt__update__tmp_h5).dtor_toolName(), (_29_dt__update__tmp_h5).dtor_displayName(), _36_dt__update_hstatus_h1, (_29_dt__update__tmp_h5).dtor_intention(), (_29_dt__update__tmp_h5).dtor_contributor(), (_29_dt__update__tmp_h5).dtor_meta(), (_29_dt__update__tmp_h5).dtor_invocationMessage(), _35_dt__update_htoolInput_h0, _34_dt__update_hconfirmationTitle_h0, _33_dt__update_hriskAssessment_h0, _32_dt__update_hedits_h0, _31_dt__update_heditable_h0, _30_dt__update_hoptions_h0, (_29_dt__update__tmp_h5).dtor_confirmed(), (_29_dt__update__tmp_h5).dtor_success(), (_29_dt__update__tmp_h5).dtor_pastTenseMessage(), (_29_dt__update__tmp_h5).dtor_reason(), (_29_dt__update__tmp_h5).dtor_reasonMessage(), (_29_dt__update__tmp_h5).dtor_userSuggestion(), (_29_dt__update__tmp_h5).dtor_selectedOption(), (_29_dt__update__tmp_h5).dtor_content(), (_29_dt__update__tmp_h5).dtor_structuredContent(), (_29_dt__update__tmp_h5).dtor_error(), (_29_dt__update__tmp_h5).dtor_auth(), (_29_dt__update__tmp_h5).dtor_partialInput());
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
))))), boxed800 -> {
      dafny.DafnySequence<? extends Part> _pat_let400_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed800));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let400_0, boxed817 -> {
        dafny.DafnySequence<? extends Part> _37_dt__update_hparts_h1 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed817));
        return Turn.create((_28_dt__update__tmp_h4).dtor_turnId(), (_28_dt__update__tmp_h4).dtor_message(), _37_dt__update_hparts_h1, (_28_dt__update__tmp_h4).dtor_state(), (_28_dt__update__tmp_h4).dtor_usage(), (_28_dt__update__tmp_h4).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))));
      _24_pendState = ChatState.create((_25_dt__update__tmp_h3).dtor_turns(), (_25_dt__update__tmp_h3).dtor_title(), _26_dt__update_hstatus_h0, (_25_dt__update__tmp_h3).dtor_modifiedAt(), (_25_dt__update__tmp_h3).dtor_draft(), (_25_dt__update__tmp_h3).dtor_activity(), _27_dt__update_hactiveTurn_h1, (_25_dt__update__tmp_h3).dtor_steeringMessage(), (_25_dt__update__tmp_h3).dtor_queuedMessages(), (_25_dt__update__tmp_h3).dtor_nextCursor());
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _38_refreshed;
      _38_refreshed = __default.firstActiveToolCall(__default.apply1(_24_pendState, ChatAction.create_CToolCallReady(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Run again")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))));
      if ((((((((((_38_refreshed).is_Some()) && ((((_38_refreshed).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("pending-confirmation")))) && ((((_38_refreshed).dtor_value()).dtor_invocationMessage()).equals(dafny.DafnySequence.asUnicodeString("Run again")))) && (java.util.Objects.equals(((_38_refreshed).dtor_value()).dtor_toolInput(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("rm -rf /tmp/test"))))) && (java.util.Objects.equals(((_38_refreshed).dtor_value()).dtor_confirmationTitle(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("Run in terminal")))))) && (java.util.Objects.equals(((_38_refreshed).dtor_value()).dtor_riskAssessment(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _21_pendingRisk)))) && (java.util.Objects.equals(((_38_refreshed).dtor_value()).dtor_edits(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _22_pendingEdits)))) && (java.util.Objects.equals(((_38_refreshed).dtor_value()).dtor_editable(), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true)))) && (java.util.Objects.equals(((_38_refreshed).dtor_value()).dtor_options(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_Some(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> of(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _23_pendingOption))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _39_forced;
      _39_forced = __default.apply1(_4_started, ChatAction.create_CTurnComplete(dafny.DafnySequence.asUnicodeString("t1")));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _40_fTc;
      _40_fTc = __default.firstStoredToolCall(_39_forced);
      if ((((((java.util.Objects.equals((_39_forced).dtor_activeTurn(), agency.open.ahp.AhpSkeleton.Option.<Turn>create_None(Turn._typeDescriptor()))) && ((java.math.BigInteger.valueOf(((_39_forced).dtor_turns()).length())).signum() == 1)) && (((((Turn)(java.lang.Object)(((_39_forced).dtor_turns()).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))))).dtor_state()).equals(dafny.DafnySequence.asUnicodeString("complete")))) && ((_40_fTc).is_Some())) && ((((_40_fTc).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("cancelled")))) && (java.util.Objects.equals(((_40_fTc).dtor_value()).dtor_reason(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("skipped"))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _41_ignoredComplete;
      _41_ignoredComplete = __default.apply1(_4_started, ChatAction.create_CToolCallComplete(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), false, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())));
      if ((__default.activeParts(_41_ignoredComplete)).equals(__default.activeParts(_4_started))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _42_authBase;
      ChatState _43_dt__update__tmp_h6 = _1_act1;
      int _44_dt__update_hstatus_h2 = 8;
      agency.open.ahp.AhpSkeleton.Option<Turn> _45_dt__update_hactiveTurn_h2 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed818 -> {
  Turn _pat_let409_0 = ((Turn)(java.lang.Object)(boxed818));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let409_0, boxed819 -> {
    Turn _46_dt__update__tmp_h7 = ((Turn)(java.lang.Object)(boxed819));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PToolCall(((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(__default.TC0(dafny.DafnySequence.asUnicodeString("tc-auth"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))), boxed821 -> {
  ToolCall _pat_let411_0 = ((ToolCall)(java.lang.Object)(boxed821));
  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(_pat_let411_0, boxed822 -> {
    ToolCall _47_dt__update__tmp_h8 = ((ToolCall)(java.lang.Object)(boxed822));
    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("partial"))), boxed823 -> {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let412_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed823));
      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ToolCall>Let(_pat_let412_0, boxed824 -> {
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _48_dt__update_hcontent_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed824));
        return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("user-action")), boxed825 -> {
          agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let413_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed825));
          return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(_pat_let413_0, boxed826 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _49_dt__update_hconfirmed_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed826));
            return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("foo")), boxed827 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let414_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed827));
              return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(_pat_let414_0, boxed828 -> {
                agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _50_dt__update_htoolInput_h1 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed828));
                return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<ToolCallContributor>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>create_Some(ToolCallContributor._typeDescriptor(), ToolCallContributor.create(dafny.DafnySequence.asUnicodeString("mcp"), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("mcp"))) })))), boxed829 -> {
                  agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> _pat_let415_0 = ((agency.open.ahp.AhpSkeleton.Option<ToolCallContributor>)(java.lang.Object)(boxed829));
                  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<ToolCallContributor>, ToolCall>Let(_pat_let415_0, boxed830 -> {
                    agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> _51_dt__update_hcontributor_h0 = ((agency.open.ahp.AhpSkeleton.Option<ToolCallContributor>)(java.lang.Object)(boxed830));
                    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(dafny.DafnySequence.asUnicodeString("running"), boxed831 -> {
                      dafny.DafnySequence<? extends dafny.CodePoint> _pat_let416_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed831));
                      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(_pat_let416_0, boxed832 -> {
                        dafny.DafnySequence<? extends dafny.CodePoint> _52_dt__update_hstatus_h3 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed832));
                        return ToolCall.create((_47_dt__update__tmp_h8).dtor_toolCallId(), (_47_dt__update__tmp_h8).dtor_toolName(), (_47_dt__update__tmp_h8).dtor_displayName(), _52_dt__update_hstatus_h3, (_47_dt__update__tmp_h8).dtor_intention(), _51_dt__update_hcontributor_h0, (_47_dt__update__tmp_h8).dtor_meta(), (_47_dt__update__tmp_h8).dtor_invocationMessage(), _50_dt__update_htoolInput_h1, (_47_dt__update__tmp_h8).dtor_confirmationTitle(), (_47_dt__update__tmp_h8).dtor_riskAssessment(), (_47_dt__update__tmp_h8).dtor_edits(), (_47_dt__update__tmp_h8).dtor_editable(), (_47_dt__update__tmp_h8).dtor_options(), _49_dt__update_hconfirmed_h0, (_47_dt__update__tmp_h8).dtor_success(), (_47_dt__update__tmp_h8).dtor_pastTenseMessage(), (_47_dt__update__tmp_h8).dtor_reason(), (_47_dt__update__tmp_h8).dtor_reasonMessage(), (_47_dt__update__tmp_h8).dtor_userSuggestion(), (_47_dt__update__tmp_h8).dtor_selectedOption(), _48_dt__update_hcontent_h0, (_47_dt__update__tmp_h8).dtor_structuredContent(), (_47_dt__update__tmp_h8).dtor_error(), (_47_dt__update__tmp_h8).dtor_auth(), (_47_dt__update__tmp_h8).dtor_partialInput());
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
    )));
  }
  )));
}
))))), boxed820 -> {
      dafny.DafnySequence<? extends Part> _pat_let410_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed820));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let410_0, boxed833 -> {
        dafny.DafnySequence<? extends Part> _53_dt__update_hparts_h2 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed833));
        return Turn.create((_46_dt__update__tmp_h7).dtor_turnId(), (_46_dt__update__tmp_h7).dtor_message(), _53_dt__update_hparts_h2, (_46_dt__update__tmp_h7).dtor_state(), (_46_dt__update__tmp_h7).dtor_usage(), (_46_dt__update__tmp_h7).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))));
      _42_authBase = ChatState.create((_43_dt__update__tmp_h6).dtor_turns(), (_43_dt__update__tmp_h6).dtor_title(), _44_dt__update_hstatus_h2, (_43_dt__update__tmp_h6).dtor_modifiedAt(), (_43_dt__update__tmp_h6).dtor_draft(), (_43_dt__update__tmp_h6).dtor_activity(), _45_dt__update_hactiveTurn_h2, (_43_dt__update__tmp_h6).dtor_steeringMessage(), (_43_dt__update__tmp_h6).dtor_queuedMessages(), (_43_dt__update__tmp_h6).dtor_nextCursor());
      agency.open.ahp.ConfluxCodec.Json _54_challenge;
      _54_challenge = agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("reason"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("required"))) }));
      ChatState _55_authState;
      _55_authState = __default.apply1(_42_authBase, ChatAction.create_CToolCallAuthRequired(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-auth"), _54_challenge, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("auth-meta")))));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _56_authTc;
      _56_authTc = __default.firstActiveToolCall(_55_authState);
      if ((((((((((_55_authState).dtor_status()) == (24)) && ((_56_authTc).is_Some())) && ((((_56_authTc).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("auth-required")))) && (java.util.Objects.equals(((_56_authTc).dtor_value()).dtor_auth(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _54_challenge)))) && (java.util.Objects.equals(((_56_authTc).dtor_value()).dtor_meta(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("auth-meta")))))) && (java.util.Objects.equals(((_56_authTc).dtor_value()).dtor_toolInput(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("foo"))))) && (java.util.Objects.equals(((_56_authTc).dtor_value()).dtor_confirmed(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("user-action"))))) && (java.util.Objects.equals(((_56_authTc).dtor_value()).dtor_content(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("partial")))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _57_resolvedState;
      _57_resolvedState = __default.apply1(_55_authState, ChatAction.create_CToolCallAuthResolved(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-auth"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("resolved-meta")))));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _58_resolvedTc;
      _58_resolvedTc = __default.firstActiveToolCall(_57_resolvedState);
      if (((((((((((_57_resolvedState).dtor_status()) == (8)) && ((_58_resolvedTc).is_Some())) && ((((_58_resolvedTc).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("running")))) && (java.util.Objects.equals(((_58_resolvedTc).dtor_value()).dtor_auth(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) && (java.util.Objects.equals(((_58_resolvedTc).dtor_value()).dtor_contributor(), agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>create_Some(ToolCallContributor._typeDescriptor(), ToolCallContributor.create(dafny.DafnySequence.asUnicodeString("mcp"), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("mcp"))) }))))))) && (java.util.Objects.equals(((_58_resolvedTc).dtor_value()).dtor_toolInput(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("foo"))))) && (java.util.Objects.equals(((_58_resolvedTc).dtor_value()).dtor_confirmed(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("user-action"))))) && (java.util.Objects.equals(((_58_resolvedTc).dtor_value()).dtor_content(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("partial")))))) && (java.util.Objects.equals(((_58_resolvedTc).dtor_value()).dtor_meta(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("resolved-meta")))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _59_clientBase;
      ChatState _60_dt__update__tmp_h9 = _42_authBase;
      agency.open.ahp.AhpSkeleton.Option<Turn> _61_dt__update_hactiveTurn_h3 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed834 -> {
  Turn _pat_let417_0 = ((Turn)(java.lang.Object)(boxed834));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let417_0, boxed835 -> {
    Turn _62_dt__update__tmp_h10 = ((Turn)(java.lang.Object)(boxed835));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PToolCall(((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(__default.TC0(dafny.DafnySequence.asUnicodeString("tc-client"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))), boxed837 -> {
  ToolCall _pat_let419_0 = ((ToolCall)(java.lang.Object)(boxed837));
  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(_pat_let419_0, boxed838 -> {
    ToolCall _63_dt__update__tmp_h11 = ((ToolCall)(java.lang.Object)(boxed838));
    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("not-needed")), boxed839 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let420_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed839));
      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(_pat_let420_0, boxed840 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _64_dt__update_hconfirmed_h1 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed840));
        return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<ToolCallContributor>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>create_Some(ToolCallContributor._typeDescriptor(), ToolCallContributor.create(dafny.DafnySequence.asUnicodeString("client"), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("client"))) })))), boxed841 -> {
          agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> _pat_let421_0 = ((agency.open.ahp.AhpSkeleton.Option<ToolCallContributor>)(java.lang.Object)(boxed841));
          return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<ToolCallContributor>, ToolCall>Let(_pat_let421_0, boxed842 -> {
            agency.open.ahp.AhpSkeleton.Option<ToolCallContributor> _65_dt__update_hcontributor_h1 = ((agency.open.ahp.AhpSkeleton.Option<ToolCallContributor>)(java.lang.Object)(boxed842));
            return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(dafny.DafnySequence.asUnicodeString("running"), boxed843 -> {
              dafny.DafnySequence<? extends dafny.CodePoint> _pat_let422_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed843));
              return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(_pat_let422_0, boxed844 -> {
                dafny.DafnySequence<? extends dafny.CodePoint> _66_dt__update_hstatus_h4 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed844));
                return ToolCall.create((_63_dt__update__tmp_h11).dtor_toolCallId(), (_63_dt__update__tmp_h11).dtor_toolName(), (_63_dt__update__tmp_h11).dtor_displayName(), _66_dt__update_hstatus_h4, (_63_dt__update__tmp_h11).dtor_intention(), _65_dt__update_hcontributor_h1, (_63_dt__update__tmp_h11).dtor_meta(), (_63_dt__update__tmp_h11).dtor_invocationMessage(), (_63_dt__update__tmp_h11).dtor_toolInput(), (_63_dt__update__tmp_h11).dtor_confirmationTitle(), (_63_dt__update__tmp_h11).dtor_riskAssessment(), (_63_dt__update__tmp_h11).dtor_edits(), (_63_dt__update__tmp_h11).dtor_editable(), (_63_dt__update__tmp_h11).dtor_options(), _64_dt__update_hconfirmed_h1, (_63_dt__update__tmp_h11).dtor_success(), (_63_dt__update__tmp_h11).dtor_pastTenseMessage(), (_63_dt__update__tmp_h11).dtor_reason(), (_63_dt__update__tmp_h11).dtor_reasonMessage(), (_63_dt__update__tmp_h11).dtor_userSuggestion(), (_63_dt__update__tmp_h11).dtor_selectedOption(), (_63_dt__update__tmp_h11).dtor_content(), (_63_dt__update__tmp_h11).dtor_structuredContent(), (_63_dt__update__tmp_h11).dtor_error(), (_63_dt__update__tmp_h11).dtor_auth(), (_63_dt__update__tmp_h11).dtor_partialInput());
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
))))), boxed836 -> {
      dafny.DafnySequence<? extends Part> _pat_let418_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed836));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let418_0, boxed845 -> {
        dafny.DafnySequence<? extends Part> _67_dt__update_hparts_h3 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed845));
        return Turn.create((_62_dt__update__tmp_h10).dtor_turnId(), (_62_dt__update__tmp_h10).dtor_message(), _67_dt__update_hparts_h3, (_62_dt__update__tmp_h10).dtor_state(), (_62_dt__update__tmp_h10).dtor_usage(), (_62_dt__update__tmp_h10).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))));
      _59_clientBase = ChatState.create((_60_dt__update__tmp_h9).dtor_turns(), (_60_dt__update__tmp_h9).dtor_title(), (_60_dt__update__tmp_h9).dtor_status(), (_60_dt__update__tmp_h9).dtor_modifiedAt(), (_60_dt__update__tmp_h9).dtor_draft(), (_60_dt__update__tmp_h9).dtor_activity(), _61_dt__update_hactiveTurn_h3, (_60_dt__update__tmp_h9).dtor_steeringMessage(), (_60_dt__update__tmp_h9).dtor_queuedMessages(), (_60_dt__update__tmp_h9).dtor_nextCursor());
      if (java.util.Objects.equals(__default.apply1(_59_clientBase, ChatAction.create_CToolCallAuthRequired(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-client"), _54_challenge, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), _59_clientBase)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _68_authFailed;
      _68_authFailed = __default.apply1(_55_authState, ChatAction.create_CToolCallComplete(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-auth"), false, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Cancelled search")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("cancelled"))), true, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("complete-meta")))));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _69_failedTc;
      _69_failedTc = __default.firstActiveToolCall(_68_authFailed);
      if ((((((((_68_authFailed).dtor_status()) == (8)) && ((_69_failedTc).is_Some())) && ((((_69_failedTc).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("completed")))) && (java.util.Objects.equals(((_69_failedTc).dtor_value()).dtor_success(), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, false)))) && (java.util.Objects.equals(((_69_failedTc).dtor_value()).dtor_pastTenseMessage(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Cancelled search"))))) && (java.util.Objects.equals(((_69_failedTc).dtor_value()).dtor_content(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("partial")))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_55_authState, ChatAction.create_CToolCallComplete(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-auth"), true, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Should not apply")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), false, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), _55_authState)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
    }
    return pass;
  }
  public static java.math.BigInteger runPendingHistory()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    if(true) {
      pass = java.math.BigInteger.ZERO;
      ChatState _0_s;
      _0_s = __default.C0();
      ChatState _1_act1;
      ChatState _2_dt__update__tmp_h0 = _0_s;
      agency.open.ahp.AhpSkeleton.Option<Turn> _3_dt__update_hactiveTurn_h0 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), __default.T0(dafny.DafnySequence.asUnicodeString("t1")));
      _1_act1 = ChatState.create((_2_dt__update__tmp_h0).dtor_turns(), (_2_dt__update__tmp_h0).dtor_title(), (_2_dt__update__tmp_h0).dtor_status(), (_2_dt__update__tmp_h0).dtor_modifiedAt(), (_2_dt__update__tmp_h0).dtor_draft(), (_2_dt__update__tmp_h0).dtor_activity(), _3_dt__update_hactiveTurn_h0, (_2_dt__update__tmp_h0).dtor_steeringMessage(), (_2_dt__update__tmp_h0).dtor_queuedMessages(), (_2_dt__update__tmp_h0).dtor_nextCursor());
      if (java.util.Objects.equals((__default.apply1(_0_s, ChatAction.create_CPendingMessageSet(dafny.DafnySequence.asUnicodeString("steering"), dafny.DafnySequence.asUnicodeString("sm-1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("Focus on tests"))))).dtor_steeringMessage(), agency.open.ahp.AhpSkeleton.Option.<QMsg>create_Some(QMsg._typeDescriptor(), QMsg.create(dafny.DafnySequence.asUnicodeString("sm-1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("Focus on tests")))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_0_s, ChatAction.create_CPendingMessageRemoved(dafny.DafnySequence.asUnicodeString("steering"), dafny.DafnySequence.asUnicodeString("sm-unknown"))), _0_s)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals((__default.apply1(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed846 -> {
        ChatState _pat_let423_0 = ((ChatState)(java.lang.Object)(boxed846));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let423_0, boxed847 -> {
          ChatState _4_dt__update__tmp_h1 = ((ChatState)(java.lang.Object)(boxed847));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<QMsg>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<QMsg>create_Some(QMsg._typeDescriptor(), QMsg.create(dafny.DafnySequence.asUnicodeString("sm-1"), agency.open.ahp.ConfluxCodec.Json.create_JNull())), boxed848 -> {
            agency.open.ahp.AhpSkeleton.Option<QMsg> _pat_let424_0 = ((agency.open.ahp.AhpSkeleton.Option<QMsg>)(java.lang.Object)(boxed848));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<QMsg>, ChatState>Let(_pat_let424_0, boxed849 -> {
              agency.open.ahp.AhpSkeleton.Option<QMsg> _5_dt__update_hsteeringMessage_h0 = ((agency.open.ahp.AhpSkeleton.Option<QMsg>)(java.lang.Object)(boxed849));
              return ChatState.create((_4_dt__update__tmp_h1).dtor_turns(), (_4_dt__update__tmp_h1).dtor_title(), (_4_dt__update__tmp_h1).dtor_status(), (_4_dt__update__tmp_h1).dtor_modifiedAt(), (_4_dt__update__tmp_h1).dtor_draft(), (_4_dt__update__tmp_h1).dtor_activity(), (_4_dt__update__tmp_h1).dtor_activeTurn(), _5_dt__update_hsteeringMessage_h0, (_4_dt__update__tmp_h1).dtor_queuedMessages(), (_4_dt__update__tmp_h1).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), ChatAction.create_CPendingMessageRemoved(dafny.DafnySequence.asUnicodeString("steering"), dafny.DafnySequence.asUnicodeString("sm-1")))).dtor_steeringMessage(), agency.open.ahp.AhpSkeleton.Option.<QMsg>create_None(QMsg._typeDescriptor()))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      QMsg _6_q;
      _6_q = QMsg.create(dafny.DafnySequence.asUnicodeString("q-1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("m")));
      if (((__default.apply1(_0_s, ChatAction.create_CPendingMessageSet(dafny.DafnySequence.asUnicodeString("queued"), dafny.DafnySequence.asUnicodeString("q-1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("m"))))).dtor_queuedMessages()).equals(dafny.DafnySequence.<QMsg> of(QMsg._typeDescriptor(), _6_q))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      QMsg _pat_let_tv0 = _6_q;
      if (((__default.apply1(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed850 -> {
        ChatState _pat_let425_0 = ((ChatState)(java.lang.Object)(boxed850));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let425_0, boxed851 -> {
          ChatState _7_dt__update__tmp_h2 = ((ChatState)(java.lang.Object)(boxed851));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(dafny.DafnySequence.<QMsg> of(QMsg._typeDescriptor(), _pat_let_tv0, QMsg.create(dafny.DafnySequence.asUnicodeString("q-2"), agency.open.ahp.ConfluxCodec.Json.create_JNull())), boxed852 -> {
            dafny.DafnySequence<? extends QMsg> _pat_let426_0 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed852));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(_pat_let426_0, boxed853 -> {
              dafny.DafnySequence<? extends QMsg> _8_dt__update_hqueuedMessages_h0 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed853));
              return ChatState.create((_7_dt__update__tmp_h2).dtor_turns(), (_7_dt__update__tmp_h2).dtor_title(), (_7_dt__update__tmp_h2).dtor_status(), (_7_dt__update__tmp_h2).dtor_modifiedAt(), (_7_dt__update__tmp_h2).dtor_draft(), (_7_dt__update__tmp_h2).dtor_activity(), (_7_dt__update__tmp_h2).dtor_activeTurn(), (_7_dt__update__tmp_h2).dtor_steeringMessage(), _8_dt__update_hqueuedMessages_h0, (_7_dt__update__tmp_h2).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), ChatAction.create_CPendingMessageRemoved(dafny.DafnySequence.asUnicodeString("queued"), dafny.DafnySequence.asUnicodeString("q-1")))).dtor_queuedMessages()).equals(dafny.DafnySequence.<QMsg> of(QMsg._typeDescriptor(), QMsg.create(dafny.DafnySequence.asUnicodeString("q-2"), agency.open.ahp.ConfluxCodec.Json.create_JNull())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _9_three;
      ChatState _10_dt__update__tmp_h3 = _0_s;
      dafny.DafnySequence<? extends Turn> _11_dt__update_hturns_h0 = dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), __default.TDone(dafny.DafnySequence.asUnicodeString("t1")), __default.TDone(dafny.DafnySequence.asUnicodeString("t2")), __default.TDone(dafny.DafnySequence.asUnicodeString("t3")));
      _9_three = ChatState.create(_11_dt__update_hturns_h0, (_10_dt__update__tmp_h3).dtor_title(), (_10_dt__update__tmp_h3).dtor_status(), (_10_dt__update__tmp_h3).dtor_modifiedAt(), (_10_dt__update__tmp_h3).dtor_draft(), (_10_dt__update__tmp_h3).dtor_activity(), (_10_dt__update__tmp_h3).dtor_activeTurn(), (_10_dt__update__tmp_h3).dtor_steeringMessage(), (_10_dt__update__tmp_h3).dtor_queuedMessages(), (_10_dt__update__tmp_h3).dtor_nextCursor());
      if (((__default.apply1(_9_three, ChatAction.create_CTruncated(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("t2"))))).dtor_turns()).equals(dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), __default.TDone(dafny.DafnySequence.asUnicodeString("t1")), __default.TDone(dafny.DafnySequence.asUnicodeString("t2"))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_9_three, boxed854 -> {
        ChatState _pat_let427_0 = ((ChatState)(java.lang.Object)(boxed854));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let427_0, boxed855 -> {
          ChatState _12_dt__update__tmp_h4 = ((ChatState)(java.lang.Object)(boxed855));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("cur")), boxed856 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let428_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed856));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(_pat_let428_0, boxed857 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _13_dt__update_hnextCursor_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed857));
              return ChatState.create((_12_dt__update__tmp_h4).dtor_turns(), (_12_dt__update__tmp_h4).dtor_title(), (_12_dt__update__tmp_h4).dtor_status(), (_12_dt__update__tmp_h4).dtor_modifiedAt(), (_12_dt__update__tmp_h4).dtor_draft(), (_12_dt__update__tmp_h4).dtor_activity(), (_12_dt__update__tmp_h4).dtor_activeTurn(), (_12_dt__update__tmp_h4).dtor_steeringMessage(), (_12_dt__update__tmp_h4).dtor_queuedMessages(), _13_dt__update_hnextCursor_h0);
            }
            )));
          }
          )));
        }
        )));
      }
      ))), ChatAction.create_CTruncated(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)))), ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_9_three, boxed858 -> {
        ChatState _pat_let429_0 = ((ChatState)(java.lang.Object)(boxed858));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let429_0, boxed859 -> {
          ChatState _14_dt__update__tmp_h5 = ((ChatState)(java.lang.Object)(boxed859));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), boxed860 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let430_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed860));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(_pat_let430_0, boxed861 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _15_dt__update_hnextCursor_h1 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed861));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(dafny.DafnySequence.<Turn> empty(Turn._typeDescriptor()), boxed862 -> {
                dafny.DafnySequence<? extends Turn> _pat_let431_0 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed862));
                return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(_pat_let431_0, boxed863 -> {
                  dafny.DafnySequence<? extends Turn> _16_dt__update_hturns_h1 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed863));
                  return ChatState.create(_16_dt__update_hturns_h1, (_14_dt__update__tmp_h5).dtor_title(), (_14_dt__update__tmp_h5).dtor_status(), (_14_dt__update__tmp_h5).dtor_modifiedAt(), (_14_dt__update__tmp_h5).dtor_draft(), (_14_dt__update__tmp_h5).dtor_activity(), (_14_dt__update__tmp_h5).dtor_activeTurn(), (_14_dt__update__tmp_h5).dtor_steeringMessage(), (_14_dt__update__tmp_h5).dtor_queuedMessages(), _15_dt__update_hnextCursor_h1);
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
      QMsg _17_qa;
      _17_qa = QMsg.create(dafny.DafnySequence.asUnicodeString("a"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("A")));
      QMsg _18_qb;
      _18_qb = QMsg.create(dafny.DafnySequence.asUnicodeString("b"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("B")));
      QMsg _19_qc;
      _19_qc = QMsg.create(dafny.DafnySequence.asUnicodeString("c"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("C")));
      QMsg _pat_let_tv1 = _17_qa;
      QMsg _pat_let_tv2 = _18_qb;
      QMsg _pat_let_tv3 = _19_qc;
      if (((__default.apply1(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed864 -> {
        ChatState _pat_let432_0 = ((ChatState)(java.lang.Object)(boxed864));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let432_0, boxed865 -> {
          ChatState _20_dt__update__tmp_h6 = ((ChatState)(java.lang.Object)(boxed865));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(dafny.DafnySequence.<QMsg> of(QMsg._typeDescriptor(), _pat_let_tv1, _pat_let_tv2, _pat_let_tv3), boxed866 -> {
            dafny.DafnySequence<? extends QMsg> _pat_let433_0 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed866));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends QMsg>, ChatState>Let(_pat_let433_0, boxed867 -> {
              dafny.DafnySequence<? extends QMsg> _21_dt__update_hqueuedMessages_h1 = ((dafny.DafnySequence<? extends QMsg>)(java.lang.Object)(boxed867));
              return ChatState.create((_20_dt__update__tmp_h6).dtor_turns(), (_20_dt__update__tmp_h6).dtor_title(), (_20_dt__update__tmp_h6).dtor_status(), (_20_dt__update__tmp_h6).dtor_modifiedAt(), (_20_dt__update__tmp_h6).dtor_draft(), (_20_dt__update__tmp_h6).dtor_activity(), (_20_dt__update__tmp_h6).dtor_activeTurn(), (_20_dt__update__tmp_h6).dtor_steeringMessage(), _21_dt__update_hqueuedMessages_h1, (_20_dt__update__tmp_h6).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), ChatAction.create_CQueuedReordered(dafny.DafnySequence.<dafny.DafnySequence<? extends dafny.CodePoint>> of(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("c"), dafny.DafnySequence.asUnicodeString("a"), dafny.DafnySequence.asUnicodeString("b"))))).dtor_queuedMessages()).equals(dafny.DafnySequence.<QMsg> of(QMsg._typeDescriptor(), _19_qc, _17_qa, _18_qb))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed868 -> {
        ChatState _pat_let434_0 = ((ChatState)(java.lang.Object)(boxed868));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let434_0, boxed869 -> {
          ChatState _22_dt__update__tmp_h7 = ((ChatState)(java.lang.Object)(boxed869));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), __default.TDone(dafny.DafnySequence.asUnicodeString("t3"))), boxed870 -> {
            dafny.DafnySequence<? extends Turn> _pat_let435_0 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed870));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(_pat_let435_0, boxed871 -> {
              dafny.DafnySequence<? extends Turn> _23_dt__update_hturns_h2 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed871));
              return ChatState.create(_23_dt__update_hturns_h2, (_22_dt__update__tmp_h7).dtor_title(), (_22_dt__update__tmp_h7).dtor_status(), (_22_dt__update__tmp_h7).dtor_modifiedAt(), (_22_dt__update__tmp_h7).dtor_draft(), (_22_dt__update__tmp_h7).dtor_activity(), (_22_dt__update__tmp_h7).dtor_activeTurn(), (_22_dt__update__tmp_h7).dtor_steeringMessage(), (_22_dt__update__tmp_h7).dtor_queuedMessages(), (_22_dt__update__tmp_h7).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), ChatAction.create_CTurnsLoaded(dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), __default.TDone(dafny.DafnySequence.asUnicodeString("t1")), __default.TDone(dafny.DafnySequence.asUnicodeString("t2"))), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("cur-0")))), ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed872 -> {
        ChatState _pat_let436_0 = ((ChatState)(java.lang.Object)(boxed872));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let436_0, boxed873 -> {
          ChatState _24_dt__update__tmp_h8 = ((ChatState)(java.lang.Object)(boxed873));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("cur-0")), boxed874 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let437_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed874));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ChatState>Let(_pat_let437_0, boxed875 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _25_dt__update_hnextCursor_h2 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed875));
              return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), __default.TDone(dafny.DafnySequence.asUnicodeString("t1")), __default.TDone(dafny.DafnySequence.asUnicodeString("t2")), __default.TDone(dafny.DafnySequence.asUnicodeString("t3"))), boxed876 -> {
                dafny.DafnySequence<? extends Turn> _pat_let438_0 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed876));
                return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(_pat_let438_0, boxed877 -> {
                  dafny.DafnySequence<? extends Turn> _26_dt__update_hturns_h3 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed877));
                  return ChatState.create(_26_dt__update_hturns_h3, (_24_dt__update__tmp_h8).dtor_title(), (_24_dt__update__tmp_h8).dtor_status(), (_24_dt__update__tmp_h8).dtor_modifiedAt(), (_24_dt__update__tmp_h8).dtor_draft(), (_24_dt__update__tmp_h8).dtor_activity(), (_24_dt__update__tmp_h8).dtor_activeTurn(), (_24_dt__update__tmp_h8).dtor_steeringMessage(), (_24_dt__update__tmp_h8).dtor_queuedMessages(), _25_dt__update_hnextCursor_h2);
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
      if (((__default.apply1(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed878 -> {
        ChatState _pat_let439_0 = ((ChatState)(java.lang.Object)(boxed878));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let439_0, boxed879 -> {
          ChatState _27_dt__update__tmp_h9 = ((ChatState)(java.lang.Object)(boxed879));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), __default.TDone(dafny.DafnySequence.asUnicodeString("t3"))), boxed880 -> {
            dafny.DafnySequence<? extends Turn> _pat_let440_0 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed880));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Turn>, ChatState>Let(_pat_let440_0, boxed881 -> {
              dafny.DafnySequence<? extends Turn> _28_dt__update_hturns_h4 = ((dafny.DafnySequence<? extends Turn>)(java.lang.Object)(boxed881));
              return ChatState.create(_28_dt__update_hturns_h4, (_27_dt__update__tmp_h9).dtor_title(), (_27_dt__update__tmp_h9).dtor_status(), (_27_dt__update__tmp_h9).dtor_modifiedAt(), (_27_dt__update__tmp_h9).dtor_draft(), (_27_dt__update__tmp_h9).dtor_activity(), (_27_dt__update__tmp_h9).dtor_activeTurn(), (_27_dt__update__tmp_h9).dtor_steeringMessage(), (_27_dt__update__tmp_h9).dtor_queuedMessages(), (_27_dt__update__tmp_h9).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), ChatAction.create_CTurnsLoaded(dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), __default.TDone(dafny.DafnySequence.asUnicodeString("t1")), __default.TDone(dafny.DafnySequence.asUnicodeString("t2")), __default.TDone(dafny.DafnySequence.asUnicodeString("t3"))), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))))).dtor_turns()).equals(dafny.DafnySequence.<Turn> of(Turn._typeDescriptor(), __default.TDone(dafny.DafnySequence.asUnicodeString("t1")), __default.TDone(dafny.DafnySequence.asUnicodeString("t2")), __default.TDone(dafny.DafnySequence.asUnicodeString("t3"))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _29_running;
      ChatState _30_dt__update__tmp_h10 = _1_act1;
      agency.open.ahp.AhpSkeleton.Option<Turn> _31_dt__update_hactiveTurn_h1 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed882 -> {
  Turn _pat_let441_0 = ((Turn)(java.lang.Object)(boxed882));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let441_0, boxed883 -> {
    Turn _32_dt__update__tmp_h11 = ((Turn)(java.lang.Object)(boxed883));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PToolCall(((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(__default.TC0(dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))), boxed885 -> {
  ToolCall _pat_let443_0 = ((ToolCall)(java.lang.Object)(boxed885));
  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(_pat_let443_0, boxed886 -> {
    ToolCall _33_dt__update__tmp_h12 = ((ToolCall)(java.lang.Object)(boxed886));
    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(dafny.DafnySequence.asUnicodeString("running"), boxed887 -> {
      dafny.DafnySequence<? extends dafny.CodePoint> _pat_let444_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed887));
      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(_pat_let444_0, boxed888 -> {
        dafny.DafnySequence<? extends dafny.CodePoint> _34_dt__update_hstatus_h0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed888));
        return ToolCall.create((_33_dt__update__tmp_h12).dtor_toolCallId(), (_33_dt__update__tmp_h12).dtor_toolName(), (_33_dt__update__tmp_h12).dtor_displayName(), _34_dt__update_hstatus_h0, (_33_dt__update__tmp_h12).dtor_intention(), (_33_dt__update__tmp_h12).dtor_contributor(), (_33_dt__update__tmp_h12).dtor_meta(), (_33_dt__update__tmp_h12).dtor_invocationMessage(), (_33_dt__update__tmp_h12).dtor_toolInput(), (_33_dt__update__tmp_h12).dtor_confirmationTitle(), (_33_dt__update__tmp_h12).dtor_riskAssessment(), (_33_dt__update__tmp_h12).dtor_edits(), (_33_dt__update__tmp_h12).dtor_editable(), (_33_dt__update__tmp_h12).dtor_options(), (_33_dt__update__tmp_h12).dtor_confirmed(), (_33_dt__update__tmp_h12).dtor_success(), (_33_dt__update__tmp_h12).dtor_pastTenseMessage(), (_33_dt__update__tmp_h12).dtor_reason(), (_33_dt__update__tmp_h12).dtor_reasonMessage(), (_33_dt__update__tmp_h12).dtor_userSuggestion(), (_33_dt__update__tmp_h12).dtor_selectedOption(), (_33_dt__update__tmp_h12).dtor_content(), (_33_dt__update__tmp_h12).dtor_structuredContent(), (_33_dt__update__tmp_h12).dtor_error(), (_33_dt__update__tmp_h12).dtor_auth(), (_33_dt__update__tmp_h12).dtor_partialInput());
      }
      )));
    }
    )));
  }
  )));
}
))))), boxed884 -> {
      dafny.DafnySequence<? extends Part> _pat_let442_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed884));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let442_0, boxed889 -> {
        dafny.DafnySequence<? extends Part> _35_dt__update_hparts_h0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed889));
        return Turn.create((_32_dt__update__tmp_h11).dtor_turnId(), (_32_dt__update__tmp_h11).dtor_message(), _35_dt__update_hparts_h0, (_32_dt__update__tmp_h11).dtor_state(), (_32_dt__update__tmp_h11).dtor_usage(), (_32_dt__update__tmp_h11).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))));
      _29_running = ChatState.create((_30_dt__update__tmp_h10).dtor_turns(), (_30_dt__update__tmp_h10).dtor_title(), (_30_dt__update__tmp_h10).dtor_status(), (_30_dt__update__tmp_h10).dtor_modifiedAt(), (_30_dt__update__tmp_h10).dtor_draft(), (_30_dt__update__tmp_h10).dtor_activity(), _31_dt__update_hactiveTurn_h1, (_30_dt__update__tmp_h10).dtor_steeringMessage(), (_30_dt__update__tmp_h10).dtor_queuedMessages(), (_30_dt__update__tmp_h10).dtor_nextCursor());
      ChatState _36_afterCc;
      _36_afterCc = __default.apply1(_29_running, ChatAction.create_CToolCallContentChanged(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("terminal-ref")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())));
      Part _37_ccTc;
      _37_ccTc = ((Part)(java.lang.Object)(((((_36_afterCc).dtor_activeTurn()).dtor_value()).dtor_parts()).select(dafny.Helpers.toInt((java.math.BigInteger.ZERO)))));
      if ((((_37_ccTc).is_PToolCall()) && ((((_37_ccTc).dtor_tc()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("running")))) && (java.util.Objects.equals(((_37_ccTc).dtor_tc()).dtor_content(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("terminal-ref")))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _38_withR;
      ChatState _39_dt__update__tmp_h13 = _1_act1;
      agency.open.ahp.AhpSkeleton.Option<Turn> _40_dt__update_hactiveTurn_h2 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed890 -> {
  Turn _pat_let445_0 = ((Turn)(java.lang.Object)(boxed890));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let445_0, boxed891 -> {
    Turn _41_dt__update__tmp_h14 = ((Turn)(java.lang.Object)(boxed891));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PReasoning(dafny.DafnySequence.asUnicodeString("r-1"), dafny.DafnySequence.asUnicodeString(""))), boxed892 -> {
      dafny.DafnySequence<? extends Part> _pat_let446_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed892));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let446_0, boxed893 -> {
        dafny.DafnySequence<? extends Part> _42_dt__update_hparts_h1 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed893));
        return Turn.create((_41_dt__update__tmp_h14).dtor_turnId(), (_41_dt__update__tmp_h14).dtor_message(), _42_dt__update_hparts_h1, (_41_dt__update__tmp_h14).dtor_state(), (_41_dt__update__tmp_h14).dtor_usage(), (_41_dt__update__tmp_h14).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))));
      _38_withR = ChatState.create((_39_dt__update__tmp_h13).dtor_turns(), (_39_dt__update__tmp_h13).dtor_title(), (_39_dt__update__tmp_h13).dtor_status(), (_39_dt__update__tmp_h13).dtor_modifiedAt(), (_39_dt__update__tmp_h13).dtor_draft(), (_39_dt__update__tmp_h13).dtor_activity(), _40_dt__update_hactiveTurn_h2, (_39_dt__update__tmp_h13).dtor_steeringMessage(), (_39_dt__update__tmp_h13).dtor_queuedMessages(), (_39_dt__update__tmp_h13).dtor_nextCursor());
      ChatState _43_r2;
      _43_r2 = __default.fold(_38_withR, dafny.DafnySequence.<ChatAction> of(ChatAction._typeDescriptor(), ChatAction.create_CReasoning(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("r-1"), dafny.DafnySequence.asUnicodeString("Thinking about ")), ChatAction.create_CReasoning(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("r-1"), dafny.DafnySequence.asUnicodeString("the plan"))));
      if (java.util.Objects.equals((_43_r2).dtor_activeTurn(), agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed894 -> {
  Turn _pat_let447_0 = ((Turn)(java.lang.Object)(boxed894));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let447_0, boxed895 -> {
    Turn _44_dt__update__tmp_h15 = ((Turn)(java.lang.Object)(boxed895));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PReasoning(dafny.DafnySequence.asUnicodeString("r-1"), dafny.DafnySequence.asUnicodeString("Thinking about the plan"))), boxed896 -> {
      dafny.DafnySequence<? extends Part> _pat_let448_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed896));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let448_0, boxed897 -> {
        dafny.DafnySequence<? extends Part> _45_dt__update_hparts_h2 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed897));
        return Turn.create((_44_dt__update__tmp_h15).dtor_turnId(), (_44_dt__update__tmp_h15).dtor_message(), _45_dt__update_hparts_h2, (_44_dt__update__tmp_h15).dtor_state(), (_44_dt__update__tmp_h15).dtor_usage(), (_44_dt__update__tmp_h15).dtor_error());
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
    }
    return pass;
  }
  public static java.math.BigInteger runInputFlow()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    if(true) {
      pass = java.math.BigInteger.ZERO;
      ChatState _0_s;
      _0_s = __default.C0();
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _1_openBody;
      _1_openBody = dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("message"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("Clarify requirements"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("questions"), agency.open.ahp.ConfluxCodec.Json.create_JArr(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> of(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("id"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("q1"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("required"), agency.open.ahp.ConfluxCodec.Json.create_JBool(true)), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("options"), agency.open.ahp.ConfluxCodec.Json.create_JArr(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> of(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("node"))))) }))))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("url"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("https://example.test/review"))) });
      ChatState _2_at1;
      ChatState _3_dt__update__tmp_h0 = _0_s;
      int _4_dt__update_hstatus_h0 = 8;
      agency.open.ahp.AhpSkeleton.Option<Turn> _5_dt__update_hactiveTurn_h0 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), __default.T0(dafny.DafnySequence.asUnicodeString("turn-1")));
      _2_at1 = ChatState.create((_3_dt__update__tmp_h0).dtor_turns(), (_3_dt__update__tmp_h0).dtor_title(), _4_dt__update_hstatus_h0, (_3_dt__update__tmp_h0).dtor_modifiedAt(), (_3_dt__update__tmp_h0).dtor_draft(), (_3_dt__update__tmp_h0).dtor_activity(), _5_dt__update_hactiveTurn_h0, (_3_dt__update__tmp_h0).dtor_steeringMessage(), (_3_dt__update__tmp_h0).dtor_queuedMessages(), (_3_dt__update__tmp_h0).dtor_nextCursor());
      ChatState _6_afterReq;
      _6_afterReq = __default.apply1(_2_at1, ChatAction.create_CInputRequested(InputReq.create(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ }), _1_openBody)));
      if (((((_6_afterReq).dtor_status()) == (24)) && (((_6_afterReq).dtor_activeTurn()).is_Some())) && (((((_6_afterReq).dtor_activeTurn()).dtor_value()).dtor_parts()).equals(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PInputRequest(InputReq.create(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ }), _1_openBody), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed898 -> {
        ChatState _pat_let449_0 = ((ChatState)(java.lang.Object)(boxed898));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let449_0, boxed899 -> {
          ChatState _7_dt__update__tmp_h1 = ((ChatState)(java.lang.Object)(boxed899));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(8, boxed900 -> {
            int _pat_let450_0 = ((int)(java.lang.Object)(boxed900));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let450_0, boxed901 -> {
              int _8_dt__update_hstatus_h1 = ((int)(java.lang.Object)(boxed901));
              return ChatState.create((_7_dt__update__tmp_h1).dtor_turns(), (_7_dt__update__tmp_h1).dtor_title(), _8_dt__update_hstatus_h1, (_7_dt__update__tmp_h1).dtor_modifiedAt(), (_7_dt__update__tmp_h1).dtor_draft(), (_7_dt__update__tmp_h1).dtor_activity(), (_7_dt__update__tmp_h1).dtor_activeTurn(), (_7_dt__update__tmp_h1).dtor_steeringMessage(), (_7_dt__update__tmp_h1).dtor_queuedMessages(), (_7_dt__update__tmp_h1).dtor_nextCursor());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), ChatAction.create_CInputRequested(InputReq.create(dafny.DafnySequence.asUnicodeString("orphan-1"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ }), _1_openBody))), ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_0_s, boxed902 -> {
        ChatState _pat_let451_0 = ((ChatState)(java.lang.Object)(boxed902));
        return ((ChatState)(java.lang.Object)(dafny.Helpers.<ChatState, ChatState>Let(_pat_let451_0, boxed903 -> {
          ChatState _9_dt__update__tmp_h2 = ((ChatState)(java.lang.Object)(boxed903));
          return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(8, boxed904 -> {
            int _pat_let452_0 = ((int)(java.lang.Object)(boxed904));
            return ((ChatState)(java.lang.Object)(dafny.Helpers.<java.lang.Integer, ChatState>Let(_pat_let452_0, boxed905 -> {
              int _10_dt__update_hstatus_h2 = ((int)(java.lang.Object)(boxed905));
              return ChatState.create((_9_dt__update__tmp_h2).dtor_turns(), (_9_dt__update__tmp_h2).dtor_title(), _10_dt__update_hstatus_h2, (_9_dt__update__tmp_h2).dtor_modifiedAt(), (_9_dt__update__tmp_h2).dtor_draft(), (_9_dt__update__tmp_h2).dtor_activity(), (_9_dt__update__tmp_h2).dtor_activeTurn(), (_9_dt__update__tmp_h2).dtor_steeringMessage(), (_9_dt__update__tmp_h2).dtor_queuedMessages(), (_9_dt__update__tmp_h2).dtor_nextCursor());
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
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _pat_let_tv0 = _1_openBody;
      ChatState _11_reqState;
      ChatState _12_dt__update__tmp_h3 = _2_at1;
      int _13_dt__update_hstatus_h3 = 24;
      agency.open.ahp.AhpSkeleton.Option<Turn> _14_dt__update_hactiveTurn_h1 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("turn-1")), boxed906 -> {
  Turn _pat_let453_0 = ((Turn)(java.lang.Object)(boxed906));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let453_0, boxed907 -> {
    Turn _15_dt__update__tmp_h4 = ((Turn)(java.lang.Object)(boxed907));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PInputRequest(InputReq.create(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ }), _pat_let_tv0), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)))), boxed908 -> {
      dafny.DafnySequence<? extends Part> _pat_let454_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed908));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let454_0, boxed909 -> {
        dafny.DafnySequence<? extends Part> _16_dt__update_hparts_h0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed909));
        return Turn.create((_15_dt__update__tmp_h4).dtor_turnId(), (_15_dt__update__tmp_h4).dtor_message(), _16_dt__update_hparts_h0, (_15_dt__update__tmp_h4).dtor_state(), (_15_dt__update__tmp_h4).dtor_usage(), (_15_dt__update__tmp_h4).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))));
      _11_reqState = ChatState.create((_12_dt__update__tmp_h3).dtor_turns(), (_12_dt__update__tmp_h3).dtor_title(), _13_dt__update_hstatus_h3, (_12_dt__update__tmp_h3).dtor_modifiedAt(), (_12_dt__update__tmp_h3).dtor_draft(), (_12_dt__update__tmp_h3).dtor_activity(), _14_dt__update_hactiveTurn_h1, (_12_dt__update__tmp_h3).dtor_steeringMessage(), (_12_dt__update__tmp_h3).dtor_queuedMessages(), (_12_dt__update__tmp_h3).dtor_nextCursor());
      if (((((__default.apply1(_11_reqState, ChatAction.create_CInputAnswerChanged(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnySequence.asUnicodeString("q1"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("node")))))).dtor_activeTurn()).dtor_value()).dtor_parts()).equals(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PInputRequest(InputReq.create(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("q1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("node"))) }), _1_openBody), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_11_reqState, ChatAction.create_CInputAnswerChanged(dafny.DafnySequence.asUnicodeString("missing"), dafny.DafnySequence.asUnicodeString("q1"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), _11_reqState)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _17_afterDone;
      _17_afterDone = __default.apply1(_11_reqState, ChatAction.create_CInputCompleted(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnySequence.asUnicodeString("accept"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ })));
      if (((((_17_afterDone).dtor_status()) == (8)) && (((_17_afterDone).dtor_activeTurn()).is_Some())) && (((((_17_afterDone).dtor_activeTurn()).dtor_value()).dtor_parts()).equals(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PInputRequest(InputReq.create(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ }), _1_openBody), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("accept"))))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_11_reqState, ChatAction.create_CInputCompleted(dafny.DafnySequence.asUnicodeString("missing"), dafny.DafnySequence.asUnicodeString("cancel"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ }))), _11_reqState)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _18_flow;
      _18_flow = __default.fold(_0_s, dafny.DafnySequence.<ChatAction> of(ChatAction._typeDescriptor(), ChatAction.create_CTurnStarted(dafny.DafnySequence.asUnicodeString("turn-1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("Hello")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))), ChatAction.create_CInputRequested(InputReq.create(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ }), _1_openBody)), ChatAction.create_CInputAnswerChanged(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnySequence.asUnicodeString("q1"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("node")))), ChatAction.create_CInputAnswerChanged(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnySequence.asUnicodeString("q2"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("small")))), ChatAction.create_CInputCompleted(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnySequence.asUnicodeString("accept"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ }))));
      if (((((_18_flow).dtor_status()) == (8)) && (((_18_flow).dtor_activeTurn()).is_Some())) && (((((_18_flow).dtor_activeTurn()).dtor_value()).dtor_parts()).equals(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PInputRequest(InputReq.create(dafny.DafnySequence.asUnicodeString("input-1"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("q1"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("node"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("q2"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("small"))) }), _1_openBody), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("accept"))))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      dafny.DafnyMap<? extends dafny.DafnySequence<? extends dafny.CodePoint>, ? extends agency.open.ahp.ConfluxCodec.Json> _pat_let_tv1 = _1_openBody;
      ChatState _19_overrideState;
      ChatState _20_dt__update__tmp_h5 = _2_at1;
      int _21_dt__update_hstatus_h4 = 24;
      agency.open.ahp.AhpSkeleton.Option<Turn> _22_dt__update_hactiveTurn_h2 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("turn-1")), boxed910 -> {
  Turn _pat_let455_0 = ((Turn)(java.lang.Object)(boxed910));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let455_0, boxed911 -> {
    Turn _23_dt__update__tmp_h6 = ((Turn)(java.lang.Object)(boxed911));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PInputRequest(InputReq.create(dafny.DafnySequence.asUnicodeString("review-1"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("approve"), agency.open.ahp.ConfluxCodec.Json.create_JBool(true)) }), _pat_let_tv1), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)))), boxed912 -> {
      dafny.DafnySequence<? extends Part> _pat_let456_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed912));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let456_0, boxed913 -> {
        dafny.DafnySequence<? extends Part> _24_dt__update_hparts_h1 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed913));
        return Turn.create((_23_dt__update__tmp_h6).dtor_turnId(), (_23_dt__update__tmp_h6).dtor_message(), _24_dt__update_hparts_h1, (_23_dt__update__tmp_h6).dtor_state(), (_23_dt__update__tmp_h6).dtor_usage(), (_23_dt__update__tmp_h6).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))));
      _19_overrideState = ChatState.create((_20_dt__update__tmp_h5).dtor_turns(), (_20_dt__update__tmp_h5).dtor_title(), _21_dt__update_hstatus_h4, (_20_dt__update__tmp_h5).dtor_modifiedAt(), (_20_dt__update__tmp_h5).dtor_draft(), (_20_dt__update__tmp_h5).dtor_activity(), _22_dt__update_hactiveTurn_h2, (_20_dt__update__tmp_h5).dtor_steeringMessage(), (_20_dt__update__tmp_h5).dtor_queuedMessages(), (_20_dt__update__tmp_h5).dtor_nextCursor());
      ChatState _25_overridden;
      _25_overridden = __default.apply1(_19_overrideState, ChatAction.create_CInputCompleted(dafny.DafnySequence.asUnicodeString("review-1"), dafny.DafnySequence.asUnicodeString("decline"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("approve"), agency.open.ahp.ConfluxCodec.Json.create_JBool(false)) })));
      if (((((_25_overridden).dtor_status()) == (8)) && (((_25_overridden).dtor_activeTurn()).is_Some())) && (((((_25_overridden).dtor_activeTurn()).dtor_value()).dtor_parts()).equals(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PInputRequest(InputReq.create(dafny.DafnySequence.asUnicodeString("review-1"), dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("approve"), agency.open.ahp.ConfluxCodec.Json.create_JBool(false)) }), _1_openBody), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("decline"))))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
    }
    return pass;
  }
  public static java.math.BigInteger runResultConfirm()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    if(true) {
      pass = java.math.BigInteger.ZERO;
      ChatState _0_s;
      _0_s = __default.C0();
      ChatState _1_act1;
      ChatState _2_dt__update__tmp_h0 = _0_s;
      agency.open.ahp.AhpSkeleton.Option<Turn> _3_dt__update_hactiveTurn_h0 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), __default.T0(dafny.DafnySequence.asUnicodeString("t1")));
      _1_act1 = ChatState.create((_2_dt__update__tmp_h0).dtor_turns(), (_2_dt__update__tmp_h0).dtor_title(), (_2_dt__update__tmp_h0).dtor_status(), (_2_dt__update__tmp_h0).dtor_modifiedAt(), (_2_dt__update__tmp_h0).dtor_draft(), (_2_dt__update__tmp_h0).dtor_activity(), _3_dt__update_hactiveTurn_h0, (_2_dt__update__tmp_h0).dtor_steeringMessage(), (_2_dt__update__tmp_h0).dtor_queuedMessages(), (_2_dt__update__tmp_h0).dtor_nextCursor());
      ChatState _4_rcLife;
      _4_rcLife = __default.fold(_1_act1, dafny.DafnySequence.<ChatAction> of(ChatAction._typeDescriptor(), ChatAction.create_CToolCallStart(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), dafny.DafnySequence.asUnicodeString("bash"), dafny.DafnySequence.asUnicodeString("Run Command"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>create_None(ToolCallContributor._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallReady(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Run")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("not-needed")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallComplete(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Done")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), false, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallResultConfirmed(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _5_rcTc;
      _5_rcTc = __default.firstActiveToolCall(_4_rcLife);
      if ((((_5_rcTc).is_Some()) && ((((_5_rcTc).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("completed")))) && (java.util.Objects.equals(((_5_rcTc).dtor_value()).dtor_success(), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true)))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _6_rdLife;
      _6_rdLife = __default.fold(_1_act1, dafny.DafnySequence.<ChatAction> of(ChatAction._typeDescriptor(), ChatAction.create_CToolCallStart(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), dafny.DafnySequence.asUnicodeString("bash"), dafny.DafnySequence.asUnicodeString("Run Command"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>create_None(ToolCallContributor._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallReady(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Run")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("not-needed")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallComplete(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Done")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), true, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallResultConfirmed(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), false, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _7_rdTc;
      _7_rdTc = __default.firstActiveToolCall(_6_rdLife);
      if ((((((_7_rdTc).is_Some()) && ((((_7_rdTc).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("cancelled")))) && (java.util.Objects.equals(((_7_rdTc).dtor_value()).dtor_confirmed(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))))) && (java.util.Objects.equals(((_7_rdTc).dtor_value()).dtor_success(), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN)))) && (java.util.Objects.equals(((_7_rdTc).dtor_value()).dtor_reason(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("result-denied"))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _8_runState;
      ChatState _9_dt__update__tmp_h1 = _1_act1;
      int _10_dt__update_hstatus_h0 = 8;
      agency.open.ahp.AhpSkeleton.Option<Turn> _11_dt__update_hactiveTurn_h1 = agency.open.ahp.AhpSkeleton.Option.<Turn>create_Some(Turn._typeDescriptor(), ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(__default.T0(dafny.DafnySequence.asUnicodeString("t1")), boxed914 -> {
  Turn _pat_let457_0 = ((Turn)(java.lang.Object)(boxed914));
  return ((Turn)(java.lang.Object)(dafny.Helpers.<Turn, Turn>Let(_pat_let457_0, boxed915 -> {
    Turn _12_dt__update__tmp_h2 = ((Turn)(java.lang.Object)(boxed915));
    return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_PToolCall(((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(__default.TC0(dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR))), boxed917 -> {
  ToolCall _pat_let459_0 = ((ToolCall)(java.lang.Object)(boxed917));
  return ((ToolCall)(java.lang.Object)(dafny.Helpers.<ToolCall, ToolCall>Let(_pat_let459_0, boxed918 -> {
    ToolCall _13_dt__update__tmp_h3 = ((ToolCall)(java.lang.Object)(boxed918));
    return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("user-action")), boxed919 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let460_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed919));
      return ((ToolCall)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, ToolCall>Let(_pat_let460_0, boxed920 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _14_dt__update_hconfirmed_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed920));
        return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(dafny.DafnySequence.asUnicodeString("running"), boxed921 -> {
          dafny.DafnySequence<? extends dafny.CodePoint> _pat_let461_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed921));
          return ((ToolCall)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ToolCall>Let(_pat_let461_0, boxed922 -> {
            dafny.DafnySequence<? extends dafny.CodePoint> _15_dt__update_hstatus_h1 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed922));
            return ToolCall.create((_13_dt__update__tmp_h3).dtor_toolCallId(), (_13_dt__update__tmp_h3).dtor_toolName(), (_13_dt__update__tmp_h3).dtor_displayName(), _15_dt__update_hstatus_h1, (_13_dt__update__tmp_h3).dtor_intention(), (_13_dt__update__tmp_h3).dtor_contributor(), (_13_dt__update__tmp_h3).dtor_meta(), (_13_dt__update__tmp_h3).dtor_invocationMessage(), (_13_dt__update__tmp_h3).dtor_toolInput(), (_13_dt__update__tmp_h3).dtor_confirmationTitle(), (_13_dt__update__tmp_h3).dtor_riskAssessment(), (_13_dt__update__tmp_h3).dtor_edits(), (_13_dt__update__tmp_h3).dtor_editable(), (_13_dt__update__tmp_h3).dtor_options(), _14_dt__update_hconfirmed_h0, (_13_dt__update__tmp_h3).dtor_success(), (_13_dt__update__tmp_h3).dtor_pastTenseMessage(), (_13_dt__update__tmp_h3).dtor_reason(), (_13_dt__update__tmp_h3).dtor_reasonMessage(), (_13_dt__update__tmp_h3).dtor_userSuggestion(), (_13_dt__update__tmp_h3).dtor_selectedOption(), (_13_dt__update__tmp_h3).dtor_content(), (_13_dt__update__tmp_h3).dtor_structuredContent(), (_13_dt__update__tmp_h3).dtor_error(), (_13_dt__update__tmp_h3).dtor_auth(), (_13_dt__update__tmp_h3).dtor_partialInput());
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
))))), boxed916 -> {
      dafny.DafnySequence<? extends Part> _pat_let458_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed916));
      return ((Turn)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, Turn>Let(_pat_let458_0, boxed923 -> {
        dafny.DafnySequence<? extends Part> _16_dt__update_hparts_h0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed923));
        return Turn.create((_12_dt__update__tmp_h2).dtor_turnId(), (_12_dt__update__tmp_h2).dtor_message(), _16_dt__update_hparts_h0, (_12_dt__update__tmp_h2).dtor_state(), (_12_dt__update__tmp_h2).dtor_usage(), (_12_dt__update__tmp_h2).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))));
      _8_runState = ChatState.create((_9_dt__update__tmp_h1).dtor_turns(), (_9_dt__update__tmp_h1).dtor_title(), _10_dt__update_hstatus_h0, (_9_dt__update__tmp_h1).dtor_modifiedAt(), (_9_dt__update__tmp_h1).dtor_draft(), (_9_dt__update__tmp_h1).dtor_activity(), _11_dt__update_hactiveTurn_h1, (_9_dt__update__tmp_h1).dtor_steeringMessage(), (_9_dt__update__tmp_h1).dtor_queuedMessages(), (_9_dt__update__tmp_h1).dtor_nextCursor());
      if (java.util.Objects.equals(__default.apply1(_8_runState, ChatAction.create_CToolCallResultConfirmed(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), _8_runState)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChatState _17_selLife;
      _17_selLife = __default.fold(_1_act1, dafny.DafnySequence.<ChatAction> of(ChatAction._typeDescriptor(), ChatAction.create_CToolCallStart(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), dafny.DafnySequence.asUnicodeString("bash"), dafny.DafnySequence.asUnicodeString("Run Command"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<ToolCallContributor>create_None(ToolCallContributor._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallReady(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Run")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends agency.open.ahp.ConfluxCodec.Json>>create_None(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json>_typeDescriptor(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallConfirmed(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("user-action")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallComplete(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("Done")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), false, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChatAction.create_CToolCallResultConfirmed(dafny.DafnySequence.asUnicodeString("t1"), dafny.DafnySequence.asUnicodeString("tc-1"), true, agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))));
      agency.open.ahp.AhpSkeleton.Option<ToolCall> _18_selTc;
      _18_selTc = __default.firstActiveToolCall(_17_selLife);
      if (((((_18_selTc).is_Some()) && ((((_18_selTc).dtor_value()).dtor_status()).equals(dafny.DafnySequence.asUnicodeString("completed")))) && (java.util.Objects.equals(((_18_selTc).dtor_value()).dtor_confirmed(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("user-action"))))) && (java.util.Objects.equals(((_18_selTc).dtor_value()).dtor_success(), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true)))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
    }
    return pass;
  }
  public static dafny.Tuple2 RunCorpus()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    java.math.BigInteger modeled = java.math.BigInteger.ZERO;
    if(true) {
      modeled = java.math.BigInteger.valueOf(54L);
      java.math.BigInteger _0_p1 = java.math.BigInteger.ZERO;
      java.math.BigInteger _out0 = java.math.BigInteger.ZERO;
      _out0 = __default.runScalarTurns();
      _0_p1 = _out0;
      java.math.BigInteger _1_p2 = java.math.BigInteger.ZERO;
      java.math.BigInteger _out1 = java.math.BigInteger.ZERO;
      _out1 = __default.runToolCalls();
      _1_p2 = _out1;
      java.math.BigInteger _2_p3 = java.math.BigInteger.ZERO;
      java.math.BigInteger _out2 = java.math.BigInteger.ZERO;
      _out2 = __default.runPendingHistory();
      _2_p3 = _out2;
      java.math.BigInteger _3_p4 = java.math.BigInteger.ZERO;
      java.math.BigInteger _out3 = java.math.BigInteger.ZERO;
      _out3 = __default.runInputFlow();
      _3_p4 = _out3;
      java.math.BigInteger _4_p5 = java.math.BigInteger.ZERO;
      java.math.BigInteger _out4 = java.math.BigInteger.ZERO;
      _out4 = __default.runResultConfirm();
      _4_p5 = _out4;
      pass = ((((_0_p1).add(_1_p2)).add(_2_p3)).add(_3_p4)).add(_4_p5);
    }
    return new dafny.Tuple2<>(pass, modeled);
  }
  public static java.util.function.Function<QMsg, dafny.DafnySequence<? extends dafny.CodePoint>> qKey()
  {
    return ((java.util.function.Function<QMsg, dafny.DafnySequence<? extends dafny.CodePoint>>)(_0_x_boxed0) -> {
      QMsg _0_x = ((QMsg)(java.lang.Object)(_0_x_boxed0));
      return (_0_x).dtor_id();
    });
  }
  public static int ERR()
  {
    return 2;
  }
  public static int INP()
  {
    return 24;
  }
  public static int GEN()
  {
    return 8;
  }
  public static int IDLE()
  {
    return 1;
  }
  public static int ACT()
  {
    return 31;
  }
  public static int READ()
  {
    return 32;
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.Chat._default";
  }
}
