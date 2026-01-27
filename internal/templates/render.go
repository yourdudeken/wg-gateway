package templates

import (
	"bytes"
	"path/filepath"
	"text/template"
)

func Render(templateName string, data interface{}) ([]byte, error) {
	tmplData, err := Templates.ReadFile(filepath.Join("files", templateName))
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New(templateName).Parse(string(tmplData))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
