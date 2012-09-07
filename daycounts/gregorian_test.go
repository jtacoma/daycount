package daycounts

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
