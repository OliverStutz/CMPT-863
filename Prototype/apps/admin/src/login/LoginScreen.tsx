import { useState, type FormEvent } from "react";
import { ApiError, login } from "../api/client";

export function LoginScreen() {
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [busy, setBusy] = useState(false);

  async function submit(e: FormEvent) {
    e.preventDefault();
    setBusy(true);
    setError(null);
    try {
      await login(password);
      setPassword("");
    } catch (err) {
      if (err instanceof ApiError && err.status === 401) setError("Wrong password.");
      else if (err instanceof ApiError && err.status === 429) setError("Too many attempts. Wait a minute and try again.");
      else setError(err instanceof Error ? err.message : "Login failed.");
    } finally {
      setBusy(false);
    }
  }

  return (
    <div className="overlay">
      <form className="login" onSubmit={submit}>
        <h2>Admin login</h2>
        <label>
          Password
          <input
            type="password"
            autoFocus
            autoComplete="current-password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </label>
        {error && <p className="error">{error}</p>}
        <button type="submit" disabled={busy || password === ""}>
          {busy ? "Logging in..." : "Log in"}
        </button>
      </form>
    </div>
  );
}
