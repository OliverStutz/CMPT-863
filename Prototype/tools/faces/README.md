# Talking heads: rendering the faces for the study

Everything the study shows as a face is rendered offline from the audio clips with this folder's scripts and
shipped as MP4 files beside the audio (`content/video/<version>/`, architecture change log 1.32). No real
person is depicted: every portrait is synthetic. The participant's own seat ("Board member (You)") never
gets a face, and the participant's spoken questions stay audio only.

The scripts run against a *lab folder* outside the repository, `FACES_LAB` (default
`/home/oliver/.cache/cmpt863-faces`), which holds the 40 GB of model files, the ComfyUI checkout, the
Python environment, staged inputs and rendered outputs. The repository holds the scripts, the cast, the
model list with checksums, and the published manifests (the MP4s themselves are gitignored like the MP3s).

## Hardware and runtime (verified 2026-10-25)

- Two AMD Radeon AI PRO R9700 (gfx1201, 32 GB each) on ROCm 7.2.4. The official PyTorch wheel for ROCm 7.2
  works out of the box (measured 130 TFLOP/s bf16 matrix multiply per card).
- The Ryzen's integrated GPU shows up as a third ROCm device and crashes PyTorch. Hide it with
  `ROCR_VISIBLE_DEVICES=<index>` **only**; setting `HIP_VISIBLE_DEVICES` as well filters twice and leaves no device.
- One ComfyUI server per card, on ports 8188 and 8189, pinned with `ROCR_VISIBLE_DEVICES=0` and `=1`.
- RAM: each server keeps a CPU copy of its weights, about 24 to 28 GB, and the machine has no swap. The
  servers start with ComfyUI's node cache off (`--cache-none`) so decoded frames are not kept between jobs,
  the batch pauses while less than 6 GB is available (`RAM_MIN_MB`), and the servers are marked as the first
  processes to be killed on an out-of-memory event (`oom_score_adj=1000`), never the user's own programs.

## One-time setup

```sh
export FACES_LAB=/home/oliver/.cache/cmpt863-faces
mkdir -p "$FACES_LAB" && cd "$FACES_LAB"
git clone https://github.com/comfyanonymous/ComfyUI.git && git -C ComfyUI checkout fe298e4e20f1ff26d30e8639267bbd0b57811390
uv venv --python 3.12 venv
uv pip install --python venv/bin/python torch torchvision torchaudio --index-url https://download.pytorch.org/whl/rocm7.2
grep -v -i -E '^(torch|torchvision|torchaudio)([<>=~ ]|$)' ComfyUI/requirements.txt > req-notorch.txt
uv pip install --python venv/bin/python -r req-notorch.txt
cp <repo>/Prototype/tools/faces/extra_model_paths.yaml .
<repo>/Prototype/tools/faces/scripts/download_models.sh     # 40 GB, resumable, verifies every SHA-256
```

Verification done before the first run, to be repeated after any update: the ComfyUI checkout matches the
upstream commit (`git ls-remote origin HEAD` against `git rev-parse HEAD`), `custom_nodes/` holds only the two
stock files, `requirements.txt` is the upstream list, and the server runs with `--offline` on localhost. Model
files are safetensors (no code) and are checked against `models.sha256`, the SHA-256 values Hugging Face
publishes in the files' LFS metadata. Licences: Wan 2.2 S2V, the Wan 2.1 VAE, umt5, wav2vec2, the lightx2v
LoRA and Flux.1-schnell are all Apache 2.0.

## Running the servers

```sh
PORT=8188 ROCR_VISIBLE_DEVICES=0 nohup tools/faces/scripts/start_server.sh > $FACES_LAB/server.log 2>&1 &
PORT=8189 ROCR_VISIBLE_DEVICES=1 nohup tools/faces/scripts/start_server.sh > $FACES_LAB/server2.log 2>&1 &
curl -s http://127.0.0.1:8188/system_stats | python3 -m json.tool | grep -E 'name|vram_total'   # 31.9 GiB = an R9700
pkill -f '[m]ain.py --listen 127.0.0.1'                                                           # stop both
```

`run_batch.py` restarts a server that stops answering. Where the weights actually are can be read from the
kernel: `/sys/class/kfd/kfd/proc/<server pid>/vram_<gpu id>` (bytes per GPU), and
`/sys/class/kfd/kfd/topology/nodes/*/properties` maps a GPU id to its render node and PCI address.

## Staging the inputs

The batch reads audio from `$FACES_LAB/inputs/`, named after the clip id with dots replaced by underscores,
for example `s3_narration.mp3`. Copy them from the audio manifest of the version with a few lines of Python:
for each clip id in `TALKING` (in `run_batch.py`) copy `content/audio/<version>/<audioKey>` to
`$FACES_LAB/inputs/<id with dots as underscores>.mp3`. The opening has two speakers and is split at the start
of the CISO's line, which the audio manifest records as `lineStarts["opening.ciso.5"]`:

```sh
ffmpeg -y -i content/audio/1.4.0/<opening audioKey> -t 56.256 -c:a libmp3lame -q:a 2 $FACES_LAB/inputs/opening_chair.mp3
ffmpeg -y -i content/audio/1.4.0/<opening audioKey> -ss 56.256 -c:a libmp3lame -q:a 2 $FACES_LAB/inputs/opening_ciso.mp3
ffmpeg -y -f lavfi -i anullsrc=r=16000:cl=mono -t 9.65 -c:a pcm_s16le $FACES_LAB/inputs/silence_9650ms.wav   # idle loops
```

The clip ids are the keys of the audio manifest; `TALKING` lists which ids get a face and who speaks them. A
new or renamed clip is added there.

## The cast

`scripts/characters.py` holds one entry per role (`ciso`, `chair`, `member2` to `member6`): the portrait
prompt, the seed, and the speaking prompt for the roles that talk. Portraits are made with Flux.1-schnell,
four candidates per role in `$FACES_LAB/out/portrait/`, and the first candidate becomes
`$FACES_LAB/inputs/ref_<role>.png`, the reference the video model animates. To use another candidate, copy
it over `ref_<role>.png`; to change a face, edit the prompt or seed, delete `ref_<role>.png`, and remove the
role's lines from `results.jsonl` (below) so its clips are rendered again.

## Rendering

```sh
cd $FACES_LAB && venv/bin/python <repo>/Prototype/tools/faces/scripts/run_batch.py        # everything missing
venv/bin/python .../run_batch.py --only s3.narration idle:member4                           # named jobs only
venv/bin/python .../run_batch.py --full                                                     # 20 steps, no speed LoRA
```

What it does: portraits first (server 1 only), then talking clips and idle loops on both servers, longest
first. A talking clip is Wan 2.2 S2V driven by the staged audio and the role's portrait, 640 by 640 at 16
frames per second, rendered in chunks of 77 frames (about 4.8 s) that the model extends one after another
while re-reading the portrait, then trimmed to the audio length. An idle loop is the same model driven by
9.65 s of silence with a "listening, mouth closed" prompt. Settings: 4 sampling steps with the lightx2v
LoRA (about 85 to 100 s per chunk per card) or `--full` for 20 steps without it (about five times slower).
Seeds are fixed per clip, so a re-run reproduces a clip exactly.

Every finished job appends a line to `$FACES_LAB/out/batch/results.jsonl` and leaves its file in
`$FACES_LAB/out/batch/`; a re-run skips jobs that have a line. Delete a line to render that job again, or
the whole file to start over. Progress is in `$FACES_LAB/out/batch.log`. The full set is 26 jobs and 94
chunks, about 75 minutes on both cards.

## Publishing

```sh
venv/bin/python <repo>/Prototype/tools/faces/scripts/publish.py 1.4.0 test-sound-1.4.0
cd <repo>/Prototype && gcloud storage rsync content gs://cmpt863-study-content --recursive --exclude='.*/\.cache/.*|.*\.py$'
# then build, push and roll a revision as in "05 - Progress.md" (the service reads the manifests at start)
e2e/security.sh https://study-api-pju4jz7m4a-nn.a.run.app <run code>          # includes three video checks
E2E_CONFIG=test-sound-1.4.0 E2E_EXPECT_VIDEO=1 npx playwright test tests/flows.spec.ts --timeout 1500000
```

`publish.py` writes `content/video/<version>/manifest.json` and copies each MP4 as `<sha256>.mp4`:

```json
{"configVersion": "1.4.0",
 "clips": {"s2.answer.a": {"role": "ciso", "videoKey": "<sha256>.mp4", "durationMs": 15562, "width": 640, "height": 640},
           "opening": {"durationMs": 59256, "segments": [{"role": "chair", "startMs": 0, "videoKey": "...", "durationMs": 56256},
                                                          {"role": "ciso", "startMs": 56256, "videoKey": "...", "durationMs": 3000}]}},
 "seats": {"member2": {"videoKey": "<sha256>.mp4", "durationMs": 9937, "width": 640, "height": 640}}}
```

The server reads it beside the audio manifest, serves `video/<version>/*.mp4` publicly (the manifest never),
puts `videoUrl` and `videoRole` (or `videoSegments` for a multi-speaker clip) on the step, and `video.seats` in
the public scene pack. `role` says whose face speaks a clip (the chair asks the routine questions, so those
clips carry `chair`); the browser lights that seat and shows the clip on it, and shows the seat's idle loop
outside the clip. Every face in `IDLE` gets a loop, the CISO included. The browser keeps the one audio element as the clock and plays the video muted in sync with it;
a missing or failing video changes nothing in the flow.

## Re-generating after a change

1. Text changed: regenerate the audio with `tools/tts`, mirror the bucket, then re-stage the changed clips
   (section "Staging the inputs"), delete their lines from `results.jsonl`, and run the batch.
2. Face changed: edit `characters.py`, delete `ref_<role>.png` and the role's result lines, run the batch.
3. Quality: `--full` for 20 steps; width and height in `graphs.s2v()`; the chunk length stays 77 (the model
   needs at least 73 frames); the four frames trimmed after decoding compensate the doubled first latent
   frame of the official workflow. Long clips drift: every extension re-reads the portrait, but colour and
   sharpness still wander after about eight chunks (the chair's 56 s opening, twelve chunks, goes lighter
   and softer in its last third, while the CISO's 41 s narration holds). A clip beyond about 40 s is better
   split at a sentence boundary into two segments of the same role, which needs `publish.py` to take the
   second segment's line offset (today it keys segments by role) and the browser already plays segments.
4. Publish, mirror, roll a revision, run the checks. Keep `content/video/<version>/manifest.json` in git.

## Preview without deploying

`scripts/make_preview.py` writes a self-contained HTML page with the real stage geometry of the participant
app, the slide 2 render, the CISO clip and the seat 3 loop embedded, to `$FACES_LAB/out/preview/`.
