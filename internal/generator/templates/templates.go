package templates

import (
	"embed"
	"fmt"
	"io"
	"text/template"
)

//go:embed *.go.tpl
var embeddedTemplates embed.FS

const (
	nameDirTpl = "directory.go.tpl"
)

var parsedTemplates *template.Template

func ExecuteDirTemplate(w io.Writer, data any) error {
	return executeTemplate(w, nameDirTpl, data)
}

func executeTemplate(w io.Writer, name string, data any) error {
	err := templates().ExecuteTemplate(w, name, data)
	if err != nil {
		return fmt.Errorf("failed to execute template %s: %w", name, err)
	}

	return nil
}

func templates() *template.Template {
	if parsedTemplates != nil {
		return parsedTemplates
	}

	var err error

	parsedTemplates, err = template.ParseFS(embeddedTemplates, "*.go.tpl")
	if err != nil {
		panic(fmt.Sprintf("Failed to parse embedded templates: %s", err.Error()))
	}

	return parsedTemplates
}
