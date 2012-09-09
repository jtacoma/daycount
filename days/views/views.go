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

package views

import (
	"github.com/jtacoma/daycount/days"
)

type Query struct {
	View     View
	Start    days.Day
	Count    int
	Template string
}

type View interface {
	Run(Query) error
}

func Resolve(viewName string) View {
	switch viewName {
	case "text":
		return new(TextView)
	case "pdf":
		return new(PdfView)
	default:
	}
	return nil
}

func (q *Query) Range() []days.Day {
	var length int
	var step int
	switch {
	case q.Count < 0:
		length = -q.Count + 1
		step = -1
	default:
		length = q.Count + 1
		step = 1
	}
	result := make([]days.Day, length)
	result[0] = q.Start
	for i := range result {
		if i > 0 {
			result[i] = result[i-1].Add(step)
		}
	}
	return result
}
