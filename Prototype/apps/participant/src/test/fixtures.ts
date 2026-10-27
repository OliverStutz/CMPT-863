// Test fixtures: a public scene pack built from the real v1.3 config, and a scripted mock server.
import { readFileSync } from "node:fs";
import { fileURLToPath } from "node:url";
import { dirname, resolve } from "node:path";
import type { PublicScenePack, Step } from "@study/shared";

function stripNotes(o: unknown): unknown {
  if (Array.isArray(o)) return o.map(stripNotes);
  if (o && typeof o === "object") {
    return Object.fromEntries(
      Object.entries(o as Record<string, unknown>)
        .filter(([k]) => !k.startsWith("_"))
        .map(([k, v]) => [k, stripNotes(v)]),
    );
  }
  return o;
}

/** The public pack as section 4.3 describes it, derived from content/study-config.v1.3.json. */
export function packFromConfig(): PublicScenePack {
  const here = dirname(fileURLToPath(import.meta.url));
  const cfg = stripNotes(
    JSON.parse(readFileSync(resolve(here, "../../../../content/study-config.v1.3.json"), "utf8")),
  ) as Record<string, any>;
  const scenes = (cfg.scenes as any[])
    .filter((s) => s.kind !== "debrief")
    .map((s) => {
      if (s.kind === "slide") return { kind: s.kind, id: s.id, template: s.template, layout: s.layout };
      if (s.kind === "opening") return { kind: s.kind, id: s.id, layout: s.layout, profileCard: s.profileCard };
      return s;
    });
  return {
    version: cfg.version,
    options: cfg.options,
    // As the server builds it: the checker panel's labels and the colours are not in the pack.
    labels: Object.fromEntries(
      Object.entries(cfg.labels).filter(([k]) => !["panelTitle", "panelIdle", "noIssues"].includes(k)),
    ),
    consent: cfg.consent,
    setup: {
      instruction: cfg.setup.instruction,
      audioCheck: { prompt: cfg.setup.audioCheck.prompt, options: cfg.setup.audioCheck.options },
      exampleScene: { layout: cfg.setup.exampleScene.layout },
      preStartText: { both: cfg.setup.preStartText.both },
    },
    rating: cfg.rating,
    scenes,
  } as PublicScenePack;
}

export interface SentEvent {
  clientSeq: number;
  type: string;
  payload: any;
}

export type Reply = { status?: number; body: unknown };

/**
 * A fake fetch. `onEvent` decides each event's reply; GET /next returns `current()`; the scene pack and
 * session creation are fixed. Every reply carries expectedClientSeq = last accepted + 1 unless set.
 */
export function mockServer(opts: {
  pack: PublicScenePack;
  first: Step;
  onEvent: (ev: SentEvent, ctx: { seq: number }) => Reply;
  /** Called on GET /next; a returned step becomes the current one (the server moved on, e.g. after a view). */
  onNext?: () => Step | undefined;
}) {
  const events: SentEvent[] = [];
  const headers: Record<string, string>[] = [];
  let current = opts.first;
  let seq = 1;
  const json = (status: number, body: unknown) =>
    new Response(JSON.stringify(body), { status, headers: { "Content-Type": "application/json" } });

  const fetchImpl = async (url: string, init?: RequestInit): Promise<Response> => {
    headers.push((init?.headers ?? {}) as Record<string, string>);
    if (url.startsWith("/a/scenes/")) return json(200, opts.pack);
    if (url.includes("/sessions") && url.startsWith("/api/v1/runs/")) {
      return json(201, {
        sessionId: "S1",
        configVersion: opts.pack.version,
        scenePackUrl: "/a/scenes/1.3.0.json",
        assetBaseUrl: "/a",
        expectedClientSeq: seq,
        next: current,
      });
    }
    if (url.endsWith("/next")) {
      const moved = opts.onNext?.();
      if (moved) current = moved;
      return json(200, { next: current, expectedClientSeq: seq, progress: { index: 1, total: 40 } });
    }
    if (url.endsWith("/events")) {
      const body = JSON.parse(String(init?.body));
      const ev: SentEvent = { clientSeq: body.clientSeq, type: body.type, payload: body.payload };
      events.push(ev);
      const r = opts.onEvent(ev, { seq });
      const status = r.status ?? 200;
      const b = r.body as any;
      if (status === 200) {
        seq = b.expectedClientSeq ?? seq + 1;
        b.expectedClientSeq = seq;
        if (b.next) current = b.next;
      } else if (b?.error) {
        if (b.error.expectedClientSeq !== undefined) seq = b.error.expectedClientSeq;
        if (b.error.next) current = b.error.next;
      }
      return json(status, b);
    }
    return new Response("", { status: 404 });
  };
  return { fetch: fetchImpl, events, headers, ok: (next: Step, extra: object = {}) => ({ body: { accepted: true, eventNo: events.length, next, ...extra } }) };
}
