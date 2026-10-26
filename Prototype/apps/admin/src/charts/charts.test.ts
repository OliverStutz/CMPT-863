import { describe, expect, it } from "vitest";
import { labelPlacement } from "./BarChart";
import { hoverProps } from "./common";
import { OTHER_ID, foldSlices } from "./Donut";
import { dumbbellDomain } from "./Dumbbell";
import { bandPolygons, orderedPoints, segments } from "./LineChart";
import { barPath, niceTicks } from "./scale";
import { stackSegments } from "./StackedBar";

const xs = [{ id: "T0" }, { id: "T1" }, { id: "T2a" }, { id: "T2b" }];

describe("line chart transforms", () => {
  it("orders points by the x axis, not by the object's key order", () => {
    const pts = orderedPoints({ points: { T2b: { y: 4 }, T0: { y: 6, lo: 5, hi: 7, n: 22 }, T1: { y: 5 } } }, xs);
    expect(pts.map((p) => p.x)).toEqual(["T0", "T1", "T2a", "T2b"]);
    expect(pts.map((p) => p.y)).toEqual([6, 5, null, 4]);
    expect(pts[0]).toMatchObject({ index: 0, lo: 5, hi: 7, n: 22 });
  });

  it("breaks the line at a missing point", () => {
    const pts = orderedPoints({ points: { T0: { y: 6 }, T1: { y: 5 }, T2b: { y: 4 } } }, xs);
    expect(segments(pts).map((s) => s.map((p) => p.x))).toEqual([["T0", "T1"], ["T2b"]]);
  });

  it("builds the CI band as upper bounds left to right, then lower bounds right to left", () => {
    const pts = orderedPoints(
      {
        points: {
          T0: { y: 6, lo: 5, hi: 7 },
          T1: { y: 5, lo: 4, hi: 6 },
          T2a: { y: 3, lo: 2, hi: 4 },
          T2b: { y: 4, lo: null, hi: null },
        },
      },
      xs,
    );
    const polys = bandPolygons(pts);
    expect(polys).toHaveLength(1);
    expect(polys[0]).toEqual([
      [0, 7],
      [1, 6],
      [2, 4],
      [2, 2],
      [1, 4],
      [0, 5],
    ]);
  });

  it("splits the band where a CI is missing", () => {
    const pts = orderedPoints(
      {
        points: { T0: { y: 6, lo: 5, hi: 7 }, T1: { y: 5 }, T2a: { y: 3, lo: 2, hi: 4 }, T2b: { y: 4, lo: 3, hi: 5 } },
      },
      xs,
    );
    expect(bandPolygons(pts).map((p) => p.length)).toEqual([2, 4]);
  });
});

describe("donut folding", () => {
  const data = [
    { id: "CA", label: "Canada", value: 10 },
    { id: "DE", label: "Germany", value: 8 },
    { id: "US", label: "United States", value: 6 },
    { id: "GB", label: "United Kingdom", value: 5 },
    { id: "CH", label: "Switzerland", value: 4 },
    { id: "IN", label: "India", value: 4 },
    { id: "BR", label: "Brazil", value: 2 },
    { id: "NL", label: "Netherlands", value: 2 },
    { id: "FR", label: "France", value: 1 },
    { id: "XX", label: "Nowhere", value: 0 },
  ];

  it("keeps the six largest and folds the rest into Other", () => {
    const f = foldSlices(data, 6);
    expect(f).toHaveLength(7);
    expect(f.slice(0, 6).map((s) => s.id)).toEqual(["CA", "DE", "US", "GB", "IN", "CH"]);
    expect(f[6]).toMatchObject({ id: OTHER_ID, label: "Other", value: 5 });
    expect(f[6].members?.map((m) => m.id)).toEqual(["BR", "NL", "FR"]);
    expect(f.reduce((a, s) => a + s.value, 0)).toBe(42);
  });

  it("has no Other when six or fewer slices remain", () => {
    expect(foldSlices(data.slice(0, 6), 6).some((s) => s.id === OTHER_ID)).toBe(false);
    expect(foldSlices([{ id: "a", label: "a", value: 0 }], 6)).toEqual([]);
  });
});

describe("bars and stacks", () => {
  it("places a value label inside, outside or only in the tooltip", () => {
    expect(labelPlacement(100, 20, 50, true)).toBe("inside");
    expect(labelPlacement(100, 20, 50, false)).toBe("outside");
    expect(labelPlacement(10, 20, 50, true)).toBe("outside");
    expect(labelPlacement(10, 20, 10, true)).toBe("tooltip");
  });

  it("leaves a 2 px gap between stacked segments and skips zeros", () => {
    const keys = [
      { key: "a", label: "A", color: "x" },
      { key: "b", label: "B", color: "x" },
      { key: "c", label: "C", color: "x" },
    ];
    const segs = stackSegments({ id: "r", label: "r", values: { a: 1, b: 0, c: 3 } }, keys, 400);
    expect(segs.map((s) => s.key)).toEqual(["a", "c"]);
    expect(segs[0]).toMatchObject({ x0: 0, x1: 100, share: 0.25 });
    expect(segs[1]).toMatchObject({ x0: 102, x1: 400, share: 0.75 });
  });

  it("draws a bar with a rounded data end and a square baseline", () => {
    expect(barPath("horizontal", 0, 0, 50, 16)).toBe("M0,0H46Q50,0 50,4V12Q50,16 46,16H0Z");
    expect(barPath("vertical", 0, 10, 50, 16)).toBe("M0,60V14Q0,10 4,10H12Q16,10 16,14V60Z");
  });

  it("covers every CI and zero in the dumbbell domain", () => {
    const d = dumbbellDomain(
      [
        {
          id: "r",
          label: "r",
          points: { a: { mean: -1.5, lo: -2.2, hi: -0.8, n: 20 }, b: { mean: 0.4, lo: 0.1, hi: 0.7, n: 20 } },
        },
      ],
      true,
    );
    expect(d[0]).toBeLessThanOrEqual(-2.2);
    expect(d[1]).toBeGreaterThanOrEqual(0.7);
  });

  it("makes round ticks", () => {
    expect(niceTicks(0, 10, 5)).toEqual([0, 2, 4, 6, 8, 10]);
    expect(niceTicks(1, 7, 6)).toEqual([1, 2, 3, 4, 5, 6, 7]);
  });
});

describe("tooltip initial state", () => {
  it("opens only on a real pointer move or on focus, and never for an empty mark", () => {
    let open = 0;
    let closed = 0;
    const h = hoverProps(
      () => open++,
      () => closed++,
    ) as Required<ReturnType<typeof hoverProps>>;
    expect("onPointerEnter" in h).toBe(false);
    h.onPointerMove({ movementX: 0, movementY: 0, isTrusted: true }); // a layout shift under a resting pointer
    h.onPointerMove({ movementX: 3, movementY: 0, isTrusted: false });
    expect(open).toBe(0);
    h.onPointerMove({ movementX: 3, movementY: 1, isTrusted: true });
    h.onFocus();
    expect(open).toBe(2);
    h.onBlur();
    expect(closed).toBe(1);
    expect(
      hoverProps(
        () => open++,
        () => closed++,
        false,
      ),
    ).toEqual({});
  });
});
