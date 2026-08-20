const { randomUUID } = require('node:crypto');
const { pathToFileURL } = require('node:url');

const toDafnyString = (value) => _dafny.Seq.UnicodeFromString(String(value));
const toDafnyStrings = (values) => _dafny.Seq.from(values, toDafnyString);
const toJsString = (value) => value.toVerbatimString(false);
const yieldToCaller = () => new Promise((resolve) => setImmediate(resolve));

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
  #state = AhpConnectionRuntime.__default.InitialState();
  #agents = [];

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
    if (!ok) throw errorFromText(toJsString(errorText));
    this.#conn = conn;
    const result = JSON.parse(toJsString(resultText));
    const root = result.snapshots?.find((snapshot) => snapshot.resource === 'ahp-root://')?.state;
    this.#agents = (root?.agents ?? [])
      .map((agent) => typeof agent === 'string' ? agent : agent?.provider)
      .filter((provider) => typeof provider === 'string');
  }

  async request(method, params = {}) {
    this.#requireConnection();
    const [ok, next, resultText, _notifications, errorText] =
      AhpConnectionRuntime.__default.PublicRequest(
        this.#conn,
        this.#state,
        toDafnyString(method),
        toDafnyString(JSON.stringify(params)),
      );
    this.#state = next;
    if (!ok) throw errorFromText(toJsString(errorText));
    return JSON.parse(toJsString(resultText));
  }

  async createChat(provider, cwd, options = {}) {
    this.#requireConnection();
    const [ok, next, chatText, _notifications, errorText] =
      AhpConnectionRuntime.__default.CreateChat(
        this.#conn,
        this.#state,
        toDafnyString(provider),
        toDafnyString(cwd),
        toDafnyString(pathToFileURL(cwd).href),
        toDafnyString(`ahp-session:/${randomUUID()}`),
        toDafnyString(JSON.stringify(options.config ?? {})),
      );
    this.#state = next;
    if (!ok) throw errorFromText(toJsString(errorText));
    return JSON.parse(toJsString(chatText));
  }

  async attachChat(chat) {
    this.#requireConnection();
    const [ok, next, _notifications, errorText] =
      AhpConnectionRuntime.__default.AttachChat(
        this.#conn,
        this.#state,
        toDafnyString(JSON.stringify(chat)),
      );
    this.#state = next;
    if (!ok) throw errorFromText(toJsString(errorText));
    return chat;
  }

  async prompt(chat, text, onAction) {
    this.#requireConnection();
    const turnId = randomUUID();
    const [ok, next, pending, view, resultText, _notifications, errorText] =
      AhpConnectionRuntime.__default.BeginPrompt(
        this.#conn,
        this.#state,
        toDafnyString(JSON.stringify(chat)),
        toDafnyString(text),
        toDafnyString(turnId),
        toDafnyString(new Date().toISOString()),
      );
    this.#state = next;
    if (!ok) throw errorFromText(toJsString(errorText));
    let currentView = view;
    let currentPending = pending;
    let observed = 0;
    let result;
    const publish = (textResult) => {
      result = JSON.parse(toJsString(textResult));
      if (onAction) {
        for (const action of result.actions.slice(observed)) onAction(action);
      }
      observed = result.actions.length;
    };
    if (!currentPending) {
      publish(resultText);
      return result;
    }

    // BeginPrompt has already recorded the active turn in the extracted state.
    // BeginPrompt and every bounded ReceiveTurn return to this interop loop, so
    // cancellation can run against the exact extracted active-turn state even
    // when the host remains silent.
    await yieldToCaller();
    while (currentPending) {
      const [received, receivedState, stillPending, nextView, nextResult,
        _receivedNotifications, receiveError] =
        AhpConnectionRuntime.__default.ReceiveTurn(
          this.#conn,
          this.#state,
          toDafnyString(JSON.stringify(chat)),
          toDafnyString(turnId),
          currentView,
        );
      this.#state = receivedState;
      if (!received) throw errorFromText(toJsString(receiveError));
      currentView = nextView;
      currentPending = stillPending;
      publish(nextResult);
      if (currentPending) await yieldToCaller();
    }
    return result;
  }

  cancel(chat) {
    if (this.#conn === undefined) return false;
    const [ok, next] = AhpConnectionRuntime.__default.Cancel(
      this.#conn,
      this.#state,
      toDafnyString(JSON.stringify(chat)),
    );
    this.#state = next;
    return ok;
  }

  close() {
    if (this.#conn !== undefined) AhpConnectionRuntime.__default.Close(this.#conn);
    this.#conn = undefined;
  }

  #requireConnection() {
    if (this.#conn === undefined) throw new AhpClientError('not connected to an AHP host');
  }
}

module.exports = {
  AhpClientError,
  AhpHostClient,
  verified: { AhpConnection, AhpConnectionRuntime, AhpSessionClient, Client, Version },
};
