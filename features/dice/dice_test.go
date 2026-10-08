// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package dice_test

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/cucumber/godog"
	"github.com/mdhender/cna/internal/dice"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: initializeScenario,
		Options: &godog.Options{
			Format:   "progress",
			Paths:    []string{"."},
			Strict:   true,
			Tags:     "~@wip",
			TestingT: t,
		},
	}
	if suite.Run() != 0 {
		t.Fatal("feature tests failed")
	}
}

// world holds the state of one scenario.
type world struct {
	src, second         dice.Source
	result              int
	throw               dice.Throw
	faces               []int
	throws, otherThrows []dice.Throw
}

func initializeScenario(sc *godog.ScenarioContext) {
	w := &world{}
	sc.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
		*w = world{}
		return ctx, nil
	})

	sc.Step(`^the die will roll (\d+)$`, w.dieWillRoll)
	sc.Step(`^the tens die will roll (\d+) and the ones die will roll (\d+)$`, w.diceWillRoll)
	sc.Step(`^a dice roller seeded with (\d+) and (\d+)$`, w.rollerSeeded)
	sc.Step(`^a second dice roller seeded with (\d+) and (\d+)$`, w.secondRollerSeeded)

	sc.Step(`^the player rolls one die$`, w.rollOne)
	sc.Step(`^the player rolls one die (\d+) times$`, w.rollOneTimes)
	sc.Step(`^the player throws two dice$`, w.throwTwo)
	sc.Step(`^the player throws two dice (\d+) times$`, w.throwTwoTimes)
	sc.Step(`^each roller throws two dice (\d+) times$`, w.eachRollerThrows)

	sc.Step(`^the result is (\d+)$`, w.resultIs)
	sc.Step(`^the sequential reading is (\d+)$`, w.sequentialIs)
	sc.Step(`^the sum is (\d+)$`, w.sumIs)
	sc.Step(`^every face from 1 to 6 is rolled$`, w.everyFaceRolled)
	sc.Step(`^no other result is rolled$`, w.noOtherFace)
	sc.Step(`^every sequential reading with both digits from 1 to 6 occurs$`, w.everyReadingOccurs)
	sc.Step(`^no other sequential reading occurs$`, w.noOtherReading)
	sc.Step(`^both rollers throw the same dice in the same order$`, w.sameThrows)
	sc.Step(`^the rollers throw different dice$`, w.differentThrows)
}

func (w *world) dieWillRoll(face int) error {
	w.src = dice.NewScript(face)
	return nil
}

func (w *world) diceWillRoll(tens, ones int) error {
	w.src = dice.NewScript(tens, ones)
	return nil
}

func (w *world) rollerSeeded(seed1, seed2 int) error {
	w.src = dice.NewRoller(uint64(seed1), uint64(seed2))
	return nil
}

func (w *world) secondRollerSeeded(seed1, seed2 int) error {
	w.second = dice.NewRoller(uint64(seed1), uint64(seed2))
	return nil
}

func (w *world) rollOne() error {
	w.result = dice.One(w.src)
	return nil
}

func (w *world) rollOneTimes(n int) error {
	for range n {
		w.faces = append(w.faces, dice.One(w.src))
	}
	return nil
}

func (w *world) throwTwo() error {
	w.throw = dice.Two(w.src)
	return nil
}

func (w *world) throwTwoTimes(n int) error {
	for range n {
		w.throws = append(w.throws, dice.Two(w.src))
	}
	return nil
}

func (w *world) eachRollerThrows(n int) error {
	for range n {
		w.throws = append(w.throws, dice.Two(w.src))
		w.otherThrows = append(w.otherThrows, dice.Two(w.second))
	}
	return nil
}

func (w *world) resultIs(want int) error {
	if w.result != want {
		return fmt.Errorf("result: got %d, want %d", w.result, want)
	}
	return nil
}

func (w *world) sequentialIs(want int) error {
	if got := w.throw.Sequential(); got != want {
		return fmt.Errorf("sequential reading of %+v: got %d, want %d", w.throw, got, want)
	}
	return nil
}

func (w *world) sumIs(want int) error {
	if got := w.throw.Sum(); got != want {
		return fmt.Errorf("sum of %+v: got %d, want %d", w.throw, got, want)
	}
	return nil
}

func (w *world) everyFaceRolled() error {
	for face := 1; face <= 6; face++ {
		if !slices.Contains(w.faces, face) {
			return fmt.Errorf("face %d was never rolled", face)
		}
	}
	return nil
}

func (w *world) noOtherFace() error {
	for _, face := range w.faces {
		if face < 1 || face > 6 {
			return fmt.Errorf("rolled %d, which is not a face of a die", face)
		}
	}
	return nil
}

func (w *world) everyReadingOccurs() error {
	seen := map[int]bool{}
	for _, t := range w.throws {
		seen[t.Sequential()] = true
	}
	for tens := 1; tens <= 6; tens++ {
		for units := 1; units <= 6; units++ {
			if reading := tens*10 + units; !seen[reading] {
				return fmt.Errorf("sequential reading %d never occurred", reading)
			}
		}
	}
	return nil
}

func (w *world) noOtherReading() error {
	for _, t := range w.throws {
		reading := t.Sequential()
		if tens, units := reading/10, reading%10; tens < 1 || tens > 6 || units < 1 || units > 6 {
			return fmt.Errorf("sequential reading %d is not possible with two dice", reading)
		}
	}
	return nil
}

func (w *world) sameThrows() error {
	if !slices.Equal(w.throws, w.otherThrows) {
		return fmt.Errorf("rollers with the same seeds threw different dice")
	}
	return nil
}

func (w *world) differentThrows() error {
	if slices.Equal(w.throws, w.otherThrows) {
		return fmt.Errorf("rollers with different seeds threw the same dice")
	}
	return nil
}
