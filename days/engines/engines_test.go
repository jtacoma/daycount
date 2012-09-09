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

package engines

import (
	"github.com/jtacoma/daycount/days"
	"testing"
)

func TestRange(t *testing.T) {
	c := new(Command)
	c.Start, _ = days.Parse("2006-01-02")
	c.Count = 3
	r := c.Range()
	if len(r) != 4 {
		t.Fatalf("expected 4 results, got %v", len(r))
	}
	if r[0] != c.Start {
		t.Fatal("expected %v got %v", c.Start, r[0])
	}
	if r[1] != c.Start.Add(1) {
		t.Fatal("expected %v got %v", c.Start.Add(1), r[1])
	}
	if r[2] != c.Start.Add(2) {
		t.Fatal("expected %v got %v", c.Start.Add(2), r[2])
	}
	if r[3] != c.Start.Add(3) {
		t.Fatal("expected %v got %v", c.Start.Add(3), r[3])
	}
}
