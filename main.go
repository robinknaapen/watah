package main

import (
	"fmt"
	"time"

	"github.com/robinknaapen/beeep"
)

func main() {
	initConfig()

	t := time.NewTimer(nextIntake())
	for range t.C {
		intake()
		t.Reset(nextIntake())
	}
}

func nextIntake() time.Duration {
	t := getTimes()

	next := t.now.Truncate(time.Hour).Add(config.Interval)
	for next.Before(t.now) {
		next = next.Add(config.Interval)
	}

	if next.After(t.end) {
		next = t.begin.AddDate(0, 0, 1)
	}

	notifyNextIntake(next)
	return time.Until(next)
}

func intake() {
	t := getTimes()
	if t.now.Before(t.begin) || t.now.After(t.end) {
		return
	}

	notifyIntake()
}

func notifyNextIntake(t time.Time) {
	err := beeep.Notify(
		beeep.AppOption(`Watah`),
		beeep.LevelOption(beeep.LevelNormal),
		beeep.MessageOption(
			fmt.Sprintf(`Next intake at %s`, t.Format(time.DateTime)),
		),
	)
	if err != nil {
		panic(err)
	}
}

func notifyIntake() {
	err := beeep.Notify(
		beeep.AppOption(`Watah`),
		beeep.LevelOption(beeep.LevelCritial),
		beeep.MessageOption(
			fmt.Sprintf(`Drink some water (%dml)`, config.dividedIntake),
		),
	)
	if err != nil {
		panic(err)
	}
}
