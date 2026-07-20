import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_ as module_
import _dafny as _dafny
import System_ as System_
import ConfluxCodec as ConfluxCodec
import ConfluxContract as ConfluxContract
import AhpSkeleton as AhpSkeleton
import ResourceWatch as ResourceWatch
import Canvas as Canvas
import ConfluxOperators as ConfluxOperators
import ConfluxPrune as ConfluxPrune
import ConfluxSeqRoute as ConfluxSeqRoute
import Changeset as Changeset

# Module: Annotations

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def AnnId(a):
        return (a).id_

    @staticmethod
    def EntryId(e):
        return (e).id_

    @staticmethod
    def upsertAnn(anns, a):
        return ConfluxSeqRoute.default__.SeqUpsertById(default__.AnnId, anns, a)

    @staticmethod
    def removeAnn(anns, id_):
        return ConfluxSeqRoute.default__.SeqRemoveById(default__.AnnId, anns, id_)

    @staticmethod
    def hasAnn(anns, id_):
        def lambda0_(exists_var_0_):
            d_0_i_: int = exists_var_0_
            return (((0) <= (d_0_i_)) and ((d_0_i_) < (len(anns)))) and ((((anns)[d_0_i_]).id_) == (id_))

        return _dafny.quantifier(_dafny.IntegerRange(0, len(anns)), False, lambda0_)

    @staticmethod
    def upsertEntry(entries, e):
        return ConfluxSeqRoute.default__.SeqUpsertById(default__.EntryId, entries, e)

    @staticmethod
    def removeEntry(entries, id_):
        return ConfluxSeqRoute.default__.SeqRemoveById(default__.EntryId, entries, id_)

    @staticmethod
    def updateAnn(anns, id_, f):
        return ConfluxSeqRoute.default__.SeqUpdateById(default__.AnnId, anns, id_, f)

    @staticmethod
    def setEntry(e):
        def lambda0_(d_0_e_):
            def lambda1_(d_1_an_):
                def iife0_(_pat_let52_0):
                    def iife1_(d_2_dt__update__tmp_h0_):
                        def iife2_(_pat_let53_0):
                            def iife3_(d_3_dt__update_hentries_h0_):
                                return Annotation_Annotation((d_2_dt__update__tmp_h0_).id_, (d_2_dt__update__tmp_h0_).turnId, (d_2_dt__update__tmp_h0_).resource, (d_2_dt__update__tmp_h0_).range_, (d_2_dt__update__tmp_h0_).resolved, d_3_dt__update_hentries_h0_, (d_2_dt__update__tmp_h0_).meta)
                            return iife3_(_pat_let53_0)
                        return iife2_(default__.upsertEntry((d_1_an_).entries, d_0_e_))
                    return iife1_(_pat_let52_0)
                return iife0_(d_1_an_)

            return lambda1_

        return lambda0_(e)

    @staticmethod
    def dropEntry(eid):
        def lambda0_(d_0_eid_):
            def lambda1_(d_1_an_):
                def iife0_(_pat_let54_0):
                    def iife1_(d_2_dt__update__tmp_h0_):
                        def iife2_(_pat_let55_0):
                            def iife3_(d_3_dt__update_hentries_h0_):
                                return Annotation_Annotation((d_2_dt__update__tmp_h0_).id_, (d_2_dt__update__tmp_h0_).turnId, (d_2_dt__update__tmp_h0_).resource, (d_2_dt__update__tmp_h0_).range_, (d_2_dt__update__tmp_h0_).resolved, d_3_dt__update_hentries_h0_, (d_2_dt__update__tmp_h0_).meta)
                            return iife3_(_pat_let55_0)
                        return iife2_(default__.removeEntry((d_1_an_).entries, d_0_eid_))
                    return iife1_(_pat_let54_0)
                return iife0_(d_1_an_)

            return lambda1_

        return lambda0_(eid)

    @staticmethod
    def reanchor(tid, res, rng, rslv):
        def lambda0_(d_0_tid_, d_1_res_, d_2_rng_, d_3_rslv_):
            def lambda1_(d_4_an_):
                return default__.applyUpdate(d_4_an_, d_0_tid_, d_1_res_, d_2_rng_, d_3_rslv_)

            return lambda1_

        return lambda0_(tid, res, rng, rslv)

    @staticmethod
    def applyUpdate(ann, tid, res, rng, rslv):
        d_0_dt__update__tmp_h0_ = ann
        d_1_dt__update_hresolved_h0_ = ((rslv).value if (rslv).is_Some else (ann).resolved)
        d_2_dt__update_hrange_h0_ = (AhpSkeleton.Option_Some((rng).value) if (rng).is_Some else (ann).range_)
        d_3_dt__update_hresource_h0_ = ((res).value if (res).is_Some else (ann).resource)
        d_4_dt__update_hturnId_h0_ = ((tid).value if (tid).is_Some else (ann).turnId)
        return Annotation_Annotation((d_0_dt__update__tmp_h0_).id_, d_4_dt__update_hturnId_h0_, d_3_dt__update_hresource_h0_, d_2_dt__update_hrange_h0_, d_1_dt__update_hresolved_h0_, (d_0_dt__update__tmp_h0_).entries, (d_0_dt__update__tmp_h0_).meta)

    @staticmethod
    def ApplyToAnnotations(s, a, now):
        source0_ = a
        if True:
            if source0_.is_Set:
                d_0_ann_ = source0_.annotation
                return (default__.upsertAnn((s), d_0_ann_), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_Removed:
                d_1_id_ = source0_.annotationId
                return (default__.removeAnn((s), d_1_id_), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_EntrySet:
                d_2_aid_ = source0_.annotationId
                d_3_e_ = source0_.entry
                if default__.hasAnn((s), d_2_aid_):
                    return (default__.updateAnn((s), d_2_aid_, default__.setEntry(d_3_e_)), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_EntryRemoved:
                d_4_aid_ = source0_.annotationId
                d_5_eid_ = source0_.entryId
                if default__.hasAnn((s), d_4_aid_):
                    return (default__.updateAnn((s), d_4_aid_, default__.dropEntry(d_5_eid_)), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_Updated:
                d_6_aid_ = source0_.annotationId
                d_7_tid_ = source0_.turnId
                d_8_res_ = source0_.resource
                d_9_rng_ = source0_.range_
                d_10_rslv_ = source0_.resolved
                if default__.hasAnn((s), d_6_aid_):
                    return (default__.updateAnn((s), d_6_aid_, default__.reanchor(d_7_tid_, d_8_res_, d_9_rng_, d_10_rslv_)), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            return (s, AhpSkeleton.ReduceOutcome_NoOp())

    @staticmethod
    def apply1(s, a):
        return (default__.ApplyToAnnotations(s, a, 9999))[0]

    @staticmethod
    def fold(s, acts):
        return ConfluxContract.default__.Fold(default__.apply1, s, acts)

    @staticmethod
    def R():
        return ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "start")): ConfluxCodec.Json_JNull()}))

    @staticmethod
    def RunCorpus():
        pass_: int = int(0)
        total: int = int(0)
        total = 10
        pass_ = 0
        d_0_e1_: Entry
        d_0_e1_ = Entry_Entry(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "original")), AhpSkeleton.Option_None())
        d_1_a1_: Annotation
        d_1_a1_ = Annotation_Annotation(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turn-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/a.ts")), AhpSkeleton.Option_Some(default__.R()), False, _dafny.SeqWithoutIsStrInference([d_0_e1_]), AhpSkeleton.Option_None())
        d_2_st1_: _dafny.Seq
        d_2_st1_ = _dafny.SeqWithoutIsStrInference([d_1_a1_])
        d_3_a2_: Annotation
        d_3_a2_ = Annotation_Annotation(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t-2")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turn-2")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/b.ts")), AhpSkeleton.Option_None(), False, _dafny.SeqWithoutIsStrInference([Entry_Entry(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c-2")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")), AhpSkeleton.Option_None())]), AhpSkeleton.Option_None())
        if (default__.apply1(d_2_st1_, AnnotationsAction_Set(d_3_a2_))) == (_dafny.SeqWithoutIsStrInference([d_1_a1_, d_3_a2_])):
            pass_ = (pass_) + (1)
        d_4_a1r_: Annotation
        d_4_a1r_ = Annotation_Annotation(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turn-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/a.ts")), AhpSkeleton.Option_None(), True, _dafny.SeqWithoutIsStrInference([d_0_e1_]), AhpSkeleton.Option_None())
        if (default__.apply1(d_2_st1_, AnnotationsAction_Set(d_4_a1r_))) == (_dafny.SeqWithoutIsStrInference([d_4_a1r_])):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_2_st1_, AnnotationsAction_Removed(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t-1"))))) == (_dafny.SeqWithoutIsStrInference([])):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_2_st1_, AnnotationsAction_Removed(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nope"))))) == (d_2_st1_):
            pass_ = (pass_) + (1)
        d_5_e2_: Entry
        d_5_e2_ = Entry_Entry(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c-2")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "reply")), AhpSkeleton.Option_None())
        pat_let_tv0_ = d_0_e1_
        pat_let_tv1_ = d_5_e2_
        def iife0_(_pat_let56_0):
            def iife1_(d_6_dt__update__tmp_h0_):
                def iife2_(_pat_let57_0):
                    def iife3_(d_7_dt__update_hentries_h0_):
                        return Annotation_Annotation((d_6_dt__update__tmp_h0_).id_, (d_6_dt__update__tmp_h0_).turnId, (d_6_dt__update__tmp_h0_).resource, (d_6_dt__update__tmp_h0_).range_, (d_6_dt__update__tmp_h0_).resolved, d_7_dt__update_hentries_h0_, (d_6_dt__update__tmp_h0_).meta)
                    return iife3_(_pat_let57_0)
                return iife2_(_dafny.SeqWithoutIsStrInference([pat_let_tv0_, pat_let_tv1_]))
            return iife1_(_pat_let56_0)
        if (default__.apply1(d_2_st1_, AnnotationsAction_EntrySet(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t-1")), d_5_e2_))) == (_dafny.SeqWithoutIsStrInference([iife0_(d_1_a1_)])):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_2_st1_, AnnotationsAction_EntrySet(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nope")), d_5_e2_))) == (d_2_st1_):
            pass_ = (pass_) + (1)
        d_8_a1two_: Annotation
        d_9_dt__update__tmp_h1_ = d_1_a1_
        d_10_dt__update_hentries_h1_ = _dafny.SeqWithoutIsStrInference([d_0_e1_, d_5_e2_])
        d_8_a1two_ = Annotation_Annotation((d_9_dt__update__tmp_h1_).id_, (d_9_dt__update__tmp_h1_).turnId, (d_9_dt__update__tmp_h1_).resource, (d_9_dt__update__tmp_h1_).range_, (d_9_dt__update__tmp_h1_).resolved, d_10_dt__update_hentries_h1_, (d_9_dt__update__tmp_h1_).meta)
        if (default__.apply1(_dafny.SeqWithoutIsStrInference([d_8_a1two_]), AnnotationsAction_EntryRemoved(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c-2"))))) == (_dafny.SeqWithoutIsStrInference([d_1_a1_])):
            pass_ = (pass_) + (1)
        def iife4_(_pat_let58_0):
            def iife5_(d_11_dt__update__tmp_h2_):
                def iife6_(_pat_let59_0):
                    def iife7_(d_12_dt__update_hresolved_h0_):
                        return Annotation_Annotation((d_11_dt__update__tmp_h2_).id_, (d_11_dt__update__tmp_h2_).turnId, (d_11_dt__update__tmp_h2_).resource, (d_11_dt__update__tmp_h2_).range_, d_12_dt__update_hresolved_h0_, (d_11_dt__update__tmp_h2_).entries, (d_11_dt__update__tmp_h2_).meta)
                    return iife7_(_pat_let59_0)
                return iife6_(True)
            return iife5_(_pat_let58_0)
        if (default__.apply1(d_2_st1_, AnnotationsAction_Updated(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t-1")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(True)))) == (_dafny.SeqWithoutIsStrInference([iife4_(d_1_a1_)])):
            pass_ = (pass_) + (1)
        d_13_newR_: ConfluxCodec.Json
        d_13_newR_ = ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "start")): ConfluxCodec.Json_JBool(True)}))
        d_14_a1re_: Annotation
        d_15_dt__update__tmp_h3_ = d_1_a1_
        d_16_dt__update_hrange_h0_ = AhpSkeleton.Option_Some(d_13_newR_)
        d_17_dt__update_hresource_h0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/b.ts"))
        d_18_dt__update_hturnId_h0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turn-2"))
        d_14_a1re_ = Annotation_Annotation((d_15_dt__update__tmp_h3_).id_, d_18_dt__update_hturnId_h0_, d_17_dt__update_hresource_h0_, d_16_dt__update_hrange_h0_, (d_15_dt__update__tmp_h3_).resolved, (d_15_dt__update__tmp_h3_).entries, (d_15_dt__update__tmp_h3_).meta)
        if (default__.apply1(d_2_st1_, AnnotationsAction_Updated(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t-1")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turn-2"))), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/b.ts"))), AhpSkeleton.Option_Some(d_13_newR_), AhpSkeleton.Option_None()))) == (_dafny.SeqWithoutIsStrInference([d_14_a1re_])):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_2_st1_, AnnotationsAction_Updated(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nope")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(True)))) == (d_2_st1_):
            pass_ = (pass_) + (1)
        return pass_, total


class Entry:
    @classmethod
    def default(cls, ):
        return lambda: Entry_Entry(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Entry(self) -> bool:
        return isinstance(self, Entry_Entry)

class Entry_Entry(Entry, NamedTuple('Entry', [('id_', Any), ('text', Any), ('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Annotations.Entry.Entry({self.id_.VerbatimString(True)}, {self.text.VerbatimString(True)}, {_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Entry_Entry) and self.id_ == __o.id_ and self.text == __o.text and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()


class Annotation:
    @classmethod
    def default(cls, ):
        return lambda: Annotation_Annotation(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()(), False, _dafny.Seq({}), AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Annotation(self) -> bool:
        return isinstance(self, Annotation_Annotation)

class Annotation_Annotation(Annotation, NamedTuple('Annotation', [('id_', Any), ('turnId', Any), ('resource', Any), ('range_', Any), ('resolved', Any), ('entries', Any), ('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Annotations.Annotation.Annotation({self.id_.VerbatimString(True)}, {self.turnId.VerbatimString(True)}, {self.resource.VerbatimString(True)}, {_dafny.string_of(self.range_)}, {_dafny.string_of(self.resolved)}, {_dafny.string_of(self.entries)}, {_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Annotation_Annotation) and self.id_ == __o.id_ and self.turnId == __o.turnId and self.resource == __o.resource and self.range_ == __o.range_ and self.resolved == __o.resolved and self.entries == __o.entries and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()


class AnnotationsState:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_AnnotationsState(self) -> bool:
        return isinstance(self, AnnotationsState_AnnotationsState)

class AnnotationsState_AnnotationsState(AnnotationsState, NamedTuple('AnnotationsState', [('annotations', Any)])):
    def __dafnystr__(self) -> str:
        return f'Annotations.AnnotationsState.AnnotationsState({_dafny.string_of(self.annotations)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AnnotationsState_AnnotationsState) and self.annotations == __o.annotations
    def __hash__(self) -> int:
        return super().__hash__()


class AnnotationsAction:
    @classmethod
    def default(cls, ):
        return lambda: AnnotationsAction_Set(Annotation.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Set(self) -> bool:
        return isinstance(self, AnnotationsAction_Set)
    @property
    def is_Removed(self) -> bool:
        return isinstance(self, AnnotationsAction_Removed)
    @property
    def is_EntrySet(self) -> bool:
        return isinstance(self, AnnotationsAction_EntrySet)
    @property
    def is_EntryRemoved(self) -> bool:
        return isinstance(self, AnnotationsAction_EntryRemoved)
    @property
    def is_Updated(self) -> bool:
        return isinstance(self, AnnotationsAction_Updated)
    @property
    def is_AnUnknown(self) -> bool:
        return isinstance(self, AnnotationsAction_AnUnknown)

class AnnotationsAction_Set(AnnotationsAction, NamedTuple('Set', [('annotation', Any)])):
    def __dafnystr__(self) -> str:
        return f'Annotations.AnnotationsAction.Set({_dafny.string_of(self.annotation)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AnnotationsAction_Set) and self.annotation == __o.annotation
    def __hash__(self) -> int:
        return super().__hash__()

class AnnotationsAction_Removed(AnnotationsAction, NamedTuple('Removed', [('annotationId', Any)])):
    def __dafnystr__(self) -> str:
        return f'Annotations.AnnotationsAction.Removed({self.annotationId.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AnnotationsAction_Removed) and self.annotationId == __o.annotationId
    def __hash__(self) -> int:
        return super().__hash__()

class AnnotationsAction_EntrySet(AnnotationsAction, NamedTuple('EntrySet', [('annotationId', Any), ('entry', Any)])):
    def __dafnystr__(self) -> str:
        return f'Annotations.AnnotationsAction.EntrySet({self.annotationId.VerbatimString(True)}, {_dafny.string_of(self.entry)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AnnotationsAction_EntrySet) and self.annotationId == __o.annotationId and self.entry == __o.entry
    def __hash__(self) -> int:
        return super().__hash__()

class AnnotationsAction_EntryRemoved(AnnotationsAction, NamedTuple('EntryRemoved', [('annotationId', Any), ('entryId', Any)])):
    def __dafnystr__(self) -> str:
        return f'Annotations.AnnotationsAction.EntryRemoved({self.annotationId.VerbatimString(True)}, {self.entryId.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AnnotationsAction_EntryRemoved) and self.annotationId == __o.annotationId and self.entryId == __o.entryId
    def __hash__(self) -> int:
        return super().__hash__()

class AnnotationsAction_Updated(AnnotationsAction, NamedTuple('Updated', [('annotationId', Any), ('turnId', Any), ('resource', Any), ('range_', Any), ('resolved', Any)])):
    def __dafnystr__(self) -> str:
        return f'Annotations.AnnotationsAction.Updated({self.annotationId.VerbatimString(True)}, {_dafny.string_of(self.turnId)}, {_dafny.string_of(self.resource)}, {_dafny.string_of(self.range_)}, {_dafny.string_of(self.resolved)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AnnotationsAction_Updated) and self.annotationId == __o.annotationId and self.turnId == __o.turnId and self.resource == __o.resource and self.range_ == __o.range_ and self.resolved == __o.resolved
    def __hash__(self) -> int:
        return super().__hash__()

class AnnotationsAction_AnUnknown(AnnotationsAction, NamedTuple('AnUnknown', [('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'Annotations.AnnotationsAction.AnUnknown({_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AnnotationsAction_AnUnknown) and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()

