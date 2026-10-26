# 06 - Analysis dashboard: plan, API contract and admin changes

Written 2026-10-25 from Oliver's request of that day, `CMPT 863 Project Abstract v8.md` (Hypotheses, Measures, Validity) and the data the service already stores (`02 - Architecture.md` sections 5 and 7). Status: contract fixed here, build in progress; see `05 - Progress.md`.

Oliver's words, in short: evaluate each study (run) with a dashboard; clicking a run such as "Pilot 1.4.0" opens it; 40 to 60 participants, so render data points at that scale, for example how trust develops across the configured slides; a pie of participation per country; key numbers as a bar chart (completed, in progress, aborted); allow deleting one participant's data on request; build a "dashboard pilot" run filled by emulating two or three users, so no old test data is mixed in; keep the export; use the Abstract to decide how the hypotheses are checked; consider the background questions (age group and so on) so the data can be analysed without an export.

## 1. What is stored per participant (already today)

| Where | What | Used for |
|---|---|---|
| `sessions/{id}` | run, config version, group, denial slide, status, opened/consent/started/finished times, audio check, reload count, flags, IP, user agent, screen, time zone, action choices | counts, durations, quality, exclusions |
| `sessions/{id}/ratings/{point}` | per rating point: trust 0 to 10, competence, benevolence, integrity 1 to 7, per-item status (answered, missing, timeout), latency | every trust chart and hypothesis |
| `sessions/{id}/answers/actions` | per flagged slide: action id, timed out, who asked (chair or member), answer type (reason or denial) | H4a, H2/H3 split, actions chart |
| `sessions/{id}/answers/budget` | approve or reject, sureness 0 to 10, status | budget panel |
| `sessions/{id}/answers/background` | field, country (ISO 3166-1 alpha-2), age band, job level, education, cyber knowledge 1 to 5, AI use (six steps), each possibly "preferNotToSay"; honeypot | demographics, moderators |
| `sessions/{id}/events/*` | the full event trail | participant drawer, not the dashboard |

Rating points, in plan order for the v1.4 deck: `T0` (baseline after the opening), `T1` (after slide 1), `T2a` (after slide 2, the flag point), `T2b` (after the CISO's answer on slide 2), `T3a`, `T3b`, `T4a`, `T4b`, `T5` (after slide 5). Phases: `baseline`, `afterSlide`, `afterAnswer`. The dashboard reads the points from the run's config, so a deck with more or fewer slides needs no code change.

Nothing new has to be collected for the dashboard. Everything below is computed from these documents.

## 2. How each hypothesis is read from the data

Trust is the 0-to-10 rating; the three METI items are shown beside it and analysed the same way. Each change is the participant's own difference between two ratings (within-subject), compared between groups (between-subject), as the Abstract's Measures section says. Participants with the flag `forced_allocation` (test links) are never analysed. Other flags are exclusion toggles (section 5).

| Hypothesis | Per participant | Per group | Comparison and test shown |
|---|---|---|---|
| **Baseline check** (Validity: "the groups start alike") | `T0` | mean, SD, 95% CI | checker vs neutral, Welch t, Hedges g |
| **Slide 1 check** (Validity: "whether the checker's presence alone moves trust") | `T1 - T0` | mean, CI | checker vs neutral |
| **H1 flag** | for each flagged slide n: `T{n}a - the rating point immediately before it in plan order` (`T2a - T1`, `T3a - T2b`, `T4a - T3b` for v1.4; the Abstract: "the change from the rating just before it"); if that earlier rating is missing (timeout), the delta is missing, no fallback | per slide and pooled over the three flagged slides (mean of the participant's available deltas) | (a) checker delta below zero: one-sample t against 0; (b) checker vs neutral: Welch t on the pooled delta; also the raw level `T{n}a` checker vs neutral |
| **H2 answer** | for the two reasoned slides: the change `T{n}b - T{n}a` and the level `T{n}b` | per group, pooled over the participant's reasoned slides (participant means) | (a) "higher than after the flag": checker change above zero, one-sample t; (b) "higher than in the neutral group": the **level** `T{n}b` checker vs neutral, Welch t (the Abstract compares the groups at the same step; a comparison of changes would favour the group that fell further, and is H4b1). Exploratory, "whether trust rises above the first rating": `T{n}b - T0` within each group, one-sample t (reviewer 2026-10-25; the earlier wording used the pre-flag rating and a between-group test, which is kept as a secondary number) |
| **H3 denial** | for the denial slide: `T{n}b - T{n}a` | per group: denial delta vs the same participant's mean reasoned delta (paired, within subject) | paired t on (reasoned mean - denial) per group; shown per group and for both groups together |
| **H4a participation** | checker: share of flagged slides where the participant chose "ask to explain" or "challenge" (not "say nothing", not timeout); neutral: share of flagged slides with "has a question" | mean share per group; also counts of slides | Welch t on the per-participant shares carries the reading; the two-proportion z on slides is shown as a descriptive number only, because a participant's slides are not independent |
| **H4b amplification (exploratory)** | b1: reasoned delta `T{n}b - T{n}a`; b2: denial delta | checker vs neutral | Welch t for each; shown as a dumbbell |
| **Research question** (Validity: "the change from the first to the last rating answers the research question") | `T_last - T0` (`T5 - T0` for v1.4) | mean, SD, CI per group | checker vs neutral, Welch t, no stated direction |
| **Budget** (Measures: "whether a change in trust reaches an actual decision") | approve or reject, sureness | share approving per group, mean sureness | two-proportion z; approve rate split by last trust rating tercile |
| **Moderators** (cyber knowledge, AI use; H1 and H2) | the participant's pooled H1 delta and H2 delta | by knowledge level 1 to 5 and by AI-use step, both groups | means with CI per level, Spearman correlation per group (descriptive) |

Tests on the dashboard are for reading the pilot and the running study, not the final analysis. The confirmatory analysis stays offline on the export (mixed model: rating ~ point x group with participant intercepts, as planned in the analysis plan to come), so the dashboard shows: n, mean, SD, 95% CI (t-based), the test statistic, p (two-sided, Welch-Satterthwaite df) and Hedges g, and a one-line reading: "supports" (sign as hypothesised and p < .05), "opposite" (p < .05 against the hypothesis), "does not support" (not significant), or "too few" (fewer than 8 contributing participants per group). Per-slide breakdowns show the numbers without a reading, so about twenty uncorrected readings do not become one chance "supports". With 40 to 60 participants (20 to 30 per group) the dashboard can show a medium effect (g about 0.6 to 0.8); smaller effects need the planned sample and the offline model. No multiple-comparison correction on the dashboard; the final analysis plan decides that.

## 3. Charts (forms chosen by the data's job, after the data-visualisation method in the dataviz skill)

Filter row at the top, one row, scoping every chart: **Participants** (finished only, default; or all consented), **Flags** (per-flag include/exclude toggles, section 5), **Group** (both, checker, neutral), and a **Table view** switch that turns every chart into its table twin.

1. **Key numbers**: a KPI row of stat tiles (finished, in progress = started or ongoing, abandoned, blocked, opened without consent, median duration of finished sessions, n per group), then the **status bar chart** Oliver asked for: one horizontal bar per status with the count, one hue, direct labels.
2. **Trust over the study**: line chart, x = rating points in plan order with the slide number and phase under each tick, y = trust 0 to 10; one line per group (checker slot 1 blue, neutral slot 2 orange), 95% CI band per group, individual participants as thin grey trajectories behind (40 to 60 lines read fine at 20% opacity), flagged points marked on the axis, the denial point of each participant drawn as a hollow marker. Hover: crosshair with both means, CI and n. Three **small multiples** beneath for competence, benevolence and integrity on their own 1-to-7 axes (never a second y axis).
3. **Hypothesis panels** H1, H2, H3, H4a, H4b, baseline: each a card with a small chart and the numbers. H1, H2, H4b: dumbbell or paired bars of the mean delta per group with CI whiskers, zero line; H3: per group two bars (reasoned, denial) with CI; H4a: two bars of the share of slides with a question. Below the chart: n per group, mean, CI, test, p, g, and the one-line reading. Per-slide breakdown behind a "by slide" toggle.
4. **Actions per slide**: stacked horizontal bars, one per flagged slide, segments say nothing / ask to explain / challenge / timed out (checker) and has question / no question / timed out (neutral), direct labels where they fit, 2 px gaps.
5. **Budget**: two stacked bars (approve, reject, missing) per group and a dot plot of sureness.
6. **Participation per country**: the donut Oliver asked for, at most six countries plus "Other", direct labels with count and share, legend, and the table twin with every country; countries from the ISO code via a small name map. Part-to-whole only; a bar chart of the same data sits in the table view for close comparisons.
7. **Background**: one bar chart each for age band, field, job level, education, cyber knowledge (ordered, ordinal ramp), AI use (ordered), one hue per chart, "prefer not to say" shown as its own bar, split by group on hover and in the table.
8. **Quality**: flags count table, audio check correct / wrong, rating timeouts per point (bar), reloads, duration distribution (histogram of finished sessions).
9. **Moderators**: two small panels, H1 and H2 pooled delta by cyber knowledge level and by AI-use step, both groups, means with CI.

Every chart: thin marks, hairline grid, legend when two series, tooltips that only add, a table twin, light and dark theme tokens from the reference palette, palette validated with `scripts/validate_palette.js` of the dataviz skill (blue `#2a78d6` and orange `#eb6834` pass adjacent-pair CVD checks in both modes).

## 4. Admin changes

- **Runs tab**: the run name becomes a link; a "Dashboard" button opens the run dashboard in place of the table (with "Back to runs"). URL hash `#run=<id>` so a dashboard can be bookmarked; the login overlay keeps working over it.
- **Run dashboard** (`RunDashboard`): header with run name, code, config version, created, status, test-overrides chip; the filter row; the panels of section 3; "Download this analysis as JSON" (the same document the charts read); link to the Export tab.
- **Participant drawer**: "Delete this participant's data" at the bottom, red. It asks for the session id to be typed (the one shown in the drawer) and a free-text reason ("participant asked by e-mail on ..."), then calls the delete endpoint, closes the drawer and reloads the list. The dashboard excludes deleted sessions at once (they are gone).
- **Configs, Export**: unchanged. The export keeps every column; deleted sessions are absent from it.

## 5. Exclusion toggles (proposal for the analysis plan; Oliver decides the defaults)

| Flag | Default on the dashboard | Why |
|---|---|---|
| `forced_allocation` | always excluded, not a toggle | test links, group chosen by hand |
| `honeypot` | excluded | bot or careless autofill |
| `audio_check_wrong` | excluded | could not hear the clips |
| `many_invalid_events` | excluded | tampering or a broken client |
| `possible_duplicate` | included | same address and browser is common in a household or an office; marked, not blocked (decision 1.21) |
| `rate_limited` | included | same reason; pilots from one network |
| `tab_hidden_during_audio` | included, shown | the participant looked away; a sensitivity check |
| `reloaded` | included | a reload replays the clip, nothing lost |
| rating timeouts | included; the missing rating is simply absent from that point | the dashboard shows how many |

Deleted participants are not an exclusion: their documents are removed and a tombstone without personal data remains (`deleted/{sessionId}`: run id, deleted at, reason), so the count of deletions stays auditable.

## 6. API contract

### `GET /api/admin/runs/{id}/analysis`

Query: `include=finished|consented` (default `finished`), `exclude=<comma-separated flags>` (default `honeypot,audio_check_wrong,many_invalid_events`; `forced_allocation` is always added), `group=both|checker|neutral` (default `both`; filters the per-participant lists, the group statistics are always computed for both groups so the comparison stands).

Response (all numbers plain JSON numbers; `null` where undefined, for example a mean of zero participants):

```ts
interface Analysis {
  run: { id: string; code: string; name: string; source: string; configVersion: string; status: string; createdAt: string; testOverrides: boolean };
  generatedAt: string;
  filters: { include: "finished" | "consented"; exclude: string[]; group: "both" | "checker" | "neutral" };
  config: {
    groups: ["checker", "neutral"];
    slides: { id: string; n: number; title: string; flagged: boolean }[];
    flaggedSlides: string[];
    ratingPoints: { point: string; slideId: string; phase: "baseline" | "afterSlide" | "afterAnswer"; label: string }[]; // plan order
    items: { id: "trust" | "competence" | "benevolence" | "integrity"; label: string; min: number; max: number }[];
    actions: { checker: { id: string; label: string }[]; neutral: { id: string; label: string }[] };
    background: { id: string; label: string; kind: "single" | "scale" | "country"; options?: { id: string; label: string }[]; scale?: { min: number; max: number; pointLabels?: Record<string, string> } }[];
    budgetOptions: { id: string; label: string }[];
  };
  counts: {
    byStatus: { opened: number; started: number; ongoing: number; finished: number; abandoned: number; blocked: number };
    total: number;                       // every session of the run
    consented: number;
    analysed: number;                    // after include and exclude
    excluded: { forced: number; byFlag: Record<string, number>; notIncluded: number };
    byGroup: Record<"checker" | "neutral", { consented: number; finished: number; analysed: number }>;
    durationsSec: { n: number; median: number | null; p25: number | null; p75: number | null; histogram: { fromSec: number; toSec: number; n: number }[] };
  };
  participants: {                        // one per analysed session, for the trajectory lines and the tables
    sessionId: string; group: "checker" | "neutral"; denialSlide: string; status: string; finishedAt: string | null; durationSec: number | null;
    ratings: Record<string, { trust: number | null; competence: number | null; benevolence: number | null; integrity: number | null; status: string } | null>; // by point
    actions: Record<string, { actionId: string | null; timedOut: boolean; questionBy: string; answerType: "reason" | "denial" }>;     // by flagged slide id
    budget: { optionId: string | null; confidence: number | null; status: string } | null;
    background: Record<string, string | number | null> | null;   // item id -> value or "preferNotToSay"
    flags: string[];
  }[];
  ratingStats: Record<string /*item*/, Record<string /*point*/, Record<"checker" | "neutral", Summary>>>;
  hypotheses: {
    baseline: Comparison;                                  // T0 checker vs neutral
    slide1: Comparison;                                    // T1 - T0
    overall: Comparison;                                   // T_last - T0, research question, no direction
    h1: { pooled: Comparison; checkerBelowZero: OneSample; perSlide: Record<string, Comparison>; levelAtFlag: Record<string, Comparison> };   // perSlide and levelAtFlag carry reading "n_a"
    h2: {
      pooled: Comparison;                                  // change T{n}b - T{n}a, checker vs neutral (same data as h4b.afterReason; shown as H4b1)
      checkerAboveZero: OneSample;                         // H2(a)
      level: Comparison;                                   // H2(b): level T{n}b over reasoned slides, checker vs neutral, direction +1; carries the H2 reading
      perSlide: Record<string, Comparison>;                // change per slide, reading "n_a"
      levelPerSlide: Record<string, Comparison>;           // level per slide, reading "n_a"
      aboveBaseline: Record<"checker" | "neutral", OneSample>; // exploratory: T{n}b - T0 within group, direction +1
      exploratoryNet: Comparison;                          // secondary: T{n}b - pre-flag rating, between groups, no direction
    };
    h3: { perGroup: Record<"checker" | "neutral", Paired>; both: Paired };
    h4a: { checker: Share; neutral: Share; slidesTest: TwoProportion; participantsTest: Comparison };   // participantsTest carries the reading; slidesTest reading is "n_a"
    h4b: { afterReason: Comparison; afterDenial: Comparison };
    budget: { byGroup: Record<"checker" | "neutral", { approve: number; reject: number; missing: number; confidence: Summary }>; approveTest: TwoProportion;
              byLastTrustTercile: { tercile: 1 | 2 | 3; from: number; to: number; approve: number; reject: number; missing: number }[] };  // terciles of the participant's last trust rating
    moderators: {
      cyberKnowledge: { level: number; h1: Record<"checker" | "neutral", Summary>; h2: Record<"checker" | "neutral", Summary> }[];
      aiUse: { optionId: string; h1: Record<"checker" | "neutral", Summary>; h2: Record<"checker" | "neutral", Summary> }[];
      spearman: { cyberKnowledgeH1: Record<"checker" | "neutral", number | null>; cyberKnowledgeH2: Record<"checker" | "neutral", number | null>; aiUseH1: Record<"checker" | "neutral", number | null>; aiUseH2: Record<"checker" | "neutral", number | null> };
    };
  };
  actions: Record<string /*slide*/, Record<"checker" | "neutral", Record<string /*actionId or "timeout"*/, number>>>;
  background: Record<string /*item*/, { id: string; label: string; n: number; byGroup: Record<"checker" | "neutral", number> }[]>; // includes "preferNotToSay" and "missing"
  countries: { code: string; n: number; byGroup: Record<"checker" | "neutral", number> }[];           // sorted by n desc
  quality: { flags: Record<string, number>; audioCheck: { correct: number; wrong: number; missing: number }; ratingTimeouts: Record<string /*point*/, number>; reloads: { sessions: number; total: number } };
}
interface Summary   { n: number; mean: number | null; sd: number | null; ci95: [number, number] | null; median: number | null }
interface OneSample { n: number; mean: number | null; ci95: [number, number] | null; t: number | null; df: number | null; p: number | null; reading: Reading }
interface Comparison { checker: Summary; neutral: Summary; diff: number | null; diffCi95: [number, number] | null; welchT: number | null; df: number | null; p: number | null; hedgesG: number | null; reading: Reading }
interface Paired    { n: number; reasoned: Summary; denial: Summary; meanDiff: number | null; ci95: [number, number] | null; t: number | null; df: number | null; p: number | null; reading: Reading }
interface Share     { slides: number; withQuestion: number; share: number | null; perParticipant: Summary }
interface TwoProportion { p1: number | null; p2: number | null; diff: number | null; z: number | null; p: number | null; reading: Reading }
type Reading = "supports" | "opposite" | "does_not_support" | "too_few" | "n_a";   // opposite: p < .05 against the hypothesis; too_few below 8 contributing participants per group
```

Deltas use only ratings whose trust value is present at both points; `n` is the number of participants with that delta. CI and p use Student's t with Welch-Satterthwaite degrees of freedom (one-sample and paired: n-1); Hedges g is Cohen's d with the small-sample correction. Readings: "supports" when the sign matches the hypothesis and p < .05; "opposite" when p < .05 with the other sign; "does_not_support" otherwise; "too_few" when either group has fewer than 8 contributing participants; "n_a" where no direction is stated or the value is undefined. Zero variance with a non-zero mean gives t and p null and a reading from the sign is not attempted ("n_a").

### `DELETE /api/admin/sessions/{id}`

Body `{ "confirm": "<sessionId>", "reason": "<text, 1 to 500 characters>" }`. Deletes `sessions/{id}` and its `events`, `ratings` and `answers`; writes `deleted/{id}` = `{ runId, group, deletedAt, reason }` (no address, agent or answers). 204; 404 if unknown; 400 if `confirm` does not match. The run's counters are untouched (the slot stays used, decision 1.6). Logged at info level with the admin's address.

## 7. Dashboard pilot with emulated participants

A run "Dashboard pilot (emulated)" on config `1.4.0`, source `Emulated`, filled by the load generator in a new `-persona` mode: each emulated participant has a baseline, a group-typical response to flags, reasons and denials with noise, a coherent METI profile, a chosen action pattern, a budget vote that follows the last trust rating, and a plausible background (country, age band, field, job level, education, knowledge, AI use; 5% "prefer not to say"). Emulated sessions come from one address, so they carry `possible_duplicate`, which the dashboard includes by default. Three participants as Oliver asked, which makes a first dashboard readable; the same mode can fill 40 in one go to see the charts at study scale (the per-group readings will then show "too few" or real tests as the numbers grow). Emulated runs are named so in the admin and are never mixed with recruitment runs.

## 8. Open for Oliver

1. The exclusion defaults in section 5.
2. Whether the dashboard should also show the METI items in the hypothesis panels (now: trust only in the panels, items in the trajectory small multiples).
3. The reading thresholds (p < .05, fewer than 8 per group) are placeholders for the analysis plan.
4. Whether to keep deleted-session tombstones or delete without trace.
5. The exploratory H2 reference point: the Abstract's "first rating" is read as T0 (within group); say if the pre-flag rating is meant instead.
6. H1 with a missing earlier rating: the delta is dropped (no fallback to an older point). Confirm.
7. Whether the research-question comparison (T_last - T0) should carry a direction.
