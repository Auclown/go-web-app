package main

import (
	"fmt"
	"net/http"

	"github.com/auclown/go-basic-web/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
