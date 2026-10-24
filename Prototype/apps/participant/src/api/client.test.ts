import { describe, expect, it } from "vitest";
import { ApiClient, ApiError } from "./client";

const json = (status: number, body: unknown) =>
  new Response(JSON.stringify(body), { status, headers: { "Content-Type": "application/json" } });

describe("ApiClient", () => {
  it("takes clientSeq from every response, including errors, and sends the study header and cookie", async () => {
    const calls: { url: string; init?: RequestInit }[] = [];
    const replies = [
      json(201, { sessionId: "S1", scenePackUrl: "/p", assetBaseUrl: "/a", configVersion: "1", expectedClientSeq: 4, next: { stepId: "p1", scene: "consent" } }),
      json(409, { error: { code: "seq_gap", message: "gap", expectedClientSeq: 9, next: { stepId: "p2", scene: "audioCheck" } } }),
      json(200, { accepted: true, eventNo: 3, expectedClientSeq: 10, next: { stepId: "p3", scene: "example" } }),
    ];
    const steps: string[] = [];
    const api = new ApiClient({
      fetch: async (url, init) => {
        calls.push({ url, init });
        return replies.shift()!;
      },
      onStep: (s) => steps.push(s.stepId),
    });
    await api.createSession("RUN1");
    expect(api.expectedClientSeq).toBe(4);
    await expect(api.send("setup_done", {})).rejects.toBeInstanceOf(ApiError);
    expect(api.expectedClientSeq).toBe(9);
    await api.send("setup_done", {});
    expect(api.expectedClientSeq).toBe(10);

    const sent = calls.slice(1).map((c) => JSON.parse(String(c.init!.body)).clientSeq);
    expect(sent).toEqual([4, 9]);
    for (const c of calls) {
      expect((c.init!.headers as Record<string, string>)["X-Study-Client"]).toBe("1");
      expect(c.init!.credentials).toBe("include");
    }
    expect(steps).toEqual(["p1", "p2", "p3"]);
  });

  it("retries a network failure with the identical body", async () => {
    const bodies: string[] = [];
    let n = 0;
    const api = new ApiClient({
      retryDelayMs: 1,
      fetch: async (_url, init) => {
        bodies.push(String(init?.body));
        if (n++ === 0) throw new TypeError("offline");
        return json(200, { accepted: true, eventNo: 1, expectedClientSeq: 2, next: { stepId: "x", scene: "done" } });
      },
    });
    api.sessionId = "S1";
    await api.send("finished", {});
    expect(bodies).toHaveLength(2);
    expect(bodies[0]).toBe(bodies[1]);
  });

  it("resends a too_early event after retryInMs with the next clientSeq", async () => {
    const seqs: number[] = [];
    let n = 0;
    const api = new ApiClient({
      fetch: async (_url, init) => {
        seqs.push(JSON.parse(String(init?.body)).clientSeq);
        if (n++ === 0) {
          return json(422, { error: { code: "too_early", message: "early", retryable: true, retryInMs: 5, expectedClientSeq: 8 } });
        }
        return json(200, { accepted: true, eventNo: 2, expectedClientSeq: 9, next: { stepId: "p", scene: "postView" } });
      },
    });
    api.sessionId = "S1";
    api.expectedClientSeq = 7;
    await api.sendUntilOnTime("audio_ended", { sceneId: "s1" });
    expect(seqs).toEqual([7, 8]);
  });
});
