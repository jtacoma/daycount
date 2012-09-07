// Package daycount provides types and functions to parse and count days.
package daycount

import (
	"time"
)

// A Day represents a day without regard to exact time, time zone, or at
// what time a day starts or ends in any given count or calendar.
//
// The zero value of type Day is January 1st, 1 AD.
type Day struct {
	time time.Time
}

// Today returns the current day in the local time zone.
func Today() (day Day) {
	day.time = time.Now()
	return day
}

// Parse parses a formatted string and returns the day it represents.
//
// The layout for Gregorian dates is the ISO 8601 Date format (e.g.
// "2006-01-02")
func Parse(value string) (day Day, err error) {
	day.time, err = time.Parse("2006-01-02", value)
	return day, err
}

// Time returns a Time struct representing a moment within this day.
func (self Day) Time() time.Time {
	return self.time
}
