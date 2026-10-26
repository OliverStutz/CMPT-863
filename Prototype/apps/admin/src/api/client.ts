import type {
  AnalysisQuery,
  ApiErrorBody,
  ConfigSummary,
  CreateRunRequest,
  CreateRunResponse,
  DeleteSessionRequest,
  ParticipantDetail,
  ParticipantsPage,
  ParticipantsQuery,
  PatchRunRequest,
  Run,
  RunAnalysis,
  SessionDoc,
  StudyConfigLite,
} from "./types";

const BASE = "/api/admin";

export class ApiError extends Error {
  constructor(
    public status: number,
    public code: string,
    message: string,
  ) {
    super(message);
  }
}

// ---- 401 handling: a request that gets 401 waits for a successful login, then is sent again. ----
type AuthListener = (loginRequired: boolean) => void;
let authListener: AuthListener = () => {};
let waiters: Array<() => void> = [];

export function onAuthChange(fn: AuthListener): void {
  authListener = fn;
}

function waitForLogin(): Promise<void> {
  authListener(true);
  return new Promise((resolve) => waiters.push(resolve));
}

// ---- gate handling: the access-key gate answers a plain-text 404 (the API's own errors are JSON).
// It means this browser's gate cookie is gone (30 days without opening the page, or a rotated key):
// the admin link with the access key must be opened again. The app shows that instead of a dead page. ----
let gateListener: () => void = () => {};

export function onGateLost(fn: () => void): void {
  gateListener = fn;
}

function isGate404(res: Response): boolean {
  return res.status === 404 && !(res.headers.get("Content-Type") ?? "").includes("application/json");
}

function loginSucceeded(): void {
  authListener(false);
  const w = waiters;
  waiters = [];
  w.forEach((resolve) => resolve());
}

async function errorFrom(res: Response): Promise<ApiError> {
  try {
    const body = (await res.json()) as ApiErrorBody;
    return new ApiError(res.status, body.error?.code ?? "unknown", body.error?.message ?? res.statusText);
  } catch {
    return new ApiError(res.status, "unknown", res.statusText || `HTTP ${res.status}`);
  }
}

/** Sends one admin request; on 401 shows the login screen and retries after login. */
export async function raw(path: string, init: RequestInit = {}): Promise<Response> {
  for (;;) {
    const headers = new Headers(init.headers);
    headers.set("X-Study-Admin", "1"); // CSRF guard, section 8.2
    if (init.body !== undefined) headers.set("Content-Type", "application/json");
    const res = await fetch(BASE + path, { ...init, headers, credentials: "same-origin" });
    if (res.status === 401) {
      await waitForLogin();
      continue;
    }
    if (isGate404(res)) {
      gateListener();
      throw new ApiError(404, "gate_lost", "Open the admin link with the access key again.");
    }
    if (!res.ok) throw await errorFrom(res);
    return res;
  }
}

async function json<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await raw(path, init);
  if (res.status === 204) return undefined as T;
  return (await res.json()) as T;
}

// ---- endpoints (section 8.2) ----

export async function login(password: string): Promise<void> {
  const res = await fetch(`${BASE}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json", "X-Study-Admin": "1" },
    credentials: "same-origin",
    body: JSON.stringify({ password }),
  });
  if (isGate404(res)) {
    gateListener();
    throw new ApiError(404, "gate_lost", "Open the admin link with the access key again.");
  }
  if (!res.ok) throw await errorFrom(res);
  loginSucceeded();
}

export async function logout(): Promise<void> {
  await fetch(`${BASE}/logout`, { method: "POST", headers: { "X-Study-Admin": "1" }, credentials: "same-origin" });
  authListener(true);
}

export const listRuns = () => json<Run[]>("/runs");

export const createRun = (req: CreateRunRequest) =>
  json<CreateRunResponse>("/runs", { method: "POST", body: JSON.stringify(req) });

export const patchRun = (id: string, req: PatchRunRequest) =>
  json<Run>(`/runs/${encodeURIComponent(id)}`, { method: "PATCH", body: JSON.stringify(req) });

export function listParticipants(q: ParticipantsQuery): Promise<ParticipantsPage> {
  const p = new URLSearchParams();
  if (q.runId) p.set("runId", q.runId);
  if (q.status) p.set("status", q.status);
  if (q.group) p.set("group", q.group);
  if (q.flagged) p.set("flagged", "true");
  if (q.updatedSince) p.set("updatedSince", q.updatedSince);
  if (q.limit) p.set("limit", String(q.limit));
  if (q.cursor) p.set("cursor", q.cursor);
  const qs = p.toString();
  return json<ParticipantsPage>(`/participants${qs ? `?${qs}` : ""}`);
}

export async function getParticipant(id: string): Promise<ParticipantDetail> {
  return normalizeDetail(await json<Record<string, unknown>>(`/participants/${encodeURIComponent(id)}`));
}

export function normalizeDetail(body: Record<string, unknown>): ParticipantDetail {
  const { session, events, ratings, answers, ...rest } = body;
  return {
    session: (session as SessionDoc | undefined) ?? (rest as SessionDoc),
    events: Array.isArray(events) ? events : [],
    ratings: (ratings as ParticipantDetail["ratings"]) ?? {},
    answers: (answers as ParticipantDetail["answers"]) ?? {},
  };
}

export const listConfigs = () => json<ConfigSummary[]>("/configs");

export const getConfig = (version: string) => json<StudyConfigLite>(`/configs/${encodeURIComponent(version)}`);

export interface ExportOptions {
  format: "csv" | "json";
  runId?: string;
  includeInvalid: boolean;
  includeIp: boolean;
}

export function exportPath(o: ExportOptions): string {
  const p = new URLSearchParams({ format: o.format });
  if (o.runId) p.set("runId", o.runId);
  if (o.includeInvalid) p.set("includeInvalid", "true");
  if (o.includeIp) p.set("includeIp", "true");
  return `/export?${p.toString()}`;
}

// DECISION: every admin call needs the X-Study-Admin header (8.2), which a plain <a href> download cannot send,
// so the export is fetched with the header and saved from a Blob. The file name comes from Content-Disposition
// when the server sets one, else study-export-<run or all>-<timestamp>.<format>.
export async function downloadExport(o: ExportOptions, runLabel: string): Promise<void> {
  const res = await raw(exportPath(o));
  const blob = await res.blob();
  const cd = res.headers.get("Content-Disposition") ?? "";
  const m = /filename\*?=(?:UTF-8'')?"?([^";]+)"?/i.exec(cd);
  const stamp = new Date().toISOString().replace(/[:.]/g, "-");
  const name = m ? decodeURIComponent(m[1]) : `study-export-${runLabel}-${stamp}.${o.format}`;
  const url = URL.createObjectURL(blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = name;
  document.body.appendChild(a);
  a.click();
  a.remove();
  setTimeout(() => URL.revokeObjectURL(url), 1000);
}

// ---- run analysis and deletion ("06 - Analysis Dashboard Plan.md" section 6) ----

/** Query string of the analysis call; a missing filter is left out so the server default applies. */
export function analysisQuery(q: AnalysisQuery): string {
  const p = new URLSearchParams();
  if (q.include) p.set("include", q.include);
  // DECISION: an empty exclude list is sent as `exclude=` (exclude nothing beyond forced_allocation), which the
  // contract distinguishes from a missing parameter (the server default). forced_allocation is never sent, the
  // server always adds it.
  if (q.exclude) p.set("exclude", q.exclude.filter((f) => f !== "forced_allocation").join(","));
  if (q.group) p.set("group", q.group);
  const qs = p.toString();
  return qs ? `?${qs}` : "";
}

export const getRunAnalysis = (id: string, q: AnalysisQuery = {}) =>
  json<RunAnalysis>(`/runs/${encodeURIComponent(id)}/analysis${analysisQuery(q)}`);

/** 204 on success; 404 unknown session; 400 when confirm does not match the id. */
export const deleteSession = (id: string, req: DeleteSessionRequest) =>
  json<void>(`/sessions/${encodeURIComponent(id)}`, { method: "DELETE", body: JSON.stringify(req) });
