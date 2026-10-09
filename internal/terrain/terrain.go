// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package terrain holds the Terrain Effects Chart [8.37] and works out what
// it costs a unit to move into a hex: the Capability Points it spends and
// the Breakdown Points it picks up.
package terrain

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Terrain is a type of terrain on the Terrain Effects Chart: what fills a
// hex, a feature along a hexside, or a fortification or minefield.
type Terrain int

const (
	Clear Terrain = iota + 1
	Gravel
	SaltMarsh
	HeavyVegetation
	Rough
	Mountain
	Delta
	Desert
	MajorCity
	Swamp
	Village
	Railroad
	Road
	Track
	Ridge
	UpSlope
	DownSlope
	UpEscarpment
	DownEscarpment
	Wadi
	MajorRiver
	MinorRiver
	FortificationLevelOne
	FortificationLevelTwo
	FortificationLevelThree
	FriendlyMinefield
	EnemyMinefield
)

var names = map[Terrain]string{
	Clear:                   "Clear",
	Gravel:                  "Gravel",
	SaltMarsh:               "Salt Marsh",
	HeavyVegetation:         "Heavy Vegetation",
	Rough:                   "Rough",
	Mountain:                "Mountain",
	Delta:                   "Delta",
	Desert:                  "Desert",
	MajorCity:               "Major City",
	Swamp:                   "Swamp",
	Village:                 "Village/Bir/Oasis",
	Railroad:                "Railroad",
	Road:                    "Road",
	Track:                   "Track",
	Ridge:                   "Ridge",
	UpSlope:                 "Up Slope",
	DownSlope:               "Down Slope",
	UpEscarpment:            "Up Escarpment",
	DownEscarpment:          "Down Escarpment",
	Wadi:                    "Wadi",
	MajorRiver:              "Major River",
	MinorRiver:              "Minor River",
	FortificationLevelOne:   "Level One Fortification",
	FortificationLevelTwo:   "Level Two Fortification",
	FortificationLevelThree: "Level Three Fortification",
	FriendlyMinefield:       "Friendly Minefield",
	EnemyMinefield:          "Enemy Minefield",
}

// All lists every type of terrain, in the order of the chart.
var All = func() []Terrain {
	var all []Terrain
	for t := Clear; t <= EnemyMinefield; t++ {
		all = append(all, t)
	}
	return all
}()

func (t Terrain) String() string {
	if name, ok := names[t]; ok {
		return name
	}
	return fmt.Sprintf("Terrain(%d)", int(t))
}

// Parse reads a terrain by its name, such as "Heavy Vegetation".
func Parse(name string) (Terrain, error) {
	for _, t := range All {
		if strings.EqualFold(names[t], name) {
			return t, nil
		}
	}
	return 0, fmt.Errorf("terrain: unknown terrain %q", name)
}

// fillsHex reports whether t is a terrain that fills a whole hex.
func (t Terrain) fillsHex() bool {
	return Clear <= t && t <= Swamp
}

// isHexside reports whether t is a feature along a hexside.
func (t Terrain) isHexside() bool {
	return Ridge <= t && t <= MinorRiver
}

// Mover is the kind of unit that is moving. Most vehicles are simply
// motorized, but the chart's notes treat a few kinds differently.
type Mover int

const (
	NonMotorized Mover = iota
	Motorized
	LightTrucks
	MotorcycleInfantry
	MotorcycleRecce
	Recce // recce-type units other than motorcycle recce
)

// vehicle reports whether the mover uses the Mot column and picks up
// Breakdown Points [21.21].
func (m Mover) vehicle() bool {
	return m != NonMotorized
}

// Route is the way a unit moves into a hex.
type Route int

const (
	// Across means moving through the terrain of the hex and hexside.
	Across Route = iota
	// AlongRoad means moving from a road hex through a road hexside into
	// a connecting road hex [8.33].
	AlongRoad
	// AlongTrack means the same along a track [8.33].
	AlongTrack
)

// Move is what it costs a unit to move into one hex.
type Move struct {
	CP        float64 // Capability Points spent
	Breakdown float64 // Breakdown Points picked up
}

// ErrProhibited is returned when a unit may not make a move.
var ErrProhibited = errors.New("prohibited")

// ErrRuling is returned for a move whose cost waits on an Open ruling in
// RULINGS.md.
var ErrRuling = errors.New("waits on a ruling")

// cost is one cost cell of the chart, read for movement.
type cost struct {
	points     float64
	prohibited bool
}

// effects is what the chart says about moving into or across a terrain.
type effects struct {
	nonMot, mot, breakdown cost
}

// moving holds the chart's movement columns for every terrain that fills
// a hex or lies along a hexside, and for roads.
var moving = func() map[Terrain]effects {
	m := map[Terrain]effects{}
	for _, r := range Chart {
		if !r.Terrain.fillsHex() && !r.Terrain.isHexside() && r.Terrain != Road {
			continue
		}
		if r.Terrain == Swamp {
			continue // it has no costs; see Enter
		}
		m[r.Terrain] = effects{
			nonMot:    mustCost(r.Cells[1]),
			mot:       mustCost(r.Cells[2]),
			breakdown: mustCost(r.Cells[3]),
		}
	}
	return m
}()

// mustCost reads a movement cell of the chart: a number of points, with or
// without a sign, ½, a dash for none, or P for prohibited. A footnote marker
// is dropped; Enter applies the notes.
func mustCost(cell string) cost {
	value, _, _ := strings.Cut(cell, " (")
	switch value {
	case "-":
		return cost{}
	case "P":
		return cost{prohibited: true}
	}
	value = strings.TrimPrefix(value, "+")
	whole, half := strings.CutSuffix(value, "½")
	var points float64
	if whole != "" {
		n, err := strconv.Atoi(whole)
		if err != nil {
			panic(fmt.Sprintf("terrain: can't read the chart cell %q", cell))
		}
		points = float64(n)
	}
	if half {
		points += 0.5
	}
	return cost{points: points}
}

func (e effects) cp(m Mover) cost {
	if m.vehicle() {
		return e.mot
	}
	return e.nonMot
}

// Step is one hex of a move: the terrain of the hex entered, the features
// of the hexside crossed to enter it, and the route taken.
type Step struct {
	Hex      Terrain
	Hexsides []Terrain
	Route    Route
}

// Path returns what it costs a unit to make a move of several hexes: the
// Capability Points spent and the Breakdown Points picked up along the way
// [8.31, 21.23].
func Path(m Mover, steps []Step) (Move, error) {
	var total Move
	for i, s := range steps {
		mv, err := Enter(m, s.Hex, s.Hexsides, s.Route)
		if err != nil {
			return Move{}, fmt.Errorf("hex %d of the move: %w", i+1, err)
		}
		total.CP += mv.CP
		total.Breakdown += mv.Breakdown
	}
	return total, nil
}

// Enter returns what it costs a unit to move into a hex of the given
// terrain, crossing a hexside with the given features [8.31, 8.37].
// A non-motorized unit picks up no Breakdown Points [21.21].
func Enter(m Mover, hex Terrain, hexsides []Terrain, route Route) (Move, error) {
	if !hex.fillsHex() {
		return Move{}, fmt.Errorf("terrain: %s does not fill a hex", hex)
	}
	for _, side := range hexsides {
		if !side.isHexside() {
			return Move{}, fmt.Errorf("terrain: %s is not a hexside feature", side)
		}
	}

	// Light Trucks and motorcycle units never enter Desert, by any route
	// [8.37 note 3, 8.45].
	if hex == Desert && (m == LightTrucks || m == MotorcycleInfantry || m == MotorcycleRecce) {
		return Move{}, fmt.Errorf("%s may not enter Desert: %w", m, ErrProhibited)
	}

	switch route {
	case AlongRoad:
		return alongRoad(m, hex, hexsides)
	case AlongTrack:
		// The chart, its errata and Cases 8.33 and 8.46 disagree on what
		// a track costs.
		return Move{}, fmt.Errorf("moving along a track: R-007 %w", ErrRuling)
	}
	return across(m, hex, hexsides)
}

// across prices a move through the terrain of the hex and hexside.
func across(m Mover, hex Terrain, hexsides []Terrain) (Move, error) {
	// A Swamp may be entered only on a road or railroad [8.37].
	if hex == Swamp {
		return Move{}, fmt.Errorf("Swamp off a road: %w", ErrProhibited)
	}
	// Vehicles other than Light Trucks, recce and motorcycle infantry
	// may not move into a Salt Marsh off a road or track [8.37 note 2, 8.44].
	if hex == SaltMarsh && m == Motorized {
		return Move{}, fmt.Errorf("%s into a Salt Marsh off a road or track: %w", m, ErrProhibited)
	}

	e := moving[hex]
	var mv Move
	mv.CP = e.cp(m).points
	if m.vehicle() {
		mv.Breakdown = e.breakdown.points
	}
	for _, side := range hexsides {
		e := moving[side]
		if e.cp(m).prohibited {
			return Move{}, fmt.Errorf("%s across %s: %w", m, side, ErrProhibited)
		}
		// Vehicles go down an escarpment only on a track [8.37 note 9, 8.42].
		if side == DownEscarpment && m.vehicle() {
			return Move{}, fmt.Errorf("%s down an escarpment off a track: %w", m, ErrProhibited)
		}
		mv.CP += e.cp(m).points
		if m.vehicle() {
			mv.Breakdown += e.breakdown.points
		}
	}
	return mv, nil
}

// alongRoad prices a move along a road. The road's costs replace those of
// the terrain in the hex [8.33], and it cancels the costs of the hexside
// it crosses [8.37 note 6], so a motorized unit may cross a Major River
// on a road [8.37 note 11].
func alongRoad(m Mover, hex Terrain, hexsides []Terrain) (Move, error) {
	for _, side := range hexsides {
		if side == UpEscarpment || side == DownEscarpment {
			// The chart lets a road cancel hexside costs, but Case 8.42
			// keeps vehicles off escarpments except down a track.
			return Move{}, fmt.Errorf("a road across an escarpment: R-008 %w", ErrRuling)
		}
	}
	if hex == SaltMarsh && m == Motorized {
		// Case 8.44 lets vehicles into a Salt Marsh on a road; note 2 of
		// the chart says only on a track.
		return Move{}, fmt.Errorf("%s along a road into a Salt Marsh: R-008 %w", m, ErrRuling)
	}
	e := moving[Road]
	mv := Move{CP: e.cp(m).points}
	if m.vehicle() {
		mv.Breakdown = e.breakdown.points
	}
	return mv, nil
}

func (m Mover) String() string {
	switch m {
	case NonMotorized:
		return "non-motorized"
	case Motorized:
		return "motorized"
	case LightTrucks:
		return "Light Trucks"
	case MotorcycleInfantry:
		return "motorcycle infantry"
	case MotorcycleRecce:
		return "motorcycle recce"
	case Recce:
		return "recce"
	}
	return fmt.Sprintf("Mover(%d)", int(m))
}
