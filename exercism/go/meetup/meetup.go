package meetup

import (
	"time"
)

type WeekSchedule string

const (
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Teenth WeekSchedule = "teenth"
	Last   WeekSchedule = "last"
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	d := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	switch wSched {
	case First:
		d = find(d, wDay, 1)
	case Second:
		d = find(d, wDay, 2)
	case Third:
		d = find(d, wDay, 3)
	case Fourth:
		d = find(d, wDay, 4)
	case Last:
		d = d.AddDate(0, 1, -1)
		for ; d.Weekday() != wDay; d = d.AddDate(0, 0, -1) {
		}
	case Teenth:
		d = d.AddDate(0, 0, 12)
		d = find(d, wDay, 1)
	}

	return d.Day()
}

func find(d time.Time, wDay time.Weekday, times int) time.Time {
	for i := 0; i < times; i++ {
		for ; d.Weekday() != wDay; d = addDay(d) {
		}
		if i+1 < times {
			d = addDay(d)
		}
	}

	return d
}

func addDay(d time.Time) time.Time {
	return d.AddDate(0, 0, 1)
}
