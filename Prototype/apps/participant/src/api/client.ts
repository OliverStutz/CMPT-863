// Participant API client ("02 - Architecture.md" sections 5.3, 8.1, 8.3).
//
// - Every call sends `X-Study-Client: 1` and the session cookie (credentials: "include").
// - The client never counts on its own: `clientSeq` is always taken from the last response's
//   `expectedClientSeq` (success, error, POST sessions and GET /next alike).
// - All calls run through one queue, so events leave in order and every event uses the number the
//   previous response gave.
// - A network failure or a retryable 5xx is retried with the identical body (same clientSeq and
//   clientTs); the server treats that as a retry and returns the stored result.

import type {
  ApiErrorBody,
  CreateSessionResponse,
  EventPayloads,
  EventResponse,
  EventType,
  NextResponse,
  Step,
} from "@study/shared";

export class ApiError extends Error {
  readonly status: number;
  readonly code: string;
  readonly retryable: boolean;
  readonly retryInMs?: number;
  readonly next?: Step;
  readonly expectedClientSeq?: number;
  constructor(status: number, body: ApiErrorBody["error"] | undefined) {
    super(body?.message ?? `HTTP ${status}`);
    this.status = status;
    this.code = body?.code ?? (status === 0 ? "network" : "unknown");
    this.retryable = body?.retryable ?? false;
    this.retryInMs = body?.retryInMs;
    this.next = body?.next;
    this.expectedClientSeq = body?.expectedClientSeq;
  }
}

export type FetchLike = (input: string, init?: RequestInit) => Promise<Response>;

export interface ApiClientOptions {
  fetch?: FetchLike;
  /** How often a network failure or retryable 5xx is retried with the same body. */
  maxRetries?: number;
  retryDelayMs?: number;
  /** Called with every step descriptor the server returns, including the `next` of an error. */
  onStep?: (step: Step) => void;
}

const sleep = (ms: number) => new Promise<void>((r) => setTimeout(r, ms));

export class ApiClient {
  sessionId: string | null = null;
  /** The clientSeq the next event must carry, from the last response. */
  expectedClientSeq = 1;
  private readonly fetchImpl: FetchLike;
  private readonly maxRetries: number;
  private readonly retryDelayMs: number;
  private queue: Promise<unknown> = Promise.resolve();
  onStep?: (step: Step) => void;

  constructor(opts: ApiClientOptions = {}) {
    this.fetchImpl = opts.fetch ?? ((i, init) => fetch(i, init));
    this.maxRetries = opts.maxRetries ?? 3;
    this.retryDelayMs = opts.retryDelayMs ?? 800;
    this.onStep = opts.onStep;
  }

  /** Runs `fn` after every earlier call has settled. */
  private enqueue<T>(fn: () => Promise<T>): Promise<T> {
    const run = this.queue.then(fn, fn);
    this.queue = run.catch(() => undefined);
    return run;
  }

  private async request<T>(method: "GET" | "POST", path: string, body?: unknown): Promise<T> {
    const init: RequestInit = {
      method,
      credentials: "include",
      headers: {
        "X-Study-Client": "1",
        Accept: "application/json",
        ...(body !== undefined ? { "Content-Type": "application/json" } : {}),
      },
      body: body !== undefined ? JSON.stringify(body) : undefined,
    };
    let attempt = 0;
    for (;;) {
      let res: Response;
      try {
        res = await this.fetchImpl(path, init);
      } catch {
        if (attempt++ < this.maxRetries) {
          await sleep(this.retryDelayMs * attempt);
          continue;
        }
        throw new ApiError(0, { code: "network", message: "The connection failed.", retryable: true });
      }
      let json: unknown = undefined;
      try {
        json = await res.json();
      } catch {
        json = undefined;
      }
      if (res.ok) {
        this.absorb(json as Partial<NextResponse>);
        return json as T;
      }
      const err = (json as ApiErrorBody | undefined)?.error;
      if (err) this.absorb(err);
      if (res.status >= 500 && (err?.retryable ?? true) && attempt++ < this.maxRetries) {
        await sleep(this.retryDelayMs * attempt);
        continue;
      }
      throw new ApiError(res.status, err);
    }
  }

  /** Takes expectedClientSeq and the step descriptor from any response or error body. */
  private absorb(body: { expectedClientSeq?: number; next?: Step } | undefined) {
    if (!body) return;
    if (typeof body.expectedClientSeq === "number") this.expectedClientSeq = body.expectedClientSeq;
    if (body.next && this.onStep) this.onStep(body.next);
  }

  createSession(code: string): Promise<CreateSessionResponse> {
    return this.enqueue(async () => {
      const res = await this.request<CreateSessionResponse>(
        "POST",
        `/api/v1/runs/${encodeURIComponent(code)}/sessions`,
        {
          screen: { w: window.screen?.width ?? window.innerWidth, h: window.screen?.height ?? window.innerHeight },
          tz: Intl.DateTimeFormat().resolvedOptions().timeZone ?? "",
        },
      );
      this.sessionId = res.sessionId;
      return res;
    });
  }

  /** Loads a JSON document (the public scene pack) with the session cookie. */
  async getJson<T>(url: string): Promise<T> {
    const res = await this.fetchImpl(url, { credentials: "include", headers: { Accept: "application/json" } });
    if (!res.ok) throw new ApiError(res.status, undefined);
    return (await res.json()) as T;
  }

  next(): Promise<NextResponse> {
    return this.enqueue(() => this.request<NextResponse>("GET", `/api/v1/sessions/${this.sid()}/next`));
  }

  /** Sends one event with the clientSeq of the last response. */
  send<T extends EventType>(type: T, payload: EventPayloads[T]): Promise<EventResponse> {
    return this.enqueue(() =>
      this.request<EventResponse>("POST", `/api/v1/sessions/${this.sid()}/events`, {
        clientSeq: this.expectedClientSeq,
        type,
        clientTs: new Date().toISOString(),
        payload,
      }),
    );
  }

  /**
   * Sends an event; on `too_early` waits `retryInMs` and sends it again with the next clientSeq
   * (section 5.3). Other errors are thrown.
   */
  async sendUntilOnTime<T extends EventType>(
    type: T,
    payload: EventPayloads[T],
    maxTries = 5,
  ): Promise<EventResponse> {
    for (let i = 0; ; i++) {
      try {
        return await this.send(type, payload);
      } catch (e) {
        if (e instanceof ApiError && e.code === "too_early" && i < maxTries) {
          await sleep(Math.max(0, e.retryInMs ?? 500) + 50);
          continue;
        }
        throw e;
      }
    }
  }

  private sid(): string {
    if (!this.sessionId) throw new Error("No session yet");
    return encodeURIComponent(this.sessionId);
  }
}
