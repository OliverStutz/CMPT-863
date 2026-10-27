import { useEffect, useRef, type ReactNode } from "react";

/**
 * A heading that takes focus when it appears, for screen readers (section 9.6). Used inside the boardroom
 * stage, where the step component's own effects run before the host has drawn the step's card, so the
 * heading focuses itself on mount instead.
 */
export function AutoFocusHeading({ className, children }: { className?: string; children: ReactNode }) {
  const ref = useRef<HTMLHeadingElement>(null);
  useEffect(() => {
    ref.current?.focus();
  }, []);
  return (
    <h1 ref={ref} tabIndex={-1} className={className}>
      {children}
    </h1>
  );
}
