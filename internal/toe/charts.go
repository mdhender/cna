// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package toe

// systems holds the Tank and Gun Characteristics Charts: Commonwealth
// [4.47], Italian [4.48] and German [4.49], as corrected by the errata.
// Each row reads across the chart: CPA, AA, Barrage, Anti-Armor, Vul,
// Armor Protection, Close Assault Off/Def, Fuel Rate and BAR.
//
// Errata applied: the German Pz III E row is replaced, and the CPAs of
// the 7.62cm Pak(R) (15) and Marder III (25) are restored [4.49]. The
// Commonwealth A9 Cruiser takes its Armor Protection from the Commonwealth
// printing [4.47].
//
// The charts' notes give two guns no anti-armor rating until the first
// Operations Stage of January 1942, marked here with antiArmorFrom
// [4.48, 4.49].
var systems = []System{
	// Commonwealth Tank [4.47]
	row(Commonwealth, Tank, "A9 Cruiser", "25", "1", "-", "3", "-", "1", "3/3", "2", "1R"),
	row(Commonwealth, Tank, "A10 Cruiser", "15", "1", "-", "3", "-", "2", "3/3", "2", "1R"),
	row(Commonwealth, Tank, "A13 Cruiser", "30", "0", "-", "3", "-", "2", "2/3", "2", "1R"),
	row(Commonwealth, Tank, "Churchill II", "15", "1", "-", "6", "-", "7", "5/6", "7", "0"),
	row(Commonwealth, Tank, "Crusader Mk.I", "25", "1", "-", "3", "-", "3", "3/4", "3", "1R"),
	row(Commonwealth, Tank, "Crusader Mk.II", "25", "1", "-", "3", "-", "4", "4/4", "3", "1R"),
	row(Commonwealth, Tank, "Crusader Mk.III", "25", "1", "-", "6", "-", "4", "4/5", "3", "1R"),
	row(Commonwealth, Tank, "Grant M3", "20", "1", "-", "7", "-", "4", "6/5", "7", "0"),
	row(Commonwealth, Tank, "Mark VI Light", "35", "1", "-", "0", "-", "1", "2/2", "1", "0"),
	row(Commonwealth, Tank, "Matilda Mk.II", "15", "0", "-", "3", "-", "6", "3/4", "3", "1R"),
	row(Commonwealth, Tank, "Sherman", "20", "1", "-", "6", "-", "6", "5/5", "6", "0"),
	row(Commonwealth, Tank, "Stuart M3", "35", "1", "-", "4", "-", "3", "4/4", "4", "0"),
	row(Commonwealth, Tank, "Valentine Mk.II", "15", "0", "-", "3", "-", "5", "3/4", "3", "0"),
	row(Commonwealth, Tank, "Scorpion", "25", "0", "-", "0", "-", "7", "0/(2)", "3", "1L"),
	// Commonwealth Artillery [4.47]
	row(Commonwealth, Artillery, "18-pounder Gun", "15", "-", "7", "(2)", "4", "-", "0/1", "1", "-"),
	row(Commonwealth, Artillery, "18/25-pounder Gun", "15", "-", "8", "2", "5", "-", "1/1", "1", "-"),
	row(Commonwealth, Artillery, "25-pounder Gun", "15", "-", "8", "5", "6", "-", "1/1", "1", "-"),
	row(Commonwealth, Artillery, "4.5\" Gun", "15", "-", "11", "1", "9", "-", "1/0", "1", "-"),
	row(Commonwealth, Artillery, "5.5\" Gun/Howitzer", "15", "-", "15", "1", "7", "-", "1/0", "1", "-"),
	row(Commonwealth, Artillery, "60-pounder Gun", "15", "-", "12", "1", "7", "-", "1/0", "1", "-"),
	sp(Commonwealth, Artillery, "SP 25-pounder Gun", "15", "-", "8", "4", "3", "5", "2/1", "3", "2R"),
	row(Commonwealth, Artillery, "3.7\" Howitzer", "15", "-", "7", "0", "3", "-", "0/0", "1", "-"),
	row(Commonwealth, Artillery, "4.5\" Howitzer", "15", "-", "9", "0", "3", "-", "1/0", "1", "-"),
	row(Commonwealth, Artillery, "6\" Howitzer", "15", "-", "15", "1", "5", "-", "1/0", "1", "-"),
	row(Commonwealth, Artillery, "155mm Howitzer", "15", "-", "15", "1", "6", "-", "1/0", "1", "-"),
	sp(Commonwealth, Artillery, "105mm SP Howitzer", "20", "1", "9", "2", "4", "4", "2/3", "3", "0"),
	// Commonwealth AntiTank [4.47]
	row(Commonwealth, AntiTank, "2-pounder", "15", "-", "-", "4", "2", "-", "1/1", "1", "-"),
	row(Commonwealth, AntiTank, "6-pounder", "15", "-", "-", "7", "2", "-", "1/1", "1", "-"),
	row(Commonwealth, AntiTank, "17-pounder", "15", "-", "-", "13", "2", "-", "1/1", "1", "-"),
	sp(Commonwealth, AntiTank, "SP 6-pounder", "20", "-", "-", "7", "2", "1", "2/2", "1", "1L"),
	// Commonwealth AntiAir [4.47]
	row(Commonwealth, AntiAir, "Light AA (Bofors 40mm)", "15", "1", "-", "(1)", "2", "-", "(1)/(1)", "1", "-"),
	row(Commonwealth, AntiAir, "Heavy AA (3.7\")", "15", "4", "-", "(7)", "2", "-", "(1)/(1)", "1", "-"),
	// Italian Tank [4.48]
	row(Italian, Tank, "CV 33(L3/35)", "25", "0", "-", "0", "-", "1", "1/2", "1", "2R"),
	row(Italian, Tank, "L 6/40", "25", "0", "-", "1", "-", "2", "2/2", "2", "0"),
	row(Italian, Tank, "M 11/39", "20", "1", "-", "2", "-", "2", "3/3", "2", "1R"),
	row(Italian, Tank, "M 13/40", "20", "1", "-", "3", "-", "3", "3/3", "2", "1R"),
	row(Italian, Tank, "M14/41", "20", "1", "-", "3", "-", "3", "3/3", "2", "0"),
	// Italian Artillery [4.48]
	row(Italian, Artillery, "65/17 Gun", "15", "-", "5", "0", "3", "-", "1/1", "1", "-"),
	row(Italian, Artillery, "75/18 Gun-Howitzer", "15", "-", "6", "0", "5", "-", "1/0", "1", "-"),
	sp(Italian, Artillery, "75/18 Gun", "20", "-", "6", "6", "4", "3", "3/3", "2", "1R"),
	antiArmorFrom("1/63", row(Italian, Artillery, "75/27 Gun", "15", "-", "6", "2", "4", "-", "1/1", "1", "-")),
	row(Italian, Artillery, "100/17 Howitzer", "15", "-", "8", "0", "5", "-", "1/0", "1", "-"),
	row(Italian, Artillery, "105/28 Gun", "15", "-", "9", "1", "7", "-", "1/1", "1", "-"),
	row(Italian, Artillery, "149/13 Howitzer", "15", "-", "15", "0", "5", "-", "1/0", "1", "-"),
	row(Italian, Artillery, "ParaArt", "15", "-", "2", "2", "1", "-", "1/1", "1", "-"),
	row(Italian, Artillery, "149mm Vichy French", "15", "-", "10", "1", "5", "-", "0/1", "1", "-"),
	row(Italian, Artillery, "155mm Rimhailo(Fr.)", "15", "-", "13", "0", "3", "-", "0/0", "1", "-"),
	// Italian AntiTank [4.48]
	row(Italian, AntiTank, "47/32 Mod. 37", "15", "-", "-", "4", "2", "-", "1/1", "1", "-"),
	// Italian AntiAir [4.48]
	row(Italian, AntiAir, "Light (20mm M/35 Breda)", "15", "1", "-", "2", "2", "-", "(1)/(1)", "1", "-"),
	row(Italian, AntiAir, "'I' Light", "0+", "1", "-", "2", "2", "-", "(1)/(1)", "1", "-"),
	row(Italian, AntiAir, "Heavy (75/46 Mod. 34)", "15", "2", "-", "(5)", "2", "-", "(1)/(1)", "1", "-"),
	row(Italian, AntiAir, "'I' Heavy", "0+", "2", "-", "(5)", "2", "-", "(1)/(1)", "0", "-"),
	row(Italian, AntiAir, "Heavy (90/53)", "15", "3", "-", "9", "2", "-", "(1)/(1)", "1", "-"),
	// German Tank [4.49]
	row(German, Tank, "PZ I", "24", "1", "-", "0", "-", "1", "2/2", "2", "0"),
	row(German, Tank, "Pz II", "35", "0", "-", "1", "-", "2", "2/3", "2", "0"),
	row(German, Tank, "Pz III E", "25", "1", "-", "4", "-", "3", "4/4", "3", "0"),
	row(German, Tank, "Pz III H", "25", "1", "-", "4", "-", "4", "4/4", "3", "0"),
	row(German, Tank, "Pz III J(Special)", "26", "1", "-", "6", "-", "4", "7/5", "3", "0"),
	row(German, Tank, "Pz IV D", "25", "1", "-", "6", "-", "3", "5/4", "4", "0"),
	row(German, Tank, "Pz IV E", "25", "1", "-", "6", "-", "3", "5/5", "4", "0"),
	row(German, Tank, "Pz IV F2 (Special)", "25", "1", "-", "8", "-", "4", "6/5", "4", "0"),
	// German Artillery [4.49]
	antiArmorFrom("1/63", row(German, Artillery, "7.5 cm(IG18) Light Infantry Gun", "15", "-", "6", "3", "2", "-", "1/1", "1", "-")),
	row(German, Artillery, "10.5cm(K18) Medium Gun", "15", "-", "9", "1", "9", "-", "1/1", "1", "-"),
	row(German, Artillery, "10.5cm(leFH18) Light Field Howitzer", "15", "-", "9", "2", "5", "-", "1/1", "1", "-"),
	sp(German, Artillery, "10.5cm SP K18 Gun", "15", "-", "9", "2", "4", "2", "2/2", "2", "0"),
	row(German, Artillery, "149mm Vichy French", "15", "-", "10", "1", "5", "-", "0/1", "1", "-"),
	row(German, Artillery, "15cm (sIG33) Medium Infantry Gun", "15", "-", "14", "1", "2", "-", "1/1", "1", "-"),
	row(German, Artillery, "15cm(sFH18) Medium Field Howitzer", "15", "-", "15", "1", "7", "-", "0/1", "1", "-"),
	row(German, Artillery, "15cm(K18) Gun", "15", "-", "15", "1", "11", "-", "1/1", "1", "-"),
	sp(German, Artillery, "15cm SP Gun", "25", "-", "14", "1", "2", "2", "2/2", "2", "0"),
	row(German, Artillery, "155mm (french 1915 short)", "15", "-", "13", "0", "5", "-", "1/1", "1", "-"),
	row(German, Artillery, "17cm (K18) Gun", "15", "-", "15", "0", "13", "-", "1/0", "1", "-"),
	row(German, Artillery, "21cm (mrs18) Howitzer", "15", "-", "18", "0", "8", "-", "1/0", "1", "-"),
	// German AntiTank [4.49]
	row(German, AntiTank, "2.8cm s.Pz.B.41 or 28/20 Pak", "15", "-", "-", "4", "2", "-", "1/1", "1", "-"),
	row(German, AntiTank, "3.7cm Pak 35/36", "15", "-", "-", "2", "1", "-", "1/1", "1", "-"),
	row(German, AntiTank, "5cm Pak 38", "15", "-", "-", "5", "2", "-", "1½/1", "1", "-"),
	row(German, AntiTank, "7.62cm Pak(R)", "15", "-", "-", "9", "2", "-", "1/1", "1", "-"),
	sp(German, AntiTank, "Pzjg 1 (SP)", "30", "-", "-", "4", "2", "1", "2/2", "1", "0"),
	sp(German, AntiTank, "Marder III (SP)", "25", "-", "-", "9", "2", "2", "3/2", "2", "0"),
	// German AntiAir [4.49]
	row(German, AntiAir, "Light (20mm & 37mm Flak)", "15", "1", "-", "(1)", "(1)", "-", "(1)/(1)", "1", "-"),
	row(German, AntiAir, "Heavy (88mm Flak)", "15", "3", "-", "13", "2", "-", "2/1", "1", "-"),
}
