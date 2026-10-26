// The public scene pack, served at the `scenePackUrl` returned by POST /api/v1/runs/{code}/sessions
// ("02 - Architecture.md" section 4.3).
//
// DECISION: section 4.3 lists what the pack contains but not its JSON shape. The simplest reading is
// used: the pack is a subset of StudyConfig with the same field names, from which everything that
// differs between groups or gives the condition away is removed (markings, panel texts, questions,
// answers, action options, preStartText.checker, the debrief, audioKey/marks of every line). The
// client reads only the fields typed below and tolerates any of the optional ones being absent.

import type {
  Block,
  BudgetScene,
  BackgroundScene,
  Choice,
  Colors,
  ConsentSpec,
  Id,
  Labels,
  LayoutTemplate,
  ProfileCard,
  RatingPoint,
  RatingSpec,
  StudyOptions,
  Timing,
} from "./config";

export interface PublicScenePack {
  version: string;
  locale?: "en";
  timing?: Partial<Timing>;
  options?: StudyOptions;
  /** The checker panel's labels (panelTitle, panelIdle, noIssues) are not in the pack; they come with the overlay. */
  labels: Omit<Labels, "panelTitle" | "panelIdle" | "noIssues"> & Partial<Pick<Labels, "panelTitle" | "panelIdle" | "noIssues">>;
  /** Not in the pack (checker-only); kept optional for older packs. */
  colors?: Colors;
  consent: ConsentSpec;
  setup: PublicSetup;
  rating: RatingSpec;
  scenes: PublicScene[];
  /** Idle loops of the silent seats (seat role -> video URL), present when the version has a video manifest (2026-10-25). */
  video?: { seats?: Record<string, string> };
}

export interface PublicSetup {
  instruction: string;
  audioCheck: { prompt: string; options: Choice[] };
  exampleScene: { template?: LayoutTemplate; layout: Block[] };
  preStartText: { both: string[] };
}

export interface PublicOpeningScene {
  kind: "opening";
  id: Id;
  template?: LayoutTemplate;
  layout?: Block[];
  profileCard: ProfileCard;
}
export interface PublicSlideScene {
  kind: "slide";
  id: Id;
  template?: LayoutTemplate;
  layout: Block[];
}
export interface PublicRatingScene {
  kind: "rating";
  id: Id;
  point: RatingPoint;
  after: Id;
}
export interface PublicForSlideScene {
  kind: "action" | "question" | "answer";
  id: Id;
  forSlide: Id;
}

export type PublicScene =
  | PublicOpeningScene
  | PublicSlideScene
  | PublicRatingScene
  | PublicForSlideScene
  | BudgetScene
  | BackgroundScene;
