const { randomUUID } = require('node:crypto');

const toDafnyString = (value) => _dafny.Seq.UnicodeFromString(String(value));
const toDafnyStrings = (values) => _dafny.Seq.from(values, toDafnyString);
const toJsString = (value) => value.toVerbatimString(false);

class AhpClientError extends Error {
  constructor(message, code, data) {
    super(message);
    this.name = 'AhpClientError';
    if (code !== undefined) this.code = code;
    if (data !== undefined) this.data = data;
  }
}

function errorFromText(text) {
  try {
    const error = JSON.parse(text);
    return new AhpClientError(error.message ?? text, error.code, error.data);
  } catch {
    return new AhpClientError(text || 'AHP client operation failed');
  }
}

class AhpHostClient {
  #url;
  #clientId;
  #conn;
  #nextId = 1000;
  #agents = [];
  #actions = [];
  #observers = new Map();

  constructor(connectInfo, label = 'ahp-client') {
    if (!connectInfo?.url) throw new AhpClientError('connect-info carries no url');
    this.#url = connectInfo.url;
    this.#clientId = `${label}#${randomUUID()}`;
  }

  get agents() { return this.#agents; }

  async connect(_timeoutMs = 15_000) {
    const [ok, conn, resultText, notifications, errorText] =
      AhpConnectionRuntime.__default.Connect(
        toDafnyString(this.#url),
        toDafnyString(this.#clientId),
        toDafnyStrings(['ahp-root://']),
      );
    this.#receiveAll(notifications);
    if (!ok) throw errorFromText(toJsString(errorText));
    this.#conn = conn;
    const result = JSON.parse(toJsString(resultText));
    const root = result.snapshots?.find((snapshot) => snapshot.resource === 'ahp-root://')?.state;
    this.#agents = (root?.agents ?? [])
      .map((agent) => typeof agent === 'string' ? agent : agent?.provider)
      .filter((provider) => typeof provider === 'string');
  }

  async request(method, params = {}) {
    if (this.#conn === undefined) throw new AhpClientError('not connected to an AHP host');
    const id = this.#nextId++;
    const [ok, resultText, notifications, errorText] =
      AhpConnectionRuntime.__default.Request(
        this.#conn,
        new BigNumber(id),
        toDafnyString(method),
        toDafnyString(JSON.stringify(params)),
      );
    this.#receiveAll(notifications);
    if (!ok) throw errorFromText(toJsString(errorText));
    return JSON.parse(toJsString(resultText));
  }

  async createChat(provider, cwd) {
    return await this.request('chat/create', { provider, cwd });
  }

  async prompt(chat, text, onAction) {
    const before = this.#actions.length;
    if (onAction) this.#observers.set(chat.channel, onAction);
    try {
      await this.request('chat/prompt', { chatId: chat.chatId, text });
    } finally {
      this.#observers.delete(chat.channel);
    }
    const actions = this.#actions.slice(before)
      .filter((action) => action.channel === chat.channel)
      .sort((left, right) => left.serverSeq - right.serverSeq);
    const join = (type) => actions
      .filter((action) => action.type === type)
      .map((action) => action.content ?? '')
      .join('');
    const terminal = actions.filter((action) =>
      action.type === 'chat/turnComplete'
      || action.type === 'chat/turnCancelled'
      || action.type === 'chat/error');
    return {
      chatId: chat.chatId,
      text: join('chat/delta'),
      reasoning: join('chat/reasoning'),
      actions,
      outcome: terminal.at(-1)?.type ?? 'chat/incomplete',
    };
  }

  cancel(chat) {
    void this.request('chat/cancel', { chatId: chat.chatId }).catch(() => undefined);
  }

  close() {
    if (this.#conn !== undefined) AhpConnectionRuntime.__default.Close(this.#conn);
    this.#conn = undefined;
  }

  #receiveAll(notifications) {
    for (const raw of notifications) this.#receive(toJsString(raw));
  }

  #receive(text) {
    let message;
    try { message = JSON.parse(text); } catch { return; }
    if (message.method !== 'action') return;
    const envelope = message.params;
    const action = envelope?.action;
    if (!action || typeof action.type !== 'string') return;
    const received = {
      channel: String(envelope?.channel ?? ''),
      serverSeq: Number(envelope?.serverSeq ?? 0),
      type: action.type,
      ...(typeof action.turnId === 'string' ? { turnId: action.turnId } : {}),
      ...(typeof action.content === 'string' ? { content: action.content } : {}),
    };
    this.#actions.push(received);
    try { this.#observers.get(received.channel)?.(received); } catch {}
  }
}

module.exports = {
  AhpClientError,
  AhpHostClient,
  verified: { AhpConnection, AhpConnectionRuntime, Client, Version },
};
