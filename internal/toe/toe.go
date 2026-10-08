// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package toe describes the weapon systems that make up TOE Strength
// Points. A unit's abilities are those of the TOE Strength Points in it
// [3.5], and each weapon system's ratings come from its nationality's
// Tank and Gun Characteristics Chart [4.47–4.49].
package toe

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mdhender/cna/internal/breakdown"
)

// Nationality is the army a weapon system belongs to.
type Nationality int

const (
	Commonwealth Nationality = iota
	Italian
	German
)

func (n Nationality) String() string {
	return [...]string{"Commonwealth", "Italian", "German"}[n]
}

// Class is the kind of weapon system, as grouped on the charts.
type Class int

const (
	Tank Class = iota
	Artillery
	AntiTank
	AntiAir
)

func (c Class) String() string {
	return [...]string{"tank", "artillery", "anti-tank", "anti-air"}[c]
}

// System is one weapon system: one row of a Tank and Gun Characteristics
// Chart.
type System struct {
	Name            string
	Nationality     Nationality
	Class           Class
	SelfPropelled   bool // a gun on a tank or similar chassis [3.4]
	CPA             CPA
	AntiAir         Rating
	Barrage         Rating
	AntiArmor       Rating
	Vulnerability   Rating
	ArmorProtection Rating
	CloseAssaultOff Rating
	CloseAssaultDef Rating
	FuelRate        Rating
	Breakdown       breakdown.Rating
}

// Systems returns the weapon systems of a nationality and class, in the
// order the chart prints them.
func Systems(n Nationality, c Class) []System {
	var list []System
	for _, s := range systems {
		if s.Nationality == n && s.Class == c {
			list = append(list, s)
		}
	}
	return list
}

// All returns every weapon system, in chart order.
func All() []System {
	return systems[:len(systems):len(systems)]
}

// Lookup finds a weapon system by nationality and name.
func Lookup(n Nationality, name string) (System, bool) {
	for _, s := range systems {
		if s.Nationality == n && s.Name == name {
			return s, true
		}
	}
	return System{}, false
}

// SubjectToBreakdown reports whether TOE Strength Points of this system
// can break down: only tanks and self-propelled guns can [21.11]. Towed
// guns move on their own transport, which never breaks down [3.4].
func (s System) SubjectToBreakdown() bool {
	return s.Class == Tank || s.SelfPropelled
}

// Rating is one numeric cell of a characteristics chart. It may be "-"
// (not applicable or zero), a whole or half number such as "3" or "1½",
// or parenthesized, as in "(2)", which limits when it may be used [3.4].
type Rating struct {
	halves        int // the value in half points
	present       bool
	parenthesized bool
}

// ParseRating parses a chart cell.
func ParseRating(s string) (Rating, error) {
	if s == "-" {
		return Rating{}, nil
	}
	r := Rating{present: true}
	if inner, ok := strings.CutPrefix(s, "("); ok {
		inner, ok = strings.CutSuffix(inner, ")")
		if !ok {
			return Rating{}, fmt.Errorf("toe: bad rating %q", s)
		}
		s, r.parenthesized = inner, true
	}
	whole, half := strings.CutSuffix(s, "½")
	n := 0
	if whole != "" || !half {
		var err error
		if n, err = strconv.Atoi(whole); err != nil || n < 0 {
			return Rating{}, fmt.Errorf("toe: bad rating %q", s)
		}
	}
	r.halves = 2 * n
	if half {
		r.halves++
	}
	return r, nil
}

// Halves returns the rating in half points; "-" is zero.
func (r Rating) Halves() int { return r.halves }

// Present reports whether the chart gives a value, rather than "-".
func (r Rating) Present() bool { return r.present }

// Parenthesized reports whether the chart prints the rating in parentheses.
func (r Rating) Parenthesized() bool { return r.parenthesized }

// String returns the rating as printed.
func (r Rating) String() string {
	if !r.present {
		return "-"
	}
	s := strconv.Itoa(r.halves / 2)
	if r.halves%2 == 1 {
		s = strings.TrimPrefix(s+"½", "0")
	}
	if r.parenthesized {
		s = "(" + s + ")"
	}
	return s
}

// CPA is a Capability Point Allowance. A "0+" CPA marks guns with no
// transport of their own, which can move only on attached trucks.
type CPA struct {
	Points int
	Plus   bool
}

// ParseCPA parses a CPA as printed, such as "25" or "0+".
func ParseCPA(s string) (CPA, error) {
	digits, plus := strings.CutSuffix(s, "+")
	n, err := strconv.Atoi(digits)
	if err != nil || n < 0 {
		return CPA{}, fmt.Errorf("toe: bad CPA %q", s)
	}
	return CPA{Points: n, Plus: plus}, nil
}

// String returns the CPA as printed.
func (c CPA) String() string {
	if c.Plus {
		return strconv.Itoa(c.Points) + "+"
	}
	return strconv.Itoa(c.Points)
}

// row builds a System from a chart row as printed. It panics on a cell it
// can't parse, so a typo in the chart data fails at startup.
func row(n Nationality, c Class, name, cpa, aa, barrage, antiArmor, vul, armor, closeAssault, fuel, bar string) System {
	off, def, ok := strings.Cut(closeAssault, "/")
	if !ok {
		panic(fmt.Sprintf("toe: %s %s: bad Close Assault %q", n, name, closeAssault))
	}
	s := System{Name: name, Nationality: n, Class: c}
	var err error
	if s.CPA, err = ParseCPA(cpa); err != nil {
		panic(fmt.Sprintf("toe: %s %s: %v", n, name, err))
	}
	for _, cell := range []struct {
		dst *Rating
		src string
	}{
		{&s.AntiAir, aa}, {&s.Barrage, barrage}, {&s.AntiArmor, antiArmor},
		{&s.Vulnerability, vul}, {&s.ArmorProtection, armor},
		{&s.CloseAssaultOff, off}, {&s.CloseAssaultDef, def}, {&s.FuelRate, fuel},
	} {
		if *cell.dst, err = ParseRating(cell.src); err != nil {
			panic(fmt.Sprintf("toe: %s %s: %v", n, name, err))
		}
	}
	if s.Breakdown, err = breakdown.ParseRating(bar); err != nil {
		panic(fmt.Sprintf("toe: %s %s: %v", n, name, err))
	}
	return s
}

// sp builds a self-propelled gun's System from its chart row.
func sp(n Nationality, c Class, name, cpa, aa, barrage, antiArmor, vul, armor, closeAssault, fuel, bar string) System {
	s := row(n, c, name, cpa, aa, barrage, antiArmor, vul, armor, closeAssault, fuel, bar)
	s.SelfPropelled = true
	return s
}
