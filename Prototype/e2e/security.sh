#!/usr/bin/env bash
# Outside-in security checks against a running API: ./security.sh <base url> <open run code>
# Prints PASS or FAIL per check and exits 1 if any check failed.
B=${1:?base url}; RUN=${2:?run code}; T=$(mktemp -d); PJ=$T/p.jar; FAILED=0
c() { curl -s -o /dev/null -w "%{http_code}" "$@"; }
check() { local want=$1 got=$2 what=$3; if [ "$want" = "$got" ]; then echo "PASS $got $what"; else echo "FAIL got $got want $want $what"; FAILED=1; fi; }

# Traversal probes: locally the service answers 404 outright. Behind Cloud Run,
# Google's front end cleans the path first and answers 302 to the cleaned path
# without the request ever reaching the service (verified in the request logs
# on 2026-10-24). A redirect whose target then answers 404 is therefore accepted.
traversal() { local what=$1; shift; local code; code=$(curl -s --path-as-is -o /dev/null -w '%{http_code}' "$@")
  if [ "$code" = "302" ] || [ "$code" = "301" ]; then
    local followed; followed=$(curl -s --path-as-is -L -o /dev/null -w '%{http_code}' "$@")
    if [ "$followed" = "404" ]; then echo "PASS 404 (after front-end redirect) $what"; return; fi
    echo "FAIL got $code then $followed want 404 $what"; FAILED=1; return
  fi
  check 404 "$code" "$what"; }

SID=$(curl -s -c $PJ -H 'X-Study-Client: 1' -H 'Content-Type: application/json' -d '{"screen":{"w":1440,"h":900},"tz":"UTC"}' $B/api/v1/runs/$RUN/sessions | python3 -c 'import json,sys;print(json.load(sys.stdin)["sessionId"])')
check 200 $(c -b $PJ -H 'X-Study-Client: 1' $B/api/v1/sessions/$SID/next) "GET next with cookie and header (control)"
check 401 $(c -H 'X-Study-Client: 1' $B/api/v1/sessions/$SID/next) "GET next without cookie"
check 403 $(c -b $PJ $B/api/v1/sessions/$SID/next) "GET next without X-Study-Client"
check 401 $(c -H 'X-Study-Client: 1' -H 'Content-Type: application/json' -d '{"clientSeq":1,"type":"device_checked","payload":{}}' $B/api/v1/sessions/$SID/events) "POST events without cookie"
check 403 $(c -b $PJ -H 'Content-Type: application/json' -d '{"clientSeq":1,"type":"device_checked","payload":{}}' $B/api/v1/sessions/$SID/events) "POST events without X-Study-Client"
check 401 $(c -H 'X-Study-Client: 1' $B/api/v1/sessions/$SID/other) "unknown session path without cookie"
check 403 $(c -H 'Content-Type: application/json' -d '{"screen":{"w":1440,"h":900},"tz":"UTC"}' $B/api/v1/runs/$RUN/sessions) "POST sessions without X-Study-Client"
check 404 $(c $B/admin/) "/admin/ without gate"
check 404 $(c $B/admin/assets/) "/admin/assets/ without gate"
check 404 $(c $B/api/admin/participants) "/api/admin/participants without gate"
check 404 $(c -H 'X-Study-Admin: 1' -H 'Content-Type: application/json' -d '{"password":"x"}' $B/api/admin/login) "/api/admin/login without gate"
check 404 $(c "$B/admin/?k=wrong-key") "/admin/?k=wrong"
check 404 $(c $B/assets/study-config.v1.4.json) "/assets/study-config.v1.4.json"
check 404 $(c $B/assets/schema/study-config.schema.json) "/assets/schema/study-config.schema.json"
check 404 $(c $B/assets/vlm-drafts/v1.3/selection.md) "/assets/vlm-drafts/v1.3/selection.md"
traversal "/assets/../services/api/go.mod (path as is)" "$B/assets/../services/api/go.mod"
check 404 $(c $B/assets/../services/api/go.mod) "/assets/../services/api/go.mod (curl-normalised)"
traversal "/assets/%2e%2e/services/api/go.mod" "$B/assets/%2e%2e/services/api/go.mod"
check 404 $(c $B/a/vlm-drafts/v1.3/s2.raw.md) "/a/vlm-drafts/v1.3/s2.raw.md"
check 404 $(c $B/assets/configs/1.4.0.json) "/assets/configs/1.4.0.json"
check 200 $(c $B/assets/slides/v1.3/s1.svg) "/assets/slides/v1.3/s1.svg"
check 200 $(c $B/assets/slides/v1.3/s1.png) "/assets/slides/v1.3/s1.png"
check 200 $(c $B/api/v1/healthz) "/api/v1/healthz"
rm -rf $T
exit $FAILED
