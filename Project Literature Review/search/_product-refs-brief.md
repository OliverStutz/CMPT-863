# Brief: verify the Meta (2024) and Microsoft (n.d.) references in v8, section 1, paragraph 1 (2026-10-20)

One Opus agent. Output: `search/product-refs-check.md`. Edit no other file.

## The sentence under check

From `Literature Review - Draft v8.md`, paragraph 1 of "Cybersecurity oversight by non-expert boards":

"The technology for an independent view is already on the market: smart glasses answer questions about what the wearer sees (Meta, 2024), and Copilot in Teams, asked during a meeting, will say where the holes in a participant's argument lie (Microsoft, n.d.)."

The two reference entries in the same file:

Meta. (2024, April 23). *New Ray-Ban | Meta smart glasses styles and Meta AI updates*. https://about.fb.com/news/2024/04/new-ray-ban-meta-smart-glasses-styles-and-meta-ai-updates/

Microsoft. (n.d.). *Catch up on meetings with Microsoft Copilot in Teams* [Support article]. Retrieved October 20, 2026, from https://support.microsoft.com/en-us/teams/copilot/catch-up-on-meetings-with-microsoft-365-copilot-in-teams

(The Microsoft entry was changed today from the older URL https://support.microsoft.com/en-us/office/use-copilot-in-microsoft-teams-meetings-0bf9dd3c-96f7-44e2-8bb8-790bedf066b1, which now redirects. An earlier check found the suggested prompt "Where are some holes in [a meeting participant's] argument?" under the heading "Have a more effective meeting", and the page calling the pane "your private chat with Copilot". Earlier search notes on both products are in `search/results-5-board-products.md`.)

## What to check

1. **Meta.** Fetch the page. Confirm the date, the exact title (including the vertical bar), the author (Meta as organisation; is there a named author or byline?), and that the page supports "smart glasses answer questions about what the wearer sees" (quote the sentence that does). Is this the best Meta source for that claim, or is there a later or more direct one (for example a 2025 or 2026 Meta newsroom post on Ray-Ban Meta or Meta AI "look and ask" features, or the product page)? If a better one exists, give its verified APA entry and the sentence it supports. Decide whether a more precise wording of the claim is needed.
2. **Microsoft.** Fetch the current page. Confirm the title, whether the page shows any date (published, last updated, applies-to) that would replace "n.d." under APA 7, the exact prompt wording and the heading it sits under, and whether the page supports "asked during a meeting, will say where the holes in a participant's argument lie" (quote the text). Note that Copilot in Teams needs a Microsoft 365 Copilot licence: does the page say so, and does the sentence's "on the market" hold? Check whether the old URL redirects to this page (curl -I). Give the correct APA 7 entry for a support article with no date and a retrieval date, and the right in-text locator form if a quotation were used (section heading).
3. **Both.** Judge whether the sentence states what the pages say, and propose exact replacement wording if not. Judge whether the two product sources belong in this paragraph at all, given that the review's section on checking tools (section 2 of `Literature Review - Draft v6.md`) also cites them.

## Rules

Verify from the live pages (WebFetch, curl). Nothing from memory. Plain wording, no em-dashes. If a page cannot be reached, say so and try the Internet Archive.

## Output layout

1. Verdict per source: entry correct / imprecise / wrong, claim supported / partly / not, with the exact fix.
2. Evidence quoted from each page, with headings.
3. Alternative sources, if any, with verified APA entries.
4. Recommended sentence and recommended entries.
