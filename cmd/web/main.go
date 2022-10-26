package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold app-wide dependencies
type application struct {
	cfg     *config
	errLog  *log.Logger
	infoLog *log.Logger
}

func main() {
	// Load config values
	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "Port to start listening on")
	flag.StringVar(&cfg.staticDir, "static-dir", "/ui.static", "Path to static assets")
	flag.Parse()

	// Create logger
	// Create INFO level logger with format: INFO (tab) DATE TIME
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Create ERROR level logger with format: ERROR (tab) DATE TIME FILE:LINENUMBER
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize our application which contains all dependencies
	app := &application{
		cfg:     &cfg,
		errLog:  errLog,
		infoLog: infoLog,
	}

	// Init new http>server struct to use custom errorLog
	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", cfg.addr)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
