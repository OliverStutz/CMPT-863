import { Fragment, useState, type FormEvent } from "react";
import { ApiError, createRun, listConfigs, listRuns, patchRun } from "../api/client";
import type { ConfigSummary, Run } from "../api/types";
import { useAsync } from "../lib/useAsync";
import { localTime } from "../lib/time";
import { CopyButton, RunQr, runUrl } from "./RunLink";

const SOURCE_SUGGESTIONS = ["LinkedIn", "WhatsApp", "Email", "Other"];

// DECISION: section 10 says "only validated versions" may be chosen. A locked version is a validated one that a
// run with sessions already uses (2.16), so it is offered too; draft versions are not.
export function selectableConfigs(configs: ConfigSummary[]): ConfigSummary[] {
  return configs.filter((c) => c.status === "validated" || c.status === "locked");
}

export function RunsTab() {
  const runs = useAsync(listRuns);
  const configs = useAsync(listConfigs);
  const [expanded, setExpanded] = useState<string | null>(null);
  const [busyId, setBusyId] = useState<string | null>(null);
  const [actionError, setActionError] = useState<string | null>(null);

  async function toggleStatus(run: Run) {
    setBusyId(run.id);
    setActionError(null);
    try {
      await patchRun(run.id, { status: run.status === "open" ? "closed" : "open" });
      runs.reload();
    } catch (e) {
      setActionError(e instanceof Error ? e.message : String(e));
    } finally {
      setBusyId(null);
    }
  }

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
                    <td>{r.name}</td>
                    <td>{r.source}</td>
                    <td className="mono">{r.code}</td>
                    <td>
                      <span className="mono small-text">{runUrl(r.code)}</span> <CopyButton text={runUrl(r.code)} />{" "}
                      <button className="small" onClick={() => setExpanded(expanded === r.id ? null : r.id)}>
                        {expanded === r.id ? "Hide QR" : "QR"}
                      </button>
                    </td>
                    <td className="mono">{r.configVersion}</td>
                    <td>
                      <span className={`chip ${r.status === "open" ? "chip-success" : "chip-neutral"}`}>{r.status}</span>
                    </td>
                    <td className="mono">
                      {r.counts
                        ? [r.counts.opened, r.counts.started, r.counts.ongoing, r.counts.finished, r.counts.abandoned, r.counts.blocked]
                            .map((n) => n ?? 0)
                            .join(" / ")
                        : "-"}
                    </td>
                    <td>
                      checker {r.finishedByGroup?.checker ?? 0}, neutral {r.finishedByGroup?.neutral ?? 0}
                    </td>
                    <td>{localTime(r.createdAt)}</td>
                    <td>
                      <button className="small" disabled={busyId === r.id} onClick={() => void toggleStatus(r)}>
                        {r.status === "open" ? "Close" : "Reopen"}
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
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function submit(e: FormEvent) {
    e.preventDefault();
    setBusy(true);
    setError(null);
    try {
      const res = await createRun({ name: name.trim(), source: source.trim(), configVersion });
      setName("");
      setSource("");
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
              {c.version} ({c.status})
            </option>
          ))}
        </select>
      </label>
      <button type="submit" disabled={busy || !name.trim() || !source.trim() || !configVersion}>
        {busy ? "Creating..." : "Create run"}
      </button>
      {props.configsError && <p className="error">Configs: {props.configsError}</p>}
      {error && <p className="error">{error}</p>}
    </form>
  );
}
