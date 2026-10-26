export interface StatTileProps {
  label: string;
  value: string;
  caption?: string;
  /** Signed change against a named reference, e.g. { text: "+3 since yesterday", good: true }. */
  delta?: { text: string; good: boolean };
}

/** A headline number: label, value in proportional figures, optional caption and delta. */
export function StatTile({ label, value, caption, delta }: StatTileProps) {
  return (
    <div className="stat-tile">
      <span className="stat-label">{label}</span>
      <span className="stat-value">{value}</span>
      {delta && <span className={delta.good ? "stat-delta-good" : "stat-delta-bad"}>{delta.text}</span>}
      {caption && <span className="stat-caption">{caption}</span>}
    </div>
  );
}
