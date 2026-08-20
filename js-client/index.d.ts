export interface AhpConnectInfo {
  readonly url: string;
  readonly token?: string;
  readonly agents?: readonly string[];
}

export interface AhpChat {
  readonly sessionId: string;
  readonly chatId: string;
  readonly agentId: string;
  readonly channel: string;
  readonly transport?: 'dispatch-action';
}

export interface AhpAction {
  readonly channel: string;
  readonly serverSeq: number;
  readonly type: string;
  readonly rawAction: Readonly<Record<string, unknown>>;
  readonly turnId?: string;
  readonly content?: string;
}

export interface AhpTurnResult {
  readonly chatId: string;
  readonly text: string;
  readonly reasoning: string;
  readonly actions: readonly AhpAction[];
  readonly outcome: string;
}

export class AhpClientError extends Error {
  readonly code?: number;
  readonly data?: unknown;
}

export class AhpHostClient {
  constructor(connectInfo: AhpConnectInfo, label?: string);
  get agents(): readonly string[];
  connect(timeoutMs?: number): Promise<void>;
  request(method: string, params?: unknown): Promise<unknown>;
  createChat(provider: string, cwd: string): Promise<AhpChat>;
  prompt(chat: AhpChat, text: string, onAction?: (action: AhpAction) => void): Promise<AhpTurnResult>;
  cancel(chat: AhpChat): void;
  close(): void;
}

export const verified: {
  readonly AhpConnection: unknown;
  readonly AhpConnectionRuntime: unknown;
  readonly Client: unknown;
  readonly Version: unknown;
};
