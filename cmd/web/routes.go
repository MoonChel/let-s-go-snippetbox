package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes(staticDir string) http.Handler {
	standartMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	fileserver := http.FileServer(http.Dir(staticDir))
	mux.Get("/static/", http.StripPrefix("/static", fileserver))

	return standartMiddleware.Then(mux)
}
