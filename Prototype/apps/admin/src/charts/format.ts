// Number formatting shared by the charts, the tables and the hypothesis cards.

const twoDecimals = new Intl.NumberFormat("en", { maximumFractionDigits: 2, minimumFractionDigits: 0 });
const intFmt = new Intl.NumberFormat("en", { maximumFractionDigits: 0 });

/** At most two decimals; "-" for null or a non-finite value. A minus sign is a plain hyphen. */
export function fmtNum(x: number | null | undefined, digits = 2): string {
  if (x === null || x === undefined || !Number.isFinite(x)) return "-";
  const f = digits === 2 ? twoDecimals : new Intl.NumberFormat("en", { maximumFractionDigits: digits });
  const s = f.format(x);
  return s === "-0" ? "0" : s;
}

/** Whole numbers with thousands separators. */
export function fmtInt(x: number | null | undefined): string {
  if (x === null || x === undefined || !Number.isFinite(x)) return "-";
  return intFmt.format(x);
}

// DECISION: p keeps three decimals without the leading zero (".032", APA style), because two decimals would
// print .004 as ".00"; below .001 it reads "< .001". p of 1 prints "1".
export function fmtP(p: number | null | undefined): string {
  if (p === null || p === undefined || !Number.isFinite(p)) return "-";
  if (p < 0.001) return "< .001";
  if (p >= 1) return "1";
  return p.toFixed(3).replace(/^0/, "");
}

/** "[lo, hi]" with at most two decimals each. */
export function fmtCi(ci: [number, number] | null | undefined): string {
  if (!ci) return "-";
  return `[${fmtNum(ci[0])}, ${fmtNum(ci[1])}]`;
}

/** A share 0..1 as a whole percentage, "45%". */
export function fmtPct(share: number | null | undefined): string {
  if (share === null || share === undefined || !Number.isFinite(share)) return "-";
  return `${Math.round(share * 100)}%`;
}

/** Seconds as m:ss, or h:mm:ss from an hour. */
export function fmtDuration(sec: number | null | undefined): string {
  if (sec === null || sec === undefined || !Number.isFinite(sec)) return "-";
  const total = Math.round(sec);
  const h = Math.floor(total / 3600);
  const m = Math.floor((total % 3600) / 60);
  const s = String(total % 60).padStart(2, "0");
  return h > 0 ? `${h}:${String(m).padStart(2, "0")}:${s}` : `${m}:${s}`;
}

/** Signed number, "+1.5" or "-0.33"; zero has no sign. */
export function fmtSigned(x: number | null | undefined): string {
  if (x === null || x === undefined || !Number.isFinite(x)) return "-";
  const s = fmtNum(x);
  return x > 0 && s !== "0" ? `+${s}` : s;
}
