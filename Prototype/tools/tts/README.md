# tools/tts

Synthesises every spoken line of a study config with Google Cloud Text-to-Speech and writes the audio set the server plays (section 11.1 of `02 - Architecture.md`):

- `content/audio/<version>/<sha256>.mp3`, one file per clip, named by a content hash so a file name never reveals the line or condition it holds;
- `content/audio/<version>/manifest.json`, in the shape of `studyconfig.Manifest`: clip id to `audioKey`, `durationMs`, `marks` (SSML mark name to offset in ms) and, for joined clips, `lines` and `lineStarts`.

`<version>` is the config's `version` field (for v1.4: `1.4.0`), the first directory the server looks in.

## Usage

From `Prototype/` (the module is in `go.work`):

```
go run ./tools/tts -config content/study-config.v1.4.json -dry-run   # lines, characters, cost; no API call
go run ./tools/tts -config content/study-config.v1.4.json            # synthesise and write the set
go run ./tools/tts -config content/study-config.v1.4.json -only s3.answer.part2
go run ./tools/tts -config content/study-config.v1.4.json -samples tools/tts/voice-candidates.json
```

Needs application-default credentials (`gcloud auth application-default login`) with a quota project on a project where the Text-to-Speech API is enabled. The quota project comes from the credentials file; `-quota-project` overrides it.

## Flags

| Flag | Meaning |
|---|---|
| `-config` | Study config file (required). |
| `-out` | Output directory. Default `<config dir>/audio/<version>`. With `-samples`: default `<config dir>/audio/samples`. |
| `-cache` | Local cache directory. Default `<config dir>/audio/.cache` (git-ignored). |
| `-voices` | A voices file (same shape as the config's `voices`) that overrides the config for every role it names. |
| `-voices-default` | Voices used where the config's voice name is a placeholder such as `TBD`. Default `tools/tts/voices.default.json`. |
| `-dry-run` | Prints every clip, its lines, character counts, whether each line is cached, and the estimated cost. Calls nothing and writes nothing. |
| `-only <id>` | Builds only the clip with this clip id or line id (plus the clip a padded denial is matched to) and updates those entries in the existing manifest. |
| `-answer-gap-ms` | Extra silence between the parts of a joined answer. Default 0: each part keeps its own leading and trailing silence, about 350 ms in total, which sounds like an ordinary sentence break. |
| `-quota-project` | Quota project for the API calls. |
| `-samples <file>` | Voice listening test: synthesises one sample sentence per candidate voice (JSON: role to list of voice names) and exits. |

## What it builds

| Clip id | Parts |
|---|---|
| `setup.audioTest`, every narration, routine question and participant question | one line |
| `setup.exampleScene`, `opening` | all lines of the scene, `timing.linePauseMs` of silence between them, `lineStarts` recorded |
| answers, `mode: "whole"` | the reason line and the denial line |
| answers, `mode: "joined"` | `<slide>.answers.reason` = opening + reasonBody + closing, `<slide>.answers.denial` = opening + closing |

Joined clips are built by concatenating MP3 frames, without re-encoding, so the shared opening and closing are byte-identical in both answers. Silence is added as MP3 frames with all-zero side information (24 ms each at 24 kHz), between lines and, with `padDenialToReason`, at the end of the denial until it matches the reason to within 12 ms. Durations are measured by decoding the MP3 completely (`github.com/hajimehoshi/go-mp3`); they match `ffprobe`.

Lines are sent as SSML: the line's `ssml`, or its `text` XML-escaped inside `<speak>`. The tool uses the `v1beta1` endpoint and asks for `SSML_MARK` timepoints when the SSML has a `<mark>`; if a mark comes back without a timepoint (Studio and Chirp voices do not return them), the run fails and names the line and voice.

## Lint

Before any API call the tool fails if a CISO answer line (reason, denial, opening, reasonBody, closing) carries `<prosody>`, `<emphasis>` or a rate, pitch or volume attribute; if a reason and its denial (or the three joined parts) use different SSML wrappers; or if any asset file name (consent PDF, images, portrait, every written MP3) contains `reason`, `denial`, `checker` or `neutral`. Run `tools/content-lint` for the full rule set.

## Hashes and cache

- A synthesised line's key is the SHA-256 of a JSON object with the voice name, language code, speaking rate, pitch, text, SSML and encoding (`MP3`, 24000 Hz). A one-line clip is that file.
- A joined or padded clip's key is the SHA-256 of its part keys, the gap and the number of padding frames.
- The cache keeps `<key>.mp3` and `<key>.json` (input, marks, duration) for every synthesised line. A line whose key is in the cache is never sent to the API again; a rerun with unchanged content makes no API calls.

## Cost

Google bills per character sent. The estimate counts the full SSML input, tags included, which slightly over-counts, at list price per million characters before the monthly free tier: Neural2 and WaveNet 16 USD, Standard 4 USD, Chirp 3 HD 30 USD, Studio 160 USD (check https://cloud.google.com/text-to-speech/pricing). The v1.4 set is 6,854 characters, about 0.11 USD.

## After a voice change

1. Change the voice name in `tools/tts/voices.default.json` (or set real names in the config's `voices`, which take precedence over the defaults).
2. `go run ./tools/tts -config content/study-config.v1.4.json -dry-run` shows which lines are no longer cached and what they cost.
3. `go run ./tools/tts -config content/study-config.v1.4.json` synthesises them and rewrites the manifest. Lines of unchanged voices come from the cache.
4. Run `go run ./tools/content-lint content/study-config.v1.4.json` again: it reads the new manifest and warns about denials more than 1 s shorter than their reasons.
5. Old MP3s in `content/audio/<version>/` are not deleted; remove files that the manifest no longer names if the directory is uploaded as is.

## Tests

```
cd tools/tts && go test ./... && go vet ./...
```

The tests use fixture MP3s in `testdata/` and a fake synthesiser; they make no API calls.
