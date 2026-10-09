// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package board names the hexes of the game-map and holds what the engine
// knows about them.
//
// The map has five sections, A to E from west to east, each numbered on
// its own. A hex is named by its section and a four-digit number: two
// digits of row, then two of column, as in C4807 (Tobruk).
package board

import (
	"fmt"
	"strconv"
	"strings"
)

// Section is a section of the game-map, A to E.
type Section byte

// Sections lists every map section, west to east.
var Sections = [...]Section{'A', 'B', 'C', 'D', 'E'}

func (s Section) String() string {
	return string(s)
}

// Hex is one hex of the game-map.
type Hex struct {
	Section Section
	Row     int
	Col     int
}

// ParseHex reads a hex written as on the map, such as "E3117" or "E 3117".
func ParseHex(s string) (Hex, error) {
	s = strings.ReplaceAll(s, " ", "")
	if len(s) != 5 || s[0] < 'A' || s[0] > 'E' {
		return Hex{}, fmt.Errorf("board: %q is not a hex such as E3117", s)
	}
	row, err1 := strconv.Atoi(s[1:3])
	col, err2 := strconv.Atoi(s[3:5])
	if err1 != nil || err2 != nil || row < 1 || col < 1 {
		return Hex{}, fmt.Errorf("board: %q is not a hex such as E3117", s)
	}
	return Hex{Section: Section(s[0]), Row: row, Col: col}, nil
}

// MustParseHex is like ParseHex but panics on error.
func MustParseHex(s string) Hex {
	h, err := ParseHex(s)
	if err != nil {
		panic(err)
	}
	return h
}

// String returns the hex as written on the map, such as "E3117".
func (h Hex) String() string {
	return fmt.Sprintf("%s%02d%02d", h.Section, h.Row, h.Col)
}

// delta holds the Delta hexes, built from deltaRows.
var delta = func() map[Hex]bool {
	m := map[Hex]bool{}
	for _, r := range deltaRows {
		for run := range strings.SplitSeq(r.cols, ", ") {
			lo, hi, isRun := strings.Cut(run, "-")
			if !isRun {
				hi = lo
			}
			first, err1 := strconv.Atoi(lo)
			last, err2 := strconv.Atoi(hi)
			if err1 != nil || err2 != nil || first > last {
				panic(fmt.Sprintf("board: bad Delta columns %q in row %d", run, r.row))
			}
			for col := first; col <= last; col++ {
				m[Hex{Section: 'E', Row: r.row, Col: col}] = true
			}
		}
	}
	return m
}()

// IsDelta reports whether the hex is a Delta hex [8.37]. The Delta is all
// on map section E [29.41].
func IsDelta(h Hex) bool {
	return delta[h]
}

// DeltaColumns returns the columns of the Delta hexes in a row of map
// section E, written as runs such as "15-16, 18-34", or "" if the row has
// none.
func DeltaColumns(row int) string {
	for _, r := range deltaRows {
		if r.row == row {
			return r.cols
		}
	}
	return ""
}

// DeltaRows returns the rows of map section E that have Delta hexes,
// north to south.
func DeltaRows() []int {
	var rows []int
	for _, r := range deltaRows {
		rows = append(rows, r.row)
	}
	return rows
}
