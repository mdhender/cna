// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package gametime counts time in the game. Each Game-Turn is about a
// week, divided into three Operations Stages; the Operations Stage is the
// basic unit of time [5.1]. The campaign runs from Stage 1 of Game-Turn 1
// to Stage 3 of Game-Turn 111 [64.2].
//
// Times are written as on the charts: the Operations Stage, a slash, then
// the Game-Turn, so 1/31 is Stage 1 of Game-Turn 31 [4.45].
package gametime

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	StagesPerTurn = 3
	FirstTurn     = 1
	LastTurn      = 111
)

// Time is one Operations Stage of the campaign.
type Time struct {
	Stage int // 1 to StagesPerTurn
	Turn  int // FirstTurn to LastTurn
}

// Start and End are the first and last Operations Stages of the campaign.
var (
	Start = Time{Stage: 1, Turn: FirstTurn}
	End   = Time{Stage: StagesPerTurn, Turn: LastTurn}
)

// New returns Stage stage of Game-Turn turn, or an error if that is not a
// time in the campaign.
func New(stage, turn int) (Time, error) {
	t := Time{Stage: stage, Turn: turn}
	if stage < 1 || stage > StagesPerTurn || turn < FirstTurn || turn > LastTurn {
		return Time{}, fmt.Errorf("gametime: %s is not a time in the campaign", t)
	}
	return t, nil
}

// Parse reads a time written stage/Game-Turn, such as "1/31".
func Parse(s string) (Time, error) {
	stage, turn, ok := strings.Cut(s, "/")
	st, err1 := strconv.Atoi(stage)
	tu, err2 := strconv.Atoi(turn)
	if !ok || err1 != nil || err2 != nil {
		return Time{}, fmt.Errorf("gametime: %q is not written stage/Game-Turn", s)
	}
	return New(st, tu)
}

// MustParse is like Parse but panics on error. It is for times written
// into the engine's data.
func MustParse(s string) Time {
	t, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return t
}

// String returns the time written stage/Game-Turn.
func (t Time) String() string {
	return fmt.Sprintf("%d/%d", t.Stage, t.Turn)
}

// index counts Operations Stages from the start of the campaign.
func (t Time) index() int {
	return (t.Turn-FirstTurn)*StagesPerTurn + t.Stage - 1
}

// Compare returns -1, 0 or +1 as t is before, the same as, or after u.
func (t Time) Compare(u Time) int {
	return cmp.Compare(t.index(), u.index())
}

// Before reports whether t is earlier than u.
func (t Time) Before(u Time) bool { return t.Compare(u) < 0 }

// After reports whether t is later than u.
func (t Time) After(u Time) bool { return t.Compare(u) > 0 }

// Next returns the Operations Stage after t, and false if t is the last
// stage of the campaign.
func (t Time) Next() (Time, bool) {
	if t == End {
		return t, false
	}
	if t.Stage < StagesPerTurn {
		return Time{Stage: t.Stage + 1, Turn: t.Turn}, true
	}
	return Time{Stage: 1, Turn: t.Turn + 1}, true
}

// Date is a Game-Turn named by its week of the month, as in "September
// III, 1940". The rules give each month four Game-Turns, weeks I to IV
// (errata 27.16, 29.1); the charts number Game-Turns instead.
type Date struct {
	Year  int
	Month time.Month
	Week  int // 1 to 4
}

// first is the date of Game-Turn 1 [6.14, 64.2].
var first = Date{Year: 1940, Month: time.September, Week: 3}

const weeksPerMonth = 4

// weeks counts Game-Turn weeks from January of year 0.
func (d Date) weeks() int {
	return (d.Year*12+int(d.Month)-1)*weeksPerMonth + d.Week - 1
}

// DateOf returns the date of a Game-Turn.
func DateOf(turn int) Date {
	w := first.weeks() + turn - FirstTurn
	m := w / weeksPerMonth
	return Date{Year: m / 12, Month: time.Month(m%12 + 1), Week: w%weeksPerMonth + 1}
}

// TurnOf returns the Game-Turn of a date, and false if the date falls
// outside the campaign.
func TurnOf(d Date) (int, bool) {
	turn := d.weeks() - first.weeks() + FirstTurn
	return turn, d.Week >= 1 && d.Week <= weeksPerMonth && turn >= FirstTurn && turn <= LastTurn
}

var weekNumerals = []string{"I", "II", "III", "IV"}

// ParseDate reads a date such as "September III, 1940".
func ParseDate(s string) (Date, error) {
	monthWeek, year, ok := strings.Cut(s, ", ")
	month, week, ok2 := strings.Cut(monthWeek, " ")
	y, err := strconv.Atoi(year)
	w := slices.Index(weekNumerals, week)
	if !ok || !ok2 || err != nil || w < 0 {
		return Date{}, fmt.Errorf("gametime: %q is not a date such as \"September III, 1940\"", s)
	}
	for m := time.January; m <= time.December; m++ {
		if m.String() == month {
			return Date{Year: y, Month: m, Week: w + 1}, nil
		}
	}
	return Date{}, fmt.Errorf("gametime: %q has no month", s)
}

// String returns the date as the rules write it, such as "September III, 1940".
func (d Date) String() string {
	return fmt.Sprintf("%s %s, %d", d.Month, weekNumerals[d.Week-1], d.Year)
}

// Season is a season of the year, which picks the row of the Weather
// Table [29.1].
type Season int

const (
	Spring Season = iota
	Summer
	Fall
	Winter
)

func (s Season) String() string {
	return [...]string{"Spring", "Summer", "Fall", "Winter"}[s]
}

// seasonStarts gives the week of the year each season starts, counting
// January I as 0: Spring from March III, Summer from June III, Fall from
// September III and Winter from December III [29.1] (ruling R-005).
var seasonStarts = [...]int{
	Spring: 2*weeksPerMonth + 2,
	Summer: 5*weeksPerMonth + 2,
	Fall:   8*weeksPerMonth + 2,
	Winter: 11*weeksPerMonth + 2,
}

// SeasonOf returns the season of a Game-Turn.
func SeasonOf(turn int) Season {
	d := DateOf(turn)
	week := (int(d.Month)-1)*weeksPerMonth + d.Week - 1
	season := Winter // from December III through March II
	for s := Spring; s <= Winter; s++ {
		if week >= seasonStarts[s] {
			season = s
		}
	}
	return season
}
