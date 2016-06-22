package main

// Pulled a large chunk of this from rest5 stuff i did in go-snippets

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"bitbucket.org/go-snippets/rest/rest5/lib"
)

// Bellow we register the required routes in main and link them to handlers
// The handlers are self contained in their own packages with their required vars
func main() {
	r := mux.NewRouter()
	// Add routes below and point to new handlers
	r.HandleFunc("/", Handlers.HomeHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}