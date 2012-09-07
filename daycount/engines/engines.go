package engines

import (
	"github.com/jtacoma/daycount/daycount"
)

type Command struct {
	Engine Engine
	Start daycount.Day
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
