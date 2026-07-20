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
import Annotations as Annotations

# Module: Terminal

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def appendData(content, d):
        if (len(content)) == (0):
            return _dafny.SeqWithoutIsStrInference([Part_Unclassified(d)])
        elif True:
            d_0_last_ = (content)[(len(content)) - (1)]
            d_1_init_ = _dafny.SeqWithoutIsStrInference((content)[:(len(content)) - (1):])
            source0_ = d_0_last_
            if True:
                if source0_.is_Command:
                    d_2_cid_ = source0_.commandId
                    d_3_cl_ = source0_.commandLine
                    d_4_out_ = source0_.output
                    d_5_ts_ = source0_.timestamp
                    d_6_comp_ = source0_.isComplete
                    d_7_ec_ = source0_.exitCode
                    d_8_dm_ = source0_.durationMs
                    if not(d_6_comp_):
                        return (d_1_init_) + (_dafny.SeqWithoutIsStrInference([Part_Command(d_2_cid_, d_3_cl_, (d_4_out_) + (d), d_5_ts_, d_6_comp_, d_7_ec_, d_8_dm_)]))
                    elif True:
                        return (content) + (_dafny.SeqWithoutIsStrInference([Part_Unclassified(d)]))
            if True:
                d_9_v_ = source0_.value
                return (d_1_init_) + (_dafny.SeqWithoutIsStrInference([Part_Unclassified((d_9_v_) + (d))]))

    @staticmethod
    def finishCommand(content, id_, code, dur):
        def lambda0_(d_0_id_, d_1_code_, d_2_dur_):
            def lambda1_(d_3_p_):
                def lambda2_():
                    source0_ = d_3_p_
                    if True:
                        if source0_.is_Command:
                            d_4_cid_ = source0_.commandId
                            d_5_cl_ = source0_.commandLine
                            d_6_out_ = source0_.output
                            d_7_ts_ = source0_.timestamp
                            d_8_comp_ = source0_.isComplete
                            d_9_ec_ = source0_.exitCode
                            d_10_dm_ = source0_.durationMs
                            if (d_4_cid_) == (d_0_id_):
                                return Part_Command(d_4_cid_, d_5_cl_, d_6_out_, d_7_ts_, True, AhpSkeleton.Option_Some(d_1_code_), AhpSkeleton.Option_Some(d_2_dur_))
                            elif True:
                                return d_3_p_
                    if True:
                        return d_3_p_

                return lambda2_()

            return lambda1_

        return ConfluxOperators.default__.Map(lambda0_(id_, code, dur), content)

    @staticmethod
    def ApplyToTerminal(s, a, now):
        pat_let_tv0_ = s
        pat_let_tv1_ = s
        pat_let_tv2_ = s
        source0_ = a
        if True:
            if source0_.is_TCwdChanged:
                d_0_c_ = source0_.cwd
                def iife0_(_pat_let60_0):
                    def iife1_(d_1_dt__update__tmp_h0_):
                        def iife2_(_pat_let61_0):
                            def iife3_(d_2_dt__update_hcwd_h0_):
                                return TerminalState_TerminalState((d_1_dt__update__tmp_h0_).title, d_2_dt__update_hcwd_h0_, (d_1_dt__update__tmp_h0_).cols, (d_1_dt__update__tmp_h0_).rows, (d_1_dt__update__tmp_h0_).content, (d_1_dt__update__tmp_h0_).claim, (d_1_dt__update__tmp_h0_).exitCode, (d_1_dt__update__tmp_h0_).supportsCommandDetection)
                            return iife3_(_pat_let61_0)
                        return iife2_(AhpSkeleton.Option_Some(d_0_c_))
                    return iife1_(_pat_let60_0)
                return (iife0_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_TTitleChanged:
                d_3_t_ = source0_.title
                def iife4_(_pat_let62_0):
                    def iife5_(d_4_dt__update__tmp_h1_):
                        def iife6_(_pat_let63_0):
                            def iife7_(d_5_dt__update_htitle_h0_):
                                return TerminalState_TerminalState(d_5_dt__update_htitle_h0_, (d_4_dt__update__tmp_h1_).cwd, (d_4_dt__update__tmp_h1_).cols, (d_4_dt__update__tmp_h1_).rows, (d_4_dt__update__tmp_h1_).content, (d_4_dt__update__tmp_h1_).claim, (d_4_dt__update__tmp_h1_).exitCode, (d_4_dt__update__tmp_h1_).supportsCommandDetection)
                            return iife7_(_pat_let63_0)
                        return iife6_(d_3_t_)
                    return iife5_(_pat_let62_0)
                return (iife4_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_TResized:
                d_6_co_ = source0_.cols
                d_7_ro_ = source0_.rows
                def iife8_(_pat_let64_0):
                    def iife9_(d_8_dt__update__tmp_h2_):
                        def iife10_(_pat_let65_0):
                            def iife11_(d_9_dt__update_hrows_h0_):
                                def iife12_(_pat_let66_0):
                                    def iife13_(d_10_dt__update_hcols_h0_):
                                        return TerminalState_TerminalState((d_8_dt__update__tmp_h2_).title, (d_8_dt__update__tmp_h2_).cwd, d_10_dt__update_hcols_h0_, d_9_dt__update_hrows_h0_, (d_8_dt__update__tmp_h2_).content, (d_8_dt__update__tmp_h2_).claim, (d_8_dt__update__tmp_h2_).exitCode, (d_8_dt__update__tmp_h2_).supportsCommandDetection)
                                    return iife13_(_pat_let66_0)
                                return iife12_(AhpSkeleton.Option_Some(d_6_co_))
                            return iife11_(_pat_let65_0)
                        return iife10_(AhpSkeleton.Option_Some(d_7_ro_))
                    return iife9_(_pat_let64_0)
                return (iife8_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_TExited:
                d_11_code_ = source0_.code
                def iife14_(_pat_let67_0):
                    def iife15_(d_12_dt__update__tmp_h3_):
                        def iife16_(_pat_let68_0):
                            def iife17_(d_13_dt__update_hexitCode_h0_):
                                return TerminalState_TerminalState((d_12_dt__update__tmp_h3_).title, (d_12_dt__update__tmp_h3_).cwd, (d_12_dt__update__tmp_h3_).cols, (d_12_dt__update__tmp_h3_).rows, (d_12_dt__update__tmp_h3_).content, (d_12_dt__update__tmp_h3_).claim, d_13_dt__update_hexitCode_h0_, (d_12_dt__update__tmp_h3_).supportsCommandDetection)
                            return iife17_(_pat_let68_0)
                        return iife16_(AhpSkeleton.Option_Some(d_11_code_))
                    return iife15_(_pat_let67_0)
                return (iife14_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_TData:
                d_14_d_ = source0_.data
                def iife18_(_pat_let69_0):
                    def iife19_(d_15_dt__update__tmp_h4_):
                        def iife20_(_pat_let70_0):
                            def iife21_(d_16_dt__update_hcontent_h0_):
                                return TerminalState_TerminalState((d_15_dt__update__tmp_h4_).title, (d_15_dt__update__tmp_h4_).cwd, (d_15_dt__update__tmp_h4_).cols, (d_15_dt__update__tmp_h4_).rows, d_16_dt__update_hcontent_h0_, (d_15_dt__update__tmp_h4_).claim, (d_15_dt__update__tmp_h4_).exitCode, (d_15_dt__update__tmp_h4_).supportsCommandDetection)
                            return iife21_(_pat_let70_0)
                        return iife20_(default__.appendData((pat_let_tv0_).content, d_14_d_))
                    return iife19_(_pat_let69_0)
                return (iife18_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_TCleared:
                def iife22_(_pat_let71_0):
                    def iife23_(d_17_dt__update__tmp_h5_):
                        def iife24_(_pat_let72_0):
                            def iife25_(d_18_dt__update_hcontent_h1_):
                                return TerminalState_TerminalState((d_17_dt__update__tmp_h5_).title, (d_17_dt__update__tmp_h5_).cwd, (d_17_dt__update__tmp_h5_).cols, (d_17_dt__update__tmp_h5_).rows, d_18_dt__update_hcontent_h1_, (d_17_dt__update__tmp_h5_).claim, (d_17_dt__update__tmp_h5_).exitCode, (d_17_dt__update__tmp_h5_).supportsCommandDetection)
                            return iife25_(_pat_let72_0)
                        return iife24_(_dafny.SeqWithoutIsStrInference([]))
                    return iife23_(_pat_let71_0)
                return (iife22_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_TClaimed:
                d_19_c_ = source0_.claim
                def iife26_(_pat_let73_0):
                    def iife27_(d_20_dt__update__tmp_h6_):
                        def iife28_(_pat_let74_0):
                            def iife29_(d_21_dt__update_hclaim_h0_):
                                return TerminalState_TerminalState((d_20_dt__update__tmp_h6_).title, (d_20_dt__update__tmp_h6_).cwd, (d_20_dt__update__tmp_h6_).cols, (d_20_dt__update__tmp_h6_).rows, (d_20_dt__update__tmp_h6_).content, d_21_dt__update_hclaim_h0_, (d_20_dt__update__tmp_h6_).exitCode, (d_20_dt__update__tmp_h6_).supportsCommandDetection)
                            return iife29_(_pat_let74_0)
                        return iife28_(AhpSkeleton.Option_Some(d_19_c_))
                    return iife27_(_pat_let73_0)
                return (iife26_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_TCommandDetectionAvailable:
                def iife30_(_pat_let75_0):
                    def iife31_(d_22_dt__update__tmp_h7_):
                        def iife32_(_pat_let76_0):
                            def iife33_(d_23_dt__update_hsupportsCommandDetection_h0_):
                                return TerminalState_TerminalState((d_22_dt__update__tmp_h7_).title, (d_22_dt__update__tmp_h7_).cwd, (d_22_dt__update__tmp_h7_).cols, (d_22_dt__update__tmp_h7_).rows, (d_22_dt__update__tmp_h7_).content, (d_22_dt__update__tmp_h7_).claim, (d_22_dt__update__tmp_h7_).exitCode, d_23_dt__update_hsupportsCommandDetection_h0_)
                            return iife33_(_pat_let76_0)
                        return iife32_(AhpSkeleton.Option_Some(True))
                    return iife31_(_pat_let75_0)
                return (iife30_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_TCommandExecuted:
                d_24_id_ = source0_.commandId
                d_25_cl_ = source0_.commandLine
                d_26_ts_ = source0_.timestamp
                def iife34_(_pat_let77_0):
                    def iife35_(d_27_dt__update__tmp_h8_):
                        def iife36_(_pat_let78_0):
                            def iife37_(d_28_dt__update_hsupportsCommandDetection_h1_):
                                def iife38_(_pat_let79_0):
                                    def iife39_(d_29_dt__update_hcontent_h2_):
                                        return TerminalState_TerminalState((d_27_dt__update__tmp_h8_).title, (d_27_dt__update__tmp_h8_).cwd, (d_27_dt__update__tmp_h8_).cols, (d_27_dt__update__tmp_h8_).rows, d_29_dt__update_hcontent_h2_, (d_27_dt__update__tmp_h8_).claim, (d_27_dt__update__tmp_h8_).exitCode, d_28_dt__update_hsupportsCommandDetection_h1_)
                                    return iife39_(_pat_let79_0)
                                return iife38_(((pat_let_tv1_).content) + (_dafny.SeqWithoutIsStrInference([Part_Command(d_24_id_, d_25_cl_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), d_26_ts_, False, AhpSkeleton.Option_None(), AhpSkeleton.Option_None())])))
                            return iife37_(_pat_let78_0)
                        return iife36_(AhpSkeleton.Option_Some(True))
                    return iife35_(_pat_let77_0)
                return (iife34_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_TCommandFinished:
                d_30_id_ = source0_.commandId
                d_31_code_ = source0_.code
                d_32_dur_ = source0_.durationMs
                def iife40_(_pat_let80_0):
                    def iife41_(d_33_dt__update__tmp_h9_):
                        def iife42_(_pat_let81_0):
                            def iife43_(d_34_dt__update_hcontent_h3_):
                                return TerminalState_TerminalState((d_33_dt__update__tmp_h9_).title, (d_33_dt__update__tmp_h9_).cwd, (d_33_dt__update__tmp_h9_).cols, (d_33_dt__update__tmp_h9_).rows, d_34_dt__update_hcontent_h3_, (d_33_dt__update__tmp_h9_).claim, (d_33_dt__update__tmp_h9_).exitCode, (d_33_dt__update__tmp_h9_).supportsCommandDetection)
                            return iife43_(_pat_let81_0)
                        return iife42_(default__.finishCommand((pat_let_tv2_).content, d_30_id_, d_31_code_, d_32_dur_))
                    return iife41_(_pat_let80_0)
                return (iife40_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_TInput:
                return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            return (s, AhpSkeleton.ReduceOutcome_NoOp())

    @staticmethod
    def apply1(s, a):
        return (default__.ApplyToTerminal(s, a, 9999))[0]

    @staticmethod
    def fold(s, acts):
        return ConfluxContract.default__.Fold(default__.apply1, s, acts)

    @staticmethod
    def T0():
        return TerminalState_TerminalState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bash")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())

    @staticmethod
    def CL():
        return ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "client"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clientId")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c1")))}))

    @staticmethod
    def RunCorpus():
        pass_: int = int(0)
        total: int = int(0)
        total = 19
        pass_ = 0
        d_0_base_: TerminalState
        d_1_dt__update__tmp_h0_ = default__.T0()
        d_2_dt__update_hclaim_h0_ = AhpSkeleton.Option_Some(default__.CL())
        d_0_base_ = TerminalState_TerminalState((d_1_dt__update__tmp_h0_).title, (d_1_dt__update__tmp_h0_).cwd, (d_1_dt__update__tmp_h0_).cols, (d_1_dt__update__tmp_h0_).rows, (d_1_dt__update__tmp_h0_).content, d_2_dt__update_hclaim_h0_, (d_1_dt__update__tmp_h0_).exitCode, (d_1_dt__update__tmp_h0_).supportsCommandDetection)
        if ((default__.apply1(d_0_base_, TerminalAction_TCwdChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/tmp"))))).cwd) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/tmp")))):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_base_, TerminalAction_TTitleChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "zsh"))))).title) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "zsh"))):
            pass_ = (pass_) + (1)
        def iife0_(_pat_let82_0):
            def iife1_(d_3_dt__update__tmp_h1_):
                def iife2_(_pat_let83_0):
                    def iife3_(d_4_dt__update_hrows_h0_):
                        def iife4_(_pat_let84_0):
                            def iife5_(d_5_dt__update_hcols_h0_):
                                return TerminalState_TerminalState((d_3_dt__update__tmp_h1_).title, (d_3_dt__update__tmp_h1_).cwd, d_5_dt__update_hcols_h0_, d_4_dt__update_hrows_h0_, (d_3_dt__update__tmp_h1_).content, (d_3_dt__update__tmp_h1_).claim, (d_3_dt__update__tmp_h1_).exitCode, (d_3_dt__update__tmp_h1_).supportsCommandDetection)
                            return iife5_(_pat_let84_0)
                        return iife4_(AhpSkeleton.Option_Some(80))
                    return iife3_(_pat_let83_0)
                return iife2_(AhpSkeleton.Option_Some(24))
            return iife1_(_pat_let82_0)
        if (default__.apply1(d_0_base_, TerminalAction_TResized(80, 24))) == (iife0_(d_0_base_)):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_base_, TerminalAction_TExited(0))).exitCode) == (AhpSkeleton.Option_Some(0)):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_0_base_, TerminalAction_TInput())) == (d_0_base_):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_0_base_, TerminalAction_TUnknown(ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "type")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "terminal/nope")))}))))) == (d_0_base_):
            pass_ = (pass_) + (1)
        def iife6_(_pat_let85_0):
            def iife7_(d_6_dt__update__tmp_h2_):
                def iife8_(_pat_let86_0):
                    def iife9_(d_7_dt__update_hcontent_h0_):
                        return TerminalState_TerminalState((d_6_dt__update__tmp_h2_).title, (d_6_dt__update__tmp_h2_).cwd, (d_6_dt__update__tmp_h2_).cols, (d_6_dt__update__tmp_h2_).rows, d_7_dt__update_hcontent_h0_, (d_6_dt__update__tmp_h2_).claim, (d_6_dt__update__tmp_h2_).exitCode, (d_6_dt__update__tmp_h2_).supportsCommandDetection)
                    return iife9_(_pat_let86_0)
                return iife8_(_dafny.SeqWithoutIsStrInference([Part_Unclassified(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))]))
            return iife7_(_pat_let85_0)
        if ((default__.apply1(iife6_(d_0_base_), TerminalAction_TCleared())).content) == (_dafny.SeqWithoutIsStrInference([])):
            pass_ = (pass_) + (1)
        def iife10_(_pat_let87_0):
            def iife11_(d_8_dt__update__tmp_h3_):
                def iife12_(_pat_let88_0):
                    def iife13_(d_9_dt__update_hsupportsCommandDetection_h0_):
                        def iife14_(_pat_let89_0):
                            def iife15_(d_10_dt__update_hcontent_h1_):
                                return TerminalState_TerminalState((d_8_dt__update__tmp_h3_).title, (d_8_dt__update__tmp_h3_).cwd, (d_8_dt__update__tmp_h3_).cols, (d_8_dt__update__tmp_h3_).rows, d_10_dt__update_hcontent_h1_, (d_8_dt__update__tmp_h3_).claim, (d_8_dt__update__tmp_h3_).exitCode, d_9_dt__update_hsupportsCommandDetection_h0_)
                            return iife15_(_pat_let89_0)
                        return iife14_(_dafny.SeqWithoutIsStrInference([Part_Unclassified(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))]))
                    return iife13_(_pat_let88_0)
                return iife12_(AhpSkeleton.Option_Some(True))
            return iife11_(_pat_let87_0)
        def iife16_(_pat_let90_0):
            def iife17_(d_11_dt__update__tmp_h4_):
                def iife18_(_pat_let91_0):
                    def iife19_(d_12_dt__update_hsupportsCommandDetection_h1_):
                        def iife20_(_pat_let92_0):
                            def iife21_(d_13_dt__update_hcontent_h2_):
                                return TerminalState_TerminalState((d_11_dt__update__tmp_h4_).title, (d_11_dt__update__tmp_h4_).cwd, (d_11_dt__update__tmp_h4_).cols, (d_11_dt__update__tmp_h4_).rows, d_13_dt__update_hcontent_h2_, (d_11_dt__update__tmp_h4_).claim, (d_11_dt__update__tmp_h4_).exitCode, d_12_dt__update_hsupportsCommandDetection_h1_)
                            return iife21_(_pat_let92_0)
                        return iife20_(_dafny.SeqWithoutIsStrInference([]))
                    return iife19_(_pat_let91_0)
                return iife18_(AhpSkeleton.Option_Some(True))
            return iife17_(_pat_let90_0)
        if (default__.apply1(iife10_(d_0_base_), TerminalAction_TCleared())) == (iife16_(d_0_base_)):
            pass_ = (pass_) + (1)
        d_14_sc_: ConfluxCodec.Json
        d_14_sc_ = ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "session"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "session")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s1")))}))
        if ((default__.apply1(d_0_base_, TerminalAction_TClaimed(d_14_sc_))).claim) == (AhpSkeleton.Option_Some(d_14_sc_)):
            pass_ = (pass_) + (1)
        d_15_scTool_: ConfluxCodec.Json
        d_15_scTool_ = ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "session"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "session")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s1"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "toolCallId")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))}))
        if ((default__.apply1(d_0_base_, TerminalAction_TClaimed(d_15_scTool_))).claim) == (AhpSkeleton.Option_Some(d_15_scTool_)):
            pass_ = (pass_) + (1)
        pat_let_tv0_ = d_14_sc_
        def iife22_(_pat_let93_0):
            def iife23_(d_16_dt__update__tmp_h5_):
                def iife24_(_pat_let94_0):
                    def iife25_(d_17_dt__update_hclaim_h1_):
                        return TerminalState_TerminalState((d_16_dt__update__tmp_h5_).title, (d_16_dt__update__tmp_h5_).cwd, (d_16_dt__update__tmp_h5_).cols, (d_16_dt__update__tmp_h5_).rows, (d_16_dt__update__tmp_h5_).content, d_17_dt__update_hclaim_h1_, (d_16_dt__update__tmp_h5_).exitCode, (d_16_dt__update__tmp_h5_).supportsCommandDetection)
                    return iife25_(_pat_let94_0)
                return iife24_(AhpSkeleton.Option_Some(pat_let_tv0_))
            return iife23_(_pat_let93_0)
        if ((default__.apply1(iife22_(d_0_base_), TerminalAction_TClaimed(default__.CL()))).claim) == (AhpSkeleton.Option_Some(default__.CL())):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_base_, TerminalAction_TCommandDetectionAvailable())).supportsCommandDetection) == (AhpSkeleton.Option_Some(True)):
            pass_ = (pass_) + (1)
        d_18_afterExec_: TerminalState
        d_18_afterExec_ = default__.apply1(d_0_base_, TerminalAction_TCommandExecuted(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npm test")), 1700000000000))
        def iife26_(_pat_let95_0):
            def iife27_(d_19_dt__update__tmp_h6_):
                def iife28_(_pat_let96_0):
                    def iife29_(d_20_dt__update_hsupportsCommandDetection_h2_):
                        def iife30_(_pat_let97_0):
                            def iife31_(d_21_dt__update_hcontent_h3_):
                                return TerminalState_TerminalState((d_19_dt__update__tmp_h6_).title, (d_19_dt__update__tmp_h6_).cwd, (d_19_dt__update__tmp_h6_).cols, (d_19_dt__update__tmp_h6_).rows, d_21_dt__update_hcontent_h3_, (d_19_dt__update__tmp_h6_).claim, (d_19_dt__update__tmp_h6_).exitCode, d_20_dt__update_hsupportsCommandDetection_h2_)
                            return iife31_(_pat_let97_0)
                        return iife30_(_dafny.SeqWithoutIsStrInference([Part_Command(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npm test")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1700000000000, False, AhpSkeleton.Option_None(), AhpSkeleton.Option_None())]))
                    return iife29_(_pat_let96_0)
                return iife28_(AhpSkeleton.Option_Some(True))
            return iife27_(_pat_let95_0)
        if (d_18_afterExec_) == (iife26_(d_0_base_)):
            pass_ = (pass_) + (1)
        d_22_withCmd_: TerminalState
        d_23_dt__update__tmp_h7_ = d_0_base_
        d_24_dt__update_hsupportsCommandDetection_h3_ = AhpSkeleton.Option_Some(True)
        d_25_dt__update_hcontent_h4_ = _dafny.SeqWithoutIsStrInference([Part_Command(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npm test")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "All tests passed\r\n")), 1700000000000, False, AhpSkeleton.Option_None(), AhpSkeleton.Option_None())])
        d_22_withCmd_ = TerminalState_TerminalState((d_23_dt__update__tmp_h7_).title, (d_23_dt__update__tmp_h7_).cwd, (d_23_dt__update__tmp_h7_).cols, (d_23_dt__update__tmp_h7_).rows, d_25_dt__update_hcontent_h4_, (d_23_dt__update__tmp_h7_).claim, (d_23_dt__update__tmp_h7_).exitCode, d_24_dt__update_hsupportsCommandDetection_h3_)
        d_26_doneCmd_: TerminalState
        d_27_dt__update__tmp_h8_ = d_0_base_
        d_28_dt__update_hsupportsCommandDetection_h4_ = AhpSkeleton.Option_Some(True)
        d_29_dt__update_hcontent_h5_ = _dafny.SeqWithoutIsStrInference([Part_Command(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npm test")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "All tests passed\r\n")), 1700000000000, True, AhpSkeleton.Option_Some(0), AhpSkeleton.Option_Some(1234))])
        d_26_doneCmd_ = TerminalState_TerminalState((d_27_dt__update__tmp_h8_).title, (d_27_dt__update__tmp_h8_).cwd, (d_27_dt__update__tmp_h8_).cols, (d_27_dt__update__tmp_h8_).rows, d_29_dt__update_hcontent_h5_, (d_27_dt__update__tmp_h8_).claim, (d_27_dt__update__tmp_h8_).exitCode, d_28_dt__update_hsupportsCommandDetection_h4_)
        if (default__.apply1(d_22_withCmd_, TerminalAction_TCommandFinished(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-1")), 0, 1234))) == (d_26_doneCmd_):
            pass_ = (pass_) + (1)
        def iife32_(_pat_let98_0):
            def iife33_(d_30_dt__update__tmp_h9_):
                def iife34_(_pat_let99_0):
                    def iife35_(d_31_dt__update_hcontent_h6_):
                        return TerminalState_TerminalState((d_30_dt__update__tmp_h9_).title, (d_30_dt__update__tmp_h9_).cwd, (d_30_dt__update__tmp_h9_).cols, (d_30_dt__update__tmp_h9_).rows, d_31_dt__update_hcontent_h6_, (d_30_dt__update__tmp_h9_).claim, (d_30_dt__update__tmp_h9_).exitCode, (d_30_dt__update__tmp_h9_).supportsCommandDetection)
                    return iife35_(_pat_let99_0)
                return iife34_(_dafny.SeqWithoutIsStrInference([Part_Unclassified(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))]))
            return iife33_(_pat_let98_0)
        if ((default__.apply1(iife32_(d_0_base_), TerminalAction_TData(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))))).content) == (_dafny.SeqWithoutIsStrInference([Part_Unclassified(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ab")))])):
            pass_ = (pass_) + (1)
        d_32_complete_: TerminalState
        d_33_dt__update__tmp_h10_ = d_0_base_
        d_34_dt__update_hsupportsCommandDetection_h5_ = AhpSkeleton.Option_Some(True)
        d_35_dt__update_hcontent_h7_ = _dafny.SeqWithoutIsStrInference([Part_Command(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "echo hi")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hi\r\n")), 1700000000000, True, AhpSkeleton.Option_Some(0), AhpSkeleton.Option_Some(50))])
        d_32_complete_ = TerminalState_TerminalState((d_33_dt__update__tmp_h10_).title, (d_33_dt__update__tmp_h10_).cwd, (d_33_dt__update__tmp_h10_).cols, (d_33_dt__update__tmp_h10_).rows, d_35_dt__update_hcontent_h7_, (d_33_dt__update__tmp_h10_).claim, (d_33_dt__update__tmp_h10_).exitCode, d_34_dt__update_hsupportsCommandDetection_h5_)
        if ((default__.apply1(d_32_complete_, TerminalAction_TData(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "$ "))))).content) == (((d_32_complete_).content) + (_dafny.SeqWithoutIsStrInference([Part_Unclassified(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "$ ")))]))):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_22_withCmd_, TerminalAction_TData(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "!"))))).content) == (_dafny.SeqWithoutIsStrInference([Part_Command(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npm test")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "All tests passed\r\n!")), 1700000000000, False, AhpSkeleton.Option_None(), AhpSkeleton.Option_None())])):
            pass_ = (pass_) + (1)
        d_36_life_: TerminalState
        d_36_life_ = default__.fold(d_0_base_, _dafny.SeqWithoutIsStrInference([TerminalAction_TCwdChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/w"))), TerminalAction_TData(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))), TerminalAction_TResized(100, 40), TerminalAction_TClaimed(d_14_sc_), TerminalAction_TExited(1)]))
        if (((((d_36_life_).cwd) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/w"))))) and (((d_36_life_).cols) == (AhpSkeleton.Option_Some(100)))) and (((d_36_life_).exitCode) == (AhpSkeleton.Option_Some(1)))) and (((d_36_life_).claim) == (AhpSkeleton.Option_Some(d_14_sc_))):
            pass_ = (pass_) + (1)
        d_37_lc_: TerminalState
        def iife36_(_pat_let100_0):
            def iife37_(d_38_dt__update__tmp_h11_):
                def iife38_(_pat_let101_0):
                    def iife39_(d_39_dt__update_hcontent_h8_):
                        return TerminalState_TerminalState((d_38_dt__update__tmp_h11_).title, (d_38_dt__update__tmp_h11_).cwd, (d_38_dt__update__tmp_h11_).cols, (d_38_dt__update__tmp_h11_).rows, d_39_dt__update_hcontent_h8_, (d_38_dt__update__tmp_h11_).claim, (d_38_dt__update__tmp_h11_).exitCode, (d_38_dt__update__tmp_h11_).supportsCommandDetection)
                    return iife39_(_pat_let101_0)
                return iife38_(_dafny.SeqWithoutIsStrInference([]))
            return iife37_(_pat_let100_0)
        d_37_lc_ = default__.fold(iife36_(d_0_base_), _dafny.SeqWithoutIsStrInference([TerminalAction_TData(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "$ "))), TerminalAction_TCommandExecuted(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "echo hello")), 1700000000000), TerminalAction_TData(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hello\r\n"))), TerminalAction_TCommandFinished(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-1")), 0, 10), TerminalAction_TData(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "$ "))), TerminalAction_TCommandExecuted(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-2")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exit 1")), 1700000001000), TerminalAction_TData(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))), TerminalAction_TCommandFinished(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-2")), 1, 5)]))
        d_40_expLc_: _dafny.Seq
        d_40_expLc_ = _dafny.SeqWithoutIsStrInference([Part_Unclassified(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "$ "))), Part_Command(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "echo hello")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hello\r\n")), 1700000000000, True, AhpSkeleton.Option_Some(0), AhpSkeleton.Option_Some(10)), Part_Unclassified(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "$ "))), Part_Command(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmd-2")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exit 1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1700000001000, True, AhpSkeleton.Option_Some(1), AhpSkeleton.Option_Some(5))])
        if (((d_37_lc_).content) == (d_40_expLc_)) and (((d_37_lc_).supportsCommandDetection) == (AhpSkeleton.Option_Some(True))):
            pass_ = (pass_) + (1)
        return pass_, total


class Part:
    @classmethod
    def default(cls, ):
        return lambda: Part_Unclassified(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Unclassified(self) -> bool:
        return isinstance(self, Part_Unclassified)
    @property
    def is_Command(self) -> bool:
        return isinstance(self, Part_Command)

class Part_Unclassified(Part, NamedTuple('Unclassified', [('value', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.Part.Unclassified({self.value.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Part_Unclassified) and self.value == __o.value
    def __hash__(self) -> int:
        return super().__hash__()

class Part_Command(Part, NamedTuple('Command', [('commandId', Any), ('commandLine', Any), ('output', Any), ('timestamp', Any), ('isComplete', Any), ('exitCode', Any), ('durationMs', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.Part.Command({self.commandId.VerbatimString(True)}, {self.commandLine.VerbatimString(True)}, {self.output.VerbatimString(True)}, {_dafny.string_of(self.timestamp)}, {_dafny.string_of(self.isComplete)}, {_dafny.string_of(self.exitCode)}, {_dafny.string_of(self.durationMs)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Part_Command) and self.commandId == __o.commandId and self.commandLine == __o.commandLine and self.output == __o.output and self.timestamp == __o.timestamp and self.isComplete == __o.isComplete and self.exitCode == __o.exitCode and self.durationMs == __o.durationMs
    def __hash__(self) -> int:
        return super().__hash__()


class TerminalState:
    @classmethod
    def default(cls, ):
        return lambda: TerminalState_TerminalState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), _dafny.Seq({}), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_TerminalState(self) -> bool:
        return isinstance(self, TerminalState_TerminalState)

class TerminalState_TerminalState(TerminalState, NamedTuple('TerminalState', [('title', Any), ('cwd', Any), ('cols', Any), ('rows', Any), ('content', Any), ('claim', Any), ('exitCode', Any), ('supportsCommandDetection', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalState.TerminalState({self.title.VerbatimString(True)}, {_dafny.string_of(self.cwd)}, {_dafny.string_of(self.cols)}, {_dafny.string_of(self.rows)}, {_dafny.string_of(self.content)}, {_dafny.string_of(self.claim)}, {_dafny.string_of(self.exitCode)}, {_dafny.string_of(self.supportsCommandDetection)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalState_TerminalState) and self.title == __o.title and self.cwd == __o.cwd and self.cols == __o.cols and self.rows == __o.rows and self.content == __o.content and self.claim == __o.claim and self.exitCode == __o.exitCode and self.supportsCommandDetection == __o.supportsCommandDetection
    def __hash__(self) -> int:
        return super().__hash__()


class TerminalAction:
    @classmethod
    def default(cls, ):
        return lambda: TerminalAction_TCwdChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_TCwdChanged(self) -> bool:
        return isinstance(self, TerminalAction_TCwdChanged)
    @property
    def is_TTitleChanged(self) -> bool:
        return isinstance(self, TerminalAction_TTitleChanged)
    @property
    def is_TResized(self) -> bool:
        return isinstance(self, TerminalAction_TResized)
    @property
    def is_TExited(self) -> bool:
        return isinstance(self, TerminalAction_TExited)
    @property
    def is_TData(self) -> bool:
        return isinstance(self, TerminalAction_TData)
    @property
    def is_TCleared(self) -> bool:
        return isinstance(self, TerminalAction_TCleared)
    @property
    def is_TClaimed(self) -> bool:
        return isinstance(self, TerminalAction_TClaimed)
    @property
    def is_TCommandDetectionAvailable(self) -> bool:
        return isinstance(self, TerminalAction_TCommandDetectionAvailable)
    @property
    def is_TCommandExecuted(self) -> bool:
        return isinstance(self, TerminalAction_TCommandExecuted)
    @property
    def is_TCommandFinished(self) -> bool:
        return isinstance(self, TerminalAction_TCommandFinished)
    @property
    def is_TInput(self) -> bool:
        return isinstance(self, TerminalAction_TInput)
    @property
    def is_TUnknown(self) -> bool:
        return isinstance(self, TerminalAction_TUnknown)

class TerminalAction_TCwdChanged(TerminalAction, NamedTuple('TCwdChanged', [('cwd', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TCwdChanged({self.cwd.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TCwdChanged) and self.cwd == __o.cwd
    def __hash__(self) -> int:
        return super().__hash__()

class TerminalAction_TTitleChanged(TerminalAction, NamedTuple('TTitleChanged', [('title', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TTitleChanged({self.title.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TTitleChanged) and self.title == __o.title
    def __hash__(self) -> int:
        return super().__hash__()

class TerminalAction_TResized(TerminalAction, NamedTuple('TResized', [('cols', Any), ('rows', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TResized({_dafny.string_of(self.cols)}, {_dafny.string_of(self.rows)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TResized) and self.cols == __o.cols and self.rows == __o.rows
    def __hash__(self) -> int:
        return super().__hash__()

class TerminalAction_TExited(TerminalAction, NamedTuple('TExited', [('code', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TExited({_dafny.string_of(self.code)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TExited) and self.code == __o.code
    def __hash__(self) -> int:
        return super().__hash__()

class TerminalAction_TData(TerminalAction, NamedTuple('TData', [('data', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TData({self.data.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TData) and self.data == __o.data
    def __hash__(self) -> int:
        return super().__hash__()

class TerminalAction_TCleared(TerminalAction, NamedTuple('TCleared', [])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TCleared'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TCleared)
    def __hash__(self) -> int:
        return super().__hash__()

class TerminalAction_TClaimed(TerminalAction, NamedTuple('TClaimed', [('claim', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TClaimed({_dafny.string_of(self.claim)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TClaimed) and self.claim == __o.claim
    def __hash__(self) -> int:
        return super().__hash__()

class TerminalAction_TCommandDetectionAvailable(TerminalAction, NamedTuple('TCommandDetectionAvailable', [])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TCommandDetectionAvailable'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TCommandDetectionAvailable)
    def __hash__(self) -> int:
        return super().__hash__()

class TerminalAction_TCommandExecuted(TerminalAction, NamedTuple('TCommandExecuted', [('commandId', Any), ('commandLine', Any), ('timestamp', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TCommandExecuted({self.commandId.VerbatimString(True)}, {self.commandLine.VerbatimString(True)}, {_dafny.string_of(self.timestamp)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TCommandExecuted) and self.commandId == __o.commandId and self.commandLine == __o.commandLine and self.timestamp == __o.timestamp
    def __hash__(self) -> int:
        return super().__hash__()

class TerminalAction_TCommandFinished(TerminalAction, NamedTuple('TCommandFinished', [('commandId', Any), ('code', Any), ('durationMs', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TCommandFinished({self.commandId.VerbatimString(True)}, {_dafny.string_of(self.code)}, {_dafny.string_of(self.durationMs)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TCommandFinished) and self.commandId == __o.commandId and self.code == __o.code and self.durationMs == __o.durationMs
    def __hash__(self) -> int:
        return super().__hash__()

class TerminalAction_TInput(TerminalAction, NamedTuple('TInput', [])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TInput'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TInput)
    def __hash__(self) -> int:
        return super().__hash__()

class TerminalAction_TUnknown(TerminalAction, NamedTuple('TUnknown', [('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'Terminal.TerminalAction.TUnknown({_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TerminalAction_TUnknown) and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()

