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

# Module: ConfluxSemanticGraphContract

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def CompleteWitness(c):
        return not ((c).complete) or ((len((c).reasons)) == (0))

    @staticmethod
    def TypedRecordWellFormed(value):
        source0_ = value
        if True:
            if source0_.is_TVList:
                d_0_items_ = source0_.items
                def lambda0_(forall_var_0_):
                    d_1_i_: int = forall_var_0_
                    return not (((0) <= (d_1_i_)) and ((d_1_i_) < (len(d_0_items_)))) or (default__.TypedRecordWellFormed((d_0_items_)[d_1_i_]))

                return _dafny.quantifier(_dafny.IntegerRange(0, len(d_0_items_)), True, lambda0_)
        if True:
            if source0_.is_TVRecord:
                d_2_names_ = source0_.fieldNames
                d_3_values_ = source0_.fieldValues
                def lambda1_(forall_var_1_):
                    def lambda2_(forall_var_2_):
                        d_5_j_: int = forall_var_2_
                        return not ((((0) <= (d_4_i_)) and ((d_4_i_) < (d_5_j_))) and ((d_5_j_) < (len(d_2_names_)))) or (((d_2_names_)[d_4_i_]) != ((d_2_names_)[d_5_j_]))

                    d_4_i_: int = forall_var_1_
                    return _dafny.quantifier(_dafny.IntegerRange((d_4_i_) + (1), len(d_2_names_)), True, lambda2_)

                def lambda3_(forall_var_3_):
                    d_6_i_: int = forall_var_3_
                    return not (((0) <= (d_6_i_)) and ((d_6_i_) < (len(d_3_values_)))) or (default__.TypedRecordWellFormed((d_3_values_)[d_6_i_]))

                return (((len(d_2_names_)) == (len(d_3_values_))) and (_dafny.quantifier(_dafny.IntegerRange(0, len(d_2_names_)), True, lambda1_))) and (_dafny.quantifier(_dafny.IntegerRange(0, len(d_3_values_)), True, lambda3_))
        if True:
            return True

    @staticmethod
    def GraphReferencesClose(graph):
        pat_let_tv0_ = graph
        pat_let_tv1_ = graph
        pat_let_tv2_ = graph
        pat_let_tv3_ = graph
        d_0_nodeIds_ = ((graph).nodes).keys
        d_1_edgeIds_ = ((graph).edges).keys
        d_2_factIds_ = ((graph).facts).keys
        d_3_findingIds_ = ((graph).findings).keys
        d_4_evidenceIds_ = ((graph).evidence).keys
        def lambda0_(forall_var_0_):
            def iife0_(_pat_let464_0):
                def iife1_(d_6_node_):
                    return ((((d_6_node_).id_) == (d_5_id_)) and (((d_6_node_).factIds).issubset(d_2_factIds_))) and (((d_6_node_).evidenceIds).issubset(d_4_evidenceIds_))
                return iife1_(_pat_let464_0)
            d_5_id_: _dafny.Seq = forall_var_0_
            return not ((d_5_id_) in (d_0_nodeIds_)) or (iife0_(((graph).nodes)[d_5_id_]))

        def lambda1_(forall_var_1_):
            def iife2_(_pat_let465_0):
                def iife3_(d_8_edge_):
                    return ((((((d_8_edge_).id_) == (d_7_id_)) and (((d_8_edge_).fromId) in ((pat_let_tv0_).nodes))) and (((d_8_edge_).toId) in ((pat_let_tv1_).nodes))) and (((d_8_edge_).factIds).issubset(d_2_factIds_))) and (((d_8_edge_).evidenceIds).issubset(d_4_evidenceIds_))
                return iife3_(_pat_let465_0)
            d_7_id_: _dafny.Seq = forall_var_1_
            return not ((d_7_id_) in (d_1_edgeIds_)) or (iife2_(((graph).edges)[d_7_id_]))

        def lambda2_(forall_var_2_):
            def iife4_(_pat_let466_0):
                def iife5_(d_10_fact_):
                    return ((((((d_10_fact_).id_) == (d_9_id_)) and (((d_10_fact_).subjectId) in ((pat_let_tv2_).nodes))) and (((d_10_fact_).supportingFactIds).issubset(d_2_factIds_))) and (((d_10_fact_).evidenceIds).issubset(d_4_evidenceIds_))) and (default__.TypedRecordWellFormed((d_10_fact_).value))
                return iife5_(_pat_let466_0)
            d_9_id_: _dafny.Seq = forall_var_2_
            return not ((d_9_id_) in (d_2_factIds_)) or (iife4_(((graph).facts)[d_9_id_]))

        def lambda3_(forall_var_3_):
            def iife6_(_pat_let467_0):
                def iife7_(d_12_finding_):
                    return (((((d_12_finding_).id_) == (d_11_id_)) and ((((d_12_finding_).subjectId) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))) or (((d_12_finding_).subjectId) in ((pat_let_tv3_).nodes)))) and (((d_12_finding_).supportingFactIds).issubset(d_2_factIds_))) and (((d_12_finding_).evidenceIds).issubset(d_4_evidenceIds_))
                return iife7_(_pat_let467_0)
            d_11_id_: _dafny.Seq = forall_var_3_
            return not ((d_11_id_) in (d_3_findingIds_)) or (iife6_(((graph).findings)[d_11_id_]))

        def lambda4_(forall_var_4_):
            def iife8_(_pat_let468_0):
                def iife9_(d_14_item_):
                    return (((d_14_item_).id_) == (d_13_id_)) and (((d_14_item_).parentEvidenceIds).issubset(d_4_evidenceIds_))
                return iife9_(_pat_let468_0)
            d_13_id_: _dafny.Seq = forall_var_4_
            return not ((d_13_id_) in (d_4_evidenceIds_)) or (iife8_(((graph).evidence)[d_13_id_]))

        return ((((_dafny.quantifier((d_0_nodeIds_).Elements, True, lambda0_)) and (_dafny.quantifier((d_1_edgeIds_).Elements, True, lambda1_))) and (_dafny.quantifier((d_2_factIds_).Elements, True, lambda2_))) and (_dafny.quantifier((d_3_findingIds_).Elements, True, lambda3_))) and (_dafny.quantifier((d_4_evidenceIds_).Elements, True, lambda4_))

    @staticmethod
    def GraphWellFormed(graph):
        return ((((((((graph).schemaVersion) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "semantic-graph/1.0.0")))) and (((graph).graphId) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))))) and (((graph).revision) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))))) and (((graph).snapshotId) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))))) and (((graph).receiptId) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))))) and (default__.CompleteWitness((graph).completeness))) and (default__.GraphReferencesClose(graph))

    @staticmethod
    def NonResetDeltaWellFormed(delta):
        return (((((((not((delta).reset)) and (((delta).resetReason).is_None)) and (((delta).replacement).is_None)) and (((((delta).nodes).added).keys).isdisjoint((((delta).nodes).changed).keys))) and (((((delta).edges).added).keys).isdisjoint((((delta).edges).changed).keys))) and (((((delta).facts).added).keys).isdisjoint((((delta).facts).changed).keys))) and (((((delta).findings).added).keys).isdisjoint((((delta).findings).changed).keys))) and (((((delta).evidence).added).keys).isdisjoint((((delta).evidence).changed).keys))

    @staticmethod
    def ResetDeltaWellFormed(delta):
        return ((((((((delta).reset) and (((delta).resetReason).is_Some)) and (((delta).replacement).is_Some)) and ((((len(((delta).nodes).added)) + (len(((delta).nodes).changed))) + (len(((delta).nodes).removed))) == (0))) and ((((len(((delta).edges).added)) + (len(((delta).edges).changed))) + (len(((delta).edges).removed))) == (0))) and ((((len(((delta).facts).added)) + (len(((delta).facts).changed))) + (len(((delta).facts).removed))) == (0))) and ((((len(((delta).findings).added)) + (len(((delta).findings).changed))) + (len(((delta).findings).removed))) == (0))) and ((((len(((delta).evidence).added)) + (len(((delta).evidence).changed))) + (len(((delta).evidence).removed))) == (0))

    @staticmethod
    def DeltaWellFormed(delta):
        return (((((((delta).schemaVersion) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "semantic-graph-delta/1.0.0")))) and (((delta).graphId) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))))) and (((delta).fromRevision) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))))) and (((delta).toRevision) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))))) and (((delta).receiptId) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))))) and ((default__.ResetDeltaWellFormed(delta) if (delta).reset else default__.NonResetDeltaWellFormed(delta)))

    @staticmethod
    def DeltaApplies(before, delta, after):
        return ((((((((default__.GraphWellFormed(before)) and (default__.GraphWellFormed(after))) and (default__.DeltaWellFormed(delta))) and (not((delta).reset))) and (((before).graphId) == ((delta).graphId))) and (((after).graphId) == ((delta).graphId))) and (((before).revision) == ((delta).fromRevision))) and (((after).revision) == ((delta).toRevision))) and (((after).previousRevision) == ((before).revision))


class SelectorKind:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [SelectorKind_Workspace(), SelectorKind_RepositorySet(), SelectorKind_Repository(), SelectorKind_Project(), SelectorKind_File(), SelectorKind_Subject()]
    @classmethod
    def default(cls, ):
        return lambda: SelectorKind_Workspace()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Workspace(self) -> bool:
        return isinstance(self, SelectorKind_Workspace)
    @property
    def is_RepositorySet(self) -> bool:
        return isinstance(self, SelectorKind_RepositorySet)
    @property
    def is_Repository(self) -> bool:
        return isinstance(self, SelectorKind_Repository)
    @property
    def is_Project(self) -> bool:
        return isinstance(self, SelectorKind_Project)
    @property
    def is_File(self) -> bool:
        return isinstance(self, SelectorKind_File)
    @property
    def is_Subject(self) -> bool:
        return isinstance(self, SelectorKind_Subject)

class SelectorKind_Workspace(SelectorKind, NamedTuple('Workspace', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SelectorKind.Workspace'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SelectorKind_Workspace)
    def __hash__(self) -> int:
        return super().__hash__()

class SelectorKind_RepositorySet(SelectorKind, NamedTuple('RepositorySet', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SelectorKind.RepositorySet'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SelectorKind_RepositorySet)
    def __hash__(self) -> int:
        return super().__hash__()

class SelectorKind_Repository(SelectorKind, NamedTuple('Repository', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SelectorKind.Repository'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SelectorKind_Repository)
    def __hash__(self) -> int:
        return super().__hash__()

class SelectorKind_Project(SelectorKind, NamedTuple('Project', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SelectorKind.Project'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SelectorKind_Project)
    def __hash__(self) -> int:
        return super().__hash__()

class SelectorKind_File(SelectorKind, NamedTuple('File', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SelectorKind.File'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SelectorKind_File)
    def __hash__(self) -> int:
        return super().__hash__()

class SelectorKind_Subject(SelectorKind, NamedTuple('Subject', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SelectorKind.Subject'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SelectorKind_Subject)
    def __hash__(self) -> int:
        return super().__hash__()


class SemanticStatus:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [SemanticStatus_Complete(), SemanticStatus_Failed(), SemanticStatus_Unavailable(), SemanticStatus_Cancelled(), SemanticStatus_Stale()]
    @classmethod
    def default(cls, ):
        return lambda: SemanticStatus_Complete()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Complete(self) -> bool:
        return isinstance(self, SemanticStatus_Complete)
    @property
    def is_Failed(self) -> bool:
        return isinstance(self, SemanticStatus_Failed)
    @property
    def is_Unavailable(self) -> bool:
        return isinstance(self, SemanticStatus_Unavailable)
    @property
    def is_Cancelled(self) -> bool:
        return isinstance(self, SemanticStatus_Cancelled)
    @property
    def is_Stale(self) -> bool:
        return isinstance(self, SemanticStatus_Stale)

class SemanticStatus_Complete(SemanticStatus, NamedTuple('Complete', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SemanticStatus.Complete'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SemanticStatus_Complete)
    def __hash__(self) -> int:
        return super().__hash__()

class SemanticStatus_Failed(SemanticStatus, NamedTuple('Failed', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SemanticStatus.Failed'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SemanticStatus_Failed)
    def __hash__(self) -> int:
        return super().__hash__()

class SemanticStatus_Unavailable(SemanticStatus, NamedTuple('Unavailable', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SemanticStatus.Unavailable'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SemanticStatus_Unavailable)
    def __hash__(self) -> int:
        return super().__hash__()

class SemanticStatus_Cancelled(SemanticStatus, NamedTuple('Cancelled', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SemanticStatus.Cancelled'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SemanticStatus_Cancelled)
    def __hash__(self) -> int:
        return super().__hash__()

class SemanticStatus_Stale(SemanticStatus, NamedTuple('Stale', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SemanticStatus.Stale'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SemanticStatus_Stale)
    def __hash__(self) -> int:
        return super().__hash__()


class Availability:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [Availability_Available(), Availability_Partial(), Availability_CapabilityUnavailable()]
    @classmethod
    def default(cls, ):
        return lambda: Availability_Available()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Available(self) -> bool:
        return isinstance(self, Availability_Available)
    @property
    def is_Partial(self) -> bool:
        return isinstance(self, Availability_Partial)
    @property
    def is_CapabilityUnavailable(self) -> bool:
        return isinstance(self, Availability_CapabilityUnavailable)

class Availability_Available(Availability, NamedTuple('Available', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.Availability.Available'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Availability_Available)
    def __hash__(self) -> int:
        return super().__hash__()

class Availability_Partial(Availability, NamedTuple('Partial', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.Availability.Partial'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Availability_Partial)
    def __hash__(self) -> int:
        return super().__hash__()

class Availability_CapabilityUnavailable(Availability, NamedTuple('CapabilityUnavailable', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.Availability.CapabilityUnavailable'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Availability_CapabilityUnavailable)
    def __hash__(self) -> int:
        return super().__hash__()


class CapabilityKind:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [CapabilityKind_Transport(), CapabilityKind_Snapshot(), CapabilityKind_View(), CapabilityKind_Analysis(), CapabilityKind_Projection(), CapabilityKind_Cache(), CapabilityKind_Lens()]
    @classmethod
    def default(cls, ):
        return lambda: CapabilityKind_Transport()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Transport(self) -> bool:
        return isinstance(self, CapabilityKind_Transport)
    @property
    def is_Snapshot(self) -> bool:
        return isinstance(self, CapabilityKind_Snapshot)
    @property
    def is_View(self) -> bool:
        return isinstance(self, CapabilityKind_View)
    @property
    def is_Analysis(self) -> bool:
        return isinstance(self, CapabilityKind_Analysis)
    @property
    def is_Projection(self) -> bool:
        return isinstance(self, CapabilityKind_Projection)
    @property
    def is_Cache(self) -> bool:
        return isinstance(self, CapabilityKind_Cache)
    @property
    def is_Lens(self) -> bool:
        return isinstance(self, CapabilityKind_Lens)

class CapabilityKind_Transport(CapabilityKind, NamedTuple('Transport', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.CapabilityKind.Transport'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CapabilityKind_Transport)
    def __hash__(self) -> int:
        return super().__hash__()

class CapabilityKind_Snapshot(CapabilityKind, NamedTuple('Snapshot', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.CapabilityKind.Snapshot'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CapabilityKind_Snapshot)
    def __hash__(self) -> int:
        return super().__hash__()

class CapabilityKind_View(CapabilityKind, NamedTuple('View', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.CapabilityKind.View'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CapabilityKind_View)
    def __hash__(self) -> int:
        return super().__hash__()

class CapabilityKind_Analysis(CapabilityKind, NamedTuple('Analysis', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.CapabilityKind.Analysis'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CapabilityKind_Analysis)
    def __hash__(self) -> int:
        return super().__hash__()

class CapabilityKind_Projection(CapabilityKind, NamedTuple('Projection', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.CapabilityKind.Projection'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CapabilityKind_Projection)
    def __hash__(self) -> int:
        return super().__hash__()

class CapabilityKind_Cache(CapabilityKind, NamedTuple('Cache', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.CapabilityKind.Cache'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CapabilityKind_Cache)
    def __hash__(self) -> int:
        return super().__hash__()

class CapabilityKind_Lens(CapabilityKind, NamedTuple('Lens', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.CapabilityKind.Lens'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CapabilityKind_Lens)
    def __hash__(self) -> int:
        return super().__hash__()


class FactLayer:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [FactLayer_Declared(), FactLayer_Discovered(), FactLayer_Observed(), FactLayer_Derived(), FactLayer_Policy()]
    @classmethod
    def default(cls, ):
        return lambda: FactLayer_Declared()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Declared(self) -> bool:
        return isinstance(self, FactLayer_Declared)
    @property
    def is_Discovered(self) -> bool:
        return isinstance(self, FactLayer_Discovered)
    @property
    def is_Observed(self) -> bool:
        return isinstance(self, FactLayer_Observed)
    @property
    def is_Derived(self) -> bool:
        return isinstance(self, FactLayer_Derived)
    @property
    def is_Policy(self) -> bool:
        return isinstance(self, FactLayer_Policy)

class FactLayer_Declared(FactLayer, NamedTuple('Declared', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FactLayer.Declared'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FactLayer_Declared)
    def __hash__(self) -> int:
        return super().__hash__()

class FactLayer_Discovered(FactLayer, NamedTuple('Discovered', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FactLayer.Discovered'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FactLayer_Discovered)
    def __hash__(self) -> int:
        return super().__hash__()

class FactLayer_Observed(FactLayer, NamedTuple('Observed', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FactLayer.Observed'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FactLayer_Observed)
    def __hash__(self) -> int:
        return super().__hash__()

class FactLayer_Derived(FactLayer, NamedTuple('Derived', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FactLayer.Derived'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FactLayer_Derived)
    def __hash__(self) -> int:
        return super().__hash__()

class FactLayer_Policy(FactLayer, NamedTuple('Policy', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FactLayer.Policy'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FactLayer_Policy)
    def __hash__(self) -> int:
        return super().__hash__()


class FindingSeverity:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [FindingSeverity_Error(), FindingSeverity_Warning(), FindingSeverity_Information(), FindingSeverity_Hint()]
    @classmethod
    def default(cls, ):
        return lambda: FindingSeverity_Error()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Error(self) -> bool:
        return isinstance(self, FindingSeverity_Error)
    @property
    def is_Warning(self) -> bool:
        return isinstance(self, FindingSeverity_Warning)
    @property
    def is_Information(self) -> bool:
        return isinstance(self, FindingSeverity_Information)
    @property
    def is_Hint(self) -> bool:
        return isinstance(self, FindingSeverity_Hint)

class FindingSeverity_Error(FindingSeverity, NamedTuple('Error', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FindingSeverity.Error'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FindingSeverity_Error)
    def __hash__(self) -> int:
        return super().__hash__()

class FindingSeverity_Warning(FindingSeverity, NamedTuple('Warning', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FindingSeverity.Warning'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FindingSeverity_Warning)
    def __hash__(self) -> int:
        return super().__hash__()

class FindingSeverity_Information(FindingSeverity, NamedTuple('Information', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FindingSeverity.Information'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FindingSeverity_Information)
    def __hash__(self) -> int:
        return super().__hash__()

class FindingSeverity_Hint(FindingSeverity, NamedTuple('Hint', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FindingSeverity.Hint'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FindingSeverity_Hint)
    def __hash__(self) -> int:
        return super().__hash__()


class FindingStatus:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [FindingStatus_Open(), FindingStatus_Resolved(), FindingStatus_FindingUnavailable()]
    @classmethod
    def default(cls, ):
        return lambda: FindingStatus_Open()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Open(self) -> bool:
        return isinstance(self, FindingStatus_Open)
    @property
    def is_Resolved(self) -> bool:
        return isinstance(self, FindingStatus_Resolved)
    @property
    def is_FindingUnavailable(self) -> bool:
        return isinstance(self, FindingStatus_FindingUnavailable)

class FindingStatus_Open(FindingStatus, NamedTuple('Open', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FindingStatus.Open'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FindingStatus_Open)
    def __hash__(self) -> int:
        return super().__hash__()

class FindingStatus_Resolved(FindingStatus, NamedTuple('Resolved', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FindingStatus.Resolved'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FindingStatus_Resolved)
    def __hash__(self) -> int:
        return super().__hash__()

class FindingStatus_FindingUnavailable(FindingStatus, NamedTuple('FindingUnavailable', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.FindingStatus.FindingUnavailable'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, FindingStatus_FindingUnavailable)
    def __hash__(self) -> int:
        return super().__hash__()


class EvidenceKind:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [EvidenceKind_Source(), EvidenceKind_Overlay(), EvidenceKind_Directory(), EvidenceKind_Config(), EvidenceKind_Std(), EvidenceKind_Doo(), EvidenceKind_Syntax(), EvidenceKind_Symbol(), EvidenceKind_CallEdge(), EvidenceKind_LexicalMatch(), EvidenceKind_Plugin(), EvidenceKind_Rule(), EvidenceKind_Derivation(), EvidenceKind_PolicyEvidence(), EvidenceKind_Receipt(), EvidenceKind_CacheDecision()]
    @classmethod
    def default(cls, ):
        return lambda: EvidenceKind_Source()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Source(self) -> bool:
        return isinstance(self, EvidenceKind_Source)
    @property
    def is_Overlay(self) -> bool:
        return isinstance(self, EvidenceKind_Overlay)
    @property
    def is_Directory(self) -> bool:
        return isinstance(self, EvidenceKind_Directory)
    @property
    def is_Config(self) -> bool:
        return isinstance(self, EvidenceKind_Config)
    @property
    def is_Std(self) -> bool:
        return isinstance(self, EvidenceKind_Std)
    @property
    def is_Doo(self) -> bool:
        return isinstance(self, EvidenceKind_Doo)
    @property
    def is_Syntax(self) -> bool:
        return isinstance(self, EvidenceKind_Syntax)
    @property
    def is_Symbol(self) -> bool:
        return isinstance(self, EvidenceKind_Symbol)
    @property
    def is_CallEdge(self) -> bool:
        return isinstance(self, EvidenceKind_CallEdge)
    @property
    def is_LexicalMatch(self) -> bool:
        return isinstance(self, EvidenceKind_LexicalMatch)
    @property
    def is_Plugin(self) -> bool:
        return isinstance(self, EvidenceKind_Plugin)
    @property
    def is_Rule(self) -> bool:
        return isinstance(self, EvidenceKind_Rule)
    @property
    def is_Derivation(self) -> bool:
        return isinstance(self, EvidenceKind_Derivation)
    @property
    def is_PolicyEvidence(self) -> bool:
        return isinstance(self, EvidenceKind_PolicyEvidence)
    @property
    def is_Receipt(self) -> bool:
        return isinstance(self, EvidenceKind_Receipt)
    @property
    def is_CacheDecision(self) -> bool:
        return isinstance(self, EvidenceKind_CacheDecision)

class EvidenceKind_Source(EvidenceKind, NamedTuple('Source', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Source'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Source)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_Overlay(EvidenceKind, NamedTuple('Overlay', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Overlay'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Overlay)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_Directory(EvidenceKind, NamedTuple('Directory', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Directory'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Directory)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_Config(EvidenceKind, NamedTuple('Config', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Config'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Config)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_Std(EvidenceKind, NamedTuple('Std', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Std'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Std)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_Doo(EvidenceKind, NamedTuple('Doo', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Doo'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Doo)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_Syntax(EvidenceKind, NamedTuple('Syntax', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Syntax'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Syntax)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_Symbol(EvidenceKind, NamedTuple('Symbol', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Symbol'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Symbol)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_CallEdge(EvidenceKind, NamedTuple('CallEdge', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.CallEdge'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_CallEdge)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_LexicalMatch(EvidenceKind, NamedTuple('LexicalMatch', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.LexicalMatch'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_LexicalMatch)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_Plugin(EvidenceKind, NamedTuple('Plugin', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Plugin'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Plugin)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_Rule(EvidenceKind, NamedTuple('Rule', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Rule'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Rule)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_Derivation(EvidenceKind, NamedTuple('Derivation', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Derivation'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Derivation)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_PolicyEvidence(EvidenceKind, NamedTuple('PolicyEvidence', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.PolicyEvidence'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_PolicyEvidence)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_Receipt(EvidenceKind, NamedTuple('Receipt', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.Receipt'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_Receipt)
    def __hash__(self) -> int:
        return super().__hash__()

class EvidenceKind_CacheDecision(EvidenceKind, NamedTuple('CacheDecision', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceKind.CacheDecision'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceKind_CacheDecision)
    def __hash__(self) -> int:
        return super().__hash__()


class LensArchetype:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [LensArchetype_Graph(), LensArchetype_Matrix(), LensArchetype_Layered(), LensArchetype_List(), LensArchetype_Table(), LensArchetype_MasterDetail()]
    @classmethod
    def default(cls, ):
        return lambda: LensArchetype_Graph()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Graph(self) -> bool:
        return isinstance(self, LensArchetype_Graph)
    @property
    def is_Matrix(self) -> bool:
        return isinstance(self, LensArchetype_Matrix)
    @property
    def is_Layered(self) -> bool:
        return isinstance(self, LensArchetype_Layered)
    @property
    def is_List(self) -> bool:
        return isinstance(self, LensArchetype_List)
    @property
    def is_Table(self) -> bool:
        return isinstance(self, LensArchetype_Table)
    @property
    def is_MasterDetail(self) -> bool:
        return isinstance(self, LensArchetype_MasterDetail)

class LensArchetype_Graph(LensArchetype, NamedTuple('Graph', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.LensArchetype.Graph'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, LensArchetype_Graph)
    def __hash__(self) -> int:
        return super().__hash__()

class LensArchetype_Matrix(LensArchetype, NamedTuple('Matrix', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.LensArchetype.Matrix'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, LensArchetype_Matrix)
    def __hash__(self) -> int:
        return super().__hash__()

class LensArchetype_Layered(LensArchetype, NamedTuple('Layered', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.LensArchetype.Layered'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, LensArchetype_Layered)
    def __hash__(self) -> int:
        return super().__hash__()

class LensArchetype_List(LensArchetype, NamedTuple('List', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.LensArchetype.List'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, LensArchetype_List)
    def __hash__(self) -> int:
        return super().__hash__()

class LensArchetype_Table(LensArchetype, NamedTuple('Table', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.LensArchetype.Table'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, LensArchetype_Table)
    def __hash__(self) -> int:
        return super().__hash__()

class LensArchetype_MasterDetail(LensArchetype, NamedTuple('MasterDetail', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.LensArchetype.MasterDetail'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, LensArchetype_MasterDetail)
    def __hash__(self) -> int:
        return super().__hash__()


class TypedValue:
    @classmethod
    def default(cls, ):
        return lambda: TypedValue_TVNull()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_TVNull(self) -> bool:
        return isinstance(self, TypedValue_TVNull)
    @property
    def is_TVBool(self) -> bool:
        return isinstance(self, TypedValue_TVBool)
    @property
    def is_TVInteger(self) -> bool:
        return isinstance(self, TypedValue_TVInteger)
    @property
    def is_TVDecimal(self) -> bool:
        return isinstance(self, TypedValue_TVDecimal)
    @property
    def is_TVString(self) -> bool:
        return isinstance(self, TypedValue_TVString)
    @property
    def is_TVRef(self) -> bool:
        return isinstance(self, TypedValue_TVRef)
    @property
    def is_TVList(self) -> bool:
        return isinstance(self, TypedValue_TVList)
    @property
    def is_TVRecord(self) -> bool:
        return isinstance(self, TypedValue_TVRecord)

class TypedValue_TVNull(TypedValue, NamedTuple('TVNull', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.TypedValue.TVNull'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TypedValue_TVNull)
    def __hash__(self) -> int:
        return super().__hash__()

class TypedValue_TVBool(TypedValue, NamedTuple('TVBool', [('boolValue', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.TypedValue.TVBool({_dafny.string_of(self.boolValue)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TypedValue_TVBool) and self.boolValue == __o.boolValue
    def __hash__(self) -> int:
        return super().__hash__()

class TypedValue_TVInteger(TypedValue, NamedTuple('TVInteger', [('decimalString', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.TypedValue.TVInteger({self.decimalString.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TypedValue_TVInteger) and self.decimalString == __o.decimalString
    def __hash__(self) -> int:
        return super().__hash__()

class TypedValue_TVDecimal(TypedValue, NamedTuple('TVDecimal', [('canonicalDecimalString', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.TypedValue.TVDecimal({self.canonicalDecimalString.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TypedValue_TVDecimal) and self.canonicalDecimalString == __o.canonicalDecimalString
    def __hash__(self) -> int:
        return super().__hash__()

class TypedValue_TVString(TypedValue, NamedTuple('TVString', [('stringValue', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.TypedValue.TVString({self.stringValue.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TypedValue_TVString) and self.stringValue == __o.stringValue
    def __hash__(self) -> int:
        return super().__hash__()

class TypedValue_TVRef(TypedValue, NamedTuple('TVRef', [('targetId', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.TypedValue.TVRef({self.targetId.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TypedValue_TVRef) and self.targetId == __o.targetId
    def __hash__(self) -> int:
        return super().__hash__()

class TypedValue_TVList(TypedValue, NamedTuple('TVList', [('items', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.TypedValue.TVList({_dafny.string_of(self.items)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TypedValue_TVList) and self.items == __o.items
    def __hash__(self) -> int:
        return super().__hash__()

class TypedValue_TVRecord(TypedValue, NamedTuple('TVRecord', [('fieldNames', Any), ('fieldValues', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.TypedValue.TVRecord({_dafny.string_of(self.fieldNames)}, {_dafny.string_of(self.fieldValues)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TypedValue_TVRecord) and self.fieldNames == __o.fieldNames and self.fieldValues == __o.fieldValues
    def __hash__(self) -> int:
        return super().__hash__()


class TypedField:
    @classmethod
    def default(cls, ):
        return lambda: TypedField_TypedField(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), TypedValue.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_TypedField(self) -> bool:
        return isinstance(self, TypedField_TypedField)

class TypedField_TypedField(TypedField, NamedTuple('TypedField', [('name', Any), ('value', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.TypedField.TypedField({self.name.VerbatimString(True)}, {_dafny.string_of(self.value)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, TypedField_TypedField) and self.name == __o.name and self.value == __o.value
    def __hash__(self) -> int:
        return super().__hash__()


class Completeness:
    @classmethod
    def default(cls, ):
        return lambda: Completeness_Completeness(False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Completeness(self) -> bool:
        return isinstance(self, Completeness_Completeness)

class Completeness_Completeness(Completeness, NamedTuple('Completeness', [('complete', Any), ('reasons', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.Completeness.Completeness({_dafny.string_of(self.complete)}, {_dafny.string_of(self.reasons)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, Completeness_Completeness) and self.complete == __o.complete and self.reasons == __o.reasons
    def __hash__(self) -> int:
        return super().__hash__()


class SourcePosition:
    @classmethod
    def default(cls, ):
        return lambda: SourcePosition_SourcePosition(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SourcePosition(self) -> bool:
        return isinstance(self, SourcePosition_SourcePosition)

class SourcePosition_SourcePosition(SourcePosition, NamedTuple('SourcePosition', [('line', Any), ('column', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SourcePosition.SourcePosition({_dafny.string_of(self.line)}, {_dafny.string_of(self.column)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SourcePosition_SourcePosition) and self.line == __o.line and self.column == __o.column
    def __hash__(self) -> int:
        return super().__hash__()


class SourceRange:
    @classmethod
    def default(cls, ):
        return lambda: SourceRange_SourceRange(SourcePosition.default()(), SourcePosition.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SourceRange(self) -> bool:
        return isinstance(self, SourceRange_SourceRange)

class SourceRange_SourceRange(SourceRange, NamedTuple('SourceRange', [('start', Any), ('end', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SourceRange.SourceRange({_dafny.string_of(self.start)}, {_dafny.string_of(self.end)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SourceRange_SourceRange) and self.start == __o.start and self.end == __o.end
    def __hash__(self) -> int:
        return super().__hash__()


class CapabilityRequirement:
    @classmethod
    def default(cls, ):
        return lambda: CapabilityRequirement_CapabilityRequirement(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_CapabilityRequirement(self) -> bool:
        return isinstance(self, CapabilityRequirement_CapabilityRequirement)

class CapabilityRequirement_CapabilityRequirement(CapabilityRequirement, NamedTuple('CapabilityRequirement', [('id_', Any), ('versionRange', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.CapabilityRequirement.CapabilityRequirement({self.id_.VerbatimString(True)}, {self.versionRange.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CapabilityRequirement_CapabilityRequirement) and self.id_ == __o.id_ and self.versionRange == __o.versionRange
    def __hash__(self) -> int:
        return super().__hash__()


class CapabilityDescriptor:
    @classmethod
    def default(cls, ):
        return lambda: CapabilityDescriptor_CapabilityDescriptor(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), CapabilityKind.default()(), _dafny.Set({}), _dafny.Seq({}), Availability.default()(), _dafny.Seq({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_CapabilityDescriptor(self) -> bool:
        return isinstance(self, CapabilityDescriptor_CapabilityDescriptor)

class CapabilityDescriptor_CapabilityDescriptor(CapabilityDescriptor, NamedTuple('CapabilityDescriptor', [('id_', Any), ('version', Any), ('schemaDigest', Any), ('implementationDigest', Any), ('kind', Any), ('scopes', Any), ('requirements', Any), ('availability', Any), ('unavailableReasons', Any), ('contentRevision', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.CapabilityDescriptor.CapabilityDescriptor({self.id_.VerbatimString(True)}, {self.version.VerbatimString(True)}, {self.schemaDigest.VerbatimString(True)}, {self.implementationDigest.VerbatimString(True)}, {_dafny.string_of(self.kind)}, {_dafny.string_of(self.scopes)}, {_dafny.string_of(self.requirements)}, {_dafny.string_of(self.availability)}, {_dafny.string_of(self.unavailableReasons)}, {self.contentRevision.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, CapabilityDescriptor_CapabilityDescriptor) and self.id_ == __o.id_ and self.version == __o.version and self.schemaDigest == __o.schemaDigest and self.implementationDigest == __o.implementationDigest and self.kind == __o.kind and self.scopes == __o.scopes and self.requirements == __o.requirements and self.availability == __o.availability and self.unavailableReasons == __o.unavailableReasons and self.contentRevision == __o.contentRevision
    def __hash__(self) -> int:
        return super().__hash__()


class AuthorityDescriptor:
    @classmethod
    def default(cls, ):
        return lambda: AuthorityDescriptor_AuthorityDescriptor(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_AuthorityDescriptor(self) -> bool:
        return isinstance(self, AuthorityDescriptor_AuthorityDescriptor)

class AuthorityDescriptor_AuthorityDescriptor(AuthorityDescriptor, NamedTuple('AuthorityDescriptor', [('authorityId', Any), ('producerId', Any), ('producerVersion', Any), ('implementationDigest', Any), ('sourceArtifactDigest', Any), ('registryDigest', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.AuthorityDescriptor.AuthorityDescriptor({self.authorityId.VerbatimString(True)}, {self.producerId.VerbatimString(True)}, {self.producerVersion.VerbatimString(True)}, {self.implementationDigest.VerbatimString(True)}, {self.sourceArtifactDigest.VerbatimString(True)}, {self.registryDigest.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, AuthorityDescriptor_AuthorityDescriptor) and self.authorityId == __o.authorityId and self.producerId == __o.producerId and self.producerVersion == __o.producerVersion and self.implementationDigest == __o.implementationDigest and self.sourceArtifactDigest == __o.sourceArtifactDigest and self.registryDigest == __o.registryDigest
    def __hash__(self) -> int:
        return super().__hash__()


class EvidenceRecord:
    @classmethod
    def default(cls, ):
        return lambda: EvidenceRecord_EvidenceRecord(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), EvidenceKind.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxContract.Option.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Set({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_EvidenceRecord(self) -> bool:
        return isinstance(self, EvidenceRecord_EvidenceRecord)

class EvidenceRecord_EvidenceRecord(EvidenceRecord, NamedTuple('EvidenceRecord', [('id_', Any), ('snapshotId', Any), ('kind', Any), ('logicalRootId', Any), ('logicalUri', Any), ('locator', Any), ('range_', Any), ('symbolId', Any), ('artifactDigest', Any), ('producerDigest', Any), ('statement', Any), ('parentEvidenceIds', Any), ('contentRevision', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.EvidenceRecord.EvidenceRecord({self.id_.VerbatimString(True)}, {self.snapshotId.VerbatimString(True)}, {_dafny.string_of(self.kind)}, {self.logicalRootId.VerbatimString(True)}, {self.logicalUri.VerbatimString(True)}, {self.locator.VerbatimString(True)}, {_dafny.string_of(self.range_)}, {self.symbolId.VerbatimString(True)}, {self.artifactDigest.VerbatimString(True)}, {self.producerDigest.VerbatimString(True)}, {self.statement.VerbatimString(True)}, {_dafny.string_of(self.parentEvidenceIds)}, {self.contentRevision.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, EvidenceRecord_EvidenceRecord) and self.id_ == __o.id_ and self.snapshotId == __o.snapshotId and self.kind == __o.kind and self.logicalRootId == __o.logicalRootId and self.logicalUri == __o.logicalUri and self.locator == __o.locator and self.range_ == __o.range_ and self.symbolId == __o.symbolId and self.artifactDigest == __o.artifactDigest and self.producerDigest == __o.producerDigest and self.statement == __o.statement and self.parentEvidenceIds == __o.parentEvidenceIds and self.contentRevision == __o.contentRevision
    def __hash__(self) -> int:
        return super().__hash__()


class SemanticNode:
    @classmethod
    def default(cls, ):
        return lambda: SemanticNode_SemanticNode(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}), _dafny.Set({}), _dafny.Set({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SemanticNode(self) -> bool:
        return isinstance(self, SemanticNode_SemanticNode)

class SemanticNode_SemanticNode(SemanticNode, NamedTuple('SemanticNode', [('id_', Any), ('kind', Any), ('displayLabel', Any), ('attributes', Any), ('factIds', Any), ('evidenceIds', Any), ('contentRevision', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SemanticNode.SemanticNode({self.id_.VerbatimString(True)}, {self.kind.VerbatimString(True)}, {self.displayLabel.VerbatimString(True)}, {_dafny.string_of(self.attributes)}, {_dafny.string_of(self.factIds)}, {_dafny.string_of(self.evidenceIds)}, {self.contentRevision.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SemanticNode_SemanticNode) and self.id_ == __o.id_ and self.kind == __o.kind and self.displayLabel == __o.displayLabel and self.attributes == __o.attributes and self.factIds == __o.factIds and self.evidenceIds == __o.evidenceIds and self.contentRevision == __o.contentRevision
    def __hash__(self) -> int:
        return super().__hash__()


class SemanticEdge:
    @classmethod
    def default(cls, ):
        return lambda: SemanticEdge_SemanticEdge(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, _dafny.Seq({}), _dafny.Set({}), _dafny.Set({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SemanticEdge(self) -> bool:
        return isinstance(self, SemanticEdge_SemanticEdge)

class SemanticEdge_SemanticEdge(SemanticEdge, NamedTuple('SemanticEdge', [('id_', Any), ('kind', Any), ('fromId', Any), ('toId', Any), ('directed', Any), ('attributes', Any), ('factIds', Any), ('evidenceIds', Any), ('contentRevision', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SemanticEdge.SemanticEdge({self.id_.VerbatimString(True)}, {self.kind.VerbatimString(True)}, {self.fromId.VerbatimString(True)}, {self.toId.VerbatimString(True)}, {_dafny.string_of(self.directed)}, {_dafny.string_of(self.attributes)}, {_dafny.string_of(self.factIds)}, {_dafny.string_of(self.evidenceIds)}, {self.contentRevision.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SemanticEdge_SemanticEdge) and self.id_ == __o.id_ and self.kind == __o.kind and self.fromId == __o.fromId and self.toId == __o.toId and self.directed == __o.directed and self.attributes == __o.attributes and self.factIds == __o.factIds and self.evidenceIds == __o.evidenceIds and self.contentRevision == __o.contentRevision
    def __hash__(self) -> int:
        return super().__hash__()


class SemanticFact:
    @classmethod
    def default(cls, ):
        return lambda: SemanticFact_SemanticFact(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), TypedValue.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), FactLayer.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Set({}), _dafny.Set({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SemanticFact(self) -> bool:
        return isinstance(self, SemanticFact_SemanticFact)

class SemanticFact_SemanticFact(SemanticFact, NamedTuple('SemanticFact', [('id_', Any), ('analysisId', Any), ('subjectId', Any), ('key', Any), ('value', Any), ('valueSchema', Any), ('layer', Any), ('ruleId', Any), ('supportingFactIds', Any), ('evidenceIds', Any), ('contentRevision', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SemanticFact.SemanticFact({self.id_.VerbatimString(True)}, {self.analysisId.VerbatimString(True)}, {self.subjectId.VerbatimString(True)}, {self.key.VerbatimString(True)}, {_dafny.string_of(self.value)}, {self.valueSchema.VerbatimString(True)}, {_dafny.string_of(self.layer)}, {self.ruleId.VerbatimString(True)}, {_dafny.string_of(self.supportingFactIds)}, {_dafny.string_of(self.evidenceIds)}, {self.contentRevision.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SemanticFact_SemanticFact) and self.id_ == __o.id_ and self.analysisId == __o.analysisId and self.subjectId == __o.subjectId and self.key == __o.key and self.value == __o.value and self.valueSchema == __o.valueSchema and self.layer == __o.layer and self.ruleId == __o.ruleId and self.supportingFactIds == __o.supportingFactIds and self.evidenceIds == __o.evidenceIds and self.contentRevision == __o.contentRevision
    def __hash__(self) -> int:
        return super().__hash__()


class SemanticFinding:
    @classmethod
    def default(cls, ):
        return lambda: SemanticFinding_SemanticFinding(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), FindingSeverity.default()(), FindingStatus.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), TypedValue.default()(), TypedValue.default()(), _dafny.Set({}), _dafny.Set({}), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SemanticFinding(self) -> bool:
        return isinstance(self, SemanticFinding_SemanticFinding)

class SemanticFinding_SemanticFinding(SemanticFinding, NamedTuple('SemanticFinding', [('id_', Any), ('analysisId', Any), ('ruleId', Any), ('subjectId', Any), ('severity', Any), ('status', Any), ('code', Any), ('message', Any), ('expected', Any), ('actual', Any), ('supportingFactIds', Any), ('evidenceIds', Any), ('blocking', Any), ('contentRevision', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SemanticFinding.SemanticFinding({self.id_.VerbatimString(True)}, {self.analysisId.VerbatimString(True)}, {self.ruleId.VerbatimString(True)}, {self.subjectId.VerbatimString(True)}, {_dafny.string_of(self.severity)}, {_dafny.string_of(self.status)}, {self.code.VerbatimString(True)}, {self.message.VerbatimString(True)}, {_dafny.string_of(self.expected)}, {_dafny.string_of(self.actual)}, {_dafny.string_of(self.supportingFactIds)}, {_dafny.string_of(self.evidenceIds)}, {_dafny.string_of(self.blocking)}, {self.contentRevision.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SemanticFinding_SemanticFinding) and self.id_ == __o.id_ and self.analysisId == __o.analysisId and self.ruleId == __o.ruleId and self.subjectId == __o.subjectId and self.severity == __o.severity and self.status == __o.status and self.code == __o.code and self.message == __o.message and self.expected == __o.expected and self.actual == __o.actual and self.supportingFactIds == __o.supportingFactIds and self.evidenceIds == __o.evidenceIds and self.blocking == __o.blocking and self.contentRevision == __o.contentRevision
    def __hash__(self) -> int:
        return super().__hash__()


class GraphLensDescriptor:
    @classmethod
    def default(cls, ):
        return lambda: GraphLensDescriptor_GraphLensDescriptor(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), LensArchetype.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}), _dafny.Seq({}), _dafny.Seq({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_GraphLensDescriptor(self) -> bool:
        return isinstance(self, GraphLensDescriptor_GraphLensDescriptor)

class GraphLensDescriptor_GraphLensDescriptor(GraphLensDescriptor, NamedTuple('GraphLensDescriptor', [('id_', Any), ('version', Any), ('archetype', Any), ('title', Any), ('nodeSelector', Any), ('edgeSelector', Any), ('factSelector', Any), ('labelField', Any), ('detailFields', Any), ('columns', Any), ('requiredCapabilities', Any), ('contentRevision', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.GraphLensDescriptor.GraphLensDescriptor({self.id_.VerbatimString(True)}, {self.version.VerbatimString(True)}, {_dafny.string_of(self.archetype)}, {self.title.VerbatimString(True)}, {self.nodeSelector.VerbatimString(True)}, {self.edgeSelector.VerbatimString(True)}, {self.factSelector.VerbatimString(True)}, {self.labelField.VerbatimString(True)}, {_dafny.string_of(self.detailFields)}, {_dafny.string_of(self.columns)}, {_dafny.string_of(self.requiredCapabilities)}, {self.contentRevision.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, GraphLensDescriptor_GraphLensDescriptor) and self.id_ == __o.id_ and self.version == __o.version and self.archetype == __o.archetype and self.title == __o.title and self.nodeSelector == __o.nodeSelector and self.edgeSelector == __o.edgeSelector and self.factSelector == __o.factSelector and self.labelField == __o.labelField and self.detailFields == __o.detailFields and self.columns == __o.columns and self.requiredCapabilities == __o.requiredCapabilities and self.contentRevision == __o.contentRevision
    def __hash__(self) -> int:
        return super().__hash__()


class SemanticGraphSnapshotV1:
    @classmethod
    def default(cls, ):
        return lambda: SemanticGraphSnapshotV1_SemanticGraphSnapshotV1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), AuthorityDescriptor.default()(), _dafny.Map({}), _dafny.Map({}), _dafny.Map({}), _dafny.Map({}), _dafny.Map({}), _dafny.Map({}), _dafny.Map({}), Completeness.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SemanticGraphSnapshotV1(self) -> bool:
        return isinstance(self, SemanticGraphSnapshotV1_SemanticGraphSnapshotV1)

class SemanticGraphSnapshotV1_SemanticGraphSnapshotV1(SemanticGraphSnapshotV1, NamedTuple('SemanticGraphSnapshotV1', [('schemaVersion', Any), ('graphId', Any), ('revision', Any), ('previousRevision', Any), ('snapshotId', Any), ('authority', Any), ('capabilities', Any), ('lenses', Any), ('nodes', Any), ('edges', Any), ('facts', Any), ('findings', Any), ('evidence', Any), ('completeness', Any), ('receiptId', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SemanticGraphSnapshotV1.SemanticGraphSnapshotV1({self.schemaVersion.VerbatimString(True)}, {self.graphId.VerbatimString(True)}, {self.revision.VerbatimString(True)}, {self.previousRevision.VerbatimString(True)}, {self.snapshotId.VerbatimString(True)}, {_dafny.string_of(self.authority)}, {_dafny.string_of(self.capabilities)}, {_dafny.string_of(self.lenses)}, {_dafny.string_of(self.nodes)}, {_dafny.string_of(self.edges)}, {_dafny.string_of(self.facts)}, {_dafny.string_of(self.findings)}, {_dafny.string_of(self.evidence)}, {_dafny.string_of(self.completeness)}, {self.receiptId.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SemanticGraphSnapshotV1_SemanticGraphSnapshotV1) and self.schemaVersion == __o.schemaVersion and self.graphId == __o.graphId and self.revision == __o.revision and self.previousRevision == __o.previousRevision and self.snapshotId == __o.snapshotId and self.authority == __o.authority and self.capabilities == __o.capabilities and self.lenses == __o.lenses and self.nodes == __o.nodes and self.edges == __o.edges and self.facts == __o.facts and self.findings == __o.findings and self.evidence == __o.evidence and self.completeness == __o.completeness and self.receiptId == __o.receiptId
    def __hash__(self) -> int:
        return super().__hash__()


class ResetReason:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [ResetReason_UnknownRevision(), ResetReason_Expired(), ResetReason_GraphMismatch(), ResetReason_SchemaEpoch(), ResetReason_CorruptBase()]
    @classmethod
    def default(cls, ):
        return lambda: ResetReason_UnknownRevision()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_UnknownRevision(self) -> bool:
        return isinstance(self, ResetReason_UnknownRevision)
    @property
    def is_Expired(self) -> bool:
        return isinstance(self, ResetReason_Expired)
    @property
    def is_GraphMismatch(self) -> bool:
        return isinstance(self, ResetReason_GraphMismatch)
    @property
    def is_SchemaEpoch(self) -> bool:
        return isinstance(self, ResetReason_SchemaEpoch)
    @property
    def is_CorruptBase(self) -> bool:
        return isinstance(self, ResetReason_CorruptBase)

class ResetReason_UnknownRevision(ResetReason, NamedTuple('UnknownRevision', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.ResetReason.UnknownRevision'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResetReason_UnknownRevision)
    def __hash__(self) -> int:
        return super().__hash__()

class ResetReason_Expired(ResetReason, NamedTuple('Expired', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.ResetReason.Expired'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResetReason_Expired)
    def __hash__(self) -> int:
        return super().__hash__()

class ResetReason_GraphMismatch(ResetReason, NamedTuple('GraphMismatch', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.ResetReason.GraphMismatch'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResetReason_GraphMismatch)
    def __hash__(self) -> int:
        return super().__hash__()

class ResetReason_SchemaEpoch(ResetReason, NamedTuple('SchemaEpoch', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.ResetReason.SchemaEpoch'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResetReason_SchemaEpoch)
    def __hash__(self) -> int:
        return super().__hash__()

class ResetReason_CorruptBase(ResetReason, NamedTuple('CorruptBase', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.ResetReason.CorruptBase'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ResetReason_CorruptBase)
    def __hash__(self) -> int:
        return super().__hash__()


class GraphChanges:
    @classmethod
    def default(cls, ):
        return lambda: GraphChanges_GraphChanges(_dafny.Map({}), _dafny.Map({}), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_GraphChanges(self) -> bool:
        return isinstance(self, GraphChanges_GraphChanges)

class GraphChanges_GraphChanges(GraphChanges, NamedTuple('GraphChanges', [('added', Any), ('changed', Any), ('removed', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.GraphChanges.GraphChanges({_dafny.string_of(self.added)}, {_dafny.string_of(self.changed)}, {_dafny.string_of(self.removed)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, GraphChanges_GraphChanges) and self.added == __o.added and self.changed == __o.changed and self.removed == __o.removed
    def __hash__(self) -> int:
        return super().__hash__()


class SemanticGraphDeltaV1:
    @classmethod
    def default(cls, ):
        return lambda: SemanticGraphDeltaV1_SemanticGraphDeltaV1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, ConfluxContract.Option.default()(), ConfluxContract.Option.default()(), GraphChanges.default()(), GraphChanges.default()(), GraphChanges.default()(), GraphChanges.default()(), GraphChanges.default()(), ConfluxContract.Option.default()(), ConfluxContract.Option.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_SemanticGraphDeltaV1(self) -> bool:
        return isinstance(self, SemanticGraphDeltaV1_SemanticGraphDeltaV1)

class SemanticGraphDeltaV1_SemanticGraphDeltaV1(SemanticGraphDeltaV1, NamedTuple('SemanticGraphDeltaV1', [('schemaVersion', Any), ('graphId', Any), ('fromRevision', Any), ('toRevision', Any), ('reset', Any), ('resetReason', Any), ('replacement', Any), ('nodes', Any), ('edges', Any), ('facts', Any), ('findings', Any), ('evidence', Any), ('capabilitiesReplacement', Any), ('lensesReplacement', Any), ('receiptId', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphContract.SemanticGraphDeltaV1.SemanticGraphDeltaV1({self.schemaVersion.VerbatimString(True)}, {self.graphId.VerbatimString(True)}, {self.fromRevision.VerbatimString(True)}, {self.toRevision.VerbatimString(True)}, {_dafny.string_of(self.reset)}, {_dafny.string_of(self.resetReason)}, {_dafny.string_of(self.replacement)}, {_dafny.string_of(self.nodes)}, {_dafny.string_of(self.edges)}, {_dafny.string_of(self.facts)}, {_dafny.string_of(self.findings)}, {_dafny.string_of(self.evidence)}, {_dafny.string_of(self.capabilitiesReplacement)}, {_dafny.string_of(self.lensesReplacement)}, {self.receiptId.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, SemanticGraphDeltaV1_SemanticGraphDeltaV1) and self.schemaVersion == __o.schemaVersion and self.graphId == __o.graphId and self.fromRevision == __o.fromRevision and self.toRevision == __o.toRevision and self.reset == __o.reset and self.resetReason == __o.resetReason and self.replacement == __o.replacement and self.nodes == __o.nodes and self.edges == __o.edges and self.facts == __o.facts and self.findings == __o.findings and self.evidence == __o.evidence and self.capabilitiesReplacement == __o.capabilitiesReplacement and self.lensesReplacement == __o.lensesReplacement and self.receiptId == __o.receiptId
    def __hash__(self) -> int:
        return super().__hash__()

