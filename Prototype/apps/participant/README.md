# Participant app

The participant-facing React app of the boardroom study ("02 - Architecture.md" section 9). Vite, React 19 and TypeScript, styled with plain CSS modules (no UI framework, no Tailwind). Shared types come from `@study/shared` in `packages/shared`, linked as a `file:` dependency.

## Run

```
cd apps/participant
npm install
npm run dev        # http://localhost:5173/s/<run code>
```

The dev server proxies `/api`, `/assets` and `/a/` to the Go service on `http://localhost:8080`. Without the backend the app shows "Loading..." and then an error, so use the tests below.

## Build, typecheck, test

```
npm run build      # tsc --noEmit, then vite build into ../../services/api/web/participant (emptied first)
npm run typecheck  # tsc --noEmit
npm test           # vitest (jsdom), mocked step descriptor sequences
```

The bundle's own files are written under `static/` (not Vite's default `assets/`), so they never collide with the content asset path `/assets/*`. The Go service embeds the output folder and serves `index.html` for `/` and every path under `/s/`.

## How it works

- `/s/:code` creates or resumes the session (`POST /api/v1/runs/{code}/sessions`), loads the public scene pack from `scenePackUrl`, then calls `GET /api/v1/sessions/{id}/next` and renders the step that comes back. The client never decides the next step: every event response (and every error carrying `next`) replaces the current step.
- `/s/:code/blocked` shows the device gate message; any other path shows "Study link needed".
- Every call sends `X-Study-Client: 1` and the cookie. `clientSeq` is always the `expectedClientSeq` of the last response. All calls go through one queue, so events leave in order. A network failure is retried with the identical body.
- One hidden `<audio>` element for the page life, no controls, no replay. Playback starts when the `scene_started` response arrives; `ended` sends `audio_ended` (resent after `too_early`). After a reload, or a step with `restart: true`, a "Continue" click provides the new user gesture first.
- Countdowns come from `opensInMs` and `closesInMs`. When a timer fires, the client sends what it has (`rating_timeout`, `action_timeout`, `budget_timeout`, `background_timeout`); the server enforces every deadline itself.

## Step kinds and screens

| `next.scene` | Component | What it does |
|---|---|---|
| `deviceGate` | `scenes/DeviceGate.tsx` | Measures width, height and touch, sends `device_checked`; below 1024 px, portrait or phone it goes to `/s/:code/blocked` |
| `consent` | `scenes/Consent.tsx` | The SFU consent form inline (iframe) with a download link, checkbox, accept button disabled until ticked and the PDF's HEAD check passed; a notice with "Try again" if the form cannot be loaded; `consent_accepted` |
| `audioCheck` | `scenes/AudioCheck.tsx` | Play button, replay as often as wanted (`audio_check_played` each time), four words, one answer (`audio_check_answered`) |
| `example`, `opening`, `slide`, `question`, `answer` | `scenes/AudioScene.tsx` | Boardroom with one clip: `scene_started`, playback, `audio_ended`; checker markings on slides when `overlay` is present (`marking_shown`) |
| `preStart` | `scenes/PreStart.tsx` | The `preStartText` lines from the descriptor (the checker sentence only when sent) and the start button; `setup_done` |
| `postView` | `scenes/PostView.tsx` | The scene stays on screen with a countdown from `opensInMs`, then `GET /next` |
| `rating` | `scenes/Rating.tsx` | Trust 0 to 10 and three 7-point antonym rows, nothing preselected, one submit; `rating_submitted` or `rating_timeout` |
| `action` | `scenes/Action.tsx` | The options the descriptor carries (three or two), one click; `action_chosen` or `action_timeout` |
| `budget` | `scenes/Budget.tsx` | Approve or reject plus 0 to 10 sureness; `budget_submitted` or `budget_timeout` |
| `background` | `scenes/Background.tsx` | Seven items, each with "Prefer not to say", hidden honeypot; `background_submitted` or `background_timeout` |
| `debrief` | `scenes/Debrief.tsx` | Paragraphs, do-not-discuss sentence, session code; sends `finished` after 6 s |
| `done` | `scenes/Debrief.tsx` (`Done`) | Keeps the debrief on screen, or a short closing line |

The group is never sent to or inferred by the client. The checker panel renders only when a step carries `overlay`, with the panel title and idle text that the overlay carries (they are not in the public pack), and the action screen renders whatever options the step lists. A `done` step with `blocked: true` goes to `/s/:code/blocked`.

## Asset paths the app expects

| Path | Source | If missing |
|---|---|---|
| `/assets/consent/<consent.pdfAssetKey>` | the SFU consent form (study 30004094, version 4) | a notice "The consent form could not be loaded" with "Try again"; the accept button stays disabled |
| `/assets/images/boardroom.jpg` | boardroom background | a drawn plain room |
| `/assets/images/ciso.png` | CISO portrait | a grey silhouette labelled "CISO" |
| `/assets/images/chair.png` | chair portrait | a grey silhouette labelled "Chair" |
| `/assets/images/member.png` | board member portrait | a grey silhouette labelled "Board member" |
| `<assetBaseUrl>/<assetKey>` | slide images from the config (for v1.3: `s3-network-diagram.svg`) | a dashed box with the image's alt text |
| `scenePackUrl` (for example `/a/scenes/1.3.0.json`) | the public scene pack | the app cannot start |
| `audioUrl` and `prefetch` of each step | hashed MP3s | `audio_problem`, then "Continue" |

## Design decisions

Every point where the architecture document is silent or contradicts itself is marked `// DECISION:` in the code. Search for it with `grep -rn "DECISION:" src ../../packages/shared/src`.
