package main

import (
	"bookingApp/cmd/pkg/handlers"
	"fmt"
	"net/http"
)

const portNum = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Application listening on ", portNum)
	http.ListenAndServe(portNum, nil)
}
