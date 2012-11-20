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
	"bytes"
	"os"
	"strings"
	"text/template"

	"bitbucket.org/zombiezen/gopdf/pdf"

	"github.com/jtacoma/daycount/pkg/days"
)

type PdfView struct {
	options  PdfViewOptions
	template *template.Template
}

type PdfViewOptions struct {
	FontName   string
	FontSize   pdf.Unit
	Margin     pdf.Unit
	Padding    pdf.Unit
	PageWidth  pdf.Unit
	PageHeight pdf.Unit
}

var (
	pdfsimple = template.Must(template.New("simple").Parse(
		"{{.Gregorian}}"))
)

func (v *PdfView) Set(options *PdfViewOptions) {
	if options != nil {
		v.options = *options
	}
	if len(v.options.FontName) == 0 {
		v.options.FontName = pdf.Helvetica
	}
	if v.options.FontSize == 0 {
		v.options.FontSize = 14
	}
	if v.options.Margin == 0 {
		v.options.Margin = 72
	}
	if v.options.Padding == 0 {
		v.options.Padding = 6
	}
	if v.options.PageWidth == 0 {
		v.options.PageWidth = pdf.USLetterWidth
	}
	if v.options.PageHeight == 0 {
		v.options.PageHeight = pdf.USLetterHeight
	}
}

func (v *PdfView) Run(q Query) (err error) {
	v.Set(nil)
	if v.template, err = q.LoadTemplate(pdfsimple); err != nil {
		return err
	}
	doc := pdf.New()
	for _, day := range q.Range() {
		canvas := doc.NewPage(v.options.PageWidth, v.options.PageHeight)
		v.renderPage(canvas, day)
		canvas.Close()
	}
	return doc.Encode(os.Stdout)
}

func (v *PdfView) renderPage(canvas *pdf.Canvas, d days.Day) {
	margin := v.options.Margin
	padding := v.options.Padding
	width := v.options.PageWidth
	height := v.options.PageHeight
	// page border:
	canvas.Translate(0, 0)
	path := new(pdf.Path)
	path.Move(pdf.Point{margin, margin})
	path.Line(pdf.Point{width - margin, margin})
	path.Line(pdf.Point{width - margin, height - margin})
	path.Line(pdf.Point{margin, height - margin})
	path.Line(pdf.Point{margin, margin})
	canvas.Stroke(path)
	// text
	canvas.Translate(margin+padding,
		height-margin-padding-v.options.FontSize/1.2)
	var buf bytes.Buffer
	v.template.Execute(&buf, d)
	text := new(pdf.Text)
	text.SetFont(v.options.FontName, v.options.FontSize)
	for i, line := range strings.Split(buf.String(), "\n") {
		if i > 0 {
			text.NextLine()
		}
		text.Text(line)
	}
	canvas.DrawText(text)
}
