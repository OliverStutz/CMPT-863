// The server-timed view after the opening, a slide or an answer (sections 5.2, 9.4): the scene stays on
// screen with a visible countdown from `opensInMs`; at zero the client asks GET /next and renders what
// comes back. If the server still answers with this step, it asks again every second.

import { useEffect, useRef } from "react";
import type { Step } from "@study/shared";
import { useStudy } from "../study";
import { useDeadline } from "../timers/useDeadline";
import { Countdown } from "../timers/Countdown";
import { BoardroomScene } from "./BoardroomScene";
import { visualForStep } from "./visual";

export function PostView({ step }: { step: Step }) {
  const { pack, refresh } = useStudy();
  const visual = visualForStep(pack, step);
  const total = useRef(step.opensInMs);
  const expired = useRef(false);
  const left = useDeadline(step.opensInMs ?? 0, () => {
    expired.current = true;
    void refresh();
  });

  useEffect(() => {
    const id = setInterval(() => {
      if (expired.current) void refresh();
    }, 1000);
    return () => clearInterval(id);
  }, [refresh]);

  return (
    <BoardroomScene
      layout={visual.layout}
      template={visual.template}
      slideNumber={visual.slideNumber}
      profileCard={visual.profileCard}
      overlay={step.overlay}
      clockMs={Infinity}
      speaking={null}
      controls={<Countdown label={pack.labels.timeLeft} leftMs={left} totalMs={total.current} />}
    />
  );
}
