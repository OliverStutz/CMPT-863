# Brief for the three independent selection reviewers (Project Literature Review)

Written 2026-10-18. You are one of three reviewers. Do not assume the others exist; give your own judgement. Your job is to pick the papers the literature review should cite, group them into thematic subsections, and say what each subsection can claim. The selection will be cross-examined: papers chosen by all three reviewers are the decisive winners.

## Files to read, in this order

1. `../Instructions.md`: the assignment. 1,000 to 1,500 words, at least 15 sources, thematic subsections that synthesize (what exists, what we know, what is unresolved, how our project differs). APA citations.
2. `../../CMPT 863 Project Abstract v7.md`: the submitted abstract with the research question, hypotheses H1 to H4, method, and the current related-work appendix (8 sources, all of which stay available for the review).
3. `_agent-briefs.md`: the project summary and the rubric used by the six search agents (sections "The project in short", "Common rules", "Rubric").
4. `_known-papers.md`: the 96 papers already known before this search (8 cited in the abstract, 75 from the earlier candidate list, 5 others). These are candidates too.
5. The six result files, `results-1-hci-checkers.md` to `results-6-factchecking.md`: about 120 new verified entries with design, finding, relation and score. Read all six in full.
6. `../../CMPT 863 Idea 2 - Related papers.md`: the earlier 75-entry list with a one-paragraph "About" and "Relation" per paper, for the known candidates.

Do not use web search. Judge from these files. Where you add knowledge from outside the files, label it as such.

## What the review has to achieve

The review situates a study in which a board member (non-expert, holds the power and the budget) uses a vision-language-model content checker during a CISO's briefing; the checker's flags are wrong for the company; the CISO answers with a reason or a bare denial; trust in the CISO is rated before and after each step. The hypotheses: H1 a flag lowers trust; H2 a reasoned answer restores it, possibly above the start; H3 a denial without a reason restores less or lowers further; H4 checker participants question more, and the same answer then moves trust further.

A marker will look for: (a) the HCI literature on AI checkers, credibility indicators and AI-mediated communication, with SIGCHI venues well represented; (b) the psychology of interpersonal trust, its dimensions and its repair; (c) evidence on what a flag, correction or AI disagreement does to the evaluation of the person flagged; (d) the boardroom and cybersecurity context, including industry sources and products where appropriate; (e) a clear gap and a clear statement of how this study differs. The strongest reviews use counter-results, not only supporting ones.

## Your task

### 1. Propose the thematic subsections

Four to six subsections, each with a one-sentence claim the subsection will make. Example form: "Monitoring and checking lower trust in the person checked even when nothing is found (Strickland 1958; ...)."

### 2. Select the papers

Choose 22 to 30 papers in total, across old and new candidates. For each: the short citation (First author et al., year), the subsection it serves, its role (anchor, supporting finding, counter-result, measure, design element, context, product, gap), and one sentence on why it earns its place over the alternatives. Mark each as **must cite** (the review is incomplete without it), **should cite** (strengthens a claim) or **optional** (if space allows).

Rules for selection:
- Keep every paper from the abstract v7 unless you argue it out.
- Prefer a paper that measures trust or credibility in a *person* over one that measures belief in content, where both exist for the same point.
- Prefer SIGCHI and HCI venues where a psychology paper and an HCI paper make the same point, but do not drop the founding psychology source for a definition or a mechanism.
- Include at least two counter-results (findings that cut against H1 to H4).
- Include the trust definition and measure the study will use (ability or competence, benevolence, integrity).
- Include at least one industry or trade source and at least one product if the result files offer a credible one; say if they do not.
- Do not pick two papers that make the same point in the same setting unless one is the replication of the other and the replication matters.

### 3. Must-read list

List the papers whose full text must be read before the review is written because the result files flag an unverified detail, a secondary-source summary, or a finding that changes a hypothesis. Give the detail that needs checking.

### 4. Gap and differentiation

In five to eight sentences, state the gap as the files support it, naming the two or three closest designs and how each differs from ours. Say whether any paper in the files closes the gap or forces a narrower claim.

### 5. Hypothesis check

For each of H1 to H4: which selected papers support it, which cut against it, and whether the files suggest a rewording. One or two sentences each.

## Output

Write Markdown with the headings 1 to 5 above to the file named in your task. Plain wording: name what was measured on whom; no statistics terms as a stand-in for a finding; no em-dashes. End with a table of your full selection: `| Short citation | Subsection | Role | Priority | New or known |`.
