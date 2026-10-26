import { useState } from "react";
import {
  Legend,
  TableTwin,
  Tooltip,
  type LegendItem,
  type TableColumn,
  type TableData,
  type TipRow,
  type TipState,
  useWidth,
  hoverProps,
} from "./common";
import { fmtNum } from "./format";
import { barPath, linear, niceTicks, textWidth } from "./scale";

export interface BarDatum {
  id: string;
  label: string;
  value: number | null;
  /** Fill override (an ordinal step, a group colour, the "Other" grey); default is the single series hue. */
  color?: string;
  /** Extra tooltip rows after the value (e.g. the split by group). */
  tip?: TipRow[];
  /** Extra table cells, keyed like `extraColumns`. */
  extra?: Record<string, string | number | null>;
}

export interface BarChartProps {
  title: string;
  data: BarDatum[];
  orientation?: "horizontal" | "vertical";
  format?: (v: number | null) => string;
  valueLabel?: string;
  /** Fixed upper end of the value axis (e.g. 1 for shares); default from the data. */
  domainMax?: number;
  extraColumns?: TableColumn[];
  legend?: LegendItem[];
  tableView?: boolean;
  /** Plot height of a vertical chart. */
  height?: number;
}

export function barTable(
  p: Pick<BarChartProps, "title" | "data" | "format" | "valueLabel" | "extraColumns">,
): TableData {
  const format = p.format ?? fmtNum;
  return {
    caption: p.title,
    columns: [
      { key: "label", label: "Category" },
      { key: "value", label: p.valueLabel ?? "Count", numeric: true },
      ...(p.extraColumns ?? []),
    ],
    rows: p.data.map((d) => ({ label: d.label, value: format(d.value), ...(d.extra ?? {}) })),
  };
}

export type LabelPlacement = "inside" | "outside" | "tooltip";

/** Where a bar's value label goes: inside the data end when it fits with padding, else past the tip, else only in the tooltip. */
export function labelPlacement(barLen: number, labelW: number, room: number, insideAllowed: boolean): LabelPlacement {
  if (insideAllowed && barLen >= labelW + 12) return "inside";
  if (room >= labelW + 6) return "outside";
  return "tooltip";
}

const ROW = 28; // band per horizontal bar: a 24 px+ hit target
const THICK = 16;

export function BarChart(props: BarChartProps) {
  const { title, data, orientation = "horizontal", tableView, legend } = props;
  const format = props.format ?? fmtNum;
  const { ref, width } = useWidth<HTMLDivElement>();
  const [hot, setHot] = useState<number | null>(null);

  if (tableView) return <TableTwin {...barTable(props)} />;

  const max = props.domainMax ?? Math.max(1, ...data.map((d) => d.value ?? 0));
  const tipFor = (i: number, x: number, y: number): TipState => ({
    x,
    y,
    title: data[i].label,
    rows: [{ value: format(data[i].value), label: props.valueLabel ?? "" }, ...(data[i].tip ?? [])],
  });

  let svg;
  let tip: TipState | null = null;
  if (orientation === "horizontal") {
    const labelW = Math.min(Math.max(...data.map((d) => textWidth(d.label)), 40) + 14, width * 0.4);
    const valueRoom = 44;
    const plotW = Math.max(60, width - labelW - valueRoom);
    const x = linear([0, max], [labelW, labelW + plotW]);
    const h = data.length * ROW + 4;
    svg = (
      <svg width={width} height={h} role="img" aria-label={title}>
        <line className="viz-axis" x1={labelW} x2={labelW} y1={0} y2={h - 4} />
        {data.map((d, i) => {
          const yTop = i * ROW + (ROW - THICK) / 2;
          const len = Math.max(0, x(d.value ?? 0) - labelW);
          const text = format(d.value);
          const tw = textWidth(text);
          const place = labelPlacement(len, tw, labelW + plotW + valueRoom - (labelW + len), !d.color);
          return (
            <g key={d.id}>
              <text className="viz-label" x={labelW - 8} y={yTop + THICK / 2 + 4} textAnchor="end">
                {clip(d.label, labelW - 10)}
              </text>
              <path
                className="viz-mark"
                d={barPath("horizontal", labelW, yTop, len, THICK)}
                fill={d.color ?? "var(--viz-single)"}
                opacity={hot === null || hot === i ? 1 : 0.55}
              />
              {place === "inside" && (
                <text className="viz-inside" x={labelW + len - 6} y={yTop + THICK / 2 + 4} textAnchor="end">
                  {text}
                </text>
              )}
              {place === "outside" && (
                <text className="viz-label" x={labelW + len + 6} y={yTop + THICK / 2 + 4}>
                  {text}
                </text>
              )}
              <rect
                className="viz-hit"
                x={0}
                y={i * ROW}
                width={width}
                height={ROW}
                tabIndex={0}
                aria-label={`${d.label}: ${text}`}
                {...hoverProps(
                  () => setHot(i),
                  () => setHot(null),
                  (d.value ?? 0) > 0,
                )}
              />
            </g>
          );
        })}
      </svg>
    );
    if (hot !== null) {
      const len = Math.max(0, x(data[hot].value ?? 0) - labelW);
      tip = tipFor(hot, Math.min(labelW + len, width - 160), hot * ROW + ROW);
    }
  } else {
    const plotH = props.height ?? 160;
    const left = 30;
    const top = 18;
    const bottom = 22;
    const plotW = Math.max(60, width - left - 4);
    const band = plotW / Math.max(1, data.length);
    const thick = Math.min(24, Math.max(6, band * 0.6));
    const y = linear([0, max], [top + plotH, top]);
    const counts = data.every((d) => d.value === null || Number.isInteger(d.value));
    const ticks = niceTicks(0, max, 4).filter((t) => !counts || Number.isInteger(t));
    svg = (
      <svg width={width} height={top + plotH + bottom} role="img" aria-label={title}>
        {ticks.map((t) => (
          <g key={t}>
            <line className="viz-grid" x1={left} x2={left + plotW} y1={y(t)} y2={y(t)} />
            <text className="viz-tick" x={left - 6} y={y(t) + 4} textAnchor="end">
              {format(t)}
            </text>
          </g>
        ))}
        <line className="viz-axis" x1={left} x2={left + plotW} y1={top + plotH} y2={top + plotH} />
        {data.map((d, i) => {
          const cx = left + band * i + band / 2;
          const yv = y(d.value ?? 0);
          const len = top + plotH - yv;
          const text = format(d.value);
          const fitsInside = !d.color && len >= 20 && thick >= textWidth(text) + 6;
          return (
            <g key={d.id}>
              <path
                className="viz-mark"
                d={barPath("vertical", cx - thick / 2, yv, len, thick)}
                fill={d.color ?? "var(--viz-single)"}
                opacity={hot === null || hot === i ? 1 : 0.55}
              />
              {d.value !== null &&
                (fitsInside ? (
                  <text className="viz-inside" x={cx} y={yv + 15} textAnchor="middle">
                    {text}
                  </text>
                ) : textWidth(text) <= band ? (
                  <text className="viz-label" x={cx} y={yv - 5} textAnchor="middle">
                    {text}
                  </text>
                ) : null)}
              <text className="viz-tick" x={cx} y={top + plotH + 15} textAnchor="middle">
                {clip(d.label, band - 2)}
              </text>
              <rect
                className="viz-hit"
                x={cx - Math.max(12, band / 2)}
                y={top}
                width={Math.max(24, band)}
                height={plotH + bottom}
                tabIndex={0}
                aria-label={`${d.label}: ${text}`}
                {...hoverProps(
                  () => setHot(i),
                  () => setHot(null),
                  (d.value ?? 0) > 0,
                )}
              />
            </g>
          );
        })}
      </svg>
    );
    if (hot !== null) tip = tipFor(hot, left + band * hot + band / 2, top);
  }

  return (
    <figure className="viz" ref={ref}>
      {legend && legend.length >= 2 && <Legend items={legend} />}
      {svg}
      <Tooltip tip={tip} width={width} />
    </figure>
  );
}

/** Shortens a label to fit `px`; the full label stays in the tooltip and the table. */
export function clip(s: string, px: number): string {
  if (textWidth(s) <= px) return s;
  let t = s;
  while (t.length > 2 && textWidth(t + "...") > px) t = t.slice(0, -1);
  return t + "...";
}
