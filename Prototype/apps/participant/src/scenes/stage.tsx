// One boardroom for all boardroom steps (Oliver, 2026-10-26: every step change re-mounted the room, the
// seats and their idle loops, which showed as a blip between a clip, the pause and the rating).
//
// The router keeps one `<BoardroomHost>` mounted across consecutive boardroom steps (example, opening,
// slide, postView, rating, action, question, answer). A step component describes what the stage shows
// with `<Stage .../>` instead of rendering its own `<BoardroomScene>`; the host renders the scene from
// the last description. The description is pushed in a layout effect, so the stage is updated before the
// browser paints and never shows a stale step. Rendered without a host (a scene component on its own,
// as in the unit tests), `<Stage>` draws the boardroom in place.

import { createContext, useContext, useLayoutEffect, useState, type ReactNode } from "react";
import { BoardroomScene, type BoardroomProps } from "./BoardroomScene";

const StageContext = createContext<((p: BoardroomProps) => void) | null>(null);

const EMPTY: BoardroomProps = { clockMs: Infinity, speaking: null };

export function BoardroomHost({ children }: { children: ReactNode }) {
  const [props, setProps] = useState<BoardroomProps>(EMPTY);
  return (
    <StageContext.Provider value={setProps}>
      <BoardroomScene {...props} />
      {children}
    </StageContext.Provider>
  );
}

export function Stage(p: BoardroomProps) {
  const set = useContext(StageContext);
  useLayoutEffect(() => {
    if (set) set(p);
  });
  if (set) return null;
  return <BoardroomScene {...p} />;
}
