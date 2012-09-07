package engines

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"fmt"
	"os"
)

type PdfEngine struct {
}

func (*PdfEngine) Run(c Command) {
	doc := pdf.New()
	canvas := doc.NewPage(pdf.USLetterWidth, pdf.USLetterHeight)
	canvas.Translate(0, 0)
	path := new(pdf.Path)
	path.Move(pdf.Point{0, 0})
	path.Line(pdf.Point{100, 0})
	canvas.Stroke(path)
	text := new(pdf.Text)
	text.SetFont(pdf.Helvetica, 14)
	text.Text("Hello, World!")
	canvas.DrawText(text)
	canvas.Close()
	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
