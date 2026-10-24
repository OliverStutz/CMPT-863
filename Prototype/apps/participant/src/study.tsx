import { createContext, useContext } from "react";
import type { Progress, PublicScene, PublicScenePack, Step } from "@study/shared";
import type { ApiClient } from "./api/client";

export interface StudyContextValue {
  api: ApiClient;
  pack: PublicScenePack;
  assetBaseUrl: string;
  code: string;
  progress?: Progress;
  /** GET /next and render whatever comes back. */
  refresh: () => Promise<void>;
  /** Shows a recoverable error banner; the router reloads the step from the server on "Try again". */
  reportError: (e: unknown) => void;
  /** Device gate failed: show the blocked page under /s/:code/blocked. */
  block: () => void;
  /** The last debrief shown in this page life, so the "done" step can keep it on screen. */
  lastDebrief?: Step["debrief"];
}

export const StudyContext = createContext<StudyContextValue | null>(null);

export function useStudy(): StudyContextValue {
  const v = useContext(StudyContext);
  if (!v) throw new Error("useStudy outside StudyContext");
  return v;
}

export function findScene(pack: PublicScenePack, id: string | undefined): PublicScene | undefined {
  if (!id) return undefined;
  return pack.scenes.find((s) => s.id === id);
}

/** Joins the server's asset base URL and an asset key. */
export function assetUrl(base: string, key: string): string {
  if (/^(https?:)?\/\//.test(key) || key.startsWith("/")) return key;
  return `${base.replace(/\/+$/, "")}/${key.replace(/^\/+/, "")}`;
}
