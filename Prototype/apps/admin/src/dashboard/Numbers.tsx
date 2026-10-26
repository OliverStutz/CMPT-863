// The numbers under each hypothesis chart: n per group, mean, CI, test statistic, df, p, Hedges g and the reading.
import type { ReactNode } from "react";
import type { Comparison, OneSample, Paired, Reading, Summary, TwoProportion } from "../api/types";
import { fmtCi, fmtNum, fmtP, fmtPct, fmtSigned } from "../charts/format";
import { Chip } from "../lib/Chip";
import { GROUPS, GROUP_LABEL, readingChip } from "./model";

export function ReadingChip({ reading, hypothesis }: { reading: Reading; hypothesis?: string }) {
  return <Chip chip={readingChip(reading, hypothesis)} />;
}

/** "t(41.3) = -6.85, p < .001" */
export function testText(stat: string, t: number | null, df: number | null, p: number | null): string {
  if (t === null) return `${stat} not computed`;
  return `${stat}(${fmtNum(df)}) = ${fmtNum(t)}, p ${fmtP(p).startsWith("<") ? fmtP(p) : `= ${fmtP(p)}`}`;
}

function SummaryRows({ rows }: { rows: { label: string; s: Summary }[] }) {
  return (
    <>
      {rows.map(({ label, s }) => (
        <tr key={label}>
          <th scope="row">{label}</th>
          <td className="num">{s.n}</td>
          <td className="num">{fmtNum(s.mean)}</td>
          <td className="num">{fmtNum(s.sd)}</td>
          <td className="num">
            <span className="nw">{fmtCi(s.ci95)}</span>
          </td>
        </tr>
      ))}
    </>
  );
}

export function SummaryTable({ rows, caption }: { rows: { label: string; s: Summary }[]; caption?: string }) {
  return (
    <table className="numbers">
      {caption && <caption>{caption}</caption>}
      <thead>
        <tr>
          <th scope="col">Group</th>
          <th scope="col" className="num">
            n
          </th>
          <th scope="col" className="num">
            Mean
          </th>
          <th scope="col" className="num">
            SD
          </th>
          <th scope="col" className="num">
            95% CI
          </th>
        </tr>
      </thead>
      <tbody>
        <SummaryRows rows={rows} />
      </tbody>
    </table>
  );
}

export function ComparisonNumbers({
  c,
  hypothesis,
  caption,
}: {
  c: Comparison;
  hypothesis?: string;
  caption?: string;
}) {
  return (
    <div className="numbers-block">
      <SummaryTable caption={caption} rows={GROUPS.map((g) => ({ label: GROUP_LABEL[g], s: c[g] }))} />
      <p className="numbers-line">
        Checker minus neutral {fmtSigned(c.diff)} {fmtCi(c.diffCi95)}; {testText("Welch t", c.welchT, c.df, c.p)};
        Hedges g = {fmtNum(c.hedgesG)} <ReadingChip reading={c.reading} hypothesis={hypothesis} />
      </p>
    </div>
  );
}

export function OneSampleLine({ o, label, hypothesis }: { o: OneSample; label: string; hypothesis?: string }) {
  return (
    <p className="numbers-line">
      {label}: n = {o.n}, mean {fmtSigned(o.mean)} {fmtCi(o.ci95)}; {testText("t", o.t, o.df, o.p)}{" "}
      <ReadingChip reading={o.reading} hypothesis={hypothesis} />
    </p>
  );
}

// The paired table folds the test and its reading into one wrapping cell, so it fits a card at 1440 px.
export function PairedTable({ rows, hypothesis }: { rows: { label: string; p: Paired }[]; hypothesis?: string }) {
  return (
    <div className="numbers-wrap">
      <table className="numbers">
        <thead>
          <tr>
            <th scope="col">Group</th>
            <th scope="col" className="num">
              n
            </th>
            <th scope="col" className="num">
              Reasoned, denial
            </th>
            <th scope="col" className="num">
              Difference [95% CI]
            </th>
            <th scope="col">Paired t and reading</th>
          </tr>
        </thead>
        <tbody>
          {rows.map(({ label, p }) => (
            <tr key={label}>
              <th scope="row">{label}</th>
              <td className="num">{p.n}</td>
              <td className="num">
                <span className="nw">{fmtSigned(p.reasoned.mean)},</span>{" "}
                <span className="nw">{fmtSigned(p.denial.mean)}</span>
              </td>
              <td className="num">
                <span className="nw">{fmtSigned(p.meanDiff)}</span> <span className="nw">{fmtCi(p.ci95)}</span>
              </td>
              <td>
                {testText("t", p.t, p.df, p.p)} <ReadingChip reading={p.reading} hypothesis={hypothesis} />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

/** Per-slide comparisons, one row per slide; they carry reading "n_a", so no chip. */
export function PerSlideTable({ rows, caption }: { rows: { label: string; c: Comparison }[]; caption: string }) {
  return (
    <div className="numbers-wrap">
      <table className="numbers">
        <caption>{caption}</caption>
        <thead>
          <tr>
            <th scope="col">Slide</th>
            <th scope="col" className="num">
              n (checker, neutral)
            </th>
            <th scope="col" className="num">
              Checker
            </th>
            <th scope="col" className="num">
              Neutral
            </th>
            <th scope="col" className="num">
              Difference [95% CI]
            </th>
            <th scope="col">Welch t</th>
            <th scope="col" className="num">
              g
            </th>
          </tr>
        </thead>
        <tbody>
          {rows.map(({ label, c }) => (
            <tr key={label}>
              <th scope="row">{label}</th>
              <td className="num">
                {c.checker.n}, {c.neutral.n}
              </td>
              <td className="num">{fmtNum(c.checker.mean)}</td>
              <td className="num">{fmtNum(c.neutral.mean)}</td>
              <td className="num">
                <span className="nw">{fmtSigned(c.diff)}</span> <span className="nw">{fmtCi(c.diffCi95)}</span>
              </td>
              <td>{testText("t", c.welchT, c.df, c.p)}</td>
              <td className="num">{fmtNum(c.hedgesG)}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export function TwoProportionLine({
  t,
  label,
  hypothesis,
  showReading = true,
}: {
  t: TwoProportion;
  label: string;
  hypothesis?: string;
  showReading?: boolean;
}) {
  return (
    <p className="numbers-line">
      {label}: {fmtPct(t.p1)} vs {fmtPct(t.p2)}, difference {fmtPct(t.diff)}; z = {fmtNum(t.z)}, p{" "}
      {fmtP(t.p).startsWith("<") ? fmtP(t.p) : `= ${fmtP(t.p)}`}{" "}
      {showReading && <ReadingChip reading={t.reading} hypothesis={hypothesis} />}
    </p>
  );
}

export function Card({
  title,
  children,
  actions,
  id,
}: {
  title: string;
  children: ReactNode;
  actions?: ReactNode;
  id?: string;
}) {
  return (
    <article className="dash-card" aria-labelledby={id}>
      <header className="dash-card-head">
        <h3 id={id}>{title}</h3>
        {actions}
      </header>
      {children}
    </article>
  );
}
