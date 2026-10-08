# Rulings

This file records every place where the rules are wrong, contradict each other or the errata, or say nothing, and how the engine resolves it.
Each ruling cites the cases involved and explains the decision in our own words.
Rulings that change engine behavior have a matching `@ruling` scenario in `features/`.

Status is one of:

- **Open**: found, but not yet decided. The engine doesn't implement the affected rule until it's decided.
- **Decided**: the engine follows the stated resolution.

## R-001: Breakdown dice are added or read sequentially

- **Status:** Decided (2026-10-08)
- **Cases:** 21.0 (introduction), 21.31, 21.34, chart 21.38
- **Docs issue:** [mdhender/cnadocs#1](https://github.com/mdhender/cnadocs/issues/1)

The introduction to Section 21.0 says the Breakdown dice are added together.
Cases 21.31 and 21.34 say they are read sequentially (11 to 66), the worked example in 21.34 uses sequential rolls, and every column of the Breakdown Table (21.38) is laid out in sequential ranges.

**Resolution:** read the Breakdown dice sequentially. The detailed cases, the example and the table all agree, and the table can't be used with a sum.
The errata doesn't address the conflict.
The `@ruling` scenario "The Breakdown dice are read sequentially, not added" in `features/breakdown/breakdown_table.feature` pins this down.

## R-002: What decides whether a Sandstorm shifts the Breakdown column

- **Status:** Open
- **Cases:** 21.37d, chart 21.38 (notes)
- **Docs issue:** [mdhender/cnadocs#2](https://github.com/mdhender/cnadocs/issues/2)

Both sources agree that a Sandstorm shifts the Breakdown column one to the right, but they disagree on when it applies.
Case 21.37d measures the unit's movement in Capability Points: the shift applies if at least half of the Capability Points spent in one movement (a Movement Phase, a Retreat, a Reaction and so on) were spent on a map section with Sandstorms.
The notes under the Breakdown Table (21.38) measure Breakdown Points instead: the shift applies if at least half of the accumulated Breakdown Points were picked up in map sections with Sandstorms.
The two can give different answers, for example when a unit moves cheaply along a road in a Sandstorm and then expensively across rough terrain in clear weather.

The Breakdown Table itself doesn't depend on this. `breakdown.Check` takes a yes-or-no `Sandstorm` flag, and the rule that sets it must be decided before movement is implemented.
