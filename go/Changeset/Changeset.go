// Package Changeset
// Dafny module Changeset compiled into Go

package Changeset

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_ConfluxCodec "github.com/joshmouch/ahp-go/ConfluxCodec"
	m_ConfluxContract "github.com/joshmouch/ahp-go/ConfluxContract"
	m_ConfluxOperators "github.com/joshmouch/ahp-go/ConfluxOperators"
	m_ConfluxPrune "github.com/joshmouch/ahp-go/ConfluxPrune"
	m_ConfluxSeqRoute "github.com/joshmouch/ahp-go/ConfluxSeqRoute"
	m_ResourceWatch "github.com/joshmouch/ahp-go/ResourceWatch"
	m__System "github.com/joshmouch/ahp-go/System_"
	_dafny "github.com/joshmouch/ahp-go/dafny"
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
	return "Changeset.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) OpSet(st _dafny.Sequence, err m_AhpSkeleton.Option) func(ChangesetOperation) ChangesetOperation {
	return (func(_0_err m_AhpSkeleton.Option, _1_st _dafny.Sequence) func(ChangesetOperation) ChangesetOperation {
		return func(_2_e ChangesetOperation) ChangesetOperation {
			return func(_pat_let26_0 ChangesetOperation) ChangesetOperation {
				return func(_3_dt__update__tmp_h0 ChangesetOperation) ChangesetOperation {
					return func(_pat_let27_0 m_AhpSkeleton.Option) ChangesetOperation {
						return func(_4_dt__update_herror_h0 m_AhpSkeleton.Option) ChangesetOperation {
							return func(_pat_let28_0 _dafny.Sequence) ChangesetOperation {
								return func(_5_dt__update_hstatus_h0 _dafny.Sequence) ChangesetOperation {
									return Companion_ChangesetOperation_.Create_ChangesetOperation_((_3_dt__update__tmp_h0).Dtor_id(), (_3_dt__update__tmp_h0).Dtor_label__(), (_3_dt__update__tmp_h0).Dtor_scopes(), _5_dt__update_hstatus_h0, _4_dt__update_herror_h0)
								}(_pat_let28_0)
							}(_1_st)
						}(_pat_let27_0)
					}(_0_err)
				}(_pat_let26_0)
			}(_2_e)
		}
	})(err, st)
}
func (_static *CompanionStruct_Default___) UpsertFile(files _dafny.Sequence, f ChangesetFile) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpsertById(func(coer8 func(ChangesetFile) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg12 interface{}) interface{} {
			return coer8(arg12.(ChangesetFile))
		}
	}(Companion_Default___.FileKey()), files, f)
}
func (_static *CompanionStruct_Default___) RemoveFile(files _dafny.Sequence, id _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqRemoveById(func(coer9 func(ChangesetFile) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg13 interface{}) interface{} {
			return coer9(arg13.(ChangesetFile))
		}
	}(Companion_Default___.FileKey()), files, id)
}
func (_static *CompanionStruct_Default___) SetReviewed(files _dafny.Sequence, ids _dafny.Sequence, rev bool) _dafny.Sequence {
	return m_ConfluxOperators.Companion_Default___.Map(func(coer10 func(ChangesetFile) ChangesetFile) func(interface{}) interface{} {
		return func(arg14 interface{}) interface{} {
			return coer10(arg14.(ChangesetFile))
		}
	}((func(_0_ids _dafny.Sequence, _1_rev bool) func(ChangesetFile) ChangesetFile {
		return func(_2_f ChangesetFile) ChangesetFile {
			return (func() ChangesetFile {
				if _dafny.Companion_Sequence_.Contains(_0_ids, (_2_f).Dtor_id()) {
					return func(_pat_let29_0 ChangesetFile) ChangesetFile {
						return func(_3_dt__update__tmp_h0 ChangesetFile) ChangesetFile {
							return func(_pat_let30_0 m_AhpSkeleton.Option) ChangesetFile {
								return func(_4_dt__update_hreviewed_h0 m_AhpSkeleton.Option) ChangesetFile {
									return Companion_ChangesetFile_.Create_ChangesetFile_((_3_dt__update__tmp_h0).Dtor_id(), _4_dt__update_hreviewed_h0, (_3_dt__update__tmp_h0).Dtor_edit())
								}(_pat_let30_0)
							}(m_AhpSkeleton.Companion_Option_.Create_Some_(_1_rev))
						}(_pat_let29_0)
					}(_2_f)
				}
				return _2_f
			})()
		}
	})(ids, rev)), files)
}
func (_static *CompanionStruct_Default___) UpdateOp(ops _dafny.Sequence, id _dafny.Sequence, status _dafny.Sequence, error_ m_AhpSkeleton.Option) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpdateById(func(coer11 func(ChangesetOperation) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg15 interface{}) interface{} {
			return coer11(arg15.(ChangesetOperation))
		}
	}(Companion_Default___.OpKey()), ops, id, func(coer12 func(ChangesetOperation) ChangesetOperation) func(interface{}) interface{} {
		return func(arg16 interface{}) interface{} {
			return coer12(arg16.(ChangesetOperation))
		}
	}(Companion_Default___.OpSet(status, error_)))
}
func (_static *CompanionStruct_Default___) HasOp(ops _dafny.Sequence, id _dafny.Sequence) bool {
	return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((ops).Cardinality())), false, func(_exists_var_0 _dafny.Int) bool {
		var _0_i _dafny.Int
		_0_i = interface{}(_exists_var_0).(_dafny.Int)
		return (((_0_i).Sign() != -1) && ((_0_i).Cmp(_dafny.IntOfUint32((ops).Cardinality())) < 0)) && (_dafny.Companion_Sequence_.Equal(((ops).Select((_0_i).Uint32()).(ChangesetOperation)).Dtor_id(), id))
	})
}
func (_static *CompanionStruct_Default___) AnyNeedsReviewed(files _dafny.Sequence, ids _dafny.Sequence, rev bool) bool {
	return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((files).Cardinality())), false, func(_exists_var_0 _dafny.Int) bool {
		var _0_i _dafny.Int
		_0_i = interface{}(_exists_var_0).(_dafny.Int)
		return ((((_0_i).Sign() != -1) && ((_0_i).Cmp(_dafny.IntOfUint32((files).Cardinality())) < 0)) && (_dafny.Companion_Sequence_.Contains(ids, ((files).Select((_0_i).Uint32()).(ChangesetFile)).Dtor_id()))) && (!(((files).Select((_0_i).Uint32()).(ChangesetFile)).Dtor_reviewed()).Equals(m_AhpSkeleton.Companion_Option_.Create_Some_(rev)))
	})
}
func (_static *CompanionStruct_Default___) ApplyToChangeset(s ChangesetState, a ChangesetAction, now _dafny.Int) _dafny.Tuple {
	var _pat_let_tv0 = s
	_ = _pat_let_tv0
	var _pat_let_tv1 = s
	_ = _pat_let_tv1
	var _pat_let_tv2 = s
	_ = _pat_let_tv2
	var _source0 ChangesetAction = a
	_ = _source0
	{
		if _source0.Is_StatusChanged() {
			var _0_st _dafny.Sequence = _source0.Get_().(ChangesetAction_StatusChanged).Status
			_ = _0_st
			var _1_err m_AhpSkeleton.Option = _source0.Get_().(ChangesetAction_StatusChanged).Error
			_ = _1_err
			return _dafny.TupleOf(func(_pat_let31_0 ChangesetState) ChangesetState {
				return func(_2_dt__update__tmp_h0 ChangesetState) ChangesetState {
					return func(_pat_let32_0 m_AhpSkeleton.Option) ChangesetState {
						return func(_3_dt__update_herror_h0 m_AhpSkeleton.Option) ChangesetState {
							return func(_pat_let33_0 _dafny.Sequence) ChangesetState {
								return func(_4_dt__update_hstatus_h0 _dafny.Sequence) ChangesetState {
									return Companion_ChangesetState_.Create_ChangesetState_(_4_dt__update_hstatus_h0, (_2_dt__update__tmp_h0).Dtor_files(), (_2_dt__update__tmp_h0).Dtor_operations(), _3_dt__update_herror_h0)
								}(_pat_let33_0)
							}(_0_st)
						}(_pat_let32_0)
					}(_1_err)
				}(_pat_let31_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_FileSet() {
			var _5_f ChangesetFile = _source0.Get_().(ChangesetAction_FileSet).File
			_ = _5_f
			return _dafny.TupleOf(func(_pat_let34_0 ChangesetState) ChangesetState {
				return func(_6_dt__update__tmp_h1 ChangesetState) ChangesetState {
					return func(_pat_let35_0 _dafny.Sequence) ChangesetState {
						return func(_7_dt__update_hfiles_h0 _dafny.Sequence) ChangesetState {
							return Companion_ChangesetState_.Create_ChangesetState_((_6_dt__update__tmp_h1).Dtor_status(), _7_dt__update_hfiles_h0, (_6_dt__update__tmp_h1).Dtor_operations(), (_6_dt__update__tmp_h1).Dtor_error())
						}(_pat_let35_0)
					}(Companion_Default___.UpsertFile((_pat_let_tv0).Dtor_files(), _5_f))
				}(_pat_let34_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_FileRemoved() {
			var _8_id _dafny.Sequence = _source0.Get_().(ChangesetAction_FileRemoved).FileId
			_ = _8_id
			return _dafny.TupleOf(func(_pat_let36_0 ChangesetState) ChangesetState {
				return func(_9_dt__update__tmp_h2 ChangesetState) ChangesetState {
					return func(_pat_let37_0 _dafny.Sequence) ChangesetState {
						return func(_10_dt__update_hfiles_h1 _dafny.Sequence) ChangesetState {
							return Companion_ChangesetState_.Create_ChangesetState_((_9_dt__update__tmp_h2).Dtor_status(), _10_dt__update_hfiles_h1, (_9_dt__update__tmp_h2).Dtor_operations(), (_9_dt__update__tmp_h2).Dtor_error())
						}(_pat_let37_0)
					}(Companion_Default___.RemoveFile((_pat_let_tv1).Dtor_files(), _8_id))
				}(_pat_let36_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_OperationsChanged() {
			var _11_ops m_AhpSkeleton.Option = _source0.Get_().(ChangesetAction_OperationsChanged).Operations
			_ = _11_ops
			return _dafny.TupleOf(func(_pat_let38_0 ChangesetState) ChangesetState {
				return func(_12_dt__update__tmp_h3 ChangesetState) ChangesetState {
					return func(_pat_let39_0 m_AhpSkeleton.Option) ChangesetState {
						return func(_13_dt__update_hoperations_h0 m_AhpSkeleton.Option) ChangesetState {
							return Companion_ChangesetState_.Create_ChangesetState_((_12_dt__update__tmp_h3).Dtor_status(), (_12_dt__update__tmp_h3).Dtor_files(), _13_dt__update_hoperations_h0, (_12_dt__update__tmp_h3).Dtor_error())
						}(_pat_let39_0)
					}(_11_ops)
				}(_pat_let38_0)
			}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_Cleared() {
			return _dafny.TupleOf(func(_pat_let40_0 ChangesetState) ChangesetState {
				return func(_14_dt__update__tmp_h4 ChangesetState) ChangesetState {
					return func(_pat_let41_0 _dafny.Sequence) ChangesetState {
						return func(_15_dt__update_hfiles_h2 _dafny.Sequence) ChangesetState {
							return Companion_ChangesetState_.Create_ChangesetState_((_14_dt__update__tmp_h4).Dtor_status(), _15_dt__update_hfiles_h2, (_14_dt__update__tmp_h4).Dtor_operations(), (_14_dt__update__tmp_h4).Dtor_error())
						}(_pat_let41_0)
					}(_dafny.SeqOf())
				}(_pat_let40_0)
			}(s), (func() m_AhpSkeleton.ReduceOutcome {
				if (_dafny.IntOfUint32(((s).Dtor_files()).Cardinality())).Sign() == 0 {
					return m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_()
				}
				return m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_()
			})())
		}
	}
	{
		if _source0.Is_OperationStatusChanged() {
			var _16_id _dafny.Sequence = _source0.Get_().(ChangesetAction_OperationStatusChanged).OperationId
			_ = _16_id
			var _17_st _dafny.Sequence = _source0.Get_().(ChangesetAction_OperationStatusChanged).Status
			_ = _17_st
			var _18_err m_AhpSkeleton.Option = _source0.Get_().(ChangesetAction_OperationStatusChanged).Error
			_ = _18_err
			var _source1 m_AhpSkeleton.Option = (s).Dtor_operations()
			_ = _source1
			{
				if _source1.Is_None() {
					return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
				}
			}
			{
				var _19_ops _dafny.Sequence = _source1.Get_().(m_AhpSkeleton.Option_Some).Value.(_dafny.Sequence)
				_ = _19_ops
				if Companion_Default___.HasOp(_19_ops, _16_id) {
					return _dafny.TupleOf(func(_pat_let42_0 ChangesetState) ChangesetState {
						return func(_20_dt__update__tmp_h5 ChangesetState) ChangesetState {
							return func(_pat_let43_0 m_AhpSkeleton.Option) ChangesetState {
								return func(_21_dt__update_hoperations_h1 m_AhpSkeleton.Option) ChangesetState {
									return Companion_ChangesetState_.Create_ChangesetState_((_20_dt__update__tmp_h5).Dtor_status(), (_20_dt__update__tmp_h5).Dtor_files(), _21_dt__update_hoperations_h1, (_20_dt__update__tmp_h5).Dtor_error())
								}(_pat_let43_0)
							}(m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Default___.UpdateOp(_19_ops, _16_id, _17_st, _18_err)))
						}(_pat_let42_0)
					}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
				} else {
					return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
				}
			}
		}
	}
	{
		if _source0.Is_ContentChanged() {
			var _22_fs m_AhpSkeleton.Option = _source0.Get_().(ChangesetAction_ContentChanged).Files
			_ = _22_fs
			var _23_ops m_AhpSkeleton.Option = _source0.Get_().(ChangesetAction_ContentChanged).Operations
			_ = _23_ops
			var _24_err m_AhpSkeleton.Option = _source0.Get_().(ChangesetAction_ContentChanged).Error
			_ = _24_err
			var _25_s1 ChangesetState = (func() ChangesetState {
				if (_22_fs).Is_Some() {
					return func(_pat_let44_0 ChangesetState) ChangesetState {
						return func(_26_dt__update__tmp_h6 ChangesetState) ChangesetState {
							return func(_pat_let45_0 _dafny.Sequence) ChangesetState {
								return func(_27_dt__update_hfiles_h3 _dafny.Sequence) ChangesetState {
									return Companion_ChangesetState_.Create_ChangesetState_((_26_dt__update__tmp_h6).Dtor_status(), _27_dt__update_hfiles_h3, (_26_dt__update__tmp_h6).Dtor_operations(), (_26_dt__update__tmp_h6).Dtor_error())
								}(_pat_let45_0)
							}((_22_fs).Dtor_value().(_dafny.Sequence))
						}(_pat_let44_0)
					}(s)
				}
				return s
			})()
			_ = _25_s1
			var _28_s2 ChangesetState = (func() ChangesetState {
				if (_23_ops).Is_Some() {
					return func(_pat_let46_0 ChangesetState) ChangesetState {
						return func(_29_dt__update__tmp_h7 ChangesetState) ChangesetState {
							return func(_pat_let47_0 m_AhpSkeleton.Option) ChangesetState {
								return func(_30_dt__update_hoperations_h2 m_AhpSkeleton.Option) ChangesetState {
									return Companion_ChangesetState_.Create_ChangesetState_((_29_dt__update__tmp_h7).Dtor_status(), (_29_dt__update__tmp_h7).Dtor_files(), _30_dt__update_hoperations_h2, (_29_dt__update__tmp_h7).Dtor_error())
								}(_pat_let47_0)
							}(_23_ops)
						}(_pat_let46_0)
					}(_25_s1)
				}
				return _25_s1
			})()
			_ = _28_s2
			var _31_s3 ChangesetState = (func() ChangesetState {
				if (_24_err).Is_Some() {
					return func(_pat_let48_0 ChangesetState) ChangesetState {
						return func(_32_dt__update__tmp_h8 ChangesetState) ChangesetState {
							return func(_pat_let49_0 m_AhpSkeleton.Option) ChangesetState {
								return func(_33_dt__update_herror_h1 m_AhpSkeleton.Option) ChangesetState {
									return Companion_ChangesetState_.Create_ChangesetState_((_32_dt__update__tmp_h8).Dtor_status(), (_32_dt__update__tmp_h8).Dtor_files(), (_32_dt__update__tmp_h8).Dtor_operations(), _33_dt__update_herror_h1)
								}(_pat_let49_0)
							}(_24_err)
						}(_pat_let48_0)
					}(_28_s2)
				}
				return _28_s2
			})()
			_ = _31_s3
			return _dafny.TupleOf(_31_s3, m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_FilesReviewedChanged() {
			var _34_ids _dafny.Sequence = _source0.Get_().(ChangesetAction_FilesReviewedChanged).FileIds
			_ = _34_ids
			var _35_rev bool = _source0.Get_().(ChangesetAction_FilesReviewedChanged).Reviewed
			_ = _35_rev
			if Companion_Default___.AnyNeedsReviewed((s).Dtor_files(), _34_ids, _35_rev) {
				return _dafny.TupleOf(func(_pat_let50_0 ChangesetState) ChangesetState {
					return func(_36_dt__update__tmp_h9 ChangesetState) ChangesetState {
						return func(_pat_let51_0 _dafny.Sequence) ChangesetState {
							return func(_37_dt__update_hfiles_h4 _dafny.Sequence) ChangesetState {
								return Companion_ChangesetState_.Create_ChangesetState_((_36_dt__update__tmp_h9).Dtor_status(), _37_dt__update_hfiles_h4, (_36_dt__update__tmp_h9).Dtor_operations(), (_36_dt__update__tmp_h9).Dtor_error())
							}(_pat_let51_0)
						}(Companion_Default___.SetReviewed((_pat_let_tv2).Dtor_files(), _34_ids, _35_rev))
					}(_pat_let50_0)
				}(s), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
	}
}
func (_static *CompanionStruct_Default___) Apply1(s ChangesetState, a ChangesetAction) ChangesetState {
	return (*(Companion_Default___.ApplyToChangeset(s, a, _dafny.IntOfInt64(9999))).IndexInt(0)).(ChangesetState)
}
func (_static *CompanionStruct_Default___) Fold(s ChangesetState, acts _dafny.Sequence) ChangesetState {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer13 func(ChangesetState, ChangesetAction) ChangesetState) func(interface{}, interface{}) interface{} {
		return func(arg17 interface{}, arg18 interface{}) interface{} {
			return coer13(arg17.(ChangesetState), arg18.(ChangesetAction))
		}
	}(Companion_Default___.Apply1), s, acts).(ChangesetState)
}
func (_static *CompanionStruct_Default___) E() m_ConfluxCodec.Json {
	return m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("diff"), m_ConfluxCodec.Companion_Json_.Create_JNull_()))
}
func (_static *CompanionStruct_Default___) RunCorpus() (_dafny.Int, _dafny.Int) {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	var total _dafny.Int = _dafny.Zero
	_ = total
	total = _dafny.IntOfInt64(15)
	pass = _dafny.Zero
	var _0_fa ChangesetFile
	_ = _0_fa
	_0_fa = Companion_ChangesetFile_.Create_ChangesetFile_(_dafny.UnicodeSeqOfUtf8Bytes("file:///src/a.ts"), m_AhpSkeleton.Companion_Option_.Create_None_(), Companion_Default___.E())
	var _1_fb ChangesetFile
	_ = _1_fb
	_1_fb = Companion_ChangesetFile_.Create_ChangesetFile_(_dafny.UnicodeSeqOfUtf8Bytes("file:///src/b.ts"), m_AhpSkeleton.Companion_Option_.Create_None_(), Companion_Default___.E())
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("computing"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("boom")))), Companion_ChangesetAction_.Create_StatusChanged_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_StatusChanged_(_dafny.UnicodeSeqOfUtf8Bytes("error"), m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("x")))))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("error"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("x"))))) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_0_fa), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_FileSet_(_1_fb))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_0_fa, _1_fb), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	var _2_fbv2 ChangesetFile
	_ = _2_fbv2
	_2_fbv2 = Companion_ChangesetFile_.Create_ChangesetFile_(_dafny.UnicodeSeqOfUtf8Bytes("file:///src/b.ts"), m_AhpSkeleton.Companion_Option_.Create_Some_(true), Companion_Default___.E())
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_0_fa, _1_fb), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_FileSet_(_2_fbv2))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_0_fa, _2_fbv2), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_0_fa, _1_fb), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_FileRemoved_(_dafny.UnicodeSeqOfUtf8Bytes("file:///src/a.ts")))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_1_fb), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_0_fa), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_FileRemoved_(_dafny.UnicodeSeqOfUtf8Bytes("nope")))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_0_fa), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	var _3_op ChangesetOperation
	_ = _3_op
	_3_op = Companion_ChangesetOperation_.Create_ChangesetOperation_(_dafny.UnicodeSeqOfUtf8Bytes("create-pr"), _dafny.UnicodeSeqOfUtf8Bytes("Create Pull Request"), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("changeset")), _dafny.UnicodeSeqOfUtf8Bytes("idle"), m_AhpSkeleton.Companion_Option_.Create_None_())
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_OperationsChanged_(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_3_op))))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_3_op)), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_3_op)), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_OperationsChanged_(m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	if _dafny.Companion_Sequence_.Equal((Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_0_fa), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_Cleared_())).Dtor_files(), _dafny.SeqOf()) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_CsUnknown_(m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("type"), m_ConfluxCodec.Companion_Json_.Create_JStr_(_dafny.UnicodeSeqOfUtf8Bytes("changeset/nope"))))))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	var _4_opR ChangesetOperation
	_ = _4_opR
	var _5_dt__update__tmp_h0 ChangesetOperation = _3_op
	_ = _5_dt__update__tmp_h0
	var _6_dt__update_hstatus_h0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("running")
	_ = _6_dt__update_hstatus_h0
	_4_opR = Companion_ChangesetOperation_.Create_ChangesetOperation_((_5_dt__update__tmp_h0).Dtor_id(), (_5_dt__update__tmp_h0).Dtor_label__(), (_5_dt__update__tmp_h0).Dtor_scopes(), _6_dt__update_hstatus_h0, (_5_dt__update__tmp_h0).Dtor_error())
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_3_op)), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_OperationStatusChanged_(_dafny.UnicodeSeqOfUtf8Bytes("create-pr"), _dafny.UnicodeSeqOfUtf8Bytes("running"), m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_4_opR)), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_3_op)), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_OperationStatusChanged_(_dafny.UnicodeSeqOfUtf8Bytes("nope"), _dafny.UnicodeSeqOfUtf8Bytes("running"), m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_3_op)), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_0_fa), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_ContentChanged_(m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_1_fb)), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_3_op)), m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_1_fb), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.SeqOf(_3_op)), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	var _7_fc ChangesetFile
	_ = _7_fc
	_7_fc = Companion_ChangesetFile_.Create_ChangesetFile_(_dafny.UnicodeSeqOfUtf8Bytes("file:///src/c.ts"), m_AhpSkeleton.Companion_Option_.Create_Some_(false), Companion_Default___.E())
	var _8_fbT ChangesetFile
	_ = _8_fbT
	_8_fbT = Companion_ChangesetFile_.Create_ChangesetFile_(_dafny.UnicodeSeqOfUtf8Bytes("file:///src/b.ts"), m_AhpSkeleton.Companion_Option_.Create_Some_(true), Companion_Default___.E())
	var _9_faT ChangesetFile
	_ = _9_faT
	_9_faT = Companion_ChangesetFile_.Create_ChangesetFile_(_dafny.UnicodeSeqOfUtf8Bytes("file:///src/a.ts"), m_AhpSkeleton.Companion_Option_.Create_Some_(true), Companion_Default___.E())
	var _10_r160 ChangesetState
	_ = _10_r160
	_10_r160 = Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_0_fa, _8_fbT, _7_fc), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_FilesReviewedChanged_(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("file:///src/a.ts"), _dafny.UnicodeSeqOfUtf8Bytes("file:///src/b.ts"), _dafny.UnicodeSeqOfUtf8Bytes("file:///src/missing.ts")), true))
	if (_10_r160).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_9_faT, _8_fbT, _7_fc), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_8_fbT), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_()), Companion_ChangesetAction_.Create_FilesReviewedChanged_(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("file:///src/b.ts")), true))).Equals(Companion_ChangesetState_.Create_ChangesetState_(_dafny.UnicodeSeqOfUtf8Bytes("ready"), _dafny.SeqOf(_8_fbT), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_())) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass, total
}
func (_static *CompanionStruct_Default___) FileKey() func(ChangesetFile) _dafny.Sequence {
	return func(_0_x ChangesetFile) _dafny.Sequence {
		return (_0_x).Dtor_id()
	}
}
func (_static *CompanionStruct_Default___) OpKey() func(ChangesetOperation) _dafny.Sequence {
	return func(_0_x ChangesetOperation) _dafny.Sequence {
		return (_0_x).Dtor_id()
	}
}

// End of class Default__

// Definition of datatype ChangesetFile
type ChangesetFile struct {
	Data_ChangesetFile_
}

func (_this ChangesetFile) Get_() Data_ChangesetFile_ {
	return _this.Data_ChangesetFile_
}

type Data_ChangesetFile_ interface {
	isChangesetFile()
}

type CompanionStruct_ChangesetFile_ struct {
}

var Companion_ChangesetFile_ = CompanionStruct_ChangesetFile_{}

type ChangesetFile_ChangesetFile struct {
	Id       _dafny.Sequence
	Reviewed m_AhpSkeleton.Option
	Edit     m_ConfluxCodec.Json
}

func (ChangesetFile_ChangesetFile) isChangesetFile() {}

func (CompanionStruct_ChangesetFile_) Create_ChangesetFile_(Id _dafny.Sequence, Reviewed m_AhpSkeleton.Option, Edit m_ConfluxCodec.Json) ChangesetFile {
	return ChangesetFile{ChangesetFile_ChangesetFile{Id, Reviewed, Edit}}
}

func (_this ChangesetFile) Is_ChangesetFile() bool {
	_, ok := _this.Get_().(ChangesetFile_ChangesetFile)
	return ok
}

func (CompanionStruct_ChangesetFile_) Default() ChangesetFile {
	return Companion_ChangesetFile_.Create_ChangesetFile_(_dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), m_ConfluxCodec.Companion_Json_.Default())
}

func (_this ChangesetFile) Dtor_id() _dafny.Sequence {
	return _this.Get_().(ChangesetFile_ChangesetFile).Id
}

func (_this ChangesetFile) Dtor_reviewed() m_AhpSkeleton.Option {
	return _this.Get_().(ChangesetFile_ChangesetFile).Reviewed
}

func (_this ChangesetFile) Dtor_edit() m_ConfluxCodec.Json {
	return _this.Get_().(ChangesetFile_ChangesetFile).Edit
}

func (_this ChangesetFile) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ChangesetFile_ChangesetFile:
		{
			return "Changeset.ChangesetFile.ChangesetFile" + "(" + data.Id.VerbatimString(true) + ", " + _dafny.String(data.Reviewed) + ", " + _dafny.String(data.Edit) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ChangesetFile) Equals(other ChangesetFile) bool {
	switch data1 := _this.Get_().(type) {
	case ChangesetFile_ChangesetFile:
		{
			data2, ok := other.Get_().(ChangesetFile_ChangesetFile)
			return ok && data1.Id.Equals(data2.Id) && data1.Reviewed.Equals(data2.Reviewed) && data1.Edit.Equals(data2.Edit)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ChangesetFile) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ChangesetFile)
	return ok && _this.Equals(typed)
}

func Type_ChangesetFile_() _dafny.TypeDescriptor {
	return type_ChangesetFile_{}
}

type type_ChangesetFile_ struct {
}

func (_this type_ChangesetFile_) Default() interface{} {
	return Companion_ChangesetFile_.Default()
}

func (_this type_ChangesetFile_) String() string {
	return "Changeset.ChangesetFile"
}
func (_this ChangesetFile) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ChangesetFile{}

// End of datatype ChangesetFile

// Definition of datatype ChangesetOperation
type ChangesetOperation struct {
	Data_ChangesetOperation_
}

func (_this ChangesetOperation) Get_() Data_ChangesetOperation_ {
	return _this.Data_ChangesetOperation_
}

type Data_ChangesetOperation_ interface {
	isChangesetOperation()
}

type CompanionStruct_ChangesetOperation_ struct {
}

var Companion_ChangesetOperation_ = CompanionStruct_ChangesetOperation_{}

type ChangesetOperation_ChangesetOperation struct {
	Id      _dafny.Sequence
	Label__ _dafny.Sequence
	Scopes  _dafny.Sequence
	Status  _dafny.Sequence
	Error   m_AhpSkeleton.Option
}

func (ChangesetOperation_ChangesetOperation) isChangesetOperation() {}

func (CompanionStruct_ChangesetOperation_) Create_ChangesetOperation_(Id _dafny.Sequence, Label__ _dafny.Sequence, Scopes _dafny.Sequence, Status _dafny.Sequence, Error m_AhpSkeleton.Option) ChangesetOperation {
	return ChangesetOperation{ChangesetOperation_ChangesetOperation{Id, Label__, Scopes, Status, Error}}
}

func (_this ChangesetOperation) Is_ChangesetOperation() bool {
	_, ok := _this.Get_().(ChangesetOperation_ChangesetOperation)
	return ok
}

func (CompanionStruct_ChangesetOperation_) Default() ChangesetOperation {
	return Companion_ChangesetOperation_.Create_ChangesetOperation_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default())
}

func (_this ChangesetOperation) Dtor_id() _dafny.Sequence {
	return _this.Get_().(ChangesetOperation_ChangesetOperation).Id
}

func (_this ChangesetOperation) Dtor_label__() _dafny.Sequence {
	return _this.Get_().(ChangesetOperation_ChangesetOperation).Label__
}

func (_this ChangesetOperation) Dtor_scopes() _dafny.Sequence {
	return _this.Get_().(ChangesetOperation_ChangesetOperation).Scopes
}

func (_this ChangesetOperation) Dtor_status() _dafny.Sequence {
	return _this.Get_().(ChangesetOperation_ChangesetOperation).Status
}

func (_this ChangesetOperation) Dtor_error() m_AhpSkeleton.Option {
	return _this.Get_().(ChangesetOperation_ChangesetOperation).Error
}

func (_this ChangesetOperation) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ChangesetOperation_ChangesetOperation:
		{
			return "Changeset.ChangesetOperation.ChangesetOperation" + "(" + data.Id.VerbatimString(true) + ", " + data.Label__.VerbatimString(true) + ", " + _dafny.String(data.Scopes) + ", " + data.Status.VerbatimString(true) + ", " + _dafny.String(data.Error) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ChangesetOperation) Equals(other ChangesetOperation) bool {
	switch data1 := _this.Get_().(type) {
	case ChangesetOperation_ChangesetOperation:
		{
			data2, ok := other.Get_().(ChangesetOperation_ChangesetOperation)
			return ok && data1.Id.Equals(data2.Id) && data1.Label__.Equals(data2.Label__) && data1.Scopes.Equals(data2.Scopes) && data1.Status.Equals(data2.Status) && data1.Error.Equals(data2.Error)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ChangesetOperation) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ChangesetOperation)
	return ok && _this.Equals(typed)
}

func Type_ChangesetOperation_() _dafny.TypeDescriptor {
	return type_ChangesetOperation_{}
}

type type_ChangesetOperation_ struct {
}

func (_this type_ChangesetOperation_) Default() interface{} {
	return Companion_ChangesetOperation_.Default()
}

func (_this type_ChangesetOperation_) String() string {
	return "Changeset.ChangesetOperation"
}
func (_this ChangesetOperation) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ChangesetOperation{}

// End of datatype ChangesetOperation

// Definition of datatype ChangesetState
type ChangesetState struct {
	Data_ChangesetState_
}

func (_this ChangesetState) Get_() Data_ChangesetState_ {
	return _this.Data_ChangesetState_
}

type Data_ChangesetState_ interface {
	isChangesetState()
}

type CompanionStruct_ChangesetState_ struct {
}

var Companion_ChangesetState_ = CompanionStruct_ChangesetState_{}

type ChangesetState_ChangesetState struct {
	Status     _dafny.Sequence
	Files      _dafny.Sequence
	Operations m_AhpSkeleton.Option
	Error      m_AhpSkeleton.Option
}

func (ChangesetState_ChangesetState) isChangesetState() {}

func (CompanionStruct_ChangesetState_) Create_ChangesetState_(Status _dafny.Sequence, Files _dafny.Sequence, Operations m_AhpSkeleton.Option, Error m_AhpSkeleton.Option) ChangesetState {
	return ChangesetState{ChangesetState_ChangesetState{Status, Files, Operations, Error}}
}

func (_this ChangesetState) Is_ChangesetState() bool {
	_, ok := _this.Get_().(ChangesetState_ChangesetState)
	return ok
}

func (CompanionStruct_ChangesetState_) Default() ChangesetState {
	return Companion_ChangesetState_.Create_ChangesetState_(_dafny.EmptySeq, _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), m_AhpSkeleton.Companion_Option_.Default())
}

func (_this ChangesetState) Dtor_status() _dafny.Sequence {
	return _this.Get_().(ChangesetState_ChangesetState).Status
}

func (_this ChangesetState) Dtor_files() _dafny.Sequence {
	return _this.Get_().(ChangesetState_ChangesetState).Files
}

func (_this ChangesetState) Dtor_operations() m_AhpSkeleton.Option {
	return _this.Get_().(ChangesetState_ChangesetState).Operations
}

func (_this ChangesetState) Dtor_error() m_AhpSkeleton.Option {
	return _this.Get_().(ChangesetState_ChangesetState).Error
}

func (_this ChangesetState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ChangesetState_ChangesetState:
		{
			return "Changeset.ChangesetState.ChangesetState" + "(" + data.Status.VerbatimString(true) + ", " + _dafny.String(data.Files) + ", " + _dafny.String(data.Operations) + ", " + _dafny.String(data.Error) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ChangesetState) Equals(other ChangesetState) bool {
	switch data1 := _this.Get_().(type) {
	case ChangesetState_ChangesetState:
		{
			data2, ok := other.Get_().(ChangesetState_ChangesetState)
			return ok && data1.Status.Equals(data2.Status) && data1.Files.Equals(data2.Files) && data1.Operations.Equals(data2.Operations) && data1.Error.Equals(data2.Error)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ChangesetState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ChangesetState)
	return ok && _this.Equals(typed)
}

func Type_ChangesetState_() _dafny.TypeDescriptor {
	return type_ChangesetState_{}
}

type type_ChangesetState_ struct {
}

func (_this type_ChangesetState_) Default() interface{} {
	return Companion_ChangesetState_.Default()
}

func (_this type_ChangesetState_) String() string {
	return "Changeset.ChangesetState"
}
func (_this ChangesetState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ChangesetState{}

// End of datatype ChangesetState

// Definition of datatype ChangesetAction
type ChangesetAction struct {
	Data_ChangesetAction_
}

func (_this ChangesetAction) Get_() Data_ChangesetAction_ {
	return _this.Data_ChangesetAction_
}

type Data_ChangesetAction_ interface {
	isChangesetAction()
}

type CompanionStruct_ChangesetAction_ struct {
}

var Companion_ChangesetAction_ = CompanionStruct_ChangesetAction_{}

type ChangesetAction_StatusChanged struct {
	Status _dafny.Sequence
	Error  m_AhpSkeleton.Option
}

func (ChangesetAction_StatusChanged) isChangesetAction() {}

func (CompanionStruct_ChangesetAction_) Create_StatusChanged_(Status _dafny.Sequence, Error m_AhpSkeleton.Option) ChangesetAction {
	return ChangesetAction{ChangesetAction_StatusChanged{Status, Error}}
}

func (_this ChangesetAction) Is_StatusChanged() bool {
	_, ok := _this.Get_().(ChangesetAction_StatusChanged)
	return ok
}

type ChangesetAction_FileSet struct {
	File ChangesetFile
}

func (ChangesetAction_FileSet) isChangesetAction() {}

func (CompanionStruct_ChangesetAction_) Create_FileSet_(File ChangesetFile) ChangesetAction {
	return ChangesetAction{ChangesetAction_FileSet{File}}
}

func (_this ChangesetAction) Is_FileSet() bool {
	_, ok := _this.Get_().(ChangesetAction_FileSet)
	return ok
}

type ChangesetAction_FileRemoved struct {
	FileId _dafny.Sequence
}

func (ChangesetAction_FileRemoved) isChangesetAction() {}

func (CompanionStruct_ChangesetAction_) Create_FileRemoved_(FileId _dafny.Sequence) ChangesetAction {
	return ChangesetAction{ChangesetAction_FileRemoved{FileId}}
}

func (_this ChangesetAction) Is_FileRemoved() bool {
	_, ok := _this.Get_().(ChangesetAction_FileRemoved)
	return ok
}

type ChangesetAction_OperationsChanged struct {
	Operations m_AhpSkeleton.Option
}

func (ChangesetAction_OperationsChanged) isChangesetAction() {}

func (CompanionStruct_ChangesetAction_) Create_OperationsChanged_(Operations m_AhpSkeleton.Option) ChangesetAction {
	return ChangesetAction{ChangesetAction_OperationsChanged{Operations}}
}

func (_this ChangesetAction) Is_OperationsChanged() bool {
	_, ok := _this.Get_().(ChangesetAction_OperationsChanged)
	return ok
}

type ChangesetAction_Cleared struct {
}

func (ChangesetAction_Cleared) isChangesetAction() {}

func (CompanionStruct_ChangesetAction_) Create_Cleared_() ChangesetAction {
	return ChangesetAction{ChangesetAction_Cleared{}}
}

func (_this ChangesetAction) Is_Cleared() bool {
	_, ok := _this.Get_().(ChangesetAction_Cleared)
	return ok
}

type ChangesetAction_OperationStatusChanged struct {
	OperationId _dafny.Sequence
	Status      _dafny.Sequence
	Error       m_AhpSkeleton.Option
}

func (ChangesetAction_OperationStatusChanged) isChangesetAction() {}

func (CompanionStruct_ChangesetAction_) Create_OperationStatusChanged_(OperationId _dafny.Sequence, Status _dafny.Sequence, Error m_AhpSkeleton.Option) ChangesetAction {
	return ChangesetAction{ChangesetAction_OperationStatusChanged{OperationId, Status, Error}}
}

func (_this ChangesetAction) Is_OperationStatusChanged() bool {
	_, ok := _this.Get_().(ChangesetAction_OperationStatusChanged)
	return ok
}

type ChangesetAction_ContentChanged struct {
	Files      m_AhpSkeleton.Option
	Operations m_AhpSkeleton.Option
	Error      m_AhpSkeleton.Option
}

func (ChangesetAction_ContentChanged) isChangesetAction() {}

func (CompanionStruct_ChangesetAction_) Create_ContentChanged_(Files m_AhpSkeleton.Option, Operations m_AhpSkeleton.Option, Error m_AhpSkeleton.Option) ChangesetAction {
	return ChangesetAction{ChangesetAction_ContentChanged{Files, Operations, Error}}
}

func (_this ChangesetAction) Is_ContentChanged() bool {
	_, ok := _this.Get_().(ChangesetAction_ContentChanged)
	return ok
}

type ChangesetAction_FilesReviewedChanged struct {
	FileIds  _dafny.Sequence
	Reviewed bool
}

func (ChangesetAction_FilesReviewedChanged) isChangesetAction() {}

func (CompanionStruct_ChangesetAction_) Create_FilesReviewedChanged_(FileIds _dafny.Sequence, Reviewed bool) ChangesetAction {
	return ChangesetAction{ChangesetAction_FilesReviewedChanged{FileIds, Reviewed}}
}

func (_this ChangesetAction) Is_FilesReviewedChanged() bool {
	_, ok := _this.Get_().(ChangesetAction_FilesReviewedChanged)
	return ok
}

type ChangesetAction_CsUnknown struct {
	Raw m_ConfluxCodec.Json
}

func (ChangesetAction_CsUnknown) isChangesetAction() {}

func (CompanionStruct_ChangesetAction_) Create_CsUnknown_(Raw m_ConfluxCodec.Json) ChangesetAction {
	return ChangesetAction{ChangesetAction_CsUnknown{Raw}}
}

func (_this ChangesetAction) Is_CsUnknown() bool {
	_, ok := _this.Get_().(ChangesetAction_CsUnknown)
	return ok
}

func (CompanionStruct_ChangesetAction_) Default() ChangesetAction {
	return Companion_ChangesetAction_.Create_StatusChanged_(_dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default())
}

func (_this ChangesetAction) Dtor_status() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case ChangesetAction_StatusChanged:
		return data.Status
	default:
		return data.(ChangesetAction_OperationStatusChanged).Status
	}
}

func (_this ChangesetAction) Dtor_error() m_AhpSkeleton.Option {
	switch data := _this.Get_().(type) {
	case ChangesetAction_StatusChanged:
		return data.Error
	case ChangesetAction_OperationStatusChanged:
		return data.Error
	default:
		return data.(ChangesetAction_ContentChanged).Error
	}
}

func (_this ChangesetAction) Dtor_file() ChangesetFile {
	return _this.Get_().(ChangesetAction_FileSet).File
}

func (_this ChangesetAction) Dtor_fileId() _dafny.Sequence {
	return _this.Get_().(ChangesetAction_FileRemoved).FileId
}

func (_this ChangesetAction) Dtor_operations() m_AhpSkeleton.Option {
	switch data := _this.Get_().(type) {
	case ChangesetAction_OperationsChanged:
		return data.Operations
	default:
		return data.(ChangesetAction_ContentChanged).Operations
	}
}

func (_this ChangesetAction) Dtor_operationId() _dafny.Sequence {
	return _this.Get_().(ChangesetAction_OperationStatusChanged).OperationId
}

func (_this ChangesetAction) Dtor_files() m_AhpSkeleton.Option {
	return _this.Get_().(ChangesetAction_ContentChanged).Files
}

func (_this ChangesetAction) Dtor_fileIds() _dafny.Sequence {
	return _this.Get_().(ChangesetAction_FilesReviewedChanged).FileIds
}

func (_this ChangesetAction) Dtor_reviewed() bool {
	return _this.Get_().(ChangesetAction_FilesReviewedChanged).Reviewed
}

func (_this ChangesetAction) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(ChangesetAction_CsUnknown).Raw
}

func (_this ChangesetAction) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ChangesetAction_StatusChanged:
		{
			return "Changeset.ChangesetAction.StatusChanged" + "(" + data.Status.VerbatimString(true) + ", " + _dafny.String(data.Error) + ")"
		}
	case ChangesetAction_FileSet:
		{
			return "Changeset.ChangesetAction.FileSet" + "(" + _dafny.String(data.File) + ")"
		}
	case ChangesetAction_FileRemoved:
		{
			return "Changeset.ChangesetAction.FileRemoved" + "(" + data.FileId.VerbatimString(true) + ")"
		}
	case ChangesetAction_OperationsChanged:
		{
			return "Changeset.ChangesetAction.OperationsChanged" + "(" + _dafny.String(data.Operations) + ")"
		}
	case ChangesetAction_Cleared:
		{
			return "Changeset.ChangesetAction.Cleared"
		}
	case ChangesetAction_OperationStatusChanged:
		{
			return "Changeset.ChangesetAction.OperationStatusChanged" + "(" + data.OperationId.VerbatimString(true) + ", " + data.Status.VerbatimString(true) + ", " + _dafny.String(data.Error) + ")"
		}
	case ChangesetAction_ContentChanged:
		{
			return "Changeset.ChangesetAction.ContentChanged" + "(" + _dafny.String(data.Files) + ", " + _dafny.String(data.Operations) + ", " + _dafny.String(data.Error) + ")"
		}
	case ChangesetAction_FilesReviewedChanged:
		{
			return "Changeset.ChangesetAction.FilesReviewedChanged" + "(" + _dafny.String(data.FileIds) + ", " + _dafny.String(data.Reviewed) + ")"
		}
	case ChangesetAction_CsUnknown:
		{
			return "Changeset.ChangesetAction.CsUnknown" + "(" + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ChangesetAction) Equals(other ChangesetAction) bool {
	switch data1 := _this.Get_().(type) {
	case ChangesetAction_StatusChanged:
		{
			data2, ok := other.Get_().(ChangesetAction_StatusChanged)
			return ok && data1.Status.Equals(data2.Status) && data1.Error.Equals(data2.Error)
		}
	case ChangesetAction_FileSet:
		{
			data2, ok := other.Get_().(ChangesetAction_FileSet)
			return ok && data1.File.Equals(data2.File)
		}
	case ChangesetAction_FileRemoved:
		{
			data2, ok := other.Get_().(ChangesetAction_FileRemoved)
			return ok && data1.FileId.Equals(data2.FileId)
		}
	case ChangesetAction_OperationsChanged:
		{
			data2, ok := other.Get_().(ChangesetAction_OperationsChanged)
			return ok && data1.Operations.Equals(data2.Operations)
		}
	case ChangesetAction_Cleared:
		{
			_, ok := other.Get_().(ChangesetAction_Cleared)
			return ok
		}
	case ChangesetAction_OperationStatusChanged:
		{
			data2, ok := other.Get_().(ChangesetAction_OperationStatusChanged)
			return ok && data1.OperationId.Equals(data2.OperationId) && data1.Status.Equals(data2.Status) && data1.Error.Equals(data2.Error)
		}
	case ChangesetAction_ContentChanged:
		{
			data2, ok := other.Get_().(ChangesetAction_ContentChanged)
			return ok && data1.Files.Equals(data2.Files) && data1.Operations.Equals(data2.Operations) && data1.Error.Equals(data2.Error)
		}
	case ChangesetAction_FilesReviewedChanged:
		{
			data2, ok := other.Get_().(ChangesetAction_FilesReviewedChanged)
			return ok && data1.FileIds.Equals(data2.FileIds) && data1.Reviewed == data2.Reviewed
		}
	case ChangesetAction_CsUnknown:
		{
			data2, ok := other.Get_().(ChangesetAction_CsUnknown)
			return ok && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ChangesetAction) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ChangesetAction)
	return ok && _this.Equals(typed)
}

func Type_ChangesetAction_() _dafny.TypeDescriptor {
	return type_ChangesetAction_{}
}

type type_ChangesetAction_ struct {
}

func (_this type_ChangesetAction_) Default() interface{} {
	return Companion_ChangesetAction_.Default()
}

func (_this type_ChangesetAction_) String() string {
	return "Changeset.ChangesetAction"
}
func (_this ChangesetAction) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ChangesetAction{}

// End of datatype ChangesetAction
