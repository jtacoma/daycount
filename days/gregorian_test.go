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
	"fmt"
	"testing"
)

func ExampleDay_Gregorian() {
	day, _ := Parse("2012-12-21")
	fmt.Println(day.Gregorian().String())
	// Output: 2012-12-21
}

func TestGregorian(t *testing.T) {
	day, _ := Parse("2012-12-21")
	if 2012 != day.Gregorian().Year() {
		t.Fatalf("expected 2012, got %v", day.Gregorian().Year())
	}
	if 12 != day.Gregorian().Month() {
		t.Fatalf("expected 12, got %v", day.Gregorian().Month())
	}
	if 21 != day.Gregorian().DayOfMonth() {
		t.Fatalf("expected 21, got %v", day.Gregorian().DayOfMonth())
	}
}
