import { describe, expect, it } from "vitest";
import { fmtCi, fmtDuration, fmtInt, fmtNum, fmtP, fmtPct, fmtSigned } from "./format";

describe("number formatting", () => {
  it("keeps at most two decimals and prints a dash for missing values", () => {
    expect(fmtNum(1.23456)).toBe("1.23");
    expect(fmtNum(6.3)).toBe("6.3");
    expect(fmtNum(-0.004)).toBe("0");
    expect(fmtNum(-1.555)).toBe("-1.56");
    expect(fmtNum(2)).toBe("2");
    expect(fmtNum(null)).toBe("-");
    expect(fmtNum(Number.NaN)).toBe("-");
    expect(fmtInt(1284)).toBe("1,284");
  });

  it("prints p without the leading zero and < .001 below that", () => {
    expect(fmtP(0.0000257)).toBe("< .001");
    expect(fmtP(0.000999)).toBe("< .001");
    expect(fmtP(0.001)).toBe(".001");
    expect(fmtP(0.0371)).toBe(".037");
    expect(fmtP(0.5)).toBe(".500");
    expect(fmtP(1)).toBe("1");
    expect(fmtP(null)).toBe("-");
  });

  it("prints a CI as [lo, hi]", () => {
    expect(fmtCi([-2.2024, -1.2976])).toBe("[-2.2, -1.3]");
    expect(fmtCi([0.123, 1.456])).toBe("[0.12, 1.46]");
    expect(fmtCi(null)).toBe("-");
  });

  it("formats shares, durations and signed deltas", () => {
    expect(fmtPct(0.6136)).toBe("61%");
    expect(fmtPct(null)).toBe("-");
    expect(fmtDuration(824.5)).toBe("13:45");
    expect(fmtDuration(3725)).toBe("1:02:05");
    expect(fmtSigned(1.09)).toBe("+1.09");
    expect(fmtSigned(-1.55)).toBe("-1.55");
    expect(fmtSigned(0)).toBe("0");
  });
});
