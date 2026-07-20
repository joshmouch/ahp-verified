// Package Terminal
// Dafny module Terminal compiled into Go

package Terminal

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-verified/go/AhpSkeleton"
	m_Annotations "github.com/joshmouch/ahp-verified/go/Annotations"
	m_Canvas "github.com/joshmouch/ahp-verified/go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-verified/go/Changeset"
	m_ConfluxCodec "github.com/joshmouch/ahp-verified/go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-verified/go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-verified/go/ConfluxOperators"
	m_ConfluxPrune "github.com/joshmouch/ahp-verified/go/ConfluxPrune"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-verified/go/ConfluxSeqRoute"
	m_ResourceWatch "github.com/joshmouch/ahp-verified/go/ResourceWatch"
	m__System "github.com/joshmouch/ahp-verified/go/System_"
	_dafny "github.com/joshmouch/ahp-verified/go/dafny"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_ConfluxCodec.Dummy__
var _ m_ConfluxContract.Dummy__
var _ m_AhpSkeleton.Dummy__
var _ m_ResourceWatch.Dummy__
var _ m_Canvas.Dummy__
var _ m_ConfluxOperators.Dummy__
var _ m_ConfluxPrune.Dummy__
var _ m_ConfluxSeqRoute.Dummy__
var _ m_Changeset.Dummy__
var _ m_Annotations.Dummy__

type Dummy__ struct{}

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "Terminal.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) AppendData(content _dafny.Sequence, d _dafny.Sequence) _dafny.Sequence {
	if (_dafny.IntOfUint32((content).Cardinality())).Sign() == 0 {
		return _dafny.SeqOf(Companion_Part_.Create_Unclassified_(d))
	} else {
		var _0_last Part = (content).Select(((_dafny.IntOfUint32((content).Cardinality())).Minus(_dafny.One)).Uint32()).(Part)
		_ = _0_last
		var _1_init _dafny.Sequence = (content).Take(((_dafny.IntOfUint32((content).Cardinality())).Minus(_dafny.One)).Uint32())
		_ = _1_init
		var _source0 Part = _0_last
		_ = _source0
		{
			if _source0.Is_Command() {
				var _2_cid _dafny.Sequence = _source0.Get_().(Part_Command).CommandId
				_ = _2_cid
				var _3_cl _dafny.Sequence = _source0.Get_().(Part_Command).CommandLine
				_ = _3_cl
				var _4_out _dafny.Sequence = _source0.Get_().(Part_Command).Output
				_ = _4_out
				var _5_ts _dafny.Int = _source0.Get_().(Part_Command).Timestamp
				_ = _5_ts
				var _6_comp bool = _source0.Get_().(Part_Command).IsComplete
				_ = _6_comp
				var _7_ec m_AhpSkeleton.Option = _source0.Get_().(Part_Command).ExitCode
				_ = _7_ec
				var _8_dm m_AhpSkeleton.Option = _source0.Get_().(Part_Command).DurationMs
				_ = _8_dm
				if !(_6_comp) {
					return _dafny.Companion_Sequence_.Concatenate(_1_init, _dafny.SeqOf(Companion_Part_.Create_Command_(_2_cid, _3_cl, _dafny.Companion_Sequence_.Concatenate(_4_out, d), _5_ts, _6_comp, _7_ec, _8_dm)))
				} else {
					return _dafny.Companion_Sequence_.Concatenate(content, _dafny.SeqOf(Companion_Part_.Create_Unclassified_(d)))
				}
			}
		}
		{
			var _9_v _dafny.Sequence = _source0.Get_().(Part_Unclassified).Value
			_ = _9_v
			return _dafny.Companion_Sequence_.Concatenate(_1_init, _dafny.SeqOf(Companion_Part_.Create_Unclassified_(_dafny.Companion_Sequence_.Concatenate(_9_v, d))))
		}
	}
}
func (_static *CompanionStruct_Default___) FinishCommand(content _dafny.Sequence, id _dafny.Sequence, code _dafny.Int, dur _dafny.Int) _dafny.Sequence {
	return m_ConfluxOperators.Companion_Default___.Map(func(coer21 func(Part) Part) func(interface{}) interface{} {
		return func(arg27 interface{}) interface{} {
			return coer21(arg27.(Part))
		}
	}((func(_0_id _dafny.Sequence, _1_code _dafny.Int, _2_dur _dafny.Int) func(Part) Part {
		return func(_3_p Part) Part {
			return func() Part {
				var _source0 Part = _3_p
				_ = _source0
				{
					if _source0.Is_Command() {
						var _4_cid _dafny.Sequence = _source0.Get_().(Part_Command).CommandId
						_ = _4_cid
						var _5_cl _dafny.Sequence = _source0.Get_().(Part_Command).CommandLine
						_ = _5_cl
						var _6_out _dafny.Sequence = _source0.Get_().(Part_Command).Output
						_ = _6_out
						var _7_ts _dafny.Int = _source0.Get_().(Part_Command).Timestamp
						_ = _7_ts
						var _8_comp bool = _source0.Get_().(Part_Command).IsComplete
						_ = _8_comp
						var _9_ec m_AhpSkeleton.Option = _source0.Get_().(Part_Command).ExitCode
						_ = _9_ec
						var _10_dm m_AhpSkeleton.Option = _source0.Get_().(Part_Command).DurationMs
						_ = _10_dm
						if _dafny.Companion_Sequence_.Equal(_4_cid, _0_id) {
							return Companion_Part_.Create_Command_(_4_cid, _5_cl, _6_out, _7_ts, true, m_AhpSkeleton.Companion_Option_.Create_Some_(_1_code), m_AhpSkeleton.Companion_Option_.Create_Some_(_2_dur))
						} else {
							return _3_p
						}
					}
				}
				{
					return _3_p
				}
			}()
		}
	})(id, code, dur)), content)
}
func (_static *CompanionStruct_Default___) ApplyToTerminal(s TerminalState, a TerminalAction, now _dafny.Int) _dafny.Tuple {
	var _pat_let_tv0 = s
	_ = _pat_let_tv0
	var _pat_let_tv1 = s
	_ = _pat_let_tv1
	var _pat_let_tv2 = s
	_ = _pat_let_tv2
	var _source0 TerminalAction = a
	_ = _source0
	{
		if _source0.Is_TCwdChanged() {
			var _0_c _dafny.Sequence = _source0.Get_().(TerminalAction_TCwdChanged).Cwd
			_ = _0_c
			return _dafny.TupleOf(func(_pat_let60_0 TerminalState) TerminalState {
				return func(_1_dt__update__tmp_h0 TerminalState) TerminalState {
					return func(_pat_let61_0 m_AhpSkeleton.Option) TerminalState {
						return func(_2_dt__update_hcwd_h0 m_AhpSkeleton.Option) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_((_1_dt__update__tmp_h0).Dtor_title(), _2_dt__update_hcwd_h0, (_1_dt__update__tmp_h0).Dtor_cols(), (_1_dt__update__tmp_h0).Dtor_rows(), (_1_dt__update__tmp_h0).Dtor_content(), (_1_dt__update__tmp_h0).Dtor_claim(), (_1_dt__update__tmp_h0).Dtor_exitCode(), (_1_dt__update__tmp_h0).Dtor_supportsCommandDetection())
						}(_pat_let61_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_0_c))
				}(_pat_let60_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_TTitleChanged() {
			var _3_t _dafny.Sequence = _source0.Get_().(TerminalAction_TTitleChanged).Title
			_ = _3_t
			return _dafny.TupleOf(func(_pat_let62_0 TerminalState) TerminalState {
				return func(_4_dt__update__tmp_h1 TerminalState) TerminalState {
					return func(_pat_let63_0 _dafny.Sequence) TerminalState {
						return func(_5_dt__update_htitle_h0 _dafny.Sequence) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_(_5_dt__update_htitle_h0, (_4_dt__update__tmp_h1).Dtor_cwd(), (_4_dt__update__tmp_h1).Dtor_cols(), (_4_dt__update__tmp_h1).Dtor_rows(), (_4_dt__update__tmp_h1).Dtor_content(), (_4_dt__update__tmp_h1).Dtor_claim(), (_4_dt__update__tmp_h1).Dtor_exitCode(), (_4_dt__update__tmp_h1).Dtor_supportsCommandDetection())
						}(_pat_let63_0)
					}(_3_t)
				}(_pat_let62_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_TResized() {
			var _6_co _dafny.Int = _source0.Get_().(TerminalAction_TResized).Cols
			_ = _6_co
			var _7_ro _dafny.Int = _source0.Get_().(TerminalAction_TResized).Rows
			_ = _7_ro
			return _dafny.TupleOf(func(_pat_let64_0 TerminalState) TerminalState {
				return func(_8_dt__update__tmp_h2 TerminalState) TerminalState {
					return func(_pat_let65_0 m_AhpSkeleton.Option) TerminalState {
						return func(_9_dt__update_hrows_h0 m_AhpSkeleton.Option) TerminalState {
							return func(_pat_let66_0 m_AhpSkeleton.Option) TerminalState {
								return func(_10_dt__update_hcols_h0 m_AhpSkeleton.Option) TerminalState {
									return Companion_TerminalState_.Create_TerminalState_((_8_dt__update__tmp_h2).Dtor_title(), (_8_dt__update__tmp_h2).Dtor_cwd(), _10_dt__update_hcols_h0, _9_dt__update_hrows_h0, (_8_dt__update__tmp_h2).Dtor_content(), (_8_dt__update__tmp_h2).Dtor_claim(), (_8_dt__update__tmp_h2).Dtor_exitCode(), (_8_dt__update__tmp_h2).Dtor_supportsCommandDetection())
								}(_pat_let66_0)
							}(m_AhpSkeleton.Companion_Option_.Create_Some_(_6_co))
						}(_pat_let65_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_7_ro))
				}(_pat_let64_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_TExited() {
			var _11_code _dafny.Int = _source0.Get_().(TerminalAction_TExited).Code
			_ = _11_code
			return _dafny.TupleOf(func(_pat_let67_0 TerminalState) TerminalState {
				return func(_12_dt__update__tmp_h3 TerminalState) TerminalState {
					return func(_pat_let68_0 m_AhpSkeleton.Option) TerminalState {
						return func(_13_dt__update_hexitCode_h0 m_AhpSkeleton.Option) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_((_12_dt__update__tmp_h3).Dtor_title(), (_12_dt__update__tmp_h3).Dtor_cwd(), (_12_dt__update__tmp_h3).Dtor_cols(), (_12_dt__update__tmp_h3).Dtor_rows(), (_12_dt__update__tmp_h3).Dtor_content(), (_12_dt__update__tmp_h3).Dtor_claim(), _13_dt__update_hexitCode_h0, (_12_dt__update__tmp_h3).Dtor_supportsCommandDetection())
						}(_pat_let68_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_11_code))
				}(_pat_let67_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_TData() {
			var _14_d _dafny.Sequence = _source0.Get_().(TerminalAction_TData).Data
			_ = _14_d
			return _dafny.TupleOf(func(_pat_let69_0 TerminalState) TerminalState {
				return func(_15_dt__update__tmp_h4 TerminalState) TerminalState {
					return func(_pat_let70_0 _dafny.Sequence) TerminalState {
						return func(_16_dt__update_hcontent_h0 _dafny.Sequence) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_((_15_dt__update__tmp_h4).Dtor_title(), (_15_dt__update__tmp_h4).Dtor_cwd(), (_15_dt__update__tmp_h4).Dtor_cols(), (_15_dt__update__tmp_h4).Dtor_rows(), _16_dt__update_hcontent_h0, (_15_dt__update__tmp_h4).Dtor_claim(), (_15_dt__update__tmp_h4).Dtor_exitCode(), (_15_dt__update__tmp_h4).Dtor_supportsCommandDetection())
						}(_pat_let70_0)
					}(Companion_Default___.AppendData((_pat_let_tv0).Dtor_content(), _14_d))
				}(_pat_let69_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_TCleared() {
			return _dafny.TupleOf(func(_pat_let71_0 TerminalState) TerminalState {
				return func(_17_dt__update__tmp_h5 TerminalState) TerminalState {
					return func(_pat_let72_0 _dafny.Sequence) TerminalState {
						return func(_18_dt__update_hcontent_h1 _dafny.Sequence) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_((_17_dt__update__tmp_h5).Dtor_title(), (_17_dt__update__tmp_h5).Dtor_cwd(), (_17_dt__update__tmp_h5).Dtor_cols(), (_17_dt__update__tmp_h5).Dtor_rows(), _18_dt__update_hcontent_h1, (_17_dt__update__tmp_h5).Dtor_claim(), (_17_dt__update__tmp_h5).Dtor_exitCode(), (_17_dt__update__tmp_h5).Dtor_supportsCommandDetection())
						}(_pat_let72_0)
					}(_dafny.SeqOf())
				}(_pat_let71_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_TClaimed() {
			var _19_c m_ConfluxCodec.Json = _source0.Get_().(TerminalAction_TClaimed).Claim
			_ = _19_c
			return _dafny.TupleOf(func(_pat_let73_0 TerminalState) TerminalState {
				return func(_20_dt__update__tmp_h6 TerminalState) TerminalState {
					return func(_pat_let74_0 m_AhpSkeleton.Option) TerminalState {
						return func(_21_dt__update_hclaim_h0 m_AhpSkeleton.Option) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_((_20_dt__update__tmp_h6).Dtor_title(), (_20_dt__update__tmp_h6).Dtor_cwd(), (_20_dt__update__tmp_h6).Dtor_cols(), (_20_dt__update__tmp_h6).Dtor_rows(), (_20_dt__update__tmp_h6).Dtor_content(), _21_dt__update_hclaim_h0, (_20_dt__update__tmp_h6).Dtor_exitCode(), (_20_dt__update__tmp_h6).Dtor_supportsCommandDetection())
						}(_pat_let74_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_19_c))
				}(_pat_let73_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_TCommandDetectionAvailable() {
			return _dafny.TupleOf(func(_pat_let75_0 TerminalState) TerminalState {
				return func(_22_dt__update__tmp_h7 TerminalState) TerminalState {
					return func(_pat_let76_0 m_AhpSkeleton.Option) TerminalState {
						return func(_23_dt__update_hsupportsCommandDetection_h0 m_AhpSkeleton.Option) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_((_22_dt__update__tmp_h7).Dtor_title(), (_22_dt__update__tmp_h7).Dtor_cwd(), (_22_dt__update__tmp_h7).Dtor_cols(), (_22_dt__update__tmp_h7).Dtor_rows(), (_22_dt__update__tmp_h7).Dtor_content(), (_22_dt__update__tmp_h7).Dtor_claim(), (_22_dt__update__tmp_h7).Dtor_exitCode(), _23_dt__update_hsupportsCommandDetection_h0)
						}(_pat_let76_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(true))
				}(_pat_let75_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_TCommandExecuted() {
			var _24_id _dafny.Sequence = _source0.Get_().(TerminalAction_TCommandExecuted).CommandId
			_ = _24_id
			var _25_cl _dafny.Sequence = _source0.Get_().(TerminalAction_TCommandExecuted).CommandLine
			_ = _25_cl
			var _26_ts _dafny.Int = _source0.Get_().(TerminalAction_TCommandExecuted).Timestamp
			_ = _26_ts
			return _dafny.TupleOf(func(_pat_let77_0 TerminalState) TerminalState {
				return func(_27_dt__update__tmp_h8 TerminalState) TerminalState {
					return func(_pat_let78_0 m_AhpSkeleton.Option) TerminalState {
						return func(_28_dt__update_hsupportsCommandDetection_h1 m_AhpSkeleton.Option) TerminalState {
							return func(_pat_let79_0 _dafny.Sequence) TerminalState {
								return func(_29_dt__update_hcontent_h2 _dafny.Sequence) TerminalState {
									return Companion_TerminalState_.Create_TerminalState_((_27_dt__update__tmp_h8).Dtor_title(), (_27_dt__update__tmp_h8).Dtor_cwd(), (_27_dt__update__tmp_h8).Dtor_cols(), (_27_dt__update__tmp_h8).Dtor_rows(), _29_dt__update_hcontent_h2, (_27_dt__update__tmp_h8).Dtor_claim(), (_27_dt__update__tmp_h8).Dtor_exitCode(), _28_dt__update_hsupportsCommandDetection_h1)
								}(_pat_let79_0)
							}(_dafny.Companion_Sequence_.Concatenate((_pat_let_tv1).Dtor_content(), _dafny.SeqOf(Companion_Part_.Create_Command_(_24_id, _25_cl, _dafny.UnicodeSeqOfUtf8Bytes(""), _26_ts, false, m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()))))
						}(_pat_let78_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(true))
				}(_pat_let77_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_TCommandFinished() {
			var _30_id _dafny.Sequence = _source0.Get_().(TerminalAction_TCommandFinished).CommandId
			_ = _30_id
			var _31_code _dafny.Int = _source0.Get_().(TerminalAction_TCommandFinished).Code
			_ = _31_code
			var _32_dur _dafny.Int = _source0.Get_().(TerminalAction_TCommandFinished).DurationMs
			_ = _32_dur
			return _dafny.TupleOf(func(_pat_let80_0 TerminalState) TerminalState {
				return func(_33_dt__update__tmp_h9 TerminalState) TerminalState {
					return func(_pat_let81_0 _dafny.Sequence) TerminalState {
						return func(_34_dt__update_hcontent_h3 _dafny.Sequence) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_((_33_dt__update__tmp_h9).Dtor_title(), (_33_dt__update__tmp_h9).Dtor_cwd(), (_33_dt__update__tmp_h9).Dtor_cols(), (_33_dt__update__tmp_h9).Dtor_rows(), _34_dt__update_hcontent_h3, (_33_dt__update__tmp_h9).Dtor_claim(), (_33_dt__update__tmp_h9).Dtor_exitCode(), (_33_dt__update__tmp_h9).Dtor_supportsCommandDetection())
						}(_pat_let81_0)
					}(Companion_Default___.FinishCommand((_pat_let_tv2).Dtor_content(), _30_id, _31_code, _32_dur))
				}(_pat_let80_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_TInput() {
			return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
		}
	}
	{
		return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
	}
}
func (_static *CompanionStruct_Default___) Apply1(s TerminalState, a TerminalAction) TerminalState {
	return (*(Companion_Default___.ApplyToTerminal(s, a, _dafny.IntOfInt64(9999))).IndexInt(0)).(TerminalState)
}
func (_static *CompanionStruct_Default___) Fold(s TerminalState, acts _dafny.Sequence) TerminalState {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer22 func(TerminalState, TerminalAction) TerminalState) func(interface{}, interface{}) interface{} {
		return func(arg28 interface{}, arg29 interface{}) interface{} {
			return coer22(arg28.(TerminalState), arg29.(TerminalAction))
		}
	}(Companion_Default___.Apply1), s, acts).(TerminalState)
}
func (_static *CompanionStruct_Default___) T0() TerminalState {
	return Companion_TerminalState_.Create_TerminalState_(_dafny.UnicodeSeqOfUtf8Bytes("bash"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())
}
func (_static *CompanionStruct_Default___) CL() m_ConfluxCodec.Json {
	return m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("client"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("clientId"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("c1"))))
}
func (_static *CompanionStruct_Default___) RunCorpus() (_dafny.Int, _dafny.Int) {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	var total _dafny.Int = _dafny.Zero
	_ = total
	total = _dafny.IntOfInt64(19)
	pass = _dafny.Zero
	var _0_base TerminalState
	_ = _0_base
	var _1_dt__update__tmp_h0 TerminalState = Companion_Default___.T0()
	_ = _1_dt__update__tmp_h0
	var _2_dt__update_hclaim_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Default___.CL())
	_ = _2_dt__update_hclaim_h0
	_0_base = Companion_TerminalState_.Create_TerminalState_((_1_dt__update__tmp_h0).Dtor_title(), (_1_dt__update__tmp_h0).Dtor_cwd(), (_1_dt__update__tmp_h0).Dtor_cols(), (_1_dt__update__tmp_h0).Dtor_rows(), (_1_dt__update__tmp_h0).Dtor_content(), _2_dt__update_hclaim_h0, (_1_dt__update__tmp_h0).Dtor_exitCode(), (_1_dt__update__tmp_h0).Dtor_supportsCommandDetection())
	if ((Companion_Default___.Apply1(_0_base, Companion_TerminalAction_.Create_TCwdChanged_(_dafny.UnicodeSeqOfUtf8Bytes("/tmp")))).Dtor_cwd()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("/tmp"))) {
		pass = (pass).Plus(_dafny.One)
	}
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_0_base, Companion_TerminalAction_.Create_TTitleChanged_(_dafny.UnicodeSeqOfUtf8Bytes("zsh")))).Dtor_title(), _dafny.UnicodeSeqOfUtf8Bytes("zsh")) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_0_base, Companion_TerminalAction_.Create_TResized_(_dafny.IntOfInt64(80), _dafny.IntOfInt64(24)))).Equals(func(_pat_let82_0 TerminalState) TerminalState {
		return func(_3_dt__update__tmp_h1 TerminalState) TerminalState {
			return func(_pat_let83_0 m_AhpSkeleton.Option) TerminalState {
				return func(_4_dt__update_hrows_h0 m_AhpSkeleton.Option) TerminalState {
					return func(_pat_let84_0 m_AhpSkeleton.Option) TerminalState {
						return func(_5_dt__update_hcols_h0 m_AhpSkeleton.Option) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_((_3_dt__update__tmp_h1).Dtor_title(), (_3_dt__update__tmp_h1).Dtor_cwd(), _5_dt__update_hcols_h0, _4_dt__update_hrows_h0, (_3_dt__update__tmp_h1).Dtor_content(), (_3_dt__update__tmp_h1).Dtor_claim(), (_3_dt__update__tmp_h1).Dtor_exitCode(), (_3_dt__update__tmp_h1).Dtor_supportsCommandDetection())
						}(_pat_let84_0)
					}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.IntOfInt64(80)))
				}(_pat_let83_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.IntOfInt64(24)))
		}(_pat_let82_0)
	}(_0_base)) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_0_base, Companion_TerminalAction_.Create_TExited_(_dafny.Zero))).Dtor_exitCode()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.Zero)) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_0_base, Companion_TerminalAction_.Create_TInput_())).Equals(_0_base) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_0_base, Companion_TerminalAction_.Create_TUnknown_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("type"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("terminal/nope"))))))).Equals(_0_base) {
		pass = (pass).Plus(_dafny.One)
	}
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let85_0 TerminalState) TerminalState {
		return func(_6_dt__update__tmp_h2 TerminalState) TerminalState {
			return func(_pat_let86_0 _dafny.Sequence) TerminalState {
				return func(_7_dt__update_hcontent_h0 _dafny.Sequence) TerminalState {
					return Companion_TerminalState_.Create_TerminalState_((_6_dt__update__tmp_h2).Dtor_title(), (_6_dt__update__tmp_h2).Dtor_cwd(), (_6_dt__update__tmp_h2).Dtor_cols(), (_6_dt__update__tmp_h2).Dtor_rows(), _7_dt__update_hcontent_h0, (_6_dt__update__tmp_h2).Dtor_claim(), (_6_dt__update__tmp_h2).Dtor_exitCode(), (_6_dt__update__tmp_h2).Dtor_supportsCommandDetection())
				}(_pat_let86_0)
			}(_dafny.SeqOf(Companion_Part_.Create_Unclassified_(_dafny.UnicodeSeqOfUtf8Bytes("x"))))
		}(_pat_let85_0)
	}(_0_base), Companion_TerminalAction_.Create_TCleared_())).Dtor_content(), _dafny.SeqOf()) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(func(_pat_let87_0 TerminalState) TerminalState {
		return func(_8_dt__update__tmp_h3 TerminalState) TerminalState {
			return func(_pat_let88_0 m_AhpSkeleton.Option) TerminalState {
				return func(_9_dt__update_hsupportsCommandDetection_h0 m_AhpSkeleton.Option) TerminalState {
					return func(_pat_let89_0 _dafny.Sequence) TerminalState {
						return func(_10_dt__update_hcontent_h1 _dafny.Sequence) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_((_8_dt__update__tmp_h3).Dtor_title(), (_8_dt__update__tmp_h3).Dtor_cwd(), (_8_dt__update__tmp_h3).Dtor_cols(), (_8_dt__update__tmp_h3).Dtor_rows(), _10_dt__update_hcontent_h1, (_8_dt__update__tmp_h3).Dtor_claim(), (_8_dt__update__tmp_h3).Dtor_exitCode(), _9_dt__update_hsupportsCommandDetection_h0)
						}(_pat_let89_0)
					}(_dafny.SeqOf(Companion_Part_.Create_Unclassified_(_dafny.UnicodeSeqOfUtf8Bytes("x"))))
				}(_pat_let88_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(true))
		}(_pat_let87_0)
	}(_0_base), Companion_TerminalAction_.Create_TCleared_())).Equals(func(_pat_let90_0 TerminalState) TerminalState {
		return func(_11_dt__update__tmp_h4 TerminalState) TerminalState {
			return func(_pat_let91_0 m_AhpSkeleton.Option) TerminalState {
				return func(_12_dt__update_hsupportsCommandDetection_h1 m_AhpSkeleton.Option) TerminalState {
					return func(_pat_let92_0 _dafny.Sequence) TerminalState {
						return func(_13_dt__update_hcontent_h2 _dafny.Sequence) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_((_11_dt__update__tmp_h4).Dtor_title(), (_11_dt__update__tmp_h4).Dtor_cwd(), (_11_dt__update__tmp_h4).Dtor_cols(), (_11_dt__update__tmp_h4).Dtor_rows(), _13_dt__update_hcontent_h2, (_11_dt__update__tmp_h4).Dtor_claim(), (_11_dt__update__tmp_h4).Dtor_exitCode(), _12_dt__update_hsupportsCommandDetection_h1)
						}(_pat_let92_0)
					}(_dafny.SeqOf())
				}(_pat_let91_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(true))
		}(_pat_let90_0)
	}(_0_base)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _14_sc m_ConfluxCodec.Json
	_ = _14_sc
	_14_sc = m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("session"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("session"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("s1"))))
	if ((Companion_Default___.Apply1(_0_base, Companion_TerminalAction_.Create_TClaimed_(_14_sc))).Dtor_claim()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_14_sc)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _15_scTool m_ConfluxCodec.Json
	_ = _15_scTool
	_15_scTool = m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kind"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("session"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("session"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("s1"))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("toolCallId"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("t1"))))
	if ((Companion_Default___.Apply1(_0_base, Companion_TerminalAction_.Create_TClaimed_(_15_scTool))).Dtor_claim()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_15_scTool)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _pat_let_tv0 = _14_sc
	_ = _pat_let_tv0
	if ((Companion_Default___.Apply1(func(_pat_let93_0 TerminalState) TerminalState {
		return func(_16_dt__update__tmp_h5 TerminalState) TerminalState {
			return func(_pat_let94_0 m_AhpSkeleton.Option) TerminalState {
				return func(_17_dt__update_hclaim_h1 m_AhpSkeleton.Option) TerminalState {
					return Companion_TerminalState_.Create_TerminalState_((_16_dt__update__tmp_h5).Dtor_title(), (_16_dt__update__tmp_h5).Dtor_cwd(), (_16_dt__update__tmp_h5).Dtor_cols(), (_16_dt__update__tmp_h5).Dtor_rows(), (_16_dt__update__tmp_h5).Dtor_content(), _17_dt__update_hclaim_h1, (_16_dt__update__tmp_h5).Dtor_exitCode(), (_16_dt__update__tmp_h5).Dtor_supportsCommandDetection())
				}(_pat_let94_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(_pat_let_tv0))
		}(_pat_let93_0)
	}(_0_base), Companion_TerminalAction_.Create_TClaimed_(Companion_Default___.CL()))).Dtor_claim()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Default___.CL())) {
		pass = (pass).Plus(_dafny.One)
	}
	if ((Companion_Default___.Apply1(_0_base, Companion_TerminalAction_.Create_TCommandDetectionAvailable_())).Dtor_supportsCommandDetection()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(true)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _18_afterExec TerminalState
	_ = _18_afterExec
	_18_afterExec = Companion_Default___.Apply1(_0_base, Companion_TerminalAction_.Create_TCommandExecuted_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-1"), _dafny.UnicodeSeqOfUtf8Bytes("npm test"), _dafny.IntOfInt64(1700000000000)))
	if (_18_afterExec).Equals(func(_pat_let95_0 TerminalState) TerminalState {
		return func(_19_dt__update__tmp_h6 TerminalState) TerminalState {
			return func(_pat_let96_0 m_AhpSkeleton.Option) TerminalState {
				return func(_20_dt__update_hsupportsCommandDetection_h2 m_AhpSkeleton.Option) TerminalState {
					return func(_pat_let97_0 _dafny.Sequence) TerminalState {
						return func(_21_dt__update_hcontent_h3 _dafny.Sequence) TerminalState {
							return Companion_TerminalState_.Create_TerminalState_((_19_dt__update__tmp_h6).Dtor_title(), (_19_dt__update__tmp_h6).Dtor_cwd(), (_19_dt__update__tmp_h6).Dtor_cols(), (_19_dt__update__tmp_h6).Dtor_rows(), _21_dt__update_hcontent_h3, (_19_dt__update__tmp_h6).Dtor_claim(), (_19_dt__update__tmp_h6).Dtor_exitCode(), _20_dt__update_hsupportsCommandDetection_h2)
						}(_pat_let97_0)
					}(_dafny.SeqOf(Companion_Part_.Create_Command_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-1"), _dafny.UnicodeSeqOfUtf8Bytes("npm test"), _dafny.UnicodeSeqOfUtf8Bytes(""), _dafny.IntOfInt64(1700000000000), false, m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())))
				}(_pat_let96_0)
			}(m_AhpSkeleton.Companion_Option_.Create_Some_(true))
		}(_pat_let95_0)
	}(_0_base)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _22_withCmd TerminalState
	_ = _22_withCmd
	var _23_dt__update__tmp_h7 TerminalState = _0_base
	_ = _23_dt__update__tmp_h7
	var _24_dt__update_hsupportsCommandDetection_h3 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(true)
	_ = _24_dt__update_hsupportsCommandDetection_h3
	var _25_dt__update_hcontent_h4 _dafny.Sequence = _dafny.SeqOf(Companion_Part_.Create_Command_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-1"), _dafny.UnicodeSeqOfUtf8Bytes("npm test"), _dafny.UnicodeSeqOfUtf8Bytes("All tests passed\r\n"), _dafny.IntOfInt64(1700000000000), false, m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()))
	_ = _25_dt__update_hcontent_h4
	_22_withCmd = Companion_TerminalState_.Create_TerminalState_((_23_dt__update__tmp_h7).Dtor_title(), (_23_dt__update__tmp_h7).Dtor_cwd(), (_23_dt__update__tmp_h7).Dtor_cols(), (_23_dt__update__tmp_h7).Dtor_rows(), _25_dt__update_hcontent_h4, (_23_dt__update__tmp_h7).Dtor_claim(), (_23_dt__update__tmp_h7).Dtor_exitCode(), _24_dt__update_hsupportsCommandDetection_h3)
	var _26_doneCmd TerminalState
	_ = _26_doneCmd
	var _27_dt__update__tmp_h8 TerminalState = _0_base
	_ = _27_dt__update__tmp_h8
	var _28_dt__update_hsupportsCommandDetection_h4 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(true)
	_ = _28_dt__update_hsupportsCommandDetection_h4
	var _29_dt__update_hcontent_h5 _dafny.Sequence = _dafny.SeqOf(Companion_Part_.Create_Command_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-1"), _dafny.UnicodeSeqOfUtf8Bytes("npm test"), _dafny.UnicodeSeqOfUtf8Bytes("All tests passed\r\n"), _dafny.IntOfInt64(1700000000000), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.Zero), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.IntOfInt64(1234))))
	_ = _29_dt__update_hcontent_h5
	_26_doneCmd = Companion_TerminalState_.Create_TerminalState_((_27_dt__update__tmp_h8).Dtor_title(), (_27_dt__update__tmp_h8).Dtor_cwd(), (_27_dt__update__tmp_h8).Dtor_cols(), (_27_dt__update__tmp_h8).Dtor_rows(), _29_dt__update_hcontent_h5, (_27_dt__update__tmp_h8).Dtor_claim(), (_27_dt__update__tmp_h8).Dtor_exitCode(), _28_dt__update_hsupportsCommandDetection_h4)
	if (Companion_Default___.Apply1(_22_withCmd, Companion_TerminalAction_.Create_TCommandFinished_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-1"), _dafny.Zero, _dafny.IntOfInt64(1234)))).Equals(_26_doneCmd) {
		pass = (pass).Plus(_dafny.One)
	}
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(func(_pat_let98_0 TerminalState) TerminalState {
		return func(_30_dt__update__tmp_h9 TerminalState) TerminalState {
			return func(_pat_let99_0 _dafny.Sequence) TerminalState {
				return func(_31_dt__update_hcontent_h6 _dafny.Sequence) TerminalState {
					return Companion_TerminalState_.Create_TerminalState_((_30_dt__update__tmp_h9).Dtor_title(), (_30_dt__update__tmp_h9).Dtor_cwd(), (_30_dt__update__tmp_h9).Dtor_cols(), (_30_dt__update__tmp_h9).Dtor_rows(), _31_dt__update_hcontent_h6, (_30_dt__update__tmp_h9).Dtor_claim(), (_30_dt__update__tmp_h9).Dtor_exitCode(), (_30_dt__update__tmp_h9).Dtor_supportsCommandDetection())
				}(_pat_let99_0)
			}(_dafny.SeqOf(Companion_Part_.Create_Unclassified_(_dafny.UnicodeSeqOfUtf8Bytes("a"))))
		}(_pat_let98_0)
	}(_0_base), Companion_TerminalAction_.Create_TData_(_dafny.UnicodeSeqOfUtf8Bytes("b")))).Dtor_content(), _dafny.SeqOf(Companion_Part_.Create_Unclassified_(_dafny.UnicodeSeqOfUtf8Bytes("ab")))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _32_complete TerminalState
	_ = _32_complete
	var _33_dt__update__tmp_h10 TerminalState = _0_base
	_ = _33_dt__update__tmp_h10
	var _34_dt__update_hsupportsCommandDetection_h5 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(true)
	_ = _34_dt__update_hsupportsCommandDetection_h5
	var _35_dt__update_hcontent_h7 _dafny.Sequence = _dafny.SeqOf(Companion_Part_.Create_Command_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-1"), _dafny.UnicodeSeqOfUtf8Bytes("echo hi"), _dafny.UnicodeSeqOfUtf8Bytes("hi\r\n"), _dafny.IntOfInt64(1700000000000), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.Zero), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.IntOfInt64(50))))
	_ = _35_dt__update_hcontent_h7
	_32_complete = Companion_TerminalState_.Create_TerminalState_((_33_dt__update__tmp_h10).Dtor_title(), (_33_dt__update__tmp_h10).Dtor_cwd(), (_33_dt__update__tmp_h10).Dtor_cols(), (_33_dt__update__tmp_h10).Dtor_rows(), _35_dt__update_hcontent_h7, (_33_dt__update__tmp_h10).Dtor_claim(), (_33_dt__update__tmp_h10).Dtor_exitCode(), _34_dt__update_hsupportsCommandDetection_h5)
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_32_complete, Companion_TerminalAction_.Create_TData_(_dafny.UnicodeSeqOfUtf8Bytes("$ ")))).Dtor_content(), _dafny.Companion_Sequence_.Concatenate((_32_complete).Dtor_content(), _dafny.SeqOf(Companion_Part_.Create_Unclassified_(_dafny.UnicodeSeqOfUtf8Bytes("$ "))))) {
		pass = (pass).Plus(_dafny.One)
	}
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(_22_withCmd, Companion_TerminalAction_.Create_TData_(_dafny.UnicodeSeqOfUtf8Bytes("!")))).Dtor_content(), _dafny.SeqOf(Companion_Part_.Create_Command_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-1"), _dafny.UnicodeSeqOfUtf8Bytes("npm test"), _dafny.UnicodeSeqOfUtf8Bytes("All tests passed\r\n!"), _dafny.IntOfInt64(1700000000000), false, m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _36_life TerminalState
	_ = _36_life
	_36_life = Companion_Default___.Fold(_0_base, _dafny.SeqOf(Companion_TerminalAction_.Create_TCwdChanged_(_dafny.UnicodeSeqOfUtf8Bytes("/w")), Companion_TerminalAction_.Create_TData_(_dafny.UnicodeSeqOfUtf8Bytes("x")), Companion_TerminalAction_.Create_TResized_(_dafny.IntOfInt64(100), _dafny.IntOfInt64(40)), Companion_TerminalAction_.Create_TClaimed_(_14_sc), Companion_TerminalAction_.Create_TExited_(_dafny.One)))
	if (((((_36_life).Dtor_cwd()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("/w")))) && (((_36_life).Dtor_cols()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.IntOfInt64(100))))) && (((_36_life).Dtor_exitCode()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.One)))) && (((_36_life).Dtor_claim()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(_14_sc))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _37_lc TerminalState
	_ = _37_lc
	_37_lc = Companion_Default___.Fold(func(_pat_let100_0 TerminalState) TerminalState {
		return func(_38_dt__update__tmp_h11 TerminalState) TerminalState {
			return func(_pat_let101_0 _dafny.Sequence) TerminalState {
				return func(_39_dt__update_hcontent_h8 _dafny.Sequence) TerminalState {
					return Companion_TerminalState_.Create_TerminalState_((_38_dt__update__tmp_h11).Dtor_title(), (_38_dt__update__tmp_h11).Dtor_cwd(), (_38_dt__update__tmp_h11).Dtor_cols(), (_38_dt__update__tmp_h11).Dtor_rows(), _39_dt__update_hcontent_h8, (_38_dt__update__tmp_h11).Dtor_claim(), (_38_dt__update__tmp_h11).Dtor_exitCode(), (_38_dt__update__tmp_h11).Dtor_supportsCommandDetection())
				}(_pat_let101_0)
			}(_dafny.SeqOf())
		}(_pat_let100_0)
	}(_0_base), _dafny.SeqOf(Companion_TerminalAction_.Create_TData_(_dafny.UnicodeSeqOfUtf8Bytes("$ ")), Companion_TerminalAction_.Create_TCommandExecuted_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-1"), _dafny.UnicodeSeqOfUtf8Bytes("echo hello"), _dafny.IntOfInt64(1700000000000)), Companion_TerminalAction_.Create_TData_(_dafny.UnicodeSeqOfUtf8Bytes("hello\r\n")), Companion_TerminalAction_.Create_TCommandFinished_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-1"), _dafny.Zero, _dafny.IntOfInt64(10)), Companion_TerminalAction_.Create_TData_(_dafny.UnicodeSeqOfUtf8Bytes("$ ")), Companion_TerminalAction_.Create_TCommandExecuted_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-2"), _dafny.UnicodeSeqOfUtf8Bytes("exit 1"), _dafny.IntOfInt64(1700000001000)), Companion_TerminalAction_.Create_TData_(_dafny.UnicodeSeqOfUtf8Bytes("")), Companion_TerminalAction_.Create_TCommandFinished_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-2"), _dafny.One, _dafny.IntOfInt64(5))))
	var _40_expLc _dafny.Sequence
	_ = _40_expLc
	_40_expLc = _dafny.SeqOf(Companion_Part_.Create_Unclassified_(_dafny.UnicodeSeqOfUtf8Bytes("$ ")), Companion_Part_.Create_Command_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-1"), _dafny.UnicodeSeqOfUtf8Bytes("echo hello"), _dafny.UnicodeSeqOfUtf8Bytes("hello\r\n"), _dafny.IntOfInt64(1700000000000), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.Zero), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.IntOfInt64(10))), Companion_Part_.Create_Unclassified_(_dafny.UnicodeSeqOfUtf8Bytes("$ ")), Companion_Part_.Create_Command_(_dafny.UnicodeSeqOfUtf8Bytes("cmd-2"), _dafny.UnicodeSeqOfUtf8Bytes("exit 1"), _dafny.UnicodeSeqOfUtf8Bytes(""), _dafny.IntOfInt64(1700000001000), true, m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.One), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.IntOfInt64(5))))
	if (_dafny.Companion_Sequence_.Equal((_37_lc).Dtor_content(), _40_expLc)) && (((_37_lc).Dtor_supportsCommandDetection()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(true))) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass, total
}

// End of class Default__

// Definition of datatype Part
type Part struct {
	Data_Part_
}

func (_this Part) Get_() Data_Part_ {
	return _this.Data_Part_
}

type Data_Part_ interface {
	isPart()
}

type CompanionStruct_Part_ struct {
}

var Companion_Part_ = CompanionStruct_Part_{}

type Part_Unclassified struct {
	Value _dafny.Sequence
}

func (Part_Unclassified) isPart() {}

func (CompanionStruct_Part_) Create_Unclassified_(Value _dafny.Sequence) Part {
	return Part{Part_Unclassified{Value}}
}

func (_this Part) Is_Unclassified() bool {
	_, ok := _this.Get_().(Part_Unclassified)
	return ok
}

type Part_Command struct {
	CommandId   _dafny.Sequence
	CommandLine _dafny.Sequence
	Output      _dafny.Sequence
	Timestamp   _dafny.Int
	IsComplete  bool
	ExitCode    m_AhpSkeleton.Option
	DurationMs  m_AhpSkeleton.Option
}

func (Part_Command) isPart() {}

func (CompanionStruct_Part_) Create_Command_(CommandId _dafny.Sequence, CommandLine _dafny.Sequence, Output _dafny.Sequence, Timestamp _dafny.Int, IsComplete bool, ExitCode m_AhpSkeleton.Option, DurationMs m_AhpSkeleton.Option) Part {
	return Part{Part_Command{CommandId, CommandLine, Output, Timestamp, IsComplete, ExitCode, DurationMs}}
}

func (_this Part) Is_Command() bool {
	_, ok := _this.Get_().(Part_Command)
	return ok
}

func (CompanionStruct_Part_) Default() Part {
	return Companion_Part_.Create_Unclassified_(_dafny.EmptySeq)
}

func (_this Part) Dtor_value() _dafny.Sequence {
	return _this.Get_().(Part_Unclassified).Value
}

func (_this Part) Dtor_commandId() _dafny.Sequence {
	return _this.Get_().(Part_Command).CommandId
}

func (_this Part) Dtor_commandLine() _dafny.Sequence {
	return _this.Get_().(Part_Command).CommandLine
}

func (_this Part) Dtor_output() _dafny.Sequence {
	return _this.Get_().(Part_Command).Output
}

func (_this Part) Dtor_timestamp() _dafny.Int {
	return _this.Get_().(Part_Command).Timestamp
}

func (_this Part) Dtor_isComplete() bool {
	return _this.Get_().(Part_Command).IsComplete
}

func (_this Part) Dtor_exitCode() m_AhpSkeleton.Option {
	return _this.Get_().(Part_Command).ExitCode
}

func (_this Part) Dtor_durationMs() m_AhpSkeleton.Option {
	return _this.Get_().(Part_Command).DurationMs
}

func (_this Part) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Part_Unclassified:
		{
			return "Terminal.Part.Unclassified" + "(" + data.Value.VerbatimString(true) + ")"
		}
	case Part_Command:
		{
			return "Terminal.Part.Command" + "(" + data.CommandId.VerbatimString(true) + ", " + data.CommandLine.VerbatimString(true) + ", " + data.Output.VerbatimString(true) + ", " + _dafny.String(data.Timestamp) + ", " + _dafny.String(data.IsComplete) + ", " + _dafny.String(data.ExitCode) + ", " + _dafny.String(data.DurationMs) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Part) Equals(other Part) bool {
	switch data1 := _this.Get_().(type) {
	case Part_Unclassified:
		{
			data2, ok := other.Get_().(Part_Unclassified)
			return ok && data1.Value.Equals(data2.Value)
		}
	case Part_Command:
		{
			data2, ok := other.Get_().(Part_Command)
			return ok && data1.CommandId.Equals(data2.CommandId) && data1.CommandLine.Equals(data2.CommandLine) && data1.Output.Equals(data2.Output) && data1.Timestamp.Cmp(data2.Timestamp) == 0 && data1.IsComplete == data2.IsComplete && data1.ExitCode.Equals(data2.ExitCode) && data1.DurationMs.Equals(data2.DurationMs)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Part) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Part)
	return ok && _this.Equals(typed)
}

func Type_Part_() _dafny.TypeDescriptor {
	return type_Part_{}
}

type type_Part_ struct {
}

func (_this type_Part_) Default() interface{} {
	return Companion_Part_.Default()
}

func (_this type_Part_) String() string {
	return "Terminal.Part"
}
func (_this Part) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Part{}

// End of datatype Part

// Definition of datatype TerminalState
type TerminalState struct {
	Data_TerminalState_
}

func (_this TerminalState) Get_() Data_TerminalState_ {
	return _this.Data_TerminalState_
}

type Data_TerminalState_ interface {
	isTerminalState()
}

type CompanionStruct_TerminalState_ struct {
}

var Companion_TerminalState_ = CompanionStruct_TerminalState_{}

type TerminalState_TerminalState struct {
	Title                    _dafny.Sequence
	Cwd                      m_AhpSkeleton.Option
	Cols                     m_AhpSkeleton.Option
	Rows                     m_AhpSkeleton.Option
	Content                  _dafny.Sequence
	Claim                    m_AhpSkeleton.Option
	ExitCode                 m_AhpSkeleton.Option
	SupportsCommandDetection m_AhpSkeleton.Option
}

func (TerminalState_TerminalState) isTerminalState() {}

func (CompanionStruct_TerminalState_) Create_TerminalState_(Title _dafny.Sequence, Cwd m_AhpSkeleton.Option, Cols m_AhpSkeleton.Option, Rows m_AhpSkeleton.Option, Content _dafny.Sequence, Claim m_AhpSkeleton.Option, ExitCode m_AhpSkeleton.Option, SupportsCommandDetection m_AhpSkeleton.Option) TerminalState {
	return TerminalState{TerminalState_TerminalState{Title, Cwd, Cols, Rows, Content, Claim, ExitCode, SupportsCommandDetection}}
}

func (_this TerminalState) Is_TerminalState() bool {
	_, ok := _this.Get_().(TerminalState_TerminalState)
	return ok
}

func (CompanionStruct_TerminalState_) Default() TerminalState {
	return Companion_TerminalState_.Create_TerminalState_(_dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default())
}

func (_this TerminalState) Dtor_title() _dafny.Sequence {
	return _this.Get_().(TerminalState_TerminalState).Title
}

func (_this TerminalState) Dtor_cwd() m_AhpSkeleton.Option {
	return _this.Get_().(TerminalState_TerminalState).Cwd
}

func (_this TerminalState) Dtor_cols() m_AhpSkeleton.Option {
	return _this.Get_().(TerminalState_TerminalState).Cols
}

func (_this TerminalState) Dtor_rows() m_AhpSkeleton.Option {
	return _this.Get_().(TerminalState_TerminalState).Rows
}

func (_this TerminalState) Dtor_content() _dafny.Sequence {
	return _this.Get_().(TerminalState_TerminalState).Content
}

func (_this TerminalState) Dtor_claim() m_AhpSkeleton.Option {
	return _this.Get_().(TerminalState_TerminalState).Claim
}

func (_this TerminalState) Dtor_exitCode() m_AhpSkeleton.Option {
	return _this.Get_().(TerminalState_TerminalState).ExitCode
}

func (_this TerminalState) Dtor_supportsCommandDetection() m_AhpSkeleton.Option {
	return _this.Get_().(TerminalState_TerminalState).SupportsCommandDetection
}

func (_this TerminalState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case TerminalState_TerminalState:
		{
			return "Terminal.TerminalState.TerminalState" + "(" + data.Title.VerbatimString(true) + ", " + _dafny.String(data.Cwd) + ", " + _dafny.String(data.Cols) + ", " + _dafny.String(data.Rows) + ", " + _dafny.String(data.Content) + ", " + _dafny.String(data.Claim) + ", " + _dafny.String(data.ExitCode) + ", " + _dafny.String(data.SupportsCommandDetection) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this TerminalState) Equals(other TerminalState) bool {
	switch data1 := _this.Get_().(type) {
	case TerminalState_TerminalState:
		{
			data2, ok := other.Get_().(TerminalState_TerminalState)
			return ok && data1.Title.Equals(data2.Title) && data1.Cwd.Equals(data2.Cwd) && data1.Cols.Equals(data2.Cols) && data1.Rows.Equals(data2.Rows) && data1.Content.Equals(data2.Content) && data1.Claim.Equals(data2.Claim) && data1.ExitCode.Equals(data2.ExitCode) && data1.SupportsCommandDetection.Equals(data2.SupportsCommandDetection)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this TerminalState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(TerminalState)
	return ok && _this.Equals(typed)
}

func Type_TerminalState_() _dafny.TypeDescriptor {
	return type_TerminalState_{}
}

type type_TerminalState_ struct {
}

func (_this type_TerminalState_) Default() interface{} {
	return Companion_TerminalState_.Default()
}

func (_this type_TerminalState_) String() string {
	return "Terminal.TerminalState"
}
func (_this TerminalState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = TerminalState{}

// End of datatype TerminalState

// Definition of datatype TerminalAction
type TerminalAction struct {
	Data_TerminalAction_
}

func (_this TerminalAction) Get_() Data_TerminalAction_ {
	return _this.Data_TerminalAction_
}

type Data_TerminalAction_ interface {
	isTerminalAction()
}

type CompanionStruct_TerminalAction_ struct {
}

var Companion_TerminalAction_ = CompanionStruct_TerminalAction_{}

type TerminalAction_TCwdChanged struct {
	Cwd _dafny.Sequence
}

func (TerminalAction_TCwdChanged) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TCwdChanged_(Cwd _dafny.Sequence) TerminalAction {
	return TerminalAction{TerminalAction_TCwdChanged{Cwd}}
}

func (_this TerminalAction) Is_TCwdChanged() bool {
	_, ok := _this.Get_().(TerminalAction_TCwdChanged)
	return ok
}

type TerminalAction_TTitleChanged struct {
	Title _dafny.Sequence
}

func (TerminalAction_TTitleChanged) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TTitleChanged_(Title _dafny.Sequence) TerminalAction {
	return TerminalAction{TerminalAction_TTitleChanged{Title}}
}

func (_this TerminalAction) Is_TTitleChanged() bool {
	_, ok := _this.Get_().(TerminalAction_TTitleChanged)
	return ok
}

type TerminalAction_TResized struct {
	Cols _dafny.Int
	Rows _dafny.Int
}

func (TerminalAction_TResized) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TResized_(Cols _dafny.Int, Rows _dafny.Int) TerminalAction {
	return TerminalAction{TerminalAction_TResized{Cols, Rows}}
}

func (_this TerminalAction) Is_TResized() bool {
	_, ok := _this.Get_().(TerminalAction_TResized)
	return ok
}

type TerminalAction_TExited struct {
	Code _dafny.Int
}

func (TerminalAction_TExited) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TExited_(Code _dafny.Int) TerminalAction {
	return TerminalAction{TerminalAction_TExited{Code}}
}

func (_this TerminalAction) Is_TExited() bool {
	_, ok := _this.Get_().(TerminalAction_TExited)
	return ok
}

type TerminalAction_TData struct {
	Data _dafny.Sequence
}

func (TerminalAction_TData) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TData_(Data _dafny.Sequence) TerminalAction {
	return TerminalAction{TerminalAction_TData{Data}}
}

func (_this TerminalAction) Is_TData() bool {
	_, ok := _this.Get_().(TerminalAction_TData)
	return ok
}

type TerminalAction_TCleared struct {
}

func (TerminalAction_TCleared) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TCleared_() TerminalAction {
	return TerminalAction{TerminalAction_TCleared{}}
}

func (_this TerminalAction) Is_TCleared() bool {
	_, ok := _this.Get_().(TerminalAction_TCleared)
	return ok
}

type TerminalAction_TClaimed struct {
	Claim m_ConfluxCodec.Json
}

func (TerminalAction_TClaimed) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TClaimed_(Claim m_ConfluxCodec.Json) TerminalAction {
	return TerminalAction{TerminalAction_TClaimed{Claim}}
}

func (_this TerminalAction) Is_TClaimed() bool {
	_, ok := _this.Get_().(TerminalAction_TClaimed)
	return ok
}

type TerminalAction_TCommandDetectionAvailable struct {
}

func (TerminalAction_TCommandDetectionAvailable) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TCommandDetectionAvailable_() TerminalAction {
	return TerminalAction{TerminalAction_TCommandDetectionAvailable{}}
}

func (_this TerminalAction) Is_TCommandDetectionAvailable() bool {
	_, ok := _this.Get_().(TerminalAction_TCommandDetectionAvailable)
	return ok
}

type TerminalAction_TCommandExecuted struct {
	CommandId   _dafny.Sequence
	CommandLine _dafny.Sequence
	Timestamp   _dafny.Int
}

func (TerminalAction_TCommandExecuted) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TCommandExecuted_(CommandId _dafny.Sequence, CommandLine _dafny.Sequence, Timestamp _dafny.Int) TerminalAction {
	return TerminalAction{TerminalAction_TCommandExecuted{CommandId, CommandLine, Timestamp}}
}

func (_this TerminalAction) Is_TCommandExecuted() bool {
	_, ok := _this.Get_().(TerminalAction_TCommandExecuted)
	return ok
}

type TerminalAction_TCommandFinished struct {
	CommandId  _dafny.Sequence
	Code       _dafny.Int
	DurationMs _dafny.Int
}

func (TerminalAction_TCommandFinished) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TCommandFinished_(CommandId _dafny.Sequence, Code _dafny.Int, DurationMs _dafny.Int) TerminalAction {
	return TerminalAction{TerminalAction_TCommandFinished{CommandId, Code, DurationMs}}
}

func (_this TerminalAction) Is_TCommandFinished() bool {
	_, ok := _this.Get_().(TerminalAction_TCommandFinished)
	return ok
}

type TerminalAction_TInput struct {
}

func (TerminalAction_TInput) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TInput_() TerminalAction {
	return TerminalAction{TerminalAction_TInput{}}
}

func (_this TerminalAction) Is_TInput() bool {
	_, ok := _this.Get_().(TerminalAction_TInput)
	return ok
}

type TerminalAction_TUnknown struct {
	Raw m_ConfluxCodec.Json
}

func (TerminalAction_TUnknown) isTerminalAction() {}

func (CompanionStruct_TerminalAction_) Create_TUnknown_(Raw m_ConfluxCodec.Json) TerminalAction {
	return TerminalAction{TerminalAction_TUnknown{Raw}}
}

func (_this TerminalAction) Is_TUnknown() bool {
	_, ok := _this.Get_().(TerminalAction_TUnknown)
	return ok
}

func (CompanionStruct_TerminalAction_) Default() TerminalAction {
	return Companion_TerminalAction_.Create_TCwdChanged_(_dafny.EmptySeq)
}

func (_this TerminalAction) Dtor_cwd() _dafny.Sequence {
	return _this.Get_().(TerminalAction_TCwdChanged).Cwd
}

func (_this TerminalAction) Dtor_title() _dafny.Sequence {
	return _this.Get_().(TerminalAction_TTitleChanged).Title
}

func (_this TerminalAction) Dtor_cols() _dafny.Int {
	return _this.Get_().(TerminalAction_TResized).Cols
}

func (_this TerminalAction) Dtor_rows() _dafny.Int {
	return _this.Get_().(TerminalAction_TResized).Rows
}

func (_this TerminalAction) Dtor_code() _dafny.Int {
	switch data := _this.Get_().(type) {
	case TerminalAction_TExited:
		return data.Code
	default:
		return data.(TerminalAction_TCommandFinished).Code
	}
}

func (_this TerminalAction) Dtor_data() _dafny.Sequence {
	return _this.Get_().(TerminalAction_TData).Data
}

func (_this TerminalAction) Dtor_claim() m_ConfluxCodec.Json {
	return _this.Get_().(TerminalAction_TClaimed).Claim
}

func (_this TerminalAction) Dtor_commandId() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case TerminalAction_TCommandExecuted:
		return data.CommandId
	default:
		return data.(TerminalAction_TCommandFinished).CommandId
	}
}

func (_this TerminalAction) Dtor_commandLine() _dafny.Sequence {
	return _this.Get_().(TerminalAction_TCommandExecuted).CommandLine
}

func (_this TerminalAction) Dtor_timestamp() _dafny.Int {
	return _this.Get_().(TerminalAction_TCommandExecuted).Timestamp
}

func (_this TerminalAction) Dtor_durationMs() _dafny.Int {
	return _this.Get_().(TerminalAction_TCommandFinished).DurationMs
}

func (_this TerminalAction) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(TerminalAction_TUnknown).Raw
}

func (_this TerminalAction) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case TerminalAction_TCwdChanged:
		{
			return "Terminal.TerminalAction.TCwdChanged" + "(" + data.Cwd.VerbatimString(true) + ")"
		}
	case TerminalAction_TTitleChanged:
		{
			return "Terminal.TerminalAction.TTitleChanged" + "(" + data.Title.VerbatimString(true) + ")"
		}
	case TerminalAction_TResized:
		{
			return "Terminal.TerminalAction.TResized" + "(" + _dafny.String(data.Cols) + ", " + _dafny.String(data.Rows) + ")"
		}
	case TerminalAction_TExited:
		{
			return "Terminal.TerminalAction.TExited" + "(" + _dafny.String(data.Code) + ")"
		}
	case TerminalAction_TData:
		{
			return "Terminal.TerminalAction.TData" + "(" + data.Data.VerbatimString(true) + ")"
		}
	case TerminalAction_TCleared:
		{
			return "Terminal.TerminalAction.TCleared"
		}
	case TerminalAction_TClaimed:
		{
			return "Terminal.TerminalAction.TClaimed" + "(" + _dafny.String(data.Claim) + ")"
		}
	case TerminalAction_TCommandDetectionAvailable:
		{
			return "Terminal.TerminalAction.TCommandDetectionAvailable"
		}
	case TerminalAction_TCommandExecuted:
		{
			return "Terminal.TerminalAction.TCommandExecuted" + "(" + data.CommandId.VerbatimString(true) + ", " + data.CommandLine.VerbatimString(true) + ", " + _dafny.String(data.Timestamp) + ")"
		}
	case TerminalAction_TCommandFinished:
		{
			return "Terminal.TerminalAction.TCommandFinished" + "(" + data.CommandId.VerbatimString(true) + ", " + _dafny.String(data.Code) + ", " + _dafny.String(data.DurationMs) + ")"
		}
	case TerminalAction_TInput:
		{
			return "Terminal.TerminalAction.TInput"
		}
	case TerminalAction_TUnknown:
		{
			return "Terminal.TerminalAction.TUnknown" + "(" + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this TerminalAction) Equals(other TerminalAction) bool {
	switch data1 := _this.Get_().(type) {
	case TerminalAction_TCwdChanged:
		{
			data2, ok := other.Get_().(TerminalAction_TCwdChanged)
			return ok && data1.Cwd.Equals(data2.Cwd)
		}
	case TerminalAction_TTitleChanged:
		{
			data2, ok := other.Get_().(TerminalAction_TTitleChanged)
			return ok && data1.Title.Equals(data2.Title)
		}
	case TerminalAction_TResized:
		{
			data2, ok := other.Get_().(TerminalAction_TResized)
			return ok && data1.Cols.Cmp(data2.Cols) == 0 && data1.Rows.Cmp(data2.Rows) == 0
		}
	case TerminalAction_TExited:
		{
			data2, ok := other.Get_().(TerminalAction_TExited)
			return ok && data1.Code.Cmp(data2.Code) == 0
		}
	case TerminalAction_TData:
		{
			data2, ok := other.Get_().(TerminalAction_TData)
			return ok && data1.Data.Equals(data2.Data)
		}
	case TerminalAction_TCleared:
		{
			_, ok := other.Get_().(TerminalAction_TCleared)
			return ok
		}
	case TerminalAction_TClaimed:
		{
			data2, ok := other.Get_().(TerminalAction_TClaimed)
			return ok && data1.Claim.Equals(data2.Claim)
		}
	case TerminalAction_TCommandDetectionAvailable:
		{
			_, ok := other.Get_().(TerminalAction_TCommandDetectionAvailable)
			return ok
		}
	case TerminalAction_TCommandExecuted:
		{
			data2, ok := other.Get_().(TerminalAction_TCommandExecuted)
			return ok && data1.CommandId.Equals(data2.CommandId) && data1.CommandLine.Equals(data2.CommandLine) && data1.Timestamp.Cmp(data2.Timestamp) == 0
		}
	case TerminalAction_TCommandFinished:
		{
			data2, ok := other.Get_().(TerminalAction_TCommandFinished)
			return ok && data1.CommandId.Equals(data2.CommandId) && data1.Code.Cmp(data2.Code) == 0 && data1.DurationMs.Cmp(data2.DurationMs) == 0
		}
	case TerminalAction_TInput:
		{
			_, ok := other.Get_().(TerminalAction_TInput)
			return ok
		}
	case TerminalAction_TUnknown:
		{
			data2, ok := other.Get_().(TerminalAction_TUnknown)
			return ok && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this TerminalAction) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(TerminalAction)
	return ok && _this.Equals(typed)
}

func Type_TerminalAction_() _dafny.TypeDescriptor {
	return type_TerminalAction_{}
}

type type_TerminalAction_ struct {
}

func (_this type_TerminalAction_) Default() interface{} {
	return Companion_TerminalAction_.Default()
}

func (_this type_TerminalAction_) String() string {
	return "Terminal.TerminalAction"
}
func (_this TerminalAction) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = TerminalAction{}

// End of datatype TerminalAction
