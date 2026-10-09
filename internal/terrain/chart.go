// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package terrain

// Headings are the column headings of the Terrain Effects Chart [8.37].
// The printed chart groups the two CP columns under "CP Cost to Enter or
// Cross" and the three combat columns under "Combat Adjustment".
var Headings = []string{
	"Terrain Type",
	"CP non-Mot",
	"CP Mot",
	"Breakdown Value",
	"Barrage",
	"Anti Armor",
	"Close Assault",
	"Stacking Limit (1)",
}

// Row is one row of the Terrain Effects Chart, with its cells as printed.
// A footnote marker is written after the cell's value in parentheses, as
// in "4 (3)". Text that spans several columns is in the first of them,
// and the cells it covers are empty.
type Row struct {
	// Terrain is the terrain the row describes. It is zero for the
	// Fortifications heading, which describes no terrain itself.
	Terrain Terrain
	Cells   []string
}

func row(t Terrain, cells ...string) Row {
	if len(cells) != len(Headings) {
		panic("terrain: a row of the Terrain Effects Chart needs a cell for each column")
	}
	return Row{Terrain: t, Cells: cells}
}

// Chart holds the Terrain Effects Chart [8.37], as corrected by the errata.
// Each row reads across the chart: Terrain Type, CP to enter or cross for
// non-motorized and motorized units, Breakdown Value, the Barrage,
// Anti Armor and Close Assault adjustments, and Stacking Limit.
//
// Errata applied [8.37]: note 4 moves from Swamp to Major City, and the
// Track row loses its printed CP cost of 1 in both CP columns, leaving
// note 8 to say what a track costs.
//
// Heavy Vegetation's Breakdown Value is printed "3." on the scan; the dot
// is a speck beside the 3.
var Chart = []Row{
	row(Clear, "Clear", "2", "2", "4", "-", "-", "-", "6"),
	row(Gravel, "Gravel", "2", "2", "6", "-", "-", "-", "6"),
	row(SaltMarsh, "Salt Marsh (2)", "3", "2", "6", "-", "-", "R1", "6"),
	row(HeavyVegetation, "Heavy Vegetation", "3", "3", "3", "-", "L1", "L1", "6"),
	row(Rough, "Rough", "3", "4", "8", "L1", "L1", "L2", "6"),
	row(Mountain, "Mountain", "4", "6", "12", "L2", "L2", "L3", "3"),
	row(Delta, "Delta", "2", "4", "2", "-", "-", "-", "6"),
	row(Desert, "Desert", "3", "4 (3)", "24", "-", "-", "-", "6"),
	row(MajorCity, "Major City (4)", "1", "½", "½", "See Fortifications", "", "", "8"),
	row(Swamp, "Swamp", "May enter only on road or railroad", "", "", "-", "-", "-", "6"),
	row(Village, "Village/Bir/Oasis", "Same as terrain in hex for all purposes", "", "", "", "", "", ""),
	row(Railroad, "Railroad (5)", "Same as terrain in hex for all purposes", "", "", "", "", "", ""),
	row(Road, "Road", "1 (6)", "½ (6)", "½ (6)", "-", "-", "-", "5 (7)"),
	row(Track, "Track", "(8)", "(8)", "- (8)", "-", "-", "-", "5 (7)"),
	row(Ridge, "Ridge", "+2", "+4", "2", "-", "L2", "L2", "-"),
	row(UpSlope, "Up Slope", "+2", "+4", "2", "-", "L1", "L2", "-"),
	row(DownSlope, "Down Slope", "+1", "+2", "2", "-", "L1", "R1", "-"),
	row(UpEscarpment, "Up Escarpment", "+6", "P", "-", "-", "P", "L3", "-"),
	row(DownEscarpment, "Down Escarpment", "+4", "+8 (9)", "+6", "-", "L2", "R1", "-"),
	row(Wadi, "Wadi", "+1 (10)", "+4 (10)", "+8", "-", "-", "L1", "-"),
	row(MajorRiver, "Major River", "+8", "P (11)", "-", "-", "-", "L6", "-"),
	row(MinorRiver, "Minor River", "+3", "+6", "+1", "-", "-", "L2", "-"),
	row(0, "Fortifications", "", "", "", "", "", "", ""),
	row(FortificationLevelOne, "Level One", "Same as terrain in hex", "", "", "L1 (12)", "L1", "L2", "-"),
	row(FortificationLevelTwo, "Level Two", "Same as terrain in hex", "", "", "L2 (12)", "L2", "L3", "-"),
	row(FortificationLevelThree, "Level Three", "Same as terrain in hex", "", "", "L2 (12)", "L2", "L4", "-"),
	row(FriendlyMinefield, "Friendly Minefield (13)", "+1", "+4", "0", "-", "L1", "L1", "-"),
	row(EnemyMinefield, "Enemy Minefield (13)", "+4", "+CPA", "+2", "-", "-", "-", "-"),
}
