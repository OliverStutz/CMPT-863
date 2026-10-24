// Action screen (2.1; content v1.3 section 7): the options the step descriptor carries for this session
// (three for checker sessions, yes/no for neutral ones; the client does not know which, it renders the
// list). One click is final. At `closesInMs` the client sends action_timeout (15 s).

import { useRef, useState } from "react";
import type { Step } from "@study/shared";
import { findScene, useStudy } from "../study";
import { useDeadline } from "../timers/useDeadline";
import { Countdown } from "../timers/Countdown";
import { Screen } from "../ui/Screen";
import styles from "./scenes.module.css";

export function slideIdForStep(step: Step, pack: Parameters<typeof findScene>[0]): string {
  const s = findScene(pack, step.sceneRef);
  if (s && "forSlide" in s) return s.forSlide;
  return step.sceneRef ?? "";
}

export function Action({ step }: { step: Step }) {
  const { api, pack, reportError } = useStudy();
  const [chosen, setChosen] = useState<string | null>(null);
  const done = useRef(false);
  const slideId = slideIdForStep(step, pack);

  const choose = async (actionId: string) => {
    if (done.current) return;
    done.current = true;
    setChosen(actionId);
    try {
      await api.send("action_chosen", { slideId, actionId });
    } catch (e) {
      reportError(e);
    }
  };
  const left = useDeadline(step.closesInMs, () => {
    if (done.current) return;
    done.current = true;
    setChosen("");
    void api.sendUntilOnTime("action_timeout", { slideId }).catch(reportError);
  });

  return (
    <Screen title={step.prompt ?? "Your choice"} titleHidden={!step.prompt}>
      <Countdown label={pack.labels.timeLeft} leftMs={left} totalMs={step.closesInMs} />
      <div role="group" aria-label={step.prompt ?? "Your choice"}>
        {(step.options ?? []).map((o) => (
          <button
            key={o.id}
            type="button"
            className={`${styles.choice} ${chosen === o.id ? styles.choiceSelected : ""}`}
            disabled={chosen !== null}
            aria-pressed={chosen === o.id}
            onClick={() => choose(o.id)}
          >
            {o.label}
            {o.detail && <span className={styles.choiceDetail}>{o.detail}</span>}
          </button>
        ))}
      </div>
      {step.note && <p className={styles.muted}>{step.note}</p>}
    </Screen>
  );
}
