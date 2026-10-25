# Study API (Go)

One Go service that serves the participant SPA, the admin SPA, the public content assets and both APIs, as described in `02 - Architecture.md` (sections 2 to 8, 12.4, 14 and 15). Module path `study/api`; the repository root `go.work` lists this module.

## Layout

| Path | What |
|---|---|
| `cmd/api` | HTTP server |
| `cmd/sweeper` | Abandonment sweep (Cloud Run job, every 5 minutes) |
| `cmd/admintool` | `password` (bcrypt hash from stdin) and `key` (random key) for the admin secrets |
| `internal/domain` | Plan compiler (5.1), state machine and event validation (5.3), lazy timeouts (5.4), allocation (5.5), step descriptor (8.1) |
| `internal/events` | Event types, payloads, request envelope |
| `internal/store` | `Store` interface, transaction buffer; `memory/` and `firestore/` implementations; `storetest/` conformance test |
| `internal/configs` | Loads config versions and audio manifests, builds the public scene packs, handles admin uploads |
| `internal/assets` | Content source (local directory or GCS) and the public asset handler |
| `internal/http` | Routes, participant and admin authentication, admin API, SPA serving |
| `internal/export` | CSV and JSON export (8.4) |
| `internal/sweep` | Sweep logic used by `cmd/sweeper` |
| `internal/secrets` | Secret Manager reader with the 5 minute refresh |
| `pkg/studyconfig` | Go types of the config schema, loader with JSON Schema validation, audio manifest type, scene pack builder |
| `web/participant`, `web/admin` | Built SPAs, embedded with `go:embed` (filled by `make build-web`) |

## Running locally

From the repository root:

```sh
make dev          # Firestore emulator in Docker (host port 8081), then the API on :8080
make dev-memory   # same API with the in-memory store, no Docker
make build-web    # build both SPAs into services/api/web/ (then restart the API)
```

`make dev` seeds a run with code `devrun` on the newest loaded config, so the study is at `http://localhost:8080/s/devrun`. The admin area is at `http://localhost:8080/admin/?k=dev-gate-key`, password `dev`.

Without an audio manifest (no TTS run yet) the server estimates each clip's duration from its word count (`timing.msPerWordEstimate`, or 400 ms per word) and points audio steps at a silent WAV of that length, `/assets/_dev/silence/<ms>.wav`, so the whole flow runs. A warning is logged per config version.

Other targets: `make test` (all tests; the Firestore store test is skipped without the emulator), `make test-emulator` (starts the emulator and runs everything), `make build` (binaries in `services/api/bin`), `make lint` (`go vet`), `make sweeper` (one sweep against the emulator), `make admin-password`, `make admin-key`.

## Environment variables

| Variable | Default | Meaning |
|---|---|---|
| `PORT` | `8080` | Listen port |
| `CONTENT_DIR` | none | Content root in dev (`../../content` from `services/api`) |
| `CONTENT_BUCKET` | none | GCS content bucket in prod; wins over `CONTENT_DIR` |
| `STORE` | auto | `memory` or `firestore`; auto picks Firestore when an emulator host or a project id is set |
| `FIRESTORE_EMULATOR_HOST` | none | e.g. `localhost:8081` |
| `FIRESTORE_PROJECT_ID`, `FIRESTORE_PROJECT`, `GOOGLE_CLOUD_PROJECT` | `demo-study` with the emulator | Firestore project (first one set wins); database `(default)` |
| `TRUSTED_PROXY_HOPS` | `0` | Client IP from `X-Forwarded-For`: the n-th entry from the right (1 on Cloud Run); 0 uses the TCP peer. The raw header is logged for the first 20 session creations per instance |
| `ALLOW_TEST_CONFIGS` | `false` | Accept configs whose version starts with `test` |
| `RUN_CODE_LENGTH` | `12` | Length of the link code of a new run (6 to 32; alphabet without look-alikes, 31 characters). Existing runs keep their codes |
| `INSECURE_COOKIES` | `false` | Drop the `Secure` cookie flag (local http only) |
| `PUBLIC_BASE_URL` | request host | Base of run links, e.g. `https://study.example.org` |
| `SEED_RUN_CODE` | none | Create a dev run with this code at start |
| `ADMIN_PASSWORD_HASH_SECRET`, `ADMIN_SESSION_KEY_SECRET`, `ADMIN_ACCESS_KEY_SECRET` | none | Secret Manager resource names, see Security |
| `ADMIN_PASSWORD_HASH_VALUE`, `ADMIN_SESSION_KEY_VALUE`, `ADMIN_ACCESS_KEY_VALUE` | none | Plain values for local dev; override the `_SECRET` variants |
| `ADMIN_PASSWORD` | none | Dev only: a plain password, hashed at start when no hash is configured |
| `ABANDON_AFTER_MS` | `1800000` | Sweeper fallback when a session's config cannot be loaded |
| `SWEEP_INTERVAL` | none | Sweeper: repeat every interval (e.g. `1m`) instead of running once |

## Routes

| Route | What |
|---|---|
| `GET /`, `GET /s/{code}` and any other path | Participant SPA (`index.html` fallback; existing bundle files, under `/static/`, served as they are; a missing path with a file extension is `404`, not the SPA page) |
| `GET /admin/...` | Admin SPA, behind the gate key |
| `GET /assets/{path}` | Participant bundle files first, then the scene packs `scenes/<version>.json`, the dev silence clips, and the public content prefixes (see Security). `/a/{path}` is the same content handler under the document's path |
| `GET /api/v1/healthz` | `{ "ok": true }` |
| `POST /api/v1/runs/{code}/sessions` | Create or resume a session (8.1) |
| `POST /api/v1/sessions/{id}/events` | One event (5.3) |
| `GET /api/v1/sessions/{id}/next` | Current step after lazy evaluation |
| `POST /api/admin/login`, `POST /api/admin/logout` | Admin session |
| `GET, POST /api/admin/runs`; `PATCH, DELETE /api/admin/runs/{id}` | Runs |
| `GET /api/admin/participants`, `GET /api/admin/participants/{id}` | List with filters `runId, status, group, flagged, updatedSince, limit, cursor`; detail with events, ratings, answers |
| `GET /api/admin/export?format=csv\|json&runId=&includeInvalid=&includeIp=` | Export (8.4) |
| `GET, POST /api/admin/configs`; `GET /api/admin/configs/{version}` | Config list, upload, full config |

Request and response shapes follow section 8 of the architecture. Additions, each marked `DECISION` in the code:

- The step descriptor carries `prompt` and `note` on action steps (the per-group wording cannot be in the public scene pack) and `blocked: true` with `scene: "done"` for a session the device gate blocked.
- `assetBaseUrl` is `/assets/images/`; the consent PDF is served from `/assets/consent/<file>.pdf`.
- `audio_problem` responses carry `continueInMs` next to the usual fields.
- `progress.total` is `null` before allocation.
- A checker overlay carries the panel texts `title` (`labels.panelTitle`), `idleText` (`labels.panelIdle`) and, on `noIssues`, `noIssuesColor`; the public scene pack no longer has `labels.panelTitle`, `labels.panelIdle`, `labels.noIssues` or `colors`, because they name the checker in a neutral session's pack.
- `debrief.sessionCodeLabel` carries the config's label for the session code.
- Background item values: `null`, `"preferNotToSay"`, an option id (with the "Other" text as `{ optionId, text }` or as a separate key `<itemId>.otherText`), an integer for a scale, a string for a country.
- Missing `X-Study-Client` or `X-Study-Admin`, and a token for another session, answer `403` with code `forbidden`. Unknown admin records answer `404 not_found`; conflicts `409 conflict`.

## Security

**Participant authentication.** The only unauthenticated participant entries are the run landing page (`GET /s/{code}`, the SPA) and `POST /api/v1/runs/{code}/sessions`, which creates the session and issues its token. The token (256 bits) is sent once, in the cookie `study_session` (`HttpOnly`, `Secure`, `SameSite=Strict`, path `/api/v1`), as `<sessionId>.<token>`; only its SHA-256 is stored. Every other request under `/api/v1/` (healthz aside) needs that cookie and `X-Study-Client: 1`:

- no cookie, an unknown session or a wrong token: `401`, nothing else touched, counted per IP (past 120 an hour: `429`);
- a valid token used on another session's path: `403`;
- a missing `X-Study-Client` header: `403` (also on session creation).

Session creation is rate limited per IP: after 5 an hour new sessions are flagged `rate_limited`, past 30 an hour `429`. Events are limited to 60 a minute and 600 in total per session.

**Admin gate key.** The admin SPA under `/admin/` and every `/api/admin/*` endpoint, login included, answer `404` unless the request carries the gate key: as `?k=<key>` on an `/admin/` page request, which sets the signed `admin_gate` cookie (`HttpOnly`, `Secure`, `SameSite=Strict`, path `/`, 12 hours) and redirects to `/admin/` without the query string, or as that cookie. The key is compared in constant time; the cookie is bound to a fingerprint of the current key, so rotating the key invalidates it. With no key configured the admin area is unreachable. Only behind the gate does the password login apply: `POST /api/admin/login { password }` (bcrypt, 5 attempts per 15 minutes per IP) sets the `admin` cookie (path `/api/admin`, 12 hours, HMAC-SHA256 signed), and every other admin call also needs `X-Study-Admin: 1`.

**Secrets.** Three secrets, each delivered the same way:

| Secret | `_SECRET` variable (prod) | `_VALUE` variable (dev) |
|---|---|---|
| bcrypt hash of the admin password | `ADMIN_PASSWORD_HASH_SECRET` | `ADMIN_PASSWORD_HASH_VALUE` |
| HMAC key for the admin and gate cookies | `ADMIN_SESSION_KEY_SECRET` | `ADMIN_SESSION_KEY_VALUE` |
| admin gate key | `ADMIN_ACCESS_KEY_SECRET` | `ADMIN_ACCESS_KEY_VALUE` |

A `_SECRET` variable holds a Secret Manager resource name, `projects/<id>/secrets/<name>`; the service reads `versions/latest` at start and again every 5 minutes (a failed refresh keeps the previous value), so a new secret version takes effect without a redeploy. A `_VALUE` variable overrides it with a plain value for local development. Without a session key a random per-process key is used (cookies then end with the process). Create the values with:

```sh
make admin-password   # reads the password from stdin, prints the bcrypt hash
make admin-key        # prints a random 32-byte base64url key (gate key or session key)
```

and add them as new secret versions, for example `make admin-key | gcloud secrets versions add admin-access-key --data-file=-`. The password is never passed as an argument.

**Public content.** The asset handler serves only these paths from the content directory or bucket: `audio/<version>/*.mp3`, `images/*` (image types), `slides/**/*.svg`, `slides/**/*.png` and `consent/*.pdf`. Everything else (the configs, the schema, `vlm-drafts/`, the audio manifests, notes and scripts) is `404` even when the file exists, and any path with `..`, `.`, `//` or a backslash is rejected before routing. Content-hashed files are `immutable`; other files and the scene packs are `no-cache`.

**What participants never receive.** The public scene pack is built from the config without actions, checker panels, routine and participant questions, answers, the checker pre-start sentence, the debrief, the voices and the audio check's clip and answer. Step descriptors never carry the group, the denial slide, the answer type, admin fields or the other group's texts; a checker overlay arrives only with its slide step; question and answer steps carry the one clip the server chose. The test `TestParticipantResponsesLeakNothing` walks a checker session over HTTP with `content/study-config.v1.4.json` and checks every body.

## Adding a config version

1. Write `content/study-config.<tag>.json` (or `content/configs/<version>.json`) against `content/schema/study-config.schema.json`. The `version` field is the version runs pin.
2. Run TTS so that `content/audio/<version>/manifest.json` exists (the server also looks in `audio/<tag>/`). Until then durations are estimated.
3. Restart the API (or publish to the bucket and start a new revision). At start the server validates every config against the schema, compiles the plan for both groups, resolves clips and marking times, and logs any config it skips. Configs whose version starts with `test` load only with `ALLOW_TEST_CONFIGS=true`.
4. Alternatively upload the JSON in the admin Configs tab (`POST /api/admin/configs`); it is validated, stored as a draft in Firestore and usable at once. A version used by a run with sessions is locked and cannot be replaced.

### Audio manifest

`audio/<version>/manifest.json`, written by `tools/tts`:

```json
{
  "configVersion": "1.4.0",
  "clips": {
    "s1.narration": { "audioKey": "<sha256>.mp3", "durationMs": 41250, "marks": { "m1": 16500 } },
    "opening": { "audioKey": "<sha256>.mp3", "durationMs": 52000, "lines": ["opening.chair.1", "..."], "lineStarts": { "opening.chair.2": 9100 } }
  }
}
```

One entry per clip, that is per audio step. A single-line clip uses the line id. The opening uses the opening scene id, the setup example `setup.exampleScene`, and joined answers `<slideId>.answers.reason` and `<slideId>.answers.denial`; whole answers use their two line ids. Audio files are then served from `/assets/audio/<version>/<audioKey>`.

## Tests

`go test ./...` covers the plan compiler (both groups, the five-slide default, four slides with two flagged), the state machine (both groups end to end, `too_early` and the resend, `seq_gap`, identical and mismatched retries, the audio check, a late rating, a lazy timeout chain, the reload restart), allocation (100 consenting sessions, rotation order), the sweeper, the store conformance test (memory always; Firestore with `FIRESTORE_EMULATOR_HOST`, including concurrent events on one session), schema round trips of the real configs, and the HTTP layer (session creation, consent, a slide and a rating; the security tests above; the admin gate; runs, participants, export and configs; the asset allow list).
