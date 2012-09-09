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

package daycount

import (
	"fmt"
)

// Gregorian represents a day in terms of the Gregorian calendar.
//
// The zero value of type Gregorian is January 1st, 1 AD.
type Gregorian struct {
	year       int
	month      int
	dayOfMonth int
	dayOfWeek  int
}

func (self Day) Gregorian() (g Gregorian) {
	g.year = self.time.Year() - 1
	g.month = int(self.time.Month()) - 1
	g.dayOfMonth = self.time.Day() - 1
	g.dayOfWeek = int(self.time.Weekday()) - 4

	// this.dayOfYear = dayOfYear[this.month - 1]
	// + this.dayOfMonth + ((this.isLeapYear && this.month > 2) ? 1 : 0);
	// var a = Math.floor((14 - this.month) / 12);
	// var y = this.year - a;
	// var m = this.month + 12 * a - 2;
	// this.dayOfWeek = ((this.dayOfMonth + y + Math.floor(y / 4)
	// - Math.floor(y / 100) + Math.floor(y / 400) + Math.floor(31 * m / 12))
	// % 7 + 7) % 7 + 1;

	return g
}

func (self Gregorian) DayOfMonth() int {
	return self.dayOfMonth + 1
}

func (self Gregorian) DayOfWeek() int {
	return self.dayOfWeek + 4
}

func (self Gregorian) IsLeapDay() bool {
	return self.IsLeapYear() &&
		(self.month == 1) && (self.dayOfMonth == 28)
}

func (self Gregorian) IsLeapYear() bool {
	year := self.Year()
	return !(year%4 != 0) &&
		(year%100 != 0) ||
		!(year%400 != 0)
}

func (self Gregorian) Month() int {
	return self.month + 1
}

func (self Gregorian) String() string {
	return fmt.Sprintf("%v-%v-%v",
		self.Year(), self.Month(), self.DayOfMonth())
}

func (self Gregorian) Year() int {
	return self.year + 1
}
