package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new HTTP server.
	server := http.Server{
		Addr: "localhost:8080",
	}

	// Create a handler for the POST /post route.
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		// Do something with the request body.
		fmt.Fprintf(w, "This is the POST /post route.")
	})

	// Create a handler for the GET /posts route.
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		// Do something with the request body.
		fmt.Fprintf(w, "This is the GET /posts route.")
	})

	// Start the HTTP server.
	server.ListenAndServe()
}
