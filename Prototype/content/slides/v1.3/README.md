# Slides v1.3: rendered images

Source: `Prototype/03 - Study Content v1.3.md`, section 5 (the five slides; the brief called it section 4) and section 1 (scenario).

## Files

| Slide | SVG | PNG |
|---|---|---|
| 1 Where we stand today | `s1.svg` | `s1.png` |
| 2 The risk and what it would cost us | `s2.svg` | `s2.png` |
| 3 How we will protect our 13 sites | `s3.svg` | `s3.png` |
| 4 What the $2.4 million pays for | `s4.svg` | `s4.png` |
| 5 The decision today | `s5.svg` | `s5.png` |

`make_slides.py` regenerates all five SVGs: `python3 make_slides.py <output folder>`. It needs Pillow and the Nimbus Sans and Noto Sans font files, which it uses only to measure text for line wrapping.

Every SVG is 1920 by 1080, self-contained, with no external fonts or images. Each text segment from the layout tables is rendered word for word, and each rendered element carries the segment id as its `id` (for example `id="s2.num3"`). Multi-element items are wrapped in a `<g id="...">`. The diagram parts on slide 3 carry sub-ids under `s3.img1` (`s3.img1.links`, `s3.img1.hq`, `s3.img1.wh1` to `s3.img1.wh12`, `s3.img1.internet`, `s3.img1.remote`). The frame elements that are not segments have the ids `brand` (company name) and `slide-number`.

## Render command

```
for N in 1 2 3 4 5; do
  rsvg-convert -w 1920 -h 1080 -o Prototype/content/slides/v1.3/s$N.png Prototype/content/slides/v1.3/s$N.svg
done
```

Run from the repository root. Rendered with `/usr/bin/rsvg-convert`.

## Font

The SVGs ask for `Inter, Helvetica, Arial, sans-serif`. Inter is not installed on this machine, so fontconfig resolves Helvetica to Nimbus Sans (a Helvetica metric clone), and the PNGs show Nimbus Sans. On a machine with Inter installed the SVGs render in Inter; Inter is slightly wider, and line wrapping was computed against the wider of Nimbus Sans and Noto Sans to leave room. Re-check the PNGs if the SVGs are rendered elsewhere.

## Design

White background, a dark navy title bar (#14325A) with an amber accent line (#D9822B), "Meridian Logistics" small at the top right with an amber square as a logo mark, a thin rule above the footer, and the slide number at the bottom right.

## Interpretations

- **Number segments.** Each `number` segment has the form "figure: label". The figure, including the colon, is set large and bold, and the label follows in regular weight (slides 1 and 2), or the label is the row text and the amount sits in the right column (slide 4, with the colon kept after the item). The text of each segment is unchanged; only the size and position differ within it. On slide 5, "$2.4 million for approval today" has no colon, so it is one line with "$2.4 million" in bold.
- **Slide 4.** "Shown as a simple cost table, one row per line, total in bold" is drawn as a borderless table with alternating row shading, amounts right-aligned, and the total row bold on a tinted background below a heavier rule. No header row was added, so no text appears that is not in the layout table.
- **Slide 3 diagram (`s3.img1`).** Head office is a dark box in the centre labelled "Head office" and "Main data centre". The 12 warehouses sit on an ellipse around it, each a small box with a factory roof icon and the label "Warehouse 1" to "Warehouse 12" (the numbering is my addition; the description only asks for warehouse icons). Each warehouse is joined to head office by a solid line. A shield icon (path with a tick) sits on the corner of every warehouse box and on head office, 13 in all. One solid line runs up from head office to a cloud labelled "Internet". A laptop labelled "Remote staff" is joined to head office by a dotted line from below. The diagram labels are not layout segments and are not part of any segment text.
- **Charts.** None of the five slides has a `chart` segment, so no chart was drawn.
- **Wrapping.** Long bullets wrap onto two lines; no text was abbreviated or left out.
