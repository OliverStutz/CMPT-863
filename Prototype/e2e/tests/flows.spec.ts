// End-to-end flows against the real Go service (README.md): one checker and one neutral session in
// headless Firefox, then the admin view, the export and a leak check of every body a participant
// session received. Every participant action is chosen from the step descriptor the server sent last.

import { test, expect, type APIRequestContext, type Browser, type Page } from "@playwright/test";
import { readFileSync, mkdirSync, writeFileSync } from "node:fs";
import { dirname, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const here = dirname(fileURLToPath(import.meta.url));
const root = resolve(here, "../..");
const OUT = resolve(here, "../test-results/bodies");

const CONFIG_VERSION = process.env.E2E_CONFIG ?? "test-fast-1.4.0";
const CONFIG_FILE = process.env.E2E_CONFIG_FILE ?? resolve(root, "content/study-config.v1.4.json");
const GATE_KEY = need("E2E_GATE_KEY");
const ADMIN_PASSWORD = need("E2E_ADMIN_PASSWORD");

function need(name: string): string {
  const v = process.env[name];
  if (!v) throw new Error(`${name} is not set (see e2e/README.md)`);
  return v;
}

// ---------- config texts (for the leak check) ----------
type Json = any;
const config: Json = JSON.parse(readFileSync(CONFIG_FILE, "utf8"));
const slides: Json[] = config.scenes.filter((s: Json) => s.kind === "slide");
function textsOf(v: Json): string[] {
  if (!v || typeof v !== "object") return [];
  const out: string[] = [];
  if (typeof v.text === "string") out.push(v.text);
  for (const [k, x] of Object.entries(v)) if (k !== "text" && !k.startsWith("_")) out.push(...textsOf(x));
  return out;
}
/**
 * slideId -> the texts that tell the two answer variants apart. "whole" answers: the reason and denial
 * texts. "joined" answers share opening and closing; the reason adds reasonBody, the denial adds
 * denialBody (or nothing, when padDenialToReason pads with silence).
 */
const answerTexts = new Map<string, { reason: string[]; denial: string[]; mode: string }>();
for (const s of slides) {
  const a = s.answers;
  if (!a) continue;
  if (a.mode === "joined") answerTexts.set(s.id, { mode: a.mode, reason: textsOf(a.reasonBody), denial: textsOf(a.denialBody) });
  else answerTexts.set(s.id, { mode: a.mode, reason: textsOf(a.reason), denial: textsOf(a.denial) });
}
/** Which variant a session heard for a slide, from the transcript of its answer step. */
function playedVariant(res: { steps: Json[] }, slideId: string): "reason" | "denial" | null {
  const t = answerTexts.get(slideId)!;
  const tr = res.steps.filter((s) => s.scene === "answer" && s.sceneRef === `${slideId}.answer` && s.transcript).map((s) => s.transcript as string);
  if (!tr.length) return null;
  const has = (xs: string[]) => xs.length > 0 && xs.every((x) => tr.some((y) => y.includes(x)));
  if (has(t.reason) && !has(t.denial)) return "reason";
  if (has(t.denial) && !has(t.reason)) return "denial";
  if (t.denial.length === 0 && !has(t.reason)) return "denial";
  return null;
}

/** Flagged slides: the checker marks stay on screen from the slide until the next slide begins. */
const flaggedSlides = new Set<string>(slides.filter((s: Json) => s.checker?.status === "flagged").map((s: Json) => s.id));
/** The slide whose marks a post-view, action, question or answer step must still show, if any. */
function persistSlideOf(step: Json): string | null {
  if (!["postView", "action", "question", "answer"].includes(step.scene) || !step.sceneRef) return null;
  const slideId = String(step.sceneRef).split(".")[0];
  return slides.some((s: Json) => s.id === slideId) ? slideId : null;
}
interface PersistCheck {
  stepId: string;
  scene: string;
  slideId: string;
  flagged: boolean;
  marks: number;
  panel: boolean;
}

// ---------- admin API ----------
async function adminContext(request: APIRequestContext) {
  const gate = await request.get(`/admin/?k=${encodeURIComponent(GATE_KEY)}`, { maxRedirects: 0 });
  expect(gate.status()).toBe(302);
  const login = await request.post("/api/admin/login", {
    headers: { "X-Study-Admin": "1" },
    data: { password: ADMIN_PASSWORD },
  });
  expect(login.status()).toBe(200);
}
const AH = { "X-Study-Admin": "1" };

// ---------- one participant session ----------
interface Captured {
  url: string;
  status: number;
  body: string;
}
interface SessionResult {
  sessionId: string;
  group: string;
  bodies: Captured[];
  steps: Json[];
  actions: string[];
  transcripts: string[];
  persist: PersistCheck[];
}

async function runSession(browser: Browser, request: APIRequestContext, runCode: string, label: string): Promise<SessionResult> {
  const ctx = await browser.newContext();
  const page = await ctx.newPage();
  const bodies: Captured[] = [];
  const steps: Json[] = [];
  let latest: Json | null = null;
  let sessionId = "";
  const pending: Promise<void>[] = [];

  page.on("response", (res) => {
    const url = new URL(res.url());
    // The JS and CSS bundle is the same for everyone (4.3); everything else the server sent is kept.
    if (url.pathname.startsWith("/static/")) return;
    pending.push(
      (async () => {
        let body = "";
        try {
          body = (await res.body()).toString("utf8");
        } catch {
          return;
        }
        if ((res.headers()["content-type"] ?? "").includes("audio")) body = "";
        bodies.push({ url: url.pathname + url.search, status: res.status(), body });
        if (!url.pathname.startsWith("/api/v1/")) return;
        let j: Json;
        try {
          j = JSON.parse(body);
        } catch {
          return;
        }
        if (j.sessionId) sessionId = j.sessionId;
        const next = j.next ?? j.error?.next;
        if (next) {
          latest = next;
          steps.push(next);
        }
      })(),
    );
  });
  page.on("pageerror", (e) => console.log(`[${label}] page error: ${e.message}`));

  // Persistence of the checker marks: on the first request a post-view, action, question or answer step
  // sends (or the request that leaves it), that step is still on screen, so the request is held while
  // the page is inspected. Checker sessions must still show the slide's marked segments (flagged slides)
  // or the panel (unflagged slides); neutral sessions never show either.
  const persist: PersistCheck[] = [];
  const inspected = new Set<string>();
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
    await Promise.all([...pending]);
    const cur: Json = latest;
    const slideId = cur ? persistSlideOf(cur) : null;
    if (cur && slideId && type !== "marking_shown" && type !== "visibility" && !inspected.has(cur.stepId)) {
      inspected.add(cur.stepId);
      const flagged = flaggedSlides.has(slideId);
      const marks = page.locator("[data-marking-id]");
      const panel = page.getByTestId("checker-panel");
      if (cur.overlay) {
        const target = flagged ? marks.first() : panel;
        await target.waitFor({ state: "visible", timeout: 5_000 }).catch(() => undefined);
      } else {
        await page.waitForTimeout(250);
      }
      persist.push({
        stepId: cur.stepId,
        scene: cur.scene,
        slideId,
        flagged,
        marks: await marks.count(),
        panel: await panel.isVisible().catch(() => false),
      });
    }
    await route.continue();
  });

  await page.goto(`/s/${runCode}`);

  let group = "";
  const actionPlan: Record<string, string[]> = {
    checker: ["challenge", "explain", "nothing"],
    neutral: ["hasQuestion", "noQuestion", "hasQuestion"],
  };
  const actions: string[] = [];
  const handled = new Set<string>();
  const deadline = Date.now() + (CONFIG_VERSION.startsWith("test") ? 8 * 60_000 : 85 * 60_000);
  let finishedSeen = false;

  while (Date.now() < deadline) {
    await Promise.all(pending.splice(0));
    const step: Json = latest;
    if (!step) {
      await page.waitForTimeout(200);
      continue;
    }
    const key = `${step.stepId}:${step.scene}`;
    if (step.scene === "done") {
      if (step.blocked) throw new Error("session was blocked by the device gate");
      finishedSeen = true;
      break;
    }
    if (handled.has(key) && !["example", "opening", "slide", "question", "answer"].includes(step.scene)) {
      await page.waitForTimeout(200);
      continue;
    }
    switch (step.scene) {
      case "deviceGate":
      case "postView":
      case "debrief":
        // The app sends device_checked, waits out the view and sends finished on its own.
        break;
      case "consent": {
        const box = page.getByRole("checkbox");
        await box.waitFor();
        const before = await box.boundingBox();
        // With the consent PDF present the viewer frame replaces the placeholder in the same area: the
        // checkbox does not move, stays inside the window, and nothing covers it.
        const pdfHead = await request.head(`/assets/consent/${encodeURIComponent(config.consent.pdfAssetKey)}`);
        if (pdfHead.ok()) {
          await expect(page.getByTestId("consent-pdf")).toBeVisible();
          const after = await box.boundingBox();
          expect(after, "consent checkbox did not move when the PDF frame appeared").toEqual(before);
          const vp = page.viewportSize()!;
          expect(after!.y + after!.height, "consent checkbox inside the window").toBeLessThanOrEqual(vp.height);
          const btn = await page.getByRole("button", { name: config.consent.acceptLabel }).boundingBox();
          expect(btn!.y + btn!.height, "accept button inside the window").toBeLessThanOrEqual(vp.height);
          const hit = await box.evaluate((el) => {
            const r = el.getBoundingClientRect();
            return document.elementFromPoint(r.x + r.width / 2, r.y + r.height / 2) === el;
          });
          expect(hit, "nothing covers the consent checkbox").toBe(true);
        }
        await box.check();
        await page.getByRole("button", { name: config.consent.acceptLabel }).click();
        handled.add(key);
        break;
      }
      case "audioCheck": {
        await page.getByRole("button", { name: config.labels.playTestSound, exact: true }).click();
        await page.waitForTimeout(500);
        await page.getByRole("button", { name: config.setup.audioCheck.correctOptionId, exact: true }).click();
        handled.add(key);
        break;
      }
      case "preStart": {
        // Group is learned from the admin API (the browser is never told).
        group = await groupOf(request, sessionId);
        await page.getByRole("button", { name: config.labels.startMeeting, exact: true }).click();
        handled.add(key);
        break;
      }
      case "example":
      case "opening":
      case "slide":
      case "question":
      case "answer": {
        if (step.transcript) handled.add(`t:${step.stepId}`);
        const cont = page.getByRole("button", { name: config.labels.continue });
        if (await cont.isVisible().catch(() => false)) {
          if (await cont.isEnabled().catch(() => false)) await cont.click().catch(() => undefined);
        }
        break;
      }
      case "rating": {
        const submit = page.getByRole("button", { name: config.labels.ratingSubmit, exact: true });
        await expect(submit).toBeEnabled({ timeout: 30_000 });
        const vals: Record<string, string> = { trust: "7", competence: "5", benevolence: "4", integrity: "6" };
        for (const [item, v] of Object.entries(vals)) {
          await pick(page, `${item}`, `${v}`);
        }
        await submit.click();
        handled.add(key);
        break;
      }
      case "action": {
        if (!group) group = await groupOf(request, sessionId);
        const want = actionPlan[group][actions.length];
        const opt = (step.options ?? []).find((o: Json) => o.id === want);
        if (!opt) throw new Error(`action ${want} not offered: ${JSON.stringify(step.options)}`);
        await page.getByRole("button", { name: opt.label }).click();
        actions.push(want);
        handled.add(key);
        break;
      }
      case "budget": {
        const scene = config.scenes.find((s: Json) => s.kind === "budget");
        await page.getByRole("radio", { name: scene.options[0].label }).check();
        await pick(page, `confidence`, `8`);
        await page.getByRole("button", { name: config.labels.ratingSubmit, exact: true }).click();
        handled.add(key);
        break;
      }
      case "background": {
        const scene = config.scenes.find((s: Json) => s.kind === "background");
        for (const it of scene.items) {
          if (it.kind === "single") {
            await page.locator(`input[type=radio][name="bg-${it.id}"]`).first().check();
          } else if (it.kind === "scale") {
            await pick(page, `${it.id}`, `${it.scale.min}`);
          } else if (it.kind === "country") {
            await page.locator(`select#bg-${it.id}`).selectOption("CA");
          }
        }
        await page.getByRole("button", { name: config.labels.ratingSubmit, exact: true }).click();
        handled.add(key);
        break;
      }
      default:
        throw new Error(`unexpected scene ${step.scene}`);
    }
    await page.waitForTimeout(150);
  }
  await Promise.all(pending.splice(0));
  expect(finishedSeen, `${label}: session reached "done"`).toBe(true);
  await expect(page.getByText(config.scenes.find((s: Json) => s.kind === "debrief").doNotDiscuss)).toBeVisible();
  const transcripts = steps.map((s) => s.transcript).filter((t): t is string => typeof t === "string");
  await ctx.close();
  mkdirSync(OUT, { recursive: true });
  writeFileSync(resolve(OUT, `${label}-${group}.json`), JSON.stringify({ sessionId, group, actions, bodies }, null, 1));
  return { sessionId, group, bodies, steps, actions, transcripts, persist };
}

/** Clicks one point of a ScaleRow (the radio is visually hidden; its label is the target). */
async function pick(page: Page, name: string, value: string) {
  const label = page.locator(`[data-testid="scale-${name}"] label`).filter({ has: page.locator(`input[value="${value}"]`) });
  await label.click();
  await expect(label.locator("input")).toBeChecked();
}

async function groupOf(request: APIRequestContext, sessionId: string): Promise<string> {
  const r = await request.get(`/api/admin/participants/${sessionId}`, { headers: AH });
  expect(r.status()).toBe(200);
  const j = await r.json();
  const g = j.group ?? j.session?.group;
  expect(["checker", "neutral"]).toContain(g);
  return g;
}

// ---------- leak check ----------
interface Hit {
  url: string;
  path: string;
  term: string;
  text: string;
}
const TERMS: [string, RegExp][] = [
  ["denial", /denial/i],
  ["checker", /checker/i],
  ["neutral", /neutral/i],
  ["denialSlide", /denialSlide/],
  ["group", /\bgroups?\b/i],
];
function walk(v: Json, path: string, visit: (path: string, s: string, isKey: boolean) => void) {
  if (typeof v === "string") visit(path, v, false);
  else if (Array.isArray(v)) v.forEach((x, i) => walk(x, `${path}[${i}]`, visit));
  else if (v && typeof v === "object")
    for (const [k, x] of Object.entries(v)) {
      visit(`${path}.${k}`, k, true);
      walk(x, `${path}.${k}`, visit);
    }
}
function scan(res: SessionResult): Hit[] {
  const hits: Hit[] = [];
  for (const b of res.bodies) {
    let j: Json;
    try {
      j = JSON.parse(b.body);
    } catch {
      for (const [term, re] of TERMS) if (re.test(b.body)) hits.push({ url: b.url, path: "(raw)", term, text: b.body.slice(0, 120) });
      continue;
    }
    walk(j, "$", (path, s, isKey) => {
      for (const [term, re] of TERMS) if (re.test(s)) hits.push({ url: b.url, path: isKey ? `${path} (key)` : path, term, text: s });
    });
  }
  return hits;
}
/** Occurrences that are the participant's own, shown content (asserted exactly in the test). */
function allowed(h: Hit, res: SessionResult): boolean {
  // Story wording of the opening ("a large hospital group"), the same for both groups.
  if (h.term === "group" && /hospital group/.test(h.text) && !h.path.endsWith("(key)")) return true;
  if (h.path.includes(".debrief.paragraphs[")) return h.term === "checker" || h.term === "group";
  if (res.group !== "checker" || h.term !== "checker") return false;
  if (/\.preStartText\[\d+\]$/.test(h.path)) return true;
  if (/\.options\[\d+\]\.label$/.test(h.path)) return /challenge|checker/i.test(h.text);
  if (/\.overlay\.title$/.test(h.path)) return true;
  if (/\.transcript$/.test(h.path)) return true; // the challenge question the participant chose to ask
  return false;
}

// ---------- the test ----------
test("checker and neutral sessions, admin view, export, leak check", async ({ browser, playwright, baseURL }) => {
  const request = await playwright.request.newContext({ baseURL });
  await adminContext(request);

  const runRes = await request.post("/api/admin/runs", {
    headers: AH,
    data: { name: `e2e ${new Date().toISOString()}`, source: "Other", configVersion: CONFIG_VERSION },
  });
  expect(runRes.status()).toBe(201);
  const run = await runRes.json();

  const results: SessionResult[] = [];
  for (let i = 1; i <= 4 && !(results.some((r) => r.group === "checker") && results.some((r) => r.group === "neutral")); i++) {
    const r = await runSession(browser, request, run.code, `s${i}`);
    console.log(`session ${i}: ${r.sessionId} group=${r.group} actions=${r.actions.join(",")}`);
    results.push(r);
  }
  const checker = results.find((r) => r.group === "checker")!;
  const neutral = results.find((r) => r.group === "neutral")!;
  expect(checker, "a checker session").toBeTruthy();
  expect(neutral, "a neutral session").toBeTruthy();

  // ---- checker marks persist after the slide (checker only) ----
  for (const res of [checker, neutral]) {
    writeFileSync(resolve(OUT, `persist-${res.group}.json`), JSON.stringify(res.persist, null, 1));
    if (res.group === "neutral") {
      for (const c of res.persist) {
        expect(c.marks, `neutral ${c.scene} ${c.stepId} (${c.slideId}): no marked segment`).toBe(0);
        expect(c.panel, `neutral ${c.scene} ${c.stepId} (${c.slideId}): no checker panel`).toBe(false);
      }
      continue;
    }
    for (const c of res.persist) {
      if (c.flagged) expect(c.marks, `checker ${c.scene} ${c.stepId} (${c.slideId}): marked segment still highlighted`).toBeGreaterThan(0);
      expect(c.panel, `checker ${c.scene} ${c.stepId} (${c.slideId}): checker panel still shown`).toBe(true);
    }
    for (const slideId of flaggedSlides) {
      const seen = new Set(res.persist.filter((c) => c.slideId === slideId).map((c) => c.scene));
      for (const scene of ["postView", "action", "question", "answer"]) {
        expect(seen.has(scene), `checker ${slideId}: ${scene} step inspected`).toBe(true);
      }
    }
  }
  // The neutral session was inspected on the same steps.
  for (const slideId of flaggedSlides) {
    expect(neutral.persist.some((c) => c.slideId === slideId && c.scene === "action"), `neutral ${slideId}: action step inspected`).toBe(true);
  }

  // ---- leak check over every body each session received ----
  const report: string[] = [];
  const leaksBy = new Map<SessionResult, Hit[]>();
  for (const res of results) {
    const hits = scan(res);
    for (const h of hits) report.push(`${res.group} ${allowed(h, res) ? "allowed" : "LEAK"} ${h.term} ${h.url} ${h.path}: ${h.text.slice(0, 100)}`);
    leaksBy.set(res, hits.filter((h) => !allowed(h, res)));
  }
  mkdirSync(OUT, { recursive: true });
  writeFileSync(resolve(OUT, "leak-report.txt"), report.join("\n") + "\n");
  console.log(report.join("\n"));
  for (const res of results) {
    expect(leaksBy.get(res), `${res.group}: forbidden terms`).toEqual([]);

    const all = res.bodies.map((b) => b.body).join("\n");
    expect(all.includes(ADMIN_PASSWORD), "admin password in a participant body").toBe(false);
    // Answer texts: per slide the played variant only; the other variant never.
    for (const [slideId, t] of answerTexts) {
      const played = playedVariant(res, slideId);
      expect(played, `${res.group} ${slideId}: exactly one answer variant played`).not.toBeNull();
      const unplayed = played === "reason" ? t.denial : t.reason;
      for (const x of unplayed) expect(all.includes(x), `${res.group} ${slideId}: unplayed answer text received`).toBe(false);
    }
  }

  // ---- admin: participants, detail, export ----
  const list = await (await request.get(`/api/admin/participants?runId=${run.id}`, { headers: AH })).json();
  for (const res of [checker, neutral]) {
    const row = list.items.find((x: Json) => x.sessionId === res.sessionId);
    expect(row, `participant row ${res.sessionId}`).toBeTruthy();
    expect(row.status.toLowerCase()).toBe("finished");
    expect(row.progress.index).toBe(row.progress.total);
    expect(row.group).toBe(res.group);
    expect(row.finishedAt).toBeTruthy();
    expect(row.consentAt).toBeTruthy();
    expect(Array.isArray(row.flags)).toBe(true);
    const detail = await (await request.get(`/api/admin/participants/${res.sessionId}`, { headers: AH })).json();
    const events = detail.events ?? detail.timeline ?? [];
    expect(events.length).toBeGreaterThan(30);
  }
  // Elapsed is frozen: finishedAt does not move between two polls.
  await new Promise((r) => setTimeout(r, 1500));
  const list2 = await (await request.get(`/api/admin/participants?runId=${run.id}`, { headers: AH })).json();
  for (const res of [checker, neutral]) {
    const a = list.items.find((x: Json) => x.sessionId === res.sessionId);
    const b = list2.items.find((x: Json) => x.sessionId === res.sessionId);
    expect(b.finishedAt).toBe(a.finishedAt);
  }

  const exp = await request.get(`/api/admin/export?format=csv&runId=${run.id}`, { headers: AH });
  expect(exp.status()).toBe(200);
  const csv = await exp.text();
  writeFileSync(resolve(OUT, "export.csv"), csv);
  const lines = csv.trim().split(/\r?\n/);
  const header = parseCsvLine(lines[0]);
  const pointCol = header.indexOf("point");
  const sidCol = header.indexOf("sessionId");
  expect(pointCol).toBeGreaterThanOrEqual(0);
  const points: string[] = config.scenes.filter((s: Json) => s.kind === "rating").map((s: Json) => s.point);
  for (const res of [checker, neutral]) {
    const rows = lines.slice(1).map(parseCsvLine).filter((r) => r[sidCol] === res.sessionId);
    expect(rows.map((r) => r[pointCol]).sort()).toEqual([...points].sort());
  }
  for (const res of [checker, neutral]) {
    for (const [slideId, t] of answerTexts) {
      const played = playedVariant(res, slideId);
      for (const x of played === "reason" ? t.denial : t.reason) expect(csv.includes(x)).toBe(false);
    }
    // Exactly one denial slide per session, and the export's answerType agrees with what was played.
    const rows = lines.slice(1).map(parseCsvLine).filter((r) => r[sidCol] === res.sessionId);
    const atCol = header.indexOf("answerType");
    const slCol = header.indexOf("slideId");
    for (const [slideId] of answerTexts) {
      const types = new Set(rows.filter((r) => r[slCol] === slideId).map((r) => r[atCol]).filter(Boolean));
      expect([...types]).toEqual([playedVariant(res, slideId)]);
    }
  }
  for (const [, t] of answerTexts) for (const x of t.denial) expect(csv.includes(x), "denial text in export").toBe(false);

  // ---- admin SPA smoke: the participants table shows both sessions as Finished ----
  const page = await browser.newPage();
  await page.goto(`/admin/?k=${encodeURIComponent(GATE_KEY)}`);
  await page.getByLabel(/password/i).fill(ADMIN_PASSWORD);
  await page.getByRole("button", { name: /log ?in|sign ?in/i }).click();
  await page.getByRole("tab", { name: /participants/i }).or(page.getByRole("button", { name: /^participants$/i })).first().click();
  for (const res of [checker, neutral]) {
    const row = page.getByRole("row").filter({ hasText: res.sessionId.slice(0, 8) });
    await expect(row).toContainText(/finished/i);
    await expect(row).toContainText(res.group, { ignoreCase: true });
  }
  await page.close();
});

function parseCsvLine(line: string): string[] {
  const out: string[] = [];
  let cur = "";
  let q = false;
  for (let i = 0; i < line.length; i++) {
    const c = line[i];
    if (q) {
      if (c === '"' && line[i + 1] === '"') {
        cur += '"';
        i++;
      } else if (c === '"') q = false;
      else cur += c;
    } else if (c === '"') q = true;
    else if (c === ",") {
      out.push(cur);
      cur = "";
    } else cur += c;
  }
  out.push(cur);
  return out;
}
