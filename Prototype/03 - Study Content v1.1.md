# Study content v1.1: slides, flags, answers and scripts

Revision written 2026-10-23 from `03 - Study Content v1.md` after the independent review `reports/crosscheck-v1.md`. Sources: `CMPT 863 Project Abstract v8.md` (the authority) and `01 - Study Flow Brief (Abstract v8 mapping).md`. v1 is left unchanged. The JSON is not updated yet; it will be rebuilt from this file once the new schema is ready. Nothing marked OPEN is decided: the brief's proposed default is used and marked. Every fact and default I set is listed for confirmation in section 13. The changes from v1 are listed in section 14.

## 1. Scenario

**Company.** Kestrel Bay Logistics (fictional). A freight and warehousing company in western Canada: head office and main data centre in the Vancouver area, 12 warehouses, 1,400 staff, about $480 million revenue a year. Everything runs on one network: orders, warehouse scanners, truck routes, payroll.

**Meeting.** The quarterly board meeting, agenda item 4: "the protection of our network". The participant sits on the board.

**Budget request.** $2.4 million (CAD), identical for both groups. What depends on it: the company's cyber insurance is renewed in 10 months only if stronger protection is in place, and the company's largest customer has written the same requirement into its new contract. If the board waits, the premium rises or cover ends, and the customer contract is at risk.

Where the amount appears: on **slide 1** (the ask, so that every rating and every flag is judged against a known amount), in the title of **slide 3** (the breakdown), and on **slide 5** (the decision, right before the budget vote). The chair also says it once in the opening. A real board deck does the same: state the ask up front, restate it at the decision.

**CISO.** Daniel Hale, Chief Information Security Officer (fictional). Pronoun "he", as in the Abstract's appendix ("our CISO answers for himself"). Profile, 4 sentences after the introduction:

> Daniel has led our security for six years. Before he joined us, he spent twelve years protecting the networks of a national bank and a large hospital group. This year he led our response to two attempted break-ins, and both were stopped before any harm was done. He holds the main professional certifications in his field.

**Proposal for the profile** (R-6, open): the chair reads it aloud in the opening scene, and a profile card with the same four points is shown on screen while she speaks. Spoken keeps it in the audio-led flow and ensures every participant receives it before the baseline rating; the card helps participants who listen less well.

**Board chair.** Margaret Lin, chair of the board (fictional). Her voice must be clearly different from the CISO's (OPEN L).

The CISO never mentions AI, checkers or tools in any line.

All scenario facts above are my choices, not the Abstract's. They are listed for confirmation in section 13 (R-5).

### Role instruction (both groups, shown before the opening scene)

> You are a member of the board of Kestrel Bay Logistics.
> Today the company's head of security asks the board to approve a budget. Listen to the presentation.
> During the meeting you will rate the presenter several times, and at the end you will vote on the budget.

This screen is my proposal (C-1).

## 2. Opening scene script

No markings. Checker group: the panel is visible but idle, with no label (OPEN A). A profile card is on screen during line 2. Total 160 words, about 64 seconds.

| # | Speaker | Line | Audio key |
|---|---|---|---|
| 1 | Chair | Good morning, everyone, and thank you for being here. Let us move to item four on our agenda: the protection of our network. | `opening.chair.1` |
| 2 | Chair | Presenting today is Daniel Hale, our Chief Information Security Officer. Daniel has led our security for six years. Before he joined us, he spent twelve years protecting the networks of a national bank and a large hospital group. This year he led our response to two attempted break-ins, and both were stopped before any harm was done. He holds the main professional certifications in his field. | `opening.chair.2` |
| 3 | Chair | Daniel is asking the board for two point four million dollars to strengthen the protection of our warehouses and our head office. Our insurer and our largest customer both expect this protection to be in place within the year. We will hear his five slides and ask our questions along the way. The board will decide on the budget today. | `opening.chair.3` |
| 4 | Chair | Daniel, the floor is yours. | `opening.chair.4` |
| 5 | CISO | Thank you, Margaret. Good morning, everyone. | `opening.ciso.5` |

Line 5 lets participants hear the CISO's voice once before the baseline rating (T0). Slide 1 then starts without a second greeting. Whether the CISO speaks before T0 is open (R-6).

**Profile card (on screen during line 2):**

- Daniel Hale, Chief Information Security Officer, Kestrel Bay Logistics
- Head of security at Kestrel Bay Logistics for 6 years
- 12 years before that protecting the networks of a national bank and a large hospital group
- Led the response to two attempted break-ins this year, both stopped before any harm
- Holds the main professional certifications in his field

## 3. Checker introduction (checker group only, OPEN A, R-1, FO-2)

Three questions about the introduction are open. This section gives content for every option and decides none of them.

- **OPEN A, the form.** Brief default: the example scene of the setup step also shows the panel and explains it in one sentence. Alternative: a separate text screen.
- **R-1, the position.** Before the baseline rating T0 (in setup, as the brief's flow table places it) or after T0 and before slide 1 (the Abstract's order: "First trust rating. Checker group: the checker is introduced."). Introducing an AI tool before T0 may shift the checker group's baseline, which the Validity check ("the groups start alike on the first rating") is meant to rule out. Open.
- **FO-2, the display.** The checker marks the slide directly, like AR glasses (Oliver's prompt), or a panel shows the camera image of each slide (Abstract). Open. The texts below do not say glasses or panel; an optional extra line per display mode is given.

**Option 1, brief default (OPEN A): one sentence in the example scene.** Shown as a caption on screen while the example scene plays, not spoken, so that neither the chair nor the CISO introduces the tool:

> In this meeting, you have a slide checker: an AI tool that reads each slide and marks what it finds.

Note: this option ties the introduction to the setup step, so it places it before T0. If R-1 is decided as "after T0", Option 1 needs a short reminder screen after T0, or Option 2 is used.

**Option 2, alternative (OPEN A): a separate text screen.** Button: "Continue". The text works in either position (before or after T0): it does not mention the example scene, the rating, or whether the meeting has started.

> During this meeting you have a slide checker. The checker is an AI tool.
> While the presenter speaks, it reads each slide and marks what it finds.
> If it finds a possible issue, it marks that part of the slide and shows a short note.
> If it finds none, it shows: No issues found.

Optional extra line, keyed to the display mode (FO-2), added after the second line:

- Overlay mode: "You see its marks directly on the slide."
- Panel mode: "A screen in front of you shows the camera image of each slide, with the checker's marks."

Neither text says anything about how accurate the checker is, in line with "no accuracy or confidence figure".

**No-issue label:** "No issues found" (Abstract v8). Configurable as `labels.noIssueLabel`. The v7-to-v8 change file records that Oliver has not decided between this wording and a supporting label, so this is a default, not a decision.

**Panel title:** "Slide checker" (`labels.panelTitle`). Used in panel mode; in overlay mode it is the heading of the short note.

## 4. The five slides

Order: slide 1 context and the ask, slide 2 architecture, slide 3 costs, slide 4 timeline and risk, slide 5 summary and decision. This is the order of a typical board approval deck (why, what, how much, when, decide), and it puts the three flags on three different kinds of content: a design point, a price point and an arithmetic point.

All five slides are correct. The numbers add up: $780,000 + $960,000 + $310,000 + $350,000 = $2,400,000. Phases run months 1 to 2, 3 to 7 and 5 to 8, so the work is done in 8 months although the durations sum to 11.

**Flag drafting status.** PLACEHOLDER. The three flags below were written by the content author, not drafted by a VLM. Each is a rule of thumb, a price comparison or a sum that a vision-language model could plausibly produce from the slide image, and each is wrong for this company for a reason the CISO states in two short sentences. The Abstract says the flags "were drafted with a real VLM from images of the slides, then fixed". That is true only after the run described in "VLM drafting run (to do)" below. Until then, the flag texts are placeholders and the debrief sentence in section 10 is marked PLACEHOLDER.

None of the flags is about honesty. None carries a confidence or accuracy figure.

### VLM drafting run (to do)

Needed so that the Abstract's Method and the checker group's debrief are true (finding 1, R-10).

1. **Render the slides.** Build the five slides as real images (the network diagram on slide 2 and the timeline on slide 4 are still descriptions) and render each slide to PNG, exactly as participants will see it.
2. **Run a real VLM.** Run one vision-language model on each PNG with one fixed prompt per slide. The prompt asks for possible issues with the slide's content and the part of the slide each one concerns. The same model and settings are used for all five slides. Which model is open (C-14).
3. **Store the provenance** under `Prototype/content/vlm-drafts/`: the five PNGs, the model name and version, the date of the run, the settings, the exact prompt for each slide, and the raw output for each slide, unedited.
4. **Choose and edit.** Choose the flag for slides 2, 3 and 4 from that output and edit it into a short label and a one-sentence detail. Each flag must still be an observation about the slide, never about honesty, with no confidence figure, and wrong for this company for a reason the CISO can state.
5. **List every edit.** In the same folder, keep a list with, per slide: the model's original text, the final label and detail, and the reason for each change.
6. **Update the content.** Replace the placeholder flags in this file and in the JSON, then re-check that each reasoned answer still answers its flag.

If the model's output contains nothing close to a plausible but wrong flag for a slide, what to do is open (C-14).

### Answer design

The CISO's answer must fit three different questions: the chair's routine question, "explain" and "challenge". In v1 the routine question was a mild probe while the answer defended against a challenge, so the same denial sounded more evasive after the chair than after a challenge (findings 12, 50). In v1.1 all three questions ask for a justification:

- **Chair, routine question:** a generic "why did you choose this" question about the slide as a whole. It does not point at the flagged element.
- **Participant, explain:** asks why the flagged element is as it is.
- **Participant, challenge:** quotes the checker's flag and asks for a reply.

So the reasoned answer answers all three, and the denial fails all three alike. The reasoned answer and the denial share the same first sentence and the same last sentence. Only the middle, the reason, differs:

- Reasoned: [claim about the slide]. [reason, 2 sentences]. That is the plan I recommend.
- Denial: [claim about the slide]. That is the plan I recommend.

The closing line "That is the plan I recommend." is neutral: it fits a routine question as well as a challenge. The narration never states the reason that the answer gives, so that the denial truly carries no information. Whether the routine question is generic (v1.1 default) or slide-specific (v1) stays open (R-4); the v1 texts are kept below as the alternative.

**Marking times** are estimated at 150 words per minute from an anchor phrase in the narration and rounded to 0.5 s. After TTS, the timing should be re-set from the real audio (the JSON stores `anchorText` for that).

### Slide 1: Why we are here today

**(a) Layout**

| Id | Kind | Text |
|---|---|---|
| `s1.title` | title | Why we are here today |
| `s1.num1` | number | $2.4 million: budget request for approval today |
| `s1.b1` | bullet | Our 12 warehouses and head office run on one network: orders, scanners, trucks and payroll |
| `s1.b2` | bullet | This year we stopped two attempted break-ins into that network |
| `s1.b3` | bullet | Our insurer requires stronger protection before renewal in 10 months |
| `s1.b4` | bullet | Our largest customer requires the same in its new contract |

**(b) Narration** (CISO, `s1.narration`, 165 words, about 66 s)

> I will keep this short and clear, because I am asking you for a decision today. Everything we do runs on one network. Orders come in, the scanners in our twelve warehouses track every pallet, our trucks get their routes, and our fourteen hundred people get paid. If that network stops, the company stops. This year, we stopped two attempted break-ins into that network. Both were caught early, and no harm was done. But attackers are getting better, and we need to stay ahead of them. Two outside parties have also told us what they expect. Our insurer will renew our cyber insurance in ten months only if stronger protection is in place. And our largest customer has written the same requirement into its new contract. So today, I am asking for two point four million dollars to put that protection in place. On the next four slides, I will show you what we will build, what it costs, and when it will be ready.

**(c) Checker panel:** no marking. "No issues found" appears at 8,000 ms (my proposal, C-3). No question or answer step.

### Slide 2: How we will protect our 13 sites

**(a) Layout**

| Id | Kind | Text |
|---|---|---|
| `s2.title` | title | How we will protect our 13 sites |
| `s2.img1` | image | Network diagram: head office with the main data centre in the middle; 12 warehouse icons around it, each joined to head office by a line; a firewall shield at every site; one line from head office to a cloud labelled Internet; a laptop icon labelled Remote staff joined to head office by a dotted line |
| `s2.b1` | bullet | A firewall at every site keeps warehouse machines apart from office computers |
| `s2.b2` | bullet | All 12 warehouses connect to our data centre at head office |
| `s2.b3` | bullet | One internet connection, at head office, watched day and night |
| `s2.b4` | bullet | Staff at home or on the road use secure remote access |

**(b) Narration** (CISO, `s2.narration`, 141 words, about 56 s)

> Let me show you how the network will look. Picture our head office in the middle of this diagram. Our main data centre sits there. All twelve warehouses connect to that data centre. That is the line you see from each warehouse to the centre. At every site, we put a firewall. Think of a firewall as a locked door with a guard. It decides what may come in and go out, and it keeps the warehouse machines apart from the office computers. So if one computer is infected, the problem stays in that room. Then we have one connection to the internet, at head office, and our security team watches it day and night. Finally, people who work from home or on the road connect through secure remote access. They get in only after we have checked who they are.

v1's last sentence, "This is a simple, proven design, and it fits the way we work.", is deleted (finding 4). It gave a general reason before the flag.

**(c) Marking and answers**

- **Target:** `s2.b2` and `s2.img1` (the lines from the warehouses to head office).
- **Appears at:** 10,000 ms (anchor: "All twelve warehouses connect").
- **Colour token:** `flag`.
- **Flag label** (PLACEHOLDER until the VLM run): Architecture likely wrong
- **Flag detail** (PLACEHOLDER until the VLM run): All 12 warehouses connect through head office, while common practice is to connect each site straight to cloud services.
- **Why it is wrong for this company:** the textbook rule (send branch traffic straight to cloud services) assumes the systems live in the cloud. Kestrel Bay's warehouse systems run in its own data centre at head office.
- **Chair's routine question** (`s2.question.routine`, R-4 default): Thank you, Daniel. Before we move on, can you tell us why you chose this design?
- **R-4 alternative (v1, slide-specific):** Thank you, Daniel. Before we move on, could you say a little more about how the warehouses connect?
- **Participant, ask to explain:** Could you explain why all the warehouses connect through head office?
- **Participant, challenge with the finding:** My checker flags this slide: "Architecture likely wrong." What do you say to that?
- **CISO, reasoned answer** (`s2.answer.reason`, 53 words): This design is the right one for us. Our warehouse systems run in our own data centre at head office, not in the cloud, so every scan goes there anyway. A direct link to the cloud at each warehouse would only add twelve more doors to guard. That is the plan I recommend.
- **CISO, denial without reason** (`s2.answer.denial`, 14 words): This design is the right one for us. That is the plan I recommend.

Remaining limitation: the diagram and the narration place the main data centre at head office, which is half of the reason. This is a fact of the layout and is kept; see C-17.

### Slide 3: What the $2.4 million pays for

**(a) Layout**

| Id | Kind | Text |
|---|---|---|
| `s3.title` | title | What the $2.4 million pays for |
| `s3.num1` | number | Firewalls for 13 sites, with 3 years of support: $780,000 |
| `s3.num2` | number | 24-hour monitoring of our network, 2 years: $960,000 |
| `s3.num3` | number | Secure remote access for 1,400 staff: $310,000 |
| `s3.num4` | number | Installation and staff training: $350,000 |
| `s3.num5` | number | Total: $2,400,000 |

Shown as a simple cost table, one row per number block, total in bold.

**(b) Narration** (CISO, `s3.narration`, 141 words, about 56 s)

> Now to the money, and where the two point four million dollars goes. The first line is the firewalls for all thirteen sites, with three years of support: seven hundred and eighty thousand dollars. The largest line is round-the-clock monitoring for two years: nine hundred and sixty thousand dollars. For that, a team of specialists watches our network every hour of every day, and they can close a door within minutes if something looks wrong. We could not build that team ourselves for this price. Next is secure remote access for all fourteen hundred staff: three hundred and ten thousand dollars. And finally, installation and staff training: three hundred and fifty thousand dollars. Most break-ins start with a fake email, so training our people is some of the best value on this page. Together, that is two point four million dollars.

v1's last sentence, "Our finance team has checked every line, and each one is backed by a written quote from a supplier.", is deleted (finding 4). It pre-empted the price flag in both groups and gave the denial an implicit reason.

**(c) Marking and answers**

- **Target:** `s3.num1`.
- **Appears at:** 11,000 ms (anchor: "seven hundred and eighty thousand dollars").
- **Colour token:** `flag`.
- **Flag label** (Oliver's wording, from his prompt; PLACEHOLDER until the VLM run): Firewalls likely too expensive, cheaper ones available
- **Flag detail** (PLACEHOLDER until the VLM run): Firewalls for small offices cost about $5,000 each, so 13 units should cost about $65,000, not $780,000.
- **Why it is wrong for this company:** the check compares with office firewalls. A warehouse with hundreds of scanners needs a larger model, and the price includes support and next-day replacement.
- **Chair's routine question** (`s3.question.routine`, R-4 default): Thank you, Daniel. Before we move on, can you tell us why you chose to spend the money this way?
- **R-4 alternative (v1, slide-specific):** Thank you, Daniel. The firewalls are a large line. Could you say a little more about that cost?
- **Participant, ask to explain:** Could you explain why the firewalls cost $780,000?
- **Participant, challenge with the finding:** My checker flags this slide: "Firewalls likely too expensive, cheaper ones available." What do you say to that?
- **CISO, reasoned answer** (`s3.answer.reason`, 51 words): This is the right way to spend the money. Take the firewalls: an office model would not cope with our warehouses, where hundreds of scanners run day and night. The price also covers next-day replacement at every site, so a warehouse is never left unprotected. That is the plan I recommend.
- **CISO, denial without reason** (`s3.answer.denial`, 15 words): This is the right way to spend the money. That is the plan I recommend.

The first sentence changed from v1's "These firewalls are the right ones for us." so that it also answers the chair's question about the whole slide. "Take the firewalls" then turns to the flagged line.

### Slide 4: Ready in 8 months, before the insurance renewal

**Choice for finding 3.** In v1 the chart drew the phase bars, so the overlap of phases 2 and 3 could be seen at a glance. Any participant could refute the flag alone, and a participant who got the denial on slide 4 had the reason on screen. v1.1 keeps the arithmetic flag and changes the slide: the phase bars are removed, and the start and end month of each phase are given only in a small-type schedule line at the foot of the slide. The slide stays correct and complete, but the overlap has to be worked out from three month ranges. The line "You can see all three on the chart." is removed from the narration.

Remaining limitation: an attentive participant who reads the small type can still work out the overlap, so the slide 4 denial is not fully free of a reason on screen. This should be reported as a limitation. Replacing the flag with one whose resolution is not on the slide at all is the alternative (C-16).

**(a) Layout**

| Id | Kind | Text |
|---|---|---|
| `s4.title` | title | Ready in 8 months, before the insurance renewal |
| `s4.chart1` | chart | Milestone line, months 1 to 10 on the horizontal axis, with two markers only: month 8 labelled Ready, month 10 labelled Insurance renewal. No phase bars |
| `s4.b1` | bullet | Phase 1, design and ordering: 2 months |
| `s4.b2` | bullet | Phase 2, firewalls installed at all 13 sites: 5 months |
| `s4.b3` | bullet | Phase 3, monitoring set up and staff trained: 4 months |
| `s4.num1` | number | Ready in 8 months, renewal in 10 months |
| `s4.num2` | number | One day of stopped warehouses: about $1.2 million in lost shipments |
| `s4.note1` | note (small type, foot of slide) | Schedule: phase 1, months 1 to 2; phase 2, months 3 to 7; phase 3, months 5 to 8 |

**(b) Narration** (CISO, `s4.narration`, 138 words, about 55 s)

> So when will this be ready? The work comes in three phases. Phase one, design and ordering, takes two months. Phase two, installing the firewalls at all thirteen sites, takes five months. Phase three, setting up the monitoring and training our staff, takes four months. Altogether, we are ready in eight months. Our insurance renewal is in ten months, so that gives us two months of margin. Why does the timing matter? Because the risk is real, and it is expensive. If our warehouses stop for a single day, we lose about one point two million dollars in shipments. That is half of this request, gone in one day. And an attack rarely stops things for only one day. If we start today, we are protected well before the renewal. If we wait another quarter, we miss it.

**(c) Marking and answers**

- **Target:** `s4.b1`, `s4.b2`, `s4.b3` and `s4.num1`.
- **Appears at:** 18,000 ms (anchor: "Altogether, we are ready in eight months"; 45 words before the anchor, v1 had 52).
- **Colour token:** `flag`.
- **Flag label** (PLACEHOLDER until the VLM run): Numbers do not add up
- **Flag detail** (PLACEHOLDER until the VLM run): The three phases add up to 11 months, but the slide says the work is done in 8 months.
- **Why it is wrong for this company:** the sum is correct arithmetic but assumes the phases run one after another. The schedule line shows that phase 3 starts in month 5, while phase 2 runs until month 7.
- **Chair's routine question** (`s4.question.routine`, R-4 default): Thank you, Daniel. Before we move on, can you tell us why you chose this timeline?
- **R-4 alternative (v1, slide-specific):** Thank you, Daniel. Could you say a little more about how the eight months come together?
- **Participant, ask to explain:** Could you explain how the three phases fit into eight months?
- **Participant, challenge with the finding:** My checker flags this slide: "Numbers do not add up." What do you say to that?
- **CISO, reasoned answer** (`s4.answer.reason`, 45 words): Eight months is the right timeline for us. Phase three starts while phase two is still running, with a different team. Our own team begins the monitoring in month five, while our network partner is still installing the firewalls. That is the plan I recommend.
- **CISO, denial without reason** (`s4.answer.denial`, 14 words): Eight months is the right timeline for us. That is the plan I recommend.

Both middle sentences are strictly true for the schedule shown: phase 3 starts in month 5, and phase 2 runs from month 3 to month 7 (finding 5). v1's "That is why the bars overlap on the chart." is gone, because the chart no longer has bars (finding 52: two middle sentences).

### Slide 5: The decision today

**(a) Layout**

| Id | Kind | Text |
|---|---|---|
| `s5.title` | title | The decision today |
| `s5.num1` | number | $2.4 million for approval today |
| `s5.b1` | bullet | A firewall at all 13 sites and one watched internet connection |
| `s5.b2` | bullet | 24-hour monitoring of our network for two years |
| `s5.b3` | bullet | Secure remote access and training for all staff |
| `s5.b4` | bullet | Ready in 8 months, before our insurance renewal |

**(b) Narration** (CISO, `s5.narration`, 130 words, about 52 s)

> Let me bring this together. Every one of our thirteen sites gets a firewall, and we keep a single, watched connection to the internet. Specialists monitor our network around the clock for two years. Everyone who works away from the office connects securely, and all staff learn to spot fake emails. The whole plan costs two point four million dollars, and it is ready in eight months, before our insurance renewal. If we wait, we face a much higher insurance premium, or no cover at all, and the contract with our largest customer is at risk. I am confident this is the right plan for this company: the right size, at the right time. So I am asking the board to approve two point four million dollars today. Thank you.

**(c) Checker panel:** no marking. "No issues found" appears at 8,000 ms (C-3). No question or answer step.

### Who hears which question and answer

- **Checker group:** "Say nothing" or a timeout plays the chair's routine question. "Ask Daniel to explain" or "Challenge" shows the participant's question as text in a speech bubble for 3.5 s (OPEN D default), then the CISO answers.
- **Neutral group:** the chair's routine question plays whatever the participant answered, and also after a timeout.
- **Answer:** the reasoned answer on two slides, the denial on the participant's assigned denial slide (rotated over slides 2, 3 and 4, balanced within each group, server-side).

## 5. Rating screen

Shown after the opening (T0), after each slide, and after each answer (OPEN B default: 9 ratings). Window 15 s, with a visible countdown.

**Instruction:** How do you see Daniel Hale right now? Go with your first impression.

| Item | Wording | Scale (OPEN I default: 0 to 10) | METI basis |
|---|---|---|---|
| Trust | How much do you trust Daniel Hale right now? | 0 Not at all, 10 Completely | (overall rating, Abstract v8) |
| Competence | How competent is Daniel Hale? | 0 Not competent at all, 10 Very competent | expertise item: competent / incompetent |
| Benevolence | How much does Daniel Hale care about the company's interests? | 0 Not at all, 10 Very much | benevolence subscale description: acts "with the interests of others at heart" |
| Integrity | How honest is Daniel Hale? | 0 Not honest at all, 10 Very honest | integrity item: honest / dishonest |

OPEN I alternative: keep the three METI items as 7-point semantic differentials, for example "incompetent 1 2 3 4 5 6 7 competent". OPEN H: on timeout, record partial answers and mark the rest as timed out (default).

**METI check (C-5, R-7).** Checked against the full text of Hendriks, Kienhues and Bromme (2015), `Project Literature Review/references/Hendriks-2015-METI-PLOSONE.pdf`:

- The items are antonym pairs rated "on a 7-point scale like a semantic differential" (p. 4).
- "competent / incompetent" is an expertise item and "honest / dishonest" an integrity item (Table 1, p. 6; Table 3, p. 8).
- The benevolence items are "moral / immoral", "ethical / unethical", "responsible / irresponsible" and "considerate / inconsiderate" (Table 1, p. 6; Table 3, p. 8).
- All three adjectives used in v1 are in the final 14-item METI: 6 expertise, 4 integrity, 4 benevolence items (p. 11).
- The paper describes benevolence as whether the expert "acts with the interests of others at heart and not just personal aims or benefit" (p. 16).

So v1's "responsible" is a real METI item. It is replaced anyway (finding 54), because lay readers may read "responsible" as "in charge of". The new wording follows the paper's description of benevolence (p. 16) and Mayer's "wants to do good" for the trustor; it is not one of the four adjectives. Alternative: METI's "considerate" ("How considerate is Daniel Hale?"). The choice stays open (C-5, OPEN I).

The integrity item "How honest is Daniel Hale?" is the one place honesty is put in front of participants. It is the measure, not a flag, and is kept neutral (finding 61).

## 6. Action screen

Shown after the post-flag rating on slides 2, 3 and 4. OPEN C default: 15 s.

**Timeouts** (finding 7):

- **Neutral group:** a timeout is recorded as timed out, with no option. It is not recorded as "No", and not as any checker-group option. The chair asks the routine question in any case. For H4a, a neutral timeout counts as missing.
- **Checker group:** OPEN C default: the chair asks the routine question, as after "Say nothing". The record marks the step as timed out, with no option, so that a timeout can be told apart from a chosen "Say nothing" (C-15).

**Checker group**

> What do you do now?
> - Say nothing
> - Ask Daniel to explain
> - Challenge Daniel with the checker's finding
>
> If you say nothing, the chair asks a question.

Proposal: under the two speaking options, show in small type the exact question that will be shown (from section 4), so participants know what they are choosing.

**Neutral group**

> Do you have a question about this slide?
> - Yes
> - No
>
> The chair will now ask a question.

## 7. Budget decision screen

OPEN E default. Untimed (R-11, open).

> Daniel Hale asks the board to approve $2.4 million for the protection plan. Our insurance renewal and our largest customer's contract depend on this protection. How do you vote?
> - Approve the full $2.4 million
> - Approve a reduced amount (follow-up: "What amount would you approve?", $0 to $2.4 million in steps of $100,000)
> - Defer the decision to a later meeting
> - Reject the request
>
> How sure are you of your vote? 0 Not sure at all, 10 Completely sure

Alternatives: (1) Approve / Reject only; (2) Approve / Reject plus the 0 to 10 confidence; (3) one slider from $0 to $2.4 million in steps of $100,000.

## 8. Background questions

OPEN F default: all items after the budget decision, so that the cybersecurity and AI items do not prime the participant. Alternative: Oliver's four (field, country, age, job level) right after consent, the other three at the end. Untimed (R-11, open).

Intro line: "Finally, a few questions about you. They help us describe who took part. You can skip any question."

| # | Item | Options | Source |
|---|---|---|---|
| 1 | What is your main field of expertise? | Business and management; Finance or accounting; Law; Information technology; Cybersecurity; Engineering; Health care or medicine; Education; Natural sciences; Social sciences or humanities; Arts, media or design; Government or public service; Student, no field yet; Other (please specify); Prefer not to say | Oliver, Abstract v8 |
| 2 | What is your current job level? | Student; Employee; Leadership; C-Level; Prefer not to say | Oliver |
| 3 | What is the highest level of education you have completed? | Secondary school or less; College diploma or trade certificate; Bachelor's degree; Master's degree; Doctorate; Other; Prefer not to say | Abstract v8 |
| 4 | What is your age? | 18 to 24; 25 to 34; 35 to 44; 45 to 54; 55 to 64; 65 or older; Prefer not to say | Oliver, Abstract v8 |
| 5 | In which country do you live? | Country list (ISO 3166); Prefer not to say | Oliver |
| 6 | How would you rate your own knowledge of cybersecurity? | 1 None; 2 Basic; 3 Moderate; 4 Good; 5 Expert (alternative: 0 to 10) (R-8, open) | Abstract v8, planned moderator |
| 7 | How often do you use AI tools such as chatbots or writing assistants? | Never; Less than once a month; About once a month; About once a week; Several times a week; Every day (R-8, open) | Abstract v8, planned moderator |
| 8 | Hidden honeypot field, must stay empty | (none shown) | OPEN U |

"Prefer not to say" on every item and "You can skip any question" are my additions for ethics; remove if the ethics form says otherwise (C-10).

## 9. Flag review screen (OPEN G, checker group, optional)

Default: included, checker group only, before the debrief. Each row shows a small image of the slide with its marking.

> During the meeting, your checker flagged three slides. For each flag, do you think it was correct?
> - Slide 2: "Architecture likely wrong": Correct / Incorrect
> - Slide 3: "Firewalls likely too expensive, cheaper ones available": Correct / Incorrect
> - Slide 4: "Numbers do not add up": Correct / Incorrect

Alternative: add "Not sure" as a third option. The labels follow the flags in section 4 and change with them after the VLM run.

## 10. Debrief screen

Shown to both groups; the middle paragraph depends on the group.

> Thank you for taking part.
>
> Here is what this study was about. None of the five slides had a real issue. The plan, the costs and the timeline were all correct for this company.
>
> Some participants saw an AI slide checker during the meeting, and others did not. The study measures how flags from an AI checker affect how much people without expert knowledge trust the presenter. It also looks at how the presenter's answers change that trust.
>
> The company, the people and the meeting are fictional.

**Checker group:**

> You were in the group with the checker. [PLACEHOLDER: The flags you saw were drafted with a real AI tool and then fixed in advance, so every participant in your group saw the same flags.] Every flag was wrong for this company.

PLACEHOLDER: the words "drafted with a real AI tool" become true only after the VLM drafting run in section 4 (render the slides, run the model, store the provenance, choose and edit the flags). Until then the sentence must not be shown to participants. If the run is not done, the fallback sentence is: "The flags you saw were fixed in advance, so every participant in your group saw the same flags." The Abstract's Method sentence "flags were drafted with a real VLM from images of the slides" also becomes true only after that run.

The JSON (`content/study-config.v1.1.json`) carries the fallback paragraph until the run is done, so the placeholder sentence cannot reach a participant by mistake:

> You were in the group with the checker. The flags you saw were fixed in advance, so every participant in your group saw the same flags. Every flag was wrong for this company.

**Neutral group:**

> You were in the group without the checker. In the other group, a checker flagged slides 2, 3 and 4. The flags were fixed in advance, and every flag was wrong for this company.

**Both, at the end:**

> If you have questions about the study, please contact [CONTACT NAME, EMAIL].
>
> [ETHICS AND DATA WITHDRAWAL TEXT, see OPEN M]

## 11. Audio test step (OPEN U)

**Instruction on screen:** Please turn on your sound. Press Play and listen. Then choose the word you heard.

**Clip** (chair voice, `setup.audioTest`): This is a short sound check. If you can hear this clearly, the word to choose is lantern.

**Options** (in this order): river, lantern, pencil, garden. Correct: lantern.

**If wrong:** That was not the word in the clip. Please check your sound and try once more. (Proposal: 2 attempts, then the result is logged and the participant continues; see C-6.)

**Example scene for the playback check** (about 14 s, chair voice, `setup.example`): Welcome to the board meeting of Kestrel Bay Logistics. This is a short example of how the meeting will look and sound. The slides appear on the screen, and you will hear each speaker.

Example slide: "Board meeting, agenda" with four bullets: "1. Minutes of the last meeting", "2. Quarterly results", "3. New warehouse lease", "4. Protection of our network". Checker group: under OPEN A Option 1, the panel is shown idle during this example with the one-sentence caption from section 3. Under Option 2, the text screen follows the example or comes after T0 (R-1).

### Interface labels

Short labels used across screens. "Continue" is from section 3; the others were in the v1 JSON only, or come from the architecture or Oliver's prompt, and are listed here so that every participant-facing text in the JSON appears in this file (table 13.1, row 36).

| Label | Text | Source |
|---|---|---|
| Rating submit button | Next | v1 JSON |
| Countdown label | Time left | v1 JSON |
| Continue button | Continue | section 3 |
| Audio check button when the word check is off | I can hear it | architecture 4.1 |
| Consent accept button | I accept the terms and would like to participate in the study | Oliver's prompt |
| Consent checkbox label | [PLACEHOLDER: checkbox text from the ethics form, see OPEN M] | OPEN M |
| Honeypot field label (never shown to people) | Hidden honeypot field, must stay empty | section 8 |

## 12. Word and timing summary

Words counted by script; seconds at 150 words per minute. Real durations come from the TTS step (`durationMs`). Rows marked * changed in v1.1.

| Clip (audio key) | Speaker | Words | Seconds |
|---|---|---|---|
| `setup.audioTest` | Chair | 18 | 7 |
| `setup.example` | Chair | 34 | 14 |
| `opening.chair.1` | Chair | 23 | 9 |
| `opening.chair.2` | Chair | 66 | 26 |
| `opening.chair.3` | Chair | 60 | 24 |
| `opening.chair.4` | Chair | 5 | 2 |
| `opening.ciso.5` | CISO | 6 | 2 |
| Opening total | | 160 | 64 |
| `s1.narration` | CISO | 165 | 66 |
| `s2.narration` * | CISO | 141 | 56 |
| `s2.question.routine` * | Chair | 16 | 6 |
| `s2.answer.reason` * | CISO | 53 | 21 |
| `s2.answer.denial` * | CISO | 14 | 6 |
| `s3.narration` * | CISO | 141 | 56 |
| `s3.question.routine` * | Chair | 20 | 8 |
| `s3.answer.reason` * | CISO | 51 | 20 |
| `s3.answer.denial` * | CISO | 15 | 6 |
| `s4.narration` * | CISO | 138 | 55 |
| `s4.question.routine` * | Chair | 16 | 6 |
| `s4.answer.reason` * | CISO | 45 | 18 |
| `s4.answer.denial` * | CISO | 14 | 6 |
| `s5.narration` | CISO | 130 | 52 |

**Per participant**, spoken audio: opening 64 s, five narrations 285 s (v1: 302 s), up to three routine questions about 20 s, two reasons and one denial about 45 s (v1: 49 s). About 6.9 minutes of audio (v1: about 7). Add 10 s viewing after each of the five narrations (50 s), up to 9 ratings of 15 s (135 s) and up to 3 actions of 15 s (45 s): about 10.7 minutes before setup, the checker introduction, budget, background and debrief, and before any view time after an answer (C-4). The Abstract plans about 10 minutes in total (C-2).

A reasoned answer is now about 18 to 21 s and a denial about 6 s (C-7).

## 13. Further open points

### 13.1 Facts and defaults to confirm in one pass (R-5, R-6, R-8, R-11; findings 13 to 17)

Every row is my choice, not the Abstract's, unless the source column says otherwise. "Keep" is the default in this file.

| # | Item | Value in v1.1 | Where |
|---|---|---|---|
| 1 | Company name | Kestrel Bay Logistics (fictional) | 1 |
| 2 | Sector | Freight and warehousing | 1 |
| 3 | Location | Western Canada, head office and data centre in the Vancouver area | 1 |
| 4 | Sites | 13: head office plus 12 warehouses | 1, slides 2 to 5 |
| 5 | Size | 1,400 staff, about $480 million revenue a year | 1, slide 1, slide 3 |
| 6 | Budget amount | $2.4 million | 1, slides 1, 3, 5, budget screen |
| 7 | Currency | Canadian dollars, shown as "$" (C-12) | 1, slides |
| 8 | Stakes | Insurance renewal depends on it; largest customer's new contract requires it | 1, slides 1 and 5, budget screen |
| 9 | Deadline | Insurance renewal in 10 months; plan ready in 8 months | slides 1, 4, 5 |
| 10 | Cost lines | Firewalls $780,000; monitoring $960,000; remote access $310,000; installation and training $350,000 | slide 3 |
| 11 | Schedule | Phase 1 months 1 to 2; phase 2 months 3 to 7; phase 3 months 5 to 8; phase 2 by a network partner, phase 3 by the company's own team | slide 4, s4 reason |
| 12 | Cost of one stopped day | About $1.2 million in lost shipments | slide 4 |
| 13 | CISO name and gender | Daniel Hale, male ("himself" in the Abstract appendix) (C-9) | 1 |
| 14 | CISO experience | 6 years in the role, 12 years before at a national bank and a hospital group, two attempted break-ins stopped this year, main certifications | 1, 2 |
| 15 | Chair name and gender | Margaret Lin, female | 1, 2 |
| 16 | Profile delivery | Read aloud by the chair, with a profile card on screen (R-6) | 1, 2 |
| 17 | CISO greeting before T0 | "Thank you, Margaret. Good morning, everyone." (R-6) | 2 |
| 18 | Role instruction screen | Three lines, both groups, before the opening (C-1) | 1 |
| 19 | Routine question | Generic per slide, asks for a justification (R-4) | 4 |
| 20 | Closing line of every answer | "That is the plan I recommend." | 4 |
| 21 | Slide 4 schedule | Phase bars removed; month ranges in small type (C-16) | 4 |
| 22 | Benevolence item | "How much does Daniel Hale care about the company's interests?" (C-5) | 5 |
| 23 | Cybersecurity knowledge scale | 1 to 5 (None, Basic, Moderate, Good, Expert); alternative 0 to 10 (R-8) | 8 |
| 24 | AI use scale | Six frequency categories, Never to Every day (R-8) | 8 |
| 25 | Untimed screens | Role instruction, checker introduction text screen, budget decision, background questions, flag review, debrief (R-11) | 1, 3, 7, 8, 9, 10 |
| 26 | "No issues found" timing | Appears at 8 s on slides 1 and 5 (C-3) | 4 |
| 27 | Speech bubble duration | 3.5 s (OPEN D default) | 4 |
| 28 | View time after an answer | 0 ms (C-4) | 4 |
| 29 | Checker action timeout | Chair asks; recorded as timed out, not as "Say nothing" (C-15) | 6 |
| 30 | Budget stakes line | "Our insurance renewal and our largest customer's contract depend on this protection." | 7 |
| 31 | Reduced amount steps | $0 to $2.4 million in steps of $100,000 | 7 |
| 32 | "Prefer not to say" and skipping | On every background item (C-10) | 8 |
| 33 | Flag review options | Correct / Incorrect (OPEN G) | 9 |
| 34 | Audio check | Word "lantern", four options, 2 attempts (C-6) | 11 |
| 35 | Voices | Chair female, CISO male; locale en-US or en-CA (OPEN L, finding 66) | 1 |
| 36 | Interface labels | Next; Time left; Continue; I can hear it; consent labels as placeholders (OPEN M) | 11 |

Application parameters set in the architecture (consent minimum time, debrief minimum time, tolerances, rate limits, the "Started" versus "Ongoing" status) are listed in `02 - Architecture.md` and the review's F 2.13, not here.

### 13.2 Open points

C-1 to C-12 keep their v1 numbers. C-13 onward are new in v1.1.

1. **C-1. Role instruction screen.** The Abstract says the participant "sits on the board and decides on the CISO's budget request" but does not say how they are told. I added a three-line text screen for both groups before the opening scene; its third line is corrected in v1.1 (finding 58). Keep, change or drop?
2. **C-2. Total length.** The audio alone is about 6.9 minutes, and the run about 10.7 minutes before setup, the checker introduction, the budget vote, the background questions and the debrief, against the Abstract's "about 10 minutes". Options: shorten the narrations toward 110 words each (saves about 1.5 minutes), drop T1 and T5 (OPEN B), or update the Abstract's estimate.
3. **C-3. Panel timing on unflagged slides.** When does "No issues found" appear on slides 1 and 5, and what does the panel show on slides 2 to 4 before the flag appears? Proposal: on 1 and 5 the label appears at 8 s; before that, and before a flag, the panel is empty with its title only. A "Checking..." state is an alternative.
4. **C-4. View time after an answer.** Oliver's 10 s view applies after a narration. After the CISO's answer, does the rating come at once, or after a pause? Default: 0 ms (`timing.postAnswerViewMs`). The architecture's default is 10 s (FO-6).
5. **C-5. METI item wording.** The three v1 adjectives are checked against Hendriks et al. (2015) (section 5). Open: the benevolence wording ("care about the company's interests", or METI's "considerate"), and whether all three items take the METI adjective word for word.
6. **C-6. Audio check failure.** What happens after two wrong answers: continue and log, or end the session?
7. **C-7. Answer length differs by design.** A reasoned answer is about 18 to 21 s, a denial about 6 s. The difference is part of "only the presence of a reason differs", but it also means the denial slide reaches its rating sooner. Flag for the analysis or the limitations.
8. **C-8. Flag label length.** v1 used an invented 6-word limit to change Oliver's label. No source sets a limit, and v1.1 uses Oliver's wording. Should flag labels have a length limit?
9. **C-9. CISO gender.** The Abstract writes "himself", so Daniel is male. Changing this changes the voice choice (OPEN L) and the pronouns in the rating items' context.
10. **C-10. "Prefer not to say" and skipping.** Added to every background item; confirm with the ethics form (OPEN M).
11. **C-11. Debrief request not to share.** Many studies ask participants not to discuss the design with people who may still take part. Not in v1.1; add if wanted.
12. **C-12. Currency.** Amounts are Canadian dollars, written "$" on the slides. Participants outside Canada may read them as US dollars; for the trust measure this probably does not matter.
13. **C-13. Checker introduction: form, position and display** [OPEN A, R-1, FO-2]. Section 3 gives content for each option. Not decided.
14. **C-14. VLM drafting run** [R-10]. Which model; whether the prompt is one shared text or written per slide; what to do if the output for a slide has no plausible but wrong flag (run another model, change the slide, or use the fallback debrief sentence).
15. **C-15. Checker group timeout on the action step** [OPEN C, finding 7]. v1.1 records it as timed out with no option, while the chair asks as after "Say nothing". Confirm that a timeout is analysed apart from a chosen "Say nothing".
16. **C-16. Slide 4 flag** [finding 3]. Keep the arithmetic flag with the month ranges in small type (v1.1), or replace it with an arithmetic flag whose resolution is not on the slide at all.
17. **C-17. Slide 2 half-reason on screen** [finding 3]. The diagram and narration place the main data centre at head office, which is half of the slide 2 reason. Kept in v1.1; remove "Our main data centre sits there." and the data centre from the diagram if wanted.
18. **C-18. Routine question** [R-4, findings 12 and 50]. Generic per slide (v1.1 default) or slide-specific (v1 wording, kept in section 4).
19. **C-19. Voice locale** [OPEN L, finding 66]. en-CA voices to match a Vancouver company, if the chosen voice family has them, or en-US.

## 14. Change log v1 to v1.1

Applied:

- **Finding 1:** flag texts and the debrief sentence marked PLACEHOLDER; subsection "VLM drafting run (to do)" added; the debrief note names the words that become true only after the run and gives a fallback sentence.
- **Finding 2:** section 3 notes that the position of the checker introduction is open (R-1); the text screen works before or after T0.
- **Finding 3:** "You can see all three on the chart." removed; slide 4 phase bars replaced by a milestone line plus month ranges in small type; marking moved from 21,000 to 18,000 ms; remaining limitation stated.
- **Finding 4:** slide 3 sentence on the finance team deleted; slide 2's "This is a simple, proven design, and it fits the way we work." deleted rather than changed to "That is the design.", because a short closing line adds nothing and still sounds like a verdict.
- **Finding 5:** slide 4 reason now "Phase three starts while phase two is still running, with a different team.", followed by "Our own team begins the monitoring in month five, while our network partner is still installing the firewalls." Both are strictly true.
- **Finding 7:** section 6 records a neutral timeout as timed out with no option; the checker group timeout is also kept apart from "Say nothing" (C-15).
- **Finding 8:** Oliver's "Firewalls likely too expensive, cheaper ones available" restored on slide 3 (label, challenge question, flag review); word limit removed; label length is now C-8.
- **Findings 9 and 10:** section 3 gives the brief's default (one caption sentence in the example scene) and the text screen as options under OPEN A; the text does not say glasses or panel, with one optional line per display mode. The caption is shown, not spoken, so that no speaker in the meeting introduces the tool.
- **Findings 12 and 50:** generic justification questions from the chair, one per slide; v1's slide-specific questions kept as the R-4 alternative; closing line "That is the plan I recommend." on all six answers; first sentences of s3 and s4 changed so they also answer the chair's question; reason and denial keep the same first and last sentence.
- **Findings 13 to 17:** all scenario facts and content defaults listed in table 13.1.
- **Finding 51 and 56:** slide 2 flag detail now "connect each site straight to cloud services"; "route through" became "connect through".
- **Finding 52:** each reason now has exactly two middle sentences.
- **Finding 54:** benevolence item now "How much does Daniel Hale care about the company's interests?"; METI full text found and cited with pages; "considerate" kept as the alternative.
- **Finding 58:** role instruction third line corrected.
- **Finding 59:** stakes line added to the budget screen.
- **Finding 60:** background intro now "They help us describe who took part."
- **Finding 61:** integrity wording kept neutral; noted in section 5.
- **Finding 62:** timing recomputed (section 12); decision stays with Oliver (C-2).
- **Finding 64:** the VLM run lists building the slide images as its first step.
- **Finding 66:** voice locale added as C-19.
- **JSON pass (added after the first v1.1 report):** debrief fallback paragraph written out in section 10, and the interface labels listed in section 11, so that every participant-facing text in `content/study-config.v1.1.json` appears in this file word for word.

No change needed: findings 53, 55, 57, 63, 65 (checks passed or "keep").

Declined: none. Partly applied: finding 3 for slide 2 (the data centre at head office stays on the slide, C-17), because it is a plain fact of the layout and removing it would make the diagram less clear.
