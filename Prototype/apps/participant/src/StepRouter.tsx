// `<StepRouter>` (section 9.2): maps `next.scene` to a scene component. The client never decides the
// next step: every screen is whatever step descriptor the server returned last (POST sessions, any
// event response or error carrying `next`, or GET /next). The clientSeq counter lives in the ApiClient
// and is always set from the last response's `expectedClientSeq`.

import { useCallback, useEffect, useMemo, useRef, useState, type ComponentType } from "react";
import type { Progress, PublicScenePack, Step, StepScene } from "@study/shared";
import type { ApiClient } from "./api/client";
import { ApiError } from "./api/client";
import { StudyContext, type StudyContextValue } from "./study";
import { player } from "./audio/player";
import { DeviceGate } from "./scenes/DeviceGate";
import { Consent } from "./scenes/Consent";
import { AudioCheck } from "./scenes/AudioCheck";
import { PreStart } from "./scenes/PreStart";
import { AudioScene } from "./scenes/AudioScene";
import { PostView } from "./scenes/PostView";
import { Rating } from "./scenes/Rating";
import { Action } from "./scenes/Action";
import { Budget } from "./scenes/Budget";
import { Background } from "./scenes/Background";
import { Debrief, Done } from "./scenes/Debrief";
import styles from "./scenes/scenes.module.css";

type SceneComponent = ComponentType<{ step: Step }>;

export const SCENES: Record<StepScene, SceneComponent> = {
  deviceGate: DeviceGate,
  consent: Consent,
  audioCheck: AudioCheck,
  example: AudioScene,
  preStart: PreStart,
  opening: AudioScene,
  slide: AudioScene,
  postView: PostView,
  rating: Rating,
  action: Action,
  question: AudioScene,
  answer: AudioScene,
  budget: Budget,
  background: Background,
  debrief: Debrief,
  done: Done,
};

export function sceneComponentFor(scene: StepScene | string): SceneComponent | undefined {
  return (SCENES as Record<string, SceneComponent | undefined>)[scene];
}

const MAX_VISIBILITY_EVENTS = 30;

interface Props {
  api: ApiClient;
  pack: PublicScenePack;
  assetBaseUrl: string;
  code: string;
  initialStep: Step;
  initialProgress?: Progress;
  onBlocked: () => void;
}

export function StepRouter({ api, pack, assetBaseUrl, code, initialStep, initialProgress, onBlocked }: Props) {
  const [step, setStep] = useState<Step>(initialStep);
  const [progress, setProgress] = useState<Progress | undefined>(initialProgress);
  const [error, setError] = useState<string | null>(null);
  const [lastDebrief, setLastDebrief] = useState<Step["debrief"]>(initialStep.debrief);

  // Every response with a step descriptor (including side events and errors) updates the screen.
  // Set during render, not in an effect, because child effects (which send events) run first.
  const applyStep = useCallback((s: Step) => {
    setStep(s);
    if (s.debrief) setLastDebrief(s.debrief);
  }, []);
  api.onStep = applyStep;

  // Prefetch the next step's assets into the HTTP cache (9.3).
  useEffect(() => {
    player.prefetch(step.prefetch);
  }, [step.stepId, step.prefetch]);

  // History: no back navigation; state is replaced, never pushed (9.1).
  useEffect(() => {
    try {
      window.history.replaceState(null, "", window.location.pathname);
    } catch {
      /* ignore */
    }
  }, [step.stepId]);

  // Tab switches are logged as side events, at most 30 per session (5.3).
  const visCount = useRef(0);
  useEffect(() => {
    const onVis = () => {
      if (visCount.current >= MAX_VISIBILITY_EVENTS) return;
      visCount.current++;
      void api.send("visibility", { hidden: document.visibilityState === "hidden" }).catch(() => undefined);
    };
    document.addEventListener("visibilitychange", onVis);
    return () => document.removeEventListener("visibilitychange", onVis);
  }, [api]);

  const refresh = useCallback(async () => {
    try {
      const r = await api.next();
      if (r.progress) setProgress(r.progress);
      setError(null);
    } catch (e) {
      setError(messageFor(e));
    }
  }, [api]);

  const reportError = useCallback((e: unknown) => {
    // State errors carry `next`, which the client has already applied; nothing more to show.
    if (e instanceof ApiError && e.next) return;
    setError(messageFor(e));
  }, []);

  const ctx = useMemo<StudyContextValue>(
    () => ({ api, pack, assetBaseUrl, code, progress, refresh, reportError, block: onBlocked, lastDebrief }),
    [api, pack, assetBaseUrl, code, progress, refresh, reportError, onBlocked, lastDebrief],
  );

  const Component = sceneComponentFor(step.scene);
  return (
    <StudyContext.Provider value={ctx}>
      {Component ? <Component key={step.stepId} step={step} /> : <UnknownStep />}
      {error && (
        <div className="error-banner" role="alert">
          <span>{error}</span>
          <button
            type="button"
            className={styles.secondary}
            onClick={() => {
              setError(null);
              void refresh();
            }}
          >
            Try again
          </button>
        </div>
      )}
    </StudyContext.Provider>
  );
}

function UnknownStep() {
  return (
    <main className="center-message">
      <p>Please wait.</p>
    </main>
  );
}

// DECISION: error messages are fixed English text; the config has no wording for them.
export function messageFor(e: unknown): string {
  if (e instanceof ApiError) {
    if (e.code === "network" || e.status === 0) return "The connection was lost. Please check your internet connection.";
    if (e.code === "rate_limited") return "Too many requests. Please wait a moment.";
    if (e.code === "no_session" || e.code === "bad_token") return "Your session could not be found. Please open the study link again.";
  }
  return "Something went wrong.";
}
