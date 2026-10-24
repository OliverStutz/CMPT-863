# Monorepo targets (02 - Architecture.md sections 13 and 14).
# The Go service lives in services/api; the SPAs in apps/participant and apps/admin.

API_DIR        := services/api
WEB_DIR        := $(API_DIR)/web
EMULATOR_HOST  := localhost:8081
DEV_ENV := FIRESTORE_EMULATOR_HOST=$(EMULATOR_HOST) FIRESTORE_PROJECT_ID=demo-study \
	CONTENT_DIR=../../content TRUSTED_PROXY_HOPS=0 INSECURE_COOKIES=true \
	ADMIN_PASSWORD=dev ADMIN_ACCESS_KEY_VALUE=dev-gate-key ADMIN_SESSION_KEY_VALUE=dev-session-key \
	SEED_RUN_CODE=devrun ALLOW_TEST_CONFIGS=true

.PHONY: dev dev-memory emulator test test-emulator build build-web lint sweeper admin-password admin-key clean

## dev: start the Firestore emulator, then run the API on :8080 against it
dev: emulator
	cd $(API_DIR) && $(DEV_ENV) go run ./cmd/api

## dev-memory: run the API with the in-memory store (no Docker needed)
dev-memory:
	cd $(API_DIR) && $(DEV_ENV) STORE=memory FIRESTORE_EMULATOR_HOST= go run ./cmd/api

emulator:
	docker compose up -d --wait firestore

## test: unit and HTTP tests (the Firestore store test is skipped without the emulator)
test:
	cd $(API_DIR) && go test ./...

## test-emulator: all tests, including the Firestore store against the emulator
test-emulator: emulator
	cd $(API_DIR) && FIRESTORE_EMULATOR_HOST=$(EMULATOR_HOST) go test ./...

## build: the api and sweeper binaries into services/api/bin
build:
	cd $(API_DIR) && CGO_ENABLED=0 go build -trimpath -o bin/api ./cmd/api && CGO_ENABLED=0 go build -trimpath -o bin/sweeper ./cmd/sweeper

## build-web: build both SPAs into services/api/web/<app> for go:embed (the Vite
## configs write there directly; a dist/ directory, if an app produces one, is copied)
build-web:
	@for app in participant admin; do \
		if [ -f apps/$$app/package.json ]; then \
			echo "building apps/$$app"; \
			(cd apps/$$app && npm ci && npm run build) || exit 1; \
			if [ -d apps/$$app/dist ]; then \
				rm -rf $(WEB_DIR)/$$app && mkdir -p $(WEB_DIR)/$$app && cp -R apps/$$app/dist/. $(WEB_DIR)/$$app/; \
			fi; \
		else \
			echo "apps/$$app not found, skipping (placeholder page stays embedded)"; \
		fi; \
	done

## lint: go vet
lint:
	cd $(API_DIR) && go vet ./...

## sweeper: run one abandonment sweep against the emulator
sweeper:
	cd $(API_DIR) && $(DEV_ENV) go run ./cmd/sweeper

## admin-password: read a password from stdin and print its bcrypt hash
admin-password:
	@cd $(API_DIR) && go run ./cmd/admintool password

## admin-key: print a random 32-byte base64url key (admin gate key or session key)
admin-key:
	@cd $(API_DIR) && go run ./cmd/admintool key

clean:
	rm -rf $(API_DIR)/bin
