// Fixed asset paths the app expects. Content assets named in the config (slide images, audio) come from
// the server: `assetBaseUrl` + key, or the step's `audioUrl`.

// The consent PDF (the SFU form, study 30004094) is served by the Go service at
// /assets/consent/<consent.pdfAssetKey> (the server's public prefix consent/*.pdf). If the file cannot be
// loaded, the consent page shows a notice and keeps the accept button disabled.
export function consentPdfUrl(pdfAssetKey: string | undefined): string {
  return `/assets/consent/${encodeURIComponent(pdfAssetKey || "placeholder.pdf")}`;
}

// Boardroom composite (section 9.2). Each is optional: until the file exists the app draws a neutral
// placeholder (SVG silhouette labelled by role, or a plain room). `member` is the participant's own seat
// ("Board member (You)", lit while the board-member voice speaks the participant's question);
// member2 to member6 are the silent board members that fill the table (Oliver, 2026-10-25).
export const IMAGE_URLS = {
  boardroom: "/assets/images/boardroom.jpg",
  ciso: "/assets/images/ciso.png",
  chair: "/assets/images/chair.png",
  member: "/assets/images/member.png",
  member2: "/assets/images/member2.png",
  member3: "/assets/images/member3.png",
  member4: "/assets/images/member4.png",
  member5: "/assets/images/member5.png",
  member6: "/assets/images/member6.png",
} as const;

export type PortraitRole = Exclude<keyof typeof IMAGE_URLS, "boardroom">;
