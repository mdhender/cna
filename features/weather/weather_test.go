// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package weather_test

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/mdhender/cna/features/steps"
	"github.com/mdhender/cna/internal/board"
	"github.com/mdhender/cna/internal/dice"
	"github.com/mdhender/cna/internal/gametime"
	"github.com/mdhender/cna/internal/weather"
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
	time   gametime.Time
	season gametime.Season
	src    dice.Source
	result weather.Result
}

const (
	seasonName  = `(Spring|Summer|Fall|Winter)`
	weatherKind = `(Normal|Hot|Sandstorm|Rainstorm)`
)

func initializeScenario(sc *godog.ScenarioContext) {
	w := &world{}
	sc.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
		*w = world{}
		return ctx, nil
	})

	steps.Dice(sc, &w.src)
	steps.Time(sc, func(t gametime.Time) {
		w.time, w.season = t, gametime.SeasonOf(t.Turn)
	})
	sc.Step(`^the season is `+seasonName+`$`, w.seasonIs)
	sc.Step(`^the weather is determined$`, w.determine)

	sc.Step(`^the Weather Table reads:$`, tableReads)
	sc.Step(`^every season of the Weather Table gives exactly one weather for each sequential reading$`, everyReadingOnce)
	sc.Step(`^the Foul Weather Location Table reads:$`, locationsRead)
	sc.Step(`^a roll of (\d+) gives `+weatherKind+` weather$`, w.rollGives)
	sc.Step(`^the weather is `+weatherKind+`$`, w.weatherIs)
	sc.Step(`^the weather is `+weatherKind+` on map sections ([A-E](?:,[A-E])*)$`, w.weatherOn)
	sc.Step(`^the weather is `+weatherKind+` in hex ([A-E] ?\d{4})$`, w.weatherInHex)
	sc.Step(`^no more dice are thrown$`, w.noMoreDice)
}

var seasons = []gametime.Season{gametime.Spring, gametime.Summer, gametime.Fall, gametime.Winter}

func parseSeason(s string) gametime.Season {
	for _, season := range seasons {
		if season.String() == s {
			return season
		}
	}
	panic("unknown season " + s)
}

func parseWeather(s string) weather.Weather {
	for _, k := range weather.Kinds {
		if k.String() == s {
			return k
		}
	}
	panic("unknown weather " + s)
}

// seasonIs sets the season when given, and checks it when asserted. A
// season set from a time is the season of that time's Game-Turn.
func (w *world) seasonIs(s string) error {
	want := parseSeason(s)
	if w.time == (gametime.Time{}) {
		w.season = want
		return nil
	}
	if w.season != want {
		return fmt.Errorf("Game-Turn %d: got %s, want %s", w.time.Turn, w.season, want)
	}
	return nil
}

func (w *world) determine() error {
	w.result = weather.Determine(w.time, w.src)
	return nil
}

// formatRange writes a range of numbers the way the charts print them.
func formatRange(lo, hi int, ok bool) string {
	switch {
	case !ok:
		return "-"
	case lo == hi:
		return strconv.Itoa(lo)
	}
	return fmt.Sprintf("%d...%d", lo, hi)
}

// turns lists the blocks of Game-Turns in a season, as "1...12, 49...60".
func turns(s gametime.Season) string {
	var blocks []string
	for turn := gametime.FirstTurn; turn <= gametime.LastTurn; turn++ {
		if gametime.SeasonOf(turn) != s {
			continue
		}
		lo := turn
		for turn < gametime.LastTurn && gametime.SeasonOf(turn+1) == s {
			turn++
		}
		blocks = append(blocks, formatRange(lo, turn, true))
	}
	return strings.Join(blocks, ", ")
}

func tableReads(table *godog.Table) error {
	want := []string{"Season", "Game-Turns", "Normal", "Hot", "Sandstorm", "Rainstorm"}
	if got := cells(table, 0); fmt.Sprint(got) != fmt.Sprint(want) {
		return fmt.Errorf("table header: got %v, want %v", got, want)
	}
	var errs []error
	for i := 1; i < len(table.Rows); i++ {
		row := cells(table, i)
		s := parseSeason(row[0])
		got := []string{turns(s)}
		for _, k := range weather.Kinds {
			got = append(got, formatRange(weather.Range(s, k)))
		}
		for j, cell := range row[1:] {
			if got[j] != cell {
				errs = append(errs, fmt.Errorf("%s %s: chart says %q, engine says %q", s, want[j+1], cell, got[j]))
			}
		}
	}
	return errors.Join(errs...)
}

func cells(table *godog.Table, i int) []string {
	var v []string
	for _, c := range table.Rows[i].Cells {
		v = append(v, c.Value)
	}
	return v
}

func everyReadingOnce() error {
	var errs []error
	for _, s := range seasons {
		for tens := 1; tens <= 6; tens++ {
			for ones := 1; ones <= 6; ones++ {
				roll := tens*10 + ones
				var hits int
				for _, k := range weather.Kinds {
					if lo, hi, ok := weather.Range(s, k); ok && lo <= roll && roll <= hi {
						hits++
					}
				}
				if hits != 1 {
					errs = append(errs, fmt.Errorf("roll %d gives %d results in %s", roll, hits, s))
				}
			}
		}
	}
	return errors.Join(errs...)
}

func locationsRead(table *godog.Table) error {
	dies, sections := cells(table, 0), cells(table, 1)
	if len(dies) != 7 || dies[0] != "Die" || sections[0] != "Map Sections" {
		return fmt.Errorf("table layout: got %v / %v", dies, sections)
	}
	var errs []error
	for i := 1; i <= 6; i++ {
		if dies[i] != strconv.Itoa(i) {
			errs = append(errs, fmt.Errorf("column %d: chart says die %s", i, dies[i]))
			continue
		}
		if got := joinSections(weather.Locate(i)); got != sections[i] {
			errs = append(errs, fmt.Errorf("die %d: chart says %s, engine says %s", i, sections[i], got))
		}
	}
	return errors.Join(errs...)
}

func joinSections(list []board.Section) string {
	var s []string
	for _, sec := range list {
		s = append(s, sec.String())
	}
	return strings.Join(s, ",")
}

func (w *world) rollGives(roll int, want string) error {
	if got := weather.Lookup(w.season, roll); got != parseWeather(want) {
		return fmt.Errorf("%s, roll %d: got %s, want %s", w.season, roll, got, want)
	}
	return nil
}

func (w *world) weatherIs(want string) error {
	if got := w.result.Weather; got != parseWeather(want) {
		return fmt.Errorf("weather: got %s, want %s (roll %d)", got, want, w.result.Throw.Sequential())
	}
	return nil
}

func (w *world) weatherOn(want, sections string) error {
	var errs []error
	for _, s := range strings.Split(sections, ",") {
		sec := board.Section(s[0])
		if got := w.result.On(sec); got != parseWeather(want) {
			errs = append(errs, fmt.Errorf("map section %s: got %s, want %s", sec, got, want))
		}
	}
	return errors.Join(errs...)
}

func (w *world) weatherInHex(want, hex string) error {
	h, err := board.ParseHex(hex)
	if err != nil {
		return err
	}
	if got := w.result.InHex(h); got != parseWeather(want) {
		return fmt.Errorf("hex %s: got %s, want %s", h, got, want)
	}
	return nil
}

func (w *world) noMoreDice() error {
	if s, ok := w.src.(*dice.Script); ok && s.Remaining() > 0 {
		return fmt.Errorf("%d scripted dice were not thrown", s.Remaining())
	}
	return nil
}
