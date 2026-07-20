// Package Annotations
// Dafny module Annotations compiled into Go

package Annotations

import (
	os "os"

	m_AhpSkeleton "github.com/joshmouch/ahp-go/AhpSkeleton"
	m_Canvas "github.com/joshmouch/ahp-go/Canvas"
	m_Changeset "github.com/joshmouch/ahp-go/Changeset"
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
var _ m_Changeset.Dummy__

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
	return "Annotations.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) AnnId(a Annotation) _dafny.Sequence {
	return (a).Dtor_id()
}
func (_static *CompanionStruct_Default___) EntryId(e Entry) _dafny.Sequence {
	return (e).Dtor_id()
}
func (_static *CompanionStruct_Default___) UpsertAnn(anns _dafny.Sequence, a Annotation) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpsertById(func(coer14 func(Annotation) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg19 interface{}) interface{} {
			return coer14(arg19.(Annotation))
		}
	}(Companion_Default___.AnnId), anns, a)
}
func (_static *CompanionStruct_Default___) RemoveAnn(anns _dafny.Sequence, id _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqRemoveById(func(coer15 func(Annotation) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg20 interface{}) interface{} {
			return coer15(arg20.(Annotation))
		}
	}(Companion_Default___.AnnId), anns, id)
}
func (_static *CompanionStruct_Default___) HasAnn(anns _dafny.Sequence, id _dafny.Sequence) bool {
	return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((anns).Cardinality())), false, func(_exists_var_0 _dafny.Int) bool {
		var _0_i _dafny.Int
		_0_i = interface{}(_exists_var_0).(_dafny.Int)
		return (((_0_i).Sign() != -1) && ((_0_i).Cmp(_dafny.IntOfUint32((anns).Cardinality())) < 0)) && (_dafny.Companion_Sequence_.Equal(((anns).Select((_0_i).Uint32()).(Annotation)).Dtor_id(), id))
	})
}
func (_static *CompanionStruct_Default___) UpsertEntry(entries _dafny.Sequence, e Entry) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpsertById(func(coer16 func(Entry) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg21 interface{}) interface{} {
			return coer16(arg21.(Entry))
		}
	}(Companion_Default___.EntryId), entries, e)
}
func (_static *CompanionStruct_Default___) RemoveEntry(entries _dafny.Sequence, id _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqRemoveById(func(coer17 func(Entry) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg22 interface{}) interface{} {
			return coer17(arg22.(Entry))
		}
	}(Companion_Default___.EntryId), entries, id)
}
func (_static *CompanionStruct_Default___) UpdateAnn(anns _dafny.Sequence, id _dafny.Sequence, f func(Annotation) Annotation) _dafny.Sequence {
	return m_ConfluxSeqRoute.Companion_Default___.SeqUpdateById(func(coer18 func(Annotation) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg23 interface{}) interface{} {
			return coer18(arg23.(Annotation))
		}
	}(Companion_Default___.AnnId), anns, id, func(coer19 func(Annotation) Annotation) func(interface{}) interface{} {
		return func(arg24 interface{}) interface{} {
			return coer19(arg24.(Annotation))
		}
	}(f))
}
func (_static *CompanionStruct_Default___) SetEntry(e Entry) func(Annotation) Annotation {
	return (func(_0_e Entry) func(Annotation) Annotation {
		return func(_1_an Annotation) Annotation {
			return func(_pat_let52_0 Annotation) Annotation {
				return func(_2_dt__update__tmp_h0 Annotation) Annotation {
					return func(_pat_let53_0 _dafny.Sequence) Annotation {
						return func(_3_dt__update_hentries_h0 _dafny.Sequence) Annotation {
							return Companion_Annotation_.Create_Annotation_((_2_dt__update__tmp_h0).Dtor_id(), (_2_dt__update__tmp_h0).Dtor_turnId(), (_2_dt__update__tmp_h0).Dtor_resource(), (_2_dt__update__tmp_h0).Dtor_range(), (_2_dt__update__tmp_h0).Dtor_resolved(), _3_dt__update_hentries_h0, (_2_dt__update__tmp_h0).Dtor_meta())
						}(_pat_let53_0)
					}(Companion_Default___.UpsertEntry((_1_an).Dtor_entries(), _0_e))
				}(_pat_let52_0)
			}(_1_an)
		}
	})(e)
}
func (_static *CompanionStruct_Default___) DropEntry(eid _dafny.Sequence) func(Annotation) Annotation {
	return (func(_0_eid _dafny.Sequence) func(Annotation) Annotation {
		return func(_1_an Annotation) Annotation {
			return func(_pat_let54_0 Annotation) Annotation {
				return func(_2_dt__update__tmp_h0 Annotation) Annotation {
					return func(_pat_let55_0 _dafny.Sequence) Annotation {
						return func(_3_dt__update_hentries_h0 _dafny.Sequence) Annotation {
							return Companion_Annotation_.Create_Annotation_((_2_dt__update__tmp_h0).Dtor_id(), (_2_dt__update__tmp_h0).Dtor_turnId(), (_2_dt__update__tmp_h0).Dtor_resource(), (_2_dt__update__tmp_h0).Dtor_range(), (_2_dt__update__tmp_h0).Dtor_resolved(), _3_dt__update_hentries_h0, (_2_dt__update__tmp_h0).Dtor_meta())
						}(_pat_let55_0)
					}(Companion_Default___.RemoveEntry((_1_an).Dtor_entries(), _0_eid))
				}(_pat_let54_0)
			}(_1_an)
		}
	})(eid)
}
func (_static *CompanionStruct_Default___) Reanchor(tid m_AhpSkeleton.Option, res m_AhpSkeleton.Option, rng m_AhpSkeleton.Option, rslv m_AhpSkeleton.Option) func(Annotation) Annotation {
	return (func(_0_tid m_AhpSkeleton.Option, _1_res m_AhpSkeleton.Option, _2_rng m_AhpSkeleton.Option, _3_rslv m_AhpSkeleton.Option) func(Annotation) Annotation {
		return func(_4_an Annotation) Annotation {
			return Companion_Default___.ApplyUpdate(_4_an, _0_tid, _1_res, _2_rng, _3_rslv)
		}
	})(tid, res, rng, rslv)
}
func (_static *CompanionStruct_Default___) ApplyUpdate(ann Annotation, tid m_AhpSkeleton.Option, res m_AhpSkeleton.Option, rng m_AhpSkeleton.Option, rslv m_AhpSkeleton.Option) Annotation {
	var _0_dt__update__tmp_h0 Annotation = ann
	_ = _0_dt__update__tmp_h0
	var _1_dt__update_hresolved_h0 bool = (func() bool {
		if (rslv).Is_Some() {
			return (rslv).Dtor_value().(bool)
		}
		return (ann).Dtor_resolved()
	})()
	_ = _1_dt__update_hresolved_h0
	var _2_dt__update_hrange_h0 m_AhpSkeleton.Option = (func() m_AhpSkeleton.Option {
		if (rng).Is_Some() {
			return m_AhpSkeleton.Companion_Option_.Create_Some_((rng).Dtor_value().(m_ConfluxCodec.Json))
		}
		return (ann).Dtor_range()
	})()
	_ = _2_dt__update_hrange_h0
	var _3_dt__update_hresource_h0 _dafny.Sequence = (func() _dafny.Sequence {
		if (res).Is_Some() {
			return (res).Dtor_value().(_dafny.Sequence)
		}
		return (ann).Dtor_resource()
	})()
	_ = _3_dt__update_hresource_h0
	var _4_dt__update_hturnId_h0 _dafny.Sequence = (func() _dafny.Sequence {
		if (tid).Is_Some() {
			return (tid).Dtor_value().(_dafny.Sequence)
		}
		return (ann).Dtor_turnId()
	})()
	_ = _4_dt__update_hturnId_h0
	return Companion_Annotation_.Create_Annotation_((_0_dt__update__tmp_h0).Dtor_id(), _4_dt__update_hturnId_h0, _3_dt__update_hresource_h0, _2_dt__update_hrange_h0, _1_dt__update_hresolved_h0, (_0_dt__update__tmp_h0).Dtor_entries(), (_0_dt__update__tmp_h0).Dtor_meta())
}
func (_static *CompanionStruct_Default___) ApplyToAnnotations(s _dafny.Sequence, a AnnotationsAction, now _dafny.Int) _dafny.Tuple {
	var _source0 AnnotationsAction = a
	_ = _source0
	{
		if _source0.Is_Set() {
			var _0_ann Annotation = _source0.Get_().(AnnotationsAction_Set).Annotation
			_ = _0_ann
			return _dafny.TupleOf(Companion_Default___.UpsertAnn((s), _0_ann), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_Removed() {
			var _1_id _dafny.Sequence = _source0.Get_().(AnnotationsAction_Removed).AnnotationId
			_ = _1_id
			return _dafny.TupleOf(Companion_Default___.RemoveAnn((s), _1_id), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
		}
	}
	{
		if _source0.Is_EntrySet() {
			var _2_aid _dafny.Sequence = _source0.Get_().(AnnotationsAction_EntrySet).AnnotationId
			_ = _2_aid
			var _3_e Entry = _source0.Get_().(AnnotationsAction_EntrySet).Entry
			_ = _3_e
			if Companion_Default___.HasAnn((s), _2_aid) {
				return _dafny.TupleOf(Companion_Default___.UpdateAnn((s), _2_aid, Companion_Default___.SetEntry(_3_e)), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_EntryRemoved() {
			var _4_aid _dafny.Sequence = _source0.Get_().(AnnotationsAction_EntryRemoved).AnnotationId
			_ = _4_aid
			var _5_eid _dafny.Sequence = _source0.Get_().(AnnotationsAction_EntryRemoved).EntryId
			_ = _5_eid
			if Companion_Default___.HasAnn((s), _4_aid) {
				return _dafny.TupleOf(Companion_Default___.UpdateAnn((s), _4_aid, Companion_Default___.DropEntry(_5_eid)), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		if _source0.Is_Updated() {
			var _6_aid _dafny.Sequence = _source0.Get_().(AnnotationsAction_Updated).AnnotationId
			_ = _6_aid
			var _7_tid m_AhpSkeleton.Option = _source0.Get_().(AnnotationsAction_Updated).TurnId
			_ = _7_tid
			var _8_res m_AhpSkeleton.Option = _source0.Get_().(AnnotationsAction_Updated).Resource
			_ = _8_res
			var _9_rng m_AhpSkeleton.Option = _source0.Get_().(AnnotationsAction_Updated).Range
			_ = _9_rng
			var _10_rslv m_AhpSkeleton.Option = _source0.Get_().(AnnotationsAction_Updated).Resolved
			_ = _10_rslv
			if Companion_Default___.HasAnn((s), _6_aid) {
				return _dafny.TupleOf(Companion_Default___.UpdateAnn((s), _6_aid, Companion_Default___.Reanchor(_7_tid, _8_res, _9_rng, _10_rslv)), m_AhpSkeleton.Companion_ReduceOutcome_.Create_Applied_())
			} else {
				return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
			}
		}
	}
	{
		return _dafny.TupleOf(s, m_AhpSkeleton.Companion_ReduceOutcome_.Create_NoOp_())
	}
}
func (_static *CompanionStruct_Default___) Apply1(s _dafny.Sequence, a AnnotationsAction) _dafny.Sequence {
	return (*(Companion_Default___.ApplyToAnnotations(s, a, _dafny.IntOfInt64(9999))).IndexInt(0)).(_dafny.Sequence)
}
func (_static *CompanionStruct_Default___) Fold(s _dafny.Sequence, acts _dafny.Sequence) _dafny.Sequence {
	return m_ConfluxContract.Companion_Default___.Fold(func(coer20 func(_dafny.Sequence, AnnotationsAction) _dafny.Sequence) func(interface{}, interface{}) interface{} {
		return func(arg25 interface{}, arg26 interface{}) interface{} {
			return coer20(arg25.(_dafny.Sequence), arg26.(AnnotationsAction))
		}
	}(Companion_Default___.Apply1), s, acts).(_dafny.Sequence)
}
func (_static *CompanionStruct_Default___) R() m_ConfluxCodec.Json {
	return m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("start"), m_ConfluxCodec.Companion_Json_.Create_JNull_()))
}
func (_static *CompanionStruct_Default___) RunCorpus() (_dafny.Int, _dafny.Int) {
	var pass _dafny.Int = _dafny.Zero
	_ = pass
	var total _dafny.Int = _dafny.Zero
	_ = total
	total = _dafny.IntOfInt64(10)
	pass = _dafny.Zero
	var _0_e1 Entry
	_ = _0_e1
	_0_e1 = Companion_Entry_.Create_Entry_(_dafny.UnicodeSeqOfUtf8Bytes("c-1"), _dafny.UnicodeSeqOfUtf8Bytes("original"), m_AhpSkeleton.Companion_Option_.Create_None_())
	var _1_a1 Annotation
	_ = _1_a1
	_1_a1 = Companion_Annotation_.Create_Annotation_(_dafny.UnicodeSeqOfUtf8Bytes("t-1"), _dafny.UnicodeSeqOfUtf8Bytes("turn-1"), _dafny.UnicodeSeqOfUtf8Bytes("file:///src/a.ts"), m_AhpSkeleton.Companion_Option_.Create_Some_(Companion_Default___.R()), false, _dafny.SeqOf(_0_e1), m_AhpSkeleton.Companion_Option_.Create_None_())
	var _2_st1 _dafny.Sequence
	_ = _2_st1
	_2_st1 = _dafny.SeqOf(_1_a1)
	var _3_a2 Annotation
	_ = _3_a2
	_3_a2 = Companion_Annotation_.Create_Annotation_(_dafny.UnicodeSeqOfUtf8Bytes("t-2"), _dafny.UnicodeSeqOfUtf8Bytes("turn-2"), _dafny.UnicodeSeqOfUtf8Bytes("file:///src/b.ts"), m_AhpSkeleton.Companion_Option_.Create_None_(), false, _dafny.SeqOf(Companion_Entry_.Create_Entry_(_dafny.UnicodeSeqOfUtf8Bytes("c-2"), _dafny.UnicodeSeqOfUtf8Bytes("x"), m_AhpSkeleton.Companion_Option_.Create_None_())), m_AhpSkeleton.Companion_Option_.Create_None_())
	if (Companion_Default___.Apply1(_2_st1, Companion_AnnotationsAction_.Create_Set_(_3_a2))).Equals(_dafny.SeqOf(_1_a1, _3_a2)) {
		pass = (pass).Plus(_dafny.One)
	}
	var _4_a1r Annotation
	_ = _4_a1r
	_4_a1r = Companion_Annotation_.Create_Annotation_(_dafny.UnicodeSeqOfUtf8Bytes("t-1"), _dafny.UnicodeSeqOfUtf8Bytes("turn-1"), _dafny.UnicodeSeqOfUtf8Bytes("file:///src/a.ts"), m_AhpSkeleton.Companion_Option_.Create_None_(), true, _dafny.SeqOf(_0_e1), m_AhpSkeleton.Companion_Option_.Create_None_())
	if (Companion_Default___.Apply1(_2_st1, Companion_AnnotationsAction_.Create_Set_(_4_a1r))).Equals(_dafny.SeqOf(_4_a1r)) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_2_st1, Companion_AnnotationsAction_.Create_Removed_(_dafny.UnicodeSeqOfUtf8Bytes("t-1")))).Equals(_dafny.SeqOf()) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_2_st1, Companion_AnnotationsAction_.Create_Removed_(_dafny.UnicodeSeqOfUtf8Bytes("nope")))).Equals(_2_st1) {
		pass = (pass).Plus(_dafny.One)
	}
	var _5_e2 Entry
	_ = _5_e2
	_5_e2 = Companion_Entry_.Create_Entry_(_dafny.UnicodeSeqOfUtf8Bytes("c-2"), _dafny.UnicodeSeqOfUtf8Bytes("reply"), m_AhpSkeleton.Companion_Option_.Create_None_())
	var _pat_let_tv0 = _0_e1
	_ = _pat_let_tv0
	var _pat_let_tv1 = _5_e2
	_ = _pat_let_tv1
	if (Companion_Default___.Apply1(_2_st1, Companion_AnnotationsAction_.Create_EntrySet_(_dafny.UnicodeSeqOfUtf8Bytes("t-1"), _5_e2))).Equals(_dafny.SeqOf(func(_pat_let56_0 Annotation) Annotation {
		return func(_6_dt__update__tmp_h0 Annotation) Annotation {
			return func(_pat_let57_0 _dafny.Sequence) Annotation {
				return func(_7_dt__update_hentries_h0 _dafny.Sequence) Annotation {
					return Companion_Annotation_.Create_Annotation_((_6_dt__update__tmp_h0).Dtor_id(), (_6_dt__update__tmp_h0).Dtor_turnId(), (_6_dt__update__tmp_h0).Dtor_resource(), (_6_dt__update__tmp_h0).Dtor_range(), (_6_dt__update__tmp_h0).Dtor_resolved(), _7_dt__update_hentries_h0, (_6_dt__update__tmp_h0).Dtor_meta())
				}(_pat_let57_0)
			}(_dafny.SeqOf(_pat_let_tv0, _pat_let_tv1))
		}(_pat_let56_0)
	}(_1_a1))) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_2_st1, Companion_AnnotationsAction_.Create_EntrySet_(_dafny.UnicodeSeqOfUtf8Bytes("nope"), _5_e2))).Equals(_2_st1) {
		pass = (pass).Plus(_dafny.One)
	}
	var _8_a1two Annotation
	_ = _8_a1two
	var _9_dt__update__tmp_h1 Annotation = _1_a1
	_ = _9_dt__update__tmp_h1
	var _10_dt__update_hentries_h1 _dafny.Sequence = _dafny.SeqOf(_0_e1, _5_e2)
	_ = _10_dt__update_hentries_h1
	_8_a1two = Companion_Annotation_.Create_Annotation_((_9_dt__update__tmp_h1).Dtor_id(), (_9_dt__update__tmp_h1).Dtor_turnId(), (_9_dt__update__tmp_h1).Dtor_resource(), (_9_dt__update__tmp_h1).Dtor_range(), (_9_dt__update__tmp_h1).Dtor_resolved(), _10_dt__update_hentries_h1, (_9_dt__update__tmp_h1).Dtor_meta())
	if (Companion_Default___.Apply1(_dafny.SeqOf(_8_a1two), Companion_AnnotationsAction_.Create_EntryRemoved_(_dafny.UnicodeSeqOfUtf8Bytes("t-1"), _dafny.UnicodeSeqOfUtf8Bytes("c-2")))).Equals(_dafny.SeqOf(_1_a1)) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_2_st1, Companion_AnnotationsAction_.Create_Updated_(_dafny.UnicodeSeqOfUtf8Bytes("t-1"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(true)))).Equals(_dafny.SeqOf(func(_pat_let58_0 Annotation) Annotation {
		return func(_11_dt__update__tmp_h2 Annotation) Annotation {
			return func(_pat_let59_0 bool) Annotation {
				return func(_12_dt__update_hresolved_h0 bool) Annotation {
					return Companion_Annotation_.Create_Annotation_((_11_dt__update__tmp_h2).Dtor_id(), (_11_dt__update__tmp_h2).Dtor_turnId(), (_11_dt__update__tmp_h2).Dtor_resource(), (_11_dt__update__tmp_h2).Dtor_range(), _12_dt__update_hresolved_h0, (_11_dt__update__tmp_h2).Dtor_entries(), (_11_dt__update__tmp_h2).Dtor_meta())
				}(_pat_let59_0)
			}(true)
		}(_pat_let58_0)
	}(_1_a1))) {
		pass = (pass).Plus(_dafny.One)
	}
	var _13_newR m_ConfluxCodec.Json
	_ = _13_newR
	_13_newR = m_ConfluxCodec.Companion_Json_.Create_JObj_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("start"), m_ConfluxCodec.Companion_Json_.Create_JBool_(true)))
	var _14_a1re Annotation
	_ = _14_a1re
	var _15_dt__update__tmp_h3 Annotation = _1_a1
	_ = _15_dt__update__tmp_h3
	var _16_dt__update_hrange_h0 m_AhpSkeleton.Option = m_AhpSkeleton.Companion_Option_.Create_Some_(_13_newR)
	_ = _16_dt__update_hrange_h0
	var _17_dt__update_hresource_h0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("file:///src/b.ts")
	_ = _17_dt__update_hresource_h0
	var _18_dt__update_hturnId_h0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("turn-2")
	_ = _18_dt__update_hturnId_h0
	_14_a1re = Companion_Annotation_.Create_Annotation_((_15_dt__update__tmp_h3).Dtor_id(), _18_dt__update_hturnId_h0, _17_dt__update_hresource_h0, _16_dt__update_hrange_h0, (_15_dt__update__tmp_h3).Dtor_resolved(), (_15_dt__update__tmp_h3).Dtor_entries(), (_15_dt__update__tmp_h3).Dtor_meta())
	if (Companion_Default___.Apply1(_2_st1, Companion_AnnotationsAction_.Create_Updated_(_dafny.UnicodeSeqOfUtf8Bytes("t-1"), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("turn-2")), m_AhpSkeleton.Companion_Option_.Create_Some_(_dafny.UnicodeSeqOfUtf8Bytes("file:///src/b.ts")), m_AhpSkeleton.Companion_Option_.Create_Some_(_13_newR), m_AhpSkeleton.Companion_Option_.Create_None_()))).Equals(_dafny.SeqOf(_14_a1re)) {
		pass = (pass).Plus(_dafny.One)
	}
	if (Companion_Default___.Apply1(_2_st1, Companion_AnnotationsAction_.Create_Updated_(_dafny.UnicodeSeqOfUtf8Bytes("nope"), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_None_(), m_AhpSkeleton.Companion_Option_.Create_Some_(true)))).Equals(_2_st1) {
		pass = (pass).Plus(_dafny.One)
	}
	return pass, total
}

// End of class Default__

// Definition of datatype Entry
type Entry struct {
	Data_Entry_
}

func (_this Entry) Get_() Data_Entry_ {
	return _this.Data_Entry_
}

type Data_Entry_ interface {
	isEntry()
}

type CompanionStruct_Entry_ struct {
}

var Companion_Entry_ = CompanionStruct_Entry_{}

type Entry_Entry struct {
	Id   _dafny.Sequence
	Text _dafny.Sequence
	Meta m_AhpSkeleton.Option
}

func (Entry_Entry) isEntry() {}

func (CompanionStruct_Entry_) Create_Entry_(Id _dafny.Sequence, Text _dafny.Sequence, Meta m_AhpSkeleton.Option) Entry {
	return Entry{Entry_Entry{Id, Text, Meta}}
}

func (_this Entry) Is_Entry() bool {
	_, ok := _this.Get_().(Entry_Entry)
	return ok
}

func (CompanionStruct_Entry_) Default() Entry {
	return Companion_Entry_.Create_Entry_(_dafny.EmptySeq, _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default())
}

func (_this Entry) Dtor_id() _dafny.Sequence {
	return _this.Get_().(Entry_Entry).Id
}

func (_this Entry) Dtor_text() _dafny.Sequence {
	return _this.Get_().(Entry_Entry).Text
}

func (_this Entry) Dtor_meta() m_AhpSkeleton.Option {
	return _this.Get_().(Entry_Entry).Meta
}

func (_this Entry) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Entry_Entry:
		{
			return "Annotations.Entry.Entry" + "(" + data.Id.VerbatimString(true) + ", " + data.Text.VerbatimString(true) + ", " + _dafny.String(data.Meta) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Entry) Equals(other Entry) bool {
	switch data1 := _this.Get_().(type) {
	case Entry_Entry:
		{
			data2, ok := other.Get_().(Entry_Entry)
			return ok && data1.Id.Equals(data2.Id) && data1.Text.Equals(data2.Text) && data1.Meta.Equals(data2.Meta)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Entry) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Entry)
	return ok && _this.Equals(typed)
}

func Type_Entry_() _dafny.TypeDescriptor {
	return type_Entry_{}
}

type type_Entry_ struct {
}

func (_this type_Entry_) Default() interface{} {
	return Companion_Entry_.Default()
}

func (_this type_Entry_) String() string {
	return "Annotations.Entry"
}
func (_this Entry) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Entry{}

// End of datatype Entry

// Definition of datatype Annotation
type Annotation struct {
	Data_Annotation_
}

func (_this Annotation) Get_() Data_Annotation_ {
	return _this.Data_Annotation_
}

type Data_Annotation_ interface {
	isAnnotation()
}

type CompanionStruct_Annotation_ struct {
}

var Companion_Annotation_ = CompanionStruct_Annotation_{}

type Annotation_Annotation struct {
	Id       _dafny.Sequence
	TurnId   _dafny.Sequence
	Resource _dafny.Sequence
	Range    m_AhpSkeleton.Option
	Resolved bool
	Entries  _dafny.Sequence
	Meta     m_AhpSkeleton.Option
}

func (Annotation_Annotation) isAnnotation() {}

func (CompanionStruct_Annotation_) Create_Annotation_(Id _dafny.Sequence, TurnId _dafny.Sequence, Resource _dafny.Sequence, Range m_AhpSkeleton.Option, Resolved bool, Entries _dafny.Sequence, Meta m_AhpSkeleton.Option) Annotation {
	return Annotation{Annotation_Annotation{Id, TurnId, Resource, Range, Resolved, Entries, Meta}}
}

func (_this Annotation) Is_Annotation() bool {
	_, ok := _this.Get_().(Annotation_Annotation)
	return ok
}

func (CompanionStruct_Annotation_) Default() Annotation {
	return Companion_Annotation_.Create_Annotation_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default(), false, _dafny.EmptySeq, m_AhpSkeleton.Companion_Option_.Default())
}

func (_this Annotation) Dtor_id() _dafny.Sequence {
	return _this.Get_().(Annotation_Annotation).Id
}

func (_this Annotation) Dtor_turnId() _dafny.Sequence {
	return _this.Get_().(Annotation_Annotation).TurnId
}

func (_this Annotation) Dtor_resource() _dafny.Sequence {
	return _this.Get_().(Annotation_Annotation).Resource
}

func (_this Annotation) Dtor_range() m_AhpSkeleton.Option {
	return _this.Get_().(Annotation_Annotation).Range
}

func (_this Annotation) Dtor_resolved() bool {
	return _this.Get_().(Annotation_Annotation).Resolved
}

func (_this Annotation) Dtor_entries() _dafny.Sequence {
	return _this.Get_().(Annotation_Annotation).Entries
}

func (_this Annotation) Dtor_meta() m_AhpSkeleton.Option {
	return _this.Get_().(Annotation_Annotation).Meta
}

func (_this Annotation) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Annotation_Annotation:
		{
			return "Annotations.Annotation.Annotation" + "(" + data.Id.VerbatimString(true) + ", " + data.TurnId.VerbatimString(true) + ", " + data.Resource.VerbatimString(true) + ", " + _dafny.String(data.Range) + ", " + _dafny.String(data.Resolved) + ", " + _dafny.String(data.Entries) + ", " + _dafny.String(data.Meta) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Annotation) Equals(other Annotation) bool {
	switch data1 := _this.Get_().(type) {
	case Annotation_Annotation:
		{
			data2, ok := other.Get_().(Annotation_Annotation)
			return ok && data1.Id.Equals(data2.Id) && data1.TurnId.Equals(data2.TurnId) && data1.Resource.Equals(data2.Resource) && data1.Range.Equals(data2.Range) && data1.Resolved == data2.Resolved && data1.Entries.Equals(data2.Entries) && data1.Meta.Equals(data2.Meta)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Annotation) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Annotation)
	return ok && _this.Equals(typed)
}

func Type_Annotation_() _dafny.TypeDescriptor {
	return type_Annotation_{}
}

type type_Annotation_ struct {
}

func (_this type_Annotation_) Default() interface{} {
	return Companion_Annotation_.Default()
}

func (_this type_Annotation_) String() string {
	return "Annotations.Annotation"
}
func (_this Annotation) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Annotation{}

// End of datatype Annotation

// Definition of datatype AnnotationsState
type AnnotationsState struct {
	Data_AnnotationsState_
}

func (_this AnnotationsState) Get_() Data_AnnotationsState_ {
	return _this.Data_AnnotationsState_
}

type Data_AnnotationsState_ interface {
	isAnnotationsState()
}

type CompanionStruct_AnnotationsState_ struct {
}

var Companion_AnnotationsState_ = CompanionStruct_AnnotationsState_{}

type AnnotationsState_AnnotationsState struct {
	Annotations _dafny.Sequence
}

func (AnnotationsState_AnnotationsState) isAnnotationsState() {}

func (CompanionStruct_AnnotationsState_) Create_AnnotationsState_(Annotations _dafny.Sequence) AnnotationsState {
	return AnnotationsState{AnnotationsState_AnnotationsState{Annotations}}
}

func (_this AnnotationsState) Is_AnnotationsState() bool {
	_, ok := _this.Get_().(AnnotationsState_AnnotationsState)
	return ok
}

func (CompanionStruct_AnnotationsState_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this AnnotationsState) Dtor_annotations() _dafny.Sequence {
	return _this.Get_().(AnnotationsState_AnnotationsState).Annotations
}

func (_this AnnotationsState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case AnnotationsState_AnnotationsState:
		{
			return "Annotations.AnnotationsState.AnnotationsState" + "(" + _dafny.String(data.Annotations) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this AnnotationsState) Equals(other AnnotationsState) bool {
	switch data1 := _this.Get_().(type) {
	case AnnotationsState_AnnotationsState:
		{
			data2, ok := other.Get_().(AnnotationsState_AnnotationsState)
			return ok && data1.Annotations.Equals(data2.Annotations)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this AnnotationsState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(AnnotationsState)
	return ok && _this.Equals(typed)
}

func Type_AnnotationsState_() _dafny.TypeDescriptor {
	return type_AnnotationsState_{}
}

type type_AnnotationsState_ struct {
}

func (_this type_AnnotationsState_) Default() interface{} {
	return Companion_AnnotationsState_.Default()
}

func (_this type_AnnotationsState_) String() string {
	return "Annotations.AnnotationsState"
}
func (_this AnnotationsState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = AnnotationsState{}

// End of datatype AnnotationsState

// Definition of datatype AnnotationsAction
type AnnotationsAction struct {
	Data_AnnotationsAction_
}

func (_this AnnotationsAction) Get_() Data_AnnotationsAction_ {
	return _this.Data_AnnotationsAction_
}

type Data_AnnotationsAction_ interface {
	isAnnotationsAction()
}

type CompanionStruct_AnnotationsAction_ struct {
}

var Companion_AnnotationsAction_ = CompanionStruct_AnnotationsAction_{}

type AnnotationsAction_Set struct {
	Annotation Annotation
}

func (AnnotationsAction_Set) isAnnotationsAction() {}

func (CompanionStruct_AnnotationsAction_) Create_Set_(Annotation Annotation) AnnotationsAction {
	return AnnotationsAction{AnnotationsAction_Set{Annotation}}
}

func (_this AnnotationsAction) Is_Set() bool {
	_, ok := _this.Get_().(AnnotationsAction_Set)
	return ok
}

type AnnotationsAction_Removed struct {
	AnnotationId _dafny.Sequence
}

func (AnnotationsAction_Removed) isAnnotationsAction() {}

func (CompanionStruct_AnnotationsAction_) Create_Removed_(AnnotationId _dafny.Sequence) AnnotationsAction {
	return AnnotationsAction{AnnotationsAction_Removed{AnnotationId}}
}

func (_this AnnotationsAction) Is_Removed() bool {
	_, ok := _this.Get_().(AnnotationsAction_Removed)
	return ok
}

type AnnotationsAction_EntrySet struct {
	AnnotationId _dafny.Sequence
	Entry        Entry
}

func (AnnotationsAction_EntrySet) isAnnotationsAction() {}

func (CompanionStruct_AnnotationsAction_) Create_EntrySet_(AnnotationId _dafny.Sequence, Entry Entry) AnnotationsAction {
	return AnnotationsAction{AnnotationsAction_EntrySet{AnnotationId, Entry}}
}

func (_this AnnotationsAction) Is_EntrySet() bool {
	_, ok := _this.Get_().(AnnotationsAction_EntrySet)
	return ok
}

type AnnotationsAction_EntryRemoved struct {
	AnnotationId _dafny.Sequence
	EntryId      _dafny.Sequence
}

func (AnnotationsAction_EntryRemoved) isAnnotationsAction() {}

func (CompanionStruct_AnnotationsAction_) Create_EntryRemoved_(AnnotationId _dafny.Sequence, EntryId _dafny.Sequence) AnnotationsAction {
	return AnnotationsAction{AnnotationsAction_EntryRemoved{AnnotationId, EntryId}}
}

func (_this AnnotationsAction) Is_EntryRemoved() bool {
	_, ok := _this.Get_().(AnnotationsAction_EntryRemoved)
	return ok
}

type AnnotationsAction_Updated struct {
	AnnotationId _dafny.Sequence
	TurnId       m_AhpSkeleton.Option
	Resource     m_AhpSkeleton.Option
	Range        m_AhpSkeleton.Option
	Resolved     m_AhpSkeleton.Option
}

func (AnnotationsAction_Updated) isAnnotationsAction() {}

func (CompanionStruct_AnnotationsAction_) Create_Updated_(AnnotationId _dafny.Sequence, TurnId m_AhpSkeleton.Option, Resource m_AhpSkeleton.Option, Range m_AhpSkeleton.Option, Resolved m_AhpSkeleton.Option) AnnotationsAction {
	return AnnotationsAction{AnnotationsAction_Updated{AnnotationId, TurnId, Resource, Range, Resolved}}
}

func (_this AnnotationsAction) Is_Updated() bool {
	_, ok := _this.Get_().(AnnotationsAction_Updated)
	return ok
}

type AnnotationsAction_AnUnknown struct {
	Raw m_ConfluxCodec.Json
}

func (AnnotationsAction_AnUnknown) isAnnotationsAction() {}

func (CompanionStruct_AnnotationsAction_) Create_AnUnknown_(Raw m_ConfluxCodec.Json) AnnotationsAction {
	return AnnotationsAction{AnnotationsAction_AnUnknown{Raw}}
}

func (_this AnnotationsAction) Is_AnUnknown() bool {
	_, ok := _this.Get_().(AnnotationsAction_AnUnknown)
	return ok
}

func (CompanionStruct_AnnotationsAction_) Default() AnnotationsAction {
	return Companion_AnnotationsAction_.Create_Set_(Companion_Annotation_.Default())
}

func (_this AnnotationsAction) Dtor_annotation() Annotation {
	return _this.Get_().(AnnotationsAction_Set).Annotation
}

func (_this AnnotationsAction) Dtor_annotationId() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case AnnotationsAction_Removed:
		return data.AnnotationId
	case AnnotationsAction_EntrySet:
		return data.AnnotationId
	case AnnotationsAction_EntryRemoved:
		return data.AnnotationId
	default:
		return data.(AnnotationsAction_Updated).AnnotationId
	}
}

func (_this AnnotationsAction) Dtor_entry() Entry {
	return _this.Get_().(AnnotationsAction_EntrySet).Entry
}

func (_this AnnotationsAction) Dtor_entryId() _dafny.Sequence {
	return _this.Get_().(AnnotationsAction_EntryRemoved).EntryId
}

func (_this AnnotationsAction) Dtor_turnId() m_AhpSkeleton.Option {
	return _this.Get_().(AnnotationsAction_Updated).TurnId
}

func (_this AnnotationsAction) Dtor_resource() m_AhpSkeleton.Option {
	return _this.Get_().(AnnotationsAction_Updated).Resource
}

func (_this AnnotationsAction) Dtor_range() m_AhpSkeleton.Option {
	return _this.Get_().(AnnotationsAction_Updated).Range
}

func (_this AnnotationsAction) Dtor_resolved() m_AhpSkeleton.Option {
	return _this.Get_().(AnnotationsAction_Updated).Resolved
}

func (_this AnnotationsAction) Dtor_raw() m_ConfluxCodec.Json {
	return _this.Get_().(AnnotationsAction_AnUnknown).Raw
}

func (_this AnnotationsAction) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case AnnotationsAction_Set:
		{
			return "Annotations.AnnotationsAction.Set" + "(" + _dafny.String(data.Annotation) + ")"
		}
	case AnnotationsAction_Removed:
		{
			return "Annotations.AnnotationsAction.Removed" + "(" + data.AnnotationId.VerbatimString(true) + ")"
		}
	case AnnotationsAction_EntrySet:
		{
			return "Annotations.AnnotationsAction.EntrySet" + "(" + data.AnnotationId.VerbatimString(true) + ", " + _dafny.String(data.Entry) + ")"
		}
	case AnnotationsAction_EntryRemoved:
		{
			return "Annotations.AnnotationsAction.EntryRemoved" + "(" + data.AnnotationId.VerbatimString(true) + ", " + data.EntryId.VerbatimString(true) + ")"
		}
	case AnnotationsAction_Updated:
		{
			return "Annotations.AnnotationsAction.Updated" + "(" + data.AnnotationId.VerbatimString(true) + ", " + _dafny.String(data.TurnId) + ", " + _dafny.String(data.Resource) + ", " + _dafny.String(data.Range) + ", " + _dafny.String(data.Resolved) + ")"
		}
	case AnnotationsAction_AnUnknown:
		{
			return "Annotations.AnnotationsAction.AnUnknown" + "(" + _dafny.String(data.Raw) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this AnnotationsAction) Equals(other AnnotationsAction) bool {
	switch data1 := _this.Get_().(type) {
	case AnnotationsAction_Set:
		{
			data2, ok := other.Get_().(AnnotationsAction_Set)
			return ok && data1.Annotation.Equals(data2.Annotation)
		}
	case AnnotationsAction_Removed:
		{
			data2, ok := other.Get_().(AnnotationsAction_Removed)
			return ok && data1.AnnotationId.Equals(data2.AnnotationId)
		}
	case AnnotationsAction_EntrySet:
		{
			data2, ok := other.Get_().(AnnotationsAction_EntrySet)
			return ok && data1.AnnotationId.Equals(data2.AnnotationId) && data1.Entry.Equals(data2.Entry)
		}
	case AnnotationsAction_EntryRemoved:
		{
			data2, ok := other.Get_().(AnnotationsAction_EntryRemoved)
			return ok && data1.AnnotationId.Equals(data2.AnnotationId) && data1.EntryId.Equals(data2.EntryId)
		}
	case AnnotationsAction_Updated:
		{
			data2, ok := other.Get_().(AnnotationsAction_Updated)
			return ok && data1.AnnotationId.Equals(data2.AnnotationId) && data1.TurnId.Equals(data2.TurnId) && data1.Resource.Equals(data2.Resource) && data1.Range.Equals(data2.Range) && data1.Resolved.Equals(data2.Resolved)
		}
	case AnnotationsAction_AnUnknown:
		{
			data2, ok := other.Get_().(AnnotationsAction_AnUnknown)
			return ok && data1.Raw.Equals(data2.Raw)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this AnnotationsAction) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(AnnotationsAction)
	return ok && _this.Equals(typed)
}

func Type_AnnotationsAction_() _dafny.TypeDescriptor {
	return type_AnnotationsAction_{}
}

type type_AnnotationsAction_ struct {
}

func (_this type_AnnotationsAction_) Default() interface{} {
	return Companion_AnnotationsAction_.Default()
}

func (_this type_AnnotationsAction_) String() string {
	return "Annotations.AnnotationsAction"
}
func (_this AnnotationsAction) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = AnnotationsAction{}

// End of datatype AnnotationsAction
