import { onCleanup, onMount } from "solid-js";
import cytoscape from "cytoscape";

// Just a fixed demo graph -- this route exists only to try out
// Cytoscape.js, not to visualize any real data.
const ELEMENTS = [
  { data: { id: "a", label: "A" } },
  { data: { id: "b", label: "B" } },
  { data: { id: "c", label: "C" } },
  { data: { id: "d", label: "D" } },
  { data: { source: "a", target: "b" } },
  { data: { source: "b", target: "c" } },
  { data: { source: "c", target: "d" } },
  { data: { source: "d", target: "a" } },
];

export default function Graph() {
  let containerRef: HTMLDivElement | undefined;

  // cytoscape needs the container element mounted in the DOM before it
  // can measure/render into it, so setup happens in onMount rather than
  // inline during render.
  onMount(() => {
    if (!containerRef) return;

    const cy = cytoscape({
      container: containerRef,
      elements: ELEMENTS,
      style: [
        {
          selector: "node",
          style: {
            "background-color": "#4ac26b",
            label: "data(label)",
            color: "#000000",
            "text-valign": "center",
            "text-halign": "center",
          },
        },
        {
          selector: "edge",
          style: {
            width: 2,
            "line-color": "#d0d7de",
            "target-arrow-color": "#d0d7de",
            "target-arrow-shape": "triangle",
            "curve-style": "bezier",
          },
        },
      ],
      layout: { name: "circle" },
    });

    // Guards against the container having zero width/height at init
    // time (e.g. if the parent flex layout hasn't reflowed yet), which
    // would otherwise render an empty canvas.
    cy.resize();
    cy.fit();

    onCleanup(() => cy.destroy());
  });

  return (
    <div class="flex w-full flex-col gap-4 xl:mx-auto xl:max-w-3xl">
      <h1 class="mb-4 font-sans text-4xl">Graph</h1>
      <div
        ref={containerRef}
        class="h-[500px] w-full rounded-md border border-border bg-field"
      />
    </div>
  );
}
