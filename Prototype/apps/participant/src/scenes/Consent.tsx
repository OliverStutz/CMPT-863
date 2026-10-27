// Consent (section 9.2 `<ConsentPage>`): the SFU consent form (study 30004094, version 4, May 2, 2025,
// served from content/consent/) inline in an iframe with a download link, a checkbox, and the accept
// button. The button is disabled until the box is ticked and the form has been confirmed present
// (HEAD on the PDF). If the form cannot be loaded the page says so and offers "Try again"; nobody can
// consent to a form they have not been shown.

import { useEffect, useState } from "react";
import type { Step } from "@study/shared";
import { useStudy } from "../study";
import { Screen } from "../ui/Screen";
import { consentPdfUrl } from "../assets";
import styles from "./scenes.module.css";

export function Consent({ step: _step }: { step: Step }) {
  const { api, pack, reportError } = useStudy();
  const [checked, setChecked] = useState(false);
  const [busy, setBusy] = useState(false);
  const [pdf, setPdf] = useState<"checking" | "ok" | "missing">("checking");
  const [attempt, setAttempt] = useState(0);
  const CONSENT_PDF_URL = consentPdfUrl(pack.consent.pdfAssetKey);

  useEffect(() => {
    let live = true;
    setPdf("checking");
    fetch(CONSENT_PDF_URL, { method: "HEAD", credentials: "include" })
      .then((r) => {
        const type = r.headers.get("content-type") ?? "";
        if (live) setPdf(r.ok && type.includes("pdf") ? "ok" : "missing");
      })
      .catch(() => live && setPdf("missing"));
    return () => {
      live = false;
    };
  }, [CONSENT_PDF_URL, attempt]);

  const accept = async () => {
    setBusy(true);
    try {
      await api.sendUntilOnTime("consent_accepted", { consentVersion: pack.consent.version, checkbox: true });
    } catch (e) {
      setBusy(false);
      reportError(e);
    }
  };

  // Layout: one fixed-height frame area (the same height while the PDF is checked, when it is shown and
  // when the notice is shown), then the checkbox and the controls row. Nothing above the checkbox
  // changes size when the PDF check finishes, so a click is never moved onto the PDF viewer, and the
  // frame height leaves the checkbox and the button inside the window. The PDF scrolls inside its frame.
  return (
    <Screen title="Consent" titleHidden wide>
      <div className={styles.pdfArea} data-testid="consent-pdf-area">
        {pdf === "ok" ? (
          <iframe className={styles.pdfFrame} src={CONSENT_PDF_URL} title="Consent form" data-testid="consent-pdf" />
        ) : (
          <div className={styles.pdfPlaceholder} data-testid="consent-placeholder" role={pdf === "missing" ? "alert" : undefined}>
            {pdf === "missing" ? (
              <div>
                <p>The consent form could not be loaded.</p>
                <p>Please check your connection and try again. You cannot take part without it.</p>
                <button type="button" className={styles.secondary} onClick={() => setAttempt((n) => n + 1)}>
                  Try again
                </button>
              </div>
            ) : (
              ""
            )}
          </div>
        )}
      </div>
      <div className={styles.consentControls}>
        <label className={styles.checkboxRow}>
          <input type="checkbox" checked={checked} onChange={(e) => setChecked(e.target.checked)} />
          <span>{pack.consent.checkboxLabel}</span>
        </label>
        <div className={styles.actions}>
          <button type="button" className={styles.primary} disabled={!checked || busy || pdf !== "ok"} onClick={accept}>
            {pack.consent.acceptLabel}
          </button>
          <a
            className={styles.muted}
            href={CONSENT_PDF_URL}
            target="_blank"
            rel="noopener"
            download
            style={{ visibility: pdf === "ok" ? "visible" : "hidden" }}
            aria-hidden={pdf === "ok" ? undefined : true}
            tabIndex={pdf === "ok" ? undefined : -1}
          >
            Download the consent form (PDF)
          </a>
        </div>
      </div>
    </Screen>
  );
}
