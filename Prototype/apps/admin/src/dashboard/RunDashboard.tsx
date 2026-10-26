import { useState } from "react";
import { getRunAnalysis } from "../api/client";
import type { AnalysisQuery, RunAnalysis } from "../api/types";
import { localTime } from "../lib/time";
import { useAsync } from "../lib/useAsync";
import { DEFAULT_FILTERS, FLAG_TOGGLES, toggleExclude, type DashboardFilters } from "./model";
import { HypothesisCards } from "./Hypotheses";
import {
  ActionsPanel,
  BackgroundPanel,
  BudgetPanel,
  CountryPanel,
  KeyNumbers,
  ModeratorsPanel,
  QualityPanel,
  TrustPanel,
} from "./Panels";
import "./dashboard.css";

export type AnalysisLoader = (runId: string, q: AnalysisQuery) => Promise<RunAnalysis>;

export interface RunDashboardProps {
  runId: string;
  onBack: () => void;
  /** Opens the Export tab; the link is hidden without it. */
  onOpenExport?: () => void;
  /** Data source; the API by default, the fixture in the preview. */
  load?: AnalysisLoader;
  /** Seeds the first render (preview and the render test). */
  initial?: RunAnalysis;
  initialTableView?: boolean;
  /** Shown above the filters, e.g. that the preview does not call the API. */
  notice?: string;
}

/** The section titles, in the order of plan section 3 (the render test checks each one appears). */
export const SECTION_TITLES = [
  "Key numbers",
  "Trust over the study",
  "Hypotheses",
  "Actions per slide",
  "Budget",
  "Participation per country",
  "Background",
  "Quality",
  "Moderators",
];

export function RunDashboard({
  runId,
  onBack,
  onOpenExport,
  load = getRunAnalysis,
  initial,
  initialTableView = false,
  notice,
}: RunDashboardProps) {
  const [filters, setFilters] = useState<DashboardFilters>(DEFAULT_FILTERS);
  const [tableView, setTableView] = useState(initialTableView);
  const analysis = useAsync(
    () => load(runId, { include: filters.include, exclude: filters.exclude, group: filters.group }),
    [runId, filters],
    initial ?? null,
  );
  const a = analysis.data;
  // Hold the previous render at reduced opacity while a filter change reloads; the seeded first render stays as is.
  const refetching = analysis.loading && a !== null && a !== initial;

  return (
    <section className="dashboard">
      <div className="dash-top">
        <button className="small" onClick={onBack}>
          Back to runs
        </button>
        {a && (
          <>
            <h2 className="dash-title">{a.run.name}</h2>
            <span className={`chip ${a.run.status === "open" ? "chip-success" : "chip-neutral"}`}>{a.run.status}</span>
            {a.run.testOverrides && <span className="chip chip-info">test overrides</span>}
          </>
        )}
        <span className="dash-spacer" />
        {a && (
          <button className="small" onClick={() => downloadJson(a)}>
            Download analysis JSON
          </button>
        )}
        {onOpenExport && (
          <button className="small" onClick={onOpenExport}>
            Go to Export
          </button>
        )}
      </div>
      {a && (
        <p className="muted small-text dash-meta">
          Code <span className="mono">{a.run.code}</span>, source {a.run.source}, config{" "}
          <span className="mono">{a.run.configVersion}</span>, created {localTime(a.run.createdAt)}; computed{" "}
          {localTime(a.generatedAt)}. Tests here are for reading the pilot and the running study; the confirmatory
          analysis runs offline on the export.
        </p>
      )}
      {notice && <p className="chip chip-info dash-notice">{notice}</p>}
      <FilterRow filters={filters} setFilters={setFilters} tableView={tableView} setTableView={setTableView} a={a} />
      {analysis.error && <p className="error">{analysis.error}</p>}
      {!a && !analysis.error && <p className="muted">Loading...</p>}
      {a && (
        <div className={refetching ? "dash-body refetching" : "dash-body"} aria-busy={refetching}>
          <h2>Key numbers</h2>
          <KeyNumbers a={a} tableView={tableView} />
          <h2>Trust over the study</h2>
          <TrustPanel a={a} tableView={tableView} group={filters.group} />
          <h2>Hypotheses</h2>
          <HypothesisCards a={a} tableView={tableView} />
          <h2>Actions per slide</h2>
          <ActionsPanel a={a} tableView={tableView} />
          <h2>Budget</h2>
          <BudgetPanel a={a} tableView={tableView} />
          <h2>Participation per country</h2>
          <CountryPanel a={a} tableView={tableView} />
          <h2>Background</h2>
          <BackgroundPanel a={a} tableView={tableView} />
          <h2>Quality</h2>
          <QualityPanel a={a} tableView={tableView} />
          <h2>Moderators</h2>
          <ModeratorsPanel a={a} tableView={tableView} />
        </div>
      )}
    </section>
  );
}

function FilterRow(props: {
  filters: DashboardFilters;
  setFilters: (f: DashboardFilters) => void;
  tableView: boolean;
  setTableView: (v: boolean) => void;
  a: RunAnalysis | null;
}) {
  const { filters, setFilters, tableView, setTableView, a } = props;
  const flagCount = (f: string) => a?.quality.flags[f] ?? 0;
  return (
    <div className="filters dash-filters" role="group" aria-label="Filters for every chart below">
      <label>
        Participants
        <select
          value={filters.include}
          onChange={(e) => setFilters({ ...filters, include: e.target.value as DashboardFilters["include"] })}
        >
          <option value="finished">Finished only</option>
          <option value="consented">All consented</option>
        </select>
      </label>
      <fieldset className="dash-flags">
        <legend>Exclude flagged sessions</legend>
        {FLAG_TOGGLES.map((t) => (
          <label
            key={t.flag}
            className="check"
            title={`${t.why}; default ${t.excludedByDefault ? "excluded" : "included"}`}
          >
            <input
              type="checkbox"
              checked={filters.exclude.includes(t.flag)}
              onChange={(e) => setFilters(toggleExclude(filters, t.flag, e.target.checked))}
            />
            <span className="mono small-text">{t.flag}</span>
            {a && <span className="muted small-text">({flagCount(t.flag)})</span>}
          </label>
        ))}
      </fieldset>
      <label>
        Group
        <select
          value={filters.group}
          onChange={(e) => setFilters({ ...filters, group: e.target.value as DashboardFilters["group"] })}
        >
          <option value="both">Both</option>
          <option value="checker">Checker</option>
          <option value="neutral">Neutral</option>
        </select>
      </label>
      <label className="check">
        <input type="checkbox" checked={tableView} onChange={(e) => setTableView(e.target.checked)} />
        Table view
      </label>
      <button className="small" onClick={() => setFilters(DEFAULT_FILTERS)}>
        Reset filters
      </button>
    </div>
  );
}

// DECISION: the download saves the document the charts read, as loaded (same filters), named
// analysis-<run code>-<generatedAt>.json; no second request.
function downloadJson(a: RunAnalysis) {
  const blob = new Blob([JSON.stringify(a, null, 2)], { type: "application/json" });
  const url = URL.createObjectURL(blob);
  const el = document.createElement("a");
  el.href = url;
  el.download = `analysis-${a.run.code}-${a.generatedAt.replace(/[:.]/g, "-")}.json`;
  document.body.appendChild(el);
  el.click();
  el.remove();
  setTimeout(() => URL.revokeObjectURL(url), 1000);
}
