/**
 * Shown when the access-key gate answers 404 while the app is open: the gate cookie is gone
 * (30 days without opening the page, or the key was rotated). The password login alone
 * cannot help here, so the app says what to do instead of showing a dead page.
 */
export function GateLostScreen() {
  return (
    <div className="overlay">
      <div className="login" role="alertdialog" aria-labelledby="gate-lost-title">
        <h2 id="gate-lost-title">Access link needed</h2>
        <p>
          Access to the admin page from this browser has expired, or the access key was changed. Open the admin
          link that carries the access key again (the one with <code>?k=</code>), then log in with the password.
        </p>
      </div>
    </div>
  );
}
