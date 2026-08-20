const assert = require('node:assert/strict');
const { spawn } = require('node:child_process');
const { AhpHostClient, verified } = require('..');

async function main() {
  for (const name of ['AhpConnection', 'AhpConnectionRuntime', 'Client', 'Version']) {
    assert.ok(verified[name], `missing extracted module ${name}`);
  }

  const server = spawn(process.execPath, [require.resolve('./server.cjs')], {
    stdio: ['ignore', 'pipe', 'inherit'],
  });
  await new Promise((resolve, reject) => {
    server.once('exit', (code) => reject(new Error(`test server exited ${code}`)));
    server.stdout.once('data', (chunk) => {
      if (String(chunk).includes('ready')) resolve();
      else reject(new Error(`unexpected server readiness: ${chunk}`));
    });
  });

  try {
    const client = new AhpHostClient({ url: 'ws://127.0.0.1:49514' }, 'smoke');
    await client.connect();
    assert.deepEqual(client.agents, ['grok']);
    const chat = await client.createChat('grok', '/tmp');
    const observed = [];
    const turn = await client.prompt(chat, 'hello', (action) => observed.push(action.type));
    assert.equal(turn.text, 'VERIFIED-CLIENT-OK');
    assert.equal(turn.outcome, 'chat/turnComplete');
    assert.deepEqual(observed, ['chat/delta', 'chat/turnComplete']);
    client.close();
  } finally {
    server.kill();
  }

  console.log('SMOKE PASSED — extracted connect/request/turn and -32005 recovery');
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
