# syntax=docker/dockerfile:1
# Study API image: one Go service serving both SPAs (go:embed), the content
# assets (read from the private GCS bucket at runtime) and both APIs.
# Architecture sections 2, 3 and 12.

########## Stage 1: build the two React apps ##########
FROM node:26-bookworm-slim AS web
WORKDIR /src
# The whole context (minus .dockerignore) so that apps/, packages/shared and a
# root package.json / lockfile are all available when they exist.
COPY . .
# Each app builds into ../../services/api/web/<app> (its vite outDir).
# - missing package-lock.json: fall back to npm install
# - missing app directory: write a placeholder index.html so the Go embed
#   directives still find a file and the image builds before the apps land.
RUN set -eu; \
    for app in participant admin; do \
      out="/src/services/api/web/${app}"; \
      if [ -f "apps/${app}/package.json" ]; then \
        cd "/src/apps/${app}"; \
        if [ -f package-lock.json ]; then npm ci; else npm install --no-audit --no-fund; fi; \
        npm run build; \
        cd /src; \
        if [ ! -f "${out}/index.html" ] && [ -f "apps/${app}/dist/index.html" ]; then \
          mkdir -p "${out}" && cp -r "apps/${app}/dist/." "${out}/"; \
        fi; \
      else \
        echo "apps/${app} not present, writing placeholder"; \
      fi; \
      if [ ! -f "${out}/index.html" ]; then \
        mkdir -p "${out}"; \
        printf '<!doctype html><meta charset="utf-8"><title>%s</title><p>%s app not built yet.</p>\n' "${app}" "${app}" > "${out}/index.html"; \
      fi; \
    done

########## Stage 2: build the Go binaries ##########
FROM golang:1.27-bookworm AS go
WORKDIR /src/services/api
ENV CGO_ENABLED=0 GOOS=linux GOWORK=off
COPY services/api/go.mod services/api/go.sum* ./
RUN go mod download
COPY services/api/ ./
# The embed directories must hold the built SPAs before go build.
RUN rm -rf web/participant web/admin
COPY --from=web /src/services/api/web/participant ./web/participant
COPY --from=web /src/services/api/web/admin ./web/admin
RUN go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api \
 && go build -trimpath -ldflags="-s -w" -o /out/sweeper ./cmd/sweeper

########## Stage 3: runtime ##########
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=go /out/api /api
# The sweeper Cloud Run job uses this same image with command /sweeper.
COPY --from=go /out/sweeper /sweeper
ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/api"]
