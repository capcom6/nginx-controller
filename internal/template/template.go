package template

import (
	"io"
	tmpl "text/template"
)

type Template struct {
	Template string
}

func (t *Template) Render(c Context, w io.Writer) error {
	prepared, err := tmpl.New("config").Parse(t.Template)
	if err != nil {
		return err
	}

	return prepared.Execute(w, c)
}
