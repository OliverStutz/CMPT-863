// Study config types, hand-written from "02 - Architecture.md" section 4.1 and
// content/schema/study-config.schema.json. Keys starting with "_" are author notes:
// the schema accepts them and no tool, server or client reads them, so they are not typed here.

export type Ms = number; // milliseconds, integer >= 0
export type Id = string; // ^[A-Za-z0-9][A-Za-z0-9._-]*$, unique within a config
export type Hex = string; // "#RRGGBB"
export type Group = "checker" | "neutral"; // server only; the participant client never receives it
export type RichText = string | Span[];
export interface Span {
  id: Id;
  text: string;
}
export interface Choice {
  id: Id;
  label: string;
}

export interface StudyConfig {
  version: string;
  locale: "en";
  timing: Timing;
  voices: Voices;
  options: StudyOptions;
  labels: Labels;
  colors: Colors;
  consent: ConsentSpec;
  setup: SetupSpec;
  rating: RatingSpec;
  actions: ActionSpec;
  scenes: Scene[];
}

export interface Timing {
  postAudioViewMs: Ms;
  postAnswerViewMs: Ms;
  ratingWindowMs: Ms;
  actionWindowMs: Ms;
  formWindowMs: Ms;
  /** Background questions screen; 0 = no time limit (the default since 2026-10-25). Absent: formWindowMs. */
  backgroundWindowMs?: Ms;
  linePauseMs: Ms;
  audioEndToleranceMs: Ms;
  latencyAllowanceMs: Ms;
  abandonAfterMs: Ms;
  msPerWordEstimate?: Ms;
}

export interface Voices {
  ciso: VoiceSpec;
  chair: VoiceSpec;
  member: VoiceSpec;
}
export interface VoiceSpec {
  provider: "google-tts";
  name: string;
  languageCode: "en-US";
  speakingRate?: number;
  pitch?: number;
}

export interface StudyOptions {
  deviceGate: { minWidthPx: number; allowPhones: boolean };
}

export interface Labels {
  panelTitle: string;
  noIssues: string;
  panelIdle: string;
  ratingSubmit: string;
  timeLeft: string;
  continue: string;
  playTestSound: string;
  replayTestSound: string;
  startMeeting: string;
}
export interface Colors {
  flag: Hex;
  noIssues: Hex;
}

export interface ConsentSpec {
  version: string;
  pdfAssetKey: string;
  checkboxLabel: string;
  acceptLabel: string;
}

// ---------- audio ----------
export type Voice = "ciso" | "chair" | "member";
export interface AudioLine {
  id: Id;
  voice: Voice;
  text: string;
  ssml?: string;
  audioKey?: string;
  durationMs?: Ms;
  marks?: Record<string, Ms>;
}

export interface SetupSpec {
  instruction: string;
  audioCheck: AudioCheck;
  exampleScene: { template?: LayoutTemplate; layout: Block[]; lines: AudioLine[] };
  preStartText: { both: string[]; checker: string[] };
}
export interface AudioCheck {
  clip: AudioLine;
  prompt: string;
  options: Choice[];
  correctOptionId: Id;
}

// ---------- slides ----------
export type LayoutTemplate = "auto" | "title-bullets" | "numbers" | "image-right" | "chart";
export interface TitleBlock {
  kind: "title";
  id: Id;
  text: RichText;
}
export interface BulletBlock {
  kind: "bullet";
  id: Id;
  text: RichText;
}
export interface NumberBlock {
  kind: "number";
  id: Id;
  value: RichText;
  caption?: RichText;
}
export interface ImageBlock {
  kind: "image";
  id: Id;
  assetKey: string;
  alt: string;
}
export interface ChartBlock {
  kind: "chart";
  id: Id;
  chartType: "bar" | "line" | "timeline";
  data: ChartData;
  alt: string;
  caption?: RichText;
}
export type Block = TitleBlock | BulletBlock | NumberBlock | ImageBlock | ChartBlock;
export interface ChartSeries {
  id: Id;
  label: string;
  value?: number;
  start?: number;
  end?: number;
  points?: number[];
}
export interface ChartData {
  axis: { min: number; max: number; label?: string };
  series: ChartSeries[];
  markers?: { id: Id; at: number; label: string }[];
}
/** Logical slide canvas, scaled to fit. */
export const SLIDE_WIDTH = 1600;
export const SLIDE_HEIGHT = 900;

export type MarkingShape = "box" | "underline" | "circle";
export interface Marking {
  id: Id;
  targetIds: Id[];
  bbox?: [x: number, y: number, w: number, h: number];
  appearAtMark?: string;
  appearAtMs?: Ms;
  anchorText?: string;
  color?: Hex;
  shape: MarkingShape;
  label: string;
  detail: string;
  vlmRun?: string;
}

export type CheckerPanel =
  | { status: "noIssues"; appearAtMark?: string; appearAtMs?: Ms; text?: string }
  | { status: "flagged"; markings: Marking[] };

export type Answers =
  | { mode: "whole"; reason: AudioLine; denial: AudioLine; padDenialToReason?: boolean }
  | { mode: "joined"; opening: AudioLine; reasonBody: AudioLine; closing: AudioLine; padDenialToReason?: boolean };

// ---------- scenes ----------
export type Scene =
  | OpeningScene
  | SlideScene
  | RatingScene
  | ActionScene
  | QuestionScene
  | AnswerScene
  | BudgetScene
  | BackgroundScene
  | DebriefScene;

export interface OpeningScene {
  kind: "opening";
  id: Id;
  template?: LayoutTemplate;
  layout?: Block[];
  profileCard: ProfileCard;
  lines: AudioLine[];
}
export interface ProfileCard {
  name: string;
  role: string;
  lines: string[];
  portraitKey?: string;
  showFromLineId: Id;
}

export interface SlideScene {
  kind: "slide";
  id: Id;
  template?: LayoutTemplate;
  layout: Block[];
  narration: AudioLine;
  checker: CheckerPanel;
  routineQuestion?: AudioLine;
  participantQuestions?: { explain: AudioLine; challenge: AudioLine };
  answers?: Answers;
}

export type RatingPoint = string; // ^T(0|[1-9][0-9]*[ab]?)$
export interface RatingScene {
  kind: "rating";
  id: Id;
  point: RatingPoint;
  after: Id;
}
export interface ActionScene {
  kind: "action";
  id: Id;
  forSlide: Id;
}
export interface QuestionScene {
  kind: "question";
  id: Id;
  forSlide: Id;
}
export interface AnswerScene {
  kind: "answer";
  id: Id;
  forSlide: Id;
}

export interface BudgetScene {
  kind: "budget";
  id: Id;
  prompt: string;
  options: Choice[];
  confidence: { label: string; scale: ScaleSpec };
}

export interface BackgroundScene {
  kind: "background";
  id: Id;
  intro: string;
  preferNotToSayLabel: string;
  items: BackgroundItem[];
}
export type BackgroundItem =
  | { kind: "single"; id: Id; label: string; options: Choice[]; otherOptionId?: Id }
  | { kind: "scale"; id: Id; label: string; scale: ScaleSpec }
  | { kind: "country"; id: Id; label: string }
  | { kind: "honeypot"; id: Id; label: string };

export interface DebriefScene {
  kind: "debrief";
  id: Id;
  paragraphs: { both: string[]; checker: string[]; neutral: string[]; end: string[] };
  doNotDiscuss: string;
  sessionCodeLabel: string;
}

// ---------- measures ----------
export type RatingItemId = "trust" | "competence" | "benevolence" | "integrity";
export interface RatingSpec {
  instruction: string;
  items: RatingItem[];
}
export interface RatingItem {
  id: RatingItemId;
  label?: string;
  scale: ScaleSpec;
}
export interface ScaleSpec {
  min: number;
  max: number;
  minLabel: string;
  maxLabel: string;
  pointLabels?: Record<string, string>;
}

export interface ActionSpec {
  checker: { prompt: string; note?: string; options: { id: "nothing" | "explain" | "challenge"; label: string }[] };
  neutral: { prompt: string; note?: string; options: { id: "hasQuestion" | "noQuestion"; label: string }[] };
}
