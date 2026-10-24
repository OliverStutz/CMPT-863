// "The meeting is about to start" (section 9.2 `<PreStart>`): the lines from the step descriptor, which
// carries the checker sentence only for checker sessions (1.1); the button sends setup_done. The click
// also primes the audio element for the opening clip.

import { useState } from "react";
import type { Step } from "@study/shared";
import { useStudy } from "../study";
import { player } from "../audio/player";
import { Screen } from "../ui/Screen";
import styles from "./scenes.module.css";

export function PreStart({ step }: { step: Step }) {
  const { api, pack, reportError } = useStudy();
  const [busy, setBusy] = useState(false);
  // DECISION: if the descriptor carries no preStartText, the public pack's `both` lines are shown.
  const lines = step.preStartText ?? pack.setup.preStartText.both;

  const start = async () => {
    setBusy(true);
    player.primeFromGesture(step.prefetch?.[0]);
    try {
      await api.sendUntilOnTime("setup_done", {});
    } catch (e) {
      setBusy(false);
      reportError(e);
    }
  };

  return (
    <Screen title={lines[0] ?? ""}>
      {lines.slice(1).map((l) => (
        <p key={l} className={styles.lead}>
          {l}
        </p>
      ))}
      <div className={styles.actions}>
        <button type="button" className={styles.primary} onClick={start} disabled={busy}>
          {pack.labels.startMeeting}
        </button>
      </div>
    </Screen>
  );
}
