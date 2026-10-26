import { useCallback, useEffect, useState, type DependencyList } from "react";

/**
 * Loads data on mount, on reload() and whenever `deps` change; errors are kept for display.
 * The previous data stays while a new load runs (`loading` is true), so a view can hold its last render.
 * `initial` seeds the data before the first load (a preview, or a server-side render in a test).
 */
export function useAsync<T>(load: () => Promise<T>, deps: DependencyList = [], initial: T | null = null) {
  const [data, setData] = useState<T | null>(initial);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);
  const [tick, setTick] = useState(0);
  const reload = useCallback(() => setTick((n) => n + 1), []);
  useEffect(() => {
    let alive = true;
    setLoading(true);
    load().then(
      (d) => {
        if (alive) {
          setData(d);
          setError(null);
          setLoading(false);
        }
      },
      (e: unknown) => {
        if (alive) {
          setError(e instanceof Error ? e.message : String(e));
          setLoading(false);
        }
      },
    );
    return () => {
      alive = false;
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [tick, ...deps]);
  return { data, error, loading, reload };
}
