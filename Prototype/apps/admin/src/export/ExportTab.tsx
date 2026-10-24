import { useState } from "react";
import { downloadExport, listRuns } from "../api/client";
import type { Run } from "../api/types";
import { useAsync } from "../lib/useAsync";

export function ExportTab() {
  const runs = useAsync(listRuns);
  const [includeInvalid, setIncludeInvalid] = useState(false);
  const [includeIp, setIncludeIp] = useState(false);
  const [busy, setBusy] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  async function go(format: "csv" | "json", run: Run | null) {
    const key = `${run?.id ?? "all"}-${format}`;
    setBusy(key);
    setError(null);
    try {
      await downloadExport({ format, runId: run?.id, includeInvalid, includeIp }, run ? run.code : "all");
    } catch (e) {
      setError(e instanceof Error ? e.message : String(e));
    } finally {
      setBusy(null);
    }
  }

  const buttons = (run: Run | null) => (
    <>
      {(["csv", "json"] as const).map((f) => {
        const key = `${run?.id ?? "all"}-${f}`;
        return (
          <button key={f} className="small" disabled={busy !== null} onClick={() => void go(f, run)}>
            {busy === key ? "Downloading..." : f.toUpperCase()}
          </button>
        );
      })}
    </>
  );

  return (
    <section>
      <h2>Export</h2>
      <p className="muted">
        CSV has one row per rating point; JSON has one object per session with ratings, answers and the full event log.
      </p>
      <div className="filters">
        <label className="check">
          <input type="checkbox" checked={includeInvalid} onChange={(e) => setIncludeInvalid(e.target.checked)} />
          Include invalid events
        </label>
        <label className="check">
          <input type="checkbox" checked={includeIp} onChange={(e) => setIncludeIp(e.target.checked)} />
          Include IP address
        </label>
      </div>
      {error && <p className="error">{error}</p>}
      {runs.error && <p className="error">{runs.error}</p>}
      <div className="table-wrap">
        <table>
          <thead>
            <tr>
              <th>Run</th>
              <th>Code</th>
              <th>Config</th>
              <th>Download</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>
                <strong>All runs</strong>
              </td>
              <td />
              <td />
              <td className="row-actions">{buttons(null)}</td>
            </tr>
            {runs.data?.map((r) => (
              <tr key={r.id}>
                <td>{r.name}</td>
                <td className="mono">{r.code}</td>
                <td className="mono">{r.configVersion}</td>
                <td className="row-actions">{buttons(r)}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </section>
  );
}
