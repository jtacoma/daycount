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
	"bitbucket.org/zombiezen/gopdf/pdf"
	"github.com/jtacoma/daycount/days"
	"os"
)

type PdfView struct {
	options PdfViewOptions
	lineheight pdf.Unit
}

type PdfViewOptions struct {
	FontName string
	FontSize pdf.Unit
}

var (
	margin = pdf.Unit(72)
	padding = pdf.Unit(6)
	pagewidth = pdf.USLetterWidth
	pageheight = pdf.USLetterHeight
)

func (v *PdfView) Run(q Query) error {
	v.Initialize(PdfViewOptions{pdf.Helvetica, 14})
	doc := pdf.New()
	for _, day := range q.Range() {
		canvas := doc.NewPage(pagewidth, pageheight)
		v.renderPage(canvas, day)
		canvas.Close()
	}
	err := doc.Encode(os.Stdout)
	return err
}

func (v *PdfView) Initialize(options PdfViewOptions) {
	v.options = options
	if len(options.FontName) == 0 {
		v.options.FontName = pdf.Helvetica
	}
	if options.FontSize == 0 {
		v.options.FontSize = 14
	}
}

func (v *PdfView) renderPage(canvas *pdf.Canvas, d days.Day) {
	// page border:
	canvas.Translate(0, 0)
	path := new(pdf.Path)
	path.Move(pdf.Point{margin, margin})
	path.Line(pdf.Point{pagewidth-margin, margin})
	path.Line(pdf.Point{pagewidth-margin, pageheight-margin})
	path.Line(pdf.Point{margin, pageheight-margin})
	path.Line(pdf.Point{margin, margin})
	canvas.Stroke(path)
	// text
	canvas.Translate(margin+padding,
		pageheight-margin-padding-v.options.FontSize/1.2)
	text := new(pdf.Text)
	text.SetFont(v.options.FontName, v.options.FontSize)
	text.Text(d.Gregorian().String())
	canvas.DrawText(text)
}
