package main

import (
	"math"
	"time"
)

var config = struct {
	Intake   int
	Begin    string
	End      string
	Interval time.Duration

	parsedBegin time.Time
	parsedEnd   time.Time

	dividedIntake int
}{
	Intake:   1200,
	Begin:    "09:00:00",
	End:      "17:00:00",
	Interval: time.Hour * 2,
}

func initConfig() {
	begin, err := time.Parse(time.TimeOnly, config.Begin)
	if err != nil {
		panic(err)
	}

	end, err := time.Parse(time.TimeOnly, config.End)
	if err != nil {
		panic(err)
	}

	config.parsedBegin = begin
	config.parsedEnd = end

	diff := time.Duration(end.Hour()-begin.Hour()) * time.Hour
	divide := diff / config.Interval

	config.dividedIntake = int(math.Round(float64(config.Intake) / float64(divide)))
}
