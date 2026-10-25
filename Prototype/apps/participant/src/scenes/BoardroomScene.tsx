// `<BoardroomScene>` (section 9.2): a fixed 16:9 stage, scaled to the viewport, as a still composite
// (2.10). Layers: boardroom background; the presentation screen with the slide; the CISO portrait
// beside the screen; at the table the chair, the participant's own seat ("Board member (You)", marked
// with a badge and a lighter frame) and three more silent board members, so the board looks full
// (Oliver, 2026-10-25); for checker sessions the checker panel at the screen's top right corner. The
// portrait of whoever speaks gets a soft glow and a speaking indicator while the clip plays; the
// participant's seat lights while the board-member voice speaks the participant's question.

import { useEffect, useRef, useState, type ReactNode } from "react";
import type {
  Block,
  LayoutTemplate,
  ProfileCard,
  StepOverlay,
} from "@study/shared";
import { SlideRenderer } from "../slide/SlideRenderer";
import { CheckerMarks, CheckerPanel } from "../slide/CheckerOverlay";
import { IMAGE_URLS, type PortraitRole } from "../assets";
import { useStudy } from "../study";
import styles from "./BoardroomScene.module.css";

export type Speaker = "ciso" | "chair" | "member" | null;

interface Props {
  layout?: Block[];
  template?: LayoutTemplate;
  slideNumber?: number;
  overlay?: StepOverlay;
  /** Audio position in ms; Infinity shows every marking (after the clip, post view, question, answer). */
  clockMs: number;
  speaking: Speaker;
  profileCard?: ProfileCard;
  onMarkingShown?: (id: string) => void;
  /** Controls drawn in the stage's lower right corner (countdown, Continue). */
  controls?: ReactNode;
  /** A larger card at the bottom of the stage, between the table portraits and the CISO (action choice). */
  dialog?: ReactNode;
}

const STAGE_W = 1600;
const STAGE_H = 900;
const SCREEN = { x: 40, y: 36, w: 1088, h: 612 };

export function BoardroomScene(p: Props) {
  const { pack, assetBaseUrl } = useStudy();
  const outerRef = useRef<HTMLDivElement>(null);
  const slideRef = useRef<HTMLDivElement>(null);
  const [scale, setScale] = useState(0.6);

  useEffect(() => {
    const el = outerRef.current;
    if (!el) return;
    const fit = () => {
      const w = el.clientWidth || window.innerWidth;
      const h = el.clientHeight || window.innerHeight;
      setScale(Math.max(0.1, Math.min(w / STAGE_W, h / STAGE_H)));
    };
    fit();
    const ro =
      typeof ResizeObserver !== "undefined" ? new ResizeObserver(fit) : null;
    ro?.observe(el);
    window.addEventListener("resize", fit);
    return () => {
      ro?.disconnect();
      window.removeEventListener("resize", fit);
    };
  }, []);

  const hasSlide = !!p.layout && p.layout.length > 0;
  const slideScale = SCREEN.w / 1600;

  return (
    <div ref={outerRef} className={styles.outer}>
      <div
        className={styles.stage}
        style={{
          width: STAGE_W,
          height: STAGE_H,
          transform: `translate(-50%, -50%) scale(${scale})`,
        }}
        data-testid="boardroom"
      >
        <Room />
        <div
          className={styles.screenFrame}
          style={{
            left: SCREEN.x - 10,
            top: SCREEN.y - 10,
            width: SCREEN.w + 20,
            height: SCREEN.h + 20,
          }}
        />
        <div
          className={styles.screen}
          style={{
            left: SCREEN.x,
            top: SCREEN.y,
            width: SCREEN.w,
            height: SCREEN.h,
          }}
        >
          {hasSlide && (
            <div
              style={{
                transform: `scale(${slideScale})`,
                transformOrigin: "0 0",
              }}
            >
              <SlideRenderer
                ref={slideRef}
                layout={p.layout!}
                template={p.template}
                slideNumber={p.slideNumber}
                assetBaseUrl={assetBaseUrl}
              >
                <CheckerMarks
                  overlay={p.overlay}
                  clockMs={p.clockMs}
                  slideRef={slideRef}
                  onShown={p.onMarkingShown}
                />
              </SlideRenderer>
            </div>
          )}
          {p.profileCard && (
            <ProfileCardView card={p.profileCard} overSlide={hasSlide} />
          )}
        </div>

        {p.overlay && (
          <div className={styles.panelSlot}>
            <CheckerPanel
              overlay={p.overlay}
              clockMs={p.clockMs}
              title={p.overlay.title ?? pack.labels.panelTitle ?? ""}
              idleText={p.overlay.idleText ?? pack.labels.panelIdle ?? ""}
              noIssuesColor={
                p.overlay.noIssuesColor ?? pack.colors?.noIssues ?? "#5B8C6A"
              }
            />
          </div>
        )}

        <Portrait
          role="ciso"
          label="CISO"
          speaking={p.speaking === "ciso"}
          className={styles.ciso}
        />
        <div className={styles.table} />
        <Portrait
          role="chair"
          label="Chair"
          speaking={p.speaking === "chair"}
          className={styles.chair}
          tone={0}
        />
        <Portrait
          role="member"
          label="Board member (You)"
          you
          speaking={p.speaking === "member"}
          className={styles.member}
          tone={1}
        />
        <Portrait
          role="member2"
          label="Board member"
          speaking={false}
          className={styles.member2}
          tone={2}
        />
        <Portrait
          role="member3"
          label="Board member"
          speaking={false}
          className={styles.member3}
          tone={3}
        />
        <Portrait
          role="member4"
          label="Board member"
          speaking={false}
          className={styles.member4}
          tone={4}
        />

        {p.controls && <div className={styles.controls}>{p.controls}</div>}
        {p.dialog && <div className={styles.dialog}>{p.dialog}</div>}
      </div>
    </div>
  );
}

function useImage(url: string): boolean {
  const [ok, setOk] = useState(false);
  useEffect(() => {
    if (typeof Image === "undefined") return;
    const img = new Image();
    img.onload = () => setOk(true);
    img.src = url;
    return () => {
      img.onload = null;
    };
  }, [url]);
  return ok;
}

function Room() {
  const hasPhoto = useImage(IMAGE_URLS.boardroom);
  if (hasPhoto)
    return <img className={styles.room} src={IMAGE_URLS.boardroom} alt="" />;
  return (
    <svg className={styles.room} viewBox="0 0 1600 900" aria-hidden="true">
      <defs>
        <linearGradient id="wall" x1="0" y1="0" x2="0" y2="1">
          <stop offset="0" stopColor="#3b4250" />
          <stop offset="1" stopColor="#2a2f39" />
        </linearGradient>
      </defs>
      <rect width="1600" height="900" fill="url(#wall)" />
      <rect y="660" width="1600" height="240" fill="#1d2129" />
      {[0, 1, 2, 3, 4, 5, 6, 7].map((i) => (
        <rect
          key={i}
          x={i * 200 + 8}
          y="0"
          width="184"
          height="660"
          fill="#ffffff"
          opacity="0.025"
        />
      ))}
    </svg>
  );
}

// Placeholder silhouettes vary a little by seat so the table does not look like one person copied.
const TONES = [
  { bg: "#4a5262", fg: "#9aa3b2", r: 38 },
  { bg: "#4d5468", fg: "#aab3c1", r: 37 },
  { bg: "#465064", fg: "#8e98a8", r: 39 },
  { bg: "#505866", fg: "#a3abb8", r: 36 },
  { bg: "#434b5c", fg: "#97a1b0", r: 40 },
];

function Portrait({
  role,
  label,
  speaking,
  className,
  you = false,
  tone = 0,
}: {
  role: PortraitRole;
  label: string;
  speaking: boolean;
  className: string;
  /** The participant's own seat: badge and lighter frame (2026-10-25). */
  you?: boolean;
  tone?: number;
}) {
  const hasPhoto = useImage(IMAGE_URLS[role]);
  const t = TONES[tone % TONES.length];
  return (
    <figure
      className={`${styles.portrait} ${className} ${speaking ? styles.speaking : ""} ${you ? styles.you : ""}`}
      data-speaking={speaking || undefined}
      data-you={you || undefined}
    >
      <div className={styles.portraitImage}>
        {hasPhoto ? (
          <img src={IMAGE_URLS[role]} alt="" />
        ) : (
          <svg viewBox="0 0 200 200" aria-hidden="true">
            <rect width="200" height="200" fill={t.bg} />
            <circle cx="100" cy="78" r={t.r} fill={t.fg} />
            <path d="M28 200c4-52 34-78 72-78s68 26 72 78z" fill={t.fg} />
          </svg>
        )}
        {you && (
          <span className={styles.youBadge} aria-hidden="true">
            You
          </span>
        )}
        {speaking && (
          <span className={styles.speakingBars} aria-hidden="true">
            <i />
            <i />
            <i />
          </span>
        )}
      </div>
      <figcaption>{label}</figcaption>
    </figure>
  );
}

function ProfileCardView({
  card,
  overSlide,
}: {
  card: ProfileCard;
  overSlide: boolean;
}) {
  return (
    <div
      className={`${styles.profileCard} ${overSlide ? styles.profileCardOver : ""}`}
      data-testid="profile-card"
    >
      <div className={styles.profileHead}>
        <svg viewBox="0 0 200 200" width="96" height="96" aria-hidden="true">
          <rect width="200" height="200" rx="16" fill="#4a5262" />
          <circle cx="100" cy="78" r="38" fill="#9aa3b2" />
          <path d="M28 200c4-52 34-78 72-78s68 26 72 78z" fill="#9aa3b2" />
        </svg>
        <div>
          <div className={styles.profileName}>{card.name}</div>
          <div className={styles.profileRole}>{card.role}</div>
        </div>
      </div>
      <ul className={styles.profileLines}>
        {card.lines.map((l) => (
          <li key={l}>{l}</li>
        ))}
      </ul>
    </div>
  );
}
