#!/usr/bin/env bash
# Starts the API locally with throwaway admin secrets (in a temp directory,
# never in a tracked file), runs the load generator against it, stops it.
#
#   ./run-local.sh memory 50 5            # memory store, 50 sessions, 5 s ramp
#   ./run-local.sh firestore 200 10       # emulator on localhost:8081, fresh project id
#
# Extra arguments after the third are passed to the generator (e.g. -think 0.5).
set -euo pipefail
cd "$(dirname "$0")/.."
store=${1:-memory}; sessions=${2:-50}; ramp=${3:-5}; shift $(( $# < 3 ? $# : 3 ))
port=${LOAD_PORT:-18080}
tmp=$(mktemp -d); trap 'kill ${api_pid:-} 2>/dev/null || true; rm -rf "$tmp"' EXIT

(cd services/api && go build -o "$tmp/api" ./cmd/api && go build -o "$tmp/admintool" ./cmd/admintool)
(cd load && go build -o "$tmp/load" .)
pw=$(head -c 18 /dev/urandom | base64 | tr -d '/+=')
hash=$(printf '%s' "$pw" | "$tmp/admintool" password 2>/dev/null | tail -1)
gate=$("$tmp/admintool" key); skey=$("$tmp/admintool" key)

env_common=(PORT=$port CONTENT_DIR=content ALLOW_TEST_CONFIGS=true INSECURE_COOKIES=true TRUSTED_PROXY_HOPS=1
  ADMIN_PASSWORD_HASH_VALUE="$hash" ADMIN_ACCESS_KEY_VALUE="$gate" ADMIN_SESSION_KEY_VALUE="$skey")
if [ "$store" = memory ]; then
  env "${env_common[@]}" STORE=memory FIRESTORE_EMULATOR_HOST= "$tmp/api" >"$tmp/api.log" 2>&1 &
else
  project=load-$(date +%s)
  echo "Firestore emulator ${FIRESTORE_EMULATOR_HOST:-localhost:8081}, project $project"
  env "${env_common[@]}" STORE=firestore FIRESTORE_EMULATOR_HOST=${FIRESTORE_EMULATOR_HOST:-localhost:8081} \
    FIRESTORE_PROJECT_ID=$project "$tmp/api" >"$tmp/api.log" 2>&1 &
fi
api_pid=$!
for _ in $(seq 1 50); do curl -sf "localhost:$port/api/v1/healthz" >/dev/null && break; sleep 0.2; done

LOAD_GATE_KEY=$gate LOAD_ADMIN_PASSWORD=$pw "$tmp/load" -base "http://localhost:$port" -sessions "$sessions" -ramp "$ramp" "$@" || status=$?
echo "Server-side transaction retries (Firestore only): $(grep -c 'transaction retried' "$tmp/api.log" || true)"
if [ "$store" != memory ]; then (cd load && go run ./cmd/docsize -emulator "${FIRESTORE_EMULATOR_HOST:-localhost:8081}" -project "$project"); fi
exit ${status:-0}
