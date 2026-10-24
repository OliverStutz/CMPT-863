import type { Group, ParticipantRow, SessionStatus } from "../api/types";

export interface Filters {
  runId: string;
  status: SessionStatus | "";
  group: Group | "";
  flaggedOnly: boolean;
}

export const NO_FILTERS: Filters = { runId: "", status: "", group: "", flaggedOnly: false };

// DECISION: polled rows (updatedSince) are merged by sessionId and the filters are re-applied in the browser, so a
// row whose status left the filter drops out; a filter change starts a full reload (cursor followed to the end).
/** Client-side check of the same filters the server applies, so a polled row that left the filter disappears. */
export function matches(row: ParticipantRow, f: Filters, runCode: string | undefined): boolean {
  if (f.runId && runCode !== undefined && row.runCode !== runCode) return false;
  if (f.status && row.status !== f.status) return false;
  if (f.group && row.group !== f.group) return false;
  if (f.flaggedOnly && (!row.flags || row.flags.length === 0)) return false;
  return true;
}

/** Merges rows changed since the last poll into the list, newest openedAt first (the 7.1 index order). */
export function mergeRows(
  current: ParticipantRow[],
  changed: ParticipantRow[],
  f: Filters,
  runCode: string | undefined,
): ParticipantRow[] {
  const byId = new Map(current.map((r) => [r.sessionId, r]));
  for (const r of changed) byId.set(r.sessionId, r);
  return [...byId.values()]
    .filter((r) => matches(r, f, runCode))
    .sort((a, b) => (a.openedAt < b.openedAt ? 1 : a.openedAt > b.openedAt ? -1 : 0));
}
