// Participant API DTOs and the step descriptor, from "02 - Architecture.md" sections 5.3, 8.1 and 8.3.

import type { Hex, Id, MarkingShape, RatingItemId } from "./config";

export type StepScene =
  | "deviceGate"
  | "consent"
  | "audioCheck"
  | "example"
  | "preStart"
  | "opening"
  | "slide"
  | "postView"
  | "rating"
  | "action"
  | "question"
  | "answer"
  | "budget"
  | "background"
  | "debrief"
  | "done";

/** Steps that play exactly one clip and use scene_started / audio_ended (section 5.3). */
export const AUDIO_STEP_SCENES: readonly StepScene[] = ["example", "opening", "slide", "question", "answer"];

/** A Marking without appearAtMark, anchorText and vlmRun, with appearAtMs resolved and color filled (8.1). */
export interface ResolvedMarking {
  id: Id;
  targetIds: Id[];
  bbox?: [x: number, y: number, w: number, h: number];
  appearAtMs: number;
  color: Hex;
  shape: MarkingShape;
  label: string;
  detail: string;
}

/**
 * The panel texts travel with the overlay (server DECISION): labels.panelTitle, labels.panelIdle and the
 * noIssues colour are checker-only, so the public scene pack no longer carries them.
 */
export interface OverlayPanelTexts {
  title?: string;
  idleText?: string;
  noIssuesColor?: Hex;
}
export type StepOverlay =
  | ({ status: "noIssues"; text: string; appearAtMs: number } & OverlayPanelTexts)
  | ({ status: "flagged"; markings: ResolvedMarking[] } & OverlayPanelTexts);

export interface StepOption {
  id: string;
  label: string;
  // DECISION: "03 - Study Content v1.3.md" section 7 shows the exact spoken question under the two
  // speaking options, but section 8.1 gives options only { id, label }. The client renders `detail`
  // in small type when the server sends it and nothing otherwise.
  detail?: string;
}

export interface Step {
  stepId: string;
  scene: StepScene;
  sceneRef?: string;
  audioUrl?: string;
  replayable?: boolean;
  restart?: boolean;
  overlay?: StepOverlay;
  options?: StepOption[];
  point?: string;
  transcript?: string;
  preStartText?: string[];
  debrief?: {
    paragraphs: string[];
    doNotDiscuss: string;
    sessionCode: string;
    // DECISION: section 8.1 gives no label for the session code and the debrief scene is not in the
    // public pack, so the client uses this when sent and the content's wording otherwise.
    sessionCodeLabel?: string;
  };
  opensInMs?: number;
  closesInMs?: number;
  prefetch?: string[];
  resumed?: boolean;
  // DECISION: the action prompt and note differ per group (config `actions.checker` / `actions.neutral`)
  // and are not in the public pack, and section 8.1 has no field for them. The client renders these
  // two optional fields when the server sends them, and no prompt otherwise. It never picks a text itself.
  prompt?: string;
  note?: string;
  /** Server DECISION: a session the device gate blocked gets scene "done" with blocked: true. */
  blocked?: boolean;
}

export interface Progress {
  index: number;
  total: number;
}

// ---------- requests and responses ----------
export interface CreateSessionRequest {
  screen: { w: number; h: number };
  tz: string;
  /** Test links only (run with testOverrides): fixes the group and the denial slide; ignored otherwise. */
  test?: { group?: string; denialSlide?: string };
}
export interface CreateSessionResponse {
  sessionId: string;
  configVersion: string;
  scenePackUrl: string;
  assetBaseUrl: string;
  expectedClientSeq: number;
  next: Step;
}
export interface NextResponse {
  next: Step;
  expectedClientSeq: number;
  progress?: Progress;
}
export interface EventRequest<T extends EventType = EventType> {
  clientSeq: number;
  type: T;
  clientTs: string;
  payload: EventPayloads[T];
}
export interface EventResponse {
  accepted: true;
  eventNo: number;
  expectedClientSeq: number;
  next: Step;
  /** audio_problem only: when the "Continue" button may be used (section 5.3). */
  continueInMs?: number;
}

export type ErrorCode =
  | "bad_request"
  | "bad_payload"
  | "no_session"
  | "bad_token"
  | "unknown_run"
  | "run_closed"
  | "wrong_state"
  | "seq_gap"
  | "duplicate_seq_mismatch"
  | "late"
  | "too_early"
  | "out_of_range"
  | "rate_limited"
  | "internal"
  | "unavailable";

export interface ApiErrorBody {
  error: {
    code: ErrorCode | string;
    message: string;
    retryable?: boolean;
    retryInMs?: number;
    expectedClientSeq?: number;
    next?: Step;
  };
}

// ---------- events (section 5.3) ----------
export type RatingValues = Record<RatingItemId, number | null>;
/** Background answers: item id -> option id, scale number, country code or "preferNotToSay"; null if untouched. */
export type BackgroundAnswers = Record<string, string | number | null>;
export const PREFER_NOT_TO_SAY = "preferNotToSay";

export interface EventPayloads {
  device_checked: { widthPx: number; heightPx: number; touch: boolean };
  consent_accepted: { consentVersion: string; checkbox: true };
  audio_check_played: Record<string, never>;
  audio_check_answered: { choiceId: string };
  scene_started: { sceneId: string };
  marking_shown: { sceneId: string; markingId: string };
  audio_ended: { sceneId: string; afterProblem?: true };
  audio_problem: { sceneId: string; kind: "error" | "stalled" | "noEnded"; atMs: number };
  setup_done: Record<string, never>;
  rating_submitted: { point: string } & RatingValues;
  rating_timeout: { point: string; partial: RatingValues };
  action_chosen: { slideId: string; actionId: string };
  action_timeout: { slideId: string };
  budget_submitted: { optionId: string | null; confidence: number | null };
  budget_timeout: { partial: { optionId: string | null; confidence: number | null } };
  background_submitted: { items: BackgroundAnswers; hp: string };
  background_timeout: { partial: BackgroundAnswers; hp: string };
  finished: Record<string, never>;
  visibility: { hidden: boolean };
}
export type EventType = keyof EventPayloads;
