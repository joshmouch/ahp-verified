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

# Module: Session

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
    def toggleEnabled(en):
        def lambda0_(d_0_en_):
            def lambda1_(d_1_e_):
                def iife0_(_pat_let102_0):
                    def iife1_(d_2_dt__update__tmp_h0_):
                        def iife2_(_pat_let103_0):
                            def iife3_(d_3_dt__update_henabled_h0_):
                                return Cust_Cust((d_2_dt__update__tmp_h0_).id_, (d_2_dt__update__tmp_h0_).kind, (d_2_dt__update__tmp_h0_).uri, (d_2_dt__update__tmp_h0_).name, d_3_dt__update_henabled_h0_, (d_2_dt__update__tmp_h0_).state, (d_2_dt__update__tmp_h0_).channel)
                            return iife3_(_pat_let103_0)
                        return iife2_(AhpSkeleton.Option_Some(d_0_en_))
                    return iife1_(_pat_let102_0)
                return iife0_(d_1_e_)

            return lambda1_

        return lambda0_(en)

    @staticmethod
    def chatPatch(st, ac, md):
        def lambda0_(d_0_md_, d_1_ac_, d_2_st_):
            def lambda1_(d_3_c_):
                def iife0_(_pat_let104_0):
                    def iife1_(d_4_dt__update__tmp_h0_):
                        def iife2_(_pat_let105_0):
                            def iife3_(d_5_dt__update_hmodifiedAt_h0_):
                                def iife4_(_pat_let106_0):
                                    def iife5_(d_6_dt__update_hactivity_h0_):
                                        def iife6_(_pat_let107_0):
                                            def iife7_(d_7_dt__update_hstatus_h0_):
                                                return Chat_Chat((d_4_dt__update__tmp_h0_).resource, (d_4_dt__update__tmp_h0_).title, d_7_dt__update_hstatus_h0_, d_6_dt__update_hactivity_h0_, d_5_dt__update_hmodifiedAt_h0_, (d_4_dt__update__tmp_h0_).origin)
                                            return iife7_(_pat_let107_0)
                                        return iife6_(((d_2_st_).value if (d_2_st_).is_Some else (d_3_c_).status))
                                    return iife5_(_pat_let106_0)
                                return iife4_((d_1_ac_ if (d_1_ac_).is_Some else (d_3_c_).activity))
                            return iife3_(_pat_let105_0)
                        return iife2_(((d_0_md_).value if (d_0_md_).is_Some else (d_3_c_).modifiedAt))
                    return iife1_(_pat_let104_0)
                return iife0_(d_3_c_)

            return lambda1_

        return lambda0_(md, ac, st)

    @staticmethod
    def mcpPatch(st, ch):
        def lambda0_(d_0_ch_, d_1_st_):
            def lambda1_(d_2_e_):
                def iife0_(_pat_let108_0):
                    def iife1_(d_3_dt__update__tmp_h0_):
                        def iife2_(_pat_let109_0):
                            def iife3_(d_4_dt__update_hchannel_h0_):
                                def iife4_(_pat_let110_0):
                                    def iife5_(d_5_dt__update_hstate_h0_):
                                        return Cust_Cust((d_3_dt__update__tmp_h0_).id_, (d_3_dt__update__tmp_h0_).kind, (d_3_dt__update__tmp_h0_).uri, (d_3_dt__update__tmp_h0_).name, (d_3_dt__update__tmp_h0_).enabled, d_5_dt__update_hstate_h0_, d_4_dt__update_hchannel_h0_)
                                    return iife5_(_pat_let110_0)
                                return iife4_(d_1_st_)
                            return iife3_(_pat_let109_0)
                        return iife2_(d_0_ch_)
                    return iife1_(_pat_let108_0)
                return iife0_(d_2_e_)

            return lambda1_

        return lambda0_(ch, st)

    @staticmethod
    def upsertCust(cs, id_, en):
        return ConfluxSeqRoute.default__.SeqUpdateById(default__.custKey, cs, id_, default__.toggleEnabled(en))

    @staticmethod
    def removeCust(cs, id_):
        return ConfluxSeqRoute.default__.SeqRemoveById(default__.custKey, cs, id_)

    @staticmethod
    def removeChat(chs, r):
        return ConfluxSeqRoute.default__.SeqRemoveById(default__.chatKey, chs, r)

    @staticmethod
    def upsertChat(chs, c):
        return ConfluxSeqRoute.default__.SeqUpsertById(default__.chatKey, chs, c)

    @staticmethod
    def updChat(chs, r, st, ac, md):
        return ConfluxSeqRoute.default__.SeqUpdateById(default__.chatKey, chs, r, default__.chatPatch(st, ac, md))

    @staticmethod
    def upsertClient(cl, c):
        return ConfluxSeqRoute.default__.SeqUpsertById(default__.clientKey, cl, c)

    @staticmethod
    def removeClient(cl, id_):
        return ConfluxSeqRoute.default__.SeqRemoveById(default__.clientKey, cl, id_)

    @staticmethod
    def upsertInput(ins, r):
        return ConfluxSeqRoute.default__.SeqUpsertById(default__.inputKey, ins, r)

    @staticmethod
    def removeInput(ins, id_):
        return ConfluxSeqRoute.default__.SeqRemoveById(default__.inputKey, ins, id_)

    @staticmethod
    def upsertCustFull(cs, c):
        return ConfluxSeqRoute.default__.SeqUpsertById(default__.custKey, cs, c)

    @staticmethod
    def setMcp(cs, id_, st, ch):
        return ConfluxSeqRoute.default__.SeqUpdateByIdWhere(default__.custKey, default__.mcpPred, cs, id_, default__.mcpPatch(st, ch))

    @staticmethod
    def isMcp(cs, id_):
        def lambda0_(exists_var_0_):
            d_0_i_: int = exists_var_0_
            return ((((0) <= (d_0_i_)) and ((d_0_i_) < (len(cs)))) and ((((cs)[d_0_i_]).id_) == (id_))) and ((((cs)[d_0_i_]).kind) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcpServer"))))

        return _dafny.quantifier(_dafny.IntegerRange(0, len(cs)), False, lambda0_)

    @staticmethod
    def ApplyToSession(s, a, now):
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
        source0_ = a
        if True:
            if source0_.is_Ready:
                def iife0_(_pat_let111_0):
                    def iife1_(d_0_dt__update__tmp_h0_):
                        def iife2_(_pat_let112_0):
                            def iife3_(d_1_dt__update_hlifecycle_h0_):
                                return SessionState_SessionState((d_0_dt__update__tmp_h0_).provider, (d_0_dt__update__tmp_h0_).title, (d_0_dt__update__tmp_h0_).status, d_1_dt__update_hlifecycle_h0_, (d_0_dt__update__tmp_h0_).activity, (d_0_dt__update__tmp_h0_).config, (d_0_dt__update__tmp_h0_).meta, (d_0_dt__update__tmp_h0_).creationError, (d_0_dt__update__tmp_h0_).serverTools, (d_0_dt__update__tmp_h0_).changesets, (d_0_dt__update__tmp_h0_).canvases, (d_0_dt__update__tmp_h0_).openCanvases, (d_0_dt__update__tmp_h0_).defaultChat, (d_0_dt__update__tmp_h0_).activeClients, (d_0_dt__update__tmp_h0_).chats, (d_0_dt__update__tmp_h0_).customizations, (d_0_dt__update__tmp_h0_).inputNeeded)
                            return iife3_(_pat_let112_0)
                        return iife2_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")))
                    return iife1_(_pat_let111_0)
                return (iife0_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CreationFailed:
                d_2_e_ = source0_.error
                def iife4_(_pat_let113_0):
                    def iife5_(d_3_dt__update__tmp_h1_):
                        def iife6_(_pat_let114_0):
                            def iife7_(d_4_dt__update_hcreationError_h0_):
                                def iife8_(_pat_let115_0):
                                    def iife9_(d_5_dt__update_hlifecycle_h1_):
                                        return SessionState_SessionState((d_3_dt__update__tmp_h1_).provider, (d_3_dt__update__tmp_h1_).title, (d_3_dt__update__tmp_h1_).status, d_5_dt__update_hlifecycle_h1_, (d_3_dt__update__tmp_h1_).activity, (d_3_dt__update__tmp_h1_).config, (d_3_dt__update__tmp_h1_).meta, d_4_dt__update_hcreationError_h0_, (d_3_dt__update__tmp_h1_).serverTools, (d_3_dt__update__tmp_h1_).changesets, (d_3_dt__update__tmp_h1_).canvases, (d_3_dt__update__tmp_h1_).openCanvases, (d_3_dt__update__tmp_h1_).defaultChat, (d_3_dt__update__tmp_h1_).activeClients, (d_3_dt__update__tmp_h1_).chats, (d_3_dt__update__tmp_h1_).customizations, (d_3_dt__update__tmp_h1_).inputNeeded)
                                    return iife9_(_pat_let115_0)
                                return iife8_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "creationFailed")))
                            return iife7_(_pat_let114_0)
                        return iife6_((AhpSkeleton.Option_None() if (d_2_e_) == (ConfluxCodec.Json_JNull()) else AhpSkeleton.Option_Some(d_2_e_)))
                    return iife5_(_pat_let113_0)
                return (iife4_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_TitleChanged:
                d_6_t_ = source0_.title
                def iife10_(_pat_let116_0):
                    def iife11_(d_7_dt__update__tmp_h2_):
                        def iife12_(_pat_let117_0):
                            def iife13_(d_8_dt__update_htitle_h0_):
                                return SessionState_SessionState((d_7_dt__update__tmp_h2_).provider, d_8_dt__update_htitle_h0_, (d_7_dt__update__tmp_h2_).status, (d_7_dt__update__tmp_h2_).lifecycle, (d_7_dt__update__tmp_h2_).activity, (d_7_dt__update__tmp_h2_).config, (d_7_dt__update__tmp_h2_).meta, (d_7_dt__update__tmp_h2_).creationError, (d_7_dt__update__tmp_h2_).serverTools, (d_7_dt__update__tmp_h2_).changesets, (d_7_dt__update__tmp_h2_).canvases, (d_7_dt__update__tmp_h2_).openCanvases, (d_7_dt__update__tmp_h2_).defaultChat, (d_7_dt__update__tmp_h2_).activeClients, (d_7_dt__update__tmp_h2_).chats, (d_7_dt__update__tmp_h2_).customizations, (d_7_dt__update__tmp_h2_).inputNeeded)
                            return iife13_(_pat_let117_0)
                        return iife12_(d_6_t_)
                    return iife11_(_pat_let116_0)
                return (iife10_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_ServerToolsChanged:
                d_9_t_ = source0_.tools
                def iife14_(_pat_let118_0):
                    def iife15_(d_10_dt__update__tmp_h3_):
                        def iife16_(_pat_let119_0):
                            def iife17_(d_11_dt__update_hserverTools_h0_):
                                return SessionState_SessionState((d_10_dt__update__tmp_h3_).provider, (d_10_dt__update__tmp_h3_).title, (d_10_dt__update__tmp_h3_).status, (d_10_dt__update__tmp_h3_).lifecycle, (d_10_dt__update__tmp_h3_).activity, (d_10_dt__update__tmp_h3_).config, (d_10_dt__update__tmp_h3_).meta, (d_10_dt__update__tmp_h3_).creationError, d_11_dt__update_hserverTools_h0_, (d_10_dt__update__tmp_h3_).changesets, (d_10_dt__update__tmp_h3_).canvases, (d_10_dt__update__tmp_h3_).openCanvases, (d_10_dt__update__tmp_h3_).defaultChat, (d_10_dt__update__tmp_h3_).activeClients, (d_10_dt__update__tmp_h3_).chats, (d_10_dt__update__tmp_h3_).customizations, (d_10_dt__update__tmp_h3_).inputNeeded)
                            return iife17_(_pat_let119_0)
                        return iife16_((AhpSkeleton.Option_None() if (d_9_t_) == (ConfluxCodec.Json_JNull()) else AhpSkeleton.Option_Some(d_9_t_)))
                    return iife15_(_pat_let118_0)
                return (iife14_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_MetaChanged:
                d_12_m_ = source0_.meta
                def iife18_(_pat_let120_0):
                    def iife19_(d_13_dt__update__tmp_h4_):
                        def iife20_(_pat_let121_0):
                            def iife21_(d_14_dt__update_hmeta_h0_):
                                return SessionState_SessionState((d_13_dt__update__tmp_h4_).provider, (d_13_dt__update__tmp_h4_).title, (d_13_dt__update__tmp_h4_).status, (d_13_dt__update__tmp_h4_).lifecycle, (d_13_dt__update__tmp_h4_).activity, (d_13_dt__update__tmp_h4_).config, d_14_dt__update_hmeta_h0_, (d_13_dt__update__tmp_h4_).creationError, (d_13_dt__update__tmp_h4_).serverTools, (d_13_dt__update__tmp_h4_).changesets, (d_13_dt__update__tmp_h4_).canvases, (d_13_dt__update__tmp_h4_).openCanvases, (d_13_dt__update__tmp_h4_).defaultChat, (d_13_dt__update__tmp_h4_).activeClients, (d_13_dt__update__tmp_h4_).chats, (d_13_dt__update__tmp_h4_).customizations, (d_13_dt__update__tmp_h4_).inputNeeded)
                            return iife21_(_pat_let121_0)
                        return iife20_((AhpSkeleton.Option_None() if (d_12_m_) == (ConfluxCodec.Json_JNull()) else AhpSkeleton.Option_Some(d_12_m_)))
                    return iife19_(_pat_let120_0)
                return (iife18_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_IsReadChanged:
                d_15_v_ = source0_.isRead
                def iife22_(_pat_let122_0):
                    def iife23_(d_16_dt__update__tmp_h5_):
                        def iife24_(_pat_let123_0):
                            def iife25_(d_17_dt__update_hstatus_h0_):
                                return SessionState_SessionState((d_16_dt__update__tmp_h5_).provider, (d_16_dt__update__tmp_h5_).title, d_17_dt__update_hstatus_h0_, (d_16_dt__update__tmp_h5_).lifecycle, (d_16_dt__update__tmp_h5_).activity, (d_16_dt__update__tmp_h5_).config, (d_16_dt__update__tmp_h5_).meta, (d_16_dt__update__tmp_h5_).creationError, (d_16_dt__update__tmp_h5_).serverTools, (d_16_dt__update__tmp_h5_).changesets, (d_16_dt__update__tmp_h5_).canvases, (d_16_dt__update__tmp_h5_).openCanvases, (d_16_dt__update__tmp_h5_).defaultChat, (d_16_dt__update__tmp_h5_).activeClients, (d_16_dt__update__tmp_h5_).chats, (d_16_dt__update__tmp_h5_).customizations, (d_16_dt__update__tmp_h5_).inputNeeded)
                            return iife25_(_pat_let123_0)
                        return iife24_((default__.setBit((pat_let_tv0_).status, default__.B__READ) if d_15_v_ else default__.clearBit((pat_let_tv1_).status, default__.B__READ)))
                    return iife23_(_pat_let122_0)
                return (iife22_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_IsArchivedChanged:
                d_18_v_ = source0_.isArchived
                def iife26_(_pat_let124_0):
                    def iife27_(d_19_dt__update__tmp_h6_):
                        def iife28_(_pat_let125_0):
                            def iife29_(d_20_dt__update_hstatus_h1_):
                                return SessionState_SessionState((d_19_dt__update__tmp_h6_).provider, (d_19_dt__update__tmp_h6_).title, d_20_dt__update_hstatus_h1_, (d_19_dt__update__tmp_h6_).lifecycle, (d_19_dt__update__tmp_h6_).activity, (d_19_dt__update__tmp_h6_).config, (d_19_dt__update__tmp_h6_).meta, (d_19_dt__update__tmp_h6_).creationError, (d_19_dt__update__tmp_h6_).serverTools, (d_19_dt__update__tmp_h6_).changesets, (d_19_dt__update__tmp_h6_).canvases, (d_19_dt__update__tmp_h6_).openCanvases, (d_19_dt__update__tmp_h6_).defaultChat, (d_19_dt__update__tmp_h6_).activeClients, (d_19_dt__update__tmp_h6_).chats, (d_19_dt__update__tmp_h6_).customizations, (d_19_dt__update__tmp_h6_).inputNeeded)
                            return iife29_(_pat_let125_0)
                        return iife28_((default__.setBit((pat_let_tv2_).status, default__.B__ARCH) if d_18_v_ else default__.clearBit((pat_let_tv3_).status, default__.B__ARCH)))
                    return iife27_(_pat_let124_0)
                return (iife26_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_ActivityChanged:
                d_21_ac_ = source0_.activity
                def iife30_(_pat_let126_0):
                    def iife31_(d_22_dt__update__tmp_h7_):
                        def iife32_(_pat_let127_0):
                            def iife33_(d_23_dt__update_hactivity_h0_):
                                return SessionState_SessionState((d_22_dt__update__tmp_h7_).provider, (d_22_dt__update__tmp_h7_).title, (d_22_dt__update__tmp_h7_).status, (d_22_dt__update__tmp_h7_).lifecycle, d_23_dt__update_hactivity_h0_, (d_22_dt__update__tmp_h7_).config, (d_22_dt__update__tmp_h7_).meta, (d_22_dt__update__tmp_h7_).creationError, (d_22_dt__update__tmp_h7_).serverTools, (d_22_dt__update__tmp_h7_).changesets, (d_22_dt__update__tmp_h7_).canvases, (d_22_dt__update__tmp_h7_).openCanvases, (d_22_dt__update__tmp_h7_).defaultChat, (d_22_dt__update__tmp_h7_).activeClients, (d_22_dt__update__tmp_h7_).chats, (d_22_dt__update__tmp_h7_).customizations, (d_22_dt__update__tmp_h7_).inputNeeded)
                            return iife33_(_pat_let127_0)
                        return iife32_(d_21_ac_)
                    return iife31_(_pat_let126_0)
                return (iife30_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_ConfigChanged:
                d_24_cfg_ = source0_.config
                d_25_repl_ = source0_.replace
                source1_ = (s).config
                if True:
                    if source1_.is_None:
                        return (s, AhpSkeleton.ReduceOutcome_NoOp())
                if True:
                    d_26_c_ = source1_.value
                    def iife34_(_pat_let128_0):
                        def iife35_(d_27_dt__update__tmp_h8_):
                            def iife36_(_pat_let129_0):
                                def iife37_(d_28_dt__update_hconfig_h0_):
                                    return SessionState_SessionState((d_27_dt__update__tmp_h8_).provider, (d_27_dt__update__tmp_h8_).title, (d_27_dt__update__tmp_h8_).status, (d_27_dt__update__tmp_h8_).lifecycle, (d_27_dt__update__tmp_h8_).activity, d_28_dt__update_hconfig_h0_, (d_27_dt__update__tmp_h8_).meta, (d_27_dt__update__tmp_h8_).creationError, (d_27_dt__update__tmp_h8_).serverTools, (d_27_dt__update__tmp_h8_).changesets, (d_27_dt__update__tmp_h8_).canvases, (d_27_dt__update__tmp_h8_).openCanvases, (d_27_dt__update__tmp_h8_).defaultChat, (d_27_dt__update__tmp_h8_).activeClients, (d_27_dt__update__tmp_h8_).chats, (d_27_dt__update__tmp_h8_).customizations, (d_27_dt__update__tmp_h8_).inputNeeded)
                                return iife37_(_pat_let129_0)
                            return iife36_(AhpSkeleton.Option_Some(SConfig_SConfig((d_26_c_).schema, (d_24_cfg_ if d_25_repl_ else ((d_26_c_).values) | (d_24_cfg_)))))
                        return iife35_(_pat_let128_0)
                    return (iife34_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CustomizationsChanged:
                d_29_cs_ = source0_.customizations
                def iife38_(_pat_let130_0):
                    def iife39_(d_30_dt__update__tmp_h9_):
                        def iife40_(_pat_let131_0):
                            def iife41_(d_31_dt__update_hcustomizations_h0_):
                                return SessionState_SessionState((d_30_dt__update__tmp_h9_).provider, (d_30_dt__update__tmp_h9_).title, (d_30_dt__update__tmp_h9_).status, (d_30_dt__update__tmp_h9_).lifecycle, (d_30_dt__update__tmp_h9_).activity, (d_30_dt__update__tmp_h9_).config, (d_30_dt__update__tmp_h9_).meta, (d_30_dt__update__tmp_h9_).creationError, (d_30_dt__update__tmp_h9_).serverTools, (d_30_dt__update__tmp_h9_).changesets, (d_30_dt__update__tmp_h9_).canvases, (d_30_dt__update__tmp_h9_).openCanvases, (d_30_dt__update__tmp_h9_).defaultChat, (d_30_dt__update__tmp_h9_).activeClients, (d_30_dt__update__tmp_h9_).chats, d_31_dt__update_hcustomizations_h0_, (d_30_dt__update__tmp_h9_).inputNeeded)
                            return iife41_(_pat_let131_0)
                        return iife40_(d_29_cs_)
                    return iife39_(_pat_let130_0)
                return (iife38_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CustomizationToggled:
                d_32_id_ = source0_.id_
                d_33_en_ = source0_.enabled
                def iife42_(_pat_let132_0):
                    def iife43_(d_34_dt__update__tmp_h10_):
                        def iife44_(_pat_let133_0):
                            def iife45_(d_35_dt__update_hcustomizations_h1_):
                                return SessionState_SessionState((d_34_dt__update__tmp_h10_).provider, (d_34_dt__update__tmp_h10_).title, (d_34_dt__update__tmp_h10_).status, (d_34_dt__update__tmp_h10_).lifecycle, (d_34_dt__update__tmp_h10_).activity, (d_34_dt__update__tmp_h10_).config, (d_34_dt__update__tmp_h10_).meta, (d_34_dt__update__tmp_h10_).creationError, (d_34_dt__update__tmp_h10_).serverTools, (d_34_dt__update__tmp_h10_).changesets, (d_34_dt__update__tmp_h10_).canvases, (d_34_dt__update__tmp_h10_).openCanvases, (d_34_dt__update__tmp_h10_).defaultChat, (d_34_dt__update__tmp_h10_).activeClients, (d_34_dt__update__tmp_h10_).chats, d_35_dt__update_hcustomizations_h1_, (d_34_dt__update__tmp_h10_).inputNeeded)
                            return iife45_(_pat_let133_0)
                        return iife44_(default__.upsertCust((pat_let_tv4_).customizations, d_32_id_, d_33_en_))
                    return iife43_(_pat_let132_0)
                return (iife42_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CustomizationRemoved:
                d_36_id_ = source0_.id_
                def iife46_(_pat_let134_0):
                    def iife47_(d_37_dt__update__tmp_h11_):
                        def iife48_(_pat_let135_0):
                            def iife49_(d_38_dt__update_hcustomizations_h2_):
                                return SessionState_SessionState((d_37_dt__update__tmp_h11_).provider, (d_37_dt__update__tmp_h11_).title, (d_37_dt__update__tmp_h11_).status, (d_37_dt__update__tmp_h11_).lifecycle, (d_37_dt__update__tmp_h11_).activity, (d_37_dt__update__tmp_h11_).config, (d_37_dt__update__tmp_h11_).meta, (d_37_dt__update__tmp_h11_).creationError, (d_37_dt__update__tmp_h11_).serverTools, (d_37_dt__update__tmp_h11_).changesets, (d_37_dt__update__tmp_h11_).canvases, (d_37_dt__update__tmp_h11_).openCanvases, (d_37_dt__update__tmp_h11_).defaultChat, (d_37_dt__update__tmp_h11_).activeClients, (d_37_dt__update__tmp_h11_).chats, d_38_dt__update_hcustomizations_h2_, (d_37_dt__update__tmp_h11_).inputNeeded)
                            return iife49_(_pat_let135_0)
                        return iife48_(default__.removeCust((pat_let_tv5_).customizations, d_36_id_))
                    return iife47_(_pat_let134_0)
                return (iife46_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_DefaultChatChanged:
                d_39_c_ = source0_.chat
                def iife50_(_pat_let136_0):
                    def iife51_(d_40_dt__update__tmp_h12_):
                        def iife52_(_pat_let137_0):
                            def iife53_(d_41_dt__update_hdefaultChat_h0_):
                                return SessionState_SessionState((d_40_dt__update__tmp_h12_).provider, (d_40_dt__update__tmp_h12_).title, (d_40_dt__update__tmp_h12_).status, (d_40_dt__update__tmp_h12_).lifecycle, (d_40_dt__update__tmp_h12_).activity, (d_40_dt__update__tmp_h12_).config, (d_40_dt__update__tmp_h12_).meta, (d_40_dt__update__tmp_h12_).creationError, (d_40_dt__update__tmp_h12_).serverTools, (d_40_dt__update__tmp_h12_).changesets, (d_40_dt__update__tmp_h12_).canvases, (d_40_dt__update__tmp_h12_).openCanvases, d_41_dt__update_hdefaultChat_h0_, (d_40_dt__update__tmp_h12_).activeClients, (d_40_dt__update__tmp_h12_).chats, (d_40_dt__update__tmp_h12_).customizations, (d_40_dt__update__tmp_h12_).inputNeeded)
                            return iife53_(_pat_let137_0)
                        return iife52_(d_39_c_)
                    return iife51_(_pat_let136_0)
                return (iife50_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_ChatAdded:
                d_42_c_ = source0_.summary
                def iife54_(_pat_let138_0):
                    def iife55_(d_43_dt__update__tmp_h13_):
                        def iife56_(_pat_let139_0):
                            def iife57_(d_44_dt__update_hchats_h0_):
                                return SessionState_SessionState((d_43_dt__update__tmp_h13_).provider, (d_43_dt__update__tmp_h13_).title, (d_43_dt__update__tmp_h13_).status, (d_43_dt__update__tmp_h13_).lifecycle, (d_43_dt__update__tmp_h13_).activity, (d_43_dt__update__tmp_h13_).config, (d_43_dt__update__tmp_h13_).meta, (d_43_dt__update__tmp_h13_).creationError, (d_43_dt__update__tmp_h13_).serverTools, (d_43_dt__update__tmp_h13_).changesets, (d_43_dt__update__tmp_h13_).canvases, (d_43_dt__update__tmp_h13_).openCanvases, (d_43_dt__update__tmp_h13_).defaultChat, (d_43_dt__update__tmp_h13_).activeClients, d_44_dt__update_hchats_h0_, (d_43_dt__update__tmp_h13_).customizations, (d_43_dt__update__tmp_h13_).inputNeeded)
                            return iife57_(_pat_let139_0)
                        return iife56_(default__.upsertChat((pat_let_tv6_).chats, d_42_c_))
                    return iife55_(_pat_let138_0)
                return (iife54_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_ChatRemoved:
                d_45_r_ = source0_.resource
                def iife58_(_pat_let140_0):
                    def iife59_(d_46_dt__update__tmp_h14_):
                        def iife60_(_pat_let141_0):
                            def iife61_(d_47_dt__update_hdefaultChat_h1_):
                                def iife62_(_pat_let142_0):
                                    def iife63_(d_48_dt__update_hchats_h1_):
                                        return SessionState_SessionState((d_46_dt__update__tmp_h14_).provider, (d_46_dt__update__tmp_h14_).title, (d_46_dt__update__tmp_h14_).status, (d_46_dt__update__tmp_h14_).lifecycle, (d_46_dt__update__tmp_h14_).activity, (d_46_dt__update__tmp_h14_).config, (d_46_dt__update__tmp_h14_).meta, (d_46_dt__update__tmp_h14_).creationError, (d_46_dt__update__tmp_h14_).serverTools, (d_46_dt__update__tmp_h14_).changesets, (d_46_dt__update__tmp_h14_).canvases, (d_46_dt__update__tmp_h14_).openCanvases, d_47_dt__update_hdefaultChat_h1_, (d_46_dt__update__tmp_h14_).activeClients, d_48_dt__update_hchats_h1_, (d_46_dt__update__tmp_h14_).customizations, (d_46_dt__update__tmp_h14_).inputNeeded)
                                    return iife63_(_pat_let142_0)
                                return iife62_(default__.removeChat((pat_let_tv9_).chats, d_45_r_))
                            return iife61_(_pat_let141_0)
                        return iife60_((AhpSkeleton.Option_None() if ((pat_let_tv7_).defaultChat) == (AhpSkeleton.Option_Some(d_45_r_)) else (pat_let_tv8_).defaultChat))
                    return iife59_(_pat_let140_0)
                return (iife58_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_ChatUpdated:
                d_49_r_ = source0_.resource
                d_50_st_ = source0_.status
                d_51_ac_ = source0_.activity
                d_52_md_ = source0_.modifiedAt
                def iife64_(_pat_let143_0):
                    def iife65_(d_53_dt__update__tmp_h15_):
                        def iife66_(_pat_let144_0):
                            def iife67_(d_54_dt__update_hchats_h2_):
                                return SessionState_SessionState((d_53_dt__update__tmp_h15_).provider, (d_53_dt__update__tmp_h15_).title, (d_53_dt__update__tmp_h15_).status, (d_53_dt__update__tmp_h15_).lifecycle, (d_53_dt__update__tmp_h15_).activity, (d_53_dt__update__tmp_h15_).config, (d_53_dt__update__tmp_h15_).meta, (d_53_dt__update__tmp_h15_).creationError, (d_53_dt__update__tmp_h15_).serverTools, (d_53_dt__update__tmp_h15_).changesets, (d_53_dt__update__tmp_h15_).canvases, (d_53_dt__update__tmp_h15_).openCanvases, (d_53_dt__update__tmp_h15_).defaultChat, (d_53_dt__update__tmp_h15_).activeClients, d_54_dt__update_hchats_h2_, (d_53_dt__update__tmp_h15_).customizations, (d_53_dt__update__tmp_h15_).inputNeeded)
                            return iife67_(_pat_let144_0)
                        return iife66_(default__.updChat((pat_let_tv10_).chats, d_49_r_, d_50_st_, d_51_ac_, d_52_md_))
                    return iife65_(_pat_let143_0)
                return (iife64_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_ChangesetsChanged:
                d_55_cs_ = source0_.changesets
                def iife68_(_pat_let145_0):
                    def iife69_(d_56_dt__update__tmp_h16_):
                        def iife70_(_pat_let146_0):
                            def iife71_(d_57_dt__update_hchangesets_h0_):
                                return SessionState_SessionState((d_56_dt__update__tmp_h16_).provider, (d_56_dt__update__tmp_h16_).title, (d_56_dt__update__tmp_h16_).status, (d_56_dt__update__tmp_h16_).lifecycle, (d_56_dt__update__tmp_h16_).activity, (d_56_dt__update__tmp_h16_).config, (d_56_dt__update__tmp_h16_).meta, (d_56_dt__update__tmp_h16_).creationError, (d_56_dt__update__tmp_h16_).serverTools, d_57_dt__update_hchangesets_h0_, (d_56_dt__update__tmp_h16_).canvases, (d_56_dt__update__tmp_h16_).openCanvases, (d_56_dt__update__tmp_h16_).defaultChat, (d_56_dt__update__tmp_h16_).activeClients, (d_56_dt__update__tmp_h16_).chats, (d_56_dt__update__tmp_h16_).customizations, (d_56_dt__update__tmp_h16_).inputNeeded)
                            return iife71_(_pat_let146_0)
                        return iife70_(d_55_cs_)
                    return iife69_(_pat_let145_0)
                return (iife68_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CanvasesChanged:
                d_58_cs_ = source0_.canvases
                def iife72_(_pat_let147_0):
                    def iife73_(d_59_dt__update__tmp_h17_):
                        def iife74_(_pat_let148_0):
                            def iife75_(d_60_dt__update_hcanvases_h0_):
                                return SessionState_SessionState((d_59_dt__update__tmp_h17_).provider, (d_59_dt__update__tmp_h17_).title, (d_59_dt__update__tmp_h17_).status, (d_59_dt__update__tmp_h17_).lifecycle, (d_59_dt__update__tmp_h17_).activity, (d_59_dt__update__tmp_h17_).config, (d_59_dt__update__tmp_h17_).meta, (d_59_dt__update__tmp_h17_).creationError, (d_59_dt__update__tmp_h17_).serverTools, (d_59_dt__update__tmp_h17_).changesets, d_60_dt__update_hcanvases_h0_, (d_59_dt__update__tmp_h17_).openCanvases, (d_59_dt__update__tmp_h17_).defaultChat, (d_59_dt__update__tmp_h17_).activeClients, (d_59_dt__update__tmp_h17_).chats, (d_59_dt__update__tmp_h17_).customizations, (d_59_dt__update__tmp_h17_).inputNeeded)
                            return iife75_(_pat_let148_0)
                        return iife74_((AhpSkeleton.Option_None() if (d_58_cs_) == (ConfluxCodec.Json_JNull()) else AhpSkeleton.Option_Some(d_58_cs_)))
                    return iife73_(_pat_let147_0)
                return (iife72_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_OpenCanvasesChanged:
                d_61_cs_ = source0_.openCanvases
                def iife76_(_pat_let149_0):
                    def iife77_(d_62_dt__update__tmp_h18_):
                        def iife78_(_pat_let150_0):
                            def iife79_(d_63_dt__update_hopenCanvases_h0_):
                                return SessionState_SessionState((d_62_dt__update__tmp_h18_).provider, (d_62_dt__update__tmp_h18_).title, (d_62_dt__update__tmp_h18_).status, (d_62_dt__update__tmp_h18_).lifecycle, (d_62_dt__update__tmp_h18_).activity, (d_62_dt__update__tmp_h18_).config, (d_62_dt__update__tmp_h18_).meta, (d_62_dt__update__tmp_h18_).creationError, (d_62_dt__update__tmp_h18_).serverTools, (d_62_dt__update__tmp_h18_).changesets, (d_62_dt__update__tmp_h18_).canvases, d_63_dt__update_hopenCanvases_h0_, (d_62_dt__update__tmp_h18_).defaultChat, (d_62_dt__update__tmp_h18_).activeClients, (d_62_dt__update__tmp_h18_).chats, (d_62_dt__update__tmp_h18_).customizations, (d_62_dt__update__tmp_h18_).inputNeeded)
                            return iife79_(_pat_let150_0)
                        return iife78_((AhpSkeleton.Option_None() if (d_61_cs_) == (ConfluxCodec.Json_JNull()) else AhpSkeleton.Option_Some(d_61_cs_)))
                    return iife77_(_pat_let149_0)
                return (iife76_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_ActiveClientSet:
                d_64_c_ = source0_.client
                def iife80_(_pat_let151_0):
                    def iife81_(d_65_dt__update__tmp_h19_):
                        def iife82_(_pat_let152_0):
                            def iife83_(d_66_dt__update_hactiveClients_h0_):
                                return SessionState_SessionState((d_65_dt__update__tmp_h19_).provider, (d_65_dt__update__tmp_h19_).title, (d_65_dt__update__tmp_h19_).status, (d_65_dt__update__tmp_h19_).lifecycle, (d_65_dt__update__tmp_h19_).activity, (d_65_dt__update__tmp_h19_).config, (d_65_dt__update__tmp_h19_).meta, (d_65_dt__update__tmp_h19_).creationError, (d_65_dt__update__tmp_h19_).serverTools, (d_65_dt__update__tmp_h19_).changesets, (d_65_dt__update__tmp_h19_).canvases, (d_65_dt__update__tmp_h19_).openCanvases, (d_65_dt__update__tmp_h19_).defaultChat, d_66_dt__update_hactiveClients_h0_, (d_65_dt__update__tmp_h19_).chats, (d_65_dt__update__tmp_h19_).customizations, (d_65_dt__update__tmp_h19_).inputNeeded)
                            return iife83_(_pat_let152_0)
                        return iife82_(default__.upsertClient((pat_let_tv11_).activeClients, d_64_c_))
                    return iife81_(_pat_let151_0)
                return (iife80_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_ActiveClientRemoved:
                d_67_id_ = source0_.clientId
                def iife84_(_pat_let153_0):
                    def iife85_(d_68_dt__update__tmp_h20_):
                        def iife86_(_pat_let154_0):
                            def iife87_(d_69_dt__update_hactiveClients_h1_):
                                return SessionState_SessionState((d_68_dt__update__tmp_h20_).provider, (d_68_dt__update__tmp_h20_).title, (d_68_dt__update__tmp_h20_).status, (d_68_dt__update__tmp_h20_).lifecycle, (d_68_dt__update__tmp_h20_).activity, (d_68_dt__update__tmp_h20_).config, (d_68_dt__update__tmp_h20_).meta, (d_68_dt__update__tmp_h20_).creationError, (d_68_dt__update__tmp_h20_).serverTools, (d_68_dt__update__tmp_h20_).changesets, (d_68_dt__update__tmp_h20_).canvases, (d_68_dt__update__tmp_h20_).openCanvases, (d_68_dt__update__tmp_h20_).defaultChat, d_69_dt__update_hactiveClients_h1_, (d_68_dt__update__tmp_h20_).chats, (d_68_dt__update__tmp_h20_).customizations, (d_68_dt__update__tmp_h20_).inputNeeded)
                            return iife87_(_pat_let154_0)
                        return iife86_(default__.removeClient((pat_let_tv12_).activeClients, d_67_id_))
                    return iife85_(_pat_let153_0)
                return (iife84_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_InputNeededSet:
                d_70_r_ = source0_.req
                def iife88_(_pat_let155_0):
                    def iife89_(d_71_dt__update__tmp_h21_):
                        def iife90_(_pat_let156_0):
                            def iife91_(d_72_dt__update_hstatus_h2_):
                                def iife92_(_pat_let157_0):
                                    def iife93_(d_73_dt__update_hinputNeeded_h0_):
                                        return SessionState_SessionState((d_71_dt__update__tmp_h21_).provider, (d_71_dt__update__tmp_h21_).title, d_72_dt__update_hstatus_h2_, (d_71_dt__update__tmp_h21_).lifecycle, (d_71_dt__update__tmp_h21_).activity, (d_71_dt__update__tmp_h21_).config, (d_71_dt__update__tmp_h21_).meta, (d_71_dt__update__tmp_h21_).creationError, (d_71_dt__update__tmp_h21_).serverTools, (d_71_dt__update__tmp_h21_).changesets, (d_71_dt__update__tmp_h21_).canvases, (d_71_dt__update__tmp_h21_).openCanvases, (d_71_dt__update__tmp_h21_).defaultChat, (d_71_dt__update__tmp_h21_).activeClients, (d_71_dt__update__tmp_h21_).chats, (d_71_dt__update__tmp_h21_).customizations, d_73_dt__update_hinputNeeded_h0_)
                                    return iife93_(_pat_let157_0)
                                return iife92_(default__.upsertInput((pat_let_tv14_).inputNeeded, d_70_r_))
                            return iife91_(_pat_let156_0)
                        return iife90_(default__.setBit((pat_let_tv13_).status, default__.B__INPUT))
                    return iife89_(_pat_let155_0)
                return (iife88_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_InputNeededRemoved:
                d_74_id_ = source0_.id_
                d_75_rem_ = default__.removeInput((s).inputNeeded, d_74_id_)
                def iife94_(_pat_let158_0):
                    def iife95_(d_76_dt__update__tmp_h22_):
                        def iife96_(_pat_let159_0):
                            def iife97_(d_77_dt__update_hstatus_h3_):
                                def iife98_(_pat_let160_0):
                                    def iife99_(d_78_dt__update_hinputNeeded_h1_):
                                        return SessionState_SessionState((d_76_dt__update__tmp_h22_).provider, (d_76_dt__update__tmp_h22_).title, d_77_dt__update_hstatus_h3_, (d_76_dt__update__tmp_h22_).lifecycle, (d_76_dt__update__tmp_h22_).activity, (d_76_dt__update__tmp_h22_).config, (d_76_dt__update__tmp_h22_).meta, (d_76_dt__update__tmp_h22_).creationError, (d_76_dt__update__tmp_h22_).serverTools, (d_76_dt__update__tmp_h22_).changesets, (d_76_dt__update__tmp_h22_).canvases, (d_76_dt__update__tmp_h22_).openCanvases, (d_76_dt__update__tmp_h22_).defaultChat, (d_76_dt__update__tmp_h22_).activeClients, (d_76_dt__update__tmp_h22_).chats, (d_76_dt__update__tmp_h22_).customizations, d_78_dt__update_hinputNeeded_h1_)
                                    return iife99_(_pat_let160_0)
                                return iife98_(d_75_rem_)
                            return iife97_(_pat_let159_0)
                        return iife96_((default__.clearBit((pat_let_tv15_).status, default__.B__INPUT) if (len(d_75_rem_)) == (0) else (pat_let_tv16_).status))
                    return iife95_(_pat_let158_0)
                return (iife94_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_CustomizationUpdated:
                d_79_c_ = source0_.customization
                def iife100_(_pat_let161_0):
                    def iife101_(d_80_dt__update__tmp_h23_):
                        def iife102_(_pat_let162_0):
                            def iife103_(d_81_dt__update_hcustomizations_h3_):
                                return SessionState_SessionState((d_80_dt__update__tmp_h23_).provider, (d_80_dt__update__tmp_h23_).title, (d_80_dt__update__tmp_h23_).status, (d_80_dt__update__tmp_h23_).lifecycle, (d_80_dt__update__tmp_h23_).activity, (d_80_dt__update__tmp_h23_).config, (d_80_dt__update__tmp_h23_).meta, (d_80_dt__update__tmp_h23_).creationError, (d_80_dt__update__tmp_h23_).serverTools, (d_80_dt__update__tmp_h23_).changesets, (d_80_dt__update__tmp_h23_).canvases, (d_80_dt__update__tmp_h23_).openCanvases, (d_80_dt__update__tmp_h23_).defaultChat, (d_80_dt__update__tmp_h23_).activeClients, (d_80_dt__update__tmp_h23_).chats, d_81_dt__update_hcustomizations_h3_, (d_80_dt__update__tmp_h23_).inputNeeded)
                            return iife103_(_pat_let162_0)
                        return iife102_(default__.upsertCustFull((pat_let_tv17_).customizations, d_79_c_))
                    return iife101_(_pat_let161_0)
                return (iife100_(s), AhpSkeleton.ReduceOutcome_Applied())
        if True:
            if source0_.is_McpServerStateChanged:
                d_82_id_ = source0_.id_
                d_83_st_ = source0_.state
                d_84_ch_ = source0_.channel
                if default__.isMcp((s).customizations, d_82_id_):
                    def iife104_(_pat_let163_0):
                        def iife105_(d_85_dt__update__tmp_h24_):
                            def iife106_(_pat_let164_0):
                                def iife107_(d_86_dt__update_hcustomizations_h4_):
                                    return SessionState_SessionState((d_85_dt__update__tmp_h24_).provider, (d_85_dt__update__tmp_h24_).title, (d_85_dt__update__tmp_h24_).status, (d_85_dt__update__tmp_h24_).lifecycle, (d_85_dt__update__tmp_h24_).activity, (d_85_dt__update__tmp_h24_).config, (d_85_dt__update__tmp_h24_).meta, (d_85_dt__update__tmp_h24_).creationError, (d_85_dt__update__tmp_h24_).serverTools, (d_85_dt__update__tmp_h24_).changesets, (d_85_dt__update__tmp_h24_).canvases, (d_85_dt__update__tmp_h24_).openCanvases, (d_85_dt__update__tmp_h24_).defaultChat, (d_85_dt__update__tmp_h24_).activeClients, (d_85_dt__update__tmp_h24_).chats, d_86_dt__update_hcustomizations_h4_, (d_85_dt__update__tmp_h24_).inputNeeded)
                                return iife107_(_pat_let164_0)
                            return iife106_(default__.setMcp((pat_let_tv18_).customizations, d_82_id_, (AhpSkeleton.Option_None() if (d_83_st_) == (ConfluxCodec.Json_JNull()) else AhpSkeleton.Option_Some(d_83_st_)), d_84_ch_))
                        return iife105_(_pat_let163_0)
                    return (iife104_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_McpServerStartRequested:
                d_87_id_ = source0_.id_
                if default__.isMcp((s).customizations, d_87_id_):
                    def iife108_(_pat_let165_0):
                        def iife109_(d_88_dt__update__tmp_h25_):
                            def iife110_(_pat_let166_0):
                                def iife111_(d_89_dt__update_hcustomizations_h5_):
                                    return SessionState_SessionState((d_88_dt__update__tmp_h25_).provider, (d_88_dt__update__tmp_h25_).title, (d_88_dt__update__tmp_h25_).status, (d_88_dt__update__tmp_h25_).lifecycle, (d_88_dt__update__tmp_h25_).activity, (d_88_dt__update__tmp_h25_).config, (d_88_dt__update__tmp_h25_).meta, (d_88_dt__update__tmp_h25_).creationError, (d_88_dt__update__tmp_h25_).serverTools, (d_88_dt__update__tmp_h25_).changesets, (d_88_dt__update__tmp_h25_).canvases, (d_88_dt__update__tmp_h25_).openCanvases, (d_88_dt__update__tmp_h25_).defaultChat, (d_88_dt__update__tmp_h25_).activeClients, (d_88_dt__update__tmp_h25_).chats, d_89_dt__update_hcustomizations_h5_, (d_88_dt__update__tmp_h25_).inputNeeded)
                                return iife111_(_pat_let166_0)
                            return iife110_(default__.setMcp((pat_let_tv19_).customizations, d_87_id_, AhpSkeleton.Option_Some(ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "starting")))}))), AhpSkeleton.Option_None()))
                        return iife109_(_pat_let165_0)
                    return (iife108_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            if source0_.is_McpServerStopRequested:
                d_90_id_ = source0_.id_
                if default__.isMcp((s).customizations, d_90_id_):
                    def iife112_(_pat_let167_0):
                        def iife113_(d_91_dt__update__tmp_h26_):
                            def iife114_(_pat_let168_0):
                                def iife115_(d_92_dt__update_hcustomizations_h6_):
                                    return SessionState_SessionState((d_91_dt__update__tmp_h26_).provider, (d_91_dt__update__tmp_h26_).title, (d_91_dt__update__tmp_h26_).status, (d_91_dt__update__tmp_h26_).lifecycle, (d_91_dt__update__tmp_h26_).activity, (d_91_dt__update__tmp_h26_).config, (d_91_dt__update__tmp_h26_).meta, (d_91_dt__update__tmp_h26_).creationError, (d_91_dt__update__tmp_h26_).serverTools, (d_91_dt__update__tmp_h26_).changesets, (d_91_dt__update__tmp_h26_).canvases, (d_91_dt__update__tmp_h26_).openCanvases, (d_91_dt__update__tmp_h26_).defaultChat, (d_91_dt__update__tmp_h26_).activeClients, (d_91_dt__update__tmp_h26_).chats, d_92_dt__update_hcustomizations_h6_, (d_91_dt__update__tmp_h26_).inputNeeded)
                                return iife115_(_pat_let168_0)
                            return iife114_(default__.setMcp((pat_let_tv20_).customizations, d_90_id_, AhpSkeleton.Option_Some(ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stopped")))}))), AhpSkeleton.Option_None()))
                        return iife113_(_pat_let167_0)
                    return (iife112_(s), AhpSkeleton.ReduceOutcome_Applied())
                elif True:
                    return (s, AhpSkeleton.ReduceOutcome_NoOp())
        if True:
            return (s, AhpSkeleton.ReduceOutcome_NoOp())

    @staticmethod
    def WfCust(c):
        return (not (((c).state).is_Some) or ((((c).state).value) != (ConfluxCodec.Json_JNull()))) and (not (((c).channel).is_Some) or ((((c).channel).value) != (ConfluxCodec.Json_JNull())))

    @staticmethod
    def WfCusts(cs):
        def lambda0_(forall_var_0_):
            d_0_i_: int = forall_var_0_
            return not (((0) <= (d_0_i_)) and ((d_0_i_) < (len(cs)))) or (default__.WfCust((cs)[d_0_i_]))

        return _dafny.quantifier(_dafny.IntegerRange(0, len(cs)), True, lambda0_)

    @staticmethod
    def SessionWf(s):
        return ((((((((((not (((s).meta).is_Some) or ((((s).meta).value) != (ConfluxCodec.Json_JNull()))) and (not (((s).creationError).is_Some) or ((((s).creationError).value) != (ConfluxCodec.Json_JNull())))) and (not (((s).serverTools).is_Some) or ((((s).serverTools).value) != (ConfluxCodec.Json_JNull())))) and (not (((s).changesets).is_Some) or ((((s).changesets).value) != (ConfluxCodec.Json_JNull())))) and (not (((s).canvases).is_Some) or ((((s).canvases).value) != (ConfluxCodec.Json_JNull())))) and (not (((s).openCanvases).is_Some) or ((((s).openCanvases).value) != (ConfluxCodec.Json_JNull())))) and (default__.WfCusts((s).customizations))) and (ConfluxSeqRoute.default__.UniqueKeys(default__.custKey, (s).customizations))) and (ConfluxSeqRoute.default__.UniqueKeys(default__.chatKey, (s).chats))) and (ConfluxSeqRoute.default__.UniqueKeys(default__.clientKey, (s).activeClients))) and (ConfluxSeqRoute.default__.UniqueKeys(default__.inputKey, (s).inputNeeded))

    @staticmethod
    def WfSessionActionInv(a):
        source0_ = a
        if True:
            if source0_.is_CustomizationsChanged:
                d_0_cs_ = source0_.customizations
                return (default__.WfCusts(d_0_cs_)) and (ConfluxSeqRoute.default__.UniqueKeys(default__.custKey, d_0_cs_))
        if True:
            if source0_.is_CustomizationUpdated:
                d_1_c_ = source0_.customization
                return default__.WfCust(d_1_c_)
        if True:
            if source0_.is_ChangesetsChanged:
                d_2_cs_ = source0_.changesets
                return not ((d_2_cs_).is_Some) or (((d_2_cs_).value) != (ConfluxCodec.Json_JNull()))
        if True:
            if source0_.is_McpServerStateChanged:
                d_3_ch_ = source0_.channel
                return not ((d_3_ch_).is_Some) or (((d_3_ch_).value) != (ConfluxCodec.Json_JNull()))
        if True:
            return True

    @staticmethod
    def apply1(s, a):
        return (default__.ApplyToSession(s, a, 9999))[0]

    @staticmethod
    def fold(s, acts):
        return ConfluxContract.default__.Fold(default__.apply1, s, acts)

    @staticmethod
    def S0():
        return SessionState_SessionState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "copilot")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Test Session")), 1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "creating")), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), AhpSkeleton.Option_None(), _dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference([]), _dafny.SeqWithoutIsStrInference([]))

    @staticmethod
    def RunCorpus():
        pass_: int = int(0)
        modeled: int = int(0)
        modeled = 36
        pass_ = 0
        d_0_s_: SessionState
        d_0_s_ = default__.S0()
        def iife0_(_pat_let169_0):
            def iife1_(d_1_dt__update__tmp_h0_):
                def iife2_(_pat_let170_0):
                    def iife3_(d_2_dt__update_hlifecycle_h0_):
                        return SessionState_SessionState((d_1_dt__update__tmp_h0_).provider, (d_1_dt__update__tmp_h0_).title, (d_1_dt__update__tmp_h0_).status, d_2_dt__update_hlifecycle_h0_, (d_1_dt__update__tmp_h0_).activity, (d_1_dt__update__tmp_h0_).config, (d_1_dt__update__tmp_h0_).meta, (d_1_dt__update__tmp_h0_).creationError, (d_1_dt__update__tmp_h0_).serverTools, (d_1_dt__update__tmp_h0_).changesets, (d_1_dt__update__tmp_h0_).canvases, (d_1_dt__update__tmp_h0_).openCanvases, (d_1_dt__update__tmp_h0_).defaultChat, (d_1_dt__update__tmp_h0_).activeClients, (d_1_dt__update__tmp_h0_).chats, (d_1_dt__update__tmp_h0_).customizations, (d_1_dt__update__tmp_h0_).inputNeeded)
                    return iife3_(_pat_let170_0)
                return iife2_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "creating")))
            return iife1_(_pat_let169_0)
        def iife4_(_pat_let171_0):
            def iife5_(d_3_dt__update__tmp_h1_):
                def iife6_(_pat_let172_0):
                    def iife7_(d_4_dt__update_hlifecycle_h1_):
                        return SessionState_SessionState((d_3_dt__update__tmp_h1_).provider, (d_3_dt__update__tmp_h1_).title, (d_3_dt__update__tmp_h1_).status, d_4_dt__update_hlifecycle_h1_, (d_3_dt__update__tmp_h1_).activity, (d_3_dt__update__tmp_h1_).config, (d_3_dt__update__tmp_h1_).meta, (d_3_dt__update__tmp_h1_).creationError, (d_3_dt__update__tmp_h1_).serverTools, (d_3_dt__update__tmp_h1_).changesets, (d_3_dt__update__tmp_h1_).canvases, (d_3_dt__update__tmp_h1_).openCanvases, (d_3_dt__update__tmp_h1_).defaultChat, (d_3_dt__update__tmp_h1_).activeClients, (d_3_dt__update__tmp_h1_).chats, (d_3_dt__update__tmp_h1_).customizations, (d_3_dt__update__tmp_h1_).inputNeeded)
                    return iife7_(_pat_let172_0)
                return iife6_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")))
            return iife5_(_pat_let171_0)
        if (default__.apply1(iife0_(d_0_s_), SessionAction_Ready())) == (iife4_(d_0_s_)):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_s_, SessionAction_CreationFailed(ConfluxCodec.Json_JNull()))).lifecycle) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "creationFailed"))):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_s_, SessionAction_TitleChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "New"))))).title) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "New"))):
            pass_ = (pass_) + (1)
        d_5_tools_: ConfluxCodec.Json
        d_5_tools_ = ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "name")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bash")))}))]))
        if ((default__.apply1(d_0_s_, SessionAction_ServerToolsChanged(d_5_tools_))).serverTools) == (AhpSkeleton.Option_Some(d_5_tools_)):
            pass_ = (pass_) + (1)
        d_6_m_: ConfluxCodec.Json
        d_6_m_ = ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "git")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "feature")))}))
        def iife8_(_pat_let173_0):
            def iife9_(d_7_dt__update__tmp_h2_):
                def iife10_(_pat_let174_0):
                    def iife11_(d_8_dt__update_hmeta_h0_):
                        return SessionState_SessionState((d_7_dt__update__tmp_h2_).provider, (d_7_dt__update__tmp_h2_).title, (d_7_dt__update__tmp_h2_).status, (d_7_dt__update__tmp_h2_).lifecycle, (d_7_dt__update__tmp_h2_).activity, (d_7_dt__update__tmp_h2_).config, d_8_dt__update_hmeta_h0_, (d_7_dt__update__tmp_h2_).creationError, (d_7_dt__update__tmp_h2_).serverTools, (d_7_dt__update__tmp_h2_).changesets, (d_7_dt__update__tmp_h2_).canvases, (d_7_dt__update__tmp_h2_).openCanvases, (d_7_dt__update__tmp_h2_).defaultChat, (d_7_dt__update__tmp_h2_).activeClients, (d_7_dt__update__tmp_h2_).chats, (d_7_dt__update__tmp_h2_).customizations, (d_7_dt__update__tmp_h2_).inputNeeded)
                    return iife11_(_pat_let174_0)
                return iife10_(AhpSkeleton.Option_Some(ConfluxCodec.Json_JNull()))
            return iife9_(_pat_let173_0)
        if ((default__.apply1(iife8_(d_0_s_), SessionAction_MetaChanged(d_6_m_))).meta) == (AhpSkeleton.Option_Some(d_6_m_)):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_s_, SessionAction_IsReadChanged(True))).status) == (33):
            pass_ = (pass_) + (1)
        def iife12_(_pat_let175_0):
            def iife13_(d_9_dt__update__tmp_h3_):
                def iife14_(_pat_let176_0):
                    def iife15_(d_10_dt__update_hstatus_h0_):
                        return SessionState_SessionState((d_9_dt__update__tmp_h3_).provider, (d_9_dt__update__tmp_h3_).title, d_10_dt__update_hstatus_h0_, (d_9_dt__update__tmp_h3_).lifecycle, (d_9_dt__update__tmp_h3_).activity, (d_9_dt__update__tmp_h3_).config, (d_9_dt__update__tmp_h3_).meta, (d_9_dt__update__tmp_h3_).creationError, (d_9_dt__update__tmp_h3_).serverTools, (d_9_dt__update__tmp_h3_).changesets, (d_9_dt__update__tmp_h3_).canvases, (d_9_dt__update__tmp_h3_).openCanvases, (d_9_dt__update__tmp_h3_).defaultChat, (d_9_dt__update__tmp_h3_).activeClients, (d_9_dt__update__tmp_h3_).chats, (d_9_dt__update__tmp_h3_).customizations, (d_9_dt__update__tmp_h3_).inputNeeded)
                    return iife15_(_pat_let176_0)
                return iife14_(33)
            return iife13_(_pat_let175_0)
        if ((default__.apply1(iife12_(d_0_s_), SessionAction_IsReadChanged(False))).status) == (1):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_s_, SessionAction_IsArchivedChanged(True))).status) == (65):
            pass_ = (pass_) + (1)
        def iife16_(_pat_let177_0):
            def iife17_(d_11_dt__update__tmp_h4_):
                def iife18_(_pat_let178_0):
                    def iife19_(d_12_dt__update_hstatus_h1_):
                        return SessionState_SessionState((d_11_dt__update__tmp_h4_).provider, (d_11_dt__update__tmp_h4_).title, d_12_dt__update_hstatus_h1_, (d_11_dt__update__tmp_h4_).lifecycle, (d_11_dt__update__tmp_h4_).activity, (d_11_dt__update__tmp_h4_).config, (d_11_dt__update__tmp_h4_).meta, (d_11_dt__update__tmp_h4_).creationError, (d_11_dt__update__tmp_h4_).serverTools, (d_11_dt__update__tmp_h4_).changesets, (d_11_dt__update__tmp_h4_).canvases, (d_11_dt__update__tmp_h4_).openCanvases, (d_11_dt__update__tmp_h4_).defaultChat, (d_11_dt__update__tmp_h4_).activeClients, (d_11_dt__update__tmp_h4_).chats, (d_11_dt__update__tmp_h4_).customizations, (d_11_dt__update__tmp_h4_).inputNeeded)
                    return iife19_(_pat_let178_0)
                return iife18_(65)
            return iife17_(_pat_let177_0)
        if ((default__.apply1(iife16_(d_0_s_), SessionAction_IsArchivedChanged(False))).status) == (1):
            pass_ = (pass_) + (1)
        def iife20_(_pat_let179_0):
            def iife21_(d_13_dt__update__tmp_h5_):
                def iife22_(_pat_let180_0):
                    def iife23_(d_14_dt__update_hactivity_h0_):
                        return SessionState_SessionState((d_13_dt__update__tmp_h5_).provider, (d_13_dt__update__tmp_h5_).title, (d_13_dt__update__tmp_h5_).status, (d_13_dt__update__tmp_h5_).lifecycle, d_14_dt__update_hactivity_h0_, (d_13_dt__update__tmp_h5_).config, (d_13_dt__update__tmp_h5_).meta, (d_13_dt__update__tmp_h5_).creationError, (d_13_dt__update__tmp_h5_).serverTools, (d_13_dt__update__tmp_h5_).changesets, (d_13_dt__update__tmp_h5_).canvases, (d_13_dt__update__tmp_h5_).openCanvases, (d_13_dt__update__tmp_h5_).defaultChat, (d_13_dt__update__tmp_h5_).activeClients, (d_13_dt__update__tmp_h5_).chats, (d_13_dt__update__tmp_h5_).customizations, (d_13_dt__update__tmp_h5_).inputNeeded)
                    return iife23_(_pat_let180_0)
                return iife22_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Running"))))
            return iife21_(_pat_let179_0)
        if (default__.apply1(d_0_s_, SessionAction_ActivityChanged(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Running")))))) == (iife20_(d_0_s_)):
            pass_ = (pass_) + (1)
        def iife24_(_pat_let181_0):
            def iife25_(d_15_dt__update__tmp_h6_):
                def iife26_(_pat_let182_0):
                    def iife27_(d_16_dt__update_hactivity_h1_):
                        return SessionState_SessionState((d_15_dt__update__tmp_h6_).provider, (d_15_dt__update__tmp_h6_).title, (d_15_dt__update__tmp_h6_).status, (d_15_dt__update__tmp_h6_).lifecycle, d_16_dt__update_hactivity_h1_, (d_15_dt__update__tmp_h6_).config, (d_15_dt__update__tmp_h6_).meta, (d_15_dt__update__tmp_h6_).creationError, (d_15_dt__update__tmp_h6_).serverTools, (d_15_dt__update__tmp_h6_).changesets, (d_15_dt__update__tmp_h6_).canvases, (d_15_dt__update__tmp_h6_).openCanvases, (d_15_dt__update__tmp_h6_).defaultChat, (d_15_dt__update__tmp_h6_).activeClients, (d_15_dt__update__tmp_h6_).chats, (d_15_dt__update__tmp_h6_).customizations, (d_15_dt__update__tmp_h6_).inputNeeded)
                    return iife27_(_pat_let182_0)
                return iife26_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))))
            return iife25_(_pat_let181_0)
        if (default__.apply1(iife24_(d_0_s_), SessionAction_ActivityChanged(AhpSkeleton.Option_None()))) == (d_0_s_):
            pass_ = (pass_) + (1)
        d_17_cfg_: SConfig
        d_17_cfg_ = SConfig_SConfig(ConfluxCodec.Json_JNull(), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "target")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "worktree"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baseBranch")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "main")))}))
        d_18_cfgExp_: SConfig
        d_18_cfgExp_ = SConfig_SConfig(ConfluxCodec.Json_JNull(), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "target")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "worktree"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baseBranch")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "develop")))}))
        pat_let_tv0_ = d_17_cfg_
        def iife28_(_pat_let183_0):
            def iife29_(d_19_dt__update__tmp_h7_):
                def iife30_(_pat_let184_0):
                    def iife31_(d_20_dt__update_hconfig_h0_):
                        return SessionState_SessionState((d_19_dt__update__tmp_h7_).provider, (d_19_dt__update__tmp_h7_).title, (d_19_dt__update__tmp_h7_).status, (d_19_dt__update__tmp_h7_).lifecycle, (d_19_dt__update__tmp_h7_).activity, d_20_dt__update_hconfig_h0_, (d_19_dt__update__tmp_h7_).meta, (d_19_dt__update__tmp_h7_).creationError, (d_19_dt__update__tmp_h7_).serverTools, (d_19_dt__update__tmp_h7_).changesets, (d_19_dt__update__tmp_h7_).canvases, (d_19_dt__update__tmp_h7_).openCanvases, (d_19_dt__update__tmp_h7_).defaultChat, (d_19_dt__update__tmp_h7_).activeClients, (d_19_dt__update__tmp_h7_).chats, (d_19_dt__update__tmp_h7_).customizations, (d_19_dt__update__tmp_h7_).inputNeeded)
                    return iife31_(_pat_let184_0)
                return iife30_(AhpSkeleton.Option_Some(pat_let_tv0_))
            return iife29_(_pat_let183_0)
        if ((default__.apply1(iife28_(d_0_s_), SessionAction_ConfigChanged(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baseBranch")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "develop")))}), False))).config) == (AhpSkeleton.Option_Some(d_18_cfgExp_)):
            pass_ = (pass_) + (1)
        d_21_cfgRepl_: SConfig
        d_21_cfgRepl_ = SConfig_SConfig(ConfluxCodec.Json_JNull(), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baseBranch")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "develop")))}))
        pat_let_tv1_ = d_17_cfg_
        def iife32_(_pat_let185_0):
            def iife33_(d_22_dt__update__tmp_h8_):
                def iife34_(_pat_let186_0):
                    def iife35_(d_23_dt__update_hconfig_h1_):
                        return SessionState_SessionState((d_22_dt__update__tmp_h8_).provider, (d_22_dt__update__tmp_h8_).title, (d_22_dt__update__tmp_h8_).status, (d_22_dt__update__tmp_h8_).lifecycle, (d_22_dt__update__tmp_h8_).activity, d_23_dt__update_hconfig_h1_, (d_22_dt__update__tmp_h8_).meta, (d_22_dt__update__tmp_h8_).creationError, (d_22_dt__update__tmp_h8_).serverTools, (d_22_dt__update__tmp_h8_).changesets, (d_22_dt__update__tmp_h8_).canvases, (d_22_dt__update__tmp_h8_).openCanvases, (d_22_dt__update__tmp_h8_).defaultChat, (d_22_dt__update__tmp_h8_).activeClients, (d_22_dt__update__tmp_h8_).chats, (d_22_dt__update__tmp_h8_).customizations, (d_22_dt__update__tmp_h8_).inputNeeded)
                    return iife35_(_pat_let186_0)
                return iife34_(AhpSkeleton.Option_Some(pat_let_tv1_))
            return iife33_(_pat_let185_0)
        if ((default__.apply1(iife32_(d_0_s_), SessionAction_ConfigChanged(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baseBranch")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "develop")))}), True))).config) == (AhpSkeleton.Option_Some(d_21_cfgRepl_)):
            pass_ = (pass_) + (1)
        if (default__.apply1(d_0_s_, SessionAction_ConfigChanged(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")): ConfluxCodec.Json_JNull()}), False))) == (d_0_s_):
            pass_ = (pass_) + (1)
        d_24_pa_: Cust
        d_24_pa_ = Cust_Cust(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plugin-a")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plugin")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "https://plugins.example/a")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Plugin A")), AhpSkeleton.Option_Some(True), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())
        d_25_pb_: Cust
        d_25_pb_ = Cust_Cust(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plugin-b")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plugin")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "https://plugins.example/b")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Plugin B")), AhpSkeleton.Option_Some(True), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())
        if ((default__.apply1(d_0_s_, SessionAction_CustomizationsChanged(_dafny.SeqWithoutIsStrInference([d_24_pa_, d_25_pb_])))).customizations) == (_dafny.SeqWithoutIsStrInference([d_24_pa_, d_25_pb_])):
            pass_ = (pass_) + (1)
        pat_let_tv2_ = d_24_pa_
        pat_let_tv3_ = d_25_pb_
        def iife36_(_pat_let187_0):
            def iife37_(d_26_dt__update__tmp_h9_):
                def iife38_(_pat_let188_0):
                    def iife39_(d_27_dt__update_hcustomizations_h0_):
                        return SessionState_SessionState((d_26_dt__update__tmp_h9_).provider, (d_26_dt__update__tmp_h9_).title, (d_26_dt__update__tmp_h9_).status, (d_26_dt__update__tmp_h9_).lifecycle, (d_26_dt__update__tmp_h9_).activity, (d_26_dt__update__tmp_h9_).config, (d_26_dt__update__tmp_h9_).meta, (d_26_dt__update__tmp_h9_).creationError, (d_26_dt__update__tmp_h9_).serverTools, (d_26_dt__update__tmp_h9_).changesets, (d_26_dt__update__tmp_h9_).canvases, (d_26_dt__update__tmp_h9_).openCanvases, (d_26_dt__update__tmp_h9_).defaultChat, (d_26_dt__update__tmp_h9_).activeClients, (d_26_dt__update__tmp_h9_).chats, d_27_dt__update_hcustomizations_h0_, (d_26_dt__update__tmp_h9_).inputNeeded)
                    return iife39_(_pat_let188_0)
                return iife38_(_dafny.SeqWithoutIsStrInference([pat_let_tv2_, pat_let_tv3_]))
            return iife37_(_pat_let187_0)
        def iife40_(_pat_let189_0):
            def iife41_(d_28_dt__update__tmp_h10_):
                def iife42_(_pat_let190_0):
                    def iife43_(d_29_dt__update_henabled_h0_):
                        return Cust_Cust((d_28_dt__update__tmp_h10_).id_, (d_28_dt__update__tmp_h10_).kind, (d_28_dt__update__tmp_h10_).uri, (d_28_dt__update__tmp_h10_).name, d_29_dt__update_henabled_h0_, (d_28_dt__update__tmp_h10_).state, (d_28_dt__update__tmp_h10_).channel)
                    return iife43_(_pat_let190_0)
                return iife42_(AhpSkeleton.Option_Some(False))
            return iife41_(_pat_let189_0)
        if ((default__.apply1(iife36_(d_0_s_), SessionAction_CustomizationToggled(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plugin-a")), False))).customizations) == (_dafny.SeqWithoutIsStrInference([iife40_(d_24_pa_), d_25_pb_])):
            pass_ = (pass_) + (1)
        pat_let_tv4_ = d_24_pa_
        pat_let_tv5_ = d_25_pb_
        def iife44_(_pat_let191_0):
            def iife45_(d_30_dt__update__tmp_h11_):
                def iife46_(_pat_let192_0):
                    def iife47_(d_31_dt__update_hcustomizations_h1_):
                        return SessionState_SessionState((d_30_dt__update__tmp_h11_).provider, (d_30_dt__update__tmp_h11_).title, (d_30_dt__update__tmp_h11_).status, (d_30_dt__update__tmp_h11_).lifecycle, (d_30_dt__update__tmp_h11_).activity, (d_30_dt__update__tmp_h11_).config, (d_30_dt__update__tmp_h11_).meta, (d_30_dt__update__tmp_h11_).creationError, (d_30_dt__update__tmp_h11_).serverTools, (d_30_dt__update__tmp_h11_).changesets, (d_30_dt__update__tmp_h11_).canvases, (d_30_dt__update__tmp_h11_).openCanvases, (d_30_dt__update__tmp_h11_).defaultChat, (d_30_dt__update__tmp_h11_).activeClients, (d_30_dt__update__tmp_h11_).chats, d_31_dt__update_hcustomizations_h1_, (d_30_dt__update__tmp_h11_).inputNeeded)
                    return iife47_(_pat_let192_0)
                return iife46_(_dafny.SeqWithoutIsStrInference([pat_let_tv4_, pat_let_tv5_]))
            return iife45_(_pat_let191_0)
        if ((default__.apply1(iife44_(d_0_s_), SessionAction_CustomizationRemoved(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plugin-a"))))).customizations) == (_dafny.SeqWithoutIsStrInference([d_25_pb_])):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_s_, SessionAction_DefaultChatChanged(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c1")))))).defaultChat) == (AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c1")))):
            pass_ = (pass_) + (1)
        def iife48_(_pat_let193_0):
            def iife49_(d_32_dt__update__tmp_h12_):
                def iife50_(_pat_let194_0):
                    def iife51_(d_33_dt__update_hdefaultChat_h0_):
                        return SessionState_SessionState((d_32_dt__update__tmp_h12_).provider, (d_32_dt__update__tmp_h12_).title, (d_32_dt__update__tmp_h12_).status, (d_32_dt__update__tmp_h12_).lifecycle, (d_32_dt__update__tmp_h12_).activity, (d_32_dt__update__tmp_h12_).config, (d_32_dt__update__tmp_h12_).meta, (d_32_dt__update__tmp_h12_).creationError, (d_32_dt__update__tmp_h12_).serverTools, (d_32_dt__update__tmp_h12_).changesets, (d_32_dt__update__tmp_h12_).canvases, (d_32_dt__update__tmp_h12_).openCanvases, d_33_dt__update_hdefaultChat_h0_, (d_32_dt__update__tmp_h12_).activeClients, (d_32_dt__update__tmp_h12_).chats, (d_32_dt__update__tmp_h12_).customizations, (d_32_dt__update__tmp_h12_).inputNeeded)
                    return iife51_(_pat_let194_0)
                return iife50_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahp-chat:/old"))))
            return iife49_(_pat_let193_0)
        if ((default__.apply1(iife48_(d_0_s_), SessionAction_DefaultChatChanged(AhpSkeleton.Option_None()))).defaultChat) == (AhpSkeleton.Option_None()):
            pass_ = (pass_) + (1)
        d_34_c1_: Chat
        d_34_c1_ = Chat_Chat(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahp-chat:/c1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Chat 1")), 1, AhpSkeleton.Option_None(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t1")), ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "user")))})))
        if ((default__.apply1(d_0_s_, SessionAction_ChatAdded(d_34_c1_))).chats) == (_dafny.SeqWithoutIsStrInference([d_34_c1_])):
            pass_ = (pass_) + (1)
        d_35_c1r_: Chat
        d_36_dt__update__tmp_h13_ = d_34_c1_
        d_37_dt__update_hmodifiedAt_h0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t2"))
        d_38_dt__update_hstatus_h2_ = 8
        d_39_dt__update_htitle_h0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Chat 1 (renamed)"))
        d_35_c1r_ = Chat_Chat((d_36_dt__update__tmp_h13_).resource, d_39_dt__update_htitle_h0_, d_38_dt__update_hstatus_h2_, (d_36_dt__update__tmp_h13_).activity, d_37_dt__update_hmodifiedAt_h0_, (d_36_dt__update__tmp_h13_).origin)
        pat_let_tv6_ = d_34_c1_
        def iife52_(_pat_let195_0):
            def iife53_(d_40_dt__update__tmp_h14_):
                def iife54_(_pat_let196_0):
                    def iife55_(d_41_dt__update_hchats_h0_):
                        return SessionState_SessionState((d_40_dt__update__tmp_h14_).provider, (d_40_dt__update__tmp_h14_).title, (d_40_dt__update__tmp_h14_).status, (d_40_dt__update__tmp_h14_).lifecycle, (d_40_dt__update__tmp_h14_).activity, (d_40_dt__update__tmp_h14_).config, (d_40_dt__update__tmp_h14_).meta, (d_40_dt__update__tmp_h14_).creationError, (d_40_dt__update__tmp_h14_).serverTools, (d_40_dt__update__tmp_h14_).changesets, (d_40_dt__update__tmp_h14_).canvases, (d_40_dt__update__tmp_h14_).openCanvases, (d_40_dt__update__tmp_h14_).defaultChat, (d_40_dt__update__tmp_h14_).activeClients, d_41_dt__update_hchats_h0_, (d_40_dt__update__tmp_h14_).customizations, (d_40_dt__update__tmp_h14_).inputNeeded)
                    return iife55_(_pat_let196_0)
                return iife54_(_dafny.SeqWithoutIsStrInference([pat_let_tv6_]))
            return iife53_(_pat_let195_0)
        if ((default__.apply1(iife52_(d_0_s_), SessionAction_ChatAdded(d_35_c1r_))).chats) == (_dafny.SeqWithoutIsStrInference([d_35_c1r_])):
            pass_ = (pass_) + (1)
        d_42_c2_: Chat
        d_42_c2_ = Chat_Chat(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahp-chat:/c2")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Chat 2")), 8, AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Thinking"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t2")), ConfluxCodec.Json_JNull())
        pat_let_tv7_ = d_34_c1_
        pat_let_tv8_ = d_42_c2_
        pat_let_tv9_ = d_42_c2_
        def iife56_(_pat_let197_0):
            def iife57_(d_43_dt__update__tmp_h15_):
                def iife58_(_pat_let198_0):
                    def iife59_(d_44_dt__update_hdefaultChat_h1_):
                        def iife60_(_pat_let199_0):
                            def iife61_(d_45_dt__update_hchats_h1_):
                                return SessionState_SessionState((d_43_dt__update__tmp_h15_).provider, (d_43_dt__update__tmp_h15_).title, (d_43_dt__update__tmp_h15_).status, (d_43_dt__update__tmp_h15_).lifecycle, (d_43_dt__update__tmp_h15_).activity, (d_43_dt__update__tmp_h15_).config, (d_43_dt__update__tmp_h15_).meta, (d_43_dt__update__tmp_h15_).creationError, (d_43_dt__update__tmp_h15_).serverTools, (d_43_dt__update__tmp_h15_).changesets, (d_43_dt__update__tmp_h15_).canvases, (d_43_dt__update__tmp_h15_).openCanvases, d_44_dt__update_hdefaultChat_h1_, (d_43_dt__update__tmp_h15_).activeClients, d_45_dt__update_hchats_h1_, (d_43_dt__update__tmp_h15_).customizations, (d_43_dt__update__tmp_h15_).inputNeeded)
                            return iife61_(_pat_let199_0)
                        return iife60_(_dafny.SeqWithoutIsStrInference([pat_let_tv7_, pat_let_tv8_]))
                    return iife59_(_pat_let198_0)
                return iife58_(AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahp-chat:/c1"))))
            return iife57_(_pat_let197_0)
        def iife62_(_pat_let200_0):
            def iife63_(d_46_dt__update__tmp_h16_):
                def iife64_(_pat_let201_0):
                    def iife65_(d_47_dt__update_hdefaultChat_h2_):
                        def iife66_(_pat_let202_0):
                            def iife67_(d_48_dt__update_hchats_h2_):
                                return SessionState_SessionState((d_46_dt__update__tmp_h16_).provider, (d_46_dt__update__tmp_h16_).title, (d_46_dt__update__tmp_h16_).status, (d_46_dt__update__tmp_h16_).lifecycle, (d_46_dt__update__tmp_h16_).activity, (d_46_dt__update__tmp_h16_).config, (d_46_dt__update__tmp_h16_).meta, (d_46_dt__update__tmp_h16_).creationError, (d_46_dt__update__tmp_h16_).serverTools, (d_46_dt__update__tmp_h16_).changesets, (d_46_dt__update__tmp_h16_).canvases, (d_46_dt__update__tmp_h16_).openCanvases, d_47_dt__update_hdefaultChat_h2_, (d_46_dt__update__tmp_h16_).activeClients, d_48_dt__update_hchats_h2_, (d_46_dt__update__tmp_h16_).customizations, (d_46_dt__update__tmp_h16_).inputNeeded)
                            return iife67_(_pat_let202_0)
                        return iife66_(_dafny.SeqWithoutIsStrInference([pat_let_tv9_]))
                    return iife65_(_pat_let201_0)
                return iife64_(AhpSkeleton.Option_None())
            return iife63_(_pat_let200_0)
        if (default__.apply1(iife56_(d_0_s_), SessionAction_ChatRemoved(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahp-chat:/c1"))))) == (iife62_(d_0_s_)):
            pass_ = (pass_) + (1)
        d_49_c1u_: Chat
        d_50_dt__update__tmp_h17_ = d_34_c1_
        d_51_dt__update_hmodifiedAt_h1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t2"))
        d_52_dt__update_hactivity_h2_ = AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Waiting for approval")))
        d_53_dt__update_hstatus_h3_ = 24
        d_49_c1u_ = Chat_Chat((d_50_dt__update__tmp_h17_).resource, (d_50_dt__update__tmp_h17_).title, d_53_dt__update_hstatus_h3_, d_52_dt__update_hactivity_h2_, d_51_dt__update_hmodifiedAt_h1_, (d_50_dt__update__tmp_h17_).origin)
        pat_let_tv10_ = d_34_c1_
        def iife68_(_pat_let203_0):
            def iife69_(d_54_dt__update__tmp_h18_):
                def iife70_(_pat_let204_0):
                    def iife71_(d_55_dt__update_hchats_h3_):
                        return SessionState_SessionState((d_54_dt__update__tmp_h18_).provider, (d_54_dt__update__tmp_h18_).title, (d_54_dt__update__tmp_h18_).status, (d_54_dt__update__tmp_h18_).lifecycle, (d_54_dt__update__tmp_h18_).activity, (d_54_dt__update__tmp_h18_).config, (d_54_dt__update__tmp_h18_).meta, (d_54_dt__update__tmp_h18_).creationError, (d_54_dt__update__tmp_h18_).serverTools, (d_54_dt__update__tmp_h18_).changesets, (d_54_dt__update__tmp_h18_).canvases, (d_54_dt__update__tmp_h18_).openCanvases, (d_54_dt__update__tmp_h18_).defaultChat, (d_54_dt__update__tmp_h18_).activeClients, d_55_dt__update_hchats_h3_, (d_54_dt__update__tmp_h18_).customizations, (d_54_dt__update__tmp_h18_).inputNeeded)
                    return iife71_(_pat_let204_0)
                return iife70_(_dafny.SeqWithoutIsStrInference([pat_let_tv10_]))
            return iife69_(_pat_let203_0)
        if ((default__.apply1(iife68_(d_0_s_), SessionAction_ChatUpdated(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahp-chat:/c1")), AhpSkeleton.Option_Some(24), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Waiting for approval"))), AhpSkeleton.Option_Some(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t2")))))).chats) == (_dafny.SeqWithoutIsStrInference([d_49_c1u_])):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_s_, SessionAction_ChangesetsChanged(AhpSkeleton.Option_Some(ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([])))))).changesets) == (AhpSkeleton.Option_Some(ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([])))):
            pass_ = (pass_) + (1)
        def iife72_(_pat_let205_0):
            def iife73_(d_56_dt__update__tmp_h19_):
                def iife74_(_pat_let206_0):
                    def iife75_(d_57_dt__update_hchangesets_h0_):
                        return SessionState_SessionState((d_56_dt__update__tmp_h19_).provider, (d_56_dt__update__tmp_h19_).title, (d_56_dt__update__tmp_h19_).status, (d_56_dt__update__tmp_h19_).lifecycle, (d_56_dt__update__tmp_h19_).activity, (d_56_dt__update__tmp_h19_).config, (d_56_dt__update__tmp_h19_).meta, (d_56_dt__update__tmp_h19_).creationError, (d_56_dt__update__tmp_h19_).serverTools, d_57_dt__update_hchangesets_h0_, (d_56_dt__update__tmp_h19_).canvases, (d_56_dt__update__tmp_h19_).openCanvases, (d_56_dt__update__tmp_h19_).defaultChat, (d_56_dt__update__tmp_h19_).activeClients, (d_56_dt__update__tmp_h19_).chats, (d_56_dt__update__tmp_h19_).customizations, (d_56_dt__update__tmp_h19_).inputNeeded)
                    return iife75_(_pat_let206_0)
                return iife74_(AhpSkeleton.Option_Some(ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([]))))
            return iife73_(_pat_let205_0)
        if ((default__.apply1(iife72_(d_0_s_), SessionAction_ChangesetsChanged(AhpSkeleton.Option_None()))).changesets) == (AhpSkeleton.Option_None()):
            pass_ = (pass_) + (1)
        d_58_canvasDecls_: ConfluxCodec.Json
        d_58_canvasDecls_ = ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "providerId")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acme.editors"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "canvasId")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "markdown")))}))]))
        if ((default__.apply1(d_0_s_, SessionAction_CanvasesChanged(d_58_canvasDecls_))).canvases) == (AhpSkeleton.Option_Some(d_58_canvasDecls_)):
            pass_ = (pass_) + (1)
        d_59_openRefs_: ConfluxCodec.Json
        d_59_openRefs_ = ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "channel")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahp-canvas:/canvas-1"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "canvasId")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "markdown"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "availability")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")))}))]))
        if ((default__.apply1(d_0_s_, SessionAction_OpenCanvasesChanged(d_59_openRefs_))).openCanvases) == (AhpSkeleton.Option_Some(d_59_openRefs_)):
            pass_ = (pass_) + (1)
        d_60_cl_: Client
        d_60_cl_ = Client_Client(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vscode-1")), ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "displayName")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "VS Code")))})))
        if ((default__.apply1(d_0_s_, SessionAction_ActiveClientSet(d_60_cl_))).activeClients) == (_dafny.SeqWithoutIsStrInference([d_60_cl_])):
            pass_ = (pass_) + (1)
        pat_let_tv11_ = d_60_cl_
        def iife76_(_pat_let207_0):
            def iife77_(d_61_dt__update__tmp_h20_):
                def iife78_(_pat_let208_0):
                    def iife79_(d_62_dt__update_hactiveClients_h0_):
                        return SessionState_SessionState((d_61_dt__update__tmp_h20_).provider, (d_61_dt__update__tmp_h20_).title, (d_61_dt__update__tmp_h20_).status, (d_61_dt__update__tmp_h20_).lifecycle, (d_61_dt__update__tmp_h20_).activity, (d_61_dt__update__tmp_h20_).config, (d_61_dt__update__tmp_h20_).meta, (d_61_dt__update__tmp_h20_).creationError, (d_61_dt__update__tmp_h20_).serverTools, (d_61_dt__update__tmp_h20_).changesets, (d_61_dt__update__tmp_h20_).canvases, (d_61_dt__update__tmp_h20_).openCanvases, (d_61_dt__update__tmp_h20_).defaultChat, d_62_dt__update_hactiveClients_h0_, (d_61_dt__update__tmp_h20_).chats, (d_61_dt__update__tmp_h20_).customizations, (d_61_dt__update__tmp_h20_).inputNeeded)
                    return iife79_(_pat_let208_0)
                return iife78_(_dafny.SeqWithoutIsStrInference([pat_let_tv11_, Client_Client(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cli-1")), ConfluxCodec.Json_JNull())]))
            return iife77_(_pat_let207_0)
        if ((default__.apply1(iife76_(d_0_s_), SessionAction_ActiveClientRemoved(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vscode-1"))))).activeClients) == (_dafny.SeqWithoutIsStrInference([Client_Client(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cli-1")), ConfluxCodec.Json_JNull())])):
            pass_ = (pass_) + (1)
        d_63_ir_: InputReq
        d_63_ir_ = InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "req-1")), ConfluxCodec.Json_JNull())
        def iife80_(_pat_let209_0):
            def iife81_(d_64_dt__update__tmp_h21_):
                def iife82_(_pat_let210_0):
                    def iife83_(d_65_dt__update_hstatus_h4_):
                        return SessionState_SessionState((d_64_dt__update__tmp_h21_).provider, (d_64_dt__update__tmp_h21_).title, d_65_dt__update_hstatus_h4_, (d_64_dt__update__tmp_h21_).lifecycle, (d_64_dt__update__tmp_h21_).activity, (d_64_dt__update__tmp_h21_).config, (d_64_dt__update__tmp_h21_).meta, (d_64_dt__update__tmp_h21_).creationError, (d_64_dt__update__tmp_h21_).serverTools, (d_64_dt__update__tmp_h21_).changesets, (d_64_dt__update__tmp_h21_).canvases, (d_64_dt__update__tmp_h21_).openCanvases, (d_64_dt__update__tmp_h21_).defaultChat, (d_64_dt__update__tmp_h21_).activeClients, (d_64_dt__update__tmp_h21_).chats, (d_64_dt__update__tmp_h21_).customizations, (d_64_dt__update__tmp_h21_).inputNeeded)
                    return iife83_(_pat_let210_0)
                return iife82_(8)
            return iife81_(_pat_let209_0)
        if ((default__.apply1(iife80_(d_0_s_), SessionAction_InputNeededSet(d_63_ir_))).status) == (24):
            pass_ = (pass_) + (1)
        d_66_i1_: InputReq
        d_66_i1_ = InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")), ConfluxCodec.Json_JNull())
        d_67_i2_: InputReq
        d_67_i2_ = InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")), ConfluxCodec.Json_JNull())
        pat_let_tv12_ = d_66_i1_
        pat_let_tv13_ = d_67_i2_
        def iife84_(_pat_let211_0):
            def iife85_(d_68_dt__update__tmp_h22_):
                def iife86_(_pat_let212_0):
                    def iife87_(d_69_dt__update_hinputNeeded_h0_):
                        def iife88_(_pat_let213_0):
                            def iife89_(d_70_dt__update_hstatus_h5_):
                                return SessionState_SessionState((d_68_dt__update__tmp_h22_).provider, (d_68_dt__update__tmp_h22_).title, d_70_dt__update_hstatus_h5_, (d_68_dt__update__tmp_h22_).lifecycle, (d_68_dt__update__tmp_h22_).activity, (d_68_dt__update__tmp_h22_).config, (d_68_dt__update__tmp_h22_).meta, (d_68_dt__update__tmp_h22_).creationError, (d_68_dt__update__tmp_h22_).serverTools, (d_68_dt__update__tmp_h22_).changesets, (d_68_dt__update__tmp_h22_).canvases, (d_68_dt__update__tmp_h22_).openCanvases, (d_68_dt__update__tmp_h22_).defaultChat, (d_68_dt__update__tmp_h22_).activeClients, (d_68_dt__update__tmp_h22_).chats, (d_68_dt__update__tmp_h22_).customizations, d_69_dt__update_hinputNeeded_h0_)
                            return iife89_(_pat_let213_0)
                        return iife88_(24)
                    return iife87_(_pat_let212_0)
                return iife86_(_dafny.SeqWithoutIsStrInference([pat_let_tv12_, pat_let_tv13_]))
            return iife85_(_pat_let211_0)
        if ((default__.apply1(iife84_(d_0_s_), SessionAction_InputNeededRemoved(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))))).status) == (24):
            pass_ = (pass_) + (1)
        d_71_mc_: Cust
        d_71_mc_ = Cust_Cust(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcp-1")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcpServer")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "file:///w")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "MCP")), AhpSkeleton.Option_Some(True), AhpSkeleton.Option_None(), AhpSkeleton.Option_None())
        pat_let_tv14_ = d_71_mc_
        def iife90_(_pat_let214_0):
            def iife91_(d_72_dt__update__tmp_h23_):
                def iife92_(_pat_let215_0):
                    def iife93_(d_73_dt__update_hcustomizations_h2_):
                        return SessionState_SessionState((d_72_dt__update__tmp_h23_).provider, (d_72_dt__update__tmp_h23_).title, (d_72_dt__update__tmp_h23_).status, (d_72_dt__update__tmp_h23_).lifecycle, (d_72_dt__update__tmp_h23_).activity, (d_72_dt__update__tmp_h23_).config, (d_72_dt__update__tmp_h23_).meta, (d_72_dt__update__tmp_h23_).creationError, (d_72_dt__update__tmp_h23_).serverTools, (d_72_dt__update__tmp_h23_).changesets, (d_72_dt__update__tmp_h23_).canvases, (d_72_dt__update__tmp_h23_).openCanvases, (d_72_dt__update__tmp_h23_).defaultChat, (d_72_dt__update__tmp_h23_).activeClients, (d_72_dt__update__tmp_h23_).chats, d_73_dt__update_hcustomizations_h2_, (d_72_dt__update__tmp_h23_).inputNeeded)
                    return iife93_(_pat_let215_0)
                return iife92_(_dafny.SeqWithoutIsStrInference([pat_let_tv14_]))
            return iife91_(_pat_let214_0)
        def iife94_(_pat_let216_0):
            def iife95_(d_74_dt__update__tmp_h24_):
                def iife96_(_pat_let217_0):
                    def iife97_(d_75_dt__update_hchannel_h0_):
                        def iife98_(_pat_let218_0):
                            def iife99_(d_76_dt__update_hstate_h0_):
                                return Cust_Cust((d_74_dt__update__tmp_h24_).id_, (d_74_dt__update__tmp_h24_).kind, (d_74_dt__update__tmp_h24_).uri, (d_74_dt__update__tmp_h24_).name, (d_74_dt__update__tmp_h24_).enabled, d_76_dt__update_hstate_h0_, d_75_dt__update_hchannel_h0_)
                            return iife99_(_pat_let218_0)
                        return iife98_(AhpSkeleton.Option_Some(ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "starting")))}))))
                    return iife97_(_pat_let217_0)
                return iife96_(AhpSkeleton.Option_None())
            return iife95_(_pat_let216_0)
        if ((default__.apply1(iife90_(d_0_s_), SessionAction_McpServerStartRequested(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcp-1"))))).customizations) == (_dafny.SeqWithoutIsStrInference([iife94_(d_71_mc_)])):
            pass_ = (pass_) + (1)
        pat_let_tv15_ = d_71_mc_
        def iife100_(_pat_let219_0):
            def iife101_(d_77_dt__update__tmp_h25_):
                def iife102_(_pat_let220_0):
                    def iife103_(d_78_dt__update_hcustomizations_h3_):
                        return SessionState_SessionState((d_77_dt__update__tmp_h25_).provider, (d_77_dt__update__tmp_h25_).title, (d_77_dt__update__tmp_h25_).status, (d_77_dt__update__tmp_h25_).lifecycle, (d_77_dt__update__tmp_h25_).activity, (d_77_dt__update__tmp_h25_).config, (d_77_dt__update__tmp_h25_).meta, (d_77_dt__update__tmp_h25_).creationError, (d_77_dt__update__tmp_h25_).serverTools, (d_77_dt__update__tmp_h25_).changesets, (d_77_dt__update__tmp_h25_).canvases, (d_77_dt__update__tmp_h25_).openCanvases, (d_77_dt__update__tmp_h25_).defaultChat, (d_77_dt__update__tmp_h25_).activeClients, (d_77_dt__update__tmp_h25_).chats, d_78_dt__update_hcustomizations_h3_, (d_77_dt__update__tmp_h25_).inputNeeded)
                    return iife103_(_pat_let220_0)
                return iife102_(_dafny.SeqWithoutIsStrInference([pat_let_tv15_]))
            return iife101_(_pat_let219_0)
        def iife104_(_pat_let221_0):
            def iife105_(d_79_dt__update__tmp_h26_):
                def iife106_(_pat_let222_0):
                    def iife107_(d_80_dt__update_hchannel_h1_):
                        def iife108_(_pat_let223_0):
                            def iife109_(d_81_dt__update_hstate_h1_):
                                return Cust_Cust((d_79_dt__update__tmp_h26_).id_, (d_79_dt__update__tmp_h26_).kind, (d_79_dt__update__tmp_h26_).uri, (d_79_dt__update__tmp_h26_).name, (d_79_dt__update__tmp_h26_).enabled, d_81_dt__update_hstate_h1_, d_80_dt__update_hchannel_h1_)
                            return iife109_(_pat_let223_0)
                        return iife108_(AhpSkeleton.Option_Some(ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")))}))))
                    return iife107_(_pat_let222_0)
                return iife106_(AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ch")))))
            return iife105_(_pat_let221_0)
        if ((default__.apply1(iife100_(d_0_s_), SessionAction_McpServerStateChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcp-1")), ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")))})), AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ch"))))))).customizations) == (_dafny.SeqWithoutIsStrInference([iife104_(d_71_mc_)])):
            pass_ = (pass_) + (1)
        pat_let_tv16_ = d_24_pa_
        pat_let_tv17_ = d_24_pa_
        def iife110_(_pat_let224_0):
            def iife111_(d_82_dt__update__tmp_h27_):
                def iife112_(_pat_let225_0):
                    def iife113_(d_83_dt__update_hcustomizations_h4_):
                        return SessionState_SessionState((d_82_dt__update__tmp_h27_).provider, (d_82_dt__update__tmp_h27_).title, (d_82_dt__update__tmp_h27_).status, (d_82_dt__update__tmp_h27_).lifecycle, (d_82_dt__update__tmp_h27_).activity, (d_82_dt__update__tmp_h27_).config, (d_82_dt__update__tmp_h27_).meta, (d_82_dt__update__tmp_h27_).creationError, (d_82_dt__update__tmp_h27_).serverTools, (d_82_dt__update__tmp_h27_).changesets, (d_82_dt__update__tmp_h27_).canvases, (d_82_dt__update__tmp_h27_).openCanvases, (d_82_dt__update__tmp_h27_).defaultChat, (d_82_dt__update__tmp_h27_).activeClients, (d_82_dt__update__tmp_h27_).chats, d_83_dt__update_hcustomizations_h4_, (d_82_dt__update__tmp_h27_).inputNeeded)
                    return iife113_(_pat_let225_0)
                return iife112_(_dafny.SeqWithoutIsStrInference([pat_let_tv16_]))
            return iife111_(_pat_let224_0)
        def iife114_(_pat_let226_0):
            def iife115_(d_84_dt__update__tmp_h28_):
                def iife116_(_pat_let227_0):
                    def iife117_(d_85_dt__update_hcustomizations_h5_):
                        return SessionState_SessionState((d_84_dt__update__tmp_h28_).provider, (d_84_dt__update__tmp_h28_).title, (d_84_dt__update__tmp_h28_).status, (d_84_dt__update__tmp_h28_).lifecycle, (d_84_dt__update__tmp_h28_).activity, (d_84_dt__update__tmp_h28_).config, (d_84_dt__update__tmp_h28_).meta, (d_84_dt__update__tmp_h28_).creationError, (d_84_dt__update__tmp_h28_).serverTools, (d_84_dt__update__tmp_h28_).changesets, (d_84_dt__update__tmp_h28_).canvases, (d_84_dt__update__tmp_h28_).openCanvases, (d_84_dt__update__tmp_h28_).defaultChat, (d_84_dt__update__tmp_h28_).activeClients, (d_84_dt__update__tmp_h28_).chats, d_85_dt__update_hcustomizations_h5_, (d_84_dt__update__tmp_h28_).inputNeeded)
                    return iife117_(_pat_let227_0)
                return iife116_(_dafny.SeqWithoutIsStrInference([pat_let_tv17_]))
            return iife115_(_pat_let226_0)
        if (default__.apply1(iife110_(d_0_s_), SessionAction_McpServerStateChanged(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plugin-a")), ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ready")))})), AhpSkeleton.Option_Some(ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nope"))))))) == (iife114_(d_0_s_)):
            pass_ = (pass_) + (1)
        pat_let_tv18_ = d_71_mc_
        def iife118_(_pat_let228_0):
            def iife119_(d_86_dt__update__tmp_h29_):
                def iife120_(_pat_let229_0):
                    def iife121_(d_87_dt__update_hcustomizations_h6_):
                        return SessionState_SessionState((d_86_dt__update__tmp_h29_).provider, (d_86_dt__update__tmp_h29_).title, (d_86_dt__update__tmp_h29_).status, (d_86_dt__update__tmp_h29_).lifecycle, (d_86_dt__update__tmp_h29_).activity, (d_86_dt__update__tmp_h29_).config, (d_86_dt__update__tmp_h29_).meta, (d_86_dt__update__tmp_h29_).creationError, (d_86_dt__update__tmp_h29_).serverTools, (d_86_dt__update__tmp_h29_).changesets, (d_86_dt__update__tmp_h29_).canvases, (d_86_dt__update__tmp_h29_).openCanvases, (d_86_dt__update__tmp_h29_).defaultChat, (d_86_dt__update__tmp_h29_).activeClients, (d_86_dt__update__tmp_h29_).chats, d_87_dt__update_hcustomizations_h6_, (d_86_dt__update__tmp_h29_).inputNeeded)
                    return iife121_(_pat_let229_0)
                return iife120_(_dafny.SeqWithoutIsStrInference([pat_let_tv18_]))
            return iife119_(_pat_let228_0)
        def iife122_(_pat_let230_0):
            def iife123_(d_88_dt__update__tmp_h30_):
                def iife124_(_pat_let231_0):
                    def iife125_(d_89_dt__update_hchannel_h2_):
                        def iife126_(_pat_let232_0):
                            def iife127_(d_90_dt__update_hstate_h2_):
                                return Cust_Cust((d_88_dt__update__tmp_h30_).id_, (d_88_dt__update__tmp_h30_).kind, (d_88_dt__update__tmp_h30_).uri, (d_88_dt__update__tmp_h30_).name, (d_88_dt__update__tmp_h30_).enabled, d_90_dt__update_hstate_h2_, d_89_dt__update_hchannel_h2_)
                            return iife127_(_pat_let232_0)
                        return iife126_(AhpSkeleton.Option_Some(ConfluxCodec.Json_JObj(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kind")): ConfluxCodec.Json_JStr(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stopped")))}))))
                    return iife125_(_pat_let231_0)
                return iife124_(AhpSkeleton.Option_None())
            return iife123_(_pat_let230_0)
        if ((default__.apply1(iife118_(d_0_s_), SessionAction_McpServerStopRequested(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcp-1"))))).customizations) == (_dafny.SeqWithoutIsStrInference([iife122_(d_71_mc_)])):
            pass_ = (pass_) + (1)
        if ((default__.apply1(d_0_s_, SessionAction_CustomizationUpdated(d_71_mc_))).customizations) == (_dafny.SeqWithoutIsStrInference([d_71_mc_])):
            pass_ = (pass_) + (1)
        return pass_, modeled

    @_dafny.classproperty
    def custKey(instance):
        def lambda0_(d_0_x_):
            return (d_0_x_).id_

        return lambda0_
    @_dafny.classproperty
    def chatKey(instance):
        def lambda0_(d_0_x_):
            return (d_0_x_).resource

        return lambda0_
    @_dafny.classproperty
    def clientKey(instance):
        def lambda0_(d_0_x_):
            return (d_0_x_).clientId

        return lambda0_
    @_dafny.classproperty
    def inputKey(instance):
        def lambda0_(d_0_x_):
            return (d_0_x_).id_

        return lambda0_
    @_dafny.classproperty
    def mcpPred(instance):
        def lambda0_(d_0_c_):
            return ((d_0_c_).kind) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcpServer")))

        return lambda0_
    @_dafny.classproperty
    def B__READ(instance):
        return 32
    @_dafny.classproperty
    def B__ARCH(instance):
        return 64
    @_dafny.classproperty
    def B__INPUT(instance):
        return 16

class Cust:
    @classmethod
    def default(cls, ):
        return lambda: Cust_Cust(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Cust(self) -> bool:
        return isinstance(self, Cust_Cust)

class Cust_Cust(Cust, NamedTuple('Cust', [('id_', Any), ('kind', Any), ('uri', Any), ('name', Any), ('enabled', Any), ('state', Any), ('channel', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.Cust.Cust({self.id_.VerbatimString(True)}, {self.kind.VerbatimString(True)}, {self.uri.VerbatimString(True)}, {self.name.VerbatimString(True)}, {_dafny.string_of(self.enabled)}, {_dafny.string_of(self.state)}, {_dafny.string_of(self.channel)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Cust_Cust) and self.id_ == __o.id_ and self.kind == __o.kind and self.uri == __o.uri and self.name == __o.name and self.enabled == __o.enabled and self.state == __o.state and self.channel == __o.channel
    def __hash__(self) -> int:
        return super().__hash__()


class Chat:
    @classmethod
    def default(cls, ):
        return lambda: Chat_Chat(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), AhpSkeleton.Option.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxCodec.Json.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Chat(self) -> bool:
        return isinstance(self, Chat_Chat)

class Chat_Chat(Chat, NamedTuple('Chat', [('resource', Any), ('title', Any), ('status', Any), ('activity', Any), ('modifiedAt', Any), ('origin', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.Chat.Chat({self.resource.VerbatimString(True)}, {self.title.VerbatimString(True)}, {_dafny.string_of(self.status)}, {_dafny.string_of(self.activity)}, {self.modifiedAt.VerbatimString(True)}, {_dafny.string_of(self.origin)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Chat_Chat) and self.resource == __o.resource and self.title == __o.title and self.status == __o.status and self.activity == __o.activity and self.modifiedAt == __o.modifiedAt and self.origin == __o.origin
    def __hash__(self) -> int:
        return super().__hash__()


class Client:
    @classmethod
    def default(cls, ):
        return lambda: Client_Client(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxCodec.Json.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Client(self) -> bool:
        return isinstance(self, Client_Client)

class Client_Client(Client, NamedTuple('Client', [('clientId', Any), ('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.Client.Client({self.clientId.VerbatimString(True)}, {_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Client_Client) and self.clientId == __o.clientId and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()


class InputReq:
    @classmethod
    def default(cls, ):
        return lambda: InputReq_InputReq(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxCodec.Json.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_InputReq(self) -> bool:
        return isinstance(self, InputReq_InputReq)

class InputReq_InputReq(InputReq, NamedTuple('InputReq', [('id_', Any), ('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.InputReq.InputReq({self.id_.VerbatimString(True)}, {_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, InputReq_InputReq) and self.id_ == __o.id_ and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()


class SConfig:
    @classmethod
    def default(cls, ):
        return lambda: SConfig_SConfig(ConfluxCodec.Json.default()(), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SConfig(self) -> bool:
        return isinstance(self, SConfig_SConfig)

class SConfig_SConfig(SConfig, NamedTuple('SConfig', [('schema', Any), ('values', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SConfig.SConfig({_dafny.string_of(self.schema)}, {_dafny.string_of(self.values)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SConfig_SConfig) and self.schema == __o.schema and self.values == __o.values
    def __hash__(self) -> int:
        return super().__hash__()


class SessionState:
    @classmethod
    def default(cls, ):
        return lambda: SessionState_SessionState(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), AhpSkeleton.Option.default()(), _dafny.Seq({}), _dafny.Seq({}), _dafny.Seq({}), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SessionState(self) -> bool:
        return isinstance(self, SessionState_SessionState)

class SessionState_SessionState(SessionState, NamedTuple('SessionState', [('provider', Any), ('title', Any), ('status', Any), ('lifecycle', Any), ('activity', Any), ('config', Any), ('meta', Any), ('creationError', Any), ('serverTools', Any), ('changesets', Any), ('canvases', Any), ('openCanvases', Any), ('defaultChat', Any), ('activeClients', Any), ('chats', Any), ('customizations', Any), ('inputNeeded', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionState.SessionState({self.provider.VerbatimString(True)}, {self.title.VerbatimString(True)}, {_dafny.string_of(self.status)}, {self.lifecycle.VerbatimString(True)}, {_dafny.string_of(self.activity)}, {_dafny.string_of(self.config)}, {_dafny.string_of(self.meta)}, {_dafny.string_of(self.creationError)}, {_dafny.string_of(self.serverTools)}, {_dafny.string_of(self.changesets)}, {_dafny.string_of(self.canvases)}, {_dafny.string_of(self.openCanvases)}, {_dafny.string_of(self.defaultChat)}, {_dafny.string_of(self.activeClients)}, {_dafny.string_of(self.chats)}, {_dafny.string_of(self.customizations)}, {_dafny.string_of(self.inputNeeded)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionState_SessionState) and self.provider == __o.provider and self.title == __o.title and self.status == __o.status and self.lifecycle == __o.lifecycle and self.activity == __o.activity and self.config == __o.config and self.meta == __o.meta and self.creationError == __o.creationError and self.serverTools == __o.serverTools and self.changesets == __o.changesets and self.canvases == __o.canvases and self.openCanvases == __o.openCanvases and self.defaultChat == __o.defaultChat and self.activeClients == __o.activeClients and self.chats == __o.chats and self.customizations == __o.customizations and self.inputNeeded == __o.inputNeeded
    def __hash__(self) -> int:
        return super().__hash__()


class SessionAction:
    @classmethod
    def default(cls, ):
        return lambda: SessionAction_Ready()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Ready(self) -> bool:
        return isinstance(self, SessionAction_Ready)
    @property
    def is_CreationFailed(self) -> bool:
        return isinstance(self, SessionAction_CreationFailed)
    @property
    def is_TitleChanged(self) -> bool:
        return isinstance(self, SessionAction_TitleChanged)
    @property
    def is_ServerToolsChanged(self) -> bool:
        return isinstance(self, SessionAction_ServerToolsChanged)
    @property
    def is_MetaChanged(self) -> bool:
        return isinstance(self, SessionAction_MetaChanged)
    @property
    def is_IsReadChanged(self) -> bool:
        return isinstance(self, SessionAction_IsReadChanged)
    @property
    def is_IsArchivedChanged(self) -> bool:
        return isinstance(self, SessionAction_IsArchivedChanged)
    @property
    def is_ActivityChanged(self) -> bool:
        return isinstance(self, SessionAction_ActivityChanged)
    @property
    def is_ConfigChanged(self) -> bool:
        return isinstance(self, SessionAction_ConfigChanged)
    @property
    def is_CustomizationsChanged(self) -> bool:
        return isinstance(self, SessionAction_CustomizationsChanged)
    @property
    def is_CustomizationToggled(self) -> bool:
        return isinstance(self, SessionAction_CustomizationToggled)
    @property
    def is_CustomizationRemoved(self) -> bool:
        return isinstance(self, SessionAction_CustomizationRemoved)
    @property
    def is_DefaultChatChanged(self) -> bool:
        return isinstance(self, SessionAction_DefaultChatChanged)
    @property
    def is_ChatAdded(self) -> bool:
        return isinstance(self, SessionAction_ChatAdded)
    @property
    def is_ChatRemoved(self) -> bool:
        return isinstance(self, SessionAction_ChatRemoved)
    @property
    def is_ChatUpdated(self) -> bool:
        return isinstance(self, SessionAction_ChatUpdated)
    @property
    def is_ChangesetsChanged(self) -> bool:
        return isinstance(self, SessionAction_ChangesetsChanged)
    @property
    def is_ActiveClientSet(self) -> bool:
        return isinstance(self, SessionAction_ActiveClientSet)
    @property
    def is_ActiveClientRemoved(self) -> bool:
        return isinstance(self, SessionAction_ActiveClientRemoved)
    @property
    def is_CanvasesChanged(self) -> bool:
        return isinstance(self, SessionAction_CanvasesChanged)
    @property
    def is_OpenCanvasesChanged(self) -> bool:
        return isinstance(self, SessionAction_OpenCanvasesChanged)
    @property
    def is_InputNeededSet(self) -> bool:
        return isinstance(self, SessionAction_InputNeededSet)
    @property
    def is_InputNeededRemoved(self) -> bool:
        return isinstance(self, SessionAction_InputNeededRemoved)
    @property
    def is_CustomizationUpdated(self) -> bool:
        return isinstance(self, SessionAction_CustomizationUpdated)
    @property
    def is_McpServerStateChanged(self) -> bool:
        return isinstance(self, SessionAction_McpServerStateChanged)
    @property
    def is_McpServerStartRequested(self) -> bool:
        return isinstance(self, SessionAction_McpServerStartRequested)
    @property
    def is_McpServerStopRequested(self) -> bool:
        return isinstance(self, SessionAction_McpServerStopRequested)
    @property
    def is_SUnknown(self) -> bool:
        return isinstance(self, SessionAction_SUnknown)

class SessionAction_Ready(SessionAction, NamedTuple('Ready', [])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.Ready'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_Ready)
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_CreationFailed(SessionAction, NamedTuple('CreationFailed', [('error', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.CreationFailed({_dafny.string_of(self.error)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_CreationFailed) and self.error == __o.error
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_TitleChanged(SessionAction, NamedTuple('TitleChanged', [('title', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.TitleChanged({self.title.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_TitleChanged) and self.title == __o.title
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_ServerToolsChanged(SessionAction, NamedTuple('ServerToolsChanged', [('tools', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.ServerToolsChanged({_dafny.string_of(self.tools)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_ServerToolsChanged) and self.tools == __o.tools
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_MetaChanged(SessionAction, NamedTuple('MetaChanged', [('meta', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.MetaChanged({_dafny.string_of(self.meta)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_MetaChanged) and self.meta == __o.meta
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_IsReadChanged(SessionAction, NamedTuple('IsReadChanged', [('isRead', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.IsReadChanged({_dafny.string_of(self.isRead)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_IsReadChanged) and self.isRead == __o.isRead
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_IsArchivedChanged(SessionAction, NamedTuple('IsArchivedChanged', [('isArchived', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.IsArchivedChanged({_dafny.string_of(self.isArchived)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_IsArchivedChanged) and self.isArchived == __o.isArchived
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_ActivityChanged(SessionAction, NamedTuple('ActivityChanged', [('activity', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.ActivityChanged({_dafny.string_of(self.activity)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_ActivityChanged) and self.activity == __o.activity
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_ConfigChanged(SessionAction, NamedTuple('ConfigChanged', [('config', Any), ('replace', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.ConfigChanged({_dafny.string_of(self.config)}, {_dafny.string_of(self.replace)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_ConfigChanged) and self.config == __o.config and self.replace == __o.replace
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_CustomizationsChanged(SessionAction, NamedTuple('CustomizationsChanged', [('customizations', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.CustomizationsChanged({_dafny.string_of(self.customizations)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_CustomizationsChanged) and self.customizations == __o.customizations
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_CustomizationToggled(SessionAction, NamedTuple('CustomizationToggled', [('id_', Any), ('enabled', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.CustomizationToggled({self.id_.VerbatimString(True)}, {_dafny.string_of(self.enabled)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_CustomizationToggled) and self.id_ == __o.id_ and self.enabled == __o.enabled
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_CustomizationRemoved(SessionAction, NamedTuple('CustomizationRemoved', [('id_', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.CustomizationRemoved({self.id_.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_CustomizationRemoved) and self.id_ == __o.id_
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_DefaultChatChanged(SessionAction, NamedTuple('DefaultChatChanged', [('chat', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.DefaultChatChanged({_dafny.string_of(self.chat)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_DefaultChatChanged) and self.chat == __o.chat
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_ChatAdded(SessionAction, NamedTuple('ChatAdded', [('summary', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.ChatAdded({_dafny.string_of(self.summary)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_ChatAdded) and self.summary == __o.summary
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_ChatRemoved(SessionAction, NamedTuple('ChatRemoved', [('resource', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.ChatRemoved({self.resource.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_ChatRemoved) and self.resource == __o.resource
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_ChatUpdated(SessionAction, NamedTuple('ChatUpdated', [('resource', Any), ('status', Any), ('activity', Any), ('modifiedAt', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.ChatUpdated({self.resource.VerbatimString(True)}, {_dafny.string_of(self.status)}, {_dafny.string_of(self.activity)}, {_dafny.string_of(self.modifiedAt)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_ChatUpdated) and self.resource == __o.resource and self.status == __o.status and self.activity == __o.activity and self.modifiedAt == __o.modifiedAt
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_ChangesetsChanged(SessionAction, NamedTuple('ChangesetsChanged', [('changesets', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.ChangesetsChanged({_dafny.string_of(self.changesets)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_ChangesetsChanged) and self.changesets == __o.changesets
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_ActiveClientSet(SessionAction, NamedTuple('ActiveClientSet', [('client', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.ActiveClientSet({_dafny.string_of(self.client)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_ActiveClientSet) and self.client == __o.client
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_ActiveClientRemoved(SessionAction, NamedTuple('ActiveClientRemoved', [('clientId', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.ActiveClientRemoved({self.clientId.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_ActiveClientRemoved) and self.clientId == __o.clientId
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_CanvasesChanged(SessionAction, NamedTuple('CanvasesChanged', [('canvases', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.CanvasesChanged({_dafny.string_of(self.canvases)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_CanvasesChanged) and self.canvases == __o.canvases
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_OpenCanvasesChanged(SessionAction, NamedTuple('OpenCanvasesChanged', [('openCanvases', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.OpenCanvasesChanged({_dafny.string_of(self.openCanvases)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_OpenCanvasesChanged) and self.openCanvases == __o.openCanvases
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_InputNeededSet(SessionAction, NamedTuple('InputNeededSet', [('req', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.InputNeededSet({_dafny.string_of(self.req)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_InputNeededSet) and self.req == __o.req
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_InputNeededRemoved(SessionAction, NamedTuple('InputNeededRemoved', [('id_', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.InputNeededRemoved({self.id_.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_InputNeededRemoved) and self.id_ == __o.id_
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_CustomizationUpdated(SessionAction, NamedTuple('CustomizationUpdated', [('customization', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.CustomizationUpdated({_dafny.string_of(self.customization)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_CustomizationUpdated) and self.customization == __o.customization
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_McpServerStateChanged(SessionAction, NamedTuple('McpServerStateChanged', [('id_', Any), ('state', Any), ('channel', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.McpServerStateChanged({self.id_.VerbatimString(True)}, {_dafny.string_of(self.state)}, {_dafny.string_of(self.channel)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_McpServerStateChanged) and self.id_ == __o.id_ and self.state == __o.state and self.channel == __o.channel
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_McpServerStartRequested(SessionAction, NamedTuple('McpServerStartRequested', [('id_', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.McpServerStartRequested({self.id_.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_McpServerStartRequested) and self.id_ == __o.id_
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_McpServerStopRequested(SessionAction, NamedTuple('McpServerStopRequested', [('id_', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.McpServerStopRequested({self.id_.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_McpServerStopRequested) and self.id_ == __o.id_
    def __hash__(self) -> int:
        return super().__hash__()

class SessionAction_SUnknown(SessionAction, NamedTuple('SUnknown', [('raw', Any)])):
    def __dafnystr__(self) -> str:
        return f'Session.SessionAction.SUnknown({_dafny.string_of(self.raw)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SessionAction_SUnknown) and self.raw == __o.raw
    def __hash__(self) -> int:
        return super().__hash__()

