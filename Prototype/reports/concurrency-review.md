# Concurrency review: participant request path, Firestore layout, admin and export paths

Written 2026-10-24 by the performance and concurrency engineer. Goal set by the owner: many participants can run sessions at the same time without contention, and nothing in the request path blocks on work that can be done asynchronously.

Read: `02 - Architecture.md` 5.3 to 5.5, 6, 7.1, 8.1, 15; `services/api/README.md`; `services/api/internal/store/` (both stores); `internal/domain/` (engine, allocation, transaction interface); `internal/http/` (participant and admin handlers); `internal/export/`; `internal/sweep/`; `infra/terraform/modules/firestore` and `modules/cloud-run`.

Tools written for this review: `load/` (load generator, `make load-*`), `load/cmd/docsize` (document sizes in the emulator), `load/txbench` (isolated transaction benchmark). See `load/README.md`.

## 1. Summary

- The single cross-session write in the event path, `counters/allocation`, was a real hotspot. At 200 sessions started within 10 s on the emulator, consent transactions collapsed: in two runs 120 and 179 of 200 sessions never got past consent (all 5 transaction attempts timed out on the counter lock), while every other event type stayed at single-digit milliseconds. Fixed in `internal/store/firestore/firestore.go` with a per-process gate that admits one allocating transaction at a time (section 4). After the fix: 200 of 200 and 400 of 400 sessions finished, zero transaction retries, consent p95 1.6 ms.
- Everything else in the event path is per session: one session document, its own subcollections, no shared documents. Contention between sessions is otherwise structurally impossible.
- The admin and export paths read more than they need: none of the three composite indexes Terraform creates is used by any query as written, the participants poll with a run filter read every session of the run every 5 s (fixed, one `switch` reorder in the same file), and the export reads each session's subcollections sequentially, which will exceed the 30 s request deadline somewhere between a few hundred and 1,000 sessions in production.

## 2. Load test results

Setup: the API built from the working tree, run locally with throwaway admin secrets, `CONTENT_DIR=content`, `TRUSTED_PROXY_HOPS=1` (each simulated participant has its own address in `X-Forwarded-For`), `INSECURE_COOKIES=true`, `ALLOW_TEST_CONFIGS=true`. Config `test-fast-1.4.0`: the server supports versions starting with `test`, so the generator used it; it has no audio manifest, so the server serves dev silence clips whose duration is in the URL, and the generator waits exactly that long (section "Timing" of `load/README.md`). Think time multiplier 1 (ratings 1.5 to 5 s, actions 1 to 5 s, budget 3 to 10 s, background 5 to 15 s, server minimum view times always respected). Each session sends 46 events and 9 `GET /next`, about 90 s per session. Emulator: `gcloud emulators firestore` in `docker-compose.yml`, host port 8081, a fresh project id per run. One machine runs generator, API and emulator.

Latency in ms over all events (the `POST /events` calls), plus the slowest endpoint:

| Run | Sessions done | Events | p50 | p95 | p99 | Max | Statuses | Rejections | Wall |
|---|---|---:|---:|---:|---:|---:|---|---|---:|
| Memory store, 50 sessions, 5 s ramp | 50 / 50 | 2,300 | 0.3 | 0.5 | 0.6 | 1.2 | 200: 2,750, 201: 50 | too_early 0, seq_gap 0, wrong_state 0 | 105 s |
| Emulator, 50 sessions, 5 s ramp | 50 / 50 | 2,300 | 1.1 | 1.7 | 2.0 | 5.8 | 200: 2,750, 201: 50 | 0, 0, 0 | 104 s |
| Emulator, 200 sessions, 10 s ramp, **before the fix** | 21 / 200 | 2,229 | 16.8 | 30,001 | 30,001 | 30,002 | 200: 1,334, 201: 200, **503: 1,084** (all `consent_accepted`) | 0, 0, 0 | 271 s |
| Emulator, 200 sessions, 10 s ramp, after the fix | 200 / 200 | 9,200 | 0.9 | 1.5 | 1.8 | 10.4 | 200: 11,000, 201: 200 | 0, 0, 0 | 115 s |
| Emulator, 400 sessions, 10 s ramp, after the fix (extra) | 400 / 400 | 18,400 | 0.9 | 1.4 | 1.8 | 11.4 | all 2xx | 0, 0, 0 | 115 s |

`consent_accepted` alone: memory p95 0.6 ms; emulator 50: 1.7 ms; 200 before the fix: p50 25,165 ms, p95 30,001 ms (the server's 30 s request deadline), 1,105 requests of which 1,084 were 503 after the generator's own retries; 200 after the fix: p95 1.6 ms; 400: p95 1.6 ms. An earlier burst run (same setup, before the retry logging was built in) ended 80 of 200 with 755 consent 503s.

Admin check after every passing run (through the admin API, JSON export with invalid events): every session `Finished`; groups 25/25, 100/100 and 200/200; denial slides per group 9/8/8, 34/33/33 and 67/67/66 (difference 1); no lost or duplicated `clientSeq`, no `eventNo` gaps; events the generator saw accepted equal the valid consumed client events stored. The failed burst left 179 sessions `Opened` with no group (no slot taken, as 5.5 intends), and the 21 allocated sessions were still balanced (11/10), so the failure mode was availability, not correctness.

**Transaction retries observed** (counted by the store, logged as `firestore transaction retried`): 0 in every run after the fix, in the 50 session emulator run, and for every non-consent transaction in every run. Before the fix, 1,086 consent transactions each ran 5 attempts and failed (`Aborted: Transaction lock timeout` 852 times, then `DeadlineExceeded` 232 times once the 30 s request deadline cut the backoff short).

Section 15's pass criteria (p95 event latency under 300 ms, no lost or duplicated events, balanced groups and denial slides, no transaction errors beyond retries) are met locally at 50, 200 and 400 sessions after the fix. Local latencies are not production latencies (section 8).

## 3. Firestore layout as implemented

Measured with `load/cmd/docsize` after the 200 session run (sizes by Firestore's storage rules: name + fields + 32 bytes).

| Path | Docs per finished session | Avg bytes | Max bytes | Written when | Shared across sessions |
|---|---:|---:|---:|---|---|
| `runs/{runId}` | (1 per run) | 209 | | admin create or patch | read on every session creation (`code ==` query) |
| `counters/allocation` | (1 in total) | 229 | | each consent | **yes: written by every consent** |
| `configs/{version}` | (per upload) | | | admin upload | read at start |
| `sessions/{id}` | 1 | 1,911 | 1,929 | rewritten (whole document `Set`) by every event, every `GET /next`, a resume | no |
| `sessions/{id}/events/{eventNo}` | 46 | 918 | 1,860 | created by every event (accepted or rejected) and every synthetic timeout | no |
| `sessions/{id}/ratings/{point}` | 9 | 367 | 369 | rating submit or timeout | no |
| `sessions/{id}/answers/{kind}` | 3 | 379 | 473 | each action (the `actions` document is rewritten 3 times), budget, background | no |

Per finished session: 59 documents, about 52 KB stored (42 KB of it events). The largest session fields are `plan` (703 bytes, written once at consent but carried in every rewrite), `lastResult` (254, the stored response for retries), `userAgent`, `tokenHash`. The largest event field is `response` (490 bytes average, the full JSON answer kept for identical retries), more than half of each event.

**Per event, in the request path:** `participantAuth` reads the session outside the transaction (token check), then the transaction reads the session again, may read one rating or answers document (late submissions) or the consumed event (a retry of an older `clientSeq`, a query), and commits: the session document, one event document, and at most one rating or answers document. That is 2 reads and 2 to 3 writes. Consent adds 1 read and 1 write of `counters/allocation`. `GET /next` is 2 reads and 1 write (the session is rewritten even when nothing changed). Session creation: 1 query on `runs`, 1 query on `sessions` by `ip`, 1 create; plus one transaction per earlier session flagged as a duplicate.

**Shared documents:** only `counters/allocation` is written by more than one session. Two other cross-session writes exist off the main path: duplicate detection writes the earlier sessions with the same IP and user agent (in the creating request, one transaction each), and the sweeper writes sessions it marks Abandoned. Neither is a hot document.

**Collection write rates at the burst:** 9,200 events in 115 s, 80 per second on average, so about 80 session rewrites and 80 event creates a second. Indexed timestamps that increase monotonically (`updatedAt`, `lastEventAt` on sessions, `serverTs` on events) are subject to Firestore's guidance of about 500 writes a second per collection for sequential indexed values. Headroom about 6 times at the 200 session burst; the 400 session run reached 160 a second (headroom 3 times). At around 1,000 concurrent participants this would need index exemptions on the timestamps not queried (`lastEventAt` once the sweeper uses `status` with an in-memory filter, `serverTs`) or none would be left.

## 4. Contention analysis: `counters/allocation`

**Write rate at the burst.** The counter is written once per consent. Consent comes 3 to 8 s after session creation (consent minimum view 3 s plus think time). In the fixed 200 session run, 200 consents fell within 13.9 s: **14.4 writes a second sustained over the burst, 28 in the busiest second** (sliding 1 s window). The 400 session run: 28.9 a second, 54 in the busiest second. The 50 session run: 6.2 a second, 14 at peak.

**Against Firestore's guidance** of about one sustained write per second per document: the 200 burst is 14 times over it for about 14 s, so there is no headroom during such a burst; the guidance is a soft limit about sustained rates, and short bursts are absorbed with higher latency and more transaction conflicts. In terms of the real study: one write per second sustained is 3,600 consents an hour. The plan is 40 participants; even 200 sign-ups in the first hour after a post is 0.06 a second, about 18 times below the guidance, and a realistic link post spreads clicks over minutes. A 200 in 10 s burst is a stress case, not an expected load.

**What actually failed, and why.** The emulator implements transactions with pessimistic locks. A consent transaction reads the session and the counter, then writes both. When two consents overlap, both hold a read lock on the counter and each waits for the other to release it before it can write: a lock upgrade deadlock that the emulator resolves only by its lock timeout (about 3 s). Every consent arriving meanwhile also takes a read lock and joins the queue, the client library retries all of them in step, and the queue never drains while new consents keep arriving. `load/txbench` reproduces it without the API: 200 workers, 15 shared transactions a second, 10 ms between read and commit: 15 of 300 shared transactions committed; with 5 ms in between none failed. The API's consent transaction holds the counter for a few milliseconds (it compiles the plan and builds the response before commit), and under 200 concurrent sessions that window was long enough to tip it.

**The fix** (`internal/store/firestore/firestore.go`, `allocGate`, marked `DECISION`): a one-slot semaphore per Store (per process). `LoadCounters`, the only read of the counter, first acquires it; it is held across the client library's retries of that transaction and released when `RunTx` returns. Effect: within one process, consents never overlap on the counter; every other transaction is untouched. With the gate, `txbench` committed 300 of 300 shared transactions at 10 ms and 286 of 286 at 30 ms hold, and the load test passed at 200 and 400 sessions with zero retries. The acquire honours the request context, so a consent can never wait past its deadline.

**Trade-off of the fix.** Consents of one instance are serialised: one instance allocates at most one session per consent transaction round trip, about 25 to 50 a second in production (4 RPCs of 5 to 10 ms), a few hundred a second on the emulator. The 400 run needed 54 in its busiest second and saw consent p95 1.6 ms. Consents on different Cloud Run instances can still overlap (`max_instances` is 10, 80 requests per instance); the transaction keeps them correct, and the production database runs in **optimistic** mode (`concurrency_mode = "OPTIMISTIC"` in Terraform), where an overlap is an abort at commit and a retry with backoff, not a lock wait, so the collapse seen on the emulator cannot occur there in the same form. Group balance and the global round robin (1.19) are unchanged.

**Alternatives considered and not taken.** Per-run counters remove cross-run contention but break the global round robin decided in 1.19 (two runs could each start with checker). Sharded counters break the strict lowest-count rule. Moving allocation into its own transaction outside the consent transaction shortens the lock but makes allocation and consent non-atomic (a slot could be taken by a consent that then fails). An increment without a read is impossible because the choice depends on the counts. A per-process gate is the smallest change that keeps every documented guarantee.

## 5. Admin query paths and indexes

| Caller | Query as implemented | Index used | Reads per call |
|---|---|---|---|
| Session creation, run lookup | `runs where code ==` limit 1 | single field (automatic) | 1 |
| Session creation, duplicate check | `sessions where ip ==`; `openedAt` and user agent filtered in memory | single field `ip` | all sessions ever from that IP |
| `GET /admin/participants` (no filter) | all of `sessions` | none | N |
| same with `runId` | `sessions where runId ==` | single field `runId` | sessions of the run |
| same with `updatedSince` (the 5 s poll) | `sessions where updatedAt >` | single field `updatedAt` | sessions changed since the last poll |
| same with `status` | `sessions where status in [..]` | single field `status` | sessions with that status |
| `GET /admin/runs`, config list, config upload | all of `sessions` (to count per run, and to find locked versions) | none | N |
| Run patch (config change), run delete | `sessions where runId ==` (delete: limit 1) | single field | sessions of the run (delete: 1) |
| Participant detail | session; events ordered by `eventNo`; ratings; answers | single field (subcollection) | about 59 |
| Export | `sessions where runId ==` or all, then per session the three subcollections | single field | about 59 per session |
| Sweeper (every 5 min) | `sessions where status in [Opened, Started, Ongoing]`; `lastEventAt` filtered in memory | single field `status` | all unfinished sessions |
| Retry of an older `clientSeq` | `events where clientSeq == and consumed == true` limit 1, in the session's subcollection | two single-field indexes merged; no composite needed | 1 |

**Terraform's three composite indexes** (`sessions_run_opened` runId asc + openedAt desc, `sessions_status_lastevent` status asc + lastEventAt asc, `sessions_ip_opened` ip asc + openedAt desc) **are used by no query as written**: every query filters on a single field and does the rest in memory (a documented `DECISION` in `ListSessions`). They still cost an index entry update on every session rewrite whose `status`, `lastEventAt` or `openedAt` changes, which for `lastEventAt` is every event. Two of them would become useful with one line each in the queries: the sweeper could filter `status in [...]` and `lastEventAt <` in the query (index 2), and the duplicate check `ip ==` and `openedAt >=` ordered by `openedAt desc` (index 3). The first one matches no query unless the participants list orders by `openedAt` in Firestore rather than in memory. Recommendation: use indexes 2 and 3 in the queries (later, when sessions number in the thousands; at 40 the in-memory filters are cheaper than the index writes), and either drop index 1 or order the run query by it.

**Change made: the poll path.** `ListSessions` chose the indexed field in the order runId, ip, updatedSince, status. The admin page polls every 5 s with `updatedSince` and whatever filters are set; with a run filter set, every poll read every session of the run. Reordered so `updatedSince` wins (same file, marked `DECISION`); the run filter is then applied in memory, as the other filters already were. A poll now reads only the sessions changed in the last 5 s. With 200 sessions in a run and the admin tab open, that is the difference between 200 reads every 5 s (3.5 million a day) and a handful.

**Still reading every session:** `GET /admin/runs` and the config list. Each call is N reads. Recommendation in section 7 (aggregation count queries).

**Single-field index exemptions** (not in Terraform today): every field of every document is indexed automatically, including `sessions.plan` (an array of about 60 step keys), `sessions.lastResult.response`, `events.response`, `events.payload`, `events.bodyHash`, `sessions.tokenHash`, `sessions.userAgent`. None is queried. Exempting them reduces write latency and index storage for every event. Recommendation in section 7.

## 6. Export cost

The export (`export.Collect`) lists the sessions, then for each session reads its events, ratings and answers with three sequential queries, holds everything in memory, and writes CSV or JSON. CSV and JSON each re-read everything.

| Sessions | Document reads per export | Reads cost | JSON size | Emulator time (measured) | Production time (estimate) |
|---:|---:|---|---:|---:|---|
| 40 | about 2,400 | free tier (50,000 a day) | about 1.2 MB | under 0.1 s | 1 to 3 s |
| 200 | about 11,800 | free tier | 6.2 MB (measured) | 0.86 s | 6 to 15 s |
| 400 | about 23,600 | free tier | 12.5 MB (measured) | 4.0 s | 12 to 30 s |
| 1,000 | about 59,000 | above the daily free tier in one export; a few cents | about 31 MB | not run | 30 to 75 s: **exceeds the 30 s request context** and approaches Cloud Run's 60 s timeout |

Reads assume finished sessions (59 documents each; sessions with timeouts have a few more synthetic events). The production estimate is 3 sequential queries per session at 10 to 25 ms each from Cloud Run to Firestore in the same region; the emulator grows faster than linearly (0.12 s at 50, 0.86 s at 200, 4.0 s at 400) because its subcollection queries scan, which production queries do not. Memory: the 1,000 session export holds about 31 MB of JSON plus the decoded structures, several times that, against the 512 MiB instance.

Recommendation (later, before a run passes about 300 sessions): read the sessions' subcollections with bounded concurrency (8 to 16 at a time) or with three collection group queries (`events`, `ratings`, `answers`) streamed and joined by session id, and stream the CSV rows as they are built. That brings 1,000 sessions to a few seconds. For a very large study, an export job that writes to the bucket is the asynchronous form; not needed at 40.

## 7. What could be asynchronous, and recommendations

Cloud Run runs with `cpu_idle = true` (CPU only during requests), so "asynchronous" means another request, a Cloud Tasks task, or the existing sweeper job; a goroutine left running after the response is not reliable.

| Candidate | What it costs in the request path now | Can it move without weakening one event at a time? | Recommendation |
|---|---|---|---|
| Allocation (counter read and write at consent) | 1 read, 1 write on the one shared document | No. The plan and the next step descriptor depend on the group and denial slide, and the slot must be taken atomically with the consent. | **Never** async. **Done now:** per-process gate (section 4). |
| Flag computation (`many_invalid_events`, `tab_hidden_during_audio`, `reloaded`, `audio_problem`, `honeypot`, `audio_check_wrong`, `rate_limited`) | in-memory checks on the already loaded session; no I/O | Could, but there is nothing to gain: no extra reads or writes. | **Never.** |
| Duplicate detection at session creation | 1 query by IP, plus one transaction per earlier matching session to add `possible_duplicate`, sequential, inside the creation request | The query is cheap and the new session's own flag can stay. Flagging the earlier sessions writes other sessions' documents, can contend with their event transactions, and grows with shared networks (a campus NAT). It can move to the sweeper pass (it already runs every 5 minutes) without touching the event guarantee. | **Later.** At 40 participants from personal networks it is rare; move the "flag the earlier sessions" part to the sweeper, and put `openedAt >=` into the query to use index 3, when a large shared network is expected. |
| Denormalised admin fields (per-run status counts, finished per group) | none today; computed at read time from all sessions | Maintaining them in the event transaction would create a new shared document per run, written on every status change: a second hotspot. | **Never** in the request path. **Later:** use Firestore aggregation count queries per run and status at read time (one read per 1,000 index entries) instead of reading all sessions. |
| `GET /next` rewrites the session even when nothing changed | 1 write per call (9 per session) | Skipping an unchanged write keeps the guarantee. | **Later:** let the transaction skip the session write when lazy evaluation changed nothing (needs a dirty flag from the domain layer). |
| Token check reads the session outside the transaction, then the transaction reads it again | 1 extra read and one round trip per event (46 per session) | Checking the token hash inside the transaction is equivalent. | **Later** (in `internal/http`, not this review's files): halves the reads per event and saves one Firestore round trip. |
| Stored `response` on every event plus `lastResult` on the session | about 490 bytes per event, half of event storage | Needed for identical retries; the latest one is on the session, older ones only for retries of older numbers. | **Later:** keep `lastResult`, drop or truncate `response` on events after the next event. Not worth it at 40. |
| Single-field index exemptions (`plan`, `lastResult`, `tokenHash`, `userAgent`, `events.response`, `events.payload`, `events.bodyHash`) | index entry updates on every event | Pure storage setting. | **Do now** in Terraform (`google_firestore_field` with no indexes); cheap, lowers write latency, no code change. |
| Unused composite indexes | index entry updates on every event (`lastEventAt`) | Pure storage setting. | **Later:** use indexes 2 and 3 in the queries or drop them (section 5). |
| Export | sequential reads, all in memory | It is an admin action, not the participant path. | **Later:** bounded parallel or collection group reads and streaming (section 6). |
| Admin poll with a run filter | read every session of the run every 5 s | read path only | **Done now:** `updatedSince` first in `ListSessions`. |

The memory store serialises every operation of every session behind one mutex. It is for tests and `make dev-memory`, and even so served 50 sessions at p99 0.6 ms; not worth changing.

## 8. Emulator versus production

- **Locking.** The emulator locks pessimistically and resolves deadlocks by timeout; the production database is configured optimistic (Terraform), where reads take no locks and conflicting commits abort and retry. The collapse in section 4 is therefore the emulator's form of the problem. In production, without the gate, concurrent consents would abort each other at commit; at 14 a second and a transaction of 20 to 50 ms, a large share would retry and some would exhaust the 5 attempts. The gate removes the intra-instance part in both modes.
- **Latency.** Emulator RPCs are sub-millisecond on localhost; production Firestore from Cloud Run in the same region is about 5 to 20 ms per RPC, and an event is 4 RPCs (auth read, transaction begin and read, commit). Expect event p50 of 20 to 60 ms and p95 under 150 ms in production, comfortably within the 300 ms criterion, but not measured here.
- **Limits not enforced by the emulator:** the 1 write per second per document guidance, the 500 writes per second per collection guidance for sequential indexed values, missing composite indexes (the emulator never asks for one), quotas and billing.
- **Scale-out.** Locally there is one API process. Cloud Run may run up to 10 instances; the per-IP rate limiters and the allocation gate are per instance.
- **Single machine.** Generator, API and emulator shared one host; the numbers show the server's own cost, not network effects.

## 9. Changes made

All in `services/api/internal/store/firestore/firestore.go`, additive, behaviour of the interface unchanged; the store conformance test passes against the emulator (`FIRESTORE_EMULATOR_HOST=localhost:8081 go test ./internal/store/...`).

1. `allocGate`: per-process one-slot gate for transactions that read `counters/allocation` (section 4).
2. `TxStats` and a log line `firestore transaction retried` (session, attempts, whether it allocated) when a transaction needed more than one attempt, so retries are visible in Cloud Logging and in the load test.
3. `ListSessions`: `updatedSince` takes precedence as the indexed field (section 5).

New: `load/` (module `study/load`, one `use ./load` line in `go.work`), `load/README.md`, `load/Makefile` (`make load`, `load-memory`, `load-emulator`, `load-burst`, `docsize`), `load/run-local.sh`, `load/cmd/docsize`, `load/txbench` (its own module, outside `go.work`, because it needs the Firestore client).
