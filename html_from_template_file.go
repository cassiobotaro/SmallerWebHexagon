package main

import (
	"bytes"
	"html/template"
)

func HTMLFromTemplateFile(tmplPath string, data any) []byte {
	buf := bytes.NewBuffer([]byte{})
	template.Must(template.ParseFiles(tmplPath)).Execute(buf, data)
	return buf.Bytes()
}
