# Study content v1: slides, flags, answers and scripts

Draft written 2026-10-23 from `CMPT 863 Project Abstract v8.md` (the authority) and `01 - Study Flow Brief (Abstract v8 mapping).md`. The machine-readable version is `content/study-config.v1.json`; every text below appears there word for word. Nothing marked OPEN is decided: the brief's proposed default is used and marked. Points I found while writing are listed at the end.

## 1. Scenario

**Company.** Kestrel Bay Logistics (fictional). A freight and warehousing company in western Canada: head office and main data centre in the Vancouver area, 12 warehouses, 1,400 staff, about $480 million revenue a year. Everything runs on one network: orders, warehouse scanners, truck routes, payroll.

**Meeting.** The quarterly board meeting, agenda item 4: "the protection of our network". The participant sits on the board.

**Budget request.** $2.4 million (CAD), identical for both groups. What depends on it: the company's cyber insurance is renewed in 10 months only if stronger protection is in place, and the company's largest customer has written the same requirement into its new contract. If the board waits, the premium rises or cover ends, and the customer contract is at risk.

Where the amount appears: on **slide 1** (the ask, so that every rating and every flag is judged against a known amount), in the title of **slide 3** (the breakdown), and on **slide 5** (the decision, right before the budget vote). The chair also says it once in the opening. Showing it on both slide 1 and slide 5 is what a real board deck does: state the ask up front, restate it at the decision.

**CISO.** Daniel Hale, Chief Information Security Officer (fictional). Pronoun "he", as in the Abstract's appendix ("our CISO answers for himself"). Profile, 4 sentences after the introduction:

> Daniel has led our security for six years. Before he joined us, he spent twelve years protecting the networks of a national bank and a large hospital group. This year he led our response to two attempted break-ins, and both were stopped before any harm was done. He holds the main professional certifications in his field.

**Proposal for the profile:** the chair reads it aloud in the opening scene, and a profile card with the same four points is shown on screen while she speaks. Spoken keeps it in the audio-led flow and ensures every participant receives it before the baseline rating; the card helps participants who listen less well.

**Board chair.** Margaret Lin, chair of the board (fictional). Her voice must be clearly different from the CISO's (OPEN L).

The CISO never mentions AI, checkers or tools in any line.

### Role instruction (both groups, shown before the opening scene)

> You are a member of the board of Kestrel Bay Logistics.
> Today the company's head of security asks the board to approve a budget. Listen to the presentation.
> After some slides you will rate the presenter, and at the end you will vote on the budget.

This screen is my proposal (further open point 1).

## 2. Opening scene script

No markings. Checker group: the panel is visible but idle, with no label (OPEN A). A profile card is on screen during line 2. Total 160 words, about 64 seconds.

| # | Speaker | Line | Audio key |
|---|---|---|---|
| 1 | Chair | Good morning, everyone, and thank you for being here. Let us move to item four on our agenda: the protection of our network. | `opening.chair.1` |
| 2 | Chair | Presenting today is Daniel Hale, our Chief Information Security Officer. Daniel has led our security for six years. Before he joined us, he spent twelve years protecting the networks of a national bank and a large hospital group. This year he led our response to two attempted break-ins, and both were stopped before any harm was done. He holds the main professional certifications in his field. | `opening.chair.2` |
| 3 | Chair | Daniel is asking the board for two point four million dollars to strengthen the protection of our warehouses and our head office. Our insurer and our largest customer both expect this protection to be in place within the year. We will hear his five slides and ask our questions along the way. The board will decide on the budget today. | `opening.chair.3` |
| 4 | Chair | Daniel, the floor is yours. | `opening.chair.4` |
| 5 | CISO | Thank you, Margaret. Good morning, everyone. | `opening.ciso.5` |

Line 5 lets participants hear the CISO's voice once before the baseline rating (T0). Slide 1 then starts without a second greeting.

**Profile card (on screen during line 2):**

- Daniel Hale, Chief Information Security Officer, Kestrel Bay Logistics
- Head of security at Kestrel Bay Logistics for 6 years
- 12 years before that protecting the networks of a national bank and a large hospital group
- Led the response to two attempted break-ins this year, both stopped before any harm
- Holds the main professional certifications in his field

## 3. Checker introduction (checker group only, OPEN A)

Default per the brief: shown after the example scene of the setup step, as a short text screen. Button: "Continue".

> In this meeting you wear smart glasses with a slide checker. The checker is an AI tool.
> While the presenter speaks, it looks at each slide through the camera in your glasses.
> If it sees a possible issue, it marks that part of the slide and shows a short note in the checker panel.
> If it sees none, the panel shows: No issues found.

The text says nothing about how accurate the checker is, in line with "no accuracy or confidence figure".

**No-issue label:** "No issues found" (Abstract v8). Configurable as `labels.noIssueLabel`. The v7-to-v8 change file records that Oliver has not decided between this wording and a supporting label, so this is a default, not a decision.

**Panel title:** "Slide checker" (`labels.panelTitle`).

## 4. The five slides

Order: slide 1 context and the ask, slide 2 architecture, slide 3 costs, slide 4 timeline and risk, slide 5 summary and decision. This is the order of a typical board approval deck (why, what, how much, when, decide), and it puts the three flags on three different kinds of content: a design point, a price point and an arithmetic point.

All five slides are correct. The numbers add up: $780,000 + $960,000 + $310,000 + $350,000 = $2,400,000. Phases run months 1 to 2, 3 to 7 and 5 to 8, so the work is done in 8 months although the durations sum to 11.

The three flags follow the literature review's point that a flag can cite a sound rule and still misread the slide. Each one is a rule of thumb, a price comparison or a sum that a vision-language model would plausibly produce from the slide image, and each is wrong for this company for a reason the CISO states in two or three short sentences. None is about honesty. None carries a confidence or accuracy figure.

**Answer design.** The CISO's answer must fit three different questions: the chair's routine question, "explain" and "challenge". So no answer starts with "No"; each starts with a claim that fits all three. The reasoned answer and the denial share the same first sentence and the same last sentence. Only the middle, the reason, differs:

- Reasoned: [claim]. [reason, 2 to 3 sentences]. We have looked at this carefully, and it stays as planned.
- Denial: [claim]. We have looked at this carefully, and it stays as planned.

The narration never states the reason that the answer gives, so that the denial truly carries no information.

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

**(c) Checker panel:** no marking. "No issues found" appears at 8,000 ms (my proposal, further open point 3). No question or answer step.

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

**(b) Narration** (CISO, `s2.narration`, 154 words, about 62 s)

> Let me show you how the network will look. Picture our head office in the middle of this diagram. Our main data centre sits there. All twelve warehouses connect to that data centre. That is the line you see from each warehouse to the centre. At every site, we put a firewall. Think of a firewall as a locked door with a guard. It decides what may come in and go out, and it keeps the warehouse machines apart from the office computers. So if one computer is infected, the problem stays in that room. Then we have one connection to the internet, at head office, and our security team watches it day and night. Finally, people who work from home or on the road connect through secure remote access. They get in only after we have checked who they are. This is a simple, proven design, and it fits the way we work.

**(c) Marking and answers**

- **Target:** `s2.b2` and `s2.img1` (the lines from the warehouses to head office).
- **Appears at:** 10,000 ms (anchor: "All twelve warehouses connect").
- **Colour token:** `flag`.
- **Flag label:** Architecture likely wrong
- **Flag detail:** All 12 warehouses route through head office, while common practice is to connect each site directly to the cloud.
- **Why it is wrong for this company:** the textbook rule (send branch traffic straight to the cloud) assumes the systems live in the cloud. Kestrel Bay's warehouse systems run in its own data centre at head office.
- **Chair's routine question** (`s2.question.routine`): Thank you, Daniel. Before we move on, could you say a little more about how the warehouses connect?
- **Participant, ask to explain:** Could you explain why all the warehouses connect through head office?
- **Participant, challenge with the finding:** My checker flags this slide: "Architecture likely wrong." What do you say to that?
- **CISO, reasoned answer** (`s2.answer.reason`, 55 words): This design is the right one for us. Our warehouse systems run in our own data centre at head office, not in the cloud, so every scan goes there anyway. A direct link at each warehouse would only add twelve more doors to guard. We have looked at this carefully, and it stays as planned.
- **CISO, denial without reason** (`s2.answer.denial`, 19 words): This design is the right one for us. We have looked at this carefully, and it stays as planned.

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

**(b) Narration** (CISO, `s3.narration`, 160 words, about 64 s)

> Now to the money, and where the two point four million dollars goes. The first line is the firewalls for all thirteen sites, with three years of support: seven hundred and eighty thousand dollars. The largest line is round-the-clock monitoring for two years: nine hundred and sixty thousand dollars. For that, a team of specialists watches our network every hour of every day, and they can close a door within minutes if something looks wrong. We could not build that team ourselves for this price. Next is secure remote access for all fourteen hundred staff: three hundred and ten thousand dollars. And finally, installation and staff training: three hundred and fifty thousand dollars. Most break-ins start with a fake email, so training our people is some of the best value on this page. Together, that is two point four million dollars. Our finance team has checked every line, and each one is backed by a written quote from a supplier.

**(c) Marking and answers**

- **Target:** `s3.num1`.
- **Appears at:** 11,000 ms (anchor: "seven hundred and eighty thousand dollars").
- **Colour token:** `flag`.
- **Flag label:** Firewalls likely overpriced, cheaper ones available
- **Flag detail:** Firewalls for small offices cost about $5,000 each, so 13 units should cost about $65,000, not $780,000.
- **Note on the label:** Oliver's wording "Firewalls likely too expensive, cheaper ones available" has 7 words, one over the 6-word limit, so v1 uses "overpriced". His wording is the alternative; the limit or the wording can change.
- **Why it is wrong for this company:** the check compares with office firewalls. A warehouse with hundreds of scanners needs a larger model, and the price includes support and next-day replacement.
- **Chair's routine question** (`s3.question.routine`): Thank you, Daniel. The firewalls are a large line. Could you say a little more about that cost?
- **Participant, ask to explain:** Could you explain why the firewalls cost $780,000?
- **Participant, challenge with the finding:** My checker flags this slide: "Firewalls likely overpriced, cheaper ones available." What do you say to that?
- **CISO, reasoned answer** (`s3.answer.reason`, 52 words): These firewalls are the right ones for us. An office model would not cope with our warehouses, where hundreds of scanners run day and night. The price also covers next-day replacement at every site, so a warehouse is never left unprotected. We have looked at this carefully, and it stays as planned.
- **CISO, denial without reason** (`s3.answer.denial`, 19 words): These firewalls are the right ones for us. We have looked at this carefully, and it stays as planned.

### Slide 4: Ready in 8 months, before the insurance renewal

**(a) Layout**

| Id | Kind | Text |
|---|---|---|
| `s4.title` | title | Ready in 8 months, before the insurance renewal |
| `s4.chart1` | chart | Timeline chart, months 1 to 10 on the horizontal axis: Phase 1 bar over months 1 to 2; Phase 2 bar over months 3 to 7; Phase 3 bar over months 5 to 8; a dashed line at month 8 labelled Ready; a dashed line at month 10 labelled Insurance renewal |
| `s4.b1` | bullet | Phase 1, design and ordering: 2 months |
| `s4.b2` | bullet | Phase 2, firewalls installed at all 13 sites: 5 months |
| `s4.b3` | bullet | Phase 3, monitoring set up and staff trained: 4 months |
| `s4.num1` | number | Ready in 8 months, renewal in 10 months |
| `s4.num2` | number | One day of stopped warehouses: about $1.2 million in lost shipments |

**(b) Narration** (CISO, `s4.narration`, 146 words, about 58 s)

> So when will this be ready? The work comes in three phases. Phase one, design and ordering, takes two months. Phase two, installing the firewalls at all thirteen sites, takes five months. Phase three, setting up the monitoring and training our staff, takes four months. You can see all three on the chart. Altogether, we are ready in eight months. Our insurance renewal is in ten months, so that gives us two months of margin. Why does the timing matter? Because the risk is real, and it is expensive. If our warehouses stop for a single day, we lose about one point two million dollars in shipments. That is half of this request, gone in one day. And an attack rarely stops things for only one day. If we start today, we are protected well before the renewal. If we wait another quarter, we miss it.

**(c) Marking and answers**

- **Target:** `s4.b1`, `s4.b2`, `s4.b3` and `s4.num1`.
- **Appears at:** 21,000 ms (anchor: "Altogether, we are ready in eight months").
- **Colour token:** `flag`.
- **Flag label:** Numbers do not add up
- **Flag detail:** The three phases add up to 11 months, but the slide says the work is done in 8 months.
- **Why it is wrong for this company:** the sum is correct arithmetic but assumes the phases run one after another. The chart shows phases 2 and 3 overlapping.
- **Chair's routine question** (`s4.question.routine`): Thank you, Daniel. Could you say a little more about how the eight months come together?
- **Participant, ask to explain:** Could you explain how the three phases fit into eight months?
- **Participant, challenge with the finding:** My checker flags this slide: "Numbers do not add up." What do you say to that?
- **CISO, reasoned answer** (`s4.answer.reason`, 51 words): The eight months are correct. Phases two and three run at the same time, with different teams. Our network partner installs the firewalls while our own team sets up the monitoring. That is why the bars overlap on the chart. We have looked at this carefully, and it stays as planned.
- **CISO, denial without reason** (`s4.answer.denial`, 16 words): The eight months are correct. We have looked at this carefully, and it stays as planned.

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

**(c) Checker panel:** no marking. "No issues found" appears at 8,000 ms (further open point 3). No question or answer step.

### Who hears which question and answer

- **Checker group:** "Say nothing" or a timeout plays the chair's routine question. "Ask Daniel to explain" or "Challenge" shows the participant's question as text in a speech bubble for 3.5 s (OPEN D default), then the CISO answers.
- **Neutral group:** the chair's routine question plays whatever the participant answered.
- **Answer:** the reasoned answer on two slides, the denial on the participant's assigned denial slide (rotated over slides 2, 3 and 4, balanced within each group, server-side).

## 5. Rating screen

Shown after the opening (T0), after each slide, and after each answer (OPEN B default: 9 ratings). Window 15 s, with a visible countdown.

**Instruction:** How do you see Daniel Hale right now? Go with your first impression.

| Item | Wording | Scale (OPEN I default: 0 to 10) | METI source item |
|---|---|---|---|
| Trust | How much do you trust Daniel Hale right now? | 0 Not at all, 10 Completely | (overall rating, Abstract v8) |
| Competence | How competent is Daniel Hale? | 0 Not competent at all, 10 Very competent | expertise: incompetent / competent |
| Benevolence | How responsible is Daniel Hale? | 0 Not responsible at all, 10 Very responsible | benevolence: irresponsible / responsible |
| Integrity | How honest is Daniel Hale? | 0 Not honest at all, 10 Very honest | integrity: dishonest / honest |

OPEN I alternative: keep the three METI items as 7-point semantic differentials, for example "incompetent 1 2 3 4 5 6 7 competent". OPEN H: on timeout, record partial answers and mark the rest as timed out (default).

The METI adjectives above are from my reading of Hendriks et al. (2015) and should be checked against the paper's item table before launch (further open point 5). Other METI benevolence adjectives are "moral", "ethical" and "considerate"; "responsible" is the one closest to Mayer's "wants to do good for the company" that a lay reader understands at a glance.

## 6. Action screen

Shown after the post-flag rating on slides 2, 3 and 4. OPEN C default: 15 s; on timeout, record "none" and treat it as "Say nothing".

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

OPEN E default. Untimed.

> Daniel Hale asks the board to approve $2.4 million for the protection plan. How do you vote?
> - Approve the full $2.4 million
> - Approve a reduced amount (follow-up: "What amount would you approve?", $0 to $2.4 million in steps of $100,000)
> - Defer the decision to a later meeting
> - Reject the request
>
> How sure are you of your vote? 0 Not sure at all, 10 Completely sure

Alternatives: (1) Approve / Reject only; (2) Approve / Reject plus the 0 to 10 confidence; (3) one slider from $0 to $2.4 million in steps of $100,000.

## 8. Background questions

OPEN F default: all items after the budget decision, so that the cybersecurity and AI items do not prime the participant. Alternative: Oliver's four (field, country, age, job level) right after consent, the other three at the end.

Intro line: "Finally, a few questions about you. They help us compare groups. You can skip any question."

| # | Item | Options | Source |
|---|---|---|---|
| 1 | What is your main field of expertise? | Business and management; Finance or accounting; Law; Information technology; Cybersecurity; Engineering; Health care or medicine; Education; Natural sciences; Social sciences or humanities; Arts, media or design; Government or public service; Student, no field yet; Other (please specify); Prefer not to say | Oliver, Abstract v8 |
| 2 | What is your current job level? | Student; Employee; Leadership; C-Level; Prefer not to say | Oliver |
| 3 | What is the highest level of education you have completed? | Secondary school or less; College diploma or trade certificate; Bachelor's degree; Master's degree; Doctorate; Other; Prefer not to say | Abstract v8 |
| 4 | What is your age? | 18 to 24; 25 to 34; 35 to 44; 45 to 54; 55 to 64; 65 or older; Prefer not to say | Oliver, Abstract v8 |
| 5 | In which country do you live? | Country list (ISO 3166); Prefer not to say | Oliver |
| 6 | How would you rate your own knowledge of cybersecurity? | 1 None; 2 Basic; 3 Moderate; 4 Good; 5 Expert (alternative: 0 to 10) | Abstract v8, planned moderator |
| 7 | How often do you use AI tools such as chatbots or writing assistants? | Never; Less than once a month; About once a month; About once a week; Several times a week; Every day | Abstract v8, planned moderator |
| 8 | Hidden honeypot field, must stay empty | (none shown) | OPEN U |

"Prefer not to say" on every item and "You can skip any question" are my additions for ethics; remove if the ethics form says otherwise.

## 9. Flag review screen (OPEN G, checker group, optional)

Default: included, checker group only, before the debrief. Each row shows a small image of the slide with its marking.

> During the meeting, your checker flagged three slides. For each flag, do you think it was correct?
> - Slide 2: "Architecture likely wrong": Correct / Incorrect
> - Slide 3: "Firewalls likely overpriced, cheaper ones available": Correct / Incorrect
> - Slide 4: "Numbers do not add up": Correct / Incorrect

Alternative: add "Not sure" as a third option.

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

> You were in the group with the checker. The flags you saw were drafted with a real AI tool and then fixed in advance, so every participant in your group saw the same flags. Every flag was wrong for this company.

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

**If wrong:** That was not the word in the clip. Please check your sound and try once more. (Proposal: 2 attempts, then the result is logged and the participant continues; see further open point 6.)

**Example scene for the playback check** (OPEN A default, about 14 s, chair voice, `setup.example`): Welcome to the board meeting of Kestrel Bay Logistics. This is a short example of how the meeting will look and sound. The slides appear on the screen, and you will hear each speaker.

Example slide: "Board meeting, agenda" with four bullets: "1. Minutes of the last meeting", "2. Quarterly results", "3. New warehouse lease", "4. Protection of our network". Checker group: the panel is shown idle during this example, then the checker introduction text follows.

## 12. Word and timing summary

Words counted by script; seconds at 150 words per minute. Real durations come from the TTS step (`durationMs`).

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
| `s2.narration` | CISO | 154 | 62 |
| `s2.question.routine` | Chair | 18 | 7 |
| `s2.answer.reason` | CISO | 55 | 22 |
| `s2.answer.denial` | CISO | 19 | 8 |
| `s3.narration` | CISO | 160 | 64 |
| `s3.question.routine` | Chair | 18 | 7 |
| `s3.answer.reason` | CISO | 52 | 21 |
| `s3.answer.denial` | CISO | 19 | 8 |
| `s4.narration` | CISO | 146 | 58 |
| `s4.question.routine` | Chair | 16 | 6 |
| `s4.answer.reason` | CISO | 51 | 20 |
| `s4.answer.denial` | CISO | 16 | 6 |
| `s5.narration` | CISO | 130 | 52 |

**Per participant**, spoken audio: opening 64 s, five narrations 302 s, three questions about 20 s, two reasons and one denial about 49 s. About 7 minutes of audio. Add 10 s viewing after each of the five narrations (50 s), up to 9 ratings of 15 s (135 s) and up to 3 actions of 15 s (45 s): about 11 minutes before setup, budget, background and debrief. The Abstract plans about 10 minutes in total (further open point 2).

## 13. Further open points found while writing

1. **Role instruction screen.** The Abstract says the participant "sits on the board and decides on the CISO's budget request" but does not say how they are told. I added a three-line text screen for both groups before the opening scene. Keep, change or drop?
2. **Total length.** The audio alone is about 7 minutes, and the full run about 13 to 15 minutes with setup and questions, against the Abstract's "about 10 minutes". Options: shorten the narrations toward 110 words each, drop T1 and T5 (OPEN B), or update the Abstract's estimate.
3. **Panel timing on unflagged slides.** When does "No issues found" appear on slides 1 and 5, and what does the panel show on slides 2 to 4 before the flag appears? Proposal: on 1 and 5 the label appears at 8 s; before that, and before a flag, the panel is empty with its title only. A "Checking..." state is an alternative.
4. **View time after an answer.** Oliver's 10 s view applies after a narration. After the CISO's answer, does the rating come at once, or after a pause? Default in the JSON: 0 ms (`timing.postAnswerViewMs`).
5. **METI item wording.** The adjectives chosen (competent, responsible, honest) should be checked against the item table of Hendriks et al. (2015) before launch.
6. **Audio check failure.** What happens after two wrong answers: continue and log, or end the session?
7. **Answer length differs by design.** A reasoned answer is about 21 s, a denial about 7 s. The difference is part of "only the presence of a reason differs", but it also means the denial slide reaches its rating sooner. Flag for the analysis or the limitations.
8. **Firewall label length.** Oliver's example has 7 words; the brief's limit is 6. v1 uses "Firewalls likely overpriced, cheaper ones available".
9. **CISO gender.** The Abstract writes "himself", so Daniel is male. Changing this changes the voice choice (OPEN L) and the pronouns in the rating items' context.
10. **"Prefer not to say" and skipping.** Added to every background item; confirm with the ethics form (OPEN M).
11. **Debrief request not to share.** Many studies ask participants not to discuss the design with people who may still take part. Not in v1; add if wanted.
12. **Currency.** Amounts are Canadian dollars, written "$" on the slides. Participants outside Canada may read them as US dollars; for the trust measure this probably does not matter.
