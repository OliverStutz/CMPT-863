// Example, opening, slide, question and answer steps: the boardroom with one clip (sections 5.3, 9.2).

import type { Step } from "@study/shared";
import { useAudioStep, sceneIdOf } from "../audio/useAudioStep";
import { useStudy } from "../study";
import { BoardroomScene } from "./BoardroomScene";
import { speakerFor, visualForStep } from "./visual";
import styles from "./scenes.module.css";

export function AudioScene({ step }: { step: Step }) {
  const { pack, api } = useStudy();
  const audio = useAudioStep(step);
  const visual = visualForStep(pack, step);
  const sceneId = sceneIdOf(step);
  const playing = audio.phase === "playing";

  // Markings follow the audio clock on the slide step; on question and answer steps they are already
  // shown in full. DECISION: the profile card is shown for the whole opening clip, because the joined
  // clip carries no offset for `showFromLineId`.
  const clockMs = step.scene === "slide" ? audio.clockMs : Infinity;

  const onMarkingShown =
    step.scene === "slide"
      ? (markingId: string) => {
          void api.send("marking_shown", { sceneId, markingId }).catch(() => undefined);
        }
      : undefined;

  let controls = null;
  if (audio.phase === "gesture") {
    controls = (
      <button type="button" className={styles.primary} onClick={audio.startFromGesture} autoFocus>
        {pack.labels.continue}
      </button>
    );
  } else if (audio.phase === "problem") {
    const wait = audio.continueInMs;
    controls = (
      <button
        type="button"
        className={styles.primary}
        onClick={audio.continueAfterProblem}
        disabled={wait === undefined || wait > 0}
        autoFocus
      >
        {pack.labels.continue}
        {wait !== undefined && wait > 0 ? ` (${Math.ceil(wait / 1000)})` : ""}
      </button>
    );
  }

  return (
    <>
      <BoardroomScene
        layout={visual.layout}
        template={visual.template}
        slideNumber={visual.slideNumber}
        profileCard={step.scene === "opening" ? visual.profileCard : undefined}
        overlay={step.overlay}
        clockMs={clockMs}
        speaking={playing ? speakerFor(step) : null}
        onMarkingShown={onMarkingShown}
        controls={controls}
      />
      {step.transcript && (
        <div className="sr-only" aria-live="off">
          {step.transcript}
        </div>
      )}
    </>
  );
}
