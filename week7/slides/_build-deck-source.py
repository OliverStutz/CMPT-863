#!/usr/bin/env python3
"""Concatenate section files 01 to 11 into deck source files (with notes, and presentation only).

Keeps each section's header line, slides, on-slide text and speaker notes.
Drops the per-file 'Verification log' and 'Open items' sections (working material).
Also emits the discussion guide (file 12) without its log, the pair cards alone, and,
if files 13 to 15 exist, the question sets in presenter and class versions.
Re-run after editing any section file:  python3 slides/_build-deck-source.py  (from week7/)
Adapted 2026-10-19 from Presentation-Week4/slides/_build-deck-source.py.
"""
import re, glob, datetime, pathlib

here = pathlib.Path(__file__).parent
out = here.parent / "Week 7 - Deck Source.md"
files = sorted(f for f in glob.glob(str(here / "[01]*-*.md")) if re.search(r"/(0[1-9]|1[01])-", f))

parts = []
parts.append("# Week 7 deck source: sections 01 to 11\n")
parts.append(f"Generated {datetime.date.today().isoformat()} from the section files in `slides/` by `slides/_build-deck-source.py`. "
             "Edit the section files, not this file, then regenerate. Verification logs and open items live in the section files. "
             "Run sheet, placeholders and figure assets are in `slides/00-Overview.md`.\n")
parts.append("""
## Slide 0: Title

**Time:** 0:30 · **Speaker:** Oliver · **Visual:** the title question large; beneath it a strip cropped from Figure 1 (p. 945, `assets/p9-09.png`), the Rams products with their annotations

**On the slide**
- **What Should We Expect From Research Through Design?**
- William Gaver, Interaction Research Studio, Goldsmiths, University of London
- CHI 2012, Austin, Texas, session Critical Perspectives on Design, pp. 937 to 946
- CMPT 863 Human-Computer Interaction · Week 7 · 22 October 2026
- Presented by Oliver Stutz

**Speaker notes**
- **Say:** Good morning. This week's paper is an essay, not a study: William Gaver, "What Should We Expect From Research Through Design?", CHI 2012. Gaver runs a design research studio at Goldsmiths in London, and he is answering colleagues who wanted research through design to adopt standards. About twenty-five minutes from me in eleven short parts, then half an hour that you will run as much as I do. We start with who is arguing with whom.
- **Transition:** First, the term, the author, and the argument he walked into.
""")

for f in files:
    text = pathlib.Path(f).read_text(encoding="utf-8")
    text = re.split(r"\n## Verification log\n", text, maxsplit=1)[0].rstrip() + "\n"
    parts.append("\n---\n\n" + text)

full = "\n".join(parts)
out.write_text(full, encoding="utf-8")
print("wrote", out, "from", len(files), "section files")

# Presentation-only variant: same slides, speaker notes removed.
pres = re.sub(r"\n\*\*Speaker notes\*\*\n.*?(?=\n## |\n---\n|\Z)", "\n", full, flags=re.S)
pres = pres.replace("# Week 7 deck source: sections 01 to 11",
                    "# Week 7 deck source, presentation only: sections 01 to 11", 1)
pres = pres.replace("Verification logs and open items live in the section files.",
                    "Speaker notes, verification logs and open items live in the section files; "
                    "the full version with notes is `Week 7 - Deck Source.md`.", 1)
# Drop the "Alternates for Oliver to swap in" and "Background for the presenter" sections (presenter material) from the presentation-only copy.
pres = re.sub(r"\n## (Alternates for Oliver to swap in|Background for the presenter)\n.*?(?=\n## |\n---\n|\Z)", "\n", pres, flags=re.S)
pres = re.sub(r"\n{3,}", "\n\n", pres)
out2 = here.parent / "Week 7 - Deck Source - Presentation only.md"
out2.write_text(pres, encoding="utf-8")
print("wrote", out2)

# ---------------------------------------------------------------------------
# Question sets 13 to 15 (optional): presenter version (full) and class version (questions only)
qfiles = sorted(glob.glob(str(here / "1[3-5]-Questions-*.md")))
if qfiles:
    qparts = ["# Week 7 question sets, presenter version\n",
              f"Generated {datetime.date.today().isoformat()} from the question-set files in `slides/` by `slides/_build-deck-source.py`. "
              "Edit the section files, then regenerate. Verification logs and open items are left in the section files. "
              "The class-facing version with question stems only is `Week 7 - Questions - Class only.md`.\n"]
    cparts = ["# Week 7 question sets, class version: the questions only\n",
              f"Generated {datetime.date.today().isoformat()} by `slides/_build-deck-source.py`. Model answers, discrimination notes, "
              "probes and the class-use plans are in `Week 7 - Questions Source.md` and in the section files.\n"]
    for f in qfiles:
        text = pathlib.Path(f).read_text(encoding="utf-8")
        text = re.split(r"\n## Verification log\n", text, maxsplit=1)[0].rstrip() + "\n"
        qparts.append("\n---\n\n" + text)
        title = re.search(r"^# (.+)$", text, re.M).group(1)
        cparts.append(f"\n---\n\n# {title}\n")
        for m in re.finditer(r"^## (Q[\d.]+): (.+?)\n(?:.*?\n)*?- \*\*Form:\*\* ([^\n]+)", text, re.M):
            num, stem, form = m.groups()
            cparts.append(f"\n**{num}** ({form.strip()})\n{stem.strip()}\n")
    (here.parent / "Week 7 - Questions Source.md").write_text("\n".join(qparts), encoding="utf-8")
    (here.parent / "Week 7 - Questions - Class only.md").write_text("\n".join(cparts), encoding="utf-8")
    print("wrote", here.parent / "Week 7 - Questions Source.md", "and", here.parent / "Week 7 - Questions - Class only.md")

# ---------------------------------------------------------------------------
# Discussion guide (file 12): presenter copy without the log, and the pair cards alone
gfile = here / "12-Discussion-Guide.md"
if gfile.exists():
    gtext = gfile.read_text(encoding="utf-8")
    gbody = re.split(r"\n## Verification log\n", gtext, maxsplit=1)[0].rstrip() + "\n"
    ghead = (f"Generated {datetime.date.today().isoformat()} from `slides/12-Discussion-Guide.md` by `slides/_build-deck-source.py`. "
             "Edit the section file, then regenerate. The verification log and open items stay in the section file.\n\n")
    gbody = gbody.replace("\n\n---\n\n## 0.", "\n\n" + ghead + "---\n\n## 0.", 1)
    (here.parent / "Week 7 - Discussion Guide.md").write_text(gbody, encoding="utf-8")
    cards = re.search(r"\n## 4\. Pair cards\n(.*?)\n---\n", gtext, re.S)
    if cards:
        ctext = ("# Week 7 pair cards\n\n" + ghead.replace("Edit the section file", "Edit section 4 of the guide") +
                 cards.group(1).strip() + "\n")
        (here.parent / "Week 7 - Discussion Cards.md").write_text(ctext, encoding="utf-8")
    print("wrote", here.parent / "Week 7 - Discussion Guide.md", "and", here.parent / "Week 7 - Discussion Cards.md")
