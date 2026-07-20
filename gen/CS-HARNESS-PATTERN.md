# C# harness pattern (reusable recipe)

C# is the PRIMARY run/proof lineage for throwaway/consumer harnesses here
(faster than JS/bignumber.js; `dafny translate cs` output is directly
consumable — real public classes, no facade patch needed unlike
`gen/build-js-facade.mjs`). JS is still generated too, for cross-lineage
parity. Worked examples: `mapper/` (single fixed harness, hand-written C#
entry point + `<StartupObject>`) and `gen/run-reducer-replay.sh` +
`gen/ReplayHarness.csproj` (N harnesses sharing ONE project, parameterized by
an MSBuild property — see §6 below).

## 1. Translate
```bash
dafny translate cs spec/client_main.dfy --no-verify --include-runtime \
  --output build/ahp_client_cs
```
Emits `build/ahp_client_cs.cs` (+ a `.dtr` cache marker). Re-run whenever the
`.dfy` sources change — the `.cs` is gitignored (`build/`, `*.cs`).

## 2. The `.csproj` shape
Target `net10.0`; no `PackageReference` needed — `BigInteger`/`ImmutableDictionary`
resolve from the shared framework (the older `dafny run`-generated
`spec/client_main.csproj`, net8.0, still carries them explicitly).
```xml
<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>net10.0</TargetFramework>
    <EnableDefaultCompileItems>false</EnableDefaultCompileItems>
    <StartupObject>Mapper.Program</StartupObject>
  </PropertyGroup>
  <ItemGroup>
    <Compile Include="AhpHostMapper.cs" />
    <Compile Include="../build/ahp_client_cs.cs" />
  </ItemGroup>
</Project>
```

## 3. Reference classes, don't run Main
`dafny translate`/`run` always emits a top-level `class __CallToMain { static
void Main(...) {...} }` that auto-runs the `.dfy` file's own `Main`. Two
`Main`s in one compile is a C# ambiguous-entry-point error — set
`<StartupObject>` to YOUR class instead. `__CallToMain` still compiles (types
stay reachable) but never runs — no text-patching needed (unlike the JS
facade, which must strip the auto-run line).

## 4. Dafny -> C# cheatsheet
| Dafny | C# |
|---|---|
| `int`/`nat` | `System.Numerics.BigInteger` |
| `string` | `Dafny.ISequence<Dafny.Rune>`; build `.UnicodeFromString("s")`; read `.ToVerbatimString(false)` |
| `seq<T>` | `Dafny.ISequence<T>`; build `Dafny.Sequence<T>.FromElements(...)`; index `.Select(int)`; len `.Count` |
| `map<K,V>` | `Dafny.IMap<K,V>`; build `Dafny.Map<K,V>.FromElements(new Dafny.Pair<K,V>(k,v),...)`; empty `.Empty` |
| datatype, 1 ctor | `Ns.T.create(f1, f2, ...)` (also `create_T(...)`) |
| datatype, N ctors | `Ns.T.create_Variant(...)`; test `.is_Variant`; read `.dtor_field` |
| `Option<T>` | `.create_None()` / `.create_Some(v)`; `.is_Some`/`.is_None`; `.dtor_value` |
| tuple `(A, B)` | `_System._ITuple2<A,B>`; read `.dtor__0`, `.dtor__1` (double underscore) |
| equality | `object.Equals(a, b)` (datatypes override `Equals`) |

## 5. Run
```bash
dotnet run --project mapper/Mapper.csproj
```

## 6. Variant: one parameterized project for N interchangeable harnesses

When N `.dfy` files share the same shape (one `Main`, no hand-written C#
entry point needed — see the reducer-replay corpus, `spec/replay/*.dfy`, 7
channels), don't hand-author N near-identical `.csproj` files and don't
regenerate a `.csproj` per run (that's the transient `dafny run` shape this
pattern replaces). Instead author ONE persistent `.csproj`, parameterized by
an MSBuild property, and pass a different value per invocation:

```xml
<!-- gen/ReplayHarness.csproj -->
<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>net10.0</TargetFramework>
    <EnableDefaultCompileItems>false</EnableDefaultCompileItems>
  </PropertyGroup>
  <ItemGroup>
    <Compile Include="../build/$(Channel)_replay_cs.cs" />
  </ItemGroup>
</Project>
```

```bash
dafny translate cs --standard-libraries --translate-standard-library --no-verify --include-runtime \
  --output build/resourcewatch_replay_cs spec/replay/resourcewatch_replay.dfy
dotnet run --project gen/ReplayHarness.csproj -p:Channel=resourcewatch -- <fixture paths...>
```

No `<StartupObject>` is needed here — each translated `.cs` file defines
exactly one `Main` (the top-level `__CallToMain` class from §3), and only one
`.cs` file is compiled per invocation (the `$(Channel)`-parameterized
`<Compile Include>` swaps which file that is). Switching `-p:Channel=` between
runs correctly forces a rebuild even when reusing the same shared `obj/`/`bin/`
— MSBuild's up-to-date check sees a different resolved `Compile` item —
verified empirically switching resourcewatch → session → resourcewatch with
no `dotnet clean` between runs. `.gitignore` carves out `!gen/*.csproj` the
same way it carves out `!mapper/*.csproj`; the generated `.cs` stays
untracked under `build/`.
