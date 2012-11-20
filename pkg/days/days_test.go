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

package days

import (
	"testing"
	"time"
)

func TestDay_Add(t *testing.T) {
	day, _ := Parse("2006-02-01")
	after, _ := Parse("2006-02-02")
	added := day.Add(1)
	if added != after {
		t.Fatalf("expected 2006-02-02 but got %v", added.Gregorian())
	}
	back := after.Add(-1)
	if back != day {
		t.Fatalf("expected 2006-02-01 but got %v", back.Gregorian())
	}
}

func TestDay_Today(t *testing.T) {
	before := time.Now()
	today := Today()
	after := time.Now()
	matchYear := today.Time().Year() == before.Year() &&
		today.Time().Year() == after.Year()
	matchMonth := today.Time().Month() == before.Month() &&
		today.Time().Month() == after.Month()
	matchDay := today.Time().Day() == before.Day() &&
		today.Time().Day() == after.Day()
	if !(matchYear && matchMonth && matchDay) {
		t.Fatalf("now is not today")
	}
}

func TestDayParse(t *testing.T) {
	day, err := Parse("2012-12-21")
	if err != nil {
		t.Fatalf("couldn't parse date: %v", err)
	}
	match := day.Time().Year() == 2012 &&
		day.Time().Month() == 12 &&
		day.Time().Day() == 21
	if !match {
		t.Fatalf("couldn't parse standard gregorian date")
	}
}
