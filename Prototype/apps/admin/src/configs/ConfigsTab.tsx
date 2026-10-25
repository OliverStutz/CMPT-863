import { useEffect, useState } from "react";
import { getConfig, listConfigs } from "../api/client";
import { AUDIO_MODE_LABEL, audioMode, type StudyConfigLite } from "../api/types";
import { localTime } from "../lib/time";
import { useAsync } from "../lib/useAsync";

// DECISION: the brief makes this tab read-only, so the upload (POST /api/admin/configs) and the slide preview
// of section 10 are left out; the preview also needs <SlideRenderer> from packages/shared, which this app does
// not depend on.

export interface SlideSummary {
  slideCount: number;
  flagged: { n: number; id: string }[];
}

/** Slide count and flagged slides of a config (4.1: a slide is flagged exactly when checker.status is "flagged"). */
export function summarizeSlides(cfg: StudyConfigLite): SlideSummary {
  const slides = (cfg.scenes ?? []).filter((s) => s.kind === "slide");
  return {
    slideCount: slides.length,
    flagged: slides.flatMap((s, i) => (s.checker?.status === "flagged" ? [{ n: i + 1, id: s.id }] : [])),
  };
}

/** Chip naming how a config's clips are served (types.ts audioMode). */
export function AudioChip({ mode }: { mode: ReturnType<typeof audioMode> }) {
  if (!mode) return <span className="muted">-</span>;
  const cls = mode === "spoken" ? "chip-success" : mode === "partly-silent" ? "chip-info" : "chip-warning";
  return (
    <span className={`chip ${cls}`} title={mode === "spoken" ? "Every clip has a TTS recording" : "No TTS recording for this version: the server plays silent placeholder clips"}>
      {AUDIO_MODE_LABEL[mode]}
    </span>
  );
}

// DECISION: the config list response (8.2) carries no slide count, so each version's full config is fetched
// from GET /api/admin/configs/{version} and only its slide scenes and checker status are read.
export function ConfigsTab() {
  const configs = useAsync(listConfigs);
  const [summaries, setSummaries] = useState<Record<string, SlideSummary | string>>({});
  const [openLint, setOpenLint] = useState<string | null>(null);

  useEffect(() => {
    if (!configs.data) return;
    let alive = true;
    for (const c of configs.data) {
      getConfig(c.version).then(
        (cfg) => alive && setSummaries((m) => ({ ...m, [c.version]: summarizeSlides(cfg) })),
        (e: unknown) => alive && setSummaries((m) => ({ ...m, [c.version]: e instanceof Error ? e.message : "error" })),
      );
    }
    return () => {
      alive = false;
    };
  }, [configs.data]);

  return (
    <section>
      <h2>Config versions</h2>
      {configs.error && <p className="error">{configs.error}</p>}
      {configs.data && configs.data.length === 0 && <p className="muted">No config versions uploaded.</p>}
      {configs.data && configs.data.length > 0 && (
        <div className="table-wrap">
          <table>
            <thead>
              <tr>
                <th>Version</th>
                <th>Status</th>
                <th>Uploaded</th>
                <th>Slides</th>
                <th>Flagged slides</th>
                <th>Audio</th>
                <th>Lint report</th>
              </tr>
            </thead>
            <tbody>
              {configs.data.map((c) => {
                const s = summaries[c.version];
                return (
                  <tr key={c.version}>
                    <td className="mono">{c.version}</td>
                    <td>
                      <span className={`chip chip-${c.status === "locked" ? "info" : c.status === "validated" ? "success" : "neutral"}`}>
                        {c.status}
                      </span>
                    </td>
                    <td>{localTime(c.uploadedAt)}</td>
                    <td className="mono">{s === undefined ? "..." : typeof s === "string" ? "-" : s.slideCount}</td>
                    <td className="mono">
                      {s === undefined
                        ? "..."
                        : typeof s === "string"
                          ? <span className="error">{s}</span>
                          : s.flagged.length
                            ? s.flagged.map((f) => `${f.n} (${f.id})`).join(", ")
                            : "none"}
                    </td>
                    <td>
                      <AudioChip mode={audioMode(c)} />
                    </td>
                    <td>
                      {c.lintReport === undefined || c.lintReport === null ? (
                        "-"
                      ) : (
                        <button className="small" onClick={() => setOpenLint(openLint === c.version ? null : c.version)}>
                          {openLint === c.version ? "Hide" : "Show"}
                        </button>
                      )}
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </div>
      )}
      {openLint && configs.data && (
        <>
          <h3>Lint report, {openLint}</h3>
          <pre className="json">
            {JSON.stringify(configs.data.find((c) => c.version === openLint)?.lintReport, null, 2)}
          </pre>
        </>
      )}
    </section>
  );
}
