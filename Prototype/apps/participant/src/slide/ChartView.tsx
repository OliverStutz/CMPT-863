// Minimal chart drawing for `chart` blocks (bar, line, timeline). No v1.3 slide uses a chart; every
// series and marker carries its id so a marking can target it.

import type { ChartBlock } from "@study/shared";

const W = 794;
const H = 500;
const PAD = { l: 70, r: 20, t: 20, b: 60 };

export function ChartView({ block }: { block: ChartBlock }) {
  const { axis, series, markers } = block.data;
  const span = axis.max - axis.min || 1;
  const y = (v: number) => PAD.t + (H - PAD.t - PAD.b) * (1 - (v - axis.min) / span);
  const x = (v: number) => PAD.l + (W - PAD.l - PAD.r) * ((v - axis.min) / span);
  const innerW = W - PAD.l - PAD.r;

  return (
    <svg viewBox={`0 0 ${W} ${H}`} width={W} height={H} role="img" aria-label={block.alt}>
      <line x1={PAD.l} y1={H - PAD.b} x2={W - PAD.r} y2={H - PAD.b} stroke="#5f6b7a" strokeWidth={2} />
      {block.chartType !== "timeline" && (
        <line x1={PAD.l} y1={PAD.t} x2={PAD.l} y2={H - PAD.b} stroke="#5f6b7a" strokeWidth={2} />
      )}
      {axis.label && (
        <text x={W / 2} y={H - 12} fontSize={20} textAnchor="middle" fill="#3e4c59">
          {axis.label}
        </text>
      )}
      {block.chartType === "bar" &&
        series.map((s, i) => {
          const bw = innerW / series.length;
          const v = s.value ?? axis.min;
          return (
            <g key={s.id} id={s.id} data-seg-id={s.id}>
              <rect
                x={PAD.l + i * bw + bw * 0.2}
                y={y(v)}
                width={bw * 0.6}
                height={H - PAD.b - y(v)}
                fill="#14325a"
              />
              <text x={PAD.l + i * bw + bw / 2} y={H - PAD.b + 24} fontSize={18} textAnchor="middle" fill="#1f2933">
                {s.label}
              </text>
            </g>
          );
        })}
      {block.chartType === "line" &&
        series.map((s, si) => {
          const pts = s.points ?? [];
          const step = pts.length > 1 ? innerW / (pts.length - 1) : 0;
          const d = pts.map((p, i) => `${i ? "L" : "M"}${PAD.l + i * step},${y(p)}`).join(" ");
          return (
            <g key={s.id} id={s.id} data-seg-id={s.id}>
              <path d={d} fill="none" stroke={si % 2 ? "#d9822b" : "#14325a"} strokeWidth={4} />
              <text x={W - PAD.r} y={pts.length ? y(pts[pts.length - 1]) - 10 : PAD.t} fontSize={18} textAnchor="end">
                {s.label}
              </text>
            </g>
          );
        })}
      {block.chartType === "timeline" &&
        series.map((s, i) => {
          const rowH = (H - PAD.t - PAD.b) / Math.max(1, series.length);
          const x0 = x(s.start ?? axis.min);
          const x1 = x(s.end ?? s.start ?? axis.min);
          return (
            <g key={s.id} id={s.id} data-seg-id={s.id}>
              <rect x={x0} y={PAD.t + i * rowH + rowH * 0.25} width={Math.max(4, x1 - x0)} height={rowH * 0.5} fill="#14325a" />
              <text x={x0 + 8} y={PAD.t + i * rowH + rowH * 0.25 - 6} fontSize={18} fill="#1f2933">
                {s.label}
              </text>
            </g>
          );
        })}
      {markers?.map((m) => (
        <g key={m.id} id={m.id} data-seg-id={m.id}>
          <line
            x1={x(m.at)}
            y1={PAD.t}
            x2={x(m.at)}
            y2={H - PAD.b}
            stroke="#d9822b"
            strokeWidth={3}
            strokeDasharray="10 8"
          />
          <text x={x(m.at) + 6} y={PAD.t + 18} fontSize={18} fill="#1f2933">
            {m.label}
          </text>
        </g>
      ))}
    </svg>
  );
}
