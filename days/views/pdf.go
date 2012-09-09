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
	"os"
)

type PdfView struct {
}

func (*PdfView) Run(q Query) error {
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
	return err
}
