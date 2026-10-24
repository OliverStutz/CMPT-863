// Consent (section 9.2 `<ConsentPage>`): the PDF inline in an iframe with a download link as fallback,
// a checkbox, and the accept button, disabled until the box is ticked. The PDF is a placeholder until
// the SFU ethics form arrives (OPEN M); if it is missing, a placeholder box is shown instead.

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
  const CONSENT_PDF_URL = consentPdfUrl(pack.consent.pdfAssetKey);

  useEffect(() => {
    let live = true;
    fetch(CONSENT_PDF_URL, { method: "HEAD", credentials: "include" })
      .then((r) => {
        const type = r.headers.get("content-type") ?? "";
        if (live) setPdf(r.ok && type.includes("pdf") ? "ok" : "missing");
      })
      .catch(() => live && setPdf("missing"));
    return () => {
      live = false;
    };
  }, [CONSENT_PDF_URL]);

  const accept = async () => {
    setBusy(true);
    try {
      await api.sendUntilOnTime("consent_accepted", { consentVersion: pack.consent.version, checkbox: true });
    } catch (e) {
      setBusy(false);
      reportError(e);
    }
  };

  return (
    <Screen title="Consent" titleHidden wide>
      {pdf === "ok" ? (
        <>
          <iframe className={styles.pdfFrame} src={CONSENT_PDF_URL} title="Consent form" />
          <p className={styles.muted}>
            <a href={CONSENT_PDF_URL} target="_blank" rel="noopener" download>
              Download the consent form (PDF)
            </a>
          </p>
        </>
      ) : (
        <div className={styles.pdfPlaceholder} data-testid="consent-placeholder">
          {pdf === "checking" ? "" : "The consent PDF goes here."}
        </div>
      )}
      <label className={styles.checkboxRow}>
        <input type="checkbox" checked={checked} onChange={(e) => setChecked(e.target.checked)} />
        <span>{pack.consent.checkboxLabel}</span>
      </label>
      <div className={styles.actions}>
        <button type="button" className={styles.primary} disabled={!checked || busy} onClick={accept}>
          {pack.consent.acceptLabel}
        </button>
      </div>
    </Screen>
  );
}
