package main

import (
	"fmt"
	"log"
	"net/http"

	"go-micro.dev/v4/web"
)

func main() {
	// Create a new web service
	service := web.NewService(
		web.Name("helloworld"),
		web.Address(":8080"),
	)

	// Initialize the service
	service.Init()
	log.Printf("Service ready")

	// Set up a route and handler
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		name := "World"
		if params.Has("name") {
			name = params.Get("name")
		}
		fmt.Fprintln(w, "Hello,", name)

	})

	// Assign the handler to the service
	service.Handle("/", http.DefaultServeMux)

	// Start the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
