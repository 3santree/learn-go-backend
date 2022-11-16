package main

import (
	"log"
	"net/http"
)

// handler function
func home(w http.ResponseWriter, r *http.Request) {
	// show 404 if it's not "/"
	if r.URL.Path != "/" {
		log.Print("404 r.URL.Path: ", r.URL.Path)
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("This is home"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Servermux in Go = router
	mux := http.NewServeMux()
	// "/" = "/*"
	// To restrict to only "/", and 404 for "/dir/not/exist"
	// check that in handler
	mux.HandleFunc("/", home)
	mux.HandleFunc("snippet.view/", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Server Listening on: 4000")

	// ListenAndServe(host:port, handler)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
