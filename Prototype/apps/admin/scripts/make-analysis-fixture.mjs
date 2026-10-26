// Deterministic fixture for the run dashboard: GET /api/admin/runs/{id}/analysis, plan "06 - Analysis Dashboard Plan.md"
// section 6. Emulated participants with plausible effects; every summary is computed here from the participant
// records, following section 2 and the notes under section 6.
// Run from apps/admin: node scripts/make-analysis-fixture.mjs
import { mkdirSync, writeFileSync } from "node:fs";
import { dirname, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const here = dirname(fileURLToPath(import.meta.url));
const OUT = resolve(here, "../src/dashboard/fixtures/analysis.sample.json");

// ---------- seeded randomness ----------
function mulberry32(seed) {
  let a = seed >>> 0;
  return () => {
    a = (a + 0x6d2b79f5) >>> 0;
    let t = a;
    t = Math.imul(t ^ (t >>> 15), t | 1);
    t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
    return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
  };
}
const rand = mulberry32(863);
function normal(mean = 0, sd = 1) {
  const u = 1 - rand();
  const v = rand();
  return mean + sd * Math.sqrt(-2 * Math.log(u)) * Math.cos(2 * Math.PI * v);
}
const chance = (p) => rand() < p;
function pick(list, weights) {
  const total = weights.reduce((a, b) => a + b, 0);
  let r = rand() * total;
  for (let i = 0; i < list.length; i++) {
    r -= weights[i];
    if (r < 0) return list[i];
  }
  return list[list.length - 1];
}
function shuffle(list) {
  const a = list.slice();
  for (let i = a.length - 1; i > 0; i--) {
    const j = Math.floor(rand() * (i + 1));
    [a[i], a[j]] = [a[j], a[i]];
  }
  return a;
}
const hex = (n) => Array.from({ length: n }, () => Math.floor(rand() * 16).toString(16)).join("");
const clamp = (x, lo, hi) => Math.min(hi, Math.max(lo, x));

// ---------- statistics ----------
function lgamma(x) {
  const c = [
    676.5203681218851, -1259.1392167224028, 771.32342877765313, -176.61502916214059, 12.507343278686905,
    -0.13857109526572012, 9.9843695780195716e-6, 1.5056327351493116e-7,
  ];
  if (x < 0.5) return Math.log(Math.PI / Math.sin(Math.PI * x)) - lgamma(1 - x);
  x -= 1;
  let a = 0.99999999999980993;
  const t = x + 7.5;
  for (let i = 0; i < 8; i++) a += c[i] / (x + i + 1);
  return 0.5 * Math.log(2 * Math.PI) + (x + 0.5) * Math.log(t) - t + Math.log(a);
}
// Regularized incomplete beta I_x(a, b), continued fraction (Numerical Recipes betacf).
function betacf(a, b, x) {
  const EPS = 1e-14;
  const FPMIN = 1e-300;
  let c = 1;
  let d = 1 - ((a + b) * x) / (a + 1);
  if (Math.abs(d) < FPMIN) d = FPMIN;
  d = 1 / d;
  let h = d;
  for (let m = 1; m <= 300; m++) {
    const m2 = 2 * m;
    let aa = (m * (b - m) * x) / ((a + m2 - 1) * (a + m2));
    d = 1 + aa * d;
    if (Math.abs(d) < FPMIN) d = FPMIN;
    c = 1 + aa / c;
    if (Math.abs(c) < FPMIN) c = FPMIN;
    d = 1 / d;
    h *= d * c;
    aa = (-(a + m) * (a + b + m) * x) / ((a + m2) * (a + m2 + 1));
    d = 1 + aa * d;
    if (Math.abs(d) < FPMIN) d = FPMIN;
    c = 1 + aa / c;
    if (Math.abs(c) < FPMIN) c = FPMIN;
    d = 1 / d;
    const del = d * c;
    h *= del;
    if (Math.abs(del - 1) < EPS) break;
  }
  return h;
}
function ibeta(x, a, b) {
  if (x <= 0) return 0;
  if (x >= 1) return 1;
  const bt = Math.exp(lgamma(a + b) - lgamma(a) - lgamma(b) + a * Math.log(x) + b * Math.log(1 - x));
  return x < (a + 1) / (a + b + 2) ? (bt * betacf(a, b, x)) / a : 1 - (bt * betacf(b, a, 1 - x)) / b;
}
/** Two-sided p of Student's t with df degrees of freedom. */
function tTwoSided(t, df) {
  if (!Number.isFinite(t) || !(df > 0)) return null;
  return ibeta(df / (df + t * t), df / 2, 0.5);
}
/** t quantile with two-sided tail mass alpha (0.05 gives the 97.5% point), by bisection. */
function tCrit(df, alpha = 0.05) {
  let lo = 0;
  let hi = 1000;
  for (let i = 0; i < 200; i++) {
    const mid = (lo + hi) / 2;
    if (tTwoSided(mid, df) > alpha) lo = mid;
    else hi = mid;
  }
  return (lo + hi) / 2;
}
function erf(x) {
  const s = Math.sign(x);
  x = Math.abs(x);
  // erf(x) = P(1/2, x^2), regularized lower incomplete gamma, series / continued fraction.
  const a = 0.5;
  const z = x * x;
  let p;
  if (z < a + 1) {
    let sum = 1 / a;
    let term = sum;
    for (let n = 1; n < 500; n++) {
      term *= z / (a + n);
      sum += term;
      if (Math.abs(term) < Math.abs(sum) * 1e-15) break;
    }
    p = sum * Math.exp(-z + a * Math.log(z) - lgamma(a));
  } else {
    let b = z + 1 - a;
    let c = 1 / 1e-300;
    let d = 1 / b;
    let h = d;
    for (let i = 1; i < 500; i++) {
      const an = -i * (i - a);
      b += 2;
      d = an * d + b;
      if (Math.abs(d) < 1e-300) d = 1e-300;
      c = b + an / c;
      if (Math.abs(c) < 1e-300) c = 1e-300;
      d = 1 / d;
      const del = d * c;
      h *= del;
      if (Math.abs(del - 1) < 1e-15) break;
    }
    p = 1 - Math.exp(-z + a * Math.log(z) - lgamma(a)) * h;
  }
  return s * p;
}
const zTwoSided = (z) => (Number.isFinite(z) ? 1 - erf(Math.abs(z) / Math.SQRT2) : null);

const MIN_PER_GROUP = 8; // "too few participants" below this (section 2)
const mean = (xs) => (xs.length ? xs.reduce((a, b) => a + b, 0) / xs.length : null);
function variance(xs) {
  if (xs.length < 2) return null;
  const m = mean(xs);
  return xs.reduce((a, x) => a + (x - m) ** 2, 0) / (xs.length - 1);
}
function quantile(xs, q) {
  if (!xs.length) return null;
  const s = xs.slice().sort((a, b) => a - b);
  const pos = (s.length - 1) * q;
  const lo = Math.floor(pos);
  const hi = Math.ceil(pos);
  return s[lo] + (s[hi] - s[lo]) * (pos - lo);
}
function summary(raw) {
  const xs = raw.filter((x) => x !== null && x !== undefined && Number.isFinite(x));
  const n = xs.length;
  const m = mean(xs);
  const v = variance(xs);
  const sd = v === null ? null : Math.sqrt(v);
  let ci95 = null;
  if (n >= 2) {
    const h = tCrit(n - 1) * (sd / Math.sqrt(n));
    ci95 = [m - h, m + h];
  }
  return { n, mean: m, sd, ci95, median: quantile(xs, 0.5) };
}
/** sign: +1 when the hypothesis expects a positive value, -1 negative, 0 for no directional hypothesis (n_a). */
function reading(value, p, enough, sign) {
  if (!enough) return "too_few";
  if (sign === 0) return "n_a";
  if (value === null || p === null) return "n_a";
  if (p >= 0.05 || Math.sign(value) === 0) return "does_not_support";
  // "opposite": significant, but against the hypothesis (contract section 6).
  return Math.sign(value) === sign ? "supports" : "opposite";
}
/** Welch comparison, diff = checker - neutral. */
function comparison(checkerVals, neutralVals, sign) {
  const a = checkerVals.filter((x) => x !== null && x !== undefined);
  const b = neutralVals.filter((x) => x !== null && x !== undefined);
  const checker = summary(a);
  const neutral = summary(b);
  let diff = null;
  let diffCi95 = null;
  let welchT = null;
  let df = null;
  let p = null;
  let hedgesG = null;
  if (a.length && b.length) diff = checker.mean - neutral.mean;
  if (a.length >= 2 && b.length >= 2) {
    const va = variance(a);
    const vb = variance(b);
    const qa = va / a.length;
    const qb = vb / b.length;
    const se = Math.sqrt(qa + qb);
    if (se > 0) {
      welchT = diff / se;
      df = (qa + qb) ** 2 / (qa ** 2 / (a.length - 1) + qb ** 2 / (b.length - 1));
      p = tTwoSided(welchT, df);
      const h = tCrit(df) * se;
      diffCi95 = [diff - h, diff + h];
    }
    const sp = Math.sqrt(((a.length - 1) * va + (b.length - 1) * vb) / (a.length + b.length - 2));
    if (sp > 0) hedgesG = (diff / sp) * (1 - 3 / (4 * (a.length + b.length) - 9));
  }
  const enough = a.length >= MIN_PER_GROUP && b.length >= MIN_PER_GROUP;
  return { checker, neutral, diff, diffCi95, welchT, df, p, hedgesG, reading: reading(diff, p, enough, sign) };
}
function oneSample(raw, sign) {
  const s = summary(raw);
  let t = null;
  let df = null;
  let p = null;
  if (s.n >= 2 && s.sd > 0) {
    t = s.mean / (s.sd / Math.sqrt(s.n));
    df = s.n - 1;
    p = tTwoSided(t, df);
  }
  return { n: s.n, mean: s.mean, ci95: s.ci95, t, df, p, reading: reading(s.mean, p, s.n >= MIN_PER_GROUP, sign) };
}
/** Paired: pairs of [reasoned, denial]; meanDiff = mean(reasoned - denial), expected > 0 (H3). */
function paired(pairs) {
  const ok = pairs.filter(([r, d]) => r !== null && d !== null);
  const diffs = ok.map(([r, d]) => r - d);
  const s = summary(diffs);
  let t = null;
  let df = null;
  let p = null;
  if (s.n >= 2 && s.sd > 0) {
    t = s.mean / (s.sd / Math.sqrt(s.n));
    df = s.n - 1;
    p = tTwoSided(t, df);
  }
  return {
    n: s.n,
    reasoned: summary(ok.map(([r]) => r)),
    denial: summary(ok.map(([, d]) => d)),
    meanDiff: s.mean,
    ci95: s.ci95,
    t,
    df,
    p,
    reading: reading(s.mean, p, s.n >= MIN_PER_GROUP, +1),
  };
}
/** Two-proportion z with pooled variance; diff = p1 - p2 (checker - neutral). */
function twoProportion(x1, n1, x2, n2, sign, enough) {
  const p1 = n1 ? x1 / n1 : null;
  const p2 = n2 ? x2 / n2 : null;
  let diff = null;
  let z = null;
  let p = null;
  if (n1 && n2) {
    diff = p1 - p2;
    const pp = (x1 + x2) / (n1 + n2);
    const se = Math.sqrt(pp * (1 - pp) * (1 / n1 + 1 / n2));
    if (se > 0) {
      z = diff / se;
      p = zTwoSided(z);
    }
  }
  return { p1, p2, diff, z, p, reading: reading(diff, p, enough, sign) };
}
function ranks(xs) {
  const idx = xs.map((x, i) => [x, i]).sort((a, b) => a[0] - b[0]);
  const r = new Array(xs.length);
  for (let i = 0; i < idx.length;) {
    let j = i;
    while (j + 1 < idx.length && idx[j + 1][0] === idx[i][0]) j++;
    for (let k = i; k <= j; k++) r[idx[k][1]] = (i + j) / 2 + 1;
    i = j + 1;
  }
  return r;
}
function spearman(pairs) {
  const ok = pairs.filter(([x, y]) => x !== null && y !== null);
  if (ok.length < 3) return null;
  const rx = ranks(ok.map(([x]) => x));
  const ry = ranks(ok.map(([, y]) => y));
  const mx = mean(rx);
  const my = mean(ry);
  let sxy = 0;
  let sxx = 0;
  let syy = 0;
  for (let i = 0; i < ok.length; i++) {
    sxy += (rx[i] - mx) * (ry[i] - my);
    sxx += (rx[i] - mx) ** 2;
    syy += (ry[i] - my) ** 2;
  }
  return sxx > 0 && syy > 0 ? sxy / Math.sqrt(sxx * syy) : null;
}

// ---------- config (content/study-config.v1.4.json, as the server would read it) ----------
const GROUPS = ["checker", "neutral"];
const slides = [
  { id: "s1", n: 1, title: "Where we stand today", flagged: false },
  { id: "s2", n: 2, title: "The risk and what it would cost us", flagged: true },
  { id: "s3", n: 3, title: "How we will protect our 13 sites", flagged: true },
  { id: "s4", n: 4, title: "What the $2.4 million pays for", flagged: true },
  { id: "s5", n: 5, title: "The decision today", flagged: false },
];
const flaggedSlides = slides.filter((s) => s.flagged).map((s) => s.id);
const ratingPoints = [{ point: "T0", slideId: "opening", phase: "baseline", label: "Baseline" }];
for (const s of slides) {
  if (s.flagged) {
    ratingPoints.push({ point: `T${s.n}a`, slideId: s.id, phase: "afterSlide", label: `After slide ${s.n}` });
    ratingPoints.push({ point: `T${s.n}b`, slideId: s.id, phase: "afterAnswer", label: `After answer ${s.n}` });
  } else {
    ratingPoints.push({ point: `T${s.n}`, slideId: s.id, phase: "afterSlide", label: `After slide ${s.n}` });
  }
}
const POINTS = ratingPoints.map((r) => r.point);
const items = [
  { id: "trust", label: "Trust", min: 0, max: 10 },
  { id: "competence", label: "Competence", min: 1, max: 7 },
  { id: "benevolence", label: "Benevolence", min: 1, max: 7 },
  { id: "integrity", label: "Integrity", min: 1, max: 7 },
];
const actionsCfg = {
  checker: [
    { id: "nothing", label: "Say nothing" },
    { id: "explain", label: "Ask Daniel to explain" },
    { id: "challenge", label: "Challenge Daniel with the checker's finding" },
  ],
  neutral: [
    { id: "hasQuestion", label: "Yes" },
    { id: "noQuestion", label: "No" },
  ],
};
const opts = (prefix, pairs) => pairs.map(([id, label]) => ({ id: `${prefix}.${id}`, label }));
const background = [
  {
    id: "field",
    label: "What is your main field of expertise?",
    kind: "single",
    options: opts("field", [
      ["business", "Business and management"],
      ["finance", "Finance or accounting"],
      ["law", "Law"],
      ["it", "Information technology"],
      ["cyber", "Cybersecurity"],
      ["engineering", "Engineering"],
      ["health", "Health care or medicine"],
      ["education", "Education"],
      ["naturalSciences", "Natural sciences"],
      ["socialSciences", "Social sciences or humanities"],
      ["arts", "Arts, media or design"],
      ["government", "Government or public service"],
      ["studentNoField", "Student, no field yet"],
      ["other", "Other (please specify)"],
    ]),
  },
  { id: "country", label: "In which country do you live?", kind: "country" },
  {
    id: "age",
    label: "What is your age?",
    kind: "single",
    options: opts("age", [
      ["a18", "18 to 24"],
      ["a25", "25 to 34"],
      ["a35", "35 to 44"],
      ["a45", "45 to 54"],
      ["a55", "55 to 64"],
      ["a65", "65 or older"],
    ]),
  },
  {
    id: "jobLevel",
    label: "What is your current job level?",
    kind: "single",
    options: opts("jobLevel", [
      ["student", "Student"],
      ["employee", "Employee"],
      ["leadership", "Leadership"],
      ["cLevel", "C-Level"],
    ]),
  },
  {
    id: "education",
    label: "What is the highest level of education you have completed?",
    kind: "single",
    options: opts("education", [
      ["secondary", "Secondary school or less"],
      ["college", "College diploma or trade certificate"],
      ["bachelor", "Bachelor's degree"],
      ["master", "Master's degree"],
      ["doctorate", "Doctorate"],
      ["other", "Other"],
    ]),
  },
  {
    id: "cyberKnowledge",
    label: "How would you rate your own knowledge of cybersecurity?",
    kind: "scale",
    scale: { min: 1, max: 5, pointLabels: { 1: "None", 2: "Basic", 3: "Some", 4: "Good", 5: "Expert" } },
  },
  {
    id: "aiUse",
    label: "How often do you use AI tools such as chatbots or writing assistants?",
    kind: "single",
    options: opts("aiUse", [
      ["never", "Never"],
      ["lessMonthly", "Less than once a month"],
      ["monthly", "About once a month"],
      ["weekly", "About once a week"],
      ["severalWeekly", "Several times a week"],
      ["daily", "Every day"],
    ]),
  },
];
const budgetOptions = [
  { id: "approve", label: "Approve the request" },
  { id: "reject", label: "Reject the request" },
];
const PNTS = "preferNotToSay";

// ---------- emulated participants ----------
// Effects per group: drop at the flag, recovery after a reason, recovery after a denial (neutral: a third).
const EFFECT = {
  checker: { flag: -1.5, reason: 1.0, denial: 0.3 },
  neutral: { flag: -0.5, reason: 0.33, denial: 0.1 },
};
const PER_GROUP = 22;
const START = Date.parse("2026-10-25T08:00:00Z");

function makeBackground() {
  const bg = {};
  const w = {
    field: [8, 4, 1, 6, 3, 3, 2, 2, 1, 2, 1, 2, 2, 1],
    age: [3, 9, 7, 5, 3, 1],
    jobLevel: [3, 9, 6, 2],
    education: [1, 2, 8, 7, 2, 1],
    aiUse: [1, 2, 2, 5, 7, 6],
  };
  for (const item of background) {
    if (chance(0.05)) {
      bg[item.id] = PNTS;
      continue;
    }
    if (item.kind === "country") {
      bg.country = pick(
        ["CA", "US", "CH", "DE", "GB", "IN", "FR", "NL", "AU", "BR", "JP"],
        [12, 8, 7, 5, 4, 4, 1, 1, 1, 1, 1],
      );
    } else if (item.kind === "scale") {
      bg[item.id] = pick([1, 2, 3, 4, 5], [2, 6, 8, 5, 2]);
    } else {
      bg[item.id] = pick(
        item.options.map((o) => o.id),
        w[item.id],
      );
    }
  }
  return bg;
}

function emulate(group, denialSlide, k) {
  const e = EFFECT[group];
  const ratings = {};
  const actions = {};
  let latent = clamp(normal(6.3, 1.3), 1.5, 9.5);
  const meti = (t) => clamp(Math.round(1 + (t / 10) * 6 + normal(0, 0.6)), 1, 7);
  const record = (point) => {
    if (chance(0.02)) {
      ratings[point] = { trust: null, competence: null, benevolence: null, integrity: null, status: "timeout" };
      return;
    }
    const trust = clamp(Math.round(latent), 0, 10);
    ratings[point] = {
      trust,
      competence: meti(latent),
      benevolence: meti(latent),
      integrity: meti(latent),
      status: "answered",
    };
  };
  record("T0");
  latent = clamp(latent + normal(group === "checker" ? 0.15 : 0.05, 0.5), 0, 10);
  record("T1");
  for (const s of flaggedSlides) {
    const n = s.slice(1);
    latent = clamp(latent + e.flag + normal(0, 0.6), 0, 10);
    record(`T${n}a`);
    // action
    let actionId;
    let timedOut = false;
    if (group === "checker") {
      actionId = pick(["nothing", "explain", "challenge", null], [31, 40, 25, 4]);
    } else {
      actionId = pick(["hasQuestion", "noQuestion", null], [40, 56, 4]);
    }
    if (actionId === null) timedOut = true;
    const asked = actionId === "explain" || actionId === "challenge" || actionId === "hasQuestion";
    const answerType = s === denialSlide ? "denial" : "reason";
    actions[s] = { actionId, timedOut, questionBy: asked ? "member" : "chair", answerType };
    latent = clamp(latent + (answerType === "denial" ? e.denial : e.reason) + normal(0, 0.6), 0, 10);
    record(`T${n}b`);
  }
  latent = clamp(latent + normal(0, 0.4), 0, 10);
  record("T5");
  const last = ratings.T5.trust ?? ratings.T4b.trust ?? latent;
  let budget;
  if (chance(0.03)) budget = { optionId: null, confidence: null, status: "timeout" };
  else {
    const pApprove = 1 / (1 + Math.exp(-(last - 5) * 0.9));
    budget = {
      optionId: chance(pApprove) ? "approve" : "reject",
      confidence: clamp(Math.round(5 + Math.abs(last - 5) * 0.6 + normal(0, 1.5)), 0, 10),
      status: "answered",
    };
  }
  const flags = ["possible_duplicate"];
  const reloads = chance(0.1) ? 1 + Math.floor(rand() * 2) : 0;
  if (chance(0.08)) flags.push("tab_hidden_during_audio");
  if (reloads) flags.push("reloaded");
  const durationSec = Math.round(600 + rand() * 500);
  const finished = START + k * 9 * 60 * 1000 + Math.round(rand() * 300000);
  return {
    sessionId: hex(20),
    group,
    denialSlide,
    status: "Finished",
    finishedAt: new Date(finished).toISOString(),
    durationSec,
    ratings,
    actions,
    budget,
    background: makeBackground(),
    flags,
    _reloads: reloads,
    _audio: "correct",
  };
}

const participants = [];
let k = 0;
for (const group of GROUPS) {
  const denials = shuffle(Array.from({ length: PER_GROUP }, (_, i) => flaggedSlides[i % 3]));
  for (let i = 0; i < PER_GROUP; i++) participants.push(emulate(group, denials[i], k++));
}
participants.sort((a, b) => a.finishedAt.localeCompare(b.finishedAt));

// Sessions that are not analysed: finished but excluded, and sessions that did not finish.
const excludedFinished = [
  { group: "checker", flags: ["possible_duplicate", "honeypot"], audio: "correct" },
  { group: "neutral", flags: ["possible_duplicate", "audio_check_wrong"], audio: "wrong" },
  { group: "checker", flags: ["forced_allocation"], audio: "correct" },
  { group: "neutral", flags: ["forced_allocation"], audio: "correct" },
];
const unfinished = [
  { status: "opened", group: null, consented: false },
  { status: "opened", group: null, consented: false },
  { status: "blocked", group: null, consented: false },
  { status: "started", group: "checker", consented: true },
  { status: "ongoing", group: "neutral", consented: true },
  { status: "abandoned", group: "checker", consented: true },
  { status: "abandoned", group: "neutral", consented: true },
];

// ---------- per-participant deltas (section 2) ----------
const trustAt = (p, point) => p.ratings[point]?.trust ?? null;
const delta = (p, from, to) => {
  const a = trustAt(p, from);
  const b = trustAt(p, to);
  return a === null || b === null ? null : b - a;
};
// DECISION: "previous rating" is the rating point immediately before T{n}a in plan order (T1 for s2, T2b for s3,
// T3b for s4). If its trust is missing, there is no delta for that slide (section 6: both points must be present).
const prevPoint = (slideId) => POINTS[POINTS.indexOf(`T${slideId.slice(1)}a`) - 1];
const pt = (slideId, phase) => `T${slideId.slice(1)}${phase}`;
const meanOrNull = (xs) => mean(xs.filter((x) => x !== null));

for (const p of participants) {
  const h1 = flaggedSlides.map((s) => delta(p, prevPoint(s), pt(s, "a")));
  const reasoned = flaggedSlides.filter((s) => p.actions[s].answerType === "reason");
  const h2 = reasoned.map((s) => delta(p, pt(s, "a"), pt(s, "b")));
  const net = reasoned.map((s) => delta(p, prevPoint(s), pt(s, "b")));
  p._h1PerSlide = Object.fromEntries(flaggedSlides.map((s, i) => [s, h1[i]]));
  p._h2PerSlide = Object.fromEntries(reasoned.map((s, i) => [s, h2[i]]));
  p._h1 = meanOrNull(h1);
  p._h2 = meanOrNull(h2);
  p._net = meanOrNull(net);
  p._denial = delta(p, pt(p.denialSlide, "a"), pt(p.denialSlide, "b"));
  const asked = flaggedSlides.filter((s) => {
    const a = p.actions[s].actionId;
    return a === "explain" || a === "challenge" || a === "hasQuestion";
  }).length;
  p._share = asked / flaggedSlides.length;
  p._asked = asked;
}
const byGroup = (g) => participants.filter((p) => p.group === g);
const C = byGroup("checker");
const N = byGroup("neutral");
const col = (ps, f) => ps.map(f);

// ---------- hypotheses ----------
/** Per-slide breakdowns carry no reading (contract section 6). */
const noReading = (obj) => Object.fromEntries(Object.entries(obj).map(([k, v]) => [k, { ...v, reading: "n_a" }]));
const LAST = POINTS[POINTS.length - 1];
const reasonedOf = (p) => flaggedSlides.filter((s) => p.actions[s].answerType === "reason");
const levelAfterReason = (p) => meanOrNull(reasonedOf(p).map((s) => trustAt(p, pt(s, "b"))));
const aboveBaseline = (p) => meanOrNull(reasonedOf(p).map((s) => delta(p, "T0", pt(s, "b"))));
const lastTrust = (p) => {
  for (let i = POINTS.length - 1; i >= 0; i--) if (trustAt(p, POINTS[i]) !== null) return trustAt(p, POINTS[i]);
  return null;
};
// DECISION: terciles of the last trust rating over all analysed participants with one: cut points at the 1/3 and
// 2/3 quantiles; tercile 1 is <= the first cut, 3 is > the second; from/to are the lowest and highest rating in it.
function lastTrustTerciles(ps) {
  const withT = ps.map((p) => ({ p, t: lastTrust(p) })).filter((r) => r.t !== null);
  const c1 = quantile(
    withT.map((r) => r.t),
    1 / 3,
  );
  const c2 = quantile(
    withT.map((r) => r.t),
    2 / 3,
  );
  return [1, 2, 3]
    .map((tercile) => {
      const rows = withT.filter((r) => (tercile === 1 ? r.t <= c1 : tercile === 2 ? r.t > c1 && r.t <= c2 : r.t > c2));
      const ts = rows.map((r) => r.t);
      return {
        tercile,
        from: ts.length ? Math.min(...ts) : null,
        to: ts.length ? Math.max(...ts) : null,
        approve: rows.filter((r) => r.p.budget?.optionId === "approve").length,
        reject: rows.filter((r) => r.p.budget?.optionId === "reject").length,
        missing: rows.filter((r) => !r.p.budget || r.p.budget.optionId === null).length,
      };
    })
    .filter((t) => t.from !== null);
}
const perSlide = (f, sign) =>
  Object.fromEntries(
    flaggedSlides.map((s) => [
      s,
      comparison(
        col(C, (p) => f(p, s)),
        col(N, (p) => f(p, s)),
        sign,
      ),
    ]),
  );

const h3Pairs = (ps) => ps.map((p) => [p._h2, p._denial]);
const share = (ps) => ({
  slides: ps.length * flaggedSlides.length,
  withQuestion: ps.reduce((a, p) => a + p._asked, 0),
  share: ps.length ? ps.reduce((a, p) => a + p._asked, 0) / (ps.length * flaggedSlides.length) : null,
  perParticipant: summary(col(ps, (p) => p._share)),
});
const enoughGroups = C.length >= MIN_PER_GROUP && N.length >= MIN_PER_GROUP;
const sC = share(C);
const sN = share(N);

const budgetOf = (ps) => ({
  approve: ps.filter((p) => p.budget?.optionId === "approve").length,
  reject: ps.filter((p) => p.budget?.optionId === "reject").length,
  missing: ps.filter((p) => !p.budget || p.budget.optionId === null).length,
  confidence: summary(col(ps, (p) => p.budget?.confidence ?? null)),
});
const bC = budgetOf(C);
const bN = budgetOf(N);

const aiOrder = background.find((b) => b.id === "aiUse").options.map((o) => o.id);
const num = (v) => (typeof v === "number" ? v : null);
const aiRank = (v) => (typeof v === "string" && aiOrder.includes(v) ? aiOrder.indexOf(v) + 1 : null);
const sumBy = (ps, f) => Object.fromEntries(GROUPS.map((g) => [g, summary(ps.filter((p) => p.group === g).map(f))]));

const hypotheses = {
  // Validity checks: no directional hypothesis, so their reading is "n_a" (or "too_few").
  baseline: comparison(
    col(C, (p) => trustAt(p, "T0")),
    col(N, (p) => trustAt(p, "T0")),
    0,
  ),
  slide1: comparison(
    col(C, (p) => delta(p, "T0", "T1")),
    col(N, (p) => delta(p, "T0", "T1")),
    0,
  ),
  // Research question: last rating minus T0, no direction.
  overall: comparison(
    col(C, (p) => delta(p, "T0", LAST)),
    col(N, (p) => delta(p, "T0", LAST)),
    0,
  ),
  h1: {
    pooled: comparison(
      col(C, (p) => p._h1),
      col(N, (p) => p._h1),
      -1,
    ),
    checkerBelowZero: oneSample(
      col(C, (p) => p._h1),
      -1,
    ),
    perSlide: noReading(perSlide((p, s) => p._h1PerSlide[s], -1)),
    levelAtFlag: noReading(perSlide((p, s) => trustAt(p, pt(s, "a")), -1)),
  },
  h2: {
    pooled: comparison(
      col(C, (p) => p._h2),
      col(N, (p) => p._h2),
      +1,
    ),
    checkerAboveZero: oneSample(
      col(C, (p) => p._h2),
      +1,
    ),
    // H2(b): level after the reasoned answer, mean over the participant's reasoned slides; carries the H2 reading.
    level: comparison(col(C, levelAfterReason), col(N, levelAfterReason), +1),
    // Per slide: only participants for whom that slide was a reasoned one.
    perSlide: noReading(perSlide((p, s) => p._h2PerSlide[s] ?? null, +1)),
    levelPerSlide: noReading(
      perSlide((p, s) => (p.actions[s].answerType === "reason" ? trustAt(p, pt(s, "b")) : null), +1),
    ),
    // Exploratory: above the first rating, within each group.
    aboveBaseline: { checker: oneSample(col(C, aboveBaseline), +1), neutral: oneSample(col(N, aboveBaseline), +1) },
    // Secondary: T{n}b minus the rating before the flag, between groups, no direction.
    exploratoryNet: comparison(
      col(C, (p) => p._net),
      col(N, (p) => p._net),
      0,
    ),
  },
  h3: {
    perGroup: { checker: paired(h3Pairs(C)), neutral: paired(h3Pairs(N)) },
    both: paired(h3Pairs(participants)),
  },
  h4a: {
    checker: sC,
    neutral: sN,
    // Slides are not independent: numbers only, reading "n_a"; participantsTest carries the H4a reading.
    slidesTest: {
      ...twoProportion(sC.withQuestion, sC.slides, sN.withQuestion, sN.slides, +1, enoughGroups),
      reading: "n_a",
    },
    participantsTest: comparison(
      col(C, (p) => p._share),
      col(N, (p) => p._share),
      +1,
    ),
  },
  // H4b amplification: after a reason checker - neutral > 0 (also H2's between-group change, H4b1).
  h4b: {
    afterReason: comparison(
      col(C, (p) => p._h2),
      col(N, (p) => p._h2),
      +1,
    ),
    // As the Go endpoint: after a denial the checker group is expected to fall further, checker - neutral < 0.
    afterDenial: comparison(
      col(C, (p) => p._denial),
      col(N, (p) => p._denial),
      -1,
    ),
  },
  budget: {
    byGroup: { checker: bC, neutral: bN },
    byLastTrustTercile: lastTrustTerciles(participants),
    // No directional hypothesis for the vote in the Abstract: reading is "n_a" (or "too_few").
    approveTest: twoProportion(
      bC.approve,
      bC.approve + bC.reject,
      bN.approve,
      bN.approve + bN.reject,
      0,
      bC.approve + bC.reject >= MIN_PER_GROUP && bN.approve + bN.reject >= MIN_PER_GROUP,
    ),
  },
  moderators: {
    cyberKnowledge: [1, 2, 3, 4, 5].map((level) => {
      const ps = participants.filter((p) => p.background?.cyberKnowledge === level);
      return { level, h1: sumBy(ps, (p) => p._h1), h2: sumBy(ps, (p) => p._h2) };
    }),
    aiUse: aiOrder.map((optionId) => {
      const ps = participants.filter((p) => p.background?.aiUse === optionId);
      return { optionId, h1: sumBy(ps, (p) => p._h1), h2: sumBy(ps, (p) => p._h2) };
    }),
    spearman: {
      cyberKnowledgeH1: Object.fromEntries(
        GROUPS.map((g) => [g, spearman(byGroup(g).map((p) => [num(p.background?.cyberKnowledge), p._h1]))]),
      ),
      cyberKnowledgeH2: Object.fromEntries(
        GROUPS.map((g) => [g, spearman(byGroup(g).map((p) => [num(p.background?.cyberKnowledge), p._h2]))]),
      ),
      aiUseH1: Object.fromEntries(
        GROUPS.map((g) => [g, spearman(byGroup(g).map((p) => [aiRank(p.background?.aiUse), p._h1]))]),
      ),
      aiUseH2: Object.fromEntries(
        GROUPS.map((g) => [g, spearman(byGroup(g).map((p) => [aiRank(p.background?.aiUse), p._h2]))]),
      ),
    },
  },
};

// ---------- descriptive blocks ----------
const ratingStats = {};
for (const item of items) {
  ratingStats[item.id] = {};
  for (const point of POINTS) {
    ratingStats[item.id][point] = Object.fromEntries(
      GROUPS.map((g) => [g, summary(byGroup(g).map((p) => p.ratings[point]?.[item.id] ?? null))]),
    );
  }
}

const actions = {};
for (const s of flaggedSlides) {
  actions[s] = {};
  for (const g of GROUPS) {
    const counts = Object.fromEntries([...actionsCfg[g].map((a) => [a.id, 0]), ["timeout", 0]]);
    for (const p of byGroup(g)) {
      const a = p.actions[s];
      counts[a.timedOut || a.actionId === null ? "timeout" : a.actionId]++;
    }
    actions[s][g] = counts;
  }
}

// As the Go endpoint: `background` holds every item. Entries follow the config's option order (for the country:
// the codes by n descending), then "preferNotToSay" and "missing", zero counts included.
const countryCodes = [
  ...new Set(participants.map((p) => p.background?.country).filter((c) => typeof c === "string" && c !== PNTS)),
];
const countryN = (c) => participants.filter((p) => p.background?.country === c).length;
countryCodes.sort((a, b) => countryN(b) - countryN(a) || a.localeCompare(b));
const bgCounts = {};
for (const item of background) {
  const entries =
    item.kind === "country"
      ? countryCodes.map((c) => ({ id: c, label: c, value: c }))
      : item.kind === "scale"
        ? Array.from({ length: item.scale.max - item.scale.min + 1 }, (_, i) => {
            const v = item.scale.min + i;
            return { id: String(v), label: item.scale.pointLabels?.[v] ?? String(v), value: v };
          })
        : item.options.map((o) => ({ id: o.id, label: o.label, value: o.id }));
  entries.push({ id: PNTS, label: "Prefer not to say", value: PNTS });
  entries.push({ id: "missing", label: "Missing", value: undefined });
  bgCounts[item.id] = entries.map((e) => {
    const match = (p) => {
      const v = p.background ? p.background[item.id] : undefined;
      return e.id === "missing" ? v === undefined || v === null : v === e.value;
    };
    const byG = Object.fromEntries(GROUPS.map((g) => [g, byGroup(g).filter(match).length]));
    return { id: e.id, label: e.label, n: byG.checker + byG.neutral, byGroup: byG };
  });
}

// As the Go endpoint: `countries` holds real ISO codes only; "Prefer not to say" and "missing" are in
// background.country.
const countryMap = new Map();
for (const p of participants) {
  const code = p.background?.country;
  if (typeof code !== "string" || code === PNTS) continue;
  if (!countryMap.has(code)) countryMap.set(code, { code, n: 0, byGroup: { checker: 0, neutral: 0 } });
  const c = countryMap.get(code);
  c.n++;
  c.byGroup[p.group]++;
}
const countries = [...countryMap.values()].sort((a, b) => b.n - a.n || a.code.localeCompare(b.code));

// As the Go endpoint: quality covers the sessions left after the forced and include rules but before the flag
// toggles, so the flags table shows what the toggles would remove. The flag-excluded sessions here are stubs
// without ratings or reloads, so they add nothing to rating timeouts and reloads.
const qualitySessions = [
  ...participants.map((p) => ({ flags: p.flags, audio: p._audio })),
  ...excludedFinished.filter((s) => !s.flags.includes("forced_allocation")),
];
const qFlags = {};
for (const s of qualitySessions) for (const f of s.flags) qFlags[f] = (qFlags[f] ?? 0) + 1;
const quality = {
  flags: Object.fromEntries(Object.entries(qFlags).sort((a, b) => b[1] - a[1] || a[0].localeCompare(b[0]))),
  audioCheck: {
    correct: qualitySessions.filter((s) => s.audio === "correct").length,
    wrong: qualitySessions.filter((s) => s.audio === "wrong").length,
    missing: qualitySessions.filter((s) => !s.audio).length,
  },
  ratingTimeouts: Object.fromEntries(
    POINTS.map((pt2) => [pt2, participants.filter((p) => p.ratings[pt2]?.status === "timeout").length]),
  ),
  reloads: {
    sessions: participants.filter((p) => p._reloads > 0).length,
    total: participants.reduce((a, p) => a + p._reloads, 0),
  },
};

// ---------- counts ----------
const durations = participants.map((p) => p.durationSec).filter((d) => d !== null);
const lo = Math.floor(Math.min(...durations) / 60) * 60;
const hi = Math.ceil((Math.max(...durations) + 1) / 60) * 60;
const histogram = [];
for (let from = lo; from < hi; from += 60) {
  histogram.push({ fromSec: from, toSec: from + 60, n: durations.filter((d) => d >= from && d < from + 60).length });
}
const finishedAll = participants.length + excludedFinished.length;
const st = (s) => unfinished.filter((u) => u.status === s).length;
const groupCounts = (g) => ({
  consented:
    byGroup(g).length +
    excludedFinished.filter((s) => s.group === g).length +
    unfinished.filter((u) => u.group === g && u.consented).length,
  finished: byGroup(g).length + excludedFinished.filter((s) => s.group === g).length,
  analysed: byGroup(g).length,
});
const counts = {
  byStatus: {
    opened: st("opened"),
    started: st("started"),
    ongoing: st("ongoing"),
    finished: finishedAll,
    abandoned: st("abandoned"),
    blocked: st("blocked"),
  },
  total: finishedAll + unfinished.length,
  consented: finishedAll + unfinished.filter((u) => u.consented).length,
  analysed: participants.length,
  excluded: {
    forced: excludedFinished.filter((s) => s.flags.includes("forced_allocation")).length,
    byFlag: Object.fromEntries(
      ["honeypot", "audio_check_wrong", "many_invalid_events"].map((f) => [
        f,
        excludedFinished.filter((s) => !s.flags.includes("forced_allocation") && s.flags.includes(f)).length,
      ]),
    ),
    // Consented sessions left out by include=finished.
    notIncluded: unfinished.filter((u) => u.consented).length,
  },
  byGroup: { checker: groupCounts("checker"), neutral: groupCounts("neutral") },
  durationsSec: {
    n: durations.length,
    median: quantile(durations, 0.5),
    p25: quantile(durations, 0.25),
    p75: quantile(durations, 0.75),
    histogram,
  },
};

// ---------- document ----------
const doc = {
  run: {
    id: "run-dashboard-pilot",
    code: "DPILOT",
    name: "Dashboard pilot (emulated)",
    source: "Emulated",
    configVersion: "1.4.0",
    status: "open",
    createdAt: "2026-10-25T07:30:00Z",
    testOverrides: false,
  },
  generatedAt: "2026-10-25T12:00:00Z",
  filters: {
    include: "finished",
    exclude: ["forced_allocation", "honeypot", "audio_check_wrong", "many_invalid_events"],
    group: "both",
  },
  config: {
    groups: GROUPS,
    slides,
    flaggedSlides,
    ratingPoints,
    items,
    actions: actionsCfg,
    background,
    budgetOptions,
  },
  counts,
  participants: participants.map((p) => {
    const out = {};
    for (const [key, v] of Object.entries(p)) if (!key.startsWith("_")) out[key] = v;
    return out;
  }),
  ratingStats,
  hypotheses,
  actions,
  background: bgCounts,
  countries,
  quality,
};

const round = (v) => {
  if (typeof v === "number") {
    if (!Number.isFinite(v)) return null;
    if (Number.isInteger(v)) return v;
    // Four decimals, but small values (p values) keep four significant digits so "< .001" stays checkable.
    return Math.abs(v) < 1e-3 ? Number(v.toPrecision(4)) : Math.round(v * 1e4) / 1e4;
  }
  if (Array.isArray(v)) return v.map(round);
  if (v && typeof v === "object") return Object.fromEntries(Object.entries(v).map(([k2, x]) => [k2, round(x)]));
  return v === undefined ? null : v;
};

mkdirSync(dirname(OUT), { recursive: true });
writeFileSync(OUT, JSON.stringify(round(doc), null, 2) + "\n");
const h = hypotheses;
const f = (x) => (x === null ? "null" : x.toFixed(3));
console.log(
  `wrote ${OUT}: ${participants.length} analysed (${C.length} checker, ${N.length} neutral); ` +
    `H1 checker ${f(h.h1.pooled.checker.mean)} p=${f(h.h1.pooled.p)} ${h.h1.pooled.reading}; ` +
    `H2 ${f(h.h2.pooled.checker.mean)} p=${f(h.h2.pooled.p)} ${h.h2.pooled.reading}; ` +
    `H3 checker p=${f(h.h3.perGroup.checker.p)} ${h.h3.perGroup.checker.reading}; ` +
    `H4a p=${f(h.h4a.participantsTest.p)} ${h.h4a.participantsTest.reading}; ` +
    `H2b level p=${f(h.h2.level.p)} ${h.h2.level.reading}`,
);
