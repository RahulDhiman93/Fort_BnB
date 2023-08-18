package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf Use for CSRF Token
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad Use for loading Session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
