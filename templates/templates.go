package templates

import (
	"text/template"

	"github.com/GibranHL0/devblog/configuration"
)

func Load(config configuration.Configuration) *template.Template{
	templates := template.Must(template.ParseGlob(config.TemplatesPath))

	return templates
}