package engines

import (
	"fmt"
)

type TextEngine struct {
}

func (*TextEngine) Run(c Command) {
	fmt.Printf("Starting at %v for %v days\n",
		c.Start.Gregorian().String(),
		c.Count)
}
