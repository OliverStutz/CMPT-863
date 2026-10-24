# Load tests (reports/concurrency-review.md). Run from load/.
# `make load` runs against a server you started; the other targets start one
# with throwaway admin secrets (run-local.sh) and stop it afterwards.
BASE ?= http://localhost:8080
SESSIONS ?= 50
RAMP ?= 5
THINK ?= 1
ARGS ?=

.PHONY: load load-memory load-emulator load-burst build vet docsize

## load: drive SESSIONS sessions against BASE (needs LOAD_GATE_KEY and LOAD_ADMIN_PASSWORD for the admin check)
load:
	go run . -base $(BASE) -sessions $(SESSIONS) -ramp $(RAMP) -think $(THINK) $(ARGS)

## load-memory: own server on the memory store, 50 sessions
load-memory:
	./run-local.sh memory $(SESSIONS) $(RAMP) -think $(THINK) $(ARGS)

## load-emulator: own server on the Firestore emulator (docker compose up firestore), 50 sessions
load-emulator:
	./run-local.sh firestore $(SESSIONS) $(RAMP) -think $(THINK) $(ARGS)

## load-burst: 200 sessions started within 10 s on the emulator
load-burst:
	./run-local.sh firestore 200 10 -think $(THINK) $(ARGS)

## docsize: document counts and sizes in the emulator (PROJECT=<id>)
docsize:
	go run ./cmd/docsize -project $(PROJECT)

build:
	go build ./...

vet:
	go vet ./...
