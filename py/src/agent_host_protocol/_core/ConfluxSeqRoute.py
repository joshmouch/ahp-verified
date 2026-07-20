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

# Module: ConfluxSeqRoute

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def SeqUpsertById(keyOf, xs, v):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([v]))
                elif (keyOf((xs)[0])) == (keyOf(v)):
                    return (d_0___accumulator_) + ((_dafny.SeqWithoutIsStrInference([v])) + (_dafny.SeqWithoutIsStrInference((xs)[1::])))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([(xs)[0]]))
                    in0_ = keyOf
                    in1_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    in2_ = v
                    keyOf = in0_
                    xs = in1_
                    v = in2_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def SeqRemoveById(keyOf, xs, k):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif (keyOf((xs)[0])) == (k):
                    in0_ = keyOf
                    in1_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    in2_ = k
                    keyOf = in0_
                    xs = in1_
                    k = in2_
                    raise _dafny.TailCall()
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([(xs)[0]]))
                    in3_ = keyOf
                    in4_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    in5_ = k
                    keyOf = in3_
                    xs = in4_
                    k = in5_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def ToMap(keyOf, xs):
        if (len(xs)) == (0):
            return _dafny.Map({})
        elif True:
            return (default__.ToMap(keyOf, _dafny.SeqWithoutIsStrInference((xs)[1::]))).set(keyOf((xs)[0]), (xs)[0])

    @staticmethod
    def UniqueKeys(keyOf, xs):
        def lambda0_(forall_var_0_):
            def lambda1_(forall_var_1_):
                d_1_j_: int = forall_var_1_
                return not (((((0) <= (d_0_i_)) and ((d_0_i_) < (len(xs)))) and (((0) <= (d_1_j_)) and ((d_1_j_) < (len(xs))))) and ((keyOf((xs)[d_0_i_])) == (keyOf((xs)[d_1_j_])))) or ((d_0_i_) == (d_1_j_))

            d_0_i_: int = forall_var_0_
            return _dafny.quantifier(_dafny.IntegerRange(0, len(xs)), True, lambda1_)

        return _dafny.quantifier(_dafny.IntegerRange(0, len(xs)), True, lambda0_)

    @staticmethod
    def InstallOp(cur, v):
        return ConfluxPrune.RouteOp_Install(v)

    @staticmethod
    def RemoveOp(cur, v):
        return ConfluxPrune.RouteOp_Remove()

    @staticmethod
    def SeqUpdateById(keyOf, xs, k, f):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif (keyOf((xs)[0])) == (k):
                    return (d_0___accumulator_) + ((_dafny.SeqWithoutIsStrInference([f((xs)[0])])) + (_dafny.SeqWithoutIsStrInference((xs)[1::])))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([(xs)[0]]))
                    in0_ = keyOf
                    in1_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    in2_ = k
                    in3_ = f
                    keyOf = in0_
                    xs = in1_
                    k = in2_
                    f = in3_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def UpdateOp(f):
        def lambda0_(d_0_f_):
            def lambda1_(d_1_cur_, d_2_a_):
                def lambda2_():
                    source0_ = d_1_cur_
                    if True:
                        if source0_.is_Some:
                            d_3_v_ = source0_.value
                            return ConfluxPrune.RouteOp_Install(d_0_f_(d_3_v_))
                    if True:
                        return ConfluxPrune.RouteOp_NoOp()

                return lambda2_()

            return lambda1_

        return lambda0_(f)

    @staticmethod
    def SeqUpdateByIdWhere(keyOf, pred, xs, k, f):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif ((keyOf((xs)[0])) == (k)) and (pred((xs)[0])):
                    return (d_0___accumulator_) + ((_dafny.SeqWithoutIsStrInference([f((xs)[0])])) + (_dafny.SeqWithoutIsStrInference((xs)[1::])))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([(xs)[0]]))
                    in0_ = keyOf
                    in1_ = pred
                    in2_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    in3_ = k
                    in4_ = f
                    keyOf = in0_
                    pred = in1_
                    xs = in2_
                    k = in3_
                    f = in4_
                    raise _dafny.TailCall()
                break

    @staticmethod
    def UpdateWhereOp(pred, f):
        def lambda0_(d_0_pred_, d_1_f_):
            def lambda1_(d_2_cur_, d_3_a_):
                def lambda2_():
                    source0_ = d_2_cur_
                    if True:
                        if source0_.is_Some:
                            d_4_v_ = source0_.value
                            if d_0_pred_(d_4_v_):
                                return ConfluxPrune.RouteOp_Install(d_1_f_(d_4_v_))
                            elif True:
                                return ConfluxPrune.RouteOp_NoOp()
                    if True:
                        return ConfluxPrune.RouteOp_NoOp()

                return lambda2_()

            return lambda1_

        return lambda0_(pred, f)

    @staticmethod
    def SeqUpdateAllWhere(keyOf, pred, xs, k, f):
        d_0___accumulator_ = _dafny.SeqWithoutIsStrInference([])
        while True:
            with _dafny.label():
                if (len(xs)) == (0):
                    return (d_0___accumulator_) + (_dafny.SeqWithoutIsStrInference([]))
                elif True:
                    d_0___accumulator_ = (d_0___accumulator_) + ((_dafny.SeqWithoutIsStrInference([f((xs)[0])]) if ((keyOf((xs)[0])) == (k)) and (pred((xs)[0])) else _dafny.SeqWithoutIsStrInference([(xs)[0]])))
                    in0_ = keyOf
                    in1_ = pred
                    in2_ = _dafny.SeqWithoutIsStrInference((xs)[1::])
                    in3_ = k
                    in4_ = f
                    keyOf = in0_
                    pred = in1_
                    xs = in2_
                    k = in3_
                    f = in4_
                    raise _dafny.TailCall()
                break

