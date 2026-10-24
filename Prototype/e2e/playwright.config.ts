import { defineConfig } from "@playwright/test";

// Firefox only, one worker: the flows share one run and the server's round-robin allocation.
// E2E_BASE_URL points at a running API (default http://localhost:8080); see README.md.
const fast = (process.env.E2E_CONFIG ?? "test-fast-1.4.0").startsWith("test");

export default defineConfig({
  testDir: "./tests",
  workers: 1,
  fullyParallel: false,
  retries: 0,
  // With the real 1.4.0 timing a session takes about 15 minutes (estimated clip lengths, 10 s views).
  timeout: fast ? 10 * 60_000 : 90 * 60_000,
  reporter: [["list"]],
  use: {
    baseURL: process.env.E2E_BASE_URL ?? "http://localhost:8080",
    browserName: "firefox",
    headless: true,
    viewport: { width: 1440, height: 900 },
    launchOptions: {
      // Silent: audio is decoded and plays (so `ended` fires and the timing logic runs) but nothing
      // reaches the speakers. Chromium, if ever added, needs args: ["--mute-audio"].
      firefoxUserPrefs: {
        "media.autoplay.default": 0,
        "media.autoplay.blocking_policy": 0,
        "media.volume_scale": "0.0",
        "media.default_volume": "0.0",
      },
    },
  },
  projects: [{ name: "firefox", use: { browserName: "firefox", headless: true } }],
});
