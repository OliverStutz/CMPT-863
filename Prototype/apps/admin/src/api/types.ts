// DTOs of the admin API, copied from "02 - Architecture.md" sections 7.1 and 8.2.
// Defined locally on purpose: this app does not depend on packages/shared.

export type SessionStatus = "Opened" | "Started" | "Ongoing" | "Finished" | "Abandoned" | "Blocked";
export const SESSION_STATUSES: SessionStatus[] = ["Opened", "Started", "Ongoing", "Finished", "Abandoned", "Blocked"];

export type Group = "checker" | "neutral";
export type RunStatus = "open" | "closed";
export type ConfigStatus = "draft" | "validated" | "locked";

export interface RunCounts {
  opened: number;
  started: number;
  ongoing: number;
  finished: number;
  abandoned: number;
  blocked: number;
}

/** GET /api/admin/runs item. */
export interface Run {
  id: string;
  code: string;
  name: string;
  source: string;
  configVersion: string;
  status: RunStatus;
  createdAt: string;
  /** Test run: links may carry ?group= and ?denial=; such sessions are flagged forced_allocation. */
  testOverrides?: boolean;
  counts: RunCounts;
  // DECISION: section 8.2 names `finishedByGroup` without a shape; read as { checker, neutral } counts,
  // matching counters/allocation.groups in 7.1.
  finishedByGroup?: Partial<Record<Group, number>>;
}

/** POST /api/admin/runs request and response. */
export interface CreateRunRequest {
  name: string;
  source: string;
  configVersion: string;
  testOverrides?: boolean;
}
export interface CreateRunResponse {
  id: string;
  code: string;
  url: string;
}

/** PATCH /api/admin/runs/{id} request. */
export interface PatchRunRequest {
  name?: string;
  source?: string;
  status?: RunStatus;
  testOverrides?: boolean;
}

export interface AudioCheckResult {
  choiceId?: string;
  correct: boolean | null;
  plays: number;
}

/** GET /api/admin/participants item. */
export interface ParticipantRow {
  sessionId: string;
  runCode: string;
  source: string;
  group: Group | null;
  openedAt: string;
  consentAt: string | null;
  finishedAt: string | null;
  progress: { index: number; total: number };
  status: SessionStatus;
  lastStep: string | null;
  lastEventAt: string | null;
  audioCheck: AudioCheckResult | null;
  reloadCount: number;
  flags: string[];
}

export interface ParticipantsPage {
  items: ParticipantRow[];
  cursor: string | null;
  serverTime: string;
}

export interface ParticipantsQuery {
  runId?: string;
  status?: SessionStatus;
  group?: Group;
  flagged?: boolean;
  updatedSince?: string;
  limit?: number;
  cursor?: string;
}

/** sessions/{id}/events/{eventNo}, section 7.1. */
export interface SessionEvent {
  eventNo: number;
  clientSeq: number | null;
  type: string;
  source: "client" | "server";
  serverTs: string;
  clientTs: string | null;
  payload: unknown;
  stateBefore: string | null;
  stateAfter: string | null;
  valid: boolean;
  reason: string | null;
}

/** sessions/{id}, section 7.1, as far as the admin shows it. */
export interface SessionDoc {
  sessionId?: string;
  id?: string;
  runId?: string;
  runCode?: string;
  source?: string;
  configVersion?: string;
  group?: Group | null;
  denialSlide?: string | null;
  status?: SessionStatus;
  currentState?: string;
  progressIndex?: number;
  progressTotal?: number;
  openedAt?: string;
  consentAt?: string | null;
  startedAt?: string | null;
  ongoingAt?: string | null;
  finishedAt?: string | null;
  lastEventAt?: string | null;
  audioCheck?: AudioCheckResult | null;
  reloadCount?: number;
  uaFamily?: string;
  flags?: string[];
  [key: string]: unknown;
}

// DECISION: section 8.2 says the detail response is "session plus full event timeline, ratings, answers"
// without field names. Read as { session, events, ratings, answers }; if the server instead returns the
// session fields at the top level, normalizeDetail() in api/client.ts copes with that too.
export interface ParticipantDetail {
  session: SessionDoc;
  events: SessionEvent[];
  ratings: Record<string, unknown> | unknown[];
  answers: Record<string, unknown> | unknown[];
}

// DECISION: section 8.2 says the config list carries "status and lint report" without field names.
// Read as the configs/{version} document of section 7.1 (version, sha256, status, uploadedAt, lintReport).
export interface ConfigSummary {
  version: string;
  sha256?: string;
  status: ConfigStatus;
  uploadedAt?: string;
  lintReport?: unknown;
  /** Path of the audio manifest the server resolved; empty when the config plays silent placeholder clips. */
  manifestPath?: string;
  /** True when at least one clip has no manifest entry (estimated duration, silent clip). */
  estimatedAudio?: boolean;
}

/**
 * How a config's audio is served. A config without an audio manifest (no TTS run for its version) plays silent
 * placeholder clips of estimated length: the study flow works, but nothing is heard. The admin view says so
 * wherever a version is shown, so a silent test run is never mistaken for a broken browser.
 */
export type AudioMode = "spoken" | "partly-silent" | "silent";

export function audioMode(c: Pick<ConfigSummary, "manifestPath" | "estimatedAudio"> | undefined): AudioMode | undefined {
  if (!c) return undefined;
  if (!c.manifestPath) return "silent";
  return c.estimatedAudio ? "partly-silent" : "spoken";
}

export const AUDIO_MODE_LABEL: Record<AudioMode, string> = {
  spoken: "spoken audio",
  "partly-silent": "some clips silent",
  silent: "no audio: silent clips",
};

// DECISION: GET /api/admin/configs/{version} ("full config", 8.2) is read as the bare StudyConfig object.
/** The parts of StudyConfig (section 4.1) the Configs tab reads from GET /api/admin/configs/{version}. */
export interface StudyConfigLite {
  version: string;
  scenes: Array<{ kind: string; id: string; checker?: { status: "noIssues" | "flagged" } }>;
}

/** Section 8.3. */
export interface ApiErrorBody {
  error: { code: string; message: string; retryable?: boolean; retryInMs?: number };
}
