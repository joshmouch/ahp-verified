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
import Chat as Chat

# Module: ClientMain

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Main(noArgsParameter__):
        d_0_rootPass_: int
        d_1_rootTotal_: int
        out0_: int
        out1_: int
        out0_, out1_ = AhpSkeleton.default__.RunCorpus()
        d_0_rootPass_ = out0_
        d_1_rootTotal_ = out1_
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ROOT CORPUS:          "))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_0_rootPass_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_1_rootTotal_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, " green against extracted code\n"))).VerbatimString(False))
        d_2_rwPass_: int
        d_3_rwTotal_: int
        out2_: int
        out3_: int
        out2_, out3_ = ResourceWatch.default__.RunCorpus()
        d_2_rwPass_ = out2_
        d_3_rwTotal_ = out3_
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "RESOURCEWATCH CORPUS: "))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_2_rwPass_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_3_rwTotal_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, " green against extracted code\n"))).VerbatimString(False))
        d_4_cvPass_: int
        d_5_cvTotal_: int
        out4_: int
        out5_: int
        out4_, out5_ = Canvas.default__.RunCorpus()
        d_4_cvPass_ = out4_
        d_5_cvTotal_ = out5_
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "CANVAS CORPUS:        "))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_4_cvPass_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_5_cvTotal_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, " green against extracted code\n"))).VerbatimString(False))
        d_6_csPass_: int
        d_7_csTotal_: int
        out6_: int
        out7_: int
        out6_, out7_ = Changeset.default__.RunCorpus()
        d_6_csPass_ = out6_
        d_7_csTotal_ = out7_
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "CHANGESET CORPUS:     "))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_6_csPass_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_7_csTotal_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, " green against extracted code\n"))).VerbatimString(False))
        d_8_anPass_: int
        d_9_anTotal_: int
        out8_: int
        out9_: int
        out8_, out9_ = Annotations.default__.RunCorpus()
        d_8_anPass_ = out8_
        d_9_anTotal_ = out9_
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ANNOTATIONS CORPUS:   "))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_8_anPass_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_9_anTotal_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, " green against extracted code\n"))).VerbatimString(False))
        d_10_tPass_: int
        d_11_tTotal_: int
        out10_: int
        out11_: int
        out10_, out11_ = Terminal.default__.RunCorpus()
        d_10_tPass_ = out10_
        d_11_tTotal_ = out11_
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "TERMINAL CORPUS:      "))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_10_tPass_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_11_tTotal_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, " green against extracted code\n"))).VerbatimString(False))
        d_12_sPass_: int
        d_13_sModeled_: int
        out12_: int
        out13_: int
        out12_, out13_ = Session.default__.RunCorpus()
        d_12_sPass_ = out12_
        d_13_sModeled_ = out13_
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "SESSION CORPUS:       "))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_12_sPass_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_13_sModeled_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, " modeled green (of 61 total; all ~25 action TYPES now modeled)\n"))).VerbatimString(False))
        d_14_chPass_: int
        d_15_chModeled_: int
        out14_: int
        out15_: int
        out14_, out15_ = Chat.default__.RunCorpus()
        d_14_chPass_ = out14_
        d_15_chModeled_ = out15_
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "CHAT CORPUS:          "))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_14_chPass_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_15_chModeled_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, " modeled green (of 97 total; full tool-call state machine + turn lifecycle modeled)\n"))).VerbatimString(False))
        d_16_pass_: int
        d_16_pass_ = (((((((d_0_rootPass_) + (d_2_rwPass_)) + (d_4_cvPass_)) + (d_6_csPass_)) + (d_8_anPass_)) + (d_10_tPass_)) + (d_12_sPass_)) + (d_14_chPass_)
        d_17_total_: int
        d_17_total_ = (((((((d_1_rootTotal_) + (d_3_rwTotal_)) + (d_5_cvTotal_)) + (d_7_csTotal_)) + (d_9_anTotal_)) + (d_11_tTotal_)) + (d_13_sModeled_)) + (d_15_chModeled_)
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "TOTAL: "))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_16_pass_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "/"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_17_total_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, " corpus fixtures green (5 full AHP channels + session/chat partial)\n"))).VerbatimString(False))
        if not((d_16_pass_) == (d_17_total_)):
            raise _dafny.HaltException("/Users/josh/Code/joshmouch/conflux/agent-host-protocol-workspace/agent-host-protocol-core-dafny/spec/client_main.dfy(41,4): " + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "expectation violation"))).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "MULTI-CHANNEL CLIENT OK — extract(cs,js) + corpus all green\n"))).VerbatimString(False))

