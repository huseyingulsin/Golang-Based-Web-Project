package render

import (
	"fmt"
	"html/template"
	"ilk/pkg/config"
	"net/http"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// not efficient way to read from memory
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache

	// get requested template from cache

	// render the template bro
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
