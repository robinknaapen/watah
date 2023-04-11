package main

import "time"

type Times struct {
	now   time.Time
	today time.Time
	begin time.Time
	end   time.Time
}

func getTimes() Times {
	t := Times{
		now: time.Now(),
	}

	year, month, day := t.now.Date()
	t.today = time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	t.begin = t.today.Add(time.Duration(config.parsedBegin.Hour()) * time.Hour)
	t.end = t.today.Add(time.Duration(config.parsedEnd.Hour()) * time.Hour)

	return t
}
