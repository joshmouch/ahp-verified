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
import Terminal as Terminal
import Session as Session
import ConfluxOrderedLog as ConfluxOrderedLog

# Module: Chat

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def setBit(s, b):
        return (s) | (b)

    @staticmethod
    def clearBit(s, b):
        return (s) & ((~(b)) & ((1 << 32) - 1))

    @staticmethod
    def jNoNull(j):
        if (j).is_JNull:
            return AhpSkeleton.Option_None()
        elif True:
            return AhpSkeleton.Option_Some(j)

    @staticmethod
    def upsertQ(qs, m):
        return ConfluxSeqRoute.default__.SeqUpsertById(default__.qKey, qs, m)

    @staticmethod
    def removeQ(qs, id_):
        return ConfluxSeqRoute.default__.SeqRemoveById(default__.qKey, qs, id_)

    @staticmethod
    def containsQ(qs, id_):
        while True:
            with _dafny.label():
                if (len(qs)) == (0):
                    return False
                elif (((qs)[0]).id_) == (id_):
                    return True
                elif True:
                    in0_ = _dafny.SeqWithoutIsStrInference((qs)[1::])
                    in1_ = id_
                    qs = in0_
                    id_ = in1_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def findQ(qs, id_):
        def lambda0_(d_1_id_):
            def lambda1_(d_2_q_):
                return ((d_2_q_).id_) == (d_1_id_)

            return lambda1_

        d_0_found_ = ConfluxOperators.default__.Filter(lambda0_(id_), qs)
        if (len(d_0_found_)) == (0):
            return _dafny.SeqWithoutIsStrInference([])
        elif True:
            return _dafny.SeqWithoutIsStrInference([(d_0_found_)[0]])

    @staticmethod
    def reorderQ(qs, order):
        return ConfluxOrderedLog.default__.SeqReorderByKeys(default__.qKey, order, qs)

    @staticmethod
    def keepUpTo(ts, id_):
        def lambda0_(d_0_t_):
            return (d_0_t_).turnId

        return ConfluxOrderedLog.default__.SeqKeepThrough(lambda0_, id_, ts)

    @staticmethod
    def hasTurn(ts, id_):
        def lambda0_(exists_var_0_):
            d_0_i_: int = exists_var_0_
            return (((0) <= (d_0_i_)) and ((d_0_i_) < (len(ts)))) and ((((ts)[d_0_i_]).turnId) == (id_))

        return _dafny.quantifier(_dafny.IntegerRange(0, len(ts)), False, lambda0_)

    @staticmethod
    def dedupPrepend(loaded, existing):
        def lambda0_(d_0_existing_):
            def lambda1_(d_1_t_):
                return not(default__.hasTurn(d_0_existing_, (d_1_t_).turnId))

            return lambda1_

        return (ConfluxOperators.default__.Filter(lambda0_(existing), loaded)) + (existing)

    @staticmethod
    def isOpenInput(p):
        return ((p).is_PInputRequest) and (((p).response).is_None)

    @staticmethod
    def hasOpenReq(ps, id_):
        if (len(ps)) == (0):
            return False
        elif True:
            return (((((ps)[0]).is_PInputRequest) and ((((ps)[0]).response).is_None)) and (((((ps)[0]).req).id_) == (id_))) or (default__.hasOpenReq(_dafny.SeqWithoutIsStrInference((ps)[1::]), id_))

    @staticmethod
    def mergeInputReq(existing, req):
        if (len((req).answers)) > (0):
            return req
        elif True:
            d_0_dt__update__tmp_h0_ = req
            d_1_dt__update_hanswers_h0_ = (existing).answers
            return InputReq_InputReq((d_0_dt__update__tmp_h0_).id_, d_1_dt__update_hanswers_h0_, (d_0_dt__update__tmp_h0_).open_)

    @staticmethod
    def mergeInputPart(req):
        def lambda0_(d_0_req_):
            def lambda1_(d_1_p_):
                def lambda2_():
                    source0_ = d_1_p_
                    if True:
                        if source0_.is_PInputRequest:
                            d_2_existing_ = source0_.req
                            return Part_PInputRequest(default__.mergeInputReq(d_2_existing_, d_0_req_), AhpSkeleton.Option_None())
                    if True:
                        return d_1_p_

                return lambda2_()

            return lambda1_

        return lambda0_(req)

    @staticmethod
    def upsertInputPart(ps, req):
        if default__.hasOpenReq(ps, (req).id_):
            return ConfluxSeqRoute.default__.SeqUpdateByIdWhere(default__.partKey, default__.isOpenInput, ps, (req).id_, default__.mergeInputPart(req))
        elif True:
            return (ps) + (_dafny.SeqWithoutIsStrInference([Part_PInputRequest(req, AhpSkeleton.Option_None())]))

    @staticmethod
    def answerInputPart(qId, ans):
        def lambda0_(d_0_ans_, d_1_qId_):
            def lambda1_(d_2_p_):
                def lambda2_():
                    source0_ = d_2_p_
                    if True:
                        if source0_.is_PInputRequest:
                            d_3_r_ = source0_.req
                            d_4_resp_ = source0_.response
                            def iife0_(_pat_let233_0):
                                def iife1_(d_5_dt__update__tmp_h0_):
                                    def iife2_(_pat_let234_0):
                                        def iife3_(d_6_dt__update_hanswers_h0_):
                                            return InputReq_InputReq((d_5_dt__update__tmp_h0_).id_, d_6_dt__update_hanswers_h0_, (d_5_dt__update__tmp_h0_).open_)
                                        return iife3_(_pat_let234_0)
                                    return iife2_((((d_3_r_).answers).set(d_1_qId_, (d_0_ans_).value) if (d_0_ans_).is_Some else ((d_3_r_).answers) - (_dafny.Set({d_1_qId_}))))
                                return iife1_(_pat_let233_0)
                            return Part_PInputRequest(iife0_(d_3_r_), d_4_resp_)
                    if True:
                        return d_2_p_

                return lambda2_()

            return lambda1_

        return lambda0_(ans, qId)

    @staticmethod
    def changeAnswerPart(ps, reqId, qId, ans):
        return ConfluxSeqRoute.default__.SeqUpdateByIdWhere(default__.partKey, default__.isOpenInput, ps, reqId, default__.answerInputPart(qId, ans))

    @staticmethod
    def mergeAnswers(drafts, completed):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: _dafny.Seq
            for compr_0_ in (((drafts).keys) | ((completed).keys)).Elements:
                d_0_k_: _dafny.Seq = compr_0_
                if (d_0_k_) in (((drafts).keys) | ((completed).keys)):
                    coll0_[d_0_k_] = ((completed)[d_0_k_] if (d_0_k_) in (completed) else (drafts)[d_0_k_])
            return _dafny.Map(coll0_)
        return iife0_()
        

    @staticmethod
    def completeInputPart(resp, answers):
        def lambda0_(d_0_answers_, d_1_resp_):
            def lambda1_(d_2_p_):
                def lambda2_():
                    source0_ = d_2_p_
                    if True:
                        if source0_.is_PInputRequest:
                            d_3_r_ = source0_.req
                            def iife0_(_pat_let235_0):
                                def iife1_(d_4_dt__update__tmp_h0_):
                                    def iife2_(_pat_let236_0):
                                        def iife3_(d_5_dt__update_hanswers_h0_):
                                            return InputReq_InputReq((d_4_dt__update__tmp_h0_).id_, d_5_dt__update_hanswers_h0_, (d_4_dt__update__tmp_h0_).open_)
                                        return iife3_(_pat_let236_0)
                                    return iife2_(default__.mergeAnswers((d_3_r_).answers, d_0_answers_))
                                return iife1_(_pat_let235_0)
                            return Part_PInputRequest(iife0_(d_3_r_), AhpSkeleton.Option_Some(d_1_resp_))
                    if True:
                        return d_2_p_

                return lambda2_()

            return lambda1_

        return lambda0_(answers, resp)

    @staticmethod
    def completeAnswerPart(ps, reqId, resp, answers):
        return ConfluxSeqRoute.default__.SeqUpdateByIdWhere(default__.partKey, default__.isOpenInput, ps, reqId, default__.completeInputPart(resp, answers))

    @staticmethod
    def turnMatches(s, id_):
        return (((s).activeTurn).is_Some) and (((((s).activeTurn).value).turnId) == (id_))

    @staticmethod
    def optOr(o, d):
        if (o).is_Some:
            return (o).value
        elif True:
            return d

    @staticmethod
    def anyPendingTC(ps):
        if (len(ps)) == (0):
            return False
        elif True:
            return ((((ps)[0]).is_PToolCall) and (((((((ps)[0]).tc).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-confirmation")))) or (((((ps)[0]).tc).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-result-confirmation"))))) or (((((ps)[0]).tc).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auth-required")))))) or (default__.anyPendingTC(_dafny.SeqWithoutIsStrInference((ps)[1::])))

    @staticmethod
    def hasPendingTCState(s):
        return (((s).activeTurn).is_Some) and (default__.anyPendingTC((((s).activeTurn).value).parts))

    @staticmethod
    def anyOpenInput(ps):
        if (len(ps)) == (0):
            return False
        elif True:
            return ((((ps)[0]).is_PInputRequest) and ((((ps)[0]).response).is_None)) or (default__.anyOpenInput(_dafny.SeqWithoutIsStrInference((ps)[1::])))

    @staticmethod
    def hasOpenInputState(s):
        return (((s).activeTurn).is_Some) and (default__.anyOpenInput((((s).activeTurn).value).parts))

    @staticmethod
    def activityBits(s, isError):
        if isError:
            return default__.ERR
        elif (default__.hasOpenInputState(s)) or (default__.hasPendingTCState(s)):
            return default__.INP
        elif ((s).activeTurn).is_Some:
            return default__.GEN
        elif True:
            return default__.IDLE

    @staticmethod
    def summaryStatus(s, isError):
        return (((s).status) & ((~(default__.ACT)) & ((1 << 32) - 1))) | (default__.activityBits(s, isError))

    @staticmethod
    def appendPart(t, p):
        d_0_dt__update__tmp_h0_ = t
        d_1_dt__update_hparts_h0_ = ((t).parts) + (_dafny.SeqWithoutIsStrInference([p]))
        return Turn_Turn((d_0_dt__update__tmp_h0_).turnId, (d_0_dt__update__tmp_h0_).message, d_1_dt__update_hparts_h0_, (d_0_dt__update__tmp_h0_).state, (d_0_dt__update__tmp_h0_).usage, (d_0_dt__update__tmp_h0_).error)

    @staticmethod
    def partKey(p):
        source0_ = p
        if True:
            if source0_.is_PMarkdown:
                d_0_pid_ = source0_.id_
                return d_0_pid_
        if True:
            if source0_.is_PReasoning:
                d_1_pid_ = source0_.id_
                return d_1_pid_
        if True:
            if source0_.is_PToolCall:
                d_2_tc_ = source0_.tc
                return (d_2_tc_).toolCallId
        if True:
            d_3_req_ = source0_.req
            return (d_3_req_).id_

    @staticmethod
    def deltaMarkdown(ps, id_, content):
        def lambda0_(d_0_p_):
            return (d_0_p_).is_PMarkdown

        def lambda1_(d_1_content_):
            def lambda2_(d_2_p_):
                def lambda3_():
                    source0_ = d_2_p_
                    if True:
                        if source0_.is_PMarkdown:
                            d_3_pid_ = source0_.id_
                            d_4_c_ = source0_.content
                            return Part_PMarkdown(d_3_pid_, (d_4_c_) + (d_1_content_))
                    if True:
                        return d_2_p_

                return lambda3_()

            return lambda2_

        return ConfluxSeqRoute.default__.SeqUpdateAllWhere(default__.partKey, lambda0_, ps, id_, lambda1_(content))

    @staticmethod
    def deltaReasoning(ps, id_, content):
        def lambda0_(d_0_p_):
            return (d_0_p_).is_PReasoning

        def lambda1_(d_1_content_):
            def lambda2_(d_2_p_):
                def lambda3_():
                    source0_ = d_2_p_
                    if True:
                        if source0_.is_PReasoning:
                            d_3_pid_ = source0_.id_
                            d_4_c_ = source0_.content
                            return Part_PReasoning(d_3_pid_, (d_4_c_) + (d_1_content_))
                    if True:
                        return d_2_p_

                return lambda3_()

            return lambda2_

        return ConfluxSeqRoute.default__.SeqUpdateAllWhere(default__.partKey, lambda0_, ps, id_, lambda1_(content))

    @staticmethod
    def mapTC(ps, id_, f):
        def lambda0_(d_0_p_):
            return (d_0_p_).is_PToolCall

        def lambda1_(d_1_f_):
            def lambda2_(d_2_p_):
                def lambda3_():
                    source0_ = d_2_p_
                    if True:
                        if source0_.is_PToolCall:
                            d_3_tc_ = source0_.tc
                            return Part_PToolCall(d_1_f_(d_3_tc_))
                    if True:
                        return d_2_p_

                return lambda3_()

            return lambda2_

        return ConfluxSeqRoute.default__.SeqUpdateAllWhere(default__.partKey, lambda0_, ps, id_, lambda1_(f))

    @staticmethod
    def metaOr(meta, prior):
        if (meta).is_Some:
            return meta
        elif True:
            return prior

    @staticmethod
    def optionId(j):
        if (((j).is_JObj) and ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "id"))) in ((j).fields))) and ((((j).fields)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "id"))]).is_JStr):
            return (((j).fields)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "id"))]).s
        elif True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))

    @staticmethod
    def selectOption(options, id_):
        while True:
            with _dafny.label():
                if (((options).is_None) or ((id_).is_None)) or ((len((options).value)) == (0)):
                    return AhpSkeleton.Option_None()
                elif (default__.optionId(((options).value)[0])) == ((id_).value):
                    return default__.jNoNull(((options).value)[0])
                elif True:
                    in0_ = AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(((options).value)[1::]))
                    in1_ = id_
                    options = in0_
                    id_ = in1_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def deltaOne(tc, content, inv, meta):
        if ((tc).status) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "streaming"))):
            return tc
        elif True:
            d_0_dt__update__tmp_h0_ = tc
            d_1_dt__update_hinvocationMessage_h0_ = default__.optOr(inv, (tc).invocationMessage)
            d_2_dt__update_hpartialInput_h0_ = AhpSkeleton.Option_Some((default__.optOr((tc).partialInput, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))) + (content))
            d_3_dt__update_hmeta_h0_ = default__.metaOr(meta, (tc).meta)
            return ToolCall_ToolCall((d_0_dt__update__tmp_h0_).toolCallId, (d_0_dt__update__tmp_h0_).toolName, (d_0_dt__update__tmp_h0_).displayName, (d_0_dt__update__tmp_h0_).status, (d_0_dt__update__tmp_h0_).intention, (d_0_dt__update__tmp_h0_).contributor, d_3_dt__update_hmeta_h0_, d_1_dt__update_hinvocationMessage_h0_, (d_0_dt__update__tmp_h0_).toolInput, (d_0_dt__update__tmp_h0_).confirmationTitle, (d_0_dt__update__tmp_h0_).riskAssessment, (d_0_dt__update__tmp_h0_).edits, (d_0_dt__update__tmp_h0_).editable, (d_0_dt__update__tmp_h0_).options, (d_0_dt__update__tmp_h0_).confirmed, (d_0_dt__update__tmp_h0_).success, (d_0_dt__update__tmp_h0_).pastTenseMessage, (d_0_dt__update__tmp_h0_).reason, (d_0_dt__update__tmp_h0_).reasonMessage, (d_0_dt__update__tmp_h0_).userSuggestion, (d_0_dt__update__tmp_h0_).selectedOption, (d_0_dt__update__tmp_h0_).content, (d_0_dt__update__tmp_h0_).structuredContent, (d_0_dt__update__tmp_h0_).error, (d_0_dt__update__tmp_h0_).auth, d_2_dt__update_hpartialInput_h0_)

    @staticmethod
    def readyOne(tc, inv, input_, title, risk, edits, editable, options, confirmed, meta):
        if ((((tc).status) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "streaming")))) and (((tc).status) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))))) and (((tc).status) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-confirmation")))):
            return tc
        elif (confirmed).is_Some:
            d_0_dt__update__tmp_h0_ = tc
            d_1_dt__update_hpartialInput_h0_ = AhpSkeleton.Option_None()
            d_2_dt__update_hauth_h0_ = AhpSkeleton.Option_None()
            d_3_dt__update_herror_h0_ = AhpSkeleton.Option_None()
            d_4_dt__update_hstructuredContent_h0_ = AhpSkeleton.Option_None()
            d_5_dt__update_hcontent_h0_ = AhpSkeleton.Option_None()
            d_6_dt__update_hselectedOption_h0_ = AhpSkeleton.Option_None()
            d_7_dt__update_huserSuggestion_h0_ = AhpSkeleton.Option_None()
            d_8_dt__update_hreasonMessage_h0_ = AhpSkeleton.Option_None()
            d_9_dt__update_hreason_h0_ = AhpSkeleton.Option_None()
            d_10_dt__update_hpastTenseMessage_h0_ = AhpSkeleton.Option_None()
            d_11_dt__update_hsuccess_h0_ = AhpSkeleton.Option_None()
            d_12_dt__update_hconfirmed_h0_ = confirmed
            d_13_dt__update_hoptions_h0_ = AhpSkeleton.Option_None()
            d_14_dt__update_heditable_h0_ = AhpSkeleton.Option_None()
            d_15_dt__update_hedits_h0_ = AhpSkeleton.Option_None()
            d_16_dt__update_hriskAssessment_h0_ = AhpSkeleton.Option_None()
            d_17_dt__update_hconfirmationTitle_h0_ = AhpSkeleton.Option_None()
            d_18_dt__update_htoolInput_h0_ = input_
            d_19_dt__update_hinvocationMessage_h0_ = default__.optOr(inv, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
            d_20_dt__update_hmeta_h0_ = default__.metaOr(meta, (tc).meta)
            d_21_dt__update_hstatus_h0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))
            return ToolCall_ToolCall((d_0_dt__update__tmp_h0_).toolCallId, (d_0_dt__update__tmp_h0_).toolName, (d_0_dt__update__tmp_h0_).displayName, d_21_dt__update_hstatus_h0_, (d_0_dt__update__tmp_h0_).intention, (d_0_dt__update__tmp_h0_).contributor, d_20_dt__update_hmeta_h0_, d_19_dt__update_hinvocationMessage_h0_, d_18_dt__update_htoolInput_h0_, d_17_dt__update_hconfirmationTitle_h0_, d_16_dt__update_hriskAssessment_h0_, d_15_dt__update_hedits_h0_, d_14_dt__update_heditable_h0_, d_13_dt__update_hoptions_h0_, d_12_dt__update_hconfirmed_h0_, d_11_dt__update_hsuccess_h0_, d_10_dt__update_hpastTenseMessage_h0_, d_9_dt__update_hreason_h0_, d_8_dt__update_hreasonMessage_h0_, d_7_dt__update_huserSuggestion_h0_, d_6_dt__update_hselectedOption_h0_, d_5_dt__update_hcontent_h0_, d_4_dt__update_hstructuredContent_h0_, d_3_dt__update_herror_h0_, d_2_dt__update_hauth_h0_, d_1_dt__update_hpartialInput_h0_)
        elif True:
            d_22_pending_ = ((tc).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-confirmation")))
            d_23_dt__update__tmp_h1_ = tc
            d_24_dt__update_hpartialInput_h1_ = AhpSkeleton.Option_None()
            d_25_dt__update_hauth_h1_ = AhpSkeleton.Option_None()
            d_26_dt__update_herror_h1_ = AhpSkeleton.Option_None()
            d_27_dt__update_hstructuredContent_h1_ = AhpSkeleton.Option_None()
            d_28_dt__update_hcontent_h1_ = AhpSkeleton.Option_None()
            d_29_dt__update_hselectedOption_h1_ = AhpSkeleton.Option_None()
            d_30_dt__update_huserSuggestion_h1_ = AhpSkeleton.Option_None()
            d_31_dt__update_hreasonMessage_h1_ = AhpSkeleton.Option_None()
            d_32_dt__update_hreason_h1_ = AhpSkeleton.Option_None()
            d_33_dt__update_hpastTenseMessage_h1_ = AhpSkeleton.Option_None()
            d_34_dt__update_hsuccess_h1_ = AhpSkeleton.Option_None()
            d_35_dt__update_hconfirmed_h1_ = AhpSkeleton.Option_None()
            d_36_dt__update_hoptions_h1_ = (options if ((options).is_Some) or (not(d_22_pending_)) else (tc).options)
            d_37_dt__update_heditable_h1_ = (editable if ((editable).is_Some) or (not(d_22_pending_)) else (tc).editable)
            d_38_dt__update_hedits_h1_ = (edits if ((edits).is_Some) or (not(d_22_pending_)) else (tc).edits)
            d_39_dt__update_hriskAssessment_h1_ = (risk if ((risk).is_Some) or (not(d_22_pending_)) else (tc).riskAssessment)
            d_40_dt__update_hconfirmationTitle_h1_ = (title if ((title).is_Some) or (not(d_22_pending_)) else (tc).confirmationTitle)
            d_41_dt__update_htoolInput_h1_ = (input_ if ((input_).is_Some) or (not(d_22_pending_)) else (tc).toolInput)
            d_42_dt__update_hinvocationMessage_h1_ = default__.optOr(inv, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
            d_43_dt__update_hmeta_h1_ = default__.metaOr(meta, (tc).meta)
            d_44_dt__update_hstatus_h1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-confirmation"))
            return ToolCall_ToolCall((d_23_dt__update__tmp_h1_).toolCallId, (d_23_dt__update__tmp_h1_).toolName, (d_23_dt__update__tmp_h1_).displayName, d_44_dt__update_hstatus_h1_, (d_23_dt__update__tmp_h1_).intention, (d_23_dt__update__tmp_h1_).contributor, d_43_dt__update_hmeta_h1_, d_42_dt__update_hinvocationMessage_h1_, d_41_dt__update_htoolInput_h1_, d_40_dt__update_hconfirmationTitle_h1_, d_39_dt__update_hriskAssessment_h1_, d_38_dt__update_hedits_h1_, d_37_dt__update_heditable_h1_, d_36_dt__update_hoptions_h1_, d_35_dt__update_hconfirmed_h1_, d_34_dt__update_hsuccess_h1_, d_33_dt__update_hpastTenseMessage_h1_, d_32_dt__update_hreason_h1_, d_31_dt__update_hreasonMessage_h1_, d_30_dt__update_huserSuggestion_h1_, d_29_dt__update_hselectedOption_h1_, d_28_dt__update_hcontent_h1_, d_27_dt__update_hstructuredContent_h1_, d_26_dt__update_herror_h1_, d_25_dt__update_hauth_h1_, d_24_dt__update_hpartialInput_h1_)

    @staticmethod
    def authRequiredOne(tc, challenge, meta):
        if ((((tc).status) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running")))) or (((tc).contributor).is_None)) or (((((tc).contributor).value).kind) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcp")))):
            return tc
        elif True:
            d_0_dt__update__tmp_h0_ = tc
            d_1_dt__update_hpartialInput_h0_ = AhpSkeleton.Option_None()
            d_2_dt__update_hauth_h0_ = default__.jNoNull(challenge)
            d_3_dt__update_herror_h0_ = AhpSkeleton.Option_None()
            d_4_dt__update_hstructuredContent_h0_ = AhpSkeleton.Option_None()
            d_5_dt__update_huserSuggestion_h0_ = AhpSkeleton.Option_None()
            d_6_dt__update_hreasonMessage_h0_ = AhpSkeleton.Option_None()
            d_7_dt__update_hreason_h0_ = AhpSkeleton.Option_None()
            d_8_dt__update_hpastTenseMessage_h0_ = AhpSkeleton.Option_None()
            d_9_dt__update_hsuccess_h0_ = AhpSkeleton.Option_None()
            d_10_dt__update_hoptions_h0_ = AhpSkeleton.Option_None()
            d_11_dt__update_heditable_h0_ = AhpSkeleton.Option_None()
            d_12_dt__update_hedits_h0_ = AhpSkeleton.Option_None()
            d_13_dt__update_hriskAssessment_h0_ = AhpSkeleton.Option_None()
            d_14_dt__update_hconfirmationTitle_h0_ = AhpSkeleton.Option_None()
            d_15_dt__update_hmeta_h0_ = default__.metaOr(meta, (tc).meta)
            d_16_dt__update_hstatus_h0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auth-required"))
            return ToolCall_ToolCall((d_0_dt__update__tmp_h0_).toolCallId, (d_0_dt__update__tmp_h0_).toolName, (d_0_dt__update__tmp_h0_).displayName, d_16_dt__update_hstatus_h0_, (d_0_dt__update__tmp_h0_).intention, (d_0_dt__update__tmp_h0_).contributor, d_15_dt__update_hmeta_h0_, (d_0_dt__update__tmp_h0_).invocationMessage, (d_0_dt__update__tmp_h0_).toolInput, d_14_dt__update_hconfirmationTitle_h0_, d_13_dt__update_hriskAssessment_h0_, d_12_dt__update_hedits_h0_, d_11_dt__update_heditable_h0_, d_10_dt__update_hoptions_h0_, (d_0_dt__update__tmp_h0_).confirmed, d_9_dt__update_hsuccess_h0_, d_8_dt__update_hpastTenseMessage_h0_, d_7_dt__update_hreason_h0_, d_6_dt__update_hreasonMessage_h0_, d_5_dt__update_huserSuggestion_h0_, (d_0_dt__update__tmp_h0_).selectedOption, (d_0_dt__update__tmp_h0_).content, d_4_dt__update_hstructuredContent_h0_, d_3_dt__update_herror_h0_, d_2_dt__update_hauth_h0_, d_1_dt__update_hpartialInput_h0_)

    @staticmethod
    def authResolvedOne(tc, meta):
        if ((tc).status) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auth-required"))):
            return tc
        elif True:
            d_0_dt__update__tmp_h0_ = tc
            d_1_dt__update_hauth_h0_ = AhpSkeleton.Option_None()
            d_2_dt__update_hmeta_h0_ = default__.metaOr(meta, (tc).meta)
            d_3_dt__update_hstatus_h0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))
            return ToolCall_ToolCall((d_0_dt__update__tmp_h0_).toolCallId, (d_0_dt__update__tmp_h0_).toolName, (d_0_dt__update__tmp_h0_).displayName, d_3_dt__update_hstatus_h0_, (d_0_dt__update__tmp_h0_).intention, (d_0_dt__update__tmp_h0_).contributor, d_2_dt__update_hmeta_h0_, (d_0_dt__update__tmp_h0_).invocationMessage, (d_0_dt__update__tmp_h0_).toolInput, (d_0_dt__update__tmp_h0_).confirmationTitle, (d_0_dt__update__tmp_h0_).riskAssessment, (d_0_dt__update__tmp_h0_).edits, (d_0_dt__update__tmp_h0_).editable, (d_0_dt__update__tmp_h0_).options, (d_0_dt__update__tmp_h0_).confirmed, (d_0_dt__update__tmp_h0_).success, (d_0_dt__update__tmp_h0_).pastTenseMessage, (d_0_dt__update__tmp_h0_).reason, (d_0_dt__update__tmp_h0_).reasonMessage, (d_0_dt__update__tmp_h0_).userSuggestion, (d_0_dt__update__tmp_h0_).selectedOption, (d_0_dt__update__tmp_h0_).content, (d_0_dt__update__tmp_h0_).structuredContent, (d_0_dt__update__tmp_h0_).error, d_1_dt__update_hauth_h0_, (d_0_dt__update__tmp_h0_).partialInput)

    @staticmethod
    def confirmOne(tc, approved, confirmedVal, reason, reasonMessage, userSuggestion, editedInput, selectedOptionId, meta):
        if ((tc).status) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-confirmation"))):
            return tc
        elif approved:
            d_0_dt__update__tmp_h0_ = tc
            d_1_dt__update_hauth_h0_ = AhpSkeleton.Option_None()
            d_2_dt__update_herror_h0_ = AhpSkeleton.Option_None()
            d_3_dt__update_hstructuredContent_h0_ = AhpSkeleton.Option_None()
            d_4_dt__update_hcontent_h0_ = AhpSkeleton.Option_None()
            d_5_dt__update_huserSuggestion_h0_ = AhpSkeleton.Option_None()
            d_6_dt__update_hreasonMessage_h0_ = AhpSkeleton.Option_None()
            d_7_dt__update_hreason_h0_ = AhpSkeleton.Option_None()
            d_8_dt__update_hpastTenseMessage_h0_ = AhpSkeleton.Option_None()
            d_9_dt__update_hsuccess_h0_ = AhpSkeleton.Option_None()
            d_10_dt__update_hpartialInput_h0_ = AhpSkeleton.Option_None()
            d_11_dt__update_hselectedOption_h0_ = default__.selectOption((tc).options, selectedOptionId)
            d_12_dt__update_hconfirmed_h0_ = confirmedVal
            d_13_dt__update_hoptions_h0_ = AhpSkeleton.Option_None()
            d_14_dt__update_heditable_h0_ = AhpSkeleton.Option_None()
            d_15_dt__update_hedits_h0_ = AhpSkeleton.Option_None()
            d_16_dt__update_hriskAssessment_h0_ = AhpSkeleton.Option_None()
            d_17_dt__update_hconfirmationTitle_h0_ = AhpSkeleton.Option_None()
            d_18_dt__update_htoolInput_h0_ = (editedInput if (editedInput).is_Some else (tc).toolInput)
            d_19_dt__update_hmeta_h0_ = default__.metaOr(meta, (tc).meta)
            d_20_dt__update_hstatus_h0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))
            return ToolCall_ToolCall((d_0_dt__update__tmp_h0_).toolCallId, (d_0_dt__update__tmp_h0_).toolName, (d_0_dt__update__tmp_h0_).displayName, d_20_dt__update_hstatus_h0_, (d_0_dt__update__tmp_h0_).intention, (d_0_dt__update__tmp_h0_).contributor, d_19_dt__update_hmeta_h0_, (d_0_dt__update__tmp_h0_).invocationMessage, d_18_dt__update_htoolInput_h0_, d_17_dt__update_hconfirmationTitle_h0_, d_16_dt__update_hriskAssessment_h0_, d_15_dt__update_hedits_h0_, d_14_dt__update_heditable_h0_, d_13_dt__update_hoptions_h0_, d_12_dt__update_hconfirmed_h0_, d_9_dt__update_hsuccess_h0_, d_8_dt__update_hpastTenseMessage_h0_, d_7_dt__update_hreason_h0_, d_6_dt__update_hreasonMessage_h0_, d_5_dt__update_huserSuggestion_h0_, d_11_dt__update_hselectedOption_h0_, d_4_dt__update_hcontent_h0_, d_3_dt__update_hstructuredContent_h0_, d_2_dt__update_herror_h0_, d_1_dt__update_hauth_h0_, d_10_dt__update_hpartialInput_h0_)
        elif True:
            d_21_dt__update__tmp_h1_ = tc
            d_22_dt__update_hpartialInput_h1_ = AhpSkeleton.Option_None()
            d_23_dt__update_hauth_h1_ = AhpSkeleton.Option_None()
            d_24_dt__update_herror_h1_ = AhpSkeleton.Option_None()
            d_25_dt__update_hstructuredContent_h1_ = AhpSkeleton.Option_None()
            d_26_dt__update_hcontent_h1_ = AhpSkeleton.Option_None()
            d_27_dt__update_hselectedOption_h1_ = default__.selectOption((tc).options, selectedOptionId)
            d_28_dt__update_huserSuggestion_h1_ = userSuggestion
            d_29_dt__update_hreasonMessage_h1_ = reasonMessage
            d_30_dt__update_hreason_h1_ = reason
            d_31_dt__update_hpastTenseMessage_h1_ = AhpSkeleton.Option_None()
            d_32_dt__update_hsuccess_h1_ = AhpSkeleton.Option_None()
            d_33_dt__update_hconfirmed_h1_ = AhpSkeleton.Option_None()
            d_34_dt__update_hoptions_h1_ = AhpSkeleton.Option_None()
            d_35_dt__update_heditable_h1_ = AhpSkeleton.Option_None()
            d_36_dt__update_hedits_h1_ = AhpSkeleton.Option_None()
            d_37_dt__update_hriskAssessment_h1_ = AhpSkeleton.Option_None()
            d_38_dt__update_hconfirmationTitle_h1_ = AhpSkeleton.Option_None()
            d_39_dt__update_hmeta_h1_ = default__.metaOr(meta, (tc).meta)
            d_40_dt__update_hstatus_h1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cancelled"))
            return ToolCall_ToolCall((d_21_dt__update__tmp_h1_).toolCallId, (d_21_dt__update__tmp_h1_).toolName, (d_21_dt__update__tmp_h1_).displayName, d_40_dt__update_hstatus_h1_, (d_21_dt__update__tmp_h1_).intention, (d_21_dt__update__tmp_h1_).contributor, d_39_dt__update_hmeta_h1_, (d_21_dt__update__tmp_h1_).invocationMessage, (d_21_dt__update__tmp_h1_).toolInput, d_38_dt__update_hconfirmationTitle_h1_, d_37_dt__update_hriskAssessment_h1_, d_36_dt__update_hedits_h1_, d_35_dt__update_heditable_h1_, d_34_dt__update_hoptions_h1_, d_33_dt__update_hconfirmed_h1_, d_32_dt__update_hsuccess_h1_, d_31_dt__update_hpastTenseMessage_h1_, d_30_dt__update_hreason_h1_, d_29_dt__update_hreasonMessage_h1_, d_28_dt__update_huserSuggestion_h1_, d_27_dt__update_hselectedOption_h1_, d_26_dt__update_hcontent_h1_, d_25_dt__update_hstructuredContent_h1_, d_24_dt__update_herror_h1_, d_23_dt__update_hauth_h1_, d_22_dt__update_hpartialInput_h1_)

    @staticmethod
    def completeOne(tc, ok, past, resultContent, structured, err, requiresResultConfirmation, meta):
        if ((tc).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auth-required"))):
            if ok:
                return tc
            elif True:
                d_0_dt__update__tmp_h0_ = tc
                d_1_dt__update_hpartialInput_h0_ = AhpSkeleton.Option_None()
                d_2_dt__update_hauth_h0_ = AhpSkeleton.Option_None()
                d_3_dt__update_herror_h0_ = err
                d_4_dt__update_hstructuredContent_h0_ = structured
                d_5_dt__update_hcontent_h0_ = (resultContent if (resultContent).is_Some else (tc).content)
                d_6_dt__update_huserSuggestion_h0_ = AhpSkeleton.Option_None()
                d_7_dt__update_hreasonMessage_h0_ = AhpSkeleton.Option_None()
                d_8_dt__update_hreason_h0_ = AhpSkeleton.Option_None()
                d_9_dt__update_hpastTenseMessage_h0_ = past
                d_10_dt__update_hsuccess_h0_ = AhpSkeleton.Option_Some(False)
                d_11_dt__update_hoptions_h0_ = AhpSkeleton.Option_None()
                d_12_dt__update_heditable_h0_ = AhpSkeleton.Option_None()
                d_13_dt__update_hedits_h0_ = AhpSkeleton.Option_None()
                d_14_dt__update_hriskAssessment_h0_ = AhpSkeleton.Option_None()
                d_15_dt__update_hconfirmationTitle_h0_ = AhpSkeleton.Option_None()
                d_16_dt__update_hmeta_h0_ = default__.metaOr(meta, (tc).meta)
                d_17_dt__update_hstatus_h0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "completed"))
                return ToolCall_ToolCall((d_0_dt__update__tmp_h0_).toolCallId, (d_0_dt__update__tmp_h0_).toolName, (d_0_dt__update__tmp_h0_).displayName, d_17_dt__update_hstatus_h0_, (d_0_dt__update__tmp_h0_).intention, (d_0_dt__update__tmp_h0_).contributor, d_16_dt__update_hmeta_h0_, (d_0_dt__update__tmp_h0_).invocationMessage, (d_0_dt__update__tmp_h0_).toolInput, d_15_dt__update_hconfirmationTitle_h0_, d_14_dt__update_hriskAssessment_h0_, d_13_dt__update_hedits_h0_, d_12_dt__update_heditable_h0_, d_11_dt__update_hoptions_h0_, (d_0_dt__update__tmp_h0_).confirmed, d_10_dt__update_hsuccess_h0_, d_9_dt__update_hpastTenseMessage_h0_, d_8_dt__update_hreason_h0_, d_7_dt__update_hreasonMessage_h0_, d_6_dt__update_huserSuggestion_h0_, (d_0_dt__update__tmp_h0_).selectedOption, d_5_dt__update_hcontent_h0_, d_4_dt__update_hstructuredContent_h0_, d_3_dt__update_herror_h0_, d_2_dt__update_hauth_h0_, d_1_dt__update_hpartialInput_h0_)
        elif (((tc).status) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running")))) and (((tc).status) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-confirmation")))):
            return tc
        elif True:
            d_18_conf_ = ((tc).confirmed if ((tc).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))) else AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "not-needed"))))
            d_19_selected_ = ((tc).selectedOption if ((tc).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))) else AhpSkeleton.Option_None())
            d_20_newStatus_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-result-confirmation")) if requiresResultConfirmation else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "completed")))
            d_21_dt__update__tmp_h1_ = tc
            d_22_dt__update_hpartialInput_h1_ = AhpSkeleton.Option_None()
            d_23_dt__update_hauth_h1_ = AhpSkeleton.Option_None()
            d_24_dt__update_herror_h1_ = err
            d_25_dt__update_hstructuredContent_h1_ = structured
            d_26_dt__update_hcontent_h1_ = resultContent
            d_27_dt__update_huserSuggestion_h1_ = AhpSkeleton.Option_None()
            d_28_dt__update_hreasonMessage_h1_ = AhpSkeleton.Option_None()
            d_29_dt__update_hreason_h1_ = AhpSkeleton.Option_None()
            d_30_dt__update_hpastTenseMessage_h1_ = past
            d_31_dt__update_hsuccess_h1_ = AhpSkeleton.Option_Some(ok)
            d_32_dt__update_hselectedOption_h0_ = d_19_selected_
            d_33_dt__update_hconfirmed_h0_ = d_18_conf_
            d_34_dt__update_hoptions_h1_ = AhpSkeleton.Option_None()
            d_35_dt__update_heditable_h1_ = AhpSkeleton.Option_None()
            d_36_dt__update_hedits_h1_ = AhpSkeleton.Option_None()
            d_37_dt__update_hriskAssessment_h1_ = AhpSkeleton.Option_None()
            d_38_dt__update_hconfirmationTitle_h1_ = AhpSkeleton.Option_None()
            d_39_dt__update_hmeta_h1_ = default__.metaOr(meta, (tc).meta)
            d_40_dt__update_hstatus_h1_ = d_20_newStatus_
            return ToolCall_ToolCall((d_21_dt__update__tmp_h1_).toolCallId, (d_21_dt__update__tmp_h1_).toolName, (d_21_dt__update__tmp_h1_).displayName, d_40_dt__update_hstatus_h1_, (d_21_dt__update__tmp_h1_).intention, (d_21_dt__update__tmp_h1_).contributor, d_39_dt__update_hmeta_h1_, (d_21_dt__update__tmp_h1_).invocationMessage, (d_21_dt__update__tmp_h1_).toolInput, d_38_dt__update_hconfirmationTitle_h1_, d_37_dt__update_hriskAssessment_h1_, d_36_dt__update_hedits_h1_, d_35_dt__update_heditable_h1_, d_34_dt__update_hoptions_h1_, d_33_dt__update_hconfirmed_h0_, d_31_dt__update_hsuccess_h1_, d_30_dt__update_hpastTenseMessage_h1_, d_29_dt__update_hreason_h1_, d_28_dt__update_hreasonMessage_h1_, d_27_dt__update_huserSuggestion_h1_, d_32_dt__update_hselectedOption_h0_, d_26_dt__update_hcontent_h1_, d_25_dt__update_hstructuredContent_h1_, d_24_dt__update_herror_h1_, d_23_dt__update_hauth_h1_, d_22_dt__update_hpartialInput_h1_)

    @staticmethod
    def resultConfirmOne(tc, approved, meta):
        if ((tc).status) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-result-confirmation"))):
            return tc
        elif approved:
            d_0_dt__update__tmp_h0_ = tc
            d_1_dt__update_hmeta_h0_ = default__.metaOr(meta, (tc).meta)
            d_2_dt__update_hstatus_h0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "completed"))
            return ToolCall_ToolCall((d_0_dt__update__tmp_h0_).toolCallId, (d_0_dt__update__tmp_h0_).toolName, (d_0_dt__update__tmp_h0_).displayName, d_2_dt__update_hstatus_h0_, (d_0_dt__update__tmp_h0_).intention, (d_0_dt__update__tmp_h0_).contributor, d_1_dt__update_hmeta_h0_, (d_0_dt__update__tmp_h0_).invocationMessage, (d_0_dt__update__tmp_h0_).toolInput, (d_0_dt__update__tmp_h0_).confirmationTitle, (d_0_dt__update__tmp_h0_).riskAssessment, (d_0_dt__update__tmp_h0_).edits, (d_0_dt__update__tmp_h0_).editable, (d_0_dt__update__tmp_h0_).options, (d_0_dt__update__tmp_h0_).confirmed, (d_0_dt__update__tmp_h0_).success, (d_0_dt__update__tmp_h0_).pastTenseMessage, (d_0_dt__update__tmp_h0_).reason, (d_0_dt__update__tmp_h0_).reasonMessage, (d_0_dt__update__tmp_h0_).userSuggestion, (d_0_dt__update__tmp_h0_).selectedOption, (d_0_dt__update__tmp_h0_).content, (d_0_dt__update__tmp_h0_).structuredContent, (d_0_dt__update__tmp_h0_).error, (d_0_dt__update__tmp_h0_).auth, (d_0_dt__update__tmp_h0_).partialInput)
        elif True:
            d_3_dt__update__tmp_h1_ = tc
            d_4_dt__update_herror_h0_ = AhpSkeleton.Option_None()
            d_5_dt__update_hstructuredContent_h0_ = AhpSkeleton.Option_None()
            d_6_dt__update_hcontent_h0_ = AhpSkeleton.Option_None()
            d_7_dt__update_hpastTenseMessage_h0_ = AhpSkeleton.Option_None()
            d_8_dt__update_hsuccess_h0_ = AhpSkeleton.Option_None()
            d_9_dt__update_hconfirmed_h0_ = AhpSkeleton.Option_None()
            d_10_dt__update_hreason_h0_ = AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "result-denied")))
            d_11_dt__update_hmeta_h1_ = default__.metaOr(meta, (tc).meta)
            d_12_dt__update_hstatus_h1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cancelled"))
            return ToolCall_ToolCall((d_3_dt__update__tmp_h1_).toolCallId, (d_3_dt__update__tmp_h1_).toolName, (d_3_dt__update__tmp_h1_).displayName, d_12_dt__update_hstatus_h1_, (d_3_dt__update__tmp_h1_).intention, (d_3_dt__update__tmp_h1_).contributor, d_11_dt__update_hmeta_h1_, (d_3_dt__update__tmp_h1_).invocationMessage, (d_3_dt__update__tmp_h1_).toolInput, (d_3_dt__update__tmp_h1_).confirmationTitle, (d_3_dt__update__tmp_h1_).riskAssessment, (d_3_dt__update__tmp_h1_).edits, (d_3_dt__update__tmp_h1_).editable, (d_3_dt__update__tmp_h1_).options, d_9_dt__update_hconfirmed_h0_, d_8_dt__update_hsuccess_h0_, d_7_dt__update_hpastTenseMessage_h0_, d_10_dt__update_hreason_h0_, (d_3_dt__update__tmp_h1_).reasonMessage, (d_3_dt__update__tmp_h1_).userSuggestion, (d_3_dt__update__tmp_h1_).selectedOption, d_6_dt__update_hcontent_h0_, d_5_dt__update_hstructuredContent_h0_, d_4_dt__update_herror_h0_, (d_3_dt__update__tmp_h1_).auth, (d_3_dt__update__tmp_h1_).partialInput)

    @staticmethod
    def contentChangedOne(tc, c, meta):
        if ((tc).status) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))):
            return tc
        elif True:
            d_0_dt__update__tmp_h0_ = tc
            d_1_dt__update_hcontent_h0_ = default__.jNoNull(c)
            d_2_dt__update_hmeta_h0_ = default__.metaOr(meta, (tc).meta)
            return ToolCall_ToolCall((d_0_dt__update__tmp_h0_).toolCallId, (d_0_dt__update__tmp_h0_).toolName, (d_0_dt__update__tmp_h0_).displayName, (d_0_dt__update__tmp_h0_).status, (d_0_dt__update__tmp_h0_).intention, (d_0_dt__update__tmp_h0_).contributor, d_2_dt__update_hmeta_h0_, (d_0_dt__update__tmp_h0_).invocationMessage, (d_0_dt__update__tmp_h0_).toolInput, (d_0_dt__update__tmp_h0_).confirmationTitle, (d_0_dt__update__tmp_h0_).riskAssessment, (d_0_dt__update__tmp_h0_).edits, (d_0_dt__update__tmp_h0_).editable, (d_0_dt__update__tmp_h0_).options, (d_0_dt__update__tmp_h0_).confirmed, (d_0_dt__update__tmp_h0_).success, (d_0_dt__update__tmp_h0_).pastTenseMessage, (d_0_dt__update__tmp_h0_).reason, (d_0_dt__update__tmp_h0_).reasonMessage, (d_0_dt__update__tmp_h0_).userSuggestion, (d_0_dt__update__tmp_h0_).selectedOption, d_1_dt__update_hcontent_h0_, (d_0_dt__update__tmp_h0_).structuredContent, (d_0_dt__update__tmp_h0_).error, (d_0_dt__update__tmp_h0_).auth, (d_0_dt__update__tmp_h0_).partialInput)

    @staticmethod
    def forceCancel(ps):
        def lambda0_(d_0_h_):
            def lambda1_():
                source0_ = d_0_h_
                if True:
                    if source0_.is_PToolCall:
                        d_1_tc_ = source0_.tc
                        if (((d_1_tc_).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "completed")))) or (((d_1_tc_).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cancelled")))):
                            return d_0_h_
                        elif True:
                            def iife0_(_pat_let237_0):
                                def iife1_(d_2_dt__update__tmp_h0_):
                                    def iife2_(_pat_let238_0):
                                        def iife3_(d_3_dt__update_hpartialInput_h0_):
                                            def iife4_(_pat_let239_0):
                                                def iife5_(d_4_dt__update_hauth_h0_):
                                                    def iife6_(_pat_let240_0):
                                                        def iife7_(d_5_dt__update_herror_h0_):
                                                            def iife8_(_pat_let241_0):
                                                                def iife9_(d_6_dt__update_hstructuredContent_h0_):
                                                                    def iife10_(_pat_let242_0):
                                                                        def iife11_(d_7_dt__update_hcontent_h0_):
                                                                            def iife12_(_pat_let243_0):
                                                                                def iife13_(d_8_dt__update_hselectedOption_h0_):
                                                                                    def iife14_(_pat_let244_0):
                                                                                        def iife15_(d_9_dt__update_huserSuggestion_h0_):
                                                                                            def iife16_(_pat_let245_0):
                                                                                                def iife17_(d_10_dt__update_hreasonMessage_h0_):
                                                                                                    def iife18_(_pat_let246_0):
                                                                                                        def iife19_(d_11_dt__update_hpastTenseMessage_h0_):
                                                                                                            def iife20_(_pat_let247_0):
                                                                                                                def iife21_(d_12_dt__update_hsuccess_h0_):
                                                                                                                    def iife22_(_pat_let248_0):
                                                                                                                        def iife23_(d_13_dt__update_hconfirmed_h0_):
                                                                                                                            def iife24_(_pat_let249_0):
                                                                                                                                def iife25_(d_14_dt__update_hoptions_h0_):
                                                                                                                                    def iife26_(_pat_let250_0):
                                                                                                                                        def iife27_(d_15_dt__update_heditable_h0_):
                                                                                                                                            def iife28_(_pat_let251_0):
                                                                                                                                                def iife29_(d_16_dt__update_hedits_h0_):
                                                                                                                                                    def iife30_(_pat_let252_0):
                                                                                                                                                        def iife31_(d_17_dt__update_hriskAssessment_h0_):
                                                                                                                                                            def iife32_(_pat_let253_0):
                                                                                                                                                                def iife33_(d_18_dt__update_hconfirmationTitle_h0_):
                                                                                                                                                                    def iife34_(_pat_let254_0):
                                                                                                                                                                        def iife35_(d_19_dt__update_htoolInput_h0_):
                                                                                                                                                                            def iife36_(_pat_let255_0):
                                                                                                                                                                                def iife37_(d_20_dt__update_hreason_h0_):
                                                                                                                                                                                    def iife38_(_pat_let256_0):
                                                                                                                                                                                        def iife39_(d_21_dt__update_hstatus_h0_):
                                                                                                                                                                                            return ToolCall_ToolCall((d_2_dt__update__tmp_h0_).toolCallId, (d_2_dt__update__tmp_h0_).toolName, (d_2_dt__update__tmp_h0_).displayName, d_21_dt__update_hstatus_h0_, (d_2_dt__update__tmp_h0_).intention, (d_2_dt__update__tmp_h0_).contributor, (d_2_dt__update__tmp_h0_).meta, (d_2_dt__update__tmp_h0_).invocationMessage, d_19_dt__update_htoolInput_h0_, d_18_dt__update_hconfirmationTitle_h0_, d_17_dt__update_hriskAssessment_h0_, d_16_dt__update_hedits_h0_, d_15_dt__update_heditable_h0_, d_14_dt__update_hoptions_h0_, d_13_dt__update_hconfirmed_h0_, d_12_dt__update_hsuccess_h0_, d_11_dt__update_hpastTenseMessage_h0_, d_20_dt__update_hreason_h0_, d_10_dt__update_hreasonMessage_h0_, d_9_dt__update_huserSuggestion_h0_, d_8_dt__update_hselectedOption_h0_, d_7_dt__update_hcontent_h0_, d_6_dt__update_hstructuredContent_h0_, d_5_dt__update_herror_h0_, d_4_dt__update_hauth_h0_, d_3_dt__update_hpartialInput_h0_)
                                                                                                                                                                                        return iife39_(_pat_let256_0)
                                                                                                                                                                                    return iife38_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cancelled")))
                                                                                                                                                                                return iife37_(_pat_let255_0)
                                                                                                                                                                            return iife36_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skipped"))))
                                                                                                                                                                        return iife35_(_pat_let254_0)
                                                                                                                                                                    return iife34_((AhpSkeleton.Option_None() if ((d_1_tc_).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "streaming"))) else (d_1_tc_).toolInput))
                                                                                                                                                                return iife33_(_pat_let253_0)
                                                                                                                                                            return iife32_(AhpSkeleton.Option_None())
                                                                                                                                                        return iife31_(_pat_let252_0)
                                                                                                                                                    return iife30_(AhpSkeleton.Option_None())
                                                                                                                                                return iife29_(_pat_let251_0)
                                                                                                                                            return iife28_(AhpSkeleton.Option_None())
                                                                                                                                        return iife27_(_pat_let250_0)
                                                                                                                                    return iife26_(AhpSkeleton.Option_None())
                                                                                                                                return iife25_(_pat_let249_0)
                                                                                                                            return iife24_(AhpSkeleton.Option_None())
                                                                                                                        return iife23_(_pat_let248_0)
                                                                                                                    return iife22_(AhpSkeleton.Option_None())
                                                                                                                return iife21_(_pat_let247_0)
                                                                                                            return iife20_(AhpSkeleton.Option_None())
                                                                                                        return iife19_(_pat_let246_0)
                                                                                                    return iife18_(AhpSkeleton.Option_None())
                                                                                                return iife17_(_pat_let245_0)
                                                                                            return iife16_(AhpSkeleton.Option_None())
                                                                                        return iife15_(_pat_let244_0)
                                                                                    return iife14_(AhpSkeleton.Option_None())
                                                                                return iife13_(_pat_let243_0)
                                                                            return iife12_(AhpSkeleton.Option_None())
                                                                        return iife11_(_pat_let242_0)
                                                                    return iife10_(AhpSkeleton.Option_None())
                                                                return iife9_(_pat_let241_0)
                                                            return iife8_(AhpSkeleton.Option_None())
                                                        return iife7_(_pat_let240_0)
                                                    return iife6_(AhpSkeleton.Option_None())
                                                return iife5_(_pat_let239_0)
                                            return iife4_(AhpSkeleton.Option_None())
                                        return iife3_(_pat_let238_0)
                                    return iife2_(AhpSkeleton.Option_None())
                                return iife1_(_pat_let237_0)
                            return Part_PToolCall(iife0_(d_1_tc_))
                if True:
                    return d_0_h_

            return lambda1_()

        return ConfluxOperators.default__.Map(lambda0_, ps)

    @staticmethod
    def endTurn(s, id_, turnState, isError, err):
        pat_let_tv0_ = err
        pat_let_tv1_ = turnState
        pat_let_tv2_ = s
        pat_let_tv3_ = isError
        if not(default__.turnMatches(s, id_)):
            return (s, AhpSkeleton.ReduceOutcome_NoOp())
        elif True:
            d_0_t_ = ((s).activeTurn).value
            def iife0_(_pat_let257_0):
                def iife1_(d_2_dt__update__tmp_h0_):
                    def iife2_(_pat_let258_0):
                        def iife3_(d_3_dt__update_herror_h0_):
                            def iife4_(_pat_let259_0):
                                def iife5_(d_4_dt__update_hparts_h0_):
                                    def iife6_(_pat_let260_0):
                                        def iife7_(d_5_dt__update_hstate_h0_):
                                            return Turn_Turn((d_2_dt__update__tmp_h0_).turnId, (d_2_dt__update__tmp_h0_).message, d_4_dt__update_hparts_h0_, d_5_dt__update_hstate_h0_, (d_2_dt__update__tmp_h0_).usage, d_3_dt__update_herror_h0_)
                                        return iife7_(_pat_let260_0)
                                    return iife6_(pat_let_tv1_)
                                return iife5_(_pat_let259_0)
                            return iife4_(default__.forceCancel((d_0_t_).parts))
                        return iife3_(_pat_let258_0)
                    return iife2_(pat_let_tv0_)
                return iife1_(_pat_let257_0)
            d_1_finalized_ = iife0_(d_0_t_)
            def iife8_(_pat_let261_0):
                def iife9_(d_7_dt__update__tmp_h1_):
                    def iife10_(_pat_let262_0):
                        def iife11_(d_8_dt__update_hactiveTurn_h0_):
                            def iife12_(_pat_let263_0):
                                def iife13_(d_9_dt__update_hturns_h0_):
                                    return ChatState_ChatState(d_9_dt__update_hturns_h0_, (d_7_dt__update__tmp_h1_).title, (d_7_dt__update__tmp_h1_).status, (d_7_dt__update__tmp_h1_).modifiedAt, (d_7_dt__update__tmp_h1_).draft, (d_7_dt__update__tmp_h1_).activity, d_8_dt__update_hactiveTurn_h0_, (d_7_dt__update__tmp_h1_).steeringMessage, (d_7_dt__update__tmp_h1_).queuedMessages, (d_7_dt__update__tmp_h1_).nextCursor)
                                return iife13_(_pat_let263_0)
                            return iife12_(((pat_let_tv2_).turns) + (_dafny.SeqWithoutIsStrInference([d_1_finalized_])))
                        return iife11_(_pat_let262_0)
                    return iife10_(AhpSkeleton.Option_None())
                return iife9_(_pat_let261_0)
            d_6_next_ = iife8_(s)
            def iife14_(_pat_let264_0):
                def iife15_(d_10_dt__update__tmp_h2_):
                    def iife16_(_pat_let265_0):
                        def iife17_(d_11_dt__update_hstatus_h0_):
                            return ChatState_ChatState((d_10_dt__update__tmp_h2_).turns, (d_10_dt__update__tmp_h2_).title, d_11_dt__update_hstatus_h0_, (d_10_dt__update__tmp_h2_).modifiedAt, (d_10_dt__update__tmp_h2_).draft, (d_10_dt__update__tmp_h2_).activity, (d_10_dt__update__tmp_h2_).activeTurn, (d_10_dt__update__tmp_h2_).steeringMessage, (d_10_dt__update__tmp_h2_).queuedMessages, (d_10_dt__update__tmp_h2_).nextCursor)
                        return iife17_(_pat_let265_0)
                    return iife16_(default__.summaryStatus(d_6_next_, pat_let_tv3_))
                return iife15_(_pat_let264_0)
            return (iife14_(d_6_next_), AhpSkeleton.ReduceOutcome_Applied())

    @staticmethod
    def ApplyToChat(s, a, now):
        pat_let_tv0_ = s
        pat_let_tv1_ = s
        pat_let_tv2_ = s
        pat_let_tv3_ = s
        pat_let_tv4_ = s
        pat_let_tv5_ = s
        pat_let_tv6_ = s
        pat_let_tv7_ = s
        pat_let_tv8_ = s
        pat_let_tv9_ = s
        pat_let_tv10_ = s
        pat_let_tv11_ = s
        pat_let_tv12_ = s
        pat_let_tv13_ = s
        pat_let_tv14_ = s
        pat_let_tv15_ = s
        pat_let_tv16_ = s
        pat_let_tv17_ = s
        pat_let_tv18_ = s
        pat_let_tv19_ = s
        pat_let_tv20_ = s
        pat_let_tv21_ = s
        pat_let_tv22_ = s
        pat_let_tv23_ = s
        pat_let_tv24_ = s
        pat_let_tv25_ = s
        pat_let_tv26_ = s
        pat_let_tv27_ = s
        pat_let_tv28_ = s
        source0_ = a
        if True:
            if source0_.is_CDraftChanged:
                d_0_d_ = source0_.draft
                def iife0_(_pat_let266_0):
                    def iife1_(d_1_dt__update__tmp_h0_):
                        def iife2_(_pat_let267_0):
                            def iife3_(d_2_dt__update_hdraft_h0_):
                                return ChatState_ChatState((d_1_dt__update__tmp_h0_).turns, (d_1_dt__update__tmp_h0_).title, (d_1_dt__update__tmp_h0_).status, (d_1_dt__update__tmp_h0_).modifiedAt, d_2_dt__update_hdraft_h0_, (d_1_dt__update__tmp_h0_).activity, (d_1_dt__update__tmp_h0_).activeTurn, (d_1_dt__update__tmp_h0_).steeringMessage, (d_1_dt__update__tmp_h0_).queuedMessages, (d_1_dt__update__tmp_h0_).nextCursor)
                            return iife3_(_pat_let267_0)
                        return iife2_(d_0_d_)
                    return iife1_(_pat_let266_0)
                return (iife0_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CActivityChanged:
                d_3_ac_ = source0_.activity
                def iife4_(_pat_let268_0):
                    def iife5_(d_4_dt__update__tmp_h1_):
                        def iife6_(_pat_let269_0):
                            def iife7_(d_5_dt__update_hactivity_h0_):
                                return ChatState_ChatState((d_4_dt__update__tmp_h1_).turns, (d_4_dt__update__tmp_h1_).title, (d_4_dt__update__tmp_h1_).status, (d_4_dt__update__tmp_h1_).modifiedAt, (d_4_dt__update__tmp_h1_).draft, d_5_dt__update_hactivity_h0_, (d_4_dt__update__tmp_h1_).activeTurn, (d_4_dt__update__tmp_h1_).steeringMessage, (d_4_dt__update__tmp_h1_).queuedMessages, (d_4_dt__update__tmp_h1_).nextCursor)
                            return iife7_(_pat_let269_0)
                        return iife6_(d_3_ac_)
                    return iife5_(_pat_let268_0)
                return (iife4_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CTurnStarted:
                d_6_id_ = source0_.turnId
                d_7_m_ = source0_.message
                d_8_qid_ = source0_.queuedMessageId
                def iife8_(_pat_let270_0):
                    def iife9_(d_10_dt__update__tmp_h2_):
                        def iife10_(_pat_let271_0):
                            def iife11_(d_11_dt__update_hactiveTurn_h0_):
                                return ChatState_ChatState((d_10_dt__update__tmp_h2_).turns, (d_10_dt__update__tmp_h2_).title, (d_10_dt__update__tmp_h2_).status, (d_10_dt__update__tmp_h2_).modifiedAt, (d_10_dt__update__tmp_h2_).draft, (d_10_dt__update__tmp_h2_).activity, d_11_dt__update_hactiveTurn_h0_, (d_10_dt__update__tmp_h2_).steeringMessage, (d_10_dt__update__tmp_h2_).queuedMessages, (d_10_dt__update__tmp_h2_).nextCursor)
                            return iife11_(_pat_let271_0)
                        return iife10_(AhpSkeleton.Option_Some(Turn_Turn(d_6_id_, d_7_m_, _dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "in-progress")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())))
                    return iife9_(_pat_let270_0)
                d_9_withTurn_ = iife8_(s)
                def iife12_(_pat_let272_0):
                    def iife13_(d_13_dt__update__tmp_h3_):
                        def iife14_(_pat_let273_0):
                            def iife15_(d_14_dt__update_hstatus_h0_):
                                return ChatState_ChatState((d_13_dt__update__tmp_h3_).turns, (d_13_dt__update__tmp_h3_).title, d_14_dt__update_hstatus_h0_, (d_13_dt__update__tmp_h3_).modifiedAt, (d_13_dt__update__tmp_h3_).draft, (d_13_dt__update__tmp_h3_).activity, (d_13_dt__update__tmp_h3_).activeTurn, (d_13_dt__update__tmp_h3_).steeringMessage, (d_13_dt__update__tmp_h3_).queuedMessages, (d_13_dt__update__tmp_h3_).nextCursor)
                            return iife15_(_pat_let273_0)
                        return iife14_(default__.clearBit(default__.summaryStatus(d_9_withTurn_, False), default__.READ))
                    return iife13_(_pat_let272_0)
                d_12_s2_ = iife12_(d_9_withTurn_)
                def iife16_(_pat_let274_0):
                    def iife17_(d_16_dt__update__tmp_h4_):
                        def iife18_(_pat_let275_0):
                            def iife19_(d_17_dt__update_hsteeringMessage_h0_):
                                return ChatState_ChatState((d_16_dt__update__tmp_h4_).turns, (d_16_dt__update__tmp_h4_).title, (d_16_dt__update__tmp_h4_).status, (d_16_dt__update__tmp_h4_).modifiedAt, (d_16_dt__update__tmp_h4_).draft, (d_16_dt__update__tmp_h4_).activity, (d_16_dt__update__tmp_h4_).activeTurn, d_17_dt__update_hsteeringMessage_h0_, (d_16_dt__update__tmp_h4_).queuedMessages, (d_16_dt__update__tmp_h4_).nextCursor)
                            return iife19_(_pat_let275_0)
                        return iife18_(AhpSkeleton.Option_None())
                    return iife17_(_pat_let274_0)
                d_15_s3_ = (iife16_(d_12_s2_) if (((d_8_qid_).is_Some) and (((d_12_s2_).steeringMessage).is_Some)) and (((((d_12_s2_).steeringMessage).value).id_) == ((d_8_qid_).value)) else d_12_s2_)
                def iife20_(_pat_let276_0):
                    def iife21_(d_19_dt__update__tmp_h5_):
                        def iife22_(_pat_let277_0):
                            def iife23_(d_20_dt__update_hqueuedMessages_h0_):
                                return ChatState_ChatState((d_19_dt__update__tmp_h5_).turns, (d_19_dt__update__tmp_h5_).title, (d_19_dt__update__tmp_h5_).status, (d_19_dt__update__tmp_h5_).modifiedAt, (d_19_dt__update__tmp_h5_).draft, (d_19_dt__update__tmp_h5_).activity, (d_19_dt__update__tmp_h5_).activeTurn, (d_19_dt__update__tmp_h5_).steeringMessage, d_20_dt__update_hqueuedMessages_h0_, (d_19_dt__update__tmp_h5_).nextCursor)
                            return iife23_(_pat_let277_0)
                        return iife22_(default__.removeQ((d_15_s3_).queuedMessages, (d_8_qid_).value))
                    return iife21_(_pat_let276_0)
                d_18_s4_ = (iife20_(d_15_s3_) if (d_8_qid_).is_Some else d_15_s3_)
                return (d_18_s4_, AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CTurnComplete:
                d_21_id_ = source0_.turnId
                return default__.endTurn(s, d_21_id_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "complete")), False, AhpSkeleton.Option_None())
        if True:
            if source0_.is_CTurnCancelled:
                d_22_id_ = source0_.turnId
                return default__.endTurn(s, d_22_id_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cancelled")), False, AhpSkeleton.Option_None())
        if True:
            if source0_.is_CError:
                d_23_id_ = source0_.turnId
                d_24_e_ = source0_.err
                return default__.endTurn(s, d_23_id_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "error")), True, default__.jNoNull(d_24_e_))
        if True:
            if source0_.is_CUsage:
                d_25_id_ = source0_.turnId
                d_26_u_ = source0_.usage
                if default__.turnMatches(s, d_25_id_):
                    def iife24_(_pat_let278_0):
                        def iife25_(d_27_dt__update__tmp_h6_):
                            def iife27_(_pat_let280_0):
                                def iife28_(d_28_dt__update__tmp_h7_):
                                    def iife29_(_pat_let281_0):
                                        def iife30_(d_29_dt__update_husage_h0_):
                                            return Turn_Turn((d_28_dt__update__tmp_h7_).turnId, (d_28_dt__update__tmp_h7_).message, (d_28_dt__update__tmp_h7_).parts, (d_28_dt__update__tmp_h7_).state, d_29_dt__update_husage_h0_, (d_28_dt__update__tmp_h7_).error)
                                        return iife30_(_pat_let281_0)
                                    return iife29_(default__.jNoNull(d_26_u_))
                                return iife28_(_pat_let280_0)
                            def iife26_(_pat_let279_0):
                                def iife31_(d_30_dt__update_hactiveTurn_h1_):
                                    return ChatState_ChatState((d_27_dt__update__tmp_h6_).turns, (d_27_dt__update__tmp_h6_).title, (d_27_dt__update__tmp_h6_).status, (d_27_dt__update__tmp_h6_).modifiedAt, (d_27_dt__update__tmp_h6_).draft, (d_27_dt__update__tmp_h6_).activity, d_30_dt__update_hactiveTurn_h1_, (d_27_dt__update__tmp_h6_).steeringMessage, (d_27_dt__update__tmp_h6_).queuedMessages, (d_27_dt__update__tmp_h6_).nextCursor)
                                return iife31_(_pat_let279_0)
                            return iife26_(AhpSkeleton.Option_Some(iife27_(((pat_let_tv0_).activeTurn).value)))
                        return iife25_(_pat_let278_0)
                    return (iife24_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CResponsePart:
                d_31_id_ = source0_.turnId
                d_32_p_ = source0_.part
                if default__.turnMatches(s, d_31_id_):
                    def iife32_(_pat_let282_0):
                        def iife33_(d_33_dt__update__tmp_h8_):
                            def iife34_(_pat_let283_0):
                                def iife35_(d_34_dt__update_hactiveTurn_h2_):
                                    return ChatState_ChatState((d_33_dt__update__tmp_h8_).turns, (d_33_dt__update__tmp_h8_).title, (d_33_dt__update__tmp_h8_).status, (d_33_dt__update__tmp_h8_).modifiedAt, (d_33_dt__update__tmp_h8_).draft, (d_33_dt__update__tmp_h8_).activity, d_34_dt__update_hactiveTurn_h2_, (d_33_dt__update__tmp_h8_).steeringMessage, (d_33_dt__update__tmp_h8_).queuedMessages, (d_33_dt__update__tmp_h8_).nextCursor)
                                return iife35_(_pat_let283_0)
                            return iife34_(AhpSkeleton.Option_Some(default__.appendPart(((pat_let_tv1_).activeTurn).value, d_32_p_)))
                        return iife33_(_pat_let282_0)
                    return (iife32_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CDelta:
                d_35_id_ = source0_.turnId
                d_36_pid_ = source0_.partId
                d_37_c_ = source0_.content
                if default__.turnMatches(s, d_35_id_):
                    def iife36_(_pat_let284_0):
                        def iife37_(d_38_dt__update__tmp_h9_):
                            def iife39_(_pat_let286_0):
                                def iife40_(d_39_dt__update__tmp_h10_):
                                    def iife41_(_pat_let287_0):
                                        def iife42_(d_40_dt__update_hparts_h0_):
                                            return Turn_Turn((d_39_dt__update__tmp_h10_).turnId, (d_39_dt__update__tmp_h10_).message, d_40_dt__update_hparts_h0_, (d_39_dt__update__tmp_h10_).state, (d_39_dt__update__tmp_h10_).usage, (d_39_dt__update__tmp_h10_).error)
                                        return iife42_(_pat_let287_0)
                                    return iife41_(default__.deltaMarkdown((((pat_let_tv3_).activeTurn).value).parts, d_36_pid_, d_37_c_))
                                return iife40_(_pat_let286_0)
                            def iife38_(_pat_let285_0):
                                def iife43_(d_41_dt__update_hactiveTurn_h3_):
                                    return ChatState_ChatState((d_38_dt__update__tmp_h9_).turns, (d_38_dt__update__tmp_h9_).title, (d_38_dt__update__tmp_h9_).status, (d_38_dt__update__tmp_h9_).modifiedAt, (d_38_dt__update__tmp_h9_).draft, (d_38_dt__update__tmp_h9_).activity, d_41_dt__update_hactiveTurn_h3_, (d_38_dt__update__tmp_h9_).steeringMessage, (d_38_dt__update__tmp_h9_).queuedMessages, (d_38_dt__update__tmp_h9_).nextCursor)
                                return iife43_(_pat_let285_0)
                            return iife38_(AhpSkeleton.Option_Some(iife39_(((pat_let_tv2_).activeTurn).value)))
                        return iife37_(_pat_let284_0)
                    return (iife36_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CReasoning:
                d_42_id_ = source0_.turnId
                d_43_pid_ = source0_.partId
                d_44_c_ = source0_.content
                if default__.turnMatches(s, d_42_id_):
                    def iife44_(_pat_let288_0):
                        def iife45_(d_45_dt__update__tmp_h11_):
                            def iife47_(_pat_let290_0):
                                def iife48_(d_46_dt__update__tmp_h12_):
                                    def iife49_(_pat_let291_0):
                                        def iife50_(d_47_dt__update_hparts_h1_):
                                            return Turn_Turn((d_46_dt__update__tmp_h12_).turnId, (d_46_dt__update__tmp_h12_).message, d_47_dt__update_hparts_h1_, (d_46_dt__update__tmp_h12_).state, (d_46_dt__update__tmp_h12_).usage, (d_46_dt__update__tmp_h12_).error)
                                        return iife50_(_pat_let291_0)
                                    return iife49_(default__.deltaReasoning((((pat_let_tv5_).activeTurn).value).parts, d_43_pid_, d_44_c_))
                                return iife48_(_pat_let290_0)
                            def iife46_(_pat_let289_0):
                                def iife51_(d_48_dt__update_hactiveTurn_h4_):
                                    return ChatState_ChatState((d_45_dt__update__tmp_h11_).turns, (d_45_dt__update__tmp_h11_).title, (d_45_dt__update__tmp_h11_).status, (d_45_dt__update__tmp_h11_).modifiedAt, (d_45_dt__update__tmp_h11_).draft, (d_45_dt__update__tmp_h11_).activity, d_48_dt__update_hactiveTurn_h4_, (d_45_dt__update__tmp_h11_).steeringMessage, (d_45_dt__update__tmp_h11_).queuedMessages, (d_45_dt__update__tmp_h11_).nextCursor)
                                return iife51_(_pat_let289_0)
                            return iife46_(AhpSkeleton.Option_Some(iife47_(((pat_let_tv4_).activeTurn).value)))
                        return iife45_(_pat_let288_0)
                    return (iife44_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CToolCallStart:
                d_49_id_ = source0_.turnId
                d_50_tcId_ = source0_.toolCallId
                d_51_tn_ = source0_.toolName
                d_52_dn_ = source0_.displayName
                d_53_intent_ = source0_.intention
                d_54_contributor_ = source0_.contributor
                d_55_meta_ = source0_.meta
                if default__.turnMatches(s, d_49_id_):
                    d_56_tc_ = ToolCall_ToolCall(d_50_tcId_, d_51_tn_, d_52_dn_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "streaming")), d_53_intent_, d_54_contributor_, d_55_meta_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())
                    def iife52_(_pat_let292_0):
                        def iife53_(d_57_dt__update__tmp_h13_):
                            def iife55_(_pat_let294_0):
                                def iife56_(d_58_dt__update__tmp_h14_):
                                    def iife57_(_pat_let295_0):
                                        def iife58_(d_59_dt__update_hparts_h2_):
                                            return Turn_Turn((d_58_dt__update__tmp_h14_).turnId, (d_58_dt__update__tmp_h14_).message, d_59_dt__update_hparts_h2_, (d_58_dt__update__tmp_h14_).state, (d_58_dt__update__tmp_h14_).usage, (d_58_dt__update__tmp_h14_).error)
                                        return iife58_(_pat_let295_0)
                                    return iife57_(((((pat_let_tv7_).activeTurn).value).parts) + (_dafny.SeqWithoutIsStrInference([Part_PToolCall(d_56_tc_)])))
                                return iife56_(_pat_let294_0)
                            def iife54_(_pat_let293_0):
                                def iife59_(d_60_dt__update_hactiveTurn_h5_):
                                    return ChatState_ChatState((d_57_dt__update__tmp_h13_).turns, (d_57_dt__update__tmp_h13_).title, (d_57_dt__update__tmp_h13_).status, (d_57_dt__update__tmp_h13_).modifiedAt, (d_57_dt__update__tmp_h13_).draft, (d_57_dt__update__tmp_h13_).activity, d_60_dt__update_hactiveTurn_h5_, (d_57_dt__update__tmp_h13_).steeringMessage, (d_57_dt__update__tmp_h13_).queuedMessages, (d_57_dt__update__tmp_h13_).nextCursor)
                                return iife59_(_pat_let293_0)
                            return iife54_(AhpSkeleton.Option_Some(iife55_(((pat_let_tv6_).activeTurn).value)))
                        return iife53_(_pat_let292_0)
                    return (iife52_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CToolCallDelta:
                d_61_id_ = source0_.turnId
                d_62_tcId_ = source0_.toolCallId
                d_63_c_ = source0_.content
                d_64_inv_ = source0_.invocationMessage
                d_65_meta_ = source0_.meta
                if default__.turnMatches(s, d_61_id_):
                    def iife60_(_pat_let296_0):
                        def iife61_(d_66_dt__update__tmp_h15_):
                            def iife63_(_pat_let298_0):
                                def iife64_(d_67_dt__update__tmp_h16_):
                                    def lambda0_(d_68_c_, d_69_inv_, d_70_meta_):
                                        def lambda1_(d_71_x_):
                                            return default__.deltaOne(d_71_x_, d_68_c_, d_69_inv_, d_70_meta_)

                                        return lambda1_

                                    def iife65_(_pat_let299_0):
                                        def iife66_(d_72_dt__update_hparts_h3_):
                                            return Turn_Turn((d_67_dt__update__tmp_h16_).turnId, (d_67_dt__update__tmp_h16_).message, d_72_dt__update_hparts_h3_, (d_67_dt__update__tmp_h16_).state, (d_67_dt__update__tmp_h16_).usage, (d_67_dt__update__tmp_h16_).error)
                                        return iife66_(_pat_let299_0)
                                    return iife65_(default__.mapTC((((pat_let_tv9_).activeTurn).value).parts, d_62_tcId_, lambda0_(d_63_c_, d_64_inv_, d_65_meta_)))
                                return iife64_(_pat_let298_0)
                            def iife62_(_pat_let297_0):
                                def iife67_(d_73_dt__update_hactiveTurn_h6_):
                                    return ChatState_ChatState((d_66_dt__update__tmp_h15_).turns, (d_66_dt__update__tmp_h15_).title, (d_66_dt__update__tmp_h15_).status, (d_66_dt__update__tmp_h15_).modifiedAt, (d_66_dt__update__tmp_h15_).draft, (d_66_dt__update__tmp_h15_).activity, d_73_dt__update_hactiveTurn_h6_, (d_66_dt__update__tmp_h15_).steeringMessage, (d_66_dt__update__tmp_h15_).queuedMessages, (d_66_dt__update__tmp_h15_).nextCursor)
                                return iife67_(_pat_let297_0)
                            return iife62_(AhpSkeleton.Option_Some(iife63_(((pat_let_tv8_).activeTurn).value)))
                        return iife61_(_pat_let296_0)
                    return (iife60_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CToolCallReady:
                d_74_id_ = source0_.turnId
                d_75_tcId_ = source0_.toolCallId
                d_76_inv_ = source0_.invocationMessage
                d_77_input_ = source0_.toolInput
                d_78_title_ = source0_.confirmationTitle
                d_79_risk_ = source0_.riskAssessment
                d_80_edits_ = source0_.edits
                d_81_editable_ = source0_.editable
                d_82_options_ = source0_.options
                d_83_confirmed_ = source0_.confirmed
                d_84_meta_ = source0_.meta
                if default__.turnMatches(s, d_74_id_):
                    def iife68_(_pat_let300_0):
                        def iife69_(d_86_dt__update__tmp_h17_):
                            def iife71_(_pat_let302_0):
                                def iife72_(d_87_dt__update__tmp_h18_):
                                    def lambda2_(d_88_inv_, d_89_input_, d_90_title_, d_91_risk_, d_92_edits_, d_93_editable_, d_94_options_, d_95_confirmed_, d_96_meta_):
                                        def lambda3_(d_97_x_):
                                            return default__.readyOne(d_97_x_, d_88_inv_, d_89_input_, d_90_title_, d_91_risk_, d_92_edits_, d_93_editable_, d_94_options_, d_95_confirmed_, d_96_meta_)

                                        return lambda3_

                                    def iife73_(_pat_let303_0):
                                        def iife74_(d_98_dt__update_hparts_h4_):
                                            return Turn_Turn((d_87_dt__update__tmp_h18_).turnId, (d_87_dt__update__tmp_h18_).message, d_98_dt__update_hparts_h4_, (d_87_dt__update__tmp_h18_).state, (d_87_dt__update__tmp_h18_).usage, (d_87_dt__update__tmp_h18_).error)
                                        return iife74_(_pat_let303_0)
                                    return iife73_(default__.mapTC((((pat_let_tv11_).activeTurn).value).parts, d_75_tcId_, lambda2_(d_76_inv_, d_77_input_, d_78_title_, d_79_risk_, d_80_edits_, d_81_editable_, d_82_options_, d_83_confirmed_, d_84_meta_)))
                                return iife72_(_pat_let302_0)
                            def iife70_(_pat_let301_0):
                                def iife75_(d_99_dt__update_hactiveTurn_h7_):
                                    return ChatState_ChatState((d_86_dt__update__tmp_h17_).turns, (d_86_dt__update__tmp_h17_).title, (d_86_dt__update__tmp_h17_).status, (d_86_dt__update__tmp_h17_).modifiedAt, (d_86_dt__update__tmp_h17_).draft, (d_86_dt__update__tmp_h17_).activity, d_99_dt__update_hactiveTurn_h7_, (d_86_dt__update__tmp_h17_).steeringMessage, (d_86_dt__update__tmp_h17_).queuedMessages, (d_86_dt__update__tmp_h17_).nextCursor)
                                return iife75_(_pat_let301_0)
                            return iife70_(AhpSkeleton.Option_Some(iife71_(((pat_let_tv10_).activeTurn).value)))
                        return iife69_(_pat_let300_0)
                    d_85_s1_ = iife68_(s)
                    def iife76_(_pat_let304_0):
                        def iife77_(d_100_dt__update__tmp_h19_):
                            def iife78_(_pat_let305_0):
                                def iife79_(d_101_dt__update_hstatus_h1_):
                                    return ChatState_ChatState((d_100_dt__update__tmp_h19_).turns, (d_100_dt__update__tmp_h19_).title, d_101_dt__update_hstatus_h1_, (d_100_dt__update__tmp_h19_).modifiedAt, (d_100_dt__update__tmp_h19_).draft, (d_100_dt__update__tmp_h19_).activity, (d_100_dt__update__tmp_h19_).activeTurn, (d_100_dt__update__tmp_h19_).steeringMessage, (d_100_dt__update__tmp_h19_).queuedMessages, (d_100_dt__update__tmp_h19_).nextCursor)
                                return iife79_(_pat_let305_0)
                            return iife78_(default__.summaryStatus(d_85_s1_, False))
                        return iife77_(_pat_let304_0)
                    return (iife76_(d_85_s1_), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CToolCallConfirmed:
                d_102_id_ = source0_.turnId
                d_103_tcId_ = source0_.toolCallId
                d_104_approved_ = source0_.approved
                d_105_cv_ = source0_.confirmedVal
                d_106_rsn_ = source0_.reason
                d_107_reasonMessage_ = source0_.reasonMessage
                d_108_userSuggestion_ = source0_.userSuggestion
                d_109_edited_ = source0_.editedToolInput
                d_110_selectedOptionId_ = source0_.selectedOptionId
                d_111_meta_ = source0_.meta
                if default__.turnMatches(s, d_102_id_):
                    def iife80_(_pat_let306_0):
                        def iife81_(d_113_dt__update__tmp_h20_):
                            def iife83_(_pat_let308_0):
                                def iife84_(d_114_dt__update__tmp_h21_):
                                    def lambda4_(d_115_approved_, d_116_cv_, d_117_rsn_, d_118_reasonMessage_, d_119_userSuggestion_, d_120_edited_, d_121_selectedOptionId_, d_122_meta_):
                                        def lambda5_(d_123_x_):
                                            return default__.confirmOne(d_123_x_, d_115_approved_, d_116_cv_, d_117_rsn_, d_118_reasonMessage_, d_119_userSuggestion_, d_120_edited_, d_121_selectedOptionId_, d_122_meta_)

                                        return lambda5_

                                    def iife85_(_pat_let309_0):
                                        def iife86_(d_124_dt__update_hparts_h5_):
                                            return Turn_Turn((d_114_dt__update__tmp_h21_).turnId, (d_114_dt__update__tmp_h21_).message, d_124_dt__update_hparts_h5_, (d_114_dt__update__tmp_h21_).state, (d_114_dt__update__tmp_h21_).usage, (d_114_dt__update__tmp_h21_).error)
                                        return iife86_(_pat_let309_0)
                                    return iife85_(default__.mapTC((((pat_let_tv13_).activeTurn).value).parts, d_103_tcId_, lambda4_(d_104_approved_, d_105_cv_, d_106_rsn_, d_107_reasonMessage_, d_108_userSuggestion_, d_109_edited_, d_110_selectedOptionId_, d_111_meta_)))
                                return iife84_(_pat_let308_0)
                            def iife82_(_pat_let307_0):
                                def iife87_(d_125_dt__update_hactiveTurn_h8_):
                                    return ChatState_ChatState((d_113_dt__update__tmp_h20_).turns, (d_113_dt__update__tmp_h20_).title, (d_113_dt__update__tmp_h20_).status, (d_113_dt__update__tmp_h20_).modifiedAt, (d_113_dt__update__tmp_h20_).draft, (d_113_dt__update__tmp_h20_).activity, d_125_dt__update_hactiveTurn_h8_, (d_113_dt__update__tmp_h20_).steeringMessage, (d_113_dt__update__tmp_h20_).queuedMessages, (d_113_dt__update__tmp_h20_).nextCursor)
                                return iife87_(_pat_let307_0)
                            return iife82_(AhpSkeleton.Option_Some(iife83_(((pat_let_tv12_).activeTurn).value)))
                        return iife81_(_pat_let306_0)
                    d_112_s1_ = iife80_(s)
                    def iife88_(_pat_let310_0):
                        def iife89_(d_126_dt__update__tmp_h22_):
                            def iife90_(_pat_let311_0):
                                def iife91_(d_127_dt__update_hstatus_h2_):
                                    return ChatState_ChatState((d_126_dt__update__tmp_h22_).turns, (d_126_dt__update__tmp_h22_).title, d_127_dt__update_hstatus_h2_, (d_126_dt__update__tmp_h22_).modifiedAt, (d_126_dt__update__tmp_h22_).draft, (d_126_dt__update__tmp_h22_).activity, (d_126_dt__update__tmp_h22_).activeTurn, (d_126_dt__update__tmp_h22_).steeringMessage, (d_126_dt__update__tmp_h22_).queuedMessages, (d_126_dt__update__tmp_h22_).nextCursor)
                                return iife91_(_pat_let311_0)
                            return iife90_(default__.summaryStatus(d_112_s1_, False))
                        return iife89_(_pat_let310_0)
                    return (iife88_(d_112_s1_), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CToolCallAuthRequired:
                d_128_id_ = source0_.turnId
                d_129_tcId_ = source0_.toolCallId
                d_130_auth_ = source0_.auth
                d_131_meta_ = source0_.meta
                if default__.turnMatches(s, d_128_id_):
                    def iife92_(_pat_let312_0):
                        def iife93_(d_133_dt__update__tmp_h23_):
                            def iife95_(_pat_let314_0):
                                def iife96_(d_134_dt__update__tmp_h24_):
                                    def lambda6_(d_135_auth_, d_136_meta_):
                                        def lambda7_(d_137_x_):
                                            return default__.authRequiredOne(d_137_x_, d_135_auth_, d_136_meta_)

                                        return lambda7_

                                    def iife97_(_pat_let315_0):
                                        def iife98_(d_138_dt__update_hparts_h6_):
                                            return Turn_Turn((d_134_dt__update__tmp_h24_).turnId, (d_134_dt__update__tmp_h24_).message, d_138_dt__update_hparts_h6_, (d_134_dt__update__tmp_h24_).state, (d_134_dt__update__tmp_h24_).usage, (d_134_dt__update__tmp_h24_).error)
                                        return iife98_(_pat_let315_0)
                                    return iife97_(default__.mapTC((((pat_let_tv15_).activeTurn).value).parts, d_129_tcId_, lambda6_(d_130_auth_, d_131_meta_)))
                                return iife96_(_pat_let314_0)
                            def iife94_(_pat_let313_0):
                                def iife99_(d_139_dt__update_hactiveTurn_h9_):
                                    return ChatState_ChatState((d_133_dt__update__tmp_h23_).turns, (d_133_dt__update__tmp_h23_).title, (d_133_dt__update__tmp_h23_).status, (d_133_dt__update__tmp_h23_).modifiedAt, (d_133_dt__update__tmp_h23_).draft, (d_133_dt__update__tmp_h23_).activity, d_139_dt__update_hactiveTurn_h9_, (d_133_dt__update__tmp_h23_).steeringMessage, (d_133_dt__update__tmp_h23_).queuedMessages, (d_133_dt__update__tmp_h23_).nextCursor)
                                return iife99_(_pat_let313_0)
                            return iife94_(AhpSkeleton.Option_Some(iife95_(((pat_let_tv14_).activeTurn).value)))
                        return iife93_(_pat_let312_0)
                    d_132_s1_ = iife92_(s)
                    def iife100_(_pat_let316_0):
                        def iife101_(d_140_dt__update__tmp_h25_):
                            def iife102_(_pat_let317_0):
                                def iife103_(d_141_dt__update_hstatus_h3_):
                                    return ChatState_ChatState((d_140_dt__update__tmp_h25_).turns, (d_140_dt__update__tmp_h25_).title, d_141_dt__update_hstatus_h3_, (d_140_dt__update__tmp_h25_).modifiedAt, (d_140_dt__update__tmp_h25_).draft, (d_140_dt__update__tmp_h25_).activity, (d_140_dt__update__tmp_h25_).activeTurn, (d_140_dt__update__tmp_h25_).steeringMessage, (d_140_dt__update__tmp_h25_).queuedMessages, (d_140_dt__update__tmp_h25_).nextCursor)
                                return iife103_(_pat_let317_0)
                            return iife102_(default__.summaryStatus(d_132_s1_, False))
                        return iife101_(_pat_let316_0)
                    return (iife100_(d_132_s1_), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CToolCallAuthResolved:
                d_142_id_ = source0_.turnId
                d_143_tcId_ = source0_.toolCallId
                d_144_meta_ = source0_.meta
                if default__.turnMatches(s, d_142_id_):
                    def iife104_(_pat_let318_0):
                        def iife105_(d_146_dt__update__tmp_h26_):
                            def iife107_(_pat_let320_0):
                                def iife108_(d_147_dt__update__tmp_h27_):
                                    def lambda8_(d_148_meta_):
                                        def lambda9_(d_149_x_):
                                            return default__.authResolvedOne(d_149_x_, d_148_meta_)

                                        return lambda9_

                                    def iife109_(_pat_let321_0):
                                        def iife110_(d_150_dt__update_hparts_h7_):
                                            return Turn_Turn((d_147_dt__update__tmp_h27_).turnId, (d_147_dt__update__tmp_h27_).message, d_150_dt__update_hparts_h7_, (d_147_dt__update__tmp_h27_).state, (d_147_dt__update__tmp_h27_).usage, (d_147_dt__update__tmp_h27_).error)
                                        return iife110_(_pat_let321_0)
                                    return iife109_(default__.mapTC((((pat_let_tv17_).activeTurn).value).parts, d_143_tcId_, lambda8_(d_144_meta_)))
                                return iife108_(_pat_let320_0)
                            def iife106_(_pat_let319_0):
                                def iife111_(d_151_dt__update_hactiveTurn_h10_):
                                    return ChatState_ChatState((d_146_dt__update__tmp_h26_).turns, (d_146_dt__update__tmp_h26_).title, (d_146_dt__update__tmp_h26_).status, (d_146_dt__update__tmp_h26_).modifiedAt, (d_146_dt__update__tmp_h26_).draft, (d_146_dt__update__tmp_h26_).activity, d_151_dt__update_hactiveTurn_h10_, (d_146_dt__update__tmp_h26_).steeringMessage, (d_146_dt__update__tmp_h26_).queuedMessages, (d_146_dt__update__tmp_h26_).nextCursor)
                                return iife111_(_pat_let319_0)
                            return iife106_(AhpSkeleton.Option_Some(iife107_(((pat_let_tv16_).activeTurn).value)))
                        return iife105_(_pat_let318_0)
                    d_145_s1_ = iife104_(s)
                    def iife112_(_pat_let322_0):
                        def iife113_(d_152_dt__update__tmp_h28_):
                            def iife114_(_pat_let323_0):
                                def iife115_(d_153_dt__update_hstatus_h4_):
                                    return ChatState_ChatState((d_152_dt__update__tmp_h28_).turns, (d_152_dt__update__tmp_h28_).title, d_153_dt__update_hstatus_h4_, (d_152_dt__update__tmp_h28_).modifiedAt, (d_152_dt__update__tmp_h28_).draft, (d_152_dt__update__tmp_h28_).activity, (d_152_dt__update__tmp_h28_).activeTurn, (d_152_dt__update__tmp_h28_).steeringMessage, (d_152_dt__update__tmp_h28_).queuedMessages, (d_152_dt__update__tmp_h28_).nextCursor)
                                return iife115_(_pat_let323_0)
                            return iife114_(default__.summaryStatus(d_145_s1_, False))
                        return iife113_(_pat_let322_0)
                    return (iife112_(d_145_s1_), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CToolCallComplete:
                d_154_id_ = source0_.turnId
                d_155_tcId_ = source0_.toolCallId
                d_156_ok_ = source0_.success
                d_157_past_ = source0_.pastTenseMessage
                d_158_resultContent_ = source0_.resultContent
                d_159_structured_ = source0_.structuredContent
                d_160_err_ = source0_.error
                d_161_rrc_ = source0_.requiresResultConfirmation
                d_162_meta_ = source0_.meta
                if default__.turnMatches(s, d_154_id_):
                    def iife116_(_pat_let324_0):
                        def iife117_(d_164_dt__update__tmp_h29_):
                            def iife119_(_pat_let326_0):
                                def iife120_(d_165_dt__update__tmp_h30_):
                                    def lambda10_(d_166_ok_, d_167_past_, d_168_resultContent_, d_169_structured_, d_170_err_, d_171_rrc_, d_172_meta_):
                                        def lambda11_(d_173_x_):
                                            return default__.completeOne(d_173_x_, d_166_ok_, d_167_past_, d_168_resultContent_, d_169_structured_, d_170_err_, d_171_rrc_, d_172_meta_)

                                        return lambda11_

                                    def iife121_(_pat_let327_0):
                                        def iife122_(d_174_dt__update_hparts_h8_):
                                            return Turn_Turn((d_165_dt__update__tmp_h30_).turnId, (d_165_dt__update__tmp_h30_).message, d_174_dt__update_hparts_h8_, (d_165_dt__update__tmp_h30_).state, (d_165_dt__update__tmp_h30_).usage, (d_165_dt__update__tmp_h30_).error)
                                        return iife122_(_pat_let327_0)
                                    return iife121_(default__.mapTC((((pat_let_tv19_).activeTurn).value).parts, d_155_tcId_, lambda10_(d_156_ok_, d_157_past_, d_158_resultContent_, d_159_structured_, d_160_err_, d_161_rrc_, d_162_meta_)))
                                return iife120_(_pat_let326_0)
                            def iife118_(_pat_let325_0):
                                def iife123_(d_175_dt__update_hactiveTurn_h11_):
                                    return ChatState_ChatState((d_164_dt__update__tmp_h29_).turns, (d_164_dt__update__tmp_h29_).title, (d_164_dt__update__tmp_h29_).status, (d_164_dt__update__tmp_h29_).modifiedAt, (d_164_dt__update__tmp_h29_).draft, (d_164_dt__update__tmp_h29_).activity, d_175_dt__update_hactiveTurn_h11_, (d_164_dt__update__tmp_h29_).steeringMessage, (d_164_dt__update__tmp_h29_).queuedMessages, (d_164_dt__update__tmp_h29_).nextCursor)
                                return iife123_(_pat_let325_0)
                            return iife118_(AhpSkeleton.Option_Some(iife119_(((pat_let_tv18_).activeTurn).value)))
                        return iife117_(_pat_let324_0)
                    d_163_s1_ = iife116_(s)
                    def iife124_(_pat_let328_0):
                        def iife125_(d_176_dt__update__tmp_h31_):
                            def iife126_(_pat_let329_0):
                                def iife127_(d_177_dt__update_hstatus_h5_):
                                    return ChatState_ChatState((d_176_dt__update__tmp_h31_).turns, (d_176_dt__update__tmp_h31_).title, d_177_dt__update_hstatus_h5_, (d_176_dt__update__tmp_h31_).modifiedAt, (d_176_dt__update__tmp_h31_).draft, (d_176_dt__update__tmp_h31_).activity, (d_176_dt__update__tmp_h31_).activeTurn, (d_176_dt__update__tmp_h31_).steeringMessage, (d_176_dt__update__tmp_h31_).queuedMessages, (d_176_dt__update__tmp_h31_).nextCursor)
                                return iife127_(_pat_let329_0)
                            return iife126_(default__.summaryStatus(d_163_s1_, False))
                        return iife125_(_pat_let328_0)
                    return (iife124_(d_163_s1_), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CToolCallResultConfirmed:
                d_178_id_ = source0_.turnId
                d_179_tcId_ = source0_.toolCallId
                d_180_approved_ = source0_.approved
                d_181_meta_ = source0_.meta
                if default__.turnMatches(s, d_178_id_):
                    def iife128_(_pat_let330_0):
                        def iife129_(d_183_dt__update__tmp_h32_):
                            def iife131_(_pat_let332_0):
                                def iife132_(d_184_dt__update__tmp_h33_):
                                    def lambda12_(d_185_approved_, d_186_meta_):
                                        def lambda13_(d_187_x_):
                                            return default__.resultConfirmOne(d_187_x_, d_185_approved_, d_186_meta_)

                                        return lambda13_

                                    def iife133_(_pat_let333_0):
                                        def iife134_(d_188_dt__update_hparts_h9_):
                                            return Turn_Turn((d_184_dt__update__tmp_h33_).turnId, (d_184_dt__update__tmp_h33_).message, d_188_dt__update_hparts_h9_, (d_184_dt__update__tmp_h33_).state, (d_184_dt__update__tmp_h33_).usage, (d_184_dt__update__tmp_h33_).error)
                                        return iife134_(_pat_let333_0)
                                    return iife133_(default__.mapTC((((pat_let_tv21_).activeTurn).value).parts, d_179_tcId_, lambda12_(d_180_approved_, d_181_meta_)))
                                return iife132_(_pat_let332_0)
                            def iife130_(_pat_let331_0):
                                def iife135_(d_189_dt__update_hactiveTurn_h12_):
                                    return ChatState_ChatState((d_183_dt__update__tmp_h32_).turns, (d_183_dt__update__tmp_h32_).title, (d_183_dt__update__tmp_h32_).status, (d_183_dt__update__tmp_h32_).modifiedAt, (d_183_dt__update__tmp_h32_).draft, (d_183_dt__update__tmp_h32_).activity, d_189_dt__update_hactiveTurn_h12_, (d_183_dt__update__tmp_h32_).steeringMessage, (d_183_dt__update__tmp_h32_).queuedMessages, (d_183_dt__update__tmp_h32_).nextCursor)
                                return iife135_(_pat_let331_0)
                            return iife130_(AhpSkeleton.Option_Some(iife131_(((pat_let_tv20_).activeTurn).value)))
                        return iife129_(_pat_let330_0)
                    d_182_s1_ = iife128_(s)
                    def iife136_(_pat_let334_0):
                        def iife137_(d_190_dt__update__tmp_h34_):
                            def iife138_(_pat_let335_0):
                                def iife139_(d_191_dt__update_hstatus_h6_):
                                    return ChatState_ChatState((d_190_dt__update__tmp_h34_).turns, (d_190_dt__update__tmp_h34_).title, d_191_dt__update_hstatus_h6_, (d_190_dt__update__tmp_h34_).modifiedAt, (d_190_dt__update__tmp_h34_).draft, (d_190_dt__update__tmp_h34_).activity, (d_190_dt__update__tmp_h34_).activeTurn, (d_190_dt__update__tmp_h34_).steeringMessage, (d_190_dt__update__tmp_h34_).queuedMessages, (d_190_dt__update__tmp_h34_).nextCursor)
                                return iife139_(_pat_let335_0)
                            return iife138_(default__.summaryStatus(d_182_s1_, False))
                        return iife137_(_pat_let334_0)
                    return (iife136_(d_182_s1_), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CToolCallContentChanged:
                d_192_id_ = source0_.turnId
                d_193_tcId_ = source0_.toolCallId
                d_194_c_ = source0_.newContent
                d_195_meta_ = source0_.meta
                if default__.turnMatches(s, d_192_id_):
                    def iife140_(_pat_let336_0):
                        def iife141_(d_196_dt__update__tmp_h35_):
                            def iife143_(_pat_let338_0):
                                def iife144_(d_197_dt__update__tmp_h36_):
                                    def lambda14_(d_198_c_, d_199_meta_):
                                        def lambda15_(d_200_x_):
                                            return default__.contentChangedOne(d_200_x_, d_198_c_, d_199_meta_)

                                        return lambda15_

                                    def iife145_(_pat_let339_0):
                                        def iife146_(d_201_dt__update_hparts_h10_):
                                            return Turn_Turn((d_197_dt__update__tmp_h36_).turnId, (d_197_dt__update__tmp_h36_).message, d_201_dt__update_hparts_h10_, (d_197_dt__update__tmp_h36_).state, (d_197_dt__update__tmp_h36_).usage, (d_197_dt__update__tmp_h36_).error)
                                        return iife146_(_pat_let339_0)
                                    return iife145_(default__.mapTC((((pat_let_tv23_).activeTurn).value).parts, d_193_tcId_, lambda14_(d_194_c_, d_195_meta_)))
                                return iife144_(_pat_let338_0)
                            def iife142_(_pat_let337_0):
                                def iife147_(d_202_dt__update_hactiveTurn_h13_):
                                    return ChatState_ChatState((d_196_dt__update__tmp_h35_).turns, (d_196_dt__update__tmp_h35_).title, (d_196_dt__update__tmp_h35_).status, (d_196_dt__update__tmp_h35_).modifiedAt, (d_196_dt__update__tmp_h35_).draft, (d_196_dt__update__tmp_h35_).activity, d_202_dt__update_hactiveTurn_h13_, (d_196_dt__update__tmp_h35_).steeringMessage, (d_196_dt__update__tmp_h35_).queuedMessages, (d_196_dt__update__tmp_h35_).nextCursor)
                                return iife147_(_pat_let337_0)
                            return iife142_(AhpSkeleton.Option_Some(iife143_(((pat_let_tv22_).activeTurn).value)))
                        return iife141_(_pat_let336_0)
                    return (iife140_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CTruncated:
                d_203_idOpt_ = source0_.upToTurnId
                if ((d_203_idOpt_).is_Some) and (not(default__.hasTurn((s).turns, (d_203_idOpt_).value))):
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
                elif True:
                    d_204_kept_ = (default__.keepUpTo((s).turns, (d_203_idOpt_).value) if (d_203_idOpt_).is_Some else _dafny.SeqWithoutIsStrInference([]))
                    def iife148_(_pat_let340_0):
                        def iife149_(d_206_dt__update__tmp_h37_):
                            def iife150_(_pat_let341_0):
                                def iife151_(d_207_dt__update_hnextCursor_h0_):
                                    def iife152_(_pat_let342_0):
                                        def iife153_(d_208_dt__update_hactiveTurn_h14_):
                                            def iife154_(_pat_let343_0):
                                                def iife155_(d_209_dt__update_hturns_h0_):
                                                    return ChatState_ChatState(d_209_dt__update_hturns_h0_, (d_206_dt__update__tmp_h37_).title, (d_206_dt__update__tmp_h37_).status, (d_206_dt__update__tmp_h37_).modifiedAt, (d_206_dt__update__tmp_h37_).draft, (d_206_dt__update__tmp_h37_).activity, d_208_dt__update_hactiveTurn_h14_, (d_206_dt__update__tmp_h37_).steeringMessage, (d_206_dt__update__tmp_h37_).queuedMessages, d_207_dt__update_hnextCursor_h0_)
                                                return iife155_(_pat_let343_0)
                                            return iife154_(d_204_kept_)
                                        return iife153_(_pat_let342_0)
                                    return iife152_(AhpSkeleton.Option_None())
                                return iife151_(_pat_let341_0)
                            return iife150_(((pat_let_tv24_).nextCursor if (d_203_idOpt_).is_Some else AhpSkeleton.Option_None()))
                        return iife149_(_pat_let340_0)
                    d_205_next_ = iife148_(s)
                    def iife156_(_pat_let344_0):
                        def iife157_(d_210_dt__update__tmp_h38_):
                            def iife158_(_pat_let345_0):
                                def iife159_(d_211_dt__update_hstatus_h7_):
                                    return ChatState_ChatState((d_210_dt__update__tmp_h38_).turns, (d_210_dt__update__tmp_h38_).title, d_211_dt__update_hstatus_h7_, (d_210_dt__update__tmp_h38_).modifiedAt, (d_210_dt__update__tmp_h38_).draft, (d_210_dt__update__tmp_h38_).activity, (d_210_dt__update__tmp_h38_).activeTurn, (d_210_dt__update__tmp_h38_).steeringMessage, (d_210_dt__update__tmp_h38_).queuedMessages, (d_210_dt__update__tmp_h38_).nextCursor)
                                return iife159_(_pat_let345_0)
                            return iife158_(default__.summaryStatus(d_205_next_, False))
                        return iife157_(_pat_let344_0)
                    return (iife156_(d_205_next_), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CQueuedReordered:
                d_212_order_ = source0_.order
                def iife160_(_pat_let346_0):
                    def iife161_(d_213_dt__update__tmp_h39_):
                        def iife162_(_pat_let347_0):
                            def iife163_(d_214_dt__update_hqueuedMessages_h1_):
                                return ChatState_ChatState((d_213_dt__update__tmp_h39_).turns, (d_213_dt__update__tmp_h39_).title, (d_213_dt__update__tmp_h39_).status, (d_213_dt__update__tmp_h39_).modifiedAt, (d_213_dt__update__tmp_h39_).draft, (d_213_dt__update__tmp_h39_).activity, (d_213_dt__update__tmp_h39_).activeTurn, (d_213_dt__update__tmp_h39_).steeringMessage, d_214_dt__update_hqueuedMessages_h1_, (d_213_dt__update__tmp_h39_).nextCursor)
                            return iife163_(_pat_let347_0)
                        return iife162_(default__.reorderQ((pat_let_tv25_).queuedMessages, d_212_order_))
                    return iife161_(_pat_let346_0)
                return (iife160_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CTurnsLoaded:
                d_215_loaded_ = source0_.loaded
                d_216_cursor_ = source0_.cursor
                def iife164_(_pat_let348_0):
                    def iife165_(d_217_dt__update__tmp_h40_):
                        def iife166_(_pat_let349_0):
                            def iife167_(d_218_dt__update_hnextCursor_h1_):
                                def iife168_(_pat_let350_0):
                                    def iife169_(d_219_dt__update_hturns_h1_):
                                        return ChatState_ChatState(d_219_dt__update_hturns_h1_, (d_217_dt__update__tmp_h40_).title, (d_217_dt__update__tmp_h40_).status, (d_217_dt__update__tmp_h40_).modifiedAt, (d_217_dt__update__tmp_h40_).draft, (d_217_dt__update__tmp_h40_).activity, (d_217_dt__update__tmp_h40_).activeTurn, (d_217_dt__update__tmp_h40_).steeringMessage, (d_217_dt__update__tmp_h40_).queuedMessages, d_218_dt__update_hnextCursor_h1_)
                                    return iife169_(_pat_let350_0)
                                return iife168_(default__.dedupPrepend(d_215_loaded_, (pat_let_tv26_).turns))
                            return iife167_(_pat_let349_0)
                        return iife166_(d_216_cursor_)
                    return iife165_(_pat_let348_0)
                return (iife164_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CInputRequested:
                d_220_req_ = source0_.req
                if ((s).activeTurn).is_Some:
                    d_221_at_ = ((s).activeTurn).value
                    def iife170_(_pat_let351_0):
                        def iife171_(d_223_dt__update__tmp_h41_):
                            def iife173_(_pat_let353_0):
                                def iife174_(d_224_dt__update__tmp_h42_):
                                    def iife175_(_pat_let354_0):
                                        def iife176_(d_225_dt__update_hparts_h11_):
                                            return Turn_Turn((d_224_dt__update__tmp_h42_).turnId, (d_224_dt__update__tmp_h42_).message, d_225_dt__update_hparts_h11_, (d_224_dt__update__tmp_h42_).state, (d_224_dt__update__tmp_h42_).usage, (d_224_dt__update__tmp_h42_).error)
                                        return iife176_(_pat_let354_0)
                                    return iife175_(default__.upsertInputPart((d_221_at_).parts, d_220_req_))
                                return iife174_(_pat_let353_0)
                            def iife172_(_pat_let352_0):
                                def iife177_(d_226_dt__update_hactiveTurn_h15_):
                                    return ChatState_ChatState((d_223_dt__update__tmp_h41_).turns, (d_223_dt__update__tmp_h41_).title, (d_223_dt__update__tmp_h41_).status, (d_223_dt__update__tmp_h41_).modifiedAt, (d_223_dt__update__tmp_h41_).draft, (d_223_dt__update__tmp_h41_).activity, d_226_dt__update_hactiveTurn_h15_, (d_223_dt__update__tmp_h41_).steeringMessage, (d_223_dt__update__tmp_h41_).queuedMessages, (d_223_dt__update__tmp_h41_).nextCursor)
                                return iife177_(_pat_let352_0)
                            return iife172_(AhpSkeleton.Option_Some(iife173_(d_221_at_)))
                        return iife171_(_pat_let351_0)
                    d_222_next_ = iife170_(s)
                    def iife178_(_pat_let355_0):
                        def iife179_(d_227_dt__update__tmp_h43_):
                            def iife180_(_pat_let356_0):
                                def iife181_(d_228_dt__update_hstatus_h8_):
                                    return ChatState_ChatState((d_227_dt__update__tmp_h43_).turns, (d_227_dt__update__tmp_h43_).title, d_228_dt__update_hstatus_h8_, (d_227_dt__update__tmp_h43_).modifiedAt, (d_227_dt__update__tmp_h43_).draft, (d_227_dt__update__tmp_h43_).activity, (d_227_dt__update__tmp_h43_).activeTurn, (d_227_dt__update__tmp_h43_).steeringMessage, (d_227_dt__update__tmp_h43_).queuedMessages, (d_227_dt__update__tmp_h43_).nextCursor)
                                return iife181_(_pat_let356_0)
                            return iife180_(default__.clearBit(default__.summaryStatus(d_222_next_, False), default__.READ))
                        return iife179_(_pat_let355_0)
                    return (iife178_(d_222_next_), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CInputAnswerChanged:
                d_229_reqId_ = source0_.requestId
                d_230_qId_ = source0_.questionId
                d_231_ans_ = source0_.answer
                if (((s).activeTurn).is_Some) and (default__.hasOpenReq((((s).activeTurn).value).parts, d_229_reqId_)):
                    d_232_at_ = ((s).activeTurn).value
                    def iife182_(_pat_let357_0):
                        def iife183_(d_233_dt__update__tmp_h44_):
                            def iife185_(_pat_let359_0):
                                def iife186_(d_234_dt__update__tmp_h45_):
                                    def iife187_(_pat_let360_0):
                                        def iife188_(d_235_dt__update_hparts_h12_):
                                            return Turn_Turn((d_234_dt__update__tmp_h45_).turnId, (d_234_dt__update__tmp_h45_).message, d_235_dt__update_hparts_h12_, (d_234_dt__update__tmp_h45_).state, (d_234_dt__update__tmp_h45_).usage, (d_234_dt__update__tmp_h45_).error)
                                        return iife188_(_pat_let360_0)
                                    return iife187_(default__.changeAnswerPart((d_232_at_).parts, d_229_reqId_, d_230_qId_, d_231_ans_))
                                return iife186_(_pat_let359_0)
                            def iife184_(_pat_let358_0):
                                def iife189_(d_236_dt__update_hactiveTurn_h16_):
                                    return ChatState_ChatState((d_233_dt__update__tmp_h44_).turns, (d_233_dt__update__tmp_h44_).title, (d_233_dt__update__tmp_h44_).status, (d_233_dt__update__tmp_h44_).modifiedAt, (d_233_dt__update__tmp_h44_).draft, (d_233_dt__update__tmp_h44_).activity, d_236_dt__update_hactiveTurn_h16_, (d_233_dt__update__tmp_h44_).steeringMessage, (d_233_dt__update__tmp_h44_).queuedMessages, (d_233_dt__update__tmp_h44_).nextCursor)
                                return iife189_(_pat_let358_0)
                            return iife184_(AhpSkeleton.Option_Some(iife185_(d_232_at_)))
                        return iife183_(_pat_let357_0)
                    return (iife182_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CInputCompleted:
                d_237_reqId_ = source0_.requestId
                d_238_resp_ = source0_.response
                d_239_answers_ = source0_.answers
                if (((s).activeTurn).is_Some) and (default__.hasOpenReq((((s).activeTurn).value).parts, d_237_reqId_)):
                    d_240_at_ = ((s).activeTurn).value
                    def iife190_(_pat_let361_0):
                        def iife191_(d_242_dt__update__tmp_h46_):
                            def iife193_(_pat_let363_0):
                                def iife194_(d_243_dt__update__tmp_h47_):
                                    def iife195_(_pat_let364_0):
                                        def iife196_(d_244_dt__update_hparts_h13_):
                                            return Turn_Turn((d_243_dt__update__tmp_h47_).turnId, (d_243_dt__update__tmp_h47_).message, d_244_dt__update_hparts_h13_, (d_243_dt__update__tmp_h47_).state, (d_243_dt__update__tmp_h47_).usage, (d_243_dt__update__tmp_h47_).error)
                                        return iife196_(_pat_let364_0)
                                    return iife195_(default__.completeAnswerPart((d_240_at_).parts, d_237_reqId_, d_238_resp_, d_239_answers_))
                                return iife194_(_pat_let363_0)
                            def iife192_(_pat_let362_0):
                                def iife197_(d_245_dt__update_hactiveTurn_h17_):
                                    return ChatState_ChatState((d_242_dt__update__tmp_h46_).turns, (d_242_dt__update__tmp_h46_).title, (d_242_dt__update__tmp_h46_).status, (d_242_dt__update__tmp_h46_).modifiedAt, (d_242_dt__update__tmp_h46_).draft, (d_242_dt__update__tmp_h46_).activity, d_245_dt__update_hactiveTurn_h17_, (d_242_dt__update__tmp_h46_).steeringMessage, (d_242_dt__update__tmp_h46_).queuedMessages, (d_242_dt__update__tmp_h46_).nextCursor)
                                return iife197_(_pat_let362_0)
                            return iife192_(AhpSkeleton.Option_Some(iife193_(d_240_at_)))
                        return iife191_(_pat_let361_0)
                    d_241_next_ = iife190_(s)
                    def iife198_(_pat_let365_0):
                        def iife199_(d_246_dt__update__tmp_h48_):
                            def iife200_(_pat_let366_0):
                                def iife201_(d_247_dt__update_hstatus_h9_):
                                    return ChatState_ChatState((d_246_dt__update__tmp_h48_).turns, (d_246_dt__update__tmp_h48_).title, d_247_dt__update_hstatus_h9_, (d_246_dt__update__tmp_h48_).modifiedAt, (d_246_dt__update__tmp_h48_).draft, (d_246_dt__update__tmp_h48_).activity, (d_246_dt__update__tmp_h48_).activeTurn, (d_246_dt__update__tmp_h48_).steeringMessage, (d_246_dt__update__tmp_h48_).queuedMessages, (d_246_dt__update__tmp_h48_).nextCursor)
                                return iife201_(_pat_let366_0)
                            return iife200_(default__.summaryStatus(d_241_next_, False))
                        return iife199_(_pat_let365_0)
                    return (iife198_(d_241_next_), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_CPendingMessageSet:
                d_248_kind_ = source0_.kind
                d_249_id_ = source0_.id_
                d_250_msg_ = source0_.message
                if (d_248_kind_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "steering"))):
                    def iife202_(_pat_let367_0):
                        def iife203_(d_251_dt__update__tmp_h49_):
                            def iife204_(_pat_let368_0):
                                def iife205_(d_252_dt__update_hsteeringMessage_h1_):
                                    return ChatState_ChatState((d_251_dt__update__tmp_h49_).turns, (d_251_dt__update__tmp_h49_).title, (d_251_dt__update__tmp_h49_).status, (d_251_dt__update__tmp_h49_).modifiedAt, (d_251_dt__update__tmp_h49_).draft, (d_251_dt__update__tmp_h49_).activity, (d_251_dt__update__tmp_h49_).activeTurn, d_252_dt__update_hsteeringMessage_h1_, (d_251_dt__update__tmp_h49_).queuedMessages, (d_251_dt__update__tmp_h49_).nextCursor)
                                return iife205_(_pat_let368_0)
                            return iife204_(AhpSkeleton.Option_Some(QMsg_QMsg(d_249_id_, d_250_msg_)))
                        return iife203_(_pat_let367_0)
                    return (iife202_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    def iife206_(_pat_let369_0):
                        def iife207_(d_253_dt__update__tmp_h50_):
                            def iife208_(_pat_let370_0):
                                def iife209_(d_254_dt__update_hqueuedMessages_h2_):
                                    return ChatState_ChatState((d_253_dt__update__tmp_h50_).turns, (d_253_dt__update__tmp_h50_).title, (d_253_dt__update__tmp_h50_).status, (d_253_dt__update__tmp_h50_).modifiedAt, (d_253_dt__update__tmp_h50_).draft, (d_253_dt__update__tmp_h50_).activity, (d_253_dt__update__tmp_h50_).activeTurn, (d_253_dt__update__tmp_h50_).steeringMessage, d_254_dt__update_hqueuedMessages_h2_, (d_253_dt__update__tmp_h50_).nextCursor)
                                return iife209_(_pat_let370_0)
                            return iife208_(default__.upsertQ((pat_let_tv27_).queuedMessages, QMsg_QMsg(d_249_id_, d_250_msg_)))
                        return iife207_(_pat_let369_0)
                    return (iife206_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CPendingMessageRemoved:
                d_255_kind_ = source0_.kind
                d_256_id_ = source0_.id_
                if (d_255_kind_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "steering"))):
                    if (((s).steeringMessage).is_Some) and (((((s).steeringMessage).value).id_) == (d_256_id_)):
                        def iife210_(_pat_let371_0):
                            def iife211_(d_257_dt__update__tmp_h51_):
                                def iife212_(_pat_let372_0):
                                    def iife213_(d_258_dt__update_hsteeringMessage_h2_):
                                        return ChatState_ChatState((d_257_dt__update__tmp_h51_).turns, (d_257_dt__update__tmp_h51_).title, (d_257_dt__update__tmp_h51_).status, (d_257_dt__update__tmp_h51_).modifiedAt, (d_257_dt__update__tmp_h51_).draft, (d_257_dt__update__tmp_h51_).activity, (d_257_dt__update__tmp_h51_).activeTurn, d_258_dt__update_hsteeringMessage_h2_, (d_257_dt__update__tmp_h51_).queuedMessages, (d_257_dt__update__tmp_h51_).nextCursor)
                                    return iife213_(_pat_let372_0)
                                return iife212_(AhpSkeleton.Option_None())
                            return iife211_(_pat_let371_0)
                        return (iife210_(s), AhpSkeleton.ReduceOutcome_Applied())
                    elif True:
                        return (s, AhpSkeleton.ReduceOutcome_NoOp())
                elif True:
                    def iife214_(_pat_let373_0):
                        def iife215_(d_259_dt__update__tmp_h52_):
                            def iife216_(_pat_let374_0):
                                def iife217_(d_260_dt__update_hqueuedMessages_h3_):
                                    return ChatState_ChatState((d_259_dt__update__tmp_h52_).turns, (d_259_dt__update__tmp_h52_).title, (d_259_dt__update__tmp_h52_).status, (d_259_dt__update__tmp_h52_).modifiedAt, (d_259_dt__update__tmp_h52_).draft, (d_259_dt__update__tmp_h52_).activity, (d_259_dt__update__tmp_h52_).activeTurn, (d_259_dt__update__tmp_h52_).steeringMessage, d_260_dt__update_hqueuedMessages_h3_, (d_259_dt__update__tmp_h52_).nextCursor)
                                return iife217_(_pat_let374_0)
                            return iife216_(default__.removeQ((pat_let_tv28_).queuedMessages, d_256_id_))
                        return iife215_(_pat_let373_0)
                    return (iife214_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            return (s, AhpSkeleton.ReduceOutcome_NoOp())

    @staticmethod
    def apply1(s, a):
        return (default__.ApplyToChat(s, a, 9999))[0]

    @staticmethod
    def fold(s, acts):
        return ConfluxContract.default__.Fold(default__.apply1, s, acts)

    @staticmethod
    def C0():
        return ChatState_ChatState(_dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Chat")), 1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), _dafny.SeqWithoutIsStrInference([]), AhpSkeleton.Option_None())

    @staticmethod
    def T0(id_):
        return Turn_Turn(id_, ConfluxCodec.Json_JNull(), _dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "in-progress")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())

    @staticmethod
    def TDone(id_):
        return Turn_Turn(id_, ConfluxCodec.Json_JNull(), _dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "complete")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())

    @staticmethod
    def TC0(id_, intent):
        return ToolCall_ToolCall(id_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bash")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run Command")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "streaming")), intent, AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())

    @staticmethod
    def activeParts(s):
        if ((s).activeTurn).is_Some:
            return (((s).activeTurn).value).parts
        elif True:
            return _dafny.SeqWithoutIsStrInference([])

    @staticmethod
    def firstActiveToolCall(s):
        d_0_ps_ = default__.activeParts(s)
        if ((len(d_0_ps_)) > (0)) and (((d_0_ps_)[0]).is_PToolCall):
            return AhpSkeleton.Option_Some(((d_0_ps_)[0]).tc)
        elif True:
            return AhpSkeleton.Option_None()

    @staticmethod
    def firstStoredToolCall(s):
        if (len((s).turns)) == (0):
            return AhpSkeleton.Option_None()
        elif True:
            d_0_ps_ = (((s).turns)[0]).parts
            if ((len(d_0_ps_)) > (0)) and (((d_0_ps_)[0]).is_PToolCall):
                return AhpSkeleton.Option_Some(((d_0_ps_)[0]).tc)
            elif True:
                return AhpSkeleton.Option_None()

    @staticmethod
    def runScalarTurns():
        pass_: int = int(0)
        pass_ = 0
        d_0_s_: ChatState
        d_0_s_ = default__.C0()
        def iife0_(_pat_let375_0):
            def iife1_(d_1_dt__update__tmp_h0_):
                def iife2_(_pat_let376_0):
                    def iife3_(d_2_dt__update_hdraft_h0_):
                        return ChatState_ChatState((d_1_dt__update__tmp_h0_).turns, (d_1_dt__update__tmp_h0_).title, (d_1_dt__update__tmp_h0_).status, (d_1_dt__update__tmp_h0_).modifiedAt, d_2_dt__update_hdraft_h0_, (d_1_dt__update__tmp_h0_).activity, (d_1_dt__update__tmp_h0_).activeTurn, (d_1_dt__update__tmp_h0_).steeringMessage, (d_1_dt__update__tmp_h0_).queuedMessages, (d_1_dt__update__tmp_h0_).nextCursor)
                    return iife3_(_pat_let376_0)
                return iife2_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))))
            return iife1_(_pat_let375_0)
        if ((default__.apply1(iife0_(d_0_s_), ChatAction_CDraftChanged(AhpSkeleton.Option_None()))).draft) == (AhpSkeleton.Option_None()):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_s_, ChatAction_CDraftChanged(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hi")))))).draft) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hi")))):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_s_, ChatAction_CActivityChanged(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Running terminal command")))))).activity) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Running terminal command")))):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_s_, ChatAction_CActivityChanged(AhpSkeleton.Option_None()))).activity) == (AhpSkeleton.Option_None()):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_s_, ChatAction_CTurnStarted(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turn-1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Hello"))), AhpSkeleton.Option_None()))).activeTurn) == (AhpSkeleton.Option_Some(Turn_Turn(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turn-1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Hello"))), _dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "in-progress")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()))):
            pass_ = (pass_) + (1)
        d_3_act1_: ChatState
        d_4_dt__update__tmp_h1_ = d_0_s_
        d_5_dt__update_hactiveTurn_h0_ = AhpSkeleton.Option_Some(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))))
        d_3_act1_ = ChatState_ChatState((d_4_dt__update__tmp_h1_).turns, (d_4_dt__update__tmp_h1_).title, (d_4_dt__update__tmp_h1_).status, (d_4_dt__update__tmp_h1_).modifiedAt, (d_4_dt__update__tmp_h1_).draft, (d_4_dt__update__tmp_h1_).activity, d_5_dt__update_hactiveTurn_h0_, (d_4_dt__update__tmp_h1_).steeringMessage, (d_4_dt__update__tmp_h1_).queuedMessages, (d_4_dt__update__tmp_h1_).nextCursor)
        if (default__.apply1(d_3_act1_, ChatAction_CTurnComplete(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrong-turn"))))) == (d_3_act1_):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_3_act1_, ChatAction_CUsage(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrong-turn")), ConfluxCodec.Json_JNull()))) == (d_3_act1_):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_3_act1_, ChatAction_CToolCallStart(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrong-turn")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bash")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run Command")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()))) == (d_3_act1_):
            pass_ = (pass_) + (1)
        def iife4_(_pat_let377_0):
            def iife5_(d_6_dt__update__tmp_h2_):
                def iife6_(_pat_let378_0):
                    def iife7_(d_7_dt__update_hactiveTurn_h1_):
                        return ChatState_ChatState((d_6_dt__update__tmp_h2_).turns, (d_6_dt__update__tmp_h2_).title, (d_6_dt__update__tmp_h2_).status, (d_6_dt__update__tmp_h2_).modifiedAt, (d_6_dt__update__tmp_h2_).draft, (d_6_dt__update__tmp_h2_).activity, d_7_dt__update_hactiveTurn_h1_, (d_6_dt__update__tmp_h2_).steeringMessage, (d_6_dt__update__tmp_h2_).queuedMessages, (d_6_dt__update__tmp_h2_).nextCursor)
                    return iife7_(_pat_let378_0)
                return iife6_(AhpSkeleton.Option_Some(Turn_Turn(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "T"))), _dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "in-progress")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())))
            return iife5_(_pat_let377_0)
        def iife8_(_pat_let379_0):
            def iife9_(d_8_dt__update__tmp_h3_):
                def iife10_(_pat_let380_0):
                    def iife11_(d_9_dt__update_hturns_h0_):
                        def iife12_(_pat_let381_0):
                            def iife13_(d_10_dt__update_hactiveTurn_h2_):
                                return ChatState_ChatState(d_9_dt__update_hturns_h0_, (d_8_dt__update__tmp_h3_).title, (d_8_dt__update__tmp_h3_).status, (d_8_dt__update__tmp_h3_).modifiedAt, (d_8_dt__update__tmp_h3_).draft, (d_8_dt__update__tmp_h3_).activity, d_10_dt__update_hactiveTurn_h2_, (d_8_dt__update__tmp_h3_).steeringMessage, (d_8_dt__update__tmp_h3_).queuedMessages, (d_8_dt__update__tmp_h3_).nextCursor)
                            return iife13_(_pat_let381_0)
                        return iife12_(AhpSkeleton.Option_None())
                    return iife11_(_pat_let380_0)
                return iife10_(_dafny.SeqWithoutIsStrInference([Turn_Turn(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "T"))), _dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "complete")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())]))
            return iife9_(_pat_let379_0)
        if (default__.apply1(iife4_(d_0_s_), ChatAction_CTurnComplete(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))))) == (iife8_(d_0_s_)):
            pass_ = (pass_) + (1)
        def iife14_(_pat_let382_0):
            def iife15_(d_11_dt__update__tmp_h4_):
                def iife16_(_pat_let383_0):
                    def iife17_(d_12_dt__update_hactiveTurn_h3_):
                        return ChatState_ChatState((d_11_dt__update__tmp_h4_).turns, (d_11_dt__update__tmp_h4_).title, (d_11_dt__update__tmp_h4_).status, (d_11_dt__update__tmp_h4_).modifiedAt, (d_11_dt__update__tmp_h4_).draft, (d_11_dt__update__tmp_h4_).activity, d_12_dt__update_hactiveTurn_h3_, (d_11_dt__update__tmp_h4_).steeringMessage, (d_11_dt__update__tmp_h4_).queuedMessages, (d_11_dt__update__tmp_h4_).nextCursor)
                    return iife17_(_pat_let383_0)
                return iife16_(AhpSkeleton.Option_Some(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))
            return iife15_(_pat_let382_0)
        if ((default__.apply1(iife14_(d_0_s_), ChatAction_CTurnCancelled(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))))).turns) == (_dafny.SeqWithoutIsStrInference([Turn_Turn(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), ConfluxCodec.Json_JNull(), _dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cancelled")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())])):
            pass_ = (pass_) + (1)
        def iife18_(_pat_let384_0):
            def iife19_(d_13_dt__update__tmp_h5_):
                def iife20_(_pat_let385_0):
                    def iife21_(d_14_dt__update_husage_h0_):
                        return Turn_Turn((d_13_dt__update__tmp_h5_).turnId, (d_13_dt__update__tmp_h5_).message, (d_13_dt__update__tmp_h5_).parts, (d_13_dt__update__tmp_h5_).state, d_14_dt__update_husage_h0_, (d_13_dt__update__tmp_h5_).error)
                    return iife21_(_pat_let385_0)
                return iife20_(AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")))))
            return iife19_(_pat_let384_0)
        if ((default__.apply1(d_3_act1_, ChatAction_CUsage(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")))))).activeTurn) == (AhpSkeleton.Option_Some(iife18_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_3_act1_, ChatAction_CError(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "boom")))))).turns) == (_dafny.SeqWithoutIsStrInference([Turn_Turn(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), ConfluxCodec.Json_JNull(), _dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "error")), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "boom")))))])):
            pass_ = (pass_) + (1)
        d_15_withMd_: ChatState
        d_15_withMd_ = default__.apply1(d_3_act1_, ChatAction_CResponsePart(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), Part_PMarkdown(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "md-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))))
        def iife22_(_pat_let386_0):
            def iife23_(d_16_dt__update__tmp_h6_):
                def iife24_(_pat_let387_0):
                    def iife25_(d_17_dt__update_hparts_h0_):
                        return Turn_Turn((d_16_dt__update__tmp_h6_).turnId, (d_16_dt__update__tmp_h6_).message, d_17_dt__update_hparts_h0_, (d_16_dt__update__tmp_h6_).state, (d_16_dt__update__tmp_h6_).usage, (d_16_dt__update__tmp_h6_).error)
                    return iife25_(_pat_let387_0)
                return iife24_(_dafny.SeqWithoutIsStrInference([Part_PMarkdown(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "md-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))]))
            return iife23_(_pat_let386_0)
        if ((d_15_withMd_).activeTurn) == (AhpSkeleton.Option_Some(iife22_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))):
            pass_ = (pass_) + (1)
        def iife26_(_pat_let388_0):
            def iife27_(d_18_dt__update__tmp_h7_):
                def iife28_(_pat_let389_0):
                    def iife29_(d_19_dt__update_hparts_h1_):
                        return Turn_Turn((d_18_dt__update__tmp_h7_).turnId, (d_18_dt__update__tmp_h7_).message, d_19_dt__update_hparts_h1_, (d_18_dt__update__tmp_h7_).state, (d_18_dt__update__tmp_h7_).usage, (d_18_dt__update__tmp_h7_).error)
                    return iife29_(_pat_let389_0)
                return iife28_(_dafny.SeqWithoutIsStrInference([Part_PMarkdown(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "md-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Hello from chat")))]))
            return iife27_(_pat_let388_0)
        if ((default__.apply1(d_15_withMd_, ChatAction_CDelta(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "md-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Hello from chat"))))).activeTurn) == (AhpSkeleton.Option_Some(iife26_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))):
            pass_ = (pass_) + (1)
        def iife30_(_pat_let390_0):
            def iife31_(d_20_dt__update__tmp_h8_):
                def iife33_(_pat_let392_0):
                    def iife34_(d_21_dt__update__tmp_h9_):
                        def iife35_(_pat_let393_0):
                            def iife36_(d_22_dt__update_hparts_h2_):
                                return Turn_Turn((d_21_dt__update__tmp_h9_).turnId, (d_21_dt__update__tmp_h9_).message, d_22_dt__update_hparts_h2_, (d_21_dt__update__tmp_h9_).state, (d_21_dt__update__tmp_h9_).usage, (d_21_dt__update__tmp_h9_).error)
                            return iife36_(_pat_let393_0)
                        return iife35_(_dafny.SeqWithoutIsStrInference([Part_PReasoning(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "th")))]))
                    return iife34_(_pat_let392_0)
                def iife32_(_pat_let391_0):
                    def iife37_(d_23_dt__update_hactiveTurn_h4_):
                        return ChatState_ChatState((d_20_dt__update__tmp_h8_).turns, (d_20_dt__update__tmp_h8_).title, (d_20_dt__update__tmp_h8_).status, (d_20_dt__update__tmp_h8_).modifiedAt, (d_20_dt__update__tmp_h8_).draft, (d_20_dt__update__tmp_h8_).activity, d_23_dt__update_hactiveTurn_h4_, (d_20_dt__update__tmp_h8_).steeringMessage, (d_20_dt__update__tmp_h8_).queuedMessages, (d_20_dt__update__tmp_h8_).nextCursor)
                    return iife37_(_pat_let391_0)
                return iife32_(AhpSkeleton.Option_Some(iife33_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))))))
            return iife31_(_pat_let390_0)
        def iife38_(_pat_let394_0):
            def iife39_(d_24_dt__update__tmp_h10_):
                def iife40_(_pat_let395_0):
                    def iife41_(d_25_dt__update_hparts_h3_):
                        return Turn_Turn((d_24_dt__update__tmp_h10_).turnId, (d_24_dt__update__tmp_h10_).message, d_25_dt__update_hparts_h3_, (d_24_dt__update__tmp_h10_).state, (d_24_dt__update__tmp_h10_).usage, (d_24_dt__update__tmp_h10_).error)
                    return iife41_(_pat_let395_0)
                return iife40_(_dafny.SeqWithoutIsStrInference([Part_PReasoning(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "th")))]))
            return iife39_(_pat_let394_0)
        if ((default__.apply1(iife30_(d_3_act1_), ChatAction_CDelta(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "inking"))))).activeTurn) == (AhpSkeleton.Option_Some(iife38_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))):
            pass_ = (pass_) + (1)
        return pass_

    @staticmethod
    def runToolCalls():
        pass_: int = int(0)
        pass_ = 0
        d_0_s_: ChatState
        d_0_s_ = default__.C0()
        d_1_act1_: ChatState
        d_2_dt__update__tmp_h0_ = d_0_s_
        d_3_dt__update_hactiveTurn_h0_ = AhpSkeleton.Option_Some(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))))
        d_1_act1_ = ChatState_ChatState((d_2_dt__update__tmp_h0_).turns, (d_2_dt__update__tmp_h0_).title, (d_2_dt__update__tmp_h0_).status, (d_2_dt__update__tmp_h0_).modifiedAt, (d_2_dt__update__tmp_h0_).draft, (d_2_dt__update__tmp_h0_).activity, d_3_dt__update_hactiveTurn_h0_, (d_2_dt__update__tmp_h0_).steeringMessage, (d_2_dt__update__tmp_h0_).queuedMessages, (d_2_dt__update__tmp_h0_).nextCursor)
        d_4_started_: ChatState
        d_4_started_ = default__.apply1(d_1_act1_, ChatAction_CToolCallStart(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bash")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run Command")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "List the files in the current directory"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()))
        def iife0_(_pat_let396_0):
            def iife1_(d_5_dt__update__tmp_h1_):
                def iife2_(_pat_let397_0):
                    def iife3_(d_6_dt__update_hparts_h0_):
                        return Turn_Turn((d_5_dt__update__tmp_h1_).turnId, (d_5_dt__update__tmp_h1_).message, d_6_dt__update_hparts_h0_, (d_5_dt__update__tmp_h1_).state, (d_5_dt__update__tmp_h1_).usage, (d_5_dt__update__tmp_h1_).error)
                    return iife3_(_pat_let397_0)
                return iife2_(_dafny.SeqWithoutIsStrInference([Part_PToolCall(default__.TC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "List the files in the current directory")))))]))
            return iife1_(_pat_let396_0)
        if ((d_4_started_).activeTurn) == (AhpSkeleton.Option_Some(iife0_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))):
            pass_ = (pass_) + (1)
        d_7_afterDelta_: ChatState
        d_7_afterDelta_ = default__.apply1(d_4_started_, ChatAction_CToolCallDelta(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ls -la")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Listing files"))), AhpSkeleton.Option_None()))
        def iife4_(_pat_let398_0):
            def iife5_(d_8_dt__update__tmp_h2_):
                def iife6_(_pat_let399_0):
                    def iife7_(d_9_dt__update_hinvocationMessage_h0_):
                        def iife8_(_pat_let400_0):
                            def iife9_(d_10_dt__update_hpartialInput_h0_):
                                return ToolCall_ToolCall((d_8_dt__update__tmp_h2_).toolCallId, (d_8_dt__update__tmp_h2_).toolName, (d_8_dt__update__tmp_h2_).displayName, (d_8_dt__update__tmp_h2_).status, (d_8_dt__update__tmp_h2_).intention, (d_8_dt__update__tmp_h2_).contributor, (d_8_dt__update__tmp_h2_).meta, d_9_dt__update_hinvocationMessage_h0_, (d_8_dt__update__tmp_h2_).toolInput, (d_8_dt__update__tmp_h2_).confirmationTitle, (d_8_dt__update__tmp_h2_).riskAssessment, (d_8_dt__update__tmp_h2_).edits, (d_8_dt__update__tmp_h2_).editable, (d_8_dt__update__tmp_h2_).options, (d_8_dt__update__tmp_h2_).confirmed, (d_8_dt__update__tmp_h2_).success, (d_8_dt__update__tmp_h2_).pastTenseMessage, (d_8_dt__update__tmp_h2_).reason, (d_8_dt__update__tmp_h2_).reasonMessage, (d_8_dt__update__tmp_h2_).userSuggestion, (d_8_dt__update__tmp_h2_).selectedOption, (d_8_dt__update__tmp_h2_).content, (d_8_dt__update__tmp_h2_).structuredContent, (d_8_dt__update__tmp_h2_).error, (d_8_dt__update__tmp_h2_).auth, d_10_dt__update_hpartialInput_h0_)
                            return iife9_(_pat_let400_0)
                        return iife8_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ls -la"))))
                    return iife7_(_pat_let399_0)
                return iife6_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Listing files")))
            return iife5_(_pat_let398_0)
        if (default__.activeParts(d_7_afterDelta_)) == (_dafny.SeqWithoutIsStrInference([Part_PToolCall(iife4_(default__.TC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "List the files in the current directory"))))))])):
            pass_ = (pass_) + (1)
        d_11_afterReady_: ChatState
        d_11_afterReady_ = default__.apply1(d_7_afterDelta_, ChatAction_CToolCallReady(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run: ls -la"))), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ls -la"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()))
        d_12_tcReady_: AhpSkeleton.Option
        d_12_tcReady_ = default__.firstActiveToolCall(d_11_afterReady_)
        if (((((d_12_tcReady_).is_Some) and ((((d_12_tcReady_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-confirmation"))))) and ((((d_12_tcReady_).value).invocationMessage) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run: ls -la"))))) and ((((d_12_tcReady_).value).toolInput) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ls -la")))))) and ((((d_12_tcReady_).value).partialInput) == (AhpSkeleton.Option_None())):
            pass_ = (pass_) + (1)
        d_13_afterConf_: ChatState
        d_13_afterConf_ = default__.apply1(d_11_afterReady_, ChatAction_CToolCallConfirmed(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "user-action"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()))
        d_14_tcConf_: AhpSkeleton.Option
        d_14_tcConf_ = default__.firstActiveToolCall(d_13_afterConf_)
        if (((d_14_tcConf_).is_Some) and ((((d_14_tcConf_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))))) and ((((d_14_tcConf_).value).confirmed) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "user-action"))))):
            pass_ = (pass_) + (1)
        d_15_afterComp_: ChatState
        d_15_afterComp_ = default__.apply1(d_13_afterConf_, ChatAction_CToolCallComplete(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Ran command"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), False, AhpSkeleton.Option_None()))
        d_16_tcComp_: AhpSkeleton.Option
        d_16_tcComp_ = default__.firstActiveToolCall(d_15_afterComp_)
        if ((((d_16_tcComp_).is_Some) and ((((d_16_tcComp_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "completed"))))) and ((((d_16_tcComp_).value).success) == (AhpSkeleton.Option_Some(True)))) and ((((d_16_tcComp_).value).pastTenseMessage) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Ran command"))))):
            pass_ = (pass_) + (1)
        d_17_life_: ChatState
        d_17_life_ = default__.fold(d_1_act1_, _dafny.SeqWithoutIsStrInference([ChatAction_CToolCallStart(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bash")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run Command")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "List the files in the current directory"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChatAction_CToolCallDelta(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ls -la")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Listing files"))), AhpSkeleton.Option_None()), ChatAction_CToolCallReady(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run: ls -la"))), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ls -la"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChatAction_CToolCallConfirmed(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "user-action"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChatAction_CToolCallComplete(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Ran command"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), False, AhpSkeleton.Option_None())]))
        d_18_lifeTc_: AhpSkeleton.Option
        d_18_lifeTc_ = default__.firstActiveToolCall(d_17_life_)
        if (((((d_18_lifeTc_).is_Some) and ((((d_18_lifeTc_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "completed"))))) and ((((d_18_lifeTc_).value).confirmed) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "user-action")))))) and ((((d_18_lifeTc_).value).success) == (AhpSkeleton.Option_Some(True)))) and ((((d_18_lifeTc_).value).toolInput) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ls -la"))))):
            pass_ = (pass_) + (1)
        d_19_auto_: ChatState
        d_19_auto_ = default__.fold(d_1_act1_, _dafny.SeqWithoutIsStrInference([ChatAction_CToolCallStart(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bash")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run Command")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChatAction_CToolCallReady(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "not-needed"))), AhpSkeleton.Option_None())]))
        d_20_autoTc_: AhpSkeleton.Option
        d_20_autoTc_ = default__.firstActiveToolCall(d_19_auto_)
        if ((((d_20_autoTc_).is_Some) and ((((d_20_autoTc_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))))) and ((((d_20_autoTc_).value).confirmed) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "not-needed")))))) and ((((d_20_autoTc_).value).invocationMessage) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run")))):
            pass_ = (pass_) + (1)
        d_21_pendingRisk_: ConfluxCodec.Json
        d_21_pendingRisk_ = ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "status")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "loading")))}))
        d_22_pendingEdits_: ConfluxCodec.Json
        d_22_pendingEdits_ = ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "items")): ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([]))}))
        d_23_pendingOption_: ConfluxCodec.Json
        d_23_pendingOption_ = ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "id")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "approve")))}))
        pat_let_tv0_ = d_23_pendingOption_
        pat_let_tv1_ = d_22_pendingEdits_
        pat_let_tv2_ = d_21_pendingRisk_
        d_24_pendState_: ChatState
        d_25_dt__update__tmp_h3_ = d_1_act1_
        d_26_dt__update_hstatus_h0_ = 24
        def iife10_(_pat_let401_0):
            def iife11_(d_28_dt__update__tmp_h4_):
                def iife13_(_pat_let403_0):
                    def iife14_(d_29_dt__update__tmp_h5_):
                        def iife15_(_pat_let404_0):
                            def iife16_(d_30_dt__update_hoptions_h0_):
                                def iife17_(_pat_let405_0):
                                    def iife18_(d_31_dt__update_heditable_h0_):
                                        def iife19_(_pat_let406_0):
                                            def iife20_(d_32_dt__update_hedits_h0_):
                                                def iife21_(_pat_let407_0):
                                                    def iife22_(d_33_dt__update_hriskAssessment_h0_):
                                                        def iife23_(_pat_let408_0):
                                                            def iife24_(d_34_dt__update_hconfirmationTitle_h0_):
                                                                def iife25_(_pat_let409_0):
                                                                    def iife26_(d_35_dt__update_htoolInput_h0_):
                                                                        def iife27_(_pat_let410_0):
                                                                            def iife28_(d_36_dt__update_hstatus_h1_):
                                                                                return ToolCall_ToolCall((d_29_dt__update__tmp_h5_).toolCallId, (d_29_dt__update__tmp_h5_).toolName, (d_29_dt__update__tmp_h5_).displayName, d_36_dt__update_hstatus_h1_, (d_29_dt__update__tmp_h5_).intention, (d_29_dt__update__tmp_h5_).contributor, (d_29_dt__update__tmp_h5_).meta, (d_29_dt__update__tmp_h5_).invocationMessage, d_35_dt__update_htoolInput_h0_, d_34_dt__update_hconfirmationTitle_h0_, d_33_dt__update_hriskAssessment_h0_, d_32_dt__update_hedits_h0_, d_31_dt__update_heditable_h0_, d_30_dt__update_hoptions_h0_, (d_29_dt__update__tmp_h5_).confirmed, (d_29_dt__update__tmp_h5_).success, (d_29_dt__update__tmp_h5_).pastTenseMessage, (d_29_dt__update__tmp_h5_).reason, (d_29_dt__update__tmp_h5_).reasonMessage, (d_29_dt__update__tmp_h5_).userSuggestion, (d_29_dt__update__tmp_h5_).selectedOption, (d_29_dt__update__tmp_h5_).content, (d_29_dt__update__tmp_h5_).structuredContent, (d_29_dt__update__tmp_h5_).error, (d_29_dt__update__tmp_h5_).auth, (d_29_dt__update__tmp_h5_).partialInput)
                                                                            return iife28_(_pat_let410_0)
                                                                        return iife27_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-confirmation")))
                                                                    return iife26_(_pat_let409_0)
                                                                return iife25_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rm -rf /tmp/test"))))
                                                            return iife24_(_pat_let408_0)
                                                        return iife23_(AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run in terminal")))))
                                                    return iife22_(_pat_let407_0)
                                                return iife21_(AhpSkeleton.Option_Some(pat_let_tv2_))
                                            return iife20_(_pat_let406_0)
                                        return iife19_(AhpSkeleton.Option_Some(pat_let_tv1_))
                                    return iife18_(_pat_let405_0)
                                return iife17_(AhpSkeleton.Option_Some(True))
                            return iife16_(_pat_let404_0)
                        return iife15_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([pat_let_tv0_])))
                    return iife14_(_pat_let403_0)
                def iife12_(_pat_let402_0):
                    def iife29_(d_37_dt__update_hparts_h1_):
                        return Turn_Turn((d_28_dt__update__tmp_h4_).turnId, (d_28_dt__update__tmp_h4_).message, d_37_dt__update_hparts_h1_, (d_28_dt__update__tmp_h4_).state, (d_28_dt__update__tmp_h4_).usage, (d_28_dt__update__tmp_h4_).error)
                    return iife29_(_pat_let402_0)
                return iife12_(_dafny.SeqWithoutIsStrInference([Part_PToolCall(iife13_(default__.TC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_None())))]))
            return iife11_(_pat_let401_0)
        d_27_dt__update_hactiveTurn_h1_ = AhpSkeleton.Option_Some(iife10_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))
        d_24_pendState_ = ChatState_ChatState((d_25_dt__update__tmp_h3_).turns, (d_25_dt__update__tmp_h3_).title, d_26_dt__update_hstatus_h0_, (d_25_dt__update__tmp_h3_).modifiedAt, (d_25_dt__update__tmp_h3_).draft, (d_25_dt__update__tmp_h3_).activity, d_27_dt__update_hactiveTurn_h1_, (d_25_dt__update__tmp_h3_).steeringMessage, (d_25_dt__update__tmp_h3_).queuedMessages, (d_25_dt__update__tmp_h3_).nextCursor)
        d_38_refreshed_: AhpSkeleton.Option
        d_38_refreshed_ = default__.firstActiveToolCall(default__.apply1(d_24_pendState_, ChatAction_CToolCallReady(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run again"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())))
        if (((((((((d_38_refreshed_).is_Some) and ((((d_38_refreshed_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pending-confirmation"))))) and ((((d_38_refreshed_).value).invocationMessage) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run again"))))) and ((((d_38_refreshed_).value).toolInput) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rm -rf /tmp/test")))))) and ((((d_38_refreshed_).value).confirmationTitle) == (AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run in terminal"))))))) and ((((d_38_refreshed_).value).riskAssessment) == (AhpSkeleton.Option_Some(d_21_pendingRisk_)))) and ((((d_38_refreshed_).value).edits) == (AhpSkeleton.Option_Some(d_22_pendingEdits_)))) and ((((d_38_refreshed_).value).editable) == (AhpSkeleton.Option_Some(True)))) and ((((d_38_refreshed_).value).options) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference([d_23_pendingOption_])))):
            pass_ = (pass_) + (1)
        d_39_forced_: ChatState
        d_39_forced_ = default__.apply1(d_4_started_, ChatAction_CTurnComplete(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))))
        d_40_fTc_: AhpSkeleton.Option
        d_40_fTc_ = default__.firstStoredToolCall(d_39_forced_)
        if (((((((d_39_forced_).activeTurn) == (AhpSkeleton.Option_None())) and ((len((d_39_forced_).turns)) > (0))) and (((((d_39_forced_).turns)[0]).state) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "complete"))))) and ((d_40_fTc_).is_Some)) and ((((d_40_fTc_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cancelled"))))) and ((((d_40_fTc_).value).reason) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skipped"))))):
            pass_ = (pass_) + (1)
        d_41_ignoredComplete_: ChatState
        d_41_ignoredComplete_ = default__.apply1(d_4_started_, ChatAction_CToolCallComplete(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), False, AhpSkeleton.Option_None()))
        if (default__.activeParts(d_41_ignoredComplete_)) == (default__.activeParts(d_4_started_)):
            pass_ = (pass_) + (1)
        d_42_authBase_: ChatState
        d_43_dt__update__tmp_h6_ = d_1_act1_
        d_44_dt__update_hstatus_h2_ = 8
        def iife30_(_pat_let411_0):
            def iife31_(d_46_dt__update__tmp_h7_):
                def iife33_(_pat_let413_0):
                    def iife34_(d_47_dt__update__tmp_h8_):
                        def iife35_(_pat_let414_0):
                            def iife36_(d_48_dt__update_hcontent_h0_):
                                def iife37_(_pat_let415_0):
                                    def iife38_(d_49_dt__update_hconfirmed_h0_):
                                        def iife39_(_pat_let416_0):
                                            def iife40_(d_50_dt__update_htoolInput_h1_):
                                                def iife41_(_pat_let417_0):
                                                    def iife42_(d_51_dt__update_hcontributor_h0_):
                                                        def iife43_(_pat_let418_0):
                                                            def iife44_(d_52_dt__update_hstatus_h3_):
                                                                return ToolCall_ToolCall((d_47_dt__update__tmp_h8_).toolCallId, (d_47_dt__update__tmp_h8_).toolName, (d_47_dt__update__tmp_h8_).displayName, d_52_dt__update_hstatus_h3_, (d_47_dt__update__tmp_h8_).intention, d_51_dt__update_hcontributor_h0_, (d_47_dt__update__tmp_h8_).meta, (d_47_dt__update__tmp_h8_).invocationMessage, d_50_dt__update_htoolInput_h1_, (d_47_dt__update__tmp_h8_).confirmationTitle, (d_47_dt__update__tmp_h8_).riskAssessment, (d_47_dt__update__tmp_h8_).edits, (d_47_dt__update__tmp_h8_).editable, (d_47_dt__update__tmp_h8_).options, d_49_dt__update_hconfirmed_h0_, (d_47_dt__update__tmp_h8_).success, (d_47_dt__update__tmp_h8_).pastTenseMessage, (d_47_dt__update__tmp_h8_).reason, (d_47_dt__update__tmp_h8_).reasonMessage, (d_47_dt__update__tmp_h8_).userSuggestion, (d_47_dt__update__tmp_h8_).selectedOption, d_48_dt__update_hcontent_h0_, (d_47_dt__update__tmp_h8_).structuredContent, (d_47_dt__update__tmp_h8_).error, (d_47_dt__update__tmp_h8_).auth, (d_47_dt__update__tmp_h8_).partialInput)
                                                            return iife44_(_pat_let418_0)
                                                        return iife43_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running")))
                                                    return iife42_(_pat_let417_0)
                                                return iife41_(AhpSkeleton.Option_Some(ToolCallContributor_ToolCallContributor(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcp")), ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcp")))})))))
                                            return iife40_(_pat_let416_0)
                                        return iife39_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "foo"))))
                                    return iife38_(_pat_let415_0)
                                return iife37_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "user-action"))))
                            return iife36_(_pat_let414_0)
                        return iife35_(AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "partial")))))
                    return iife34_(_pat_let413_0)
                def iife32_(_pat_let412_0):
                    def iife45_(d_53_dt__update_hparts_h2_):
                        return Turn_Turn((d_46_dt__update__tmp_h7_).turnId, (d_46_dt__update__tmp_h7_).message, d_53_dt__update_hparts_h2_, (d_46_dt__update__tmp_h7_).state, (d_46_dt__update__tmp_h7_).usage, (d_46_dt__update__tmp_h7_).error)
                    return iife45_(_pat_let412_0)
                return iife32_(_dafny.SeqWithoutIsStrInference([Part_PToolCall(iife33_(default__.TC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-auth")), AhpSkeleton.Option_None())))]))
            return iife31_(_pat_let411_0)
        d_45_dt__update_hactiveTurn_h2_ = AhpSkeleton.Option_Some(iife30_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))
        d_42_authBase_ = ChatState_ChatState((d_43_dt__update__tmp_h6_).turns, (d_43_dt__update__tmp_h6_).title, d_44_dt__update_hstatus_h2_, (d_43_dt__update__tmp_h6_).modifiedAt, (d_43_dt__update__tmp_h6_).draft, (d_43_dt__update__tmp_h6_).activity, d_45_dt__update_hactiveTurn_h2_, (d_43_dt__update__tmp_h6_).steeringMessage, (d_43_dt__update__tmp_h6_).queuedMessages, (d_43_dt__update__tmp_h6_).nextCursor)
        d_54_challenge_: ConfluxCodec.Json
        d_54_challenge_ = ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "reason")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "required")))}))
        d_55_authState_: ChatState
        d_55_authState_ = default__.apply1(d_42_authBase_, ChatAction_CToolCallAuthRequired(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-auth")), d_54_challenge_, AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auth-meta"))))))
        d_56_authTc_: AhpSkeleton.Option
        d_56_authTc_ = default__.firstActiveToolCall(d_55_authState_)
        if (((((((((d_55_authState_).status) == (24)) and ((d_56_authTc_).is_Some)) and ((((d_56_authTc_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auth-required"))))) and ((((d_56_authTc_).value).auth) == (AhpSkeleton.Option_Some(d_54_challenge_)))) and ((((d_56_authTc_).value).meta) == (AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auth-meta"))))))) and ((((d_56_authTc_).value).toolInput) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "foo")))))) and ((((d_56_authTc_).value).confirmed) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "user-action")))))) and ((((d_56_authTc_).value).content) == (AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "partial")))))):
            pass_ = (pass_) + (1)
        d_57_resolvedState_: ChatState
        d_57_resolvedState_ = default__.apply1(d_55_authState_, ChatAction_CToolCallAuthResolved(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-auth")), AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "resolved-meta"))))))
        d_58_resolvedTc_: AhpSkeleton.Option
        d_58_resolvedTc_ = default__.firstActiveToolCall(d_57_resolvedState_)
        if ((((((((((d_57_resolvedState_).status) == (8)) and ((d_58_resolvedTc_).is_Some)) and ((((d_58_resolvedTc_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))))) and ((((d_58_resolvedTc_).value).auth) == (AhpSkeleton.Option_None()))) and ((((d_58_resolvedTc_).value).contributor) == (AhpSkeleton.Option_Some(ToolCallContributor_ToolCallContributor(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcp")), ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcp")))}))))))) and ((((d_58_resolvedTc_).value).toolInput) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "foo")))))) and ((((d_58_resolvedTc_).value).confirmed) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "user-action")))))) and ((((d_58_resolvedTc_).value).content) == (AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "partial"))))))) and ((((d_58_resolvedTc_).value).meta) == (AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "resolved-meta")))))):
            pass_ = (pass_) + (1)
        d_59_clientBase_: ChatState
        d_60_dt__update__tmp_h9_ = d_42_authBase_
        def iife46_(_pat_let419_0):
            def iife47_(d_62_dt__update__tmp_h10_):
                def iife49_(_pat_let421_0):
                    def iife50_(d_63_dt__update__tmp_h11_):
                        def iife51_(_pat_let422_0):
                            def iife52_(d_64_dt__update_hconfirmed_h1_):
                                def iife53_(_pat_let423_0):
                                    def iife54_(d_65_dt__update_hcontributor_h1_):
                                        def iife55_(_pat_let424_0):
                                            def iife56_(d_66_dt__update_hstatus_h4_):
                                                return ToolCall_ToolCall((d_63_dt__update__tmp_h11_).toolCallId, (d_63_dt__update__tmp_h11_).toolName, (d_63_dt__update__tmp_h11_).displayName, d_66_dt__update_hstatus_h4_, (d_63_dt__update__tmp_h11_).intention, d_65_dt__update_hcontributor_h1_, (d_63_dt__update__tmp_h11_).meta, (d_63_dt__update__tmp_h11_).invocationMessage, (d_63_dt__update__tmp_h11_).toolInput, (d_63_dt__update__tmp_h11_).confirmationTitle, (d_63_dt__update__tmp_h11_).riskAssessment, (d_63_dt__update__tmp_h11_).edits, (d_63_dt__update__tmp_h11_).editable, (d_63_dt__update__tmp_h11_).options, d_64_dt__update_hconfirmed_h1_, (d_63_dt__update__tmp_h11_).success, (d_63_dt__update__tmp_h11_).pastTenseMessage, (d_63_dt__update__tmp_h11_).reason, (d_63_dt__update__tmp_h11_).reasonMessage, (d_63_dt__update__tmp_h11_).userSuggestion, (d_63_dt__update__tmp_h11_).selectedOption, (d_63_dt__update__tmp_h11_).content, (d_63_dt__update__tmp_h11_).structuredContent, (d_63_dt__update__tmp_h11_).error, (d_63_dt__update__tmp_h11_).auth, (d_63_dt__update__tmp_h11_).partialInput)
                                            return iife56_(_pat_let424_0)
                                        return iife55_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running")))
                                    return iife54_(_pat_let423_0)
                                return iife53_(AhpSkeleton.Option_Some(ToolCallContributor_ToolCallContributor(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "client")), ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "client")))})))))
                            return iife52_(_pat_let422_0)
                        return iife51_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "not-needed"))))
                    return iife50_(_pat_let421_0)
                def iife48_(_pat_let420_0):
                    def iife57_(d_67_dt__update_hparts_h3_):
                        return Turn_Turn((d_62_dt__update__tmp_h10_).turnId, (d_62_dt__update__tmp_h10_).message, d_67_dt__update_hparts_h3_, (d_62_dt__update__tmp_h10_).state, (d_62_dt__update__tmp_h10_).usage, (d_62_dt__update__tmp_h10_).error)
                    return iife57_(_pat_let420_0)
                return iife48_(_dafny.SeqWithoutIsStrInference([Part_PToolCall(iife49_(default__.TC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-client")), AhpSkeleton.Option_None())))]))
            return iife47_(_pat_let419_0)
        d_61_dt__update_hactiveTurn_h3_ = AhpSkeleton.Option_Some(iife46_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))
        d_59_clientBase_ = ChatState_ChatState((d_60_dt__update__tmp_h9_).turns, (d_60_dt__update__tmp_h9_).title, (d_60_dt__update__tmp_h9_).status, (d_60_dt__update__tmp_h9_).modifiedAt, (d_60_dt__update__tmp_h9_).draft, (d_60_dt__update__tmp_h9_).activity, d_61_dt__update_hactiveTurn_h3_, (d_60_dt__update__tmp_h9_).steeringMessage, (d_60_dt__update__tmp_h9_).queuedMessages, (d_60_dt__update__tmp_h9_).nextCursor)
        if (default__.apply1(d_59_clientBase_, ChatAction_CToolCallAuthRequired(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-client")), d_54_challenge_, AhpSkeleton.Option_None()))) == (d_59_clientBase_):
            pass_ = (pass_) + (1)
        d_68_authFailed_: ChatState
        d_68_authFailed_ = default__.apply1(d_55_authState_, ChatAction_CToolCallComplete(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-auth")), False, AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Cancelled search"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cancelled")))), True, AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "complete-meta"))))))
        d_69_failedTc_: AhpSkeleton.Option
        d_69_failedTc_ = default__.firstActiveToolCall(d_68_authFailed_)
        if (((((((d_68_authFailed_).status) == (8)) and ((d_69_failedTc_).is_Some)) and ((((d_69_failedTc_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "completed"))))) and ((((d_69_failedTc_).value).success) == (AhpSkeleton.Option_Some(False)))) and ((((d_69_failedTc_).value).pastTenseMessage) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Cancelled search")))))) and ((((d_69_failedTc_).value).content) == (AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "partial")))))):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_55_authState_, ChatAction_CToolCallComplete(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-auth")), True, AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Should not apply"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), False, AhpSkeleton.Option_None()))) == (d_55_authState_):
            pass_ = (pass_) + (1)
        return pass_

    @staticmethod
    def runPendingHistory():
        pass_: int = int(0)
        pass_ = 0
        d_0_s_: ChatState
        d_0_s_ = default__.C0()
        d_1_act1_: ChatState
        d_2_dt__update__tmp_h0_ = d_0_s_
        d_3_dt__update_hactiveTurn_h0_ = AhpSkeleton.Option_Some(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))))
        d_1_act1_ = ChatState_ChatState((d_2_dt__update__tmp_h0_).turns, (d_2_dt__update__tmp_h0_).title, (d_2_dt__update__tmp_h0_).status, (d_2_dt__update__tmp_h0_).modifiedAt, (d_2_dt__update__tmp_h0_).draft, (d_2_dt__update__tmp_h0_).activity, d_3_dt__update_hactiveTurn_h0_, (d_2_dt__update__tmp_h0_).steeringMessage, (d_2_dt__update__tmp_h0_).queuedMessages, (d_2_dt__update__tmp_h0_).nextCursor)
        if ((default__.apply1(d_0_s_, ChatAction_CPendingMessageSet(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "steering")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sm-1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Focus on tests")))))).steeringMessage) == (AhpSkeleton.Option_Some(QMsg_QMsg(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sm-1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Focus on tests")))))):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_0_s_, ChatAction_CPendingMessageRemoved(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "steering")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sm-unknown"))))) == (d_0_s_):
            pass_ = (pass_) + (1)
        def iife0_(_pat_let425_0):
            def iife1_(d_4_dt__update__tmp_h1_):
                def iife2_(_pat_let426_0):
                    def iife3_(d_5_dt__update_hsteeringMessage_h0_):
                        return ChatState_ChatState((d_4_dt__update__tmp_h1_).turns, (d_4_dt__update__tmp_h1_).title, (d_4_dt__update__tmp_h1_).status, (d_4_dt__update__tmp_h1_).modifiedAt, (d_4_dt__update__tmp_h1_).draft, (d_4_dt__update__tmp_h1_).activity, (d_4_dt__update__tmp_h1_).activeTurn, d_5_dt__update_hsteeringMessage_h0_, (d_4_dt__update__tmp_h1_).queuedMessages, (d_4_dt__update__tmp_h1_).nextCursor)
                    return iife3_(_pat_let426_0)
                return iife2_(AhpSkeleton.Option_Some(QMsg_QMsg(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sm-1")), ConfluxCodec.Json_JNull())))
            return iife1_(_pat_let425_0)
        if ((default__.apply1(iife0_(d_0_s_), ChatAction_CPendingMessageRemoved(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "steering")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sm-1"))))).steeringMessage) == (AhpSkeleton.Option_None()):
            pass_ = (pass_) + (1)
        d_6_q_: QMsg
        d_6_q_ = QMsg_QMsg(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q-1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))))
        if ((default__.apply1(d_0_s_, ChatAction_CPendingMessageSet(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "queued")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q-1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")))))).queuedMessages) == (_dafny.SeqWithoutIsStrInference([d_6_q_])):
            pass_ = (pass_) + (1)
        pat_let_tv0_ = d_6_q_
        def iife4_(_pat_let427_0):
            def iife5_(d_7_dt__update__tmp_h2_):
                def iife6_(_pat_let428_0):
                    def iife7_(d_8_dt__update_hqueuedMessages_h0_):
                        return ChatState_ChatState((d_7_dt__update__tmp_h2_).turns, (d_7_dt__update__tmp_h2_).title, (d_7_dt__update__tmp_h2_).status, (d_7_dt__update__tmp_h2_).modifiedAt, (d_7_dt__update__tmp_h2_).draft, (d_7_dt__update__tmp_h2_).activity, (d_7_dt__update__tmp_h2_).activeTurn, (d_7_dt__update__tmp_h2_).steeringMessage, d_8_dt__update_hqueuedMessages_h0_, (d_7_dt__update__tmp_h2_).nextCursor)
                    return iife7_(_pat_let428_0)
                return iife6_(_dafny.SeqWithoutIsStrInference([pat_let_tv0_, QMsg_QMsg(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q-2")), ConfluxCodec.Json_JNull())]))
            return iife5_(_pat_let427_0)
        if ((default__.apply1(iife4_(d_0_s_), ChatAction_CPendingMessageRemoved(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "queued")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q-1"))))).queuedMessages) == (_dafny.SeqWithoutIsStrInference([QMsg_QMsg(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q-2")), ConfluxCodec.Json_JNull())])):
            pass_ = (pass_) + (1)
        d_9_three_: ChatState
        d_10_dt__update__tmp_h3_ = d_0_s_
        d_11_dt__update_hturns_h0_ = _dafny.SeqWithoutIsStrInference([default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))), default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t2"))), default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t3")))])
        d_9_three_ = ChatState_ChatState(d_11_dt__update_hturns_h0_, (d_10_dt__update__tmp_h3_).title, (d_10_dt__update__tmp_h3_).status, (d_10_dt__update__tmp_h3_).modifiedAt, (d_10_dt__update__tmp_h3_).draft, (d_10_dt__update__tmp_h3_).activity, (d_10_dt__update__tmp_h3_).activeTurn, (d_10_dt__update__tmp_h3_).steeringMessage, (d_10_dt__update__tmp_h3_).queuedMessages, (d_10_dt__update__tmp_h3_).nextCursor)
        if ((default__.apply1(d_9_three_, ChatAction_CTruncated(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t2")))))).turns) == (_dafny.SeqWithoutIsStrInference([default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))), default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t2")))])):
            pass_ = (pass_) + (1)
        def iife8_(_pat_let429_0):
            def iife9_(d_12_dt__update__tmp_h4_):
                def iife10_(_pat_let430_0):
                    def iife11_(d_13_dt__update_hnextCursor_h0_):
                        return ChatState_ChatState((d_12_dt__update__tmp_h4_).turns, (d_12_dt__update__tmp_h4_).title, (d_12_dt__update__tmp_h4_).status, (d_12_dt__update__tmp_h4_).modifiedAt, (d_12_dt__update__tmp_h4_).draft, (d_12_dt__update__tmp_h4_).activity, (d_12_dt__update__tmp_h4_).activeTurn, (d_12_dt__update__tmp_h4_).steeringMessage, (d_12_dt__update__tmp_h4_).queuedMessages, d_13_dt__update_hnextCursor_h0_)
                    return iife11_(_pat_let430_0)
                return iife10_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cur"))))
            return iife9_(_pat_let429_0)
        def iife12_(_pat_let431_0):
            def iife13_(d_14_dt__update__tmp_h5_):
                def iife14_(_pat_let432_0):
                    def iife15_(d_15_dt__update_hnextCursor_h1_):
                        def iife16_(_pat_let433_0):
                            def iife17_(d_16_dt__update_hturns_h1_):
                                return ChatState_ChatState(d_16_dt__update_hturns_h1_, (d_14_dt__update__tmp_h5_).title, (d_14_dt__update__tmp_h5_).status, (d_14_dt__update__tmp_h5_).modifiedAt, (d_14_dt__update__tmp_h5_).draft, (d_14_dt__update__tmp_h5_).activity, (d_14_dt__update__tmp_h5_).activeTurn, (d_14_dt__update__tmp_h5_).steeringMessage, (d_14_dt__update__tmp_h5_).queuedMessages, d_15_dt__update_hnextCursor_h1_)
                            return iife17_(_pat_let433_0)
                        return iife16_(_dafny.SeqWithoutIsStrInference([]))
                    return iife15_(_pat_let432_0)
                return iife14_(AhpSkeleton.Option_None())
            return iife13_(_pat_let431_0)
        if (default__.apply1(iife8_(d_9_three_), ChatAction_CTruncated(AhpSkeleton.Option_None()))) == (iife12_(d_9_three_)):
            pass_ = (pass_) + (1)
        d_17_qa_: QMsg
        d_17_qa_ = QMsg_QMsg(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "A"))))
        d_18_qb_: QMsg
        d_18_qb_ = QMsg_QMsg(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "B"))))
        d_19_qc_: QMsg
        d_19_qc_ = QMsg_QMsg(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "C"))))
        pat_let_tv1_ = d_17_qa_
        pat_let_tv2_ = d_18_qb_
        pat_let_tv3_ = d_19_qc_
        def iife18_(_pat_let434_0):
            def iife19_(d_20_dt__update__tmp_h6_):
                def iife20_(_pat_let435_0):
                    def iife21_(d_21_dt__update_hqueuedMessages_h1_):
                        return ChatState_ChatState((d_20_dt__update__tmp_h6_).turns, (d_20_dt__update__tmp_h6_).title, (d_20_dt__update__tmp_h6_).status, (d_20_dt__update__tmp_h6_).modifiedAt, (d_20_dt__update__tmp_h6_).draft, (d_20_dt__update__tmp_h6_).activity, (d_20_dt__update__tmp_h6_).activeTurn, (d_20_dt__update__tmp_h6_).steeringMessage, d_21_dt__update_hqueuedMessages_h1_, (d_20_dt__update__tmp_h6_).nextCursor)
                    return iife21_(_pat_let435_0)
                return iife20_(_dafny.SeqWithoutIsStrInference([pat_let_tv1_, pat_let_tv2_, pat_let_tv3_]))
            return iife19_(_pat_let434_0)
        if ((default__.apply1(iife18_(d_0_s_), ChatAction_CQueuedReordered(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))])))).queuedMessages) == (_dafny.SeqWithoutIsStrInference([d_19_qc_, d_17_qa_, d_18_qb_])):
            pass_ = (pass_) + (1)
        def iife22_(_pat_let436_0):
            def iife23_(d_22_dt__update__tmp_h7_):
                def iife24_(_pat_let437_0):
                    def iife25_(d_23_dt__update_hturns_h2_):
                        return ChatState_ChatState(d_23_dt__update_hturns_h2_, (d_22_dt__update__tmp_h7_).title, (d_22_dt__update__tmp_h7_).status, (d_22_dt__update__tmp_h7_).modifiedAt, (d_22_dt__update__tmp_h7_).draft, (d_22_dt__update__tmp_h7_).activity, (d_22_dt__update__tmp_h7_).activeTurn, (d_22_dt__update__tmp_h7_).steeringMessage, (d_22_dt__update__tmp_h7_).queuedMessages, (d_22_dt__update__tmp_h7_).nextCursor)
                    return iife25_(_pat_let437_0)
                return iife24_(_dafny.SeqWithoutIsStrInference([default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t3")))]))
            return iife23_(_pat_let436_0)
        def iife26_(_pat_let438_0):
            def iife27_(d_24_dt__update__tmp_h8_):
                def iife28_(_pat_let439_0):
                    def iife29_(d_25_dt__update_hnextCursor_h2_):
                        def iife30_(_pat_let440_0):
                            def iife31_(d_26_dt__update_hturns_h3_):
                                return ChatState_ChatState(d_26_dt__update_hturns_h3_, (d_24_dt__update__tmp_h8_).title, (d_24_dt__update__tmp_h8_).status, (d_24_dt__update__tmp_h8_).modifiedAt, (d_24_dt__update__tmp_h8_).draft, (d_24_dt__update__tmp_h8_).activity, (d_24_dt__update__tmp_h8_).activeTurn, (d_24_dt__update__tmp_h8_).steeringMessage, (d_24_dt__update__tmp_h8_).queuedMessages, d_25_dt__update_hnextCursor_h2_)
                            return iife31_(_pat_let440_0)
                        return iife30_(_dafny.SeqWithoutIsStrInference([default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))), default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t2"))), default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t3")))]))
                    return iife29_(_pat_let439_0)
                return iife28_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cur-0"))))
            return iife27_(_pat_let438_0)
        if (default__.apply1(iife22_(d_0_s_), ChatAction_CTurnsLoaded(_dafny.SeqWithoutIsStrInference([default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))), default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t2")))]), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cur-0")))))) == (iife26_(d_0_s_)):
            pass_ = (pass_) + (1)
        def iife32_(_pat_let441_0):
            def iife33_(d_27_dt__update__tmp_h9_):
                def iife34_(_pat_let442_0):
                    def iife35_(d_28_dt__update_hturns_h4_):
                        return ChatState_ChatState(d_28_dt__update_hturns_h4_, (d_27_dt__update__tmp_h9_).title, (d_27_dt__update__tmp_h9_).status, (d_27_dt__update__tmp_h9_).modifiedAt, (d_27_dt__update__tmp_h9_).draft, (d_27_dt__update__tmp_h9_).activity, (d_27_dt__update__tmp_h9_).activeTurn, (d_27_dt__update__tmp_h9_).steeringMessage, (d_27_dt__update__tmp_h9_).queuedMessages, (d_27_dt__update__tmp_h9_).nextCursor)
                    return iife35_(_pat_let442_0)
                return iife34_(_dafny.SeqWithoutIsStrInference([default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t3")))]))
            return iife33_(_pat_let441_0)
        if ((default__.apply1(iife32_(d_0_s_), ChatAction_CTurnsLoaded(_dafny.SeqWithoutIsStrInference([default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))), default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t2"))), default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t3")))]), AhpSkeleton.Option_None()))).turns) == (_dafny.SeqWithoutIsStrInference([default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))), default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t2"))), default__.TDone(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t3")))])):
            pass_ = (pass_) + (1)
        d_29_running_: ChatState
        d_30_dt__update__tmp_h10_ = d_1_act1_
        def iife36_(_pat_let443_0):
            def iife37_(d_32_dt__update__tmp_h11_):
                def iife39_(_pat_let445_0):
                    def iife40_(d_33_dt__update__tmp_h12_):
                        def iife41_(_pat_let446_0):
                            def iife42_(d_34_dt__update_hstatus_h0_):
                                return ToolCall_ToolCall((d_33_dt__update__tmp_h12_).toolCallId, (d_33_dt__update__tmp_h12_).toolName, (d_33_dt__update__tmp_h12_).displayName, d_34_dt__update_hstatus_h0_, (d_33_dt__update__tmp_h12_).intention, (d_33_dt__update__tmp_h12_).contributor, (d_33_dt__update__tmp_h12_).meta, (d_33_dt__update__tmp_h12_).invocationMessage, (d_33_dt__update__tmp_h12_).toolInput, (d_33_dt__update__tmp_h12_).confirmationTitle, (d_33_dt__update__tmp_h12_).riskAssessment, (d_33_dt__update__tmp_h12_).edits, (d_33_dt__update__tmp_h12_).editable, (d_33_dt__update__tmp_h12_).options, (d_33_dt__update__tmp_h12_).confirmed, (d_33_dt__update__tmp_h12_).success, (d_33_dt__update__tmp_h12_).pastTenseMessage, (d_33_dt__update__tmp_h12_).reason, (d_33_dt__update__tmp_h12_).reasonMessage, (d_33_dt__update__tmp_h12_).userSuggestion, (d_33_dt__update__tmp_h12_).selectedOption, (d_33_dt__update__tmp_h12_).content, (d_33_dt__update__tmp_h12_).structuredContent, (d_33_dt__update__tmp_h12_).error, (d_33_dt__update__tmp_h12_).auth, (d_33_dt__update__tmp_h12_).partialInput)
                            return iife42_(_pat_let446_0)
                        return iife41_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running")))
                    return iife40_(_pat_let445_0)
                def iife38_(_pat_let444_0):
                    def iife43_(d_35_dt__update_hparts_h0_):
                        return Turn_Turn((d_32_dt__update__tmp_h11_).turnId, (d_32_dt__update__tmp_h11_).message, d_35_dt__update_hparts_h0_, (d_32_dt__update__tmp_h11_).state, (d_32_dt__update__tmp_h11_).usage, (d_32_dt__update__tmp_h11_).error)
                    return iife43_(_pat_let444_0)
                return iife38_(_dafny.SeqWithoutIsStrInference([Part_PToolCall(iife39_(default__.TC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_None())))]))
            return iife37_(_pat_let443_0)
        d_31_dt__update_hactiveTurn_h1_ = AhpSkeleton.Option_Some(iife36_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))
        d_29_running_ = ChatState_ChatState((d_30_dt__update__tmp_h10_).turns, (d_30_dt__update__tmp_h10_).title, (d_30_dt__update__tmp_h10_).status, (d_30_dt__update__tmp_h10_).modifiedAt, (d_30_dt__update__tmp_h10_).draft, (d_30_dt__update__tmp_h10_).activity, d_31_dt__update_hactiveTurn_h1_, (d_30_dt__update__tmp_h10_).steeringMessage, (d_30_dt__update__tmp_h10_).queuedMessages, (d_30_dt__update__tmp_h10_).nextCursor)
        d_36_afterCc_: ChatState
        d_36_afterCc_ = default__.apply1(d_29_running_, ChatAction_CToolCallContentChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "terminal-ref"))), AhpSkeleton.Option_None()))
        d_37_ccTc_: Part
        d_37_ccTc_ = ((((d_36_afterCc_).activeTurn).value).parts)[0]
        if (((d_37_ccTc_).is_PToolCall) and ((((d_37_ccTc_).tc).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running"))))) and ((((d_37_ccTc_).tc).content) == (AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "terminal-ref")))))):
            pass_ = (pass_) + (1)
        d_38_withR_: ChatState
        d_39_dt__update__tmp_h13_ = d_1_act1_
        def iife44_(_pat_let447_0):
            def iife45_(d_41_dt__update__tmp_h14_):
                def iife46_(_pat_let448_0):
                    def iife47_(d_42_dt__update_hparts_h1_):
                        return Turn_Turn((d_41_dt__update__tmp_h14_).turnId, (d_41_dt__update__tmp_h14_).message, d_42_dt__update_hparts_h1_, (d_41_dt__update__tmp_h14_).state, (d_41_dt__update__tmp_h14_).usage, (d_41_dt__update__tmp_h14_).error)
                    return iife47_(_pat_let448_0)
                return iife46_(_dafny.SeqWithoutIsStrInference([Part_PReasoning(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))]))
            return iife45_(_pat_let447_0)
        d_40_dt__update_hactiveTurn_h2_ = AhpSkeleton.Option_Some(iife44_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))
        d_38_withR_ = ChatState_ChatState((d_39_dt__update__tmp_h13_).turns, (d_39_dt__update__tmp_h13_).title, (d_39_dt__update__tmp_h13_).status, (d_39_dt__update__tmp_h13_).modifiedAt, (d_39_dt__update__tmp_h13_).draft, (d_39_dt__update__tmp_h13_).activity, d_40_dt__update_hactiveTurn_h2_, (d_39_dt__update__tmp_h13_).steeringMessage, (d_39_dt__update__tmp_h13_).queuedMessages, (d_39_dt__update__tmp_h13_).nextCursor)
        d_43_r2_: ChatState
        d_43_r2_ = default__.fold(d_38_withR_, _dafny.SeqWithoutIsStrInference([ChatAction_CReasoning(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Thinking about "))), ChatAction_CReasoning(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "the plan")))]))
        def iife48_(_pat_let449_0):
            def iife49_(d_44_dt__update__tmp_h15_):
                def iife50_(_pat_let450_0):
                    def iife51_(d_45_dt__update_hparts_h2_):
                        return Turn_Turn((d_44_dt__update__tmp_h15_).turnId, (d_44_dt__update__tmp_h15_).message, d_45_dt__update_hparts_h2_, (d_44_dt__update__tmp_h15_).state, (d_44_dt__update__tmp_h15_).usage, (d_44_dt__update__tmp_h15_).error)
                    return iife51_(_pat_let450_0)
                return iife50_(_dafny.SeqWithoutIsStrInference([Part_PReasoning(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Thinking about the plan")))]))
            return iife49_(_pat_let449_0)
        if ((d_43_r2_).activeTurn) == (AhpSkeleton.Option_Some(iife48_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))):
            pass_ = (pass_) + (1)
        return pass_

    @staticmethod
    def runInputFlow():
        pass_: int = int(0)
        pass_ = 0
        d_0_s_: ChatState
        d_0_s_ = default__.C0()
        d_1_openBody_: _dafny.Map
        d_1_openBody_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "message")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Clarify requirements"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "questions")): ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "id")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q1"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "required")): ConfluxCodec.Json_JBool(True), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "options")): ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "node")))]))}))])), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "url")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "https://example.test/review")))})
        d_2_at1_: ChatState
        d_3_dt__update__tmp_h0_ = d_0_s_
        d_4_dt__update_hstatus_h0_ = 8
        d_5_dt__update_hactiveTurn_h0_ = AhpSkeleton.Option_Some(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turn-1"))))
        d_2_at1_ = ChatState_ChatState((d_3_dt__update__tmp_h0_).turns, (d_3_dt__update__tmp_h0_).title, d_4_dt__update_hstatus_h0_, (d_3_dt__update__tmp_h0_).modifiedAt, (d_3_dt__update__tmp_h0_).draft, (d_3_dt__update__tmp_h0_).activity, d_5_dt__update_hactiveTurn_h0_, (d_3_dt__update__tmp_h0_).steeringMessage, (d_3_dt__update__tmp_h0_).queuedMessages, (d_3_dt__update__tmp_h0_).nextCursor)
        d_6_afterReq_: ChatState
        d_6_afterReq_ = default__.apply1(d_2_at1_, ChatAction_CInputRequested(InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.Map({}), d_1_openBody_)))
        if ((((d_6_afterReq_).status) == (24)) and (((d_6_afterReq_).activeTurn).is_Some)) and (((((d_6_afterReq_).activeTurn).value).parts) == (_dafny.SeqWithoutIsStrInference([Part_PInputRequest(InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.Map({}), d_1_openBody_), AhpSkeleton.Option_None())]))):
            pass_ = (pass_) + (1)
        def iife0_(_pat_let451_0):
            def iife1_(d_7_dt__update__tmp_h1_):
                def iife2_(_pat_let452_0):
                    def iife3_(d_8_dt__update_hstatus_h1_):
                        return ChatState_ChatState((d_7_dt__update__tmp_h1_).turns, (d_7_dt__update__tmp_h1_).title, d_8_dt__update_hstatus_h1_, (d_7_dt__update__tmp_h1_).modifiedAt, (d_7_dt__update__tmp_h1_).draft, (d_7_dt__update__tmp_h1_).activity, (d_7_dt__update__tmp_h1_).activeTurn, (d_7_dt__update__tmp_h1_).steeringMessage, (d_7_dt__update__tmp_h1_).queuedMessages, (d_7_dt__update__tmp_h1_).nextCursor)
                    return iife3_(_pat_let452_0)
                return iife2_(8)
            return iife1_(_pat_let451_0)
        def iife4_(_pat_let453_0):
            def iife5_(d_9_dt__update__tmp_h2_):
                def iife6_(_pat_let454_0):
                    def iife7_(d_10_dt__update_hstatus_h2_):
                        return ChatState_ChatState((d_9_dt__update__tmp_h2_).turns, (d_9_dt__update__tmp_h2_).title, d_10_dt__update_hstatus_h2_, (d_9_dt__update__tmp_h2_).modifiedAt, (d_9_dt__update__tmp_h2_).draft, (d_9_dt__update__tmp_h2_).activity, (d_9_dt__update__tmp_h2_).activeTurn, (d_9_dt__update__tmp_h2_).steeringMessage, (d_9_dt__update__tmp_h2_).queuedMessages, (d_9_dt__update__tmp_h2_).nextCursor)
                    return iife7_(_pat_let454_0)
                return iife6_(8)
            return iife5_(_pat_let453_0)
        if (default__.apply1(iife0_(d_0_s_), ChatAction_CInputRequested(InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orphan-1")), _dafny.Map({}), d_1_openBody_)))) == (iife4_(d_0_s_)):
            pass_ = (pass_) + (1)
        pat_let_tv0_ = d_1_openBody_
        d_11_reqState_: ChatState
        d_12_dt__update__tmp_h3_ = d_2_at1_
        d_13_dt__update_hstatus_h3_ = 24
        def iife8_(_pat_let455_0):
            def iife9_(d_15_dt__update__tmp_h4_):
                def iife10_(_pat_let456_0):
                    def iife11_(d_16_dt__update_hparts_h0_):
                        return Turn_Turn((d_15_dt__update__tmp_h4_).turnId, (d_15_dt__update__tmp_h4_).message, d_16_dt__update_hparts_h0_, (d_15_dt__update__tmp_h4_).state, (d_15_dt__update__tmp_h4_).usage, (d_15_dt__update__tmp_h4_).error)
                    return iife11_(_pat_let456_0)
                return iife10_(_dafny.SeqWithoutIsStrInference([Part_PInputRequest(InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.Map({}), pat_let_tv0_), AhpSkeleton.Option_None())]))
            return iife9_(_pat_let455_0)
        d_14_dt__update_hactiveTurn_h1_ = AhpSkeleton.Option_Some(iife8_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turn-1")))))
        d_11_reqState_ = ChatState_ChatState((d_12_dt__update__tmp_h3_).turns, (d_12_dt__update__tmp_h3_).title, d_13_dt__update_hstatus_h3_, (d_12_dt__update__tmp_h3_).modifiedAt, (d_12_dt__update__tmp_h3_).draft, (d_12_dt__update__tmp_h3_).activity, d_14_dt__update_hactiveTurn_h1_, (d_12_dt__update__tmp_h3_).steeringMessage, (d_12_dt__update__tmp_h3_).queuedMessages, (d_12_dt__update__tmp_h3_).nextCursor)
        if ((((default__.apply1(d_11_reqState_, ChatAction_CInputAnswerChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q1")), AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "node"))))))).activeTurn).value).parts) == (_dafny.SeqWithoutIsStrInference([Part_PInputRequest(InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q1")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "node")))}), d_1_openBody_), AhpSkeleton.Option_None())])):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_11_reqState_, ChatAction_CInputAnswerChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "missing")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q1")), AhpSkeleton.Option_None()))) == (d_11_reqState_):
            pass_ = (pass_) + (1)
        d_17_afterDone_: ChatState
        d_17_afterDone_ = default__.apply1(d_11_reqState_, ChatAction_CInputCompleted(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "accept")), _dafny.Map({})))
        if ((((d_17_afterDone_).status) == (8)) and (((d_17_afterDone_).activeTurn).is_Some)) and (((((d_17_afterDone_).activeTurn).value).parts) == (_dafny.SeqWithoutIsStrInference([Part_PInputRequest(InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.Map({}), d_1_openBody_), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "accept"))))]))):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_11_reqState_, ChatAction_CInputCompleted(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "missing")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cancel")), _dafny.Map({})))) == (d_11_reqState_):
            pass_ = (pass_) + (1)
        d_18_flow_: ChatState
        d_18_flow_ = default__.fold(d_0_s_, _dafny.SeqWithoutIsStrInference([ChatAction_CTurnStarted(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turn-1")), ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Hello"))), AhpSkeleton.Option_None()), ChatAction_CInputRequested(InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.Map({}), d_1_openBody_)), ChatAction_CInputAnswerChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q1")), AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "node"))))), ChatAction_CInputAnswerChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q2")), AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "small"))))), ChatAction_CInputCompleted(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "accept")), _dafny.Map({}))]))
        if ((((d_18_flow_).status) == (8)) and (((d_18_flow_).activeTurn).is_Some)) and (((((d_18_flow_).activeTurn).value).parts) == (_dafny.SeqWithoutIsStrInference([Part_PInputRequest(InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "input-1")), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q1")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "node"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q2")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "small")))}), d_1_openBody_), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "accept"))))]))):
            pass_ = (pass_) + (1)
        pat_let_tv1_ = d_1_openBody_
        d_19_overrideState_: ChatState
        d_20_dt__update__tmp_h5_ = d_2_at1_
        d_21_dt__update_hstatus_h4_ = 24
        def iife12_(_pat_let457_0):
            def iife13_(d_23_dt__update__tmp_h6_):
                def iife14_(_pat_let458_0):
                    def iife15_(d_24_dt__update_hparts_h1_):
                        return Turn_Turn((d_23_dt__update__tmp_h6_).turnId, (d_23_dt__update__tmp_h6_).message, d_24_dt__update_hparts_h1_, (d_23_dt__update__tmp_h6_).state, (d_23_dt__update__tmp_h6_).usage, (d_23_dt__update__tmp_h6_).error)
                    return iife15_(_pat_let458_0)
                return iife14_(_dafny.SeqWithoutIsStrInference([Part_PInputRequest(InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "review-1")), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "approve")): ConfluxCodec.Json_JBool(True)}), pat_let_tv1_), AhpSkeleton.Option_None())]))
            return iife13_(_pat_let457_0)
        d_22_dt__update_hactiveTurn_h2_ = AhpSkeleton.Option_Some(iife12_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turn-1")))))
        d_19_overrideState_ = ChatState_ChatState((d_20_dt__update__tmp_h5_).turns, (d_20_dt__update__tmp_h5_).title, d_21_dt__update_hstatus_h4_, (d_20_dt__update__tmp_h5_).modifiedAt, (d_20_dt__update__tmp_h5_).draft, (d_20_dt__update__tmp_h5_).activity, d_22_dt__update_hactiveTurn_h2_, (d_20_dt__update__tmp_h5_).steeringMessage, (d_20_dt__update__tmp_h5_).queuedMessages, (d_20_dt__update__tmp_h5_).nextCursor)
        d_25_overridden_: ChatState
        d_25_overridden_ = default__.apply1(d_19_overrideState_, ChatAction_CInputCompleted(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "review-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "decline")), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "approve")): ConfluxCodec.Json_JBool(False)})))
        if ((((d_25_overridden_).status) == (8)) and (((d_25_overridden_).activeTurn).is_Some)) and (((((d_25_overridden_).activeTurn).value).parts) == (_dafny.SeqWithoutIsStrInference([Part_PInputRequest(InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "review-1")), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "approve")): ConfluxCodec.Json_JBool(False)}), d_1_openBody_), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "decline"))))]))):
            pass_ = (pass_) + (1)
        return pass_

    @staticmethod
    def runResultConfirm():
        pass_: int = int(0)
        pass_ = 0
        d_0_s_: ChatState
        d_0_s_ = default__.C0()
        d_1_act1_: ChatState
        d_2_dt__update__tmp_h0_ = d_0_s_
        d_3_dt__update_hactiveTurn_h0_ = AhpSkeleton.Option_Some(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1"))))
        d_1_act1_ = ChatState_ChatState((d_2_dt__update__tmp_h0_).turns, (d_2_dt__update__tmp_h0_).title, (d_2_dt__update__tmp_h0_).status, (d_2_dt__update__tmp_h0_).modifiedAt, (d_2_dt__update__tmp_h0_).draft, (d_2_dt__update__tmp_h0_).activity, d_3_dt__update_hactiveTurn_h0_, (d_2_dt__update__tmp_h0_).steeringMessage, (d_2_dt__update__tmp_h0_).queuedMessages, (d_2_dt__update__tmp_h0_).nextCursor)
        d_4_rcLife_: ChatState
        d_4_rcLife_ = default__.fold(d_1_act1_, _dafny.SeqWithoutIsStrInference([ChatAction_CToolCallStart(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bash")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run Command")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChatAction_CToolCallReady(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "not-needed"))), AhpSkeleton.Option_None()), ChatAction_CToolCallComplete(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Done"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), False, AhpSkeleton.Option_None()), ChatAction_CToolCallResultConfirmed(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_None())]))
        d_5_rcTc_: AhpSkeleton.Option
        d_5_rcTc_ = default__.firstActiveToolCall(d_4_rcLife_)
        if (((d_5_rcTc_).is_Some) and ((((d_5_rcTc_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "completed"))))) and ((((d_5_rcTc_).value).success) == (AhpSkeleton.Option_Some(True))):
            pass_ = (pass_) + (1)
        d_6_rdLife_: ChatState
        d_6_rdLife_ = default__.fold(d_1_act1_, _dafny.SeqWithoutIsStrInference([ChatAction_CToolCallStart(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bash")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run Command")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChatAction_CToolCallReady(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "not-needed"))), AhpSkeleton.Option_None()), ChatAction_CToolCallComplete(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Done"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), True, AhpSkeleton.Option_None()), ChatAction_CToolCallResultConfirmed(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), False, AhpSkeleton.Option_None())]))
        d_7_rdTc_: AhpSkeleton.Option
        d_7_rdTc_ = default__.firstActiveToolCall(d_6_rdLife_)
        if (((((d_7_rdTc_).is_Some) and ((((d_7_rdTc_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cancelled"))))) and ((((d_7_rdTc_).value).confirmed) == (AhpSkeleton.Option_None()))) and ((((d_7_rdTc_).value).success) == (AhpSkeleton.Option_None()))) and ((((d_7_rdTc_).value).reason) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "result-denied"))))):
            pass_ = (pass_) + (1)
        d_8_runState_: ChatState
        d_9_dt__update__tmp_h1_ = d_1_act1_
        d_10_dt__update_hstatus_h0_ = 8
        def iife0_(_pat_let459_0):
            def iife1_(d_12_dt__update__tmp_h2_):
                def iife3_(_pat_let461_0):
                    def iife4_(d_13_dt__update__tmp_h3_):
                        def iife5_(_pat_let462_0):
                            def iife6_(d_14_dt__update_hconfirmed_h0_):
                                def iife7_(_pat_let463_0):
                                    def iife8_(d_15_dt__update_hstatus_h1_):
                                        return ToolCall_ToolCall((d_13_dt__update__tmp_h3_).toolCallId, (d_13_dt__update__tmp_h3_).toolName, (d_13_dt__update__tmp_h3_).displayName, d_15_dt__update_hstatus_h1_, (d_13_dt__update__tmp_h3_).intention, (d_13_dt__update__tmp_h3_).contributor, (d_13_dt__update__tmp_h3_).meta, (d_13_dt__update__tmp_h3_).invocationMessage, (d_13_dt__update__tmp_h3_).toolInput, (d_13_dt__update__tmp_h3_).confirmationTitle, (d_13_dt__update__tmp_h3_).riskAssessment, (d_13_dt__update__tmp_h3_).edits, (d_13_dt__update__tmp_h3_).editable, (d_13_dt__update__tmp_h3_).options, d_14_dt__update_hconfirmed_h0_, (d_13_dt__update__tmp_h3_).success, (d_13_dt__update__tmp_h3_).pastTenseMessage, (d_13_dt__update__tmp_h3_).reason, (d_13_dt__update__tmp_h3_).reasonMessage, (d_13_dt__update__tmp_h3_).userSuggestion, (d_13_dt__update__tmp_h3_).selectedOption, (d_13_dt__update__tmp_h3_).content, (d_13_dt__update__tmp_h3_).structuredContent, (d_13_dt__update__tmp_h3_).error, (d_13_dt__update__tmp_h3_).auth, (d_13_dt__update__tmp_h3_).partialInput)
                                    return iife8_(_pat_let463_0)
                                return iife7_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "running")))
                            return iife6_(_pat_let462_0)
                        return iife5_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "user-action"))))
                    return iife4_(_pat_let461_0)
                def iife2_(_pat_let460_0):
                    def iife9_(d_16_dt__update_hparts_h0_):
                        return Turn_Turn((d_12_dt__update__tmp_h2_).turnId, (d_12_dt__update__tmp_h2_).message, d_16_dt__update_hparts_h0_, (d_12_dt__update__tmp_h2_).state, (d_12_dt__update__tmp_h2_).usage, (d_12_dt__update__tmp_h2_).error)
                    return iife9_(_pat_let460_0)
                return iife2_(_dafny.SeqWithoutIsStrInference([Part_PToolCall(iife3_(default__.TC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_None())))]))
            return iife1_(_pat_let459_0)
        d_11_dt__update_hactiveTurn_h1_ = AhpSkeleton.Option_Some(iife0_(default__.T0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")))))
        d_8_runState_ = ChatState_ChatState((d_9_dt__update__tmp_h1_).turns, (d_9_dt__update__tmp_h1_).title, d_10_dt__update_hstatus_h0_, (d_9_dt__update__tmp_h1_).modifiedAt, (d_9_dt__update__tmp_h1_).draft, (d_9_dt__update__tmp_h1_).activity, d_11_dt__update_hactiveTurn_h1_, (d_9_dt__update__tmp_h1_).steeringMessage, (d_9_dt__update__tmp_h1_).queuedMessages, (d_9_dt__update__tmp_h1_).nextCursor)
        if (default__.apply1(d_8_runState_, ChatAction_CToolCallResultConfirmed(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_None()))) == (d_8_runState_):
            pass_ = (pass_) + (1)
        d_17_selLife_: ChatState
        d_17_selLife_ = default__.fold(d_1_act1_, _dafny.SeqWithoutIsStrInference([ChatAction_CToolCallStart(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bash")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run Command")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChatAction_CToolCallReady(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Run"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChatAction_CToolCallConfirmed(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "user-action"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None()), ChatAction_CToolCallComplete(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Done"))), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), False, AhpSkeleton.Option_None()), ChatAction_CToolCallResultConfirmed(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc-1")), True, AhpSkeleton.Option_None())]))
        d_18_selTc_: AhpSkeleton.Option
        d_18_selTc_ = default__.firstActiveToolCall(d_17_selLife_)
        if ((((d_18_selTc_).is_Some) and ((((d_18_selTc_).value).status) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "completed"))))) and ((((d_18_selTc_).value).confirmed) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "user-action")))))) and ((((d_18_selTc_).value).success) == (AhpSkeleton.Option_Some(True))):
            pass_ = (pass_) + (1)
        return pass_

    @staticmethod
    def RunCorpus():
        pass_: int = int(0)
        modeled: int = int(0)
        modeled = 54
        d_0_p1_: int
        out0_: int
        out0_ = default__.runScalarTurns()
        d_0_p1_ = out0_
        d_1_p2_: int
        out1_: int
        out1_ = default__.runToolCalls()
        d_1_p2_ = out1_
        d_2_p3_: int
        out2_: int
        out2_ = default__.runPendingHistory()
        d_2_p3_ = out2_
        d_3_p4_: int
        out3_: int
        out3_ = default__.runInputFlow()
        d_3_p4_ = out3_
        d_4_p5_: int
        out4_: int
        out4_ = default__.runResultConfirm()
        d_4_p5_ = out4_
        pass_ = ((((d_0_p1_) + (d_1_p2_)) + (d_2_p3_)) + (d_3_p4_)) + (d_4_p5_)
        return pass_, modeled

    @_dafny.classproperty
    def qKey(instance):
        def lambda0_(d_0_x_):
            return (d_0_x_).id_

        return lambda0_
    @_dafny.classproperty
    def ERR(instance):
        return 2
    @_dafny.classproperty
    def INP(instance):
        return 24
    @_dafny.classproperty
    def GEN(instance):
        return 8
    @_dafny.classproperty
    def IDLE(instance):
        return 1
    @_dafny.classproperty
    def ACT(instance):
        return 31
    @_dafny.classproperty
    def READ(instance):
        return 32

class ToolCallContributor:
    @classmethod
    def default(cls, ):
        return lambda: ToolCallContributor_ToolCallContributor(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxCodec.Json.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ToolCallContributor(self) -> bool:
        return isinstance(self, ToolCallContributor_ToolCallContributor)

class ToolCallContributor_ToolCallContributor(ToolCallContributor, NamedTuple('ToolCallContributor', [('kind', Any), ('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ToolCallContributor.ToolCallContributor({self.kind.VerbatimString(True)}, {_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ToolCallContributor_ToolCallContributor) and self.kind == __o.kind and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()


class ToolCall:
    @classmethod
    def default(cls, ):
        return lambda: ToolCall_ToolCall(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ToolCall(self) -> bool:
        return isinstance(self, ToolCall_ToolCall)

class ToolCall_ToolCall(ToolCall, NamedTuple('ToolCall', [('toolCallId', Any), ('toolName', Any), ('displayName', Any), ('status', Any), ('intention', Any), ('contributor', Any), ('meta', Any), ('invocationMessage', Any), ('toolInput', Any), ('confirmationTitle', Any), ('riskAssessment', Any), ('edits', Any), ('editable', Any), ('options', Any), ('confirmed', Any), ('success', Any), ('pastTenseMessage', Any), ('reason', Any), ('reasonMessage', Any), ('userSuggestion', Any), ('selectedOption', Any), ('content', Any), ('structuredContent', Any), ('error', Any), ('auth', Any), ('partialInput', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ToolCall.ToolCall({self.toolCallId.VerbatimString(True)}, {self.toolName.VerbatimString(True)}, {self.displayName.VerbatimString(True)}, {self.status.VerbatimString(True)}, {_dafny.string_of(self.intention)}, {_dafny.string_of(self.contributor)}, {_dafny.string_of(self.meta)}, {self.invocationMessage.VerbatimString(True)}, {_dafny.string_of(self.toolInput)}, {_dafny.string_of(self.confirmationTitle)}, {_dafny.string_of(self.riskAssessment)}, {_dafny.string_of(self.edits)}, {_dafny.string_of(self.editable)}, {_dafny.string_of(self.options)}, {_dafny.string_of(self.confirmed)}, {_dafny.string_of(self.success)}, {_dafny.string_of(self.pastTenseMessage)}, {_dafny.string_of(self.reason)}, {_dafny.string_of(self.reasonMessage)}, {_dafny.string_of(self.userSuggestion)}, {_dafny.string_of(self.selectedOption)}, {_dafny.string_of(self.content)}, {_dafny.string_of(self.structuredContent)}, {_dafny.string_of(self.error)}, {_dafny.string_of(self.auth)}, {_dafny.string_of(self.partialInput)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ToolCall_ToolCall) and self.toolCallId == __o.toolCallId and self.toolName == __o.toolName and self.displayName == __o.displayName and self.status == __o.status and self.intention == __o.intention and self.contributor == __o.contributor and self.meta == __o.meta and self.invocationMessage == __o.invocationMessage and self.toolInput == __o.toolInput and self.confirmationTitle == __o.confirmationTitle and self.riskAssessment == __o.riskAssessment and self.edits == __o.edits and self.editable == __o.editable and self.options == __o.options and self.confirmed == __o.confirmed and self.success == __o.success and self.pastTenseMessage == __o.pastTenseMessage and self.reason == __o.reason and self.reasonMessage == __o.reasonMessage and self.userSuggestion == __o.userSuggestion and self.selectedOption == __o.selectedOption and self.content == __o.content and self.structuredContent == __o.structuredContent and self.error == __o.error and self.auth == __o.auth and self.partialInput == __o.partialInput
    def __hash__(self) -> int:
        return super().__hash__()


class Part:
    @classmethod
    def default(cls, ):
        return lambda: Part_PMarkdown(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_PMarkdown(self) -> bool:
        return isinstance(self, Part_PMarkdown)
    @property
    def is_PReasoning(self) -> bool:
        return isinstance(self, Part_PReasoning)
    @property
    def is_PToolCall(self) -> bool:
        return isinstance(self, Part_PToolCall)
    @property
    def is_PInputRequest(self) -> bool:
        return isinstance(self, Part_PInputRequest)

class Part_PMarkdown(Part, NamedTuple('PMarkdown', [('id_', Any), ('content', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.Part.PMarkdown({self.id_.VerbatimString(True)}, {self.content.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Part_PMarkdown) and self.id_ == __o.id_ and self.content == __o.content
    def __hash__(self) -> int:
        return super().__hash__()

class Part_PReasoning(Part, NamedTuple('PReasoning', [('id_', Any), ('content', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.Part.PReasoning({self.id_.VerbatimString(True)}, {self.content.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Part_PReasoning) and self.id_ == __o.id_ and self.content == __o.content
    def __hash__(self) -> int:
        return super().__hash__()

class Part_PToolCall(Part, NamedTuple('PToolCall', [('tc', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.Part.PToolCall({_dafny.string_of(self.tc)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Part_PToolCall) and self.tc == __o.tc
    def __hash__(self) -> int:
        return super().__hash__()

class Part_PInputRequest(Part, NamedTuple('PInputRequest', [('req', Any), ('response', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.Part.PInputRequest({_dafny.string_of(self.req)}, {_dafny.string_of(self.response)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Part_PInputRequest) and self.req == __o.req and self.response == __o.response
    def __hash__(self) -> int:
        return super().__hash__()


class Turn:
    @classmethod
    def default(cls, ):
        return lambda: Turn_Turn(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxCodec.Json.default()(), _dafny.Seq({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Turn(self) -> bool:
        return isinstance(self, Turn_Turn)

class Turn_Turn(Turn, NamedTuple('Turn', [('turnId', Any), ('message', Any), ('parts', Any), ('state', Any), ('usage', Any), ('error', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.Turn.Turn({self.turnId.VerbatimString(True)}, {_dafny.string_of(self.message)}, {_dafny.string_of(self.parts)}, {self.state.VerbatimString(True)}, {_dafny.string_of(self.usage)}, {_dafny.string_of(self.error)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Turn_Turn) and self.turnId == __o.turnId and self.message == __o.message and self.parts == __o.parts and self.state == __o.state and self.usage == __o.usage and self.error == __o.error
    def __hash__(self) -> int:
        return super().__hash__()


class QMsg:
    @classmethod
    def default(cls, ):
        return lambda: QMsg_QMsg(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxCodec.Json.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_QMsg(self) -> bool:
        return isinstance(self, QMsg_QMsg)

class QMsg_QMsg(QMsg, NamedTuple('QMsg', [('id_', Any), ('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.QMsg.QMsg({self.id_.VerbatimString(True)}, {_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, QMsg_QMsg) and self.id_ == __o.id_ and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()


class InputReq:
    @classmethod
    def default(cls, ):
        return lambda: InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Map({}), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_InputReq(self) -> bool:
        return isinstance(self, InputReq_InputReq)

class InputReq_InputReq(InputReq, NamedTuple('InputReq', [('id_', Any), ('answers', Any), ('open_', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.InputReq.InputReq({self.id_.VerbatimString(True)}, {_dafny.string_of(self.answers)}, {_dafny.string_of(self.open_)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, InputReq_InputReq) and self.id_ == __o.id_ and self.answers == __o.answers and self.open_ == __o.open_
    def __hash__(self) -> int:
        return super().__hash__()


class ChatState:
    @classmethod
    def default(cls, ):
        return lambda: ChatState_ChatState(_dafny.Seq({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), _dafny.Seq({}), AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ChatState(self) -> bool:
        return isinstance(self, ChatState_ChatState)

class ChatState_ChatState(ChatState, NamedTuple('ChatState', [('turns', Any), ('title', Any), ('status', Any), ('modifiedAt', Any), ('draft', Any), ('activity', Any), ('activeTurn', Any), ('steeringMessage', Any), ('queuedMessages', Any), ('nextCursor', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatState.ChatState({_dafny.string_of(self.turns)}, {self.title.VerbatimString(True)}, {_dafny.string_of(self.status)}, {self.modifiedAt.VerbatimString(True)}, {_dafny.string_of(self.draft)}, {_dafny.string_of(self.activity)}, {_dafny.string_of(self.activeTurn)}, {_dafny.string_of(self.steeringMessage)}, {_dafny.string_of(self.queuedMessages)}, {_dafny.string_of(self.nextCursor)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatState_ChatState) and self.turns == __o.turns and self.title == __o.title and self.status == __o.status and self.modifiedAt == __o.modifiedAt and self.draft == __o.draft and self.activity == __o.activity and self.activeTurn == __o.activeTurn and self.steeringMessage == __o.steeringMessage and self.queuedMessages == __o.queuedMessages and self.nextCursor == __o.nextCursor
    def __hash__(self) -> int:
        return super().__hash__()


class ChatAction:
    @classmethod
    def default(cls, ):
        return lambda: ChatAction_CDraftChanged(AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_CDraftChanged(self) -> bool:
        return isinstance(self, ChatAction_CDraftChanged)
    @property
    def is_CActivityChanged(self) -> bool:
        return isinstance(self, ChatAction_CActivityChanged)
    @property
    def is_CTurnStarted(self) -> bool:
        return isinstance(self, ChatAction_CTurnStarted)
    @property
    def is_CTurnComplete(self) -> bool:
        return isinstance(self, ChatAction_CTurnComplete)
    @property
    def is_CTurnCancelled(self) -> bool:
        return isinstance(self, ChatAction_CTurnCancelled)
    @property
    def is_CUsage(self) -> bool:
        return isinstance(self, ChatAction_CUsage)
    @property
    def is_CError(self) -> bool:
        return isinstance(self, ChatAction_CError)
    @property
    def is_CResponsePart(self) -> bool:
        return isinstance(self, ChatAction_CResponsePart)
    @property
    def is_CDelta(self) -> bool:
        return isinstance(self, ChatAction_CDelta)
    @property
    def is_CToolCallStart(self) -> bool:
        return isinstance(self, ChatAction_CToolCallStart)
    @property
    def is_CToolCallDelta(self) -> bool:
        return isinstance(self, ChatAction_CToolCallDelta)
    @property
    def is_CToolCallReady(self) -> bool:
        return isinstance(self, ChatAction_CToolCallReady)
    @property
    def is_CToolCallConfirmed(self) -> bool:
        return isinstance(self, ChatAction_CToolCallConfirmed)
    @property
    def is_CToolCallAuthRequired(self) -> bool:
        return isinstance(self, ChatAction_CToolCallAuthRequired)
    @property
    def is_CToolCallAuthResolved(self) -> bool:
        return isinstance(self, ChatAction_CToolCallAuthResolved)
    @property
    def is_CToolCallComplete(self) -> bool:
        return isinstance(self, ChatAction_CToolCallComplete)
    @property
    def is_CToolCallResultConfirmed(self) -> bool:
        return isinstance(self, ChatAction_CToolCallResultConfirmed)
    @property
    def is_CToolCallContentChanged(self) -> bool:
        return isinstance(self, ChatAction_CToolCallContentChanged)
    @property
    def is_CReasoning(self) -> bool:
        return isinstance(self, ChatAction_CReasoning)
    @property
    def is_CTruncated(self) -> bool:
        return isinstance(self, ChatAction_CTruncated)
    @property
    def is_CQueuedReordered(self) -> bool:
        return isinstance(self, ChatAction_CQueuedReordered)
    @property
    def is_CTurnsLoaded(self) -> bool:
        return isinstance(self, ChatAction_CTurnsLoaded)
    @property
    def is_CInputRequested(self) -> bool:
        return isinstance(self, ChatAction_CInputRequested)
    @property
    def is_CInputAnswerChanged(self) -> bool:
        return isinstance(self, ChatAction_CInputAnswerChanged)
    @property
    def is_CInputCompleted(self) -> bool:
        return isinstance(self, ChatAction_CInputCompleted)
    @property
    def is_CPendingMessageSet(self) -> bool:
        return isinstance(self, ChatAction_CPendingMessageSet)
    @property
    def is_CPendingMessageRemoved(self) -> bool:
        return isinstance(self, ChatAction_CPendingMessageRemoved)
    @property
    def is_CUnknown(self) -> bool:
        return isinstance(self, ChatAction_CUnknown)

class ChatAction_CDraftChanged(ChatAction, NamedTuple('CDraftChanged', [('draft', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CDraftChanged({_dafny.string_of(self.draft)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CDraftChanged) and self.draft == __o.draft
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CActivityChanged(ChatAction, NamedTuple('CActivityChanged', [('activity', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CActivityChanged({_dafny.string_of(self.activity)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CActivityChanged) and self.activity == __o.activity
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CTurnStarted(ChatAction, NamedTuple('CTurnStarted', [('turnId', Any), ('message', Any), ('queuedMessageId', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CTurnStarted({self.turnId.VerbatimString(True)}, {_dafny.string_of(self.message)}, {_dafny.string_of(self.queuedMessageId)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CTurnStarted) and self.turnId == __o.turnId and self.message == __o.message and self.queuedMessageId == __o.queuedMessageId
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CTurnComplete(ChatAction, NamedTuple('CTurnComplete', [('turnId', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CTurnComplete({self.turnId.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CTurnComplete) and self.turnId == __o.turnId
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CTurnCancelled(ChatAction, NamedTuple('CTurnCancelled', [('turnId', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CTurnCancelled({self.turnId.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CTurnCancelled) and self.turnId == __o.turnId
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CUsage(ChatAction, NamedTuple('CUsage', [('turnId', Any), ('usage', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CUsage({self.turnId.VerbatimString(True)}, {_dafny.string_of(self.usage)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CUsage) and self.turnId == __o.turnId and self.usage == __o.usage
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CError(ChatAction, NamedTuple('CError', [('turnId', Any), ('err', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CError({self.turnId.VerbatimString(True)}, {_dafny.string_of(self.err)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CError) and self.turnId == __o.turnId and self.err == __o.err
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CResponsePart(ChatAction, NamedTuple('CResponsePart', [('turnId', Any), ('part', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CResponsePart({self.turnId.VerbatimString(True)}, {_dafny.string_of(self.part)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CResponsePart) and self.turnId == __o.turnId and self.part == __o.part
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CDelta(ChatAction, NamedTuple('CDelta', [('turnId', Any), ('partId', Any), ('content', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CDelta({self.turnId.VerbatimString(True)}, {self.partId.VerbatimString(True)}, {self.content.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CDelta) and self.turnId == __o.turnId and self.partId == __o.partId and self.content == __o.content
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CToolCallStart(ChatAction, NamedTuple('CToolCallStart', [('turnId', Any), ('toolCallId', Any), ('toolName', Any), ('displayName', Any), ('intention', Any), ('contributor', Any), ('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CToolCallStart({self.turnId.VerbatimString(True)}, {self.toolCallId.VerbatimString(True)}, {self.toolName.VerbatimString(True)}, {self.displayName.VerbatimString(True)}, {_dafny.string_of(self.intention)}, {_dafny.string_of(self.contributor)}, {_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CToolCallStart) and self.turnId == __o.turnId and self.toolCallId == __o.toolCallId and self.toolName == __o.toolName and self.displayName == __o.displayName and self.intention == __o.intention and self.contributor == __o.contributor and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CToolCallDelta(ChatAction, NamedTuple('CToolCallDelta', [('turnId', Any), ('toolCallId', Any), ('content', Any), ('invocationMessage', Any), ('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CToolCallDelta({self.turnId.VerbatimString(True)}, {self.toolCallId.VerbatimString(True)}, {self.content.VerbatimString(True)}, {_dafny.string_of(self.invocationMessage)}, {_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CToolCallDelta) and self.turnId == __o.turnId and self.toolCallId == __o.toolCallId and self.content == __o.content and self.invocationMessage == __o.invocationMessage and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CToolCallReady(ChatAction, NamedTuple('CToolCallReady', [('turnId', Any), ('toolCallId', Any), ('invocationMessage', Any), ('toolInput', Any), ('confirmationTitle', Any), ('riskAssessment', Any), ('edits', Any), ('editable', Any), ('options', Any), ('confirmed', Any), ('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CToolCallReady({self.turnId.VerbatimString(True)}, {self.toolCallId.VerbatimString(True)}, {_dafny.string_of(self.invocationMessage)}, {_dafny.string_of(self.toolInput)}, {_dafny.string_of(self.confirmationTitle)}, {_dafny.string_of(self.riskAssessment)}, {_dafny.string_of(self.edits)}, {_dafny.string_of(self.editable)}, {_dafny.string_of(self.options)}, {_dafny.string_of(self.confirmed)}, {_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CToolCallReady) and self.turnId == __o.turnId and self.toolCallId == __o.toolCallId and self.invocationMessage == __o.invocationMessage and self.toolInput == __o.toolInput and self.confirmationTitle == __o.confirmationTitle and self.riskAssessment == __o.riskAssessment and self.edits == __o.edits and self.editable == __o.editable and self.options == __o.options and self.confirmed == __o.confirmed and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CToolCallConfirmed(ChatAction, NamedTuple('CToolCallConfirmed', [('turnId', Any), ('toolCallId', Any), ('approved', Any), ('confirmedVal', Any), ('reason', Any), ('reasonMessage', Any), ('userSuggestion', Any), ('editedToolInput', Any), ('selectedOptionId', Any), ('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CToolCallConfirmed({self.turnId.VerbatimString(True)}, {self.toolCallId.VerbatimString(True)}, {_dafny.string_of(self.approved)}, {_dafny.string_of(self.confirmedVal)}, {_dafny.string_of(self.reason)}, {_dafny.string_of(self.reasonMessage)}, {_dafny.string_of(self.userSuggestion)}, {_dafny.string_of(self.editedToolInput)}, {_dafny.string_of(self.selectedOptionId)}, {_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CToolCallConfirmed) and self.turnId == __o.turnId and self.toolCallId == __o.toolCallId and self.approved == __o.approved and self.confirmedVal == __o.confirmedVal and self.reason == __o.reason and self.reasonMessage == __o.reasonMessage and self.userSuggestion == __o.userSuggestion and self.editedToolInput == __o.editedToolInput and self.selectedOptionId == __o.selectedOptionId and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CToolCallAuthRequired(ChatAction, NamedTuple('CToolCallAuthRequired', [('turnId', Any), ('toolCallId', Any), ('auth', Any), ('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CToolCallAuthRequired({self.turnId.VerbatimString(True)}, {self.toolCallId.VerbatimString(True)}, {_dafny.string_of(self.auth)}, {_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CToolCallAuthRequired) and self.turnId == __o.turnId and self.toolCallId == __o.toolCallId and self.auth == __o.auth and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CToolCallAuthResolved(ChatAction, NamedTuple('CToolCallAuthResolved', [('turnId', Any), ('toolCallId', Any), ('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CToolCallAuthResolved({self.turnId.VerbatimString(True)}, {self.toolCallId.VerbatimString(True)}, {_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CToolCallAuthResolved) and self.turnId == __o.turnId and self.toolCallId == __o.toolCallId and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CToolCallComplete(ChatAction, NamedTuple('CToolCallComplete', [('turnId', Any), ('toolCallId', Any), ('success', Any), ('pastTenseMessage', Any), ('resultContent', Any), ('structuredContent', Any), ('error', Any), ('requiresResultConfirmation', Any), ('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CToolCallComplete({self.turnId.VerbatimString(True)}, {self.toolCallId.VerbatimString(True)}, {_dafny.string_of(self.success)}, {_dafny.string_of(self.pastTenseMessage)}, {_dafny.string_of(self.resultContent)}, {_dafny.string_of(self.structuredContent)}, {_dafny.string_of(self.error)}, {_dafny.string_of(self.requiresResultConfirmation)}, {_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CToolCallComplete) and self.turnId == __o.turnId and self.toolCallId == __o.toolCallId and self.success == __o.success and self.pastTenseMessage == __o.pastTenseMessage and self.resultContent == __o.resultContent and self.structuredContent == __o.structuredContent and self.error == __o.error and self.requiresResultConfirmation == __o.requiresResultConfirmation and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CToolCallResultConfirmed(ChatAction, NamedTuple('CToolCallResultConfirmed', [('turnId', Any), ('toolCallId', Any), ('approved', Any), ('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CToolCallResultConfirmed({self.turnId.VerbatimString(True)}, {self.toolCallId.VerbatimString(True)}, {_dafny.string_of(self.approved)}, {_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CToolCallResultConfirmed) and self.turnId == __o.turnId and self.toolCallId == __o.toolCallId and self.approved == __o.approved and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CToolCallContentChanged(ChatAction, NamedTuple('CToolCallContentChanged', [('turnId', Any), ('toolCallId', Any), ('newContent', Any), ('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CToolCallContentChanged({self.turnId.VerbatimString(True)}, {self.toolCallId.VerbatimString(True)}, {_dafny.string_of(self.newContent)}, {_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CToolCallContentChanged) and self.turnId == __o.turnId and self.toolCallId == __o.toolCallId and self.newContent == __o.newContent and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CReasoning(ChatAction, NamedTuple('CReasoning', [('turnId', Any), ('partId', Any), ('content', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CReasoning({self.turnId.VerbatimString(True)}, {self.partId.VerbatimString(True)}, {self.content.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CReasoning) and self.turnId == __o.turnId and self.partId == __o.partId and self.content == __o.content
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CTruncated(ChatAction, NamedTuple('CTruncated', [('upToTurnId', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CTruncated({_dafny.string_of(self.upToTurnId)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CTruncated) and self.upToTurnId == __o.upToTurnId
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CQueuedReordered(ChatAction, NamedTuple('CQueuedReordered', [('order', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CQueuedReordered({_dafny.string_of(self.order)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CQueuedReordered) and self.order == __o.order
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CTurnsLoaded(ChatAction, NamedTuple('CTurnsLoaded', [('loaded', Any), ('cursor', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CTurnsLoaded({_dafny.string_of(self.loaded)}, {_dafny.string_of(self.cursor)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CTurnsLoaded) and self.loaded == __o.loaded and self.cursor == __o.cursor
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CInputRequested(ChatAction, NamedTuple('CInputRequested', [('req', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CInputRequested({_dafny.string_of(self.req)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CInputRequested) and self.req == __o.req
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CInputAnswerChanged(ChatAction, NamedTuple('CInputAnswerChanged', [('requestId', Any), ('questionId', Any), ('answer', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CInputAnswerChanged({self.requestId.VerbatimString(True)}, {self.questionId.VerbatimString(True)}, {_dafny.string_of(self.answer)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CInputAnswerChanged) and self.requestId == __o.requestId and self.questionId == __o.questionId and self.answer == __o.answer
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CInputCompleted(ChatAction, NamedTuple('CInputCompleted', [('requestId', Any), ('response', Any), ('answers', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CInputCompleted({self.requestId.VerbatimString(True)}, {self.response.VerbatimString(True)}, {_dafny.string_of(self.answers)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CInputCompleted) and self.requestId == __o.requestId and self.response == __o.response and self.answers == __o.answers
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CPendingMessageSet(ChatAction, NamedTuple('CPendingMessageSet', [('kind', Any), ('id_', Any), ('message', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CPendingMessageSet({self.kind.VerbatimString(True)}, {self.id_.VerbatimString(True)}, {_dafny.string_of(self.message)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CPendingMessageSet) and self.kind == __o.kind and self.id_ == __o.id_ and self.message == __o.message
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CPendingMessageRemoved(ChatAction, NamedTuple('CPendingMessageRemoved', [('kind', Any), ('id_', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CPendingMessageRemoved({self.kind.VerbatimString(True)}, {self.id_.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CPendingMessageRemoved) and self.kind == __o.kind and self.id_ == __o.id_
    def __hash__(self) -> int:
        return super().__hash__()

class ChatAction_CUnknown(ChatAction, NamedTuple('CUnknown', [('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'Chat.ChatAction.CUnknown({_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ChatAction_CUnknown) and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()

