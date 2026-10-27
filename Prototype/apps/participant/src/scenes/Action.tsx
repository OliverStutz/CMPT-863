// Action screen (2.1; content v1.3 section 7): the options the step descriptor carries for this session
// (three for checker sessions, yes/no for neutral ones; the client does not know which, it renders the
// list). One click is final. At `closesInMs` the client sends action_timeout (15 s).
// The choice is made in the boardroom, with the slide still on the screen and, in checker sessions, the
// checker's marks and panel still shown (the descriptor carries the slide's overlay), so the participant
// decides while looking at what was flagged.

import { useRef, useState } from "react";
import type { Step } from "@study/shared";
import { findScene, useStudy } from "../study";
import { useDeadline } from "../timers/useDeadline";
import { Countdown } from "../timers/Countdown";
import { AutoFocusHeading } from "../ui/AutoFocusHeading";
import { Stage } from "./stage";
import { visualForStep } from "./visual";
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

  const visual = visualForStep(pack, step);
  const title = step.prompt ?? "Your choice";
  const options = step.options ?? [];

  return (
    <Stage
      sceneKind="action"
      layout={visual.layout}
      template={visual.template}
      slideNumber={visual.slideNumber}
      overlay={step.overlay}
      clockMs={Infinity}
      speaking={null}
      dialog={
        <div className={styles.actionCard} data-testid="action-card">
          <AutoFocusHeading className={step.prompt ? styles.actionTitle : "sr-only"}>{title}</AutoFocusHeading>
          <Countdown label={pack.labels.timeLeft} leftMs={left} totalMs={step.closesInMs} />
          <div role="group" aria-label={title} className={styles.actionOptions} style={{ gridTemplateColumns: `repeat(${Math.max(1, options.length)}, minmax(0, 1fr))` }}>
            {options.map((o) => (
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
          {step.note && <p className={styles.actionNote}>{step.note}</p>}
        </div>
      }
    />
  );
}
