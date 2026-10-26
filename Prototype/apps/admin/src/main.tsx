import { StrictMode, type ReactNode } from "react";
import { createRoot } from "react-dom/client";
import { App } from "./App";
import type { RunAnalysis } from "./api/types";
import "./styles.css";

const root = createRoot(document.getElementById("root")!);
const render = (node: ReactNode) => root.render(<StrictMode>{node}</StrictMode>);

// DECISION: the fixture preview exists only in the dev server (`npm run dev`, then /admin/#preview=fixture, or
// #preview=fixture&table for the table view); the production build drops this branch and the fixture with it.
if (import.meta.env.DEV && window.location.hash.startsWith("#preview=fixture")) {
  void Promise.all([import("./dashboard/Preview"), import("./dashboard/fixtures/analysis.sample.json")]).then(
    ([m, fx]) => render(<m.DashboardPreview fixture={fx.default as unknown as RunAnalysis} />),
  );
} else {
  render(<App />);
}
