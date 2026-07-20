// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.Terminal;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static dafny.DafnySequence<? extends Part> appendData(dafny.DafnySequence<? extends Part> content, dafny.DafnySequence<? extends dafny.CodePoint> d)
  {
    if ((java.math.BigInteger.valueOf((content).length())).signum() == 0) {
      return dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Unclassified(d));
    } else {
      Part _0_last = ((Part)(java.lang.Object)((content).select(dafny.Helpers.toInt(((java.math.BigInteger.valueOf((content).length())).subtract(java.math.BigInteger.ONE))))));
      dafny.DafnySequence<? extends Part> _1_init = (content).take((java.math.BigInteger.valueOf((content).length())).subtract(java.math.BigInteger.ONE));
      Part _source0 = _0_last;
      if (_source0.is_Unclassified()) {
        dafny.DafnySequence<? extends dafny.CodePoint> _2___mcc_h0 = ((agency.open.ahp.Terminal.Part_Unclassified)_source0)._value;
        dafny.DafnySequence<? extends dafny.CodePoint> _3_v = _2___mcc_h0;
        return dafny.DafnySequence.<Part>concatenate(_1_init, dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Unclassified(dafny.DafnySequence.<dafny.CodePoint>concatenate(_3_v, d))));
      } else {
        dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h1 = ((agency.open.ahp.Terminal.Part_Command)_source0)._commandId;
        dafny.DafnySequence<? extends dafny.CodePoint> _5___mcc_h2 = ((agency.open.ahp.Terminal.Part_Command)_source0)._commandLine;
        dafny.DafnySequence<? extends dafny.CodePoint> _6___mcc_h3 = ((agency.open.ahp.Terminal.Part_Command)_source0)._output;
        java.math.BigInteger _7___mcc_h4 = ((agency.open.ahp.Terminal.Part_Command)_source0)._timestamp;
        boolean _8___mcc_h5 = ((agency.open.ahp.Terminal.Part_Command)_source0)._isComplete;
        agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _9___mcc_h6 = ((agency.open.ahp.Terminal.Part_Command)_source0)._exitCode;
        agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _10___mcc_h7 = ((agency.open.ahp.Terminal.Part_Command)_source0)._durationMs;
        agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _11_dm = _10___mcc_h7;
        agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _12_ec = _9___mcc_h6;
        boolean _13_comp = _8___mcc_h5;
        java.math.BigInteger _14_ts = _7___mcc_h4;
        dafny.DafnySequence<? extends dafny.CodePoint> _15_out = _6___mcc_h3;
        dafny.DafnySequence<? extends dafny.CodePoint> _16_cl = _5___mcc_h2;
        dafny.DafnySequence<? extends dafny.CodePoint> _17_cid = _4___mcc_h1;
        if (!(_13_comp)) {
          return dafny.DafnySequence.<Part>concatenate(_1_init, dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Command(_17_cid, _16_cl, dafny.DafnySequence.<dafny.CodePoint>concatenate(_15_out, d), _14_ts, _13_comp, _12_ec, _11_dm)));
        } else {
          return dafny.DafnySequence.<Part>concatenate(content, dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Unclassified(d)));
        }
      }
    }
  }
  public static dafny.DafnySequence<? extends Part> finishCommand(dafny.DafnySequence<? extends Part> content, dafny.DafnySequence<? extends dafny.CodePoint> id, java.math.BigInteger code, java.math.BigInteger dur)
  {
    return agency.open.ahp.ConfluxOperators.__default.<Part, Part>Map(Part._typeDescriptor(), Part._typeDescriptor(), ((dafny.Function3<dafny.DafnySequence<? extends dafny.CodePoint>, java.math.BigInteger, java.math.BigInteger, java.util.function.Function<Part, Part>>)(_0_id, _1_code, _2_dur) -> {return ((java.util.function.Function<Part, Part>)(_3_p_boxed0) -> {
      Part _3_p = ((Part)(java.lang.Object)(_3_p_boxed0));
      return ((java.util.function.Function<Part, Part>)(_source0_boxed0) -> {
        Part _source0 = ((Part)(java.lang.Object)(_source0_boxed0));
        if (_source0.is_Unclassified()) {
          dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h0 = ((agency.open.ahp.Terminal.Part_Unclassified)_source0)._value;
          return _3_p;
        } else {
          dafny.DafnySequence<? extends dafny.CodePoint> _5___mcc_h1 = ((agency.open.ahp.Terminal.Part_Command)_source0)._commandId;
          dafny.DafnySequence<? extends dafny.CodePoint> _6___mcc_h2 = ((agency.open.ahp.Terminal.Part_Command)_source0)._commandLine;
          dafny.DafnySequence<? extends dafny.CodePoint> _7___mcc_h3 = ((agency.open.ahp.Terminal.Part_Command)_source0)._output;
          java.math.BigInteger _8___mcc_h4 = ((agency.open.ahp.Terminal.Part_Command)_source0)._timestamp;
          boolean _9___mcc_h5 = ((agency.open.ahp.Terminal.Part_Command)_source0)._isComplete;
          agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _10___mcc_h6 = ((agency.open.ahp.Terminal.Part_Command)_source0)._exitCode;
          agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _11___mcc_h7 = ((agency.open.ahp.Terminal.Part_Command)_source0)._durationMs;
          agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _12_dm = _11___mcc_h7;
          agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _13_ec = _10___mcc_h6;
          boolean _14_comp = _9___mcc_h5;
          java.math.BigInteger _15_ts = _8___mcc_h4;
          dafny.DafnySequence<? extends dafny.CodePoint> _16_out = _7___mcc_h3;
          dafny.DafnySequence<? extends dafny.CodePoint> _17_cl = _6___mcc_h2;
          dafny.DafnySequence<? extends dafny.CodePoint> _18_cid = _5___mcc_h1;
          if ((_18_cid).equals(_0_id)) {
            return Part.create_Command(_18_cid, _17_cl, _16_out, _15_ts, true, agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, _1_code), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, _2_dur));
          } else {
            return _3_p;
          }
        }
      }).apply(_3_p);
    });}).apply(id, code, dur), content);
  }
  public static dafny.Tuple2<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome> ApplyToTerminal(TerminalState s, TerminalAction a, java.math.BigInteger now)
  {
    TerminalState _pat_let_tv0 = s;
    TerminalState _pat_let_tv1 = s;
    TerminalState _pat_let_tv2 = s;
    TerminalAction _source0 = a;
    if (_source0.is_TCwdChanged()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _0___mcc_h0 = ((agency.open.ahp.Terminal.TerminalAction_TCwdChanged)_source0)._cwd;
      dafny.DafnySequence<? extends dafny.CodePoint> _1_c = _0___mcc_h0;
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(s, boxed116 -> {
  TerminalState _pat_let58_0 = ((TerminalState)(java.lang.Object)(boxed116));
  return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let58_0, boxed117 -> {
    TerminalState _2_dt__update__tmp_h0 = ((TerminalState)(java.lang.Object)(boxed117));
    return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), _1_c), boxed118 -> {
      agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _pat_let59_0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed118));
      return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>, TerminalState>Let(_pat_let59_0, boxed119 -> {
        agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>> _3_dt__update_hcwd_h0 = ((agency.open.ahp.AhpSkeleton.Option<dafny.DafnySequence<? extends dafny.CodePoint>>)(java.lang.Object)(boxed119));
        return TerminalState.create((_2_dt__update__tmp_h0).dtor_title(), _3_dt__update_hcwd_h0, (_2_dt__update__tmp_h0).dtor_cols(), (_2_dt__update__tmp_h0).dtor_rows(), (_2_dt__update__tmp_h0).dtor_content(), (_2_dt__update__tmp_h0).dtor_claim(), (_2_dt__update__tmp_h0).dtor_exitCode(), (_2_dt__update__tmp_h0).dtor_supportsCommandDetection());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_TTitleChanged()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _4___mcc_h1 = ((agency.open.ahp.Terminal.TerminalAction_TTitleChanged)_source0)._title;
      dafny.DafnySequence<? extends dafny.CodePoint> _5_t = _4___mcc_h1;
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(s, boxed120 -> {
  TerminalState _pat_let60_0 = ((TerminalState)(java.lang.Object)(boxed120));
  return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let60_0, boxed121 -> {
    TerminalState _6_dt__update__tmp_h1 = ((TerminalState)(java.lang.Object)(boxed121));
    return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, TerminalState>Let(_5_t, boxed122 -> {
      dafny.DafnySequence<? extends dafny.CodePoint> _pat_let61_0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed122));
      return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends dafny.CodePoint>, TerminalState>Let(_pat_let61_0, boxed123 -> {
        dafny.DafnySequence<? extends dafny.CodePoint> _7_dt__update_htitle_h0 = ((dafny.DafnySequence<? extends dafny.CodePoint>)(java.lang.Object)(boxed123));
        return TerminalState.create(_7_dt__update_htitle_h0, (_6_dt__update__tmp_h1).dtor_cwd(), (_6_dt__update__tmp_h1).dtor_cols(), (_6_dt__update__tmp_h1).dtor_rows(), (_6_dt__update__tmp_h1).dtor_content(), (_6_dt__update__tmp_h1).dtor_claim(), (_6_dt__update__tmp_h1).dtor_exitCode(), (_6_dt__update__tmp_h1).dtor_supportsCommandDetection());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_TResized()) {
      java.math.BigInteger _8___mcc_h2 = ((agency.open.ahp.Terminal.TerminalAction_TResized)_source0)._cols;
      java.math.BigInteger _9___mcc_h3 = ((agency.open.ahp.Terminal.TerminalAction_TResized)_source0)._rows;
      java.math.BigInteger _10_ro = _9___mcc_h3;
      java.math.BigInteger _11_co = _8___mcc_h2;
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(s, boxed124 -> {
  TerminalState _pat_let62_0 = ((TerminalState)(java.lang.Object)(boxed124));
  return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let62_0, boxed125 -> {
    TerminalState _12_dt__update__tmp_h2 = ((TerminalState)(java.lang.Object)(boxed125));
    return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, _10_ro), boxed126 -> {
      agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _pat_let63_0 = ((agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>)(java.lang.Object)(boxed126));
      return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>, TerminalState>Let(_pat_let63_0, boxed127 -> {
        agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _13_dt__update_hrows_h0 = ((agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>)(java.lang.Object)(boxed127));
        return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, _11_co), boxed128 -> {
          agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _pat_let64_0 = ((agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>)(java.lang.Object)(boxed128));
          return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>, TerminalState>Let(_pat_let64_0, boxed129 -> {
            agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _14_dt__update_hcols_h0 = ((agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>)(java.lang.Object)(boxed129));
            return TerminalState.create((_12_dt__update__tmp_h2).dtor_title(), (_12_dt__update__tmp_h2).dtor_cwd(), _14_dt__update_hcols_h0, _13_dt__update_hrows_h0, (_12_dt__update__tmp_h2).dtor_content(), (_12_dt__update__tmp_h2).dtor_claim(), (_12_dt__update__tmp_h2).dtor_exitCode(), (_12_dt__update__tmp_h2).dtor_supportsCommandDetection());
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
    } else if (_source0.is_TExited()) {
      java.math.BigInteger _15___mcc_h4 = ((agency.open.ahp.Terminal.TerminalAction_TExited)_source0)._code;
      java.math.BigInteger _16_code = _15___mcc_h4;
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(s, boxed130 -> {
  TerminalState _pat_let65_0 = ((TerminalState)(java.lang.Object)(boxed130));
  return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let65_0, boxed131 -> {
    TerminalState _17_dt__update__tmp_h3 = ((TerminalState)(java.lang.Object)(boxed131));
    return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, _16_code), boxed132 -> {
      agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _pat_let66_0 = ((agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>)(java.lang.Object)(boxed132));
      return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>, TerminalState>Let(_pat_let66_0, boxed133 -> {
        agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _18_dt__update_hexitCode_h0 = ((agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>)(java.lang.Object)(boxed133));
        return TerminalState.create((_17_dt__update__tmp_h3).dtor_title(), (_17_dt__update__tmp_h3).dtor_cwd(), (_17_dt__update__tmp_h3).dtor_cols(), (_17_dt__update__tmp_h3).dtor_rows(), (_17_dt__update__tmp_h3).dtor_content(), (_17_dt__update__tmp_h3).dtor_claim(), _18_dt__update_hexitCode_h0, (_17_dt__update__tmp_h3).dtor_supportsCommandDetection());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_TData()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _19___mcc_h5 = ((agency.open.ahp.Terminal.TerminalAction_TData)_source0)._data;
      dafny.DafnySequence<? extends dafny.CodePoint> _20_d = _19___mcc_h5;
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(s, boxed134 -> {
  TerminalState _pat_let67_0 = ((TerminalState)(java.lang.Object)(boxed134));
  return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let67_0, boxed135 -> {
    TerminalState _21_dt__update__tmp_h4 = ((TerminalState)(java.lang.Object)(boxed135));
    return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(__default.appendData((_pat_let_tv0).dtor_content(), _20_d), boxed136 -> {
      dafny.DafnySequence<? extends Part> _pat_let68_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed136));
      return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(_pat_let68_0, boxed137 -> {
        dafny.DafnySequence<? extends Part> _22_dt__update_hcontent_h0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed137));
        return TerminalState.create((_21_dt__update__tmp_h4).dtor_title(), (_21_dt__update__tmp_h4).dtor_cwd(), (_21_dt__update__tmp_h4).dtor_cols(), (_21_dt__update__tmp_h4).dtor_rows(), _22_dt__update_hcontent_h0, (_21_dt__update__tmp_h4).dtor_claim(), (_21_dt__update__tmp_h4).dtor_exitCode(), (_21_dt__update__tmp_h4).dtor_supportsCommandDetection());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_TCleared()) {
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(s, boxed138 -> {
  TerminalState _pat_let69_0 = ((TerminalState)(java.lang.Object)(boxed138));
  return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let69_0, boxed139 -> {
    TerminalState _23_dt__update__tmp_h5 = ((TerminalState)(java.lang.Object)(boxed139));
    return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), boxed140 -> {
      dafny.DafnySequence<? extends Part> _pat_let70_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed140));
      return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(_pat_let70_0, boxed141 -> {
        dafny.DafnySequence<? extends Part> _24_dt__update_hcontent_h1 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed141));
        return TerminalState.create((_23_dt__update__tmp_h5).dtor_title(), (_23_dt__update__tmp_h5).dtor_cwd(), (_23_dt__update__tmp_h5).dtor_cols(), (_23_dt__update__tmp_h5).dtor_rows(), _24_dt__update_hcontent_h1, (_23_dt__update__tmp_h5).dtor_claim(), (_23_dt__update__tmp_h5).dtor_exitCode(), (_23_dt__update__tmp_h5).dtor_supportsCommandDetection());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_TClaimed()) {
      agency.open.ahp.ConfluxCodec.Json _25___mcc_h6 = ((agency.open.ahp.Terminal.TerminalAction_TClaimed)_source0)._claim;
      agency.open.ahp.ConfluxCodec.Json _26_c = _25___mcc_h6;
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(s, boxed142 -> {
  TerminalState _pat_let71_0 = ((TerminalState)(java.lang.Object)(boxed142));
  return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let71_0, boxed143 -> {
    TerminalState _27_dt__update__tmp_h6 = ((TerminalState)(java.lang.Object)(boxed143));
    return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _26_c), boxed144 -> {
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let72_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed144));
      return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, TerminalState>Let(_pat_let72_0, boxed145 -> {
        agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _28_dt__update_hclaim_h0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed145));
        return TerminalState.create((_27_dt__update__tmp_h6).dtor_title(), (_27_dt__update__tmp_h6).dtor_cwd(), (_27_dt__update__tmp_h6).dtor_cols(), (_27_dt__update__tmp_h6).dtor_rows(), (_27_dt__update__tmp_h6).dtor_content(), _28_dt__update_hclaim_h0, (_27_dt__update__tmp_h6).dtor_exitCode(), (_27_dt__update__tmp_h6).dtor_supportsCommandDetection());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_TCommandDetectionAvailable()) {
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(s, boxed146 -> {
  TerminalState _pat_let73_0 = ((TerminalState)(java.lang.Object)(boxed146));
  return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let73_0, boxed147 -> {
    TerminalState _29_dt__update__tmp_h7 = ((TerminalState)(java.lang.Object)(boxed147));
    return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), boxed148 -> {
      agency.open.ahp.AhpSkeleton.Option<Boolean> _pat_let74_0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed148));
      return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, TerminalState>Let(_pat_let74_0, boxed149 -> {
        agency.open.ahp.AhpSkeleton.Option<Boolean> _30_dt__update_hsupportsCommandDetection_h0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed149));
        return TerminalState.create((_29_dt__update__tmp_h7).dtor_title(), (_29_dt__update__tmp_h7).dtor_cwd(), (_29_dt__update__tmp_h7).dtor_cols(), (_29_dt__update__tmp_h7).dtor_rows(), (_29_dt__update__tmp_h7).dtor_content(), (_29_dt__update__tmp_h7).dtor_claim(), (_29_dt__update__tmp_h7).dtor_exitCode(), _30_dt__update_hsupportsCommandDetection_h0);
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_TCommandExecuted()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _31___mcc_h7 = ((agency.open.ahp.Terminal.TerminalAction_TCommandExecuted)_source0)._commandId;
      dafny.DafnySequence<? extends dafny.CodePoint> _32___mcc_h8 = ((agency.open.ahp.Terminal.TerminalAction_TCommandExecuted)_source0)._commandLine;
      java.math.BigInteger _33___mcc_h9 = ((agency.open.ahp.Terminal.TerminalAction_TCommandExecuted)_source0)._timestamp;
      java.math.BigInteger _34_ts = _33___mcc_h9;
      dafny.DafnySequence<? extends dafny.CodePoint> _35_cl = _32___mcc_h8;
      dafny.DafnySequence<? extends dafny.CodePoint> _36_id = _31___mcc_h7;
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(s, boxed150 -> {
  TerminalState _pat_let75_0 = ((TerminalState)(java.lang.Object)(boxed150));
  return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let75_0, boxed151 -> {
    TerminalState _37_dt__update__tmp_h8 = ((TerminalState)(java.lang.Object)(boxed151));
    return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), boxed152 -> {
      agency.open.ahp.AhpSkeleton.Option<Boolean> _pat_let76_0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed152));
      return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, TerminalState>Let(_pat_let76_0, boxed153 -> {
        agency.open.ahp.AhpSkeleton.Option<Boolean> _38_dt__update_hsupportsCommandDetection_h1 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed153));
        return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(dafny.DafnySequence.<Part>concatenate((_pat_let_tv1).dtor_content(), dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Command(_36_id, _35_cl, dafny.DafnySequence.asUnicodeString(""), _34_ts, false, agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER)))), boxed154 -> {
          dafny.DafnySequence<? extends Part> _pat_let77_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed154));
          return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(_pat_let77_0, boxed155 -> {
            dafny.DafnySequence<? extends Part> _39_dt__update_hcontent_h2 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed155));
            return TerminalState.create((_37_dt__update__tmp_h8).dtor_title(), (_37_dt__update__tmp_h8).dtor_cwd(), (_37_dt__update__tmp_h8).dtor_cols(), (_37_dt__update__tmp_h8).dtor_rows(), _39_dt__update_hcontent_h2, (_37_dt__update__tmp_h8).dtor_claim(), (_37_dt__update__tmp_h8).dtor_exitCode(), _38_dt__update_hsupportsCommandDetection_h1);
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
    } else if (_source0.is_TCommandFinished()) {
      dafny.DafnySequence<? extends dafny.CodePoint> _40___mcc_h10 = ((agency.open.ahp.Terminal.TerminalAction_TCommandFinished)_source0)._commandId;
      java.math.BigInteger _41___mcc_h11 = ((agency.open.ahp.Terminal.TerminalAction_TCommandFinished)_source0)._code;
      java.math.BigInteger _42___mcc_h12 = ((agency.open.ahp.Terminal.TerminalAction_TCommandFinished)_source0)._durationMs;
      java.math.BigInteger _43_dur = _42___mcc_h12;
      java.math.BigInteger _44_code = _41___mcc_h11;
      dafny.DafnySequence<? extends dafny.CodePoint> _45_id = _40___mcc_h10;
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(s, boxed156 -> {
  TerminalState _pat_let78_0 = ((TerminalState)(java.lang.Object)(boxed156));
  return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let78_0, boxed157 -> {
    TerminalState _46_dt__update__tmp_h9 = ((TerminalState)(java.lang.Object)(boxed157));
    return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(__default.finishCommand((_pat_let_tv2).dtor_content(), _45_id, _44_code, _43_dur), boxed158 -> {
      dafny.DafnySequence<? extends Part> _pat_let79_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed158));
      return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(_pat_let79_0, boxed159 -> {
        dafny.DafnySequence<? extends Part> _47_dt__update_hcontent_h3 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed159));
        return TerminalState.create((_46_dt__update__tmp_h9).dtor_title(), (_46_dt__update__tmp_h9).dtor_cwd(), (_46_dt__update__tmp_h9).dtor_cols(), (_46_dt__update__tmp_h9).dtor_rows(), _47_dt__update_hcontent_h3, (_46_dt__update__tmp_h9).dtor_claim(), (_46_dt__update__tmp_h9).dtor_exitCode(), (_46_dt__update__tmp_h9).dtor_supportsCommandDetection());
      }
      )));
    }
    )));
  }
  )));
}
))), agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else if (_source0.is_TInput()) {
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
    } else {
      agency.open.ahp.ConfluxCodec.Json _48___mcc_h13 = ((agency.open.ahp.Terminal.TerminalAction_TUnknown)_source0)._raw;
      return dafny.Tuple2.<TerminalState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
    }
  }
  public static TerminalState apply1(TerminalState s, TerminalAction a)
  {
    return (__default.ApplyToTerminal(s, a, java.math.BigInteger.valueOf(9999L))).dtor__0();
  }
  public static TerminalState fold(TerminalState s, dafny.DafnySequence<? extends TerminalAction> acts)
  {
    return agency.open.ahp.ConfluxContract.__default.<TerminalState, TerminalAction>Fold(TerminalState._typeDescriptor(), TerminalAction._typeDescriptor(), __default::apply1, s, acts);
  }
  public static TerminalState T0() {
    return TerminalState.create(dafny.DafnySequence.asUnicodeString("bash"), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_None(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR)), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_None(agency.open.ahp.ConfluxCodec.Json._typeDescriptor()), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_None(dafny.TypeDescriptor.BOOLEAN));
  }
  public static agency.open.ahp.ConfluxCodec.Json CL() {
    return agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("client"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("clientId"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("c1"))) }));
  }
  public static dafny.Tuple2 RunCorpus()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    java.math.BigInteger total = java.math.BigInteger.ZERO;
    if(true) {
      total = java.math.BigInteger.valueOf(19L);
      pass = java.math.BigInteger.ZERO;
      TerminalState _0_base;
      TerminalState _1_dt__update__tmp_h0 = __default.T0();
      agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _2_dt__update_hclaim_h0 = agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), __default.CL());
      _0_base = TerminalState.create((_1_dt__update__tmp_h0).dtor_title(), (_1_dt__update__tmp_h0).dtor_cwd(), (_1_dt__update__tmp_h0).dtor_cols(), (_1_dt__update__tmp_h0).dtor_rows(), (_1_dt__update__tmp_h0).dtor_content(), _2_dt__update_hclaim_h0, (_1_dt__update__tmp_h0).dtor_exitCode(), (_1_dt__update__tmp_h0).dtor_supportsCommandDetection());
      if (java.util.Objects.equals((__default.apply1(_0_base, TerminalAction.create_TCwdChanged(dafny.DafnySequence.asUnicodeString("/tmp")))).dtor_cwd(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("/tmp")))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (((__default.apply1(_0_base, TerminalAction.create_TTitleChanged(dafny.DafnySequence.asUnicodeString("zsh")))).dtor_title()).equals(dafny.DafnySequence.asUnicodeString("zsh"))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_0_base, TerminalAction.create_TResized(java.math.BigInteger.valueOf(80L), java.math.BigInteger.valueOf(24L))), ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_0_base, boxed160 -> {
        TerminalState _pat_let80_0 = ((TerminalState)(java.lang.Object)(boxed160));
        return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let80_0, boxed161 -> {
          TerminalState _3_dt__update__tmp_h1 = ((TerminalState)(java.lang.Object)(boxed161));
          return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.valueOf(24L)), boxed162 -> {
            agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _pat_let81_0 = ((agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>)(java.lang.Object)(boxed162));
            return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>, TerminalState>Let(_pat_let81_0, boxed163 -> {
              agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _4_dt__update_hrows_h0 = ((agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>)(java.lang.Object)(boxed163));
              return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.valueOf(80L)), boxed164 -> {
                agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _pat_let82_0 = ((agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>)(java.lang.Object)(boxed164));
                return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>, TerminalState>Let(_pat_let82_0, boxed165 -> {
                  agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger> _5_dt__update_hcols_h0 = ((agency.open.ahp.AhpSkeleton.Option<java.math.BigInteger>)(java.lang.Object)(boxed165));
                  return TerminalState.create((_3_dt__update__tmp_h1).dtor_title(), (_3_dt__update__tmp_h1).dtor_cwd(), _5_dt__update_hcols_h0, _4_dt__update_hrows_h0, (_3_dt__update__tmp_h1).dtor_content(), (_3_dt__update__tmp_h1).dtor_claim(), (_3_dt__update__tmp_h1).dtor_exitCode(), (_3_dt__update__tmp_h1).dtor_supportsCommandDetection());
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
      if (java.util.Objects.equals((__default.apply1(_0_base, TerminalAction.create_TExited(java.math.BigInteger.ZERO))).dtor_exitCode(), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.ZERO))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_0_base, TerminalAction.create_TInput()), _0_base)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(_0_base, TerminalAction.create_TUnknown(agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("type"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("terminal/nope"))) })))), _0_base)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (((__default.apply1(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_0_base, boxed166 -> {
        TerminalState _pat_let83_0 = ((TerminalState)(java.lang.Object)(boxed166));
        return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let83_0, boxed167 -> {
          TerminalState _6_dt__update__tmp_h2 = ((TerminalState)(java.lang.Object)(boxed167));
          return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Unclassified(dafny.DafnySequence.asUnicodeString("x"))), boxed168 -> {
            dafny.DafnySequence<? extends Part> _pat_let84_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed168));
            return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(_pat_let84_0, boxed169 -> {
              dafny.DafnySequence<? extends Part> _7_dt__update_hcontent_h0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed169));
              return TerminalState.create((_6_dt__update__tmp_h2).dtor_title(), (_6_dt__update__tmp_h2).dtor_cwd(), (_6_dt__update__tmp_h2).dtor_cols(), (_6_dt__update__tmp_h2).dtor_rows(), _7_dt__update_hcontent_h0, (_6_dt__update__tmp_h2).dtor_claim(), (_6_dt__update__tmp_h2).dtor_exitCode(), (_6_dt__update__tmp_h2).dtor_supportsCommandDetection());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), TerminalAction.create_TCleared())).dtor_content()).equals(dafny.DafnySequence.<Part> empty(Part._typeDescriptor()))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals(__default.apply1(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_0_base, boxed170 -> {
        TerminalState _pat_let85_0 = ((TerminalState)(java.lang.Object)(boxed170));
        return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let85_0, boxed171 -> {
          TerminalState _8_dt__update__tmp_h3 = ((TerminalState)(java.lang.Object)(boxed171));
          return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), boxed172 -> {
            agency.open.ahp.AhpSkeleton.Option<Boolean> _pat_let86_0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed172));
            return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, TerminalState>Let(_pat_let86_0, boxed173 -> {
              agency.open.ahp.AhpSkeleton.Option<Boolean> _9_dt__update_hsupportsCommandDetection_h0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed173));
              return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Unclassified(dafny.DafnySequence.asUnicodeString("x"))), boxed174 -> {
                dafny.DafnySequence<? extends Part> _pat_let87_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed174));
                return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(_pat_let87_0, boxed175 -> {
                  dafny.DafnySequence<? extends Part> _10_dt__update_hcontent_h1 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed175));
                  return TerminalState.create((_8_dt__update__tmp_h3).dtor_title(), (_8_dt__update__tmp_h3).dtor_cwd(), (_8_dt__update__tmp_h3).dtor_cols(), (_8_dt__update__tmp_h3).dtor_rows(), _10_dt__update_hcontent_h1, (_8_dt__update__tmp_h3).dtor_claim(), (_8_dt__update__tmp_h3).dtor_exitCode(), _9_dt__update_hsupportsCommandDetection_h0);
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
      ))), TerminalAction.create_TCleared()), ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_0_base, boxed176 -> {
        TerminalState _pat_let88_0 = ((TerminalState)(java.lang.Object)(boxed176));
        return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let88_0, boxed177 -> {
          TerminalState _11_dt__update__tmp_h4 = ((TerminalState)(java.lang.Object)(boxed177));
          return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), boxed178 -> {
            agency.open.ahp.AhpSkeleton.Option<Boolean> _pat_let89_0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed178));
            return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, TerminalState>Let(_pat_let89_0, boxed179 -> {
              agency.open.ahp.AhpSkeleton.Option<Boolean> _12_dt__update_hsupportsCommandDetection_h1 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed179));
              return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), boxed180 -> {
                dafny.DafnySequence<? extends Part> _pat_let90_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed180));
                return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(_pat_let90_0, boxed181 -> {
                  dafny.DafnySequence<? extends Part> _13_dt__update_hcontent_h2 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed181));
                  return TerminalState.create((_11_dt__update__tmp_h4).dtor_title(), (_11_dt__update__tmp_h4).dtor_cwd(), (_11_dt__update__tmp_h4).dtor_cols(), (_11_dt__update__tmp_h4).dtor_rows(), _13_dt__update_hcontent_h2, (_11_dt__update__tmp_h4).dtor_claim(), (_11_dt__update__tmp_h4).dtor_exitCode(), _12_dt__update_hsupportsCommandDetection_h1);
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
      agency.open.ahp.ConfluxCodec.Json _14_sc;
      _14_sc = agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("session"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("session"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("s1"))) }));
      if (java.util.Objects.equals((__default.apply1(_0_base, TerminalAction.create_TClaimed(_14_sc))).dtor_claim(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _14_sc))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      agency.open.ahp.ConfluxCodec.Json _15_scTool;
      _15_scTool = agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("kind"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("session"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("session"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("s1"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("toolCallId"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("t1"))) }));
      if (java.util.Objects.equals((__default.apply1(_0_base, TerminalAction.create_TClaimed(_15_scTool))).dtor_claim(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _15_scTool))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      agency.open.ahp.ConfluxCodec.Json _pat_let_tv0 = _14_sc;
      if (java.util.Objects.equals((__default.apply1(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_0_base, boxed182 -> {
        TerminalState _pat_let91_0 = ((TerminalState)(java.lang.Object)(boxed182));
        return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let91_0, boxed183 -> {
          TerminalState _16_dt__update__tmp_h5 = ((TerminalState)(java.lang.Object)(boxed183));
          return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _pat_let_tv0), boxed184 -> {
            agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _pat_let92_0 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed184));
            return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>, TerminalState>Let(_pat_let92_0, boxed185 -> {
              agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json> _17_dt__update_hclaim_h1 = ((agency.open.ahp.AhpSkeleton.Option<agency.open.ahp.ConfluxCodec.Json>)(java.lang.Object)(boxed185));
              return TerminalState.create((_16_dt__update__tmp_h5).dtor_title(), (_16_dt__update__tmp_h5).dtor_cwd(), (_16_dt__update__tmp_h5).dtor_cols(), (_16_dt__update__tmp_h5).dtor_rows(), (_16_dt__update__tmp_h5).dtor_content(), _17_dt__update_hclaim_h1, (_16_dt__update__tmp_h5).dtor_exitCode(), (_16_dt__update__tmp_h5).dtor_supportsCommandDetection());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), TerminalAction.create_TClaimed(__default.CL()))).dtor_claim(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), __default.CL()))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (java.util.Objects.equals((__default.apply1(_0_base, TerminalAction.create_TCommandDetectionAvailable())).dtor_supportsCommandDetection(), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      TerminalState _18_afterExec;
      _18_afterExec = __default.apply1(_0_base, TerminalAction.create_TCommandExecuted(dafny.DafnySequence.asUnicodeString("cmd-1"), dafny.DafnySequence.asUnicodeString("npm test"), java.math.BigInteger.valueOf(1700000000000L)));
      if (java.util.Objects.equals(_18_afterExec, ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_0_base, boxed186 -> {
        TerminalState _pat_let93_0 = ((TerminalState)(java.lang.Object)(boxed186));
        return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let93_0, boxed187 -> {
          TerminalState _19_dt__update__tmp_h6 = ((TerminalState)(java.lang.Object)(boxed187));
          return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, TerminalState>Let(agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true), boxed188 -> {
            agency.open.ahp.AhpSkeleton.Option<Boolean> _pat_let94_0 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed188));
            return ((TerminalState)(java.lang.Object)(dafny.Helpers.<agency.open.ahp.AhpSkeleton.Option<Boolean>, TerminalState>Let(_pat_let94_0, boxed189 -> {
              agency.open.ahp.AhpSkeleton.Option<Boolean> _20_dt__update_hsupportsCommandDetection_h2 = ((agency.open.ahp.AhpSkeleton.Option<Boolean>)(java.lang.Object)(boxed189));
              return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Command(dafny.DafnySequence.asUnicodeString("cmd-1"), dafny.DafnySequence.asUnicodeString("npm test"), dafny.DafnySequence.asUnicodeString(""), java.math.BigInteger.valueOf(1700000000000L), false, agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER))), boxed190 -> {
                dafny.DafnySequence<? extends Part> _pat_let95_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed190));
                return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(_pat_let95_0, boxed191 -> {
                  dafny.DafnySequence<? extends Part> _21_dt__update_hcontent_h3 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed191));
                  return TerminalState.create((_19_dt__update__tmp_h6).dtor_title(), (_19_dt__update__tmp_h6).dtor_cwd(), (_19_dt__update__tmp_h6).dtor_cols(), (_19_dt__update__tmp_h6).dtor_rows(), _21_dt__update_hcontent_h3, (_19_dt__update__tmp_h6).dtor_claim(), (_19_dt__update__tmp_h6).dtor_exitCode(), _20_dt__update_hsupportsCommandDetection_h2);
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
      TerminalState _22_withCmd;
      TerminalState _23_dt__update__tmp_h7 = _0_base;
      agency.open.ahp.AhpSkeleton.Option<Boolean> _24_dt__update_hsupportsCommandDetection_h3 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true);
      dafny.DafnySequence<? extends Part> _25_dt__update_hcontent_h4 = dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Command(dafny.DafnySequence.asUnicodeString("cmd-1"), dafny.DafnySequence.asUnicodeString("npm test"), dafny.DafnySequence.asUnicodeString("All tests passed\r\n"), java.math.BigInteger.valueOf(1700000000000L), false, agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER)));
      _22_withCmd = TerminalState.create((_23_dt__update__tmp_h7).dtor_title(), (_23_dt__update__tmp_h7).dtor_cwd(), (_23_dt__update__tmp_h7).dtor_cols(), (_23_dt__update__tmp_h7).dtor_rows(), _25_dt__update_hcontent_h4, (_23_dt__update__tmp_h7).dtor_claim(), (_23_dt__update__tmp_h7).dtor_exitCode(), _24_dt__update_hsupportsCommandDetection_h3);
      TerminalState _26_doneCmd;
      TerminalState _27_dt__update__tmp_h8 = _0_base;
      agency.open.ahp.AhpSkeleton.Option<Boolean> _28_dt__update_hsupportsCommandDetection_h4 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true);
      dafny.DafnySequence<? extends Part> _29_dt__update_hcontent_h5 = dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Command(dafny.DafnySequence.asUnicodeString("cmd-1"), dafny.DafnySequence.asUnicodeString("npm test"), dafny.DafnySequence.asUnicodeString("All tests passed\r\n"), java.math.BigInteger.valueOf(1700000000000L), true, agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.ZERO), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.valueOf(1234L))));
      _26_doneCmd = TerminalState.create((_27_dt__update__tmp_h8).dtor_title(), (_27_dt__update__tmp_h8).dtor_cwd(), (_27_dt__update__tmp_h8).dtor_cols(), (_27_dt__update__tmp_h8).dtor_rows(), _29_dt__update_hcontent_h5, (_27_dt__update__tmp_h8).dtor_claim(), (_27_dt__update__tmp_h8).dtor_exitCode(), _28_dt__update_hsupportsCommandDetection_h4);
      if (java.util.Objects.equals(__default.apply1(_22_withCmd, TerminalAction.create_TCommandFinished(dafny.DafnySequence.asUnicodeString("cmd-1"), java.math.BigInteger.ZERO, java.math.BigInteger.valueOf(1234L))), _26_doneCmd)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (((__default.apply1(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_0_base, boxed192 -> {
        TerminalState _pat_let96_0 = ((TerminalState)(java.lang.Object)(boxed192));
        return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let96_0, boxed193 -> {
          TerminalState _30_dt__update__tmp_h9 = ((TerminalState)(java.lang.Object)(boxed193));
          return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Unclassified(dafny.DafnySequence.asUnicodeString("a"))), boxed194 -> {
            dafny.DafnySequence<? extends Part> _pat_let97_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed194));
            return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(_pat_let97_0, boxed195 -> {
              dafny.DafnySequence<? extends Part> _31_dt__update_hcontent_h6 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed195));
              return TerminalState.create((_30_dt__update__tmp_h9).dtor_title(), (_30_dt__update__tmp_h9).dtor_cwd(), (_30_dt__update__tmp_h9).dtor_cols(), (_30_dt__update__tmp_h9).dtor_rows(), _31_dt__update_hcontent_h6, (_30_dt__update__tmp_h9).dtor_claim(), (_30_dt__update__tmp_h9).dtor_exitCode(), (_30_dt__update__tmp_h9).dtor_supportsCommandDetection());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), TerminalAction.create_TData(dafny.DafnySequence.asUnicodeString("b")))).dtor_content()).equals(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Unclassified(dafny.DafnySequence.asUnicodeString("ab"))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      TerminalState _32_complete;
      TerminalState _33_dt__update__tmp_h10 = _0_base;
      agency.open.ahp.AhpSkeleton.Option<Boolean> _34_dt__update_hsupportsCommandDetection_h5 = agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true);
      dafny.DafnySequence<? extends Part> _35_dt__update_hcontent_h7 = dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Command(dafny.DafnySequence.asUnicodeString("cmd-1"), dafny.DafnySequence.asUnicodeString("echo hi"), dafny.DafnySequence.asUnicodeString("hi\r\n"), java.math.BigInteger.valueOf(1700000000000L), true, agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.ZERO), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.valueOf(50L))));
      _32_complete = TerminalState.create((_33_dt__update__tmp_h10).dtor_title(), (_33_dt__update__tmp_h10).dtor_cwd(), (_33_dt__update__tmp_h10).dtor_cols(), (_33_dt__update__tmp_h10).dtor_rows(), _35_dt__update_hcontent_h7, (_33_dt__update__tmp_h10).dtor_claim(), (_33_dt__update__tmp_h10).dtor_exitCode(), _34_dt__update_hsupportsCommandDetection_h5);
      if (((__default.apply1(_32_complete, TerminalAction.create_TData(dafny.DafnySequence.asUnicodeString("$ ")))).dtor_content()).equals(dafny.DafnySequence.<Part>concatenate((_32_complete).dtor_content(), dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Unclassified(dafny.DafnySequence.asUnicodeString("$ ")))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      if (((__default.apply1(_22_withCmd, TerminalAction.create_TData(dafny.DafnySequence.asUnicodeString("!")))).dtor_content()).equals(dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Command(dafny.DafnySequence.asUnicodeString("cmd-1"), dafny.DafnySequence.asUnicodeString("npm test"), dafny.DafnySequence.asUnicodeString("All tests passed\r\n!"), java.math.BigInteger.valueOf(1700000000000L), false, agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_None(dafny.TypeDescriptor.BIG_INTEGER))))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      TerminalState _36_life;
      _36_life = __default.fold(_0_base, dafny.DafnySequence.<TerminalAction> of(TerminalAction._typeDescriptor(), TerminalAction.create_TCwdChanged(dafny.DafnySequence.asUnicodeString("/w")), TerminalAction.create_TData(dafny.DafnySequence.asUnicodeString("x")), TerminalAction.create_TResized(java.math.BigInteger.valueOf(100L), java.math.BigInteger.valueOf(40L)), TerminalAction.create_TClaimed(_14_sc), TerminalAction.create_TExited(java.math.BigInteger.ONE)));
      if ((((java.util.Objects.equals((_36_life).dtor_cwd(), agency.open.ahp.AhpSkeleton.Option.<dafny.DafnySequence<? extends dafny.CodePoint>>create_Some(dafny.DafnySequence.<dafny.CodePoint>_typeDescriptor(dafny.TypeDescriptor.UNICODE_CHAR), dafny.DafnySequence.asUnicodeString("/w")))) && (java.util.Objects.equals((_36_life).dtor_cols(), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.valueOf(100L))))) && (java.util.Objects.equals((_36_life).dtor_exitCode(), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.ONE)))) && (java.util.Objects.equals((_36_life).dtor_claim(), agency.open.ahp.AhpSkeleton.Option.<agency.open.ahp.ConfluxCodec.Json>create_Some(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), _14_sc)))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      TerminalState _37_lc;
      _37_lc = __default.fold(((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_0_base, boxed196 -> {
        TerminalState _pat_let98_0 = ((TerminalState)(java.lang.Object)(boxed196));
        return ((TerminalState)(java.lang.Object)(dafny.Helpers.<TerminalState, TerminalState>Let(_pat_let98_0, boxed197 -> {
          TerminalState _38_dt__update__tmp_h11 = ((TerminalState)(java.lang.Object)(boxed197));
          return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(dafny.DafnySequence.<Part> empty(Part._typeDescriptor()), boxed198 -> {
            dafny.DafnySequence<? extends Part> _pat_let99_0 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed198));
            return ((TerminalState)(java.lang.Object)(dafny.Helpers.<dafny.DafnySequence<? extends Part>, TerminalState>Let(_pat_let99_0, boxed199 -> {
              dafny.DafnySequence<? extends Part> _39_dt__update_hcontent_h8 = ((dafny.DafnySequence<? extends Part>)(java.lang.Object)(boxed199));
              return TerminalState.create((_38_dt__update__tmp_h11).dtor_title(), (_38_dt__update__tmp_h11).dtor_cwd(), (_38_dt__update__tmp_h11).dtor_cols(), (_38_dt__update__tmp_h11).dtor_rows(), _39_dt__update_hcontent_h8, (_38_dt__update__tmp_h11).dtor_claim(), (_38_dt__update__tmp_h11).dtor_exitCode(), (_38_dt__update__tmp_h11).dtor_supportsCommandDetection());
            }
            )));
          }
          )));
        }
        )));
      }
      ))), dafny.DafnySequence.<TerminalAction> of(TerminalAction._typeDescriptor(), TerminalAction.create_TData(dafny.DafnySequence.asUnicodeString("$ ")), TerminalAction.create_TCommandExecuted(dafny.DafnySequence.asUnicodeString("cmd-1"), dafny.DafnySequence.asUnicodeString("echo hello"), java.math.BigInteger.valueOf(1700000000000L)), TerminalAction.create_TData(dafny.DafnySequence.asUnicodeString("hello\r\n")), TerminalAction.create_TCommandFinished(dafny.DafnySequence.asUnicodeString("cmd-1"), java.math.BigInteger.ZERO, java.math.BigInteger.valueOf(10L)), TerminalAction.create_TData(dafny.DafnySequence.asUnicodeString("$ ")), TerminalAction.create_TCommandExecuted(dafny.DafnySequence.asUnicodeString("cmd-2"), dafny.DafnySequence.asUnicodeString("exit 1"), java.math.BigInteger.valueOf(1700000001000L)), TerminalAction.create_TData(dafny.DafnySequence.asUnicodeString("")), TerminalAction.create_TCommandFinished(dafny.DafnySequence.asUnicodeString("cmd-2"), java.math.BigInteger.ONE, java.math.BigInteger.valueOf(5L))));
      dafny.DafnySequence<? extends Part> _40_expLc;
      _40_expLc = dafny.DafnySequence.<Part> of(Part._typeDescriptor(), Part.create_Unclassified(dafny.DafnySequence.asUnicodeString("$ ")), Part.create_Command(dafny.DafnySequence.asUnicodeString("cmd-1"), dafny.DafnySequence.asUnicodeString("echo hello"), dafny.DafnySequence.asUnicodeString("hello\r\n"), java.math.BigInteger.valueOf(1700000000000L), true, agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.ZERO), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.valueOf(10L))), Part.create_Unclassified(dafny.DafnySequence.asUnicodeString("$ ")), Part.create_Command(dafny.DafnySequence.asUnicodeString("cmd-2"), dafny.DafnySequence.asUnicodeString("exit 1"), dafny.DafnySequence.asUnicodeString(""), java.math.BigInteger.valueOf(1700000001000L), true, agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.ONE), agency.open.ahp.AhpSkeleton.Option.<java.math.BigInteger>create_Some(dafny.TypeDescriptor.BIG_INTEGER, java.math.BigInteger.valueOf(5L))));
      if ((((_37_lc).dtor_content()).equals(_40_expLc)) && (java.util.Objects.equals((_37_lc).dtor_supportsCommandDetection(), agency.open.ahp.AhpSkeleton.Option.<Boolean>create_Some(dafny.TypeDescriptor.BOOLEAN, true)))) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
    }
    return new dafny.Tuple2<>(pass, total);
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.Terminal._default";
  }
}
