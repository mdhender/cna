# CLAUDE.md

This project is the game engine for SPI's **The Campaign for North Africa** (*CNA*), written in Go.

- Module: `github.com/mdhender/cna`, Go 1.26.
- License: AGPLv3 for the Go code and web design. The game's rules, map and counters are © 1979 SPI and are not ours.

## Layout

- The engine is an internal package (`internal/...`). Nothing outside this module imports it.
- `cmd/cna` is the command-line front end.
- `cmd/cnad` is the web server. Its pages use HTMX: the server renders HTML fragments, and there is no client-side framework.
- Both front ends are thin. Game logic goes in the engine, never in a `cmd/` package.
- `features/<area>/` holds the Gherkin feature files for one rules area (`dice`, `combat`, `movement`, ...), with their godog step definitions beside them. See [Behavior-driven development](#behavior-driven-development).
- `internal/dice` rolls dice. Every random outcome goes through its `Source` interface.
- `internal/breakdown` resolves Breakdown checks on the Breakdown Table [21.38].
- `internal/gametime` counts game time: Operations Stages within Game-Turns, written `stage/turn`, and the month-and-week dates the rules sometimes use.
- `internal/board` names the hexes of the game-map (section plus RRCC number, as in C4807) and holds what the engine knows about them, starting with the Delta hexes. Map geometry will come from `github.com/maloquacious/hexg`.
- `internal/weather` determines the weather for an Operations Stage: the Weather Table [29.61] and the Foul Weather Location Table [29.7].
- `internal/terrain` holds the Terrain Effects Chart [8.37] and prices a move into a hex: the Capability Points spent and the Breakdown Points picked up.
- `internal/toe` holds the weapon systems that make up TOE Strength Points, with the Tank and Gun Characteristics Charts [4.47–4.49] as data.
- `features/steps` holds step definitions shared by more than one feature area: the steps that fix the dice, and the step that sets the game time.
- `RULINGS.md` records how we resolve errors, conflicts and gaps in the rules. See [The tests define the game](#the-tests-define-the-game).
- `version.go` holds the version (`cna.Version()`, using `github.com/maloquacious/semver`).

## Commands

```sh
go build ./...
go test ./...                                   # includes every feature file
go test ./features/dice -run 'TestFeatures/Adding_two_dice'   # one scenario
go vet ./...
go mod tidy    # after adding or removing imports
```

## The rules are a reference, not source material

The rulebooks and charts are converted to Markdown in a separate repository at `../docs/`:

- `../docs/land-game-rules.md`: Land Game rules (sections 1.0 to 32.0)
- `../docs/air-logistics-rules.md`: Air and Logistics rules and the scenarios (sections 33.0 to 65.0)
- `../docs/errata.md`: the September 1979 errata. It overrides the rulebooks.
- `../docs/charts-<case>-<name>.md`: one file per chart, named by its case number (for example, `charts-15.79-close-assault-combat-results.md`)

Read them to learn how a rule works.

**Mechanics and chart data are fair game.** Implement any game mechanic. Copy chart and table data, such as combat results, terrain effects, costs and unit characteristics, into Go data structs and functions.

**Rule text is not. Never quote the rules.** This applies to code, comments, tests, commit messages, and any documentation in this repository.

- Cite rules by section or case number, such as `[15.79]` or "case 8.37".
- Summarize or paraphrase in your own words, briefly enough to say what the code does and why.
- Do not copy the prose of the rules, the examples of play, the designer's notes, or the explanatory notes printed with the charts.
- Name the source chart in a comment above each data table, so the values can be checked against it.
- When code implements a rule, put the case number in the comment above it so a reader can look it up:

  ```go
  // resolveCloseAssault applies the Close Assault results table [15.79].
  ```

- When the errata changes a rule, cite both the case and the errata.

## Randomness must be deterministic and repeatable

Every die roll and other random outcome uses `math/rand/v2` with a source that is seeded explicitly.

- Never call the top-level functions in `math/rand/v2` (`rand.IntN`, `rand.Shuffle`, ...). They use a randomly seeded global source, so their results can't be replayed. Never use `math/rand` (v1) or `crypto/rand` for game outcomes either.
- Build generators from an explicit seed, using `rand.NewPCG(seed1, seed2)` or `rand.NewChaCha8(seed)`, and wrap the source with `rand.New`.
- Draw dice through `dice.Source` (`internal/dice`). The engine uses a seeded `dice.Roller`; tests use a `dice.Script` of preset faces. Pass the source into the code that needs it. Don't store it in a package-level variable.
- The same seed and the same sequence of player inputs must produce the same game. Record the seed with the game state so any game can be replayed.
- Keep the order of rolls stable. Ranging over a map gives a random order, so sort the keys first whenever a loop draws random numbers.
- In tests, use fixed seeds and assert on exact outcomes.

## The tests define the game

The feature files are the authority on how the game plays. The rulebooks, charts and errata are evidence for what the features should say, and the engine is whatever makes them pass.

When a feature disagrees with one of the others, record the conflict and open an issue. Don't silently follow one source.

- **The rules, charts or errata conflict with each other, or are wrong or silent:**
  1. Add an entry to `RULINGS.md`: the cases involved, what each says (in our words), and the resolution. Use status Open until the user decides, and Decided after. Rulings are the user's call: propose a resolution, but don't decide one yourself.
  2. Write `@ruling` scenarios that pin the decision down.
  3. Open an issue on the docs repository, giving the file and line of each source, and link it from the ruling:
     ```sh
     gh issue create -R mdhender/cnadocs --assignee @me --label bug ...
     ```
     The docs record the text as printed, so suggest a `<!-- check: -->` comment rather than a change to the text.
- **The engine disagrees with a feature,** and the fix isn't part of the current change: open an issue on this repository (`gh issue create -R mdhender/cna --assignee @me --label bug ...`), tag the failing scenario `@wip`, and add a comment linking the issue.

Issues follow the no-quoting rule too: cite case numbers and line numbers, and describe the text in your own words.

## Behavior-driven development

The feature files are the specification. They're written so players can check them against the rulebooks and errata without reading Go, and godog runs them directly as tests, so what players check is exactly what is tested.

- Write a feature file first, then the step definitions, then the engine code that makes them pass.
- One folder per rules area: `features/<area>/<topic>.feature`, with steps in `features/<area>/<area>_test.go` (package `<area>_test`). Each folder has its own `TestFeatures` runner.
- The runner is strict: a step without a definition fails the build. Scenarios tagged `@wip` are skipped, so unfinished work can be merged without breaking `main`.
- Write steps in plain game language ("the tens die will roll 3"), not in terms of Go types or functions. Reuse an existing step's wording before inventing a new one. When a step is needed in more than one area, move it to `features/steps` and register it from each runner.
- godog matches step text without regard to `Given`, `When` or `Then`, so an action and an assertion need different wording ("the unit suffers 10% Breakdown" versus "the result is 10% Breakdown").
- Group a feature's scenarios with `Rule:`, one rule per idea, so a reader can check each idea against the rulebook on its own.
- Quote names that can contain spaces or punctuation (`the Commonwealth "Heavy AA (3.7")" never breaks down`), and match them with `"(.+)"`, not `"([^"]+)"`, since some names contain quotes.
- Write game times the way the charts do, Operations Stage then Game-Turn: `1/31` is Stage 1 of Game-Turn 31 [4.45].
- Write a scenario for behavior that depends on something not yet built (such as the game calendar) and tag it `@wip`, so the requirement is recorded where players can see it.
- Step definitions are thin. They set up state, call the engine, and compare results. Game logic belongs in `internal/`, never in steps.
- Keep each scenario's state in a `world` struct that is reset before every scenario.
- Fix the dice in rule scenarios with a scripted step such as "the tens die will roll 3 and the ones die will roll 4". Dice steps queue their faces in order, so a scenario can fix several rolls (the Weather Table's two dice, then the Foul Weather Location Table's one). Assert "no more dice are thrown" when a rule must not roll again. Use a seeded roller only for scenarios about the dice themselves.

### Tags

- `@case-<n>` on every scenario that comes from a rule, one tag per case it depends on (for example `@case-15.73`). This is how a reader finds the rule to check it against.
- `@errata` when the errata changes the behavior.
- `@ruling` when the rules are wrong, contradictory or silent, and the scenario pins down our decision. Every `@ruling` scenario has a matching entry in `RULINGS.md`.
- `@engine` for behavior the engine needs that no rule states, such as repeatable dice.
- `@wip` for scenarios that are written but not yet passing.

Feature files are documentation, so the no-quoting rule applies to them. Describe rules in your own words. Chart data may go into `Examples:` tables.

### Map data

Map data is brought in as Go data, a piece at a time, as features need it. Michael Miller's Hex Database (2015) is a useful first draft but is known to have errors, and it stays out of this repository. The official map sections are the authority. A feature that adds map data starts with a scenario holding the data in a form a player can desk-check against the map (see `features/board/delta.feature`), plus scenarios for hexes already checked. A hex on a seam between sections is named by the eastern section's number (B3701, not A3734).

### Charts

A feature that implements a chart starts with a scenario holding a copy of the whole chart as a Gherkin data table, laid out like the printed chart. Its step compares the engine's data with the copy cell by cell and reports every cell that differs. A player can check the copy against the printed chart by eye, and the engine can't drift from it. `features/breakdown/breakdown_table.feature` is the example to follow.

Keep the chart data in Go in the chart's own layout, one call per row with the cells as printed (see `internal/toe/charts.go`), so it can be read against the chart line by line. Apply errata in the data, say so in a comment, and add `@errata` scenarios for each correction.

Then cover how the chart is used: the boundaries of each column and row, the shifts, and what happens at the edges of the table. Use an `@engine` scenario to check the data's integrity, such as every dice reading giving exactly one result.

## Versioning and commits

`version.go` follows semantic versioning.

- A code change always bumps the version. A new feature bumps the minor version. A bugfix, or a change to an existing feature with no external API change, bumps the patch version.
- A documentation-only change doesn't need a bump, but may have one.
- You may commit to `main` and push upstream without asking, as long as any required version bump is made in `version.go` and that file is part of the same commit.

## Conventions

- Every Go file starts with this header:

  ```go
  // Copyright (c) 2026 Michael D Henderson.
  // SPDX-License-Identifier: AGPL-3.0-or-later
  ```
- Follow the Modern Go Guidelines (the `use-modern-go` skill) when writing Go.
