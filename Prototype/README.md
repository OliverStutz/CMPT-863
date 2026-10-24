# Prototype: automated browser study

Started 2026-10-23. The prototype runs the study described in `CMPT 863 Project Abstract v8.md` as a self-contained web application: a participant follows a link, consents, hears a CISO present five slides in a simulated boardroom, rates trust at each step, and the backend records every step live.

## Files

| File | What it is | Status |
|---|---|---|
| `00 - Original Prompt (2026-10-23).md` | Oliver's prompt, verbatim. Authority on the application shape, timing and admin needs | Fixed |
| `01 - Study Flow Brief (Abstract v8 mapping).md` | Maps the prompt onto Abstract v8. Fixed constraints, canonical flow, OPEN points A to U | Superseded for decisions by `04`; still the flow reference |
| `02 - Architecture.md` | Technical architecture: React/TS front end, Go on Cloud Run serving app, assets and audio (no CDN, no load balancer), Firestore, shared admin password, raw IP stored, Terraform with region variable (default Montreal), TTS pipeline with a board-member voice, state machine, API, types. Section 18 is the change log: cross-check fixes, then Oliver's decisions of 2026-10-23 and 2026-10-24 | Revised 2026-10-24 |
| `03 - Study Content v1.md` | First content draft: scenario, scripts, five slides, flags, answers, screen wording | Superseded by v1.1, kept for comparison |
| `03 - Study Content v1.1.md` | Content after the cross-check: slide 3 and 4 giveaways removed, routine questions ask for a justification, Oliver's firewall label restored, METI wording checked against Hendriks 2015, scenario facts tabled for confirmation | Superseded by v1.2 |
| `03 - Study Content v1.2.md` | v1.1 plus Oliver's decisions of 2026-10-23 | Superseded by v1.3 |
| `03 - Study Content v1.3.md` | All 53 decisions applied: Meridian Logistics (neutral name), five-stage deck (where we stand, risk and impact, how we protect, what it costs, the decision), narrations about 110 words, denials at reason length where credible, 7-point METI items, voiced participant questions, no role screen, no flag review, checker sentence on the last setup screen. Flags are placeholders until the VLM run is folded in | Current |
| `04 - Open Decisions.md` | The 53 points with options, recommendation and Oliver's answer each (eight on 2026-10-23, the rest on 2026-10-24). Ends with the table of facts the drafts set | All answered |
| `05 - Progress.md` | Build and deployment state in one place: component table, deployment facts, security rules in force, what is still needed, dated log | Kept current |
| `content/schema/study-config.schema.json` | JSON Schema (draft 2020-12) mirroring the types in `02` section 4.1; the single source of truth for the content file shape | Current |
| `content/study-config.v1.json` | First machine-readable content; does not match the schema | Superseded |
| `content/study-config.v1.1.json` | Content v1.1 encoded to the schema as it was before the slide-count and checker-intro changes; fails the final schema at three paths | Superseded |
| `content/study-config.v1.2.json` | Content v1.2; fails the schema after the 2026-10-24 decisions | Superseded |
| `content/study-config.v1.3.json` | Content v1.3 encoded to the final schema, validated; flags are placeholders | Current once present |
| `content/slides/v1.3/` | The five slides as SVG (segment ids as element ids) and 1920 by 1080 PNG, plus the generator script and the render command | Current |
| `content/vlm-drafts/v1.3/` | Provenance of the VLM drafting run: fixed prompt (written before any model saw the slides), input hashes, raw output per slide, run record. A fresh Claude Opus agent is the vision-language model | Current |
| `reports/crosscheck-v1.md` | Independent review of `02` and `03 v1` against Abstract v8: 66 findings, merged open-point list | Done; all technical findings applied, content findings applied except those needing Oliver |

## Rules for this folder

- Abstract v8 decides the study design. The prompt decides the application. Where they disagree, the point is OPEN until Oliver decides; see `04`.
- No em-dashes in documents. Plain wording.
- Content is versioned. A study run always points at one config version. The schema is the contract between content and code.
- Flags are worded about the slide, never about honesty, and carry no confidence or accuracy figure. The CISO never mentions AI. A lint step enforces both.
- The debrief sentence about flags drafted with a real vision-language model becomes true once the flags in the content are taken from `content/vlm-drafts/v1.3/` and the edits are logged there.
