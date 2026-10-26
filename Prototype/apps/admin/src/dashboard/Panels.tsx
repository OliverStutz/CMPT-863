import type { AnalysisGroupFilter, Group, RunAnalysis } from "../api/types";
import { BarChart } from "../charts/BarChart";
import { TableTwin } from "../charts/common";
import { Donut } from "../charts/Donut";
import { Dumbbell, type DumbbellRow, type DumbbellSeries } from "../charts/Dumbbell";
import { fmtDuration, fmtInt, fmtNum } from "../charts/format";
import { Histogram } from "../charts/Histogram";
import { LineChart } from "../charts/LineChart";
import { StackedBar } from "../charts/StackedBar";
import { StatTile } from "../charts/StatTile";
import {
  BACKGROUND_TITLE,
  FLAG_TOGGLES,
  GROUPS,
  GROUP_COLOR,
  GROUP_LABEL,
  actionKeys,
  actionRows,
  backgroundBars,
  budgetKeys,
  budgetRows,
  ciPoint,
  countryNonAnswers,
  countrySlices,
  cyberLevelLabel,
  groupSeries,
  optionLabel,
  ratingXs,
  statusBars,
  tercileRows,
  trajectories,
} from "./model";
import { Card, TwoProportionLine } from "./Numbers";

interface PanelProps {
  a: RunAnalysis;
  tableView: boolean;
}

const GROUP_SERIES: DumbbellSeries[] = GROUPS.map((g) => ({ key: g, label: GROUP_LABEL[g], color: GROUP_COLOR[g] }));
const groupExtraColumns = GROUPS.map((g) => ({ key: g, label: GROUP_LABEL[g], numeric: true }));

export function KeyNumbers({ a, tableView }: PanelProps) {
  const c = a.counts;
  const s = c.byStatus;
  return (
    <>
      <div className="kpi-row">
        <StatTile label="Finished" value={fmtInt(s.finished)} caption={`of ${fmtInt(c.total)} sessions`} />
        <StatTile label="In progress" value={fmtInt(s.started + s.ongoing)} caption="started or ongoing" />
        <StatTile label="Abandoned" value={fmtInt(s.abandoned)} />
        <StatTile label="Blocked" value={fmtInt(s.blocked)} />
        <StatTile label="Opened without consent" value={fmtInt(s.opened)} />
        <StatTile
          label="Median duration"
          value={fmtDuration(c.durationsSec.median)}
          caption={`finished, middle half ${fmtDuration(c.durationsSec.p25)} to ${fmtDuration(c.durationsSec.p75)}`}
        />
        <StatTile
          label="Analysed per group"
          value={`${fmtInt(c.byGroup.checker.analysed)} / ${fmtInt(c.byGroup.neutral.analysed)}`}
          caption={`checker / neutral, ${fmtInt(c.analysed)} in all`}
        />
      </div>
      <Card title="Sessions by status" id="panel-status">
        <BarChart title="Sessions by status" data={statusBars(a)} valueLabel="Sessions" tableView={tableView} />
        <p className="muted small-text">
          Excluded from the analysis: {fmtInt(c.excluded.forced)} test link{c.excluded.forced === 1 ? "" : "s"} (forced
          allocation),{" "}
          {Object.entries(c.excluded.byFlag)
            .filter(([, n]) => n > 0)
            .map(([f, n]) => `${n} ${f}`)
            .join(", ") || "none by flag"}
          , {fmtInt(c.excluded.notIncluded)} not finished.
        </p>
      </Card>
    </>
  );
}

export function TrustPanel({ a, tableView, group }: PanelProps & { group: AnalysisGroupFilter }) {
  const xs = ratingXs(a);
  const items = a.config.items.filter((i) => i.id !== "trust");
  const trust = a.config.items.find((i) => i.id === "trust") ?? { id: "trust", label: "Trust", min: 0, max: 10 };
  return (
    <>
      <Card title="Trust over the study" id="panel-trust">
        <LineChart
          title={`${trust.label} (${trust.min} to ${trust.max}) at each rating point, group means with 95% CI`}
          xs={xs}
          series={groupSeries(a, "trust")}
          trajectories={trajectories(a, "trust", group)}
          yDomain={[trust.min, trust.max]}
          height={300}
          tableView={tableView}
          note="Thin grey lines: each participant (hollow ring: after that participant's denial). Triangle on the axis: a flagged slide."
        />
      </Card>
      <div className="dash-grid three">
        {items.map((it) => (
          <Card key={it.id} title={it.label} id={`panel-${it.id}`}>
            <LineChart
              title={`${it.label} (${it.min} to ${it.max}), group means with 95% CI`}
              xs={xs}
              series={groupSeries(a, it.id)}
              yDomain={[it.min, it.max]}
              height={190}
              tableView={tableView}
            />
          </Card>
        ))}
      </div>
    </>
  );
}

export function ActionsPanel({ a, tableView }: PanelProps) {
  return (
    <div className="dash-grid">
      {GROUPS.map((g) => (
        <Card key={g} title={`Actions per slide, ${GROUP_LABEL[g].toLowerCase()} group`} id={`panel-actions-${g}`}>
          <StackedBar
            title={`What the ${GROUP_LABEL[g].toLowerCase()} group did at each flagged slide`}
            keys={actionKeys(a, g)}
            rows={actionRows(a, g)}
            tableView={tableView}
          />
        </Card>
      ))}
    </div>
  );
}

export function BudgetPanel({ a, tableView }: PanelProps) {
  const b = a.hypotheses.budget;
  const sureRows: DumbbellRow[] = GROUPS.map((g) => ({
    id: g,
    label: GROUP_LABEL[g],
    color: GROUP_COLOR[g],
    points: { sure: ciPoint(b.byGroup[g].confidence) },
  }));
  return (
    <div className="dash-grid">
      <Card title="Budget vote by group" id="panel-budget">
        <StackedBar title="Budget vote by group" keys={budgetKeys(a)} rows={budgetRows(a)} tableView={tableView} />
        <TwoProportionLine t={b.approveTest} label="Share approving, checker vs neutral" />
      </Card>
      <Card title="Sureness of the vote" id="panel-sureness">
        <Dumbbell
          title="How sure (0 to 10), mean with 95% CI by group"
          series={[{ key: "sure", label: "Mean sureness", color: "var(--viz-ink-2)" }]}
          rows={sureRows}
          zeroLine={false}
          domain={[0, 10]}
          tableView={tableView}
        />
        <h4>Vote by last trust rating</h4>
        <StackedBar
          title="Budget vote by tercile of the participant's last trust rating"
          keys={budgetKeys(a)}
          rows={tercileRows(a)}
          tableView={tableView}
        />
      </Card>
    </div>
  );
}

export function CountryPanel({ a, tableView }: PanelProps) {
  const slices = countrySlices(a);
  const extraColumns = [{ key: "code", label: "Code" }, ...groupExtraColumns];
  const non = countryNonAnswers(a);
  return (
    <Card title="Participation per country" id="panel-country">
      <Donut title="Participants per country" data={slices} extraColumns={extraColumns} tableView={tableView} />
      <p className="muted small-text">
        Not in the chart: {fmtInt(non.preferNotToSay)} prefer not to say, {fmtInt(non.missing)} missing.
      </p>
      {tableView && (
        <BarChart
          title="Participants per country"
          data={slices.map((s) => ({
            id: s.id,
            label: s.label,
            value: s.value,
            tip: GROUPS.map((g) => ({
              color: GROUP_COLOR[g],
              value: String(s.extra?.[g] ?? 0),
              label: GROUP_LABEL[g],
            })),
          }))}
          valueLabel="Participants"
        />
      )}
    </Card>
  );
}

export function BackgroundPanel({ a, tableView }: PanelProps) {
  const items = a.config.background.filter((b) => b.kind !== "country" && a.background[b.id]);
  return (
    <div className="dash-grid">
      {items.map((it) => (
        <Card key={it.id} title={BACKGROUND_TITLE[it.id] ?? it.label} id={`panel-bg-${it.id}`}>
          <BarChart
            title={it.label}
            data={backgroundBars(a, it.id)}
            valueLabel="Participants"
            extraColumns={groupExtraColumns}
            tableView={tableView}
          />
        </Card>
      ))}
    </div>
  );
}

export function QualityPanel({ a, tableView }: PanelProps) {
  const q = a.quality;
  const toggles = new Map(FLAG_TOGGLES.map((t) => [t.flag, t]));
  const flags = Object.keys(q.flags).sort((x, y) => q.flags[y] - q.flags[x]);
  return (
    <div className="dash-grid">
      <Card title="Flags" id="panel-flags">
        <TableTwin
          caption="Sessions with each flag, before the flag toggles (what each toggle would remove)"
          columns={[
            { key: "flag", label: "Flag" },
            { key: "n", label: "Sessions", numeric: true },
            { key: "now", label: "In this view" },
          ]}
          rows={flags.map((f) => ({
            flag: f,
            n: fmtInt(q.flags[f]),
            now:
              f === "forced_allocation"
                ? "always excluded"
                : a.filters.exclude.includes(f)
                  ? "excluded"
                  : toggles.has(f)
                    ? "included"
                    : "included (no toggle)",
          }))}
        />
      </Card>
      <Card title="Audio check" id="panel-audio">
        <StackedBar
          title="Audio check answers"
          keys={[
            { key: "correct", label: "Right word", color: "var(--viz-good)" },
            { key: "wrong", label: "Wrong word", color: "var(--viz-critical)" },
            { key: "missing", label: "No answer", color: "var(--viz-other)", ink: "var(--viz-on-other)" },
          ]}
          rows={[{ id: "audio", label: "Sessions", values: { ...q.audioCheck } }]}
          tableView={tableView}
        />
        <div className="kpi-row">
          <StatTile
            label="Sessions with a reload"
            value={fmtInt(q.reloads.sessions)}
            caption={`${fmtInt(q.reloads.total)} reloads in all`}
          />
        </div>
      </Card>
      <Card title="Rating timeouts per point" id="panel-timeouts">
        <BarChart
          title="Rating timeouts per point"
          orientation="vertical"
          height={110}
          valueLabel="Timeouts"
          data={a.config.ratingPoints.map((rp) => ({
            id: rp.point,
            label: rp.point,
            value: q.ratingTimeouts[rp.point] ?? 0,
          }))}
          tableView={tableView}
        />
      </Card>
      <Card title="Duration of finished sessions" id="panel-duration">
        <Histogram
          title="Duration of finished sessions (minutes)"
          bins={a.counts.durationsSec.histogram.map((b) => ({ from: b.fromSec, to: b.toSec, n: b.n }))}
          formatEdge={(v) => fmtNum(v / 60, 1)}
          tableView={tableView}
        />
      </Card>
    </div>
  );
}

function moderatorRows(a: RunAnalysis, which: "h1" | "h2", by: "cyberKnowledge" | "aiUse"): DumbbellRow[] {
  const m = a.hypotheses.moderators;
  if (by === "cyberKnowledge")
    return m.cyberKnowledge.map((r) => ({
      id: String(r.level),
      label: cyberLevelLabel(a, r.level),
      points: { checker: ciPoint(r[which].checker), neutral: ciPoint(r[which].neutral) },
    }));
  return m.aiUse.map((r) => ({
    id: r.optionId,
    label: optionLabel(a, "aiUse", r.optionId),
    points: { checker: ciPoint(r[which].checker), neutral: ciPoint(r[which].neutral) },
  }));
}

export function ModeratorsPanel({ a, tableView }: PanelProps) {
  const sp = a.hypotheses.moderators.spearman;
  const rho = (r: Record<Group, number | null>) => `checker ${fmtNum(r.checker)}, neutral ${fmtNum(r.neutral)}`;
  const panels: { key: string; title: string; which: "h1" | "h2"; by: "cyberKnowledge" | "aiUse"; rho: string }[] = [
    {
      key: "h1-ck",
      title: "H1 change at the flags by cyber security knowledge",
      which: "h1",
      by: "cyberKnowledge",
      rho: rho(sp.cyberKnowledgeH1),
    },
    { key: "h1-ai", title: "H1 change at the flags by AI use", which: "h1", by: "aiUse", rho: rho(sp.aiUseH1) },
    {
      key: "h2-ck",
      title: "H2 change after a reason by cyber security knowledge",
      which: "h2",
      by: "cyberKnowledge",
      rho: rho(sp.cyberKnowledgeH2),
    },
    { key: "h2-ai", title: "H2 change after a reason by AI use", which: "h2", by: "aiUse", rho: rho(sp.aiUseH2) },
  ];
  return (
    <div className="dash-grid">
      {panels.map((p) => (
        <Card key={p.key} title={p.title} id={`panel-mod-${p.key}`}>
          <Dumbbell
            title={p.title}
            series={GROUP_SERIES}
            rows={moderatorRows(a, p.which, p.by)}
            tableView={tableView}
          />
          <p className="numbers-line muted">Spearman correlation (descriptive): {p.rho}</p>
        </Card>
      ))}
    </div>
  );
}
