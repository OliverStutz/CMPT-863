import { useEffect, useRef, useState } from "react";

/**
 * Counts down `ms` from mount (section 9.4: countdowns come from the server's relative opensInMs /
 * closesInMs, so the client clock is never trusted). Calls `onExpire` once at zero. Returns the
 * remaining milliseconds, updated a few times a second. `ms` undefined means no deadline.
 */
export function useDeadline(ms: number | undefined, onExpire?: () => void): number | undefined {
  const deadline = useRef<number | undefined>(undefined);
  const expired = useRef(false);
  const cb = useRef(onExpire);
  cb.current = onExpire;
  const [left, setLeft] = useState<number | undefined>(ms === undefined ? undefined : Math.max(0, ms));

  useEffect(() => {
    if (ms === undefined) return;
    deadline.current = performance.now() + Math.max(0, ms);
    expired.current = false;
    const tick = () => {
      const rem = Math.max(0, (deadline.current ?? 0) - performance.now());
      setLeft(rem);
      if (rem <= 0 && !expired.current) {
        expired.current = true;
        cb.current?.();
      }
    };
    tick();
    const t = setInterval(tick, 200);
    return () => clearInterval(t);
  }, [ms]);

  return left;
}
