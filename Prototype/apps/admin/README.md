# Admin app

React, Vite and TypeScript, styled with plain global CSS (`src/styles.css`, no CSS framework). It is the researcher's view of the study (section 10 of `02 - Architecture.md`): Runs, Participants, Export and Configs tabs behind the shared admin password. The bundle holds no study content; slide text, flags, ratings and answers come only from the admin API.

## Run

```sh
npm install
npm run dev        # http://localhost:5174/admin/, proxies /api to http://localhost:8080
```

Start the Go service on port 8080 first. The `admin` cookie is set with `Secure`; browsers accept Secure cookies on `http://localhost`, so login works through the dev proxy.

## Build and check

```sh
npm run build      # typecheck, then writes ../../services/api/web/admin (emptied first), base /admin/
npm run typecheck
npm test           # Vitest: elapsed timer and status chip mapping
```

The Go service embeds `services/api/web/admin` and serves it under `/admin/`. The build empties that folder, so nothing else should be kept there.

## Endpoints used

Every call sends `X-Study-Admin: 1` and the `admin` cookie. A `401` shows the login screen; the request is sent again after a successful login.

| Method | Path | Used by |
|---|---|---|
| POST | `/api/admin/login` | login screen, body `{ password }` |
| POST | `/api/admin/logout` | header, "Log out" |
| GET | `/api/admin/runs` | Runs, Participants filter, Export |
| POST | `/api/admin/runs` | Runs, "Create run", body `{ name, source, configVersion }` |
| PATCH | `/api/admin/runs/{id}` | Runs, close and reopen, body `{ status }` |
| GET | `/api/admin/participants` | Participants, query `runId, status, group, flagged=true, updatedSince, limit, cursor`, polled every 5 s |
| GET | `/api/admin/participants/{id}` | Participants, detail drawer |
| GET | `/api/admin/export` | Export, query `format=csv or json, runId, includeInvalid=true, includeIp=true` |
| GET | `/api/admin/configs` | Configs, and the config choice in "New run" |
| GET | `/api/admin/configs/{version}` | Configs, slide count and flagged slides |

## Notes

- Times are shown in the browser's time zone, named once in the header.
- The run link is `<origin>/s/<code>`, with a copy button and a QR code made in the browser (`qrcode`), downloadable as PNG.
- Elapsed time is `(finishedAt or now) - consentAt`, ticks every second between polls, freezes when Finished and shows `-` before consent.
- Choices made where the architecture is silent are marked `// DECISION:` in the code.
