// Budget decision (1.5, 2.12; content v1.3 section 8): approve or reject, then "How sure are you?" 0 to
// 10 as clickable numbers, nothing preselected, 30 s countdown; budget_submitted, or budget_timeout with
// what was given.
// DECISION: the submit button uses `labels.ratingSubmit` ("Next"); the config has no budget button label.

import { useRef, useState } from "react";
import type { BudgetScene, Step } from "@study/shared";
import { findScene, useStudy } from "../study";
import { useDeadline } from "../timers/useDeadline";
import { Countdown } from "../timers/Countdown";
import { ScaleRow } from "../ui/ScaleRow";
import { Screen } from "../ui/Screen";
import styles from "./scenes.module.css";

export function Budget({ step }: { step: Step }) {
  const { api, pack, reportError } = useStudy();
  const scene = (findScene(pack, step.sceneRef) ?? pack.scenes.find((s) => s.kind === "budget")) as BudgetScene;
  const [optionId, setOptionId] = useState<string | null>(null);
  const [confidence, setConfidence] = useState<number | null>(null);
  const state = useRef({ optionId, confidence });
  state.current = { optionId, confidence };
  const [sent, setSent] = useState(false);
  const sentRef = useRef(false);

  const finish = async (timeout: boolean) => {
    if (sentRef.current) return;
    sentRef.current = true;
    setSent(true);
    const v = { ...state.current };
    try {
      if (timeout) await api.sendUntilOnTime("budget_timeout", { partial: v });
      else await api.send("budget_submitted", v);
    } catch (e) {
      reportError(e);
    }
  };
  const left = useDeadline(step.closesInMs, () => void finish(true));

  return (
    <Screen title="Budget decision" titleHidden wide>
      <Countdown label={pack.labels.timeLeft} leftMs={left} totalMs={step.closesInMs} />
      <fieldset className={styles.question} disabled={sent}>
        <legend className={styles.lead}>{scene.prompt}</legend>
        <div className={styles.radioList}>
          {scene.options.map((o) => (
            <label key={o.id} className={styles.radioItem}>
              <input type="radio" name="budget-option" checked={optionId === o.id} onChange={() => setOptionId(o.id)} />
              <span>{o.label}</span>
            </label>
          ))}
        </div>
      </fieldset>
      <ScaleRow
        name="confidence"
        legend={scene.confidence.label}
        min={scene.confidence.scale.min}
        max={scene.confidence.scale.max}
        minLabel={scene.confidence.scale.minLabel}
        maxLabel={scene.confidence.scale.maxLabel}
        pointLabels={scene.confidence.scale.pointLabels}
        variant="ends"
        value={confidence}
        onChange={(v) => setConfidence(v as number)}
        disabled={sent}
      />
      <div className={styles.actions}>
        <button type="button" className={styles.primary} onClick={() => void finish(false)} disabled={sent}>
          {pack.labels.ratingSubmit}
        </button>
      </div>
    </Screen>
  );
}
