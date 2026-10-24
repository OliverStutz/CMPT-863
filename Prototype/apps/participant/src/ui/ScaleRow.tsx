import styles from "./ScaleRow.module.css";

export interface ScaleExtra {
  value: string;
  label: string;
}

interface Props {
  name: string;
  /** Accessible group name. Shown as a heading unless `legendHidden`. */
  legend: string;
  legendHidden?: boolean;
  min: number;
  max: number;
  minLabel: string;
  maxLabel: string;
  pointLabels?: Record<string, string>;
  /** "ends": numbers in a row with the two end labels below; "antonym": left word, numbers, right word. */
  variant: "ends" | "antonym";
  value: number | string | null;
  onChange: (v: number | string) => void;
  /** Extra choices after the numbers in the same radio group, e.g. "Prefer not to say". */
  extras?: ScaleExtra[];
  disabled?: boolean;
}

/**
 * A row of clickable numbers with nothing preselected (2.8), built as a native radio group so it is
 * keyboard reachable and labelled (section 9.6).
 */
export function ScaleRow(p: Props) {
  const points: number[] = [];
  for (let n = p.min; n <= p.max; n++) points.push(n);
  const id = `scale-${p.name}`;
  return (
    <fieldset className={styles.fieldset} disabled={p.disabled} data-testid={id}>
      <legend className={p.legendHidden ? "sr-only" : styles.legend}>{p.legend}</legend>
      <div className={p.variant === "antonym" ? styles.antonymRow : styles.row}>
        {p.variant === "antonym" && (
          <span className={`${styles.pole} ${styles.poleLeft}`} aria-hidden="true">
            {p.minLabel}
          </span>
        )}
        <div className={styles.options}>
          {points.map((n) => {
            const pl = p.pointLabels?.[String(n)];
            const aria =
              n === p.min ? `${n}, ${p.minLabel}` : n === p.max ? `${n}, ${p.maxLabel}` : pl ? `${n}, ${pl}` : String(n);
            return (
              <label key={n} className={styles.option}>
                <input
                  type="radio"
                  className="sr-only"
                  name={id}
                  value={n}
                  checked={p.value === n}
                  onChange={() => p.onChange(n)}
                  aria-label={aria}
                />
                <span className={styles.face} aria-hidden="true">
                  {n}
                </span>
                {p.variant === "ends" && p.pointLabels && (
                  <span className={styles.pointLabel} aria-hidden="true">
                    {pl ?? (n === p.min ? p.minLabel : n === p.max ? p.maxLabel : "")}
                  </span>
                )}
              </label>
            );
          })}
          {p.extras?.map((x) => (
            <label key={x.value} className={`${styles.option} ${styles.extra}`}>
              <input
                type="radio"
                className="sr-only"
                name={id}
                value={x.value}
                checked={p.value === x.value}
                onChange={() => p.onChange(x.value)}
              />
              <span className={styles.face}>{x.label}</span>
            </label>
          ))}
        </div>
        {p.variant === "antonym" && (
          <span className={styles.pole} aria-hidden="true">
            {p.maxLabel}
          </span>
        )}
      </div>
      {p.variant === "ends" && !p.pointLabels && (
        <div className={styles.ends} aria-hidden="true">
          <span>
            {p.min} = {p.minLabel}
          </span>
          <span>
            {p.max} = {p.maxLabel}
          </span>
        </div>
      )}
    </fieldset>
  );
}
