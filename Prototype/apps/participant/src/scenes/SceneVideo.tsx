// Video layers of the boardroom (talking heads, 2026-10-25).
//
// `SyncedVideo` is a muted video of the clip the page's one audio element is playing (player.ts). The
// audio element stays the source of truth for timing, events and the autoplay unlock; the video only
// follows its clock, so a video that fails to load or stalls changes nothing in the study flow. A muted
// video may start without a user gesture in every browser.
//
// `IdleVideo` is a silent loop for a seat that never speaks.

import { useEffect, useRef } from "react";
import { player } from "../audio/player";

/** Re-seek the video when it drifts further than this from the audio clock. */
const DRIFT_MS = 150;

export function SyncedVideo({
  src,
  startMs = 0,
  endMs,
  onFailed,
}: {
  src: string;
  /** Offset of this video in the audio clip (a segment of a multi-speaker clip); 0 for a whole clip. */
  startMs?: number;
  /** Where the segment ends in the audio clip; undefined = the end of the audio. */
  endMs?: number;
  onFailed: () => void;
}) {
  const ref = useRef<HTMLVideoElement>(null);
  useEffect(() => {
    const v = ref.current;
    if (!v) return;
    v.muted = true;
    const audio = player.element;
    const follow = () => {
      const at = audio.currentTime * 1000;
      const inSegment = at >= startMs && (endMs === undefined || at < endMs);
      if (audio.paused || audio.ended || !inSegment) {
        if (!v.paused) v.pause();
        if (at < startMs && v.currentTime !== 0) v.currentTime = 0; // before the segment: its first frame
        return;
      }
      const want = (at - startMs) / 1000;
      if (Math.abs(v.currentTime - want) * 1000 > DRIFT_MS) v.currentTime = want;
      if (v.paused) void v.play().catch(() => undefined);
    };
    const unsubscribe = player.subscribe(follow);
    const id = setInterval(follow, 500);
    follow();
    return () => {
      unsubscribe();
      clearInterval(id);
      v.pause();
    };
  }, [src, startMs, endMs]);
  return <video ref={ref} src={src} muted playsInline preload="auto" onError={onFailed} data-testid="speaker-video" />;
}

export function IdleVideo({ src, onFailed }: { src: string; onFailed: () => void }) {
  const ref = useRef<HTMLVideoElement>(null);
  useEffect(() => {
    const v = ref.current;
    if (!v) return;
    v.muted = true;
    void v.play().catch(() => undefined);
  }, [src]);
  return <video ref={ref} src={src} muted loop autoPlay playsInline preload="auto" onError={onFailed} data-testid="idle-video" />;
}
