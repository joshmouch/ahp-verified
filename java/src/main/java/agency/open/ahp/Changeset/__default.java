// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.Changeset;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static java.util.function.Function<ChangesetOperation, ChangesetOperation> opSet(dafny.DafnySequence<? extends dafny.CodePoint> st, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> err)
  {
    return ((dafny.Function2<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, dafny.DafnySequence<? extends dafny.CodePoint>, java.util.function.Function<ChangesetOperation, ChangesetOperation>>)(_0_err, _1_st) -> {return ((java.util.function.Function<ChangesetOperation, ChangesetOperation>)(_2_e_boxed0) -> {
      ChangesetOperation _2_e = ((ChangesetOperation)(java.lang.Object)(_2_e_boxed0));
      return ((ChangesetOperation)(java.lang.Object)(dafny.Helpers.<ChangesetOperation, ChangesetOperation>Let(_2_e, boxed48 -> {
        ChangesetOperation _pat_let24_0 = ((ChangesetOperation)(java.lang.Object)(boxed48));
        return ((ChangesetOperation)(java.lang.Object)(dafny.Helpers.<ChangesetOperation, ChangesetOperation>Let(_pat_let24_0, boxed49 -> {
          ChangesetOperation _3_dt__update__tmp_h0 = ((ChangesetOperation)(java.lang.Object)(boxed49));
          return ((ChangesetOperation)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ChangesetOperation>Let(_0_err, boxed50 -> {
            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let25_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed50));
            return ((ChangesetOperation)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ChangesetOperation>Let(_pat_let25_0, boxed51 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _4_dt__update_herror_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed51));
              return ((ChangesetOperation)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ChangesetOperation>Let(_1_st, boxed52 -> {
                dafny.DafnySequence<? extends dafny.CodePoint> _pat_let26_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed52));
                return ((ChangesetOperation)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ChangesetOperation>Let(_pat_let26_0, boxed53 -> {
                  dafny.DafnySequence<? extends dafny.CodePoint> _5_dt__update_hstatus_h0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed53));
                  return ChangesetOperation.create((_3_dt__update__tmp_h0).dtor_id(), (_3_dt__update__tmp_h0).dtor_label__(), (_3_dt__update__tmp_h0).dtor_scopes(), _5_dt__update_hstatus_h0, _4_dt__update_herror_h0);
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
    });}).apply(err, st);
  }
  public static dafny.DafnySequence<? extends ChangesetFile> upsertFile(dafny.DafnySequence<? extends ChangesetFile> files, ChangesetFile f)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, ChangesetFile>SeqUpsertById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), ChangesetFile._typeDescriptor(), __default.fileKey(), files, f);
  }
  public static dafny.DafnySequence<? extends ChangesetFile> removeFile(dafny.DafnySequence<? extends ChangesetFile> files, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, ChangesetFile>SeqRemoveById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), ChangesetFile._typeDescriptor(), __default.fileKey(), files, id);
  }
  public static dafny.DafnySequence<? extends ChangesetFile> setReviewed(dafny.DafnySequence<? extends ChangesetFile> files, dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> ids, boolean rev)
  {
    return agency.open.ahp.ConfluxOperators.__default.<ChangesetFile, ChangesetFile>Map(ChangesetFile._typeDescriptor(), ChangesetFile._typeDescriptor(), ((dafny.Function2<dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>>, Boolean, java.util.function.Function<ChangesetFile, ChangesetFile>>)(_0_ids, __1_rev0) -> {boolean _1_rev = ((boolean)(java.lang.Object)(__1_rev0));
    return ((java.util.function.Function<ChangesetFile, ChangesetFile>)(_2_f_boxed0) -> {
      ChangesetFile _2_f = ((ChangesetFile)(java.lang.Object)(_2_f_boxed0));
      return (((_0_ids).contains((_2_f).dtor_id())) ? (((ChangesetFile)(java.lang.Object)(dafny.Helpers.<ChangesetFile, ChangesetFile>Let(_2_f, boxed54 -> {
        ChangesetFile _pat_let27_0 = ((ChangesetFile)(java.lang.Object)(boxed54));
        return ((ChangesetFile)(java.lang.Object)(dafny.Helpers.<ChangesetFile, ChangesetFile>Let(_pat_let27_0, boxed55 -> {
          ChangesetFile _3_dt__update__tmp_h0 = ((ChangesetFile)(java.lang.Object)(boxed55));
          return ((ChangesetFile)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, ChangesetFile>Let(agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, _1_rev), boxed56 -> {
            agency.open.ahp.AhpSkeleton.Option<Boolean> _pat_let28_0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed56));
            return ((ChangesetFile)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, ChangesetFile>Let(_pat_let28_0, boxed57 -> {
              agency.open.ahp.AhpSkeleton.Option<Boolean> _4_dt__update_hreviewed_h0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed57));
              return ChangesetFile.create((_3_dt__update__tmp_h0).dtor_id(), _4_dt__update_hreviewed_h0, (_3_dt__update__tmp_h0).dtor_edit());
            }
            )));
          }
          )));
        }
        )));
      }
      )))) : (_2_f));
    });}).apply(ids, rev), files);
  }
  public static dafny.DafnySequence<? extends ChangesetOperation> updateOp(dafny.DafnySequence<? extends ChangesetOperation> ops, dafny.DafnySequence<? extends dafny.CodePoint> id, dafny.DafnySequence<? extends dafny.CodePoint> status, agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> error)
  {
    return agency.open.ahp.ConfluxSeqRoute.__default.<dafny.DafnySequence<? extends dafny.CodePoint>, ChangesetOperation>SeqUpdateById(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), ChangesetOperation._typeDescriptor(), __default.opKey(), ops, id, __default.opSet(status, error));
  }
  public static boolean hasOp(dafny.DafnySequence<? extends ChangesetOperation> ops, dafny.DafnySequence<? extends dafny.CodePoint> id)
  {
    return ((dafny.Function2<dafny.DafnySequence<? extends ChangesetOperation>, dafny.DafnySequence<? extends dafny.CodePoint>, Boolean>)(_0_ops, _1_id) -> {return dafny.Helpers.Quantifier(dafny.Helpers.IntegerRange(java.math.BigInteger.ZERO, java.math.BigInteger.valueOf((_0_ops).length())), false, ((_exists_var_0_boxed0) -> {
      java.math.BigInteger _exists_var_0 = ((java.math.BigInteger)(java.lang.Object)(_exists_var_0_boxed0));
      java.math.BigInteger _2_i = (java.math.BigInteger)_exists_var_0;
      return (((_2_i).signum() != -1) && ((_2_i).compareTo(java.math.BigInteger.valueOf((_0_ops).length())) < 0)) && (((((ChangesetOperation)(java.lang.Object)((_0_ops).select(dafny.Helpers.toInt((_2_i)))))).dtor_id()).equals(_1_id));
    }));}).apply(ops, id);
  }
  public static boolean anyNeedsReviewed(dafny.DafnySequence<? extends ChangesetFile> files, dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> ids, boolean rev)
  {
    return ((dafny.Function3<dafny.DafnySequence<? extends ChangesetFile>, dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>>, Boolean, Boolean>)(_0_files, _1_ids, __2_rev0) -> {boolean _2_rev = ((boolean)(java.lang.Object)(__2_rev0));
    return dafny.Helpers.Quantifier(dafny.Helpers.IntegerRange(java.math.BigInteger.ZERO, java.math.BigInteger.valueOf((_0_files).length())), false, ((_exists_var_0_boxed0) -> {
      java.math.BigInteger _exists_var_0 = ((java.math.BigInteger)(java.lang.Object)(_exists_var_0_boxed0));
      java.math.BigInteger _3_i = (java.math.BigInteger)_exists_var_0;
      return ((((_3_i).signum() != -1) && ((_3_i).compareTo(java.math.BigInteger.valueOf((_0_files).length())) < 0)) && ((_1_ids).contains((((ChangesetFile)(java.lang.Object)((_0_files).select(dafny.Helpers.toInt((_3_i)))))).dtor_id()))) && (!java.util.Objects.equals((((ChangesetFile)(java.lang.Object)((_0_files).select(dafny.Helpers.toInt((_3_i)))))).dtor_reviewed(), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, _2_rev)));
    }));}).apply(files, ids, rev);
  }
  public static dafny.Tuple2<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome> ApplyToChangeset(ChangesetState s, ChangesetAction a, java.math.BigInteger now)
  {
    ChangesetState _pat_let_tv0 = s;
    ChangesetState _pat_let_tv1 = s;
    ChangesetState _pat_let_tv2 = s;
    ChangesetAction _source0 = a;
    if (_source0.is_StatusChanged()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _0___mcc_h0 = ((agency.open.ahp.Changeset.ChangesetAction_StatusChanged)_source0)._status;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _1___mcc_h1 = ((agency.open.ahp.Changeset.ChangesetAction_StatusChanged)_source0)._error;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _2_err = _1___mcc_h1;
      dafny.DafnySequence<? extends dafny.CodePoint> _3_st = _0___mcc_h0;
      return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(s, boxed58 -> {
  ChangesetState _pat_let29_0 = ((ChangesetState)(java.lang.Object)(boxed58));
  return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_pat_let29_0, boxed59 -> {
    ChangesetState _4_dt__update__tmp_h0 = ((ChangesetState)(java.lang.Object)(boxed59));
    return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ChangesetState>Let(_2_err, boxed60 -> {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let30_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed60));
      return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ChangesetState>Let(_pat_let30_0, boxed61 -> {
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _5_dt__update_herror_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed61));
        return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ChangesetState>Let(_3_st, boxed62 -> {
          dafny.DafnySequence<? extends dafny.CodePoint> _pat_let31_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed62));
          return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, ChangesetState>Let(_pat_let31_0, boxed63 -> {
            dafny.DafnySequence<? extends dafny.CodePoint> _6_dt__update_hstatus_h0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed63));
            return ChangesetState.create(_6_dt__update_hstatus_h0, (_4_dt__update__tmp_h0).dtor_files(), (_4_dt__update__tmp_h0).dtor_operations(), _5_dt__update_herror_h0);
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
    } else if (_source0.is_FileSet()) {
      ChangesetFile _7___mcc_h2 = ((agency.open.ahp.Changeset.ChangesetAction_FileSet)_source0)._file;
      ChangesetFile _8_f = _7___mcc_h2;
      return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(s, boxed64 -> {
  ChangesetState _pat_let32_0 = ((ChangesetState)(java.lang.Object)(boxed64));
  return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_pat_let32_0, boxed65 -> {
    ChangesetState _9_dt__update__tmp_h1 = ((ChangesetState)(java.lang.Object)(boxed65));
    return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends ChangesetFile>, ChangesetState>Let(__default.upsertFile((_pat_let_tv0).dtor_files(), _8_f), boxed66 -> {
      dafny.DafnySequence<? extends ChangesetFile> _pat_let33_0 = ((dafny.DafnySequence<? extends ChangesetFile>)(java.lang.Object)(boxed66));
      return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends ChangesetFile>, ChangesetState>Let(_pat_let33_0, boxed67 -> {
        dafny.DafnySequence<? extends ChangesetFile> _10_dt__update_hfiles_h0 = ((dafny.DafnySequence<? extends ChangesetFile>)(java.lang.Object)(boxed67));
        return ChangesetState.create((_9_dt__update__tmp_h1).dtor_status(), _10_dt__update_hfiles_h0, (_9_dt__update__tmp_h1).dtor_operations(), (_9_dt__update__tmp_h1).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_FileRemoved()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _11___mcc_h3 = ((agency.open.ahp.Changeset.ChangesetAction_FileRemoved)_source0)._fileId;
      dafny.DafnySequence<? extends dafny.CodePoint> _12_id = _11___mcc_h3;
      return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(s, boxed68 -> {
  ChangesetState _pat_let34_0 = ((ChangesetState)(java.lang.Object)(boxed68));
  return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_pat_let34_0, boxed69 -> {
    ChangesetState _13_dt__update__tmp_h2 = ((ChangesetState)(java.lang.Object)(boxed69));
    return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends ChangesetFile>, ChangesetState>Let(__default.removeFile((_pat_let_tv1).dtor_files(), _12_id), boxed70 -> {
      dafny.DafnySequence<? extends ChangesetFile> _pat_let35_0 = ((dafny.DafnySequence<? extends ChangesetFile>)(java.lang.Object)(boxed70));
      return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends ChangesetFile>, ChangesetState>Let(_pat_let35_0, boxed71 -> {
        dafny.DafnySequence<? extends ChangesetFile> _14_dt__update_hfiles_h1 = ((dafny.DafnySequence<? extends ChangesetFile>)(java.lang.Object)(boxed71));
        return ChangesetState.create((_13_dt__update__tmp_h2).dtor_status(), _14_dt__update_hfiles_h1, (_13_dt__update__tmp_h2).dtor_operations(), (_13_dt__update__tmp_h2).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_OperationsChanged()) {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _15___mcc_h4 = ((agency.open.ahp.Changeset.ChangesetAction_OperationsChanged)_source0)._operations;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _16_ops = _15___mcc_h4;
      return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(s, boxed72 -> {
  ChangesetState _pat_let36_0 = ((ChangesetState)(java.lang.Object)(boxed72));
  return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_pat_let36_0, boxed73 -> {
    ChangesetState _17_dt__update__tmp_h3 = ((ChangesetState)(java.lang.Object)(boxed73));
    return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>, ChangesetState>Let(_16_ops, boxed74 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _pat_let37_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>)(java.lang.Object)(boxed74));
      return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>, ChangesetState>Let(_pat_let37_0, boxed75 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _18_dt__update_hoperations_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>)(java.lang.Object)(boxed75));
        return ChangesetState.create((_17_dt__update__tmp_h3).dtor_status(), (_17_dt__update__tmp_h3).dtor_files(), _18_dt__update_hoperations_h0, (_17_dt__update__tmp_h3).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_Cleared()) {
      return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(s, boxed76 -> {
  ChangesetState _pat_let38_0 = ((ChangesetState)(java.lang.Object)(boxed76));
  return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_pat_let38_0, boxed77 -> {
    ChangesetState _19_dt__update__tmp_h4 = ((ChangesetState)(java.lang.Object)(boxed77));
    return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends ChangesetFile>, ChangesetState>Let(dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), boxed78 -> {
      dafny.DafnySequence<? extends ChangesetFile> _pat_let39_0 = ((dafny.DafnySequence<? extends ChangesetFile>)(java.lang.Object)(boxed78));
      return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends ChangesetFile>, ChangesetState>Let(_pat_let39_0, boxed79 -> {
        dafny.DafnySequence<? extends ChangesetFile> _20_dt__update_hfiles_h2 = ((dafny.DafnySequence<? extends ChangesetFile>)(java.lang.Object)(boxed79));
        return ChangesetState.create((_19_dt__update__tmp_h4).dtor_status(), _20_dt__update_hfiles_h2, (_19_dt__update__tmp_h4).dtor_operations(), (_19_dt__update__tmp_h4).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))), (((java.math.BigInteger.valueOf(((s).dtor_files()).length())).signum() == 0) ? (agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp()) : (agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied())));
    } else if (_source0.is_OperationStatusChanged()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _21___mcc_h5 = ((agency.open.ahp.Changeset.ChangesetAction_OperationStatusChanged)_source0)._operationId;
      dafny.DafnySequence<? extends dafny.CodePoint> _22___mcc_h6 = ((agency.open.ahp.Changeset.ChangesetAction_OperationStatusChanged)_source0)._status;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _23___mcc_h7 = ((agency.open.ahp.Changeset.ChangesetAction_OperationStatusChanged)_source0)._error;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _24_err = _23___mcc_h7;
      dafny.DafnySequence<? extends dafny.CodePoint> _25_st = _22___mcc_h6;
      dafny.DafnySequence<? extends dafny.CodePoint> _26_id = _21___mcc_h5;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _source1 = (s).dtor_operations();
      if (_source1.is_None()) {
        return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      } else {
        dafny.DafnySequence<? extends ChangesetOperation> _27___mcc_h14 = ((agency.open.ahp.AhpSkeleton.Option_Some<dafny.DafnySequence<? extends ChangesetOperation>>)_source1)._value;
        dafny.DafnySequence<? extends ChangesetOperation> _28_ops = _27___mcc_h14;
        if (__default.hasOp(_28_ops, _26_id)) {
          return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(s, boxed80 -> {
  ChangesetState _pat_let40_0 = ((ChangesetState)(java.lang.Object)(boxed80));
  return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_pat_let40_0, boxed81 -> {
    ChangesetState _29_dt__update__tmp_h5 = ((ChangesetState)(java.lang.Object)(boxed81));
    return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>, ChangesetState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_Some(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor()), __default.updateOp(_28_ops, _26_id, _25_st, _24_err)), boxed82 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _pat_let41_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>)(java.lang.Object)(boxed82));
      return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>, ChangesetState>Let(_pat_let41_0, boxed83 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _30_dt__update_hoperations_h1 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>)(java.lang.Object)(boxed83));
        return ChangesetState.create((_29_dt__update__tmp_h5).dtor_status(), (_29_dt__update__tmp_h5).dtor_files(), _30_dt__update_hoperations_h1, (_29_dt__update__tmp_h5).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
        } else {
          return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
        }
      }
    } else if (_source0.is_ContentChanged()) {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetFile>> _31___mcc_h8 = ((agency.open.ahp.Changeset.ChangesetAction_ContentChanged)_source0)._files;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _32___mcc_h9 = ((agency.open.ahp.Changeset.ChangesetAction_ContentChanged)_source0)._operations;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _33___mcc_h10 = ((agency.open.ahp.Changeset.ChangesetAction_ContentChanged)_source0)._error;
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _34_err = _33___mcc_h10;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _35_ops = _32___mcc_h9;
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetFile>> _36_fs = _31___mcc_h8;
      ChangesetState _37_s1 = (((_36_fs).is_Some()) ? (((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(s, boxed84 -> {
        ChangesetState _pat_let42_0 = ((ChangesetState)(java.lang.Object)(boxed84));
        return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_pat_let42_0, boxed85 -> {
          ChangesetState _38_dt__update__tmp_h6 = ((ChangesetState)(java.lang.Object)(boxed85));
          return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends ChangesetFile>, ChangesetState>Let((_36_fs).dtor_value(), boxed86 -> {
            dafny.DafnySequence<? extends ChangesetFile> _pat_let43_0 = ((dafny.DafnySequence<? extends ChangesetFile>)(java.lang.Object)(boxed86));
            return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends ChangesetFile>, ChangesetState>Let(_pat_let43_0, boxed87 -> {
              dafny.DafnySequence<? extends ChangesetFile> _39_dt__update_hfiles_h3 = ((dafny.DafnySequence<? extends ChangesetFile>)(java.lang.Object)(boxed87));
              return ChangesetState.create((_38_dt__update__tmp_h6).dtor_status(), _39_dt__update_hfiles_h3, (_38_dt__update__tmp_h6).dtor_operations(), (_38_dt__update__tmp_h6).dtor_error());
            }
            )));
          }
          )));
        }
        )));
      }
      )))) : (s));
      ChangesetState _40_s2 = (((_35_ops).is_Some()) ? (((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_37_s1, boxed88 -> {
        ChangesetState _pat_let44_0 = ((ChangesetState)(java.lang.Object)(boxed88));
        return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_pat_let44_0, boxed89 -> {
          ChangesetState _41_dt__update__tmp_h7 = ((ChangesetState)(java.lang.Object)(boxed89));
          return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>, ChangesetState>Let(_35_ops, boxed90 -> {
            agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _pat_let45_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>)(java.lang.Object)(boxed90));
            return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>, ChangesetState>Let(_pat_let45_0, boxed91 -> {
              agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>> _42_dt__update_hoperations_h2 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends ChangesetOperation>>)(java.lang.Object)(boxed91));
              return ChangesetState.create((_41_dt__update__tmp_h7).dtor_status(), (_41_dt__update__tmp_h7).dtor_files(), _42_dt__update_hoperations_h2, (_41_dt__update__tmp_h7).dtor_error());
            }
            )));
          }
          )));
        }
        )));
      }
      )))) : (_37_s1));
      ChangesetState _43_s3 = (((_34_err).is_Some()) ? (((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_40_s2, boxed92 -> {
        ChangesetState _pat_let46_0 = ((ChangesetState)(java.lang.Object)(boxed92));
        return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_pat_let46_0, boxed93 -> {
          ChangesetState _44_dt__update__tmp_h8 = ((ChangesetState)(java.lang.Object)(boxed93));
          return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ChangesetState>Let(_34_err, boxed94 -> {
            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let47_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed94));
            return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, ChangesetState>Let(_pat_let47_0, boxed95 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _45_dt__update_herror_h1 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed95));
              return ChangesetState.create((_44_dt__update__tmp_h8).dtor_status(), (_44_dt__update__tmp_h8).dtor_files(), (_44_dt__update__tmp_h8).dtor_operations(), _45_dt__update_herror_h1);
            }
            )));
          }
          )));
        }
        )));
      }
      )))) : (_40_s2));
      return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(_43_s3, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_FilesReviewedChanged()) {
      dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> _46___mcc_h11 = ((agency.open.ahp.Changeset.ChangesetAction_FilesReviewedChanged)_source0)._fileIds;
      boolean _47___mcc_h12 = ((agency.open.ahp.Changeset.ChangesetAction_FilesReviewedChanged)_source0)._reviewed;
      boolean _48_rev = _47___mcc_h12;
      dafny.DafnySequence<? extends dafny.DafnySequence<? extends dafny.CodePoint>> _49_ids = _46___mcc_h11;
      if (__default.anyNeedsReviewed((s).dtor_files(), _49_ids, _48_rev)) {
        return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(s, boxed96 -> {
  ChangesetState _pat_let48_0 = ((ChangesetState)(java.lang.Object)(boxed96));
  return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<ChangesetState, ChangesetState>Let(_pat_let48_0, boxed97 -> {
    ChangesetState _50_dt__update__tmp_h9 = ((ChangesetState)(java.lang.Object)(boxed97));
    return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends ChangesetFile>, ChangesetState>Let(__default.setReviewed((_pat_let_tv2).dtor_files(), _49_ids, _48_rev), boxed98 -> {
      dafny.DafnySequence<? extends ChangesetFile> _pat_let49_0 = ((dafny.DafnySequence<? extends ChangesetFile>)(java.lang.Object)(boxed98));
      return ((ChangesetState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends ChangesetFile>, ChangesetState>Let(_pat_let49_0, boxed99 -> {
        dafny.DafnySequence<? extends ChangesetFile> _51_dt__update_hfiles_h4 = ((dafny.DafnySequence<? extends ChangesetFile>)(java.lang.Object)(boxed99));
        return ChangesetState.create((_50_dt__update__tmp_h9).dtor_status(), _51_dt__update_hfiles_h4, (_50_dt__update__tmp_h9).dtor_operations(), (_50_dt__update__tmp_h9).dtor_error());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
      } else {
        return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
      }
    } else {
      agency.open.ahp.ConfluxCodec.Json _52___mcc_h13 = ((agency.open.ahp.Changeset.ChangesetAction_CsUnknown)_source0)._raw;
      return dafny.Tuple2.<ChangesetState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
    }
  }
  public static ChangesetState apply1(ChangesetState s, ChangesetAction a)
  {
    return (__default.ApplyToChangeset(s, a, java.math.BigInteger.valueOf(9999L))).dtor__0();
  }
  public static ChangesetState fold(ChangesetState s, dafny.DafnySequence<? extends ChangesetAction> acts)
  {
    return agency.open.ahp.ConfluxContract.__default.<ChangesetState, ChangesetAction>Fold(ChangesetState._typeDescriptor(), ChangesetAction._typeDescriptor(), __default::apply1, s, acts);
  }
  public static agency.open.ahp.ConfluxCodec.Json E() {
    return agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("diff"), agency.open.ahp.ConfluxCodec.Json.create_JNull()) }));
  }
  public static dafny.Tuple2 RunCorpus()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    java.math.BigInteger total = java.math.BigInteger.ZERO;
    if(true) {
      total = java.math.BigInteger.valueOf(15L);
      pass = java.math.BigInteger.ZERO;
      ChangesetFile _0_fa;
      _0_fa = ChangesetFile.create(dafny.DafnySequence.asUnicodeString("file:///src/a.ts"), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), __default.E());
      ChangesetFile _1_fb;
      _1_fb = ChangesetFile.create(dafny.DafnySequence.asUnicodeString("file:///src/b.ts"), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN), __default.E());
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("computing"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("boom")))), ChangesetAction.create_StatusChanged(dafny.DafnySequence.asUnicodeString("ready"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_StatusChanged(dafny.DafnySequence.asUnicodeString("error"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("x"))))), ChangesetState.create(dafny.DafnySequence.asUnicodeString("error"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("x")))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _0_fa), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_FileSet(_1_fb)), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _0_fa, _1_fb), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChangesetFile _2_fbv2;
      _2_fbv2 = ChangesetFile.create(dafny.DafnySequence.asUnicodeString("file:///src/b.ts"), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), __default.E());
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _0_fa, _1_fb), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_FileSet(_2_fbv2)), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _0_fa, _2_fbv2), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _0_fa, _1_fb), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_FileRemoved(dafny.DafnySequence.asUnicodeString("file:///src/a.ts"))), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _1_fb), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _0_fa), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_FileRemoved(dafny.DafnySequence.asUnicodeString("nope"))), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _0_fa), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChangesetOperation _3_op;
      _3_op = ChangesetOperation.create(dafny.DafnySequence.asUnicodeString("create-pr"), dafny.DafnySequence.asUnicodeString("Create Pull Request"), dafny.DafnySequence.<dafny.DafnySequence<? extends dafny.CodePoint>> of(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("changeset")), dafny.DafnySequence.asUnicodeString("idle"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()));
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_OperationsChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_Some(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor()), dafny.DafnySequence.<ChangesetOperation> of(ChangesetOperation._typeDescriptor(), _3_op)))), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_Some(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor()), dafny.DafnySequence.<ChangesetOperation> of(ChangesetOperation._typeDescriptor(), _3_op)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_Some(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor()), dafny.DafnySequence.<ChangesetOperation> of(ChangesetOperation._typeDescriptor(), _3_op)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_OperationsChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())))), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (((__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _0_fa), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_Cleared())).dtor_files()).equals(dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_CsUnknown(agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("type"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("changeset/nope"))) })))), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChangesetOperation _4_opR;
      ChangesetOperation _5_dt__update__tmp_h0 = _3_op;
      dafny.DafnySequence<? extends dafny.CodePoint> _6_dt__update_hstatus_h0 = dafny.DafnySequence.asUnicodeString("running");
      _4_opR = ChangesetOperation.create((_5_dt__update__tmp_h0).dtor_id(), (_5_dt__update__tmp_h0).dtor_label__(), (_5_dt__update__tmp_h0).dtor_scopes(), _6_dt__update_hstatus_h0, (_5_dt__update__tmp_h0).dtor_error());
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_Some(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor()), dafny.DafnySequence.<ChangesetOperation> of(ChangesetOperation._typeDescriptor(), _3_op)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_OperationStatusChanged(dafny.DafnySequence.asUnicodeString("create-pr"), dafny.DafnySequence.asUnicodeString("running"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_Some(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor()), dafny.DafnySequence.<ChangesetOperation> of(ChangesetOperation._typeDescriptor(), _4_opR)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_Some(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor()), dafny.DafnySequence.<ChangesetOperation> of(ChangesetOperation._typeDescriptor(), _3_op)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_OperationStatusChanged(dafny.DafnySequence.asUnicodeString("nope"), dafny.DafnySequence.asUnicodeString("running"), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> empty(ChangesetFile._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_Some(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor()), dafny.DafnySequence.<ChangesetOperation> of(ChangesetOperation._typeDescriptor(), _3_op)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _0_fa), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_ContentChanged(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetFile>>create_Some(dafny.DafnySequence.<ChangesetFile>_typeDescriptor(ChangesetFile._typeDescriptor()), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _1_fb)), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_Some(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor()), dafny.DafnySequence.<ChangesetOperation> of(ChangesetOperation._typeDescriptor(), _3_op)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()))), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _1_fb), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_Some(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor()), dafny.DafnySequence.<ChangesetOperation> of(ChangesetOperation._typeDescriptor(), _3_op)), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ChangesetFile _7_fc;
      _7_fc = ChangesetFile.create(dafny.DafnySequence.asUnicodeString("file:///src/c.ts"), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, false), __default.E());
      ChangesetFile _8_fbT;
      _8_fbT = ChangesetFile.create(dafny.DafnySequence.asUnicodeString("file:///src/b.ts"), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), __default.E());
      ChangesetFile _9_faT;
      _9_faT = ChangesetFile.create(dafny.DafnySequence.asUnicodeString("file:///src/a.ts"), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), __default.E());
      ChangesetState _10_r160;
      _10_r160 = __default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _0_fa, _8_fbT, _7_fc), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_FilesReviewedChanged(dafny.DafnySequence.<dafny.DafnySequence<? extends dafny.CodePoint>> of(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("file:///src/a.ts"), dafny.DafnySequence.asUnicodeString("file:///src/b.ts"), dafny.DafnySequence.asUnicodeString("file:///src/missing.ts")), true));
      if (java.util.Objects.equals(_10_r160, ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _9_faT, _8_fbT, _7_fc), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _8_fbT), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())), ChangesetAction.create_FilesReviewedChanged(dafny.DafnySequence.<dafny.DafnySequence<? extends dafny.CodePoint>> of(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("file:///src/b.ts")), true)), ChangesetState.create(dafny.DafnySequence.asUnicodeString("ready"), dafny.DafnySequence.<ChangesetFile> of(ChangesetFile._typeDescriptor(), _8_fbT), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends ChangesetOperation>>create_None(dafny.DafnySequence.<ChangesetOperation>_typeDescriptor(ChangesetOperation._typeDescriptor())), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor())))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
    }
    return new dafny.Tuple2<>(pass, total);
  }
  public static java.util.function.Function<ChangesetFile, dafny.DafnySequence<? extends dafny.CodePoint>> fileKey()
  {
    return ((java.util.function.Function<ChangesetFile, dafny.DafnySequence<? extends dafny.CodePoint>>)(_0_x_boxed0) -> {
      ChangesetFile _0_x = ((ChangesetFile)(java.lang.Object)(_0_x_boxed0));
      return (_0_x).dtor_id();
    });
  }
  public static java.util.function.Function<ChangesetOperation, dafny.DafnySequence<? extends dafny.CodePoint>> opKey()
  {
    return ((java.util.function.Function<ChangesetOperation, dafny.DafnySequence<? extends dafny.CodePoint>>)(_0_x_boxed0) -> {
      ChangesetOperation _0_x = ((ChangesetOperation)(java.lang.Object)(_0_x_boxed0));
      return (_0_x).dtor_id();
    });
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.Changeset._default";
  }
}
