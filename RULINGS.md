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

- **Status:** Decided (2026-10-08)
- **Cases:** 21.37d, 29.45, 29.41, chart 21.38 (notes)
- **Docs issue:** [mdhender/cnadocs#2](https://github.com/mdhender/cnadocs/issues/2)

All the sources agree that a Sandstorm shifts the Breakdown column one to the right, but they disagree on when it applies.

- Case 21.37d measures the unit's movement in Capability Points: the shift applies if at least half of the Capability Points spent in one movement (a Movement Phase, a Retreat, a Reaction and so on) were spent on a map section with Sandstorms.
- Case 29.45 also measures Capability Points: the shift applies if the unit spends half or more of its Capability Points for movement in a Movement Phase in hexes with a Sandstorm.
- The notes under the Breakdown Table (21.38) measure Breakdown Points instead: the shift applies if at least half of the accumulated Breakdown Points were picked up in map sections with Sandstorms.

The two measures can give different answers, for example when a unit moves cheaply along a road in a Sandstorm and then expensively across rough terrain in clear weather.
The two cases also differ slightly on where: 21.37d counts movement on a map section with Sandstorms, while 29.45 counts movement in Sandstorm hexes. They differ in practice only where a Sandstorm doesn't reach every hex of a struck section: it never covers the Delta hexes of map section E [29.41, 29.7], and stops at the coastal hexes [29.46].

**Resolution:** measure Breakdown Points, as the notes under the Breakdown Table do.
A Sandstorm shifts the Breakdown column one to the right if at least half of the Breakdown Points the unit has accumulated this Operations Stage were picked up in map sections with a Sandstorm.
This overrides the Capability Point measure in Cases 21.37d and 29.45.

`breakdown.SandstormShift` decides the shift from the points accumulated and the points picked up in a Sandstorm; its answer goes in `breakdown.Check.Sandstorm`.
The `@ruling` scenario "Breakdown Points decide whether a Sandstorm shifts the column" in `features/breakdown/breakdown_table.feature` pins this down, including a unit that picks up most of its points outside the Sandstorm.

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

- **Status:** Decided (2026-10-08)
- **Cases:** 29.1, chart 29.61, errata 29.1 and 29.61
- **Docs issues:** [mdhender/cnadocs#5](https://github.com/mdhender/cnadocs/issues/5) and [#6](https://github.com/mdhender/cnadocs/issues/6), both closed. The docs flag the conflicts with check comments (cnadocs commit `3949e92`).

The seasons are named by month and week (the dating the errata calls abandoned), and three sources say when they start:

- The table in Case 29.1 starts Spring with March III and ends it with June II, and the other seasons follow the same pattern (Summer from June III, Fall from September III, Winter from December III).
- The errata to 29.1 agrees: Spring runs from the 3rd week of March through the 2nd week of June.
- The errata to the Weather Table (29.61) says Spring runs from the 4th week of March through the 3rd week of June.

With four Game-Turns to a month and Game-Turn 1 in September III, 1940, the 29.1 table gives every season exactly 12 Game-Turns, starting with Fall at Game-Turn 1. The Weather Table's Game-Turn columns come in blocks of 12 (1–12, 13–24 and so on), so they match 29.1's starts, not the 29.61 errata's.

The Weather Table also prints each Game-Turn block against the opposite season, which the errata acknowledges: the block 1–12 is printed on the Spring row but is Fall, and the Summer row's blocks are Winter. Its last block stops at Game-Turn 110, though the campaign runs to 111 (January I, 1943, in Winter).
The weather results on each row fit the season named on that row (Summer is mostly Hot; Winter has no Hot and the most rain), so only the Game-Turn blocks are misplaced.

**Resolution:** follow the 29.1 table and its errata. Fall is Game-Turns 1–12, 49–60 and 97–108; Winter is 13–24, 61–72 and 109–111; Spring is 25–36 and 73–84; Summer is 37–48 and 85–96. The weather results in each row stay with the season named on that row.
The calendar (`features/calendar/game_time.feature`) already checks that each season boundary in the 29.1 table falls 12 Game-Turns after the last. Weather isn't implemented yet; when it is, its feature starts with this table as its full-chart scenario, with `@ruling` scenarios for the season boundaries and Game-Turn 111.

The Weather Table with the resolution applied:

| Season | Starts | Game-Turns | | | Normal | Hot | Sandstorm | Rainstorm |
|---|---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| Fall | September III | 1...12 | 49...60 | 97...108 | 11...35 | 36...54 | 55...61 | 62...66 |
| Winter | December III | 13...24 | 61...72 | 109...111 | 11...52 | - | - | 53...66 |
| Spring | March III | 25...36 | 73...84 | - | 11...42 | 43...55 | 56...64 | 65...66 |
| Summer | June III | 37...48 | 85...96 | - | 11...23 | 24...55 | 56...66 | - |

A Sandstorm or Rainstorm result leads to the Foul Weather Location Table. The Weather Table's footnote calls it Table 29.62, but no such table exists: it is 29.7 (cnadocs#6). The Weather Table is 29.61, as numbered by the charts booklet and the errata; the docs now point the rules' references to it there (cnadocs#9). The engine cites 29.61 and 29.7.

## R-006: The terrain of a hex that is partly water

- **Status:** Decided (2026-10-08)
- **Cases:** 8.37

Some hexes on the map are partly water and partly land. The rules and the Terrain Effects Chart give each hex one terrain type, but don't say how to type a hex that shows more than one, or one that is mostly water.

On map section E, hexes E3717 and E3718 look alike: mostly water, with a Delta marker reaching into the hex. The engine's map data (first transcribed from Michael Miller's Hex Database) types E3718 as Delta but E3717 as Clear, with sea hexsides. E0819 and E0820 are partly water too: a lake runs through both, with Delta land in E0819 and Rough in E0820, and the data types them Delta and Rough.

**Resolution:** a hex that is partly water takes the terrain of its land. A hex showing a Delta marker is a Delta hex, so E3717 becomes Delta like E3718. The water is shown by the hex's sea hexsides, not by its terrain.

The engine's map data now has E3717 as Delta. The `@ruling` scenario "A partly water hex with a Delta marker is a Delta hex" in `features/board/delta.feature` pins this down.

## R-007: What it costs to move along a track

- **Status:** Open
- **Cases:** 8.33, 8.37 (Track row and note 8), 8.42, 8.46, errata 8.37
- **Docs issue:** [mdhender/cnadocs#11](https://github.com/mdhender/cnadocs/issues/11)

The sources give three different costs for moving along a track:

- The Terrain Effects Chart prints a cost of 1 CP for a track, and its note 8 says a track halves the costs of terrain features and the Breakdown values of hexes and hexsides, except for a vehicle going down an escarpment.
- The errata to 8.37 says the printed 1 is wrong and note 8 is right: a track halves the cost of the terrain it runs through.
- Case 8.46 says a track costs 1 CP per hex, halves most hexside costs, and halves the Breakdown cost of the hex. Case 8.33 says a unit on a track ignores the other terrain in the hex and hexside, except for vehicles crossing escarpments.

**Proposed resolution:** follow the errata and note 8, overriding Cases 8.33 and 8.46.
Along a track, a unit spends half the CP of the hex's terrain and half the CP of each hexside feature, and a vehicle picks up half their Breakdown Points.
A vehicle going down an escarpment on a track pays the escarpment in full: 8 CP and 6 Breakdown Points on top of the halved hex [8.42].
Vehicles still never go up an escarpment, even on a track [8.42].

Tracks do cross escarpments on the map, so this matters in play. Tracks cross escarpments at seven hexsides in Michael Miller's hex database, and all seven have been checked against the official map. Going down the escarpment, the track runs:

- from A4133 to A4132
- from A4432 to A4431
- from B3901 (A3934 on sheet A) to A3933
- from C4020 to C4121
- from D1308 to D1208
- from D1815 to D1716
- from D2429 to D2330

Miller's data has the direction right for all but C4020 to C4121, which it marks as down from both sides.

Until this is decided, `terrain.Enter` refuses moves along a track.
The `@ruling @wip` scenarios under "A track halves the cost of the terrain it runs through" in `features/terrain/terrain_effects_chart.feature` hold the proposal.

## R-008: Where vehicles may go on a road

- **Status:** Open
- **Cases:** 8.37 (notes 2, 6, 9 and 11), 8.42, 8.44
- **Docs issue:** [mdhender/cnadocs#12](https://github.com/mdhender/cnadocs/issues/12)

A road cancels the costs of every hexside it crosses (note 6), and lets motorized units cross a Major River (note 11). But the sources don't agree on whether it opens two other kinds of terrain to vehicles:

- **Escarpments.** The chart prohibits motorized units from going up an escarpment, and lets them go down one only on a track (note 9). Case 8.42 says no vehicle ever goes up an escarpment, and vehicles go down only along a track. Neither mentions roads.
  No road on the map has yet been found crossing an escarpment, so this part may never arise in play. The road from C4020 into Sollum (C4021) crosses a slope, only a track crosses at Halfaya Pass (over a slope too), and the one road-and-escarpment hexside in Miller's hex database (C4120's northeast side) is an error: the road there runs beside the escarpment.
- **Salt Marsh.** Note 2 lets most vehicles into or out of a Salt Marsh only on a track. Case 8.44 says on a road or track.

**Proposed resolution:** a road opens both to vehicles.
A unit moving along a road crosses an escarpment either way at the road's cost, as note 6 says for every hexside; the prohibitions in 8.42 and note 9 apply to units off the road.
Vehicles may enter a Salt Marsh along a road, as Case 8.44 says.

Until this is decided, `terrain.Enter` refuses vehicles along a road into a Salt Marsh and any unit along a road across an escarpment.
The `@ruling @wip` scenarios "Vehicles use a road across an escarpment" and "Vehicles use a road into a Salt Marsh" in `features/terrain/terrain_effects_chart.feature` hold the proposal.
