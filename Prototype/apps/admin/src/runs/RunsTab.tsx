import { Fragment, useEffect, useState, type FormEvent } from "react";
import { ApiError, createRun, listConfigs, listRuns, patchRun } from "../api/client";
import { AUDIO_MODE_LABEL, audioMode, type ConfigSummary, type Run } from "../api/types";
import { AudioChip } from "../configs/ConfigsTab";
import { useAsync } from "../lib/useAsync";
import { localTime } from "../lib/time";
import { CopyButton, RunQr, runUrl } from "./RunLink";
import { RunDashboard } from "../dashboard/RunDashboard";

const SOURCE_SUGGESTIONS = ["LinkedIn", "WhatsApp", "Email", "Other"];

// DECISION: section 10 says "only validated versions" may be chosen. A locked version is a validated one that a
// run with sessions already uses (2.16), so it is offered too; draft versions are not.
export function selectableConfigs(configs: ConfigSummary[]): ConfigSummary[] {
  return configs.filter((c) => c.status === "validated" || c.status === "locked");
}

/** Run id from a `#run=<id>` URL hash, or null. */
export function runFromHash(hash: string): string | null {
  const m = /^#run=([^&]+)/.exec(hash);
  if (!m) return null;
  try {
    return decodeURIComponent(m[1]);
  } catch {
    return null;
  }
}

// DECISION: the dashboard is addressed by `#run=<id>` (no router library). Opening it pushes a history entry, so the
// browser's Back button returns to the table; the tab stays mounted under the login overlay, so a 401 resumes in place.
export function RunsTab({ onOpenExport }: { onOpenExport?: () => void } = {}) {
  const [dashboardRun, setDashboardRun] = useState<string | null>(() => runFromHash(window.location.hash));
  useEffect(() => {
    const onHash = () => setDashboardRun(runFromHash(window.location.hash));
    window.addEventListener("hashchange", onHash);
    return () => window.removeEventListener("hashchange", onHash);
  }, []);
  const openDashboard = (id: string) => {
    window.location.hash = `run=${encodeURIComponent(id)}`;
    setDashboardRun(id);
  };
  const closeDashboard = () => {
    if (window.location.hash) window.history.pushState(null, "", window.location.pathname + window.location.search);
    setDashboardRun(null);
  };

  if (dashboardRun) return <RunDashboard runId={dashboardRun} onBack={closeDashboard} onOpenExport={onOpenExport} />;
  return <RunsTable onOpenDashboard={openDashboard} />;
}

function RunsTable({ onOpenDashboard }: { onOpenDashboard: (id: string) => void }) {
  const runs = useAsync(listRuns);
  const configs = useAsync(listConfigs);
  const [expanded, setExpanded] = useState<string | null>(null);
  const [busyId, setBusyId] = useState<string | null>(null);
  const [actionError, setActionError] = useState<string | null>(null);

  async function patch(run: Run, req: Parameters<typeof patchRun>[1]) {
    setBusyId(run.id);
    setActionError(null);
    try {
      await patchRun(run.id, req);
      runs.reload();
    } catch (e) {
      setActionError(e instanceof Error ? e.message : String(e));
    } finally {
      setBusyId(null);
    }
  }
  const toggleStatus = (run: Run) => patch(run, { status: run.status === "open" ? "closed" : "open" });
  const toggleTest = (run: Run) => patch(run, { testOverrides: !run.testOverrides });

  return (
    <section>
      <NewRunForm
        configs={configs.data ? selectableConfigs(configs.data) : []}
        configsError={configs.error}
        onCreated={(id) => {
          runs.reload();
          setExpanded(id);
        }}
      />
      <h2>Runs</h2>
      {runs.error && <p className="error">{runs.error}</p>}
      {actionError && <p className="error">{actionError}</p>}
      {runs.data && runs.data.length === 0 && <p className="muted">No runs yet.</p>}
      {runs.data && runs.data.length > 0 && (
        <div className="table-wrap">
          <table>
            <thead>
              <tr>
                <th>Name</th>
                <th>Source</th>
                <th>Code</th>
                <th>Link</th>
                <th>Config</th>
                <th>Status</th>
                <th title="Opened / Started / Ongoing / Finished / Abandoned / Blocked">O / S / On / F / A / B</th>
                <th>Finished per group</th>
                <th>Created</th>
                <th />
              </tr>
            </thead>
            <tbody>
              {runs.data.map((r) => (
                <Fragment key={r.id}>
                  <tr>
                    <td>
                      <button
                        className="link"
                        title="Open the dashboard of this run"
                        onClick={() => onOpenDashboard(r.id)}
                      >
                        {r.name}
                      </button>
                    </td>
                    <td>{r.source}</td>
                    <td className="mono">{r.code}</td>
                    <td>
                      <span className="mono small-text">{runUrl(r.code)}</span> <CopyButton text={runUrl(r.code)} />{" "}
                      <button className="small" onClick={() => setExpanded(expanded === r.id ? null : r.id)}>
                        {expanded === r.id ? "Hide QR" : "QR"}
                      </button>
                    </td>
                    <td className="mono">
                      {r.configVersion}
                      {audioMode(configs.data?.find((c) => c.version === r.configVersion)) !== "spoken" && (
                        <>
                          {" "}
                          <AudioChip mode={audioMode(configs.data?.find((c) => c.version === r.configVersion))} />
                        </>
                      )}
                    </td>
                    <td>
                      <span className={`chip ${r.status === "open" ? "chip-success" : "chip-neutral"}`}>
                        {r.status}
                      </span>
                      {r.testOverrides && (
                        <>
                          {" "}
                          <span
                            className="chip chip-info"
                            title="Links may carry ?group=checker|neutral&denial=<slide id>; such sessions are flagged forced_allocation and skip the round robin"
                          >
                            test overrides
                          </span>
                        </>
                      )}
                    </td>
                    <td className="mono">
                      {r.counts
                        ? [
                            r.counts.opened,
                            r.counts.started,
                            r.counts.ongoing,
                            r.counts.finished,
                            r.counts.abandoned,
                            r.counts.blocked,
                          ]
                            .map((n) => n ?? 0)
                            .join(" / ")
                        : "-"}
                    </td>
                    <td>
                      checker {r.finishedByGroup?.checker ?? 0}, neutral {r.finishedByGroup?.neutral ?? 0}
                    </td>
                    <td>{localTime(r.createdAt)}</td>
                    <td>
                      <button className="small" onClick={() => onOpenDashboard(r.id)}>
                        Dashboard
                      </button>{" "}
                      <button className="small" disabled={busyId === r.id} onClick={() => void toggleStatus(r)}>
                        {r.status === "open" ? "Close" : "Reopen"}
                      </button>{" "}
                      <button
                        className="small"
                        disabled={busyId === r.id}
                        title="Test runs only: allow ?group= and ?denial= in the link"
                        onClick={() => void toggleTest(r)}
                      >
                        {r.testOverrides ? "Overrides off" : "Overrides on"}
                      </button>
                    </td>
                  </tr>
                  {expanded === r.id && (
                    <tr className="expanded">
                      <td colSpan={10}>
                        <RunQr code={r.code} />
                      </td>
                    </tr>
                  )}
                </Fragment>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </section>
  );
}

function NewRunForm(props: { configs: ConfigSummary[]; configsError: string | null; onCreated: (id: string) => void }) {
  const [name, setName] = useState("");
  const [source, setSource] = useState("");
  const [configVersion, setConfigVersion] = useState("");
  const [testOverrides, setTestOverrides] = useState(false);
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function submit(e: FormEvent) {
    e.preventDefault();
    setBusy(true);
    setError(null);
    try {
      const res = await createRun({ name: name.trim(), source: source.trim(), configVersion, testOverrides });
      setName("");
      setSource("");
      setTestOverrides(false);
      props.onCreated(res.id);
    } catch (err) {
      setError(err instanceof ApiError ? `${err.code}: ${err.message}` : String(err));
    } finally {
      setBusy(false);
    }
  }

  return (
    <form className="card form-inline" onSubmit={submit}>
      <h2>New run</h2>
      <label>
        Name
        <input value={name} onChange={(e) => setName(e.target.value)} placeholder="LinkedIn post 1" required />
      </label>
      <label>
        Source
        <input list="source-suggestions" value={source} onChange={(e) => setSource(e.target.value)} required />
        <datalist id="source-suggestions">
          {SOURCE_SUGGESTIONS.map((s) => (
            <option key={s} value={s} />
          ))}
        </datalist>
      </label>
      <label>
        Config version
        <select value={configVersion} onChange={(e) => setConfigVersion(e.target.value)} required>
          <option value="">Choose...</option>
          {props.configs.map((c) => (
            <option key={c.version} value={c.version}>
              {c.version} ({c.status}
              {audioMode(c) && audioMode(c) !== "spoken" ? `, ${AUDIO_MODE_LABEL[audioMode(c)!]}` : ""})
            </option>
          ))}
        </select>
      </label>
      <label
        className="checkbox"
        title="Sessions opened with ?group=checker|neutral&denial=<slide id> get that allocation, are flagged forced_allocation and do not count in the round robin. Never for a real run."
      >
        <input type="checkbox" checked={testOverrides} onChange={(e) => setTestOverrides(e.target.checked)} />
        Test run: allow group and denial overrides in the link
      </label>
      <button type="submit" disabled={busy || !name.trim() || !source.trim() || !configVersion}>
        {busy ? "Creating..." : "Create run"}
      </button>
      {props.configsError && <p className="error">Configs: {props.configsError}</p>}
      {error && <p className="error">{error}</p>}
    </form>
  );
}
