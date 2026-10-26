import { useState } from "react";
import { TableTwin, Tooltip, useWidth, type TableData, type TipState, hoverProps } from "./common";
import { fmtInt } from "./format";
import { linear, niceTicks } from "./scale";

export interface Bin {
  from: number;
  to: number;
  n: number;
}
export interface HistogramProps {
  title: string;
  bins: Bin[];
  /** Formats a bin edge for the axis, the tooltip and the table. */
  formatEdge?: (v: number) => string;
  height?: number;
  tableView?: boolean;
}

export function histogramTable(p: Pick<HistogramProps, "title" | "bins" | "formatEdge">): TableData {
  const f = p.formatEdge ?? String;
  return {
    caption: p.title,
    columns: [
      { key: "range", label: "From, to" },
      { key: "n", label: "n", numeric: true },
    ],
    rows: p.bins.map((b) => ({ range: `${f(b.from)} to ${f(b.to)}`, n: fmtInt(b.n) })),
  };
}

export function Histogram(props: HistogramProps) {
  const { title, bins, tableView, height = 140 } = props;
  const f = props.formatEdge ?? String;
  const { ref, width } = useWidth<HTMLDivElement>();
  const [hot, setHot] = useState<number | null>(null);
  if (tableView) return <TableTwin {...histogramTable(props)} />;
  if (bins.length === 0) return <p className="muted">No finished sessions.</p>;

  const left = 30;
  const top = 8;
  const bottom = 22;
  const plotW = Math.max(60, width - left - 8);
  const lo = bins[0].from;
  const hi = bins[bins.length - 1].to;
  const x = linear([lo, hi], [left, left + plotW]);
  const max = Math.max(1, ...bins.map((b) => b.n));
  const y = linear([0, max], [top + height, top]);
  const yTicks = niceTicks(0, max, 3).filter((t) => Number.isInteger(t));
  const xTicks = bins.length <= 12 ? bins.map((b) => b.from).concat(hi) : niceTicks(lo, hi, 6);

  const tip: TipState | null =
    hot === null
      ? null
      : {
          x: x((bins[hot].from + bins[hot].to) / 2),
          y: y(bins[hot].n),
          title: `${f(bins[hot].from)} to ${f(bins[hot].to)}`,
          rows: [{ value: fmtInt(bins[hot].n), label: "sessions" }],
        };

  return (
    <figure className="viz" ref={ref}>
      <svg width={width} height={top + height + bottom} role="img" aria-label={title}>
        {yTicks.map((t) => (
          <g key={t}>
            <line className="viz-grid" x1={left} x2={left + plotW} y1={y(t)} y2={y(t)} />
            <text className="viz-tick" x={left - 6} y={y(t) + 4} textAnchor="end">
              {t}
            </text>
          </g>
        ))}
        <line className="viz-axis" x1={left} x2={left + plotW} y1={top + height} y2={top + height} />
        {xTicks.map((t) => (
          <text key={t} className="viz-tick" x={x(t)} y={top + height + 15} textAnchor="middle">
            {f(t)}
          </text>
        ))}
        {bins.map((b, i) => {
          const x0 = x(b.from) + 1;
          const w = Math.max(1, x(b.to) - x(b.from) - 2); // 2 px surface gap between touching bins
          return (
            <g key={i}>
              {b.n > 0 && (
                <rect
                  className="viz-mark"
                  x={x0}
                  y={y(b.n)}
                  width={w}
                  height={top + height - y(b.n)}
                  rx={Math.min(2, w / 2)}
                  fill="var(--viz-single)"
                  opacity={hot === null || hot === i ? 1 : 0.55}
                />
              )}
              <rect
                className="viz-hit"
                x={x0 - 1 - Math.max(0, (24 - w - 2) / 2)}
                y={top}
                width={Math.max(24, w + 2)}
                height={height}
                tabIndex={0}
                aria-label={`${f(b.from)} to ${f(b.to)}: ${b.n}`}
                {...hoverProps(
                  () => setHot(i),
                  () => setHot(null),
                  b.n > 0,
                )}
              />
            </g>
          );
        })}
      </svg>
      <Tooltip tip={tip} width={width} />
    </figure>
  );
}
