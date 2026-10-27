// Step router against a mocked step descriptor sequence: the client renders whatever step the server
// returns, numbers events from expectedClientSeq, and never decides the next step itself.

import { describe, expect, it, beforeEach } from "vitest";
import { act, fireEvent, render, screen, waitFor, within } from "@testing-library/react";
import type { Step } from "@study/shared";
import { ApiClient } from "./api/client";
import { StudyRoute } from "./routes/StudyRoute";
import { sceneComponentFor, SCENES } from "./StepRouter";
import { player } from "./audio/player";
import { mockServer, packFromConfig, type SentEvent } from "./test/fixtures";

const pack = packFromConfig();

function setWindow(w: number, h: number) {
  Object.defineProperty(window, "innerWidth", { configurable: true, value: w });
  Object.defineProperty(window, "innerHeight", { configurable: true, value: h });
}

beforeEach(() => setWindow(1440, 900));

const step = (s: Partial<Step> & Pick<Step, "stepId" | "scene">): Step => s as Step;

/**
 * Answers the consent page's HEAD check on the PDF (jsdom cannot fetch a relative URL). `status` is
 * the current answer and can be changed between attempts; returns the restore function.
 */
function mockConsentPdf(status: { code: number }) {
  const realFetch = globalThis.fetch;
  globalThis.fetch = (async (url: RequestInfo | URL, init?: RequestInit) => {
    if (String(url).startsWith("/assets/consent/") && init?.method === "HEAD") {
      return new Response(null, {
        status: status.code,
        headers: status.code === 200 ? { "Content-Type": "application/pdf" } : { "Content-Type": "text/plain" },
      });
    }
    return realFetch(url, init);
  }) as typeof fetch;
  return () => {
    globalThis.fetch = realFetch;
  };
}

describe("sceneComponentFor", () => {
  it("maps every step kind of section 8.1 to a component", () => {
    const kinds = [
      "deviceGate", "consent", "audioCheck", "example", "preStart", "opening", "slide", "postView",
      "rating", "action", "question", "answer", "budget", "background", "debrief", "done",
    ];
    for (const k of kinds) expect(sceneComponentFor(k)).toBeTypeOf("function");
    expect(Object.keys(SCENES).sort()).toEqual([...kinds].sort());
    expect(sceneComponentFor("checkerIntro")).toBeUndefined();
  });
});

describe("StepRouter flow", () => {
  it("walks device gate, consent, audio check and a rating as the server dictates", async () => {
    const steps: Record<string, Step> = {
      consent: step({ stepId: "p2", scene: "consent" }),
      audioCheck: step({ stepId: "p3", scene: "audioCheck", audioUrl: "/a/test.mp3", replayable: true }),
      rating: step({ stepId: "p9", scene: "rating", sceneRef: "rating.T0", point: "T0", closesInMs: 15000 }),
      action: step({
        stepId: "p10", scene: "action", sceneRef: "s2.action", closesInMs: 15000,
        options: [{ id: "nothing", label: "Say nothing" }, { id: "explain", label: "Ask Daniel to explain" }, { id: "challenge", label: "Challenge Daniel with the checker's finding" }],
      }),
    };
    const srv = mockServer({
      pack,
      first: step({ stepId: "p1", scene: "deviceGate" }),
      onEvent: (ev) => {
        switch (ev.type) {
          case "device_checked": return srv.ok(steps.consent);
          case "consent_accepted": return srv.ok(steps.audioCheck);
          case "audio_check_played": return srv.ok(steps.audioCheck);
          case "audio_check_answered": return srv.ok(steps.rating);
          case "rating_submitted": return srv.ok(steps.action);
          default: return srv.ok(steps.action);
        }
      },
    });
    const api = new ApiClient({ fetch: srv.fetch });
    const restoreFetch = mockConsentPdf({ code: 200 });
    render(<StudyRoute code="RUN1" onBlocked={() => undefined} api={api} />);

    // Device gate passes locally and is reported; the server answers with consent.
    const accept = await screen.findByRole("button", { name: pack.consent.acceptLabel });
    expect(srv.events[0]).toMatchObject({ type: "device_checked", clientSeq: 1, payload: { widthPx: 1440, heightPx: 900 } });
    await screen.findByTestId("consent-pdf");
    restoreFetch();
    expect(accept).toBeDisabled();
    fireEvent.click(screen.getByRole("checkbox"));
    expect(accept).toBeEnabled();
    fireEvent.click(accept);

    // Audio check: options locked until the clip was played once; one answer.
    const lantern = await screen.findByRole("button", { name: "lantern" });
    expect(srv.events[1]).toMatchObject({ type: "consent_accepted", clientSeq: 2, payload: { consentVersion: pack.consent.version, checkbox: true } });
    expect(lantern).toBeDisabled();
    fireEvent.click(screen.getByRole("button", { name: pack.labels.playTestSound }));
    await waitFor(() => expect(srv.events.at(-1)?.type).toBe("audio_check_played"));
    expect(lantern).toBeEnabled();
    fireEvent.click(lantern);

    // Rating comes next because the server said so.
    await screen.findByText(pack.rating.instruction);
    const ev = srv.events.find((e) => e.type === "audio_check_answered")!;
    expect(ev).toMatchObject({ clientSeq: 4, payload: { choiceId: "lantern" } });
    fireEvent.click(screen.getByRole("radio", { name: "7" }));
    fireEvent.click(within(screen.getByTestId("scale-competence")).getByRole("radio", { name: "5" }));
    fireEvent.click(screen.getByRole("button", { name: pack.labels.ratingSubmit }));

    await screen.findByRole("button", { name: "Challenge Daniel with the checker's finding" });
    const rated = srv.events.find((e) => e.type === "rating_submitted")!;
    expect(rated).toEqual<SentEvent>({
      clientSeq: 5,
      type: "rating_submitted",
      payload: { point: "T0", trust: 7, competence: 5, benevolence: null, integrity: null },
    });
    // Nothing about the group is sent or shown: the action options come from the step.
    expect(screen.getAllByRole("button")).toHaveLength(3);
  });

  it("applies the step carried by a state error and continues numbering from it", async () => {
    const srv = mockServer({
      pack,
      first: step({ stepId: "p5", scene: "preStart", preStartText: ["The meeting is about to start.", "Checker line."] }),
      onEvent: () => ({
        status: 409,
        body: { error: { code: "wrong_state", message: "x", expectedClientSeq: 12, next: { stepId: "p30", scene: "budget", sceneRef: "budget", closesInMs: 30000 } } },
      }),
    });
    const api = new ApiClient({ fetch: srv.fetch });
    render(<StudyRoute code="RUN1" onBlocked={() => undefined} api={api} />);
    expect(await screen.findByText("Checker line.")).toBeInTheDocument();
    fireEvent.click(screen.getByRole("button", { name: pack.labels.startMeeting }));
    await screen.findByText(/How do you vote\?/);
    expect(api.expectedClientSeq).toBe(12);
  });

  it("shows the checker sentence only when the descriptor carries it", async () => {
    const srv = mockServer({ pack, first: step({ stepId: "p5", scene: "preStart", preStartText: pack.setup.preStartText.both }), onEvent: () => ({ body: {} }) });
    render(<StudyRoute code="RUN1" onBlocked={() => undefined} api={new ApiClient({ fetch: srv.fetch })} />);
    await screen.findByText("The meeting is about to start.");
    expect(screen.queryByText(/checker/i)).toBeNull();
  });

  it("sends the device check and goes to the blocked page when the screen is too small", async () => {
    setWindow(800, 1200);
    const srv = mockServer({ pack, first: step({ stepId: "p1", scene: "deviceGate" }), onEvent: () => srv.ok(step({ stepId: "p1", scene: "deviceGate" })) });
    let blocked = "";
    render(<StudyRoute code="RUN9" onBlocked={(c) => (blocked = c)} api={new ApiClient({ fetch: srv.fetch })} />);
    await waitFor(() => expect(blocked).toBe("RUN9"));
    expect(srv.events[0]).toMatchObject({ type: "device_checked", payload: { widthPx: 800, heightPx: 1200 } });
  });

  it("plays a slide only after scene_started, sends audio_ended on `ended`, and resends after too_early", async () => {
    player.unlocked = true;
    let endedTries = 0;
    const slide = step({
      stepId: "p20", scene: "slide", sceneRef: "s2", audioUrl: "/a/abc.mp3",
      overlay: { status: "flagged", markings: [{ id: "s2.m1", targetIds: ["s2.num1"], appearAtMs: 0, color: "#E8A33D", shape: "box", label: "Numbers do not add up", detail: "Detail." }] },
    });
    const post = step({ stepId: "p21", scene: "postView", sceneRef: "s2", opensInMs: 10000, overlay: slide.overlay });
    const srv = mockServer({
      pack,
      first: slide,
      onEvent: (ev, { seq }) => {
        if (ev.type === "audio_ended" && endedTries++ === 0) {
          return { status: 422, body: { error: { code: "too_early", message: "early", retryable: true, retryInMs: 5, expectedClientSeq: seq + 1, next: slide } } };
        }
        if (ev.type === "audio_ended") return srv.ok(post);
        return srv.ok(slide);
      },
    });
    render(<StudyRoute code="RUN1" onBlocked={() => undefined} api={new ApiClient({ fetch: srv.fetch })} />);
    await waitFor(() => expect(srv.events.some((e) => e.type === "scene_started")).toBe(true));
    expect(srv.events.find((e) => e.type === "scene_started")!.payload).toEqual({ sceneId: "s2" });
    const boardroomEl = screen.getByTestId("boardroom");
    expect(boardroomEl).toHaveAttribute("data-scene", "slide");
    // Slide rendered from the layout, with block ids as element ids.
    expect(document.getElementById("s2.num3")).toHaveTextContent("$8 million");
    // Checker panel: marking label and detail at its time; marking_shown reported once.
    const panel = await screen.findByTestId("checker-panel");
    await waitFor(() => expect(panel).toHaveTextContent("Numbers do not add up"));
    await waitFor(() => expect(srv.events.filter((e) => e.type === "marking_shown")).toHaveLength(1));

    await act(async () => {
      player.element.dispatchEvent(new Event("ended"));
    });
    await waitFor(() => expect(srv.events.filter((e) => e.type === "audio_ended")).toHaveLength(2));
    const ended = srv.events.filter((e) => e.type === "audio_ended");
    expect(ended[1].clientSeq).toBe(ended[0].clientSeq + 1);
    // The pause after the clip: the same boardroom element stays mounted (no blip), the slide and the panel
    // stay, nobody speaks, and there is no countdown card (Oliver, 2026-10-26).
    await waitFor(() => expect(boardroomEl).toHaveAttribute("data-scene", "postView"));
    expect(screen.getByTestId("boardroom")).toBe(boardroomEl);
    expect(document.getElementById("s2.num3")).toHaveTextContent("$8 million");
    expect(screen.getByTestId("checker-panel")).toHaveTextContent("Numbers do not add up");
    expect(screen.queryByRole("timer")).toBeNull();
    expect(document.querySelector("[data-speaking]")).toBeNull();
  });

  it("keeps one boardroom from the pause to the rating, which takes the meeting-room screen", async () => {
    const post = step({ stepId: "p21", scene: "postView", sceneRef: "s2", opensInMs: 600 });
    const rating = step({ stepId: "p22", scene: "rating", sceneRef: "rating.T2a", point: "T2a", closesInMs: 20000 });
    // The route asks GET /next once at start; only the poll at the end of the pause may move to the rating.
    let armed = false;
    const srv = mockServer({
      pack,
      first: post,
      onEvent: () => srv.ok(step({ stepId: "p23", scene: "done" })),
      onNext: () => (armed ? rating : undefined),
    });
    render(<StudyRoute code="RUN1" onBlocked={() => undefined} api={new ApiClient({ fetch: srv.fetch })} />);
    const boardroomEl = await screen.findByTestId("boardroom");
    expect(boardroomEl).toHaveAttribute("data-scene", "postView");
    expect(document.getElementById("s2.num3")).toHaveTextContent("$8 million");
    // The server moves on to the rating: same boardroom element, the form on the room's screen, slide gone.
    armed = true;
    await waitFor(() => expect(boardroomEl).toHaveAttribute("data-scene", "rating"), { timeout: 3000 });
    expect(screen.getByTestId("boardroom")).toBe(boardroomEl);
    const screenContent = screen.getByTestId("screen-content");
    expect(boardroomEl).toContainElement(screenContent);
    expect(screenContent).toContainElement(screen.getByTestId("scale-trust"));
    expect(within(screenContent).getByText(pack.rating.instruction)).toBeVisible();
    expect(document.getElementById("s2.num3")).toBeNull();
    expect(screen.getByRole("timer")).toHaveAccessibleName(/Time left: 20 seconds/);
    fireEvent.click(within(screen.getByTestId("scale-trust")).getByRole("radio", { name: "8" }));
    fireEvent.click(screen.getByRole("button", { name: pack.labels.ratingSubmit }));
    await waitFor(() => expect(srv.events.some((e) => e.type === "rating_submitted")).toBe(true));
    expect(srv.events.find((e) => e.type === "rating_submitted")!.payload).toEqual({ point: "T2a", trust: 8, competence: null, benevolence: null, integrity: null });
  });

  it("sends rating_timeout with the partial values when the window closes", async () => {
    const srv = mockServer({
      pack,
      first: step({ stepId: "p9", scene: "rating", sceneRef: "rating.T1", point: "T1", closesInMs: 400 }),
      onEvent: () => srv.ok(step({ stepId: "p10", scene: "done" })),
    });
    render(<StudyRoute code="RUN1" onBlocked={() => undefined} api={new ApiClient({ fetch: srv.fetch })} />);
    await screen.findByText(pack.rating.instruction);
    fireEvent.click(within(screen.getByTestId("scale-integrity")).getByRole("radio", { name: "2" }));
    await waitFor(() => expect(srv.events.some((e) => e.type === "rating_timeout")).toBe(true), { timeout: 2000 });
    expect(srv.events.find((e) => e.type === "rating_timeout")!.payload).toEqual({
      point: "T1",
      partial: { trust: null, competence: null, benevolence: null, integrity: 2 },
    });
  });

  it("keeps the slide and the checker overlay on the action screen when the descriptor carries it", async () => {
    const overlay = { status: "flagged" as const, markings: [{ id: "s2.m1", targetIds: ["s2.num1"], appearAtMs: 0, color: "#E8A33D", shape: "box" as const, label: "Numbers do not add up", detail: "Detail." }] };
    const action = step({
      stepId: "p10", scene: "action", sceneRef: "s2.action", closesInMs: 15000, overlay,
      options: [{ id: "nothing", label: "Say nothing" }, { id: "challenge", label: "Challenge Daniel with the checker's finding" }],
    });
    const srv = mockServer({ pack, first: action, onEvent: () => srv.ok(action) });
    render(<StudyRoute code="RUN1" onBlocked={() => undefined} api={new ApiClient({ fetch: srv.fetch })} />);
    await screen.findByRole("button", { name: "Say nothing" });
    expect(document.getElementById("s2.num3")).toHaveTextContent("$8 million");
    expect(await screen.findByTestId("checker-panel")).toHaveTextContent("Numbers do not add up");
    // Shown again, not newly: no marking_shown outside the slide step.
    expect(srv.events.filter((e) => e.type === "marking_shown")).toHaveLength(0);
  });

  it("keeps the consent checkbox checkable when the PDF frame is shown, in the same frame area", async () => {
    const restoreFetch = mockConsentPdf({ code: 200 });
    try {
      const srv = mockServer({ pack, first: step({ stepId: "p2", scene: "consent" }), onEvent: () => ({ body: {} }) });
      render(<StudyRoute code="RUN1" onBlocked={() => undefined} api={new ApiClient({ fetch: srv.fetch })} />);
      const area = await screen.findByTestId("consent-pdf-area");
      const frame = await screen.findByTestId("consent-pdf");
      // The frame replaces the placeholder inside the same area: the checkbox and button are not moved.
      expect(area).toContainElement(frame);
      expect(screen.queryByTestId("consent-placeholder")).toBeNull();
      const box = screen.getByRole("checkbox");
      const accept = screen.getByRole("button", { name: pack.consent.acceptLabel });
      expect(area.compareDocumentPosition(box) & Node.DOCUMENT_POSITION_FOLLOWING).toBeTruthy();
      expect(accept).toBeDisabled();
      fireEvent.click(box);
      expect(box).toBeChecked();
      expect(accept).toBeEnabled();
      expect(screen.getByRole("link", { name: /Download the consent form/ })).toBeVisible();
    } finally {
      restoreFetch();
    }
  });

  it("does not let anyone accept while the consent form cannot be loaded, and recovers on Try again", async () => {
    const status = { code: 404 };
    const restoreFetch = mockConsentPdf(status);
    try {
      const srv = mockServer({ pack, first: step({ stepId: "p2", scene: "consent" }), onEvent: () => ({ body: {} }) });
      render(<StudyRoute code="RUN1" onBlocked={() => undefined} api={new ApiClient({ fetch: srv.fetch })} />);
      const notice = await screen.findByRole("alert");
      expect(notice).toHaveTextContent("The consent form could not be loaded");
      expect(screen.queryByTestId("consent-pdf")).toBeNull();
      const accept = screen.getByRole("button", { name: pack.consent.acceptLabel });
      fireEvent.click(screen.getByRole("checkbox"));
      // Ticking the box is not enough: the form was never shown.
      expect(accept).toBeDisabled();
      expect(srv.events.filter((e) => e.type === "consent_accepted")).toHaveLength(0);

      status.code = 200;
      fireEvent.click(screen.getByRole("button", { name: "Try again" }));
      await screen.findByTestId("consent-pdf");
      expect(screen.queryByRole("alert")).toBeNull();
      expect(accept).toBeEnabled();
    } finally {
      restoreFetch();
    }
  });
});
