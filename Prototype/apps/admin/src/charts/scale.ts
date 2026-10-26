// Small scale helpers for the hand-written SVG charts (no chart library).

export type Scale = ((v: number) => number) & { domain: [number, number]; range: [number, number] };

export function linear(domain: [number, number], range: [number, number]): Scale {
  const [d0, d1] = domain;
  const [r0, r1] = range;
  const span = d1 - d0 || 1;
  const f = ((v: number) => r0 + ((v - d0) / span) * (r1 - r0)) as Scale;
  f.domain = domain;
  f.range = range;
  return f;
}

/** About `count` round tick values covering [min, max] (steps of 1, 2 or 5 times a power of ten). */
export function niceTicks(min: number, max: number, count = 5): number[] {
  if (!Number.isFinite(min) || !Number.isFinite(max)) return [];
  if (min === max) return [min];
  const raw = (max - min) / Math.max(1, count);
  const pow = Math.pow(10, Math.floor(Math.log10(raw)));
  const step = [1, 2, 5, 10].map((m) => m * pow).find((s) => s >= raw) ?? 10 * pow;
  const start = Math.ceil(min / step - 1e-9) * step;
  const out: number[] = [];
  for (let v = start; v <= max + step * 1e-9; v += step) out.push(Math.round(v / step) * step);
  return out.map((v) => (Math.abs(v) < 1e-12 ? 0 : Number(v.toFixed(10))));
}

/** Extends [min, max] outward to round numbers, so the axis ends on a tick. */
export function niceDomain(min: number, max: number, count = 5): [number, number] {
  if (min === max) return [min - 1, max + 1];
  const raw = (max - min) / Math.max(1, count);
  const pow = Math.pow(10, Math.floor(Math.log10(raw)));
  const step = [1, 2, 5, 10].map((m) => m * pow).find((s) => s >= raw) ?? 10 * pow;
  return [Math.floor(min / step) * step, Math.ceil(max / step) * step];
}

/** Rough rendered width of a label in px, for fit checks (system sans, about 0.6 em per character). */
export function textWidth(text: string, fontSize = 12): number {
  return text.length * fontSize * 0.6;
}

/**
 * Path of a bar with a 4 px rounded data end and a square end at the baseline.
 * Horizontal bars grow to the right from x; vertical bars grow upward from the baseline y (y + length = baseline).
 */
export function barPath(
  orientation: "horizontal" | "vertical",
  x: number,
  y: number,
  length: number,
  thickness: number,
  radius = 4,
): string {
  const len = Math.max(0, length);
  const r = Math.max(0, Math.min(radius, len, thickness / 2));
  if (orientation === "horizontal") {
    const x1 = x + len;
    return [
      `M${x},${y}`,
      `H${x1 - r}`,
      `Q${x1},${y} ${x1},${y + r}`,
      `V${y + thickness - r}`,
      `Q${x1},${y + thickness} ${x1 - r},${y + thickness}`,
      `H${x}`,
      "Z",
    ].join("");
  }
  const base = y + len;
  return [
    `M${x},${base}`,
    `V${y + r}`,
    `Q${x},${y} ${x + r},${y}`,
    `H${x + thickness - r}`,
    `Q${x + thickness},${y} ${x + thickness},${y + r}`,
    `V${base}`,
    "Z",
  ].join("");
}

export const clamp = (v: number, lo: number, hi: number) => Math.min(hi, Math.max(lo, v));
