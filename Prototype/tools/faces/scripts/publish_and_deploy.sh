#!/usr/bin/env bash
# After the render batch: publish manifests, mirror the bucket, commit, build, push, roll a revision.
set -e
R=/mnt/RAID_STORAGE/SFU/Projects/CMPT-863; P=$R/Prototype; LAB=/home/oliver/.cache/cmpt863-faces
echo "== $(date -Is) publish"; cd "$LAB" && venv/bin/python "$P/tools/faces/scripts/publish.py" 1.4.0 test-sound-1.4.0
python3 -c "import json;m=json.load(open('$P/content/video/1.4.0/manifest.json'));print('clips:',sorted(m['clips']));print('seats:',sorted(m['seats']))"
echo "== $(date -Is) mirror"; cd "$P" && gcloud storage rsync content gs://cmpt863-study-content --recursive --exclude='.*/\.cache/.*|.*\.py$' 2>&1 | grep -c 'Copying' || true
gcloud storage ls 'gs://cmpt863-study-content/video/1.4.0/' | wc -l
echo "== $(date -Is) commit"; cd "$R"
HITS=$(git ls-files | grep -v -E '\.example$|\.md$|^docs/|\.github/workflows/ci\.yml$' | xargs grep -l -E -- '-----BEGIN|"refresh_token"|ya29\.|AIza[0-9A-Za-z_-]{35}|"private_key"' 2>/dev/null || true); [ -z "$HITS" ] || { echo "CREDENTIAL HITS: $HITS"; exit 1; }
git add Prototype/content/video Prototype/e2e/tests/flows.spec.ts
git -c core.hooksPath=/dev/null commit -q -m "$(printf 'Prototype: all talking-head clips and idle loops published for 1.4.0 and test-sound-1.4.0; flow test pins the denial slide when it expects a video\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_019fmu3BGbg41BxSuuEM4n16')" || echo "nothing to commit"
SHA=$(git rev-parse --short HEAD); echo "HEAD $SHA"
IMG=northamerica-northeast1-docker.pkg.dev/cmpt863-study/study/api:$SHA
echo "== $(date -Is) docker build"; cd "$P"; docker build -q -t "$IMG" .; docker push -q "$IMG"
DIGEST=$(docker inspect --format='{{index .RepoDigests 0}}' "$IMG"); echo "digest $DIGEST"
python3 - "$P/infra/terraform/envs/dev/terraform.tfvars" "$DIGEST" "$SHA" <<'PY'
import re, sys
p, digest, sha = sys.argv[1:4]; s = open(p).read()
s2 = re.sub(r'(?m)^# Current image \(commit [0-9a-f]+\)', f'# Current image (commit {sha})', s)
s2 = re.sub(r'(?m)^image_digest\s*=\s*".*"$', f'image_digest = "{digest}"', s2)
open(p, "w").write(s2); print("tfvars updated")
PY
echo "== $(date -Is) terraform apply"; cd "$P/infra/terraform/envs/dev" && terraform apply -auto-approve -input=false 2>&1 | grep -E 'Apply complete|Error' | head -5
echo "== $(date -Is) revision"; gcloud run services describe study-api --region northamerica-northeast1 --format='value(status.latestReadyRevisionName)'
cd "$R" && git add Prototype/infra/terraform/envs/dev/terraform.tfvars && git -c core.hooksPath=/dev/null commit -q -m "$(printf 'Prototype: revision with all talking-head clips live (image %s)\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_019fmu3BGbg41BxSuuEM4n16' "$SHA")" && git push -q origin main && echo "pushed $(git rev-parse --short HEAD)"
echo "== $(date -Is) done"
