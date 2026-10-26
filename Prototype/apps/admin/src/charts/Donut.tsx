import { useState } from "react";
import {
  Legend,
  TableTwin,
  Tooltip,
  useWidth,
  type TableColumn,
  type TableData,
  type TipState,
  hoverProps,
} from "./common";
import { fmtInt, fmtPct } from "./format";
import { textWidth } from "./scale";

export interface Slice {
  id: string;
  label: string;
  value: number;
  extra?: Record<string, string | number | null>;
}
export interface FoldedSlice extends Slice {
  /** The slices folded into "Other". */
  members?: Slice[];
}

export const OTHER_ID = "__other";

/**
 * At most `max` slices (largest first, ties by label) plus one "Other" holding the rest.
 * Zero slices are dropped; when nothing is left over there is no "Other".
 */
export function foldSlices(data: Slice[], max = 6, otherLabel = "Other"): FoldedSlice[] {
  const sorted = data.filter((d) => d.value > 0).sort((a, b) => b.value - a.value || a.label.localeCompare(b.label));
  if (sorted.length <= max) return sorted;
  const head = sorted.slice(0, max);
  const tail = sorted.slice(max);
  return [...head, { id: OTHER_ID, label: otherLabel, value: tail.reduce((a, s) => a + s.value, 0), members: tail }];
}

// DECISION: slices take categorical slots 1 to 6 in fixed order (validated adjacent, light and dark) and "Other" the
// grey; slots 1 and 2 are the group hues elsewhere, but a donut slice is never a group, and its legend names it.
const SLOT_COLORS = [
  "var(--viz-cat-1)",
  "var(--viz-cat-2)",
  "var(--viz-cat-3)",
  "var(--viz-cat-4)",
  "var(--viz-cat-5)",
  "var(--viz-cat-6)",
];

export function sliceColor(s: FoldedSlice, i: number): string {
  return s.id === OTHER_ID ? "var(--viz-other)" : SLOT_COLORS[i % SLOT_COLORS.length];
}

export interface DonutProps {
  title: string;
  data: Slice[];
  maxSlices?: number;
  extraColumns?: TableColumn[];
  tableView?: boolean;
}

export function donutTable(p: Pick<DonutProps, "title" | "data" | "extraColumns">): TableData {
  const total = p.data.reduce((a, s) => a + s.value, 0);
  return {
    caption: p.title,
    columns: [
      { key: "label", label: "Category" },
      { key: "n", label: "n", numeric: true },
      { key: "share", label: "Share", numeric: true },
      ...(p.extraColumns ?? []),
    ],
    rows: [...p.data]
      .sort((a, b) => b.value - a.value || a.label.localeCompare(b.label))
      .map((s) => ({
        label: s.label,
        n: fmtInt(s.value),
        share: fmtPct(total ? s.value / total : null),
        ...(s.extra ?? {}),
      })),
  };
}

function arcPath(cx: number, cy: number, r0: number, r1: number, a0: number, a1: number): string {
  const large = a1 - a0 > Math.PI ? 1 : 0;
  const p = (r: number, a: number) => `${cx + r * Math.sin(a)},${cy - r * Math.cos(a)}`;
  if (a1 - a0 >= Math.PI * 2 - 1e-6) {
    // A full ring: two half arcs.
    return `M${p(r1, 0)}A${r1},${r1} 0 1 1 ${p(r1, Math.PI)}A${r1},${r1} 0 1 1 ${p(r1, 0)}M${p(r0, 0)}A${r0},${r0} 0 1 0 ${p(r0, Math.PI)}A${r0},${r0} 0 1 0 ${p(r0, 0)}Z`;
  }
  return `M${p(r1, a0)}A${r1},${r1} 0 ${large} 1 ${p(r1, a1)}L${p(r0, a1)}A${r0},${r0} 0 ${large} 0 ${p(r0, a0)}Z`;
}

export function Donut(props: DonutProps) {
  const { title, data, maxSlices = 6, tableView } = props;
  const { ref, width } = useWidth<HTMLDivElement>();
  const [hot, setHot] = useState<number | null>(null);

  if (tableView) return <TableTwin {...donutTable(props)} />;

  const slices = foldSlices(data, maxSlices);
  const total = slices.reduce((a, s) => a + s.value, 0);
  const r1 = Math.max(50, Math.min(100, (width - 260) / 2));
  const r0 = r1 - 26;
  const h = r1 * 2 + 40;
  const cx = width / 2;
  const cy = h / 2;
  const pad = slices.length > 1 ? 2 / r1 : 0; // a 2 px surface gap at the outer edge

  let a = 0;
  const arcs = slices.map((s, i) => {
    const span = total ? (s.value / total) * Math.PI * 2 : 0;
    const a0 = a;
    a += span;
    return { s, i, a0, a1: a, mid: a0 + span / 2, share: total ? s.value / total : 0 };
  });

  // Direct labels outside the ring; skipped for thin slices and where they would collide (legend and table carry them).
  const placed: { side: number; y: number }[] = [];
  const labels = arcs.map((arc) => {
    if (arc.share < 0.05) return null;
    const lx = cx + (r1 + 12) * Math.sin(arc.mid);
    const ly = cy - (r1 + 12) * Math.cos(arc.mid);
    const side = Math.sin(arc.mid) >= 0 ? 1 : -1;
    const text = `${arc.s.label} ${fmtInt(arc.s.value)} (${fmtPct(arc.share)})`;
    if (placed.some((p) => p.side === side && Math.abs(p.y - ly) < 14)) return null;
    if (side > 0 ? lx + textWidth(text) > width : lx - textWidth(text) < 0) return null;
    placed.push({ side, y: ly });
    return { lx, ly, side, text };
  });

  let tip: TipState | null = null;
  if (hot !== null) {
    const arc = arcs[hot];
    tip = {
      x: cx + r1 * Math.sin(arc.mid) * 0.6,
      y: cy - r1 * Math.cos(arc.mid) * 0.6,
      title: arc.s.label,
      rows: [
        { color: sliceColor(arc.s, arc.i), key: "rect", value: `${fmtInt(arc.s.value)} (${fmtPct(arc.share)})` },
        ...(arc.s.members ?? []).map((m) => ({ value: fmtInt(m.value), label: m.label })),
      ],
    };
  }

  return (
    <figure className="viz" ref={ref}>
      <Legend
        items={arcs.map((arc) => ({
          id: arc.s.id,
          label: `${arc.s.label} ${fmtInt(arc.s.value)} (${fmtPct(arc.share)})`,
          color: sliceColor(arc.s, arc.i),
          kind: "rect",
        }))}
      />
      <svg width={width} height={h} role="img" aria-label={title}>
        {arcs.map((arc) => (
          <path
            key={arc.s.id}
            className="viz-mark viz-hit-arc"
            d={arcPath(
              cx,
              cy,
              r0,
              r1,
              arc.a0 + (arc.a1 - arc.a0 > pad * 2 ? pad / 2 : 0),
              arc.a1 - (arc.a1 - arc.a0 > pad * 2 ? pad / 2 : 0),
            )}
            fill={sliceColor(arc.s, arc.i)}
            opacity={hot === null || hot === arc.i ? 1 : 0.5}
            tabIndex={0}
            aria-label={`${arc.s.label}: ${fmtInt(arc.s.value)}, ${fmtPct(arc.share)}`}
            {...hoverProps(
              () => setHot(arc.i),
              () => setHot(null),
              arc.s.value > 0,
            )}
            style={{ outline: "none" }}
          />
        ))}
        <text className="viz-label" x={cx} y={cy + 4} textAnchor="middle" style={{ fontSize: 14, fontWeight: 600 }}>
          {fmtInt(total)}
        </text>
        {labels.map((l, i) =>
          l ? (
            <text key={i} className="viz-label" x={l.lx} y={l.ly + 4} textAnchor={l.side > 0 ? "start" : "end"}>
              {l.text}
            </text>
          ) : null,
        )}
      </svg>
      <Tooltip tip={tip} width={width} />
    </figure>
  );
}
