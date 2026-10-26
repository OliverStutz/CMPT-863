// Pure transforms from the analysis document (RunAnalysis) to chart data. No React here, so they are unit-tested.
import type {
  AnalysisGroupFilter,
  AnalysisInclude,
  Group,
  LastTrustTercile,
  Reading,
  RunAnalysis,
  Summary,
} from "../api/types";
import type { BarDatum } from "../charts/BarChart";
import type { CiPoint } from "../charts/Dumbbell";
import type { LineSeries, Trajectory, XPoint } from "../charts/LineChart";
import type { Slice } from "../charts/Donut";
import type { StackKey, StackRow } from "../charts/StackedBar";
import type { Chip } from "../lib/status";
import { shortId } from "../lib/time";

export const GROUPS: Group[] = ["checker", "neutral"];
export const GROUP_LABEL: Record<Group, string> = { checker: "Checker", neutral: "Neutral" };
export const GROUP_COLOR: Record<Group, string> = { checker: "var(--viz-checker)", neutral: "var(--viz-neutral)" };

// ---- filters (plan sections 3 and 5) ----

export interface FlagToggle {
  flag: string;
  excludedByDefault: boolean;
  why: string;
}

/** Plan section 5; forced_allocation is always excluded by the server and is no toggle. */
export const FLAG_TOGGLES: FlagToggle[] = [
  { flag: "honeypot", excludedByDefault: true, why: "bot or careless autofill" },
  { flag: "audio_check_wrong", excludedByDefault: true, why: "could not hear the clips" },
  { flag: "many_invalid_events", excludedByDefault: true, why: "tampering or a broken client" },
  { flag: "possible_duplicate", excludedByDefault: false, why: "same address and browser, common in a household" },
  { flag: "rate_limited", excludedByDefault: false, why: "pilots from one network" },
  { flag: "tab_hidden_during_audio", excludedByDefault: false, why: "looked away; a sensitivity check" },
  { flag: "reloaded", excludedByDefault: false, why: "a reload replays the clip, nothing lost" },
];

export interface DashboardFilters {
  include: AnalysisInclude;
  exclude: string[];
  group: AnalysisGroupFilter;
}

export const DEFAULT_FILTERS: DashboardFilters = {
  include: "finished",
  exclude: FLAG_TOGGLES.filter((t) => t.excludedByDefault).map((t) => t.flag),
  group: "both",
};

export function toggleExclude(f: DashboardFilters, flag: string, excluded: boolean): DashboardFilters {
  const set = new Set(f.exclude);
  if (excluded) set.add(flag);
  else set.delete(flag);
  // Keep the toggle order, so the query string is stable.
  return { ...f, exclude: FLAG_TOGGLES.map((t) => t.flag).filter((x) => set.has(x)) };
}

// ---- readings ----

/**
 * The one-line reading as a chip: supports = success, opposite = danger, does_not_support = neutral,
 * too_few = info, n_a = muted.
 */
export function readingChip(reading: Reading | null | undefined, hypothesis?: string): Chip {
  switch (reading) {
    case "supports":
      return { label: hypothesis ? `supports ${hypothesis} so far` : "supports so far", tone: "success" };
    case "opposite":
      return { label: "opposite", tone: "danger" };
    case "does_not_support":
      return { label: hypothesis ? `does not support ${hypothesis}` : "does not support", tone: "neutral" };
    case "too_few":
      return { label: "too few participants", tone: "info" };
    default:
      return { label: "no directional test", tone: "muted" };
  }
}

// ---- summaries ----

export function ciPoint(s: Summary | null | undefined): CiPoint | undefined {
  if (!s) return undefined;
  return { mean: s.mean, lo: s.ci95?.[0] ?? null, hi: s.ci95?.[1] ?? null, n: s.n };
}

export function slideNumber(a: RunAnalysis, slideId: string): number | null {
  return a.config.slides.find((s) => s.id === slideId)?.n ?? null;
}

export function slideLabel(a: RunAnalysis, slideId: string): string {
  const n = slideNumber(a, slideId);
  return n === null ? slideId : `Slide ${n}`;
}

// ---- trust over the study ----

/** X axis of the trajectory charts: rating points in plan order, slide number and phase under each tick. */
export function ratingXs(a: RunAnalysis): XPoint[] {
  const flagged = new Set(a.config.flaggedSlides);
  return a.config.ratingPoints.map((rp) => {
    const n = slideNumber(a, rp.slideId);
    const sub =
      rp.phase === "baseline"
        ? "baseline"
        : rp.phase === "afterAnswer"
          ? `answer ${n ?? ""}`.trim()
          : `slide ${n ?? ""}`.trim();
    return { id: rp.point, label: rp.point, sub, marked: rp.phase === "afterSlide" && flagged.has(rp.slideId) };
  });
}

/** One series per group: the mean of `item` at each point with its 95% CI and n. */
export function groupSeries(a: RunAnalysis, item: string): LineSeries[] {
  const stats = a.ratingStats[item] ?? {};
  return GROUPS.map((g) => {
    const points: LineSeries["points"] = {};
    for (const rp of a.config.ratingPoints) {
      const s = stats[rp.point]?.[g];
      points[rp.point] = s ? { y: s.mean, lo: s.ci95?.[0] ?? null, hi: s.ci95?.[1] ?? null, n: s.n } : undefined;
    }
    return { id: g, label: GROUP_LABEL[g], color: GROUP_COLOR[g], points };
  });
}

/** The point that follows the answer on the participant's denial slide, or null. */
export function denialPoint(a: RunAnalysis, denialSlide: string): string | null {
  return a.config.ratingPoints.find((rp) => rp.slideId === denialSlide && rp.phase === "afterAnswer")?.point ?? null;
}

/** Each analysed participant's ratings of `item`, with the denial point marked for a hollow marker. */
export function trajectories(a: RunAnalysis, item: string, group: AnalysisGroupFilter = "both"): Trajectory[] {
  return a.participants
    .filter((p) => group === "both" || p.group === group)
    .map((p) => {
      const values: Trajectory["values"] = {};
      for (const rp of a.config.ratingPoints) {
        const r = p.ratings[rp.point];
        values[rp.point] = r ? ((r as unknown as Record<string, number | null>)[item] ?? null) : null;
      }
      const dp = denialPoint(a, p.denialSlide);
      return { id: `${shortId(p.sessionId)} (${p.group})`, values, hollow: dp ? [dp] : [] };
    });
}

// ---- status ----

export const STATUS_ORDER: { key: keyof RunAnalysis["counts"]["byStatus"]; label: string }[] = [
  { key: "opened", label: "Opened, no consent" },
  { key: "started", label: "Started" },
  { key: "ongoing", label: "Ongoing" },
  { key: "finished", label: "Finished" },
  { key: "abandoned", label: "Abandoned" },
  { key: "blocked", label: "Blocked" },
];

export function statusBars(a: RunAnalysis): BarDatum[] {
  return STATUS_ORDER.map((s) => ({ id: s.key, label: s.label, value: a.counts.byStatus[s.key] ?? 0 }));
}

// ---- background ----

export const BACKGROUND_TITLE: Record<string, string> = {
  age: "Age",
  field: "Field of expertise",
  jobLevel: "Job level",
  education: "Education",
  cyberKnowledge: "Cyber security knowledge",
  aiUse: "AI use",
  country: "Country",
};

// DECISION: age bands, cyber knowledge and AI use are ordered scales and take the ordinal ramp; field, job level
// and education are read as nominal (education has "Other") and take the single hue.
export const ORDINAL_ITEMS = new Set(["age", "cyberKnowledge", "aiUse"]);
export const NOT_AN_ANSWER = new Set(["preferNotToSay", "missing"]);

/** Ordinal step i of k on the six-step ramp, spread so neighbours stay at least one validated step apart. */
export function ordinalColor(i: number, k: number): string {
  const step = k <= 1 ? 6 : 1 + Math.round((i * 5) / (k - 1));
  return `var(--viz-ord-${step})`;
}

/** Bars of one background item in config order; "prefer not to say" is its own grey bar, "missing" only when > 0. */
export function backgroundBars(a: RunAnalysis, itemId: string): BarDatum[] {
  const rows = (a.background[itemId] ?? []).filter((r) => r.id !== "missing" || r.n > 0);
  const answers = rows.filter((r) => !NOT_AN_ANSWER.has(r.id));
  const ordinal = ORDINAL_ITEMS.has(itemId);
  return rows.map((r) => {
    const i = answers.indexOf(r);
    return {
      id: r.id,
      label: r.label,
      value: r.n,
      color: NOT_AN_ANSWER.has(r.id) ? "var(--viz-other)" : ordinal ? ordinalColor(i, answers.length) : undefined,
      tip: GROUPS.map((g) => ({ color: GROUP_COLOR[g], value: String(r.byGroup[g] ?? 0), label: GROUP_LABEL[g] })),
      extra: { checker: r.byGroup.checker ?? 0, neutral: r.byGroup.neutral ?? 0 },
    };
  });
}

// ---- countries ----

const COUNTRY_NAMES: Record<string, string> = {
  AR: "Argentina",
  AT: "Austria",
  AU: "Australia",
  BE: "Belgium",
  BR: "Brazil",
  CA: "Canada",
  CH: "Switzerland",
  CL: "Chile",
  CN: "China",
  CO: "Colombia",
  CZ: "Czechia",
  DE: "Germany",
  DK: "Denmark",
  EG: "Egypt",
  ES: "Spain",
  FI: "Finland",
  FR: "France",
  GB: "United Kingdom",
  GR: "Greece",
  HK: "Hong Kong",
  HU: "Hungary",
  ID: "Indonesia",
  IE: "Ireland",
  IL: "Israel",
  IN: "India",
  IR: "Iran",
  IT: "Italy",
  JP: "Japan",
  KE: "Kenya",
  KR: "South Korea",
  LU: "Luxembourg",
  MX: "Mexico",
  MY: "Malaysia",
  NG: "Nigeria",
  NL: "Netherlands",
  NO: "Norway",
  NZ: "New Zealand",
  PH: "Philippines",
  PK: "Pakistan",
  PL: "Poland",
  PT: "Portugal",
  RO: "Romania",
  RU: "Russia",
  SA: "Saudi Arabia",
  SE: "Sweden",
  SG: "Singapore",
  TH: "Thailand",
  TR: "Turkey",
  TW: "Taiwan",
  UA: "Ukraine",
  AE: "United Arab Emirates",
  US: "United States",
  VN: "Vietnam",
  ZA: "South Africa",
};

export function countryName(code: string): string {
  if (code === "preferNotToSay") return "Prefer not to say";
  if (code === "missing" || !code) return "Missing";
  return COUNTRY_NAMES[code.toUpperCase()] ?? code.toUpperCase();
}

export function countrySlices(a: RunAnalysis): Slice[] {
  return a.countries
    .filter((c) => /^[A-Za-z]{2}$/.test(c.code))
    .map((c) => ({
      id: c.code,
      label: countryName(c.code),
      value: c.n,
      extra: { code: c.code, checker: c.byGroup.checker ?? 0, neutral: c.byGroup.neutral ?? 0 },
    }));
}

// ---- actions ----

const TIMEOUT_KEY = { key: "timeout", label: "Timed out", color: "var(--viz-other)", ink: "var(--viz-on-other)" };

// DECISION: the action segments are ordered by how much the participant speaks up, so they take the ordinal ramp
// (light = says nothing, dark = challenges) and timeouts the "Other" grey; the group hues stay reserved for groups.
export function actionKeys(a: RunAnalysis, g: Group): StackKey[] {
  const opts = a.config.actions[g];
  const ladder: Record<string, number> =
    g === "checker" ? { nothing: 2, explain: 4, challenge: 6 } : { hasQuestion: 6, noQuestion: 2 };
  const keys = opts.map((o, i) => {
    const step = ladder[o.id] ?? 1 + Math.round((i * 5) / Math.max(1, opts.length - 1));
    return { key: o.id, label: o.label, color: `var(--viz-ord-${step})`, ink: `var(--viz-on-ord-${step})` };
  });
  return [...keys, TIMEOUT_KEY];
}

export function actionRows(a: RunAnalysis, g: Group): StackRow[] {
  return a.config.flaggedSlides.map((sid) => ({
    id: sid,
    label: slideLabel(a, sid),
    values: a.actions[sid]?.[g] ?? {},
  }));
}

// ---- budget ----

// DECISION: approve and reject are opposite poles, so they take the diverging pair (blue, red) of palette.md and
// "no answer" the Other grey; the rows name the group, so the blue here is not read as "checker".
export function budgetKeys(a: RunAnalysis): StackKey[] {
  const label = (id: string, fallback: string) => a.config.budgetOptions.find((o) => o.id === id)?.label ?? fallback;
  return [
    { key: "approve", label: label("approve", "Approve"), color: "var(--viz-diverge-pos)" },
    { key: "reject", label: label("reject", "Reject"), color: "var(--viz-diverge-neg)" },
    { key: "missing", label: "No answer", color: "var(--viz-other)", ink: "var(--viz-on-other)" },
  ];
}

export function budgetRows(a: RunAnalysis): StackRow[] {
  return GROUPS.map((g) => {
    const b = a.hypotheses.budget.byGroup[g];
    return { id: g, label: GROUP_LABEL[g], values: { approve: b.approve, reject: b.reject, missing: b.missing } };
  });
}

/** Last trust rating present per participant (plan order), for the approve rate by tercile (plan section 2). */
export function lastTrust(a: RunAnalysis, p: RunAnalysis["participants"][number]): number | null {
  for (let i = a.config.ratingPoints.length - 1; i >= 0; i--) {
    const v = p.ratings[a.config.ratingPoints[i].point]?.trust;
    if (v !== null && v !== undefined) return v;
  }
  return null;
}

export interface TercileRow {
  id: string;
  label: string;
  approve: number;
  answered: number;
  byGroup: Record<Group, { approve: number; answered: number }>;
}

// DECISION: the server sends budget.byLastTrustTercile (ties go to the lower tercile; from/to are the rating
// bounds). Until every server does, the same split is computed here from the participant list as a fallback.
export function lastTrustTerciles(a: RunAnalysis): LastTrustTercile[] {
  if (a.hypotheses.budget.byLastTrustTercile) return a.hypotheses.budget.byLastTrustTercile;
  const rows = a.participants.map((p) => ({ p, t: lastTrust(a, p) })).filter((r) => r.t !== null) as {
    p: RunAnalysis["participants"][number];
    t: number;
  }[];
  if (rows.length === 0) return [];
  const sorted = rows.map((r) => r.t).sort((x, y) => x - y);
  const q = (f: number) => {
    const pos = (sorted.length - 1) * f;
    const lo = Math.floor(pos);
    return sorted[lo] + (sorted[Math.ceil(pos)] - sorted[lo]) * (pos - lo);
  };
  const c1 = q(1 / 3);
  const c2 = q(2 / 3);
  const tercileOf = (t: number): 1 | 2 | 3 => (t <= c1 ? 1 : t <= c2 ? 2 : 3);
  const out: LastTrustTercile[] = [];
  for (const tercile of [1, 2, 3] as const) {
    const inT = rows.filter((r) => tercileOf(r.t) === tercile);
    if (inT.length === 0) continue;
    const ts = inT.map((r) => r.t);
    out.push({
      tercile,
      from: Math.min(...ts),
      to: Math.max(...ts),
      approve: inT.filter((r) => r.p.budget?.optionId === "approve").length,
      reject: inT.filter((r) => r.p.budget?.optionId === "reject").length,
      missing: inT.filter((r) => !r.p.budget || r.p.budget.optionId === null).length,
    });
  }
  return out;
}

export function tercileRows(a: RunAnalysis): StackRow[] {
  return lastTrustTerciles(a).map((t) => ({
    id: `t${t.tercile}`,
    label: `${["Low", "Middle", "High"][t.tercile - 1]} (trust ${t.from === t.to ? t.from : `${t.from} to ${t.to}`})`,
    values: { approve: t.approve, reject: t.reject, missing: t.missing },
  }));
}

/** "Prefer not to say" and "missing" for the country question; they are in background.country, not in countries. */
export function countryNonAnswers(a: RunAnalysis): { preferNotToSay: number; missing: number } {
  const rows = a.background.country ?? [];
  return {
    preferNotToSay: rows.find((r) => r.id === "preferNotToSay")?.n ?? 0,
    missing: rows.find((r) => r.id === "missing")?.n ?? 0,
  };
}

// ---- moderators ----

export function cyberLevelLabel(a: RunAnalysis, level: number): string {
  const item = a.config.background.find((b) => b.id === "cyberKnowledge");
  const pl = item?.scale?.pointLabels?.[String(level)];
  return pl ? `${level} ${pl}` : String(level);
}

export function optionLabel(a: RunAnalysis, itemId: string, optionId: string): string {
  return a.config.background.find((b) => b.id === itemId)?.options?.find((o) => o.id === optionId)?.label ?? optionId;
}

// ---- the preview (no API) ----

/** Applies the group filter to the participant list, as the server does; the group statistics stay as they are. */
export function filterParticipants(a: RunAnalysis, group: AnalysisGroupFilter): RunAnalysis {
  if (group === "both") return { ...a, filters: { ...a.filters, group } };
  return { ...a, filters: { ...a.filters, group }, participants: a.participants.filter((p) => p.group === group) };
}
