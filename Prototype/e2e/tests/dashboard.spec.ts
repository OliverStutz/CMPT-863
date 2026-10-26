// Screenshots of the run dashboard of a running service (skipped unless DASH_RUN_CODE is set).
//
//   E2E_BASE_URL=https://... E2E_GATE_KEY=... E2E_ADMIN_PASSWORD=... DASH_RUN_CODE=<run code> \
//   SHOTS_OUT=<dir> npx playwright test tests/dashboard.spec.ts
//
// Logs in through the admin UI, opens the dashboard of the run with that code (hash route #run=<id>),
// and saves full-page screenshots in light and dark mode, plus the table view. Nothing is changed.

import { mkdirSync } from "node:fs";
import { resolve } from "node:path";
import { expect, test, type Page } from "@playwright/test";

const RUN_CODE = process.env.DASH_RUN_CODE ?? "";
const OUT = resolve(process.env.SHOTS_OUT ?? resolve(__dirname, "..", "..", "reports", "screenshots", "dashboard"));
const GATE_KEY = process.env.E2E_GATE_KEY ?? "";
const ADMIN_PASSWORD = process.env.E2E_ADMIN_PASSWORD ?? "";
const AH = { "X-Study-Admin": "1" };

test.skip(!RUN_CODE, "DASH_RUN_CODE not set (dashboard screenshots only)");

async function settle(page: Page, ms = 800) {
  await page
    .waitForFunction(() => document.fonts.status === "loaded", null, { timeout: 15_000 })
    .catch(() => undefined);
  await page.waitForTimeout(ms);
}

test("run dashboard screenshots", async ({ browser, playwright, baseURL }) => {
  test.setTimeout(5 * 60_000);
  mkdirSync(OUT, { recursive: true });
  const request = await playwright.request.newContext({ baseURL: baseURL! });
  expect((await request.get(`/admin/?k=${encodeURIComponent(GATE_KEY)}`, { maxRedirects: 0 })).status()).toBe(302);
  expect((await request.post("/api/admin/login", { headers: AH, data: { password: ADMIN_PASSWORD } })).status()).toBe(200);
  const runs = (await (await request.get("/api/admin/runs", { headers: AH })).json()) as { id: string; code: string }[];
  const run = runs.find((r) => r.code === RUN_CODE);
  expect(run, `run ${RUN_CODE} exists`).toBeTruthy();

  for (const scheme of ["light", "dark"] as const) {
    const ctx = await browser.newContext({ viewport: { width: 1440, height: 900 }, colorScheme: scheme });
    const page = await ctx.newPage();
    await page.goto(`/admin/?k=${encodeURIComponent(GATE_KEY)}`);
    await page.getByLabel(/password/i).fill(ADMIN_PASSWORD);
    await page.getByRole("button", { name: /log in/i }).click();
    await page.getByRole("tab", { name: /runs/i }).waitFor({ state: "visible", timeout: 20_000 });
    await page.goto(`/admin/#run=${encodeURIComponent(run!.id)}`);
    await page.getByText(/back to runs/i).first().waitFor({ state: "visible", timeout: 30_000 });
    // Wait for the analysis to load: the KPI tiles appear once the document arrived.
    await page.locator("svg").first().waitFor({ state: "visible", timeout: 30_000 }).catch(() => undefined);
    await settle(page, 1500);
    await page.screenshot({ path: resolve(OUT, `dashboard-${scheme}.png`), fullPage: true });
    if (scheme === "light") {
      const table = page.getByLabel(/table view/i).or(page.getByText(/table view/i)).first();
      if (await table.count()) {
        await table.click();
        await settle(page, 800);
        await page.screenshot({ path: resolve(OUT, `dashboard-${scheme}-tables.png`), fullPage: true });
      }
    }
    await ctx.close();
  }
  await request.dispose();
});
