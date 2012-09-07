package daycounts

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
