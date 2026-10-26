import { defineConfig } from "vitest/config";
import react from "@vitejs/plugin-react";

// Built into the Go service, which embeds services/api/web/admin and serves it under /admin/ (section 3).
export default defineConfig({
  base: "/admin/",
  plugins: [react()],
  build: {
    outDir: "../../services/api/web/admin",
    emptyOutDir: true,
  },
  server: {
    port: 5174,
    strictPort: true,
    proxy: {
      "/api": { target: "http://localhost:8080", changeOrigin: false },
    },
  },
  test: {
    environment: "node",
    include: ["src/**/*.test.ts", "src/**/*.test.tsx"],
  },
});
