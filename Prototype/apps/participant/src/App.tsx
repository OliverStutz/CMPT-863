// Routes (section 9.1): /s/:code is the study, /s/:code/blocked the device gate message, anything else
// a neutral "study link needed" page. No router library: the server, not the URL, decides the screen.

import { useCallback, useState } from "react";
import { StudyRoute } from "./routes/StudyRoute";
import { BlockedPage } from "./routes/BlockedPage";
import { NoLinkPage } from "./routes/NoLinkPage";

export type Route = { kind: "study"; code: string } | { kind: "blocked"; code: string } | { kind: "none" };

export function parseRoute(pathname: string): Route {
  const m = /^\/s\/([^/]+)\/?(blocked)?\/?$/.exec(pathname);
  if (!m) return { kind: "none" };
  const code = decodeURIComponent(m[1]);
  return m[2] ? { kind: "blocked", code } : { kind: "study", code };
}

export function App() {
  const [route, setRoute] = useState<Route>(() => parseRoute(window.location.pathname));
  const onBlocked = useCallback((code: string) => {
    try {
      window.history.replaceState(null, "", `/s/${encodeURIComponent(code)}/blocked`);
    } catch {
      /* ignore */
    }
    setRoute({ kind: "blocked", code });
  }, []);

  if (route.kind === "study") return <StudyRoute code={route.code} onBlocked={onBlocked} />;
  if (route.kind === "blocked") return <BlockedPage code={route.code} />;
  return <NoLinkPage />;
}
