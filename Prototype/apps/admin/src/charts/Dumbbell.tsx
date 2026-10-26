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
import { fmtCi, fmtNum } from "./format";
import { clip } from "./BarChart";
import { linear, niceDomain, niceTicks, textWidth } from "./scale";

export interface DumbbellSeries {
  key: string;
  label: string;
  color: string;
  /** Hollow marker (e.g. the denial point), so the two series differ by shape as well. */
  hollow?: boolean;
}
export interface CiPoint {
  mean: number | null;
  lo: number | null;
  hi: number | null;
  n: number;
}
export interface DumbbellRow {
  id: string;
  label: string;
  points: Record<string, CiPoint | undefined>;
  /** Colour for every point of this row (the row is a group), overriding the series colour. */
  color?: string;
}
export interface DumbbellProps {
  title: string;
  series: DumbbellSeries[];
  rows: DumbbellRow[];
  /** Draw the zero line (deltas); off for levels. */
  zeroLine?: boolean;
  domain?: [number, number];
  tableView?: boolean;
}

/** Value domain covering every CI (and zero when the zero line is drawn), rounded outward. */
export function dumbbellDomain(rows: DumbbellRow[], withZero: boolean): [number, number] {
  const vals: number[] = withZero ? [0] : [];
  for (const r of rows)
    for (const p of Object.values(r.points)) {
      if (!p) continue;
      for (const v of [p.mean, p.lo, p.hi]) if (v !== null && Number.isFinite(v)) vals.push(v);
    }
  if (vals.length === 0) return [-1, 1];
  return niceDomain(Math.min(...vals), Math.max(...vals), 4);
}

export function dumbbellTable(p: Pick<DumbbellProps, "title" | "series" | "rows">): TableData {
  const columns: TableColumn[] = [{ key: "label", label: "Row" }];
  for (const s of p.series)
    columns.push(
      { key: `${s.key}.mean`, label: `${s.label} mean`, numeric: true },
      { key: `${s.key}.ci`, label: `${s.label} 95% CI`, numeric: true },
      { key: `${s.key}.n`, label: `${s.label} n`, numeric: true },
    );
  return {
    caption: p.title,
    columns,
    rows: p.rows.map((r) => {
      const cells: Record<string, string | number> = { label: r.label };
      for (const s of p.series) {
        const pt = r.points[s.key];
        cells[`${s.key}.mean`] = fmtNum(pt?.mean);
        cells[`${s.key}.ci`] = pt && pt.lo !== null && pt.hi !== null ? fmtCi([pt.lo, pt.hi]) : "-";
        cells[`${s.key}.n`] = pt?.n ?? 0;
      }
      return cells;
    }),
  };
}

const ROW = 40;

export function Dumbbell(props: DumbbellProps) {
  const { title, series, rows, zeroLine = true, tableView } = props;
  const { ref, width } = useWidth<HTMLDivElement>(360);
  const [hot, setHot] = useState<{ row: number; key: string } | null>(null);

  if (tableView) return <TableTwin {...dumbbellTable(props)} />;

  const domain = props.domain ?? dumbbellDomain(rows, zeroLine);
  const labelW = Math.min(Math.max(...rows.map((r) => textWidth(r.label)), 30) + 16, width * 0.35);
  const plotW = Math.max(80, width - labelW - 14);
  const x = linear(domain, [labelW, labelW + plotW]);
  const ticks = niceTicks(domain[0], domain[1], 4);
  const h = rows.length * ROW + 24;
  const offset = (si: number) => (series.length < 2 ? 0 : (si - (series.length - 1) / 2) * 10);
  const colorOf = (r: DumbbellRow, s: DumbbellSeries) => r.color ?? s.color;

  let tip: TipState | null = null;
  if (hot) {
    const r = rows[hot.row];
    const s = series.find((q) => q.key === hot.key)!;
    const pt = r.points[s.key];
    if (pt)
      tip = {
        x: x(pt.mean ?? 0),
        y: hot.row * ROW + ROW,
        title: `${r.label}, ${s.label}`,
        rows: [
          {
            color: colorOf(r, s),
            key: "dot",
            value: fmtNum(pt.mean),
            label: "mean",
            sub: `CI ${fmtCi(pt.lo !== null && pt.hi !== null ? [pt.lo, pt.hi] : null)}, n = ${pt.n}`,
          },
        ],
      };
  }

  return (
    <figure className="viz" ref={ref}>
      {series.length >= 2 && (
        <Legend
          items={series.map((s) => ({ id: s.key, label: s.label, color: s.color, kind: "dot", hollow: s.hollow }))}
        />
      )}
      <svg width={width} height={h} role="img" aria-label={title}>
        {ticks.map((t) => (
          <g key={t}>
            <line className="viz-grid" x1={x(t)} x2={x(t)} y1={0} y2={rows.length * ROW} />
            <text className="viz-tick" x={x(t)} y={rows.length * ROW + 16} textAnchor="middle">
              {fmtNum(t)}
            </text>
          </g>
        ))}
        {zeroLine && domain[0] <= 0 && domain[1] >= 0 && (
          <line className="viz-zero" x1={x(0)} x2={x(0)} y1={0} y2={rows.length * ROW} />
        )}
        {rows.map((r, ri) => {
          const cy = ri * ROW + ROW / 2;
          const means = series
            .map((s) => r.points[s.key]?.mean)
            .filter((m): m is number => m !== null && m !== undefined);
          return (
            <g key={r.id}>
              <text className="viz-label" x={labelW - 10} y={cy + 4} textAnchor="end">
                {clip(r.label, labelW - 12)}
              </text>
              {means.length >= 2 && (
                <line
                  x1={x(Math.min(...means))}
                  x2={x(Math.max(...means))}
                  y1={cy}
                  y2={cy}
                  stroke="var(--viz-axis)"
                  strokeWidth={2}
                />
              )}
              {series.map((s, si) => {
                const pt = r.points[s.key];
                if (!pt || pt.mean === null) return null;
                const py = cy + offset(si);
                const c = colorOf(r, s);
                const isHot = hot?.row === ri && hot.key === s.key;
                return (
                  <g key={s.key} opacity={!hot || isHot ? 1 : 0.5}>
                    {pt.lo !== null && pt.hi !== null && (
                      <g stroke={c} strokeWidth={1.5}>
                        <line x1={x(pt.lo)} x2={x(pt.hi)} y1={py} y2={py} />
                        <line x1={x(pt.lo)} x2={x(pt.lo)} y1={py - 3} y2={py + 3} />
                        <line x1={x(pt.hi)} x2={x(pt.hi)} y1={py - 3} y2={py + 3} />
                      </g>
                    )}
                    <circle
                      cx={x(pt.mean)}
                      cy={py}
                      r={s.hollow ? 4 : 5}
                      fill={s.hollow ? "var(--viz-surface)" : c}
                      stroke={s.hollow ? c : "var(--viz-ring)"}
                      strokeWidth={2}
                    />
                    <rect
                      className="viz-hit"
                      x={x(pt.mean) - 12}
                      y={py - 12}
                      width={24}
                      height={24}
                      tabIndex={0}
                      aria-label={`${r.label}, ${s.label}: mean ${fmtNum(pt.mean)}`}
                      {...hoverProps(
                        () => setHot({ row: ri, key: s.key }),
                        () => setHot(null),
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
