// Host mapper (agent-host-protocol-core-dafny) — the host-boundary consumer shape, C# lineage.
//
// A SEPARATE compilation unit from the Dafny build tool invocation: this project
// compiles the file `dafny translate cs` emitted (build/ahp_client_cs.cs) alongside
// this hand-written harness, then DRIVES the PROVEN reducers through the host
// boundary — constructing Dafny-shaped wire values, calling a reducer, and
// reading results back out. Proves the full chain end-to-end:
//
//   Dafny model (proven)  →  dafny translate cs  →  build/ahp_client_cs.cs
//       →  separate .NET console app  →  reducer runs  →  proven outcome observed
//
// This is the C# twin of mapper/ahp-host-mapper.mjs (the JS/Node mapper). Same
// four scenarios, same assertions — ported to the real generated C# API instead
// of the real generated JS API. No stubs, no mocks: every value below is a real
// Dafny-C# datatype instance and every reducer call is the real ApplyToRoot.
//
// C# needs no facade-patching step (unlike `gen/build-js-facade.mjs`, which must
// strip the JS bundle's auto-run Main and append a CommonJS export list): the
// generated .cs already exposes real public classes/namespaces, and the
// auto-generated `__CallToMain.Main` entry point is simply never invoked because
// Mapper.csproj's <StartupObject> names OUR Main instead. See
// gen/CS-HARNESS-PATTERN.md for the reusable recipe.
using System;
using System.Numerics;

namespace Mapper
{
    public static class Program
    {
        static Dafny.ISequence<Dafny.Rune> Str(string s) => Dafny.Sequence<Dafny.Rune>.UnicodeFromString(s);
        static string FromStr(Dafny.ISequence<Dafny.Rune> s) => s.ToVerbatimString(false);
        static BigInteger Int(int n) => new BigInteger(n);

        static AhpSkeleton._IRootState RootInit() =>
            AhpSkeleton.RootState.create(
                Dafny.Sequence<AhpSkeleton._IAgentInfo>.FromElements(),
                AhpSkeleton.Option<BigInteger>.create_None(),
                AhpSkeleton.Option<Dafny.ISequence<ConfluxCodec._IJson>>.create_None(),
                AhpSkeleton.Option<AhpSkeleton._IRootConfig>.create_None());

        static AhpSkeleton._IRootState ApplyRoot(AhpSkeleton._IRootState state, AhpSkeleton._IRootAction action) =>
            AhpSkeleton.__default.ApplyToRoot(state, action, Int(0)).dtor__0;

        static int failures = 0;
        static void Check(string name, bool cond)
        {
            Console.WriteLine((cond ? "PASS" : "FAIL") + "  " + name);
            if (!cond) failures++;
        }

        public static int Main(string[] args)
        {
            // scenario 1: activeSessionsChanged sets the count (proven Applied path, observed via mapper)
            var s = ApplyRoot(RootInit(), AhpSkeleton.RootAction.create_RootActiveSessionsChanged(Int(5)));
            Check("activeSessionsChanged -> activeSessions = Some(5)",
                s.dtor_activeSessions.is_Some && s.dtor_activeSessions.dtor_value == 5);

            // scenario 2: agentsChanged replaces the agent list
            var agent = AhpSkeleton.AgentInfo.create(
                Str("copilot"), Str("Copilot"), Str("AI"),
                Dafny.Sequence<Dafny.ISequence<Dafny.Rune>>.FromElements());
            var s2 = ApplyRoot(RootInit(),
                AhpSkeleton.RootAction.create_RootAgentsChanged(Dafny.Sequence<AhpSkeleton._IAgentInfo>.FromElements(agent)));
            Check("agentsChanged -> one agent, provider \"copilot\"",
                s2.dtor_agents.Count == 1 && FromStr(s2.dtor_agents.Select(0).dtor_provider) == "copilot");

            // scenario 3: an unknown action is a no-op (loud-negative — reducer returns NoOp, state unchanged)
            var unknownRes = AhpSkeleton.__default.ApplyToRoot(
                RootInit(), AhpSkeleton.RootAction.create_RootUnknown(ConfluxCodec.Json.create_JNull()), Int(0));
            Check("unknown action -> NoOp outcome + state unchanged",
                unknownRes.dtor__1.is_NoOp && object.Equals(unknownRes.dtor__0, RootInit()));

            // scenario 4: configChanged with no config present -> no-op (precondition guard, via mapper)
            var cfgRes = AhpSkeleton.__default.ApplyToRoot(
                RootInit(),
                AhpSkeleton.RootAction.create_RootConfigChanged(
                    Dafny.Map<Dafny.ISequence<Dafny.Rune>, ConfluxCodec._IJson>.Empty, false),
                Int(0));
            Check("configChanged with no config -> NoOp (precondition guard)", cfgRes.dtor__1.is_NoOp);

            if (failures == 0)
            {
                Console.WriteLine();
                Console.WriteLine("MAPPER OK — separate .NET process drove the extracted+proven Dafny reducers to their proven outcomes");
                return 0;
            }
            else
            {
                Console.WriteLine();
                Console.WriteLine($"MAPPER FAILED — {failures} check(s) failed");
                return 1;
            }
        }
    }
}
