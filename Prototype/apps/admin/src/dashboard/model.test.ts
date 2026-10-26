import { describe, expect, it } from "vitest";
import type { RunAnalysis } from "../api/types";
import { analysisQuery } from "../api/client";
import { canDelete } from "../participants/ParticipantDrawer";
import { runFromHash } from "../runs/RunsTab";
import fixtureJson from "./fixtures/analysis.sample.json";
import {
  DEFAULT_FILTERS,
  actionKeys,
  backgroundBars,
  countryName,
  countryNonAnswers,
  countrySlices,
  denialPoint,
  filterParticipants,
  groupSeries,
  lastTrustTerciles,
  ordinalColor,
  ratingXs,
  readingChip,
  toggleExclude,
  trajectories,
} from "./model";

const a = fixtureJson as unknown as RunAnalysis;

describe("reading chips", () => {
  it("maps each reading to its tone", () => {
    expect(readingChip("supports", "H1")).toEqual({ label: "supports H1 so far", tone: "success" });
    expect(readingChip("does_not_support", "H2")).toEqual({ label: "does not support H2", tone: "neutral" });
    expect(readingChip("opposite", "H2")).toEqual({ label: "opposite", tone: "danger" });
    expect(readingChip("too_few")).toEqual({ label: "too few participants", tone: "info" });
    expect(readingChip("n_a").tone).toBe("muted");
    expect(readingChip(undefined).tone).toBe("muted");
  });
});

describe("filters and query", () => {
  it("defaults to the plan's exclusions", () => {
    expect(DEFAULT_FILTERS).toEqual({
      include: "finished",
      exclude: ["honeypot", "audio_check_wrong", "many_invalid_events"],
      group: "both",
    });
  });

  it("toggles a flag and keeps a stable order", () => {
    const f = toggleExclude(toggleExclude(DEFAULT_FILTERS, "reloaded", true), "honeypot", false);
    expect(f.exclude).toEqual(["audio_check_wrong", "many_invalid_events", "reloaded"]);
  });

  it("builds the analysis query; an empty exclude list is sent, forced_allocation never", () => {
    expect(analysisQuery({})).toBe("");
    expect(analysisQuery({ include: "consented", exclude: ["honeypot", "forced_allocation"], group: "checker" })).toBe(
      "?include=consented&exclude=honeypot&group=checker",
    );
    expect(analysisQuery({ exclude: [] })).toBe("?exclude=");
  });

  it("reads the run id from the hash", () => {
    expect(runFromHash("#run=abc123")).toBe("abc123");
    expect(runFromHash("#run=a%20b")).toBe("a b");
    expect(runFromHash("#preview=fixture")).toBeNull();
    expect(runFromHash("")).toBeNull();
  });

  it("allows a deletion only with the exact id and a reason of 1 to 500 characters", () => {
    expect(canDelete("abc", "abc", "asked by e-mail")).toBe(true);
    expect(canDelete("abc", "ABC", "asked")).toBe(false);
    expect(canDelete("abc", "abc ", "asked")).toBe(false);
    expect(canDelete("abc", "abc", "   ")).toBe(false);
    expect(canDelete("abc", "abc", "x".repeat(500))).toBe(true);
    expect(canDelete("abc", "abc", "x".repeat(501))).toBe(false);
  });
});

describe("chart data from the fixture", () => {
  it("puts the rating points in plan order with slide and phase under each tick, flags marked", () => {
    const xs = ratingXs(a);
    expect(xs.map((x) => x.id)).toEqual(["T0", "T1", "T2a", "T2b", "T3a", "T3b", "T4a", "T4b", "T5"]);
    expect(xs[0].sub).toBe("baseline");
    expect(xs[2]).toMatchObject({ sub: "slide 2", marked: true });
    expect(xs[3]).toMatchObject({ sub: "answer 2", marked: false });
    expect(xs[1].marked).toBe(false);
  });

  it("gives one series per group with the group's colour and the stats' CI", () => {
    const s = groupSeries(a, "trust");
    expect(s.map((x) => x.id)).toEqual(["checker", "neutral"]);
    expect(s[0].color).toBe("var(--viz-checker)");
    const t0 = a.ratingStats.trust.T0.checker;
    expect(s[0].points.T0).toEqual({ y: t0.mean, lo: t0.ci95?.[0], hi: t0.ci95?.[1], n: t0.n });
  });

  it("marks each participant's denial point hollow", () => {
    expect(denialPoint(a, "s3")).toBe("T3b");
    const t = trajectories(a, "trust");
    expect(t).toHaveLength(a.participants.length);
    const p = a.participants[0];
    expect(t[0].hollow).toEqual([denialPoint(a, p.denialSlide)]);
    expect(trajectories(a, "trust", "neutral").length).toBe(a.participants.filter((x) => x.group === "neutral").length);
  });

  it("keeps background bars in config order with prefer-not-to-say grey and missing only when present", () => {
    const bars = backgroundBars(a, "cyberKnowledge");
    expect(bars.map((b) => b.id)).toEqual(["1", "2", "3", "4", "5", "preferNotToSay"]);
    expect(bars[0].color).toBe("var(--viz-ord-1)");
    expect(bars[4].color).toBe("var(--viz-ord-6)");
    expect(bars[5].color).toBe("var(--viz-other)");
    expect(bars[1].extra).toEqual({ checker: 6, neutral: 8 });
    expect(backgroundBars(a, "field").every((b) => b.color === undefined || b.color === "var(--viz-other)")).toBe(true);
  });

  it("spreads ordinal steps over the six-step ramp", () => {
    expect([0, 1, 2, 3, 4, 5].map((i) => ordinalColor(i, 6))).toEqual(
      [1, 2, 3, 4, 5, 6].map((s) => `var(--viz-ord-${s})`),
    );
    expect(ordinalColor(0, 1)).toBe("var(--viz-ord-6)");
  });

  it("names countries and the non-answers", () => {
    expect(countryName("CA")).toBe("Canada");
    expect(countryName("preferNotToSay")).toBe("Prefer not to say");
    expect(countryName("zz")).toBe("ZZ");
  });

  it("adds a timeout segment to the action keys", () => {
    expect(actionKeys(a, "checker").map((k) => k.key)).toEqual(["nothing", "explain", "challenge", "timeout"]);
    expect(actionKeys(a, "neutral").map((k) => k.key)).toEqual(["hasQuestion", "noQuestion", "timeout"]);
  });

  it("reads the terciles from the server and falls back to the participant list", () => {
    const server = a.hypotheses.budget.byLastTrustTercile!;
    expect(lastTrustTerciles(a)).toBe(server);
    const noField = {
      ...a,
      hypotheses: { ...a.hypotheses, budget: { ...a.hypotheses.budget, byLastTrustTercile: undefined } },
    };
    const local = lastTrustTerciles(noField);
    expect(local).toEqual(server);
    const withTrust = a.participants.filter((p) => p.ratings.T5?.trust != null || p.ratings.T4b?.trust != null);
    expect(local.reduce((n, t) => n + t.approve + t.reject + t.missing, 0)).toBeGreaterThanOrEqual(withTrust.length);
  });

  it("keeps the donut to real ISO codes and takes the non-answers from background.country", () => {
    expect(countrySlices(a).every((s) => /^[A-Z]{2}$/.test(s.id))).toBe(true);
    const non = countryNonAnswers(a);
    const total = countrySlices(a).reduce((n, s) => n + s.value, 0) + non.preferNotToSay + non.missing;
    expect(total).toBe(a.participants.length);
  });

  it("filters the participant list by group in the preview, statistics untouched", () => {
    const f = filterParticipants(a, "checker");
    expect(f.participants.every((p) => p.group === "checker")).toBe(true);
    expect(f.ratingStats).toBe(a.ratingStats);
  });
});
