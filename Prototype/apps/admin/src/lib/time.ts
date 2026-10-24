/** Times are shown in the browser's local zone; the zone is named once in the header. */
export const localZone = Intl.DateTimeFormat().resolvedOptions().timeZone;

const fmt = new Intl.DateTimeFormat(undefined, {
  year: "numeric",
  month: "2-digit",
  day: "2-digit",
  hour: "2-digit",
  minute: "2-digit",
  second: "2-digit",
  hour12: false,
});

export function localTime(iso: string | null | undefined): string {
  if (!iso) return "-";
  const d = new Date(iso);
  return Number.isNaN(d.getTime()) ? "-" : fmt.format(d);
}

// DECISION: the session column shows the first 8 characters of the session id; the full id is in the
// row's title and in the detail drawer.
export function shortId(id: string): string {
  return id.length > 8 ? id.slice(0, 8) : id;
}
