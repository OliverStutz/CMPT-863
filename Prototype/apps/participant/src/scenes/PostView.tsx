// The server-timed view after the opening, a slide or an answer (sections 5.2, 9.4): the boardroom stays
// exactly as it was when the clip ended (slide, marks and panel; nobody speaking) with no countdown and no
// screen change; the pause itself is the cue that the rating is coming (Oliver, 2026-10-26: the former
// countdown card and the re-mounted room read as a blip). At `opensInMs` the client asks GET /next and
// renders what comes back. If the server still answers with this step, it asks again every second.

import { useEffect, useRef } from "react";
import type { Step } from "@study/shared";
import { useStudy } from "../study";
import { useDeadline } from "../timers/useDeadline";
import { Stage } from "./stage";
import { visualForStep } from "./visual";

export function PostView({ step }: { step: Step }) {
  const { pack, refresh } = useStudy();
  const visual = visualForStep(pack, step);
  const expired = useRef(false);
  useDeadline(step.opensInMs ?? 0, () => {
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
    <Stage
      sceneKind="postView"
      layout={visual.layout}
      template={visual.template}
      slideNumber={visual.slideNumber}
      profileCard={visual.profileCard}
      overlay={step.overlay}
      clockMs={Infinity}
      speaking={null}
    />
  );
}
