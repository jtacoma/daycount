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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jtacoma/daycount/pkg/days"
	"github.com/jtacoma/daycount/pkg/views"
)

var (
	today = days.Today()
	start = flag.String("d", today.Gregorian().String(),
		"starting day")
	count = flag.Int("c", 0,
		"count from starting day to ending day")
	format = flag.String("f", "text",
		"output format (text or pdf)")
	template = flag.String("t", "", "file containing a text template.")
)

func getQuery() *views.Query {
	var err error
	var command views.Query
	flag.Parse()
	command.Count = *count
	command.Template = *template
	command.Start, err = days.Parse(*start)
	command.Out = os.Stdout
	if err != nil {
		fmt.Printf("error: failed to parse: %v\n", *start)
		os.Exit(1)
	}
	command.View = views.Resolve(*format)
	if command.View == nil {
		fmt.Printf("error: unsupported format: %v\n", *format)
		os.Exit(1)
	}
	return &command
}

func main() {
	q := getQuery()
	err := q.View.Run(*q)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
