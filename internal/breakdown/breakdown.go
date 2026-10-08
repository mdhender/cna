// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package breakdown determines how many vehicles break down when a unit
// stops moving [21.0].
//
// A unit accumulates Breakdown Points as it moves. When it stops, the
// total picks a column of the Breakdown Table [21.38], the column shifts
// for the vehicles' Breakdown Adjustment Rating and the weather, and two
// dice read sequentially give the percentage of TOE Strength Points that
// break down [21.34] (see ruling R-001).
package breakdown

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mdhender/cna/internal/dice"
)

// Column is a column of the Breakdown Table, from 0 (0...3 points) to
// 8 (71+ points).
type Column int

// NumColumns is the number of columns in the Breakdown Table.
const NumColumns = 9

// Percents lists the results of the Breakdown Table, one per row.
var Percents = [...]int{0, 10, 25, 33, 50, 75}

// rollRange is a range of sequential dice readings. The zero value means
// the column has no entry for that row.
type rollRange struct{ lo, hi int }

// columns holds the Breakdown Table [21.38], with the labels of its
// columns and, for each column, the range of rolls that gives each
// percentage in Percents.
var columns = [NumColumns]struct {
	label string
	max   int // the most points in the column; the last column has no limit
	rows  [len(Percents)]rollRange
}{
	{"0...3", 3, [...]rollRange{{11, 66}, {}, {}, {}, {}, {}}},
	{"4...10", 10, [...]rollRange{{11, 42}, {43, 64}, {65, 65}, {66, 66}, {}, {}}},
	{"11...20", 20, [...]rollRange{{11, 32}, {33, 62}, {63, 64}, {65, 65}, {66, 66}, {}}},
	{"21...30", 30, [...]rollRange{{11, 26}, {31, 55}, {56, 62}, {63, 65}, {66, 66}, {}}},
	{"31...40", 40, [...]rollRange{{11, 23}, {24, 53}, {54, 61}, {62, 64}, {65, 66}, {}}},
	{"41...50", 50, [...]rollRange{{11, 16}, {21, 46}, {51, 56}, {61, 63}, {64, 65}, {66, 66}}},
	{"51...60", 60, [...]rollRange{{11, 14}, {15, 42}, {43, 54}, {55, 63}, {64, 65}, {66, 66}}},
	{"61...70", 70, [...]rollRange{{}, {11, 33}, {34, 52}, {53, 62}, {63, 64}, {65, 66}}},
	{"71+", math.MaxInt, [...]rollRange{{}, {11, 25}, {26, 43}, {44, 55}, {56, 63}, {64, 66}}},
}

// ColumnFor returns the column for the points accumulated, before any
// shift. Fractions of a point round up [21.31].
func ColumnFor(points float64) Column {
	p := int(math.Ceil(points))
	for c, col := range columns {
		if p <= col.max {
			return Column(c)
		}
	}
	return NumColumns - 1
}

// Shift moves the column n columns to the right (n > 0) or left (n < 0).
// It stops at 71+ on the right [21.33]. On the left it stops at 0...3,
// where nothing breaks down.
func (c Column) Shift(n int) Column {
	return min(max(c+Column(n), 0), NumColumns-1)
}

// String returns the column's label as printed on the table.
func (c Column) String() string {
	return columns[c].label
}

// Range returns the range of rolls that gives percent in the column, or
// false if the column has no entry for percent.
func (c Column) Range(percent int) (lo, hi int, ok bool) {
	for i, p := range Percents {
		if p == percent {
			r := columns[c].rows[i]
			return r.lo, r.hi, r != rollRange{}
		}
	}
	return 0, 0, false
}

// Percent returns the percentage of TOE Strength Points that break down
// for a sequential dice reading in the column. It panics if the reading
// is not on the table.
func (c Column) Percent(roll int) int {
	for i, r := range columns[c].rows {
		if r.lo <= roll && roll <= r.hi {
			return Percents[i]
		}
	}
	panic(fmt.Sprintf("breakdown: roll %d is not on the %s column", roll, c))
}

// BAR is a Breakdown Adjustment Rating: the number of columns a type of
// vehicle shifts the column, positive to the right [21.13].
type BAR int

// ParseBAR parses a rating written as on the charts: "0", or a number of
// columns followed by L (left) or R (right), such as "2L" or "1R".
func ParseBAR(s string) (BAR, error) {
	if s == "0" {
		return 0, nil
	}
	digits, sign := s, 0
	if n, ok := strings.CutSuffix(s, "L"); ok {
		digits, sign = n, -1
	} else if n, ok := strings.CutSuffix(s, "R"); ok {
		digits, sign = n, 1
	}
	n, err := strconv.Atoi(digits)
	if sign == 0 || err != nil || n < 0 {
		return 0, fmt.Errorf("breakdown: bad Breakdown Adjustment Rating %q", s)
	}
	return BAR(sign * n), nil
}

// String returns the rating as written on the charts.
func (b BAR) String() string {
	switch {
	case b < 0:
		return strconv.Itoa(int(-b)) + "L"
	case b > 0:
		return strconv.Itoa(int(b)) + "R"
	}
	return "0"
}

// Rating is a Breakdown Adjustment Rating as printed on the Tank and Gun
// Characteristics Charts [4.47–4.49]: a column shift, or Exempt for
// vehicles that never break down. A Rating of 0 is not Exempt: those
// vehicles break down without shifting the column. See ruling R-003.
type Rating struct {
	bar    BAR
	exempt bool
}

// Exempt is the Rating of vehicles that never break down, printed "-".
var Exempt = Rating{exempt: true}

// RatingOf returns the Rating for vehicles that break down with BAR b.
func RatingOf(b BAR) Rating {
	return Rating{bar: b}
}

// ParseRating parses a Rating as printed: "-" for Exempt, or a BAR such as
// "0", "2L" or "1R".
func ParseRating(s string) (Rating, error) {
	if s == "-" {
		return Exempt, nil
	}
	b, err := ParseBAR(s)
	return RatingOf(b), err
}

// BAR returns the column shift, and false if the vehicles never break down.
func (r Rating) BAR() (BAR, bool) {
	return r.bar, !r.exempt
}

// String returns the Rating as printed.
func (r Rating) String() string {
	if r.exempt {
		return "-"
	}
	return r.bar.String()
}

// Check is one Breakdown check for one group of TOE Strength Points that
// share a Breakdown Adjustment Rating [21.28].
type Check struct {
	Points    float64 // Breakdown Points accumulated this Operations Stage
	BAR       BAR     // Breakdown Adjustment Rating of the vehicles
	Hot       bool    // the weather is Hot [21.37]
	Sandstorm bool    // a Sandstorm shifts the column [21.37] (see ruling R-002)
}

// Column returns the column the check uses after all shifts, which are
// cumulative [21.32]. It returns false if the unit does not check at all,
// because it has three or fewer points [21.27].
func (c Check) Column() (Column, bool) {
	if c.Points <= 3 {
		return 0, false
	}
	shift := int(c.BAR)
	if c.Hot {
		shift++
	}
	if c.Sandstorm {
		shift++
	}
	return ColumnFor(c.Points).Shift(shift), true
}

// Result is the outcome of a Breakdown check.
type Result struct {
	Checked bool       // false if the unit did not need to check
	Column  Column     // the column used
	Throw   dice.Throw // the dice thrown
	Percent int        // percentage of TOE Strength Points broken down
	Broken  int        // TOE Strength Points broken down
}

// Resolve makes the check for toe TOE Strength Points, throwing the dice
// from src only if the unit needs to check.
func Resolve(c Check, toe int, src dice.Source) Result {
	col, ok := c.Column()
	if !ok {
		return Result{}
	}
	throw := dice.Two(src)
	percent := col.Percent(throw.Sequential())
	return Result{
		Checked: true,
		Column:  col,
		Throw:   throw,
		Percent: percent,
		Broken:  Broken(toe, percent),
	}
}

// Broken returns how many of toe TOE Strength Points break down at
// percent, rounding up. A unit with one TOE Strength Point ignores a 10%
// result [21.35].
func Broken(toe, percent int) int {
	if toe == 1 && percent == 10 {
		return 0
	}
	return (toe*percent + 99) / 100
}
