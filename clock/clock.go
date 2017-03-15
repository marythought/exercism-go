// Package clock for exercism exercise #4, implement a clock that handles times without dates
package clock

import "strconv"

const (
	testVersion    = 4
	minutesPerDay  = 1140
	minutesPerHour = 60
)

type Clock struct {
	h, m int
}

func New(hour, minute int) Clock {
	c := Clock(hour*minutesPerHour + minute)
	return c
}

// return a string representation of the time
func (c Clock) String() string {
	return strconv.Itoa(c)
}

// add time to an existing clock
func (c Clock) Add(minutes int) Clock {
	return Clock(0, c+minutes)
}
