package main

import (
	"ilk/pkg/config"
	"ilk/pkg/handlers"
	"ilk/pkg/render"
	"net/http"
)

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.ListenAndServe(":8080", nil).
}
