package templates

import (
	"embed"
	"fmt"
	"text/template"
)

//go:embed *.go.tpl
var embeddedTemplates embed.FS

const (
	TemplateNameDir = "directory.go.tpl"
)

func ParseEmbeddedTemplates() *template.Template {
	tmpl, err := template.ParseFS(embeddedTemplates, "*.go.tpl")
	if err != nil {
		panic(fmt.Sprintf("Failed to parse embedded templates: %s", err.Error()))
	}

	return tmpl
}
