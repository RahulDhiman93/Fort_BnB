package main

import (
	"fmt"
	"net/http"
)

func startWebApp() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Application listening on ", portNum)
	http.ListenAndServe(portNum, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home_page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about_page.tmpl")
}
