# Flag selection from the VLM run v1.3

Written 2026-10-24. Oliver delegated the choice to the coordinator; this file records it. Inputs: `prompt.md` (the fixed prompt, the same for every slide), `run.json` (Claude Opus via a Claude Code agent, 2026-10-24, one pass per slide, answers not revised), and the raw outputs `s1.raw.md` to `s5.raw.md`. The raw outputs are not edited. The flags below are used in `Prototype/03 - Study Content v1.4.md` and `Prototype/content/study-config.v1.4.json`.

Every flag is shown as an observation about the slide. None mentions the presenter or honesty, and none carries a confidence or accuracy figure; the model's output had none either.

## Slide 2: The risk and what it would cost us

**Chosen:** finding 1 (the model's most important issue on this slide).

**Model's label, verbatim:** $8 million does not follow from shown figures

**Model's detail, verbatim:** Five days at $1.2 million of goods per day gives $6 million, so the slide leaves $2 million of the $8 million unexplained. Also, the value of goods shipped is not the same as money lost, since delayed shipments may be recovered and only margin or penalties are truly lost.

**Final label:** $8 million does not follow from shown figures

**Final detail:** Five days at $1.2 million of goods per day gives $6 million, so the slide leaves $2 million of the $8 million unexplained.

**Edits:**

1. Label: none.
2. Detail: second sentence deleted ("Also, the value of goods shipped is not the same as money lost, since delayed shipments may be recovered and only margin or penalties are truly lost."). Reason: it raises a second, separate concern, and the CISO's answer addresses the $2 million gap.

**Why it is wrong for Meridian:** the $8 million is the full cost of one serious attack. Five days of lost shipments are $6 million; the other $2 million is recovery: outside specialists, overtime, replaced equipment and penalties to customers. The slide gives the total without itemising it, which is correct, so the arithmetic on the two figures shown misses part of the cost.

**Alternatives considered:**

- Finding 6, "Working days versus calendar days ambiguous": not chosen, because it points at a wording ambiguity rather than a claim the CISO can answer with a clear reason, and its effect on the total is uncertain.

## Slide 3: How we will protect our 13 sites

**Chosen:** finding 1.

**Model's label, verbatim:** Head office is a single point of failure

**Model's detail, verbatim:** All 12 warehouses depend on one data centre and one internet connection, both at head office, with no second link or backup site shown. If head office is hit by an attack, power loss or line outage, every site stops at once, which is the outcome the previous risk framing warns about.

**Final label:** Head office is a single point of failure

**Final detail:** All 12 warehouses depend on one data centre and one internet connection, both at head office, with no second link or backup site shown.

**Edits:**

1. Label: none.
2. Detail: second sentence deleted ("If head office is hit by an attack, power loss or line outage, every site stops at once, which is the outcome the previous risk framing warns about."). Reason: it refers to the previous slide and states a consequence ("every site stops at once") that is the part the CISO's answer corrects; the first sentence alone is the observation.

**Why it is wrong for Meridian:** every warehouse keeps a local copy of its scanning system and runs on its own for days if the link to head office drops, then catches up when the link returns. A head-office outage slows the company; it does not stop it. The slide stays correct: nothing on it claims that the warehouses stop without head office.

**Alternatives considered:**

- Finding 4, "Perimeter-only approach misses common controls": not chosen, because the plan does include training and secure remote access (slide 4), so the flag would be answered by another slide rather than by a reason about this company, and it reads as a list of omissions rather than one claim.

## Slide 4: What the $2.4 million pays for

**Chosen:** finding 3.

**Model's label, verbatim:** Firewall cost per site looks high

**Model's detail, verbatim:** $780,000 across 13 sites is about $60,000 per site including support, which is high for typical warehouse-sized sites, though it may be reasonable if the head office unit is much larger. A per-site or head office versus warehouse split would make this checkable.

**Final label:** Firewall cost per site looks high

**Final detail:** $780,000 across 13 sites is about $60,000 per site including support, which is high for typical warehouse-sized sites.

**Edits:**

1. Label: none.
2. Detail: hedge deleted (", though it may be reasonable if the head office unit is much larger"). Reason: it half-answers the flag on the checker's behalf.
3. Detail: second sentence deleted ("A per-site or head office versus warehouse split would make this checkable."). Reason: it is a request for more information, not part of the observation.

**Why it is wrong for Meridian:** each warehouse runs hundreds of scanners day and night and needs a warehouse-grade unit, much larger than a typical site's, and the price covers three years of support and next-day replacement at every site.

**Alternatives considered:**

- Finding 1, "Support and monitoring periods do not match": not chosen, because it is about how the costs are presented over time, and the model itself notes the lines add up; it is harder for a lay board member to judge and does not fit a short cost reason.

## Slides 1 and 5: output on record, not shown

The model reported issues on slides 1 and 5 as well. By design, the checker shows "No issues found" on these slides (Abstract v8: slides 1 and 5 show "No issues found"; the output is fixed in advance so all participants see the same). The model's first label on each, for the record:

- Slide 1, finding 1: "Budget figure has no breakdown or timeframe" (6 findings in total).
- Slide 5, finding 1: "Two-month margin before renewal is thin" (5 findings in total).

These are not used. The debrief says that the flags were drafted from images of the slides and fixed in advance; it does not claim the model found nothing on slides 1 and 5.
