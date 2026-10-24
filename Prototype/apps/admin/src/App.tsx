import { useEffect, useState } from "react";
import { logout, onAuthChange } from "./api/client";
import { LoginScreen } from "./login/LoginScreen";
import { RunsTab } from "./runs/RunsTab";
import { ParticipantsTab } from "./participants/ParticipantsTab";
import { ExportTab } from "./export/ExportTab";
import { ConfigsTab } from "./configs/ConfigsTab";
import { localZone } from "./lib/time";

type Tab = "runs" | "participants" | "export" | "configs";
const TABS: { id: Tab; label: string }[] = [
  { id: "runs", label: "Runs" },
  { id: "participants", label: "Participants" },
  { id: "export", label: "Export" },
  { id: "configs", label: "Configs" },
];

export function App() {
  const [loginRequired, setLoginRequired] = useState(false);
  const [tab, setTab] = useState<Tab>("runs");

  useEffect(() => onAuthChange(setLoginRequired), []);

  // The tabs stay mounted under the login overlay, so a request that got 401 resumes after login.
  return (
    <div className="app">
      <header className="header">
        <h1>Study admin</h1>
        <nav className="tabs" role="tablist">
          {TABS.map((t) => (
            <button
              key={t.id}
              role="tab"
              aria-selected={tab === t.id}
              className={tab === t.id ? "tab active" : "tab"}
              onClick={() => setTab(t.id)}
            >
              {t.label}
            </button>
          ))}
        </nav>
        <span className="zone" title="All times are shown in this time zone">
          Times in {localZone}
        </span>
        <button className="link" onClick={() => void logout()}>
          Log out
        </button>
      </header>
      <main className="main">
        {tab === "runs" && <RunsTab />}
        {tab === "participants" && <ParticipantsTab />}
        {tab === "export" && <ExportTab />}
        {tab === "configs" && <ConfigsTab />}
      </main>
      {loginRequired && <LoginScreen />}
    </div>
  );
}
