// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package steps holds step definitions shared by more than one feature
// area, so a step reads the same wherever it appears.
package steps

import (
	"github.com/cucumber/godog"
	"github.com/mdhender/cna/internal/dice"
	"github.com/mdhender/cna/internal/gametime"
)

// Dice registers the steps that fix the dice for a scenario. The faces
// are queued in the order the steps give them, so a scenario can fix the
// dice for more than one roll. If *src is not already a script, it is
// replaced by one.
func Dice(sc *godog.ScenarioContext, src *dice.Source) {
	add := func(faces ...int) error {
		if s, ok := (*src).(*dice.Script); ok {
			s.Add(faces...)
		} else {
			*src = dice.NewScript(faces...)
		}
		return nil
	}
	sc.Step(`^the die will roll (\d+)$`, func(face int) error {
		return add(face)
	})
	sc.Step(`^the tens die will roll (\d+) and the ones die will roll (\d+)$`, func(tens, ones int) error {
		return add(tens, ones)
	})
}

// Time registers the step that sets the game time, written stage/Game-Turn
// as in "the time is 1/31", passing the time to set.
func Time(sc *godog.ScenarioContext, set func(gametime.Time)) {
	sc.Step(`^the time is (\S+)$`, func(s string) error {
		t, err := gametime.Parse(s)
		if err == nil {
			set(t)
		}
		return err
	})
}
