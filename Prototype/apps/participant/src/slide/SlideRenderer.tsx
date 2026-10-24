// Renders a slide layout on the 1600 x 900 logical canvas ("02 - Architecture.md" sections 4.1, 9.2).
// Every block and span carries its config id as both `id` and `data-seg-id`, so the checker overlay can
// find its box with getBoundingClientRect. The look follows content/slides/v1.3/*.svg (navy title bar,
// amber accent line, square amber bullets, number panels, cost table, footer rule and slide number).

import { forwardRef, useState, type ReactNode } from "react";
import {
  SLIDE_HEIGHT,
  SLIDE_WIDTH,
  type Block,
  type BulletBlock,
  type ChartBlock,
  type ImageBlock,
  type LayoutTemplate,
  type NumberBlock,
  type RichText,
} from "@study/shared";
import styles from "./SlideRenderer.module.css";
import { ChartView } from "./ChartView";

export function Rich({ text }: { text: RichText | undefined }) {
  if (text === undefined) return null;
  if (typeof text === "string") return <>{text}</>;
  return (
    <>
      {text.map((s) => (
        <span key={s.id} id={s.id} data-seg-id={s.id}>
          {s.text}
        </span>
      ))}
    </>
  );
}

export function richToString(text: RichText | undefined): string {
  if (text === undefined) return "";
  return typeof text === "string" ? text : text.map((s) => s.text).join("");
}

interface Props {
  layout: Block[];
  template?: LayoutTemplate;
  slideNumber?: number;
  assetBaseUrl: string;
  /** Rendered on top of the slide, in slide coordinates (the checker markings). */
  children?: ReactNode;
}

type Arrangement = "table" | "media" | "bulletsAndNumbers" | "numberRows" | "bullets";

// DECISION: the config has no field that asks for a bullet reveal, so all bullets are shown at once.
// DECISION: the layout is chosen from the template and the blocks, matching the rendered v1.3 SVGs:
//   "numbers" with only number blocks (4 or more) -> cost table, last row bold as the total (slide 4);
//   any image or chart -> media on the left, bullets on the right (slide 3, whose SVG puts the image
//     on the left although the template is called "image-right");
//   numbers and bullets under "auto"/"title-bullets" -> bullets left, numbers in a side panel (slide 1);
//   "numbers" with bullets, or numbers alone -> full-width number rows, bullets below (slide 2);
//   otherwise a bullet list.
// The "figure: label" colon of the content tables is not added; the config splits value and caption.
// The "Meridian Logistics" brand mark of the SVGs is not drawn because it is not in the config.
export function arrangementFor(layout: Block[], template: LayoutTemplate | undefined): Arrangement {
  const bullets = layout.filter((b) => b.kind === "bullet").length;
  const numbers = layout.filter((b) => b.kind === "number").length;
  const media = layout.some((b) => b.kind === "image" || b.kind === "chart");
  const tpl = template ?? "auto";
  if (tpl === "numbers" && bullets === 0 && numbers >= 4 && !media) return "table";
  if (media) return "media";
  if (numbers > 0 && bullets > 0 && tpl !== "numbers") return "bulletsAndNumbers";
  if (numbers > 0) return "numberRows";
  return "bullets";
}

export const SlideRenderer = forwardRef<HTMLDivElement, Props>(function SlideRenderer(
  { layout, template, slideNumber, assetBaseUrl, children },
  ref,
) {
  const title = layout.find((b) => b.kind === "title");
  const bullets = layout.filter((b): b is BulletBlock => b.kind === "bullet");
  const numbers = layout.filter((b): b is NumberBlock => b.kind === "number");
  const media = layout.filter((b): b is ImageBlock | ChartBlock => b.kind === "image" || b.kind === "chart");
  const arrangement = arrangementFor(layout, template);

  return (
    <div
      ref={ref}
      className={styles.slide}
      style={{ width: SLIDE_WIDTH, height: SLIDE_HEIGHT }}
      data-arrangement={arrangement}
    >
      <div className={styles.titleBar}>
        {title && (
          <h2 className={styles.title} id={title.id} data-seg-id={title.id}>
            <Rich text={title.kind === "title" ? title.text : undefined} />
          </h2>
        )}
      </div>
      <div className={styles.accent} />

      {arrangement === "table" && <CostTable numbers={numbers} />}

      {arrangement === "media" && (
        <div className={styles.mediaGrid}>
          <div className={styles.mediaCol}>
            {media.map((m) =>
              m.kind === "image" ? (
                <SlideImage key={m.id} block={m} assetBaseUrl={assetBaseUrl} />
              ) : (
                <div key={m.id} id={m.id} data-seg-id={m.id} className={styles.chartBox}>
                  <ChartView block={m} />
                  {m.caption !== undefined && (
                    <div className={styles.chartCaption}>
                      <Rich text={m.caption} />
                    </div>
                  )}
                </div>
              ),
            )}
          </div>
          <div className={styles.sideCol}>
            {numbers.map((n) => (
              <NumberPanel key={n.id} block={n} />
            ))}
            <BulletList bullets={bullets} />
          </div>
        </div>
      )}

      {arrangement === "bulletsAndNumbers" && (
        <div className={styles.twoCol}>
          <div className={styles.leftCol}>
            <BulletList bullets={bullets} />
          </div>
          <div className={styles.rightCol}>
            {numbers.map((n) => (
              <NumberPanel key={n.id} block={n} />
            ))}
          </div>
        </div>
      )}

      {arrangement === "numberRows" && (
        <div className={styles.body}>
          {numbers.map((n) => (
            <div key={n.id} id={n.id} data-seg-id={n.id} className={styles.numberRow}>
              <span className={styles.numberRowValue}>
                <Rich text={n.value} />
              </span>
              <span className={styles.numberRowCaption}>
                <Rich text={n.caption} />
              </span>
            </div>
          ))}
          <BulletList bullets={bullets} />
        </div>
      )}

      {arrangement === "bullets" && (
        <div className={styles.body}>
          <BulletList bullets={bullets} />
        </div>
      )}

      <div className={styles.footerRule} />
      {slideNumber !== undefined && <div className={styles.slideNumber}>{slideNumber}</div>}
      {children}
    </div>
  );
});

function BulletList({ bullets }: { bullets: BulletBlock[] }) {
  if (bullets.length === 0) return null;
  return (
    <ul className={styles.bullets}>
      {bullets.map((b) => (
        <li key={b.id} id={b.id} data-seg-id={b.id} className={styles.bullet}>
          <Rich text={b.text} />
        </li>
      ))}
    </ul>
  );
}

function NumberPanel({ block }: { block: NumberBlock }) {
  return (
    <div id={block.id} data-seg-id={block.id} className={styles.numberPanel}>
      <div className={styles.numberPanelValue}>
        <Rich text={block.value} />
      </div>
      {block.caption !== undefined && (
        <div className={styles.numberPanelCaption}>
          <Rich text={block.caption} />
        </div>
      )}
    </div>
  );
}

function CostTable({ numbers }: { numbers: NumberBlock[] }) {
  return (
    <div className={styles.table} role="table">
      {numbers.map((n, i) => {
        const total = i === numbers.length - 1;
        return (
          <div
            key={n.id}
            id={n.id}
            data-seg-id={n.id}
            role="row"
            className={`${styles.tableRow} ${total ? styles.tableTotal : i % 2 === 0 ? styles.tableShaded : ""}`}
          >
            <span role="cell" className={styles.tableCaption}>
              <Rich text={n.caption} />
            </span>
            <span role="cell" className={styles.tableValue}>
              <Rich text={n.value} />
            </span>
          </div>
        );
      })}
    </div>
  );
}

function SlideImage({ block, assetBaseUrl }: { block: ImageBlock; assetBaseUrl: string }) {
  const [failed, setFailed] = useState(false);
  const src = `${assetBaseUrl.replace(/\/+$/, "")}/${block.assetKey}`;
  return (
    <div id={block.id} data-seg-id={block.id} className={styles.imageBox}>
      {failed ? (
        <div className={styles.imagePlaceholder} role="img" aria-label={block.alt}>
          <span>{block.alt}</span>
        </div>
      ) : (
        <img className={styles.image} src={src} alt={block.alt} onError={() => setFailed(true)} />
      )}
    </div>
  );
}
