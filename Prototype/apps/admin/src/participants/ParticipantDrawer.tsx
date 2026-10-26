import { useEffect, useState, type FormEvent } from "react";
import { ApiError, deleteSession, getParticipant } from "../api/client";
import type { ParticipantDetail } from "../api/types";
import { Chip } from "../lib/Chip";
import { flagChip, statusChip } from "../lib/status";
import { localTime } from "../lib/time";
import "./drawer.css";

export function ParticipantDrawer({
  sessionId,
  onClose,
  onDeleted,
}: {
  sessionId: string;
  onClose: () => void;
  /** Called after the session was deleted; the drawer closes itself. */
  onDeleted?: () => void;
}) {
  const [detail, setDetail] = useState<ParticipantDetail | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [tick, setTick] = useState(0);
  const [onlyRejected, setOnlyRejected] = useState(false);

  useEffect(() => {
    let alive = true;
    getParticipant(sessionId).then(
      (d) => alive && (setDetail(d), setError(null)),
      (e: unknown) => alive && setError(e instanceof Error ? e.message : String(e)),
    );
    return () => {
      alive = false;
    };
  }, [sessionId, tick]);

  useEffect(() => {
    const onKey = (e: KeyboardEvent) => e.key === "Escape" && onClose();
    window.addEventListener("keydown", onKey);
    return () => window.removeEventListener("keydown", onKey);
  }, [onClose]);

  const s = detail?.session;
  const events = (detail?.events ?? []).slice().sort((a, b) => a.eventNo - b.eventNo);
  const shown = onlyRejected ? events.filter((e) => !e.valid) : events;
  const rejected = events.filter((e) => !e.valid).length;

  return (
    <div className="drawer-backdrop" onClick={onClose}>
      <aside className="drawer" onClick={(e) => e.stopPropagation()} aria-label="Session detail">
        <div className="drawer-head">
          <h2 className="mono">{sessionId}</h2>
          <button className="small" onClick={() => setTick((n) => n + 1)}>
            Refresh
          </button>
          <button className="small" onClick={onClose}>
            Close
          </button>
        </div>
        {error && <p className="error">{error}</p>}
        {!detail && !error && <p className="muted">Loading...</p>}
        {s && (
          <dl className="facts">
            <dt>Status</dt>
            <dd>{s.status ? <Chip chip={statusChip(s.status)} /> : "-"}</dd>
            <dt>Run</dt>
            <dd className="mono">{s.runCode ?? "-"}</dd>
            <dt>Source</dt>
            <dd>{s.source ?? "-"}</dd>
            <dt>Config</dt>
            <dd className="mono">{s.configVersion ?? "-"}</dd>
            <dt>Group</dt>
            <dd>{s.group ?? "-"}</dd>
            <dt>Denial slide</dt>
            <dd className="mono">{s.denialSlide ?? "-"}</dd>
            <dt>Current state</dt>
            <dd className="mono">{s.currentState ?? "-"}</dd>
            <dt>Opened</dt>
            <dd>{localTime(s.openedAt)}</dd>
            <dt>Consent</dt>
            <dd>{localTime(s.consentAt)}</dd>
            <dt>Started</dt>
            <dd>{localTime(s.startedAt)}</dd>
            <dt>Ongoing</dt>
            <dd>{localTime(s.ongoingAt)}</dd>
            <dt>Finished</dt>
            <dd>{localTime(s.finishedAt)}</dd>
            <dt>Last event</dt>
            <dd>{localTime(s.lastEventAt)}</dd>
            <dt>Browser</dt>
            <dd>{s.uaFamily ?? "-"}</dd>
            <dt>Flags</dt>
            <dd>{(s.flags ?? []).length ? (s.flags ?? []).map((f) => <Chip key={f} chip={flagChip(f)} />) : "-"}</dd>
          </dl>
        )}
        {detail && (
          <>
            <div className="drawer-sub">
              <h3>
                Event timeline ({events.length} events, {rejected} rejected)
              </h3>
              <label className="check">
                <input type="checkbox" checked={onlyRejected} onChange={(e) => setOnlyRejected(e.target.checked)} />
                Rejected only
              </label>
            </div>
            <div className="table-wrap">
              <table className="events">
                <thead>
                  <tr>
                    <th>#</th>
                    <th>Server time</th>
                    <th>Seq</th>
                    <th>From</th>
                    <th>Type</th>
                    <th>State</th>
                    <th>Result</th>
                    <th>Payload</th>
                  </tr>
                </thead>
                <tbody>
                  {shown.map((e) => (
                    <tr key={e.eventNo} className={e.valid ? "" : "rejected"}>
                      <td className="mono">{e.eventNo}</td>
                      <td>{localTime(e.serverTs)}</td>
                      <td className="mono">{e.clientSeq ?? "-"}</td>
                      <td>{e.source}</td>
                      <td className="mono">{e.type}</td>
                      <td className="mono small-text">
                        {e.stateBefore ?? "-"} to {e.stateAfter ?? "-"}
                      </td>
                      <td>
                        {e.valid ? "accepted" : <span className="chip chip-danger">{e.reason ?? "rejected"}</span>}
                      </td>
                      <td className="mono small-text payload">{payloadText(e.payload)}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
            <h3>Ratings</h3>
            <pre className="json">{JSON.stringify(detail.ratings, null, 2)}</pre>
            <h3>Answers</h3>
            <pre className="json">{JSON.stringify(detail.answers, null, 2)}</pre>
          </>
        )}
        <DeleteSession
          sessionId={sessionId}
          onDeleted={() => {
            onDeleted?.();
            onClose();
          }}
        />
      </aside>
    </div>
  );
}

function payloadText(p: unknown): string {
  if (p === null || p === undefined) return "";
  const t = JSON.stringify(p);
  return t === "{}" ? "" : t;
}

export const REASON_MAX = 500;

/** The confirm box may submit only when the typed id matches exactly and the reason has 1 to 500 characters. */
export function canDelete(sessionId: string, typedId: string, reason: string): boolean {
  const r = reason.trim();
  return typedId === sessionId && r.length >= 1 && r.length <= REASON_MAX;
}

// DECISION: deletion is the last section of the drawer and needs two steps: open the confirm box, then type the
// session id exactly (no trimming, no case folding) and give a reason. The reason is sent trimmed.
function DeleteSession({ sessionId, onDeleted }: { sessionId: string; onDeleted: () => void }) {
  const [open, setOpen] = useState(false);
  const [typedId, setTypedId] = useState("");
  const [reason, setReason] = useState("");
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function submit(e: FormEvent) {
    e.preventDefault();
    if (!canDelete(sessionId, typedId, reason)) return;
    setBusy(true);
    setError(null);
    try {
      await deleteSession(sessionId, { confirm: typedId, reason: reason.trim() });
      onDeleted();
    } catch (err) {
      setError(err instanceof ApiError ? `${err.code}: ${err.message}` : String(err));
      setBusy(false);
    }
  }

  if (!open)
    return (
      <div className="delete-zone">
        <button className="danger" onClick={() => setOpen(true)}>
          Delete this participant's data
        </button>
      </div>
    );
  return (
    <form className="delete-zone delete-box" onSubmit={submit}>
      <h3>Delete this participant's data</h3>
      <p>
        This removes the session, its events, ratings and answers for good. A record without personal data stays (run,
        group, time and your reason), so the deletion can be counted. The export and the dashboard will no longer show
        this session.
      </p>
      <label>
        Type the session id to confirm
        <input
          className="mono"
          value={typedId}
          onChange={(e) => setTypedId(e.target.value)}
          placeholder={sessionId}
          autoComplete="off"
          spellCheck={false}
        />
      </label>
      <label>
        Reason ({reason.trim().length} of {REASON_MAX} characters)
        <textarea
          value={reason}
          maxLength={REASON_MAX}
          rows={3}
          onChange={(e) => setReason(e.target.value)}
          placeholder="Participant asked by e-mail on ..."
        />
      </label>
      <div className="row-actions">
        <button type="submit" className="danger" disabled={busy || !canDelete(sessionId, typedId, reason)}>
          {busy ? "Deleting..." : "Delete permanently"}
        </button>
        <button type="button" disabled={busy} onClick={() => setOpen(false)}>
          Cancel
        </button>
      </div>
      {error && <p className="error">{error}</p>}
    </form>
  );
}
