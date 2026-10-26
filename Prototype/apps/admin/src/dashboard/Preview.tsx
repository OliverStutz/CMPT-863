import type { RunAnalysis } from "../api/types";
import { filterParticipants } from "./model";
import { RunDashboard } from "./RunDashboard";

/**
 * Dev-only preview: `#preview=fixture` renders the dashboard from the generated fixture, without the API or the
 * login screen, so it can be screenshotted. Only the group filter acts (on the participant list, as the server does).
 */
export function DashboardPreview({ fixture }: { fixture: RunAnalysis }) {
  return (
    <div className="app">
      <header className="header">
        <h1>Study admin</h1>
        <span className="zone">Preview from the fixture</span>
      </header>
      <main className="main">
        <RunDashboard
          runId={fixture.run.id}
          onBack={() => (window.location.hash = "")}
          initial={fixture}
          initialTableView={window.location.hash.includes("&table")}
          load={async (_id, q) => filterParticipants(fixture, q.group ?? "both")}
          notice="Preview: data from src/dashboard/fixtures/analysis.sample.json. The participant and flag filters do not change it."
        />
      </main>
    </div>
  );
}
