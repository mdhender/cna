# CLAUDE.md

This project is the game engine for SPI's **The Campaign for North Africa** (*CNA*), written in Go.

- Module: `github.com/mdhender/cna`, Go 1.26.
- License: AGPLv3 for the Go code and web design. The game's rules, map and counters are © 1979 SPI and are not ours.

## Layout

- The engine is an internal package (`internal/...`). Nothing outside this module imports it.
- `cmd/cna` is the command-line front end.
- `cmd/cnad` is the web server. Its pages use HTMX: the server renders HTML fragments, and there is no client-side framework.
- Both front ends are thin. Game logic goes in the engine, never in a `cmd/` package.
- `features/` holds work organized by rules area (`combat`, `movement`, ...).
- `version.go` holds the version (`cna.Version()`, using `github.com/maloquacious/semver`).

## Commands

```sh
go build ./...
go test ./...
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
- Pass the `*rand.Rand` (or an interface around it) into the code that needs it. Don't store it in a package-level variable.
- The same seed and the same sequence of player inputs must produce the same game. Record the seed with the game state so any game can be replayed.
- Keep the order of rolls stable. Ranging over a map gives a random order, so sort the keys first whenever a loop draws random numbers.
- In tests, use fixed seeds and assert on exact outcomes.

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
