// jsdom has no media playback, CSS.escape or matchMedia; stub them.
import { afterEach, vi } from "vitest";
import { cleanup } from "@testing-library/react";
import "@testing-library/jest-dom/vitest";

afterEach(() => cleanup());

Object.defineProperty(HTMLMediaElement.prototype, "play", {
  configurable: true,
  value: vi.fn(function (this: HTMLMediaElement) {
    return Promise.resolve();
  }),
});
Object.defineProperty(HTMLMediaElement.prototype, "pause", { configurable: true, value: vi.fn() });
Object.defineProperty(HTMLMediaElement.prototype, "load", { configurable: true, value: vi.fn() });

if (!(globalThis as { CSS?: { escape?: unknown } }).CSS?.escape) {
  (globalThis as unknown as { CSS: { escape: (s: string) => string } }).CSS = {
    escape: (s: string) => s.replace(/["\\]/g, "\\$&"),
  };
}
if (!window.matchMedia) {
  window.matchMedia = ((q: string) => ({
    matches: false,
    media: q,
    onchange: null,
    addListener: () => undefined,
    removeListener: () => undefined,
    addEventListener: () => undefined,
    removeEventListener: () => undefined,
    dispatchEvent: () => false,
  })) as unknown as typeof window.matchMedia;
}

// No real network in tests: anything not going through an injected ApiClient fetch gets a 404.
globalThis.fetch = vi.fn(async () => new Response("", { status: 404 })) as unknown as typeof fetch;
