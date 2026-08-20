const assert = require('node:assert/strict');
const { spawn } = require('node:child_process');
const { AhpHostClient, verified } = require('..');

async function main() {
  for (const name of ['AhpConnection', 'AhpConnectionRuntime', 'AhpSessionClient', 'Client', 'Version']) {
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
    await assert.rejects(
      client.request('missing-method'),
      (error) => error.code === -32601 && error.message === 'unknown method missing-method',
    );
    const chat = await client.createChat('grok', '/tmp');
    const observed = [];
    const turn = await client.prompt(chat, 'hello', (action) => observed.push(action.type));
    assert.equal(turn.text, 'VERIFIED-CLIENT-OK');
    assert.equal(turn.outcome, 'chat/turnComplete');
    assert.deepEqual(observed, ['chat/delta', 'chat/turnComplete']);
    assert.equal(client.cancel(chat), true);
    client.close();

    const standard = new AhpHostClient({ url: 'ws://127.0.0.1:49514' }, 'standard-smoke');
    await standard.connect();
    const standardChat = await standard.createChat(
      'copilotcli',
      '/tmp',
      { config: { isolation: 'folder' } },
    );
    assert.equal(standardChat.transport, 'dispatch-action');
    assert.equal(standardChat.chatId, 'ahp-chat:/standard');
    const standardObserved = [];
    const standardTurn = await standard.prompt(
      standardChat,
      'standard hello',
      (action) => standardObserved.push(action.type),
    );
    assert.equal(standardTurn.text, 'VERIFIED-STANDARD-AHP-OK');
    assert.equal(standardTurn.reasoning, 'verified-reasoning');
    assert.equal(standardTurn.outcome, 'chat/turnComplete');
    assert.deepEqual(standardObserved, ['chat/delta', 'chat/reasoning', 'chat/turnComplete']);
    assert.deepEqual(standardTurn.actions.map((action) => action.serverSeq), [1, 2, 3]);
    assert.equal(standardTurn.actions[0].rawAction.content, 'VERIFIED-STANDARD-AHP-OK');
    assert.equal(standard.cancel(standardChat), true);
    standard.close();

    const cancellable = new AhpHostClient({ url: 'ws://127.0.0.1:49514' }, 'cancel-smoke');
    await cancellable.connect();
    const cancellableChat = await cancellable.createChat(
      'copilotcli',
      '/tmp',
      { config: { isolation: 'folder' } },
    );
    const cancelObserved = [];
    const cancelledPromise = cancellable.prompt(
      cancellableChat,
      'cancel me',
      (action) => cancelObserved.push(action.type),
    );
    await assert.rejects(
      cancellable.prompt(cancellableChat, 'concurrent turn'),
      /turn already active/,
    );
    await new Promise((resolve) => setTimeout(resolve, 75));
    assert.equal(cancellable.cancel(cancellableChat), true);
    const cancelledTurn = await cancelledPromise;
    assert.equal(cancelledTurn.outcome, 'chat/turnCancelled');
    assert.deepEqual(cancelObserved, ['chat/turnCancelled']);
    cancellable.close();

    const resumed = new AhpHostClient({ url: 'ws://127.0.0.1:49514' }, 'resume-smoke');
    await resumed.connect();
    await resumed.attachChat(standardChat);
    const resumedTurn = await resumed.prompt(standardChat, 'resumed hello');
    assert.equal(resumedTurn.text, 'VERIFIED-STANDARD-RESUME-OK');
    assert.equal(resumedTurn.outcome, 'chat/turnComplete');
    resumed.close();

    const repeated = new AhpHostClient(
      { url: 'ws://127.0.0.1:49514' },
      'repeat-version',
    );
    await assert.rejects(
      repeated.connect(),
      /host repeated an unsupported version offer/,
    );
  } finally {
    server.kill();
  }

  console.log('SMOKE PASSED — extracted connect/request, both turn surfaces, ordered actions, live standard cancel, resume, and reconnecting -32005 recovery');
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
