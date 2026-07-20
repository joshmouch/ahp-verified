#!/usr/bin/env node
// Phase 5 facade builder (ahp-dafny).
//
// `dafny translate js` emits a single file-scoped bundle whose modules are
// `let X = (function(){...})()` bindings and which auto-runs ClientMain.Main.
// That shape is runnable but NOT consumable by a separate process (nothing is
// exported). This step turns the translated bundle into a CommonJS LIBRARY: it
// strips the auto-run Main line and appends a `module.exports` of the channel
// modules + the Dafny runtime handles a consumer needs to build arguments.
//
// This is a reusable build step (a facade GENERATOR — per gate-as-feature —
// not a one-off script): the exported surface is derived from the translated
// bundle's own top-level `let <Module> =` bindings, so it tracks new channels
// automatically. Output: build/ahp_client_facade.cjs.
//
// Usage: node gen/build-js-facade.mjs
import { execFileSync } from 'node:child_process';
import { readFileSync, writeFileSync, mkdirSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

const root = dirname(dirname(fileURLToPath(import.meta.url)));  // ahp-dafny/
const buildDir = join(root, 'build');
const packageDir = join(root, '.conflux/runtime/dependencies/conflux-runtime', 'current');
mkdirSync(buildDir, { recursive: true });

// 1. Translate the multi-channel client to a persisted JS bundle (runtime included).
const rawOut = join(buildDir, 'ahp_client');
execFileSync('dafny', ['translate', 'js', join(root, 'spec', 'client_main.dfy'),
  '--no-verify',
  '--library', join(packageDir, 'conflux-runtime.doo'),
  '--translation-record', join(packageDir, 'conflux-runtime-js.dtr'),
  '--output', rawOut], { stdio: 'inherit' });

// 2. Read the bundle; discover its exported module names from `let X = (function` bindings.
const bundlePath = join(buildDir, 'ahp_client.js');
let src = readFileSync(bundlePath, 'utf8');
if (/^let Conflux(?:Contract|Codec|Channel|SeqRoute|OrderedLog)\b/m.test(src)) {
  throw new Error('facade builder: Conflux source was regenerated instead of linked from the target package');
}
const moduleNames = [...src.matchAll(/^let ([A-Z][A-Za-z0-9_]*) = \(function\(\)/gm)].map((m) => m[1]);
if (moduleNames.length === 0) throw new Error('facade builder: no top-level module bindings found — bundle shape changed');

// 3. Strip the auto-run Main line (a library must not execute on require).
src = src.replace(/^_dafny\.HandleHaltExceptions\(\(\) => ClientMain\.__default\.Main\(.*\);\s*$/m,
  '// [facade] auto-run Main stripped — consumed as a library');

// 4. Append a CommonJS export of every channel module + the runtime handles.
const exportList = ['_dafny', 'BigNumber', ...moduleNames].join(', ');
src = `${readFileSync(join(packageDir, 'conflux-runtime.cjs'), 'utf8').trimEnd()}\n\n${src}`;
src += `\nmodule.exports = { ${exportList}, ConfluxCodec: ConfluxRuntime_ConfluxCodec };\n`;

const facadePath = join(buildDir, 'ahp_client_facade.cjs');
writeFileSync(facadePath, src);
console.log(`[facade] wrote ${facadePath}`);
console.log(`[facade] exported modules: ${moduleNames.join(', ')}`);
