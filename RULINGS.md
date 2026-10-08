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
- **Docs issue:** [mdhender/cnadocs#3](https://github.com/mdhender/cnadocs/issues/3)

The key to the Tank and Gun Characteristics Charts says a dash means "not applicable or zero", and one of the order of battle spreadsheets records the guns' BAR as 0.
Read as zero, a dash would make towed guns break down without a column shift.

**Resolution:** a dash in the BAR column means the weapon system never breaks down, which is different from a rating of 0.
Case 21.11 lists the only vehicles that break down (trucks, tanks, armored cars and recce, and self-propelled guns), and Case 3.4 says towed artillery isn't subject to Breakdown.
On all three charts, every tank and self-propelled gun has a rating and every towed gun has a dash, so this reading matches the charts exactly.
The `@ruling` scenarios in `features/breakdown/rating.feature` and `features/units/tank_and_gun_characteristics.feature` pin this down.

## R-004: When German tanks stop having a BAR of 1R

- **Status:** Decided (2026-10-08)
- **Cases:** chart 4.49 (notes), 4.45, 20.11, 20.55, 43.12, 21.36, errata 27.16 and 29.1
- **Docs issue:** [mdhender/cnadocs#4](https://github.com/mdhender/cnadocs/issues/4)

A note to the German Tank and Gun Characteristics Chart gives all German tanks a BAR of 1R until the start of the "1/31" Game-Turn, after which the printed rating (0) applies.
It wasn't clear which turn "1/31" names, and Case 21.36 says all German tanks have a BAR of 0 without mentioning the early rating.

**Resolution:** "1/31" is Stage 1 of Game-Turn 31. German tanks have a BAR of 1R through Stage 3 of Game-Turn 30, and the printed rating from Stage 1 of Game-Turn 31.

- Case 4.45 defines the notation: the OA Charts give a unit's arrival as the Operations Stage, a slash, then the Game-Turn.
- Cases 20.11, 20.55 and 43.12 use the same notation, and the stage number is never more than 3, the number of Operations Stages in a Game-Turn.
- Case 21.36's BAR of 0 is the printed rating that applies from Game-Turn 31 on, so it doesn't conflict.
- Case 43.13 writes "June I/35", where the "I" is a week of the month, not a stage. The errata to 27.16 and 29.1 says that month-and-week dating was abandoned, which supports reading "1/31" in the newer stage/Game-Turn notation. Either reading starts the printed rating at the beginning of Game-Turn 31.

The `@ruling` scenario "German tanks have a rating of 1R until Game-Turn 31" in `features/units/tank_and_gun_characteristics.feature` pins this down.

## R-005: When the seasons start

- **Status:** Open
- **Cases:** 29.1, chart 29.61, errata 29.1 and 29.61
- **Docs issue:** [mdhender/cnadocs#5](https://github.com/mdhender/cnadocs/issues/5)

The seasons are named by month and week (the dating the errata calls abandoned), and three sources say when they start:

- The table in Case 29.1 starts Spring with March III and ends it with June II, and the other seasons follow the same pattern (Summer from June III, Fall from September III, Winter from December III).
- The errata to 29.1 agrees: Spring runs from the 3rd week of March through the 2nd week of June.
- The errata to the Weather Table (29.61) says Spring runs from the 4th week of March through the 3rd week of June.

With four Game-Turns to a month and Game-Turn 1 in September III, 1940, the 29.1 table gives every season exactly 12 Game-Turns, starting with Fall at Game-Turn 1. The Weather Table's Game-Turn columns come in blocks of 12 (1–12, 13–24 and so on), so they match 29.1's starts, not the 29.61 errata's.

The Weather Table also lists its Game-Turn blocks against the wrong seasons, which the errata acknowledges: the block 1–12 is printed on the Spring row but is Fall. Its last block stops at Game-Turn 110, though the campaign runs to 111.

**Proposed resolution:** follow the 29.1 table and its errata. Fall is Game-Turns 1–12, 49–60 and 97–108; Winter is 13–24, 61–72 and 109–111; Spring is 25–36 and 73–84; Summer is 37–48 and 85–96. The weather results in each row stay with the season named on that row.
The calendar (`features/calendar/game_time.feature`) already checks that each season boundary in the 29.1 table falls 12 Game-Turns after the last. Weather isn't implemented yet.
