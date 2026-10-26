"""Builds content/video/<version>/manifest.json from the batch results and copies the MP4s as <sha256>.mp4.

Reads out/batch/results.jsonl (run_batch.py) and the audio manifest of the version (for the line offset of the
split opening). Talking clips are keyed by their audio clip id; the opening becomes `segments` (chair, then the
CISO at the recorded start of his line); idle loops go to `seats`. Then mirror the bucket and roll a revision
(see ../README.md). Usage: venv/bin/python scripts/publish.py 1.4.0 test-sound-1.4.0"""
import datetime, hashlib, json, os, shutil, sys
LAB = os.environ.get("FACES_LAB", "/home/oliver/.cache/cmpt863-faces")
P = os.environ.get("PROTOTYPE_DIR", os.path.join(os.path.dirname(os.path.abspath(__file__)), "..", "..", ".."))
versions = [a for a in sys.argv[1:] if not a.startswith("--")] or ["1.4.0", "test-sound-1.4.0"]
results = {}
for line in open(f"{LAB}/out/batch/results.jsonl"):
    if line.strip():
        r = json.loads(line); results[r["job"]] = r  # the last line for a job wins
def sha(p): return hashlib.sha256(open(p, "rb").read()).hexdigest()
def entry(r): return {"videoKey": sha(r["file"]) + ".mp4", "durationMs": r["durationMs"], "width": r["width"], "height": r["height"]}
for version in versions:
    audio = json.load(open(f"{P}/content/audio/{version}/manifest.json"))["clips"]
    d = f"{P}/content/video/{version}"; os.makedirs(d, exist_ok=True)
    man = {"configVersion": version, "generatedAt": datetime.datetime.now(datetime.timezone.utc).isoformat(timespec="seconds"),
           "_note": "Talking heads rendered offline with Wan 2.2 S2V from the audio clips and synthetic portraits (tools/faces). "
                    "Clips carry the same audio as the MP3; the browser plays them muted in sync with the audio element.",
           "clips": {}, "seats": {}}
    files = {}
    for job, r in results.items():
        if r["kind"] == "portrait": continue
        e = entry(r); files[e["videoKey"]] = r["file"]
        if r["kind"] == "idle":
            man["seats"][r["role"]] = e
        elif r["clip"] in audio:
            man["clips"][r["clip"]] = e
        elif r["clip"].startswith("opening."):
            role = r["clip"].split(".", 1)[1]
            starts = audio["opening"]["lineStarts"]
            first = min(ms for lid, ms in starts.items() if f".{role}." in lid)
            seg = man["clips"].setdefault("opening", {"segments": []})
            seg["segments"].append({"role": role, "startMs": first, **e})
        else:
            print("skipped (no audio clip):", r["clip"])
    if "opening" in man["clips"]:
        man["clips"]["opening"]["segments"].sort(key=lambda s: s["startMs"])
        man["clips"]["opening"]["durationMs"] = audio["opening"]["durationMs"]
    for key, src in files.items():
        dst = os.path.join(d, key)
        if not os.path.exists(dst): shutil.copy(src, dst)
    json.dump(man, open(os.path.join(d, "manifest.json"), "w"), indent=1)
    print(f"{version}: {len(man['clips'])} clip(s), {len(man['seats'])} seat loop(s), {len(files)} file(s) in {d}")
