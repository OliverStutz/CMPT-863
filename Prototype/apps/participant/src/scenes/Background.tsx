// Background questions (1.9, 2.3, 2.12, 2.14; content v1.3 section 9): every answerable item with
// "Prefer not to say" (added by the renderer), the hidden honeypot field, 30 s countdown;
// background_submitted, or background_timeout with what was given.
//
// DECISION: section 5.3 does not fix the shape of `items`. It is sent as { [itemId]: value }, where the
// value is the chosen option id (single), the number (scale), the ISO 3166-1 alpha-2 code (country), or
// "preferNotToSay", and null when untouched. The free text of an "Other (please specify)" option is
// sent under "<itemId>.otherText". The honeypot value goes in `hp`, never in `items`.
// DECISION: the submit button uses `labels.ratingSubmit` ("Next"); the config has no label of its own
// for the budget and background forms.

import { useRef, useState } from "react";
import { PREFER_NOT_TO_SAY, type BackgroundAnswers, type BackgroundScene, type Step } from "@study/shared";
import { findScene, useStudy } from "../study";
import { useDeadline } from "../timers/useDeadline";
import { Countdown } from "../timers/Countdown";
import { ScaleRow } from "../ui/ScaleRow";
import { Screen } from "../ui/Screen";
import { countries } from "../ui/countries";
import styles from "./scenes.module.css";

export function Background({ step }: { step: Step }) {
  const { api, pack, reportError } = useStudy();
  const scene = (findScene(pack, step.sceneRef) ?? pack.scenes.find((s) => s.kind === "background")) as BackgroundScene;
  const answerable = scene.items.filter((i) => i.kind !== "honeypot");
  const honeypot = scene.items.find((i) => i.kind === "honeypot");

  const [answers, setAnswers] = useState<BackgroundAnswers>(() =>
    Object.fromEntries(answerable.map((i) => [i.id, null])),
  );
  const [hp, setHp] = useState("");
  const state = useRef({ answers, hp });
  state.current = { answers, hp };
  const [sent, setSent] = useState(false);
  const sentRef = useRef(false);
  const pnts = scene.preferNotToSayLabel;

  const set = (key: string, v: string | number | null) => setAnswers((a) => ({ ...a, [key]: v }));

  const finish = async (timeout: boolean) => {
    if (sentRef.current) return;
    sentRef.current = true;
    setSent(true);
    const { answers: items, hp: h } = state.current;
    try {
      if (timeout) await api.sendUntilOnTime("background_timeout", { partial: { ...items }, hp: h });
      else await api.send("background_submitted", { items: { ...items }, hp: h });
    } catch (e) {
      reportError(e);
    }
  };
  const left = useDeadline(step.closesInMs, () => void finish(true));

  return (
    <Screen title="Background questions" titleHidden wide>
      <Countdown label={pack.labels.timeLeft} leftMs={left} totalMs={step.closesInMs} />
      <p className={styles.lead}>{scene.intro}</p>
      {answerable.map((item) => {
        if (item.kind === "single") {
          return (
            <fieldset key={item.id} className={styles.question} disabled={sent}>
              <legend>{item.label}</legend>
              <div className={styles.radioList}>
                {[...item.options, { id: PREFER_NOT_TO_SAY, label: pnts }].map((o) => (
                  <div key={o.id}>
                    <label className={styles.radioItem}>
                      <input
                        type="radio"
                        name={`bg-${item.id}`}
                        checked={answers[item.id] === o.id}
                        onChange={() => set(item.id, o.id)}
                      />
                      <span>{o.label}</span>
                    </label>
                    {item.otherOptionId === o.id && answers[item.id] === o.id && (
                      <input
                        type="text"
                        className={`${styles.textInput} ${styles.otherInput}`}
                        aria-label={o.label}
                        maxLength={200}
                        value={(answers[`${item.id}.otherText`] as string | null) ?? ""}
                        onChange={(e) => set(`${item.id}.otherText`, e.target.value)}
                      />
                    )}
                  </div>
                ))}
              </div>
            </fieldset>
          );
        }
        if (item.kind === "scale") {
          return (
            <ScaleRow
              key={item.id}
              name={item.id}
              legend={item.label}
              min={item.scale.min}
              max={item.scale.max}
              minLabel={item.scale.minLabel}
              maxLabel={item.scale.maxLabel}
              pointLabels={item.scale.pointLabels ?? {}}
              variant="ends"
              value={answers[item.id]}
              onChange={(v) => set(item.id, v)}
              extras={[{ value: PREFER_NOT_TO_SAY, label: pnts }]}
              disabled={sent}
            />
          );
        }
        // country
        return (
          <div key={item.id} className={styles.question}>
            <label htmlFor={`bg-${item.id}`} style={{ display: "block", fontWeight: 600, fontSize: "1.08rem", marginBottom: "0.6rem" }}>
              {item.label}
            </label>
            <select
              id={`bg-${item.id}`}
              className={styles.select}
              disabled={sent}
              value={(answers[item.id] as string | null) ?? ""}
              onChange={(e) => set(item.id, e.target.value === "" ? null : e.target.value)}
            >
              <option value="" />
              <option value={PREFER_NOT_TO_SAY}>{pnts}</option>
              {countries().map((c) => (
                <option key={c.code} value={c.code}>
                  {c.name}
                </option>
              ))}
            </select>
          </div>
        );
      })}
      {honeypot && (
        <div className={styles.honeypot} aria-hidden="true">
          <label htmlFor={`hp-${honeypot.id}`}>{honeypot.label}</label>
          <input
            id={`hp-${honeypot.id}`}
            name={honeypot.id}
            type="text"
            tabIndex={-1}
            autoComplete="off"
            value={hp}
            onChange={(e) => setHp(e.target.value)}
          />
        </div>
      )}
      <div className={styles.actions}>
        <button type="button" className={styles.primary} onClick={() => void finish(false)} disabled={sent}>
          {pack.labels.ratingSubmit}
        </button>
      </div>
    </Screen>
  );
}
