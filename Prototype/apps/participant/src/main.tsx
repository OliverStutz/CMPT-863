import { createRoot } from "react-dom/client";
import { App } from "./App";
import "./global.css";

// No React.StrictMode: its doubled effects in development would send study events twice.
createRoot(document.getElementById("root")!).render(<App />);
