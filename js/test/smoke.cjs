// Smoke test for the @open-agency/ahp distribution.
//
// Proves three things about the shipped artifact:
//   1. it LOADS as a CommonJS library without executing anything
//   2. every AHP channel module is present on the export surface
//   3. the extracted reducers actually RUN — each channel replays its slice of
//      the fixture corpus that the Dafny core is proven against, and every
//      fixture must come back green

const assert = require('node:assert');
const ahp = require('..');

const CHANNELS = [
  'AhpSkeleton',
  'ResourceWatch',
  'Canvas',
  'Changeset',
  'Annotations',
  'Terminal',
  'Session',
  'Chat',
];

// 1 + 2. Export surface.
for (const name of CHANNELS) {
  assert.ok(ahp[name], `missing channel module: ${name}`);
  assert.strictEqual(
    typeof ahp[name].__default.RunCorpus,
    'function',
    `${name} has no RunCorpus entry point`,
  );
}
assert.ok(ahp._dafny, 'missing _dafny runtime handle');
assert.ok(ahp.BigNumber, 'missing BigNumber handle');
assert.ok(ahp.ConfluxCodec, 'missing ConfluxCodec handle');
console.log(`export surface OK — ${CHANNELS.length} channel modules + runtime handles`);

// 3. Replay the reducer corpus through the extracted code.
let totalPass = 0;
let totalCount = 0;
for (const name of CHANNELS) {
  const [pass, count] = ahp[name].__default.RunCorpus();
  const p = Number(pass.toString());
  const c = Number(count.toString());
  totalPass += p;
  totalCount += c;
  console.log(`  ${name.padEnd(14)} ${p}/${c}`);
  assert.strictEqual(p, c, `${name}: ${c - p} fixture(s) failed`);
}
console.log(`corpus OK — ${totalPass}/${totalCount} fixtures green against extracted code`);

// 4. Construct a datatype + call a pure reducer helper directly, proving the
//    extracted value layer is usable and not just loadable.
const emptyTurns = ahp._dafny.Seq.of();
assert.strictEqual(
  ahp.Chat.__default.hasTurn(emptyTurns, ahp._dafny.Seq.UnicodeFromString('missing-id')),
  false,
  'Chat.hasTurn should be false on an empty turn sequence',
);
console.log('direct reducer call OK — Chat.hasTurn([], "missing-id") === false');

console.log('\nSMOKE TEST PASSED');
