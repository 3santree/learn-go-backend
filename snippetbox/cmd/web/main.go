package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// -flag --flag -flag=x
	// -flag x  // non-boolean flags only
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP netword address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assests")
	flag.Parse()

	// logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// shortfile will log line number of the error
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()

	// existed file: show
	// existed folder: no show
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/ui/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Start Server at %s", cfg.addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
