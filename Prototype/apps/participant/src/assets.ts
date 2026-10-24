// Fixed asset paths the app expects. Content assets named in the config (slide images, audio) come from
// the server: `assetBaseUrl` + key, or the step's `audioUrl`.

// DECISION: the integration brief fixes the consent PDF at this path. Section 4.3 would name it by hash
// under the asset base URL (config `consent.pdfAssetKey`); the brief's fixed path is used until the
// SFU ethics form arrives (OPEN M). If the file is missing, the consent page shows a placeholder box.
export const CONSENT_PDF_URL = "/assets/consent/placeholder.pdf";

// Boardroom composite (section 9.2). Each is optional: until the file exists the app draws a neutral
// placeholder (SVG silhouette labelled by role, or a plain room).
export const IMAGE_URLS = {
  boardroom: "/assets/images/boardroom.jpg",
  ciso: "/assets/images/ciso.png",
  chair: "/assets/images/chair.png",
  member: "/assets/images/member.png",
} as const;
