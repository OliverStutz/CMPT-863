import { describe, expect, it } from "vitest";
import { elapsedMs, elapsedText, formatElapsed, isTicking } from "./elapsed";

const consentAt = "2026-10-23T10:00:00.000Z";
const t = (iso: string) => Date.parse(iso);

describe("elapsed timer", () => {
  it("shows a dash before consent", () => {
    const row = { consentAt: null, finishedAt: null, status: "Opened" as const };
    expect(elapsedMs(row, t("2026-10-23T10:05:00Z"))).toBeNull();
    expect(elapsedText(row, t("2026-10-23T10:05:00Z"))).toBe("-");
    expect(isTicking(row)).toBe(false);
  });

  it("ticks with the clock while the session runs", () => {
    const row = { consentAt, finishedAt: null, status: "Ongoing" as const };
    expect(isTicking(row)).toBe(true);
    expect(elapsedText(row, t("2026-10-23T10:00:05Z"))).toBe("0:05");
    expect(elapsedText(row, t("2026-10-23T10:00:06Z"))).toBe("0:06");
    expect(elapsedText(row, t("2026-10-23T10:12:34Z"))).toBe("12:34");
    expect(elapsedText(row, t("2026-10-23T11:02:03Z"))).toBe("1:02:03");
  });

  it("freezes at finishedAt when Finished", () => {
    const row = { consentAt, finishedAt: "2026-10-23T10:20:00Z", status: "Finished" as const };
    expect(isTicking(row)).toBe(false);
    expect(elapsedText(row, t("2026-10-23T10:30:00Z"))).toBe("20:00");
    expect(elapsedText(row, t("2026-10-24T09:00:00Z"))).toBe("20:00");
  });

  it("freezes at lastEventAt for a Finished row without finishedAt", () => {
    const row = { consentAt, finishedAt: null, lastEventAt: "2026-10-23T10:01:00Z", status: "Finished" as const };
    expect(isTicking(row)).toBe(false);
    expect(elapsedText(row, t("2026-10-23T12:00:00Z"))).toBe("1:00");
  });

  it("keeps ticking for Abandoned, as (finishedAt or now) - consentAt", () => {
    const row = { consentAt, finishedAt: null, status: "Abandoned" as const };
    expect(isTicking(row)).toBe(true);
    expect(elapsedText(row, t("2026-10-23T10:45:00Z"))).toBe("45:00");
  });

  it("never goes negative when the browser clock lags the server", () => {
    const row = { consentAt, finishedAt: null, status: "Started" as const };
    expect(elapsedMs(row, t("2026-10-23T09:59:58Z"))).toBe(0);
  });

  it("formats null as a dash", () => {
    expect(formatElapsed(null)).toBe("-");
    expect(formatElapsed(0)).toBe("0:00");
  });
});
