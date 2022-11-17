package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	// Handle Get Parameter
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// fmt.Fprintf(w io.Writer, format string, a ...any) (n int, err error)
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// Must before Write()/WriteHeader(), otherwise no effect
		w.Header().Set("Allow", http.MethodPost)
		// Custom header, avoid canonicalization():
		//			“converts the first letter and any letter following a hyphen to upper case,
		// 			and the rest of the letters to lowercase.”
		w.Header()["custom-header"] = []string{"1; mode=block"}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		// Header, If not set, default as 200
		// Only can be set once in a handler
		// w.WriteHeader(405)
		// Body
		// w.Write([]byte("Method not Allowed, You idoit"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Servermux in Go = router
	mux := http.NewServeMux()
	// "/" = "/*"
	// To restrict to only "/", and 404 for "/dir/not/exist"
	// check that in handler
	mux.HandleFunc("/", home)
	// vhost
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Server Listening on: 4000")

	// ListenAndServe(host:port, handler)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
