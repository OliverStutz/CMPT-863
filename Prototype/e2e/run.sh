#!/usr/bin/env bash
# Runs the end-to-end flows and the security checks (see README.md).
#
# Without E2E_BASE_URL: starts the Firestore emulator, builds and starts the API on E2E_PORT
# (default 18080) with throwaway admin secrets made for this run only, runs everything, stops the API.
# With E2E_BASE_URL: uses that running API; E2E_GATE_KEY and E2E_ADMIN_PASSWORD must then be set.
set -euo pipefail

ROOT=$(cd "$(dirname "$0")/.." && pwd)
E2E="$ROOT/e2e"
TMP=$(mktemp -d)
API_PID=""
cleanup() {
  if [ -n "$API_PID" ]; then kill "$API_PID" 2>/dev/null || true; fi
  rm -rf "$TMP"
}
trap cleanup EXIT

cd "$E2E"
[ -d node_modules ] || npm ci --no-audit --no-fund
npx playwright install firefox >/dev/null

if [ -z "${E2E_BASE_URL:-}" ]; then
  PORT=${E2E_PORT:-18080}
  export E2E_BASE_URL="http://localhost:$PORT"
  (cd "$ROOT" && docker compose up -d --wait firestore)

  # Throwaway secrets: generated here, kept in this process's environment only.
  E2E_ADMIN_PASSWORD=$(head -c 24 /dev/urandom | base64 | tr -d '/+=')
  export E2E_ADMIN_PASSWORD
  cd "$ROOT/services/api"
  ADMIN_PASSWORD_HASH_VALUE=$(printf '%s\n' "$E2E_ADMIN_PASSWORD" | go run ./cmd/admintool password 2>/dev/null)
  ADMIN_ACCESS_KEY_VALUE=$(go run ./cmd/admintool key)
  ADMIN_SESSION_KEY_VALUE=$(go run ./cmd/admintool key)
  export ADMIN_PASSWORD_HASH_VALUE ADMIN_ACCESS_KEY_VALUE ADMIN_SESSION_KEY_VALUE
  export E2E_GATE_KEY="$ADMIN_ACCESS_KEY_VALUE"

  go build -o "$TMP/api" ./cmd/api
  PORT=$PORT FIRESTORE_EMULATOR_HOST=localhost:8081 FIRESTORE_PROJECT_ID=demo-study \
    CONTENT_DIR="$ROOT/content" TRUSTED_PROXY_HOPS=0 INSECURE_COOKIES=true ALLOW_TEST_CONFIGS=true \
    "$TMP/api" >"$TMP/api.log" 2>&1 &
  API_PID=$!
  for _ in $(seq 60); do
    curl -sf "$E2E_BASE_URL/api/v1/healthz" >/dev/null && break
    sleep 1
  done
  curl -sf "$E2E_BASE_URL/api/v1/healthz" >/dev/null || { cat "$TMP/api.log"; exit 1; }
fi
: "${E2E_GATE_KEY:?set E2E_GATE_KEY (the admin gate key of the running API)}"
: "${E2E_ADMIN_PASSWORD:?set E2E_ADMIN_PASSWORD (the admin password of the running API)}"

# A run for the security checks (the Playwright test creates its own).
JAR="$TMP/admin.jar"
curl -sf -o /dev/null -c "$JAR" "$E2E_BASE_URL/admin/?k=$E2E_GATE_KEY"
printf '{"password":%s}' "$(printf '%s' "$E2E_ADMIN_PASSWORD" | python3 -c 'import json,sys;print(json.dumps(sys.stdin.read()))')" >"$TMP/login.json"
curl -sf -o /dev/null -b "$JAR" -c "$JAR" -H 'X-Study-Admin: 1' -H 'Content-Type: application/json' \
  --data @"$TMP/login.json" "$E2E_BASE_URL/api/admin/login"
RUN_CODE=$(curl -sf -b "$JAR" -H 'X-Study-Admin: 1' -H 'Content-Type: application/json' \
  -d "{\"name\":\"e2e security\",\"source\":\"Other\",\"configVersion\":\"${E2E_CONFIG:-test-fast-1.4.0}\"}" \
  "$E2E_BASE_URL/api/admin/runs" | python3 -c 'import json,sys;print(json.load(sys.stdin)["code"])')

cd "$E2E"
status=0
npx playwright test || status=1
echo "--- security checks against $E2E_BASE_URL"
./security.sh "$E2E_BASE_URL" "$RUN_CODE" || status=1
exit $status
