package clock

import "fmt"

type Clock struct {
	time int // in minutes
}

func New(h, m int) Clock {
	minutes := h*60 + m // That's the minute representation of HH:MM
	minutes %= 24 * 60  // We can only store one day worth of minutes, we drop any extra information

	// Negative minutes must be subtracted to the end of the day
	if minutes < 0 {
		minutes = 24*60 + minutes
	}

	return Clock{time: minutes}
}

func (c Clock) Add(m int) Clock {
	return New(c.hours(), c.minutes()+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours(), c.minutes()-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours(), c.minutes())
}

func (c Clock) hours() int {
	return c.time / 60
}

func (c Clock) minutes() int {
	return c.time % 60
}
