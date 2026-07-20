// Class __default
// Dafny class __default compiled into Java
package agency.open.ahp.ResourceWatch;

@SuppressWarnings({"unchecked", "deprecation"})
public class __default {
  public __default() {
  }
  public static dafny.Tuple2<ResourceWatchState, agency.open.ahp.AhpSkeleton.ReduceOutcome> ApplyToResourceWatch(ResourceWatchState s, ResourceWatchAction a, java.math.BigInteger now)
  {
    ResourceWatchAction _source0 = a;
    if (_source0.is_RWChanged()) {
      agency.open.ahp.ConfluxCodec.Json _0___mcc_h0 = ((agency.open.ahp.ResourceWatch.ResourceWatchAction_RWChanged)_source0)._changes;
      return dafny.Tuple2.<ResourceWatchState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_Applied());
    } else {
      agency.open.ahp.ConfluxCodec.Json _1___mcc_h1 = ((agency.open.ahp.ResourceWatch.ResourceWatchAction_RWUnknown)_source0)._raw;
      return dafny.Tuple2.<ResourceWatchState, agency.open.ahp.AhpSkeleton.ReduceOutcome>create(s, agency.open.ahp.AhpSkeleton.ReduceOutcome.create_NoOp());
    }
  }
  public static boolean IsUnknownRW(ResourceWatchAction a) {
    return (a).is_RWUnknown();
  }
  public static ResourceWatchState apply1(ResourceWatchState s, ResourceWatchAction a)
  {
    return (__default.ApplyToResourceWatch(s, a, java.math.BigInteger.valueOf(9999L))).dtor__0();
  }
  public static ResourceWatchState fold(ResourceWatchState s, dafny.DafnySequence<? extends ResourceWatchAction> acts)
  {
    return agency.open.ahp.ConfluxContract.__default.<ResourceWatchState, ResourceWatchAction>Fold(ResourceWatchState._typeDescriptor(), ResourceWatchAction._typeDescriptor(), __default::apply1, s, acts);
  }
  public static dafny.Tuple2 RunCorpus()
  {
    java.math.BigInteger pass = java.math.BigInteger.ZERO;
    java.math.BigInteger total = java.math.BigInteger.ZERO;
    if(true) {
      total = java.math.BigInteger.valueOf(2L);
      pass = java.math.BigInteger.ZERO;
      ResourceWatchState _0_s0;
      _0_s0 = ResourceWatchState.create(dafny.DafnySequence.asUnicodeString("file:///workspace"), true);
      agency.open.ahp.ConfluxCodec.Json _1_changes;
      _1_changes = agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("items"), agency.open.ahp.ConfluxCodec.Json.create_JArr(dafny.DafnySequence.<agency.open.ahp.ConfluxCodec.Json> of(agency.open.ahp.ConfluxCodec.Json._typeDescriptor(), agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("uri"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("file:///workspace/a.txt"))), new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("type"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("added"))) }))))) }));
      if (java.util.Objects.equals(__default.apply1(_0_s0, ResourceWatchAction.create_RWChanged(_1_changes)), _0_s0)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
      ResourceWatchState _2_s1;
      _2_s1 = ResourceWatchState.create(dafny.DafnySequence.asUnicodeString("file:///workspace"), false);
      if (java.util.Objects.equals(__default.apply1(_2_s1, ResourceWatchAction.create_RWUnknown(agency.open.ahp.ConfluxCodec.Json.create_JObj(dafny.DafnyMap.fromElements((dafny.Tuple2<dafny.DafnySequence<? extends dafny.CodePoint>, agency.open.ahp.ConfluxCodec.Json>[])new dafny.Tuple2[]{ new dafny.Tuple2(dafny.DafnySequence.asUnicodeString("type"), agency.open.ahp.ConfluxCodec.Json.create_JStr(dafny.DafnySequence.asUnicodeString("resourceWatch/unknownAction"))) })))), _2_s1)) {
        pass = (pass).add(java.math.BigInteger.ONE);
      }
    }
    return new dafny.Tuple2<>(pass, total);
  }
  @Override
  public java.lang.String toString() {
    return "agency.open.ahp.ResourceWatch._default";
  }
}
