import { useCallback, useEffect, useRef, useState } from "react";
import { listParticipants, listRuns } from "../api/client";
import { SESSION_STATUSES, type ParticipantRow, type ParticipantsQuery, type Run } from "../api/types";
import { Chip } from "../lib/Chip";
import { elapsedText } from "../lib/elapsed";
import { flagChip, statusChip } from "../lib/status";
import { localTime, shortId } from "../lib/time";
import { useAsync } from "../lib/useAsync";
import { useNow } from "../lib/useNow";
import { mergeRows, NO_FILTERS, type Filters } from "./merge";
import { ParticipantDrawer } from "./ParticipantDrawer";

const POLL_MS = 5000; // section 8.2
const PAGE_LIMIT = 200;

function baseQuery(f: Filters): ParticipantsQuery {
  return {
    runId: f.runId || undefined,
    status: f.status || undefined,
    group: f.group || undefined,
    flagged: f.flaggedOnly || undefined,
  };
}

export function ParticipantsTab() {
  const runs = useAsync(listRuns);
  const [filters, setFilters] = useState<Filters>(NO_FILTERS);
  const [rows, setRows] = useState<ParticipantRow[] | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [lastPoll, setLastPoll] = useState<string | null>(null);
  const [selected, setSelected] = useState<string | null>(null);
  const [reloadKey, setReloadKey] = useState(0); // bumped after a deletion: the delta poll cannot report removals
  const now = useNow(1000);
  const serverTime = useRef<string | null>(null);

  const runCodeOf = useCallback(
    (runId: string): string | undefined => runs.data?.find((r: Run) => r.id === runId)?.code,
    [runs.data],
  );

  // Full load on mount and on filter change, then "changed since serverTime" every 5 s, merged by sessionId.
  useEffect(() => {
    let alive = true;
    let timer: ReturnType<typeof setTimeout> | undefined;
    serverTime.current = null;
    setRows(null);
    const q = baseQuery(filters);
    const code = filters.runId ? runCodeOf(filters.runId) : undefined;

    async function fullLoad() {
      const all: ParticipantRow[] = [];
      let cursor: string | undefined;
      let st = "";
      do {
        const page = await listParticipants({ ...q, limit: PAGE_LIMIT, cursor });
        all.push(...page.items);
        cursor = page.cursor ?? undefined;
        if (!st) st = page.serverTime; // the first page's time, so nothing changed during paging is missed
      } while (cursor);
      if (!alive) return;
      serverTime.current = st;
      setRows(mergeRows([], all, filters, code));
    }

    async function poll() {
      try {
        if (serverTime.current === null) {
          await fullLoad();
        } else {
          const changed: ParticipantRow[] = [];
          let cursor: string | undefined;
          let st = "";
          do {
            const page = await listParticipants({ ...q, updatedSince: serverTime.current, limit: PAGE_LIMIT, cursor });
            changed.push(...page.items);
            cursor = page.cursor ?? undefined;
            if (!st) st = page.serverTime;
          } while (cursor);
          if (!alive) return;
          serverTime.current = st;
          setRows((cur) => mergeRows(cur ?? [], changed, filters, code));
        }
        if (alive) {
          setError(null);
          setLastPoll(new Date().toISOString());
        }
      } catch (e) {
        if (alive) setError(e instanceof Error ? e.message : String(e));
      } finally {
        if (alive) timer = setTimeout(poll, POLL_MS);
      }
    }
    void poll();
    return () => {
      alive = false;
      if (timer) clearTimeout(timer);
    };
  }, [filters, runCodeOf, reloadKey]);

  const set = <K extends keyof Filters>(k: K, v: Filters[K]) => setFilters((f) => ({ ...f, [k]: v }));

  return (
    <section>
      <div className="filters">
        <label>
          Run
          <select value={filters.runId} onChange={(e) => set("runId", e.target.value)}>
            <option value="">All runs</option>
            {runs.data?.map((r) => (
              <option key={r.id} value={r.id}>
                {r.name} ({r.code})
              </option>
            ))}
          </select>
        </label>
        <label>
          Status
          <select value={filters.status} onChange={(e) => set("status", e.target.value as Filters["status"])}>
            <option value="">All</option>
            {SESSION_STATUSES.map((s) => (
              <option key={s} value={s}>
                {s}
              </option>
            ))}
          </select>
        </label>
        <label>
          Group
          <select value={filters.group} onChange={(e) => set("group", e.target.value as Filters["group"])}>
            <option value="">All</option>
            <option value="checker">checker</option>
            <option value="neutral">neutral</option>
          </select>
        </label>
        <label className="check">
          <input type="checkbox" checked={filters.flaggedOnly} onChange={(e) => set("flaggedOnly", e.target.checked)} />
          Flagged only
        </label>
        <span className="muted small-text">
          {rows ? `${rows.length} sessions` : "Loading..."}
          {lastPoll && `, updated ${localTime(lastPoll)}`}
        </span>
      </div>
      {error && <p className="error">{error}</p>}
      {rows && (
        <div className="table-wrap">
          <table className="participants">
            <thead>
              <tr>
                <th>Session</th>
                <th>Run</th>
                <th>Source</th>
                <th>Group</th>
                <th>Opened at</th>
                <th>Consent at</th>
                <th>Progress</th>
                <th>Status</th>
                <th>Elapsed</th>
                <th>Last step</th>
                <th>Audio check</th>
                <th>Reloads</th>
                <th>Flags</th>
              </tr>
            </thead>
            <tbody>
              {rows.map((r) => (
                <tr key={r.sessionId} className="clickable" onClick={() => setSelected(r.sessionId)}>
                  <td className="mono" title={r.sessionId}>
                    {shortId(r.sessionId)}
                  </td>
                  <td className="mono">{r.runCode}</td>
                  <td>{r.source}</td>
                  <td>{r.group ?? "-"}</td>
                  <td>{localTime(r.openedAt)}</td>
                  <td>{localTime(r.consentAt)}</td>
                  <td className="mono">{progressText(r)}</td>
                  <td>
                    <Chip chip={statusChip(r.status)} />
                  </td>
                  <td className="mono">{elapsedText(r, now)}</td>
                  <td className="mono">{r.lastStep ?? "-"}</td>
                  <td>{audioText(r)}</td>
                  <td className="mono">{r.reloadCount ?? 0}</td>
                  <td className="flags">
                    {(r.flags ?? []).map((f) => (
                      <Chip key={f} chip={flagChip(f)} />
                    ))}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
      {selected && (
        <ParticipantDrawer
          sessionId={selected}
          onClose={() => setSelected(null)}
          onDeleted={() => {
            setRows((cur) => cur?.filter((r) => r.sessionId !== selected) ?? cur);
            setReloadKey((n) => n + 1);
          }}
        />
      )}
    </section>
  );
}

/** k/N; N is "-" before allocation (section 5.1). */
export function progressText(r: ParticipantRow): string {
  const total = r.progress?.total;
  return `${r.progress?.index ?? 0}/${total ? total : "-"}`;
}

/** Right or wrong and number of plays (section 10); "-" until answered. */
export function audioText(r: ParticipantRow): string {
  const a = r.audioCheck;
  if (!a) return "-";
  const result = a.correct === true ? "right" : a.correct === false ? "wrong" : "not answered";
  return `${result}, ${a.plays ?? 0} play${a.plays === 1 ? "" : "s"}`;
}
