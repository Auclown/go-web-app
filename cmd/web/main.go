package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/auclown/go-basic-web/pkg/config"
	"github.com/auclown/go-basic-web/pkg/handlers"
	"github.com/auclown/go-basic-web/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache:\n", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
