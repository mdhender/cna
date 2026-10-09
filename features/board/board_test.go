// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package board_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/cucumber/godog"
	"github.com/mdhender/cna/internal/board"
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

func initializeScenario(sc *godog.ScenarioContext) {
	sc.Step(`^the Delta hexes of map section E are:$`, deltaHexesAre)
	sc.Step(`^hex ([A-E] ?\d{4}) (is|is not) a Delta hex$`, isDelta)
}

// deltaHexesAre compares the engine's Delta hexes with the table, row by
// row, and reports every row that differs.
func deltaHexesAre(table *godog.Table) error {
	var errs []error
	listed := map[int]bool{}
	for _, row := range table.Rows[1:] {
		r, err := strconv.Atoi(row.Cells[0].Value)
		if err != nil {
			return fmt.Errorf("bad row %q", row.Cells[0].Value)
		}
		listed[r] = true
		if got, want := board.DeltaColumns(r), row.Cells[1].Value; got != want {
			errs = append(errs, fmt.Errorf("row %d: map says %q, engine says %q", r, want, got))
		}
	}
	for _, r := range board.DeltaRows() {
		if !listed[r] {
			errs = append(errs, fmt.Errorf("row %d: engine has Delta hexes %q the map doesn't list", r, board.DeltaColumns(r)))
		}
	}
	return errors.Join(errs...)
}

func isDelta(hex, is string) error {
	h, err := board.ParseHex(hex)
	if err != nil {
		return err
	}
	if want := is == "is"; board.IsDelta(h) != want {
		return fmt.Errorf("hex %s: IsDelta is %v, want %v", h, board.IsDelta(h), want)
	}
	return nil
}
