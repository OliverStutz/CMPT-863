import { defineConfig } from "vitest/config";
import react from "@vitejs/plugin-react";

// Built into the Go service, which embeds services/api/web/participant and serves it at / and /s/*
// ("02 - Architecture.md" sections 3 and 9.1).
export default defineConfig({
  base: "/",
  plugins: [react()],
  build: {
    outDir: "../../services/api/web/participant",
    emptyOutDir: true,
    // DECISION: bundle files go under /static/, not Vite's default /assets/, because /assets/* is the
    // content asset path (consent PDF, images) in the integration brief.
    assetsDir: "static",
  },
  server: {
    port: 5173,
    strictPort: true,
    // DECISION: section 2 serves assets at /a/*, the integration brief at /assets/*. Both are proxied;
    // the app itself uses the assetBaseUrl the server returns, plus the fixed /assets paths in src/assets.ts.
    proxy: {
      "/api": { target: "http://localhost:8080", changeOrigin: false },
      "/assets": { target: "http://localhost:8080", changeOrigin: false },
      "/a/": { target: "http://localhost:8080", changeOrigin: false },
    },
  },
  test: {
    environment: "jsdom",
    include: ["src/**/*.test.{ts,tsx}"],
    setupFiles: ["src/test/setup.ts"],
    css: { modules: { classNameStrategy: "non-scoped" } },
  },
});
