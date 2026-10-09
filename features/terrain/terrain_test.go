// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package terrain_test

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/mdhender/cna/internal/terrain"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: initializeScenario,
		Options: &godog.Options{
			Format:   "progress",
			Paths:    []string{"."},
			Strict:   true,
			Tags:     "~@wip",
			TestingT: t,
		},
	}
	if suite.Run() != 0 {
		t.Fatal("feature tests failed")
	}
}

// world holds the state of one scenario.
type world struct {
	mover terrain.Mover
	move  terrain.Move
	err   error
}

// points matches a number of points as the chart writes them: 2, ½ or 2½.
const points = `(\d+½?|½)`

func initializeScenario(sc *godog.ScenarioContext) {
	w := &world{}
	sc.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
		*w = world{}
		return ctx, nil
	})

	sc.Step(`^the unit is (non-motorized|motorized|Light Trucks|motorcycle infantry|motorcycle recce|recce)$`, w.unitIs)
	sc.Step(`^it moves (?:(across country|along a road|along a track) )?into an? (.+?) hex(?: across (.+))?$`, w.moves)

	sc.Step(`^it moves through these hexes:$`, w.movesThrough)

	sc.Step(`^the Terrain Effects Chart reads:$`, w.chartReads)
	sc.Step(`^note (\d+) is on the "(.+)" row only$`, w.noteOnRow)
	sc.Step(`^the "(.+)" row gives no CP cost of its own$`, w.noCPCost)
	sc.Step(`^every type of terrain has exactly one row on the Terrain Effects Chart$`, w.oneRowEach)
	sc.Step(`^it spends `+points+` Capability Points?$`, w.spends)
	sc.Step(`^it picks up `+points+` Breakdown Points?$`, w.picksUp)
	sc.Step(`^it spends `+points+` Capability Points? and picks up `+points+` Breakdown Points?$`, w.spendsAndPicksUp)
	sc.Step(`^it may not move there$`, w.mayNotMove)
}

var movers = map[string]terrain.Mover{
	"non-motorized":       terrain.NonMotorized,
	"motorized":           terrain.Motorized,
	"Light Trucks":        terrain.LightTrucks,
	"motorcycle infantry": terrain.MotorcycleInfantry,
	"motorcycle recce":    terrain.MotorcycleRecce,
	"recce":               terrain.Recce,
}

func (w *world) unitIs(kind string) error {
	w.mover = movers[kind]
	return nil
}

var routes = map[string]terrain.Route{
	"":               terrain.Across,
	"across country": terrain.Across,
	"along a road":   terrain.AlongRoad,
	"along a track":  terrain.AlongTrack,
}

// movesThrough reads a move of several hexes from a table with the
// columns route, terrain and hexsides, and asks the engine for its cost.
func (w *world) movesThrough(table *godog.Table) error {
	var path []terrain.Step
	for _, row := range table.Rows[1:] {
		route, ok := routes[row.Cells[0].Value]
		if !ok {
			return fmt.Errorf("unknown route %q", row.Cells[0].Value)
		}
		hex, err := terrain.Parse(row.Cells[1].Value)
		if err != nil {
			return err
		}
		sides, err := parseHexsides(row.Cells[2].Value)
		if err != nil {
			return err
		}
		path = append(path, terrain.Step{Hex: hex, Hexsides: sides, Route: route})
	}
	w.move, w.err = terrain.Path(w.mover, path)
	return nil
}

// moves reads the hex's terrain and a list of hexside features written
// "a Wadi and a Minor River", and asks the engine for the cost.
func (w *world) moves(route, hex, across string) error {
	t, err := terrain.Parse(hex)
	if err != nil {
		return err
	}
	sides, err := parseHexsides(across)
	if err != nil {
		return err
	}
	w.move, w.err = terrain.Enter(w.mover, t, sides, routes[route])
	return nil
}

// parseHexsides reads a list of hexside features written "a Wadi and a
// Minor River". An empty list means no features.
func parseHexsides(s string) ([]terrain.Terrain, error) {
	if s == "" {
		return nil, nil
	}
	var sides []terrain.Terrain
	for name := range strings.SplitSeq(s, " and ") {
		name = strings.TrimPrefix(strings.TrimPrefix(name, "an "), "a ")
		side, err := terrain.Parse(name)
		if err != nil {
			return nil, err
		}
		sides = append(sides, side)
	}
	return sides, nil
}

// parsePoints reads a number of points written 2, ½ or 2½.
func parsePoints(s string) float64 {
	whole, half := strings.CutSuffix(s, "½")
	var n float64
	if whole != "" {
		i, _ := strconv.Atoi(whole)
		n = float64(i)
	}
	if half {
		n += 0.5
	}
	return n
}

// formatPoints writes a number of points the way the chart does.
func formatPoints(n float64) string {
	whole := int(n)
	switch {
	case n == float64(whole):
		return strconv.Itoa(whole)
	case n == float64(whole)+0.5 && whole == 0:
		return "½"
	case n == float64(whole)+0.5:
		return strconv.Itoa(whole) + "½"
	}
	return strconv.FormatFloat(n, 'f', -1, 64)
}

func (w *world) spends(want string) error {
	if w.err != nil {
		return fmt.Errorf("the move failed: %w", w.err)
	}
	if got := formatPoints(w.move.CP); got != formatPoints(parsePoints(want)) {
		return fmt.Errorf("Capability Points: got %s, want %s", got, want)
	}
	return nil
}

func (w *world) picksUp(want string) error {
	if w.err != nil {
		return fmt.Errorf("the move failed: %w", w.err)
	}
	if got := formatPoints(w.move.Breakdown); got != formatPoints(parsePoints(want)) {
		return fmt.Errorf("Breakdown Points: got %s, want %s", got, want)
	}
	return nil
}

func (w *world) spendsAndPicksUp(cp, bp string) error {
	return errors.Join(w.spends(cp), w.picksUp(bp))
}

func (w *world) mayNotMove() error {
	switch {
	case w.err == nil:
		return fmt.Errorf("the move was allowed, spending %s CP", formatPoints(w.move.CP))
	case !errors.Is(w.err, terrain.ErrProhibited):
		return fmt.Errorf("the move failed for another reason: %w", w.err)
	}
	return nil
}

// chartReads compares the engine's chart with a copy of the printed chart,
// cell by cell, and reports every cell that differs.
func (w *world) chartReads(table *godog.Table) error {
	if got, want := len(table.Rows)-1, len(terrain.Chart); got != want {
		return fmt.Errorf("the chart has %d rows, the engine has %d", got, want)
	}
	var errs []error
	for c, cell := range table.Rows[0].Cells {
		if c >= len(terrain.Headings) {
			errs = append(errs, fmt.Errorf("the chart has a column %q the engine lacks", cell.Value))
		} else if got := terrain.Headings[c]; got != cell.Value {
			errs = append(errs, fmt.Errorf("column %d: chart says %q, engine says %q", c+1, cell.Value, got))
		}
	}
	for r, row := range table.Rows[1:] {
		engine := terrain.Chart[r].Cells
		if len(row.Cells) != len(engine) {
			errs = append(errs, fmt.Errorf("%s: the chart has %d cells, the engine has %d", row.Cells[0].Value, len(row.Cells), len(engine)))
			continue
		}
		for c, cell := range row.Cells {
			if engine[c] != cell.Value {
				errs = append(errs, fmt.Errorf("%s, %s: chart says %q, engine says %q", row.Cells[0].Value, terrain.Headings[c], cell.Value, engine[c]))
			}
		}
	}
	return errors.Join(errs...)
}

// noteOnRow checks that a footnote marker appears in the named row and in
// no other.
func (w *world) noteOnRow(note int, name string) error {
	marker := fmt.Sprintf("(%d)", note)
	var errs []error
	found := false
	for _, r := range terrain.Chart {
		has := false
		for _, cell := range r.Cells {
			if strings.Contains(cell, marker) {
				has = true
			}
		}
		rowName, _, _ := strings.Cut(r.Cells[0], " (")
		switch {
		case rowName == name:
			found = has
		case has:
			errs = append(errs, fmt.Errorf("note %d is also on the %s row", note, rowName))
		}
	}
	if !found {
		errs = append(errs, fmt.Errorf("note %d is not on the %s row", note, name))
	}
	return errors.Join(errs...)
}

// noCPCost checks that both CP cells of a row hold only a footnote marker.
func (w *world) noCPCost(name string) error {
	for _, r := range terrain.Chart {
		if r.Terrain.String() != name {
			continue
		}
		var errs []error
		for c := 1; c <= 2; c++ {
			if !strings.HasPrefix(r.Cells[c], "(") {
				errs = append(errs, fmt.Errorf("%s, %s: the engine gives a cost of %q", name, terrain.Headings[c], r.Cells[c]))
			}
		}
		return errors.Join(errs...)
	}
	return fmt.Errorf("no %s row on the chart", name)
}

func (w *world) oneRowEach() error {
	count := map[terrain.Terrain]int{}
	for _, r := range terrain.Chart {
		if r.Terrain != 0 {
			count[r.Terrain]++
		}
	}
	var errs []error
	for _, t := range terrain.All {
		if count[t] != 1 {
			errs = append(errs, fmt.Errorf("%s has %d rows", t, count[t]))
		}
	}
	return errors.Join(errs...)
}
