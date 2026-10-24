import { useEffect, useState } from "react";
import QRCode from "qrcode";

// DECISION: the run link is built as <origin>/s/<code> from the admin page's own origin, because the admin SPA
// is served by the same service as the participant app (section 3.4 of the change log) and GET /runs carries no url.
export function runUrl(code: string): string {
  return `${window.location.origin}/s/${code}`;
}

export function CopyButton({ text }: { text: string }) {
  const [copied, setCopied] = useState(false);
  async function copy() {
    try {
      await navigator.clipboard.writeText(text);
    } catch {
      const ta = document.createElement("textarea");
      ta.value = text;
      document.body.appendChild(ta);
      ta.select();
      document.execCommand("copy");
      ta.remove();
    }
    setCopied(true);
    setTimeout(() => setCopied(false), 1500);
  }
  return (
    <button className="small" onClick={() => void copy()}>
      {copied ? "Copied" : "Copy"}
    </button>
  );
}

/** QR code generated in the browser, downloadable as PNG (section 10). */
export function RunQr({ code }: { code: string }) {
  const url = runUrl(code);
  const [png, setPng] = useState<string | null>(null);
  useEffect(() => {
    let alive = true;
    QRCode.toDataURL(url, { width: 320, margin: 2, errorCorrectionLevel: "M" }).then(
      (d) => alive && setPng(d),
      () => alive && setPng(null),
    );
    return () => {
      alive = false;
    };
  }, [url]);
  return (
    <div className="qr">
      {png ? <img src={png} width={160} height={160} alt={`QR code for ${url}`} /> : <span>Generating QR code...</span>}
      <div>
        <div className="mono">{url}</div>
        <div className="row-actions">
          <CopyButton text={url} />
          {png && (
            <a className="button small" href={png} download={`run-${code}-qr.png`}>
              Download PNG
            </a>
          )}
        </div>
      </div>
    </div>
  );
}
