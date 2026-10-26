import { useState } from "react";
import type { Comparison, RunAnalysis } from "../api/types";
import { BarChart } from "../charts/BarChart";
import { Dumbbell, type DumbbellRow, type DumbbellSeries } from "../charts/Dumbbell";
import { fmtCi, fmtNum, fmtPct } from "../charts/format";
import { GROUPS, GROUP_COLOR, GROUP_LABEL, ciPoint, slideLabel } from "./model";
import {
  Card,
  ComparisonNumbers,
  OneSampleLine,
  ReadingChip,
  PairedTable,
  PerSlideTable,
  SummaryTable,
  TwoProportionLine,
  testText,
} from "./Numbers";

const GROUP_SERIES: DumbbellSeries[] = GROUPS.map((g) => ({ key: g, label: GROUP_LABEL[g], color: GROUP_COLOR[g] }));

function compRow(id: string, label: string, c: Comparison): DumbbellRow {
  return { id, label, points: { checker: ciPoint(c.checker), neutral: ciPoint(c.neutral) } };
}

function BySlideToggle({ on, set }: { on: boolean; set: (v: boolean) => void }) {
  return (
    <label className="check small-text">
      <input type="checkbox" checked={on} onChange={(e) => set(e.target.checked)} />
      By slide
    </label>
  );
}

export function HypothesisCards({ a, tableView }: { a: RunAnalysis; tableView: boolean }) {
  const h = a.hypotheses;
  const [h1BySlide, setH1BySlide] = useState(false);
  const [h2BySlide, setH2BySlide] = useState(false);
  const slides = a.config.flaggedSlides;
  const lastPoint = a.config.ratingPoints[a.config.ratingPoints.length - 1]?.point ?? "last";
  const reasonedSlides = Object.keys(h.h2.perSlide);

  const h1Rows: DumbbellRow[] = [
    compRow("pooled", "Pooled", h.h1.pooled),
    ...(h1BySlide
      ? slides.filter((s) => h.h1.perSlide[s]).map((s) => compRow(s, slideLabel(a, s), h.h1.perSlide[s]))
      : []),
  ];
  const h2LevelRows: DumbbellRow[] = [
    compRow("level", "Pooled", h.h2.level),
    ...(h2BySlide
      ? reasonedSlides
          .filter((s) => h.h2.levelPerSlide?.[s])
          .map((s) => compRow(s, slideLabel(a, s), h.h2.levelPerSlide[s]))
      : []),
  ];
  const h2Rows: DumbbellRow[] = [
    compRow("pooled", "Pooled", h.h2.pooled),
    ...(h2BySlide
      ? reasonedSlides.filter((s) => h.h2.perSlide[s]).map((s) => compRow(s, slideLabel(a, s), h.h2.perSlide[s]))
      : []),
    compRow("net", "Net (exploratory)", h.h2.exploratoryNet),
  ];

  // DECISION: H3 rows are the groups (coloured by group, "both" in grey); the two points are the reasoned mean
  // (filled) and the denial (hollow, as in the trajectory chart), so the series differ by shape, not by a third hue.
  const h3Series: DumbbellSeries[] = [
    { key: "reasoned", label: "After a reason", color: "var(--viz-ink-2)" },
    { key: "denial", label: "After the denial", color: "var(--viz-ink-2)", hollow: true },
  ];
  const h3Rows: DumbbellRow[] = [
    ...GROUPS.map((g) => ({
      id: g,
      label: GROUP_LABEL[g],
      color: GROUP_COLOR[g],
      points: { reasoned: ciPoint(h.h3.perGroup[g].reasoned), denial: ciPoint(h.h3.perGroup[g].denial) },
    })),
    {
      id: "both",
      label: "Both groups",
      color: "var(--viz-deemph)",
      points: { reasoned: ciPoint(h.h3.both.reasoned), denial: ciPoint(h.h3.both.denial) },
    },
  ];

  return (
    <div className="dash-grid">
      <Card title="First to last rating" id="hyp-overall">
        <p className="muted small-text">
          Trust at the last rating point minus T0: how far trust moved over the whole meeting (research question, no
          direction).
        </p>
        <Dumbbell
          title="Change in trust from the first to the last rating by group"
          series={GROUP_SERIES}
          rows={[compRow("overall", `${lastPoint} - T0`, h.overall)]}
          tableView={tableView}
        />
        <ComparisonNumbers c={h.overall} />
      </Card>

      <Card title="Baseline check" id="hyp-baseline">
        <p className="muted small-text">Trust at T0. The groups should start alike.</p>
        <Dumbbell
          title="Baseline trust by group"
          series={GROUP_SERIES}
          rows={[compRow("t0", "T0", h.baseline)]}
          zeroLine={false}
          tableView={tableView}
        />
        <ComparisonNumbers c={h.baseline} />
      </Card>

      <Card title="Slide 1 check" id="hyp-slide1">
        <p className="muted small-text">T1 minus T0: does the checker's presence alone move trust?</p>
        <Dumbbell
          title="Change in trust after slide 1 by group"
          series={GROUP_SERIES}
          rows={[compRow("d1", "T1 - T0", h.slide1)]}
          tableView={tableView}
        />
        <ComparisonNumbers c={h.slide1} />
      </Card>

      <Card title="H1 flag" id="hyp-h1" actions={<BySlideToggle on={h1BySlide} set={setH1BySlide} />}>
        <p className="muted small-text">
          Change in trust at each flag (after the slide minus the rating before it). H1: the checker group drops, and
          more than the neutral group.
        </p>
        <Dumbbell
          title="Change in trust at the flags by group"
          series={GROUP_SERIES}
          rows={h1Rows}
          tableView={tableView}
        />
        <ComparisonNumbers c={h.h1.pooled} hypothesis="H1" caption="Pooled over the flagged slides" />
        <OneSampleLine o={h.h1.checkerBelowZero} label="Checker below zero" hypothesis="H1" />
        {h1BySlide && (
          <>
            <PerSlideTable
              caption="Change at each flag"
              rows={slides
                .filter((s) => h.h1.perSlide[s])
                .map((s) => ({ label: slideLabel(a, s), c: h.h1.perSlide[s] }))}
            />
            <PerSlideTable
              caption="Trust level right after each flag"
              rows={slides
                .filter((s) => h.h1.levelAtFlag[s])
                .map((s) => ({ label: slideLabel(a, s), c: h.h1.levelAtFlag[s] }))}
            />
          </>
        )}
      </Card>

      <Card title="H2 answer" id="hyp-h2" actions={<BySlideToggle on={h2BySlide} set={setH2BySlide} />}>
        <p className="muted small-text">
          H2: after a reasoned answer the checker group recovers (a) and ends higher than the neutral group (b). The
          headline is (b), the trust level after the answer.
        </p>
        <h4>Trust after the reasoned answer</h4>
        <Dumbbell
          title="Trust level after a reasoned answer by group"
          series={GROUP_SERIES}
          rows={h2LevelRows}
          zeroLine={false}
          tableView={tableView}
        />
        <ComparisonNumbers
          c={h.h2.level}
          hypothesis="H2"
          caption="H2(b): level after the reasoned answer, between groups"
        />
        <OneSampleLine
          o={h.h2.checkerAboveZero}
          label="H2(a): checker change after the reason above zero"
          hypothesis="H2"
        />
        <h4>Change after the reasoned answer</h4>
        <Dumbbell
          title="Change in trust after a reasoned answer by group"
          series={GROUP_SERIES}
          rows={h2Rows}
          tableView={tableView}
        />
        <ComparisonNumbers c={h.h2.pooled} caption="Change after the reason, between groups (also H4b1)" />
        {GROUPS.map((g) =>
          h.h2.aboveBaseline?.[g] ? (
            <OneSampleLine
              key={g}
              o={h.h2.aboveBaseline[g]}
              label={`${GROUP_LABEL[g]} above the first rating (exploratory)`}
            />
          ) : null,
        )}
        <p className="numbers-line muted">
          Exploratory net change (after the answer minus before the flag): checker{" "}
          {fmtNum(h.h2.exploratoryNet.checker.mean)}, neutral {fmtNum(h.h2.exploratoryNet.neutral.mean)};{" "}
          {testText("Welch t", h.h2.exploratoryNet.welchT, h.h2.exploratoryNet.df, h.h2.exploratoryNet.p)}
        </p>
        {h2BySlide && (
          <>
            <PerSlideTable
              caption="Level after the answer on each reasoned slide"
              rows={reasonedSlides
                .filter((s) => h.h2.levelPerSlide?.[s])
                .map((s) => ({ label: slideLabel(a, s), c: h.h2.levelPerSlide[s] }))}
            />
            <PerSlideTable
              caption="Change after the answer on each reasoned slide"
              rows={reasonedSlides
                .filter((s) => h.h2.perSlide[s])
                .map((s) => ({ label: slideLabel(a, s), c: h.h2.perSlide[s] }))}
            />
          </>
        )}
      </Card>

      <Card title="H3 denial" id="hyp-h3">
        <p className="muted small-text">
          Change in trust after the denial compared with the same participant's mean change after a reason. H3: trust
          recovers less after a denial.
        </p>
        <Dumbbell
          title="Recovery after a reason and after the denial"
          series={h3Series}
          rows={h3Rows}
          tableView={tableView}
        />
        <PairedTable
          hypothesis="H3"
          rows={[
            ...GROUPS.map((g) => ({ label: GROUP_LABEL[g], p: h.h3.perGroup[g] })),
            { label: "Both groups", p: h.h3.both },
          ]}
        />
      </Card>

      <Card title="H4a participation" id="hyp-h4a">
        <p className="muted small-text">
          Share of flagged slides where the participant spoke up (checker: asked to explain or challenged; neutral: had
          a question). H4a: the checker group speaks up more often.
        </p>
        <BarChart
          title="Share of flagged slides with a question by group"
          orientation="vertical"
          height={120}
          domainMax={1}
          format={fmtPct}
          valueLabel="Share of slides"
          data={GROUPS.map((g) => ({
            id: g,
            label: GROUP_LABEL[g],
            value: h.h4a[g].share,
            color: GROUP_COLOR[g],
            tip: [{ value: `${h.h4a[g].withQuestion} of ${h.h4a[g].slides}`, label: "slides" }],
            extra: { slides: h.h4a[g].slides, withQuestion: h.h4a[g].withQuestion },
          }))}
          extraColumns={[
            { key: "withQuestion", label: "Slides with a question", numeric: true },
            { key: "slides", label: "Flagged slides", numeric: true },
          ]}
          legend={GROUPS.map((g) => ({ id: g, label: GROUP_LABEL[g], color: GROUP_COLOR[g], kind: "rect" }))}
          tableView={tableView}
        />
        <SummaryTable
          caption="Per participant share of slides"
          rows={GROUPS.map((g) => ({ label: GROUP_LABEL[g], s: h.h4a[g].perParticipant }))}
        />
        <p className="numbers-line">
          Participants, checker vs neutral:{" "}
          {testText("Welch t", h.h4a.participantsTest.welchT, h.h4a.participantsTest.df, h.h4a.participantsTest.p)};
          Hedges g = {fmtNum(h.h4a.participantsTest.hedgesG)}; difference {fmtCi(h.h4a.participantsTest.diffCi95)}{" "}
          <ReadingChip reading={h.h4a.participantsTest.reading} hypothesis="H4a" />
        </p>
        <TwoProportionLine
          t={h.h4a.slidesTest}
          label="Slides, checker vs neutral (slides are not independent)"
          showReading={false}
        />
      </Card>

      <Card title="H4b amplification (exploratory)" id="hyp-h4b">
        <p className="muted small-text">
          Do the checker's flags amplify the answer? After a reason the checker group should recover more; after the
          denial it should fall further. Change after the answer, checker vs neutral.
        </p>
        <Dumbbell
          title="Change after the answer by group"
          series={GROUP_SERIES}
          rows={[
            compRow("reason", "After a reason", h.h4b.afterReason),
            compRow("denial", "After the denial", h.h4b.afterDenial),
          ]}
          tableView={tableView}
        />
        <ComparisonNumbers c={h.h4b.afterReason} caption="After a reason" hypothesis="H4b" />
        <ComparisonNumbers c={h.h4b.afterDenial} caption="After the denial" hypothesis="H4b" />
      </Card>
    </div>
  );
}
