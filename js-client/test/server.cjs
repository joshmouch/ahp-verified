const { WebSocketServer } = require('ws');

const expectedInitial = ['0.8.0', '0.7.0', '0.6.0', '0.5.2', '0.5.1'];
const server = new WebSocketServer({ host: '127.0.0.1', port: 49514 });

server.on('connection', (socket) => {
  let initializeAttempts = 0;
  socket.on('message', (raw) => {
    const request = JSON.parse(String(raw));
    if (request.method === 'initialize') {
      initializeAttempts += 1;
      const offered = request.params.protocolVersions;
      if (initializeAttempts === 1) {
        if (JSON.stringify(offered) !== JSON.stringify(expectedInitial)) process.exit(21);
        socket.send(JSON.stringify({
          jsonrpc: '2.0', id: request.id,
          error: {
            code: -32005,
            message: 'unsupported protocol version',
            data: { supportedVersions: ['9.0.0'] },
          },
        }));
        return;
      }
      if (JSON.stringify(offered) !== JSON.stringify(['9.0.0'])) process.exit(22);
      socket.send(JSON.stringify({
        jsonrpc: '2.0', id: request.id,
        result: {
          protocolVersion: '9.0.0', serverSeq: 0,
          snapshots: [{ resource: 'ahp-root://', state: { agents: [{ provider: 'grok' }] } }],
        },
      }));
      return;
    }
    if (request.method === 'chat/create') {
      if (request.params.provider === 'copilotcli') {
        socket.send(JSON.stringify({
          jsonrpc: '2.0', id: request.id,
          error: { code: -32601, message: 'Method not found: chat/create' },
        }));
        return;
      }
      socket.send(JSON.stringify({
        jsonrpc: '2.0', id: request.id,
        result: { sessionId: 's1', chatId: 'c1', agentId: 'grok', channel: 'ahp-chat:/c1' },
      }));
      return;
    }
    if (request.method === 'createSession') {
      if (request.params.provider !== 'copilotcli'
          || request.params.channel.indexOf('ahp-session:/') !== 0
          || JSON.stringify(request.params.workingDirectories) !== JSON.stringify(['file:///tmp'])) {
        process.exit(23);
      }
      socket.send(JSON.stringify({ jsonrpc: '2.0', id: request.id, result: null }));
      return;
    }
    if (request.method === 'subscribe') {
      const channel = request.params.channel;
      const state = channel.indexOf('ahp-session:/') === 0
        ? { defaultChat: 'ahp-chat:/standard' }
        : { status: 'idle' };
      socket.send(JSON.stringify({
        jsonrpc: '2.0', id: request.id,
        result: { snapshot: { resource: channel, state } },
      }));
      return;
    }
    if (request.method === 'dispatchAction') {
      const action = request.params.action;
      if (request.id !== undefined
          || request.params.channel !== 'ahp-chat:/standard'
          || request.params.clientSeq !== 1
          || action.type !== 'chat/turnStarted'
          || action.message.text !== 'standard hello'
          || action.message.origin.kind !== 'user') {
        process.exit(24);
      }
      socket.send(JSON.stringify({
        jsonrpc: '2.0', method: 'action',
        params: { channel: request.params.channel, serverSeq: 1,
          action: { type: 'chat/delta', turnId: action.turnId, content: 'VERIFIED-STANDARD-AHP-OK' } },
      }));
      socket.send(JSON.stringify({
        jsonrpc: '2.0', method: 'action',
        params: { channel: request.params.channel, serverSeq: 2,
          action: { type: 'chat/turnComplete', turnId: action.turnId } },
      }));
      return;
    }
    if (request.method === 'chat/prompt') {
      socket.send(JSON.stringify({
        jsonrpc: '2.0', method: 'action',
        params: { channel: 'ahp-chat:/c1', serverSeq: 1,
          action: { type: 'chat/delta', turnId: 't1', content: 'VERIFIED-CLIENT-OK' } },
      }));
      socket.send(JSON.stringify({
        jsonrpc: '2.0', method: 'action',
        params: { channel: 'ahp-chat:/c1', serverSeq: 2,
          action: { type: 'chat/turnComplete', turnId: 't1' } },
      }));
      socket.send(JSON.stringify({ jsonrpc: '2.0', id: request.id, result: {} }));
      return;
    }
    socket.send(JSON.stringify({
      jsonrpc: '2.0', id: request.id,
      error: { code: -32601, message: `unknown method ${request.method}` },
    }));
  });
});

server.on('listening', () => process.stdout.write('ready\n'));
