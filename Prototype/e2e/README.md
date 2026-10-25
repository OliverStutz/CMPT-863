# End-to-end tests

Playwright (headless Firefox only) against the real Go service and the Firestore emulator. One test,
`tests/flows.spec.ts`, does the whole round trip:

1. Logs in to the admin API (gate key, then password) and creates a run.
2. Runs participant sessions on that run until it has one checker and one neutral session (allocation
   is round robin, so usually two). Each session goes through the device gate, consent, the audio check
   (silent dev clip, correct word), the example, pre-start, the opening, the baseline rating, all five
   slides with their ratings, three action rounds, the questions and answers, budget, background and
   the debrief until the server says `done`. The checker session chooses "challenge", "explain" and
   "nothing"; the neutral session chooses "Yes", "No", "Yes". The test never decides the next screen:
   it reads the step descriptor of the last API response and acts on the scene it names. The group is
   read from the admin API, never from the participant side.
3. Leak check: every response body the participant page received (API, scene pack, HTML; not the JS
   and CSS bundle, which is the same for everyone) is scanned for `denial`, `checker`, `neutral`,
   `denialSlide` and `group`, as keys and in text, for the answer variant that was not played on each
   slide, and for the admin password. Allowed, and listed in `test-results/bodies/leak-report.txt`:
   the checker session's pre-start sentence, the "Challenge" option label, the panel title and the
   challenge question it chose to ask; the debrief paragraphs of both groups; the story wording
   "hospital group" in the opening.
4. Admin checks: both sessions are Finished with progress N/N, group and flags; `finishedAt` does not
   move between two polls; the detail has the event timeline; the CSV export has one row per rating
   point for both sessions, an `answerType` that matches what was played, and no answer text of a
   variant that was not played. Then the admin SPA is opened in Firefox and the participants table
   must show both sessions as finished.

Every body is saved under `test-results/bodies/` (git-ignored) for inspection, with `export.csv`.

`security.sh` checks the server from the outside with curl: participant calls without the cookie
(401) or without `X-Study-Client` (403), the admin gate (404 without it), and the asset allow list
(configs, schema, VLM drafts and path traversal are 404, slide SVG and PNG are 200).

## Run everything

From the repository root:

```sh
make e2e
```

This builds both SPAs, starts the Firestore emulator (`docker compose`, host port 8081), builds the API
and starts it on port 18080 (`E2E_PORT` to change it) with `CONTENT_DIR=content` and
`ALLOW_TEST_CONFIGS=true`, generates a throwaway admin password, gate key and session key for this run
only (they live in the script's environment and are never written to a file), runs the Playwright
test and then `security.sh`, and stops the API. The emulator keeps running; stop it with
`docker compose down`.

## Against a running API

```sh
cd e2e
npm ci
npx playwright install firefox
E2E_BASE_URL=http://localhost:8080 E2E_GATE_KEY=<gate key> E2E_ADMIN_PASSWORD=<password> ./run.sh
```

Or only the browser test: the same variables with `npx playwright test`. The API must run with
`ALLOW_TEST_CONFIGS=true` and `INSECURE_COOKIES=true` (plain http). The container image works the
same way, for example:

```sh
docker run --rm -p 8095:8080 --add-host host.docker.internal:host-gateway \
  -v "$PWD/content:/content:ro" -e CONTENT_DIR=/content \
  -e FIRESTORE_EMULATOR_HOST=host.docker.internal:8081 -e FIRESTORE_PROJECT_ID=demo-study \
  -e INSECURE_COOKIES=true -e ALLOW_TEST_CONFIGS=true \
  -e ADMIN_PASSWORD_HASH_VALUE -e ADMIN_ACCESS_KEY_VALUE -e ADMIN_SESSION_KEY_VALUE study-api:dev
```

## Fast and real timing

The clips are silent placeholders whose length the server estimates from the word count. The server
loads `content/study-config.<tag>.json` from `CONTENT_DIR` and accepts versions starting with `test`
only with `ALLOW_TEST_CONFIGS=true`; it has no other fast mode, the timing comes from the config. So
`content/study-config.test-fast.json` (version `test-fast-1.4.0`) is a copy of
`study-config.v1.4.json` with every text unchanged and short timings: 15 ms per word for the clips and
300 ms views after audio. With it both sessions take about a minute and a half. Its clips are silent
placeholders: no TTS set exists for the version, so the server serves `/assets/_dev/silence/<ms>.wav`
and the admin view marks the config "no audio: silent clips". For a walk-through with the spoken lines
use `content/study-config.test-sound.json` (version `test-sound-1.4.0`): the same texts, the real clips
(`content/audio/test-sound-1.4.0/`, built from the TTS cache at no cost) and 2 s views; one session takes
about seven minutes, almost all of it audio.

`E2E_CONFIG=1.4.0` runs the same test on the real 1.4 timing (10 s views and the real clip lengths
from `content/audio/1.4.0/manifest.json`, or about 400 ms per word without it); expect about 15
minutes per session. `E2E_CONFIG_FILE` names the config file whose texts the leak
check uses (default `content/study-config.v1.4.json`).

## Files

| File | What |
|---|---|
| `playwright.config.ts` | Firefox only, headless, one worker, autoplay allowed, volume 0 (`media.volume_scale` and `media.default_volume`: clips still decode and fire `ended`, nothing reaches the speakers), 1440 x 900 viewport |
| `tests/flows.spec.ts` | The flows, the leak check and the admin checks |
| `security.sh` | The outside-in curl checks |
| `run.sh` | Starts what is needed and runs both (used by `make e2e`) |
