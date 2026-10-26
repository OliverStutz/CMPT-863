import { useState, type KeyboardEvent, type PointerEvent } from "react";
import {
  Legend,
  TableTwin,
  Tooltip,
  localPoint,
  useWidth,
  type TableColumn,
  type TableData,
  type TipState,
} from "./common";
import { fmtCi, fmtNum } from "./format";
import { linear, niceTicks, textWidth } from "./scale";

export interface XPoint {
  id: string;
  label: string;
  /** Second tick line, e.g. "slide 2". */
  sub?: string;
  /** Marked on the axis (a flag point). */
  marked?: boolean;
}
export interface LinePoint {
  y: number | null;
  lo?: number | null;
  hi?: number | null;
  n?: number;
}
export interface LineSeries {
  id: string;
  label: string;
  color: string;
  /** By x id. */
  points: Record<string, LinePoint | undefined>;
}
export interface Trajectory {
  id: string;
  values: Record<string, number | null | undefined>;
  /** X ids drawn as a hollow marker on this trajectory (the participant's denial point). */
  hollow?: string[];
}

export interface LineChartProps {
  title: string;
  xs: XPoint[];
  series: LineSeries[];
  trajectories?: Trajectory[];
  yDomain: [number, number];
  yLabel?: string;
  height?: number;
  tableView?: boolean;
  /** Text under the plot, e.g. what the axis marker means. */
  note?: string;
}

export interface OrderedPoint extends LinePoint {
  index: number;
  x: string;
}

/** The series' points in the order of `xs` (plan order), with undefined points as y = null. */
export function orderedPoints(series: Pick<LineSeries, "points">, xs: Pick<XPoint, "id">[]): OrderedPoint[] {
  return xs.map((x, index) => {
    const p = series.points[x.id];
    return { index, x: x.id, y: p?.y ?? null, lo: p?.lo ?? null, hi: p?.hi ?? null, n: p?.n };
  });
}

/** Runs of consecutive points with a value; a missing point breaks the line instead of bridging it. */
export function segments<T extends { y: number | null }>(points: T[]): T[][] {
  const out: T[][] = [];
  let cur: T[] = [];
  for (const p of points) {
    if (p.y === null || !Number.isFinite(p.y)) {
      if (cur.length) out.push(cur);
      cur = [];
    } else cur.push(p);
  }
  if (cur.length) out.push(cur);
  return out;
}

/**
 * The 95% CI band as closed polygons in data space ([index, value] pairs): along the upper bounds left to right,
 * then back along the lower bounds right to left. A point without both bounds breaks the band.
 */
export function bandPolygons(points: OrderedPoint[]): [number, number][][] {
  const withCi = points.map((p) => ({ ...p, y: p.lo != null && p.hi != null ? p.y : null }));
  return segments(withCi).map((seg) => [
    ...seg.map((p) => [p.index, p.hi as number] as [number, number]),
    ...seg
      .slice()
      .reverse()
      .map((p) => [p.index, p.lo as number] as [number, number]),
  ]);
}

export function lineTable(p: Pick<LineChartProps, "title" | "xs" | "series">): TableData {
  const columns: TableColumn[] = [{ key: "point", label: "Point" }];
  for (const s of p.series) {
    columns.push(
      { key: `${s.id}.mean`, label: `${s.label} mean`, numeric: true },
      { key: `${s.id}.ci`, label: `${s.label} 95% CI`, numeric: true },
      { key: `${s.id}.n`, label: `${s.label} n`, numeric: true },
    );
  }
  const rows = p.xs.map((x) => {
    const r: Record<string, string | number | null> = { point: x.sub ? `${x.label} (${x.sub})` : x.label };
    for (const s of p.series) {
      const pt = s.points[x.id];
      r[`${s.id}.mean`] = fmtNum(pt?.y);
      r[`${s.id}.ci`] = pt?.lo != null && pt?.hi != null ? fmtCi([pt.lo, pt.hi]) : "-";
      r[`${s.id}.n`] = pt?.n ?? null;
    }
    return r;
  });
  return { caption: p.title, columns, rows };
}

export function trajectoryTable(title: string, xs: XPoint[], trajectories: Trajectory[]): TableData {
  return {
    caption: `${title}: each participant`,
    columns: [{ key: "id", label: "Participant" }, ...xs.map((x) => ({ key: x.id, label: x.label, numeric: true }))],
    rows: trajectories.map((t) => {
      const r: Record<string, string> = { id: t.id };
      for (const x of xs) r[x.id] = fmtNum(t.values[x.id] ?? null) + (t.hollow?.includes(x.id) ? " (denial)" : "");
      return r;
    }),
  };
}

const M = { top: 10, right: 112, bottom: 44, left: 34 };

export function LineChart(props: LineChartProps) {
  const { title, xs, series, trajectories = [], yDomain, height = 260, tableView, note, yLabel } = props;
  const { ref, width } = useWidth<HTMLDivElement>();
  const [idx, setIdx] = useState<number | null>(null);

  if (tableView) {
    return (
      <div>
        <TableTwin {...lineTable(props)} />
        {trajectories.length > 0 && <TableTwin {...trajectoryTable(title, xs, trajectories)} />}
      </div>
    );
  }

  const plotW = Math.max(120, width - M.left - M.right);
  const plotH = height - M.top - M.bottom;
  const x = (i: number) => M.left + (xs.length > 1 ? (i * plotW) / (xs.length - 1) : plotW / 2);
  const y = linear(yDomain, [M.top + plotH, M.top]);
  const span = yDomain[1] - yDomain[0];
  const ticks = niceTicks(yDomain[0], yDomain[1], Number.isInteger(span) && span <= 6 ? span : 5);
  const step = xs.length > 1 ? plotW / (xs.length - 1) : plotW;
  const ordered = series.map((s) => ({ s, pts: orderedPoints(s, xs) }));

  const linePath = (pts: { index: number; y: number | null }[]) =>
    segments(pts)
      .map((seg) => seg.map((p, i) => `${i === 0 ? "M" : "L"}${x(p.index)},${y(p.y as number)}`).join(""))
      .join("");

  // Direct end labels unless two of them would collide (then the legend and tooltip carry identity).
  const ends = ordered
    .map(({ s, pts }) => {
      const last = [...pts].reverse().find((p) => p.y !== null);
      return last ? { s, p: last, ty: y(last.y as number) } : null;
    })
    .filter((e): e is NonNullable<typeof e> => e !== null);
  const sortedEnds = [...ends].sort((a, b) => a.ty - b.ty);
  const collide = sortedEnds.some((e, i) => i > 0 && e.ty - sortedEnds[i - 1].ty < 14);

  const tip: TipState | null =
    idx === null
      ? null
      : {
          x: x(idx),
          y: M.top,
          title: xs[idx].sub ? `${xs[idx].label}, ${xs[idx].sub}` : xs[idx].label,
          rows: ordered.map(({ s, pts }) => {
            const p = pts[idx];
            return {
              color: s.color,
              key: "line" as const,
              value: fmtNum(p.y),
              label: s.label,
              sub: `${p.lo != null && p.hi != null ? `CI ${fmtCi([p.lo, p.hi])}, ` : ""}n = ${p.n ?? 0}`,
            };
          }),
        };

  function onMove(e: PointerEvent<SVGRectElement>) {
    // Only a real move shows the crosshair; a layout shift under a resting pointer must not (see hoverProps).
    if (!e.isTrusted || (e.movementX === 0 && e.movementY === 0)) return;
    const { x: px } = localPoint(e, ref.current);
    const i = Math.round((px - M.left) / step);
    setIdx(Math.max(0, Math.min(xs.length - 1, i)));
  }
  function onKey(e: KeyboardEvent<SVGGElement>) {
    if (e.key === "ArrowRight") setIdx((i) => Math.min(xs.length - 1, (i ?? -1) + 1));
    else if (e.key === "ArrowLeft") setIdx((i) => Math.max(0, (i ?? 1) - 1));
    else if (e.key === "Escape") setIdx(null);
    else return;
    e.preventDefault();
  }

  return (
    <figure className="viz" ref={ref}>
      {series.length >= 2 && (
        <Legend items={series.map((s) => ({ id: s.id, label: s.label, color: s.color, kind: "line" }))} />
      )}
      <svg width={width} height={height} role="img" aria-label={title}>
        {yLabel && (
          <text className="viz-tick" x={0} y={M.top - 2} style={{ fontSize: 11 }}>
            {yLabel}
          </text>
        )}
        {ticks.map((t) => (
          <g key={t}>
            <line className="viz-grid" x1={M.left} x2={M.left + plotW} y1={y(t)} y2={y(t)} />
            <text className="viz-tick" x={M.left - 6} y={y(t) + 4} textAnchor="end">
              {fmtNum(t)}
            </text>
          </g>
        ))}
        <line className="viz-axis" x1={M.left} x2={M.left + plotW} y1={M.top + plotH} y2={M.top + plotH} />
        {xs.map((p, i) => (
          <g key={p.id}>
            <text className="viz-tick" x={x(i)} y={M.top + plotH + 16} textAnchor="middle">
              {p.label}
            </text>
            {p.sub && step >= 48 && (
              <text className="viz-tick" x={x(i)} y={M.top + plotH + 30} textAnchor="middle" style={{ fontSize: 10 }}>
                {p.sub}
              </text>
            )}
            {p.marked && (
              <path
                d={`M${x(i) - 5},${M.top + plotH - 7}L${x(i) + 5},${M.top + plotH - 7}L${x(i)},${M.top + plotH}Z`}
                fill="var(--viz-ink-2)"
              >
                <title>flag point</title>
              </path>
            )}
          </g>
        ))}
        <g aria-hidden>
          {trajectories.map((t) => {
            const pts = xs.map((p, index) => ({ index, y: t.values[p.id] ?? null }));
            return (
              <g key={t.id}>
                <path
                  d={linePath(pts)}
                  fill="none"
                  stroke="var(--viz-deemph)"
                  strokeWidth={1}
                  strokeOpacity={0.22}
                  strokeLinejoin="round"
                />
                {(t.hollow ?? []).map((xid) => {
                  const i = xs.findIndex((p) => p.id === xid);
                  const v = t.values[xid];
                  if (i < 0 || v === null || v === undefined) return null;
                  return (
                    <circle
                      key={xid}
                      cx={x(i)}
                      cy={y(v)}
                      r={3.5}
                      fill="var(--viz-surface)"
                      stroke="var(--viz-deemph)"
                      strokeOpacity={0.7}
                      strokeWidth={1}
                    />
                  );
                })}
              </g>
            );
          })}
        </g>
        {ordered.map(({ s, pts }) =>
          bandPolygons(pts).map((poly, k) => (
            <path
              key={`${s.id}-band-${k}`}
              d={poly.map(([i, v], j) => `${j === 0 ? "M" : "L"}${x(i)},${y(v)}`).join("") + "Z"}
              fill={s.color}
              fillOpacity={0.12}
            />
          )),
        )}
        {ordered.map(({ s, pts }) => (
          <path
            key={s.id}
            d={linePath(pts)}
            fill="none"
            stroke={s.color}
            strokeWidth={2}
            strokeLinejoin="round"
            strokeLinecap="round"
          />
        ))}
        {idx !== null && <line className="viz-crosshair" x1={x(idx)} x2={x(idx)} y1={M.top} y2={M.top + plotH} />}
        {ordered.map(({ s, pts }) => {
          const p = idx !== null ? pts[idx] : null;
          return p && p.y !== null ? (
            <circle key={s.id} cx={x(idx!)} cy={y(p.y)} r={4} fill={s.color} stroke="var(--viz-ring)" strokeWidth={2} />
          ) : null;
        })}
        {!collide &&
          ends.map(({ s, p, ty }) => (
            <g key={s.id}>
              <circle cx={x(p.index)} cy={ty} r={4} fill={s.color} stroke="var(--viz-ring)" strokeWidth={2} />
              <text className="viz-label" x={x(p.index) + 9} y={ty + 4}>
                {truncate(`${s.label} ${fmtNum(p.y)}`, M.right - 12)}
              </text>
            </g>
          ))}
        <g
          className="viz-plot"
          tabIndex={0}
          role="group"
          aria-label={`${title}: use the arrow keys to read each point`}
          onFocus={() => setIdx((i) => i ?? 0)}
          onBlur={() => setIdx(null)}
          onKeyDown={onKey}
        >
          <rect
            className="viz-hit"
            x={M.left - step / 2}
            y={M.top}
            width={plotW + step}
            height={plotH}
            onPointerMove={onMove}
            onPointerLeave={() => setIdx(null)}
          />
        </g>
      </svg>
      {note && <figcaption className="muted small-text">{note}</figcaption>}
      <Tooltip tip={tip} width={width} />
    </figure>
  );
}

function truncate(s: string, px: number): string {
  if (textWidth(s) <= px) return s;
  let t = s;
  while (t.length > 3 && textWidth(t + "...") > px) t = t.slice(0, -1);
  return t + "...";
}
