// One audio step (example, opening, slide, question, answer), sections 5.3, 9.3 and 9.5:
//  1. If the page has no unlocked audio yet (reload) or the step says `restart`, wait for a click on
//     "Continue" (a new user gesture), which primes the one audio element.
//  2. Send scene_started; start playback only when its response arrives.
//  3. `ended` sends audio_ended; on too_early wait retryInMs and send again with the next clientSeq.
//  4. `error`, a `stalled` that does not recover, or no `ended` within duration + 5 s sends
//     audio_problem; "Continue" becomes active after `continueInMs`, then audio_ended with afterProblem.

import { useCallback, useEffect, useRef, useState } from "react";
import type { Step } from "@study/shared";
import { player } from "./player";
import { useStudy } from "../study";

export type AudioPhase = "gesture" | "starting" | "playing" | "problem" | "ending";

export interface AudioStepState {
  phase: AudioPhase;
  /** Audio position in ms for markings; Infinity once the clip is over. */
  clockMs: number;
  /** Remaining ms before "Continue" may be used after a problem; undefined when not in problem. */
  continueInMs?: number;
  /** Click handler for the gesture "Continue" (before playback). */
  startFromGesture: () => void;
  /** Click handler for the "Continue" after an audio problem. */
  continueAfterProblem: () => void;
}

const NO_ENDED_GRACE_MS = 5000;
const STALL_GRACE_MS = 3000;

/** DECISION: Step.sceneRef is used as the `sceneId` of audio events; the stepId is the fallback. */
export function sceneIdOf(step: Step): string {
  return step.sceneRef ?? step.stepId;
}

export function useAudioStep(step: Step): AudioStepState {
  const { api, reportError } = useStudy();
  const sceneId = sceneIdOf(step);
  const needsGesture = !!step.restart || !player.unlocked;
  const [phase, setPhase] = useState<AudioPhase>(needsGesture ? "gesture" : "starting");
  const [clockMs, setClockMs] = useState(0);
  const [continueAt, setContinueAt] = useState<number | undefined>(undefined);
  const [now, setNow] = useState(() => performance.now());
  const phaseRef = useRef(phase);
  phaseRef.current = phase;
  const started = useRef(false);
  const ending = useRef(false);
  const problemSent = useRef(false);

  const sendEnded = useCallback(
    async (afterProblem: boolean) => {
      if (ending.current) return;
      ending.current = true;
      player.settle();
      setPhase("ending");
      setClockMs(Infinity);
      try {
        await api.sendUntilOnTime("audio_ended", afterProblem ? { sceneId, afterProblem: true } : { sceneId });
      } catch (e) {
        ending.current = false;
        reportError(e);
      }
    },
    [api, reportError, sceneId],
  );

  const problem = useCallback(
    async (kind: "error" | "stalled" | "noEnded") => {
      if (problemSent.current || ending.current) return;
      problemSent.current = true;
      setPhase("problem");
      try {
        const res = await api.send("audio_problem", { sceneId, kind, atMs: Math.round(player.currentTimeMs) });
        setContinueAt(performance.now() + Math.max(0, res.continueInMs ?? 0));
      } catch (e) {
        setContinueAt(performance.now());
        reportError(e);
      }
    },
    [api, reportError, sceneId],
  );

  const start = useCallback(async () => {
    if (started.current) return;
    started.current = true;
    setPhase("starting");
    if (step.audioUrl) player.load(step.audioUrl);
    try {
      await api.send("scene_started", { sceneId });
    } catch (e) {
      started.current = false;
      reportError(e);
      return;
    }
    try {
      await player.playFromStart();
      setPhase("playing");
    } catch {
      // Playback refused (autoplay lost) or failed: ask for a gesture once, then report a problem.
      if (!player.unlocked) {
        started.current = false;
        setPhase("gesture");
      } else {
        void problem("error");
      }
    }
  }, [api, problem, reportError, sceneId, step.audioUrl]);

  // Load the clip at once; start without a gesture when audio is already unlocked.
  useEffect(() => {
    if (step.audioUrl) player.load(step.audioUrl);
    if (!needsGesture) void start();
    return () => player.stop();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // Element events.
  useEffect(() => {
    const el = player.element;
    let stallTimer: ReturnType<typeof setTimeout> | undefined;
    const onEnded = () => {
      if (phaseRef.current === "playing" || phaseRef.current === "problem") void sendEnded(phaseRef.current === "problem");
    };
    const onError = () => {
      if (phaseRef.current === "playing" || phaseRef.current === "starting") void problem("error");
    };
    const onStalled = () => {
      if (phaseRef.current !== "playing") return;
      const at = el.currentTime;
      clearTimeout(stallTimer);
      stallTimer = setTimeout(() => {
        if (phaseRef.current === "playing" && el.currentTime === at && !el.ended) void problem("stalled");
      }, STALL_GRACE_MS);
    };
    el.addEventListener("ended", onEnded);
    el.addEventListener("error", onError);
    el.addEventListener("stalled", onStalled);
    return () => {
      clearTimeout(stallTimer);
      el.removeEventListener("ended", onEnded);
      el.removeEventListener("error", onError);
      el.removeEventListener("stalled", onStalled);
    };
  }, [problem, sendEnded]);

  // Clock for markings, and the no-`ended` watchdog (duration + 5 s of wall time from the start).
  useEffect(() => {
    if (phase !== "playing") return;
    const t0 = performance.now();
    const id = setInterval(() => {
      setClockMs(player.currentTimeMs);
      const d = player.durationMs;
      if (Number.isFinite(d) && performance.now() - t0 > d + NO_ENDED_GRACE_MS) void problem("noEnded");
    }, 100);
    return () => clearInterval(id);
  }, [phase, problem]);

  // Countdown to "Continue" after a problem.
  useEffect(() => {
    if (phase !== "problem" || continueAt === undefined) return;
    const id = setInterval(() => setNow(performance.now()), 250);
    return () => clearInterval(id);
  }, [phase, continueAt]);

  const startFromGesture = useCallback(() => {
    player.primeFromGesture(step.audioUrl);
    void start();
  }, [start, step.audioUrl]);

  const continueAfterProblem = useCallback(() => {
    void sendEnded(true);
  }, [sendEnded]);

  return {
    phase,
    clockMs: phase === "problem" || phase === "ending" ? Math.max(clockMs, player.currentTimeMs) : clockMs,
    continueInMs: phase === "problem" && continueAt !== undefined ? Math.max(0, continueAt - now) : undefined,
    startFromGesture,
    continueAfterProblem,
  };
}
