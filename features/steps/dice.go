// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package steps holds step definitions shared by more than one feature
// area, so a step reads the same wherever it appears.
package steps

import (
	"github.com/cucumber/godog"
	"github.com/mdhender/cna/internal/dice"
)

// Dice registers the steps that fix the dice for a scenario. They set
// *src to a script of the faces given.
func Dice(sc *godog.ScenarioContext, src *dice.Source) {
	sc.Step(`^the die will roll (\d+)$`, func(face int) error {
		*src = dice.NewScript(face)
		return nil
	})
	sc.Step(`^the tens die will roll (\d+) and the ones die will roll (\d+)$`, func(tens, ones int) error {
		*src = dice.NewScript(tens, ones)
		return nil
	})
}
