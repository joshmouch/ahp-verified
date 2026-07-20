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

# Module: ConfluxContentAddress

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def SameContent(left, right):
        return ((left).address) == ((right).address)

    @staticmethod
    def SizeMatches(address, bytes_):
        source0_ = (address).exactSize
        if True:
            if source0_.is_None:
                return True
        if True:
            d_0_n_ = source0_.value
            return (len(bytes_)) == (d_0_n_)

    @staticmethod
    def VerifyMaterialization(reference, bytes_, digestBytes):
        return (default__.SizeMatches((reference).address, bytes_)) and ((digestBytes(((reference).address).digestAlgorithm, bytes_)) == (((reference).address).digest))


class DigestAlgorithm:
    @_dafny.classproperty
    def AllSingletonConstructors(cls):
        return [DigestAlgorithm_Sha256()]
    @classmethod
    def default(cls, ):
        return lambda: DigestAlgorithm_Sha256()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_Sha256(self) -> bool:
        return isinstance(self, DigestAlgorithm_Sha256)

class DigestAlgorithm_Sha256(DigestAlgorithm, NamedTuple('Sha256', [])):
    def __dafnystr__(self) -> str:
        return f'ConfluxContentAddress.DigestAlgorithm.Sha256'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, DigestAlgorithm_Sha256)
    def __hash__(self) -> int:
        return super().__hash__()


class ContentAddress:
    @classmethod
    def default(cls, ):
        return lambda: ContentAddress_ContentAddress(DigestAlgorithm.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxContract.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ContentAddress(self) -> bool:
        return isinstance(self, ContentAddress_ContentAddress)

class ContentAddress_ContentAddress(ContentAddress, NamedTuple('ContentAddress', [('digestAlgorithm', Any), ('digest', Any), ('exactSize', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxContentAddress.ContentAddress.ContentAddress({_dafny.string_of(self.digestAlgorithm)}, {self.digest.VerbatimString(True)}, {_dafny.string_of(self.exactSize)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ContentAddress_ContentAddress) and self.digestAlgorithm == __o.digestAlgorithm and self.digest == __o.digest and self.exactSize == __o.exactSize
    def __hash__(self) -> int:
        return super().__hash__()


class ContentLocator:
    @classmethod
    def default(cls, ):
        return lambda: ContentLocator_ContentLocator(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), ConfluxContract.Option.default()(), ConfluxContract.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ContentLocator(self) -> bool:
        return isinstance(self, ContentLocator_ContentLocator)

class ContentLocator_ContentLocator(ContentLocator, NamedTuple('ContentLocator', [('uri', Any), ('contentType', Any), ('nonce', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxContentAddress.ContentLocator.ContentLocator({self.uri.VerbatimString(True)}, {_dafny.string_of(self.contentType)}, {_dafny.string_of(self.nonce)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ContentLocator_ContentLocator) and self.uri == __o.uri and self.contentType == __o.contentType and self.nonce == __o.nonce
    def __hash__(self) -> int:
        return super().__hash__()


class VerifiedContentRef:
    @classmethod
    def default(cls, ):
        return lambda: VerifiedContentRef_VerifiedContentRef(ContentAddress.default()(), ConfluxContract.Option.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_VerifiedContentRef(self) -> bool:
        return isinstance(self, VerifiedContentRef_VerifiedContentRef)

class VerifiedContentRef_VerifiedContentRef(VerifiedContentRef, NamedTuple('VerifiedContentRef', [('address', Any), ('locator', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxContentAddress.VerifiedContentRef.VerifiedContentRef({_dafny.string_of(self.address)}, {_dafny.string_of(self.locator)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, VerifiedContentRef_VerifiedContentRef) and self.address == __o.address and self.locator == __o.locator
    def __hash__(self) -> int:
        return super().__hash__()


class ContentCodecs:
    @classmethod
    def default(cls, default_W):
        return lambda: ContentCodecs_ContentCodecs((lambda x2: default_W()), (lambda x3: ConfluxContract.Option.default()()), (lambda x4: default_W()), (lambda x5: ConfluxContract.Option.default()()), (lambda x6: default_W()), (lambda x7: ConfluxContract.Option.default()()))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_ContentCodecs(self) -> bool:
        return isinstance(self, ContentCodecs_ContentCodecs)

class ContentCodecs_ContentCodecs(ContentCodecs, NamedTuple('ContentCodecs', [('encodeAddress', Any), ('decodeAddress', Any), ('encodeLocator', Any), ('decodeLocator', Any), ('encodeRef', Any), ('decodeRef', Any)])):
    def __dafnystr__(self) -> str:
        return f'ConfluxContentAddress.ContentCodecs.ContentCodecs({_dafny.string_of(self.encodeAddress)}, {_dafny.string_of(self.decodeAddress)}, {_dafny.string_of(self.encodeLocator)}, {_dafny.string_of(self.decodeLocator)}, {_dafny.string_of(self.encodeRef)}, {_dafny.string_of(self.decodeRef)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, ContentCodecs_ContentCodecs) and self.encodeAddress == __o.encodeAddress and self.decodeAddress == __o.decodeAddress and self.encodeLocator == __o.encodeLocator and self.decodeLocator == __o.decodeLocator and self.encodeRef == __o.encodeRef and self.decodeRef == __o.decodeRef
    def __hash__(self) -> int:
        return super().__hash__()

