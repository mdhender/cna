// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package board

// deltaRows lists the Delta hexes of map section E [8.37, 29.41], row by
// row from north to south, as runs of columns. The hexes were first
// transcribed from Michael Miller's Hex Database (2015) and are checked
// against the official map in features/board/delta.feature.
var deltaRows = []struct {
	row  int
	cols string
}{
	{42, "28"},
	{41, "29"},
	{39, "20, 27-34"},
	{38, "19-33"},
	{37, "15-16, 18-34"},
	{36, "14-33"},
	{35, "14-34"},
	{34, "15-33"},
	{33, "15-34"},
	{32, "14-33"},
	{31, "17-34"},
	{30, "17-33"},
	{29, "22-34"},
	{28, "23-33"},
	{27, "25-34"},
	{26, "24-33"},
	{25, "25-34"},
	{24, "24-33"},
	{23, "25-33"},
	{22, "24-31"},
	{21, "27-31"},
	{20, "27-30"},
	{19, "28-29"},
	{18, "28"},
	{17, "29"},
	{16, "29"},
	{15, "30"},
	{13, "30"},
	{12, "29-30"},
	{11, "30"},
	{10, "29-30"},
	{9, "26, 30"},
	{8, "19, 23-27, 29"},
	{7, "20-27, 29"},
	{6, "23-26, 28"},
	{5, "24-26, 28-29"},
	{4, "23-24, 26-28"},
	{3, "22-24, 27-28"},
	{2, "26-27"},
	{1, "26-27"},
}
