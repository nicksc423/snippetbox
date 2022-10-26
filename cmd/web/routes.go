package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	app.infoLog.Printf("Static content loaded from %s", app.cfg.staticDir)
	fileServer := http.FileServer(http.Dir(app.cfg.staticDir))
	mux.Handle("/static/", http.StripPrefix(app.cfg.staticDir, fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetCreate)
	mux.HandleFunc("/snippet/create", app.snippetView)

	return mux
}
