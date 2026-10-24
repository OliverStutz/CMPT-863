// Device gate message (2.9, section 9.1). Keeps the link visible so it can be opened later.

export function BlockedPage({ code }: { code: string }) {
  const link = `${window.location.origin}/s/${encodeURIComponent(code)}`;
  return (
    <main className="center-message">
      <h1>Please open this link on a computer or a tablet in landscape.</h1>
      <p>This study needs a screen at least 1,024 pixels wide.</p>
      <p>
        Your link: <span className="link-text">{link}</span>
      </p>
    </main>
  );
}
