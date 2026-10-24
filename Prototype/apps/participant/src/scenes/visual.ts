import type { Block, LayoutTemplate, ProfileCard, PublicScenePack, Step } from "@study/shared";
import { findScene } from "../study";
import type { Speaker } from "./BoardroomScene";

export interface Visual {
  layout?: Block[];
  template?: LayoutTemplate;
  slideNumber?: number;
  profileCard?: ProfileCard;
}

/** Slide number n = position among the slide scenes of the pack (1, 2, 3, ...). */
export function slideNumberOf(pack: PublicScenePack, slideId: string): number | undefined {
  const slides = pack.scenes.filter((s) => s.kind === "slide");
  const i = slides.findIndex((s) => s.id === slideId);
  return i >= 0 ? i + 1 : undefined;
}

/**
 * What the presentation screen shows for a scene id: the slide itself, the slide an action, question or
 * answer belongs to, the slide or answer a rating follows, or the opening.
 */
export function visualFor(pack: PublicScenePack, sceneRef: string | undefined, depth = 0): Visual {
  const s = findScene(pack, sceneRef);
  if (!s || depth > 4) return {};
  switch (s.kind) {
    case "slide":
      return { layout: s.layout, template: s.template, slideNumber: slideNumberOf(pack, s.id) };
    case "opening":
      return { layout: s.layout, template: s.template, profileCard: s.profileCard };
    case "action":
    case "question":
    case "answer":
      return visualFor(pack, s.forSlide, depth + 1);
    case "rating":
      return visualFor(pack, s.after, depth + 1);
    default:
      return {};
  }
}

export function visualForStep(pack: PublicScenePack, step: Step): Visual {
  if (step.scene === "example") {
    return { layout: pack.setup.exampleScene.layout, template: pack.setup.exampleScene.template };
  }
  return visualFor(pack, step.sceneRef);
}

// DECISION: the step descriptor does not say which voice speaks. The lit portrait is chosen from the
// step kind only: example -> chair; opening -> chair (one joined clip, the CISO's closing greeting has
// no offset in the descriptor); slide and answer -> CISO; question -> none, because the client would
// otherwise have to map its own action choice to a voice, which is the server's decision.
export function speakerFor(step: Step): Speaker {
  switch (step.scene) {
    case "example":
    case "opening":
      return "chair";
    case "slide":
    case "answer":
      return "ciso";
    default:
      return null;
  }
}
