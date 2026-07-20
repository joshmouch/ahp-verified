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

# Module: ConfluxJsonText

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def IsWs(c):
        return ((((c) == (_dafny.CodePoint(' '))) or ((c) == (_dafny.CodePoint('\n')))) or ((c) == (_dafny.CodePoint('\t')))) or ((c) == (_dafny.CodePoint('\r')))

    @staticmethod
    def IsDigit(c):
        return ((_dafny.CodePoint('0')) <= (c)) and ((c) <= (_dafny.CodePoint('9')))

    @staticmethod
    def SkipWs(s, pos):
        while True:
            with _dafny.label():
                if ((pos) < (len(s))) and (default__.IsWs((s)[pos])):
                    in0_ = s
                    in1_ = (pos) + (1)
                    s = in0_
                    pos = in1_
                    raise _dafny.TailCall()
                elif True:
                    return pos
                break

    @staticmethod
    def HexVal(c):
        if ((_dafny.CodePoint('0')) <= (c)) and ((c) <= (_dafny.CodePoint('9'))):
            return (ord(c)) - (ord(_dafny.CodePoint('0')))
        elif ((_dafny.CodePoint('a')) <= (c)) and ((c) <= (_dafny.CodePoint('f'))):
            return ((ord(c)) - (ord(_dafny.CodePoint('a')))) + (10)
        elif ((_dafny.CodePoint('A')) <= (c)) and ((c) <= (_dafny.CodePoint('F'))):
            return ((ord(c)) - (ord(_dafny.CodePoint('A')))) + (10)
        elif True:
            return -1

    @staticmethod
    def HexDigit(n):
        if (n) < (10):
            return _dafny.CodePoint(chr((ord(_dafny.CodePoint('0'))) + (n)))
        elif True:
            return _dafny.CodePoint(chr((ord(_dafny.CodePoint('A'))) + ((n) - (10))))

    @staticmethod
    def Hex4(s, p):
        d_0_d0_ = default__.HexVal((s)[p])
        d_1_d1_ = default__.HexVal((s)[(p) + (1)])
        d_2_d2_ = default__.HexVal((s)[(p) + (2)])
        d_3_d3_ = default__.HexVal((s)[(p) + (3)])
        if ((((d_0_d0_) < (0)) or ((d_1_d1_) < (0))) or ((d_2_d2_) < (0))) or ((d_3_d3_) < (0)):
            return -1
        elif True:
            return ((((d_0_d0_) * (4096)) + ((d_1_d1_) * (256))) + ((d_2_d2_) * (16))) + (d_3_d3_)

    @staticmethod
    def IsScalar(v):
        return (((0) <= (v)) and ((v) < (55296))) or (((57344) <= (v)) and ((v) <= (1114111)))

    @staticmethod
    def ScalarToChar(v):
        if default__.IsScalar(v):
            return _dafny.CodePoint(chr(v))
        elif True:
            return _dafny.CodePoint('\U0000FFFD')

    @staticmethod
    def ScanDigits(s, pos, acc, count):
        while True:
            with _dafny.label():
                if ((pos) < (len(s))) and (default__.IsDigit((s)[pos])):
                    in0_ = s
                    in1_ = (pos) + (1)
                    in2_ = ((acc) * (10)) + ((ord((s)[pos])) - (ord(_dafny.CodePoint('0'))))
                    in3_ = (count) + (1)
                    s = in0_
                    pos = in1_
                    acc = in2_
                    count = in3_
                    raise _dafny.TailCall()
                elif True:
                    return DigitCursor_DigitCursor(acc, pos, count)
                break

    @staticmethod
    def ParseStringTail(s, pos, acc):
        while True:
            with _dafny.label():
                if (pos) >= (len(s)):
                    return StringCursor_StringCursor(acc, pos, False)
                elif ((s)[pos]) == (_dafny.CodePoint('"')):
                    return StringCursor_StringCursor(acc, (pos) + (1), True)
                elif ((s)[pos]) != (_dafny.CodePoint('\\')):
                    if (ord((s)[pos])) < (32):
                        return StringCursor_StringCursor(acc, pos, False)
                    elif True:
                        in0_ = s
                        in1_ = (pos) + (1)
                        in2_ = (acc) + (_dafny.SeqWithoutIsStrInference([(s)[pos]]))
                        s = in0_
                        pos = in1_
                        acc = in2_
                        raise _dafny.TailCall()
                elif ((pos) + (1)) >= (len(s)):
                    return StringCursor_StringCursor(acc, pos, False)
                elif True:
                    d_0_e_ = (s)[(pos) + (1)]
                    if (((d_0_e_) == (_dafny.CodePoint('"'))) or ((d_0_e_) == (_dafny.CodePoint('\\')))) or ((d_0_e_) == (_dafny.CodePoint('/'))):
                        in3_ = s
                        in4_ = (pos) + (2)
                        in5_ = (acc) + (_dafny.SeqWithoutIsStrInference([d_0_e_]))
                        s = in3_
                        pos = in4_
                        acc = in5_
                        raise _dafny.TailCall()
                    elif (d_0_e_) == (_dafny.CodePoint('n')):
                        in6_ = s
                        in7_ = (pos) + (2)
                        in8_ = (acc) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('\n')]))
                        s = in6_
                        pos = in7_
                        acc = in8_
                        raise _dafny.TailCall()
                    elif (d_0_e_) == (_dafny.CodePoint('t')):
                        in9_ = s
                        in10_ = (pos) + (2)
                        in11_ = (acc) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('\t')]))
                        s = in9_
                        pos = in10_
                        acc = in11_
                        raise _dafny.TailCall()
                    elif (d_0_e_) == (_dafny.CodePoint('r')):
                        in12_ = s
                        in13_ = (pos) + (2)
                        in14_ = (acc) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('\r')]))
                        s = in12_
                        pos = in13_
                        acc = in14_
                        raise _dafny.TailCall()
                    elif (d_0_e_) == (_dafny.CodePoint('b')):
                        in15_ = s
                        in16_ = (pos) + (2)
                        in17_ = (acc) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('\U00000008')]))
                        s = in15_
                        pos = in16_
                        acc = in17_
                        raise _dafny.TailCall()
                    elif (d_0_e_) == (_dafny.CodePoint('f')):
                        in18_ = s
                        in19_ = (pos) + (2)
                        in20_ = (acc) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('\U0000000C')]))
                        s = in18_
                        pos = in19_
                        acc = in20_
                        raise _dafny.TailCall()
                    elif (((d_0_e_) == (_dafny.CodePoint('u'))) and (((pos) + (6)) <= (len(s)))) and ((default__.Hex4(s, (pos) + (2))) >= (0)):
                        d_1_code_ = default__.Hex4(s, (pos) + (2))
                        if ((((((55296) <= (d_1_code_)) and ((d_1_code_) <= (56319))) and (((pos) + (12)) <= (len(s)))) and (((s)[(pos) + (6)]) == (_dafny.CodePoint('\\')))) and (((s)[(pos) + (7)]) == (_dafny.CodePoint('u')))) and (((56320) <= (default__.Hex4(s, (pos) + (8)))) and ((default__.Hex4(s, (pos) + (8))) <= (57343))):
                            d_2_lo_ = default__.Hex4(s, (pos) + (8))
                            in21_ = s
                            in22_ = (pos) + (12)
                            in23_ = (acc) + (_dafny.SeqWithoutIsStrInference([default__.ScalarToChar(((65536) + (((d_1_code_) - (55296)) * (1024))) + ((d_2_lo_) - (56320)))]))
                            s = in21_
                            pos = in22_
                            acc = in23_
                            raise _dafny.TailCall()
                        elif True:
                            in24_ = s
                            in25_ = (pos) + (6)
                            in26_ = (acc) + (_dafny.SeqWithoutIsStrInference([default__.ScalarToChar(d_1_code_)]))
                            s = in24_
                            pos = in25_
                            acc = in26_
                            raise _dafny.TailCall()
                    elif True:
                        return StringCursor_StringCursor(acc, pos, False)
                break

    @staticmethod
    def ParseStringModel(s, pos):
        if ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint('"'))):
            return default__.ParseStringTail(s, (pos) + (1), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
        elif True:
            return StringCursor_StringCursor(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), pos, False)

    @staticmethod
    def ParseNumberModel(s, pos0):
        d_0_neg_ = ((pos0) < (len(s))) and (((s)[pos0]) == (_dafny.CodePoint('-')))
        d_1_first_ = ((pos0) + (1) if d_0_neg_ else pos0)
        d_2_whole_ = default__.ScanDigits(s, d_1_first_, 0, 0)
        if (((d_2_whole_).count) == (0)) or ((((d_2_whole_).count) > (1)) and (((s)[d_1_first_]) == (_dafny.CodePoint('0')))):
            return JsonCursor_JsonCursor(ConfluxCodec.Json_JNull(), (d_2_whole_).pos, False)
        elif True:
            d_3_frac_ = (default__.ScanDigits(s, ((d_2_whole_).pos) + (1), (d_2_whole_).value, 0) if (((d_2_whole_).pos) < (len(s))) and (((s)[(d_2_whole_).pos]) == (_dafny.CodePoint('.'))) else DigitCursor_DigitCursor((d_2_whole_).value, (d_2_whole_).pos, 0))
            d_4_hadDot_ = (((d_2_whole_).pos) < (len(s))) and (((s)[(d_2_whole_).pos]) == (_dafny.CodePoint('.')))
            if (d_4_hadDot_) and (((d_3_frac_).count) == (0)):
                return JsonCursor_JsonCursor(ConfluxCodec.Json_JNull(), (d_3_frac_).pos, False)
            elif True:
                d_5_expMark_ = (((d_3_frac_).pos) < (len(s))) and ((((s)[(d_3_frac_).pos]) == (_dafny.CodePoint('e'))) or (((s)[(d_3_frac_).pos]) == (_dafny.CodePoint('E'))))
                d_6_expSignPos_ = (((d_3_frac_).pos) + (1) if d_5_expMark_ else (d_3_frac_).pos)
                d_7_expNeg_ = ((d_6_expSignPos_) < (len(s))) and (((s)[d_6_expSignPos_]) == (_dafny.CodePoint('-')))
                d_8_expPlus_ = ((d_6_expSignPos_) < (len(s))) and (((s)[d_6_expSignPos_]) == (_dafny.CodePoint('+')))
                d_9_expFirst_ = ((d_6_expSignPos_) + (1) if (d_7_expNeg_) or (d_8_expPlus_) else d_6_expSignPos_)
                d_10_expDigits_ = (default__.ScanDigits(s, d_9_expFirst_, 0, 0) if d_5_expMark_ else DigitCursor_DigitCursor(0, (d_3_frac_).pos, 0))
                if (d_5_expMark_) and (((d_10_expDigits_).count) == (0)):
                    return JsonCursor_JsonCursor(ConfluxCodec.Json_JNull(), (d_10_expDigits_).pos, False)
                elif True:
                    d_11_mantissa_ = ((0) - ((d_3_frac_).value) if d_0_neg_ else (d_3_frac_).value)
                    d_12_exponent_ = ((0) - ((d_3_frac_).count)) + (((0) - ((d_10_expDigits_).value) if d_7_expNeg_ else (d_10_expDigits_).value))
                    d_13_value_ = (ConfluxCodec.Json_JNum(d_11_mantissa_) if (not(d_4_hadDot_)) and (not(d_5_expMark_)) else ConfluxCodec.Json_JDec(d_11_mantissa_, d_12_exponent_))
                    return JsonCursor_JsonCursor(d_13_value_, (d_10_expDigits_).pos, True)

    @staticmethod
    def ParseValueModel(s, pos0, budget):
        d_0_pos_ = default__.SkipWs(s, pos0)
        if ((budget) == (0)) or ((d_0_pos_) >= (len(s))):
            return JsonCursor_JsonCursor(ConfluxCodec.Json_JNull(), d_0_pos_, False)
        elif True:
            d_1_c_ = (s)[d_0_pos_]
            if (d_1_c_) == (_dafny.CodePoint('"')):
                d_2_r_ = default__.ParseStringModel(s, d_0_pos_)
                return JsonCursor_JsonCursor(ConfluxCodec.Json_JStr((d_2_r_).value), (d_2_r_).pos, (d_2_r_).ok)
            elif (d_1_c_) == (_dafny.CodePoint('{')):
                return default__.ParseObjectModel(s, d_0_pos_, budget)
            elif (d_1_c_) == (_dafny.CodePoint('[')):
                return default__.ParseArrayModel(s, d_0_pos_, budget)
            elif (((d_1_c_) == (_dafny.CodePoint('t'))) and (((d_0_pos_) + (4)) <= (len(s)))) and ((_dafny.SeqWithoutIsStrInference((s)[d_0_pos_:(d_0_pos_) + (4):])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "true")))):
                return JsonCursor_JsonCursor(ConfluxCodec.Json_JBool(True), (d_0_pos_) + (4), True)
            elif (((d_1_c_) == (_dafny.CodePoint('f'))) and (((d_0_pos_) + (5)) <= (len(s)))) and ((_dafny.SeqWithoutIsStrInference((s)[d_0_pos_:(d_0_pos_) + (5):])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "false")))):
                return JsonCursor_JsonCursor(ConfluxCodec.Json_JBool(False), (d_0_pos_) + (5), True)
            elif (((d_1_c_) == (_dafny.CodePoint('n'))) and (((d_0_pos_) + (4)) <= (len(s)))) and ((_dafny.SeqWithoutIsStrInference((s)[d_0_pos_:(d_0_pos_) + (4):])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "null")))):
                return JsonCursor_JsonCursor(ConfluxCodec.Json_JNull(), (d_0_pos_) + (4), True)
            elif ((d_1_c_) == (_dafny.CodePoint('-'))) or (default__.IsDigit(d_1_c_)):
                return default__.ParseNumberModel(s, d_0_pos_)
            elif True:
                return JsonCursor_JsonCursor(ConfluxCodec.Json_JNull(), (d_0_pos_) + (1), False)

    @staticmethod
    def ParseArrayModel(s, pos, budget):
        d_0_next_ = default__.SkipWs(s, (pos) + (1))
        if ((d_0_next_) < (len(s))) and (((s)[d_0_next_]) == (_dafny.CodePoint(']'))):
            return JsonCursor_JsonCursor(ConfluxCodec.Json_JArr(_dafny.SeqWithoutIsStrInference([])), (d_0_next_) + (1), True)
        elif True:
            return default__.ParseArrayTailModel(s, d_0_next_, budget, _dafny.SeqWithoutIsStrInference([]))

    @staticmethod
    def ParseArrayTailModel(s, pos, budget, acc):
        d_0_item_ = default__.ParseValueModel(s, pos, (budget) - (1))
        if not((d_0_item_).ok):
            return JsonCursor_JsonCursor(ConfluxCodec.Json_JArr(acc), (d_0_item_).pos, False)
        elif True:
            d_1_next_ = default__.SkipWs(s, (d_0_item_).pos)
            if ((d_1_next_) < (len(s))) and (((s)[d_1_next_]) == (_dafny.CodePoint(']'))):
                return JsonCursor_JsonCursor(ConfluxCodec.Json_JArr((acc) + (_dafny.SeqWithoutIsStrInference([(d_0_item_).value]))), (d_1_next_) + (1), True)
            elif (((d_1_next_) < (len(s))) and (((s)[d_1_next_]) == (_dafny.CodePoint(',')))) and (((d_1_next_) + (1)) > (pos)):
                return default__.ParseArrayTailModel(s, (d_1_next_) + (1), budget, (acc) + (_dafny.SeqWithoutIsStrInference([(d_0_item_).value])))
            elif True:
                return JsonCursor_JsonCursor(ConfluxCodec.Json_JArr((acc) + (_dafny.SeqWithoutIsStrInference([(d_0_item_).value]))), d_1_next_, False)

    @staticmethod
    def ParseObjectModel(s, pos, budget):
        d_0_next_ = default__.SkipWs(s, (pos) + (1))
        if ((d_0_next_) < (len(s))) and (((s)[d_0_next_]) == (_dafny.CodePoint('}'))):
            return JsonCursor_JsonCursor(ConfluxCodec.Json_JObj(_dafny.Map({})), (d_0_next_) + (1), True)
        elif True:
            return default__.ParseObjectTailModel(s, d_0_next_, budget, _dafny.Map({}))

    @staticmethod
    def ParseObjectTailModel(s, pos, budget, acc):
        d_0_next_ = default__.SkipWs(s, pos)
        if ((d_0_next_) >= (len(s))) or (((s)[d_0_next_]) != (_dafny.CodePoint('"'))):
            return JsonCursor_JsonCursor(ConfluxCodec.Json_JObj(acc), d_0_next_, False)
        elif True:
            d_1_key_ = default__.ParseStringModel(s, d_0_next_)
            d_2_colon_ = default__.SkipWs(s, (d_1_key_).pos)
            if ((not((d_1_key_).ok)) or ((d_2_colon_) >= (len(s)))) or (((s)[d_2_colon_]) != (_dafny.CodePoint(':'))):
                return JsonCursor_JsonCursor(ConfluxCodec.Json_JObj(acc), d_2_colon_, False)
            elif True:
                d_3_value_ = default__.ParseValueModel(s, (d_2_colon_) + (1), (budget) - (1))
                if not((d_3_value_).ok):
                    return JsonCursor_JsonCursor(ConfluxCodec.Json_JObj(acc), (d_3_value_).pos, False)
                elif True:
                    d_4_after_ = default__.SkipWs(s, (d_3_value_).pos)
                    d_5_fields_ = (acc).set((d_1_key_).value, (d_3_value_).value)
                    if ((d_4_after_) < (len(s))) and (((s)[d_4_after_]) == (_dafny.CodePoint('}'))):
                        return JsonCursor_JsonCursor(ConfluxCodec.Json_JObj(d_5_fields_), (d_4_after_) + (1), True)
                    elif (((d_4_after_) < (len(s))) and (((s)[d_4_after_]) == (_dafny.CodePoint(',')))) and (((d_4_after_) + (1)) > (pos)):
                        return default__.ParseObjectTailModel(s, (d_4_after_) + (1), budget, d_5_fields_)
                    elif True:
                        return JsonCursor_JsonCursor(ConfluxCodec.Json_JObj(d_5_fields_), d_4_after_, False)

    @staticmethod
    def ParseString(s, pos0):
        val: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pos: int = int(0)
        ok: bool = False
        val = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        ok = False
        pos = pos0
        if ((pos) >= (len(s))) or (((s)[pos]) != (_dafny.CodePoint('"'))):
            return val, pos, ok
        pos = (pos) + (1)
        d_0_runStart_: int
        d_0_runStart_ = pos
        while ((pos) < (len(s))) and (((s)[pos]) != (_dafny.CodePoint('"'))):
            if (((s)[pos]) == (_dafny.CodePoint('\\'))) and (((pos) + (1)) < (len(s))):
                val = (val) + (_dafny.SeqWithoutIsStrInference((s)[d_0_runStart_:pos:]))
                d_1_e_: str
                d_1_e_ = (s)[(pos) + (1)]
                if (((d_1_e_) == (_dafny.CodePoint('"'))) or ((d_1_e_) == (_dafny.CodePoint('\\')))) or ((d_1_e_) == (_dafny.CodePoint('/'))):
                    val = (val) + (_dafny.SeqWithoutIsStrInference([d_1_e_]))
                    pos = (pos) + (2)
                elif (d_1_e_) == (_dafny.CodePoint('n')):
                    val = (val) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('\n')]))
                    pos = (pos) + (2)
                elif (d_1_e_) == (_dafny.CodePoint('t')):
                    val = (val) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('\t')]))
                    pos = (pos) + (2)
                elif (d_1_e_) == (_dafny.CodePoint('r')):
                    val = (val) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('\r')]))
                    pos = (pos) + (2)
                elif (d_1_e_) == (_dafny.CodePoint('b')):
                    val = (val) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('\U00000008')]))
                    pos = (pos) + (2)
                elif (d_1_e_) == (_dafny.CodePoint('f')):
                    val = (val) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('\U0000000C')]))
                    pos = (pos) + (2)
                elif (((d_1_e_) == (_dafny.CodePoint('u'))) and (((pos) + (6)) <= (len(s)))) and ((default__.Hex4(s, (pos) + (2))) >= (0)):
                    d_2_code_: int
                    d_2_code_ = default__.Hex4(s, (pos) + (2))
                    if (((((((55296) <= (d_2_code_)) and ((d_2_code_) <= (56319))) and (((pos) + (12)) <= (len(s)))) and (((s)[(pos) + (6)]) == (_dafny.CodePoint('\\')))) and (((s)[(pos) + (7)]) == (_dafny.CodePoint('u')))) and ((default__.Hex4(s, (pos) + (8))) >= (56320))) and ((default__.Hex4(s, (pos) + (8))) <= (57343)):
                        d_3_lo_: int
                        d_3_lo_ = default__.Hex4(s, (pos) + (8))
                        val = (val) + (_dafny.SeqWithoutIsStrInference([default__.ScalarToChar(((65536) + (((d_2_code_) - (55296)) * (1024))) + ((d_3_lo_) - (56320)))]))
                        pos = (pos) + (12)
                    elif True:
                        val = (val) + (_dafny.SeqWithoutIsStrInference([default__.ScalarToChar(d_2_code_)]))
                        pos = (pos) + (6)
                elif True:
                    return val, pos, ok
                d_0_runStart_ = pos
            elif True:
                if (ord((s)[pos])) < (32):
                    return val, pos, ok
                pos = (pos) + (1)
        val = (val) + (_dafny.SeqWithoutIsStrInference((s)[d_0_runStart_:pos:]))
        if ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint('"'))):
            pos = (pos) + (1)
            ok = True
        return val, pos, ok

    @staticmethod
    def AdvanceOver(s, pos, n):
        if ((pos) + (n)) <= (len(s)):
            return (pos) + (n)
        elif True:
            return len(s)

    @staticmethod
    def ParseNumber(s, pos0):
        j: ConfluxCodec.Json = ConfluxCodec.Json.default()()
        pos: int = int(0)
        ok: bool = False
        pos = pos0
        ok = False
        d_0_neg_: bool
        d_0_neg_ = False
        if ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint('-'))):
            d_0_neg_ = True
            pos = (pos) + (1)
        d_1_whole_: int
        d_1_whole_ = 0
        d_2_sawDigit_: bool
        d_2_sawDigit_ = False
        d_3_firstDigit_: int
        d_3_firstDigit_ = pos
        while ((pos) < (len(s))) and (default__.IsDigit((s)[pos])):
            d_1_whole_ = ((d_1_whole_) * (10)) + ((ord((s)[pos])) - (ord(_dafny.CodePoint('0'))))
            d_2_sawDigit_ = True
            pos = (pos) + (1)
        d_4_mantissa_: int
        d_4_mantissa_ = d_1_whole_
        d_5_exp_: int
        d_5_exp_ = 0
        d_6_decimalDigits_: int
        d_6_decimalDigits_ = 0
        if (not(d_2_sawDigit_)) or ((((pos) - (d_3_firstDigit_)) > (1)) and (((s)[d_3_firstDigit_]) == (_dafny.CodePoint('0')))):
            j = ConfluxCodec.Json_JNull()
            return j, pos, ok
        if ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint('.'))):
            pos = (pos) + (1)
            d_7_fractionStart_: int
            d_7_fractionStart_ = pos
            while ((pos) < (len(s))) and (default__.IsDigit((s)[pos])):
                d_4_mantissa_ = ((d_4_mantissa_) * (10)) + ((ord((s)[pos])) - (ord(_dafny.CodePoint('0'))))
                d_6_decimalDigits_ = (d_6_decimalDigits_) + (1)
                d_2_sawDigit_ = True
                pos = (pos) + (1)
            if (pos) == (d_7_fractionStart_):
                j = ConfluxCodec.Json_JNull()
                return j, pos, ok
            d_5_exp_ = (0) - (d_6_decimalDigits_)
        d_8_explicitExp_: bool
        d_8_explicitExp_ = False
        if ((pos) < (len(s))) and ((((s)[pos]) == (_dafny.CodePoint('e'))) or (((s)[pos]) == (_dafny.CodePoint('E')))):
            d_8_explicitExp_ = True
            pos = (pos) + (1)
            d_9_expNeg_: bool
            d_9_expNeg_ = False
            if ((pos) < (len(s))) and ((((s)[pos]) == (_dafny.CodePoint('+'))) or (((s)[pos]) == (_dafny.CodePoint('-')))):
                d_9_expNeg_ = ((s)[pos]) == (_dafny.CodePoint('-'))
                pos = (pos) + (1)
            d_10_e_: int
            d_10_e_ = 0
            d_11_exponentStart_: int
            d_11_exponentStart_ = pos
            while ((pos) < (len(s))) and (default__.IsDigit((s)[pos])):
                d_10_e_ = ((d_10_e_) * (10)) + ((ord((s)[pos])) - (ord(_dafny.CodePoint('0'))))
                pos = (pos) + (1)
            if (pos) == (d_11_exponentStart_):
                j = ConfluxCodec.Json_JNull()
                return j, pos, ok
            d_5_exp_ = (d_5_exp_) + (((0) - (d_10_e_) if d_9_expNeg_ else d_10_e_))
        d_12_signed_: int
        if d_0_neg_:
            d_12_signed_ = (0) - (d_4_mantissa_)
        elif True:
            d_12_signed_ = d_4_mantissa_
        if not(d_2_sawDigit_):
            j = ConfluxCodec.Json_JNull()
        elif ((d_6_decimalDigits_) == (0)) and (not(d_8_explicitExp_)):
            j = ConfluxCodec.Json_JNum(d_12_signed_)
        elif True:
            j = ConfluxCodec.Json_JDec(d_12_signed_, d_5_exp_)
        ok = True
        return j, pos, ok

    @staticmethod
    def ParseValue(s, pos0, budget):
        j: ConfluxCodec.Json = ConfluxCodec.Json.default()()
        pos: int = int(0)
        ok: bool = False
        pos = default__.SkipWs(s, pos0)
        ok = False
        if ((budget) == (0)) or ((pos) >= (len(s))):
            j = ConfluxCodec.Json_JNull()
            return j, pos, ok
        d_0_c_: str
        d_0_c_ = (s)[pos]
        if (d_0_c_) == (_dafny.CodePoint('"')):
            d_1_v_: _dafny.Seq
            d_2_p_: int
            d_3_good_: bool
            out0_: _dafny.Seq
            out1_: int
            out2_: bool
            out0_, out1_, out2_ = default__.ParseString(s, pos)
            d_1_v_ = out0_
            d_2_p_ = out1_
            d_3_good_ = out2_
            j = ConfluxCodec.Json_JStr(d_1_v_)
            pos = d_2_p_
            ok = d_3_good_
        elif (d_0_c_) == (_dafny.CodePoint('{')):
            d_4_m_: _dafny.Map
            d_5_p_: int
            d_6_good_: bool
            out3_: _dafny.Map
            out4_: int
            out5_: bool
            out3_, out4_, out5_ = default__.ParseObject(s, pos, budget)
            d_4_m_ = out3_
            d_5_p_ = out4_
            d_6_good_ = out5_
            j = ConfluxCodec.Json_JObj(d_4_m_)
            pos = d_5_p_
            ok = d_6_good_
        elif (d_0_c_) == (_dafny.CodePoint('[')):
            d_7_a_: _dafny.Seq
            d_8_p_: int
            d_9_good_: bool
            out6_: _dafny.Seq
            out7_: int
            out8_: bool
            out6_, out7_, out8_ = default__.ParseArray(s, pos, budget)
            d_7_a_ = out6_
            d_8_p_ = out7_
            d_9_good_ = out8_
            j = ConfluxCodec.Json_JArr(d_7_a_)
            pos = d_8_p_
            ok = d_9_good_
        elif (((d_0_c_) == (_dafny.CodePoint('t'))) and (((pos) + (4)) <= (len(s)))) and ((_dafny.SeqWithoutIsStrInference((s)[pos:(pos) + (4):])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "true")))):
            j = ConfluxCodec.Json_JBool(True)
            pos = (pos) + (4)
            ok = True
        elif (((d_0_c_) == (_dafny.CodePoint('f'))) and (((pos) + (5)) <= (len(s)))) and ((_dafny.SeqWithoutIsStrInference((s)[pos:(pos) + (5):])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "false")))):
            j = ConfluxCodec.Json_JBool(False)
            pos = (pos) + (5)
            ok = True
        elif (((d_0_c_) == (_dafny.CodePoint('n'))) and (((pos) + (4)) <= (len(s)))) and ((_dafny.SeqWithoutIsStrInference((s)[pos:(pos) + (4):])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "null")))):
            j = ConfluxCodec.Json_JNull()
            pos = (pos) + (4)
            ok = True
        elif ((d_0_c_) == (_dafny.CodePoint('-'))) or (default__.IsDigit(d_0_c_)):
            d_10_n_: ConfluxCodec.Json
            d_11_p_: int
            d_12_good_: bool
            out9_: ConfluxCodec.Json
            out10_: int
            out11_: bool
            out9_, out10_, out11_ = default__.ParseNumber(s, pos)
            d_10_n_ = out9_
            d_11_p_ = out10_
            d_12_good_ = out11_
            j = d_10_n_
            pos = d_11_p_
            ok = d_12_good_
        elif True:
            j = ConfluxCodec.Json_JNull()
            pos = (pos) + (1)
        return j, pos, ok

    @staticmethod
    def ParseObject(s, pos0, budget):
        m: _dafny.Map = _dafny.Map({})
        pos: int = int(0)
        ok: bool = False
        m = _dafny.Map({})
        pos = pos0
        ok = False
        if (budget) == (0):
            return m, pos, ok
        if ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint('{'))):
            pos = (pos) + (1)
        pos = default__.SkipWs(s, pos)
        if ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint('}'))):
            pos = (pos) + (1)
            ok = True
            return m, pos, ok
        with _dafny.label("0"):
            while (pos) < (len(s)):
                with _dafny.c_label("0"):
                    d_0_before_: int
                    d_0_before_ = pos
                    pos = default__.SkipWs(s, pos)
                    if ((pos) >= (len(s))) or (((s)[pos]) != (_dafny.CodePoint('"'))):
                        raise _dafny.Break("0")
                    d_1_key_: _dafny.Seq
                    d_2_p1_: int
                    d_3_keyOk_: bool
                    out0_: _dafny.Seq
                    out1_: int
                    out2_: bool
                    out0_, out1_, out2_ = default__.ParseString(s, pos)
                    d_1_key_ = out0_
                    d_2_p1_ = out1_
                    d_3_keyOk_ = out2_
                    pos = d_2_p1_
                    if not(d_3_keyOk_):
                        return m, pos, ok
                    pos = default__.SkipWs(s, pos)
                    if ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint(':'))):
                        pos = (pos) + (1)
                    elif True:
                        raise _dafny.Break("0")
                    d_4_v_: ConfluxCodec.Json
                    d_5_p2_: int
                    d_6_valueOk_: bool
                    out3_: ConfluxCodec.Json
                    out4_: int
                    out5_: bool
                    out3_, out4_, out5_ = default__.ParseValue(s, pos, (budget) - (1))
                    d_4_v_ = out3_
                    d_5_p2_ = out4_
                    d_6_valueOk_ = out5_
                    pos = d_5_p2_
                    if not(d_6_valueOk_):
                        return m, pos, ok
                    m = (m).set(d_1_key_, d_4_v_)
                    pos = default__.SkipWs(s, pos)
                    if ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint(','))):
                        pos = (pos) + (1)
                    elif ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint('}'))):
                        pos = (pos) + (1)
                        ok = True
                        raise _dafny.Break("0")
                    elif True:
                        raise _dafny.Break("0")
                    if (pos) <= (d_0_before_):
                        raise _dafny.Break("0")
                    pass
            pass
        return m, pos, ok

    @staticmethod
    def ParseArray(s, pos0, budget):
        a: _dafny.Seq = _dafny.Seq({})
        pos: int = int(0)
        ok: bool = False
        a = _dafny.SeqWithoutIsStrInference([])
        pos = pos0
        ok = False
        if (budget) == (0):
            return a, pos, ok
        if ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint('['))):
            pos = (pos) + (1)
        pos = default__.SkipWs(s, pos)
        if ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint(']'))):
            pos = (pos) + (1)
            ok = True
            return a, pos, ok
        with _dafny.label("1"):
            while (pos) < (len(s)):
                with _dafny.c_label("1"):
                    d_0_before_: int
                    d_0_before_ = pos
                    d_1_v_: ConfluxCodec.Json
                    d_2_p_: int
                    d_3_valueOk_: bool
                    out0_: ConfluxCodec.Json
                    out1_: int
                    out2_: bool
                    out0_, out1_, out2_ = default__.ParseValue(s, pos, (budget) - (1))
                    d_1_v_ = out0_
                    d_2_p_ = out1_
                    d_3_valueOk_ = out2_
                    pos = d_2_p_
                    if not(d_3_valueOk_):
                        return a, pos, ok
                    a = (a) + (_dafny.SeqWithoutIsStrInference([d_1_v_]))
                    pos = default__.SkipWs(s, pos)
                    if ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint(','))):
                        pos = (pos) + (1)
                    elif ((pos) < (len(s))) and (((s)[pos]) == (_dafny.CodePoint(']'))):
                        pos = (pos) + (1)
                        ok = True
                        raise _dafny.Break("1")
                    elif True:
                        raise _dafny.Break("1")
                    if (pos) <= (d_0_before_):
                        raise _dafny.Break("1")
                    pass
            pass
        return a, pos, ok

    @staticmethod
    def ParseJsonModel(s):
        d_0_parsed_ = default__.ParseValueModel(s, 0, default__.PARSE__BUDGET)
        d_1_end_ = default__.SkipWs(s, (d_0_parsed_).pos)
        if ((d_0_parsed_).ok) and ((d_1_end_) == (len(s))):
            return ParseResult_Parsed((d_0_parsed_).value)
        elif True:
            return ParseResult_Invalid()

    @staticmethod
    def ParseJsonChecked(s):
        result: ParseResult = ParseResult.default()()
        result = default__.ParseJsonModel(s)
        return result

    @staticmethod
    def ParseJson(s):
        j: ConfluxCodec.Json = ConfluxCodec.Json.default()()
        d_0_result_: ParseResult
        out0_: ParseResult
        out0_ = default__.ParseJsonChecked(s)
        d_0_result_ = out0_
        if (d_0_result_).is_Parsed:
            j = (d_0_result_).value
        elif True:
            j = ConfluxCodec.Json_JNull()
        return j

    @staticmethod
    def CharLt(a, b):
        return (ord(a)) < (ord(b))

    @staticmethod
    def StrLt(a, b):
        while True:
            with _dafny.label():
                if (len(a)) == (0):
                    return (len(b)) > (0)
                elif (len(b)) == (0):
                    return False
                elif ((a)[0]) == ((b)[0]):
                    in0_ = _dafny.SeqWithoutIsStrInference((a)[1::])
                    in1_ = _dafny.SeqWithoutIsStrInference((b)[1::])
                    a = in0_
                    b = in1_
                    raise _dafny.TailCall()
                elif True:
                    return default__.CharLt((a)[0], (b)[0])
                break

    @staticmethod
    def Sorted(xs):
        def lambda0_(forall_var_0_):
            d_0_i_: int = forall_var_0_
            return not (((0) <= (d_0_i_)) and (((d_0_i_) + (1)) < (len(xs)))) or (default__.StrLt((xs)[d_0_i_], (xs)[(d_0_i_) + (1)]))

        return _dafny.quantifier(_dafny.IntegerRange(0, (len(xs)) - (1)), True, lambda0_)

    @staticmethod
    def Elems(xs):
        d_0___accumulator_ = _dafny.Set({})
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (_dafny.Set({})) | (d_0___accumulator_)
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) | (_dafny.Set({(xs)[0]}))
                    in0_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    xs = in0_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def SortedKeySetRuntime(s):
        keys: _dafny.Seq = _dafny.Seq({})
        if (s) == (_dafny.Set({})):
            keys = _dafny.SeqWithoutIsStrInference([])
        elif True:
            d_0_k_: _dafny.Seq
            with _dafny.label("_ASSIGN_SUCH_THAT_d_0"):
                assign_such_that_0_: _dafny.Seq
                for assign_such_that_0_ in (s).Elements:
                    d_0_k_ = assign_such_that_0_
                    def lambda0_(forall_var_0_):
                        d_1_y_: _dafny.Seq = forall_var_0_
                        return not ((d_1_y_) in (s)) or (not(default__.StrLt(d_1_y_, d_0_k_)))

                    if ((d_0_k_) in (s)) and (_dafny.quantifier((s).Elements, True, lambda0_)):
                        raise _dafny.Break("_ASSIGN_SUCH_THAT_d_0")
                raise Exception("assign-such-that search produced no value")
                pass
            d_2_tail_: _dafny.Seq
            out0_: _dafny.Seq
            out0_ = default__.SortedKeySetRuntime((s) - (_dafny.Set({d_0_k_})))
            d_2_tail_ = out0_
            keys = (_dafny.SeqWithoutIsStrInference([d_0_k_])) + (d_2_tail_)
        return keys

    @staticmethod
    def InsertKey(k, xs):
        ys: _dafny.Seq = _dafny.Seq({})
        if (len(xs)) == (0):
            ys = _dafny.SeqWithoutIsStrInference([k])
        elif default__.StrLt(k, (xs)[0]):
            ys = (_dafny.SeqWithoutIsStrInference([k])) + (xs)
        elif True:
            d_0_tail_: _dafny.Seq
            out0_: _dafny.Seq
            out0_ = default__.InsertKey(k, _dafny.SeqWithoutIsStrInference((xs)[1::]))
            d_0_tail_ = out0_
            ys = (_dafny.SeqWithoutIsStrInference([(xs)[0]])) + (d_0_tail_)
        return ys

    @staticmethod
    def SortedKeys(m):
        keys: _dafny.Seq = _dafny.Seq({})
        out0_: _dafny.Seq
        out0_ = default__.SortedKeySetRuntime((m).keys)
        keys = out0_
        return keys

    @staticmethod
    def IntText(n):
        if (n) < (0):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "-"))) + (ConfluxDecimalText.default__.NatText((0) - (n)))
        elif True:
            return ConfluxDecimalText.default__.NatText(n)

    @staticmethod
    def EscapeChar(c):
        if (c) == (_dafny.CodePoint('"')):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\\\""))
        elif (c) == (_dafny.CodePoint('\\')):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\\\\"))
        elif (c) == (_dafny.CodePoint('\n')):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\\n"))
        elif (c) == (_dafny.CodePoint('\r')):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\\r"))
        elif (c) == (_dafny.CodePoint('\t')):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\\t"))
        elif (c) == (_dafny.CodePoint('\U00000008')):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\\b"))
        elif (c) == (_dafny.CodePoint('\U0000000C')):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\\f"))
        elif (ord(c)) < (32):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\\u00"))) + (_dafny.SeqWithoutIsStrInference([default__.HexDigit(_dafny.euclidian_division(ord(c), 16)), default__.HexDigit(_dafny.euclidian_modulus(ord(c), 16))]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([c])

    @staticmethod
    def Escape(s):
        hresult_: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        if (len(s)) == (0):
            hresult_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
            return hresult_
        d_0_n_: int
        d_0_n_ = len(s)
        d_1_arr_: _dafny.Array
        nw0_ = _dafny.Array(_dafny.CodePoint('D'), (6) * (d_0_n_))
        d_1_arr_ = nw0_
        d_2_off_: int
        d_2_off_ = 0
        d_3_i_: int
        d_3_i_ = 0
        while (d_3_i_) < (d_0_n_):
            d_4_e_: _dafny.Seq
            d_4_e_ = default__.EscapeChar((s)[d_3_i_])
            d_5_j_: int
            d_5_j_ = 0
            while (d_5_j_) < (len(d_4_e_)):
                index0_ = (d_2_off_) + (d_5_j_)
                (d_1_arr_)[index0_] = (d_4_e_)[d_5_j_]
                d_5_j_ = (d_5_j_) + (1)
            d_2_off_ = (d_2_off_) + (len(d_4_e_))
            d_3_i_ = (d_3_i_) + (1)
        d_6_result_: _dafny.Seq
        d_6_result_ = _dafny.SeqWithoutIsStrInference([(d_1_arr_)[d_7_k_] for d_7_k_ in range(d_2_off_)])
        hresult_ = d_6_result_
        return hresult_
        return hresult_

    @staticmethod
    def AllDigits(ds):
        def lambda0_(forall_var_0_):
            d_0_i_: int = forall_var_0_
            return not (((0) <= (d_0_i_)) and ((d_0_i_) < (len(ds)))) or (default__.IsDigit((ds)[d_0_i_]))

        return _dafny.quantifier(_dafny.IntegerRange(0, len(ds)), True, lambda0_)

    @staticmethod
    def FoldDigits(ds, acc):
        while True:
            with _dafny.label():
                if (len(ds)) == (0):
                    return acc
                elif True:
                    in0_ = _dafny.SeqWithoutIsStrInference((ds)[1::])
                    in1_ = ((acc) * (10)) + ((ord((ds)[0])) - (ord(_dafny.CodePoint('0'))))
                    ds = in0_
                    acc = in1_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def DigitStep(acc, d):
        return ((acc) * (10)) + ((ord(d)) - (ord(_dafny.CodePoint('0'))))

    @staticmethod
    def ValueStop(suffix):
        return ((((len(suffix)) == (0)) or (((suffix)[0]) == (_dafny.CodePoint(',')))) or (((suffix)[0]) == (_dafny.CodePoint(']')))) or (((suffix)[0]) == (_dafny.CodePoint('}')))

    @staticmethod
    def SignLen(n):
        if (n) < (0):
            return 1
        elif True:
            return 0

    @staticmethod
    def Magnitude(n):
        if (n) < (0):
            return (0) - (n)
        elif True:
            return n

    @staticmethod
    def StringifySeqRuntime(xs, budget):
        s: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        if (len(xs)) == (0):
            s = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        elif (len(xs)) == (1):
            out0_: _dafny.Seq
            out0_ = default__.StringifyValue((xs)[0], budget)
            s = out0_
        elif True:
            d_0_split_: int
            d_0_split_ = _dafny.euclidian_division(len(xs), 2)
            d_1_left_: _dafny.Seq
            out1_: _dafny.Seq
            out1_ = default__.StringifySeqRuntime(_dafny.SeqWithoutIsStrInference((xs)[:d_0_split_:]), budget)
            d_1_left_ = out1_
            d_2_right_: _dafny.Seq
            out2_: _dafny.Seq
            out2_ = default__.StringifySeqRuntime(_dafny.SeqWithoutIsStrInference((xs)[d_0_split_::]), budget)
            d_2_right_ = out2_
            s = ((d_1_left_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ",")))) + (d_2_right_)
        return s

    @staticmethod
    def StringifyFieldsRuntime(fields, keys, budget):
        s: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        if (len(keys)) == (0):
            s = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        elif True:
            d_0_value_: _dafny.Seq
            out0_: _dafny.Seq
            out0_ = default__.StringifyValue((fields)[(keys)[0]], budget)
            d_0_value_ = out0_
            d_1_head_: _dafny.Seq
            d_1_head_ = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\""))) + (default__.Escape((keys)[0]))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\":")))) + (d_0_value_)
            if (len(keys)) == (1):
                s = d_1_head_
            elif True:
                d_2_tail_: _dafny.Seq
                out1_: _dafny.Seq
                out1_ = default__.StringifyFieldsRuntime(fields, _dafny.SeqWithoutIsStrInference((keys)[1::]), budget)
                d_2_tail_ = out1_
                s = ((d_1_head_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ",")))) + (d_2_tail_)
        return s

    @staticmethod
    def StringifyValue(j, budget):
        s: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        if (budget) == (0):
            s = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "null"))
            return s
        if (j).is_JNull:
            s = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "null"))
            return s
        if (j).is_JBool:
            if (j).b:
                s = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "true"))
            elif True:
                s = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "false"))
            return s
        if (j).is_JNum:
            s = default__.IntText((j).n)
            return s
        if (j).is_JDec:
            s = ((default__.IntText((j).mantissa)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")))) + (default__.IntText((j).exp))
            return s
        if (j).is_JStr:
            s = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\""))) + (default__.Escape((j).s))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\"")))
            return s
        if (j).is_JArr:
            d_0_body_: _dafny.Seq
            out0_: _dafny.Seq
            out0_ = default__.StringifySeqRuntime((j).elems, (budget) - (1))
            d_0_body_ = out0_
            s = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "["))) + (d_0_body_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "]")))
            return s
        d_1_keys_: _dafny.Seq
        out1_: _dafny.Seq
        out1_ = default__.SortedKeys((j).fields)
        d_1_keys_ = out1_
        d_2_body_: _dafny.Seq
        out2_: _dafny.Seq
        out2_ = default__.StringifyFieldsRuntime((j).fields, d_1_keys_, (budget) - (1))
        d_2_body_ = out2_
        s = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "{"))) + (d_2_body_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "}")))
        return s

    @staticmethod
    def Stringify(j):
        s: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        out0_: _dafny.Seq
        out0_ = default__.StringifyValue(j, default__.STRINGIFY__BUDGET)
        s = out0_
        return s

    @_dafny.classproperty
    def PARSE__BUDGET(instance):
        return 4000
    @_dafny.classproperty
    def STRINGIFY__BUDGET(instance):
        return 4000

class StringCursor:
    @classmethod
    def default(cls, ):
        return lambda: StringCursor_StringCursor(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_StringCursor(self) -> bool:
        return isinstance(self, StringCursor_StringCursor)

class StringCursor_StringCursor(StringCursor, NamedTuple('StringCursor', [('value', Any), ('pos', Any), ('ok', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxJsonText.StringCursor.StringCursor({self.value.VerbatimString(True)}, {_dafny.string_of(self.pos)}, {_dafny.string_of(self.ok)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, StringCursor_StringCursor) and self.value == __o.value and self.pos == __o.pos and self.ok == __o.ok
    def __hash__(self) -> int:
        return super().__hash__()


class JsonCursor:
    @classmethod
    def default(cls, ):
        return lambda: JsonCursor_JsonCursor(ConfluxCodec.Json.default()(), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_JsonCursor(self) -> bool:
        return isinstance(self, JsonCursor_JsonCursor)

class JsonCursor_JsonCursor(JsonCursor, NamedTuple('JsonCursor', [('value', Any), ('pos', Any), ('ok', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxJsonText.JsonCursor.JsonCursor({_dafny.string_of(self.value)}, {_dafny.string_of(self.pos)}, {_dafny.string_of(self.ok)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, JsonCursor_JsonCursor) and self.value == __o.value and self.pos == __o.pos and self.ok == __o.ok
    def __hash__(self) -> int:
        return super().__hash__()


class DigitCursor:
    @classmethod
    def default(cls, ):
        return lambda: DigitCursor_DigitCursor(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DigitCursor(self) -> bool:
        return isinstance(self, DigitCursor_DigitCursor)

class DigitCursor_DigitCursor(DigitCursor, NamedTuple('DigitCursor', [('value', Any), ('pos', Any), ('count', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxJsonText.DigitCursor.DigitCursor({_dafny.string_of(self.value)}, {_dafny.string_of(self.pos)}, {_dafny.string_of(self.count)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, DigitCursor_DigitCursor) and self.value == __o.value and self.pos == __o.pos and self.count == __o.count
    def __hash__(self) -> int:
        return super().__hash__()


class ParseResult:
    @classmethod
    def default(cls, ):
        return lambda: ParseResult_Invalid()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Invalid(self) -> bool:
        return isinstance(self, ParseResult_Invalid)
    @property
    def is_Parsed(self) -> bool:
        return isinstance(self, ParseResult_Parsed)

class ParseResult_Invalid(ParseResult, NamedTuple('Invalid', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxJsonText.ParseResult.Invalid'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ParseResult_Invalid)
    def __hash__(self) -> int:
        return super().__hash__()

class ParseResult_Parsed(ParseResult, NamedTuple('Parsed', [('value', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxJsonText.ParseResult.Parsed({_dafny.string_of(self.value)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ParseResult_Parsed) and self.value == __o.value
    def __hash__(self) -> int:
        return super().__hash__()

