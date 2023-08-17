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

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Application listening on ", portNum)
	http.ListenAndServe(portNum, nil)
}
