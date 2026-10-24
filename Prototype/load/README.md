# Load generator (`study/load`)

Drives N concurrent participant sessions through the complete study flow using only the public participant API (`02 - Architecture.md` 8.1), then checks the outcome through the admin API. Results and analysis: `reports/concurrency-review.md`.

## What one simulated participant does

Create the session with the run code (own cookie jar, own client address in `X-Forwarded-For`), then follow the server's step descriptor until `scene: "done"`:

| Step | Events and timing |
|---|---|
| deviceGate | `device_checked` 1440 x 900 |
| consent | waits at least 3.1 s (server minimum 3 s), `consent_accepted` with the scene pack's `consent.version` |
| audioCheck | `audio_check_played`, waits the clip, `audio_check_answered` with a random option |
| example, opening, slide, question, answer | `scene_started`, waits the clip duration (checker slides: `marking_shown` at each marking's `appearAtMs` within the clip), `audio_ended` |
| preStart | waits at least 2.1 s, `setup_done` |
| postView | waits `opensInMs`, then `GET /next` (as the participant app does) |
| rating | `rating_submitted` with random values within each item's scale |
| action | `action_chosen` with a random option of the descriptor (`slideId` from the scene pack's `forSlide`) |
| budget, background | random option and confidence; random background answers, about 10 % "preferNotToSay", empty honeypot |
| debrief | waits at least 5.1 s, `finished` |

**Timing.** The generator uses the test-fast config (`test-fast-1.4.0`, accepted when the server runs with `ALLOW_TEST_CONFIGS=true`) when it creates the run itself. Clip durations come from the step descriptor's `audioUrl`: a dev silence clip carries its length in the URL (`/assets/_dev/silence/<ms>.wav`); for a real MP3 the generator downloads it once and reads the duration from the frame header (Xing frame count or constant bitrate). Post-audio views use the descriptor's `opensInMs`, windows its `closesInMs` (think time is capped 1.5 s before a window closes, so no timeouts are produced). If a duration were wrong, `too_early` with `retryInMs` is honoured and the event resent with the next `clientSeq`, and the rejection is counted. Think times are random within fixed ranges times `-think`; the server's minimum view times always apply.

Every response's `expectedClientSeq` and `next` are taken as the new state. A 5xx or network error is retried with the identical body (a retry for the server), at most 5 times with backoff.

## Flags

| Flag | Default | Meaning |
|---|---|---|
| `-base` | `http://localhost:8080` | Server |
| `-run-code` | empty | Run to use; empty creates one through the admin API |
| `-config` | `test-fast-1.4.0` | Config version of a created run |
| `-sessions` | 50 | Concurrent sessions |
| `-ramp` | 5 | Seconds over which the sessions start, evenly spaced |
| `-think` | 1 | Think time multiplier (0 = as fast as the server allows) |
| `-check` | true | Admin check at the end |
| `-ip-base` | 1 | Varies the fake client addresses between runs |
| `-seed` | time | Random seed |

Environment for the admin check (and run creation): `LOAD_GATE_KEY` (the admin gate key) and `LOAD_ADMIN_PASSWORD`.

**Server settings the generator needs:** `TRUSTED_PROXY_HOPS=1`, so each session's `X-Forwarded-For` address is its client IP (otherwise all sessions share 127.0.0.1 and the per-IP limit of 30 sessions an hour returns 429, and every session after the first is flagged `possible_duplicate`); `INSECURE_COOKIES=true` over plain http; `ALLOW_TEST_CONFIGS=true` for the fast config.

## Output

A table per endpoint (`POST sessions`, `POST events <type>`, `GET next`) with request count, p50, p95, p99 and max latency and the HTTP statuses; totals over all events and all requests; server rejections (`too_early`, `seq_gap`, `wrong_state`, others); network errors and repeated requests; completed sessions and wall time. The admin check then reads `GET /api/admin/runs`, `GET /api/admin/participants?runId=`, and both exports (timed), and from the JSON export (with invalid events) checks that

- every session of this load run is `Finished`;
- group counts differ by at most one;
- within each group, denial slide counts differ by at most one;
- each session's consumed `clientSeq` values are exactly 1..`lastClientSeq` once each, its `eventNo` values exactly 1..`lastEventNo`, and the server's valid consumed client events equal the events the generator saw accepted (no lost or duplicated events).

It prints `RESULT: PASS` or `RESULT: FAIL` (exit status 1).

## Running

```sh
cd load
make load-memory                 # own API on the memory store, 50 sessions, 5 s ramp
make load-emulator               # own API on the Firestore emulator (docker compose up firestore; host port 8081)
make load-burst                  # 200 sessions started within 10 s on the emulator
make load BASE=http://localhost:8080 SESSIONS=50   # against a server you started (LOAD_GATE_KEY, LOAD_ADMIN_PASSWORD)
```

`run-local.sh` (used by the `load-*` targets) builds the API, makes throwaway admin secrets in a temporary directory (never in a tracked file), starts the API on port 18080 (`LOAD_PORT`) with `CONTENT_DIR=content`, `TRUSTED_PROXY_HOPS=1`, `INSECURE_COOKIES=true`, `ALLOW_TEST_CONFIGS=true`, and for the emulator a fresh project id (`load-<unix time>`), so counters and data start empty and do not mix with `make dev` data. After the run it prints the number of Firestore transaction retries the server logged and, for the emulator, the document sizes (`cmd/docsize`).

## Tools

- `cmd/docsize`: document counts and sizes per collection in the emulator, by Firestore's storage size rules, through the emulator's REST API (`go run ./cmd/docsize -project <id>`).
- `txbench/` (own module, not in `go.work`, because it needs the Firestore client): transactions on per-worker documents plus a share on one shared document, to isolate the allocation counter's behaviour. `cd txbench && GOWORK=off FIRESTORE_EMULATOR_HOST=localhost:8081 go run . -workers 200 -rate 0.5 -shared-rate 15 -hold 5ms`.
