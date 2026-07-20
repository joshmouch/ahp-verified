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

# Module: ConfluxSemanticGraphIdentity

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def StableId(canonicalEncode, sha256, input_):
        return sha256((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "archai:stable-id:1.0.0\n"))) + (canonicalEncode(input_)))

    @staticmethod
    def ContentRevision(kind, schemaVersion, canonicalSemanticContent, sha256, content):
        return sha256((((((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "archai:"))) + (kind)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ":")))) + (schemaVersion)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n")))) + (canonicalSemanticContent(content)))

    @staticmethod
    def GraphRevision(canonicalEncode, sha256, input_):
        return sha256((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "archai:semantic-graph:1.0.0\n"))) + (canonicalEncode(input_)))

    @staticmethod
    def GraphId(sha256, authorityId, selectorKind, selectorId, schemaEpoch):
        return sha256((((((((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "archai:graph-id:1.0.0\n"))) + (authorityId)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n")))) + (selectorKind)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n")))) + (selectorId)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n")))) + (schemaEpoch))


class StableIdentityInput:
    @classmethod
    def default(cls, ):
        return lambda: StableIdentityInput_StableIdentityInput(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_StableIdentityInput(self) -> bool:
        return isinstance(self, StableIdentityInput_StableIdentityInput)

class StableIdentityInput_StableIdentityInput(StableIdentityInput, NamedTuple('StableIdentityInput', [('kind', Any), ('authorityId', Any), ('naturalKey', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphIdentity.StableIdentityInput.StableIdentityInput({self.kind.VerbatimString(True)}, {self.authorityId.VerbatimString(True)}, {self.naturalKey.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, StableIdentityInput_StableIdentityInput) and self.kind == __o.kind and self.authorityId == __o.authorityId and self.naturalKey == __o.naturalKey
    def __hash__(self) -> int:
        return super().__hash__()


class GraphIdentityInput:
    @classmethod
    def default(cls, ):
        return lambda: GraphIdentityInput_GraphIdentityInput(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}), _dafny.Seq({}), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_GraphIdentityInput(self) -> bool:
        return isinstance(self, GraphIdentityInput_GraphIdentityInput)

class GraphIdentityInput_GraphIdentityInput(GraphIdentityInput, NamedTuple('GraphIdentityInput', [('graphId', Any), ('authorityArtifactIdentity', Any), ('capabilityRevisions', Any), ('lensRevisions', Any), ('recordRevisions', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphIdentity.GraphIdentityInput.GraphIdentityInput({self.graphId.VerbatimString(True)}, {self.authorityArtifactIdentity.VerbatimString(True)}, {_dafny.string_of(self.capabilityRevisions)}, {_dafny.string_of(self.lensRevisions)}, {_dafny.string_of(self.recordRevisions)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, GraphIdentityInput_GraphIdentityInput) and self.graphId == __o.graphId and self.authorityArtifactIdentity == __o.authorityArtifactIdentity and self.capabilityRevisions == __o.capabilityRevisions and self.lensRevisions == __o.lensRevisions and self.recordRevisions == __o.recordRevisions
    def __hash__(self) -> int:
        return super().__hash__()


class NonSemanticMetadata:
    @classmethod
    def default(cls, ):
        return lambda: NonSemanticMetadata_NonSemanticMetadata(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_NonSemanticMetadata(self) -> bool:
        return isinstance(self, NonSemanticMetadata_NonSemanticMetadata)

class NonSemanticMetadata_NonSemanticMetadata(NonSemanticMetadata, NamedTuple('NonSemanticMetadata', [('absoluteLocator', Any), ('gitHead', Any), ('listenerPid', Any), ('wallClock', Any), ('receiptLocator', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxSemanticGraphIdentity.NonSemanticMetadata.NonSemanticMetadata({self.absoluteLocator.VerbatimString(True)}, {self.gitHead.VerbatimString(True)}, {self.listenerPid.VerbatimString(True)}, {self.wallClock.VerbatimString(True)}, {self.receiptLocator.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, NonSemanticMetadata_NonSemanticMetadata) and self.absoluteLocator == __o.absoluteLocator and self.gitHead == __o.gitHead and self.listenerPid == __o.listenerPid and self.wallClock == __o.wallClock and self.receiptLocator == __o.receiptLocator
    def __hash__(self) -> int:
        return super().__hash__()

