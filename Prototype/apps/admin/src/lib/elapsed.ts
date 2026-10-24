import type { SessionStatus } from "../api/types";

export interface ElapsedInput {
  consentAt: string | null | undefined;
  finishedAt: string | null | undefined;
  status: SessionStatus;
  lastEventAt?: string | null;
}

/**
 * Elapsed time since consent in ms, per section 7.1: (finishedAt or now) - consentAt; null before consent.
 * Finished freezes the timer at finishedAt.
 */
// DECISION: the document's formula ticks for every status without finishedAt, so Abandoned and Blocked
// sessions keep counting (an Abandoned session can still resume, section 5.2). Only Finished freezes.
export function elapsedMs(row: ElapsedInput, nowMs: number): number | null {
  if (!row.consentAt) return null;
  const start = Date.parse(row.consentAt);
  if (Number.isNaN(start)) return null;
  let end: number;
  if (row.finishedAt) end = Date.parse(row.finishedAt);
  // DECISION: a Finished row without finishedAt (should not happen) freezes at lastEventAt, else shows "-".
  else if (row.status === "Finished") {
    if (!row.lastEventAt) return null;
    end = Date.parse(row.lastEventAt);
  }
  else end = nowMs;
  if (Number.isNaN(end)) return null;
  return Math.max(0, end - start);
}

/** True while the timer should tick in the browser between polls. */
export function isTicking(row: ElapsedInput): boolean {
  return !!row.consentAt && !row.finishedAt && row.status !== "Finished";
}

/** "-" before consent; m:ss under an hour, h:mm:ss after. */
export function formatElapsed(ms: number | null): string {
  if (ms === null) return "-";
  const total = Math.floor(ms / 1000);
  const h = Math.floor(total / 3600);
  const m = Math.floor((total % 3600) / 60);
  const s = total % 60;
  const ss = String(s).padStart(2, "0");
  return h > 0 ? `${h}:${String(m).padStart(2, "0")}:${ss}` : `${m}:${ss}`;
}

export function elapsedText(row: ElapsedInput, nowMs: number): string {
  return formatElapsed(elapsedMs(row, nowMs));
}
