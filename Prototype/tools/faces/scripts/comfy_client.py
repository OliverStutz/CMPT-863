"""Minimal ComfyUI API client: submit a graph in API format, wait, return output files."""
import json, sys, time, urllib.request, urllib.error, uuid

import os
BASE = "http://127.0.0.1:" + os.environ.get("COMFY_PORT", "8188")

def get(path, base=None):
    with urllib.request.urlopen((base or BASE) + path, timeout=60) as r:
        return json.loads(r.read())

def object_info(base=None):
    return get("/object_info", base)

def validate(graph, info):
    """Check every node type and input name against the server's node registry."""
    problems = []
    for nid, node in graph.items():
        t = node["class_type"]
        if t not in info:
            problems.append(f"node {nid}: unknown class {t}"); continue
        spec = info[t]["input"]
        allowed = set(spec.get("required", {})) | set(spec.get("optional", {}))
        for k in node["inputs"]:
            if k not in allowed:
                problems.append(f"node {nid} ({t}): unknown input '{k}'; allowed {sorted(allowed)}")
        for k in spec.get("required", {}):
            if k not in node["inputs"]:
                problems.append(f"node {nid} ({t}): missing required input '{k}'")
    return problems

def run(graph, label="job", poll=5, base=None):
    """base: the server, e.g. "http://127.0.0.1:8189"; defaults to COMFY_PORT. Safe to call from several threads."""
    base = base or BASE
    info = object_info(base)
    problems = validate(graph, info)
    if problems:
        print("GRAPH INVALID:\n  " + "\n  ".join(problems)); sys.exit(2)
    body = json.dumps({"prompt": graph, "client_id": str(uuid.uuid4())}).encode()
    req = urllib.request.Request(base + "/prompt", data=body, headers={"Content-Type": "application/json"})
    try:
        with urllib.request.urlopen(req, timeout=60) as r:
            resp = json.loads(r.read())
    except urllib.error.HTTPError as e:
        print("SUBMIT FAILED", e.code, e.read().decode()[:2000]); sys.exit(2)
    if resp.get("node_errors"):
        print("NODE ERRORS", json.dumps(resp["node_errors"], indent=1)[:3000]); sys.exit(2)
    pid = resp["prompt_id"]; t0 = time.time()
    print(f"[{label}] submitted {pid}", flush=True)
    while True:
        time.sleep(poll)
        hist = get(f"/history/{pid}", base)
        if pid in hist:
            h = hist[pid]; status = h.get("status", {})
            el = time.time() - t0
            if status.get("status_str") == "error":
                msgs = [m for m in status.get("messages", []) if m[0] == "execution_error"]
                print(f"[{label}] ERROR after {el:.0f}s:", json.dumps(msgs, indent=1)[:4000]); sys.exit(1)
            outs = []
            for nid, o in h.get("outputs", {}).items():
                for kind in ("images", "videos", "audio", "gifs"):
                    for f in o.get(kind, []):
                        outs.append((kind, f.get("subfolder", ""), f["filename"]))
            print(f"[{label}] done in {el:.0f}s; outputs: {outs}", flush=True)
            return outs, el
        q = get("/queue", base)
        running = len(q.get("queue_running", [])); pending = len(q.get("queue_pending", []))
        print(f"[{label}] {time.time()-t0:.0f}s running={running} pending={pending}", flush=True)
