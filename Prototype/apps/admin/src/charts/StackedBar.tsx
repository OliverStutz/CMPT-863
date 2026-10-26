import { useState } from "react";
import { Legend, TableTwin, Tooltip, useWidth, type TableData, type TipState, hoverProps } from "./common";
import { fmtInt, fmtPct } from "./format";
import { barPath, textWidth } from "./scale";
import { clip } from "./BarChart";

export interface StackKey {
  key: string;
  label: string;
  color: string;
  /** Label colour inside the fill (a CSS colour or token): white by default, near-black on a light fill. */
  ink?: string;
}
export interface StackRow {
  id: string;
  label: string;
  values: Record<string, number | undefined>;
}
export interface StackedBarProps {
  title: string;
  keys: StackKey[];
  rows: StackRow[];
  /** 100% bars (default) or bars on a common count axis. */
  normalize?: boolean;
  tableView?: boolean;
}

export interface Segment {
  key: string;
  value: number;
  share: number;
  x0: number;
  x1: number;
}

const GAP = 2;

/** Segment geometry of one row: widths proportional to the values, a 2 px surface gap between neighbours. */
export function stackSegments(row: StackRow, keys: StackKey[], width: number, scaleTotal?: number): Segment[] {
  const total = keys.reduce((a, k) => a + (row.values[k.key] ?? 0), 0);
  const denom = scaleTotal ?? total;
  const out: Segment[] = [];
  let x = 0;
  for (const k of keys) {
    const v = row.values[k.key] ?? 0;
    if (v <= 0 || denom <= 0) continue;
    const w = (v / denom) * width;
    const start = out.length === 0 ? x : x + GAP;
    out.push({ key: k.key, value: v, share: total > 0 ? v / total : 0, x0: start, x1: Math.max(start, x + w) });
    x += w;
  }
  return out;
}

export function stackTable(p: Pick<StackedBarProps, "title" | "keys" | "rows">): TableData {
  return {
    caption: p.title,
    columns: [
      { key: "label", label: "Row" },
      ...p.keys.map((k) => ({ key: k.key, label: k.label, numeric: true })),
      { key: "total", label: "Total", numeric: true },
    ],
    rows: p.rows.map((r) => {
      const total = p.keys.reduce((a, k) => a + (r.values[k.key] ?? 0), 0);
      const cells: Record<string, string> = { label: r.label, total: fmtInt(total) };
      for (const k of p.keys) {
        const v = r.values[k.key] ?? 0;
        cells[k.key] = `${fmtInt(v)} (${fmtPct(total ? v / total : null)})`;
      }
      return cells;
    }),
  };
}

const ROW = 30;
const THICK = 20;

export function StackedBar(props: StackedBarProps) {
  const { title, keys, rows, normalize = true, tableView } = props;
  const { ref, width } = useWidth<HTMLDivElement>();
  const [hot, setHot] = useState<{ row: number; key: string } | null>(null);

  if (tableView) return <TableTwin {...stackTable(props)} />;

  const labelW = Math.min(Math.max(...rows.map((r) => textWidth(r.label)), 40) + 14, width * 0.35);
  const plotW = Math.max(60, width - labelW - 8);
  const maxTotal = Math.max(1, ...rows.map((r) => keys.reduce((a, k) => a + (r.values[k.key] ?? 0), 0)));
  const keyOf = (k: string) => keys.find((x) => x.key === k)!;
  const h = rows.length * ROW + 4;

  let tip: TipState | null = null;
  if (hot) {
    const r = rows[hot.row];
    const segs = stackSegments(r, keys, plotW, normalize ? undefined : maxTotal);
    const s = segs.find((x) => x.key === hot.key);
    if (s)
      tip = {
        x: labelW + (s.x0 + s.x1) / 2,
        y: hot.row * ROW + ROW,
        title: r.label,
        rows: [
          {
            color: keyOf(s.key).color,
            key: "rect",
            value: `${fmtInt(s.value)} (${fmtPct(s.share)})`,
            label: keyOf(s.key).label,
          },
        ],
      };
  }

  return (
    <figure className="viz" ref={ref}>
      <Legend items={keys.map((k) => ({ id: k.key, label: k.label, color: k.color, kind: "rect" }))} />
      <svg width={width} height={h} role="img" aria-label={title}>
        {rows.map((r, ri) => {
          const yTop = ri * ROW + (ROW - THICK) / 2;
          const segs = stackSegments(r, keys, plotW, normalize ? undefined : maxTotal);
          return (
            <g key={r.id}>
              <text className="viz-label" x={labelW - 8} y={yTop + THICK / 2 + 4} textAnchor="end">
                {clip(r.label, labelW - 10)}
              </text>
              {segs.map((s, si) => {
                const k = keyOf(s.key);
                const w = s.x1 - s.x0;
                const last = si === segs.length - 1;
                const text = fmtInt(s.value);
                const fits = w >= textWidth(text) + 10;
                const isHot = hot?.row === ri && hot.key === s.key;
                return (
                  <g key={s.key}>
                    {last ? (
                      <path
                        className="viz-mark"
                        d={barPath("horizontal", labelW + s.x0, yTop, w, THICK)}
                        fill={k.color}
                        opacity={!hot || isHot ? 1 : 0.55}
                      />
                    ) : (
                      <rect
                        className="viz-mark"
                        x={labelW + s.x0}
                        y={yTop}
                        width={w}
                        height={THICK}
                        fill={k.color}
                        opacity={!hot || isHot ? 1 : 0.55}
                      />
                    )}
                    {fits && (
                      <text
                        className="viz-inside"
                        x={labelW + s.x0 + w / 2}
                        y={yTop + THICK / 2 + 4}
                        textAnchor="middle"
                        style={k.ink ? { fill: k.ink } : undefined}
                      >
                        {text}
                      </text>
                    )}
                    <rect
                      className="viz-hit"
                      x={labelW + s.x0}
                      y={ri * ROW}
                      width={Math.max(w, 24)}
                      height={ROW}
                      tabIndex={0}
                      aria-label={`${r.label}, ${k.label}: ${text} (${fmtPct(s.share)})`}
                      {...hoverProps(
                        () => setHot({ row: ri, key: s.key }),
                        () => setHot(null),
                        s.value > 0,
                      )}
                    />
                  </g>
                );
              })}
            </g>
          );
        })}
      </svg>
      <Tooltip tip={tip} width={width} />
    </figure>
  );
}
