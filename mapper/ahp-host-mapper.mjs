#!/usr/bin/env node
// Host mapper (agent-host-protocol-core-dafny) — the host-boundary consumer shape.
//
// A SEPARATE process from the Dafny build: it `require`s the generated facade
// (build/ahp_client_facade.cjs) and drives the PROVEN reducers through the host
// boundary — constructing Dafny-shaped wire values, calling a reducer, and
// reading results back out. Proves the full chain end-to-end:
//
//   Dafny model (proven)  →  dafny translate js  →  facade (module.exports)
//       →  separate node consumer  →  reducer runs  →  proven outcome observed
//
// This provider-side mapper drives the pure-AHP `root` channel. (The OpenAgency
// extension repo has its own mapper driving the OA channels.)
import { createRequire } from 'node:module';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

const require = createRequire(import.meta.url);
const root = dirname(dirname(fileURLToPath(import.meta.url)));
const { _dafny, BigNumber, AhpSkeleton, ConfluxCodec } = require(join(root, 'build', 'ahp_client_facade.cjs'));

const str = (s) => _dafny.Seq.UnicodeFromString(s);
const fromStr = (d) => d.toVerbatimString(false);
const int = (n) => new BigNumber(n);
const None = () => AhpSkeleton.Option.create_None();
const applyRoot = (state, action) => AhpSkeleton.__default.ApplyToRoot(state, action, int(0))[0];
const rootInit = () => AhpSkeleton.RootState.create_RootState(_dafny.Seq.of(), None(), None(), None());

let failures = 0;
const check = (name, cond) => { console.log(`${cond ? 'PASS' : 'FAIL'}  ${name}`); if (!cond) failures++; };

// scenario 1: activeSessionsChanged sets the count (proven Applied path, observed via mapper)
let s = applyRoot(rootInit(), AhpSkeleton.RootAction.create_RootActiveSessionsChanged(int(5)));
check('activeSessionsChanged → activeSessions = Some(5)',
  s.dtor_activeSessions.is_Some && s.dtor_activeSessions.dtor_value.toNumber() === 5);

// scenario 2: agentsChanged replaces the agent list
const agent = AhpSkeleton.AgentInfo.create_AgentInfo(str('copilot'), str('Copilot'), str('AI'), _dafny.Seq.of());
const s2 = applyRoot(rootInit(), AhpSkeleton.RootAction.create_RootAgentsChanged(_dafny.Seq.of(agent)));
check('agentsChanged → one agent, provider "copilot"',
  s2.dtor_agents.length === 1 && fromStr(s2.dtor_agents[0].dtor_provider) === 'copilot');

// scenario 3: an unknown action is a no-op (loud-negative — reducer returns NoOp, state unchanged)
const unknownRes = AhpSkeleton.__default.ApplyToRoot(rootInit(), AhpSkeleton.RootAction.create_RootUnknown(ConfluxCodec.Json.create_JNull()), int(0));
check('unknown action → NoOp outcome + state unchanged',
  unknownRes[1].is_NoOp && _dafny.areEqual(unknownRes[0], rootInit()));

// scenario 4: configChanged with no config present → no-op (precondition guard, via mapper)
const cfgRes = AhpSkeleton.__default.ApplyToRoot(rootInit(), AhpSkeleton.RootAction.create_RootConfigChanged(_dafny.Map.Empty, false), int(0));
check('configChanged with no config → NoOp (precondition guard)', cfgRes[1].is_NoOp);

console.log(failures === 0
  ? '\nMAPPER OK — separate node process drove the extracted+proven Dafny reducers to their proven outcomes'
  : `\nMAPPER FAILED — ${failures} check(s) failed`);
process.exit(failures === 0 ? 0 : 1);
