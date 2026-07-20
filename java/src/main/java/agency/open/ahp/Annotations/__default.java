// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.Annotations;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static dafny.DafnySequence<? extends dafny.CodePoint> AnnId(Annotation a) {
    return (a).dtor_id();
  }
  public static dafny.DafnySequence<? extends dafny.CodePoint> EntryId(Entry e) {
    return (e).dtor_id();
  }
  public static dafny.DafnySequence<? extends Annotation> upsertAnn(dafny.DafnySequence<? extends Annotation> anns, Annotation a)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Annotation>SeqUpsertById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Annotation._typeDescriptor(), __default::AnnId, anns, a);
  }
  public static dafny.DafnySequence<? extends Annotation> removeAnn(dafny.DafnySequence<? extends Annotation> anns, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Annotation>SeqRemoveById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Annotation._typeDescriptor(), __default::AnnId, anns, id);
  }
  public static boolean hasAnn(dafny.DafnySequence<? extends Annotation> anns, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return ((dafny.Function2<dafny.DafnySequence<? extends Annotation>, dafny.DafnySequence<? extends dafny.CodePoint>, Boolean>)(_0_anns, _1_id) -> {return dafny.Helpers.Quantifier(dafny.Helpers.IntegerRange(java.math.BigInteger.ZERO, java.math.BigInteger.valueOf((_0_anns).length())), false, ((_exists_var_0_boxed0) -> {
      java.math.BigInteger _exists_var_0 = ((java.math.BigInteger)(java.lang.Object)(_exists_var_0_boxed0));
      java.math.BigInteger _2_i = (java.math.BigInteger)_exists_var_0;
      return (((_2_i).signum() != -1) && ((_2_i).compareTo(java.math.BigInteger.valueOf((_0_anns).length())) < 0)) && (((((Annotation)(java.lang.Object)((_0_anns).select(dafny.Helpers.toInt((_2_i)))))).dtor_id()).equals(_1_id));
    }));}).apply(anns, id);
  }
  public static dafny.DafnySequence<? extends Entry> upsertEntry(dafny.DafnySequence<? extends Entry> entries, Entry e)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Entry>SeqUpsertById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Entry._typeDescriptor(), __default::EntryId, entries, e);
  }
  public static dafny.DafnySequence<? extends Entry> removeEntry(dafny.DafnySequence<? extends Entry> entries, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Entry>SeqRemoveById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Entry._typeDescriptor(), __default::EntryId, entries, id);
  }
  public static dafny.DafnySequence<? extends Annotation> updateAnn(dafny.DafnySequence<? extends Annotation> anns, dafny.DafnySequence<? extends dafny.CodePoint> id, java.util.function.Function<Annotation, Annotation> f)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, Annotation>SeqUpdateById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), Annotation._typeDescriptor(), __default::AnnId, anns, id, f);
  }
  public static java.util.function.Function<Annotation, Annotation> setEntry(Entry e) {
    return ((java.util.function.Function<Entry, java.util.function.Function<Annotation, Annotation>>)(_0_e) -> {return ((java.util.function.Function<Annotation, Annotation>)(_1_an_boxed0) -> {
      Annotation _1_an = ((Annotation)(java.lang.Object)(_1_an_boxed0));
      return ((Annotation)(java.lang.Object)(dafny.Helpers.<Annotation, Annotation>Let(_1_an, boxed100 -> {
        Annotation _pat_let50_0 = ((Annotation)(java.lang.Object)(boxed100));
        return ((Annotation)(java.lang.Object)(dafny.Helpers.<Annotation, Annotation>Let(_pat_let50_0, boxed101 -> {
          Annotation _2_dt__update__tmp_h0 = ((Annotation)(java.lang.Object)(boxed101));
          return ((Annotation)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Entry>, Annotation>Let(__default.upsertEntry((_1_an).dtor_entries(), _0_e), boxed102 -> {
            dafny.DafnySequence<? extends Entry> _pat_let51_0 = ((dafny.DafnySequence<? extends Entry>)(java.lang.Object)(boxed102));
            return ((Annotation)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Entry>, Annotation>Let(_pat_let51_0, boxed103 -> {
              dafny.DafnySequence<? extends Entry> _3_dt__update_hentries_h0 = ((dafny.DafnySequence<? extends Entry>)(java.lang.Object)(boxed103));
              return Annotation.create((_2_dt__update__tmp_h0).dtor_id(), (_2_dt__update__tmp_h0).dtor_turnId(), (_2_dt__update__tmp_h0).dtor_resource(), (_2_dt__update__tmp_h0).dtor_range(), (_2_dt__update__tmp_h0).dtor_resolved(), _3_dt__update_hentries_h0, (_2_dt__update__tmp_h0).dtor_meta());
            }
            )));
          }
          )));
        }
        )));
      }
      )));
    });}).apply(e);
  }
  public static java.util.function.Function<Annotation, Annotation> dropEntry(dafny.DafnySequence<? extends dafny.CodePoint> eid) {
    return ((java.util.function.Function<dafny.DafnySequence<? extends dafny.CodePoint>, java.util.function.Function<Annotation, Annotation>>)(_0_eid) -> {return ((java.util.function.Function<Annotation, Annotation>)(_1_an_boxed0) -> {
      Annotation _1_an = ((Annotation)(java.lang.Object)(_1_an_boxed0));
      return ((Annotation)(java.lang.Object)(dafny.Helpers.<Annotation, Annotation>Let(_1_an, boxed104 -> {
        Annotation _pat_let52_0 = ((Annotation)(java.lang.Object)(boxed104));
        return ((Annotation)(java.lang.Object)(dafny.Helpers.<Annotation, Annotation>Let(_pat_let52_0, boxed105 -> {
          Annotation _2_dt__update__tmp_h0 = ((Annotation)(java.lang.Object)(boxed105));
          return ((Annotation)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Entry>, Annotation>Let(__default.removeEntry((_1_an).dtor_entries(), _0_eid), boxed106 -> {
            dafny.DafnySequence<? extends Entry> _pat_let53_0 = ((dafny.DafnySequence<? extends Entry>)(java.lang.Object)(boxed106));
            return ((Annotation)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Entry>, Annotation>Let(_pat_let53_0, boxed107 -> {
              dafny.DafnySequence<? extends Entry> _3_dt__update_hentries_h0 = ((dafny.DafnySequence<? extends Entry>)(java.lang.Object)(boxed107));
              return Annotation.create((_2_dt__update__tmp_h0).dtor_id(), (_2_dt__update__tmp_h0).dtor_turnId(), (_2_dt__update__tmp_h0).dtor_resource(), (_2_dt__update__tmp_h0).dtor_range(), (_2_dt__update__tmp_h0).dtor_resolved(), _3_dt__update_hentries_h0, (_2_dt__update__tmp_h0).dtor_meta());
            }
            )));
          }
          )));
        }
        )));
      }
      )));
    });}).apply(eid);
  }
  public static java.util.function.Function<Annotation, Annotation> reanchor(agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> tid, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> res, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> rng, agency.open.ahp.AhpSkeleton.Option<Boolean> rslv)
  {
    return ((dafny.Function4<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, agency.open.ahp.AhpSkeleton.Option<Boolean>, java.util.function.Function<Annotation, Annotation>>)(_0_tid, _1_res, _2_rng, _3_rslv) -> {return ((java.util.function.Function<Annotation, Annotation>)(_4_an_boxed0) -> {
      Annotation _4_an = ((Annotation)(java.lang.Object)(_4_an_boxed0));
      return __default.applyUpdate(_4_an, _0_tid, _1_res, _2_rng, _3_rslv);
    });}).apply(tid, res, rng, rslv);
  }
  public static Annotation applyUpdate(Annotation ann, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> tid, agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> res, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> rng, agency.open.ahp.AhpSkeleton.Option<Boolean> rslv)
  {
    Annotation _0_dt__update__tmp_h0 = ann;
    boolean _1_dt__update_hresolved_h0 = (((rslv).is_Some()) ? ((rslv).dtor_value()) : ((ann).dtor_resolved()));
    agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _2_dt__update_hrange_h0 = (((rng).is_Some()) ? (agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), (rng).dtor_value())) : ((ann).dtor_range()));
    dafny.DafnySequence<? extends dafny.CodePoint> _3_dt__update_hresource_h0 = (((res).is_Some()) ? ((res).dtor_value()) : ((ann).dtor_resource()));
    dafny.DafnySequence<? extends dafny.CodePoint> _4_dt__update_hturnId_h0 = (((tid).is_Some()) ? ((tid).dtor_value()) : ((ann).dtor_turnId()));
    return Annotation.create((_0_dt__update__tmp_h0).dtor_id(), _4_dt__update_hturnId_h0, _3_dt__update_hresource_h0, _2_dt__update_hrange_h0, _1_dt__update_hresolved_h0, (_0_dt__update__tmp_h0).dtor_entries(), (_0_dt__update__tmp_h0).dtor_meta());
  }
  public static dafny.Tuple2<dafny.DafnySequence<? extends Annotation>, agency.open.ahp.AhpSkeleton.ReduceOutcome> ApplyToAnnotations(dafny.DafnySequence<? extends Annotation> s, AnnotationsAction a, java.math.BigInteger now)
  {
    AnnotationsAction _source0 = a;
    if (_source0.is_Set()) {
      Annotation _0___mcc_h0 = ((agency.open.ahp.Annotations.AnnotationsAction_Set)_source0)._annotation;
      Annotation _1_ann = _0___mcc_h0;
      return dafny.Tuple2.<dafny.DafnySequence<? extends Annotation>, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(__default.upsertAnn((s), _1_ann), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_Removed()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _2___mcc_h1 = ((agency.open.ahp.Annotations.AnnotationsAction_Removed)_source0)._annotationId;
      dafny.DafnySequence<? extends dafny.CodePoint> _3_id = _2___mcc_h1;
      return dafny.Tuple2.<dafny.DafnySequence<? extends Annotation>, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(__default.removeAnn((s), _3_id), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_EntrySet()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h2 = ((agency.open.ahp.Annotations.AnnotationsAction_EntrySet)_source0)._annotationId;
      Entry _5___mcc_h3 = ((agency.open.ahp.Annotations.AnnotationsAction_EntrySet)_source0)._entry;
      Entry _6_e = _5___mcc_h3;
      dafny.DafnySequence<? extends dafny.CodePoint> _7_aid = _4___mcc_h2;
      if (__default.hasAnn((s), _7_aid)) {
        return dafny.Tuple2.<dafny.DafnySequence<? extends Annotation>, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(__default.updateAnn((s), _7_aid, __default.setEntry(_6_e)), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<dafny.DafnySequence<? extends Annotation>, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_EntryRemoved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _8___mcc_h4 = ((agency.open.ahp.Annotations.AnnotationsAction_EntryRemoved)_source0)._annotationId;
      dafny.DafnySequence<? extends dafny.CodePoint> _9___mcc_h5 = ((agency.open.ahp.Annotations.AnnotationsAction_EntryRemoved)_source0)._entryId;
      dafny.DafnySequence<? extends dafny.CodePoint> _10_eid = _9___mcc_h5;
      dafny.DafnySequence<? extends dafny.CodePoint> _11_aid = _8___mcc_h4;
      if (__default.hasAnn((s), _11_aid)) {
        return dafny.Tuple2.<dafny.DafnySequence<? extends Annotation>, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(__default.updateAnn((s), _11_aid, __default.dropEntry(_10_eid)), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<dafny.DafnySequence<? extends Annotation>, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else if (_source0.is_Updated()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _12___mcc_h6 = ((agency.open.ahp.Annotations.AnnotationsAction_Updated)_source0)._annotationId;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _13___mcc_h7 = ((agency.open.ahp.Annotations.AnnotationsAction_Updated)_source0)._turnId;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _14___mcc_h8 = ((agency.open.ahp.Annotations.AnnotationsAction_Updated)_source0)._resource;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _15___mcc_h9 = ((agency.open.ahp.Annotations.AnnotationsAction_Updated)_source0)._range;
      agency.open.ahp.AhpSkeleton.Option<Boolean> _16___mcc_h10 = ((agency.open.ahp.Annotations.AnnotationsAction_Updated)_source0)._resolved;
      agency.open.ahp.AhpSkeleton.Option<Boolean> _17_rslv = _16___mcc_h10;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _18_rng = _15___mcc_h9;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _19_res = _14___mcc_h8;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _20_tid = _13___mcc_h7;
      dafny.DafnySequence<? extends dafny.CodePoint> _21_aid = _12___mcc_h6;
      if (__default.hasAnn((s), _21_aid)) {
        return dafny.Tuple2.<dafny.DafnySequence<? extends Annotation>, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(__default.updateAnn((s), _21_aid, __default.reanchor(_20_tid, _19_res, _18_rng, _17_rslv)), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<dafny.DafnySequence<? extends Annotation>, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else {
      agency.open.ahp.ConfluxCodec.Json _22___mcc_h11 = ((agency.open.ahp.Annotations.AnnotationsAction_AnUnknown)_source0)._raw;
      return dafny.Tuple2.<dafny.DafnySequence<? extends Annotation>, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
    }
  }
  public static dafny.DafnySequence<? extends Annotation> apply1(dafny.DafnySequence<? extends Annotation> s, AnnotationsAction a)
  {
    return (__default.ApplyToAnnotations(s, a, java.math.BigInteger.valueOf(9999L))).dtor__0();
  }
  public static dafny.DafnySequence<? extends Annotation> fold(dafny.DafnySequence<? extends Annotation> s, dafny.DafnySequence<? extends AnnotationsAction> acts)
  {
    return agency.open.ahp.ConfluxContract.__default.<dafny.DafnySequence<? extends Annotation>, AnnotationsAction>Fold(dafny.DafnySequence.<Annotation>_typeDescriptor(Annotation._typeDescriptor()), AnnotationsAction._typeDescriptor(), __default::apply1, s, acts);
  }
  public static agency.open.ahp.ConfluxCodec.Json R() {
    return agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("start"), agency.open.ahp.ConfluxCodec.Json.create_JNull()) }));
  }
  public static dafny.Tuple2 RunCorpus()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    java.math.BigInteger total = java.math.BigInteger.ZERO;
    if(true) {
      total = java.math.BigInteger.valueOf(10L);
      pass = java.math.BigInteger.ZERO;
      Entry _0_e1;
      _0_e1 = Entry.create(dafny.DafnySequence.asUnicodeString("c-1"), dafny.DafnySequence.asUnicodeString("original"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      Annotation _1_a1;
      _1_a1 = Annotation.create(dafny.DafnySequence.asUnicodeString("t-1"), dafny.DafnySequence.asUnicodeString("turn-1"), dafny.DafnySequence.asUnicodeString("file:///src/a.ts"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), __default.R()), false, dafny.DafnySequence.<Entry> of(Entry._typeDescriptor(), _0_e1), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      dafny.DafnySequence<? extends Annotation> _2_st1;
      _2_st1 = dafny.DafnySequence.<Annotation> of(Annotation._typeDescriptor(), _1_a1);
      Annotation _3_a2;
      _3_a2 = Annotation.create(dafny.DafnySequence.asUnicodeString("t-2"), dafny.DafnySequence.asUnicodeString("turn-2"), dafny.DafnySequence.asUnicodeString("file:///src/b.ts"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), false, dafny.DafnySequence.<Entry> of(Entry._typeDescriptor(), Entry.create(dafny.DafnySequence.asUnicodeString("c-2"), dafny.DafnySequence.asUnicodeString("x"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      if (java.util.Objects.equals(__default.apply1(_2_st1, AnnotationsAction.create_Set(_3_a2)), dafny.DafnySequence.<Annotation> of(Annotation._typeDescriptor(), _1_a1, _3_a2))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Annotation _4_a1r;
      _4_a1r = Annotation.create(dafny.DafnySequence.asUnicodeString("t-1"), dafny.DafnySequence.asUnicodeString("turn-1"), dafny.DafnySequence.asUnicodeString("file:///src/a.ts"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), true, dafny.DafnySequence.<Entry> of(Entry._typeDescriptor(), _0_e1), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      if (java.util.Objects.equals(__default.apply1(_2_st1, AnnotationsAction.create_Set(_4_a1r)), dafny.DafnySequence.<Annotation> of(Annotation._typeDescriptor(), _4_a1r))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_2_st1, AnnotationsAction.create_Removed(dafny.DafnySequence.asUnicodeString("t-1"))), dafny.DafnySequence.<Annotation> empty(Annotation._typeDescriptor()))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_2_st1, AnnotationsAction.create_Removed(dafny.DafnySequence.asUnicodeString("nope"))), _2_st1)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Entry _5_e2;
      _5_e2 = Entry.create(dafny.DafnySequence.asUnicodeString("c-2"), dafny.DafnySequence.asUnicodeString("reply"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      Entry _pat_let_tv0 = _0_e1;
      Entry _pat_let_tv1 = _5_e2;
      if (java.util.Objects.equals(__default.apply1(_2_st1, AnnotationsAction.create_EntrySet(dafny.DafnySequence.asUnicodeString("t-1"), _5_e2)), dafny.DafnySequence.<Annotation> of(Annotation._typeDescriptor(), ((Annotation)(java.lang.Object)(dafny.Helpers.<Annotation, Annotation>Let(_1_a1, boxed108 -> {
        Annotation _pat_let54_0 = ((Annotation)(java.lang.Object)(boxed108));
        return ((Annotation)(java.lang.Object)(dafny.Helpers.<Annotation, Annotation>Let(_pat_let54_0, boxed109 -> {
          Annotation _6_dt__update__tmp_h0 = ((Annotation)(java.lang.Object)(boxed109));
          return ((Annotation)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Entry>, Annotation>Let(dafny.DafnySequence.<Entry> of(Entry._typeDescriptor(), _pat_let_tv0, _pat_let_tv1), boxed110 -> {
            dafny.DafnySequence<? extends Entry> _pat_let55_0 = ((dafny.DafnySequence<? extends Entry>)(java.lang.Object)(boxed110));
            return ((Annotation)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Entry>, Annotation>Let(_pat_let55_0, boxed111 -> {
              dafny.DafnySequence<? extends Entry> _7_dt__update_hentries_h0 = ((dafny.DafnySequence<? extends Entry>)(java.lang.Object)(boxed111));
              return Annotation.create((_6_dt__update__tmp_h0).dtor_id(), (_6_dt__update__tmp_h0).dtor_turnId(), (_6_dt__update__tmp_h0).dtor_resource(), (_6_dt__update__tmp_h0).dtor_range(), (_6_dt__update__tmp_h0).dtor_resolved(), _7_dt__update_hentries_h0, (_6_dt__update__tmp_h0).dtor_meta());
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
      if (java.util.Objects.equals(__default.apply1(_2_st1, AnnotationsAction.create_EntrySet(dafny.DafnySequence.asUnicodeString("nope"), _5_e2)), _2_st1)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      Annotation _8_a1two;
      Annotation _9_dt__update__tmp_h1 = _1_a1;
      dafny.DafnySequence<? extends Entry> _10_dt__update_hentries_h1 = dafny.DafnySequence.<Entry> of(Entry._typeDescriptor(), _0_e1, _5_e2);
      _8_a1two = Annotation.create((_9_dt__update__tmp_h1).dtor_id(), (_9_dt__update__tmp_h1).dtor_turnId(), (_9_dt__update__tmp_h1).dtor_resource(), (_9_dt__update__tmp_h1).dtor_range(), (_9_dt__update__tmp_h1).dtor_resolved(), _10_dt__update_hentries_h1, (_9_dt__update__tmp_h1).dtor_meta());
      if (java.util.Objects.equals(__default.apply1(dafny.DafnySequence.<Annotation> of(Annotation._typeDescriptor(), _8_a1two), AnnotationsAction.create_EntryRemoved(dafny.DafnySequence.asUnicodeString("t-1"), dafny.DafnySequence.asUnicodeString("c-2"))), dafny.DafnySequence.<Annotation> of(Annotation._typeDescriptor(), _1_a1))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_2_st1, AnnotationsAction.create_Updated(dafny.DafnySequence.asUnicodeString("t-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true))), dafny.DafnySequence.<Annotation> of(Annotation._typeDescriptor(), ((Annotation)(java.lang.Object)(dafny.Helpers.<Annotation, Annotation>Let(_1_a1, boxed112 -> {
        Annotation _pat_let56_0 = ((Annotation)(java.lang.Object)(boxed112));
        return ((Annotation)(java.lang.Object)(dafny.Helpers.<Annotation, Annotation>Let(_pat_let56_0, boxed113 -> {
          Annotation _11_dt__update__tmp_h2 = ((Annotation)(java.lang.Object)(boxed113));
          return ((Annotation)(java.lang.Object)(dafny.Helpers.<Boolean, Annotation>Let(true, boxed114 -> {
            boolean _pat_let57_0 = ((boolean)(java.lang.Object)(boxed114));
            return ((Annotation)(java.lang.Object)(dafny.Helpers.<Boolean, Annotation>Let(_pat_let57_0, boxed115 -> {
              boolean _12_dt__update_hresolved_h0 = ((boolean)(java.lang.Object)(boxed115));
              return Annotation.create((_11_dt__update__tmp_h2).dtor_id(), (_11_dt__update__tmp_h2).dtor_turnId(), (_11_dt__update__tmp_h2).dtor_resource(), (_11_dt__update__tmp_h2).dtor_range(), _12_dt__update_hresolved_h0, (_11_dt__update__tmp_h2).dtor_entries(), (_11_dt__update__tmp_h2).dtor_meta());
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
      agency.open.ahp.ConfluxCodec.Json _13_newR;
      _13_newR = agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("start"), agency.open.ahp.ConfluxCodec.Json.create_JBool(true)) }));
      Annotation _14_a1re;
      Annotation _15_dt__update__tmp_h3 = _1_a1;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _16_dt__update_hrange_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _13_newR);
      dafny.DafnySequence<? extends dafny.CodePoint> _17_dt__update_hresource_h0 = dafny.DafnySequence.asUnicodeString("file:///src/b.ts");
      dafny.DafnySequence<? extends dafny.CodePoint> _18_dt__update_hturnId_h0 = dafny.DafnySequence.asUnicodeString("turn-2");
      _14_a1re = Annotation.create((_15_dt__update__tmp_h3).dtor_id(), _18_dt__update_hturnId_h0, _17_dt__update_hresource_h0, _16_dt__update_hrange_h0, (_15_dt__update__tmp_h3).dtor_resolved(), (_15_dt__update__tmp_h3).dtor_entries(), (_15_dt__update__tmp_h3).dtor_meta());
      if (java.util.Objects.equals(__default.apply1(_2_st1, AnnotationsAction.create_Updated(dafny.DafnySequence.asUnicodeString("t-1"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("turn-2")), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("file:///src/b.ts")), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _13_newR), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN))), dafny.DafnySequence.<Annotation> of(Annotation._typeDescriptor(), _14_a1re))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_2_st1, AnnotationsAction.create_Updated(dafny.DafnySequence.asUnicodeString("nope"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true))), _2_st1)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
    }
    return new dafny.Tuple2<>(pass, total);
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.Annotations._default";
  }
}
