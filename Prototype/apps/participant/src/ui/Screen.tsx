import { useEffect, useRef, type ReactNode } from "react";
import styles from "./Screen.module.css";

/** A centred page for form screens. Focus moves to the heading on mount, for screen readers. */
export function Screen({ title, titleHidden, children, wide }: { title: string; titleHidden?: boolean; children: ReactNode; wide?: boolean }) {
  const h = useRef<HTMLHeadingElement>(null);
  useEffect(() => {
    h.current?.focus();
  }, []);
  return (
    <main className={styles.page}>
      <div className={`${styles.card} ${wide ? styles.wide : ""}`}>
        <h1 ref={h} tabIndex={-1} className={titleHidden ? "sr-only" : styles.title}>
          {title}
        </h1>
        {children}
      </div>
    </main>
  );
}
