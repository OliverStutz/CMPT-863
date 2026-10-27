// Rating (section 9.2 `<RatingPanel>`; content v1.3 section 6): shown on the meeting-room screen inside the
// boardroom, not on a page of its own, so the room and the seats stay where they are (Oliver, 2026-10-26,
// 2.17). The instruction and four items in fixed order (1.23). Trust is a row of eleven numbers 0 to 10
// with its two end labels; the three METI items are 7-point antonym rows (left word, seven numbers, right
// word) (1.13). Nothing is preselected; an item not clicked is sent as null (2.8). One submit button, and
// at `closesInMs` (20 s since 2026-10-26) the client sends rating_timeout with whatever was given (2.2).
// The point comes from the descriptor.

import { useRef, useState } from "react";
import type { RatingItemId, RatingValues, Step } from "@study/shared";
import { useStudy } from "../study";
import { useDeadline } from "../timers/useDeadline";
import { Countdown } from "../timers/Countdown";
import { ScaleRow } from "../ui/ScaleRow";
import { AutoFocusHeading } from "../ui/AutoFocusHeading";
import { Stage } from "./stage";
import styles from "./scenes.module.css";

const EMPTY: RatingValues = { trust: null, competence: null, benevolence: null, integrity: null };

export function Rating({ step }: { step: Step }) {
  const { api, pack, reportError } = useStudy();
  const [values, setValues] = useState<RatingValues>(EMPTY);
  const valuesRef = useRef(values);
  valuesRef.current = values;
  const [sent, setSent] = useState(false);
  const sentRef = useRef(false);
  const point = step.point ?? "";

  // DECISION: if the rating window has not opened yet (opensInMs > 0), the items wait until it opens.
  const opensLeft = useDeadline(step.opensInMs && step.opensInMs > 0 ? step.opensInMs : undefined);
  const open = opensLeft === undefined || opensLeft <= 0;

  const finish = async (timeout: boolean) => {
    if (sentRef.current) return;
    sentRef.current = true;
    setSent(true);
    const v = valuesRef.current;
    try {
      if (timeout) await api.sendUntilOnTime("rating_timeout", { point, partial: { ...v } });
      else await api.send("rating_submitted", { point, ...v });
    } catch (e) {
      reportError(e);
    }
  };

  const closesLeft = useDeadline(step.closesInMs, () => void finish(true));
  const set = (id: RatingItemId, v: number) => setValues((old) => ({ ...old, [id]: v }));
  const items = pack.rating.items;

  const form = (
    <div className={styles.ratingScreen} data-testid="rating-screen">
      <AutoFocusHeading className="sr-only">Rating</AutoFocusHeading>
      <Countdown label={pack.labels.timeLeft} leftMs={closesLeft} totalMs={step.closesInMs} />
      <p className={styles.lead}>{pack.rating.instruction}</p>
      {items.map((item) => {
        const antonym = !item.label;
        return (
          <ScaleRow
            key={item.id}
            name={item.id}
            legend={item.label ?? `${item.scale.minLabel} to ${item.scale.maxLabel}`}
            legendHidden={antonym}
            min={item.scale.min}
            max={item.scale.max}
            minLabel={item.scale.minLabel}
            maxLabel={item.scale.maxLabel}
            pointLabels={item.scale.pointLabels}
            variant={antonym ? "antonym" : "ends"}
            value={values[item.id]}
            onChange={(v) => set(item.id, v as number)}
            disabled={sent || !open}
          />
        );
      })}
      <div className={`${styles.actions} ${styles.ratingFoot}`}>
        <button type="button" className={styles.primary} onClick={() => void finish(false)} disabled={sent || !open}>
          {pack.labels.ratingSubmit}
        </button>
      </div>
    </div>
  );

  // The rating takes the meeting-room screen; rating steps carry no overlay, so the checker panel is away
  // while the participant rates, as on the former rating page.
  return <Stage sceneKind="rating" clockMs={Infinity} speaking={null} screenContent={form} />;
}
