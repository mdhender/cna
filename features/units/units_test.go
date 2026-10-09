// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package units_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/mdhender/cna/features/steps"
	"github.com/mdhender/cna/internal/gametime"
	"github.com/mdhender/cna/internal/toe"
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

const nationality = `(Commonwealth|Italian|German)`

// world holds the state of one scenario.
type world struct {
	// time is the game time, or nil to use the ratings printed on the charts.
	time *gametime.Time
}

func initializeScenario(sc *godog.ScenarioContext) {
	w := &world{}
	sc.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
		*w = world{}
		return ctx, nil
	})

	steps.Time(sc, func(t gametime.Time) { w.time = &t })
	sc.Step(`^the `+nationality+` (tank|artillery|anti-tank|anti-air) weapon systems read:$`, systemsRead)
	sc.Step(`^the `+nationality+` "(.+)" reads:$`, systemReads)
	sc.Step(`^the `+nationality+` "(.+)" has a CPA of (\S+)$`, hasCPA)
	sc.Step(`^the `+nationality+` "(.+)" has an Armor Protection Rating of (\S+)$`, hasArmorProtection)
	sc.Step(`^the `+nationality+` "(.+)" breaks down with a rating of (\S+)$`, w.breaksDownWith)
	sc.Step(`^the `+nationality+` "(.+)" has an Anti-Armor Rating of (\S+)$`, w.hasAntiArmor)
	sc.Step(`^the `+nationality+` "(.+)" never breaks down$`, neverBreaksDown)
	sc.Step(`^every tank and self-propelled gun has a Breakdown Adjustment Rating$`, everyVehicleHasRating)
	sc.Step(`^no other weapon system has one$`, noOtherHasRating)
}

// header is the column layout of the Tank and Gun Characteristics Charts.
var header = []string{"CPA", "AA", "Barrage", "Anti-Armor", "Vul", "Armor Prtctn", "Close Assault", "Fuel Rate", "BAR"}

// cells renders a weapon system's ratings the way the chart prints them.
func cells(s toe.System) []string {
	return []string{
		s.CPA.String(), s.AntiAir.String(), s.Barrage.String(), s.AntiArmor.String(),
		s.Vulnerability.String(), s.ArmorProtection.String(),
		s.CloseAssaultOff.String() + "/" + s.CloseAssaultDef.String(),
		s.FuelRate.String(), s.Breakdown.String(),
	}
}

func parseNationality(s string) toe.Nationality {
	for _, n := range []toe.Nationality{toe.Commonwealth, toe.Italian, toe.German} {
		if n.String() == s {
			return n
		}
	}
	panic("unknown nationality " + s)
}

func parseClass(s string) toe.Class {
	for _, c := range []toe.Class{toe.Tank, toe.Artillery, toe.AntiTank, toe.AntiAir} {
		if c.String() == s {
			return c
		}
	}
	panic("unknown class " + s)
}

func lookup(nat, name string) (toe.System, error) {
	s, ok := toe.Lookup(parseNationality(nat), name)
	if !ok {
		return s, fmt.Errorf("no %s weapon system named %q", nat, name)
	}
	return s, nil
}

// compareRow checks one chart row against a weapon system, cell by cell.
func compareRow(s toe.System, row []string) []error {
	var errs []error
	for i, got := range cells(s) {
		if got != row[i] {
			errs = append(errs, fmt.Errorf("%s %s: chart says %q, engine says %q", s.Name, header[i], row[i], got))
		}
	}
	return errs
}

// values returns the cells of row i of a table as strings.
func values(table *godog.Table, i int) []string {
	var v []string
	for _, c := range table.Rows[i].Cells {
		v = append(v, c.Value)
	}
	return v
}

// checkHeader checks a table's header row, ignoring its first skip cells.
func checkHeader(table *godog.Table, skip int) error {
	got := values(table, 0)[skip:]
	if fmt.Sprint(got) != fmt.Sprint(header) {
		return fmt.Errorf("table header: got %v, want %v", got, header)
	}
	return nil
}

func systemsRead(nat, class string, table *godog.Table) error {
	if err := checkHeader(table, 1); err != nil {
		return err
	}
	systems := toe.Systems(parseNationality(nat), parseClass(class))
	rows := len(table.Rows) - 1
	var errs []error
	for i := range rows {
		v := values(table, i+1)
		if i >= len(systems) {
			errs = append(errs, fmt.Errorf("the chart lists %s, the engine has no more %s %s weapon systems", v[0], nat, class))
			continue
		}
		if systems[i].Name != v[0] {
			errs = append(errs, fmt.Errorf("row %d: chart says %q, engine says %q", i+1, v[0], systems[i].Name))
			continue
		}
		errs = append(errs, compareRow(systems[i], v[1:])...)
	}
	for _, s := range systems[min(rows, len(systems)):] {
		errs = append(errs, fmt.Errorf("the engine has %s, which the chart doesn't list", s.Name))
	}
	return errors.Join(errs...)
}

func systemReads(nat, name string, table *godog.Table) error {
	if err := checkHeader(table, 0); err != nil {
		return err
	}
	s, err := lookup(nat, name)
	if err != nil {
		return err
	}
	return errors.Join(compareRow(s, values(table, 1))...)
}

func hasCPA(nat, name, want string) error {
	s, err := lookup(nat, name)
	if err != nil {
		return err
	}
	if got := s.CPA.String(); got != want {
		return fmt.Errorf("%s CPA: got %s, want %s", name, got, want)
	}
	return nil
}

func hasArmorProtection(nat, name, want string) error {
	s, err := lookup(nat, name)
	if err != nil {
		return err
	}
	if got := s.ArmorProtection.String(); got != want {
		return fmt.Errorf("%s Armor Protection: got %s, want %s", name, got, want)
	}
	return nil
}

func (w *world) breaksDownWith(nat, name, want string) error {
	s, err := lookup(nat, name)
	if err != nil {
		return err
	}
	if !s.SubjectToBreakdown() {
		return fmt.Errorf("%s is not subject to Breakdown", name)
	}
	rating := s.Breakdown
	if w.time != nil {
		rating = s.BreakdownAt(*w.time)
	}
	bar, ok := rating.BAR()
	if !ok {
		return fmt.Errorf("%s has no Breakdown Adjustment Rating", name)
	}
	if got := bar.String(); got != want {
		return fmt.Errorf("%s rating: got %s, want %s", name, got, want)
	}
	return nil
}

func (w *world) hasAntiArmor(nat, name, want string) error {
	s, err := lookup(nat, name)
	if err != nil {
		return err
	}
	rating := s.AntiArmor
	if w.time != nil {
		rating = s.AntiArmorAt(*w.time)
	}
	if got := rating.String(); got != want {
		return fmt.Errorf("%s Anti-Armor: got %s, want %s", name, got, want)
	}
	return nil
}

func neverBreaksDown(nat, name string) error {
	s, err := lookup(nat, name)
	if err != nil {
		return err
	}
	if s.SubjectToBreakdown() {
		return fmt.Errorf("%s is subject to Breakdown", name)
	}
	if _, ok := s.Breakdown.BAR(); ok {
		return fmt.Errorf("%s has a Breakdown Adjustment Rating of %s", name, s.Breakdown)
	}
	return nil
}

func everyVehicleHasRating() error {
	var errs []error
	for _, s := range toe.All() {
		if _, ok := s.Breakdown.BAR(); s.SubjectToBreakdown() && !ok {
			errs = append(errs, fmt.Errorf("%s %s breaks down but has no rating", s.Nationality, s.Name))
		}
	}
	return errors.Join(errs...)
}

func noOtherHasRating() error {
	var errs []error
	for _, s := range toe.All() {
		if _, ok := s.Breakdown.BAR(); !s.SubjectToBreakdown() && ok {
			errs = append(errs, fmt.Errorf("%s %s never breaks down but has a rating of %s", s.Nationality, s.Name, s.Breakdown))
		}
	}
	return errors.Join(errs...)
}
