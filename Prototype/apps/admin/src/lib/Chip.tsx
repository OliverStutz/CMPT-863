import type { Chip as ChipT } from "./status";

export function Chip({ chip }: { chip: ChipT }) {
  return <span className={`chip chip-${chip.tone}`}>{chip.label}</span>;
}
