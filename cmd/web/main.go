package main

import (
	"bookingApp/cmd/pkg/config"
	"bookingApp/cmd/pkg/handlers"
	"bookingApp/cmd/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNum = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false //False for DEV //True for Production

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("Application listening on ", portNum)

	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
