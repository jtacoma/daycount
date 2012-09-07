package main

import (
	"flag"
	"fmt"
	"github.com/jtacoma/daycount/daycount"
	"github.com/jtacoma/daycount/daycount/engines"
	"os"
)

func getCommand() *engines.Command {
	var err error
	var command engines.Command
	today := daycount.Today()
	start := flag.String("d",
		today.Gregorian().String(), "starting day")
	flag.IntVar(&command.Count, "c",
		0, "count from starting day to ending day")
	format := flag.String("f",
		"text", "output format (text or pdf)")
	flag.Parse()
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
