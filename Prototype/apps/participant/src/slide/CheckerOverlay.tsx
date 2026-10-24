// The checker (section 9.2, `<CheckerOverlay>`; content v1.3 section 4). Present only when the step
// descriptor carries `overlay`, which the server sends for checker sessions only; the component itself
// is in every bundle, so the code does not tell a participant their group.
//
// Two parts:
//  - CheckerMarks, inside the slide: at each marking's appearAtMs (read from the audio element's
//    currentTime, not a separate timer) it draws the marking over every target element (outline plus a
//    translucent fill in the marking's colour, with a short draw-on animation). Shape plus colour, so
//    colour is not the only cue (9.6).
//  - CheckerPanel, beside the slide's top right corner: the title before anything is shown
//    (`labels.panelIdle` under it if that is non-empty, C-3), then one label and detail per marking,
//    or the "No issues found" text with a tick at its time on unflagged slides.

import { useEffect, useRef, useState, type RefObject } from "react";
import type { ResolvedMarking, StepOverlay } from "@study/shared";
import { SLIDE_WIDTH } from "@study/shared";
import styles from "./CheckerOverlay.module.css";

interface Rect {
  x: number;
  y: number;
  w: number;
  h: number;
}

const PAD = 10;

/** Finds each target in the slide and returns its box in slide coordinates; falls back to bbox. */
function measure(root: HTMLElement, m: ResolvedMarking): Rect[] {
  const rootBox = root.getBoundingClientRect();
  const scale = rootBox.width > 0 ? rootBox.width / SLIDE_WIDTH : 1;
  const rects: Rect[] = [];
  for (const id of m.targetIds) {
    const el = root.querySelector<HTMLElement>(`[data-seg-id="${CSS.escape(id)}"]`);
    if (!el) continue;
    const b = el.getBoundingClientRect();
    if (b.width === 0 && b.height === 0) continue;
    rects.push({
      x: (b.left - rootBox.left) / scale - PAD,
      y: (b.top - rootBox.top) / scale - PAD,
      w: b.width / scale + 2 * PAD,
      h: b.height / scale + 2 * PAD,
    });
  }
  if (rects.length === 0 && m.bbox) {
    const [x, y, w, h] = m.bbox;
    rects.push({ x, y, w, h });
  }
  return rects;
}

export function visibleMarkings(overlay: StepOverlay | undefined, clockMs: number): ResolvedMarking[] {
  if (!overlay || overlay.status !== "flagged") return [];
  return overlay.markings.filter((m) => clockMs >= m.appearAtMs);
}

interface MarksProps {
  overlay: StepOverlay | undefined;
  clockMs: number;
  slideRef: RefObject<HTMLDivElement | null>;
  onShown?: (markingId: string) => void;
}

export function CheckerMarks({ overlay, clockMs, slideRef, onShown }: MarksProps) {
  const shown = visibleMarkings(overlay, clockMs);
  const [rects, setRects] = useState<Record<string, Rect[]>>({});
  const reported = useRef(new Set<string>());
  const shownKey = shown.map((m) => m.id).join("|");

  // Measured after commit (a passive effect), when the slide's ref is attached; again on resize.
  useEffect(() => {
    const run = () => {
      const root = slideRef.current;
      if (!root) return;
      const next: Record<string, Rect[]> = {};
      for (const m of shown) next[m.id] = measure(root, m);
      setRects(next);
    };
    run();
    const raf = typeof requestAnimationFrame === "function" ? requestAnimationFrame(run) : undefined;
    window.addEventListener("resize", run);
    return () => {
      if (raf !== undefined) cancelAnimationFrame(raf);
      window.removeEventListener("resize", run);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [shownKey, slideRef]);

  useEffect(() => {
    for (const m of shown) {
      if (!reported.current.has(m.id)) {
        reported.current.add(m.id);
        onShown?.(m.id);
      }
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [shownKey]);

  if (shown.length === 0) return null;
  return (
    <svg className={styles.marks} viewBox="0 0 1600 900" width={1600} height={900} aria-hidden="true">
      {shown.map((m) =>
        (rects[m.id] ?? []).map((r, i) => <Shape key={`${m.id}-${i}`} m={m} r={r} />),
      )}
    </svg>
  );
}

function Shape({ m, r }: { m: ResolvedMarking; r: Rect }) {
  const common = { stroke: m.color, strokeWidth: 5, pathLength: 1, className: styles.draw };
  if (m.shape === "circle") {
    return (
      <g data-marking-id={m.id}>
        <ellipse cx={r.x + r.w / 2} cy={r.y + r.h / 2} rx={r.w / 2 + 14} ry={r.h / 2 + 14} fill={m.color} fillOpacity={0.16} {...common} />
      </g>
    );
  }
  if (m.shape === "underline") {
    return (
      <g data-marking-id={m.id}>
        <rect x={r.x} y={r.y} width={r.w} height={r.h} fill={m.color} fillOpacity={0.14} stroke="none" />
        <line x1={r.x} y1={r.y + r.h} x2={r.x + r.w} y2={r.y + r.h} strokeLinecap="round" {...common} strokeWidth={7} />
      </g>
    );
  }
  return (
    <g data-marking-id={m.id}>
      <rect x={r.x} y={r.y} width={r.w} height={r.h} rx={8} fill={m.color} fillOpacity={0.16} {...common} />
    </g>
  );
}

interface PanelProps {
  overlay: StepOverlay;
  clockMs: number;
  title: string;
  idleText: string;
  noIssuesColor: string;
}

export function CheckerPanel({ overlay, clockMs, title, idleText, noIssuesColor }: PanelProps) {
  const marks = visibleMarkings(overlay, clockMs);
  const noIssuesShown = overlay.status === "noIssues" && clockMs >= overlay.appearAtMs;
  const idle = marks.length === 0 && !noIssuesShown;
  return (
    <aside className={styles.panel} aria-label={title} aria-live="polite" data-testid="checker-panel">
      <h3 className={styles.panelTitle}>{title}</h3>
      {idle && idleText && <p className={styles.idle}>{idleText}</p>}
      {noIssuesShown && overlay.status === "noIssues" && (
        <p className={styles.noIssues}>
          <svg width="28" height="28" viewBox="0 0 28 28" aria-hidden="true">
            <circle cx="14" cy="14" r="13" fill={noIssuesColor} />
            <path d="M8 14.5l4 4 8-9" fill="none" stroke="#fff" strokeWidth="3" strokeLinecap="round" strokeLinejoin="round" />
          </svg>
          <span>{overlay.text}</span>
        </p>
      )}
      {marks.map((m) => (
        <div key={m.id} className={styles.callout} style={{ borderLeftColor: m.color }}>
          <div className={styles.calloutLabel}>
            <ShapeIcon m={m} />
            <span>{m.label}</span>
          </div>
          <p className={styles.calloutDetail}>{m.detail}</p>
        </div>
      ))}
    </aside>
  );
}

function ShapeIcon({ m }: { m: ResolvedMarking }) {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" aria-hidden="true" className={styles.icon}>
      {m.shape === "circle" ? (
        <ellipse cx="12" cy="12" rx="9" ry="9" fill={m.color} fillOpacity={0.25} stroke={m.color} strokeWidth="3" />
      ) : m.shape === "underline" ? (
        <line x1="3" y1="18" x2="21" y2="18" stroke={m.color} strokeWidth="4" strokeLinecap="round" />
      ) : (
        <rect x="3" y="4" width="18" height="16" rx="2" fill={m.color} fillOpacity={0.25} stroke={m.color} strokeWidth="3" />
      )}
    </svg>
  );
}
