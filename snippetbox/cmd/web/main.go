package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// existed file: show
	// existed folder: no show
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/ui/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Start Server at: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
