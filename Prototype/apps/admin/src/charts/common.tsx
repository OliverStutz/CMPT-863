import { useLayoutEffect, useRef, useState, type ReactNode } from "react";
import "./charts.css";

/** Width of a container in px, measured before the first paint and followed with a ResizeObserver; `fallback` in tests. */
export function useWidth<T extends HTMLElement>(fallback = 640) {
  const ref = useRef<T | null>(null);
  const [width, setWidth] = useState(fallback);
  useLayoutEffect(() => {
    const el = ref.current;
    if (!el || typeof ResizeObserver === "undefined") return;
    const measure = () => {
      const w = Math.floor(el.getBoundingClientRect().width);
      if (w > 0) setWidth(w);
    };
    measure();
    const ro = new ResizeObserver(measure);
    ro.observe(el);
    return () => ro.disconnect();
  }, []);
  return { ref, width };
}

export interface TipRow {
  /** CSS colour of the series key; omitted for a row without a series. */
  color?: string;
  key?: "line" | "rect" | "dot";
  value: string;
  label?: string;
  sub?: string;
}
export interface TipState {
  x: number;
  y: number;
  title?: string;
  rows: TipRow[];
}

/** One tooltip: values lead (strong), labels follow; series keyed by a short line, never a filled box. */
export function Tooltip({ tip, width }: { tip: TipState | null; width: number }) {
  if (!tip) return null;
  const left = Math.min(Math.max(0, tip.x + 12), Math.max(0, width - 200));
  return (
    <div className="viz-tooltip" style={{ left, top: Math.max(0, tip.y - 8) }} role="status" aria-live="polite">
      {tip.title && <div className="viz-tt-title">{tip.title}</div>}
      {tip.rows.map((r, i) => (
        <div className="viz-tt-row" key={i}>
          {r.color && <Key color={r.color} kind={r.key ?? "line"} />}
          <strong>{r.value}</strong>
          {r.label && <span>{r.label}</span>}
          {r.sub && <span className="viz-tt-sub">{r.sub}</span>}
        </div>
      ))}
    </div>
  );
}

export function Key({ color, kind, hollow }: { color: string; kind: "line" | "rect" | "dot"; hollow?: boolean }) {
  if (kind === "rect") return <span className="viz-key-rect" style={{ background: color }} aria-hidden />;
  if (kind === "dot")
    return (
      <span
        className="viz-key-dot"
        style={hollow ? { border: `2px solid ${color}` } : { background: color }}
        aria-hidden
      />
    );
  return <span className="viz-key-line" style={{ background: color }} aria-hidden />;
}

export interface LegendItem {
  id: string;
  label: string;
  color: string;
  kind?: "line" | "rect" | "dot";
  hollow?: boolean;
}

/** Legend for two or more series; mirrors the mark (line for lines, rect for bars, dot for dots). */
export function Legend({ items }: { items: LegendItem[] }) {
  if (items.length === 0) return null;
  return (
    <ul className="viz-legend">
      {items.map((it) => (
        <li key={it.id}>
          <Key color={it.color} kind={it.kind ?? "rect"} hollow={it.hollow} />
          {it.label}
        </li>
      ))}
    </ul>
  );
}

export interface TableColumn {
  key: string;
  label: string;
  numeric?: boolean;
}
export type TableCell = string | number | null | undefined | ReactNode;
export interface TableData {
  caption: string;
  columns: TableColumn[];
  rows: Record<string, TableCell>[];
}

/** The accessible twin of a chart: the same data as a plain table, numbers in tabular figures. */
export function TableTwin({ caption, columns, rows }: TableData) {
  return (
    <div className="viz-table table-wrap">
      <table>
        <caption>{caption}</caption>
        <thead>
          <tr>
            {columns.map((c) => (
              <th key={c.key} scope="col" className={c.numeric ? "num" : undefined}>
                {c.label}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rows.map((r, i) => (
            <tr key={i}>
              {columns.map((c, j) =>
                j === 0 ? (
                  <th key={c.key} scope="row">
                    {cell(r[c.key])}
                  </th>
                ) : (
                  <td key={c.key} className={c.numeric ? "num" : undefined}>
                    {cell(r[c.key])}
                  </td>
                ),
              )}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

function cell(v: TableCell): ReactNode {
  if (v === null || v === undefined || v === "") return "-";
  return v as ReactNode;
}

/** Position of a pointer event relative to the chart's wrapper. */
export function localPoint(e: { clientX: number; clientY: number }, el: Element | null): { x: number; y: number } {
  if (!el) return { x: 0, y: 0 };
  const r = el.getBoundingClientRect();
  return { x: e.clientX - r.left, y: e.clientY - r.top };
}

/**
 * Hover and focus handlers of one mark. The tooltip opens on a real pointer move over the mark or on keyboard
 * focus, never on pointerenter: a layout shift under a resting pointer (page load, a refetch) fires pointerenter
 * without any movement, and the tooltip must stay hidden until the reader reaches a mark. A mark with nothing to
 * show (`enabled` false, e.g. a zero-length bar) never opens one; its value stays in the label and the table.
 */
export function hoverProps(open: () => void, close: () => void, enabled = true) {
  if (!enabled) return {};
  return {
    onPointerMove: (e: { movementX: number; movementY: number; isTrusted: boolean }) => {
      if (e.isTrusted && (e.movementX !== 0 || e.movementY !== 0)) open();
    },
    onPointerLeave: close,
    onFocus: open,
    onBlur: close,
  };
}
