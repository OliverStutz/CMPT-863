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

export function audioMode(
  c: Pick<ConfigSummary, "manifestPath" | "estimatedAudio"> | undefined,
): AudioMode | undefined {
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

// ---- Run analysis, "06 - Analysis Dashboard Plan.md" section 6 (binding contract). ----

export type AnalysisInclude = "finished" | "consented";
export type AnalysisGroupFilter = "both" | Group;
// opposite: p < .05 against the hypothesis; does_not_support: not significant; too_few below 8 contributing per group;
// n_a where no direction is stated or the value is undefined.
export type Reading = "supports" | "opposite" | "does_not_support" | "too_few" | "n_a";
export type RatingItemId = "trust" | "competence" | "benevolence" | "integrity";
export type RatingPhase = "baseline" | "afterSlide" | "afterAnswer";

export interface Summary {
  n: number;
  mean: number | null;
  sd: number | null;
  ci95: [number, number] | null;
  median: number | null;
}
export interface OneSample {
  n: number;
  mean: number | null;
  ci95: [number, number] | null;
  t: number | null;
  df: number | null;
  p: number | null;
  reading: Reading;
}
export interface Comparison {
  checker: Summary;
  neutral: Summary;
  diff: number | null;
  diffCi95: [number, number] | null;
  welchT: number | null;
  df: number | null;
  p: number | null;
  hedgesG: number | null;
  reading: Reading;
}
export interface Paired {
  n: number;
  reasoned: Summary;
  denial: Summary;
  meanDiff: number | null;
  ci95: [number, number] | null;
  t: number | null;
  df: number | null;
  p: number | null;
  reading: Reading;
}
export interface Share {
  slides: number;
  withQuestion: number;
  share: number | null;
  perParticipant: Summary;
}
export interface TwoProportion {
  p1: number | null;
  p2: number | null;
  diff: number | null;
  z: number | null;
  p: number | null;
  reading: Reading;
}

export type ByGroup<T> = Record<Group, T>;

export interface RatingValue {
  trust: number | null;
  competence: number | null;
  benevolence: number | null;
  integrity: number | null;
  status: string;
}

export interface AnalysisParticipant {
  sessionId: string;
  group: Group;
  denialSlide: string;
  status: string;
  finishedAt: string | null;
  durationSec: number | null;
  ratings: Record<string, RatingValue | null>; // by point
  actions: Record<
    string,
    { actionId: string | null; timedOut: boolean; questionBy: string; answerType: "reason" | "denial" }
  >; // by flagged slide id
  budget: { optionId: string | null; confidence: number | null; status: string } | null;
  background: Record<string, string | number | null> | null; // item id -> value or "preferNotToSay"
  flags: string[];
}

export interface BackgroundItemConfig {
  id: string;
  label: string;
  kind: "single" | "scale" | "country";
  options?: { id: string; label: string }[];
  scale?: { min: number; max: number; pointLabels?: Record<string, string> };
}

export interface LastTrustTercile {
  tercile: 1 | 2 | 3;
  from: number;
  to: number;
  approve: number;
  reject: number;
  missing: number;
}

export interface ModeratorRow {
  h1: ByGroup<Summary>;
  h2: ByGroup<Summary>;
}

/** GET /api/admin/runs/{id}/analysis. */
export interface RunAnalysis {
  run: {
    id: string;
    code: string;
    name: string;
    source: string;
    configVersion: string;
    status: string;
    createdAt: string;
    testOverrides: boolean;
  };
  generatedAt: string;
  filters: { include: AnalysisInclude; exclude: string[]; group: AnalysisGroupFilter };
  config: {
    groups: ["checker", "neutral"];
    slides: { id: string; n: number; title: string; flagged: boolean }[];
    flaggedSlides: string[];
    ratingPoints: { point: string; slideId: string; phase: RatingPhase; label: string }[]; // plan order
    items: { id: RatingItemId; label: string; min: number; max: number }[];
    actions: { checker: { id: string; label: string }[]; neutral: { id: string; label: string }[] };
    background: BackgroundItemConfig[];
    budgetOptions: { id: string; label: string }[];
  };
  counts: {
    byStatus: {
      opened: number;
      started: number;
      ongoing: number;
      finished: number;
      abandoned: number;
      blocked: number;
    };
    total: number; // every session of the run
    consented: number;
    analysed: number; // after include and exclude
    excluded: { forced: number; byFlag: Record<string, number>; notIncluded: number };
    byGroup: ByGroup<{ consented: number; finished: number; analysed: number }>;
    durationsSec: {
      n: number;
      median: number | null;
      p25: number | null;
      p75: number | null;
      histogram: { fromSec: number; toSec: number; n: number }[];
    };
  };
  participants: AnalysisParticipant[];
  ratingStats: Record<string /* item */, Record<string /* point */, ByGroup<Summary>>>;
  hypotheses: {
    baseline: Comparison; // T0 checker vs neutral
    slide1: Comparison; // T1 - T0
    overall: Comparison; // last rating - T0, research question, no direction
    h1: {
      pooled: Comparison;
      checkerBelowZero: OneSample;
      perSlide: Record<string, Comparison>; // reading "n_a"
      levelAtFlag: Record<string, Comparison>; // reading "n_a"
    };
    h2: {
      pooled: Comparison; // change T{n}b - T{n}a between groups (same data as h4b.afterReason, H4b1)
      checkerAboveZero: OneSample; // H2(a)
      level: Comparison; // H2(b): level T{n}b over reasoned slides; carries the H2 reading
      perSlide: Record<string, Comparison>; // change per slide, reading "n_a"
      levelPerSlide: Record<string, Comparison>; // level per slide, reading "n_a"
      aboveBaseline: ByGroup<OneSample>; // exploratory: T{n}b - T0 within group
      exploratoryNet: Comparison; // secondary: T{n}b - pre-flag rating, no direction
    };
    h3: { perGroup: ByGroup<Paired>; both: Paired };
    h4a: { checker: Share; neutral: Share; slidesTest: TwoProportion; participantsTest: Comparison }; // slidesTest reading "n_a"
    h4b: { afterReason: Comparison; afterDenial: Comparison };
    budget: {
      byGroup: ByGroup<{ approve: number; reject: number; missing: number; confidence: Summary }>;
      approveTest: TwoProportion;
      /** Terciles of the last trust rating (ties to the lower tercile); optional until every server sends it. */
      byLastTrustTercile?: LastTrustTercile[];
    };
    moderators: {
      cyberKnowledge: ({ level: number } & ModeratorRow)[];
      aiUse: ({ optionId: string } & ModeratorRow)[];
      spearman: {
        cyberKnowledgeH1: ByGroup<number | null>;
        cyberKnowledgeH2: ByGroup<number | null>;
        aiUseH1: ByGroup<number | null>;
        aiUseH2: ByGroup<number | null>;
      };
    };
  };
  actions: Record<string /* slide */, ByGroup<Record<string /* actionId or "timeout" */, number>>>;
  background: Record<string /* item */, { id: string; label: string; n: number; byGroup: ByGroup<number> }[]>; // every item, country too; with "preferNotToSay" and "missing"
  countries: { code: string; n: number; byGroup: ByGroup<number> }[]; // real ISO codes only, sorted by n desc
  quality: {
    flags: Record<string, number>;
    audioCheck: { correct: number; wrong: number; missing: number };
    ratingTimeouts: Record<string /* point */, number>;
    reloads: { sessions: number; total: number };
  };
}

export interface AnalysisQuery {
  include?: AnalysisInclude;
  exclude?: string[];
  group?: AnalysisGroupFilter;
}

/** DELETE /api/admin/sessions/{id} body. */
export interface DeleteSessionRequest {
  confirm: string;
  reason: string;
}
