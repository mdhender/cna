// Copyright (c) 2026 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package calendar_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/mdhender/cna/internal/gametime"
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
	time gametime.Time
	err  error
}

const time = `(\S+/\S+)`

func initializeScenario(sc *godog.ScenarioContext) {
	w := &world{}
	sc.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
		*w = world{}
		return ctx, nil
	})

	sc.Step(`^the time (\S+) is read$`, w.read)
	sc.Step(`^it is Operations Stage (\d+) of Game-Turn (\d+)$`, w.isStageOf)
	sc.Step(`^it is not a time in the campaign$`, w.isNotATime)
	sc.Step(`^the Operations Stage after `+time+` is `+time+`$`, next)
	sc.Step(`^there is no Operations Stage after `+time+`$`, noNext)
	sc.Step(`^`+time+` comes (before|after) `+time+`$`, comes)
	sc.Step(`^the campaign (starts|ends) at `+time+`$`, campaign)
	sc.Step(`^Game-Turn (\d+) is (.+)$`, turnIsDate)
	sc.Step(`^(.+, \d+) is Game-Turn (\d+)$`, dateIsTurn)
}

func (w *world) read(s string) error {
	w.time, w.err = gametime.Parse(s)
	return nil
}

func (w *world) isStageOf(stage, turn int) error {
	if w.err != nil {
		return w.err
	}
	if w.time.Stage != stage || w.time.Turn != turn {
		return fmt.Errorf("got Stage %d of Game-Turn %d, want Stage %d of Game-Turn %d", w.time.Stage, w.time.Turn, stage, turn)
	}
	return nil
}

func (w *world) isNotATime() error {
	if w.err == nil {
		return fmt.Errorf("read as %s, want an error", w.time)
	}
	return nil
}

func parse(s string) gametime.Time {
	return gametime.MustParse(s)
}

func next(from, want string) error {
	got, ok := parse(from).Next()
	if !ok {
		return fmt.Errorf("there is no Operations Stage after %s", from)
	}
	if got.String() != want {
		return fmt.Errorf("after %s: got %s, want %s", from, got, want)
	}
	return nil
}

func noNext(from string) error {
	if got, ok := parse(from).Next(); ok {
		return fmt.Errorf("after %s: got %s, want none", from, got)
	}
	return nil
}

func comes(a, order, b string) error {
	ta, tb := parse(a), parse(b)
	if ok := (order == "before" && ta.Before(tb)) || (order == "after" && ta.After(tb)); !ok {
		return fmt.Errorf("%s does not come %s %s", a, order, b)
	}
	return nil
}

func campaign(which, want string) error {
	got := gametime.Start
	if which == "ends" {
		got = gametime.End
	}
	if got.String() != want {
		return fmt.Errorf("the campaign %s at %s, want %s", which, got, want)
	}
	return nil
}

func turnIsDate(turn int, want string) error {
	if got := gametime.DateOf(turn).String(); got != want {
		return fmt.Errorf("Game-Turn %d: got %s, want %s", turn, got, want)
	}
	return nil
}

func dateIsTurn(date string, want int) error {
	d, err := gametime.ParseDate(date)
	if err != nil {
		return err
	}
	got, ok := gametime.TurnOf(d)
	if !ok {
		return fmt.Errorf("%s is not in the campaign", date)
	}
	if got != want {
		return fmt.Errorf("%s: got Game-Turn %d, want %d", date, got, want)
	}
	return nil
}
