#!/usr/bin/env python3
"""Concatenate section files 01 to 12 into deck source files (with notes, and presentation only).

Keeps each section's header line, slides, on-slide text and speaker notes.
Drops the per-file 'Verification log' and 'Open items' sections (working material).
Re-run after editing any section file:  python3 _build-deck-source.py
"""
import re, glob, datetime, pathlib

here = pathlib.Path(__file__).parent
out = here.parent / "Week 4 - Deck Source.md"
files = sorted(f for f in glob.glob(str(here / "[01]*-*.md")) if re.search(r"/(0[1-9]|1[0-2])-", f))

parts = []
parts.append("# Week 4 deck source: sections 01 to 12\n")
parts.append(f"Generated {datetime.date.today().isoformat()} from the section files in `slides/` by `slides/_build-deck-source.py`. "
             "Edit the section files, not this file, then regenerate. Verification logs and open items live in the section files. "
             "Run sheet, placeholders and figure assets are in `slides/00-Overview.md`.\n")
parts.append("""
## Slide 0: Title

**Time:** 0:30 · **Speaker:** Oliver · **Visual:** Figure 1 (PDF page 1, printed p. 1083), the participant panning with a flat hand

**On the slide**
- **User-Defined Gestures for Surface Computing**
- Jacob O. Wobbrock (University of Washington, Information School), Meredith Ringel Morris and Andrew D. Wilson (Microsoft Research)
- CHI 2009, Boston, pp. 1083 to 1092
- CMPT 863 Human-Computer Interaction · Week 4 · [date]
- Presented by Oliver Stutz

**Speaker notes**
- **Say:** Good morning. This week's paper is Wobbrock, Morris and Wilson, "User-Defined Gestures for Surface Computing," CHI 2009. Wobbrock was at the University of Washington; Morris and Wilson at Microsoft Research, which will matter. About twenty-five minutes of talk in twelve short parts, then half an hour of discussion that you will run as much as I do. Section one puts us in 2008.
- **Transition:** Let's go back to 2008.
""")

for f in files:
    text = pathlib.Path(f).read_text(encoding="utf-8")
    # cut from '## Verification log' to end, but keep nothing after it
    text = re.split(r"\n## Verification log\n", text, maxsplit=1)[0].rstrip() + "\n"
    parts.append("\n---\n\n" + text)

full = "\n".join(parts)
out.write_text(full, encoding="utf-8")
print("wrote", out, "from", len(files), "section files")

# Presentation-only variant: same slides, speaker notes removed.
pres = re.sub(r"\n\*\*Speaker notes\*\*\n.*?(?=\n## |\n---\n|\Z)", "\n", full, flags=re.S)
pres = pres.replace("# Week 4 deck source: sections 01 to 12",
                    "# Week 4 deck source, presentation only: sections 01 to 12", 1)
pres = pres.replace("Verification logs and open items live in the section files.",
                    "Speaker notes, verification logs and open items live in the section files; "
                    "the full version with notes is `Week 4 - Deck Source.md`.", 1)
pres = re.sub(r"\n{3,}", "\n\n", pres)
out2 = here.parent / "Week 4 - Deck Source - Presentation only.md"
out2.write_text(pres, encoding="utf-8")
print("wrote", out2)

# ---------------------------------------------------------------------------
# Question sets 13 to 15: presenter version (full) and class version (questions only)
qfiles = sorted(glob.glob(str(here / "1[3-5]-Questions-*.md")))
qparts = ["# Week 4 question sets, presenter version: files 13 to 15\n",
          f"Generated {datetime.date.today().isoformat()} from `slides/13-Questions-Paper.md`, `14-Questions-Presentation.md` and "
          "`15-Questions-Synthesis.md` by `slides/_build-deck-source.py`. Edit the section files, then regenerate. "
          "Verification logs and open items are left in the section files. The class-facing version with question stems only is "
          "`Week 4 - Questions - Class only.md`.\n"]
cparts = ["# Week 4 question sets, class version: the questions only\n",
          f"Generated {datetime.date.today().isoformat()} by `slides/_build-deck-source.py`. Model answers, discrimination notes, "
          "probes and the class-use plans are in `Week 4 - Questions Source.md` and in the section files.\n"]
for f in qfiles:
    text = pathlib.Path(f).read_text(encoding="utf-8")
    text = re.split(r"\n## Verification log\n", text, maxsplit=1)[0].rstrip() + "\n"
    qparts.append("\n---\n\n" + text)
    title = re.search(r"^# (.+)$", text, re.M).group(1)
    cparts.append(f"\n---\n\n# {title}\n")
    for m in re.finditer(r"^## (Q[\d.]+): (.+?)\n(?:.*?\n)*?- \*\*Form:\*\* ([^\n]+)", text, re.M):
        num, stem, form = m.groups()
        cparts.append(f"\n**{num}** ({form.strip()})\n{stem.strip()}\n")
(here.parent / "Week 4 - Questions Source.md").write_text("\n".join(qparts), encoding="utf-8")
(here.parent / "Week 4 - Questions - Class only.md").write_text("\n".join(cparts), encoding="utf-8")
print("wrote", here.parent / "Week 4 - Questions Source.md", "and", here.parent / "Week 4 - Questions - Class only.md")

# ---------------------------------------------------------------------------
# Discussion guide (file 16): presenter copy without the log, and the pair cards alone
gfile = here / "16-Discussion-Guide.md"
if gfile.exists():
    gtext = gfile.read_text(encoding="utf-8")
    gbody = re.split(r"\n## Verification log\n", gtext, maxsplit=1)[0].rstrip() + "\n"
    ghead = (f"Generated {datetime.date.today().isoformat()} from `slides/16-Discussion-Guide.md` by `slides/_build-deck-source.py`. "
             "Edit the section file, then regenerate. The verification log and open items stay in the section file.\n\n")
    gbody = gbody.replace("\n\n---\n\n## 0.", "\n\n" + ghead + "---\n\n## 0.", 1)
    (here.parent / "Week 4 - Discussion Guide.md").write_text(gbody, encoding="utf-8")
    cards = re.search(r"\n## 4\. Pair cards\n(.*?)\n---\n", gtext, re.S)
    if cards:
        ctext = ("# Week 4 pair cards\n\n" + ghead.replace("Edit the section file", "Edit section 4 of the guide") +
                 cards.group(1).strip() + "\n")
        (here.parent / "Week 4 - Discussion Cards.md").write_text(ctext, encoding="utf-8")
    print("wrote", here.parent / "Week 4 - Discussion Guide.md", "and", here.parent / "Week 4 - Discussion Cards.md")
