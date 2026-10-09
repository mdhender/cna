// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package weather determines the weather for an Operations Stage [29.0].
//
// Two dice, read sequentially, are rolled on the season's row of the
// Weather Table [29.61]. A Sandstorm or Rainstorm strikes only some map
// sections: one die on the Foul Weather Location Table [29.7] picks them,
// and the rest have Normal weather [29.1]. Hot weather covers every map
// section [29.31].
package weather

import (
	"fmt"
	"slices"

	"github.com/mdhender/cna/internal/board"
	"github.com/mdhender/cna/internal/dice"
	"github.com/mdhender/cna/internal/gametime"
)

// Weather is the weather on a map section for one Operations Stage.
type Weather int

const (
	Normal Weather = iota
	Hot
	Sandstorm
	Rainstorm
)

// Kinds lists the weather in the order of the Weather Table's columns.
var Kinds = [...]Weather{Normal, Hot, Sandstorm, Rainstorm}

func (w Weather) String() string {
	return [...]string{"Normal", "Hot", "Sandstorm", "Rainstorm"}[w]
}

// Foul reports whether the weather strikes only the map sections given
// by the Foul Weather Location Table.
func (w Weather) Foul() bool {
	return w == Sandstorm || w == Rainstorm
}

// rollRange is a range of sequential dice readings. The zero value means
// the season has no entry for that weather.
type rollRange struct{ lo, hi int }

// table is the Weather Table [29.61], as corrected by the errata and
// ruling R-005: each row keeps the weather results printed on it, and the
// Game-Turns come from the season's dates in 29.1 (gametime.SeasonOf).
// The columns are Normal, Hot, Sandstorm and Rainstorm.
var table = [...][len(Kinds)]rollRange{
	gametime.Fall:   {{11, 35}, {36, 54}, {55, 61}, {62, 66}},
	gametime.Winter: {{11, 52}, {}, {}, {53, 66}},
	gametime.Spring: {{11, 42}, {43, 55}, {56, 64}, {65, 66}},
	gametime.Summer: {{11, 23}, {24, 55}, {56, 66}, {}},
}

// Range returns the range of rolls that gives w in season s, or false if
// the season has no entry for w.
func Range(s gametime.Season, w Weather) (lo, hi int, ok bool) {
	r := table[s][w]
	return r.lo, r.hi, r != rollRange{}
}

// Lookup returns the weather for a sequential dice reading in season s.
// It panics if the reading is not on the table.
func Lookup(s gametime.Season, roll int) Weather {
	for _, w := range Kinds {
		if r := table[s][w]; r.lo <= roll && roll <= r.hi {
			return w
		}
	}
	panic(fmt.Sprintf("weather: roll %d is not on the %s row", roll, s))
}

// locations is the Foul Weather Location Table [29.7]: the map sections
// struck by a Sandstorm or Rainstorm for each roll of one die.
var locations = [6][]board.Section{
	{'A', 'B'},
	{'C', 'D'},
	{'D', 'E'},
	{'B', 'C'},
	{'B', 'D'},
	{'B', 'C', 'D'},
}

// Locate returns the map sections struck by foul weather for a roll of
// one die.
func Locate(die int) []board.Section {
	return slices.Clone(locations[die-1])
}

// Result is the weather determined for one Operations Stage.
type Result struct {
	Time    gametime.Time
	Season  gametime.Season
	Throw   dice.Throw // the two dice rolled on the Weather Table
	Weather Weather    // the weather the table gives
	// Die is the die rolled on the Foul Weather Location Table, and
	// Struck the map sections it gives. Both are zero unless the weather
	// is foul.
	Die    int
	Struck []board.Section
}

// Determine rolls the weather for time t [29.1]. It throws two dice, and
// one more die only if the weather is foul.
func Determine(t gametime.Time, src dice.Source) Result {
	r := Result{Time: t, Season: gametime.SeasonOf(t.Turn), Throw: dice.Two(src)}
	r.Weather = Lookup(r.Season, r.Throw.Sequential())
	if r.Weather.Foul() {
		r.Die = dice.One(src)
		r.Struck = Locate(r.Die)
	}
	return r
}

// On returns the weather on map section s. Foul weather strikes only the
// sections it was located on; the others have Normal weather [29.1].
func (r Result) On(s board.Section) Weather {
	if r.Weather.Foul() && !slices.Contains(r.Struck, s) {
		return Normal
	}
	return r.Weather
}

// InHex returns the weather in hex h. It is the weather on the hex's map
// section, except that a Sandstorm never reaches a Delta hex [29.41, 29.7].
func (r Result) InHex(h board.Hex) Weather {
	w := r.On(h.Section)
	if w == Sandstorm && board.IsDelta(h) {
		return Normal
	}
	return w
}
