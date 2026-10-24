// The one audio element for the whole page life ("02 - Architecture.md" section 9.3).
// No controls, no seek, no keyboard shortcut. Each step swaps `src`. The next clip is prefetched with
// fetch() into the HTTP cache, not into a second element (Safari unlocks per element).

type Listener = () => void;

class Player {
  private el: HTMLAudioElement | null = null;
  /** True once play() has succeeded after a user gesture in this page life. */
  unlocked = false;
  /** True while a study clip should be playing; an unexpected pause (OS media keys) is undone. */
  private wantPlaying = false;
  private listeners = new Set<Listener>();
  private prefetched = new Set<string>();
  /** Bumped by every real playback so a pending prime does not pause it. */
  private primeId = 0;

  get element(): HTMLAudioElement {
    if (!this.el) {
      const el = document.createElement("audio");
      el.preload = "auto";
      el.controls = false;
      el.setAttribute("aria-hidden", "true");
      el.style.display = "none";
      document.body.appendChild(el);
      el.addEventListener("pause", () => {
        if (this.wantPlaying && !el.ended) void el.play().catch(() => undefined);
      });
      for (const ev of ["play", "playing", "pause", "ended"]) el.addEventListener(ev, () => this.emit());
      this.el = el;
      this.neutraliseMediaSession();
    }
    return this.el;
  }

  private neutraliseMediaSession() {
    const ms = (navigator as Navigator & { mediaSession?: MediaSession }).mediaSession;
    if (!ms) return;
    const noop = () => undefined;
    for (const a of ["play", "pause", "stop", "seekbackward", "seekforward", "seekto", "previoustrack", "nexttrack"]) {
      try {
        ms.setActionHandler(a as MediaSessionAction, noop);
      } catch {
        /* action not supported */
      }
    }
  }

  private emit() {
    for (const l of this.listeners) l();
  }
  subscribe(l: Listener): () => void {
    this.listeners.add(l);
    return () => this.listeners.delete(l);
  }

  get isPlaying(): boolean {
    return !!this.el && !this.el.paused && !this.el.ended;
  }
  get currentTimeMs(): number {
    return this.el ? this.el.currentTime * 1000 : 0;
  }
  get durationMs(): number {
    const d = this.el?.duration;
    return d && Number.isFinite(d) ? d * 1000 : NaN;
  }

  /** Loads a clip without playing it. */
  load(url: string) {
    const el = this.element;
    if (el.getAttribute("src") !== url) {
      this.wantPlaying = false;
      el.pause();
      el.src = url;
      el.load();
    }
  }

  /**
   * Must run inside a click handler: plays and immediately pauses the current element so that later
   * play() calls, made after a network round trip, are allowed (autoplay unlock).
   */
  primeFromGesture(url?: string) {
    const el = this.element;
    if (url) this.load(url);
    const id = ++this.primeId;
    el.muted = true;
    const p = el.play();
    const done = () => {
      if (id !== this.primeId) return;
      el.pause();
      el.currentTime = 0;
      el.muted = false;
    };
    if (p && typeof p.then === "function") {
      p.then(() => {
        this.unlocked = true;
        done();
      }).catch(done);
    } else {
      done();
    }
  }

  /** Plays the loaded clip from the start. */
  async playFromStart(): Promise<void> {
    const el = this.element;
    this.primeId++;
    el.muted = false;
    el.currentTime = 0;
    this.wantPlaying = true;
    await el.play();
    this.unlocked = true;
  }

  stop() {
    this.wantPlaying = false;
    if (this.el) this.el.pause();
  }

  /** Marks the clip as ended so a pause event is not undone. */
  settle() {
    this.wantPlaying = false;
  }

  prefetch(urls: string[] | undefined) {
    for (const u of urls ?? []) {
      if (this.prefetched.has(u)) continue;
      this.prefetched.add(u);
      void fetch(u, { credentials: "include" }).catch(() => this.prefetched.delete(u));
    }
  }
}

export const player = new Player();
