// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package breakdown_test

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/cucumber/godog"
	"github.com/mdhender/cna/features/steps"
	"github.com/mdhender/cna/internal/breakdown"
	"github.com/mdhender/cna/internal/dice"
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
	check  breakdown.Check
	toe    int
	src    dice.Source
	result breakdown.Result
	broken int
}

func initializeScenario(sc *godog.ScenarioContext) {
	w := &world{}
	sc.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
		// Scenarios that don't fix the dice still need some to throw.
		*w = world{src: dice.NewRoller(0, 0)}
		return ctx, nil
	})

	steps.Dice(sc, &w.src)
	sc.Step(`^the unit has accumulated (\d+(?:\.\d+)?) Breakdown Points$`, w.accumulated)
	sc.Step(`^the unit's Breakdown Adjustment Rating is (\w+)$`, w.rating)
	sc.Step(`^the weather is (Normal|Hot)$`, w.weather)
	sc.Step(`^a Sandstorm shifts the column$`, w.sandstorm)
	sc.Step(`^the unit has (\d+) TOE Strength Points?$`, w.hasTOE)

	sc.Step(`^the unit stops moving$`, w.stops)
	sc.Step(`^the unit suffers (\d+)% Breakdown$`, w.suffers)

	sc.Step(`^the Breakdown Table reads:$`, w.tableReads)
	sc.Step(`^every column of the Breakdown Table gives exactly one result for each sequential reading$`, w.everyReadingOnce)
	sc.Step(`^it checks for Breakdown on the (\S+) column$`, w.checksOn)
	sc.Step(`^it does not check for Breakdown$`, w.doesNotCheck)
	sc.Step(`^the dice read (\d+)$`, w.diceRead)
	sc.Step(`^the result is (\d+)% Breakdown$`, w.resultIs)
	sc.Step(`^(\d+) TOE Strength Points? breaks? down$`, w.pointsBreakDown)
}

func (w *world) accumulated(points float64) error {
	w.check.Points = points
	return nil
}

func (w *world) rating(s string) error {
	bar, err := breakdown.ParseBAR(s)
	w.check.BAR = bar
	return err
}

func (w *world) weather(weather string) error {
	w.check.Hot = weather == "Hot"
	return nil
}

func (w *world) sandstorm() error {
	w.check.Sandstorm = true
	return nil
}

func (w *world) hasTOE(toe int) error {
	w.toe = toe
	return nil
}

func (w *world) stops() error {
	w.result = breakdown.Resolve(w.check, w.toe, w.src)
	w.broken = w.result.Broken
	return nil
}

func (w *world) suffers(percent int) error {
	w.broken = breakdown.Broken(w.toe, percent)
	return nil
}

// tableReads compares the engine's table with a copy of the printed
// chart, cell by cell, and reports every cell that differs.
func (w *world) tableReads(table *godog.Table) error {
	if got, want := len(table.Rows[0].Cells)-1, breakdown.NumColumns; got != want {
		return fmt.Errorf("the chart has %d columns, the engine has %d", got, want)
	}
	if got, want := len(table.Rows)-1, len(breakdown.Percents); got != want {
		return fmt.Errorf("the chart has %d rows, the engine has %d", got, want)
	}
	var errs []error
	for c, cell := range table.Rows[0].Cells[1:] {
		if got := breakdown.Column(c).String(); got != cell.Value {
			errs = append(errs, fmt.Errorf("column %d: chart says %q, engine says %q", c+1, cell.Value, got))
		}
	}
	for r, row := range table.Rows[1:] {
		percent := breakdown.Percents[r]
		if got := strconv.Itoa(percent); got != row.Cells[0].Value {
			errs = append(errs, fmt.Errorf("row %d: chart says %s%%, engine says %s%%", r+1, row.Cells[0].Value, got))
			continue
		}
		for c, cell := range row.Cells[1:] {
			col := breakdown.Column(c)
			if got := formatRange(col.Range(percent)); got != cell.Value {
				errs = append(errs, fmt.Errorf("%d%% on the %s column: chart says %q, engine says %q", percent, col, cell.Value, got))
			}
		}
	}
	return errors.Join(errs...)
}

// formatRange writes a range of rolls the way the chart prints it.
func formatRange(lo, hi int, ok bool) string {
	switch {
	case !ok:
		return "-"
	case lo == hi:
		return strconv.Itoa(lo)
	}
	return fmt.Sprintf("%d...%d", lo, hi)
}

func (w *world) everyReadingOnce() error {
	var errs []error
	for c := range breakdown.NumColumns {
		col := breakdown.Column(c)
		for tens := 1; tens <= 6; tens++ {
			for ones := 1; ones <= 6; ones++ {
				roll := tens*10 + ones
				var hits int
				for _, percent := range breakdown.Percents {
					if lo, hi, ok := col.Range(percent); ok && lo <= roll && roll <= hi {
						hits++
					}
				}
				if hits != 1 {
					errs = append(errs, fmt.Errorf("roll %d gives %d results on the %s column", roll, hits, col))
				}
			}
		}
	}
	return errors.Join(errs...)
}

func (w *world) checksOn(column string) error {
	if !w.result.Checked {
		return fmt.Errorf("the unit did not check for Breakdown, want the %s column", column)
	}
	if got := w.result.Column.String(); got != column {
		return fmt.Errorf("column: got %s, want %s", got, column)
	}
	return nil
}

func (w *world) doesNotCheck() error {
	if w.result.Checked {
		return fmt.Errorf("the unit checked for Breakdown on the %s column", w.result.Column)
	}
	return nil
}

func (w *world) diceRead(want int) error {
	if got := w.result.Throw.Sequential(); got != want {
		return fmt.Errorf("dice: got %d, want %d", got, want)
	}
	return nil
}

func (w *world) resultIs(want int) error {
	if !w.result.Checked {
		return fmt.Errorf("the unit did not check for Breakdown, want %d%%", want)
	}
	if w.result.Percent != want {
		return fmt.Errorf("result: got %d%%, want %d%%", w.result.Percent, want)
	}
	return nil
}

func (w *world) pointsBreakDown(want int) error {
	if w.broken != want {
		return fmt.Errorf("TOE Strength Points broken down: got %d, want %d", w.broken, want)
	}
	return nil
}
