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
A `@ruling` scenario will pin this down when Breakdown is implemented.
