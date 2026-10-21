# Product references check: Meta (2024) and Microsoft (n.d.) in v8, section 1, paragraph 1 (2026-10-20)

Brief: `search/_product-refs-brief.md`. All pages fetched live on 2026-10-20 with curl and WebFetch. Nothing from memory.

**Note before the verdicts.** The sentence in the current `Literature Review - Draft v8.md` (line 7) is not the one quoted in the brief. It now ends with four more words:

> "... and Copilot in Teams, asked during a meeting, will say where the holes in a participant's argument lie (Microsoft, n.d.) no matter the actual context."

Neither page supports "no matter the actual context" (see 1.2). The checks below cover both the brief's version and the current text.

## 1. Verdict per source

### 1.1 Meta (2024, April 23)

- **Entry: correct.** The title, date, author and URL all match the live page. The URL returns HTTP 200 with no redirect. The page carries a second date, "October 22, 2025" (a last-updated date; `article:modified_time` 2025-10-22). The cited passage describes the April 2024 rollout ("Starting today ..."), so the original date of 2024, April 23 is the right one to keep. The byline shown on the page is "Meta" (the structured data gives "Meta Newsroom"). No person is named; the HTML `author` meta tag "isolomons" is a CMS account, not a byline. The title shows as "New Ray-Ban | Meta Smart Glasses Styles and Meta AI Updates", with the vertical bar. In sentence case the entry's version is correct.
- **Claim "smart glasses answer questions about what the wearer sees": supported.** The page says the feature reached all Ray-Ban Meta glasses in the US and Canada, in beta. "Answer questions" is a fair plain wording of "ask your glasses about what you're seeing, and they'll give you smart, helpful answers".
- **Fix:** none to the entry. Naming the product would be more precise ("Ray-Ban Meta smart glasses ..."), and v6 section 2 already does that.

### 1.2 Microsoft (n.d.)

- **Entry: correct, with one optional tidy-up.**
  - The title matches the page `<title>` and H1: "Catch up on meetings with Microsoft Copilot in Teams".
  - The page shows no date to the reader: no published, last-updated or copyright-year line for the article. "Applies To" names only "Microsoft Teams". The HTML holds hidden metadata (`ms.date` 12/23/2025; `updated_at` 2026-09-24), but nothing on the page shows it, so "n.d." plus a retrieval date is right under APA 7.
  - The old URL redirects with `HTTP/2 301` to the new path, which returns `HTTP/2 200`.
  - The optional tidy-up: the bracketed "[Support article]" is not needed for an ordinary webpage under APA 7. It is harmless, but it can be dropped.
  - The site name ("Microsoft Support") can be left out because it is the same as the author.
- **Claim "asked during a meeting, will say where the holes in a participant's argument lie": partly supported.** The page lists "Where are some holes in [a meeting participant's] argument?" as a suggested prompt, and says Copilot answers questions "in real time during or after a meeting" by "reasoning over the discussion and meeting chat". The page does not show or promise what the answer contains. "Will say where the holes ... lie" turns a suggested prompt into a promised result. The page does not give the private-pane detail, which is the one that matters for the review: only the asker sees the answer.
- **Added clause "no matter the actual context": not supported.** If anything the page says the opposite: Copilot answers by "reasoning over the discussion and meeting chat", and without live transcription it cannot be prompted about the meeting. The clause reads as the author's opinion presented as the source's content. Delete it, or move it into a separate sentence with its own evidence (for example the accuracy figures for checkers in v6 section 2).
- **Licence and "on the market".** The support page does not say outright that a user needs a Microsoft 365 Copilot licence. It only implies one: "Even if you don't have a license, you're in control of Copilot for meetings you organize", and "Intelligent recap is also available as part of the Microsoft Copilot license". Microsoft's admin page on Microsoft Learn states it directly (quoted in 2.3). "On the market" holds: it is a paid add-on that businesses can buy. It is not a consumer feature that every Teams user has.

## 2. Evidence quoted from the pages

### 2.1 Meta newsroom post

URL: https://about.fb.com/news/2024/04/new-ray-ban-meta-smart-glasses-styles-and-meta-ai-updates/ (HTTP 200, no redirect)

- Byline and dates as displayed under the H1: "Meta", "April 23, 2024", "October 22, 2025".
- Under **"Takeaways"**: "We're rolling out Meta AI with Vision, so you can ask your glasses about what you're seeing and get helpful information" (followed on the page by "completely hands-free").
- Under **"Meta AI Makes Your Smart Glasses Even Smarter"**: "We started testing a multimodal AI update in December, so you can ask your glasses about what you're seeing, and they'll give you smart, helpful answers or suggestions. That means you can do more with your glasses because now they can see what you see. Starting today, we're rolling this functionality out to all Ray-Ban Meta smart glasses in the US and Canada in beta."
- Same section: "Say you're traveling and trying to read a menu in French. Your smart glasses can use their built-in camera and Meta AI to translate the text for you".

### 2.2 Microsoft support article

URL: https://support.microsoft.com/en-us/teams/copilot/catch-up-on-meetings-with-microsoft-365-copilot-in-teams (HTTP 200). The old URL https://support.microsoft.com/en-us/office/use-copilot-in-microsoft-teams-meetings-0bf9dd3c-96f7-44e2-8bb8-790bedf066b1 returns `HTTP/2 301` with `location: /en-us/teams/copilot/catch-up-on-meetings-with-microsoft-365-copilot-in-teams`.

- H1: "Catch up on meetings with Microsoft Copilot in Teams". "Applies To: Microsoft Teams". The page shows no date.
- Under **"How to use Copilot in a meeting"**: "Copilot summarizes key discussion points, including who spoke and what they said. It also suggests action items and answers any questions you have, in real time during or after a meeting." Note: "to prompt Copilot about the meeting or see Copilot's conversation history after the meeting ends, make sure to turn on live transcription during the meeting." "Copilot won't work in meetings that are hosted outside the participant's organization."
- Under **"Set up meetings with Copilot"**: "Even if you don't have a license, you're in control of Copilot for meetings you organize."
- Under **"Have a more effective meeting"**: "Copilot helps you stay on track by reasoning over the discussion and meeting chat to answer to your questions about the meeting. Select Copilot from the meeting controls to start your private chat with Copilot." Then: "Select View prompts in the Copilot pane to see prompt suggestions or ask Copilot questions of your own using the compose box. Here are some ideas:", and among the six listed ideas: "Where are some holes in [a meeting participant's] argument?" The ideas just before it are "Where do we disagree on this topic?", "How did [a meeting participant] respond to this proposal?" and "What questions can I ask to move the meeting forward?".
- Under **"From the Recap tab"** (note): "Intelligent recap is also available as part of the Microsoft Copilot license."

### 2.3 Licence statement (supporting page, Microsoft Learn)

URL: https://learn.microsoft.com/en-us/microsoftteams/copilot-teams-transcription (HTTP 200; the page shows "Last updated on 2026-08-11")

- Opening paragraph: "Each participant with a Microsoft Copilot license in a meeting or in an event with up to 1,000 attendees can ask prompts that are only visible to them. Participants and organizers can learn things like who said what and where people agree or disagree."
- Under **"Prerequisites"**: "An add-on Microsoft Copilot license for intended users."

## 3. Alternative sources

**Meta.** I found no later Meta source that is more direct. The April 2024 post remains the best one: it states the "ask about what you're seeing" feature, gives a general rollout date, and names the regions.

- The Meta newsroom search lists 2025 and 2026 posts about accessibility uses and Meta AI updates for India. None of them is a better primary source for the general feature.
- One later post was verified (via WebFetch; the server refused curl with HTTP 400). It covers continuous video ("live AI"), but only for early-access users. Use it only if the sentence is about continuous, conversational viewing:
  - Meta. (2024, December 16). *Ray-Ban Meta glasses add live AI, live translation, & Shazam support*. https://www.meta.com/blog/ray-ban-meta-v11-software-update-live-ai-translation-shazam/
  - Supports, under "Live AI & Live Translation via Our Early Access Program": "During a live AI session, Meta AI can see what you see continuously and converse with you more naturally than ever before." The page also says this is limited to Early Access Program members in the US and Canada, which makes it weaker for an "on the market" claim. No byline is shown.

**Microsoft.** The support article remains the right source for the prompt. If the review wants the licence or "only visible to them" point backed in the text, add:

- Microsoft. (2026, August 11). *Manage Microsoft Copilot in Teams meetings and events*. Microsoft Learn. https://learn.microsoft.com/en-us/microsoftteams/copilot-teams-transcription
- Supports: "Each participant with a Microsoft Copilot license ... can ask prompts that are only visible to them" and "An add-on Microsoft Copilot license for intended users" (Prerequisites). Two Microsoft entries would then need different years (2026 vs n.d.), which APA already separates. The site name is kept because "Microsoft Learn" is not identical to the author.

## 4. Recommended sentence and entries

### 4.1 Recommended sentence (paragraph 1, section 1)

> The technology for an independent view is already on the market: Ray-Ban Meta smart glasses answer questions about what the wearer is seeing (Meta, 2024), and Copilot in Teams suggests to a meeting participant, in a private chat, the prompt "Where are some holes in [a meeting participant's] argument?" (Microsoft, n.d., "Have a more effective meeting" section).

If a paraphrase is preferred to the quotation:

> ..., and Copilot in Teams offers a meeting participant, in a private chat, a suggested prompt asking where the holes in another participant's argument lie (Microsoft, n.d.).

Either way, delete "no matter the actual context". If the point behind it is that the tool may answer without enough grounding, state it as the author's claim and back it with the evaluation evidence already in v6 section 2 (Danry et al., 2020, 63.5% correct in live use).

### 4.2 Do the two sources belong in this paragraph?

They overlap with v6 section 2, paragraph 1, which already cites both, quotes the Copilot prompt with the same section locator, and puts them in the right context ("Listener-held checking tools already exist as commercial products", followed by the accuracy evidence). In section 1 the products serve only as a bridge: the board has a filtering problem, and tools for an independent view exist. My judgement: keep the detail and the quotation in section 2 only. In section 1, keep a short pointer without the quotation, for example "Tools that could give directors an independent view are already sold (Meta, 2024; Microsoft, n.d.), and section 2 asks what they do to trust." This avoids saying the same thing twice in a 1,300-word review. If the author keeps the full sentence in section 1, shorten section 2's version to avoid repeating the quote.

### 4.3 Recommended entries

Meta. (2024, April 23). *New Ray-Ban | Meta smart glasses styles and Meta AI updates*. https://about.fb.com/news/2024/04/new-ray-ban-meta-smart-glasses-styles-and-meta-ai-updates/

Microsoft. (n.d.). *Catch up on meetings with Microsoft Copilot in Teams*. Retrieved October 20, 2026, from https://support.microsoft.com/en-us/teams/copilot/catch-up-on-meetings-with-microsoft-365-copilot-in-teams

(Keeping "[Support article]" after the title is acceptable, but not needed.)

Optional, only if the licence or private-answer point is cited:

Microsoft. (2026, August 11). *Manage Microsoft Copilot in Teams meetings and events*. Microsoft Learn. https://learn.microsoft.com/en-us/microsoftteams/copilot-teams-transcription

In-text forms: (Meta, 2024); (Microsoft, n.d.); with a quotation from the support page, (Microsoft, n.d., "Have a more effective meeting" section); optional Learn page, (Microsoft, 2026).
