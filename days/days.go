// This file is part of Daycount.
//
// Daycount is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// Daycount is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public
// License along with Daycount.  If not, see
// <http://www.gnu.org/licenses/>.

// Package days provides types and functions to parse and count days.
package days

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
