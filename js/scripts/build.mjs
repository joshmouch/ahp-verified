#!/usr/bin/env node
// Build the @open-agency/ahp CommonJS distribution from the formally-verified
// Dafny core.
//
// `dafny translate js` emits a single file-scoped bundle whose modules are
// `let X = (function(){...})()` bindings, which references the Dafny + Conflux
// runtime as free variables (`_dafny`, `BigNumber`, `ConfluxRuntime_*`) and
// which auto-runs ClientMain.Main on load. That shape is runnable as a script
// but is NOT consumable as a library: nothing is exported and requiring it
// would execute the corpus runner.
//
// This step turns the translated bundle into a CommonJS LIBRARY:
//   1. translate the multi-channel client against the Conflux runtime library
//   2. prepend the runtime bundle so the free variables resolve
//   3. strip the auto-run Main line
//   4. append a module.exports of the channel modules + the runtime handles
//
// The exported surface is DERIVED from the translated bundle's own top-level
// `let <Module> =` bindings, so new channels are picked up automatically
// rather than needing a hand-maintained export list.
//
// Usage: node scripts/build.mjs

import { execFileSync } from 'node:child_process';
import { readFileSync, writeFileSync, mkdirSync, statSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

const here = dirname(fileURLToPath(import.meta.url));
const pkgRoot = dirname(here);

// The formally-verified Dafny core (read-only input to this build).
const coreRoot =
  process.env.AHP_DAFNY_CORE ??
  '/Users/josh/Code/joshmouch/conflux/agent-host-protocol-workspace/agent-host-protocol-core-dafny';

// The vendored Conflux runtime the core depends on. `conflux-runtime.doo` is the
// library the translation links against; `conflux-runtime-js.dtr` is the JS
// translation record that tells Dafny the library was already translated for
// this backend (without it the translation re-emits Conflux source instead of
// linking to it); `conflux-runtime.cjs` is the already-translated runtime body.
const runtimeDir = join(coreRoot, '.conflux/runtime/dependencies/conflux-runtime/current');
const libraryDoo = join(runtimeDir, 'conflux-runtime.doo');
const translationRecord = join(runtimeDir, 'conflux-runtime-js.dtr');
const runtimeCjs = join(runtimeDir, 'conflux-runtime.cjs');

const workDir = join(pkgRoot, '.build');
const distDir = join(pkgRoot, 'dist');
mkdirSync(workDir, { recursive: true });
mkdirSync(distDir, { recursive: true });

// 1. Translate the multi-channel client to a persisted JS bundle.
const rawOut = join(workDir, 'ahp_client');
execFileSync(
  'dafny',
  [
    'translate', 'js', join(coreRoot, 'spec', 'client_main.dfy'),
    '--no-verify',
    '--library', libraryDoo,
    '--translation-record', translationRecord,
    '--output', rawOut,
  ],
  { stdio: 'inherit' },
);

// 2. Read the bundle; discover its module names from `let X = (function` bindings.
const bundlePath = join(workDir, 'ahp_client.js');
let src = readFileSync(bundlePath, 'utf8');

// Guard: if the Conflux runtime source was re-emitted into our bundle rather
// than linked from the target package, the translation record was not honored
// and we would ship two divergent copies of the runtime.
if (/^let Conflux(?:Contract|Codec|Channel|SeqRoute|OrderedLog)\b/m.test(src)) {
  throw new Error('build: Conflux source was regenerated instead of linked from the target package');
}

const moduleNames = [...src.matchAll(/^let ([A-Z][A-Za-z0-9_]*) = \(function\(\)/gm)].map((m) => m[1]);
if (moduleNames.length === 0) {
  throw new Error('build: no top-level module bindings found — bundle shape changed');
}

// 3. Strip the auto-run Main line (a library must not execute on require).
const mainLine = /^_dafny\.HandleHaltExceptions\(\(\) => ClientMain\.__default\.Main\(.*\);\s*$/m;
if (!mainLine.test(src)) {
  throw new Error('build: auto-run Main line not found — bundle shape changed');
}
src = src.replace(mainLine, '// [dist] auto-run Main stripped — consumed as a library');

// 4. Prepend the runtime, append the CommonJS export surface.
const exportList = ['_dafny', 'BigNumber', ...moduleNames].join(', ');
src = `${readFileSync(runtimeCjs, 'utf8').trimEnd()}\n\n${src}`;
src += `\nmodule.exports = { ${exportList}, ConfluxCodec: ConfluxRuntime_ConfluxCodec };\n`;

const outPath = join(distDir, 'ahp.cjs');
writeFileSync(outPath, src);

console.log(`[build] wrote ${outPath} (${statSync(outPath).size} bytes)`);
console.log(`[build] exported modules: ${moduleNames.join(', ')}`);
