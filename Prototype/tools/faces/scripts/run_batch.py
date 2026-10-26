"""Render every talking-head clip and idle loop of the study on the two cards.

Jobs: portraits (Flux, server 1 only, four candidates, the first is used), talking clips (Wan 2.2 S2V from the
staged audio and the role's portrait), idle loops (same model with a silent track). Jobs of a role wait for its
portrait. Longest jobs first. A RAM guard pauses submissions while less than RAM_MIN_MB is available. A server
that stops answering is restarted. Results go to out/batch/results.jsonl, one line per finished job, so a
re-run skips what is done; delete a line (or the file) to render again. Usage:
    venv/bin/python scripts/run_batch.py            # everything missing
    venv/bin/python scripts/run_batch.py --only s3.narration idle:member4 portrait:chair
    venv/bin/python scripts/run_batch.py --full     # 20 steps, no speed LoRA (about five times slower)
"""
import json, math, os, shutil, subprocess, sys, threading, time, urllib.request, zlib
sys.path.insert(0, os.path.dirname(__file__))
import graphs
from characters import CAST, portrait_prompt, idle_prompt

LAB = os.environ.get("FACES_LAB", "/home/oliver/.cache/cmpt863-faces")
OUT, INP, BATCH = f"{LAB}/out", f"{LAB}/inputs", f"{LAB}/out/batch"
PORTS = [int(p) for p in os.environ.get("COMFY_PORTS", "8188,8189").split(",")]
FPS, CHUNK, RAM_MIN_MB = 16, 77, int(os.environ.get("RAM_MIN_MB", "6000"))
FULL = "--full" in sys.argv
STEPS, CFG, LORA = (20, 6.0, False) if FULL else (4, 1.0, True)
SILENCE = "silence_9650ms.wav"

# clip id (= audio manifest key, or opening.<role> for the split opening) -> speaker and staged audio file
TALKING = {
    "s1.narration": ("ciso", "s1_narration.mp3"), "s2.narration": ("ciso", "s2_narration.mp3"),
    "s3.narration": ("ciso", "s3_narration.mp3"), "s4.narration": ("ciso", "s4_narration.mp3"),
    "s5.narration": ("ciso", "s5_narration.mp3"),
    "s2.answer.a": ("ciso", "s2_answer_a.mp3"), "s2.answer.b": ("ciso", "s2_answer_b.mp3"),
    "s3.answers.reason": ("ciso", "s3_answers_reason.mp3"), "s3.answers.denial": ("ciso", "s3_answers_denial.mp3"),
    "s4.answer.a": ("ciso", "s4_answer_a.mp3"), "s4.answer.b": ("ciso", "s4_answer_b.mp3"),
    "opening.chair": ("chair", "opening_chair.mp3"), "opening.ciso": ("ciso", "opening_ciso.mp3"),
    "setup.exampleScene": ("chair", "setup_exampleScene.mp3"),
}
IDLE = ["ciso", "chair", "member2", "member3", "member4", "member5", "member6"]  # every face gets a listening loop

def audio_ms(path):
    out = subprocess.run(["ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "csv=p=0", path], capture_output=True, text=True).stdout.strip()
    return int(round(float(out) * 1000))

def mem_available_mb():
    for line in open("/proc/meminfo"):
        if line.startswith("MemAvailable:"):
            return int(line.split()[1]) // 1024
    return 0

def log(msg):
    print(time.strftime("%H:%M:%S"), msg, flush=True)

def load_results():
    done = {}
    if os.path.exists(f"{BATCH}/results.jsonl"):
        for line in open(f"{BATCH}/results.jsonl"):
            if line.strip():
                r = json.loads(line); done[r["job"]] = r
    return done

def record(r):
    with lock:
        with open(f"{BATCH}/results.jsonl", "a") as f:
            f.write(json.dumps(r) + "\n")

lock = threading.Lock()

def ensure_server(port):
    """Returns when the server answers; restarts it if it does not."""
    for attempt in range(2):
        try:
            urllib.request.urlopen(f"http://127.0.0.1:{port}/system_stats", timeout=10).read(); return
        except Exception:
            log(f"[{port}] server not answering, restarting")
            idx = PORTS.index(port)
            subprocess.Popen(["bash", "-c", f"PORT={port} ROCR_VISIBLE_DEVICES={idx} nohup {LAB}/scripts/start_server.sh > {LAB}/server-{port}.log 2>&1 &"])
            for _ in range(180):
                time.sleep(1)
                try:
                    urllib.request.urlopen(f"http://127.0.0.1:{port}/system_stats", timeout=5).read(); break
                except Exception:
                    pass
            for pid in subprocess.run(["pgrep", "-f", f"[m]ain.py --listen 127.0.0.1 --port {port}"], capture_output=True, text=True).stdout.split():
                try: open(f"/proc/{pid}/oom_score_adj", "w").write("1000")
                except OSError: pass
    raise RuntimeError(f"server on port {port} did not come back")

def ram_guard(port):
    waited = 0
    while mem_available_mb() < RAM_MIN_MB:
        if waited == 0: log(f"[{port}] RAM guard: {mem_available_mb()} MB available, waiting")
        time.sleep(15); waited += 15
    if waited: log(f"[{port}] RAM guard released after {waited}s ({mem_available_mb()} MB available)")

def run_job(port, job):
    import comfy_client
    base = f"http://127.0.0.1:{port}"
    kind, name = job["kind"], job["name"]
    t0 = time.time()
    if kind == "portrait":
        outs, _ = comfy_client.run(graphs.portrait(portrait_prompt(name), CAST[name]["seed"], prefix=f"portrait/{name}"), label=f"[{port}] portrait:{name}", base=base)
        files = sorted(f for k, sub, f in outs if k == "images")
        shutil.copy(os.path.join(OUT, "portrait", files[0]), os.path.join(INP, f"ref_{name}.png"))
        return {"job": job["id"], "kind": kind, "role": name, "candidates": files, "used": files[0], "seconds": round(time.time() - t0)}
    if kind == "talk":
        role, audio = TALKING[name]
        ms = audio_ms(os.path.join(INP, audio)); chunks = max(1, math.ceil(ms / 1000 * FPS / CHUNK))
        g = graphs.s2v(audio, f"ref_{role}.png", CAST[role]["talk"], chunks, steps=STEPS, cfg=CFG, lora=LORA, seed=zlib.crc32(name.encode()) % 100000, prefix=f"s2v/{name}")
    else:
        role = name
        ms = 9625; chunks = 2
        g = graphs.s2v(SILENCE, f"ref_{role}.png", idle_prompt(role), chunks, steps=STEPS, cfg=CFG, lora=LORA, seed=zlib.crc32(("idle" + name).encode()) % 100000, prefix=f"s2v/idle_{name}")
    outs, _ = comfy_client.run(g, label=f"[{port}] {kind}:{name} ({chunks} chunks)", poll=15, base=base)
    vids = [os.path.join(OUT, sub, f) for k, sub, f in outs if f.endswith(".mp4")]
    if not vids: raise RuntimeError(f"no mp4 output for {name}: {outs}")
    os.makedirs(BATCH, exist_ok=True)
    final = os.path.join(BATCH, f"{job['id'].replace(':', '_')}.mp4")
    if kind == "talk":
        subprocess.run(["ffmpeg", "-v", "error", "-y", "-i", vids[0], "-t", f"{ms/1000:.3f}", "-c:v", "libx264", "-preset", "medium", "-crf", "18",
                        "-pix_fmt", "yuv420p", "-c:a", "copy", "-movflags", "+faststart", final], check=True)
    else:
        subprocess.run(["ffmpeg", "-v", "error", "-y", "-i", vids[0], "-an", "-c:v", "copy", "-movflags", "+faststart", final], check=True)
    probe = json.loads(subprocess.run(["ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height,nb_frames,duration", "-of", "json", final], capture_output=True, text=True).stdout)["streams"][0]
    return {"job": job["id"], "kind": kind, "role": role, "clip": name if kind == "talk" else None, "audioMs": ms, "chunks": chunks, "steps": STEPS, "lora": LORA,
            "file": final, "width": int(probe["width"]), "height": int(probe["height"]), "frames": int(probe.get("nb_frames", 0)),
            "durationMs": int(round(float(probe["duration"]) * 1000)), "seconds": round(time.time() - t0), "server": port}

def build_jobs(only):
    jobs = []
    for role in CAST:
        jobs.append({"id": f"portrait:{role}", "kind": "portrait", "name": role, "chunks": 0, "needs": None})
    for name, (role, audio) in TALKING.items():
        ms = audio_ms(os.path.join(INP, audio)); chunks = max(1, math.ceil(ms / 1000 * FPS / CHUNK))
        jobs.append({"id": name, "kind": "talk", "name": name, "chunks": chunks, "needs": f"portrait:{role}"})
    for role in IDLE:
        jobs.append({"id": f"idle:{role}", "kind": "idle", "name": role, "chunks": 2, "needs": f"portrait:{role}"})
    if only:
        jobs = [j for j in jobs if j["id"] in only or (j["kind"] == "portrait" and any(o == j["id"] for o in only))]
    return jobs

def main():
    only = [a for a in sys.argv[1:] if not a.startswith("--")]
    os.makedirs(BATCH, exist_ok=True)
    done = load_results()
    # portraits that already have a reference image count as done
    for role in CAST:
        if os.path.exists(os.path.join(INP, f"ref_{role}.png")) and f"portrait:{role}" not in done:
            done[f"portrait:{role}"] = {"job": f"portrait:{role}", "kind": "portrait", "existing": True}
    jobs = [j for j in build_jobs(only) if j["id"] not in done]
    total_chunks = sum(j["chunks"] for j in jobs)
    log(f"{len(jobs)} job(s), {total_chunks} chunks, steps={STEPS} lora={LORA}, servers {PORTS}, RAM guard {RAM_MIN_MB} MB")
    # portraits first (every talking clip and idle loop of a role waits for its reference face), then longest first
    pending = sorted(jobs, key=lambda j: (j["kind"] != "portrait", -j["chunks"], j["id"]))
    running = set(); failures = []

    def worker(port):
        ensure_server(port)
        while True:
            with lock:
                pick = None
                for j in pending:
                    if j["kind"] == "portrait" and port != PORTS[0]: continue  # Flux lives on server 1 only
                    if j["needs"] and j["needs"] not in done: continue
                    pick = j; break
                if not pending:
                    return
                if pick is None:
                    waiting = True
                else:
                    pending.remove(pick); running.add(pick["id"]); waiting = False
            if waiting:
                time.sleep(10); continue
            ram_guard(port)
            try:
                ensure_server(port)
                r = run_job(port, pick)
                with lock:
                    done[pick["id"]] = r; running.discard(pick["id"])
                record(r)
                log(f"[{port}] done {pick['id']} in {r.get('seconds')}s; RAM available {mem_available_mb()} MB")
            except (SystemExit, Exception) as e:
                reason = f"exit {e.code}" if isinstance(e, SystemExit) else str(e)[:300]
                with lock:
                    running.discard(pick["id"]); failures.append((pick["id"], reason))
                    if pick["kind"] == "portrait":  # nothing can be rendered for this role
                        for j in [j for j in pending if j["needs"] == pick["id"]]:
                            pending.remove(j); failures.append((j["id"], f"portrait failed: {reason}"))
                log(f"[{port}] FAILED {pick['id']}: {reason}")

    threads = [threading.Thread(target=worker, args=(p,), daemon=True) for p in PORTS]
    for t in threads: t.start()
    for t in threads: t.join()
    log(f"batch finished; {len(failures)} failure(s): {failures}")
    sys.exit(1 if failures else 0)

if __name__ == "__main__":
    main()
