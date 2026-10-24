// Rating, action, budget and background components on their own, inside a StudyContext.

import { describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen, waitFor, within } from "@testing-library/react";
import type { ReactElement } from "react";
import type { Step } from "@study/shared";
import { StudyContext, type StudyContextValue } from "../study";
import { ApiClient } from "../api/client";
import { packFromConfig } from "../test/fixtures";
import { Rating } from "./Rating";
import { Action } from "./Action";
import { Budget } from "./Budget";
import { Background } from "./Background";

const pack = packFromConfig();

function harness(ui: ReactElement) {
  const sent: { type: string; payload: unknown }[] = [];
  const api = new ApiClient();
  api.send = vi.fn(async (type: string, payload: unknown) => {
    sent.push({ type, payload });
    return { accepted: true, eventNo: 1, expectedClientSeq: 2, next: { stepId: "x", scene: "done" } } as never;
  }) as never;
  api.sendUntilOnTime = api.send as never;
  const ctx: StudyContextValue = {
    api, pack, assetBaseUrl: "/a", code: "RUN1",
    refresh: async () => undefined, reportError: () => undefined, block: () => undefined,
  };
  render(<StudyContext.Provider value={ctx}>{ui}</StudyContext.Provider>);
  return sent;
}

describe("Rating", () => {
  const rating: Step = { stepId: "p1", scene: "rating", point: "T2a", closesInMs: 15000 };

  it("shows trust as eleven numbers and three 7-point antonym rows, nothing preselected, in fixed order", () => {
    harness(<Rating step={rating} />);
    const trust = screen.getByTestId("scale-trust");
    expect(within(trust).getAllByRole("radio")).toHaveLength(11);
    expect(within(trust).getByText("Your trust in Daniel Hale")).toBeVisible();
    for (const id of ["competence", "benevolence", "integrity"]) {
      expect(within(screen.getByTestId(`scale-${id}`)).getAllByRole("radio")).toHaveLength(7);
    }
    const ints = screen.getByTestId("scale-competence");
    expect(ints.textContent).toMatch(/^incompetent.*1.*7.*competent$/s);
    expect(screen.getAllByRole("radio").every((r) => !(r as HTMLInputElement).checked)).toBe(true);
    const order = screen.getAllByTestId(/^scale-/).map((e) => e.getAttribute("data-testid"));
    expect(order).toEqual(["scale-trust", "scale-competence", "scale-benevolence", "scale-integrity"]);
    expect(screen.getByRole("timer")).toHaveAccessibleName(/Time left: 15 seconds/);
  });

  it("submits untouched items as null, once", async () => {
    const sent = harness(<Rating step={rating} />);
    fireEvent.click(within(screen.getByTestId("scale-benevolence")).getByRole("radio", { name: /^6/ }));
    const btn = screen.getByRole("button", { name: "Next" });
    fireEvent.click(btn);
    fireEvent.click(btn);
    await waitFor(() => expect(sent).toHaveLength(1));
    expect(sent[0]).toEqual({
      type: "rating_submitted",
      payload: { point: "T2a", trust: null, competence: null, benevolence: 6, integrity: null },
    });
  });
});

describe("Action", () => {
  it("renders yes/no when the step carries two options and sends the choice for the slide", async () => {
    const sent = harness(
      <Action step={{ stepId: "p2", scene: "action", sceneRef: "s3.action", closesInMs: 15000, options: [{ id: "hasQuestion", label: "Yes" }, { id: "noQuestion", label: "No" }] }} />,
    );
    expect(screen.getAllByRole("button")).toHaveLength(2);
    fireEvent.click(screen.getByRole("button", { name: "No" }));
    await waitFor(() => expect(sent).toHaveLength(1));
    expect(sent[0]).toEqual({ type: "action_chosen", payload: { slideId: "s3", actionId: "noQuestion" } });
    expect(screen.getByRole("button", { name: "Yes" })).toBeDisabled();
  });

  it("sends action_timeout when the window closes", async () => {
    const sent = harness(<Action step={{ stepId: "p2", scene: "action", sceneRef: "s4.action", closesInMs: 100, options: [{ id: "hasQuestion", label: "Yes" }] }} />);
    await waitFor(() => expect(sent).toEqual([{ type: "action_timeout", payload: { slideId: "s4" } }]), { timeout: 2000 });
  });
});

describe("Budget", () => {
  it("sends approve or reject and the 0 to 10 sureness, null when untouched", async () => {
    const sent = harness(<Budget step={{ stepId: "p3", scene: "budget", sceneRef: "budget", closesInMs: 30000 }} />);
    expect(within(screen.getByTestId("scale-confidence")).getAllByRole("radio")).toHaveLength(11);
    fireEvent.click(screen.getByRole("radio", { name: "Reject the request" }));
    fireEvent.click(screen.getByRole("button", { name: "Next" }));
    await waitFor(() => expect(sent).toEqual([{ type: "budget_submitted", payload: { optionId: "reject", confidence: null } }]));
  });
});

describe("Background", () => {
  it("offers Prefer not to say on every item and keeps the honeypot hidden", async () => {
    const sent = harness(<Background step={{ stepId: "p4", scene: "background", sceneRef: "background", closesInMs: 30000 }} />);
    expect(screen.getAllByRole("radio", { name: "Prefer not to say" })).toHaveLength(6);
    expect(screen.getByRole("combobox")).toBeInTheDocument();
    expect(screen.queryByRole("textbox", { name: /honeypot/i })).toBeNull();
    fireEvent.click(screen.getAllByRole("radio", { name: "Prefer not to say" })[0]);
    fireEvent.change(screen.getByRole("combobox"), { target: { value: "CA" } });
    fireEvent.click(screen.getByRole("button", { name: "Next" }));
    await waitFor(() => expect(sent).toHaveLength(1));
    expect(sent[0]).toEqual({
      type: "background_submitted",
      payload: {
        items: { field: "preferNotToSay", country: "CA", age: null, jobLevel: null, education: null, cyberKnowledge: null, aiUse: null },
        hp: "",
      },
    });
  });
});
