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
import ClientMain as ClientMain
import ConfluxSemanticGraphIdentity as ConfluxSemanticGraphIdentity
import ConfluxContentAddress as ConfluxContentAddress
import ConfluxSemanticGraphContract as ConfluxSemanticGraphContract
import ConfluxResourceCeilings as ConfluxResourceCeilings
import ConfluxSupervisedOperationResult as ConfluxSupervisedOperationResult
import ConfluxTree as ConfluxTree
import ConfluxMerge as ConfluxMerge
import ConfluxRoute as ConfluxRoute
import ConfluxMapValues as ConfluxMapValues
import ConfluxDelimited as ConfluxDelimited
import ConfluxFixpoint as ConfluxFixpoint
import ConfluxGenericCodec as ConfluxGenericCodec
import ConfluxResolve as ConfluxResolve
import ConfluxWatermark as ConfluxWatermark
import ConfluxStateMachine as ConfluxStateMachine
import ConfluxProduct as ConfluxProduct
import ConfluxJoin as ConfluxJoin
import ConfluxDedupe as ConfluxDedupe
import ConfluxBatchRoute as ConfluxBatchRoute
import ConfluxSeqRouteMerge as ConfluxSeqRouteMerge
import ConfluxKeyedOps as ConfluxKeyedOps
import ConfluxFilterKeys as ConfluxFilterKeys
import ConfluxChannel as ConfluxChannel
import ConfluxDurableHistory as ConfluxDurableHistory
import ConfluxCommand as ConfluxCommand
import ConfluxChannelManifest as ConfluxChannelManifest
import ConfluxConvergence as ConfluxConvergence
import ConfluxBudgetConvergence as ConfluxBudgetConvergence
import ConfluxExtern as ConfluxExtern
import ConfluxCommandIdentityCapability as ConfluxCommandIdentityCapability
import ConfluxDecimalText as ConfluxDecimalText
import ConfluxJsonText as ConfluxJsonText
import ConfluxCommandIdentity as ConfluxCommandIdentity
import ConfluxIoCapability as ConfluxIoCapability
import ConfluxResourceSupervisor as ConfluxResourceSupervisor
import ConfluxSocketCapability as ConfluxSocketCapability
import ConfluxWebSocketCapability as ConfluxWebSocketCapability
import ConfluxHttpCapability as ConfluxHttpCapability
import ConfluxServiceLifecycle as ConfluxServiceLifecycle
import ConfluxSequence as ConfluxSequence
import ConfluxJsonRpc as ConfluxJsonRpc
import ConfluxHttpService as ConfluxHttpService

# Module: ConfluxContentLength

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def Token():
        return _dafny.SeqWithoutIsStrInference([67, 111, 110, 116, 101, 110, 116, 45, 76, 101, 110, 103, 116, 104, 58])

    @staticmethod
    def Separator():
        return _dafny.SeqWithoutIsStrInference([13, 10, 13, 10])

    @staticmethod
    def SkipSpaces(header, at):
        while True:
            with _dafny.label():
                if ((at) < (len(header))) and ((((header)[at]) == (32)) or (((header)[at]) == (9))):
                    in0_ = header
                    in1_ = (at) + (1)
                    header = in0_
                    at = in1_
                    raise _dafny.TailCall()
                elif True:
                    return at
                break

    @staticmethod
    def IsDigit(value):
        return ((48) <= (value)) and ((value) <= (57))

    @staticmethod
    def ReadLength(header, at, value, seen):
        while True:
            with _dafny.label():
                if ((at) < (len(header))) and (default__.IsDigit((header)[at])):
                    in0_ = header
                    in1_ = (at) + (1)
                    in2_ = ((value) * (10)) + (((header)[at]) - (48))
                    in3_ = True
                    header = in0_
                    at = in1_
                    value = in2_
                    seen = in3_
                    raise _dafny.TailCall()
                elif seen:
                    return ConfluxContract.Option_Some(value)
                elif True:
                    return ConfluxContract.Option_None()
                break

    @staticmethod
    def Parse(header):
        d_0_token_ = default__.Token()
        d_1_at_ = ConfluxSequence.default__.FindSubsequenceFrom(header, d_0_token_, 0)
        if (d_1_at_) < (0):
            return ConfluxContract.Option_None()
        elif ((d_1_at_) + (len(d_0_token_))) > (len(header)):
            return ConfluxContract.Option_None()
        elif True:
            return default__.ReadLength(header, default__.SkipSpaces(header, (d_1_at_) + (len(d_0_token_))), 0, False)

    @staticmethod
    def HeaderText(length):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "Content-Length: "))) + (ConfluxDecimalText.default__.NatText(length))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\r\n\r\n")))

    @staticmethod
    def Frame(header, body):
        return (header) + (body)

    @staticmethod
    def Extract(data):
        rest: _dafny.Seq = _dafny.Seq({})
        bodies: _dafny.Seq = _dafny.Seq({})
        d_0_separator_: _dafny.Seq
        d_0_separator_ = default__.Separator()
        d_1_separatorAt_: int
        d_1_separatorAt_ = ConfluxSequence.default__.FindSubsequenceFrom(data, d_0_separator_, 0)
        if (d_1_separatorAt_) < (0):
            rhs0_ = data
            rhs1_ = _dafny.SeqWithoutIsStrInference([])
            rest = rhs0_
            bodies = rhs1_
            return rest, bodies
        d_2_bodyStart_: int
        d_2_bodyStart_ = (d_1_separatorAt_) + (len(d_0_separator_))
        source0_ = default__.Parse(_dafny.SeqWithoutIsStrInference((data)[:d_1_separatorAt_:]))
        with _dafny.label("match0"):
            if True:
                if source0_.is_None:
                    out0_: _dafny.Seq
                    out1_: _dafny.Seq
                    out0_, out1_ = default__.Extract(_dafny.SeqWithoutIsStrInference((data)[d_2_bodyStart_::]))
                    rest = out0_
                    bodies = out1_
                    raise _dafny.Break("match0")
            if True:
                d_3_length_ = source0_.value
                if (len(data)) < ((d_2_bodyStart_) + (d_3_length_)):
                    rhs2_ = data
                    rhs3_ = _dafny.SeqWithoutIsStrInference([])
                    rest = rhs2_
                    bodies = rhs3_
                elif True:
                    d_4_tailBodies_: _dafny.Seq = _dafny.Seq({})
                    out2_: _dafny.Seq
                    out3_: _dafny.Seq
                    out2_, out3_ = default__.Extract(_dafny.SeqWithoutIsStrInference((data)[(d_2_bodyStart_) + (d_3_length_)::]))
                    rest = out2_
                    d_4_tailBodies_ = out3_
                    bodies = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference((data)[d_2_bodyStart_:(d_2_bodyStart_) + (d_3_length_):])])) + (d_4_tailBodies_)
            pass
        return rest, bodies

