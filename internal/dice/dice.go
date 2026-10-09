// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package dice rolls the six-sided dice the game uses.
//
// The rules read dice three ways: one die (1 to 6), two dice added
// together (2 to 12), and two dice read sequentially (11 to 66). For a
// sequential read, one die is designated the tens die and the other the
// ones die before the throw. The reading is the tens die then the ones
// die, whichever shows more, so all 36 combinations are possible and
// equally likely [3.1]. A single throw of two dice may be read both
// sequentially and as a sum [15.73].
package dice

import (
	"fmt"
	"math/rand/v2"
)

// Source supplies die faces. Every random outcome in the engine draws
// from a Source, so a game can be replayed from its seed.
type Source interface {
	// Face returns the face of the next die rolled, from 1 to 6.
	Face() int
}

// Roller is a seeded Source. Rollers created with the same seeds roll
// the same faces in the same order.
type Roller struct {
	rng *rand.Rand
}

// NewRoller returns a Roller seeded with seed1 and seed2.
func NewRoller(seed1, seed2 uint64) *Roller {
	return &Roller{rng: rand.New(rand.NewPCG(seed1, seed2))}
}

// Face implements Source.
func (r *Roller) Face() int {
	return r.rng.IntN(6) + 1
}

// Script is a Source that returns preset faces in order. Tests use it to
// fix the outcome of a roll.
type Script struct {
	faces []int
}

// NewScript returns a Script that rolls faces in order. It panics if a
// face is not from 1 to 6.
func NewScript(faces ...int) *Script {
	for _, face := range faces {
		if face < 1 || face > 6 {
			panic(fmt.Sprintf("dice: scripted face %d is not from 1 to 6", face))
		}
	}
	return &Script{faces: faces}
}

// Add appends faces to the end of the script. It panics if a face is not
// from 1 to 6.
func (s *Script) Add(faces ...int) {
	s.faces = append(s.faces, NewScript(faces...).faces...)
}

// Remaining returns the number of faces not yet rolled.
func (s *Script) Remaining() int {
	return len(s.faces)
}

// Face implements Source. It panics when the script runs out of faces.
func (s *Script) Face() int {
	if len(s.faces) == 0 {
		panic("dice: script has no faces left")
	}
	face := s.faces[0]
	s.faces = s.faces[1:]
	return face
}

// Throw is one throw of two dice.
type Throw struct {
	Tens int // face of the die designated as the tens digit
	Ones int // face of the die designated as the ones digit
}

// One rolls one die.
func One(src Source) int {
	return src.Face()
}

// Two throws two dice. The tens die is drawn from src first.
func Two(src Source) Throw {
	tens := src.Face()
	return Throw{Tens: tens, Ones: src.Face()}
}

// Sum reads the throw by adding the dice, giving 2 to 12.
func (t Throw) Sum() int {
	return t.Tens + t.Ones
}

// Sequential reads the throw as a two-digit number, the tens die then
// the ones die, giving 11 to 66 [3.1].
func (t Throw) Sequential() int {
	return t.Tens*10 + t.Ones
}
