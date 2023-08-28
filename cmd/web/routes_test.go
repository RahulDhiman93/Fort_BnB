package main

import (
	"bookingApp/internal/config"
	"fmt"
	"github.com/go-chi/chi"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//Do Nothing test passed
	default:
		t.Error(fmt.Sprintf("type is not CHI.MUX but, its %T", v))
	}
}
