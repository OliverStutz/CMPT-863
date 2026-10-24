import { describe, expect, it } from "vitest";
import { flagChip, statusChip } from "./status";
import { SESSION_STATUSES } from "../api/types";

describe("status chip mapping", () => {
  it("maps every session status to its tone", () => {
    expect(statusChip("Opened")).toEqual({ label: "Opened", tone: "neutral" });
    expect(statusChip("Started")).toEqual({ label: "Started", tone: "info" });
    expect(statusChip("Ongoing")).toEqual({ label: "Ongoing", tone: "active" });
    expect(statusChip("Finished")).toEqual({ label: "Finished", tone: "success" });
    expect(statusChip("Abandoned")).toEqual({ label: "Abandoned", tone: "warning" });
    expect(statusChip("Blocked")).toEqual({ label: "Blocked", tone: "danger" });
  });

  it("gives every status a distinct tone", () => {
    const tones = SESSION_STATUSES.map((s) => statusChip(s).tone);
    expect(new Set(tones).size).toBe(SESSION_STATUSES.length);
  });

  it("falls back to neutral for an unknown status", () => {
    expect(statusChip("Weird")).toEqual({ label: "Weird", tone: "neutral" });
    expect(statusChip("")).toEqual({ label: "-", tone: "neutral" });
  });

  it("maps known flags and keeps unknown flags by name", () => {
    expect(flagChip("possible_duplicate").tone).toBe("danger");
    expect(flagChip("reloaded").tone).toBe("info");
    expect(flagChip("audio_check_wrong").tone).toBe("warning");
    expect(flagChip("something_new")).toEqual({ label: "something_new", tone: "warning" });
  });
});
