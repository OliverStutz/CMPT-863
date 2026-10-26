import type { SessionStatus } from "../api/types";

export type ChipTone = "neutral" | "info" | "active" | "success" | "warning" | "danger" | "muted";

export interface Chip {
  label: string;
  tone: ChipTone;
}

const STATUS_TONES: Record<SessionStatus, ChipTone> = {
  Opened: "neutral",
  Started: "info",
  Ongoing: "active",
  Finished: "success",
  Abandoned: "warning",
  Blocked: "danger",
};

/** Status chip for the participants table (section 7.1 status mapping). Unknown values get a neutral chip. */
export function statusChip(status: string): Chip {
  const tone = (STATUS_TONES as Record<string, ChipTone | undefined>)[status];
  return { label: status || "-", tone: tone ?? "neutral" };
}

// Flags named in the architecture (sections 5.3, 6, 10). Any other flag string from the server still shows,
// as a neutral chip with its raw name; the app holds no list of study content.
const FLAG_TONES: Record<string, ChipTone> = {
  possible_duplicate: "danger",
  audio_check_wrong: "warning",
  reloaded: "info",
  forced_allocation: "info",
};

export function flagChip(flag: string): Chip {
  return { label: flag, tone: FLAG_TONES[flag] ?? "warning" };
}
