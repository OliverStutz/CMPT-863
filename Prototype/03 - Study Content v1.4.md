# Study content v1.4: slides, flags, answers and scripts

Revision written 2026-10-24 from `03 - Study Content v1.3.md`. It puts in the three flags selected from the VLM drafting run (`content/vlm-drafts/v1.3/selection.md`) and adjusts the answers to them. Earlier versions are left unchanged. The machine-readable version is `content/study-config.v1.4.json`. Oliver's answers in `04 - Open Decisions.md` are applied as in v1.3. Abstract v8 remains the authority on the study design. Where an answer left me a choice, I made it and listed it in section 14 (choices made in v1.3; still valid). Section numbers in square brackets, such as [1.13], refer to `04 - Open Decisions.md`.

**Participant flow (decided):** consent, device check, setup (audio check, example scene, "The meeting is about to start"), the opening scene, baseline rating T0, then CISO slides 1 to 5. Every slide is followed by a rating. A flagged slide also gets the action, the question, the CISO's answer and a second rating. Then the budget decision, the background questions and the debrief. There is no role instruction screen [4.3] and no flag review screen [1.15].

The number of slides and which slides carry a flag are configurable [1.24]. Abstract v8's five slides with flags on slides 2 to 4 remain the default content below.

## 1. Scenario

**Company.** Meridian Logistics (fictional). A freight and warehousing company: a head office with the main data centre, 12 warehouses, 1,400 staff, about $480 million revenue a year. Everything runs on one network: orders, warehouse scanners, truck routes, payroll. No country, city or region is named anywhere [1.11]. Amounts are written with a plain "$" [4.4].

**Meeting.** The quarterly board meeting, agenda item 4: "the protection of our network". The participant sits on the board. Without a role screen, the participant learns this from the scene itself: the chair addresses the board, and says the board will decide on the budget today.

**Budget request.** $2.4 million, identical for both groups. What depends on it: the company's cyber insurance is renewed in 10 months only if stronger protection is in place, and the company's largest customer has written the same requirement into its new contract. If the board waits, the premium rises or cover ends, and the customer contract is at risk. The amount is on slide 1 (the ask) and slide 5 (the decision), and in the title of slide 4 (the costs). The chair also says it once in the opening.

**Scene (decided) [2.10].** Every scene is a still composite with light animation: a boardroom image, a CISO portrait that lights up while he speaks, the slide rendered live, and, in the checker group, the checker's highlight on the slide. There is no video. There are no captions; the spoken text is available to screen readers only [2.7].

**CISO.** Daniel Hale, Chief Information Security Officer (fictional), "he". Six years in the role; twelve years before that at a national bank and a large hospital group; led the response to two attempted break-ins this year; holds the main certifications in his field. The CISO never mentions AI, checkers or tools in any line.

**Board chair.** Margaret Lin, chair of the board (fictional), "she".

**Voices [3.6].** US English. Three roles: CISO (male), chair (clearly different), and a board member who speaks the participant's question [1.4]. Two or three Google voices per role are shortlisted from a family with word timing marks; Oliver picks by listening.

## 2. Setup screens (both groups)

### 2.1 Audio check [2.3, 2.4]

**Screen text:**

> Please turn on your sound.
> Press Play and listen. You can play the clip as often as you like.
> Then choose the word you heard. You can answer once.

**Clip** (chair voice, `setup.audioTest`, replayable, the only replayable clip): This is a short sound check. If you can hear this clearly, the word to choose is lantern.

**Options** (in this order): river, lantern, pencil, garden. Correct: lantern.

One answer. Right or wrong, the answer is recorded and shown on the session record in the admin view. The participant always continues; no feedback is shown. A wrong answer flags the session and never blocks it.

### 2.2 Example scene (playback check)

About 13 s, chair voice, `setup.example`, played once: Welcome to the board meeting of Meridian Logistics. This is a short example of how the meeting will look and sound. The slides appear on the screen, and you will hear each speaker.

Example slide: "Board meeting, agenda" with four bullets: "1. Minutes of the last meeting", "2. Quarterly results", "3. New warehouse lease", "4. Protection of our network". Both groups see the same example scene. The checker does not appear in it.

### 2.3 "The meeting is about to start" (last setup screen)

**Both groups:**

> The meeting is about to start.
> You will hear each speaker once. The audio cannot be replayed.

**Checker group only**, one more line after the first one [1.1, 1.25]:

> During this meeting, a checker reads each slide and marks anything it finds.

The neutral group sees the same screen without that line. Button: "Start". The sentence does not say AI, glasses, camera or panel, and says nothing about whether the checker's findings are right or wrong. Participants learn that the checker was an AI at the debrief [1.25]. The sentence comes before the baseline rating T0 for the checker group; the Validity check that the groups start alike at T0 still applies.

## 3. Opening scene

No markings. The checker shows nothing during the opening. A profile card is on screen during line 2 [1.12]. Total 160 words, about 64 s. A 10 s view follows before T0 [1.8].

| # | Speaker | Line | Audio key |
|---|---|---|---|
| 1 | Chair | Good morning, everyone, and thank you for being here. Let us move to item four on our agenda: the protection of our network. | `opening.chair.1` |
| 2 | Chair | Presenting today is Daniel Hale, our Chief Information Security Officer. Daniel has led our security for six years. Before he joined us, he spent twelve years protecting the networks of a national bank and a large hospital group. This year he led our response to two attempted break-ins, and both were stopped before any harm was done. He holds the main professional certifications in his field. | `opening.chair.2` |
| 3 | Chair | Daniel is asking the board for two point four million dollars to strengthen the protection of our warehouses and our head office. Our insurer and our largest customer both expect this protection to be in place within the year. We will hear his five slides and ask our questions along the way. The board will decide on the budget today. | `opening.chair.3` |
| 4 | Chair | Daniel, the floor is yours. | `opening.chair.4` |
| 5 | CISO | Thank you, Margaret. Good morning, everyone. | `opening.ciso.5` |

Line 5 is the CISO's one greeting before T0 [1.12]. Slide 1 then starts without a second greeting.

**Profile card (on screen during line 2):**

- Daniel Hale, Chief Information Security Officer, Meridian Logistics
- Head of security at Meridian Logistics for 6 years
- 12 years before that protecting the networks of a national bank and a large hospital group
- Led the response to two attempted break-ins this year, both stopped before any harm
- Holds the main professional certifications in his field

## 4. The checker (checker group only)

The checker shows as highlighting on the slide while the CISO talks [1.16]. A flag is a coloured box around the marked part of the slide, with a short note next to it: the heading "Slide checker", the flag label, and one sentence of detail. On unflagged slides, a note with "No issues found" appears in the top right corner of the slide [4.2]. The checker is present from slide 1 [1.18].

The neutral group sees the identical simulation without any highlighting or note [1.16].

**Timing (open, brief defaults) [1.18]:** "No issues found" appears 8 s into slides 1 and 5. Before a flag or label appears, nothing is shown.

Flags are observations about the slide, never about the CISO's honesty, with no accuracy or confidence figure.

## 5. The five slides

The deck tells the CISO's case for funding in stages [4.1]: slide 1 what is going on, slide 2 the risk and its impact, slide 3 how we solve it, slide 4 what it costs, slide 5 the decision. Slides 1 and 5 are unflagged. Slides 2 to 4 each carry one flag: on the risk and impact figures, on the architecture, and on the costs.

All five slides are correct for this company:

- Slide 2: $1.2 million a day for 5 days is $6 million in lost shipments; with about $2 million for recovery, one serious attack costs about $8 million.
- Slide 4: $780,000 + $960,000 + $310,000 + $350,000 = $2,400,000.
- Slide 5: ready in 8 months, insurance renewal in 10 months.

Each narration is about 110 words, about 45 s [1.7]. The narration never states the reason that the CISO's answer gives, so that the denial truly carries no information.

### 5.1 Flags from the VLM drafting run [1.17]

The run is done. Folder: `Prototype/content/vlm-drafts/v1.3/`. A Claude Opus agent, run fresh with no other project file, read each slide image with one fixed prompt (`prompt.md`), on 2026-10-24 (`run.json`). The raw output per slide is kept unedited (`s1.raw.md` to `s5.raw.md`). Oliver delegated the choice of flags to the coordinator; `selection.md` records, per flagged slide, the finding chosen, the model's text, every edit, why the flag is wrong for Meridian, and the alternatives considered.

| Slide | Finding | Final label (verbatim from the model) | Edit to the detail |
|---|---|---|---|
| 2 | 1 | $8 million does not follow from shown figures | second sentence deleted |
| 3 | 1 | Head office is a single point of failure | second sentence deleted |
| 4 | 3 | Firewall cost per site looks high | hedge and second sentence deleted |

The model also reported issues on slides 1 and 5. By design the checker shows "No issues found" there (Abstract v8); the model's output for those slides is on record in `selection.md` and not shown.

### 5.2 Answer design [1.10, 1.22]

Each flagged slide has three questions, and all three ask for a justification: the chair's slide-specific "why" question, the board member's "explain" question, and the board member's challenge. The reasoned answer answers all three; the denial fails all three alike.

The reasoned answer and the denial share the same first sentence and the same last sentence. Only the middle differs. The last sentence is neutral: "That is the plan I recommend." on slides 3 and 4, and "That is the risk we face." on slide 2, where no plan is shown.

**Denial length [1.22].** Each denial is first written at the same length as its reasoned answer, with a middle that carries no reason: only confidence and process words, no information about the slide. A short denial plus silence is given as the fallback for each slide. My judgement per slide is at the end of each slide below.

Remaining risk: process words such as "I have gone over every line" may be read as a weak generic reason. They give no information that answers the flag, which is what Abstract v8's "denial with no reason" requires, but a reviewer may see it differently.

### Slide 1: Where we stand today

**(a) Layout**

| Id | Kind | Text |
|---|---|---|
| `s1.title` | title | Where we stand today |
| `s1.num1` | number | $2.4 million: budget request for approval today |
| `s1.b1` | bullet | Our 12 warehouses and head office run on one network: orders, scanners, trucks and payroll |
| `s1.b2` | bullet | This year we stopped two attempted break-ins into that network |
| `s1.b3` | bullet | Attackers are getting better |
| `s1.b4` | bullet | Our insurer and our largest customer now expect stronger protection |

**(b) Narration** (CISO, `s1.narration`, 115 words, about 46 s)

> I will keep this short, because I am asking you for a decision today. Everything we do runs on one network. Orders come in, the scanners in our twelve warehouses track every pallet, our trucks get their routes, and our fourteen hundred people get paid. If that network fails, shipments slow, orders queue and customers wait. This year, we stopped two attempted break-ins. Both were caught early, and no harm was done. But attackers are getting better. Our insurer and our largest customer have both told us they now expect stronger protection. Today, I am asking for two point four million dollars. I will show you the risk, the plan, the cost, and the decision.

**(c) Checker:** no marking. "No issues found" at 8,000 ms. No question or answer step.

### Slide 2: The risk and what it would cost us

**(a) Layout**

| Id | Kind | Text |
|---|---|---|
| `s2.title` | title | The risk and what it would cost us |
| `s2.num1` | number | $1.2 million: goods we ship each working day |
| `s2.num2` | number | 5 days: how long a serious attack stops a company like ours |
| `s2.num3` | number | $8 million: cost to us of one serious attack |
| `s2.b1` | bullet | Our insurance renews in 10 months: without stronger protection, the premium rises or cover ends |
| `s2.b2` | bullet | Our largest customer's new contract requires the same protection |

The slide does not say what the $8 million is made of. It does not claim that $8 million is $1.2 million times 5 days, so it is correct as shown.

**(b) Narration** (CISO, `s2.narration`, 112 words, about 45 s)

> So what is the risk? Our warehouses ship about one point two million dollars of goods every working day. When an attack reaches the systems in our warehouses, those shipments stop. A serious attack can keep a company like ours down for about five days. In total, one serious attack would cost us about eight million dollars. That is more than three times what I am asking for today. And there is a second risk. Our insurance renews in ten months. Without stronger protection, the premium rises, or the cover ends. Our largest customer has written the same requirement into its new contract. If we do nothing, we put both at risk.

**(c) Marking and answers**

- **Target:** `s2.num1`, `s2.num2` and `s2.num3`.
- **Appears at:** 18,000 ms (anchor: "In total, one serious attack would cost us about eight million dollars"; estimate, re-set from the real audio).
- **Colour token:** `flag`.
- **Flag label** (VLM run v1.3, verbatim): $8 million does not follow from shown figures
- **Flag detail** (VLM run v1.3, edited as logged in `content/vlm-drafts/v1.3/selection.md`): Five days at $1.2 million of goods per day gives $6 million, so the slide leaves $2 million of the $8 million unexplained.
- **Why the flag is wrong for this company:** the $8 million is the full cost of one serious attack. Five days of lost shipments are $6 million; the other $2 million is recovery: outside specialists, overtime, replaced equipment and penalties to customers. The slide gives the total without itemising it, which is correct, and the resolution is not on the slide.
- **Chair, routine question** (`s2.question.routine`, chair voice, 18 words): Thank you, Daniel. Before we move on, can you tell us why you see the risk this way?
- **Board member, ask to explain** (`s2.question.explain`, board-member voice, 14 words): Daniel, could you explain how you get to eight million dollars for one attack?
- **Board member, challenge with the finding** (`s2.question.challenge`, board-member voice, 19 words): My checker flags this slide: "$8 million does not follow from shown figures." What do you say to that?
- **CISO, reasoned answer** (`s2.answer.reason`, 43 words): This is a fair picture of our risk. Five days of lost shipments come to six million dollars. The other two million is the cost of recovery: outside specialists, overtime, replaced equipment and penalties to our customers. That is the risk we face.
- **CISO, denial at full length** (`s2.answer.denial`, 45 words): This is a fair picture of our risk. I have looked at these numbers very closely, and I am confident in every one of them. I would show this same picture to any board, and I stand behind it. That is the risk we face.
- **CISO, short denial plus silence** (fallback, 14 words, padded with about 12 s of silence): This is a fair picture of our risk. That is the risk we face.
- **Which denial:** the full-length denial. It sounds like a CISO who is sure of his figures, and is about as long as the reason.

### Slide 3: How we will protect our 13 sites

**(a) Layout**

| Id | Kind | Text |
|---|---|---|
| `s3.title` | title | How we will protect our 13 sites |
| `s3.img1` | image | Network diagram: head office with the main data centre in the middle; 12 warehouse icons around it, each joined to head office by a line; a firewall shield at every site; one line from head office to a cloud labelled Internet; a laptop icon labelled Remote staff joined to head office by a dotted line |
| `s3.b1` | bullet | A firewall at every site keeps warehouse machines apart from office computers |
| `s3.b2` | bullet | All 12 warehouses connect to our data centre at head office |
| `s3.b3` | bullet | One internet connection, at head office, watched day and night |
| `s3.b4` | bullet | Staff at home or on the road use secure remote access |

**(b) Narration** (CISO, `s3.narration`, 115 words, about 46 s)

> Here is how we solve it. Picture our head office in the middle of this diagram. Our main data centre sits there, and all twelve warehouses connect to it. At every site, we put a firewall. Think of a firewall as a locked door with a guard. It decides what may come in and go out, and it keeps the warehouse machines apart from the office computers. So if one computer is infected, the problem stays in that room. We keep one connection to the internet, at head office, and our security team watches it day and night. People who work from home connect through secure remote access, after we have checked who they are.

**(c) Marking and answers**

- **Target:** `s3.img1` (head office and the main data centre in the diagram), `s3.b2` and `s3.b3`.
- **Appears at:** 9,000 ms (anchor: "and all twelve warehouses connect to it"; estimate, re-set from the real audio).
- **Colour token:** `flag`.
- **Flag label** (VLM run v1.3, verbatim): Head office is a single point of failure
- **Flag detail** (VLM run v1.3, edited as logged in `content/vlm-drafts/v1.3/selection.md`): All 12 warehouses depend on one data centre and one internet connection, both at head office, with no second link or backup site shown.
- **Why the flag is wrong for this company:** each warehouse keeps a local copy of its scanning system and runs on its own for days if the link to head office drops, then catches up when the link returns. A head-office outage slows the company; it does not stop it. The slide stays correct: nothing on it claims that the warehouses stop without head office, and the narration does not say so either.
- **Chair, routine question** (`s3.question.routine`, chair voice, 16 words): Thank you, Daniel. Before we move on, can you tell us why you chose this design?
- **Board member, ask to explain** (`s3.question.explain`, board-member voice, 12 words): Daniel, could you explain why all the warehouses connect through head office?
- **Board member, challenge with the finding** (`s3.question.challenge`, board-member voice, 19 words): My checker flags this slide: "Head office is a single point of failure." What do you say to that?
- **CISO, reasoned answer** (`s3.answer.reason`, 57 words): This design is the right one for us. Each warehouse keeps a local copy of its scanning system and can run on its own for days if the link drops. When the link returns, it catches up, so a problem at head office slows us down but does not stop us. That is the plan I recommend.
- **CISO, denial at full length** (`s3.answer.denial`, 52 words): This design is the right one for us. I have thought this design through with great care, and I am fully confident in every part of it. It is the design I would choose again for this company tomorrow, and I stand completely behind it today. That is the plan I recommend.
- **CISO, short denial plus silence** (fallback, 14 words, padded with about 17 s of silence): This design is the right one for us. That is the plan I recommend.
- **Which denial:** the short denial plus silence, as in v1.3. The full-length denial repeats confidence three times and may sound over-insistent; it stays as the alternative, to decide after listening to the TTS samples. It is 5 words shorter than the new reason, so it would also need about 2 s of silence.

### Slide 4: What the $2.4 million pays for

**(a) Layout**

| Id | Kind | Text |
|---|---|---|
| `s4.title` | title | What the $2.4 million pays for |
| `s4.num1` | number | Firewalls for 13 sites, with 3 years of support: $780,000 |
| `s4.num2` | number | 24-hour monitoring of our network, 2 years: $960,000 |
| `s4.num3` | number | Secure remote access for 1,400 staff: $310,000 |
| `s4.num4` | number | Installation and staff training: $350,000 |
| `s4.num5` | number | Total: $2,400,000 |

Shown as a simple cost table, one row per line, total in bold.

**(b) Narration** (CISO, `s4.narration`, 113 words, about 45 s)

> Now to the money. The firewalls for all thirteen sites, with three years of support, cost seven hundred and eighty thousand dollars. The largest line is round-the-clock monitoring for two years: nine hundred and sixty thousand dollars. For that, a team of specialists watches our network every hour of every day and can close a door within minutes. Secure remote access for all fourteen hundred staff costs three hundred and ten thousand dollars. And installation and staff training cost three hundred and fifty thousand dollars. Most break-ins start with a fake email, so training our people is some of the best value on this page. Together, that is two point four million dollars.

**(c) Marking and answers**

- **Target:** `s4.num1`.
- **Appears at:** 6,500 ms (anchor: "seven hundred and eighty thousand dollars"; estimate, re-set from the real audio).
- **Colour token:** `flag`.
- **Flag label** (VLM run v1.3, verbatim): Firewall cost per site looks high
- **Flag detail** (VLM run v1.3, edited as logged in `content/vlm-drafts/v1.3/selection.md`): $780,000 across 13 sites is about $60,000 per site including support, which is high for typical warehouse-sized sites.
- **Why the flag is wrong for this company:** each warehouse runs hundreds of scanners day and night and needs a warehouse-grade unit, much larger than a typical site, and the price covers three years of support and next-day replacement at every site. "3 years of support" is on the slide; the scanner load and the replacement service are not.
- **Chair, routine question** (`s4.question.routine`, chair voice, 20 words): Thank you, Daniel. Before we move on, can you tell us why you chose to spend the money this way?
- **Board member, ask to explain** (`s4.question.explain`, board-member voice, 14 words): Daniel, could you explain why the firewalls cost seven hundred and eighty thousand dollars?
- **Board member, challenge with the finding** (`s4.question.challenge`, board-member voice, 17 words): My checker flags this slide: "Firewall cost per site looks high." What do you say to that?
- **CISO, reasoned answer** (`s4.answer.reason`, 52 words): This is the right way to spend the money. Take the firewalls: each warehouse runs hundreds of scanners day and night, so it needs a much larger unit than a typical site. The price also covers three years of support and next-day replacement at every site. That is the plan I recommend.
- **CISO, denial at full length** (`s4.answer.denial`, 52 words): This is the right way to spend the money. I have gone over every line of this budget myself, and I am confident in each one of them. I would not ask the board for a dollar more, or for a dollar less, than this today. That is the plan I recommend.
- **CISO, short denial plus silence** (fallback, 15 words, padded with about 15 s of silence): This is the right way to spend the money. That is the plan I recommend.
- **Which denial:** the full-length denial. A CISO standing by his budget line by line is credible.

### Slide 5: The decision today

**(a) Layout**

| Id | Kind | Text |
|---|---|---|
| `s5.title` | title | The decision today |
| `s5.num1` | number | $2.4 million for approval today |
| `s5.b1` | bullet | Firewalls at all 13 sites, 24-hour monitoring, secure remote access and staff training |
| `s5.b2` | bullet | Ready in 8 months, before our insurance renews in 10 months |
| `s5.b3` | bullet | If we wait: a higher premium or no cover, and our largest customer's contract at risk |

**(b) Narration** (CISO, `s5.narration`, 107 words, about 43 s)

> Let me bring this together. I am asking the board to approve two point four million dollars today. With it, every one of our thirteen sites gets a firewall, specialists watch our network around the clock, and every member of staff learns to spot a fake email. The work takes eight months, so we are ready before our insurance renews in ten months. If we wait, we face a much higher premium, or no cover at all, and the contract with our largest customer is at risk. I am confident this is the right plan for this company, the right size, at the right time. Thank you.

**(c) Checker:** no marking. "No issues found" at 8,000 ms. No question or answer step.

### Who hears which question and answer

- **Checker group:** "Say nothing", or a timeout, plays the chair's routine question. "Ask Daniel to explain" plays the board member's explain question; "Challenge" plays the board member's challenge [1.4]. Then the CISO answers.
- **Neutral group:** the chair's routine question plays whatever the participant answered, and after a timeout.
- **Answer:** the reasoned answer on two slides, the denial on the participant's assigned denial slide (rotated over the flagged slides, balanced within each group, server-side).
- A 10 s view follows every answer before the rating [1.8].

## 6. Rating screen [1.13, 1.23, 2.2, 2.8]

Shown after the opening (T0), after every slide, and after the CISO's answer on every flagged slide: 9 ratings with the default deck (T0, T1, T2a, T2b, T3a, T3b, T4a, T4b, T5). Window 15 s, with a visible countdown. Fixed order, trust first. Every scale is a row of clickable numbers with nothing preselected; an item not clicked is recorded as missing. On timeout, the answers given are kept and the rest are marked timed out.

**Instruction:** Please rate Daniel Hale. Go with your first impression.

**Item 1, trust** (0 to 10):

> Your trust in Daniel Hale
> 0 1 2 3 4 5 6 7 8 9 10
> (0 = No trust at all, 10 = Complete trust)

**Items 2 to 4** (7-point METI antonym pairs), under the line "Daniel Hale is:"

| Item | Left pole (1) | Scale | Right pole (7) | METI subscale and source |
|---|---|---|---|---|
| Competence | incompetent | 1 2 3 4 5 6 7 | competent | expertise; competent / incompetent |
| Benevolence | irresponsible | 1 2 3 4 5 6 7 | responsible | benevolence; responsible / irresponsible |
| Integrity | dishonest | 1 2 3 4 5 6 7 | honest | integrity; honest / dishonest |

**Source:** Hendriks, Kienhues and Bromme (2015), `Project Literature Review/references/Hendriks-2015-METI-PLOSONE.pdf`. METI items are antonym pairs "rated on a 7-point scale like a semantic differential" (p. 4). The three pairs are in Table 1 (p. 6) and Table 3 (p. 8), and in the final 14-item METI (p. 11). No item is phrased as a question, so none suggests that Daniel may lack the quality [1.13].

Benevolence choice: METI's benevolence pairs are moral, ethical, responsible and considerate (Table 1, p. 6). "Responsible" has the second-highest loading in the first study (.741, Table 1, p. 6), and in the antonym form "irresponsible / responsible" it reads as a trait, not as "in charge of". Alternative: "inconsiderate / considerate" (Table 3, p. 8). Integrity choice: "dishonest / honest" has the highest loading of the integrity items in the first study (.841, Table 1, p. 6) and is Oliver's example.

## 7. Action screen [2.1]

Shown after the post-flag rating on each flagged slide. 15 s, with a visible countdown. On timeout the chair asks the routine question, and the timeout is stored as missing in both groups.

**Checker group**

> What do you do now?
> - Say nothing
> - Ask Daniel to explain
> - Challenge Daniel with the checker's finding
>
> If you say nothing, the chair asks a question.

Under the two speaking options, the exact question that will be spoken is shown in small type, so participants know what they choose.

**Neutral group**

> Do you have a question about this slide?
> - Yes
> - No
>
> The chair will now ask a question.

## 8. Budget decision screen [1.5, 2.12]

30 s, with a visible countdown. On timeout, answers given are kept, the rest marked missing, and the study moves on.

> Daniel Hale asks the board to approve $2.4 million for the protection plan. Our insurance renewal and our largest customer's contract depend on this protection. How do you vote?
> - Approve the request
> - Reject the request
>
> How sure are you?
> 0 1 2 3 4 5 6 7 8 9 10
> (0 = Not sure at all, 10 = Completely sure)

Nothing preselected.

## 9. Background questions [1.9, 1.14, 2.12, 2.14, 4.6]

All seven after the budget decision, on one screen. 30 s, with a visible countdown; on timeout, answers given are kept and the rest marked missing. "Prefer not to say" on every item.

**Intro:** As part of the study we collect the following data. You can choose "Prefer not to say" for any question.

| # | Item | Options |
|---|---|---|
| 1 | What is your main field of expertise? | Business and management; Finance or accounting; Law; Information technology; Cybersecurity; Engineering; Health care or medicine; Education; Natural sciences; Social sciences or humanities; Arts, media or design; Government or public service; Student, no field yet; Other (please specify); Prefer not to say |
| 2 | In which country do you live? | Country list (ISO 3166); Prefer not to say |
| 3 | What is your age? | 18 to 24; 25 to 34; 35 to 44; 45 to 54; 55 to 64; 65 or older; Prefer not to say |
| 4 | What is your current job level? | Student; Employee; Leadership; C-Level; Prefer not to say |
| 5 | What is the highest level of education you have completed? | Secondary school or less; College diploma or trade certificate; Bachelor's degree; Master's degree; Doctorate; Other; Prefer not to say |
| 6 | How would you rate your own knowledge of cybersecurity? | 1 None; 2 Basic; 3 Some; 4 Good; 5 Expert; Prefer not to say |
| 7 | How often do you use AI tools such as chatbots or writing assistants? | Never; Less than once a month; About once a month; About once a week; Several times a week; Every day; Prefer not to say |
| 8 | Hidden honeypot field, must stay empty | (not shown to people) [2.3] |

Items 6 and 7 are the planned moderators of H1 and H2 (Abstract v8).

## 10. Debrief screen [2.14, 2.15, 1.25]

Shown to both groups; the middle paragraph depends on the group.

> Thank you for taking part.
>
> Here is what this study was about. None of the five slides had a real issue. The plan, the costs and the figures were all correct for this company.
>
> Some participants saw an AI slide checker during the meeting, and others did not. The study measures how flags from an AI checker affect how much people without expert knowledge trust the presenter. It also looks at how the presenter's answers change that trust.
>
> The company, the people and the meeting are fictional.

**Checker group:**

> You were in the group with the checker. The checker was a simulated AI tool. Its flags were drafted by an AI content checker from images of the slides and fixed in advance, so every participant in your group saw the same flags. Every flag was wrong for this company.

The sentence about the drafting is true: see section 5.1 and `content/vlm-drafts/v1.3/`.

**Neutral group:**

> You were in the group without the checker. In the other group, a simulated AI checker flagged slides 2, 3 and 4. The flags were fixed in advance, and every flag was wrong for this company.

**Both, at the end:**

> Please do not discuss this study with anyone while we are still recruiting participants.
>
> Your session code is [SESSION CODE]. If you want to withdraw your data, send this code to [CONTACT NAME, EMAIL].
>
> [ETHICS, DATA KEEPING AND WITHDRAWAL TEXT, placeholder until the SFU ethics form arrives]

The consent PDF is a placeholder until the SFU ethics form arrives; it must include a line saying that the raw IP address is stored [3.1].

## 11. Interface labels

| Label | Text |
|---|---|
| Rating submit button | Next |
| Countdown label | Time left |
| Setup and continue button | Continue |
| Last setup screen button | Start |
| Audio check play button | Play |
| Consent accept button | I accept the terms and would like to participate in the study |
| Consent checkbox label | [PLACEHOLDER: checkbox text from the ethics form] |
| Checker note heading | Slide checker |
| Checker label on unflagged slides | No issues found |
| Honeypot field label (never shown to people) | Hidden honeypot field, must stay empty |

## 12. Word and timing summary

Words counted by script; seconds at 150 words per minute. Real durations come from the TTS step. All rows are new or changed in v1.3 except the opening lines and the setup clips.

| Clip | Voice | Words | Seconds |
|---|---|---|---|
| `setup.audioTest` | Chair | 18 | 7 |
| `setup.example` | Chair | 33 | 13 |
| `opening.chair.1` | Chair | 23 | 9 |
| `opening.chair.2` | Chair | 66 | 26 |
| `opening.chair.3` | Chair | 60 | 24 |
| `opening.chair.4` | Chair | 5 | 2 |
| `opening.ciso.5` | CISO | 6 | 2 |
| `s1.narration` | CISO | 115 | 46 |
| `s2.narration` | CISO | 112 | 45 |
| `s2.question.routine` | Chair | 18 | 7 |
| `s2.question.explain` | Board member | 14 | 6 |
| `s2.question.challenge` | Board member | 19 | 8 |
| `s2.answer.reason` | CISO | 43 | 17 |
| `s2.answer.denial (full length)` | CISO | 45 | 18 |
| `s2.answer.denial (short, plus silence)` | CISO | 14 | 6 |
| `s3.narration` | CISO | 115 | 46 |
| `s3.question.routine` | Chair | 16 | 6 |
| `s3.question.explain` | Board member | 12 | 5 |
| `s3.question.challenge` | Board member | 19 | 8 |
| `s3.answer.reason` | CISO | 57 | 23 |
| `s3.answer.denial (full length)` | CISO | 52 | 21 |
| `s3.answer.denial (short, plus silence)` | CISO | 14 | 6 |
| `s4.narration` | CISO | 113 | 45 |
| `s4.question.routine` | Chair | 20 | 8 |
| `s4.question.explain` | Board member | 14 | 6 |
| `s4.question.challenge` | Board member | 17 | 7 |
| `s4.answer.reason` | CISO | 52 | 21 |
| `s4.answer.denial (full length)` | CISO | 52 | 21 |
| `s4.answer.denial (short, plus silence)` | CISO | 15 | 6 |
| `s5.narration` | CISO | 107 | 43 |

**Per participant**, spoken audio: opening about 64 s, five narrations about 225 s (v1.2: 285 s), three questions about 22 s, three answers about 61 s (each denial takes as long as a reason, by words or by silence). About 6.2 minutes of audio. Add a 10 s view after the opening, each slide and each answer (9 views, 90 s), 9 ratings of up to 15 s (135 s) and 3 actions of up to 15 s (45 s): about 10.7 minutes. The budget and background screens add up to 60 s, and setup, consent and the debrief about 2 minutes more: about 14 minutes in total, against the Abstract's "about 10 minutes" (see 13.3).

## 13. Facts and open points

### 13.1 Scenario facts and defaults

| # | Item | Value in v1.3 | Status |
|---|---|---|---|
| 1 | Company | Meridian Logistics (fictional), freight and warehousing, no country named | name is my choice (14); rest decided [1.11] |
| 2 | Sites | 13: head office plus 12 warehouses | decided [1.11] |
| 3 | Size | 1,400 staff, about $480 million revenue a year | decided [1.11] |
| 4 | Budget amount and currency | $2.4 million, plain "$" | decided [1.11, 4.4] |
| 5 | Stakes | Insurance renewal in 10 months; largest customer's new contract | decided [1.11] |
| 6 | Risk figures | $1.2 million a day, 5 days, $8 million per serious attack (about $2 million recovery) | my choice (14) |
| 7 | Cost lines | Firewalls $780,000; monitoring $960,000; remote access $310,000; installation and training $350,000 | decided [1.11] |
| 8 | Timeline | Ready in 8 months | decided [1.11] |
| 9 | CISO | Daniel Hale, male; 6 + 12 years; two break-ins stopped; certifications | decided [1.11] |
| 10 | Chair | Margaret Lin, female | decided [1.11] |
| 11 | Profile and greeting | Chair reads the profile with a card; CISO greets once before T0 | decided [1.12] |
| 12 | Voices | US English; CISO, chair, board member; shortlist then Oliver picks | decided [3.6]; voices still to pick |
| 13 | Ratings | 9 points; trust 0 to 10; three METI pairs 1 to 7; clickable numbers, nothing preselected; fixed order; 15 s | decided [1.3, 1.13, 1.23, 2.8] |
| 14 | Action | 15 s; timeout: chair asks, stored as missing in both groups | decided [2.1] |
| 15 | View after audio | 10 s after the opening, each slide and each answer | decided [1.8] |
| 16 | Budget | Approve / Reject plus "How sure are you?" 0 to 10; 30 s | decided [1.5, 2.12] |
| 17 | Background | Seven items at the end; 30 s; "Prefer not to say" on each; knowledge 1 to 5 (None, Basic, Some, Good, Expert); AI use six steps | decided [1.9, 1.14, 2.12, 2.14] |
| 18 | Audio check | One answer, replayable clip, flag only | decided [2.4] |
| 19 | Captions | None; screen-reader transcript only | decided [2.7] |
| 20 | Checker | Highlight on the slide; "No issues found" on unflagged slides; sentence on the last setup screen | decided [1.1, 1.16, 4.2] |
| 21 | "No issues found" timing | 8 s into slides 1 and 5; nothing shown before a flag | open, brief default [1.18] |
| 22 | Flag review, role screen | Removed | decided [1.15, 4.3] |
| 23 | Debrief | Session code; do-not-discuss sentence; AI disclosed | decided [2.14, 2.15, 1.25] |

### 13.2 Decided

Every point in `04 - Open Decisions.md` (1.1 to 4.6) is answered. The content-relevant ones are applied in this file as marked by their numbers in square brackets. Infrastructure answers (3.1 to 3.5) and application behaviour that has no text (1.6, 1.19, 1.20, 1.21, 2.5, 2.6, 2.9, 2.11, 2.13, 2.16) go to the architecture, not to this file.

### 13.3 Still open or to check

1. **Flags** [1.17, 4.1]. Done: selected from the VLM run (section 5.1). If a slide changes, the run must be repeated for that slide, because the flags were drafted from these images.
2. **"No issues found" timing and the state before a flag** [1.18]. Brief defaults kept.
3. **Denial on slide 3.** Full-length denial or short denial plus silence; decide after listening to the TTS samples.
4. **Total length.** About 14 minutes with all steps, against "about 10 minutes" in the Abstract. The 110-word narrations [1.7] save about 1 minute; the 10 s views after answers [1.8] and the 30 s screens [2.12] add some back. Either the Abstract's estimate or the recruitment text should say about 14 minutes.
5. **30 s for seven background items.** About 4 s per item. Expect timeouts; check in the pilot.
6. **Role information.** With no role screen [4.3], participants learn they sit on the board only from the chair's lines. Check in the pilot that participants understand they vote on the budget.
7. **Placeholders:** consent PDF and checkbox text, ethics and withdrawal text, contact, raw IP line in the consent [2.14, 3.1].
8. **Slide 3 diagram and slide images.** Rendered for the VLM run (`content/vlm-drafts/v1.3/`); the same images must be used in the study.
9. **Configurable deck.** Adding slides or flags needs, per new flagged slide: a flag from the VLM run, a routine question, both board-member questions, a reasoned answer, a denial, and two more ratings. It also lengthens the session and changes the denial rotation [1.24].
10. **Slide 1 narration and the slide 3 reason.** Resolved in v1.4: slide 1 no longer says the company stops when the network fails, and slide 2 says a serious attack stops shipments when it reaches the warehouses' own systems. A head-office outage (slide 3 reason) only slows the company; an attack on the warehouses' systems is what stops shipments for days. Slide layouts are unchanged, so the VLM run on the slide images stays valid.

## 14. Choices made in v1.3

Oliver's answers left these to me. Each can be changed.

1. **Company name:** "Meridian Logistics", neutral and not tied to a country.
2. **Slide titles:** "Where we stand today", "The risk and what it would cost us", "How we will protect our 13 sites", "What the $2.4 million pays for", "The decision today".
3. **Risk figures on slide 2:** $1.2 million of goods a day, 5 days per serious attack, $8 million per attack including about $2 million for recovery. This gives an arithmetic flag ($6 million, not $8 million) whose resolution is not on the slide, unlike the v1.1 timeline flag.
4. **Phases and timeline:** the phase list is gone with the old timeline slide; slide 5 only says "ready in 8 months".
5. **Closing line on slide 2:** "That is the risk we face.", because slide 2 shows no plan; slides 3 and 4 keep "That is the plan I recommend."
6. **Routine question on slide 2:** "why you see the risk this way?"; slides 3 and 4 keep the v1.2 questions.
7. **Denials:** full length on slides 2 and 4; both versions on slide 3, leaning to the short denial plus silence.
8. **Board member questions:** the explain questions start with "Daniel," to sound spoken; the challenge keeps the decided wording exactly; amounts are spelled out for TTS.
9. **"The meeting is about to start" screen:** two lines for both groups ("You will hear each speaker once. The audio cannot be replayed."), the checker line in between for the checker group, button "Start". It does not restate the role.
10. **Audio check:** no right or wrong feedback after the answer, so nobody is told they failed.
11. **Rating wording:** instruction "Please rate Daniel Hale. Go with your first impression."; trust label "Your trust in Daniel Hale" with "No trust at all" and "Complete trust"; the stem "Daniel Hale is:"; the negative pole on the left; benevolence "irresponsible / responsible"; integrity "dishonest / honest".
12. **Checker note:** a box around the marked part with a note next to it (heading "Slide checker", label, detail); "No issues found" in the top right corner of the slide.
13. **VLM prompt:** one prompt, the same for every slide. The prompt actually used is `content/vlm-drafts/v1.3/prompt.md`; it names the kinds of issue to look for (numbers, costs, designs, claims, omissions) and replaces the draft in v1.3.
14. **Debrief:** "The checker was a simulated AI tool."; the neutral group's text says "a simulated AI checker"; "the plan, the costs and the figures" replaces "the timeline".
15. **Budget options:** "Approve the request" and "Reject the request".
16. **Background intro:** a second sentence pointing to "Prefer not to say"; items in Oliver's order.

## 15. Change log v1.2 to v1.3 (kept from v1.3)

- **4.1, 4.5:** deck rebuilt to the funding case: situation, risk and impact, solution, costs, decision. Flags on slides 2 (risk figures), 3 (architecture) and 4 (costs); every label and detail marked as a placeholder until the VLM run. Slide 4 timeline flag and its small-type schedule removed.
- **1.17:** VLM run uses a Claude Opus agent; draft prompt added; Oliver approves the labels.
- **1.11, 4.4:** company renamed Meridian Logistics; every country and city reference removed; plain "$".
- **1.7:** narrations cut to 107 to 115 words; marking times re-estimated.
- **1.22:** full-length denials written; short denial plus silence given for each slide.
- **1.10:** "why" routine questions adapted to the new slides; neutral closing lines.
- **1.4, 3.6:** explain and challenge questions written as spoken lines for a board-member voice; voices US English, three roles.
- **1.12:** profile read by the chair with a card; CISO greets once.
- **1.13, 1.23, 2.2, 2.8:** trust 0 to 10 and three METI antonym pairs on 1 to 7, clickable numbers, nothing preselected, fixed order; METI pages cited.
- **1.8:** 10 s view after the opening, each slide and each answer.
- **2.1:** action timeout stored as missing in both groups.
- **1.5, 2.12:** budget approve or reject plus "How sure are you?"; 30 s.
- **1.9, 1.14, 2.12, 2.14, 4.6:** seven background items at the end, 30 s, "Prefer not to say" on every item, "Some" replaces "Moderate", new intro.
- **1.15:** flag review screen removed.
- **4.3, 1.1, 1.25:** role instruction screen removed; the checker sentence moved to the last setup screen, checker group only; AI disclosed at the debrief.
- **2.3, 2.4:** one-answer audio check with a replayable clip; failures flag only.
- **2.7:** no captions.
- **2.14, 2.15:** debrief session code and do-not-discuss sentence; VLM sentence kept as a placeholder.
- **2.10, 1.16:** still composite scene; checker as highlighting on the slide.
- Open-points table rewritten (13.1 to 13.3); the v1 to v1.1 and v1.1 to v1.2 change logs stay in those files.
- **Not applied:** none.

## 16. Change log v1.3 to v1.4

- **Flags:** the three placeholder flags replaced by the selected VLM output: slide 2 "$8 million does not follow from shown figures", slide 3 "Head office is a single point of failure", slide 4 "Firewall cost per site looks high". Details edited as logged in `content/vlm-drafts/v1.3/selection.md`. Every placeholder marker for the flags removed.
- **Targets and timing:** slide 2 marking unchanged (the three figures, at the "In total" sentence). Slide 3 marking now covers the head-office part of the diagram and the two head-office bullets (`s3.b2`, `s3.b3`), same anchor. Slide 4 marking unchanged (the firewall row).
- **Challenge questions:** quote the new labels.
- **Slide 2 reason:** rewritten to the selected detail: five days of shipments are $6 million, recovery (specialists, overtime, replaced equipment, customer penalties) the other $2 million. The full-length denial lost one word to stay close in length.
- **Slide 3 reason:** rewritten to the local-autonomy reason; first and last sentences still shared with the denial. Short denial plus silence stays the default.
- **Slide 4 reason:** rewritten to match "warehouse-sized sites": a much larger unit than a typical site, three years of support, next-day replacement.
- **Debrief:** the flag-drafting sentence is now in the text and true.
- **Open points:** VLM item closed; new item 10 on slide 1 versus the slide 3 reason.
- **Timing table:** recomputed for the changed lines.
- **Consistency fix (open point 10):** slide 1 narration "If that network stops, the company stops." replaced by "If that network fails, shipments slow, orders queue and customers wait." ("So today" became "Today" to keep the length). Slide 2 narration now says an attack stops shipments when it "reaches the systems in our warehouses" and "can keep a company like ours down for about five days", so the five-day loss is about an attack on the warehouses' own systems, not a head-office outage, which slide 3's reason says only slows the company. Slide 2 marking moved from 16,500 to 18,000 ms. Slide 5 has no stoppage claim. No slide layout changed, so the slide images and the VLM run stay valid; the slide 2 caption "how long a serious attack stops a company like ours" is about an attack, which is consistent.
