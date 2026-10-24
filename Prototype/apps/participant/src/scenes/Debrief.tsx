// Debrief (2.14, 2.15, 1.25; section 9.2 `<Debrief>`): the paragraphs from the step descriptor (both
// groups, own group, end), the do-not-discuss sentence and the session code.
//
// DECISION: section 5.3 requires at least 5 s on the debrief before `finished` and names no button, so
// the client sends `finished` by itself after 6 s and keeps the debrief on screen afterwards (the "done"
// step that follows renders the same debrief). On a later reload, "done" with no debrief shows a short
// closing line.

import { useEffect, useRef } from "react";
import type { Step } from "@study/shared";
import { useStudy } from "../study";
import { Screen } from "../ui/Screen";
import styles from "./scenes.module.css";

const FINISH_AFTER_MS = 6000;

export function DebriefView({ debrief }: { debrief: NonNullable<Step["debrief"]> }) {
  return (
    <Screen title="Thank you" titleHidden wide>
      {debrief.paragraphs.map((p) => (
        <p key={p} className={styles.lead}>
          {p}
        </p>
      ))}
      <p className={styles.lead}>
        <strong>{debrief.doNotDiscuss}</strong>
      </p>
      <p className={styles.lead}>
        {debrief.sessionCodeLabel ?? "Your session code is"} <span className={styles.code}>{debrief.sessionCode}</span>
      </p>
    </Screen>
  );
}

export function Debrief({ step }: { step: Step }) {
  const { api } = useStudy();
  const sent = useRef(false);
  useEffect(() => {
    const t = setTimeout(() => {
      if (sent.current) return;
      sent.current = true;
      void api.sendUntilOnTime("finished", {}).catch(() => undefined);
    }, FINISH_AFTER_MS);
    return () => clearTimeout(t);
  }, [api]);
  if (!step.debrief) return <Done step={step} />;
  return <DebriefView debrief={step.debrief} />;
}

export function Done({ step }: { step: Step }) {
  const { lastDebrief } = useStudy();
  const d = step.debrief ?? lastDebrief;
  if (d) return <DebriefView debrief={d} />;
  // DECISION: no config text exists for a finished session reopened later; this line is fixed.
  return (
    <Screen title="Thank you">
      <p className={styles.lead}>You have completed this study. You can close this window.</p>
    </Screen>
  );
}
