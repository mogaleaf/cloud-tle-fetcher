package util

import (
	"bytes"
	"text/template"
)

func ExecTempl(t *template.Template, model interface{}) (string, error) {
	buf := &bytes.Buffer{}
	err := t.Execute(buf, model)
	return buf.String(), err
}
