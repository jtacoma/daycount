package engines

import (
	"github.com/jtacoma/go-daycount/daycounts"
)

type Command struct {
	Engine Engine
	Start daycounts.Day
	Count int
}

type Engine interface {
	Run(Command)
}

func Resolve(engineName string) Engine {
	switch engineName {
	case "text":
		return new(TextEngine)
	case "pdf":
		return new(PdfEngine)
	default:
	}
	return nil
}
