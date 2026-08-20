#!/usr/bin/env node
import { execFileSync } from 'node:child_process';
import { mkdirSync, readFileSync, statSync, writeFileSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

const here = dirname(fileURLToPath(import.meta.url));
const pkgRoot = dirname(here);
const workspace = process.env.AHP_WORKSPACE
  ?? '/Users/josh/Code/joshmouch/conflux/agent-host-protocol-workspace';
const core = process.env.AHP_DAFNY_CORE
  ?? join(workspace, 'agent-host-protocol-core-dafny');
const client = process.env.AHP_DAFNY_CLIENT
  ?? join(workspace, 'agent-host-protocol-client-dafny');
const runtime = join(core, '.conflux/runtime/dependencies/conflux-runtime/current');
const runtimeDoo = join(runtime, 'conflux-runtime.doo');
const runtimeRecord = join(runtime, 'conflux-runtime-js.dtr');
const runtimeCjs = join(runtime, 'conflux-runtime.cjs');
const coreDoo = join(client, 'vendor/agent-host-protocol-core.doo');
const clientDoo = join(client, 'agent-host-protocol-client.doo');
const build = join(pkgRoot, '.build');
const dist = join(pkgRoot, 'dist');
mkdirSync(build, { recursive: true });
mkdirSync(dist, { recursive: true });

execFileSync('bash', [join(client, 'gen/build-doo.sh')], { stdio: 'inherit' });

const translate = (source, output, libraries = []) => execFileSync('dafny', [
  'translate', 'js', source, '--no-verify',
  ...libraries.flatMap((library) => ['--library', library]),
  '--library', runtimeDoo,
  '--translation-record', runtimeRecord,
  '--output', join(build, output),
], { stdio: 'inherit' });

translate(join(core, 'spec/core_lib.dfy'), 'core');
translate(join(client, 'spec/client.dfy'), 'client', [coreDoo]);
translate(join(client, 'run/connection_runtime.dfy'), 'runtime', [clientDoo, coreDoo]);

const generated = ['core.js', 'client.js', 'runtime.js'].map((name) =>
  readFileSync(join(build, name), 'utf8')
    .replace(/let _module = \(function\(\) \{\n  let \$module = \{\};\n\n  return \$module;\n\}\)\(\); \/\/ end of module _module\n?/g, ''));
const combined = generated.join('\n');
for (const required of [
  'let Version =', 'let Client =', 'let AhpConnection =',
  'let AhpSessionClient =', 'let AhpConnectionRuntime =',
]) {
  if (!combined.includes(required)) throw new Error(`missing extracted module: ${required}`);
}
if (/^let ConfluxRuntime_/m.test(combined)) {
  throw new Error('Conflux runtime source was regenerated instead of linked');
}

const wrapper = readFileSync(join(pkgRoot, 'src/wrapper.cjs'), 'utf8');
const output = join(dist, 'client.cjs');
writeFileSync(output, `${readFileSync(runtimeCjs, 'utf8').trimEnd()}\n\n${combined}\n${wrapper}`);
console.log(`[build] wrote ${output} (${statSync(output).size} bytes)`);
