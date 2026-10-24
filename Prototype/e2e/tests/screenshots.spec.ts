// Screenshot walk-through of a running service (not part of `make e2e`: skipped unless SHOTS_RUN_CODE
// is set). One checker and one neutral participant session on an existing run, then the admin SPA.
//
//   E2E_BASE_URL=https://... E2E_GATE_KEY=... E2E_ADMIN_PASSWORD=... SHOTS_RUN_CODE=<code> \
//   SHOTS_OUT=../reports/screenshots/<date> npx playwright test screenshots.spec.ts
//
// The step driving is the one of flows.spec.ts: the screen is whatever step descriptor the server
// returned last. To catch short-lived screens (device gate, the fast test timing), every participant
// API call is routed through a gate: the request that would leave the current screen is held until the
// screenshot of that screen is taken. On a slide step, scene_started waits for the first shot (slide
// rendered, before the clip) and audio_ended waits for the "-marked" shot (all markings drawn, since
// the client sets the marking clock to the end when the clip ends). Holding only makes events later.

import { test, expect, type APIRequestContext, type Browser, type Page } from "@playwright/test";
import { readFileSync, mkdirSync, renameSync, readdirSync, rmSync } from "node:fs";
import { dirname, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const here = dirname(fileURLToPath(import.meta.url));
const root = resolve(here, "../..");
const RUN_CODE = process.env.SHOTS_RUN_CODE ?? "";
const OUT = resolve(process.env.SHOTS_OUT ?? resolve(root, "reports/screenshots/shots"));
const CONFIG_FILE = process.env.E2E_CONFIG_FILE ?? resolve(root, "content/study-config.test-fast.json");
const GATE_KEY = process.env.E2E_GATE_KEY ?? "";
const ADMIN_PASSWORD = process.env.E2E_ADMIN_PASSWORD ?? "";
const AH = { "X-Study-Admin": "1" };

type Json = any;
const config: Json = JSON.parse(readFileSync(CONFIG_FILE, "utf8"));

test.use({ viewport: { width: 1366, height: 768 } });
test.skip(!RUN_CODE, "SHOTS_RUN_CODE not set (screenshot walk-through only)");

// ---------- gates ----------
class Gates {
  private m = new Map<string, { p: Promise<void>; r: () => void }>();
  private get(k: string) {
    let g = this.m.get(k);
    if (!g) {
      let r!: () => void;
      const p = new Promise<void>((res) => (r = res));
      g = { p, r };
      this.m.set(k, g);
    }
    return g;
  }
  wait(k: string, ms = 90_000): Promise<void> {
    const g = this.get(k);
    let t: ReturnType<typeof setTimeout> | undefined;
    const timeout = new Promise<void>((res) => {
      t = setTimeout(() => {
        console.log(`gate ${k} timed out, releasing`);
        res();
      }, ms);
    });
    return Promise.race([g.p, timeout]).finally(() => clearTimeout(t));
  }
  open(k: string) {
    this.get(k).r();
  }
}
const SIDE = new Set(["marking_shown", "visibility"]);

async function adminLogin(request: APIRequestContext) {
  const gate = await request.get(`/admin/?k=${encodeURIComponent(GATE_KEY)}`, { maxRedirects: 0 });
  expect(gate.status()).toBe(302);
  const login = await request.post("/api/admin/login", { headers: AH, data: { password: ADMIN_PASSWORD } });
  expect(login.status()).toBe(200);
}
async function groupOf(request: APIRequestContext, sessionId: string): Promise<string> {
  const r = await request.get(`/api/admin/participants/${sessionId}`, { headers: AH });
  expect(r.status()).toBe(200);
  const j = await r.json();
  return j.group ?? j.session?.group ?? "";
}

/** Waits until fonts and every image on the page are loaded, then a short settle. */
async function settle(page: Page, ms = 400) {
  await page
    .waitForFunction(() => document.fonts.status === "loaded" && [...document.images].every((i) => i.complete), null, { timeout: 15_000 })
    .catch(() => undefined);
  await page.waitForTimeout(ms);
}

/** Waits until the checker overlay a step carries is drawn (marks or the panel's tick) and animated in. */
async function overlayDrawn(page: Page, step: Json) {
  if (!step.overlay) return;
  const sel = step.overlay.status === "flagged" ? "[data-marking-id]" : `[data-testid="checker-panel"] svg`;
  await page.locator(sel).first().waitFor({ state: "visible", timeout: 15_000 }).catch(() => console.log(`no marking visible on ${step.stepId}`));
  await page.waitForTimeout(1500); // draw-on animation
}

async function pick(page: Page, name: string, value: string) {
  const label = page.locator(`[data-testid="scale-${name}"] label`).filter({ has: page.locator(`input[value="${value}"]`) });
  await label.click();
  await expect(label.locator("input")).toBeChecked();
}

// ---------- one participant session ----------
async function runSession(browser: Browser, request: APIRequestContext, dir: string) {
  const ctx = await browser.newContext({ viewport: { width: 1366, height: 768 } });
  const page = await ctx.newPage();
  const gates = new Gates();
  const endPending = new Set<string>();
  let latest: Json | null = null;
  let sessionId = "";
  const files: string[] = [];

  await page.route(/\/api\/v1\//, async (route) => {
    const req = route.request();
    let type = "";
    if (req.method() === "POST" && req.url().includes("/events")) {
      try {
        type = JSON.parse(req.postData() ?? "{}").type ?? "";
      } catch {
        /* ignore */
      }
    }
    const cur = latest;
    // The landing's GET /next (after the session and the scene pack) is what mounts the device gate, so
    // on that step only the POST that leaves it (device_checked) is held.
    const landing = cur?.scene === "deviceGate" && req.method() !== "POST";
    if (cur && !SIDE.has(type) && !landing) {
      if (cur.scene === "slide" && type === "audio_ended") {
        endPending.add(cur.stepId);
        await gates.wait(`${cur.stepId}:marked`);
      } else await gates.wait(`${cur.stepId}:base`);
    }
    const resp = await route.fetch();
    const body = await resp.body();
    try {
      const j = JSON.parse(body.toString("utf8"));
      if (j.sessionId) sessionId = j.sessionId;
      const n = j.next ?? j.error?.next;
      if (n) latest = n;
    } catch {
      /* not JSON */
    }
    await route.fulfill({ response: resp, body });
  });
  page.on("pageerror", (e) => console.log(`page error: ${e.message}`));

  let nn = 0;
  const shot = async (step: Json, suffix = "") => {
    const name = `${String(nn).padStart(2, "0")}-${step.scene}-${step.sceneRef ?? step.stepId}${suffix}.png`;
    await page.screenshot({ path: resolve(dir, name), fullPage: true, animations: "allow" });
    files.push(name);
  };

  await page.goto(`/s/${RUN_CODE}`);
  let group = "";
  const plan: Record<string, string[]> = { checker: ["challenge", "explain", "nothing"], neutral: ["hasQuestion", "noQuestion", "hasQuestion"] };
  let actions = 0;
  const done = new Set<string>();
  const deadline = Date.now() + 20 * 60_000;
  const cont = page.getByRole("button", { name: config.labels.continue });

  while (Date.now() < deadline) {
    const step: Json = latest;
    if (!step || done.has(step.stepId)) {
      await page.waitForTimeout(100);
      continue;
    }
    done.add(step.stepId);
    nn++;
    const base = `${step.stepId}:base`;
    switch (step.scene) {
      case "deviceGate":
        // The route shows "Loading..." until the scene pack is in; the gate screen mounts after it.
        await expect(page.getByText("Loading...")).toBeHidden({ timeout: 30_000 });
        await settle(page, 300);
        await shot(step);
        gates.open(base);
        break;
      case "consent":
        await page.getByRole("checkbox").check();
        await expect(page.getByRole("button", { name: config.consent.acceptLabel })).toBeEnabled();
        await settle(page);
        await shot(step);
        gates.open(base);
        await page.getByRole("button", { name: config.consent.acceptLabel }).click();
        break;
      case "audioCheck":
        await expect(page.getByRole("button", { name: config.labels.playTestSound, exact: true })).toBeVisible();
        await settle(page);
        await shot(step);
        gates.open(base);
        await page.getByRole("button", { name: config.labels.playTestSound, exact: true }).click();
        await page.waitForTimeout(500);
        await page.getByRole("button", { name: config.setup.audioCheck.correctOptionId, exact: true }).click();
        break;
      case "preStart":
        group = await groupOf(request, sessionId);
        await expect(page.getByRole("button", { name: config.labels.startMeeting, exact: true })).toBeVisible();
        await settle(page);
        await shot(step);
        gates.open(base);
        await page.getByRole("button", { name: config.labels.startMeeting, exact: true }).click();
        break;
      case "example":
      case "opening":
      case "slide":
      case "question":
      case "answer": {
        await expect(page.getByTestId("boardroom")).toBeVisible();
        await settle(page, 600);
        if (step.scene !== "slide") await overlayDrawn(page, step);
        await shot(step);
        gates.open(base);
        // Gesture or problem "Continue", as in flows.spec.ts, until the step changes.
        const stepId = step.stepId;
        while (latest?.stepId === stepId && Date.now() < deadline) {
          if (step.scene === "slide" && endPending.has(stepId)) break;
          if ((await cont.isVisible().catch(() => false)) && (await cont.isEnabled().catch(() => false))) await cont.click().catch(() => undefined);
          await page.waitForTimeout(100);
        }
        if (step.scene === "slide") {
          if (step.overlay) {
            const sel = step.overlay.status === "flagged" ? "[data-marking-id]" : `[data-testid="checker-panel"] svg`;
            await page.locator(sel).first().waitFor({ state: "visible", timeout: 15_000 }).catch(() => console.log(`no marking visible on ${stepId}`));
            await page.waitForTimeout(1500); // draw-on animation
          } else await page.waitForTimeout(300);
          await shot(step, "-marked");
          gates.open(`${stepId}:marked`);
        }
        break;
      }
      case "postView":
        await expect(page.getByTestId("boardroom")).toBeVisible();
        await settle(page, 200);
        await shot(step);
        // Checker sessions keep the slide's marks and panel through the view (and the action, question
        // and answer): a second shot once the marks are drawn shows that they persist.
        if (step.overlay) {
          await overlayDrawn(page, step);
          await shot(step, "-marked");
        }
        gates.open(base);
        break;
      case "rating": {
        const submit = page.getByRole("button", { name: config.labels.ratingSubmit, exact: true });
        await expect(submit).toBeVisible();
        await settle(page);
        await shot(step);
        gates.open(base);
        await expect(submit).toBeEnabled({ timeout: 30_000 });
        for (const [item, v] of Object.entries({ trust: "7", competence: "5", benevolence: "4", integrity: "6" })) await pick(page, item, v);
        await submit.click();
        break;
      }
      case "action": {
        if (!group) group = await groupOf(request, sessionId);
        const want = plan[group][actions++];
        const opt = (step.options ?? []).find((o: Json) => o.id === want);
        if (!opt) throw new Error(`action ${want} not offered: ${JSON.stringify(step.options)}`);
        await expect(page.getByRole("button", { name: opt.label })).toBeVisible();
        await settle(page);
        await overlayDrawn(page, step);
        await shot(step);
        gates.open(base);
        await page.getByRole("button", { name: opt.label }).click();
        break;
      }
      case "budget": {
        const scene = config.scenes.find((s: Json) => s.kind === "budget");
        await expect(page.getByRole("radio", { name: scene.options[0].label })).toBeVisible();
        await settle(page);
        await shot(step);
        gates.open(base);
        await page.getByRole("radio", { name: scene.options[0].label }).check();
        await pick(page, "confidence", "8");
        await page.getByRole("button", { name: config.labels.ratingSubmit, exact: true }).click();
        break;
      }
      case "background": {
        const scene = config.scenes.find((s: Json) => s.kind === "background");
        await expect(page.getByRole("button", { name: config.labels.ratingSubmit, exact: true })).toBeVisible();
        await settle(page);
        await shot(step);
        gates.open(base);
        for (const it of scene.items) {
          if (it.kind === "single") await page.locator(`input[type=radio][name="bg-${it.id}"]`).first().check();
          else if (it.kind === "scale") await pick(page, it.id, `${it.scale.min}`);
          else if (it.kind === "country") await page.locator(`select#bg-${it.id}`).selectOption("CA");
        }
        await page.getByRole("button", { name: config.labels.ratingSubmit, exact: true }).click();
        break;
      }
      case "debrief":
        await expect(page.getByText(config.scenes.find((s: Json) => s.kind === "debrief").doNotDiscuss)).toBeVisible();
        await settle(page);
        await shot(step);
        gates.open(base);
        break;
      case "done":
        if (step.blocked) throw new Error("blocked by the device gate");
        await settle(page);
        await shot(step);
        break;
      default:
        throw new Error(`unexpected scene ${step.scene}`);
    }
    if (step.scene === "done") break;
  }
  await ctx.close();
  if (!group) group = await groupOf(request, sessionId);
  return { sessionId, group, files };
}

// ---------- the walk ----------
test("screenshot walk-through: participant sessions and admin", async ({ browser, playwright, baseURL }) => {
  test.setTimeout(60 * 60_000);
  mkdirSync(OUT, { recursive: true });
  const request = await playwright.request.newContext({ baseURL });
  await adminLogin(request);

  const got = new Map<string, string>();
  for (let i = 1; i <= 4 && got.size < 2; i++) {
    const tmp = resolve(OUT, `.tmp-s${i}`);
    rmSync(tmp, { recursive: true, force: true });
    mkdirSync(tmp, { recursive: true });
    const r = await runSession(browser, request, tmp);
    console.log(`session ${i}: ${r.sessionId} group=${r.group} shots=${r.files.length}`);
    if (!got.has(r.group) && (r.group === "checker" || r.group === "neutral")) {
      got.set(r.group, r.sessionId);
      for (const f of readdirSync(tmp)) renameSync(resolve(tmp, f), resolve(OUT, `${r.group}-${f}`));
    }
    rmSync(tmp, { recursive: true, force: true });
  }
  expect(got.size).toBe(2);

  // ---- admin SPA ----
  const ctx = await browser.newContext({ viewport: { width: 1366, height: 768 } });
  const page = await ctx.newPage();
  const ashot = (name: string) => page.screenshot({ path: resolve(OUT, `admin-${name}.png`), fullPage: true });

  await page.goto(`/admin/`);
  await settle(page);
  await ashot("01-404-without-key");

  await page.goto(`/admin/?k=${encodeURIComponent(GATE_KEY)}`);
  const pw = page.getByLabel(/password/i);
  await expect(pw).toBeVisible();
  await settle(page);
  await ashot("02-login");
  await pw.fill(ADMIN_PASSWORD);
  await page.getByRole("button", { name: /log ?in/i }).click();
  await expect(pw).toBeHidden();

  await page.getByRole("tab", { name: "Runs" }).click();
  await expect(page.getByText(RUN_CODE).first()).toBeVisible();
  await settle(page);
  await ashot("03-runs");

  await page.getByRole("tab", { name: "Participants" }).click();
  const runSelect = page.locator("select").first();
  await expect(runSelect.locator("option", { hasText: `(${RUN_CODE})` })).toHaveCount(1);
  const runLabel = (await runSelect.locator("option", { hasText: `(${RUN_CODE})` }).textContent()) ?? "";
  await runSelect.selectOption({ label: runLabel });
  for (const sid of got.values()) {
    await expect(page.getByRole("row").filter({ hasText: sid.slice(0, 8) })).toContainText(/finished/i, { timeout: 30_000 });
  }
  await settle(page);
  await ashot("04-participants");

  await page.getByRole("row").filter({ hasText: got.get("checker")!.slice(0, 8) }).click();
  const drawer = page.getByLabel("Session detail");
  await expect(drawer).toBeVisible();
  await expect(drawer.locator("table, li").first()).toBeVisible({ timeout: 15_000 }).catch(() => undefined);
  await settle(page, 800);
  await ashot("05-session-detail");
  // The timeline further down the drawer.
  await drawer.evaluate((el) => (el.scrollTop = el.scrollHeight / 2));
  await page.waitForTimeout(400);
  await ashot("06-session-detail-timeline");
  await page.keyboard.press("Escape");

  await page.getByRole("tab", { name: "Export" }).click();
  await settle(page);
  await ashot("07-export");

  await page.getByRole("tab", { name: "Configs" }).click();
  await settle(page, 800);
  await ashot("08-configs");
  await ctx.close();
});
