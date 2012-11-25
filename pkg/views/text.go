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
	"text/template"
)

var textsimple = template.Must(
	template.New("simple").Parse("{{range .}}{{.Gregorian}}\n{{end}}"))

type TextView struct {
}

func (v *TextView) Run(q Query) (err error) {
	if t, err := q.LoadTemplate(textsimple); err != nil {
		return err
	} else {
		t.Execute(q.Out, q.Range())
	}
	return nil
}
