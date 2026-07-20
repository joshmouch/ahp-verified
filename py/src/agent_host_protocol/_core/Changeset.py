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

# Module: Changeset

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def opSet(st, err):
        def lambda0_(d_0_err_, d_1_st_):
            def lambda1_(d_2_e_):
                def iife0_(_pat_let26_0):
                    def iife1_(d_3_dt__update__tmp_h0_):
                        def iife2_(_pat_let27_0):
                            def iife3_(d_4_dt__update_herror_h0_):
                                def iife4_(_pat_let28_0):
                                    def iife5_(d_5_dt__update_hstatus_h0_):
                                        return ChangesetOperation_ChangesetOperation((d_3_dt__update__tmp_h0_).id_, (d_3_dt__update__tmp_h0_).label__, (d_3_dt__update__tmp_h0_).scopes, d_5_dt__update_hstatus_h0_, d_4_dt__update_herror_h0_)
                                    return iife5_(_pat_let28_0)
                                return iife4_(d_1_st_)
                            return iife3_(_pat_let27_0)
                        return iife2_(d_0_err_)
                    return iife1_(_pat_let26_0)
                return iife0_(d_2_e_)

            return lambda1_

        return lambda0_(err, st)

    @staticmethod
    def upsertFile(files, f):
        return ConfluxSeqRoute.default__.SeqUpsertById(default__.fileKey, files, f)

    @staticmethod
    def removeFile(files, id_):
        return ConfluxSeqRoute.default__.SeqRemoveById(default__.fileKey, files, id_)

    @staticmethod
    def setReviewed(files, ids, rev):
        def lambda0_(d_0_ids_, d_1_rev_):
            def lambda1_(d_2_f_):
                def iife0_(_pat_let29_0):
                    def iife1_(d_3_dt__update__tmp_h0_):
                        def iife2_(_pat_let30_0):
                            def iife3_(d_4_dt__update_hreviewed_h0_):
                                return ChangesetFile_ChangesetFile((d_3_dt__update__tmp_h0_).id_, d_4_dt__update_hreviewed_h0_, (d_3_dt__update__tmp_h0_).edit)
                            return iife3_(_pat_let30_0)
                        return iife2_(AhpSkeleton.Option_Some(d_1_rev_))
                    return iife1_(_pat_let29_0)
                return (iife0_(d_2_f_) if ((d_2_f_).id_) in (d_0_ids_) else d_2_f_)

            return lambda1_

        return ConfluxOperators.default__.Map(lambda0_(ids, rev), files)

    @staticmethod
    def updateOp(ops, id_, status, error):
        return ConfluxSeqRoute.default__.SeqUpdateById(default__.opKey, ops, id_, default__.opSet(status, error))

    @staticmethod
    def hasOp(ops, id_):
        def lambda0_(exists_var_0_):
            d_0_i_: int = exists_var_0_
            return (((0) <= (d_0_i_)) and ((d_0_i_) < (len(ops)))) and ((((ops)[d_0_i_]).id_) == (id_))

        return _dafny.quantifier(_dafny.IntegerRange(0, len(ops)), False, lambda0_)

    @staticmethod
    def anyNeedsReviewed(files, ids, rev):
        def lambda0_(exists_var_0_):
            d_0_i_: int = exists_var_0_
            return ((((0) <= (d_0_i_)) and ((d_0_i_) < (len(files)))) and ((((files)[d_0_i_]).id_) in (ids))) and ((((files)[d_0_i_]).reviewed) != (AhpSkeleton.Option_Some(rev)))

        return _dafny.quantifier(_dafny.IntegerRange(0, len(files)), False, lambda0_)

    @staticmethod
    def ApplyToChangeset(s, a, now):
        pat_let_tv0_ = s
        pat_let_tv1_ = s
        pat_let_tv2_ = s
        source0_ = a
        if True:
            if source0_.is_StatusChanged:
                d_0_st_ = source0_.status
                d_1_err_ = source0_.error
                def iife0_(_pat_let31_0):
                    def iife1_(d_2_dt__update__tmp_h0_):
                        def iife2_(_pat_let32_0):
                            def iife3_(d_3_dt__update_herror_h0_):
                                def iife4_(_pat_let33_0):
                                    def iife5_(d_4_dt__update_hstatus_h0_):
                                        return ChangesetState_ChangesetState(d_4_dt__update_hstatus_h0_, (d_2_dt__update__tmp_h0_).files, (d_2_dt__update__tmp_h0_).operations, d_3_dt__update_herror_h0_)
                                    return iife5_(_pat_let33_0)
                                return iife4_(d_0_st_)
                            return iife3_(_pat_let32_0)
                        return iife2_(d_1_err_)
                    return iife1_(_pat_let31_0)
                return (iife0_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_FileSet:
                d_5_f_ = source0_.file
                def iife6_(_pat_let34_0):
                    def iife7_(d_6_dt__update__tmp_h1_):
                        def iife8_(_pat_let35_0):
                            def iife9_(d_7_dt__update_hfiles_h0_):
                                return ChangesetState_ChangesetState((d_6_dt__update__tmp_h1_).status, d_7_dt__update_hfiles_h0_, (d_6_dt__update__tmp_h1_).operations, (d_6_dt__update__tmp_h1_).error)
                            return iife9_(_pat_let35_0)
                        return iife8_(default__.upsertFile((pat_let_tv0_).files, d_5_f_))
                    return iife7_(_pat_let34_0)
                return (iife6_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_FileRemoved:
                d_8_id_ = source0_.fileId
                def iife10_(_pat_let36_0):
                    def iife11_(d_9_dt__update__tmp_h2_):
                        def iife12_(_pat_let37_0):
                            def iife13_(d_10_dt__update_hfiles_h1_):
                                return ChangesetState_ChangesetState((d_9_dt__update__tmp_h2_).status, d_10_dt__update_hfiles_h1_, (d_9_dt__update__tmp_h2_).operations, (d_9_dt__update__tmp_h2_).error)
                            return iife13_(_pat_let37_0)
                        return iife12_(default__.removeFile((pat_let_tv1_).files, d_8_id_))
                    return iife11_(_pat_let36_0)
                return (iife10_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_OperationsChanged:
                d_11_ops_ = source0_.operations
                def iife14_(_pat_let38_0):
                    def iife15_(d_12_dt__update__tmp_h3_):
                        def iife16_(_pat_let39_0):
                            def iife17_(d_13_dt__update_hoperations_h0_):
                                return ChangesetState_ChangesetState((d_12_dt__update__tmp_h3_).status, (d_12_dt__update__tmp_h3_).files, d_13_dt__update_hoperations_h0_, (d_12_dt__update__tmp_h3_).error)
                            return iife17_(_pat_let39_0)
                        return iife16_(d_11_ops_)
                    return iife15_(_pat_let38_0)
                return (iife14_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_Cleared:
                def iife18_(_pat_let40_0):
                    def iife19_(d_14_dt__update__tmp_h4_):
                        def iife20_(_pat_let41_0):
                            def iife21_(d_15_dt__update_hfiles_h2_):
                                return ChangesetState_ChangesetState((d_14_dt__update__tmp_h4_).status, d_15_dt__update_hfiles_h2_, (d_14_dt__update__tmp_h4_).operations, (d_14_dt__update__tmp_h4_).error)
                            return iife21_(_pat_let41_0)
                        return iife20_(_dafny.SeqWithoutIsStrInference([]))
                    return iife19_(_pat_let40_0)
                return (iife18_(s), (AhpSkeleton.ReduceOutcome_NoOp() if (len((s).files)) == (0) else AhpSkeleton.ReduceOutcome_Applied()))
        if True:
            if source0_.is_OperationStatusChanged:
                d_16_id_ = source0_.operationId
                d_17_st_ = source0_.status
                d_18_err_ = source0_.error
                source1_ = (s).operations
                if True:
                    if source1_.is_None:
                        return (s, AhpSkeleton.ReduceOutcome_NoOp())
                if True:
                    d_19_ops_ = source1_.value
                    if default__.hasOp(d_19_ops_, d_16_id_):
                        def iife22_(_pat_let42_0):
                            def iife23_(d_20_dt__update__tmp_h5_):
                                def iife24_(_pat_let43_0):
                                    def iife25_(d_21_dt__update_hoperations_h1_):
                                        return ChangesetState_ChangesetState((d_20_dt__update__tmp_h5_).status, (d_20_dt__update__tmp_h5_).files, d_21_dt__update_hoperations_h1_, (d_20_dt__update__tmp_h5_).error)
                                    return iife25_(_pat_let43_0)
                                return iife24_(AhpSkeleton.Option_Some(default__.updateOp(d_19_ops_, d_16_id_, d_17_st_, d_18_err_)))
                            return iife23_(_pat_let42_0)
                        return (iife22_(s), AhpSkeleton.ReduceOutcome_Applied())
                    elif True:
                        return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_ContentChanged:
                d_22_fs_ = source0_.files
                d_23_ops_ = source0_.operations
                d_24_err_ = source0_.error
                def iife26_(_pat_let44_0):
                    def iife27_(d_26_dt__update__tmp_h6_):
                        def iife28_(_pat_let45_0):
                            def iife29_(d_27_dt__update_hfiles_h3_):
                                return ChangesetState_ChangesetState((d_26_dt__update__tmp_h6_).status, d_27_dt__update_hfiles_h3_, (d_26_dt__update__tmp_h6_).operations, (d_26_dt__update__tmp_h6_).error)
                            return iife29_(_pat_let45_0)
                        return iife28_((d_22_fs_).value)
                    return iife27_(_pat_let44_0)
                d_25_s1_ = (iife26_(s) if (d_22_fs_).is_Some else s)
                def iife30_(_pat_let46_0):
                    def iife31_(d_29_dt__update__tmp_h7_):
                        def iife32_(_pat_let47_0):
                            def iife33_(d_30_dt__update_hoperations_h2_):
                                return ChangesetState_ChangesetState((d_29_dt__update__tmp_h7_).status, (d_29_dt__update__tmp_h7_).files, d_30_dt__update_hoperations_h2_, (d_29_dt__update__tmp_h7_).error)
                            return iife33_(_pat_let47_0)
                        return iife32_(d_23_ops_)
                    return iife31_(_pat_let46_0)
                d_28_s2_ = (iife30_(d_25_s1_) if (d_23_ops_).is_Some else d_25_s1_)
                def iife34_(_pat_let48_0):
                    def iife35_(d_32_dt__update__tmp_h8_):
                        def iife36_(_pat_let49_0):
                            def iife37_(d_33_dt__update_herror_h1_):
                                return ChangesetState_ChangesetState((d_32_dt__update__tmp_h8_).status, (d_32_dt__update__tmp_h8_).files, (d_32_dt__update__tmp_h8_).operations, d_33_dt__update_herror_h1_)
                            return iife37_(_pat_let49_0)
                        return iife36_(d_24_err_)
                    return iife35_(_pat_let48_0)
                d_31_s3_ = (iife34_(d_28_s2_) if (d_24_err_).is_Some else d_28_s2_)
                return (d_31_s3_, AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_FilesReviewedChanged:
                d_34_ids_ = source0_.fileIds
                d_35_rev_ = source0_.reviewed
                if default__.anyNeedsReviewed((s).files, d_34_ids_, d_35_rev_):
                    def iife38_(_pat_let50_0):
                        def iife39_(d_36_dt__update__tmp_h9_):
                            def iife40_(_pat_let51_0):
                                def iife41_(d_37_dt__update_hfiles_h4_):
                                    return ChangesetState_ChangesetState((d_36_dt__update__tmp_h9_).status, d_37_dt__update_hfiles_h4_, (d_36_dt__update__tmp_h9_).operations, (d_36_dt__update__tmp_h9_).error)
                                return iife41_(_pat_let51_0)
                            return iife40_(default__.setReviewed((pat_let_tv2_).files, d_34_ids_, d_35_rev_))
                        return iife39_(_pat_let50_0)
                    return (iife38_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            return (s, AhpSkeleton.ReduceOutcome_NoOp())

    @staticmethod
    def apply1(s, a):
        return (default__.ApplyToChangeset(s, a, 9999))[0]

    @staticmethod
    def fold(s, acts):
        return ConfluxContract.default__.Fold(default__.apply1, s, acts)

    @staticmethod
    def E():
        return ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "diff")): ConfluxCodec.Json_JNull()}))

    @staticmethod
    def RunCorpus():
        pass_: int = int(0)
        total: int = int(0)
        total = 15
        pass_ = 0
        d_0_fa_: ChangesetFile
        d_0_fa_ = ChangesetFile_ChangesetFile(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/a.ts")), AhpSkeleton.Option_None(), default__.E())
        d_1_fb_: ChangesetFile
        d_1_fb_ = ChangesetFile_ChangesetFile(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/b.ts")), AhpSkeleton.Option_None(), default__.E())
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "computing")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "boom"))))), ChangesetAction_StatusChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), AhpSkeleton.Option_None()))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChangesetAction_StatusChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "error")), AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))))))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "error")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))))):
            pass_ = (pass_) + (1)
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_0_fa_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChangesetAction_FileSet(d_1_fb_))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_0_fa_, d_1_fb_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        d_2_fbv2_: ChangesetFile
        d_2_fbv2_ = ChangesetFile_ChangesetFile(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/b.ts")), AhpSkeleton.Option_Some(True), default__.E())
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_0_fa_, d_1_fb_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChangesetAction_FileSet(d_2_fbv2_))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_0_fa_, d_2_fbv2_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_0_fa_, d_1_fb_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChangesetAction_FileRemoved(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/a.ts"))))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_1_fb_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_0_fa_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChangesetAction_FileRemoved(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nope"))))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_0_fa_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        d_3_op_: ChangesetOperation
        d_3_op_ = ChangesetOperation_ChangesetOperation(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "create-pr")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Create Pull Request")), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "changeset"))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idle")), AhpSkeleton.Option_None())
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChangesetAction_OperationsChanged(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([d_3_op_]))))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([d_3_op_])), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([d_3_op_])), AhpSkeleton.Option_None()), ChangesetAction_OperationsChanged(AhpSkeleton.Option_None()))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        if ((default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_0_fa_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChangesetAction_Cleared())).files) == (_dafny.SeqWithoutIsStrInference([])):
            pass_ = (pass_) + (1)
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChangesetAction_CsUnknown(ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "type")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "changeset/nope")))}))))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        d_4_opR_: ChangesetOperation
        d_5_dt__update__tmp_h0_ = d_3_op_
        d_6_dt__update_hstatus_h0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))
        d_4_opR_ = ChangesetOperation_ChangesetOperation((d_5_dt__update__tmp_h0_).id_, (d_5_dt__update__tmp_h0_).label__, (d_5_dt__update__tmp_h0_).scopes, d_6_dt__update_hstatus_h0_, (d_5_dt__update__tmp_h0_).error)
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([d_3_op_])), AhpSkeleton.Option_None()), ChangesetAction_OperationStatusChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "create-pr")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running")), AhpSkeleton.Option_None()))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([d_4_opR_])), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([d_3_op_])), AhpSkeleton.Option_None()), ChangesetAction_OperationStatusChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nope")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running")), AhpSkeleton.Option_None()))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([d_3_op_])), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_0_fa_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChangesetAction_ContentChanged(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([d_1_fb_])), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([d_3_op_])), AhpSkeleton.Option_None()))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_1_fb_]), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([d_3_op_])), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        d_7_fc_: ChangesetFile
        d_7_fc_ = ChangesetFile_ChangesetFile(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/c.ts")), AhpSkeleton.Option_Some(False), default__.E())
        d_8_fbT_: ChangesetFile
        d_8_fbT_ = ChangesetFile_ChangesetFile(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/b.ts")), AhpSkeleton.Option_Some(True), default__.E())
        d_9_faT_: ChangesetFile
        d_9_faT_ = ChangesetFile_ChangesetFile(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/a.ts")), AhpSkeleton.Option_Some(True), default__.E())
        d_10_r160_: ChangesetState
        d_10_r160_ = default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_0_fa_, d_8_fbT_, d_7_fc_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChangesetAction_FilesReviewedChanged(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/a.ts")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/b.ts")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/missing.ts"))]), True))
        if (d_10_r160_) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_9_faT_, d_8_fbT_, d_7_fc_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        if (default__.apply1(ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_8_fbT_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChangesetAction_FilesReviewedChanged(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///src/b.ts"))]), True))) == (ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")), _dafny.SeqWithoutIsStrInference([d_8_fbT_]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        return pass_, total

    @_dafny.classproperty
    def fileKey(instance):
        def lambda0_(d_0_x_):
            return (d_0_x_).id_

        return lambda0_
    @_dafny.classproperty
    def opKey(instance):
        def lambda0_(d_0_x_):
            return (d_0_x_).id_

        return lambda0_

class ChangesetFile:
    @classmethod
    def default(cls, ):
        return lambda: ChangesetFile_ChangesetFile(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()(), ConfluxCodec.Json.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ChangesetFile(self) -> bool:
        return isinstance(self, ChangesetFile_ChangesetFile)

class ChangesetFile_ChangesetFile(ChangesetFile, NamedTuple('ChangesetFile', [('id_', Any), ('reviewed', Any), ('edit', Any)])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetFile.ChangesetFile({self.id_.VerbatimString(True)}, {_dafny.string_of(self.reviewed)}, {_dafny.string_of(self.edit)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetFile_ChangesetFile) and self.id_ == __o.id_ and self.reviewed == __o.reviewed and self.edit == __o.edit
    def __hash__(self) -> int:
        return super().__hash__()


class ChangesetOperation:
    @classmethod
    def default(cls, ):
        return lambda: ChangesetOperation_ChangesetOperation(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ChangesetOperation(self) -> bool:
        return isinstance(self, ChangesetOperation_ChangesetOperation)

class ChangesetOperation_ChangesetOperation(ChangesetOperation, NamedTuple('ChangesetOperation', [('id_', Any), ('label__', Any), ('scopes', Any), ('status', Any), ('error', Any)])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetOperation.ChangesetOperation({self.id_.VerbatimString(True)}, {self.label__.VerbatimString(True)}, {_dafny.string_of(self.scopes)}, {self.status.VerbatimString(True)}, {_dafny.string_of(self.error)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetOperation_ChangesetOperation) and self.id_ == __o.id_ and self.label__ == __o.label__ and self.scopes == __o.scopes and self.status == __o.status and self.error == __o.error
    def __hash__(self) -> int:
        return super().__hash__()


class ChangesetState:
    @classmethod
    def default(cls, ):
        return lambda: ChangesetState_ChangesetState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ChangesetState(self) -> bool:
        return isinstance(self, ChangesetState_ChangesetState)

class ChangesetState_ChangesetState(ChangesetState, NamedTuple('ChangesetState', [('status', Any), ('files', Any), ('operations', Any), ('error', Any)])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetState.ChangesetState({self.status.VerbatimString(True)}, {_dafny.string_of(self.files)}, {_dafny.string_of(self.operations)}, {_dafny.string_of(self.error)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetState_ChangesetState) and self.status == __o.status and self.files == __o.files and self.operations == __o.operations and self.error == __o.error
    def __hash__(self) -> int:
        return super().__hash__()


class ChangesetAction:
    @classmethod
    def default(cls, ):
        return lambda: ChangesetAction_StatusChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_StatusChanged(self) -> bool:
        return isinstance(self, ChangesetAction_StatusChanged)
    @property
    def is_FileSet(self) -> bool:
        return isinstance(self, ChangesetAction_FileSet)
    @property
    def is_FileRemoved(self) -> bool:
        return isinstance(self, ChangesetAction_FileRemoved)
    @property
    def is_OperationsChanged(self) -> bool:
        return isinstance(self, ChangesetAction_OperationsChanged)
    @property
    def is_Cleared(self) -> bool:
        return isinstance(self, ChangesetAction_Cleared)
    @property
    def is_OperationStatusChanged(self) -> bool:
        return isinstance(self, ChangesetAction_OperationStatusChanged)
    @property
    def is_ContentChanged(self) -> bool:
        return isinstance(self, ChangesetAction_ContentChanged)
    @property
    def is_FilesReviewedChanged(self) -> bool:
        return isinstance(self, ChangesetAction_FilesReviewedChanged)
    @property
    def is_CsUnknown(self) -> bool:
        return isinstance(self, ChangesetAction_CsUnknown)

class ChangesetAction_StatusChanged(ChangesetAction, NamedTuple('StatusChanged', [('status', Any), ('error', Any)])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetAction.StatusChanged({self.status.VerbatimString(True)}, {_dafny.string_of(self.error)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetAction_StatusChanged) and self.status == __o.status and self.error == __o.error
    def __hash__(self) -> int:
        return super().__hash__()

class ChangesetAction_FileSet(ChangesetAction, NamedTuple('FileSet', [('file', Any)])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetAction.FileSet({_dafny.string_of(self.file)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetAction_FileSet) and self.file == __o.file
    def __hash__(self) -> int:
        return super().__hash__()

class ChangesetAction_FileRemoved(ChangesetAction, NamedTuple('FileRemoved', [('fileId', Any)])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetAction.FileRemoved({self.fileId.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetAction_FileRemoved) and self.fileId == __o.fileId
    def __hash__(self) -> int:
        return super().__hash__()

class ChangesetAction_OperationsChanged(ChangesetAction, NamedTuple('OperationsChanged', [('operations', Any)])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetAction.OperationsChanged({_dafny.string_of(self.operations)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetAction_OperationsChanged) and self.operations == __o.operations
    def __hash__(self) -> int:
        return super().__hash__()

class ChangesetAction_Cleared(ChangesetAction, NamedTuple('Cleared', [])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetAction.Cleared'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetAction_Cleared)
    def __hash__(self) -> int:
        return super().__hash__()

class ChangesetAction_OperationStatusChanged(ChangesetAction, NamedTuple('OperationStatusChanged', [('operationId', Any), ('status', Any), ('error', Any)])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetAction.OperationStatusChanged({self.operationId.VerbatimString(True)}, {self.status.VerbatimString(True)}, {_dafny.string_of(self.error)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetAction_OperationStatusChanged) and self.operationId == __o.operationId and self.status == __o.status and self.error == __o.error
    def __hash__(self) -> int:
        return super().__hash__()

class ChangesetAction_ContentChanged(ChangesetAction, NamedTuple('ContentChanged', [('files', Any), ('operations', Any), ('error', Any)])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetAction.ContentChanged({_dafny.string_of(self.files)}, {_dafny.string_of(self.operations)}, {_dafny.string_of(self.error)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetAction_ContentChanged) and self.files == __o.files and self.operations == __o.operations and self.error == __o.error
    def __hash__(self) -> int:
        return super().__hash__()

class ChangesetAction_FilesReviewedChanged(ChangesetAction, NamedTuple('FilesReviewedChanged', [('fileIds', Any), ('reviewed', Any)])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetAction.FilesReviewedChanged({_dafny.string_of(self.fileIds)}, {_dafny.string_of(self.reviewed)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetAction_FilesReviewedChanged) and self.fileIds == __o.fileIds and self.reviewed == __o.reviewed
    def __hash__(self) -> int:
        return super().__hash__()

class ChangesetAction_CsUnknown(ChangesetAction, NamedTuple('CsUnknown', [('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'Changeset.ChangesetAction.CsUnknown({_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChangesetAction_CsUnknown) and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()

