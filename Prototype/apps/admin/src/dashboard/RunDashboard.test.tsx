import { describe, expect, it, vi } from "vitest";
import { renderToString } from "react-dom/server";
import type { RunAnalysis } from "../api/types";
import fixtureJson from "./fixtures/analysis.sample.json";

const fixture = fixtureJson as unknown as RunAnalysis;

vi.mock("../api/client", () => ({
  getRunAnalysis: vi.fn(async () => fixture),
}));

const { RunDashboard, SECTION_TITLES } = await import("./RunDashboard");
const { getRunAnalysis } = await import("../api/client");

// The node test environment has no DOM, so the dashboard is rendered on the server with the fixture as its first
// data (effects do not run there); the mocked client stands in for the API the dashboard would call next.
function render(tableView: boolean): string {
  return renderToString(
    <RunDashboard runId={fixture.run.id} onBack={() => {}} initial={fixture} initialTableView={tableView} />,
  );
}

const PANEL_TITLES = [
  "Sessions by status",
  "Trust over the study",
  "Competence",
  "Benevolence",
  "Integrity",
  "First to last rating",
  "Baseline check",
  "Slide 1 check",
  "H1 flag",
  "H2 answer",
  "H3 denial",
  "H4a participation",
  "H4b amplification (exploratory)",
  "Actions per slide, checker group",
  "Actions per slide, neutral group",
  "Budget vote by group",
  "Sureness of the vote",
  "Participation per country",
  "Age",
  "Field of expertise",
  "Job level",
  "Education",
  "Cyber security knowledge",
  "AI use",
  "Flags",
  "Audio check",
  "Rating timeouts per point",
  "Duration of finished sessions",
  "H1 change at the flags by cyber security knowledge",
  "H2 change after a reason by AI use",
];

describe("RunDashboard", () => {
  it("renders every section and panel title, in plan order", () => {
    const html = render(false);
    let at = -1;
    for (const t of SECTION_TITLES) {
      const i = html.indexOf(`<h2>${t}</h2>`);
      expect(i, t).toBeGreaterThan(at);
      at = i;
    }
    for (const t of PANEL_TITLES) expect(html, t).toContain(`>${t}</h3>`);
    expect(html).toContain(fixture.run.name);
    expect(html).toContain("Download analysis JSON");
    expect(html).toContain("<svg");
    expect(getRunAnalysis).not.toHaveBeenCalled(); // no effect runs in a server render
  });

  it("shows the reading chips with their tones", () => {
    const html = render(false);
    expect(html).toContain('class="chip chip-success">supports H1 so far');
    expect(html).toContain("chip chip-muted");
    // The first hypothesis card is the research question.
    expect(html.indexOf(">First to last rating</h3>")).toBeLessThan(html.indexOf(">Baseline check</h3>"));
    expect(html).toContain("slides are not independent");
  });

  it("turns every chart into a table in table view", () => {
    const charts = render(false);
    const tables = render(true);
    const svgCount = (s: string) => (s.match(/<svg/g) ?? []).length;
    expect(svgCount(charts)).toBeGreaterThan(20);
    // Only the bar chart that sits beside the country table stays a chart in table view.
    expect(svgCount(tables)).toBe(1);
    expect((tables.match(/<caption>/g) ?? []).length).toBeGreaterThan(svgCount(charts));
    expect(tables).toContain("<caption>Participants per country</caption>");
    expect(tables).toContain("Canada");
  });
});
