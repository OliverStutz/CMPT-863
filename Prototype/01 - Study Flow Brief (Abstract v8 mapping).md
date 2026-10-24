# Study flow brief: mapping Oliver's prototype prompt onto Abstract v8

Written 2026-10-23. This is the shared brief for the architecture and content work. Abstract v8 is the authority on the study design; Oliver's prompt (see `00 - Original Prompt (2026-10-23).md`) is the authority on the application shape, timing and admin needs. Where the two disagree, the point is listed under "Open points" below and the architecture treats it as a configurable parameter with the Abstract v8 reading as default. Nothing under "Open points" is decided.

## 1. Fixed by Abstract v8 (not negotiable in the prototype)

- **Two groups**: checker group and neutral group. Same slides, same CISO narration, same answers. Only the checker panel/markings and who asks the question differ. Participants do not know their group.
- **Allocation**: round robin as sign-ups arrive, so groups stay equal in size. Done server-side at session creation. (Reviewers asked for random; Oliver has not changed this.)
- **Five CISO slides**, all error-free, identical in both groups. Checker group sees "No issues found" on slides 1 and 5 and a marked region with a flag on slides 2, 3 and 4.
- **Flags**: worded as observations about the slide, never about the CISO's honesty. No accuracy or confidence figure anywhere on the panel. Every flag is plausible but wrong for this company. Examples in v8: "Architecture likely wrong", "Numbers do not add up". Oliver's example: "Firewalls likely too expensive, cheaper ones available".
- **Per flagged slide (2, 3, 4)**: trust rating (after the slide/flag) -> action -> question is asked -> CISO answers -> trust rating (after the answer).
  - Checker group actions: say nothing / ask the CISO to explain / challenge the CISO with the checker's finding. If nothing, the board chair asks a routine question.
  - Neutral group action: say whether they have a question about the slide. The chair asks the routine question either way.
  - The CISO's answer is the same whoever asked. Two of the three answers give the company's reason; one is a denial with no reason. Which slide carries the denial is balanced across participants within each group (so: rotate 2, 3, 4 within each group, server-side).
  - The CISO is the same person, same tone, for reasons and denial. The CISO never mentions AI.
- **Trust measure at every rating point**: 0-to-10 trust rating plus one item each for competence, benevolence, integrity (Mayer 1995; wording after METI, Hendriks 2015). Four quick responses; participants should react, not deliberate.
- **Scenario**: participant sits on the board and decides the CISO's budget request. Amount and what depends on it are the same for both groups and shown on a slide. A CISO profile establishes experience before the first (baseline) rating. Checker group: the checker is introduced before the slides.
- **After the slides**: budget decision, then background questions.
- **Debrief**: no slide had a real issue, the flags were fixed in advance, the study measures how LLM flags affect non-experts' trust in the presenter.
- **Measures to log**: every trust rating; the action chosen at each flag; the budget decision; education, age group, field of expertise, self-rated cybersecurity knowledge, own AI use (the last two are planned moderators of H1 and H2).
- **Validity**: a rating after slide 1 (no flag) is needed to see whether the checker's presence alone moves trust. The groups' baseline ratings are compared.

## 2. Fixed by Oliver's prompt (application shape)

- Consent page with inline PDF, a checkbox that must be ticked before the "I accept the terms and would like to participate in the study" button is enabled.
- Passive capture of IP address and browser type (user agent) for bot detection.
- Setup step: confirm audio works, play a short example to test playback.
- Boardroom scene per slide: the slide, the CISO and the boardroom visible; the checker overlays markings on the slide "like AR glasses", in a configurable colour, with configurable flag text, appearing at a configurable time within the slide's audio.
- Audio is text-to-speech, generated in advance, identical for all participants. No replay. No back navigation.
- After the audio ends: 10 s to look at the slide, then automatic switch to the rating screen, which is open for 15 s.
- Every step (presentation start, audio finished, rating submitted or timed out) is reported to the backend as it happens, so progress is live and a client cannot post a fabricated full record.
- Slides, markings, flag texts, colours, timings: all configurable so the content can evolve.
- Last screen: debrief that there were no issues on the slides.
- Study runs with a specific identifier (link) that tells the app how to load; admin creates runs with a source string (LinkedIn, WhatsApp, ...).
- Admin participants tab: when they took it, progress k/N, status Started / Ongoing / Finished, elapsed time (live while running).
- Stack: TypeScript React app, Go backend on Cloud Run (GCP), static artefacts on Google Cloud CDN, Terraform for the full setup.

## 3. Canonical participant flow (v8 order, Oliver's mechanics)

| # | Screen | Group difference | Logged |
|---|---|---|---|
| 0 | Landing via run link; server creates session, assigns group (round robin) and denial slide (balanced within group), issues session token | none | ip, user agent, run, source, group, denial slide, created_at |
| 1 | Consent: inline PDF, checkbox, accept button | none | consent_accepted |
| 2 | Setup: audio check (play clip, participant confirms what they heard), playback check with a short example scene | Checker group: the example scene also introduces the checker panel (OPEN A) | setup_done |
| 3 | Opening: boardroom, chair opens the meeting, CISO profile, topic and budget amount. No markings | Checker group: panel present, "No issues found" or idle (OPEN A) | scene_started, audio_ended |
| 4 | Baseline rating T0 (trust 0-10, competence, benevolence, integrity), 15 s | none | rating |
| 5 | Slide 1: narration, no flag | Checker: panel "No issues found" | scene_started, audio_ended |
| 6 | Rating after slide 1 (OPEN B) | none | rating |
| 7 | Slide 2: narration; marking + flag appear at configured time | Neutral: no panel, no marking | scene_started, marking_shown, audio_ended |
| 8 | Rating after slide 2 (post-flag), 15 s | none | rating |
| 9 | Action: checker group chooses say nothing / ask to explain / challenge with the finding; neutral group answers "Do you have a question about this slide?" (OPEN C: time limit) | wording differs | action |
| 10 | Question is voiced: chair's routine question, or the participant's chosen question (OPEN D: how the participant's question is presented) | who asks | question_played |
| 11 | CISO answer audio: company's reason, or denial without reason on the participant's assigned denial slide | none | answer_ended |
| 12 | Rating after the answer, 15 s | none | rating |
| 13 | Slides 3 and 4: repeat steps 7 to 12 | | |
| 14 | Slide 5: narration, no flag | Checker: "No issues found" | scene_started, audio_ended |
| 15 | Rating after slide 5 (OPEN B) | none | rating |
| 16 | Budget decision (OPEN E: response format) | none | budget_decision |
| 17 | Background questions (OPEN F: item list and position) | none | background |
| 18 | Optional end-of-study items (OPEN G) | | |
| 19 | Debrief, thank you, finished | none | finished |

Progress k/N in the admin view counts the logged steps above for the participant's group.

## 4. Open points (ask Oliver; do not assume)

Each has a proposed default so that the architecture can be parameterised now.

- **OPEN A. Checker introduction.** v8 says the checker is introduced in step 1 (checker group only). Oliver's setup step plays an example to test playback. Proposal: the example scene is a 20 s neutral boardroom clip; for the checker group it additionally shows the panel and explains it in one sentence. Alternative: a separate text screen.
- **OPEN B. Ratings after slides 1 and 5.** v8 Method lists ratings after slides 2 to 4 only; v8 Validity needs a slide-1 reading; the literature review says "before and after every slide"; Oliver's prompt says every slide has a rating. Proposal: rate after every slide (T0 baseline, T1, T2a/T2b, T3a/T3b, T4a/T4b, T5 = 9 ratings). Alternative: skip T1 and T5 (7 ratings).
- **OPEN C. Action step timing.** Oliver's prompt gives 10 s view + 15 s rating but says nothing about the action choice. Proposal: 15 s for the action; on timeout record "no action" and treat as "say nothing" (chair asks). Alternative: untimed.
- **OPEN D. How the participant's question is voiced.** When the participant asks or challenges, is the question shown as text only, or spoken (TTS, a "board member" voice)? Proposal: shown as text in a speech bubble for 3 to 4 s, then the CISO answers; the chair's routine question is spoken by a chair voice.
- **OPEN E. Budget decision format.** v8: "budget decision". Proposal: approve the full request / approve a reduced amount / defer / reject, plus a 0-to-10 confidence. Alternative: approve / reject only.
- **OPEN F. Background questions.** Oliver's prompt asks field of expertise, country, age bracket, job level (Student, Employee, Leadership, C-Level) right after consent. v8 asks education, age group, field of expertise, self-rated cybersecurity knowledge, own AI use, after the budget decision; the last two are planned moderators. Proposal: ask all of them after the budget decision (so that the cybersecurity and AI items do not prime the participant), with Oliver's four plus v8's three. Alternative: Oliver's four up front, v8's three at the end.
- **OPEN G. End-of-study items.** v8 appendix borrows from Prike "the end-of-study rating of how true each checked claim is". Not in the Method steps. Include a screen where checker participants rate each of the three flags as correct/incorrect before the debrief? Proposal: include, checker group only, optional in config.
- **OPEN H. Rating timeout.** If 15 s expire without all four items: record what was given and mark the rest as timed out, or discard the whole rating? Proposal: record partial, mark missing items; the backend marks the rating as "timeout".
- **OPEN I. Scale for the three items.** v8: 0-to-10 for trust; competence, benevolence, integrity "worded after METI" (METI uses a 7-point semantic differential). Proposal: all four on 0-to-10 sliders for speed and consistency. Alternative: 7-point for the three items.
- **OPEN J. IP address storage.** IP is personal data. Proposal: store a salted hash plus country/ASN derived at request time; keep the raw IP only for 7 days in logs. Alternative: store raw IP (needs ethics wording).
- **OPEN K. Scene visuals.** "Showing the presentation, the CISO and the boardroom": is a static composite (boardroom photo/illustration, CISO portrait that lights up while speaking, slide rendered live, slide animations) acceptable, or does Oliver want generated video (for example an avatar video)? Proposal: static composite with lightweight animation; video is a later upgrade.
- **OPEN L. Voices and language.** Google Cloud Text-to-Speech, English. Which voices for CISO and chair (gender, accent)? Proposal: CISO one Neural2/Studio voice, chair a clearly different one; both fixed in config.
- **OPEN M. Consent PDF.** Who supplies it (SFU ethics form)? Placeholder until then.
- **OPEN N. Admin login.** Proposal: Google Identity-Aware Proxy in front of the admin routes, allowlist of Google accounts. Alternative: shared admin password.
- **OPEN O. Database.** Proposal: Firestore (serverless, free tier, fits Cloud Run). Alternative: Cloud SQL Postgres (easier SQL export, about CAD 15/month minimum).
- **OPEN P. Static hosting.** Google Cloud CDN needs a global external HTTPS load balancer (about USD 18/month fixed). Alternative: Firebase Hosting (CDN included, free tier) or serve the SPA from Cloud Run. Oliver asked for Cloud CDN; confirm the cost is acceptable.
- **OPEN Q. GCP project, region, domain.** Project ID, region (northamerica-northeast1 Montreal or us-west1 are closest to SFU), custom domain and DNS provider.
- **OPEN R. Devices.** Desktop only (minimum width enforced) or allow phones? Proposal: desktop and tablet landscape only; phones see a "please open on a larger screen" page.
- **OPEN S. Run-level config versioning.** Each run points at one content config version so results are comparable. Proposal: config is a versioned JSON file in the repo, uploaded to GCS; a run stores the version it uses.
- **OPEN T. "Slide 1-4" in the prompt.** Oliver wrote "Slide 1-4 with the Questions ahead of the start". Read as: the opening scene is before slide 1, slides 1 to 5 are the CISO's deck, and "questions ahead of the start" is the baseline rating. Confirm.
- **OPEN U. Attention or bot check.** Beyond IP and user agent: an audio check where the participant picks the word they heard, and a honeypot field on the background form. Include?
