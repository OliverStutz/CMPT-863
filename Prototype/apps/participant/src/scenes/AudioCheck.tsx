// Audio check (2.4, section 9.2 `<AudioCheck>`): play button, then play again as often as wanted (each
// play sends audio_check_played), then the four word choices; one choice is final, no feedback. The
// first click is the user gesture that unlocks the one audio element (9.3).

import { useState } from "react";
import type { Step } from "@study/shared";
import { useStudy } from "../study";
import { player } from "../audio/player";
import { Screen } from "../ui/Screen";
import styles from "./scenes.module.css";

export function AudioCheck({ step }: { step: Step }) {
  const { api, pack, reportError } = useStudy();
  const [plays, setPlays] = useState(0);
  const [answer, setAnswer] = useState<string | null>(null);
  const ac = pack.setup.audioCheck;

  const play = () => {
    if (step.audioUrl) player.load(step.audioUrl);
    void player.playFromStart().catch(() => undefined);
    setPlays((n) => n + 1);
    void api.send("audio_check_played", {}).catch(reportError);
  };

  const choose = async (choiceId: string) => {
    if (answer) return;
    setAnswer(choiceId);
    player.stop();
    try {
      await api.send("audio_check_answered", { choiceId });
    } catch (e) {
      reportError(e);
    }
  };

  // DECISION: the setup instruction is one config string; it is split into sentences on screen, as in
  // the content document's screen text.
  const sentences = pack.setup.instruction.split(/(?<=\.)\s+/);

  return (
    <Screen title="Sound check" titleHidden>
      {sentences.map((s) => (
        <p key={s} className={styles.lead}>
          {s}
        </p>
      ))}
      <div className={styles.actions}>
        <button type="button" className={styles.primary} onClick={play} disabled={!!answer}>
          {plays === 0 ? pack.labels.playTestSound : pack.labels.replayTestSound}
        </button>
      </div>
      <p className={styles.subheading} id="audio-check-prompt">
        {ac.prompt}
      </p>
      <div className={styles.choiceGrid} role="group" aria-labelledby="audio-check-prompt">
        {ac.options.map((o) => (
          <button
            key={o.id}
            type="button"
            className={`${styles.choice} ${answer === o.id ? styles.choiceSelected : ""}`}
            disabled={plays === 0 || !!answer}
            aria-pressed={answer === o.id}
            onClick={() => choose(o.id)}
          >
            {o.label}
          </button>
        ))}
      </div>
    </Screen>
  );
}
