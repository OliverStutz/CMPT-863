# Progress: prototype build and deployment

Single place for the state of the prototype. Updated whenever a component, the deployment or the content changes. Last updated 2026-10-24, after revision 6 (consent layout fix) passed the live walk.

## Status at a glance

| Area | State | Next step |
|---|---|---|
| Study design and content | v1.4 content with flags from the Opus VLM run; JSON validates against the schema | TTS audio, voice samples for Oliver, profile card and boardroom images |
| Architecture | Revised to all 53 decisions; change log in `02` section 18 | Keep in step with integration findings |
| Go backend (`services/api`) | Built and integrated; go test and go vet clean; security tests pass; deployed | Wording lint lives in `tools/content-lint`; admin config upload to GCS still open |
| Participant app (`apps/participant`) | Built, integrated and deployed; 17 unit tests; Playwright flows pass locally and live; marks and panel persist through post-view, action, question and answer; action screen inside the boardroom; slide 3 diagram asset in place | Real images for boardroom, chair, CISO and member (drawn placeholders now) |
| Admin app (`apps/admin`) | Built, integrated and deployed behind the URL gate and password; 11 unit tests | Config upload tab is read-only for now |
| Terraform (`infra/terraform`) | Bootstrap and dev applied, 86 resources; state in the bucket | Prod environment once a second project exists |
| Dockerfile and CI | Image builds and runs; pushed as `api:d259725` (digest in `envs/dev/terraform.tfvars`); GitHub workflows written but GitHub environment variables not yet set | Set the GitHub variables listed in `infra/README.md` to enable deploys from main |
| Cloud Run (dev) | Revision `study-api-00006-68s` live on image `api:62a9155` (consent layout fix) with `ALLOW_TEST_CONFIGS=true`; configs 1.3.0, 1.4.0 and test-fast-1.4.0 loaded from the bucket; live security checks pass; live Playwright flows pass (checker and neutral) | Pilot with 3 to 5 people on run `kyfxfg` |
| Secrets | Three containers, version 1 each (password hash, access key, session key) | None; rotate by adding versions |
| Content in the bucket | Mirrored from `content/` (configs, schema, slides, audio 1.4.0 with manifest, samples, VLM drafts); only the allowlisted paths are reachable through the service | Re-mirror after any content change: `gcloud storage rsync content gs://cmpt863-study-content --recursive --exclude='.*/\.cache/.*|.*\.py$'` |
| TTS tool, content lint | Built; 18 and 11 tests pass; lint on v1.4: 0 failures. Audio set `content/audio/1.4.0/` (23 clips, 6 min 53 s, about USD 0.11) with voices CISO en-US-Neural2-J, chair en-US-Neural2-F, member en-US-Neural2-E (picked without listening). Samples in `content/audio/samples/`. Denials padded to their reasons on all three slides (s2 15.6 s, s3 17.3 s, s4 17.0 s); lint with manifest: 0 failures, 0 warnings; audio set 6 min 58 s | Oliver listens to the samples and confirms or changes the voices; `tools/tts` regenerates from cache |
| End-to-end tests | `e2e/` Playwright (Firefox, muted) and `e2e/security.sh` (23 checks); pass locally on the fast and the real timing, and live on the fast config | Run live on the real timing before the pilot (about 15 min) |
| Load test and concurrency review | Done. `load/` generator; results: 50 sessions clean on memory and emulator; a 200-session burst failed 1,084 consents on the allocation counter, fixed by a per-instance gate in the Firestore store; 200 and 400 bursts then clean, event p99 under 2 ms on the emulator, groups and denial slides balanced. `reports/concurrency-review.md` has the layout, indexes and async recommendations | Later: parallel streamed export before 1,000 sessions; use or drop the three composite indexes |
| Git | Snapshots on origin main: 6bb9d0c, d259725, cc97907, 753f193, 288b2df, 62f6fb1, 62a9155 (consent layout fix) | Keep snapshotting after each change |

## Deployment facts

- Project `cmpt863-study` under the organisation, billing linked (CAD), region `northamerica-northeast1`.
- Terraform state: `gs://cmpt863-study-tfstate`, prefixes `study/bootstrap` and `study/dev`, versioned.
- Service URL: `https://study-api-pju4jz7m4a-nn.a.run.app`, revision `study-api-00006-68s`, image `api@sha256:86857ed7...` from commit 62a9155.
- Runs created 2026-10-24: `kyfxfg` (name "Pilot 1.4.0", config 1.4.0, real timing, link `https://study-api-pju4jz7m4a-nn.a.run.app/s/kyfxfg`) and `tvfe8g` ("Fast walk-through", test-fast-1.4.0, link `https://study-api-pju4jz7m4a-nn.a.run.app/s/tvfe8g`). Admin: `https://study-api-pju4jz7m4a-nn.a.run.app/admin/?k=<access key>` once per browser, then the password; both in the credentials file.
- Buckets: `cmpt863-study-content` (private), `cmpt863-study-tts-cache` (private), `cmpt863-study-tfstate`.
- Admin plaintext credentials: only in `/home/oliver/.config/cmpt863-study/admin-credentials.txt`, owner-only, outside the repository. Secret Manager holds the bcrypt hash, the access key and the session key.
- CI identity: Workload Identity Federation for `OliverStutz/CMPT-863`; GitHub environment variables not yet set.

## Parallelism rule (Oliver, 2026-10-24)

Many participants must run at the same time without contention, and the build runs in parallel streams. Layout: one document per session, an append-only event subcollection per session, the plan compiled server-side; the only cross-session write is the round-robin allocation counter, which strict alternation requires. The load test proves the headroom; the review names what to make asynchronous.

## Security rules in force (Oliver, 2026-10-24)

1. No credentials in the repository, ever; public exposure of config is tested, not assumed.
2. Only the run landing page is unauthenticated. Opening it creates the session and its token; every other participant request without that token is dropped.
3. The admin page renders only when a configured access key is provided in the URL once per browser (`/admin/?k=...`), never prompted; the password login applies on top. Without the key, admin paths answer 404.
4. The content bucket and the asset handler never expose the study config, the schema, the drafting files or anything that reveals the denial slide or the group.

## Still needed from Oliver

- A domain and who runs its DNS (optional for the pilot).
- The SFU consent PDF and ethics wording (data retention, raw IP, browser string). A one-page placeholder is served as `consent/consent-placeholder.pdf` until then.
- Voice choice from the samples once the TTS tool exists.
- Real images for the boardroom, chair, CISO and board member, or approval of drawn placeholders.

## Resume from here (if the session crashes)

Running when this was written: the integration agent (edits under `services/`, `apps/`, `packages/`, `e2e/`, `Dockerfile`, `Makefile`) and the load-test agent (edits under `load/`, `reports/concurrency-review.md`, possibly minimal additive changes under `services/api/internal/store/`). The audio agent is finished. If either stopped mid-way, run the checks below first and finish by hand.

1. Verify the pieces: `cd services/api && go test ./... && go vet ./...`; `cd apps/participant && npm ci && npm run typecheck && npm run build && npm test`; same in `apps/admin`; `cd tools/tts && go test ./...`; `cd tools/content-lint && go test ./... && go run . -config ../../content/study-config.v1.4.json -manifest ../../content/audio/1.4.0/manifest.json`.
2. Run locally: `make dev` (Firestore emulator on 8081, API on 8080 with `CONTENT_DIR=content`), then open `http://localhost:8080/admin/?k=dev-gate-key`, password `dev`, create a run on config `1.4.0`, open its link.
3. Build and push the image (Oliver allowed this on 2026-10-24; done once for d259725, repeat for a new commit):
   `gcloud auth configure-docker northamerica-northeast1-docker.pkg.dev --quiet`
   `docker build -t northamerica-northeast1-docker.pkg.dev/cmpt863-study/study/api:$(git rev-parse --short HEAD) .`
   `docker push northamerica-northeast1-docker.pkg.dev/cmpt863-study/study/api:$(git rev-parse --short HEAD)` and note the digest it prints.
4. Roll out: `cd infra/terraform/envs/dev && terraform apply -input=false -auto-approve -var="image_digest=northamerica-northeast1-docker.pkg.dev/cmpt863-study/study/api@sha256:<digest>"`.
5. Upload content (Oliver allowed this): `gcloud storage rsync content gs://cmpt863-study-content --recursive --exclude='.*\.cache.*'` (the service's allowlist decides what is reachable; the bucket itself is private).
6. Verify live: `curl -i https://study-api-pju4jz7m4a-nn.a.run.app/api/v1/healthz`; `/admin/` must be 404 without the key; `/assets/study-config.v1.4.json` must be 404; `/assets/slides/v1.3/s1.png` must be 200; admin key and password are in `/home/oliver/.config/cmpt863-study/admin-credentials.txt`.
7. Commit and push a snapshot (exclude `.obsidian/workspace.json`), then update this file.

## Log

- 2026-10-23: prompt saved; flow brief; architecture and content v1 by two Opus agents; independent cross-check (66 findings); technical fixes applied; content v1.1 and v1.2; schema; decisions document; Oliver's first eight answers.
- 2026-10-24 (consent): the real PDF caused a layout shift that moved the checkbox under the frame; fixed height for the PDF area; image 62a9155 rolled out as revision 6; live walk passes. Note: headless Firefox does not render PDFs, so the viewer must be checked in a real browser.
- 2026-10-24 (fixes): diagram asset, persistent marks, action screen in the boardroom, wrapped flag chips, consent placeholder; image 62f6fb1 rolled out as revision 5; content re-mirrored.
- 2026-10-24 (screens): screenshot walk of both groups and the admin (104 frames, contact sheet, index in `reports/screenshots/2026-10-24/README.md`). Found: slide 3 network diagram asset missing in the deployed content; checker marks vanish after the narration instead of staying through the viewing pause, action, question and answer; admin Flags column clipped at 1366 px. Fix round started.
- 2026-10-24 (checks): `e2e/security.sh` run live: 23 of 23 pass; the two traversal probes are answered by Google's front end with a redirect to the cleaned path, which the service then answers 404 (request logs confirm the uncleaned path never reaches the service); the script accepts that form. Snapshots dc4165f and the following commit.
- 2026-10-24 (load): 200-session burst exposed the allocation counter; per-instance gate added; image 753f193 rolled out as revision 4; live flows and the 23 security checks re-run.
- 2026-10-24 (deploy): image d259725 pushed and rolled out; content mirrored to the bucket; first revision had started before the upload and held no configs, fixed by wiring `extra_env` (ALLOW_TEST_CONFIGS=true on dev) through Terraform, which rolled revision 3; live security checks and the live Playwright flows pass; runs kyfxfg and tvfe8g created.
- 2026-10-24 (later): denial padding enabled on s2 and s4, audio regenerated from cache, lint clean. Audio pipeline delivered (tools/tts, tools/content-lint, voice samples, v1.4.0 audio set). Oliver allowed image push, apply with digest and content upload.
- 2026-10-24 (late): first snapshot pushed to origin main (6bb9d0c). Backend delivered with the token trail, admin gate and asset allowlist; integration pass started.
- 2026-10-24: all 53 decisions answered; content v1.3 (five-stage deck, neutral company) and v1.4 (flags from the VLM run, selection logged); slides rendered; Opus VLM drafting run archived with its fixed prompt; architecture revised to all decisions; builds started for backend, both apps and infrastructure; project created, bootstrap and dev environment applied; secrets set; security rules added to the backend scope.
