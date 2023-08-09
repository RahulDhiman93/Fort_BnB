package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNum = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home_page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about_page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemp, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemp.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Application listening on ", portNum)
	http.ListenAndServe(portNum, nil)
}
