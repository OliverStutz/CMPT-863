// Device gate (2.9, section 9.6): width at least `deviceGate.minWidthPx` (1024), landscape, not a phone.
// The client measures and sends device_checked; the server decides. DECISION: when the local check
// already fails, the client still sends device_checked (so the server records Blocked) and then shows
// the blocked page at /s/:code/blocked without waiting for the server's step.

import { useEffect, useRef } from "react";
import type { Step } from "@study/shared";
import { useStudy } from "../study";
import { Screen } from "../ui/Screen";
import styles from "./scenes.module.css";

export interface DeviceInfo {
  widthPx: number;
  heightPx: number;
  touch: boolean;
  phone: boolean;
  landscape: boolean;
}

export function measureDevice(): DeviceInfo {
  const widthPx = window.innerWidth;
  const heightPx = window.innerHeight;
  const touch = (navigator.maxTouchPoints ?? 0) > 0 || "ontouchstart" in window;
  const coarse = typeof window.matchMedia === "function" && window.matchMedia("(pointer: coarse)").matches;
  const shortSide = Math.min(window.screen?.width ?? widthPx, window.screen?.height ?? heightPx);
  return { widthPx, heightPx, touch, phone: coarse && shortSide < 600, landscape: widthPx > heightPx };
}

export function devicePasses(d: DeviceInfo, gate: { minWidthPx: number; allowPhones: boolean }): boolean {
  return d.widthPx >= gate.minWidthPx && d.landscape && (gate.allowPhones || !d.phone);
}

export function DeviceGate({ step: _step }: { step: Step }) {
  const { api, pack, reportError, block } = useStudy();
  const sent = useRef(false);
  useEffect(() => {
    if (sent.current) return;
    sent.current = true;
    const gate = pack.options?.deviceGate ?? { minWidthPx: 1024, allowPhones: false };
    const d = measureDevice();
    const pass = devicePasses(d, gate);
    api
      .send("device_checked", { widthPx: d.widthPx, heightPx: d.heightPx, touch: d.touch })
      .then(() => {
        if (!pass) block();
      })
      .catch((e) => (pass ? reportError(e) : block()));
  }, [api, pack, reportError, block]);
  return (
    <Screen title="Checking your screen" titleHidden>
      <p className={styles.waiting} aria-live="polite">
        &nbsp;
      </p>
    </Screen>
  );
}
