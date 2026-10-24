# tools/content-lint

Validates a study config against `content/schema/study-config.schema.json` (draft 2020-12, via `github.com/santhosh-tekuri/jsonschema/v6` through `studyconfig.NewValidator`) and applies the rules of section 4.2 of `02 - Architecture.md`. Exit code 1 on any failure; 0 when there are only warnings or nothing at all; 2 on a usage error.

## Usage

From `Prototype/`:

```
go run ./tools/content-lint content/study-config.v1.4.json
go run ./tools/content-lint -no-manifest content/study-config.v1.4.json
go run ./tools/content-lint -strict content/study-config.v1.4.json
```

Several configs can be given at once.

## Flags

| Flag | Meaning |
|---|---|
| `-schema` | JSON Schema file. Default `<config dir>/schema/study-config.schema.json`. |
| `-rules` | Rules file with the deny lists. Default: the built-in `rules.json`. |
| `-manifest` | Audio manifest for the post-TTS checks. Default `<config dir>/audio/<version>/manifest.json` when it exists. |
| `-no-manifest` | Skip the post-TTS checks. |
| `-resolved` | The config is a resolved config: `audioKey`, `durationMs` and `marks` may be set, and every `audioKey` must be `<sha256>.mp3`. |
| `-strict` | A flag without `vlmRun` fails instead of warning (use before a config is locked for a run). |

## Failures

- The schema rejects the config.
- A marking `label` or `detail` mentions honesty, lying, deceiving, misleading, hiding, truth or trust, or carries a confidence or accuracy figure (a percentage, "confidence", "accuracy", "probability", "certain", "likely to be correct").
- A CISO line says "AI", "artificial intelligence", "AI model", "language model", LLM, ChatGPT, "algorithm", "machine learning" or "checker".
- Text that describes the checker before the debrief (`labels`, marking labels and details, `setup.preStartText`) says "AI", "artificial intelligence", LLM, "language model" or "machine learning".
- `labels.noIssues` is not the configured label (`"No issues found"` in `rules.json`), or a no-issues panel sets a `text` different from it.
- A CISO answer line carries `<prosody>`, `<emphasis>` or a rate, pitch or volume change, or a slide's reason and denial (or joined parts) use different SSML wrappers.
- A marking's `anchorText` does not occur in the narration (case and spacing ignored); an `appearAtMark` has no `<mark name>` in the narration SSML; a marking or no-issues panel has neither `appearAtMark` nor `appearAtMs`; a target id is not in the layout; a bbox lies outside 1600 x 900.
- A slide's first block is not its only title block; there is no slide or no flagged slide; a flagged slide lacks its questions or answers, or an unflagged slide has them; action, question and answer scenes are not exactly one per flagged slide.
- A rating point is duplicated or differs from the id the plan derives (T0, T<n>, T<n>a, T<n>b), or the opening, a slide or an answer scene is not followed by its rating.
- `correctOptionId` is not an option, the background has no honeypot, `doNotDiscuss` is empty.
- A source config sets `audioKey`, `durationMs` or `marks`.
- An asset file name contains `reason`, `denial`, `checker` or `neutral`.
- With a manifest: a flag or panel appears after the narration ends.

## Warnings

- The number of slides is not five, the number of flagged slides is not three, or the flagged slides are not 2, 3 and 4: "differs from Abstract v8". With fewer than three flagged slides a second warning names the design limit.
- Rating scales or order, or the budget's option count, differ from Oliver's decisions.
- A flag has no `vlmRun`.
- With a manifest: a denial is more than 1,000 ms shorter than its reason and `padDenialToReason` is off.

## Extending the deny lists

Edit `rules.json` (Go regular expressions, `(?i)` for case-insensitive) or pass another file with `-rules`. No code change is needed. A lint cannot judge tone; that review stays with the authors.

## Tests

```
cd tools/content-lint && go test ./... && go vet ./...
```
