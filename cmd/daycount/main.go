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
	"github.com/jtacoma/daycount/daycount"
	"github.com/jtacoma/daycount/daycount/engines"
	"os"
)

var (
	today = daycount.Today()
	start = flag.String("d", today.Gregorian().String(),
		"starting day")
	count = flag.Int("c", 0,
		"count from starting day to ending day")
	format = flag.String("f", "text",
		"output format (text or pdf)")
)

func getCommand() *engines.Command {
	var err error
	var command engines.Command
	flag.Parse()
	command.Count = *count
	command.Start, err = daycount.Parse(*start)
	if err != nil {
		fmt.Printf("error: failed to parse: %v\n", *start)
		os.Exit(1)
	}
	command.Engine = engines.Resolve(*format)
	if command.Engine == nil {
		fmt.Printf("error: unsupported format: %v\n", *format)
		os.Exit(1)
	}
	return &command
}

func main() {
	c := getCommand()
	if c != nil {
		c.Engine.Run(*c)
	}
}
