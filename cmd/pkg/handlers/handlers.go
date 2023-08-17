package handlers

import (
	"bookingApp/cmd/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
