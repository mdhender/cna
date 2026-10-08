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

## R-003: A dash in the BAR column means the vehicles never break down

- **Status:** Decided (2026-10-08)
- **Cases:** 3.4, 21.11, charts 4.47–4.49 (BAR column and key)

The key to the Tank and Gun Characteristics Charts says a dash means "not applicable or zero", and one of the order of battle spreadsheets records the guns' BAR as 0.
Read as zero, a dash would make towed guns break down without a column shift.

**Resolution:** a dash in the BAR column means the weapon system never breaks down, which is different from a rating of 0.
Case 21.11 lists the only vehicles that break down (trucks, tanks, armored cars and recce, and self-propelled guns), and Case 3.4 says towed artillery isn't subject to Breakdown.
On all three charts, every tank and self-propelled gun has a rating and every towed gun has a dash, so this reading matches the charts exactly.
The `@ruling` scenarios in `features/breakdown/rating.feature` and `features/units/tank_and_gun_characteristics.feature` pin this down.

## R-004: When German tanks stop having a BAR of 1R

- **Status:** Open
- **Cases:** chart 4.49 (notes), 21.36

A note to the German Tank and Gun Characteristics Chart gives all German tanks a BAR of 1R until the start of the "1/31" Game-Turn, after which the printed rating (0) applies.
Case 21.36 says all German tanks have a BAR of 0, and doesn't mention the early rating.
It isn't clear which Game-Turn "1/31" names: Game-Turn 31, or a date such as the turn containing January 31.

The engine uses the printed rating until this is decided. The `@wip` scenario "German tanks have a rating of 1R early in the campaign" waits on this ruling and on the game calendar.
