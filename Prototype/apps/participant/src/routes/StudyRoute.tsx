// Landing for /s/:code: creates the session (or resumes it, if this browser's cookie belongs to an
// unfinished session of the run), loads the public scene pack, then asks GET /next and hands the step
// to the StepRouter (sections 8.1, 9.5).

import { useEffect, useRef, useState } from "react";
import type { Progress, PublicScenePack, Step } from "@study/shared";
import { ApiClient, ApiError } from "../api/client";
import { StepRouter, messageFor } from "../StepRouter";

type State =
  | { kind: "loading" }
  | { kind: "error"; message: string }
  | { kind: "ready"; pack: PublicScenePack; assetBaseUrl: string; step: Step; progress?: Progress };

export interface StudyRouteProps {
  code: string;
  onBlocked: (code: string) => void;
  /** For tests. */
  api?: ApiClient;
}

/**
 * Test links: /s/<code>?group=checker|neutral&denial=<slide id>. The server honours these only on a run
 * the admin marked for test overrides and flags such sessions forced_allocation; on any other run they
 * are ignored (2026-10-25).
 */
export function testOverridesFromLocation(search: string = window.location.search): { group?: string; denialSlide?: string } | undefined {
  const q = new URLSearchParams(search);
  const group = q.get("group") ?? undefined;
  const denialSlide = q.get("denial") ?? undefined;
  return group || denialSlide ? { group, denialSlide } : undefined;
}

export function StudyRoute({ code, onBlocked, api: injected }: StudyRouteProps) {
  const apiRef = useRef<ApiClient>(injected ?? new ApiClient());
  const [state, setState] = useState<State>({ kind: "loading" });
  const started = useRef(false);

  useEffect(() => {
    if (started.current) return;
    started.current = true;
    const api = apiRef.current;
    (async () => {
      try {
        const created = await api.createSession(code, testOverridesFromLocation());
        const pack = await api.getJson<PublicScenePack>(created.scenePackUrl);
        const next = await api.next();
        setState({ kind: "ready", pack, assetBaseUrl: created.assetBaseUrl, step: next.next, progress: next.progress });
      } catch (e) {
        setState({ kind: "error", message: landingMessage(e) });
      }
    })();
  }, [code]);

  if (state.kind === "loading") {
    return (
      <main className="center-message" aria-busy="true">
        <p>Loading...</p>
      </main>
    );
  }
  if (state.kind === "error") {
    return (
      <main className="center-message">
        <p role="alert">{state.message}</p>
      </main>
    );
  }
  return (
    <StepRouter
      api={apiRef.current}
      pack={state.pack}
      assetBaseUrl={state.assetBaseUrl}
      code={code}
      initialStep={state.step}
      initialProgress={state.progress}
      onBlocked={() => onBlocked(code)}
    />
  );
}

// DECISION: fixed English wording for link errors; the config has none.
function landingMessage(e: unknown): string {
  if (e instanceof ApiError) {
    if (e.code === "unknown_run" || e.status === 404) return "This study link is not valid. Please check the link you received.";
    if (e.code === "run_closed" || e.status === 410) return "This study is closed. Thank you for your interest.";
  }
  return messageFor(e);
}
