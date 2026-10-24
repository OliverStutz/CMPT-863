import styles from "./Countdown.module.css";

interface Props {
  label: string;
  leftMs: number | undefined;
  totalMs: number | undefined;
}

/** Visible countdown: label, whole seconds left and a shrinking bar. */
export function Countdown({ label, leftMs, totalMs }: Props) {
  if (leftMs === undefined) return null;
  const secs = Math.ceil(leftMs / 1000);
  const frac = totalMs && totalMs > 0 ? Math.max(0, Math.min(1, leftMs / totalMs)) : 0;
  return (
    <div className={styles.countdown} role="timer" aria-label={`${label}: ${secs} seconds`}>
      <div className={styles.text} aria-hidden="true">
        {label}: <strong>{secs}</strong> s
      </div>
      <div className={styles.track} aria-hidden="true">
        <div className={`${styles.bar} ${secs <= 5 ? styles.low : ""}`} style={{ width: `${frac * 100}%` }} />
      </div>
    </div>
  );
}
