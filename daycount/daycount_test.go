package daycount

import (
	"testing"
	"time"
)

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
