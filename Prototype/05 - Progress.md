# Progress: prototype build and deployment

Single place for the state of the prototype. Updated whenever a component, the deployment or the content changes. Last updated 2026-10-24 (late evening).

## Status at a glance

| Area | State | Next step |
|---|---|---|
| Study design and content | v1.4 content with flags from the Opus VLM run; JSON validates against the schema | TTS audio, voice samples for Oliver, profile card and boardroom images |
| Architecture | Revised to all 53 decisions; change log in `02` section 18 | Keep in step with integration findings |
| Go backend (`services/api`) | Built; go test and go vet clean (domain, http, both stores, sweeper, config); eight security tests pass; live run against the Firestore emulator worked | Integration in progress; wording lint and GCS config upload still open |
| Participant app (`apps/participant`) | Built; typecheck, build and 16 tests pass; not yet run against the backend | Integration: reconcile step descriptor fields and scene pack shape |
| Admin app (`apps/admin`) | Built; typecheck, build and 11 tests pass | Integration: endpoint check against the backend, admin gate behaviour |
| Terraform (`infra/terraform`) | Bootstrap and dev applied, 86 resources; state in the bucket | Prod environment once a second project exists |
| Dockerfile and CI | Written; docker build is part of the integration pass now running | Push and rollout after integration, with Oliver's permission for the push |
| Cloud Run (dev) | Service live on Google's placeholder image | Deploy the real image after integration and security checks |
| Secrets | Three containers, version 1 each (password hash, access key, session key) | None; rotate by adding versions |
| Content in the bucket | Empty | Upload after the asset allowlist is verified |
| TTS tool, content lint | Being built against the backend's config package; voice samples and a first v1.4 audio set with recommended voices are part of the job | Oliver picks voices from the samples; regenerate if the pick differs |
| End-to-end tests, load test | Not started | After integration |
| Git | Prototype folder untracked, nothing committed | Commit and push after integration |

## Deployment facts

- Project `cmpt863-study` under the organisation, billing linked (CAD), region `northamerica-northeast1`.
- Terraform state: `gs://cmpt863-study-tfstate`, prefixes `study/bootstrap` and `study/dev`, versioned.
- Service URL: `https://study-api-pju4jz7m4a-nn.a.run.app` (placeholder image until the first real deploy).
- Buckets: `cmpt863-study-content` (private), `cmpt863-study-tts-cache` (private), `cmpt863-study-tfstate`.
- Admin plaintext credentials: only in `/home/oliver/.config/cmpt863-study/admin-credentials.txt`, owner-only, outside the repository. Secret Manager holds the bcrypt hash, the access key and the session key.
- CI identity: Workload Identity Federation for `OliverStutz/CMPT-863`; GitHub environment variables not yet set.

## Security rules in force (Oliver, 2026-10-24)

1. No credentials in the repository, ever; public exposure of config is tested, not assumed.
2. Only the run landing page is unauthenticated. Opening it creates the session and its token; every other participant request without that token is dropped.
3. The admin page renders only when a configured access key is provided in the URL once per browser (`/admin/?k=...`), never prompted; the password login applies on top. Without the key, admin paths answer 404.
4. The content bucket and the asset handler never expose the study config, the schema, the drafting files or anything that reveals the denial slide or the group.

## Still needed from Oliver

- A domain and who runs its DNS (optional for the pilot).
- The SFU consent PDF and ethics wording (data retention, raw IP, browser string).
- Voice choice from the samples once the TTS tool exists.
- Real images for the boardroom, chair, CISO and board member, or approval of drawn placeholders.

## Log

- 2026-10-23: prompt saved; flow brief; architecture and content v1 by two Opus agents; independent cross-check (66 findings); technical fixes applied; content v1.1 and v1.2; schema; decisions document; Oliver's first eight answers.
- 2026-10-24 (late): backend delivered with the token trail, admin gate and asset allowlist; integration pass started.
- 2026-10-24: all 53 decisions answered; content v1.3 (five-stage deck, neutral company) and v1.4 (flags from the VLM run, selection logged); slides rendered; Opus VLM drafting run archived with its fixed prompt; architecture revised to all decisions; builds started for backend, both apps and infrastructure; project created, bootstrap and dev environment applied; secrets set; security rules added to the backend scope.
