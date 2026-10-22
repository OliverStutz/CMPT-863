#!/usr/bin/env python3
"""Concatenate slides2 section files 01 to 09 into deck source files (with notes, and presentation only).

Keeps each section's header line, slides, on-slide text and speaker notes.
Drops the per-file 'Verification log' and 'Open items' sections (working material).
The presentation-only copy also drops the 'If asked' and 'Alternates' blocks.
Re-run after editing any section file:  python3 slides2/_build-deck-source.py  (from week7/)
Adapted 2026-10-21 from slides/_build-deck-source.py. The discussion guide is built by the slides/ script.
"""
import re, glob, datetime, pathlib

here = pathlib.Path(__file__).parent
out = here.parent / "Week 7 - Deck Source v2.md"
files = sorted(f for f in glob.glob(str(here / "0*-*.md")) if re.search(r"/0[1-9]-", f))

parts = []
parts.append("# Week 7 deck source, version 2: sections 01 to 09\n")
parts.append(f"Generated {datetime.date.today().isoformat()} from the section files in `slides2/` by `slides2/_build-deck-source.py`. "
             "Edit the section files, not this file, then regenerate. Verification logs and open items live in the section files. "
             "Run sheet, placeholders and the rationale for version 2 are in `slides2/00-Overview.md`.\n")
parts.append("""
## Slide 0: Title

**Time:** 0:30 · **Speaker:** Oliver · **Visual:** the title question large; beneath it a strip cropped from Figure 1 (p. 945, `assets/rams-fig1-annotated.png`), the Rams products with their annotations

**On the slide**
- **What Should We Expect From Research Through Design?**
- William Gaver, Interaction Research Studio, Goldsmiths, University of London
- CHI 2012, Austin, Texas, session Critical Perspectives on Design, pp. 937 to 946
- CMPT 863 Human-Computer Interaction · Week 7 · 22 October 2026
- Presented by Oliver Stutz

**Speaker notes**
- **Say:** Good morning. This week's paper is an essay, not a study. William Gaver runs a design research studio at Goldsmiths in London, and in 2012 he answered colleagues who wanted research through design to adopt standards. I will take about twenty-five minutes, with three moments where I need your hands, and then the half hour is yours. We start with an object.
- **Transition:** Before the paper, a coffee table.
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
pres = pres.replace("# Week 7 deck source, version 2: sections 01 to 09",
                    "# Week 7 deck source, version 2, presentation only: sections 01 to 09", 1)
pres = pres.replace("Verification logs and open items live in the section files.",
                    "Speaker notes, verification logs and open items live in the section files; "
                    "the full version with notes is `Week 7 - Deck Source v2.md`.", 1)
pres = re.sub(r"\n{3,}", "\n\n", pres)
out2 = here.parent / "Week 7 - Deck Source v2 - Presentation only.md"
out2.write_text(pres, encoding="utf-8")
print("wrote", out2)
